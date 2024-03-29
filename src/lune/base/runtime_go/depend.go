/*
MIT License

Copyright (c) 2018 ifritJP

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

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

var lns_traceOn = false

// func Lns_Depend_init() {
// }

/*
*
path の最終更新日時を取得する。

@param path ファイルパス
@return 1970/1/1 0:0:0 からの秒数。 取得失敗した場合は nil。
*/
func depend_getFileLastModifiedTime(path string) LnsAny {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil
	}
	return LnsReal(fileinfo.ModTime().Unix())
}

func depend_getLoadedMod() *Lns_luaValue {
	return Lns_getVM().GetEmptyMap()
}

func printMemInfo(mess string) {
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	fmt.Printf("----- %s\n", mess)
	fmt.Printf("Alloc: %d MB\n", ms.Alloc/1024/1024)
	fmt.Printf("HeapAlloc: %d MB\n", ms.HeapAlloc/1024/1024)
	fmt.Printf("TotalAlloc: %d MB\n", ms.TotalAlloc/1024/1024)
	fmt.Printf("Sys: %d MB\n", ms.Sys/1024/1024)
	fmt.Printf("HeapObjects: %d\n", ms.HeapObjects)
	fmt.Printf("NumGC: %d\n", ms.NumGC)
}

func depend_profileSub(
	_env *LnsEnv, validTest bool, path string, work func() LnsAny) LnsAny {
	if validTest {
		{
			// start cpu profile
			profWriter, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			if err := pprof.StartCPUProfile(profWriter); err != nil {
				panic(err)
			}
			defer pprof.StopCPUProfile()
		}

		{
			LnsStartRunnerLog(_env, true)
			profRunnerWriter, err := os.Create(path + "runner")
			if err != nil {
				panic(err)
			}
			defer profRunnerWriter.Close()
			bufWriter := bufio.NewWriter(profRunnerWriter)
			defer bufWriter.Flush()

			defer lns_threadMgrInfo.dumpEventLog(func(txt string) {
				bufWriter.Write([]byte(txt))
			})
		}

		if lns_traceOn {
			// trace
			proftrace, err := os.Create(path + "trace")
			if err != nil {
				panic(err)
			}
			if err := trace.Start(proftrace); err != nil {
				panic(fmt.Sprintf("failed to start trace: %v", err))
			}
			defer trace.Stop()
		}

		// mem
		printMemInfo("start")
	}

	ret := work()

	if validTest {
		// write mem profile
		prof, err := os.Create(path + "mem")
		if err != nil {
			panic(err)
		}
		defer prof.Close()
		printMemInfo("before gc")
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(prof); err != nil {
			panic(err)
		}
		printMemInfo("end")
	}

	return ret
}

// func Depend_getStackTrace() string {
// 	return ""
// }

func depend_searchpath(mod string, pathPattern string) LnsAny {
	return Lns_getVM().Package_searchpath(mod, pathPattern)
}

func depend_existFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// func Depend_getLuaVersion() string {
//     return "5.3"
// }
// type Depend_UpdateVer func ( ver LnsInt )
// func Depend_setup( callback Depend_UpdateVer) {
//     callback( 53 );
// }

// func Depend_canUseChannel() bool {
// 	return true
// }
// func Depend_canUseAsync() bool {
// 	return false
// }

//var dependLuaOnLns_runLuaOnLnsFunc func(luaCode string) (LnsAny, string) = nil

func dependLuaOnLns_runLuaOnLns(
	_env *LnsEnv, luaCode string, baseDir LnsAny, async bool) (LnsAny, string) {

	setBindListStr := ""
	for key, _ := range lnsSrcMap {
		setBindListStr = fmt.Sprintf(
			"%s\ntable.insert( bindModuleList, '%s' )", setBindListStr, key)
	}

	txt := fmt.Sprintf(`
local DependLuaOnLns = require( 'lune.base.DependLuaOnLns' )
local Depend = require( 'lune.base.Depend' )

local bindModuleList = {}
%s

local function newFront()
   local Front = require( 'lune.base.front' )
   local Option = require( 'lune.base.Option' )
   local option = Option.createDefaultOption( {"dummy.lns"}, nil )
   return Front.Front._new(option, bindModuleList)
end

-- 高速性最重視のため、 __luneScript に直接アクセスする
local WrapFront = {}
function WrapFront:setupFront()
   if self.readyFront then
      return
   end
   -- require( 'lune.base.front' ) 内の loadModule 処理で
   -- 再度 WrapFront:loadModule() がコールされて無限ループなるので、
   -- __luneScript を nil でクリアしておく
   __luneScript = nil
   self.readyFront = true

   __luneScript = newFront()
end
function WrapFront:loadModule( mod )
   self:setupFront()
   return __luneScript:loadModule( mod )
end
function WrapFront:loadMeta( importModuleInfo, mod, orgMod, baseDir )
   self:setupFront()
   return __luneScript:loadMeta( importModuleInfo, mod, orgMod, baseDir )
end
function WrapFront:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )
   self:setupFront()
   return __luneScript:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )
end
function WrapFront:searchModule( mod, baseDir )
   self:setupFront()
   return __luneScript:searchModule( mod, baseDir )
end
function WrapFront:error( message )
   self:setupFront()
   return __luneScript:error( message )
end
function WrapFront.setmeta( obj )
  setmetatable( obj, { __index = WrapFront } )
end
function WrapFront.new(  )
   local obj = {}
   WrapFront.setmeta( obj )
   return obj
end

local wrapFront = WrapFront.new()
__luneScript = wrapFront

_lnsLoad = function( name, code )
   __luneScript = nil
   local frontInterface = require( 'lune.base.frontInterface' )
   __luneScript = newFront()

   local importModuleInfo = frontInterface.ImportModuleInfo.new();
   return frontInterface.loadFromLnsTxt( importModuleInfo, name, code )
end

local txt=[===[
%s
]===]

return DependLuaOnLns.runLuaOnLns( wrapFront, txt )
    `, setBindListStr, luaCode)

	var loaded LnsAny
	var err LnsAny
	var ret []LnsAny

	if async {
		luaVM := _env.GetVM()
		loaded, err = luaVM.Load(txt, nil)
		if loaded != nil {
			ret = luaVM.RunLoadedfunc(loaded.(*Lns_luaValue), []LnsAny{})
		}
	} else {
		Lns_LockEnvSync(_env, -254, func() {
			luaVM := _env.GetVM()
			loaded, err = luaVM.Load(txt, nil)
			if loaded != nil {
				ret = luaVM.RunLoadedfunc(loaded.(*Lns_luaValue), []LnsAny{})
			}
		})
	}
	if loaded != nil {
		return ret[0], ""
	}
	if err != nil {
		return nil, err.(string)
	}
	return nil, ""
}
func lns_DependLuaOnLns_init() {
}

func depend_runMain(mainFunc LnsAny, argList *LnsList) LnsInt {
	if !Lns_IsNil(mainFunc) {
		luaVM := Lns_getVM()
		ret := luaVM.RunLoadedfunc(mainFunc.(*Lns_luaValue), []LnsAny{argList})
		return ret[0].(LnsInt)
	}
	return -1
}

func depend_getGOPATH() LnsAny {
	val, exist := os.LookupEnv("GOPATH")
	if !exist {
		val, exist = os.LookupEnv("HOME")
		if !exist {
			return nil
		}
		return path.Join(val, "go")
	}
	return val
}

var validRuntimeLog = false

func depend_setRuntimeLog(valid bool) {
	validRuntimeLog = valid
}
