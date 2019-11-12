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

#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>
#include <lunescript.h>
#include <math.h>

#define LUNE_LUA_NAME_LEN 30

static const char * lune_valTableName = "__lune_vals";
static const char * lune_swapWorkName = "__lune_swapwork";

void lune_setLuaWapper( lune_env_t * _pEnv )
{
    luaL_openlibs( _pEnv->pLua );
    
    lua_newtable( _pEnv->pLua );
    lua_setglobal( _pEnv->pLua, lune_valTableName );
}

static void lune_getAccessName( const void * pKey, char * pBuf )
{
    sprintf( pBuf, "%p", pKey );
}

/**
pAny が保持する Lua モジュールを push する。
 */
void lune_pushAnyVal( lune_env_t * _pEnv, void * pKey )
{
    // __lune_vals[ name ] に保持している値を取り出す。
    // ただし、このままだと __lune_vals 自体がスタックに残るので、
    // 一旦 __lune_swapwork に格納して、スタックをクリアしてから
    // __lune_swapwork をスタックに戻すことで、スタックに値だけつむように処理している。

    
    // __lune_swapwork = __lune_vals[ name ]
    char name[ LUNE_LUA_NAME_LEN ];
    lune_getAccessName( pKey, name );
    lua_getglobal( _pEnv->pLua, lune_valTableName );
    lua_getfield( _pEnv->pLua, -1, name );

    // push __lune_swapwork
    lua_setglobal( _pEnv->pLua, lune_swapWorkName );
    lua_pop( _pEnv->pLua, 2 );
    lua_getglobal( _pEnv->pLua, lune_swapWorkName );
}

/**
スタック top の値を __lune_vals[ name ] に保持する。
スタック top の値は取り除かれる。
 */
static void lune_setAnyVal( lune_env_t * _pEnv, void * pKey )
{
    // __lune_vals[ name ] = stack[ index ]
    char name[ LUNE_LUA_NAME_LEN ];
    lune_getAccessName( pKey, name );
    lua_getglobal( _pEnv->pLua, lune_valTableName );
    lua_pushvalue( _pEnv->pLua, -2 );
    lua_setfield( _pEnv->pLua, -2, name );
    lua_pop( _pEnv->pLua, 2 );
}

/**
スタックの index の位置にある文字列を lune_stem_t にセットする。
 */
void lune_lua_stack2str( lune_env_t * _pEnv, int index, lune_stem_t * pStem )
{
    size_t len;
    const char * pMess = lua_tolstring( _pEnv->pLua, index, &len );
    pStem->type = lune_stem_type_any;
    pStem->val.pAny = lune_cloneBin2any( _pEnv, pMess, len );
}

/**
スタックの top の値から pStem を設定する。
スタック top の値は取り除かれる。
 */
static void lune_setupFromStack( lune_env_t * _pEnv, int index, lune_stem_t * pStem )
{
    lua_State * pLua = _pEnv->pLua;
    switch ( lua_type( pLua, index ) ) {
    case LUA_TNIL:
        *pStem = lune_global.nilStem;
        lua_pop( pLua, 1 );
        break;
    case LUA_TNUMBER:
        {
            double work;
            lune_real_t realVal = lua_tonumber( pLua, index );
            if ( modf( realVal, &work ) == 0 ) {
                *pStem = LUNE_STEM_INT( lua_tointeger( pLua, index ) );
            }
            else {
                *pStem = LUNE_STEM_REAL( realVal );
            }
            lua_pop( pLua, 1 );
        }
        break;
    case LUA_TBOOLEAN:
        *pStem = LUNE_STEM_BOOL( lua_toboolean( pLua, index ) );
        lua_pop( pLua, 1 );
        break;
    case LUA_TSTRING:
        lune_lua_stack2str( _pEnv, index, pStem );
        lua_pop( pLua, 1 );
        break;
    default:
        {
            *pStem = LUNE_STEM_ANY( lune_luaVal_new( _pEnv, lune_value_type_luaVal ) );
            lua_pushvalue( pLua, index );
            lune_setAnyVal( _pEnv, pStem->val.pAny );
        }
        break;
    }
}


static lune_stem_t lune_lua_loadedFunc( lune_env_t * _pEnv, lune_any_t * _pForm )
{
    lua_State * pLua = _pEnv->pLua;
    lune_stem_t retObj;
    
    int top = lua_gettop( pLua );
    
    lune_pushAnyVal( _pEnv, _pForm );
    int luaRet = lua_pcall( pLua, 0, LUA_MULTRET, 0 );
    if ( luaRet == LUA_OK ) {
        if ( lua_isnil( pLua, top ) ) {
            // ロード失敗
            lune_stem_t mess;
            lune_lua_stack2str( _pEnv, top + 1, &mess );
            retObj = lune_createMRet( _pEnv, false, 2, lune_global.nilStem, mess );
        }
        else {
            // ロード成功
            lune_stem_t mapObj;
            lune_setupFromStack( _pEnv, top, &mapObj );
            retObj = lune_createMRet( _pEnv, false, 2, mapObj, lune_global.nilStem );
        }
    }
    else {
        lune_stem_t mess;
        lune_lua_stack2str( _pEnv, top, &mess );
        retObj = lune_createMRet( _pEnv, false, 2, lune_global.nilStem, mess );
    }

    lua_settop( pLua, top );

    return retObj;
}

lune_stem_t lune__load( lune_env_t * _pEnv, lune_any_t * code, lune_stem_t newEnv )
{
    int top = lua_gettop( _pEnv->pLua );

    int result = luaL_loadstring (_pEnv->pLua, code->val.str.pStr );

    lune_stem_t retObj;
    
    if ( result == LUA_OK ) {
        lune_any_t * formObj = lune_func2any(
            _pEnv, (lune_closure_t * )lune_lua_loadedFunc, 0, false, 0 );
        lune_setClosure( formObj );

        lune_setAnyVal( _pEnv, formObj );

        if ( newEnv.type != lune_stem_type_nil ) {
            lune_abort( "not suport" );
        }

        retObj = lune_createMRet( _pEnv, false, 1,  LUNE_STEM_ANY( formObj ) );
    }
    else { 
        retObj = lune_createMRet(
            _pEnv, false, 2, lune_global.nilStem,
            LUNE_STEM_ANY( lune_litStr2any( _pEnv, "failed to load" ) ) );
    }

    lua_settop( _pEnv->pLua, top );

    return retObj;
}

/**

 */
lune_any_t * lune_lua_itMap_new( lune_env_t * _pEnv, lune_any_t * _obj )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, nil }
    lua_newtable( pLua );
    // map のオブジェクトをプッシュする
    lune_pushAnyVal( _pEnv, _obj );
    lua_seti( pLua, -2, 1 );


    // { map, nil } を保持
    lune_any_t * itObj = lune_luaVal_new( _pEnv, lune_value_type_itMap );
    lune_setAnyVal( _pEnv, itObj );

    lua_settop( pLua, top );

    return itObj;
}

bool lune_lua_itMap_hasNext(
    lune_env_t * _pEnv, lune_any_t * _itAny )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, key, val } を push
    lune_pushAnyVal( _pEnv, _itAny );
    // map を push
    lua_geti( pLua, -1, 1 );
    // key を push
    lua_geti( pLua, -2, 2 );

    int ret = lua_next( pLua, -2 );
    if ( ret != 0 ) {
        // val を保持
        lua_seti( pLua, top, 3 );
        // key を保持
        lua_seti( pLua, top, 2 );
    }

    lua_settop( pLua, top );
    
    return ret != 0;
}


void lune_lua_itMap_getEntry(
    lune_env_t * _pEnv, lune_any_t * _itAny, lune_Map_entry_t * pEntry )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, key } を push
    lune_pushAnyVal( _pEnv, _itAny );
    // key を push
    lua_geti( pLua, top, 2 );
    // key を取り出して登録
    lune_setupFromStack( _pEnv, -1, &pEntry->key );
    // val を push
    lua_geti( pLua, top, 3 );
    // val を取り出して登録
    lune_setupFromStack( _pEnv, -1, &pEntry->val );

    lua_settop( pLua, top );
}
