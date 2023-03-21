// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Ebnf bool
var Ebnf__mod__ string
// decl alge -- ParseCodeRet
type Ebnf_ParseCodeRet = LnsAny
type Ebnf_ParseCodeRet__Abbr struct{
}
var Ebnf_ParseCodeRet__Abbr_Obj = &Ebnf_ParseCodeRet__Abbr{}
func (self *Ebnf_ParseCodeRet__Abbr) GetTxt() string {
return "ParseCodeRet.Abbr"
}
type Ebnf_ParseCodeRet__Detect struct{
Val1 Code_CodeCore
}
func (self *Ebnf_ParseCodeRet__Detect) GetTxt() string {
return "ParseCodeRet.Detect"
}
type Ebnf_ParseCodeRet__Eof struct{
}
var Ebnf_ParseCodeRet__Eof_Obj = &Ebnf_ParseCodeRet__Eof{}
func (self *Ebnf_ParseCodeRet__Eof) GetTxt() string {
return "ParseCodeRet.Eof"
}
type Ebnf_ParseCodeRet__Unmatch struct{
}
var Ebnf_ParseCodeRet__Unmatch_Obj = &Ebnf_ParseCodeRet__Unmatch{}
func (self *Ebnf_ParseCodeRet__Unmatch) GetTxt() string {
return "ParseCodeRet.Unmatch"
}
// 9: decl @lns.@Ebnf.log
func Ebnf_log_0_(_env *LnsEnv, ddd []LnsAny) {
}

// 609: decl @lns.@Ebnf.isElement
func Ebnf_isElement(_env *LnsEnv, token *LnsTypes.Types_Token) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == LnsTypes.Types_TokenKind__Type) &&
        _env.SetStackVal( LnsInt(token.Txt[1-1]) == 60) ).(bool)
}

// 23: decl @lns.@Ebnf.EbnfTokenizer.create
func Ebnf_EbnfTokenizer_create(_env *LnsEnv) *Ebnf_EbnfTokenizer {
    var path string
    path = "../../../ifritJP.github.io/hugo/content/LuneScript/ebnf.ja.org"
    var tokenizer *LnsTokenizer.Tokenizer_StreamTokenizer
    tokenizer = LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, &LnsTypes.Types_TokenizerSrc__LnsPath{nil, path, "ebnf", nil}, false, nil, nil)
    return NewEbnf_EbnfTokenizer(_env, &tokenizer.Tokenizer_Tokenizer)
}
// 30: decl @lns.@Ebnf.EbnfTokenizer.checkRawNext
func (self *Ebnf_EbnfTokenizer) checkRawNext(_env *LnsEnv, txt string) *LnsTypes.Types_Token {
    var token *LnsTypes.Types_Token
    token = self.tokenizer.GetTokenNoErr(_env, false)
    if token.Txt == txt{
        return token
    }
    Ebnf_log_0_(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d: Illegal token. expects '%s' but '%s'", Lns_2DDD(token.Pos.LineNo, token.Pos.Column, txt, token.Txt))))
    _env.GetVM().OS_exit(1)
// insert a dummy
    return nil
}
// 41: decl @lns.@Ebnf.EbnfTokenizer.getToken
func (self *Ebnf_EbnfTokenizer) GetToken(_env *LnsEnv) *LnsTypes.Types_Token {
    var token *LnsTypes.Types_Token
    token = self.tokenizer.GetTokenNoErr(_env, false)
    if _switch0 := token.Txt; _switch0 == "#" {
        var nextToken *LnsTypes.Types_Token
        nextToken = self.tokenizer.GetTokenNoErr(_env, false)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nextToken.Txt == "+") &&
            _env.SetStackVal( nextToken.Consecutive) ).(bool)){
            return LnsTypes.NewTypes_Token(_env, LnsTypes.Types_TokenKind__Dlmt, "#+", token.Pos, false, nil)
        }
        self.tokenizer.Pushback(_env)
    } else if _switch0 == ":" {
        var nextToken *LnsTypes.Types_Token
        nextToken = self.tokenizer.GetTokenNoErr(_env, false)
        if nextToken.Txt == ":"{
            self.FP.checkRawNext(_env, "=")
            return LnsTypes.NewTypes_Token(_env, LnsTypes.Types_TokenKind__Dlmt, "::=", token.Pos, false, nil)
        } else { 
            self.tokenizer.Pushback(_env)
        }
    } else if _switch0 == "<" {
        if self.inEbnf{
            var nextToken *LnsTypes.Types_Token
            nextToken = self.tokenizer.GetTokenNoErr(_env, false)
            self.FP.checkRawNext(_env, ">")
            return LnsTypes.NewTypes_Token(_env, LnsTypes.Types_TokenKind__Type, _env.GetVM().String_format("<%s>", Lns_2DDD(nextToken.Txt)), token.Pos, false, nil)
        }
    }
    return token
}
// 72: decl @lns.@Ebnf.EbnfTokenizer.checkNext
func (self *Ebnf_EbnfTokenizer) CheckNext(_env *LnsEnv, txt string) *LnsTypes.Types_Token {
    var token *LnsTypes.Types_Token
    token = self.FP.GetToken(_env)
    if token.Txt == txt{
        return token
    }
    Ebnf_log_0_(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d: Illegal token. expects '%s' but '%s'", Lns_2DDD(token.Pos.LineNo, token.Pos.Column, txt, token.Txt))))
    _env.GetVM().OS_exit(1)
// insert a dummy
    return nil
}
// 83: decl @lns.@Ebnf.EbnfTokenizer.pushBack
func (self *Ebnf_EbnfTokenizer) PushBack(_env *LnsEnv, token *LnsTypes.Types_Token) {
    self.tokenizer.PushbackToken(_env, token)
}
// 87: decl @lns.@Ebnf.EbnfTokenizer.skipLine
func (self *Ebnf_EbnfTokenizer) SkipLine(_env *LnsEnv, curToken *LnsTypes.Types_Token) {
    var token *LnsTypes.Types_Token
    token = self.FP.GetToken(_env)
    for token.Pos.LineNo == curToken.Pos.LineNo {
        token = self.FP.GetToken(_env)
    }
    self.FP.PushBack(_env, token)
}
// 255: decl @lns.@Ebnf.Candidate.addTxt
func (self *Ebnf_Candidate) AddTxt(_env *LnsEnv, txt string) {
    self.txtSet.Add(txt)
}
// 258: decl @lns.@Ebnf.Candidate.addTokenKind
func (self *Ebnf_Candidate) AddTokenKind(_env *LnsEnv, kind LnsInt) {
    self.tokenKindSet.Add(kind)
}
// 261: decl @lns.@Ebnf.Candidate.add
func (self *Ebnf_Candidate) Add(_env *LnsEnv, other *Ebnf_Candidate,validAbbr bool) {
    self.txtSet.Or(other.txtSet)
    self.tokenKindSet.Or(other.tokenKindSet)
    if validAbbr{
        self.canAbbr = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.canAbbr) &&
            _env.SetStackVal( other.canAbbr) ).(bool)
    }
    self.any = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.any) &&
        _env.SetStackVal( other.any) ).(bool)
}
// 270: decl @lns.@Ebnf.Candidate.hasTxt
func (self *Ebnf_Candidate) HasTxt(_env *LnsEnv, txt string) bool {
    return self.txtSet.Has(txt)
}
// 273: decl @lns.@Ebnf.Candidate.hasKind
func (self *Ebnf_Candidate) HasKind(_env *LnsEnv, kind LnsInt) bool {
    return self.tokenKindSet.Has(kind)
}
// 276: decl @lns.@Ebnf.Candidate.has
func (self *Ebnf_Candidate) Has(_env *LnsEnv, token *LnsTypes.Types_Token) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.HasTxt(_env, token.Txt)) ||
        _env.SetStackVal( self.FP.HasKind(_env, token.Kind)) ).(bool)
}
// 280: decl @lns.@Ebnf.Candidate.dump
func (self *Ebnf_Candidate) Dump(_env *LnsEnv, stream Lns_oStream) {
    stream.Write(_env, "txt: ")
    for _txt := range( self.txtSet.Items ) {
        txt := _txt.(string)
        stream.Write(_env, _env.GetVM().String_format("%s, ", Lns_2DDD(txt)))
    }
    stream.Write(_env, "\n")
    stream.Write(_env, "kind: ")
    for _kind := range( self.tokenKindSet.Items ) {
        kind := _kind.(LnsInt)
        stream.Write(_env, _env.GetVM().String_format("%s, ", Lns_2DDD(LnsTypes.Types_TokenKind_getTxt( kind))))
    }
    stream.Write(_env, "\n")
    stream.Write(_env, _env.GetVM().String_format("abbr: %s, any: %s\n", Lns_2DDD(self.canAbbr, self.any)))
}
// 395: decl @lns.@Ebnf.TokenCore.setupCandidate
func (self *Ebnf_TokenCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    return self.candidate
}
// 398: decl @lns.@Ebnf.TokenCore.parseCode
func (self *Ebnf_TokenCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Ebnf.TokenCore.parseCode"
    var token *LnsTypes.Types_Token
    token = tokenizer.FP.GetToken(_env)
    Ebnf_log_0_(_env, Lns_2DDD(__func__, self.token.Txt, token.Txt, LnsTypes.Types_TokenKind_getTxt( token.Kind)))
    if token.Txt == self.rawTxt{
        Ebnf_log_0_(_env, Lns_2DDD(__func__, "detect", token.Txt, _env.GetVM().String_format("%d:%d", Lns_2DDD(token.Pos.LineNo, token.Pos.Column))))
        return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreToken(_env, token).FP}
    }
    tokenizer.FP.Pushback(_env)
    if token.Kind == LnsTypes.Types_TokenKind__Eof{
        return Ebnf_ParseCodeRet__Eof_Obj
    }
    return Ebnf_ParseCodeRet__Unmatch_Obj
}
// 431: decl @lns.@Ebnf.BuiltinTokenKindCore.setupCandidate
func (self *Ebnf_BuiltinTokenKindCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    return self.candidate
}
// 434: decl @lns.@Ebnf.BuiltinTokenKindCore.parseCode
func (self *Ebnf_BuiltinTokenKindCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Ebnf.BuiltinTokenKindCore.parseCode"
    var token *LnsTypes.Types_Token
    token = tokenizer.FP.GetToken(_env)
    if self.tokenKindSet.Has(token.Kind){
        Ebnf_log_0_(_env, Lns_2DDD(__func__, "detect", self.elementName))
        return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreBuiltin(_env, token).FP}
    }
    tokenizer.FP.Pushback(_env)
    return Ebnf_ParseCodeRet__Unmatch_Obj
}
// 462: decl @lns.@Ebnf.BuiltinTokenCore.setupCandidate
func (self *Ebnf_BuiltinTokenCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    return self.candidate
}
// 465: decl @lns.@Ebnf.BuiltinTokenCore.parseCode
func (self *Ebnf_BuiltinTokenCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Ebnf.BuiltinTokenCore.parseCode"
    var token *LnsTypes.Types_Token
    token = tokenizer.FP.GetToken(_env)
    if Lns_op_not(self.tokenTxtSet.Has(token.Txt)){
        Ebnf_log_0_(_env, Lns_2DDD(__func__, "detect", self.elementName))
        return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreBuiltin(_env, token).FP}
    }
    tokenizer.FP.Pushback(_env)
    return Ebnf_ParseCodeRet__Unmatch_Obj
}
// 486: decl @lns.@Ebnf.BuiltinStatCore.setupCandidate
func (self *Ebnf_BuiltinStatCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    return self.candidate
}
// 489: decl @lns.@Ebnf.BuiltinStatCore.parseCode
func (self *Ebnf_BuiltinStatCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    var depth LnsInt
    depth = 1
    for  {
        var workToken *LnsTypes.Types_Token
        workToken = tokenizer.FP.GetToken(_env)
        if workToken.Kind == LnsTypes.Types_TokenKind__Eof{
            return Ebnf_ParseCodeRet__Eof_Obj
        }
        if _switch0 := workToken.Txt; _switch0 == "}" {
            depth = depth - 1
            if depth == 0{
                break
            }
        } else if _switch0 == "{" || _switch0 == "`{" {
            depth = depth + 1
        }
        list.Insert(LnsTypes.Types_Token2Stem(workToken))
    }
    tokenizer.FP.Pushback(_env)
    return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreStat(_env, list).FP}
}
// 550: decl @lns.@Ebnf.BuiltinCore.get
func Ebnf_BuiltinCore_get_1_(_env *LnsEnv, elementName string) LnsAny {
    return Ebnf_BuiltinCore__elementNameMap.Get(elementName)
}
// 569: decl @lns.@Ebnf.ElementCore.create
func Ebnf_ElementCore_create(_env *LnsEnv, elementName2Core *LnsMap,elementName string) Ebnf_Core {
    var core Ebnf_Core
    
    {
        _core := elementName2Core.Get(elementName)
        if _core == nil{
            Lns_LockEnvSync( _env, 573, func () {
                {
                    __exp := Ebnf_BuiltinCore_get_1_(_env, elementName)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(Ebnf_Core)
                        core = _exp
                    } else {
                        core = NewEbnf_ElementCore(_env, elementName).FP
                    }
                }
            })
            elementName2Core.Set(elementName,core)
        } else {
            core = _core.(Ebnf_Core)
        }
    }
    return core
}
// 584: decl @lns.@Ebnf.ElementCore.setupCandidate
func (self *Ebnf_ElementCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    __func__ := "@lns.@Ebnf.ElementCore.setupCandidate"
    Lns_print(Lns_2DDD(__func__, self.elementName))
    if Lns_op_not(usedCoreSet.Has(Ebnf_ElementCore2Stem(self))){
        usedCoreSet.Add(Ebnf_ElementCore2Stem(self))
        {
            _ruleList := ctrl.GetRuleList(_env, self.elementName)
            if !Lns_IsNil( _ruleList ) {
                ruleList := _ruleList.(Ebnf_RuleListIF)
                self.candidate.FP.Add(_env, ruleList.SetupCandidate(_env, ctrl, usedCoreSet), true)
            } else {
                panic(_env.GetVM().String_format("unknown -- %s", Lns_2DDD(self.elementName)))
            }
        }
    } else { 
        Ebnf_log_0_(_env, Lns_2DDD("already:", self.elementName))
    }
    return self.candidate
}
// 598: decl @lns.@Ebnf.ElementCore.parseCode
func (self *Ebnf_ElementCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    var ruleList Ebnf_RuleListIF
    
    {
        _ruleList := ctrl.GetRuleList(_env, self.elementName)
        if _ruleList == nil{
            return Ebnf_ParseCodeRet__Unmatch_Obj
        } else {
            ruleList = _ruleList.(Ebnf_RuleListIF)
        }
    }
    return ruleList.ParseCode(_env, ctrl, tokenizer, usedElementSet)
}
// 627: decl @lns.@Ebnf.DeclTokenizer.getToken
func (self *Ebnf_DeclTokenizer) GetToken(_env *LnsEnv) LnsAny {
    var nextToken *LnsTypes.Types_Token
    nextToken = self.tokenizer.FP.GetToken(_env)
    if nextToken.Kind == LnsTypes.Types_TokenKind__Eof{
        return nil
    }
    if nextToken.Pos.LineNo != self.prevToken.Pos.LineNo{
        if nextToken.Txt == "#+"{
            self.tokenizer.FP.PushBack(_env, nextToken)
            return nil
        }
        var checkToken *LnsTypes.Types_Token
        checkToken = self.tokenizer.FP.GetToken(_env)
        self.tokenizer.FP.PushBack(_env, checkToken)
        if checkToken.Txt == "::="{
            self.tokenizer.FP.PushBack(_env, nextToken)
            return nil
        }
    }
    if Ebnf_isElement(_env, nextToken){
        self.allElementSet.Add(nextToken.Txt)
    }
    self.prevToken = nextToken
    return nextToken
}
// 651: decl @lns.@Ebnf.DeclTokenizer.err
func (self *Ebnf_DeclTokenizer) Err(_env *LnsEnv, mess string) {
    Ebnf_log_0_(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d:", Lns_2DDD(self.prevToken.Pos.LineNo, self.prevToken.Pos.Column)), mess))
    _env.GetVM().OS_exit(1)
}
// declaration Class -- EbnfTokenizer
type Ebnf_EbnfTokenizerMtd interface {
    CheckNext(_env *LnsEnv, arg1 string) *LnsTypes.Types_Token
    checkRawNext(_env *LnsEnv, arg1 string) *LnsTypes.Types_Token
    GetToken(_env *LnsEnv) *LnsTypes.Types_Token
    Get_inEbnf(_env *LnsEnv) bool
    PushBack(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
    Set_inEbnf(_env *LnsEnv, arg1 bool)
    SkipLine(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
}
type Ebnf_EbnfTokenizer struct {
    tokenizer LnsTokenizer.Tokenizer_PushbackTokenizer
    inEbnf bool
    FP Ebnf_EbnfTokenizerMtd
}
func Ebnf_EbnfTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_EbnfTokenizer).FP
}
func Ebnf_EbnfTokenizer_toSlice(slice []LnsAny) []*Ebnf_EbnfTokenizer {
    ret := make([]*Ebnf_EbnfTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_EbnfTokenizerDownCast).ToEbnf_EbnfTokenizer()
    }
    return ret
}
type Ebnf_EbnfTokenizerDownCast interface {
    ToEbnf_EbnfTokenizer() *Ebnf_EbnfTokenizer
}
func Ebnf_EbnfTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_EbnfTokenizerDownCast)
    if ok { return work.ToEbnf_EbnfTokenizer() }
    return nil
}
func (obj *Ebnf_EbnfTokenizer) ToEbnf_EbnfTokenizer() *Ebnf_EbnfTokenizer {
    return obj
}
func NewEbnf_EbnfTokenizer(_env *LnsEnv, arg1 *LnsTokenizer.Tokenizer_Tokenizer) *Ebnf_EbnfTokenizer {
    obj := &Ebnf_EbnfTokenizer{}
    obj.FP = obj
    obj.InitEbnf_EbnfTokenizer(_env, arg1)
    return obj
}
func (self *Ebnf_EbnfTokenizer) Get_inEbnf(_env *LnsEnv) bool{ return self.inEbnf }
func (self *Ebnf_EbnfTokenizer) Set_inEbnf(_env *LnsEnv, arg1 bool){ self.inEbnf = arg1 }
// 18: DeclConstr
func (self *Ebnf_EbnfTokenizer) InitEbnf_EbnfTokenizer(_env *LnsEnv, tokenizer *LnsTokenizer.Tokenizer_Tokenizer) {
    self.tokenizer = LnsTokenizer.NewTokenizer_DefaultPushbackTokenizer(_env, tokenizer).FP
    self.inEbnf = false
}


// declaration Class -- Candidate
type Ebnf_CandidateMtd interface {
    Add(_env *LnsEnv, arg1 *Ebnf_Candidate, arg2 bool)
    AddTokenKind(_env *LnsEnv, arg1 LnsInt)
    AddTxt(_env *LnsEnv, arg1 string)
    Dump(_env *LnsEnv, arg1 Lns_oStream)
    Get_any(_env *LnsEnv) bool
    Get_canAbbr(_env *LnsEnv) bool
    Has(_env *LnsEnv, arg1 *LnsTypes.Types_Token) bool
    HasKind(_env *LnsEnv, arg1 LnsInt) bool
    HasTxt(_env *LnsEnv, arg1 string) bool
}
type Ebnf_Candidate struct {
    txtSet *LnsSet
    tokenKindSet *LnsSet
    canAbbr bool
    any bool
    FP Ebnf_CandidateMtd
}
func Ebnf_Candidate2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_Candidate).FP
}
func Ebnf_Candidate_toSlice(slice []LnsAny) []*Ebnf_Candidate {
    ret := make([]*Ebnf_Candidate, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_CandidateDownCast).ToEbnf_Candidate()
    }
    return ret
}
type Ebnf_CandidateDownCast interface {
    ToEbnf_Candidate() *Ebnf_Candidate
}
func Ebnf_CandidateDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_CandidateDownCast)
    if ok { return work.ToEbnf_Candidate() }
    return nil
}
func (obj *Ebnf_Candidate) ToEbnf_Candidate() *Ebnf_Candidate {
    return obj
}
func NewEbnf_Candidate(_env *LnsEnv, arg1 bool, arg2 bool) *Ebnf_Candidate {
    obj := &Ebnf_Candidate{}
    obj.FP = obj
    obj.InitEbnf_Candidate(_env, arg1, arg2)
    return obj
}
func (self *Ebnf_Candidate) Get_canAbbr(_env *LnsEnv) bool{ return self.canAbbr }
func (self *Ebnf_Candidate) Get_any(_env *LnsEnv) bool{ return self.any }
// 248: DeclConstr
func (self *Ebnf_Candidate) InitEbnf_Candidate(_env *LnsEnv, canAbbr bool,any bool) {
    self.canAbbr = canAbbr
    self.txtSet = NewLnsSet([]LnsAny{})
    self.tokenKindSet = NewLnsSet([]LnsAny{})
    self.any = any
}


type Ebnf_RuleListIF interface {
        ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
        SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
func Lns_cast2Ebnf_RuleListIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ebnf_RuleListIF); ok { 
        return obj
    }
    return nil
}

type Ebnf_RuleIF interface {
        ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
        SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
func Lns_cast2Ebnf_RuleIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ebnf_RuleIF); ok { 
        return obj
    }
    return nil
}

type Ebnf_EbnfCtrlIF interface {
        GetRuleList(_env *LnsEnv, arg1 string) LnsAny
}
func Lns_cast2Ebnf_EbnfCtrlIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ebnf_EbnfCtrlIF); ok { 
        return obj
    }
    return nil
}

type Ebnf_Core interface {
        Get_candidate(_env *LnsEnv) *Ebnf_Candidate
        ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
        SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
func Lns_cast2Ebnf_Core( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ebnf_Core); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TokenCore
type Ebnf_TokenCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    Get_token(_env *LnsEnv) *LnsTypes.Types_Token
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Ebnf_TokenCore struct {
    token *LnsTypes.Types_Token
    rawTxt string
    candidate *Ebnf_Candidate
    FP Ebnf_TokenCoreMtd
}
func Ebnf_TokenCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_TokenCore).FP
}
func Ebnf_TokenCore_toSlice(slice []LnsAny) []*Ebnf_TokenCore {
    ret := make([]*Ebnf_TokenCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_TokenCoreDownCast).ToEbnf_TokenCore()
    }
    return ret
}
type Ebnf_TokenCoreDownCast interface {
    ToEbnf_TokenCore() *Ebnf_TokenCore
}
func Ebnf_TokenCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_TokenCoreDownCast)
    if ok { return work.ToEbnf_TokenCore() }
    return nil
}
func (obj *Ebnf_TokenCore) ToEbnf_TokenCore() *Ebnf_TokenCore {
    return obj
}
func NewEbnf_TokenCore(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *Ebnf_TokenCore {
    obj := &Ebnf_TokenCore{}
    obj.FP = obj
    obj.InitEbnf_TokenCore(_env, arg1)
    return obj
}
func (self *Ebnf_TokenCore) Get_token(_env *LnsEnv) *LnsTypes.Types_Token{ return self.token }
func (self *Ebnf_TokenCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 389: DeclConstr
func (self *Ebnf_TokenCore) InitEbnf_TokenCore(_env *LnsEnv, token *LnsTypes.Types_Token) {
    self.token = token
    self.rawTxt = token.FP.GetExcludedDelimitTxt(_env)
    self.candidate = NewEbnf_Candidate(_env, false, false)
    self.candidate.FP.AddTxt(_env, self.rawTxt)
}


// declaration Class -- BuiltinTokenKindCore
type Ebnf_BuiltinTokenKindCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Ebnf_BuiltinTokenKindCore struct {
    elementName string
    tokenKindSet *LnsSet2_[LnsInt]
    candidate *Ebnf_Candidate
    FP Ebnf_BuiltinTokenKindCoreMtd
}
func Ebnf_BuiltinTokenKindCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_BuiltinTokenKindCore).FP
}
func Ebnf_BuiltinTokenKindCore_toSlice(slice []LnsAny) []*Ebnf_BuiltinTokenKindCore {
    ret := make([]*Ebnf_BuiltinTokenKindCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_BuiltinTokenKindCoreDownCast).ToEbnf_BuiltinTokenKindCore()
    }
    return ret
}
type Ebnf_BuiltinTokenKindCoreDownCast interface {
    ToEbnf_BuiltinTokenKindCore() *Ebnf_BuiltinTokenKindCore
}
func Ebnf_BuiltinTokenKindCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_BuiltinTokenKindCoreDownCast)
    if ok { return work.ToEbnf_BuiltinTokenKindCore() }
    return nil
}
func (obj *Ebnf_BuiltinTokenKindCore) ToEbnf_BuiltinTokenKindCore() *Ebnf_BuiltinTokenKindCore {
    return obj
}
func NewEbnf_BuiltinTokenKindCore(_env *LnsEnv, arg1 string, arg2 *LnsSet2_[LnsInt]) *Ebnf_BuiltinTokenKindCore {
    obj := &Ebnf_BuiltinTokenKindCore{}
    obj.FP = obj
    obj.InitEbnf_BuiltinTokenKindCore(_env, arg1, arg2)
    return obj
}
func (self *Ebnf_BuiltinTokenKindCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 422: DeclConstr
func (self *Ebnf_BuiltinTokenKindCore) InitEbnf_BuiltinTokenKindCore(_env *LnsEnv, elementName string,tokenKindSet *LnsSet2_[LnsInt]) {
    self.elementName = elementName
    self.tokenKindSet = tokenKindSet
    self.candidate = NewEbnf_Candidate(_env, false, false)
    for _kind := range( tokenKindSet.Items ) {
        kind := _kind
        self.candidate.FP.AddTokenKind(_env, kind)
    }
}


// declaration Class -- BuiltinTokenCore
type Ebnf_BuiltinTokenCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Ebnf_BuiltinTokenCore struct {
    elementName string
    tokenTxtSet *LnsSet2_[string]
    candidate *Ebnf_Candidate
    FP Ebnf_BuiltinTokenCoreMtd
}
func Ebnf_BuiltinTokenCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_BuiltinTokenCore).FP
}
func Ebnf_BuiltinTokenCore_toSlice(slice []LnsAny) []*Ebnf_BuiltinTokenCore {
    ret := make([]*Ebnf_BuiltinTokenCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_BuiltinTokenCoreDownCast).ToEbnf_BuiltinTokenCore()
    }
    return ret
}
type Ebnf_BuiltinTokenCoreDownCast interface {
    ToEbnf_BuiltinTokenCore() *Ebnf_BuiltinTokenCore
}
func Ebnf_BuiltinTokenCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_BuiltinTokenCoreDownCast)
    if ok { return work.ToEbnf_BuiltinTokenCore() }
    return nil
}
func (obj *Ebnf_BuiltinTokenCore) ToEbnf_BuiltinTokenCore() *Ebnf_BuiltinTokenCore {
    return obj
}
func NewEbnf_BuiltinTokenCore(_env *LnsEnv, arg1 string, arg2 *LnsSet2_[string]) *Ebnf_BuiltinTokenCore {
    obj := &Ebnf_BuiltinTokenCore{}
    obj.FP = obj
    obj.InitEbnf_BuiltinTokenCore(_env, arg1, arg2)
    return obj
}
func (self *Ebnf_BuiltinTokenCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 453: DeclConstr
func (self *Ebnf_BuiltinTokenCore) InitEbnf_BuiltinTokenCore(_env *LnsEnv, elementName string,tokenTxtSet *LnsSet2_[string]) {
    self.elementName = elementName
    self.tokenTxtSet = tokenTxtSet
    self.candidate = NewEbnf_Candidate(_env, false, false)
    for _kind := range( tokenTxtSet.Items ) {
        kind := _kind
        self.candidate.FP.AddTxt(_env, kind)
    }
}


// declaration Class -- BuiltinStatCore
type Ebnf_BuiltinStatCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Ebnf_BuiltinStatCore struct {
    candidate *Ebnf_Candidate
    FP Ebnf_BuiltinStatCoreMtd
}
func Ebnf_BuiltinStatCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_BuiltinStatCore).FP
}
func Ebnf_BuiltinStatCore_toSlice(slice []LnsAny) []*Ebnf_BuiltinStatCore {
    ret := make([]*Ebnf_BuiltinStatCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_BuiltinStatCoreDownCast).ToEbnf_BuiltinStatCore()
    }
    return ret
}
type Ebnf_BuiltinStatCoreDownCast interface {
    ToEbnf_BuiltinStatCore() *Ebnf_BuiltinStatCore
}
func Ebnf_BuiltinStatCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_BuiltinStatCoreDownCast)
    if ok { return work.ToEbnf_BuiltinStatCore() }
    return nil
}
func (obj *Ebnf_BuiltinStatCore) ToEbnf_BuiltinStatCore() *Ebnf_BuiltinStatCore {
    return obj
}
func NewEbnf_BuiltinStatCore(_env *LnsEnv) *Ebnf_BuiltinStatCore {
    obj := &Ebnf_BuiltinStatCore{}
    obj.FP = obj
    obj.InitEbnf_BuiltinStatCore(_env)
    return obj
}
func (self *Ebnf_BuiltinStatCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 482: DeclConstr
func (self *Ebnf_BuiltinStatCore) InitEbnf_BuiltinStatCore(_env *LnsEnv) {
    self.candidate = NewEbnf_Candidate(_env, false, true)
}


// declaration Class -- BuiltinCore
var Ebnf_BuiltinCore__elementNameMap *LnsMap
// 519: decl @lns.@Ebnf.BuiltinCore.___init
func Ebnf_BuiltinCore____init_0_(_env *LnsEnv) {
    Ebnf_BuiltinCore__elementNameMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        var _map *LnsMap2_[string,*LnsSet2_[LnsInt]]
        _map = NewLnsMap2_[string,*LnsSet2_[LnsInt]]( map[string]*LnsSet2_[LnsInt]{"<shebang>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Sheb)),"<sym>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Symb, LnsTypes.Types_TokenKind__Type, LnsTypes.Types_TokenKind__Kywd, LnsTypes.Types_TokenKind__Ope)),"<literal_str>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Str)),"<literal_int>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Int)),"<literal_real>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Real)),"<literal_char>":NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](LnsTypes.Types_TokenKind__Char)),})
        for _elementName, _tokenKindSet := range( _map.Items ) {
            elementName := _elementName
            tokenKindSet := _tokenKindSet
            Ebnf_BuiltinCore__elementNameMap.Set(elementName,NewEbnf_BuiltinTokenKindCore(_env, elementName, tokenKindSet).FP)
        }
    }
    {
        var _map *LnsMap2_[string,*LnsSet2_[string]]
        _map = NewLnsMap2_[string,*LnsSet2_[string]]( map[string]*LnsSet2_[string]{"<token>":NewLnsSet2_[string](Lns_2DDDGen[string](";")),})
        for _elementName, _tokenTxtSet := range( _map.Items ) {
            elementName := _elementName
            tokenTxtSet := _tokenTxtSet
            Ebnf_BuiltinCore__elementNameMap.Set("<token>",NewEbnf_BuiltinTokenCore(_env, elementName, tokenTxtSet).FP)
        }
    }
    Ebnf_BuiltinCore__elementNameMap.Set("<stat>",NewEbnf_BuiltinStatCore(_env).FP)
}

type Ebnf_BuiltinCoreMtd interface {
}
type Ebnf_BuiltinCore struct {
    elementNameMap *LnsMap
    FP Ebnf_BuiltinCoreMtd
}
func Ebnf_BuiltinCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_BuiltinCore).FP
}
func Ebnf_BuiltinCore_toSlice(slice []LnsAny) []*Ebnf_BuiltinCore {
    ret := make([]*Ebnf_BuiltinCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_BuiltinCoreDownCast).ToEbnf_BuiltinCore()
    }
    return ret
}
type Ebnf_BuiltinCoreDownCast interface {
    ToEbnf_BuiltinCore() *Ebnf_BuiltinCore
}
func Ebnf_BuiltinCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_BuiltinCoreDownCast)
    if ok { return work.ToEbnf_BuiltinCore() }
    return nil
}
func (obj *Ebnf_BuiltinCore) ToEbnf_BuiltinCore() *Ebnf_BuiltinCore {
    return obj
}
func NewEbnf_BuiltinCore(_env *LnsEnv) *Ebnf_BuiltinCore {
    obj := &Ebnf_BuiltinCore{}
    obj.FP = obj
    obj.InitEbnf_BuiltinCore(_env)
    return obj
}
func (self *Ebnf_BuiltinCore) InitEbnf_BuiltinCore(_env *LnsEnv) {
}

// declaration Class -- ElementCore
type Ebnf_ElementCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Ebnf_ElementCore struct {
    elementName string
    candidate *Ebnf_Candidate
    FP Ebnf_ElementCoreMtd
}
func Ebnf_ElementCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_ElementCore).FP
}
func Ebnf_ElementCore_toSlice(slice []LnsAny) []*Ebnf_ElementCore {
    ret := make([]*Ebnf_ElementCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_ElementCoreDownCast).ToEbnf_ElementCore()
    }
    return ret
}
type Ebnf_ElementCoreDownCast interface {
    ToEbnf_ElementCore() *Ebnf_ElementCore
}
func Ebnf_ElementCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_ElementCoreDownCast)
    if ok { return work.ToEbnf_ElementCore() }
    return nil
}
func (obj *Ebnf_ElementCore) ToEbnf_ElementCore() *Ebnf_ElementCore {
    return obj
}
func NewEbnf_ElementCore(_env *LnsEnv, arg1 string) *Ebnf_ElementCore {
    obj := &Ebnf_ElementCore{}
    obj.FP = obj
    obj.InitEbnf_ElementCore(_env, arg1)
    return obj
}
func (self *Ebnf_ElementCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 565: DeclConstr
func (self *Ebnf_ElementCore) InitEbnf_ElementCore(_env *LnsEnv, elementName string) {
    self.elementName = elementName
    self.candidate = NewEbnf_Candidate(_env, true, false)
}


// declaration Class -- DeclTokenizer
type Ebnf_DeclTokenizerMtd interface {
    Err(_env *LnsEnv, arg1 string)
    GetToken(_env *LnsEnv) LnsAny
}
type Ebnf_DeclTokenizer struct {
    tokenizer *Ebnf_EbnfTokenizer
    prevToken *LnsTypes.Types_Token
    allElementSet *LnsSet
    FP Ebnf_DeclTokenizerMtd
}
func Ebnf_DeclTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ebnf_DeclTokenizer).FP
}
func Ebnf_DeclTokenizer_toSlice(slice []LnsAny) []*Ebnf_DeclTokenizer {
    ret := make([]*Ebnf_DeclTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Ebnf_DeclTokenizerDownCast).ToEbnf_DeclTokenizer()
    }
    return ret
}
type Ebnf_DeclTokenizerDownCast interface {
    ToEbnf_DeclTokenizer() *Ebnf_DeclTokenizer
}
func Ebnf_DeclTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ebnf_DeclTokenizerDownCast)
    if ok { return work.ToEbnf_DeclTokenizer() }
    return nil
}
func (obj *Ebnf_DeclTokenizer) ToEbnf_DeclTokenizer() *Ebnf_DeclTokenizer {
    return obj
}
func NewEbnf_DeclTokenizer(_env *LnsEnv, arg1 *Ebnf_EbnfTokenizer, arg2 *LnsTypes.Types_Token, arg3 *LnsSet) *Ebnf_DeclTokenizer {
    obj := &Ebnf_DeclTokenizer{}
    obj.FP = obj
    obj.InitEbnf_DeclTokenizer(_env, arg1, arg2, arg3)
    return obj
}
// 619: DeclConstr
func (self *Ebnf_DeclTokenizer) InitEbnf_DeclTokenizer(_env *LnsEnv, tokenizer *Ebnf_EbnfTokenizer,prevToken *LnsTypes.Types_Token,allElementSet *LnsSet) {
    self.tokenizer = tokenizer
    self.prevToken = prevToken
    self.allElementSet = allElementSet
}


func Lns_Ebnf_init(_env *LnsEnv) {
    if init_Ebnf { return }
    init_Ebnf = true
    Ebnf__mod__ = "@lns.@Ebnf"
    Lns_InitMod()
    Lns_Code_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
    Ebnf_BuiltinCore____init_0_(_env)
}
func init() {
    init_Ebnf = false
}
