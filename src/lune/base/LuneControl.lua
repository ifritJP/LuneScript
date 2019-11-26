--lune/base/LuneControl.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@LuneControl'
local _lune = {}
if _lune2 then
   _lune = _lune2
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

if not _lune2 then
   _lune2 = _lune
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

Pragma.can_not_conv_code = { "can_not_conv_code", {{ func=_lune._toSet, nilable=false, child={ func = Code._from, nilable = false, child = {} } }}}
Pragma._name2Val["can_not_conv_code"] = Pragma.can_not_conv_code
Pragma.default__init = { "default__init"}
Pragma._name2Val["default__init"] = Pragma.default__init
Pragma.default__init_old = { "default__init_old"}
Pragma._name2Val["default__init_old"] = Pragma.default__init_old
Pragma.disable_mut_control = { "disable_mut_control"}
Pragma._name2Val["disable_mut_control"] = Pragma.disable_mut_control
Pragma.ignore_symbol_ = { "ignore_symbol_"}
Pragma._name2Val["ignore_symbol_"] = Pragma.ignore_symbol_
Pragma.load__lune_module = { "load__lune_module"}
Pragma._name2Val["load__lune_module"] = Pragma.load__lune_module


return _moduleObj
