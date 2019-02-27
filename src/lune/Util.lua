--lune/Util.lns
local _moduleObj = {}
local __mod__ = 'lune.Util'
if not _lune then
   _lune = {}
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

local function SetOr( val1, val2 )

   local set = {}
   for val, __val in pairs( val1 ) do
      set[val]= true
   end
   
   for val, __val in pairs( val2 ) do
      set[val]= true
   end
   
   return set
end
_moduleObj.SetOr = SetOr
local function SetAnd( val1, val2 )

   local set = {}
   for val, __val in pairs( val1 ) do
      if val2[val] then
         set[val]= true
      end
      
   end
   
   return set
end
_moduleObj.SetAnd = SetAnd
local function SetSub( val1, val2 )

   local set = {}
   for val, __val in pairs( val1 ) do
      if not val2[val] then
         set[val]= true
      end
      
   end
   
   return set
end
_moduleObj.SetSub = SetSub
local function SetLen( set )

   local count = 0
   for val, __val in pairs( set ) do
      count = count + 1
   end
   
   return count
end
_moduleObj.SetLen = SetLen
return _moduleObj
