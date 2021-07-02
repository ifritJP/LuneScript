// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Macro bool
var Macro__mod__ string
var Macro_toLuavalNoasync *Lns_luaValue
type Macro_toListEmptyLua func (_env *LnsEnv) *LnsList
type Macro_toLuavalLuaFunc func (_env *LnsEnv, arg1 LnsAny) LnsAny
type Macro_EvalMacroCallback func (_env *LnsEnv)
// for 40
func Macro_convExp1_40(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 722
func Macro_convExp3_17(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 768
func Macro_convExp3_359(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 804
func Macro_convExp3_527(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 38
func Macro_convExp1_12(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 53
func Macro_convExp1_80(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 76
func Macro_convExp1_163(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 595
func Macro_convExp2_477(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 669
func Macro_convExp2_820(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 9
func Macro_convExp2_919(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 84
func Macro_convExp2_1262(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 9
func Macro_convExp2_1362(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 84
func Macro_convExp2_1705(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 35: decl @lune.@base.@Macro.loadCode
func Macro_loadCode_0_(_env *LnsEnv, code string) LnsAny {
    var ret LnsAny
        var loaded LnsAny
        var mess LnsAny
        loaded,mess = _env.GetVM().Load(code, nil)
        if loaded != nil{
            loaded_6 := loaded.(*Lns_luaValue)
            {
                _obj := Macro_convExp1_40(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded_6,Lns_2DDD([]LnsAny{}))[0]))
                if !Lns_IsNil( _obj ) {
                    obj := _obj
                    ret = obj
                } else {
                    panic("failed to load")
                }
            }
        } else {
            panic(_env.GetVM().String_format("%s -- %s", []LnsAny{mess, code}))
        }
    return ret
}

// 52: decl @lune.@base.@Macro.runLuaOnLns
func Macro_runLuaOnLns_1_(_env *LnsEnv, code string,baseDir LnsAny,async bool)(LnsAny, string) {
    var loadFunc LnsAny
    var err string
    loadFunc,err = DependLuaOnLns_runLuaOnLns(_env, code, baseDir, async)
    if loadFunc != nil{
        loadFunc_15 := loadFunc.(*Lns_luaValue)
        var mod LnsAny
        mod = nil
        if async{
                mod = _env.GetVM().RunLoadedfunc(loadFunc_15,Lns_2DDD([]LnsAny{}))[0]
        } else { 
            Lns_LockEnvSync( _env, 61, func () {
                mod = _env.GetVM().RunLoadedfunc(loadFunc_15,Lns_2DDD([]LnsAny{}))[0]
            })
        }
        if mod != nil{
            mod_22 := mod
            return mod_22, ""
        }
        return nil, "load error"
    }
    return nil, err
}

// 73: decl @lune.@base.@Macro.runLuaOnLnsToMacroProc
func Macro_runLuaOnLnsToMacroProc_2_(_env *LnsEnv, code string,baseDir LnsAny,async bool)(LnsAny, string) {
    var luaObj LnsAny
    var err string
    luaObj,err = Macro_runLuaOnLns_1_(_env, code, baseDir, async)
    if luaObj != nil{
        luaObj_27 := luaObj
        return luaObj_27.(*Lns_luaValue), ""
    }
    return nil, err
}

// 151: decl @lune.@base.@Macro.getLiteralMacroVal
func Macro_getLiteralMacroVal_8_(_env *LnsEnv, obj LnsAny) LnsAny {
    switch _matchExp0 := obj.(type) {
    case *Nodes_Literal__Nil:
        return nil
    case *Nodes_Literal__Int:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Real:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Str:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Bool:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Symbol:
    val := _matchExp0.Val1
        return NewLnsList([]LnsAny{val})
    case *Nodes_Literal__Field:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__LIST:
    list := _matchExp0.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_8_(_env, item))
        }
        return newList
    case *Nodes_Literal__ARRAY:
    list := _matchExp0.Val1
        var newList *LnsList
        newList = NewLnsList([]LnsAny{})
        for _index, _item := range( list.Items ) {
            index := _index + 1
            item := _item
            newList.Set(index,Macro_getLiteralMacroVal_8_(_env, item))
        }
        return newList
    case *Nodes_Literal__SET:
    list := _matchExp0.Val1
        var newSet *LnsSet
        newSet = NewLnsSet([]LnsAny{})
        for _, _item := range( list.Items ) {
            item := _item
            {
                __exp := Macro_getLiteralMacroVal_8_(_env, item)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp
                    newSet.Add(_exp)
                }
            }
        }
        return newSet
    case *Nodes_Literal__MAP:
    _map := _matchExp0.Val1
        var newMap *LnsMap
        newMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _key, _val := range( _map.Items ) {
            key := _key
            val := _val
            var keyObj LnsAny
            keyObj = Macro_getLiteralMacroVal_8_(_env, key)
            var valObj LnsAny
            valObj = Macro_getLiteralMacroVal_8_(_env, val)
            if keyObj != nil && valObj != nil{
                keyObj_77 := keyObj
                valObj_78 := valObj
                newMap.Set(keyObj_77,valObj_78)
            }
        }
        return newMap
    }
// insert a dummy
    return nil
}

// 458: decl @lune.@base.@Macro.equalsType
func Macro_equalsType_14_(_env *LnsEnv, typeInfo *Ast_TypeInfo,builtinType *Ast_TypeInfo) bool {
    return typeInfo.FP.Get_srcTypeInfo(_env) == builtinType
}

// 858: decl @lune.@base.@Macro.expandVal
func Macro_expandVal_15_(_env *LnsEnv, tokenList *LnsList,workval LnsAny,pos Types_Position) LnsAny {
    if workval != nil{
        workval_63 := workval
        var val LnsAny
        val = workval_63
        if _switch0 := Lns_type(val); _switch0 == "boolean" {
            var token string
            token = _env.GetVM().String_format("%s", []LnsAny{val})
            var kind LnsInt
            kind = Types_TokenKind__Kywd
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, kind, token, pos, false, nil)))
        } else if _switch0 == "number" {
            var num string
            num = _env.GetVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
            var kind LnsInt
            kind = Types_TokenKind__Int
            if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(num, ".", 1, true))){
                kind = Types_TokenKind__Real
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, kind, num, pos, false, nil)))
        } else if _switch0 == "string" {
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, Types_TokenKind__Str, Parser_quoteStr(_env, val.(string)), pos, false, nil)))
        } else {
            return _env.GetVM().String_format("not support ,, List -- %s", []LnsAny{Lns_type(val)})
        }
    }
    return nil
}

// 904: decl @lune.@base.@Macro.pushbackTxt
func Macro_pushbackTxt_16_(_env *LnsEnv, pushbackParser Parser_PushbackParser,txtList *LnsList,streamName string,pos Types_Position) {
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    for _, _txt := range( txtList.Items ) {
        txt := _txt.(string)
        var parser *Parser_StreamParser
        parser = Parser_StreamParser_create(_env, &Types_ParserSrc__LnsCode{txt, _env.GetVM().String_format("macro symbol -- %s", []LnsAny{streamName}), nil}, false, nil, pos.Get_RawOrgPos(_env))
        var workParser *Parser_DefaultPushbackParser
        workParser = NewParser_DefaultPushbackParser(_env, &parser.Parser_Parser)
        for  {
            var worktoken *Types_Token
            worktoken = workParser.FP.GetTokenNoErr(_env, nil)
            if worktoken.Kind == Types_TokenKind__Eof{
                break
            }
            tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, worktoken.Kind, worktoken.Txt, pos, false, nil)))
        }
    }
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
            pushbackParser.PushbackToken(_env, tokenList.GetAt(index).(Types_TokenDownCast).ToTypes_Token())
            _forWork0 += _forDelta0
        }
    }
}

// 1297: decl @lune.@base.@Macro.nodeToCodeTxt
func Macro_nodeToCodeTxt(_env *LnsEnv, node *Nodes_Node,moduleTypeInfo *Ast_TypeInfo) string {
    var code string
    var memStream *Util_memStream
    memStream = NewUtil_memStream(_env)
    var formatter *Nodes_Filter
    formatter = Formatter_createFilter(_env, moduleTypeInfo, memStream.FP)
    node.FP.ProcessFilter(_env, formatter, Formatter_Opt2Stem(NewFormatter_Opt(_env, node)))
    code = memStream.FP.Get_txt(_env)
    return code
}



// 949: decl @lune.@base.@Macro.MacroCtrl.expandMacroVal.macroVal2strList
func MacroCtrl_expandMacroVal__macroVal2strList_1_(_env *LnsEnv, name string,macroVal *Nodes_MacroValInfo,workParser Parser_PushbackParser) *LnsList {
    var val LnsAny
    
    {
        _val := macroVal.Val
        if _val == nil{
            workParser.Error(_env, _env.GetVM().String_format("macroVal is nil -- %s", []LnsAny{name}))
        } else {
            val = _val
        }
    }
    if Lns_isCondTrue( macroVal.ArgNode){
        return val.(*LnsList)
    }
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    {
        __forsortCollection0 := val.(*LnsMap)
        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            item := __forsortCollection0.Items[ ___forsortKey0 ].(string)
            list.Insert(item)
        }
    }
    return list
}

// 123: decl @lune.@base.@Macro.MacroParser.createPosition
func (self *Macro_MacroParser) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return Types_Position_create(_env, lineNo, column, self.FP.GetStreamName(_env), self.overridePos)
}
// 128: decl @lune.@base.@Macro.MacroParser.getToken
func (self *Macro_MacroParser) GetToken(_env *LnsEnv) LnsAny {
    if self.tokenList.Len() < self.pos{
        return nil
    }
    var token *Types_Token
    token = self.tokenList.GetAt(self.pos).(Types_TokenDownCast).ToTypes_Token()
    self.pos = self.pos + 1
    {
        __ := self.overridePos
        if !Lns_IsNil( __ ) {
            return NewTypes_Token(_env, token.Kind, token.Txt, self.FP.CreatePosition(_env, token.Pos.LineNo, token.Pos.Column), token.Consecutive, token.FP.Get_commentList(_env))
        }
    }
    return token
}
// 145: decl @lune.@base.@Macro.MacroParser.getStreamName
func (self *Macro_MacroParser) GetStreamName(_env *LnsEnv) string {
    return self.name
}
// 217: decl @lune.@base.@Macro.ExtMacroInfo.getArgList
func (self *Macro_ExtMacroInfo) GetArgList(_env *LnsEnv) *LnsList {
    return self.argList
}
// 220: decl @lune.@base.@Macro.ExtMacroInfo.getTokenList
func (self *Macro_ExtMacroInfo) GetTokenList(_env *LnsEnv) *LnsList {
    return self.tokenList
}
// 249: decl @lune.@base.@Macro.MacroAnalyzeInfo.clone
func (self *Macro_MacroAnalyzeInfo) Clone(_env *LnsEnv) *Macro_MacroAnalyzeInfo {
    var obj *Macro_MacroAnalyzeInfo
    obj = NewMacro_MacroAnalyzeInfo(_env, self.typeInfo, self.mode)
    obj.argIndex = self.argIndex
    return obj
}
// 255: decl @lune.@base.@Macro.MacroAnalyzeInfo.equalsArgTypeList
func (self *Macro_MacroAnalyzeInfo) EqualsArgTypeList(_env *LnsEnv, argTypeList LnsAny) bool {
    return self.typeInfo.FP.Get_argTypeInfoList(_env) == argTypeList
}
// 258: decl @lune.@base.@Macro.MacroAnalyzeInfo.getCurArgType
func (self *Macro_MacroAnalyzeInfo) GetCurArgType(_env *LnsEnv) *Ast_TypeInfo {
    if self.typeInfo.FP.Get_argTypeInfoList(_env).Len() < self.argIndex{
        return Ast_builtinTypeNone
    }
    return self.typeInfo.FP.Get_argTypeInfoList(_env).GetAt(self.argIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}
// 264: decl @lune.@base.@Macro.MacroAnalyzeInfo.nextArg
func (self *Macro_MacroAnalyzeInfo) NextArg(_env *LnsEnv) {
    self.argIndex = self.argIndex + 1
}
// 268: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingSymArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingSymArg(_env *LnsEnv) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        _env.SetStackVal( self.FP.GetCurArgType(_env) == Ast_builtinTypeSymbol) ).(bool)
}
// 271: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingExpArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingExpArg(_env *LnsEnv) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        _env.SetStackVal( self.FP.GetCurArgType(_env) == Ast_builtinTypeExp) ).(bool)
}
// 274: decl @lune.@base.@Macro.MacroAnalyzeInfo.isAnalyzingBlockArg
func (self *Macro_MacroAnalyzeInfo) IsAnalyzingBlockArg(_env *LnsEnv) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.mode == Nodes_MacroMode__AnalyzeArg) &&
        _env.SetStackVal( self.FP.GetCurArgType(_env) == Ast_builtinTypeBlockArg) ).(bool)
}
// 336: decl @lune.@base.@Macro.MacroCtrl.get_isDeclaringMacro
func (self *Macro_MacroCtrl) Get_isDeclaringMacro(_env *LnsEnv) bool {
    return Ast_TypeInfo2Stem(self.declaringType) != nil
}
// 343: decl @lune.@base.@Macro.MacroCtrl.isUsing__var
func (self *Macro_MacroCtrl) IsUsing__var(_env *LnsEnv, macroType *Ast_TypeInfo) bool {
    var typeId *Ast_IdInfo
    typeId = macroType.FP.Get_typeId(_env)
    if Lns_isCondTrue( self.typeId2ImportedMacroInfo.Get(typeId)){
        return false
    }
    {
        __exp := self.id2use___var.Get(typeId)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(bool)
            return _exp
        }
    }
    Util_err(_env, _env.GetVM().String_format("unknown macro -- %s", []LnsAny{macroType.FP.GetTxt(_env, nil, nil, nil)}))
// insert a dummy
    return false
}
// 354: decl @lune.@base.@Macro.MacroCtrl.setUsing__var
func (self *Macro_MacroCtrl) SetUsing__var(_env *LnsEnv) {
    {
        __exp := self.declaringType
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            self.id2use___var.Set(_exp.FP.Get_typeId(_env),true)
        }
    }
}
// 391: decl @lune.@base.@Macro.MacroCtrl.clone
func (self *Macro_MacroCtrl) Clone(_env *LnsEnv) *Macro_MacroCtrl {
    var obj *Macro_MacroCtrl
    obj = NewMacro_MacroCtrl(_env, self.macroEval, self.validAsyncMacro)
    if Lns_op_not(self.validAsyncMacro){
        obj.toLuavalLuaAsync = self.toLuavalLuaAsync
    }
    obj.useLnsLoad = self.useLnsLoad
    obj.tokenExpanding = self.tokenExpanding
    obj.useModuleMacroSet.Or(self.useModuleMacroSet)
    {
        for _key, _val := range( self.typeId2ImportedMacroInfo.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            val := _val.(Nodes_MacroInfoDownCast).ToNodes_MacroInfo()
            obj.typeId2ImportedMacroInfo.Set(key,val)
        }
    }
    
    {
        for _key, _val := range( self.declPubMacroInfoMap.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            val := _val.(Nodes_MacroInfoDownCast).ToNodes_MacroInfo()
            obj.declPubMacroInfoMap.Set(key,val)
        }
    }
    
    {
        for _key, _val := range( self.id2use___var.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            val := _val.(bool)
            obj.id2use___var.Set(key,val)
        }
    }
    
    if self.validAsyncMacro{
        for _key, _defInfo := range( self.declMacroInfoMap.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            defInfo := _defInfo.(Macro_DefMacroInfoWithSrcDownCast).ToMacro_DefMacroInfoWithSrc()
            var srcInfo *Macro_DefMacroSrc
            srcInfo = defInfo.FP.get_srcInfo(_env)
            var stmtFunc LnsAny
            stmtFunc = nil
            obj.declMacroInfoMap.Set(key,NewMacro_DefMacroInfoWithSrc(_env, stmtFunc, srcInfo, srcInfo.FP.Get_symbol2MacroValInfoMap(_env)))
        }
    } else { 
        {
            for _key, _val := range( self.declMacroInfoMap.Items ) {
                key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
                val := _val.(Macro_DefMacroInfoWithSrcDownCast).ToMacro_DefMacroInfoWithSrc()
                obj.declMacroInfoMap.Set(key,val)
            }
        }
        
    }
    {
        for _key, _val := range( self.symbol2ValueMapForMacro.Items ) {
            key := _key.(string)
            val := _val.(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
            obj.symbol2ValueMapForMacro.Set(key,val)
        }
    }
    
    obj.analyzeInfo = self.analyzeInfo.FP.Clone(_env)
    obj.macroCallLineNo = self.macroCallLineNo
    for _, _val := range( self.macroAnalyzeInfoStack.Items ) {
        val := _val.(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
        obj.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(val))
    }
    return obj
}
// 447: decl @lune.@base.@Macro.MacroCtrl.mergeFrom
func (self *Macro_MacroCtrl) MergeFrom(_env *LnsEnv, macroCtrl *Macro_MacroCtrl) {
    self.useModuleMacroSet.Or(macroCtrl.useModuleMacroSet)
}
// 451: decl @lune.@base.@Macro.MacroCtrl.setToUseLnsLoad
func (self *Macro_MacroCtrl) SetToUseLnsLoad(_env *LnsEnv) {
    if Lns_isCondTrue( self.declaringType){
        self.useLnsLoad = true
    }
}
// 472: decl @lune.@base.@Macro.MacroCtrl.evalMacroOp
func (self *Macro_MacroCtrl) EvalMacroOp(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,streamName string,firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny)(LnsAny, LnsAny) {
    self.useModuleMacroSet.Add(Ast_TypeInfo2Stem(macroTypeInfo.FP.GetModule(_env)))
    if expList != nil{
        expList_6 := expList.(*Nodes_ExpListNode)
        for _, _exp := range( expList_6.FP.Get_expList(_env).Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            var kind LnsInt
            kind = exp.FP.Get_kind(_env)
            if _switch0 := kind; _switch0 == Nodes_NodeKind_get_LiteralNil(_env) || _switch0 == Nodes_NodeKind_get_LiteralChar(_env) || _switch0 == Nodes_NodeKind_get_LiteralInt(_env) || _switch0 == Nodes_NodeKind_get_LiteralReal(_env) || _switch0 == Nodes_NodeKind_get_LiteralArray(_env) || _switch0 == Nodes_NodeKind_get_LiteralList(_env) || _switch0 == Nodes_NodeKind_get_LiteralMap(_env) || _switch0 == Nodes_NodeKind_get_LiteralString(_env) || _switch0 == Nodes_NodeKind_get_LiteralBool(_env) || _switch0 == Nodes_NodeKind_get_LiteralSymbol(_env) || _switch0 == Nodes_NodeKind_get_RefField(_env) || _switch0 == Nodes_NodeKind_get_ExpMacroStat(_env) || _switch0 == Nodes_NodeKind_get_ExpMacroArgExp(_env) || _switch0 == Nodes_NodeKind_get_ExpOmitEnum(_env) || _switch0 == Nodes_NodeKind_get_ExpCast(_env) || _switch0 == Nodes_NodeKind_get_ExpOp2(_env) {
            } else {
                var mess string
                mess = _env.GetVM().String_format("Macro arguments must be literal value. -- %d:%d:%s", []LnsAny{exp.FP.Get_pos(_env).LineNo, exp.FP.Get_pos(_env).Column, Nodes_getNodeKindName(_env, kind)})
                return nil, mess
            }
        }
    }
    var macroInfo *Nodes_MacroInfo
    {
        __exp := _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.typeId2ImportedMacroInfo.Get(macroTypeInfo.FP.Get_typeId(_env))) ||
            _env.SetStackVal( self.declPubMacroInfoMap.Get(macroTypeInfo.FP.Get_typeId(_env))) )
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_MacroInfo)
            macroInfo = _exp
        } else {
            {
                __exp := self.declMacroInfoMap.Get(macroTypeInfo.FP.Get_typeId(_env))
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Macro_DefMacroInfoWithSrc)
                    var defInfo *Macro_DefMacroInfoWithSrc
                    defInfo = _exp
                    var srcInfo *Macro_DefMacroSrc
                    srcInfo = defInfo.FP.get_srcInfo(_env)
                    if Lns_op_not(defInfo.FP.Get_func(_env)){
                        {
                            _luaCode := srcInfo.FP.Get_luaCode(_env)
                            if !Lns_IsNil( _luaCode ) {
                                luaCode := _luaCode.(string)
                                var stmtFunc *Lns_luaValue
                                stmtFunc = Lns_unwrap( Lns_car(Macro_runLuaOnLnsToMacroProc_2_(_env, luaCode, srcInfo.FP.Get_baseDir(_env), srcInfo.FP.Get_asyncFlag(_env)))).(*Lns_luaValue)
                                defInfo.FP.Set_func(_env, stmtFunc)
                            }
                        }
                    }
                    macroInfo = &defInfo.Nodes_MacroInfo
                } else {
                    Util_err(_env, _env.GetVM().String_format("not found macroInfo -- %d", []LnsAny{macroTypeInfo.FP.Get_typeId(_env).Id}))
                }
            }
        }
    }
    var process func(_env *LnsEnv) LnsAny
    process = func(_env *LnsEnv) LnsAny {
        var argValMap *LnsMap
        argValMap = NewLnsMap( map[LnsAny]LnsAny{})
        var macroArgNodeList *LnsList
        macroArgNodeList = macroInfo.FP.GetArgList(_env)
        var macroArgName2ArgNode *LnsMap
        macroArgName2ArgNode = NewLnsMap( map[LnsAny]LnsAny{})
        var errmess LnsAny
        errmess = nil
        var innerMacro bool
        innerMacro = macroTypeInfo.FP.GetModule(_env) == moduleTypeInfo
        var asyncMacro bool
        asyncMacro = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.validAsyncMacro) &&
            _env.SetStackVal( innerMacro) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, macroTypeInfo.FP.Get_accessMode(_env)))) ).(bool)
        var valid__var bool
        valid__var = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( innerMacro) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, macroTypeInfo.FP.Get_accessMode(_env)))) ).(bool)
        var toLuaval *Lns_luaValue
        if asyncMacro{
            var work *Lns_luaValue
            
            {
                _work := self.toLuavalLuaAsync
                if _work == nil{
                    work = Macro_loadCode_0_(_env, "return function( val ) return val end").(*Lns_luaValue)
                    self.toLuavalLuaAsync = work
                } else {
                    work = _work.(*Lns_luaValue)
                }
            }
            toLuaval = work
        } else { 
            toLuaval = Macro_toLuavalNoasync
        }
        var macroArgValMap *LnsMap
        macroArgValMap = NewLnsMap( map[LnsAny]LnsAny{})
        if asyncMacro{
                {
                    {
                        __func := macroInfo.FP.Get_func(_env)
                        if !Lns_IsNil( __func ) {
                            _func := __func.(*Lns_luaValue)
                            if expList != nil{
                                expList_104 := expList.(*Nodes_ExpListNode)
                                for _index, _argNode := range( expList_104.FP.Get_expList(_env).Items ) {
                                    index := _index + 1
                                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                                    var declArgNode *Nodes_MacroArgInfo
                                    declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                                    macroArgName2ArgNode.Set(declArgNode.FP.Get_name(_env),argNode)
                                    var literal LnsAny
                                    var mess LnsAny
                                    literal,mess = argNode.FP.GetLiteral(_env)
                                    if literal != nil{
                                        literal_112 := literal
                                        {
                                            _val := Macro_getLiteralMacroVal_8_(_env, literal_112)
                                            if !Lns_IsNil( _val ) {
                                                val := _val
                                                argValMap.Set(index,val)
                                                if argNode.FP.Get_expType(_env) == Ast_builtinTypeSymbol{
                                                    macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.GetVM().RunLoadedfunc(toLuaval,Lns_2DDD(Lns_FromStemGetAt(val,1, false )))[0])
                                                } else { 
                                                    macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.GetVM().RunLoadedfunc(toLuaval,Lns_2DDD(val))[0])
                                                }
                                            }
                                        }
                                    } else {
                                        errmess = _env.GetVM().String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env)), mess})
                                        break
                                    }
                                }
                            }
                            if Lns_op_not(errmess){
                                var varMap LnsAny
                                
                                {
                                    _varMap := self.macroLocalVarMap
                                    if _varMap == nil{
                                        var toListEmpty *Lns_luaValue
                                        toListEmpty = Macro_loadCode_0_(_env, "return function() return {} end").(*Lns_luaValue)
                                        varMap = _env.GetVM().RunLoadedfunc(toListEmpty,[]LnsAny{})[0].(*Lns_luaValue)
                                    } else {
                                        varMap = _varMap
                                    }
                                }
                                if valid__var{
                                    macroArgValMap.Set("__var",varMap)
                                }
                                var macroVars LnsAny
                                macroVars = Lns_unwrap( _env.GetVM().ExpandLuavalMap(_env.GetVM().RunLoadedfunc(_func,Lns_2DDD(macroArgValMap))[0].(*Lns_luaValue)))
                                if valid__var{
                                    self.macroLocalVarMap = Lns_unwrap( Lns_FromStemGetAt(macroVars,"__var", false ))
                                }
                                for _, _name := range( (Lns_unwrap( Lns_FromStemGetAt(macroVars,"__names", false ))).(*LnsMap).Items ) {
                                    name := _name.(string)
                                    var valInfo *Nodes_MacroValInfo
                                    
                                    {
                                        _valInfo := macroInfo.FP.Get_symbol2MacroValInfoMap(_env).Get(name)
                                        if _valInfo == nil{
                                            Util_err(_env, _env.GetVM().String_format("not found macro symbol -- %s", []LnsAny{name}))
                                        } else {
                                            valInfo = _valInfo.(*Nodes_MacroValInfo)
                                        }
                                    }
                                    var typeInfo *Ast_TypeInfo
                                    typeInfo = valInfo.TypeInfo
                                    var valMap LnsAny
                                    {
                                        _val := Lns_FromStemGetAt(macroVars,name, false )
                                        if !Lns_IsNil( _val ) {
                                            val := _val
                                            if Macro_equalsType_14_(_env, typeInfo, Ast_builtinTypeSymbol){
                                                valMap = NewLnsMap( map[LnsAny]LnsAny{1:val,})
                                            } else { 
                                                valMap = val
                                            }
                                        } else {
                                            valMap = NewLnsMap( map[LnsAny]LnsAny{})
                                        }
                                    }
                                    self.symbol2ValueMapForMacro.Set(name,NewNodes_MacroValInfo(_env, valMap, typeInfo, nil))
                                }
                            }
                        } else {
                            if expList != nil{
                                expList_140 := expList.(*Nodes_ExpListNode)
                                for _index, _argNode := range( expList_140.FP.Get_expList(_env).Items ) {
                                    index := _index + 1
                                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                                    var declArgNode *Nodes_MacroArgInfo
                                    declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                                    macroArgName2ArgNode.Set(declArgNode.FP.Get_name(_env),argNode)
                                    var literal LnsAny
                                    var mess LnsAny
                                    literal,mess = argNode.FP.GetLiteral(_env)
                                    if literal != nil{
                                        literal_148 := literal
                                        {
                                            _val := Macro_getLiteralMacroVal_8_(_env, literal_148)
                                            if !Lns_IsNil( _val ) {
                                                val := _val
                                                argValMap.Set(index,val)
                                            }
                                        }
                                    } else {
                                        errmess = _env.GetVM().String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env)), mess})
                                        break
                                    }
                                }
                            }
                        }
                    }
                }
        } else { 
            Lns_LockEnvSync( _env, 583, func () {
                {
                    {
                        __func := macroInfo.FP.Get_func(_env)
                        if !Lns_IsNil( __func ) {
                            _func := __func.(*Lns_luaValue)
                            if expList != nil{
                                expList_158 := expList.(*Nodes_ExpListNode)
                                for _index, _argNode := range( expList_158.FP.Get_expList(_env).Items ) {
                                    index := _index + 1
                                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                                    var declArgNode *Nodes_MacroArgInfo
                                    declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                                    macroArgName2ArgNode.Set(declArgNode.FP.Get_name(_env),argNode)
                                    var literal LnsAny
                                    var mess LnsAny
                                    literal,mess = argNode.FP.GetLiteral(_env)
                                    if literal != nil{
                                        literal_166 := literal
                                        {
                                            _val := Macro_getLiteralMacroVal_8_(_env, literal_166)
                                            if !Lns_IsNil( _val ) {
                                                val := _val
                                                argValMap.Set(index,val)
                                                if argNode.FP.Get_expType(_env) == Ast_builtinTypeSymbol{
                                                    macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.GetVM().RunLoadedfunc(toLuaval,Lns_2DDD(Lns_FromStemGetAt(val,1, false )))[0])
                                                } else { 
                                                    macroArgValMap.Set(declArgNode.FP.Get_name(_env),_env.GetVM().RunLoadedfunc(toLuaval,Lns_2DDD(val))[0])
                                                }
                                            }
                                        }
                                    } else {
                                        errmess = _env.GetVM().String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env)), mess})
                                        break
                                    }
                                }
                            }
                            if Lns_op_not(errmess){
                                var varMap LnsAny
                                
                                {
                                    _varMap := self.macroLocalVarMap
                                    if _varMap == nil{
                                        var toListEmpty *Lns_luaValue
                                        toListEmpty = Macro_loadCode_0_(_env, "return function() return {} end").(*Lns_luaValue)
                                        varMap = _env.GetVM().RunLoadedfunc(toListEmpty,[]LnsAny{})[0].(*Lns_luaValue)
                                    } else {
                                        varMap = _varMap
                                    }
                                }
                                if valid__var{
                                    macroArgValMap.Set("__var",varMap)
                                }
                                var macroVars LnsAny
                                macroVars = Lns_unwrap( _env.GetVM().ExpandLuavalMap(_env.GetVM().RunLoadedfunc(_func,Lns_2DDD(macroArgValMap))[0].(*Lns_luaValue)))
                                if valid__var{
                                    self.macroLocalVarMap = Lns_unwrap( Lns_FromStemGetAt(macroVars,"__var", false ))
                                }
                                for _, _name := range( (Lns_unwrap( Lns_FromStemGetAt(macroVars,"__names", false ))).(*LnsMap).Items ) {
                                    name := _name.(string)
                                    var valInfo *Nodes_MacroValInfo
                                    
                                    {
                                        _valInfo := macroInfo.FP.Get_symbol2MacroValInfoMap(_env).Get(name)
                                        if _valInfo == nil{
                                            Util_err(_env, _env.GetVM().String_format("not found macro symbol -- %s", []LnsAny{name}))
                                        } else {
                                            valInfo = _valInfo.(*Nodes_MacroValInfo)
                                        }
                                    }
                                    var typeInfo *Ast_TypeInfo
                                    typeInfo = valInfo.TypeInfo
                                    var valMap LnsAny
                                    {
                                        _val := Lns_FromStemGetAt(macroVars,name, false )
                                        if !Lns_IsNil( _val ) {
                                            val := _val
                                            if Macro_equalsType_14_(_env, typeInfo, Ast_builtinTypeSymbol){
                                                valMap = NewLnsMap( map[LnsAny]LnsAny{1:val,})
                                            } else { 
                                                valMap = val
                                            }
                                        } else {
                                            valMap = NewLnsMap( map[LnsAny]LnsAny{})
                                        }
                                    }
                                    self.symbol2ValueMapForMacro.Set(name,NewNodes_MacroValInfo(_env, valMap, typeInfo, nil))
                                }
                            }
                        } else {
                            if expList != nil{
                                expList_194 := expList.(*Nodes_ExpListNode)
                                for _index, _argNode := range( expList_194.FP.Get_expList(_env).Items ) {
                                    index := _index + 1
                                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                                    var declArgNode *Nodes_MacroArgInfo
                                    declArgNode = macroArgNodeList.GetAt(index).(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
                                    macroArgName2ArgNode.Set(declArgNode.FP.Get_name(_env),argNode)
                                    var literal LnsAny
                                    var mess LnsAny
                                    literal,mess = argNode.FP.GetLiteral(_env)
                                    if literal != nil{
                                        literal_202 := literal
                                        {
                                            _val := Macro_getLiteralMacroVal_8_(_env, literal_202)
                                            if !Lns_IsNil( _val ) {
                                                val := _val
                                                argValMap.Set(index,val)
                                            }
                                        }
                                    } else {
                                        errmess = _env.GetVM().String_format("not support node at arg(%d) -- %s:%s", []LnsAny{index, Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env)), mess})
                                        break
                                    }
                                }
                            }
                        }
                    }
                }
            })
        }
        
        if errmess != nil{
            errmess_207 := errmess.(string)
            return errmess_207
        }
        for _index, _arg := range( macroInfo.FP.GetArgList(_env).Items ) {
            index := _index + 1
            arg := _arg.(Nodes_MacroArgInfoDownCast).ToNodes_MacroArgInfo()
            if arg.FP.Get_typeInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
                var argType *Ast_TypeInfo
                argType = arg.FP.Get_typeInfo(_env)
                var argName string
                argName = arg.FP.Get_name(_env)
                self.symbol2ValueMapForMacro.Set(argName,NewNodes_MacroValInfo(_env, argValMap.Get(index), argType, macroArgName2ArgNode.Get(argName)))
            } else { 
                return "not support ... in macro"
            }
        }
        return nil
    }
    {
        _mess := process(_env)
        if !Lns_IsNil( _mess ) {
            mess := _mess.(string)
            return nil, mess
        }
    }
    return &NewMacro_MacroParser(_env, macroInfo.FP.GetTokenList(_env), _env.GetVM().String_format("%s:%d:%d: (macro %s)", []LnsAny{streamName, firstToken.Pos.LineNo, firstToken.Pos.Column, macroTypeInfo.FP.GetTxt(_env, nil, nil, nil)}), firstToken.Pos.Get_orgPos(_env)).Parser_Parser, nil
}
// 715: decl @lune.@base.@Macro.MacroCtrl.importMacro
func (self *Macro_MacroCtrl) ImportMacro(_env *LnsEnv, processInfo *Ast_ProcessInfo,lnsPath string,macroInfoStem LnsAny,macroTypeInfo *Ast_TypeInfo,typeId2TypeInfo *LnsMap,importedMacroInfoMap *LnsMap,baseDir LnsAny) {
    var macroInfo LnsAny
    var err LnsAny
    macroInfo, err = Macro_MacroMetaInfo__fromStem_3_(_env, macroInfoStem,nil)
    if macroInfo != nil{
        macroInfo_4 := macroInfo.(*Macro_MacroMetaInfo)
        var orgPos Types_Position
        if macroInfo_4.Pos.Len() == 2{
            orgPos = NewTypes_Position(_env, macroInfo_4.Pos.GetAt(1).(LnsInt), macroInfo_4.Pos.GetAt(2).(LnsInt), lnsPath, nil)
        } else { 
            Util_err(_env, "macroInfo.pos is illegal")
        }
        var argList *LnsList
        argList = NewLnsList([]LnsAny{})
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        var symbol2MacroValInfoMap *LnsMap
        symbol2MacroValInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _argInfo := range( macroInfo_4.ArgList.Items ) {
            argInfo := _argInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var argTypeInfo *Ast_TypeInfo
            argTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(argInfo.TypeId)).(*Ast_TypeInfo)
            argList.Insert(Nodes_MacroArgInfo2Stem(NewNodes_MacroArgInfo(_env, argInfo.Name, argTypeInfo)))
            argNameList.Insert(argInfo.Name)
        }
        for _, _symInfo := range( macroInfo_4.SymList.Items ) {
            symInfo := _symInfo.(Macro_MacroMetaArgInfoDownCast).ToMacro_MacroMetaArgInfo()
            var symTypeInfo *Ast_TypeInfo
            symTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(symInfo.TypeId)).(*Ast_TypeInfo)
            symbol2MacroValInfoMap.Set(symInfo.Name,NewNodes_MacroValInfo(_env, nil, symTypeInfo, nil))
        }
        var tokenList *LnsList
        tokenList = NewLnsList([]LnsAny{})
        var lineNo LnsInt
        lineNo = 0
        var column LnsInt
        column = 1
        for _, _tokenInfo := range( macroInfo_4.TokenList.Items ) {
            tokenInfo := _tokenInfo.(*LnsList)
            var txt string
            txt = tokenInfo.GetAt(2).(string)
            if txt == "\n"{
                lineNo = lineNo + 1
                column = 1
            } else { 
                var pos Types_Position
                pos = Types_Position_create(_env, lineNo, column, _env.GetVM().String_format("macro:%s", []LnsAny{macroInfo_4.Name}), orgPos)
                tokenList.Insert(Types_Token2Stem(NewTypes_Token(_env, Lns_unwrap( Types_TokenKind__from(_env, Lns_forceCastInt(tokenInfo.GetAt(1)))).(LnsInt), txt, pos, false, nil)))
                column = column + len(txt) + 1
            }
        }
        var luaCode string
        luaCode = self.macroEval.FP.EvalFromCodeToLuaCode(_env, processInfo, macroInfo_4.Name, argNameList, macroInfo_4.StmtBlock)
        var stmtFunc LnsAny
        stmtFunc, err = Macro_runLuaOnLnsToMacroProc_2_(_env, luaCode, baseDir, false)
        if stmtFunc != nil{
            stmtFunc_29 := stmtFunc.(*Lns_luaValue)
            var extMacroInfo *Macro_ExtMacroInfo
            extMacroInfo = NewMacro_ExtMacroInfo(_env, macroInfo_4.Name, stmtFunc_29, symbol2MacroValInfoMap, argList, tokenList, baseDir)
            self.typeId2ImportedMacroInfo.Set(macroTypeInfo.FP.Get_typeId(_env),&extMacroInfo.Nodes_MacroInfo)
            importedMacroInfoMap.Set(macroTypeInfo.FP.Get_typeId(_env),&extMacroInfo.Nodes_MacroInfo)
            return 
        }
    }
    Util_errorLog(_env, _env.GetVM().String_format("macro load fail -- %s: %s ", []LnsAny{macroTypeInfo.FP.GetTxt(_env, nil, nil, nil), Lns_unwrapDefault( err, "").(string)}))
}
// 783: decl @lune.@base.@Macro.MacroCtrl.importMacroInfo
func (self *Macro_MacroCtrl) ImportMacroInfo(_env *LnsEnv, importedMacroInfoMap *LnsMap) {
    for _typeId, _macroInfo := range( importedMacroInfoMap.Items ) {
        typeId := _typeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
        macroInfo := _macroInfo.(Nodes_MacroInfoDownCast).ToNodes_MacroInfo()
        self.typeId2ImportedMacroInfo.Set(typeId,macroInfo)
    }
}
// 792: decl @lune.@base.@Macro.MacroCtrl.regist
func (self *Macro_MacroCtrl) Regist(_env *LnsEnv, processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode,macroScope *Ast_Scope,baseDir LnsAny) LnsAny {
    var luaCode LnsAny
    var stmtFunc LnsAny
    var err LnsAny
    var asyncFlag bool
    asyncFlag = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.validAsyncMacro) &&
        _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, node.FP.Get_expType(_env).FP.Get_accessMode(_env)))) ).(bool)
    var ok bool
    if Lns_isCondTrue( node.FP.Get_declInfo(_env).FP.Get_stmtBlock(_env)){
        var workCode string
        workCode = self.macroEval.FP.EvalToLuaCode(_env, processInfo, node)
        luaCode = workCode
        stmtFunc, err = Macro_runLuaOnLnsToMacroProc_2_(_env, workCode, baseDir, asyncFlag)
        if Lns_isCondTrue( stmtFunc){
            ok = true
            err = nil
        } else { 
            ok = false
        }
    } else { 
        ok = true
        stmtFunc, err = nil, nil
        luaCode = nil
    }
    if ok{
        var remap *LnsMap
        remap = NewLnsMap( map[LnsAny]LnsAny{})
        for _name, _macroValInfo := range( self.symbol2ValueMapForMacro.Items ) {
            name := _name.(string)
            macroValInfo := _macroValInfo.(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
            if Macro_equalsType_14_(_env, macroValInfo.TypeInfo, Ast_builtinTypeEmpty){
                {
                    _typeInfo := macroScope.FP.GetTypeInfoChild(_env, name)
                    if !Lns_IsNil( _typeInfo ) {
                        typeInfo := _typeInfo.(*Ast_TypeInfo)
                        remap.Set(name,NewNodes_MacroValInfo(_env, macroValInfo.Val, typeInfo, macroValInfo.ArgNode))
                    } else {
                        remap.Set(name,macroValInfo)
                    }
                }
            } else { 
                remap.Set(name,macroValInfo)
            }
        }
        var srcInfo *Macro_DefMacroSrc
        srcInfo = NewMacro_DefMacroSrc(_env, luaCode, node.FP.Get_declInfo(_env), remap, baseDir, asyncFlag)
        var macroInfo *Macro_DefMacroInfoWithSrc
        macroInfo = NewMacro_DefMacroInfoWithSrc(_env, stmtFunc, srcInfo, remap)
        if Ast_isPubToExternal(_env, node.FP.Get_expType(_env).FP.Get_accessMode(_env)){
            self.declPubMacroInfoMap.Set(node.FP.Get_expType(_env).FP.Get_typeId(_env),&macroInfo.Nodes_MacroInfo)
        } else { 
            self.declMacroInfoMap.Set(node.FP.Get_expType(_env).FP.Get_typeId(_env),macroInfo)
        }
    }
    if Lns_op_not(self.id2use___var.Get(node.FP.Get_expType(_env).FP.Get_typeId(_env))){
        self.id2use___var.Set(node.FP.Get_expType(_env).FP.Get_typeId(_env),false)
    }
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    self.declaringType = nil
    return err
}
// 934: decl @lune.@base.@Macro.MacroCtrl.expandMacroVal
func (self *Macro_MacroCtrl) ExpandMacroVal(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,scope *Ast_Scope,parser Parser_PushbackParser,token *Types_Token) *Types_Token {
    if self.tokenExpanding{
        return token
    }
    var getToken func(_env *LnsEnv) *Types_Token
    getToken = func(_env *LnsEnv) *Types_Token {
        self.tokenExpanding = true
        var work *Types_Token
        work = parser.GetTokenNoErr(_env, nil)
        self.tokenExpanding = false
        return work
    }
    var tokenTxt string
    tokenTxt = token.Txt
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( tokenTxt == ",,") ||
        _env.SetStackVal( tokenTxt == ",,,") ||
        _env.SetStackVal( tokenTxt == ",,,,") ).(bool){
        var nextToken *Types_Token
        nextToken = getToken(_env)
        var macroVal *Nodes_MacroValInfo
        
        {
            _macroVal := self.symbol2ValueMapForMacro.Get(nextToken.Txt)
            if _macroVal == nil{
                parser.Error(_env, _env.GetVM().String_format("unknown macro val -- %s", []LnsAny{nextToken.Txt}))
            } else {
                macroVal = _macroVal.(*Nodes_MacroValInfo)
            }
        }
        if tokenTxt == ",,"{
            if Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *LnsList
                txtList = NewLnsList([]LnsAny{})
                for _, _txt := range( MacroCtrl_expandMacroVal__macroVal2strList_1_(_env, nextToken.Txt, macroVal, parser).Items ) {
                    txt := _txt.(string)
                    txtList.Insert(txt)
                }
                Macro_pushbackTxt_16_(_env, parser, txtList, nextToken.Txt, nextToken.Pos)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeBlockArg)) ).(bool){
                var pos Types_Position
                pos = _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(macroVal.ArgNode) && 
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.Get_pos(_env)})&&
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(Types_Position).Get_RawOrgPos(_env)}))) ||
                    _env.SetStackVal( nextToken.Pos.Get_RawOrgPos(_env)) ||
                    _env.SetStackVal( token.Pos.Get_orgPos(_env)) ).(Types_Position)
                parser.PushbackStr(_env, nil, _env.GetVM().String_format("macroVal %s", []LnsAny{nextToken.Txt}), (Lns_unwrap( macroVal.Val)).(string), pos)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Array) ||
                _env.SetStackVal( macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool){
                if Macro_equalsType_14_(_env, macroVal.TypeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat){
                    var pos Types_Position
                    pos = _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(macroVal.ArgNode) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.Get_pos(_env)})&&
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(Types_Position).Get_RawOrgPos(_env)}))) ||
                        _env.SetStackVal( nextToken.Pos.Get_RawOrgPos(_env)) ||
                        _env.SetStackVal( token.Pos.Get_orgPos(_env)) ).(Types_Position)
                    var strList *LnsList
                    strList = MacroCtrl_expandMacroVal__macroVal2strList_1_(_env, nextToken.Txt, macroVal, parser)
                    {
                        var _forFrom0 LnsInt = strList.Len()
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
                            parser.PushbackStr(_env, nil, _env.GetVM().String_format("macroVal %s[%d]", []LnsAny{nextToken.Txt, index}), strList.GetAt(index).(string), pos)
                            _forWork0 += _forDelta0
                        }
                    }
                } else { 
                    var tokenList *LnsList
                    tokenList = NewLnsList([]LnsAny{})
                    {
                        _argNode := macroVal.ArgNode
                        if !Lns_IsNil( _argNode ) {
                            argNode := _argNode.(*Nodes_Node)
                            if Lns_op_not(argNode.FP.SetupLiteralTokenList(_env, tokenList)){
                                parser.Error(_env, _env.GetVM().String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                            }
                        } else {
                            parser.Error(_env, _env.GetVM().String_format("not support ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    }
                    parser.NewPushback(_env, tokenList)
                }
            } else if macroVal.TypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum{
                var enumTypeInfo *Ast_EnumTypeInfo
                enumTypeInfo = Lns_unwrap( Ast_EnumTypeInfoDownCastF(macroVal.TypeInfo.FP.Get_aliasSrc(_env).FP)).(*Ast_EnumTypeInfo)
                var fullname string
                fullname = macroVal.TypeInfo.FP.GetFullName(_env, typeNameCtrl, scope.FP, true)
                var nameList *LnsList
                nameList = Util_splitStr(_env, fullname, "[^%.]+")
                var enumValInfo *Ast_EnumValInfo
                enumValInfo = Lns_unwrap( enumTypeInfo.FP.Get_val2EnumValInfo(_env).Get(Lns_unwrap( macroVal.Val))).(*Ast_EnumValInfo)
                nextToken = NewTypes_Token(_env, Types_TokenKind__Symb, enumValInfo.FP.Get_name(_env), nextToken.Pos, false, nil)
                parser.PushbackToken(_env, nextToken)
                {
                    var _forFrom1 LnsInt = nameList.Len()
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
                        nextToken = NewTypes_Token(_env, Types_TokenKind__Dlmt, ".", nextToken.Pos, false, nil)
                        parser.PushbackToken(_env, nextToken)
                        nextToken = NewTypes_Token(_env, Types_TokenKind__Symb, nameList.GetAt(index).(string), nextToken.Pos, false, nil)
                        parser.PushbackToken(_env, nextToken)
                        _forWork1 += _forDelta1
                    }
                }
            } else { 
                var tokenList *LnsList
                tokenList = NewLnsList([]LnsAny{})
                {
                    _argNode := macroVal.ArgNode
                    if !Lns_IsNil( _argNode ) {
                        argNode := _argNode.(*Nodes_Node)
                        if Lns_op_not(argNode.FP.SetupLiteralTokenList(_env, tokenList)){
                            parser.Error(_env, _env.GetVM().String_format("illegal macro val ,, -- %s", []LnsAny{nextToken.Txt}))
                        }
                    } else {
                        Macro_expandVal_15_(_env, tokenList, macroVal.Val, nextToken.Pos)
                    }
                }
                parser.NewPushback(_env, tokenList)
            }
        } else if tokenTxt == ",,,"{
            if Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeString){
                Macro_pushbackTxt_16_(_env, parser, NewLnsList([]LnsAny{(Lns_unwrap( macroVal.Val)).(string)}), nextToken.Txt, nextToken.Pos)
            } else { 
                parser.Error(_env, _env.GetVM().String_format("',,,' does not support this type -- %s", []LnsAny{macroVal.TypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else if tokenTxt == ",,,,"{
            if Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeSymbol){
                var txtList *LnsList
                txtList = (Lns_unwrap( macroVal.Val)).(*LnsList)
                var newToken string
                newToken = ""
                for _, _txt := range( txtList.Items ) {
                    txt := _txt.(string)
                    newToken = _env.GetVM().String_format("%s%s", []LnsAny{newToken, txt})
                }
                nextToken = NewTypes_Token(_env, Types_TokenKind__Str, _env.GetVM().String_format("'%s'", []LnsAny{newToken}), nextToken.Pos, false, nil)
                parser.PushbackToken(_env, nextToken)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeStat)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeExp)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeMultiExp)) ||
                _env.SetStackVal( Macro_equalsType_14_(_env, macroVal.TypeInfo, Ast_builtinTypeBlockArg)) ).(bool){
                var txt string
                txt = (Lns_unwrap( macroVal.Val)).(string)
                var rawTxt string
                if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(txt,"^```", nil, nil))){
                    rawTxt = Parser_quoteStr(_env, txt)
                } else { 
                    rawTxt = Parser_quoteStr(_env, txt)
                }
                nextToken = NewTypes_Token(_env, Types_TokenKind__Str, rawTxt, nextToken.Pos, false, nil)
                parser.PushbackToken(_env, nextToken)
            } else { 
                parser.Error(_env, _env.GetVM().String_format("not support this symbol -- %s%s", []LnsAny{tokenTxt, nextToken.Txt}))
            }
        }
        nextToken = getToken(_env)
        token = nextToken
    }
    self.tokenExpanding = false
    return token
}
// 1120: decl @lune.@base.@Macro.MacroCtrl.expandSymbol
func (self *Macro_MacroCtrl) ExpandSymbol(_env *LnsEnv, parser Parser_PushbackParser,inTestBlock bool,prefixToken *Types_Token,exp *Nodes_Node,nodeManager *Nodes_NodeManager,errMessList *LnsList) *Nodes_LiteralStringNode {
    var nextToken *Types_Token
    nextToken = parser.GetTokenNoErr(_env, nil)
    if nextToken.Txt != "~~"{
        parser.PushbackToken(_env, nextToken)
    }
    var format string
    format = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( prefixToken.Txt == ",,,") &&
        _env.SetStackVal( "' %s '") ||
        _env.SetStackVal( "\"'%s'\"") ).(string)
    if prefixToken.Txt == ",,"{
        {
            _refNode := Nodes_ExpRefNodeDownCastF(exp.FP)
            if !Lns_IsNil( _refNode ) {
                refNode := _refNode.(*Nodes_ExpRefNode)
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = refNode.FP.Get_symbolInfo(_env)
                var macroInfo LnsAny
                macroInfo = self.symbol2ValueMapForMacro.Get(symbolInfo.FP.Get_name(_env))
                if macroInfo != nil{
                    macroInfo_325 := macroInfo.(*Nodes_MacroValInfo)
                    var valType *Ast_TypeInfo
                    valType = macroInfo_325.TypeInfo
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeSymbol)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeExp)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeMultiExp)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeBlockArg)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( valType.FP.Get_kind(_env) == Ast_TypeInfoKind__List) &&
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_builtinTypeStat)) ).(bool)){
                        format = "' %s '"
                        exp = &Nodes_ExpMacroStatListNode_create(_env, nodeManager, prefixToken.Pos, inTestBlock, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), exp).Nodes_Node
                    } else if Macro_equalsType_14_(_env, Ast_builtinTypeString, valType){
                    } else if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeInt)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, valType, Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                    } else { 
                        errMessList.Insert(Macro_ErrorMess2Stem(NewMacro_ErrorMess(_env, Lns_unwrap( symbolInfo.FP.Get_pos(_env)).(Types_Position), _env.GetVM().String_format("not support ,, -- %s", []LnsAny{valType.FP.GetTxt(_env, nil, nil, nil)}))))
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeInt)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeReal)) ).(bool){
                        format = "' %s' "
                    } else if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeSymbol)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeExp)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeMultiExp)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeBlockArg)) ||
                        _env.SetStackVal( Macro_equalsType_14_(_env, exp.FP.Get_expType(_env), Ast_builtinTypeStat)) ).(bool){
                        format = "' %s '"
                    }
                }
            }
        }
    }
    var newToken *Types_Token
    newToken = NewTypes_Token(_env, Types_TokenKind__Str, format, prefixToken.Pos, prefixToken.Consecutive, nil)
    var expListNode *Nodes_ExpListNode
    expListNode = Nodes_ExpListNode_create(_env, nodeManager, exp.FP.Get_pos(_env), inTestBlock, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), NewLnsList([]LnsAny{Nodes_Node2Stem(exp)}), nil, false)
    var dddNode *Nodes_ExpToDDDNode
    dddNode = Nodes_ExpToDDDNode_create(_env, nodeManager, exp.FP.Get_pos(_env), inTestBlock, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), expListNode)
    var literalStr *Nodes_LiteralStringNode
    literalStr = Nodes_LiteralStringNode_create(_env, nodeManager, prefixToken.Pos, inTestBlock, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), newToken, expListNode, Nodes_ExpListNode_create(_env, nodeManager, exp.FP.Get_pos(_env), inTestBlock, self.analyzeInfo.FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg, exp.FP.Get_expTypeList(_env), NewLnsList([]LnsAny{Nodes_ExpToDDDNode2Stem(dddNode)}), nil, false))
    return literalStr
}
// 1219: decl @lune.@base.@Macro.MacroCtrl.registVar
func (self *Macro_MacroCtrl) RegistVar(_env *LnsEnv, symbolList *LnsList) {
    for _, _symbolInfo := range( symbolList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var info *Nodes_MacroValInfo
        info = NewNodes_MacroValInfo(_env, nil, symbolInfo.FP.Get_typeInfo(_env), nil)
        self.symbol2ValueMapForMacro.Set(symbolInfo.FP.Get_name(_env),info)
    }
}
// 1227: decl @lune.@base.@Macro.MacroCtrl.startDecl
func (self *Macro_MacroCtrl) StartDecl(_env *LnsEnv, declaringType *Ast_TypeInfo) {
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    self.declaringType = declaringType
    self.useLnsLoad = false
}
// 1235: decl @lune.@base.@Macro.MacroCtrl.finishMacroMode
func (self *Macro_MacroCtrl) FinishMacroMode(_env *LnsEnv) {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
}
// 1241: decl @lune.@base.@Macro.MacroCtrl.startExpandMode
func (self *Macro_MacroCtrl) StartExpandMode(_env *LnsEnv, pos Types_Position,typeInfo *Ast_TypeInfo,callback Macro_EvalMacroCallback) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, typeInfo, Nodes_MacroMode__Expand)
    self.macroCallLineNo = pos
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
    callback(_env)
    self.FP.FinishMacroMode(_env)
}
// 1253: decl @lune.@base.@Macro.MacroCtrl.startAnalyzeArgMode
func (self *Macro_MacroCtrl) StartAnalyzeArgMode(_env *LnsEnv, macroFuncType *Ast_TypeInfo) {
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, macroFuncType, Nodes_MacroMode__AnalyzeArg)
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}
// 1258: decl @lune.@base.@Macro.MacroCtrl.switchMacroMode
func (self *Macro_MacroCtrl) SwitchMacroMode(_env *LnsEnv) {
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len() - 1).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
    self.macroAnalyzeInfoStack.Insert(Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo))
}
// 1263: decl @lune.@base.@Macro.MacroCtrl.restoreMacroMode
func (self *Macro_MacroCtrl) RestoreMacroMode(_env *LnsEnv) {
    self.macroAnalyzeInfoStack.Remove(nil)
    self.analyzeInfo = self.macroAnalyzeInfoStack.GetAt(self.macroAnalyzeInfoStack.Len()).(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
}
// 1268: decl @lune.@base.@Macro.MacroCtrl.isInMode
func (self *Macro_MacroCtrl) IsInMode(_env *LnsEnv, mode LnsInt) bool {
    if self.macroAnalyzeInfoStack.Len() == 0{
        return false
    }
    for _, _info := range( self.macroAnalyzeInfoStack.Items ) {
        info := _info.(Macro_MacroAnalyzeInfoDownCast).ToMacro_MacroAnalyzeInfo()
        if info.FP.Get_mode(_env) == mode{
            return true
        }
    }
    return false
}
// 1284: decl @lune.@base.@Macro.MacroCtrl.isInAnalyzeArgMode
func (self *Macro_MacroCtrl) IsInAnalyzeArgMode(_env *LnsEnv) bool {
    return self.FP.IsInMode(_env, Nodes_MacroMode__AnalyzeArg)
}
// 1292: decl @lune.@base.@Macro.MacroCtrl.isInExpandMode
func (self *Macro_MacroCtrl) IsInExpandMode(_env *LnsEnv) bool {
    return self.FP.IsInMode(_env, Nodes_MacroMode__Expand)
}
// declaration Class -- MacroMetaArgInfo
type Macro_MacroMetaArgInfoMtd interface {
    ToMap() *LnsMap
}
type Macro_MacroMetaArgInfo struct {
    Name string
    TypeId LnsInt
    FP Macro_MacroMetaArgInfoMtd
}
func Macro_MacroMetaArgInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroMetaArgInfo).FP
}
type Macro_MacroMetaArgInfoDownCast interface {
    ToMacro_MacroMetaArgInfo() *Macro_MacroMetaArgInfo
}
func Macro_MacroMetaArgInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroMetaArgInfoDownCast)
    if ok { return work.ToMacro_MacroMetaArgInfo() }
    return nil
}
func (obj *Macro_MacroMetaArgInfo) ToMacro_MacroMetaArgInfo() *Macro_MacroMetaArgInfo {
    return obj
}
func NewMacro_MacroMetaArgInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) *Macro_MacroMetaArgInfo {
    obj := &Macro_MacroMetaArgInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaArgInfo(_env, arg1, arg2)
    return obj
}
func (self *Macro_MacroMetaArgInfo) InitMacro_MacroMetaArgInfo(_env *LnsEnv, arg1 string, arg2 LnsInt) {
    self.Name = arg1
    self.TypeId = arg2
}
func (self *Macro_MacroMetaArgInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["typeId"] = Lns_ToCollection( self.TypeId )
    return obj
}
func (self *Macro_MacroMetaArgInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Macro_MacroMetaArgInfo__fromMap_2_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaArgInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaArgInfo__fromStem_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaArgInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaArgInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Macro_MacroMetaArgInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Macro_MacroMetaArgInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Macro_MacroMetaArgInfo{}
    newObj.FP = newObj
    return Macro_MacroMetaArgInfo_FromMapMain( newObj, objMap, paramList )
}
func Macro_MacroMetaArgInfo_FromMapMain( newObj *Macro_MacroMetaArgInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["typeId"], false, nil); !ok {
       return false,nil,"typeId:" + mess.(string)
    } else {
       newObj.TypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- MacroMetaInfo
type Macro_MacroMetaInfoMtd interface {
    ToMap() *LnsMap
}
type Macro_MacroMetaInfo struct {
    Name string
    Pos *LnsList
    ArgList *LnsList
    SymList *LnsList
    StmtBlock LnsAny
    TokenList *LnsList
    FP Macro_MacroMetaInfoMtd
}
func Macro_MacroMetaInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroMetaInfo).FP
}
type Macro_MacroMetaInfoDownCast interface {
    ToMacro_MacroMetaInfo() *Macro_MacroMetaInfo
}
func Macro_MacroMetaInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroMetaInfoDownCast)
    if ok { return work.ToMacro_MacroMetaInfo() }
    return nil
}
func (obj *Macro_MacroMetaInfo) ToMacro_MacroMetaInfo() *Macro_MacroMetaInfo {
    return obj
}
func NewMacro_MacroMetaInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) *Macro_MacroMetaInfo {
    obj := &Macro_MacroMetaInfo{}
    obj.FP = obj
    obj.InitMacro_MacroMetaInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Macro_MacroMetaInfo) InitMacro_MacroMetaInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *LnsList, arg4 *LnsList, arg5 LnsAny, arg6 *LnsList) {
    self.Name = arg1
    self.Pos = arg2
    self.ArgList = arg3
    self.SymList = arg4
    self.StmtBlock = arg5
    self.TokenList = arg6
}
func (self *Macro_MacroMetaInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["pos"] = Lns_ToCollection( self.Pos )
    obj.Items["argList"] = Lns_ToCollection( self.ArgList )
    obj.Items["symList"] = Lns_ToCollection( self.SymList )
    obj.Items["stmtBlock"] = Lns_ToCollection( self.StmtBlock )
    obj.Items["tokenList"] = Lns_ToCollection( self.TokenList )
    return obj
}
func (self *Macro_MacroMetaInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Macro_MacroMetaInfo__fromMap_2_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaInfo__fromStem_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Macro_MacroMetaInfo_FromMap( arg1, paramList )
}
func Macro_MacroMetaInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Macro_MacroMetaInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Macro_MacroMetaInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Macro_MacroMetaInfo{}
    newObj.FP = newObj
    return Macro_MacroMetaInfo_FromMapMain( newObj, objMap, paramList )
}
func Macro_MacroMetaInfo_FromMapMain( newObj *Macro_MacroMetaInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["pos"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"pos:" + mess.(string)
    } else {
       newObj.Pos = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["argList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Macro_MacroMetaArgInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"argList:" + mess.(string)
    } else {
       newObj.ArgList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["symList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Macro_MacroMetaArgInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"symList:" + mess.(string)
    } else {
       newObj.SymList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["stmtBlock"], true, nil); !ok {
       return false,nil,"stmtBlock:" + mess.(string)
    } else {
       newObj.StmtBlock = conv
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["tokenList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToListSub, false,[]Lns_ToObjParam{Lns_ToObjParam{
                    Lns_ToStemSub, false,nil}}}}); !ok {
       return false,nil,"tokenList:" + mess.(string)
    } else {
       newObj.TokenList = conv.(*LnsList)
    }
    return true, newObj, nil
}

// declaration Class -- MacroParser
type Macro_MacroParserMtd interface {
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    GetStreamName(_env *LnsEnv) string
    GetToken(_env *LnsEnv) LnsAny
}
type Macro_MacroParser struct {
    Parser_Parser
    tokenList *LnsList
    pos LnsInt
    name string
    overridePos LnsAny
    FP Macro_MacroParserMtd
}
func Macro_MacroParser2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroParser).FP
}
type Macro_MacroParserDownCast interface {
    ToMacro_MacroParser() *Macro_MacroParser
}
func Macro_MacroParserDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroParserDownCast)
    if ok { return work.ToMacro_MacroParser() }
    return nil
}
func (obj *Macro_MacroParser) ToMacro_MacroParser() *Macro_MacroParser {
    return obj
}
func NewMacro_MacroParser(_env *LnsEnv, arg1 *LnsList, arg2 string, arg3 LnsAny) *Macro_MacroParser {
    obj := &Macro_MacroParser{}
    obj.FP = obj
    obj.Parser_Parser.FP = obj
    obj.InitMacro_MacroParser(_env, arg1, arg2, arg3)
    return obj
}
// 113: DeclConstr
func (self *Macro_MacroParser) InitMacro_MacroParser(_env *LnsEnv, tokenList *LnsList,name string,overridePos LnsAny) {
    self.InitParser_Parser(_env)
    self.pos = 1
    self.tokenList = tokenList
    self.name = name
    self.overridePos = _env.NilAccFin(_env.NilAccPush(overridePos) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(Types_Position).Get_orgPos(_env)}))
}


// declaration Class -- ExtMacroInfo
type Macro_ExtMacroInfoMtd interface {
    GetArgList(_env *LnsEnv) *LnsList
    GetTokenList(_env *LnsEnv) *LnsList
    Get_baseDir(_env *LnsEnv) LnsAny
    Get_func(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) string
    Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap
    Set_func(_env *LnsEnv, arg1 LnsAny)
}
type Macro_ExtMacroInfo struct {
    Nodes_MacroInfo
    name string
    argList *LnsList
    tokenList *LnsList
    baseDir LnsAny
    FP Macro_ExtMacroInfoMtd
}
func Macro_ExtMacroInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_ExtMacroInfo).FP
}
type Macro_ExtMacroInfoDownCast interface {
    ToMacro_ExtMacroInfo() *Macro_ExtMacroInfo
}
func Macro_ExtMacroInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_ExtMacroInfoDownCast)
    if ok { return work.ToMacro_ExtMacroInfo() }
    return nil
}
func (obj *Macro_ExtMacroInfo) ToMacro_ExtMacroInfo() *Macro_ExtMacroInfo {
    return obj
}
func NewMacro_ExtMacroInfo(_env *LnsEnv, arg1 string, arg2 *Lns_luaValue, arg3 *LnsMap, arg4 *LnsList, arg5 *LnsList, arg6 LnsAny) *Macro_ExtMacroInfo {
    obj := &Macro_ExtMacroInfo{}
    obj.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitMacro_ExtMacroInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Macro_ExtMacroInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Macro_ExtMacroInfo) Get_baseDir(_env *LnsEnv) LnsAny{ return self.baseDir }
// 224: DeclConstr
func (self *Macro_ExtMacroInfo) InitMacro_ExtMacroInfo(_env *LnsEnv, name string,_func *Lns_luaValue,symbol2MacroValInfoMap *LnsMap,argList *LnsList,tokenList *LnsList,baseDir LnsAny) {
    self.InitNodes_MacroInfo(_env, _func, symbol2MacroValInfoMap)
    self.name = name
    self.argList = argList
    self.tokenList = tokenList
    self.baseDir = baseDir
}


// declaration Class -- MacroAnalyzeInfo
type Macro_MacroAnalyzeInfoMtd interface {
    Clone(_env *LnsEnv) *Macro_MacroAnalyzeInfo
    EqualsArgTypeList(_env *LnsEnv, arg1 LnsAny) bool
    GetCurArgType(_env *LnsEnv) *Ast_TypeInfo
    Get_argIndex(_env *LnsEnv) LnsInt
    Get_mode(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    IsAnalyzingBlockArg(_env *LnsEnv) bool
    IsAnalyzingExpArg(_env *LnsEnv) bool
    IsAnalyzingSymArg(_env *LnsEnv) bool
    NextArg(_env *LnsEnv)
}
type Macro_MacroAnalyzeInfo struct {
    typeInfo *Ast_TypeInfo
    mode LnsInt
    argIndex LnsInt
    FP Macro_MacroAnalyzeInfoMtd
}
func Macro_MacroAnalyzeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroAnalyzeInfo).FP
}
type Macro_MacroAnalyzeInfoDownCast interface {
    ToMacro_MacroAnalyzeInfo() *Macro_MacroAnalyzeInfo
}
func Macro_MacroAnalyzeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroAnalyzeInfoDownCast)
    if ok { return work.ToMacro_MacroAnalyzeInfo() }
    return nil
}
func (obj *Macro_MacroAnalyzeInfo) ToMacro_MacroAnalyzeInfo() *Macro_MacroAnalyzeInfo {
    return obj
}
func NewMacro_MacroAnalyzeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Macro_MacroAnalyzeInfo {
    obj := &Macro_MacroAnalyzeInfo{}
    obj.FP = obj
    obj.InitMacro_MacroAnalyzeInfo(_env, arg1, arg2)
    return obj
}
func (self *Macro_MacroAnalyzeInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *Macro_MacroAnalyzeInfo) Get_mode(_env *LnsEnv) LnsInt{ return self.mode }
func (self *Macro_MacroAnalyzeInfo) Get_argIndex(_env *LnsEnv) LnsInt{ return self.argIndex }
// 243: DeclConstr
func (self *Macro_MacroAnalyzeInfo) InitMacro_MacroAnalyzeInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,mode LnsInt) {
    self.typeInfo = typeInfo
    self.mode = mode
    self.argIndex = 1
}


// declaration Class -- DefMacroSrc
type Macro_DefMacroSrcMtd interface {
    Get_asyncFlag(_env *LnsEnv) bool
    Get_baseDir(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclMacroInfo
    Get_luaCode(_env *LnsEnv) LnsAny
    Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap
}
type Macro_DefMacroSrc struct {
    luaCode LnsAny
    declInfo *Nodes_DeclMacroInfo
    symbol2MacroValInfoMap *LnsMap
    baseDir LnsAny
    asyncFlag bool
    FP Macro_DefMacroSrcMtd
}
func Macro_DefMacroSrc2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_DefMacroSrc).FP
}
type Macro_DefMacroSrcDownCast interface {
    ToMacro_DefMacroSrc() *Macro_DefMacroSrc
}
func Macro_DefMacroSrcDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_DefMacroSrcDownCast)
    if ok { return work.ToMacro_DefMacroSrc() }
    return nil
}
func (obj *Macro_DefMacroSrc) ToMacro_DefMacroSrc() *Macro_DefMacroSrc {
    return obj
}
func NewMacro_DefMacroSrc(_env *LnsEnv, arg1 LnsAny, arg2 *Nodes_DeclMacroInfo, arg3 *LnsMap, arg4 LnsAny, arg5 bool) *Macro_DefMacroSrc {
    obj := &Macro_DefMacroSrc{}
    obj.FP = obj
    obj.InitMacro_DefMacroSrc(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Macro_DefMacroSrc) InitMacro_DefMacroSrc(_env *LnsEnv, arg1 LnsAny, arg2 *Nodes_DeclMacroInfo, arg3 *LnsMap, arg4 LnsAny, arg5 bool) {
    self.luaCode = arg1
    self.declInfo = arg2
    self.symbol2MacroValInfoMap = arg3
    self.baseDir = arg4
    self.asyncFlag = arg5
}
func (self *Macro_DefMacroSrc) Get_luaCode(_env *LnsEnv) LnsAny{ return self.luaCode }
func (self *Macro_DefMacroSrc) Get_declInfo(_env *LnsEnv) *Nodes_DeclMacroInfo{ return self.declInfo }
func (self *Macro_DefMacroSrc) Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap{ return self.symbol2MacroValInfoMap }
func (self *Macro_DefMacroSrc) Get_baseDir(_env *LnsEnv) LnsAny{ return self.baseDir }
func (self *Macro_DefMacroSrc) Get_asyncFlag(_env *LnsEnv) bool{ return self.asyncFlag }

// declaration Class -- DefMacroInfoWithSrc
type Macro_DefMacroInfoWithSrcMtd interface {
    GetArgList(_env *LnsEnv) *LnsList
    GetTokenList(_env *LnsEnv) *LnsList
    Get_func(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) string
    get_srcInfo(_env *LnsEnv) *Macro_DefMacroSrc
    Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap
    Set_func(_env *LnsEnv, arg1 LnsAny)
}
type Macro_DefMacroInfoWithSrc struct {
    Nodes_DefMacroInfo
    srcInfo *Macro_DefMacroSrc
    FP Macro_DefMacroInfoWithSrcMtd
}
func Macro_DefMacroInfoWithSrc2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_DefMacroInfoWithSrc).FP
}
type Macro_DefMacroInfoWithSrcDownCast interface {
    ToMacro_DefMacroInfoWithSrc() *Macro_DefMacroInfoWithSrc
}
func Macro_DefMacroInfoWithSrcDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_DefMacroInfoWithSrcDownCast)
    if ok { return work.ToMacro_DefMacroInfoWithSrc() }
    return nil
}
func (obj *Macro_DefMacroInfoWithSrc) ToMacro_DefMacroInfoWithSrc() *Macro_DefMacroInfoWithSrc {
    return obj
}
func NewMacro_DefMacroInfoWithSrc(_env *LnsEnv, arg1 LnsAny, arg2 *Macro_DefMacroSrc, arg3 *LnsMap) *Macro_DefMacroInfoWithSrc {
    obj := &Macro_DefMacroInfoWithSrc{}
    obj.FP = obj
    obj.Nodes_DefMacroInfo.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitMacro_DefMacroInfoWithSrc(_env, arg1, arg2, arg3)
    return obj
}
func (self *Macro_DefMacroInfoWithSrc) get_srcInfo(_env *LnsEnv) *Macro_DefMacroSrc{ return self.srcInfo }
// 291: DeclConstr
func (self *Macro_DefMacroInfoWithSrc) InitMacro_DefMacroInfoWithSrc(_env *LnsEnv, _func LnsAny,srcInfo *Macro_DefMacroSrc,symbol2MacroValInfoMap *LnsMap) {
    self.InitNodes_DefMacroInfo(_env, _func, srcInfo.FP.Get_declInfo(_env), symbol2MacroValInfoMap)
    self.srcInfo = srcInfo
}


// declaration Class -- MacroCtrl
type Macro_MacroCtrlMtd interface {
    Clone(_env *LnsEnv) *Macro_MacroCtrl
    EvalMacroOp(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 string, arg3 *Types_Token, arg4 *Ast_TypeInfo, arg5 LnsAny)(LnsAny, LnsAny)
    ExpandMacroVal(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 *Ast_Scope, arg3 Parser_PushbackParser, arg4 *Types_Token) *Types_Token
    ExpandSymbol(_env *LnsEnv, arg1 Parser_PushbackParser, arg2 bool, arg3 *Types_Token, arg4 *Nodes_Node, arg5 *Nodes_NodeManager, arg6 *LnsList) *Nodes_LiteralStringNode
    FinishMacroMode(_env *LnsEnv)
    Get_analyzeInfo(_env *LnsEnv) *Macro_MacroAnalyzeInfo
    Get_declPubMacroInfoMap(_env *LnsEnv) *LnsMap
    Get_declaringType(_env *LnsEnv) LnsAny
    Get_isDeclaringMacro(_env *LnsEnv) bool
    Get_macroCallLineNo(_env *LnsEnv) LnsAny
    Get_tokenExpanding(_env *LnsEnv) bool
    Get_useModuleMacroSet(_env *LnsEnv) *LnsSet
    ImportMacro(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 *LnsMap, arg6 *LnsMap, arg7 LnsAny)
    ImportMacroInfo(_env *LnsEnv, arg1 *LnsMap)
    IsInAnalyzeArgMode(_env *LnsEnv) bool
    IsInExpandMode(_env *LnsEnv) bool
    IsInMode(_env *LnsEnv, arg1 LnsInt) bool
    IsUsing__var(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    MergeFrom(_env *LnsEnv, arg1 *Macro_MacroCtrl)
    Regist(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode, arg3 *Ast_Scope, arg4 LnsAny) LnsAny
    RegistVar(_env *LnsEnv, arg1 *LnsList)
    RestoreMacroMode(_env *LnsEnv)
    SetToUseLnsLoad(_env *LnsEnv)
    SetUsing__var(_env *LnsEnv)
    StartAnalyzeArgMode(_env *LnsEnv, arg1 *Ast_TypeInfo)
    StartDecl(_env *LnsEnv, arg1 *Ast_TypeInfo)
    StartExpandMode(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 Macro_EvalMacroCallback)
    SwitchMacroMode(_env *LnsEnv)
}
type Macro_MacroCtrl struct {
    useModuleMacroSet *LnsSet
    macroEval *Nodes_MacroEval
    typeId2ImportedMacroInfo *LnsMap
    declPubMacroInfoMap *LnsMap
    declMacroInfoMap *LnsMap
    symbol2ValueMapForMacro *LnsMap
    analyzeInfo *Macro_MacroAnalyzeInfo
    macroAnalyzeInfoStack *LnsList
    tokenExpanding bool
    macroCallLineNo LnsAny
    macroLocalVarMap LnsAny
    declaringType LnsAny
    id2use___var *LnsMap
    useLnsLoad bool
    toLuavalLuaAsync LnsAny
    validAsyncMacro bool
    FP Macro_MacroCtrlMtd
}
func Macro_MacroCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_MacroCtrl).FP
}
type Macro_MacroCtrlDownCast interface {
    ToMacro_MacroCtrl() *Macro_MacroCtrl
}
func Macro_MacroCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_MacroCtrlDownCast)
    if ok { return work.ToMacro_MacroCtrl() }
    return nil
}
func (obj *Macro_MacroCtrl) ToMacro_MacroCtrl() *Macro_MacroCtrl {
    return obj
}
func NewMacro_MacroCtrl(_env *LnsEnv, arg1 *Nodes_MacroEval, arg2 bool) *Macro_MacroCtrl {
    obj := &Macro_MacroCtrl{}
    obj.FP = obj
    obj.InitMacro_MacroCtrl(_env, arg1, arg2)
    return obj
}
func (self *Macro_MacroCtrl) Get_useModuleMacroSet(_env *LnsEnv) *LnsSet{ return self.useModuleMacroSet }
func (self *Macro_MacroCtrl) Get_declPubMacroInfoMap(_env *LnsEnv) *LnsMap{ return self.declPubMacroInfoMap }
func (self *Macro_MacroCtrl) Get_analyzeInfo(_env *LnsEnv) *Macro_MacroAnalyzeInfo{ return self.analyzeInfo }
func (self *Macro_MacroCtrl) Get_tokenExpanding(_env *LnsEnv) bool{ return self.tokenExpanding }
func (self *Macro_MacroCtrl) Get_macroCallLineNo(_env *LnsEnv) LnsAny{ return self.macroCallLineNo }
func (self *Macro_MacroCtrl) Get_declaringType(_env *LnsEnv) LnsAny{ return self.declaringType }
// 369: DeclConstr
func (self *Macro_MacroCtrl) InitMacro_MacroCtrl(_env *LnsEnv, macroEval *Nodes_MacroEval,validAsyncMacro bool) {
    self.id2use___var = NewLnsMap( map[LnsAny]LnsAny{})
    self.validAsyncMacro = validAsyncMacro
    self.toLuavalLuaAsync = nil
    self.useLnsLoad = false
    self.declMacroInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    self.declPubMacroInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    self.declaringType = nil
    self.tokenExpanding = false
    self.useModuleMacroSet = NewLnsSet([]LnsAny{})
    self.typeId2ImportedMacroInfo = NewLnsMap( map[LnsAny]LnsAny{})
    self.symbol2ValueMapForMacro = NewLnsMap( map[LnsAny]LnsAny{})
    self.macroEval = macroEval
    self.analyzeInfo = NewMacro_MacroAnalyzeInfo(_env, Ast_builtinTypeNone, Nodes_MacroMode__None)
    self.macroCallLineNo = nil
    self.macroAnalyzeInfoStack = NewLnsList([]LnsAny{Macro_MacroAnalyzeInfo2Stem(self.analyzeInfo)})
    self.macroLocalVarMap = nil
}


// declaration Class -- ErrorMess
type Macro_ErrorMessMtd interface {
}
type Macro_ErrorMess struct {
    Pos Types_Position
    Mess string
    FP Macro_ErrorMessMtd
}
func Macro_ErrorMess2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Macro_ErrorMess).FP
}
type Macro_ErrorMessDownCast interface {
    ToMacro_ErrorMess() *Macro_ErrorMess
}
func Macro_ErrorMessDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Macro_ErrorMessDownCast)
    if ok { return work.ToMacro_ErrorMess() }
    return nil
}
func (obj *Macro_ErrorMess) ToMacro_ErrorMess() *Macro_ErrorMess {
    return obj
}
func NewMacro_ErrorMess(_env *LnsEnv, arg1 Types_Position, arg2 string) *Macro_ErrorMess {
    obj := &Macro_ErrorMess{}
    obj.FP = obj
    obj.InitMacro_ErrorMess(_env, arg1, arg2)
    return obj
}
func (self *Macro_ErrorMess) InitMacro_ErrorMess(_env *LnsEnv, arg1 Types_Position, arg2 string) {
    self.Pos = arg1
    self.Mess = arg2
}

func Lns_Macro_init(_env *LnsEnv) {
    if init_Macro { return }
    init_Macro = true
    Macro__mod__ = "@lune.@base.@Macro"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Nodes_init(_env)
    Lns_Ast_init(_env)
    Lns_Parser_init(_env)
    Lns_Types_init(_env)
    Lns_Formatter_init(_env)
    Lns_DependLuaOnLns_init(_env)
    Macro_toLuavalNoasync = Macro_loadCode_0_(_env, "return function( val ) return val end").(*Lns_luaValue)
}
func init() {
    init_Macro = false
}
