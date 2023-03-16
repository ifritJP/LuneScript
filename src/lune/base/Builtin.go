// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Builtin bool
var Builtin__mod__ string
var Builtin_builtinFunc *Builtin_BuiltinFuncType
var Builtin_readyBuiltin bool
// for 613
func Builtin_convExp0_15591(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 605
func Builtin_convExp0_15472(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 633
func Builtin_convExp0_15657(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 75
func Builtin_convExp0_217(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 93
func Builtin_convExp0_340(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 136
func Builtin_convExp0_626(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 608
func Builtin_convExp0_15495(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1047
func Builtin_convExp0_17702(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 779
func Builtin_convExp0_16388(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}
// for 907
func Builtin_convExp0_17075(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}



// 956: decl @lune.@base.@Builtin.Builtin.registBuiltInScope.processCopyAlterList
func Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env *LnsEnv, alterList *LnsList2_[*Ast_AlternateTypeInfo],typeList *LnsList2_[*Ast_TypeInfo]) {
    for _, _typeInfo := range( typeList.Items ) {
        typeInfo := _typeInfo
        alterList.Insert(Lns_unwrap( Ast_AlternateTypeInfoDownCastF(typeInfo.FP)).(*Ast_AlternateTypeInfo))
    }
}

// 206: decl @lune.@base.@Builtin.getBuiltinFunc
func Builtin_getBuiltinFunc(_env *LnsEnv) *Builtin_BuiltinFuncType {
    return Builtin_builtinFunc
}

// 210: decl @lune.@base.@Builtin.setupBuiltinTypeInfo
func Builtin_setupBuiltinTypeInfo_4_(_env *LnsEnv, name string,fieldName string,symInfo *Ast_SymbolInfo) {
    var typeInfo *Ast_TypeInfo
    typeInfo = symInfo.FP.Get_typeInfo(_env)
    var Builtin_process___Er func(_env *LnsEnv)
    Builtin_process___Er = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "get_txt" {
            Builtin_builtinFunc.G__er_get_txt = typeInfo
            Builtin_builtinFunc.G__er_get_txt_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_ func(_env *LnsEnv)
    Builtin_process_ = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "__join" {
            Builtin_builtinFunc.Lns___join = typeInfo
            Builtin_builtinFunc.Lns___join_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "__run" {
            Builtin_builtinFunc.Lns___run = typeInfo
            Builtin_builtinFunc.Lns___run_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "__serr" {
            Builtin_builtinFunc.Lns___serr = typeInfo
            Builtin_builtinFunc.Lns___serr_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "_fcall" {
            Builtin_builtinFunc.Lns__fcall = typeInfo
            Builtin_builtinFunc.Lns__fcall_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "_kind" {
            Builtin_builtinFunc.Lns__kind = typeInfo
            Builtin_builtinFunc.Lns__kind_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "_load" {
            Builtin_builtinFunc.Lns__load = typeInfo
            Builtin_builtinFunc.Lns__load_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "collectgarbage" {
            Builtin_builtinFunc.Lns_collectgarbage = typeInfo
            Builtin_builtinFunc.Lns_collectgarbage_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "error" {
            Builtin_builtinFunc.Lns_error = typeInfo
            Builtin_builtinFunc.Lns_error_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "expandLuavalMap" {
            Builtin_builtinFunc.Lns_expandLuavalMap = typeInfo
            Builtin_builtinFunc.Lns_expandLuavalMap_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "loadfile" {
            Builtin_builtinFunc.Lns_loadfile = typeInfo
            Builtin_builtinFunc.Lns_loadfile_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "print" {
            Builtin_builtinFunc.Lns_print = typeInfo
            Builtin_builtinFunc.Lns_print_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "require" {
            Builtin_builtinFunc.Lns_require = typeInfo
            Builtin_builtinFunc.Lns_require_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "tonumber" {
            Builtin_builtinFunc.Lns_tonumber = typeInfo
            Builtin_builtinFunc.Lns_tonumber_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "tostring" {
            Builtin_builtinFunc.Lns_tostring = typeInfo
            Builtin_builtinFunc.Lns_tostring_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "type" {
            Builtin_builtinFunc.Lns_type = typeInfo
            Builtin_builtinFunc.Lns_type_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_iStream func(_env *LnsEnv)
    Builtin_process_iStream = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "close" {
            Builtin_builtinFunc.Istream_close = typeInfo
            Builtin_builtinFunc.Istream_close_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "read" {
            Builtin_builtinFunc.Istream_read = typeInfo
            Builtin_builtinFunc.Istream_read_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_oStream func(_env *LnsEnv)
    Builtin_process_oStream = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "close" {
            Builtin_builtinFunc.Ostream_close = typeInfo
            Builtin_builtinFunc.Ostream_close_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "flush" {
            Builtin_builtinFunc.Ostream_flush = typeInfo
            Builtin_builtinFunc.Ostream_flush_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "write" {
            Builtin_builtinFunc.Ostream_write = typeInfo
            Builtin_builtinFunc.Ostream_write_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___pipe func(_env *LnsEnv)
    Builtin_process___pipe = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "get" {
            Builtin_builtinFunc.G__pipe_get = typeInfo
            Builtin_builtinFunc.G__pipe_get_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "put" {
            Builtin_builtinFunc.G__pipe_put = typeInfo
            Builtin_builtinFunc.G__pipe_put_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___lns_runMode func(_env *LnsEnv)
    Builtin_process___lns_runMode = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "Queue" {
            Builtin_builtinFunc.G__lns_runmode_Queue = typeInfo
            Builtin_builtinFunc.G__lns_runmode_Queue_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "Skip" {
            Builtin_builtinFunc.G__lns_runmode_Skip = typeInfo
            Builtin_builtinFunc.G__lns_runmode_Skip_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "Sync" {
            Builtin_builtinFunc.G__lns_runmode_Sync = typeInfo
            Builtin_builtinFunc.G__lns_runmode_Sync_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___lns_runtime func(_env *LnsEnv)
    Builtin_process___lns_runtime = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "dumpLog" {
            Builtin_builtinFunc.G__lns_runtime_dumpLog = typeInfo
            Builtin_builtinFunc.G__lns_runtime_dumpLog_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "enableDebugLog" {
            Builtin_builtinFunc.G__lns_runtime_enableDebugLog = typeInfo
            Builtin_builtinFunc.G__lns_runtime_enableDebugLog_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "enableLog" {
            Builtin_builtinFunc.G__lns_runtime_enableLog = typeInfo
            Builtin_builtinFunc.G__lns_runtime_enableLog_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "log" {
            Builtin_builtinFunc.G__lns_runtime_log = typeInfo
            Builtin_builtinFunc.G__lns_runtime_log_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___lns_capability func(_env *LnsEnv)
    Builtin_process___lns_capability = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "async" {
            Builtin_builtinFunc.G__lns_capability_async = typeInfo
            Builtin_builtinFunc.G__lns_capability_async_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___lns_Sync_Flag func(_env *LnsEnv)
    Builtin_process___lns_Sync_Flag = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "set" {
            Builtin_builtinFunc.G__lns_sync_flag_set = typeInfo
            Builtin_builtinFunc.G__lns_sync_flag_set_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "wait" {
            Builtin_builtinFunc.G__lns_sync_flag_wait = typeInfo
            Builtin_builtinFunc.G__lns_sync_flag_wait_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___lns_Sync func(_env *LnsEnv)
    Builtin_process___lns_Sync = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "_createPipe" {
            Builtin_builtinFunc.G__lns_sync__createPipe = typeInfo
            Builtin_builtinFunc.G__lns_sync__createPipe_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "createFlag" {
            Builtin_builtinFunc.G__lns_sync_createFlag = typeInfo
            Builtin_builtinFunc.G__lns_sync_createFlag_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "createProcesser" {
            Builtin_builtinFunc.G__lns_sync_createProcesser = typeInfo
            Builtin_builtinFunc.G__lns_sync_createProcesser_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_luaStream func(_env *LnsEnv)
    Builtin_process_luaStream = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "close" {
            Builtin_builtinFunc.Luastream_close = typeInfo
            Builtin_builtinFunc.Luastream_close_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "flush" {
            Builtin_builtinFunc.Luastream_flush = typeInfo
            Builtin_builtinFunc.Luastream_flush_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "read" {
            Builtin_builtinFunc.Luastream_read = typeInfo
            Builtin_builtinFunc.Luastream_read_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "seek" {
            Builtin_builtinFunc.Luastream_seek = typeInfo
            Builtin_builtinFunc.Luastream_seek_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "write" {
            Builtin_builtinFunc.Luastream_write = typeInfo
            Builtin_builtinFunc.Luastream_write_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_Mapping func(_env *LnsEnv)
    Builtin_process_Mapping = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "_toMap" {
            Builtin_builtinFunc.Mapping__toMap = typeInfo
            Builtin_builtinFunc.Mapping__toMap_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___Runner func(_env *LnsEnv)
    Builtin_process___Runner = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "run" {
            Builtin_builtinFunc.G__runner_run = typeInfo
            Builtin_builtinFunc.G__runner_run_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___Processor func(_env *LnsEnv)
    Builtin_process___Processor = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "end" {
            Builtin_builtinFunc.G__processor_end = typeInfo
            Builtin_builtinFunc.G__processor_end_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_io func(_env *LnsEnv)
    Builtin_process_io = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "open" {
            Builtin_builtinFunc.Io_open = typeInfo
            Builtin_builtinFunc.Io_open_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "popen" {
            Builtin_builtinFunc.Io_popen = typeInfo
            Builtin_builtinFunc.Io_popen_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "stderr" {
            Builtin_builtinFunc.Io_stderr = typeInfo
            Builtin_builtinFunc.Io_stderr_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "stdin" {
            Builtin_builtinFunc.Io_stdin = typeInfo
            Builtin_builtinFunc.Io_stdin_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "stdout" {
            Builtin_builtinFunc.Io_stdout = typeInfo
            Builtin_builtinFunc.Io_stdout_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_package func(_env *LnsEnv)
    Builtin_process_package = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "path" {
            Builtin_builtinFunc.Package_path = typeInfo
            Builtin_builtinFunc.Package_path_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "searchpath" {
            Builtin_builtinFunc.Package_searchpath = typeInfo
            Builtin_builtinFunc.Package_searchpath_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_os func(_env *LnsEnv)
    Builtin_process_os = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "clock" {
            Builtin_builtinFunc.Os_clock = typeInfo
            Builtin_builtinFunc.Os_clock_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "date" {
            Builtin_builtinFunc.Os_date = typeInfo
            Builtin_builtinFunc.Os_date_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "difftime" {
            Builtin_builtinFunc.Os_difftime = typeInfo
            Builtin_builtinFunc.Os_difftime_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "exit" {
            Builtin_builtinFunc.Os_exit = typeInfo
            Builtin_builtinFunc.Os_exit_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "remove" {
            Builtin_builtinFunc.Os_remove = typeInfo
            Builtin_builtinFunc.Os_remove_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "rename" {
            Builtin_builtinFunc.Os_rename = typeInfo
            Builtin_builtinFunc.Os_rename_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "time" {
            Builtin_builtinFunc.Os_time = typeInfo
            Builtin_builtinFunc.Os_time_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_string func(_env *LnsEnv)
    Builtin_process_string = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "byte" {
            Builtin_builtinFunc.String_byte = typeInfo
            Builtin_builtinFunc.String_byte_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "dump" {
            Builtin_builtinFunc.String_dump = typeInfo
            Builtin_builtinFunc.String_dump_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "find" {
            Builtin_builtinFunc.String_find = typeInfo
            Builtin_builtinFunc.String_find_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "format" {
            Builtin_builtinFunc.String_format = typeInfo
            Builtin_builtinFunc.String_format_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "gmatch" {
            Builtin_builtinFunc.String_gmatch = typeInfo
            Builtin_builtinFunc.String_gmatch_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "gsub" {
            Builtin_builtinFunc.String_gsub = typeInfo
            Builtin_builtinFunc.String_gsub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "lower" {
            Builtin_builtinFunc.String_lower = typeInfo
            Builtin_builtinFunc.String_lower_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "rep" {
            Builtin_builtinFunc.String_rep = typeInfo
            Builtin_builtinFunc.String_rep_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "reverse" {
            Builtin_builtinFunc.String_reverse = typeInfo
            Builtin_builtinFunc.String_reverse_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sub" {
            Builtin_builtinFunc.String_sub = typeInfo
            Builtin_builtinFunc.String_sub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "upper" {
            Builtin_builtinFunc.String_upper = typeInfo
            Builtin_builtinFunc.String_upper_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_str func(_env *LnsEnv)
    Builtin_process_str = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "byte" {
            Builtin_builtinFunc.Str_byte = typeInfo
            Builtin_builtinFunc.Str_byte_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "find" {
            Builtin_builtinFunc.Str_find = typeInfo
            Builtin_builtinFunc.Str_find_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "format" {
            Builtin_builtinFunc.Str_format = typeInfo
            Builtin_builtinFunc.Str_format_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "gmatch" {
            Builtin_builtinFunc.Str_gmatch = typeInfo
            Builtin_builtinFunc.Str_gmatch_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "gsub" {
            Builtin_builtinFunc.Str_gsub = typeInfo
            Builtin_builtinFunc.Str_gsub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "lower" {
            Builtin_builtinFunc.Str_lower = typeInfo
            Builtin_builtinFunc.Str_lower_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "rep" {
            Builtin_builtinFunc.Str_rep = typeInfo
            Builtin_builtinFunc.Str_rep_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "replace" {
            Builtin_builtinFunc.Str_replace = typeInfo
            Builtin_builtinFunc.Str_replace_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "reverse" {
            Builtin_builtinFunc.Str_reverse = typeInfo
            Builtin_builtinFunc.Str_reverse_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sub" {
            Builtin_builtinFunc.Str_sub = typeInfo
            Builtin_builtinFunc.Str_sub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "upper" {
            Builtin_builtinFunc.Str_upper = typeInfo
            Builtin_builtinFunc.Str_upper_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process__List func(_env *LnsEnv)
    Builtin_process__List = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "__less" {
            Builtin_builtinFunc.G_list___less = typeInfo
            Builtin_builtinFunc.G_list___less_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "insert" {
            Builtin_builtinFunc.G_list_insert = typeInfo
            Builtin_builtinFunc.G_list_insert_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "remove" {
            Builtin_builtinFunc.G_list_remove = typeInfo
            Builtin_builtinFunc.G_list_remove_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sort" {
            Builtin_builtinFunc.G_list_sort = typeInfo
            Builtin_builtinFunc.G_list_sort_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "unpack" {
            Builtin_builtinFunc.G_list_unpack = typeInfo
            Builtin_builtinFunc.G_list_unpack_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___List func(_env *LnsEnv)
    Builtin_process___List = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "__less" {
            Builtin_builtinFunc.G__list___less = typeInfo
            Builtin_builtinFunc.G__list___less_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "insert" {
            Builtin_builtinFunc.G__list_insert = typeInfo
            Builtin_builtinFunc.G__list_insert_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "remove" {
            Builtin_builtinFunc.G__list_remove = typeInfo
            Builtin_builtinFunc.G__list_remove_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sort" {
            Builtin_builtinFunc.G__list_sort = typeInfo
            Builtin_builtinFunc.G__list_sort_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "unpack" {
            Builtin_builtinFunc.G__list_unpack = typeInfo
            Builtin_builtinFunc.G__list_unpack_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_Array func(_env *LnsEnv)
    Builtin_process_Array = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "__less" {
            Builtin_builtinFunc.Array___less = typeInfo
            Builtin_builtinFunc.Array___less_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sort" {
            Builtin_builtinFunc.Array_sort = typeInfo
            Builtin_builtinFunc.Array_sort_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "unpack" {
            Builtin_builtinFunc.Array_unpack = typeInfo
            Builtin_builtinFunc.Array_unpack_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process__Set func(_env *LnsEnv)
    Builtin_process__Set = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "add" {
            Builtin_builtinFunc.G_set_add = typeInfo
            Builtin_builtinFunc.G_set_add_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "and" {
            Builtin_builtinFunc.G_set_and = typeInfo
            Builtin_builtinFunc.G_set_and_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "clone" {
            Builtin_builtinFunc.G_set_clone = typeInfo
            Builtin_builtinFunc.G_set_clone_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "del" {
            Builtin_builtinFunc.G_set_del = typeInfo
            Builtin_builtinFunc.G_set_del_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "has" {
            Builtin_builtinFunc.G_set_has = typeInfo
            Builtin_builtinFunc.G_set_has_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "len" {
            Builtin_builtinFunc.G_set_len = typeInfo
            Builtin_builtinFunc.G_set_len_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "or" {
            Builtin_builtinFunc.G_set_or = typeInfo
            Builtin_builtinFunc.G_set_or_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sub" {
            Builtin_builtinFunc.G_set_sub = typeInfo
            Builtin_builtinFunc.G_set_sub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process___Set func(_env *LnsEnv)
    Builtin_process___Set = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "add" {
            Builtin_builtinFunc.G__set_add = typeInfo
            Builtin_builtinFunc.G__set_add_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "and" {
            Builtin_builtinFunc.G__set_and = typeInfo
            Builtin_builtinFunc.G__set_and_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "clone" {
            Builtin_builtinFunc.G__set_clone = typeInfo
            Builtin_builtinFunc.G__set_clone_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "del" {
            Builtin_builtinFunc.G__set_del = typeInfo
            Builtin_builtinFunc.G__set_del_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "has" {
            Builtin_builtinFunc.G__set_has = typeInfo
            Builtin_builtinFunc.G__set_has_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "len" {
            Builtin_builtinFunc.G__set_len = typeInfo
            Builtin_builtinFunc.G__set_len_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "or" {
            Builtin_builtinFunc.G__set_or = typeInfo
            Builtin_builtinFunc.G__set_or_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sub" {
            Builtin_builtinFunc.G__set_sub = typeInfo
            Builtin_builtinFunc.G__set_sub_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_math func(_env *LnsEnv)
    Builtin_process_math = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "random" {
            Builtin_builtinFunc.Math_random = typeInfo
            Builtin_builtinFunc.Math_random_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "randomseed" {
            Builtin_builtinFunc.Math_randomseed = typeInfo
            Builtin_builtinFunc.Math_randomseed_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_debug func(_env *LnsEnv)
    Builtin_process_debug = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "getinfo" {
            Builtin_builtinFunc.Debug_getinfo = typeInfo
            Builtin_builtinFunc.Debug_getinfo_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "getlocal" {
            Builtin_builtinFunc.Debug_getlocal = typeInfo
            Builtin_builtinFunc.Debug_getlocal_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    var Builtin_process_Nilable func(_env *LnsEnv)
    Builtin_process_Nilable = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "val" {
            Builtin_builtinFunc.Nilable_val = typeInfo
            Builtin_builtinFunc.Nilable_val_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        }
    }
    if _switch0 := name; _switch0 == "__Er" {
        Builtin_process___Er(_env)
    } else if _switch0 == "" {
        Builtin_process_(_env)
    } else if _switch0 == "iStream" {
        Builtin_process_iStream(_env)
    } else if _switch0 == "oStream" {
        Builtin_process_oStream(_env)
    } else if _switch0 == "__pipe" {
        Builtin_process___pipe(_env)
    } else if _switch0 == "__lns_runMode" {
        Builtin_process___lns_runMode(_env)
    } else if _switch0 == "__lns_runtime" {
        Builtin_process___lns_runtime(_env)
    } else if _switch0 == "__lns_capability" {
        Builtin_process___lns_capability(_env)
    } else if _switch0 == "__lns_Sync_Flag" {
        Builtin_process___lns_Sync_Flag(_env)
    } else if _switch0 == "__lns_Sync" {
        Builtin_process___lns_Sync(_env)
    } else if _switch0 == "luaStream" {
        Builtin_process_luaStream(_env)
    } else if _switch0 == "Mapping" {
        Builtin_process_Mapping(_env)
    } else if _switch0 == "__Runner" {
        Builtin_process___Runner(_env)
    } else if _switch0 == "__Processor" {
        Builtin_process___Processor(_env)
    } else if _switch0 == "io" {
        Builtin_process_io(_env)
    } else if _switch0 == "package" {
        Builtin_process_package(_env)
    } else if _switch0 == "os" {
        Builtin_process_os(_env)
    } else if _switch0 == "string" {
        Builtin_process_string(_env)
    } else if _switch0 == "str" {
        Builtin_process_str(_env)
    } else if _switch0 == "_List" {
        Builtin_process__List(_env)
    } else if _switch0 == "__List" {
        Builtin_process___List(_env)
    } else if _switch0 == "Array" {
        Builtin_process_Array(_env)
    } else if _switch0 == "_Set" {
        Builtin_process__Set(_env)
    } else if _switch0 == "__Set" {
        Builtin_process___Set(_env)
    } else if _switch0 == "math" {
        Builtin_process_math(_env)
    } else if _switch0 == "debug" {
        Builtin_process_debug(_env)
    } else if _switch0 == "Nilable" {
        Builtin_process_Nilable(_env)
    }
}

// 220: decl @lune.@base.@Builtin.getBuiltInInfo
func Builtin_getBuiltInInfo_5_(_env *LnsEnv) *LnsList2_[*LnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]] {
    return NewLnsList2_[*LnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]](Lns_2DDDGen[*LnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]](NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__Er":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"get_txt":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__join":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&__Runner")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("")),}),"__run":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("__Runner", "int", "str!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),}),"__serr":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Er")),}),"_fcall":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("form", "&...")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("")),}),"_kind":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("int")),}),"_load":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<form>!", "str!")),}),"collectgarbage":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),}),"error":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__")),}),"expandLuavalMap":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<&stem>!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("&stem!")),}),"loadfile":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<form>!", "str!")),}),"print":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&...")),"ret":NewLnsList2_[string]([]string{}),}),"require":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<stem>")),}),"tonumber":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("real!")),}),"tostring":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&stem")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"type":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"iStream":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"close":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"read":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"oStream":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"close":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"flush":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"write":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("stem!", "str!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__pipe<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"get":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("T!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"put":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T!")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__lns.runMode":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"Queue":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("int")),}),"Skip":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("int")),}),"Sync":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("int")),}),"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("class")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__lns.runtime":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"dumpLog":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("oStream")),"ret":NewLnsList2_[string]([]string{}),}),"enableDebugLog":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"ret":NewLnsList2_[string]([]string{}),}),"enableLog":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"ret":NewLnsList2_[string]([]string{}),}),"log":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string]([]string{}),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__lns.capability":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"async":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__lns.Sync.Flag":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"set":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"wait":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__lns.Sync":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("class")),}),"_createPipe":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("__exp", "int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__pipe!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("macro")),}),"createFlag":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__lns.Sync.Flag!")),}),"createProcesser":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Processor")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"luaStream":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"implements":NewLnsList2_[string](Lns_2DDDGen[string]("iStream", "oStream")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"close":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"flush":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"read":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"seek":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("int!", "str!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"write":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("stem!", "str!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"Mapping":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"_toMap":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("&Map<str,&stem>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__Runner":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"run":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__Processor":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"implements":NewLnsList2_[string](Lns_2DDDGen[string]("__Runner")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("interface")),}),"end":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"io":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"open":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("luaStream!", "str!")),}),"popen":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<luaStream>!")),}),"stderr":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("oStream")),}),"stdin":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("iStream")),}),"stdout":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("oStream")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"package":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"path":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("var")),"typeInfo":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"searchpath":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str!")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"os":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"clock":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("real")),}),"date":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str!", "stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<stem>!")),}),"difftime":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&stem", "&stem")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("int")),}),"exit":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__")),}),"remove":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool!", "str!")),}),"rename":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("stem!", "str!")),}),"time":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("stem!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("stem!")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"string":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"byte":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int!", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...<int>")),}),"dump":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&Luaval<form>", "bool!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"find":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str", "int!", "bool!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...<int>")),}),"format":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "&...")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"gmatch":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<form>", "stem!", "stem!")),}),"gsub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int")),}),"lower":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"rep":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"reverse":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"sub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),"upper":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"str":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"implements":NewLnsList2_[string](Lns_2DDDGen[string]("Mapping")),}),"byte":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int!", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...<int!>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"find":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int!", "bool!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...<int>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"format":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&...")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"gmatch":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Luaval<form>", "stem!", "stem!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"gsub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str", "int")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"lower":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"rep":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"replace":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("str", "str")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"reverse":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"sub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"upper":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"_List<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__less":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T", "T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("formfunc")),}),"insert":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"remove":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("T!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"sort":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("__less!")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"unpack":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__List<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__less":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T", "T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("formfunc")),}),"insert":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"remove":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("T!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"sort":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("__less!")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"unpack":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"Array<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__less":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T", "T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("formfunc")),}),"sort":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("__less!")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"unpack":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("...")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"_Set<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"add":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"and":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&_Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("_Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"clone":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("_Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"del":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"has":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"len":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("int")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"or":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&_Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("_Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"sub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&_Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("_Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"__Set<T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"add":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"and":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&__Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"clone":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"del":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string]([]string{}),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"has":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("T")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("bool")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"len":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("int")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),"or":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&__Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),"sub":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("&__Set<T>")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("__Set<T>")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("mut")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"math":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"random":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int!", "int!")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("real")),}),"randomseed":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int")),"ret":NewLnsList2_[string]([]string{}),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"debug":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"__attrib":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"type":NewLnsList2_[string](Lns_2DDDGen[string]("module")),}),"getinfo":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("Map<str,stem>!")),}),"getlocal":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string](Lns_2DDDGen[string]("int", "int")),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("str!", "stem!")),}),}),}), NewLnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]{"Nilable<_T>":NewLnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]( map[string]*LnsMap2_[string,*LnsList2_[string]]{"val":NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"arg":NewLnsList2_[string]([]string{}),"ret":NewLnsList2_[string](Lns_2DDDGen[string]("_T!")),"type":NewLnsList2_[string](Lns_2DDDGen[string]("method")),}),}),})))
}

// 223: decl @lune.@base.@Builtin.getBuiltinAlgeInfo
func Builtin_getBuiltinAlgeInfo_6_(_env *LnsEnv) *LnsList2_[*LnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]] {
    return NewLnsList2_[*LnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]](Lns_2DDDGen[*LnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]](NewLnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]( map[string]*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]{"__Ret<T1,T2>":NewLnsList2_[*LnsMap2_[string,*LnsList2_[string]]](Lns_2DDDGen[*LnsMap2_[string,*LnsList2_[string]]](NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"Ok":NewLnsList2_[string](Lns_2DDDGen[string]("T1")),}), NewLnsMap2_[string,*LnsList2_[string]]( map[string]*LnsList2_[string]{"Err":NewLnsList2_[string](Lns_2DDDGen[string]("T2")),}))),})))
}




























// 500: decl @lune.@base.@Builtin.Builtin.getTypeInfo
func (self *Builtin_Builtin) getTypeInfo(_env *LnsEnv, typeName string,targetScope LnsAny) *Ast_TypeInfo {
    if _switch0 := typeName; _switch0 == "_T" {
        return Ast_builtinTypeBox.FP.Get_boxingType(_env)
    } else if _switch0 == "_T!" {
        return Ast_builtinTypeBox.FP.Get_boxingType(_env).FP.Get_nilableTypeInfo(_env)
    }
    var Builtin_getTypeInfoFromScope func(_env *LnsEnv, scope *Ast_Scope,symbol string,genTypeList *LnsList2_[*Ast_TypeInfo]) LnsAny
    Builtin_getTypeInfoFromScope = func(_env *LnsEnv, scope *Ast_Scope,symbol string,genTypeList *LnsList2_[*Ast_TypeInfo]) LnsAny {
        var Builtin_getTypeInfo func(_env *LnsEnv, workScope *Ast_Scope) LnsAny
        Builtin_getTypeInfo = func(_env *LnsEnv, workScope *Ast_Scope) LnsAny {
            var nameList *LnsList2_[string]
            nameList = Util_splitModule(_env, symbol)
            if nameList.Len() <= 1{
                return workScope.FP.GetTypeInfo(_env, symbol, workScope, false, Ast_ScopeAccess__Normal)
            } else { 
                for _index, _name := range( nameList.Items ) {
                    index := _index + 1
                    name := _name
                    var workTypeInfo LnsAny
                    if index == 1{
                        workTypeInfo = workScope.FP.GetTypeInfo(_env, name, workScope, false, Ast_ScopeAccess__Normal)
                    } else { 
                        workTypeInfo = workScope.FP.GetTypeInfoChild(_env, name)
                    }
                    if workTypeInfo != nil{
                        workTypeInfo_691 := workTypeInfo.(*Ast_TypeInfo)
                        if index == nameList.Len(){
                            return workTypeInfo_691
                        }
                        {
                            __exp := workTypeInfo_691.FP.Get_scope(_env)
                            if !Lns_IsNil( __exp ) {
                                _exp := __exp.(*Ast_Scope)
                                workScope = _exp
                            } else {
                                return nil
                            }
                        }
                    } else {
                        return nil
                    }
                }
            }
            return nil
        }
        var typeInfo *Ast_TypeInfo
        
        {
            _typeInfo := Builtin_getTypeInfo(_env, scope)
            if _typeInfo == nil{
                return nil
            } else {
                typeInfo = _typeInfo.(*Ast_TypeInfo)
            }
        }
        var validGenType bool
        validGenType = false
        for _, _genType := range( genTypeList.Items ) {
            genType := _genType
            if genType.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate{
                validGenType = true
                break
            }
        }
        if validGenType{
            if _switch0 := typeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Map {
                if genTypeList.Len() != 2{
                    Util_err(_env, _env.GetVM().String_format("illegal map param -- %d", Lns_2DDD(genTypeList.Len())))
                }
                var keyType *Ast_TypeInfo
                keyType = genTypeList.GetAt(1)
                var valType *Ast_TypeInfo
                valType = genTypeList.GetAt(2)
                return self.processInfo.FP.CreateMap_(_env, typeInfo.FP.Get_canDealGenInherit(_env), Ast_AccessMode__Pub, typeInfo.FP.Get_parentInfo(_env), keyType, valType, typeInfo.FP.Get_mutMode(_env))
            } else if _switch0 == Ast_TypeInfoKind__Ext {
                self.hasLuaval = true
                if genTypeList.Len() != 1{
                    Util_err(_env, _env.GetVM().String_format("illegal param -- %d", Lns_2DDD(genTypeList.Len())))
                }
                switch _matchExp0 := self.processInfo.FP.CreateLuaval(_env, genTypeList.GetAt(1), true).(type) {
                case *Ast_LuavalResult__OK:
                    workType := _matchExp0.Val1
                    if self.ctrl_info.ValidLuaval{
                        return workType
                    }
                    return genTypeList.GetAt(1)
                case *Ast_LuavalResult__Err:
                    mess := _matchExp0.Val1
                    Util_err(_env, mess)
                }
            } else if _switch0 == Ast_TypeInfoKind__DDD {
                if genTypeList.Len() != 1{
                    Util_err(_env, _env.GetVM().String_format("illegal map param -- %d", Lns_2DDD(genTypeList.Len())))
                }
                return &self.processInfo.FP.CreateDDD(_env, genTypeList.GetAt(1), true, false).Ast_TypeInfo
            } else {
                Util_err(_env, _env.GetVM().String_format("not support type -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
            }
        }
        return typeInfo
    }
    var mutable bool
    mutable = true
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(typeName,"^&", nil, nil))){
        mutable = false
        typeName = Builtin_convExp0_15472(Lns_2DDD(_env.GetVM().String_gsub(typeName,"^&", "")))
    }
    var genTypeList *LnsList2_[*Ast_TypeInfo]
    genTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    var endIndex LnsAny
    _,endIndex = Builtin_convExp0_15495(Lns_2DDD(_env.GetVM().String_find(typeName,"[%w%.]+<", nil, nil)))
    var suffix string
    suffix = ""
    if endIndex != nil{
        endIndex_727 := endIndex.(LnsInt)
        var genTypeName string
        genTypeName = _env.GetVM().String_sub(typeName,endIndex_727 + 1, nil)
        for  {
            {
                _tailIndex := Builtin_convExp0_15591(Lns_2DDD(_env.GetVM().String_find(genTypeName,"[,>]", nil, nil)))
                if !Lns_IsNil( _tailIndex ) {
                    tailIndex := _tailIndex.(LnsInt)
                    var genType *Ast_TypeInfo
                    genType = self.FP.getTypeInfo(_env, _env.GetVM().String_sub(genTypeName,1, tailIndex - 1), targetScope)
                    genTypeList.Insert(genType)
                    genTypeName = _env.GetVM().String_sub(genTypeName,tailIndex + 1, nil)
                } else {
                    suffix = _env.GetVM().String_sub(genTypeName,1, nil)
                    break
                }
            }
        }
        typeName = _env.GetVM().String_sub(typeName,1, endIndex_727 - 1) + suffix
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var nilable bool
    var orgTypeName string
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(typeName,"!$", nil, nil))){
        nilable = true
        orgTypeName = Builtin_convExp0_15657(Lns_2DDD(_env.GetVM().String_gsub(typeName,"!$", "")))
    } else { 
        nilable = false
        orgTypeName = typeName
    }
    {
        __exp := Builtin_getTypeInfoFromScope(_env, self.transUnit.FP.Get_scope(_env), orgTypeName, genTypeList)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
        } else {
            if targetScope != nil{
                targetScope_743 := targetScope.(*Ast_Scope)
                {
                    __exp := Builtin_getTypeInfoFromScope(_env, targetScope_743, orgTypeName, genTypeList)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*Ast_TypeInfo)
                        typeInfo = _exp
                    }
                }
            }
        }
    }
    if typeInfo == Ast_headTypeInfo{
        Util_err(_env, _env.GetVM().String_format("not found builtin -- %s", Lns_2DDD(orgTypeName)))
    }
    if nilable{
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
    }
    if mutable{
        return typeInfo
    }
    typeInfo = self.modifier.FP.CreateModifier(_env, typeInfo, Ast_MutMode__IMut)
    return typeInfo
}
// 661: decl @lune.@base.@Builtin.Builtin.createGenType
func (self *Builtin_Builtin) createGenType(_env *LnsEnv, typeName string,genTypeList *LnsList2_[*Ast_AlternateTypeInfo],workParentInfo LnsAny) string {
    var parentInfo *Ast_TypeInfo
    parentInfo = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( workParentInfo) ||
        _env.SetStackVal( Ast_headTypeInfo) ).(*Ast_TypeInfo)
    var name string
    name = typeName
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(typeName,"<", nil, nil))){
        name = ""
            {
                _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(typeName,"[^<>,%s]+")
                for {
                    _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                    _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                    if Lns_IsNil( _applyPrev0 ) { break }
                    token := _applyPrev0.(string)
                    if len(name) == 0{
                        name = token
                    } else { 
                        genTypeList.Insert(Lns_car(self.processInfo.FP.CreateAlternate(_env, true, genTypeList.Len() + 1, token, Ast_AccessMode__Pri, parentInfo, nil, nil, nil)).(*Ast_AlternateTypeInfo))
                    }
                }
            }
    }
    return name
}
// 686: decl @lune.@base.@Builtin.Builtin.processField
func (self *Builtin_Builtin) processField(_env *LnsEnv, prefixName string,orgName string,fieldName string,info *LnsMap2_[string,*LnsList2_[string]],parentInfo *Ast_TypeInfo) {
    self.hasLuaval = false
    if self.targetLuaVer.FP.IsSupport(_env, _env.GetVM().String_format("%s.%s", Lns_2DDD(prefixName, fieldName))){
        if _switch1 := _env.NilAccFin( _env.NilAccPush( info.Get("type")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsList2_[string]).GetAt(1))); _switch1 == "var" {
            var symbol *Ast_SymbolInfo
            symbol = Lns_unwrap( Lns_car(self.transUnit.FP.Get_scope(_env).FP.Add(_env, self.processInfo, Ast_SymbolKind__Var, false, true, fieldName, nil, self.FP.getTypeInfo(_env, Lns_unwrap( _env.NilAccFin( _env.NilAccPush( info.Get("typeInfo")) && 
            _env.NilAccPush( _env.NilAccPop().(*LnsList2_[string]).GetAt(1)))).(string), nil), Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_4_(_env, prefixName, fieldName, symbol)
        } else {
            var funcType LnsAny
            funcType = _env.NilAccFin( _env.NilAccPush( info.Get("type")) && 
            _env.NilAccPush( _env.NilAccPop().(*LnsList2_[string]).GetAt(1)))
            var mutable bool
            var staticFlag bool
            var kind LnsInt
            var symbolKind LnsInt
            var abstractFlag bool
            var accessMode LnsInt
            if _switch0 := funcType; _switch0 == "method" || _switch0 == "mut" {
                staticFlag = false
                kind = Ast_TypeInfoKind__Method
                symbolKind = Ast_SymbolKind__Mtd
                abstractFlag = false
                accessMode = Ast_AccessMode__Pub
                mutable = funcType == "mut"
            } else if _switch0 == "STATIC" {
                staticFlag = true
                kind = Ast_TypeInfoKind__Func
                symbolKind = Ast_SymbolKind__Fun
                abstractFlag = false
                accessMode = Ast_AccessMode__Pri
                mutable = false
            } else if _switch0 == "abstract" {
                staticFlag = false
                kind = Ast_TypeInfoKind__Method
                symbolKind = Ast_SymbolKind__Mtd
                abstractFlag = true
                accessMode = Ast_AccessMode__Pro
                mutable = true
            } else if _switch0 == "macro" {
                staticFlag = true
                kind = Ast_TypeInfoKind__Macro
                symbolKind = Ast_SymbolKind__Mac
                abstractFlag = false
                accessMode = Ast_AccessMode__Pub
                mutable = false
            } else if _switch0 == "formfunc" {
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
            self.transUnit.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
            var scope *Ast_Scope
            scope = self.transUnit.FP.Get_scope(_env)
            var asyncMode LnsInt
            asyncMode = Ast_Async__Async
            var fieldGenTypeList *LnsList2_[*Ast_AlternateTypeInfo]
            fieldGenTypeList = NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})
            var argTypeList *LnsList2_[*Ast_TypeInfo]
            argTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
            var retTypeList *LnsList2_[*Ast_TypeInfo]
            retTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
            var typeInfo *Ast_NormalTypeInfo
            typeInfo = self.processInfo.FP.CreateFuncAsync(_env, abstractFlag, true, scope, kind, parentInfo, Ast_getBuiltinMut(_env, parentInfo).FP, false, true, staticFlag, accessMode, fieldName, asyncMode, NewLnsList2_[*Ast_TypeInfo](Ast_TypeInfo_toSlice( Lns_2DDD(fieldGenTypeList.Unpack()))), argTypeList, retTypeList, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mutable) &&
                _env.SetStackVal( Ast_MutMode__Mut) ||
                _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
            self.FP.createGenType(_env, orgName, fieldGenTypeList, &typeInfo.Ast_TypeInfo)
            for _, _altType := range( fieldGenTypeList.Items ) {
                altType := _altType
                scope.FP.AddAlternate(_env, self.processInfo, accessMode, altType.FP.Get_rawTxt(_env), Types_nonePos, &altType.Ast_TypeInfo)
            }
            for _, _argType := range( Lns_unwrap( info.Get("arg")).(*LnsList2_[string]).Items ) {
                argType := _argType
                argTypeList.Insert(self.FP.getTypeInfo(_env, argType, scope))
            }
            for _, _retType := range( Lns_unwrap( info.Get("ret")).(*LnsList2_[string]).Items ) {
                retType := _retType
                var retTypeInfo *Ast_TypeInfo
                retTypeInfo = self.FP.getTypeInfo(_env, retType, scope)
                retTypeList.Insert(retTypeInfo)
            }
            if self.hasLuaval{
                Builtin_builtinFunc.FP.addLuavalFunc(_env, &typeInfo.Ast_TypeInfo)
            }
            self.transUnit.FP.PopScope(_env)
            Builtin_builtinFunc.FP.Get_allFuncTypeSet(_env).Add(&typeInfo.Ast_TypeInfo)
            Ast_addBuiltinMut(_env, &typeInfo.Ast_TypeInfo, scope)
            var symInfo *Ast_SymbolInfo
            symInfo = Lns_unwrap( Lns_car(self.transUnit.FP.Get_scope(_env).FP.Add(_env, self.processInfo, symbolKind, false, kind == Ast_TypeInfoKind__Func, fieldName, nil, &typeInfo.Ast_TypeInfo, accessMode, staticFlag, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mutable) &&
                _env.SetStackVal( Ast_MutMode__Mut) ||
                _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_4_(_env, prefixName, fieldName, symInfo)
        }
    }
}
// 818: decl @lune.@base.@Builtin.Builtin.registClass
func (self *Builtin_Builtin) registClass(_env *LnsEnv, nameList *LnsList2_[string],name2FieldInfo *LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]],pos Types_Position,genTypeList *LnsList2_[*Ast_AlternateTypeInfo]) *Ast_TypeInfo {
    var classKind LnsInt
    classKind = TransUnitIF_DeclClassMode__Class
    {
        _types := _env.NilAccFin( _env.NilAccPush( name2FieldInfo.Get("__attrib")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsMap2_[string,*LnsList2_[string]]).Get("type")))
        if !Lns_IsNil( _types ) {
            types := _types.(*LnsList2_[string])
            if types.Len() > 0{
                if _switch0 := types.GetAt(1); _switch0 == "interface" {
                    classKind = TransUnitIF_DeclClassMode__Interface
                } else if _switch0 == "module" {
                    classKind = TransUnitIF_DeclClassMode__Module
                }
            }
        }
    }
    var interfaceList *LnsList2_[*Ast_TypeInfo]
    interfaceList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    {
        __exp := _env.NilAccFin( _env.NilAccPush( name2FieldInfo.Get("__attrib")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsMap2_[string,*LnsList2_[string]]).Get("implements")))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*LnsList2_[string])
            for _, _ifname := range( _exp.Items ) {
                ifname := _ifname
                var ifType *Ast_TypeInfo
                ifType = self.FP.getTypeInfo(_env, ifname, nil)
                interfaceList.Insert(ifType)
            }
        }
    }
    var Builtin_registerClass func(_env *LnsEnv, regName string,index LnsInt,last bool) *Ast_TypeInfo
    Builtin_registerClass = func(_env *LnsEnv, regName string,index LnsInt,last bool) *Ast_TypeInfo {
        var workClassKind LnsInt
        if last{
            workClassKind = classKind
        } else { 
            if index == 1{
                workClassKind = TransUnitIF_DeclClassMode__Module
            } else { 
                workClassKind = TransUnitIF_DeclClassMode__Class
            }
        }
        var classType *Ast_TypeInfo
        classType = Ast_headTypeInfo
        if _switch0 := workClassKind; _switch0 == TransUnitIF_DeclClassMode__Class || _switch0 == TransUnitIF_DeclClassMode__Interface {
            var declMode LnsInt
            if workClassKind == TransUnitIF_DeclClassMode__Class{
                declMode = TransUnitIF_DeclClassMode__Class
            } else { 
                declMode = TransUnitIF_DeclClassMode__Interface
            }
            classType = self.transUnit.FP.PushClassLow(_env, self.processInfo, pos, declMode, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( declMode != TransUnitIF_DeclClassMode__Interface) ||
                _env.SetStackVal( regName == "Nilable") ).(bool), false, nil, interfaceList, genTypeList, true, regName, true, Ast_AccessMode__Pub, nil)
            Builtin_builtinFunc.FP.registerClass(_env, classType)
        } else if _switch0 == TransUnitIF_DeclClassMode__Module {
            classType = self.transUnit.FP.PushModuleLow(_env, self.processInfo, true, regName, true)
            self.transUnit.FP.Get_scope(_env).FP.Get_outerScope(_env).FP.Add(_env, self.processInfo, Ast_SymbolKind__Typ, false, false, regName, nil, classType, Ast_AccessMode__Local, true, Ast_MutMode__Mut, true, false)
        }
        return classType
    }
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    for _index, _name := range( nameList.Items ) {
        index := _index + 1
        name := _name
        parentInfo = Builtin_registerClass(_env, name, index, index == nameList.Len())
    }
    return parentInfo
}
// 888: decl @lune.@base.@Builtin.Builtin.createBuiltinAlge
func (self *Builtin_Builtin) createBuiltinAlge(_env *LnsEnv, algesymList *LnsList2_[*LnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]]) {
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    for _, _algeMap := range( algesymList.Items ) {
        algeMap := _algeMap
        for _genAlgeName, _valMapList := range( algeMap.Items ) {
            genAlgeName := _genAlgeName
            valMapList := _valMapList
            var genTypeList *LnsList2_[*Ast_AlternateTypeInfo]
            genTypeList = NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})
            var algeName string
            algeName = self.FP.createGenType(_env, genAlgeName, genTypeList, nil)
            var typename2typeInfo *LnsMap2_[string,*Ast_TypeInfo]
            typename2typeInfo = NewLnsMap2_[string,*Ast_TypeInfo]( map[string]*Ast_TypeInfo{})
            for _, _genType := range( genTypeList.Items ) {
                genType := _genType
                typename2typeInfo.Set(genType.FP.Get_rawTxt(_env),&genType.Ast_TypeInfo)
            }
            var algeScope *Ast_Scope
            algeScope = NewAst_Scope(_env, self.processInfo, self.transUnit.FP.Get_scope(_env), Ast_ScopeKind__Class, nil, nil)
            var algeTypeInfo *Ast_AlgeTypeInfo
            algeTypeInfo = self.processInfo.FP.CreateAlge(_env, algeScope, parentInfo, Ast_getBuiltinMut(_env, parentInfo).FP, false, Ast_AccessMode__Pub, algeName, NewLnsList2_[*Ast_TypeInfo](Ast_TypeInfo_toSlice( Lns_2DDD(genTypeList.Unpack()))))
            Ast_addBuiltinMut(_env, &algeTypeInfo.Ast_TypeInfo, algeScope)
            Builtin_builtinFunc.FP.registerAlge(_env, &algeTypeInfo.Ast_TypeInfo)
            self.transUnit.FP.Get_scope(_env).FP.AddAlge(_env, self.processInfo, Ast_AccessMode__Pub, algeName, nil, &algeTypeInfo.Ast_TypeInfo)
            for _, _valMap := range( valMapList.Items ) {
                valMap := _valMap
                for _valName, _valTypeNameList := range( valMap.Items ) {
                    valName := _valName
                    valTypeNameList := _valTypeNameList
                    var algeValSym *Ast_SymbolInfo
                    algeValSym = Lns_unwrap( Lns_car(algeScope.FP.AddAlgeVal(_env, self.processInfo, valName, Types_nonePos, &algeTypeInfo.Ast_TypeInfo))).(*Ast_SymbolInfo)
                    var valTypeList *LnsList2_[*Ast_TypeInfo]
                    valTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
                    for _, _typeName := range( valTypeNameList.Items ) {
                        typeName := _typeName
                        valTypeList.Insert(_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( typename2typeInfo.Get(typeName)) ||
                            _env.SetStackVal( self.FP.getTypeInfo(_env, typeName, nil)) ).(*Ast_TypeInfo))
                    }
                    var algeValInfo *Ast_AlgeValInfo
                    algeValInfo = NewAst_AlgeValInfo(_env, algeValSym.FP.Get_name(_env), valTypeList, algeTypeInfo, algeValSym)
                    algeTypeInfo.FP.AddValInfo(_env, algeValInfo)
                }
            }
        }
    }
}
// 934: decl @lune.@base.@Builtin.Builtin.registBuiltInScope
func (self *Builtin_Builtin) RegistBuiltInScope(_env *LnsEnv) *Builtin_BuiltinFuncType {
    if Builtin_readyBuiltin{
        return Builtin_builtinFunc
    }
    Builtin_readyBuiltin = true
    var builtInInfo *LnsList2_[*LnsMap2_[string,*LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]]]]
    builtInInfo = Builtin_getBuiltInInfo_5_(_env)
    var builtinModuleName2Scope *LnsMap2_[string,*Ast_Scope]
    builtinModuleName2Scope = NewLnsMap2_[string,*Ast_Scope]( map[string]*Ast_Scope{})
    var mapType *Ast_TypeInfo
    mapType = self.processInfo.FP.CreateMap_(_env, true, Ast_AccessMode__Pub, Ast_headTypeInfo, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "_ENV", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "_G", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "__line__", nil, Ast_builtinTypeInt, Ast_MutMode__IMut, true)
    var pos Types_Position
    pos = NewTypes_Position(_env, 0, 0, "@builtin@", nil)
    for _, _builtinClassInfo := range( builtInInfo.Items ) {
        builtinClassInfo := _builtinClassInfo
        for _className, _name2FieldInfo := range( builtinClassInfo.Items ) {
            className := _className
            name2FieldInfo := _name2FieldInfo
            var name string
            name = className
            var genTypeList *LnsList2_[*Ast_AlternateTypeInfo]
            genTypeList = NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})
            if _switch0 := className; _switch0 == "_List<T>" {
                name = "_List"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeList_.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "__List<T>" {
                name = "__List"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeList__.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Array<T>" {
                name = "Array"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeArray.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "_Set<T>" {
                name = "_Set"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeSet_.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "__Set<T>" {
                name = "__Set"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeSet__.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Map<K,V>" {
                name = "Map"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeMap_.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "__Map<K,V>" {
                name = "__Map"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeMap__.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Nilable<_T>" {
                name = "Nilable"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeBox.FP.Get_itemTypeInfoList(_env))
            } else {
                name = self.FP.createGenType(_env, className, genTypeList, nil)
            }
            var nameList *LnsList2_[string]
            nameList = Util_splitModule(_env, name)
            var parentInfo *Ast_TypeInfo
            parentInfo = Ast_headTypeInfo
            if name != ""{
                parentInfo = self.FP.registClass(_env, nameList, name2FieldInfo, pos, genTypeList)
            }
            if Lns_op_not(builtinModuleName2Scope.Get(name)){
                if name != ""{
                    builtinModuleName2Scope.Set(name,self.transUnit.FP.Get_scope(_env))
                }
                {
                    __forsortCollection0 := name2FieldInfo
                    __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
                    __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
                    for _, _orgFieldName := range( __forsortSorted0.Items ) {
                        info := __forsortCollection0.Items[ _orgFieldName ]
                        orgFieldName := _orgFieldName
                        var fieldName string
                        fieldName = Builtin_convExp0_17702(Lns_2DDD(_env.GetVM().String_gsub(orgFieldName,"<.*", "")))
                        if _switch1 := fieldName; _switch1 == "__attrib" {
                        } else {
                            self.FP.processField(_env, Lns_car(_env.GetVM().String_gsub(name,"%.", "_")).(string), orgFieldName, fieldName, info, parentInfo)
                        }
                    }
                }
            }
            if name != ""{
                for range( nameList.Items ) {
                    self.transUnit.FP.PopClass(_env)
                }
            }
        }
    }
    self.FP.createBuiltinAlge(_env, Builtin_getBuiltinAlgeInfo_6_(_env))
    var threadSafeSet *LnsSet2_[*Ast_TypeInfo]
    threadSafeSet = NewLnsSet2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Builtin_builtinFunc.Lns_error, Builtin_builtinFunc.Lns_print, Builtin_builtinFunc.Lns_type, Builtin_builtinFunc.Lns_tonumber, Builtin_builtinFunc.Io_open, Builtin_builtinFunc.G_set_has, Builtin_builtinFunc.G_set_add, Builtin_builtinFunc.G__set_has, Builtin_builtinFunc.G__set_add))
    for _typeInfo := range( Builtin_builtinFunc.FP.Get_allFuncTypeSet(_env).Items ) {
        typeInfo := _typeInfo
        if Lns_op_not(threadSafeSet.Has(typeInfo)){
            Builtin_builtinFunc.FP.Get_needThreadingTypes(_env).Add(typeInfo)
        }
    }
    return Builtin_builtinFunc
}
// 1093: decl @lune.@base.@Builtin.Builtin.getBuiltInFuncType
func (self *Builtin_Builtin) GetBuiltInFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType {
    if Builtin_readyBuiltin{
        return Builtin_builtinFunc
    }
    Util_err(_env, "builtinFunc is not ready")
// insert a dummy
    return nil
}
// 178: decl @lune.@base.@Builtin.BuiltinFuncType.addLuavalFunc
func (self *Builtin_BuiltinFuncType) addLuavalFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    self.luavalFuncTypeSet.Add(typeInfo)
}
// 182: decl @lune.@base.@Builtin.BuiltinFuncType.isLuavalFunc
func (self *Builtin_BuiltinFuncType) IsLuavalFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return self.luavalFuncTypeSet.Has(typeInfo)
}
// 186: decl @lune.@base.@Builtin.BuiltinFuncType.register
func (self *Builtin_BuiltinFuncType) register(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) {
    self.allSymbol.Insert(symbolInfo)
    self.allSymbolSet.Add(symbolInfo)
}
// 190: decl @lune.@base.@Builtin.BuiltinFuncType.registerClass
func (self *Builtin_BuiltinFuncType) registerClass(_env *LnsEnv, classInfo *Ast_TypeInfo) {
    self.allClass.Insert(classInfo)
    if _switch0 := classInfo.FP.Get_rawTxt(_env); _switch0 == "__Er" {
        self.G__er_ = classInfo
    } else if _switch0 == "" {
        self.Lns_ = classInfo
    } else if _switch0 == "iStream" {
        self.Istream_ = classInfo
    } else if _switch0 == "oStream" {
        self.Ostream_ = classInfo
    } else if _switch0 == "__pipe" {
        self.G__pipe_ = classInfo
    } else if _switch0 == "__lns_runMode" {
        self.G__lns_runmode_ = classInfo
    } else if _switch0 == "__lns_runtime" {
        self.G__lns_runtime_ = classInfo
    } else if _switch0 == "__lns_capability" {
        self.G__lns_capability_ = classInfo
    } else if _switch0 == "__lns_Sync_Flag" {
        self.G__lns_sync_flag_ = classInfo
    } else if _switch0 == "__lns_Sync" {
        self.G__lns_sync_ = classInfo
    } else if _switch0 == "luaStream" {
        self.Luastream_ = classInfo
    } else if _switch0 == "Mapping" {
        self.Mapping_ = classInfo
    } else if _switch0 == "__Runner" {
        self.G__runner_ = classInfo
    } else if _switch0 == "__Processor" {
        self.G__processor_ = classInfo
    } else if _switch0 == "io" {
        self.Io_ = classInfo
    } else if _switch0 == "package" {
        self.Package_ = classInfo
    } else if _switch0 == "os" {
        self.Os_ = classInfo
    } else if _switch0 == "string" {
        self.String_ = classInfo
    } else if _switch0 == "str" {
        self.Str_ = classInfo
    } else if _switch0 == "_List" {
        self.G_list_ = classInfo
    } else if _switch0 == "__List" {
        self.G__list_ = classInfo
    } else if _switch0 == "Array" {
        self.Array_ = classInfo
    } else if _switch0 == "_Set" {
        self.G_set_ = classInfo
    } else if _switch0 == "__Set" {
        self.G__set_ = classInfo
    } else if _switch0 == "math" {
        self.Math_ = classInfo
    } else if _switch0 == "debug" {
        self.Debug_ = classInfo
    } else if _switch0 == "Nilable" {
        self.Nilable_ = classInfo
    }
}
// 196: decl @lune.@base.@Builtin.BuiltinFuncType.registerAlge
func (self *Builtin_BuiltinFuncType) registerAlge(_env *LnsEnv, algeInfo *Ast_TypeInfo) {
    self.allClass.Insert(algeInfo)
    if _switch0 := algeInfo.FP.Get_rawTxt(_env); _switch0 == "__Ret" {
        self.G__ret_ = algeInfo
    }
}
// 493: decl @lune.@base.@Builtin.BuiltinFuncType.isStrFormFunc
func (self *Builtin_BuiltinFuncType) IsStrFormFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_srcTypeInfo(_env) == self.String_format{
        return true
    }
    return false
}
// declaration Class -- Builtin
type Builtin_BuiltinMtd interface {
    createBuiltinAlge(_env *LnsEnv, arg1 *LnsList2_[*LnsMap2_[string,*LnsList2_[*LnsMap2_[string,*LnsList2_[string]]]]])
    createGenType(_env *LnsEnv, arg1 string, arg2 *LnsList2_[*Ast_AlternateTypeInfo], arg3 LnsAny) string
    GetBuiltInFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType
    getTypeInfo(_env *LnsEnv, arg1 string, arg2 LnsAny) *Ast_TypeInfo
    processField(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 *LnsMap2_[string,*LnsList2_[string]], arg5 *Ast_TypeInfo)
    RegistBuiltInScope(_env *LnsEnv) *Builtin_BuiltinFuncType
    registClass(_env *LnsEnv, arg1 *LnsList2_[string], arg2 *LnsMap2_[string,*LnsMap2_[string,*LnsList2_[string]]], arg3 Types_Position, arg4 *LnsList2_[*Ast_AlternateTypeInfo]) *Ast_TypeInfo
}
type Builtin_Builtin struct {
    transUnit *BuiltinTransUnit_TransUnit
    targetLuaVer *LuaVer_LuaVerInfo
    ctrl_info *Types_TransCtrlInfo
    processInfo *Ast_ProcessInfo
    modifier *TransUnitIF_Modifier
    hasLuaval bool
    FP Builtin_BuiltinMtd
}
func Builtin_Builtin2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Builtin_Builtin).FP
}
func Builtin_Builtin_toSlice(slice []LnsAny) []*Builtin_Builtin {
    ret := make([]*Builtin_Builtin, len(slice))
    for index, val := range slice {
        ret[index] = val.(Builtin_BuiltinDownCast).ToBuiltin_Builtin()
    }
    return ret
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
func NewBuiltin_Builtin(_env *LnsEnv, arg1 *LuaVer_LuaVerInfo, arg2 *Types_TransCtrlInfo) *Builtin_Builtin {
    obj := &Builtin_Builtin{}
    obj.FP = obj
    obj.InitBuiltin_Builtin(_env, arg1, arg2)
    return obj
}
// 51: DeclConstr
func (self *Builtin_Builtin) InitBuiltin_Builtin(_env *LnsEnv, targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo) {
    self.hasLuaval = false
    self.processInfo = Ast_getRootProcessInfo(_env)
    self.transUnit = NewBuiltinTransUnit_TransUnit(_env, ctrl_info, self.processInfo)
    self.targetLuaVer = targetLuaVer
    self.ctrl_info = ctrl_info
    self.modifier = NewTransUnitIF_Modifier(_env, true, self.processInfo)
}


// declaration Class -- BuiltinFuncType
type Builtin_BuiltinFuncTypeMtd interface {
    addLuavalFunc(_env *LnsEnv, arg1 *Ast_TypeInfo)
    Get_allClass(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]
    Get_allFuncTypeSet(_env *LnsEnv) *LnsSet2_[*Ast_TypeInfo]
    Get_allSymbol(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]
    Get_allSymbolSet(_env *LnsEnv) *LnsSet2_[*Ast_SymbolInfo]
    Get_needThreadingTypes(_env *LnsEnv) *LnsSet2_[*Ast_TypeInfo]
    IsLuavalFunc(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsStrFormFunc(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    register(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    registerAlge(_env *LnsEnv, arg1 *Ast_TypeInfo)
    registerClass(_env *LnsEnv, arg1 *Ast_TypeInfo)
}
type Builtin_BuiltinFuncType struct {
    G__er_ *Ast_TypeInfo
    G__er_get_txt *Ast_TypeInfo
    G__er_get_txt_sym *Ast_SymbolInfo
    Lns_ *Ast_TypeInfo
    Lns___join *Ast_TypeInfo
    Lns___join_sym *Ast_SymbolInfo
    Lns___run *Ast_TypeInfo
    Lns___run_sym *Ast_SymbolInfo
    Lns___serr *Ast_TypeInfo
    Lns___serr_sym *Ast_SymbolInfo
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
    Istream_close *Ast_TypeInfo
    Istream_close_sym *Ast_SymbolInfo
    Istream_read *Ast_TypeInfo
    Istream_read_sym *Ast_SymbolInfo
    Ostream_ *Ast_TypeInfo
    Ostream_close *Ast_TypeInfo
    Ostream_close_sym *Ast_SymbolInfo
    Ostream_flush *Ast_TypeInfo
    Ostream_flush_sym *Ast_SymbolInfo
    Ostream_write *Ast_TypeInfo
    Ostream_write_sym *Ast_SymbolInfo
    G__pipe_ *Ast_TypeInfo
    G__pipe_get *Ast_TypeInfo
    G__pipe_get_sym *Ast_SymbolInfo
    G__pipe_put *Ast_TypeInfo
    G__pipe_put_sym *Ast_SymbolInfo
    G__lns_runmode_ *Ast_TypeInfo
    G__lns_runmode_Queue *Ast_TypeInfo
    G__lns_runmode_Queue_sym *Ast_SymbolInfo
    G__lns_runmode_Skip *Ast_TypeInfo
    G__lns_runmode_Skip_sym *Ast_SymbolInfo
    G__lns_runmode_Sync *Ast_TypeInfo
    G__lns_runmode_Sync_sym *Ast_SymbolInfo
    G__lns_runtime_ *Ast_TypeInfo
    G__lns_runtime_dumpLog *Ast_TypeInfo
    G__lns_runtime_dumpLog_sym *Ast_SymbolInfo
    G__lns_runtime_enableDebugLog *Ast_TypeInfo
    G__lns_runtime_enableDebugLog_sym *Ast_SymbolInfo
    G__lns_runtime_enableLog *Ast_TypeInfo
    G__lns_runtime_enableLog_sym *Ast_SymbolInfo
    G__lns_runtime_log *Ast_TypeInfo
    G__lns_runtime_log_sym *Ast_SymbolInfo
    G__lns_capability_ *Ast_TypeInfo
    G__lns_capability_async *Ast_TypeInfo
    G__lns_capability_async_sym *Ast_SymbolInfo
    G__lns_sync_flag_ *Ast_TypeInfo
    G__lns_sync_flag_set *Ast_TypeInfo
    G__lns_sync_flag_set_sym *Ast_SymbolInfo
    G__lns_sync_flag_wait *Ast_TypeInfo
    G__lns_sync_flag_wait_sym *Ast_SymbolInfo
    G__lns_sync_ *Ast_TypeInfo
    G__lns_sync__createPipe *Ast_TypeInfo
    G__lns_sync__createPipe_sym *Ast_SymbolInfo
    G__lns_sync_createFlag *Ast_TypeInfo
    G__lns_sync_createFlag_sym *Ast_SymbolInfo
    G__lns_sync_createProcesser *Ast_TypeInfo
    G__lns_sync_createProcesser_sym *Ast_SymbolInfo
    Luastream_ *Ast_TypeInfo
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
    Mapping__toMap *Ast_TypeInfo
    Mapping__toMap_sym *Ast_SymbolInfo
    G__runner_ *Ast_TypeInfo
    G__runner_run *Ast_TypeInfo
    G__runner_run_sym *Ast_SymbolInfo
    G__processor_ *Ast_TypeInfo
    G__processor_end *Ast_TypeInfo
    G__processor_end_sym *Ast_SymbolInfo
    Io_ *Ast_TypeInfo
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
    Package_path *Ast_TypeInfo
    Package_path_sym *Ast_SymbolInfo
    Package_searchpath *Ast_TypeInfo
    Package_searchpath_sym *Ast_SymbolInfo
    Os_ *Ast_TypeInfo
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
    Str_replace *Ast_TypeInfo
    Str_replace_sym *Ast_SymbolInfo
    Str_reverse *Ast_TypeInfo
    Str_reverse_sym *Ast_SymbolInfo
    Str_sub *Ast_TypeInfo
    Str_sub_sym *Ast_SymbolInfo
    Str_upper *Ast_TypeInfo
    Str_upper_sym *Ast_SymbolInfo
    G_list_ *Ast_TypeInfo
    G_list___less *Ast_TypeInfo
    G_list___less_sym *Ast_SymbolInfo
    G_list_insert *Ast_TypeInfo
    G_list_insert_sym *Ast_SymbolInfo
    G_list_remove *Ast_TypeInfo
    G_list_remove_sym *Ast_SymbolInfo
    G_list_sort *Ast_TypeInfo
    G_list_sort_sym *Ast_SymbolInfo
    G_list_unpack *Ast_TypeInfo
    G_list_unpack_sym *Ast_SymbolInfo
    G__list_ *Ast_TypeInfo
    G__list___less *Ast_TypeInfo
    G__list___less_sym *Ast_SymbolInfo
    G__list_insert *Ast_TypeInfo
    G__list_insert_sym *Ast_SymbolInfo
    G__list_remove *Ast_TypeInfo
    G__list_remove_sym *Ast_SymbolInfo
    G__list_sort *Ast_TypeInfo
    G__list_sort_sym *Ast_SymbolInfo
    G__list_unpack *Ast_TypeInfo
    G__list_unpack_sym *Ast_SymbolInfo
    Array_ *Ast_TypeInfo
    Array___less *Ast_TypeInfo
    Array___less_sym *Ast_SymbolInfo
    Array_sort *Ast_TypeInfo
    Array_sort_sym *Ast_SymbolInfo
    Array_unpack *Ast_TypeInfo
    Array_unpack_sym *Ast_SymbolInfo
    G_set_ *Ast_TypeInfo
    G_set_add *Ast_TypeInfo
    G_set_add_sym *Ast_SymbolInfo
    G_set_and *Ast_TypeInfo
    G_set_and_sym *Ast_SymbolInfo
    G_set_clone *Ast_TypeInfo
    G_set_clone_sym *Ast_SymbolInfo
    G_set_del *Ast_TypeInfo
    G_set_del_sym *Ast_SymbolInfo
    G_set_has *Ast_TypeInfo
    G_set_has_sym *Ast_SymbolInfo
    G_set_len *Ast_TypeInfo
    G_set_len_sym *Ast_SymbolInfo
    G_set_or *Ast_TypeInfo
    G_set_or_sym *Ast_SymbolInfo
    G_set_sub *Ast_TypeInfo
    G_set_sub_sym *Ast_SymbolInfo
    G__set_ *Ast_TypeInfo
    G__set_add *Ast_TypeInfo
    G__set_add_sym *Ast_SymbolInfo
    G__set_and *Ast_TypeInfo
    G__set_and_sym *Ast_SymbolInfo
    G__set_clone *Ast_TypeInfo
    G__set_clone_sym *Ast_SymbolInfo
    G__set_del *Ast_TypeInfo
    G__set_del_sym *Ast_SymbolInfo
    G__set_has *Ast_TypeInfo
    G__set_has_sym *Ast_SymbolInfo
    G__set_len *Ast_TypeInfo
    G__set_len_sym *Ast_SymbolInfo
    G__set_or *Ast_TypeInfo
    G__set_or_sym *Ast_SymbolInfo
    G__set_sub *Ast_TypeInfo
    G__set_sub_sym *Ast_SymbolInfo
    Math_ *Ast_TypeInfo
    Math_random *Ast_TypeInfo
    Math_random_sym *Ast_SymbolInfo
    Math_randomseed *Ast_TypeInfo
    Math_randomseed_sym *Ast_SymbolInfo
    Debug_ *Ast_TypeInfo
    Debug_getinfo *Ast_TypeInfo
    Debug_getinfo_sym *Ast_SymbolInfo
    Debug_getlocal *Ast_TypeInfo
    Debug_getlocal_sym *Ast_SymbolInfo
    Nilable_ *Ast_TypeInfo
    Nilable_val *Ast_TypeInfo
    Nilable_val_sym *Ast_SymbolInfo
    G__ret_ *Ast_TypeInfo
    allSymbol *LnsList2_[*Ast_SymbolInfo]
    allClass *LnsList2_[*Ast_TypeInfo]
    allFuncTypeSet *LnsSet2_[*Ast_TypeInfo]
    allSymbolSet *LnsSet2_[*Ast_SymbolInfo]
    needThreadingTypes *LnsSet2_[*Ast_TypeInfo]
    luavalFuncTypeSet *LnsSet2_[*Ast_TypeInfo]
    FP Builtin_BuiltinFuncTypeMtd
}
func Builtin_BuiltinFuncType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Builtin_BuiltinFuncType).FP
}
func Builtin_BuiltinFuncType_toSlice(slice []LnsAny) []*Builtin_BuiltinFuncType {
    ret := make([]*Builtin_BuiltinFuncType, len(slice))
    for index, val := range slice {
        ret[index] = val.(Builtin_BuiltinFuncTypeDownCast).ToBuiltin_BuiltinFuncType()
    }
    return ret
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
func NewBuiltin_BuiltinFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType {
    obj := &Builtin_BuiltinFuncType{}
    obj.FP = obj
    obj.InitBuiltin_BuiltinFuncType(_env)
    return obj
}
func (self *Builtin_BuiltinFuncType) Get_allSymbol(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]{ return self.allSymbol }
func (self *Builtin_BuiltinFuncType) Get_allClass(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]{ return self.allClass }
func (self *Builtin_BuiltinFuncType) Get_allFuncTypeSet(_env *LnsEnv) *LnsSet2_[*Ast_TypeInfo]{ return self.allFuncTypeSet }
func (self *Builtin_BuiltinFuncType) Get_allSymbolSet(_env *LnsEnv) *LnsSet2_[*Ast_SymbolInfo]{ return self.allSymbolSet }
func (self *Builtin_BuiltinFuncType) Get_needThreadingTypes(_env *LnsEnv) *LnsSet2_[*Ast_TypeInfo]{ return self.needThreadingTypes }
// 167: DeclConstr
func (self *Builtin_BuiltinFuncType) InitBuiltin_BuiltinFuncType(_env *LnsEnv) {
    self.G__er_ = Ast_headTypeInfo
    self.G__er_get_txt = Ast_headTypeInfo
    self.G__er_get_txt_sym = Ast_dummySymbol
    self.Lns_ = Ast_headTypeInfo
    self.Lns___join = Ast_headTypeInfo
    self.Lns___join_sym = Ast_dummySymbol
    self.Lns___run = Ast_headTypeInfo
    self.Lns___run_sym = Ast_dummySymbol
    self.Lns___serr = Ast_headTypeInfo
    self.Lns___serr_sym = Ast_dummySymbol
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
    self.Istream_close = Ast_headTypeInfo
    self.Istream_close_sym = Ast_dummySymbol
    self.Istream_read = Ast_headTypeInfo
    self.Istream_read_sym = Ast_dummySymbol
    self.Ostream_ = Ast_headTypeInfo
    self.Ostream_close = Ast_headTypeInfo
    self.Ostream_close_sym = Ast_dummySymbol
    self.Ostream_flush = Ast_headTypeInfo
    self.Ostream_flush_sym = Ast_dummySymbol
    self.Ostream_write = Ast_headTypeInfo
    self.Ostream_write_sym = Ast_dummySymbol
    self.G__pipe_ = Ast_headTypeInfo
    self.G__pipe_get = Ast_headTypeInfo
    self.G__pipe_get_sym = Ast_dummySymbol
    self.G__pipe_put = Ast_headTypeInfo
    self.G__pipe_put_sym = Ast_dummySymbol
    self.G__lns_runmode_ = Ast_headTypeInfo
    self.G__lns_runmode_Queue = Ast_headTypeInfo
    self.G__lns_runmode_Queue_sym = Ast_dummySymbol
    self.G__lns_runmode_Skip = Ast_headTypeInfo
    self.G__lns_runmode_Skip_sym = Ast_dummySymbol
    self.G__lns_runmode_Sync = Ast_headTypeInfo
    self.G__lns_runmode_Sync_sym = Ast_dummySymbol
    self.G__lns_runtime_ = Ast_headTypeInfo
    self.G__lns_runtime_dumpLog = Ast_headTypeInfo
    self.G__lns_runtime_dumpLog_sym = Ast_dummySymbol
    self.G__lns_runtime_enableDebugLog = Ast_headTypeInfo
    self.G__lns_runtime_enableDebugLog_sym = Ast_dummySymbol
    self.G__lns_runtime_enableLog = Ast_headTypeInfo
    self.G__lns_runtime_enableLog_sym = Ast_dummySymbol
    self.G__lns_runtime_log = Ast_headTypeInfo
    self.G__lns_runtime_log_sym = Ast_dummySymbol
    self.G__lns_capability_ = Ast_headTypeInfo
    self.G__lns_capability_async = Ast_headTypeInfo
    self.G__lns_capability_async_sym = Ast_dummySymbol
    self.G__lns_sync_flag_ = Ast_headTypeInfo
    self.G__lns_sync_flag_set = Ast_headTypeInfo
    self.G__lns_sync_flag_set_sym = Ast_dummySymbol
    self.G__lns_sync_flag_wait = Ast_headTypeInfo
    self.G__lns_sync_flag_wait_sym = Ast_dummySymbol
    self.G__lns_sync_ = Ast_headTypeInfo
    self.G__lns_sync__createPipe = Ast_headTypeInfo
    self.G__lns_sync__createPipe_sym = Ast_dummySymbol
    self.G__lns_sync_createFlag = Ast_headTypeInfo
    self.G__lns_sync_createFlag_sym = Ast_dummySymbol
    self.G__lns_sync_createProcesser = Ast_headTypeInfo
    self.G__lns_sync_createProcesser_sym = Ast_dummySymbol
    self.Luastream_ = Ast_headTypeInfo
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
    self.Mapping__toMap = Ast_headTypeInfo
    self.Mapping__toMap_sym = Ast_dummySymbol
    self.G__runner_ = Ast_headTypeInfo
    self.G__runner_run = Ast_headTypeInfo
    self.G__runner_run_sym = Ast_dummySymbol
    self.G__processor_ = Ast_headTypeInfo
    self.G__processor_end = Ast_headTypeInfo
    self.G__processor_end_sym = Ast_dummySymbol
    self.Io_ = Ast_headTypeInfo
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
    self.Package_path = Ast_headTypeInfo
    self.Package_path_sym = Ast_dummySymbol
    self.Package_searchpath = Ast_headTypeInfo
    self.Package_searchpath_sym = Ast_dummySymbol
    self.Os_ = Ast_headTypeInfo
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
    self.Str_replace = Ast_headTypeInfo
    self.Str_replace_sym = Ast_dummySymbol
    self.Str_reverse = Ast_headTypeInfo
    self.Str_reverse_sym = Ast_dummySymbol
    self.Str_sub = Ast_headTypeInfo
    self.Str_sub_sym = Ast_dummySymbol
    self.Str_upper = Ast_headTypeInfo
    self.Str_upper_sym = Ast_dummySymbol
    self.G_list_ = Ast_headTypeInfo
    self.G_list___less = Ast_headTypeInfo
    self.G_list___less_sym = Ast_dummySymbol
    self.G_list_insert = Ast_headTypeInfo
    self.G_list_insert_sym = Ast_dummySymbol
    self.G_list_remove = Ast_headTypeInfo
    self.G_list_remove_sym = Ast_dummySymbol
    self.G_list_sort = Ast_headTypeInfo
    self.G_list_sort_sym = Ast_dummySymbol
    self.G_list_unpack = Ast_headTypeInfo
    self.G_list_unpack_sym = Ast_dummySymbol
    self.G__list_ = Ast_headTypeInfo
    self.G__list___less = Ast_headTypeInfo
    self.G__list___less_sym = Ast_dummySymbol
    self.G__list_insert = Ast_headTypeInfo
    self.G__list_insert_sym = Ast_dummySymbol
    self.G__list_remove = Ast_headTypeInfo
    self.G__list_remove_sym = Ast_dummySymbol
    self.G__list_sort = Ast_headTypeInfo
    self.G__list_sort_sym = Ast_dummySymbol
    self.G__list_unpack = Ast_headTypeInfo
    self.G__list_unpack_sym = Ast_dummySymbol
    self.Array_ = Ast_headTypeInfo
    self.Array___less = Ast_headTypeInfo
    self.Array___less_sym = Ast_dummySymbol
    self.Array_sort = Ast_headTypeInfo
    self.Array_sort_sym = Ast_dummySymbol
    self.Array_unpack = Ast_headTypeInfo
    self.Array_unpack_sym = Ast_dummySymbol
    self.G_set_ = Ast_headTypeInfo
    self.G_set_add = Ast_headTypeInfo
    self.G_set_add_sym = Ast_dummySymbol
    self.G_set_and = Ast_headTypeInfo
    self.G_set_and_sym = Ast_dummySymbol
    self.G_set_clone = Ast_headTypeInfo
    self.G_set_clone_sym = Ast_dummySymbol
    self.G_set_del = Ast_headTypeInfo
    self.G_set_del_sym = Ast_dummySymbol
    self.G_set_has = Ast_headTypeInfo
    self.G_set_has_sym = Ast_dummySymbol
    self.G_set_len = Ast_headTypeInfo
    self.G_set_len_sym = Ast_dummySymbol
    self.G_set_or = Ast_headTypeInfo
    self.G_set_or_sym = Ast_dummySymbol
    self.G_set_sub = Ast_headTypeInfo
    self.G_set_sub_sym = Ast_dummySymbol
    self.G__set_ = Ast_headTypeInfo
    self.G__set_add = Ast_headTypeInfo
    self.G__set_add_sym = Ast_dummySymbol
    self.G__set_and = Ast_headTypeInfo
    self.G__set_and_sym = Ast_dummySymbol
    self.G__set_clone = Ast_headTypeInfo
    self.G__set_clone_sym = Ast_dummySymbol
    self.G__set_del = Ast_headTypeInfo
    self.G__set_del_sym = Ast_dummySymbol
    self.G__set_has = Ast_headTypeInfo
    self.G__set_has_sym = Ast_dummySymbol
    self.G__set_len = Ast_headTypeInfo
    self.G__set_len_sym = Ast_dummySymbol
    self.G__set_or = Ast_headTypeInfo
    self.G__set_or_sym = Ast_dummySymbol
    self.G__set_sub = Ast_headTypeInfo
    self.G__set_sub_sym = Ast_dummySymbol
    self.Math_ = Ast_headTypeInfo
    self.Math_random = Ast_headTypeInfo
    self.Math_random_sym = Ast_dummySymbol
    self.Math_randomseed = Ast_headTypeInfo
    self.Math_randomseed_sym = Ast_dummySymbol
    self.Debug_ = Ast_headTypeInfo
    self.Debug_getinfo = Ast_headTypeInfo
    self.Debug_getinfo_sym = Ast_dummySymbol
    self.Debug_getlocal = Ast_headTypeInfo
    self.Debug_getlocal_sym = Ast_dummySymbol
    self.Nilable_ = Ast_headTypeInfo
    self.Nilable_val = Ast_headTypeInfo
    self.Nilable_val_sym = Ast_dummySymbol
    self.G__ret_ = Ast_headTypeInfo
    self.allSymbol = NewLnsList2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
    self.allClass = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.allSymbolSet = NewLnsSet2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
    self.allFuncTypeSet = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.needThreadingTypes = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.luavalFuncTypeSet = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
}


func Lns_Builtin_init(_env *LnsEnv) {
    if init_Builtin { return }
    init_Builtin = true
    Builtin__mod__ = "@lune.@base.@Builtin"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Tokenizer_init(_env)
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Log_init(_env)
    Lns_TransUnitIF_init(_env)
    Lns_BuiltinTransUnit_init(_env)
    Builtin_builtinFunc = NewBuiltin_BuiltinFuncType(_env)
    
    Builtin_readyBuiltin = false
}
func init() {
    init_Builtin = false
}
