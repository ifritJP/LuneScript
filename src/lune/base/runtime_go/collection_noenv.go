// +build noEnvArg

package runtimelns

import "sort"

type LnsComp func( val1, val2 LnsAny) bool

func (self *LnsList) Sort(_env *LnsEnv, kind int, comp LnsAny) {
	if self.lnsItemKind == LnsItemKindUnknown || self.lnsItemKind == LnsItemKindStem {
		self.lnsItemKind = kind
	}
	if Lns_IsNil(comp) {
		if self.lnsItemKind == LnsItemKindStem {
			hasInt := 0
			hasReal := 0
			hasStr := 0
			hasStem := false
			for _, val := range self.Items {
				switch val.(type) {
				case LnsInt:
					hasInt = 1
				case LnsReal:
					hasReal = 1
				case string:
					hasStr = 1
				default:
					break
				}
			}
			if !hasStem && (hasInt+hasReal+hasStr) == 1 {
				if hasInt == 1 {
					self.lnsItemKind = LnsItemKindInt
				} else if hasReal == 1 {
					self.lnsItemKind = LnsItemKindReal
				} else if hasStr == 1 {
					self.lnsItemKind = LnsItemKindStr
				}
			}
		}
		sort.Sort(self)
	} else {
        callback := comp.(LnsComp)
        sort.Slice(
            self.Items,
            func(idx1, idx2 int) bool {
                return callback(_env, self.Items[idx1], self.Items[idx2])
            })
	}
}
