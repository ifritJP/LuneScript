// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Rule bool
var Rule__mod__ string
// decl enum -- RuleKind 
type Rule_RuleKind = LnsInt
const Rule_RuleKind__Option = 0
const Rule_RuleKind__Repeat = 1
var Rule_RuleKindList_ = NewLnsList( []LnsAny {
  Rule_RuleKind__Option,
  Rule_RuleKind__Repeat,
})
func Rule_RuleKind_get__allList(_env *LnsEnv) *LnsList{
    return Rule_RuleKindList_
}
var Rule_RuleKindMap_ = map[LnsInt]string {
  Rule_RuleKind__Option: "RuleKind.Option",
  Rule_RuleKind__Repeat: "RuleKind.Repeat",
}
func Rule_RuleKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Rule_RuleKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Rule_RuleKind_getTxt(arg1 LnsInt) string {
    return Rule_RuleKindMap_[arg1];
}

// 36: decl @lns.@Rule.RuleCore.setupCandidate
func (self *Rule_RuleCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    self.candidate.FP.Add(_env, self.rule.SetupCandidate(_env, ctrl, usedCoreSet), false)
    return self.candidate
}
// 42: decl @lns.@Rule.RuleCore.parseCode
func (self *Rule_RuleCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer Code_CodeTokenizerIF,hook Code_ParseCodeHookIF,depth LnsInt) LnsAny {
    __func__ := "@lns.@Rule.RuleCore.parseCode"
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        var result LnsAny
        result = hook.Process(_env, self.rule.ParseCode(_env, ctrl, tokenizer, hook, depth + 1), depth + 1)
        switch _matchExp0 := result.(type) {
        case *Code_ParseCodeRet__Eof:
            if list.Len() == 0{
                return Code_ParseCodeRet__Abbr_Obj
            } else { 
                return &Code_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Code_ParseCodeRet__Unmatch:
            if list.Len() == 0{
                return Code_ParseCodeRet__Abbr_Obj
            } else { 
                return &Code_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Code_ParseCodeRet__Abbr:
            if list.Len() == 0{
                return Code_ParseCodeRet__Abbr_Obj
            } else { 
                return &Code_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Code_ParseCodeRet__Detect:
            codeCore := _matchExp0.Val1
            Util_log(_env, Lns_2DDD(__func__, Rule_RuleKind_getTxt( self.kind)))
            if self.kind == Rule_RuleKind__Option{
                return result
            }
            list.Insert(codeCore)
        }
    }
// insert a dummy
    return nil
}
// 102: decl @lns.@Rule.Rule.setupCandidate
func (self *Rule_Rule) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    for _, _core := range( self.coreList.Items ) {
        core := _core.(Ebnf_Core)
        var coreSet *LnsSet
        coreSet = usedCoreSet.Clone()
        var work *Ebnf_Candidate
        work = core.SetupCandidate(_env, ctrl, coreSet)
        self.candidate.FP.Add(_env, work, true)
        if Lns_op_not(work.FP.Get_canAbbr(_env)){
            break
        }
    }
    return self.candidate
}
// 116: decl @lns.@Rule.Rule.parseCodeSub
func (self *Rule_Rule) parseCodeSub(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer Code_CodeTokenizerIF,hook Code_ParseCodeHookIF,nextDepth LnsInt) LnsAny {
    __func__ := "@lns.@Rule.Rule.parseCodeSub"
    var codeCoreList *LnsList
    codeCoreList = NewLnsList([]LnsAny{})
    var Rule_pushbackToken func(_env *LnsEnv)
    Rule_pushbackToken = func(_env *LnsEnv) {
        {
            var _forFrom0 LnsInt = codeCoreList.Len()
            var _forTo0 LnsInt = 1
            _forWork0 := _forFrom0
            _forDelta0 := -1
            for {
                if _forDelta0 > 0 {
                   if _forWork0 > _forTo0 { break }
                } else {
                   if _forWork0 < _forTo0 { break }
                }
                coreIndex := _forWork0
                var codeCore Code_CodeCore
                codeCore = codeCoreList.GetAt(coreIndex).(Code_CodeCore)
                codeCore.Pushback(_env, tokenizer)
                _forWork0 += _forDelta0
            }
        }
    }
    for _index, _core := range( self.coreList.Items ) {
        index := _index + 1
        core := _core.(Ebnf_Core)
        var token *LnsTypes.Types_Token
        token = tokenizer.PeekToken(_env)
        Util_log(_env, Lns_2DDD(__func__, self.elementName, index, self.coreList.Len(), token.Txt))
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( core.Get_candidate(_env).FP.Has(_env, token)) ||
            _env.SetStackVal( core.Get_candidate(_env).FP.Get_canAbbr(_env)) ).(bool){
            switch _matchExp0 := hook.Process(_env, core.ParseCode(_env, ctrl, tokenizer, hook, nextDepth), nextDepth).(type) {
            case *Code_ParseCodeRet__Eof:
                return Code_ParseCodeRet__Eof_Obj
            case *Code_ParseCodeRet__Unmatch:
                Rule_pushbackToken(_env)
                return Code_ParseCodeRet__Unmatch_Obj
            case *Code_ParseCodeRet__Abbr:
            case *Code_ParseCodeRet__Detect:
                codeCore := _matchExp0.Val1
                codeCoreList.Insert(codeCore)
            }
        } else { 
            Util_log(_env, Lns_2DDD(__func__, self.elementName, "Unmatch -- candidate", index))
            Rule_pushbackToken(_env)
            return Code_ParseCodeRet__Unmatch_Obj
        }
    }
    if codeCoreList.Len() == 0{
        return Code_ParseCodeRet__Abbr_Obj
    }
    return &Code_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, self.elementName, codeCoreList).FP}
}
// 158: decl @lns.@Rule.Rule.parseCode
func (self *Rule_Rule) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer Code_CodeTokenizerIF,hook Code_ParseCodeHookIF,depth LnsInt) LnsAny {
    {
        _elementName := self.elementName
        if !Lns_IsNil( _elementName ) {
            elementName := _elementName.(string)
            hook.Prepare(_env, elementName, depth + 1, tokenizer.PeekToken(_env))
        }
    }
    return hook.Process(_env, self.FP.parseCodeSub(_env, ctrl, tokenizer, hook, depth + 2), depth + 1)
}
// 188: decl @lns.@Rule.RuleList.setupCandidate
func (self *Rule_RuleList) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    for _, _rule := range( self.list.Items ) {
        rule := _rule.(Rule_RuleDownCast).ToRule_Rule()
        self.candidate.FP.Add(_env, rule.FP.SetupCandidate(_env, ctrl, usedCoreSet), true)
    }
    return self.candidate
}
// 196: decl @lns.@Rule.RuleList.parseCode
func (self *Rule_RuleList) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer Code_CodeTokenizerIF,hook Code_ParseCodeHookIF,depth LnsInt) LnsAny {
    __func__ := "@lns.@Rule.RuleList.parseCode"
    if Lns_op_not(self.candidate.FP.Has(_env, tokenizer.PeekToken(_env))){
        Util_log(_env, Lns_2DDD(__func__, "not found candidate -- ", tokenizer.PeekToken(_env).Txt, LnsTypes.Types_TokenKind_getTxt( tokenizer.PeekToken(_env).Kind), self.elementName))
        if self.candidate.FP.Get_canAbbr(_env){
            return Code_ParseCodeRet__Abbr_Obj
        }
        return Code_ParseCodeRet__Unmatch_Obj
    }
    var result LnsAny
    result = Code_ParseCodeRet__Unmatch_Obj
    for _index, _rule := range( self.list.Items ) {
        index := _index + 1
        rule := _rule.(Rule_RuleDownCast).ToRule_Rule()
        if rule.FP.Get_candidate(_env).FP.Has(_env, tokenizer.PeekToken(_env)){
            Util_log(_env, Lns_2DDD("check -- ", index, self.list.Len(), self.elementName))
            switch _matchExp0 := rule.FP.ParseCode(_env, ctrl, tokenizer, hook, depth + 1).(type) {
            case *Code_ParseCodeRet__Eof:
                return Code_ParseCodeRet__Eof_Obj
            case *Code_ParseCodeRet__Unmatch:
            case *Code_ParseCodeRet__Abbr:
                result = Code_ParseCodeRet__Abbr_Obj
            case *Code_ParseCodeRet__Detect:
                codeCore := _matchExp0.Val1
                Util_log(_env, Lns_2DDD("detect -- ", self.elementName))
                return &Code_ParseCodeRet__Detect{codeCore}
            }
        }
    }
    Util_log(_env, Lns_2DDD(result.(LnsAlgeVal).GetTxt(), " -- ", self.elementName))
    return result
}
// declaration Class -- RuleCore
type Rule_RuleCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 Code_CodeTokenizerIF, arg3 Code_ParseCodeHookIF, arg4 LnsInt) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Rule_RuleCore struct {
    kind LnsInt
    rule Ebnf_RuleIF
    candidate *Ebnf_Candidate
    FP Rule_RuleCoreMtd
}
func Rule_RuleCore2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Rule_RuleCore).FP
}
func Rule_RuleCore_toSlice(slice []LnsAny) []*Rule_RuleCore {
    ret := make([]*Rule_RuleCore, len(slice))
    for index, val := range slice {
        ret[index] = val.(Rule_RuleCoreDownCast).ToRule_RuleCore()
    }
    return ret
}
type Rule_RuleCoreDownCast interface {
    ToRule_RuleCore() *Rule_RuleCore
}
func Rule_RuleCoreDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Rule_RuleCoreDownCast)
    if ok { return work.ToRule_RuleCore() }
    return nil
}
func (obj *Rule_RuleCore) ToRule_RuleCore() *Rule_RuleCore {
    return obj
}
func NewRule_RuleCore(_env *LnsEnv, arg1 LnsInt, arg2 Ebnf_RuleIF) *Rule_RuleCore {
    obj := &Rule_RuleCore{}
    obj.FP = obj
    obj.InitRule_RuleCore(_env, arg1, arg2)
    return obj
}
func (self *Rule_RuleCore) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 30: DeclConstr
func (self *Rule_RuleCore) InitRule_RuleCore(_env *LnsEnv, kind LnsInt,rule Ebnf_RuleIF) {
    self.kind = kind
    self.rule = rule
    self.candidate = NewEbnf_Candidate(_env, true, false)
}


// declaration Class -- Rule
type Rule_RuleMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 Code_CodeTokenizerIF, arg3 Code_ParseCodeHookIF, arg4 LnsInt) LnsAny
    parseCodeSub(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 Code_CodeTokenizerIF, arg3 Code_ParseCodeHookIF, arg4 LnsInt) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Rule_Rule struct {
    elementName LnsAny
    coreList *LnsList
    candidate *Ebnf_Candidate
    FP Rule_RuleMtd
}
func Rule_Rule2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Rule_Rule).FP
}
func Rule_Rule_toSlice(slice []LnsAny) []*Rule_Rule {
    ret := make([]*Rule_Rule, len(slice))
    for index, val := range slice {
        ret[index] = val.(Rule_RuleDownCast).ToRule_Rule()
    }
    return ret
}
type Rule_RuleDownCast interface {
    ToRule_Rule() *Rule_Rule
}
func Rule_RuleDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Rule_RuleDownCast)
    if ok { return work.ToRule_Rule() }
    return nil
}
func (obj *Rule_Rule) ToRule_Rule() *Rule_Rule {
    return obj
}
func NewRule_Rule(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList) *Rule_Rule {
    obj := &Rule_Rule{}
    obj.FP = obj
    obj.InitRule_Rule(_env, arg1, arg2)
    return obj
}
func (self *Rule_Rule) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 96: DeclConstr
func (self *Rule_Rule) InitRule_Rule(_env *LnsEnv, elementName LnsAny,coreList *LnsList) {
    self.elementName = elementName
    self.coreList = coreList
    self.candidate = NewEbnf_Candidate(_env, true, true)
}


// declaration Class -- RuleList
type Rule_RuleListMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 Code_CodeTokenizerIF, arg3 Code_ParseCodeHookIF, arg4 LnsInt) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Rule_RuleList struct {
    elementName string
    list *LnsList
    candidate *Ebnf_Candidate
    FP Rule_RuleListMtd
}
func Rule_RuleList2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Rule_RuleList).FP
}
func Rule_RuleList_toSlice(slice []LnsAny) []*Rule_RuleList {
    ret := make([]*Rule_RuleList, len(slice))
    for index, val := range slice {
        ret[index] = val.(Rule_RuleListDownCast).ToRule_RuleList()
    }
    return ret
}
type Rule_RuleListDownCast interface {
    ToRule_RuleList() *Rule_RuleList
}
func Rule_RuleListDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Rule_RuleListDownCast)
    if ok { return work.ToRule_RuleList() }
    return nil
}
func (obj *Rule_RuleList) ToRule_RuleList() *Rule_RuleList {
    return obj
}
func NewRule_RuleList(_env *LnsEnv, arg1 string, arg2 *LnsList) *Rule_RuleList {
    obj := &Rule_RuleList{}
    obj.FP = obj
    obj.InitRule_RuleList(_env, arg1, arg2)
    return obj
}
func (self *Rule_RuleList) Get_candidate(_env *LnsEnv) *Ebnf_Candidate{ return self.candidate }
// 182: DeclConstr
func (self *Rule_RuleList) InitRule_RuleList(_env *LnsEnv, elementName string,list *LnsList) {
    self.elementName = elementName
    self.list = list
    self.candidate = NewEbnf_Candidate(_env, true, true)
}


func Lns_Rule_init(_env *LnsEnv) {
    if init_Rule { return }
    init_Rule = true
    Rule__mod__ = "@lns.@Rule"
    Lns_InitMod()
    Lns_Ebnf_init(_env)
    Lns_Code_init(_env)
    Lns_Util_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Rule = false
}
