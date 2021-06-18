// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Json bool
var Json__mod__ string
// for 145
func Json_convExp55(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 150
func Json_convExp124(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 155
func Json_convExp193(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 160
func Json_convExp268(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 165
func Json_convExp337(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 170
func Json_convExp416(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 175
func Json_convExp495(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 180
func Json_convExp574(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 185
func Json_convExp643(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 202
func Json_convExp932(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 210
func Json_convExp1053(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 45
func Json_convExp1321(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 67
func Json_convExp1412(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 124
func Json_convExp1630(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// 6: decl @lune.@base.@Json.getRawTxt
func Json_getRawTxt_0_(_env *LnsEnv, token *Types_Token) string {
    return _env.GetVM().String_sub(token.Txt,2, -2)
}

// 10: decl @lune.@base.@Json.getVal
func Json_getVal_1_(_env *LnsEnv, parser *Parser_DefaultPushbackParser)(LnsAny, bool) {
    var token *Types_Token
    token = parser.FP.GetTokenNoErr(_env)
    if _switch1 := token.Kind; _switch1 == Types_TokenKind__Dlmt {
        if _switch2 := token.Txt; _switch2 == "{" {
            var _map *LnsMap
            _map = NewLnsMap( map[LnsAny]LnsAny{})
            for  {
                var key *Types_Token
                key = parser.FP.GetTokenNoErr(_env)
                if _switch3 := key.Kind; _switch3 == Types_TokenKind__Str {
                } else if _switch3 == Types_TokenKind__Dlmt {
                    if _switch4 := key.Txt; _switch4 == "}" {
                        return _map, true
                    } else if _switch4 == "," {
                        key = parser.FP.GetTokenNoErr(_env)
                        if key.Kind != Types_TokenKind__Str{
                            return nil, false
                        }
                    } else {
                        return nil, false
                    }
                } else {
                    return nil, false
                }
                if parser.FP.GetTokenNoErr(_env).Txt != ":"{
                    return nil, false
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                _map.Set(Json_getRawTxt_0_(_env, key),val)
            }
        } else if _switch2 == "[" {
            var list *LnsList
            list = NewLnsList([]LnsAny{})
            var count LnsInt
            count = 1
            for  {
                var nextToken *Types_Token
                nextToken = parser.FP.GetTokenNoErr(_env)
                if _switch5 := nextToken.Txt; _switch5 == "]" {
                    return list, true
                } else if _switch5 == "," {
                } else {
                    parser.FP.Pushback(_env)
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                list.Set(count,val)
                count = count + 1
            }
        }
        return nil, false
    } else if _switch1 == Types_TokenKind__Int {
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
    } else if _switch1 == Types_TokenKind__Real {
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
    } else if _switch1 == Types_TokenKind__Str {
        return Json_getRawTxt_0_(_env, token), true
    } else if _switch1 == Types_TokenKind__Kywd {
        if _switch6 := token.Txt; _switch6 == "true" {
            return true, true
        } else if _switch6 == "false" {
            return false, true
        } else if _switch6 == "nil" || _switch6 == "null" {
            return nil, true
        }
        return nil, false
    }
    return nil, false
}

// 119: decl @lune.@base.@Json.fromStr
func Json_fromStr(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(_env, &Parser_StreamParser_create(_env, &Types_ParserSrc__LnsCode{txt, "json", nil}, false, nil, nil).Parser_Parser)
    var val LnsAny
    var ok bool
    val,ok = Json_getVal_1_(_env, parser)
    if Lns_op_not(ok){
        return nil, parser.FP.GetLastPos(_env)
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
