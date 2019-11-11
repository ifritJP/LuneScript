// lune/base/lns_builtin.c
#include <lunescript.h>
#include <lns_builtin.h>
static lune_any_t ** lune_module_globalStemList;
static void u_mtd_luaStream__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_io__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_package__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_os__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_string__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_math__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_debug__del( lune_env_t * _pEnv, lune_any_t * pObj );
static void u_mtd_luaStream_close( lune_env_t * _pEnv, lune_any_t * pObj);
static void u_mtd_luaStream_flush( lune_env_t * _pEnv, lune_any_t * pObj);
static lune_stem_t u_mtd_luaStream_read( lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1);
static lune_stem_t u_mtd_luaStream_seek( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1, lune_int_t arg2);
static lune_stem_t u_mtd_luaStream_write( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1);
lune_type_meta_t lune_type_meta_iStream = { "iStream" };
lune_type_meta_t lune_type_meta_oStream = { "oStream" };
static lune_mtd_iStream_t lune_if_luaStream_imp_iStream = {
   (lune_method_t *)u_mtd_luaStream_close,
   (lune_method_t *)u_mtd_luaStream_read,
};
static lune_mtd_oStream_t lune_if_luaStream_imp_oStream = {
   (lune_method_t *)u_mtd_luaStream_close,
   (lune_method_t *)u_mtd_luaStream_flush,
   (lune_method_t *)u_mtd_luaStream_write,
};
lune_type_meta_t lune_type_meta_luaStream = { "luaStream" };
lune_mtd_luaStream_t lune_mtd_luaStream = {
   u_mtd_luaStream__del,
   NULL,
   (lune_method_t *)u_mtd_luaStream_close,
   (lune_method_t *)u_mtd_luaStream_flush,
   (lune_method_t *)u_mtd_luaStream_read,
   (lune_method_t *)u_mtd_luaStream_seek,
   (lune_method_t *)u_mtd_luaStream_write,
};
lune_type_meta_t lune_type_meta_Mapping = { "Mapping" };
lune_type_meta_t lune_type_meta_io = { "io" };
lune_mtd_io_t lune_mtd_io = {
   u_mtd_io__del,
   NULL,
};
lune_any_t * lune_var_io_stdin;
lune_any_t * lune_var_io_stderr;
lune_any_t * lune_var_io_stdout;
lune_type_meta_t lune_type_meta_package = { "package" };
lune_mtd_package_t lune_mtd_package = {
   u_mtd_package__del,
   NULL,
};
lune_any_t * lune_var_package_path;
lune_type_meta_t lune_type_meta_os = { "os" };
lune_mtd_os_t lune_mtd_os = {
   u_mtd_os__del,
   NULL,
};
lune_type_meta_t lune_type_meta_string = { "string" };
lune_mtd_string_t lune_mtd_string = {
   u_mtd_string__del,
   NULL,
};
lune_type_meta_t lune_type_meta_math = { "math" };
lune_mtd_math_t lune_mtd_math = {
   u_mtd_math__del,
   NULL,
};
lune_type_meta_t lune_type_meta_debug = { "debug" };
lune_mtd_debug_t lune_mtd_debug = {
   u_mtd_debug__del,
   NULL,
};
static void u_mtd_luaStream__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_luaStream_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, luaStream, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   lune_init_if( &pObj->imp.iStream, _pEnv, pAny, &lune_if_luaStream_imp_iStream, iStream );
   lune_init_if( &pObj->imp.oStream, _pEnv, pAny, &lune_if_luaStream_imp_oStream, oStream );
   return pAny;
}
static void u_mtd_io__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_io_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, io, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
static void u_mtd_package__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_package_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, package, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
static void u_mtd_os__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_os_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, os, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
static void u_mtd_string__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_string_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, string, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
static void u_mtd_math__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_math_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, math, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
static void u_mtd_debug__del( lune_env_t * _pEnv, lune_any_t * pObj ) {
}
lune_any_t * lune_class_debug_new( lune_env_t * _pEnv){
   lune_class_new_( _pEnv, debug, pAny, pObj );
   pObj->pImp = &pObj->imp;
   pObj->imp.sentinel.type = lune_value_type_none;
   return pAny;
}
void u_call_mtd_iStream_close( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_iStream( pObj )->close( _pEnv, lune_getImpObj( pObj ) );
}
lune_stem_t u_call_mtd_iStream_read( lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1){
return lune_mtd_iStream( pObj )->read( _pEnv, lune_getImpObj( pObj ) , arg1);
}
void u_call_mtd_oStream_close( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_oStream( pObj )->close( _pEnv, lune_getImpObj( pObj ) );
}
void u_call_mtd_oStream_flush( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_oStream( pObj )->flush( _pEnv, lune_getImpObj( pObj ) );
}
lune_stem_t u_call_mtd_oStream_write( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1){
return lune_mtd_oStream( pObj )->write( _pEnv, lune_getImpObj( pObj ) , arg1);
}
void u_call_mtd_luaStream_close( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_luaStream( pObj )->close( _pEnv, pObj );
}
void u_call_mtd_luaStream_flush( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_luaStream( pObj )->flush( _pEnv, pObj );
}
lune_stem_t u_call_mtd_luaStream_read( lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1){
return lune_mtd_luaStream( pObj )->read( _pEnv, pObj , arg1);
}
lune_stem_t u_call_mtd_luaStream_seek( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1, lune_int_t arg2){
return lune_mtd_luaStream( pObj )->seek( _pEnv, pObj , arg1, arg2);
}
lune_stem_t u_call_mtd_luaStream_write( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1){
return lune_mtd_luaStream( pObj )->write( _pEnv, pObj , arg1);
}
void u_call_mtd_Mapping__toMap( lune_env_t * _pEnv, lune_any_t * pObj){
lune_mtd_Mapping( pObj )->_toMap( _pEnv, lune_getImpObj( pObj ) );
}
static void lune_init_builtinSub( lune_env_t * _pEnv );
void lune_init_lns_builtin( lune_env_t * _pEnv ){
   lune_block_t * pBlock_62 = lune_enter_module( 1, 0, 0 );
   lune_set_block_any( pBlock_62, 0, lune_module_globalStemList);
   lune_setQ_any( lune_module_globalStemList, lune_class_List_new( _pEnv ));
   lune_enter_block( _pEnv, 0, 0, 0 );
   lune_init_builtinSub( _pEnv );
   lune_leave_block( _pEnv );
}
#include "lns_builtinInc.c"
