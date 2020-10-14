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

import "os"
import "io"
import "bytes"

type Lns_iStream interface {
    Read( arg LnsAny ) LnsAny
    Close()
}
type Lns_stdin_t struct {
    FP Lns_iStream
}
var Lns_io_stdin *Lns_stdin_t

func (self *Lns_stdin_t) Read( arg LnsAny) LnsAny {
    buffer := bytes.Buffer{}
    buffer.Reset()

    buf := make([]byte,1024 * 100)

    for {
        readSize, err := os.Stdin.Read( buf )
        if err == io.EOF {
            break
        }
        buffer.Write( buf[:readSize] )
    }
    
    return string(buffer.Bytes())
}
func (self *Lns_stdin_t) Close() {
}

type Lns_oStream interface {
    Write( arg string ) (LnsAny, LnsAny)
    Flush()
    Close()
}
type Lns_stdout_t struct {
    FP Lns_oStream
}
var Lns_io_stdout *Lns_stdout_t
var Lns_io_stderr *Lns_stdout_t

func (self *Lns_stdout_t) Write( arg string ) (LnsAny, LnsAny) {
    var stream io.Writer
    if self == Lns_io_stdout {
        stream = os.Stdout
    } else {
        stream = os.Stderr
    }
    _, err := stream.Write( []byte(arg) )
    return self, err
}
func (self *Lns_stdout_t) Flush() {
}
func (self *Lns_stdout_t) Close() {
}




func init() {
    Lns_io_stdin = &Lns_stdin_t{}
    Lns_io_stdin.FP = Lns_io_stdin;
    Lns_io_stdout = &Lns_stdout_t{}
    Lns_io_stdout.FP = Lns_io_stdout
    Lns_io_stderr = &Lns_stdout_t{}
    Lns_io_stderr.FP = Lns_io_stderr
}
