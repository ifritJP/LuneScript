#include <lunescript.h>


    /**
     * リストの末尾に STEM を追加。
     */
#define __lune_add2list( TOP, STEM )            \
    STEM->pNext = TOP;                          \
    STEM->pPrev = (TOP)->pPrev;                 \
    (TOP)->pPrev->pNext = STEM;                 \
    (TOP)->pPrev = STEM;

    /**
     * リストから STEM を除外。
     */
#define __lune_rmFromList( STEM )               \
    if ( (STEM)->pNext != NULL ) {              \
        (STEM)->pPrev->pNext = (STEM)->pNext;   \
        (STEM)->pNext->pPrev = (STEM)->pPrev;   \
        (STEM)->pNext = NULL;                   \
    }


/**
 * type の値を保持する stem 型を生成する。
 *
 * 実際の type の値は別途セットする必要がある。
 *
 * この関数で確保した値は、 _pEnv で管理するブロックに紐付けられる。
 * この時点では、参照カウントは 0。
 * __lune_setq() で、何らかの変数に格納しない限り、ブロック終了時に開放される。
 *
 * @return 生成した stem 値
 */
static __lune_stem_t * __lune_alloc_stem( __lune_env_t * _pEnv, __lune_value_type_t type )
{
    _pEnv->allocNum++;
    
    __lune_stem_t * pStem = (__lune_stem_t *)malloc( sizeof( __lune_stem_t ) );
    pStem->type = type;
    pStem->refCount = 0;

    __lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    __lune_add2list( &pBlock->unassignStemTop, pStem );

    return pStem;
}


/**
 * pStem を開放する。
 *
 * 参照カウントが 0 になっていることが前提条件。
 */
static void __lune_release_stem( __lune_env_t * _pEnv, __lune_stem_t * pStem ) {
    free( pStem );
    _pEnv->allocNum--;
}

static void __lune_gc_stem( __lune_env_t * _pEnv, __lune_stem_t * pStem ) {
    switch ( pStem->type ) {
    case __lune_value_type_str:
        if ( !pStem->val.str.staticFlag ) {
            free( (char *)pStem->val.str.pStr );
        }
        break;
    case __lune_value_type_class:
        if ( ((__lune_Class_t*)pStem->val.classVal)->pMtd->_gc != NULL ) {
            ((__lune_Class_t*)pStem->val.classVal)->pMtd->_gc( _pEnv, pStem, true );
        }
        break;
    case __lune_value_type_ddd:
        free( pStem->val.ddd.pStemList );
        _pEnv->allocNum--;
        break;
    case __lune_value_type_form:
        free( pStem->val.form.pStemList );
        _pEnv->allocNum--;
        break;
    default:
        break;
    }
    __lune_release_stem( _pEnv, pStem );
}

/**
 * pStem の参照カウントをデクリメントする。
 *
 * 参照カウントが 0 になった場合は、開放処理を行なう。
 */
void __lune_decre_ref( __lune_env_t * _pEnv, __lune_stem_t * pStem ) {
    pStem->refCount--;
    if ( pStem->refCount == 0 && !pStem->retValFlag ) {
        __lune_gc_stem( _pEnv, pStem );
    }
}

void __lune_setQ_( __lune_stem_t * pStem )
{
    pStem->refCount++;
    if ( pStem->retValFlag ) {
        pStem->retValFlag = false;
    }
    __lune_rmFromList( pStem );
}

__lune_stem_t * __lune_setRet( __lune_env_t * _pEnv, __lune_stem_t * pStem )
{
    __lune_rmFromList( pStem );

    __lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth - 1 ];
    __lune_add2list( &pBlock->retValUnassignStemTop, pStem );
    pStem->retValFlag = true;

    return pStem;
}


/**
 * 新しくブロックを開始する。
 *
 * @param stemVerNum ブロックで管理する stem 型の値の数
 * @return ブロック情報
 */
__lune_block_t * __lune_enter_block( __lune_env_t * _pEnv, int stemVerNum )
{
    int dummy;
    _pEnv->blockDepth++;
    __lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];

    pBlock->unassignStemTop.pPrev = &pBlock->unassignStemTop;
    pBlock->unassignStemTop.pNext = &pBlock->unassignStemTop;

    pBlock->retValUnassignStemTop.pPrev = &pBlock->retValUnassignStemTop;
    pBlock->retValUnassignStemTop.pNext = &pBlock->retValUnassignStemTop;

    pBlock->pStemBuf = _pEnv->stemPPool + _pEnv->useStemPoolNum;
    memset( pBlock->pStemBuf, 0, sizeof( *pBlock->pStemBuf ) * stemVerNum );
    pBlock->len = stemVerNum;
    pBlock->blockDepth = _pEnv->blockDepth;
    pBlock->pStackAddr = &dummy;
    
    _pEnv->useStemPoolNum += stemVerNum;

    return pBlock;
}

/**
 * 現在のブロックを終了する。
 *
 * 次の処理を行なう。
 * - ブロックで管理している stem 型の値の参照カウンタをデクリメント
 * - 変数にアサインされていない stem 型の値を開放
 */
void __lune_leave_block( __lune_env_t * _pEnv )
{
    __lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    _pEnv->blockDepth--;
    

    int index;
    for ( index = 0; index < pBlock->len; index++ ) {
        __lune_decre_ref( _pEnv, pBlock->pStemBuf[ index ] );
    }

    {
        __lune_stem_t * pWork = pBlock->unassignStemTop.pPrev;
        while ( pWork != &pBlock->unassignStemTop ) {
            __lune_stem_t * pPrev = pWork->pPrev;

            __lune_gc_stem( _pEnv, pWork );
            pWork = pPrev;
        }
    }

    {
        __lune_stem_t * pWork = pBlock->retValUnassignStemTop.pPrev;
        while ( pWork != &pBlock->retValUnassignStemTop ) {
            __lune_stem_t * pPrev = pWork->pPrev;

            pWork->retValFlag = false;
            if ( pWork->refCount == 0 ) {
                __lune_gc_stem( _pEnv, pWork );
            }
            pWork = pPrev;
        }
    }
    

    _pEnv->useStemPoolNum -= pBlock->len;
}

/**
 * 複数のブロックを抜ける。
 *
 * break や return の時に、一度に複数のブロックを抜けるときの処理。
 *
 * @param num ブロックの数
 */
void __lune_leave_blockMulti( __lune_env_t * _pEnv, int num )
{
    while ( num > 0 ) {
        __lune_leave_block( _pEnv );
        num--;
    }
}

/**
 * 関数開始時の処理
 *
 * __lune_enter_block() とほぼ同じ。
 * ただし、関数の引数の処理が追加。
 * 関数終了時は、 __lune_leave_block() をコールする。
 *
 * @param num __lune_enter_block() の stemVerNum と同じ。
 *   ブロック内の stem 値の数と、 argNum を合せた数が num になる。
 * @param argNum 引数の stem 型の値の数
 * @param ... 引数の stem 型の値。
 */
__lune_block_t * __lune_enter_func( __lune_env_t * _pEnv, int num, int argNum, ... )
{
    __lune_enter_block( _pEnv, num );
    __lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];


    va_list ap;
    va_start( ap, argNum );

    int index;
    for ( index = 0; index < argNum; index++ ) {
        __lune_stem_t * pStem = va_arg( ap, __lune_stem_t * );
        __lune_setq( _pEnv, pBlock->pStemBuf[ index ], pStem );
    }
    va_end(ap);

    return pBlock;
}

/**
 * ... の値を生成する
 *
 * ... に含める値は全て stem に変換する必要がある。
 *
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
__lune_stem_t * __lune_createDDD( __lune_env_t * _pEnv, int num, ... ) {
    __lune_stem_t * pDDDStem = __lune_alloc_stem( _pEnv, __lune_value_type_ddd );
    __lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    pDDD->len = num;
    pDDD->pStemList = (__lune_stem_t **)malloc( sizeof( __lune_stem_t * ) * num );
    _pEnv->allocNum++;

    va_list ap;
    va_start( ap, num );

    int index;
    for ( index = 0; index < num; index++ ) {
        __lune_stem_t * pStem = va_arg( ap, __lune_stem_t * );
        pDDD->pStemList[ index ] = pStem;
    }
    va_end(ap);

    return pDDDStem;
}


/**
 * ... の値を生成する
 *
 * ... に含める値は全て stem に変換する必要がある。
 *
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
__lune_stem_t * __lune_createDDDOnly( __lune_env_t * _pEnv, int num ) {
    __lune_stem_t * pDDDStem = __lune_alloc_stem( _pEnv, __lune_value_type_ddd );
    __lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    pDDD->len = num;
    pDDD->pStemList = (__lune_stem_t **)malloc( sizeof( __lune_stem_t * ) * num );
    _pEnv->allocNum++;

    return pDDDStem;
}

/**
 * クラスのインスタンスを保持する stem を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return stem
 */
__lune_stem_t * __lune_class_new( __lune_env_t * _pEnv, int size )
{
    __lune_stem_t * pStem = __lune_alloc_stem( _pEnv, __lune_value_type_class );
    void * pObj = malloc( size );
    _pEnv->allocNum++;
    pStem->val.classVal = pObj;
    return pStem;
}

/**
 * クラスのインスタンスを開放する
 *
 * @param pObj クラスのインスタンス
 */
void __lune_class_del( __lune_env_t * _pEnv, void * pObj )
{
    _pEnv->allocNum--;
    free( pObj );
}

/**
 * 関数 pFunc を保持する stem を生成する
 *
 * @param pFunc 関数
 * @param num フォーム内でアクセスする外部変数の数
 * @param ... フォーム内でアクセスする外部変数
 * @return stem
 */
__lune_stem_t * __lune_func2stem(
    __lune_env_t * _pEnv, __lune_func_t * pFunc, int num, ... )
{
    __lune_stem_t * pFormStem = __lune_alloc_stem( _pEnv, __lune_value_type_form );
    __lune_form_t * pForm = &pFormStem->val.form;
    pForm->len = num;
    pForm->pStemList = (__lune_stem_t **)malloc( sizeof( __lune_stem_t * ) * num );
    _pEnv->allocNum++;

    pForm->pFunc = pFunc;
    
    va_list ap;
    va_start( ap, num );

    int index;
    for ( index = 0; index < num; index++ ) {
        __lune_stem_t * pStem = va_arg( ap, __lune_stem_t * );
        pForm->pStemList[ index ] = pStem;
    }
    va_end(ap);

    return pFormStem;
}


/**
 * int 値 val を保持する stem を生成する
 *
 * @param val int 値
 * @return stem
 */
__lune_stem_t * __lune_int2stem( __lune_env_t * _pEnv, __lune_int_t val ) {
    __lune_stem_t * pStem = __lune_alloc_stem( _pEnv, __lune_value_type_int );
    pStem->val.intVal = val;
    return pStem;
}

/**
 * リテラルな文字列型を生成する
 *
 * @param pStr 文字列
 * @return 文字列型データ
 */
__lune_str_t __lune_createLiteralStr( const char * pStr ) {
    __lune_str_t str;
    str.len = strlen( pStr );
    str.pStr = pStr;
    str.staticFlag = true;
    return str;
}

/**
 * 文字列データを保持する stem を生成する
 *
 * @param val 文字列型データ
 * @return stem
 */
__lune_stem_t * __lune_str2stem( __lune_env_t * _pEnv, __lune_str_t val ) {
    __lune_stem_t * pStem = __lune_alloc_stem( _pEnv, __lune_value_type_str );
    pStem->val.str = val;
    return pStem;
}

/**
 * 環境を生成する。
 *
 * この環境は、スレッド毎に 1 つ。
 *
 * @return 環境。
 */
static __lune_env_t * __lune_createEnv() {
    __lune_env_t * _pEnv = (__lune_env_t *)malloc( sizeof( __lune_env_t ) );
    _pEnv->useStemPoolNum = 0;
    _pEnv->allocNum = 0;
    _pEnv->blockDepth = 0;

    __lune_enter_block( _pEnv, 0 );
    _pEnv->pNoneStem = __lune_alloc_stem( _pEnv, __lune_value_type_none );

    __lune_enter_block( _pEnv, 0 );
    return _pEnv;
}

/**
 * 環境を開放する。
 *
 * @param _pEnv 環境
 */
static void __lune_deleteEnv( __lune_env_t * _pEnv ) {
    __lune_leave_block( _pEnv );
    __lune_leave_block( _pEnv );

    printf( "-------------\n" );
    printf( "allocNum = %d\n", _pEnv->allocNum );
    printf( "useStemPoolNum = %d\n", _pEnv->useStemPoolNum );
    printf( "blockDepth = %d\n", _pEnv->blockDepth );

    free( _pEnv );
}


/**
 * lua の print() に相当する処理。
 *
 * @param pArg は ddd 型。
 */
void __lune_print( __lune_env_t * _pEnv, __lune_stem_t * pArg ) {

    __lune_enter_func( _pEnv, 1, 1, pArg );

    __lune_ddd_t * pDDD = &pArg->val.ddd;
    int index;
    for ( index = 0; index < pDDD->len; index++ ) {
        if ( index > 0 ) {
            printf( "\t" );
        }
        __lune_stem_t * pStem = pDDD->pStemList[ index ];
        switch ( pStem->type ) {
        case __lune_value_type_int:
            printf( "%d", pStem->val.intVal );
            break;
        case __lune_value_type_real:
            printf( "%g", pStem->val.realVal );
            break;
        case __lune_value_type_str:
            printf( "%s", pStem->val.str.pStr );
            break;
        case __lune_value_type_form:
            printf( "form: %p", pStem );
            break;
        default:
            printf( "unknown type -- %d", pStem->type );
            break;
        }
    }
    printf( "\n" );

    __lune_leave_block( _pEnv );
}



int main() {
    __lune_env_t * _pEnv = __lune_createEnv();
    
    __lune_init_test( _pEnv );

    __lune_deleteEnv( _pEnv );

    return 0;
}
