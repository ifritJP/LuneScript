// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
import LnsWriter "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
import LnsAsyncTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
var init_Formatter bool
var Formatter__mod__ string
var Formatter_termTxtSet *LnsSet2_[string]
var Formatter_beginTxtSet *LnsSet2_[string]
var Formatter_parenMap *LnsMap2_[string,string]
var Formatter_indentTermTxtSet *LnsSet2_[string]
var Formatter_indentDelta LnsInt
// decl alge -- ParenInfoIndent
type Formatter_ParenInfoIndent = LnsAny
type Formatter_ParenInfoIndent__Abs struct{
Val1 LnsInt
}
func (self *Formatter_ParenInfoIndent__Abs) GetTxt() string {
return "ParenInfoIndent.Abs"
}
type Formatter_ParenInfoIndent__RelIndent struct{
Val1 *Formatter_ParenInfo
Val2 LnsInt
}
func (self *Formatter_ParenInfoIndent__RelIndent) GetTxt() string {
return "ParenInfoIndent.RelIndent"
}
type Formatter_ParenInfoIndent__RelTerm struct{
Val1 *Formatter_ParenInfo
Val2 LnsInt
}
func (self *Formatter_ParenInfoIndent__RelTerm) GetTxt() string {
return "ParenInfoIndent.RelTerm"
}
type Formatter_ParenInfoIndent__RelTermPos struct{
Val1 *Formatter_ParenInfo
Val2 LnsInt
}
func (self *Formatter_ParenInfoIndent__RelTermPos) GetTxt() string {
return "ParenInfoIndent.RelTermPos"
}
// for 536
func Formatter_convExp3_405(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 []LnsAny) (*LnsEnv, *LnsTypes.Types_Token, LnsInt, LnsInt) {
    return _env, arg1, Lns_getFromMulti( arg2, 0 ).(LnsInt), Lns_getFromMulti( arg2, 1 ).(LnsInt)
}
// for 750
func Formatter_convExp4_390(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 *Formatter_ParenInfo, arg3 []LnsAny) (*LnsEnv, *LnsTypes.Types_Token, *Formatter_ParenInfo, LnsInt, LnsInt) {
    return _env, arg1, arg2, Lns_getFromMulti( arg3, 0 ).(LnsInt), Lns_getFromMulti( arg3, 1 ).(LnsInt)
}
// 33: decl @lns.@Formatter.getTokenTxt
func Formatter_getTokenTxt_0_(_env *LnsEnv, token *LnsTypes.Types_Token) string {
    var pos LnsTypes.Types_Position
    pos = token.Pos
    return _env.GetVM().String_format("%d:%d:%s", Lns_2DDD(pos.LineNo, pos.Column, token.Txt))
}





// 46: decl @lns.@Formatter.PickCoreToken.collect
func (self *Formatter_PickCoreToken) Collect(_env *LnsEnv, pickedTokenList *LnsList2_[*LnsTypes.Types_Token]) {
    if self.token.Kind != LnsTypes.Types_TokenKind__Eof{
        pickedTokenList.Insert(self.token)
    }
}
// 70: decl @lns.@Formatter.PickPareInfo.add
func (self *Formatter_PickPareInfo) Add(_env *LnsEnv, token *LnsTypes.Types_Token) {
    self.coreList.Insert(NewFormatter_PickCoreToken(_env, token).FP)
}
// 73: decl @lns.@Formatter.PickPareInfo.addCore
func (self *Formatter_PickPareInfo) AddCore(_env *LnsEnv, core Formatter_PickCore) {
    self.coreList.Insert(core)
}
// 76: decl @lns.@Formatter.PickPareInfo.clearSub
func (self *Formatter_PickPareInfo) clearSub(_env *LnsEnv) {
    var list *LnsList2_[Formatter_PickCore]
    list = NewLnsList2_[Formatter_PickCore]([]Formatter_PickCore{})
    list.Insert(self.coreList.GetAt(1))
    if self.closed{
        list.Insert(self.coreList.GetAt(self.coreList.Len()))
    }
    self.coreList = list
}
// 84: decl @lns.@Formatter.PickPareInfo.clear
func (self *Formatter_PickPareInfo) Clear(_env *LnsEnv) {
    for _, _core := range( self.coreList.Items ) {
        core := _core
        {
            _pick := Formatter_PickPareInfoDownCastF(core)
            if !Lns_IsNil( _pick ) {
                pick := _pick.(*Formatter_PickPareInfo)
                pick.FP.clearSub(_env)
            }
        }
    }
}
// 92: decl @lns.@Formatter.PickPareInfo.collect
func (self *Formatter_PickPareInfo) Collect(_env *LnsEnv, pickedTokenList *LnsList2_[*LnsTypes.Types_Token]) {
    for _, _core := range( self.coreList.Items ) {
        core := _core
        core.Collect(_env, pickedTokenList)
    }
}
// 114: decl @lns.@Formatter.CodePicker.getTokenMain
func (self *Formatter_CodePicker) GetTokenMain(_env *LnsEnv) LnsAny {
    if self.pickedTokenList.Len() < self.curIndex + 1{
        return nil
    }
    self.curIndex = self.curIndex + 1
    return self.pickedTokenList.GetAt(self.curIndex)
}
// 122: decl @lns.@Formatter.CodePicker.dump
func (self *Formatter_CodePicker) dump(_env *LnsEnv) {
    for range( self.pickedTokenList.Items ) {
        var prev *LnsTypes.Types_Token
        prev = self.FP.GetPrevToken(_env)
        var token *LnsTypes.Types_Token
        token = self.FP.GetToken(_env)
        if prev.Pos.LineNo != token.Pos.LineNo{
            Lns_print(Lns_2DDD(""))
        }
        Lns_io_stdout.Write(_env, token.Txt + " ")
    }
}
// 133: decl @lns.@Formatter.CodePicker.createFrom
func Formatter_CodePicker_createFrom(_env *LnsEnv, tokenizer Code_CodeTokenizerIF,startLineNo LnsInt,endLineNo LnsInt) *Formatter_CodePicker {
    var Formatter_process func(_env *LnsEnv, curInfo *Formatter_PickPareInfo) bool
    Formatter_process = func(_env *LnsEnv, curInfo *Formatter_PickPareInfo) bool {
        for  {
            var token *LnsTypes.Types_Token
            token = tokenizer.GetToken(_env)
            if token.Kind == LnsTypes.Types_TokenKind__Eof{
                return false
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( token.Txt != "case") &&
                _env.SetStackVal( Formatter_parenMap.Get(token.Txt)) )){
                var nextInfo *Formatter_PickPareInfo
                nextInfo = NewFormatter_PickPareInfo(_env, token)
                var result bool
                result = Formatter_process(_env, nextInfo)
                if result{
                    curInfo.FP.Clear(_env)
                }
                curInfo.FP.AddCore(_env, nextInfo.FP)
                if result{
                    return true
                }
            } else if token.Txt == curInfo.FP.Get_termTxt(_env){
                curInfo.FP.Add(_env, token)
                curInfo.FP.Set_closed(_env, true)
                if token.Pos.LineNo >= startLineNo{
                    return true
                }
                return false
            } else { 
                curInfo.FP.Add(_env, token)
            }
            if token.Pos.LineNo >= startLineNo{
                return true
            }
        }
    // insert a dummy
        return false
    }
    var pickPareInfoStack *LnsList2_[*Formatter_PickPareInfo]
    pickPareInfoStack = NewLnsList2_[*Formatter_PickPareInfo]([]*Formatter_PickPareInfo{})
    var curInfo *Formatter_PickPareInfo
    curInfo = NewFormatter_PickPareInfo(_env, LnsTypes.Types_noneToken)
    pickPareInfoStack.Insert(curInfo)
    Formatter_process(_env, curInfo)
    var codePicker *Formatter_CodePicker
    codePicker = NewFormatter_CodePicker(_env)
    pickPareInfoStack.GetAt(1).FP.Collect(_env, codePicker.pickedTokenList)
    for  {
        var token *LnsTypes.Types_Token
        token = tokenizer.GetToken(_env)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Kind == LnsTypes.Types_TokenKind__Eof) ||
            _env.SetStackVal( token.Pos.LineNo > endLineNo) ).(bool){
            break
        }
        codePicker.pickedTokenList.Insert(token)
    }
    return codePicker
}
// 227: decl @lns.@Formatter.Line2Indent.add
func (self *Formatter_Line2Indent) Add(_env *LnsEnv, lineNo LnsInt,lineIndent *Formatter_LineIndent) {
    if Lns_op_not(self.lineno2indent.Get(lineNo)){
        self.lineno2indent.Set(lineNo,lineIndent)
        self.indentList.Insert(lineIndent)
    }
}
// 233: decl @lns.@Formatter.Line2Indent.getPrevIndent
func (self *Formatter_Line2Indent) GetPrevIndent(_env *LnsEnv) *Formatter_LineIndent {
    return self.indentList.GetAt(self.indentList.Len())
}
// 236: decl @lns.@Formatter.Line2Indent.get
func (self *Formatter_Line2Indent) Get(_env *LnsEnv, lineNo LnsInt,column LnsInt) LnsInt {
    var lineIndent *Formatter_LineIndent
    
    {
        _lineIndent := self.lineno2indent.Get(lineNo)
        if _lineIndent == nil{
            return column
        } else {
            lineIndent = _lineIndent.(*Formatter_LineIndent)
        }
    }
    return lineIndent.FP.Get_delta(_env) + column
}
// 243: decl @lns.@Formatter.Line2Indent.outputResult
func (self *Formatter_Line2Indent) OutputResult(_env *LnsEnv, stream Lns_oStream) {
    var json *LnsWriter.Writer_JSON
    json = LnsWriter.NewWriter_JSON(_env, stream)
    json.FP.StartParent(_env, "indent", false)
    {
        _latest := self.latestLineIndent
        if !Lns_IsNil( _latest ) {
            latest := _latest.(*Formatter_LineIndent)
            json.FP.StartParent(_env, "latest", false)
            json.FP.Write(_env, "debugLineNo", latest.FP.Get_decidePos(_env))
            json.FP.Write(_env, "column", latest.FP.Get_indent(_env))
            json.FP.EndElement(_env)
        }
    }
    json.FP.StartParent(_env, "lines", true)
    {
        __forsortCollection0 := self.FP.Get_lineno2indent(_env)
        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
        for _, _line := range( __forsortSorted0.Items ) {
            lineIndent := __forsortCollection0.Items[ _line ]
            line := _line
            json.FP.StartParent(_env, "info", false)
            json.FP.Write(_env, "debugLineNo", lineIndent.FP.Get_decidePos(_env))
            json.FP.Write(_env, "column", lineIndent.FP.Get_indent(_env))
            json.FP.Write(_env, "lineNo", line)
            json.FP.Write(_env, "same", lineIndent.FP.Get_token(_env).Pos.Column == lineIndent.FP.Get_indent(_env))
            json.FP.EndElement(_env)
        }
    }
    json.FP.EndElement(_env)
    json.FP.EndElement(_env)
    json.FP.EndLayer(_env)
}
// 302: decl @lns.@Formatter.ParenInfo.set_frontToken
func (self *Formatter_ParenInfo) Set_frontToken(_env *LnsEnv, token LnsAny) {
    self.frontToken = token
    self.isWaitingFront = LnsTypes.Types_Token2Stem(token) == nil
}
// 309: decl @lns.@Formatter.ParenInfo.get_stmtIndent
func (self *Formatter_ParenInfo) Get_stmtIndent(_env *LnsEnv) LnsInt {
    return self.FP.get_indentPos(_env, self.indent)
}
// 312: decl @lns.@Formatter.ParenInfo.get_indent
func (self *Formatter_ParenInfo) Get_indent(_env *LnsEnv) LnsInt {
    var indent LnsInt
    indent = self.FP.get_indentPos(_env, self.indent)
    if self.isWaitingFront{
        return indent
    }
    return indent + Formatter_indentDelta
}
// 319: decl @lns.@Formatter.ParenInfo.get_termIndent
func (self *Formatter_ParenInfo) Get_termIndent(_env *LnsEnv) LnsInt {
    return self.FP.get_indentPos(_env, self.termIndent)
}
// 323: decl @lns.@Formatter.ParenInfo.get_indentPos
func (self *Formatter_ParenInfo) get_indentPos(_env *LnsEnv, parenInfoIndent LnsAny) LnsInt {
    switch _matchExp0 := parenInfoIndent.(type) {
    case *Formatter_ParenInfoIndent__Abs:
        column := _matchExp0.Val1
        return column
    case *Formatter_ParenInfoIndent__RelIndent:
        tokenIndent := _matchExp0.Val1
        offset := _matchExp0.Val2
        return tokenIndent.FP.Get_stmtIndent(_env) + offset
    case *Formatter_ParenInfoIndent__RelTerm:
        tokenIndent := _matchExp0.Val1
        offset := _matchExp0.Val2
        return tokenIndent.FP.Get_termIndent(_env) + offset
    case *Formatter_ParenInfoIndent__RelTermPos:
        tokenIndent := _matchExp0.Val1
        offset := _matchExp0.Val2
        var correct LnsInt
        correct = self.line2Indent.FP.Get(_env, self.parentInfo.startToken.Pos.LineNo, 0)
        return tokenIndent.FP.Get_termIndent(_env) + correct + offset
    }
// insert a dummy
    return 0
}
// 342: decl @lns.@Formatter.ParenInfo.set_contToken
func (self *Formatter_ParenInfo) Set_contToken(_env *LnsEnv, contToken *LnsTypes.Types_Token) {
    if self.isWaitingFront{
        self.FP.Set_frontToken(_env, contToken)
    }
    if Lns_isCondTrue( self.contToken){
        return 
    }
    if self.startToken.Pos.LineNo != contToken.Pos.LineNo{
        return 
    }
    self.contToken = contToken
    self.termIndent = &Formatter_ParenInfoIndent__RelTerm{self.parentInfo, len(self.startToken.Txt) + Formatter_indentDelta}
    if self.startToken.Txt == "case"{
    } else { 
        self.indent = &Formatter_ParenInfoIndent__RelTermPos{self.parentInfo, contToken.Pos.Column - self.parentInfo.FP.Get_termIndent(_env)}
        self.termIndent = &Formatter_ParenInfoIndent__RelTermPos{self.parentInfo, contToken.Pos.Column - self.parentInfo.FP.Get_termIndent(_env) - Formatter_indentDelta * 2}
    }
}
// 448: decl @lns.@Formatter.ParseCodeHook.prepare
func (self *Formatter_ParseCodeHook) Prepare(_env *LnsEnv, elementName string,depth LnsInt,token *LnsTypes.Types_Token) {
}
// 451: decl @lns.@Formatter.ParseCodeHook.dump
func (self *Formatter_ParseCodeHook) dump(_env *LnsEnv) {
    for _index, _parenInfo := range( self.parenInfoStack.Items ) {
        index := _index + 1
        parenInfo := _parenInfo
        if index > 1{
            Util_log(_env, Lns_2DDD("-------", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env)), parenInfo.FP.Get_termIndent(_env), parenInfo.FP.Get_stmtIndent(_env)))
        }
    }
}
// 461: decl @lns.@Formatter.ParseCodeHook.output
func (self *Formatter_ParseCodeHook) output(_env *LnsEnv, token *LnsTypes.Types_Token,lineNo LnsInt,indent LnsInt) {
    var pos LnsTypes.Types_Position
    pos = token.Pos
    {
        var commentIndent LnsInt
        if Formatter_termTxtSet.Has(token.Txt){
            var prev *Formatter_LineIndent
            prev = self.line2Indent.FP.GetPrevIndent(_env)
            if prev.FP.Get_indent(_env) == indent{
                commentIndent = indent + Formatter_indentDelta
            } else { 
                commentIndent = prev.FP.Get_indent(_env)
            }
        } else { 
            commentIndent = indent
        }
        for _, _comment := range( token.FP.Get_commentList(_env).Items ) {
            comment := _comment
            var work *Formatter_LineIndent
            work = NewFormatter_LineIndent(_env, comment, commentIndent, lineNo, commentIndent - pos.Column)
            self.line2Indent.FP.Add(_env, comment.Pos.LineNo, work)
        }
    }
    var lineIndent *Formatter_LineIndent
    lineIndent = NewFormatter_LineIndent(_env, token, indent, lineNo, indent - pos.Column)
    self.line2Indent.FP.Add(_env, pos.LineNo, lineIndent)
    self.targetLineNo = pos.LineNo + 1
    if pos.LineNo < self.endLineNo{
        return 
    }
    self.FP.dump(_env)
    self.line2Indent.FP.OutputResult(_env, Lns_io_stdout)
}
// 504: decl @lns.@Formatter.ParseCodeHook.checkKeyword
func (self *Formatter_ParseCodeHook) checkKeyword(_env *LnsEnv, token *LnsTypes.Types_Token) *LnsTypes.Types_Token {
    __func__ := "@lns.@Formatter.ParseCodeHook.checkKeyword"
    {
        if _switch0 := token.Kind; _switch0 == LnsTypes.Types_TokenKind__Str || _switch0 == LnsTypes.Types_TokenKind__Cmnt {
            var list *LnsList2_[string]
            list = LnsUtil.Util_splitStr(_env, token.Txt, "\n")
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( list.Len() > 0) &&
                _env.SetStackVal( token.Pos.LineNo <= self.targetLineNo) &&
                _env.SetStackVal( token.Pos.LineNo + list.Len() >= self.targetLineNo) ).(bool)){
                Util_log(_env, Lns_2DDD(__func__, _env.GetVM().String_format("skip -- %s", Lns_2DDD(LnsTypes.Types_TokenKind_getTxt( token.Kind)))))
                self.targetLineNo = token.Pos.LineNo + list.Len() + 1
                return token
            }
        }
    }
    var Formatter_process func(_env *LnsEnv)(LnsInt, LnsInt)
    Formatter_process = func(_env *LnsEnv)(LnsInt, LnsInt) {
        var parenInfo *Formatter_ParenInfo
        parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        return 526, parenInfo.FP.Get_indent(_env)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Txt == "___LNS___") ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Pos.LineNo > self.targetLineNo) ||
            _env.SetStackVal( token.Pos.LineNo == self.targetLineNo) ).(bool))) &&
        _env.SetStackVal( Lns_op_not(Formatter_termTxtSet.Has(token.Txt))) &&
        _env.SetStackVal( Lns_op_not(Formatter_beginTxtSet.Has(token.Txt))) ).(bool){
        self.FP.output(Formatter_convExp3_405(_env, token, Lns_2DDD(Formatter_process(_env))))
    }
    return token
}
// 542: decl @lns.@Formatter.ParseCodeHook.getToken
func (self *Formatter_ParseCodeHook) GetToken(_env *LnsEnv) *LnsTypes.Types_Token {
    return self.FP.checkKeyword(_env, self.tokenizer.GetToken(_env))
}
// 545: decl @lns.@Formatter.ParseCodeHook.peekToken
func (self *Formatter_ParseCodeHook) PeekToken(_env *LnsEnv) *LnsTypes.Types_Token {
    return self.FP.checkKeyword(_env, self.tokenizer.PeekToken(_env))
}
// 549: decl @lns.@Formatter.ParseCodeHook.process
func (self *Formatter_ParseCodeHook) Process(_env *LnsEnv, parseCodeRet LnsAny,depth LnsInt) LnsAny {
    __func__ := "@lns.@Formatter.ParseCodeHook.process"
    switch _matchExp0 := parseCodeRet.(type) {
    case *Code_ParseCodeRet__Detect:
        codeCore := _matchExp0.Val1
        var nowParenInfo *Formatter_ParenInfo
        nowParenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        var parenInfo *Formatter_ParenInfo
        parenInfo = nowParenInfo
        var Formatter_processTokens func(_env *LnsEnv, token *LnsTypes.Types_Token)
        Formatter_processTokens = func(_env *LnsEnv, token *LnsTypes.Types_Token) {
            __func__ := "@lns.@Formatter.ParseCodeHook.process.processTokens"
            var isWaitingFront bool
            isWaitingFront = false
            if Lns_isCondTrue( Formatter_parenMap.Get(token.Txt)){
                parenInfo = NewFormatter_ParenInfo(_env, self.line2Indent, self.parenInfoStack.GetAt(self.parenInfoStack.Len()), depth, token)
                self.parenInfoStack.Insert(parenInfo)
                Util_log(_env, Lns_2DDD(__func__, "push", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env))))
            } else if parenInfo.FP.Get_termTxt(_env) == token.Txt{
                Util_log(_env, Lns_2DDD(__func__, "pop", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env))))
                if parenInfo.FP.Get_depth(_env) >= depth{
                    self.parenInfoStack.Remove(nil)
                    parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
                }
                if parenInfo.FP.Get_startToken(_env).Txt == "case"{
                    self.parenInfoStack.Remove(nil)
                    parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
                }
                if token.Txt == "}"{
                    isWaitingFront = true
                }
            } else { 
                if Formatter_indentTermTxtSet.Has(token.Txt){
                    isWaitingFront = true
                } else { 
                    parenInfo.FP.Set_contToken(_env, token)
                }
            }
            if (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( token.Pos.LineNo > self.targetLineNo) ||
                _env.SetStackVal( token.Pos.LineNo == self.targetLineNo) ).(bool)){
                if Formatter_termTxtSet.Has(token.Txt){
                    var indent LnsInt
                    indent = nowParenInfo.FP.Get_termIndent(_env)
                    self.FP.output(_env, token, 590, indent)
                } else if Formatter_beginTxtSet.Has(token.Txt){
                    var indent LnsInt
                    indent = parenInfo.FP.Get_termIndent(_env)
                    self.FP.output(_env, token, 593, indent)
                }
            }
            if isWaitingFront{
                parenInfo.FP.Set_frontToken(_env, nil)
            }
        }
        {
            _codeCoreToken := Code_CodeCoreTokenDownCastF(codeCore)
            if !Lns_IsNil( _codeCoreToken ) {
                codeCoreToken := _codeCoreToken.(*Code_CodeCoreToken)
                var token *LnsTypes.Types_Token
                token = codeCoreToken.FP.Get_token(_env)
                Formatter_processTokens(_env, token)
            } else {
                {
                    _workCoreToken := Code_CodeCoreBuiltinDownCastF(codeCore)
                    if !Lns_IsNil( _workCoreToken ) {
                        workCoreToken := _workCoreToken.(*Code_CodeCoreBuiltin)
                        Formatter_processTokens(_env, workCoreToken.FP.Get_token(_env))
                    } else {
                        {
                            _workCoreToken := Code_CodeCoreStatDownCastF(codeCore)
                            if !Lns_IsNil( _workCoreToken ) {
                                workCoreToken := _workCoreToken.(*Code_CodeCoreStat)
                                for _, _token := range( workCoreToken.FP.Get_list(_env).Items ) {
                                    token := _token.(LnsTypes.Types_TokenDownCast).ToTypes_Token()
                                    Formatter_processTokens(_env, token)
                                }
                            } else {
                                {
                                    _workCoreToken := Code_CodeCoreStatOneDownCastF(codeCore)
                                    if !Lns_IsNil( _workCoreToken ) {
                                        workCoreToken := _workCoreToken.(*Code_CodeCoreStatOne)
                                        Formatter_processTokens(_env, workCoreToken.FP.Get_token(_env))
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    default:
        var parenInfo *Formatter_ParenInfo
        parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        for parenInfo.FP.Get_depth(_env) > depth {
            Util_log(_env, Lns_2DDD(__func__, "pushback", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env))))
            self.parenInfoStack.Remove(nil)
            parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        }
    }
    if depth == 0{
        self.line2Indent.FP.OutputResult(_env, Lns_io_stdout)
    }
    return parseCodeRet
}
// 674: decl @lns.@Formatter.SimpleParser.dump
func (self *Formatter_SimpleParser) dump(_env *LnsEnv) {
    for _index, _parenInfo := range( self.parenInfoStack.Items ) {
        index := _index + 1
        parenInfo := _parenInfo
        if index > 1{
            Util_log(_env, Lns_2DDD("-------", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env)), parenInfo.FP.Get_termIndent(_env), parenInfo.FP.Get_stmtIndent(_env)))
        }
    }
}
// 684: decl @lns.@Formatter.SimpleParser.output
func (self *Formatter_SimpleParser) output(_env *LnsEnv, token *LnsTypes.Types_Token,parenInfo *Formatter_ParenInfo,lineNo LnsInt,indent LnsInt) bool {
    var pos LnsTypes.Types_Position
    pos = token.Pos
    {
        var commentIndent LnsInt
        commentIndent = parenInfo.FP.Get_stmtIndent(_env)
        for _, _comment := range( token.FP.Get_commentList(_env).Items ) {
            comment := _comment
            var work *Formatter_LineIndent
            work = NewFormatter_LineIndent(_env, comment, commentIndent, lineNo, commentIndent - pos.Column)
            var list *LnsList2_[string]
            list = LnsUtil.Util_splitStr(_env, comment.Txt, "\n")
            if list.Len() > 0{
                for _index, _ := range( list.Items ) {
                    index := _index + 1
                    self.line2Indent.FP.Add(_env, comment.Pos.LineNo + index - 1, work)
                }
            } else { 
                self.line2Indent.FP.Add(_env, comment.Pos.LineNo, work)
            }
        }
    }
    var lineIndent *Formatter_LineIndent
    lineIndent = NewFormatter_LineIndent(_env, token, indent, lineNo, indent - pos.Column)
    self.line2Indent.FP.Add(_env, pos.LineNo, lineIndent)
    self.targetLineNo = pos.LineNo + 1
    return pos.LineNo < self.endLineNo
}
// 719: decl @lns.@Formatter.SimpleParser.checkKeyword
func (self *Formatter_SimpleParser) checkKeyword(_env *LnsEnv, token *LnsTypes.Types_Token) LnsAny {
    __func__ := "@lns.@Formatter.SimpleParser.checkKeyword"
    {
        if _switch0 := token.Kind; _switch0 == LnsTypes.Types_TokenKind__Str || _switch0 == LnsTypes.Types_TokenKind__Cmnt {
            var list *LnsList2_[string]
            list = LnsUtil.Util_splitStr(_env, token.Txt, "\n")
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( list.Len() > 0) &&
                _env.SetStackVal( token.Pos.LineNo <= self.targetLineNo) &&
                _env.SetStackVal( token.Pos.LineNo + list.Len() >= self.targetLineNo) ).(bool)){
                Util_log(_env, Lns_2DDD(__func__, _env.GetVM().String_format("skip -- %s", Lns_2DDD(LnsTypes.Types_TokenKind_getTxt( token.Kind)))))
                self.targetLineNo = token.Pos.LineNo + list.Len() + 1
                return token
            }
        }
    }
    var Formatter_process func(_env *LnsEnv)(LnsInt, LnsInt)
    Formatter_process = func(_env *LnsEnv)(LnsInt, LnsInt) {
        var parenInfo *Formatter_ParenInfo
        parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        return 741, parenInfo.FP.Get_indent(_env)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Txt == "___LNS___") ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Pos.LineNo > self.targetLineNo) ||
            _env.SetStackVal( token.Pos.LineNo == self.targetLineNo) ).(bool))) &&
        _env.SetStackVal( Lns_op_not(Formatter_termTxtSet.Has(token.Txt))) &&
        _env.SetStackVal( Lns_op_not(Formatter_beginTxtSet.Has(token.Txt))) ).(bool){
        if Lns_op_not(self.FP.output(Formatter_convExp4_390(_env, token, self.parenInfoStack.GetAt(self.parenInfoStack.Len()), Lns_2DDD(Formatter_process(_env))))){
            return nil
        }
    }
    return token
}
// 759: decl @lns.@Formatter.SimpleParser.getToken
func (self *Formatter_SimpleParser) getToken(_env *LnsEnv) LnsAny {
    return self.FP.checkKeyword(_env, self.tokenizer.GetToken(_env))
}
// 763: decl @lns.@Formatter.SimpleParser.process
func (self *Formatter_SimpleParser) process(_env *LnsEnv, token *LnsTypes.Types_Token) bool {
    __func__ := "@lns.@Formatter.SimpleParser.process"
    var nowParenInfo *Formatter_ParenInfo
    nowParenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
    var parenInfo *Formatter_ParenInfo
    parenInfo = nowParenInfo
    var depth LnsInt
    depth = 1
    var isWaitingFront bool
    isWaitingFront = false
    if Lns_isCondTrue( Formatter_parenMap.Get(token.Txt)){
        parenInfo = NewFormatter_ParenInfo(_env, self.line2Indent, self.parenInfoStack.GetAt(self.parenInfoStack.Len()), depth, token)
        self.parenInfoStack.Insert(parenInfo)
        Util_log(_env, Lns_2DDD(__func__, "push", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env))))
    } else if parenInfo.FP.Get_termTxt(_env) == token.Txt{
        Util_log(_env, Lns_2DDD(__func__, "pop", Formatter_getTokenTxt_0_(_env, parenInfo.FP.Get_startToken(_env))))
        if parenInfo.FP.Get_depth(_env) >= depth{
            self.parenInfoStack.Remove(nil)
            parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        }
        if parenInfo.FP.Get_startToken(_env).Txt == "case"{
            self.parenInfoStack.Remove(nil)
            parenInfo = self.parenInfoStack.GetAt(self.parenInfoStack.Len())
        }
        if token.Txt == "}"{
            isWaitingFront = true
        }
    } else { 
        if Formatter_indentTermTxtSet.Has(token.Txt){
            isWaitingFront = true
        } else { 
            parenInfo.FP.Set_contToken(_env, token)
        }
    }
    var _continue bool
    if (_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Pos.LineNo > self.targetLineNo) ||
        _env.SetStackVal( token.Pos.LineNo == self.targetLineNo) ).(bool)){
        if Formatter_termTxtSet.Has(token.Txt){
            var indent LnsInt
            indent = nowParenInfo.FP.Get_termIndent(_env)
            _continue = self.FP.output(_env, token, nowParenInfo, 804, indent)
        } else if Formatter_beginTxtSet.Has(token.Txt){
            var indent LnsInt
            indent = parenInfo.FP.Get_termIndent(_env)
            _continue = self.FP.output(_env, token, nowParenInfo, 807, indent)
        } else { 
            _continue = true
        }
    } else { 
        _continue = true
    }
    if isWaitingFront{
        parenInfo.FP.Set_frontToken(_env, nil)
    }
    return _continue
}
// 821: decl @lns.@Formatter.SimpleParser.analyze
func (self *Formatter_SimpleParser) Analyze(_env *LnsEnv) *Formatter_Line2Indent {
    for  {
        var token *LnsTypes.Types_Token
        
        {
            _token := self.FP.getToken(_env)
            if _token == nil{
                break
            } else {
                token = _token.(*LnsTypes.Types_Token)
            }
        }
        if token.Kind == LnsTypes.Types_TokenKind__Eof{
            break
        }
        if Lns_op_not(self.FP.process(_env, token)){
            break
        }
    }
    return self.line2Indent
}
type Formatter_PickCore interface {
        Collect(_env *LnsEnv, arg1 *LnsList2_[*LnsTypes.Types_Token])
}
func Lns_cast2Formatter_PickCore( obj LnsAny ) LnsAny {
    if _, ok := obj.(Formatter_PickCore); ok { 
        return obj
    }
    return nil
}

// declaration Class -- PickCoreToken
type Formatter_PickCoreTokenMtd interface {
    Collect(_env *LnsEnv, arg1 *LnsList2_[*LnsTypes.Types_Token])
    Get_token(_env *LnsEnv) *LnsTypes.Types_Token
}
type Formatter_PickCoreToken struct {
    token *LnsTypes.Types_Token
    FP Formatter_PickCoreTokenMtd
}
func Formatter_PickCoreToken2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_PickCoreToken).FP
}
func Formatter_PickCoreToken_toSlice(slice []LnsAny) []*Formatter_PickCoreToken {
    ret := make([]*Formatter_PickCoreToken, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_PickCoreTokenDownCast).ToFormatter_PickCoreToken()
    }
    return ret
}
type Formatter_PickCoreTokenDownCast interface {
    ToFormatter_PickCoreToken() *Formatter_PickCoreToken
}
func Formatter_PickCoreTokenDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_PickCoreTokenDownCast)
    if ok { return work.ToFormatter_PickCoreToken() }
    return nil
}
func (obj *Formatter_PickCoreToken) ToFormatter_PickCoreToken() *Formatter_PickCoreToken {
    return obj
}
func NewFormatter_PickCoreToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *Formatter_PickCoreToken {
    obj := &Formatter_PickCoreToken{}
    obj.FP = obj
    obj.InitFormatter_PickCoreToken(_env, arg1)
    return obj
}
func (self *Formatter_PickCoreToken) InitFormatter_PickCoreToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token) {
    self.token = arg1
}
func (self *Formatter_PickCoreToken) Get_token(_env *LnsEnv) *LnsTypes.Types_Token{ return self.token }

// declaration Class -- PickPareInfo
type Formatter_PickPareInfoMtd interface {
    Add(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
    AddCore(_env *LnsEnv, arg1 Formatter_PickCore)
    Clear(_env *LnsEnv)
    clearSub(_env *LnsEnv)
    Collect(_env *LnsEnv, arg1 *LnsList2_[*LnsTypes.Types_Token])
    Get_closed(_env *LnsEnv) bool
    Get_startToken(_env *LnsEnv) *LnsTypes.Types_Token
    Get_termTxt(_env *LnsEnv) string
    Set_closed(_env *LnsEnv, arg1 bool)
}
type Formatter_PickPareInfo struct {
    startToken *LnsTypes.Types_Token
    termTxt string
    coreList *LnsList2_[Formatter_PickCore]
    closed bool
    FP Formatter_PickPareInfoMtd
}
func Formatter_PickPareInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_PickPareInfo).FP
}
func Formatter_PickPareInfo_toSlice(slice []LnsAny) []*Formatter_PickPareInfo {
    ret := make([]*Formatter_PickPareInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_PickPareInfoDownCast).ToFormatter_PickPareInfo()
    }
    return ret
}
type Formatter_PickPareInfoDownCast interface {
    ToFormatter_PickPareInfo() *Formatter_PickPareInfo
}
func Formatter_PickPareInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_PickPareInfoDownCast)
    if ok { return work.ToFormatter_PickPareInfo() }
    return nil
}
func (obj *Formatter_PickPareInfo) ToFormatter_PickPareInfo() *Formatter_PickPareInfo {
    return obj
}
func NewFormatter_PickPareInfo(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *Formatter_PickPareInfo {
    obj := &Formatter_PickPareInfo{}
    obj.FP = obj
    obj.InitFormatter_PickPareInfo(_env, arg1)
    return obj
}
func (self *Formatter_PickPareInfo) Get_startToken(_env *LnsEnv) *LnsTypes.Types_Token{ return self.startToken }
func (self *Formatter_PickPareInfo) Get_termTxt(_env *LnsEnv) string{ return self.termTxt }
func (self *Formatter_PickPareInfo) Get_closed(_env *LnsEnv) bool{ return self.closed }
func (self *Formatter_PickPareInfo) Set_closed(_env *LnsEnv, arg1 bool){ self.closed = arg1 }
// 62: DeclConstr
func (self *Formatter_PickPareInfo) InitFormatter_PickPareInfo(_env *LnsEnv, startToken *LnsTypes.Types_Token) {
    self.startToken = startToken
    self.termTxt = Lns_unwrap( Formatter_parenMap.Get(startToken.Txt)).(string)
    self.coreList = NewLnsList2_[Formatter_PickCore]([]Formatter_PickCore{})
    self.coreList.Insert(NewFormatter_PickCoreToken(_env, startToken).FP)
    self.closed = false
}


// declaration Class -- CodePicker
type Formatter_CodePickerMtd interface {
    dump(_env *LnsEnv)
    GetPrevToken(_env *LnsEnv) *LnsTypes.Types_Token
    GetTailCommentList(_env *LnsEnv, arg1 *LnsTypes.Types_Token) LnsAny
    GetToken(_env *LnsEnv) *LnsTypes.Types_Token
    GetTokenMain(_env *LnsEnv) LnsAny
    PeekToken(_env *LnsEnv) *LnsTypes.Types_Token
    Pushback(_env *LnsEnv)
}
type Formatter_CodePicker struct {
    Code_CodeTokenizerBase
    pickedTokenList *LnsList2_[*LnsTypes.Types_Token]
    curIndex LnsInt
    FP Formatter_CodePickerMtd
}
func Formatter_CodePicker2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_CodePicker).FP
}
func Formatter_CodePicker_toSlice(slice []LnsAny) []*Formatter_CodePicker {
    ret := make([]*Formatter_CodePicker, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_CodePickerDownCast).ToFormatter_CodePicker()
    }
    return ret
}
type Formatter_CodePickerDownCast interface {
    ToFormatter_CodePicker() *Formatter_CodePicker
}
func Formatter_CodePickerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_CodePickerDownCast)
    if ok { return work.ToFormatter_CodePicker() }
    return nil
}
func (obj *Formatter_CodePicker) ToFormatter_CodePicker() *Formatter_CodePicker {
    return obj
}
func NewFormatter_CodePicker(_env *LnsEnv) *Formatter_CodePicker {
    obj := &Formatter_CodePicker{}
    obj.FP = obj
    obj.Code_CodeTokenizerBase.FP = obj
    obj.InitFormatter_CodePicker(_env)
    return obj
}
// 108: DeclConstr
func (self *Formatter_CodePicker) InitFormatter_CodePicker(_env *LnsEnv) {
    self.InitCode_CodeTokenizerBase(_env)
    self.pickedTokenList = NewLnsList2_[*LnsTypes.Types_Token]([]*LnsTypes.Types_Token{})
    self.curIndex = 0
}


// declaration Class -- LineIndent
type Formatter_LineIndentMtd interface {
    Get_decidePos(_env *LnsEnv) LnsInt
    Get_delta(_env *LnsEnv) LnsInt
    Get_indent(_env *LnsEnv) LnsInt
    Get_token(_env *LnsEnv) *LnsTypes.Types_Token
}
type Formatter_LineIndent struct {
    token *LnsTypes.Types_Token
    indent LnsInt
    decidePos LnsInt
    delta LnsInt
    FP Formatter_LineIndentMtd
}
func Formatter_LineIndent2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_LineIndent).FP
}
func Formatter_LineIndent_toSlice(slice []LnsAny) []*Formatter_LineIndent {
    ret := make([]*Formatter_LineIndent, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_LineIndentDownCast).ToFormatter_LineIndent()
    }
    return ret
}
type Formatter_LineIndentDownCast interface {
    ToFormatter_LineIndent() *Formatter_LineIndent
}
func Formatter_LineIndentDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_LineIndentDownCast)
    if ok { return work.ToFormatter_LineIndent() }
    return nil
}
func (obj *Formatter_LineIndent) ToFormatter_LineIndent() *Formatter_LineIndent {
    return obj
}
func NewFormatter_LineIndent(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) *Formatter_LineIndent {
    obj := &Formatter_LineIndent{}
    obj.FP = obj
    obj.InitFormatter_LineIndent(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Formatter_LineIndent) InitFormatter_LineIndent(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt) {
    self.token = arg1
    self.indent = arg2
    self.decidePos = arg3
    self.delta = arg4
}
func (self *Formatter_LineIndent) Get_token(_env *LnsEnv) *LnsTypes.Types_Token{ return self.token }
func (self *Formatter_LineIndent) Get_indent(_env *LnsEnv) LnsInt{ return self.indent }
func (self *Formatter_LineIndent) Get_decidePos(_env *LnsEnv) LnsInt{ return self.decidePos }
func (self *Formatter_LineIndent) Get_delta(_env *LnsEnv) LnsInt{ return self.delta }

// declaration Class -- Line2Indent
type Formatter_Line2IndentMtd interface {
    Add(_env *LnsEnv, arg1 LnsInt, arg2 *Formatter_LineIndent)
    Get(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) LnsInt
    GetPrevIndent(_env *LnsEnv) *Formatter_LineIndent
    Get_lineno2indent(_env *LnsEnv) *LnsMap2_[LnsInt,*Formatter_LineIndent]
    OutputResult(_env *LnsEnv, arg1 Lns_oStream)
}
type Formatter_Line2Indent struct {
    lineno2indent *LnsMap2_[LnsInt,*Formatter_LineIndent]
    indentList *LnsList2_[*Formatter_LineIndent]
    latestLineIndent LnsAny
    FP Formatter_Line2IndentMtd
}
func Formatter_Line2Indent2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_Line2Indent).FP
}
func Formatter_Line2Indent_toSlice(slice []LnsAny) []*Formatter_Line2Indent {
    ret := make([]*Formatter_Line2Indent, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_Line2IndentDownCast).ToFormatter_Line2Indent()
    }
    return ret
}
type Formatter_Line2IndentDownCast interface {
    ToFormatter_Line2Indent() *Formatter_Line2Indent
}
func Formatter_Line2IndentDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_Line2IndentDownCast)
    if ok { return work.ToFormatter_Line2Indent() }
    return nil
}
func (obj *Formatter_Line2Indent) ToFormatter_Line2Indent() *Formatter_Line2Indent {
    return obj
}
func NewFormatter_Line2Indent(_env *LnsEnv) *Formatter_Line2Indent {
    obj := &Formatter_Line2Indent{}
    obj.FP = obj
    obj.InitFormatter_Line2Indent(_env)
    return obj
}
func (self *Formatter_Line2Indent) Get_lineno2indent(_env *LnsEnv) *LnsMap2_[LnsInt,*Formatter_LineIndent]{ return self.lineno2indent }
// 220: DeclConstr
func (self *Formatter_Line2Indent) InitFormatter_Line2Indent(_env *LnsEnv) {
    self.latestLineIndent = nil
    self.lineno2indent = NewLnsMap2_[LnsInt,*Formatter_LineIndent]( map[LnsInt]*Formatter_LineIndent{})
    self.indentList = NewLnsList2_[*Formatter_LineIndent]([]*Formatter_LineIndent{})
    self.indentList.Insert(NewFormatter_LineIndent(_env, LnsTypes.Types_noneToken, 1, 1, 0))
}


// declaration Class -- ParenInfo
type Formatter_ParenInfoMtd interface {
    Get_depth(_env *LnsEnv) LnsInt
    Get_indent(_env *LnsEnv) LnsInt
    get_indentPos(_env *LnsEnv, arg1 LnsAny) LnsInt
    Get_startToken(_env *LnsEnv) *LnsTypes.Types_Token
    Get_stmtIndent(_env *LnsEnv) LnsInt
    Get_termIndent(_env *LnsEnv) LnsInt
    Get_termTxt(_env *LnsEnv) string
    Set_contToken(_env *LnsEnv, arg1 *LnsTypes.Types_Token)
    Set_frontToken(_env *LnsEnv, arg1 LnsAny)
}
type Formatter_ParenInfo struct {
    line2Indent *Formatter_Line2Indent
    parentInfo *Formatter_ParenInfo
    depth LnsInt
    startToken *LnsTypes.Types_Token
    termTxt string
    contToken LnsAny
    isWaitingFront bool
    frontToken LnsAny
    termIndent LnsAny
    indent LnsAny
    FP Formatter_ParenInfoMtd
}
func Formatter_ParenInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_ParenInfo).FP
}
func Formatter_ParenInfo_toSlice(slice []LnsAny) []*Formatter_ParenInfo {
    ret := make([]*Formatter_ParenInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_ParenInfoDownCast).ToFormatter_ParenInfo()
    }
    return ret
}
type Formatter_ParenInfoDownCast interface {
    ToFormatter_ParenInfo() *Formatter_ParenInfo
}
func Formatter_ParenInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_ParenInfoDownCast)
    if ok { return work.ToFormatter_ParenInfo() }
    return nil
}
func (obj *Formatter_ParenInfo) ToFormatter_ParenInfo() *Formatter_ParenInfo {
    return obj
}
func NewFormatter_ParenInfo(_env *LnsEnv, arg1 *Formatter_Line2Indent, arg2 LnsAny, arg3 LnsInt, arg4 *LnsTypes.Types_Token) *Formatter_ParenInfo {
    obj := &Formatter_ParenInfo{}
    obj.FP = obj
    obj.InitFormatter_ParenInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Formatter_ParenInfo) Get_depth(_env *LnsEnv) LnsInt{ return self.depth }
func (self *Formatter_ParenInfo) Get_startToken(_env *LnsEnv) *LnsTypes.Types_Token{ return self.startToken }
func (self *Formatter_ParenInfo) Get_termTxt(_env *LnsEnv) string{ return self.termTxt }

// 366: DeclConstr
func (self *Formatter_ParenInfo) InitFormatter_ParenInfo(_env *LnsEnv, line2Indent *Formatter_Line2Indent,parentInfo LnsAny,depth LnsInt,startToken *LnsTypes.Types_Token) {
    self.line2Indent = line2Indent
    self.parentInfo = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( parentInfo) ||
        _env.SetStackVal( self) ).(*Formatter_ParenInfo)
    self.depth = depth
    self.startToken = startToken
    {
        __exp := Formatter_parenMap.Get(startToken.Txt)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            self.termTxt = _exp
        } else {
            panic(_env.GetVM().String_format("illegal paren -- %s", Lns_2DDD(Formatter_getTokenTxt_0_(_env, startToken))))
        }
    }
    self.contToken = nil
    self.isWaitingFront = true
    self.frontToken = nil
    var termIndent LnsAny
    termIndent = &Formatter_ParenInfoIndent__Abs{1}
    var offset LnsInt
    offset = 0
    if Lns_isCondTrue( parentInfo){
        if self.parentInfo.startToken.Txt == ""{
            offset = 0
        } else { 
            offset = Formatter_indentDelta
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.parentInfo.frontToken) && 
            _env.NilAccPush(_env.NilAccPop().(*LnsTypes.Types_Token).Pos)&&
            _env.NilAccPush(_env.NilAccPop().(LnsTypes.Types_Position).LineNo)) != startToken.Pos.LineNo) &&
            _env.SetStackVal( Lns_op_not(self.parentInfo.isWaitingFront)) &&
            _env.SetStackVal( Lns_op_not((_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( startToken.Txt == "{") ||
                _env.SetStackVal( startToken.Txt == "`{") ).(bool)))) ).(bool)){
            offset = offset + Formatter_indentDelta
        }
        termIndent = &Formatter_ParenInfoIndent__RelTerm{self.parentInfo, offset}
    }
    var indent LnsAny
    indent = &Formatter_ParenInfoIndent__Abs{1}
    if Lns_isCondTrue( parentInfo){
        indent = &Formatter_ParenInfoIndent__RelTerm{self.parentInfo, offset + Formatter_indentDelta}
    }
    if startToken.Txt == "case"{
        indent = &Formatter_ParenInfoIndent__RelTerm{self.parentInfo, Formatter_indentDelta + 5}
    } else { 
        if parentInfo != nil{
            parentInfo_184 := parentInfo.(*Formatter_ParenInfo)
            if parentInfo_184.startToken.Txt == "case"{
                termIndent = parentInfo_184.parentInfo.indent
                indent = &Formatter_ParenInfoIndent__RelTerm{self.parentInfo.parentInfo, Formatter_indentDelta * 2}
            }
        }
    }
    self.termIndent = termIndent
    self.indent = indent
}


// declaration Class -- ParseCodeHook
type Formatter_ParseCodeHookMtd interface {
    checkKeyword(_env *LnsEnv, arg1 *LnsTypes.Types_Token) *LnsTypes.Types_Token
    dump(_env *LnsEnv)
    GetPrevToken(_env *LnsEnv) *LnsTypes.Types_Token
    GetTailCommentList(_env *LnsEnv, arg1 *LnsTypes.Types_Token) LnsAny
    GetToken(_env *LnsEnv) *LnsTypes.Types_Token
    output(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 LnsInt, arg3 LnsInt)
    PeekToken(_env *LnsEnv) *LnsTypes.Types_Token
    Prepare(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 *LnsTypes.Types_Token)
    Process(_env *LnsEnv, arg1 LnsAny, arg2 LnsInt) LnsAny
    Pushback(_env *LnsEnv)
}
type Formatter_ParseCodeHook struct {
    tokenizer Code_CodeTokenizerIF
    targetLineNo LnsInt
    endLineNo LnsInt
    multiLine bool
    line2Indent *Formatter_Line2Indent
    parenInfoStack *LnsList2_[*Formatter_ParenInfo]
    FP Formatter_ParseCodeHookMtd
}
func Formatter_ParseCodeHook2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_ParseCodeHook).FP
}
func Formatter_ParseCodeHook_toSlice(slice []LnsAny) []*Formatter_ParseCodeHook {
    ret := make([]*Formatter_ParseCodeHook, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_ParseCodeHookDownCast).ToFormatter_ParseCodeHook()
    }
    return ret
}
type Formatter_ParseCodeHookDownCast interface {
    ToFormatter_ParseCodeHook() *Formatter_ParseCodeHook
}
func Formatter_ParseCodeHookDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_ParseCodeHookDownCast)
    if ok { return work.ToFormatter_ParseCodeHook() }
    return nil
}
func (obj *Formatter_ParseCodeHook) ToFormatter_ParseCodeHook() *Formatter_ParseCodeHook {
    return obj
}
func NewFormatter_ParseCodeHook(_env *LnsEnv, arg1 Code_CodeTokenizerIF, arg2 LnsInt, arg3 LnsInt) *Formatter_ParseCodeHook {
    obj := &Formatter_ParseCodeHook{}
    obj.FP = obj
    obj.InitFormatter_ParseCodeHook(_env, arg1, arg2, arg3)
    return obj
}
// advertise -- 420
func (self *Formatter_ParseCodeHook) GetPrevToken(_env *LnsEnv) *LnsTypes.Types_Token {
    return self.tokenizer. GetPrevToken( _env)
}
// advertise -- 420
func (self *Formatter_ParseCodeHook) GetTailCommentList(_env *LnsEnv, arg1 *LnsTypes.Types_Token) LnsAny {
    return self.tokenizer. GetTailCommentList( _env, arg1)
}
// advertise -- 420
func (self *Formatter_ParseCodeHook) Pushback(_env *LnsEnv) {
self.tokenizer. Pushback( _env)
}
// 433: DeclConstr
func (self *Formatter_ParseCodeHook) InitFormatter_ParseCodeHook(_env *LnsEnv, tokenizer Code_CodeTokenizerIF,startLineNo LnsInt,endLineNo LnsInt) {
    var line2Indent *Formatter_Line2Indent
    line2Indent = NewFormatter_Line2Indent(_env)
    self.line2Indent = line2Indent
    self.parenInfoStack = NewLnsList2_[*Formatter_ParenInfo]([]*Formatter_ParenInfo{})
    self.parenInfoStack.Insert(NewFormatter_ParenInfo(_env, line2Indent, nil, 1, LnsTypes.Types_noneToken))
    self.tokenizer = tokenizer
    self.targetLineNo = startLineNo
    self.endLineNo = endLineNo
    self.multiLine = startLineNo != endLineNo
}


// declaration Class -- SimpleParser
type Formatter_SimpleParserMtd interface {
    Analyze(_env *LnsEnv) *Formatter_Line2Indent
    checkKeyword(_env *LnsEnv, arg1 *LnsTypes.Types_Token) LnsAny
    dump(_env *LnsEnv)
    getToken(_env *LnsEnv) LnsAny
    output(_env *LnsEnv, arg1 *LnsTypes.Types_Token, arg2 *Formatter_ParenInfo, arg3 LnsInt, arg4 LnsInt) bool
    process(_env *LnsEnv, arg1 *LnsTypes.Types_Token) bool
}
type Formatter_SimpleParser struct {
    tokenizer Code_CodeTokenizerIF
    targetLineNo LnsInt
    endLineNo LnsInt
    line2Indent *Formatter_Line2Indent
    parenInfoStack *LnsList2_[*Formatter_ParenInfo]
    FP Formatter_SimpleParserMtd
}
func Formatter_SimpleParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_SimpleParser).FP
}
func Formatter_SimpleParser_toSlice(slice []LnsAny) []*Formatter_SimpleParser {
    ret := make([]*Formatter_SimpleParser, len(slice))
    for index, val := range slice {
        ret[index] = val.(Formatter_SimpleParserDownCast).ToFormatter_SimpleParser()
    }
    return ret
}
type Formatter_SimpleParserDownCast interface {
    ToFormatter_SimpleParser() *Formatter_SimpleParser
}
func Formatter_SimpleParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_SimpleParserDownCast)
    if ok { return work.ToFormatter_SimpleParser() }
    return nil
}
func (obj *Formatter_SimpleParser) ToFormatter_SimpleParser() *Formatter_SimpleParser {
    return obj
}
func NewFormatter_SimpleParser(_env *LnsEnv, arg1 Code_CodeTokenizerIF, arg2 LnsInt, arg3 LnsInt) *Formatter_SimpleParser {
    obj := &Formatter_SimpleParser{}
    obj.FP = obj
    obj.InitFormatter_SimpleParser(_env, arg1, arg2, arg3)
    return obj
}
// 659: DeclConstr
func (self *Formatter_SimpleParser) InitFormatter_SimpleParser(_env *LnsEnv, tokenizer Code_CodeTokenizerIF,startLineNo LnsInt,endLineNo LnsInt) {
    var line2Indent *Formatter_Line2Indent
    line2Indent = NewFormatter_Line2Indent(_env)
    self.line2Indent = line2Indent
    self.parenInfoStack = NewLnsList2_[*Formatter_ParenInfo]([]*Formatter_ParenInfo{})
    self.parenInfoStack.Insert(NewFormatter_ParenInfo(_env, line2Indent, nil, 1, LnsTypes.Types_noneToken))
    self.tokenizer = tokenizer
    self.targetLineNo = startLineNo
    self.endLineNo = endLineNo
}


func Lns_Formatter_init(_env *LnsEnv) {
    if init_Formatter { return }
    init_Formatter = true
    Formatter__mod__ = "@lns.@Formatter"
    Lns_InitMod()
    Lns_Code_init(_env)
    Lns_Util_init(_env)
    LnsTypes.Lns_Types_init(_env)
    LnsWriter.Lns_Writer_init(_env)
    LnsUtil.Lns_Util_init(_env)
    LnsAsyncTokenizer.Lns_AsyncTokenizer_init(_env)
    Formatter_termTxtSet = NewLnsSet2_[string](Lns_2DDDGen[string]("]", ")", "}"))
    Formatter_beginTxtSet = NewLnsSet2_[string](Lns_2DDDGen[string]("[", "(", "{", "[@", "(@", "[=", "`{"))
    Formatter_parenMap = NewLnsMap2_[string,string]( map[string]string{"[":"]","(":")","{":"}","[@":"]","(@":")","[=":"]","`{":"}","case":"","":"",})
    Formatter_indentTermTxtSet = (NewLnsSet2_[string](Lns_2DDDGen[string](",", ";")))
    Formatter_indentDelta = 3
}
func init() {
    init_Formatter = false
}
