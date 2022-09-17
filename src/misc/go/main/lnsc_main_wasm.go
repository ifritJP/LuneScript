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

package main

import (
	"syscall/js"

	lnsc "github.com/ifritJP/LuneScript/src/lune/base"
	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

//import bind "github.com/ifritJP/LuneScript/src/lune"

var JSNone interface{}

func lns2lua(this js.Value, args []js.Value) interface{} {

	lnsCode := args[0].String()
	path := "lnsc_frontend.lns"

	env := Lns_GetEnv()

	luaCode := ""
	Lns_LockEnvSync(env, 0, func() {
		option := lnsc.Option_createDefaultOption(env, NewLnsList([]LnsAny{path}), nil)

		luaCode = lnsc.Front_convertLnsCode2LuaCodeWithOpt(
			env, option, lnsCode, path, nil)
	})

	return luaCode
}

func exeLua(this js.Value, args []js.Value) interface{} {

	result := ""

	env := Lns_GetEnv()
	Lns_LockEnvSync(env, 0, func() {
		// luaCode :=
		luaCode := `
local writeList = { "" }
local function write( self, txt )
   table.insert( writeList, txt )
   return self
end
io.stdout = write
io.stderr = write
local __org_print = print
print = function( ... )
   local args = { ... }
   for index, val in ipairs( args ) do
      write( nil, string.format( "%s", val ) )
      if index < #args then
         write( nil, "\t" )
      end
   end
   write( nil, "\n" )
end
local function luaCode()
`
		luaCode += args[0].String()
		luaCode += `
end
luaCode()
return table.concat( writeList )`

		result = lnsc.Front_loadFromLuaTxt(env, luaCode).(string)
	})

	return result
}

func setConsoleWriter(this js.Value, args []js.Value) interface{} {

	writerObj := args[0]

	env := Lns_GetEnv()

	lnsc.Util_setConsoleOStreamWithWriter(
		env,
		func(env *LnsEnv, txt string) {
			writerObj.Invoke(txt)
		},
		func(env *LnsEnv, txt string) {
			writerObj.Invoke(txt)
		})

	return JSNone
}

func Setup(this js.Value, args []js.Value) interface{} {
	obj := map[string]interface{}{}
	obj["lns2lua"] = js.FuncOf(lns2lua)
	obj["exeLua"] = js.FuncOf(exeLua)
	obj["setConsoleWriter"] = js.FuncOf(setConsoleWriter)
	return obj
}

func main() {

	Lns_InitModOnce()
	env := Lns_GetEnv()
	lnsc.Lns_front_init(env)
	lnsc.Util_setDebugFlag(env, false)
	lnsc.Util_setConsoleOStreamWithWriter(
		env,
		func(env *LnsEnv, txt string) {
			print(txt)
		},
		func(env *LnsEnv, txt string) {
			print(txt)
		})

	//Lns_RunMainNoExit(lnsc.Front___main)

	js.Global().Set("__lnsc", js.FuncOf(Setup))
	<-make(chan bool)

}
