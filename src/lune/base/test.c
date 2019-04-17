#include <lunescript.h>


/**
class Test {
   pri let val:int;
   pub fn func() {
      print( self.val );
   }
}
*/
typedef struct __mtd_Test_t {
    __lune_gc_t * _gc;
    __lune_method_t * func;
} __mtd_Test_t;

typedef struct Test {
    __mtd_Test_t * pMtd;
    __lune_int_t val;
    __lune_stem_t * val2;
} Test;

#define __lune_mtd_Test( OBJ ) \
    ((Test*)OBJ->val.classVal)->pMtd
#define __lune_obj_Test( OBJ ) \
    ((Test*)OBJ->val.classVal)


static void __mtd_Test_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag );
static __lune_stem_t * __mtd_Test_func( __lune_env_t * _pEnv, __lune_stem_t * pObj );

__mtd_Test_t __mtd_Test = {
    __mtd_Test_gc,
    (__lune_method_t*)__mtd_Test_func
};


void __class_Test_init( __lune_env_t * _pEnv, Test * pObj, int val ) {
    pObj->val = val;
    __lune_setQ( pObj->val2, __lune_int2stem( _pEnv, 123 + val ) );
}


__lune_stem_t * __class_Test_new( __lune_env_t * _pEnv, int val ) {
    __lune_class_new_( _pEnv, Test, pStem, pObj );

    __class_Test_init( _pEnv, pObj, val );

    return pStem;
}

static void __mtd_Test_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag )
{
    printf( "%s\n", __func__ );
    __lune_decre_ref( _pEnv, __lune_obj_Test( pObj )->val2 );

    if ( freeFlag ) {
        __lune_class_del( _pEnv, __lune_obj_Test( pObj ) );
    }
}


static __lune_stem_t * __mtd_Test_func( __lune_env_t * _pEnv, __lune_stem_t * pObj )
{
    __lune_print(
        _pEnv,
        __lune_createDDD(
            _pEnv, 3,
            __lune_str2stem( _pEnv, __lune_createLiteralStr( __func__ ) ),
            __lune_int2stem( _pEnv, __lune_obj_Test( pObj )->val ),
            __lune_obj_Test( pObj )->val2 ) );
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
typedef struct __mtd_Sub_t {
    __lune_gc_t * _gc;
    __lune_method_t * _super_func;
    __lune_method_t * func;
} __mtd_Sub_t;

typedef struct Sub {
    __mtd_Sub_t * pMtd;
    __lune_int_t val;
    __lune_stem_t * val2;
    __lune_stem_t * val3;
} Sub;

#define __lune_mtd_Sub( OBJ ) \
    ((Sub*)OBJ->val.classVal)->pMtd
#define __lune_obj_Sub( OBJ ) \
    ((Sub*)OBJ->val.classVal)


static void __mtd_Sub_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag );
static __lune_stem_t * __mtd_Sub_func( __lune_env_t * _pEnv, __lune_stem_t * pObj );

__mtd_Sub_t __mtd_Sub = {
    __mtd_Sub_gc,
    (__lune_method_t *)__mtd_Test_func,
    (__lune_method_t *)__mtd_Sub_func
};

void __class_Sub_init( __lune_env_t * _pEnv, Sub * pObj, __lune_stem_t * val ) {
    __lune_setQ( pObj->val3, val );
}

__lune_stem_t * __class_Sub_new(
    __lune_env_t * _pEnv, int val, __lune_stem_t * val3 )
{
    __lune_class_new_( _pEnv, Sub, pStem, pObj );

    __class_Test_init( _pEnv, (Test *)pObj, val );
    __class_Sub_init( _pEnv, pObj, val3 );

    return pStem;
}

static void __mtd_Sub_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag )
{
    
    printf( "%s\n", __func__ );
    __lune_decre_ref( _pEnv, __lune_obj_Sub( pObj )->val3 );

    __mtd_Test_gc( _pEnv, pObj, false );

    if ( freeFlag ) {
        __lune_class_del( _pEnv, __lune_obj_Sub( pObj ) );
    }
}


static __lune_stem_t * __mtd_Sub_func( __lune_env_t * _pEnv, __lune_stem_t * pObj )
{
    __mtd_Test_func( _pEnv, pObj );
    
    __lune_print(
        _pEnv,
        __lune_createDDD(
            _pEnv, 1,
            __lune_obj_Sub( pObj )->val3 ) );
    return _pEnv->pNoneStem;
}

static __lune_stem_t * __form_test( __lune_env_t * _pEnv, __lune_stem_t * pForm )
{
    __lune_block_t * pBlock = __lune_enter_func( _pEnv, 2, 0 );

    __lune_stem_t * test;
    __lune_initVal( test, pBlock, 0, __class_Test_new( _pEnv, 10 ) );
    __lune_mtd_Test( test )->func( _pEnv, test );
    

    __lune_stem_t * sub;
    __lune_initVal(
        sub, pBlock, 1,
        __class_Sub_new(
            _pEnv, 20, __lune_str2stem( _pEnv, __lune_createLiteralStr( "xyz" ) )) );
    __lune_mtd_Sub( sub )->func( _pEnv, sub );
    
    
    __lune_print(
        _pEnv,
        __lune_createDDD(
            _pEnv, 3,
            __lune_int2stem( _pEnv, 1 ),
            __lune_int2stem( _pEnv, 2 ),
            __lune_str2stem( _pEnv, __lune_createLiteralStr( "abc" ) )));

    __lune_stem_t * _ret = __lune_setRet( _pEnv, __lune_int2stem( _pEnv, 100 ) );
    
    __lune_leave_block( _pEnv );

    return _ret;
}

static __lune_stem_t * __form_test2(
    __lune_env_t * _pEnv, __lune_stem_t * _pForm, __lune_stem_t * pVal )
{
    __lune_block_t * pBlock = __lune_enter_func( _pEnv, 2, 1, pVal );


    pVal->val.form.pFunc( _pEnv, pVal );


    __lune_stem_t * pList;
    __lune_initVal( pList, pBlock, 1, __class_List_new( _pEnv ) );

    __lune_mtd_List( pList )->insert( _pEnv, pList, pVal );
    
    __lune_leave_block( _pEnv );

    return _pEnv->pNoneStem;
}


void __lune_init_test( __lune_env_t * _pEnv )
{
    __lune_block_t * pBlock = __lune_enter_func( _pEnv, 1, 0 );

    
    __lune_stem_t * pVal;
    __lune_initVal( pVal, pBlock, 0, __form_test( _pEnv, NULL ) );

    __lune_print( _pEnv, __lune_createDDD( _pEnv, 1, pVal ) );

    __form_test2( _pEnv, NULL, __lune_func2stem( _pEnv, (__lune_func_t *)__form_test, 0 ) );

    __lune_leave_block( _pEnv );
}
