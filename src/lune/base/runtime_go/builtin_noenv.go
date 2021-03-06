// +build noEnvArg

package runtimelns

import (
	"os"
)

type LnsForm func(val []LnsAny) []LnsAny

func Depend_profile(validTest bool, work LnsForm, path string) LnsAny {
	return Depend_profileSub(
		Lns_GetEnv(),
		validTest, path,
		func() LnsAny {
			return work([]LnsAny{})
		})
}

func Lns_RunMain(mainFunc func(args *LnsList) LnsInt) {
	args := []LnsAny{}
	for _, arg := range os.Args {
		args = append(args, arg)
	}

	exitRuntime(mainFunc(NewLnsList(args)))
}
