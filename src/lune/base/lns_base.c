#include <lua.h>
#include <lauxlib.h>
#include <lunescript.h>
#include <math.h>
#include <stdint.h>

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

#define lune_malloc( ALLOCATOR, SIZE ) _lune_malloc( ALLOCATOR, SIZE, __FILE__, __LINE__ )
#define lune_free( ALLOCATOR, ADDR ) _lune_free( ALLOCATOR, ADDR, __FILE__, __LINE__ )

typedef struct lune_globalEnv_t {
    lune_env_t * pEnv;
    /** モジュールの init ブロックのバッファ */
    lune_block_t moduleInitBlockBuf[ LUNE_MODULE_MAX_NUM ];
    /** 現在のモジュール数 */
    int moduleNum;
    /** 現在確保している stem の数 */
    int allocNum;
} lune_globalEnv_t;

static lune_globalEnv_t s_globalEnv;

#define lune_alloc_stem( ENV, TYPE, FILE, LINE )             \
    _lune_alloc_stem( ENV, TYPE, FILE, LINE )
#define lune_alloc_stem_op( ENV, TYPE ) \
    _lune_alloc_stem( ENV, TYPE, __FILE__, __LINE__ )

#define lune_abort( MESS ) _lune_abort( MESS, LUNE_DEBUG_POS )

static void _lune_abort( const char * pMessage, const char * pFile, int lineNo )
{
    fprintf( stderr, "abort:%s:%d:%s\n", pFile, lineNo, pMessage );
    abort();
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
static lune_stem_t * _lune_alloc_stem(
    lune_env_t * _pEnv, lune_value_type_t type, const char * pFile, int lineNo )
{
    s_globalEnv.allocNum++;
    
    lune_stem_t * pStem = (lune_stem_t *)_lune_malloc(
        _pEnv->allocateor, sizeof( lune_stem_t ), pFile, lineNo );
    lune_init_stem( pStem, type, _pEnv );
    pStem->refCount = 1;

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
    if ( pStem->type == lune_value_type_str ) {
        pStem->type = lune_value_type_str;
    }
    pStem->type = lune_value_type_none; // このセットは必須ではないが、
                                        // セットしておいた方が不具合を見つけ易い。
    lune_free( _pEnv->allocateor, pStem );
    s_globalEnv.allocNum--;
}

static void lune_gc_stem( lune_env_t * _pEnv, lune_stem_t * pStem, bool freeFlag ) {
    switch ( pStem->type ) {
    case lune_value_type_str:
        if ( !pStem->val.str.staticFlag ) {
            lune_free( _pEnv->allocateor, (char *)pStem->val.str.pStr );
        }
        break;
    case lune_value_type_class:
        if ( ((lune_Class_t*)pStem->val.classVal)->pMtd->_gc != NULL ) {
            ((lune_Class_t*)pStem->val.classVal)->pMtd->_gc( _pEnv, pStem );
        }
        if ( ((lune_Class_t*)pStem->val.classVal)->pMtd->_del != NULL ) {
            ((lune_Class_t*)pStem->val.classVal)->pMtd->_del( _pEnv, pStem );
        }
        lune_class_del( _pEnv, pStem->val.classVal );
        break;
    case lune_value_type_ddd: // fall-through
    case lune_value_type_mRet:
        lune_free( _pEnv->allocateor, pStem->val.ddd.pStemList );
        s_globalEnv.allocNum--;
        break;
    case lune_value_type_form:
        {
            int index;
            for ( index = 0; index < pStem->val.form.len; index++ ) {
                lune_var_t * pVar = lune_form_closure( pStem, index );
                pVar->refCount--;
                if ( pVar->refCount == 0 ) {
                    lune_decre_ref( _pEnv, pVar->pStem );
                }
            }
            lune_free( _pEnv->allocateor, pStem->val.form.ppClosureValList );
            s_globalEnv.allocNum--;
        }
        break;
    case lune_value_type_itSet:
        lune_itSet__del( _pEnv, pStem );
        break;
    case lune_value_type_if:
        lune_decre_ref( _pEnv, pStem->val.ifVal.pObj );
        // if は class 内のメンバなので開放しないで return する。
        return;
    default:
        break;
    }
    if ( freeFlag ) {
        lune_release_stem( _pEnv, pStem );
    }
}

/**
 * pStem の参照カウントをデクリメントする。
 *
 * 参照カウントが 0 になった場合は、開放処理を行なう。
 */
void lune_decre_ref( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    pStem->refCount--;
    if ( pStem->refCount == 0 ) {
        lune_gc_stem( _pEnv, pStem, true );
    }
}

lune_stem_t * lune_setMRet( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    _pEnv->pMRet = pStem;
    return lune_fromDDD( pStem, 0 );
}

lune_stem_t * lune_getIF( lune_env_t * _pEnv, lune_stem_t * pIFStem )
{
    if ( pIFStem->refCount == 0 ) {
        // interface の stem は共有なので、
        // 参照回数が 0 の状態から 1 になったタイミングで、
        // 元インスタンスの参照回数をインクリメントする。
        pIFStem->val.ifVal.pObj->refCount++;

        lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
        lune_add2list( &pBlock->managedStemTop, pIFStem );
        pIFStem->refCount++;
    }
    return pIFStem;
}

/**
 * pStem が class ならば、指定の pMeta の IF のインスタンスを返す。
 * pStem が IF なら、そのまま pStem を返す。
 */
lune_stem_t * lune_toIF(
    lune_env_t * _pEnv, lune_stem_t * pStem, const lune_type_meta_t * pMeta )
{
    if ( pStem->type == lune_value_type_if ) {
        return pStem;
    }
    if ( pStem->type == lune_value_type_class ) {
        lune_stem_t * pIFStem = (lune_stem_t *)pStem->val.classVal->pIFdummy;
        for ( ; pIFStem->type != lune_value_type_nil; pIFStem++ ) {
            if ( pIFStem->val.ifVal.pMeta == pMeta ) {
                return lune_getIF( _pEnv, pIFStem );
            }
        }
    }
    lune_abort( "not found interface" );
    return NULL;
}

void lune_setQ_( lune_stem_t * pStem )
{
    pStem->refCount++;
}

lune_stem_t * lune_setRet( lune_env_t * _pEnv, lune_stem_t * pStem )
{
    if ( pStem->type == lune_value_type_bool ) {
        return pStem;
    }
    
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

void lune_setup_block( lune_env_t * _pEnv, lune_block_t * pBlock, int stemVerNum )
{
    int dummy;

    pBlock->managedStemTop.pPrev = &pBlock->managedStemTop;
    pBlock->managedStemTop.pNext = &pBlock->managedStemTop;

    pBlock->pVarList = _pEnv->varPPool + _pEnv->useStemPoolNum;
    int index;
    for ( index = 0; index < stemVerNum; index++ ) {
        lune_var_t * pVar = (lune_var_t *)lune_malloc(
            _pEnv->allocateor, sizeof( lune_var_t ) );
        s_globalEnv.allocNum++;
        pVar->pStem = NULL;
        pBlock->pVarList[ index ] = pVar;
        pVar->refCount = 1;
    }
    pBlock->len = stemVerNum;
    pBlock->blockDepth = _pEnv->blockDepth;
    pBlock->pStackAddr = &dummy;
    
    _pEnv->useStemPoolNum += stemVerNum;
}


/**
 * 新しくブロックを開始する。
 *
 * @param stemVerNum ブロックで管理する stem 型の値の数
 * @return ブロック情報
 */
lune_block_t * lune_enter_module( int stemVerNum )
{
    lune_block_t * pBlock = &s_globalEnv.moduleInitBlockBuf[ s_globalEnv.moduleNum ];
    s_globalEnv.moduleNum++;

    lune_setup_block( s_globalEnv.pEnv, pBlock, stemVerNum );

    return pBlock;
}


/**
 * 新しくブロックを開始する。
 *
 * @param stemVerNum ブロックで管理する stem 型の値の数
 * @return ブロック情報
 */
lune_block_t * lune_enter_block( lune_env_t * _pEnv, int stemVerNum )
{
    _pEnv->blockDepth++;
    
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lune_setup_block( _pEnv, pBlock, stemVerNum );

    return pBlock;
}


static inline void lune_reset_blockSub( lune_env_t * _pEnv, lune_block_t * pBlock ) {
    lune_stem_t * pWork = pBlock->managedStemTop.pPrev;
    while ( pWork != &pBlock->managedStemTop ) {
        lune_stem_t * pPrev = pWork->pPrev;

        if ( pWork->refCount == 1 ) {
            lune_gc_stem( _pEnv, pWork, true );
        }
        else {
            pWork->refCount--;
        }
        pWork = pPrev;
    }
}

/**
 * 現在のブロックを終了する。
 *
 * 次の処理を行なう。
 * - ブロックで管理している stem 型の値の参照カウンタをデクリメント
 * - 変数にアサインされていない stem 型の値を開放
 */
static void lune_leave_blockSub( lune_env_t * _pEnv, lune_block_t * pBlock )
{
    int index;
    for ( index = pBlock->len - 1; index >= 0; index-- ) {
        lune_var_t * pVar = pBlock->pVarList[ index ];
        pVar->refCount--;
        if ( pVar->refCount == 0 ) {
            lune_stem_t * pStem = pVar->pStem;
            if ( pStem != NULL ) {
                lune_decre_ref( _pEnv, pStem );
            }
        }
        lune_free( _pEnv->allocateor, pVar );
        s_globalEnv.allocNum--;
    }

    lune_reset_blockSub( _pEnv, pBlock );
    
    _pEnv->useStemPoolNum -= pBlock->len;
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

    lune_leave_blockSub( _pEnv, pBlock );
}

/**
 * 現在のブロックをクリアする。
 */
void lune_reset_block( lune_env_t * _pEnv )
{
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];

    int index;
    for ( index = pBlock->len - 1; index >= 0; index-- ) {
        lune_var_t * pVar = pBlock->pVarList[ index ];
        if ( pVar->refCount == 1 ) {
            lune_stem_t * pStem = pVar->pStem;
            if ( pStem != NULL ) {
                lune_decre_ref( _pEnv, pStem );
            }
        }
        else {
            pVar->refCount--;
        }
    }

    lune_reset_blockSub( _pEnv, pBlock );

    pBlock->managedStemTop.pPrev = &pBlock->managedStemTop;
    pBlock->managedStemTop.pNext = &pBlock->managedStemTop;
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
        lune_setq( _pEnv, &pBlock->pVarList[ index ]->pStem, pStem );
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
    pDDD->pStemList = (lune_stem_t **)lune_malloc(
        _pEnv->allocateor, sizeof( lune_stem_t * ) * argNum );
    s_globalEnv.allocNum++;

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
lune_stem_t * _lune_createMRet(
    const char * pFile, int lineNo, lune_env_t * _pEnv, bool hasDDD, int num, ... )
{
    lune_stem_t * pDDDStem =
        lune_alloc_stem( _pEnv, lune_value_type_mRet, pFile, lineNo );

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
lune_stem_t * _lune_createDDD(
    const char * pFile, int lineNo, lune_env_t * _pEnv, bool hasDDD, int num, ... )
{
    lune_stem_t * pDDDStem =
        lune_alloc_stem( _pEnv, lune_value_type_ddd, pFile, lineNo );

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
lune_stem_t * _lune_createDDDOnly(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int num ) {
    lune_stem_t * pDDDStem =
        lune_alloc_stem( _pEnv, lune_value_type_ddd, pFile, lineNo );
    lune_ddd_t * pDDD = &pDDDStem->val.ddd;
    pDDD->len = num;
    pDDD->pStemList = (lune_stem_t **)lune_malloc(
        _pEnv->allocateor, sizeof( lune_stem_t * ) * num );
    s_globalEnv.allocNum++;

    return pDDDStem;
}


lune_stem_t * _lune_it_new(
    const char * pFile, int lineNo, 
    lune_env_t * _pEnv, lune_value_type_t type, void * pVal )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, type, pFile, lineNo );
    switch ( type ) {
    case lune_value_type_itList:
        pStem->val.itList = pVal;
        break;
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
    lune_gc_stem( _pEnv, pStem, true );
}

/**
 * クラスのインスタンスを保持する stem を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return stem
 */
lune_stem_t * _lune_class_new(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int size )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, lune_value_type_class, pFile, lineNo );
    void * pObj = lune_malloc( _pEnv->allocateor, size );
    s_globalEnv.allocNum++;
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
    s_globalEnv.allocNum--;
    lune_free( _pEnv->allocateor, pObj );
}

/**
 * 関数 pFunc を保持する stem を生成する
 *
 * @param pFunc 関数
 * @param num フォーム内でアクセスする外部変数の数
 * @param ... フォーム内でアクセスする外部変数
 * @return stem
 */
lune_stem_t * _lune_func2stem(
    const char * pFile, int lineNo, 
    lune_env_t * _pEnv, lune_func_t * pFunc, int argNum, bool hasDDD, int num, ... )
{
    lune_stem_t * pFormStem =
        lune_alloc_stem( _pEnv, lune_value_type_form, pFile, lineNo );
    lune_form_t * pForm = &pFormStem->val.form;
    pForm->len = num;
    pForm->ppClosureValList = (lune_var_t **)lune_malloc(
        _pEnv->allocateor, sizeof( lune_var_t * ) * num );
    s_globalEnv.allocNum++;
    /* pForm->pOrgClosureValList = (lune_stem_t **)lune_malloc( */
    /*     _pEnv->allocateor, sizeof( lune_stem_t * ) * num ); */
    /* s_globalEnv.allocNum++; */

    pForm->pFunc = pFunc;
    pForm->argNum = argNum;
    pForm->hasDDD = hasDDD;
    
    va_list ap;
    va_start( ap, num );

    int index;
    for ( index = 0; index < num; index++ ) {
        lune_var_t * pVar = va_arg( ap, lune_var_t * );
        pVar->refCount++;
        pForm->ppClosureValList[ index ] = pVar;
        // pVar->pStem->refCount++;
        //pForm->pOrgClosureValList[ index ] = pVar->pStem;
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
lune_stem_t * _lune_bool2stem(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_bool_t val )
{
    if ( val ) {
        return _pEnv->pTrueStem;
    }
    return _pEnv->pFalseStem;
}

/**
 * bool 値 val を保持する stem を生成する
 *
 * @param val bool 値
 * @return stem
 */
static lune_stem_t * lune_createBoolStem( lune_env_t * _pEnv, lune_bool_t val )
{
    lune_stem_t * pStem = _lune_alloc_stem( _pEnv, lune_value_type_bool, LUNE_DEBUG_POS );
    pStem->val.boolVal = val;
    return pStem;
}


/**
 * int 値 val を保持する stem を生成する
 *
 * @param val int 値
 * @return stem
 */
lune_stem_t * _lune_int2stem(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_int_t val )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, lune_value_type_int, pFile, lineNo );
    pStem->val.intVal = val;
    return pStem;
}

/**
 * real 値 val を保持する stem を生成する
 *
 * @param val real 値
 * @return stem
 */
lune_stem_t * _lune_real2stem(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_real_t val )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, lune_value_type_real, pFile, lineNo );
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
lune_stem_t * _lune_str2stem(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_str_t val )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, lune_value_type_str, pFile, lineNo );
    pStem->val.str = val;
    return pStem;
}

lune_stem_t * _lune_litStr2stem(
    const char * pFile, int lineNo, lune_env_t * _pEnv, const char * pStr )
{
    lune_stem_t * pStem =
        lune_alloc_stem( _pEnv, lune_value_type_str, pFile, lineNo );
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
static lune_env_t * lune_createEnv()
{
    lune_allocator_t allocateor = lune_createAllocator();
    
    lune_env_t * _pEnv = (lune_env_t *)lune_malloc( allocateor, sizeof( lune_env_t ) );
    s_globalEnv.allocNum++;
    
    _pEnv->allocateor = allocateor;
    _pEnv->useStemPoolNum = 0;
    _pEnv->blockDepth = 0;
    _pEnv->stackPos = 0;

    lune_enter_block( _pEnv, 0 );
    _pEnv->pNoneStem = lune_alloc_stem_op( _pEnv, lune_value_type_none );
    _pEnv->pNilStem = lune_alloc_stem_op( _pEnv, lune_value_type_nil );
    _pEnv->pTrueStem = lune_createBoolStem( _pEnv, true );
    _pEnv->pTrueStem->refCount++;
    _pEnv->pFalseStem = lune_createBoolStem( _pEnv, false );
    _pEnv->pFalseStem->refCount++;
    _pEnv->pMRet = NULL;

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

    _pEnv->pTrueStem->refCount--;
    _pEnv->pFalseStem->refCount--;
    
    lune_leave_block( _pEnv );
    lune_leave_block( _pEnv );

    s_globalEnv.allocNum--;

    printf( ":debug:useStemPoolNum = %d\n", _pEnv->useStemPoolNum );
    printf( ":debug:blockDepth = %d\n", _pEnv->blockDepth );
    
    lune_free( _pEnv->allocateor, _pEnv );
}

static void lune_createGlobalEnv(void) {
    s_globalEnv.allocNum = 0;
    s_globalEnv.pEnv = lune_createEnv();
    s_globalEnv.moduleNum = 0;
}

static void lune_releaseGlobalEnv(void) {
    int index;
    for ( index = s_globalEnv.moduleNum - 1; index >= 0; index-- ) {
        lune_leave_blockSub( s_globalEnv.pEnv,
                             &s_globalEnv.moduleInitBlockBuf[ index ] );
    }
    lune_deleteEnv( s_globalEnv.pEnv );
    printf( ":debug:allocNum = %d\n", s_globalEnv.allocNum );
    lune_checkMem();
}

lune_stem_t * lune_call_form( lune_env_t * _pEnv, lune_stem_t * _pForm, int num, ... )
{
    lune_stem_t * pRet;
    lune_form_t * pFormInfo = &_pForm->val.form;
    
    if ( pFormInfo->hasDDD ) {
        pRet = NULL;
    }
    else {
        lune_stem_t * pStemArray[ 10 ];

        va_list ap;
        va_start( ap, num );

        int index;
        for ( index = 0; index < num; index++ ) {
            pStemArray[ index ] = va_arg( ap, lune_stem_t * );
        }
        va_end(ap);

        switch ( pFormInfo->argNum ) {
        case 0:
            pRet = pFormInfo->pFunc( _pEnv, _pForm );
            break;
        case 1:
            pRet = pFormInfo->pFunc(
                _pEnv, _pForm, pStemArray[ 0 ] );
            break;
        case 2:
            pRet = pFormInfo->pFunc(
                _pEnv, _pForm, pStemArray[ 0 ], pStemArray[ 1 ] );
            break;
        }
    }

    return pRet;
}

lune_stem_t * lune_op_not( lune_env_t * _pEnv, lune_stem_t * pStem ) {
    if ( lune_isCondTrue( pStem ) ) {
        return _pEnv->pFalseStem;
    }
    return _pEnv->pTrueStem;
}

/**
 * pStem の条件判定を行なう
 *
 * false か nil の場合 false、それ以外は true
 */
bool lune_isCondTrue( const lune_stem_t * pStem ) {
    if ( pStem->type == lune_value_type_bool ) {
        return pStem->val.boolVal;
    }
    else if ( pStem->type == lune_value_type_nil ) {
        return false;
    }
    return true;
}

/**
 * スタックを一段上げる
 */
bool lune_incStack( lune_env_t * _pEnv ) {
    _pEnv->stackPos++;
    return false;
}

/**
 * 値 pVal をスタックの top にセットし、値 pVal の条件判定結果を返す
 *
 * スタックに lune_ddd_t は詰めない。
 * 呼び出し側で lune_ddd_t の先頭要素を指定すること。
 *
 * @param pVal スタックに詰む値
 * @return pVal の条件判定結果。 lune_isCondTrue()。
 */
bool lune_setStackVal( lune_env_t * _pEnv, lune_stem_t * pVal )
{
    _pEnv->pValStack[ _pEnv->stackPos ] = pVal;
    return lune_isCondTrue( pVal );
}

/**
 * スタックから値を pop する。
 *
 * @return pop した値。
 */
lune_stem_t * lune_popVal( lune_env_t * _pEnv, bool dummy ) {
    lune_stem_t * pVal = _pEnv->pValStack[ _pEnv->stackPos ];
    _pEnv->stackPos--;
    return pVal;
}

/**
 * int! の unwrap 処理
 */
lune_stem_t * lune_unwrap_stem( lune_stem_t * pStem, lune_stem_t * pDefVal )
{
    if ( pStem->type != lune_value_type_nil ) {
        return pStem;
    }
    if ( pDefVal != NULL ) {
        return pDefVal;
    }
    lune_abort( __func__ );
    return NULL; // dummy
}



/**
 * int! の unwrap 処理
 */
lune_int_t lune_unwrap_int( lune_stem_t * pStem )
{
    if ( pStem->type == lune_value_type_int ) {
        return pStem->val.intVal;
    }
    lune_abort( __func__ );
    return 0; // dummy
}

/**
 * int! の unwrap 処理。 default 値付き。
 */
lune_int_t lune_unwrap_intDefault( lune_stem_t * pStem, lune_int_t val )
{
    if ( pStem->type == lune_value_type_int ) {
        return pStem->val.intVal;
    }
    return val;
}


/**
 * real! の unwrap 処理
 */
lune_real_t lune_unwrap_real( lune_stem_t * pStem )
{
    if ( pStem->type == lune_value_type_real ) {
        return pStem->val.realVal;
    }
    lune_abort( __func__ );
}

/**
 * real! の unwrap 処理。 default 値付き。
 */
lune_real_t lune_unwrap_realDefault( lune_stem_t * pStem, lune_real_t val )
{
    if ( pStem->type == lune_value_type_real ) {
        return pStem->val.realVal;
    }
    return val;
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
            {
                double work;
                if ( modf( pStem->val.realVal, &work ) == 0 ) {
                    printf( "%g.0", pStem->val.realVal );
                }
                else {
                    printf( "%.14g", pStem->val.realVal );
                }
            }
            break;
        case lune_value_type_str:
            printf( "%s", pStem->val.str.pStr );
            break;
        case lune_value_type_form:
            printf( "form: %p", pStem );
            break;
        case lune_value_type_class:
            printf( "class: %p(%s)", pStem, pStem->val.classVal->pMeta->pName );
            break;
        case lune_value_type_if:
            printf( "if: %p(%s): %p(%s)",
                    pStem, pStem->val.ifVal.pMeta->pName,
                    pStem->val.ifVal.pObj,
                    pStem->val.ifVal.pObj->val.classVal->pMeta->pName );
            break;
        default:
            printf( "unknown type -- %d", pStem->type );
            break;
        }
    }
    printf( "\n" );

    lune_leave_block( _pEnv );
}



static int lua_main( lua_State *L) {
    int argc = lua_tointeger(L, 1);
    char ** pArgv = (char **)lua_touserdata(L, 2);
    int script;
    
    lune_createGlobalEnv();
    
    lune_env_t * _pEnv = lune_createEnv();
    
    lune_init_test( _pEnv );

    lune_deleteEnv( _pEnv );

    lune_releaseGlobalEnv();
  

    lua_pushboolean(L, 1);
    return 1;
}

int main (int argc, char * pArgv[] )  {
    lua_State * pLua = luaL_newstate();
    if ( pLua == NULL ) {
        printf( "failed to create a Lua VM.\n" );
        return 1;
    }
    lua_pushcfunction( pLua, lua_main );
    lua_pushinteger( pLua, argc );
    lua_pushlightuserdata( pLua, pArgv );
    int status = lua_pcall( pLua, 2, 1, 0 );
    int result = lua_toboolean( pLua, -1 );
    lua_close( pLua );
    if ( !status ) {
        if ( result ) {
            return 0;
        }
    }
    return 1;
}

