#ifndef __lns_builtin__
#define __lns_builtin__
       
typedef struct lune_mtd_iStream_t {
   lune_method_t * close;
   lune_method_t * read;
} lune_mtd_iStream_t;
typedef struct iStream {
   lune_type_meta_t * pMeta;
   lune_any_t * pObj;
   lune_mtd_iStream_t * pMtd;
} iStream;
#define lune_mtd_iStream( OBJ )                     \
          ((iStream*)&OBJ->val.ifVal)->pMtd
typedef struct lune_mtd_oStream_t {
   lune_method_t * close;
   lune_method_t * flush;
   lune_method_t * write;
} lune_mtd_oStream_t;
typedef struct oStream {
   lune_type_meta_t * pMeta;
   lune_any_t * pObj;
   lune_mtd_oStream_t * pMtd;
} oStream;
#define lune_mtd_oStream( OBJ )                     \
          ((oStream*)&OBJ->val.ifVal)->pMtd
typedef struct lune_mtd_luaStream_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
   lune_method_t * close;
   lune_method_t * flush;
   lune_method_t * read;
   lune_method_t * seek;
   lune_method_t * write;
} lune_mtd_luaStream_t;
typedef struct u_if_imp_luaStream_t {
   lune_any_t iStream;
   lune_any_t oStream;
   lune_any_t sentinel;
} u_if_imp_luaStream_t;
typedef struct luaStream {
   lune_type_meta_t * pMeta;
   u_if_imp_luaStream_t * pImp;
   lune_mtd_luaStream_t * pMtd;
   // member
   // interface implements
   u_if_imp_luaStream_t imp;
} luaStream;
#define lune_mtd_luaStream( OBJ )                     \
                (((luaStream*)OBJ->val.classVal)->pMtd )
#define lune_obj_luaStream( OBJ ) ((luaStream*)OBJ->val.classVal)
#define lune_if_luaStream( OBJ ) ((luaStream*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_luaStream_new( lune_env_t * _pEnv);
typedef struct lune_mtd_Mapping_t {
   lune_method_t * _toMap;
} lune_mtd_Mapping_t;
typedef struct Mapping {
   lune_type_meta_t * pMeta;
   lune_any_t * pObj;
   lune_mtd_Mapping_t * pMtd;
} Mapping;
#define lune_mtd_Mapping( OBJ )                     \
          ((Mapping*)&OBJ->val.ifVal)->pMtd
typedef struct lune_mtd_io_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_io_t;
typedef struct u_if_imp_io_t {
   lune_any_t sentinel;
} u_if_imp_io_t;
typedef struct io {
   lune_type_meta_t * pMeta;
   u_if_imp_io_t * pImp;
   lune_mtd_io_t * pMtd;
   // member
   // interface implements
   u_if_imp_io_t imp;
} io;
#define lune_mtd_io( OBJ )                     \
                (((io*)OBJ->val.classVal)->pMtd )
#define lune_obj_io( OBJ ) ((io*)OBJ->val.classVal)
#define lune_if_io( OBJ ) ((io*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_io_new( lune_env_t * _pEnv);
typedef struct lune_mtd_package_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_package_t;
typedef struct u_if_imp_package_t {
   lune_any_t sentinel;
} u_if_imp_package_t;
typedef struct package {
   lune_type_meta_t * pMeta;
   u_if_imp_package_t * pImp;
   lune_mtd_package_t * pMtd;
   // member
   // interface implements
   u_if_imp_package_t imp;
} package;
#define lune_mtd_package( OBJ )                     \
                (((package*)OBJ->val.classVal)->pMtd )
#define lune_obj_package( OBJ ) ((package*)OBJ->val.classVal)
#define lune_if_package( OBJ ) ((package*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_package_new( lune_env_t * _pEnv);
typedef struct lune_mtd_os_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_os_t;
typedef struct u_if_imp_os_t {
   lune_any_t sentinel;
} u_if_imp_os_t;
typedef struct os {
   lune_type_meta_t * pMeta;
   u_if_imp_os_t * pImp;
   lune_mtd_os_t * pMtd;
   // member
   // interface implements
   u_if_imp_os_t imp;
} os;
#define lune_mtd_os( OBJ )                     \
                (((os*)OBJ->val.classVal)->pMtd )
#define lune_obj_os( OBJ ) ((os*)OBJ->val.classVal)
#define lune_if_os( OBJ ) ((os*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_os_new( lune_env_t * _pEnv);
typedef struct lune_mtd_string_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_string_t;
typedef struct u_if_imp_string_t {
   lune_any_t sentinel;
} u_if_imp_string_t;
typedef struct string {
   lune_type_meta_t * pMeta;
   u_if_imp_string_t * pImp;
   lune_mtd_string_t * pMtd;
   // member
   // interface implements
   u_if_imp_string_t imp;
} string;
#define lune_mtd_string( OBJ )                     \
                (((string*)OBJ->val.classVal)->pMtd )
#define lune_obj_string( OBJ ) ((string*)OBJ->val.classVal)
#define lune_if_string( OBJ ) ((string*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_string_new( lune_env_t * _pEnv);
typedef struct lune_mtd_math_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_math_t;
typedef struct u_if_imp_math_t {
   lune_any_t sentinel;
} u_if_imp_math_t;
typedef struct math {
   lune_type_meta_t * pMeta;
   u_if_imp_math_t * pImp;
   lune_mtd_math_t * pMtd;
   // member
   // interface implements
   u_if_imp_math_t imp;
} math;
#define lune_mtd_math( OBJ )                     \
                (((math*)OBJ->val.classVal)->pMtd )
#define lune_obj_math( OBJ ) ((math*)OBJ->val.classVal)
#define lune_if_math( OBJ ) ((math*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_math_new( lune_env_t * _pEnv);
typedef struct lune_mtd_debug_t {
   lune_del_t * _del;
   lune_gc_t * _gc;
} lune_mtd_debug_t;
typedef struct u_if_imp_debug_t {
   lune_any_t sentinel;
} u_if_imp_debug_t;
typedef struct debug {
   lune_type_meta_t * pMeta;
   u_if_imp_debug_t * pImp;
   lune_mtd_debug_t * pMtd;
   // member
   // interface implements
   u_if_imp_debug_t imp;
} debug;
#define lune_mtd_debug( OBJ )                     \
                (((debug*)OBJ->val.classVal)->pMtd )
#define lune_obj_debug( OBJ ) ((debug*)OBJ->val.classVal)
#define lune_if_debug( OBJ ) ((debug*)OBJ->val.classVal)->pImp
extern lune_any_t * lune_class_debug_new( lune_env_t * _pEnv);
extern void u_call_mtd_iStream_close( lune_env_t * _pEnv, lune_any_t * pObj);
extern lune_stem_t u_call_mtd_iStream_read( lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1);
extern void u_call_mtd_oStream_close( lune_env_t * _pEnv, lune_any_t * pObj);
extern void u_call_mtd_oStream_flush( lune_env_t * _pEnv, lune_any_t * pObj);
extern lune_stem_t u_call_mtd_oStream_write( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1);
extern void u_call_mtd_luaStream_close( lune_env_t * _pEnv, lune_any_t * pObj);
extern void u_call_mtd_luaStream_flush( lune_env_t * _pEnv, lune_any_t * pObj);
extern lune_stem_t u_call_mtd_luaStream_read( lune_env_t * _pEnv, lune_any_t * pObj, lune_stem_t arg1);
extern lune_stem_t u_call_mtd_luaStream_seek( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1, lune_int_t arg2);
extern lune_stem_t u_call_mtd_luaStream_write( lune_env_t * _pEnv, lune_any_t * pObj, lune_any_t * arg1);
extern void u_call_mtd_Mapping__toMap( lune_env_t * _pEnv, lune_any_t * pObj);
extern lune_stem_t u_mtd_io_open( lune_env_t * _pEnv, lune_any_t * arg1, lune_stem_t arg2);
extern lune_stem_t u_mtd_io_popen( lune_env_t * _pEnv, lune_any_t * arg1);
extern lune_stem_t u_mtd_package_searchpath( lune_env_t * _pEnv, lune_any_t * arg1, lune_any_t * arg2);
extern lune_real_t u_mtd_os_clock( lune_env_t * _pEnv);
extern lune_stem_t u_mtd_os_date( lune_env_t * _pEnv, lune_stem_t arg1, lune_stem_t arg2);
extern lune_int_t u_mtd_os_difftime( lune_env_t * _pEnv, lune_stem_t arg1, lune_stem_t arg2);
extern lune_any_t * u_mtd_os_exit( lune_env_t * _pEnv, lune_stem_t arg1);
extern lune_stem_t u_mtd_os_remove( lune_env_t * _pEnv, lune_any_t * arg1);
extern lune_stem_t u_mtd_os_rename( lune_env_t * _pEnv, lune_any_t * arg1, lune_any_t * arg2);
extern lune_stem_t u_mtd_os_time( lune_env_t * _pEnv, lune_stem_t arg1);
extern lune_stem_t u_mtd_string_byte( lune_env_t * _pEnv, lune_any_t * arg1, lune_stem_t arg2, lune_stem_t arg3);
extern lune_any_t * u_mtd_string_dump( lune_env_t * _pEnv, lune_any_t * arg1, lune_stem_t arg2);
extern lune_stem_t u_mtd_string_find( lune_env_t * _pEnv, lune_any_t * arg1, lune_any_t * arg2, lune_stem_t arg3, lune_stem_t arg4);
extern lune_any_t * u_mtd_string_format( lune_env_t * _pEnv, lune_any_t * arg1, lune_stem_t arg2);
extern lune_stem_t u_mtd_string_gmatch( lune_env_t * _pEnv, lune_any_t * arg1, lune_any_t * arg2);
extern lune_stem_t u_mtd_string_gsub( lune_env_t * _pEnv, lune_any_t * arg1, lune_any_t * arg2, lune_any_t * arg3);
extern lune_any_t * u_mtd_string_lower( lune_env_t * _pEnv, lune_any_t * arg1);
extern lune_any_t * u_mtd_string_rep( lune_env_t * _pEnv, lune_any_t * arg1, lune_int_t arg2);
extern lune_any_t * u_mtd_string_reverse( lune_env_t * _pEnv, lune_any_t * arg1);
extern lune_any_t * u_mtd_string_sub( lune_env_t * _pEnv, lune_any_t * arg1, lune_int_t arg2, lune_stem_t arg3);
extern lune_any_t * u_mtd_string_upper( lune_env_t * _pEnv, lune_any_t * arg1);
extern lune_real_t u_mtd_math_random( lune_env_t * _pEnv, lune_stem_t arg1, lune_stem_t arg2);
extern void u_mtd_math_randomseed( lune_env_t * _pEnv, lune_stem_t arg1);
extern lune_stem_t u_mtd_debug_getinfo( lune_env_t * _pEnv, lune_int_t arg1);
extern lune_stem_t u_mtd_debug_getlocal( lune_env_t * _pEnv, lune_int_t arg1, lune_int_t arg2);
extern lune_any_t * lune_var_io_stdin;
extern lune_any_t * lune_var_io_stderr;
extern lune_any_t * lune_var_io_stdout;
extern lune_any_t * lune_var_package_path;
extern void lune_init_lns_builtin( lune_env_t * _pEnv );
#endif
