--lune/base/frontInterface.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@frontInterface'
local _lune = {}
if _lune8 then
   _lune = _lune8
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

function _lune.loadstring51( txt, env )
   local func = loadstring( txt )
   if func and env then
      setfenv( func, env )
   end
   return func
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
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Runner = _lune.loadModule( 'lune.base.Runner' )




local ModuleId = {}
_moduleObj.ModuleId = ModuleId
function ModuleId._new( modTime, buildCount )
   local obj = {}
   ModuleId._setmeta( obj )
   if obj.__init then obj:__init( modTime, buildCount ); end
   return obj
end
function ModuleId:__init(modTime, buildCount) 
   self.modTime = modTime
   self.buildCount = buildCount
   self.idStr = string.format( "%f:%d", modTime, buildCount)
end
function ModuleId:getNextModuleId(  )

   return ModuleId._new(self.modTime, self.buildCount + 1)
end
function ModuleId._setmeta( obj )
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
   ModuleId.tempId = ModuleId._new(0.0, 0)
end

function ModuleId.createId( modTime, buildCount )

   return ModuleId._new(modTime, buildCount)
end

function ModuleId.createIdFromTxt( idStr )

   local modTime = _lune.unwrapDefault( tonumber( (idStr:gsub( ":.*", "" ) ) ), 0.0)
   local buildCount = _lune.unwrapDefault( tonumber( (idStr:gsub( ".*:", "" ) ) ), 0.0)
   return ModuleId._new(modTime, math.floor(buildCount))
end


local ModuleProvideInfo = {}
_moduleObj.ModuleProvideInfo = ModuleProvideInfo
function ModuleProvideInfo._setmeta( obj )
  setmetatable( obj, { __index = ModuleProvideInfo  } )
end
function ModuleProvideInfo._new( typeInfo, symbolKind, mutable )
   local obj = {}
   ModuleProvideInfo._setmeta( obj )
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


local LuneHelperInfo = {}
_moduleObj.LuneHelperInfo = LuneHelperInfo
function LuneHelperInfo._new(  )
   local obj = {}
   LuneHelperInfo._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function LuneHelperInfo:__init() 
   self.useStrReplace = false
   self.useNilAccess = false
   self.useUnwrapExp = false
   self.hasMappingClassDef = false
   self.useLoad = false
   self.useUnpack = false
   self.useAlge = false
   self.useSet = false
   self.callAnonymous = false
   self.pragmaSet = {}
   self.useLazyLoad = false
   self.useLazyRequire = false
   self.useRun = false
   self.useResult = false
   self.useError = false
end
function LuneHelperInfo:mergeFrom( src )

   
   
   self.useNilAccess = self.useNilAccess or src.useNilAccess
   
   self.useUnwrapExp = self.useUnwrapExp or src.useUnwrapExp
   
   self.hasMappingClassDef = self.hasMappingClassDef or src.hasMappingClassDef
   
   self.useLoad = self.useLoad or src.useLoad
   
   self.useUnpack = self.useUnpack or src.useUnpack
   
   self.useAlge = self.useAlge or src.useAlge
   
   self.useSet = self.useSet or src.useSet
   
   self.callAnonymous = self.callAnonymous or src.callAnonymous
   
   self.useLazyLoad = self.useLazyLoad or src.useLazyLoad
   
   self.useLazyRequire = self.useLazyRequire or src.useLazyRequire
   
   self.useRun = self.useRun or src.useRun
   
   self.useStrReplace = self.useStrReplace or src.useStrReplace
   
   self.useResult = self.useResult or src.useResult
   
   self.useError = self.useError or src.useError
   
   for val, __val in pairs( src.pragmaSet ) do
      self.pragmaSet[val]= true
   end
   
end
function LuneHelperInfo._setmeta( obj )
  setmetatable( obj, { __index = LuneHelperInfo  } )
end


local function getRootDependModId(  )

   return -1
end
_moduleObj.getRootDependModId = getRootDependModId

local ExportInfo = {}
setmetatable( ExportInfo, { ifList = {Ast.ModuleInfoIF,} } )
_moduleObj.ExportInfo = ExportInfo
function ExportInfo._new( moduleTypeInfo, provideInfo, processInfo, globalSymbolList, importedAliasMap, moduleId, fullName, assignName, streamName, idMap )
   local obj = {}
   ExportInfo._setmeta( obj )
   if obj.__init then obj:__init( moduleTypeInfo, provideInfo, processInfo, globalSymbolList, importedAliasMap, moduleId, fullName, assignName, streamName, idMap ); end
   return obj
end
function ExportInfo:__init(moduleTypeInfo, provideInfo, processInfo, globalSymbolList, importedAliasMap, moduleId, fullName, assignName, streamName, idMap) 
   self.moduleTypeInfo = moduleTypeInfo
   self.provideInfo = provideInfo
   self.processInfo = processInfo
   self.globalSymbolList = globalSymbolList
   self.importedAliasMap = importedAliasMap
   self.moduleId = moduleId
   self.fullName = fullName
   self.assignName = assignName
   self.streamName = streamName
   
   local importId2localTypeInfoMap = {}
   for typeInfo, importId in pairs( idMap ) do
      importId2localTypeInfoMap[importId] = typeInfo
   end
   
   self.importId2localTypeInfoMap = importId2localTypeInfoMap
end
function ExportInfo:get_modulePath(  )

   return self.fullName
end
function ExportInfo:getTypeInfo( localTypeId )

   do
      local typeInfo = self.importId2localTypeInfoMap[localTypeId]
      if typeInfo ~= nil then
         return typeInfo
      end
   end
   
   return nil
end
function ExportInfo._setmeta( obj )
  setmetatable( obj, { __index = ExportInfo  } )
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
function ExportInfo:get_importedAliasMap()
   return self.importedAliasMap
end
function ExportInfo:get_moduleId()
   return self.moduleId
end
function ExportInfo:get_fullName()
   return self.fullName
end
function ExportInfo:get_assignName()
   return self.assignName
end
function ExportInfo:get_streamName()
   return self.streamName
end
function ExportInfo:get_importId2localTypeInfoMap()
   return self.importId2localTypeInfoMap
end
function ExportInfo:set_importId2localTypeInfoMap( importId2localTypeInfoMap )
   self.importId2localTypeInfoMap = importId2localTypeInfoMap
end


function ExportInfo:assign( assignName )

   local info = ExportInfo._new(self.moduleTypeInfo, self.provideInfo, self.processInfo, self.globalSymbolList, self.importedAliasMap, self.moduleId, self.fullName, assignName, self.streamName, {})
   info.importId2localTypeInfoMap = self.importId2localTypeInfoMap
   return info
end

local ModuleInfo = {}
_moduleObj.ModuleInfo = ModuleInfo
function ModuleInfo._new( exportInfo )
   local obj = {}
   ModuleInfo._setmeta( obj )
   if obj.__init then obj:__init( exportInfo ); end
   return obj
end
function ModuleInfo:__init(exportInfo) 
   self.exportInfo = exportInfo
end
function ModuleInfo._setmeta( obj )
  setmetatable( obj, { __index = ModuleInfo  } )
end
function ModuleInfo:get_exportInfo()
   return self.exportInfo
end
function ModuleInfo:assign( ... )
   return self.exportInfo:assign( ... )
end

function ModuleInfo:getTypeInfo( ... )
   return self.exportInfo:getTypeInfo( ... )
end

function ModuleInfo:get_assignName( ... )
   return self.exportInfo:get_assignName( ... )
end

function ModuleInfo:get_fullName( ... )
   return self.exportInfo:get_fullName( ... )
end

function ModuleInfo:get_globalSymbolList( ... )
   return self.exportInfo:get_globalSymbolList( ... )
end

function ModuleInfo:get_importId2localTypeInfoMap( ... )
   return self.exportInfo:get_importId2localTypeInfoMap( ... )
end

function ModuleInfo:get_importedAliasMap( ... )
   return self.exportInfo:get_importedAliasMap( ... )
end

function ModuleInfo:get_moduleId( ... )
   return self.exportInfo:get_moduleId( ... )
end

function ModuleInfo:get_modulePath( ... )
   return self.exportInfo:get_modulePath( ... )
end

function ModuleInfo:get_moduleTypeInfo( ... )
   return self.exportInfo:get_moduleTypeInfo( ... )
end

function ModuleInfo:get_processInfo( ... )
   return self.exportInfo:get_processInfo( ... )
end

function ModuleInfo:get_provideInfo( ... )
   return self.exportInfo:get_provideInfo( ... )
end

function ModuleInfo:get_streamName( ... )
   return self.exportInfo:get_streamName( ... )
end

function ModuleInfo:set_importId2localTypeInfoMap( ... )
   return self.exportInfo:set_importId2localTypeInfoMap( ... )
end



local MetaOrModule = {}
MetaOrModule._name2Val = {}
_moduleObj.MetaOrModule = MetaOrModule
function MetaOrModule:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "MetaOrModule.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function MetaOrModule._from( val )
   return _lune._AlgeFrom( MetaOrModule, val )
end

MetaOrModule.Export = { "Export", {{}}}
MetaOrModule._name2Val["Export"] = MetaOrModule.Export
MetaOrModule.MetaRaw = { "MetaRaw", {{}}}
MetaOrModule._name2Val["MetaRaw"] = MetaOrModule.MetaRaw
MetaOrModule.Module = { "Module", {{},{}}}
MetaOrModule._name2Val["Module"] = MetaOrModule.Module


local ModuleMeta = {}
_moduleObj.ModuleMeta = ModuleMeta
function ModuleMeta._setmeta( obj )
  setmetatable( obj, { __index = ModuleMeta  } )
end
function ModuleMeta._new( lnsPath, metaOrModule )
   local obj = {}
   ModuleMeta._setmeta( obj )
   if obj.__init then
      obj:__init( lnsPath, metaOrModule )
   end
   return obj
end
function ModuleMeta:__init( lnsPath, metaOrModule )

   self.lnsPath = lnsPath
   self.metaOrModule = metaOrModule
end
function ModuleMeta:get_lnsPath()
   return self.lnsPath
end
function ModuleMeta:get_metaOrModule()
   return self.metaOrModule
end
function ModuleMeta:set_metaOrModule( metaOrModule )
   self.metaOrModule = metaOrModule
end

local ImportModuleInfo = {}
_moduleObj.ImportModuleInfo = ImportModuleInfo
function ImportModuleInfo._new(  )
   local obj = {}
   ImportModuleInfo._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function ImportModuleInfo:__init() 
   self.orderedSet = Util.OrderedSet._new()
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
function ImportModuleInfo:clone(  )

   local info = ImportModuleInfo._new()
   for __index, mod in ipairs( self.orderedSet:get_list() ) do
      info:add( mod )
   end
   
   return info
end
function ImportModuleInfo:len(  )

   return #self.orderedSet:get_list()
end
function ImportModuleInfo:list(  )

   return self.orderedSet:get_list()
end
function ImportModuleInfo._setmeta( obj )
  setmetatable( obj, { __index = ImportModuleInfo  } )
end


local ModuleLoader = {}
_moduleObj.ModuleLoader = ModuleLoader
function ModuleLoader._setmeta( obj )
  setmetatable( obj, { __index = ModuleLoader  } )
end
function ModuleLoader._new(  )
   local obj = {}
   ModuleLoader._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function ModuleLoader:__init(  )

end


local frontInterface = {}
_moduleObj.frontInterface = frontInterface
function frontInterface._setmeta( obj )
  setmetatable( obj, { __index = frontInterface  } )
end
function frontInterface._new(  )
   local obj = {}
   frontInterface._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function frontInterface:__init(  )

end


local function dummyLoadModule( mod )

   
   local modVal, moduleMeta
   
   do
      local emptyTable
      
      local loaded = _lune.loadstring51( "return {}" )
      if loaded ~= nil then
         emptyTable = _lune.unwrap( loaded(  ))
      else
         Util.err( "load error" )
      end
      
      moduleMeta = ModuleMeta._new(mod:gsub( "%.", "/" ) .. ".lns", _lune.newAlge( MetaOrModule.MetaRaw, {emptyTable}))
      modVal = require( mod )
   end
   
   return modVal, moduleMeta
end
_moduleObj.dummyLoadModule = dummyLoadModule

local dummyFront = {}
setmetatable( dummyFront, { ifList = {frontInterface,} } )
function dummyFront:loadModule( mod )

   return dummyLoadModule( mod )
end
function dummyFront:loadMeta( importModuleInfo, mod, orgMod, baseDir, loader )
   local __func__ = '@lune.@base.@frontInterface.dummyFront.loadMeta'

   Util.err( string.format( "not implements: %s", __func__) )
end
function dummyFront:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )
   local __func__ = '@lune.@base.@frontInterface.dummyFront.loadFromLnsTxt'

   Util.err( string.format( "not implements: %s", __func__) )
end
function dummyFront:getLuaModulePath( mod, baseDir )

   return mod, nil, mod
end
function dummyFront:searchModule( mod, baseDir, addSearchPath )
   local __func__ = '@lune.@base.@frontInterface.dummyFront.searchModule'

   Util.err( string.format( "not implements: %s", __func__) )
end
function dummyFront:error( message )

   Util.err( message )
end
function dummyFront._setmeta( obj )
  setmetatable( obj, { __index = dummyFront  } )
end
function dummyFront._new(  )
   local obj = {}
   dummyFront._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function dummyFront:__init(  )

end


__luneScript = dummyFront._new()
_moduleObj.__luneScript = __luneScript


local function setFront( newFront )

   __luneScript = newFront
end
_moduleObj.setFront = setFront

local function loadModule( mod )

   return __luneScript:loadModule( mod )
end
_moduleObj.loadModule = loadModule

local function loadFromLnsTxt( importModuleInfo, baseDir, name, txt )

   return __luneScript:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )
end
_moduleObj.loadFromLnsTxt = loadFromLnsTxt

local function loadMeta( importModuleInfo, mod, orgMod, baseDir, loader )

   return __luneScript:loadMeta( importModuleInfo, mod, orgMod, baseDir, loader )
end
_moduleObj.loadMeta = loadMeta

local function searchModule( mod, baseDir, addSearchPath )

   return __luneScript:searchModule( mod, baseDir, addSearchPath )
end
_moduleObj.searchModule = searchModule

local function getLuaModulePath( mod, baseDir )

   return __luneScript:getLuaModulePath( mod, baseDir )
end
_moduleObj.getLuaModulePath = getLuaModulePath

return _moduleObj
