// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Parser bool
var Parser__mod__ string
var Parser_noneToken *Types_Token
var Parser_eofToken *Types_Token
// for 152
func Parser_convExp525(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 219
func Parser_convExp871(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
type Parser_Position = Types_Position
type Parser_TokenKind = Types_TokenKind
type Parser_Token = Types_Token
// 31: decl @lune.@base.@Parser.isLuaKeyword
func Parser_isLuaKeyword(txt string) bool {
    return AsyncParser_isLuaKeyword(txt)
}

// 134: decl @lune.@base.@Parser.convFromRawToStr
func Parser_convFromRawToStr(txt string) string {
    if len(txt) == 0{
        return txt
    }
    if _switch466 := LnsInt(txt[1-1]); _switch466 == 39 || _switch466 == 34 {
    } else {
        return Lns_getVM().String_sub(txt,4, len(txt) - 3)
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
            _endIndex := Parser_convExp525(Lns_2DDD(Lns_getVM().String_find(workTxt, workPattern, workIndex, nil)))
            if _endIndex == nil{
                Util_err(Lns_getVM().String_format("error: illegal string -- %s", []LnsAny{workTxt}))
            } else {
                endIndex = _endIndex.(LnsInt)
            }
        }
        var workChar LnsInt
        workChar = LnsInt(workTxt[endIndex-1])
        if workChar == findChar{
            return retTxt + Lns_getVM().String_sub(workTxt,setIndex, endIndex - 1)
        } else if workChar == 92{
            var quote LnsInt
            quote = LnsInt(workTxt[endIndex + 1-1])
            if _switch632 := quote; _switch632 == 39 || _switch632 == 34 {
                retTxt = Lns_getVM().String_format("%s%s%c", []LnsAny{retTxt, Lns_getVM().String_sub(workTxt,setIndex, endIndex - 1), quote})
                
            } else {
                retTxt = Lns_getVM().String_format("%s%s", []LnsAny{retTxt, Lns_getVM().String_sub(workTxt,setIndex, endIndex + 1)})
                
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

// 374: decl @lune.@base.@Parser.isOp2
func Parser_isOp2(ope string) bool {
    return AsyncParser_isOp2(ope)
}

// 378: decl @lune.@base.@Parser.isOp1
func Parser_isOp1(ope string) bool {
    return AsyncParser_isOp1(ope)
}

// 387: decl @lune.@base.@Parser.getEofToken
func Parser_getEofToken() *Types_Token {
    return Parser_eofToken
}

// 474: decl @lune.@base.@Parser.quoteStr
func Parser_quoteStr(txt string) string {
    var work string
    work = txt
    var part string
    part = "\""
    {
        var _from2146 LnsInt = 1
        var _to2146 LnsInt = len(work)
        for _work2146 := _from2146; _work2146 <= _to2146; _work2146++ {
            index := _work2146
            var char LnsInt
            char = LnsInt(work[index-1])
            if _switch2144 := char; _switch2144 == 10 {
                part = part + "\\n"
                
            } else if _switch2144 == 13 {
                part = part + "\\r"
                
            } else if _switch2144 == 9 {
                part = part + "\\t"
                
            } else if _switch2144 == 34 {
                part = part + "\\\""
                
            } else if _switch2144 == 92 {
                part = part + "\\\\"
                
            } else {
                part = part + Lns_getVM().String_format("%c", []LnsAny{char})
                
            }
        }
    }
    work = part + "\""
    
    return work
}

// declaration Class -- TxtStream
type Parser_TxtStreamMtd interface {
    Close()
    GetSubstring(arg1 LnsInt, arg2 LnsAny) string
    Get_lineNo() LnsInt
    Get_txt() string
    Read(arg1 LnsAny) LnsAny
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
func NewParser_TxtStream(arg1 string) *Parser_TxtStream {
    obj := &Parser_TxtStream{}
    obj.FP = obj
    obj.InitParser_TxtStream(arg1)
    return obj
}
func (self *Parser_TxtStream) Get_txt() string{ return self.txt }
func (self *Parser_TxtStream) Get_lineNo() LnsInt{ return self.lineNo }
// 42: DeclConstr
func (self *Parser_TxtStream) InitParser_TxtStream(txt string) {
    self.txt = txt
    
    self.start = 1
    
    self.eof = false
    
    self.lineList = Str_getLineList(self.txt)
    
    self.lineNo = 1
    
}

// 61: decl @lune.@base.@Parser.TxtStream.getSubstring
func (self *Parser_TxtStream) GetSubstring(fromLineNo LnsInt,toLineNo LnsAny) string {
    var txt string
    txt = ""
    var to LnsInt
    to = Lns_unwrapDefault( toLineNo, self.lineList.Len() + 1).(LnsInt)
    {
        var _from167 LnsInt = fromLineNo
        var _to167 LnsInt = to - 1
        for _work167 := _from167; _work167 <= _to167; _work167++ {
            index := _work167
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( index < 1) ||
                Lns_GetEnv().SetStackVal( index > self.lineList.Len()) ).(bool){
                break
            }
            txt = Lns_getVM().String_format("%s%s", []LnsAny{txt, self.lineList.GetAt(index).(string)})
            
        }
    }
    return txt
}

// 73: decl @lune.@base.@Parser.TxtStream.read
func (self *Parser_TxtStream) Read(mode LnsAny) LnsAny {
    if mode != "*l"{
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{mode}))
    }
    if self.lineNo > self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    
    var line string
    line = self.lineList.GetAt(self.lineNo - 1).(string)
    if Str_endsWith(line, "\n"){
        return Lns_getVM().String_sub(line,1, len(line) - 1)
    }
    return line
}

// 87: decl @lune.@base.@Parser.TxtStream.close
func (self *Parser_TxtStream) Close() {
}


// declaration Class -- Parser
type Parser_ParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
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
func (self *Parser_Parser) InitParser_Parser() {
}


// 99: decl @lune.@base.@Parser.Parser.createPosition
func (self *Parser_Parser) CreatePosition(lineNo LnsInt,column LnsInt) *Types_Position {
    return NewTypes_Position(lineNo, column, self.FP.GetStreamName())
}


type Parser_PushbackParser interface {
        CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
        Error(arg1 string)
        GetStreamName() string
        GetTokenNoErr() *Types_Token
        NewPushback(arg1 *LnsList)
        Pushback()
        PushbackStr(arg1 string, arg2 string, arg3 *Types_Position)
        PushbackToken(arg1 *Types_Token)
}
func Lns_cast2Parser_PushbackParser( obj LnsAny ) LnsAny {
    if _, ok := obj.(Parser_PushbackParser); ok { 
        return obj
    }
    return nil
}

// declaration Class -- WrapParser
type Parser_WrapParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
}
type Parser_WrapParser struct {
    Parser_Parser
    parser *Parser_Parser
    name string
    FP Parser_WrapParserMtd
}
func Parser_WrapParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_WrapParser).FP
}
type Parser_WrapParserDownCast interface {
    ToParser_WrapParser() *Parser_WrapParser
}
func Parser_WrapParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_WrapParserDownCast)
    if ok { return work.ToParser_WrapParser() }
    return nil
}
func (obj *Parser_WrapParser) ToParser_WrapParser() *Parser_WrapParser {
    return obj
}
func NewParser_WrapParser(arg1 *Parser_Parser, arg2 string) *Parser_WrapParser {
    obj := &Parser_WrapParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_WrapParser(arg1, arg2)
    return obj
}
func (self *Parser_WrapParser) InitParser_WrapParser(arg1 *Parser_Parser, arg2 string) {
    self.Parser_Parser.InitParser_Parser( )
    self.parser = arg1
    self.name = arg2
}
// 118: decl @lune.@base.@Parser.WrapParser.getToken
func (self *Parser_WrapParser) GetToken() LnsAny {
    var token LnsAny
    token = self.parser.FP.GetToken()
    return token
}

// 122: decl @lune.@base.@Parser.WrapParser.getStreamName
func (self *Parser_WrapParser) GetStreamName() string {
    return self.name
}


// declaration Class -- StreamParser
var Parser_StreamParser__stdinStreamModuleName LnsAny
var Parser_StreamParser__stdinTxt string
// 179: decl @lune.@base.@Parser.StreamParser.___init
func Parser_StreamParser____init_1134_() {
    Parser_StreamParser__stdinStreamModuleName = nil
    
    Parser_StreamParser__stdinTxt = ""
    
}

type Parser_StreamParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
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
func NewParser_StreamParser(arg1 Lns_iStream, arg2 string, arg3 LnsAny, arg4 LnsAny) *Parser_StreamParser {
    obj := &Parser_StreamParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_StreamParser(arg1, arg2, arg3, arg4)
    return obj
}
// 188: decl @lune.@base.@Parser.StreamParser.setStdinStream
func Parser_StreamParser_setStdinStream(moduleName string) {
    Parser_StreamParser__stdinStreamModuleName = moduleName
    
    Parser_StreamParser__stdinTxt = Lns_unwrapDefault( Lns_io_stdin.Read("*a"), "").(string)
    
}

// 200: DeclConstr
func (self *Parser_StreamParser) InitParser_StreamParser(stream Lns_iStream,name string,luaMode LnsAny,pos LnsAny) {
    self.InitParser_Parser()
    self.streamName = name
    
    self.pos = 1
    
    self.lineTokenList = NewLnsList([]LnsAny{})
    
    self.overridePos = pos
    
    self.asyncParser = NewAsyncParser_Parser(stream, name, luaMode, pos)
    
}

// 212: decl @lune.@base.@Parser.StreamParser.getStreamName
func (self *Parser_StreamParser) GetStreamName() string {
    return self.streamName
}

// 216: decl @lune.@base.@Parser.StreamParser.create
func Parser_StreamParser_create(path string,luaMode bool,moduleName string) LnsAny {
    var stream Lns_iStream
    if Parser_StreamParser__stdinStreamModuleName != moduleName{
        {
            var _stream LnsAny
            _stream = Parser_convExp871(Lns_2DDD(Lns_io_open(path, "r")))
            if _stream == nil {
                return nil
            }
            stream = _stream.(Lns_iStream)
        }
    } else { 
        stream = NewParser_TxtStream(Parser_StreamParser__stdinTxt).FP
        
    }
    return NewParser_StreamParser(stream, path, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( luaMode) ||
        Lns_GetEnv().SetStackVal( Str_endsWith(path, ".lua")) &&
        Lns_GetEnv().SetStackVal( true) ).(bool), nil)
}

// 230: decl @lune.@base.@Parser.StreamParser.getToken
func (self *Parser_StreamParser) GetToken() LnsAny {
    if self.lineTokenList.Len() < self.pos{
        self.pos = 1
        
        self.lineTokenList = NewLnsList([]LnsAny{})
        
        for self.lineTokenList.Len() == 0 {
            var pipeItem *Async_PipeItem
            
            {
                _pipeItem := self.asyncParser.FP.GetNext()
                if _pipeItem == nil{
                    return nil
                } else {
                    pipeItem = _pipeItem.(*Async_PipeItem)
                }
            }
            self.lineTokenList = pipeItem.FP.Get_item().(*AsyncParser_AsyncItem).List
            
        }
    }
    var token *Types_Token
    token = self.lineTokenList.GetAt(self.pos).(Types_TokenDownCast).ToTypes_Token()
    self.pos = self.pos + 1
    
    return token
}


// declaration Class -- DefaultPushbackParser
type Parser_DefaultPushbackParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    Error(arg1 string)
    GetLastPos() *Types_Position
    GetNearCode() string
    GetStreamName() string
    GetTokenNoErr() *Types_Token
    Get_currentToken() *Types_Token
    NewPushback(arg1 *LnsList)
    Pushback()
    PushbackStr(arg1 string, arg2 string, arg3 *Types_Position)
    PushbackToken(arg1 *Types_Token)
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
func NewParser_DefaultPushbackParser(arg1 *Parser_Parser) *Parser_DefaultPushbackParser {
    obj := &Parser_DefaultPushbackParser{}
    obj.FP = obj
    obj.InitParser_DefaultPushbackParser(arg1)
    return obj
}
func (self *Parser_DefaultPushbackParser) Get_currentToken() *Types_Token{ return self.currentToken }
// 256: DeclConstr
func (self *Parser_DefaultPushbackParser) InitParser_DefaultPushbackParser(parser *Parser_Parser) {
    self.parser = parser
    
    self.pushbackedList = NewLnsList([]LnsAny{})
    
    self.usedTokenList = NewLnsList([]LnsAny{})
    
    self.currentToken = Types_noneToken
    
}

// 263: decl @lune.@base.@Parser.DefaultPushbackParser.createPosition
func (self *Parser_DefaultPushbackParser) CreatePosition(lineNo LnsInt,column LnsInt) *Types_Position {
    return self.parser.FP.CreatePosition(lineNo, column)
}

// 267: decl @lune.@base.@Parser.DefaultPushbackParser.getTokenNoErr
func (self *Parser_DefaultPushbackParser) GetTokenNoErr() *Types_Token {
    if self.pushbackedList.Len() > 0{
        self.currentToken = self.pushbackedList.GetAt(self.pushbackedList.Len()).(Types_TokenDownCast).ToTypes_Token()
        
        self.pushbackedList.Remove(nil)
    } else { 
        {
            _token := self.parser.FP.GetToken()
            if _token != nil {
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

// 285: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackToken
func (self *Parser_DefaultPushbackParser) PushbackToken(token *Types_Token) {
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

// 308: decl @lune.@base.@Parser.DefaultPushbackParser.pushback
func (self *Parser_DefaultPushbackParser) Pushback() {
    self.FP.PushbackToken(self.currentToken)
}

// 311: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackStr
func (self *Parser_DefaultPushbackParser) PushbackStr(name string,statement string,pos *Types_Position) {
    var memStream *Parser_TxtStream
    memStream = NewParser_TxtStream(statement)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(memStream.FP, name, false, pos)
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        {
            __exp := parser.FP.GetToken()
            if __exp != nil {
                _exp := __exp.(*Types_Token)
                list.Insert(Types_Token2Stem(_exp))
            } else {
                break
            }
        }
    }
    {
        var _from1415 LnsInt = list.Len()
        var _to1415 LnsInt = 1
        _work1415 := _from1415
        _delta1415 := -1
        for {
            if _delta1415 > 0 {
               if _work1415 > _to1415 { break }
            } else {
               if _work1415 < _to1415 { break }
            }
            index := _work1415
            self.FP.PushbackToken(list.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work1415 += _delta1415
        }
    }
}

// 328: decl @lune.@base.@Parser.DefaultPushbackParser.newPushback
func (self *Parser_DefaultPushbackParser) NewPushback(tokenList *LnsList) {
    {
        var _from1444 LnsInt = tokenList.Len()
        var _to1444 LnsInt = 1
        _work1444 := _from1444
        _delta1444 := -1
        for {
            if _delta1444 > 0 {
               if _work1444 > _to1444 { break }
            } else {
               if _work1444 < _to1444 { break }
            }
            index := _work1444
            self.FP.PushbackToken(tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work1444 += _delta1444
        }
    }
}

// 333: decl @lune.@base.@Parser.DefaultPushbackParser.error
func (self *Parser_DefaultPushbackParser) Error(message string) {
    Util_err(message)
}

// 338: decl @lune.@base.@Parser.DefaultPushbackParser.getLastPos
func (self *Parser_DefaultPushbackParser) GetLastPos() *Types_Position {
    var pos *Types_Position
    pos = self.parser.FP.CreatePosition(0, 0)
    if self.FP.Get_currentToken().Kind != Types_TokenKind__Eof{
        pos = self.FP.Get_currentToken().Pos
        
    } else { 
        if self.usedTokenList.Len() > 0{
            var token *Types_Token
            token = self.usedTokenList.GetAt(self.usedTokenList.Len()).(Types_TokenDownCast).ToTypes_Token()
            pos = token.Pos
            
        }
    }
    return pos
}

// 354: decl @lune.@base.@Parser.DefaultPushbackParser.getNearCode
func (self *Parser_DefaultPushbackParser) GetNearCode() string {
    var code string
    code = ""
    {
        var _from1619 LnsInt = self.usedTokenList.Len() - 30
        var _to1619 LnsInt = self.usedTokenList.Len()
        for _work1619 := _from1619; _work1619 <= _to1619; _work1619++ {
            index := _work1619
            if index > 1{
                var token *Types_Token
                token = self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token()
                if token.Consecutive{
                    code = Lns_getVM().String_format("%s%s", []LnsAny{code, self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token().Txt})
                    
                } else { 
                    code = Lns_getVM().String_format("%s %s", []LnsAny{code, self.usedTokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token().Txt})
                    
                }
            }
        }
    }
    return Lns_getVM().String_format("%s -- current '%s'", []LnsAny{code, self.currentToken.Txt})
}

// 369: decl @lune.@base.@Parser.DefaultPushbackParser.getStreamName
func (self *Parser_DefaultPushbackParser) GetStreamName() string {
    return self.parser.FP.GetStreamName()
}


// declaration Class -- DummyParser
type Parser_DummyParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
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
func NewParser_DummyParser() *Parser_DummyParser {
    obj := &Parser_DummyParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_DummyParser()
    return obj
}
func (self *Parser_DummyParser) InitParser_DummyParser() {
    self.Parser_Parser.InitParser_Parser( )
}
// 391: decl @lune.@base.@Parser.DummyParser.getToken
func (self *Parser_DummyParser) GetToken() LnsAny {
    return Parser_eofToken
}

// 394: decl @lune.@base.@Parser.DummyParser.getStreamName
func (self *Parser_DummyParser) GetStreamName() string {
    return "dummy"
}


// declaration Class -- CommentLayer
type Parser_CommentLayerMtd interface {
    Add(arg1 *Types_Token)
    AddDirect(arg1 *LnsList)
    Clear()
    Get_commentList() *LnsList
    HasInvalidComment() LnsAny
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
func NewParser_CommentLayer() *Parser_CommentLayer {
    obj := &Parser_CommentLayer{}
    obj.FP = obj
    obj.InitParser_CommentLayer()
    return obj
}
func (self *Parser_CommentLayer) Get_commentList() *LnsList{ return self.commentList }
// 407: DeclConstr
func (self *Parser_CommentLayer) InitParser_CommentLayer() {
    self.commentList = NewLnsList([]LnsAny{})
    
    self.tokenSet = NewLnsSet([]LnsAny{})
    
    self.tokenList = NewLnsList([]LnsAny{})
    
}

// 413: decl @lune.@base.@Parser.CommentLayer.addDirect
func (self *Parser_CommentLayer) AddDirect(commentList *LnsList) {
    for _, _comment := range( commentList.Items ) {
        comment := _comment.(Types_TokenDownCast).ToTypes_Token()
        self.commentList.Insert(Types_Token2Stem(comment))
    }
}

// 419: decl @lune.@base.@Parser.CommentLayer.add
func (self *Parser_CommentLayer) Add(token *Types_Token) {
    if Lns_op_not(self.tokenSet.Has(Types_Token2Stem(token))){
        self.tokenSet.Add(Types_Token2Stem(token))
        self.tokenList.Insert(Types_Token2Stem(token))
        self.FP.AddDirect(token.FP.Get_commentList())
    }
}

// 428: decl @lune.@base.@Parser.CommentLayer.clear
func (self *Parser_CommentLayer) Clear() {
    if self.commentList.Len() != 0{
        self.commentList = NewLnsList([]LnsAny{})
        
        self.tokenSet = NewLnsSet([]LnsAny{})
        
        self.tokenList = NewLnsList([]LnsAny{})
        
    }
}

// 444: decl @lune.@base.@Parser.CommentLayer.hasInvalidComment
func (self *Parser_CommentLayer) HasInvalidComment() LnsAny {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.tokenList.Len() > 1) &&
        Lns_GetEnv().SetStackVal( self.tokenList.GetAt(2).(Types_TokenDownCast).ToTypes_Token().FP.Get_commentList().GetAt(1).(Types_TokenDownCast).ToTypes_Token()) ||
        Lns_GetEnv().SetStackVal( nil) )
}


// declaration Class -- CommentCtrl
type Parser_CommentCtrlMtd interface {
    Add(arg1 *Types_Token)
    AddDirect(arg1 *LnsList)
    Clear()
    Get_commentList() *LnsList
    HasInvalidComment() LnsAny
    Pop()
    Push()
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
func NewParser_CommentCtrl() *Parser_CommentCtrl {
    obj := &Parser_CommentCtrl{}
    obj.FP = obj
    obj.InitParser_CommentCtrl()
    return obj
}
func (self *Parser_CommentCtrl) Add(arg1 *Types_Token) {
self.layer. FP.Add( arg1)
}
func (self *Parser_CommentCtrl) AddDirect(arg1 *LnsList) {
self.layer. FP.AddDirect( arg1)
}
func (self *Parser_CommentCtrl) Clear() {
self.layer. FP.Clear( )
}
func (self *Parser_CommentCtrl) Get_commentList() *LnsList {
    return self.layer. FP.Get_commentList( )
}
func (self *Parser_CommentCtrl) HasInvalidComment() LnsAny {
    return self.layer. FP.HasInvalidComment( )
}
// 453: DeclConstr
func (self *Parser_CommentCtrl) InitParser_CommentCtrl() {
    self.layer = NewParser_CommentLayer()
    
    self.layerStack = NewLnsList([]LnsAny{Parser_CommentLayer2Stem(self.layer)})
    
}

// 458: decl @lune.@base.@Parser.CommentCtrl.push
func (self *Parser_CommentCtrl) Push() {
    self.layer = NewParser_CommentLayer()
    
    self.layerStack.Insert(Parser_CommentLayer2Stem(self.layer))
}

// 463: decl @lune.@base.@Parser.CommentCtrl.pop
func (self *Parser_CommentCtrl) Pop() {
    self.layer = self.layerStack.GetAt(self.layerStack.Len()).(Parser_CommentLayerDownCast).ToParser_CommentLayer()
    
    self.layerStack.Remove(nil)
}


// declaration Class -- LowToken
type Parser_LowToken1360Mtd interface {
}
type Parser_LowToken1360 struct {
    Txt string
    Kind LnsInt
    FP Parser_LowToken1360Mtd
}
func Parser_LowToken13602Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_LowToken1360).FP
}
type Parser_LowToken1360DownCast interface {
    ToParser_LowToken1360() *Parser_LowToken1360
}
func Parser_LowToken1360DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_LowToken1360DownCast)
    if ok { return work.ToParser_LowToken1360() }
    return nil
}
func (obj *Parser_LowToken1360) ToParser_LowToken1360() *Parser_LowToken1360 {
    return obj
}
func NewParser_LowToken1360(arg1 string, arg2 LnsInt) *Parser_LowToken1360 {
    obj := &Parser_LowToken1360{}
    obj.FP = obj
    obj.InitParser_LowToken1360(arg1, arg2)
    return obj
}
func (self *Parser_LowToken1360) InitParser_LowToken1360(arg1 string, arg2 LnsInt) {
    self.Txt = arg1
    self.Kind = arg2
}

func Lns_Parser_init() {
    if init_Parser { return }
    init_Parser = true
    Parser__mod__ = "@lune.@base.@Parser"
    Lns_InitMod()
    Lns_Util_init()
    Lns_Str_init()
    Lns_Types_init()
    Lns_Async_init()
    Lns_AsyncParser_init()
    Parser_noneToken = Types_noneToken
    Parser_StreamParser____init_1134_()
    Parser_eofToken = NewTypes_Token(Types_TokenKind__Eof, "<EOF>", NewTypes_Position(0, 0, "eof"), false, NewLnsList([]LnsAny{}))
}
func init() {
    init_Parser = false
}
