#include <lunescript.h>

#ifdef __cplusplus
extern "C" {
#endif

#define __lune_mtd_List( OBJ ) \
    ((__lune_List_t*)OBJ->val.classVal)->pMtd
    
    struct __mtd_List_t {
        __lune_gc_t * _gc;
        __lune_method_t * insert;
        __lune_method_t * refAt;
        __lune_method_t * unpack;
    };
    
    extern __lune_stem_t * __class_List_new( __lune_env_t * _pEnv );

#ifdef __cplusplus
}
#endif
