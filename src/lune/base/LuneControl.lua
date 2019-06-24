--lune/base/LuneControl.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@LuneControl'
local _lune = {}
if _lune1 then
   _lune = _lune1
end
if not _lune1 then
   _lune1 = _lune
end
local Pragma = {}
_moduleObj.Pragma = Pragma
Pragma._val2NameMap = {}
function Pragma:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Pragma.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Pragma._from( val )
   if Pragma._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Pragma.__allList = {}
function Pragma.get__allList()
   return Pragma.__allList
end

Pragma.default__init = 'default__init'
Pragma._val2NameMap['default__init'] = 'default__init'
Pragma.__allList[1] = Pragma.default__init
Pragma.default__init_old = 'default__init_old'
Pragma._val2NameMap['default__init_old'] = 'default__init_old'
Pragma.__allList[2] = Pragma.default__init_old
Pragma.disable_mut_control = 'disable_mut_control'
Pragma._val2NameMap['disable_mut_control'] = 'disable_mut_control'
Pragma.__allList[3] = Pragma.disable_mut_control
Pragma.ignore_symbol_ = 'ignore_symbol_'
Pragma._val2NameMap['ignore_symbol_'] = 'ignore_symbol_'
Pragma.__allList[4] = Pragma.ignore_symbol_
Pragma.load__lune_module = 'load__lune_module'
Pragma._val2NameMap['load__lune_module'] = 'load__lune_module'
Pragma.__allList[5] = Pragma.load__lune_module

return _moduleObj
