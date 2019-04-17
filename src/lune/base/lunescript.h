#ifndef __LUNESCRIPT__
#define __LUNESCRIPT__

#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>

#ifdef __cplusplus
extern "C" {
#endif


    typedef int __lune_int_t;
    typedef double __lune_real_t;
    typedef void * __lune_class_t;


    typedef struct __lune_stem_t __lune_stem_t;
    typedef struct __lune_block_t __lune_block_t;
    typedef struct __lune_env_t __lune_env_t;
    typedef struct __lune_form_t __lune_form_t;
    typedef struct __lune_List_t __lune_List_t;

    /**
     * ブロックの最大深度。
     * 深い関数の再帰呼び出しをすると消費するが、 1000 もあれば十分。
     */
#define __LUNE_BLOCK_MAX_DEPTH 1000

    /**
     * 全ブロックで保持する stem 型の値の最大数。
     *
     * 要は、main から関数コールしていった時に、ローカル変数を何個使用できるか。
     */
#define __LUNE_STEM_POOL_MAX_NUM 100000

#define __lune_set_block_stem( BLOCK, INDEX, STEM )     \
    (BLOCK)->pStemBuf[ INDEX ] = STEM



#define __lune_initVal( SYMBOL, BLOCK, INDEX, VAL )     \
    __lune_setQ( SYMBOL, VAL );                         \
    __lune_set_block_stem( BLOCK, INDEX, SYMBOL );
    
    
    /**
       STEM 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に VAL をセットする前に、
       変数 SYM が保持する値の参照カウントをデクリメント。
    */
#define __lune_setq( ENV, SYM, VAL )            \
    if ( SYM != NULL ) {                        \
        __lune_decre_ref( ENV, SYM );           \
    }                                           \
    __lune_setQ( SYM, VAL );

    /**
       STEM 型の値 VAL を、 変数 SYM に代入する。

       - VAL の参照カウントインクリメント
       - VAL を unassignStem から除外
    */
#define __lune_setQ( SYM, VAL )                 \
    SYM = VAL;                                  \
    __lune_setQ_( SYM );
    
#if 0
    SYM->refCount++;                            \
    if ( SYM->retValFlag ) {                    \
        SYM->retValFlag = FALSE;                \
    }                                           \
    __lune_rmFromList( SYM );
#endif
    

    /**
       クラスのインスタンスを生成する。

       @param CLASS 生成するクラスシンボル
       @param STEMVAL 生成した stem 値を格納する変数シンボル
       @param CLASSVAL 生成したクラスインスタンスを格納する変数シンボル
    */
#define __lune_class_new_( ENV, CLASS, STEMVAL, CLASSVAL )              \
    __lune_class_new2_( ENV, CLASS, CLASS, STEMVAL, CLASSVAL )
    // __lune_stem_t * STEMVAL = __lune_class_new( ENV, sizeof( CLASS ) );
    // CLASS * CLASSVAL = __lune_obj_##CLASS( pStem );                   
    // CLASSVAL->pMtd = &__mtd_##CLASS;


#define __lune_class_new2_( ENV, CLASS, SCLASS, STEMVAL, CLASSVAL )     \
    __lune_stem_t * STEMVAL = __lune_class_new( ENV, sizeof( CLASS ) ); \
    CLASS * CLASSVAL = __lune_obj_##SCLASS( pStem );                    \
    CLASSVAL->pMtd = &__mtd_##SCLASS
    



    typedef enum {
        __lune_value_type_none,
        __lune_value_type_int,
        __lune_value_type_real,
        __lune_value_type_str,
        __lune_value_type_class,
        __lune_value_type_ddd,
        __lune_value_type_form,
    } __lune_value_type_t;



    /**
     * 関数の型
     *
     * @param pInfo フォーム情報
     * @param ... 関数の引数
     * @return 関数の戻り値
     */
    typedef __lune_stem_t * __lune_func_t( __lune_env_t * _pEnv, __lune_stem_t * pInfo, ... );

    /**
     * メソッドの型
     *
     * @param pObj クラスのインスタンスを保持する stem
     * @param ... メソッドの引数
     * @return メソッドの戻り値
     */
    typedef __lune_stem_t * __lune_method_t( __lune_env_t * _pEnv, __lune_stem_t * pObj, ... );



    /**
     * クラスのインスタンスの gc の型
     *
     * @param pObj クラスのインスタンスを保持する stem
     * @param freeFlag インスタンスを開放する場合 true
     */
    typedef void __lune_gc_t( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag );

    /**
     * クラスのメソッドの最小構造。
     *
     * 各クラスのインスタンスは、メソッドのポインタを保持する構造体を持つ。
     * メソッドのポインタを保持する構造体の先頭は、必ず gc のメソッドでなければならない。
     */
    typedef struct __mtd__Class_t {
        __lune_gc_t * _gc;
    } __mtd__Class_t;

    /**
     * クラスのインスタンスの最小構造。
     *
     * 各クラスのインスタンスは、
     * メソッドのポインタを保持する構造体のポインタ pMtd を先頭に持つ。
     * クラスのメンバは、 pMtd の次に宣言する。
     */
    typedef struct __lune_Class_t {
        __mtd__Class_t * pMtd;
    } __lune_Class_t;


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
    } __lune_str_t;

    /**
     * ... を保持する型
     */
    typedef struct {
        /** ... に含む stem を管理するバッファ */
        __lune_stem_t ** pStemList;
        /** pStemList で管理している stem の数 */
        int len;
    } __lune_ddd_t;

    /**
     * form の情報。
     */
    struct __lune_form_t {
        /** 関数 */
        __lune_func_t * pFunc;
        /** form 内でアクセスする外部変数を管理するバッファ */
        __lune_stem_t ** pStemList;
        /** pStemList で管理している stem の数 */
        int len;
    };

    typedef void __lune_listObj_t;

    typedef struct __mtd_List_t __mtd_List_t;
    struct __lune_List_t {
        __mtd_List_t * pMtd;
        __lune_listObj_t * pObj;
    };

    /** stem 型データ */
    struct __lune_stem_t {
        /** 保持しているデータの型 */
        __lune_value_type_t type;
        /** このデータを参照している数 */
        int refCount;
        /** 戻り値として処理中か */
        bool retValFlag;
        /** 実データ */
        union {
            __lune_int_t intVal;
            __lune_real_t realVal;
            __lune_str_t str;
            __lune_ddd_t ddd;
            __lune_form_t form;
            __lune_class_t classVal;
            __lune_List_t list;
        } val;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。アサイン済みの場合 NULL。 */
        struct __lune_stem_t * pNext;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。  */
        struct __lune_stem_t * pPrev;
    };

    struct __lune_block_t {
        /** このブロックの深度を判定するための目安 */
        void * pStackAddr;
        /** ブロック深度 */
        int blockDepth;
        /** このブロックで管理する stem 型を保持するバッファ */
        __lune_stem_t ** pStemBuf;
        /** pStemBuf で管理する値の数 */
        int len;
        /**
         * このブロック内で生成され、
         * 変数にアサインされていない stem 型データの双方向リスト。
         * 実際の先頭要素は unassignStemTop.pNext。
         */
        __lune_stem_t unassignStemTop;
        /**
         * このブロックでコールしている関数の戻り値で、
         * 変数にアサインされていない stem 型データの双方向リスト。
         * 実際の先頭要素は retValUnassignStemTop.pNext。
         */
        __lune_stem_t retValUnassignStemTop;
    };

    struct __lune_env_t {
        /** 値を返さない関数の戻り値 */
        __lune_stem_t * pNoneStem;
        /**
         * ブロック情報で利用する pStemBuf のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        __lune_stem_t * stemPPool[ __LUNE_STEM_POOL_MAX_NUM ];
        /** stemPPool をどこまで使用しているか個数を示す */
        int useStemPoolNum;
        /** ブロック情報の Queue。*/
        __lune_block_t blockQueue[ __LUNE_BLOCK_MAX_DEPTH ];
        /** 現在のブロックの深度 */
        int blockDepth;
        /** 現在確保している stem の数 */
        int allocNum;
    };

    /**
     * リストの末尾に STEM を追加。
     */
#define __lune_add2list( TOP, STEM )            \
    STEM->pNext = TOP;                          \
    STEM->pPrev = (TOP)->pPrev;                 \
    (TOP)->pPrev->pNext = STEM;                 \
    (TOP)->pPrev = STEM;

    /**
     * リストから STEM を除外。
     */
#define __lune_rmFromList( STEM )               \
    if ( (STEM)->pNext != NULL ) {              \
        (STEM)->pPrev->pNext = (STEM)->pNext;   \
        (STEM)->pNext->pPrev = (STEM)->pPrev;   \
        (STEM)->pNext = NULL;                   \
    }
  
    extern __lune_stem_t * __lune_int2stem( __lune_env_t * _pEnv, __lune_int_t val );
    extern __lune_str_t __lune_createLiteralStr( const char * pStr );
    extern __lune_stem_t * __lune_str2stem( __lune_env_t * _pEnv, __lune_str_t val );
    extern __lune_stem_t * __lune_func2stem( __lune_env_t * _pEnv, __lune_func_t * pFunc, int num, ... );
    extern __lune_stem_t * __lune_createDDD( __lune_env_t * _pEnv, int num, ... );


    extern __lune_stem_t * __lune_setRet( __lune_env_t * __pEnv, __lune_stem_t * pStem );
    extern void __lune_setQ_( __lune_stem_t * pStem );
    extern __lune_block_t * __lune_enter_func( __lune_env_t * _pEnv, int num, int argNum, ... );
    extern void __lune_leave_block( __lune_env_t * _pEnv );
    extern __lune_block_t * __lune_enter_block( __lune_env_t * _pEnv, int stemVerNum );
    extern void __lune_decre_ref( __lune_env_t * _pEnv, __lune_stem_t * pStem );
    extern __lune_stem_t * __lune_class_new( __lune_env_t * _pEnv, int size );
    extern void __lune_class_del( __lune_env_t * _pEnv, void * pObj );



    extern void __lune_print( __lune_env_t * _pEnv, __lune_stem_t * pArg );


#ifdef __cplusplus
}
#endif

#include <lns_collection.h>

#endif
