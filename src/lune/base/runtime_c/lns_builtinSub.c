#include <lunescript.h>
#include <lauxlib.h>
#include <lns_luaWrapper.h>

#define LUNE_PUSH_STR( LUA, STR )                               \
    lua_pushlstring( LUA, STR->val.str.pStr, STR->val.str.len )

static int lune_string_call_setup(
    lua_State * pLua, const char * pFuncName, lune_any_t * pStrAny ) {
    int top = lua_gettop( pLua );

    lua_getglobal( pLua, "string" );
    lua_getfield( pLua, -1, pFuncName );
    LUNE_PUSH_STR( pLua, pStrAny );
    
    return top;
}


lune_any_t * u_mtd_string_format(
    lune_env_t * _pEnv, lune_any_t * pFormat, lune_stem_t ddd)
{
    lua_State * pLua = _pEnv->pLua;
    int stackTop = lune_string_call_setup( pLua, "format", pFormat );

    int index;
    // ... の値を push する
    for ( index = 0; index < lune_lenDDD( ddd.val.pAny ); index++ ) {
        lune_stem_t stem = lune_fromDDD( ddd.val.pAny, index );
        switch ( stem.type ) {
        case lune_stem_type_none: // fall-through
        case lune_stem_type_nil:
            lua_pushnil( pLua );
            break;
        case lune_stem_type_int:
            lua_pushinteger( pLua, stem.val.intVal );
            break;
        case lune_stem_type_real:
            lua_pushnumber( pLua, stem.val.realVal );
            break;
        case lune_stem_type_bool:
            lua_pushboolean( pLua, stem.val.boolVal );
            break;
        case lune_stem_type_any:
            {
                lune_any_t * pAny = stem.val.pAny;
                switch ( pAny->type ) {
                case lune_value_type_str:
                    LUNE_PUSH_STR( pLua, pAny );
                    break;
                case lune_value_type_form:
                    lua_pushfstring( pLua, "form: %p", pAny );
                    break;
                case lune_value_type_class:
                    lua_pushfstring( pLua, "class: %p(%s)",
                                     pAny, pAny->val.classVal->pMeta->pName );
                    break;
                case lune_value_type_if:
                    lua_pushfstring(
                        pLua, "if: %p(%s): %p(%s)",
                        pAny, pAny->val.ifVal.pMeta->pName,
                        pAny->val.ifVal.pObj,
                        pAny->val.ifVal.pObj->val.classVal->pMeta->pName );
                    break;
                case lune_value_type_luaVal:
                    lune_pushAnyVal( _pEnv, pAny );
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
    lua_call( pLua, 1 + lune_lenDDD( ddd.val.pAny ), LUA_MULTRET );

    lune_stem_t result;
    lune_lua_stack2str( _pEnv, stackTop + 2, &result );
    lune_setRet( _pEnv, result );

    lua_settop( pLua, stackTop );

    return result.val.pAny;
}


extern lune_stem_t u_mtd_string_find(
    lune_env_t * _pEnv, lune_any_t * pTarget,
    lune_any_t * pPattern, lune_stem_t init, lune_stem_t plain)
{
    lua_State * pLua = _pEnv->pLua;
    int stackTop = lune_string_call_setup( pLua, "format", pTarget );

    LUNE_PUSH_STR( pLua, pPattern );
    int argNum = 2;
    if ( init.type != lune_stem_type_none ) {
        if ( init.type == lune_stem_type_nil ) {
            lua_pushnil( pLua );
        }
        else {
            lua_pushinteger( pLua, init.val.intVal );
        }
        argNum++;
    }
    if ( plain.type != lune_stem_type_none ) {
        if ( init.type == lune_stem_type_nil ) {
            lua_pushnil( pLua );
        }
        else {
            lua_pushboolean( pLua, plain.val.boolVal );
        }
        argNum++;
    }
    
    // 関数を実行
    lua_call( pLua, argNum, LUA_MULTRET );

    lune_stem_t findIndex;
    lune_setupFromStack( _pEnv, stackTop + 2, &findIndex );
    lune_stem_t sIndex;
    lune_setupFromStack( _pEnv, stackTop + 3, &sIndex );

    lune_stem_t result = lune_createMRet( _pEnv, false, 2, findIndex, sIndex );
    lune_setMRet( _pEnv, result.val.pAny );

    lua_settop( pLua, stackTop );

    return result;
}