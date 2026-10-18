/*
  LuneScript から変換した wasm (WasmGC) が利用するランタイム (入出力など)。

  wasi-sdk で WASI のコマンド (main を持つ wasm) にし、このファイルの関数を export する。
  変換した wasm は、この wasm の memory と関数 (モジュール名 "lnsrt") を import して使い、
  main() は変換した wasm の main (モジュール名 "lnsmain") を呼び出す。
  Lua を使う場合は、lnslua.c と Lua を一緒にリンクした wasm (lnslua.wasm) を使う。

  ファイル・時刻・引数などは wasi-libc (WASI) で処理するので、
  Node.js (node:wasi) と WASI のランタイム (wasmtime など) のどちらでも動く。
  Node.js でしかできない処理 (io.popen) は、モジュール "lnsjs" の関数を使う
  (WASI のランタイムでは、何もしない関数 (lnsjs_stub.wat) を使う)。

  文字列の受け渡し:
   - 変換した wasm → C: lnsr_buf() の領域に書き込み、(アドレス, 長さ) を渡す
   - C → 変換した wasm: 長さを返し (nil は -1)、内容は lnsr_res() の領域に置く
*/

#include <errno.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <time.h>
#include <unistd.h>

#include "lnsrt.h"

#define IMPORT(module, name) __attribute__((import_module(#module), import_name(#name)))

/* 変換した wasm の main */
IMPORT(lnsmain, main) void lnsmain_main( void );

/*
  io.popen (Node.js でだけ使える)。
  lnsjs_popen() でコマンドを実行し、出力の長さを返す (失敗は -1)。
  lnsjs_popen_data() で、その出力を dst に書き込む。
*/
IMPORT(lnsjs, popen) int32_t lnsjs_popen( const char * cmd, int32_t len );
IMPORT(lnsjs, popen_data) void lnsjs_popen_data( char * dst );

/* 文字列の受け渡し用の領域 (変換した wasm → C) */
static char * s_buf;
static size_t s_bufSize;

/* 文字列の受け渡し用の領域 (C → 変換した wasm) */
static char * s_res;
static size_t s_resSize;

static int s_argc;
static char ** s_argv;

/* 直前のファイル操作の失敗のメッセージ */
static char s_fileError[ 1024 ];

/* io.open などで開いたファイル (ファイル番号が添字) */
static FILE ** s_fileList;
static int32_t s_fileNum;

/* size バイト以上の領域を確保する */
static char * reserve( char ** buf, size_t * bufSize, size_t size )
{
    if ( size > *bufSize ) {
        size_t newSize = *bufSize == 0 ? 256 : *bufSize;
        while ( newSize < size ) {
            newSize *= 2;
        }
        *buf = realloc( *buf, newSize );
        if ( *buf == NULL ) {
            abort();
        }
        *bufSize = newSize;
    }
    return *buf;
}

/* size バイト以上の受け渡し用の領域 (変換した wasm → C) を返す */
char * EXPORT(lnsr_buf)( int32_t size )
{
    return reserve( &s_buf, &s_bufSize, size );
}

/* 戻り値の文字列の領域 (C → 変換した wasm) を返す */
char * EXPORT(lnsr_res)( void )
{
    return s_res;
}

/* txt の len バイトを戻り値の文字列にして、長さを返す */
static int32_t setRes( const char * txt, size_t len )
{
    memcpy( reserve( &s_res, &s_resSize, len + 1 ), txt, len );
    return (int32_t)len;
}

/* (アドレス, 長さ) の文字列を、終端が '\0' の文字列にする (free() する) */
static char * toCStr( const char * txt, int32_t len )
{
    char * str = malloc( len + 1 );
    memcpy( str, txt, len );
    str[ len ] = '\0';
    return str;
}

/* Lua の luaL_fileresult() と同じエラーのメッセージ ("名前: 理由") */
static void setFileError( const char * name )
{
    snprintf( s_fileError, sizeof( s_fileError ), "%s: %s", name, strerror( errno ) );
}

void EXPORT(lnsr_write)( const char * txt, int32_t len )
{
    fwrite( txt, 1, len, stdout );
}

void EXPORT(lnsr_write_err)( const char * txt, int32_t len )
{
    fflush( stdout );
    fwrite( txt, 1, len, stderr );
}

/*
  LuneScript の error()。メッセージを出力して終了する。

  環境変数 LNS_WASM_STACK=1 の場合は、wasm の関数のスタックを確認できるように trap する。
*/
void EXPORT(lnsr_error)( const char * txt, int32_t len )
{
    fflush( stdout );
    fprintf( stderr, "lua: %.*s\n", (int)len, txt );
    const char * stack = getenv( "LNS_WASM_STACK" );
    if ( stack != NULL && strcmp( stack, "1" ) == 0 ) {
        fflush( stderr );
        __builtin_trap();
    }
    exit( 1 );
}

/* ファイルを登録し、ファイル番号を返す */
static int32_t addFile( FILE * fp )
{
    s_fileList = realloc( s_fileList, sizeof( FILE * ) * ( s_fileNum + 1 ) );
    s_fileList[ s_fileNum ] = fp;
    return s_fileNum++;
}

/* io.open。ファイル番号を返す (失敗は -1 で、メッセージは lnsr_file_error() で取得する) */
int32_t EXPORT(lnsr_file_open)( const char * path, int32_t pathLen,
                                const char * mode, int32_t modeLen )
{
    char * pathStr = toCStr( path, pathLen );
    /* Lua の io.open と同じモード ("b" は無視する) */
    char modeStr[ 4 ];
    int pos = 0;
    int valid = modeLen <= 3;
    for ( int index = 0; valid && index < modeLen; index++ ) {
        if ( mode[ index ] != 'b' ) {
            modeStr[ pos++ ] = mode[ index ];
        }
    }
    modeStr[ pos ] = '\0';
    if ( valid ) {
        valid = ( pos == 1 || ( pos == 2 && modeStr[ 1 ] == '+' ) ) &&
            strchr( "rwa", modeStr[ 0 ] ) != NULL;
    }
    if ( !valid ) {
        char * modeAll = toCStr( mode, modeLen );
        snprintf( s_fileError, sizeof( s_fileError ), "invalid mode '%s'", modeAll );
        free( modeAll );
        free( pathStr );
        return -1;
    }
    FILE * fp = fopen( pathStr, modeStr );
    if ( fp == NULL ) {
        setFileError( pathStr );
        free( pathStr );
        return -1;
    }
    free( pathStr );
    return addFile( fp );
}

/* io.popen。コマンドの出力を読むファイルの番号を返す (失敗は -1) */
int32_t EXPORT(lnsr_file_popen)( const char * cmd, int32_t len )
{
    int32_t size = lnsjs_popen( cmd, len );
    if ( size < 0 ) {
        snprintf( s_fileError, sizeof( s_fileError ), "'popen' not supported" );
        return -1;
    }
    char * data = malloc( size + 1 );
    lnsjs_popen_data( data );
    /* 出力を読み終わるまで data は解放しない (fclose しても解放されない) */
    FILE * fp = fmemopen( data, size, "r" );
    if ( fp == NULL ) {
        setFileError( "popen" );
        free( data );
        return -1;
    }
    return addFile( fp );
}

/* 直前のファイル操作の失敗のメッセージ */
int32_t EXPORT(lnsr_file_error)( void )
{
    return setRes( s_fileError, strlen( s_fileError ) );
}

/* io.stdin */
int32_t EXPORT(lnsr_file_stdin)( void )
{
    return addFile( stdin );
}

/*
  ファイルを読む。読んだ文字列の長さを返す (無い場合は -1)。

  kind: 0: 全て ("a")、1: 1 行 ("l")、2: 改行付きの 1 行 ("L")、3: num バイト
*/
int32_t EXPORT(lnsr_file_read)( int32_t fileNo, int32_t kind, int64_t num )
{
    FILE * fp = s_fileList[ fileNo ];
    size_t len = 0;
    switch ( kind ) {
    case 0: {
        for (;;) {
            char * dst = reserve( &s_res, &s_resSize, len + 4096 );
            size_t size = fread( dst + len, 1, 4096, fp );
            len += size;
            if ( size < 4096 ) {
                break;
            }
        }
        if ( ferror( fp ) ) {
            clearerr( fp );
            return -1;
        }
        return (int32_t)len;
    }
    case 1:
    case 2: {
        int ch;
        while ( ( ch = getc( fp ) ) != EOF ) {
            if ( ch == '\n' && kind == 1 ) {
                break;
            }
            reserve( &s_res, &s_resSize, len + 1 )[ len++ ] = (char)ch;
            if ( ch == '\n' ) {
                break;
            }
        }
        if ( ferror( fp ) ) {
            clearerr( fp );
            return -1;
        }
        if ( ch == EOF && len == 0 ) {
            return -1;
        }
        return (int32_t)len;
    }
    default: {
        if ( num <= 0 ) {
            return 0;
        }
        char * dst = reserve( &s_res, &s_resSize, (size_t)num );
        len = fread( dst, 1, (size_t)num, fp );
        if ( ferror( fp ) ) {
            clearerr( fp );
            return -1;
        }
        if ( len == 0 ) {
            return -1;
        }
        return (int32_t)len;
    }
    }
}

void EXPORT(lnsr_file_write)( int32_t fileNo, const char * txt, int32_t len )
{
    fwrite( txt, 1, len, s_fileList[ fileNo ] );
}

/* ファイルの seek。whence: 0: set、1: cur、2: end。位置を返す (失敗は -1) */
int64_t EXPORT(lnsr_file_seek)( int32_t fileNo, int32_t whence, int64_t offset )
{
    static const int whenceList[] = { SEEK_SET, SEEK_CUR, SEEK_END };
    FILE * fp = s_fileList[ fileNo ];
    if ( fseeko( fp, offset, whenceList[ whence ] ) != 0 ) {
        return -1;
    }
    return ftello( fp );
}

void EXPORT(lnsr_file_close)( int32_t fileNo )
{
    FILE * fp = s_fileList[ fileNo ];
    if ( fp != NULL && fp != stdin ) {
        fclose( fp );
    }
    s_fileList[ fileNo ] = NULL;
}

/* 最終更新日時 (1970/1/1 からの秒数)。取得できない場合は NaN */
double EXPORT(lnsr_file_mtime)( const char * path, int32_t len )
{
    char * pathStr = toCStr( path, len );
    struct stat info;
    int result = stat( pathStr, &info );
    free( pathStr );
    if ( result != 0 ) {
        return __builtin_nan( "" );
    }
    return (double)info.st_mtim.tv_sec;
}

/* コマンドの引数の数 (先頭は wasm のパス) */
int32_t EXPORT(lnsr_arg_count)( void )
{
    return s_argc;
}

int32_t EXPORT(lnsr_arg_get)( int32_t index )
{
    return setRes( s_argv[ index ], strlen( s_argv[ index ] ) );
}

void EXPORT(lnsr_os_exit)( int64_t code )
{
    exit( (int)code );
}

/* os.remove。成功は -1 (nil)、失敗はエラーのメッセージの長さ */
int32_t EXPORT(lnsr_os_remove)( const char * path, int32_t len )
{
    char * pathStr = toCStr( path, len );
    int32_t result = -1;
    if ( remove( pathStr ) != 0 ) {
        setFileError( pathStr );
        result = setRes( s_fileError, strlen( s_fileError ) );
    }
    free( pathStr );
    return result;
}

/* os.rename。成功は -1 (nil)、失敗はエラーのメッセージの長さ */
int32_t EXPORT(lnsr_os_rename)( const char * src, int32_t srcLen,
                                const char * dst, int32_t dstLen )
{
    char * srcStr = toCStr( src, srcLen );
    char * dstStr = toCStr( dst, dstLen );
    int32_t result = -1;
    if ( rename( srcStr, dstStr ) != 0 ) {
        setFileError( srcStr );
        result = setRes( s_fileError, strlen( s_fileError ) );
    }
    free( srcStr );
    free( dstStr );
    return result;
}

/* os.time() (引数無し) */
int64_t EXPORT(lnsr_os_time)( void )
{
    return (int64_t)time( NULL );
}

/* os.clock (プロセスの CPU 時間、秒) */
double EXPORT(lnsr_os_clock)( void )
{
    return (double)clock() / CLOCKS_PER_SEC;
}

__attribute__((weak)) void lnsr_initExt( void )
{
}

int main( int argc, char ** argv )
{
    s_argc = argc;
    s_argv = argv;
    /*
      WASI のカレントディレクトリは / なので、ホストが環境変数 LNS_WASM_CWD で指定した場合は、
      カレントディレクトリをそこにする (ホストは、そのディレクトリを含むディレクトリを preopen する)。
    */
    const char * cwd = getenv( "LNS_WASM_CWD" );
    if ( cwd != NULL ) {
        chdir( cwd );
    }
    lnsr_initExt();
    lnsmain_main();
    return 0;
}
