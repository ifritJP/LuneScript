--lune/base/Util.lns
local moduleObj = {}
local function _lune_nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   end
   error( string.format( "illegal access -- %s", access ) )
end

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
  local stderr = (io ).stderr or _luneScript.error( 'unwrap val is nil' )
  
  local write = (stderr.write or _luneScript.error( 'unwrap val is nil' ) )
  
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
  local ProFi = require( 'ProFi' )
  
  ProFi.start(  )
  local result = func(  )
  
  ProFi.stop(  )
  ProFi.writeReport( path )
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
