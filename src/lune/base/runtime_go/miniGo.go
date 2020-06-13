package main

import "fmt"

func sub()  {
    Lns_print("hello world", "abc", "xyz", 1, 1.5, 10 + 100)
}
func sub2() LnsInt {
    return 1
}
func sub3() (LnsInt, string) {
    return 100, "abc"
}
func Lns_init() {
    fmt.Println( 1, sub2() )
    Lns_print( 1, sub3() )
}
