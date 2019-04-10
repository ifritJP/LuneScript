#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>

typedef struct __lune_stem_t __lune_stem_t;

typedef struct __lune_block_t __lune_block_t;

#define __LUNE_BLOCK_MAX_DEPTH 10000
#define __LUNE_STEM_POOL_MAX_NUM 100000

typedef struct {
    __lune_stem_t * stemPPool[ __LUNE_STEM_POOL_MAX_NUM ];
    int useStemPoolNum;
    __lune_block_t * pBlockQueue;
    int blockDepth;
    int allocNum;
} __lune_env_t;


typedef enum {
    __lune_value_type_int,
    __lune_value_type_real,
    __lune_value_type_str,
    __lune_value_type_ddd,
} __lune_value_type_t;

typedef int __lune_int_t;
typedef double __lune_real_t;

typedef struct {
    char * pStr;
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

#define __lune_rmFromList( STEM )               \
    (STEM)->pPrev->pNext = (STEM)->pNext;       \
    (STEM)->pNext->pPrev = (STEM)->pPrev;

  


__lune_env_t * __lune_createEnv() {
    __lune_env_t * pEnv = (__lune_env_t *)malloc( sizeof( __lune_env_t ) );
    pEnv->useStemPoolNum = 0;
    pEnv->allocNum = 0;
    pEnv->blockDepth = 0;
    pEnv->pBlockQueue =
        (__lune_block_t *)malloc( sizeof( __lune_block_t ) * __LUNE_BLOCK_MAX_DEPTH );
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
                free( pStem->val.str.pStr );
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
    
    pBlock->pStemBuf = pEnv->stemPPool + pEnv->useStemPoolNum + 1;
    pBlock->len = stemVerNum;
    pBlock->blockDepth = pEnv->blockDepth;
    pBlock->pStackAddr = &dummy;
    
    pEnv->useStemPoolNum += stemVerNum + 1;
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

    pEnv->useStemPoolNum -= (pBlock->len + 1);
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
        pBlock->pStemBuf[ index ] = pStem;
        __lune_rmFromList( pStem );
        pStem->refCount++;
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

__lune_str_t __lune_createLiteralStr( char * pStr ) {
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

void test( __lune_env_t * pEnv ) {
    __lune_enter_block( pEnv, 0 );
    
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

    printf( "allocNum = %d\n", pEnv->allocNum );
    printf( "useStemPoolNum = %d\n", pEnv->useStemPoolNum );
    printf( "blockDepth = %d\n", pEnv->blockDepth );
    
}
