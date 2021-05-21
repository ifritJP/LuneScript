// +build !noEnvArg

package runtimelns

import (
	"os"
)

type LnsForm func(_env *LnsEnv, val []LnsAny) []LnsAny

func Depend_profile(_env *LnsEnv, validTest bool, work LnsForm, path string) LnsAny {
	return Depend_profileSub(
		validTest, path,
		func() LnsAny {
			return work(_env, []LnsAny{})
		})
}

func Lns_RunMain(mainFunc func(_env *LnsEnv, args *LnsList) LnsInt) {
	args := []LnsAny{}
	for _, arg := range os.Args {
		args = append(args, arg)
	}

	os.Exit(mainFunc(Lns_GetEnv(), NewLnsList(args)))
}
