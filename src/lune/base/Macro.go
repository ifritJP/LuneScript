// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Macro bool
var Macro__mod__ string
var Macro_toListEmpty *Lns_luaValue
var Macro_toLuaval *Lns_luaValue
type Macro_toListEmptyLua func (_env *LnsEnv) *LnsList
type Macro_toLuavalLua func (_env *LnsEnv, arg1 LnsAny) LnsAny
type Macro_EvalMacroCallback func (_env *LnsEnv)
// for 39
func Macro_convExp44(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 37
func Macro_convExp27(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 356
func Macro_convExp1296(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 445
func Macro_convExp1730(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 36: decl @lune.@base.@Macro.loadCode
func Macro_loadCode_1008_(_env *LnsEnv, code string) LnsAny {
    var loaded LnsAny
    var mess LnsAny
    loaded,mess = _env.CommonLuaVM.Load(code, nil)
    if loaded != nil{
        loaded_18 := loaded.(*Lns_luaValue)
        {
            _obj := Macro_convExp44(Lns_2DDD(_env.CommonLuaVM.RunLoadedfunc(loaded_18,Lns_2DDD([]LnsAny{}))[0]))
            if !Lns_IsNil( _obj ) {
                obj := _obj
                return obj
            }
        }
        panic("failed to load")
    }
    panic(_env.LuaVM.String_format("%s -- %s", []LnsAny{mess, code}))
// insert a dummy
    return nil
}

// 114: decl @lune.@base.@Macro.getLiteralMacroVal
func Macro_getLiteralMacroVal_1150_(_env *LnsEnv, obj LnsAny) LnsAny {
    switch _exp578 := obj.(type) {
    case *Nodes_Literal__Nil:
        return nil
    case *Nodes_Literal__Int:
    val := _exp578.Val1
        return val
    case *Nodes_Literal__Real:
    val := _exp578.Val1
        return val
    case *Nodes_Literal__Str:
    val := _exp578.Val1
        return val
    case *Nodes_Literal__Bool:
    val := _exp578.Val1
        return val
    case *Nodes_Literal__Symbol:
    val := _exp578.Val1
        return NewLnsList([]LnsAny{val})
    case *Nodes_Literal__Field:
    val := _exp578.Val1
        return val
    case *Nodes_Literal__LIST:
    list := _exp578.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_1150_(_env, item))
        }
        return newList
    case *Nodes_Literal__ARRAY:
    list := _exp578.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_1150_(_env, item))
        }
        return newList
    case *Nodes_Literal__SET:
    list := _exp578.Val1
        var newSet *LnsSet
        newSet = NewLnsSet([]LnsAny{})
        for _, _item := range( list.Items ) {
            item := _item
            {
                __exp := Macro_getLiteralMacroVal_1150_(_env, item)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp
                    newSet.Add(_exp)
                }
            }
        }
        return newSet
    case *Nodes_Literal__MAP:
    _map := _exp578.Val1
        var newMap *LnsMap
        newMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _key, _val := range( _map.Items ) {
            key := _key
            val := _val
            var keyObj LnsAny
            keyObj = Macro_getLiteralMacroVal_1150_(_env, key)
            var valObj LnsAny
            valObj = Macro_getLiteralMacroVal_1150_(_env, val)
            if keyObj != nil && valObj != nil{
                keyObj_120 := keyObj
                valObj_121 := valObj
                newMap.Set(keyObj_120,valObj_121)
            }
        }
        return newMap
    }
// insert a dummy
    return nil
}

// 288: decl @lune.@base.@Macro.equalsType
func Macro_equalsType_1404_(_env *LnsEnv, typeInfo *Ast_TypeInfo,builtinType *Ast_TypeInfo) bool {
    return typeInfo.FP.Get_srcTypeInfo(_env) == builtinType
}


// 544: decl @lune.@base.@Macro.expandVal
func Macro_expandVal_1649_(_env *LnsEnv, tokenList *LnsList,workval LnsAny,pos *Types_Position) LnsAny {
    if workval != nil{
        workval_344 := workval
        var val LnsAny
        val = workval_344
        if _switch2546 := Lns_type(val); _switch2546 == "boolean" {
            var token string
            token = _env.LuaVM.String_format("%s", []LnsAny{val})
            var kind LnsInt
            kind = Types_TokenKind__Kywd
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, kind, token, pos, false, nil)))
        } else if _switch2546 == "number" {
            var num string
            num = _env.LuaVM.String_format("%g", []LnsAny{Lns_forceCastReal(val)})
            var kind LnsInt
            kind = Types_TokenKind__Int
            if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(num, ".", 1, true))){
                kind = Types_TokenKind__Real
                
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, kind, num, pos, false, nil)))
        } else if _switch2546 == "string" {
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, Types_TokenKind__Str, Parser_quoteStr(_env, val.(string)), pos, false, nil)))
        } else {
            return _env.LuaVM.String_format("not support ,, List -- %s", []LnsAny{Lns_type(val)})
        }
    }
    return nil
}

// 590: decl @lune.@base.@Macro.pushbackTxt
func Macro_pushbackTxt_1676_(_env *LnsEnv, pushbackParser Parser_PushbackParser,txtList *LnsList,streamName string,pos *Types_Position) {
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    for _, _txt := range( txtList.Items ) {
        txt := _txt.(string)
        var parser *Parser_StreamParser
        parser = Parser_StreamParser_create(_env, &Types_ParserSrc__LnsCode{txt, _env.LuaVM.String_format("macro symbol -- %s", []LnsAny{streamName})}, nil, pos.FP.Get_RawOrgPos(_env))
        var workParser *Parser_DefaultPushbackParser
        workParser = NewParser_DefaultPushbackParser(_env, &parser.Parser_Parser)
        for  {
            var worktoken *Types_Token
            worktoken = workParser.FP.GetTokenNoErr(_env)
            if worktoken.Kind == Types_TokenKind__Eof{
                break
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, worktoken.Kind, worktoken.Txt, pos, false, nil)))
        }
    }
    {
        var _from2692 LnsInt = tokenList.Len()
        var _to2692 LnsInt = 1
        _work2692 := _from2692
        _delta2692 := -1
        for {
            if _delta2692 > 0 {
               if _work2692 > _to2692 { break }
            } else {
               if _work2692 < _to2692 { break }
            }
            index := _work2692
            pushbackParser.PushbackToken(_env, tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work2692 += _delta2692
        }
    }
}


// 635: decl @lune.@base.@Macro.MacroCtrl.expandMacroVal.macroVal2strList
func MacroCtrl_expandMacroVal__macroVal2strList_1716_(_env *LnsEnv, name string,macroVal *Nodes_MacroValInfo,workParser Parser_PushbackParser) *LnsList {
    var val LnsAny
    
    {
        _val := macroVal.Val
        if _val == nil{
            workParser.Error(_env, _env.LuaVM.String_format("macroVal is nil -- %s", []LnsAny{name}))
        } else {
            val = _val
        }
    }
    if Lns_isCondTrue( macroVal.ArgNode){
        return val.(*LnsList)
    }
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    {
        __collection2837 := val.(*LnsMap)
        __sorted2837 := __collection2837.CreateKeyListInt()
        __sorted2837.Sort( LnsItemKindInt, nil )
        for _, ___key2837 := range( __sorted2837.Items ) {
            item := __collection2837.Items[ ___key2837 ].(string)
            list.Insert(item)
        }
    }
    return list
}

// 962: decl @lune.@base.@Macro.nodeToCodeTxt
func Macro_nodeToCodeTxt(_env *LnsEnv, node *Nodes_Node,moduleTypeInfo *Ast_TypeInfo) string {
    var code string
    Lns_LockEnvSync( _env, func () {
        var memStream *Util_memStream
        memStream = NewUtil_memStream(_env)
        var formatter *Nodes_Filter
        formatter = Formatter_createFilter(_env, moduleTypeInfo, memStream.FP)
        node.FP.ProcessFilter(_env, formatter, Formatter_Opt2Stem(NewFormatter_Opt(_env, node)))
        code = memStream.FP.Get_txt(_env)
        
    })
    return code
}

// declaration Class -- MacroMetaArgInfo
type Macro_MacroMetaArgInfoMtd interface {
    ToMap() *LnsMap
}
type Macro_MacroMetaArgInfo struct {
    Name string
    TypeId LnsInt
    FP Macro_MacroMetaArgInfoMtd
}
func Macro_MacroMetaArgInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroMetaArgInfo).FP
}
type Macro_MacroMetaArgInfoDownCast interface {
    ToMacro_MacroMetaArgInfo() *Macro_MacroMetaArgInfo
}
func Macro_MacroMetaArgInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroMetaArgInfoDownCast)
    if ok { return work.ToMacro_MacroMetaArgInfo() }
    return nil
}
func (obj *Macro_MacroMetaArgInfo) ToMacro_MacroMetaArgInfo() *Macro_MacroMetaArgInfo {
    return obj
}
func NewMacro_MacroMetaArgInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) *Macro_MacroMetaArgInfo {
    obj := &Macro_MacroMetaArgInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaArgInfo(_env, arg1, arg2)
    return obj
}
func (self *Macro_MacroMetaArgInfo) InitMacro_MacroMetaArgInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) {
    self.Name = arg1
    self.TypeId = arg2
}
func (self *Macro_MacroMetaArgInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["typeId"] = Lns_ToCollection( self.TypeId )
    return obj
}
func (self *Macro_MacroMetaArgInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Macro_MacroMetaArgInfo__fromMap_1053_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaArgInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaArgInfo__fromStem_1057_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaArgInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaArgInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Macro_MacroMetaArgInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Macro_MacroMetaArgInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Macro_MacroMetaArgInfo{}
    newObj.FP = newObj
    return Macro_MacroMetaArgInfo_FromMapMain( newObj, objMap, paramList )
}
func Macro_MacroMetaArgInfo_FromMapMain( newObj *Macro_MacroMetaArgInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["typeId"], false, nil); !ok {
       return false,nil,"typeId:" + mess.(string)
    } else {
       newObj.TypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- MacroMetaInfo
type Macro_MacroMetaInfoMtd interface {
    ToMap() *LnsMap
}
type Macro_MacroMetaInfo struct {
    Name string
    Pos *LnsList
    ArgList *LnsList
    SymList *LnsList
    StmtBlock LnsAny
    TokenList *LnsList
    FP Macro_MacroMetaInfoMtd
}
func Macro_MacroMetaInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroMetaInfo).FP
}
type Macro_MacroMetaInfoDownCast interface {
    ToMacro_MacroMetaInfo() *Macro_MacroMetaInfo
}
func Macro_MacroMetaInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroMetaInfoDownCast)
    if ok { return work.ToMacro_MacroMetaInfo() }
    return nil
}
func (obj *Macro_MacroMetaInfo) ToMacro_MacroMetaInfo() *Macro_MacroMetaInfo {
    return obj
}
func NewMacro_MacroMetaInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) *Macro_MacroMetaInfo {
    obj := &Macro_MacroMetaInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Macro_MacroMetaInfo) InitMacro_MacroMetaInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) {
    self.Name = arg1
    self.Pos = arg2
    self.ArgList = arg3
    self.SymList = arg4
    self.StmtBlock = arg5
    self.TokenList = arg6
}
func (self *Macro_MacroMetaInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["pos"] = Lns_ToCollection( self.Pos )
    obj.Items["argList"] = Lns_ToCollection( self.ArgList )
    obj.Items["symList"] = Lns_ToCollection( self.SymList )
    obj.Items["stmtBlock"] = Lns_ToCollection( self.StmtBlock )
    obj.Items["tokenList"] = Lns_ToCollection( self.TokenList )
    return obj
}
func (self *Macro_MacroMetaInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Macro_MacroMetaInfo__fromMap_1093_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaInfo__fromStem_1097_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Macro_MacroMetaInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Macro_MacroMetaInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Macro_MacroMetaInfo{}
    newObj.FP = newObj
    return Macro_MacroMetaInfo_FromMapMain( newObj, objMap, paramList )
}
func Macro_MacroMetaInfo_FromMapMain( newObj *Macro_MacroMetaInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["pos"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"pos:" + mess.(string)
    } else {
       newObj.Pos = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["argList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Macro_MacroMetaArgInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"argList:" + mess.(string)
    } else {
       newObj.ArgList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["symList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Macro_MacroMetaArgInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"symList:" + mess.(string)
    } else {
       newObj.SymList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["stmtBlock"], true, nil); !ok {
       return false,nil,"stmtBlock:" + mess.(string)
    } else {
       newObj.StmtBlock = conv
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["tokenList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToListSub, false,[]Lns_ToObjParam{Lns_ToObjParam{
                    Lns_ToStemSub, false,nil}}}}); !ok {
       return false,nil,"tokenList:" + mess.(string)
    } else {
       newObj.TokenList = conv.(*LnsList)
    }
    return true, newObj, nil
}

// declaration Class -- MacroParser
type Macro_MacroParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Macro_MacroParser struct {
    Parser_Parser
    tokenList *LnsList
    pos LnsInt
    name string
    overridePos LnsAny
    FP Macro_MacroParserMtd
}
func Macro_MacroParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroParser).FP
}
type Macro_MacroParserDownCast interface {
    ToMacro_MacroParser() *Macro_MacroParser
}
func Macro_MacroParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroParserDownCast)
    if ok { return work.ToMacro_MacroParser() }
    return nil
}
func (obj *Macro_MacroParser) ToMacro_MacroParser() *Macro_MacroParser {
    return obj
}
func NewMacro_MacroParser(_env *LnsEnv, arg1 *LnsList, arg2 string, arg3 LnsAny) *Macro_MacroParser {
    obj := &Macro_MacroParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitMacro_MacroParser(_env, arg1, arg2, arg3)
    return obj
}
// 76: DeclConstr
func (self *Macro_MacroParser) InitMacro_MacroParser(_env *LnsEnv, tokenList *LnsList,name string,overridePos LnsAny) {
    self.InitParser_Parser(_env)
    self.pos = 1
    
    self.tokenList = tokenList
    
    self.name = name
    
    self.overridePos = overridePos
    
}

// 86: decl @lune.@base.@Macro.MacroParser.createPosition
func (self *Macro_MacroParser) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) *Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}

// 91: decl @lune.@base.@Macro.MacroParser.getToken
func (self *Macro_MacroParser) GetToken(_env *LnsEnv) LnsAny {
    if self.tokenList.Len() < self.pos{
        return nil
    }
    var token *Types_Token
    token = self.tokenList.GetAt(self.pos).(Types_TokenDownCast).ToTypes_Token()
    self.pos = self.pos + 1
    
    {
        __ := self.overridePos
        if !Lns_IsNil( __ ) {
            return NewTypes_Token(_env, token.Kind, token.Txt, self.FP.CreatePosition(_env, token.Pos.LineNo, token.Pos.Column), token.Consecutive, token.FP.Get_commentList(_env))
        }
    }
    return token
}

// 108: decl @lune.@base.@Macro.MacroParser.getStreamName
func (self *Macro_MacroParser) GetStreamName(_env *LnsEnv) string {
    return self.name
}


// declaration Class -- ExtMacroInfo
type Macro_ExtMacroInfoMtd interface {
    GetArgList(_env *LnsEnv) *LnsList
    GetTokenList(_env *LnsEnv) *LnsList
    Get_name(_env *LnsEnv) string
}
type Macro_ExtMacroInfo struct {
    Nodes_MacroInfo
    name string
    argList *LnsList
    tokenList *LnsList
    FP Macro_ExtMacroInfoMtd
}
func Macro_ExtMacroInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_ExtMacroInfo).FP
}
type Macro_ExtMacroInfoDownCast interface {
    ToMacro_ExtMacroInfo() *Macro_ExtMacroInfo
}
func Macro_ExtMacroInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_ExtMacroInfoDownCast)
    if ok { return work.ToMacro_ExtMacroInfo() }
    return nil
}
func (obj *Macro_ExtMacroInfo) ToMacro_ExtMacroInfo() *Macro_ExtMacroInfo {
    return obj
}
func NewMacro_ExtMacroInfo(_env *LnsEnv, arg1 string, arg2 *Lns_luaValue, arg3 *LnsMap, arg4 *LnsList, arg5 *LnsList) *Macro_ExtMacroInfo {
    obj := &Macro_ExtMacroInfo{}
    obj.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitMacro_ExtMacroInfo(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Macro_ExtMacroInfo) Get_name(_env *LnsEnv) string{ return self.name }
// 179: decl @lune.@base.@Macro.ExtMacroInfo.getArgList
func (self *Macro_ExtMacroInfo) GetArgList(_env *LnsEnv) *LnsList {
    return self.argList
}

// 182: decl @lune.@base.@Macro.ExtMacroInfo.getTokenList
func (self *Macro_ExtMacroInfo) GetTokenList(_env *LnsEnv) *LnsList {
    return self.tokenList
}

// 186: DeclConstr
func (self *Macro_ExtMacroInfo) InitMacro_ExtMacroInfo(_env *LnsEnv, name string,_func *Lns_luaValue,symbol2MacroValInfoMap *LnsMap,argList *LnsList,tokenList *LnsList) {
    self.InitNodes_MacroInfo(_env, _func, symbol2MacroValInfoMap)
    self.name = name
    
    self.argList = argList
    
    self.tokenList = tokenList
    
}


// declaration Class -- MacroAnalyzeInfo
type Macro_MacroAnalyzeInfoMtd interface {
    EqualsArgTypeList(_env *LnsEnv, arg1 LnsAny) bool
    GetCurArgType(_env *LnsEnv) *Ast_TypeInfo
    Get_argIndex(_env *LnsEnv) LnsInt
    Get_mode(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    IsAnalyzingExpArg(_env *LnsEnv) bool
    IsAnalyzingSymArg(_env *LnsEnv) bool
    NextArg(_env *LnsEnv)
}
type Macro_MacroAnalyzeInfo struct {
    typeInfo *Ast_TypeInfo
    mode LnsInt
    argIndex LnsInt
    FP Macro_MacroAnalyzeInfoMtd
}
func Macro_MacroAnalyzeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroAnalyzeInfo).FP
}
type Macro_MacroAnalyzeInfoDownCast interface {
    ToMacro_MacroAnalyzeInfo() *Macro_MacroAnalyzeInfo
}
func Macro_MacroAnalyzeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroAnalyzeInfoDownCast)
    if ok { return work.ToMacro_MacroAnalyzeInfo() }
    return nil
}
func (obj *Macro_MacroAnalyzeInfo) ToMacro_MacroAnalyzeInfo() *Macro_MacroAnalyzeInfo {
    return obj
}
func NewMacro_MacroAnalyzeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Macro_MacroAnalyzeInfo {
    obj := &Macro_MacroAnalyzeInfo{}
    obj.FP = obj
    obj.InitMacro_MacroAnalyzeInfo(_env, arg1, arg2)
    return obj
}
func (self *Macro_MacroAnalyzeInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *Macro_MacroAnalyzeInfo) Get_mode(_env *LnsEnv) LnsInt{ return self.mode }
func (self *Macro_MacroAnalyzeInfo) Get_argIndex(_env *LnsEnv) LnsInt{ return self.argIndex }
// 203: DeclConstr
func (self *Macro_MacroAnalyzeInfo) InitMacro_MacroAnalyzeInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,mode LnsInt) {
    self.typeInfo = typeInfo
    
    self.mode = mode
    
    self.argIndex = 1
    
}

// 209: decl @lune.@base.@Macro.MacroAnalyzeInfo.equalsArgTypeList
func (self *Macro_MacroAnalyzeInfo) EqualsArgTypeList(_env *LnsEnv, argTypeList LnsAny) bool {
    return self.typeInfo.FP.Get_argTypeInfoList(_env) == argTypeList
}

// 212: decl @lune.@base.@Macro.MacroAnalyzeInfo.getCurArgType
func (self *Macro_MacroAnalyzeInfo) GetCurArgType(_env *LnsEnv) *Ast_TypeInfo {
    if self.typeInfo.FP.Get_argTypeInfoList(_env).Len() < self.argIndex{
        return Ast_builtinTypeNone
    }
    return self.typeInfo.FP.Get_argTypeInfoList(_env).GetAt(self.argIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}

// 218: decl @lune.@base.@Macro.MacroAnalyzeInfo.nextArg
func (self *Macro_MacroAnalyzeInfo) NextArg(_env *LnsEnv) {
    self.argIndex = self.argIndex + 1
    
}

// 222: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingSymArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingSymArg(_env *LnsEnv) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        _env.SetStackVal( self.FP.GetCurArgType(_env) == Ast_builtinTypeSymbol) ).(bool)
}

// 235: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingExpArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingExpArg(_env *LnsEnv) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        _env.SetStackVal( self.FP.GetCurArgType(_env) == Ast_builtinTypeExp) ).(bool)
}


// declaration Class -- MacroCtrl
type Macro_MacroCtrlMtd interface {
    EvalMacroOp(_env *LnsEnv, arg1 string, arg2 *Types_Token, arg3 *Ast_TypeInfo, arg4 LnsAny)(LnsAny, LnsAny)
    ExpandMacroVal(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 *Ast_Scope, arg3 Parser_PushbackParser, arg4 *Types_Token) *Types_Token
    ExpandSymbol(_env *LnsEnv, arg1 Parser_PushbackParser, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_NodeManager, arg5 *LnsList) *Nodes_LiteralStringNode
    FinishMacroMode(_env *LnsEnv)
    Get_analyzeInfo(_env *LnsEnv) *Macro_MacroAnalyzeInfo
    Get_declMacroInfoMap(_env *LnsEnv) *LnsMap
    Get_isDeclaringMacro(_env *LnsEnv) bool
    Get_macroCallLineNo(_env *LnsEnv) LnsInt
    Get_tokenExpanding(_env *LnsEnv) bool
    Get_typeId2MacroInfo(_env *LnsEnv) *LnsMap
    Get_useModuleMacroSet(_env *LnsEnv) *LnsSet
    ImportMacro(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 *LnsMap, arg6 *LnsMap, arg7 LnsAny)
    ImportMacroInfo(_env *LnsEnv, arg1 *LnsMap)
    IsInAnalyzeArgMode(_env *LnsEnv) bool
    Regist(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode, arg3 *Ast_Scope, arg4 LnsAny)
    RegistVar(_env *LnsEnv, arg1 *LnsList)
    RestoreMacroMode(_env *LnsEnv)
    StartAnalyzeArgMode(_env *LnsEnv, arg1 *Ast_TypeInfo)
    StartDecl(_env *LnsEnv)
    StartExpandMode(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 Macro_EvalMacroCallback)
    SwitchMacroMode(_env *LnsEnv)
}
type Macro_MacroCtrl struct {
    useModuleMacroSet *LnsSet
    macroEval *Nodes_MacroEval
    typeId2MacroInfo *LnsMap
    declMacroInfoMap *LnsMap
    symbol2ValueMapForMacro *LnsMap
    analyzeInfo *Macro_MacroAnalyzeInfo
    macroAnalyzeInfoStack *LnsList
    tokenExpanding bool
    macroCallLineNo LnsInt
    macroLocalVarMap LnsAny
    isDeclaringMacro bool
    FP Macro_MacroCtrlMtd
}
func Macro_MacroCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroCtrl).FP
}
type Macro_MacroCtrlDownCast interface {
    ToMacro_MacroCtrl() *Macro_MacroCtrl
}
func Macro_MacroCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroCtrlDownCast)
    if ok { return work.ToMacro_MacroCtrl() }
    return nil
}
func (obj *Macro_MacroCtrl) ToMacro_MacroCtrl() *Macro_MacroCtrl {
    return obj
}
func NewMacro_MacroCtrl(_env *LnsEnv, arg1 *Nodes_MacroEval) *Macro_MacroCtrl {
    obj := &Macro_MacroCtrl{}
    obj.FP = obj
    obj.InitMacro_MacroCtrl(_env, arg1)
    return obj
}
func (self *Macro_MacroCtrl) Get_useModuleMacroSet(_env *LnsEnv) *LnsSet{ return self.useModuleMacroSet }
func (self *Macro_MacroCtrl) Get_typeId2MacroInfo(_env *LnsEnv) *LnsMap{ return self.typeId2MacroInfo }
func (self *Macro_MacroCtrl) Get_declMacroInfoMap(_env *LnsEnv) *LnsMap{ return self.declMacroInfoMap }
func (self *Macro_MacroCtrl) Get_analyzeInfo(_env *LnsEnv) *Macro_MacroAnalyzeInfo{ return self.analyzeInfo }
func (self *Macro_MacroCtrl) Get_tokenExpanding(_env *LnsEnv) bool{ return self.tokenExpanding }
func (self *Macro_MacroCtrl) Get_macroCallLineNo(_env *LnsEnv) LnsInt{ return self.macroCallLineNo }
func (self *Macro_MacroCtrl) Get_isDeclaringMacro(_env *LnsEnv) bool{ return self.isDeclaringMacro }
// 269: DeclConstr
func (self *Macro_MacroCtrl) InitMacro_MacroCtrl(_env *LnsEnv, macroEval *Nodes_MacroEval) {
    self.declMacroInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = false
    
    self.tokenExpanding = false
    
    self.useModuleMacroSet = NewLnsSet([]LnsAny{})
    
    self.typeId2MacroInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.macroEval = macroEval
    
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, Ast_builtinTypeNone, Nodes_MacroMode__None)
    
    self.macroCallLineNo = 0
    
    self.macroAnalyzeInfoStack = NewLnsList([]LnsAny{Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo)})
    
    Lns_LockEnvSync( _env, func () {
        self.macroLocalVarMap = _env.CommonLuaVM.RunLoadedfunc(Macro_toListEmpty,[]LnsAny{})[0].(*Lns_luaValue)
        
    })
}

// 302: decl @lune.@base.@Macro.MacroCtrl.evalMacroOp
func (self *Macro_MacroCtrl) EvalMacroOp(_env *LnsEnv, streamName string,firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny)(LnsAny, LnsAny) {
    self.useModuleMacroSet.Add(Ast_TypeInfo2Stem(macroTypeInfo.FP.GetModule(_env)))
    if expList != nil{
        expList_219 := expList.(*Nodes_ExpListNode)
        for _, _exp := range( expList_219.FP.Get_expList(_env).Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            var kind LnsInt
            kind = exp.FP.Get_kind(_env)
            if _switch1197 := kind; _switch1197 == Nodes_NodeKind_get_LiteralNil(_env) || _switch1197 == Nodes_NodeKind_get_LiteralChar(_env) || _switch1197 == Nodes_NodeKind_get_LiteralInt(_env) || _switch1197 == Nodes_NodeKind_get_LiteralReal(_env) || _switch1197 == Nodes_NodeKind_get_LiteralArray(_env) || _switch1197 == Nodes_NodeKind_get_LiteralList(_env) || _switch1197 == Nodes_NodeKind_get_LiteralMap(_env) || _switch1197 == Nodes_NodeKind_get_LiteralString(_env) || _switch1197 == Nodes_NodeKind_get_LiteralBool(_env) || _switch1197 == Nodes_NodeKind_get_LiteralSymbol(_env) || _switch1197 == Nodes_NodeKind_get_RefField(_env) || _switch1197 == Nodes_NodeKind_get_ExpMacroStat(_env) || _switch1197 == Nodes_NodeKind_get_ExpMacroArgExp(_env) || _switch1197 == Nodes_NodeKind_get_ExpOmitEnum(_env) || _switch1197 == Nodes_NodeKind_get_ExpCast(_env) || _switch1197 == Nodes_NodeKind_get_ExpOp2(_env) {
            } else {
                var mess string
                mess = _env.LuaVM.String_format("Macro arguments must be literal value. -- %d:%d:%s", []LnsAny{exp.FP.Get_pos(_env).LineNo, exp.FP.Get_pos(_env).Column, Nodes_getNodeKindName(_env, kind)})
                return nil, mess
            }
        }
    }
    var macroInfo *Nodes_MacroInfo
    macroInfo = Lns_unwrap( self.typeId2MacroInfo.Get(macroTypeInfo.FP.Get_typeId(_env))).(*Nodes_MacroInfo)
    var process func(_env *LnsEnv) LnsAny
    process = func(_env *LnsEnv) LnsAny {
        var argValMap *LnsMap
        argValMap = NewLnsMap( map[LnsAny]LnsAny{})
        var macroArgValMap *LnsMap
        macroArgValMap = NewLnsMap( map[LnsAny]LnsAny{})
        var macroArgNodeList *LnsList
        macroArgNodeList = macroInfo.FP.GetArgList(_env)
        var macroArgName2ArgNode *LnsMap
        macroArgName2ArgNode = NewLnsMap( map[LnsAny]LnsAny{})
        if expList != nil{
            expList_235 := expList.(*Nodes_ExpListNode)
            for _index, _argNode := range( expList_235.FP.Get_expList(_env).Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                var declArgNode *Nodes_MacroArgInfo
                declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                macroArgName2ArgNode.Set(declArgNode.FP.Get_name(_env),argNode)
                var literal LnsAny
                var mess LnsAny
                literal,mess = argNode.FP.GetLiteral(_env)
                if literal != nil{
                    literal_243 := literal
                    {
                        _val := Macro_getLiteralMacroVal_1150_(_env, literal_243)
                        if !Lns_IsNil( _val ) {
                            val := _val
                            argValMap.Set(index,val)
                            if argNode.FP.Get_expType(_env) == Ast_builtinTypeSymbol{
                                macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.CommonLuaVM.RunLoadedfunc(Macro_toLuaval,Lns_2DDD(Lns_FromStemGetAt(val,1, false )))[0])
                            } else { 
                                macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.CommonLuaVM.RunLoadedfunc(Macro_toLuaval,Lns_2DDD(val))[0])
                            }
                        }
                    }
                } else {
                    var errmess string
                    errmess = _env.LuaVM.String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env)), mess})
                    return errmess
                }
            }
        }
        macroArgValMap.Set("__var",self.macroLocalVarMap)
        var _func *Lns_luaValue
        _func = macroInfo.G_func
        var macroVars LnsAny
        macroVars = Lns_unwrap( _env.CommonLuaVM.ExpandLuavalMap(_env.CommonLuaVM.RunLoadedfunc(_func,Lns_2DDD(macroArgValMap))[0].(*Lns_luaValue)))
        self.macroLocalVarMap = Lns_unwrap( Lns_FromStemGetAt(macroVars,"__var", false ))
        
        for _, _name := range( (Lns_unwrap( Lns_FromStemGetAt(macroVars,"__names", false ))).(*LnsMap).Items ) {
            name := _name.(string)
            var valInfo *Nodes_MacroValInfo
            
            {
                _valInfo := macroInfo.Symbol2MacroValInfoMap.Get(name)
                if _valInfo == nil{
                    Util_err(_env, _env.LuaVM.String_format("not found macro symbol -- %s", []LnsAny{name}))
                } else {
                    valInfo = _valInfo.(*Nodes_MacroValInfo)
                }
            }
            var typeInfo *Ast_TypeInfo
            typeInfo = valInfo.TypeInfo
            var valMap LnsAny
            {
                _val := Lns_FromStemGetAt(macroVars,name, false )
                if !Lns_IsNil( _val ) {
                    val := _val
                    if Macro_equalsType_1404_(_env, typeInfo, Ast_builtinTypeSymbol){
                        valMap = NewLnsMap( map[LnsAny]LnsAny{1:val,})
                        
                    } else { 
                        valMap = val
                        
                    }
                } else {
                    valMap = NewLnsMap( map[LnsAny]LnsAny{})
                    
                }
            }
            self.symbol2ValueMapForMacro.Set(name,NewNodes_MacroValInfo(_env, valMap, typeInfo, nil))
        }
        for _index, _arg := range( macroInfo.FP.GetArgList(_env).Items ) {
            index := _index + 1
            arg := _arg.(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
            if arg.FP.Get_typeInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
                var argType *Ast_TypeInfo
                argType = arg.FP.Get_typeInfo(_env)
                var argName string
                argName = arg.FP.Get_name(_env)
                self.symbol2ValueMapForMacro.Set(argName,NewNodes_MacroValInfo(_env, argValMap.Get(index), argType, macroArgName2ArgNode.Get(argName)))
            } else { 
                return "not support ... in macro"
            }
        }
        return nil
    }
    Lns_LockEnvSync( _env, func () {
        process(_env)
    })
    return &NewMacro_MacroParser(_env, macroInfo.FP.GetTokenList(_env), _env.LuaVM.String_format("%s:%d:%d: (macro %s)", []LnsAny{streamName, firstToken.Pos.LineNo, firstToken.Pos.Column, macroTypeInfo.FP.GetTxt(_env, nil, nil, nil)}), firstToken.Pos.FP.Get_orgPos(_env)).Parser_Parser, nil
}

// 439: decl @lune.@base.@Macro.MacroCtrl.importMacro
func (self *Macro_MacroCtrl) ImportMacro(_env *LnsEnv, processInfo *Ast_ProcessInfo,lnsPath string,macroInfoStem LnsAny,macroTypeInfo *Ast_TypeInfo,typeId2TypeInfo *LnsMap,importedMacroInfoMap *LnsMap,baseDir LnsAny) {
    var macroInfo LnsAny
    var err LnsAny
    macroInfo,err = Macro_MacroMetaInfo__fromStem_1097_(_env, macroInfoStem,nil)
    if macroInfo != nil{
        macroInfo_286 := macroInfo.(*Macro_MacroMetaInfo)
        var orgPos *Types_Position
        if macroInfo_286.Pos.Len() == 2{
            orgPos = NewTypes_Position(_env, macroInfo_286.Pos.GetAt(1).(LnsInt), macroInfo_286.Pos.GetAt(2).(LnsInt), lnsPath)
            
        } else { 
            Util_err(_env, "macroInfo.pos is illegal")
        }
        var argList *LnsList
        argList = NewLnsList([]LnsAny{})
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        var symbol2MacroValInfoMap *LnsMap
        symbol2MacroValInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _argInfo := range( macroInfo_286.ArgList.Items ) {
            argInfo := _argInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var argTypeInfo *Ast_TypeInfo
            argTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(argInfo.TypeId)).(*Ast_TypeInfo)
            argList.Insert(Nodes_MacroArgInfo2Stem(NewNodes_MacroArgInfo(_env, argInfo.Name, argTypeInfo)))
            argNameList.Insert(argInfo.Name)
        }
        for _, _symInfo := range( macroInfo_286.SymList.Items ) {
            symInfo := _symInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var symTypeInfo *Ast_TypeInfo
            symTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(symInfo.TypeId)).(*Ast_TypeInfo)
            symbol2MacroValInfoMap.Set(symInfo.Name,NewNodes_MacroValInfo(_env, nil, symTypeInfo, nil))
        }
        var tokenList *LnsList
        tokenList = NewLnsList([]LnsAny{})
        var lineNo LnsInt
        lineNo = 0
        var column LnsInt
        column = 1
        for _, _tokenInfo := range( macroInfo_286.TokenList.Items ) {
            tokenInfo := _tokenInfo.(*LnsList)
            var txt string
            txt = tokenInfo.GetAt(2).(string)
            if txt == "\n"{
                lineNo = lineNo + 1
                
                column = 1
                
            } else { 
                var pos *Types_Position
                pos = Types_Position_create(_env, lineNo, column, _env.LuaVM.String_format("macro:%s", []LnsAny{macroInfo_286.Name}), orgPos)
                tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, Lns_unwrap( Types_TokenKind__from(_env, Lns_forceCastInt(tokenInfo.GetAt(1)))).(LnsInt), txt, pos, false, nil)))
                column = column + len(txt) + 1
                
            }
        }
        var extMacroInfo *Macro_ExtMacroInfo
        extMacroInfo = NewMacro_ExtMacroInfo(_env, macroInfo_286.Name, self.macroEval.FP.EvalFromCode(_env, processInfo, macroInfo_286.Name, argNameList, macroInfo_286.StmtBlock, baseDir), symbol2MacroValInfoMap, argList, tokenList)
        self.typeId2MacroInfo.Set(macroTypeInfo.FP.Get_typeId(_env),&extMacroInfo.Nodes_MacroInfo)
        importedMacroInfoMap.Set(macroTypeInfo.FP.Get_typeId(_env),&extMacroInfo.Nodes_MacroInfo)
    } else {
        Util_errorLog(_env, _env.LuaVM.String_format("macro load fail -- %s: %s ", []LnsAny{macroTypeInfo.FP.GetTxt(_env, nil, nil, nil), Lns_unwrapDefault( err, "").(string)}))
    }
}

// 506: decl @lune.@base.@Macro.MacroCtrl.importMacroInfo
func (self *Macro_MacroCtrl) ImportMacroInfo(_env *LnsEnv, importedMacroInfoMap *LnsMap) {
    for _typeId, _macroInfo := range( importedMacroInfoMap.Items ) {
        typeId := _typeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
        macroInfo := _macroInfo.(Nodes_MacroInfoDownCast).ToNodes_MacroInfo()
        self.typeId2MacroInfo.Set(typeId,macroInfo)
    }
}

// 515: decl @lune.@base.@Macro.MacroCtrl.regist
func (self *Macro_MacroCtrl) Regist(_env *LnsEnv, processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode,macroScope *Ast_Scope,baseDir LnsAny) {
    var macroObj *Lns_luaValue
    macroObj = self.macroEval.FP.Eval(_env, processInfo, node, baseDir)
    var remap *LnsMap
    remap = NewLnsMap( map[LnsAny]LnsAny{})
    for _name, _macroValInfo := range( self.symbol2ValueMapForMacro.Items ) {
        name := _name.(string)
        macroValInfo := _macroValInfo.(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
        if Macro_equalsType_1404_(_env, macroValInfo.TypeInfo, Ast_builtinTypeEmpty){
            {
                _typeInfo := macroScope.FP.GetTypeInfoChild(_env, name)
                if !Lns_IsNil( _typeInfo ) {
                    typeInfo := _typeInfo.(*Ast_TypeInfo)
                    remap.Set(name,NewNodes_MacroValInfo(_env, macroValInfo.Val, typeInfo, macroValInfo.ArgNode))
                } else {
                    remap.Set(name,macroValInfo)
                }
            }
        } else { 
            remap.Set(name,macroValInfo)
        }
    }
    var macroInfo *Nodes_DefMacroInfo
    macroInfo = NewNodes_DefMacroInfo(_env, macroObj, node.FP.Get_declInfo(_env), remap)
    self.typeId2MacroInfo.Set(node.FP.Get_expType(_env).FP.Get_typeId(_env),&macroInfo.Nodes_MacroInfo)
    self.declMacroInfoMap.Set(node.FP.Get_expType(_env).FP.Get_typeId(_env),&macroInfo.Nodes_MacroInfo)
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = false
    
}

// 620: decl @lune.@base.@Macro.MacroCtrl.expandMacroVal
func (self *Macro_MacroCtrl) ExpandMacroVal(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,scope *Ast_Scope,parser Parser_PushbackParser,token *Types_Token) *Types_Token {
    if self.tokenExpanding{
        return token
    }
    var getToken func(_env *LnsEnv) *Types_Token
    getToken = func(_env *LnsEnv) *Types_Token {
        self.tokenExpanding = true
        
        var work *Types_Token
        work = parser.GetTokenNoErr(_env)
        self.tokenExpanding = false
        
        return work
    }
    var tokenTxt string
    tokenTxt = token.Txt
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( tokenTxt == ",,") ||
        _env.SetStackVal( tokenTxt == ",,,") ||
        _env.SetStackVal( tokenTxt == ",,,,") ).(bool){
        var nextToken *Types_Token
        nextToken = getToken(_env)
        var macroVal *Nodes_MacroValInfo
        
        {
            _macroVal := self.symbol2ValueMapForMacro.Get(nextToken.Txt)
            if _macroVal == nil{
                parser.Error(_env, _env.LuaVM.String_format("unknown macro val -- %s", []LnsAny{nextToken.Txt}))
            } else {
                macroVal = _macroVal.(*Nodes_MacroValInfo)
            }
        }
        if tokenTxt == ",,"{
            if Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *LnsList
                txtList = NewLnsList([]LnsAny{})
                for _, _txt := range( MacroCtrl_expandMacroVal__macroVal2strList_1716_(_env, nextToken.Txt, macroVal, parser).Items ) {
                    txt := _txt.(string)
                    txtList.Insert(txt)
                }
                Macro_pushbackTxt_1676_(_env, parser, txtList, nextToken.Txt, nextToken.Pos)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ).(bool){
                var pos *Types_Position
                pos = _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(macroVal.ArgNode) && 
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.Get_pos(_env)})&&
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Types_Position).FP.Get_RawOrgPos(_env)}))) ||
                    _env.SetStackVal( nextToken.Pos.FP.Get_RawOrgPos(_env)) ||
                    _env.SetStackVal( token.Pos.FP.Get_orgPos(_env)) ).(*Types_Position)
                parser.PushbackStr(_env, _env.LuaVM.String_format("macroVal %s", []LnsAny{nextToken.Txt}), (Lns_unwrap( macroVal.Val)).(string), pos)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Array) ||
                _env.SetStackVal( macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool){
                if Macro_equalsType_1404_(_env, macroVal.TypeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat){
                    var pos *Types_Position
                    pos = _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(macroVal.ArgNode) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.Get_pos(_env)})&&
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Types_Position).FP.Get_RawOrgPos(_env)}))) ||
                        _env.SetStackVal( nextToken.Pos.FP.Get_RawOrgPos(_env)) ||
                        _env.SetStackVal( token.Pos.FP.Get_orgPos(_env)) ).(*Types_Position)
                    var strList *LnsList
                    strList = MacroCtrl_expandMacroVal__macroVal2strList_1716_(_env, nextToken.Txt, macroVal, parser)
                    {
                        var _from3134 LnsInt = strList.Len()
                        var _to3134 LnsInt = 1
                        _work3134 := _from3134
                        _delta3134 := -1
                        for {
                            if _delta3134 > 0 {
                               if _work3134 > _to3134 { break }
                            } else {
                               if _work3134 < _to3134 { break }
                            }
                            index := _work3134
                            parser.PushbackStr(_env, _env.LuaVM.String_format("macroVal %s[%d]", []LnsAny{nextToken.Txt, index}), strList.GetAt(index).(string), pos)
                            _work3134 += _delta3134
                        }
                    }
                } else { 
                    var tokenList *LnsList
                    tokenList = NewLnsList([]LnsAny{})
                    {
                        _argNode := macroVal.ArgNode
                        if !Lns_IsNil( _argNode ) {
                            argNode := _argNode.(*Nodes_Node)
                            if Lns_op_not(argNode.FP.SetupLiteralTokenList(_env, tokenList)){
                                parser.Error(_env, _env.LuaVM.String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                            }
                        } else {
                            parser.Error(_env, _env.LuaVM.String_format("not support ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    }
                    parser.NewPushback(_env, tokenList)
                }
            } else if macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum{
                var enumTypeInfo *Ast_EnumTypeInfo
                enumTypeInfo = Lns_unwrap( Ast_EnumTypeInfoDownCastF(macroVal.TypeInfo.FP.Get_aliasSrc(_env).FP)).(*Ast_EnumTypeInfo)
                var fullname string
                fullname = macroVal.TypeInfo.FP.GetFullName(_env, typeNameCtrl, scope.FP, true)
                var nameList *LnsList
                nameList = Util_splitStr(_env, fullname, "[^%.]+")
                var enumValInfo *Ast_EnumValInfo
                enumValInfo = Lns_unwrap( enumTypeInfo.FP.Get_val2EnumValInfo(_env).Get(Lns_unwrap( macroVal.Val))).(*Ast_EnumValInfo)
                nextToken = NewTypes_Token(_env, Types_TokenKind__Symb, enumValInfo.FP.Get_name(_env), nextToken.Pos, false, nil)
                
                parser.PushbackToken(_env, nextToken)
                {
                    var _from3375 LnsInt = nameList.Len()
                    var _to3375 LnsInt = 1
                    _work3375 := _from3375
                    _delta3375 := -1
                    for {
                        if _delta3375 > 0 {
                           if _work3375 > _to3375 { break }
                        } else {
                           if _work3375 < _to3375 { break }
                        }
                        index := _work3375
                        nextToken = NewTypes_Token(_env, Types_TokenKind__Dlmt, ".", nextToken.Pos, false, nil)
                        
                        parser.PushbackToken(_env, nextToken)
                        nextToken = NewTypes_Token(_env, Types_TokenKind__Symb, nameList.GetAt(index).(string), nextToken.Pos, false, nil)
                        
                        parser.PushbackToken(_env, nextToken)
                        _work3375 += _delta3375
                    }
                }
            } else { 
                var tokenList *LnsList
                tokenList = NewLnsList([]LnsAny{})
                {
                    _argNode := macroVal.ArgNode
                    if !Lns_IsNil( _argNode ) {
                        argNode := _argNode.(*Nodes_Node)
                        if Lns_op_not(argNode.FP.SetupLiteralTokenList(_env, tokenList)){
                            parser.Error(_env, _env.LuaVM.String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    } else {
                        Macro_expandVal_1649_(_env, tokenList, macroVal.Val, nextToken.Pos)
                    }
                }
                parser.NewPushback(_env, tokenList)
            }
        } else if tokenTxt == ",,,"{
            if Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeString){
                Macro_pushbackTxt_1676_(_env, parser, NewLnsList([]LnsAny{(Lns_unwrap( macroVal.Val)).(string)}), nextToken.Txt, nextToken.Pos)
            } else { 
                parser.Error(_env, _env.LuaVM.String_format("',,,' does not support this type -- %s", []LnsAny{macroVal.TypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else if tokenTxt == ",,,,"{
            if Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *LnsList
                txtList = (Lns_unwrap( macroVal.Val)).(*LnsList)
                var newToken string
                newToken = ""
                for _, _txt := range( txtList.Items ) {
                    txt := _txt.(string)
                    newToken = _env.LuaVM.String_format("%s%s", []LnsAny{newToken, txt})
                    
                }
                nextToken = NewTypes_Token(_env, Types_TokenKind__Str, _env.LuaVM.String_format("'%s'", []LnsAny{newToken}), nextToken.Pos, false, nil)
                
                parser.PushbackToken(_env, nextToken)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                _env.SetStackVal( Macro_equalsType_1404_(_env, macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ).(bool){
                var txt string
                txt = (Lns_unwrap( macroVal.Val)).(string)
                var rawTxt string
                if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(txt,"^```", nil, nil))){
                    rawTxt = Parser_quoteStr(_env, txt)
                    
                } else { 
                    rawTxt = Parser_quoteStr(_env, txt)
                    
                }
                nextToken = NewTypes_Token(_env, Types_TokenKind__Str, rawTxt, nextToken.Pos, false, nil)
                
                parser.PushbackToken(_env, nextToken)
            } else { 
                parser.Error(_env, _env.LuaVM.String_format("not support this symbol -- %s%s", []LnsAny{tokenTxt, nextToken.Txt}))
            }
        }
        nextToken = getToken(_env)
        
        token = nextToken
        
    }
    self.tokenExpanding = false
    
    return token
}

// 804: decl @lune.@base.@Macro.MacroCtrl.expandSymbol
func (self *Macro_MacroCtrl) ExpandSymbol(_env *LnsEnv, parser Parser_PushbackParser,prefixToken *Types_Token,exp *Nodes_Node,nodeManager *Nodes_NodeManager,errMessList *LnsList) *Nodes_LiteralStringNode {
    var nextToken *Types_Token
    nextToken = parser.GetTokenNoErr(_env)
    if nextToken.Txt != "~~"{
        parser.PushbackToken(_env, nextToken)
    }
    var format string
    format = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( prefixToken.Txt == ",,,") &&
        _env.SetStackVal( "' %s '") ||
        _env.SetStackVal( "\"'%s'\"") ).(string)
    if prefixToken.Txt == ",,"{
        {
            _refNode := Nodes_ExpRefNodeDownCastF(exp.FP)
            if !Lns_IsNil( _refNode ) {
                refNode := _refNode.(*Nodes_ExpRefNode)
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = refNode.FP.Get_symbolInfo(_env)
                var macroInfo LnsAny
                macroInfo = self.symbol2ValueMapForMacro.Get(symbolInfo.FP.Get_name(_env))
                if macroInfo != nil{
                    macroInfo_473 := macroInfo.(*Nodes_MacroValInfo)
                    var valType *Ast_TypeInfo
                    valType = macroInfo_473.TypeInfo
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeSymbol)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeExp)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeMultiExp)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                        
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( valType.FP.Get_kind(_env) == Ast_TypeInfoKind__List) &&
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat)) ).(bool)){
                        format = "' %s '"
                        
                        exp = &Nodes_ExpMacroStatListNode_create(_env, nodeManager, prefixToken.Pos, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), exp).Nodes_Node
                        
                    } else if Macro_equalsType_1404_(_env, Ast_builtinTypeString, valType){
                    } else if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeInt)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, valType, Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                        
                    } else { 
                        errMessList.Insert(Macro_ErrorMess2Stem(NewMacro_ErrorMess(_env, Lns_unwrap( symbolInfo.FP.Get_pos(_env)).(*Types_Position), _env.LuaVM.String_format("not support ,, -- %s", []LnsAny{valType.FP.GetTxt(_env, nil, nil, nil)}))))
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeInt)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                        
                    } else if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeSymbol)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeExp)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeMultiExp)) ||
                        _env.SetStackVal( Macro_equalsType_1404_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                        
                    }
                }
            }
        }
    }
    var newToken *Types_Token
    newToken = NewTypes_Token(_env, Types_TokenKind__Str, format, prefixToken.Pos, prefixToken.Consecutive, nil)
    var expListNode *Nodes_ExpListNode
    expListNode = Nodes_ExpListNode_create(_env, nodeManager, exp.FP.Get_pos(_env), self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), NewLnsList([]LnsAny{Nodes_Node2Stem(exp)}), nil, false)
    var dddNode *Nodes_ExpToDDDNode
    dddNode = Nodes_ExpToDDDNode_create(_env, nodeManager, exp.FP.Get_pos(_env), self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), expListNode)
    var literalStr *Nodes_LiteralStringNode
    literalStr = Nodes_LiteralStringNode_create(_env, nodeManager, prefixToken.Pos, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), newToken, expListNode, Nodes_ExpListNode_create(_env, nodeManager, exp.FP.Get_pos(_env), self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), NewLnsList([]LnsAny{Nodes_ExpToDDDNode2Stem(dddNode)}), nil, false))
    return literalStr
}

// 897: decl @lune.@base.@Macro.MacroCtrl.registVar
func (self *Macro_MacroCtrl) RegistVar(_env *LnsEnv, symbolList *LnsList) {
    for _, _symbolInfo := range( symbolList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var info *Nodes_MacroValInfo
        info = NewNodes_MacroValInfo(_env, nil, symbolInfo.FP.Get_typeInfo(_env), nil)
        self.symbol2ValueMapForMacro.Set(symbolInfo.FP.Get_name(_env),info)
    }
}

// 905: decl @lune.@base.@Macro.MacroCtrl.startDecl
func (self *Macro_MacroCtrl) StartDecl(_env *LnsEnv) {
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = true
    
}

// 912: decl @lune.@base.@Macro.MacroCtrl.finishMacroMode
func (self *Macro_MacroCtrl) FinishMacroMode(_env *LnsEnv) {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
}

// 918: decl @lune.@base.@Macro.MacroCtrl.startExpandMode
func (self *Macro_MacroCtrl) StartExpandMode(_env *LnsEnv, lineNo LnsInt,typeInfo *Ast_TypeInfo,callback Macro_EvalMacroCallback) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, typeInfo, Nodes_MacroMode__Expand)
    
    self.macroCallLineNo = lineNo
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
    callback(_env)
    self.FP.FinishMacroMode(_env)
}

// 930: decl @lune.@base.@Macro.MacroCtrl.startAnalyzeArgMode
func (self *Macro_MacroCtrl) StartAnalyzeArgMode(_env *LnsEnv, macroFuncType *Ast_TypeInfo) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, macroFuncType, Nodes_MacroMode__AnalyzeArg)
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}

// 935: decl @lune.@base.@Macro.MacroCtrl.switchMacroMode
func (self *Macro_MacroCtrl) SwitchMacroMode(_env *LnsEnv) {
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len() - 1).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}

// 940: decl @lune.@base.@Macro.MacroCtrl.restoreMacroMode
func (self *Macro_MacroCtrl) RestoreMacroMode(_env *LnsEnv) {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
}

// 949: decl @lune.@base.@Macro.MacroCtrl.isInAnalyzeArgMode
func (self *Macro_MacroCtrl) IsInAnalyzeArgMode(_env *LnsEnv) bool {
    if self.macroAnalyzeInfoStack.Len() == 0{
        return false
    }
    for _, _info := range( self.macroAnalyzeInfoStack.Items ) {
        info := _info.(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
        if info.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg{
            return true
        }
    }
    return false
}


// declaration Class -- ErrorMess
type Macro_ErrorMessMtd interface {
}
type Macro_ErrorMess struct {
    Pos *Types_Position
    Mess string
    FP Macro_ErrorMessMtd
}
func Macro_ErrorMess2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_ErrorMess).FP
}
type Macro_ErrorMessDownCast interface {
    ToMacro_ErrorMess() *Macro_ErrorMess
}
func Macro_ErrorMessDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_ErrorMessDownCast)
    if ok { return work.ToMacro_ErrorMess() }
    return nil
}
func (obj *Macro_ErrorMess) ToMacro_ErrorMess() *Macro_ErrorMess {
    return obj
}
func NewMacro_ErrorMess(_env *LnsEnv, arg1 *Types_Position, arg2 string) *Macro_ErrorMess {
    obj := &Macro_ErrorMess{}
    obj.FP = obj
    obj.InitMacro_ErrorMess(_env, arg1, arg2)
    return obj
}
func (self *Macro_ErrorMess) InitMacro_ErrorMess(_env *LnsEnv, arg1 *Types_Position, arg2 string) {
    self.Pos = arg1
    self.Mess = arg2
}

func Lns_Macro_init(_env *LnsEnv) {
    if init_Macro { return }
    init_Macro = true
    Macro__mod__ = "@lune.@base.@Macro"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Nodes_init(_env)
    Lns_Ast_init(_env)
    Lns_Parser_init(_env)
    Lns_Types_init(_env)
    Lns_Formatter_init(_env)
    Macro_toListEmpty = Macro_loadCode_1008_(_env, "return function() return {} end").(*Lns_luaValue)
    Macro_toLuaval = Macro_loadCode_1008_(_env, "return function( val ) return val end").(*Lns_luaValue)
}
func init() {
    init_Macro = false
}
