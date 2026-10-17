/*
  LuneScript から変換した wasm (WasmGC) が利用する Lua ランタイム。

  Lua 5.3 を wasi-sdk で wasm にし、このファイルの関数を export する。
  変換した wasm は、この wasm の memory と関数を import して使う。

  呼び出し方:
   1. lnsl_push_builtin() で呼び出す関数を push する
   2. lnsl_push_xxx() で引数を push する
      (文字列は lnsl_buf() で得た領域に書き込んでから lnsl_push_str() する)
   3. lnsl_call() で呼び出す。戻り値の数が返る (エラー時は -1 で、メッセージが 1 つ残る)
   4. lnsl_kind() / lnsl_to_xxx() で戻り値を取得する
   5. lnsl_settop( 0 ) で戻り値を捨てる
*/

#include <stdint.h>
#include <stdlib.h>

#include "lua.h"
#include "lauxlib.h"
#include "lualib.h"

#define EXPORT(name) __attribute__((export_name(#name))) name

static lua_State * s_lua;

/* 文字列の受け渡し用の領域 */
static char * s_buf;
static size_t s_bufSize;

/* lnsl_to_str() で返した文字列の長さ */
static size_t s_strLen;

/* lnsl_push_builtin() で push する関数 (この順番が id になる) */
static const char * s_builtinList[][ 2 ] = {
    { "string", "format" },
    { "string", "find" },
    { "string", "gsub" },
    { NULL, "tostring" },
    { "string", "gmatch" },
    { "string", "match" },
    { NULL, "tonumber" },
};
#define BUILTIN_NUM (int)(sizeof( s_builtinList ) / sizeof( s_builtinList[ 0 ] ))

/* builtin の関数を入れたテーブルの registry の参照 */
static int s_builtinRef;

void EXPORT(lnsl_init)( void )
{
    s_lua = luaL_newstate();
    /* io・os・package などのライブラリはホスト側で提供するので開かない */
    static const luaL_Reg libList[] = {
        { "_G", luaopen_base },
        { LUA_TABLIBNAME, luaopen_table },
        { LUA_STRLIBNAME, luaopen_string },
        { LUA_MATHLIBNAME, luaopen_math },
        { LUA_UTF8LIBNAME, luaopen_utf8 },
        { NULL, NULL },
    };
    for ( const luaL_Reg * lib = libList; lib->func != NULL; lib++ ) {
        luaL_requiref( s_lua, lib->name, lib->func, 1 );
        lua_pop( s_lua, 1 );
    }

    lua_createtable( s_lua, BUILTIN_NUM, 0 );
    for ( int index = 0; index < BUILTIN_NUM; index++ ) {
        if ( s_builtinList[ index ][ 0 ] != NULL ) {
            lua_getglobal( s_lua, s_builtinList[ index ][ 0 ] );
            lua_getfield( s_lua, -1, s_builtinList[ index ][ 1 ] );
            lua_remove( s_lua, -2 );
        }
        else {
            lua_getglobal( s_lua, s_builtinList[ index ][ 1 ] );
        }
        lua_rawseti( s_lua, -2, index );
    }
    s_builtinRef = luaL_ref( s_lua, LUA_REGISTRYINDEX );
}

/* size バイト以上の受け渡し用の領域を返す */
char * EXPORT(lnsl_buf)( int32_t size )
{
    if ( (size_t)size > s_bufSize ) {
        size_t newSize = s_bufSize == 0 ? 256 : s_bufSize;
        while ( newSize < (size_t)size ) {
            newSize *= 2;
        }
        s_buf = realloc( s_buf, newSize );
        if ( s_buf == NULL ) {
            abort();
        }
        s_bufSize = newSize;
    }
    return s_buf;
}

int32_t EXPORT(lnsl_gettop)( void )
{
    return lua_gettop( s_lua );
}

void EXPORT(lnsl_settop)( int32_t top )
{
    lua_settop( s_lua, top );
}

void EXPORT(lnsl_push_builtin)( int32_t id )
{
    lua_rawgeti( s_lua, LUA_REGISTRYINDEX, s_builtinRef );
    lua_rawgeti( s_lua, -1, id );
    lua_remove( s_lua, -2 );
}

void EXPORT(lnsl_push_nil)( void )
{
    lua_pushnil( s_lua );
}

void EXPORT(lnsl_push_bool)( int32_t val )
{
    lua_pushboolean( s_lua, val );
}

void EXPORT(lnsl_push_int)( int64_t val )
{
    lua_pushinteger( s_lua, val );
}

void EXPORT(lnsl_push_real)( double val )
{
    lua_pushnumber( s_lua, val );
}

/* lnsl_buf() の領域の先頭 len バイトを文字列として push する */
void EXPORT(lnsl_push_str)( int32_t len )
{
    lua_pushlstring( s_lua, s_buf, len );
}

static int callFunc( lua_State * lua )
{
    lua_call( lua, lua_gettop( lua ) - 1, LUA_MULTRET );
    return lua_gettop( lua );
}

/*
  スタックの先頭から関数と nargs 個の引数を取り出して呼び出す。

  戻り値の数を返す。エラー時は -1 を返し、エラーメッセージを 1 つ残す。
  スタックは、関数を push する前の状態が空であることを前提とする。
*/
int32_t EXPORT(lnsl_call)( int32_t nargs )
{
    int base = lua_gettop( s_lua ) - nargs - 1;
    if ( lua_pcall( s_lua, nargs, LUA_MULTRET, 0 ) != LUA_OK ) {
        if ( lua_type( s_lua, -1 ) != LUA_TSTRING ) {
            luaL_tolstring( s_lua, -1, NULL );
            lua_remove( s_lua, -2 );
        }
        return -1;
    }
    return lua_gettop( s_lua ) - base;
}

/* 値の種類 (0: nil, 1: int, 2: real, 3: bool, 4: str, 5: その他) */
int32_t EXPORT(lnsl_kind)( int32_t index )
{
    switch ( lua_type( s_lua, index ) ) {
    case LUA_TNIL:
        return 0;
    case LUA_TNUMBER:
        return lua_isinteger( s_lua, index ) ? 1 : 2;
    case LUA_TBOOLEAN:
        return 3;
    case LUA_TSTRING:
        return 4;
    default:
        return 5;
    }
}

int64_t EXPORT(lnsl_to_int)( int32_t index )
{
    return lua_tointeger( s_lua, index );
}

double EXPORT(lnsl_to_real)( int32_t index )
{
    return lua_tonumber( s_lua, index );
}

int32_t EXPORT(lnsl_to_bool)( int32_t index )
{
    return lua_toboolean( s_lua, index );
}

/* 文字列の先頭のアドレスを返す。長さは lnsl_str_len() で取得する */
const char * EXPORT(lnsl_to_str)( int32_t index )
{
    return lua_tolstring( s_lua, index, &s_strLen );
}

int32_t EXPORT(lnsl_str_len)( void )
{
    return (int32_t)s_strLen;
}
