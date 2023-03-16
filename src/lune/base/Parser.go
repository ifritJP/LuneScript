// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Parser bool
var Parser__mod__ string
var Parser_noneToken *Types_Token
var Parser_eofToken *Types_Token
// for 87
func Parser_convExp0_827(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 182
func Parser_convExp0_312(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
type Parser_TxtStream = Util_TxtStream
type Parser_Position = Types_Position
type Parser_TokenKind = Types_TokenKind
type Parser_Token = Types_Token
// 34: decl @lune.@base.@Parser.isLuaKeyword
func Parser_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncParser_isLuaKeyword(_env, txt)
}

// 69: decl @lune.@base.@Parser.convFromRawToStr
func Parser_convFromRawToStr(_env *LnsEnv, txt string) string {
    if len(txt) == 0{
        return txt
    }
    if _switch0 := LnsInt(txt[1-1]); _switch0 == 39 || _switch0 == 34 {
    } else {
        return _env.GetVM().String_sub(txt,4, len(txt) - 3)
    }
    var findChar LnsInt
    findChar = LnsInt(txt[1-1])
    var workTxt string
    workTxt = txt
    var retTxt string
    retTxt = ""
    var workIndex LnsInt
    workIndex = 2
    var setIndex LnsInt
    setIndex = 2
    var workPattern string
    workPattern = "[\"'\\]"
    for  {
        var endIndex LnsInt
        
        {
            _endIndex := Parser_convExp0_827(Lns_2DDD(_env.GetVM().String_find(workTxt, workPattern, workIndex, nil)))
            if _endIndex == nil{
                Util_err(_env, _env.GetVM().String_format("error: illegal string -- %s", Lns_2DDD(workTxt)))
            } else {
                endIndex = _endIndex.(LnsInt)
            }
        }
        var workChar LnsInt
        workChar = LnsInt(workTxt[endIndex-1])
        if workChar == findChar{
            return retTxt + _env.GetVM().String_sub(workTxt,setIndex, endIndex - 1)
        } else if workChar == 92{
            var quote LnsInt
            quote = LnsInt(workTxt[endIndex + 1-1])
            if _switch1 := quote; _switch1 == 39 || _switch1 == 34 {
                retTxt = _env.GetVM().String_format("%s%s%c", Lns_2DDD(retTxt, _env.GetVM().String_sub(workTxt,setIndex, endIndex - 1), quote))
            } else {
                retTxt = _env.GetVM().String_format("%s%s", Lns_2DDD(retTxt, _env.GetVM().String_sub(workTxt,setIndex, endIndex + 1)))
            }
            workIndex = endIndex + 2
            setIndex = workIndex
        } else { 
            workIndex = endIndex + 1
        }
    }
// insert a dummy
    return ""
}

// 360: decl @lune.@base.@Parser.isOp2
func Parser_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncParser_isOp2(_env, ope)
}

// 364: decl @lune.@base.@Parser.isOp1
func Parser_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncParser_isOp1(_env, ope)
}

// 373: decl @lune.@base.@Parser.getEofToken
func Parser_getEofToken(_env *LnsEnv) *Types_Token {
    return Parser_eofToken
}

// 463: decl @lune.@base.@Parser.quoteStr
func Parser_quoteStr(_env *LnsEnv, txt string) string {
    var work string
    work = txt
    var part string
    part = "\""
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = len(work)
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            var char LnsInt
            char = LnsInt(work[index-1])
            if _switch0 := char; _switch0 == 10 {
                part = part + "\\n"
            } else if _switch0 == 13 {
                part = part + "\\r"
            } else if _switch0 == 9 {
                part = part + "\\t"
            } else if _switch0 == 34 {
                part = part + "\\\""
            } else if _switch0 == 92 {
                part = part + "\\\\"
            } else {
                part = part + _env.GetVM().String_format("%c", Lns_2DDD(char))
            }
        }
    }
    work = part + "\""
    return work
}

// 493: decl @lune.@base.@Parser.createTokenizerFrom
func Parser_createTokenizerFrom(_env *LnsEnv, src LnsAny,async bool,stdinFile LnsAny) *Parser_Tokenizer {
    return &NewParser_StreamTokenizer(_env, src, async, stdinFile, nil).Parser_Tokenizer
}

// 131: decl @lune.@base.@Parser.TokenListTokenizer.createPosition
func (self *Parser_TokenListTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}
// 135: decl @lune.@base.@Parser.TokenListTokenizer.getStreamName
func (self *Parser_TokenListTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.streamName
}
// 139: decl @lune.@base.@Parser.TokenListTokenizer.getToken
func (self *Parser_TokenListTokenizer) GetToken(_env *LnsEnv) LnsAny {
    if self.tokenList.Len() < self.index{
        return nil
    }
    var token *Types_Token
    token = self.tokenList.GetAt(self.index)
    self.index = self.index + 1
    return token
}
// 160: decl @lune.@base.@Parser.StreamTokenizer.setStdinStream
func Parser_StreamTokenizer_setStdinStream(_env *LnsEnv, moduleName string) {
    Parser_StreamTokenizer__stdinStreamModuleName = moduleName
    Parser_StreamTokenizer__stdinTxt = Lns_unwrapDefault( Lns_io_stdin.Read(_env, "*a"), "").(string)
}
// 191: decl @lune.@base.@Parser.StreamTokenizer.createPosition
func (self *Parser_StreamTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}
// 196: decl @lune.@base.@Parser.StreamTokenizer.getStreamName
func (self *Parser_StreamTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.streamName
}
// 200: decl @lune.@base.@Parser.StreamTokenizer.create
func Parser_StreamTokenizer_create(_env *LnsEnv, tokenizerSrc LnsAny,async bool,stdinFile LnsAny,pos LnsAny) *Parser_StreamTokenizer {
    return NewParser_StreamTokenizer(_env, tokenizerSrc, async, stdinFile, pos)
}
// 207: decl @lune.@base.@Parser.StreamTokenizer.getToken
func (self *Parser_StreamTokenizer) GetToken(_env *LnsEnv) LnsAny {
    if self.lineTokenList.Len() < self.pos{
        self.pos = 1
        self.lineTokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
        for self.lineTokenList.Len() == 0 {
            var pipeItem *Async_PipeItem
            
            {
                _pipeItem := self.asyncTokenizer.FP.GetNext(_env)
                if _pipeItem == nil{
                    return nil
                } else {
                    pipeItem = _pipeItem.(*Async_PipeItem)
                }
            }
            self.lineTokenList = pipeItem.FP.Get_item(_env).(*AsyncParser_AsyncItem).List
        }
    }
    var token *Types_Token
    token = self.lineTokenList.GetAt(self.pos)
    self.pos = self.pos + 1
    return token
}
// 239: decl @lune.@base.@Parser.DefaultPushbackTokenizer.getUsedTokenListLen
func (self *Parser_DefaultPushbackTokenizer) GetUsedTokenListLen(_env *LnsEnv) LnsInt {
    return self.usedTokenList.Len()
}
// 243: decl @lune.@base.@Parser.DefaultPushbackTokenizer.createFromLnsCode
func Parser_DefaultPushbackTokenizer_createFromLnsCode(_env *LnsEnv, code string,name string) *Parser_DefaultPushbackTokenizer {
    return NewParser_DefaultPushbackTokenizer(_env, &NewParser_StreamTokenizer(_env, &Types_TokenizerSrc__LnsCode{code, name, nil}, false, nil, nil).Parser_Tokenizer)
}
// 250: decl @lune.@base.@Parser.DefaultPushbackTokenizer.createPosition
func (self *Parser_DefaultPushbackTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return self.tokenizer.FP.CreatePosition(_env, lineNo, column)
}
// 254: decl @lune.@base.@Parser.DefaultPushbackTokenizer.getTokenNoErr
func (self *Parser_DefaultPushbackTokenizer) GetTokenNoErr(_env *LnsEnv, skipFlag LnsAny) *Types_Token {
    if self.pushbackedList.Len() > 0{
        self.currentToken = self.pushbackedList.GetAt(self.pushbackedList.Len())
        self.pushbackedList.Remove(nil)
    } else { 
        {
            _token := self.tokenizer.FP.GetToken(_env)
            if !Lns_IsNil( _token ) {
                token := _token.(*Types_Token)
                self.currentToken = token
            } else {
                self.currentToken = Types_noneToken
            }
        }
    }
    if self.currentToken.Kind != Types_TokenKind__Eof{
        self.usedTokenList.Insert(self.currentToken)
    }
    return self.currentToken
}
// 272: decl @lune.@base.@Parser.DefaultPushbackTokenizer.pushbackToken
func (self *Parser_DefaultPushbackTokenizer) PushbackToken(_env *LnsEnv, token *Types_Token) {
    if token.Kind != Types_TokenKind__Eof{
        self.pushbackedList.Insert(token)
    }
    if token == self.currentToken{
        if self.usedTokenList.Len() > 0{
            var used *Types_Token
            used = self.usedTokenList.GetAt(self.usedTokenList.Len())
            if used == token{
                self.usedTokenList.Remove(nil)
            }
            if self.usedTokenList.Len() > 0{
                self.currentToken = self.usedTokenList.GetAt(self.usedTokenList.Len())
            } else { 
                self.currentToken = Types_noneToken
            }
        } else { 
            self.currentToken = Types_noneToken
        }
    }
}
// 295: decl @lune.@base.@Parser.DefaultPushbackTokenizer.pushback
func (self *Parser_DefaultPushbackTokenizer) Pushback(_env *LnsEnv) {
    self.FP.PushbackToken(_env, self.currentToken)
}
// 298: decl @lune.@base.@Parser.DefaultPushbackTokenizer.pushbackStr
func (self *Parser_DefaultPushbackTokenizer) PushbackStr(_env *LnsEnv, asyncParse LnsAny,name string,statement string,pos Types_Position) {
    var tokenizer *Parser_StreamTokenizer
    tokenizer = NewParser_StreamTokenizer(_env, &Types_TokenizerSrc__LnsCode{statement, name, nil}, asyncParse == true, nil, pos)
    var list *LnsList2_[*Types_Token]
    list = NewLnsList2_[*Types_Token]([]*Types_Token{})
    for  {
        {
            __exp := tokenizer.FP.GetToken(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Types_Token)
                list.Insert(_exp)
            } else {
                break
            }
        }
    }
    {
        var _forFrom0 LnsInt = list.Len()
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
            self.FP.PushbackToken(_env, list.GetAt(index))
            _forWork0 += _forDelta0
        }
    }
}
// 318: decl @lune.@base.@Parser.DefaultPushbackTokenizer.newPushback
func (self *Parser_DefaultPushbackTokenizer) NewPushback(_env *LnsEnv, tokenList *LnsList2_[*Types_Token]) {
    {
        var _forFrom0 LnsInt = tokenList.Len()
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
            self.FP.PushbackToken(_env, tokenList.GetAt(index))
            _forWork0 += _forDelta0
        }
    }
}
// 324: decl @lune.@base.@Parser.DefaultPushbackTokenizer.getLastPos
func (self *Parser_DefaultPushbackTokenizer) GetLastPos(_env *LnsEnv) Types_Position {
    var pos Types_Position
    pos = self.tokenizer.FP.CreatePosition(_env, 0, 0)
    if self.FP.Get_currentToken(_env).Kind != Types_TokenKind__Eof{
        pos = self.FP.Get_currentToken(_env).Pos
    } else { 
        if self.usedTokenList.Len() > 0{
            var token *Types_Token
            token = self.usedTokenList.GetAt(self.usedTokenList.Len())
            pos = token.Pos
        }
    }
    return pos
}
// 340: decl @lune.@base.@Parser.DefaultPushbackTokenizer.getNearCode
func (self *Parser_DefaultPushbackTokenizer) GetNearCode(_env *LnsEnv) string {
    var code string
    code = ""
    {
        var _forFrom0 LnsInt = self.usedTokenList.Len() - 30
        var _forTo0 LnsInt = self.usedTokenList.Len()
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            if index > 1{
                var token *Types_Token
                token = self.usedTokenList.GetAt(index)
                if token.Consecutive{
                    code = _env.GetVM().String_format("%s%s", Lns_2DDD(code, self.usedTokenList.GetAt(index).Txt))
                } else { 
                    code = _env.GetVM().String_format("%s %s", Lns_2DDD(code, self.usedTokenList.GetAt(index).Txt))
                }
            }
        }
    }
    return _env.GetVM().String_format("%s -- current '%s'", Lns_2DDD(code, self.currentToken.Txt))
}
// 355: decl @lune.@base.@Parser.DefaultPushbackTokenizer.getStreamName
func (self *Parser_DefaultPushbackTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.tokenizer.FP.GetStreamName(_env)
}
// 377: decl @lune.@base.@Parser.DummyTokenizer.getToken
func (self *Parser_DummyTokenizer) GetToken(_env *LnsEnv) LnsAny {
    return Parser_eofToken
}
// 380: decl @lune.@base.@Parser.DummyTokenizer.getStreamName
func (self *Parser_DummyTokenizer) GetStreamName(_env *LnsEnv) string {
    return "dummy"
}
// 383: decl @lune.@base.@Parser.DummyTokenizer.createPosition
func (self *Parser_DummyTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), nil)
}
// 402: decl @lune.@base.@Parser.CommentLayer.addDirect
func (self *Parser_CommentLayer) AddDirect(_env *LnsEnv, commentList *LnsList2_[*Types_Token]) {
    for _, _comment := range( commentList.Items ) {
        comment := _comment
        self.commentList.Insert(comment)
    }
}
// 408: decl @lune.@base.@Parser.CommentLayer.add
func (self *Parser_CommentLayer) Add(_env *LnsEnv, token *Types_Token) {
    if Lns_op_not(self.tokenSet.Has(token)){
        self.tokenSet.Add(token)
        self.tokenList.Insert(token)
        self.FP.AddDirect(_env, token.FP.Get_commentList(_env))
    }
}
// 417: decl @lune.@base.@Parser.CommentLayer.clear
func (self *Parser_CommentLayer) Clear(_env *LnsEnv) {
    if self.commentList.Len() != 0{
        self.commentList = NewLnsList2_[*Types_Token]([]*Types_Token{})
        self.tokenSet = NewLnsSet2_[*Types_Token]([]*Types_Token{})
        self.tokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    }
}
// 433: decl @lune.@base.@Parser.CommentLayer.hasInvalidComment
func (self *Parser_CommentLayer) HasInvalidComment(_env *LnsEnv) LnsAny {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.tokenList.Len() > 1) &&
        _env.SetStackVal( self.tokenList.GetAt(2).FP.Get_commentList(_env).GetAt(1)) ||
        _env.SetStackVal( nil) )
}
// 447: decl @lune.@base.@Parser.CommentCtrl.push
func (self *Parser_CommentCtrl) Push(_env *LnsEnv) {
    self.layer = NewParser_CommentLayer(_env)
    self.layerStack.Insert(self.layer)
}
// 452: decl @lune.@base.@Parser.CommentCtrl.pop
func (self *Parser_CommentCtrl) Pop(_env *LnsEnv) {
    self.layer = self.layerStack.GetAt(self.layerStack.Len())
    self.layerStack.Remove(nil)
}
// declaration Class -- Tokenizer
type Parser_TokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_Tokenizer struct {
    FP Parser_TokenizerMtd
}
func Parser_Tokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_Tokenizer).FP
}
func Parser_Tokenizer_toSlice(slice []LnsAny) []*Parser_Tokenizer {
    ret := make([]*Parser_Tokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_TokenizerDownCast).ToParser_Tokenizer()
    }
    return ret
}
type Parser_TokenizerDownCast interface {
    ToParser_Tokenizer() *Parser_Tokenizer
}
func Parser_TokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_TokenizerDownCast)
    if ok { return work.ToParser_Tokenizer() }
    return nil
}
func (obj *Parser_Tokenizer) ToParser_Tokenizer() *Parser_Tokenizer {
    return obj
}
func (self *Parser_Tokenizer) InitParser_Tokenizer(_env *LnsEnv) {
}




type Parser_PushbackTokenizer interface {
        CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
        GetStreamName(_env *LnsEnv) string
        GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
        NewPushback(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
        Pushback(_env *LnsEnv)
        PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position)
        PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
func Lns_cast2Parser_PushbackTokenizer( obj LnsAny ) LnsAny {
    if _, ok := obj.(Parser_PushbackTokenizer); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TokenListTokenizer
type Parser_TokenListTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_TokenListTokenizer struct {
    Parser_Tokenizer
    streamName string
    tokenList *LnsList2_[*Types_Token]
    index LnsInt
    overridePos LnsAny
    FP Parser_TokenListTokenizerMtd
}
func Parser_TokenListTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_TokenListTokenizer).FP
}
func Parser_TokenListTokenizer_toSlice(slice []LnsAny) []*Parser_TokenListTokenizer {
    ret := make([]*Parser_TokenListTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_TokenListTokenizerDownCast).ToParser_TokenListTokenizer()
    }
    return ret
}
type Parser_TokenListTokenizerDownCast interface {
    ToParser_TokenListTokenizer() *Parser_TokenListTokenizer
}
func Parser_TokenListTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_TokenListTokenizerDownCast)
    if ok { return work.ToParser_TokenListTokenizer() }
    return nil
}
func (obj *Parser_TokenListTokenizer) ToParser_TokenListTokenizer() *Parser_TokenListTokenizer {
    return obj
}
func NewParser_TokenListTokenizer(_env *LnsEnv, arg1 *LnsList2_[*Types_Token], arg2 string, arg3 LnsAny) *Parser_TokenListTokenizer {
    obj := &Parser_TokenListTokenizer{}
    obj.FP = obj
    obj.Parser_Tokenizer.FP = obj
    obj.InitParser_TokenListTokenizer(_env, arg1, arg2, arg3)
    return obj
}
// 120: DeclConstr
func (self *Parser_TokenListTokenizer) InitParser_TokenListTokenizer(_env *LnsEnv, tokenList *LnsList2_[*Types_Token],streamName string,overridePos LnsAny) {
    self.InitParser_Tokenizer(_env)
    self.index = 1
    self.tokenList = tokenList
    self.streamName = streamName
    self.overridePos = overridePos
}


// declaration Class -- StreamTokenizer
var Parser_StreamTokenizer__stdinStreamModuleName LnsAny
var Parser_StreamTokenizer__stdinTxt string
// 151: decl @lune.@base.@Parser.StreamTokenizer.___init
func Parser_StreamTokenizer____init_0_(_env *LnsEnv) {
    Parser_StreamTokenizer__stdinStreamModuleName = nil
    Parser_StreamTokenizer__stdinTxt = ""
}

type Parser_StreamTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_StreamTokenizer struct {
    Parser_Tokenizer
    stdinStreamModuleName LnsAny
    stdinTxt string
    streamName string
    pos LnsInt
    lineTokenList *LnsList2_[*Types_Token]
    asyncTokenizer *AsyncParser_Tokenizer
    overridePos LnsAny
    FP Parser_StreamTokenizerMtd
}
func Parser_StreamTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_StreamTokenizer).FP
}
func Parser_StreamTokenizer_toSlice(slice []LnsAny) []*Parser_StreamTokenizer {
    ret := make([]*Parser_StreamTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_StreamTokenizerDownCast).ToParser_StreamTokenizer()
    }
    return ret
}
type Parser_StreamTokenizerDownCast interface {
    ToParser_StreamTokenizer() *Parser_StreamTokenizer
}
func Parser_StreamTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_StreamTokenizerDownCast)
    if ok { return work.ToParser_StreamTokenizer() }
    return nil
}
func (obj *Parser_StreamTokenizer) ToParser_StreamTokenizer() *Parser_StreamTokenizer {
    return obj
}
func NewParser_StreamTokenizer(_env *LnsEnv, arg1 LnsAny, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Parser_StreamTokenizer {
    obj := &Parser_StreamTokenizer{}
    obj.FP = obj
    obj.Parser_Tokenizer.FP = obj
    obj.InitParser_StreamTokenizer(_env, arg1, arg2, arg3, arg4)
    return obj
}
// 172: DeclConstr
func (self *Parser_StreamTokenizer) InitParser_StreamTokenizer(_env *LnsEnv, tokenizerSrc LnsAny,async bool,stdinFile LnsAny,pos LnsAny) {
    self.InitParser_Tokenizer(_env)
    self.pos = 1
    self.lineTokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.overridePos = pos
    var asyncTokenizer LnsAny
    var errMess string
    asyncTokenizer,errMess = AsyncParser_create(_env, tokenizerSrc, stdinFile, pos, async)
    {
        __exp := asyncTokenizer
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncParser_Tokenizer)
            self.asyncTokenizer = _exp
        } else {
            Util_err(_env, errMess)
        }
    }
    self.streamName = self.asyncTokenizer.FP.Get_streamName(_env)
}


// declaration Class -- DefaultPushbackTokenizer
type Parser_DefaultPushbackTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetLastPos(_env *LnsEnv) Types_Position
    GetNearCode(_env *LnsEnv) string
    GetStreamName(_env *LnsEnv) string
    GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
    GetUsedTokenListLen(_env *LnsEnv) LnsInt
    Get_currentToken(_env *LnsEnv) *Types_Token
    NewPushback(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
type Parser_DefaultPushbackTokenizer struct {
    tokenizer *Parser_Tokenizer
    pushbackedList *LnsList2_[*Types_Token]
    usedTokenList *LnsList2_[*Types_Token]
    currentToken *Types_Token
    FP Parser_DefaultPushbackTokenizerMtd
}
func Parser_DefaultPushbackTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_DefaultPushbackTokenizer).FP
}
func Parser_DefaultPushbackTokenizer_toSlice(slice []LnsAny) []*Parser_DefaultPushbackTokenizer {
    ret := make([]*Parser_DefaultPushbackTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_DefaultPushbackTokenizerDownCast).ToParser_DefaultPushbackTokenizer()
    }
    return ret
}
type Parser_DefaultPushbackTokenizerDownCast interface {
    ToParser_DefaultPushbackTokenizer() *Parser_DefaultPushbackTokenizer
}
func Parser_DefaultPushbackTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_DefaultPushbackTokenizerDownCast)
    if ok { return work.ToParser_DefaultPushbackTokenizer() }
    return nil
}
func (obj *Parser_DefaultPushbackTokenizer) ToParser_DefaultPushbackTokenizer() *Parser_DefaultPushbackTokenizer {
    return obj
}
func NewParser_DefaultPushbackTokenizer(_env *LnsEnv, arg1 *Parser_Tokenizer) *Parser_DefaultPushbackTokenizer {
    obj := &Parser_DefaultPushbackTokenizer{}
    obj.FP = obj
    obj.InitParser_DefaultPushbackTokenizer(_env, arg1)
    return obj
}
func (self *Parser_DefaultPushbackTokenizer) Get_currentToken(_env *LnsEnv) *Types_Token{ return self.currentToken }
// 232: DeclConstr
func (self *Parser_DefaultPushbackTokenizer) InitParser_DefaultPushbackTokenizer(_env *LnsEnv, tokenizer *Parser_Tokenizer) {
    self.tokenizer = tokenizer
    self.pushbackedList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.usedTokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.currentToken = Types_noneToken
}


// declaration Class -- DummyTokenizer
type Parser_DummyTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_DummyTokenizer struct {
    Parser_Tokenizer
    FP Parser_DummyTokenizerMtd
}
func Parser_DummyTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_DummyTokenizer).FP
}
func Parser_DummyTokenizer_toSlice(slice []LnsAny) []*Parser_DummyTokenizer {
    ret := make([]*Parser_DummyTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_DummyTokenizerDownCast).ToParser_DummyTokenizer()
    }
    return ret
}
type Parser_DummyTokenizerDownCast interface {
    ToParser_DummyTokenizer() *Parser_DummyTokenizer
}
func Parser_DummyTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_DummyTokenizerDownCast)
    if ok { return work.ToParser_DummyTokenizer() }
    return nil
}
func (obj *Parser_DummyTokenizer) ToParser_DummyTokenizer() *Parser_DummyTokenizer {
    return obj
}
func NewParser_DummyTokenizer(_env *LnsEnv) *Parser_DummyTokenizer {
    obj := &Parser_DummyTokenizer{}
    obj.FP = obj
    obj.Parser_Tokenizer.FP = obj
    obj.InitParser_DummyTokenizer(_env)
    return obj
}
func (self *Parser_DummyTokenizer) InitParser_DummyTokenizer(_env *LnsEnv) {
    self.Parser_Tokenizer.InitParser_Tokenizer( _env)
}

// declaration Class -- CommentLayer
type Parser_CommentLayerMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]
    HasInvalidComment(_env *LnsEnv) LnsAny
}
type Parser_CommentLayer struct {
    commentList *LnsList2_[*Types_Token]
    tokenSet *LnsSet2_[*Types_Token]
    tokenList *LnsList2_[*Types_Token]
    FP Parser_CommentLayerMtd
}
func Parser_CommentLayer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_CommentLayer).FP
}
func Parser_CommentLayer_toSlice(slice []LnsAny) []*Parser_CommentLayer {
    ret := make([]*Parser_CommentLayer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_CommentLayerDownCast).ToParser_CommentLayer()
    }
    return ret
}
type Parser_CommentLayerDownCast interface {
    ToParser_CommentLayer() *Parser_CommentLayer
}
func Parser_CommentLayerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_CommentLayerDownCast)
    if ok { return work.ToParser_CommentLayer() }
    return nil
}
func (obj *Parser_CommentLayer) ToParser_CommentLayer() *Parser_CommentLayer {
    return obj
}
func NewParser_CommentLayer(_env *LnsEnv) *Parser_CommentLayer {
    obj := &Parser_CommentLayer{}
    obj.FP = obj
    obj.InitParser_CommentLayer(_env)
    return obj
}
func (self *Parser_CommentLayer) Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]{ return self.commentList }
// 396: DeclConstr
func (self *Parser_CommentLayer) InitParser_CommentLayer(_env *LnsEnv) {
    self.commentList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.tokenSet = NewLnsSet2_[*Types_Token]([]*Types_Token{})
    self.tokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
}


// declaration Class -- CommentCtrl
type Parser_CommentCtrlMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]
    HasInvalidComment(_env *LnsEnv) LnsAny
    Pop(_env *LnsEnv)
    Push(_env *LnsEnv)
}
type Parser_CommentCtrl struct {
    layerStack *LnsList2_[*Parser_CommentLayer]
    layer *Parser_CommentLayer
    FP Parser_CommentCtrlMtd
}
func Parser_CommentCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_CommentCtrl).FP
}
func Parser_CommentCtrl_toSlice(slice []LnsAny) []*Parser_CommentCtrl {
    ret := make([]*Parser_CommentCtrl, len(slice))
    for index, val := range slice {
        ret[index] = val.(Parser_CommentCtrlDownCast).ToParser_CommentCtrl()
    }
    return ret
}
type Parser_CommentCtrlDownCast interface {
    ToParser_CommentCtrl() *Parser_CommentCtrl
}
func Parser_CommentCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_CommentCtrlDownCast)
    if ok { return work.ToParser_CommentCtrl() }
    return nil
}
func (obj *Parser_CommentCtrl) ToParser_CommentCtrl() *Parser_CommentCtrl {
    return obj
}
func NewParser_CommentCtrl(_env *LnsEnv) *Parser_CommentCtrl {
    obj := &Parser_CommentCtrl{}
    obj.FP = obj
    obj.InitParser_CommentCtrl(_env)
    return obj
}
// advertise -- 438
func (self *Parser_CommentCtrl) Add(_env *LnsEnv, arg1 *Types_Token) {
self.layer. FP.Add( _env, arg1)
}
// advertise -- 438
func (self *Parser_CommentCtrl) AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) {
self.layer. FP.AddDirect( _env, arg1)
}
// advertise -- 438
func (self *Parser_CommentCtrl) Clear(_env *LnsEnv) {
self.layer. FP.Clear( _env)
}
// advertise -- 438
func (self *Parser_CommentCtrl) Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token] {
    return self.layer. FP.Get_commentList( _env)
}
// advertise -- 438
func (self *Parser_CommentCtrl) HasInvalidComment(_env *LnsEnv) LnsAny {
    return self.layer. FP.HasInvalidComment( _env)
}
// 442: DeclConstr
func (self *Parser_CommentCtrl) InitParser_CommentCtrl(_env *LnsEnv) {
    self.layer = NewParser_CommentLayer(_env)
    self.layerStack = NewLnsList2_[*Parser_CommentLayer](Lns_2DDDGen[*Parser_CommentLayer](self.layer))
}


func Lns_Parser_init(_env *LnsEnv) {
    if init_Parser { return }
    init_Parser = true
    Parser__mod__ = "@lune.@base.@Parser"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Str_init(_env)
    Lns_Types_init(_env)
    Lns_Async_init(_env)
    Lns_AsyncParser_init(_env)
    Parser_noneToken = Types_noneToken
    Parser_StreamTokenizer____init_0_(_env)
    Parser_eofToken = NewTypes_Token(_env, Types_TokenKind__Eof, "<EOF>", NewTypes_Position(_env, 0, 0, "eof", nil), false, NewLnsList2_[*Types_Token]([]*Types_Token{}))
}
func init() {
    init_Parser = false
}
