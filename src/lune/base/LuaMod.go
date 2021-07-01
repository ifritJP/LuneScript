// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_LuaMod bool
var LuaMod__mod__ string
// decl enum -- CodeKind 
type LuaMod_CodeKind = LnsInt
const LuaMod_CodeKind__Alge = 8
const LuaMod_CodeKind__AlgeMapping = 9
const LuaMod_CodeKind__Cast = 13
const LuaMod_CodeKind__Finalize = 18
const LuaMod_CodeKind__Init = 0
const LuaMod_CodeKind__InstanceOf = 12
const LuaMod_CodeKind__LazyLoad = 14
const LuaMod_CodeKind__LazyRequire = 15
const LuaMod_CodeKind__LoadModule = 4
const LuaMod_CodeKind__LoadStr51 = 5
const LuaMod_CodeKind__LoadStr52 = 6
const LuaMod_CodeKind__Mapping = 7
const LuaMod_CodeKind__NilAcc = 2
const LuaMod_CodeKind__Run = 16
const LuaMod_CodeKind__SetMapping = 10
const LuaMod_CodeKind__SetOp = 11
const LuaMod_CodeKind__StrReplace = 17
const LuaMod_CodeKind__Unpack = 1
const LuaMod_CodeKind__Unwrap = 3
var LuaMod_CodeKindList_ = NewLnsList( []LnsAny {
  LuaMod_CodeKind__Init,
  LuaMod_CodeKind__Unpack,
  LuaMod_CodeKind__NilAcc,
  LuaMod_CodeKind__Unwrap,
  LuaMod_CodeKind__LoadModule,
  LuaMod_CodeKind__LoadStr51,
  LuaMod_CodeKind__LoadStr52,
  LuaMod_CodeKind__Mapping,
  LuaMod_CodeKind__Alge,
  LuaMod_CodeKind__AlgeMapping,
  LuaMod_CodeKind__SetMapping,
  LuaMod_CodeKind__SetOp,
  LuaMod_CodeKind__InstanceOf,
  LuaMod_CodeKind__Cast,
  LuaMod_CodeKind__LazyLoad,
  LuaMod_CodeKind__LazyRequire,
  LuaMod_CodeKind__Run,
  LuaMod_CodeKind__StrReplace,
  LuaMod_CodeKind__Finalize,
})
func LuaMod_CodeKind_get__allList(_env *LnsEnv) *LnsList{
    return LuaMod_CodeKindList_
}
var LuaMod_CodeKindMap_ = map[LnsInt]string {
  LuaMod_CodeKind__Alge: "CodeKind.Alge",
  LuaMod_CodeKind__AlgeMapping: "CodeKind.AlgeMapping",
  LuaMod_CodeKind__Cast: "CodeKind.Cast",
  LuaMod_CodeKind__Finalize: "CodeKind.Finalize",
  LuaMod_CodeKind__Init: "CodeKind.Init",
  LuaMod_CodeKind__InstanceOf: "CodeKind.InstanceOf",
  LuaMod_CodeKind__LazyLoad: "CodeKind.LazyLoad",
  LuaMod_CodeKind__LazyRequire: "CodeKind.LazyRequire",
  LuaMod_CodeKind__LoadModule: "CodeKind.LoadModule",
  LuaMod_CodeKind__LoadStr51: "CodeKind.LoadStr51",
  LuaMod_CodeKind__LoadStr52: "CodeKind.LoadStr52",
  LuaMod_CodeKind__Mapping: "CodeKind.Mapping",
  LuaMod_CodeKind__NilAcc: "CodeKind.NilAcc",
  LuaMod_CodeKind__Run: "CodeKind.Run",
  LuaMod_CodeKind__SetMapping: "CodeKind.SetMapping",
  LuaMod_CodeKind__SetOp: "CodeKind.SetOp",
  LuaMod_CodeKind__StrReplace: "CodeKind.StrReplace",
  LuaMod_CodeKind__Unpack: "CodeKind.Unpack",
  LuaMod_CodeKind__Unwrap: "CodeKind.Unwrap",
}
func LuaMod_CodeKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := LuaMod_CodeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func LuaMod_CodeKind_getTxt(arg1 LnsInt) string {
    return LuaMod_CodeKindMap_[arg1];
}
// decl enum -- CastKind 
type LuaMod_CastKind = LnsInt
const LuaMod_CastKind__Class = 3
const LuaMod_CastKind__Int = 0
const LuaMod_CastKind__Real = 1
const LuaMod_CastKind__Str = 2
var LuaMod_CastKindList_ = NewLnsList( []LnsAny {
  LuaMod_CastKind__Int,
  LuaMod_CastKind__Real,
  LuaMod_CastKind__Str,
  LuaMod_CastKind__Class,
})
func LuaMod_CastKind_get__allList(_env *LnsEnv) *LnsList{
    return LuaMod_CastKindList_
}
var LuaMod_CastKindMap_ = map[LnsInt]string {
  LuaMod_CastKind__Class: "CastKind.Class",
  LuaMod_CastKind__Int: "CastKind.Int",
  LuaMod_CastKind__Real: "CastKind.Real",
  LuaMod_CastKind__Str: "CastKind.Str",
}
func LuaMod_CastKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := LuaMod_CastKindMap_[arg1]; ok { return arg1 }
    return nil
}

func LuaMod_CastKind_getTxt(arg1 LnsInt) string {
    return LuaMod_CastKindMap_[arg1];
}
var LuaMod_codeMap *LnsMap
// 479: decl @lune.@base.@LuaMod.getCode
func LuaMod_getCode(_env *LnsEnv, kind LnsInt) string {
    return Lns_unwrap( LuaMod_codeMap.Get(kind)).(string)
}

func Lns_LuaMod_init(_env *LnsEnv) {
    if init_LuaMod { return }
    init_LuaMod = true
    LuaMod__mod__ = "@lune.@base.@LuaMod"
    Lns_InitMod()
    {
        var work *LnsMap
        work = NewLnsMap( map[LnsAny]LnsAny{})
        work.Set(LuaMod_CodeKind__Init,"local _lune = {}\n")
        work.Set(LuaMod_CodeKind__Unpack,"if not table.unpack then\n   table.unpack = unpack\nend\n")
        work.Set(LuaMod_CodeKind__NilAcc,"function _lune.nilacc( val, fieldName, access, ... )\n   if not val then\n      return nil\n   end\n   if fieldName then\n      local field = val[ fieldName ]\n      if not field then\n         return nil\n      end\n      if access == \"item\" then\n         local typeId = type( field )\n         if typeId == \"table\" then\n            return field[ ... ]\n         elseif typeId == \"string\" then\n            return string.byte( field, ... )\n         end\n      elseif access == \"call\" then\n         return field( ... )\n      elseif access == \"callmtd\" then\n         return field( val, ... )\n      end\n      return field\n   end\n   if access == \"item\" then\n      local typeId = type( val )\n      if typeId == \"table\" then\n         return val[ ... ]\n      elseif typeId == \"string\" then\n         return string.byte( val, ... )\n      end\n   elseif access == \"call\" then\n      return val( ... )\n   elseif access == \"list\" then\n      local list, arg = ...\n      if not list then\n         return nil\n      end\n      return val( list, arg )\n   end\n   error( string.format( \"illegal access -- %s\", access ) )\nend\n")
        work.Set(LuaMod_CodeKind__Unwrap,"function _lune.unwrap( val )\n   if val == nil then\n      __luneScript:error( 'unwrap val is nil' )\n   end\n   return val\nend\nfunction _lune.unwrapDefault( val, defval )\n   if val == nil then\n      return defval\n   end\n   return val\nend\n")
        work.Set(LuaMod_CodeKind__LoadModule,"function _lune.loadModule( mod )\n   if __luneScript then\n      return  __luneScript:loadModule( mod )\n   end\n   return require( mod )\nend\n")
        work.Set(LuaMod_CodeKind__LoadStr51,"function _lune.loadstring51( txt, env )\n   local func = loadstring( txt )\n   if func and env then\n      setfenv( func, env )\n   end\n   return func\nend\n")
        work.Set(LuaMod_CodeKind__LoadStr52,"function _lune.loadstring52( txt, env )\n   if not env then\n      return load( txt )\n   end\n   return load( txt, \"\", \"bt\", env )\nend\n")
        work.Set(LuaMod_CodeKind__Mapping,"function _lune._toStem( val )\n   return val\nend\nfunction _lune._toInt( val )\n   if type( val ) == \"number\" then\n      return math.floor( val )\n   end\n   return nil\nend\nfunction _lune._toReal( val )\n   if type( val ) == \"number\" then\n      return val\n   end\n   return nil\nend\nfunction _lune._toBool( val )\n   if type( val ) == \"boolean\" then\n      return val\n   end\n   return nil\nend\nfunction _lune._toStr( val )\n   if type( val ) == \"string\" then\n      return val\n   end\n   return nil\nend\nfunction _lune._toList( val, toValInfoList )\n   if type( val ) == \"table\" then\n      local tbl = {}\n      local toValInfo = toValInfoList[ 1 ]\n      for index, mem in ipairs( val ) do\n         local memval, mess = toValInfo.func( mem, toValInfo.child )\n         if memval == nil and not toValInfo.nilable then\n            if mess then\n              return nil, string.format( \"%d.%s\", index, mess )\n            end\n            return nil, index\n         end\n         tbl[ index ] = memval\n      end\n      return tbl\n   end\n   return nil\nend\nfunction _lune._toMap( val, toValInfoList )\n   if type( val ) == \"table\" then\n      local tbl = {}\n      local toKeyInfo = toValInfoList[ 1 ]\n      local toValInfo = toValInfoList[ 2 ]\n      for key, mem in pairs( val ) do\n         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )\n         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )\n         if mapKey == nil or mapVal == nil then\n            if mapKey == nil then\n               return nil\n            end\n            if keySub == nil then\n               return nil, mapKey\n            end\n            return nil, string.format( \"%s.%s\", mapKey, keySub)\n         end\n         tbl[ mapKey ] = mapVal\n      end\n      return tbl\n   end\n   return nil\nend\nfunction _lune._fromMap( obj, map, memInfoList )\n   if type( map ) ~= \"table\" then\n      return false\n   end\n   for index, memInfo in ipairs( memInfoList ) do\n      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )\n      if val == nil and not memInfo.nilable then\n         return false, key and string.format( \"%s.%s\", memInfo.name, key) or memInfo.name\n      end\n      obj[ memInfo.name ] = val\n   end\n   return true\nend\n")
        work.Set(LuaMod_CodeKind__SetOp,"function _lune._Set_or( setObj, otherSet )\n   for val in pairs( otherSet ) do\n      setObj[ val ] = true\n   end\n   return setObj\nend\nfunction _lune._Set_and( setObj, otherSet )\n   local delValList = {}\n   for val in pairs( setObj ) do\n      if not otherSet[ val ] then\n         table.insert( delValList, val )\n      end\n   end\n   for index, val in ipairs( delValList ) do\n      setObj[ val ] = nil\n   end\n   return setObj\nend\nfunction _lune._Set_has( setObj, val )\n   return setObj[ val ] ~= nil\nend\nfunction _lune._Set_sub( setObj, otherSet )\n   local delValList = {}\n   for val in pairs( setObj ) do\n      if otherSet[ val ] then\n         table.insert( delValList, val )\n      end\n   end\n   for index, val in ipairs( delValList ) do\n      setObj[ val ] = nil\n   end\n   return setObj\nend\nfunction _lune._Set_len( setObj )\n   local total = 0\n   for val in pairs( setObj ) do\n      total = total + 1\n   end\n   return total\nend\nfunction _lune._Set_clone( setObj )\n   local obj = {}\n   for val in pairs( setObj ) do\n      obj[ val ] = true\n   end\n   return obj\nend\n")
        work.Set(LuaMod_CodeKind__SetMapping,"function _lune._toSet( val, toKeyInfo )\n   if type( val ) == \"table\" then\n      local tbl = {}\n      for key, mem in pairs( val ) do\n         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )\n         local mapVal = _lune._toBool( mem )\n         if mapKey == nil or mapVal == nil then\n            if mapKey == nil then\n               return nil\n            end\n            if keySub == nil then\n               return nil, mapKey\n            end\n            return nil, string.format( \"%s.%s\", mapKey, keySub)\n         end\n         tbl[ mapKey ] = mapVal\n      end\n      return tbl\n   end\n   return nil\nend\n")
        work.Set(LuaMod_CodeKind__AlgeMapping,"function _lune._fromList( obj, list, memInfoList )\n   if type( list ) ~= \"table\" then\n      return false\n   end\n   for index, memInfo in ipairs( memInfoList ) do\n      local val, key = memInfo.func( list[ index ], memInfo.child )\n      if val == nil and not memInfo.nilable then\n         return false, key and string.format( \"%s[%s]\", memInfo.name, key) or memInfo.name\n      end\n      obj[ index ] = val\n   end\n   return true\nend\nfunction _lune._AlgeFrom( Alge, val )\n   local work = Alge._name2Val[ val[ 1 ] ]\n   if not work then\n      return nil\n   end\n   if #work == 1 then\n     return work\n   end\n   local paramList = {}\n   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )\n   if not result then\n      return nil, mess\n   end\n   return { work[ 1 ], paramList }\nend\n")
        work.Set(LuaMod_CodeKind__Alge,"function _lune.newAlge( kind, vals )\n   local memInfoList = kind[ 2 ]\n   if not memInfoList then\n      return kind\n   end\n   return { kind[ 1 ], vals }\nend\n")
        work.Set(LuaMod_CodeKind__InstanceOf,"function _lune.__isInstanceOf( obj, class )\n   while obj do\n      local meta = getmetatable( obj )\n      if not meta then\n\t return false\n      end\n      local indexTbl = meta.__index\n      if indexTbl == class then\n\t return true\n      end\n      if meta.ifList then\n         for index, ifType in ipairs( meta.ifList ) do\n            if ifType == class then\n               return true\n            end\n            if _lune.__isInstanceOf( ifType, class ) then\n               return true\n            end\n         end\n      end\n      obj = indexTbl\n   end\n   return false\nend\n")
        work.Set(LuaMod_CodeKind__Cast,_env.GetVM().String_format("function _lune.__Cast( obj, kind, class )\n   if kind == %d then -- int\n      if type( obj ) ~= \"number\" then\n         return nil\n      end\n      if math.floor( obj ) ~= obj then\n         return nil\n      end\n      return obj\n   elseif kind == %d then -- real\n      if type( obj ) ~= \"number\" then\n         return nil\n      end\n      return obj\n   elseif kind == %d then -- str\n      if type( obj ) ~= \"string\" then\n         return nil\n      end\n      return obj\n   elseif kind == %d then -- class\n      return _lune.__isInstanceOf( obj, class ) and obj or nil\n   end\n   return nil\nend\n", []LnsAny{LuaMod_CastKind__Int, LuaMod_CastKind__Real, LuaMod_CastKind__Str, LuaMod_CastKind__Class}))
        work.Set(LuaMod_CodeKind__LazyLoad,"function _lune._lazyImport( modName )\n  local mod\n  return function()\n    if mod then\n       return mod\n    end\n    mod = _lune.loadModule( modName )\n    return mod\n  end\nend\n")
        work.Set(LuaMod_CodeKind__LazyRequire,"function _lune._lazyRequire( modName )\n  local mod\n  return function()\n    if mod then\n       return mod\n    end\n    mod = require( modName )\n    return mod\n  end\nend\n")
        work.Set(LuaMod_CodeKind__Run,"function _lune._run( runner, mod )\n    if mod == 2 then\n      return false\n    end\n    runner:run()\n    return true\nend\n")
        work.Set(LuaMod_CodeKind__StrReplace,"function _lune.replace( txt, src, dst )\n   local result = \"\"\n   local index = 1\n   while index <= #txt do\n      local findIndex = string.find( txt, src, index, true )\n      if not findIndex then\n         result = result .. string.sub( txt, index )\n         break\n      end\n      if findIndex ~= index then\n         result = result .. (string.sub( txt, index, findIndex - 1 ) .. dst)\n      else\n         result = result .. dst\n      end\n      index = findIndex + #src\n   end\n   return result\nend\n")
        work.Set(LuaMod_CodeKind__Finalize,"return _lune\n")
        LuaMod_codeMap = work
    }
}
func init() {
    init_LuaMod = false
}
