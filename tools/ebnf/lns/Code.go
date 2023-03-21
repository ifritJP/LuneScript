// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Code bool
var Code__mod__ string
// 8: decl @lns.@Code.log
func Code_log_0_(_env *LnsEnv, ddd []LnsAny) {
}


// 40: decl @lns.@Code.CodeTokenizer.getToken
func (self *Code_CodeTokenizer) GetToken(_env *LnsEnv) *LnsTypes.Types_Token {
    if self.tokenIndex <= self.tokenList.Len(){
        var token *LnsTypes.Types_Token
        token = self.tokenList.GetAt(self.tokenIndex).(LnsTypes.Types_TokenDownCast).ToTypes_Token()
        self.tokenIndex = self.tokenIndex + 1
        return token
    }
    var Code_process func(_env *LnsEnv, token *LnsTypes.Types_Token) *LnsTypes.Types_Token
    Code_process = func(_env *LnsEnv, token *LnsTypes.Types_Token) *LnsTypes.Types_Token {
        var result *LnsTypes.Types_Token
        if self.commentList.Len() == 0{
            result = token
        } else { 
            var newToken *LnsTypes.Types_Token
            newToken = LnsTypes.NewTypes_Token(_env, token.Kind, token.Txt, token.Pos, token.Consecutive, self.commentList)
            self.commentList = NewLnsList2_[*LnsTypes.Types_Token]([]*LnsTypes.Types_Token{})
            result = newToken
        }
        self.tokenList.Insert(LnsTypes.Types_Token2Stem(result))
        self.tokenIndex = self.tokenList.Len() + 1
        return result
    }
    for  {
        var token *LnsTypes.Types_Token
        
        {
            _token := self.tokenizer.FP.GetToken(_env)
            if _token == nil{
                token = LnsTypes.Types_noneToken
            } else {
                token = _token.(*LnsTypes.Types_Token)
            }
        }
        if _switch0 := token.Kind; _switch0 == LnsTypes.Types_TokenKind__Eof {
            return Code_process(_env, token)
        } else if _switch0 == LnsTypes.Types_TokenKind__Cmnt {
            self.commentList.Insert(token)
        } else {
            return Code_process(_env, token)
        }
    }
// insert a dummy
    return nil
}
// 83: decl @lns.@Code.CodeTokenizer.pushback
func (self *Code_CodeTokenizer) Pushback(_env *LnsEnv) {
    __func__ := "@lns.@Code.CodeTokenizer.pushback"
    Code_log_0_(_env, Lns_2DDD(__func__, self.tokenList.Len(), self.tokenIndex))
    if self.tokenIndex <= 1{
        panic("illegal tokenIndex")
    }
    self.tokenIndex = self.tokenIndex - 1
}
// 91: decl @lns.@Code.CodeTokenizer.peekToken
func (self *Code_CodeTokenizer) PeekToken(_env *LnsEnv) *LnsTypes.Types_Token {
    if self.tokenIndex <= self.tokenList.Len(){
        var token *LnsTypes.Types_Token
        token = self.tokenList.GetAt(self.tokenIndex).(LnsTypes.Types_TokenDownCast).ToTypes_Token()
        return token
    }
    var token *LnsTypes.Types_Token
    token = self.FP.GetToken(_env)
    self.FP.Pushback(_env)
    return token
}
// 121: decl @lns.@Code.CodeGenerator.output
func (self *Code_CodeGenerator) Output(_env *LnsEnv, codeCore Code_CodeCore) {
    codeCore.OutputCode(_env, self)
}
// 124: decl @lns.@Code.CodeGenerator.outputElement
func (self *Code_CodeGenerator) OutputElement(_env *LnsEnv, elementName LnsAny,codeCoreList *LnsList) {
    for _, _core := range( codeCoreList.Items ) {
        core := _core.(Code_CodeCore)
        core.OutputCode(_env, self)
    }
}
// 129: decl @lns.@Code.CodeGenerator.outputToken
func (self *Code_CodeGenerator) OutputToken(_env *LnsEnv, token *LnsTypes.Types_Token) {
    if self.prevToken.Pos.LineNo != token.Pos.LineNo{
        self.stream.Write(_env, "\n")
    } else { 
        if Lns_op_not(token.Consecutive){
            self.stream.Write(_env, " ")
        }
    }
    self.stream.Write(_env, token.Txt)
    self.tokenList.Insert(LnsTypes.Types_Token2Stem(token))
    self.prevToken = token
}
// 146: decl @lns.@Code.CodeCoreStat.outputCode
func (self *Code_CodeCoreStat) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    for _, _token := range( self.list.Items ) {
        token := _token.(LnsTypes.Types_TokenDownCast).ToTypes_Token()
        stream.FP.OutputToken(_env, token)
    }
}
// 151: decl @lns.@Code.CodeCoreStat.pushback
func (self *Code_CodeCoreStat) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    for range( self.list.Items ) {
        tokenizer.FP.Pushback(_env)
    }
}
// 161: decl @lns.@Code.CodeCoreBuiltin.outputCode
func (self *Code_CodeCoreBuiltin) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputToken(_env, self.token)
}
// 164: decl @lns.@Code.CodeCoreBuiltin.pushback
func (self *Code_CodeCoreBuiltin) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    tokenizer.FP.Pushback(_env)
}
// 172: decl @lns.@Code.CodeCoreToken.outputCode
func (self *Code_CodeCoreToken) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputToken(_env, self.token)
}
// 175: decl @lns.@Code.CodeCoreToken.pushback
func (self *Code_CodeCoreToken) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    __func__ := "@lns.@Code.CodeCoreToken.pushback"
    Code_log_0_(_env, Lns_2DDD(__func__, self.token.Txt))
    tokenizer.FP.Pushback(_env)
}
// 183: decl @lns.@Code.CodeCoreList.outputCode
func (self *Code_CodeCoreList) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputElement(_env, self.elementName, self.codeCoreList)
}
// 186: decl @lns.@Code.CodeCoreList.pushback
func (self *Code_CodeCoreList) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    __func__ := "@lns.@Code.CodeCoreList.pushback"
    Code_log_0_(_env, Lns_2DDD(__func__, self.elementName))
    {
        var _forFrom0 LnsInt = self.codeCoreList.Len()
        var _forTo0 LnsInt = 1
        _forWork0 := _forFrom0
        _forDelta0 := -1
        for {
            if _forDelta0 > 0 {
               if _forWork0 > _forTo0 { break }
            } else {
               if _forWork0 < _forTo0 { break }
            }
            index := _forWork0
            var codeCore Code_CodeCore
            codeCore = self.codeCoreList.GetAt(index).(Code_CodeCore)
            codeCore.Pushback(_env, tokenizer)
            _forWork0 += _forDelta0
        }
    }
}
// declaration Class -- CodeTokenizer
type Code_CodeTokenizerMtd interface {
    GetToken(_env *LnsEnv) *LnsTypes.Types_Token
    PeekToken(_env *LnsEnv) *LnsTypes.Types_Token
    Pushback(_env *LnsEnv)
}
type Code_CodeTokenizer struct {
    tokenizer *LnsTokenizer.Tokenizer_Tokenizer
    tokenList *LnsList
    tokenIndex LnsInt
    commentList *LnsList2_[*LnsTypes.Types_Token]
    FP Code_CodeTokenizerMtd
}
func Code_CodeTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeTokenizer).FP
}
func Code_CodeTokenizer_toSlice(slice []LnsAny) []*Code_CodeTokenizer {
    ret := make([]*Code_CodeTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeTokenizerDownCast).ToCode_CodeTokenizer()
    }
    return ret
}
type Code_CodeTokenizerDownCast interface {
    ToCode_CodeTokenizer() *Code_CodeTokenizer
}
func Code_CodeTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeTokenizerDownCast)
    if ok { return work.ToCode_CodeTokenizer() }
    return nil
}
func (obj *Code_CodeTokenizer) ToCode_CodeTokenizer() *Code_CodeTokenizer {
    return obj
}
func NewCode_CodeTokenizer(_env *LnsEnv, arg1 *LnsTokenizer.Tokenizer_Tokenizer) *Code_CodeTokenizer {
    obj := &Code_CodeTokenizer{}
    obj.FP = obj
    obj.InitCode_CodeTokenizer(_env, arg1)
    return obj
}
// 28: DeclConstr
func (self *Code_CodeTokenizer) InitCode_CodeTokenizer(_env *LnsEnv, tokenizer *LnsTokenizer.Tokenizer_Tokenizer) {
    self.tokenizer = tokenizer
    self.tokenList = NewLnsList([]LnsAny{})
    self.commentList = NewLnsList2_[*LnsTypes.Types_Token]([]*LnsTypes.Types_Token{})
    self.tokenIndex = 1
}


type Code_CodeCore interface {
        OutputCode(_env *LnsEnv, arg1 *Code_CodeGenerator)
        Pushback(_env *LnsEnv, arg1 *Code_CodeTokenizer)
}
func Lns_cast2Code_CodeCore( obj LnsAny ) LnsAny {
    if _, ok := obj.(Code_CodeCore); ok { 
        return obj
    }
    return nil
}

// declaration Class -- CodeGenerator
type Code_CodeGeneratorMtd interface {
    Output(_env *LnsEnv, arg1 Code_CodeCore)
    OutputElement(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList)
    OutputToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
}
type Code_CodeGenerator struct {
    stream Lns_oStream
    tokenList *LnsList
    prevToken *LnsTypes.Types_Token
    FP Code_CodeGeneratorMtd
}
func Code_CodeGenerator2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeGenerator).FP
}
func Code_CodeGenerator_toSlice(slice []LnsAny) []*Code_CodeGenerator {
    ret := make([]*Code_CodeGenerator, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeGeneratorDownCast).ToCode_CodeGenerator()
    }
    return ret
}
type Code_CodeGeneratorDownCast interface {
    ToCode_CodeGenerator() *Code_CodeGenerator
}
func Code_CodeGeneratorDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeGeneratorDownCast)
    if ok { return work.ToCode_CodeGenerator() }
    return nil
}
func (obj *Code_CodeGenerator) ToCode_CodeGenerator() *Code_CodeGenerator {
    return obj
}
func NewCode_CodeGenerator(_env *LnsEnv, arg1 Lns_oStream) *Code_CodeGenerator {
    obj := &Code_CodeGenerator{}
    obj.FP = obj
    obj.InitCode_CodeGenerator(_env, arg1)
    return obj
}
// 114: DeclConstr
func (self *Code_CodeGenerator) InitCode_CodeGenerator(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.tokenList = NewLnsList([]LnsAny{})
    self.prevToken = LnsTypes.Types_noneToken
}


// declaration Class -- CodeCoreStat
type Code_CodeCoreStatMtd interface {
    OutputCode(_env *LnsEnv, arg1 *Code_CodeGenerator)
    Pushback(_env *LnsEnv, arg1 *Code_CodeTokenizer)
}
type Code_CodeCoreStat struct {
    list *LnsList
    FP Code_CodeCoreStatMtd
}
func Code_CodeCoreStat2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeCoreStat).FP
}
func Code_CodeCoreStat_toSlice(slice []LnsAny) []*Code_CodeCoreStat {
    ret := make([]*Code_CodeCoreStat, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeCoreStatDownCast).ToCode_CodeCoreStat()
    }
    return ret
}
type Code_CodeCoreStatDownCast interface {
    ToCode_CodeCoreStat() *Code_CodeCoreStat
}
func Code_CodeCoreStatDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeCoreStatDownCast)
    if ok { return work.ToCode_CodeCoreStat() }
    return nil
}
func (obj *Code_CodeCoreStat) ToCode_CodeCoreStat() *Code_CodeCoreStat {
    return obj
}
func NewCode_CodeCoreStat(_env *LnsEnv, arg1 *LnsList) *Code_CodeCoreStat {
    obj := &Code_CodeCoreStat{}
    obj.FP = obj
    obj.InitCode_CodeCoreStat(_env, arg1)
    return obj
}
func (self *Code_CodeCoreStat) InitCode_CodeCoreStat(_env *LnsEnv, arg1 *LnsList) {
    self.list = arg1
}

// declaration Class -- CodeCoreBuiltin
type Code_CodeCoreBuiltinMtd interface {
    OutputCode(_env *LnsEnv, arg1 *Code_CodeGenerator)
    Pushback(_env *LnsEnv, arg1 *Code_CodeTokenizer)
}
type Code_CodeCoreBuiltin struct {
    token *LnsTypes.Types_Token
    FP Code_CodeCoreBuiltinMtd
}
func Code_CodeCoreBuiltin2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeCoreBuiltin).FP
}
func Code_CodeCoreBuiltin_toSlice(slice []LnsAny) []*Code_CodeCoreBuiltin {
    ret := make([]*Code_CodeCoreBuiltin, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeCoreBuiltinDownCast).ToCode_CodeCoreBuiltin()
    }
    return ret
}
type Code_CodeCoreBuiltinDownCast interface {
    ToCode_CodeCoreBuiltin() *Code_CodeCoreBuiltin
}
func Code_CodeCoreBuiltinDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeCoreBuiltinDownCast)
    if ok { return work.ToCode_CodeCoreBuiltin() }
    return nil
}
func (obj *Code_CodeCoreBuiltin) ToCode_CodeCoreBuiltin() *Code_CodeCoreBuiltin {
    return obj
}
func NewCode_CodeCoreBuiltin(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *Code_CodeCoreBuiltin {
    obj := &Code_CodeCoreBuiltin{}
    obj.FP = obj
    obj.InitCode_CodeCoreBuiltin(_env, arg1)
    return obj
}
func (self *Code_CodeCoreBuiltin) InitCode_CodeCoreBuiltin(_env *LnsEnv, arg1 *LnsTypes.Types_Token) {
    self.token = arg1
}

// declaration Class -- CodeCoreToken
type Code_CodeCoreTokenMtd interface {
    OutputCode(_env *LnsEnv, arg1 *Code_CodeGenerator)
    Pushback(_env *LnsEnv, arg1 *Code_CodeTokenizer)
}
type Code_CodeCoreToken struct {
    token *LnsTypes.Types_Token
    FP Code_CodeCoreTokenMtd
}
func Code_CodeCoreToken2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeCoreToken).FP
}
func Code_CodeCoreToken_toSlice(slice []LnsAny) []*Code_CodeCoreToken {
    ret := make([]*Code_CodeCoreToken, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeCoreTokenDownCast).ToCode_CodeCoreToken()
    }
    return ret
}
type Code_CodeCoreTokenDownCast interface {
    ToCode_CodeCoreToken() *Code_CodeCoreToken
}
func Code_CodeCoreTokenDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeCoreTokenDownCast)
    if ok { return work.ToCode_CodeCoreToken() }
    return nil
}
func (obj *Code_CodeCoreToken) ToCode_CodeCoreToken() *Code_CodeCoreToken {
    return obj
}
func NewCode_CodeCoreToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *Code_CodeCoreToken {
    obj := &Code_CodeCoreToken{}
    obj.FP = obj
    obj.InitCode_CodeCoreToken(_env, arg1)
    return obj
}
func (self *Code_CodeCoreToken) InitCode_CodeCoreToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token) {
    self.token = arg1
}

// declaration Class -- CodeCoreList
type Code_CodeCoreListMtd interface {
    OutputCode(_env *LnsEnv, arg1 *Code_CodeGenerator)
    Pushback(_env *LnsEnv, arg1 *Code_CodeTokenizer)
}
type Code_CodeCoreList struct {
    elementName LnsAny
    codeCoreList *LnsList
    FP Code_CodeCoreListMtd
}
func Code_CodeCoreList2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeCoreList).FP
}
func Code_CodeCoreList_toSlice(slice []LnsAny) []*Code_CodeCoreList {
    ret := make([]*Code_CodeCoreList, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeCoreListDownCast).ToCode_CodeCoreList()
    }
    return ret
}
type Code_CodeCoreListDownCast interface {
    ToCode_CodeCoreList() *Code_CodeCoreList
}
func Code_CodeCoreListDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeCoreListDownCast)
    if ok { return work.ToCode_CodeCoreList() }
    return nil
}
func (obj *Code_CodeCoreList) ToCode_CodeCoreList() *Code_CodeCoreList {
    return obj
}
func NewCode_CodeCoreList(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList) *Code_CodeCoreList {
    obj := &Code_CodeCoreList{}
    obj.FP = obj
    obj.InitCode_CodeCoreList(_env, arg1, arg2)
    return obj
}
func (self *Code_CodeCoreList) InitCode_CodeCoreList(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList) {
    self.elementName = arg1
    self.codeCoreList = arg2
}

func Lns_Code_init(_env *LnsEnv) {
    if init_Code { return }
    init_Code = true
    Code__mod__ = "@lns.@Code"
    Lns_InitMod()
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Code = false
}
