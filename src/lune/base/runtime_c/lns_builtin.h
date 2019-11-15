#ifndef __lns_builtin__
#define __lns_builtin__
       
extern lns_any_t * lns_f__fcall( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2 );
extern lns_int_t lns_f__kind( lns_env_t * _pEnv, lns_stem_t arg1 );
extern lns_stem_t lns_f__load( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2 );
extern void lns_f_collectgarbage( lns_env_t * _pEnv );
extern lns_any_t * lns_f_error( lns_env_t * _pEnv, lns_any_t * arg1 );
extern lns_stem_t lns_f_load( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2, lns_stem_t arg3, lns_stem_t arg4 );
extern lns_stem_t lns_f_loadfile( lns_env_t * _pEnv, lns_any_t * arg1 );
extern void lns_f_print( lns_env_t * _pEnv, lns_stem_t arg1 );
extern lns_stem_t lns_f_require( lns_env_t * _pEnv, lns_any_t * arg1 );
extern lns_stem_t lns_f_tonumber( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2 );
extern lns_any_t * lns_f_tostring( lns_env_t * _pEnv, lns_stem_t arg1 );
extern lns_any_t * lns_f_type( lns_env_t * _pEnv, lns_stem_t arg1 );
typedef struct lns_mtd_lns_iStream_t {
   lns_method_t * close;
   lns_method_t * read;
} lns_mtd_lns_iStream_t;
typedef struct lns_iStream {
   lns_type_meta_t * pMeta;
   lns_any_t * pObj;
   lns_mtd_lns_iStream_t * pMtd;
} lns_iStream;
#define lns_mtd_lns_iStream( OBJ )                     \
             ((lns_iStream*)&OBJ->val.ifVal)->pMtd
extern lns_type_meta_t lns_type_meta_lns_iStream;
typedef struct lns_mtd_lns_oStream_t {
   lns_method_t * close;
   lns_method_t * flush;
   lns_method_t * write;
} lns_mtd_lns_oStream_t;
typedef struct lns_oStream {
   lns_type_meta_t * pMeta;
   lns_any_t * pObj;
   lns_mtd_lns_oStream_t * pMtd;
} lns_oStream;
#define lns_mtd_lns_oStream( OBJ )                     \
             ((lns_oStream*)&OBJ->val.ifVal)->pMtd
extern lns_type_meta_t lns_type_meta_lns_oStream;
typedef struct lns_mtd_lns_luaStream_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
   lns_method_t * close;
   lns_method_t * flush;
   lns_method_t * read;
   lns_method_t * seek;
   lns_method_t * write;
} lns_mtd_lns_luaStream_t;
typedef struct u_if_imp_lns_luaStream_t {
   lns_any_t lns_iStream;
   lns_any_t lns_oStream;
   lns_any_t sentinel;
} u_if_imp_lns_luaStream_t;
typedef struct lns_luaStream {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_luaStream_t * pImp;
   lns_mtd_lns_luaStream_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_luaStream_t imp;
} lns_luaStream;
#define lns_mtd_lns_luaStream( OBJ )                     \
                (((lns_luaStream*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_luaStream( OBJ ) ((lns_luaStream*)OBJ->val.classVal)
#define lns_if_lns_luaStream( OBJ ) ((lns_luaStream*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_luaStream_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_Mapping_t {
   lns_method_t * _toMap;
} lns_mtd_lns_Mapping_t;
typedef struct lns_Mapping {
   lns_type_meta_t * pMeta;
   lns_any_t * pObj;
   lns_mtd_lns_Mapping_t * pMtd;
} lns_Mapping;
#define lns_mtd_lns_Mapping( OBJ )                     \
             ((lns_Mapping*)&OBJ->val.ifVal)->pMtd
extern lns_type_meta_t lns_type_meta_lns_Mapping;
typedef struct lns_mtd_lns_io_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_io_t;
typedef struct u_if_imp_lns_io_t {
   lns_any_t sentinel;
} u_if_imp_lns_io_t;
typedef struct lns_io {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_io_t * pImp;
   lns_mtd_lns_io_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_io_t imp;
} lns_io;
#define lns_mtd_lns_io( OBJ )                     \
                (((lns_io*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_io( OBJ ) ((lns_io*)OBJ->val.classVal)
#define lns_if_lns_io( OBJ ) ((lns_io*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_io_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_package_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_package_t;
typedef struct u_if_imp_lns_package_t {
   lns_any_t sentinel;
} u_if_imp_lns_package_t;
typedef struct lns_package {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_package_t * pImp;
   lns_mtd_lns_package_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_package_t imp;
} lns_package;
#define lns_mtd_lns_package( OBJ )                     \
                (((lns_package*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_package( OBJ ) ((lns_package*)OBJ->val.classVal)
#define lns_if_lns_package( OBJ ) ((lns_package*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_package_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_os_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_os_t;
typedef struct u_if_imp_lns_os_t {
   lns_any_t sentinel;
} u_if_imp_lns_os_t;
typedef struct lns_os {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_os_t * pImp;
   lns_mtd_lns_os_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_os_t imp;
} lns_os;
#define lns_mtd_lns_os( OBJ )                     \
                (((lns_os*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_os( OBJ ) ((lns_os*)OBJ->val.classVal)
#define lns_if_lns_os( OBJ ) ((lns_os*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_os_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_string_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_string_t;
typedef struct u_if_imp_lns_string_t {
   lns_any_t sentinel;
} u_if_imp_lns_string_t;
typedef struct lns_string {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_string_t * pImp;
   lns_mtd_lns_string_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_string_t imp;
} lns_string;
#define lns_mtd_lns_string( OBJ )                     \
                (((lns_string*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_string( OBJ ) ((lns_string*)OBJ->val.classVal)
#define lns_if_lns_string( OBJ ) ((lns_string*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_string_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_math_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_math_t;
typedef struct u_if_imp_lns_math_t {
   lns_any_t sentinel;
} u_if_imp_lns_math_t;
typedef struct lns_math {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_math_t * pImp;
   lns_mtd_lns_math_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_math_t imp;
} lns_math;
#define lns_mtd_lns_math( OBJ )                     \
                (((lns_math*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_math( OBJ ) ((lns_math*)OBJ->val.classVal)
#define lns_if_lns_math( OBJ ) ((lns_math*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_math_new( lns_env_t * _pEnv);
typedef struct lns_mtd_lns_debug_t {
   lns_del_t * _del;
   lns_gc_t * _gc;
} lns_mtd_lns_debug_t;
typedef struct u_if_imp_lns_debug_t {
   lns_any_t sentinel;
} u_if_imp_lns_debug_t;
typedef struct lns_debug {
   lns_type_meta_t * pMeta;
   u_if_imp_lns_debug_t * pImp;
   lns_mtd_lns_debug_t * pMtd;
   // member
   // interface implements
   u_if_imp_lns_debug_t imp;
} lns_debug;
#define lns_mtd_lns_debug( OBJ )                     \
                (((lns_debug*)OBJ->val.classVal)->pMtd )
#define lns_obj_lns_debug( OBJ ) ((lns_debug*)OBJ->val.classVal)
#define lns_if_lns_debug( OBJ ) ((lns_debug*)OBJ->val.classVal)->pImp
extern lns_any_t * lns_class_lns_debug_new( lns_env_t * _pEnv);
extern void l_call_mtd_lns_iStream_close( lns_env_t * _pEnv, lns_any_t * pObj);
extern lns_stem_t l_call_mtd_lns_iStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1);
extern void l_call_mtd_lns_oStream_close( lns_env_t * _pEnv, lns_any_t * pObj);
extern void l_call_mtd_lns_oStream_flush( lns_env_t * _pEnv, lns_any_t * pObj);
extern lns_stem_t l_call_mtd_lns_oStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1);
extern void l_call_mtd_lns_luaStream_close( lns_env_t * _pEnv, lns_any_t * pObj);
extern void l_call_mtd_lns_luaStream_flush( lns_env_t * _pEnv, lns_any_t * pObj);
extern lns_stem_t l_call_mtd_lns_luaStream_read( lns_env_t * _pEnv, lns_any_t * pObj, lns_stem_t arg1);
extern lns_stem_t l_call_mtd_lns_luaStream_seek( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1, lns_int_t arg2);
extern lns_stem_t l_call_mtd_lns_luaStream_write( lns_env_t * _pEnv, lns_any_t * pObj, lns_any_t * arg1);
extern void l_call_mtd_lns_Mapping__toMap( lns_env_t * _pEnv, lns_any_t * pObj);
extern lns_stem_t mtd_lns_io_open( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2);
extern lns_stem_t mtd_lns_io_popen( lns_env_t * _pEnv, lns_any_t * arg1);
extern lns_stem_t mtd_lns_package_searchpath( lns_env_t * _pEnv, lns_any_t * arg1, lns_any_t * arg2);
extern lns_real_t mtd_lns_os_clock( lns_env_t * _pEnv);
extern lns_stem_t mtd_lns_os_date( lns_env_t * _pEnv, lns_stem_t arg1, lns_stem_t arg2);
extern lns_int_t mtd_lns_os_difftime( lns_env_t * _pEnv, lns_stem_t arg1, lns_stem_t arg2);
extern lns_any_t * mtd_lns_os_exit( lns_env_t * _pEnv, lns_stem_t arg1);
extern lns_stem_t mtd_lns_os_remove( lns_env_t * _pEnv, lns_any_t * arg1);
extern lns_stem_t mtd_lns_os_rename( lns_env_t * _pEnv, lns_any_t * arg1, lns_any_t * arg2);
extern lns_stem_t mtd_lns_os_time( lns_env_t * _pEnv, lns_stem_t arg1);
extern lns_stem_t mtd_lns_string_byte( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2, lns_stem_t arg3);
extern lns_any_t * mtd_lns_string_dump( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2);
extern lns_stem_t mtd_lns_string_find( lns_env_t * _pEnv, lns_any_t * arg1, lns_any_t * arg2, lns_stem_t arg3, lns_stem_t arg4);
extern lns_any_t * mtd_lns_string_format( lns_env_t * _pEnv, lns_any_t * arg1, lns_stem_t arg2);
extern lns_stem_t mtd_lns_string_gmatch( lns_env_t * _pEnv, lns_any_t * arg1, lns_any_t * arg2);
extern lns_stem_t mtd_lns_string_gsub( lns_env_t * _pEnv, lns_any_t * arg1, lns_any_t * arg2, lns_any_t * arg3);
extern lns_any_t * mtd_lns_string_lower( lns_env_t * _pEnv, lns_any_t * arg1);
extern lns_any_t * mtd_lns_string_rep( lns_env_t * _pEnv, lns_any_t * arg1, lns_int_t arg2);
extern lns_any_t * mtd_lns_string_reverse( lns_env_t * _pEnv, lns_any_t * arg1);
extern lns_any_t * mtd_lns_string_sub( lns_env_t * _pEnv, lns_any_t * arg1, lns_int_t arg2, lns_stem_t arg3);
extern lns_any_t * mtd_lns_string_upper( lns_env_t * _pEnv, lns_any_t * arg1);
extern lns_real_t mtd_lns_math_random( lns_env_t * _pEnv, lns_stem_t arg1, lns_stem_t arg2);
extern void mtd_lns_math_randomseed( lns_env_t * _pEnv, lns_stem_t arg1);
extern lns_stem_t mtd_lns_debug_getinfo( lns_env_t * _pEnv, lns_int_t arg1);
extern lns_stem_t mtd_lns_debug_getlocal( lns_env_t * _pEnv, lns_int_t arg1, lns_int_t arg2);
extern lns_any_t * l_var_lns_io_stderr;
extern lns_any_t * l_var_lns_io_stdin;
extern lns_any_t * l_var_lns_io_stdout;
extern lns_any_t * l_var_lns_package_path;
extern void lns_init_lns_builtin( lns_env_t * _pEnv );
#endif
