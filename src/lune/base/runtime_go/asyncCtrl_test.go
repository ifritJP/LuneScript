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

import (
	"runtime"
	"testing"
)

// put した値を、終わり (nil) の前に全て get できることを確認する。
// 以前の get() は、チャネルが空かどうかを確認してから終わりを確認するまでの間に
// 残りの値と終わりが送られると、残りの値を捨てて nil を返していた。
func TestPipeGetAllItems(t *testing.T) {
	const items = 50
	for trial := 0; trial < 5000; trial++ {
		pipe := NewLnspipe(3)
		go func() {
			for index := 0; index < items; index++ {
				pipe.put(LnsInt(index))
				if index%7 == 0 {
					runtime.Gosched()
				}
			}
			pipe.put(nil)
		}()
		count := 0
		for pipe.get() != nil {
			count++
		}
		if count != items {
			t.Fatalf("trial %d: got %d items, expected %d", trial, count, items)
		}
	}
}

// 終わりの後の get() は nil を返し、終わりの後の put() は無視することを確認する。
func TestPipeAfterEnd(t *testing.T) {
	pipe := NewLnspipe(2)
	pipe.put(LnsInt(1))
	pipe.put(nil)
	pipe.put(LnsInt(2))
	pipe.put(nil)
	if val := pipe.get(); val != LnsInt(1) {
		t.Fatalf("got %v, expected 1", val)
	}
	for index := 0; index < 3; index++ {
		if val := pipe.get(); val != nil {
			t.Fatalf("got %v after the end, expected nil", val)
		}
	}
}
