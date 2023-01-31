// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Types bool
var Types__mod__ string
// decl enum -- Lang 
type Types_Lang = LnsInt
const Types_Lang__C = 4
const Types_Lang__Go = 2
const Types_Lang__Lua = 1
const Types_Lang__Python = 3
const Types_Lang__Same = 0
var Types_LangList_ = NewLnsList( []LnsAny {
  Types_Lang__Same,
  Types_Lang__Lua,
  Types_Lang__Go,
  Types_Lang__Python,
  Types_Lang__C,
})
func Types_Lang_get__allList(_env *LnsEnv) *LnsList{
    return Types_LangList_
}
var Types_LangMap_ = map[LnsInt]string {
  Types_Lang__C: "Lang.C",
  Types_Lang__Go: "Lang.Go",
  Types_Lang__Lua: "Lang.Lua",
  Types_Lang__Python: "Lang.Python",
  Types_Lang__Same: "Lang.Same",
}
func Types_Lang__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func Types_TokenKind_get__allList(_env *LnsEnv) *LnsList{
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
func Types_TokenKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Types_TokenKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Types_TokenKind_getTxt(arg1 LnsInt) string {
    return Types_TokenKindMap_[arg1];
}
var Types_nonePos Types_Position
var Types_noneToken *Types_Token
var Types_defaultParserPipeSize LnsInt
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
// decl alge -- ParserSrc
type Types_ParserSrc = LnsAny
type Types_ParserSrc__LnsCode struct{
Val1 string
Val2 string
Val3 LnsAny
}
func (self *Types_ParserSrc__LnsCode) GetTxt() string {
return "ParserSrc.LnsCode"
}
type Types_ParserSrc__LnsPath struct{
Val1 LnsAny
Val2 string
Val3 string
Val4 LnsAny
}
func (self *Types_ParserSrc__LnsPath) GetTxt() string {
return "ParserSrc.LnsPath"
}
type Types_ParserSrc__Parser struct{
Val1 string
Val2 bool
Val3 string
Val4 LnsAny
}
func (self *Types_ParserSrc__Parser) GetTxt() string {
return "ParserSrc.Parser"
}
// 144: decl @lune.@base.@Types.TransCtrlInfo.create_normal
func Types_TransCtrlInfo_create_normal(_env *LnsEnv) *Types_TransCtrlInfo {
    return NewTypes_TransCtrlInfo(_env)
}
// 170: decl @lune.@base.@Types.Position.get_orgPos
func (self Types_Position) Get_orgPos(_env *LnsEnv) Types_Position {
    {
        __exp := self.OrgPos
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(Types_Position)
            return _exp.Get_orgPos(_env)
        }
    }
    return self
}
// 177: decl @lune.@base.@Types.Position.get_RawOrgPos
func (self Types_Position) Get_RawOrgPos(_env *LnsEnv) LnsAny {
    return self.OrgPos
}
// 181: decl @lune.@base.@Types.Position.create
func Types_Position_create(_env *LnsEnv, lineNo LnsInt,column LnsInt,streamName string,orgPos LnsAny) Types_Position {
    return NewTypes_Position(_env, lineNo, column, streamName, orgPos)
}
// 187: decl @lune.@base.@Types.Position.getDisplayTxt
func (self Types_Position) GetDisplayTxt(_env *LnsEnv) string {
    var txt string
    txt = _env.GetVM().String_format("%s:%d:%d", []LnsAny{self.StreamName, self.LineNo, self.Column})
    var orgPos Types_Position
    orgPos = self.Get_orgPos(_env)
    if self != orgPos{
        var txt2 string
        txt2 = _env.GetVM().String_format("%s:%d:%d", []LnsAny{orgPos.StreamName, orgPos.LineNo, orgPos.Column})
        return _env.GetVM().String_format("%s: (%s)", []LnsAny{txt2, txt})
    }
    return txt
}
// 197: decl @lune.@base.@Types.Position.comp
func (self Types_Position) Comp(_env *LnsEnv, other Types_Position) LnsInt {
    if self.StreamName < other.StreamName{
        return -1
    }
    if self.StreamName > other.StreamName{
        return 1
    }
    if self.LineNo < other.LineNo{
        return -1
    }
    if self.LineNo > other.LineNo{
        return 1
    }
    if self.Column < other.Column{
        return -1
    }
    if self.Column > other.Column{
        return 1
    }
    var orgPos Types_Position
    var otherOrgPos Types_Position
    
    {
        _orgPos, _otherOrgPos := self.OrgPos, other.OrgPos
        if _orgPos == nil || _otherOrgPos == nil{
            return 0
        } else {
            orgPos = _orgPos.(Types_Position)
            otherOrgPos = _otherOrgPos.(Types_Position)
        }
    }
    return orgPos.Comp(_env, otherOrgPos)
}
// 268: decl @lune.@base.@Types.Token.getExcludedDelimitTxt
func (self *Types_Token) GetExcludedDelimitTxt(_env *LnsEnv) string {
    if self.Kind != Types_TokenKind__Str{
        return self.Txt
    }
    if _switch0 := LnsInt(self.Txt[1-1]); _switch0 == 39 || _switch0 == 34 {
        return _env.GetVM().String_sub(self.Txt,2, len(self.Txt) - 1)
    } else if _switch0 == 96 {
        return _env.GetVM().String_sub(self.Txt,1 + 3, len(self.Txt) - 3)
    }
    Util_err(_env, _env.GetVM().String_format("illegal delimit -- %s", []LnsAny{self.Txt}))
// insert a dummy
    return ""
}
// 288: decl @lune.@base.@Types.Token.getLineCount
func (self *Types_Token) GetLineCount(_env *LnsEnv) LnsInt {
    var count LnsInt
    count = 1
        {
            _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(self.Txt,"\n")
            for {
                _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                if Lns_IsNil( _applyPrev0 ) { break }
                count = count + 1
            }
        }
    return count
}
// declaration Class -- AltBase
type Types_AltBaseMtd interface {
}
type Types_AltBase struct {
    val LnsAny
    mess LnsAny
    FP Types_AltBaseMtd
}
func Types_AltBase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_AltBase).FP
}
type Types_AltBaseDownCast interface {
    ToTypes_AltBase() *Types_AltBase
}
func Types_AltBaseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_AltBaseDownCast)
    if ok { return work.ToTypes_AltBase() }
    return nil
}
func (obj *Types_AltBase) ToTypes_AltBase() *Types_AltBase {
    return obj
}
func NewTypes_AltBase(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny) *Types_AltBase {
    obj := &Types_AltBase{}
    obj.FP = obj
    obj.InitTypes_AltBase(_env, arg1, arg2)
    return obj
}
func (self *Types_AltBase) InitTypes_AltBase(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny) {
    self.val = arg1
    self.mess = arg2
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
    ValidCheckingMutable bool
    LegacyMutableControl bool
    ValidAstDetailError bool
    ValidAsyncCtrl bool
    DefaultAsync bool
    Testing bool
    ValidMultiPhaseTransUnit bool
    ThreadPerUnitThread LnsInt
    MacroAsyncParseStmtLen LnsInt
    ValidMacroAsync bool
    UseWaiter bool
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
func NewTypes_TransCtrlInfo(_env *LnsEnv) *Types_TransCtrlInfo {
    obj := &Types_TransCtrlInfo{}
    obj.FP = obj
    obj.InitTypes_TransCtrlInfo(_env)
    return obj
}
// 122: DeclConstr
func (self *Types_TransCtrlInfo) InitTypes_TransCtrlInfo(_env *LnsEnv) {
    self.UseWaiter = true
    self.MacroAsyncParseStmtLen = 500
    self.WarningShadowing = false
    self.CompatComment = false
    self.CheckingDefineAbbr = true
    self.StopByWarning = false
    self.UptodateMode = Types_CheckingUptodateMode__Touch_Obj
    self.ValidLuaval = false
    self.DefaultLazy = false
    self.ValidCheckingMutable = true
    self.LegacyMutableControl = false
    self.ValidAstDetailError = false
    self.ValidAsyncCtrl = false
    self.DefaultAsync = false
    self.Testing = false
    self.ValidMultiPhaseTransUnit = true
    self.ThreadPerUnitThread = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( true) &&
        _env.SetStackVal( 5) ||
        _env.SetStackVal( 0) ).(LnsInt)
    self.ValidMacroAsync = false
}


// declaration Class -- Position
type Types_PositionMtd interface {
    ToMap() *LnsMap
    Comp(_env *LnsEnv, arg1 Types_Position) LnsInt
    GetDisplayTxt(_env *LnsEnv) string
    Get_RawOrgPos(_env *LnsEnv) LnsAny
    Get_orgPos(_env *LnsEnv) Types_Position
}
type Types_Position struct {
    LineNo LnsInt
    Column LnsInt
    StreamName string
    OrgPos LnsAny
}
func Types_Position2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(Types_Position)
}
type Types_PositionDownCast interface {
    ToTypes_Position() Types_Position
}
func Types_PositionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_PositionDownCast)
    if ok { return work.ToTypes_Position() }
    return nil
}
func (obj Types_Position) ToTypes_Position() Types_Position {
    return obj
}
func NewTypes_Position(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 string, arg4 LnsAny) Types_Position {
    obj := Types_Position{}
    obj.InitTypes_Position(_env, arg1, arg2, arg3, arg4)
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
func Types_Position__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Position_FromMap( arg1, paramList )
}
func Types_Position__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    newObj := Types_Position{}
    return Types_Position_FromMapMain( newObj, objMap, paramList )
}
func Types_Position_FromMapMain( newObj Types_Position, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
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
// 163: DeclConstr
func (self *Types_Position) InitTypes_Position(_env *LnsEnv, lineNo LnsInt,column LnsInt,streamName string,orgPos LnsAny) {
    self.LineNo = lineNo
    self.Column = column
    self.StreamName = streamName
    self.OrgPos = orgPos
}


// declaration Class -- Token
type Types_TokenMtd interface {
    ToMap() *LnsMap
    GetExcludedDelimitTxt(_env *LnsEnv) string
    GetLineCount(_env *LnsEnv) LnsInt
    Get_commentList(_env *LnsEnv) *LnsList
}
type Types_Token struct {
    Kind LnsInt
    Txt string
    Pos Types_Position
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
func NewTypes_Token(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 Types_Position, arg4 bool, arg5 LnsAny) *Types_Token {
    obj := &Types_Token{}
    obj.FP = obj
    obj.InitTypes_Token(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Types_Token) Get_commentList(_env *LnsEnv) *LnsList{ return self.commentList }
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
func Types_Token__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Types_Token_FromMap( arg1, paramList )
}
func Types_Token__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
       newObj.Pos = conv.(Types_Position)
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
// 258: DeclConstr
func (self *Types_Token) InitTypes_Token(_env *LnsEnv, kind LnsInt,txt string,pos Types_Position,consecutive bool,commentList LnsAny) {
    self.Kind = kind
    self.Txt = txt
    self.Pos = pos
    self.Consecutive = consecutive
    self.commentList = Lns_unwrapDefault( commentList, NewLnsList([]LnsAny{})).(*LnsList)
}


// declaration Class -- StdinFile
type Types_StdinFileMtd interface {
    Get_mod(_env *LnsEnv) string
    Get_txt(_env *LnsEnv) string
}
type Types_StdinFile struct {
    mod string
    txt string
    FP Types_StdinFileMtd
}
func Types_StdinFile2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_StdinFile).FP
}
type Types_StdinFileDownCast interface {
    ToTypes_StdinFile() *Types_StdinFile
}
func Types_StdinFileDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_StdinFileDownCast)
    if ok { return work.ToTypes_StdinFile() }
    return nil
}
func (obj *Types_StdinFile) ToTypes_StdinFile() *Types_StdinFile {
    return obj
}
func NewTypes_StdinFile(_env *LnsEnv, arg1 string, arg2 string) *Types_StdinFile {
    obj := &Types_StdinFile{}
    obj.FP = obj
    obj.InitTypes_StdinFile(_env, arg1, arg2)
    return obj
}
func (self *Types_StdinFile) InitTypes_StdinFile(_env *LnsEnv, arg1 string, arg2 string) {
    self.mod = arg1
    self.txt = arg2
}
func (self *Types_StdinFile) Get_mod(_env *LnsEnv) string{ return self.mod }
func (self *Types_StdinFile) Get_txt(_env *LnsEnv) string{ return self.txt }

func Lns_Types_init(_env *LnsEnv) {
    if init_Types { return }
    init_Types = true
    Types__mod__ = "@lune.@base.@Types"
    Lns_InitMod()
    Lns_Util_init(_env)
    Types_nonePos = NewTypes_Position(_env, 0, -1, "eof", nil)
    Types_noneToken = NewTypes_Token(_env, Types_TokenKind__Eof, "", Types_nonePos, false, NewLnsList([]LnsAny{}))
    Types_defaultParserPipeSize = 100
}
func init() {
    init_Types = false
}
