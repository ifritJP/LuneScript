--lune/base/Util.lns
local moduleObj = {}
local Depend = require( 'lune.base.Depend' )

local outStream = {}
moduleObj.outStream = outStream
-- none
function outStream.new(  )
  local obj = {}
  setmetatable( obj, { __index = outStream } )
  if obj.__init then
    obj:__init(  )
  end        
  return obj 
 end         
function outStream:__init(  ) 
            
end
do
  end

local memStream = {}
setmetatable( memStream, { __index = outStream } )
moduleObj.memStream = memStream
function memStream.new(  )
  local obj = {}
  setmetatable( obj, { __index = memStream } )
  if obj.__init then obj:__init(  ); end
return obj
end
function memStream:__init() 
  self.txt = ""
end
function memStream:write( val )
  self.txt = self.txt .. val
end
function memStream:get_txt()
  return self.txt
end
do
  end

local function errorLog( message )
  local stderr = (io ).stderr
  
  local write = (stderr ).write
  
  write( stderr, message .. "\n" )
end
moduleObj.errorLog = errorLog
local function debugLog(  )
  for level = 2, 6 do
    local debugInfo = debug.getinfo( level )
    
    if debugInfo then
      errorLog( string.format( "-- %s %s", debugInfo["short_src"], debugInfo['currentline']) )
    end
  end
end
moduleObj.debugLog = debugLog
local function profile( validTest, func, path )
  if not validTest then
    return func(  )
  end
  local profiler = require( 'ProFi' ) or _luneScript.error( 'unwrap val is nil' )
  
  ((profiler.start or _luneScript.error( 'unwrap val is nil' ) ) )(  )
  local result = func(  )
  
  ((profiler.stop or _luneScript.error( 'unwrap val is nil' ) ) )(  )
  ((profiler.writeReport or _luneScript.error( 'unwrap val is nil' ) ) )( path )
  return result
end
moduleObj.profile = profile
local function getReadyCode( lnsPath, luaPath )
  local luaTime, lnsTime = Depend.getFileLastModifiedTime( luaPath ), Depend.getFileLastModifiedTime( lnsPath )
  
      if  not luaTime or  not lnsTime then
        local _luaTime = luaTime
        local _lnsTime = lnsTime
        
        return false
      end
    
  return luaTime >= lnsTime
end
moduleObj.getReadyCode = getReadyCode
return moduleObj
