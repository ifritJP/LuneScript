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

#define LNS_LUA_NAME_LEN 30

static const char * lns_valTableName = "__lns_vals";
static const char * lns_swapWorkName = "__lns_swapwork";

void lns_setLuaWapper( lns_env_t * _pEnv )
{
    luaL_openlibs( _pEnv->pLua );
    
    lua_newtable( _pEnv->pLua );
    lua_setglobal( _pEnv->pLua, lns_valTableName );
}

static void lns_getAccessName( const void * pKey, char * pBuf )
{
    sprintf( pBuf, "%p", pKey );
}

void lns_releaeAnyVal( lns_env_t * _pEnv, void * pKey )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );
    
    // __lns_swapwork = __lns_vals[ name ]
    char name[ LNS_LUA_NAME_LEN ];
    lns_getAccessName( pKey, name );
    lua_getglobal( pLua, lns_valTableName );
    lua_pushnil( pLua );
    lua_setfield( pLua, -2, name );

    lua_settop( pLua, top );
}

/**
pAny が保持する Lua モジュールを push する。
 */
void lns_pushAnyVal( lns_env_t * _pEnv, void * pKey )
{
    // __lns_vals[ name ] に保持している値を取り出す。
    // ただし、このままだと __lns_vals 自体がスタックに残るので、
    // 一旦 __lns_swapwork に格納して、スタックをクリアしてから
    // __lns_swapwork をスタックに戻すことで、スタックに値だけつむように処理している。

    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );
    
    // __lns_swapwork = __lns_vals[ name ]
    char name[ LNS_LUA_NAME_LEN ];
    lns_getAccessName( pKey, name );
    lua_getglobal( _pEnv->pLua, lns_valTableName );
    lua_getfield( _pEnv->pLua, -1, name );

    // push __lns_swapwork
    lua_setglobal( _pEnv->pLua, lns_swapWorkName );
    lua_pop( _pEnv->pLua, 1 );
    lua_getglobal( _pEnv->pLua, lns_swapWorkName );

    int after = lua_gettop( pLua );
    
}

/**
スタック top の値を __lns_vals[ name ] に保持する。
スタック top の値は取り除かれる。
 */
static void lns_setAnyVal( lns_env_t * _pEnv, void * pKey )
{
    // __lns_vals[ name ] = stack[ index ]
    char name[ LNS_LUA_NAME_LEN ];
    lns_getAccessName( pKey, name );
    lua_getglobal( _pEnv->pLua, lns_valTableName );
    lua_pushvalue( _pEnv->pLua, -2 );
    lua_setfield( _pEnv->pLua, -2, name );
    lua_pop( _pEnv->pLua, 2 );
}

/**
スタックの index の位置にある文字列を lns_stem_t にセットする。
 */
void lns_lua_stack2str( lns_env_t * _pEnv, int index, lns_stem_t * pStem )
{
    size_t len;
    const char * pMess = lua_tolstring( _pEnv->pLua, index, &len );
    pStem->type = lns_stem_type_any;
    pStem->val.pAny = lns_cloneBin2any( _pEnv, pMess, len );
}

/**
スタックの top の値から pStem を設定する。
 */
void lns_setupFromStack( lns_env_t * _pEnv, int index, lns_stem_t * pStem )
{
    lua_State * pLua = _pEnv->pLua;
    switch ( lua_type( pLua, index ) ) {
    case LUA_TNIL:
        *pStem = lns_global.nilStem;
        //lua_pop( pLua, 1 );
        break;
    case LUA_TNUMBER:
        {
            double work;
            lns_real_t realVal = lua_tonumber( pLua, index );
            if ( modf( realVal, &work ) == 0 ) {
                *pStem = LNS_STEM_INT( lua_tointeger( pLua, index ) );
            }
            else {
                *pStem = LNS_STEM_REAL( realVal );
            }
            //lua_pop( pLua, 1 );
        }
        break;
    case LUA_TBOOLEAN:
        *pStem = LNS_STEM_BOOL( lua_toboolean( pLua, index ) );
        //lua_pop( pLua, 1 );
        break;
    case LUA_TSTRING:
        lns_lua_stack2str( _pEnv, index, pStem );
        // lua_pop( pLua, 1 );
        break;
    case LUA_TFUNCTION:
        {
            *pStem = LNS_STEM_ANY( lns_createFormAny( _pEnv ) );
            lua_pushvalue( pLua, index );
            lns_setAnyVal( _pEnv, pStem->val.pAny );
        }
        break;
    default:
        {
            *pStem = LNS_STEM_ANY( lns_luaVal_new( _pEnv, lns_value_type_luaVal ) );
            lua_pushvalue( pLua, index );
            lns_setAnyVal( _pEnv, pStem->val.pAny );
        }
        break;
    }
}


static lns_stem_t lns_lua_loadedFunc( lns_env_t * _pEnv, lns_any_t * _pForm )
{
    lua_State * pLua = _pEnv->pLua;
    lns_stem_t retObj;
    
    int top = lua_gettop( pLua );
    
    lns_pushAnyVal( _pEnv, _pForm );
    int luaRet = lua_pcall( pLua, 0, LUA_MULTRET, 0 );
    if ( luaRet == LUA_OK ) {
        if ( lua_isnil( pLua, top ) ) {
            // ロード失敗
            lns_stem_t mess;
            lns_lua_stack2str( _pEnv, top + 2, &mess );
            retObj = lns_createMRet( _pEnv, false, 2, lns_global.nilStem, mess );
        }
        else {
            // ロード成功
            lns_stem_t mapObj;
            lns_setupFromStack( _pEnv, top + 1, &mapObj );
            retObj = lns_createMRet( _pEnv, false, 2, mapObj, lns_global.nilStem );
        }
    }
    else {
        lns_stem_t mess;
        lns_lua_stack2str( _pEnv, top + 2, &mess );
        retObj = lns_createMRet( _pEnv, false, 2, lns_global.nilStem, mess );
    }

    lua_settop( pLua, top );

    return retObj;
}

lns_stem_t lns_f__load( lns_env_t * _pEnv, lns_any_t * code, lns_stem_t newEnv )
{
    int top = lua_gettop( _pEnv->pLua );

    int result = luaL_loadstring (_pEnv->pLua, code->val.str.pStr );

    lns_stem_t retObj;
    
    if ( result == LUA_OK ) {
        lns_any_t * formObj = lns_func2any(
            _pEnv, (lns_closure_t * )lns_lua_loadedFunc, 0, false, 0 );
        lns_setClosure( formObj );

        lns_setAnyVal( _pEnv, formObj );

        if ( newEnv.type != lns_stem_type_nil && newEnv.type != lns_stem_type_none )
        {
            lns_abort( "not suport" );
        }

        retObj = lns_createMRet( _pEnv, false, 1,  LNS_STEM_ANY( formObj ) );
    }
    else { 
        retObj = lns_createMRet(
            _pEnv, false, 2, lns_global.nilStem,
            LNS_STEM_ANY( lns_litStr2any( _pEnv, "failed to load" ) ) );
    }

    lua_settop( _pEnv->pLua, top );

    return retObj;
}

/**

 */
lns_any_t * lns_lua_itMap_new( lns_env_t * _pEnv, lns_any_t * _obj )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, nil }
    lua_newtable( pLua );
    // map のオブジェクトをプッシュする
    lns_pushAnyVal( _pEnv, _obj );

    lua_seti( pLua, top + 1, 1 );


    // { map, nil } を保持
    lns_any_t * itObj = lns_luaVal_new( _pEnv, lns_value_type_itMap );
    lns_setAnyVal( _pEnv, itObj );

    lua_settop( pLua, top );

    return itObj;
}

bool lns_lua_itMap_hasNext(
    lns_env_t * _pEnv, lns_any_t * _itAny )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, key, val } を push
    lns_pushAnyVal( _pEnv, _itAny );

    // map を push
    lua_geti( pLua, top + 1, 1 );
    // key を push
    lua_geti( pLua, top + 1, 2 );

    int ret = lua_next( pLua, -2 );
    if ( ret != 0 ) {
        // val を保持
        lua_seti( pLua, top + 1, 3 );
        // key を保持
        lua_seti( pLua, top + 1, 2 );
    }

    lua_settop( pLua, top );
    
    return ret != 0;
}


void lns_lua_itMap_getEntry(
    lns_env_t * _pEnv, lns_any_t * _itAny, lns_Map_entry_t * pEntry )
{
    lua_State * pLua = _pEnv->pLua;
    int top = lua_gettop( pLua );

    // { map, key } を push
    lns_pushAnyVal( _pEnv, _itAny );
    // key を push
    lua_geti( pLua, top + 1, 2 );
    // key を登録
    lns_setupFromStack( _pEnv, -1, &pEntry->key );
    // val を push
    lua_geti( pLua, top + 1, 3 );
    // val を登録
    lns_setupFromStack( _pEnv, -1, &pEntry->val );

    lua_settop( pLua, top );
}

/**
 * Lua の関数コール処理。
 *
 * @param stackBase 関数コール前の基準スタック位置。
 *    この関数処理後は、stackBase 位置をスタックトップとする。
 * @param argNum Lua 関数に渡す引数の数
 * @param retNum Lua 関数の戻り値の数。 不定の場合 LUA_MULTRET。
 * @param Lua 関数の戻り値。
 */
lns_stem_t lns_lua_call( lns_env_t * _pEnv, int stackBase, int argNum, int retNum )
{
    lua_State * pLua = _pEnv->pLua;

    int top = lua_gettop( pLua ) - argNum - 1;
    lua_call( pLua, argNum, LUA_MULTRET );

    int lastRet = lua_gettop( pLua );

    lns_stem_t result;

    int stackNum = lastRet - top;
    if ( retNum >= 2 || retNum == LUA_MULTRET ) {
        lns_any_t * pMRet = lns_createMRetOnly( _pEnv, stackNum );

        int index;
        lns_stem_t * pStem = &pMRet->val.ddd.stemList[ 0 ];
        for ( index = top + 1; index <= lastRet; index++, pStem++ ) {
            lns_setupFromStack( _pEnv, index, pStem );
        }
        result = LNS_STEM_ANY( pMRet );
        lns_setRetInBlock( _pEnv, result );
    }
    else if ( stackNum == 1 ) {
        lns_setupFromStack( _pEnv, -1, &result );
        lns_setRetInBlock( _pEnv, result );
    }
    else {
        result = lns_global.noneStem;
    }

    lua_settop( pLua, stackBase );

    return result;
}

/**
 * Lua の関数をコールする
 *
 */
lns_stem_t lns_lua_callForm(
    lns_env_t * _pEnv, lns_any_t * _pForm, lns_stem_t ddd )
{
    lua_State * pLua = _pEnv->pLua;
    int index;

    int stackTop = lua_gettop( pLua );

    lns_pushAnyVal( _pEnv, _pForm );
    
    lns_any_t * pAny = ddd.val.pAny;
    for ( index = 0; index < lns_lenDDD( pAny ); index++ ) {
        lns_stem_t stem = lns_fromDDD( pAny, index );
        switch ( stem.type ) {
        case lns_stem_type_nil:
            lua_pushnil( pLua );
            break;
        case lns_stem_type_int:
            lua_pushinteger( pLua, stem.val.intVal );
            break;
        case lns_stem_type_real:
            lua_pushinteger( pLua, stem.val.realVal );
            break;
        case lns_stem_type_bool:
            lua_pushboolean( pLua, stem.val.boolVal );
            break;
        case lns_stem_type_any:
            {
                lns_any_t * pAny = stem.val.pAny;
                switch ( pAny->type ) {
                case lns_value_type_str:
                    lua_pushlstring( pLua, pAny->val.str.pStr, pAny->val.str.len );
                    break;
                default:
                    lns_pushAnyVal( _pEnv, pAny );
                    break;
                }
            }
            break;
        default:
            lns_abort( "illegal type" );
            break;
        }
    }

    return lns_lua_call( _pEnv, stackTop, lns_lenDDD( pAny ), LUA_MULTRET );
}

lns_any_t * lns_createFormAny( lns_env_t * _pEnv )
{
    lns_any_t * pFormAny = lns_luaForm_new( _pEnv );
    lns_form_t * pForm = &pFormAny->val.form;
    pForm->len = 0;
    pForm->needFormParam = true;
    pForm->pClosureValList = NULL;
    pForm->pFunc = lns_lua_callForm;
    pForm->hasDDD = false;
    return pFormAny;
}
    
