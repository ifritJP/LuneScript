--lune/base/Str.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Str'
local _lune = {}
if _lune2 then
   _lune = _lune2
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

if not _lune2 then
   _lune2 = _lune
end


local function startsWith( txt, ptn )

   return txt:find( ptn, 1, true ) ~= nil
end
_moduleObj.startsWith = startsWith
local function endsWith( txt, ptn )

   return txt:find( ptn, #txt - #ptn + 1, true ) ~= nil
end
_moduleObj.endsWith = endsWith
local function getLineList( txt )

   local list = {}
   for line in string.gmatch( txt, "[^\n]*\n" ) do
      table.insert( list, line )
   end
   
   for last in string.gmatch( txt, "[^\n]+$" ) do
      table.insert( list, last )
   end
   
   return list
end
_moduleObj.getLineList = getLineList

return _moduleObj
