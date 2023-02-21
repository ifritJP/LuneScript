// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_AsyncParser bool
var AsyncParser__mod__ string
var AsyncParser_luaKeywordSet *LnsSet2_[string]
var AsyncParser_quotedCharSet *LnsSet2_[string]
var AsyncParser_op2Set *LnsSet2_[string]
var AsyncParser_op1Set *LnsSet2_[string]
var AsyncParser_defaultPipeSize LnsInt
// for 276
func AsyncParser_convExp0_1487(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 503
func AsyncParser_convExp0_2783(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 589
func AsyncParser_convExp0_2928(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 331
func AsyncParser_convExp0_522(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 247
func AsyncParser_convExp0_415(arg1 []LnsAny) (*LnsSet2_[string], *LnsSet2_[string], *LnsSet2_[string], *LnsMap) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 1 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 2 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 3 ).(*LnsMap)
}
// for 296
func AsyncParser_convExp0_1605(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 300
func AsyncParser_convExp0_1631(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 306
func AsyncParser_convExp0_1659(arg1 []LnsAny) (string, bool, LnsAny, string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(string), Lns_getFromMulti( arg1, 4 )
}
// for 357
func AsyncParser_convExp0_1772(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 407
func AsyncParser_convExp0_1967(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 415
func AsyncParser_convExp0_2018(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 423
func AsyncParser_convExp0_2069(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 434
func AsyncParser_convExp0_2149(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 439
func AsyncParser_convExp0_2179(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 477
func AsyncParser_convExp0_2281(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 494
func AsyncParser_convExp0_2399(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 620
func AsyncParser_convExp0_3071(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 643
func AsyncParser_convExp0_3216(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 667
func AsyncParser_convExp0_3335(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 705
func AsyncParser_convExp0_3591(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// 59: decl @lune.@base.@AsyncParser.isLuaKeyword
func AsyncParser_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncParser_luaKeywordSet.Has(txt)
}

// 64: decl @lune.@base.@AsyncParser.createReserveInfo
func AsyncParser_createReserveInfo_1_(_env *LnsEnv, luaMode LnsAny)(*LnsSet2_[string], *LnsSet2_[string], *LnsSet2_[string], *LnsMap) {
    var keywordSet *LnsSet2_[string]
    keywordSet = NewLnsSet2_[string]([]string{})
    var typeSet *LnsSet2_[string]
    typeSet = NewLnsSet2_[string]([]string{})
    var builtInSet *LnsSet2_[string]
    builtInSet = NewLnsSet2_[string]([]string{})
    builtInSet.Add("require")
    for _key := range( AsyncParser_luaKeywordSet.Items ) {
        key := _key
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
    multiCharDelimitMap.Set("=",NewLnsList(Lns_2DDD("==")))
    multiCharDelimitMap.Set("<",NewLnsList(Lns_2DDD("<=")))
    multiCharDelimitMap.Set(">",NewLnsList(Lns_2DDD(">=")))
    if Lns_op_not(luaMode){
        multiCharDelimitMap.Set("|",NewLnsList(Lns_2DDD("|<", "|>")))
        multiCharDelimitMap.Set("|<",NewLnsList(Lns_2DDD("|<<")))
        multiCharDelimitMap.Set("|>",NewLnsList(Lns_2DDD("|>>")))
        multiCharDelimitMap.Set("[",NewLnsList(Lns_2DDD("[@")))
        multiCharDelimitMap.Set("(",NewLnsList(Lns_2DDD("(@", "(=")))
        multiCharDelimitMap.Set("~",NewLnsList(Lns_2DDD("~=", "~~")))
        multiCharDelimitMap.Set("$",NewLnsList(Lns_2DDD("$[", "$.", "$(")))
        multiCharDelimitMap.Set("$.",NewLnsList(Lns_2DDD("$.$")))
        multiCharDelimitMap.Set(".",NewLnsList(Lns_2DDD("..", ".$")))
        multiCharDelimitMap.Set("..",NewLnsList(Lns_2DDD("...")))
        multiCharDelimitMap.Set(",",NewLnsList(Lns_2DDD(",,")))
        multiCharDelimitMap.Set(",,",NewLnsList(Lns_2DDD(",,,")))
        multiCharDelimitMap.Set(",,,",NewLnsList(Lns_2DDD(",,,,")))
        multiCharDelimitMap.Set("@",NewLnsList(Lns_2DDD("@@")))
        multiCharDelimitMap.Set("@@",NewLnsList(Lns_2DDD("@@@", "@@=")))
        multiCharDelimitMap.Set("#",NewLnsList(Lns_2DDD("##")))
        multiCharDelimitMap.Set("*",NewLnsList(Lns_2DDD("**")))
    } else { 
        multiCharDelimitMap.Set(".",NewLnsList(Lns_2DDD("..")))
        multiCharDelimitMap.Set("~",NewLnsList(Lns_2DDD("~=")))
    }
    return keywordSet, typeSet, builtInSet, multiCharDelimitMap
}

// 202: decl @lune.@base.@AsyncParser.isOp2
func AsyncParser_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncParser_op2Set.Has(ope)
}

// 206: decl @lune.@base.@AsyncParser.isOp1
func AsyncParser_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncParser_op1Set.Has(ope)
}

// 215: decl @lune.@base.@AsyncParser.setDefaultPipeSize
func AsyncParser_setDefaultPipeSize(_env *LnsEnv, size LnsInt) {
    AsyncParser_defaultPipeSize = size
}




// 350: decl @lune.@base.@AsyncParser.create
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
        parser_204 := parser.(*AsyncParser_Parser)
        parser_204.FP.Stop(_env)
    }
    return parser, mess
}



// 254: decl @lune.@base.@AsyncParser.Parser.setup
func (self *AsyncParser_Parser) Setup(_env *LnsEnv) {
    var lineList *LnsList2_[string]
    lineList = NewLnsList2_[string]([]string{})
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
    var AsyncParser_createStream func(_env *LnsEnv, mod string,path string)(LnsAny, string)
    AsyncParser_createStream = func(_env *LnsEnv, mod string,path string)(LnsAny, string) {
        __func__ := "@lune.@base.@AsyncParser.Parser.create.createStream"
        if stdinFile != nil{
            stdinFile_150 := stdinFile.(*Types_StdinFile)
            if stdinFile_150.FP.Get_mod(_env) == mod{
                return NewUtil_TxtStream(_env, stdinFile_150.FP.Get_txt(_env)).FP, ""
            }
        }
        {
            __exp := AsyncParser_convExp0_1487(Lns_2DDD(Lns_io_open(path, "r")))
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(Lns_luaStream)
                return _exp, ""
            }
        }
        return nil, _env.GetVM().String_format("%s: failed to open -- %s", Lns_2DDD(__func__, path))
    }
    var AsyncParser_createStreamWithBaseDir func(_env *LnsEnv, mod string,baseDir LnsAny,path string)(LnsAny, string)
    AsyncParser_createStreamWithBaseDir = func(_env *LnsEnv, mod string,baseDir LnsAny,path string)(LnsAny, string) {
        if baseDir != nil{
            baseDir_161 := baseDir.(string)
            return AsyncParser_createStream(_env, mod, Util_pathJoin(_env, baseDir_161, path))
        } else {
            return AsyncParser_createStream(_env, mod, path)
        }
    // insert a dummy
        return nil,""
    }
    var AsyncParser_createStreamFrom func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny)
    AsyncParser_createStreamFrom = func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny) {
        switch _matchExp0 := parserSrc.(type) {
        case *Types_ParserSrc__LnsCode:
            txt := _matchExp0.Val1
            path := _matchExp0.Val2
            pipeSize := _matchExp0.Val3
            return path, false, NewUtil_TxtStream(_env, txt).FP, "", pipeSize
        case *Types_ParserSrc__LnsPath:
            baseDir := _matchExp0.Val1
            path := _matchExp0.Val2
            mod := _matchExp0.Val3
            pipeSize := _matchExp0.Val4
            var stream LnsAny
            var mess string
            stream,mess = AsyncParser_createStreamWithBaseDir(_env, mod, baseDir, path)
            return path, false, stream, mess, pipeSize
        case *Types_ParserSrc__Parser:
            path := _matchExp0.Val1
            luaMode := _matchExp0.Val2
            mod := _matchExp0.Val3
            pipeSize := _matchExp0.Val4
            var stream LnsAny
            var mess string
            stream,mess = AsyncParser_createStream(_env, mod, path)
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
    streamName,luaMode,stream,mess,pipeSize = AsyncParser_createStreamFrom(_env)
    if stream != nil{
        stream_190 := stream.(Lns_iStream)
        return NewAsyncParser_Parser(_env, streamName, stream_190, luaMode, overridePos, pipeSize), ""
    }
    return nil, mess
}
// 315: decl @lune.@base.@AsyncParser.Parser.access
func (self *AsyncParser_Parser) Access(_env *LnsEnv) LnsAny {
    var tokenList *LnsList2_[*Types_Token]
    
    {
        _tokenList := self.FP.Parse(_env)
        if _tokenList == nil{
            return nil
        } else {
            tokenList = _tokenList.(*LnsList2_[*Types_Token])
        }
    }
    return NewAsync_PipeItem(_env, AsyncParser_AsyncItem2Stem(NewAsyncParser_AsyncItem(_env, tokenList)))
}
// 368: decl @lune.@base.@AsyncParser.Parser.createInfo
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
    newToken = NewTypes_Token(_env, tokenKind, token, Types_Position_create(_env, self.lineNo, tokenColumn, self.streamName, self.overridePos), consecutive, NewLnsList2_[*Types_Token]([]*Types_Token{}))
    self.prevToken = newToken
    return newToken
}
// 406: decl @lune.@base.@AsyncParser.Parser.analyzeNumber
func (self *AsyncParser_Parser) analyzeNumber(_env *LnsEnv, token string,beginIndex LnsInt)(LnsInt, bool) {
    var nonNumIndex LnsInt
    
    {
        _nonNumIndex := AsyncParser_convExp0_1967(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", beginIndex, nil)))
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
            _nonNumIndex = AsyncParser_convExp0_2018(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
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
            _nonNumIndex = AsyncParser_convExp0_2069(Lns_2DDD(_env.GetVM().String_find(token,"[^%da-fA-F]", nonNumIndex + 1, nil)))
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
                _nonNumIndex = AsyncParser_convExp0_2149(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 2, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        } else { 
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncParser_convExp0_2179(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        }
    }
    return nonNumIndex - 1, intFlag
}
// 447: decl @lune.@base.@AsyncParser.Parser.readLine
func (self *AsyncParser_Parser) readLine(_env *LnsEnv) LnsAny {
    if self.lineNo >= self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    return self.lineList.GetAt(self.lineNo)
}
// 466: decl @lune.@base.@AsyncParser.Parser.addVal
func (self *AsyncParser_Parser) addVal(_env *LnsEnv, list *LnsList2_[*Types_Token],kind LnsInt,val string,column LnsInt) {
    if kind != Types_TokenKind__Symb{
        list.Insert(self.FP.createInfo(_env, kind, val, column))
        return 
    }
    var searchIndex LnsInt
    searchIndex = 1
    for  {
        var tokenIndex LnsInt
        var tokenEndIndex LnsInt
        
        {
            _tokenIndex, _tokenEndIndex := AsyncParser_convExp0_2281(Lns_2DDD(_env.GetVM().String_find(val, "[%p%w]+", searchIndex, nil)))
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
                list.Insert(info)
                subIndex = endIndex + 1
            } else { 
                {
                    __exp := AsyncParser_convExp0_2783(Lns_2DDD(_env.GetVM().String_find(token, "[^%w_]", subIndex, nil)))
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(LnsInt)
                        var index LnsInt
                        index = _exp
                        if index > subIndex{
                            var info *Types_Token
                            info = self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, index - 1), columnIndex + subIndex)
                            list.Insert(info)
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
                            list.Insert(self.FP.createInfo(_env, Types_TokenKind__Char, nextChar, columnIndex + subIndex))
                            subIndex = subIndex + 1
                        } else { 
                            list.Insert(self.FP.createInfo(_env, workKind, delimit, columnIndex + index))
                        }
                    } else {
                        if subIndex <= len(token){
                            list.Insert(self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, nil), columnIndex + subIndex))
                        }
                        break
                    }
                }
            }
        }
    }
}
// 564: decl @lune.@base.@AsyncParser.Parser.parse
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
            token = NewTypes_Token(_env, Types_TokenKind__Sheb, rawLine, NewTypes_Position(_env, self.lineNo, 1, self.streamName, nil), false, NewLnsList2_[*Types_Token]([]*Types_Token{}))
            return NewLnsList2_[*Types_Token](Lns_2DDDGen[*Types_Token](token))
        }
    }
    var AsyncParser_multiComment func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt)
    AsyncParser_multiComment = func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt) {
        var searchIndex LnsInt
        searchIndex = comIndex
        var comment string
        comment = ""
        for  {
            {
                _, _termEndIndex := AsyncParser_convExp0_2928(Lns_2DDD(_env.GetVM().String_find(rawLine, termStr, searchIndex, true)))
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
    var AsyncParser_getChar func(_env *LnsEnv, index LnsInt) LnsInt
    AsyncParser_getChar = func(_env *LnsEnv, index LnsInt) LnsInt {
        if len(rawLine) >= index{
            return LnsInt(rawLine[index-1])
        }
        return 0
    }
    var list *LnsList2_[*Types_Token]
    list = NewLnsList2_[*Types_Token]([]*Types_Token{})
    for  {
        var syncIndexFlag bool
        syncIndexFlag = true
        var pattern string
        pattern = "[/%-%?\"%'%`]."
        var index LnsInt
        
        {
            _index := AsyncParser_convExp0_3071(Lns_2DDD(_env.GetVM().String_find(rawLine, pattern, searchIndex, nil)))
            if _index == nil{
                self.FP.addVal(_env, list, Types_TokenKind__Symb, _env.GetVM().String_sub(rawLine,startIndex, nil), startIndex)
                return list
            } else {
                index = _index.(LnsInt)
            }
        }
        var findChar LnsInt
        findChar = AsyncParser_getChar(_env, index)
        var nextChar LnsInt
        nextChar = AsyncParser_getChar(_env, index + 1)
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
                        _endIndex := AsyncParser_convExp0_3216(Lns_2DDD(_env.GetVM().String_find(rawLine, workPattern, workIndex, nil)))
                        if _endIndex == nil{
                            Util_err(_env, _env.GetVM().String_format("%s:%d:%d: error: illegal string -- %s", Lns_2DDD(self.streamName, self.lineNo, index, rawLine)))
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
                    _env.SetStackVal( AsyncParser_getChar(_env, index + 2) == 96) ).(bool))){
                    var txt string
                    var nextIndex LnsInt
                    txt,nextIndex = AsyncParser_multiComment(_env, index + 3, "```")
                    self.FP.addVal(_env, list, Types_TokenKind__Str, "```" + txt, index)
                    searchIndex = nextIndex
                } else if nextChar == 123{
                    self.FP.addVal(_env, list, Types_TokenKind__Ope, "`{", index)
                    searchIndex = index + 2
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
                        comment,nextIndex = AsyncParser_multiComment(_env, index + 2, "*/")
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
                    Util_err(_env, _env.GetVM().String_format("%s:%d:%d: error: illegal syntax -- %s", Lns_2DDD(self.streamName, self.lineNo, index, rawLine)))
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
// 343: decl @lune.@base.@AsyncParser.Runner.run
func (self *AsyncParser_Runner) Run(_env *LnsEnv) {
    {
        __exp := self.parser
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncParser_Parser)
            _exp.FP.Run(_env)
        }
    }
}
// declaration Class -- AsyncItem
type AsyncParser_AsyncItemMtd interface {
    ToMap() *LnsMap
}
type AsyncParser_AsyncItem struct {
    List *LnsList2_[*Types_Token]
    FP AsyncParser_AsyncItemMtd
}
func AsyncParser_AsyncItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncParser_AsyncItem).FP
}
func AsyncParser_AsyncItem_toSlice(slice []LnsAny) []*AsyncParser_AsyncItem {
    ret := make([]*AsyncParser_AsyncItem, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncParser_AsyncItemDownCast).ToAsyncParser_AsyncItem()
    }
    return ret
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
func NewAsyncParser_AsyncItem(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) *AsyncParser_AsyncItem {
    obj := &AsyncParser_AsyncItem{}
    obj.FP = obj
    obj.InitAsyncParser_AsyncItem(_env, arg1)
    return obj
}
func (self *AsyncParser_AsyncItem) InitAsyncParser_AsyncItem(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) {
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
    if ok,conv,mess := Lns_ToList2Sub[*Types_Token]( objMap.Items["list"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Types_Token_FromMapSub, false,nil}}); !ok {
       return false,nil,"list:" + mess.(string)
    } else {
       newObj.List = conv.(*LnsList2_[*Types_Token])
    }
    return true, newObj, nil
}

// declaration Class -- Parser
type AsyncParser_ParserMtd interface {
    Access(_env *LnsEnv) LnsAny
    addVal(_env *LnsEnv, arg1 *LnsList2_[*Types_Token], arg2 LnsInt, arg3 string, arg4 LnsInt)
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
    keywordSet *LnsSet2_[string]
    typeSet *LnsSet2_[string]
    multiCharDelimitMap *LnsMap
    luaMode bool
    lineList *LnsList2_[string]
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
func AsyncParser_Parser_toSlice(slice []LnsAny) []*AsyncParser_Parser {
    ret := make([]*AsyncParser_Parser, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncParser_ParserDownCast).ToAsyncParser_Parser()
    }
    return ret
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
// 232: DeclConstr
func (self *AsyncParser_Parser) InitAsyncParser_Parser(_env *LnsEnv, streamName string,stream Lns_iStream,luaMode bool,overridePos LnsAny,pipeSize LnsAny) {
    self.InitAsync_Pipe(_env, LnsAny(NewLnspipe( Lns_unwrapDefault( pipeSize, AsyncParser_defaultPipeSize).(LnsInt))))
    self.stream = stream
    self.lineList = NewLnsList2_[string]([]string{})
    self.streamName = streamName
    self.overridePos = overridePos
    self.firstLine = true
    self.lineNo = 0
    self.prevToken = Types_noneToken
    self.luaMode = luaMode
    var keywordSet *LnsSet2_[string]
    var typeSet *LnsSet2_[string]
    var multiCharDelimitMap *LnsMap
    keywordSet,typeSet,_,multiCharDelimitMap = AsyncParser_createReserveInfo_1_(_env, luaMode)
    self.keywordSet = keywordSet
    self.typeSet = typeSet
    self.multiCharDelimitMap = multiCharDelimitMap
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
func AsyncParser_Runner_toSlice(slice []LnsAny) []*AsyncParser_Runner {
    ret := make([]*AsyncParser_Runner, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncParser_RunnerDownCast).ToAsyncParser_Runner()
    }
    return ret
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
// 328: DeclConstr
func (self *AsyncParser_Runner) InitAsyncParser_Runner(_env *LnsEnv, parserSrc LnsAny,stdinFile LnsAny,overridePos LnsAny) {
    self._syncFlag = &Lns_syncFlag{}
    self.parser, self.errMess = AsyncParser_Parser_create_3_(_env, parserSrc, stdinFile, overridePos)
    {
        __exp := self.parser
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncParser_Parser)
            _exp.FP.Start(_env)
            if Lns_op_not(LnsRun(_env, self.FP, 2, _env.GetVM().String_format("parser - %s", Lns_2DDD(_exp.FP.Get_streamName(_env))))){
                _exp.FP.Stop(_env)
            }
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
    AsyncParser_luaKeywordSet = NewLnsSet2_[string](Lns_2DDDGen[string]("if", "else", "elseif", "while", "for", "in", "return", "break", "nil", "true", "false", "{", "}", "do", "require", "function", "then", "end", "repeat", "until", "goto", "local"))
    AsyncParser_quotedCharSet = NewLnsSet2_[string](Lns_2DDDGen[string]("a", "b", "f", "n", "r", "t", "v", "\\", "\"", "'"))
    AsyncParser_op2Set = NewLnsSet2_[string](Lns_2DDDGen[string]("+", "-", "*", "/", "^", "%", "&", "~", "|", "|>>", "|<<", "..", "<", "<=", ">", ">=", "==", "~=", "and", "or", "@", "@@", "@@@", "="))
    AsyncParser_op1Set = NewLnsSet2_[string](Lns_2DDDGen[string]("-", "not", "#", "~", "`{", ",,", ",,,", ",,,,"))
    AsyncParser_defaultPipeSize = 100
}
func init() {
    init_AsyncParser = false
}
