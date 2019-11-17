// lune/base/runtime_c/lns_builtin.c
#include <lunescript.h>
#include <lns_builtin.h>
static lns_module_t s_module = {NULL,NULL,false};
static lns_any_t ** lns_module_globalStemList;
static lns_any_t ** lns_module_path = NULL;
static void mtd_lns_luaStream__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_io__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_package__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_os__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_string__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_math__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_debug__del( lns_env_t * _pEnv, lns_any_t * pObj );
static void mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj);
static void mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj);
static lns_stem_t mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1);
static lns_stem_t mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2);
static lns_stem_t mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1);
lns_type_meta_t lns_type_meta_lns_iStream = { "lns_iStream" };
lns_type_meta_t lns_type_meta_lns_oStream = { "lns_oStream" };
static lns_mtd_lns_iStream_t lns_if_lns_luaStream_imp_lns_iStream = {
   (lns_method_t *)mtd_lns_luaStream_close,
   (lns_method_t *)mtd_lns_luaStream_read,
};
static lns_mtd_lns_oStream_t lns_if_lns_luaStream_imp_lns_oStream = {
   (lns_method_t *)mtd_lns_luaStream_close,
   (lns_method_t *)mtd_lns_luaStream_flush,
   (lns_method_t *)mtd_lns_luaStream_write,
};
lns_type_meta_t lns_type_meta_lns_luaStream = { "lns_luaStream" };
lns_mtd_lns_luaStream_t lns_mtd_lns_luaStream = {
   mtd_lns_luaStream__del,
   NULL,
   (lns_method_t *)mtd_lns_luaStream_close,
   (lns_method_t *)mtd_lns_luaStream_flush,
   (lns_method_t *)mtd_lns_luaStream_read,
   (lns_method_t *)mtd_lns_luaStream_seek,
   (lns_method_t *)mtd_lns_luaStream_write,
};
lns_type_meta_t lns_type_meta_lns_Mapping = { "lns_Mapping" };
lns_type_meta_t lns_type_meta_lns_io = { "lns_io" };
lns_mtd_lns_io_t lns_mtd_lns_io = {
   mtd_lns_io__del,
   NULL,
};
lns_any_t * l_var_lns_io_stdout;
lns_any_t * l_var_lns_io_stdin;
lns_any_t * l_var_lns_io_stderr;
lns_type_meta_t lns_type_meta_lns_package = { "lns_package" };
lns_mtd_lns_package_t lns_mtd_lns_package = {
   mtd_lns_package__del,
   NULL,
};
lns_any_t * l_var_lns_package_path;
lns_type_meta_t lns_type_meta_lns_os = { "lns_os" };
lns_mtd_lns_os_t lns_mtd_lns_os = {
   mtd_lns_os__del,
   NULL,
};
lns_type_meta_t lns_type_meta_lns_string = { "lns_string" };
lns_mtd_lns_string_t lns_mtd_lns_string = {
   mtd_lns_string__del,
   NULL,
};
lns_type_meta_t lns_type_meta_lns_math = { "lns_math" };
lns_mtd_lns_math_t lns_mtd_lns_math = {
   mtd_lns_math__del,
   NULL,
};
lns_type_meta_t lns_type_meta_lns_debug = { "lns_debug" };
lns_mtd_lns_debug_t lns_mtd_lns_debug = {
   mtd_lns_debug__del,
   NULL,
};
static void mtd_lns_luaStream__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_luaStream_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_luaStream, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   lns_init_if( &pObj->imp.lns_iStream, _pEnv, pAny, &lns_if_lns_luaStream_imp_lns_iStream, lns_iStream );
   lns_init_if( &pObj->imp.lns_oStream, _pEnv, pAny, &lns_if_lns_luaStream_imp_lns_oStream, lns_oStream );
   return pAny;
}
static void mtd_lns_io__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_io_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_io, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
static void mtd_lns_package__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_package_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_package, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
static void mtd_lns_os__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_os_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_os, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
static void mtd_lns_string__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_string_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_string, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
static void mtd_lns_math__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_math_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_math, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
static void mtd_lns_debug__del( lns_env_t * _pEnv, lns_any_t * pObj ) {
}
lns_any_t * lns_class_lns_debug_new( lns_env_t * _pEnv){
   lns_class_new_( _pEnv, lns_debug, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lns_value_type_none;
   return pAny;
}
void l_call_mtd_lns_iStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_iStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
lns_stem_t l_call_mtd_lns_iStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1){
return lns_mtd_lns_iStream( pObj )->read( _pEnv, lns_getImpObj( pObj ) , arg1);
}
void l_call_mtd_lns_oStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_oStream( pObj )->close( _pEnv, lns_getImpObj( pObj ) );
}
void l_call_mtd_lns_oStream_flush( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_oStream( pObj )->flush( _pEnv, lns_getImpObj( pObj ) );
}
lns_stem_t l_call_mtd_lns_oStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1){
return lns_mtd_lns_oStream( pObj )->write( _pEnv, lns_getImpObj( pObj ) , arg1);
}
void l_call_mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_luaStream( pObj )->close( _pEnv, pObj );
}
void l_call_mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_luaStream( pObj )->flush( _pEnv, pObj );
}
lns_stem_t l_call_mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1){
return lns_mtd_lns_luaStream( pObj )->read( _pEnv, pObj , arg1);
}
lns_stem_t l_call_mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2){
return lns_mtd_lns_luaStream( pObj )->seek( _pEnv, pObj , arg1, arg2);
}
lns_stem_t l_call_mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1){
return lns_mtd_lns_luaStream( pObj )->write( _pEnv, pObj , arg1);
}
void l_call_mtd_lns_Mapping__toMap( lns_env_t * _pEnv, lns_any_t * pObj){
lns_mtd_lns_Mapping( pObj )->_toMap( _pEnv, lns_getImpObj( pObj ) );
}
static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv );
void lns_init_lns_builtin( lns_env_t * _pEnv ){
   if ( s_module.readyFlag ) {
      return;
   }
   s_module.readyFlag = true;
   lns_add2list( &_pEnv->loadModuleTop, &s_module);
   
   lns_block_t * pBlock_62 = lns_enter_module( _pEnv, 2, 0, 0 );
   s_module.pBlock = pBlock_62;
   lns_set_block_any( pBlock_62, 0, lns_module_globalStemList);
   lns_setQ_any( lns_module_globalStemList, lns_class_List_new( _pEnv ));
   lns_set_block_any( pBlock_62, 1, lns_module_path);
   lns_setQ_any( lns_module_path, lns_litStr2any( _pEnv, "lns_builtin"));
   lns_enter_block( _pEnv, 0, 0, 0 );
   
   lns_init_lns_builtin_Sub( _pEnv );
   
   lns_leave_block( _pEnv );
}
#include "lns_builtinInc.c"
