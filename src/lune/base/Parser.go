// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Parser bool
var Parser__mod__ string
var Parser_luaKeywordSet *LnsSet
var Parser_noneToken *Types_Token
var Parser_quotedCharSet *LnsSet
var Parser_op2Set *LnsSet
var Parser_op1Set *LnsSet
var Parser_eofToken *Types_Token
// for 692
func Parser_convExp2827(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 816
func Parser_convExp3815(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 339
func Parser_convExp1291(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 446
func Parser_convExp1802(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 750
func Parser_convExp3074(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 757
func Parser_convExp3124(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 764
func Parser_convExp3174(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 773
func Parser_convExp3252(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 778
func Parser_convExp3281(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 790
func Parser_convExp3324(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 895
func Parser_convExp3927(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 918
func Parser_convExp4067(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
type Parser_Position = Types_Position
type Parser_Token = Types_Token
// 56: decl @lune.@base.@Parser.isLuaKeyword
func Parser_isLuaKeyword(txt string) bool {
    return Parser_luaKeywordSet.Has(txt)
}

// 61: decl @lune.@base.@Parser.createReserveInfo
func Parser_createReserveInfo_1032_(luaMode LnsAny)(*LnsSet, *LnsSet, *LnsSet, *LnsMap) {
    var keywordSet *LnsSet
    keywordSet = NewLnsSet([]LnsAny{})
    var typeSet *LnsSet
    typeSet = NewLnsSet([]LnsAny{})
    var builtInSet *LnsSet
    builtInSet = NewLnsSet([]LnsAny{})
    builtInSet.Add("require")
    for _key := range( Parser_luaKeywordSet.Items ) {
        key := _key.(string)
        if Lns_op_not(builtInSet.Has(key)){
            keywordSet.Add(key)
        }
    }
    if Lns_op_not(luaMode){
        keywordSet.Add("null")
        keywordSet.Add("let")
        keywordSet.Add("mut")
        keywordSet.Add("pub")
        keywordSet.Add("pro")
        keywordSet.Add("pri")
        keywordSet.Add("fn")
        keywordSet.Add("each")
        keywordSet.Add("form")
        keywordSet.Add("class")
        builtInSet.Add("super")
        keywordSet.Add("static")
        keywordSet.Add("advertise")
        keywordSet.Add("import")
        keywordSet.Add("new")
        keywordSet.Add("!")
        keywordSet.Add("unwrap")
        keywordSet.Add("sync")
        typeSet.Add("int")
        typeSet.Add("real")
        typeSet.Add("stem")
        typeSet.Add("str")
        typeSet.Add("Map")
        typeSet.Add("bool")
    }
    var multiCharDelimitMap *LnsMap
    multiCharDelimitMap = NewLnsMap( map[LnsAny]LnsAny{})
    multiCharDelimitMap.Set("=",NewLnsList([]LnsAny{"=="}))
    multiCharDelimitMap.Set("<",NewLnsList([]LnsAny{"<="}))
    multiCharDelimitMap.Set(">",NewLnsList([]LnsAny{">="}))
    if Lns_op_not(luaMode){
        multiCharDelimitMap.Set("|",NewLnsList([]LnsAny{"|<", "|>"}))
        multiCharDelimitMap.Set("|<",NewLnsList([]LnsAny{"|<<"}))
        multiCharDelimitMap.Set("|>",NewLnsList([]LnsAny{"|>>"}))
        multiCharDelimitMap.Set("[",NewLnsList([]LnsAny{"[@"}))
        multiCharDelimitMap.Set("(",NewLnsList([]LnsAny{"(@"}))
        multiCharDelimitMap.Set("~",NewLnsList([]LnsAny{"~=", "~~"}))
        multiCharDelimitMap.Set("$",NewLnsList([]LnsAny{"$[", "$.", "$("}))
        multiCharDelimitMap.Set("$.",NewLnsList([]LnsAny{"$.$"}))
        multiCharDelimitMap.Set(".",NewLnsList([]LnsAny{"..", ".$"}))
        multiCharDelimitMap.Set("..",NewLnsList([]LnsAny{"..."}))
        multiCharDelimitMap.Set(",",NewLnsList([]LnsAny{",,"}))
        multiCharDelimitMap.Set(",,",NewLnsList([]LnsAny{",,,"}))
        multiCharDelimitMap.Set(",,,",NewLnsList([]LnsAny{",,,,"}))
        multiCharDelimitMap.Set("@",NewLnsList([]LnsAny{"@@"}))
        multiCharDelimitMap.Set("@@",NewLnsList([]LnsAny{"@@@", "@@="}))
        multiCharDelimitMap.Set("#",NewLnsList([]LnsAny{"##"}))
        multiCharDelimitMap.Set("*",NewLnsList([]LnsAny{"**"}))
    } else { 
        multiCharDelimitMap.Set(".",NewLnsList([]LnsAny{".."}))
        multiCharDelimitMap.Set("~",NewLnsList([]LnsAny{"~="}))
    }
    return keywordSet, typeSet, builtInSet, multiCharDelimitMap
}

// 321: decl @lune.@base.@Parser.convFromRawToStr
func Parser_convFromRawToStr(txt string) string {
    if len(txt) == 0{
        return txt
    }
    if _switch1232 := LnsInt(txt[1-1]); _switch1232 == 39 || _switch1232 == 34 {
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
            _endIndex := Parser_convExp1291(Lns_2DDD(Lns_getVM().String_find(workTxt, workPattern, workIndex, nil)))
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
            if _switch1398 := quote; _switch1398 == 39 || _switch1398 == 34 {
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

// 653: decl @lune.@base.@Parser.isOp2
func Parser_isOp2(ope string) bool {
    return Parser_op2Set.Has(ope)
}

// 657: decl @lune.@base.@Parser.isOp1
func Parser_isOp1(ope string) bool {
    return Parser_op1Set.Has(ope)
}




// 749: decl @lune.@base.@Parser.StreamParser.parse.addVal.analyzeNumber
func StreamParser_parse_addVal__analyzeNumber_1500_(token string,beginIndex LnsInt)(LnsInt, bool) {
    var nonNumIndex LnsInt
    
    {
        _nonNumIndex := Parser_convExp3074(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", beginIndex, nil)))
        if _nonNumIndex == nil{
            return len(token), true
        } else {
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
    }
    var intFlag bool
    intFlag = true
    var nonNumChar LnsInt
    nonNumChar = LnsInt(token[nonNumIndex-1])
    if nonNumChar == 46{
        intFlag = false
        
        {
            var _nonNumIndex LnsAny
            _nonNumIndex = Parser_convExp3124(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
            if _nonNumIndex == nil {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
        
    }
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( nonNumChar == 88) ||
        Lns_setStackVal( nonNumChar == 120) ).(bool){
        {
            var _nonNumIndex LnsAny
            _nonNumIndex = Parser_convExp3174(Lns_2DDD(Lns_getVM().String_find(token,"[^%da-fA-F]", nonNumIndex + 1, nil)))
            if _nonNumIndex == nil {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
        
    }
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( nonNumChar == 69) ||
        Lns_setStackVal( nonNumChar == 101) ).(bool){
        intFlag = false
        
        var nextChar LnsInt
        nextChar = LnsInt(token[nonNumIndex + 1-1])
        if Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( nextChar == 45) ||
            Lns_setStackVal( nextChar == 43) ).(bool){
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = Parser_convExp3252(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 2, nil)))
                if _nonNumIndex == nil {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        } else { 
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = Parser_convExp3281(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
                if _nonNumIndex == nil {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        }
    }
    return nonNumIndex - 1, intFlag
}



// 1024: decl @lune.@base.@Parser.getEofToken
func Parser_getEofToken() *Types_Token {
    return Parser_eofToken
}

// 1111: decl @lune.@base.@Parser.quoteStr
func Parser_quoteStr(txt string) string {
    var work string
    work = txt
    var part string
    part = "\""
    {
        var _from5105 LnsInt = 1
        var _to5105 LnsInt = len(work)
        for _work5105 := _from5105; _work5105 <= _to5105; _work5105++ {
            index := _work5105
            var char LnsInt
            char = LnsInt(work[index-1])
            if _switch5103 := char; _switch5103 == 10 {
                part = part + "\\n"
                
            } else if _switch5103 == 13 {
                part = part + "\\r"
                
            } else if _switch5103 == 9 {
                part = part + "\\t"
                
            } else if _switch5103 == 34 {
                part = part + "\\\""
                
            } else if _switch5103 == 92 {
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
// 148: DeclConstr
func (self *Parser_TxtStream) InitParser_TxtStream(txt string) {
    self.txt = txt
    
    self.start = 1
    
    self.eof = false
    
    self.lineList = Str_getLineList(self.txt)
    
    self.lineNo = 1
    
}

// 167: decl @lune.@base.@Parser.TxtStream.getSubstring
func (self *Parser_TxtStream) GetSubstring(fromLineNo LnsInt,toLineNo LnsAny) string {
    var txt string
    txt = ""
    var to LnsInt
    to = Lns_unwrapDefault( toLineNo, self.lineList.Len() + 1).(LnsInt)
    {
        var _from915 LnsInt = fromLineNo
        var _to915 LnsInt = to - 1
        for _work915 := _from915; _work915 <= _to915; _work915++ {
            index := _work915
            if Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( index < 1) ||
                Lns_setStackVal( index > self.lineList.Len()) ).(bool){
                break
            }
            txt = Lns_getVM().String_format("%s%s", []LnsAny{txt, self.lineList.GetAt(index).(string)})
            
        }
    }
    return txt
}

// 182: decl @lune.@base.@Parser.TxtStream.read
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

// 208: decl @lune.@base.@Parser.TxtStream.close
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


// 286: decl @lune.@base.@Parser.Parser.createPosition
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
        PushbackStr(arg1 string, arg2 string)
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
// 305: decl @lune.@base.@Parser.WrapParser.getToken
func (self *Parser_WrapParser) GetToken() LnsAny {
    var token LnsAny
    token = self.parser.FP.GetToken()
    return token
}

// 309: decl @lune.@base.@Parser.WrapParser.getStreamName
func (self *Parser_WrapParser) GetStreamName() string {
    return self.name
}


// declaration Class -- AsyncItem
type Parser_AsyncItemMtd interface {
    ToMap() *LnsMap
}
type Parser_AsyncItem struct {
    List *LnsList
    FP Parser_AsyncItemMtd
}
func Parser_AsyncItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_AsyncItem).FP
}
type Parser_AsyncItemDownCast interface {
    ToParser_AsyncItem() *Parser_AsyncItem
}
func Parser_AsyncItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_AsyncItemDownCast)
    if ok { return work.ToParser_AsyncItem() }
    return nil
}
func (obj *Parser_AsyncItem) ToParser_AsyncItem() *Parser_AsyncItem {
    return obj
}
func NewParser_AsyncItem(arg1 *LnsList) *Parser_AsyncItem {
    obj := &Parser_AsyncItem{}
    obj.FP = obj
    obj.InitParser_AsyncItem(arg1)
    return obj
}
func (self *Parser_AsyncItem) InitParser_AsyncItem(arg1 *LnsList) {
    self.List = arg1
}
func (self *Parser_AsyncItem) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["list"] = Lns_ToCollection( self.List )
    return obj
}
func (self *Parser_AsyncItem) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Parser_AsyncItem__fromMap_1269_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Parser_AsyncItem_FromMap( arg1, paramList )
}
func Parser_AsyncItem__fromStem_1272_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Parser_AsyncItem_FromMap( arg1, paramList )
}
func Parser_AsyncItem_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Parser_AsyncItem_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Parser_AsyncItem_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Parser_AsyncItem{}
    newObj.FP = newObj
    return Parser_AsyncItem_FromMapMain( newObj, objMap, paramList )
}
func Parser_AsyncItem_FromMapMain( newObj *Parser_AsyncItem, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToListSub( objMap.Items["list"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Types_Token_FromMapSub, false,nil}}); !ok {
       return false,nil,"list:" + mess.(string)
    } else {
       newObj.List = conv.(*LnsList)
    }
    return true, newObj, nil
}
func Parser_AsyncItem__createPipe_1278_(arg1 LnsInt) LnsAny{
   return NewLnspipe( arg1 )
}

// declaration Class -- AsyncStreamParser
type Parser_AsyncStreamParserMtd interface {
    Access() LnsAny
    GetNext() LnsAny
    Start()
}
type Parser_AsyncStreamParser struct {
    Async_Pipe
    streamParser *Parser_StreamParser
    FP Parser_AsyncStreamParserMtd
}
func Parser_AsyncStreamParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_AsyncStreamParser).FP
}
type Parser_AsyncStreamParserDownCast interface {
    ToParser_AsyncStreamParser() *Parser_AsyncStreamParser
}
func Parser_AsyncStreamParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_AsyncStreamParserDownCast)
    if ok { return work.ToParser_AsyncStreamParser() }
    return nil
}
func (obj *Parser_AsyncStreamParser) ToParser_AsyncStreamParser() *Parser_AsyncStreamParser {
    return obj
}
func NewParser_AsyncStreamParser(arg1 LnsAny, arg2 *Parser_StreamParser) *Parser_AsyncStreamParser {
    obj := &Parser_AsyncStreamParser{}
    obj.FP = obj
    obj.Async_Pipe.FP = obj
    obj.InitParser_AsyncStreamParser(arg1, arg2)
    return obj
}
// 376: DeclConstr
func (self *Parser_AsyncStreamParser) InitParser_AsyncStreamParser(pipe LnsAny,streamParser *Parser_StreamParser) {
    self.InitAsync_Pipe(pipe)
    self.streamParser = streamParser
    
}

// 460: decl @lune.@base.@Parser.AsyncStreamParser.access
func (self *Parser_AsyncStreamParser) Access() LnsAny {
    var tokenList *LnsList
    
    {
        _tokenList := self.streamParser.FP.parse()
        if _tokenList == nil{
            return nil
        } else {
            tokenList = _tokenList.(*LnsList)
        }
    }
    return NewAsync_PipeItem(Parser_AsyncItem2Stem(NewParser_AsyncItem(tokenList)))
}


// declaration Class -- StreamParser
var Parser_StreamParser__stdinStreamModuleName LnsAny
var Parser_StreamParser__stdinTxt string
// 382: decl @lune.@base.@Parser.StreamParser.___init
func Parser_StreamParser____init_1301_() {
    Parser_StreamParser__stdinStreamModuleName = nil
    
    Parser_StreamParser__stdinTxt = ""
    
}

type Parser_StreamParserMtd interface {
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    GetStreamName() string
    GetToken() LnsAny
    parse() LnsAny
}
type Parser_StreamParser struct {
    Parser_Parser
    stdinStreamModuleName LnsAny
    stdinTxt string
    eof bool
    stream Lns_iStream
    streamName string
    lineNo LnsInt
    prevToken *Types_Token
    pos LnsInt
    lineTokenList *LnsList
    keywordSet *LnsSet
    typeSet *LnsSet
    builtInSet *LnsSet
    multiCharDelimitMap *LnsMap
    luaMode bool
    pipe *Parser_AsyncStreamParser
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
func NewParser_StreamParser(arg1 Lns_iStream, arg2 string, arg3 LnsAny) *Parser_StreamParser {
    obj := &Parser_StreamParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitParser_StreamParser(arg1, arg2, arg3)
    return obj
}
// 391: decl @lune.@base.@Parser.StreamParser.setStdinStream
func Parser_StreamParser_setStdinStream(moduleName string) {
    Parser_StreamParser__stdinStreamModuleName = moduleName
    
    Parser_StreamParser__stdinTxt = Lns_unwrapDefault( Lns_io_stdin.Read("*a"), "").(string)
    
}

// 410: DeclConstr
func (self *Parser_StreamParser) InitParser_StreamParser(stream Lns_iStream,name string,luaMode LnsAny) {
    self.InitParser_Parser()
    self.pipe = NewParser_AsyncStreamParser(Parser_AsyncItem__createPipe_1278_(100), self)
    
    self.eof = false
    
    self.stream = stream
    
    self.streamName = name
    
    self.lineNo = 0
    
    self.pos = 1
    
    self.lineTokenList = NewLnsList([]LnsAny{})
    
    self.prevToken = Parser_noneToken
    
    self.luaMode = Lns_unwrapDefault( luaMode, false).(bool)
    
    var keywordSet *LnsSet
    var typeSet *LnsSet
    var builtInSet *LnsSet
    var multiCharDelimitMap *LnsMap
    keywordSet,typeSet,builtInSet,multiCharDelimitMap = Parser_createReserveInfo_1032_(luaMode)
    self.keywordSet = keywordSet
    
    self.typeSet = typeSet
    
    self.builtInSet = builtInSet
    
    self.multiCharDelimitMap = multiCharDelimitMap
    
}


// 439: decl @lune.@base.@Parser.StreamParser.getStreamName
func (self *Parser_StreamParser) GetStreamName() string {
    return self.streamName
}

// 443: decl @lune.@base.@Parser.StreamParser.create
func Parser_StreamParser_create(path string,luaMode bool,moduleName string) LnsAny {
    var stream Lns_iStream
    if Parser_StreamParser__stdinStreamModuleName != moduleName{
        {
            var _stream LnsAny
            _stream = Parser_convExp1802(Lns_2DDD(Lns_io_open(path, "r")))
            if _stream == nil {
                return nil
            }
            stream = _stream.(Lns_iStream)
        }
    } else { 
        stream = NewParser_TxtStream(Parser_StreamParser__stdinTxt).FP
        
    }
    return NewParser_StreamParser(stream, path, Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( luaMode) ||
        Lns_setStackVal( Str_endsWith(path, ".lua")) &&
        Lns_setStackVal( true) ).(bool))
}


// 665: decl @lune.@base.@Parser.StreamParser.parse
func (self *Parser_StreamParser) parse() LnsAny {
    var readLine func() LnsAny
    readLine = func() LnsAny {
        if self.eof{
            return nil
        }
        self.lineNo = self.lineNo + 1
        
        var line LnsAny
        line = self.stream.Read("*l")
        if Lns_op_not(line){
            self.eof = true
            
        }
        return line
    }
    var rawLine string
    
    {
        _rawLine := readLine()
        if _rawLine == nil{
            return nil
        } else {
            rawLine = _rawLine.(string)
        }
    }
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    var startIndex LnsInt
    startIndex = 1
    var multiComment func(comIndex LnsInt,termStr string)(string, LnsInt)
    multiComment = func(comIndex LnsInt,termStr string)(string, LnsInt) {
        var searchIndex LnsInt
        searchIndex = comIndex
        var comment string
        comment = ""
        for  {
            {
                _, _termEndIndex := Parser_convExp2827(Lns_2DDD(Lns_getVM().String_find(rawLine, termStr, searchIndex, true)))
                if _termEndIndex != nil {
                    termEndIndex := _termEndIndex.(LnsInt)
                    comment = comment + Lns_getVM().String_sub(rawLine,searchIndex, termEndIndex)
                    
                    return comment, termEndIndex + 1
                }
            }
            comment = comment + Lns_getVM().String_sub(rawLine,searchIndex, nil) + "\n"
            
            searchIndex = 1
            
            rawLine = Lns_unwrap( readLine()).(string)
            
        }
    // insert a dummy
        return "",0
    }
    var addVal func(kind LnsInt,val string,column LnsInt)
    addVal = func(kind LnsInt,val string,column LnsInt) {
        var createInfo func(tokenKind LnsInt,token string,tokenColumn LnsInt) *Types_Token
        createInfo = func(tokenKind LnsInt,token string,tokenColumn LnsInt) *Types_Token {
            if tokenKind == Types_TokenKind__Symb{
                if self.keywordSet.Has(token){
                    tokenKind = Types_TokenKind__Kywd
                    
                } else if self.typeSet.Has(token){
                    tokenKind = Types_TokenKind__Type
                    
                } else if Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( Parser_op2Set.Has(token)) ||
                    Lns_setStackVal( Parser_op1Set.Has(token)) ).(bool){
                    tokenKind = Types_TokenKind__Ope
                    
                }
            }
            var consecutive bool
            consecutive = false
            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( self.prevToken.Pos.LineNo == self.lineNo) &&
                Lns_setStackVal( self.prevToken.Pos.Column + len(self.prevToken.Txt) == tokenColumn) ).(bool)){
                consecutive = true
                
            }
            var newToken *Types_Token
            newToken = NewTypes_Token(tokenKind, token, self.FP.CreatePosition(self.lineNo, tokenColumn), consecutive, NewLnsList([]LnsAny{}))
            self.prevToken = newToken
            
            return newToken
        }
        if kind == Types_TokenKind__Symb{
            var searchIndex LnsInt
            searchIndex = 1
            for  {
                var tokenIndex LnsInt
                var tokenEndIndex LnsInt
                
                {
                    _tokenIndex, _tokenEndIndex := Parser_convExp3324(Lns_2DDD(Lns_getVM().String_find(val, "[%p%w]+", searchIndex, nil)))
                    if _tokenIndex == nil || _tokenEndIndex == nil{
                        break
                    } else {
                        tokenIndex = _tokenIndex.(LnsInt)
                        tokenEndIndex = _tokenEndIndex.(LnsInt)
                    }
                }
                var columnIndex LnsInt
                columnIndex = column + tokenIndex - 2
                searchIndex = tokenEndIndex + 1
                
                var token string
                token = Lns_getVM().String_sub(val,tokenIndex, tokenEndIndex)
                var subIndex LnsInt
                subIndex = 1
                for len(token) >= subIndex {
                    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                        Lns_setStackVal( Lns_car(Lns_getVM().String_find(token,"^[%d]", subIndex, nil))) ||
                        Lns_setStackVal( LnsInt(token[subIndex-1]) == 45) &&
                        Lns_setStackVal( Lns_car(Lns_getVM().String_find(token,"^[%d]", subIndex + 1, nil))) )){
                        var checkIndex LnsInt
                        checkIndex = subIndex
                        if LnsInt(token[checkIndex-1]) == 45{
                            checkIndex = checkIndex + 1
                            
                        }
                        var endIndex LnsInt
                        var intFlag bool
                        endIndex,intFlag = StreamParser_parse_addVal__analyzeNumber_1500_(token, checkIndex)
                        var info *Types_Token
                        info = createInfo(Lns_popVal( Lns_incStack() ||
                            Lns_setStackVal( intFlag) &&
                            Lns_setStackVal( Types_TokenKind__Int) ||
                            Lns_setStackVal( Types_TokenKind__Real) ).(LnsInt), Lns_getVM().String_sub(token,subIndex, endIndex), columnIndex + subIndex)
                        list.Insert(Types_Token2Stem(info))
                        subIndex = endIndex + 1
                        
                    } else { 
                        {
                            __exp := Parser_convExp3815(Lns_2DDD(Lns_getVM().String_find(token, "[^%w_]", subIndex, nil)))
                            if __exp != nil {
                                _exp := __exp.(LnsInt)
                                var index LnsInt
                                index = _exp
                                if index > subIndex{
                                    var info *Types_Token
                                    info = createInfo(Types_TokenKind__Symb, Lns_getVM().String_sub(token,subIndex, index - 1), columnIndex + subIndex)
                                    list.Insert(Types_Token2Stem(info))
                                }
                                var delimit string
                                delimit = Lns_getVM().String_sub(token,index, index)
                                var candidateList LnsAny
                                candidateList = self.multiCharDelimitMap.Items[delimit]
                                for Lns_isCondTrue( candidateList) {
                                    var findFlag bool
                                    findFlag = false
                                    for _, _candidate := range( Lns_unwrap( (candidateList)).(*LnsList).Items ) {
                                        candidate := _candidate.(string)
                                        if candidate == Lns_getVM().String_sub(token,index, index + len(candidate) - 1){
                                            delimit = candidate
                                            
                                            candidateList = self.multiCharDelimitMap.Items[delimit]
                                            
                                            findFlag = true
                                            
                                            break
                                        }
                                    }
                                    if Lns_op_not(findFlag){
                                        break
                                    }
                                }
                                subIndex = index + len(delimit)
                                
                                var workKind LnsInt
                                workKind = Types_TokenKind__Dlmt
                                if Lns_popVal( Lns_incStack() ||
                                    Lns_setStackVal( Parser_op2Set.Has(delimit)) ||
                                    Lns_setStackVal( Parser_op1Set.Has(delimit)) ).(bool){
                                    workKind = Types_TokenKind__Ope
                                    
                                }
                                if delimit == "..."{
                                    workKind = Types_TokenKind__Symb
                                    
                                }
                                if delimit == "?"{
                                    var nextChar string
                                    nextChar = Lns_getVM().String_sub(token,index, subIndex)
                                    list.Insert(Types_Token2Stem(createInfo(Types_TokenKind__Char, nextChar, columnIndex + subIndex)))
                                    subIndex = subIndex + 1
                                    
                                } else { 
                                    list.Insert(Types_Token2Stem(createInfo(workKind, delimit, columnIndex + index)))
                                }
                            } else {
                                if subIndex <= len(token){
                                    list.Insert(Types_Token2Stem(createInfo(Types_TokenKind__Symb, Lns_getVM().String_sub(token,subIndex, nil), columnIndex + subIndex)))
                                }
                                break
                            }
                        }
                    }
                }
            }
        } else { 
            list.Insert(Types_Token2Stem(createInfo(kind, val, column)))
        }
    }
    var searchIndex LnsInt
    searchIndex = startIndex
    var getChar func(index LnsInt) LnsInt
    getChar = func(index LnsInt) LnsInt {
        if len(rawLine) >= index{
            return LnsInt(rawLine[index-1])
        }
        return 0
    }
    for  {
        var syncIndexFlag bool
        syncIndexFlag = true
        var pattern string
        pattern = "[/%-%?\"%'%`]."
        var index LnsInt
        
        {
            _index := Parser_convExp3927(Lns_2DDD(Lns_getVM().String_find(rawLine, pattern, searchIndex, nil)))
            if _index == nil{
                addVal(Types_TokenKind__Symb, Lns_getVM().String_sub(rawLine,startIndex, nil), startIndex)
                return list
            } else {
                index = _index.(LnsInt)
            }
        }
        var findChar LnsInt
        findChar = getChar(index)
        var nextChar LnsInt
        nextChar = getChar(index + 1)
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( findChar == 45) &&
            Lns_setStackVal( nextChar != 45) ).(bool)){
            searchIndex = index + 1
            
            syncIndexFlag = false
            
        } else { 
            if startIndex < index{
                addVal(Types_TokenKind__Symb, Lns_getVM().String_sub(rawLine,startIndex, index - 1), startIndex)
            }
            if Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( findChar == 39) ||
                Lns_setStackVal( findChar == 34) ).(bool){
                var workIndex LnsInt
                workIndex = index + 1
                var workPattern string
                workPattern = "[\"'\\]"
                for  {
                    var endIndex LnsInt
                    
                    {
                        _endIndex := Parser_convExp4067(Lns_2DDD(Lns_getVM().String_find(rawLine, workPattern, workIndex, nil)))
                        if _endIndex == nil{
                            Util_err(Lns_getVM().String_format("%s:%d:%d: error: illegal string -- %s", []LnsAny{self.FP.GetStreamName(), self.lineNo, index, rawLine}))
                        } else {
                            endIndex = _endIndex.(LnsInt)
                        }
                    }
                    var workChar LnsInt
                    workChar = LnsInt(rawLine[endIndex-1])
                    if workChar == findChar{
                        addVal(Types_TokenKind__Str, Lns_getVM().String_sub(rawLine,index, endIndex), index)
                        searchIndex = endIndex + 1
                        
                        break
                    } else if workChar == 92{
                        workIndex = endIndex + 2
                        
                    } else { 
                        workIndex = endIndex + 1
                        
                    }
                }
            } else if findChar == 96{
                if Lns_isCondTrue( (Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( nextChar == findChar) &&
                    Lns_setStackVal( LnsInt(rawLine[index + 2-1]) == 96) ).(bool))){
                    var txt string
                    var nextIndex LnsInt
                    txt,nextIndex = multiComment(index + 3, "```")
                    addVal(Types_TokenKind__Str, "```" + txt, index)
                    searchIndex = nextIndex
                    
                } else { 
                    addVal(Types_TokenKind__Ope, "`", index)
                    searchIndex = index + 1
                    
                }
            } else if findChar == 63{
                var codeChar string
                codeChar = Lns_getVM().String_sub(rawLine,index + 1, index + 1)
                if nextChar == 92{
                    var quoted string
                    quoted = Lns_getVM().String_sub(rawLine,index + 2, index + 2)
                    if Parser_quotedCharSet.Has(quoted){
                        codeChar = Lns_getVM().String_sub(rawLine,index + 1, index + 2)
                        
                    } else { 
                        codeChar = quoted
                        
                    }
                    searchIndex = index + 3
                    
                } else { 
                    searchIndex = index + 2
                    
                }
                addVal(Types_TokenKind__Char, codeChar, index)
            } else { 
                if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( self.luaMode) &&
                    Lns_setStackVal( findChar == 45) &&
                    Lns_setStackVal( nextChar == 45) ).(bool)){
                    addVal(Types_TokenKind__Cmnt, Lns_getVM().String_sub(rawLine,index, nil), index)
                    searchIndex = len(rawLine) + 1
                    
                } else if findChar == 47{
                    if nextChar == 42{
                        var comment string
                        var nextIndex LnsInt
                        comment,nextIndex = multiComment(index + 2, "*/")
                        addVal(Types_TokenKind__Cmnt, "/*" + comment, index)
                        searchIndex = nextIndex
                        
                    } else if nextChar == 47{
                        addVal(Types_TokenKind__Cmnt, Lns_getVM().String_sub(rawLine,index, nil), index)
                        searchIndex = len(rawLine) + 1
                        
                    } else { 
                        addVal(Types_TokenKind__Ope, "/", index)
                        searchIndex = index + 1
                        
                    }
                } else { 
                    Util_err(Lns_getVM().String_format("%s:%d:%d: error: illegal syntax -- %s", []LnsAny{self.FP.GetStreamName(), self.lineNo, index, rawLine}))
                }
            }
        }
        if syncIndexFlag{
            startIndex = searchIndex
            
        }
    }
// insert a dummy
    return nil
}

// 1005: decl @lune.@base.@Parser.StreamParser.getToken
func (self *Parser_StreamParser) GetToken() LnsAny {
    if self.lineTokenList.Len() < self.pos{
        self.pos = 1
        
        self.lineTokenList = NewLnsList([]LnsAny{})
        
        for self.lineTokenList.Len() == 0 {
            var pipeItem *Async_PipeItem
            
            {
                _pipeItem := self.pipe.FP.GetNext()
                if _pipeItem == nil{
                    return nil
                } else {
                    pipeItem = _pipeItem.(*Async_PipeItem)
                }
            }
            self.lineTokenList = pipeItem.FP.Get_item().(*Parser_AsyncItem).List
            
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
    PushbackStr(arg1 string, arg2 string)
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
// 474: DeclConstr
func (self *Parser_DefaultPushbackParser) InitParser_DefaultPushbackParser(parser *Parser_Parser) {
    self.parser = parser
    
    self.pushbackedList = NewLnsList([]LnsAny{})
    
    self.usedTokenList = NewLnsList([]LnsAny{})
    
    self.currentToken = Parser_noneToken
    
}

// 481: decl @lune.@base.@Parser.DefaultPushbackParser.createPosition
func (self *Parser_DefaultPushbackParser) CreatePosition(lineNo LnsInt,column LnsInt) *Types_Position {
    return self.parser.FP.CreatePosition(lineNo, column)
}

// 485: decl @lune.@base.@Parser.DefaultPushbackParser.getTokenNoErr
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
                self.currentToken = Parser_noneToken
                
            }
        }
    }
    if self.currentToken.Kind != Types_TokenKind__Eof{
        self.usedTokenList.Insert(Types_Token2Stem(self.currentToken))
    }
    return self.currentToken
}

// 503: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackToken
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
                self.currentToken = Parser_noneToken
                
            }
        } else { 
            self.currentToken = Parser_noneToken
            
        }
    }
}

// 526: decl @lune.@base.@Parser.DefaultPushbackParser.pushback
func (self *Parser_DefaultPushbackParser) Pushback() {
    self.FP.PushbackToken(self.currentToken)
}

// 529: decl @lune.@base.@Parser.DefaultPushbackParser.pushbackStr
func (self *Parser_DefaultPushbackParser) PushbackStr(name string,statement string) {
    var memStream *Parser_TxtStream
    memStream = NewParser_TxtStream(statement)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(memStream.FP, name, false)
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
        var _from2283 LnsInt = list.Len()
        var _to2283 LnsInt = 1
        _work2283 := _from2283
        _delta2283 := -1
        for {
            if _delta2283 > 0 {
               if _work2283 > _to2283 { break }
            } else {
               if _work2283 < _to2283 { break }
            }
            index := _work2283
            self.FP.PushbackToken(list.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work2283 += _delta2283
        }
    }
}

// 546: decl @lune.@base.@Parser.DefaultPushbackParser.newPushback
func (self *Parser_DefaultPushbackParser) NewPushback(tokenList *LnsList) {
    {
        var _from2312 LnsInt = tokenList.Len()
        var _to2312 LnsInt = 1
        _work2312 := _from2312
        _delta2312 := -1
        for {
            if _delta2312 > 0 {
               if _work2312 > _to2312 { break }
            } else {
               if _work2312 < _to2312 { break }
            }
            index := _work2312
            self.FP.PushbackToken(tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _work2312 += _delta2312
        }
    }
}

// 551: decl @lune.@base.@Parser.DefaultPushbackParser.error
func (self *Parser_DefaultPushbackParser) Error(message string) {
    Util_err(message)
}

// 556: decl @lune.@base.@Parser.DefaultPushbackParser.getLastPos
func (self *Parser_DefaultPushbackParser) GetLastPos() *Types_Position {
    var pos *Types_Position
    pos = self.parser.FP.CreatePosition(0, 0)
    if self.currentToken.Kind != Types_TokenKind__Eof{
        pos = self.currentToken.Pos
        
    } else { 
        if self.usedTokenList.Len() > 0{
            var token *Types_Token
            token = self.usedTokenList.GetAt(self.usedTokenList.Len()).(Types_TokenDownCast).ToTypes_Token()
            pos = token.Pos
            
        }
    }
    return pos
}

// 572: decl @lune.@base.@Parser.DefaultPushbackParser.getNearCode
func (self *Parser_DefaultPushbackParser) GetNearCode() string {
    var code string
    code = ""
    {
        var _from2487 LnsInt = self.usedTokenList.Len() - 30
        var _to2487 LnsInt = self.usedTokenList.Len()
        for _work2487 := _from2487; _work2487 <= _to2487; _work2487++ {
            index := _work2487
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

// 587: decl @lune.@base.@Parser.DefaultPushbackParser.getStreamName
func (self *Parser_DefaultPushbackParser) GetStreamName() string {
    return self.parser.FP.GetStreamName()
}


// declaration Class -- TokenList
type Parser_TokenListMtd interface {
}
type Parser_TokenList struct {
    List *LnsList
    FP Parser_TokenListMtd
}
func Parser_TokenList2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_TokenList).FP
}
type Parser_TokenListDownCast interface {
    ToParser_TokenList() *Parser_TokenList
}
func Parser_TokenListDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_TokenListDownCast)
    if ok { return work.ToParser_TokenList() }
    return nil
}
func (obj *Parser_TokenList) ToParser_TokenList() *Parser_TokenList {
    return obj
}
func NewParser_TokenList(arg1 *LnsList) *Parser_TokenList {
    obj := &Parser_TokenList{}
    obj.FP = obj
    obj.InitParser_TokenList(arg1)
    return obj
}
func (self *Parser_TokenList) InitParser_TokenList(arg1 *LnsList) {
    self.List = arg1
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
// 1028: decl @lune.@base.@Parser.DummyParser.getToken
func (self *Parser_DummyParser) GetToken() LnsAny {
    return Parser_eofToken
}

// 1031: decl @lune.@base.@Parser.DummyParser.getStreamName
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
// 1044: DeclConstr
func (self *Parser_CommentLayer) InitParser_CommentLayer() {
    self.commentList = NewLnsList([]LnsAny{})
    
    self.tokenSet = NewLnsSet([]LnsAny{})
    
    self.tokenList = NewLnsList([]LnsAny{})
    
}

// 1050: decl @lune.@base.@Parser.CommentLayer.addDirect
func (self *Parser_CommentLayer) AddDirect(commentList *LnsList) {
    for _, _comment := range( commentList.Items ) {
        comment := _comment.(Types_TokenDownCast).ToTypes_Token()
        self.commentList.Insert(Types_Token2Stem(comment))
    }
}

// 1056: decl @lune.@base.@Parser.CommentLayer.add
func (self *Parser_CommentLayer) Add(token *Types_Token) {
    if Lns_op_not(self.tokenSet.Has(Types_Token2Stem(token))){
        self.tokenSet.Add(Types_Token2Stem(token))
        self.tokenList.Insert(Types_Token2Stem(token))
        self.FP.AddDirect(token.FP.Get_commentList())
    }
}

// 1065: decl @lune.@base.@Parser.CommentLayer.clear
func (self *Parser_CommentLayer) Clear() {
    if self.commentList.Len() != 0{
        self.commentList = NewLnsList([]LnsAny{})
        
        self.tokenSet = NewLnsSet([]LnsAny{})
        
        self.tokenList = NewLnsList([]LnsAny{})
        
    }
}

// 1081: decl @lune.@base.@Parser.CommentLayer.hasInvalidComment
func (self *Parser_CommentLayer) HasInvalidComment() LnsAny {
    return Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( self.tokenList.Len() > 1) &&
        Lns_setStackVal( self.tokenList.GetAt(2).(Types_TokenDownCast).ToTypes_Token().FP.Get_commentList().GetAt(1).(Types_TokenDownCast).ToTypes_Token()) ||
        Lns_setStackVal( nil) )
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
// 1090: DeclConstr
func (self *Parser_CommentCtrl) InitParser_CommentCtrl() {
    self.layer = NewParser_CommentLayer()
    
    self.layerStack = NewLnsList([]LnsAny{Parser_CommentLayer2Stem(self.layer)})
    
}

// 1095: decl @lune.@base.@Parser.CommentCtrl.push
func (self *Parser_CommentCtrl) Push() {
    self.layer = NewParser_CommentLayer()
    
    self.layerStack.Insert(Parser_CommentLayer2Stem(self.layer))
}

// 1100: decl @lune.@base.@Parser.CommentCtrl.pop
func (self *Parser_CommentCtrl) Pop() {
    self.layer = self.layerStack.GetAt(self.layerStack.Len()).(Parser_CommentLayerDownCast).ToParser_CommentLayer()
    
    self.layerStack.Remove(nil)
}


// declaration Class -- LowToken
type Parser_LowToken1638Mtd interface {
}
type Parser_LowToken1638 struct {
    Txt string
    Kind LnsInt
    FP Parser_LowToken1638Mtd
}
func Parser_LowToken16382Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Parser_LowToken1638).FP
}
type Parser_LowToken1638DownCast interface {
    ToParser_LowToken1638() *Parser_LowToken1638
}
func Parser_LowToken1638DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Parser_LowToken1638DownCast)
    if ok { return work.ToParser_LowToken1638() }
    return nil
}
func (obj *Parser_LowToken1638) ToParser_LowToken1638() *Parser_LowToken1638 {
    return obj
}
func NewParser_LowToken1638(arg1 string, arg2 LnsInt) *Parser_LowToken1638 {
    obj := &Parser_LowToken1638{}
    obj.FP = obj
    obj.InitParser_LowToken1638(arg1, arg2)
    return obj
}
func (self *Parser_LowToken1638) InitParser_LowToken1638(arg1 string, arg2 LnsInt) {
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
    Lns_Async_init()
    Lns_Types_init()
    Parser_luaKeywordSet = NewLnsSet([]LnsAny{"if", "else", "elseif", "while", "for", "in", "return", "break", "nil", "true", "false", "{", "}", "do", "require", "function", "then", "end", "repeat", "until", "goto", "local"})
    Parser_noneToken = NewTypes_Token(Types_TokenKind__Eof, "", NewTypes_Position(0, -1, "eof"), false, NewLnsList([]LnsAny{}))
    Parser_StreamParser____init_1301_()
    Parser_quotedCharSet = NewLnsSet([]LnsAny{"a", "b", "f", "n", "r", "t", "v", "\\", "\"", "'"})
    Parser_op2Set = NewLnsSet([]LnsAny{"+", "-", "*", "/", "^", "%", "&", "~", "|", "|>>", "|<<", "..", "<", "<=", ">", ">=", "==", "~=", "and", "or", "@", "@@", "@@@", "="})
    Parser_op1Set = NewLnsSet([]LnsAny{"-", "not", "#", "~", "*", "`", ",,", ",,,", ",,,,"})
    Parser_eofToken = NewTypes_Token(Types_TokenKind__Eof, "<EOF>", NewTypes_Position(0, 0, "eof"), false, NewLnsList([]LnsAny{}))
}
func init() {
    init_Parser = false
}
