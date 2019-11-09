--lune/base/Meta.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Meta'
local _lune = {}
if _lune1 then
   _lune = _lune1
end
if not _lune1 then
   _lune1 = _lune
end



local _MetaInfo = {}
_moduleObj._MetaInfo = _MetaInfo
function _MetaInfo.setmeta( obj )
  setmetatable( obj, { __index = _MetaInfo  } )
end
function _MetaInfo.new( __formatVersion, __enableTest, __buildId, __typeId2ClassInfoMap, __typeInfoList, __varName2InfoMap, __funcName2InfoMap, __moduleTypeId, __moduleSymbolKind, __moduleMutable, __dependModuleMap, __dependIdMap, __macroName2InfoMap )
   local obj = {}
   _MetaInfo.setmeta( obj )
   if obj.__init then
      obj:__init( __formatVersion, __enableTest, __buildId, __typeId2ClassInfoMap, __typeInfoList, __varName2InfoMap, __funcName2InfoMap, __moduleTypeId, __moduleSymbolKind, __moduleMutable, __dependModuleMap, __dependIdMap, __macroName2InfoMap )
   end
   return obj
end
function _MetaInfo:__init( __formatVersion, __enableTest, __buildId, __typeId2ClassInfoMap, __typeInfoList, __varName2InfoMap, __funcName2InfoMap, __moduleTypeId, __moduleSymbolKind, __moduleMutable, __dependModuleMap, __dependIdMap, __macroName2InfoMap )

   self.__formatVersion = __formatVersion
   self.__enableTest = __enableTest
   self.__buildId = __buildId
   self.__typeId2ClassInfoMap = __typeId2ClassInfoMap
   self.__typeInfoList = __typeInfoList
   self.__varName2InfoMap = __varName2InfoMap
   self.__funcName2InfoMap = __funcName2InfoMap
   self.__moduleTypeId = __moduleTypeId
   self.__moduleSymbolKind = __moduleSymbolKind
   self.__moduleMutable = __moduleMutable
   self.__dependModuleMap = __dependModuleMap
   self.__dependIdMap = __dependIdMap
   self.__macroName2InfoMap = __macroName2InfoMap
end


return _moduleObj
