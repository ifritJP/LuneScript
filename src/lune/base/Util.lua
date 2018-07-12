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

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfomemStream = {}
_className2InfoMap.memStream = _classInfomemStream
local _classInfooutStream = {}
_className2InfoMap.outStream = _classInfooutStream
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 96, baseId = 1, txt = 'outStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {98}, }
_typeInfoList[2] = { parentId = 1, typeId = 100, baseId = 96, txt = 'memStream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {102, 104, 106}, }
_typeInfoList[3] = { parentId = 96, typeId = 98, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[4] = { parentId = 100, typeId = 102, baseId = 1, txt = '__init',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[5] = { parentId = 100, typeId = 104, baseId = 1, txt = 'write',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[6] = { parentId = 100, typeId = 106, baseId = 1, txt = 'get_txt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
----- meta -----
return moduleObj
