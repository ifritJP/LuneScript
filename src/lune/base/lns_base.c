#include <lunescript.h>

    /**
     * リストの末尾に STEM を追加。
     */
#define lune_add2list( TOP, STEM )            \
    STEM->pNext = TOP;                          \
    STEM->pPrev = (TOP)->pPrev;                 \
    (TOP)->pPrev->pNext = STEM;                 \
    (TOP)->pPrev = STEM;

    /**
     * リストから STEM を除外。
     */
#define lune_rmFromList( STEM )               \
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
 * lune_setq() で、何らかの変数に格納しない限り、ブロック終了時に開放される。
 *
 * @return 生成した stem 値
 */
static lune_stem_t * lune_alloc_stem( lune_env_t * _pEnv, lune_value_type_t type )
{
    _pEnv->allocNum++;
    
    lune_stem_t * pStem = (lune_stem_t *)malloc( sizeof( lune_stem_t ) );
    pStem->type = type;
    pStem->refCount = 1;
    pStem->pEnv = _pEnv;

    // 現在のブロックに登録
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lune_add2list( &pBlock->managedStemTop, pStem );

    return pStem;
}


/**
 * pStem を開放する。
 *
 * 参照カウントが 0 になっていることが前提条件。
 */
static void lune_release_stem( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    free( pStem );
    _pEnv->allocNum--;
}

static void lune_gc_stem( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    switch ( pStem->type ) {
    case lune_value_type_str:
        if ( !pStem->val.str.staticFlag ) {
            free( (char *)pStem->val.str.pStr );
        }
        break;
    case lune_value_type_class:
        if ( ((lune_Class_t*)pStem->val.classVal)->pMtd->_gc != NULL ) {
            ((lune_Class_t*)pStem->val.classVal)->pMtd->_gc( _pEnv, pStem, true );
        }
        break;
    case lune_value_type_ddd: // fall-through
    case lune_value_type_mRet:
        free( pStem->val.ddd.pStemList );
        _pEnv->allocNum--;
        break;
    case lune_value_type_form:
        {
            int index;
            for ( index = 0; index < pStem->val.form.len; index++ ) {
                lune_decre_ref( _pEnv, pStem->val.form.pStemList[ index ] );
            }
            free( pStem->val.form.pStemList );
            _pEnv->allocNum--;
        }
        break;
    case lune_value_type_itSet:
        lune_itSet_gc( _pEnv, pStem );
        break;
    default:
        break;
    }
    lune_release_stem( _pEnv, pStem );
}

/**
 * pStem の参照カウントをデクリメントする。
 *
 * 参照カウントが 0 になった場合は、開放処理を行なう。
 */
void lune_decre_ref( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    pStem->refCount--;
    if ( pStem->refCount == 0 ) {
        lune_gc_stem( _pEnv, pStem );
    }
}

void lune_setQ_( lune_stem_t * pStem )
{
    pStem->refCount++;
}

lune_stem_t * lune_setRet( lune_env_t * _pEnv, lune_stem_t * pStem )
{
    lune_rmFromList( pStem );

    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth - 1 ];
    lune_add2list( &pBlock->managedStemTop, pStem );

    if ( pStem->type == lune_value_type_ddd ||
         pStem->type == lune_value_type_mRet )
    {
        int index;
        for ( index = 0; index < pStem->val.ddd.len; index++ ) {
            lune_setRet( _pEnv, lune_fromDDD( pStem, index ) );
        }
    }
    return pStem;
}


/**
 * 新しくブロックを開始する。
 *
 * @param stemVerNum ブロックで管理する stem 型の値の数
 * @return ブロック情報
 */
lune_block_t * lune_enter_block( lune_env_t * _pEnv, int stemVerNum )
{
    int dummy;
    _pEnv->blockDepth++;
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];

    pBlock->managedStemTop.pPrev = &pBlock->managedStemTop;
    pBlock->managedStemTop.pNext = &pBlock->managedStemTop;

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
void lune_leave_block( lune_env_t * _pEnv )
{
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    _pEnv->blockDepth--;
    

    int index;
    for ( index = 0; index < pBlock->len; index++ ) {
        lune_stem_t * pStem = pBlock->pStemBuf[ index ];
        if ( pStem != NULL ) {
            lune_decre_ref( _pEnv, pStem );
        }
    }

    {
        lune_stem_t * pWork = pBlock->managedStemTop.pPrev;
        while ( pWork != &pBlock->managedStemTop ) {
            lune_stem_t * pPrev = pWork->pPrev;

            if ( pWork->refCount == 1 ) {
                lune_gc_stem( _pEnv, pWork );
            }
            else {
                pWork->refCount--;
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
void lune_leave_blockMulti( lune_env_t * _pEnv, int num )
{
    while ( num > 0 ) {
        lune_leave_block( _pEnv );
        num--;
    }
}

/**
 * 関数開始時の処理
 *
 * lune_enter_block() とほぼ同じ。
 * ただし、関数の引数の処理が追加。
 * 関数終了時は、 lune_leave_block() をコールする。
 *
 * @param num lune_enter_block() の stemVerNum と同じ。
 *   ブロック内の stem 値の数と、 argNum を合せた数が num になる。
 * @param argNum 引数の stem 型の値の数
 * @param ... 引数の stem 型の値。
 */
lune_block_t * lune_enter_func( lune_env_t * _pEnv, int num, int argNum, ... )
{
    lune_enter_block( _pEnv, num );
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];


    va_list ap;
    va_start( ap, argNum );

    int index;
    for ( index = 0; index < argNum; index++ ) {
        lune_stem_t * pStem = va_arg( ap, lune_stem_t * );
        lune_setq( _pEnv, pBlock->pStemBuf[ index ], pStem );
    }
    va_end(ap);

    return pBlock;
}

lune_stem_t * lune_getValFromDDD( lune_env_t * _pEnv, lune_stem_t * pStem, int index )
{
    if ( pStem->val.ddd.len <= index ) {
        return _pEnv->pNilStem;
    }
    return pStem->val.ddd.pStemList[ index ];
}


static lune_stem_t * lune_createDDDSub(
    lune_env_t * _pEnv, bool hasDDD, int num, va_list apSrc, lune_stem_t * pDDDStem ) {
    lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    int argNum = num;

    if ( hasDDD ) {
        argNum = num;

        va_list ap;
        va_copy( ap, apSrc );

        int index;
        for ( index = 0; index < num - 1; index++ ) {
            va_arg( ap, lune_stem_t * );
        }
        lune_stem_t * pDDDStem = va_arg( ap, lune_stem_t * );
        argNum += pDDDStem->val.ddd.len - 1;
        va_end(ap);
    }
    
    pDDD->len = argNum;
    pDDD->pStemList = (lune_stem_t **)malloc( sizeof( lune_stem_t * ) * argNum );
    _pEnv->allocNum++;

    va_list ap;
    va_copy( ap, apSrc );

    int index;
    for ( index = 0; index < num - 1; index++ ) {
        lune_stem_t * pStem = va_arg( ap, lune_stem_t * );
        if ( pStem->type == lune_value_type_ddd ||
             pStem->type == lune_value_type_mRet ) {
            pDDD->pStemList[ index ] = lune_getValFromDDD( _pEnv, pStem, 0 );
        }
        else {
            pDDD->pStemList[ index ] = pStem;
        }
    }

    if ( hasDDD ) {
        lune_stem_t * pStem = va_arg( ap, lune_stem_t * );
        lune_stem_t ** ppStem = &pDDD->pStemList[ num - 1 ];
        for ( index = 0; index < pStem->val.ddd.len; index++ ) {
            *ppStem = pStem->val.ddd.pStemList[ index ];
            ppStem++;
        }
    }
    else {
        lune_stem_t * pStem = va_arg( ap, lune_stem_t * );
        pDDD->pStemList[ num - 1 ] = pStem;
    }

    
    va_end(ap);

    return pDDDStem;
}

/**
 * ... の値を生成する
 *
 * ... に含める値は全て stem に変換する必要がある。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t * lune_createMRet( lune_env_t * _pEnv, bool hasDDD, int num, ... ) {
    lune_stem_t * pDDDStem = lune_alloc_stem( _pEnv, lune_value_type_mRet );

    va_list ap;
    va_start( ap, num );

    lune_createDDDSub( _pEnv, hasDDD, num, ap, pDDDStem );

    va_end(ap);

    return pDDDStem;
}


/**
 * ... の値を生成する
 *
 * ... に含める値は全て stem に変換する必要がある。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t * lune_createDDD( lune_env_t * _pEnv, bool hasDDD, int num, ... ) {
    lune_stem_t * pDDDStem = lune_alloc_stem( _pEnv, lune_value_type_ddd );

    va_list ap;
    va_start( ap, num );

    lune_createDDDSub( _pEnv, hasDDD, num, ap, pDDDStem );

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
lune_stem_t * lune_createDDDOnly( lune_env_t * _pEnv, int num ) {
    lune_stem_t * pDDDStem = lune_alloc_stem( _pEnv, lune_value_type_ddd );
    lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    pDDD->len = num;
    pDDD->pStemList = (lune_stem_t **)malloc( sizeof( lune_stem_t * ) * num );
    _pEnv->allocNum++;

    return pDDDStem;
}


lune_stem_t * lune_it_new(
    lune_env_t * _pEnv, lune_value_type_t type, void * pVal )
{
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, type );
    switch ( type ) {
    case lune_value_type_itSet:
        pStem->val.itSet = pVal;
        break;
    case lune_value_type_itMap:
        pStem->val.itMap = pVal;
        break;
    default:
        break;
    }
    return pStem;
}

void lune_it_delete( lune_env_t * _pEnv, lune_stem_t * pStem )
{
    lune_rmFromList( pStem );
    lune_gc_stem( _pEnv, pStem );
}

/**
 * クラスのインスタンスを保持する stem を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return stem
 */
lune_stem_t * lune_class_new( lune_env_t * _pEnv, int size )
{
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_class );
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
void lune_class_del( lune_env_t * _pEnv, void * pObj )
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
lune_stem_t * lune_func2stem(
    lune_env_t * _pEnv, lune_func_t * pFunc, int num, ... )
{
    lune_stem_t * pFormStem = lune_alloc_stem( _pEnv, lune_value_type_form );
    lune_form_t * pForm = &pFormStem->val.form;
    pForm->len = num;
    pForm->pStemList = (lune_stem_t **)malloc( sizeof( lune_stem_t * ) * num );
    _pEnv->allocNum++;

    pForm->pFunc = pFunc;
    
    va_list ap;
    va_start( ap, num );

    int index;
    for ( index = 0; index < num; index++ ) {
        lune_stem_t * pStem = va_arg( ap, lune_stem_t * );
        pForm->pStemList[ index ] = pStem;
    }
    va_end(ap);

    return pFormStem;
}


/**
 * bool 値 val を保持する stem を生成する
 *
 * @param val bool 値
 * @return stem
 */
lune_stem_t * lune_bool2stem( lune_env_t * _pEnv, lune_bool_t val )
{
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_bool );
    pStem->val.boolVal = val;
    return pStem;
}

/**
 * int 値 val を保持する stem を生成する
 *
 * @param val int 値
 * @return stem
 */
lune_stem_t * lune_int2stem( lune_env_t * _pEnv, lune_int_t val ) {
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_int );
    pStem->val.intVal = val;
    return pStem;
}

/**
 * real 値 val を保持する stem を生成する
 *
 * @param val real 値
 * @return stem
 */
lune_stem_t * lune_real2stem( lune_env_t * _pEnv, lune_real_t val ) {
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_real );
    pStem->val.realVal = val;
    return pStem;
}

/**
 * リテラルな文字列型を生成する
 *
 * @param pStr 文字列
 * @return 文字列型データ
 */
lune_str_t lune_createLiteralStr( const char * pStr ) {
    lune_str_t str;
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
lune_stem_t * lune_str2stem( lune_env_t * _pEnv, lune_str_t val ) {
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_str );
    pStem->val.str = val;
    return pStem;
}

lune_stem_t * lune_litStr2stem( lune_env_t * _pEnv, const char * pStr )
{
    lune_stem_t * pStem = lune_alloc_stem( _pEnv, lune_value_type_str );
    pStem->val.str.len = strlen( pStr );
    pStem->val.str.pStr = pStr;
    pStem->val.str.staticFlag = true;
    return pStem;
}

/**
 * 環境を生成する。
 *
 * この環境は、スレッド毎に 1 つ。
 *
 * @return 環境。
 */
static lune_env_t * lune_createEnv() {
    lune_env_t * _pEnv = (lune_env_t *)malloc( sizeof( lune_env_t ) );
    _pEnv->useStemPoolNum = 0;
    _pEnv->allocNum = 0;
    _pEnv->blockDepth = 0;

    lune_enter_block( _pEnv, 0 );
    _pEnv->pNoneStem = lune_alloc_stem( _pEnv, lune_value_type_none );
    _pEnv->pNilStem = lune_alloc_stem( _pEnv, lune_value_type_nil );

    _pEnv->pSortCallback = NULL;

    lune_enter_block( _pEnv, 0 );
    return _pEnv;
}

/**
 * 環境を開放する。
 *
 * @param _pEnv 環境
 */
static void lune_deleteEnv( lune_env_t * _pEnv ) {
    lune_leave_block( _pEnv );
    lune_leave_block( _pEnv );

    printf( ":debug:allocNum = %d\n", _pEnv->allocNum );
    printf( ":debug:useStemPoolNum = %d\n", _pEnv->useStemPoolNum );
    printf( ":debug:blockDepth = %d\n", _pEnv->blockDepth );

    free( _pEnv );
}

/**
 * lua の print() に相当する処理。
 *
 * @param pArg は ddd 型。
 */
void lune_print( lune_env_t * _pEnv, lune_stem_t * _pForm, lune_stem_t * pArg ) {

    lune_enter_func( _pEnv, 1, 1, pArg );

    lune_ddd_t * pDDD = &pArg->val.ddd;
    int index;
    for ( index = 0; index < pDDD->len; index++ ) {
        if ( index > 0 ) {
            printf( "\t" );
        }
        lune_stem_t * pStem = pDDD->pStemList[ index ];
        switch ( pStem->type ) {
        case lune_value_type_nil:
            printf( "nil" );
            break;
        case lune_value_type_bool:
            printf( pStem->val.boolVal ? "true" : "false" );
            break;
        case lune_value_type_int:
            printf( "%d", pStem->val.intVal );
            break;
        case lune_value_type_real:
            printf( "%g", pStem->val.realVal );
            break;
        case lune_value_type_str:
            printf( "%s", pStem->val.str.pStr );
            break;
        case lune_value_type_form:
            printf( "form: %p", pStem );
            break;
        default:
            printf( "unknown type -- %d", pStem->type );
            break;
        }
    }
    printf( "\n" );

    lune_leave_block( _pEnv );
}



int main() {
    lune_env_t * _pEnv = lune_createEnv();
    
    lune_init_test( _pEnv );

    lune_deleteEnv( _pEnv );

    return 0;
}
