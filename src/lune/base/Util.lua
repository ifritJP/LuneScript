--lune/base/Util.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Util'
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

local Depend = require( 'lune.base.Depend' )
local memStream = {}
_moduleObj.memStream = memStream
function memStream.new(  )
   local obj = {}
   memStream.setmeta( obj )
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
function memStream:flush(  )
end
function memStream.setmeta( obj )
  setmetatable( obj, { __index = memStream  } )
end
function memStream:get_txt()       
   return self.txt         
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
local function existFile( path )
   local fileObj = io.open( path )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return false
   end
   
   fileObj:close(  )
   return true
end
_moduleObj.existFile = existFile
return _moduleObj
