module github.com/ifritJP/lnstags

go 1.18

require (
	github.com/ifritJP/LuneScript/src v0.0.0-20230304114223-928f73113833
	github.com/ifritJP/lnssqlite3/src v0.0.0-20230225123857-cfb8496afd8c
)

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/mattn/go-sqlite3 v1.14.7 // indirect
	github.com/yuin/gopher-lua v0.0.0-20220413183635-c841877397d8 // indirect
)

replace github.com/ifritJP/LuneScript/src => ../../src

//replace github.com/ifritJP/lnssqlite3/src => ../lnssqlite3/src
