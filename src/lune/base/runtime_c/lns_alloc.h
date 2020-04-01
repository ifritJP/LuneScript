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

#ifndef __LUNEALLOCATOR__
#define __LUNEALLOCATOR__

#ifdef __cplusplus
extern "C" {
#endif

#define lns_malloc( ALLOCATOR, SIZE ) _lns_malloc( ALLOCATOR, SIZE, __FILE__, __LINE__ )
#define lns_free( ALLOCATOR, ADDR ) _lns_free( ALLOCATOR, ADDR, __FILE__, __LINE__ )

    
    typedef void * lns_allocator_t;

    extern lns_allocator_t lns_createAllocator( void );
    extern void * _lns_malloc( lns_allocator_t allocateor,
                                int size, const char * pName, int lineNo );
    extern void _lns_free( lns_allocator_t allocateor,
                            void * pAddr, const char * pName, int lineNo );
    extern void lns_checkMem( void );

#ifdef __cplusplus
}
#endif

#endif
