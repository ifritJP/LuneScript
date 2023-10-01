// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Tokenizer bool
var Tokenizer__mod__ string
var Tokenizer_noneToken *Types_Token
var Tokenizer_eofToken *Types_Token
type Tokenizer_Pushbacked func (_env *LnsEnv, arg1 *Types_Token)
// for 91
func Tokenizer_convExp0_837(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 186
func Tokenizer_convExp0_319(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
type Tokenizer_TxtStream = Util_TxtStream
type Tokenizer_Position = Types_Position
type Tokenizer_TokenKind = Types_TokenKind
type Tokenizer_Token = Types_Token
// 34: decl @lune.@base.@Tokenizer.isLuaKeyword
func Tokenizer_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncTokenizer_isLuaKeyword(_env, txt)
}

// 73: decl @lune.@base.@Tokenizer.convFromRawToStr
func Tokenizer_convFromRawToStr(_env *LnsEnv, txt string) string {
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
            _endIndex := Tokenizer_convExp0_837(Lns_2DDD(_env.GetVM().String_find(workTxt, workPattern, workIndex, nil)))
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

// 369: decl @lune.@base.@Tokenizer.isOp2
func Tokenizer_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncTokenizer_isOp2(_env, ope)
}

// 373: decl @lune.@base.@Tokenizer.isOp1
func Tokenizer_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncTokenizer_isOp1(_env, ope)
}

// 382: decl @lune.@base.@Tokenizer.getEofToken
func Tokenizer_getEofToken(_env *LnsEnv) *Types_Token {
    return Tokenizer_eofToken
}

// 472: decl @lune.@base.@Tokenizer.quoteStr
func Tokenizer_quoteStr(_env *LnsEnv, txt string) string {
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

// 502: decl @lune.@base.@Tokenizer.createTokenizerFrom
func Tokenizer_createTokenizerFrom(_env *LnsEnv, src LnsAny,async bool,stdinFile LnsAny) *Tokenizer_Tokenizer {
    return &NewTokenizer_StreamTokenizer(_env, src, async, stdinFile, nil).Tokenizer_Tokenizer
}

// 135: decl @lune.@base.@Tokenizer.TokenListTokenizer.createPosition
func (self *Tokenizer_TokenListTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}
// 139: decl @lune.@base.@Tokenizer.TokenListTokenizer.getStreamName
func (self *Tokenizer_TokenListTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.streamName
}
// 143: decl @lune.@base.@Tokenizer.TokenListTokenizer.getToken
func (self *Tokenizer_TokenListTokenizer) GetToken(_env *LnsEnv) LnsAny {
    if self.tokenList.Len() < self.index{
        return nil
    }
    var token *Types_Token
    token = self.tokenList.GetAt(self.index)
    self.index = self.index + 1
    return token
}
// 164: decl @lune.@base.@Tokenizer.StreamTokenizer.setStdinStream
func Tokenizer_StreamTokenizer_setStdinStream(_env *LnsEnv, moduleName string) {
    Tokenizer_StreamTokenizer__stdinStreamModuleName = moduleName
    Tokenizer_StreamTokenizer__stdinTxt = Lns_unwrapDefault( Lns_io_stdin.Read(_env, "*a"), "").(string)
}
// 195: decl @lune.@base.@Tokenizer.StreamTokenizer.createPosition
func (self *Tokenizer_StreamTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}
// 200: decl @lune.@base.@Tokenizer.StreamTokenizer.getStreamName
func (self *Tokenizer_StreamTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.streamName
}
// 204: decl @lune.@base.@Tokenizer.StreamTokenizer.create
func Tokenizer_StreamTokenizer_create(_env *LnsEnv, tokenizerSrc LnsAny,async bool,stdinFile LnsAny,pos LnsAny) *Tokenizer_StreamTokenizer {
    return NewTokenizer_StreamTokenizer(_env, tokenizerSrc, async, stdinFile, pos)
}
// 211: decl @lune.@base.@Tokenizer.StreamTokenizer.getToken
func (self *Tokenizer_StreamTokenizer) GetToken(_env *LnsEnv) LnsAny {
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
            self.lineTokenList = pipeItem.FP.Get_item(_env).(*AsyncTokenizer_AsyncItem).List
        }
    }
    var token *Types_Token
    token = self.lineTokenList.GetAt(self.pos)
    self.pos = self.pos + 1
    return token
}
// 243: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.getUsedTokenListLen
func (self *Tokenizer_DefaultPushbackTokenizer) GetUsedTokenListLen(_env *LnsEnv) LnsInt {
    return self.usedTokenList.Len()
}
// 247: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.createFromLnsCode
func Tokenizer_DefaultPushbackTokenizer_createFromLnsCode(_env *LnsEnv, code string,name string) *Tokenizer_DefaultPushbackTokenizer {
    return NewTokenizer_DefaultPushbackTokenizer(_env, &NewTokenizer_StreamTokenizer(_env, &Types_TokenizerSrc__LnsCode{code, name, nil}, false, nil, nil).Tokenizer_Tokenizer)
}
// 254: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.createPosition
func (self *Tokenizer_DefaultPushbackTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return self.tokenizer.FP.CreatePosition(_env, lineNo, column)
}
// 258: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.getTokenNoErr
func (self *Tokenizer_DefaultPushbackTokenizer) GetTokenNoErr(_env *LnsEnv, skipFlag LnsAny) *Types_Token {
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
// 276: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.pushbackToken
func (self *Tokenizer_DefaultPushbackTokenizer) PushbackToken(_env *LnsEnv, token *Types_Token) {
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
// 299: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.pushback
func (self *Tokenizer_DefaultPushbackTokenizer) Pushback(_env *LnsEnv) {
    self.FP.PushbackToken(_env, self.currentToken)
}
// 302: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.pushbackStr
func (self *Tokenizer_DefaultPushbackTokenizer) PushbackStr(_env *LnsEnv, asyncParse LnsAny,name string,statement string,pos Types_Position,pushbacked LnsAny) {
    var tokenizer *Tokenizer_StreamTokenizer
    tokenizer = NewTokenizer_StreamTokenizer(_env, &Types_TokenizerSrc__LnsCode{statement, name, nil}, asyncParse == true, nil, pos)
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
    if pushbacked != nil{
        pushbacked_323 := pushbacked.(Tokenizer_Pushbacked)
        {
            var _forFrom1 LnsInt = list.Len()
            var _forTo1 LnsInt = 1
            _forWork1 := _forFrom1
            _forDelta1 := -1
            for {
                if _forDelta1 > 0 {
                   if _forWork1 > _forTo1 { break }
                } else {
                   if _forWork1 < _forTo1 { break }
                }
                index := _forWork1
                pushbacked_323(_env, list.GetAt(index))
                _forWork1 += _forDelta1
            }
        }
    }
}
// 327: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.newPushback
func (self *Tokenizer_DefaultPushbackTokenizer) NewPushback(_env *LnsEnv, tokenList *LnsList2_[*Types_Token]) {
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
// 333: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.getLastPos
func (self *Tokenizer_DefaultPushbackTokenizer) GetLastPos(_env *LnsEnv) Types_Position {
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
// 349: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.getNearCode
func (self *Tokenizer_DefaultPushbackTokenizer) GetNearCode(_env *LnsEnv) string {
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
// 364: decl @lune.@base.@Tokenizer.DefaultPushbackTokenizer.getStreamName
func (self *Tokenizer_DefaultPushbackTokenizer) GetStreamName(_env *LnsEnv) string {
    return self.tokenizer.FP.GetStreamName(_env)
}
// 386: decl @lune.@base.@Tokenizer.DummyTokenizer.getToken
func (self *Tokenizer_DummyTokenizer) GetToken(_env *LnsEnv) LnsAny {
    return Tokenizer_eofToken
}
// 389: decl @lune.@base.@Tokenizer.DummyTokenizer.getStreamName
func (self *Tokenizer_DummyTokenizer) GetStreamName(_env *LnsEnv) string {
    return "dummy"
}
// 392: decl @lune.@base.@Tokenizer.DummyTokenizer.createPosition
func (self *Tokenizer_DummyTokenizer) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), nil)
}
// 411: decl @lune.@base.@Tokenizer.CommentLayer.addDirect
func (self *Tokenizer_CommentLayer) AddDirect(_env *LnsEnv, commentList *LnsList2_[*Types_Token]) {
    for _, _comment := range( commentList.Items ) {
        comment := _comment
        self.commentList.Insert(comment)
    }
}
// 417: decl @lune.@base.@Tokenizer.CommentLayer.add
func (self *Tokenizer_CommentLayer) Add(_env *LnsEnv, token *Types_Token) {
    if Lns_op_not(self.tokenSet.Has(token)){
        self.tokenSet.Add(token)
        self.tokenList.Insert(token)
        self.FP.AddDirect(_env, token.FP.Get_commentList(_env))
    }
}
// 426: decl @lune.@base.@Tokenizer.CommentLayer.clear
func (self *Tokenizer_CommentLayer) Clear(_env *LnsEnv) {
    if self.commentList.Len() != 0{
        self.commentList = NewLnsList2_[*Types_Token]([]*Types_Token{})
        self.tokenSet = NewLnsSet2_[*Types_Token]([]*Types_Token{})
        self.tokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    }
}
// 442: decl @lune.@base.@Tokenizer.CommentLayer.hasInvalidComment
func (self *Tokenizer_CommentLayer) HasInvalidComment(_env *LnsEnv) LnsAny {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.tokenList.Len() > 1) &&
        _env.SetStackVal( self.tokenList.GetAt(2).FP.Get_commentList(_env).GetAt(1)) ||
        _env.SetStackVal( nil) )
}
// 456: decl @lune.@base.@Tokenizer.CommentCtrl.push
func (self *Tokenizer_CommentCtrl) Push(_env *LnsEnv) {
    self.layer = NewTokenizer_CommentLayer(_env)
    self.layerStack.Insert(self.layer)
}
// 461: decl @lune.@base.@Tokenizer.CommentCtrl.pop
func (self *Tokenizer_CommentCtrl) Pop(_env *LnsEnv) {
    self.layer = self.layerStack.GetAt(self.layerStack.Len())
    self.layerStack.Remove(nil)
}
// declaration Class -- Tokenizer
type Tokenizer_TokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Tokenizer_Tokenizer struct {
    FP Tokenizer_TokenizerMtd
}
func Tokenizer_Tokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_Tokenizer).FP
}
func Tokenizer_Tokenizer_toSlice(slice []LnsAny) []*Tokenizer_Tokenizer {
    ret := make([]*Tokenizer_Tokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_TokenizerDownCast).ToTokenizer_Tokenizer()
    }
    return ret
}
type Tokenizer_TokenizerDownCast interface {
    ToTokenizer_Tokenizer() *Tokenizer_Tokenizer
}
func Tokenizer_TokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_TokenizerDownCast)
    if ok { return work.ToTokenizer_Tokenizer() }
    return nil
}
func (obj *Tokenizer_Tokenizer) ToTokenizer_Tokenizer() *Tokenizer_Tokenizer {
    return obj
}
func (self *Tokenizer_Tokenizer) InitTokenizer_Tokenizer(_env *LnsEnv) {
}




type Tokenizer_PushbackTokenizer interface {
        CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
        GetStreamName(_env *LnsEnv) string
        GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
        NewPushback(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
        Pushback(_env *LnsEnv)
        PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position, arg5 LnsAny)
        PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
func Lns_cast2Tokenizer_PushbackTokenizer( obj LnsAny ) LnsAny {
    if _, ok := obj.(Tokenizer_PushbackTokenizer); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TokenListTokenizer
type Tokenizer_TokenListTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Tokenizer_TokenListTokenizer struct {
    Tokenizer_Tokenizer
    streamName string
    tokenList *LnsList2_[*Types_Token]
    index LnsInt
    overridePos LnsAny
    FP Tokenizer_TokenListTokenizerMtd
}
func Tokenizer_TokenListTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_TokenListTokenizer).FP
}
func Tokenizer_TokenListTokenizer_toSlice(slice []LnsAny) []*Tokenizer_TokenListTokenizer {
    ret := make([]*Tokenizer_TokenListTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_TokenListTokenizerDownCast).ToTokenizer_TokenListTokenizer()
    }
    return ret
}
type Tokenizer_TokenListTokenizerDownCast interface {
    ToTokenizer_TokenListTokenizer() *Tokenizer_TokenListTokenizer
}
func Tokenizer_TokenListTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_TokenListTokenizerDownCast)
    if ok { return work.ToTokenizer_TokenListTokenizer() }
    return nil
}
func (obj *Tokenizer_TokenListTokenizer) ToTokenizer_TokenListTokenizer() *Tokenizer_TokenListTokenizer {
    return obj
}
func NewTokenizer_TokenListTokenizer(_env *LnsEnv, arg1 *LnsList2_[*Types_Token], arg2 string, arg3 LnsAny) *Tokenizer_TokenListTokenizer {
    obj := &Tokenizer_TokenListTokenizer{}
    obj.FP = obj
    obj.Tokenizer_Tokenizer.FP = obj
    obj.InitTokenizer_TokenListTokenizer(_env, arg1, arg2, arg3)
    return obj
}
// 124: DeclConstr
func (self *Tokenizer_TokenListTokenizer) InitTokenizer_TokenListTokenizer(_env *LnsEnv, tokenList *LnsList2_[*Types_Token],streamName string,overridePos LnsAny) {
    self.InitTokenizer_Tokenizer(_env)
    self.index = 1
    self.tokenList = tokenList
    self.streamName = streamName
    self.overridePos = overridePos
}


// declaration Class -- StreamTokenizer
var Tokenizer_StreamTokenizer__stdinStreamModuleName LnsAny
var Tokenizer_StreamTokenizer__stdinTxt string
// 155: decl @lune.@base.@Tokenizer.StreamTokenizer.___init
func Tokenizer_StreamTokenizer____init_0_(_env *LnsEnv) {
    Tokenizer_StreamTokenizer__stdinStreamModuleName = nil
    Tokenizer_StreamTokenizer__stdinTxt = ""
}

type Tokenizer_StreamTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Tokenizer_StreamTokenizer struct {
    Tokenizer_Tokenizer
    stdinStreamModuleName LnsAny
    stdinTxt string
    streamName string
    pos LnsInt
    lineTokenList *LnsList2_[*Types_Token]
    asyncTokenizer *AsyncTokenizer_Tokenizer
    overridePos LnsAny
    FP Tokenizer_StreamTokenizerMtd
}
func Tokenizer_StreamTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_StreamTokenizer).FP
}
func Tokenizer_StreamTokenizer_toSlice(slice []LnsAny) []*Tokenizer_StreamTokenizer {
    ret := make([]*Tokenizer_StreamTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_StreamTokenizerDownCast).ToTokenizer_StreamTokenizer()
    }
    return ret
}
type Tokenizer_StreamTokenizerDownCast interface {
    ToTokenizer_StreamTokenizer() *Tokenizer_StreamTokenizer
}
func Tokenizer_StreamTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_StreamTokenizerDownCast)
    if ok { return work.ToTokenizer_StreamTokenizer() }
    return nil
}
func (obj *Tokenizer_StreamTokenizer) ToTokenizer_StreamTokenizer() *Tokenizer_StreamTokenizer {
    return obj
}
func NewTokenizer_StreamTokenizer(_env *LnsEnv, arg1 LnsAny, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Tokenizer_StreamTokenizer {
    obj := &Tokenizer_StreamTokenizer{}
    obj.FP = obj
    obj.Tokenizer_Tokenizer.FP = obj
    obj.InitTokenizer_StreamTokenizer(_env, arg1, arg2, arg3, arg4)
    return obj
}
// 176: DeclConstr
func (self *Tokenizer_StreamTokenizer) InitTokenizer_StreamTokenizer(_env *LnsEnv, tokenizerSrc LnsAny,async bool,stdinFile LnsAny,pos LnsAny) {
    self.InitTokenizer_Tokenizer(_env)
    self.pos = 1
    self.lineTokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.overridePos = pos
    var asyncTokenizer LnsAny
    var errMess string
    asyncTokenizer,errMess = AsyncTokenizer_create(_env, tokenizerSrc, stdinFile, pos, async)
    {
        __exp := asyncTokenizer
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncTokenizer_Tokenizer)
            self.asyncTokenizer = _exp
        } else {
            Util_err(_env, errMess)
        }
    }
    self.streamName = self.asyncTokenizer.FP.Get_streamName(_env)
}


// declaration Class -- DefaultPushbackTokenizer
type Tokenizer_DefaultPushbackTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetLastPos(_env *LnsEnv) Types_Position
    GetNearCode(_env *LnsEnv) string
    GetStreamName(_env *LnsEnv) string
    GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
    GetUsedTokenListLen(_env *LnsEnv) LnsInt
    Get_currentToken(_env *LnsEnv) *Types_Token
    NewPushback(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position, arg5 LnsAny)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
type Tokenizer_DefaultPushbackTokenizer struct {
    tokenizer *Tokenizer_Tokenizer
    pushbackedList *LnsList2_[*Types_Token]
    usedTokenList *LnsList2_[*Types_Token]
    currentToken *Types_Token
    FP Tokenizer_DefaultPushbackTokenizerMtd
}
func Tokenizer_DefaultPushbackTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_DefaultPushbackTokenizer).FP
}
func Tokenizer_DefaultPushbackTokenizer_toSlice(slice []LnsAny) []*Tokenizer_DefaultPushbackTokenizer {
    ret := make([]*Tokenizer_DefaultPushbackTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_DefaultPushbackTokenizerDownCast).ToTokenizer_DefaultPushbackTokenizer()
    }
    return ret
}
type Tokenizer_DefaultPushbackTokenizerDownCast interface {
    ToTokenizer_DefaultPushbackTokenizer() *Tokenizer_DefaultPushbackTokenizer
}
func Tokenizer_DefaultPushbackTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_DefaultPushbackTokenizerDownCast)
    if ok { return work.ToTokenizer_DefaultPushbackTokenizer() }
    return nil
}
func (obj *Tokenizer_DefaultPushbackTokenizer) ToTokenizer_DefaultPushbackTokenizer() *Tokenizer_DefaultPushbackTokenizer {
    return obj
}
func NewTokenizer_DefaultPushbackTokenizer(_env *LnsEnv, arg1 *Tokenizer_Tokenizer) *Tokenizer_DefaultPushbackTokenizer {
    obj := &Tokenizer_DefaultPushbackTokenizer{}
    obj.FP = obj
    obj.InitTokenizer_DefaultPushbackTokenizer(_env, arg1)
    return obj
}
func (self *Tokenizer_DefaultPushbackTokenizer) Get_currentToken(_env *LnsEnv) *Types_Token{ return self.currentToken }
// 236: DeclConstr
func (self *Tokenizer_DefaultPushbackTokenizer) InitTokenizer_DefaultPushbackTokenizer(_env *LnsEnv, tokenizer *Tokenizer_Tokenizer) {
    self.tokenizer = tokenizer
    self.pushbackedList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.usedTokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.currentToken = Types_noneToken
}


// declaration Class -- DummyTokenizer
type Tokenizer_DummyTokenizerMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Tokenizer_DummyTokenizer struct {
    Tokenizer_Tokenizer
    FP Tokenizer_DummyTokenizerMtd
}
func Tokenizer_DummyTokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_DummyTokenizer).FP
}
func Tokenizer_DummyTokenizer_toSlice(slice []LnsAny) []*Tokenizer_DummyTokenizer {
    ret := make([]*Tokenizer_DummyTokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_DummyTokenizerDownCast).ToTokenizer_DummyTokenizer()
    }
    return ret
}
type Tokenizer_DummyTokenizerDownCast interface {
    ToTokenizer_DummyTokenizer() *Tokenizer_DummyTokenizer
}
func Tokenizer_DummyTokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_DummyTokenizerDownCast)
    if ok { return work.ToTokenizer_DummyTokenizer() }
    return nil
}
func (obj *Tokenizer_DummyTokenizer) ToTokenizer_DummyTokenizer() *Tokenizer_DummyTokenizer {
    return obj
}
func NewTokenizer_DummyTokenizer(_env *LnsEnv) *Tokenizer_DummyTokenizer {
    obj := &Tokenizer_DummyTokenizer{}
    obj.FP = obj
    obj.Tokenizer_Tokenizer.FP = obj
    obj.InitTokenizer_DummyTokenizer(_env)
    return obj
}
func (self *Tokenizer_DummyTokenizer) InitTokenizer_DummyTokenizer(_env *LnsEnv) {
    self.Tokenizer_Tokenizer.InitTokenizer_Tokenizer( _env)
}

// declaration Class -- CommentLayer
type Tokenizer_CommentLayerMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]
    HasInvalidComment(_env *LnsEnv) LnsAny
}
type Tokenizer_CommentLayer struct {
    commentList *LnsList2_[*Types_Token]
    tokenSet *LnsSet2_[*Types_Token]
    tokenList *LnsList2_[*Types_Token]
    FP Tokenizer_CommentLayerMtd
}
func Tokenizer_CommentLayer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_CommentLayer).FP
}
func Tokenizer_CommentLayer_toSlice(slice []LnsAny) []*Tokenizer_CommentLayer {
    ret := make([]*Tokenizer_CommentLayer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_CommentLayerDownCast).ToTokenizer_CommentLayer()
    }
    return ret
}
type Tokenizer_CommentLayerDownCast interface {
    ToTokenizer_CommentLayer() *Tokenizer_CommentLayer
}
func Tokenizer_CommentLayerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_CommentLayerDownCast)
    if ok { return work.ToTokenizer_CommentLayer() }
    return nil
}
func (obj *Tokenizer_CommentLayer) ToTokenizer_CommentLayer() *Tokenizer_CommentLayer {
    return obj
}
func NewTokenizer_CommentLayer(_env *LnsEnv) *Tokenizer_CommentLayer {
    obj := &Tokenizer_CommentLayer{}
    obj.FP = obj
    obj.InitTokenizer_CommentLayer(_env)
    return obj
}
func (self *Tokenizer_CommentLayer) Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]{ return self.commentList }
// 405: DeclConstr
func (self *Tokenizer_CommentLayer) InitTokenizer_CommentLayer(_env *LnsEnv) {
    self.commentList = NewLnsList2_[*Types_Token]([]*Types_Token{})
    self.tokenSet = NewLnsSet2_[*Types_Token]([]*Types_Token{})
    self.tokenList = NewLnsList2_[*Types_Token]([]*Types_Token{})
}


// declaration Class -- CommentCtrl
type Tokenizer_CommentCtrlMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token])
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]
    HasInvalidComment(_env *LnsEnv) LnsAny
    Pop(_env *LnsEnv)
    Push(_env *LnsEnv)
}
type Tokenizer_CommentCtrl struct {
    layerStack *LnsList2_[*Tokenizer_CommentLayer]
    layer *Tokenizer_CommentLayer
    FP Tokenizer_CommentCtrlMtd
}
func Tokenizer_CommentCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Tokenizer_CommentCtrl).FP
}
func Tokenizer_CommentCtrl_toSlice(slice []LnsAny) []*Tokenizer_CommentCtrl {
    ret := make([]*Tokenizer_CommentCtrl, len(slice))
    for index, val := range slice {
        ret[index] = val.(Tokenizer_CommentCtrlDownCast).ToTokenizer_CommentCtrl()
    }
    return ret
}
type Tokenizer_CommentCtrlDownCast interface {
    ToTokenizer_CommentCtrl() *Tokenizer_CommentCtrl
}
func Tokenizer_CommentCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Tokenizer_CommentCtrlDownCast)
    if ok { return work.ToTokenizer_CommentCtrl() }
    return nil
}
func (obj *Tokenizer_CommentCtrl) ToTokenizer_CommentCtrl() *Tokenizer_CommentCtrl {
    return obj
}
func NewTokenizer_CommentCtrl(_env *LnsEnv) *Tokenizer_CommentCtrl {
    obj := &Tokenizer_CommentCtrl{}
    obj.FP = obj
    obj.InitTokenizer_CommentCtrl(_env)
    return obj
}
// advertise -- 447
func (self *Tokenizer_CommentCtrl) Add(_env *LnsEnv, arg1 *Types_Token) {
self.layer. FP.Add( _env, arg1)
}
// advertise -- 447
func (self *Tokenizer_CommentCtrl) AddDirect(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) {
self.layer. FP.AddDirect( _env, arg1)
}
// advertise -- 447
func (self *Tokenizer_CommentCtrl) Clear(_env *LnsEnv) {
self.layer. FP.Clear( _env)
}
// advertise -- 447
func (self *Tokenizer_CommentCtrl) Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token] {
    return self.layer. FP.Get_commentList( _env)
}
// advertise -- 447
func (self *Tokenizer_CommentCtrl) HasInvalidComment(_env *LnsEnv) LnsAny {
    return self.layer. FP.HasInvalidComment( _env)
}
// 451: DeclConstr
func (self *Tokenizer_CommentCtrl) InitTokenizer_CommentCtrl(_env *LnsEnv) {
    self.layer = NewTokenizer_CommentLayer(_env)
    self.layerStack = NewLnsList2_[*Tokenizer_CommentLayer](Lns_2DDDGen[*Tokenizer_CommentLayer](self.layer))
}


func Lns_Tokenizer_init(_env *LnsEnv) {
    if init_Tokenizer { return }
    init_Tokenizer = true
    Tokenizer__mod__ = "@lune.@base.@Tokenizer"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Str_init(_env)
    Lns_Types_init(_env)
    Lns_Async_init(_env)
    Lns_AsyncTokenizer_init(_env)
    Tokenizer_noneToken = Types_noneToken
    Tokenizer_StreamTokenizer____init_0_(_env)
    Tokenizer_eofToken = NewTypes_Token(_env, Types_TokenKind__Eof, "<EOF>", NewTypes_Position(_env, 0, 0, "eof", nil), false, NewLnsList2_[*Types_Token]([]*Types_Token{}))
}
func init() {
    init_Tokenizer = false
}
