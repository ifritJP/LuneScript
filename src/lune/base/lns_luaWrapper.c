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
#include <lauxlib.h>
#include <lunescript.h>

static const char * lune_valTableName = "__lune_vals";

void lune_setLuaWapper( lune_env_t * _pEnv )
{
    lua_newtable( _pEnv->pLua );
    lua_setglobal( _pEnv->pLua, lune_valTableName );
}

lune_stem_t lune__load( lune_env_t * _pEnv, lune_stem_t code, lune_stem_t newEnv )
{
    
    int result = luaL_loadstring (_pEnv->pLua, code.val.pAny->val.str.pStr );
    if ( result == LUA_OK ) {
        char name[ 30 ];
        lune_any_t * pAny = lune_luaVal_new( _pEnv );
        pAny->val.luaVal.type = lune_value_type_form;
        
        sprintf( name, "load%p", pAny );
        // __lune_vals.load%x = load();
        lua_getglobal( _pEnv->pLua, lune_valTableName );
        lua_pushvalue( _pEnv->pLua, -2 );
        lua_setfield( _pEnv->pLua, -2, name );
        lua_pop( _pEnv->pLua, 1 );
        
        if ( newEnv.type != lune_stem_type_nil ) {
            lune_abort( "not suport" );
        }

        return lune_createMRet(
            _pEnv, false, 1,  LUNE_STEM_ANY( pAny ) );
    }
    return lune_createMRet( _pEnv, false, 2, lune_global.nilStem,
                            lune_litStr2stem( _pEnv, "failed to load" ) );
}
