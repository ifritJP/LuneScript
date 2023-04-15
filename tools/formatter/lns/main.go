// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
import LnsDepend "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
func main_expTuple0_327(tuple *LnsTuple3[LnsInt,LnsInt,bool]) (LnsInt,LnsInt,bool) {
    return tuple.Val1,tuple.Val2,tuple.Val3
}
func main_expTuple0_412(tuple *LnsTuple3[LnsInt,LnsInt,bool]) (LnsInt,LnsInt,bool) {
    return tuple.Val1,tuple.Val2,tuple.Val3
}
// for 238: ExpCast
func conv2Form0_1065( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 156
func main_convExp0_673(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 107
func main_convExp0_332(arg1 []LnsAny) (LnsInt, LnsInt, bool) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(LnsInt), Lns_getFromMulti( arg1, 2 ).(bool)
}
// for 129
func main_convExp0_417(arg1 []LnsAny) (LnsInt, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 218
func main_convExp0_916(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 219
func main_convExp0_929(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 13: decl @lns.@main.createTokenizer
func main_createTokenizer_0_(_env *LnsEnv) LnsAny {
    var lnsCode string
    lnsCode = "{\n   if val.val2\n {\n   }\n\n\n\n\n   else {\n}\n\n\n}\nlet val = [ \n  a\n\n];\nmatch hoge {\n\n   case .a  {\n\n}\n}\n\nswitch hoge {\n    case .a {\n}\n    case .b,\n.c ___LNS___\n{ \n\n}\n"
    return &LnsTypes.Types_TokenizerSrc__LnsCode{lnsCode, "test", nil}
}

// 51: decl @lns.@main.createTokenizerFromFile
func main_createTokenizerFromFile_1_(_env *LnsEnv, path string) LnsAny {
    return &LnsTypes.Types_TokenizerSrc__LnsPath{nil, path, "test", nil}
}

// 55: decl @lns.@main.getIndent
func main_getIndent_2_(_env *LnsEnv, tokenizerSrc LnsAny,targetLineNo LnsInt,endLineNo LnsInt) *Formatter_Line2Indent {
    var tokenizer *LnsTokenizer.Tokenizer_StreamTokenizer
    tokenizer = LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, tokenizerSrc, true, nil, nil)
    var codeTokenizer Code_CodeTokenizerIF
    codeTokenizer = NewCode_CodeTokenizer(_env, &tokenizer.Tokenizer_Tokenizer).FP
    var codePicker *Formatter_CodePicker
    codePicker = Formatter_CodePicker_createFrom(_env, codeTokenizer, targetLineNo, endLineNo)
    var parser *Formatter_SimpleParser
    parser = NewFormatter_SimpleParser(_env, codePicker.FP, targetLineNo, endLineNo)
    return parser.FP.Analyze(_env)
}

// 71: decl @lns.@main.getIndentWithCode
func Main_getIndentWithCode(_env *LnsEnv, lnsCode string,targetLineNo LnsInt,endLineNo LnsInt) string {
    var line2indent *Formatter_Line2Indent
    line2indent = main_getIndent_2_(_env, &LnsTypes.Types_TokenizerSrc__LnsCode{lnsCode, "code", nil}, targetLineNo, endLineNo)
    var stream *LnsUtil.Util_memStream
    stream = LnsUtil.NewUtil_memStream(_env)
    line2indent.FP.OutputResult(_env, stream.FP)
    return stream.FP.Get_txt(_env)
}

// 90: decl @lns.@main.format
func main_format_5_(_env *LnsEnv, option *main_Option) bool {
    var tokenizerSrc LnsAny
    if _switch0 := option.FP.Get_inpath(_env); _switch0 == "" {
        tokenizerSrc = main_createTokenizer_0_(_env)
    } else if _switch0 == "@-" {
        var lnsCode string
        lnsCode = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_io_stdin.Read(_env, "*a")) ||
            _env.SetStackVal( "") ).(string)
        tokenizerSrc = &LnsTypes.Types_TokenizerSrc__LnsCode{lnsCode, "test", nil}
    } else {
        tokenizerSrc = main_createTokenizerFromFile_1_(_env, option.FP.Get_inpath(_env))
    }
    {
        _lineNo := option.FP.Get_lineNo(_env)
        if !Lns_IsNil( _lineNo ) {
            lineNo := _lineNo.(*LnsTuple3[LnsInt,LnsInt,bool])
            var targetLineNo LnsInt
            var endLineNo LnsInt
            var useEbnf bool
            targetLineNo,endLineNo,useEbnf = main_expTuple0_327(lineNo)
            if Lns_op_not(useEbnf){
                var line2indent *Formatter_Line2Indent
                line2indent = main_getIndent_2_(_env, tokenizerSrc, targetLineNo, endLineNo)
                line2indent.FP.OutputResult(_env, Lns_io_stdout)
                return true
            }
        }
    }
    var tokenizer *LnsTokenizer.Tokenizer_StreamTokenizer
    tokenizer = LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, tokenizerSrc, true, nil, nil)
    var codeTokenizer Code_CodeTokenizerIF
    codeTokenizer = NewCode_CodeTokenizer(_env, &tokenizer.Tokenizer_Tokenizer).FP
    var hookIf Code_ParseCodeHookIF
    {
        _lineNo := option.FP.Get_lineNo(_env)
        if !Lns_IsNil( _lineNo ) {
            lineNo := _lineNo.(*LnsTuple3[LnsInt,LnsInt,bool])
            var targetLineNo LnsInt
            var endLineNo LnsInt
            targetLineNo,endLineNo = main_convExp0_417(Lns_2DDD(main_expTuple0_412(lineNo)))
            var hook *Formatter_ParseCodeHook
            hook = NewFormatter_ParseCodeHook(_env, codeTokenizer, targetLineNo, endLineNo)
            codeTokenizer = hook.FP
            hookIf = hook.FP
        } else {
            hookIf = NewCode_DummyParseCodeHook(_env).FP
        }
    }
    var ebnfCtrl *Parser_EbnfCtrl
    ebnfCtrl = Parser_analyze_ebnf(_env, Ebnf_EbnfTokenizer_create(_env, option.FP.Get_ebnfPath(_env)))
    if option.FP.Get_debugFlag(_env){
        ebnfCtrl.FP.Dump(_env)
    }
    switch _matchExp0 := ebnfCtrl.FP.Parse(_env, "<code>", codeTokenizer, hookIf).(type) {
    case *Code_ParseCodeRet__Eof:
        Util_log(_env, Lns_2DDD("eof"))
    case *Code_ParseCodeRet__Unmatch:
        Util_log(_env, Lns_2DDD("unmatch"))
    case *Code_ParseCodeRet__Abbr:
        var pos LnsTypes.Types_Position
        pos = codeTokenizer.GetToken(_env).Pos
        Util_log(_env, Lns_2DDD("illegal", pos.LineNo, pos.Column))
    case *Code_ParseCodeRet__Detect:
        codeCore := _matchExp0.Val1
        Util_log(_env, Lns_2DDD("outputCode"))
        {
            _fileObj := main_convExp0_673(Lns_2DDD(Lns_io_open(option.FP.Get_outpath(_env), "w")))
            if !Lns_IsNil( _fileObj ) {
                fileObj := _fileObj.(Lns_luaStream)
                var stream *LnsUtil.Util_memStream
                stream = LnsUtil.NewUtil_memStream(_env)
                Code_outputCode(_env, codeCore, NewCode_CodeGenerator(_env, stream.FP, codeTokenizer))
                fileObj.Write(_env, stream.FP.Get_txt(_env))
                fileObj.Close(_env)
                Util_log(_env, Lns_2DDD(""))
                var token *LnsTypes.Types_Token
                token = codeTokenizer.GetToken(_env)
                var pos LnsTypes.Types_Position
                pos = token.Pos
                Util_log(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d", Lns_2DDD(pos.LineNo, pos.Column)), LnsTypes.Types_TokenKind_getTxt( token.Kind)))
                return true
            }
        }
    }
    return false
}

// 172: decl @lns.@main.usage
func main_usage_6_(_env *LnsEnv, mess string) {
    Lns_print(Lns_2DDD(mess))
    Lns_print(Lns_2DDD(""))
    Lns_print(Lns_2DDD("usage: command [-log] [-ebnf path] [-prof] [-i lineno] inpath outpath"))
    _env.GetVM().OS_exit(1)
}

// 179: decl @lns.@main.__main
func Main___main(_env *LnsEnv, argList *LnsList) LnsInt {
    Lns_main_init( _env )
    var optList *LnsList
    optList = NewLnsList([]LnsAny{})
    var ebnfPath LnsAny
    ebnfPath = nil
    var index LnsInt
    index = 1
    var prof bool
    prof = false
    var debugFlag bool
    debugFlag = false
    var region LnsAny
    region = nil
    var main_getNext func(_env *LnsEnv) string
    main_getNext = func(_env *LnsEnv) string {
        if index > argList.Len(){
            main_usage_6_(_env, "illegal option")
        }
        var arg string
        arg = argList.GetAt(index).(string)
        index = index + 1
        return arg
    }
    for index <= argList.Len() {
        var arg string
        arg = main_getNext(_env)
        if _switch0 := arg; _switch0 == "-log" {
            debugFlag = true
            Util_setLog(_env, true)
        } else if _switch0 == "-prof" {
            prof = true
        } else if _switch0 == "-ebnf" {
            ebnfPath = main_getNext(_env)
        } else if _switch0 == "-i" || _switch0 == "-I" {
            var nextArg string
            nextArg = main_getNext(_env)
            var pre string
            pre = main_convExp0_916(Lns_2DDD(_env.GetVM().String_gsub(nextArg,":.*", "")))
            var post string
            post = main_convExp0_929(Lns_2DDD(_env.GetVM().String_gsub(nextArg,".*:", "")))
            var startLineNo LnsInt
            startLineNo = main___main__convInt_1_(_env, pre, _env.GetVM().String_format("%s is not number", Lns_2DDD(pre)))
            var endLineNo LnsInt
            endLineNo = main___main__convInt_1_(_env, post, _env.GetVM().String_format("%s is not number", Lns_2DDD(post)))
            region = &LnsTuple3[LnsInt,LnsInt,bool]{startLineNo, endLineNo, arg == "-I"}
        } else {
            optList.Insert(arg)
        }
    }
    if optList.Len() < 3{
        main_usage_6_(_env, "illegal argument")
    }
    var exitCode LnsInt
    exitCode = 1
    LnsDepend.Depend_profile(_env, prof, conv2Form0_1065(func(_env *LnsEnv) {
        if main_format_5_(_env, Newmain_Option(_env, optList.GetAt(2).(string), region, optList.GetAt(3).(string), ebnfPath, debugFlag)){
            exitCode = 0
        }
    }), "prof")
    return exitCode
}


// 195: decl @lns.@main.__main.convInt
func main___main__convInt_1_(_env *LnsEnv, txt string,mess string) LnsInt {
    {
        _num := Lns_tonumber(txt, nil)
        if !Lns_IsNil( _num ) {
            num := _num.(LnsReal)
            return (LnsInt)(num)
        } else {
            main_usage_6_(_env, mess)
        }
    }
// insert a dummy
    return 0
}


// declaration Class -- Option
type main_OptionMtd interface {
    Get_debugFlag(_env *LnsEnv) bool
    Get_ebnfPath(_env *LnsEnv) LnsAny
    Get_inpath(_env *LnsEnv) string
    Get_lineNo(_env *LnsEnv) LnsAny
    Get_outpath(_env *LnsEnv) string
}
type main_Option struct {
    inpath string
    lineNo LnsAny
    outpath string
    ebnfPath LnsAny
    debugFlag bool
    FP main_OptionMtd
}
func main_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*main_Option).FP
}
func main_Option_toSlice(slice []LnsAny) []*main_Option {
    ret := make([]*main_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(main_OptionDownCast).Tomain_Option()
    }
    return ret
}
type main_OptionDownCast interface {
    Tomain_Option() *main_Option
}
func main_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(main_OptionDownCast)
    if ok { return work.Tomain_Option() }
    return nil
}
func (obj *main_Option) Tomain_Option() *main_Option {
    return obj
}
func Newmain_Option(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string, arg4 LnsAny, arg5 bool) *main_Option {
    obj := &main_Option{}
    obj.FP = obj
    obj.Initmain_Option(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *main_Option) Initmain_Option(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string, arg4 LnsAny, arg5 bool) {
    self.inpath = arg1
    self.lineNo = arg2
    self.outpath = arg3
    self.ebnfPath = arg4
    self.debugFlag = arg5
}
func (self *main_Option) Get_inpath(_env *LnsEnv) string{ return self.inpath }
func (self *main_Option) Get_lineNo(_env *LnsEnv) LnsAny{ return self.lineNo }
func (self *main_Option) Get_outpath(_env *LnsEnv) string{ return self.outpath }
func (self *main_Option) Get_ebnfPath(_env *LnsEnv) LnsAny{ return self.ebnfPath }
func (self *main_Option) Get_debugFlag(_env *LnsEnv) bool{ return self.debugFlag }

func Lns_main_init(_env *LnsEnv) {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@main"
    Lns_InitMod()
    Lns_Ebnf_init(_env)
    Lns_Parser_init(_env)
    Lns_Code_init(_env)
    Lns_Util_init(_env)
    Lns_Formatter_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
    LnsDepend.Lns_Depend_init(_env)
    LnsUtil.Lns_Util_init(_env)
}
func init() {
    init_main = false
}
