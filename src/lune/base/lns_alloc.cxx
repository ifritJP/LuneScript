#include <lunescript.h>
#include <map>
using namespace std;

class AllocInfo {
public:
    const char * pName;
    int lineNo;
    AllocInfo(const char * pName, int lineNo) {
        this->pName = pName;
        this->lineNo = lineNo;
    }
};

lune_allocator_t lune_createAllocator( void ) {
    return new map<void*,AllocInfo *>();
}

void * _lune_malloc( lune_allocator_t allocateor,
                     int size, const char * pName, int lineNo )
{
    map<void*,AllocInfo *> * pMap = (map<void*,AllocInfo *> *)allocateor;
    void * pAddr = malloc( size );
    (*pMap)[ pAddr ] = new AllocInfo( pName, lineNo );
    return pAddr;
}

void _lune_free( lune_allocator_t allocateor,
                 void * pAddr, const char * pName, int lineNo )
{
    map<void*,AllocInfo *> * pMap = (map<void*,AllocInfo *> *)allocateor;
    map<void*,AllocInfo *>::iterator it = pMap->find( pAddr );
    if ( it != pMap->end() ) {
        delete it->second;
        pMap->erase( pAddr );
    }
    else {
        printf( "error!!\n" );
    }
}

void lune_checkMem( lune_allocator_t allocateor )
{
    map<void*,AllocInfo *> * pMap = (map<void*,AllocInfo *> *)allocateor;
    map<void*,AllocInfo *>::iterator it = pMap->begin();
    map<void*,AllocInfo *>::iterator end = pMap->end();

    for ( ; it != end; it++ ) {
        printf( "allocated = %p,%s,%d\n",
                it->first, it->second->pName, it->second->lineNo );
    }
}
    
