#include <lunescript.h>
#include <vector>
#include <algorithm>
#include <set>
#include <map>

using namespace std;


/// ================ List ======================

#define lune_obj_List( OBJ ) \
    ((lune_List_t*)OBJ->val.classVal)

#define lune_obj_List_obj( OBJ ) \
    ((vector<lune_stem_t*>*)((lune_List_t*)OBJ->val.classVal)->pObj)


#define lune_ListClass vector<lune_stem_t*>
#define lune_ListIterator vector<lune_stem_t*>::iterator


static void lune_mtd_List_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );
static lune_stem_t * lune_mtd_List_insert(
    lune_env_t * _pEnv, lune_stem_t * pListStem, lune_stem_t * pVal );
static lune_stem_t * lune_mtd_List_refAt(
    lune_env_t * _pEnv, lune_stem_t * pListStem, int index );
static lune_stem_t * lune_mtd_List_unpack(
    lune_env_t * _pEnv, lune_stem_t * pListStem );
static lune_stem_t * lune_mtd_List_sort(
    lune_env_t * _pEnv, lune_stem_t * pListStem, lune_stem_t * pForm );

lune_mtd_List_t lune_mtd_List = {
    lune_mtd_List_gc,
    (lune_method_t*)lune_mtd_List_insert,
    (lune_method_t*)lune_mtd_List_refAt,
    (lune_method_t*)lune_mtd_List_unpack,
    (lune_method_t*)lune_mtd_List_sort,
};


void lune_class_List_init( lune_env_t * _pEnv, lune_List_t * pObj )
{
    pObj->pObj = new vector<lune_stem_t*>();
}


lune_stem_t * lune_class_List_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_List_t, List, pStem, pObj );
    
    lune_class_List_init( _pEnv, pObj );

    return pStem;
}

lune_stem_t * lune_List_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem )
{
    lune_stem_t * pStem = lune_class_List_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDStem->val.ddd.len; index++ ) {
        lune_stem_t * pWork;
        lune_setQ( (&pWork), lune_fromDDD( pDDDStem, index ) );
    
        lune_obj_List_obj( pStem )->push_back( pWork );
    }
    
    return pStem;
}

/**
List インスタンスの開放処理
 */
static void lune_mtd_List_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag )
{
    LUNE_DEBUG_CALL_LOG;

    lune_ListIterator it;
    lune_ListIterator end = lune_obj_List_obj( pObj )->end();
    
    for ( it = lune_obj_List_obj( pObj )->begin(); it != end; it++ )
    {
        lune_decre_ref( _pEnv, *it );
    }
    
    delete lune_obj_List_obj( pObj );

    if ( freeFlag ) {
        lune_class_del( _pEnv, lune_obj_List( pObj ) );
    }
}

/**
List.insert()  処理
 */
static lune_stem_t * lune_mtd_List_insert(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal )
{
    // pVal の参照カウントをインクリメントするため、 work に setq する
    lune_stem_t * pWork;
    lune_setQ( (&pWork), pVal );
    
    lune_obj_List_obj( pObj )->push_back( pVal );
    
    return NULL;
}

/**
リストの要素アクセス。
 */
static lune_stem_t * lune_mtd_List_refAt(
    lune_env_t * _pEnv, lune_stem_t * pObj, int index )
{
    return lune_obj_List_obj( pObj )->at( index - 1 );
}


/**
  List.unpack() 処理
 */
static lune_stem_t * lune_mtd_List_unpack( lune_env_t * _pEnv, lune_stem_t * pObj )
{
    lune_stem_t * pDDD =
        lune_createDDDOnly( _pEnv, lune_obj_List_obj( pObj )->size() );

    lune_ListIterator it;
    lune_ListIterator end = lune_obj_List_obj( pObj )->end();

    int index = 0;
    for ( it = lune_obj_List_obj( pObj )->begin(); it != end; it++, index++ )
    {
        lune_set2DDDArg( pDDD, index, *it );
    }

    return pDDD;
}

/**
sort() 用のデフォルト比較関数 (昇順でソート)

@param pVal1 比較対象1
@param pVal2 比較対象2
@return int 比較対象1 < 比較対象2 ならば 0 以外
 */
static bool lune_mtd__Cmp( lune_stem_t * pVal1, lune_stem_t * pVal2 )
{
    if ( pVal1->type != pVal2->type ) {
        return pVal1 < pVal2;
    }

    switch ( pVal1->type ) {
    case lune_value_type_int:
        return !(pVal1->val.intVal >= pVal2->val.intVal);
    case lune_value_type_real:
        return !(pVal1->val.realVal >= pVal2->val.realVal);
    case lune_value_type_str:
        return !(strcmp( pVal1->val.str.pStr, pVal2->val.str.pStr ) >= 0);
    default:
        return !(pVal1 >= pVal2);
    }
}

/**
sort() のユーザフォーム実行用関数

@param pVal1 比較対象1
@param pVal2 比較対象2
@return 比較対象1 < 比較対象2 ならば true
 */
static bool lune_mtd__CmpLune( lune_stem_t * pVal1, lune_stem_t * pVal2 )
{
    if ( pVal1->type != pVal2->type ) {
        return pVal1 < pVal2;
    }

    lune_env_t * pEnv = pVal1->pEnv;
    lune_stem_t * pRet;
    lune_setQ( (&pRet), pEnv->pSortCallback->val.form.pFunc( pEnv, pEnv->pSortCallback,
                                                            pVal1, pVal2 ) );

    int ret = pRet->val.intVal;
    lune_decre_ref( pEnv, pRet );
    
    return ret;
}

/**
List.sort() 
 */
static lune_stem_t * lune_mtd_List_sort(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pForm )
{
    if ( pForm == _pEnv->pNilStem ) {
        std::sort( lune_obj_List_obj( pObj )->begin(),
                   lune_obj_List_obj( pObj )->end(), lune_mtd__Cmp );
    }
    else {
        _pEnv->pSortCallback = pForm;
        std::sort( lune_obj_List_obj( pObj )->begin(),
                   lune_obj_List_obj( pObj )->end(),
                   lune_mtd__CmpLune );
        _pEnv->pSortCallback = NULL;
    }

    return NULL;
}

/// ================ Array ======================



/// ================ Set ======================
#define lune_SetClass set<lune_stem_t*,lune_SetComp>
#define lune_SetIterator lune_SetClass::iterator


class lune_SetComp {
public:
    bool operator()(const lune_stem_t * pVal1, const lune_stem_t * pVal2) {
        if ( pVal1->type != pVal2->type ) {
            return pVal1 < pVal2;
        }
        switch ( pVal1->type ) {
        case lune_value_type_int:
            return pVal1->val.intVal < pVal2->val.intVal;
        case lune_value_type_real:
            return pVal1->val.realVal < pVal2->val.realVal;
        case lune_value_type_bool:
            return pVal1->val.boolVal < pVal2->val.boolVal;
        case lune_value_type_str:
            return memcmp( pVal1->val.str.pStr, pVal2->val.str.pStr,
                           pVal1->val.str.len + 1 ) < 0;
        default:
            return pVal1 < pVal2;
        }
    }
};

#define lune_obj_Set( OBJ ) \
    ((lune_Set_t*)OBJ->val.classVal)

#define lune_obj_Set_obj( OBJ ) \
    ((lune_SetClass*)((lune_Set_t*)OBJ->val.classVal)->pObj)

#define lune_mtd_Set_has_( OBJ, VAL )                                          \
    (lune_obj_Set_obj( OBJ )->find( VAL ) != lune_obj_Set_obj( OBJ )->end())

struct lune_itSet_t {
    lune_SetIterator it;
    lune_SetIterator end;
};


static void lune_mtd_Set_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );


static lune_stem_t * lune_mtd_Set_add( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal );
static lune_stem_t * lune_mtd_Set_del( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal );
static lune_stem_t * lune_mtd_Set_has( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal );
static lune_stem_t * lune_mtd_Set_and_( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet );
static lune_stem_t * lune_mtd_Set_or_( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet );
static lune_stem_t * lune_mtd_Set_sub( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet );
static lune_stem_t * lune_mtd_Set_clone( lune_env_t * _pEnv, lune_stem_t * pObj );
static lune_stem_t * lune_mtd_Set_len( lune_env_t * _pEnv, lune_stem_t * pObj );


lune_mtd_Set_t lune_mtd_Set = {
    lune_mtd_Set_gc,
    (lune_method_t*)lune_mtd_Set_add,
    (lune_method_t*)lune_mtd_Set_del,
    (lune_method_t*)lune_mtd_Set_has,
    (lune_method_t*)lune_mtd_Set_and_,
    (lune_method_t*)lune_mtd_Set_or_,
    (lune_method_t*)lune_mtd_Set_sub,
    (lune_method_t*)lune_mtd_Set_clone,
    (lune_method_t*)lune_mtd_Set_len,
};


void lune_class_Set_init( lune_env_t * _pEnv, lune_Set_t * pObj )
{
    pObj->pObj = new lune_SetClass();
}


lune_stem_t * lune_class_Set_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_Set_t, Set, pStem, pObj );
    
    lune_class_Set_init( _pEnv, pObj );

    return pStem;
}

lune_stem_t * lune_Set_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem )
{
    lune_stem_t * pStem = lune_class_Set_new( _pEnv );

    int index;
    for ( index = 0; index < pDDDStem->val.ddd.len; index++ ) {
        lune_stem_t * pWork;
        lune_setQ( (&pWork), lune_fromDDD( pDDDStem, index ) );
    
        lune_obj_Set_obj( pStem )->insert( pWork );
    }
    
    return pStem;
}

/**
List インスタンスの開放処理
 */
static void lune_mtd_Set_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag )
{
    LUNE_DEBUG_CALL_LOG;

    lune_SetIterator it;
    lune_SetIterator end = lune_obj_Set_obj( pObj )->end();
    
    for ( it = lune_obj_Set_obj( pObj )->begin(); it != end; it++ )
    {
        lune_decre_ref( _pEnv, *it );
    }
    
    delete lune_obj_Set_obj( pObj );

    if ( freeFlag ) {
        lune_class_del( _pEnv, lune_obj_Set( pObj ) );
    }
}

lune_stem_t * lune_itSet_new( lune_env_t * _pEnv, lune_stem_t * pSet )
{
    lune_stem_t * pStem = lune_it_new(
        _pEnv, lune_value_type_itSet, new lune_itSet_t );
    pStem->val.itSet->it = lune_obj_Set_obj( pSet )->begin();
    pStem->val.itSet->end = lune_obj_Set_obj( pSet )->end();
    return pStem;
}

void lune_itSet_gc( lune_env_t * _pEnv, lune_stem_t * it )
{
    delete it->val.itSet;
}

void lune_itSet_inc( lune_env_t * _pEnv, lune_stem_t * it )
{
    it->val.itSet->it++;
}

bool lune_itSet_hasNext( lune_env_t * _pEnv, lune_stem_t * it, lune_stem_t ** ppVal )
{
    if ( it->val.itSet->it != it->val.itSet->end ) {
        *ppVal = *it->val.itSet->it;
        return true;
    }
    *ppVal = NULL;
    return false;
}

static lune_stem_t * lune_mtd_Set_add( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal )
{
    if ( pVal->type != lune_value_type_nil &&
         lune_obj_Set_obj( pObj )->find( pVal ) == lune_obj_Set_obj( pObj )->end() )
    {
        // 次の val は追加しない
        // - nil
        // - 登録済み
        
        // pVal の参照カウントをインクリメントするため、 work に setq する
        lune_stem_t * pWork;
        lune_setQ( (&pWork), pVal );
    
        lune_obj_Set_obj( pObj )->insert( pVal );
    }
    
    return NULL;
}

static lune_stem_t * lune_mtd_Set_del( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal )
{

    lune_obj_Set_obj( pObj )->erase( pVal );
    lune_decre_ref( _pEnv, pVal );
    
    return NULL;
}

static lune_stem_t * lune_mtd_Set_has( lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pVal )
{
    bool contains =
        lune_obj_Set_obj( pObj )->find( pVal ) != lune_obj_Set_obj( pObj )->end();
    return lune_bool2stem( _pEnv, contains );
}

static lune_stem_t * lune_mtd_Set_and_(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet )
{
    lune_ListClass list;
    {
        // A - A ∩ B の要素を list に格納
        lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
        lune_SetIterator end = lune_obj_Set_obj( pObj )->begin();
    
        for ( ; it != end; it++ ) {
            if ( !lune_mtd_Set_has_( pSet, *it ) ) {
                list.push_back( *it );
            }
        }
    }

    {
        // 集合から list の要素を除外
        lune_ListIterator it = lune_obj_List_obj( pObj )->begin();
        lune_ListIterator end = lune_obj_List_obj( pObj )->end();
        
        for ( ; it != end; it++ ) {
            lune_obj_Set_obj( pObj )->erase( *it );
            lune_decre_ref( _pEnv, *it );
        }
    }
    
    return pObj;
}

static lune_stem_t * lune_mtd_Set_or_(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet )
{
    lune_SetIterator it = lune_obj_Set_obj( pSet )->begin();
    lune_SetIterator end = lune_obj_Set_obj( pSet )->begin();
    
    for ( ; it != end; it++ ) {
        lune_mtd_Set_add( _pEnv, pObj, *it );
    }
    
    return pObj;
}

static lune_stem_t * lune_mtd_Set_sub(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pSet )
{
    lune_ListClass list;

    {
        // and 集合の要素を抽出
        lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
        lune_SetIterator end = lune_obj_Set_obj( pObj )->begin();
    
        for ( ; it != end; it++ ) {
            if ( lune_mtd_Set_has_( pSet, *it ) ) {
                list.push_back( *it );
            }
        }
    }

    {
        // 抽出した要素を除外
        lune_ListIterator it = lune_obj_List_obj( pObj )->begin();
        lune_ListIterator end = lune_obj_List_obj( pObj )->end();
        
        for ( ; it != end; it++ ) {
            lune_obj_Set_obj( pObj )->erase( *it );
            lune_decre_ref( _pEnv, *it );
        }
    }
    
    
    return pObj;
}

static lune_stem_t * lune_mtd_Set_clone( lune_env_t * _pEnv, lune_stem_t * pObj )
{
    lune_stem_t * pStem = lune_class_Set_new( _pEnv );

    lune_SetIterator it = lune_obj_Set_obj( pObj )->begin();
    lune_SetIterator end = lune_obj_Set_obj( pObj )->begin();
    
    for ( ; it != end; it++ ) {
        lune_stem_t * pWork;
        lune_setQ( (&pWork), *it );

        lune_obj_Set_obj( pStem )->insert( pWork );
    }
    
    return NULL;
}

static lune_stem_t * lune_mtd_Set_len( lune_env_t * _pEnv, lune_stem_t * pObj )
{
    return lune_int2stem( _pEnv, (lune_int_t)lune_obj_Set_obj( pObj )->size() );
}


/// ================ Map ======================
#define lune_MapClass map<lune_stem_t*,lune_stem_t*,lune_SetComp>
#define lune_MapIterator lune_MapClass::iterator

#define lune_obj_Map( OBJ ) \
    ((lune_Map_t*)OBJ->val.classVal)

#define lune_obj_Map_obj( OBJ ) \
    ((lune_MapClass*)((lune_Map_t*)OBJ->val.classVal)->pObj)

#define lune_mtd_Map_has_( OBJ, VAL )                                          \
    (lune_obj_Map_obj( OBJ )->find( VAL ) != lune_obj_Map_obj( OBJ )->end())

struct lune_itMap_t {
    lune_MapIterator it;
    lune_MapIterator end;
};


static void lune_mtd_Map_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );


static lune_stem_t * lune_mtd_Map_add( lune_env_t * _pEnv, lune_stem_t * pObj,
                                      lune_stem_t * pKey, lune_stem_t * pVal );
static lune_stem_t * lune_mtd_Map_get( lune_env_t * _pEnv, lune_stem_t * pObj,
                                      lune_stem_t * pKey );


lune_mtd_Map_t lune_mtd_Map = {
    lune_mtd_Map_gc,
    (lune_method_t*)lune_mtd_Map_add,
    (lune_method_t*)lune_mtd_Map_get,
};


void lune_class_Map_init( lune_env_t * _pEnv, lune_Map_t * pObj )
{
    pObj->pObj = new lune_MapClass();
}


lune_stem_t * lune_class_Map_new( lune_env_t * _pEnv )
{
    lune_class_new2_( _pEnv, lune_Map_t, Map, pStem, pObj );
    
    lune_class_Map_init( _pEnv, pObj );
    
    return pStem;
}

lune_stem_t * lune_Map_ctor( lune_env_t * _pEnv, lune_stem_t * pDDDStem )
{
    lune_stem_t * pStem = lune_class_Map_new( _pEnv );
    lune_MapClass * pMap = lune_obj_Map_obj( pStem );

    int index;
    for ( index = 0; index < pDDDStem->val.ddd.len; index += 2 ) {
        lune_stem_t * pKey;
        lune_setQ( (&pKey), lune_fromDDD( pDDDStem, index ) );
        lune_stem_t * pVal;
        lune_setQ( (&pVal), lune_fromDDD( pDDDStem, index + 1 ) );

        pMap->insert( lune_MapClass::value_type( pKey, pVal ) );
    }
    
    return pStem;
}

/**
List インスタンスの開放処理
 */
static void lune_mtd_Map_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag )
{
    LUNE_DEBUG_CALL_LOG;

    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    
    lune_MapIterator it;
    lune_MapIterator end = pMap->end();
    
    for ( it = pMap->begin(); it != end; it++ ) {
        lune_decre_ref( _pEnv, it->first );
        lune_decre_ref( _pEnv, it->second );
    }
    
    delete pMap;

    if ( freeFlag ) {
        lune_class_del( _pEnv, lune_obj_Map( pObj ) );
    }
}

lune_stem_t * lune_itMap_new( lune_env_t * _pEnv, lune_stem_t * pMap )
{
    lune_stem_t * pStem = lune_it_new(
        _pEnv, lune_value_type_itMap, new lune_itMap_t );
    pStem->val.itMap->it = lune_obj_Map_obj( pMap )->begin();
    pStem->val.itMap->end = lune_obj_Map_obj( pMap )->end();

    return pStem;
}

void lune_itMap_gc( lune_env_t * _pEnv, lune_stem_t * it )
{
    delete it->val.itMap;
}

void lune_itMap_inc( lune_env_t * _pEnv, lune_stem_t * it )
{
    it->val.itMap->it++;
}

bool lune_itMap_hasNext( lune_env_t * _pEnv, lune_stem_t * it, lune_Map_entry_t * pEntry )
{
    if ( it->val.itMap->it != it->val.itMap->end ) {
        pEntry->pKey = it->val.itMap->it->first;
        pEntry->pVal = it->val.itMap->it->second;
        return true;
    }
    return false;
}

static lune_stem_t * lune_mtd_Map_add( lune_env_t * _pEnv, lune_stem_t * pObj,
                                      lune_stem_t * pKey, lune_stem_t * pVal )
{
    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    if ( pVal->type == lune_value_type_nil ) {
        // nil の場合は、削除

        lune_MapIterator it = pMap->find( pKey );
        if ( it != pMap->end() ) {
            lune_decre_ref( _pEnv, it->first );
            lune_decre_ref( _pEnv, it->second );
            
            pMap->erase( pVal );
        }
        
    }
    else if ( pMap->find( pKey ) == pMap->end() )
    {
        // 次の val は追加しない
        // - 登録済み
        
        // pVal の参照カウントをインクリメントするため、 work に setq する
        lune_stem_t * pWorkKey;
        lune_setQ( (&pWorkKey), pKey );
        lune_stem_t * pWorkVal;
        lune_setQ( (&pWorkVal), pVal );
    
        (*pMap)[ pKey ] = pVal;
    }
    
    return pObj;
}

static lune_stem_t * lune_mtd_Map_get(
    lune_env_t * _pEnv, lune_stem_t * pObj, lune_stem_t * pKey )
{
    lune_MapClass * pMap = lune_obj_Map_obj( pObj );
    lune_MapIterator it = pMap->find( pKey );
    
    if ( it == pMap->end() ) {
        return _pEnv->pNilStem;
    }

    return it->second;
}

