--lune/base/Util.lns
local _moduleObj = {}
local Depend = require( 'lune.base.Depend' )

local memStream = {}
_moduleObj.memStream = memStream
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
function memStream:close(  )
end
function memStream:get_txt()       
  return self.txt         
end
do
  end

local debugFlag = true

local function setDebugFlag( flag )
  debugFlag = flag
end
_moduleObj.setDebugFlag = setDebugFlag
local errorCode = 1

local function setErrorCode( code )
  errorCode = code
end
_moduleObj.setErrorCode = setErrorCode
local function errorLog( message )
  io.stderr:write( message .. "\n" )
end
_moduleObj.errorLog = errorLog
local function err( message )
  if debugFlag then
    error( message )
  end
  errorLog( message )
  os.exit( errorCode )
end
_moduleObj.err = err
local function log( message )
  if debugFlag then
    errorLog( message )
  end
end
_moduleObj.log = log
local function printStackTrace(  )
  for level = 2, 6 do
    local debugInfo = debug.getinfo( level )
    
    if debugInfo then
      errorLog( string.format( "-- %s %s", debugInfo["short_src"], debugInfo['currentline']) )
    end
  end
end
_moduleObj.printStackTrace = printStackTrace
local function profile( validTest, func, path )
  if not validTest then
    return func(  )
  end
  local ProFi = require( 'ProFi' )
  _moduleObj.ProFi = ProFi
  
  ProFi.start(  )
  local result = func(  )
  
  ProFi.stop(  )
  ProFi.writeReport( path )
  return result
end
_moduleObj.profile = profile
local function getReadyCode( lnsPath, luaPath )
  local luaTime, lnsTime = Depend.getFileLastModifiedTime( luaPath ), Depend.getFileLastModifiedTime( lnsPath )
  
      if  nil == luaTime or  nil == lnsTime then
        local _luaTime = luaTime
        local _lnsTime = lnsTime
        
        return false
      end
    
  return luaTime >= lnsTime
end
_moduleObj.getReadyCode = getReadyCode
return _moduleObj
