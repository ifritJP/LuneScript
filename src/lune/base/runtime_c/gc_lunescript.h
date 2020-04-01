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
#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>
#include <stdarg.h>
#include <stddef.h>
#include <lns_alloc.h>

#ifdef __cplusplus
extern "C" {
#endif

/** minor で管理するオブジェクト数 */
#define LNS_GC_MINOR_MAX 1000

#define LNS_GC_STATE_REFED 0x1
#define LNS_GC_STATE_MINOR 0x2

#define LNS_IS_GC_REFED( ANY ) ( ANY->state & LNS_GC_STATE_REFED )
#define LNS_IS_GC_MINOR( ANY ) ( ANY->state & LNS_GC_STATE_MINOR )
    

    
#define LNS_DEBUG_POS __FILE__, __LINE__
#define LNS_DEBUG_DECL const char * pFile, int lineNo

#ifdef LNS_DEBUG
#define LNS_DEBUG_CALL_LOG printf( "%s -- %p\n", __func__, pObj )
#else
#define LNS_DEBUG_CALL_LOG
#endif
    
    typedef bool lns_bool_t;
    typedef int lns_int_t;
    typedef double lns_real_t;

    typedef struct lns_envAlloc_t lns_envAlloc_t;
    
    typedef struct lns_varLink_t lns_varLink_t;
    typedef struct lns_var_t lns_var_t;
    typedef struct lns_retObj_t  lns_retObj_t;
    typedef struct lns_any_t lns_any_t;
    typedef struct lns_block_t lns_block_t;
    typedef struct lns_env_t lns_env_t;
    typedef struct lns_form_t lns_form_t;


    typedef enum {
        lns_imdType_sentinel,
        lns_imdType_stem,
        lns_imdType_int,
        lns_imdType_real,
        lns_imdType_bool,
        lns_imdType_str,
        lns_imdType_list,
        lns_imdType_map,
        lns_imdType_set,
        lns_imdType_any,
    } lns_imdType_t;
    typedef struct lns_imdEntry_t lns_imdEntry_t;

#define lns_imdSet_t lns_imdList_t

#define lns_imdSentinel { .type = lns_imdType_sentinel }
#define lns_imdStem( VAL ) { .type = lns_imdType_stem, .val.stem = (VAL) }
#define lns_imdInt( VAL ) { .type = lns_imdType_int, .val.valInt = (VAL) }
#define lns_imdReal( VAL ) { .type = lns_imdType_real, .val.valReal = (VAL) }
#define lns_imdBool( VAL ) { .type = lns_imdType_bool, .val.valBool = (VAL) }
#define lns_imdStr( VAL ) { .type = lns_imdType_str, .val.str = (VAL) }
#define lns_imdAny( VAL ) { .type = lns_imdType_any, .val.any = (VAL) }
#define lns_imdList( LIST, ... )                          \
    lns_imdVal_t LIST[] = { __VA_ARGS__, lns_imdSentinel };
#define lns_imdSet( SET, ... )                          \
    lns_imdVal_t SET[] = { __VA_ARGS__, lns_imdSentinel };
#define lns_imdMap( MAP, ... )                          \
    lns_imdEntry_t MAP[] = { __VA_ARGS__, { lns_imdSentinel, lns_imdSentinel } };


    /**
       将来のマルチスレッド対応時の排他制御範囲。
     */
#define lns_lock( ... )    __VA_ARGS__


    typedef enum {
        lns_stem_type_none,
        /** スタックの区切り。 */
        lns_stem_type_stack,
        lns_stem_type_nil,
        lns_stem_type_int,
        lns_stem_type_real,
        lns_stem_type_bool,
        lns_stem_type_any,
    } lns_stem_type_t;

#define LNS_IS_NILNONE( STEM ) \
    ( STEM.type == lns_stem_type_nil || STEM.type == lns_stem_type_none )

    typedef struct {
        lns_stem_type_t type;
        union {
            lns_int_t intVal;
            lns_real_t realVal;
            lns_bool_t boolVal;
            lns_any_t * pAny;
        } val;
    } lns_stem_t;

    typedef struct lns_imdVal_t {
        lns_imdType_t type;
        union {
            lns_int_t valInt;
            lns_real_t valReal;
            bool valBool;
            const char * str;
            struct lns_imdVal_t * list;
            struct lns_imdEntry_t * map;
            struct lns_imdVal_t * set;
            lns_any_t * any;
            lns_stem_t stem;
        } val;
    } lns_imdVal_t;

    struct lns_imdEntry_t {
        lns_imdVal_t key;
        lns_imdVal_t val;
    };


#define LNS_STEM_BASE(TYPE) (lns_stem_t){ .type = TYPE }
#define LNS_STEM_INT(VAL) \
    (lns_stem_t){ .type = lns_stem_type_int, .val = { .intVal = VAL } }
#define LNS_STEM_REAL(VAL) \
    (lns_stem_t){ .type = lns_stem_type_real, .val = { .realVal = VAL } }
#define LNS_STEM_BOOL(VAL) \
    (lns_stem_t){ .type = lns_stem_type_bool, .val = { .boolVal = VAL } }
#define LNS_STEM_ANY(VAL) \
    (lns_stem_t){ .type = lns_stem_type_any, .val = { .pAny = VAL } }
    

/**
 * リストの末尾に ANY を追加。
 */
#define lns_add2list( TOP, ANY )                \
    (ANY)->pNext = TOP;                         \
    (ANY)->pPrev = (TOP)->pPrev;                \
    (TOP)->pPrev->pNext = ANY;                  \
    (TOP)->pPrev = ANY;

    /**
     * ブロックの最大深度。
     * 深い関数の再帰呼び出しをすると消費するが、 1000 もあれば十分。
     */
#define LNS_BLOCK_MAX_DEPTH 1000

    /**
     * モジュールの最大個数
     */
#define LNS_MODULE_MAX_NUM 10000

    /**
     * and or 演算子の評価中に利用するスタック数
     */
#define LNS_VAL_STACK_MAX 1000
    
    

    /**
     * 全ブロックで保持する any 型の値の最大数。
     *
     * 要は、main から関数コールしていった時に、ローカル変数を何個使用できるか。
     */
#define LNS_ANY_POOL_MAX_NUM 100000


#define lns_check_err_from_map( ERR, ENV, MAP, NILABLE, MBR, KIND, ACC ) \
    if ( ERR == NULL ) {                                                \
        lns_stem_t _work;                                               \
        if ( lns_fromMapPrim( _pEnv, &ERR, &_work, MAP, NILABLE,        \
                              lns_litStr2any( ENV, #MBR ), KIND ) ) {   \
            MBR = _work ACC;                                            \
        }                                                               \
    }

#define lns_check_err_from_map_str( ERR, ENV, MAP, NILABLE, MBR, ACC )  \
    if ( ERR == NULL ) {                                                \
        lns_stem_t _work;                                               \
        if ( lns_fromMapStr( _pEnv, &ERR, &_work, MAP, NILABLE,        \
                             lns_litStr2any( ENV, #MBR ) ) ) {          \
            MBR = _work ACC;                                            \
        }                                                               \
    }

#define lns_check_err_from_map_stem( ERR, ENV, MAP, NILABLE, MBR )      \
    if ( ERR == NULL ) {                                                \
        lns_stem_t _name = LNS_STEM_ANY( lns_litStr2any( ENV, #MBR ) ); \
        MBR = lns_mtd_Map_get( _pEnv, MAP, _name );                     \
        if ( !NILABLE && MBR.type == lns_stem_type_nil ) {              \
            ERR = _name.val.pAny;                                       \
        }                                                               \
    }
    

#define lns_check_err_from_map_class( ERR, ENV, MAP, NILABLE, MBR, FromMap, INFO, ACC ) \
    if ( ERR == NULL ) {                                                \
        lns_stem_t _name = LNS_STEM_ANY( lns_litStr2any( ENV, #MBR ) ); \
        lns_stem_t _work = lns_mtd_Map_get( _pEnv, MAP, _name );        \
        if ( _work.type != lns_stem_type_nil ) {                        \
            lns_stem_t mret = FromMap( _pEnv, INFO, _work);             \
            _work = lns_fromDDD( mret.val.pAny, 0 );                    \
            if ( _work.type == lns_stem_type_nil ) {                    \
                lns_stem_t err = lns_fromDDD( mret.val.pAny, 1 );       \
                if ( err.type != lns_stem_type_nil ) {                  \
                    ERR = lns_strconcat(                                \
                        _pEnv, lns_litStr2any( _pEnv, #MBR "." ), err.val.pAny ); \
                }                                                       \
                else {                                                  \
                    ERR = _name.val.pAny;                               \
                }                                                       \
            }                                                           \
            else {                                                      \
                MBR = _work ACC;                                        \
            }                                                           \
        }                                                               \
        else {                                                          \
            if ( !NILABLE || _work.type != lns_stem_type_nil ) {        \
                ERR = _name.val.pAny;                                   \
            }                                                           \
        }                                                               \
    }

    

#define lns_set_block_var( BLOCK, INDEX, TYPE, VAR )    \
    VAR = lns_var_alloc( _pEnv, BLOCK, INDEX, TYPE );
    
#define lns_set_block_stem( BLOCK, INDEX, VAR )     \
    (BLOCK)->pStemList[ INDEX ] = &(VAR);           \
    VAR.type = lns_stem_type_nil;

#define lns_set_block_any( BLOCK, INDEX, SYMBOL )     \
    SYMBOL = &(BLOCK)->pAnyList[ INDEX ];

    
#define lns_initVal_any( SYMBOL, BLOCK, INDEX, ANY )       \
    lns_set_block_any( BLOCK, INDEX, SYMBOL );             \
    *SYMBOL = ANY;                                          \
    lns_setQ_( *SYMBOL );                                  \
    
    

#define lns_initVal_stem( SYMBOL, BLOCK, INDEX, VAL )                \
    SYMBOL = (BLOCK)->pVarList[ INDEX ];                              \
    SYMBOL->stem = VAL
    
#define lns_initVal_var( SYMBOL, BLOCK, INDEX, VAL )                   \
    lns_set_block_var( BLOCK, INDEX, lns_stem_type_any, SYMBOL );     \
    lns_setQ( SYMBOL->stem, (VAL) );

#define lns_decre_ref_stem( ENV, STEM )               \
    {                                                   \
        lns_stem_t _work = STEM;                       \
        if ( _work.type == lns_stem_type_any ) {       \
            lns_decre_ref( ENV, _work.val.pAny );      \
        }                                               \
    }
    
    


    /**
       STEM 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に VAL をセットする前に、
       変数 SYM が保持する値の参照カウントをデクリメント。
    */
#define lns_setq( ENV, STEM, VAL )                      \
    {                                                   \
        if ( STEM.type == lns_stem_type_any ) {         \
            if ( STEM.val.pAny != NULL ) {              \
                lns_decre_ref( ENV, STEM.val.pAny );    \
            }                                           \
        }                                               \
        STEM = VAL;                                     \
        if ( STEM.type == lns_stem_type_any ) {         \
            lns_setQ_( STEM.val.pAny );                 \
        }                                               \
    }
    
    /**
       ANY 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に VAL をセットする前に、
       変数 SYM が保持する値の参照カウントをデクリメント。
    */
#define lns_setq_any( ENV, SYM, VAL )           \
    {                                           \
        lns_any_t * __WORK = VAL;               \
        if ( (*SYM) != __WORK ) {               \
            if ( (*SYM) != NULL ) {             \
                lns_decre_ref( ENV, (*SYM) );   \
            }                                   \
            (*SYM) = __WORK;                    \
            lns_setQ_( (*SYM) );                \
        }                                       \
    }



    /**
       変数 SYM を STEM 型の値 VAL で初期化する。

       変数 SYM に値が事前にセットされている場合は、 lns_setq() を使用する。

       - VAL の参照カウントインクリメント
    */
#define lns_setQ( STEM, VAL )                           \
    {                                                   \
        STEM = VAL;                                     \
        if ( STEM.type == lns_stem_type_any ) {         \
            lns_setQ_( STEM.val.pAny );                 \
        }                                               \
    }
    
    /**
       ANY 型の値 VAL を、 変数 SYM に代入する。

       変数 SYM に値が事前にセットされている場合は、 lns_setq_any() を使用する。
       
       - VAL の参照カウントインクリメント
    */
#define lns_setQ_any( SYM, VAL )                \
    {                                           \
        lns_any_t * __WORK = VAL;               \
        if ( (*(SYM)) != __WORK ) {             \
            (*(SYM)) = __WORK;                  \
            lns_setQ_( (*(SYM)) );              \
        }                                       \
    }

#define lns_getMRet( ENV, INDEX ) (lns_fromDDD( ENV->pMRet, INDEX ))
    
    

    /**
       クラスのインスタンスを生成する。

       @param CLASS 生成するクラスシンボル
       @param ANYVAL 生成した any 値を格納する変数シンボル
       @param CLASSVAL 生成したクラスインスタンスを格納する変数シンボル
    */
#define lns_class_new_( ENV, CLASS, ANYVAL, CLASSVAL )              \
    lns_class_new2_( ENV, CLASS, CLASS, ANYVAL, CLASSVAL )
    // lns_any_t * ANYVAL = lns_class_new( ENV, sizeof( CLASS ) );
    // CLASS * CLASSVAL = lns_obj_##CLASS( pAny );                   
    // CLASSVAL->pMtd = &lns_mtd_##CLASS;


#define lns_class_new2_( ENV, CLASS, SCLASS, ANYVAL, CLASSVAL )      \
    lns_any_t * ANYVAL = lns_class_new( ENV, sizeof( CLASS ) );      \
    CLASS * CLASSVAL = lns_obj_##SCLASS( ANYVAL );          \
    CLASSVAL->pMtd = &lns_mtd_##SCLASS;                              \
    CLASSVAL->pMeta = &lns_type_meta_##SCLASS;
    

#define lns_decl_enum_get__text( ENUM, VALTYPE )                       \
    lns_stem_t ENUM##_get__txt( lns_env_t * _pEnv, lns_any_t * _pForm, \
                                   lns_##VALTYPE##_t val )             \
    {                                                                   \
        return lns_mtd_Map_get(                                        \
            _pEnv, ENUM##_val2NameMap, lns_##VALTYPE##2stem( _pEnv, val ) ); \
    }
#define lns_decl_enum_get__text_any( ENUM )                           \
    lns_stem_t ENUM##_get__txt( lns_env_t * _pEnv, lns_any_t * _pForm, \
                                   lns_stem_t val )             \
    {                                                                   \
        return lns_mtd_Map_get( _pEnv, ENUM##_val2NameMap, val );   \
    }

#define lns_init_if( ANY, ENV, OBJ, MTD, IF )                    \
    {                                                             \
        lns_any_t * __pAny = ANY;                                \
        lns_init_any( __pAny, lns_value_type_if, ENV );         \
        __pAny->refCount = 0;                                     \
        __pAny->val.ifVal.pMeta = &lns_type_meta_##IF;           \
        __pAny->val.ifVal.pObj = OBJ;                    \
        __pAny->val.ifVal.pMtd = MTD;                             \
    }

#define lns_init_any( ANY, TYPE, ENV )            \
    {                                                \
        lns_any_t * _pAny = ANY;                 \
        _pAny->type = TYPE;                         \
    }

#define lns_getImpObj( ANY ) (ANY)->val.ifVal.pObj


    typedef enum {
        lns_value_type_none,
        lns_value_type_block,
        lns_value_type_str,
        lns_value_type_class,
        lns_value_type_if,
        lns_value_type_ddd,
        lns_value_type_mRet,
        lns_value_type_form,
        //lns_value_type_List,
        //lns_value_type_Array,
        //lns_value_type_Set,
        //lns_value_type_Map,
        lns_value_type_itList,
        lns_value_type_itSet,
        lns_value_type_itMap,
        lns_value_type_alge,
        lns_value_type_luaVal,
        lns_value_type_luaForm,
    } lns_value_type_t;
    typedef struct lns_itList_t lns_itList_t;
    typedef struct lns_itSet_t lns_itSet_t;
    typedef struct lns_itMap_t lns_itMap_t;
    

    typedef lns_any_t * lns_newfunc_t( lns_env_t * _pEnv, ... );


    struct lns_fromVal_info_t;
        
    typedef lns_stem_t lns_fromMap_t( lns_env_t * _pEnv, const struct lns_fromVal_info_t * pInfoArray, lns_stem_t val);

    typedef struct lns_fromVal_info_t {
        bool nilable;
        lns_fromMap_t * pFromMap;
        const struct lns_fromVal_info_t * pInfoArray[];
    } lns_fromVal_info_t;

    

    /**
     * 関数の型
     *
     * @param pInfo フォーム情報
     * @param ... 関数の引数
     * @return 関数の戻り値
     */
    typedef lns_stem_t lns_closure_t( lns_env_t * _pEnv, lns_any_t * pInfo, ... );
    typedef lns_any_t * lns_closureAny_t( lns_env_t * _pEnv, lns_any_t * pInfo, ... );
    typedef lns_int_t lns_closureInt_t( lns_env_t * _pEnv, lns_any_t * pInfo, ... );
    typedef lns_real_t lns_closureReal_t( lns_env_t * _pEnv, lns_any_t * pInfo, ... );
    typedef bool lns_closureBool_t( lns_env_t * _pEnv, lns_any_t * pInfo, ... );
    
    typedef lns_stem_t lns_func_t( lns_env_t * _pEnv, ... );
    typedef lns_any_t * lns_funcAny_t( lns_env_t * _pEnv, ... );
    typedef lns_int_t lns_funcInt_t( lns_env_t * _pEnv, ... );
    typedef lns_real_t lns_funcReal_t( lns_env_t * _pEnv, ... );
    typedef bool lns_funcBool_t( lns_env_t * _pEnv, ... );

    
    /**
     * メソッドの型
     *
     * @param pObj クラスのインスタンスを保持する any
     * @param ... メソッドの引数
     * @return メソッドの戻り値
     */
    typedef lns_stem_t lns_method_t( lns_env_t * _pEnv, lns_any_t * pObj, ... );
    typedef lns_any_t * lns_method_any_t( lns_env_t * _pEnv, lns_any_t * pObj, ... );
    typedef lns_int_t lns_method_int_t( lns_env_t * _pEnv, lns_any_t * pObj, ... );
    typedef lns_real_t lns_method_real_t( lns_env_t * _pEnv, lns_any_t * pObj, ... );
    typedef bool lns_method_bool_t( lns_env_t * _pEnv, lns_any_t * pObj, ... );



    /**
     * クラスのインスタンスの gc の型
     *
     * @param pObj クラスのインスタンスを保持する any
     */
    typedef void lns_gc_t( lns_env_t * _pEnv, lns_any_t * pObj );

    /**
     * クラスのインスタンスの開放処理
     *
     * @param pObj クラスのインスタンスを保持する any
     */
    typedef void lns_del_t( lns_env_t * _pEnv, lns_any_t * pObj );

    /**
     * 型のメタ情報
     */
    typedef struct lns_type_meta_t {
        /** 型名 */
        const char * pName;
        /** 親クラスのメタ */
        struct lns_type_meta_t * pSuper;
        /**
           impliment している interface の Meta のリスト。
           NULL で終端する。
         */
        struct lns_type_meta_t * pIfList[];
    } lns_type_meta_t;

    /** 全クラスの super */
    extern lns_type_meta_t lns_type_meta_lns__root;
    
    /**
     * クラスのメソッドの最小構造。
     *
     * 各クラスのインスタンスは、メソッドのポインタを保持する構造体を持つ。
     * メソッドのポインタを保持する構造体の先頭は、
     * 必ず del, gc のメソッドでなければならない。
     */
    typedef struct lns_mtd__Class_t {
        lns_del_t * _del;
        lns_gc_t * _gc;
    } lns_mtd__Class_t;

    /**
     * クラスのインスタンスの最小構造。
     *
     * 各クラスのインスタンスは、
     * メソッドのポインタを保持する構造体のポインタ pMtd を先頭に持つ。
     * クラスのメンバは、 pMtd の次に宣言する。
     */
    typedef struct lns_Class_t {
        lns_type_meta_t * pMeta;
        void * pIFdummy;
        lns_mtd__Class_t * pMtd;
    } lns_Class_t;


    /**
     * alge 型を開放する際のコールバック
     */
    typedef void (lns_algeVal_gc_t)( lns_env_t * _pEnv, void * pVal );
    /**
     * alge 型のインスタンス。
     */
    typedef struct lns_Alge_t {
        int type;
        void * pVal;
        lns_algeVal_gc_t * gc;
    } lns_Alge_t;

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
    } lns_str_t;

    /**
     * ... を保持する型。
     *
     * 関数の戻り値の多値にも使用する。
     */
    typedef struct {
        /** ... に含む stem を管理するバッファ */
        lns_stem_t * stemList;
        /** pAnyList で管理している any の数 */
        int len;
    } lns_ddd_t;

    typedef struct {
        lns_value_type_t type;
        lns_any_t * pAny;
    } lns_luaVal_t;


#define lns_isClosure( FORM ) \
    ((FORM)->val.form.needFormParam)
#define lns_setClosure( FORM ) \
    (FORM)->val.form.needFormParam = true;
    
#define lns_closure( FORM )                  \
    (FORM)->val.form.pFunc
#define lns_closure_any( FORM )                  \
    ((lns_closureAny_t *)(FORM)->val.form.pFunc)
#define lns_closure_int( FORM )                  \
    ((lns_closureInt_t *)(FORM)->val.form.pFunc)
#define lns_closure_real( FORM )                  \
    ((lns_closureReal_t *)(FORM)->val.form.pFunc)
#define lns_closure_bool( FORM )                  \
    ((lns_closureBool_t *)(FORM)->val.form.pFunc)

#define lns_func( FORM )                  \
    ((lns_func_t *)(FORM)->val.form.pFunc)
#define lns_func_any( FORM )                  \
    ((lns_funcAny_t *)(FORM)->val.form.pFunc)
#define lns_func_int( FORM )                  \
    ((lns_funcInt_t *)(FORM)->val.form.pFunc)
#define lns_func_real( FORM )                  \
    ((lns_funcReal_t *)(FORM)->val.form.pFunc)
#define lns_func_bool( FORM )                  \
    ((lns_funcBool_t *)(FORM)->val.form.pFunc)
    
#define l_form_closure_var( FORM, INDEX )        \
    (FORM)->val.form.pClosureValList[ INDEX ].pVar
#define l_form_closure( FORM, INDEX )                        \
    &((FORM)->val.form.pClosureValList[ INDEX ])

    /** クロージャアクセスされる変数 */
    typedef struct lns_closureVar_t {
        //lns_any_t * pAny;
        lns_stem_t stem;
        int refCount;
    } lns_closureVar_t;

    /** クロージャ変数参照情報 */
    typedef struct lns_closureRef_t {
        lns_closureVar_t * pVar;
    } lns_closureRef_t;
    
    /**
     * form の情報。
     */
    struct lns_form_t {
        /** 関数 */
        lns_closure_t * pFunc;
        /** form 内でアクセスする外部変数を管理するバッファ */
        lns_var_t * pClosureValList;
        /** pClosureValList で管理している any の数 */
        int len;
        /** 関数コール時に、この構造体を必要とする場合 true*/
        bool needFormParam;
        /** 引数の数*/
        //int argNum;
        /** 引数に ... を持つかどうか */
        bool hasDDD;
    };

    typedef void lns_listObj_t;
    typedef struct lns_mtd_List_t lns_mtd_List_t;
    typedef struct lns_List_t {
        lns_type_meta_t * pMeta;
        void * pIFdummy;
        lns_mtd_List_t * pMtd;
        lns_listObj_t * pObj;
    } lns_List_t;
    extern lns_type_meta_t lns_type_meta_List;


    typedef void lns_setObj_t;
    typedef struct lns_mtd_Set_t lns_mtd_Set_t;
    typedef struct lns_Set_t {
        lns_type_meta_t * pMeta;
        void * pIFdummy;
        lns_mtd_Set_t * pMtd;
        lns_setObj_t * pObj;
    } lns_Set_t;
    extern lns_type_meta_t lns_type_meta_Set;

    typedef void lns_mapObj_t;
    typedef struct lns_mtd_Map_t lns_mtd_Map_t;
    typedef struct lns_Map_t {
        lns_type_meta_t * pMeta;
        void * pIFdummy;
        lns_mtd_Map_t * pMtd;
        lns_mapObj_t * pObj;
    } lns_Map_t;
    extern lns_type_meta_t lns_type_meta_Map;
    

    typedef struct lns_if_t {
        lns_type_meta_t * pMeta;
        lns_any_t * pObj;
        void * pMtd;
    } lns_if_t;

    
    /** any 型データ */
    struct lns_any_t {
        uint8_t state;
        /** 保持しているデータの型 */
        lns_value_type_t type;
        /** このデータを参照している数 */
        int refCount;
        /** 実データ */
        union {
            lns_str_t str;
            lns_ddd_t ddd;
            lns_form_t form;
            lns_Class_t * classVal;
            lns_itList_t * itList;
            lns_itSet_t * itSet;
            lns_itMap_t * itMap;
            lns_if_t ifVal;
            lns_Alge_t alge;
            lns_luaVal_t luaVal;
        } val;
        /**
         * オブジェクト同士をリンクする。
         *
         * major の場合
         *  - 全 major オブジェクトと繋ぐ双方向リンク。
         *  - s_majorObjTop からアクセスする。
         * minor の場合
         *  - 全 minor オブジェクトへの片方向リンク。 retObjTop で使用。
         *  - このオブジェクトを参照している var への varLink。
         */
        union {
            struct {
                struct lns_any_t * pNext;
                struct lns_any_t * pPrev;
            } major;
            struct {
                struct lns_any_t * pNext;
                struct lns_varLink_t * pLink;
            } minor;
        };
    };

    struct lns_varLink_t {
        /** gc で開放されなかった回数 */
        int age;
        /** newvar 同士をリンクする双方向リスト。 */
        struct lns_varLink_t * pNext;
        struct lns_varLink_t * pPrev;
        /** この newvar のオリジナル var */
        struct lns_var_t * pVar;
    };
    /** 任意の値を管理する変数 */
    typedef struct lns_var_t {
        lns_stem_t stem;

        /**
         * newvar のリスト。
         *
         * pLink が NULL の場合、 stem は major。
         * pLink が NULL 以外の場合、 stem は minor。
         */
        lns_varLink_t * pLink;
    } lns_var_t;
    /** 戻り値を管理する */
    typedef struct lns_retObj_t {
        lns_var_t var;

        struct lns_retObj_t * pNext;
    } lns_retObj_t;
    
    

    struct lns_block_t {
        /** このブロックの深度を判定するための目安 */
        void * pStackAddr;
        /** ブロック深度 */
        int blockDepth;
        /** このブロックで管理する var 型を保持するバッファ */
        lns_closureVar_t ** pVarList;
        /** このブロックで管理する stem 型を保持するバッファ */
        lns_stem_t ** pStemList;
        /** このブロックで管理する stem 型を保持するバッファ */
        lns_any_t ** pAnyList;
        /** pAnyList で管理する値の数 */
        int anyLen;
        /** pStemList で管理する値の数 */
        int stemLen;
        /** pVarBuf で管理する値の数 */
        int varLen;
        /**
         * このブロック内で生成された any 型データの双方向リスト。
         * 実際の先頭要素は managedAnyTop.pNext。
         */
        lns_any_t managedAnyTop;


        /** ブロック開始時の useStackVarNum */
        int useStackVarNum;
        
   };

    typedef struct {
        /** 値を返さない関数の戻り値 */
        lns_stem_t noneStem;
        /** nil */
        lns_stem_t nilStem;
        /** true */
        lns_stem_t trueStem;
        /** false */
        lns_stem_t falseStem;

        /** 要素 0 個の DDD*/
        lns_stem_t ddd0;

        /** nil の var */
        lns_var_t nilVar;
        
    } lns_global_t;

    extern lns_global_t lns_global;

    typedef void * (lns_testFunc)( lns_env_t * _pEnv );
    
    
    typedef struct lns_module_t {
        struct lns_module_t * pPrev;
        struct lns_module_t * pNext;

        /** モジュールを初期化済みかどうか */
        bool readyFlag;
        /** モジュールの init ブロック */
        lns_block_t * pBlock;
        /** モジュール名 */
        const char * pName;
        /** pending */
        lns_testFunc * pTestFunc[];
    } lns_module_t;

#define LNS_BLOCK_AT( ENV, DEPTH ) (&ENV->blockQueue[ ENV->blockDepth - DEPTH ])
    
    struct lua_State;
    
    struct lns_env_t {
        struct lua_State * pLua;
        lns_allocator_t allocateor;
        /**
         * ブロック情報で利用する pVarList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lns_closureVar_t * varPPool[ LNS_ANY_POOL_MAX_NUM ];
        /** varPPool をどこまで使用しているか個数を示す */
        int useVarPoolNum;
        /**
         * ブロック情報で利用する pAnyList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lns_stem_t * stemPPool[ LNS_ANY_POOL_MAX_NUM ];
        /** anyPPool をどこまで使用しているか個数を示す */
        int useStemPoolNum;
        /**
         * ブロック情報で利用する pAnyList のバッファ。
         * ブロック開始時に、ここから割り当てる。
         */
        lns_any_t * anyPPool[ LNS_ANY_POOL_MAX_NUM ];
        /** anyPPool をどこまで使用しているか個数を示す */
        int useAnyPoolNum;
        /** ブロック情報の Queue。*/
        lns_block_t blockQueue[ LNS_BLOCK_MAX_DEPTH ];
        /** 現在のブロックの深度 */
        int blockDepth;

        /** sort callback */
        lns_stem_t pSortCallback;

        /** 値のスタック */
        lns_stem_t pValStack[ LNS_VAL_STACK_MAX ];
        /** pValStack の現在位置 */
        int stackPos;

        /** 処理中の MRet */
        lns_any_t * pMRet;

        lns_module_t loadModuleTop;


        /** 戻り値 Queue の先頭 */
        lns_retObj_t retObjQueue;


        ///// 動的に alloc するデータ
        lns_envAlloc_t * pEnvAlloc;
    };


#define lns_set2DDDArg( ANY, INDEX, VAL )  \
    ANY->val.ddd.stemList[ INDEX ] = VAL;

#if 0
#define lns_fromDDD( ANY, INDEX )  \
//     (ANY->val.ddd.stemList[ INDEX ])
#else
#define lns_fromDDD( ANY, INDEX ) \
    lns_getValFromDDD( ANY, INDEX )
#endif
    

#define lns_lenDDD( ANY ) ANY->val.ddd.len

#define lns_str2any( ENV, VAL )               \
    _lns_str2any( LNS_DEBUG_POS, ENV, VAL )
#define lns_litStr2any( ENV, STR )            \
    _lns_litStr2any( LNS_DEBUG_POS, ENV, STR )
#define lns_cloneBin2any( ENV, STR, LEN )               \
    _lns_cloneBin2any( LNS_DEBUG_POS, ENV, STR, LEN )
#define lns_func2any( ENV, FUNC, ARGNUM, HASDDD, NUM, ... )            \
    _lns_func2any( LNS_DEBUG_POS, ENV, FUNC, ARGNUM, HASDDD, NUM, ##__VA_ARGS__ )
#define lns_createDDD( ENV, HASDDD, NUM, ... )         \
    _lns_createDDD( LNS_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__)
#define lns_createDDDOnly( ENV, NUM )          \
    _lns_createDDDOnly( LNS_DEBUG_POS, ENV, NUM )
#define lns_createSubDDD( ENV, OFFSET, DDD ) \
    _lns_createSubDDD( LNS_DEBUG_POS, ENV, OFFSET, DDD )
#define lns_createMRet( ENV, HASDDD, NUM, ... )        \
    _lns_createMRet( LNS_DEBUG_POS, ENV, HASDDD, NUM, ##__VA_ARGS__ )
#define lns_class_new( ENV, SIZE )             \
    _lns_class_new( LNS_DEBUG_POS, ENV, SIZE )
#define lns_alge_new( ENV, VALTYPE, SIZE, GC )                    \
    _lns_alge_new( LNS_DEBUG_POS, ENV, VALTYPE, SIZE, GC )
#define lns_it_new( ENV, TYPE, VAL )                  \
    _lns_it_new( LNS_DEBUG_POS, ENV, TYPE, VAL )
#define lns_createList( ENV, LIST )                  \
    _lns_createList( LNS_DEBUG_POS, ENV, LIST )
#define lns_createSet( ENV, SET )                  \
    _lns_createSet( LNS_DEBUG_POS, ENV, SET )
#define lns_createMap( ENV, ENTRY )                  \
    _lns_createMap( LNS_DEBUG_POS, ENV, ENTRY )
#define lns_createImmediateVal( ENV, VAL )                     \
    _lns_createImmediateVal( LNS_DEBUG_POS, ENV, VAL )
#define lns_luaVal_new( ENV, TYPE )            \
    _lns_luaVal_new( LNS_DEBUG_POS, ENV, TYPE )
    
    

#define lns_abort( MESS ) _lns_abort( MESS, LNS_DEBUG_POS )


    typedef lns_stem_t (lns_mtd_getter_t)( lns_env_t * _pEnv, lns_any_t * pObj);
    typedef lns_any_t * (lns_mtd_getter_any_t)( lns_env_t * _pEnv, lns_any_t * pObj);
    typedef lns_int_t (lns_mtd_getter_int_t)( lns_env_t * _pEnv, lns_any_t * pObj);
    typedef lns_real_t (lns_mtd_getter_real_t)( lns_env_t * _pEnv, lns_any_t * pObj);
    typedef bool (lns_mtd_getter_bool_t)( lns_env_t * _pEnv, lns_any_t * pObj);

    extern lns_stem_t l_nil_mtd_getter(
        lns_env_t * _pEnv, lns_stem_t obj, lns_mtd_getter_t * pGetter );
    extern lns_stem_t l_nil_mtd_getter_any(
        lns_env_t * _pEnv, lns_stem_t obj, lns_mtd_getter_any_t * pGetter );
    extern lns_stem_t l_nil_mtd_getter_int(
        lns_env_t * _pEnv, lns_stem_t obj, lns_mtd_getter_int_t * pGetter );
    extern lns_stem_t l_nil_mtd_getter_real(
        lns_env_t * _pEnv, lns_stem_t obj, lns_mtd_getter_real_t * pGetter );
    extern lns_stem_t l_nil_mtd_getter_bool(
        lns_env_t * _pEnv, lns_stem_t obj, lns_mtd_getter_bool_t * pGetter );

    

    extern void _lns_abort( const char * pMessage, const char * pFile, int lineNo );
    
    extern lns_any_t * _lns_str2any( LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_str_t val );
    extern lns_any_t * _lns_litStr2any( LNS_DEBUG_DECL, lns_env_t * _pEnv, const char * pStr );
    extern lns_any_t * _lns_func2any( LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_closure_t * pFunc, int argNum, bool hasDDD, int num, ... );
    extern lns_stem_t _lns_createDDD( LNS_DEBUG_DECL, lns_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lns_stem_t _lns_createDDDOnly( LNS_DEBUG_DECL, lns_env_t * _pEnv, int num );
    extern lns_stem_t _lns_createSubDDD(
        const char * pFile, int lineNo, lns_env_t * _pEnv, int offset, lns_stem_t pDDD );
    extern lns_stem_t _lns_createMRet( LNS_DEBUG_DECL, lns_env_t * _pEnv, bool hasDDD, int num, ... );
    extern lns_any_t * _lns_class_new( LNS_DEBUG_DECL, lns_env_t * _pEnv, int size );
    extern lns_any_t * _lns_alge_new( LNS_DEBUG_DECL, lns_env_t * _pEnv, int valType, int size, lns_algeVal_gc_t * gc );
    extern lns_any_t * _lns_it_new(
        LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_value_type_t type, void * pVal );
    extern lns_any_t * _lns_cloneBin2any(
        const char * pFile, int lineNo, lns_env_t * _pEnv, const void * pBuf, int len );

    extern lns_any_t * _lns_luaVal_new(
        const char * pFile, int lineNo, lns_env_t * _pEnv, lns_value_type_t type );

    extern lns_any_t * lns_createMRetOnly( lns_env_t * _pEnv, int num );

    extern lns_int_t lns_stem2int( lns_stem_t stem );
    extern lns_real_t lns_stem2real( lns_stem_t stem );
    extern lns_bool_t lns_stem2bool( lns_stem_t stem );

    extern void lns_init_alge( lns_stem_t * pStem, lns_any_t * pAny, int valType );

    
    extern lns_any_t * _lns_createList( LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pList );
    extern lns_any_t * _lns_createSet(
        LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pSet );
    extern lns_any_t * _lns_createMap(
        LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdEntry_t * pEntry );

    extern lns_stem_t _lns_createImmediateVal(
        LNS_DEBUG_DECL, lns_env_t * _pEnv, lns_imdVal_t * pVal );
    

    extern lns_any_t * lns_strconcat(
        lns_env_t * _pEnv, lns_any_t * pAny1, lns_any_t * pAny2 );
    extern lns_any_t * lns_luaForm_new( lns_env_t * _pEnv );
    

    extern lns_str_t lns_createLiteralStr( const char * pStr );

    extern lns_closureVar_t * lns_var_alloc(
        lns_env_t * _pEnv, lns_block_t * pBlock, int index, lns_stem_type_t type );

    extern lns_block_t * lns_enter_module( lns_env_t * _pEnv, int anyNum, int stemNum, int varNum );
    
    extern void lns_reset_block( lns_env_t * _pEnv );

    extern void lns_setRet( lns_env_t * _pEnv, lns_stem_t stem );
    extern void lns_setRetInBlock( lns_env_t * _pEnv, lns_stem_t stem );
    extern void lns_setRetAtBlock( lns_block_t * pBlock, lns_stem_t pStem );
    extern void lns_setQ_( lns_any_t * pAny );
    extern void lns_setOverwrite( lns_any_t * pAny );
    extern lns_block_t * lns_enter_func( lns_env_t * _pEnv, int anyNum, int stemNum, int varNum, int argNum, ... );
    extern void lns_leave_block( lns_env_t * _pEnv, int block );
    extern void lns_leave_blockMulti( lns_env_t * _pEnv, int num );
    extern int lns_enter_block( lns_env_t * _pEnv );
    extern bool lns_decre_ref( lns_env_t * _pEnv, lns_any_t * pAny );
    extern void lns_it_delete( lns_env_t * _pEnv, lns_any_t * pAny );
    extern lns_stem_t lns_call_form( lns_env_t * _pEnv, lns_any_t * _pForm, lns_stem_t ddd );
    extern lns_stem_t lns_nil_call_form( lns_env_t * _pEnv, lns_stem_t form, lns_stem_t ddd );
    
    extern lns_any_t * lns_getIF( lns_env_t * _pEnv, lns_any_t * pIFAny );
    extern lns_stem_t lns_toIF( lns_env_t * _pEnv, lns_any_t * pAny, const lns_type_meta_t * pMeta );
    extern lns_stem_t lns_setMRet( lns_env_t * _pEnv, lns_any_t * pAny );
    extern lns_stem_t lns_getValFromDDD( lns_any_t * pAny, int index );
    extern lns_stem_t lns_getValFromDDDStem( lns_stem_t stem, int index );

    extern lns_varLink_t * lns_newVarLink( lns_var_t * pVar );
    extern void lns_initVar( lns_var_t * pVar, lns_stem_t stem );
    extern void lns_unuseVar( lns_var_t * pVar );



    extern lns_stem_t lns_unwrap_stem( lns_stem_t pAny, lns_stem_t pDefVal );
    extern lns_any_t * lns_unwrap_any( lns_stem_t pAny, lns_stem_t pDefVal );
    extern lns_int_t lns_unwrap_int( lns_stem_t pAny );
    extern lns_int_t lns_unwrap_intDefault( lns_stem_t pAny, lns_int_t val );
    extern lns_real_t lns_unwrap_real( lns_stem_t pAny );
    extern lns_real_t lns_unwrap_realDefault( lns_stem_t pAny, lns_real_t val );

    

    extern bool lns_isCondTrue( const lns_stem_t pAny );
    extern bool lns_incStack( lns_env_t * _pEnv );
    extern bool lns_setStackVal( lns_env_t * _pEnv, lns_stem_t pVal );
    extern lns_stem_t lns_popVal( lns_env_t * _pEnv, bool dummy );
    extern bool lns_op_not( lns_env_t * _pEnv, lns_stem_t pAny );
    extern bool lns_equals( lns_stem_t stem1, lns_stem_t stem2 );
    extern bool lns_equals_any( const lns_any_t * pAny1, const lns_any_t * pAny2 );

    extern lns_stem_t lns_castAny( lns_stem_t stem, lns_value_type_t kind );
    extern lns_stem_t lns_castStem( lns_stem_t stem, lns_stem_type_t kind );
    extern lns_stem_t lns_castIf( lns_env_t * _pEnv, lns_stem_t stem, const lns_type_meta_t * pMeta );
    extern lns_stem_t lns_castClass( lns_stem_t stem, const lns_type_meta_t * pMeta );

    extern lns_stem_t lns_refFieldNil(
        lns_env_t * _pEnv, lns_stem_t obj, int offset, lns_stem_type_t stemType );
    extern lns_stem_t lns_refItemNil( lns_env_t * _pEnv, lns_stem_t prefix, lns_stem_t index );
    extern lns_stem_t lns_stem_refAt( lns_env_t * _pEnv, lns_stem_t stem, lns_stem_t key );

    typedef lns_any_t * lns_toMap_t( lns_env_t * _pEnv, lns_any_t * pAny );
    
    extern lns_stem_t lns_toMapFromStem( lns_env_t * _pEnv, lns_stem_t stem );
    extern lns_stem_t lns_toMapFromList( lns_env_t * _pEnv, lns_stem_t stem );
    extern lns_stem_t lns_toMapFromSet( lns_env_t * _pEnv, lns_stem_t stem );
    extern lns_stem_t lns_toMapFromMap( lns_env_t * _pEnv, lns_stem_t stem );

    extern lns_stem_t lns_fromMapToStemSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToIntSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToRealSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToBoolSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToStrSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    
    extern lns_stem_t lns_fromMapToListSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToSetSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    extern lns_stem_t lns_fromMapToMapSub(
        lns_env_t * _pEnv, const lns_fromVal_info_t * pInfoArray, lns_stem_t stem );
    

    extern lns_any_t * lns_string_format( lns_env_t * _pEnv, const char * pFormat, lns_stem_t ddd );

    extern void lns_run_module( lns_env_t * _pEnv );



    extern void lns_f_print( lns_env_t * _pEnv, lns_stem_t pArg );

#ifdef __cplusplus
}
#endif

#include <lns_collection.h>
#include <lns_luaWrapper.h>
#include <lns_builtin.h>
#include <lns_envAlloc.h>

#endif


