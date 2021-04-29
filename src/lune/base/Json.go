// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Json bool
var Json__mod__ string
// 4: decl @lune.@base.@Json.getRawTxt
func Json_getRawTxt_1007_(token *Types_Token) string {
    return Lns_getVM().String_sub(token.Txt,2, -2)
}

// 8: decl @lune.@base.@Json.getVal
func Json_getVal_1010_(parser *Parser_DefaultPushbackParser)(LnsAny, bool) {
    var token *Types_Token
    token = parser.FP.GetTokenNoErr()
    if _switch428 := token.Kind; _switch428 == Types_TokenKind__Dlmt {
        if _switch293 := token.Txt; _switch293 == "{" {
            var _map *LnsMap
            _map = NewLnsMap( map[LnsAny]LnsAny{})
            for  {
                var key *Types_Token
                key = parser.FP.GetTokenNoErr()
                if _switch140 := key.Kind; _switch140 == Types_TokenKind__Str {
                } else if _switch140 == Types_TokenKind__Dlmt {
                    if _switch131 := key.Txt; _switch131 == "}" {
                        return _map, true
                    } else if _switch131 == "," {
                        key = parser.FP.GetTokenNoErr()
                        
                        if key.Kind != Types_TokenKind__Str{
                            return nil, false
                        }
                    } else {
                        return nil, false
                    }
                } else {
                    return nil, false
                }
                if parser.FP.GetTokenNoErr().Txt != ":"{
                    return nil, false
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1010_(parser)
                if Lns_op_not(ok){
                    return nil, false
                }
                _map.Set(Json_getRawTxt_1007_(key),val)
            }
        } else if _switch293 == "[" {
            var list *LnsList
            list = NewLnsList([]LnsAny{})
            var count LnsInt
            count = 1
            for  {
                var nextToken *Types_Token
                nextToken = parser.FP.GetTokenNoErr()
                if _switch247 := nextToken.Txt; _switch247 == "]" {
                    return list, true
                } else if _switch247 == "," {
                } else {
                    parser.FP.Pushback()
                }
                var val LnsAny
                var ok bool
                val,ok = Json_getVal_1010_(parser)
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
        return Json_getRawTxt_1007_(token), true
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
func Json_fromStr(txt string)(LnsAny, LnsAny) {
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(txt)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(&NewParser_StreamParser(stream.FP, "json", false, nil).Parser_Parser)
    var val LnsAny
    var ok bool
    val,ok = Json_getVal_1010_(parser)
    if Lns_op_not(ok){
        return nil, parser.FP.GetLastPos()
    }
    return val, nil
}

// 131: decl @lune.@base.@Json.lenMap
func _lenMap_1028_(_map LnsAny) LnsInt {
    if _map != nil{
        map_339 := _map
        var count LnsInt
        count = 0
        for range( map_339.(*LnsMap).Items ) {
            count = count + 1
            
        }
        return count
    }
    return -1
}

func Lns_Json_init() {
    if init_Json { return }
    init_Json = true
    Json__mod__ = "@lune.@base.@Json"
    Lns_InitMod()
    Lns_Parser_init()
    Lns_Types_init()
}
func init() {
    init_Json = false
}
