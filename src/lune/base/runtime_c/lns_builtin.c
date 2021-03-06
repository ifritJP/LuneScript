// lune/base/runtime_c/lns_builtin.c
#include <lunescript.h>
#include <lns_builtin.h>
static lns_module_t lns_moduleInfo_lns_builtin = {NULL,NULL,false, NULL, "lns_builtin", {NULL } };
static lns_any_t ** lns_module_globalStemList;
static lns_any_t ** lns_module_path = NULL;
// decl class iStream (ProcessMode.Prototype)-->
// <--- decl class iStream (ProcessMode.Prototype)
// decl class oStream (ProcessMode.Prototype)-->
// <--- decl class oStream (ProcessMode.Prototype)
// decl class luaStream (ProcessMode.Prototype)-->
static void mtd_lns_luaStream__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_luaStream__delExt( lns_env_t * _pEnv, lns_any_t * pObj );
// <--- decl class luaStream (ProcessMode.Prototype)
// decl class Mapping (ProcessMode.Prototype)-->
// <--- decl class Mapping (ProcessMode.Prototype)
static void mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj);
static void mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj);
static lns_stem_t mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1);
static lns_stem_t mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2);
static lns_stem_t mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1);
lns_any_t * lns_io_stderr;
lns_any_t * lns_io_stdin;
lns_any_t * lns_io_stdout;
lns_any_t * lns_package_path;
// decl class iStream (ProcessMode.WideScopeVer)-->
lns_type_meta_t lns_type_meta_lns_iStream = { "lns_iStream", &lns_type_meta_lns__root, {NULL } };
// <--- decl class iStream (ProcessMode.WideScopeVer)
// decl class oStream (ProcessMode.WideScopeVer)-->
lns_type_meta_t lns_type_meta_lns_oStream = { "lns_oStream", &lns_type_meta_lns__root, {NULL } };
// <--- decl class oStream (ProcessMode.WideScopeVer)
// decl class luaStream (ProcessMode.WideScopeVer)-->
static lns_mtd_lns_iStream_t lns_if_lns_luaStream_imp_lns_iStream = {
    (lns_method_t *)mtd_lns_luaStream_close,
    (lns_method_t *)mtd_lns_luaStream_read,
};
static lns_mtd_lns_oStream_t lns_if_lns_luaStream_imp_lns_oStream = {
    (lns_method_t *)mtd_lns_luaStream_close,
    (lns_method_t *)mtd_lns_luaStream_flush,
    (lns_method_t *)mtd_lns_luaStream_write,
};
lns_type_meta_t lns_type_meta_lns_luaStream = { "lns_luaStream", &lns_type_meta_lns__root, {&lns_type_meta_lns_iStream, &lns_type_meta_lns_oStream, NULL } };
lns_mtd_lns_luaStream_t lns_mtd_lns_luaStream = {
    mtd_lns_luaStream__del,
    NULL,
    (lns_method_t *)mtd_lns_luaStream_close,
    (lns_method_t *)mtd_lns_luaStream_flush,
    (lns_method_t *)mtd_lns_luaStream_read,
    (lns_method_t *)mtd_lns_luaStream_seek,
    (lns_method_t *)mtd_lns_luaStream_write,
};
// <--- decl class luaStream (ProcessMode.WideScopeVer)
// decl class Mapping (ProcessMode.WideScopeVer)-->
lns_type_meta_t lns_type_meta_lns_Mapping = { "lns_Mapping", &lns_type_meta_lns__root, {NULL } };
// <--- decl class Mapping (ProcessMode.WideScopeVer)
// decl class iStream (ProcessMode.DefClass)-->
// <--- decl class iStream (ProcessMode.DefClass)
// decl class oStream (ProcessMode.DefClass)-->
// <--- decl class oStream (ProcessMode.DefClass)
// decl class luaStream (ProcessMode.DefClass)-->
static void mtd_lns_luaStream__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
    mtd_lns_luaStream__delExt( _pEnv, pObj );
}
lns_any_t * lns_class_lns_luaStream_new( lns_env_t * _pEnv){ // 0
    lns_class_new_( _pEnv, lns_luaStream, pAny, pObj );
    pObj->pExt = NULL;
    pObj->pImp = &pObj->imp;
    pObj->imp.sentinel.type = lns_value_type_none;
    lns_init_if( &pObj->imp.lns_iStream, _pEnv, pAny, &lns_if_lns_luaStream_imp_lns_iStream, lns_iStream );
    lns_init_if( &pObj->imp.lns_oStream, _pEnv, pAny, &lns_if_lns_luaStream_imp_lns_oStream, lns_oStream );
    return pAny;
}
// <--- decl class luaStream (ProcessMode.DefClass)
// decl class Mapping (ProcessMode.DefClass)-->
// <--- decl class Mapping (ProcessMode.DefClass)
static void initFuncSym( lns_env_t * _pEnv, lns_block_t * pBlock )
{
}
void l_call_mtd_lns_iStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_iStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
void l_nil_mtd_lns_iStream_close( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_iStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
lns_stem_t l_call_mtd_lns_iStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1){
return lns_mtd_lns_iStream( pObj )->read( _pEnv, lns_getImpObj( pObj ) , arg1);
}
lns_stem_t l_nil_mtd_lns_iStream_read( lns_env_t * _pEnv, lns_stem_t obj, lns_stem_t arg1){
if ( obj.type == lns_stem_type_nil ) { return lns_global.nilStem; }
lns_any_t * pObj = obj.val.pAny;
return lns_mtd_lns_iStream( pObj )->read( _pEnv, lns_getImpObj( pObj ) , arg1);
}
void l_call_mtd_lns_oStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_oStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
void l_nil_mtd_lns_oStream_close( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_oStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
void l_call_mtd_lns_oStream_flush( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_oStream( pObj )->flush( _pEnv, lns_getImpObj( pObj ) );
}
void l_nil_mtd_lns_oStream_flush( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_oStream( pObj )->flush( _pEnv, lns_getImpObj( pObj ) );
}
lns_stem_t l_call_mtd_lns_oStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1){
return lns_mtd_lns_oStream( pObj )->write( _pEnv, lns_getImpObj( pObj ) , arg1);
}
lns_stem_t l_nil_mtd_lns_oStream_write( lns_env_t * _pEnv, lns_stem_t obj, lns_any_t * arg1){
if ( obj.type == lns_stem_type_nil ) { return lns_global.ddd0; }
lns_any_t * pObj = obj.val.pAny;
return lns_mtd_lns_oStream( pObj )->write( _pEnv, lns_getImpObj( pObj ) , arg1);
}
void l_call_mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_luaStream( pObj )->close( _pEnv, pObj );
}
void l_nil_mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_luaStream( pObj )->close( _pEnv, pObj );
}
void l_call_mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_luaStream( pObj )->flush( _pEnv, pObj );
}
void l_nil_mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_luaStream( pObj )->flush( _pEnv, pObj );
}
lns_stem_t l_call_mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1){
return lns_mtd_lns_luaStream( pObj )->read( _pEnv, pObj , arg1);
}
lns_stem_t l_nil_mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_stem_t obj, lns_stem_t arg1){
if ( obj.type == lns_stem_type_nil ) { return lns_global.nilStem; }
lns_any_t * pObj = obj.val.pAny;
return lns_mtd_lns_luaStream( pObj )->read( _pEnv, pObj , arg1);
}
lns_stem_t l_call_mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2){
return lns_mtd_lns_luaStream( pObj )->seek( _pEnv, pObj , arg1, arg2);
}
lns_stem_t l_nil_mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_stem_t obj, lns_any_t * arg1, lns_int_t arg2){
if ( obj.type == lns_stem_type_nil ) { return lns_global.ddd0; }
lns_any_t * pObj = obj.val.pAny;
return lns_mtd_lns_luaStream( pObj )->seek( _pEnv, pObj , arg1, arg2);
}
lns_stem_t l_call_mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1){
return lns_mtd_lns_luaStream( pObj )->write( _pEnv, pObj , arg1);
}
lns_stem_t l_nil_mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_stem_t obj, lns_any_t * arg1){
if ( obj.type == lns_stem_type_nil ) { return lns_global.ddd0; }
lns_any_t * pObj = obj.val.pAny;
return lns_mtd_lns_luaStream( pObj )->write( _pEnv, pObj , arg1);
}
void l_call_mtd_lns_Mapping__toMap( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_Mapping( pObj )->_toMap( _pEnv, lns_getImpObj( pObj ) );
}
void l_nil_mtd_lns_Mapping__toMap( lns_env_t * _pEnv, lns_stem_t obj){
if ( obj.type == lns_stem_type_nil ) { return ; }
lns_any_t * pObj = obj.val.pAny;
lns_mtd_lns_Mapping( pObj )->_toMap( _pEnv, lns_getImpObj( pObj ) );
}
static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv );
lns_module_t * lns_init_lns_builtin( lns_env_t * _pEnv ){
    if ( lns_moduleInfo_lns_builtin.readyFlag ) {
        return &lns_moduleInfo_lns_builtin;
    }
    lns_moduleInfo_lns_builtin.readyFlag = true;
    lns_add2list( &_pEnv->loadModuleTop, &lns_moduleInfo_lns_builtin);
    
    lns_block_t * pBlock_62 = lns_enter_module( _pEnv, 2, 0, 0 );
    lns_moduleInfo_lns_builtin.pBlock = pBlock_62;
    lns_set_block_any( pBlock_62, 0, lns_module_globalStemList);
    lns_setQ_any( lns_module_globalStemList, lns_class_List_new( _pEnv ));
    lns_set_block_any( pBlock_62, 1, lns_module_path);
    lns_setQ_any( lns_module_path, lns_litStr2any( _pEnv, "@lns_builtin"));
    lns_enter_block( _pEnv, 0, 0, 0 );
    initFuncSym( _pEnv, pBlock_62 );
    
    lns_init_lns_builtin_Sub( _pEnv );
