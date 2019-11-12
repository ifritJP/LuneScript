#include <lunescript.h>
#include <lauxlib.h>
#include <lns_luaWrapper.h>

lune_any_t * u_mtd_string_format(
    lune_env_t * _pEnv, lune_any_t * pFormat, lune_stem_t ddd)
{
    lua_State * pLua = _pEnv->pLua;

    int stackTop = lua_gettop( pLua );
    
    lua_getglobal( pLua, "string" );
    lua_getfield( _pEnv->pLua, -1, "format" );
    lua_pushlstring( pLua, pFormat->val.str.pStr, pFormat->val.str.len );
    int index;

    // ... の値を push する
    for ( index = 0; index < lune_lenDDD( ddd.val.pAny ); index++ ) {
        lune_stem_t stem = lune_fromDDD( ddd.val.pAny, index );
        switch ( stem.type ) {
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
                    lua_pushlstring( pLua, pAny->val.str.pStr, pAny->val.str.len );
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
    lua_call( pLua, 1 + lune_lenDDD( ddd.val.pAny ), 1 );

    lune_stem_t result;
    lune_lua_stack2str( _pEnv, -1, &result );
    lune_setRet( _pEnv, result );

    lua_settop( pLua, stackTop );

    return result.val.pAny;
}
