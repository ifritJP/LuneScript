package main

import "os"

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"

func main() {
    args := []LnsAny{}
    for _, arg := range( os.Args ) {
        args = append( args, arg )
    }

    Lns_InitModOnce()
    //callmain
}
