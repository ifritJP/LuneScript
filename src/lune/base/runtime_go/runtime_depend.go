/*
MIT License

Copyright (c) 2018 ifritJP

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

func Lns_Depend_init() {
}


/**
path の最終更新日時を取得する。

@param path ファイルパス
@return 1970/1/1 0:0:0 からの秒数。 取得失敗した場合は nil。
*/
func Depend_getFileLastModifiedTime( path string ) LnsAny {
    fileinfo, err := os.Stat( "miniGo.lns" )
    if err != nil {
        return nil
    }
    return fileinfo.ModTime().Unix()
}


func Depend_getLoadedMod() *LnsMap {
    return NewLnsMap( map[LnsAny]LnsAny{} )
}

func Depend_profile( validTest bool, work func(), path string ) LnsAny {
    work()
    return nil
}

func Depend_getStackTrace() string {
   return ""
}

func Depend_searchpath( mod string, pathPattern string ) LnsAny {
    return Lns_getVM().Package_searchpath( mod, pathPattern );
}

func Depend_existFile( path string ) bool {
    if _, err := os.Stat( path ); err != nil {
        return false
    }
    return true;
}

func Depend_getLuaVersion() string {
    return "5.3"
}
func Depend_setup( callback func( ver LnsInt )) {
    callback( 53 );
}

