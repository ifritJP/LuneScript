//go:build !noEnvArg
// +build !noEnvArg

package runtimelns

import (
	"os"
)

type LnsForm func(_env *LnsEnv, val []LnsAny) []LnsAny

func Lns_Depend_profile(_env *LnsEnv, validTest bool, work LnsForm, path string) LnsAny {
	return depend_profileSub(
		_env, validTest, path,
		func() LnsAny {
			return work(_env, []LnsAny{})
		})
}

func Lns_RunMain(mainFunc func(_env *LnsEnv, args *LnsList) LnsInt) {
	exitRuntime(Lns_RunMainNoExit(mainFunc))
}

func Lns_RunMainNoExit(mainFunc func(_env *LnsEnv, args *LnsList) LnsInt) LnsInt {
	args := []LnsAny{}
	for _, arg := range os.Args {
		args = append(args, arg)
	}

	return mainFunc(Lns_GetEnv(), NewLnsList(args))
}
