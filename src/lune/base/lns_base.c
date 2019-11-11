/*
  MIT License

  Copyright (c) 2018,2019 ifritJP

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.
*/

#include <lns_base.h>
#include <lauxlib.h>
#include <lunescript.h>
#include <math.h>
#include <stdint.h>


typedef struct lune_globalEnv_t {
    lune_env_t * pEnv;
    /** モジュールの init ブロックのバッファ */
    lune_block_t moduleInitBlockBuf[ LUNE_MODULE_MAX_NUM ];
    /** 現在のモジュール数 */
    int moduleNum;
    /** 現在確保している any の数 */
    int allocNum;
} lune_globalEnv_t;


/**
 * リストの末尾に ANY を追加。
 */
#define lune_add2list( TOP, ANY )               \
    ANY->pNext = TOP;                           \
    ANY->pPrev = (TOP)->pPrev;                  \
    (TOP)->pPrev->pNext = ANY;                  \
    (TOP)->pPrev = ANY;

/**
 * リストから ANY を除外。
 */
#define lune_rmFromList( ANY )                  \
    if ( (ANY)->pNext != NULL ) {               \
        (ANY)->pPrev->pNext = (ANY)->pNext;     \
        (ANY)->pNext->pPrev = (ANY)->pPrev;     \
        (ANY)->pNext = NULL;                    \
    }

#define DECRE_STEM( ENV, STEM )                          \
    if ( STEM.type == lune_stem_type_any ) {             \
        if ( STEM.val.pAny != NULL ) {                   \
            lune_decre_ref( ENV, STEM.val.pAny );        \
        }                                                \
    }


#define lune_malloc( ALLOCATOR, SIZE ) _lune_malloc( ALLOCATOR, SIZE, __FILE__, __LINE__ )
#define lune_free( ALLOCATOR, ADDR ) _lune_free( ALLOCATOR, ADDR, __FILE__, __LINE__ )

static lune_globalEnv_t s_globalEnv;
lune_global_t lune_global;

#define lune_alloc_any( ENV, TYPE, FILE, LINE ) \
    _lune_alloc_any( ENV, TYPE, FILE, LINE )
#define lune_alloc_any_op( ENV, TYPE )                  \
    _lune_alloc_any( ENV, TYPE, __FILE__, __LINE__ )

#define LUNE_FORM_MAX_ARG 20

static void lune_class_del( lune_env_t * _pEnv, void * pObj );
static void lune_alge_del( lune_env_t * _pEnv, void * pObj );


void _lune_abort( const char * pMessage, const char * pFile, int lineNo )
{
    fprintf( stderr, "abort:%s:%d:%s\n", pFile, lineNo, pMessage );
    abort();
}


/**
 * type の値を保持する any 型を生成する。
 *
 * 実際の type の値は別途セットする必要がある。
 *
 * この関数で確保した値は、 _pEnv で管理するブロックに紐付けられる。
 * この時点では、参照カウントは 0。
 * lune_setq() で、何らかの変数に格納しない限り、ブロック終了時に開放される。
 *
 * @return 生成した any 値
 */
static lune_any_t * _lune_alloc_any(
    lune_env_t * _pEnv, lune_value_type_t type, const char * pFile, int lineNo )
{
    s_globalEnv.allocNum++;
    
    lune_any_t * pAny = (lune_any_t *)_lune_malloc(
        _pEnv->allocateor, sizeof( lune_any_t ), pFile, lineNo );
    lune_init_any( pAny, type, _pEnv );
    pAny->refCount = 1;

    // 現在のブロックに登録
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lune_add2list( &pBlock->managedAnyTop, pAny );

    return pAny;
}


/**
 * pAny を開放する。
 *
 * 参照カウントが 0 になっていることが前提条件。
 */
static void lune_release_any( lune_env_t * _pEnv, lune_any_t * pAny ) {
    if ( pAny->type == lune_value_type_str ) {
        pAny->type = lune_value_type_str;
    }
    pAny->type = lune_value_type_none; // このセットは必須ではないが、
    // セットしておいた方が不具合を見つけ易い。
    lune_free( _pEnv->allocateor, pAny );
    s_globalEnv.allocNum--;
}

lune_var_t * lune_var_alloc(
    lune_env_t * _pEnv, lune_block_t * pBlock, int index, lune_stem_type_t type )
{
    lune_var_t * pVar =
        (lune_var_t *)lune_malloc( _pEnv->allocateor, sizeof( lune_var_t ) );
    s_globalEnv.allocNum++;
    
    pVar->stem.type = type;
    pBlock->pVarList[ index ] = pVar;
    pVar->refCount = 1;
    return pVar;
}

static void lune_var_decre( lune_env_t * _pEnv, lune_var_t * pVar )
{
    lune_lock(
        pVar->refCount--;
        if ( pVar->refCount == 0 ) {
            DECRE_STEM( _pEnv, pVar->stem );
            lune_free( _pEnv->allocateor, pVar );
            s_globalEnv.allocNum--;
        }
    );
}

static void lune_gc_any( lune_env_t * _pEnv, lune_any_t * pAny, bool freeFlag ) {
    switch ( pAny->type ) {
    case lune_value_type_str:
        if ( !pAny->val.str.staticFlag ) {
            lune_free( _pEnv->allocateor, (char *)pAny->val.str.pStr );
            s_globalEnv.allocNum--;
        }
        break;
    case lune_value_type_class:
        if ( ((lune_Class_t*)pAny->val.classVal)->pMtd->_gc != NULL ) {
            ((lune_Class_t*)pAny->val.classVal)->pMtd->_gc( _pEnv, pAny );
        }
        if ( ((lune_Class_t*)pAny->val.classVal)->pMtd->_del != NULL ) {
            ((lune_Class_t*)pAny->val.classVal)->pMtd->_del( _pEnv, pAny );
        }
        lune_class_del( _pEnv, pAny->val.classVal );
        break;
    case lune_value_type_alge:
        if ( pAny->val.alge.gc != NULL ) {
            pAny->val.alge.gc( _pEnv, pAny->val.alge.pVal );
        }
        lune_alge_del( _pEnv, pAny->val.alge.pVal );
        break;
    case lune_value_type_ddd: // fall-through
    case lune_value_type_mRet:
        if ( pAny->val.ddd.len != 0 ) {
            lune_free( _pEnv->allocateor, pAny->val.ddd.stemList );
            s_globalEnv.allocNum--;
        }
        break;
    case lune_value_type_form:
        {
            int index;
            for ( index = 0; index < pAny->val.form.len; index++ ) {
                lune_closureVal_t * pVal = lune_form_closure( pAny, index );
                lune_var_decre( _pEnv, pVal->pVar );
            }
            lune_free( _pEnv->allocateor, pAny->val.form.pClosureValList );
            s_globalEnv.allocNum--;
        }
        break;
    case lune_value_type_itSet:
        lune_itSet__del( _pEnv, pAny );
        break;
    case lune_value_type_if:
        lune_decre_ref( _pEnv, pAny->val.ifVal.pObj );
        // if は class 内のメンバなので開放しないで return する。
        return;
    default:
        break;
    }

    if ( freeFlag ) {
        lune_release_any( _pEnv, pAny );
    }
}

/**
 * pAny の参照カウントをデクリメントする。
 *
 * 参照カウントが 0 になった場合は、開放処理を行なう。
 */
bool lune_decre_ref( lune_env_t * _pEnv, lune_any_t * pAny ) {
    lune_lock(
        pAny->refCount--;
        if ( pAny->refCount == 0 ) {
            lune_gc_any( _pEnv, pAny, true );
            return true;
        }
    );
    return false;
}


/**
MRet の戻り値を設定し、先頭の値を返す。
 */
lune_stem_t lune_setMRet( lune_env_t * _pEnv, lune_any_t * pAny ) {
    _pEnv->pMRet = pAny;
    return lune_fromDDD( pAny, 0 );
}


lune_any_t * lune_getIF( lune_env_t * _pEnv, lune_any_t * pIFAny )
{
    lune_lock(
        if ( pIFAny->refCount == 0 ) {
            // interface の any は共有なので、
            // 参照回数が 0 の状態から 1 になったタイミングで、
            // 元インスタンスの参照回数をインクリメントする。
            pIFAny->val.ifVal.pObj->refCount++;

            lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
            lune_add2list( &pBlock->managedAnyTop, pIFAny );
            pIFAny->refCount++;
        }
    );
    return pIFAny;
}

/**
 * pAny が class ならば、指定の pMeta の IF のインスタンスを返す。
 * pAny が IF なら、そのまま pAny を返す。
 */
lune_stem_t lune_toIF(
    lune_env_t * _pEnv, lune_any_t * pAny, const lune_type_meta_t * pMeta )
{
    if ( pAny->type == lune_value_type_if ) {
        return LUNE_STEM_ANY( pAny );
    }
    if ( pAny->type == lune_value_type_class ) {
        lune_any_t * pIFAny = (lune_any_t *)pAny->val.classVal->pIFdummy;
        for ( ; pIFAny->type != lune_value_type_none; pIFAny++ ) {
            if ( pIFAny->val.ifVal.pMeta == pMeta ) {
                return LUNE_STEM_ANY( lune_getIF( _pEnv, pIFAny ) );
            }
        }
    }
    lune_abort( "not found interface" );
    return lune_global.nilStem;
}

void lune_setQ_( lune_any_t * pAny )
{
    if ( pAny != NULL ) {
        lune_lock( 
            pAny->refCount++;
        );
    }
}

void lune_setRet( lune_env_t * _pEnv, lune_stem_t stem )
{
    if ( stem.type != lune_stem_type_any ) {
        return;
    }

    lune_any_t * pAny = stem.val.pAny;
    
    lune_rmFromList( pAny );

    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth - 1 ];
    lune_add2list( &pBlock->managedAnyTop, pAny );

    if ( pAny->type == lune_value_type_ddd ||
         pAny->type == lune_value_type_mRet )
    {
        int index;
        for ( index = 0; index < pAny->val.ddd.len; index++ ) {
            lune_stem_t item = lune_fromDDD( pAny, index );
            lune_setRet( _pEnv, item );
        }
    }
    return;
}




void lune_setup_block(
    lune_env_t * _pEnv, lune_block_t * pBlock, int anyNum, int stemNum, int varNum )
{
    int dummy;

    pBlock->managedAnyTop.pPrev = &pBlock->managedAnyTop;
    pBlock->managedAnyTop.pNext = &pBlock->managedAnyTop;

    pBlock->pVarList = _pEnv->varPPool + _pEnv->useVarPoolNum;
    pBlock->varLen = varNum;
    int index;
    for ( index = 0; index < pBlock->varLen; index++ ) {
        pBlock->pVarList[ index ] = NULL;
    }
    _pEnv->useVarPoolNum += varNum;
    
    pBlock->stemLen = stemNum;
    pBlock->pStemList = _pEnv->stemPPool + _pEnv->useStemPoolNum;
    for ( index = 0; index < stemNum; index++ ) {
        pBlock->pStemList[ index ] = NULL;
    }
    _pEnv->useStemPoolNum += stemNum;

    pBlock->anyLen = anyNum;
    pBlock->pAnyList = _pEnv->anyPPool + _pEnv->useAnyPoolNum;
    for ( index = 0; index < anyNum; index++ ) {
        pBlock->pAnyList[ index ] = NULL;
    }
    _pEnv->useAnyPoolNum += anyNum;

    
    pBlock->blockDepth = _pEnv->blockDepth;
    pBlock->pStackAddr = &dummy;
    
}


/**
 * 新しくブロックを開始する。
 *
 * @param anyVerNum ブロックで管理する any 型の値の数
 * @return ブロック情報
 */
lune_block_t * lune_enter_module( int anyNum, int stemNum, int varNum )
{
    lune_block_t * pBlock = &s_globalEnv.moduleInitBlockBuf[ s_globalEnv.moduleNum ];
    s_globalEnv.moduleNum++;

    lune_setup_block( s_globalEnv.pEnv, pBlock, anyNum, stemNum, varNum );

    return pBlock;
}


/**
 * 新しくブロックを開始する。
 *
 * @param anyVerNum ブロックで管理する any 型の値の数
 * @return ブロック情報
 */
lune_block_t * lune_enter_block(
    lune_env_t * _pEnv, int anyNum, int stemNum, int varNum )
{
    _pEnv->blockDepth++;
    
    lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lune_setup_block( _pEnv, pBlock, anyNum, stemNum, varNum );

    return pBlock;
}


static inline void lune_reset_blockSub( lune_env_t * _pEnv, lune_block_t * pBlock ) {
    lune_any_t * pWork = pBlock->managedAnyTop.pPrev;
    while ( pWork != &pBlock->managedAnyTop ) {
        lune_any_t * pPrev = pWork->pPrev;

        lune_lock( 
            if ( pWork->refCount == 1 ) {
                lune_gc_any( _pEnv, pWork, true );
            }
            else {
                pWork->refCount--;
            }
        );
        pWork = pPrev;
    }
}

/**
 * 現在のブロックを終了する。
 *
 * 次の処理を行なう。
 * - ブロックで管理している any 型の値の参照カウンタをデクリメント
 * - 変数にアサインされていない any 型の値を開放
 */
static void lune_leave_blockSub( lune_env_t * _pEnv, lune_block_t * pBlock )
{
    int index;
    for ( index = pBlock->varLen - 1; index >= 0; index-- ) {
        lune_var_t * pVar = pBlock->pVarList[ index ];
        if ( pVar != NULL ) {
            lune_var_decre( _pEnv, pVar );
        }
    }
    for ( index = pBlock->stemLen - 1; index >= 0; index-- ) {
        lune_stem_t * pStem = pBlock->pStemList[ index ];
        if ( pStem != NULL && pStem->type == lune_stem_type_any ) {
            if ( lune_decre_ref( _pEnv, pStem->val.pAny ) ) {
                pStem->type = lune_stem_type_none;
            }
        }
    }
    for ( index = pBlock->anyLen - 1; index >= 0; index-- ) {
        lune_any_t * pAny = pBlock->pAnyList[ index ];
        if ( pAny != NULL ) {
            lune_decre_ref( _pEnv, pAny );
        }
    }

    lune_reset_blockSub( _pEnv, pBlock );
    
    _pEnv->useAnyPoolNum -= pBlock->anyLen;
    _pEnv->useStemPoolNum -= pBlock->stemLen;
    _pEnv->useVarPoolNum -= pBlock->varLen;
}


/**
 * 現在のブロックを終了する。
 *
 * 次の処理を行なう。
 * - ブロックで管理している any 型の値の参照カウンタをデクリメント
 * - 変数にアサインされていない any 型の値を開放
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

    for ( index = pBlock->varLen - 1; index >= 0; index-- ) {
        lune_var_t * pVar = pBlock->pVarList[ index ];
        if ( pVar != NULL ) {
            lune_var_decre( _pEnv, pVar );
        }
    }
    
    for ( index = pBlock->stemLen - 1; index >= 0; index-- ) {
        lune_stem_t * pStem = pBlock->pStemList[ index ];
        if ( pStem != NULL && pStem->type == lune_stem_type_any ) {
            lune_decre_ref( _pEnv, pStem->val.pAny );
        }
        
    }

    for ( index = pBlock->anyLen - 1; index >= 0; index-- ) {
        lune_any_t * pAny = pBlock->pAnyList[ index ];
        if ( pAny != NULL ) {
            lune_decre_ref( _pEnv, pAny );
        }
    }

    lune_reset_blockSub( _pEnv, pBlock );

    pBlock->managedAnyTop.pPrev = &pBlock->managedAnyTop;
    pBlock->managedAnyTop.pNext = &pBlock->managedAnyTop;
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
 * @param num lune_enter_block() の anyVerNum と同じ。
 *   ブロック内の any 値の数と、 argNum を合せた数が num になる。
 * @param argNum 引数の any 型の値の数
 * @param ... 引数の any 型の値。
 */
/* lune_block_t * lune_enter_func( */
/*     lune_env_t * _pEnv, int stemNum, int varNum, int argNum, ... ) */
/* { */
/*     lune_enter_block( _pEnv, stemNum, varNum ); */
/*     lune_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ]; */


/*     va_list ap; */
/*     va_start( ap, argNum ); */

/*     int index; */
/*     for ( index = 0; index < argNum; index++ ) { */
/*         lune_any_t * pAny = va_arg( ap, lune_any_t * ); */
/*         lune_setq( _pEnv, &pBlock->pVarList[ index ]->stem.val.pAny, pAny ); */
/*         pBlock->pVarList[ index ]->stem.type = lune_stem_type_any; */
/*     } */
/*     va_end(ap); */

/*     return pBlock; */
/* } */

lune_stem_t lune_getValFromDDD( lune_any_t * pAny, int index )
{
    if ( pAny->val.ddd.len <= index ) {
        return lune_global.nilStem;
    }
    return pAny->val.ddd.stemList[ index ];
}


static lune_any_t * lune_createDDDSub(
    lune_env_t * _pEnv, bool hasDDD, int num, va_list apSrc, lune_any_t * pDDDAny ) {
    lune_ddd_t * pDDD = &pDDDAny->val.ddd;
    int argNum = num;

    if ( hasDDD ) {
        // 最終要素が複数要素の場合、その複数要素が何個あるのか調べる
        argNum = num;

        va_list ap;
        va_copy( ap, apSrc );

        int index;
        for ( index = 0; index < num - 1; index++ ) {
            va_arg( ap, lune_stem_t );
        }
        lune_stem_t dddStem = va_arg( ap, lune_stem_t );
        argNum += dddStem.val.pAny->val.ddd.len - 1;
        va_end(ap);
    }
    
    pDDD->len = argNum;
    pDDD->stemList = (lune_stem_t *)lune_malloc(
        _pEnv->allocateor, sizeof( lune_stem_t ) * argNum );
    s_globalEnv.allocNum++;

    va_list ap;
    va_copy( ap, apSrc );

    int index;
    for ( index = 0; index < num - 1; index++ ) {
        lune_stem_t stem = va_arg( ap, lune_stem_t );
        if ( stem.type == lune_stem_type_any ) {
            lune_any_t * pAny = stem.val.pAny;
            if ( pAny->type == lune_value_type_ddd ||
                 pAny->type == lune_value_type_mRet )
            {
                pDDD->stemList[ index ] = lune_getValFromDDD( pAny, 0 );
            }
            else {
                pDDD->stemList[ index ] = stem;
            }
        }
        else {
            pDDD->stemList[ index ] = stem;
        }
    }

    if ( hasDDD ) {
        lune_stem_t stem = va_arg( ap, lune_stem_t );
        lune_any_t * pAny = stem.val.pAny;
        lune_stem_t* pStem = &pDDD->stemList[ num - 1 ];
        for ( index = 0; index < lune_lenDDD( pAny ); index++ ) {
            *pStem = lune_fromDDD( pAny, index );
            pStem++;
        }
    }
    else {
        pDDD->stemList[ num - 1 ] = va_arg( ap, lune_stem_t );
    }

    
    va_end(ap);

    return pDDDAny;
}

/**
 * ... の値を生成する
 *
 * ... に含める値は全て any に変換する必要がある。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t _lune_createMRet(
    const char * pFile, int lineNo, lune_env_t * _pEnv, bool hasDDD, int num, ... )
{
    lune_any_t * pDDDAny =
        lune_alloc_any( _pEnv, lune_value_type_mRet, pFile, lineNo );

    va_list ap;
    va_start( ap, num );

    lune_createDDDSub( _pEnv, hasDDD, num, ap, pDDDAny );

    va_end(ap);

    return LUNE_STEM_ANY( pDDDAny );
}


/**
 * ... の値を生成する
 *
 * ... に含める値は全て any に変換する必要がある。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t _lune_createDDD(
    const char * pFile, int lineNo, lune_env_t * _pEnv, bool hasDDD, int num, ... )
{
    if ( num > 0 ) {
        lune_any_t * pDDDAny =
            lune_alloc_any( _pEnv, lune_value_type_ddd, pFile, lineNo );

        va_list ap;
        va_start( ap, num );
        
        lune_createDDDSub( _pEnv, hasDDD, num, ap, pDDDAny );

        va_end(ap);

        return LUNE_STEM_ANY( pDDDAny );
    }
    else {
        return lune_createDDDOnly( _pEnv, 0 );
    }

}


/**
 * ... の値を生成する
 *
 * ... に含める値は全て any に変換する必要がある。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t _lune_createSubDDD(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int offset, lune_stem_t ddd )
{
    lune_any_t * pDDD = ddd.val.pAny;
    
    int len;
    len = lune_lenDDD( pDDD ) - offset;
    if ( len < 0 ) {
        return _lune_createDDDOnly( pFile, lineNo, _pEnv, 0 );
    }

    lune_stem_t pSubDDD = _lune_createDDDOnly( pFile, lineNo, _pEnv, len );
        
    int index;
    for ( index = 0; index < len; index++ ) {
        lune_set2DDDArg( pSubDDD.val.pAny, index, lune_fromDDD( pDDD, index + offset ) );
    }
    return pSubDDD;
}




/**
 * ... の値を生成する
 *
 * ... に含める値は全て any に変換する必要がある。
 *
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lune_stem_t _lune_createDDDOnly(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int num ) {
    lune_any_t * pDDDAny =
        lune_alloc_any( _pEnv, lune_value_type_ddd, pFile, lineNo );
    lune_ddd_t * pDDD = &pDDDAny->val.ddd;
    pDDD->len = num;
    if ( num != 0 ) {
        pDDD->stemList = (lune_stem_t *)lune_malloc(
            _pEnv->allocateor, sizeof( lune_stem_t ) * num );
        s_globalEnv.allocNum++;
    }
    else {
        pDDD->stemList = NULL;
    }

    return LUNE_STEM_ANY( pDDDAny );
}


lune_any_t * _lune_luaVal_new(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_value_type_t type )
{
    lune_any_t * pAny = lune_alloc_any( _pEnv, lune_value_type_luaVal, pFile, lineNo );
    pAny->val.luaVal.type = type;
    return pAny;
}


lune_any_t * _lune_it_new(
    const char * pFile, int lineNo, 
    lune_env_t * _pEnv, lune_value_type_t type, void * pVal )
{
    lune_any_t * pAny =
        lune_alloc_any( _pEnv, type, pFile, lineNo );
    switch ( type ) {
    case lune_value_type_itList:
        pAny->val.itList = pVal;
        break;
    case lune_value_type_itSet:
        pAny->val.itSet = pVal;
        break;
    case lune_value_type_itMap:
        pAny->val.itMap = pVal;
        break;
    default:
        break;
    }
    return pAny;
}

void lune_it_delete( lune_env_t * _pEnv, lune_any_t * pAny )
{
    lune_rmFromList( pAny );
    lune_gc_any( _pEnv, pAny, true );
}


/**
 * クラスのインスタンスを保持する any を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return any
 */
lune_any_t * _lune_alge_new(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int valType,
    int size, lune_algeVal_gc_t * gc )
{
    lune_any_t * pAny =
        lune_alloc_any( _pEnv, lune_value_type_alge, pFile, lineNo );
    void * pObj = lune_malloc( _pEnv->allocateor, size );
    s_globalEnv.allocNum++;
    pAny->val.alge.type = valType;
    pAny->val.alge.pVal = pObj;
    pAny->val.alge.gc = gc;
    return pAny;
}

/**
 * クラスのインスタンスを開放する
 *
 * @param pObj クラスのインスタンス
 */
static void lune_alge_del( lune_env_t * _pEnv, void * pObj )
{
    s_globalEnv.allocNum--;
    lune_free( _pEnv->allocateor, pObj );
}

/**
 * クラスのインスタンスを保持する any を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return any
 */
lune_any_t * _lune_class_new(
    const char * pFile, int lineNo, lune_env_t * _pEnv, int size )
{
    lune_any_t * pAny =
        lune_alloc_any( _pEnv, lune_value_type_class, pFile, lineNo );
    void * pObj = lune_malloc( _pEnv->allocateor, size );
    s_globalEnv.allocNum++;
    pAny->val.classVal = pObj;
    return pAny;
}

/**
 * クラスのインスタンスを開放する
 *
 * @param pObj クラスのインスタンス
 */
static void lune_class_del( lune_env_t * _pEnv, void * pObj )
{
    s_globalEnv.allocNum--;
    lune_free( _pEnv->allocateor, pObj );
}

/**
 * 関数 pFunc を保持する any を生成する
 *
 * @param pFunc 関数
 * @param num フォーム内でアクセスする外部変数の数
 * @param ... フォーム内でアクセスする外部変数
 * @return any
 */
lune_any_t * _lune_func2any(
    const char * pFile, int lineNo, 
    lune_env_t * _pEnv, lune_closure_t * pFunc, int argNum, bool hasDDD, int num, ... )
{
    lune_any_t * pFormAny =
        lune_alloc_any( _pEnv, lune_value_type_form, pFile, lineNo );
    lune_form_t * pForm = &pFormAny->val.form;
    pForm->len = num;
    if ( num == 0 ) {
        pForm->needFormParam = false;
    }
    else {
        pForm->needFormParam = true;
    }
    pForm->pClosureValList = (lune_closureVal_t*)lune_malloc(
        _pEnv->allocateor, sizeof( lune_closureVal_t ) * num );
    s_globalEnv.allocNum++;
    /* pForm->pOrgClosureValList = (lune_any_t **)lune_malloc( */
    /*     _pEnv->allocateor, sizeof( lune_any_t * ) * num ); */
    /* s_globalEnv.allocNum++; */

    pForm->pFunc = pFunc;
    //pForm->argNum = argNum;
    pForm->hasDDD = hasDDD;
    
    va_list ap;
    va_start( ap, num );

    lune_lock( 
        int index;
        for ( index = 0; index < num; index++ ) {
            lune_closureVal_t * pClosureVal = &pForm->pClosureValList[ index ];
            lune_var_t * pVar = va_arg( ap, lune_var_t * );
            pVar->refCount++;
            pForm->pClosureValList[ index ].pVar = pVar;
        }
    );
    va_end(ap);

    return pFormAny;
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
 * 文字列データを保持する any を生成する
 *
 * @param val 文字列型データ
 * @return any
 */
lune_any_t * _lune_str2any(
    const char * pFile, int lineNo, lune_env_t * _pEnv, lune_str_t val )
{
    lune_any_t * pAny =
        lune_alloc_any( _pEnv, lune_value_type_str, pFile, lineNo );
    pAny->val.str = val;
    return pAny;
}

lune_any_t * _lune_litStr2any(
    const char * pFile, int lineNo, lune_env_t * _pEnv, const char * pStr )
{
    lune_any_t * pAny =
        lune_alloc_any( _pEnv, lune_value_type_str, pFile, lineNo );
    pAny->val.str.len = strlen( pStr );
    pAny->val.str.pStr = pStr;
    pAny->val.str.staticFlag = true;
    return pAny;
}

lune_int_t lune_stem2int( lune_stem_t stem )
{
    switch ( stem.type ) {
    case lune_stem_type_int:
        return stem.val.intVal;
    case lune_stem_type_real:
        return (lune_int_t)stem.val.realVal;
    default:
        break;
    }
    lune_abort( "convert error stem2int" );
    return 0;
}

lune_real_t lune_stem2real( lune_stem_t stem )
{
    switch ( stem.type ) {
    case lune_stem_type_int:
        return stem.val.intVal;
    case lune_stem_type_real:
        return stem.val.realVal;
    default:
        break;
    }
    lune_abort( "convert error stem2real" );
    return 0.0;
}

lune_bool_t lune_stem2bool( lune_stem_t stem )
{
    switch ( stem.type ) {
    case lune_stem_type_bool:
        return stem.val.boolVal;
    default:
        break;
    }
    lune_abort( "convert error stem2real" );
    return 0.0;
}


lune_any_t * _lune_cloneBin2any(
    const char * pFile, int lineNo, lune_env_t * _pEnv, const void * pBuf, int len )
{
    lune_any_t * pAny = lune_alloc_any( _pEnv, lune_value_type_str, pFile, lineNo );

    char * pAlloc = (char *)lune_malloc( _pEnv->allocateor, len + 1 );
    s_globalEnv.allocNum++;
    if ( pAlloc == NULL ) {
        lune_abort( "alloc error" );
    }
    memcpy( pAlloc, pBuf, len );
    pAlloc[ len ] = '\0';
    pAny->val.str.pStr = pAlloc;
    pAny->val.str.len = len;
    pAny->val.str.staticFlag = false;
    return pAny;
}

bool lune_equals_any( const lune_any_t * pAny1, const lune_any_t * pAny2 )
{
    if ( pAny1->type == pAny2->type ) {
        if ( pAny1 == pAny2 ) {
            // 同じ any データなら equal
            return true;
        }
        switch ( pAny1->type ) {
        case lune_value_type_str:
            {
                // 文字列は同じデータなら equal
                const lune_str_t * pStr1 = &pAny1->val.str;
                const lune_str_t * pStr2 = &pAny2->val.str;
                if ( pStr1->len != pStr2->len ) {
                    // 長さが違えば false
                    return false;
                }
                return memcmp( pStr1->pStr, pStr2->pStr, pStr1->len ) == 0;
            }
            break;
        case lune_value_type_if:
            // インタフェースは、同じインスタンスを指している場合は equal
            return pAny1->val.ifVal.pObj == pAny2->val.ifVal.pObj;
        default:
            break;
        }
        return false;
    }
    // 異なる any タイプなら
    switch ( pAny1->type ) {
    case lune_value_type_if:
        if ( pAny2->type == lune_value_type_class ) {
            return pAny1->val.ifVal.pObj == pAny2;
        }
        break;
    case lune_value_type_class:
        if ( pAny2->type == lune_value_type_if ) {
            return pAny2->val.ifVal.pObj == pAny1;
        }
        break;
    default:
        break;
    }
    return false;
}

bool lune_equals( const lune_stem_t stem1, const lune_stem_t stem2 ) {
    if ( stem1.type == stem2.type ) {
        switch ( stem1.type ) {
        case lune_stem_type_nil:
            return true;
        case lune_stem_type_int:
            return stem1.val.intVal == stem2.val.intVal;
        case lune_stem_type_real:
            return stem1.val.realVal == stem2.val.realVal;
        case lune_stem_type_bool:
            return stem1.val.boolVal == stem2.val.boolVal;
        default:
            break;
        }
        return lune_equals_any( stem1.val.pAny, stem2.val.pAny );
    }

    // 異なる stem タイプなら
    switch ( stem1.type ) {
    case lune_stem_type_int:
        if ( stem2.type == lune_stem_type_real ) {
            return stem1.val.intVal == stem2.val.realVal;
        }
        break;
    case lune_stem_type_real:
        if ( stem2.type == lune_stem_type_int ) {
            return stem2.val.intVal == stem1.val.realVal;
        }
        break;
    default:
        break;
    }

    return false;
}



/**
 * 環境を生成する。
 *
 * この環境は、スレッド毎に 1 つ。
 *
 * @return 環境。
 */
static lune_env_t * lune_createEnv( lua_State * pLua )
{
    lune_allocator_t allocateor = lune_createAllocator();
    
    lune_env_t * _pEnv =
        (lune_env_t *)lune_malloc( allocateor, sizeof( lune_env_t ) );
    s_globalEnv.allocNum++;
    
    _pEnv->allocateor = allocateor;
    _pEnv->pLua = pLua;
    _pEnv->useAnyPoolNum = 0;
    _pEnv->useStemPoolNum = 0;
    _pEnv->useVarPoolNum = 0;
    _pEnv->blockDepth = 0;
    _pEnv->stackPos = 0;

    lune_enter_block( _pEnv, 0, 0, 0 );
    _pEnv->pMRet = NULL;

    _pEnv->pSortCallback = lune_global.nilStem;

    lune_enter_block( _pEnv, 0, 0, 0 );

    if ( pLua != NULL ) {
        lune_setLuaWapper( _pEnv );
    }
    
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

    s_globalEnv.allocNum--;

    printf( ":debug:useAnyPoolNum = %d\n", _pEnv->useAnyPoolNum );
    printf( ":debug:useStemPoolNum = %d\n", _pEnv->useStemPoolNum );
    printf( ":debug:useVarPoolNum = %d\n", _pEnv->useVarPoolNum );
    printf( ":debug:blockDepth = %d\n", _pEnv->blockDepth );
    
    lune_free( _pEnv->allocateor, _pEnv );
}

static void lune_createGlobalEnv() {
    s_globalEnv.allocNum = 0;
    s_globalEnv.pEnv = lune_createEnv( NULL );
    s_globalEnv.moduleNum = 0;

    
    lune_global.noneStem = LUNE_STEM_BASE( lune_stem_type_none );
    lune_global.nilStem = LUNE_STEM_BASE( lune_stem_type_nil );
    lune_global.trueStem = LUNE_STEM_BOOL( true );
    lune_global.falseStem = LUNE_STEM_BOOL( false );

    lune_setQ( lune_global.ddd0, lune_createDDD( s_globalEnv.pEnv, false, 0 ) );
}

static void lune_releaseGlobalEnv(void) {
    int index;
    for ( index = s_globalEnv.moduleNum - 1; index >= 0; index-- ) {
        lune_leave_blockSub( s_globalEnv.pEnv,
                             &s_globalEnv.moduleInitBlockBuf[ index ] );
    }
    lune_decre_ref( s_globalEnv.pEnv, lune_global.ddd0.val.pAny );
    lune_deleteEnv( s_globalEnv.pEnv );
    printf( ":debug:allocNum = %d\n", s_globalEnv.allocNum );
    lune_checkMem();
}

void lune_init_alge( lune_stem_t * pStem, lune_any_t * pAny, int valType )
{
    pStem->type = lune_stem_type_any;
    pStem->val.pAny = pAny;
    lune_init_any( pAny, lune_value_type_alge, ENV );
    pAny->refCount = 1;                                      
    pAny->val.alge.type = valType;                           
    pAny->val.alge.pVal = NULL;                              
    pAny->val.alge.gc = NULL;
}

lune_stem_t lune_call_form(
    lune_env_t * _pEnv, lune_any_t * _pForm, lune_stem_t ddd )
{
    if ( lune_isClosure( _pForm ) ) {
        return lune_closure( _pForm )( _pEnv, _pForm, ddd );
    }
    else {
        return lune_func( _pForm )( _pEnv, ddd );
    }
}

bool lune_op_not( lune_env_t * _pEnv, lune_stem_t stem ) {
    return !lune_isCondTrue( stem );
}

/**
 * pAny の条件判定を行なう
 *
 * false か nil の場合 false、それ以外は true
 */
bool lune_isCondTrue( const lune_stem_t stem ) {
    if ( stem.type == lune_stem_type_bool ) {
        return stem.val.boolVal;
    }
    else if ( stem.type == lune_stem_type_nil ) {
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
bool lune_setStackVal( lune_env_t * _pEnv, lune_stem_t val )
{
    _pEnv->pValStack[ _pEnv->stackPos ] = val;
    return lune_isCondTrue( val );
}

/**
 * スタックから値を pop する。
 *
 * @return pop した値。
 */
lune_stem_t lune_popVal( lune_env_t * _pEnv, bool dummy ) {
    lune_stem_t val = _pEnv->pValStack[ _pEnv->stackPos ];
    _pEnv->stackPos--;
    return val;
}

/**
 * int! の unwrap 処理
 */
lune_stem_t lune_unwrap_stem( lune_stem_t stem, lune_stem_t defVal )
{
    if ( stem.type != lune_stem_type_nil ) {
        return stem;
    }
    if ( defVal.type != lune_stem_type_none ) {
        return defVal;
    }
    lune_abort( __func__ );
    return lune_global.noneStem; // dummy
}

lune_any_t * lune_unwrap_any( lune_stem_t stem, lune_stem_t defVal )
{
    if ( stem.type != lune_stem_type_nil ) {
        return stem.val.pAny;
    }
    if ( defVal.type != lune_stem_type_none ) {
        return defVal.val.pAny;
    }
    lune_abort( __func__ );
    return NULL; // dummy
}



/**
 * int! の unwrap 処理
 */
lune_int_t lune_unwrap_int( lune_stem_t stem )
{
    if ( stem.type == lune_stem_type_int ) {
        return stem.val.intVal;
    }
    lune_abort( __func__ );
    return 0; // dummy
}

/**
 * int! の unwrap 処理。 default 値付き。
 */
lune_int_t lune_unwrap_intDefault( lune_stem_t stem, lune_int_t val )
{
    if ( stem.type == lune_stem_type_int ) {
        return stem.val.intVal;
    }
    return val;
}


/**
 * real! の unwrap 処理
 */
lune_real_t lune_unwrap_real( lune_stem_t stem )
{
    if ( stem.type == lune_stem_type_real ) {
        return stem.val.realVal;
    }
    lune_abort( __func__ );
    return 0.0;
}

/**
 * real! の unwrap 処理。 default 値付き。
 */
lune_real_t lune_unwrap_realDefault( lune_stem_t stem, lune_real_t val )
{
    if ( stem.type == lune_stem_type_real ) {
        return stem.val.realVal;
    }
    return val;
}


/**
 * lua の print() に相当する処理。
 *
 * @param pArg は ddd 型。
 */
void lune_print( lune_env_t * _pEnv, lune_stem_t ddd ) {

    //lune_enter_func( _pEnv, 0, 0, 1, pArg );
    //lune_enter_block( _pEnv, 0, 0 );

    lune_any_t * pArg = ddd.val.pAny;

    int index;
    for ( index = 0; index < lune_lenDDD( pArg ); index++ ) {
        if ( index > 0 ) {
            printf( "\t" );
        }
        const lune_stem_t stem = lune_fromDDD( pArg, index );
        switch ( stem.type ) {
        case lune_stem_type_nil:
            printf( "nil" );
            break;
        case lune_stem_type_bool:
            printf( stem.val.boolVal ? "true" : "false" );
            break;
        case lune_stem_type_int:
            printf( "%d", stem.val.intVal );
            break;
        case lune_stem_type_real:
            {
                double work;
                if ( modf( stem.val.realVal, &work ) == 0 ) {
                    printf( "%g.0", stem.val.realVal );
                }
                else {
                    printf( "%.14g", stem.val.realVal );
                }
            }
            break;
            
        case lune_stem_type_any:
            {
                lune_any_t * pAny = stem.val.pAny;
                switch ( pAny->type ) {
                case lune_value_type_str:
                    printf( "%s", pAny->val.str.pStr );
                    break;
                case lune_value_type_form:
                    printf( "form: %p", pAny );
                    break;
                case lune_value_type_class:
                    printf( "class: %p(%s)", pAny, pAny->val.classVal->pMeta->pName );
                    break;
                case lune_value_type_if:
                    printf( "if: %p(%s): %p(%s)",
                            pAny, pAny->val.ifVal.pMeta->pName,
                            pAny->val.ifVal.pObj,
                            pAny->val.ifVal.pObj->val.classVal->pMeta->pName );
                    break;
                case lune_value_type_luaVal:
                    switch( pAny->val.luaVal.type ) {
                    case lune_value_type_none:
                        printf( "lua: none %p", pAny );
                        break;
                    case lune_value_type_str:
                        printf( "lua: str %p", pAny );
                        break;
                    case lune_value_type_class:
                        printf( "lua: class %p", pAny );
                        break;
                    case lune_value_type_if:
                        printf( "lua: if %p", pAny );
                        break;
                    case lune_value_type_ddd:
                        printf( "lua: ddd %p", pAny );
                        break;
                    case lune_value_type_mRet:
                        printf( "lua: mRet %p", pAny );
                        break;
                    case lune_value_type_form:
                        printf( "lua: form %p", pAny );
                        break;
                    case lune_value_type_List:
                        printf( "lua: List %p", pAny );
                        break;
                    case lune_value_type_Array:
                        printf( "lua: Array %p", pAny );
                        break;
                    case lune_value_type_Set:
                        printf( "lua: Set %p", pAny );
                        break;
                    case lune_value_type_Map:
                        printf( "lua: Map %p", pAny );
                        break;
                    case lune_value_type_itList:
                        printf( "lua: itList %p", pAny );
                        break;
                    case lune_value_type_itSet:
                        printf( "lua: itSet %p", pAny );
                        break;
                    case lune_value_type_itMap:
                        printf( "lua: itMap %p", pAny );
                        break;
                    case lune_value_type_alge:
                        printf( "lua: alge %p", pAny );
                        break;
                    case lune_value_type_luaVal:
                        printf( "lua: luaVal %p", pAny );
                        break;
                    default:
                        printf( "unknown type -- %d", pAny->val.luaVal.type );
                        break;
                    }
                    break;
                default:
                    printf( "unknown type -- %d", pAny->type );
                    break;
                }
            }
            break;
        default:
            printf( "unknown stem type -- %d", stem.type );
            break;
        }
    }
    printf( "\n" );

    //lune_leave_block( _pEnv );
}

int lua_main( lua_State * pLua ) {
    int argc = lua_tointeger( pLua, 1);
    char ** pArgv = (char **)lua_touserdata( pLua, 2);
    int script;

    lune_collection_init();
    
    lune_createGlobalEnv();
    
    lune_env_t * _pEnv = lune_createEnv( pLua );

    lune_init_lns_builtin( _pEnv );
    
    lune_run_module( _pEnv );

    lune_deleteEnv( _pEnv );

    lune_releaseGlobalEnv();
  

    lua_pushboolean( pLua, 1);
    return 1;
}
