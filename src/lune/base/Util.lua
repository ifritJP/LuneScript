--lune/base/Util.lns
local moduleObj = {}
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

local function errorLog( message )
  local write = (io ).stderr.write
  write( (io ).stderr, message .. "\n" )
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
  local profiler = require( 'ProFi' )
  profiler.start(  )
  local result = func(  )
  profiler.stop(  )
  profiler.writeReport( path )
  return result
end
moduleObj.profile = profile
----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfomemStream = {}
  _className2InfoMap.memStream = _classInfomemStream
  end
do
  local _classInfooutStream = {}
  _className2InfoMap.outStream = _classInfooutStream
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 100, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {102}, }
_typeInfoList[2] = { parentId = 100, typeId = 102, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[3] = { parentId = 102, typeId = 104, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {106, 110, 118, 120, 122}, }
_typeInfoList[4] = { parentId = 104, typeId = 106, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {108}, }
_typeInfoList[5] = { parentId = 106, typeId = 108, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[6] = { parentId = 104, typeId = 110, baseId = 106, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {112, 114, 116}, }
_typeInfoList[7] = { parentId = 110, typeId = 112, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[8] = { parentId = 110, typeId = 114, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[9] = { parentId = 110, typeId = 116, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[10] = { parentId = 104, typeId = 118, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[11] = { parentId = 104, typeId = 120, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[12] = { parentId = 104, typeId = 122, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
