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
nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

type Lns_iStream interface {
	Read(_env *LnsEnv, arg LnsAny) LnsAny
	Close(_env *LnsEnv)
}

// func NewLnsReader(reader io.Reader) *LnsReader {
// 	lnsReader := LnsReader{}
// 	lnsReader.orgReader = reader
// 	lnsReader.bufReader = nil
// 	return &lnsReader
// }

func (self *LnsReader) Read(_env *LnsEnv, arg LnsAny) LnsAny {
	return self.read(arg)
}
func (self *LnsReader) Close(_env *LnsEnv) {
	self.close()
}

func (self *Lns_stdin_t) Read(_env *LnsEnv, arg LnsAny) LnsAny {
	return self.read(arg)
}
func (self *Lns_stdin_t) Close(_env *LnsEnv) {
	self.close()
}

type Lns_oStream interface {
	Write(_env *LnsEnv, arg string) (LnsAny, LnsAny)
	Flush(_env *LnsEnv)
	Close(_env *LnsEnv)
}

func (self *Lns_stdout_t) Write(_env *LnsEnv, arg string) (LnsAny, LnsAny) {
	return self.write(arg)
}
func (self *Lns_stdout_t) Flush(_env *LnsEnv) {
	self.flush()
}
func (self *Lns_stdout_t) Close(_env *LnsEnv) {
	self.close()
}

func (self *Lns_FileObj_t) Read(_env *LnsEnv, arg LnsAny) LnsAny {
	return self.read(arg)
}
func (self *Lns_FileObj_t) Write(_env *LnsEnv, arg string) (LnsAny, LnsAny) {
	return self.write(arg)
}
func (self *Lns_FileObj_t) Close(_env *LnsEnv) {
	self.close()
}
func (self *Lns_FileObj_t) Flush(_env *LnsEnv) {
	self.flush()
}
func (self *Lns_FileObj_t) Seek(_env *LnsEnv, kind string, pos LnsInt) (LnsAny, LnsAny) {
	return self.seek(kind, pos)
}

//================== io.open

type Lns_luaStream interface {
	Read(_env *LnsEnv, arg LnsAny) LnsAny
	Write(_env *LnsEnv, arg string) (LnsAny, LnsAny)
	Flush(_env *LnsEnv)
	Close(_env *LnsEnv)
	Seek(_env *LnsEnv, kind string, pos LnsInt) (LnsAny, LnsAny)
}

func Lns_io_open(path string, mode LnsAny) (LnsAny, LnsAny) {
	return lns_io_open(path, mode)
}
