#include <gc_lunescript.h>
#include <lauxlib.h>

#define LNS_GC_SET_BIT( ANY, BIT ) (ANY)->state = (ANY)->state | BIT;
#define LNS_GC_CLEAR_BIT( ANY, BIT ) (ANY)->state = (ANY)->state & ~(BIT);

/** minor から major に promote させるための age */
#define LNS_GC_AGE_TO_PROMOTE 2 

/**
 * リストから ANY を除外。
 */
#define lns_rmFromList( ANY )                  \
    if ( (ANY)->pNext != NULL ) {               \
        (ANY)->pPrev->pNext = (ANY)->pNext;     \
        (ANY)->pNext->pPrev = (ANY)->pPrev;     \
        (ANY)->pNext = NULL;                    \
    }

/**
 * リストの初期化
 */
#define lns_initList( ITEM, TOP )              \
    (ITEM)->pNext = TOP;                       \
    (ITEM)->pPrev = TOP;



#define lns_alloc_any( ENV, TYPE, FILE, LINE ) \
    _lns_alloc_any( ENV, TYPE, FILE, LINE )
#define lns_alloc_any_op( ENV, TYPE )                  \
    _lns_alloc_any( ENV, TYPE, __FILE__, __LINE__ )



    
typedef struct {
    /** 確保したオブジェクトのポインタを保持する */
    lns_any_t * pPool[ LNS_GC_MINOR_MAX ];
    /** pPool に格納しているオブジェクトの数 */
    int count;
} lns_newpool_t;


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
static lns_any_t s_majorObjTop;


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


lns_type_meta_t lns_type_meta_lns__root = {
    "_root",
    &lns_type_meta_lns__root,
    { NULL }
};


lns_type_meta_t lns_type_meta_List = { "List", &lns_type_meta_lns__root, { NULL } };
lns_type_meta_t lns_type_meta_Set = { "Set", &lns_type_meta_lns__root, { NULL } };
lns_type_meta_t lns_type_meta_Map = { "Map", &lns_type_meta_lns__root, { NULL } };


static void lns_gc_collect_minor( lns_env_t * _pEnv );


void lns_releaeAnyVal( lns_env_t * _pEnv, void * pKey ) { /* dummy */ }
lns_stem_t lns_getValFromDDD( lns_any_t * pAny, int index ) { /* dummy */ }
void lns_setQ_( lns_any_t * pAny ) { /* dummy */}
void lns_setRet( lns_env_t * _pEnv, lns_stem_t stem ) {}
lns_stem_t _lns_createDDDOnly(
    const char * pFile, int lineNo, lns_env_t * _pEnv, int num ) {}



void init( void ) {
    int index;
    s_pNewPool = &s_poolBuf[ 0 ];
    s_pWorkPool = &s_poolBuf[ 1 ];
    s_pNewPool->count = 0;
    s_pWorkPool->count = 0;

    lns_initList( &s_freeVarLinkTop, &s_freeVarLinkTop );
    lns_initList( &s_minorVarLinkTop, &s_freeVarLinkTop );
    lns_initList( &s_majorObjTop.major, &s_majorObjTop );
    
    for ( index = 0; index < LNS_GC_MINOR_MAX; index++ ) {
        lns_varLink_t * pVarLink = &s_varLinkBuf[ index ];
        lns_add2list( &s_freeVarLinkTop, pVarLink );
    }
}

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

    pAny->state = LNS_GC_STATE_MINOR;

    if ( s_pNewPool->count >= LNS_GC_MINOR_MAX ) {
        lns_gc_collect_minor( _pEnv );
    }
    s_pNewPool->pPool[ s_pNewPool->count ] = pAny;
    s_pNewPool->count++;


    
    /* // 現在のブロックに登録 */
    /* lns_block_t * pBlock = &_pEnv->blockQueue[ _pEnv->blockDepth ]; */
    /* lns_add2list( &pBlock->managedAnyTop, pAny ); */

    return pAny;
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
 * pAny を開放する。
 *
 * 参照カウントが 0 になっていることが前提条件。
 */
static void lns_release_any( lns_env_t * _pEnv, lns_any_t * pAny ) {
    pAny->type = lns_value_type_none; // このセットは必須ではないが、
    // セットしておいた方が不具合を見つけ易い。
    lns_free( _pEnv->allocateor, pAny );
    s_globalEnv.allocNum--;
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
                lns_var_t * pVar = l_form_closure( pAny, index );
                if ( pVar->pLink ) {
                    lns_unuseVar( pVar );
                }
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


lns_any_t * _lns_litStr2any(
    const char * pFile, int lineNo, lns_env_t * _pEnv, const char * pStr )
{
    lns_any_t * pAny =
        lns_alloc_any( _pEnv, lns_value_type_str, pFile, lineNo );
    pAny->val.str.len = strlen( pStr );
    pAny->val.str.pStr = pStr;
    pAny->val.str.staticFlag = true;

    lns_pushRetAny( _pEnv, pAny );
    
    return pAny;
}

/**
pVar で参照しているオブジェクトを minor から major に変更する。
ただし、 pVar で使用していた varlink は、空きリストに戻さない。
 */
static void lns_promoteObjOnly( lns_var_t * pVar )
{
    // varlink から除去
    lns_rmFromList( pVar->pLink );
    pVar->pLink = NULL;

    lns_any_t * pAny = pVar->stem.val.pAny;
    // major オブジェクトのリンクに追加
    lns_add2list( &s_majorObjTop, pAny );
    // major にセット
    LNS_GC_CLEAR_BIT( pAny, LNS_GC_STATE_MINOR );
}

/**
pVar で参照しているオブジェクトを minor から major に変更する。

@param pVar minor オブジェクトを保持する var
 */
void lns_promoteObj( lns_var_t * pVar )
{
    lns_varLink_t * pVarLink = pVar->pLink;
    lns_promoteObjOnly( pVar );

    // varlink の空きリストに追加
    lns_add2list( &s_freeVarLinkTop, pVarLink );
}

lns_varLink_t * lns_newVarLink( lns_var_t * pVar )
{
    lns_varLink_t * pVarLink = s_freeVarLinkTop.pNext;
    if ( pVarLink == &s_freeVarLinkTop ) {
        // 空きがない場合、一番古いオブジェクトを promote
        pVarLink = s_freeVarLinkTop.pPrev;
        lns_promoteObjOnly( pVarLink->pVar );
    }
    else {
        // 空きリストから除外
        lns_rmFromList( pVarLink );
    }
    // varLink を初期化
    pVarLink->age = 0;
    lns_add2list( &s_minorVarLinkTop, pVarLink );


    pVarLink->pVar = pVar;
    pVar->pLink = pVarLink;
    
    return pVarLink;
}

void lns_unuseVar( lns_var_t * pVar )
{
    lns_varLink_t * pLink = pVar->pLink;
    if ( pLink != NULL ) {
        lns_rmFromList( pLink );
        // 空き varlink に追加
        lns_add2list( &s_freeVarLinkTop, pLink );
    }
}

void lns_initVar( lns_var_t * pVar, lns_stem_t stem )
{
    pVar->stem = stem;
    if ( pVar->stem.type == lns_stem_type_any ) {
        if ( LNS_IS_GC_MINOR( pVar->stem.val.pAny ) ) {
            lns_newVarLink( pVar );
        }
    }
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

    lns_initList( &_pEnv->loadModuleTop, &_pEnv->loadModuleTop );

    _pEnv->pMRet = NULL;

    _pEnv->pSortCallback = lns_global.nilStem;

    lns_initList( &_pEnv->retObjTop, &_pEnv->retObjTop );
    
    if ( pLua != NULL ) {
        //lns_setLuaWapper( _pEnv );
    }
    
    return _pEnv;
}

/**
 * 新しくブロックを開始する。
 *
 * @return ブロックインデックス
 */
int32_t lns_enter_block( lns_env_t * _pEnv )
{
    lns_var_t * pVar = lns_allocVar( _pEnv, lns_stem_type_stack );
    int32_t ret = _pEnv->useStackVarNum - 1;


    lns_any_t * pAny = &_pEnv->blockAnyBuf[ _pEnv->blockDepth ];
    pAny->type = lns_value_type_block;
    pAny->pNext = _pEnv->retObjTop.pNext;
    _pEnv->retObjTop.pNext = pAny;

    _pEnv->blockDepth++;
                  
    return ret;
}


/**
 * 現在のブロックを終了する。
 *
 * blockIndex 以降の var の参照を外す
 */
void lns_leave_block( lns_env_t * _pEnv, int32_t blockIndex )
{
    _pEnv->blockDepth--;

    lns_any_t * pAny = &_pEnv->blockAnyBuf[ _pEnv->blockDepth ];
    _pEnv->retObjTop.pNext = pAny->pNext;


    
    _pEnv->useStackVarNum--;
    for ( ; _pEnv->useStackVarNum >= blockIndex; _pEnv->useStackVarNum-- ) {
        lns_unuseVar( &_pEnv->stackVarBuf[ _pEnv->useStackVarNum ] );
    }
    _pEnv->useStackVarNum++;
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
static bool lns_deleteEnv( lns_env_t * _pEnv )
{
    lns_module_t * pModule = _pEnv->loadModuleTop.pNext;
    for ( ; pModule != &_pEnv->loadModuleTop; pModule = pModule->pNext )
    {
        //lns_leave_blockSub( _pEnv, pModule->pBlock );
        // lns_leave_blockSub( s_globalEnv.pEnv, pModule->pBlock );
    }

    lns_gc_collect_minor( _pEnv );

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

    lns_initList( &pBlock->managedAnyTop, &pBlock->managedAnyTop );

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

static void lns_createGlobalEnv() {
    s_globalEnv.allocNum = 0;
    s_globalEnv.pEnv = lns_createEnv( NULL );
    s_globalEnv.moduleNum = 0;
    
    lns_global.noneStem = LNS_STEM_BASE( lns_stem_type_none );
    lns_global.nilStem = LNS_STEM_BASE( lns_stem_type_nil );
    lns_global.trueStem = LNS_STEM_BOOL( true );
    lns_global.falseStem = LNS_STEM_BOOL( false );

    lns_global.nilVar.pLink = NULL;
    lns_global.nilVar.stem = lns_global.nilStem;
}

static bool lns_releaseGlobalEnv(void) {
    /* int index; */
    /* for ( index = s_globalEnv.moduleNum - 1; index >= 0; index-- ) { */
    /*     lns_leave_blockSub( s_globalEnv.pEnv, */
    /*                          &s_globalEnv.moduleInitBlockBuf[ index ] ); */
    /* } */
    //lns_decre_ref( s_globalEnv.pEnv, lns_global.ddd0.val.pAny );
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


static void lns_gc_collect_minor( lns_env_t * _pEnv )
{
    int index;
    // newpool のオブジェクトの REFED をクリア
    for ( index = 0; index < s_pNewPool->count; index++ ) {
        LNS_GC_CLEAR_BIT( s_pNewPool->pPool[ index ], LNS_GC_STATE_REFED );
    }

    // newvar で参照している全オブジェクトについて以下を実行
    lns_varLink_t * pLink = s_minorVarLinkTop.pNext;
    lns_varLink_t * pNext;
    for ( ; pLink != &s_minorVarLinkTop; pLink = pNext ) {
        // リンク情報が変わる可能性があるので、ここで次を取っておく
        pNext = pLink->pNext;
        
        lns_var_t * pVar = pLink->pVar;
        lns_any_t * pAny = pVar->stem.val.pAny;
        if ( !LNS_IS_GC_MINOR( pAny ) ) {
            // オブジェクトが MAJOR だった場合、 newvar から除外する
            lns_rmFromList( pLink );
            // 空き varlink に追加
            lns_add2list( &s_freeVarLinkTop, pLink );
            continue;
        }
        pLink->age++;
        if ( pLink->age > LNS_GC_AGE_TO_PROMOTE ) {
            // age をインクリメントし所定の値より大きくなった場合、 promote
            lns_promoteObj( pVar );
        }
        else {
            // 所定の値以下の場合、 REFED をセット
            LNS_GC_SET_BIT( pAny, LNS_GC_STATE_REFED );

            if ( pAny->type == lns_value_type_if ) {
                // interface の場合、元のオブジェクトも REFED をセットする
                lns_any_t * pObjForIF = pAny->val.ifVal.pObj;
                if ( LNS_IS_GC_MINOR( pObjForIF ) ) {
                    LNS_GC_SET_BIT( pObjForIF, LNS_GC_STATE_REFED );
                }
            }
        }
    }
    // ret に登録されているオブジェクトについて実行
    lns_any_t * pRet = _pEnv->retObjTop.pNext;
    while ( pRet != NULL ) {
        if ( LNS_IS_GC_MINOR( pObjForIF ) ) {
            LNS_GC_SET_BIT( pObjForIF, LNS_GC_STATE_REFED );
        }
    }

    
    // newpool のオブジェクトで REFED がセットされていないオブジェクトを破棄。
    // REFED がセットされている場合 workpool にコピー
    for ( index = 0; index < s_pNewPool->count; index++ ) {
        lns_any_t * pAny = s_pNewPool->pPool[ index ];
        if ( !LNS_IS_GC_REFED( pAny ) ) {
            lns_gc_any( _pEnv, pAny, true );
        }
        else {
            s_pWorkPool->pPool[ s_pWorkPool->count ] = pAny;
            s_pWorkPool->count++;
        }
    }
    if ( s_pWorkPool->count == LNS_GC_MINOR_MAX ) {
        // 全 minor オブジェクトが参照されている場合、
        
    }
    
    // newpool を workpool で置き換え
    lns_newpool_t * pTemp = s_pNewPool;
    s_pNewPool = s_pWorkPool;
    s_pWorkPool = pTemp;
}


void test() {
    lua_State * pLua = NULL;

    lns_createGlobalEnv();
    lns_env_t * _pEnv = lns_createEnv( pLua );

    int _block = lns_enter_block( _pEnv );
    lns_var_t * val1 = lns_allocVar( _pEnv, lns_stem_type_any );
    val1->stem.val.pAny = lns_litStr2any( _pEnv, "a" );
    

    lns_leave_block( _pEnv, _block );
    
    lns_deleteEnv( _pEnv );

    lns_releaseGlobalEnv();
}

int main( int argc, const char * pArgv[] ) {
    init();
    test();
    return 0;
}
