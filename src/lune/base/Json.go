// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Json bool
var Json__mod__ string
// for 43
func Json_convExp166(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 65
func Json_convExp257(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 121
func Json_convExp492(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 142
func Json_convExp573(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 147
func Json_convExp642(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 152
func Json_convExp711(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 157
func Json_convExp786(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 162
func Json_convExp855(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 167
func Json_convExp934(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 172
func Json_convExp1013(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 177
func Json_convExp1092(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 182
func Json_convExp1161(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 199
func Json_convExp1445(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 207
func Json_convExp1566(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 4: decl @lune.@base.@Json.getRawTxt
func Json_getRawTxt_1004_(_env *LnsEnv, token *Types_Token) string {
    return _env.LuaVM.String_sub(token.Txt,2, -2)
}

// 8: decl @lune.@base.@Json.getVal
func Json_getVal_1006_(_env *LnsEnv, parser *Parser_DefaultPushbackParser)(LnsAny, bool) {
    var token *Types_Token
    token = parser.FP.GetTokenNoErr(_env)
    if _switch428 := token.Kind; _switch428 == Types_TokenKind__Dlmt {
        if _switch293 := token.Txt; _switch293 == "{" {
            var _map *LnsMap
            _map = NewLnsMap( map[LnsAny]LnsAny{})
            for  {
                var key *Types_Token
                key = parser.FP.GetTokenNoErr(_env)
                if _switch140 := key.Kind; _switch140 == Types_TokenKind__Str {
                } else if _switch140 == Types_TokenKind__Dlmt {
                    if _switch131 := key.Txt; _switch131 == "}" {
                        return _map, true
                    } else if _switch131 == "," {
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
                val,ok = Json_getVal_1006_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                _map.Set(Json_getRawTxt_1004_(_env, key),val)
            }
        } else if _switch293 == "[" {
            var list *LnsList
            list = NewLnsList([]LnsAny{})
            var count LnsInt
            count = 1
            for  {
                var nextToken *Types_Token
                nextToken = parser.FP.GetTokenNoErr(_env)
                if _switch247 := nextToken.Txt; _switch247 == "]" {
                    return list, true
                } else if _switch247 == "," {
                } else {
                    parser.FP.Pushback(_env)
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1006_(_env, parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                list.Set(count,val)
                count = count + 1
                
            }
        }
        return nil, false
    } else if _switch428 == Types_TokenKind__Int {
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
    } else if _switch428 == Types_TokenKind__Real {
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
    } else if _switch428 == Types_TokenKind__Str {
        return Json_getRawTxt_1004_(_env, token), true
    } else if _switch428 == Types_TokenKind__Kywd {
        if _switch420 := token.Txt; _switch420 == "true" {
            return true, true
        } else if _switch420 == "false" {
            return false, true
        } else if _switch420 == "nil" || _switch420 == "null" {
            return nil, true
        }
        return nil, false
    }
    return nil, false
}

// 117: decl @lune.@base.@Json.fromStr
func Json_fromStr(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(_env, txt)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(_env, &NewParser_StreamParser(_env, stream.FP, "json", false, nil).Parser_Parser)
    var val LnsAny
    var ok bool
    val,ok = Json_getVal_1006_(_env, parser)
    if Lns_op_not(ok){
        return nil, parser.FP.GetLastPos(_env)
    }
    return val, nil
}

// 131: decl @lune.@base.@Json.lenMap
func _lenMap_1020_(_env *LnsEnv, _map LnsAny) LnsInt {
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
