/*
MIT License

Copyright (c) 2018,2019 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

#include <lunescript.h>
#include <map>
using namespace std;


//#define DEBUG_LOG


class AllocInfo {
public:
    const char * pName;
    int lineNo;
    AllocInfo(const char * pName, int lineNo) {
        this->pName = pName;
        this->lineNo = lineNo;
    }
};

static map<void*,AllocInfo *> s_map;

lns_allocator_t lns_createAllocator( void ) {
    return &s_map;
}

void * _lns_malloc( lns_allocator_t allocateor,
                     int size, const char * pName, int lineNo )
{
    map<void*,AllocInfo *> * pMap = (map<void*,AllocInfo *> *)allocateor;
    void * pAddr = malloc( size );
    (*pMap)[ pAddr ] = new AllocInfo( pName, lineNo );
#ifdef DEBUG_LOG
    printf( "alloc = %p, %p,%s,%d\n", pMap, pAddr, pName, lineNo );
#endif
    return pAddr;
}

void _lns_free( lns_allocator_t allocateor,
                 void * pAddr, const char * pName, int lineNo )
{
    map<void*,AllocInfo *> * pMap = (map<void*,AllocInfo *> *)allocateor;
    map<void*,AllocInfo *>::iterator it = pMap->find( pAddr );
    if ( it != pMap->end() ) {
#ifdef DEBUG_LOG
        printf( "free = %p, %p,%s,%d:%s:%d\n",
                pMap, pAddr, it->second->pName, it->second->lineNo, pName, lineNo );
#endif
        delete it->second;
        pMap->erase( pAddr );
    }
    else {
        printf( "free error!! -- %p, %p:%s:%d\n", pMap, pAddr, pName, lineNo );
    }
}

void lns_checkMem( void )
{
    map<void*,AllocInfo *> * pMap = &s_map;
    map<void*,AllocInfo *>::iterator it = pMap->begin();
    map<void*,AllocInfo *>::iterator end = pMap->end();

    for ( ; it != end; it++ ) {
        printf( "allocated = %p,%s,%d\n",
                it->first, it->second->pName, it->second->lineNo );
    }
}
    
