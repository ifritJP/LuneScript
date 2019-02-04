--lune/base/frontInterface.lns
local _moduleObj = {}
local __mod__ = 'lune.base.frontInterface'
if not _lune then
   _lune = {}
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
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

   local modTime = tonumber( (idStr:gsub( ":.*", "" ) ) )
   local buildCount = tonumber( (idStr:gsub( ".*:", "" ) ) )
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
   self.moduleSet = {}
   self.moduleList = {}
end
function ImportModuleInfo:add( modulePath )

   if self.moduleSet[modulePath] then
      return false
   end
   
   self.moduleSet[modulePath] = true
   table.insert( self.moduleList, modulePath )
   return true
end
function ImportModuleInfo:remove(  )

   if #self.moduleList == 0 then
      Util.err( "self.moduleList is 0" )
   end
   
   self.moduleSet[self.moduleList[#self.moduleList]] = nil
   table.remove( self.moduleList )
end
function ImportModuleInfo:getFull(  )

   local txt = ""
   for __index, modulePath in pairs( self.moduleList ) do
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
local function error( message )

   __luneScript:error( message )
end
_moduleObj.error = error
return _moduleObj
