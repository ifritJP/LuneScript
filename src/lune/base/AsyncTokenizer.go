// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_AsyncTokenizer bool
var AsyncTokenizer__mod__ string
var AsyncTokenizer_luaKeywordSet *LnsSet2_[string]
var AsyncTokenizer_quotedCharSet *LnsSet2_[string]
var AsyncTokenizer_op2Set *LnsSet2_[string]
var AsyncTokenizer_op1Set *LnsSet2_[string]
var AsyncTokenizer_defaultPipeSize LnsInt
// for 282
func AsyncTokenizer_convExp0_1503(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 528
func AsyncTokenizer_convExp0_2931(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 614
func AsyncTokenizer_convExp0_3076(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 337
func AsyncTokenizer_convExp0_532(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 253
func AsyncTokenizer_convExp0_425(arg1 []LnsAny) (*LnsSet2_[string], *LnsSet2_[string], *LnsSet2_[string], *LnsMap2_[string,*LnsList]) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 1 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 2 ).(*LnsSet2_[string]), Lns_getFromMulti( arg1, 3 ).(*LnsMap2_[string,*LnsList])
}
// for 302
func AsyncTokenizer_convExp0_1621(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 306
func AsyncTokenizer_convExp0_1647(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 312
func AsyncTokenizer_convExp0_1675(arg1 []LnsAny) (string, bool, LnsAny, string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(string), Lns_getFromMulti( arg1, 4 )
}
// for 363
func AsyncTokenizer_convExp0_1788(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 432
func AsyncTokenizer_convExp0_2109(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 440
func AsyncTokenizer_convExp0_2160(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 448
func AsyncTokenizer_convExp0_2211(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 459
func AsyncTokenizer_convExp0_2291(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 464
func AsyncTokenizer_convExp0_2321(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 502
func AsyncTokenizer_convExp0_2424(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 519
func AsyncTokenizer_convExp0_2542(arg1 []LnsAny) (LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 645
func AsyncTokenizer_convExp0_3219(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 668
func AsyncTokenizer_convExp0_3364(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 693
func AsyncTokenizer_convExp0_3489(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 733
func AsyncTokenizer_convExp0_3757(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// 60: decl @lune.@base.@AsyncTokenizer.isLuaKeyword
func AsyncTokenizer_isLuaKeyword(_env *LnsEnv, txt string) bool {
    return AsyncTokenizer_luaKeywordSet.Has(txt)
}

// 65: decl @lune.@base.@AsyncTokenizer.createReserveInfo
func AsyncTokenizer_createReserveInfo_1_(_env *LnsEnv, luaMode LnsAny)(*LnsSet2_[string], *LnsSet2_[string], *LnsSet2_[string], *LnsMap2_[string,*LnsList]) {
    var keywordSet *LnsSet2_[string]
    keywordSet = NewLnsSet2_[string]([]string{})
    var typeSet *LnsSet2_[string]
    typeSet = NewLnsSet2_[string]([]string{})
    var builtInSet *LnsSet2_[string]
    builtInSet = NewLnsSet2_[string]([]string{})
    builtInSet.Add("require")
    for _key := range( AsyncTokenizer_luaKeywordSet.Items ) {
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
    var multiCharDelimitMap *LnsMap2_[string,*LnsList]
    multiCharDelimitMap = NewLnsMap2_[string,*LnsList]( map[string]*LnsList{})
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

// 203: decl @lune.@base.@AsyncTokenizer.isOp2
func AsyncTokenizer_isOp2(_env *LnsEnv, ope string) bool {
    return AsyncTokenizer_op2Set.Has(ope)
}

// 207: decl @lune.@base.@AsyncTokenizer.isOp1
func AsyncTokenizer_isOp1(_env *LnsEnv, ope string) bool {
    return AsyncTokenizer_op1Set.Has(ope)
}

// 216: decl @lune.@base.@AsyncTokenizer.setDefaultPipeSize
func AsyncTokenizer_setDefaultPipeSize(_env *LnsEnv, size LnsInt) {
    AsyncTokenizer_defaultPipeSize = size
}




// 356: decl @lune.@base.@AsyncTokenizer.create
func AsyncTokenizer_create(_env *LnsEnv, tokenizerSrc LnsAny,stdinFile LnsAny,overridePos LnsAny,async bool)(LnsAny, string) {
    if async{
        var runner *AsyncTokenizer_Runner
        runner = NewAsyncTokenizer_Runner(_env, tokenizerSrc, stdinFile, overridePos)
        return runner.FP.Get_tokenizer(_env), runner.FP.Get_errMess(_env)
    }
    var tokenizer LnsAny
    var mess string
    tokenizer,mess = AsyncTokenizer_Tokenizer_create_3_(_env, tokenizerSrc, stdinFile, overridePos)
    if tokenizer != nil{
        tokenizer_213 := tokenizer.(*AsyncTokenizer_Tokenizer)
        tokenizer_213.FP.Stop(_env)
    }
    return tokenizer, mess
}



// 260: decl @lune.@base.@AsyncTokenizer.Tokenizer.setup
func (self *AsyncTokenizer_Tokenizer) Setup(_env *LnsEnv) {
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
// 272: decl @lune.@base.@AsyncTokenizer.Tokenizer.create
func AsyncTokenizer_Tokenizer_create_3_(_env *LnsEnv, tokenizerSrc LnsAny,stdinFile LnsAny,overridePos LnsAny)(LnsAny, string) {
    var AsyncTokenizer_createStream func(_env *LnsEnv, mod string,path string)(LnsAny, string)
    AsyncTokenizer_createStream = func(_env *LnsEnv, mod string,path string)(LnsAny, string) {
        __func__ := "@lune.@base.@AsyncTokenizer.Tokenizer.create.createStream"
        if stdinFile != nil{
            stdinFile_159 := stdinFile.(*Types_StdinFile)
            if stdinFile_159.FP.Get_mod(_env) == mod{
                return NewUtil_TxtStream(_env, stdinFile_159.FP.Get_txt(_env)).FP, ""
            }
        }
        {
            __exp := AsyncTokenizer_convExp0_1503(Lns_2DDD(Lns_io_open(path, "r")))
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(Lns_luaStream)
                return _exp, ""
            }
        }
        return nil, _env.GetVM().String_format("%s: failed to open -- %s", Lns_2DDD(__func__, path))
    }
    var AsyncTokenizer_createStreamWithBaseDir func(_env *LnsEnv, mod string,baseDir LnsAny,path string)(LnsAny, string)
    AsyncTokenizer_createStreamWithBaseDir = func(_env *LnsEnv, mod string,baseDir LnsAny,path string)(LnsAny, string) {
        if baseDir != nil{
            baseDir_170 := baseDir.(string)
            return AsyncTokenizer_createStream(_env, mod, Util_pathJoin(_env, baseDir_170, path))
        } else {
            return AsyncTokenizer_createStream(_env, mod, path)
        }
    // insert a dummy
        return nil,""
    }
    var AsyncTokenizer_createStreamFrom func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny)
    AsyncTokenizer_createStreamFrom = func(_env *LnsEnv)(string, bool, LnsAny, string, LnsAny) {
        switch _matchExp0 := tokenizerSrc.(type) {
        case *Types_TokenizerSrc__LnsCode:
            txt := _matchExp0.Val1
            path := _matchExp0.Val2
            pipeSize := _matchExp0.Val3
            return path, false, NewUtil_TxtStream(_env, txt).FP, "", pipeSize
        case *Types_TokenizerSrc__LnsPath:
            baseDir := _matchExp0.Val1
            path := _matchExp0.Val2
            mod := _matchExp0.Val3
            pipeSize := _matchExp0.Val4
            var stream LnsAny
            var mess string
            stream,mess = AsyncTokenizer_createStreamWithBaseDir(_env, mod, baseDir, path)
            return path, false, stream, mess, pipeSize
        case *Types_TokenizerSrc__Tokenizer:
            path := _matchExp0.Val1
            luaMode := _matchExp0.Val2
            mod := _matchExp0.Val3
            pipeSize := _matchExp0.Val4
            var stream LnsAny
            var mess string
            stream,mess = AsyncTokenizer_createStream(_env, mod, path)
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
    streamName,luaMode,stream,mess,pipeSize = AsyncTokenizer_createStreamFrom(_env)
    if stream != nil{
        stream_199 := stream.(Lns_iStream)
        return NewAsyncTokenizer_Tokenizer(_env, streamName, stream_199, luaMode, overridePos, pipeSize), ""
    }
    return nil, mess
}
// 321: decl @lune.@base.@AsyncTokenizer.Tokenizer.access
func (self *AsyncTokenizer_Tokenizer) Access(_env *LnsEnv) LnsAny {
    var tokenList *LnsList2_[*Types_Token]
    
    {
        _tokenList := self.FP.Parse(_env)
        if _tokenList == nil{
            return nil
        } else {
            tokenList = _tokenList.(*LnsList2_[*Types_Token])
        }
    }
    return NewAsync_PipeItem(_env, AsyncTokenizer_AsyncItem2Stem(NewAsyncTokenizer_AsyncItem(_env, tokenList)))
}
// 371: decl @lune.@base.@AsyncTokenizer.Tokenizer.createInfo
func (self *AsyncTokenizer_Tokenizer) createInfo(_env *LnsEnv, tokenKind LnsInt,token string,tokenColumn LnsInt,tokenLineNo LnsAny,endColumn LnsAny) *Types_Token {
    var lineNo LnsInt
    
    {
        _lineNo := tokenLineNo
        if _lineNo == nil{
            lineNo = self.lineNo
        } else {
            lineNo = _lineNo.(LnsInt)
        }
    }
    if tokenKind == Types_TokenKind__Symb{
        if self.keywordSet.Has(token){
            tokenKind = Types_TokenKind__Kywd
        } else if self.typeSet.Has(token){
            tokenKind = Types_TokenKind__Type
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( AsyncTokenizer_op2Set.Has(token)) ||
            _env.SetStackVal( AsyncTokenizer_op1Set.Has(token)) ).(bool){
            tokenKind = Types_TokenKind__Ope
        }
    }
    var consecutive bool
    consecutive = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.prevToken.Pos.LineNo == lineNo) &&
        _env.SetStackVal( self.prevToken.Pos.Column + len(self.prevToken.Txt) == tokenColumn) ).(bool)){
        consecutive = true
    } else if self.prevToken.Kind == Types_TokenKind__Str{
        {
            _multiLineToken := AsyncTokenizer_MultiLineTokenDownCastF(self.prevToken.FP)
            if !Lns_IsNil( _multiLineToken ) {
                multiLineToken := _multiLineToken.(*AsyncTokenizer_MultiLineToken)
                var endPos Types_Position
                endPos = multiLineToken.FP.Get_endPos(_env)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( endPos.LineNo == lineNo) &&
                    _env.SetStackVal( endPos.Column + 1 == tokenColumn) ).(bool)){
                    consecutive = true
                }
            }
        }
    }
    var newToken *Types_Token
    if tokenLineNo != nil && endColumn != nil{
        endColumn_232 := endColumn.(LnsInt)
        newToken = &NewAsyncTokenizer_MultiLineToken(_env, tokenKind, token, Types_Position_create(_env, lineNo, tokenColumn, self.streamName, self.overridePos), consecutive, NewLnsList2_[*Types_Token]([]*Types_Token{}), Types_Position_create(_env, self.lineNo, endColumn_232, self.streamName, self.overridePos)).Types_Token
    } else {
        newToken = NewTypes_Token(_env, tokenKind, token, Types_Position_create(_env, lineNo, tokenColumn, self.streamName, self.overridePos), consecutive, NewLnsList2_[*Types_Token]([]*Types_Token{}))
    }
    self.prevToken = newToken
    return newToken
}
// 431: decl @lune.@base.@AsyncTokenizer.Tokenizer.analyzeNumber
func (self *AsyncTokenizer_Tokenizer) analyzeNumber(_env *LnsEnv, token string,beginIndex LnsInt)(LnsInt, bool) {
    var nonNumIndex LnsInt
    
    {
        _nonNumIndex := AsyncTokenizer_convExp0_2109(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", beginIndex, nil)))
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
            _nonNumIndex = AsyncTokenizer_convExp0_2160(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
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
            _nonNumIndex = AsyncTokenizer_convExp0_2211(Lns_2DDD(_env.GetVM().String_find(token,"[^%da-fA-F]", nonNumIndex + 1, nil)))
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
                _nonNumIndex = AsyncTokenizer_convExp0_2291(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 2, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        } else { 
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncTokenizer_convExp0_2321(Lns_2DDD(_env.GetVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
                if Lns_IsNil( _nonNumIndex ) {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        }
    }
    return nonNumIndex - 1, intFlag
}
// 472: decl @lune.@base.@AsyncTokenizer.Tokenizer.readLine
func (self *AsyncTokenizer_Tokenizer) readLine(_env *LnsEnv) LnsAny {
    if self.lineNo >= self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    return self.lineList.GetAt(self.lineNo)
}
// 491: decl @lune.@base.@AsyncTokenizer.Tokenizer.addVal
func (self *AsyncTokenizer_Tokenizer) addVal(_env *LnsEnv, list *LnsList2_[*Types_Token],kind LnsInt,val string,column LnsInt) {
    if kind != Types_TokenKind__Symb{
        list.Insert(self.FP.createInfo(_env, kind, val, column, nil, nil))
        return 
    }
    var searchIndex LnsInt
    searchIndex = 1
    for  {
        var tokenIndex LnsInt
        var tokenEndIndex LnsInt
        
        {
            _tokenIndex, _tokenEndIndex := AsyncTokenizer_convExp0_2424(Lns_2DDD(_env.GetVM().String_find(val, "[%p%w]+", searchIndex, nil)))
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
                    _env.SetStackVal( Types_TokenKind__Real) ).(LnsInt), _env.GetVM().String_sub(token,subIndex, endIndex), columnIndex + subIndex, nil, nil)
                list.Insert(info)
                subIndex = endIndex + 1
            } else { 
                {
                    __exp := AsyncTokenizer_convExp0_2931(Lns_2DDD(_env.GetVM().String_find(token, "[^%w_]", subIndex, nil)))
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(LnsInt)
                        var index LnsInt
                        index = _exp
                        if index > subIndex{
                            var info *Types_Token
                            info = self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, index - 1), columnIndex + subIndex, nil, nil)
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
                            _env.SetStackVal( AsyncTokenizer_op2Set.Has(delimit)) ||
                            _env.SetStackVal( AsyncTokenizer_op1Set.Has(delimit)) ).(bool){
                            workKind = Types_TokenKind__Ope
                        }
                        if delimit == "..."{
                            workKind = Types_TokenKind__Symb
                        }
                        if delimit == "?"{
                            var nextChar string
                            nextChar = _env.GetVM().String_sub(token,index, subIndex)
                            list.Insert(self.FP.createInfo(_env, Types_TokenKind__Char, nextChar, columnIndex + subIndex, nil, nil))
                            subIndex = subIndex + 1
                        } else { 
                            list.Insert(self.FP.createInfo(_env, workKind, delimit, columnIndex + index, nil, nil))
                        }
                    } else {
                        if subIndex <= len(token){
                            list.Insert(self.FP.createInfo(_env, Types_TokenKind__Symb, _env.GetVM().String_sub(token,subIndex, nil), columnIndex + subIndex, nil, nil))
                        }
                        break
                    }
                }
            }
        }
    }
}
// 589: decl @lune.@base.@AsyncTokenizer.Tokenizer.parse
func (self *AsyncTokenizer_Tokenizer) Parse(_env *LnsEnv) LnsAny {
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
    var AsyncTokenizer_multiComment func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt)
    AsyncTokenizer_multiComment = func(_env *LnsEnv, comIndex LnsInt,termStr string)(string, LnsInt) {
        var searchIndex LnsInt
        searchIndex = comIndex
        var comment string
        comment = ""
        for  {
            {
                _, _termEndIndex := AsyncTokenizer_convExp0_3076(Lns_2DDD(_env.GetVM().String_find(rawLine, termStr, searchIndex, true)))
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
    var AsyncTokenizer_getChar func(_env *LnsEnv, index LnsInt) LnsInt
    AsyncTokenizer_getChar = func(_env *LnsEnv, index LnsInt) LnsInt {
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
            _index := AsyncTokenizer_convExp0_3219(Lns_2DDD(_env.GetVM().String_find(rawLine, pattern, searchIndex, nil)))
            if _index == nil{
                self.FP.addVal(_env, list, Types_TokenKind__Symb, _env.GetVM().String_sub(rawLine,startIndex, nil), startIndex)
                return list
            } else {
                index = _index.(LnsInt)
            }
        }
        var findChar LnsInt
        findChar = AsyncTokenizer_getChar(_env, index)
        var nextChar LnsInt
        nextChar = AsyncTokenizer_getChar(_env, index + 1)
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
                        _endIndex := AsyncTokenizer_convExp0_3364(Lns_2DDD(_env.GetVM().String_find(rawLine, workPattern, workIndex, nil)))
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
                    _env.SetStackVal( AsyncTokenizer_getChar(_env, index + 2) == 96) ).(bool))){
                    var lineNo LnsInt
                    lineNo = self.lineNo
                    var txt string
                    var nextIndex LnsInt
                    txt,nextIndex = AsyncTokenizer_multiComment(_env, index + 3, "```")
                    list.Insert(self.FP.createInfo(_env, Types_TokenKind__Str, "```" + txt, index, lineNo, nextIndex - 1))
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
                    if AsyncTokenizer_quotedCharSet.Has(quoted){
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
                        comment,nextIndex = AsyncTokenizer_multiComment(_env, index + 2, "*/")
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
// 349: decl @lune.@base.@AsyncTokenizer.Runner.run
func (self *AsyncTokenizer_Runner) Run(_env *LnsEnv) {
    {
        __exp := self.tokenizer
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncTokenizer_Tokenizer)
            _exp.FP.Run(_env)
        }
    }
}
// declaration Class -- AsyncItem
type AsyncTokenizer_AsyncItemMtd interface {
    ToMap() *LnsMap
}
type AsyncTokenizer_AsyncItem struct {
    List *LnsList2_[*Types_Token]
    FP AsyncTokenizer_AsyncItemMtd
}
func AsyncTokenizer_AsyncItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncTokenizer_AsyncItem).FP
}
func AsyncTokenizer_AsyncItem_toSlice(slice []LnsAny) []*AsyncTokenizer_AsyncItem {
    ret := make([]*AsyncTokenizer_AsyncItem, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncTokenizer_AsyncItemDownCast).ToAsyncTokenizer_AsyncItem()
    }
    return ret
}
type AsyncTokenizer_AsyncItemDownCast interface {
    ToAsyncTokenizer_AsyncItem() *AsyncTokenizer_AsyncItem
}
func AsyncTokenizer_AsyncItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncTokenizer_AsyncItemDownCast)
    if ok { return work.ToAsyncTokenizer_AsyncItem() }
    return nil
}
func (obj *AsyncTokenizer_AsyncItem) ToAsyncTokenizer_AsyncItem() *AsyncTokenizer_AsyncItem {
    return obj
}
func NewAsyncTokenizer_AsyncItem(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) *AsyncTokenizer_AsyncItem {
    obj := &AsyncTokenizer_AsyncItem{}
    obj.FP = obj
    obj.InitAsyncTokenizer_AsyncItem(_env, arg1)
    return obj
}
func (self *AsyncTokenizer_AsyncItem) InitAsyncTokenizer_AsyncItem(_env *LnsEnv, arg1 *LnsList2_[*Types_Token]) {
    self.List = arg1
}
func (self *AsyncTokenizer_AsyncItem) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["list"] = Lns_ToCollection( self.List )
    return obj
}
func (self *AsyncTokenizer_AsyncItem) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func AsyncTokenizer_AsyncItem__fromMap(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncTokenizer_AsyncItem_FromMap( arg1, paramList )
}
func AsyncTokenizer_AsyncItem__fromStem(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncTokenizer_AsyncItem_FromMap( arg1, paramList )
}
func AsyncTokenizer_AsyncItem_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := AsyncTokenizer_AsyncItem_FromMapSub(obj,false, paramList);
    return conv,mess
}
func AsyncTokenizer_AsyncItem_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &AsyncTokenizer_AsyncItem{}
    newObj.FP = newObj
    return AsyncTokenizer_AsyncItem_FromMapMain( newObj, objMap, paramList )
}
func AsyncTokenizer_AsyncItem_FromMapMain( newObj *AsyncTokenizer_AsyncItem, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToList2Sub[*Types_Token]( objMap.Items["list"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Types_Token_FromMapSub, false,nil}}); !ok {
       return false,nil,"list:" + mess.(string)
    } else {
       newObj.List = conv.(*LnsList2_[*Types_Token])
    }
    return true, newObj, nil
}

// declaration Class -- MultiLineToken
type AsyncTokenizer_MultiLineTokenMtd interface {
    ToMap() *LnsMap
    GetExcludedDelimitTxt(_env *LnsEnv) string
    GetLineCount(_env *LnsEnv) LnsInt
    Get_commentList(_env *LnsEnv) *LnsList2_[*Types_Token]
    Get_endPos(_env *LnsEnv) Types_Position
}
type AsyncTokenizer_MultiLineToken struct {
    Types_Token
    endPos Types_Position
    FP AsyncTokenizer_MultiLineTokenMtd
}
func AsyncTokenizer_MultiLineToken2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncTokenizer_MultiLineToken).FP
}
func AsyncTokenizer_MultiLineToken_toSlice(slice []LnsAny) []*AsyncTokenizer_MultiLineToken {
    ret := make([]*AsyncTokenizer_MultiLineToken, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncTokenizer_MultiLineTokenDownCast).ToAsyncTokenizer_MultiLineToken()
    }
    return ret
}
type AsyncTokenizer_MultiLineTokenDownCast interface {
    ToAsyncTokenizer_MultiLineToken() *AsyncTokenizer_MultiLineToken
}
func AsyncTokenizer_MultiLineTokenDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncTokenizer_MultiLineTokenDownCast)
    if ok { return work.ToAsyncTokenizer_MultiLineToken() }
    return nil
}
func (obj *AsyncTokenizer_MultiLineToken) ToAsyncTokenizer_MultiLineToken() *AsyncTokenizer_MultiLineToken {
    return obj
}
func NewAsyncTokenizer_MultiLineToken(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 Types_Position, arg4 bool, arg5 LnsAny, arg6 Types_Position) *AsyncTokenizer_MultiLineToken {
    obj := &AsyncTokenizer_MultiLineToken{}
    obj.FP = obj
    obj.Types_Token.FP = obj
    obj.InitAsyncTokenizer_MultiLineToken(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *AsyncTokenizer_MultiLineToken) InitAsyncTokenizer_MultiLineToken(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 Types_Position, arg4 bool, arg5 LnsAny, arg6 Types_Position) {
    self.Types_Token.InitTypes_Token( _env, arg1,arg2,arg3,arg4,arg5)
    self.endPos = arg6
}
func (self *AsyncTokenizer_MultiLineToken) Get_endPos(_env *LnsEnv) Types_Position{ return self.endPos }
func (self *AsyncTokenizer_MultiLineToken) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["endPos"] = Lns_ToCollection( self.endPos )
    return self.Types_Token.ToMapSetup( obj )
}
func (self *AsyncTokenizer_MultiLineToken) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func AsyncTokenizer_MultiLineToken__fromMap_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncTokenizer_MultiLineToken_FromMap( arg1, paramList )
}
func AsyncTokenizer_MultiLineToken__fromStem_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncTokenizer_MultiLineToken_FromMap( arg1, paramList )
}
func AsyncTokenizer_MultiLineToken_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := AsyncTokenizer_MultiLineToken_FromMapSub(obj,false, paramList);
    return conv,mess
}
func AsyncTokenizer_MultiLineToken_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &AsyncTokenizer_MultiLineToken{}
    newObj.FP = newObj
    newObj.Types_Token.FP = newObj
    return AsyncTokenizer_MultiLineToken_FromMapMain( newObj, objMap, paramList )
}
func AsyncTokenizer_MultiLineToken_FromMapMain( newObj *AsyncTokenizer_MultiLineToken, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Types_Token_FromMapMain( &newObj.Types_Token, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Types_Position_FromMapSub( objMap.Items["endPos"], false, nil); !ok {
       return false,nil,"endPos:" + mess.(string)
    } else {
       newObj.endPos = conv.(Types_Position)
    }
    return true, newObj, nil
}

// declaration Class -- Tokenizer
type AsyncTokenizer_TokenizerMtd interface {
    Access(_env *LnsEnv) LnsAny
    addVal(_env *LnsEnv, arg1 *LnsList2_[*Types_Token], arg2 LnsInt, arg3 string, arg4 LnsInt)
    analyzeNumber(_env *LnsEnv, arg1 string, arg2 LnsInt)(LnsInt, bool)
    createInfo(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsAny, arg5 LnsAny) *Types_Token
    GetNext(_env *LnsEnv) LnsAny
    Get_streamName(_env *LnsEnv) string
    Parse(_env *LnsEnv) LnsAny
    readLine(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
    Setup(_env *LnsEnv)
    Start(_env *LnsEnv)
    Stop(_env *LnsEnv)
}
type AsyncTokenizer_Tokenizer struct {
    Async_Pipe
    streamName string
    lineNo LnsInt
    prevToken *Types_Token
    keywordSet *LnsSet2_[string]
    typeSet *LnsSet2_[string]
    multiCharDelimitMap *LnsMap2_[string,*LnsList]
    luaMode bool
    lineList *LnsList2_[string]
    firstLine bool
    overridePos LnsAny
    stream Lns_iStream
    FP AsyncTokenizer_TokenizerMtd
}
func AsyncTokenizer_Tokenizer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncTokenizer_Tokenizer).FP
}
func AsyncTokenizer_Tokenizer_toSlice(slice []LnsAny) []*AsyncTokenizer_Tokenizer {
    ret := make([]*AsyncTokenizer_Tokenizer, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncTokenizer_TokenizerDownCast).ToAsyncTokenizer_Tokenizer()
    }
    return ret
}
type AsyncTokenizer_TokenizerDownCast interface {
    ToAsyncTokenizer_Tokenizer() *AsyncTokenizer_Tokenizer
}
func AsyncTokenizer_TokenizerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncTokenizer_TokenizerDownCast)
    if ok { return work.ToAsyncTokenizer_Tokenizer() }
    return nil
}
func (obj *AsyncTokenizer_Tokenizer) ToAsyncTokenizer_Tokenizer() *AsyncTokenizer_Tokenizer {
    return obj
}
func NewAsyncTokenizer_Tokenizer(_env *LnsEnv, arg1 string, arg2 Lns_iStream, arg3 bool, arg4 LnsAny, arg5 LnsAny) *AsyncTokenizer_Tokenizer {
    obj := &AsyncTokenizer_Tokenizer{}
    obj.FP = obj
    obj.Async_Pipe.FP = obj
    obj.InitAsyncTokenizer_Tokenizer(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *AsyncTokenizer_Tokenizer) Get_streamName(_env *LnsEnv) string{ return self.streamName }
// 238: DeclConstr
func (self *AsyncTokenizer_Tokenizer) InitAsyncTokenizer_Tokenizer(_env *LnsEnv, streamName string,stream Lns_iStream,luaMode bool,overridePos LnsAny,pipeSize LnsAny) {
    self.InitAsync_Pipe(_env, LnsAny(NewLnspipe( Lns_unwrapDefault( pipeSize, AsyncTokenizer_defaultPipeSize).(LnsInt))))
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
    var multiCharDelimitMap *LnsMap2_[string,*LnsList]
    keywordSet,typeSet,_,multiCharDelimitMap = AsyncTokenizer_createReserveInfo_1_(_env, luaMode)
    self.keywordSet = keywordSet
    self.typeSet = typeSet
    self.multiCharDelimitMap = multiCharDelimitMap
}



// declaration Class -- Runner
type AsyncTokenizer_RunnerMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Get_errMess(_env *LnsEnv) string
    Get_tokenizer(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
}
type AsyncTokenizer_Runner struct {
    _syncFlag *Lns_syncFlag
    tokenizer LnsAny
    errMess string
    FP AsyncTokenizer_RunnerMtd
}
func (self *AsyncTokenizer_Runner) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
func AsyncTokenizer_Runner2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*AsyncTokenizer_Runner).FP
}
func AsyncTokenizer_Runner_toSlice(slice []LnsAny) []*AsyncTokenizer_Runner {
    ret := make([]*AsyncTokenizer_Runner, len(slice))
    for index, val := range slice {
        ret[index] = val.(AsyncTokenizer_RunnerDownCast).ToAsyncTokenizer_Runner()
    }
    return ret
}
type AsyncTokenizer_RunnerDownCast interface {
    ToAsyncTokenizer_Runner() *AsyncTokenizer_Runner
}
func AsyncTokenizer_RunnerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(AsyncTokenizer_RunnerDownCast)
    if ok { return work.ToAsyncTokenizer_Runner() }
    return nil
}
func (obj *AsyncTokenizer_Runner) ToAsyncTokenizer_Runner() *AsyncTokenizer_Runner {
    return obj
}
func NewAsyncTokenizer_Runner(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) *AsyncTokenizer_Runner {
    obj := &AsyncTokenizer_Runner{}
    obj.FP = obj
    obj.InitAsyncTokenizer_Runner(_env, arg1, arg2, arg3)
    return obj
}
func (self *AsyncTokenizer_Runner) Get_tokenizer(_env *LnsEnv) LnsAny{ return self.tokenizer }
func (self *AsyncTokenizer_Runner) Get_errMess(_env *LnsEnv) string{ return self.errMess }
// 334: DeclConstr
func (self *AsyncTokenizer_Runner) InitAsyncTokenizer_Runner(_env *LnsEnv, tokenizerSrc LnsAny,stdinFile LnsAny,overridePos LnsAny) {
    self._syncFlag = &Lns_syncFlag{}
    self.tokenizer, self.errMess = AsyncTokenizer_Tokenizer_create_3_(_env, tokenizerSrc, stdinFile, overridePos)
    {
        __exp := self.tokenizer
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*AsyncTokenizer_Tokenizer)
            _exp.FP.Start(_env)
            if Lns_op_not(LnsRun(_env, self.FP, 2, _env.GetVM().String_format("tokenizer - %s", Lns_2DDD(_exp.FP.Get_streamName(_env))))){
                _exp.FP.Stop(_env)
            }
        }
    }
}


func Lns_AsyncTokenizer_init(_env *LnsEnv) {
    if init_AsyncTokenizer { return }
    init_AsyncTokenizer = true
    AsyncTokenizer__mod__ = "@lune.@base.@AsyncTokenizer"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Types_init(_env)
    Lns_Async_init(_env)
    AsyncTokenizer_luaKeywordSet = NewLnsSet2_[string](Lns_2DDDGen[string]("if", "else", "elseif", "while", "for", "in", "return", "break", "nil", "true", "false", "{", "}", "do", "require", "function", "then", "end", "repeat", "until", "goto", "local"))
    AsyncTokenizer_quotedCharSet = NewLnsSet2_[string](Lns_2DDDGen[string]("a", "b", "f", "n", "r", "t", "v", "\\", "\"", "'"))
    AsyncTokenizer_op2Set = NewLnsSet2_[string](Lns_2DDDGen[string]("+", "-", "*", "/", "^", "%", "&", "~", "|", "|>>", "|<<", "..", "<", "<=", ">", ">=", "==", "~=", "and", "or", "@", "@@", "@@@", "="))
    AsyncTokenizer_op1Set = NewLnsSet2_[string](Lns_2DDDGen[string]("-", "not", "#", "~", "`{", ",,", ",,,", ",,,,"))
    AsyncTokenizer_defaultPipeSize = 100
}
func init() {
    init_AsyncTokenizer = false
}
