// +build !addEnvArg

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

func Lns_Str_init() {
	lns_Str_init()
}

func Str_getLineList(txt string) *LnsList {
	return str_getLineList(txt)
}

func Str_startsWith(txt, ptn string) bool {
	return str_startsWith(txt, ptn)
}

func Str_endsWith(txt, ptn string) bool {
	return str_endsWith(txt, ptn)
}

func Str_isValidStrBuilder() bool {
	return true
}

// ==== Builder
type Str_BuilderMtd interface {
	Get_txt() string
	Add(arg1 string)
	Len() LnsInt
	Clear()
}

func Str_Builder2Stem(obj LnsAny) LnsAny {
	return str_Builder2Stem(obj)
}

func (obj *Str_Builder) Str_ToBuilder() *Str_Builder {
	return obj.str_ToBuilder()
}

func NewStr_Builder() *Str_Builder {
	return newStr_Builder()
}
func (self *Str_Builder) InitStr_Builder() {
	self.initStr_Builder()
}

func (self *Str_Builder) Get_txt() string {
	return self.get_txt()
}

func (self *Str_Builder) Add(val string) {
	self.add(val)
}

func (self *Str_Builder) Len() LnsInt {
	return self.len()
}

func (self *Str_Builder) Clear() {
	self.clear()
}
