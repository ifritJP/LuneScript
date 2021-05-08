// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Macro bool
var Macro__mod__ string
var Macro_toList *Lns_luaValue
var Macro_toListEmpty *Lns_luaValue
var Macro_toLuaval *Lns_luaValue
type Macro_toListLua func (arg1 LnsAny) *LnsList
type Macro_toListEmptyLua func () *LnsList
type Macro_toLuavalLua func (arg1 LnsAny) LnsAny
type Macro_EvalMacroCallback func ()
// for 37
func Macro_convExp42(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 34: decl @lune.@base.@Macro.loadCode
func Macro_loadCode_1004_(code string) LnsAny {
    var loaded LnsAny
    var mess LnsAny
    loaded,mess = Lns_getVM().Load(code, nil)
    if loaded != nil{
        loaded_3238 := loaded.(*Lns_luaValue)
        {
            _obj := Macro_convExp42(Lns_2DDD(Lns_getVM().RunLoadedfunc(loaded_3238,Lns_2DDD([]LnsAny{}))[0]))
            if !Lns_IsNil( _obj ) {
                obj := _obj
                return obj
            }
        }
        panic("failed to load")
    }
    panic(Lns_getVM().String_format("%s -- %s", []LnsAny{mess, code}))
// insert a dummy
    return nil
}

// 112: decl @lune.@base.@Macro.getLiteralMacroVal
func Macro_getLiteralMacroVal_1090_(obj LnsAny) LnsAny {
    switch _exp597 := obj.(type) {
    case *Nodes_Literal__Nil:
        return nil
    case *Nodes_Literal__Int:
    val := _exp597.Val1
        return val
    case *Nodes_Literal__Real:
    val := _exp597.Val1
        return val
    case *Nodes_Literal__Str:
    val := _exp597.Val1
        return val
    case *Nodes_Literal__Bool:
    val := _exp597.Val1
        return val
    case *Nodes_Literal__Symbol:
    val := _exp597.Val1
        return NewLnsList([]LnsAny{val})
    case *Nodes_Literal__Field:
    val := _exp597.Val1
        return val
    case *Nodes_Literal__LIST:
    list := _exp597.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_1090_(item))
        }
        return newList
    case *Nodes_Literal__ARRAY:
    list := _exp597.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_1090_(item))
        }
        return newList
    case *Nodes_Literal__SET:
    list := _exp597.Val1
        var newSet *LnsSet
        newSet = NewLnsSet([]LnsAny{})
        for _, _item := range( list.Items ) {
            item := _item
            {
                __exp := Macro_getLiteralMacroVal_1090_(item)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp
                    newSet.Add(_exp)
                }
            }
        }
        return newSet
    case *Nodes_Literal__MAP:
    _map := _exp597.Val1
        var newMap *LnsMap
        newMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _key, _val := range( _map.Items ) {
            key := _key
            val := _val
            var keyObj LnsAny
            keyObj = Macro_getLiteralMacroVal_1090_(key)
            var valObj LnsAny
            valObj = Macro_getLiteralMacroVal_1090_(val)
            if keyObj != nil && valObj != nil{
                keyObj_3343 := keyObj
                valObj_3344 := valObj
                newMap.Set(keyObj_3343,valObj_3344)
            }
        }
        return newMap
    }
// insert a dummy
    return nil
}

// 283: decl @lune.@base.@Macro.equalsType
func Macro_equalsType_1219_(typeInfo *Ast_TypeInfo,builtinType *Ast_TypeInfo) bool {
    return typeInfo.FP.Get_srcTypeInfo() == builtinType
}

// 532: decl @lune.@base.@Macro.expandVal
func Macro_expandVal_1292_(tokenList *LnsList,workval LnsAny,pos *Types_Position) LnsAny {
    if workval != nil{
        workval_3560 := workval
        var val LnsAny
        val = workval_3560
        if _switch2523 := Lns_type(val); _switch2523 == "boolean" {
            var token string
            token = Lns_getVM().String_format("%s", []LnsAny{val})
            var kind LnsInt
            kind = Types_TokenKind__Kywd
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(kind, token, pos, false, nil)))
        } else if _switch2523 == "number" {
            var num string
            num = Lns_getVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
            var kind LnsInt
            kind = Types_TokenKind__Int
            if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(num, ".", 1, true))){
                kind = Types_TokenKind__Real
                
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(kind, num, pos, false, nil)))
        } else if _switch2523 == "string" {
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(Types_TokenKind__Str, Parser_quoteStr(val.(string)), pos, false, nil)))
        } else {
            return Lns_getVM().String_format("not support ,, List -- %s", []LnsAny{Lns_type(val)})
        }
    }
    return nil
}

// 578: decl @lune.@base.@Macro.pushbackTxt
func Macro_pushbackTxt_1297_(pushbackParser Parser_PushbackParser,txtList *LnsList,streamName string,pos *Types_Position) {
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    for _, _txt := range( txtList.Items ) {
        txt := _txt.(string)
        var stream *Parser_TxtStream
        stream = NewParser_TxtStream(txt)
        var parser *Parser_StreamParser
        parser = NewParser_StreamParser(stream.FP, Lns_getVM().String_format("macro symbol -- %s", []LnsAny{streamName}), false, pos.FP.Get_RawOrgPos())
        var workParser *Parser_DefaultPushbackParser
        workParser = NewParser_DefaultPushbackParser(&parser.Parser_Parser)
        for  {
            var worktoken *Types_Token
            worktoken = workParser.FP.GetTokenNoErr()
            if worktoken.Kind == Types_TokenKind__Eof{
                break
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(worktoken.Kind, worktoken.Txt, pos, false, nil)))
        }
    }
    {
        var _from2676 LnsInt = tokenList.Len()
        var _to2676 LnsInt = 1
        _work2676 := _from2676
        _delta2676 := -1
        for {
            if _delta2676 > 0 {
               if _work2676 > _to2676 { break }
            } else {
               if _work2676 < _to2676 { break }
            }
            index := _work2676
            pushbackParser.PushbackToken(tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work2676 += _delta2676
        }
    }
}


// 938: decl @lune.@base.@Macro.nodeToCodeTxt
func Macro_nodeToCodeTxt(node *Nodes_Node,moduleTypeInfo *Ast_TypeInfo) string {
    var memStream *Util_memStream
    memStream = NewUtil_memStream()
    var formatter *Nodes_Filter
    formatter = Formatter_createFilter(moduleTypeInfo, memStream.FP)
    node.FP.ProcessFilter(formatter, Formatter_Opt2Stem(NewFormatter_Opt(node)))
    return memStream.FP.Get_txt()
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
func NewMacro_MacroMetaArgInfo(arg1 string, arg2 LnsInt) *Macro_MacroMetaArgInfo {
    obj := &Macro_MacroMetaArgInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaArgInfo(arg1, arg2)
    return obj
}
func (self *Macro_MacroMetaArgInfo) InitMacro_MacroMetaArgInfo(arg1 string, arg2 LnsInt) {
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
func Macro_MacroMetaArgInfo__fromMap_1037_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaArgInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaArgInfo__fromStem_1039_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
func NewMacro_MacroMetaInfo(arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) *Macro_MacroMetaInfo {
    obj := &Macro_MacroMetaInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaInfo(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Macro_MacroMetaInfo) InitMacro_MacroMetaInfo(arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) {
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
func Macro_MacroMetaInfo__fromMap_1066_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaInfo__fromStem_1068_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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

// declaration Class -- MacroPaser
type Macro_MacroPaserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
}
type Macro_MacroPaser struct {
    Parser_Parser
    tokenList *LnsList
    pos LnsInt
    name string
    overridePos LnsAny
    FP Macro_MacroPaserMtd
}
func Macro_MacroPaser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroPaser).FP
}
type Macro_MacroPaserDownCast interface {
    ToMacro_MacroPaser() *Macro_MacroPaser
}
func Macro_MacroPaserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroPaserDownCast)
    if ok { return work.ToMacro_MacroPaser() }
    return nil
}
func (obj *Macro_MacroPaser) ToMacro_MacroPaser() *Macro_MacroPaser {
    return obj
}
func NewMacro_MacroPaser(arg1 *LnsList, arg2 string, arg3 LnsAny) *Macro_MacroPaser {
    obj := &Macro_MacroPaser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitMacro_MacroPaser(arg1, arg2, arg3)
    return obj
}
// 74: DeclConstr
func (self *Macro_MacroPaser) InitMacro_MacroPaser(tokenList *LnsList,name string,overridePos LnsAny) {
    self.InitParser_Parser()
    self.pos = 1
    
    self.tokenList = tokenList
    
    self.name = name
    
    self.overridePos = overridePos
    
}

// 84: decl @lune.@base.@Macro.MacroPaser.createPosition
func (self *Macro_MacroPaser) CreatePosition(lineNo LnsInt,column LnsInt) *Types_Position {
    return Types_Position_create(lineNo, column, self.FP.GetStreamName(), self.overridePos)
}

// 89: decl @lune.@base.@Macro.MacroPaser.getToken
func (self *Macro_MacroPaser) GetToken() LnsAny {
    if self.tokenList.Len() < self.pos{
        return nil
    }
    var token *Types_Token
    token = self.tokenList.GetAt(self.pos).(Types_TokenDownCast).ToTypes_Token()
    self.pos = self.pos + 1
    
    {
        __ := self.overridePos
        if !Lns_IsNil( __ ) {
            return NewTypes_Token(token.Kind, token.Txt, self.FP.CreatePosition(token.Pos.LineNo, token.Pos.Column), token.Consecutive, token.FP.Get_commentList())
        }
    }
    return token
}

// 106: decl @lune.@base.@Macro.MacroPaser.getStreamName
func (self *Macro_MacroPaser) GetStreamName() string {
    return self.name
}


// declaration Class -- ExtMacroInfo
type Macro_ExtMacroInfoMtd interface {
    GetArgList() *LnsList
    GetTokenList() *LnsList
    Get_name() string
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
func NewMacro_ExtMacroInfo(arg1 string, arg2 *Lns_luaValue, arg3 *LnsMap, arg4 *LnsList, arg5 *LnsList) *Macro_ExtMacroInfo {
    obj := &Macro_ExtMacroInfo{}
    obj.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitMacro_ExtMacroInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Macro_ExtMacroInfo) Get_name() string{ return self.name }
// 177: decl @lune.@base.@Macro.ExtMacroInfo.getArgList
func (self *Macro_ExtMacroInfo) GetArgList() *LnsList {
    return self.argList
}

// 180: decl @lune.@base.@Macro.ExtMacroInfo.getTokenList
func (self *Macro_ExtMacroInfo) GetTokenList() *LnsList {
    return self.tokenList
}

// 184: DeclConstr
func (self *Macro_ExtMacroInfo) InitMacro_ExtMacroInfo(name string,_func *Lns_luaValue,symbol2MacroValInfoMap *LnsMap,argList *LnsList,tokenList *LnsList) {
    self.InitNodes_MacroInfo(_func, symbol2MacroValInfoMap)
    self.name = name
    
    self.argList = argList
    
    self.tokenList = tokenList
    
}


// declaration Class -- MacroAnalyzeInfo
type Macro_MacroAnalyzeInfoMtd interface {
    EqualsArgTypeList(arg1 LnsAny) bool
    GetCurArgType() *Ast_TypeInfo
    Get_argIndex() LnsInt
    Get_mode() LnsInt
    Get_typeInfo() *Ast_TypeInfo
    IsAnalyzingExpArg() bool
    IsAnalyzingSymArg() bool
    NextArg()
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
func NewMacro_MacroAnalyzeInfo(arg1 *Ast_TypeInfo, arg2 LnsInt) *Macro_MacroAnalyzeInfo {
    obj := &Macro_MacroAnalyzeInfo{}
    obj.FP = obj
    obj.InitMacro_MacroAnalyzeInfo(arg1, arg2)
    return obj
}
func (self *Macro_MacroAnalyzeInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *Macro_MacroAnalyzeInfo) Get_mode() LnsInt{ return self.mode }
func (self *Macro_MacroAnalyzeInfo) Get_argIndex() LnsInt{ return self.argIndex }
// 201: DeclConstr
func (self *Macro_MacroAnalyzeInfo) InitMacro_MacroAnalyzeInfo(typeInfo *Ast_TypeInfo,mode LnsInt) {
    self.typeInfo = typeInfo
    
    self.mode = mode
    
    self.argIndex = 1
    
}

// 207: decl @lune.@base.@Macro.MacroAnalyzeInfo.equalsArgTypeList
func (self *Macro_MacroAnalyzeInfo) EqualsArgTypeList(argTypeList LnsAny) bool {
    return self.typeInfo.FP.Get_argTypeInfoList() == argTypeList
}

// 210: decl @lune.@base.@Macro.MacroAnalyzeInfo.getCurArgType
func (self *Macro_MacroAnalyzeInfo) GetCurArgType() *Ast_TypeInfo {
    if self.typeInfo.FP.Get_argTypeInfoList().Len() < self.argIndex{
        return Ast_builtinTypeNone
    }
    return self.typeInfo.FP.Get_argTypeInfoList().GetAt(self.argIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}

// 216: decl @lune.@base.@Macro.MacroAnalyzeInfo.nextArg
func (self *Macro_MacroAnalyzeInfo) NextArg() {
    self.argIndex = self.argIndex + 1
    
}

// 220: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingSymArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingSymArg() bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        Lns_GetEnv().SetStackVal( self.FP.GetCurArgType() == Ast_builtinTypeSymbol) ).(bool)
}

// 233: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingExpArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingExpArg() bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        Lns_GetEnv().SetStackVal( self.FP.GetCurArgType() == Ast_builtinTypeExp) ).(bool)
}


// declaration Class -- MacroCtrl
type Macro_MacroCtrlMtd interface {
    EvalMacroOp(arg1 string, arg2 *Types_Token, arg3 *Ast_TypeInfo, arg4 LnsAny)(LnsAny, LnsAny)
    ExpandMacroVal(arg1 *Ast_TypeNameCtrl, arg2 *Ast_Scope, arg3 Parser_PushbackParser, arg4 *Types_Token) *Types_Token
    ExpandSymbol(arg1 Parser_PushbackParser, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_NodeManager, arg5 *LnsList) *Nodes_LiteralStringNode
    FinishMacroMode()
    Get_analyzeInfo() *Macro_MacroAnalyzeInfo
    Get_declMacroInfoMap() *LnsMap
    Get_isDeclaringMacro() bool
    Get_macroCallLineNo() LnsInt
    Get_tokenExpanding() bool
    Get_typeId2MacroInfo() *LnsMap
    Get_useModuleMacroSet() *LnsSet
    ImportMacro(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 *LnsMap, arg6 *LnsMap)
    ImportMacroInfo(arg1 *LnsMap)
    IsInAnalyzeArgMode() bool
    Regist(arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode, arg3 *Ast_Scope)
    RegistVar(arg1 *LnsList)
    RestoreMacroMode()
    StartAnalyzeArgMode(arg1 *Ast_TypeInfo)
    StartDecl()
    StartExpandMode(arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 Macro_EvalMacroCallback)
    SwitchMacroMode()
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
func NewMacro_MacroCtrl(arg1 *Nodes_MacroEval) *Macro_MacroCtrl {
    obj := &Macro_MacroCtrl{}
    obj.FP = obj
    obj.InitMacro_MacroCtrl(arg1)
    return obj
}
func (self *Macro_MacroCtrl) Get_useModuleMacroSet() *LnsSet{ return self.useModuleMacroSet }
func (self *Macro_MacroCtrl) Get_typeId2MacroInfo() *LnsMap{ return self.typeId2MacroInfo }
func (self *Macro_MacroCtrl) Get_declMacroInfoMap() *LnsMap{ return self.declMacroInfoMap }
func (self *Macro_MacroCtrl) Get_analyzeInfo() *Macro_MacroAnalyzeInfo{ return self.analyzeInfo }
func (self *Macro_MacroCtrl) Get_tokenExpanding() bool{ return self.tokenExpanding }
func (self *Macro_MacroCtrl) Get_macroCallLineNo() LnsInt{ return self.macroCallLineNo }
func (self *Macro_MacroCtrl) Get_isDeclaringMacro() bool{ return self.isDeclaringMacro }
// 267: DeclConstr
func (self *Macro_MacroCtrl) InitMacro_MacroCtrl(macroEval *Nodes_MacroEval) {
    self.declMacroInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = false
    
    self.tokenExpanding = false
    
    self.useModuleMacroSet = NewLnsSet([]LnsAny{})
    
    self.typeId2MacroInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.macroEval = macroEval
    
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(Ast_builtinTypeNone, Nodes_MacroMode__None)
    
    self.macroCallLineNo = 0
    
    self.macroAnalyzeInfoStack = NewLnsList([]LnsAny{Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo)})
    
    self.macroLocalVarMap = Lns_getVM().RunLoadedfunc(Macro_toListEmpty,[]LnsAny{})[0].(*Lns_luaValue)
    
}

// 297: decl @lune.@base.@Macro.MacroCtrl.evalMacroOp
func (self *Macro_MacroCtrl) EvalMacroOp(streamName string,firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny)(LnsAny, LnsAny) {
    self.useModuleMacroSet.Add(Ast_TypeInfo2Stem(macroTypeInfo.FP.GetModule()))
    if expList != nil{
        expList_3441 := expList.(*Nodes_ExpListNode)
        for _, _exp := range( expList_3441.FP.Get_expList().Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            var kind LnsInt
            kind = exp.FP.Get_kind()
            if _switch1213 := kind; _switch1213 == Nodes_NodeKind_get_LiteralNil() || _switch1213 == Nodes_NodeKind_get_LiteralChar() || _switch1213 == Nodes_NodeKind_get_LiteralInt() || _switch1213 == Nodes_NodeKind_get_LiteralReal() || _switch1213 == Nodes_NodeKind_get_LiteralArray() || _switch1213 == Nodes_NodeKind_get_LiteralList() || _switch1213 == Nodes_NodeKind_get_LiteralMap() || _switch1213 == Nodes_NodeKind_get_LiteralString() || _switch1213 == Nodes_NodeKind_get_LiteralBool() || _switch1213 == Nodes_NodeKind_get_LiteralSymbol() || _switch1213 == Nodes_NodeKind_get_RefField() || _switch1213 == Nodes_NodeKind_get_ExpMacroStat() || _switch1213 == Nodes_NodeKind_get_ExpMacroArgExp() || _switch1213 == Nodes_NodeKind_get_ExpOmitEnum() || _switch1213 == Nodes_NodeKind_get_ExpCast() || _switch1213 == Nodes_NodeKind_get_ExpOp2() {
            } else {
                var mess string
                mess = Lns_getVM().String_format("Macro arguments must be literal value. -- %d:%d:%s", []LnsAny{exp.FP.Get_pos().LineNo, exp.FP.Get_pos().Column, Nodes_getNodeKindName(kind)})
                return nil, mess
            }
        }
    }
    var macroInfo *Nodes_MacroInfo
    macroInfo = Lns_unwrap( self.typeId2MacroInfo.Get(macroTypeInfo.FP.Get_typeId())).(*Nodes_MacroInfo)
    var argValMap *LnsMap
    argValMap = NewLnsMap( map[LnsAny]LnsAny{})
    var macroArgValMap *LnsMap
    macroArgValMap = NewLnsMap( map[LnsAny]LnsAny{"__var":self.macroLocalVarMap,})
    var macroArgNodeList *LnsList
    macroArgNodeList = macroInfo.FP.GetArgList()
    var macroArgName2ArgNode *LnsMap
    macroArgName2ArgNode = NewLnsMap( map[LnsAny]LnsAny{})
    if expList != nil{
        expList_3454 := expList.(*Nodes_ExpListNode)
        for _index, _argNode := range( expList_3454.FP.Get_expList().Items ) {
            index := _index + 1
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            var literal LnsAny
            var mess LnsAny
            literal,mess = argNode.FP.GetLiteral()
            if literal != nil{
                literal_3461 := literal
                {
                    _val := Macro_getLiteralMacroVal_1090_(literal_3461)
                    if !Lns_IsNil( _val ) {
                        val := _val
                        argValMap.Set(index,Lns_getVM().RunLoadedfunc(Macro_toLuaval,Lns_2DDD(val))[0])
                        var declArgNode *Nodes_MacroArgInfo
                        declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                        if argNode.FP.Get_expType() == Ast_builtinTypeSymbol{
                            macroArgValMap.Set(declArgNode.FP.Get_name(),Lns_FromStemGetAt(val,1, false ))
                        } else { 
                            macroArgValMap.Set(declArgNode.FP.Get_name(),val)
                        }
                        macroArgName2ArgNode.Set(declArgNode.FP.Get_name(),argNode)
                    }
                }
            } else {
                var errmess string
                errmess = Lns_getVM().String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(argNode.FP.Get_kind()), mess})
                return nil, errmess
            }
        }
    }
    var _func *Lns_luaValue
    _func = macroInfo.G_func
    var macroVars *Lns_luaValue
    macroVars = Lns_getVM().RunLoadedfunc(_func,Lns_2DDD(macroArgValMap))[0].(*Lns_luaValue)
    self.macroLocalVarMap = Lns_unwrap( macroVars.GetAt("__var"))
    
    {
        _exp1564 := (Lns_unwrap( macroVars.GetAt("__names"))).(*Lns_luaValue)
        _key1564, _val1564 := _exp1564.Get1stFromMap()
        for _key1564 != nil {
            name := _val1564.(string)
            var valInfo *Nodes_MacroValInfo
            
            {
                _valInfo := macroInfo.Symbol2MacroValInfoMap.Get(name)
                if _valInfo == nil{
                    Util_err(Lns_getVM().String_format("not found macro symbol -- %s", []LnsAny{name}))
                } else {
                    valInfo = _valInfo.(*Nodes_MacroValInfo)
                }
            }
            var typeInfo *Ast_TypeInfo
            typeInfo = valInfo.TypeInfo
            var valList LnsAny
            {
                _val := macroVars.GetAt(name)
                if !Lns_IsNil( _val ) {
                    val := _val
                    if Macro_equalsType_1219_(typeInfo, Ast_builtinTypeSymbol){
                        valList = Lns_getVM().RunLoadedfunc(Macro_toList,Lns_2DDD(val))[0].(*Lns_luaValue)
                        
                    } else { 
                        valList = val
                        
                    }
                } else {
                    valList = Lns_getVM().RunLoadedfunc(Macro_toListEmpty,[]LnsAny{})[0].(*Lns_luaValue)
                    
                }
            }
            self.symbol2ValueMapForMacro.Set(name,NewNodes_MacroValInfo(valList, typeInfo, macroArgName2ArgNode.Get(name)))
            _key1564, _val1564 = _exp1564.NextFromMap( _key1564 )
        }
    }
    for _index, _arg := range( macroInfo.FP.GetArgList().Items ) {
        index := _index + 1
        arg := _arg.(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
        if arg.FP.Get_typeInfo().FP.Get_kind() != Ast_TypeInfoKind__DDD{
            var argType *Ast_TypeInfo
            argType = arg.FP.Get_typeInfo()
            var argName string
            argName = arg.FP.Get_name()
            self.symbol2ValueMapForMacro.Set(argName,NewNodes_MacroValInfo(argValMap.Get(index), argType, macroArgName2ArgNode.Get(argName)))
        } else { 
            return nil, "not support ... in macro"
        }
    }
    return &NewMacro_MacroPaser(macroInfo.FP.GetTokenList(), Lns_getVM().String_format("%s:%d:%d: (macro %s)", []LnsAny{streamName, firstToken.Pos.LineNo, firstToken.Pos.Column, macroTypeInfo.FP.GetTxt(nil, nil, nil)}), firstToken.Pos.FP.Get_orgPos()).Parser_Parser, nil
}

// 426: decl @lune.@base.@Macro.MacroCtrl.importMacro
func (self *Macro_MacroCtrl) ImportMacro(processInfo *Ast_ProcessInfo,lnsPath string,macroInfoStem LnsAny,macroTypeInfo *Ast_TypeInfo,typeId2TypeInfo *LnsMap,importedMacroInfoMap *LnsMap) {
    var macroInfo LnsAny
    var err LnsAny
    macroInfo,err = Macro_MacroMetaInfo__fromStem_1068_(macroInfoStem,nil)
    if macroInfo != nil{
        macroInfo_3503 := macroInfo.(*Macro_MacroMetaInfo)
        var orgPos *Types_Position
        if macroInfo_3503.Pos.Len() == 2{
            orgPos = NewTypes_Position(macroInfo_3503.Pos.GetAt(1).(LnsInt), macroInfo_3503.Pos.GetAt(2).(LnsInt), lnsPath)
            
        } else { 
            Util_err("macroInfo.pos is illegal")
        }
        var argList *LnsList
        argList = NewLnsList([]LnsAny{})
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        var symbol2MacroValInfoMap *LnsMap
        symbol2MacroValInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _argInfo := range( macroInfo_3503.ArgList.Items ) {
            argInfo := _argInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var argTypeInfo *Ast_TypeInfo
            argTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(argInfo.TypeId)).(*Ast_TypeInfo)
            argList.Insert(Nodes_MacroArgInfo2Stem(NewNodes_MacroArgInfo(argInfo.Name, argTypeInfo)))
            argNameList.Insert(argInfo.Name)
        }
        for _, _symInfo := range( macroInfo_3503.SymList.Items ) {
            symInfo := _symInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var symTypeInfo *Ast_TypeInfo
            symTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(symInfo.TypeId)).(*Ast_TypeInfo)
            symbol2MacroValInfoMap.Set(symInfo.Name,NewNodes_MacroValInfo(nil, symTypeInfo, nil))
        }
        var tokenList *LnsList
        tokenList = NewLnsList([]LnsAny{})
        var lineNo LnsInt
        lineNo = 0
        var column LnsInt
        column = 1
        for _, _tokenInfo := range( macroInfo_3503.TokenList.Items ) {
            tokenInfo := _tokenInfo.(*LnsList)
            var txt string
            txt = tokenInfo.GetAt(2).(string)
            if txt == "\n"{
                lineNo = lineNo + 1
                
                column = 1
                
            } else { 
                var pos *Types_Position
                pos = Types_Position_create(lineNo, column, Lns_getVM().String_format("macro:%s", []LnsAny{macroInfo_3503.Name}), orgPos)
                tokenList.Insert(Types_Token2Stem(NewTypes_Token(Lns_unwrap( Types_TokenKind__from(Lns_forceCastInt(tokenInfo.GetAt(1)))).(LnsInt), txt, pos, false, nil)))
                column = column + len(txt) + 1
                
            }
        }
        var extMacroInfo *Macro_ExtMacroInfo
        extMacroInfo = NewMacro_ExtMacroInfo(macroInfo_3503.Name, self.macroEval.FP.EvalFromCode(processInfo, macroInfo_3503.Name, argNameList, macroInfo_3503.StmtBlock), symbol2MacroValInfoMap, argList, tokenList)
        self.typeId2MacroInfo.Set(macroTypeInfo.FP.Get_typeId(),&extMacroInfo.Nodes_MacroInfo)
        importedMacroInfoMap.Set(macroTypeInfo.FP.Get_typeId(),&extMacroInfo.Nodes_MacroInfo)
    } else {
        Util_errorLog(Lns_getVM().String_format("macro load fail -- %s: %s ", []LnsAny{macroTypeInfo.FP.GetTxt(nil, nil, nil), Lns_unwrapDefault( err, "").(string)}))
    }
}

// 495: decl @lune.@base.@Macro.MacroCtrl.importMacroInfo
func (self *Macro_MacroCtrl) ImportMacroInfo(importedMacroInfoMap *LnsMap) {
    for _typeId, _macroInfo := range( importedMacroInfoMap.Items ) {
        typeId := _typeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
        macroInfo := _macroInfo.(Nodes_MacroInfoDownCast).ToNodes_MacroInfo()
        self.typeId2MacroInfo.Set(typeId,macroInfo)
    }
}

// 504: decl @lune.@base.@Macro.MacroCtrl.regist
func (self *Macro_MacroCtrl) Regist(processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode,macroScope *Ast_Scope) {
    var macroObj *Lns_luaValue
    macroObj = self.macroEval.FP.Eval(processInfo, node)
    var remap *LnsMap
    remap = NewLnsMap( map[LnsAny]LnsAny{})
    for _name, _macroValInfo := range( self.symbol2ValueMapForMacro.Items ) {
        name := _name.(string)
        macroValInfo := _macroValInfo.(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
        if Macro_equalsType_1219_(macroValInfo.TypeInfo, Ast_builtinTypeEmpty){
            {
                _typeInfo := macroScope.FP.GetTypeInfoChild(name)
                if !Lns_IsNil( _typeInfo ) {
                    typeInfo := _typeInfo.(*Ast_TypeInfo)
                    remap.Set(name,NewNodes_MacroValInfo(macroValInfo.Val, typeInfo, macroValInfo.ArgNode))
                } else {
                    remap.Set(name,macroValInfo)
                }
            }
        } else { 
            remap.Set(name,macroValInfo)
        }
    }
    var macroInfo *Nodes_DefMacroInfo
    macroInfo = NewNodes_DefMacroInfo(macroObj, node.FP.Get_declInfo(), remap)
    self.typeId2MacroInfo.Set(node.FP.Get_expType().FP.Get_typeId(),&macroInfo.Nodes_MacroInfo)
    self.declMacroInfoMap.Set(node.FP.Get_expType().FP.Get_typeId(),&macroInfo.Nodes_MacroInfo)
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = false
    
}

// 608: decl @lune.@base.@Macro.MacroCtrl.expandMacroVal
func (self *Macro_MacroCtrl) ExpandMacroVal(typeNameCtrl *Ast_TypeNameCtrl,scope *Ast_Scope,parser Parser_PushbackParser,token *Types_Token) *Types_Token {
    if self.tokenExpanding{
        return token
    }
    var getToken func() *Types_Token
    getToken = func() *Types_Token {
        self.tokenExpanding = true
        
        var work *Types_Token
        work = parser.GetTokenNoErr()
        self.tokenExpanding = false
        
        return work
    }
    var tokenTxt string
    tokenTxt = token.Txt
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( tokenTxt == ",,") ||
        Lns_GetEnv().SetStackVal( tokenTxt == ",,,") ||
        Lns_GetEnv().SetStackVal( tokenTxt == ",,,,") ).(bool){
        var nextToken *Types_Token
        nextToken = getToken()
        var macroVal *Nodes_MacroValInfo
        
        {
            _macroVal := self.symbol2ValueMapForMacro.Get(nextToken.Txt)
            if _macroVal == nil{
                parser.Error(Lns_getVM().String_format("unknown macro val -- %s", []LnsAny{nextToken.Txt}))
            } else {
                macroVal = _macroVal.(*Nodes_MacroValInfo)
            }
        }
        if tokenTxt == ",,"{
            if Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *LnsList
                txtList = NewLnsList([]LnsAny{})
                {
                    _exp2842 := (Lns_unwrap( macroVal.Val)).(*Lns_luaValue)
                    _key2842, _val2842 := _exp2842.Get1stFromMap()
                    for _key2842 != nil {
                        txt := _val2842.(string)
                        txtList.Insert(txt)
                        _key2842, _val2842 = _exp2842.NextFromMap( _key2842 )
                    }
                }
                Macro_pushbackTxt_1297_(parser, txtList, nextToken.Txt, nextToken.Pos)
            } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ).(bool){
                var pos *Types_Position
                pos = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(macroVal.ArgNode) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_Node).FP.Get_pos()})&&
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Types_Position).FP.Get_RawOrgPos()}))) ||
                    Lns_GetEnv().SetStackVal( nextToken.Pos.FP.Get_RawOrgPos()) ||
                    Lns_GetEnv().SetStackVal( token.Pos.FP.Get_orgPos()) ).(*Types_Position)
                parser.PushbackStr(Lns_getVM().String_format("macroVal %s", []LnsAny{nextToken.Txt}), (Lns_unwrap( macroVal.Val)).(string), pos)
            } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( macroVal.TypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Array) ||
                Lns_GetEnv().SetStackVal( macroVal.TypeInfo.FP.Get_kind() == Ast_TypeInfoKind__List) ).(bool){
                if Macro_equalsType_1219_(macroVal.TypeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat){
                    var pos *Types_Position
                    pos = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(macroVal.ArgNode) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_Node).FP.Get_pos()})&&
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Types_Position).FP.Get_RawOrgPos()}))) ||
                        Lns_GetEnv().SetStackVal( nextToken.Pos.FP.Get_RawOrgPos()) ||
                        Lns_GetEnv().SetStackVal( token.Pos.FP.Get_orgPos()) ).(*Types_Position)
                    var strList *Lns_luaValue
                    strList = (Lns_unwrap( macroVal.Val)).(*Lns_luaValue)
                    {
                        var _from3028 LnsInt = strList.Len()
                        var _to3028 LnsInt = 1
                        _work3028 := _from3028
                        _delta3028 := -1
                        for {
                            if _delta3028 > 0 {
                               if _work3028 > _to3028 { break }
                            } else {
                               if _work3028 < _to3028 { break }
                            }
                            index := _work3028
                            parser.PushbackStr(Lns_getVM().String_format("macroVal %s[%d]", []LnsAny{nextToken.Txt, index}), strList.GetAt(index).(string), pos)
                            _work3028 += _delta3028
                        }
                    }
                } else { 
                    var tokenList *LnsList
                    tokenList = NewLnsList([]LnsAny{})
                    {
                        _argNode := macroVal.ArgNode
                        if !Lns_IsNil( _argNode ) {
                            argNode := _argNode.(*Nodes_Node)
                            if Lns_op_not(argNode.FP.SetupLiteralTokenList(tokenList)){
                                parser.Error(Lns_getVM().String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                            }
                        } else {
                            parser.Error(Lns_getVM().String_format("not support ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    }
                    parser.NewPushback(tokenList)
                }
            } else if macroVal.TypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Enum{
                var enumTypeInfo *Ast_EnumTypeInfo
                enumTypeInfo = Lns_unwrap( Ast_EnumTypeInfoDownCastF(macroVal.TypeInfo.FP.Get_aliasSrc().FP)).(*Ast_EnumTypeInfo)
                var fullname string
                fullname = macroVal.TypeInfo.FP.GetFullName(typeNameCtrl, scope.FP, true)
                var nameList *LnsList
                nameList = NewLnsList([]LnsAny{})
                {
                    _form3171, _param3171, _prev3171 := Lns_getVM().String_gmatch(fullname,"[^%.]+")
                    for {
                        _work3171 := _form3171.(*Lns_luaValue).Call( Lns_2DDD( _param3171, _prev3171 ) )
                        _prev3171 = Lns_getFromMulti(_work3171,0)
                        if Lns_IsNil( _prev3171 ) { break }
                        name := _prev3171.(string)
                        nameList.Insert(name)
                    }
                }
                var enumValInfo *Ast_EnumValInfo
                enumValInfo = Lns_unwrap( enumTypeInfo.FP.Get_val2EnumValInfo().Get(Lns_unwrap( macroVal.Val))).(*Ast_EnumValInfo)
                nextToken = NewTypes_Token(Types_TokenKind__Symb, enumValInfo.FP.Get_name(), nextToken.Pos, false, nil)
                
                parser.PushbackToken(nextToken)
                {
                    var _from3291 LnsInt = nameList.Len()
                    var _to3291 LnsInt = 1
                    _work3291 := _from3291
                    _delta3291 := -1
                    for {
                        if _delta3291 > 0 {
                           if _work3291 > _to3291 { break }
                        } else {
                           if _work3291 < _to3291 { break }
                        }
                        index := _work3291
                        nextToken = NewTypes_Token(Types_TokenKind__Dlmt, ".", nextToken.Pos, false, nil)
                        
                        parser.PushbackToken(nextToken)
                        nextToken = NewTypes_Token(Types_TokenKind__Symb, nameList.GetAt(index).(string), nextToken.Pos, false, nil)
                        
                        parser.PushbackToken(nextToken)
                        _work3291 += _delta3291
                    }
                }
            } else { 
                var tokenList *LnsList
                tokenList = NewLnsList([]LnsAny{})
                {
                    _argNode := macroVal.ArgNode
                    if !Lns_IsNil( _argNode ) {
                        argNode := _argNode.(*Nodes_Node)
                        if Lns_op_not(argNode.FP.SetupLiteralTokenList(tokenList)){
                            parser.Error(Lns_getVM().String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    } else {
                        Macro_expandVal_1292_(tokenList, macroVal.Val, nextToken.Pos)
                    }
                }
                parser.NewPushback(tokenList)
            }
        } else if tokenTxt == ",,,"{
            if Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeString){
                Macro_pushbackTxt_1297_(parser, NewLnsList([]LnsAny{(Lns_unwrap( macroVal.Val)).(string)}), nextToken.Txt, nextToken.Pos)
            } else { 
                parser.Error(Lns_getVM().String_format("',,,' does not support this type -- %s", []LnsAny{macroVal.TypeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        } else if tokenTxt == ",,,,"{
            if Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *Lns_luaValue
                txtList = (Lns_unwrap( macroVal.Val)).(*Lns_luaValue)
                var newToken string
                newToken = ""
                {
                    _exp3479 := txtList
                    _key3479, _val3479 := _exp3479.Get1stFromMap()
                    for _key3479 != nil {
                        txt := _val3479.(string)
                        newToken = Lns_getVM().String_format("%s%s", []LnsAny{newToken, txt})
                        
                        _key3479, _val3479 = _exp3479.NextFromMap( _key3479 )
                    }
                }
                nextToken = NewTypes_Token(Types_TokenKind__Str, Lns_getVM().String_format("'%s'", []LnsAny{newToken}), nextToken.Pos, false, nil)
                
                parser.PushbackToken(nextToken)
            } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ).(bool){
                var txt string
                txt = (Lns_unwrap( macroVal.Val)).(string)
                var rawTxt string
                if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(txt,"^```", nil, nil))){
                    rawTxt = Parser_quoteStr(txt)
                    
                } else { 
                    rawTxt = Parser_quoteStr(txt)
                    
                }
                nextToken = NewTypes_Token(Types_TokenKind__Str, rawTxt, nextToken.Pos, false, nil)
                
                parser.PushbackToken(nextToken)
            } else { 
                parser.Error(Lns_getVM().String_format("not support this symbol -- %s%s", []LnsAny{tokenTxt, nextToken.Txt}))
            }
        }
        nextToken = getToken()
        
        token = nextToken
        
    }
    self.tokenExpanding = false
    
    return token
}

// 780: decl @lune.@base.@Macro.MacroCtrl.expandSymbol
func (self *Macro_MacroCtrl) ExpandSymbol(parser Parser_PushbackParser,prefixToken *Types_Token,exp *Nodes_Node,nodeManager *Nodes_NodeManager,errMessList *LnsList) *Nodes_LiteralStringNode {
    var nextToken *Types_Token
    nextToken = parser.GetTokenNoErr()
    if nextToken.Txt != "~~"{
        parser.PushbackToken(nextToken)
    }
    var format string
    format = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( prefixToken.Txt == ",,,") &&
        Lns_GetEnv().SetStackVal( "' %s '") ||
        Lns_GetEnv().SetStackVal( "\"'%s'\"") ).(string)
    if prefixToken.Txt == ",,"{
        {
            _refNode := Nodes_ExpRefNodeDownCastF(exp.FP)
            if !Lns_IsNil( _refNode ) {
                refNode := _refNode.(*Nodes_ExpRefNode)
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = refNode.FP.Get_symbolInfo()
                var macroInfo LnsAny
                macroInfo = self.symbol2ValueMapForMacro.Get(symbolInfo.FP.Get_name())
                if macroInfo != nil{
                    macroInfo_3679 := macroInfo.(*Nodes_MacroValInfo)
                    var valType *Ast_TypeInfo
                    valType = macroInfo_3679.TypeInfo
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeSymbol)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeExp)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeMultiExp)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                        
                    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( valType.FP.Get_kind() == Ast_TypeInfoKind__List) &&
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat)) ).(bool)){
                        format = "' %s '"
                        
                        exp = &Nodes_ExpMacroStatListNode_create(nodeManager, prefixToken.Pos, self.analyzeInfo.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), exp).Nodes_Node
                        
                    } else if Macro_equalsType_1219_(Ast_builtinTypeString, valType){
                    } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeInt)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(valType, Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                        
                    } else { 
                        errMessList.Insert(Macro_ErrorMess2Stem(NewMacro_ErrorMess(Lns_unwrap( symbolInfo.FP.Get_pos()).(*Types_Position), Lns_getVM().String_format("not support ,, -- %s", []LnsAny{valType.FP.GetTxt(nil, nil, nil)}))))
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeInt)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                        
                    } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeSymbol)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeExp)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeMultiExp)) ||
                        Lns_GetEnv().SetStackVal( Macro_equalsType_1219_(exp.FP.Get_expType(), Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                        
                    }
                }
            }
        }
    }
    var newToken *Types_Token
    newToken = NewTypes_Token(Types_TokenKind__Str, format, prefixToken.Pos, prefixToken.Consecutive, nil)
    var expListNode *Nodes_ExpListNode
    expListNode = Nodes_ExpListNode_create(nodeManager, exp.FP.Get_pos(), self.analyzeInfo.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(), NewLnsList([]LnsAny{Nodes_Node2Stem(exp)}), nil, false)
    var dddNode *Nodes_ExpToDDDNode
    dddNode = Nodes_ExpToDDDNode_create(nodeManager, exp.FP.Get_pos(), self.analyzeInfo.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(), expListNode)
    var literalStr *Nodes_LiteralStringNode
    literalStr = Nodes_LiteralStringNode_create(nodeManager, prefixToken.Pos, self.analyzeInfo.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), newToken, expListNode, Nodes_ExpListNode_create(nodeManager, exp.FP.Get_pos(), self.analyzeInfo.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(), NewLnsList([]LnsAny{Nodes_ExpToDDDNode2Stem(dddNode)}), nil, false), false)
    return literalStr
}

// 873: decl @lune.@base.@Macro.MacroCtrl.registVar
func (self *Macro_MacroCtrl) RegistVar(symbolList *LnsList) {
    for _, _symbolInfo := range( symbolList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var info *Nodes_MacroValInfo
        info = NewNodes_MacroValInfo(nil, symbolInfo.FP.Get_typeInfo(), nil)
        self.symbol2ValueMapForMacro.Set(symbolInfo.FP.Get_name(),info)
    }
}

// 881: decl @lune.@base.@Macro.MacroCtrl.startDecl
func (self *Macro_MacroCtrl) StartDecl() {
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.isDeclaringMacro = true
    
}

// 888: decl @lune.@base.@Macro.MacroCtrl.finishMacroMode
func (self *Macro_MacroCtrl) FinishMacroMode() {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
}

// 894: decl @lune.@base.@Macro.MacroCtrl.startExpandMode
func (self *Macro_MacroCtrl) StartExpandMode(lineNo LnsInt,typeInfo *Ast_TypeInfo,callback Macro_EvalMacroCallback) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(typeInfo, Nodes_MacroMode__Expand)
    
    self.macroCallLineNo = lineNo
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
    callback()
    self.FP.FinishMacroMode()
}

// 906: decl @lune.@base.@Macro.MacroCtrl.startAnalyzeArgMode
func (self *Macro_MacroCtrl) StartAnalyzeArgMode(macroFuncType *Ast_TypeInfo) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(macroFuncType, Nodes_MacroMode__AnalyzeArg)
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}

// 911: decl @lune.@base.@Macro.MacroCtrl.switchMacroMode
func (self *Macro_MacroCtrl) SwitchMacroMode() {
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len() - 1).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}

// 916: decl @lune.@base.@Macro.MacroCtrl.restoreMacroMode
func (self *Macro_MacroCtrl) RestoreMacroMode() {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    
}

// 925: decl @lune.@base.@Macro.MacroCtrl.isInAnalyzeArgMode
func (self *Macro_MacroCtrl) IsInAnalyzeArgMode() bool {
    if self.macroAnalyzeInfoStack.Len() == 0{
        return false
    }
    for _, _info := range( self.macroAnalyzeInfoStack.Items ) {
        info := _info.(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
        if info.FP.Get_mode() == Nodes_MacroMode__AnalyzeArg{
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
func NewMacro_ErrorMess(arg1 *Types_Position, arg2 string) *Macro_ErrorMess {
    obj := &Macro_ErrorMess{}
    obj.FP = obj
    obj.InitMacro_ErrorMess(arg1, arg2)
    return obj
}
func (self *Macro_ErrorMess) InitMacro_ErrorMess(arg1 *Types_Position, arg2 string) {
    self.Pos = arg1
    self.Mess = arg2
}

func Lns_Macro_init() {
    if init_Macro { return }
    init_Macro = true
    Macro__mod__ = "@lune.@base.@Macro"
    Lns_InitMod()
    Lns_Util_init()
    Lns_Nodes_init()
    Lns_Ast_init()
    Lns_Parser_init()
    Lns_Types_init()
    Lns_Formatter_init()
    Macro_toList = Macro_loadCode_1004_("return function( ... ) return { ... } end").(*Lns_luaValue)
    Macro_toListEmpty = Macro_loadCode_1004_("return function() return {} end").(*Lns_luaValue)
    Macro_toLuaval = Macro_loadCode_1004_("return function( val ) return val end").(*Lns_luaValue)
}
func init() {
    init_Macro = false
}
