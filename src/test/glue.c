#include "test_glueTest_glue.h"
#include <stdio.h>

typedef struct {
    int val;
} hoge_t;

int hoge_create( lua_State * pLua, int val )
{
    hoge_t * pInfo = lns_glue_new_test_glueTest( pLua, sizeof( hoge_t ) );
    if ( pInfo == NULL ) {
        return 0;
    }
    pInfo->val = val;
    return 1;
}

int hoge_func1( lua_State * pLua, int val, const char * txt, int size_txt )
{
    printf( "func1: %d, %s, %d\n", val, txt, size_txt );
    lua_pushinteger( pLua, 1 );
    return 1;
}

int hoge_func2( lua_State * pLua, int val, int * val2,
                double * val3, const char * val4, int size_val4 )
/* int hoge_func2( lua_State * pLua, int val ) */
{
    hoge_t * pInfo = lns_glue_get_test_glueTest( pLua, 1 );
    char val2buf[10] = "nil";
    char val3buf[10] = "nil";
    if ( val2 != NULL ) {
        sprintf( val2buf, "%d", *val2 );
    }
    if ( val3 != NULL ) {
        sprintf( val3buf, "%g", *val3 );
    }
    printf( "func2: %d, %d, %s, %s, %s(%d) \n",
            pInfo->val, val,
            val2buf, val3buf, val4 ? val4 : "nil", size_val4 );
    lua_pushinteger( pLua, 2 );
    return 1;
}
