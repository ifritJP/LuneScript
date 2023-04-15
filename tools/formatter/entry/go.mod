module github.com/ifritJP/LuneScript/tools

go 1.18


require (
	github.com/ifritJP/LuneScript/src v0.0.0-20230409073752-33f3ab0a6b08
	github.com/ifritJP/LuneScript/tools/formatter v0.0.0-20230409073752-33f3ab0a6b08
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/yuin/gopher-lua v0.0.0-20220413183635-c841877397d8 // indirect
)

replace github.com/ifritJP/LuneScript/src => ../../../src
replace github.com/ifritJP/LuneScript/tools/formatter => ../

//replace github.com/ifritJP/lnssqlite3/src => ../lnssqlite3/src
