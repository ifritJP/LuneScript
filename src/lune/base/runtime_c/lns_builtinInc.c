#include <stdio.h>
#include <unistd.h>

static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv )
{
    lns_any_t * lns_stdin = lns_class_lns_luaStream_new( _pEnv );
    lns_obj_lns_luaStream( lns_stdin )->pExt = stdin;
    lns_setQ_any( &lns_io_stdin, &lns_if_lns_luaStream( lns_stdin )->lns_iStream );

    lns_any_t * lns_stderr = lns_class_lns_luaStream_new( _pEnv );
    lns_obj_lns_luaStream( lns_stderr )->pExt = stderr;
    lns_setQ_any( &lns_io_stderr, &lns_if_lns_luaStream( lns_stderr )->lns_oStream );

    lns_any_t * lns_stdout = lns_class_lns_luaStream_new( _pEnv );
    lns_obj_lns_luaStream( lns_stdout )->pExt = stdout;
    lns_setQ_any( &lns_io_stdout, &lns_if_lns_luaStream( lns_stdout )->lns_oStream );
}

static void mtd_lns_luaStream__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_io__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_package__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_os__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_string__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_math__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_debug__delExt( lns_env_t * _pEnv, lns_any_t * pObj )
{
}

static void mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj)
{
    void * pExt = lns_obj_lns_luaStream( pObj )->pExt;
    if ( pExt != NULL ) {
        fclose( pExt );
    }
    else {
        lns_abort( __func__ );
    }
}

static void mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj)
{
    if ( pObj == lns_io_stdin ) {
        lns_abort( __func__ );
    }
    void * pExt = lns_obj_lns_luaStream( pObj )->pExt;
    if ( pExt != NULL ) {
        fflush( pExt );
    }
    else {
        lns_abort( __func__ );
    }
}

static lns_stem_t mtd_lns_luaStream_read(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1)
{
    void * pExt = lns_obj_lns_luaStream( pObj )->pExt;
    if ( pExt == NULL ) {
        lns_abort( __func__ );
    }

    if ( arg1.type != lns_stem_type_int ) {
        lns_abort( __func__ );
    }

    int size = arg1.val.intVal;
    char * pBuf = malloc( size );
    if ( pBuf == NULL ) {
        lns_abort( __func__ );
    }
    int readSize = fread( pBuf, 1, size, pExt );

    lns_stem_t result = lns_global.nilStem;
    if ( readSize != 0 ) {
        result = LNS_STEM_ANY( lns_cloneBin2any( _pEnv, pBuf, readSize ) );
    }
    free( pBuf );

    return result;
}


static lns_stem_t mtd_lns_luaStream_seek(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * kind, lns_int_t pos)
{
    void * pExt = lns_obj_lns_luaStream( pObj )->pExt;
    if ( pExt == NULL ) {
        lns_abort( __func__ );
    }

    int whence;
    if ( strcmp( kind->val.str.pStr, "set" ) == 0 ) {
        whence = SEEK_SET;
    }
    else if ( strcmp( kind->val.str.pStr, "cur" ) == 0 ) {
        whence = SEEK_CUR;
    }
    else if ( strcmp( kind->val.str.pStr, "end" ) == 0 ) {
        whence = SEEK_END;
    }
    else {
        lns_abort( "illegal kind" );
    }
    
    int result = fseek( pExt, pos, whence );
    if ( result >= 0 ) {
        return lns_createMRet( _pEnv, false, 1, LNS_STEM_INT( result ) );
    }

    return lns_createMRet( _pEnv, false, 1, lns_global.nilStem );
}

static lns_stem_t mtd_lns_luaStream_write(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1)
{
    void * pExt = lns_obj_lns_luaStream( pObj )->pExt;
    if ( pExt == NULL ) {
        lns_abort( __func__ );
    }

    if ( arg1->type != lns_value_type_str ) {
        lns_abort( __func__ );
    }

    lns_str_t * pStr = &arg1->val.str;
    fwrite( pStr->pStr, 1, pStr->len, pExt );

    return LNS_STEM_ANY( pObj );
}


lns_stem_t mtd_lns_io_open( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2)
{
    const char * pMode = "r";
    if ( arg2.type == lns_stem_type_any &&
         arg2.val.pAny->type == lns_value_type_str )
    {
        pMode = arg2.val.pAny->val.str.pStr;
    }
    
    FILE * pHandle = fopen( arg1->val.str.pStr, pMode );
    if ( pHandle == NULL ) {
        return lns_global.nilStem;
    }

    lns_any_t * pObj = lns_class_lns_luaStream_new( _pEnv );
    lns_obj_lns_luaStream( pObj )->pExt = pHandle;

    return LNS_STEM_ANY( pObj );
}
