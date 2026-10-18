/*
  LuneScript から変換した wasm (WasmGC) が利用する Lua ランタイム。

  Lua 5.3 と lnsrt.c (入出力などのランタイム) を wasi-sdk で 1 つの wasm にし、
  このファイルの関数を export する。
  変換した wasm は、この wasm の memory と関数 (モジュール名 "lnsrt") を import して使う。
  Lua の初期化は、lnsrt.c の main() から呼ばれる lnsr_initExt() で行う。

  呼び出し方:
   1. lnsl_push_builtin() で呼び出す関数を push する
   2. lnsl_push_xxx() で引数を push する
      (文字列は lnsl_buf() で得た領域に書き込んでから lnsl_push_str() する)
   3. lnsl_call() で呼び出す。戻り値の数が返る (エラー時は -1 で、メッセージが 1 つ残る)
   4. lnsl_kind() / lnsl_to_xxx() で戻り値を取得する
   5. lnsl_settop( 0 ) で戻り値を捨てる
*/

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#include <string.h>
#include <unistd.h>
#include "lua.h"
#include "lauxlib.h"
#include "lualib.h"

#include "lnsrt.h"

/*
  WASI に無い C の関数。Lua の io.tmpfile・os.execute・os.tmpname は失敗を返す。
*/
FILE * tmpfile( void )
{
    return NULL;
}
int system( const char * command )
{
    return command == NULL ? 0 : -1;
}
char * tmpnam( char * buf )
{
    return NULL;
}

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
    { "string", "byte" },
    { "string", "rep" },
    { "string", "lower" },
    { "string", "upper" },
    { "string", "reverse" },
    { "string", "len" },
    { "string", "char" },
    /*
      gmatch の全ての一致を返す (apply ... of str.gmatch() 用)。
      戻り値: 1 回の一致の値の数、全ての一致の値 (一致の順)
    */
    { "=", "return function( s, p )\n"
           "  local it = string.gmatch( s, p )\n"
           "  local result, n, num = {}, 0, nil\n"
           "  local function add( ... )\n"
           "    local cnt = select( '#', ... )\n"
           "    if cnt == 0 or (...) == nil then return false end\n"
           "    num = cnt\n"
           "    for i = 1, cnt do n = n + 1; result[ n ] = (select( i, ... )) end\n"
           "    return true\n"
           "  end\n"
           "  while add( it() ) do end\n"
           "  return num or 1, table.unpack( result, 1, n )\n"
           "end\n" },
    /* str.replace (runtime.lua の _lune.replace と同じ) */
    { "=", "return function( txt, src, dst )\n"
           "  local result = ''\n"
           "  local index = 1\n"
           "  while index <= #txt do\n"
           "    local findIndex = string.find( txt, src, index, true )\n"
           "    if not findIndex then\n"
           "      result = result .. string.sub( txt, index )\n"
           "      break\n"
           "    end\n"
           "    if findIndex ~= index then\n"
           "      result = result .. (string.sub( txt, index, findIndex - 1 ) .. dst)\n"
           "    else\n"
           "      result = result .. dst\n"
           "    end\n"
           "    index = findIndex + #src\n"
           "  end\n"
           "  return result\n"
           "end\n" },
    { "string", "sub" },
    { "math", "random" },
    { "math", "randomseed" },
    /*
      _load (runtime の _lune.loadstring52 と同じ) と loadfile。
      chunkname は loadfile のファイル名 ("@" + パス)
    */
    { "=", "return function( txt, env, chunkname )\n"
           "  if not env then return load( txt, chunkname ) end\n"
           "  return load( txt, chunkname or '', 'bt', env )\n"
           "end\n" },
    { NULL, "require" },
    { "=", "return function() return package.path end\n" },
    { "package", "searchpath" },
    { "string", "dump" },
    { "=", "return function() return _G end\n" },
    /*
      DependLuaOnLns.runLuaOnLns (マクロの実行)。
      Go 版 (runtime_go/depend.go の dependLuaOnLns_runLuaOnLns) と同じく、
      Lua 版の LuneScript (lune/base/*.lua) を使って Lua のコードを実行する。
      Lua のコードから LuneScript のコードを読み込む場合 (_lnsLoad) は、Lua 版の front を使う。
    */
    { "=", "local function newFront()\n"
           "  local Front = require( 'lune.base.front' )\n"
           "  local Option = require( 'lune.base.Option' )\n"
           "  local option = Option.createDefaultOption( {'dummy.lns'}, nil )\n"
           "  return Front.Front._new( option, {} )\n"
           "end\n"
           "local WrapFront = {}\n"
           "function WrapFront:setupFront()\n"
           "  if self.readyFront then return end\n"
           "  __luneScript = nil\n"
           "  self.readyFront = true\n"
           "  __luneScript = newFront()\n"
           "end\n"
           "function WrapFront:loadModule( mod )\n"
           "  self:setupFront()\n"
           "  return __luneScript:loadModule( mod )\n"
           "end\n"
           "function WrapFront:loadMeta( importModuleInfo, mod, orgMod, baseDir )\n"
           "  self:setupFront()\n"
           "  return __luneScript:loadMeta( importModuleInfo, mod, orgMod, baseDir )\n"
           "end\n"
           "function WrapFront:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )\n"
           "  self:setupFront()\n"
           "  return __luneScript:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )\n"
           "end\n"
           "function WrapFront:searchModule( mod, baseDir )\n"
           "  self:setupFront()\n"
           "  return __luneScript:searchModule( mod, baseDir )\n"
           "end\n"
           "function WrapFront:error( message )\n"
           "  self:setupFront()\n"
           "  return __luneScript:error( message )\n"
           "end\n"
           "local wrapFront = setmetatable( {}, { __index = WrapFront } )\n"
           "return function( luaCode )\n"
           "  local DependLuaOnLns = require( 'lune.base.DependLuaOnLns' )\n"
           "  __luneScript = wrapFront\n"
           "  _lnsLoad = function( name, code )\n"
           "    __luneScript = nil\n"
           "    local frontInterface = require( 'lune.base.frontInterface' )\n"
           "    __luneScript = newFront()\n"
           "    local importModuleInfo = frontInterface.ImportModuleInfo.new()\n"
           "    return frontInterface.loadFromLnsTxt( importModuleInfo, name, code )\n"
           "  end\n"
           "  return DependLuaOnLns.runLuaOnLns( wrapFront, luaCode )\n"
           "end\n" },
    /*
      wasm に埋め込んだ Lua のモジュール (wasm 版の lnsc の、Lua 版の lnsc) を
      package.preload に登録する (Go 版の Lns_initPreload と同じ)。
    */
    { "=", "return function( name, code, chunkname )\n"
           "  package.preload[ name ] = function( ... )\n"
           "    local chunk = assert( load( code, chunkname ) )\n"
           "    return chunk( ... )\n"
           "  end\n"
           "end\n" },
};
#define BUILTIN_NUM (int)(sizeof( s_builtinList ) / sizeof( s_builtinList[ 0 ] ))

/* builtin の関数を入れたテーブルの registry の参照 */
static int s_builtinRef;


void lnsr_initExt( void )
{
    s_lua = luaL_newstate();
    /*
      io・os・package・debug は、Lua で動かす LuneScript (マクロの実行) が使う。
      ファイルアクセスなどは WASI で行う (ホストがルートを preopen する)。
    */
    static const luaL_Reg libList[] = {
        { "_G", luaopen_base },
        { LUA_LOADLIBNAME, luaopen_package },
        { LUA_TABLIBNAME, luaopen_table },
        { LUA_IOLIBNAME, luaopen_io },
        { LUA_OSLIBNAME, luaopen_os },
        { LUA_STRLIBNAME, luaopen_string },
        { LUA_MATHLIBNAME, luaopen_math },
        { LUA_UTF8LIBNAME, luaopen_utf8 },
        { LUA_DBLIBNAME, luaopen_debug },
        { NULL, NULL },
    };
    for ( const luaL_Reg * lib = libList; lib->func != NULL; lib++ ) {
        luaL_requiref( s_lua, lib->name, lib->func, 1 );
        lua_pop( s_lua, 1 );
    }

    lua_createtable( s_lua, BUILTIN_NUM, 0 );
    for ( int index = 0; index < BUILTIN_NUM; index++ ) {
        if ( s_builtinList[ index ][ 0 ] != NULL &&
             strcmp( s_builtinList[ index ][ 0 ], "=" ) == 0 )
        {
            /* Lua で書いた関数 */
            if ( luaL_dostring( s_lua, s_builtinList[ index ][ 1 ] ) != LUA_OK ) {
                abort();
            }
        }
        else if ( s_builtinList[ index ][ 0 ] != NULL ) {
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

/* 値の Lua の型 (lua_type() の値) */
int32_t EXPORT(lnsl_type)( int32_t index )
{
    return lua_type( s_lua, index );
}

/*
  Lua の値 (テーブル・関数など) を、LuneScript から参照するために registry に登録する。
  戻り値は参照番号。
*/
int32_t EXPORT(lnsl_ref)( int32_t index )
{
    lua_pushvalue( s_lua, index );
    return luaL_ref( s_lua, LUA_REGISTRYINDEX );
}

/*
  lnsl_ref() の登録を解除する。
  LuneScript 側の値 ($LuaVal) が GC で回収された時に、ホストから呼ばれる。
*/
void EXPORT(lnsl_unref)( int32_t ref )
{
    luaL_unref( s_lua, LUA_REGISTRYINDEX, ref );
}

/* lnsl_ref() で登録した値を push する */
void EXPORT(lnsl_push_ref)( int32_t ref )
{
    lua_rawgeti( s_lua, LUA_REGISTRYINDEX, ref );
}

void EXPORT(lnsl_new_table)( void )
{
    lua_newtable( s_lua );
}

/* index のテーブルに、スタックの先頭の 2 つ (キー、値) を設定して取り除く */
void EXPORT(lnsl_set_table)( int32_t index )
{
    lua_settable( s_lua, index );
}

/* index のテーブルの、スタックの先頭のキーの値に置き換える */
void EXPORT(lnsl_get_table)( int32_t index )
{
    lua_gettable( s_lua, index );
}

/* index の値の長さ (# 演算子) */
int64_t EXPORT(lnsl_len)( int32_t index )
{
    return luaL_len( s_lua, index );
}

/*
  index のテーブルの、スタックの先頭のキーの次の要素を push する (lua_next())。
  要素が無い場合は 0 を返す。
*/
int32_t EXPORT(lnsl_next)( int32_t index )
{
    return lua_next( s_lua, index );
}
