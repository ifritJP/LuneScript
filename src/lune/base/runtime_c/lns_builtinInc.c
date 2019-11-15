#include <stdio.h>
#include <unistd.h>

static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv )
{
    lns_any_t * lns_stdin = lns_class_lns_luaStream_new( _pEnv );
    lns_setQ_any( &l_var_lns_io_stdin, &lns_if_lns_luaStream( lns_stdin )->lns_iStream );
    lns_any_t * lns_stderr = lns_class_lns_luaStream_new( _pEnv );
    lns_setQ_any( &l_var_lns_io_stderr, &lns_if_lns_luaStream( lns_stderr )->lns_oStream );
    lns_any_t * lns_stdout = lns_class_lns_luaStream_new( _pEnv );
    lns_setQ_any( &l_var_lns_io_stdout, &lns_if_lns_luaStream( lns_stdout )->lns_oStream );
}

static void mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj)
{
    if ( pObj == l_var_lns_io_stdin ) {
        fclose( stdin );
    }
    else if ( pObj == l_var_lns_io_stderr ) {
        fclose( stderr );
    }
    else if ( pObj == l_var_lns_io_stdout ) {
        fclose( stdout );
    }
    else {
        lns_abort( __func__ );
    }
}

static void mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj)
{
    if ( pObj == l_var_lns_io_stdin ) {
        lns_abort( __func__ );
    }
    else if ( pObj == l_var_lns_io_stderr ) {
        fflush( stderr );
    }
    else if ( pObj == l_var_lns_io_stdout ) {
        fflush( stdout );
    }
    else {
        lns_abort( __func__ );
    }
}

static lns_stem_t mtd_lns_luaStream_read(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1)
{
    if ( pObj != l_var_lns_io_stdin ) {
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
    int readSize = fread( pBuf, 1, size, stdin );

    lns_stem_t result = lns_global.nilStem;
    if ( readSize != 0 ) {
        result = LNS_STEM_ANY( lns_cloneBin2any( _pEnv, pBuf, readSize ) );
    }
    free( pBuf );

    return result;
}


static lns_stem_t mtd_lns_luaStream_seek(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2)
{
    lns_abort( __func__ );
    return lns_global.nilStem;
}

static lns_stem_t mtd_lns_luaStream_write(
    lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1)
{
    if ( pObj == l_var_lns_io_stdin ) {
        lns_abort( __func__ );
    }
    if ( arg1->type != lns_value_type_str ) {
        lns_abort( __func__ );
    }

    FILE * handle;
    if ( pObj == l_var_lns_io_stderr ) {
        handle = stderr;
    }
    else {
        handle = stdout;
    }

    lns_str_t * pStr = &arg1->val.str;
    fwrite( pStr->pStr, 1, pStr->len, handle );

    return LNS_STEM_ANY( pObj );
}
