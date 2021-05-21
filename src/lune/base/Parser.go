// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Parser bool
var Parser__mod__ string
var Parser_noneToken *Types_Token
var Parser_eofToken *Types_Token
// for 141
func Parser_convExp473(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 213
func Parser_convExp848(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
type Parser_Position = Types_Position
type Parser_TokenKind = Types_TokenKind
type Parser_Token = Types_Token
// 31: decl @lune.@base.@Parser.isLuaKeyword
func Parser_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncParser_isLuaKeyword(_env, txt)
}

// 123: decl @lune.@base.@Parser.convFromRawToStr
func Parser_convFromRawToStr(_env *LnsEnv, txt string) string {
    if len(txt) == 0{
        return txt
    }
    if _switch414 := LnsInt(txt[1-1]); _switch414 == 39 || _switch414 == 34 {
    } else {
        return _env.LuaVM.String_sub(txt,4, len(txt) - 3)
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
            _endIndex := Parser_convExp473(Lns_2DDD(_env.LuaVM.String_find(workTxt, workPattern, workIndex, nil)))
            if _endIndex == nil{
                Util_err(_env, _env.LuaVM.String_format("error: illegal string -- %s", []LnsAny{workTxt}))
            } else {
                endIndex = _endIndex.(LnsInt)
            }
        }
        var workChar LnsInt
        workChar = LnsInt(workTxt[endIndex-1])
        if workChar == findChar{
            return retTxt + _env.LuaVM.String_sub(workTxt,setIndex, endIndex - 1)
        } else if workChar == 92{
            var quote LnsInt
            quote = LnsInt(workTxt[endIndex + 1-1])
            if _switch580 := quote; _switch580 == 39 || _switch580 == 34 {
                retTxt = _env.LuaVM.String_format("%s%s%c", []LnsAny{retTxt, _env.LuaVM.String_sub(workTxt,setIndex, endIndex - 1), quote})
                
            } else {
                retTxt = _env.LuaVM.String_format("%s%s", []LnsAny{retTxt, _env.LuaVM.String_sub(workTxt,setIndex, endIndex + 1)})
                
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

// 368: decl @lune.@base.@Parser.isOp2
func Parser_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncParser_isOp2(_env, ope)
}

// 372: decl @lune.@base.@Parser.isOp1
func Parser_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncParser_isOp1(_env, ope)
}

// 381: decl @lune.@base.@Parser.getEofToken
func Parser_getEofToken(_env *LnsEnv) *Types_Token {
    return Parser_eofToken
}

// 471: decl @lune.@base.@Parser.quoteStr
func Parser_quoteStr(_env *LnsEnv, txt string) string {
    var work string
    work = txt
    var part string
    part = "\""
    {
        var _from2151 LnsInt = 1
        var _to2151 LnsInt = len(work)
        for _work2151 := _from2151; _work2151 <= _to2151; _work2151++ {
            index := _work2151
            var char LnsInt
            char = LnsInt(work[index-1])
            if _switch2149 := char; _switch2149 == 10 {
                part = part + "\\n"
                
            } else if _switch2149 == 13 {
                part = part + "\\r"
                
            } else if _switch2149 == 9 {
                part = part + "\\t"
                
            } else if _switch2149 == 34 {
                part = part + "\\\""
                
            } else if _switch2149 == 92 {
                part = part + "\\\\"
                
            } else {
                part = part + _env.LuaVM.String_format("%c", []LnsAny{char})
                
            }
        }
    }
    work = part + "\""
    
    return work
}

// declaration Class -- TxtStream
type Parser_TxtStreamMtd interface {
    Close(_env *LnsEnv)
    GetSubstring(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) string
    Get_lineNo(_env *LnsEnv) LnsInt
    Get_txt(_env *LnsEnv) string
    Read(_env *LnsEnv, arg1 LnsAny) LnsAny
}
type Parser_TxtStream struct {
    txt string
    lineList *LnsList
    start LnsInt
    eof bool
    lineNo LnsInt
    FP Parser_TxtStreamMtd
}
func Parser_TxtStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_TxtStream).FP
}
type Parser_TxtStreamDownCast interface {
    ToParser_TxtStream() *Parser_TxtStream
}
func Parser_TxtStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_TxtStreamDownCast)
    if ok { return work.ToParser_TxtStream() }
    return nil
}
func (obj *Parser_TxtStream) ToParser_TxtStream() *Parser_TxtStream {
    return obj
}
func NewParser_TxtStream(_env *LnsEnv, arg1 string) *Parser_TxtStream {
    obj := &Parser_TxtStream{}
    obj.FP = obj
    obj.InitParser_TxtStream(_env, arg1)
    return obj
}
func (self *Parser_TxtStream) Get_txt(_env *LnsEnv) string{ return self.txt }
func (self *Parser_TxtStream) Get_lineNo(_env *LnsEnv) LnsInt{ return self.lineNo }
// 42: DeclConstr
func (self *Parser_TxtStream) InitParser_TxtStream(_env *LnsEnv, txt string) {
    self.txt = txt
    
    self.start = 1
    
    self.eof = false
    
    self.lineList = Str_getLineList(_env, self.txt)
    
    self.lineNo = 1
    
}

// 61: decl @lune.@base.@Parser.TxtStream.getSubstring
func (self *Parser_TxtStream) GetSubstring(_env *LnsEnv, fromLineNo LnsInt,toLineNo LnsAny) string {
    var txt string
    txt = ""
    var to LnsInt
    to = Lns_unwrapDefault( toLineNo, self.lineList.Len() + 1).(LnsInt)
    {
        var _from167 LnsInt = fromLineNo
        var _to167 LnsInt = to - 1
        for _work167 := _from167; _work167 <= _to167; _work167++ {
            index := _work167
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index < 1) ||
                _env.SetStackVal( index > self.lineList.Len()) ).(bool){
                break
            }
            txt = _env.LuaVM.String_format("%s%s", []LnsAny{txt, self.lineList.GetAt(index).(string)})
            
        }
    }
    return txt
}

// 73: decl @lune.@base.@Parser.TxtStream.read
func (self *Parser_TxtStream) Read(_env *LnsEnv, mode LnsAny) LnsAny {
    if mode != "*l"{
        Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{mode}))
    }
    if self.lineNo > self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    
    var line string
    line = self.lineList.GetAt(self.lineNo - 1).(string)
    if Str_endsWith(_env, line, "\n"){
        return _env.LuaVM.String_sub(line,1, len(line) - 1)
    }
    return line
}

// 87: decl @lune.@base.@Parser.TxtStream.close
func (self *Parser_TxtStream) Close(_env *LnsEnv) {
}


// declaration Class -- Parser
type Parser_ParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_Parser struct {
    FP Parser_ParserMtd
}
func Parser_Parser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_Parser).FP
}
type Parser_ParserDownCast interface {
    ToParser_Parser() *Parser_Parser
}
func Parser_ParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_ParserDownCast)
    if ok { return work.ToParser_Parser() }
    return nil
}
func (obj *Parser_Parser) ToParser_Parser() *Parser_Parser {
    return obj
}
func (self *Parser_Parser) InitParser_Parser(_env *LnsEnv) {
}




type Parser_PushbackParser interface {
        CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
        Error(_env *LnsEnv, arg1 string)
        GetStreamName(_env *LnsEnv) string
        GetTokenNoErr(_env *LnsEnv) *Types_Token
        NewPushback(_env *LnsEnv, arg1 *LnsList)
        Pushback(_env *LnsEnv)
        PushbackStr(_env *LnsEnv, arg1 string, arg2 string, arg3 *Types_Position)
        PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
func Lns_cast2Parser_PushbackParser( obj LnsAny ) LnsAny {
    if _, ok := obj.(Parser_PushbackParser); ok { 
        return obj
    }
    return nil
}

// declaration Class -- StreamParser
var Parser_StreamParser__stdinStreamModuleName LnsAny
var Parser_StreamParser__stdinTxt string
// 168: decl @lune.@base.@Parser.StreamParser.___init
func Parser_StreamParser____init_1084_(_env *LnsEnv) {
    Parser_StreamParser__stdinStreamModuleName = nil
    
    Parser_StreamParser__stdinTxt = ""
    
}

type Parser_StreamParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_StreamParser struct {
    Parser_Parser
    stdinStreamModuleName LnsAny
    stdinTxt string
    streamName string
    pos LnsInt
    lineTokenList *LnsList
    asyncParser *AsyncParser_Parser
    overridePos LnsAny
    FP Parser_StreamParserMtd
}
func Parser_StreamParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_StreamParser).FP
}
type Parser_StreamParserDownCast interface {
    ToParser_StreamParser() *Parser_StreamParser
}
func Parser_StreamParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_StreamParserDownCast)
    if ok { return work.ToParser_StreamParser() }
    return nil
}
func (obj *Parser_StreamParser) ToParser_StreamParser() *Parser_StreamParser {
    return obj
}
func NewParser_StreamParser(_env *LnsEnv, arg1 Lns_iStream, arg2 string, arg3 LnsAny, arg4 LnsAny) *Parser_StreamParser {
    obj := &Parser_StreamParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_StreamParser(_env, arg1, arg2, arg3, arg4)
    return obj
}
// 177: decl @lune.@base.@Parser.StreamParser.setStdinStream
func Parser_StreamParser_setStdinStream(_env *LnsEnv, moduleName string) {
    Parser_StreamParser__stdinStreamModuleName = moduleName
    
    Parser_StreamParser__stdinTxt = Lns_unwrapDefault( Lns_io_stdin.Read(_env, "*a"), "").(string)
    
}

// 189: DeclConstr
func (self *Parser_StreamParser) InitParser_StreamParser(_env *LnsEnv, stream Lns_iStream,name string,luaMode LnsAny,pos LnsAny) {
    self.InitParser_Parser(_env)
    self.streamName = name
    
    self.pos = 1
    
    self.lineTokenList = NewLnsList([]LnsAny{})
    
    self.overridePos = pos
    
    self.asyncParser = NewAsyncParser_Parser(_env, stream, name, luaMode, pos)
    
}

// 201: decl @lune.@base.@Parser.StreamParser.createPosition
func (self *Parser_StreamParser) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) *Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}

// 206: decl @lune.@base.@Parser.StreamParser.getStreamName
func (self *Parser_StreamParser) GetStreamName(_env *LnsEnv) string {
    return self.streamName
}

// 210: decl @lune.@base.@Parser.StreamParser.create
func Parser_StreamParser_create(_env *LnsEnv, path string,luaMode bool,moduleName string) LnsAny {
    var stream Lns_iStream
    if Parser_StreamParser__stdinStreamModuleName != moduleName{
        {
            var _stream LnsAny
            _stream = Parser_convExp848(Lns_2DDD(Lns_io_open(path, "r")))
            if Lns_IsNil( _stream ) {
                return nil
            }
            stream = _stream.(Lns_iStream)
        }
    } else { 
        stream = NewParser_TxtStream(_env, Parser_StreamParser__stdinTxt).FP
        
    }
    return NewParser_StreamParser(_env, stream, path, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( luaMode) ||
        _env.SetStackVal( Str_endsWith(_env, path, ".lua")) &&
        _env.SetStackVal( true) ).(bool), nil)
}

// 224: decl @lune.@base.@Parser.StreamParser.getToken
func (self *Parser_StreamParser) GetToken(_env *LnsEnv) LnsAny {
    if self.lineTokenList.Len() < self.pos{
        self.pos = 1
        
        self.lineTokenList = NewLnsList([]LnsAny{})
        
        for self.lineTokenList.Len() == 0 {
            var pipeItem *Async_PipeItem
            
            {
                _pipeItem := self.asyncParser.FP.GetNext(_env)
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
    token = self.lineTokenList.GetAt(self.pos).(Types_TokenDownCast).ToTypes_Token()
    self.pos = self.pos + 1
    
    return token
}


// declaration Class -- DefaultPushbackParser
type Parser_DefaultPushbackParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
    Error(_env *LnsEnv, arg1 string)
    GetLastPos(_env *LnsEnv) *Types_Position
    GetNearCode(_env *LnsEnv) string
    GetStreamName(_env *LnsEnv) string
    GetTokenNoErr(_env *LnsEnv) *Types_Token
    Get_currentToken(_env *LnsEnv) *Types_Token
    NewPushback(_env *LnsEnv, arg1 *LnsList)
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 string, arg2 string, arg3 *Types_Position)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
}
type Parser_DefaultPushbackParser struct {
    parser *Parser_Parser
    pushbackedList *LnsList
    usedTokenList *LnsList
    currentToken *Types_Token
    FP Parser_DefaultPushbackParserMtd
}
func Parser_DefaultPushbackParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_DefaultPushbackParser).FP
}
type Parser_DefaultPushbackParserDownCast interface {
    ToParser_DefaultPushbackParser() *Parser_DefaultPushbackParser
}
func Parser_DefaultPushbackParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_DefaultPushbackParserDownCast)
    if ok { return work.ToParser_DefaultPushbackParser() }
    return nil
}
func (obj *Parser_DefaultPushbackParser) ToParser_DefaultPushbackParser() *Parser_DefaultPushbackParser {
    return obj
}
func NewParser_DefaultPushbackParser(_env *LnsEnv, arg1 *Parser_Parser) *Parser_DefaultPushbackParser {
    obj := &Parser_DefaultPushbackParser{}
    obj.FP = obj
    obj.InitParser_DefaultPushbackParser(_env, arg1)
    return obj
}
func (self *Parser_DefaultPushbackParser) Get_currentToken(_env *LnsEnv) *Types_Token{ return self.currentToken }
// 250: DeclConstr
func (self *Parser_DefaultPushbackParser) InitParser_DefaultPushbackParser(_env *LnsEnv, parser *Parser_Parser) {
    self.parser = parser
    
    self.pushbackedList = NewLnsList([]LnsAny{})
    
    self.usedTokenList = NewLnsList([]LnsAny{})
    
    self.currentToken = Types_noneToken
    
}

// 257: decl @lune.@base.@Parser.DefaultPushbackParser.createPosition
func (self *Parser_DefaultPushbackParser) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) *Types_Position {
    return self.parser.FP.CreatePosition(_env, lineNo, column)
}

// 261: decl @lune.@base.@Parser.DefaultPushbackParser.getTokenNoErr
func (self *Parser_DefaultPushbackParser) GetTokenNoErr(_env *LnsEnv) *Types_Token {
    if self.pushbackedList.Len() > 0{
        self.currentToken = self.pushbackedList.GetAt(self.pushbackedList.Len()).(Types_TokenDownCast).ToTypes_Token()
        
        self.pushbackedList.Remove(nil)
    } else { 
        {
            _token := self.parser.FP.GetToken(_env)
            if !Lns_IsNil( _token ) {
                token := _token.(*Types_Token)
                self.currentToken = token
                
            } else {
                self.currentToken = Types_noneToken
                
            }
        }
    }
    if self.currentToken.Kind != Types_TokenKind__Eof{
        self.usedTokenList.Insert(Types_Token2Stem(self.currentToken))
    }
    return self.currentToken
}

// 279: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackToken
func (self *Parser_DefaultPushbackParser) PushbackToken(_env *LnsEnv, token *Types_Token) {
    if token.Kind != Types_TokenKind__Eof{
        self.pushbackedList.Insert(Types_Token2Stem(token))
    }
    if token == self.currentToken{
        if self.usedTokenList.Len() > 0{
            var used *Types_Token
            used = self.usedTokenList.GetAt(self.usedTokenList.Len()).(Types_TokenDownCast).ToTypes_Token()
            if used == token{
                self.usedTokenList.Remove(nil)
            }
            if self.usedTokenList.Len() > 0{
                self.currentToken = self.usedTokenList.GetAt(self.usedTokenList.Len()).(Types_TokenDownCast).ToTypes_Token()
                
            } else { 
                self.currentToken = Types_noneToken
                
            }
        } else { 
            self.currentToken = Types_noneToken
            
        }
    }
}

// 302: decl @lune.@base.@Parser.DefaultPushbackParser.pushback
func (self *Parser_DefaultPushbackParser) Pushback(_env *LnsEnv) {
    self.FP.PushbackToken(_env, self.currentToken)
}

// 305: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackStr
func (self *Parser_DefaultPushbackParser) PushbackStr(_env *LnsEnv, name string,statement string,pos *Types_Position) {
    var memStream *Parser_TxtStream
    memStream = NewParser_TxtStream(_env, statement)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(_env, memStream.FP, name, false, pos)
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        {
            __exp := parser.FP.GetToken(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Types_Token)
                list.Insert(Types_Token2Stem(_exp))
            } else {
                break
            }
        }
    }
    {
        var _from1392 LnsInt = list.Len()
        var _to1392 LnsInt = 1
        _work1392 := _from1392
        _delta1392 := -1
        for {
            if _delta1392 > 0 {
               if _work1392 > _to1392 { break }
            } else {
               if _work1392 < _to1392 { break }
            }
            index := _work1392
            self.FP.PushbackToken(_env, list.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work1392 += _delta1392
        }
    }
}

// 322: decl @lune.@base.@Parser.DefaultPushbackParser.newPushback
func (self *Parser_DefaultPushbackParser) NewPushback(_env *LnsEnv, tokenList *LnsList) {
    {
        var _from1421 LnsInt = tokenList.Len()
        var _to1421 LnsInt = 1
        _work1421 := _from1421
        _delta1421 := -1
        for {
            if _delta1421 > 0 {
               if _work1421 > _to1421 { break }
            } else {
               if _work1421 < _to1421 { break }
            }
            index := _work1421
            self.FP.PushbackToken(_env, tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work1421 += _delta1421
        }
    }
}

// 327: decl @lune.@base.@Parser.DefaultPushbackParser.error
func (self *Parser_DefaultPushbackParser) Error(_env *LnsEnv, message string) {
    Util_err(_env, message)
}

// 332: decl @lune.@base.@Parser.DefaultPushbackParser.getLastPos
func (self *Parser_DefaultPushbackParser) GetLastPos(_env *LnsEnv) *Types_Position {
    var pos *Types_Position
    pos = self.parser.FP.CreatePosition(_env, 0, 0)
    if self.FP.Get_currentToken(_env).Kind != Types_TokenKind__Eof{
        pos = self.FP.Get_currentToken(_env).Pos
        
    } else { 
        if self.usedTokenList.Len() > 0{
            var token *Types_Token
            token = self.usedTokenList.GetAt(self.usedTokenList.Len()).(Types_TokenDownCast).ToTypes_Token()
            pos = token.Pos
            
        }
    }
    return pos
}

// 348: decl @lune.@base.@Parser.DefaultPushbackParser.getNearCode
func (self *Parser_DefaultPushbackParser) GetNearCode(_env *LnsEnv) string {
    var code string
    code = ""
    {
        var _from1596 LnsInt = self.usedTokenList.Len() - 30
        var _to1596 LnsInt = self.usedTokenList.Len()
        for _work1596 := _from1596; _work1596 <= _to1596; _work1596++ {
            index := _work1596
            if index > 1{
                var token *Types_Token
                token = self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token()
                if token.Consecutive{
                    code = _env.LuaVM.String_format("%s%s", []LnsAny{code, self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token().Txt})
                    
                } else { 
                    code = _env.LuaVM.String_format("%s %s", []LnsAny{code, self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token().Txt})
                    
                }
            }
        }
    }
    return _env.LuaVM.String_format("%s -- current '%s'", []LnsAny{code, self.currentToken.Txt})
}

// 363: decl @lune.@base.@Parser.DefaultPushbackParser.getStreamName
func (self *Parser_DefaultPushbackParser) GetStreamName(_env *LnsEnv) string {
    return self.parser.FP.GetStreamName(_env)
}


// declaration Class -- DummyParser
type Parser_DummyParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Parser_DummyParser struct {
    Parser_Parser
    FP Parser_DummyParserMtd
}
func Parser_DummyParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_DummyParser).FP
}
type Parser_DummyParserDownCast interface {
    ToParser_DummyParser() *Parser_DummyParser
}
func Parser_DummyParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_DummyParserDownCast)
    if ok { return work.ToParser_DummyParser() }
    return nil
}
func (obj *Parser_DummyParser) ToParser_DummyParser() *Parser_DummyParser {
    return obj
}
func NewParser_DummyParser(_env *LnsEnv) *Parser_DummyParser {
    obj := &Parser_DummyParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_DummyParser(_env)
    return obj
}
func (self *Parser_DummyParser) InitParser_DummyParser(_env *LnsEnv) {
    self.Parser_Parser.InitParser_Parser( _env)
}
// 385: decl @lune.@base.@Parser.DummyParser.getToken
func (self *Parser_DummyParser) GetToken(_env *LnsEnv) LnsAny {
    return Parser_eofToken
}

// 388: decl @lune.@base.@Parser.DummyParser.getStreamName
func (self *Parser_DummyParser) GetStreamName(_env *LnsEnv) string {
    return "dummy"
}

// 391: decl @lune.@base.@Parser.DummyParser.createPosition
func (self *Parser_DummyParser) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) *Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), nil)
}


// declaration Class -- CommentLayer
type Parser_CommentLayerMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList)
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList
    HasInvalidComment(_env *LnsEnv) LnsAny
}
type Parser_CommentLayer struct {
    commentList *LnsList
    tokenSet *LnsSet
    tokenList *LnsList
    FP Parser_CommentLayerMtd
}
func Parser_CommentLayer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_CommentLayer).FP
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
func (self *Parser_CommentLayer) Get_commentList(_env *LnsEnv) *LnsList{ return self.commentList }
// 404: DeclConstr
func (self *Parser_CommentLayer) InitParser_CommentLayer(_env *LnsEnv) {
    self.commentList = NewLnsList([]LnsAny{})
    
    self.tokenSet = NewLnsSet([]LnsAny{})
    
    self.tokenList = NewLnsList([]LnsAny{})
    
}

// 410: decl @lune.@base.@Parser.CommentLayer.addDirect
func (self *Parser_CommentLayer) AddDirect(_env *LnsEnv, commentList *LnsList) {
    for _, _comment := range( commentList.Items ) {
        comment := _comment.(Types_TokenDownCast).ToTypes_Token()
        self.commentList.Insert(Types_Token2Stem(comment))
    }
}

// 416: decl @lune.@base.@Parser.CommentLayer.add
func (self *Parser_CommentLayer) Add(_env *LnsEnv, token *Types_Token) {
    if Lns_op_not(self.tokenSet.Has(Types_Token2Stem(token))){
        self.tokenSet.Add(Types_Token2Stem(token))
        self.tokenList.Insert(Types_Token2Stem(token))
        self.FP.AddDirect(_env, token.FP.Get_commentList(_env))
    }
}

// 425: decl @lune.@base.@Parser.CommentLayer.clear
func (self *Parser_CommentLayer) Clear(_env *LnsEnv) {
    if self.commentList.Len() != 0{
        self.commentList = NewLnsList([]LnsAny{})
        
        self.tokenSet = NewLnsSet([]LnsAny{})
        
        self.tokenList = NewLnsList([]LnsAny{})
        
    }
}

// 441: decl @lune.@base.@Parser.CommentLayer.hasInvalidComment
func (self *Parser_CommentLayer) HasInvalidComment(_env *LnsEnv) LnsAny {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.tokenList.Len() > 1) &&
        _env.SetStackVal( self.tokenList.GetAt(2).(Types_TokenDownCast).ToTypes_Token().FP.Get_commentList(_env).GetAt(1).(Types_TokenDownCast).ToTypes_Token()) ||
        _env.SetStackVal( nil) )
}


// declaration Class -- CommentCtrl
type Parser_CommentCtrlMtd interface {
    Add(_env *LnsEnv, arg1 *Types_Token)
    AddDirect(_env *LnsEnv, arg1 *LnsList)
    Clear(_env *LnsEnv)
    Get_commentList(_env *LnsEnv) *LnsList
    HasInvalidComment(_env *LnsEnv) LnsAny
    Pop(_env *LnsEnv)
    Push(_env *LnsEnv)
}
type Parser_CommentCtrl struct {
    layerStack *LnsList
    layer *Parser_CommentLayer
    FP Parser_CommentCtrlMtd
}
func Parser_CommentCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_CommentCtrl).FP
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
// advertise -- 446
func (self *Parser_CommentCtrl) Add(_env *LnsEnv, arg1 *Types_Token) {
self.layer. FP.Add( _env, arg1)
}
// advertise -- 446
func (self *Parser_CommentCtrl) AddDirect(_env *LnsEnv, arg1 *LnsList) {
self.layer. FP.AddDirect( _env, arg1)
}
// advertise -- 446
func (self *Parser_CommentCtrl) Clear(_env *LnsEnv) {
self.layer. FP.Clear( _env)
}
// advertise -- 446
func (self *Parser_CommentCtrl) Get_commentList(_env *LnsEnv) *LnsList {
    return self.layer. FP.Get_commentList( _env)
}
// advertise -- 446
func (self *Parser_CommentCtrl) HasInvalidComment(_env *LnsEnv) LnsAny {
    return self.layer. FP.HasInvalidComment( _env)
}
// 450: DeclConstr
func (self *Parser_CommentCtrl) InitParser_CommentCtrl(_env *LnsEnv) {
    self.layer = NewParser_CommentLayer(_env)
    
    self.layerStack = NewLnsList([]LnsAny{Parser_CommentLayer2Stem(self.layer)})
    
}

// 455: decl @lune.@base.@Parser.CommentCtrl.push
func (self *Parser_CommentCtrl) Push(_env *LnsEnv) {
    self.layer = NewParser_CommentLayer(_env)
    
    self.layerStack.Insert(Parser_CommentLayer2Stem(self.layer))
}

// 460: decl @lune.@base.@Parser.CommentCtrl.pop
func (self *Parser_CommentCtrl) Pop(_env *LnsEnv) {
    self.layer = self.layerStack.GetAt(self.layerStack.Len()).(Parser_CommentLayerDownCast).ToParser_CommentLayer()
    
    self.layerStack.Remove(nil)
}


// declaration Class -- LowToken
type Parser_LowToken1244Mtd interface {
}
type Parser_LowToken1244 struct {
    Txt string
    Kind LnsInt
    FP Parser_LowToken1244Mtd
}
func Parser_LowToken12442Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_LowToken1244).FP
}
type Parser_LowToken1244DownCast interface {
    ToParser_LowToken1244() *Parser_LowToken1244
}
func Parser_LowToken1244DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_LowToken1244DownCast)
    if ok { return work.ToParser_LowToken1244() }
    return nil
}
func (obj *Parser_LowToken1244) ToParser_LowToken1244() *Parser_LowToken1244 {
    return obj
}
func NewParser_LowToken1244(_env *LnsEnv, arg1 string, arg2 LnsInt) *Parser_LowToken1244 {
    obj := &Parser_LowToken1244{}
    obj.FP = obj
    obj.InitParser_LowToken1244(_env, arg1, arg2)
    return obj
}
func (self *Parser_LowToken1244) InitParser_LowToken1244(_env *LnsEnv, arg1 string, arg2 LnsInt) {
    self.Txt = arg1
    self.Kind = arg2
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
    Parser_StreamParser____init_1084_(_env)
    Parser_eofToken = NewTypes_Token(_env, Types_TokenKind__Eof, "<EOF>", NewTypes_Position(_env, 0, 0, "eof"), false, NewLnsList([]LnsAny{}))
}
func init() {
    init_Parser = false
}
