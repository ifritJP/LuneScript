// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Json bool
var Json__mod__ string
// for 45
func Json_convExp168(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 67
func Json_convExp259(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 123
func Json_convExp494(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 144
func Json_convExp575(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 149
func Json_convExp644(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 154
func Json_convExp713(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 159
func Json_convExp788(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 164
func Json_convExp857(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 169
func Json_convExp936(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 174
func Json_convExp1015(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 179
func Json_convExp1094(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 184
func Json_convExp1163(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 201
func Json_convExp1447(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 209
func Json_convExp1568(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 6: decl @lune.@base.@Json.getRawTxt
func Json_getRawTxt_1007_(_env *LnsEnv, token *Types_Token) string {
    return _env.LuaVM.String_sub(token.Txt,2, -2)
}

// 10: decl @lune.@base.@Json.getVal
func Json_getVal_1011_(_env *LnsEnv, parser *Parser_DefaultPushbackParser)(LnsAny, bool) {
    var token *Types_Token
    token = parser.FP.GetTokenNoErr(_env)
    if _switch430 := token.Kind; _switch430 == Types_TokenKind__Dlmt {
        if _switch295 := token.Txt; _switch295 == "{" {
            var _map *LnsMap
            _map = NewLnsMap( map[LnsAny]LnsAny{})
            for  {
                var key *Types_Token
                key = parser.FP.GetTokenNoErr(_env)
                if _switch142 := key.Kind; _switch142 == Types_TokenKind__Str {
                } else if _switch142 == Types_TokenKind__Dlmt {
                    if _switch133 := key.Txt; _switch133 == "}" {
                        return _map, true
                    } else if _switch133 == "," {
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
                val,ok = Json_getVal_1011_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                _map.Set(Json_getRawTxt_1007_(_env, key),val)
            }
        } else if _switch295 == "[" {
            var list *LnsList
            list = NewLnsList([]LnsAny{})
            var count LnsInt
            count = 1
            for  {
                var nextToken *Types_Token
                nextToken = parser.FP.GetTokenNoErr(_env)
                if _switch249 := nextToken.Txt; _switch249 == "]" {
                    return list, true
                } else if _switch249 == "," {
                } else {
                    parser.FP.Pushback(_env)
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1011_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                list.Set(count,val)
                count = count + 1
                
            }
        }
        return nil, false
    } else if _switch430 == Types_TokenKind__Int {
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
    } else if _switch430 == Types_TokenKind__Real {
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
    } else if _switch430 == Types_TokenKind__Str {
        return Json_getRawTxt_1007_(_env, token), true
    } else if _switch430 == Types_TokenKind__Kywd {
        if _switch422 := token.Txt; _switch422 == "true" {
            return true, true
        } else if _switch422 == "false" {
            return false, true
        } else if _switch422 == "nil" || _switch422 == "null" {
            return nil, true
        }
        return nil, false
    }
    return nil, false
}

// 119: decl @lune.@base.@Json.fromStr
func Json_fromStr(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(_env, txt)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(_env, &NewParser_StreamParser(_env, stream.FP, "json", false, nil).Parser_Parser)
    var val LnsAny
    var ok bool
    val,ok = Json_getVal_1011_(_env, parser)
    if Lns_op_not(ok){
        return nil, parser.FP.GetLastPos(_env)
    }
    return val, nil
}

// 133: decl @lune.@base.@Json.lenMap
func _lenMap_1091_(_env *LnsEnv, _map LnsAny) LnsInt {
    if _map != nil{
        map_73 := _map
        var count LnsInt
        count = 0
        for range( map_73.(*LnsMap).Items ) {
            count = count + 1
            
        }
        return count
    }
    return -1
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
