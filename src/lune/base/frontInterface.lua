--lune/base/frontInterface.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@frontInterface'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
end

function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end

function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune3 then
   _lune3 = _lune
end
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )


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

   local modTime = _lune.unwrapDefault( tonumber( (idStr:gsub( ":.*", "" ) ) ), 0.0)
   local buildCount = _lune.unwrapDefault( tonumber( (idStr:gsub( ".*:", "" ) ) ), 0.0)
   return ModuleId.new(modTime, math.floor(buildCount))
end


local ModuleProvideInfo = {}
_moduleObj.ModuleProvideInfo = ModuleProvideInfo
function ModuleProvideInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleProvideInfo  } )
end
function ModuleProvideInfo.new( typeInfo, symbolKind, mutable )
   local obj = {}
   ModuleProvideInfo.setmeta( obj )
   if obj.__init then
      obj:__init( typeInfo, symbolKind, mutable )
   end
   return obj
end
function ModuleProvideInfo:__init( typeInfo, symbolKind, mutable )

   self.typeInfo = typeInfo
   self.symbolKind = symbolKind
   self.mutable = mutable
end
function ModuleProvideInfo:get_typeInfo()
   return self.typeInfo
end
function ModuleProvideInfo:get_symbolKind()
   return self.symbolKind
end
function ModuleProvideInfo:get_mutable()
   return self.mutable
end


local ExportInfo = {}
_moduleObj.ExportInfo = ExportInfo
function ExportInfo.setmeta( obj )
  setmetatable( obj, { __index = ExportInfo  } )
end
function ExportInfo.new( moduleTypeInfo, provideInfo, processInfo, globalSymbolList )
   local obj = {}
   ExportInfo.setmeta( obj )
   if obj.__init then
      obj:__init( moduleTypeInfo, provideInfo, processInfo, globalSymbolList )
   end
   return obj
end
function ExportInfo:__init( moduleTypeInfo, provideInfo, processInfo, globalSymbolList )

   self.moduleTypeInfo = moduleTypeInfo
   self.provideInfo = provideInfo
   self.processInfo = processInfo
   self.globalSymbolList = globalSymbolList
end
function ExportInfo:get_moduleTypeInfo()
   return self.moduleTypeInfo
end
function ExportInfo:get_provideInfo()
   return self.provideInfo
end
function ExportInfo:get_processInfo()
   return self.processInfo
end
function ExportInfo:get_globalSymbolList()
   return self.globalSymbolList
end


local ModuleInfo = {}
setmetatable( ModuleInfo, { ifList = {Ast.ModuleInfoIF,} } )
_moduleObj.ModuleInfo = ModuleInfo
function ModuleInfo.new( fullName, assignName, idMap, moduleId, exportInfo, importedAliasMap )
   local obj = {}
   ModuleInfo.setmeta( obj )
   if obj.__init then obj:__init( fullName, assignName, idMap, moduleId, exportInfo, importedAliasMap ); end
   return obj
end
function ModuleInfo:__init(fullName, assignName, idMap, moduleId, exportInfo, importedAliasMap) 
   self.exportInfo = exportInfo
   self.moduleId = moduleId
   self.fullName = fullName
   self.assignName = assignName
   self.localTypeInfo2importIdMap = idMap
   self.importId2localTypeInfoMap = {}
   for typeInfo, importId in pairs( idMap ) do
      self.importId2localTypeInfoMap[importId] = typeInfo
   end
   
   self.importedAliasMap = importedAliasMap
end
function ModuleInfo:getImportTypeId( typeInfo )

   do
      local typeId = self.localTypeInfo2importIdMap[typeInfo]
      if typeId ~= nil then
         return typeId
      end
   end
   
   return nil
end
function ModuleInfo:getTypeInfo( localTypeId )

   do
      local typeInfo = self.importId2localTypeInfoMap[localTypeId]
      if typeInfo ~= nil then
         return typeInfo
      end
   end
   
   return nil
end
function ModuleInfo:get_modulePath(  )

   return self.fullName
end
function ModuleInfo:assign( assignName )

   return ModuleInfo.new(self.fullName, assignName, self.localTypeInfo2importIdMap, self.moduleId, self.exportInfo, self.importedAliasMap)
end
function ModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfo  } )
end
function ModuleInfo:get_fullName()
   return self.fullName
end
function ModuleInfo:get_localTypeInfo2importIdMap()
   return self.localTypeInfo2importIdMap
end
function ModuleInfo:get_importId2localTypeInfoMap()
   return self.importId2localTypeInfoMap
end
function ModuleInfo:get_assignName()
   return self.assignName
end
function ModuleInfo:get_moduleId()
   return self.moduleId
end
function ModuleInfo:get_importedAliasMap()
   return self.importedAliasMap
end
function ModuleInfo:get_exportInfo()
   return self.exportInfo
end


local ModuleMeta = {}
_moduleObj.ModuleMeta = ModuleMeta
function ModuleMeta.setmeta( obj )
  setmetatable( obj, { __index = ModuleMeta  } )
end
function ModuleMeta.new( metaInfo, lnsPath, moduleInfo )
   local obj = {}
   ModuleMeta.setmeta( obj )
   if obj.__init then
      obj:__init( metaInfo, lnsPath, moduleInfo )
   end
   return obj
end
function ModuleMeta:__init( metaInfo, lnsPath, moduleInfo )

   self.metaInfo = metaInfo
   self.lnsPath = lnsPath
   self.moduleInfo = moduleInfo
end
function ModuleMeta:get_metaInfo()
   return self.metaInfo
end
function ModuleMeta:get_lnsPath()
   return self.lnsPath
end
function ModuleMeta:get_moduleInfo()
   return self.moduleInfo
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
   self.orderedSet = Util.OrderedSet.new()
end
function ImportModuleInfo:add( modulePath )

   return self.orderedSet:add( modulePath )
end
function ImportModuleInfo:remove(  )

   self.orderedSet:removeLast(  )
end
function ImportModuleInfo:getFull(  )

   local txt = ""
   for __index, modulePath in ipairs( self.orderedSet:get_list() ) do
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
setmetatable( dummyFront, { ifList = {frontInterface,} } )
function dummyFront:loadModule( mod )

   local loaded = _lune.loadstring52( "return {}" )
   local emptyTable
   
   if loaded ~= nil then
      emptyTable = _lune.unwrap( loaded(  ))
   else
      error( "load error" )
   end
   
   local meta = ModuleMeta.new(emptyTable, mod:gsub( "%.", "/" ) .. ".lns", nil)
   return require( mod ), meta
end
function dummyFront:loadMeta( importModuleInfo, mod )

   error( "not implements" )
end
function dummyFront:loadFromLnsTxt( importModuleInfo, name, txt )

   error( "not implements" )
end
function dummyFront:getLuaModulePath( mod )

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

local function getLuaModulePath( mod )

   return __luneScript:getLuaModulePath( mod )
end
_moduleObj.getLuaModulePath = getLuaModulePath

return _moduleObj
