// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTokenizer "github.com/ifritJP/LuneScript/src/lune/base"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_main bool
var main__mod__ string
var main_ebnfCtrl *Parser_EbnfCtrl
var main_codeTokenizer *LnsTokenizer.Tokenizer_Tokenizer

// 25: decl @lns.@main.createTokenizerFromFile
func main_createTokenizerFromFile_1_(_env *LnsEnv, path string) *LnsTokenizer.Tokenizer_Tokenizer {
    return &LnsTokenizer.Tokenizer_StreamTokenizer_create(_env, &LnsTypes.Types_TokenizerSrc__LnsPath{nil, path, "test", nil}, false, nil, nil).Tokenizer_Tokenizer
}

func Lns_main_init(_env *LnsEnv) {
    if init_main { return }
    init_main = true
    main__mod__ = "@lns.@main"
    Lns_InitMod()
    Lns_Ebnf_init(_env)
    Lns_Parser_init(_env)
    Lns_Code_init(_env)
    LnsTokenizer.Lns_Tokenizer_init(_env)
    LnsTypes.Lns_Types_init(_env)
    main_ebnfCtrl = Parser_analyze_ebnf(_env, Ebnf_EbnfTokenizer_create(_env))
    main_ebnfCtrl.FP.Dump(_env)
    main_codeTokenizer = main_createTokenizerFromFile_1_(_env, "../../src/lune/base/front.lns")
    main_ebnfCtrl.FP.Parse(_env, "<code>", NewCode_CodeTokenizer(_env, main_codeTokenizer))
}
func Main___main( _env *LnsEnv, args *LnsList ) LnsInt {
Lns_main_init( _env )
return 0
}
func init() {
    init_main = false
}
