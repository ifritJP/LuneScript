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
// 10: decl @lns.@Rule.log
func Rule_log_0_(_env *LnsEnv, ddd []LnsAny) {
}

// 31: decl @lns.@Rule.RuleCore.setupCandidate
func (self *Rule_RuleCore) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    self.candidate.FP.Add(_env, self.rule.SetupCandidate(_env, ctrl, usedCoreSet), false)
    return self.candidate
}
// 37: decl @lns.@Rule.RuleCore.parseCode
func (self *Rule_RuleCore) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Rule.RuleCore.parseCode"
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        var result LnsAny
        result = self.rule.ParseCode(_env, ctrl, tokenizer, usedElementSet)
        switch _matchExp0 := result.(type) {
        case *Ebnf_ParseCodeRet__Eof:
            if list.Len() == 0{
                return Ebnf_ParseCodeRet__Abbr_Obj
            } else { 
                return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Ebnf_ParseCodeRet__Unmatch:
            if list.Len() == 0{
                return Ebnf_ParseCodeRet__Abbr_Obj
            } else { 
                return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Ebnf_ParseCodeRet__Abbr:
            if list.Len() == 0{
                return Ebnf_ParseCodeRet__Abbr_Obj
            } else { 
                return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, nil, list).FP}
            }
        case *Ebnf_ParseCodeRet__Detect:
            codeCore := _matchExp0.Val1
            Rule_log_0_(_env, Lns_2DDD(__func__, Rule_RuleKind_getTxt( self.kind)))
            if self.kind == Rule_RuleKind__Option{
                return result
            }
            list.Insert(codeCore)
        }
    }
// insert a dummy
    return nil
}
// 95: decl @lns.@Rule.Rule.setupCandidate
func (self *Rule_Rule) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    var candidate *Ebnf_Candidate
    candidate = NewEbnf_Candidate(_env, true, true)
    for _, _core := range( self.coreList.Items ) {
        core := _core.(Ebnf_Core)
        var coreSet *LnsSet
        coreSet = usedCoreSet.Clone()
        var work *Ebnf_Candidate
        work = core.SetupCandidate(_env, ctrl, coreSet)
        candidate.FP.Add(_env, work, true)
        if Lns_op_not(work.FP.Get_canAbbr(_env)){
            break
        }
    }
    return candidate
}
// 110: decl @lns.@Rule.Rule.parseCode
func (self *Rule_Rule) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Rule.Rule.parseCode"
    var codeCoreList *LnsList
    codeCoreList = NewLnsList([]LnsAny{})
    for _index, _core := range( self.coreList.Items ) {
        index := _index + 1
        core := _core.(Ebnf_Core)
        Rule_log_0_(_env, Lns_2DDD(__func__, index, self.coreList.Len()))
        switch _matchExp0 := core.ParseCode(_env, ctrl, tokenizer, usedElementSet).(type) {
        case *Ebnf_ParseCodeRet__Eof:
            return Ebnf_ParseCodeRet__Eof_Obj
        case *Ebnf_ParseCodeRet__Unmatch:
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
            return Ebnf_ParseCodeRet__Unmatch_Obj
        case *Ebnf_ParseCodeRet__Abbr:
        case *Ebnf_ParseCodeRet__Detect:
            codeCore := _matchExp0.Val1
            codeCoreList.Insert(codeCore)
        }
    }
    if codeCoreList.Len() == 0{
        return Ebnf_ParseCodeRet__Abbr_Obj
    }
    return &Ebnf_ParseCodeRet__Detect{NewCode_CodeCoreList(_env, self.elementName, codeCoreList).FP}
}
// 161: decl @lns.@Rule.RuleList.setupCandidate
func (self *Rule_RuleList) SetupCandidate(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,usedCoreSet *LnsSet) *Ebnf_Candidate {
    for _, _rule := range( self.list.Items ) {
        rule := _rule.(Rule_RuleDownCast).ToRule_Rule()
        self.candidate.FP.Add(_env, rule.FP.SetupCandidate(_env, ctrl, usedCoreSet), true)
    }
    return self.candidate
}
// 169: decl @lns.@Rule.RuleList.parseCode
func (self *Rule_RuleList) ParseCode(_env *LnsEnv, ctrl Ebnf_EbnfCtrlIF,tokenizer *Code_CodeTokenizer,usedElementSet *LnsSet) LnsAny {
    __func__ := "@lns.@Rule.RuleList.parseCode"
    if Lns_op_not(self.candidate.FP.Has(_env, tokenizer.FP.PeekToken(_env))){
        Rule_log_0_(_env, Lns_2DDD(__func__, "not found candidate -- ", tokenizer.FP.PeekToken(_env).Txt, LnsTypes.Types_TokenKind_getTxt( tokenizer.FP.PeekToken(_env).Kind), self.elementName))
        if self.candidate.FP.Get_canAbbr(_env){
            return Ebnf_ParseCodeRet__Abbr_Obj
        }
        return Ebnf_ParseCodeRet__Unmatch_Obj
    }
    var result LnsAny
    result = Ebnf_ParseCodeRet__Unmatch_Obj
    for _index, _rule := range( self.list.Items ) {
        index := _index + 1
        rule := _rule.(Rule_RuleDownCast).ToRule_Rule()
        Rule_log_0_(_env, Lns_2DDD("check -- ", index, self.list.Len(), self.elementName))
        switch _matchExp0 := rule.FP.ParseCode(_env, ctrl, tokenizer, usedElementSet).(type) {
        case *Ebnf_ParseCodeRet__Eof:
            return Ebnf_ParseCodeRet__Eof_Obj
        case *Ebnf_ParseCodeRet__Unmatch:
        case *Ebnf_ParseCodeRet__Abbr:
            result = Ebnf_ParseCodeRet__Abbr_Obj
        case *Ebnf_ParseCodeRet__Detect:
            codeCore := _matchExp0.Val1
            Rule_log_0_(_env, Lns_2DDD("detect -- ", self.elementName))
            return &Ebnf_ParseCodeRet__Detect{codeCore}
        }
    }
    Rule_log_0_(_env, Lns_2DDD(result.(LnsAlgeVal).GetTxt(), " -- ", self.elementName))
    return result
}
// declaration Class -- RuleCore
type Rule_RuleCoreMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
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
// 25: DeclConstr
func (self *Rule_RuleCore) InitRule_RuleCore(_env *LnsEnv, kind LnsInt,rule Ebnf_RuleIF) {
    self.kind = kind
    self.rule = rule
    self.candidate = NewEbnf_Candidate(_env, true, false)
}


// declaration Class -- Rule
type Rule_RuleMtd interface {
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
    SetupCandidate(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *LnsSet) *Ebnf_Candidate
}
type Rule_Rule struct {
    elementName LnsAny
    coreList *LnsList
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
// 90: DeclConstr
func (self *Rule_Rule) InitRule_Rule(_env *LnsEnv, elementName LnsAny,coreList *LnsList) {
    self.elementName = elementName
    self.coreList = coreList
}


// declaration Class -- RuleList
type Rule_RuleListMtd interface {
    Get_candidate(_env *LnsEnv) *Ebnf_Candidate
    ParseCode(_env *LnsEnv, arg1 Ebnf_EbnfCtrlIF, arg2 *Code_CodeTokenizer, arg3 *LnsSet) LnsAny
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
// 155: DeclConstr
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
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Rule = false
}
