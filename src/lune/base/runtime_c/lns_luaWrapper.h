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

extern void lns_setLuaWapper( lns_env_t * _pEnv );
extern lns_stem_t lns_f__load( lns_env_t * _pEnv, lns_any_t * code, lns_stem_t newEnv );

extern lns_any_t * lns_lua_itMap_new( lns_env_t * _pEnv, lns_any_t * _obj );
extern bool lns_lua_itMap_hasNext(
    lns_env_t * _pEnv, lns_any_t * _itAny );
extern void lns_lua_itMap_getEntry(
    lns_env_t * _pEnv, lns_any_t * _itAny, lns_Map_entry_t * pEntry );

extern void lns_pushAnyVal( lns_env_t * _pEnv, void * pKey );
extern void lns_lua_stack2str( lns_env_t * _pEnv, int index, lns_stem_t * pStem );
extern void lns_setupFromStack( lns_env_t * _pEnv, int index, lns_stem_t * pStem );

#endif
