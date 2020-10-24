/*
MIT License

Copyright (c) 2018,2020 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

// #include <stdlib.h>
import "C"

func (luaVM *Lns_luaVM) RunCode( code string ) (bool,[]LnsAny, LnsAny) {
    loaded, mess := luaVM.Load( code, nil )
    if loaded == nil {
        return false, []LnsAny{}, mess
    }
    ret := luaVM.RunLoadedfunc( loaded.(*Lns_luaValue), []LnsAny{} )
    return true, ret, nil
}

func (luaVM *Lns_luaVM) GetEmptyMap() *Lns_luaValue {
    code := "return {}"
    ok, ret, mess := luaVM.RunCode( code )
    if !ok {
        panic( mess )
    }
    return ret[0].(*Lns_luaValue)
}

func (luaVM *Lns_luaVM) GetPackagePath() string {
    code := "return package.path"
    ok, ret, mess := luaVM.RunCode( code )
    if !ok {
        panic( mess )
    }
    return ret[0].(string)
}


func (luaVM *Lns_luaVM) SortMapKeyList( mapObj *Lns_luaValue ) *Lns_luaValue {
    code := `
return function( map )
  local keys = {}
  for key in pairs( map ) do
    table.insert( keys, key )
  end
  table.sort(keys)
  return keys
end
`
    ok, ret, mess := luaVM.RunCode( code )
    if !ok {
        panic( mess )
    }
    loaded := ret[0].(*Lns_luaValue)
    vals := luaVM.RunLoadedfunc( loaded, []LnsAny{ mapObj } )
    return vals[ 0 ].(*Lns_luaValue)
}
