// +build addEnvArg

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

//import "runtime"

func Lns_Str_init(_env *LnsEnv) {
	lns_Str_init()
}

func Str_getLineList(_env *LnsEnv, txt string) *LnsList {
	return str_getLineList(txt)
}

func Str_startsWith(_env *LnsEnv, txt, ptn string) bool {
	return str_startsWith(txt, ptn)
}

func Str_endsWith(_env *LnsEnv, txt, ptn string) bool {
	return str_endsWith(txt, ptn)
}

func Str_isValidStrBuilder(_env *LnsEnv) bool {
	return true
}

// ==== Builder
type Str_BuilderMtd interface {
	Get_txt(_env *LnsEnv) string
	Add(_env *LnsEnv, arg1 string)
	Len(_env *LnsEnv) LnsInt
	Clear(_env *LnsEnv)
}

func Str_Builder2Stem(_env *LnsEnv, obj LnsAny) LnsAny {
	return str_Builder2Stem(obj)
}

func (obj *Str_Builder) Str_ToBuilder(_env *LnsEnv) *Str_Builder {
	return obj.str_ToBuilder()
}

func NewStr_Builder(_env *LnsEnv) *Str_Builder {
	return newStr_Builder()
}
func (self *Str_Builder) InitStr_Builder(_env *LnsEnv) {
	self.initStr_Builder()
}

func (self *Str_Builder) Get_txt(_env *LnsEnv) string {
	return self.get_txt()
}

func (self *Str_Builder) Add(_env *LnsEnv, val string) {
	self.add(val)
}

func (self *Str_Builder) Len(_env *LnsEnv) LnsInt {
	return self.len()
}

func (self *Str_Builder) Clear(_env *LnsEnv) {
	self.clear()
}
