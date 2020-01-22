#include <gc_lunescript.h>
#include <lauxlib.h>


/**
 * リストから ANY を除外。
 */
#define lns_rmFromList( ANY )                  \
    if ( (ANY)->pNext != NULL ) {               \
        (ANY)->pPrev->pNext = (ANY)->pNext;     \
        (ANY)->pNext->pPrev = (ANY)->pPrev;     \
        (ANY)->pNext = NULL;                    \
    }

#define lns_malloc( ALLOCATOR, SIZE ) _lns_malloc( ALLOCATOR, SIZE, __FILE__, __LINE__ )
#define lns_free( ALLOCATOR, ADDR ) _lns_free( ALLOCATOR, ADDR, __FILE__, __LINE__ )

#define lns_alloc_any( ENV, TYPE, FILE, LINE ) \
    _lns_alloc_any( ENV, TYPE, FILE, LINE )
#define lns_alloc_any_op( ENV, TYPE )                  \
    _lns_alloc_any( ENV, TYPE, __FILE__, __LINE__ )



static lns_newpool_t * s_pNewPool;
static lns_newpool_t * s_pWorkPool;
/** s_pNewPool, s_pWorkPool の実体 */
static lns_newpool_t s_poolBuf[ 2 ];
static lns_varLink_t s_varLinkBuf[ LNS_GC_MINOR_MAX ];

/** 空き varlink の先頭を管理。 */
static lns_varLink_t s_freeVarLinkTop;
/** minor オブジェクトを参照する varlink の先頭を管理。 */
static lns_varLink_t s_minorVarLinkTop;
/** major オブジェクトの先頭を管理 */
static lns_varLink_t s_majorObjTop;


typedef struct lns_globalEnv_t {
    lns_env_t * pEnv;
    /** モジュールの init ブロックのバッファ */
    lns_block_t moduleInitBlockBuf[ LNS_MODULE_MAX_NUM ];
    /** 現在のモジュール数 */
    int moduleNum;
    /** 現在確保している any の数 */
    int allocNum;
} lns_globalEnv_t;

static lns_globalEnv_t s_globalEnv;
lns_global_t lns_global;


void init( void ) {
    int index;
    s_pNewPool = &s_poolBuf[ 0 ];
    s_pWorkPool = &s_poolBuf[ 1 ];
    s_pNewPool->count = 0;
    s_pWorkPool->count = 0;

    s_freeVarLinkTop.pNext = &s_freeVarLinkTop;
    s_freeVarLinkTop.pPrev = &s_freeVarLinkTop;

    s_minorVarLinkTop.pNext = &s_minorVarLinkTop;
    s_minorVarLinkTop.pPrev = &s_minorVarLinkTop;
    
    for ( index = 0; index < LNS_GC_MINOR_MAX; index++ ) {
        lns_varLink_t * pVarLink = &s_varLinkBuf[ index ];
        lns_add2list( &s_freeVarLinkTop, pVarLink );
    }

    s_majorObjTop.pNext = &s_majorObjTop;
    s_majorObjTop.pPrev = &s_majorObjTop;
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


/**
pVar で参照しているオブジェクトを minor から major に変更する。
 */
void lns_promoteObj( lns_env_t * _pEnv, lns_var_t * pVar )
{
    // varlink から除去
    lns_rmFromList( pVar->pLink );

    // major オブジェクトのリンクに追加
    lns_add2list( &s_majorObjTop, pVar->pLink );
}

lns_var_t * lns_allocVar( lns_env_t * _pEnv, lns_stem_type_t type )
{
    lns_varLink_t * pVarLink = s_freeVarLinkTop.pNext;
    if ( pVarLink == &s_freeVarLinkTop ) {
        // 空きがない場合、一番古いオブジェクトを promote
        pVarLink = s_freeVarLinkTop.pPrev;
        lns_promoteObj( _pEnv, pVarLink->pVar );
    }
    else {
        // 空きリストから除外
        lns_rmFromList( pVarLink );
    }
    // varLink を初期化
    pVarLink->age = 0;
    lns_add2list( &s_minorVarLinkTop, pVarLink );

    lns_var_t * pVar = &_pEnv->stackVarBuf[ _pEnv->useStackVarNum ];
    pVarLink->pVar = pVar;
    pVar->pLink = pVarLink;
    _pEnv->useStackVarNum++;

    return pVarLink->pVar;
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

    _pEnv->useStackVarNum = 0;
    
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
        //lns_setLuaWapper( _pEnv );
    }
    
    return _pEnv;
}

static inline void lns_reset_blockSub( lns_env_t * _pEnv, lns_block_t * pBlock ) {
    /* lns_any_t * pWork = pBlock->managedAnyTop.pPrev; */
    /* while ( pWork != &pBlock->managedAnyTop ) { */
    /*     lns_any_t * pPrev = pWork->pPrev; */

    /*     lns_lock(  */
    /*         if ( pWork->refCount == 1 ) { */
    /*             lns_gc_any( _pEnv, pWork, true ); */
    /*         } */
    /*         else { */
    /*             pWork->refCount--; */
    /*             // オブジェクトが残る場合は、チェーンから除外する */
    /*             pWork->pNext = NULL; */
    /*         } */
    /*     ); */
    /*     pWork = pPrev; */
    /* } */
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
    /* int index; */
    /* for ( index = pBlock->varLen - 1; index >= 0; index-- ) { */
    /*     lns_closureVar_t * pVar = pBlock->pVarList[ index ]; */
    /*     if ( pVar != NULL ) { */
    /*         lns_var_decre( _pEnv, pVar ); */
    /*     } */
    /* } */
    /* for ( index = pBlock->stemLen - 1; index >= 0; index-- ) { */
    /*     lns_stem_t * pStem = pBlock->pStemList[ index ]; */
    /*     if ( pStem != NULL && pStem->type == lns_stem_type_any ) { */
    /*         if ( lns_decre_ref( _pEnv, pStem->val.pAny ) ) { */
    /*             pStem->type = lns_stem_type_none; */
    /*             pBlock->pStemList[ index ] = NULL; */
    /*         } */
    /*     } */
    /* } */
    /* for ( index = pBlock->anyLen - 1; index >= 0; index-- ) { */
    /*     lns_any_t * pAny = pBlock->pAnyList[ index ]; */
    /*     if ( pAny != NULL ) { */
    /*         lns_decre_ref( _pEnv, pAny ); */
    /*         pBlock->pAnyList[ index ] = NULL; */
    /*     } */
    /* } */

    /* lns_reset_blockSub( _pEnv, pBlock ); */
    
    /* _pEnv->useAnyPoolNum -= pBlock->anyLen; */
    /* _pEnv->useStemPoolNum -= pBlock->stemLen; */
    /* _pEnv->useVarPoolNum -= pBlock->varLen; */
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
    DEBUG_MEM_LOG( useStackVarNum );
    
    lns_free( _pEnv->allocateor, _pEnv );

    return result;
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
lns_block_t * lns_enter_block(
    lns_env_t * _pEnv, int anyNum, int stemNum, int varNum )
{
    _pEnv->blockDepth++;
    
    lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ];
    lns_setup_block( _pEnv, pBlock, anyNum, stemNum, varNum );

    return pBlock;
}


void lns_gc_collect_minor( lns_env_t * _pEnv )
{
    
}


void test() {
    lua_State * pLua = NULL;
    
    lns_env_t * _pEnv = lns_createEnv( pLua );
    
    lns_var_t * val1 = lns_allocVar( _pEnv, lns_stem_type_any );
    val1->stem.val.pAny = lns_litStr2any( _pEnv, "a" );
    

    lns_deleteEnv( _pEnv );

    printf( ":debug: allocNum = %d\n", s_globalEnv.allocNum );
}

int main( int argc, const char * pArgv[] ) {
    init();
    test();
    return 0;
}
