// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Code bool
var Code__mod__ string
// decl alge -- ParseCodeRet
type Code_ParseCodeRet = LnsAny
type Code_ParseCodeRet__Abbr struct{
}
var Code_ParseCodeRet__Abbr_Obj = &Code_ParseCodeRet__Abbr{}
func (self *Code_ParseCodeRet__Abbr) GetTxt() string {
return "ParseCodeRet.Abbr"
}
type Code_ParseCodeRet__Detect struct{
Val1 Code_CodeCore
}
func (self *Code_ParseCodeRet__Detect) GetTxt() string {
return "ParseCodeRet.Detect"
}
type Code_ParseCodeRet__Eof struct{
}
var Code_ParseCodeRet__Eof_Obj = &Code_ParseCodeRet__Eof{}
func (self *Code_ParseCodeRet__Eof) GetTxt() string {
return "ParseCodeRet.Eof"
}
type Code_ParseCodeRet__Unmatch struct{
}
var Code_ParseCodeRet__Unmatch_Obj = &Code_ParseCodeRet__Unmatch{}
func (self *Code_ParseCodeRet__Unmatch) GetTxt() string {
return "ParseCodeRet.Unmatch"
}

// 384: decl @lns.@Code.outputCode
func Code_outputCode(_env *LnsEnv, codeCore Code_CodeCore,stream *Code_CodeGenerator) {
    codeCore.OutputCode(_env, stream)
    stream.FP.Writeln(_env)
}

// 36: decl @lns.@Code.CodeTokenizer.getToken
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
// 79: decl @lns.@Code.CodeTokenizer.pushback
func (self *Code_CodeTokenizer) Pushback(_env *LnsEnv) {
    __func__ := "@lns.@Code.CodeTokenizer.pushback"
    Util_log(_env, Lns_2DDD(__func__, self.tokenList.Len(), self.tokenIndex))
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
// 151: decl @lns.@Code.CodeGenerator.pushElementIndent
func (self *Code_CodeGenerator) pushElementIndent(_env *LnsEnv) {
    self.elementIndent.Insert(Code_ElementIndentInfo2Stem(NewCode_ElementIndentInfo(_env, self.colmun + 1, false)))
}
// 154: decl @lns.@Code.CodeGenerator.popElementIndent
func (self *Code_CodeGenerator) popElementIndent(_env *LnsEnv) {
    self.elementIndent.Remove(nil)
}
// 157: decl @lns.@Code.CodeGenerator.peekElementIndent
func (self *Code_CodeGenerator) peekElementIndent(_env *LnsEnv) *Code_ElementIndentInfo {
    return self.elementIndent.GetAt(self.elementIndent.Len()).(Code_ElementIndentInfoDownCast).ToCode_ElementIndentInfo()
}
// 160: decl @lns.@Code.CodeGenerator.peekElementIndentRO
func (self *Code_CodeGenerator) peekElementIndentRO(_env *LnsEnv) *Code_ElementIndentInfo {
    return self.elementIndent.GetAt(self.elementIndent.Len()).(Code_ElementIndentInfoDownCast).ToCode_ElementIndentInfo()
}
// 165: decl @lns.@Code.CodeGenerator.curIndent
func (self *Code_CodeGenerator) curIndent(_env *LnsEnv, term bool) LnsInt {
    var indent LnsInt
    var elementIndent *Code_ElementIndentInfo
    elementIndent = self.FP.peekElementIndentRO(_env)
    if elementIndent.FP.Get_hasToken(_env){
        indent = elementIndent.FP.Get_indent(_env)
    } else { 
        indent = 0
    }
    var work LnsInt
    work = self.indentStack.GetAt(self.indentStack.Len()).(Code_CodeIndentInfoDownCast).ToCode_CodeIndentInfo().FP.Get_indent(_env)
    if term{
        return work
    }
    if indent > work{
        return indent
    }
    return work
}
// 182: decl @lns.@Code.CodeGenerator.isIndentEnd
func (self *Code_CodeGenerator) isIndentEnd(_env *LnsEnv, token *LnsTypes.Types_Token) bool {
    var endTxt string
    endTxt = self.indentStack.GetAt(self.indentStack.Len()).(Code_CodeIndentInfoDownCast).ToCode_CodeIndentInfo().FP.Get_endTxt(_env)
    return token.Txt == endTxt
}
// 186: decl @lns.@Code.CodeGenerator.popIndent
func (self *Code_CodeGenerator) popIndent(_env *LnsEnv) {
    __func__ := "@lns.@Code.CodeGenerator.popIndent"
    var indentInfo *Code_CodeIndentInfo
    indentInfo = self.indentStack.GetAt(self.indentStack.Len()).(Code_CodeIndentInfoDownCast).ToCode_CodeIndentInfo()
    Util_log(_env, Lns_2DDD(__func__, indentInfo.FP.Get_endTxt(_env)))
    self.indentStack.Remove(nil)
}
// 191: decl @lns.@Code.CodeGenerator.pushIndent
func (self *Code_CodeGenerator) pushIndent(_env *LnsEnv, endTxt string,pos LnsInt) {
    __func__ := "@lns.@Code.CodeGenerator.pushIndent"
    Util_log(_env, Lns_2DDD(__func__, endTxt, pos))
    self.indentStack.Insert(Code_CodeIndentInfo2Stem(NewCode_CodeIndentInfo(_env, endTxt, pos)))
}
// 196: decl @lns.@Code.CodeGenerator.outputElement
func (self *Code_CodeGenerator) OutputElement(_env *LnsEnv, elementName LnsAny,codeCoreList *LnsList) {
    var pushedElementIndent bool
    if elementName == "<exp>"{
        pushedElementIndent = true
        self.FP.pushElementIndent(_env)
    } else { 
        pushedElementIndent = false
    }
    for _, _core := range( codeCoreList.Items ) {
        core := _core.(Code_CodeCore)
        core.OutputCode(_env, self)
    }
    if pushedElementIndent{
        self.FP.popElementIndent(_env)
    }
}
// 213: decl @lns.@Code.CodeGenerator.processIndentCur
func (self *Code_CodeGenerator) processIndentCur(_env *LnsEnv, token *LnsTypes.Types_Token,newLine bool) LnsInt {
    {
        var indentPos LnsInt
        if newLine{
            indentPos = self.FP.curIndent(_env, false) + self.baseIndent
        } else { 
            indentPos = self.colmun + 1
        }
        {
            _endTxt := self.indentPairMap.Get(self.prevToken.Txt)
            if !Lns_IsNil( _endTxt ) {
                endTxt := _endTxt.(string)
                self.FP.pushIndent(_env, endTxt, indentPos)
            }
        }
    }
    if self.FP.isIndentEnd(_env, token){
        self.FP.popIndent(_env)
        return self.FP.curIndent(_env, true)
    }
    return self.FP.curIndent(_env, false)
}
// 233: decl @lns.@Code.CodeGenerator.writeRaw
func (self *Code_CodeGenerator) writeRaw(_env *LnsEnv, txt string,validToken bool) {
    if validToken{
        self.FP.peekElementIndent(_env).FP.set_hasToken(_env, true)
    }
    self.colmun = self.colmun + len(txt)
    self.stream.Write(_env, txt)
}
// 240: decl @lns.@Code.CodeGenerator.writeRawln
func (self *Code_CodeGenerator) writeRawln(_env *LnsEnv, num LnsInt) {
    self.colmun = 0
    self.stream.Write(_env, "\n")
    if num > 1{
        self.stream.Write(_env, "\n")
    }
    var elementIndent *Code_ElementIndentInfo
    elementIndent = self.FP.peekElementIndent(_env)
    if Lns_op_not(elementIndent.FP.Get_hasToken(_env)){
        elementIndent.FP.set_indent(_env, self.FP.curIndent(_env, false))
    }
}
// 252: decl @lns.@Code.CodeGenerator.writeln
func (self *Code_CodeGenerator) Writeln(_env *LnsEnv) {
    self.FP.writeRawln(_env, 0)
}
// 256: decl @lns.@Code.CodeGenerator.isNewLine
func (self *Code_CodeGenerator) isNewLine(_env *LnsEnv, token *LnsTypes.Types_Token) bool {
    if self.prevToken == LnsTypes.Types_noneToken{
        return false
    }
    var prevLineNo LnsInt
    prevLineNo = self.prevToken.Pos.LineNo
    if _switch0 := self.prevToken.Kind; _switch0 == LnsTypes.Types_TokenKind__Str {
        if LnsInt(self.prevToken.Txt[1-1]) == 96{
            Lns_LockEnvSync( _env, 264, func () {
                {
                    _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(self.prevToken.Txt,"\n")
                    for {
                        _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                        _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                        if Lns_IsNil( _applyPrev0 ) { break }
                        prevLineNo = prevLineNo + 1
                    }
                }
            })
        }
    }
    return prevLineNo != token.Pos.LineNo
}
// 275: decl @lns.@Code.CodeGenerator.outputToken
func (self *Code_CodeGenerator) OutputToken(_env *LnsEnv, token *LnsTypes.Types_Token) {
    __func__ := "@lns.@Code.CodeGenerator.outputToken"
    for _, _comment := range( token.FP.Get_commentList(_env).Items ) {
        comment := _comment
        self.FP.OutputToken(_env, comment)
    }
    if token.Kind == LnsTypes.Types_TokenKind__Eof{
        return 
    }
    if self.FP.isNewLine(_env, token){
        Util_log(_env, Lns_2DDD(__func__, LnsTypes.Types_TokenKind_getTxt( self.prevToken.Kind), self.prevToken.Pos.LineNo, LnsTypes.Types_TokenKind_getTxt( token.Kind), token.Pos.LineNo))
        self.FP.writeRawln(_env, token.Pos.LineNo - self.prevToken.Pos.LineNo)
        self.FP.writeRaw(_env, _env.GetVM().String_rep(" ", self.FP.processIndentCur(_env, token, true)), false)
    } else { 
        self.FP.processIndentCur(_env, token, false)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(token.Consecutive)) &&
            _env.SetStackVal( self.prevToken != LnsTypes.Types_noneToken) ).(bool)){
            self.FP.writeRaw(_env, " ", false)
        }
    }
    if token.Kind == LnsTypes.Types_TokenKind__Char{
        if _switch0 := token.Txt; _switch0 == "'" || _switch0 == "\"" || _switch0 == "\\" {
            self.FP.writeRaw(_env, _env.GetVM().String_format("?\\%s", Lns_2DDD(token.Txt)), true)
        } else {
            self.FP.writeRaw(_env, _env.GetVM().String_format("?%s", Lns_2DDD(token.Txt)), true)
        }
    } else { 
        self.FP.writeRaw(_env, token.Txt, token.Kind != LnsTypes.Types_TokenKind__Cmnt)
    }
    self.prevToken = token
}
// 327: decl @lns.@Code.CodeCoreStat.outputCode
func (self *Code_CodeCoreStat) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    for _, _token := range( self.list.Items ) {
        token := _token.(LnsTypes.Types_TokenDownCast).ToTypes_Token()
        stream.FP.OutputToken(_env, token)
    }
}
// 332: decl @lns.@Code.CodeCoreStat.pushback
func (self *Code_CodeCoreStat) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    for range( self.list.Items ) {
        tokenizer.FP.Pushback(_env)
    }
}
// 344: decl @lns.@Code.CodeCoreBuiltin.outputCode
func (self *Code_CodeCoreBuiltin) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputToken(_env, self.token)
}
// 347: decl @lns.@Code.CodeCoreBuiltin.pushback
func (self *Code_CodeCoreBuiltin) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    tokenizer.FP.Pushback(_env)
}
// 357: decl @lns.@Code.CodeCoreToken.outputCode
func (self *Code_CodeCoreToken) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputToken(_env, self.token)
}
// 360: decl @lns.@Code.CodeCoreToken.pushback
func (self *Code_CodeCoreToken) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    __func__ := "@lns.@Code.CodeCoreToken.pushback"
    Util_log(_env, Lns_2DDD(__func__, self.token.Txt))
    tokenizer.FP.Pushback(_env)
}
// 372: decl @lns.@Code.CodeCoreList.outputCode
func (self *Code_CodeCoreList) OutputCode(_env *LnsEnv, stream *Code_CodeGenerator) {
    stream.FP.OutputElement(_env, self.elementName, self.codeCoreList)
}
// 375: decl @lns.@Code.CodeCoreList.pushback
func (self *Code_CodeCoreList) Pushback(_env *LnsEnv, tokenizer *Code_CodeTokenizer) {
    __func__ := "@lns.@Code.CodeCoreList.pushback"
    Util_log(_env, Lns_2DDD(__func__, self.elementName))
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
// 395: decl @lns.@Code.ParseCodeHook.process
func (self *Code_ParseCodeHook) Process(_env *LnsEnv, parseCodeRet LnsAny,depth LnsInt) LnsAny {
    switch parseCodeRet.(type) {
    case *Code_ParseCodeRet__Eof:
    case *Code_ParseCodeRet__Abbr:
    case *Code_ParseCodeRet__Unmatch:
    case *Code_ParseCodeRet__Detect:
    }
    return parseCodeRet
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
// 24: DeclConstr
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

// declaration Class -- CodeIndentInfo
type Code_CodeIndentInfoMtd interface {
    Get_endTxt(_env *LnsEnv) string
    Get_indent(_env *LnsEnv) LnsInt
}
type Code_CodeIndentInfo struct {
    endTxt string
    indent LnsInt
    FP Code_CodeIndentInfoMtd
}
func Code_CodeIndentInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_CodeIndentInfo).FP
}
func Code_CodeIndentInfo_toSlice(slice []LnsAny) []*Code_CodeIndentInfo {
    ret := make([]*Code_CodeIndentInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_CodeIndentInfoDownCast).ToCode_CodeIndentInfo()
    }
    return ret
}
type Code_CodeIndentInfoDownCast interface {
    ToCode_CodeIndentInfo() *Code_CodeIndentInfo
}
func Code_CodeIndentInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_CodeIndentInfoDownCast)
    if ok { return work.ToCode_CodeIndentInfo() }
    return nil
}
func (obj *Code_CodeIndentInfo) ToCode_CodeIndentInfo() *Code_CodeIndentInfo {
    return obj
}
func NewCode_CodeIndentInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) *Code_CodeIndentInfo {
    obj := &Code_CodeIndentInfo{}
    obj.FP = obj
    obj.InitCode_CodeIndentInfo(_env, arg1, arg2)
    return obj
}
func (self *Code_CodeIndentInfo) InitCode_CodeIndentInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) {
    self.endTxt = arg1
    self.indent = arg2
}
func (self *Code_CodeIndentInfo) Get_endTxt(_env *LnsEnv) string{ return self.endTxt }
func (self *Code_CodeIndentInfo) Get_indent(_env *LnsEnv) LnsInt{ return self.indent }

// declaration Class -- ElementIndentInfo
type Code_ElementIndentInfoMtd interface {
    Get_hasToken(_env *LnsEnv) bool
    Get_indent(_env *LnsEnv) LnsInt
    set_hasToken(_env *LnsEnv, arg1 bool)
    set_indent(_env *LnsEnv, arg1 LnsInt)
}
type Code_ElementIndentInfo struct {
    indent LnsInt
    hasToken bool
    FP Code_ElementIndentInfoMtd
}
func Code_ElementIndentInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_ElementIndentInfo).FP
}
func Code_ElementIndentInfo_toSlice(slice []LnsAny) []*Code_ElementIndentInfo {
    ret := make([]*Code_ElementIndentInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_ElementIndentInfoDownCast).ToCode_ElementIndentInfo()
    }
    return ret
}
type Code_ElementIndentInfoDownCast interface {
    ToCode_ElementIndentInfo() *Code_ElementIndentInfo
}
func Code_ElementIndentInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_ElementIndentInfoDownCast)
    if ok { return work.ToCode_ElementIndentInfo() }
    return nil
}
func (obj *Code_ElementIndentInfo) ToCode_ElementIndentInfo() *Code_ElementIndentInfo {
    return obj
}
func NewCode_ElementIndentInfo(_env *LnsEnv, arg1 LnsInt, arg2 bool) *Code_ElementIndentInfo {
    obj := &Code_ElementIndentInfo{}
    obj.FP = obj
    obj.InitCode_ElementIndentInfo(_env, arg1, arg2)
    return obj
}
func (self *Code_ElementIndentInfo) InitCode_ElementIndentInfo(_env *LnsEnv, arg1 LnsInt, arg2 bool) {
    self.indent = arg1
    self.hasToken = arg2
}
func (self *Code_ElementIndentInfo) Get_indent(_env *LnsEnv) LnsInt{ return self.indent }
func (self *Code_ElementIndentInfo) set_indent(_env *LnsEnv, arg1 LnsInt){ self.indent = arg1 }
func (self *Code_ElementIndentInfo) Get_hasToken(_env *LnsEnv) bool{ return self.hasToken }
func (self *Code_ElementIndentInfo) set_hasToken(_env *LnsEnv, arg1 bool){ self.hasToken = arg1 }

// declaration Class -- CodeGenerator
type Code_CodeGeneratorMtd interface {
    curIndent(_env *LnsEnv, arg1 bool) LnsInt
    isIndentEnd(_env *LnsEnv, arg1 *LnsTypes.Types_Token) bool
    isNewLine(_env *LnsEnv, arg1 *LnsTypes.Types_Token) bool
    OutputElement(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList)
    OutputToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
    peekElementIndent(_env *LnsEnv) *Code_ElementIndentInfo
    peekElementIndentRO(_env *LnsEnv) *Code_ElementIndentInfo
    popElementIndent(_env *LnsEnv)
    popIndent(_env *LnsEnv)
    processIndentCur(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 bool) LnsInt
    pushElementIndent(_env *LnsEnv)
    pushIndent(_env *LnsEnv, arg1 string, arg2 LnsInt)
    writeRaw(_env *LnsEnv, arg1 string, arg2 bool)
    writeRawln(_env *LnsEnv, arg1 LnsInt)
    Writeln(_env *LnsEnv)
}
type Code_CodeGenerator struct {
    stream Lns_oStream
    prevToken *LnsTypes.Types_Token
    indentStack *LnsList
    indentPairMap *LnsMap
    baseIndent LnsInt
    colmun LnsInt
    elementIndent *LnsList
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
// 134: DeclConstr
func (self *Code_CodeGenerator) InitCode_CodeGenerator(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.prevToken = LnsTypes.Types_noneToken
    self.indentStack = NewLnsList(Lns_2DDD(Code_CodeIndentInfo2Stem(NewCode_CodeIndentInfo(_env, "", 0))))
    self.elementIndent = NewLnsList(Lns_2DDD(Code_ElementIndentInfo2Stem(NewCode_ElementIndentInfo(_env, 0, false))))
    self.baseIndent = 3
    self.colmun = 0
    self.indentPairMap = NewLnsMap( map[LnsAny]LnsAny{"{":"}","`{":"}","[":"]","(":")","(@":")","(=":")",})
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

type Code_ParseCodeHookIF interface {
        Process(_env *LnsEnv, arg1 LnsAny, arg2 LnsInt) LnsAny
}
func Lns_cast2Code_ParseCodeHookIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Code_ParseCodeHookIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- ParseCodeHook
type Code_ParseCodeHookMtd interface {
    Process(_env *LnsEnv, arg1 LnsAny, arg2 LnsInt) LnsAny
}
type Code_ParseCodeHook struct {
    FP Code_ParseCodeHookMtd
}
func Code_ParseCodeHook2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Code_ParseCodeHook).FP
}
func Code_ParseCodeHook_toSlice(slice []LnsAny) []*Code_ParseCodeHook {
    ret := make([]*Code_ParseCodeHook, len(slice))
    for index, val := range slice {
        ret[index] = val.(Code_ParseCodeHookDownCast).ToCode_ParseCodeHook()
    }
    return ret
}
type Code_ParseCodeHookDownCast interface {
    ToCode_ParseCodeHook() *Code_ParseCodeHook
}
func Code_ParseCodeHookDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Code_ParseCodeHookDownCast)
    if ok { return work.ToCode_ParseCodeHook() }
    return nil
}
func (obj *Code_ParseCodeHook) ToCode_ParseCodeHook() *Code_ParseCodeHook {
    return obj
}
func NewCode_ParseCodeHook(_env *LnsEnv) *Code_ParseCodeHook {
    obj := &Code_ParseCodeHook{}
    obj.FP = obj
    obj.InitCode_ParseCodeHook(_env)
    return obj
}
func (self *Code_ParseCodeHook) InitCode_ParseCodeHook(_env *LnsEnv) {
}

func Lns_Code_init(_env *LnsEnv) {
    if init_Code { return }
    init_Code = true
    Code__mod__ = "@lns.@Code"
    Lns_InitMod()
    Lns_Util_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Code = false
}
