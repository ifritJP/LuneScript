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


typedef struct lns_globalEnv_t {
    lns_env_t * pEnv;
    /** モジュールの init ブロックのバッファ */
    lns_block_t moduleInitBlockBuf[ LNS_MODULE_MAX_NUM ];
    /** 現在のモジュール数 */
    int moduleNum;
    /** 現在確保している any の数 */
    int allocNum;
} lns_globalEnv_t;


/**
 * リストから ANY を除外。
 */
#define lns_rmFromList( ANY )                  \
    if ( (ANY)->pNext != NULL ) {               \
        (ANY)->pPrev->pNext = (ANY)->pNext;     \
        (ANY)->pNext->pPrev = (ANY)->pPrev;     \
        (ANY)->pNext = NULL;                    \
    }

#define DECRE_STEM( ENV, STEM )                          \
    if ( STEM.type == lns_stem_type_any ) {             \
        if ( STEM.val.pAny != NULL ) {                   \
            lns_decre_ref( ENV, STEM.val.pAny );        \
        }                                                \
    }


#define lns_malloc( ALLOCATOR, SIZE ) _lns_malloc( ALLOCATOR, SIZE, __FILE__, __LINE__ )
#define lns_free( ALLOCATOR, ADDR ) _lns_free( ALLOCATOR, ADDR, __FILE__, __LINE__ )

static lns_globalEnv_t s_globalEnv;
lns_global_t lns_global;

#define lns_alloc_any( ENV, TYPE, FILE, LINE ) \
    _lns_alloc_any( ENV, TYPE, FILE, LINE )
#define lns_alloc_any_op( ENV, TYPE )                  \
    _lns_alloc_any( ENV, TYPE, __FILE__, __LINE__ )

#define LNS_FORM_MAX_ARG 20

static void lns_class_del( lns_env_t * _pEnv, void * pObj );
static void lns_alge_del( lns_env_t * _pEnv, void * pObj );



lns_type_meta_t lns_type_meta_lns__root = {
    "_root",
    &lns_type_meta_lns__root,
    { NULL }
};


lns_type_meta_t lns_type_meta_List = { "List", &lns_type_meta_lns__root, { NULL } };
lns_type_meta_t lns_type_meta_Set = { "Set", &lns_type_meta_lns__root, { NULL } };
lns_type_meta_t lns_type_meta_Map = { "Map", &lns_type_meta_lns__root, { NULL } };


void _lns_abort( const char * pMessage, const char * pFile, int lineNo )
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
 * lns_setq() で、何らかの変数に格納しない限り、ブロック終了時に開放される。
 *
 * @return 生成した any 値
 */
static lns_any_t * _lns_alloc_any(
    lns_env_t * _pEnv, lns_value_type_t type, const char * pFile, int lineNo )
{
    s_globalEnv.allocNum++;
    
    lns_any_t * pAny = (lns_any_t *)_lns_malloc(
        _pEnv->allocateor, sizeof( lns_any_t ), pFile, lineNo );
    lns_init_any( pAny, type, _pEnv );
    pAny->refCount = 1;

    // 現在のブロックに登録
    lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lns_add2list( &pBlock->managedAnyTop, pAny );

    return pAny;
}


/**
 * pAny を開放する。
 *
 * 参照カウントが 0 になっていることが前提条件。
 */
static void lns_release_any( lns_env_t * _pEnv, lns_any_t * pAny ) {
    if ( pAny->type == lns_value_type_str ) {
        pAny->type = lns_value_type_str;
    }
    pAny->type = lns_value_type_none; // このセットは必須ではないが、
    // セットしておいた方が不具合を見つけ易い。
    lns_free( _pEnv->allocateor, pAny );
    s_globalEnv.allocNum--;
}

lns_var_t * lns_var_alloc(
    lns_env_t * _pEnv, lns_block_t * pBlock, int index, lns_stem_type_t type )
{
    lns_var_t * pVar =
        (lns_var_t *)lns_malloc( _pEnv->allocateor, sizeof( lns_var_t ) );
    s_globalEnv.allocNum++;
    
    pVar->stem.type = type;
    pBlock->pVarList[ index ] = pVar;
    pVar->refCount = 1;
    return pVar;
}

static void lns_var_decre( lns_env_t * _pEnv, lns_var_t * pVar )
{
    lns_lock(
        pVar->refCount--;
        if ( pVar->refCount == 0 ) {
            DECRE_STEM( _pEnv, pVar->stem );
            lns_free( _pEnv->allocateor, pVar );
            s_globalEnv.allocNum--;
        }
    );
}

static void lns_gc_any( lns_env_t * _pEnv, lns_any_t * pAny, bool freeFlag ) {
    switch ( pAny->type ) {
    case lns_value_type_str:
        if ( !pAny->val.str.staticFlag ) {
            lns_free( _pEnv->allocateor, (char *)pAny->val.str.pStr );
            s_globalEnv.allocNum--;
        }
        break;
    case lns_value_type_class:
        if ( ((lns_Class_t*)pAny->val.classVal)->pMtd->_gc != NULL ) {
            ((lns_Class_t*)pAny->val.classVal)->pMtd->_gc( _pEnv, pAny );
        }
        if ( ((lns_Class_t*)pAny->val.classVal)->pMtd->_del != NULL ) {
            ((lns_Class_t*)pAny->val.classVal)->pMtd->_del( _pEnv, pAny );
        }
        lns_class_del( _pEnv, pAny->val.classVal );
        break;
    case lns_value_type_alge:
        if ( pAny->val.alge.gc != NULL ) {
            pAny->val.alge.gc( _pEnv, pAny->val.alge.pVal );
        }
        lns_alge_del( _pEnv, pAny->val.alge.pVal );
        break;
    case lns_value_type_ddd: // fall-through
    case lns_value_type_mRet:
        if ( pAny->val.ddd.len != 0 ) {
            lns_free( _pEnv->allocateor, pAny->val.ddd.stemList );
            s_globalEnv.allocNum--;
        }
        break;
    case lns_value_type_form:
        {
            int index;
            for ( index = 0; index < pAny->val.form.len; index++ ) {
                lns_closureVal_t * pVal = l_form_closure( pAny, index );
                lns_var_decre( _pEnv, pVal->pVar );
            }
            lns_free( _pEnv->allocateor, pAny->val.form.pClosureValList );
            s_globalEnv.allocNum--;
        }
        break;
    case lns_value_type_itList:
        lns_itList__del( _pEnv, pAny );
        break;
    case lns_value_type_itSet:
        lns_itSet__del( _pEnv, pAny );
        break;
    case lns_value_type_itMap:
        lns_itMap__del( _pEnv, pAny );
        break;
    case lns_value_type_if:
        lns_decre_ref( _pEnv, pAny->val.ifVal.pObj );
        // if は class 内のメンバなので開放しないで return する。
        return;
    case lns_value_type_luaVal:
        lns_releaeAnyVal( _pEnv, pAny );
        break;
    case lns_value_type_luaForm:
        lns_releaeAnyVal( _pEnv, pAny );
        break;
    default:
        break;
    }

    if ( freeFlag ) {
        lns_release_any( _pEnv, pAny );
    }
}

/**
 * pAny の参照カウントをデクリメントする。
 *
 * 参照カウントが 0 になった場合は、開放処理を行なう。
 */
bool lns_decre_ref( lns_env_t * _pEnv, lns_any_t * pAny ) {
    lns_lock(
        pAny->refCount--;
        if ( pAny->refCount == 0 ) {
            lns_gc_any( _pEnv, pAny, true );
            return true;
        }
    );
    return false;
}


/**
MRet の戻り値を設定し、先頭の値を返す。
 */
lns_stem_t lns_setMRet( lns_env_t * _pEnv, lns_any_t * pAny ) {
    _pEnv->pMRet = pAny;
    return lns_fromDDD( pAny, 0 );
}


lns_any_t * lns_getIF( lns_env_t * _pEnv, lns_any_t * pIFAny )
{
    lns_lock(
        if ( pIFAny->refCount == 0 ) {
            // interface の any は共有なので、
            // 参照回数が 0 の状態から 1 になったタイミングで、
            // 元インスタンスの参照回数をインクリメントする。
            pIFAny->val.ifVal.pObj->refCount++;

            lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
            lns_add2list( &pBlock->managedAnyTop, pIFAny );
            pIFAny->refCount++;
        }
    );
    return pIFAny;
}

/**
 * pAny が class ならば、指定の pMeta の IF のインスタンスを返す。
 * pAny が IF なら、そのまま pAny を返す。
 */
lns_stem_t lns_toIF(
    lns_env_t * _pEnv, lns_any_t * pAny, const lns_type_meta_t * pMeta )
{
    if ( pAny->type == lns_value_type_if ) {
        return LNS_STEM_ANY( pAny );
    }
    if ( pAny->type == lns_value_type_class ) {
        lns_any_t * pIFAny = (lns_any_t *)pAny->val.classVal->pIFdummy;
        for ( ; pIFAny->type != lns_value_type_none; pIFAny++ ) {
            if ( pIFAny->val.ifVal.pMeta == pMeta ) {
                return LNS_STEM_ANY( lns_getIF( _pEnv, pIFAny ) );
            }
        }
    }
    lns_abort( "not found interface" );
    return lns_global.nilStem;
}

static lns_any_t * lns_toIFDirectly(
    lns_env_t * _pEnv, lns_any_t * pAny, const lns_type_meta_t * pMeta )
{
    if ( pAny->type == lns_value_type_class ) {
        lns_any_t * pIFAny = (lns_any_t *)pAny->val.classVal->pIFdummy;
        for ( ; pIFAny->type != lns_value_type_none; pIFAny++ ) {
            if ( pIFAny->val.ifVal.pMeta == pMeta ) {
                return pIFAny;
            }
        }
    }
    lns_abort( "not found interface" );
    return NULL;
}

void lns_setQ_( lns_any_t * pAny )
{
    if ( pAny != NULL ) {
        lns_lock( 
            pAny->refCount++;
        );
    }
}

void lns_setRetAtBlockSub( lns_block_t * pBlock, lns_stem_t * pStem )
{
    if ( pStem->type != lns_stem_type_any ) {
        return;
    }

    lns_any_t * pAny = pStem->val.pAny;

    if ( pAny->pNext == NULL ) {
        return;
    }
    lns_rmFromList( pAny );

    lns_add2list( &pBlock->managedAnyTop, pAny );

    if ( pAny->type == lns_value_type_ddd ||
         pAny->type == lns_value_type_mRet )
    {
        int index;
        for ( index = 0; index < pAny->val.ddd.len; index++ ) {
            lns_stem_t item = lns_fromDDD( pAny, index );
            lns_setRetAtBlockSub( pBlock, &item );
        }
    }
    return;
}

void lns_setRetAtBlock( lns_block_t * pBlock, lns_stem_t stem )
{
    lns_setRetAtBlockSub( pBlock, &stem );
}

void lns_setRet( lns_env_t * _pEnv, lns_stem_t stem )
{
    lns_setRetAtBlockSub( LNS_BLOCK_AT( _pEnv, 1 ), &stem );
}

void lns_setRetInBlock( lns_env_t * _pEnv, lns_stem_t stem )
{
    lns_setRetAtBlockSub( LNS_BLOCK_AT( _pEnv, 0 ), &stem );
}


void lns_setup_block(
    lns_env_t * _pEnv, lns_block_t * pBlock, int anyNum, int stemNum, int varNum )
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
lns_block_t * lns_enter_module( lns_env_t * _pEnv, int anyNum, int stemNum, int varNum )
{
    lns_block_t * pBlock = &s_globalEnv.moduleInitBlockBuf[ s_globalEnv.moduleNum ];
    s_globalEnv.moduleNum++;
    
    lns_setup_block( s_globalEnv.pEnv, pBlock, anyNum, stemNum, varNum );
    //lns_setup_block( _pEnv, pBlock, anyNum, stemNum, varNum );

    return pBlock;
}

static inline void lns_reset_blockSub( lns_env_t * _pEnv, lns_block_t * pBlock ) {
    lns_any_t * pWork = pBlock->managedAnyTop.pPrev;
    while ( pWork != &pBlock->managedAnyTop ) {
        lns_any_t * pPrev = pWork->pPrev;

        lns_lock( 
            if ( pWork->refCount == 1 ) {
                lns_gc_any( _pEnv, pWork, true );
            }
            else {
                pWork->refCount--;
                // オブジェクトが残る場合は、チェーンから除外する
                pWork->pNext = NULL;
            }
        );
        pWork = pPrev;
    }
}

/**
 * 新しくブロックを開始する。
 *
 * @param anyVerNum ブロックで管理する any 型の値の数
 * @return ブロック情報
 */
lns_block_t * lns_enter_block(
    lns_env_t * _pEnv, int anyNum, int stemNum, int varNum )
{
    _pEnv->blockDepth++;
    
    lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lns_setup_block( _pEnv, pBlock, anyNum, stemNum, varNum );

    return pBlock;
}


/**
 * 現在のブロックを終了する。
 *
 * 次の処理を行なう。
 * - ブロックで管理している any 型の値の参照カウンタをデクリメント
 * - 変数にアサインされていない any 型の値を開放
 */
static void lns_leave_blockSub( lns_env_t * _pEnv, lns_block_t * pBlock )
{
    int index;
    for ( index = pBlock->varLen - 1; index >= 0; index-- ) {
        lns_var_t * pVar = pBlock->pVarList[ index ];
        if ( pVar != NULL ) {
            lns_var_decre( _pEnv, pVar );
        }
    }
    for ( index = pBlock->stemLen - 1; index >= 0; index-- ) {
        lns_stem_t * pStem = pBlock->pStemList[ index ];
        if ( pStem != NULL && pStem->type == lns_stem_type_any ) {
            if ( lns_decre_ref( _pEnv, pStem->val.pAny ) ) {
                pStem->type = lns_stem_type_none;
                pBlock->pStemList[ index ] = NULL;
            }
        }
    }
    for ( index = pBlock->anyLen - 1; index >= 0; index-- ) {
        lns_any_t * pAny = pBlock->pAnyList[ index ];
        if ( pAny != NULL ) {
            lns_decre_ref( _pEnv, pAny );
            pBlock->pAnyList[ index ] = NULL;
        }
    }

    lns_reset_blockSub( _pEnv, pBlock );
    
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
void lns_leave_block( lns_env_t * _pEnv )
{
    lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    _pEnv->blockDepth--;

    lns_leave_blockSub( _pEnv, pBlock );
}

/**
 * 現在のブロックをクリアする。
 */
void lns_reset_block( lns_env_t * _pEnv )
{
    lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    int index;

    for ( index = pBlock->varLen - 1; index >= 0; index-- ) {
        lns_var_t * pVar = pBlock->pVarList[ index ];
        if ( pVar != NULL ) {
            lns_var_decre( _pEnv, pVar );
        }
    }
    
    for ( index = pBlock->stemLen - 1; index >= 0; index-- ) {
        lns_stem_t * pStem = pBlock->pStemList[ index ];
        if ( pStem != NULL && pStem->type == lns_stem_type_any ) {
            lns_decre_ref( _pEnv, pStem->val.pAny );
            pBlock->pStemList[ index ] = NULL;
        }
        
    }

    for ( index = pBlock->anyLen - 1; index >= 0; index-- ) {
        lns_any_t * pAny = pBlock->pAnyList[ index ];
        if ( pAny != NULL ) {
            lns_decre_ref( _pEnv, pAny );
            pBlock->pAnyList[ index ] = NULL;
        }
    }

    lns_reset_blockSub( _pEnv, pBlock );

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
void lns_leave_blockMulti( lns_env_t * _pEnv, int num )
{
    while ( num > 0 ) {
        lns_leave_block( _pEnv );
        num--;
    }
}

/**
 * 関数開始時の処理
 *
 * lns_enter_block() とほぼ同じ。
 * ただし、関数の引数の処理が追加。
 * 関数終了時は、 lns_leave_block() をコールする。
 *
 * @param num lns_enter_block() の anyVerNum と同じ。
 *   ブロック内の any 値の数と、 argNum を合せた数が num になる。
 * @param argNum 引数の any 型の値の数
 * @param ... 引数の any 型の値。
 */
/* lns_block_t * lns_enter_func( */
/*     lns_env_t * _pEnv, int stemNum, int varNum, int argNum, ... ) */
/* { */
/*     lns_enter_block( _pEnv, stemNum, varNum ); */
/*     lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ]; */


/*     va_list ap; */
/*     va_start( ap, argNum ); */

/*     int index; */
/*     for ( index = 0; index < argNum; index++ ) { */
/*         lns_any_t * pAny = va_arg( ap, lns_any_t * ); */
/*         lns_setq( _pEnv, &pBlock->pVarList[ index ]->stem.val.pAny, pAny ); */
/*         pBlock->pVarList[ index ]->stem.type = lns_stem_type_any; */
/*     } */
/*     va_end(ap); */

/*     return pBlock; */
/* } */

lns_stem_t lns_getValFromDDD( lns_any_t * pAny, int index )
{
    if ( pAny->val.ddd.len <= index ) {
        return lns_global.nilStem;
    }
    return pAny->val.ddd.stemList[ index ];
}


static lns_any_t * lns_createDDDSub(
    lns_env_t * _pEnv, bool hasDDD, int num, va_list apSrc, lns_any_t * pDDDAny ) {
    lns_ddd_t * pDDD = &pDDDAny->val.ddd;
    int argNum = num;

    if ( hasDDD ) {
        // 最終要素が複数要素の場合、その複数要素が何個あるのか調べる
        argNum = num;

        va_list ap;
        va_copy( ap, apSrc );

        int index;
        for ( index = 0; index < num - 1; index++ ) {
            va_arg( ap, lns_stem_t );
        }
        lns_stem_t dddStem = va_arg( ap, lns_stem_t );
        argNum += dddStem.val.pAny->val.ddd.len - 1;
        va_end(ap);
    }
    
    pDDD->len = argNum;
    pDDD->stemList = (lns_stem_t *)lns_malloc(
        _pEnv->allocateor, sizeof( lns_stem_t ) * argNum );
    s_globalEnv.allocNum++;

    va_list ap;
    va_copy( ap, apSrc );

    int index;
    for ( index = 0; index < num - 1; index++ ) {
        lns_stem_t stem = va_arg( ap, lns_stem_t );
        if ( stem.type == lns_stem_type_any ) {
            lns_any_t * pAny = stem.val.pAny;
            if ( pAny->type == lns_value_type_ddd ||
                 pAny->type == lns_value_type_mRet )
            {
                pDDD->stemList[ index ] = lns_getValFromDDD( pAny, 0 );
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
        lns_stem_t stem = va_arg( ap, lns_stem_t );
        lns_any_t * pAny = stem.val.pAny;
        lns_stem_t* pStem = &pDDD->stemList[ num - 1 ];
        for ( index = 0; index < lns_lenDDD( pAny ); index++ ) {
            *pStem = lns_fromDDD( pAny, index );
            pStem++;
        }
    }
    else {
        pDDD->stemList[ num - 1 ] = va_arg( ap, lns_stem_t );
    }

    
    va_end(ap);

    return pDDDAny;
}


/**
 * 多値返却を生成する。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lns_any_t * lns_createMRetOnly( lns_env_t * _pEnv, int num )
{
    lns_any_t * pDDDAny = _lns_alloc_any( _pEnv, lns_value_type_mRet, LNS_DEBUG_POS );

    lns_ddd_t * pDDD = &pDDDAny->val.ddd;
    pDDD->len = num;
    if ( pDDD->len != 0 ) {
        pDDD->stemList = (lns_stem_t *)lns_malloc(
            _pEnv->allocateor, sizeof( lns_stem_t ) * num );
        s_globalEnv.allocNum++;
    }
    else {
        pDDD->stemList = NULL;
    }
    
    return pDDDAny;
}


/**
 * 多値返却を生成する。
 *
 * @param hasDDD ... の最後が ... 要素の場合 true。
 * @param num 値の数
 * @param ... 含める値
 * @return ... の値
 */
lns_stem_t _lns_createMRet(
    const char * pFile, int lineNo, lns_env_t * _pEnv, bool hasDDD, int num, ... )
{
    lns_any_t * pDDDAny =
        _lns_alloc_any( _pEnv, lns_value_type_mRet, pFile, lineNo );

    va_list ap;
    va_start( ap, num );

    lns_createDDDSub( _pEnv, hasDDD, num, ap, pDDDAny );

    va_end(ap);

    return LNS_STEM_ANY( pDDDAny );
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
lns_stem_t _lns_createDDD(
    const char * pFile, int lineNo, lns_env_t * _pEnv, bool hasDDD, int num, ... )
{
    if ( num > 0 ) {
        lns_any_t * pDDDAny =
            lns_alloc_any( _pEnv, lns_value_type_ddd, pFile, lineNo );

        va_list ap;
        va_start( ap, num );
        
        lns_createDDDSub( _pEnv, hasDDD, num, ap, pDDDAny );

        va_end(ap);

        return LNS_STEM_ANY( pDDDAny );
    }
    else {
        return lns_createDDDOnly( _pEnv, 0 );
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
lns_stem_t _lns_createSubDDD(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int offset, lns_stem_t ddd )
{
    lns_any_t * pDDD = ddd.val.pAny;
    
    int len;
    len = lns_lenDDD( pDDD ) - offset;
    if ( len < 0 ) {
        return _lns_createDDDOnly( pFile, lineNo, _pEnv, 0 );
    }

    lns_stem_t pSubDDD = _lns_createDDDOnly( pFile, lineNo, _pEnv, len );
        
    int index;
    for ( index = 0; index < len; index++ ) {
        lns_set2DDDArg( pSubDDD.val.pAny, index, lns_fromDDD( pDDD, index + offset ) );
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
lns_stem_t _lns_createDDDOnly(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int num ) {
    lns_any_t * pDDDAny =
        lns_alloc_any( _pEnv, lns_value_type_ddd, pFile, lineNo );
    lns_ddd_t * pDDD = &pDDDAny->val.ddd;
    pDDD->len = num;
    if ( num != 0 ) {
        pDDD->stemList = (lns_stem_t *)lns_malloc(
            _pEnv->allocateor, sizeof( lns_stem_t ) * num );
        s_globalEnv.allocNum++;
    }
    else {
        pDDD->stemList = NULL;
    }

    return LNS_STEM_ANY( pDDDAny );
}


lns_any_t * _lns_luaVal_new(
    const char * pFile, int lineNo, lns_env_t * _pEnv, lns_value_type_t type )
{
    lns_any_t * pAny = lns_alloc_any( _pEnv, lns_value_type_luaVal, pFile, lineNo );
    pAny->val.luaVal.type = type;
    return pAny;
}

lns_any_t * lns_luaForm_new( lns_env_t * _pEnv )
{
    return _lns_alloc_any( _pEnv, lns_value_type_luaForm, LNS_DEBUG_POS );
}


lns_any_t * _lns_it_new(
    const char * pFile, int lineNo, 
    lns_env_t * _pEnv, lns_value_type_t type, void * pVal )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, type, pFile, lineNo );
    switch ( type ) {
    case lns_value_type_itList:
        pAny->val.itList = pVal;
        break;
    case lns_value_type_itSet:
        pAny->val.itSet = pVal;
        break;
    case lns_value_type_itMap:
        pAny->val.itMap = pVal;
        break;
    default:
        break;
    }
    return pAny;
}

/* void lns_it_delete( lns_env_t * _pEnv, lns_any_t * pAny ) */
/* { */
/*     lns_rmFromList( pAny ); */
/*     lns_gc_any( _pEnv, pAny, true ); */
/* } */


/**
 * クラスのインスタンスを保持する any を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return any
 */
lns_any_t * _lns_alge_new(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int valType,
    int size, lns_algeVal_gc_t * gc )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, lns_value_type_alge, pFile, lineNo );
    void * pObj = lns_malloc( _pEnv->allocateor, size );
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
static void lns_alge_del( lns_env_t * _pEnv, void * pObj )
{
    s_globalEnv.allocNum--;
    lns_free( _pEnv->allocateor, pObj );
}

/**
 * クラスのインスタンスを保持する any を生成する
 *
 * @param size クラスインスタンスのサイズ
 * @return any
 */
lns_any_t * _lns_class_new(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int size )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, lns_value_type_class, pFile, lineNo );
    void * pObj = lns_malloc( _pEnv->allocateor, size );
    s_globalEnv.allocNum++;
    pAny->val.classVal = pObj;
    return pAny;
}

/**
 * クラスのインスタンスを開放する
 *
 * @param pObj クラスのインスタンス
 */
static void lns_class_del( lns_env_t * _pEnv, void * pObj )
{
    s_globalEnv.allocNum--;
    lns_free( _pEnv->allocateor, pObj );
}

/**
 * 関数 pFunc を保持する any を生成する
 *
 * @param pFunc 関数
 * @param num フォーム内でアクセスする外部変数の数
 * @param ... フォーム内でアクセスする外部変数
 * @return any
 */
lns_any_t * _lns_func2any(
    const char * pFile, int lineNo, 
    lns_env_t * _pEnv, lns_closure_t * pFunc, int argNum, bool hasDDD, int num, ... )
{
    lns_any_t * pFormAny =
        lns_alloc_any( _pEnv, lns_value_type_form, pFile, lineNo );
    lns_form_t * pForm = &pFormAny->val.form;
    pForm->len = num;
    if ( num == 0 ) {
        pForm->needFormParam = false;
    }
    else {
        pForm->needFormParam = true;
    }
    pForm->pClosureValList = (lns_closureVal_t*)lns_malloc(
        _pEnv->allocateor, sizeof( lns_closureVal_t ) * num );
    s_globalEnv.allocNum++;
    /* pForm->pOrgClosureValList = (lns_any_t **)lns_malloc( */
    /*     _pEnv->allocateor, sizeof( lns_any_t * ) * num ); */
    /* s_globalEnv.allocNum++; */

    pForm->pFunc = pFunc;
    //pForm->argNum = argNum;
    pForm->hasDDD = hasDDD;
    
    va_list ap;
    va_start( ap, num );

    lns_lock( 
        int index;
        for ( index = 0; index < num; index++ ) {
            lns_closureVal_t * pClosureVal = &pForm->pClosureValList[ index ];
            lns_var_t * pVar = va_arg( ap, lns_var_t * );
            pVar->refCount++;
            pForm->pClosureValList[ index ].pVar = pVar;
        }
    );
    va_end(ap);

    return pFormAny;
}


bool lns_fromMapPrim(
    lns_env_t * _pEnv, lns_any_t ** ppErr, lns_stem_t * pWork,
    lns_any_t * pMap, bool nilable, lns_any_t * pMbr, lns_stem_type_t kind )
{
    lns_stem_t _name = LNS_STEM_ANY( pMbr );
    *pWork = lns_mtd_Map_get( _pEnv, pMap, _name );
    if ( pWork->type != kind ) {
        if ( !nilable || pWork->type != lns_stem_type_nil ) {
            *ppErr = pMbr;
            return false;
        }
    }
    return true;
}

bool lns_fromMapStr(
    lns_env_t * _pEnv, lns_any_t ** ppErr, lns_stem_t * pWork,
    lns_any_t * pMap, bool nilable, lns_any_t * pMbr )
{
    lns_stem_t _name = LNS_STEM_ANY( pMbr );
    *pWork = lns_mtd_Map_get( _pEnv, pMap, _name );
    if ( pWork->type == lns_stem_type_any &&
         pWork->val.pAny->type == lns_value_type_str ) {
    }
    else if ( !nilable || pWork->type != lns_stem_type_nil ) {
        *ppErr = pMbr;
        return false;
    }
    return true;
}

lns_stem_t lns_fromMapToClass(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ||
         stem.val.pAny->type != lns_value_type_class ||
         stem.val.pAny->val.classVal->pMeta != &lns_type_meta_List )
    {
        return lns_global.ddd0;
    }
    return pInfo->pFromMap( _pEnv, pInfo, stem );
}

lns_stem_t lns_toMapFromStem( lns_env_t * _pEnv, lns_stem_t item )
{
    switch ( item.type )
    {
    case lns_stem_type_nil: // fall-through
    case lns_stem_type_int: // fall-through
    case lns_stem_type_real: // fall-through
    case lns_stem_type_bool:
        return item;
    case lns_stem_type_any:
        switch ( item.val.pAny->type )
        {
        case lns_value_type_class:
            {
                lns_Class_t * pClass = item.val.pAny->val.classVal;
                if ( pClass->pMeta == &lns_type_meta_List ) {
                    return lns_toMapFromList( _pEnv, item );
                }
                else if ( pClass->pMeta == &lns_type_meta_Set ) {
                    return lns_toMapFromSet( _pEnv, item );
                }
                else if ( pClass->pMeta == &lns_type_meta_Map ) {
                    return lns_toMapFromMap( _pEnv, item );
                }
                else {
                    // Mapping IF を取得する
                    lns_any_t * pMapping = lns_toIFDirectly(
                        _pEnv, item.val.pAny, &lns_type_meta_lns_Mapping );
                    lns_toMap_t * pToMap = (lns_toMap_t*)lns_mtd_lns_Mapping( pMapping )->_toMap;
                    return LNS_STEM_ANY( pToMap( _pEnv, item.val.pAny ) );
                }
            }
            break;
        case lns_value_type_str:
            return item;
        default:
            break;
        }
        break;
    default:
        break;
    }
    lns_abort( "illegal type" );
    return lns_global.nilStem;
}


lns_stem_t lns_toMapFromList( lns_env_t * _pEnv, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ) {
        return lns_global.nilStem;
    }
    lns_any_t * pList = stem.val.pAny;

    lns_any_t * pNewList = lns_class_List_new( _pEnv );
    
    lns_stem_t item;
    lns_any_t * pIt = lns_itList_new( _pEnv, pList );
    for ( ; lns_itList_hasNext( _pEnv, pIt, &item ); lns_itList_inc( _pEnv, pIt ) ) {
        lns_mtd_List_insert( _pEnv, pNewList, lns_toMapFromStem( _pEnv, item ) );
    }
    // pIt は managedAnyTop で管理されているので、ここでは開放しない

    return LNS_STEM_ANY( pNewList );
}

lns_stem_t lns_toMapFromSet( lns_env_t * _pEnv, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ) {
        return lns_global.nilStem;
    }
    lns_any_t * pSet = stem.val.pAny;

    lns_any_t * pNewSet = lns_class_Set_new( _pEnv );
    
    lns_stem_t item;
    lns_any_t * pIt = lns_itSet_new( _pEnv, pSet );
    for ( ; lns_itSet_hasNext( _pEnv, pIt, &item ); lns_itSet_inc( _pEnv, pIt ) ) {
        lns_mtd_Set_add( _pEnv, pNewSet, lns_toMapFromStem( _pEnv, item ) );
    }
    // pIt は managedAnyTop で管理されているので、ここでは開放しない

    return LNS_STEM_ANY( pNewSet );
}

lns_stem_t lns_toMapFromMap( lns_env_t * _pEnv, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ) {
        return lns_global.nilStem;
    }
    lns_any_t * pMap = stem.val.pAny;

    lns_any_t * pNewMap = lns_class_Map_new( _pEnv );
    
    lns_Map_entry_t entry;
    lns_any_t * pIt = lns_itMap_new( _pEnv, pMap );
    for ( ; lns_itMap_hasNext( _pEnv, pIt, &entry ); lns_itMap_inc( _pEnv, pIt ) ) {
        lns_mtd_Map_add( _pEnv, pNewMap,
                         lns_toMapFromStem( _pEnv, entry.key ),
                         lns_toMapFromStem( _pEnv, entry.val ) );
    }
    // pIt は managedAnyTop で管理されているので、ここでは開放しない

    return LNS_STEM_ANY( pNewMap );
}

lns_stem_t lns_fromMapToStemSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    return lns_createMRet(_pEnv, false, 1, stem );
}
lns_stem_t lns_fromMapToIntSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_int ) {
        return lns_global.ddd0;
    }
    return lns_createMRet(_pEnv, false, 1, stem );
}
lns_stem_t lns_fromMapToRealSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_real ) {
        return lns_global.ddd0;
    }
    return lns_createMRet(_pEnv, false, 1, stem );
}
lns_stem_t lns_fromMapToBoolSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_bool ) {
        return lns_global.ddd0;
    }
    return lns_createMRet(_pEnv, false, 1, stem );
}
lns_stem_t lns_fromMapToStrSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ||
         stem.val.pAny->type != lns_value_type_str ) {
        return lns_global.ddd0;
    }
    return lns_createMRet( _pEnv, false, 1, stem );
}

lns_stem_t lns_fromMapToListSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ||
         stem.val.pAny->type != lns_value_type_class ||
         stem.val.pAny->val.classVal->pMeta != &lns_type_meta_List )
    {
        return lns_global.ddd0;
    }
    lns_any_t * pList = stem.val.pAny;

    lns_any_t * pNewList = lns_class_List_new( _pEnv );

    lns_any_t * pIt = lns_itList_new( _pEnv, pList );

    const struct lns_fromVal_info_t * pGenInfo = pInfo->pInfoArray[0];
    lns_stem_t val;
    lns_any_t * pErr = NULL;
    int index = 0;
    for ( ; lns_itList_hasNext( _pEnv, pIt, &val );
          lns_itList_inc( _pEnv, pIt ), index++ )
    {
        if ( val.type == lns_stem_type_nil ) {
            if ( pInfo->nilable ) {
                lns_mtd_List_insert( _pEnv, pNewList, lns_global.nilStem );
            }
            else {
                pErr = lns_string_format(
                    _pEnv, "%d", lns_createDDD( _pEnv, false, 1,
                                                LNS_STEM_INT( index + 1 ) ) );
                break;
            }
        }   
        else {
            lns_stem_t ddd = pGenInfo->pFromMap( _pEnv, pGenInfo, val );
            lns_stem_t work = lns_fromDDD( ddd.val.pAny, 0 );
            
            if ( work.type != lns_stem_type_nil ) {
                lns_mtd_List_insert( _pEnv, pNewList, work );
            }
            else {
                lns_stem_t err = lns_fromDDD( ddd.val.pAny, 1 );
                if ( err.type == lns_stem_type_any ) {
                    pErr = lns_string_format(
                        _pEnv, "%d.%s",
                        lns_createDDD( _pEnv, false, 2,
                                       LNS_STEM_INT( index + 1 ), err ) );
                }
                else {
                    pErr = lns_string_format(
                        _pEnv, "%d", lns_createDDD( _pEnv, false, 1,
                                                    LNS_STEM_INT( index + 1 ) ) );
                }
                break;
            }
        }
    }

    if ( pErr != NULL ) {
        lns_rmFromList( pNewList );
        lns_gc_any( _pEnv, pNewList, true );
        pNewList = NULL;
        return lns_createMRet( _pEnv, false, 2, lns_global.nilStem,
                               LNS_STEM_ANY( pErr ) );
    }
    
    return lns_createMRet( _pEnv, false, 1, LNS_STEM_ANY( pNewList ) );
}

lns_stem_t lns_fromMapToSetSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ||
         stem.val.pAny->type != lns_value_type_class ||
         stem.val.pAny->val.classVal->pMeta != &lns_type_meta_Set )
    {
        return lns_global.ddd0;
    }
    lns_any_t * pSet = stem.val.pAny;

    lns_any_t * pNewSet = lns_class_Set_new( _pEnv );

    lns_any_t * pIt = lns_itSet_new( _pEnv, pSet );

    const struct lns_fromVal_info_t * pGenInfo = pInfo->pInfoArray[0];
    lns_stem_t val;
    lns_stem_t errStem = lns_global.noneStem;
    for ( ; lns_itSet_hasNext( _pEnv, pIt, &val ); lns_itSet_inc( _pEnv, pIt ) )
    {
        if ( val.type == lns_stem_type_nil ) {
            if ( pInfo->nilable ) {
                lns_mtd_Set_add( _pEnv, pNewSet, lns_global.nilStem );
            }
            else {
                // set は順番が固定でないのでインデックス情報を付加しない
                errStem = lns_global.nilStem;
                break;
            }
        }
        else {
            lns_stem_t ddd = pGenInfo->pFromMap( _pEnv, pGenInfo, val );
            lns_stem_t work = lns_fromDDD( ddd.val.pAny, 0 );
            
            if ( work.type != lns_stem_type_nil ) {
                lns_mtd_Set_add( _pEnv, pNewSet, work );
            }
            else {
                lns_stem_t err = lns_fromDDD( ddd.val.pAny, 1 );
                if ( err.type == lns_stem_type_any ) {
                    errStem = LNS_STEM_ANY(
                        lns_string_format(
                            _pEnv, ".%s",
                            lns_createDDD( _pEnv, false, 1, err ) ) );
                }
                else {
                    errStem = lns_global.nilStem;
                }
                break;
            }
        }
    }

    if ( errStem.type != lns_stem_type_none ) {
        lns_rmFromList( pNewSet );
        lns_gc_any( _pEnv, pNewSet, true );
        pNewSet = NULL;
        return lns_createMRet( _pEnv, false, 2, lns_global.nilStem, errStem );
    }
    
    return lns_createMRet( _pEnv, false, 1, LNS_STEM_ANY( pNewSet ) );
}

lns_stem_t lns_fromMapToMapSub(
    lns_env_t * _pEnv, const lns_fromVal_info_t * pInfo, lns_stem_t stem )
{
    if ( stem.type != lns_stem_type_any ||
         stem.val.pAny->type != lns_value_type_class ||
         stem.val.pAny->val.classVal->pMeta != &lns_type_meta_Map )
    {
        return lns_global.ddd0;
    }
    lns_any_t * pMap = stem.val.pAny;

    lns_any_t * pNewMap = lns_class_Map_new( _pEnv );

    lns_any_t * pIt = lns_itMap_new( _pEnv, pMap );

    const struct lns_fromVal_info_t * pKeyGenInfo = pInfo->pInfoArray[0];
    const struct lns_fromVal_info_t * pValGenInfo = pInfo->pInfoArray[1];

    lns_Map_entry_t entry;
    bool success = true;
    for ( ; lns_itMap_hasNext( _pEnv, pIt, &entry ); lns_itMap_inc( _pEnv, pIt ) ) {
        lns_stem_t keyStem;
        if ( entry.key.type != lns_stem_type_nil ) {
            lns_stem_t ddd = pKeyGenInfo->pFromMap( _pEnv, pKeyGenInfo, entry.key );
            keyStem = lns_fromDDD( ddd.val.pAny, 0 );
            if ( keyStem.type == lns_stem_type_nil ) {
                success = false;
                break;
            }
        }
        lns_stem_t valStem;
        if ( entry.val.type == lns_stem_type_nil ) {
            if ( pInfo->nilable ) {
                valStem = entry.val;
            }
            else {
                success = false;
                break;
            }
        }
        else {
            lns_stem_t ddd = pValGenInfo->pFromMap( _pEnv, pValGenInfo, entry.val );
            valStem = lns_fromDDD( ddd.val.pAny, 0 );
            
            if ( valStem.type == lns_stem_type_nil ) {
                success = false;
                break;
            }
        }
        lns_mtd_Map_add( _pEnv, pNewMap, keyStem, valStem );        
    }

    if ( !success ) {
        lns_rmFromList( pNewMap );
        lns_gc_any( _pEnv, pNewMap, true );
        pNewMap = NULL;
        return lns_global.ddd0;
    }
    
    return lns_createMRet( _pEnv, false, 1, LNS_STEM_ANY( pNewMap ) );
}


/**
 * リテラルな文字列型を生成する
 *
 * @param pStr 文字列
 * @return 文字列型データ
 */
lns_str_t lns_createLiteralStr( const char * pStr ) {
    lns_str_t str;
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
lns_any_t * _lns_str2any(
    const char * pFile, int lineNo, lns_env_t * _pEnv, lns_str_t val )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, lns_value_type_str, pFile, lineNo );
    pAny->val.str = val;
    return pAny;
}

lns_any_t * _lns_litStr2any(
    const char * pFile, int lineNo, lns_env_t * _pEnv, const char * pStr )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, lns_value_type_str, pFile, lineNo );
    pAny->val.str.len = strlen( pStr );
    pAny->val.str.pStr = pStr;
    pAny->val.str.staticFlag = true;
    return pAny;
}

lns_any_t * _lns_createBinAny(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int len )
{
    lns_any_t * pAny = lns_alloc_any( _pEnv, lns_value_type_str, pFile, lineNo );

    char * pAlloc = (char *)lns_malloc( _pEnv->allocateor, len + 1 );
    s_globalEnv.allocNum++;
    if ( pAlloc == NULL ) {
        lns_abort( "alloc error" );
    }
    pAlloc[ len ] = '\0';
    pAny->val.str.pStr = pAlloc;
    pAny->val.str.len = len;
    pAny->val.str.staticFlag = false;
    return pAny;
}


lns_any_t * _lns_cloneBin2any(
    const char * pFile, int lineNo, lns_env_t * _pEnv, const void * pBuf, int len )
{
    lns_any_t * pAny = _lns_createBinAny( pFile, lineNo, _pEnv, len );
    memcpy( (void *)pAny->val.str.pStr, pBuf, len );
    return pAny;
}

lns_any_t * lns_strconcat(
    lns_env_t * _pEnv, lns_any_t * pAny1, lns_any_t * pAny2 )
{
    lns_any_t * pResult = _lns_createBinAny(
        LNS_DEBUG_POS, _pEnv, pAny1->val.str.len + pAny2->val.str.len );
    memcpy( (void*)pResult->val.str.pStr, pAny1->val.str.pStr, pAny1->val.str.len );
    memcpy( (void*)(pResult->val.str.pStr + pAny1->val.str.len),
            pAny2->val.str.pStr, pAny2->val.str.len );
    return pResult;
}

/**
stem を pMeta のクラスにキャストする。   
 */
lns_stem_t lns_castClass( lns_stem_t stem, const lns_type_meta_t * pMeta )
{
    if ( stem.type == lns_stem_type_any ) {
        const lns_any_t * pAny = stem.val.pAny;
        if ( pAny->type == lns_value_type_class || pAny->type == lns_value_type_if )
        {
            const lns_type_meta_t * pWorkMeta = pAny->val.classVal->pMeta;
            while ( pWorkMeta != &lns_type_meta_lns__root ) {
                if ( pWorkMeta == pMeta ) {
                    return stem;
                }
                pWorkMeta = pWorkMeta->pSuper;
            }
        }
    }
    return lns_global.nilStem;
}

/**
stem を pMeta のインタフェースにキャストする。   
   
 */
lns_stem_t lns_castIf(
    lns_env_t * _pEnv, lns_stem_t stem, const lns_type_meta_t * pMeta )
{
    if ( stem.type == lns_stem_type_any ) {
        lns_any_t * pAny = stem.val.pAny;
        switch ( pAny->type ) {
        case lns_value_type_if:
            pAny = pAny->val.ifVal.pObj;
            break;
        case lns_value_type_class:
            break;
        default:
            return lns_global.nilStem;
        }
        lns_any_t * pIFAny = (lns_any_t *)pAny->val.classVal->pIFdummy;
        for ( ; pIFAny->type != lns_value_type_none; pIFAny++ ) {
            if ( pIFAny->val.ifVal.pMeta == pMeta ) {
                return LNS_STEM_ANY( lns_getIF( _pEnv, pIFAny ) );
            }
        }
    }
    return lns_global.nilStem;
}

lns_stem_t lns_castStem( lns_stem_t stem, lns_stem_type_t kind )
{
    if ( stem.type == kind ) {
        return stem;
    }
    else if ( kind == lns_stem_type_int && stem.type == lns_stem_type_real ) {
        if ( (lns_int_t)stem.val.realVal == stem.val.realVal ) {
            return LNS_STEM_INT( stem.val.realVal );
        }
    }
    else if ( kind == lns_stem_type_real && stem.type == lns_stem_type_int ) {
        return LNS_STEM_REAL( stem.val.intVal );
    }
    return lns_global.nilStem;
}

lns_stem_t lns_castAny( lns_stem_t stem, lns_value_type_t kind )
{
    if ( stem.type == lns_stem_type_any && stem.val.pAny->type == kind ) {
        return stem;
    }
    return lns_global.nilStem;
}

lns_int_t lns_stem2int( lns_stem_t stem )
{
    switch ( stem.type ) {
    case lns_stem_type_int:
        return stem.val.intVal;
    case lns_stem_type_real:
        return (lns_int_t)stem.val.realVal;
    default:
        break;
    }
    lns_abort( "convert error stem2int" );
    return 0;
}

lns_real_t lns_stem2real( lns_stem_t stem )
{
    switch ( stem.type ) {
    case lns_stem_type_int:
        return stem.val.intVal;
    case lns_stem_type_real:
        return stem.val.realVal;
    default:
        break;
    }
    lns_abort( "convert error stem2real" );
    return 0.0;
}

lns_bool_t lns_stem2bool( lns_stem_t stem )
{
    switch ( stem.type ) {
    case lns_stem_type_bool:
        return stem.val.boolVal;
    default:
        break;
    }
    lns_abort( "convert error stem2real" );
    return false;
}

bool lns_equals_any( const lns_any_t * pAny1, const lns_any_t * pAny2 )
{
    if ( pAny1->type == pAny2->type ) {
        if ( pAny1 == pAny2 ) {
            // 同じ any データなら equal
            return true;
        }
        switch ( pAny1->type ) {
        case lns_value_type_str:
            {
                // 文字列は同じデータなら equal
                const lns_str_t * pStr1 = &pAny1->val.str;
                const lns_str_t * pStr2 = &pAny2->val.str;
                if ( pStr1->len != pStr2->len ) {
                    // 長さが違えば false
                    return false;
                }
                return memcmp( pStr1->pStr, pStr2->pStr, pStr1->len ) == 0;
            }
            break;
        case lns_value_type_if:
            // インタフェースは、同じインスタンスを指している場合は equal
            return pAny1->val.ifVal.pObj == pAny2->val.ifVal.pObj;
        default:
            break;
        }
        return false;
    }
    // 異なる any タイプなら
    switch ( pAny1->type ) {
    case lns_value_type_if:
        if ( pAny2->type == lns_value_type_class ) {
            return pAny1->val.ifVal.pObj == pAny2;
        }
        break;
    case lns_value_type_class:
        if ( pAny2->type == lns_value_type_if ) {
            return pAny2->val.ifVal.pObj == pAny1;
        }
        break;
    default:
        break;
    }
    return false;
}

bool lns_equals( const lns_stem_t stem1, const lns_stem_t stem2 ) {
    if ( stem1.type == stem2.type ) {
        switch ( stem1.type ) {
        case lns_stem_type_nil:
            return true;
        case lns_stem_type_int:
            return stem1.val.intVal == stem2.val.intVal;
        case lns_stem_type_real:
            return stem1.val.realVal == stem2.val.realVal;
        case lns_stem_type_bool:
            return stem1.val.boolVal == stem2.val.boolVal;
        default:
            break;
        }
        return lns_equals_any( stem1.val.pAny, stem2.val.pAny );
    }

    // 異なる stem タイプなら
    switch ( stem1.type ) {
    case lns_stem_type_int:
        if ( stem2.type == lns_stem_type_real ) {
            return stem1.val.intVal == stem2.val.realVal;
        }
        break;
    case lns_stem_type_real:
        if ( stem2.type == lns_stem_type_int ) {
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
static lns_env_t * lns_createEnv( lua_State * pLua )
{
    lns_allocator_t allocateor = lns_createAllocator();
    
    lns_env_t * _pEnv =
        (lns_env_t *)lns_malloc( allocateor, sizeof( lns_env_t ) );
    s_globalEnv.allocNum++;
    
    _pEnv->allocateor = allocateor;
    _pEnv->pLua = pLua;
    _pEnv->useAnyPoolNum = 0;
    _pEnv->useStemPoolNum = 0;
    _pEnv->useVarPoolNum = 0;
    _pEnv->blockDepth = 0;
    _pEnv->stackPos = 0;
    _pEnv->loadModuleTop.pNext = &_pEnv->loadModuleTop;
    _pEnv->loadModuleTop.pPrev = &_pEnv->loadModuleTop;

    lns_enter_block( _pEnv, 0, 0, 0 );
    _pEnv->pMRet = NULL;

    _pEnv->pSortCallback = lns_global.nilStem;

    lns_enter_block( _pEnv, 0, 0, 0 );

    if ( pLua != NULL ) {
        lns_setLuaWapper( _pEnv );
    }
    
    return _pEnv;
}

#define DEBUG_MEM_LOG( SYM )                                            \
    if ( _pEnv->SYM != 0 ) {                                            \
        result = false;                                                 \
    }                                                                   \
    printf( ":debug:" #SYM "= %d\n", _pEnv->SYM );


/**
 * 環境を開放する。
 *
 * @param _pEnv 環境
 */
static bool lns_deleteEnv( lns_env_t * _pEnv ) {

    lns_leave_block( _pEnv );
    lns_leave_block( _pEnv );

    lns_module_t * pModule = _pEnv->loadModuleTop.pNext;
    for ( ; pModule != &_pEnv->loadModuleTop; pModule = pModule->pNext )
    {
        //lns_leave_blockSub( _pEnv, pModule->pBlock );
        lns_leave_blockSub( s_globalEnv.pEnv, pModule->pBlock );
    }
    
    
    s_globalEnv.allocNum--;

    bool result = true;
    DEBUG_MEM_LOG( useAnyPoolNum );
    DEBUG_MEM_LOG( useStemPoolNum );
    DEBUG_MEM_LOG( useVarPoolNum );
    DEBUG_MEM_LOG( blockDepth );
    
    lns_free( _pEnv->allocateor, _pEnv );

    return result;
}

static void lns_createGlobalEnv() {
    s_globalEnv.allocNum = 0;
    s_globalEnv.pEnv = lns_createEnv( NULL );
    s_globalEnv.moduleNum = 0;

    
    lns_global.noneStem = LNS_STEM_BASE( lns_stem_type_none );
    lns_global.nilStem = LNS_STEM_BASE( lns_stem_type_nil );
    lns_global.trueStem = LNS_STEM_BOOL( true );
    lns_global.falseStem = LNS_STEM_BOOL( false );

    lns_setQ( lns_global.ddd0, lns_createDDD( s_globalEnv.pEnv, false, 0 ) );
}

static bool lns_releaseGlobalEnv(void) {
    /* int index; */
    /* for ( index = s_globalEnv.moduleNum - 1; index >= 0; index-- ) { */
    /*     lns_leave_blockSub( s_globalEnv.pEnv, */
    /*                          &s_globalEnv.moduleInitBlockBuf[ index ] ); */
    /* } */
    lns_decre_ref( s_globalEnv.pEnv, lns_global.ddd0.val.pAny );
    bool result = lns_deleteEnv( s_globalEnv.pEnv );
    printf( ":debug:allocNum = %d\n", s_globalEnv.allocNum );
    if ( s_globalEnv.allocNum != 0 ) {
        result = false;
    }
    printf( ":debug:collection_allocNum = %d\n", lns_collection_getAllocNum() );
    if ( lns_collection_getAllocNum() != 0 ) {
        result = false;
    }
    lns_checkMem();

    return result;
}

void lns_init_alge( lns_stem_t * pStem, lns_any_t * pAny, int valType )
{
    pStem->type = lns_stem_type_any;
    pStem->val.pAny = pAny;
    lns_init_any( pAny, lns_value_type_alge, ENV );
    pAny->refCount = 1;                                      
    pAny->val.alge.type = valType;                           
    pAny->val.alge.pVal = NULL;                              
    pAny->val.alge.gc = NULL;
}

lns_stem_t lns_call_form( lns_env_t * _pEnv, lns_any_t * _pForm, lns_stem_t ddd )
{
    if ( lns_isClosure( _pForm ) ) {
        return lns_closure( _pForm )( _pEnv, _pForm, ddd );
    }
    else {
        return lns_func( _pForm )( _pEnv, ddd );
    }
}

bool lns_op_not( lns_env_t * _pEnv, lns_stem_t stem ) {
    return !lns_isCondTrue( stem );
}

/**
 * pAny の条件判定を行なう
 *
 * false か nil の場合 false、それ以外は true
 */
bool lns_isCondTrue( const lns_stem_t stem ) {
    if ( stem.type == lns_stem_type_bool ) {
        return stem.val.boolVal;
    }
    else if ( stem.type == lns_stem_type_nil ) {
        return false;
    }
    return true;
}

/**
 * スタックを一段上げる
 */
bool lns_incStack( lns_env_t * _pEnv ) {
    _pEnv->stackPos++;
    return false;
}

/**
 * 値 pVal をスタックの top にセットし、値 pVal の条件判定結果を返す
 *
 * スタックに lns_ddd_t は詰めない。
 * 呼び出し側で lns_ddd_t の先頭要素を指定すること。
 *
 * @param pVal スタックに詰む値
 * @return pVal の条件判定結果。 lns_isCondTrue()。
 */
bool lns_setStackVal( lns_env_t * _pEnv, lns_stem_t val )
{
    _pEnv->pValStack[ _pEnv->stackPos ] = val;
    return lns_isCondTrue( val );
}

/**
 * スタックから値を pop する。
 *
 * @return pop した値。
 */
lns_stem_t lns_popVal( lns_env_t * _pEnv, bool dummy ) {
    lns_stem_t val = _pEnv->pValStack[ _pEnv->stackPos ];
    _pEnv->stackPos--;
    return val;
}

/**
 * int! の unwrap 処理
 */
lns_stem_t lns_unwrap_stem( lns_stem_t stem, lns_stem_t defVal )
{
    if ( stem.type != lns_stem_type_nil ) {
        return stem;
    }
    if ( defVal.type != lns_stem_type_none && defVal.type != lns_stem_type_nil ) {
        return defVal;
    }
    lns_abort( __func__ );
    return lns_global.noneStem; // dummy
}

lns_any_t * lns_unwrap_any( lns_stem_t stem, lns_stem_t defVal )
{
    if ( stem.type != lns_stem_type_nil ) {
        return stem.val.pAny;
    }
    if ( defVal.type != lns_stem_type_none && defVal.type != lns_stem_type_nil ) {
        return defVal.val.pAny;
    }
    lns_abort( __func__ );
    return NULL; // dummy
}



/**
 * int! の unwrap 処理
 */
lns_int_t lns_unwrap_int( lns_stem_t stem )
{
    if ( stem.type == lns_stem_type_int ) {
        return stem.val.intVal;
    }
    lns_abort( __func__ );
    return 0; // dummy
}

/**
 * int! の unwrap 処理。 default 値付き。
 */
lns_int_t lns_unwrap_intDefault( lns_stem_t stem, lns_int_t val )
{
    if ( stem.type == lns_stem_type_int ) {
        return stem.val.intVal;
    }
    return val;
}


/**
 * real! の unwrap 処理
 */
lns_real_t lns_unwrap_real( lns_stem_t stem )
{
    if ( stem.type == lns_stem_type_real ) {
        return stem.val.realVal;
    }
    lns_abort( __func__ );
    return 0.0;
}

/**
 * real! の unwrap 処理。 default 値付き。
 */
lns_real_t lns_unwrap_realDefault( lns_stem_t stem, lns_real_t val )
{
    if ( stem.type == lns_stem_type_real ) {
        return stem.val.realVal;
    }
    return val;
}


lns_stem_t lns_stem_refAt( lns_env_t * _pEnv, lns_stem_t stem, lns_stem_t key )
{
    if ( stem.type == lns_stem_type_any &&
         stem.val.pAny->type == lns_value_type_class &&
         stem.val.pAny->val.classVal->pMeta == &lns_type_meta_Map )
    {
        return lns_mtd_Map_get( _pEnv, stem.val.pAny, key );
    }
    return lns_global.nilStem;
}


/**
 * lua の print() に相当する処理。
 *
 * @param pArg は ddd 型。
 */
void lns_f_print( lns_env_t * _pEnv, lns_stem_t ddd ) {

    //lns_enter_func( _pEnv, 0, 0, 1, pArg );
    //lns_enter_block( _pEnv, 0, 0 );

    lns_any_t * pArg = ddd.val.pAny;

    int index;
    for ( index = 0; index < lns_lenDDD( pArg ); index++ ) {
        if ( index > 0 ) {
            printf( "\t" );
        }
        const lns_stem_t stem = lns_fromDDD( pArg, index );
        switch ( stem.type ) {
        case lns_stem_type_nil:
            printf( "nil" );
            break;
        case lns_stem_type_bool:
            printf( stem.val.boolVal ? "true" : "false" );
            break;
        case lns_stem_type_int:
            printf( "%d", stem.val.intVal );
            break;
        case lns_stem_type_real:
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
            
        case lns_stem_type_any:
            {
                lns_any_t * pAny = stem.val.pAny;
                switch ( pAny->type ) {
                case lns_value_type_str:
                    printf( "%s", pAny->val.str.pStr );
                    break;
                case lns_value_type_form:
                    printf( "form: %p", pAny );
                    break;
                case lns_value_type_class:
                    printf( "class: %p(%s)", pAny, pAny->val.classVal->pMeta->pName );
                    break;
                case lns_value_type_if:
                    printf( "if: %p(%s): %p(%s)",
                            pAny, pAny->val.ifVal.pMeta->pName,
                            pAny->val.ifVal.pObj,
                            pAny->val.ifVal.pObj->val.classVal->pMeta->pName );
                    break;
                case lns_value_type_luaVal:
                    switch( pAny->val.luaVal.type ) {
                    case lns_value_type_none:
                        printf( "lua: none %p", pAny );
                        break;
                    case lns_value_type_str:
                        printf( "lua: str %p", pAny );
                        break;
                    case lns_value_type_class:
                        printf( "lua: class %p", pAny );
                        break;
                    case lns_value_type_if:
                        printf( "lua: if %p", pAny );
                        break;
                    case lns_value_type_ddd:
                        printf( "lua: ddd %p", pAny );
                        break;
                    case lns_value_type_mRet:
                        printf( "lua: mRet %p", pAny );
                        break;
                    case lns_value_type_form:
                        printf( "lua: form %p", pAny );
                        break;
                    /* case lns_value_type_List: */
                    /*     printf( "lua: List %p", pAny ); */
                    /*     break; */
                    /* case lns_value_type_Array: */
                    /*     printf( "lua: Array %p", pAny ); */
                    /*     break; */
                    /* case lns_value_type_Set: */
                    /*     printf( "lua: Set %p", pAny ); */
                    /*     break; */
                    /* case lns_value_type_Map: */
                    /*     printf( "lua: Map %p", pAny ); */
                    /*     break; */
                    case lns_value_type_itList:
                        printf( "lua: itList %p", pAny );
                        break;
                    case lns_value_type_itSet:
                        printf( "lua: itSet %p", pAny );
                        break;
                    case lns_value_type_itMap:
                        printf( "lua: itMap %p", pAny );
                        break;
                    case lns_value_type_alge:
                        printf( "lua: alge %p", pAny );
                        break;
                    case lns_value_type_luaVal:
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

    //lns_leave_block( _pEnv );
}

int lua_main( lua_State * pLua ) {
    int argc = lua_tointeger( pLua, 1);
    char ** pArgv = (char **)lua_touserdata( pLua, 2);
    int script;

    lns_collection_init();
    
    lns_createGlobalEnv();
    
    lns_env_t * _pEnv = lns_createEnv( pLua );

    lns_init_lns_builtin( _pEnv );

    // エントリモジュールを実行する
    lns_run_module( _pEnv );

    lns_deleteEnv( _pEnv );

    if ( lns_releaseGlobalEnv() ) {
        lua_pushboolean( pLua, 1);
    }
    else {
        lua_pushboolean( pLua, 0);
    }
  
    return 1;
}
