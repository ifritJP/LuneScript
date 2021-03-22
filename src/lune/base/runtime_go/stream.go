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
import "strings"
import "fmt"
import "bufio"

type Lns_iStream interface {
    Read( arg LnsAny ) LnsAny
    Close()
}

type LnsReader struct {
    bufReader *bufio.Reader
    orgReader io.Reader
}

func NewLnsReader( reader io.Reader ) *LnsReader {
    lnsReader := LnsReader{}
    lnsReader.orgReader = reader
    lnsReader.bufReader = nil
    return &lnsReader
}

func (self *LnsReader) Sync() {
    self.bufReader = nil
}

func (self *LnsReader) Read( arg LnsAny) LnsAny {
    if self.bufReader == nil {
        self.bufReader = bufio.NewReader( self.orgReader )
    }
    
    switch arg.(type) {
    case string:
        mode := arg.(string)
        switch mode {
        case "*a":
            buffer := bytes.Buffer{}
            buffer.Reset()

            buf := make([]byte,1024 * 100)

            for {
                readSize, err := self.bufReader.Read( buf )
                if err != nil && err != io.EOF {
                    break
                }
                if err == io.EOF && readSize == 0 {
                    break
                }
                
                buffer.Write( buf[:readSize] )
            }

            ret := buffer.Bytes()
            if len( ret ) == 0 {
                return nil
            }
            return string(ret)
        case "*l":
            buf, err := self.bufReader.ReadBytes( '\n' )
            if err != nil && err != io.EOF {
                return nil
            }
            if len( buf ) == 0 {
                return nil
            }
            if buf[ len( buf ) - 1 ] == '\n' {
                // 改行を除外して返す
                return string( buf[ : len( buf ) - 1 ] )
            }
            return string( buf )
        default:
            panic( fmt.Sprintf( "not support -- %s", mode ) );
        }
    case LnsInt:
        buf := make([]byte, arg.(LnsInt) )
        size, err := self.bufReader.Read( buf )
        if err != nil && err != io.EOF {
            return nil
        }
        if size == 0 {
            return nil
        }
        return string( buf[ :size] )
    default:
        panic( fmt.Sprintf( "not support -- %s", arg ) );
    }
}


type Lns_stdin_t struct {
    reader *LnsReader
    FP Lns_iStream
}
var Lns_io_stdin *Lns_stdin_t

func (self *Lns_stdin_t) Read( arg LnsAny) LnsAny {
    return self.reader.Read( arg )
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


//================== io.open

type Lns_luaStream interface {
    Read( arg LnsAny ) LnsAny
    Write( arg string ) (LnsAny, LnsAny)
    Flush()
    Close()
    Seek( kind string, pos LnsInt ) (LnsAny, LnsAny)
}


type Lns_FileObj_t struct {
    fileObj *os.File
    reader *LnsReader
    
    FP Lns_luaStream
}

func Lns_io_open( path string, mode LnsAny ) (LnsAny,LnsAny) {
    var modeStr string
    if mode == nil {
        modeStr = "r";
    } else {
        modeStr = mode.(string)
    }

    flag := 0
    
    if strings.Index( modeStr, "a" ) != -1 || strings.Index( modeStr, "w" ) != -1 {
        // a, w いずれかの場合
        flag |= os.O_RDWR
        if strings.Index( modeStr, "a" ) == -1 {
            // a がなければ create
            flag |= os.O_CREATE
        }
        if strings.Index( modeStr, "a" ) != -1 && strings.Index( modeStr, "+" ) != -1 {
            // a+ は append
            flag |= os.O_APPEND
        } else {
            flag |= os.O_TRUNC
        }
    } else {
        flag |= os.O_RDONLY
    }

    if fileObj, err := os.OpenFile( path, flag, 0666 ); err != nil {
        return nil, err.Error()
    } else {
        obj := &Lns_FileObj_t{ fileObj, nil, nil }
        obj.reader = NewLnsReader( fileObj )
        obj.FP = obj
        return obj, nil
    }
}

func (self *Lns_FileObj_t) Read( arg LnsAny) LnsAny {
    return self.reader.Read( arg )
}
func (self *Lns_FileObj_t) Write( arg string ) (LnsAny, LnsAny) {
    self.reader.Sync()
    _, err := self.fileObj.Write( []byte(arg) )
    return self, err
}
func (self *Lns_FileObj_t) Close() {
    if self.fileObj != nil {
        self.fileObj.Close()
        self.fileObj = nil
    }
}
func (self *Lns_FileObj_t) Flush() {
    self.fileObj.Sync()
}
func (self *Lns_FileObj_t) Seek( kind string, pos LnsInt ) (LnsAny, LnsAny) {
    self.reader.Sync()
    
    var ret int64
    var err error
    switch kind {
    case "set":
        ret, err = self.fileObj.Seek( int64(pos), os.SEEK_SET )
    case "cur":
        ret, err = self.fileObj.Seek( int64(pos), os.SEEK_CUR )
    case "end":
        ret, err = self.fileObj.Seek( int64(pos), os.SEEK_END )
    default:
        panic( fmt.Sprintf( "not support -- %s", kind ) )
    }
    if err != nil {
        return nil, err.Error()
    }
    return LnsInt(ret), nil
}




func init() {
    Lns_io_stdin = &Lns_stdin_t{}
    Lns_io_stdin.FP = Lns_io_stdin;
    Lns_io_stdin.reader = NewLnsReader( os.Stdin )
    Lns_io_stdout = &Lns_stdout_t{}
    Lns_io_stdout.FP = Lns_io_stdout
    Lns_io_stderr = &Lns_stdout_t{}
    Lns_io_stderr.FP = Lns_io_stderr
}
