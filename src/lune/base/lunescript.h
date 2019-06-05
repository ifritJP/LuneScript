#ifndef __LUNESCRIPT__
#define __LUNESCRIPT__

#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>
#include <lns_alloc.h>

#ifdef __cplusplus
extern "C" {
#endif


#define LUNE_DEBUG_POS __FILE__, __LINE__
#define LUNE_DEBUG_DECL const char * pFile, int lineNo
#define LUNE_DEBUG_CALL_LOG printf( "%s -- %p\n", __func__, pObj )
    
    typedef int lune_bool_t;
    typedef int lune_int_t;
    typedef double lune_real_t;
    typedef void * lune_class_t;


    typedef struct lune_stem_t lune_stem_t;
    typedef struct lune_block_t lune_block_t;
    typedef struct lune_env_t lune_env_t;
    typedef struct lune_form_t lune_form_t;

    /**
     * ブロックの最大深度。
     * 深い関数の再帰呼び出しをすると消費するが、 1000 もあれば十分。
     */
#define LUNE_BLOCK_MAX_DEPTH 1000

    /**
     * モジュールの最大個数
     */
#define LUNE_MODULE_MAX_NUM 10000

    
    /**
     * 全ブロックで保持する stem 型の値の最大数。
     *
     * 要は、main から関数コールしていった時に、ローカル変数を何個使用できるか。
     */
#define LUNE_STEM_POOL_MAX_NUM 100000

#define lune_set_block_stem( BLOCK, INDEX, VAR )     \
    VAR = (BLOCK)->pVarList[ INDEX ]



#define lune_initVal( SYMBOL, BLOCK, INDEX, VAL )     \
    lune_set_block_stem( BLOCK, INDEX, SYMBOL );   \
    lune_setQ( (&SYMBOL->pStem), VAL );                         

    /**
       STEM 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に VAL をセットする前に、
       変数 SYM が保持する値の参照カウントをデクリメント。
    */
#define lune_setq( ENV, SYM, VAL )            \
    if ( (*SYM) != NULL ) {                   \
        lune_decre_ref( ENV, (*SYM) );          \
    }                                           \
    lune_setQ( SYM, VAL );

    /**
       STEM 型の値 VAL を、 変数 SYM に代入する。

       - VAL の参照カウントインクリメント
       - VAL を managedStemTop から除外
    */
#define lune_setQ( SYM, VAL )           \
    if ( (*SYM) != VAL ) {              \
        (*SYM) = VAL;                   \
        lune_setQ_( (*SYM) );           \
    }
    
    

    /**
       クラスのインスタンスを生成する。

       @param CLASS 生成するクラスシンボル
       @param STEMVAL 生成した stem 値を格納する変数シンボル
       @param CLASSVAL 生成したクラスインスタンスを格納する変数シンボル
    */
#define lune_class_new_( ENV, CLASS, STEMVAL, CLASSVAL )              \
    lune_class_new2_( ENV, CLASS, CLASS, STEMVAL, CLASSVAL )
    // lune_stem_t * STEMVAL = lune_class_new( ENV, sizeof( CLASS ) );
    // CLASS * CLASSVAL = lune_obj_##CLASS( pStem );                   
    // CLASSVAL->pMtd = &lune_mtd_##CLASS;


#define lune_class_new2_( ENV, CLASS, SCLASS, STEMVAL, CLASSVAL )     \
    lune_stem_t * STEMVAL = lune_class_new( ENV, sizeof( CLASS ) ); \
    CLASS * CLASSVAL = lune_obj_##SCLASS( pStem );                    \
    CLASSVAL->pMtd = &lune_mtd_##SCLASS
    

    typedef enum {
        lune_value_type_none,
        lune_value_type_nil,
        lune_value_type_bool,
        lune_value_type_int,
        lune_value_type_real,
        lune_value_type_str,
        lune_value_type_class,
        lune_value_type_ddd,
        lune_value_type_mRet,
        lune_value_type_form,
        lune_value_type_List,
        lune_value_type_Array,
        lune_value_type_Set,
        lune_value_type_Map,
        lune_value_type_itSet,
        lune_value_type_itMap,
    } lune_value_type_t;

    typedef struct lune_itSet_t lune_itSet_t;
    typedef struct lune_itMap_t lune_itMap_t;
    

    typedef lune_stem_t * lune_newfunc_t( lune_env_t * _pEnv, ... );


    /**
     * 関数の型
     *
     * @param pInfo フォーム情報
     * @param ... 関数の引数
     * @return 関数の戻り値
     */
    typedef lune_stem_t * lune_func_t( lune_env_t * _pEnv, lune_stem_t * pInfo, ... );

    /**
     * メソッドの型
     *
     * @param pObj クラスのインスタンスを保持する stem
     * @param ... メソッドの引数
     * @return メソッドの戻り値
     */
    typedef lune_stem_t * lune_method_t( lune_env_t * _pEnv, lune_stem_t * pObj, ... );



    /**
     * クラスのインスタンスの gc の型
     *
     * @param pObj クラスのインスタンスを保持する stem
     * @param freeFlag インスタンスを開放する場合 true
     */
    typedef void lune_gc_t( lune_env_t * _pEnv, lune_stem_t * pObj, bool freeFlag );

    /**
     * クラスのメソッドの最小構造。
     *
     * 各クラスのインスタンスは、メソッドのポインタを保持する構造体を持つ。
     * メソッドのポインタを保持する構造体の先頭は、必ず gc のメソッドでなければならない。
     */
    typedef struct lune_mtd__Class_t {
        lune_gc_t * _gc;
    } lune_mtd__Class_t;

    /**
     * クラスのインスタンスの最小構造。
     *
     * 各クラスのインスタンスは、
     * メソッドのポインタを保持する構造体のポインタ pMtd を先頭に持つ。
     * クラスのメンバは、 pMtd の次に宣言する。
     */
    typedef struct lune_Class_t {
        lune_mtd__Class_t * pMtd;
    } lune_Class_t;


    /**
     * 文字列型データ。
     *
     * NULL 文字で終端されているとは限らない。
     */
    typedef struct {
        /** データ */
        const char * pStr;
        /** pStr のサイズ */
        int len;
        /** pStr のデータが static かどうか。 */
        bool staticFlag;
    } lune_str_t;

    /**
     * ... を保持する型。
     *
     * 関数の戻り値の多値にも使用する。
     */
    typedef struct {
        /** ... に含む stem を管理するバッファ */
        lune_stem_t ** pStemList;
        /** pStemList で管理している stem の数 */
        int len;
    } lune_ddd_t;


#define lune_form_closure( FORM, INDEX )        \
    (FORM)->val.form.ppClosureValList[ INDEX ]

    typedef struct lune_var_t {
        lune_stem_t * pStem;
        int refCount;
    } lune_var_t;
    
    
    /**
     * form の情報。
     */
    struct lune_form_t {
        /** 関数 */
        lune_func_t * pFunc;
        /** form 内でアクセスする外部変数を管理するバッファ */
        lune_var_t ** ppClosureValList;
        /** ppClosureValList で管理している stem の数 */
        int len;
        /** 引数の数*/
        int argNum;
        /** 引数に ... を持つかどうか */
        bool hasDDD;
    };

    typedef void lune_listObj_t;
    typedef struct lune_mtd_List_t lune_mtd_List_t;
    typedef struct lune_List_t {
        lune_mtd_List_t * pMtd;
        lune_listObj_t * pObj;
    } lune_List_t;


    typedef void lune_setObj_t;
    typedef struct lune_mtd_Set_t lune_mtd_Set_t;
    typedef struct lune_Set_t {
        lune_mtd_Set_t * pMtd;
        lune_setObj_t * pObj;
    } lune_Set_t;

    typedef void lune_mapObj_t;
    typedef struct lune_mtd_Map_t lune_mtd_Map_t;
    typedef struct lune_Map_t {
        lune_mtd_Map_t * pMtd;
        lune_mapObj_t * pObj;
    } lune_Map_t;
    
    /** stem 型データ */
    struct lune_stem_t {
        /** 保持しているデータの型 */
        lune_value_type_t type;
        /** このデータを参照している数 */
        int refCount;
        /** ENV */
        lune_env_t * pEnv;
        /** 実データ */
        union {
            lune_bool_t boolVal;
            lune_int_t intVal;
            lune_real_t realVal;
            lune_str_t str;
            lune_ddd_t ddd;
            lune_form_t form;
            lune_class_t classVal;
            lune_itSet_t * itSet;
            lune_itMap_t * itMap;
        } val;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。アサイン済みの場合 NULL。 */
        struct lune_stem_t * pNext;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。  */
        struct lune_stem_t * pPrev;
    };

    struct lune_block_t {
        /** このブロックの深度を判定するための目安 */
        void * pStackAddr;
        /** ブロック深度 */
        int blockDepth;
        /** このブロックで管理する stem 型を保持するバッファ */
        lune_var_t ** pVarList;
        /** pStemBuf で管理する値の数 */
        int len;
        /**
         * このブロック内で生成された stem 型データの双方向リスト。
         * 実際の先頭要素は managedStemTop.pNext。
         */
        lune_stem_t managedStemTop;
   };

    struct lune_env_t {
        lune_allocator_t allocateor;
        /** 値を返さない関数の戻り値 */
        lune_stem_t * pNoneStem;
        /** nil */
        lune_stem_t * pNilStem;
        /**
         * ブロック情報で利用する pVarList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lune_var_t * varPPool[ LUNE_STEM_POOL_MAX_NUM ];
        /** stemPPool をどこまで使用しているか個数を示す */
        int useStemPoolNum;
        /** ブロック情報の Queue。*/
        lune_block_t blockQueue[ LUNE_BLOCK_MAX_DEPTH ];
        /** 現在のブロックの深度 */
        int blockDepth;

        /** sort callback */
        lune_stem_t * pSortCallback;
    };


#define lune_set2DDDArg( STEM, INDEX, VAL )  \
    STEM->val.ddd.pStemList[ INDEX ] = VAL;

#define lune_fromDDD( STEM, INDEX )  \
    STEM->val.ddd.pStemList[ INDEX ]

#define lune_bool2stem( ENV, VAL )              \
    _lune_bool2stem( LUNE_DEBUG_POS, ENV, VAL )
#define lune_int2stem( ENV, VAL )               \
    _lune_int2stem( LUNE_DEBUG_POS, ENV, VAL )
#define lune_real2stem( ENV, VAL )              \
    _lune_real2stem( LUNE_DEBUG_POS, ENV, VAL )
#define lune_str2stem( ENV, VAL )               \
    _lune_str2stem( LUNE_DEBUG_POS, ENV, VAL )
#define lune_litStr2stem( ENV, STR )            \
    _lune_litStr2stem( LUNE_DEBUG_POS, ENV, STR )
#define lune_func2stem( ENV, FUNC, ARGNUM, HASDDD, NUM, ... )            \
    _lune_func2stem( LUNE_DEBUG_POS, ENV, FUNC, ARGNUM, HASDDD, NUM, ##__VA_ARGS__ )
#define lune_createDDD( ENV, HASDDD, NUM, ... )         \
    _lune_createDDD( LUNE_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__)
#define lune_createDDDOnly( ENV, NUM )          \
    _lune_createDDDOnly( LUNE_DEBUG_POS, ENV, NUM )
#define lune_createMRet( ENV, HASDDD, NUM, ... )        \
    _lune_createMRet( LUNE_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__ )
#define lune_class_new( ENV, SIZE )             \
    _lune_class_new( LUNE_DEBUG_POS, ENV, SIZE )
#define lune_it_new( ENV, TYPE, VAL )                  \
    _lune_it_new( LUNE_DEBUG_POS, ENV, TYPE, VAL )

    
    extern lune_stem_t * _lune_bool2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_bool_t val );
    extern lune_stem_t * _lune_int2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_int_t val );
    extern lune_stem_t * _lune_real2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_real_t val );
    extern lune_stem_t * _lune_str2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_str_t val );
    extern lune_stem_t * _lune_litStr2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, const char * pStr );
    extern lune_stem_t * _lune_func2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_func_t * pFunc, int argNum, bool hasDDD, int num, ... );
    extern lune_stem_t * _lune_createDDD( LUNE_DEBUG_DECL, lune_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lune_stem_t * _lune_createDDDOnly( LUNE_DEBUG_DECL, lune_env_t * _pEnv, int num );
    extern lune_stem_t * _lune_createMRet( LUNE_DEBUG_DECL, lune_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lune_stem_t * _lune_class_new( LUNE_DEBUG_DECL, lune_env_t * _pEnv, int size );
    extern lune_stem_t * _lune_it_new(
        LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_value_type_t type, void * pVal );

    

    extern lune_str_t lune_createLiteralStr( const char * pStr );

    extern lune_block_t * lune_enter_module( int stemVerNum );

    extern lune_stem_t * lune_setRet( lune_env_t * __pEnv, lune_stem_t * pStem );
    extern void lune_setQ_( lune_stem_t * pStem );
    extern void lune_setOverwrite( lune_stem_t * pStem );
    extern lune_block_t * lune_enter_func( lune_env_t * _pEnv, int num, int argNum, ... );
    extern void lune_leave_block( lune_env_t * _pEnv );
    extern lune_block_t * lune_enter_block( lune_env_t * _pEnv, int stemVerNum );
    extern void lune_decre_ref( lune_env_t * _pEnv, lune_stem_t * pStem );
    extern void lune_class_del( lune_env_t * _pEnv, void * pObj );
    extern void lune_it_delete( lune_env_t * _pEnv, lune_stem_t * pStem );
    extern lune_stem_t * lune_call_form( lune_env_t * _pEnv, lune_stem_t * _pForm, int num, ... );




    extern void lune_print( lune_env_t * _pEnv, lune_stem_t * _pForm, lune_stem_t * pArg );


#ifdef __cplusplus
}
#endif

#include <lns_collection.h>

#endif
