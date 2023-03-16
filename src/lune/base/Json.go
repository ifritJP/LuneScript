// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Json bool
var Json__mod__ string
// for 70
func Json_convExp0_197(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 94
func Json_convExp0_294(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 151
func Json_convExp0_512(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// 31: decl @lune.@base.@Json.getRawTxt
func Json_getRawTxt_0_(_env *LnsEnv, token *Types_Token) string {
    return _env.GetVM().String_sub(token.Txt,2, -2)
}

// 35: decl @lune.@base.@Json.getVal
func Json_getVal_1_(_env *LnsEnv, tokenizer *Parser_DefaultPushbackTokenizer)(LnsAny, bool) {
    var token *Types_Token
    token = tokenizer.FP.GetTokenNoErr(_env, nil)
    if _switch5 := token.Kind; _switch5 == Types_TokenKind__Dlmt {
        if _switch3 := token.Txt; _switch3 == "{" {
            var _map *LnsMap
            _map = NewLnsMap( map[LnsAny]LnsAny{})
            for  {
                var key *Types_Token
                key = tokenizer.FP.GetTokenNoErr(_env, nil)
                if _switch1 := key.Kind; _switch1 == Types_TokenKind__Str {
                } else if _switch1 == Types_TokenKind__Dlmt {
                    if _switch0 := key.Txt; _switch0 == "}" {
                        return _map, true
                    } else if _switch0 == "," {
                        key = tokenizer.FP.GetTokenNoErr(_env, nil)
                        if key.Kind != Types_TokenKind__Str{
                            return nil, false
                        }
                    } else {
                        return nil, false
                    }
                } else {
                    return nil, false
                }
                if tokenizer.FP.GetTokenNoErr(_env, nil).Txt != ":"{
                    return nil, false
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1_(_env, tokenizer)
                if Lns_op_not(ok){
                    return nil, false
                }
                if val != nil{
                    val_38 := val
                    _map.Set(Json_getRawTxt_0_(_env, key),val_38)
                }
            }
        } else if _switch3 == "[" {
            var list *LnsList
            list = NewLnsList([]LnsAny{})
            var count LnsInt
            count = 1
            for  {
                var nextToken *Types_Token
                nextToken = tokenizer.FP.GetTokenNoErr(_env, nil)
                if _switch2 := nextToken.Txt; _switch2 == "]" {
                    return list, true
                } else if _switch2 == "," {
                } else {
                    tokenizer.FP.Pushback(_env)
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1_(_env, tokenizer)
                if Lns_op_not(ok){
                    return nil, false
                }
                list.Set(count,val)
                count = count + 1
            }
        }
        return nil, false
    } else if _switch5 == Types_TokenKind__Int {
        var num LnsReal
        
        {
            _num := Lns_tonumber(token.Txt, nil)
            if _num == nil{
                return nil, false
            } else {
                num = _num.(LnsReal)
            }
        }
        return (LnsInt)(num), true
    } else if _switch5 == Types_TokenKind__Real {
        var num LnsReal
        
        {
            _num := Lns_tonumber(token.Txt, nil)
            if _num == nil{
                return nil, false
            } else {
                num = _num.(LnsReal)
            }
        }
        return num, true
    } else if _switch5 == Types_TokenKind__Str {
        return Json_getRawTxt_0_(_env, token), true
    } else if _switch5 == Types_TokenKind__Kywd {
        if _switch4 := token.Txt; _switch4 == "true" {
            return true, true
        } else if _switch4 == "false" {
            return false, true
        } else if _switch4 == "nil" || _switch4 == "null" {
            return nil, true
        }
        return nil, false
    }
    return nil, false
}

// 146: decl @lune.@base.@Json.fromStr
func Json_fromStr(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var tokenizer *Parser_DefaultPushbackTokenizer
    tokenizer = NewParser_DefaultPushbackTokenizer(_env, &Parser_StreamTokenizer_create(_env, &Types_TokenizerSrc__LnsCode{txt, "json", nil}, false, nil, nil).Parser_Tokenizer)
    var val LnsAny
    var ok bool
    val,ok = Json_getVal_1_(_env, tokenizer)
    if Lns_op_not(ok){
        return nil, tokenizer.FP.GetLastPos(_env)
    }
    return val, nil
}

func Lns_Json_init(_env *LnsEnv) {
    if init_Json { return }
    init_Json = true
    Json__mod__ = "@lune.@base.@Json"
    Lns_InitMod()
    Lns_Parser_init(_env)
    Lns_Types_init(_env)
}
func init() {
    init_Json = false
}
