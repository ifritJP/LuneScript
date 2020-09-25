package main

import "os"
import "io"
import "bytes"

type Lns_iStream interface {
    Gread( arg LnsAny ) LnsAny
    Gclose()
}
type Lns_stdin_t struct {
    FP Lns_iStream
}
var Lns_io_stdin *Lns_stdin_t

func (self *Lns_stdin_t) Gread( arg LnsAny) LnsAny {
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
func (self *Lns_stdin_t) Gclose() {
}

type Lns_oStream interface {
    Gwrite( arg string ) (LnsAny, LnsAny)
    Gflush()
    Gclose()
}
type Lns_stdout_t struct {
    FP Lns_oStream
}
var Lns_io_stdout *Lns_stdout_t

func (self *Lns_stdout_t) Gwrite( arg string ) (LnsAny, LnsAny) {
    _, err := os.Stdout.Write( []byte(arg) )
    return self, err
}
func (self *Lns_stdout_t) Gflush() {
}
func (self *Lns_stdout_t) Gclose() {
}




func init() {
    Lns_io_stdin = &Lns_stdin_t{}
    Lns_io_stdin.FP = Lns_io_stdin;
    Lns_io_stdout = &Lns_stdout_t{}
    Lns_io_stdout.FP = Lns_io_stdout
}
