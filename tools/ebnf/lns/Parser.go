// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_Parser bool
var Parser__mod__ string
// for 92
func Parser_convExp0_443(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// 11: decl @lns.@Parser.log
func Parser_log_0_(_env *LnsEnv, ddd []LnsAny) {
}

// 153: decl @lns.@Parser.analyze_ebnf
func Parser_analyze_ebnf(_env *LnsEnv, tokenizer *Ebnf_EbnfTokenizer) *Parser_EbnfCtrl {
    var ebnfCtrl *Parser_EbnfCtrl
    ebnfCtrl = NewParser_EbnfCtrl(_env)
    for  {
        var token *LnsTypes.Types_Token
        token = tokenizer.FP.GetToken(_env)
        if token.Kind == LnsTypes.Types_TokenKind__Eof{
            break
        }
        if token.Txt == "#"{
            tokenizer.FP.SkipLine(_env, token)
        } else if token.Txt == "#+"{
            var nextToken *LnsTypes.Types_Token
            nextToken = tokenizer.FP.GetToken(_env)
            if _switch0 := _env.GetVM().String_upper(nextToken.Txt); _switch0 == "BEGIN_SRC" {
                tokenizer.FP.Set_inEbnf(_env, true)
            } else if _switch0 == "END_SRC" {
                tokenizer.FP.Set_inEbnf(_env, false)
            }
            tokenizer.FP.SkipLine(_env, token)
        } else if tokenizer.FP.Get_inEbnf(_env){
            if Ebnf_isElement(_env, token){
                ebnfCtrl.FP.processDecl(_env, tokenizer, token)
            }
        }
    }
    ebnfCtrl.FP.SetupCandidate(_env)
    return ebnfCtrl
}

// 28: decl @lns.@Parser.EbnfCtrl.getRuleList
func (self *Parser_EbnfCtrl) GetRuleList(_env *LnsEnv, name string) LnsAny {
    return self.element2ruleList.Get(name)
}
// 34: decl @lns.@Parser.EbnfCtrl.processRule
func (self *Parser_EbnfCtrl) processRule(_env *LnsEnv, elementName LnsAny,tokenizer *Ebnf_DeclTokenizer,termTxt LnsAny)(LnsAny, string) {
    var coreList *LnsList
    coreList = NewLnsList([]LnsAny{})
    var endsTerm bool
    endsTerm = false
    for  {
        var token *LnsTypes.Types_Token
        
        {
            _token := tokenizer.FP.GetToken(_env)
            if _token == nil{
                break
            } else {
                token = _token.(*LnsTypes.Types_Token)
            }
        }
        if _switch0 := token.Txt; _switch0 == "[" {
            var _cond1 *Rule_Rule
            {
                __condWork1, __condWork2 := self.FP.processRule(_env, nil, tokenizer, "]")
                if __condWork1 == nil { return nil, __condWork2 }
                _cond1 = __condWork1.(*Rule_Rule)
            }
            var optRule *Rule_Rule
            optRule = _cond1
            coreList.Insert(Rule_RuleCore2Stem(NewRule_RuleCore(_env, Rule_RuleKind__Option, optRule.FP)))
        } else if _switch0 == "{" {
            var _cond2 *Rule_Rule
            {
                __condWork1, __condWork2 := self.FP.processRule(_env, nil, tokenizer, "}")
                if __condWork1 == nil { return nil, __condWork2 }
                _cond2 = __condWork1.(*Rule_Rule)
            }
            var optRule *Rule_Rule
            optRule = _cond2
            coreList.Insert(Rule_RuleCore2Stem(NewRule_RuleCore(_env, Rule_RuleKind__Repeat, optRule.FP)))
        } else if _switch0 == "|" {
            break
        } else if _switch0 == termTxt {
            endsTerm = true
            break
        } else {
            if Ebnf_isElement(_env, token){
                coreList.Insert(Ebnf_ElementCore_create(_env, self.elementName2Core, token.Txt))
            } else { 
                coreList.Insert(Ebnf_TokenCore2Stem(NewEbnf_TokenCore(_env, token)))
            }
        }
    }
    if coreList.Len() == 0{
        return nil, "coreList is 0"
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( termTxt) &&
        _env.SetStackVal( Lns_op_not(endsTerm)) )){
        return nil, _env.GetVM().String_format("not end term -- %s", Lns_2DDD(termTxt))
    }
    return NewRule_Rule(_env, elementName, coreList), ""
}
// 79: decl @lns.@Parser.EbnfCtrl.processDecl
func (self *Parser_EbnfCtrl) processDecl(_env *LnsEnv, tokenizer *Ebnf_EbnfTokenizer,symbolToken *LnsTypes.Types_Token) *Rule_RuleList {
    __func__ := "@lns.@Parser.EbnfCtrl.processDecl"
    var elementName string
    elementName = symbolToken.Txt
    Parser_log_0_(_env, Lns_2DDD(__func__, elementName))
    tokenizer.FP.CheckNext(_env, "::=")
    self.allElementSet.Add(elementName)
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        var declTokenizer *Ebnf_DeclTokenizer
        declTokenizer = NewEbnf_DeclTokenizer(_env, tokenizer, symbolToken, self.allElementSet)
        var rule LnsAny
        var mess string
        rule,mess = self.FP.processRule(_env, elementName, declTokenizer, nil)
        if rule != nil{
            rule_77 := rule.(*Rule_Rule)
            list.Insert(Rule_Rule2Stem(rule_77))
        } else {
            if list.Len() == 0{
                declTokenizer.FP.Err(_env, mess)
            }
            break
        }
    }
    var ruleList *Rule_RuleList
    ruleList = NewRule_RuleList(_env, elementName, list)
    self.element2ruleList.Set(elementName,ruleList)
    return ruleList
}
// 107: decl @lns.@Parser.EbnfCtrl.dump
func (self *Parser_EbnfCtrl) Dump(_env *LnsEnv) {
    Parser_log_0_(_env, Lns_2DDD("================="))
    {
        __forsortCollection0 := self.allElementSet
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _element := range( __forsortSorted0.Items ) {
            element := _element.(string)
            var ruleList LnsAny
            ruleList = self.element2ruleList.Get(element)
            Parser_log_0_(_env, Lns_2DDD(Rule_RuleList2Stem(ruleList) != nil, element))
            if ruleList != nil{
                ruleList_86 := ruleList.(*Rule_RuleList)
                ruleList_86.FP.Get_candidate(_env).FP.Dump(_env, Lns_io_stdout)
            }
        }
    }
}
// 118: decl @lns.@Parser.EbnfCtrl.parse
func (self *Parser_EbnfCtrl) Parse(_env *LnsEnv, elementName string,tokenizer *Code_CodeTokenizer) {
    var ruleList *Rule_RuleList
    
    {
        _ruleList := self.element2ruleList.Get(elementName)
        if _ruleList == nil{
            Parser_log_0_(_env, Lns_2DDD("error parse"))
            return 
        } else {
            ruleList = _ruleList.(*Rule_RuleList)
        }
    }
    switch _matchExp0 := ruleList.FP.ParseCode(_env, self.FP, tokenizer, NewLnsSet([]LnsAny{})).(type) {
    case *Ebnf_ParseCodeRet__Eof:
        Parser_log_0_(_env, Lns_2DDD("eof"))
    case *Ebnf_ParseCodeRet__Unmatch:
        Parser_log_0_(_env, Lns_2DDD("unmatch"))
    case *Ebnf_ParseCodeRet__Abbr:
        var pos LnsTypes.Types_Position
        pos = tokenizer.FP.GetToken(_env).Pos
        Parser_log_0_(_env, Lns_2DDD("illegal", pos.LineNo, pos.Column))
    case *Ebnf_ParseCodeRet__Detect:
        codeCore := _matchExp0.Val1
        Parser_log_0_(_env, Lns_2DDD("outputCode"))
        codeCore.OutputCode(_env, NewCode_CodeGenerator(_env, Lns_io_stdout))
        Lns_io_stdout.Flush(_env)
        Parser_log_0_(_env, Lns_2DDD(""))
        var token *LnsTypes.Types_Token
        token = tokenizer.FP.GetToken(_env)
        var pos LnsTypes.Types_Position
        pos = token.Pos
        Parser_log_0_(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d", Lns_2DDD(pos.LineNo, pos.Column)), LnsTypes.Types_TokenKind_getTxt( token.Kind)))
    }
}
// 146: decl @lns.@Parser.EbnfCtrl.setupCandidate
func (self *Parser_EbnfCtrl) SetupCandidate(_env *LnsEnv) {
    var fusedCoreSet *LnsSet
    fusedCoreSet = NewLnsSet([]LnsAny{})
    for _, _ruleList := range( self.element2ruleList.Items ) {
        ruleList := _ruleList.(Rule_RuleListDownCast).ToRule_RuleList()
        ruleList.FP.SetupCandidate(_env, self.FP, fusedCoreSet)
    }
}
// declaration Class -- EbnfCtrl
type Parser_EbnfCtrlMtd interface {
    Dump(_env *LnsEnv)
    GetRuleList(_env *LnsEnv, arg1 string) LnsAny
    Parse(_env *LnsEnv, arg1 string, arg2 *Code_CodeTokenizer)
    processDecl(_env *LnsEnv, arg1 *Ebnf_EbnfTokenizer, arg2 *LnsTypes.Types_Token) *Rule_RuleList
    processRule(_env *LnsEnv, arg1 LnsAny, arg2 *Ebnf_DeclTokenizer, arg3 LnsAny)(LnsAny, string)
    SetupCandidate(_env *LnsEnv)
}
type Parser_EbnfCtrl struct {
    element2ruleList *LnsMap
    allElementSet *LnsSet
    elementName2Core *LnsMap
    FP Parser_EbnfCtrlMtd
}
func Parser_EbnfCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_EbnfCtrl).FP
}
func Parser_EbnfCtrl_toSlice(slice []LnsAny) []*Parser_EbnfCtrl {
    ret := make([]*Parser_EbnfCtrl, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_EbnfCtrlDownCast).ToParser_EbnfCtrl()
    }
    return ret
}
type Parser_EbnfCtrlDownCast interface {
    ToParser_EbnfCtrl() *Parser_EbnfCtrl
}
func Parser_EbnfCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_EbnfCtrlDownCast)
    if ok { return work.ToParser_EbnfCtrl() }
    return nil
}
func (obj *Parser_EbnfCtrl) ToParser_EbnfCtrl() *Parser_EbnfCtrl {
    return obj
}
func NewParser_EbnfCtrl(_env *LnsEnv) *Parser_EbnfCtrl {
    obj := &Parser_EbnfCtrl{}
    obj.FP = obj
    obj.InitParser_EbnfCtrl(_env)
    return obj
}
// 22: DeclConstr
func (self *Parser_EbnfCtrl) InitParser_EbnfCtrl(_env *LnsEnv) {
    self.element2ruleList = NewLnsMap( map[LnsAny]LnsAny{})
    self.allElementSet = NewLnsSet([]LnsAny{})
    self.elementName2Core = NewLnsMap( map[LnsAny]LnsAny{})
}


func Lns_Parser_init(_env *LnsEnv) {
    if init_Parser { return }
    init_Parser = true
    Parser__mod__ = "@lns.@Parser"
    Lns_InitMod()
    Lns_Ebnf_init(_env)
    Lns_Rule_init(_env)
    Lns_Code_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
}
func init() {
    init_Parser = false
}
