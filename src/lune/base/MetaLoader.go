// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_MetaLoader bool
var MetaLoader__mod__ string
var MetaLoader_quotedChar2Code *LnsMap
// for 378
func MetaLoader_convExp2307(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 307: decl @lune.@base.@MetaLoader.loadFromStream
func MetaLoader_loadFromStream(path string,stream Lns_iStream)(LnsAny, LnsAny) {
    var streamParser *Parser_StreamParser
    streamParser = NewParser_StreamParser(stream, path, true)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(&streamParser.Parser_Parser)
    var loader *MetaLoader_Loader
    loader = NewMetaLoader_Loader(parser.FP)
    var item LnsAny
    var err LnsAny
    item,err = loader.FP.Process()
    if Lns_isCondTrue( err){
        return nil, err
    }
    if item != nil{
        item_553 := item
        var _map *LnsMap
        _map = item_553.(*LnsMap)
        var formatVersion string
        {
            _workVal := _map.Items["__formatVersion"]
            if _workVal != nil {
                workVal := _workVal
                formatVersion = workVal.(string)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"formatVersion"})
            }
        }
        
        var enableTest bool
        {
            _workVal := _map.Items["__enableTest"]
            if _workVal != nil {
                workVal := _workVal
                enableTest = workVal.(bool)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"enableTest"})
            }
        }
        
        var buildId string
        {
            _workVal := _map.Items["__buildId"]
            if _workVal != nil {
                workVal := _workVal
                buildId = workVal.(string)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"buildId"})
            }
        }
        
        var typeId2ClassInfoMap *LnsMap
        {
            _workVal := _map.Items["__typeId2ClassInfoMap"]
            if _workVal != nil {
                workVal := _workVal
                typeId2ClassInfoMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"typeId2ClassInfoMap"})
            }
        }
        
        var typeInfoList *LnsMap
        {
            _workVal := _map.Items["__typeInfoList"]
            if _workVal != nil {
                workVal := _workVal
                typeInfoList = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"typeInfoList"})
            }
        }
        
        var varName2InfoMap *LnsMap
        {
            _workVal := _map.Items["__varName2InfoMap"]
            if _workVal != nil {
                workVal := _workVal
                varName2InfoMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"varName2InfoMap"})
            }
        }
        
        var moduleTypeId LnsInt
        {
            _workVal := _map.Items["__moduleTypeId"]
            if _workVal != nil {
                workVal := _workVal
                moduleTypeId = Lns_forceCastInt(workVal)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"moduleTypeId"})
            }
        }
        
        var moduleSymbolKind LnsInt
        {
            _workVal := _map.Items["__moduleSymbolKind"]
            if _workVal != nil {
                workVal := _workVal
                moduleSymbolKind = Lns_forceCastInt(workVal)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"moduleSymbolKind"})
            }
        }
        
        var moduleMutable bool
        {
            _workVal := _map.Items["__moduleMutable"]
            if _workVal != nil {
                workVal := _workVal
                moduleMutable = workVal.(bool)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"moduleMutable"})
            }
        }
        
        var dependModuleMap *LnsMap
        {
            _workVal := _map.Items["__dependModuleMap"]
            if _workVal != nil {
                workVal := _workVal
                dependModuleMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"dependModuleMap"})
            }
        }
        
        var dependIdMap *LnsMap
        {
            _workVal := _map.Items["__dependIdMap"]
            if _workVal != nil {
                workVal := _workVal
                dependIdMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"dependIdMap"})
            }
        }
        
        var macroName2InfoMap *LnsMap
        {
            _workVal := _map.Items["__macroName2InfoMap"]
            if _workVal != nil {
                workVal := _workVal
                macroName2InfoMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"macroName2InfoMap"})
            }
        }
        
        var subModuleMap *LnsMap
        {
            _workVal := _map.Items["__subModuleMap"]
            if _workVal != nil {
                workVal := _workVal
                subModuleMap = workVal.(*LnsMap)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"subModuleMap"})
            }
        }
        
        var hasTest bool
        {
            _workVal := _map.Items["__hasTest"]
            if _workVal != nil {
                workVal := _workVal
                hasTest = workVal.(bool)
                
            } else {
                return nil, Lns_getVM().String_format("error -- %s", []LnsAny{"hasTest"})
            }
        }
        
        return NewMeta__MetaInfo(formatVersion, enableTest, buildId, typeId2ClassInfoMap, typeInfoList, varName2InfoMap, moduleTypeId, moduleSymbolKind, moduleMutable, dependModuleMap, dependIdMap, macroName2InfoMap, hasTest, subModuleMap), nil
    }
    return nil, "error"
}

// 373: decl @lune.@base.@MetaLoader.loadFromCode
func MetaLoader_loadFromCode(path string,code string)(LnsAny, LnsAny) {
    return MetaLoader_loadFromStream(path, NewParser_TxtStream(code).FP)
}

// 377: decl @lune.@base.@MetaLoader.load
func MetaLoader_load(path string)(LnsAny, LnsAny) {
    var stream Lns_luaStream
    
    {
        _stream := MetaLoader_convExp2307(Lns_2DDD(Lns_io_open(path, "r")))
        if _stream == nil{
            return nil, Lns_getVM().String_format("open error -- %s", []LnsAny{path})
        } else {
            stream = _stream.(Lns_luaStream)
        }
    }
    return MetaLoader_loadFromStream(path, stream)
}

// declaration Class -- SymbolInfo
type MetaLoader_SymbolInfoMtd interface {
    Get_name() string
    Get_val() LnsAny
}
type MetaLoader_SymbolInfo struct {
    name string
    val LnsAny
    FP MetaLoader_SymbolInfoMtd
}
func MetaLoader_SymbolInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*MetaLoader_SymbolInfo).FP
}
type MetaLoader_SymbolInfoDownCast interface {
    ToMetaLoader_SymbolInfo() *MetaLoader_SymbolInfo
}
func MetaLoader_SymbolInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(MetaLoader_SymbolInfoDownCast)
    if ok { return work.ToMetaLoader_SymbolInfo() }
    return nil
}
func (obj *MetaLoader_SymbolInfo) ToMetaLoader_SymbolInfo() *MetaLoader_SymbolInfo {
    return obj
}
func NewMetaLoader_SymbolInfo(arg1 string, arg2 LnsAny) *MetaLoader_SymbolInfo {
    obj := &MetaLoader_SymbolInfo{}
    obj.FP = obj
    obj.InitMetaLoader_SymbolInfo(arg1, arg2)
    return obj
}
func (self *MetaLoader_SymbolInfo) InitMetaLoader_SymbolInfo(arg1 string, arg2 LnsAny) {
    self.name = arg1
    self.val = arg2
}
func (self *MetaLoader_SymbolInfo) Get_name() string{ return self.name }
func (self *MetaLoader_SymbolInfo) Get_val() LnsAny{ return self.val }

// declaration Class -- Scope
type MetaLoader_ScopeMtd interface {
    Add(arg1 string, arg2 LnsAny) *MetaLoader_SymbolInfo
    Get(arg1 string) LnsAny
    Get_parent() LnsAny
}
type MetaLoader_Scope struct {
    parent LnsAny
    name2val *LnsMap
    FP MetaLoader_ScopeMtd
}
func MetaLoader_Scope2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*MetaLoader_Scope).FP
}
type MetaLoader_ScopeDownCast interface {
    ToMetaLoader_Scope() *MetaLoader_Scope
}
func MetaLoader_ScopeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(MetaLoader_ScopeDownCast)
    if ok { return work.ToMetaLoader_Scope() }
    return nil
}
func (obj *MetaLoader_Scope) ToMetaLoader_Scope() *MetaLoader_Scope {
    return obj
}
func NewMetaLoader_Scope(arg1 LnsAny) *MetaLoader_Scope {
    obj := &MetaLoader_Scope{}
    obj.FP = obj
    obj.InitMetaLoader_Scope(arg1)
    return obj
}
func (self *MetaLoader_Scope) Get_parent() LnsAny{ return self.parent }
// 15: DeclConstr
func (self *MetaLoader_Scope) InitMetaLoader_Scope(parent LnsAny) {
    self.parent = parent
    
    self.name2val = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 19: decl @lune.@base.@MetaLoader.Scope.add
func (self *MetaLoader_Scope) Add(name string,val LnsAny) *MetaLoader_SymbolInfo {
    var symbolInfo *MetaLoader_SymbolInfo
    symbolInfo = NewMetaLoader_SymbolInfo(name, val)
    self.name2val.Set(name,symbolInfo)
    return symbolInfo
}

// 24: decl @lune.@base.@MetaLoader.Scope.get
func (self *MetaLoader_Scope) Get(name string) LnsAny {
    {
        __exp := self.name2val.Items[name]
        if __exp != nil {
            _exp := __exp.(*MetaLoader_SymbolInfo)
            return _exp
        }
    }
    {
        __exp := self.parent
        if __exp != nil {
            _exp := __exp.(*MetaLoader_Scope)
            return _exp.FP.Get(name)
        }
    }
    return nil
}


// declaration Class -- Loader
type MetaLoader_LoaderMtd interface {
    analyzeExp()(LnsAny, LnsAny)
    analyzeTable()(LnsAny, LnsAny)
    Process()(LnsAny, LnsAny)
    processDeclVar() LnsAny
    processExpSym(arg1 *Parser_Token) LnsAny
    processStmt()(LnsAny, LnsAny)
}
type MetaLoader_Loader struct {
    parser Parser_PushbackParser
    curScope *MetaLoader_Scope
    FP MetaLoader_LoaderMtd
}
func MetaLoader_Loader2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*MetaLoader_Loader).FP
}
type MetaLoader_LoaderDownCast interface {
    ToMetaLoader_Loader() *MetaLoader_Loader
}
func MetaLoader_LoaderDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(MetaLoader_LoaderDownCast)
    if ok { return work.ToMetaLoader_Loader() }
    return nil
}
func (obj *MetaLoader_Loader) ToMetaLoader_Loader() *MetaLoader_Loader {
    return obj
}
func NewMetaLoader_Loader(arg1 Parser_PushbackParser) *MetaLoader_Loader {
    obj := &MetaLoader_Loader{}
    obj.FP = obj
    obj.InitMetaLoader_Loader(arg1)
    return obj
}
// 52: DeclConstr
func (self *MetaLoader_Loader) InitMetaLoader_Loader(parser Parser_PushbackParser) {
    self.parser = parser
    
    self.curScope = NewMetaLoader_Scope(nil)
    
}


// 59: decl @lune.@base.@MetaLoader.Loader.analyzeTable
func (self *MetaLoader_Loader) analyzeTable()(LnsAny, LnsAny) {
    var _map *LnsMap
    _map = NewLnsMap( map[LnsAny]LnsAny{})
    var index LnsInt
    index = 1
    for  {
        var token *Parser_Token
        token = self.parser.GetTokenNoErr()
        if token.Kind == Parser_TokenKind__Eof{
            return nil, "eof"
        }
        if _switch545 := token.Txt; _switch545 == "}" {
            return _map, nil
        } else if _switch545 == "," {
        } else {
            var item LnsAny
            item = nil
            if _switch445 := token.Kind; _switch445 == Parser_TokenKind__Symb {
                var nextToken *Parser_Token
                nextToken = self.parser.GetTokenNoErr()
                self.parser.PushbackToken(nextToken)
                if nextToken.Txt == "="{
                    item = token.Txt
                    
                } else { 
                    return nil, Lns_getVM().String_format("not support -- %s", []LnsAny{token.Txt})
                }
            } else if _switch445 == Parser_TokenKind__Dlmt {
                if token.Txt == "{"{
                    var err LnsAny
                    item, err = self.FP.analyzeTable()
                    
                    if Lns_isCondTrue( err){
                        return nil, err
                    }
                } else { 
                    return nil, Lns_getVM().String_format("not support -- %s", []LnsAny{token.Txt})
                }
            } else {
                self.parser.Pushback()
                var err LnsAny
                item, err = self.FP.analyzeExp()
                
                if Lns_isCondTrue( err){
                    return nil, err
                }
            }
            if Lns_op_not(item){
                self.parser.PushbackToken(token)
            }
            var nextToken *Parser_Token
            nextToken = self.parser.GetTokenNoErr()
            if nextToken.Txt == "="{
                var val LnsAny
                var work LnsAny
                val,work = self.FP.analyzeExp()
                if Lns_isCondTrue( work){
                    return nil, work
                }
                if item != nil && val != nil{
                    item_427 := item
                    val_428 := val
                    _map.Set(item_427,val_428)
                }
            } else { 
                self.parser.Pushback()
                if item != nil{
                    item_431 := item
                    _map.Set(index,item_431)
                }
                index = index + 1
                
            }
        }
    }
// insert a dummy
    return nil,nil
}

// 130: decl @lune.@base.@MetaLoader.Loader.analyzeExp
func (self *MetaLoader_Loader) analyzeExp()(LnsAny, LnsAny) {
    for  {
        var token *Parser_Token
        token = self.parser.GetTokenNoErr()
        if token.Kind == Parser_TokenKind__Eof{
            return nil, "eof"
        }
        var txt string
        txt = token.Txt
        if _switch874 := token.Kind; _switch874 == Parser_TokenKind__Dlmt {
            if txt == "{"{
                return self.FP.analyzeTable()
            }
            return nil, Lns_getVM().String_format("illegal delimit -- %s", []LnsAny{txt})
        } else if _switch874 == Parser_TokenKind__Char {
            if len(txt) == 1{
                return LnsInt(txt[1-1]), nil
            }
            return Lns_unwrap( MetaLoader_quotedChar2Code.Items[Lns_getVM().String_sub(txt,2, 2)]).(LnsInt), nil
        } else if _switch874 == Parser_TokenKind__Int {
            return Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(txt, nil), 0))), nil
        } else if _switch874 == Parser_TokenKind__Real {
            return Lns_unwrapDefault( Lns_tonumber(txt, nil), 0.0).(LnsReal), nil
        } else if _switch874 == Parser_TokenKind__Str {
            if LnsInt(txt[1-1]) == 96{
                return Lns_getVM().String_sub(txt,4, -4), nil
            }
            return Lns_getVM().String_sub(txt,2, -2), nil
        } else if _switch874 == Parser_TokenKind__Symb {
            {
                _symbolInfo := self.curScope.FP.Get(txt)
                if _symbolInfo != nil {
                    symbolInfo := _symbolInfo.(*MetaLoader_SymbolInfo)
                    return symbolInfo.FP.Get_val(), nil
                }
            }
            return nil, Lns_getVM().String_format("not found -- %s", []LnsAny{txt})
        } else if _switch874 == Parser_TokenKind__Kywd {
            if _switch819 := token.Txt; _switch819 == "true" {
                return true, nil
            } else if _switch819 == "false" {
                return false, nil
            }
            return nil, Lns_getVM().String_format("illegal keyword -- %s", []LnsAny{token.Txt})
        } else if _switch874 == Parser_TokenKind__Cmnt {
        } else if _switch874 == Parser_TokenKind__Eof || _switch874 == Parser_TokenKind__Type {
            return nil, Lns_getVM().String_format("illegal kind -- %s", []LnsAny{Parser_TokenKind_getTxt( token.Kind)})
        } else if _switch874 == Parser_TokenKind__Ope {
            return nil, Lns_getVM().String_format("not support -- %s", []LnsAny{txt})
        }
    }
// insert a dummy
    return nil,nil
}

// 191: decl @lune.@base.@MetaLoader.Loader.processDeclVar
func (self *MetaLoader_Loader) processDeclVar() LnsAny {
    var token *Parser_Token
    token = self.parser.GetTokenNoErr()
    if token.Kind != Parser_TokenKind__Symb{
        return Lns_getVM().String_format("no synbol -- %s", []LnsAny{token.Txt})
    }
    var nextToken *Parser_Token
    nextToken = self.parser.GetTokenNoErr()
    if nextToken.Txt != "="{
        self.parser.Pushback()
        return nil
    }
    var val LnsAny
    var err LnsAny
    val,err = self.FP.analyzeExp()
    if Lns_isCondTrue( err){
        return err
    }
    self.curScope.FP.Add(token.Txt, val)
    return nil
}

// 210: decl @lune.@base.@MetaLoader.Loader.processExpSym
func (self *MetaLoader_Loader) processExpSym(symToken *Parser_Token) LnsAny {
    var symbolInfo *MetaLoader_SymbolInfo
    
    {
        _symbolInfo := self.curScope.FP.Get(symToken.Txt)
        if _symbolInfo == nil{
            return Lns_getVM().String_format("not found -- %s", []LnsAny{symToken.Txt})
        } else {
            symbolInfo = _symbolInfo.(*MetaLoader_SymbolInfo)
        }
    }
    var nextToken *Parser_Token
    nextToken = self.parser.GetTokenNoErr()
    var exp LnsAny
    exp = symbolInfo.FP.Get_val()
    var index LnsAny
    index = nil
    for  {
        if nextToken.Kind == Parser_TokenKind__Eof{
            return "eof"
        }
        if _switch1177 := nextToken.Txt; _switch1177 == "." {
            var fieldToken *Parser_Token
            fieldToken = self.parser.GetTokenNoErr()
            if exp != nil{
                index = fieldToken.Txt
                
            } else {
                return "nil access"
            }
        } else if _switch1177 == "[" {
            var indexExp LnsAny
            var err LnsAny
            indexExp,err = self.FP.analyzeExp()
            if exp != nil{
                index = indexExp
                
            } else {
                return Lns_getVM().String_format("illegal index -- %s", []LnsAny{err})
            }
            var closeToken *Parser_Token
            closeToken = self.parser.GetTokenNoErr()
            if closeToken.Txt != "]"{
                return Lns_getVM().String_format("illegal token -- %s", []LnsAny{closeToken.Txt})
            }
        } else if _switch1177 == "=" {
            var val LnsAny
            var err LnsAny
            val,err = self.FP.analyzeExp()
            if Lns_isCondTrue( err){
                return err
            }
            if index != nil && exp != nil{
                index_499 := index
                exp_500 := exp
                var _map *LnsMap
                _map = exp_500.(*LnsMap)
                _map.Set(index_499,val)
            }
            return nil
        }
        nextToken = self.parser.GetTokenNoErr()
        
    }
// insert a dummy
    return nil
}

// 259: decl @lune.@base.@MetaLoader.Loader.processStmt
func (self *MetaLoader_Loader) processStmt()(LnsAny, LnsAny) {
    for  {
        var token *Parser_Token
        token = self.parser.GetTokenNoErr()
        if token.Kind == Parser_TokenKind__Eof{
            return nil, "eof"
        }
        if _switch1325 := token.Kind; _switch1325 == Parser_TokenKind__Symb {
            {
                __exp := self.FP.processExpSym(token)
                if __exp != nil {
                    _exp := __exp.(string)
                    return nil, _exp
                }
            }
        } else if _switch1325 == Parser_TokenKind__Kywd {
            if _switch1323 := token.Txt; _switch1323 == "local" {
                {
                    _err := self.FP.processDeclVar()
                    if _err != nil {
                        err := _err.(string)
                        return nil, err
                    }
                }
            } else if _switch1323 == "do" {
                self.curScope = NewMetaLoader_Scope(self.curScope)
                
            } else if _switch1323 == "end" {
                {
                    _parent := self.curScope.FP.Get_parent()
                    if _parent != nil {
                        parent := _parent.(*MetaLoader_Scope)
                        self.curScope = parent
                        
                    } else {
                        return nil, "underflow scope"
                    }
                }
            } else if _switch1323 == "return" {
                return self.FP.analyzeExp()
            }
        }
    }
// insert a dummy
    return nil,nil
}

// 297: decl @lune.@base.@MetaLoader.Loader.process
func (self *MetaLoader_Loader) Process()(LnsAny, LnsAny) {
    var val LnsAny
    var err LnsAny
    val,err = self.FP.processStmt()
    if err != nil{
        err_529 := err.(string)
        var pos *Parser_Position
        pos = self.parser.GetTokenNoErr().Pos
        return nil, Lns_getVM().String_format("%d:%d:%s", []LnsAny{pos.LineNo, pos.Column, err_529})
    }
    return val, err
}


func Lns_MetaLoader_init() {
    if init_MetaLoader { return }
    init_MetaLoader = true
    MetaLoader__mod__ = "@lune.@base.@MetaLoader"
    Lns_InitMod()
    Lns_Str_init()
    Lns_Util_init()
    Lns_Meta_init()
    Lns_Parser_init()
    Lns_LuaVer_init()
    Lns_Depend_init()
    MetaLoader_quotedChar2Code = NewLnsMap( map[LnsAny]LnsAny{"a":7,"b":8,"t":9,"n":10,"v":11,"f":12,"r":13,"\\":92,"\"":34,"'":39,})
}
func init() {
    init_MetaLoader = false
}
