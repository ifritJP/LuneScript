// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_AsyncParser bool
var AsyncParser__mod__ string
var AsyncParser_luaKeywordSet *LnsSet
var AsyncParser_quotedCharSet *LnsSet
var AsyncParser_op2Set *LnsSet
var AsyncParser_op1Set *LnsSet
var AsyncParser_defaultPipeSize LnsInt
// for 276
func AsyncParser_convExp0_1486(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 496
func AsyncParser_convExp0_2727(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 582
func AsyncParser_convExp0_2873(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 324
func AsyncParser_convExp0_515(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 247
func AsyncParser_convExp0_410(arg1 []LnsAny) (*LnsSet, *LnsSet, *LnsSet, *LnsMap) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsSet), Lns_getFromMulti( arg1, 1 ).(*LnsSet), Lns_getFromMulti( arg1, 2 ).(*LnsSet), Lns_getFromMulti( arg1, 3 ).(*LnsMap)
}
// for 289
func AsyncParser_convExp0_1547(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 293
func AsyncParser_convExp0_1573(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 299
func AsyncParser_convExp0_1601(arg1 []LnsAny) (string, bool, LnsAny, string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(string), Lns_getFromMulti( arg1, 4 )
}
// for 350
func AsyncParser_convExp0_1714(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 400
func AsyncParser_convExp0_1911(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 408
func AsyncParser_convExp0_1962(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 416
func AsyncParser_convExp0_2013(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 427
func AsyncParser_convExp0_2093(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 432
func AsyncParser_convExp0_2123(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 470
func AsyncParser_convExp0_2225(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 487
func AsyncParser_convExp0_2343(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 613
func AsyncParser_convExp0_3018(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 636
func AsyncParser_convExp0_3163(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 662
func AsyncParser_convExp0_3281(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 698
func AsyncParser_convExp0_3533(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// 59: decl @lune.@base.@AsyncParser.isLuaKeyword
func AsyncParser_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncParser_luaKeywordSet.Has(txt)
}

// 64: decl @lune.@base.@AsyncParser.createReserveInfo
func AsyncParser_createReserveInfo_1_(_env *LnsEnv, luaMode LnsAny)(*LnsSet, *LnsSet, *LnsSet, *LnsMap) {
    var keywordSet *LnsSet
    keywordSet = NewLnsSet([]LnsAny{})
    var typeSet *LnsSet
    typeSet = NewLnsSet([]LnsAny{})
    var builtInSet *LnsSet
    builtInSet = NewLnsSet([]LnsAny{})
    builtInSet.Add("require")
    for _key := range( AsyncParser_luaKeywordSet.Items ) {
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
        keywordSet.Add("__request")
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

// 203: decl @lune.@base.@AsyncParser.isOp2
func AsyncParser_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncParser_op2Set.Has(ope)
}

// 207: decl @lune.@base.@AsyncParser.isOp1
func AsyncParser_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncParser_op1Set.Has(ope)
}

// 216: decl @lune.@base.@AsyncParser.setDefaultPipeSize
func AsyncParser_setDefaultPipeSize(_env *LnsEnv, size LnsInt) {
    AsyncParser_defaultPipeSize = size
}

// 343: decl @lune.@base.@AsyncParser.create
func AsyncParser_create(_env *LnsEnv, parserSrc LnsAny,stdinFile LnsAny,overridePos LnsAny,async bool)(LnsAny, string) {
    if async{
        var runner *AsyncParser_Runner
        runner = NewAsyncParser_Runner(_env, parserSrc, stdinFile, overridePos)
        return runner.FP.Get_parser(_env), runner.FP.Get_errMess(_env)
    }
    var parser LnsAny
    var mess string
    parser,mess = AsyncParser_Parser_create_3_(_env, parserSrc, stdinFile, overridePos)
    if parser != nil{
        parser_195 := parser.(*AsyncParser_Parser)
        parser_195.FP.Stop(_env)
    }
    return parser, mess
}





// declaration Class -- AsyncItem
type AsyncParser_AsyncItemMtd interface {
    ToMap() *LnsMap
}
type AsyncParser_AsyncItem struct {
    List *LnsList
    FP AsyncParser_AsyncItemMtd
}
func AsyncParser_AsyncItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncParser_AsyncItem).FP
}
type AsyncParser_AsyncItemDownCast interface {
    ToAsyncParser_AsyncItem() *AsyncParser_AsyncItem
}
func AsyncParser_AsyncItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncParser_AsyncItemDownCast)
    if ok { return work.ToAsyncParser_AsyncItem() }
    return nil
}
func (obj *AsyncParser_AsyncItem) ToAsyncParser_AsyncItem() *AsyncParser_AsyncItem {
    return obj
}
func NewAsyncParser_AsyncItem(_env *LnsEnv, arg1 *LnsList) *AsyncParser_AsyncItem {
    obj := &AsyncParser_AsyncItem{}
    obj.FP = obj
    obj.InitAsyncParser_AsyncItem(_env, arg1)
    return obj
}
func (self *AsyncParser_AsyncItem) InitAsyncParser_AsyncItem(_env *LnsEnv, arg1 *LnsList) {
    self.List = arg1
}
func (self *AsyncParser_AsyncItem) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["list"] = Lns_ToCollection( self.List )
    return obj
}
func (self *AsyncParser_AsyncItem) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func AsyncParser_AsyncItem__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncParser_AsyncItem_FromMap( arg1, paramList )
}
func AsyncParser_AsyncItem__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncParser_AsyncItem_FromMap( arg1, paramList )
}
func AsyncParser_AsyncItem_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := AsyncParser_AsyncItem_FromMapSub(obj,false, paramList);
    return conv,mess
}
func AsyncParser_AsyncItem_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &AsyncParser_AsyncItem{}
    newObj.FP = newObj
    return AsyncParser_AsyncItem_FromMapMain( newObj, objMap, paramList )
}
func AsyncParser_AsyncItem_FromMapMain( newObj *AsyncParser_AsyncItem, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToListSub( objMap.Items["list"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Types_Token_FromMapSub, false,nil}}); !ok {
       return false,nil,"list:" + mess.(string)
    } else {
       newObj.List = conv.(*LnsList)
    }
    return true, newObj, nil
}
func AsyncParser_AsyncItem__createPipe(_env *LnsEnv, arg1 LnsInt) LnsAny{
   return NewLnspipe( arg1 )
}

// declaration Class -- Parser
type AsyncParser_ParserMtd interface {
    Access(_env *LnsEnv) LnsAny
    addVal(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string, arg4 LnsInt)
    analyzeNumber(_env *LnsEnv, arg1 string, arg2 LnsInt)(LnsInt, bool)
    createInfo(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt) *Types_Token
    GetNext(_env *LnsEnv) LnsAny
    Get_streamName(_env *LnsEnv) string
    Parse(_env *LnsEnv) LnsAny
    readLine(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
    Setup(_env *LnsEnv)
    Start(_env *LnsEnv)
    Stop(_env *LnsEnv)
}
type AsyncParser_Parser struct {
    Async_Pipe
    streamName string
    lineNo LnsInt
    prevToken *Types_Token
    keywordSet *LnsSet
    typeSet *LnsSet
    multiCharDelimitMap *LnsMap
    luaMode bool
    lineList *LnsList
    firstLine bool
    overridePos LnsAny
    stream Lns_iStream
    FP AsyncParser_ParserMtd
}
func AsyncParser_Parser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncParser_Parser).FP
}
type AsyncParser_ParserDownCast interface {
    ToAsyncParser_Parser() *AsyncParser_Parser
}
func AsyncParser_ParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncParser_ParserDownCast)
    if ok { return work.ToAsyncParser_Parser() }
    return nil
}
func (obj *AsyncParser_Parser) ToAsyncParser_Parser() *AsyncParser_Parser {
    return obj
}
func NewAsyncParser_Parser(_env *LnsEnv, arg1 string, arg2 Lns_iStream, arg3 bool, arg4 LnsAny, arg5 LnsAny) *AsyncParser_Parser {
    obj := &AsyncParser_Parser{}
    obj.FP = obj
    obj.Async_Pipe.FP = obj
    obj.InitAsyncParser_Parser(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *AsyncParser_Parser) Get_streamName(_env *LnsEnv) string{ return self.streamName }
// 233: DeclConstr
func (self *AsyncParser_Parser) InitAsyncParser_Parser(_env *LnsEnv, streamName string,stream Lns_iStream,luaMode bool,overridePos LnsAny,pipeSize LnsAny) {
    self.InitAsync_Pipe(_env, AsyncParser_AsyncItem__createPipe(_env, Lns_unwrapDefault( pipeSize, AsyncParser_defaultPipeSize).(LnsInt)))
    self.stream = stream
    self.lineList = NewLnsList([]LnsAny{})
    self.streamName = streamName
    self.overridePos = overridePos
    self.firstLine = true
    self.lineNo = 0
    self.prevToken = Types_noneToken
    self.luaMode = luaMode
    var keywordSet *LnsSet
    var typeSet *LnsSet
    var multiCharDelimitMap *LnsMap
    keywordSet,typeSet,_,multiCharDelimitMap = AsyncParser_createReserveInfo_1_(_env, luaMode)
    self.keywordSet = keywordSet
    self.typeSet = typeSet
    self.multiCharDelimitMap = multiCharDelimitMap
}

// 254: decl @lune.@base.@AsyncParser.Parser.setup
func (self *AsyncParser_Parser) Setup(_env *LnsEnv) {
    var lineList *LnsList
    lineList = NewLnsList([]LnsAny{})
    for  {
        var line string
        
        {
            _line := self.stream.Read(_env, "*l")
            if _line == nil{
                break
            } else {
                line = _line.(string)
            }
        }
        lineList.Insert(line)
    }
    self.lineList = lineList
    self.stream.Close(_env)
}

// 266: decl @lune.@base.@AsyncParser.Parser.create
func AsyncParser_Parser_create_3_(_env *LnsEnv, parserSrc LnsAny,stdinFile LnsAny,overridePos LnsAny)(LnsAny, string) {
    var createStream func(_env *LnsEnv, mod string,path string)(LnsAny, string)
    createStream = func(_env *LnsEnv, mod string,path string)(LnsAny, string) {
        if stdinFile != nil{
            stdinFile_151 := stdinFile.(*Types_StdinFile)
            if stdinFile_151.FP.Get_mod(_env) == mod{
                return NewUtil_TxtStream(_env, stdinFile_151.FP.Get_txt(_env)).FP, ""
            }
        }
        {
            __exp := AsyncParser_convExp0_1486(Lns_2DDD(Lns_io_open(path, "r")))
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(Lns_luaStream)
                return _exp, ""
            }
        }
        return nil, _env.GetVM().String_format("failed to open -- %s", []LnsAny{path})
    }
    var createStreamFrom func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny)
    createStreamFrom = func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny) {
        switch _matchExp1 := parserSrc.(type) {
        case *Types_ParserSrc__LnsCode:
        txt := _matchExp1.Val1
        path := _matchExp1.Val2
        pipeSize := _matchExp1.Val3
            return path, false, NewUtil_TxtStream(_env, txt).FP, "", pipeSize
        case *Types_ParserSrc__LnsPath:
        path := _matchExp1.Val1
        mod := _matchExp1.Val2
        pipeSize := _matchExp1.Val3
            var stream LnsAny
            var mess string
            stream,mess = createStream(_env, mod, path)
            return path, false, stream, mess, pipeSize
        case *Types_ParserSrc__Parser:
        path := _matchExp1.Val1
        luaMode := _matchExp1.Val2
        mod := _matchExp1.Val3
        pipeSize := _matchExp1.Val4
            var stream LnsAny
            var mess string
            stream,mess = createStream(_env, mod, path)
            return path, luaMode, stream, mess, pipeSize
        }
    // insert a dummy
        return "",false,nil,"",nil
    }
    var streamName string
    var luaMode bool
    var stream LnsAny
    var mess string
    var pipeSize LnsAny
    streamName,luaMode,stream,mess,pipeSize = createStreamFrom(_env)
    if stream != nil{
        stream_181 := stream.(Lns_iStream)
        return NewAsyncParser_Parser(_env, streamName, stream_181, luaMode, overridePos, pipeSize), ""
    }
    return nil, mess
}


// 308: decl @lune.@base.@AsyncParser.Parser.access
func (self *AsyncParser_Parser) Access(_env *LnsEnv) LnsAny {
    var tokenList *LnsList
    
    {
        _tokenList := self.FP.Parse(_env)
        if _tokenList == nil{
            return nil
        } else {
            tokenList = _tokenList.(*LnsList)
        }
    }
    return NewAsync_PipeItem(_env, AsyncParser_AsyncItem2Stem(NewAsyncParser_AsyncItem(_env, tokenList)))
}

// 361: decl @lune.@base.@AsyncParser.Parser.createInfo
func (self *AsyncParser_Parser) createInfo(_env *LnsEnv, tokenKind LnsInt,token string,tokenColumn LnsInt) *Types_Token {
    if tokenKind == Types_TokenKind__Symb{
        if self.keywordSet.Has(token){
            tokenKind = Types_TokenKind__Kywd
        } else if self.typeSet.Has(token){
            tokenKind = Types_TokenKind__Type
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( AsyncParser_op2Set.Has(token)) ||
            _env.SetStackVal( AsyncParser_op1Set.Has(token)) ).(bool){
            tokenKind = Types_TokenKind__Ope
        }
    }
    var consecutive bool
    consecutive = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.prevToken.Pos.LineNo == self.lineNo) &&
        _env.SetStackVal( self.prevToken.Pos.Column + len(self.prevToken.Txt) == tokenColumn) ).(bool)){
        consecutive = true
    }
    var newToken *Types_Token
    newToken = NewTypes_Token(_env, tokenKind, token, Types_Position_create(_env, self.lineNo, tokenColumn, self.streamName, self.overridePos), consecutive, NewLnsList([]LnsAny{}))
    self.prevToken = newToken
    return newToken
}

// 399: decl @lune.@base.@AsyncParser.Parser.analyzeNumber
func (self *AsyncParser_Parser) analyzeNumber(_env *LnsEnv, token string,beginIndex LnsInt)(LnsInt, bool) {
    var nonNumIndex LnsInt
    
    {
        _nonNumIndex := AsyncParser_convExp0_1911(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", beginIndex, nil)))
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
            _nonNumIndex = AsyncParser_convExp0_1962(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
            if Lns_IsNil( _nonNumIndex ) {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( nonNumChar == 88) ||
        _env.SetStackVal( nonNumChar == 120) ).(bool){
        {
            var _nonNumIndex LnsAny
            _nonNumIndex = AsyncParser_convExp0_2013(Lns_2DDD(_env.GetVM().String_find(token,"[^%da-fA-F]", nonNumIndex + 1, nil)))
            if Lns_IsNil( _nonNumIndex ) {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( nonNumChar == 69) ||
        _env.SetStackVal( nonNumChar == 101) ).(bool){
        intFlag = false
        var nextChar LnsInt
        nextChar = LnsInt(token[nonNumIndex + 1-1])
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nextChar == 45) ||
            _env.SetStackVal( nextChar == 43) ).(bool){
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncParser_convExp0_2093(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 2, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        } else { 
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncParser_convExp0_2123(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        }
    }
    return nonNumIndex - 1, intFlag
}

// 440: decl @lune.@base.@AsyncParser.Parser.readLine
func (self *AsyncParser_Parser) readLine(_env *LnsEnv) LnsAny {
    if self.lineNo >= self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    return self.lineList.GetAt(self.lineNo).(string)
}

// 459: decl @lune.@base.@AsyncParser.Parser.addVal
func (self *AsyncParser_Parser) addVal(_env *LnsEnv, list *LnsList,kind LnsInt,val string,column LnsInt) {
    if kind != Types_TokenKind__Symb{
        list.Insert(Types_Token2Stem(self.FP.createInfo(_env, kind, val, column)))
        return 
    }
    var searchIndex LnsInt
    searchIndex = 1
    for  {
        var tokenIndex LnsInt
        var tokenEndIndex LnsInt
        
        {
            _tokenIndex, _tokenEndIndex := AsyncParser_convExp0_2225(Lns_2DDD(_env.GetVM().String_find(val, "[%p%w]+", searchIndex, nil)))
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
        token = _env.GetVM().String_sub(val,tokenIndex, tokenEndIndex)
        var subIndex LnsInt
        subIndex = 1
        for len(token) >= subIndex {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_car(_env.GetVM().String_find(token,"^[%d]", subIndex, nil))) ||
                _env.SetStackVal( LnsInt(token[subIndex-1]) == 45) &&
                _env.SetStackVal( Lns_car(_env.GetVM().String_find(token,"^[%d]", subIndex + 1, nil))) )){
                var checkIndex LnsInt
                checkIndex = subIndex
                if LnsInt(token[checkIndex-1]) == 45{
                    checkIndex = checkIndex + 1
                }
                var endIndex LnsInt
                var intFlag bool
                endIndex,intFlag = self.FP.analyzeNumber(_env, token, checkIndex)
                var info *Types_Token
                info = self.FP.createInfo(_env, _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( intFlag) &&
                    _env.SetStackVal( Types_TokenKind__Int) ||
                    _env.SetStackVal( Types_TokenKind__Real) ).(LnsInt), _env.GetVM().String_sub(token,subIndex, endIndex), columnIndex + subIndex)
                list.Insert(Types_Token2Stem(info))
                subIndex = endIndex + 1
            } else { 
                {
                    __exp := AsyncParser_convExp0_2727(Lns_2DDD(_env.GetVM().String_find(token, "[^%w_]", subIndex, nil)))
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(LnsInt)
                        var index LnsInt
                        index = _exp
                        if index > subIndex{
                            var info *Types_Token
                            info = self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, index - 1), columnIndex + subIndex)
                            list.Insert(Types_Token2Stem(info))
                        }
                        var delimit string
                        delimit = _env.GetVM().String_sub(token,index, index)
                        var candidateList LnsAny
                        candidateList = self.multiCharDelimitMap.Get(delimit)
                        for Lns_isCondTrue( candidateList) {
                            var findFlag bool
                            findFlag = false
                            for _, _candidate := range( Lns_unwrap( (candidateList)).(*LnsList).Items ) {
                                candidate := _candidate.(string)
                                if candidate == _env.GetVM().String_sub(token,index, index + len(candidate) - 1){
                                    delimit = candidate
                                    candidateList = self.multiCharDelimitMap.Get(delimit)
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
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( AsyncParser_op2Set.Has(delimit)) ||
                            _env.SetStackVal( AsyncParser_op1Set.Has(delimit)) ).(bool){
                            workKind = Types_TokenKind__Ope
                        }
                        if delimit == "..."{
                            workKind = Types_TokenKind__Symb
                        }
                        if delimit == "?"{
                            var nextChar string
                            nextChar = _env.GetVM().String_sub(token,index, subIndex)
                            list.Insert(Types_Token2Stem(self.FP.createInfo(_env, Types_TokenKind__Char, nextChar, columnIndex + subIndex)))
                            subIndex = subIndex + 1
                        } else { 
                            list.Insert(Types_Token2Stem(self.FP.createInfo(_env, workKind, delimit, columnIndex + index)))
                        }
                    } else {
                        if subIndex <= len(token){
                            list.Insert(Types_Token2Stem(self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, nil), columnIndex + subIndex)))
                        }
                        break
                    }
                }
            }
        }
    }
}

// 557: decl @lune.@base.@AsyncParser.Parser.parse
func (self *AsyncParser_Parser) Parse(_env *LnsEnv) LnsAny {
    var rawLine string
    
    {
        _rawLine := self.FP.readLine(_env)
        if _rawLine == nil{
            return nil
        } else {
            rawLine = _rawLine.(string)
        }
    }
    if self.firstLine{
        self.firstLine = false
        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(rawLine,"^#!", nil, nil))){
            var token *Types_Token
            token = NewTypes_Token(_env, Types_TokenKind__Sheb, rawLine, NewTypes_Position(_env, self.lineNo, 1, self.streamName), false, NewLnsList([]LnsAny{}))
            return NewLnsList([]LnsAny{Types_Token2Stem(token)})
        }
    }
    var multiComment func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt)
    multiComment = func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt) {
        var searchIndex LnsInt
        searchIndex = comIndex
        var comment string
        comment = ""
        for  {
            {
                _, _termEndIndex := AsyncParser_convExp0_2873(Lns_2DDD(_env.GetVM().String_find(rawLine, termStr, searchIndex, true)))
                if !Lns_IsNil( _termEndIndex ) {
                    termEndIndex := _termEndIndex.(LnsInt)
                    comment = comment + _env.GetVM().String_sub(rawLine,searchIndex, termEndIndex)
                    return comment, termEndIndex + 1
                }
            }
            comment = comment + _env.GetVM().String_sub(rawLine,searchIndex, nil) + "\n"
            searchIndex = 1
            rawLine = Lns_unwrap( self.FP.readLine(_env)).(string)
        }
    // insert a dummy
        return "",0
    }
    var startIndex LnsInt
    startIndex = 1
    var searchIndex LnsInt
    searchIndex = startIndex
    var getChar func(_env *LnsEnv, index LnsInt) LnsInt
    getChar = func(_env *LnsEnv, index LnsInt) LnsInt {
        if len(rawLine) >= index{
            return LnsInt(rawLine[index-1])
        }
        return 0
    }
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for  {
        var syncIndexFlag bool
        syncIndexFlag = true
        var pattern string
        pattern = "[/%-%?\"%'%`]."
        var index LnsInt
        
        {
            _index := AsyncParser_convExp0_3018(Lns_2DDD(_env.GetVM().String_find(rawLine, pattern, searchIndex, nil)))
            if _index == nil{
                self.FP.addVal(_env, list, Types_TokenKind__Symb, _env.GetVM().String_sub(rawLine,startIndex, nil), startIndex)
                return list
            } else {
                index = _index.(LnsInt)
            }
        }
        var findChar LnsInt
        findChar = getChar(_env, index)
        var nextChar LnsInt
        nextChar = getChar(_env, index + 1)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( findChar == 45) &&
            _env.SetStackVal( nextChar != 45) ).(bool)){
            searchIndex = index + 1
            syncIndexFlag = false
        } else { 
            if startIndex < index{
                self.FP.addVal(_env, list, Types_TokenKind__Symb, _env.GetVM().String_sub(rawLine,startIndex, index - 1), startIndex)
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( findChar == 39) ||
                _env.SetStackVal( findChar == 34) ).(bool){
                var workIndex LnsInt
                workIndex = index + 1
                var workPattern string
                workPattern = "[\"'\\]"
                for  {
                    var endIndex LnsInt
                    
                    {
                        _endIndex := AsyncParser_convExp0_3163(Lns_2DDD(_env.GetVM().String_find(rawLine, workPattern, workIndex, nil)))
                        if _endIndex == nil{
                            Util_err(_env, _env.GetVM().String_format("%s:%d:%d: error: illegal string -- %s", []LnsAny{self.streamName, self.lineNo, index, rawLine}))
                        } else {
                            endIndex = _endIndex.(LnsInt)
                        }
                    }
                    var workChar LnsInt
                    workChar = LnsInt(rawLine[endIndex-1])
                    if workChar == findChar{
                        self.FP.addVal(_env, list, Types_TokenKind__Str, _env.GetVM().String_sub(rawLine,index, endIndex), index)
                        searchIndex = endIndex + 1
                        break
                    } else if workChar == 92{
                        workIndex = endIndex + 2
                    } else { 
                        workIndex = endIndex + 1
                    }
                }
            } else if findChar == 96{
                if Lns_isCondTrue( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( nextChar == findChar) &&
                    _env.SetStackVal( LnsInt(rawLine[index + 2-1]) == 96) ).(bool))){
                    var txt string
                    var nextIndex LnsInt
                    txt,nextIndex = multiComment(_env, index + 3, "```")
                    self.FP.addVal(_env, list, Types_TokenKind__Str, "```" + txt, index)
                    searchIndex = nextIndex
                } else { 
                    self.FP.addVal(_env, list, Types_TokenKind__Ope, "`", index)
                    searchIndex = index + 1
                }
            } else if findChar == 63{
                var codeChar string
                codeChar = _env.GetVM().String_sub(rawLine,index + 1, index + 1)
                if nextChar == 92{
                    var quoted string
                    quoted = _env.GetVM().String_sub(rawLine,index + 2, index + 2)
                    if AsyncParser_quotedCharSet.Has(quoted){
                        codeChar = _env.GetVM().String_sub(rawLine,index + 1, index + 2)
                    } else { 
                        codeChar = quoted
                    }
                    searchIndex = index + 3
                } else { 
                    searchIndex = index + 2
                }
                self.FP.addVal(_env, list, Types_TokenKind__Char, codeChar, index)
            } else { 
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( self.luaMode) &&
                    _env.SetStackVal( findChar == 45) &&
                    _env.SetStackVal( nextChar == 45) ).(bool)){
                    self.FP.addVal(_env, list, Types_TokenKind__Cmnt, _env.GetVM().String_sub(rawLine,index, nil), index)
                    searchIndex = len(rawLine) + 1
                } else if findChar == 47{
                    if nextChar == 42{
                        var comment string
                        var nextIndex LnsInt
                        comment,nextIndex = multiComment(_env, index + 2, "*/")
                        self.FP.addVal(_env, list, Types_TokenKind__Cmnt, "/*" + comment, index)
                        searchIndex = nextIndex
                    } else if nextChar == 47{
                        self.FP.addVal(_env, list, Types_TokenKind__Cmnt, _env.GetVM().String_sub(rawLine,index, nil), index)
                        searchIndex = len(rawLine) + 1
                    } else { 
                        self.FP.addVal(_env, list, Types_TokenKind__Ope, "/", index)
                        searchIndex = index + 1
                    }
                } else { 
                    Util_err(_env, _env.GetVM().String_format("%s:%d:%d: error: illegal syntax -- %s", []LnsAny{self.streamName, self.lineNo, index, rawLine}))
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


// declaration Class -- Runner
type AsyncParser_RunnerMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Get_errMess(_env *LnsEnv) string
    Get_parser(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
}
type AsyncParser_Runner struct {
    _syncFlag *Lns_syncFlag
    parser LnsAny
    errMess string
    FP AsyncParser_RunnerMtd
}
func (self *AsyncParser_Runner) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
func AsyncParser_Runner2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncParser_Runner).FP
}
type AsyncParser_RunnerDownCast interface {
    ToAsyncParser_Runner() *AsyncParser_Runner
}
func AsyncParser_RunnerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncParser_RunnerDownCast)
    if ok { return work.ToAsyncParser_Runner() }
    return nil
}
func (obj *AsyncParser_Runner) ToAsyncParser_Runner() *AsyncParser_Runner {
    return obj
}
func NewAsyncParser_Runner(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) *AsyncParser_Runner {
    obj := &AsyncParser_Runner{}
    obj.FP = obj
    obj.InitAsyncParser_Runner(_env, arg1, arg2, arg3)
    return obj
}
func (self *AsyncParser_Runner) Get_parser(_env *LnsEnv) LnsAny{ return self.parser }
func (self *AsyncParser_Runner) Get_errMess(_env *LnsEnv) string{ return self.errMess }
// 321: DeclConstr
func (self *AsyncParser_Runner) InitAsyncParser_Runner(_env *LnsEnv, parserSrc LnsAny,stdinFile LnsAny,overridePos LnsAny) {
    self._syncFlag = &Lns_syncFlag{}
    self.parser, self.errMess = AsyncParser_Parser_create_3_(_env, parserSrc, stdinFile, overridePos)
    {
        __exp := self.parser
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncParser_Parser)
            _exp.FP.Start(_env)
            if Lns_op_not(LnsRun(_env, self.FP, 2, _env.GetVM().String_format("parser - %s", []LnsAny{_exp.FP.Get_streamName(_env)}))){
                _exp.FP.Stop(_env)
            }
        }
    }
}

// 336: decl @lune.@base.@AsyncParser.Runner.run
func (self *AsyncParser_Runner) Run(_env *LnsEnv) {
    {
        __exp := self.parser
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncParser_Parser)
            _exp.FP.Run(_env)
        }
    }
}


func Lns_AsyncParser_init(_env *LnsEnv) {
    if init_AsyncParser { return }
    init_AsyncParser = true
    AsyncParser__mod__ = "@lune.@base.@AsyncParser"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Types_init(_env)
    Lns_Async_init(_env)
    AsyncParser_luaKeywordSet = NewLnsSet([]LnsAny{"if", "else", "elseif", "while", "for", "in", "return", "break", "nil", "true", "false", "{", "}", "do", "require", "function", "then", "end", "repeat", "until", "goto", "local"})
    AsyncParser_quotedCharSet = NewLnsSet([]LnsAny{"a", "b", "f", "n", "r", "t", "v", "\\", "\"", "'"})
    AsyncParser_op2Set = NewLnsSet([]LnsAny{"+", "-", "*", "/", "^", "%", "&", "~", "|", "|>>", "|<<", "..", "<", "<=", ">", ">=", "==", "~=", "and", "or", "@", "@@", "@@@", "="})
    AsyncParser_op1Set = NewLnsSet([]LnsAny{"-", "not", "#", "~", "*", "`", ",,", ",,,", ",,,,"})
    AsyncParser_defaultPipeSize = 20
}
func init() {
    init_AsyncParser = false
}
