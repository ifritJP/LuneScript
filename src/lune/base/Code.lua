--lune/base/Code.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Code'
local _lune = {}
if _lune0 then
   _lune = _lune0
end
if not _lune0 then
   _lune0 = _lune
end
local ID = {}
_moduleObj.ID = ID
ID._val2NameMap = {}
function ID:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ID.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ID._from( val )
   if ID._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ID.__allList = {}
function ID.get__allList()
   return ID.__allList
end

ID.nothing_define_abbr = 0
ID._val2NameMap[0] = 'nothing_define_abbr'
ID.__allList[1] = ID.nothing_define_abbr

local function format( id, mess )

   return string.format( "%05d:%s", id, mess )
end
_moduleObj.format = format
local function isMessageOf( id, mess )

   local pat = string.format( "^%05d:", id)
   if mess:find( pat ) then
      return true
   end
   
   return false
end
_moduleObj.isMessageOf = isMessageOf
return _moduleObj
