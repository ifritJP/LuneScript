// +build noEnvArg

package runtimelns

func Lns_Depend_init() {
}

func Depend_getFileLastModifiedTime(path string) LnsAny {
	return depend_getFileLastModifiedTime(path)
}

func Depend_getLoadedMod() *Lns_luaValue {
	return depend_getLoadedMod()
}

func Depend_getStackTrace() string {
	return ""
}

func Depend_searchpath(mod string, pathPattern string) LnsAny {
	return depend_searchpath(mod, pathPattern)
}

func Depend_existFile(path string) bool {
	return depend_existFile(path)
}

func Depend_canUseChannel() bool {
	return true
}

func Depend_canUseAsync() bool {
	return false
}

var DependLuaOnLns_runLuaOnLnsFunc func(luaCode string) (LnsAny, string) = nil

func DependLuaOnLns_runLuaOnLns(luaCode string) (LnsAny, string) {
	return dependLuaOnLns_runLuaOnLns(luaCode)
}

func Lns_DependLuaOnLns_init() {
}

func Depend_getGOPATH() LnsAny {
	return depend_getGOPATH()
}

type Depend_UpdateVer func(ver LnsInt)

func Depend_getLuaVersion() string {
	return depend_getLuaVersion()
}

func Depend_setup(callback Depend_UpdateVer) {
	depend_setup(
		func(ver LnsInt) {
			callback(ver)
		})
}

func Depend_getBindLns(mod string) LnsAny {
	return depend_getBindLns(mod)
}

func Depend_runMain(mainFunc LnsAny, argList *LnsList) LnsInt {
	return depend_runMain(mainFunc, argList)
}

func Depend_setRuntimeLog(valid bool) {
	depend_setRuntimeLog(valid)
}

func Depend_setRuntimeThreadLimit(limit int) {
	Lns_setThreadLimit(limit)
}
