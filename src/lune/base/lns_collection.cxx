#include <lunescript.h>
#include <vector>

using namespace std;

#if 0
{
   "List<T>": {
      "insert": { "type": [ "mut" ], "arg": [ "&T" ], "ret": [""] },
      "remove": { "type": [ "mut" ], "arg": [ "int!" ], "ret": ["T!"] },
      "unpack": { "type": [ "method" ], "arg": [ ], "ret": ["..."] },
      "sort": { "type": [ "mut" ], "arg": [ "form!" ], "ret": [] },
   },
},
{
   "Array<T>": {
      "unpack": { "type": [ "method" ], "arg": [ ], "ret": ["..."] },
      "sort": { "type": [ "mut" ], "arg": [ "form!" ], "ret": [] },
   },
},
{
   "Set<T>": {
      "add": { "type": [ "mut" ], "arg": [ "T" ], "ret": [] },
      "del": { "type": [ "mut" ], "arg": [ "T" ], "ret": [] },
      "has": { "type": [ "method" ], "arg": [ "T" ], "ret": [ "bool" ] },
      "and": { "type": [ "mut" ], "arg": [ "&Set<T>" ], "ret": [ "Set<T>" ] },
      "or": { "type": [ "mut" ], "arg": [ "&Set<T>" ], "ret": [ "Set<T>" ] },
      "sub": { "type": [ "mut" ], "arg": [ "&Set<T>" ], "ret": [ "Set<T>" ] },
      "clone": { "type": [ "method" ], "arg": [], "ret": [ "Set<T>" ] },
      "len": { "type": [ "method" ], "arg": [], "ret": [ "int" ] },
   },
}
#endif

#define __lune_obj_List( OBJ ) \
    ((__lune_List_t*)OBJ->val.classVal)

#define __lune_obj_List_obj( OBJ ) \
    ((vector<__lune_stem_t*>*)((__lune_List_t*)OBJ->val.classVal)->pObj)


static void __mtd_List_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag );
static __lune_stem_t * __mtd_List_insert(
    __lune_env_t * _pEnv, __lune_stem_t * pListStem, __lune_stem_t * pVal );
static __lune_stem_t * __mtd_List_refAt(
    __lune_env_t * _pEnv, __lune_stem_t * pListStem, int index );
static __lune_stem_t * __mtd_List_unpack(
    __lune_env_t * _pEnv, __lune_stem_t * pListStem );

__mtd_List_t __mtd_List = {
    __mtd_List_gc,
    (__lune_method_t*)__mtd_List_insert,
    (__lune_method_t*)__mtd_List_refAt,
    (__lune_method_t*)__mtd_List_unpack,
};


void __class_List_init( __lune_env_t * _pEnv, __lune_List_t * pObj )
{
    pObj->pObj = new vector<__lune_stem_t*>();
}


__lune_stem_t * __class_List_new( __lune_env_t * _pEnv )
{
    __lune_class_new2_( _pEnv, __lune_List_t, List, pStem, pObj );
    
    __class_List_init( _pEnv, pObj );

    return pStem;
}

static void __mtd_List_gc( __lune_env_t * _pEnv, __lune_stem_t * pObj, bool freeFlag )
{
    printf( "%s\n", __func__ );

    vector<__lune_stem_t*>::iterator it;
    vector<__lune_stem_t*>::iterator end = __lune_obj_List_obj( pObj )->end();
    
    for ( it = __lune_obj_List_obj( pObj )->begin(); it != end; it++ )
    {
        __lune_decre_ref( _pEnv, *it );
    }
    
    delete __lune_obj_List_obj( pObj );

    if ( freeFlag ) {
        __lune_class_del( _pEnv, __lune_obj_List( pObj ) );
    }
}

static __lune_stem_t * __mtd_List_insert(
    __lune_env_t * _pEnv, __lune_stem_t * pObj, __lune_stem_t * pVal )
{
    // pVal の参照カウントをインクリメントするため、 work に setq する
    __lune_stem_t * pWork;
    __lune_setQ( pWork, pVal );
    
    __lune_obj_List_obj( pObj )->push_back( pVal );
    
    return NULL;
}

static __lune_stem_t * __mtd_List_refAt(
    __lune_env_t * _pEnv, __lune_stem_t * pObj, int index )
{
    return __lune_obj_List_obj( pObj )->at( index - 1 );
}

static __lune_stem_t * __mtd_List_unpack(
    __lune_env_t * _pEnv, __lune_stem_t * pObj )
{
    __lune_stem_t * pDDD =
        __lune_createDDDOnly( _pEnv, __lune_obj_List_obj( pObj )->size() );

    vector<__lune_stem_t*>::iterator it;
    vector<__lune_stem_t*>::iterator end = __lune_obj_List_obj( pObj )->end();

    int index = 0;
    for ( it = __lune_obj_List_obj( pObj )->begin(); it != end; it++, index++ )
    {
        __lune_set2DDDArg( pDDD, index, *it );
    }

    return pDDD;
}
