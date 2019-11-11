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

#include <lns_base.h>
#include <lauxlib.h>


int main (int argc, char * pArgv[] )  {

    lua_State * pLua = luaL_newstate();
    if ( pLua == NULL ) {
        printf( "failed to create a Lua VM.\n" );
        return 1;
    }
    lua_pushcfunction( pLua, lua_main );
    lua_pushinteger( pLua, argc );
    lua_pushlightuserdata( pLua, pArgv );
    int status = lua_pcall( pLua, 2, 1, 0 );
    int result = lua_toboolean( pLua, -1 );
    if ( !status ) {
        lua_close( pLua );
        if ( result ) {
            return 0;
        }
    }
    else {
        printf( "lua error: %s\n", lua_tostring( pLua, -1 ) );
        lua_close( pLua );
    }
    return 1;
}

