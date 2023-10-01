package main

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import . "github.com/ifritJP/LuneScript/tools/runnerAnalyzer/lns"

//IMPORT:
////TEST:import . "lns/lune/base"

func main() {
    Lns_InitModOnce(LnsRuntimeOpt{ Int2strModeDepend })
    //TEST:Lns_Testing_init()
    Lns_RunMain( Main___main )
    //TEST:Testing_run( "" )
    //TEST:Testing_outputAllResult(Lns_io_stdout)
}
