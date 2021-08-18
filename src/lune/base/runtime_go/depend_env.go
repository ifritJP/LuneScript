// +build !noEnvArg

package runtimelns

func Lns_Depend_init(_env *LnsEnv) {
}

func Depend_getFileLastModifiedTime(_env *LnsEnv, path string) LnsAny {
	return depend_getFileLastModifiedTime(path)
}

func Depend_getLoadedMod(_env *LnsEnv) *Lns_luaValue {
	return depend_getLoadedMod()
}

func Depend_getStackTrace(_env *LnsEnv) string {
	return ""
}

func Depend_searchpath(_env *LnsEnv, mod string, pathPattern string) LnsAny {
	return depend_searchpath(mod, pathPattern)
}

func Depend_existFile(_env *LnsEnv, path string) bool {
	return depend_existFile(path)
}

func Depend_canUseChannel(_env *LnsEnv) bool {
	return true
}

func Depend_canUseAsync(_env *LnsEnv) bool {
	return false
}

var DependLuaOnLns_runLuaOnLnsFunc func(_env *LnsEnv, luaCode string) (LnsAny, string) = nil

func DependLuaOnLns_runLuaOnLns(
	_env *LnsEnv, luaCode string, baseDir LnsAny, async bool) (LnsAny, string) {
	return dependLuaOnLns_runLuaOnLns(_env, luaCode, baseDir, async)
}

func Lns_DependLuaOnLns_init(_env *LnsEnv) {
}

func Depend_getGOPATH(_env *LnsEnv) LnsAny {
	return depend_getGOPATH()
}

type Depend_UpdateVer func(_env *LnsEnv, ver LnsInt)

func Depend_getLuaVersion(_env *LnsEnv) string {
	return depend_getLuaVersion()
}

func Depend_setup(_env *LnsEnv, callback Depend_UpdateVer) {
	depend_setup(
		func(ver LnsInt) {
			callback(nil, ver)
		})
}

func Depend_getBindLns(_env *LnsEnv, mod string) LnsAny {
	return depend_getBindLns(mod)
}

func Depend_runMain(_env *LnsEnv, mainFunc LnsAny, argList *LnsList) LnsInt {
	return depend_runMain(mainFunc, argList)
}

func Depend_setupShebang(_env *LnsEnv) {
	_env.GetVM().RunCode(" _lune = _lune or {}; _lune._shebang = true")
}

func Depend_setRuntimeLog(_env *LnsEnv, valid bool) {
	depend_setRuntimeLog(valid)
}

func Depend_setRuntimeThreadLimit(_env *LnsEnv, limit int) {
	Lns_setThreadLimit(limit)
}

func Depend_setRunnerLog(_env *LnsEnv, valid bool) {
	LnsStartRunnerLog(_env, valid)
}
func Depend_dumpRunnerLog(_env *LnsEnv, stream Lns_oStream) {
	lns_threadMgrInfo.dumpEventLog(func(txt string) {
		stream.Write(_env, txt)
	})
}
