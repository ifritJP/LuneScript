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
nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

type Lns_iStream interface {
	Read(arg LnsAny) LnsAny
	Close()
}

// func NewLnsReader(reader io.Reader) *LnsReader {
// 	lnsReader := LnsReader{}
// 	lnsReader.orgReader = reader
// 	lnsReader.bufReader = nil
// 	return &lnsReader
// }

func (self *LnsReader) Read(arg LnsAny) LnsAny {
	return self.read(arg)
}
func (self *LnsReader) Close() {
	self.close()
}

func (self *Lns_stdin_t) Read(arg LnsAny) LnsAny {
	return self.reader.read(arg)
}
func (self *Lns_stdin_t) Close() {
	self.close()
}

type Lns_oStream interface {
	Write(arg string) (LnsAny, LnsAny)
	Flush()
	Close()
}

func (self *Lns_stdout_t) Write(arg string) (LnsAny, LnsAny) {
	return self.write(arg)
}
func (self *Lns_stdout_t) Flush() {
	self.flush()
}
func (self *Lns_stdout_t) Close() {
	self.close()
}

//================== io.open

type Lns_luaStream interface {
	Read(arg LnsAny) LnsAny
	Write(arg string) (LnsAny, LnsAny)
	Flush()
	Close()
	Seek(kind string, pos LnsInt) (LnsAny, LnsAny)
}

func Lns_io_open(path string, mode LnsAny) (LnsAny, LnsAny) {
	return lns_io_open(path, mode)
}

func (self *Lns_FileObj_t) Read(arg LnsAny) LnsAny {
	return self.read(arg)
}
func (self *Lns_FileObj_t) Write(arg string) (LnsAny, LnsAny) {
	return self.write(arg)
}
func (self *Lns_FileObj_t) Close() {
	self.close()
}
func (self *Lns_FileObj_t) Flush() {
	self.flush()
}
func (self *Lns_FileObj_t) Seek(kind string, pos LnsInt) (LnsAny, LnsAny) {
	return self.seek(kind, pos)
}
