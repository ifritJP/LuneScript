// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Builtin bool
var Builtin__mod__ string
var Builtin_builtinFunc *Builtin_BuiltinFuncType
var Builtin_readyBuiltin bool
// for 586
func Builtin_convExp0_13970(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 578
func Builtin_convExp0_13851(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 73
func Builtin_convExp0_226(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 133
func Builtin_convExp0_626(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 581
func Builtin_convExp0_13876(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 602
func Builtin_convExp0_14025(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}



// 904: decl @lune.@base.@Builtin.Builtin.registBuiltInScope.processCopyAlterList
func Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env *LnsEnv, alterList *LnsList,typeList *LnsList) {
    for _, _typeInfo := range( typeList.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        alterList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( Ast_AlternateTypeInfoDownCastF(typeInfo.FP)).(*Ast_AlternateTypeInfo)))
    }
}

// 203: decl @lune.@base.@Builtin.getBuiltinFunc
func Builtin_getBuiltinFunc(_env *LnsEnv) *Builtin_BuiltinFuncType {
    return Builtin_builtinFunc
}

// 207: decl @lune.@base.@Builtin.setupBuiltinTypeInfo
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
    var Builtin_process_List func(_env *LnsEnv)
    Builtin_process_List = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "__less" {
            Builtin_builtinFunc.List___less = typeInfo
            Builtin_builtinFunc.List___less_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "insert" {
            Builtin_builtinFunc.List_insert = typeInfo
            Builtin_builtinFunc.List_insert_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "remove" {
            Builtin_builtinFunc.List_remove = typeInfo
            Builtin_builtinFunc.List_remove_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sort" {
            Builtin_builtinFunc.List_sort = typeInfo
            Builtin_builtinFunc.List_sort_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "unpack" {
            Builtin_builtinFunc.List_unpack = typeInfo
            Builtin_builtinFunc.List_unpack_sym = symInfo
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
    var Builtin_process_Set func(_env *LnsEnv)
    Builtin_process_Set = func(_env *LnsEnv) {
        if _switch0 := fieldName; _switch0 == "add" {
            Builtin_builtinFunc.Set_add = typeInfo
            Builtin_builtinFunc.Set_add_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "and" {
            Builtin_builtinFunc.Set_and = typeInfo
            Builtin_builtinFunc.Set_and_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "clone" {
            Builtin_builtinFunc.Set_clone = typeInfo
            Builtin_builtinFunc.Set_clone_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "del" {
            Builtin_builtinFunc.Set_del = typeInfo
            Builtin_builtinFunc.Set_del_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "has" {
            Builtin_builtinFunc.Set_has = typeInfo
            Builtin_builtinFunc.Set_has_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "len" {
            Builtin_builtinFunc.Set_len = typeInfo
            Builtin_builtinFunc.Set_len_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "or" {
            Builtin_builtinFunc.Set_or = typeInfo
            Builtin_builtinFunc.Set_or_sym = symInfo
            Builtin_builtinFunc.FP.register(_env, symInfo)
        } else if _switch0 == "sub" {
            Builtin_builtinFunc.Set_sub = typeInfo
            Builtin_builtinFunc.Set_sub_sym = symInfo
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
    } else if _switch0 == "List" {
        Builtin_process_List(_env)
    } else if _switch0 == "Array" {
        Builtin_process_Array(_env)
    } else if _switch0 == "Set" {
        Builtin_process_Set(_env)
    } else if _switch0 == "math" {
        Builtin_process_math(_env)
    } else if _switch0 == "debug" {
        Builtin_process_debug(_env)
    } else if _switch0 == "Nilable" {
        Builtin_process_Nilable(_env)
    }
}

// 217: decl @lune.@base.@Builtin.getBuiltInInfo
func Builtin_getBuiltInInfo_5_(_env *LnsEnv) *LnsList {
    return NewLnsList([]LnsAny{NewLnsMap( map[LnsAny]LnsAny{"__Er":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"get_txt":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"":NewLnsMap( map[LnsAny]LnsAny{"__join":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&__Runner"}),"ret":NewLnsList([]LnsAny{""}),}),"__run":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__Runner", "int", "str!"}),"ret":NewLnsList([]LnsAny{"bool"}),}),"__serr":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"__Er"}),}),"_fcall":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"form", "&..."}),"ret":NewLnsList([]LnsAny{""}),}),"_kind":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"int"}),}),"_load":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"collectgarbage":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),}),"error":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"__"}),}),"expandLuavalMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"Luaval<&stem>!"}),"ret":NewLnsList([]LnsAny{"&stem!"}),}),"loadfile":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"print":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{}),}),"require":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>"}),}),"tonumber":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!"}),"ret":NewLnsList([]LnsAny{"real!"}),}),"tostring":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem"}),"ret":NewLnsList([]LnsAny{"str"}),}),"type":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem!"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"iStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"oStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__pipe<T>":NewLnsMap( map[LnsAny]LnsAny{"get":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"method"}),}),"put":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__lns.runMode":NewLnsMap( map[LnsAny]LnsAny{"Queue":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"int"}),}),"Skip":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"int"}),}),"Sync":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"int"}),}),"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"class"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__lns.runtime":NewLnsMap( map[LnsAny]LnsAny{"dumpLog":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"oStream"}),"ret":NewLnsList([]LnsAny{}),}),"enableLog":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"bool"}),"ret":NewLnsList([]LnsAny{}),}),"log":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__lns.capability":NewLnsMap( map[LnsAny]LnsAny{"async":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"bool"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__lns.Sync.Flag":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"set":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"wait":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__lns.Sync":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"class"}),}),"_createPipe":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__exp", "int"}),"ret":NewLnsList([]LnsAny{"__pipe!"}),"type":NewLnsList([]LnsAny{"macro"}),}),"createFlag":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"__lns.Sync.Flag!"}),}),"createProcesser":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"__Processor"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"luaStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"implements":NewLnsList([]LnsAny{"iStream", "oStream"}),"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"seek":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"int!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Mapping":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"_toMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"&Map<str,&stem>"}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__Runner":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"run":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__Processor":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"implements":NewLnsList([]LnsAny{"__Runner"}),"type":NewLnsList([]LnsAny{"interface"}),}),"end":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"io":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"open":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str!"}),"ret":NewLnsList([]LnsAny{"luaStream!", "str!"}),}),"popen":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<luaStream>!"}),}),"stderr":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),"stdin":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"iStream"}),}),"stdout":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"package":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"path":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"str"}),}),"searchpath":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"os":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"clock":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"real"}),}),"date":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str!", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>!"}),}),"difftime":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem", "&stem"}),"ret":NewLnsList([]LnsAny{"int"}),}),"exit":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"__"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"bool!", "str!"}),}),"rename":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),}),"time":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"string":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"dump":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Luaval<form>", "bool!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "&..."}),"ret":NewLnsList([]LnsAny{"str"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"str"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"str":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"implements":NewLnsList([]LnsAny{"Mapping"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int!>"}),"type":NewLnsList([]LnsAny{"method"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),"type":NewLnsList([]LnsAny{"method"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),"type":NewLnsList([]LnsAny{"method"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),"type":NewLnsList([]LnsAny{"method"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"replace":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"List<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"insert":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{""}),"type":NewLnsList([]LnsAny{"mut"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Array<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Set<T>":NewLnsMap( map[LnsAny]LnsAny{"add":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"and":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"clone":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"method"}),}),"del":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"has":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"method"}),}),"len":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"int"}),"type":NewLnsList([]LnsAny{"method"}),}),"or":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"math":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"random":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"real"}),}),"randomseed":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"debug":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"getinfo":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"Map<str,stem>!"}),}),"getlocal":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int"}),"ret":NewLnsList([]LnsAny{"str!", "stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Nilable<_T>":NewLnsMap( map[LnsAny]LnsAny{"val":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"_T!"}),"type":NewLnsList([]LnsAny{"method"}),}),}),})})
}

// 220: decl @lune.@base.@Builtin.getBuiltinAlgeInfo
func Builtin_getBuiltinAlgeInfo_6_(_env *LnsEnv) *LnsList {
    return NewLnsList([]LnsAny{NewLnsMap( map[LnsAny]LnsAny{"__Ret<T1,T2>":NewLnsList([]LnsAny{NewLnsMap( map[LnsAny]LnsAny{"Ok":NewLnsList([]LnsAny{"T1"}),}), NewLnsMap( map[LnsAny]LnsAny{"Err":NewLnsList([]LnsAny{"T2"}),})}),})})
}


























// 473: decl @lune.@base.@Builtin.Builtin.getTypeInfo
func (self *Builtin_Builtin) getTypeInfo(_env *LnsEnv, typeName string) *Ast_TypeInfo {
    if _switch0 := typeName; _switch0 == "_T" {
        return Ast_builtinTypeBox.FP.Get_boxingType(_env)
    } else if _switch0 == "_T!" {
        return Ast_builtinTypeBox.FP.Get_boxingType(_env).FP.Get_nilableTypeInfo(_env)
    }
    var Builtin_getTypeInfoFromScope func(_env *LnsEnv, scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny
    Builtin_getTypeInfoFromScope = func(_env *LnsEnv, scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny {
        var Builtin_getTypeInfo func(_env *LnsEnv, workScope *Ast_Scope) LnsAny
        Builtin_getTypeInfo = func(_env *LnsEnv, workScope *Ast_Scope) LnsAny {
            var nameList *LnsList
            nameList = Util_splitModule(_env, symbol)
            if nameList.Len() <= 1{
                return workScope.FP.GetTypeInfo(_env, symbol, workScope, false, Ast_ScopeAccess__Normal)
            } else { 
                for _index, _name := range( nameList.Items ) {
                    index := _index + 1
                    name := _name.(string)
                    var workTypeInfo LnsAny
                    if index == 1{
                        workTypeInfo = workScope.FP.GetTypeInfo(_env, name, workScope, false, Ast_ScopeAccess__Normal)
                    } else { 
                        workTypeInfo = workScope.FP.GetTypeInfoChild(_env, name)
                    }
                    if workTypeInfo != nil{
                        workTypeInfo_633 := workTypeInfo.(*Ast_TypeInfo)
                        if index == nameList.Len(){
                            return workTypeInfo_633
                        }
                        {
                            __exp := workTypeInfo_633.FP.Get_scope(_env)
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
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if genType.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate{
                validGenType = true
                break
            }
        }
        if validGenType{
            if _switch0 := typeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Map {
                if genTypeList.Len() != 2{
                    Util_err(_env, _env.GetVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                }
                var keyType *Ast_TypeInfo
                keyType = genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var valType *Ast_TypeInfo
                valType = genTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                return self.processInfo.FP.CreateMap(_env, Ast_AccessMode__Pub, typeInfo.FP.Get_parentInfo(_env), keyType, valType, typeInfo.FP.Get_mutMode(_env))
            } else if _switch0 == Ast_TypeInfoKind__Ext {
                self.hasLuaval = true
                if genTypeList.Len() != 1{
                    Util_err(_env, _env.GetVM().String_format("illegal param -- %d", []LnsAny{genTypeList.Len()}))
                }
                switch _matchExp0 := self.processInfo.FP.CreateLuaval(_env, genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true).(type) {
                case *Ast_LuavalResult__OK:
                    workType := _matchExp0.Val1
                    if self.ctrl_info.ValidLuaval{
                        return workType
                    }
                    return genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                case *Ast_LuavalResult__Err:
                    mess := _matchExp0.Val1
                    Util_err(_env, mess)
                }
            } else if _switch0 == Ast_TypeInfoKind__DDD {
                if genTypeList.Len() != 1{
                    Util_err(_env, _env.GetVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                }
                return &self.processInfo.FP.CreateDDD(_env, genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true, false).Ast_TypeInfo
            } else {
                Util_err(_env, _env.GetVM().String_format("not support type -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        return typeInfo
    }
    var mutable bool
    mutable = true
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(typeName,"^&", nil, nil))){
        mutable = false
        typeName = Builtin_convExp0_13851(Lns_2DDD(_env.GetVM().String_gsub(typeName,"^&", "")))
    }
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    var endIndex LnsAny
    _,endIndex = Builtin_convExp0_13876(Lns_2DDD(_env.GetVM().String_find(typeName,"[%w%.]+<", nil, nil)))
    var suffix string
    suffix = ""
    if endIndex != nil{
        endIndex_669 := endIndex.(LnsInt)
        var genTypeName string
        genTypeName = _env.GetVM().String_sub(typeName,endIndex_669 + 1, nil)
        for  {
            {
                _tailIndex := Builtin_convExp0_13970(Lns_2DDD(_env.GetVM().String_find(genTypeName,"[,>]", nil, nil)))
                if !Lns_IsNil( _tailIndex ) {
                    tailIndex := _tailIndex.(LnsInt)
                    var genType *Ast_TypeInfo
                    genType = self.FP.getTypeInfo(_env, _env.GetVM().String_sub(genTypeName,1, tailIndex - 1))
                    genTypeList.Insert(Ast_TypeInfo2Stem(genType))
                    genTypeName = _env.GetVM().String_sub(genTypeName,tailIndex + 1, nil)
                } else {
                    suffix = _env.GetVM().String_sub(genTypeName,1, nil)
                    break
                }
            }
        }
        typeName = _env.GetVM().String_sub(typeName,1, endIndex_669 - 1) + suffix
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(typeName,"!$", nil, nil))){
        var orgTypeName string
        orgTypeName = Builtin_convExp0_14025(Lns_2DDD(_env.GetVM().String_gsub(typeName,"!$", "")))
        {
            __exp := Builtin_getTypeInfoFromScope(_env, self.transUnit.FP.Get_scope(_env), orgTypeName, genTypeList)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                typeInfo = _exp
            } else {
                Util_err(_env, _env.GetVM().String_format("not found builtin -- %s", []LnsAny{orgTypeName}))
            }
        }
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
    } else { 
        {
            __exp := Builtin_getTypeInfoFromScope(_env, self.transUnit.FP.Get_scope(_env), typeName, genTypeList)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                typeInfo = _exp
            } else {
                Util_err(_env, _env.GetVM().String_format("not found builtin -- %s", []LnsAny{typeName}))
            }
        }
    }
    if mutable{
        return typeInfo
    }
    typeInfo = self.modifier.FP.CreateModifier(_env, typeInfo, Ast_MutMode__IMut)
    return typeInfo
}
// 626: decl @lune.@base.@Builtin.Builtin.processField
func (self *Builtin_Builtin) processField(_env *LnsEnv, name string,fieldName string,info *LnsMap,parentInfo *Ast_TypeInfo) {
    self.hasLuaval = false
    if self.targetLuaVer.FP.IsSupport(_env, _env.GetVM().String_format("%s.%s", []LnsAny{name, fieldName})){
        if _switch1 := _env.NilAccFin( _env.NilAccPush(
        info.Get("type")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsList).GetAt(1))); _switch1 == "var" {
            var symbol *Ast_SymbolInfo
            symbol = Lns_unwrap( Lns_car(self.transUnit.FP.Get_scope(_env).FP.Add(_env, self.processInfo, Ast_SymbolKind__Var, false, true, fieldName, nil, self.FP.getTypeInfo(_env, Lns_unwrap( _env.NilAccFin( _env.NilAccPush(
            info.Get("typeInfo")) && 
            _env.NilAccPush( _env.NilAccPop().(*LnsList).GetAt(1)))).(string)), Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_4_(_env, name, fieldName, symbol)
        } else {
            var argTypeList *LnsList
            argTypeList = NewLnsList([]LnsAny{})
            for _, _argType := range( Lns_unwrap( info.Get("arg")).(*LnsList).Items ) {
                argType := _argType.(string)
                argTypeList.Insert(Ast_TypeInfo2Stem(self.FP.getTypeInfo(_env, argType)))
            }
            var retTypeList *LnsList
            retTypeList = NewLnsList([]LnsAny{})
            for _, _retType := range( Lns_unwrap( info.Get("ret")).(*LnsList).Items ) {
                retType := _retType.(string)
                var retTypeInfo *Ast_TypeInfo
                retTypeInfo = self.FP.getTypeInfo(_env, retType)
                retTypeList.Insert(Ast_TypeInfo2Stem(retTypeInfo))
            }
            var funcType LnsAny
            funcType = _env.NilAccFin( _env.NilAccPush(
            info.Get("type")) && 
            _env.NilAccPush( _env.NilAccPop().(*LnsList).GetAt(1)))
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
            var typeInfo *Ast_NormalTypeInfo
            typeInfo = self.processInfo.FP.CreateFuncAsync(_env, abstractFlag, true, scope, kind, parentInfo, Ast_getBuiltinMut(_env, parentInfo).FP, false, true, staticFlag, accessMode, fieldName, asyncMode, nil, argTypeList, retTypeList, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mutable) &&
                _env.SetStackVal( Ast_MutMode__Mut) ||
                _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
            if self.hasLuaval{
                Builtin_builtinFunc.FP.addLuavalFunc(_env, &typeInfo.Ast_TypeInfo)
            }
            self.transUnit.FP.PopScope(_env)
            Builtin_builtinFunc.FP.Get_allFuncTypeSet(_env).Add(Ast_NormalTypeInfo2Stem(typeInfo))
            Ast_addBuiltinMut(_env, &typeInfo.Ast_TypeInfo, scope)
            var symInfo *Ast_SymbolInfo
            symInfo = Lns_unwrap( Lns_car(self.transUnit.FP.Get_scope(_env).FP.Add(_env, self.processInfo, symbolKind, false, kind == Ast_TypeInfoKind__Func, fieldName, nil, &typeInfo.Ast_TypeInfo, accessMode, staticFlag, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mutable) &&
                _env.SetStackVal( Ast_MutMode__Mut) ||
                _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false))).(*Ast_SymbolInfo)
            Builtin_setupBuiltinTypeInfo_4_(_env, name, fieldName, symInfo)
        }
    }
}
// 741: decl @lune.@base.@Builtin.Builtin.registClass
func (self *Builtin_Builtin) registClass(_env *LnsEnv, nameList *LnsList,name2FieldInfo *LnsMap,pos Types_Position,genTypeList *LnsList) *Ast_TypeInfo {
    var classKind LnsInt
    classKind = TransUnitIF_DeclClassMode__Class
    {
        _types := _env.NilAccFin( _env.NilAccPush(
        name2FieldInfo.Get("__attrib")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsMap).Get("type")))
        if !Lns_IsNil( _types ) {
            types := _types.(*LnsList)
            if types.Len() > 0{
                if _switch0 := types.GetAt(1).(string); _switch0 == "interface" {
                    classKind = TransUnitIF_DeclClassMode__Interface
                } else if _switch0 == "module" {
                    classKind = TransUnitIF_DeclClassMode__Module
                }
            }
        }
    }
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    {
        __exp := _env.NilAccFin( _env.NilAccPush(
        name2FieldInfo.Get("__attrib")) && 
        _env.NilAccPush( _env.NilAccPop().(*LnsMap).Get("implements")))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*LnsList)
            for _, _ifname := range( _exp.Items ) {
                ifname := _ifname.(string)
                var ifType *Ast_TypeInfo
                ifType = self.FP.getTypeInfo(_env, ifname)
                interfaceList.Insert(Ast_TypeInfo2Stem(ifType))
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
        name := _name.(string)
        parentInfo = Builtin_registerClass(_env, name, index, index == nameList.Len())
    }
    return parentInfo
}
// 811: decl @lune.@base.@Builtin.Builtin.createGenType
func (self *Builtin_Builtin) createGenType(_env *LnsEnv, typeName string,genTypeList *LnsList) string {
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
                        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Lns_car(self.processInfo.FP.CreateAlternate(_env, true, genTypeList.Len() + 1, token, Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil, nil)).(*Ast_AlternateTypeInfo)))
                    }
                }
            }
    }
    return name
}
// 834: decl @lune.@base.@Builtin.Builtin.createBuiltinAlge
func (self *Builtin_Builtin) createBuiltinAlge(_env *LnsEnv, algesymList *LnsList) {
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    for _, _algeMap := range( algesymList.Items ) {
        algeMap := _algeMap.(*LnsMap)
        for _genAlgeName, _valMapList := range( algeMap.Items ) {
            genAlgeName := _genAlgeName.(string)
            valMapList := _valMapList.(*LnsList)
            var genTypeList *LnsList
            genTypeList = NewLnsList([]LnsAny{})
            var algeName string
            algeName = self.FP.createGenType(_env, genAlgeName, genTypeList)
            var typename2typeInfo *LnsMap
            typename2typeInfo = NewLnsMap( map[LnsAny]LnsAny{})
            for _, _genType := range( genTypeList.Items ) {
                genType := _genType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
                typename2typeInfo.Set(genType.FP.Get_rawTxt(_env),&genType.Ast_TypeInfo)
            }
            var algeScope *Ast_Scope
            algeScope = NewAst_Scope(_env, self.processInfo, self.transUnit.FP.Get_scope(_env), Ast_ScopeKind__Class, nil, nil)
            var algeTypeInfo *Ast_AlgeTypeInfo
            algeTypeInfo = self.processInfo.FP.CreateAlge(_env, algeScope, parentInfo, Ast_getBuiltinMut(_env, parentInfo).FP, false, Ast_AccessMode__Pub, algeName, genTypeList)
            Ast_addBuiltinMut(_env, &algeTypeInfo.Ast_TypeInfo, algeScope)
            Builtin_builtinFunc.FP.registerAlge(_env, &algeTypeInfo.Ast_TypeInfo)
            self.transUnit.FP.Get_scope(_env).FP.AddAlge(_env, self.processInfo, Ast_AccessMode__Pub, algeName, nil, &algeTypeInfo.Ast_TypeInfo)
            for _, _valMap := range( valMapList.Items ) {
                valMap := _valMap.(*LnsMap)
                for _valName, _valTypeNameList := range( valMap.Items ) {
                    valName := _valName.(string)
                    valTypeNameList := _valTypeNameList.(*LnsList)
                    var algeValSym *Ast_SymbolInfo
                    algeValSym = Lns_unwrap( Lns_car(algeScope.FP.AddAlgeVal(_env, self.processInfo, valName, Types_nonePos, &algeTypeInfo.Ast_TypeInfo))).(*Ast_SymbolInfo)
                    var valTypeList *LnsList
                    valTypeList = NewLnsList([]LnsAny{})
                    for _, _typeName := range( valTypeNameList.Items ) {
                        typeName := _typeName.(string)
                        valTypeList.Insert(Ast_TypeInfo2Stem(_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( typename2typeInfo.Get(typeName)) ||
                            _env.SetStackVal( self.FP.getTypeInfo(_env, typeName)) ).(*Ast_TypeInfo)))
                    }
                    var algeValInfo *Ast_AlgeValInfo
                    algeValInfo = NewAst_AlgeValInfo(_env, algeValSym.FP.Get_name(_env), valTypeList, algeTypeInfo, algeValSym)
                    algeTypeInfo.FP.AddValInfo(_env, algeValInfo)
                }
            }
        }
    }
}
// 880: decl @lune.@base.@Builtin.Builtin.registBuiltInScope
func (self *Builtin_Builtin) RegistBuiltInScope(_env *LnsEnv) *Builtin_BuiltinFuncType {
    if Builtin_readyBuiltin{
        return Builtin_builtinFunc
    }
    Builtin_readyBuiltin = true
    var builtInInfo *LnsList
    builtInInfo = Builtin_getBuiltInInfo_5_(_env)
    var builtinModuleName2Scope *LnsMap
    builtinModuleName2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    var mapType *Ast_TypeInfo
    mapType = self.processInfo.FP.CreateMap(_env, Ast_AccessMode__Pub, Ast_headTypeInfo, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "_ENV", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "_G", nil, mapType, Ast_MutMode__IMutRe, true)
    self.transUnit.FP.Get_scope(_env).FP.AddVar(_env, self.processInfo, Ast_AccessMode__Global, "__line__", nil, Ast_builtinTypeInt, Ast_MutMode__IMut, true)
    var pos Types_Position
    pos = NewTypes_Position(_env, 0, 0, "@builtin@", nil)
    for _, _builtinClassInfo := range( builtInInfo.Items ) {
        builtinClassInfo := _builtinClassInfo.(*LnsMap)
        for _className, _name2FieldInfo := range( builtinClassInfo.Items ) {
            className := _className.(string)
            name2FieldInfo := _name2FieldInfo.(*LnsMap)
            var name string
            name = className
            var genTypeList *LnsList
            genTypeList = NewLnsList([]LnsAny{})
            if _switch0 := className; _switch0 == "List<T>" {
                name = "List"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeList.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Array<T>" {
                name = "Array"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeArray.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Set<T>" {
                name = "Set"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeSet.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Map<K,V>" {
                name = "Map"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeMap.FP.Get_itemTypeInfoList(_env))
            } else if _switch0 == "Nilable<_T>" {
                name = "Nilable"
                Builtin_Builtin_registBuiltInScope__processCopyAlterList_0_(_env, genTypeList, Ast_builtinTypeBox.FP.Get_itemTypeInfoList(_env))
            } else {
                name = self.FP.createGenType(_env, className, genTypeList)
            }
            var nameList *LnsList
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
                    for _, _fieldName := range( __forsortSorted0.Items ) {
                        info := __forsortCollection0.Items[ _fieldName ].(*LnsMap)
                        fieldName := _fieldName.(string)
                        if _switch1 := fieldName; _switch1 == "__attrib" {
                        } else {
                            self.FP.processField(_env, Lns_car(_env.GetVM().String_gsub(name,"%.", "_")).(string), fieldName, info, parentInfo)
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
    var threadSafeSet *LnsSet
    threadSafeSet = NewLnsSet([]LnsAny{Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_error), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_print), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_type), Ast_TypeInfo2Stem(Builtin_builtinFunc.Lns_tonumber), Ast_TypeInfo2Stem(Builtin_builtinFunc.Io_open), Ast_TypeInfo2Stem(Builtin_builtinFunc.Set_has), Ast_TypeInfo2Stem(Builtin_builtinFunc.Set_add)})
    for _typeInfo := range( Builtin_builtinFunc.FP.Get_allFuncTypeSet(_env).Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(threadSafeSet.Has(Ast_TypeInfo2Stem(typeInfo))){
            Builtin_builtinFunc.FP.Get_needThreadingTypes(_env).Add(Ast_TypeInfo2Stem(typeInfo))
        }
    }
    return Builtin_builtinFunc
}
// 1023: decl @lune.@base.@Builtin.Builtin.getBuiltInFuncType
func (self *Builtin_Builtin) GetBuiltInFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType {
    if Builtin_readyBuiltin{
        return Builtin_builtinFunc
    }
    Util_err(_env, "builtinFunc is not ready")
// insert a dummy
    return nil
}
// 175: decl @lune.@base.@Builtin.BuiltinFuncType.addLuavalFunc
func (self *Builtin_BuiltinFuncType) addLuavalFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    self.luavalFuncTypeSet.Add(Ast_TypeInfo2Stem(typeInfo))
}
// 179: decl @lune.@base.@Builtin.BuiltinFuncType.isLuavalFunc
func (self *Builtin_BuiltinFuncType) IsLuavalFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return self.luavalFuncTypeSet.Has(Ast_TypeInfo2Stem(typeInfo))
}
// 183: decl @lune.@base.@Builtin.BuiltinFuncType.register
func (self *Builtin_BuiltinFuncType) register(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) {
    self.allSymbol.Insert(Ast_SymbolInfo2Stem(symbolInfo))
    self.allSymbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
}
// 187: decl @lune.@base.@Builtin.BuiltinFuncType.registerClass
func (self *Builtin_BuiltinFuncType) registerClass(_env *LnsEnv, classInfo *Ast_TypeInfo) {
    self.allClass.Insert(Ast_TypeInfo2Stem(classInfo))
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
    } else if _switch0 == "List" {
        self.List_ = classInfo
    } else if _switch0 == "Array" {
        self.Array_ = classInfo
    } else if _switch0 == "Set" {
        self.Set_ = classInfo
    } else if _switch0 == "math" {
        self.Math_ = classInfo
    } else if _switch0 == "debug" {
        self.Debug_ = classInfo
    } else if _switch0 == "Nilable" {
        self.Nilable_ = classInfo
    }
}
// 193: decl @lune.@base.@Builtin.BuiltinFuncType.registerAlge
func (self *Builtin_BuiltinFuncType) registerAlge(_env *LnsEnv, algeInfo *Ast_TypeInfo) {
    self.allClass.Insert(Ast_TypeInfo2Stem(algeInfo))
    if _switch0 := algeInfo.FP.Get_rawTxt(_env); _switch0 == "__Ret" {
        self.G__ret_ = algeInfo
    }
}
// 466: decl @lune.@base.@Builtin.BuiltinFuncType.isStrFormFunc
func (self *Builtin_BuiltinFuncType) IsStrFormFunc(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_srcTypeInfo(_env) == self.String_format{
        return true
    }
    return false
}
// declaration Class -- Builtin
type Builtin_BuiltinMtd interface {
    createBuiltinAlge(_env *LnsEnv, arg1 *LnsList)
    createGenType(_env *LnsEnv, arg1 string, arg2 *LnsList) string
    GetBuiltInFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType
    getTypeInfo(_env *LnsEnv, arg1 string) *Ast_TypeInfo
    processField(_env *LnsEnv, arg1 string, arg2 string, arg3 *LnsMap, arg4 *Ast_TypeInfo)
    RegistBuiltInScope(_env *LnsEnv) *Builtin_BuiltinFuncType
    registClass(_env *LnsEnv, arg1 *LnsList, arg2 *LnsMap, arg3 Types_Position, arg4 *LnsList) *Ast_TypeInfo
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
// 49: DeclConstr
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
    Get_allClass(_env *LnsEnv) *LnsList
    Get_allFuncTypeSet(_env *LnsEnv) *LnsSet
    Get_allSymbol(_env *LnsEnv) *LnsList
    Get_allSymbolSet(_env *LnsEnv) *LnsSet
    Get_needThreadingTypes(_env *LnsEnv) *LnsSet
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
    allSymbol *LnsList
    allClass *LnsList
    allFuncTypeSet *LnsSet
    allSymbolSet *LnsSet
    needThreadingTypes *LnsSet
    luavalFuncTypeSet *LnsSet
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
func NewBuiltin_BuiltinFuncType(_env *LnsEnv) *Builtin_BuiltinFuncType {
    obj := &Builtin_BuiltinFuncType{}
    obj.FP = obj
    obj.InitBuiltin_BuiltinFuncType(_env)
    return obj
}
func (self *Builtin_BuiltinFuncType) Get_allSymbol(_env *LnsEnv) *LnsList{ return self.allSymbol }
func (self *Builtin_BuiltinFuncType) Get_allClass(_env *LnsEnv) *LnsList{ return self.allClass }
func (self *Builtin_BuiltinFuncType) Get_allFuncTypeSet(_env *LnsEnv) *LnsSet{ return self.allFuncTypeSet }
func (self *Builtin_BuiltinFuncType) Get_allSymbolSet(_env *LnsEnv) *LnsSet{ return self.allSymbolSet }
func (self *Builtin_BuiltinFuncType) Get_needThreadingTypes(_env *LnsEnv) *LnsSet{ return self.needThreadingTypes }
// 164: DeclConstr
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
    self.allSymbol = NewLnsList([]LnsAny{})
    self.allClass = NewLnsList([]LnsAny{})
    self.allSymbolSet = NewLnsSet([]LnsAny{})
    self.allFuncTypeSet = NewLnsSet([]LnsAny{})
    self.needThreadingTypes = NewLnsSet([]LnsAny{})
    self.luavalFuncTypeSet = NewLnsSet([]LnsAny{})
}


func Lns_Builtin_init(_env *LnsEnv) {
    if init_Builtin { return }
    init_Builtin = true
    Builtin__mod__ = "@lune.@base.@Builtin"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Parser_init(_env)
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
