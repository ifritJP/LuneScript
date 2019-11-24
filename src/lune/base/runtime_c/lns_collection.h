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

#ifndef __LNS_COLLECTION__
#define __LNS_COLLECTION__

#ifdef __cplusplus
extern "C" {
#endif

    extern void lns_collection_init();
    

    // ========== List ==========
#define lns_mtd_List( OBJ ) \
    ((lns_List_t*)OBJ->val.classVal)->pMtd
    
    struct lns_mtd_List_t {
        lns_del_t * _del;
        lns_gc_t * _gc;
        lns_method_t * insert;
        lns_method_t * refAt;
        lns_method_t * unpack;
        lns_method_t * sort;
    };
    
    extern lns_any_t * lns_class_List_new( lns_env_t * _pEnv );
    extern lns_any_t * lns_List_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny );


    extern lns_stem_t lns_mtd_List_insert(
        lns_env_t * _pEnv, lns_any_t * pListAny, lns_stem_t pVal );
    extern lns_stem_t lns_mtd_List_remove(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val );
    extern lns_stem_t lns_mtd_List_refAt(
        lns_env_t * _pEnv, lns_any_t * pListAny, int index );
    extern void lns_mtd_List_setAt(
        lns_env_t * _pEnv, lns_any_t * pListAny, int index, lns_stem_t val );
    extern lns_stem_t lns_mtd_List_unpack(
        lns_env_t * _pEnv, lns_any_t * pListAny );
    extern lns_stem_t lns_mtd_List_sort(
        lns_env_t * _pEnv, lns_any_t * pListAny, lns_stem_t pForm );
    extern lns_int_t lns_mtd_List_len( lns_env_t * _pEnv, lns_any_t * pListAny );

    
    extern lns_any_t * lns_itList_new( lns_env_t * _pEnv, lns_any_t * pList );
    extern void lns_itList__del( lns_env_t * _pEnv, lns_any_t * it );
    extern void lns_itList_inc( lns_env_t * _pEnv, lns_any_t * it );
    extern bool lns_itList_hasNext( lns_env_t * _pEnv, lns_any_t * it, lns_stem_t* ppVal );
    

    // ========== Set ==========
#define lns_mtd_Set( OBJ ) \
    ((lns_Set_t*)OBJ->val.classVal)->pMtd
    
    struct lns_mtd_Set_t {
        lns_del_t * _del;
        lns_gc_t * _gc;
        lns_method_t * add;
        lns_method_t * del;
        lns_method_t * has;
        lns_method_t * and_;
        lns_method_t * or_;
        lns_method_t * sub;
        lns_method_t * clone;
        lns_method_t * len;
    };

    
    extern lns_any_t * lns_class_Set_new( lns_env_t * _pEnv );
    extern lns_any_t * lns_Set_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny );

    extern lns_any_t * lns_mtd_Set_createList( lns_env_t * _pEnv, lns_any_t * pObj );

    extern void lns_mtd_Set_add(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val );
    extern void lns_mtd_Set_del(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val );
    extern bool lns_mtd_Set_has(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val );
    extern lns_any_t *lns_mtd_Set_and(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet );
    extern lns_any_t *lns_mtd_Set_or(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet );
    extern lns_any_t * lns_mtd_Set_sub(
        lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet );
    extern lns_any_t * lns_mtd_Set_clone( lns_env_t * _pEnv, lns_any_t * pObj );
    extern lns_int_t lns_mtd_Set_len( lns_env_t * _pEnv, lns_any_t * pObj );
    


    extern lns_any_t * lns_itSet_new( lns_env_t * _pEnv, lns_any_t * pSet );
    extern void lns_itSet__del( lns_env_t * _pEnv, lns_any_t * it );
    extern void lns_itSet_inc( lns_env_t * _pEnv, lns_any_t * it );
    extern bool lns_itSet_hasNext( lns_env_t * _pEnv, lns_any_t * it, lns_stem_t* ppVal );


    // ========== Map ==========
#define lns_mtd_Map( OBJ ) \
    ((lns_Map_t*)OBJ->val.classVal)->pMtd
    
    struct lns_mtd_Map_t {
        lns_del_t * _del;
        lns_gc_t * _gc;
        lns_method_t * add;
        lns_method_t * get;
    };

    typedef struct lns_Map_entry_t {
        lns_stem_t key;
        lns_stem_t val;
    } lns_Map_entry_t;

    
    extern lns_any_t * lns_class_Map_new( lns_env_t * _pEnv );
    extern lns_any_t * lns_Map_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny );
    extern lns_any_t * lns_mtd_Map_createKeyList(
        lns_env_t * _pEnv, lns_any_t * pObj );
    extern void lns_mtd_Map_add( lns_env_t * _pEnv, lns_any_t * pObj,
                                  lns_stem_t pKey, lns_stem_t pVal );
    extern lns_stem_t lns_mtd_Map_get( lns_env_t * _pEnv, lns_any_t * pObj,
                                           lns_stem_t pKey );

    

    extern lns_any_t * lns_itMap_new( lns_env_t * _pEnv, lns_any_t * pMap );
    extern void lns_itMap__del( lns_env_t * _pEnv, lns_any_t * it );
    extern void lns_itMap_inc( lns_env_t * _pEnv, lns_any_t * it );
    extern bool lns_itMap_hasNext( lns_env_t * _pEnv, lns_any_t * it,
                                    lns_Map_entry_t * pEntry );

    

    
#ifdef __cplusplus
}
#endif

#endif

