--lune/base/LuneControl.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@LuneControl'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
end

function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end

if not _lune6 then
   _lune6 = _lune
end


local Code = {}
_moduleObj.Code = Code
Code._val2NameMap = {}
function Code:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Code.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Code._from( val )
   if Code._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Code.__allList = {}
function Code.get__allList()
   return Code.__allList
end

Code.Lua = 'lua'
Code._val2NameMap['lua'] = 'Lua'
Code.__allList[1] = Code.Lua
Code.C = 'c'
Code._val2NameMap['c'] = 'C'
Code.__allList[2] = Code.C
Code.Go = 'go'
Code._val2NameMap['go'] = 'Go'
Code.__allList[3] = Code.Go
Code.Python = 'py'
Code._val2NameMap['py'] = 'Python'
Code.__allList[4] = Code.Python


local Pragma = {}
Pragma._name2Val = {}
_moduleObj.Pragma = Pragma
function Pragma:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "Pragma.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function Pragma._from( val )
   return _lune._AlgeFrom( Pragma, val )
end

Pragma.default__init = { "default__init"}
Pragma._name2Val["default__init"] = Pragma.default__init
Pragma.default__init_old = { "default__init_old"}
Pragma._name2Val["default__init_old"] = Pragma.default__init_old
Pragma.default_async_all = { "default_async_all"}
Pragma._name2Val["default_async_all"] = Pragma.default_async_all
Pragma.default_async_func = { "default_async_func"}
Pragma._name2Val["default_async_func"] = Pragma.default_async_func
Pragma.default_async_this_class = { "default_async_this_class"}
Pragma._name2Val["default_async_this_class"] = Pragma.default_async_this_class
Pragma.default_noasync_this_class = { "default_noasync_this_class"}
Pragma._name2Val["default_noasync_this_class"] = Pragma.default_noasync_this_class
Pragma.disable_mut_control = { "disable_mut_control"}
Pragma._name2Val["disable_mut_control"] = Pragma.disable_mut_control
Pragma.ignore_symbol_ = { "ignore_symbol_"}
Pragma._name2Val["ignore_symbol_"] = Pragma.ignore_symbol_
Pragma.limit_conv_code = { "limit_conv_code", {{ func=_lune._toSet, nilable=false, child={ func = Code._from, nilable = false, child = {} } }}}
Pragma._name2Val["limit_conv_code"] = Pragma.limit_conv_code
Pragma.load__lune_module = { "load__lune_module"}
Pragma._name2Val["load__lune_module"] = Pragma.load__lune_module
Pragma.run_async_pipe = { "run_async_pipe"}
Pragma._name2Val["run_async_pipe"] = Pragma.run_async_pipe
Pragma.single_phase_ast = { "single_phase_ast"}
Pragma._name2Val["single_phase_ast"] = Pragma.single_phase_ast
Pragma.use_async = { "use_async"}
Pragma._name2Val["use_async"] = Pragma.use_async
Pragma.use_macro_special_var = { "use_macro_special_var"}
Pragma._name2Val["use_macro_special_var"] = Pragma.use_macro_special_var


return _moduleObj
