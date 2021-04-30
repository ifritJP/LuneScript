// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Types bool
var Types__mod__ string
// decl enum -- Lang 
type Types_Lang = LnsInt
const Types_Lang__C = 3
const Types_Lang__Go = 2
const Types_Lang__Lua = 1
const Types_Lang__Same = 0
var Types_LangList_ = NewLnsList( []LnsAny {
  Types_Lang__Same,
  Types_Lang__Lua,
  Types_Lang__Go,
  Types_Lang__C,
})
func Types_Lang_get__allList() *LnsList{
    return Types_LangList_
}
var Types_LangMap_ = map[LnsInt]string {
  Types_Lang__C: "Lang.C",
  Types_Lang__Go: "Lang.Go",
  Types_Lang__Lua: "Lang.Lua",
  Types_Lang__Same: "Lang.Same",
}
func Types_Lang__from(arg1 LnsInt) LnsAny{
    if _, ok := Types_LangMap_[arg1]; ok { return arg1 }
    return nil
}

func Types_Lang_getTxt(arg1 LnsInt) string {
    return Types_LangMap_[arg1];
}
// decl enum -- TokenKind 
type Types_TokenKind = LnsInt
const Types_TokenKind__Char = 4
const Types_TokenKind__Cmnt = 0
const Types_TokenKind__Dlmt = 6
const Types_TokenKind__Eof = 11
const Types_TokenKind__Int = 2
const Types_TokenKind__Kywd = 7
const Types_TokenKind__Ope = 8
const Types_TokenKind__Real = 3
const Types_TokenKind__Sheb = 10
const Types_TokenKind__Str = 1
const Types_TokenKind__Symb = 5
const Types_TokenKind__Type = 9
var Types_TokenKindList_ = NewLnsList( []LnsAny {
  Types_TokenKind__Cmnt,
  Types_TokenKind__Str,
  Types_TokenKind__Int,
  Types_TokenKind__Real,
  Types_TokenKind__Char,
  Types_TokenKind__Symb,
  Types_TokenKind__Dlmt,
  Types_TokenKind__Kywd,
  Types_TokenKind__Ope,
  Types_TokenKind__Type,
  Types_TokenKind__Sheb,
  Types_TokenKind__Eof,
})
func Types_TokenKind_get__allList() *LnsList{
    return Types_TokenKindList_
}
var Types_TokenKindMap_ = map[LnsInt]string {
  Types_TokenKind__Char: "TokenKind.Char",
  Types_TokenKind__Cmnt: "TokenKind.Cmnt",
  Types_TokenKind__Dlmt: "TokenKind.Dlmt",
  Types_TokenKind__Eof: "TokenKind.Eof",
  Types_TokenKind__Int: "TokenKind.Int",
  Types_TokenKind__Kywd: "TokenKind.Kywd",
  Types_TokenKind__Ope: "TokenKind.Ope",
  Types_TokenKind__Real: "TokenKind.Real",
  Types_TokenKind__Sheb: "TokenKind.Sheb",
  Types_TokenKind__Str: "TokenKind.Str",
  Types_TokenKind__Symb: "TokenKind.Symb",
  Types_TokenKind__Type: "TokenKind.Type",
}
func Types_TokenKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Types_TokenKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Types_TokenKind_getTxt(arg1 LnsInt) string {
    return Types_TokenKindMap_[arg1];
}
var Types_noneToken *Types_Token
// decl alge -- CheckingUptodateMode
type Types_CheckingUptodateMode = LnsAny
type Types_CheckingUptodateMode__Force1 struct{
Val1 string
}
func (self *Types_CheckingUptodateMode__Force1) GetTxt() string {
return "CheckingUptodateMode.Force1"
}
type Types_CheckingUptodateMode__ForceAll struct{
}
var Types_CheckingUptodateMode__ForceAll_Obj = &Types_CheckingUptodateMode__ForceAll{}
func (self *Types_CheckingUptodateMode__ForceAll) GetTxt() string {
return "CheckingUptodateMode.ForceAll"
}
type Types_CheckingUptodateMode__Normal struct{
}
var Types_CheckingUptodateMode__Normal_Obj = &Types_CheckingUptodateMode__Normal{}
func (self *Types_CheckingUptodateMode__Normal) GetTxt() string {
return "CheckingUptodateMode.Normal"
}
type Types_CheckingUptodateMode__Touch struct{
}
var Types_CheckingUptodateMode__Touch_Obj = &Types_CheckingUptodateMode__Touch{}
func (self *Types_CheckingUptodateMode__Touch) GetTxt() string {
return "CheckingUptodateMode.Touch"
}
// declaration Class -- TransCtrlInfo
type Types_TransCtrlInfoMtd interface {
}
type Types_TransCtrlInfo struct {
    WarningShadowing bool
    CompatComment bool
    CheckingDefineAbbr bool
    StopByWarning bool
    UptodateMode LnsAny
    ValidLuaval bool
    DefaultLazy bool
    FP Types_TransCtrlInfoMtd
}
func Types_TransCtrlInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_TransCtrlInfo).FP
}
type Types_TransCtrlInfoDownCast interface {
    ToTypes_TransCtrlInfo() *Types_TransCtrlInfo
}
func Types_TransCtrlInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_TransCtrlInfoDownCast)
    if ok { return work.ToTypes_TransCtrlInfo() }
    return nil
}
func (obj *Types_TransCtrlInfo) ToTypes_TransCtrlInfo() *Types_TransCtrlInfo {
    return obj
}
func NewTypes_TransCtrlInfo(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 LnsAny, arg6 bool, arg7 bool) *Types_TransCtrlInfo {
    obj := &Types_TransCtrlInfo{}
    obj.FP = obj
    obj.InitTypes_TransCtrlInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Types_TransCtrlInfo) InitTypes_TransCtrlInfo(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 LnsAny, arg6 bool, arg7 bool) {
    self.WarningShadowing = arg1
    self.CompatComment = arg2
    self.CheckingDefineAbbr = arg3
    self.StopByWarning = arg4
    self.UptodateMode = arg5
    self.ValidLuaval = arg6
    self.DefaultLazy = arg7
}
// 60: decl @lune.@base.@Types.TransCtrlInfo.create_normal
func Types_TransCtrlInfo_create_normal() *Types_TransCtrlInfo {
    return NewTypes_TransCtrlInfo(false, false, true, false, Types_CheckingUptodateMode__Touch_Obj, false, false)
}


// declaration Class -- Position
type Types_PositionMtd interface {
    ToMap() *LnsMap
    Get_RawOrgPos() LnsAny
    Get_orgPos() *Types_Position
}
type Types_Position struct {
    LineNo LnsInt
    Column LnsInt
    StreamName string
    OrgPos LnsAny
    FP Types_PositionMtd
}
func Types_Position2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_Position).FP
}
type Types_PositionDownCast interface {
    ToTypes_Position() *Types_Position
}
func Types_PositionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_PositionDownCast)
    if ok { return work.ToTypes_Position() }
    return nil
}
func (obj *Types_Position) ToTypes_Position() *Types_Position {
    return obj
}
func NewTypes_Position(arg1 LnsInt, arg2 LnsInt, arg3 string) *Types_Position {
    obj := &Types_Position{}
    obj.FP = obj
    obj.InitTypes_Position(arg1, arg2, arg3)
    return obj
}
func (self *Types_Position) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["lineNo"] = Lns_ToCollection( self.LineNo )
    obj.Items["column"] = Lns_ToCollection( self.Column )
    obj.Items["streamName"] = Lns_ToCollection( self.StreamName )
    obj.Items["orgPos"] = Lns_ToCollection( self.OrgPos )
    return obj
}
func (self *Types_Position) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Types_Position__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Position_FromMap( arg1, paramList )
}
func Types_Position__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Position_FromMap( arg1, paramList )
}
func Types_Position_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Types_Position_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Types_Position_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Types_Position{}
    newObj.FP = newObj
    return Types_Position_FromMapMain( newObj, objMap, paramList )
}
func Types_Position_FromMapMain( newObj *Types_Position, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["lineNo"], false, nil); !ok {
       return false,nil,"lineNo:" + mess.(string)
    } else {
       newObj.LineNo = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["column"], false, nil); !ok {
       return false,nil,"column:" + mess.(string)
    } else {
       newObj.Column = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["streamName"], false, nil); !ok {
       return false,nil,"streamName:" + mess.(string)
    } else {
       newObj.StreamName = conv.(string)
    }
    if ok,conv,mess := Types_Position_FromMapSub( objMap.Items["orgPos"], true, nil); !ok {
       return false,nil,"orgPos:" + mess.(string)
    } else {
       newObj.OrgPos = conv
    }
    return true, newObj, nil
}
// 79: DeclConstr
func (self *Types_Position) InitTypes_Position(lineNo LnsInt,column LnsInt,streamName string) {
    self.LineNo = lineNo
    
    self.Column = column
    
    self.StreamName = streamName
    
    self.OrgPos = nil
    
}

// 86: decl @lune.@base.@Types.Position.get_orgPos
func (self *Types_Position) Get_orgPos() *Types_Position {
    {
        __exp := self.OrgPos
        if __exp != nil {
            _exp := __exp.(*Types_Position)
            return _exp.FP.Get_orgPos()
        }
    }
    return self
}

// 93: decl @lune.@base.@Types.Position.get_RawOrgPos
func (self *Types_Position) Get_RawOrgPos() LnsAny {
    return self.OrgPos
}

// 97: decl @lune.@base.@Types.Position.create
func Types_Position_create(lineNo LnsInt,column LnsInt,streamName string,orgPos LnsAny) *Types_Position {
    var pos *Types_Position
    pos = NewTypes_Position(lineNo, column, streamName)
    pos.OrgPos = orgPos
    
    return pos
}


// declaration Class -- Token
type Types_TokenMtd interface {
    ToMap() *LnsMap
    GetExcludedDelimitTxt() string
    GetLineCount() LnsInt
    Get_commentList() *LnsList
    Set_commentList(arg1 *LnsList)
}
type Types_Token struct {
    Kind LnsInt
    Txt string
    Pos *Types_Position
    Consecutive bool
    commentList *LnsList
    FP Types_TokenMtd
}
func Types_Token2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_Token).FP
}
type Types_TokenDownCast interface {
    ToTypes_Token() *Types_Token
}
func Types_TokenDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_TokenDownCast)
    if ok { return work.ToTypes_Token() }
    return nil
}
func (obj *Types_Token) ToTypes_Token() *Types_Token {
    return obj
}
func NewTypes_Token(arg1 LnsInt, arg2 string, arg3 *Types_Position, arg4 bool, arg5 LnsAny) *Types_Token {
    obj := &Types_Token{}
    obj.FP = obj
    obj.InitTypes_Token(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Types_Token) Get_commentList() *LnsList{ return self.commentList }
func (self *Types_Token) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["kind"] = Lns_ToCollection( self.Kind )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["pos"] = Lns_ToCollection( self.Pos )
    obj.Items["consecutive"] = Lns_ToCollection( self.Consecutive )
    obj.Items["commentList"] = Lns_ToCollection( self.commentList )
    return obj
}
func (self *Types_Token) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Types_Token__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Token_FromMap( arg1, paramList )
}
func Types_Token__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Token_FromMap( arg1, paramList )
}
func Types_Token_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Types_Token_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Types_Token_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Types_Token{}
    newObj.FP = newObj
    return Types_Token_FromMapMain( newObj, objMap, paramList )
}
func Types_Token_FromMapMain( newObj *Types_Token, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["kind"], false, nil); !ok {
       return false,nil,"kind:" + mess.(string)
    } else {
       newObj.Kind = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Types_Position_FromMapSub( objMap.Items["pos"], false, nil); !ok {
       return false,nil,"pos:" + mess.(string)
    } else {
       newObj.Pos = conv.(*Types_Position)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["consecutive"], false, nil); !ok {
       return false,nil,"consecutive:" + mess.(string)
    } else {
       newObj.Consecutive = conv.(bool)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["commentList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Types_Token_FromMapSub, false,nil}}); !ok {
       return false,nil,"commentList:" + mess.(string)
    } else {
       newObj.commentList = conv.(*LnsList)
    }
    return true, newObj, nil
}
// 130: DeclConstr
func (self *Types_Token) InitTypes_Token(kind LnsInt,txt string,pos *Types_Position,consecutive bool,commentList LnsAny) {
    self.Kind = kind
    
    self.Txt = txt
    
    self.Pos = pos
    
    self.Consecutive = consecutive
    
    self.commentList = Lns_unwrapDefault( commentList, NewLnsList([]LnsAny{})).(*LnsList)
    
}

// 140: decl @lune.@base.@Types.Token.getExcludedDelimitTxt
func (self *Types_Token) GetExcludedDelimitTxt() string {
    if self.Kind != Types_TokenKind__Str{
        return self.Txt
    }
    if _switch342 := LnsInt(self.Txt[1-1]); _switch342 == 39 || _switch342 == 34 {
        return Lns_getVM().String_sub(self.Txt,2, len(self.Txt) - 1)
    } else if _switch342 == 96 {
        return Lns_getVM().String_sub(self.Txt,1 + 3, len(self.Txt) - 3)
    }
    panic(Lns_getVM().String_format("illegal delimit -- %s", []LnsAny{self.Txt}))
// insert a dummy
    return ""
}

// 155: decl @lune.@base.@Types.Token.set_commentList
func (self *Types_Token) Set_commentList(commentList *LnsList) {
    self.commentList = commentList
    
}

// 159: decl @lune.@base.@Types.Token.getLineCount
func (self *Types_Token) GetLineCount() LnsInt {
    var count LnsInt
    count = 1
    {
        _form405, _param405, _prev405 := Lns_getVM().String_gmatch(self.Txt,"\n")
        for {
            _work405 := _form405.(*Lns_luaValue).Call( Lns_2DDD( _param405, _prev405 ) )
            _prev405 = Lns_getFromMulti(_work405,0)
            if Lns_IsNil( _prev405 ) { break }
            count = count + 1
            
        }
    }
    return count
}


func Lns_Types_init() {
    if init_Types { return }
    init_Types = true
    Types__mod__ = "@lune.@base.@Types"
    Lns_InitMod()
    Types_noneToken = NewTypes_Token(Types_TokenKind__Eof, "", NewTypes_Position(0, -1, "eof"), false, NewLnsList([]LnsAny{}))
}
func init() {
    init_Types = false
}
