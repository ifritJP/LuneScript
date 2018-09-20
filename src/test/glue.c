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

int hoge_func2( lua_State * pLua, int val )
{
    hoge_t * pInfo = lns_glue_get_test_glueTest( pLua, 1 );
    printf( "func2: %d, %d\n", pInfo->val, val );
    lua_pushinteger( pLua, 2 );
    return 1;
}
