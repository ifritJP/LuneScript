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
#include <lns_collection.h>
#include <vector>
#include <algorithm>
#include <set>
#include <map>

using namespace std;

void lune_collection_init( void )
{
}


lune_stem_t _lune_createImmediateVal(
    LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pVal )
{
    switch ( pVal->type ) {
    case lune_imdType_int:
        return LUNE_STEM_INT( pVal->val.valInt );
    case lune_imdType_real:
        return LUNE_STEM_REAL( pVal->val.valReal );
    case lune_imdType_str:
        return _lune_litStr2stem( LUNE_DEBUG_POS, _pEnv, pVal->val.str );
    case lune_imdType_list:
        return _lune_createList( LUNE_DEBUG_POS, _pEnv, pVal->val.list );
    case lune_imdType_map:
        return _lune_createMap( LUNE_DEBUG_POS, _pEnv, pVal->val.map );
    case lune_imdType_set:
        return _lune_createSet( LUNE_DEBUG_POS, _pEnv, pVal->val.set );
    case lune_imdType_any:
        return LUNE_STEM_ANY( pVal->val.any );
    default:
        abort();
    }
}


#define DECRE_STEM( ENV, STEM )                          \
    {                                                    \
        lune_stem_t _work = STEM;                        \
        if ( _work.type == lune_stem_type_any ) {        \
            lune_decre_ref( ENV, _work.val.pAny );       \
        }                                                \
    }

#define SETQ( STEM )                                              \
    if ( (STEM)->type == lune_stem_type_any ) {                   \
        lune_setQ_( (STEM)->val.pAny );                          \
    }




/// ================ List ======================

#define lune_obj_List( OBJ )                    \
    ((lune_List_t*)OBJ->val.classVal)

#define lune_obj_List_obj( OBJ )                                        \
    ((vector<lune_stem_t>*)((lune_List_t*)OBJ->val.classVal)->pObj)


#define lune_ListClass vector<lune_stem_t>
#define lune_ListIterator vector<lune_stem_t>::iterator

struct lune_itList_t {
    lune_ListIterator it;
    lune_ListIterator end;
};



static void lune_mtd_List__del( lune_env_t * _pEnv, lune_any_t * pObj );

lune_mtd_List_t lune_mtd_List = {
    lune_mtd_List__del,
    NULL,
    (lune_method_t*)lune_mtd_List_insert,
    (lune_method_t*)lune_mtd_List_refAt,
    (lune_method_t*)lune_mtd_List_unpack,
    (lune_method_t*)lune_mtd_List_sort,
};

lune_type_meta_t lune_type_meta_List = { "List" };


lune_any_t * lune_itList_new( lune_env_t * _pEnv, lune_any_t * pList )
{
    lune_any_t * pAny = lune_it_new(
        _pEnv, lune_value_type_itList, new lune_itList_t );
    pAny->val.itList->it = lune_obj_List_obj( pList )->begin();
    pAny->val.itList->end = lune_obj_List_obj( pList )->end();
    return pAny;
}

void lune_itList__del( lune_env_t * _pEnv, lune_any_t * it )
{
    delete it->val.itList;
}

void lune_itList_inc( lune_env_t * _pEnv, lune_any_t * it )
{
    it->val.itList->it++;
}

bool lune_itList_hasNext( lune_env_t * _pEnv, lune_any_t * it, lune_stem_t * pVal )
{
    if ( it->val.itList->it != it->val.itList->end ) {
        *pVal = *it->val.itList->it;
        return true;
    }
    return false;
}


void lune_class_List_init( lune_env_t * _pEnv, lune_List_t * pObj )
{
    pObj->pObj = new vector<lune_stem_t>();
}


lune_stem_t lune_class_List_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_List_t, List, pAny, pObj );
    
    lune_class_List_init( _pEnv, pObj );

    return pAny;
}

lune_stem_t lune_List_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny )
{
    lune_stem_t pAny = lune_class_List_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index++ ) {
        lune_stem_t stem = lune_fromDDD( pDDDAny, index );
        SETQ( &stem );
        
        lune_obj_List_obj( pAny.val.pAny )->push_back( stem );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lune_mtd_List__del( lune_env_t * _pEnv, lune_any_t * pObj )
{
    LUNE_DEBUG_CALL_LOG;

    lune_ListIterator it;
    lune_ListIterator end = lune_obj_List_obj( pObj )->end();
    
    for ( it = lune_obj_List_obj( pObj )->begin(); it != end; it++ )
    {
        lune_stem_t stem = *it;
        DECRE_STEM( _pEnv, stem );
    }
    
    delete lune_obj_List_obj( pObj );
}

/**
   List.insert()  処理
*/
lune_stem_t lune_mtd_List_insert(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val )
{
    // pVal の参照カウントをインクリメントするため、 work に setq する
    SETQ( &val );
    
    lune_obj_List_obj( pObj )->push_back( val );
    
    return lune_global.nilStem;
}

/**
   リストの要素アクセス。
*/
lune_stem_t lune_mtd_List_refAt(
    lune_env_t * _pEnv, lune_any_t * pObj, int index )
{
    if ( (int)lune_obj_List_obj( pObj )->size() >= index ) {
        return lune_obj_List_obj( pObj )->at( index - 1 );
    }
    return lune_global.nilStem;
}


/**
   List.unpack() 処理
*/
lune_stem_t lune_mtd_List_unpack( lune_env_t * _pEnv, lune_any_t * pObj )
{
    lune_stem_t pDDD =
        lune_createDDDOnly( _pEnv, lune_obj_List_obj( pObj )->size() );

    lune_ListIterator it;
    lune_ListIterator end = lune_obj_List_obj( pObj )->end();

    int index = 0;
    for ( it = lune_obj_List_obj( pObj )->begin(); it != end; it++, index++ )
    {
        lune_set2DDDArg( pDDD.val.pAny, index, *it );
    }

    return pDDD;
}

/**
   sort() 用のデフォルト比較関数 (昇順でソート)

   @param pVal1 比較対象1
   @param pVal2 比較対象2
   @return int 比較対象1 < 比較対象2 ならば 0 以外
*/
static bool lune_mtd__Cmp( lune_stem_t val1, lune_stem_t val2 )
{
    if ( val1.type != val2.type ) {
        lune_real_t work1;
        switch ( val1.type ) {
        case lune_stem_type_int:
            work1 = val1.val.intVal;
            break;
        case lune_stem_type_real:
            work1 = val1.val.realVal;
            break;
        default:
            return val1.type < val2.type;
        }
        switch ( val2.type ) {
        case lune_stem_type_int:
            return work1 < val2.val.intVal;
        case lune_stem_type_real:
            return work1 < val2.val.realVal;
        default:
            return val1.type < val2.type;
        }
    }

    switch ( val1.type ) {
    case lune_stem_type_int:
        return val1.val.intVal < val2.val.intVal;
    case lune_stem_type_real:
        return val1.val.realVal < val2.val.realVal;
    default:
        break;
    }

    lune_any_t * pAny1 = val1.val.pAny;
    lune_any_t * pAny2 = val2.val.pAny;

    if ( pAny1->type != pAny2->type ) {
        return pAny1->type < pAny2->type;
    }

    switch ( pAny1->type ) {
    case lune_value_type_str:
        return strcmp( pAny1->val.str.pStr, pAny2->val.str.pStr ) < 0;
    default:
        break;
    }

    return pAny1 < pAny2;
}

/**
   lune_mtd__CmpLune 用のコールバック環境。
*/
static lune_env_t * s_pSortEnv;

/**
   sort() のユーザフォーム実行用関数

   @param pVal1 比較対象1
   @param pVal2 比較対象2
   @return 比較対象1 < 比較対象2 ならば true
*/
static bool lune_mtd__CmpLune( lune_stem_t val1, lune_stem_t val2 )
{
    lune_any_t * pRet;
    lune_stem_t pCond =
        s_pSortEnv->pSortCallback.val.pAny->val.form.pFunc(
            s_pSortEnv, s_pSortEnv->pSortCallback.val.pAny, val1, val2 );

    lune_setQ_any( (&pRet), pCond.val.pAny );

    bool ret = lune_isCondTrue( pCond );

    lune_decre_ref( s_pSortEnv, pRet );
    
    return ret;
}

/**
   List.sort() 
*/
lune_stem_t lune_mtd_List_sort(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pForm )
{
    if ( pForm.type == lune_stem_type_nil ) {
        std::sort( lune_obj_List_obj( pObj )->begin(),
                   lune_obj_List_obj( pObj )->end(), lune_mtd__Cmp );
    }
    else {
        lune_lock(
            s_pSortEnv = _pEnv;
            _pEnv->pSortCallback = pForm;
            std::sort( lune_obj_List_obj( pObj )->begin(),
                       lune_obj_List_obj( pObj )->end(), lune_mtd__CmpLune );
            s_pSortEnv = NULL;
        );
    }

    return lune_global.nilStem;
}


/**
 * List を生成する
 *
 * @param pList 生成するリストの要素
 * @return 生成したリスト
 */
lune_stem_t _lune_createList(
    LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pList )
{
    lune_stem_t pAny = lune_class_List_new( _pEnv );

    for ( ; pList->type != lune_imdType_sentinel; pList++ ) {
        lune_mtd_List_insert(
            _pEnv, pAny.val.pAny,
            _lune_createImmediateVal( LUNE_DEBUG_POS, _pEnv, pList ) );
    }
    return pAny;
}

/// ================ Array ======================



/// ================ Set ======================
#define lune_SetClass set<lune_stem_t,lune_SetComp>
#define lune_SetIterator lune_SetClass::iterator


class lune_SetComp {
public:
    bool operator()(const lune_stem_t val1, const lune_stem_t val2) {
        if ( val1.type != val2.type ) {
            return val1.val.pAny < val2.val.pAny;
        }
        switch ( val1.type ) {
        case lune_stem_type_int:
            return val1.val.intVal < val2.val.intVal;
        case lune_stem_type_real:
            return val1.val.realVal < val2.val.realVal;
        case lune_stem_type_bool:
            return val1.val.boolVal < val2.val.boolVal;
        default:
            break;
        }
        const lune_any_t * pAny1 = val1.val.pAny;
        const lune_any_t * pAny2 = val2.val.pAny;
        if ( pAny1->type != pAny2->type ) {
            return pAny1 < pAny2;
        }
        switch ( pAny1->type ) {
        case lune_value_type_str:
            return memcmp( pAny1->val.str.pStr, pAny2->val.str.pStr,
                           pAny1->val.str.len + 1 ) < 0;
        default:
            return pAny1 < pAny2;
        }
    }
};

#define lune_obj_Set( OBJ )                     \
    ((lune_Set_t*)OBJ->val.classVal)

#define lune_obj_Set_obj( OBJ )                                 \
    ((lune_SetClass*)((lune_Set_t*)OBJ->val.classVal)->pObj)

#define lune_mtd_Set_has_( OBJ, VAL )                                   \
    (lune_obj_Set_obj( OBJ )->find( VAL ) != lune_obj_Set_obj( OBJ )->end())

struct lune_itSet_t {
    lune_SetIterator it;
    lune_SetIterator end;
};


static void lune_mtd_Set__del( lune_env_t * _pEnv, lune_any_t * pObj );


void lune_mtd_Set_add(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val );
void lune_mtd_Set_del(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val );
lune_stem_t lune_mtd_Set_has(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val );
lune_stem_t lune_mtd_Set_and_(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet );
lune_stem_t lune_mtd_Set_or_(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet );
lune_stem_t lune_mtd_Set_sub(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet );
lune_stem_t lune_mtd_Set_clone( lune_env_t * _pEnv, lune_any_t * pObj );
lune_int_t lune_mtd_Set_len( lune_env_t * _pEnv, lune_any_t * pObj );


lune_mtd_Set_t lune_mtd_Set = {
    lune_mtd_Set__del,
    NULL,
    (lune_method_t*)lune_mtd_Set_add,
    (lune_method_t*)lune_mtd_Set_del,
    (lune_method_t*)lune_mtd_Set_has,
    (lune_method_t*)lune_mtd_Set_and_,
    (lune_method_t*)lune_mtd_Set_or_,
    (lune_method_t*)lune_mtd_Set_sub,
    (lune_method_t*)lune_mtd_Set_clone,
    (lune_method_t*)lune_mtd_Set_len,
};

lune_type_meta_t lune_type_meta_Set = { "Set" };



void lune_class_Set_init( lune_env_t * _pEnv, lune_Set_t * pObj )
{
    pObj->pObj = new lune_SetClass();
}


lune_stem_t lune_class_Set_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_Set_t, Set, pAny, pObj );
    
    lune_class_Set_init( _pEnv, pObj );

    return pAny;
}

lune_stem_t lune_Set_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny )
{
    lune_stem_t pAny = lune_class_Set_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index++ ) {
        lune_stem_t stem = lune_fromDDD( pDDDAny, index );
        SETQ( &stem );
    
        lune_obj_Set_obj( pAny.val.pAny )->insert( stem );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lune_mtd_Set__del( lune_env_t * _pEnv, lune_any_t * pObj )
{
    LUNE_DEBUG_CALL_LOG;

    lune_SetIterator it;
    lune_SetIterator end = lune_obj_Set_obj( pObj )->end();
    
    for ( it = lune_obj_Set_obj( pObj )->begin(); it != end; it++ )
    {
        DECRE_STEM( _pEnv, *it );
    }
    
    delete lune_obj_Set_obj( pObj );
}

lune_any_t * lune_itSet_new( lune_env_t * _pEnv, lune_any_t * pSet )
{
    lune_any_t * pAny = lune_it_new(
        _pEnv, lune_value_type_itSet, new lune_itSet_t );
    pAny->val.itSet->it = lune_obj_Set_obj( pSet )->begin();
    pAny->val.itSet->end = lune_obj_Set_obj( pSet )->end();
    return pAny;
}

void lune_itSet__del( lune_env_t * _pEnv, lune_any_t * it )
{
    delete it->val.itSet;
}

void lune_itSet_inc( lune_env_t * _pEnv, lune_any_t * it )
{
    it->val.itSet->it++;
}

bool lune_itSet_hasNext( lune_env_t * _pEnv, lune_any_t * it, lune_stem_t * pVal )
{
    if ( it->val.itSet->it != it->val.itSet->end ) {
        *pVal = *it->val.itSet->it;
        return true;
    }
    return false;
}

void lune_mtd_Set_add(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val )
{
    if ( val.type != lune_stem_type_nil ) {
        if ( lune_obj_Set_obj( pObj )->find( val ) == lune_obj_Set_obj( pObj )->end() )
        {
            // 次の val は追加しない
            // - nil
            // - 登録済み
        
            // pVal の参照カウントをインクリメントするため、 work に setq する
            SETQ( &val );
    
            lune_obj_Set_obj( pObj )->insert( val );
        }
    }
}

void lune_mtd_Set_del(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val )
{
    lune_SetIterator it = lune_obj_Set_obj( pObj )->find( val );

    if ( it != lune_obj_Set_obj( pObj )->end() ) {
        DECRE_STEM( _pEnv, *it );
        lune_obj_Set_obj( pObj )->erase( it );
    }
}

lune_stem_t lune_mtd_Set_has(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t val )
{
    bool contains =
        lune_obj_Set_obj( pObj )->find( val ) != lune_obj_Set_obj( pObj )->end();
    return LUNE_STEM_BOOL( contains );
}

lune_stem_t lune_mtd_Set_and_(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet )
{
    lune_ListClass list;
    {
        // A - A ∩ B の要素を list に格納
        lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
        lune_SetIterator end = lune_obj_Set_obj( pObj )->end();
    
        for ( ; it != end; it++ ) {
            if ( !lune_mtd_Set_has_( pSet.val.pAny, *it ) ) {
                list.push_back( *it );
            }
        }
    }

    {
        // 集合から list の要素を除外
        lune_ListIterator it = lune_obj_List_obj( pObj )->begin();
        lune_ListIterator end = lune_obj_List_obj( pObj )->end();
        
        for ( ; it != end; it++ ) {
            lune_SetIterator setIt = lune_obj_Set_obj( pObj )->find( *it );
            if ( setIt != lune_obj_Set_obj( pObj )->end() ) {
                DECRE_STEM( _pEnv, *setIt );
                lune_obj_Set_obj( pObj )->erase( setIt );
            }
        }
    }
    
    return LUNE_STEM_ANY( pObj );
}

lune_stem_t lune_mtd_Set_or_(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet )
{
    lune_SetIterator it = lune_obj_Set_obj( pSet.val.pAny )->begin();
    lune_SetIterator end = lune_obj_Set_obj( pSet.val.pAny )->begin();
    
    for ( ; it != end; it++ ) {
        lune_mtd_Set_add( _pEnv, pObj, *it );
    }
    
    return LUNE_STEM_ANY( pObj );
}

lune_stem_t lune_mtd_Set_sub(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t pSet )
{
    lune_ListClass list;

    {
        // and 集合の要素を抽出
        lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
        lune_SetIterator end = lune_obj_Set_obj( pObj )->begin();
    
        for ( ; it != end; it++ ) {
            if ( lune_mtd_Set_has_( pSet.val.pAny, *it ) ) {
                list.push_back( *it );
            }
        }
    }

    {
        // 抽出した要素を除外
        lune_ListIterator it = lune_obj_List_obj( pObj )->begin();
        lune_ListIterator end = lune_obj_List_obj( pObj )->end();
        
        for ( ; it != end; it++ ) {
            DECRE_STEM( _pEnv, *it );
            
            lune_obj_Set_obj( pObj )->erase( *it );
        }
    }
    
    
    return LUNE_STEM_ANY( pObj );
}

lune_stem_t lune_mtd_Set_clone( lune_env_t * _pEnv, lune_any_t * pObj )
{
    lune_stem_t pAny = lune_class_Set_new( _pEnv );

    lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
    lune_SetIterator end = lune_obj_Set_obj( pObj )->begin();
    
    for ( ; it != end; it++ ) {
        lune_stem_t stem = *it;
        SETQ( &stem );

        lune_obj_Set_obj( pAny.val.pAny )->insert( stem );
    }
    
    return pAny;
}

lune_int_t lune_mtd_Set_len( lune_env_t * _pEnv, lune_any_t * pObj )
{
    return (lune_int_t)lune_obj_Set_obj( pObj )->size();
}

lune_stem_t lune_mtd_Set_createList( lune_env_t * _pEnv, lune_any_t * pObj )
{
    lune_stem_t pList = lune_class_List_new( _pEnv );

    lune_SetClass * pSet = lune_obj_Set_obj( pObj );
    lune_SetIterator it;
    lune_SetIterator end = pSet->end();
    
    for ( it = pSet->begin(); it != end; it++ ) {
        lune_mtd_List_insert( _pEnv, pList.val.pAny, *it );
    }

    return pList;
}


/**
 * Set を生成する
 *
 * @param pSet 生成する Set の要素
 * @return 生成した Set
 */
lune_stem_t _lune_createSet(
    LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pSet )
{
    lune_stem_t pAny = lune_class_Set_new( _pEnv );

    for ( ; pSet->type != lune_imdType_sentinel; pSet++ ) {
        lune_mtd_Set_add(
            _pEnv, pAny.val.pAny,
            _lune_createImmediateVal( LUNE_DEBUG_POS, _pEnv, pSet ) );
    }
    return pAny;
}

/// ================ Map ======================
#define lune_MapClass map<lune_stem_t,lune_stem_t,lune_SetComp>
#define lune_MapIterator lune_MapClass::iterator

#define lune_obj_Map( OBJ )                     \
    ((lune_Map_t*)OBJ->val.classVal)

#define lune_obj_Map_obj( OBJ )                                 \
    ((lune_MapClass*)((lune_Map_t*)OBJ->val.classVal)->pObj)

#define lune_mtd_Map_has_( OBJ, VAL )                                   \
    (lune_obj_Map_obj( OBJ )->find( VAL ) != lune_obj_Map_obj( OBJ )->end())

struct lune_itMap_t {
    lune_MapIterator it;
    lune_MapIterator end;
};


static void lune_mtd_Map__del( lune_env_t * _pEnv, lune_any_t * pObj );



lune_mtd_Map_t lune_mtd_Map = {
    lune_mtd_Map__del,
    NULL,
    (lune_method_t*)lune_mtd_Map_add,
    (lune_method_t*)lune_mtd_Map_get,
};

lune_type_meta_t lune_type_meta_Map = { "Map" };



void lune_class_Map_init( lune_env_t * _pEnv, lune_Map_t * pObj )
{
    pObj->pObj = new lune_MapClass();
}


lune_stem_t lune_class_Map_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_Map_t, Map, pAny, pObj );
    
    lune_class_Map_init( _pEnv, pObj );
    
    return pAny;
}

lune_stem_t lune_Map_ctor( lune_env_t * _pEnv, lune_any_t * pDDDAny )
{
    lune_stem_t pAny = lune_class_Map_new( _pEnv );
    lune_MapClass * pMap = lune_obj_Map_obj( pAny.val.pAny );

    int index;
    for ( index = 0; index < pDDDAny->val.ddd.len; index += 2 ) {
        lune_stem_t keyStem = lune_fromDDD( pDDDAny, index );
        SETQ( &keyStem );
        
        lune_stem_t valStem = lune_fromDDD( pDDDAny, index + 1 );
        SETQ( &valStem );

        pMap->insert( lune_MapClass::value_type( keyStem, valStem ) );
    }
    
    return pAny;
}

/**
   List インスタンスの開放処理
*/
static void lune_mtd_Map__del( lune_env_t * _pEnv, lune_any_t * pObj )
{
    LUNE_DEBUG_CALL_LOG;

    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    
    lune_MapIterator it;
    lune_MapIterator end = pMap->end();
    
    for ( it = pMap->begin(); it != end; it++ ) {
        DECRE_STEM( _pEnv, it->first );
        DECRE_STEM( _pEnv, it->second );
    }
    
    delete pMap;
}

lune_any_t * lune_itMap_new( lune_env_t * _pEnv, lune_any_t * pMap )
{
    lune_any_t * pAny = lune_it_new(
        _pEnv, lune_value_type_itMap, new lune_itMap_t );
    pAny->val.itMap->it = lune_obj_Map_obj( pMap )->begin();
    pAny->val.itMap->end = lune_obj_Map_obj( pMap )->end();

    return pAny;
}

void lune_itMap__del( lune_env_t * _pEnv, lune_any_t * it )
{
    delete it->val.itMap;
}

void lune_itMap_inc( lune_env_t * _pEnv, lune_any_t * it )
{
    it->val.itMap->it++;
}

bool lune_itMap_hasNext( lune_env_t * _pEnv, lune_any_t * it, lune_Map_entry_t * pEntry )
{
    if ( it->val.itMap->it != it->val.itMap->end ) {
        pEntry->key = it->val.itMap->it->first;
        pEntry->val = it->val.itMap->it->second;
        return true;
    }
    return false;
}

void lune_mtd_Map_add( lune_env_t * _pEnv, lune_any_t * pObj,
                               lune_stem_t key, lune_stem_t val )
{
    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    if ( val.type == lune_stem_type_nil )
    {
        // nil の場合は、削除

        lune_MapIterator it = pMap->find( key );
        if ( it != pMap->end() ) {
            DECRE_STEM( _pEnv, it->first );
            DECRE_STEM( _pEnv, it->second );
            
            pMap->erase( it );
        }
        
    }
    else if ( pMap->find( key ) == pMap->end() )
    {
        // 次の val は追加しない
        // - 登録済み
        
        // pVal の参照カウントをインクリメントするため、 work に setq する
        SETQ( &key );
        SETQ( &val );
    
        (*pMap)[ key ] = val;
    }
}

lune_stem_t lune_mtd_Map_get(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t key )
{
    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    lune_MapIterator it = pMap->find( key );
    
    if ( it == pMap->end() ) {
        return lune_global.nilStem;
    }

    return it->second;
}

lune_stem_t lune_mtd_Map_createKeyList(
    lune_env_t * _pEnv, lune_any_t * pObj )
{
    lune_stem_t pList = lune_class_List_new( _pEnv );

    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    lune_MapIterator it;
    lune_MapIterator end = pMap->end();
    
    for ( it = pMap->begin(); it != end; it++ ) {
        lune_mtd_List_insert( _pEnv, pList.val.pAny, it->first );
    }

    return pList;
}



/**
 * Map を生成する
 *
 * @param pMap 生成する Map の要素
 * @return 生成した Map
 */
lune_stem_t _lune_createMap(
    LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdEntry_t * pEntry )
{
    lune_stem_t pAny = lune_class_Map_new( _pEnv );

    for ( ; pEntry->key.type != lune_imdType_sentinel; pEntry++ ) {
        lune_mtd_Map_add(
            _pEnv, pAny.val.pAny,
            _lune_createImmediateVal( LUNE_DEBUG_POS, _pEnv, &pEntry->key ),
            _lune_createImmediateVal( LUNE_DEBUG_POS, _pEnv, &pEntry->val ) );
    }
    return pAny;
}

