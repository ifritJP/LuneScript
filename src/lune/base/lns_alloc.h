#ifndef __LUNEALLOCATOR__
#define __LUNEALLOCATOR__

#ifdef __cplusplus
extern "C" {
#endif

    typedef void * lune_allocator_t;

    extern lune_allocator_t lune_createAllocator( void );
    extern void * _lune_malloc( lune_allocator_t allocateor,
                                int size, const char * pName, int lineNo );
    extern void _lune_free( lune_allocator_t allocateor,
                            void * pAddr, const char * pName, int lineNo );
    extern void lune_checkMem( void );

#ifdef __cplusplus
}
#endif

#endif
