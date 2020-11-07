// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_AsyncParser bool
var AsyncParser__mod__ string
var AsyncParser_luaKeywordSet *LnsSet
var AsyncParser_quotedCharSet *LnsSet
var AsyncParser_op2Set *LnsSet
var AsyncParser_op1Set *LnsSet
// for 395
func AsyncParser_convExp2206(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 471
func AsyncParser_convExp2294(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 299
func AsyncParser_convExp1374(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 307
func AsyncParser_convExp1424(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 315
func AsyncParser_convExp1474(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 326
func AsyncParser_convExp1552(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 331
func AsyncParser_convExp1581(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 369
func AsyncParser_convExp1707(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 502
func AsyncParser_convExp2436(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 525
func AsyncParser_convExp2578(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 55: decl @lune.@base.@AsyncParser.isLuaKeyword
func AsyncParser_isLuaKeyword(txt string) bool {
    return AsyncParser_luaKeywordSet.Has(txt)
}

// 60: decl @lune.@base.@AsyncParser.createReserveInfo
func AsyncParser_createReserveInfo_1032_(luaMode LnsAny)(*LnsSet, *LnsSet, *LnsSet, *LnsMap) {
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

// 198: decl @lune.@base.@AsyncParser.isOp2
func AsyncParser_isOp2(ope string) bool {
    return AsyncParser_op2Set.Has(ope)
}

// 202: decl @lune.@base.@AsyncParser.isOp1
func AsyncParser_isOp1(ope string) bool {
    return AsyncParser_op1Set.Has(ope)
}

// 298: decl @lune.@base.@AsyncParser.analyzeNumber
func AsyncParser_analyzeNumber_1249_(token string,beginIndex LnsInt)(LnsInt, bool) {
    var nonNumIndex LnsInt
    
    {
        _nonNumIndex := AsyncParser_convExp1374(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", beginIndex, nil)))
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
            _nonNumIndex = AsyncParser_convExp1424(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
            if _nonNumIndex == nil {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
        
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( nonNumChar == 88) ||
        Lns_GetEnv().SetStackVal( nonNumChar == 120) ).(bool){
        {
            var _nonNumIndex LnsAny
            _nonNumIndex = AsyncParser_convExp1474(Lns_2DDD(Lns_getVM().String_find(token,"[^%da-fA-F]", nonNumIndex + 1, nil)))
            if _nonNumIndex == nil {
                return len(token), intFlag
            }
            nonNumIndex = _nonNumIndex.(LnsInt)
        }
        nonNumChar = LnsInt(token[nonNumIndex-1])
        
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( nonNumChar == 69) ||
        Lns_GetEnv().SetStackVal( nonNumChar == 101) ).(bool){
        intFlag = false
        
        var nextChar LnsInt
        nextChar = LnsInt(token[nonNumIndex + 1-1])
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nextChar == 45) ||
            Lns_GetEnv().SetStackVal( nextChar == 43) ).(bool){
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncParser_convExp1552(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 2, nil)))
                if _nonNumIndex == nil {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        } else { 
            {
                var _nonNumIndex LnsAny
                _nonNumIndex = AsyncParser_convExp1581(Lns_2DDD(Lns_getVM().String_find(token,"[^%d]", nonNumIndex + 1, nil)))
                if _nonNumIndex == nil {
                    return len(token), intFlag
                }
                nonNumIndex = _nonNumIndex.(LnsInt)
            }
        }
    }
    return nonNumIndex - 1, intFlag
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
func NewAsyncParser_AsyncItem(arg1 *LnsList) *AsyncParser_AsyncItem {
    obj := &AsyncParser_AsyncItem{}
    obj.FP = obj
    obj.InitAsyncParser_AsyncItem(arg1)
    return obj
}
func (self *AsyncParser_AsyncItem) InitAsyncParser_AsyncItem(arg1 *LnsList) {
    self.List = arg1
}
func (self *AsyncParser_AsyncItem) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["list"] = Lns_ToCollection( self.List )
    return obj
}
func (self *AsyncParser_AsyncItem) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func AsyncParser_AsyncItem__fromMap(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return AsyncParser_AsyncItem_FromMap( arg1, paramList )
}
func AsyncParser_AsyncItem__fromStem(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
func AsyncParser_AsyncItem__createPipe(arg1 LnsInt) LnsAny{
   return NewLnspipe( arg1 )
}

// declaration Class -- Parser
type AsyncParser_ParserMtd interface {
    Access() LnsAny
    addVal(arg1 *LnsList, arg2 LnsInt, arg3 string, arg4 LnsInt)
    createInfo(arg1 LnsInt, arg2 string, arg3 LnsInt) *Types_Token
    GetNext() LnsAny
    Loop()
    Parse() LnsAny
    readLine() LnsAny
    Start()
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
func NewAsyncParser_Parser(arg1 Lns_iStream, arg2 string, arg3 LnsAny) *AsyncParser_Parser {
    obj := &AsyncParser_Parser{}
    obj.FP = obj
    obj.Async_Pipe.FP = obj
    obj.LnsThread.FP = obj
    obj.InitAsyncParser_Parser(arg1, arg2, arg3)
    return obj
}
// 220: DeclConstr
func (self *AsyncParser_Parser) InitAsyncParser_Parser(stream Lns_iStream,name string,luaMode LnsAny) {
    self.InitAsync_Pipe(AsyncParser_AsyncItem__createPipe(100))
    self.streamName = name
    
    self.lineNo = 0
    
    self.prevToken = Types_noneToken
    
    self.luaMode = Lns_unwrapDefault( luaMode, false).(bool)
    
    var keywordSet *LnsSet
    var typeSet *LnsSet
    var multiCharDelimitMap *LnsMap
    keywordSet,typeSet,_,multiCharDelimitMap = AsyncParser_createReserveInfo_1032_(luaMode)
    self.keywordSet = keywordSet
    
    self.typeSet = typeSet
    
    self.multiCharDelimitMap = multiCharDelimitMap
    
    var lineList *LnsList
    lineList = NewLnsList([]LnsAny{})
    for  {
        var line string
        
        {
            _line := stream.Read("*l")
            if _line == nil{
                break
            } else {
                line = _line.(string)
            }
        }
        lineList.Insert(line)
    }
    self.lineList = lineList
    
}


// 250: decl @lune.@base.@AsyncParser.Parser.access
func (self *AsyncParser_Parser) Access() LnsAny {
    var tokenList *LnsList
    
    {
        _tokenList := self.FP.Parse()
        if _tokenList == nil{
            return nil
        } else {
            tokenList = _tokenList.(*LnsList)
        }
    }
    return NewAsync_PipeItem(AsyncParser_AsyncItem2Stem(NewAsyncParser_AsyncItem(tokenList)))
}

// 261: decl @lune.@base.@AsyncParser.Parser.createInfo
func (self *AsyncParser_Parser) createInfo(tokenKind LnsInt,token string,tokenColumn LnsInt) *Types_Token {
    if tokenKind == Types_TokenKind__Symb{
        if self.keywordSet.Has(token){
            tokenKind = Types_TokenKind__Kywd
            
        } else if self.typeSet.Has(token){
            tokenKind = Types_TokenKind__Type
            
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( AsyncParser_op2Set.Has(token)) ||
            Lns_GetEnv().SetStackVal( AsyncParser_op1Set.Has(token)) ).(bool){
            tokenKind = Types_TokenKind__Ope
            
        }
    }
    var consecutive bool
    consecutive = false
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.prevToken.Pos.LineNo == self.lineNo) &&
        Lns_GetEnv().SetStackVal( self.prevToken.Pos.Column + len(self.prevToken.Txt) == tokenColumn) ).(bool)){
        consecutive = true
        
    }
    var newToken *Types_Token
    newToken = NewTypes_Token(tokenKind, token, NewTypes_Position(self.lineNo, tokenColumn, self.streamName), consecutive, NewLnsList([]LnsAny{}))
    self.prevToken = newToken
    
    return newToken
}

// 339: decl @lune.@base.@AsyncParser.Parser.readLine
func (self *AsyncParser_Parser) readLine() LnsAny {
    if self.lineNo >= self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    
    return self.lineList.GetAt(self.lineNo).(string)
}

// 358: decl @lune.@base.@AsyncParser.Parser.addVal
func (self *AsyncParser_Parser) addVal(list *LnsList,kind LnsInt,val string,column LnsInt) {
    if kind != Types_TokenKind__Symb{
        list.Insert(Types_Token2Stem(self.FP.createInfo(kind, val, column)))
        return 
    }
    var searchIndex LnsInt
    searchIndex = 1
    for  {
        var tokenIndex LnsInt
        var tokenEndIndex LnsInt
        
        {
            _tokenIndex, _tokenEndIndex := AsyncParser_convExp1707(Lns_2DDD(Lns_getVM().String_find(val, "[%p%w]+", searchIndex, nil)))
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
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(token,"^[%d]", subIndex, nil))) ||
                Lns_GetEnv().SetStackVal( LnsInt(token[subIndex-1]) == 45) &&
                Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(token,"^[%d]", subIndex + 1, nil))) )){
                var checkIndex LnsInt
                checkIndex = subIndex
                if LnsInt(token[checkIndex-1]) == 45{
                    checkIndex = checkIndex + 1
                    
                }
                var endIndex LnsInt
                var intFlag bool
                endIndex,intFlag = AsyncParser_analyzeNumber_1249_(token, checkIndex)
                var info *Types_Token
                info = self.FP.createInfo(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( intFlag) &&
                    Lns_GetEnv().SetStackVal( Types_TokenKind__Int) ||
                    Lns_GetEnv().SetStackVal( Types_TokenKind__Real) ).(LnsInt), Lns_getVM().String_sub(token,subIndex, endIndex), columnIndex + subIndex)
                list.Insert(Types_Token2Stem(info))
                subIndex = endIndex + 1
                
            } else { 
                {
                    __exp := AsyncParser_convExp2206(Lns_2DDD(Lns_getVM().String_find(token, "[^%w_]", subIndex, nil)))
                    if __exp != nil {
                        _exp := __exp.(LnsInt)
                        var index LnsInt
                        index = _exp
                        if index > subIndex{
                            var info *Types_Token
                            info = self.FP.createInfo(Types_TokenKind__Symb, Lns_getVM().String_sub(token,subIndex, index - 1), columnIndex + subIndex)
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
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( AsyncParser_op2Set.Has(delimit)) ||
                            Lns_GetEnv().SetStackVal( AsyncParser_op1Set.Has(delimit)) ).(bool){
                            workKind = Types_TokenKind__Ope
                            
                        }
                        if delimit == "..."{
                            workKind = Types_TokenKind__Symb
                            
                        }
                        if delimit == "?"{
                            var nextChar string
                            nextChar = Lns_getVM().String_sub(token,index, subIndex)
                            list.Insert(Types_Token2Stem(self.FP.createInfo(Types_TokenKind__Char, nextChar, columnIndex + subIndex)))
                            subIndex = subIndex + 1
                            
                        } else { 
                            list.Insert(Types_Token2Stem(self.FP.createInfo(workKind, delimit, columnIndex + index)))
                        }
                    } else {
                        if subIndex <= len(token){
                            list.Insert(Types_Token2Stem(self.FP.createInfo(Types_TokenKind__Symb, Lns_getVM().String_sub(token,subIndex, nil), columnIndex + subIndex)))
                        }
                        break
                    }
                }
            }
        }
    }
}

// 456: decl @lune.@base.@AsyncParser.Parser.parse
func (self *AsyncParser_Parser) Parse() LnsAny {
    var rawLine string
    
    {
        _rawLine := self.FP.readLine()
        if _rawLine == nil{
            return nil
        } else {
            rawLine = _rawLine.(string)
        }
    }
    var multiComment func(comIndex LnsInt,termStr string)(string, LnsInt)
    multiComment = func(comIndex LnsInt,termStr string)(string, LnsInt) {
        var searchIndex LnsInt
        searchIndex = comIndex
        var comment string
        comment = ""
        for  {
            {
                _, _termEndIndex := AsyncParser_convExp2294(Lns_2DDD(Lns_getVM().String_find(rawLine, termStr, searchIndex, true)))
                if _termEndIndex != nil {
                    termEndIndex := _termEndIndex.(LnsInt)
                    comment = comment + Lns_getVM().String_sub(rawLine,searchIndex, termEndIndex)
                    
                    return comment, termEndIndex + 1
                }
            }
            comment = comment + Lns_getVM().String_sub(rawLine,searchIndex, nil) + "\n"
            
            searchIndex = 1
            
            rawLine = Lns_unwrap( self.FP.readLine()).(string)
            
        }
    // insert a dummy
        return "",0
    }
    var startIndex LnsInt
    startIndex = 1
    var searchIndex LnsInt
    searchIndex = startIndex
    var getChar func(index LnsInt) LnsInt
    getChar = func(index LnsInt) LnsInt {
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
            _index := AsyncParser_convExp2436(Lns_2DDD(Lns_getVM().String_find(rawLine, pattern, searchIndex, nil)))
            if _index == nil{
                self.FP.addVal(list, Types_TokenKind__Symb, Lns_getVM().String_sub(rawLine,startIndex, nil), startIndex)
                return list
            } else {
                index = _index.(LnsInt)
            }
        }
        var findChar LnsInt
        findChar = getChar(index)
        var nextChar LnsInt
        nextChar = getChar(index + 1)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( findChar == 45) &&
            Lns_GetEnv().SetStackVal( nextChar != 45) ).(bool)){
            searchIndex = index + 1
            
            syncIndexFlag = false
            
        } else { 
            if startIndex < index{
                self.FP.addVal(list, Types_TokenKind__Symb, Lns_getVM().String_sub(rawLine,startIndex, index - 1), startIndex)
            }
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( findChar == 39) ||
                Lns_GetEnv().SetStackVal( findChar == 34) ).(bool){
                var workIndex LnsInt
                workIndex = index + 1
                var workPattern string
                workPattern = "[\"'\\]"
                for  {
                    var endIndex LnsInt
                    
                    {
                        _endIndex := AsyncParser_convExp2578(Lns_2DDD(Lns_getVM().String_find(rawLine, workPattern, workIndex, nil)))
                        if _endIndex == nil{
                            Util_err(Lns_getVM().String_format("%s:%d:%d: error: illegal string -- %s", []LnsAny{self.streamName, self.lineNo, index, rawLine}))
                        } else {
                            endIndex = _endIndex.(LnsInt)
                        }
                    }
                    var workChar LnsInt
                    workChar = LnsInt(rawLine[endIndex-1])
                    if workChar == findChar{
                        self.FP.addVal(list, Types_TokenKind__Str, Lns_getVM().String_sub(rawLine,index, endIndex), index)
                        searchIndex = endIndex + 1
                        
                        break
                    } else if workChar == 92{
                        workIndex = endIndex + 2
                        
                    } else { 
                        workIndex = endIndex + 1
                        
                    }
                }
            } else if findChar == 96{
                if Lns_isCondTrue( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( nextChar == findChar) &&
                    Lns_GetEnv().SetStackVal( LnsInt(rawLine[index + 2-1]) == 96) ).(bool))){
                    var txt string
                    var nextIndex LnsInt
                    txt,nextIndex = multiComment(index + 3, "```")
                    self.FP.addVal(list, Types_TokenKind__Str, "```" + txt, index)
                    searchIndex = nextIndex
                    
                } else { 
                    self.FP.addVal(list, Types_TokenKind__Ope, "`", index)
                    searchIndex = index + 1
                    
                }
            } else if findChar == 63{
                var codeChar string
                codeChar = Lns_getVM().String_sub(rawLine,index + 1, index + 1)
                if nextChar == 92{
                    var quoted string
                    quoted = Lns_getVM().String_sub(rawLine,index + 2, index + 2)
                    if AsyncParser_quotedCharSet.Has(quoted){
                        codeChar = Lns_getVM().String_sub(rawLine,index + 1, index + 2)
                        
                    } else { 
                        codeChar = quoted
                        
                    }
                    searchIndex = index + 3
                    
                } else { 
                    searchIndex = index + 2
                    
                }
                self.FP.addVal(list, Types_TokenKind__Char, codeChar, index)
            } else { 
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( self.luaMode) &&
                    Lns_GetEnv().SetStackVal( findChar == 45) &&
                    Lns_GetEnv().SetStackVal( nextChar == 45) ).(bool)){
                    self.FP.addVal(list, Types_TokenKind__Cmnt, Lns_getVM().String_sub(rawLine,index, nil), index)
                    searchIndex = len(rawLine) + 1
                    
                } else if findChar == 47{
                    if nextChar == 42{
                        var comment string
                        var nextIndex LnsInt
                        comment,nextIndex = multiComment(index + 2, "*/")
                        self.FP.addVal(list, Types_TokenKind__Cmnt, "/*" + comment, index)
                        searchIndex = nextIndex
                        
                    } else if nextChar == 47{
                        self.FP.addVal(list, Types_TokenKind__Cmnt, Lns_getVM().String_sub(rawLine,index, nil), index)
                        searchIndex = len(rawLine) + 1
                        
                    } else { 
                        self.FP.addVal(list, Types_TokenKind__Ope, "/", index)
                        searchIndex = index + 1
                        
                    }
                } else { 
                    Util_err(Lns_getVM().String_format("%s:%d:%d: error: illegal syntax -- %s", []LnsAny{self.streamName, self.lineNo, index, rawLine}))
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


func Lns_AsyncParser_init() {
    if init_AsyncParser { return }
    init_AsyncParser = true
    AsyncParser__mod__ = "@lune.@base.@AsyncParser"
    Lns_InitMod()
    Lns_Util_init()
    Lns_Types_init()
    Lns_Async_init()
    AsyncParser_luaKeywordSet = NewLnsSet([]LnsAny{"if", "else", "elseif", "while", "for", "in", "return", "break", "nil", "true", "false", "{", "}", "do", "require", "function", "then", "end", "repeat", "until", "goto", "local"})
    AsyncParser_quotedCharSet = NewLnsSet([]LnsAny{"a", "b", "f", "n", "r", "t", "v", "\\", "\"", "'"})
    AsyncParser_op2Set = NewLnsSet([]LnsAny{"+", "-", "*", "/", "^", "%", "&", "~", "|", "|>>", "|<<", "..", "<", "<=", ">", ">=", "==", "~=", "and", "or", "@", "@@", "@@@", "="})
    AsyncParser_op1Set = NewLnsSet([]LnsAny{"-", "not", "#", "~", "*", "`", ",,", ",,,", ",,,,"})
}
func init() {
    init_AsyncParser = false
}
