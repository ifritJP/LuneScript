package lnsc

import (
	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

var init_DependLuaOnLns bool

var DependLuaOnLns_runLuaOnLnsFunc func(_env *LnsEnv, luaCode string) (LnsAny, string) = nil

func DependLuaOnLns_runLuaOnLns(
	_env *LnsEnv, frontAccessor *FrontInterface_FrontAccessor,
	luaCode string, baseDir LnsAny, async bool) (LnsAny, string) {
	return Lns_dependLuaOnLns_runLuaOnLns(_env, luaCode, baseDir, async)
}

func Depend_profile(_env *LnsEnv, validTest bool, work LnsForm, path string) LnsAny {
	return Lns_Depend_profile(_env, validTest, work, path)
}

func Lns_DependLuaOnLns_init(_env *LnsEnv) {
	if init_DependLuaOnLns {
		return
	}
	init_DependLuaOnLns = true
}
func Lns_Depend_init(_env *LnsEnv) {
	Lns_DependLuaOnLns_init(_env)
}
