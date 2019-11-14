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

#ifndef __LUNE_COLLECTION__
#define __LUNE_COLLECTION__

#ifdef __cplusplus
extern "C" {
#endif

    extern void lune_collection_init();
    

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
    
    extern lune_any_t * lune_class_List_new( lune_env_t * _pEnv );
    extern lune_any_t * lune_List_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny );


    extern lune_stem_t lune_mtd_List_insert(
        lune_env_t * _pEnv, lune_any_t * pListAny, lune_stem_t pVal );
    extern lune_stem_t lune_mtd_List_refAt(
        lune_env_t * _pEnv, lune_any_t * pListAny, int index );
    extern lune_any_t * lune_mtd_List_unpack(
        lune_env_t * _pEnv, lune_any_t * pListAny );
    extern lune_stem_t lune_mtd_List_sort(
        lune_env_t * _pEnv, lune_any_t * pListAny, lune_stem_t pForm );
    extern lune_int_t lune_mtd_List_len( lune_env_t * _pEnv, lune_any_t * pListAny );

    
    extern lune_any_t * lune_itList_new( lune_env_t * _pEnv, lune_any_t * pList );
    extern void lune_itList__del( lune_env_t * _pEnv, lune_any_t * it );
    extern void lune_itList_inc( lune_env_t * _pEnv, lune_any_t * it );
    extern bool lune_itList_hasNext( lune_env_t * _pEnv, lune_any_t * it, lune_stem_t* ppVal );
    

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

    
    extern lune_any_t * lune_class_Set_new( lune_env_t * _pEnv );
    extern lune_any_t * lune_Set_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny );

    extern lune_any_t * lune_mtd_Set_createList( lune_env_t * _pEnv, lune_any_t * pObj );


    extern lune_any_t * lune_itSet_new( lune_env_t * _pEnv, lune_any_t * pSet );
    extern void lune_itSet__del( lune_env_t * _pEnv, lune_any_t * it );
    extern void lune_itSet_inc( lune_env_t * _pEnv, lune_any_t * it );
    extern bool lune_itSet_hasNext( lune_env_t * _pEnv, lune_any_t * it, lune_stem_t* ppVal );


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
        lune_stem_t key;
        lune_stem_t val;
    } lune_Map_entry_t;

    
    extern lune_any_t * lune_class_Map_new( lune_env_t * _pEnv );
    extern lune_any_t * lune_Map_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny );
    extern lune_any_t * lune_mtd_Map_createKeyList(
        lune_env_t * _pEnv, lune_any_t * pObj );
    extern void lune_mtd_Map_add( lune_env_t * _pEnv, lune_any_t * pObj,
                                  lune_stem_t pKey, lune_stem_t pVal );
    extern lune_stem_t lune_mtd_Map_get( lune_env_t * _pEnv, lune_any_t * pObj,
                                           lune_stem_t pKey );

    

    extern lune_any_t * lune_itMap_new( lune_env_t * _pEnv, lune_any_t * pMap );
    extern void lune_itMap__del( lune_env_t * _pEnv, lune_any_t * it );
    extern void lune_itMap_inc( lune_env_t * _pEnv, lune_any_t * it );
    extern bool lune_itMap_hasNext( lune_env_t * _pEnv, lune_any_t * it,
                                    lune_Map_entry_t * pEntry );

    

    
#ifdef __cplusplus
}
#endif

#endif

