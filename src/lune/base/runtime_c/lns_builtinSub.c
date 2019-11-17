#include <lunescript.h>
#include <lauxlib.h>
#include <lns_luaWrapper.h>
#include <stdlib.h>

#define LNS_PUSH_STR( LUA, STR )                               \
    lua_pushlstring( LUA, STR->val.str.pStr, STR->val.str.len )

lns_any_t * lns_f_error( lns_env_t * _pEnv, lns_any_t * arg1 )
{
    lua_State * pLua = _pEnv->pLua;
    LNS_PUSH_STR( pLua, arg1 );
    lua_error( pLua );
    return NULL;
}

lns_any_t * mtd_lns_os_exit( lns_env_t * _pEnv, lns_stem_t arg1)
{
    if ( arg1.type == lns_stem_type_int ) {
        exit( arg1.val.intVal );
    }
    else {
        exit( 0 );
    }
}


static int lns_string_call_setup(
    lua_State * pLua, const char * pFuncName, lns_any_t * pStrAny ) {
    int top = lua_gettop( pLua );

    lua_getglobal( pLua, "string" );
    lua_getfield( pLua, -1, pFuncName );
    LNS_PUSH_STR( pLua, pStrAny );
    
    return top;
}


lns_any_t * mtd_lns_string_format(
    lns_env_t * _pEnv, lns_any_t * pFormat, lns_stem_t ddd)
{
    lua_State * pLua = _pEnv->pLua;
    int stackTop = lns_string_call_setup( pLua, "format", pFormat );

    int index;
    // ... の値を push する
    for ( index = 0; index < lns_lenDDD( ddd.val.pAny ); index++ ) {
        lns_stem_t stem = lns_fromDDD( ddd.val.pAny, index );
        switch ( stem.type ) {
        case lns_stem_type_none: // fall-through
        case lns_stem_type_nil:
            lua_pushnil( pLua );
            break;
        case lns_stem_type_int:
            lua_pushinteger( pLua, stem.val.intVal );
            break;
        case lns_stem_type_real:
            lua_pushnumber( pLua, stem.val.realVal );
            break;
        case lns_stem_type_bool:
            lua_pushboolean( pLua, stem.val.boolVal );
            break;
        case lns_stem_type_any:
            {
                lns_any_t * pAny = stem.val.pAny;
                switch ( pAny->type ) {
                case lns_value_type_str:
                    LNS_PUSH_STR( pLua, pAny );
                    break;
                case lns_value_type_form:
                    lua_pushfstring( pLua, "form: %p", pAny );
                    break;
                case lns_value_type_class:
                    lua_pushfstring( pLua, "class: %p(%s)",
                                     pAny, pAny->val.classVal->pMeta->pName );
                    break;
                case lns_value_type_if:
                    lua_pushfstring(
                        pLua, "if: %p(%s): %p(%s)",
                        pAny, pAny->val.ifVal.pMeta->pName,
                        pAny->val.ifVal.pObj,
                        pAny->val.ifVal.pObj->val.classVal->pMeta->pName );
                    break;
                case lns_value_type_luaVal:
                    lns_pushAnyVal( _pEnv, pAny );
                default:
                    printf( "unknown type -- %d", pAny->type );
                    break;
                }
            }
            break;
        default:
            lua_pushstring( pLua, "<none>" );
        }
    }

    // 関数を実行
    return lns_lua_call( _pEnv, stackTop, 1 + lns_lenDDD( ddd.val.pAny ), 1 ).val.pAny;
}


extern lns_stem_t mtd_lns_string_find(
    lns_env_t * _pEnv, lns_any_t * pTarget,
    lns_any_t * pPattern, lns_stem_t init, lns_stem_t plain)
{
    lua_State * pLua = _pEnv->pLua;

    // string.find と pTarget を push
    int stackTop = lns_string_call_setup( pLua, "find", pTarget );

    // 残りの引数を push
    LNS_PUSH_STR( pLua, pPattern );
    int argNum = 2;
    if ( init.type != lns_stem_type_none ) {
        if ( init.type == lns_stem_type_nil ) {
            lua_pushnil( pLua );
        }
        else {
            lua_pushinteger( pLua, init.val.intVal );
        }
        argNum++;
    }
    if ( plain.type != lns_stem_type_none ) {
        if ( init.type == lns_stem_type_nil ) {
            lua_pushnil( pLua );
        }
        else {
            lua_pushboolean( pLua, plain.val.boolVal );
        }
        argNum++;
    }
    
    // Lua 関数を実行
    return lns_lua_call( _pEnv, stackTop, argNum, 2 );
    
}


lns_any_t * mtd_lns_string_rep( lns_env_t * _pEnv, lns_any_t * pTxt, lns_int_t num)
{
    lua_State * pLua = _pEnv->pLua;
    int stackTop = lns_string_call_setup( pLua, "find", pTxt );

    lua_pushinteger( pLua, num );

    return lns_lua_call( _pEnv, stackTop, 2, 1 ).val.pAny;
}

lns_stem_t mtd_lns_string_gmatch(
    lns_env_t * _pEnv, lns_any_t * pTarget, lns_any_t * pPattern)
{
    lua_State * pLua = _pEnv->pLua;
    int stackTop = lns_string_call_setup( pLua, "gmatch", pTarget );

    LNS_PUSH_STR( pLua, pPattern );

    // Lua 関数を実行
    return lns_lua_call( _pEnv, stackTop, 2, 3 );
}
