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
  (profiler.writeReport )( path )
  return result
end
moduleObj.profile = profile
----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo114 = {}
  _className2InfoMap.memStream = _classInfo114
  _typeId2ClassInfoMap[ 114 ] = _classInfo114
  end
do
  local _classInfo108 = {}
  _className2InfoMap.outStream = _classInfo108
  _typeId2ClassInfoMap[ 108 ] = _classInfo108
  end
do
  local _classInfo106 = {}
  _className2InfoMap.Util = _classInfo106
  _typeId2ClassInfoMap[ 106 ] = _classInfo106
  _classInfo106.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 124 }
  _classInfo106.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 122 }
  _classInfo106.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 114 }
  _classInfo106.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 108 }
  _classInfo106.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 128 }
  end
do
  local _classInfo104 = {}
  _className2InfoMap.base = _classInfo104
  _typeId2ClassInfoMap[ 104 ] = _classInfo104
  _classInfo104.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 106 }
  end
do
  local _classInfo102 = {}
  _className2InfoMap.lune = _classInfo102
  _typeId2ClassInfoMap[ 102 ] = _classInfo102
  _classInfo102.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 104 }
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 9, nilable = true, orgTypeId = 8 }
_typeInfoList[2] = { parentId = 1, typeId = 11, nilable = true, orgTypeId = 10 }
_typeInfoList[3] = { parentId = 1, typeId = 13, nilable = true, orgTypeId = 12 }
_typeInfoList[4] = { parentId = 1, typeId = 19, nilable = true, orgTypeId = 18 }
_typeInfoList[5] = { parentId = 1, typeId = 102, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[6] = { parentId = 1, typeId = 103, nilable = true, orgTypeId = 102 }
_typeInfoList[7] = { parentId = 102, typeId = 104, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106}, }
_typeInfoList[8] = { parentId = 102, typeId = 105, nilable = true, orgTypeId = 104 }
_typeInfoList[9] = { parentId = 104, typeId = 106, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {108, 114, 122, 124, 128}, }
_typeInfoList[10] = { parentId = 104, typeId = 107, nilable = true, orgTypeId = 106 }
_typeInfoList[11] = { parentId = 106, typeId = 108, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {110, 112}, }
_typeInfoList[12] = { parentId = 106, typeId = 109, nilable = true, orgTypeId = 108 }
_typeInfoList[13] = { parentId = 108, typeId = 110, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 108, typeId = 112, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 106, typeId = 114, baseId = 108, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {116, 118, 120}, }
_typeInfoList[16] = { parentId = 106, typeId = 115, nilable = true, orgTypeId = 114 }
_typeInfoList[17] = { parentId = 114, typeId = 116, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[18] = { parentId = 114, typeId = 118, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[19] = { parentId = 114, typeId = 120, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[20] = { parentId = 106, typeId = 122, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[21] = { parentId = 106, typeId = 124, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[22] = { parentId = 106, typeId = 128, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10, 6, 18}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
