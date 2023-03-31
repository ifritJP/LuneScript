// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
import LnsDepend "github.com/ifritJP/LuneScript/src/lune/base"
import LnsUtil "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
// for 125: ExpCast
func conv2Form0_630( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 66
func main_convExp0_371(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 12: decl @lns.@main.createTokenizer
func main_createTokenizer_0_(_env *LnsEnv) *LnsTokenizer.Tokenizer_Tokenizer {
    var lnsCode string
    lnsCode = "      macro _clone( map:sym ) {\n         foreach val, key in self.,,map {\n            obj.,,map[ key ] = val;\n         }\n      }\n"
    return &LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, &LnsTypes.Types_TokenizerSrc__LnsCode{lnsCode, "test", nil}, false, nil, nil).Tokenizer_Tokenizer
}

// 24: decl @lns.@main.createTokenizerFromFile
func main_createTokenizerFromFile_1_(_env *LnsEnv, path string) *LnsTokenizer.Tokenizer_Tokenizer {
    return &LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, &LnsTypes.Types_TokenizerSrc__LnsPath{nil, path, "test", nil}, false, nil, nil).Tokenizer_Tokenizer
}

// 30: decl @lns.@main.format
func main_format_2_(_env *LnsEnv, inpath string,outpath string,ebnfPath LnsAny,debugFlag bool) bool {
    var tokenizer *LnsTokenizer.Tokenizer_Tokenizer
    if inpath == ""{
        tokenizer = main_createTokenizer_0_(_env)
    } else { 
        tokenizer = main_createTokenizerFromFile_1_(_env, inpath)
    }
    var ebnfCtrl *Parser_EbnfCtrl
    ebnfCtrl = Parser_analyze_ebnf(_env, Ebnf_EbnfTokenizer_create(_env, ebnfPath))
    if debugFlag{
        ebnfCtrl.FP.Dump(_env)
    }
    var hook *Code_ParseCodeHook
    hook = NewCode_ParseCodeHook(_env)
    var codeTokenizer *Code_CodeTokenizer
    codeTokenizer = NewCode_CodeTokenizer(_env, tokenizer)
    switch _matchExp0 := ebnfCtrl.FP.Parse(_env, "<code>", codeTokenizer, hook.FP).(type) {
    case *Code_ParseCodeRet__Eof:
        Util_log(_env, Lns_2DDD("eof"))
    case *Code_ParseCodeRet__Unmatch:
        Util_log(_env, Lns_2DDD("unmatch"))
    case *Code_ParseCodeRet__Abbr:
        var pos LnsTypes.Types_Position
        pos = codeTokenizer.FP.GetToken(_env).Pos
        Util_log(_env, Lns_2DDD("illegal", pos.LineNo, pos.Column))
    case *Code_ParseCodeRet__Detect:
        codeCore := _matchExp0.Val1
        Util_log(_env, Lns_2DDD("outputCode"))
        {
            _fileObj := main_convExp0_371(Lns_2DDD(Lns_io_open(outpath, "w")))
            if !Lns_IsNil( _fileObj ) {
                fileObj := _fileObj.(Lns_luaStream)
                var stream *LnsUtil.Util_memStream
                stream = LnsUtil.NewUtil_memStream(_env)
                Code_outputCode(_env, codeCore, NewCode_CodeGenerator(_env, stream.FP))
                fileObj.Write(_env, stream.FP.Get_txt(_env))
                fileObj.Close(_env)
                Util_log(_env, Lns_2DDD(""))
                var token *LnsTypes.Types_Token
                token = codeTokenizer.FP.GetToken(_env)
                var pos LnsTypes.Types_Position
                pos = token.Pos
                Util_log(_env, Lns_2DDD(_env.GetVM().String_format("%d:%d", Lns_2DDD(pos.LineNo, pos.Column)), LnsTypes.Types_TokenKind_getTxt( token.Kind)))
                return true
            }
        }
    }
    return false
}

// 82: decl @lns.@main.usage
func main_usage_3_(_env *LnsEnv, mess string) {
    Lns_print(Lns_2DDD(mess))
    Lns_print(Lns_2DDD(""))
    Lns_print(Lns_2DDD("usage: command inpath outpath"))
    _env.GetVM().OS_exit(1)
}

// 89: decl @lns.@main.__main
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
    var main_getNext func(_env *LnsEnv) string
    main_getNext = func(_env *LnsEnv) string {
        if index > argList.Len(){
            main_usage_3_(_env, "illegal option")
        }
        var arg string
        arg = argList.GetAt(index).(string)
        index = index + 1
        return arg
    }
    for index <= argList.Len() {
        var arg string
        arg = main_getNext(_env)
        if arg == "-log"{
            debugFlag = true
            Util_setLog(_env, true)
        } else if arg == "-prof"{
            prof = true
        } else if arg == "-ebnf"{
            ebnfPath = main_getNext(_env)
        } else { 
            optList.Insert(arg)
        }
    }
    if optList.Len() < 3{
        Lns_print(Lns_2DDD("usage: command [-log] [-ebnf path] [-prof] inpath outpath"))
        return 1
    }
    var exitCode LnsInt
    exitCode = 1
    LnsDepend.Depend_profile(_env, prof, conv2Form0_630(func(_env *LnsEnv) {
        if main_format_2_(_env, optList.GetAt(2).(string), optList.GetAt(3).(string), ebnfPath, debugFlag){
            exitCode = 0
        }
    }), "prof")
    return exitCode
}



func Lns_main_init(_env *LnsEnv) {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@main"
    Lns_InitMod()
    Lns_Ebnf_init(_env)
    Lns_Parser_init(_env)
    Lns_Code_init(_env)
    Lns_Util_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
    LnsDepend.Lns_Depend_init(_env)
    LnsUtil.Lns_Util_init(_env)
}
func init() {
    init_main = false
}
