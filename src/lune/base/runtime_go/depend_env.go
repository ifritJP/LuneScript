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

func DependLuaOnLns_runLuaOnLns(_env *LnsEnv, luaCode string) (LnsAny, string) {
	return dependLuaOnLns_runLuaOnLns(luaCode)
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
