--lune/base/Code.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Code'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune3 then
   _lune3 = _lune
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
