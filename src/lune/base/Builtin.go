// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Builtin bool
var Builtin__mod__ string
var Builtin_builtinFunc *Builtin_BuiltinFuncType
var Builtin_readyBuiltin bool
// for 435
func Builtin_convExp11696(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 427
func Builtin_convExp11577(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 68
func Builtin_convExp184(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 430
func Builtin_convExp11602(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 451
func Builtin_convExp11751(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 159: decl @lune.@base.@Builtin.getBuiltinFunc
func Builtin_getBuiltinFunc() *Builtin_BuiltinFuncType {
    return Builtin_builtinFunc
}



















// 163: decl @lune.@base.@Builtin.setupBuiltinTypeInfo
func Builtin_setupBuiltinTypeInfo_1796_(name string,fieldName string,symInfo *Ast_SymbolInfo) {
    var typeInfo *Ast_TypeInfo
    typeInfo = symInfo.FP.Get_typeInfo()
    var process_ func()
    process_ = func() {
        if _switch6301 := fieldName; _switch6301 == "_fcall" {
            Builtin_builtinFunc.Lns__fcall = typeInfo
            
            Builtin_builtinFunc.Lns__fcall_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "_kind" {
            Builtin_builtinFunc.Lns__kind = typeInfo
            
            Builtin_builtinFunc.Lns__kind_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "_load" {
            Builtin_builtinFunc.Lns__load = typeInfo
            
            Builtin_builtinFunc.Lns__load_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "collectgarbage" {
            Builtin_builtinFunc.Lns_collectgarbage = typeInfo
            
            Builtin_builtinFunc.Lns_collectgarbage_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "error" {
            Builtin_builtinFunc.Lns_error = typeInfo
            
            Builtin_builtinFunc.Lns_error_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "expandLuavalMap" {
            Builtin_builtinFunc.Lns_expandLuavalMap = typeInfo
            
            Builtin_builtinFunc.Lns_expandLuavalMap_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "loadfile" {
            Builtin_builtinFunc.Lns_loadfile = typeInfo
            
            Builtin_builtinFunc.Lns_loadfile_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "print" {
            Builtin_builtinFunc.Lns_print = typeInfo
            
            Builtin_builtinFunc.Lns_print_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "require" {
            Builtin_builtinFunc.Lns_require = typeInfo
            
            Builtin_builtinFunc.Lns_require_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "tonumber" {
            Builtin_builtinFunc.Lns_tonumber = typeInfo
            
            Builtin_builtinFunc.Lns_tonumber_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "tostring" {
            Builtin_builtinFunc.Lns_tostring = typeInfo
            
            Builtin_builtinFunc.Lns_tostring_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6301 == "type" {
            Builtin_builtinFunc.Lns_type = typeInfo
            
            Builtin_builtinFunc.Lns_type_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_iStream func()
    process_iStream = func() {
        if _switch6391 := fieldName; _switch6391 == "__attrib" {
            Builtin_builtinFunc.Istream___attrib = typeInfo
            
            Builtin_builtinFunc.Istream___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6391 == "close" {
            Builtin_builtinFunc.Istream_close = typeInfo
            
            Builtin_builtinFunc.Istream_close_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6391 == "read" {
            Builtin_builtinFunc.Istream_read = typeInfo
            
            Builtin_builtinFunc.Istream_read_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_oStream func()
    process_oStream = func() {
        if _switch6509 := fieldName; _switch6509 == "__attrib" {
            Builtin_builtinFunc.Ostream___attrib = typeInfo
            
            Builtin_builtinFunc.Ostream___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6509 == "close" {
            Builtin_builtinFunc.Ostream_close = typeInfo
            
            Builtin_builtinFunc.Ostream_close_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6509 == "flush" {
            Builtin_builtinFunc.Ostream_flush = typeInfo
            
            Builtin_builtinFunc.Ostream_flush_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6509 == "write" {
            Builtin_builtinFunc.Ostream_write = typeInfo
            
            Builtin_builtinFunc.Ostream_write_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_luaStream func()
    process_luaStream = func() {
        if _switch6683 := fieldName; _switch6683 == "__attrib" {
            Builtin_builtinFunc.Luastream___attrib = typeInfo
            
            Builtin_builtinFunc.Luastream___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6683 == "close" {
            Builtin_builtinFunc.Luastream_close = typeInfo
            
            Builtin_builtinFunc.Luastream_close_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6683 == "flush" {
            Builtin_builtinFunc.Luastream_flush = typeInfo
            
            Builtin_builtinFunc.Luastream_flush_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6683 == "read" {
            Builtin_builtinFunc.Luastream_read = typeInfo
            
            Builtin_builtinFunc.Luastream_read_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6683 == "seek" {
            Builtin_builtinFunc.Luastream_seek = typeInfo
            
            Builtin_builtinFunc.Luastream_seek_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6683 == "write" {
            Builtin_builtinFunc.Luastream_write = typeInfo
            
            Builtin_builtinFunc.Luastream_write_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Mapping func()
    process_Mapping = func() {
        if _switch6745 := fieldName; _switch6745 == "__attrib" {
            Builtin_builtinFunc.Mapping___attrib = typeInfo
            
            Builtin_builtinFunc.Mapping___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6745 == "_toMap" {
            Builtin_builtinFunc.Mapping__toMap = typeInfo
            
            Builtin_builtinFunc.Mapping__toMap_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process___pipe func()
    process___pipe = func() {
        if _switch6807 := fieldName; _switch6807 == "get" {
            Builtin_builtinFunc.G__pipe_get = typeInfo
            
            Builtin_builtinFunc.G__pipe_get_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6807 == "put" {
            Builtin_builtinFunc.G__pipe_put = typeInfo
            
            Builtin_builtinFunc.G__pipe_put_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_LnsThread func()
    process_LnsThread = func() {
        if _switch6869 := fieldName; _switch6869 == "__init" {
            Builtin_builtinFunc.Lnsthread___init = typeInfo
            
            Builtin_builtinFunc.Lnsthread___init_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch6869 == "loop" {
            Builtin_builtinFunc.Lnsthread_loop = typeInfo
            
            Builtin_builtinFunc.Lnsthread_loop_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_io func()
    process_io = func() {
        if _switch7043 := fieldName; _switch7043 == "__attrib" {
            Builtin_builtinFunc.Io___attrib = typeInfo
            
            Builtin_builtinFunc.Io___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7043 == "open" {
            Builtin_builtinFunc.Io_open = typeInfo
            
            Builtin_builtinFunc.Io_open_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7043 == "popen" {
            Builtin_builtinFunc.Io_popen = typeInfo
            
            Builtin_builtinFunc.Io_popen_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7043 == "stderr" {
            Builtin_builtinFunc.Io_stderr = typeInfo
            
            Builtin_builtinFunc.Io_stderr_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7043 == "stdin" {
            Builtin_builtinFunc.Io_stdin = typeInfo
            
            Builtin_builtinFunc.Io_stdin_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7043 == "stdout" {
            Builtin_builtinFunc.Io_stdout = typeInfo
            
            Builtin_builtinFunc.Io_stdout_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_package func()
    process_package = func() {
        if _switch7133 := fieldName; _switch7133 == "__attrib" {
            Builtin_builtinFunc.Package___attrib = typeInfo
            
            Builtin_builtinFunc.Package___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7133 == "path" {
            Builtin_builtinFunc.Package_path = typeInfo
            
            Builtin_builtinFunc.Package_path_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7133 == "searchpath" {
            Builtin_builtinFunc.Package_searchpath = typeInfo
            
            Builtin_builtinFunc.Package_searchpath_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_os func()
    process_os = func() {
        if _switch7363 := fieldName; _switch7363 == "__attrib" {
            Builtin_builtinFunc.Os___attrib = typeInfo
            
            Builtin_builtinFunc.Os___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "clock" {
            Builtin_builtinFunc.Os_clock = typeInfo
            
            Builtin_builtinFunc.Os_clock_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "date" {
            Builtin_builtinFunc.Os_date = typeInfo
            
            Builtin_builtinFunc.Os_date_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "difftime" {
            Builtin_builtinFunc.Os_difftime = typeInfo
            
            Builtin_builtinFunc.Os_difftime_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "exit" {
            Builtin_builtinFunc.Os_exit = typeInfo
            
            Builtin_builtinFunc.Os_exit_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "remove" {
            Builtin_builtinFunc.Os_remove = typeInfo
            
            Builtin_builtinFunc.Os_remove_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "rename" {
            Builtin_builtinFunc.Os_rename = typeInfo
            
            Builtin_builtinFunc.Os_rename_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7363 == "time" {
            Builtin_builtinFunc.Os_time = typeInfo
            
            Builtin_builtinFunc.Os_time_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_string func()
    process_string = func() {
        if _switch7705 := fieldName; _switch7705 == "__attrib" {
            Builtin_builtinFunc.String___attrib = typeInfo
            
            Builtin_builtinFunc.String___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "byte" {
            Builtin_builtinFunc.String_byte = typeInfo
            
            Builtin_builtinFunc.String_byte_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "dump" {
            Builtin_builtinFunc.String_dump = typeInfo
            
            Builtin_builtinFunc.String_dump_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "find" {
            Builtin_builtinFunc.String_find = typeInfo
            
            Builtin_builtinFunc.String_find_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "format" {
            Builtin_builtinFunc.String_format = typeInfo
            
            Builtin_builtinFunc.String_format_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "gmatch" {
            Builtin_builtinFunc.String_gmatch = typeInfo
            
            Builtin_builtinFunc.String_gmatch_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "gsub" {
            Builtin_builtinFunc.String_gsub = typeInfo
            
            Builtin_builtinFunc.String_gsub_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "lower" {
            Builtin_builtinFunc.String_lower = typeInfo
            
            Builtin_builtinFunc.String_lower_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "rep" {
            Builtin_builtinFunc.String_rep = typeInfo
            
            Builtin_builtinFunc.String_rep_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "reverse" {
            Builtin_builtinFunc.String_reverse = typeInfo
            
            Builtin_builtinFunc.String_reverse_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "sub" {
            Builtin_builtinFunc.String_sub = typeInfo
            
            Builtin_builtinFunc.String_sub_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch7705 == "upper" {
            Builtin_builtinFunc.String_upper = typeInfo
            
            Builtin_builtinFunc.String_upper_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_str func()
    process_str = func() {
        if _switch8019 := fieldName; _switch8019 == "__attrib" {
            Builtin_builtinFunc.Str___attrib = typeInfo
            
            Builtin_builtinFunc.Str___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "byte" {
            Builtin_builtinFunc.Str_byte = typeInfo
            
            Builtin_builtinFunc.Str_byte_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "find" {
            Builtin_builtinFunc.Str_find = typeInfo
            
            Builtin_builtinFunc.Str_find_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "format" {
            Builtin_builtinFunc.Str_format = typeInfo
            
            Builtin_builtinFunc.Str_format_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "gmatch" {
            Builtin_builtinFunc.Str_gmatch = typeInfo
            
            Builtin_builtinFunc.Str_gmatch_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "gsub" {
            Builtin_builtinFunc.Str_gsub = typeInfo
            
            Builtin_builtinFunc.Str_gsub_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "lower" {
            Builtin_builtinFunc.Str_lower = typeInfo
            
            Builtin_builtinFunc.Str_lower_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "rep" {
            Builtin_builtinFunc.Str_rep = typeInfo
            
            Builtin_builtinFunc.Str_rep_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "reverse" {
            Builtin_builtinFunc.Str_reverse = typeInfo
            
            Builtin_builtinFunc.Str_reverse_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "sub" {
            Builtin_builtinFunc.Str_sub = typeInfo
            
            Builtin_builtinFunc.Str_sub_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8019 == "upper" {
            Builtin_builtinFunc.Str_upper = typeInfo
            
            Builtin_builtinFunc.Str_upper_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_List func()
    process_List = func() {
        if _switch8165 := fieldName; _switch8165 == "__less" {
            Builtin_builtinFunc.List___less = typeInfo
            
            Builtin_builtinFunc.List___less_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8165 == "insert" {
            Builtin_builtinFunc.List_insert = typeInfo
            
            Builtin_builtinFunc.List_insert_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8165 == "remove" {
            Builtin_builtinFunc.List_remove = typeInfo
            
            Builtin_builtinFunc.List_remove_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8165 == "sort" {
            Builtin_builtinFunc.List_sort = typeInfo
            
            Builtin_builtinFunc.List_sort_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8165 == "unpack" {
            Builtin_builtinFunc.List_unpack = typeInfo
            
            Builtin_builtinFunc.List_unpack_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Array func()
    process_Array = func() {
        if _switch8255 := fieldName; _switch8255 == "__less" {
            Builtin_builtinFunc.Array___less = typeInfo
            
            Builtin_builtinFunc.Array___less_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8255 == "sort" {
            Builtin_builtinFunc.Array_sort = typeInfo
            
            Builtin_builtinFunc.Array_sort_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8255 == "unpack" {
            Builtin_builtinFunc.Array_unpack = typeInfo
            
            Builtin_builtinFunc.Array_unpack_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Set func()
    process_Set = func() {
        if _switch8485 := fieldName; _switch8485 == "add" {
            Builtin_builtinFunc.Set_add = typeInfo
            
            Builtin_builtinFunc.Set_add_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "and" {
            Builtin_builtinFunc.Set_and = typeInfo
            
            Builtin_builtinFunc.Set_and_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "clone" {
            Builtin_builtinFunc.Set_clone = typeInfo
            
            Builtin_builtinFunc.Set_clone_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "del" {
            Builtin_builtinFunc.Set_del = typeInfo
            
            Builtin_builtinFunc.Set_del_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "has" {
            Builtin_builtinFunc.Set_has = typeInfo
            
            Builtin_builtinFunc.Set_has_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "len" {
            Builtin_builtinFunc.Set_len = typeInfo
            
            Builtin_builtinFunc.Set_len_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "or" {
            Builtin_builtinFunc.Set_or = typeInfo
            
            Builtin_builtinFunc.Set_or_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8485 == "sub" {
            Builtin_builtinFunc.Set_sub = typeInfo
            
            Builtin_builtinFunc.Set_sub_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_math func()
    process_math = func() {
        if _switch8575 := fieldName; _switch8575 == "__attrib" {
            Builtin_builtinFunc.Math___attrib = typeInfo
            
            Builtin_builtinFunc.Math___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8575 == "random" {
            Builtin_builtinFunc.Math_random = typeInfo
            
            Builtin_builtinFunc.Math_random_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8575 == "randomseed" {
            Builtin_builtinFunc.Math_randomseed = typeInfo
            
            Builtin_builtinFunc.Math_randomseed_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_debug func()
    process_debug = func() {
        if _switch8665 := fieldName; _switch8665 == "__attrib" {
            Builtin_builtinFunc.Debug___attrib = typeInfo
            
            Builtin_builtinFunc.Debug___attrib_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8665 == "getinfo" {
            Builtin_builtinFunc.Debug_getinfo = typeInfo
            
            Builtin_builtinFunc.Debug_getinfo_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        } else if _switch8665 == "getlocal" {
            Builtin_builtinFunc.Debug_getlocal = typeInfo
            
            Builtin_builtinFunc.Debug_getlocal_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Nilable func()
    process_Nilable = func() {
        if _switch8699 := fieldName; _switch8699 == "val" {
            Builtin_builtinFunc.Nilable_val = typeInfo
            
            Builtin_builtinFunc.Nilable_val_sym = symInfo
            
            Builtin_builtinFunc.FP.Register(symInfo)
        }
    }
    if _switch8850 := name; _switch8850 == "" {
        process_()
    } else if _switch8850 == "iStream" {
        process_iStream()
    } else if _switch8850 == "oStream" {
        process_oStream()
    } else if _switch8850 == "luaStream" {
        process_luaStream()
    } else if _switch8850 == "Mapping" {
        process_Mapping()
    } else if _switch8850 == "__pipe" {
        process___pipe()
    } else if _switch8850 == "LnsThread" {
        process_LnsThread()
    } else if _switch8850 == "io" {
        process_io()
    } else if _switch8850 == "package" {
        process_package()
    } else if _switch8850 == "os" {
        process_os()
    } else if _switch8850 == "string" {
        process_string()
    } else if _switch8850 == "str" {
        process_str()
    } else if _switch8850 == "List" {
        process_List()
    } else if _switch8850 == "Array" {
        process_Array()
    } else if _switch8850 == "Set" {
        process_Set()
    } else if _switch8850 == "math" {
        process_math()
    } else if _switch8850 == "debug" {
        process_debug()
    } else if _switch8850 == "Nilable" {
        process_Nilable()
    }
}

// 173: decl @lune.@base.@Builtin.getBuiltInInfo
func Builtin_getBuiltInInfo_1849_() *LnsList {
    return NewLnsList([]LnsAny{NewLnsMap( map[LnsAny]LnsAny{"":NewLnsMap( map[LnsAny]LnsAny{"_fcall":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"form", "&..."}),"ret":NewLnsList([]LnsAny{""}),}),"_kind":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"int"}),}),"_load":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"collectgarbage":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),}),"error":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"__"}),}),"expandLuavalMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"Luaval<&stem>!"}),"ret":NewLnsList([]LnsAny{"&stem!"}),}),"loadfile":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"print":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{}),}),"require":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>"}),}),"tonumber":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!"}),"ret":NewLnsList([]LnsAny{"real!"}),}),"tostring":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem"}),"ret":NewLnsList([]LnsAny{"str"}),}),"type":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem!"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"iStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"oStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"luaStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"inplements":NewLnsList([]LnsAny{"iStream", "oStream"}),"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"seek":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"int!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Mapping":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"_toMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__pipe<T>":NewLnsMap( map[LnsAny]LnsAny{"get":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"put":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"LnsThread":NewLnsMap( map[LnsAny]LnsAny{"__init":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),"loop":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"abstract"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"io":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"open":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str!"}),"ret":NewLnsList([]LnsAny{"luaStream!", "str!"}),}),"popen":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<luaStream>!"}),}),"stderr":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),"stdin":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"iStream"}),}),"stdout":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"package":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"path":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"str"}),}),"searchpath":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"os":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"clock":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"real"}),}),"date":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str!", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>!"}),}),"difftime":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem", "&stem"}),"ret":NewLnsList([]LnsAny{"int"}),}),"exit":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"__"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"bool!", "str!"}),}),"rename":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),}),"time":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"string":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"dump":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Luaval<form>", "bool!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "..."}),"ret":NewLnsList([]LnsAny{"str"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"str"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"str":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"inplements":NewLnsList([]LnsAny{"Mapping"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int!>"}),"type":NewLnsList([]LnsAny{"method"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),"type":NewLnsList([]LnsAny{"method"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),"type":NewLnsList([]LnsAny{"method"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),"type":NewLnsList([]LnsAny{"method"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"List<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"insert":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{""}),"type":NewLnsList([]LnsAny{"mut"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Array<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Set<T>":NewLnsMap( map[LnsAny]LnsAny{"add":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"and":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"clone":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"method"}),}),"del":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"has":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"method"}),}),"len":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"int"}),"type":NewLnsList([]LnsAny{"method"}),}),"or":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"math":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"random":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"real"}),}),"randomseed":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"debug":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"getinfo":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"Map<str,stem>!"}),}),"getlocal":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int"}),"ret":NewLnsList([]LnsAny{"str!", "stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Nilable<_T>":NewLnsMap( map[LnsAny]LnsAny{"val":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"_T!"}),"type":NewLnsList([]LnsAny{"method"}),}),}),})})
}


// 600: decl @lune.@base.@Builtin.Builtin.registBuiltInScope.processCopyAlterList
func Builtin_registBuiltInScope__processCopyAlterList_2593_(alterList *LnsList,typeList *LnsList) {
    for _, _typeInfo := range( typeList.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        alterList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( Ast_AlternateTypeInfoDownCastF(typeInfo.FP)).(*Ast_AlternateTypeInfo)))
    }
}

// declaration Class -- Builtin
type Builtin_BuiltinMtd interface {
    getTypeInfo(arg1 string) *Ast_TypeInfo
    processField(arg1 string, arg2 string, arg3 *LnsMap, arg4 *Ast_TypeInfo)
    RegistBuiltInScope() *Builtin_BuiltinFuncType
}
type Builtin_Builtin struct {
    transUnitIF TransUnitIF_TransUnitIF
    targetLuaVer *LuaVer_LuaVerInfo
    ctrl_info *Types_TransCtrlInfo
    processInfo *Ast_ProcessInfo
    modifier *TransUnitIF_Modifier
    FP Builtin_BuiltinMtd
}
func Builtin_Builtin2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Builtin_Builtin).FP
}
type Builtin_BuiltinDownCast interface {
    ToBuiltin_Builtin() *Builtin_Builtin
}
func Builtin_BuiltinDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Builtin_BuiltinDownCast)
    if ok { return work.ToBuiltin_Builtin() }
    return nil
}
func (obj *Builtin_Builtin) ToBuiltin_Builtin() *Builtin_Builtin {
    return obj
}
func NewBuiltin_Builtin(arg1 TransUnitIF_TransUnitIF, arg2 *LuaVer_LuaVerInfo, arg3 *Types_TransCtrlInfo) *Builtin_Builtin {
    obj := &Builtin_Builtin{}
    obj.FP = obj
    obj.InitBuiltin_Builtin(arg1, arg2, arg3)
    return obj
}
// 46: DeclConstr
func (self *Builtin_Builtin) InitBuiltin_Builtin(transUnitIF TransUnitIF_TransUnitIF,targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo) {
    self.transUnitIF = transUnitIF
    
    self.targetLuaVer = targetLuaVer
    
    self.ctrl_info = ctrl_info
    
    self.processInfo = Ast_getRootProcessInfo()
    
    self.modifier = NewTransUnitIF_Modifier(true, self.processInfo)
    
}

// 352: decl @lune.@base.@Builtin.Builtin.getTypeInfo
func (self *Builtin_Builtin) getTypeInfo(typeName string) *Ast_TypeInfo {
    if _switch11260 := typeName; _switch11260 == "_T" {
        return Ast_builtinTypeBox.FP.Get_boxingType()
    } else if _switch11260 == "_T!" {
        return Ast_builtinTypeBox.FP.Get_boxingType().FP.Get_nilableTypeInfo()
    }
    var getTypeInfoFromScope func(scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny
    getTypeInfoFromScope = func(scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny {
        var typeInfo *Ast_TypeInfo
        
        {
            _typeInfo := scope.FP.GetTypeInfo(symbol, scope, false, Ast_ScopeAccess__Normal)
            if _typeInfo == nil{
                return nil
            } else {
                typeInfo = _typeInfo.(*Ast_TypeInfo)
            }
        }
        var validGenType bool
        validGenType = false
        for _, _genType := range( genTypeList.Items ) {
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if genType.FP.Get_kind() != Ast_TypeInfoKind__Alternate{
                validGenType = true
                
                break
            }
        }
        if validGenType{
            if _switch11531 := typeInfo.FP.Get_kind(); _switch11531 == Ast_TypeInfoKind__Map {
                if genTypeList.Len() != 2{
                    Util_err(Lns_getVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                }
                var keyType *Ast_TypeInfo
                keyType = genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var valType *Ast_TypeInfo
                valType = genTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                return self.processInfo.FP.CreateMap(Ast_AccessMode__Pub, typeInfo.FP.Get_parentInfo(), keyType, valType, typeInfo.FP.Get_mutMode())
            } else if _switch11531 == Ast_TypeInfoKind__Ext {
                if genTypeList.Len() != 1{
                    Util_err(Lns_getVM().String_format("illegal param -- %d", []LnsAny{genTypeList.Len()}))
                }
                switch _exp11461 := self.processInfo.FP.CreateLuaval(genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true).(type) {
                case *Ast_LuavalResult__OK:
                workType := _exp11461.Val1
                    if self.ctrl_info.ValidLuaval{
                        return workType
                    }
                    return genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                case *Ast_LuavalResult__Err:
                mess := _exp11461.Val1
                    Util_err(mess)
                }
            } else if _switch11531 == Ast_TypeInfoKind__DDD {
                if genTypeList.Len() != 1{
                    Util_err(Lns_getVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                }
                return &self.processInfo.FP.CreateDDD(genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true, false).Ast_TypeInfo
            } else {
                Util_err(Lns_getVM().String_format("not support type -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        }
        return typeInfo
    }
    var mutable bool
    mutable = true
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(typeName,"^&", nil, nil))){
        mutable = false
        
        typeName = Builtin_convExp11577(Lns_2DDD(Lns_getVM().String_gsub(typeName,"^&", "")))
        
    }
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    var endIndex LnsAny
    _,endIndex = Builtin_convExp11602(Lns_2DDD(Lns_getVM().String_find(typeName,"[%w%.]+<", nil, nil)))
    var suffix string
    suffix = ""
    if endIndex != nil{
        endIndex_3816 := endIndex.(LnsInt)
        var genTypeName string
        genTypeName = Lns_getVM().String_sub(typeName,endIndex_3816 + 1, nil)
        for  {
            {
                _tailIndex := Builtin_convExp11696(Lns_2DDD(Lns_getVM().String_find(genTypeName,"[,>]", nil, nil)))
                if !Lns_IsNil( _tailIndex ) {
                    tailIndex := _tailIndex.(LnsInt)
                    var genType *Ast_TypeInfo
                    genType = self.FP.getTypeInfo(Lns_getVM().String_sub(genTypeName,1, tailIndex - 1))
                    genTypeList.Insert(Ast_TypeInfo2Stem(genType))
                    genTypeName = Lns_getVM().String_sub(genTypeName,tailIndex + 1, nil)
                    
                } else {
                    suffix = Lns_getVM().String_sub(genTypeName,1, nil)
                    
                    break
                }
            }
        }
        typeName = Lns_getVM().String_sub(typeName,1, endIndex_3816 - 1) + suffix
        
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(typeName,"!$", nil, nil))){
        var orgTypeName string
        orgTypeName = Builtin_convExp11751(Lns_2DDD(Lns_getVM().String_gsub(typeName,"!$", "")))
        {
            __exp := getTypeInfoFromScope(self.transUnitIF.Get_scope(), orgTypeName, genTypeList)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                typeInfo = _exp
                
            } else {
                Util_err(Lns_getVM().String_format("not found builtin -- %s", []LnsAny{orgTypeName}))
            }
        }
        typeInfo = typeInfo.FP.Get_nilableTypeInfo()
        
    } else { 
        {
            __exp := getTypeInfoFromScope(self.transUnitIF.Get_scope(), typeName, genTypeList)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                typeInfo = _exp
                
            } else {
                Util_err(Lns_getVM().String_format("not found builtin -- %s", []LnsAny{typeName}))
            }
        }
    }
    if mutable{
        return typeInfo
    }
    typeInfo = self.modifier.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
    
    return typeInfo
}

// 475: decl @lune.@base.@Builtin.Builtin.processField
func (self *Builtin_Builtin) processField(name string,fieldName string,info *LnsMap,parentInfo *Ast_TypeInfo) {
    if self.targetLuaVer.FP.IsSupport(Lns_getVM().String_format("%s.%s", []LnsAny{name, fieldName})){
        if Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
        info.Get("type")) && 
        Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1))) == "var"{
            var symbol *Ast_SymbolInfo
            symbol = Lns_unwrap( Lns_car(self.transUnitIF.Get_scope().FP.Add(self.processInfo, Ast_SymbolKind__Var, false, true, fieldName, nil, self.FP.getTypeInfo(Lns_unwrap( Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
            info.Get("typeInfo")) && 
            Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1)))).(string)), Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_1796_(name, fieldName, symbol)
        } else { 
            var argTypeList *LnsList
            argTypeList = NewLnsList([]LnsAny{})
            for _, _argType := range( Lns_unwrap( info.Get("arg")).(*LnsList).Items ) {
                argType := _argType.(string)
                argTypeList.Insert(Ast_TypeInfo2Stem(self.FP.getTypeInfo(argType)))
            }
            var retTypeList *LnsList
            retTypeList = NewLnsList([]LnsAny{})
            for _, _retType := range( Lns_unwrap( info.Get("ret")).(*LnsList).Items ) {
                retType := _retType.(string)
                var retTypeInfo *Ast_TypeInfo
                retTypeInfo = self.FP.getTypeInfo(retType)
                retTypeList.Insert(Ast_TypeInfo2Stem(retTypeInfo))
            }
            var funcType LnsAny
            funcType = Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
            info.Get("type")) && 
            Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1)))
            var mutable bool
            var staticFlag bool
            var kind LnsInt
            var symbolKind LnsInt
            var abstractFlag bool
            var accessMode LnsInt
            if _switch12312 := funcType; _switch12312 == "method" || _switch12312 == "mut" {
                staticFlag = false
                
                kind = Ast_TypeInfoKind__Method
                
                symbolKind = Ast_SymbolKind__Mtd
                
                abstractFlag = false
                
                accessMode = Ast_AccessMode__Pub
                
                mutable = funcType == "mut"
                
            } else if _switch12312 == "abstract" {
                staticFlag = false
                
                kind = Ast_TypeInfoKind__Method
                
                symbolKind = Ast_SymbolKind__Mtd
                
                abstractFlag = true
                
                accessMode = Ast_AccessMode__Pro
                
                mutable = true
                
            } else if _switch12312 == "macro" {
                staticFlag = true
                
                kind = Ast_TypeInfoKind__Macro
                
                symbolKind = Ast_SymbolKind__Mac
                
                abstractFlag = false
                
                accessMode = Ast_AccessMode__Pub
                
                mutable = false
                
            } else if _switch12312 == "formfunc" {
                staticFlag = true
                
                kind = Ast_TypeInfoKind__FormFunc
                
                symbolKind = Ast_SymbolKind__Typ
                
                abstractFlag = false
                
                accessMode = Ast_AccessMode__Pub
                
                mutable = false
                
            } else {
                staticFlag = true
                
                kind = Ast_TypeInfoKind__Func
                
                symbolKind = Ast_SymbolKind__Fun
                
                abstractFlag = false
                
                accessMode = Ast_AccessMode__Pub
                
                mutable = false
                
            }
            self.transUnitIF.PushScope(false, nil, nil)
            var typeInfo *Ast_NormalTypeInfo
            typeInfo = self.processInfo.FP.CreateFunc(abstractFlag, true, self.transUnitIF.Get_scope(), kind, parentInfo, false, true, staticFlag, accessMode, fieldName, nil, argTypeList, retTypeList, mutable)
            self.transUnitIF.PopScope()
            Builtin_builtinFunc.FP.Get_allFuncTypeSet().Add(Ast_NormalTypeInfo2Stem(typeInfo))
            Ast_addBuiltin(&typeInfo.Ast_TypeInfo)
            if typeInfo.FP.Get_nilableTypeInfo() != Ast_headTypeInfo{
                Ast_addBuiltin(typeInfo.FP.Get_nilableTypeInfo())
            }
            var symInfo *Ast_SymbolInfo
            symInfo = Lns_unwrap( Lns_car(self.transUnitIF.Get_scope().FP.Add(self.processInfo, symbolKind, false, kind == Ast_TypeInfoKind__Func, fieldName, nil, &typeInfo.Ast_TypeInfo, accessMode, staticFlag, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( mutable) &&
                Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
                Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_1796_(name, fieldName, symInfo)
        }
    }
}

// 576: decl @lune.@base.@Builtin.Builtin.registBuiltInScope
func (self *Builtin_Builtin) RegistBuiltInScope() *Builtin_BuiltinFuncType {
    if Builtin_readyBuiltin{
        return Builtin_builtinFunc
    }
    Builtin_readyBuiltin = true
    
    var builtInInfo *LnsList
    builtInInfo = Builtin_getBuiltInInfo_1849_()
    var builtinModuleName2Scope *LnsMap
    builtinModuleName2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    var mapType *Ast_TypeInfo
    mapType = self.processInfo.FP.CreateMap(Ast_AccessMode__Pub, Ast_headTypeInfo, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
    self.transUnitIF.Get_scope().FP.AddVar(self.processInfo, Ast_AccessMode__Global, "_ENV", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnitIF.Get_scope().FP.AddVar(self.processInfo, Ast_AccessMode__Global, "_G", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnitIF.Get_scope().FP.AddVar(self.processInfo, Ast_AccessMode__Global, "__line__", nil, Ast_builtinTypeInt, Ast_MutMode__IMut, true)
    for _, _builtinClassInfo := range( builtInInfo.Items ) {
        builtinClassInfo := _builtinClassInfo.(*LnsMap)
        for _className, _name2FieldInfo := range( builtinClassInfo.Items ) {
            className := _className.(string)
            name2FieldInfo := _name2FieldInfo.(*LnsMap)
            var name string
            name = className
            var genTypeList *LnsList
            genTypeList = NewLnsList([]LnsAny{})
            if _switch12842 := className; _switch12842 == "List<T>" {
                name = "List"
                
                Builtin_registBuiltInScope__processCopyAlterList_2593_(genTypeList, Ast_builtinTypeList.FP.Get_itemTypeInfoList())
            } else if _switch12842 == "Array<T>" {
                name = "Array"
                
                Builtin_registBuiltInScope__processCopyAlterList_2593_(genTypeList, Ast_builtinTypeArray.FP.Get_itemTypeInfoList())
            } else if _switch12842 == "Set<T>" {
                name = "Set"
                
                Builtin_registBuiltInScope__processCopyAlterList_2593_(genTypeList, Ast_builtinTypeSet.FP.Get_itemTypeInfoList())
            } else if _switch12842 == "Nilable<_T>" {
                name = "Nilable"
                
                Builtin_registBuiltInScope__processCopyAlterList_2593_(genTypeList, Ast_builtinTypeBox.FP.Get_itemTypeInfoList())
            } else {
                if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(className,"<", nil, nil))){
                    name = ""
                    
                    {
                        _form12838, _param12838, _prev12838 := Lns_getVM().String_gmatch(className,"[^<>,%s]+")
                        for {
                            _work12838 := _form12838.(*Lns_luaValue).Call( Lns_2DDD( _param12838, _prev12838 ) )
                            _prev12838 = Lns_getFromMulti(_work12838,0)
                            if Lns_IsNil( _prev12838 ) { break }
                            token := _prev12838.(string)
                            if len(name) == 0{
                                name = token
                                
                            } else { 
                                genTypeList.Insert(Ast_AlternateTypeInfo2Stem(self.processInfo.FP.CreateAlternate(true, genTypeList.Len() + 1, token, Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
                            }
                        }
                    }
                }
            }
            var parentInfo *Ast_TypeInfo
            parentInfo = Ast_headTypeInfo
            if name != ""{
                var classKind LnsInt
                classKind = TransUnitIF_DeclClassMode__Class
                {
                    _types := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                    name2FieldInfo.Get("__attrib")) && 
                    Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Get("type")))
                    if !Lns_IsNil( _types ) {
                        types := _types.(*LnsList)
                        if types.Len() > 0{
                            if _switch12902 := types.GetAt(1).(string); _switch12902 == "interface" {
                                classKind = TransUnitIF_DeclClassMode__Interface
                                
                            } else if _switch12902 == "module" {
                                classKind = TransUnitIF_DeclClassMode__Module
                                
                            }
                        }
                    }
                }
                var interfaceList *LnsList
                interfaceList = NewLnsList([]LnsAny{})
                {
                    __exp := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                    name2FieldInfo.Get("__attrib")) && 
                    Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Get("inplements")))
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*LnsList)
                        for _, _ifname := range( _exp.Items ) {
                            ifname := _ifname.(string)
                            var ifType *Ast_TypeInfo
                            ifType = self.FP.getTypeInfo(ifname)
                            interfaceList.Insert(Ast_TypeInfo2Stem(ifType))
                        }
                    }
                }
                if _switch13089 := classKind; _switch13089 == TransUnitIF_DeclClassMode__Class || _switch13089 == TransUnitIF_DeclClassMode__Interface {
                    var declMode LnsInt
                    if classKind == TransUnitIF_DeclClassMode__Class{
                        declMode = TransUnitIF_DeclClassMode__Class
                        
                    } else { 
                        declMode = TransUnitIF_DeclClassMode__Interface
                        
                    }
                    parentInfo = self.transUnitIF.PushClass(self.processInfo, self.transUnitIF.GetLatestPos(), declMode, false, nil, interfaceList, genTypeList, true, name, true, Ast_AccessMode__Pub, nil)
                    
                    Builtin_builtinFunc.FP.RegisterClass(parentInfo)
                } else if _switch13089 == TransUnitIF_DeclClassMode__Module {
                    parentInfo = self.transUnitIF.PushModule(self.processInfo, true, name, true)
                    
                    self.transUnitIF.Get_scope().FP.Get_parent().FP.Add(self.processInfo, Ast_SymbolKind__Typ, false, false, name, nil, parentInfo, Ast_AccessMode__Local, true, Ast_MutMode__Mut, true, false)
                }
                Ast_addBuiltin(parentInfo)
                Ast_addBuiltin(parentInfo.FP.Get_nilableTypeInfo())
            }
            if Lns_op_not(builtinModuleName2Scope.Get(name)){
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( name != "") &&
                    Lns_GetEnv().SetStackVal( self.FP.getTypeInfo(name)) )){
                    builtinModuleName2Scope.Set(name,self.transUnitIF.Get_scope())
                }
                {
                    __collection13171 := name2FieldInfo
                    __sorted13171 := __collection13171.CreateKeyListStr()
                    __sorted13171.Sort( LnsItemKindStr, nil )
                    for _, _fieldName := range( __sorted13171.Items ) {
                        info := __collection13171.Items[ _fieldName ].(*LnsMap)
                        fieldName := _fieldName.(string)
                        if _switch13169 := fieldName; _switch13169 == "__attrib" {
                        } else {
                            self.FP.processField(name, fieldName, info, parentInfo)
                        }
                    }
                }
            }
            if name != ""{
                self.transUnitIF.PopClass()
            }
        }
    }
    var threadSafeSet *LnsSet
    threadSafeSet = NewLnsSet([]LnsAny{Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_error), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_print), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_type), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_tonumber), Ast_TypeInfo2Stem(Builtin_builtinFunc.Io_open), Ast_TypeInfo2Stem(Builtin_builtinFunc.Set_has), Ast_TypeInfo2Stem(Builtin_builtinFunc.Set_add)})
    for _typeInfo := range( Builtin_builtinFunc.FP.Get_allFuncTypeSet().Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(threadSafeSet.Has(Ast_TypeInfo2Stem(typeInfo))){
            Builtin_builtinFunc.FP.Get_needThreadingTypes().Add(Ast_TypeInfo2Stem(typeInfo))
        }
    }
    return Builtin_builtinFunc
}


// declaration Class -- BuiltinFuncType
type Builtin_BuiltinFuncTypeMtd interface {
    Get_allClass() *LnsList
    Get_allFuncTypeSet() *LnsSet
    Get_allSymbol() *LnsList
    Get_allSymbolSet() *LnsSet
    Get_needThreadingTypes() *LnsSet
    Register(arg1 *Ast_SymbolInfo)
    RegisterClass(arg1 *Ast_TypeInfo)
}
type Builtin_BuiltinFuncType struct {
    Lns_ *Ast_TypeInfo
    Lns__fcall *Ast_TypeInfo
    Lns__fcall_sym *Ast_SymbolInfo
    Lns__kind *Ast_TypeInfo
    Lns__kind_sym *Ast_SymbolInfo
    Lns__load *Ast_TypeInfo
    Lns__load_sym *Ast_SymbolInfo
    Lns_collectgarbage *Ast_TypeInfo
    Lns_collectgarbage_sym *Ast_SymbolInfo
    Lns_error *Ast_TypeInfo
    Lns_error_sym *Ast_SymbolInfo
    Lns_expandLuavalMap *Ast_TypeInfo
    Lns_expandLuavalMap_sym *Ast_SymbolInfo
    Lns_loadfile *Ast_TypeInfo
    Lns_loadfile_sym *Ast_SymbolInfo
    Lns_print *Ast_TypeInfo
    Lns_print_sym *Ast_SymbolInfo
    Lns_require *Ast_TypeInfo
    Lns_require_sym *Ast_SymbolInfo
    Lns_tonumber *Ast_TypeInfo
    Lns_tonumber_sym *Ast_SymbolInfo
    Lns_tostring *Ast_TypeInfo
    Lns_tostring_sym *Ast_SymbolInfo
    Lns_type *Ast_TypeInfo
    Lns_type_sym *Ast_SymbolInfo
    Istream_ *Ast_TypeInfo
    Istream___attrib *Ast_TypeInfo
    Istream___attrib_sym *Ast_SymbolInfo
    Istream_close *Ast_TypeInfo
    Istream_close_sym *Ast_SymbolInfo
    Istream_read *Ast_TypeInfo
    Istream_read_sym *Ast_SymbolInfo
    Ostream_ *Ast_TypeInfo
    Ostream___attrib *Ast_TypeInfo
    Ostream___attrib_sym *Ast_SymbolInfo
    Ostream_close *Ast_TypeInfo
    Ostream_close_sym *Ast_SymbolInfo
    Ostream_flush *Ast_TypeInfo
    Ostream_flush_sym *Ast_SymbolInfo
    Ostream_write *Ast_TypeInfo
    Ostream_write_sym *Ast_SymbolInfo
    Luastream_ *Ast_TypeInfo
    Luastream___attrib *Ast_TypeInfo
    Luastream___attrib_sym *Ast_SymbolInfo
    Luastream_close *Ast_TypeInfo
    Luastream_close_sym *Ast_SymbolInfo
    Luastream_flush *Ast_TypeInfo
    Luastream_flush_sym *Ast_SymbolInfo
    Luastream_read *Ast_TypeInfo
    Luastream_read_sym *Ast_SymbolInfo
    Luastream_seek *Ast_TypeInfo
    Luastream_seek_sym *Ast_SymbolInfo
    Luastream_write *Ast_TypeInfo
    Luastream_write_sym *Ast_SymbolInfo
    Mapping_ *Ast_TypeInfo
    Mapping___attrib *Ast_TypeInfo
    Mapping___attrib_sym *Ast_SymbolInfo
    Mapping__toMap *Ast_TypeInfo
    Mapping__toMap_sym *Ast_SymbolInfo
    G__pipe_ *Ast_TypeInfo
    G__pipe_get *Ast_TypeInfo
    G__pipe_get_sym *Ast_SymbolInfo
    G__pipe_put *Ast_TypeInfo
    G__pipe_put_sym *Ast_SymbolInfo
    Lnsthread_ *Ast_TypeInfo
    Lnsthread___init *Ast_TypeInfo
    Lnsthread___init_sym *Ast_SymbolInfo
    Lnsthread_loop *Ast_TypeInfo
    Lnsthread_loop_sym *Ast_SymbolInfo
    Io_ *Ast_TypeInfo
    Io___attrib *Ast_TypeInfo
    Io___attrib_sym *Ast_SymbolInfo
    Io_open *Ast_TypeInfo
    Io_open_sym *Ast_SymbolInfo
    Io_popen *Ast_TypeInfo
    Io_popen_sym *Ast_SymbolInfo
    Io_stderr *Ast_TypeInfo
    Io_stderr_sym *Ast_SymbolInfo
    Io_stdin *Ast_TypeInfo
    Io_stdin_sym *Ast_SymbolInfo
    Io_stdout *Ast_TypeInfo
    Io_stdout_sym *Ast_SymbolInfo
    Package_ *Ast_TypeInfo
    Package___attrib *Ast_TypeInfo
    Package___attrib_sym *Ast_SymbolInfo
    Package_path *Ast_TypeInfo
    Package_path_sym *Ast_SymbolInfo
    Package_searchpath *Ast_TypeInfo
    Package_searchpath_sym *Ast_SymbolInfo
    Os_ *Ast_TypeInfo
    Os___attrib *Ast_TypeInfo
    Os___attrib_sym *Ast_SymbolInfo
    Os_clock *Ast_TypeInfo
    Os_clock_sym *Ast_SymbolInfo
    Os_date *Ast_TypeInfo
    Os_date_sym *Ast_SymbolInfo
    Os_difftime *Ast_TypeInfo
    Os_difftime_sym *Ast_SymbolInfo
    Os_exit *Ast_TypeInfo
    Os_exit_sym *Ast_SymbolInfo
    Os_remove *Ast_TypeInfo
    Os_remove_sym *Ast_SymbolInfo
    Os_rename *Ast_TypeInfo
    Os_rename_sym *Ast_SymbolInfo
    Os_time *Ast_TypeInfo
    Os_time_sym *Ast_SymbolInfo
    String_ *Ast_TypeInfo
    String___attrib *Ast_TypeInfo
    String___attrib_sym *Ast_SymbolInfo
    String_byte *Ast_TypeInfo
    String_byte_sym *Ast_SymbolInfo
    String_dump *Ast_TypeInfo
    String_dump_sym *Ast_SymbolInfo
    String_find *Ast_TypeInfo
    String_find_sym *Ast_SymbolInfo
    String_format *Ast_TypeInfo
    String_format_sym *Ast_SymbolInfo
    String_gmatch *Ast_TypeInfo
    String_gmatch_sym *Ast_SymbolInfo
    String_gsub *Ast_TypeInfo
    String_gsub_sym *Ast_SymbolInfo
    String_lower *Ast_TypeInfo
    String_lower_sym *Ast_SymbolInfo
    String_rep *Ast_TypeInfo
    String_rep_sym *Ast_SymbolInfo
    String_reverse *Ast_TypeInfo
    String_reverse_sym *Ast_SymbolInfo
    String_sub *Ast_TypeInfo
    String_sub_sym *Ast_SymbolInfo
    String_upper *Ast_TypeInfo
    String_upper_sym *Ast_SymbolInfo
    Str_ *Ast_TypeInfo
    Str___attrib *Ast_TypeInfo
    Str___attrib_sym *Ast_SymbolInfo
    Str_byte *Ast_TypeInfo
    Str_byte_sym *Ast_SymbolInfo
    Str_find *Ast_TypeInfo
    Str_find_sym *Ast_SymbolInfo
    Str_format *Ast_TypeInfo
    Str_format_sym *Ast_SymbolInfo
    Str_gmatch *Ast_TypeInfo
    Str_gmatch_sym *Ast_SymbolInfo
    Str_gsub *Ast_TypeInfo
    Str_gsub_sym *Ast_SymbolInfo
    Str_lower *Ast_TypeInfo
    Str_lower_sym *Ast_SymbolInfo
    Str_rep *Ast_TypeInfo
    Str_rep_sym *Ast_SymbolInfo
    Str_reverse *Ast_TypeInfo
    Str_reverse_sym *Ast_SymbolInfo
    Str_sub *Ast_TypeInfo
    Str_sub_sym *Ast_SymbolInfo
    Str_upper *Ast_TypeInfo
    Str_upper_sym *Ast_SymbolInfo
    List_ *Ast_TypeInfo
    List___less *Ast_TypeInfo
    List___less_sym *Ast_SymbolInfo
    List_insert *Ast_TypeInfo
    List_insert_sym *Ast_SymbolInfo
    List_remove *Ast_TypeInfo
    List_remove_sym *Ast_SymbolInfo
    List_sort *Ast_TypeInfo
    List_sort_sym *Ast_SymbolInfo
    List_unpack *Ast_TypeInfo
    List_unpack_sym *Ast_SymbolInfo
    Array_ *Ast_TypeInfo
    Array___less *Ast_TypeInfo
    Array___less_sym *Ast_SymbolInfo
    Array_sort *Ast_TypeInfo
    Array_sort_sym *Ast_SymbolInfo
    Array_unpack *Ast_TypeInfo
    Array_unpack_sym *Ast_SymbolInfo
    Set_ *Ast_TypeInfo
    Set_add *Ast_TypeInfo
    Set_add_sym *Ast_SymbolInfo
    Set_and *Ast_TypeInfo
    Set_and_sym *Ast_SymbolInfo
    Set_clone *Ast_TypeInfo
    Set_clone_sym *Ast_SymbolInfo
    Set_del *Ast_TypeInfo
    Set_del_sym *Ast_SymbolInfo
    Set_has *Ast_TypeInfo
    Set_has_sym *Ast_SymbolInfo
    Set_len *Ast_TypeInfo
    Set_len_sym *Ast_SymbolInfo
    Set_or *Ast_TypeInfo
    Set_or_sym *Ast_SymbolInfo
    Set_sub *Ast_TypeInfo
    Set_sub_sym *Ast_SymbolInfo
    Math_ *Ast_TypeInfo
    Math___attrib *Ast_TypeInfo
    Math___attrib_sym *Ast_SymbolInfo
    Math_random *Ast_TypeInfo
    Math_random_sym *Ast_SymbolInfo
    Math_randomseed *Ast_TypeInfo
    Math_randomseed_sym *Ast_SymbolInfo
    Debug_ *Ast_TypeInfo
    Debug___attrib *Ast_TypeInfo
    Debug___attrib_sym *Ast_SymbolInfo
    Debug_getinfo *Ast_TypeInfo
    Debug_getinfo_sym *Ast_SymbolInfo
    Debug_getlocal *Ast_TypeInfo
    Debug_getlocal_sym *Ast_SymbolInfo
    Nilable_ *Ast_TypeInfo
    Nilable_val *Ast_TypeInfo
    Nilable_val_sym *Ast_SymbolInfo
    allSymbol *LnsList
    allClass *LnsList
    allFuncTypeSet *LnsSet
    allSymbolSet *LnsSet
    needThreadingTypes *LnsSet
    FP Builtin_BuiltinFuncTypeMtd
}
func Builtin_BuiltinFuncType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Builtin_BuiltinFuncType).FP
}
type Builtin_BuiltinFuncTypeDownCast interface {
    ToBuiltin_BuiltinFuncType() *Builtin_BuiltinFuncType
}
func Builtin_BuiltinFuncTypeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Builtin_BuiltinFuncTypeDownCast)
    if ok { return work.ToBuiltin_BuiltinFuncType() }
    return nil
}
func (obj *Builtin_BuiltinFuncType) ToBuiltin_BuiltinFuncType() *Builtin_BuiltinFuncType {
    return obj
}
func NewBuiltin_BuiltinFuncType() *Builtin_BuiltinFuncType {
    obj := &Builtin_BuiltinFuncType{}
    obj.FP = obj
    obj.InitBuiltin_BuiltinFuncType()
    return obj
}
func (self *Builtin_BuiltinFuncType) Get_allSymbol() *LnsList{ return self.allSymbol }
func (self *Builtin_BuiltinFuncType) Get_allClass() *LnsList{ return self.allClass }
func (self *Builtin_BuiltinFuncType) Get_allFuncTypeSet() *LnsSet{ return self.allFuncTypeSet }
func (self *Builtin_BuiltinFuncType) Get_allSymbolSet() *LnsSet{ return self.allSymbolSet }
func (self *Builtin_BuiltinFuncType) Get_needThreadingTypes() *LnsSet{ return self.needThreadingTypes }
// 135: DeclConstr
func (self *Builtin_BuiltinFuncType) InitBuiltin_BuiltinFuncType() {
    self.Lns_ = Ast_headTypeInfo
    
    self.Lns__fcall = Ast_headTypeInfo
    
    self.Lns__fcall_sym = Ast_dummySymbol
    
    self.Lns__kind = Ast_headTypeInfo
    
    self.Lns__kind_sym = Ast_dummySymbol
    
    self.Lns__load = Ast_headTypeInfo
    
    self.Lns__load_sym = Ast_dummySymbol
    
    self.Lns_collectgarbage = Ast_headTypeInfo
    
    self.Lns_collectgarbage_sym = Ast_dummySymbol
    
    self.Lns_error = Ast_headTypeInfo
    
    self.Lns_error_sym = Ast_dummySymbol
    
    self.Lns_expandLuavalMap = Ast_headTypeInfo
    
    self.Lns_expandLuavalMap_sym = Ast_dummySymbol
    
    self.Lns_loadfile = Ast_headTypeInfo
    
    self.Lns_loadfile_sym = Ast_dummySymbol
    
    self.Lns_print = Ast_headTypeInfo
    
    self.Lns_print_sym = Ast_dummySymbol
    
    self.Lns_require = Ast_headTypeInfo
    
    self.Lns_require_sym = Ast_dummySymbol
    
    self.Lns_tonumber = Ast_headTypeInfo
    
    self.Lns_tonumber_sym = Ast_dummySymbol
    
    self.Lns_tostring = Ast_headTypeInfo
    
    self.Lns_tostring_sym = Ast_dummySymbol
    
    self.Lns_type = Ast_headTypeInfo
    
    self.Lns_type_sym = Ast_dummySymbol
    
    self.Istream_ = Ast_headTypeInfo
    
    self.Istream___attrib = Ast_headTypeInfo
    
    self.Istream___attrib_sym = Ast_dummySymbol
    
    self.Istream_close = Ast_headTypeInfo
    
    self.Istream_close_sym = Ast_dummySymbol
    
    self.Istream_read = Ast_headTypeInfo
    
    self.Istream_read_sym = Ast_dummySymbol
    
    self.Ostream_ = Ast_headTypeInfo
    
    self.Ostream___attrib = Ast_headTypeInfo
    
    self.Ostream___attrib_sym = Ast_dummySymbol
    
    self.Ostream_close = Ast_headTypeInfo
    
    self.Ostream_close_sym = Ast_dummySymbol
    
    self.Ostream_flush = Ast_headTypeInfo
    
    self.Ostream_flush_sym = Ast_dummySymbol
    
    self.Ostream_write = Ast_headTypeInfo
    
    self.Ostream_write_sym = Ast_dummySymbol
    
    self.Luastream_ = Ast_headTypeInfo
    
    self.Luastream___attrib = Ast_headTypeInfo
    
    self.Luastream___attrib_sym = Ast_dummySymbol
    
    self.Luastream_close = Ast_headTypeInfo
    
    self.Luastream_close_sym = Ast_dummySymbol
    
    self.Luastream_flush = Ast_headTypeInfo
    
    self.Luastream_flush_sym = Ast_dummySymbol
    
    self.Luastream_read = Ast_headTypeInfo
    
    self.Luastream_read_sym = Ast_dummySymbol
    
    self.Luastream_seek = Ast_headTypeInfo
    
    self.Luastream_seek_sym = Ast_dummySymbol
    
    self.Luastream_write = Ast_headTypeInfo
    
    self.Luastream_write_sym = Ast_dummySymbol
    
    self.Mapping_ = Ast_headTypeInfo
    
    self.Mapping___attrib = Ast_headTypeInfo
    
    self.Mapping___attrib_sym = Ast_dummySymbol
    
    self.Mapping__toMap = Ast_headTypeInfo
    
    self.Mapping__toMap_sym = Ast_dummySymbol
    
    self.G__pipe_ = Ast_headTypeInfo
    
    self.G__pipe_get = Ast_headTypeInfo
    
    self.G__pipe_get_sym = Ast_dummySymbol
    
    self.G__pipe_put = Ast_headTypeInfo
    
    self.G__pipe_put_sym = Ast_dummySymbol
    
    self.Lnsthread_ = Ast_headTypeInfo
    
    self.Lnsthread___init = Ast_headTypeInfo
    
    self.Lnsthread___init_sym = Ast_dummySymbol
    
    self.Lnsthread_loop = Ast_headTypeInfo
    
    self.Lnsthread_loop_sym = Ast_dummySymbol
    
    self.Io_ = Ast_headTypeInfo
    
    self.Io___attrib = Ast_headTypeInfo
    
    self.Io___attrib_sym = Ast_dummySymbol
    
    self.Io_open = Ast_headTypeInfo
    
    self.Io_open_sym = Ast_dummySymbol
    
    self.Io_popen = Ast_headTypeInfo
    
    self.Io_popen_sym = Ast_dummySymbol
    
    self.Io_stderr = Ast_headTypeInfo
    
    self.Io_stderr_sym = Ast_dummySymbol
    
    self.Io_stdin = Ast_headTypeInfo
    
    self.Io_stdin_sym = Ast_dummySymbol
    
    self.Io_stdout = Ast_headTypeInfo
    
    self.Io_stdout_sym = Ast_dummySymbol
    
    self.Package_ = Ast_headTypeInfo
    
    self.Package___attrib = Ast_headTypeInfo
    
    self.Package___attrib_sym = Ast_dummySymbol
    
    self.Package_path = Ast_headTypeInfo
    
    self.Package_path_sym = Ast_dummySymbol
    
    self.Package_searchpath = Ast_headTypeInfo
    
    self.Package_searchpath_sym = Ast_dummySymbol
    
    self.Os_ = Ast_headTypeInfo
    
    self.Os___attrib = Ast_headTypeInfo
    
    self.Os___attrib_sym = Ast_dummySymbol
    
    self.Os_clock = Ast_headTypeInfo
    
    self.Os_clock_sym = Ast_dummySymbol
    
    self.Os_date = Ast_headTypeInfo
    
    self.Os_date_sym = Ast_dummySymbol
    
    self.Os_difftime = Ast_headTypeInfo
    
    self.Os_difftime_sym = Ast_dummySymbol
    
    self.Os_exit = Ast_headTypeInfo
    
    self.Os_exit_sym = Ast_dummySymbol
    
    self.Os_remove = Ast_headTypeInfo
    
    self.Os_remove_sym = Ast_dummySymbol
    
    self.Os_rename = Ast_headTypeInfo
    
    self.Os_rename_sym = Ast_dummySymbol
    
    self.Os_time = Ast_headTypeInfo
    
    self.Os_time_sym = Ast_dummySymbol
    
    self.String_ = Ast_headTypeInfo
    
    self.String___attrib = Ast_headTypeInfo
    
    self.String___attrib_sym = Ast_dummySymbol
    
    self.String_byte = Ast_headTypeInfo
    
    self.String_byte_sym = Ast_dummySymbol
    
    self.String_dump = Ast_headTypeInfo
    
    self.String_dump_sym = Ast_dummySymbol
    
    self.String_find = Ast_headTypeInfo
    
    self.String_find_sym = Ast_dummySymbol
    
    self.String_format = Ast_headTypeInfo
    
    self.String_format_sym = Ast_dummySymbol
    
    self.String_gmatch = Ast_headTypeInfo
    
    self.String_gmatch_sym = Ast_dummySymbol
    
    self.String_gsub = Ast_headTypeInfo
    
    self.String_gsub_sym = Ast_dummySymbol
    
    self.String_lower = Ast_headTypeInfo
    
    self.String_lower_sym = Ast_dummySymbol
    
    self.String_rep = Ast_headTypeInfo
    
    self.String_rep_sym = Ast_dummySymbol
    
    self.String_reverse = Ast_headTypeInfo
    
    self.String_reverse_sym = Ast_dummySymbol
    
    self.String_sub = Ast_headTypeInfo
    
    self.String_sub_sym = Ast_dummySymbol
    
    self.String_upper = Ast_headTypeInfo
    
    self.String_upper_sym = Ast_dummySymbol
    
    self.Str_ = Ast_headTypeInfo
    
    self.Str___attrib = Ast_headTypeInfo
    
    self.Str___attrib_sym = Ast_dummySymbol
    
    self.Str_byte = Ast_headTypeInfo
    
    self.Str_byte_sym = Ast_dummySymbol
    
    self.Str_find = Ast_headTypeInfo
    
    self.Str_find_sym = Ast_dummySymbol
    
    self.Str_format = Ast_headTypeInfo
    
    self.Str_format_sym = Ast_dummySymbol
    
    self.Str_gmatch = Ast_headTypeInfo
    
    self.Str_gmatch_sym = Ast_dummySymbol
    
    self.Str_gsub = Ast_headTypeInfo
    
    self.Str_gsub_sym = Ast_dummySymbol
    
    self.Str_lower = Ast_headTypeInfo
    
    self.Str_lower_sym = Ast_dummySymbol
    
    self.Str_rep = Ast_headTypeInfo
    
    self.Str_rep_sym = Ast_dummySymbol
    
    self.Str_reverse = Ast_headTypeInfo
    
    self.Str_reverse_sym = Ast_dummySymbol
    
    self.Str_sub = Ast_headTypeInfo
    
    self.Str_sub_sym = Ast_dummySymbol
    
    self.Str_upper = Ast_headTypeInfo
    
    self.Str_upper_sym = Ast_dummySymbol
    
    self.List_ = Ast_headTypeInfo
    
    self.List___less = Ast_headTypeInfo
    
    self.List___less_sym = Ast_dummySymbol
    
    self.List_insert = Ast_headTypeInfo
    
    self.List_insert_sym = Ast_dummySymbol
    
    self.List_remove = Ast_headTypeInfo
    
    self.List_remove_sym = Ast_dummySymbol
    
    self.List_sort = Ast_headTypeInfo
    
    self.List_sort_sym = Ast_dummySymbol
    
    self.List_unpack = Ast_headTypeInfo
    
    self.List_unpack_sym = Ast_dummySymbol
    
    self.Array_ = Ast_headTypeInfo
    
    self.Array___less = Ast_headTypeInfo
    
    self.Array___less_sym = Ast_dummySymbol
    
    self.Array_sort = Ast_headTypeInfo
    
    self.Array_sort_sym = Ast_dummySymbol
    
    self.Array_unpack = Ast_headTypeInfo
    
    self.Array_unpack_sym = Ast_dummySymbol
    
    self.Set_ = Ast_headTypeInfo
    
    self.Set_add = Ast_headTypeInfo
    
    self.Set_add_sym = Ast_dummySymbol
    
    self.Set_and = Ast_headTypeInfo
    
    self.Set_and_sym = Ast_dummySymbol
    
    self.Set_clone = Ast_headTypeInfo
    
    self.Set_clone_sym = Ast_dummySymbol
    
    self.Set_del = Ast_headTypeInfo
    
    self.Set_del_sym = Ast_dummySymbol
    
    self.Set_has = Ast_headTypeInfo
    
    self.Set_has_sym = Ast_dummySymbol
    
    self.Set_len = Ast_headTypeInfo
    
    self.Set_len_sym = Ast_dummySymbol
    
    self.Set_or = Ast_headTypeInfo
    
    self.Set_or_sym = Ast_dummySymbol
    
    self.Set_sub = Ast_headTypeInfo
    
    self.Set_sub_sym = Ast_dummySymbol
    
    self.Math_ = Ast_headTypeInfo
    
    self.Math___attrib = Ast_headTypeInfo
    
    self.Math___attrib_sym = Ast_dummySymbol
    
    self.Math_random = Ast_headTypeInfo
    
    self.Math_random_sym = Ast_dummySymbol
    
    self.Math_randomseed = Ast_headTypeInfo
    
    self.Math_randomseed_sym = Ast_dummySymbol
    
    self.Debug_ = Ast_headTypeInfo
    
    self.Debug___attrib = Ast_headTypeInfo
    
    self.Debug___attrib_sym = Ast_dummySymbol
    
    self.Debug_getinfo = Ast_headTypeInfo
    
    self.Debug_getinfo_sym = Ast_dummySymbol
    
    self.Debug_getlocal = Ast_headTypeInfo
    
    self.Debug_getlocal_sym = Ast_dummySymbol
    
    self.Nilable_ = Ast_headTypeInfo
    
    self.Nilable_val = Ast_headTypeInfo
    
    self.Nilable_val_sym = Ast_dummySymbol
    
    self.allSymbol = NewLnsList([]LnsAny{})
    
    self.allClass = NewLnsList([]LnsAny{})
    
    self.allSymbolSet = NewLnsSet([]LnsAny{})
    
    self.allFuncTypeSet = NewLnsSet([]LnsAny{})
    
    self.needThreadingTypes = NewLnsSet([]LnsAny{})
    
}

// 145: decl @lune.@base.@Builtin.BuiltinFuncType.register
func (self *Builtin_BuiltinFuncType) Register(symbolInfo *Ast_SymbolInfo) {
    self.allSymbol.Insert(Ast_SymbolInfo2Stem(symbolInfo))
    self.allSymbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
}

// 149: decl @lune.@base.@Builtin.BuiltinFuncType.registerClass
func (self *Builtin_BuiltinFuncType) RegisterClass(classInfo *Ast_TypeInfo) {
    self.allClass.Insert(Ast_TypeInfo2Stem(classInfo))
    if _switch5925 := classInfo.FP.Get_rawTxt(); _switch5925 == "" {
        self.Lns_ = classInfo
        
    } else if _switch5925 == "iStream" {
        self.Istream_ = classInfo
        
    } else if _switch5925 == "oStream" {
        self.Ostream_ = classInfo
        
    } else if _switch5925 == "luaStream" {
        self.Luastream_ = classInfo
        
    } else if _switch5925 == "Mapping" {
        self.Mapping_ = classInfo
        
    } else if _switch5925 == "__pipe" {
        self.G__pipe_ = classInfo
        
    } else if _switch5925 == "LnsThread" {
        self.Lnsthread_ = classInfo
        
    } else if _switch5925 == "io" {
        self.Io_ = classInfo
        
    } else if _switch5925 == "package" {
        self.Package_ = classInfo
        
    } else if _switch5925 == "os" {
        self.Os_ = classInfo
        
    } else if _switch5925 == "string" {
        self.String_ = classInfo
        
    } else if _switch5925 == "str" {
        self.Str_ = classInfo
        
    } else if _switch5925 == "List" {
        self.List_ = classInfo
        
    } else if _switch5925 == "Array" {
        self.Array_ = classInfo
        
    } else if _switch5925 == "Set" {
        self.Set_ = classInfo
        
    } else if _switch5925 == "math" {
        self.Math_ = classInfo
        
    } else if _switch5925 == "debug" {
        self.Debug_ = classInfo
        
    } else if _switch5925 == "Nilable" {
        self.Nilable_ = classInfo
        
    }
}


func Lns_Builtin_init() {
    if init_Builtin { return }
    init_Builtin = true
    Builtin__mod__ = "@lune.@base.@Builtin"
    Lns_InitMod()
    Lns_Types_init()
    Lns_Parser_init()
    Lns_Util_init()
    Lns_Ast_init()
    Lns_LuaVer_init()
    Lns_Log_init()
    Lns_TransUnitIF_init()
    Builtin_builtinFunc = NewBuiltin_BuiltinFuncType()
    
    Builtin_readyBuiltin = false
}
func init() {
    init_Builtin = false
}
