#include <lunescript.h>


/**
class Test {
   pri let val:int;
   pub fn func() {
      print( self.val );
   }
}
*/
typedef struct lune_mtd_Test_t {
    lune_gc_t * _gc;
    lune_method_t * func;
} lune_mtd_Test_t;

typedef struct Test {
    lune_mtd_Test_t * pMtd;
    lune_int_t val;
    lune_stem_t * val2;
} Test;

#define lune_mtd_Test( OBJ ) \
    ((Test*)OBJ->val.classVal)->pMtd
#define lune_obj_Test( OBJ ) \
    ((Test*)OBJ->val.classVal)


static void u_mtd_Test_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );
static lune_stem_t * u_mtd_Test_func( lune_env_t * _pEnv, lune_stem_t * pObj );

lune_mtd_Test_t lune_mtd_Test = {
    u_mtd_Test_gc,
    (lune_method_t*)u_mtd_Test_func
};


void u_class_Test_init( lune_env_t * _pEnv, Test * pObj, int val ) {
    pObj->val = val;
    lune_setQ( (&pObj->val2), lune_int2stem( _pEnv, 123 + val ) );
}


lune_stem_t * u_class_Test_new( lune_env_t * _pEnv, int val ) {
    lune_class_new_( _pEnv, Test, pStem, pObj );

    u_class_Test_init( _pEnv, pObj, val );

    return pStem;
}

static void u_mtd_Test_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag )
{
    LUNE_DEBUG_CALL_LOG;
    lune_decre_ref( _pEnv, lune_obj_Test( pObj )->val2 );

    if ( freeFlag ) {
        lune_class_del( _pEnv, lune_obj_Test( pObj ) );
    }
}


static lune_stem_t * u_mtd_Test_func( lune_env_t * _pEnv, lune_stem_t * pObj )
{
    lune_print(
        _pEnv, NULL,
        lune_createDDD(
            _pEnv, false, 3,
            lune_litStr2stem( _pEnv, __func__ ),
            lune_int2stem( _pEnv, lune_obj_Test( pObj )->val ),
            lune_obj_Test( pObj )->val2 ) );
    return _pEnv->pNoneStem;
}



/**
class Sub extend Test {
   pri let val:int;
   pub fn func() {
      print( self.val );
   }
}
*/
typedef struct lune_mtd_Sub_t {
    lune_gc_t * _gc;
    lune_method_t * _super_func;
    lune_method_t * func;
} lune_mtd_Sub_t;

typedef struct Sub {
    lune_mtd_Sub_t * pMtd;
    lune_int_t val;
    lune_stem_t * val2;
    lune_stem_t * val3;
} Sub;

#define lune_mtd_Sub( OBJ ) \
    ((Sub*)OBJ->val.classVal)->pMtd
#define lune_obj_Sub( OBJ ) \
    ((Sub*)OBJ->val.classVal)


static void u_mtd_Sub_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );
static lune_stem_t * u_mtd_Sub_func( lune_env_t * _pEnv, lune_stem_t * pObj );

lune_mtd_Sub_t lune_mtd_Sub = {
    u_mtd_Sub_gc,
    (lune_method_t *)u_mtd_Test_func,
    (lune_method_t *)u_mtd_Sub_func
};

void u_class_Sub_init( lune_env_t * _pEnv, Sub * pObj, lune_stem_t * val ) {
    lune_setQ( (&pObj->val3), val );
}

lune_stem_t * lune_class_Sub_new(
    lune_env_t * _pEnv, int val, lune_stem_t * val3 )
{
    lune_class_new_( _pEnv, Sub, pStem, pObj );

    u_class_Test_init( _pEnv, (Test *)pObj, val );
    u_class_Sub_init( _pEnv, pObj, val3 );

    return pStem;
}

static void u_mtd_Sub_gc( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag )
{
    
    LUNE_DEBUG_CALL_LOG;
    lune_decre_ref( _pEnv, lune_obj_Sub( pObj )->val3 );

    u_mtd_Test_gc( _pEnv, pObj, false );

    if ( freeFlag ) {
        lune_class_del( _pEnv, lune_obj_Sub( pObj ) );
    }
}


static lune_stem_t * u_mtd_Sub_func( lune_env_t * _pEnv, lune_stem_t * pObj )
{
    u_mtd_Test_func( _pEnv, pObj );
    
    lune_print(
        _pEnv, NULL,
        lune_createDDD(
            _pEnv, false, 1,
            lune_obj_Sub( pObj )->val3 ) );
    return _pEnv->pNoneStem;
}

static lune_stem_t * u_lune_form_test( lune_env_t * _pEnv, lune_stem_t * pForm )
{
    lune_block_t * pBlock = lune_enter_block( _pEnv, 2 );

    lune_var_t * test;
    lune_initVal( test, pBlock, 0, u_class_Test_new( _pEnv, 10 ) );
    lune_mtd_Test( test->pStem )->func( _pEnv, test->pStem );
    

    lune_var_t * sub;
    lune_initVal(
        sub, pBlock, 1,
        lune_class_Sub_new(
            _pEnv, 20, lune_litStr2stem( _pEnv, "xyz" )) );
    lune_mtd_Sub( sub->pStem )->func( _pEnv, sub->pStem );
    
    
    lune_print(
        _pEnv, NULL,
        lune_createDDD(
            _pEnv, false, 3,
            lune_int2stem( _pEnv, 1 ),
            lune_int2stem( _pEnv, 2 ),
            lune_litStr2stem( _pEnv, "abc" )));

    lune_stem_t * _ret = lune_setRet( _pEnv, lune_int2stem( _pEnv, 100 ) );
    
    lune_leave_block( _pEnv );

    return _ret;
}

static lune_stem_t * u_lune_form_test2(
    lune_env_t * _pEnv, lune_stem_t * _pForm, lune_stem_t * pVal )
{
    lune_block_t * pBlock = lune_enter_block( _pEnv, 2 );

    pVal->val.form.pFunc( _pEnv, pVal );


    lune_var_t * pList;
    {
        lune_imdList( list, lune_imdInt( 1000 ), lune_imdReal( 10.5 ) );
        lune_initVal( pList, pBlock, 1, lune_createList( _pEnv, &list ) );
    }
        //lune_class_List_new( _pEnv ) );

    lune_mtd_List( pList->pStem )->insert( _pEnv, pList->pStem, pVal );
    
    lune_print( _pEnv, NULL,
                lune_mtd_List( pList->pStem )->unpack( _pEnv, pList->pStem ) );
    
    lune_leave_block( _pEnv );

    return _pEnv->pNoneStem;
}

static lune_stem_t * u_lune_form_test3( lune_env_t * _pEnv, lune_stem_t * _pForm )
{
    lune_enter_block( _pEnv, 0 );

    lune_var_t * pVal = lune_form_closure( _pForm, 0 );
    pVal->pStem->val.intVal += 1000;
    
    
    lune_leave_block( _pEnv );

    return _pEnv->pNoneStem;
}

// fn test4( val1:int, val2:int ): int { return val1 + val2; }
static lune_stem_t * u_lune_form_test4(
    lune_env_t * _pEnv, lune_stem_t * _pForm, int val1, int val2 )
{
    lune_enter_block( _pEnv, 0 );

    lune_stem_t * pRet = lune_setRet( _pEnv, lune_int2stem( _pEnv, val1 + val2 ) );
    
    lune_leave_block( _pEnv );

    return pRet;
}

static lune_stem_t * u_lune_form__test4(
    lune_env_t * _pEnv, lune_stem_t * _pForm, lune_stem_t * pDDD )
{
    return u_lune_form_test4( _pEnv, _pForm,
                              pDDD->val.ddd.pStemList[ 0 ]->val.intVal,
                              pDDD->val.ddd.pStemList[ 1 ]->val.intVal );
}

// fn test5(): int, int { return 1, 2; }
static lune_stem_t * u_lune_form_test5(
    lune_env_t * _pEnv, lune_stem_t * _pForm )
{
    lune_enter_block( _pEnv, 0 );

    lune_stem_t * pRet = lune_setRet(
        _pEnv,
        lune_createDDD( _pEnv, false, 2,
                        lune_int2stem( _pEnv, 1 ),
                        lune_int2stem( _pEnv, 2 ) ) );
    
    lune_leave_block( _pEnv );

    return pRet;
}



static lune_stem_t * u_lune_form_comp(
    lune_env_t * _pEnv, lune_stem_t * _pForm,
    lune_stem_t * pVal1, lune_stem_t * pVal2 )
{
    return lune_int2stem( _pEnv, pVal1->val.intVal > pVal2->val.intVal );
}


void lune_init_test( lune_env_t * _pEnv )
{
    lune_block_t * pBlock = lune_enter_block( _pEnv, 5 );

    // let val = test();
    lune_var_t * pVal;
    lune_initVal( pVal, pBlock, 0, u_lune_form_test( _pEnv, NULL ) );

    // print( val );
    lune_print( _pEnv, NULL, lune_createDDD( _pEnv, false, 1, pVal->pStem ) );

    // test2( test );
    u_lune_form_test2(
        _pEnv, NULL,
        lune_func2stem( _pEnv, (lune_func_t *)u_lune_form_test, 0, false, 0 ) );


    // let val2 = 99;
    lune_var_t * pVal2;
    lune_initVal( pVal2, pBlock, 1, lune_int2stem( _pEnv, 0 ) );
    pVal2->pStem->val.intVal= 99;
    // fn () { val2 = val2 + 1000; } ();
    lune_stem_t * pClosure = lune_func2stem(
        _pEnv, (lune_func_t *)u_lune_form_test3, 1, pVal2, 0, false );
    lune_call_form( _pEnv, pClosure, 0 );
    // print( val2 );
    lune_print( _pEnv, NULL, lune_createDDD( _pEnv, false, 1, pVal2->pStem ) );
    
    
    

    // let list = [10,20,30];
    lune_var_t * pList;
    lune_initVal(
        pList, pBlock, 2, 
        lune_List_ctor(
            _pEnv, lune_createDDD(
                _pEnv, false, 3,
                lune_int2stem( _pEnv, 10 ),
                lune_int2stem( _pEnv, 20 ),
                lune_int2stem( _pEnv, 30 ) ) ) );
    
    
    
    // let list:List<int> = [];
    /* lune_stem_t * pList; */
    /* lune_initVal( pList, pBlock, 1, lune_class_List_new( _pEnv ) ); */

    /* // list.insert( 10 ), list.insert( 30 ), list.insert( 20 ) */
    /* lune_mtd_List( pList )->insert( _pEnv, pList, lune_int2stem( _pEnv, 10 ) ); */
    /* lune_mtd_List( pList )->insert( _pEnv, pList, lune_int2stem( _pEnv, 30 ) ); */
    /* lune_mtd_List( pList )->insert( _pEnv, pList, lune_int2stem( _pEnv, 20 ) ); */


    // list.sort();
    // print( list.unpack() );
    lune_mtd_List( pList->pStem )->sort( _pEnv, pList->pStem, _pEnv->pNilStem );
    lune_print( _pEnv, NULL, lune_mtd_List( pList->pStem )->unpack( _pEnv, pList->pStem ) );

    // list.sort( fn (val1:int,val2:int):bool { return val1 > val2; } );
    lune_mtd_List( pList->pStem )->sort(
        _pEnv, pList->pStem,
        lune_func2stem( _pEnv, (lune_func_t *)u_lune_form_comp, 2, false, 0 ) );
    // print( list.unpack() );
    lune_print( _pEnv, NULL, lune_mtd_List( pList->pStem )->unpack( _pEnv, pList->pStem ) );

    // foreach val in list { print( val ); }
    {
        lune_stem_t * itStem;
        lune_stem_t * pVal;
        for ( itStem = lune_itList_new( _pEnv, pList->pStem );
              lune_itList_hasNext( _pEnv, itStem, &pVal );
              lune_itList_inc( _pEnv, itStem ) )
        {
            lune_print( _pEnv, NULL, lune_createDDD( _pEnv, false, 1, pVal ) );
        }
        lune_it_delete( _pEnv, itStem );
    }
    

    

    
    // let set = (@ 100, 200, 300 )
    lune_var_t * pSet;
    lune_imdSet( set, lune_imdStr( "123" ), lune_imdStr( "456" ), lune_imdStr( "789" ) );
    lune_initVal(
        pSet, pBlock, 3,
        lune_createSet( _pEnv, set ) );
        /* lune_Set_ctor( */
        /*     _pEnv, lune_createDDD( */
        /*         _pEnv, false, 3, */
        /*         lune_litStr2stem( _pEnv, "100" ), */
        /*         lune_litStr2stem( _pEnv, "200" ), */
        /*         lune_litStr2stem( _pEnv, "300" ) ) ) ); */

    // set.add( 100 )
    lune_mtd_Set( pSet->pStem )->add(
        _pEnv, pSet->pStem, lune_litStr2stem( _pEnv, "456" ) );

    // foreach val in set { print( val ); }
    {
        lune_stem_t * itStem;
        lune_stem_t * pVal;
        for ( itStem = lune_itSet_new( _pEnv, pSet->pStem );
              lune_itSet_hasNext( _pEnv, itStem, &pVal ); lune_itSet_inc( _pEnv, itStem ) )
        {
            lune_print( _pEnv, NULL, lune_createDDD( _pEnv, false, 1, pVal ) );
        }
        lune_it_delete( _pEnv, itStem );
    }


    // let map = { "abc":123, "def":456 }
    lune_var_t * pMap;
    {
        lune_stem_t * pWorkVal = lune_int2stem( _pEnv, 456 );
        lune_imdMap( map,
                     { lune_imdStr( "abc" ), lune_imdInt( 123 ) },
                     { lune_imdStr( "def" ),  lune_imdStem( pWorkVal ) } );
        lune_initVal(
            pMap, pBlock, 4,
            lune_createMap(
                _pEnv, map ));
        /* lune_Map_ctor( */
        /*     _pEnv, lune_createDDD( */
        /*         _pEnv, false, 4, */
        /*         lune_litStr2stem( _pEnv, "abc" ), lune_int2stem( _pEnv, 123 ), */
        /*         lune_litStr2stem( _pEnv, "def" ), lune_int2stem( _pEnv, 456 ) ) ) ); */
    }

    // foreach val, key in map { print( key, val ); }
    {
        lune_stem_t * itStem;
        lune_Map_entry_t entry;
        for ( itStem = lune_itMap_new( _pEnv, pMap->pStem );
              lune_itMap_hasNext( _pEnv, itStem, &entry ); lune_itMap_inc( _pEnv, itStem ) )
        {
            lune_print(
                _pEnv, NULL, lune_createDDD( _pEnv, false, 2, entry.pKey, entry.pVal ) );
        }
        lune_it_delete( _pEnv, itStem );
    }


    // print( "multi-val", test4( 1, 2 ), test4( test5() ) )
    lune_print( _pEnv, NULL,
                lune_createDDD(
                    _pEnv, false, 3,
                    lune_litStr2stem( _pEnv, "multi-val" ),
                    u_lune_form_test4( _pEnv, NULL, 1, 2 ),
                    u_lune_form__test4(
                       _pEnv, NULL, u_lune_form_test5( _pEnv, NULL ) ) ) );
    
    
    lune_leave_block( _pEnv );
}
