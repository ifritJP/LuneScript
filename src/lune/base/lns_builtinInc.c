#include <stdio.h>
#include <unistd.h>

static void lune_init_builtinSub( lune_env_t * _pEnv )
{
    lune_any_t * lune_stdin = lune_class_luaStream_new( _pEnv );
    lune_setQ_any( &lune_var_io_stdin, &lune_if_luaStream( lune_stdin )->iStream );
    lune_any_t * lune_stderr = lune_class_luaStream_new( _pEnv );
    lune_setQ_any( &lune_var_io_stderr, &lune_if_luaStream( lune_stderr )->oStream );
    lune_any_t * lune_stdout = lune_class_luaStream_new( _pEnv );
    lune_setQ_any( &lune_var_io_stdout, &lune_if_luaStream( lune_stdout )->oStream );
}

static void u_mtd_luaStream_close( lune_env_t * _pEnv, lune_any_t * pObj)
{
    if ( pObj == lune_var_io_stdin ) {
        fclose( stdin );
    }
    else if ( pObj == lune_var_io_stderr ) {
        fclose( stderr );
    }
    else if ( pObj == lune_var_io_stdout ) {
        fclose( stdout );
    }
    else {
        lune_abort( __func__ );
    }
}

static void u_mtd_luaStream_flush( lune_env_t * _pEnv, lune_any_t * pObj)
{
    if ( pObj == lune_var_io_stdin ) {
        lune_abort( __func__ );
    }
    else if ( pObj == lune_var_io_stderr ) {
        fflush( stderr );
    }
    else if ( pObj == lune_var_io_stdout ) {
        fflush( stdout );
    }
    else {
        lune_abort( __func__ );
    }
}

static lune_stem_t u_mtd_luaStream_read(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1)
{
    if ( pObj != lune_var_io_stdin ) {
        lune_abort( __func__ );
    }

    if ( arg1.type != lune_stem_type_int ) {
        lune_abort( __func__ );
    }

    int size = arg1.val.intVal;
    char * pBuf = malloc( size );
    if ( pBuf == NULL ) {
        lune_abort( __func__ );
    }
    int readSize = fread( pBuf, 1, size, stdin );

    lune_stem_t result = lune_global.nilStem;
    if ( readSize != 0 ) {
        result = LUNE_STEM_ANY( lune_cloneBin2any( _pEnv, pBuf, readSize ) );
    }
    free( pBuf );

    return result;
}


static lune_stem_t u_mtd_luaStream_seek(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1, lune_int_t arg2)
{
    lune_abort( __func__ );
    return lune_global.nilStem;
}

static lune_stem_t u_mtd_luaStream_write(
    lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1)
{
    if ( pObj == lune_var_io_stdin ) {
        lune_abort( __func__ );
    }
    if ( arg1->type != lune_value_type_str ) {
        lune_abort( __func__ );
    }

    FILE * handle;
    if ( pObj == lune_var_io_stderr ) {
        handle = stderr;
    }
    else {
        handle = stdout;
    }

    lune_str_t * pStr = &arg1->val.str;
    fwrite( pStr->pStr, 1, pStr->len, handle );

    return LUNE_STEM_ANY( pObj );
}
