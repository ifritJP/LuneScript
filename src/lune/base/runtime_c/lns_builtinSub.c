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

static int lns_string_call_setup_native(
    lua_State * pLua, const char * pFuncName, const char * pStr ) {
    int top = lua_gettop( pLua );

    lua_getglobal( pLua, "string" );
    lua_getfield( pLua, -1, pFuncName );
    lua_pushstring( pLua, pStr );
    
    return top;
}



static lns_any_t * lns_string_formatSub( lns_env_t * _pEnv, lns_stem_t ddd, int stackTop )
{
    lua_State * pLua = _pEnv->pLua;
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

lns_any_t * lns_string_format(
    lns_env_t * _pEnv, const char * pFormat, lns_stem_t ddd)
{
    int stackTop = lns_string_call_setup_native( _pEnv->pLua, "format", pFormat );
    
    return lns_string_formatSub( _pEnv, ddd, stackTop );
}

lns_any_t * mtd_lns_string_format(
    lns_env_t * _pEnv, lns_any_t * pFormat, lns_stem_t ddd)
{
    int stackTop = lns_string_call_setup( _pEnv->pLua, "format", pFormat );
    
    return lns_string_formatSub( _pEnv, ddd, stackTop );
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
    int stackTop = lns_string_call_setup( pLua, "rep", pTxt );

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

lns_stem_t mtd_lns_string_byte(
    lns_env_t * _pEnv, lns_any_t * str, lns_stem_t index1, lns_stem_t index2)
{
    int startIndex = 1;
    if ( index1.type == lns_stem_type_int ) {
        startIndex = index1.val.intVal;
    }
    int endIndex = startIndex;
    if ( index2.type == lns_stem_type_int ) {
        endIndex = index2.val.intVal;
        if ( endIndex > str->val.str.len ) {
            endIndex = str->val.str.len;
        }
    }
    int len = endIndex - startIndex + 1;
    if ( len <= 0 || startIndex > str->val.str.len ) {
        return lns_global.ddd0;
    }

    lns_any_t * pRetObj = lns_createMRetOnly( _pEnv, len );

    int index;
    for ( index = startIndex; index <= endIndex; index++ ) {
        lns_set2DDDArg(
            pRetObj, index - startIndex,
            LNS_STEM_INT( (unsigned char)str->val.str.pStr[ index - 1 ] ) );
    }

    return LNS_STEM_ANY( pRetObj );
}

static int fix_string_index( int index, int len, bool correctMinus ) {
    int workIndex;
    if ( index < 0 ) {
        workIndex = len + index;
    }
    else if ( index > len ) {
        workIndex = len;
    }
    else {
        workIndex = index - 1;
    }
    if ( correctMinus && workIndex < 0 ) {
        return 0;
    }
    return workIndex;
}

lns_any_t * mtd_lns_string_sub(
    lns_env_t * _pEnv, lns_any_t * arg1, lns_int_t arg2, lns_stem_t arg3)
{
    const lns_str_t * pStr = &arg1->val.str;
    int lastIndex;
    if ( LNS_IS_NILNONE( arg3 ) ) {
        lastIndex = pStr->len - 1;
    }
    else {
        lastIndex = fix_string_index( arg3.val.intVal, pStr->len, false );
    }
    int startIndex = fix_string_index( arg2, pStr->len, true );
    if ( lastIndex < startIndex ) {
        return lns_litStr2any( _pEnv, "" );
    }
    int len = lastIndex - startIndex + 1;
    if ( len < 0 ) {
        len = 0;
    }
    
    lns_any_t * pResult = lns_cloneBin2any( _pEnv, pStr->pStr + startIndex, len );
    return pResult;
}
