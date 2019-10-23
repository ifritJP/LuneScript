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

#ifndef __LNS_LUAWRAPPER__
#define __LNS_LUAWRAPPER__

extern void lune_setLuaWapper( lune_env_t * _pEnv );
extern lune_stem_t lune__load( lune_env_t * _pEnv, lune_stem_t code, lune_stem_t newEnv );

extern lune_any_t * lune_lua_itMap_new( lune_env_t * _pEnv, lune_any_t * _obj );
extern bool lune_lua_itMap_hasNext(
    lune_env_t * _pEnv, lune_any_t * _itAny );
extern void lune_lua_itMap_getEntry(
    lune_env_t * _pEnv, lune_any_t * _itAny, lune_Map_entry_t * pEntry );


#endif
