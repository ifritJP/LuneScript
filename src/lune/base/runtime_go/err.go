/*
MIT License

Copyright (c) 2018,2020 ifritJP

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

type G__Er interface {
	Get_txt(_env *LnsEnv) string
}

func Lns_cast2G__Er(obj LnsAny) LnsAny {
	if _, ok := obj.(G__Er); ok {
		return obj
	}
	return nil
}

// 2: decl @miniGo.LnsErr.get_txt
func (self *LnsErr) Get_txt(_env *LnsEnv) string {
	return self.txt
}

// declaration Class -- LnsErr
type LnsErrMtd interface {
	Get_txt(_env *LnsEnv) string
}
type LnsErr struct {
	FP  LnsErrMtd
	txt string
}

func LnsErr2Stem(obj LnsAny) LnsAny {
	if obj == nil {
		return nil
	}
	return obj.(*LnsErr).FP
}

type LnsErrDownCast interface {
	ToLnsErr() *LnsErr
}

func LnsErrDownCastF(multi ...LnsAny) LnsAny {
	if len(multi) == 0 {
		return nil
	}
	obj := multi[0]
	if ddd, ok := multi[0].([]LnsAny); ok {
		obj = ddd[0]
	}
	work, ok := obj.(LnsErrDownCast)
	if ok {
		return work.ToLnsErr()
	}
	return nil
}
func (obj *LnsErr) ToLnsErr() *LnsErr {
	return obj
}
func NewLnsErr(_env *LnsEnv, txt string) *LnsErr {
	obj := &LnsErr{}
	obj.FP = obj
	obj.InitLnsErr(_env, txt)
	return obj
}
func (self *LnsErr) InitLnsErr(_env *LnsEnv, txt string) {
	self.txt = txt
}

func LnsErr_create(_env *LnsEnv, txt string) G__Er {
	return NewLnsErr(_env, txt).FP
}

type G__Ret = LnsAny
type G__Ret__Err struct {
	Val1 LnsAny
}

func (self *G__Ret__Err) GetTxt() string {
	return "Result.Err"
}

type G__Ret__Ok struct {
	Val1 LnsAny
}

func (self *G__Ret__Ok) GetTxt() string {
	return "Result.Ok"
}
