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

#include <gc_lunescript.h>
#include <lns_collection.h>
#include <vector>
#include <algorithm>
#include <set>
#include <map>

using namespace std;

static int s_collection_allocNum = 0;

// 一時的に stem から var を使用するためのマクロ。
// stem の find など var でアクセスするため使う。
#define WORK_VAR( STEM ) (lns_var_t){ .stem = STEM }

//#define DEBUG_LOG


void lns_collection_init( void )
{
}

int lns_collection_getAllocNum( void ) {
    return s_collection_allocNum;
}


lns_stem_t _lns_createImmediateVal(
    LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pVal )
{
    switch ( pVal->type ) {
    case lns_imdType_stem:
        return pVal->val.stem;
    case lns_imdType_int:
        return LNS_STEM_INT( pVal->val.valInt );
    case lns_imdType_real:
        return LNS_STEM_REAL( pVal->val.valReal );
    case lns_imdType_bool:
        return LNS_STEM_BOOL( pVal->val.valBool );
    case lns_imdType_str:
        return LNS_STEM_ANY( _lns_litStr2any( LNS_DEBUG_POS, _pEnv, pVal->val.str ) );
    case lns_imdType_list:
        return LNS_STEM_ANY( _lns_createList( LNS_DEBUG_POS, _pEnv, pVal->val.list ) );
    case lns_imdType_map:
        return LNS_STEM_ANY( _lns_createMap( LNS_DEBUG_POS, _pEnv, pVal->val.map ) );
    case lns_imdType_set:
        return LNS_STEM_ANY( _lns_createSet( LNS_DEBUG_POS, _pEnv, pVal->val.set ) );
    case lns_imdType_any:
        return LNS_STEM_ANY( pVal->val.any );
    default:
        abort();
    }
}


#define SETQ( STEM )                                              \
    if ( (STEM)->type == lns_stem_type_any ) {                   \
        lns_setQ_( (STEM)->val.pAny );                          \
    }




/// ================ List ======================

#define lns_obj_List( OBJ )                    \
    ((lns_List_t*)OBJ->val.classVal)

#define lns_obj_List_obj( OBJ )                                        \
    ((vector<lns_var_t>*)((lns_List_t*)OBJ->val.classVal)->pObj)


#define lns_ListClass vector<lns_var_t>
#define lns_ListIterator vector<lns_var_t>::iterator

struct lns_itList_t {
    lns_ListIterator it;
    lns_ListIterator end;
};



static void lns_mtd_List__del( lns_env_t * _pEnv, lns_any_t * pObj );

lns_mtd_List_t lns_mtd_List = {
    lns_mtd_List__del,
    NULL,
    (lns_method_t*)lns_mtd_List_insert,
    (lns_method_t*)lns_mtd_List_refAt,
    (lns_method_t*)lns_mtd_List_unpack,
    (lns_method_t*)lns_mtd_List_sort,
};


lns_any_t * lns_itList_new( lns_env_t * _pEnv, lns_any_t * pList )
{
    lns_lock( s_collection_allocNum++; );
    lns_any_t * pAny = lns_it_new(
        _pEnv, lns_value_type_itList, new lns_itList_t );
#ifdef DEBUG_LOG
    printf( "new = %p, %s\n", pAny->val.itList, __func__ );
#endif
    
    pAny->val.itList->it = lns_obj_List_obj( pList )->begin();
    pAny->val.itList->end = lns_obj_List_obj( pList )->end();
    return pAny;
}

void lns_itList__del( lns_env_t * _pEnv, lns_any_t * it )
{
    delete it->val.itList;
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", it->val.itList, __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

void lns_itList_inc( lns_env_t * _pEnv, lns_any_t * it )
{
    it->val.itList->it++;
}

bool lns_itList_hasNext( lns_env_t * _pEnv, lns_any_t * it, lns_stem_t * pVal )
{
    if ( it->val.itList->it != it->val.itList->end ) {
        *pVal = (*it->val.itList->it).stem;
        return true;
    }
    return false;
}


void lns_class_List_init( lns_env_t * _pEnv, lns_List_t * pObj )
{
    lns_lock( s_collection_allocNum++; );
    pObj->pObj = new vector<lns_var_t>();
#ifdef DEBUG_LOG
    printf( "new = %p, %s\n", pObj->pObj, __func__ );
#endif
}


lns_any_t * lns_class_List_new( lns_env_t * _pEnv )
{
    lns_class_new2_( _pEnv, lns_List_t, List, pAny, pObj );
    
    lns_class_List_init( _pEnv, pObj );

    return pAny;
}

lns_any_t * lns_List_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny )
{
    lns_any_t * pAny = lns_class_List_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index++ ) {
        lns_var_t var;
        lns_initVar( &var, lns_fromDDD( pDDDAny, index ) );
        
        lns_obj_List_obj( pAny )->push_back( var );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lns_mtd_List__del( lns_env_t * _pEnv, lns_any_t * pObj )
{
    LNS_DEBUG_CALL_LOG;

    lns_ListIterator it;
    lns_ListIterator end = lns_obj_List_obj( pObj )->end();
    
    delete lns_obj_List_obj( pObj );
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", lns_obj_List_obj( pObj ), __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

/**
   List.insert()  処理
*/
lns_stem_t lns_mtd_List_insert(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val )
{
    // pVal の参照カウントをインクリメントするため、 work に setq する

    lns_var_t var;
    lns_initVar( &var, val );
    
    lns_obj_List_obj( pObj )->push_back( var );
    
    return lns_global.nilStem;
}

/**
   List.remove()  処理
*/
lns_stem_t lns_mtd_List_remove(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val )
{
    int len = lns_mtd_List_len( _pEnv, pObj );

    lns_ListIterator it;
    
    if ( val.type == lns_stem_type_int ) {
        // インデックスが指定された場合は、その場所を削除
        int index = val.val.intVal;
        if ( index < len ) {
            it = lns_obj_List_obj( pObj )->begin() + index;
        }
        else {
            return lns_global.nilStem;
        }
    }
    else {
        // インデックス指定されていない場合は、末尾を削除
        if ( len > 0 ) {
            it = lns_obj_List_obj( pObj )->begin() + len - 1;
        }
        else {
            return lns_global.nilStem;
        }
    }

    
    lns_var_t var = *it;
    lns_obj_List_obj( pObj )->erase( it );
    
    lns_setRet( _pEnv, var.stem );
    lns_unuseVar( &var );
    return var.stem;
}


/**
   リストの要素アクセス。
*/
lns_stem_t lns_mtd_List_refAt(
    lns_env_t * _pEnv, lns_any_t * pObj, int index )
{
    if ( (int)lns_obj_List_obj( pObj )->size() >= index ) {
        return lns_obj_List_obj( pObj )->at( index - 1 ).stem;
    }
    return lns_global.nilStem;
}

void lns_mtd_List_setAt(
    lns_env_t * _pEnv, lns_any_t * pListAny, int index, lns_stem_t val )
{
    if ( index <= 0 ) {
        lns_abort( "illegal index" );
    }

    lns_var_t var;
    lns_initVar( &var, val );

    
    if ( (int)lns_obj_List_obj( pListAny )->size() >= index ) {
        // リストの要素の間に一旦挿入する
        lns_ListIterator it = lns_obj_List_obj( pListAny )->begin() + index - 1;
        lns_obj_List_obj( pListAny )->insert( it, var );
        it++;

        // 挿入した後ろの要素を削除する
        lns_var_t work = *it;
        lns_obj_List_obj( pListAny )->erase( it );
        lns_unuseVar( &work );
    }
    else {
        int count = index - lns_obj_List_obj( pListAny )->size() - 1;
        for ( ; count > 0; count-- ) {
            lns_obj_List_obj( pListAny )->push_back( lns_global.nilVar );
        }
        lns_obj_List_obj( pListAny )->push_back( var );
    }
}



/**
   List.unpack() 処理
*/
lns_stem_t lns_mtd_List_unpack( lns_env_t * _pEnv, lns_any_t * pObj )
{
    lns_stem_t ddd = lns_createDDDOnly( _pEnv, lns_obj_List_obj( pObj )->size() );
    lns_any_t * pDDD = ddd.val.pAny;
        

    lns_ListIterator it;
    lns_ListIterator end = lns_obj_List_obj( pObj )->end();

    int index = 0;
    for ( it = lns_obj_List_obj( pObj )->begin(); it != end; it++, index++ )
    {
        lns_set2DDDArg( pDDD, index, (*it).stem );
    }

    return LNS_STEM_ANY( pDDD );
}

lns_int_t lns_mtd_List_len( lns_env_t * _pEnv, lns_any_t * pListAny )
{
    return (lns_int_t)lns_obj_List_obj( pListAny )->size();
}


/**
   sort() 用のデフォルト比較関数 (昇順でソート)

   @param pVal1 比較対象1
   @param pVal2 比較対象2
   @return int 比較対象1 < 比較対象2 ならば 0 以外
*/
static bool lns_mtd__Cmp( lns_var_t _val1, lns_var_t _val2 )
{
    lns_stem_t * pStem1 = &_val1.stem;
    lns_stem_t * pStem2 = &_val2.stem;
    
    if ( pStem1->type != pStem2->type ) {
        lns_real_t work1;
        switch ( pStem1->type ) {
        case lns_stem_type_int:
            work1 = pStem1->val.intVal;
            break;
        case lns_stem_type_real:
            work1 = pStem1->val.realVal;
            break;
        default:
            return pStem1->type < pStem2->type;
        }
        switch ( pStem2->type ) {
        case lns_stem_type_int:
            return work1 < pStem2->val.intVal;
        case lns_stem_type_real:
            return work1 < pStem2->val.realVal;
        default:
            return pStem1->type < pStem2->type;
        }
    }

    // val1, pStem2 が同じタイプの場合

    switch ( pStem1->type ) {
    case lns_stem_type_int:
        return pStem1->val.intVal < pStem2->val.intVal;
    case lns_stem_type_real:
        return pStem1->val.realVal < pStem2->val.realVal;
    case lns_stem_type_bool:
        return pStem1->val.boolVal < pStem2->val.boolVal;
    default:
        break;
    }

    lns_any_t * pAny1 = pStem1->val.pAny;
    lns_any_t * pAny2 = pStem2->val.pAny;

    if ( pAny1->type != pAny2->type ) {
        return pAny1->type < pAny2->type;
    }

    switch ( pAny1->type ) {
    case lns_value_type_str:
        return strcmp( pAny1->val.str.pStr, pAny2->val.str.pStr ) < 0;
    default:
        break;
    }

    return pAny1 < pAny2;
}

/**
   lns_mtd__CmpLune 用のコールバック環境。
*/
static lns_env_t * s_pSortEnv;

/**
   sort() のユーザフォーム実行用関数

   @param pVal1 比較対象1
   @param pVal2 比較対象2
   @return 比較対象1 < 比較対象2 ならば true
*/
static bool lns_mtd__CmpLune( lns_var_t val1, lns_var_t val2 )
{
    lns_stem_t pCond =
        s_pSortEnv->pSortCallback.val.pAny->val.form.pFunc(
            s_pSortEnv, s_pSortEnv->pSortCallback.val.pAny, val1.stem, val2.stem );

    return lns_isCondTrue( pCond );
}

/**
   List.sort() 
*/
lns_stem_t lns_mtd_List_sort(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t pForm )
{
    if ( pForm.type == lns_stem_type_nil || pForm.type == lns_stem_type_none ) {
        std::sort( lns_obj_List_obj( pObj )->begin(),
                   lns_obj_List_obj( pObj )->end(), lns_mtd__Cmp );
    }
    else {
        lns_lock(
            s_pSortEnv = _pEnv;
            _pEnv->pSortCallback = pForm;
            std::sort( lns_obj_List_obj( pObj )->begin(),
                       lns_obj_List_obj( pObj )->end(), lns_mtd__CmpLune );
            s_pSortEnv = NULL;
        );
    }

    return lns_global.nilStem;
}


/**
 * List を生成する
 *
 * @param pList 生成するリストの要素
 * @return 生成したリスト
 */
lns_any_t * _lns_createList(
    LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pList )
{
    lns_any_t * pAny = lns_class_List_new( _pEnv );

    for ( ; pList->type != lns_imdType_sentinel; pList++ ) {
        if ( (pList+1)->type == lns_imdType_sentinel &&
             pList->type == lns_imdType_any &&
             pList->val.any->type == lns_value_type_ddd )
        {
            // 最終が ... の処理
            lns_any_t * pDDD = pList->val.any;
            int index;
            for ( index = 0; index < lns_lenDDD( pDDD ); index++ ) {
                lns_mtd_List_insert( _pEnv, pAny, lns_getValFromDDD( pDDD, index ) );
            }
            break;
        }
        lns_stem_t val = _lns_createImmediateVal( LNS_DEBUG_POS, _pEnv, pList );
        lns_mtd_List_insert( _pEnv, pAny, val );
    }
    return pAny;
}

/// ================ Array ======================



/// ================ Set ======================
#define lns_SetClass set<lns_var_t,lns_SetComp>
#define lns_SetIterator lns_SetClass::iterator


class lns_SetComp {
public:
    bool operator()(const lns_var_t _val1, const lns_var_t _val2) {
        const lns_stem_t * pVal1 = &_val1.stem;
        const lns_stem_t * pVal2 = &_val2.stem;
        
        if ( pVal1->type != pVal2->type ) {
            return pVal1->val.pAny < pVal2->val.pAny;
        }
        switch ( pVal1->type ) {
        case lns_stem_type_int:
            return pVal1->val.intVal < pVal2->val.intVal;
        case lns_stem_type_real:
            return pVal1->val.realVal < pVal2->val.realVal;
        case lns_stem_type_bool:
            return pVal1->val.boolVal < pVal2->val.boolVal;
        default:
            break;
        }
        const lns_any_t * pAny1 = pVal1->val.pAny;
        const lns_any_t * pAny2 = pVal2->val.pAny;
        if ( pAny1->type != pAny2->type ) {
            return pAny1 < pAny2;
        }
        switch ( pAny1->type ) {
        case lns_value_type_str:
            return memcmp( pAny1->val.str.pStr, pAny2->val.str.pStr,
                           pAny1->val.str.len + 1 ) < 0;
        default:
            return pAny1 < pAny2;
        }
    }
};

#define lns_obj_Set( OBJ )                     \
    ((lns_Set_t*)OBJ->val.classVal)

#define lns_obj_Set_obj( OBJ )                                 \
    ((lns_SetClass*)((lns_Set_t*)OBJ->val.classVal)->pObj)

#define lns_mtd_Set_has_( OBJ, VAL )                                   \
    (lns_obj_Set_obj( OBJ )->find( VAL ) != lns_obj_Set_obj( OBJ )->end())

struct lns_itSet_t {
    lns_SetIterator it;
    lns_SetIterator end;
};


static void lns_mtd_Set__del( lns_env_t * _pEnv, lns_any_t * pObj );


lns_mtd_Set_t lns_mtd_Set = {
    lns_mtd_Set__del,
    NULL,
    (lns_method_t*)lns_mtd_Set_add,
    (lns_method_t*)lns_mtd_Set_del,
    (lns_method_t*)lns_mtd_Set_has,
    (lns_method_t*)lns_mtd_Set_and,
    (lns_method_t*)lns_mtd_Set_or,
    (lns_method_t*)lns_mtd_Set_sub,
    (lns_method_t*)lns_mtd_Set_clone,
    (lns_method_t*)lns_mtd_Set_len,
};




void lns_class_Set_init( lns_env_t * _pEnv, lns_Set_t * pObj )
{
    lns_lock( s_collection_allocNum++; );
    pObj->pObj = new lns_SetClass();
#ifdef DEBUG_LOG
    printf( "new = %p, %s\n", pObj->pObj, __func__ );
#endif
}


lns_any_t * lns_class_Set_new( lns_env_t * _pEnv )
{
    lns_class_new2_( _pEnv, lns_Set_t, Set, pAny, pObj );
    
    lns_class_Set_init( _pEnv, pObj );

    return pAny;
}

lns_any_t * lns_Set_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny )
{
    lns_any_t * pAny = lns_class_Set_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index++ ) {
        lns_stem_t stem = lns_fromDDD( pDDDAny, index );
        SETQ( &stem );

        lns_var_t var;
        lns_initVar( &var, stem );
        lns_obj_Set_obj( pAny )->insert( var );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lns_mtd_Set__del( lns_env_t * _pEnv, lns_any_t * pObj )
{
    LNS_DEBUG_CALL_LOG;

    lns_SetIterator it;
    lns_SetIterator end = lns_obj_Set_obj( pObj )->end();
    
    for ( it = lns_obj_Set_obj( pObj )->begin(); it != end; it++ )
    {
        lns_var_t var = *it;
        lns_unuseVar( &var );
    }
    
    delete lns_obj_Set_obj( pObj );
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", lns_obj_Set_obj( pObj ), __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

lns_any_t * lns_itSet_new( lns_env_t * _pEnv, lns_any_t * pSet )
{
    lns_lock( s_collection_allocNum++; );
    lns_any_t * pAny = lns_it_new(
        _pEnv, lns_value_type_itSet, new lns_itSet_t );
#ifdef DEBUG_LOG
    printf( "new = %p, %s\n", pAny->val.itSet, __func__ );
#endif
    pAny->val.itSet->it = lns_obj_Set_obj( pSet )->begin();
    pAny->val.itSet->end = lns_obj_Set_obj( pSet )->end();
    return pAny;
}

void lns_itSet__del( lns_env_t * _pEnv, lns_any_t * it )
{
    delete it->val.itSet;
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", it->val.itSet, __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

void lns_itSet_inc( lns_env_t * _pEnv, lns_any_t * it )
{
    it->val.itSet->it++;
}

bool lns_itSet_hasNext( lns_env_t * _pEnv, lns_any_t * it, lns_stem_t * pVal )
{
    if ( it->val.itSet->it != it->val.itSet->end ) {
        *pVal = (*it->val.itSet->it).stem;
        return true;
    }
    return false;
}

void lns_mtd_Set_add(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val )
{
    if ( val.type != lns_stem_type_nil ) {
        if ( lns_obj_Set_obj( pObj )->find( WORK_VAR( val ) ) ==
             lns_obj_Set_obj( pObj )->end() )
        {
            // 次の val は追加しない
            // - nil
            // - 登録済み

            lns_var_t var;
            lns_initVar( &var, val );
            
            lns_obj_Set_obj( pObj )->insert( var );
        }
    }
}

void lns_mtd_Set_del(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val )
{
    lns_SetIterator it = lns_obj_Set_obj( pObj )->find( WORK_VAR( val ) );

    if ( it != lns_obj_Set_obj( pObj )->end() ) {
        lns_var_t var = *it;
        lns_unuseVar( &var );
        lns_obj_Set_obj( pObj )->erase( it );
    }
}

bool lns_mtd_Set_has(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t val )
{
    bool contains = lns_obj_Set_obj( pObj )->find( WORK_VAR( val ) ) !=
        lns_obj_Set_obj( pObj )->end();
    return contains;
}

lns_any_t * lns_mtd_Set_and(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet )
{
    lns_ListClass list;
    {
        // A - A ∩ B の要素を list に格納
        lns_SetIterator it = lns_obj_Set_obj( pObj )->begin();
        lns_SetIterator end = lns_obj_Set_obj( pObj )->end();
    
        for ( ; it != end; it++ ) {
            if ( !lns_mtd_Set_has_( pSet, *it ) ) {
                list.push_back( *it );
            }
        }
    }

    {
        // 集合から list の要素を除外
        lns_ListIterator it = list.begin();
        lns_ListIterator end = list.end();
        
        for ( ; it != end; it++ ) {
            lns_SetIterator setIt = lns_obj_Set_obj( pObj )->find( *it );
            if ( setIt != lns_obj_Set_obj( pObj )->end() ) {
                lns_var_t var = *setIt;
                lns_unuseVar( &var );
                lns_obj_Set_obj( pObj )->erase( setIt );
            }
        }
    }
    
    return pObj;
}

lns_any_t * lns_mtd_Set_or(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet )
{
    lns_SetIterator it = lns_obj_Set_obj( pSet )->begin();
    lns_SetIterator end = lns_obj_Set_obj( pSet )->end();
    
    for ( ; it != end; it++ ) {
        lns_mtd_Set_add( _pEnv, pObj, (*it).stem );
    }
    
    return pObj;
}

lns_any_t * lns_mtd_Set_sub(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * pSet )
{
    lns_ListClass list;

    {
        // and 集合の要素を抽出
        lns_SetIterator it = lns_obj_Set_obj( pObj )->begin();
        lns_SetIterator end = lns_obj_Set_obj( pObj )->end();
    
        for ( ; it != end; it++ ) {
            if ( lns_mtd_Set_has_( pSet, *it ) ) {
                lns_var_t var;
                lns_initVar( &var, (*it).stem );
                list.push_back( var );
            }
        }
    }

    {
        // 抽出した要素を除外
        lns_ListIterator it = list.begin();
        lns_ListIterator end = list.end();
        
        for ( ; it != end; it++ ) {
            lns_var_t var = *it;
            lns_unuseVar( &var );
            
            lns_obj_Set_obj( pObj )->erase( var );
        }
    }
    
    
    return pObj;
}

lns_any_t * lns_mtd_Set_clone( lns_env_t * _pEnv, lns_any_t * pObj )
{
    lns_any_t * pAny = lns_class_Set_new( _pEnv );

    lns_SetIterator it = lns_obj_Set_obj( pObj )->begin();
    lns_SetIterator end = lns_obj_Set_obj( pObj )->end();
    
    for ( ; it != end; it++ ) {
        lns_var_t var;
        lns_initVar( &var, (*it).stem );

        lns_obj_Set_obj( pAny )->insert( var );
    }
    
    return pAny;
}

lns_int_t lns_mtd_Set_len( lns_env_t * _pEnv, lns_any_t * pObj )
{
    return (lns_int_t)lns_obj_Set_obj( pObj )->size();
}

lns_any_t * lns_mtd_Set_createList( lns_env_t * _pEnv, lns_any_t * pObj )
{
    lns_any_t * pList = lns_class_List_new( _pEnv );

    lns_SetClass * pSet = lns_obj_Set_obj( pObj );
    lns_SetIterator it;
    lns_SetIterator end = pSet->end();
    
    for ( it = pSet->begin(); it != end; it++ ) {
        lns_mtd_List_insert( _pEnv, pList, (*it).stem );
    }

    return pList;
}


/**
 * Set を生成する
 *
 * @param pSet 生成する Set の要素
 * @return 生成した Set
 */
lns_any_t * _lns_createSet(
    LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pSet )
{
    lns_any_t * pAny = lns_class_Set_new( _pEnv );

    for ( ; pSet->type != lns_imdType_sentinel; pSet++ ) {
        lns_mtd_Set_add(
            _pEnv, pAny, _lns_createImmediateVal( LNS_DEBUG_POS, _pEnv, pSet ) );
    }
    return pAny;
}

/// ================ Map ======================
#define lns_MapClass map<lns_var_t,lns_var_t,lns_SetComp>
#define lns_MapIterator lns_MapClass::iterator

#define lns_obj_Map( OBJ )                     \
    ((lns_Map_t*)OBJ->val.classVal)

#define lns_obj_Map_obj( OBJ )                                 \
    ((lns_MapClass*)((lns_Map_t*)OBJ->val.classVal)->pObj)

#define lns_mtd_Map_has_( OBJ, VAL )                                   \
    (lns_obj_Map_obj( OBJ )->find( VAL ) != lns_obj_Map_obj( OBJ )->end())

struct lns_itMap_t {
    lns_MapIterator it;
    lns_MapIterator end;
};


static void lns_mtd_Map__del( lns_env_t * _pEnv, lns_any_t * pObj );



lns_mtd_Map_t lns_mtd_Map = {
    lns_mtd_Map__del,
    NULL,
    (lns_method_t*)lns_mtd_Map_add,
    (lns_method_t*)lns_mtd_Map_get,
};




void lns_class_Map_init( lns_env_t * _pEnv, lns_Map_t * pObj )
{
    lns_lock( s_collection_allocNum++; );
    pObj->pObj = new lns_MapClass();
#ifdef DEBUG_LOG
    printf( "new = %p, %s\n", pObj->pObj, __func__ );
#endif
}


lns_any_t * lns_class_Map_new( lns_env_t * _pEnv )
{
    lns_class_new2_( _pEnv, lns_Map_t, Map, pAny, pObj );
    
    lns_class_Map_init( _pEnv, pObj );
    
    return pAny;
}

lns_any_t * lns_Map_ctor( lns_env_t * _pEnv, lns_any_t * pDDDAny )
{
    lns_any_t * pAny = lns_class_Map_new( _pEnv );
    lns_MapClass * pMap = lns_obj_Map_obj( pAny );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index += 2 ) {
        lns_var_t keyVar;
        lns_initVar( &keyVar, lns_fromDDD( pDDDAny, index ) );
        lns_var_t valVar;
        lns_initVar( &valVar, lns_fromDDD( pDDDAny, index + 1 ) );

        pMap->insert( lns_MapClass::value_type( keyVar, valVar ) );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lns_mtd_Map__del( lns_env_t * _pEnv, lns_any_t * pObj )
{
    LNS_DEBUG_CALL_LOG;

    lns_MapClass * pMap = lns_obj_Map_obj( pObj );
    
    lns_MapIterator it;
    lns_MapIterator end = pMap->end();
    
    for ( it = pMap->begin(); it != end; it++ ) {
        lns_var_t work = it->first;
        lns_unuseVar( &work );
        work = it->second;
        lns_unuseVar( &work );
    }
    
    delete pMap;
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", pMap, __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

lns_any_t * lns_itMap_new( lns_env_t * _pEnv, lns_any_t * pMap )
{
    lns_lock( s_collection_allocNum++; );
    lns_any_t * pAny = lns_it_new(
        _pEnv, lns_value_type_itMap, new lns_itMap_t );
#ifdef DEBUG_LOG
    printf( "alloc = %p, %s\n", pAny->val.itMap, __func__ );
#endif
    pAny->val.itMap->it = lns_obj_Map_obj( pMap )->begin();
    pAny->val.itMap->end = lns_obj_Map_obj( pMap )->end();

    return pAny;
}

void lns_itMap__del( lns_env_t * _pEnv, lns_any_t * it )
{
    delete it->val.itMap;
#ifdef DEBUG_LOG
    printf( "delete = %p, %s\n", it->val.itMap, __func__ );
#endif
    lns_lock( s_collection_allocNum--; );    
}

void lns_itMap_inc( lns_env_t * _pEnv, lns_any_t * it )
{
    it->val.itMap->it++;
}

bool lns_itMap_hasNext( lns_env_t * _pEnv, lns_any_t * it, lns_Map_entry_t * pEntry )
{
    if ( it->val.itMap->it != it->val.itMap->end ) {
        pEntry->key = it->val.itMap->it->first.stem;
        pEntry->val = it->val.itMap->it->second.stem;
        return true;
    }
    return false;
}

void lns_mtd_Map_add( lns_env_t * _pEnv, lns_any_t * pObj,
                       lns_stem_t key, lns_stem_t val )
{
    lns_MapClass * pMap = lns_obj_Map_obj( pObj );
    if ( val.type == lns_stem_type_nil )
    {
        // nil の場合は、削除

        lns_MapIterator it = pMap->find( WORK_VAR( key ) );
        if ( it != pMap->end() ) {
            lns_var_t work = it->first;
            lns_unuseVar( &work );
            work = it->second;
            lns_unuseVar( &work );
            
            pMap->erase( it );
        }
        
    }
    else {
        lns_var_t keyVar = WORK_VAR( key );
        if ( pMap->find( keyVar ) == pMap->end() ) {
            // 次の val は追加しない
            // - 登録済み
        
            // pVal の参照カウントをインクリメントするため、 work に setq する
            lns_var_t var;
            lns_initVar( &var, val );
            (*pMap)[ keyVar ] = var;
        }
    }
}

lns_stem_t lns_mtd_Map_get(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t key )
{
    lns_MapClass * pMap = lns_obj_Map_obj( pObj );
    lns_MapIterator it = pMap->find( WORK_VAR( key ) );
    
    if ( it == pMap->end() ) {
        return lns_global.nilStem;
    }

    return it->second.stem;
}

lns_any_t * lns_mtd_Map_createKeyList(
    lns_env_t * _pEnv, lns_any_t * pObj )
{
    lns_any_t * pList = lns_class_List_new( _pEnv );

    lns_MapClass * pMap = lns_obj_Map_obj( pObj );
    lns_MapIterator it;
    lns_MapIterator end = pMap->end();
    
    for ( it = pMap->begin(); it != end; it++ ) {
        lns_mtd_List_insert( _pEnv, pList, it->first.stem );
    }

    return pList;
}



/**
 * Map を生成する
 *
 * @param pMap 生成する Map の要素
 * @return 生成した Map
 */
lns_any_t * _lns_createMap(
    LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdEntry_t * pEntry )
{
    lns_any_t * pAny = lns_class_Map_new( _pEnv );

    for ( ; pEntry->key.type != lns_imdType_sentinel; pEntry++ ) {
        lns_mtd_Map_add(
            _pEnv, pAny,
            _lns_createImmediateVal( LNS_DEBUG_POS, _pEnv, &pEntry->key ),
            _lns_createImmediateVal( LNS_DEBUG_POS, _pEnv, &pEntry->val ) );
    }
    return pAny;
}

