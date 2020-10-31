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


type Lns_pipeMtd interface {
    Put( val LnsAny )
    Get() LnsAny
}
type Lns__pipe struct {
    ch chan LnsAny
    FP Lns_pipeMtd
}

func NewLnspipe( count int ) *Lns__pipe {
    pipe := &Lns__pipe{}
    pipe.ch = make(chan LnsAny,count)
    pipe.FP = pipe
    return pipe
}

func (self *Lns__pipe) Put( val LnsAny ) {
    self.ch<- val
}
func (self *Lns__pipe) Get() LnsAny {
    return <-self.ch
}
