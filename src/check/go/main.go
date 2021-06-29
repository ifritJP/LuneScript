package main

import (
	"fmt"
	"runtime"
	"time"
)

type Test struct {
	val int
}

func (self *Test) sub() int {
	return self.val
}

var list1 = make([]Test, 1)
var list2 = make([]interface{}, 1)

func sub1(test Test) {
	list1[0] = test
}
func sub2(test *Test) {
	list2[0] = test
}
func sub3(test Test) {
	list2[0] = test // escape
}
func sub4(test *Test) {
	list1[0] = *test
}

func profile(name string, callback func()) {
	runtime.GC()
	prev := time.Now()

	callback()

	fmt.Printf("%s: time = %v\n", name, time.Now().Sub(prev).Milliseconds())
}

func main() {
	maxCount := 100000 * 5000
	profile("sub1", func() {
		test := Test{}
		for count := 0; count < maxCount; count++ {
			sub1(test)
			test.sub()
		}
	})
	profile("sub2", func() {
		test := Test{}
		for count := 0; count < maxCount; count++ {
			sub2(&test)
		}
	})
	profile("sub3", func() {
		test := Test{}
		for count := 0; count < maxCount; count++ {
			sub3(test) // escape
		}
	})
	profile("sub4", func() {
		test := Test{}
		for count := 0; count < maxCount; count++ {
			sub4(&test)
		}
	})
}
