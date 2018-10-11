--lune/base/frontInterface.lns
local _moduleObj = {}
local __mod__ = 'lune.base.frontInterface'
if not _ENV._lune then
   _lune = {}
end
      
function _lune._fromMapSub( val, memKind )
   if type( memKind ) == "function" then
      return memKind( val )
   end
   if string.find( memKind, "!$" ) then
      if val == nil then
         return nil, true
      end
      memKind = memKind:sub( 1, #memKind - 1 )
   end
   local valType = type( val )
   if memKind == "stem" then
      if valType == "number" or valType == "string" or valType == "boolean" then
         return val
      end
      return nil
   end
   if memKind == "int" then
      if valType == "number" then
         return math.floor( val )
      end
      return nil
   end
   if memKind == "real" then
      if valType == "number" then
         return val
      end
      return nil
   end
   if memKind == "bool" then
      if valType == "boolean" then
         return val
      end
      return nil
   end
   if memKind == "str" then
      if valType == "string" then
         return val
      end
      return nil
   end
   if string.find( memKind, "^Array" ) or string.find( memKind, "^List" )
   then
      if valType == "table" then
         local tbl = {}
         for index, mem in ipairs( val ) do
            local kind = string.gsub( memKind, "^[%a]+<", "" )
            kind = string.gsub( kind, ">$", ""  )
            local memval, valid = _lune._fromMapSub( mem, kind )
            if memval == nil and not valid then
               return nil
            end
            tbl[ index ] = memval
         end
         return tbl
      end
      return nil
   end
   if string.find( memKind, "^Map" ) then
      if valType == "table" then
         local tbl = {}
         for key, mem in pairs( val ) do
            local kind = string.gsub( memKind, "^%a+<", "" )
            kind = string.gsub( kind, ">$", ""  )
            local delimitIndex = string.find( kind, ",", 1, true )
            local keyKind = string.sub( kind, 1, delimitIndex - 1 )
            local valKind = string.sub( kind, delimitIndex + 1 )
            local mapKey = _lune._fromMapSub( key, keyKind )
            local mapVal = _lune._fromMapSub( mem, valKind )
            if mapKey == nil or mapVal == nil then
               return nil
            end
            tbl[ mapKey ] = mapVal
         end
         return tbl
      end
      return nil
   end
end


function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, valid = _lune._fromMapSub( map[ memInfo.name ], memInfo.kind )
      if val == nil and not valid then
         return false
      end
      obj[ memInfo.name ] = val
   end
   return true
end

local frontInterface = {}
_moduleObj.frontInterface = frontInterface
function frontInterface.setmeta( obj )
  setmetatable( obj, { __index = frontInterface  } )
end
function frontInterface.new(  )
   local obj = {}
   frontInterface.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function frontInterface:__init(  ) 

end

local dummyFront = {}
function dummyFront:loadModule( mod )
   error( "not implements" )
end
function dummyFront:loadMeta( mod )
   error( "not implements" )
end
function dummyFront:searchModule( mod )
   error( "not implements" )
end
function dummyFront:error( message )
   error( "not implements" )
end
function dummyFront.setmeta( obj )
  setmetatable( obj, { __index = dummyFront  } )
end
function dummyFront.new(  )
   local obj = {}
   dummyFront.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function dummyFront:__init(  ) 

end

__luneScript = dummyFront.new()
local function setFront( newFront )
   __luneScript = newFront
end
_moduleObj.setFront = setFront
local function loadModule( mod )
   return __luneScript:loadModule( mod )
end
_moduleObj.loadModule = loadModule
local function loadMeta( mod )
   return __luneScript:loadMeta( mod )
end
_moduleObj.loadMeta = loadMeta
local function searchModule( mod )
   return __luneScript:searchModule( mod )
end
_moduleObj.searchModule = searchModule
local function error( message )
   __luneScript:error( message )
end
_moduleObj.error = error
return _moduleObj
