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

#ifndef __LUNESCRIPT__
#define __LUNESCRIPT__

#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>
#include <stddef.h>
#include <lns_alloc.h>

#ifdef __cplusplus
extern "C" {
#endif

#define LUNE_DEBUG_POS __FILE__, __LINE__
#define LUNE_DEBUG_DECL const char * pFile, int lineNo

#ifdef LUNE_DEBUG
#define LUNE_DEBUG_CALL_LOG printf( "%s -- %p\n", __func__, pObj )
#else
#define LUNE_DEBUG_CALL_LOG
#endif
    
    typedef int lune_bool_t;
    typedef int lune_int_t;
    typedef double lune_real_t;


    typedef struct lune_any_t lune_any_t;
    typedef struct lune_block_t lune_block_t;
    typedef struct lune_env_t lune_env_t;
    typedef struct lune_form_t lune_form_t;


    typedef enum {
        lune_imdType_sentinel,
        lune_imdType_int,
        lune_imdType_real,
        lune_imdType_str,
        lune_imdType_list,
        lune_imdType_map,
        lune_imdType_set,
        lune_imdType_any,
    } lune_imdType_t;
    typedef struct lune_imdEntry_t lune_imdEntry_t;

#define lune_imdSet_t lune_imdList_t

#define lune_imdSentinel { .type = lune_imdType_sentinel }
#define lune_imdInt( VAL ) { .type = lune_imdType_int, .val.valInt = (VAL) }
#define lune_imdReal( VAL ) { .type = lune_imdType_real, .val.valReal = (VAL) }
#define lune_imdStr( VAL ) { .type = lune_imdType_str, .val.str = (VAL) }
#define lune_imdAny( VAL ) { .type = lune_imdType_any, .val.any = (VAL) }
#define lune_imdList( LIST, ... )                          \
    lune_imdVal_t LIST[] = { __VA_ARGS__, lune_imdSentinel };
#define lune_imdSet( SET, ... )                          \
    lune_imdVal_t SET[] = { __VA_ARGS__, lune_imdSentinel };
#define lune_imdMap( MAP, ... )                          \
    lune_imdEntry_t MAP[] = { __VA_ARGS__, { lune_imdSentinel, lune_imdSentinel } };


    /**
       将来のマルチスレッド対応時の排他制御範囲。
     */
#define lune_lock( ... )    __VA_ARGS__


    typedef struct lune_imdVal_t {
        lune_imdType_t type;
        union {
            lune_int_t valInt;
            lune_real_t valReal;
            const char * str;
            struct lune_imdVal_t * list;
            struct lune_imdEntry_t * map;
            struct lune_imdVal_t * set;
            lune_any_t * any;
        } val;
    } lune_imdVal_t;

    struct lune_imdEntry_t {
        lune_imdVal_t key;
        lune_imdVal_t val;
    };


    typedef enum {
        lune_stem_type_none,
        lune_stem_type_nil,
        lune_stem_type_int,
        lune_stem_type_real,
        lune_stem_type_bool,
        lune_stem_type_any,
    } lune_stem_type_t;

    typedef struct {
        lune_stem_type_t type;
        union {
            lune_int_t intVal;
            lune_real_t realVal;
            lune_bool_t boolVal;
            lune_any_t * pAny;
        } val;
    } lune_stem_t;

#define LUNE_STEM_BASE(TYPE) (lune_stem_t){ .type = TYPE }
#define LUNE_STEM_INT(VAL) \
    (lune_stem_t){ .type = lune_stem_type_int, .val = { .intVal = VAL } }
#define LUNE_STEM_REAL(VAL) \
    (lune_stem_t){ .type = lune_stem_type_real, .val = { .realVal = VAL } }
#define LUNE_STEM_BOOL(VAL) \
    (lune_stem_t){ .type = lune_stem_type_bool, .val = { .boolVal = VAL } }
#define LUNE_STEM_ANY(VAL) \
    (lune_stem_t){ .type = lune_stem_type_any, .val = { .pAny = VAL } }
    

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
     * and or 演算子の評価中に利用するスタック数
     */
#define LUNE_VAL_STACK_MAX 10000
    
    

    /**
     * 全ブロックで保持する any 型の値の最大数。
     *
     * 要は、main から関数コールしていった時に、ローカル変数を何個使用できるか。
     */
#define LUNE_ANY_POOL_MAX_NUM 100000

#define lune_set_block_var( BLOCK, INDEX, TYPE, VAR )    \
    VAR = lune_var_alloc( _pEnv, BLOCK, INDEX, TYPE );   \
    VAR->stem.type = TYPE
    
#define lune_set_block_stem( BLOCK, INDEX, VAR )     \
    (BLOCK)->pStemList[ INDEX ] = &(VAR);           \
    VAR.type = lune_stem_type_nil;


#define lune_initVal_stem( SYMBOL, BLOCK, INDEX, VAL )                \
    SYMBOL = (BLOCK)->pVarList[ INDEX ];                              \
    SYMBOL->stem = VAL
    
#define lune_initVal_var( SYMBOL, BLOCK, INDEX, VAL )                   \
    lune_set_block_var( BLOCK, INDEX, lune_stem_type_any, SYMBOL );     \
    lune_setQ( SYMBOL->stem, (VAL) );

#define lune_decre_ref_alter( ENV, STEM )               \
    {                                                   \
        lune_stem_t _work = STEM;                       \
        if ( _work.type == lune_stem_type_any ) {       \
            lune_decre_ref( ENV, _work.val.pAny );      \
        }                                               \
    }
    
    

#define lune_setq( ENV, STEM, VAL )                                     \
    {                                                                   \
        if ( STEM.type == lune_stem_type_any ) {                        \
            if ( STEM.val.pAny != NULL ) {                              \
                lune_decre_ref( ENV, STEM.val.pAny );                   \
            }                                                           \
        }                                                               \
        STEM = VAL;                                                     \
        lune_setQ_( STEM.val.pAny );                                    \
    }
    
    /**
       ANY 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に VAL をセットする前に、
       変数 SYM が保持する値の参照カウントをデクリメント。
    */
#define lune_setq_any( ENV, SYM, VAL )          \
    {                                           \
        lune_any_t * __WORK = VAL;             \
        if ( (*SYM) != __WORK ) {               \
            if ( (*SYM) != NULL ) {             \
                lune_decre_ref( ENV, (*SYM) );  \
            }                                   \
            (*SYM) = __WORK;                    \
            lune_setQ_( (*SYM) );               \
        }                                       \
    }



#define lune_setQ( STEM, VAL )                                  \
    {                                                           \
        lune_stem_t ___work = VAL;                              \
        if ( ___work.type == lune_stem_type_any ) {             \
            lune_setQ_any( &STEM.val.pAny, ___work.val.pAny );  \
            STEM.type = lune_stem_type_any;                     \
        }                                                       \
        else {                                                  \
            STEM = ___work;                                     \
        }                                                       \
    }
    
    /**
       ANY 型の値 VAL を、 変数 SYM に代入する。

       - VAL の参照カウントインクリメント
       - VAL を managedAnyTop から除外
    */
#define lune_setQ_any( SYM, VAL )               \
    {                                           \
        lune_any_t * __WORK = VAL;              \
        if ( (*(SYM)) != __WORK ) {             \
            (*(SYM)) = __WORK;                  \
            lune_setQ_( (*(SYM)) );             \
        }                                       \
    }

#define lune_getMRet( ENV, INDEX ) (lune_fromDDD( ENV->pMRet, INDEX ))
    
    

    /**
       クラスのインスタンスを生成する。

       @param CLASS 生成するクラスシンボル
       @param ANYVAL 生成した any 値を格納する変数シンボル
       @param CLASSVAL 生成したクラスインスタンスを格納する変数シンボル
    */
#define lune_class_new_( ENV, CLASS, ANYVAL, CLASSVAL )              \
    lune_class_new2_( ENV, CLASS, CLASS, ANYVAL, CLASSVAL )
    // lune_any_t * ANYVAL = lune_class_new( ENV, sizeof( CLASS ) );
    // CLASS * CLASSVAL = lune_obj_##CLASS( pAny );                   
    // CLASSVAL->pMtd = &lune_mtd_##CLASS;


#define lune_class_new2_( ENV, CLASS, SCLASS, ANYVAL, CLASSVAL )      \
    lune_stem_t ANYVAL = lune_class_new( ENV, sizeof( CLASS ) );      \
    CLASS * CLASSVAL = lune_obj_##SCLASS( ANYVAL.val.pAny );          \
    CLASSVAL->pMtd = &lune_mtd_##SCLASS;                              \
    CLASSVAL->pMeta = &lune_type_meta_##SCLASS;
    

#define lune_decl_enum_get__text( ENUM, VALTYPE )                       \
    lune_stem_t ENUM##_get__txt( lune_env_t * _pEnv, lune_any_t * _pForm, \
                                   lune_##VALTYPE##_t val )             \
    {                                                                   \
        return lune_mtd_Map_get(                                        \
            _pEnv, ENUM##_val2NameMap, lune_##VALTYPE##2stem( _pEnv, val ) ); \
    }
#define lune_decl_enum_get__text_any( ENUM )                           \
    lune_stem_t ENUM##_get__txt( lune_env_t * _pEnv, lune_any_t * _pForm, \
                                   lune_stem_t val )             \
    {                                                                   \
        return lune_mtd_Map_get( _pEnv, ENUM##_val2NameMap, val );   \
    }

#define lune_init_if( ANY, ENV, OBJ, MTD, IF )                    \
    {                                                             \
        lune_any_t * __pAny = ANY;                                \
        lune_init_any( __pAny, lune_value_type_if, ENV );         \
        __pAny->refCount = 0;                                     \
        __pAny->val.ifVal.pMeta = &lune_type_meta_##IF;           \
        __pAny->val.ifVal.pObj = OBJ.val.pAny;                    \
        __pAny->val.ifVal.pMtd = MTD;                             \
    }

#define lune_init_any( ANY, TYPE, ENV )            \
    {                                                \
        lune_any_t * _pAny = ANY;                 \
        _pAny->type = TYPE;                         \
    }

#define lune_getImpObj( ANY ) (ANY)->val.ifVal.pObj


    typedef enum {
        lune_value_type_none,
        lune_value_type_str,
        lune_value_type_class,
        lune_value_type_if,
        lune_value_type_ddd,
        lune_value_type_mRet,
        lune_value_type_form,
        lune_value_type_List,
        lune_value_type_Array,
        lune_value_type_Set,
        lune_value_type_Map,
        lune_value_type_itList,
        lune_value_type_itSet,
        lune_value_type_itMap,
        lune_value_type_alge,
        lune_value_type_luaVal,
    } lune_value_type_t;
    typedef struct lune_itList_t lune_itList_t;
    typedef struct lune_itSet_t lune_itSet_t;
    typedef struct lune_itMap_t lune_itMap_t;
    

    typedef lune_any_t * lune_newfunc_t( lune_env_t * _pEnv, ... );


    /**
     * 関数の型
     *
     * @param pInfo フォーム情報
     * @param ... 関数の引数
     * @return 関数の戻り値
     */
    typedef lune_stem_t lune_closure_t( lune_env_t * _pEnv, lune_any_t * pInfo, ... );
    typedef lune_int_t lune_closureInt_t( lune_env_t * _pEnv, lune_any_t * pInfo, ... );
    typedef lune_real_t lune_closureReal_t( lune_env_t * _pEnv, lune_any_t * pInfo, ... );
    
    typedef lune_stem_t lune_func_t( lune_env_t * _pEnv, ... );
    typedef lune_int_t lune_funcInt_t( lune_env_t * _pEnv, ... );
    typedef lune_real_t lune_funcReal_t( lune_env_t * _pEnv, ... );

    
    /**
     * メソッドの型
     *
     * @param pObj クラスのインスタンスを保持する any
     * @param ... メソッドの引数
     * @return メソッドの戻り値
     */
    typedef lune_stem_t lune_method_t( lune_env_t * _pEnv, lune_any_t * pObj, ... );
    typedef lune_int_t lune_method_int_t( lune_env_t * _pEnv, lune_any_t * pObj, ... );
    typedef lune_real_t lune_method_real_t( lune_env_t * _pEnv, lune_any_t * pObj, ... );



    /**
     * クラスのインスタンスの gc の型
     *
     * @param pObj クラスのインスタンスを保持する any
     */
    typedef void lune_gc_t( lune_env_t * _pEnv, lune_any_t * pObj );

    /**
     * クラスのインスタンスの開放処理
     *
     * @param pObj クラスのインスタンスを保持する any
     */
    typedef void lune_del_t( lune_env_t * _pEnv, lune_any_t * pObj );

    /**
     * 型のメタ情報
     */
    typedef struct lune_type_meta_t {
        /** 型名 */
        const char * pName;        
    } lune_type_meta_t;
    
    /**
     * クラスのメソッドの最小構造。
     *
     * 各クラスのインスタンスは、メソッドのポインタを保持する構造体を持つ。
     * メソッドのポインタを保持する構造体の先頭は、
     * 必ず del, gc のメソッドでなければならない。
     */
    typedef struct lune_mtd__Class_t {
        lune_del_t * _del;
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
        lune_type_meta_t * pMeta;
        void * pIFdummy;
        lune_mtd__Class_t * pMtd;
    } lune_Class_t;


    /**
     * alge 型を開放する際のコールバック
     */
    typedef void (lune_algeVal_gc_t)( lune_env_t * _pEnv, void * pVal );
    /**
     * alge 型のインスタンス。
     */
    typedef struct lune_Alge_t {
        int type;
        void * pVal;
        lune_algeVal_gc_t * gc;
    } lune_Alge_t;

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
        lune_stem_t * stemList;
        /** pAnyList で管理している any の数 */
        int len;
    } lune_ddd_t;

    typedef struct {
        lune_value_type_t type;
        lune_any_t * pAny;
    } lune_luaVal_t;


#define lune_isClosure( FORM ) \
    ((FORM)->val.form.needFormParam)
#define lune_setClosure( FORM ) \
    (FORM)->val.form.needFormParam = true;
    
#define lune_closure( FORM )                  \
    (FORM)->val.form.pFunc
#define lune_closure_int( FORM )                  \
    ((lune_closureInt_t *)(FORM)->val.form.pFunc)
#define lune_closure_real( FORM )                  \
    ((lune_closureReal_t *)(FORM)->val.form.pFunc)

#define lune_func( FORM )                  \
    ((lune_func_t *)(FORM)->val.form.pFunc)
#define lune_func_int( FORM )                  \
    ((lune_funcInt_t *)(FORM)->val.form.pFunc)
#define lune_func_real( FORM )                  \
    ((lune_funcReal_t *)(FORM)->val.form.pFunc)
    
#define lune_form_closure_var( FORM, INDEX )        \
    (FORM)->val.form.pClosureValList[ INDEX ].pVar
#define lune_form_closure( FORM, INDEX )                        \
    &((FORM)->val.form.pClosureValList[ INDEX ])

    typedef struct lune_var_t {
        //lune_any_t * pAny;
        lune_stem_t stem;
        int refCount;
    } lune_var_t;

    typedef struct lune_closureVal_t {
        lune_var_t * pVar;
    } lune_closureVal_t;
    
    /**
     * form の情報。
     */
    struct lune_form_t {
        /** 関数 */
        lune_closure_t * pFunc;
        /** form 内でアクセスする外部変数を管理するバッファ */
        lune_closureVal_t * pClosureValList;
        /** pClosureValList で管理している any の数 */
        int len;
        /** 関数コール時に、この構造体を必要とする場合 true*/
        bool needFormParam;
        /** 引数の数*/
        //int argNum;
        /** 引数に ... を持つかどうか */
        bool hasDDD;
    };

    typedef void lune_listObj_t;
    typedef struct lune_mtd_List_t lune_mtd_List_t;
    typedef struct lune_List_t {
        lune_type_meta_t * pMeta;
        void * pIFdummy;
        lune_mtd_List_t * pMtd;
        lune_listObj_t * pObj;
    } lune_List_t;


    typedef void lune_setObj_t;
    typedef struct lune_mtd_Set_t lune_mtd_Set_t;
    typedef struct lune_Set_t {
        lune_type_meta_t * pMeta;
        void * pIFdummy;
        lune_mtd_Set_t * pMtd;
        lune_setObj_t * pObj;
    } lune_Set_t;

    typedef void lune_mapObj_t;
    typedef struct lune_mtd_Map_t lune_mtd_Map_t;
    typedef struct lune_Map_t {
        lune_type_meta_t * pMeta;
        void * pIFdummy;
        lune_mtd_Map_t * pMtd;
        lune_mapObj_t * pObj;
    } lune_Map_t;

    typedef struct lune_if_t {
        lune_type_meta_t * pMeta;
        lune_any_t * pObj;
        void * pMtd;
    } lune_if_t;
    
    /** any 型データ */
    struct lune_any_t {
        /** 保持しているデータの型 */
        lune_value_type_t type;
        /** このデータを参照している数 */
        int refCount;
        /** 実データ */
        union {
            lune_str_t str;
            lune_ddd_t ddd;
            lune_form_t form;
            lune_Class_t * classVal;
            lune_itList_t * itList;
            lune_itSet_t * itSet;
            lune_itMap_t * itMap;
            lune_if_t ifVal;
            lune_Alge_t alge;
            lune_luaVal_t luaVal;
        } val;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。アサイン済みの場合 NULL。 */
        struct lune_any_t * pNext;
        /** 変数にアサインされる前の値を管理する双方向リスト構造。  */
        struct lune_any_t * pPrev;
    };

    struct lune_block_t {
        /** このブロックの深度を判定するための目安 */
        void * pStackAddr;
        /** ブロック深度 */
        int blockDepth;
        /** このブロックで管理する var 型を保持するバッファ */
        lune_var_t ** pVarList;
        /** このブロックで管理する stem 型を保持するバッファ */
        lune_stem_t ** pStemList;
        /** pStemList で管理する値の数 */
        int stemLen;
        /** pVarBuf で管理する値の数 */
        int varLen;
        /**
         * このブロック内で生成された any 型データの双方向リスト。
         * 実際の先頭要素は managedAnyTop.pNext。
         */
        lune_any_t managedAnyTop;
   };

    typedef struct {
        /** 値を返さない関数の戻り値 */
        lune_stem_t noneStem;
        /** nil */
        lune_stem_t nilStem;
        /** true */
        lune_stem_t trueStem;
        /** false */
        lune_stem_t falseStem;

        /** 要素 0 個の DDD*/
        lune_stem_t ddd0;
        
    } lune_global_t;

    extern lune_global_t lune_global;


    struct lua_State;
    
    struct lune_env_t {
        struct lua_State * pLua;
        lune_allocator_t allocateor;
        /**
         * ブロック情報で利用する pVarList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lune_var_t * varPPool[ LUNE_ANY_POOL_MAX_NUM ];
        /** varPPool をどこまで使用しているか個数を示す */
        int useVarPoolNum;
        /**
         * ブロック情報で利用する pAnyList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lune_stem_t * stemPPool[ LUNE_ANY_POOL_MAX_NUM ];
        /** anyPPool をどこまで使用しているか個数を示す */
        int useStemPoolNum;
        /** ブロック情報の Queue。*/
        lune_block_t blockQueue[ LUNE_BLOCK_MAX_DEPTH ];
        /** 現在のブロックの深度 */
        int blockDepth;

        /** sort callback */
        lune_stem_t pSortCallback;

        /** 値のスタック */
        lune_stem_t pValStack[ LUNE_VAL_STACK_MAX ];
        /** pValStack の現在位置 */
        int stackPos;

        /** 処理中の MRet */
        lune_any_t * pMRet;
    };


#define lune_set2DDDArg( ANY, INDEX, VAL )  \
    ANY->val.ddd.stemList[ INDEX ] = VAL;

#if 0
#define lune_fromDDD( ANY, INDEX )  \
//     (ANY->val.ddd.stemList[ INDEX ])
#else
#define lune_fromDDD( ANY, INDEX ) \
    lune_getValFromDDD( ANY, INDEX )
#endif
    

#define lune_lenDDD( ANY ) ANY->val.ddd.len

#define lune_str2stem( ENV, VAL )               \
    _lune_str2stem( LUNE_DEBUG_POS, ENV, VAL )
#define lune_litStr2stem( ENV, STR )            \
    _lune_litStr2stem( LUNE_DEBUG_POS, ENV, STR )
#define lune_cloneBin2stem( ENV, STR, LEN )               \
    _lune_cloneBin2stem( LUNE_DEBUG_POS, ENV, STR, LEN )
#define lune_func2stem( ENV, FUNC, ARGNUM, HASDDD, NUM, ... )            \
    _lune_func2stem( LUNE_DEBUG_POS, ENV, FUNC, ARGNUM, HASDDD, NUM, ##__VA_ARGS__ )
#define lune_createDDD( ENV, HASDDD, NUM, ... )         \
    _lune_createDDD( LUNE_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__)
#define lune_createDDDOnly( ENV, NUM )          \
    _lune_createDDDOnly( LUNE_DEBUG_POS, ENV, NUM )
#define lune_createSubDDD( ENV, OFFSET, DDD ) \
    _lune_createSubDDD( LUNE_DEBUG_POS, ENV, OFFSET, DDD )
#define lune_createMRet( ENV, HASDDD, NUM, ... )        \
    _lune_createMRet( LUNE_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__ )
#define lune_class_new( ENV, SIZE )             \
    _lune_class_new( LUNE_DEBUG_POS, ENV, SIZE )
#define lune_alge_new( ENV, VALTYPE, SIZE, GC )                    \
    _lune_alge_new( LUNE_DEBUG_POS, ENV, VALTYPE, SIZE, GC )
#define lune_it_new( ENV, TYPE, VAL )                  \
    _lune_it_new( LUNE_DEBUG_POS, ENV, TYPE, VAL )
#define lune_createList( ENV, LIST )                  \
    _lune_createList( LUNE_DEBUG_POS, ENV, LIST )
#define lune_createSet( ENV, SET )                  \
    _lune_createSet( LUNE_DEBUG_POS, ENV, SET )
#define lune_createMap( ENV, ENTRY )                  \
    _lune_createMap( LUNE_DEBUG_POS, ENV, ENTRY )
#define lune_createImmediateVal( ENV, VAL )                     \
    _lune_createImmediateVal( LUNE_DEBUG_POS, ENV, VAL )
#define lune_luaVal_new( ENV, TYPE )            \
    _lune_luaVal_new( LUNE_DEBUG_POS, ENV, TYPE )
    
    

#define lune_abort( MESS ) _lune_abort( MESS, LUNE_DEBUG_POS )
    

    extern void _lune_abort( const char * pMessage, const char * pFile, int lineNo );
    
    extern lune_stem_t _lune_str2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_str_t val );
    extern lune_stem_t _lune_litStr2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, const char * pStr );
    extern lune_stem_t _lune_func2stem( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_closure_t * pFunc, int argNum, bool hasDDD, int num, ... );
    extern lune_stem_t _lune_createDDD( LUNE_DEBUG_DECL, lune_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lune_stem_t _lune_createDDDOnly( LUNE_DEBUG_DECL, lune_env_t * _pEnv, int num );
    extern lune_stem_t _lune_createSubDDD(
        const char * pFile, int lineNo, lune_env_t * _pEnv, int offset, lune_any_t * pDDD );
    extern lune_stem_t _lune_createMRet( LUNE_DEBUG_DECL, lune_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lune_stem_t _lune_class_new( LUNE_DEBUG_DECL, lune_env_t * _pEnv, int size );
    extern lune_stem_t _lune_alge_new( LUNE_DEBUG_DECL, lune_env_t * _pEnv, int valType, int size, lune_algeVal_gc_t * gc );
    extern lune_any_t * _lune_it_new(
        LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_value_type_t type, void * pVal );
    extern lune_stem_t _lune_cloneBin2stem(
        const char * pFile, int lineNo, lune_env_t * _pEnv, const void * pBuf, int len );

    extern lune_stem_t _lune_luaVal_new(
        const char * pFile, int lineNo, lune_env_t * _pEnv, lune_value_type_t type );
    

    extern lune_int_t lune_stem2int( lune_stem_t stem );
    extern lune_real_t lune_stem2real( lune_stem_t stem );

    extern void lune_init_alge( lune_stem_t * pStem, lune_any_t * pAny, int valType );

    
    extern lune_stem_t _lune_createList( LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pList );
    extern lune_stem_t _lune_createSet(
        LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pSet );
    extern lune_stem_t _lune_createMap(
        LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdEntry_t * pEntry );

    extern lune_stem_t _lune_createImmediateVal(
        LUNE_DEBUG_DECL, lune_env_t * _pEnv, lune_imdVal_t * pVal );
    
    

    extern lune_str_t lune_createLiteralStr( const char * pStr );

    extern lune_var_t * lune_var_alloc(
        lune_env_t * _pEnv, lune_block_t * pBlock, int index, lune_stem_type_t type );

    extern lune_block_t * lune_enter_module( int anyNum, int varNum );
    extern void lune_reset_block( lune_env_t * _pEnv );

    extern void lune_setRet( lune_env_t * _pEnv, lune_stem_t pAny );
    extern void lune_setQ_( lune_any_t * pAny );
    extern void lune_setOverwrite( lune_any_t * pAny );
    extern lune_block_t * lune_enter_func( lune_env_t * _pEnv, int anyNum, int varNum, int argNum, ... );
    extern void lune_leave_block( lune_env_t * _pEnv );
    extern void lune_leave_blockMulti( lune_env_t * _pEnv, int num );
    extern lune_block_t * lune_enter_block( lune_env_t * _pEnv, int anyNum, int varNum );
    extern bool lune_decre_ref( lune_env_t * _pEnv, lune_any_t * pAny );
    extern void lune_it_delete( lune_env_t * _pEnv, lune_any_t * pAny );
    extern lune_stem_t lune_call_form( lune_env_t * _pEnv, lune_any_t * _pForm, lune_stem_t _pDDD );
    extern lune_stem_t lune_getIF( lune_env_t * _pEnv, lune_any_t * pIFAny );
    extern lune_stem_t lune_toIF( lune_env_t * _pEnv, lune_any_t * pAny, const lune_type_meta_t * pMeta );
    extern lune_stem_t lune_setMRet( lune_env_t * _pEnv, lune_any_t * pAny );
    extern lune_stem_t lune_getValFromDDD( lune_any_t * pAny, int index );



    extern lune_stem_t lune_unwrap_stem( lune_stem_t pAny, lune_stem_t pDefVal );
    extern lune_int_t lune_unwrap_int( lune_stem_t pAny );
    extern lune_int_t lune_unwrap_intDefault( lune_stem_t pAny, lune_int_t val );
    extern lune_real_t lune_unwrap_real( lune_stem_t pAny );
    extern lune_real_t lune_unwrap_realDefault( lune_stem_t pAny, lune_real_t val );

    

    extern bool lune_isCondTrue( const lune_stem_t pAny );
    extern bool lune_incStack( lune_env_t * _pEnv );
    extern bool lune_setStackVal( lune_env_t * _pEnv, lune_stem_t pVal );
    extern lune_stem_t lune_popVal( lune_env_t * _pEnv, bool dummy );
    extern lune_stem_t lune_op_not( lune_env_t * _pEnv, lune_stem_t pAny );
    extern bool lune_equals( lune_stem_t stem1, lune_stem_t stem2 );

    extern void lune_run_module( lune_env_t * _pEnv );



    extern void lune_print( lune_env_t * _pEnv, lune_stem_t pArg );

#ifdef __cplusplus
}
#endif

#include <lns_collection.h>
#include <lns_luaWrapper.h>

#endif


