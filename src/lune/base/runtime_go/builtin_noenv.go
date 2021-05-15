// +build !addEnvArg

package runtimelns

import (
	"os"
)

type LnsForm func(val []LnsAny) []LnsAny

func Depend_profile(validTest bool, work LnsForm, path string) LnsAny {
	return Depend_profileSub(
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

	os.Exit(mainFunc(NewLnsList(args)))
}
