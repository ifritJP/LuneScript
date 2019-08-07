#include <lunescript.h>

#ifdef __cplusplus
extern "C" {
#endif

    // ========== List ==========
#define lune_mtd_List( OBJ ) \
    ((lune_List_t*)OBJ->val.classVal)->pMtd
    
    struct lune_mtd_List_t {
        lune_del_t * _del;
        lune_gc_t * _gc;
        lune_method_t * insert;
        lune_method_t * refAt;
        lune_method_t * unpack;
        lune_method_t * sort;
    };
    
    extern lune_stem_t * lune_class_List_new( lune_env_t * _pEnv );
    extern lune_stem_t * lune_List_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem );

    extern lune_stem_t * lune_itList_new( lune_env_t * _pEnv, lune_stem_t * pList );
    extern void lune_itList__del( lune_env_t * _pEnv, lune_stem_t * it );
    extern void lune_itList_inc( lune_env_t * _pEnv, lune_stem_t * it );
    extern bool lune_itList_hasNext( lune_env_t * _pEnv, lune_stem_t * it, lune_stem_t ** ppVal );
    

    // ========== Set ==========
#define lune_mtd_Set( OBJ ) \
    ((lune_Set_t*)OBJ->val.classVal)->pMtd
    
    struct lune_mtd_Set_t {
        lune_del_t * _del;
        lune_gc_t * _gc;
        lune_method_t * add;
        lune_method_t * del;
        lune_method_t * has;
        lune_method_t * and_;
        lune_method_t * or_;
        lune_method_t * sub;
        lune_method_t * clone;
        lune_method_t * len;
    };

    
    extern lune_stem_t * lune_class_Set_new( lune_env_t * _pEnv );
    extern lune_stem_t * lune_Set_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem );

    extern lune_stem_t * lune_mtd_Set_createList( lune_env_t * _pEnv, lune_stem_t * pObj );


    extern lune_stem_t * lune_itSet_new( lune_env_t * _pEnv, lune_stem_t * pSet );
    extern void lune_itSet__del( lune_env_t * _pEnv, lune_stem_t * it );
    extern void lune_itSet_inc( lune_env_t * _pEnv, lune_stem_t * it );
    extern bool lune_itSet_hasNext( lune_env_t * _pEnv, lune_stem_t * it, lune_stem_t ** ppVal );


    // ========== Map ==========
#define lune_mtd_Map( OBJ ) \
    ((lune_Map_t*)OBJ->val.classVal)->pMtd
    
    struct lune_mtd_Map_t {
        lune_del_t * _del;
        lune_gc_t * _gc;
        lune_method_t * add;
        lune_method_t * get;
    };

    typedef struct lune_Map_entry_t {
        lune_stem_t * pKey;
        lune_stem_t * pVal;
    } lune_Map_entry_t;

    
    extern lune_stem_t * lune_class_Map_new( lune_env_t * _pEnv );
    extern lune_stem_t * lune_Map_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem );
    extern lune_stem_t * lune_mtd_Map_createKeyList(
        lune_env_t * _pEnv, lune_stem_t * pObj );
    extern lune_stem_t * lune_mtd_Map_add( lune_env_t * _pEnv, lune_stem_t * pObj,
                                           lune_stem_t * pKey, lune_stem_t * pVal );
    extern lune_stem_t * lune_mtd_Map_get( lune_env_t * _pEnv, lune_stem_t * pObj,
                                           lune_stem_t * pKey );

    

    extern lune_stem_t * lune_itMap_new( lune_env_t * _pEnv, lune_stem_t * pMap );
    extern void lune_itMap__del( lune_env_t * _pEnv, lune_stem_t * it );
    extern void lune_itMap_inc( lune_env_t * _pEnv, lune_stem_t * it );
    extern bool lune_itMap_hasNext( lune_env_t * _pEnv, lune_stem_t * it,
                                    lune_Map_entry_t * pEntry );

    

    
#ifdef __cplusplus
}
#endif
