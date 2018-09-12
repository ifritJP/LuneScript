--lune/base/frontInterface.lns
local _moduleObj = {}
local frontInterface = {}
_moduleObj.frontInterface = frontInterface
-- none
-- none
-- none
-- none
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
do
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
do
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
