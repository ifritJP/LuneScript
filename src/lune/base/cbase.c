#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>

typedef struct __lune_stem_t __lune_stem_t;

typedef struct __lune_block_t __lune_block_t;

#define __LUNE_BLOCK_MAX_DEPTH 10000
#define __LUNE_STEM_POOL_MAX_NUM 100000

#define __lune_set_block_stem( BLOCK, INDEX, STEM ) \
    (BLOCK)->pStemBuf[ INDEX ] = STEM

/**
STEM 型の値 VAL を、 変数 SYM に代入する。

- 参照カウントインクリメント
- unassignStem から除外
 */
#define __lune_setq( SYM, VAL )                 \
    SYM = VAL;                                  \
    SYM->refCount++;                            \
    __lune_rmFromList( SYM );

typedef struct {
    __lune_stem_t * stemPPool[ __LUNE_STEM_POOL_MAX_NUM ];
    __lune_stem_t * pNoneStem;
    int useStemPoolNum;
    __lune_block_t * pBlockQueue;
    int blockDepth;
    int allocNum;
} __lune_env_t;


typedef enum {
    __lune_value_type_none,
    __lune_value_type_int,
    __lune_value_type_real,
    __lune_value_type_str,
    __lune_value_type_class,
    __lune_value_type_ddd,
} __lune_value_type_t;

typedef int __lune_int_t;
typedef double __lune_real_t;
typedef void * __lune_class_t;

typedef __lune_stem_t * __lune_form_t( __lune_env_t * pEnv, __lune_stem_t * pObj, ... );
typedef void __lune_gc_t( __lune_env_t * pEnv, __lune_stem_t * pObj, bool freeFlag );

typedef struct __mtd__Class_t {
    __lune_gc_t * _gc;
} __mtd__Class_t;

typedef struct __lune_Class_t {
    __mtd__Class_t * pMtd;
} __lune_Class_t;


typedef struct {
    const char * pStr;
    int len;
    bool staticFlag;
} __lune_str_t;

typedef struct {
    int len;
    __lune_stem_t ** pStemList;
} __lune_ddd_t;

struct __lune_stem_t {
    __lune_value_type_t type;
    int refCount;
    union {
        __lune_int_t intVal;
        __lune_real_t realVal;
        __lune_str_t str;
        __lune_ddd_t ddd;
        __lune_class_t classVal;
    } val;
    struct __lune_stem_t * pNext;
    struct __lune_stem_t * pPrev;
};

struct __lune_block_t {
    void * pStackAddr;
    int blockDepth;
    int len;
    __lune_stem_t ** pStemBuf;
    __lune_stem_t unassignStemTop;
};




#define __lune_add2list( TOP, STEM )    \
        STEM->pNext = TOP;              \
        STEM->pPrev = (TOP)->pPrev;     \
        (TOP)->pPrev->pNext = STEM;     \
        (TOP)->pPrev = STEM;

#define __lune_rmFromList( STEM )                       \
    if ( (STEM)->pNext != NULL ) {                      \
        (STEM)->pPrev->pNext = (STEM)->pNext;           \
        (STEM)->pNext->pPrev = (STEM)->pPrev;           \
        (STEM)->pNext = NULL;                           \
    }
  


__lune_stem_t * __lune_alloc_stem( __lune_env_t * pEnv, __lune_value_type_t type ) {
    __lune_stem_t * pStem = (__lune_stem_t *)malloc( sizeof( __lune_stem_t ) );
    pStem->type = type;
    pStem->refCount = 0;

    __lune_block_t * pBlock = &pEnv->pBlockQueue[ pEnv->blockDepth - 1 ];
    __lune_add2list( &pBlock->unassignStemTop, pStem );
    pEnv->allocNum++;

    return pStem;
}


void __lune_release_stem( __lune_env_t * pEnv, __lune_stem_t * pStem ) {
    free( pStem );
    pEnv->allocNum--;
}

void __lune_decre_ref( __lune_env_t * pEnv, __lune_stem_t * pStem ) {
    pStem->refCount--;
    if ( pStem->refCount == 0 ) {
        switch ( pStem->type ) {
        case __lune_value_type_str:
            if ( !pStem->val.str.staticFlag ) {
                free( (char *)pStem->val.str.pStr );
            }
            break;
        case __lune_value_type_class:
            if ( ((__lune_Class_t*)pStem->val.classVal)->pMtd->_gc != NULL ) {
                ((__lune_Class_t*)pStem->val.classVal)->pMtd->_gc( pEnv, pStem, true );
            }
            break;
        }
        __lune_release_stem( pEnv, pStem );
    }
}

void __lune_enter_block( __lune_env_t * pEnv, int stemVerNum )
{
    int dummy;
    __lune_block_t * pBlock = &pEnv->pBlockQueue[ pEnv->blockDepth ];
    pEnv->blockDepth++;

    pBlock->unassignStemTop.pPrev = &pBlock->unassignStemTop;
    pBlock->unassignStemTop.pNext = &pBlock->unassignStemTop;
    
    pBlock->pStemBuf = pEnv->stemPPool + pEnv->useStemPoolNum;
    pBlock->len = stemVerNum;
    pBlock->blockDepth = pEnv->blockDepth;
    pBlock->pStackAddr = &dummy;
    
    pEnv->useStemPoolNum += stemVerNum;
}

void __lune_leave_block( __lune_env_t * pEnv )
{
    pEnv->blockDepth--;
    
    __lune_block_t * pBlock = &pEnv->pBlockQueue[ pEnv->blockDepth ];

    int index;
    for ( index = 0; index < pBlock->len; index++ ) {
        __lune_decre_ref( pEnv, pBlock->pStemBuf[ index ] );
    }

    __lune_stem_t * pWork = pBlock->unassignStemTop.pPrev;
    while ( pWork != &pBlock->unassignStemTop ) {
        pWork->refCount = 1;
        __lune_stem_t * pPrev = pWork->pPrev;

        __lune_decre_ref( pEnv, pWork );
        pWork = pPrev;
    }

    pEnv->useStemPoolNum -= pBlock->len;
}

void __lune_enter_func( __lune_env_t * pEnv, int num, int argNum, ... )
{
    __lune_enter_block( pEnv, num );
    __lune_block_t * pBlock = &pEnv->pBlockQueue[ pEnv->blockDepth - 1 ];


    va_list ap;
    va_start( ap, argNum );

    int index;
    for ( index = 0; index < argNum; index++ ) {
        __lune_stem_t * pStem = va_arg( ap, __lune_stem_t * );
        __lune_setq( pBlock->pStemBuf[ index ], pStem );
    }
    va_end(ap);
}

__lune_stem_t * __lune_createDDD( __lune_env_t * pEnv, int num, ... ) {
    __lune_stem_t * pDDDStem = __lune_alloc_stem( pEnv, __lune_value_type_ddd );
    __lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    pDDD->len = num;
    pDDD->pStemList = (__lune_stem_t **)malloc( sizeof( __lune_stem_t * ) * num );

    va_list ap;
    va_start( ap, num );

    int index;
    for ( index = 0; index < num; index++ ) {
        __lune_stem_t * pStem = va_arg( ap, __lune_stem_t * );
        pDDD->pStemList[ index ] = pStem;
    }
    va_end(ap);

    return pDDDStem;
}

__lune_stem_t * __lune_int2stem( __lune_env_t * pEnv, __lune_int_t val ) {
    __lune_stem_t * pStem = __lune_alloc_stem( pEnv, __lune_value_type_int );
    pStem->val.intVal = val;
    return pStem;
}

__lune_str_t __lune_createLiteralStr( const char * pStr ) {
    __lune_str_t str;
    str.len = strlen( pStr );
    str.pStr = pStr;
    str.staticFlag = true;
    return str;
}

__lune_stem_t * __lune_str2stem( __lune_env_t * pEnv, __lune_str_t val ) {
    __lune_stem_t * pStem = __lune_alloc_stem( pEnv, __lune_value_type_str );
    pStem->val.str = val;
    return pStem;
}

__lune_env_t * __lune_createEnv() {
    __lune_env_t * pEnv = (__lune_env_t *)malloc( sizeof( __lune_env_t ) );
    pEnv->useStemPoolNum = 0;
    pEnv->allocNum = 0;
    pEnv->blockDepth = 0;
    pEnv->pBlockQueue =
        (__lune_block_t *)malloc( sizeof( __lune_block_t ) * __LUNE_BLOCK_MAX_DEPTH );

    __lune_enter_block( pEnv, 0 );
    pEnv->pNoneStem = __lune_alloc_stem( pEnv, __lune_value_type_none );
}

__lune_env_t * __lune_deleteEnv( __lune_env_t * pEnv ) {
    __lune_leave_block( pEnv );

    printf( "allocNum = %d\n", pEnv->allocNum );
    printf( "useStemPoolNum = %d\n", pEnv->useStemPoolNum );
    printf( "blockDepth = %d\n", pEnv->blockDepth );

    free( pEnv->pBlockQueue );
    free( pEnv );
}

void __lune_print( __lune_env_t * pEnv, __lune_stem_t * pArg ) {

    __lune_enter_func( pEnv, 1, 1, pArg );

    __lune_ddd_t * pDDD = &pArg->val.ddd;
    int index;
    for ( index = 0; index < pDDD->len; index++ ) {
        if ( index > 0 ) {
            printf( "\t" );
        }
        __lune_stem_t * pStem = pDDD->pStemList[ index ];
        switch ( pStem->type ) {
        case __lune_value_type_int:
            printf( "%d", pStem->val.intVal );
            break;
        case __lune_value_type_real:
            printf( "%g", pStem->val.realVal );
            break;
        case __lune_value_type_str:
            printf( "%s", pStem->val.str.pStr );
            break;
        default:
            printf( "unknown type -- %d", pStem->type );
            break;
        }
    }
    printf( "\n" );

    __lune_leave_block( pEnv );
}

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
    __lune_form_t * func;
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


static void __mtd_Test_gc( __lune_env_t * pEnv, __lune_stem_t * pObj, bool freeFlag );
static __lune_stem_t * __mtd_Test_func( __lune_env_t * pEnv, __lune_stem_t * pObj );

__mtd_Test_t __mtd_Test = {
    __mtd_Test_gc,
    (__lune_form_t*)__mtd_Test_func
};


void __lune_class_Test_init( __lune_env_t * pEnv, Test * pObj, int val ) {
    pObj->val = val;
    __lune_setq( pObj->val2, __lune_int2stem( pEnv, 123 + val ) );
}


__lune_stem_t * __lune_class_Test_new( __lune_env_t * pEnv, int val ) {
    __lune_stem_t * pStem = __lune_alloc_stem( pEnv, __lune_value_type_class );
    Test * pObj = malloc( sizeof( Test ) );
    pStem->val.classVal = pObj;
    pObj->pMtd = &__mtd_Test;

    __lune_class_Test_init( pEnv, pObj, val );

    return pStem;
}

static void __mtd_Test_gc( __lune_env_t * pEnv, __lune_stem_t * pObj, bool freeFlag )
{
    printf( "%s\n", __func__ );
    __lune_decre_ref( pEnv, __lune_obj_Test( pObj )->val2 );

    if ( freeFlag ) {
        free( __lune_obj_Test( pObj ) );
    }
}


static __lune_stem_t * __mtd_Test_func( __lune_env_t * pEnv, __lune_stem_t * pObj )
{
    __lune_print(
        pEnv,
        __lune_createDDD(
            pEnv, 3,
            __lune_str2stem( pEnv, __lune_createLiteralStr( __func__ ) ),
            __lune_int2stem( pEnv, __lune_obj_Test( pObj )->val ),
            __lune_obj_Test( pObj )->val2 ) );
    return pEnv->pNoneStem;
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
    __lune_form_t * _super_func;
    __lune_form_t * func;
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


static void __mtd_Sub_gc( __lune_env_t * pEnv, __lune_stem_t * pObj, bool freeFlag );
static __lune_stem_t * __mtd_Sub_func( __lune_env_t * pEnv, __lune_stem_t * pObj );

__mtd_Sub_t __mtd_Sub = {
    __mtd_Sub_gc,
    (__lune_form_t *)__mtd_Test_func,
    (__lune_form_t *)__mtd_Sub_func
};

void __lune_class_Sub_init( __lune_env_t * pEnv, Sub * pObj, __lune_stem_t * val ) {
    __lune_setq( pObj->val3, val );
}

__lune_stem_t * __lune_class_Sub_new(
    __lune_env_t * pEnv, int val, __lune_stem_t * val3 )
{
    __lune_stem_t * pStem = __lune_alloc_stem( pEnv, __lune_value_type_class );
    Sub * pObj = malloc( sizeof( Sub ) );
    pStem->val.classVal = pObj;
    pObj->pMtd = &__mtd_Sub;

    __lune_class_Test_init( pEnv, (Test *)pObj, val );
    __lune_class_Sub_init( pEnv, pObj, val3 );

    return pStem;
}

static void __mtd_Sub_gc( __lune_env_t * pEnv, __lune_stem_t * pObj, bool freeFlag )
{
    
    printf( "%s\n", __func__ );
    __lune_decre_ref( pEnv, __lune_obj_Sub( pObj )->val3 );

    __mtd_Test_gc( pEnv, pObj, false );

    if ( freeFlag ) {
        free( __lune_obj_Sub( pObj ) );
    }
}


static __lune_stem_t * __mtd_Sub_func( __lune_env_t * pEnv, __lune_stem_t * pObj )
{
    __mtd_Test_func( pEnv, pObj );
    
    __lune_print(
        pEnv,
        __lune_createDDD(
            pEnv, 1,
            __lune_obj_Sub( pObj )->val3 ) );
    return pEnv->pNoneStem;
}



void test( __lune_env_t * pEnv ) {
    __lune_enter_block( pEnv, 2 );
    __lune_block_t * pBlock = &pEnv->pBlockQueue[ pEnv->blockDepth - 1 ];

    __lune_stem_t * test = __lune_class_Test_new( pEnv, 10 );
    __lune_set_block_stem( pBlock, 0, test );
    __lune_mtd_Test( test )->func( pEnv, test );

    __lune_stem_t * sub = __lune_class_Sub_new(
        pEnv, 20, __lune_str2stem( pEnv, __lune_createLiteralStr( "xyz" ) ));
    __lune_set_block_stem( pBlock, 1, sub );
    __lune_mtd_Sub( sub )->func( pEnv, sub );
    
    __lune_print(
        pEnv,
        __lune_createDDD(
            pEnv, 3,
            __lune_int2stem( pEnv, 1 ),
            __lune_int2stem( pEnv, 2 ),
            __lune_str2stem( pEnv, __lune_createLiteralStr( "abc" ) )));
    
    __lune_leave_block( pEnv );
}


int main() {
    __lune_env_t * pEnv = __lune_createEnv();
    
    test( pEnv );

    __lune_deleteEnv( pEnv );
}
