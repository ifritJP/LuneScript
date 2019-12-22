--lune/base/frontInterface.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@frontInterface'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
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

if not _lune2 then
   _lune2 = _lune
end

local Util = _lune.loadModule( 'lune.base.Util' )



local ModuleId = {}
_moduleObj.ModuleId = ModuleId
function ModuleId.new( modTime, buildCount )
   local obj = {}
   ModuleId.setmeta( obj )
   if obj.__init then obj:__init( modTime, buildCount ); end
   return obj
end
function ModuleId:__init(modTime, buildCount) 
   self.modTime = modTime
   self.buildCount = buildCount
   self.idStr = string.format( "%f:%d", modTime, buildCount)
end
function ModuleId:getNextModuleId(  )

   return ModuleId.new(self.modTime, self.buildCount + 1)
end
function ModuleId.setmeta( obj )
  setmetatable( obj, { __index = ModuleId  } )
end
function ModuleId:get_modTime()
   return self.modTime
end
function ModuleId:get_buildCount()
   return self.buildCount
end
function ModuleId:get_idStr()
   return self.idStr
end
do
   ModuleId.tempId = ModuleId.new(0.0, 0)
end

function ModuleId.createId( modTime, buildCount )

   return ModuleId.new(modTime, buildCount)
end

function ModuleId.createIdFromTxt( idStr )

   local modTime = _lune.unwrapDefault( tonumber( (idStr:gsub( ":.*", "" ) ) ), 0.0)
   local buildCount = _lune.unwrapDefault( tonumber( (idStr:gsub( ".*:", "" ) ) ), 0.0)
   return ModuleId.new(modTime, math.floor(buildCount))
end


local ImportModuleInfo = {}
_moduleObj.ImportModuleInfo = ImportModuleInfo
function ImportModuleInfo.new(  )
   local obj = {}
   ImportModuleInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function ImportModuleInfo:__init() 
   self.orderedSet = Util.OrderedSet.new()
end
function ImportModuleInfo:add( modulePath )

   return self.orderedSet:add( modulePath )
end
function ImportModuleInfo:remove(  )

   self.orderedSet:removeLast(  )
end
function ImportModuleInfo:getFull(  )

   local txt = ""
   for __index, modulePath in pairs( self.orderedSet:get_list() ) do
      txt = string.format( "%s -> %s", txt, modulePath)
   end
   
   return txt
end
function ImportModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = ImportModuleInfo  } )
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
setmetatable( dummyFront, { ifList = {frontInterface,} } )
function dummyFront:loadModule( mod )

   
   return require( mod ), {}
end
function dummyFront:loadMeta( importModuleInfo, mod )

   error( "not implements" )
end
function dummyFront:loadFromLnsTxt( importModuleInfo, name, txt )

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

local function loadFromLnsTxt( importModuleInfo, name, txt )

   return __luneScript:loadFromLnsTxt( importModuleInfo, name, txt )
end
_moduleObj.loadFromLnsTxt = loadFromLnsTxt

local function loadMeta( importModuleInfo, mod )

   return __luneScript:loadMeta( importModuleInfo, mod )
end
_moduleObj.loadMeta = loadMeta

local function searchModule( mod )

   return __luneScript:searchModule( mod )
end
_moduleObj.searchModule = searchModule

return _moduleObj
