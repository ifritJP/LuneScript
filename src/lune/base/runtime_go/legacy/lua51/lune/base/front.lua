--lune/base/front.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@front'
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

function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
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
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
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
local Types = _lune.loadModule( 'lune.base.Types' )
local Str = _lune.loadModule( 'lune.base.Str' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local convLua = _lune.loadModule( 'lune.base.convLua' )
local convGo = _lune.loadModule( 'lune.base.convGo' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Option = _lune.loadModule( 'lune.base.Option' )
local dumpNode = _lune.loadModule( 'lune.base.dumpNode' )
local glueFilter = _lune.loadModule( 'lune.base.glueFilter' )
local Depend = _lune.loadModule( 'lune.base.Depend' )
local OutputDepend = _lune.loadModule( 'lune.base.OutputDepend' )
local Ver = _lune.loadModule( 'lune.base.Ver' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Formatter = _lune.loadModule( 'lune.base.Formatter' )
local Testing = _lune.loadModule( 'lune.base.Testing' )
local GoMod = _lune.loadModule( 'lune.base.GoMod' )
local Meta = _lune.loadModule( 'lune.base.Meta' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )


Depend.setup( function ( ver )

   LuaVer.setCurVer( ver )
end )

local forceUpdateMeta = true

local LoadInfo = {}
function LoadInfo.setmeta( obj )
  setmetatable( obj, { __index = LoadInfo  } )
end
function LoadInfo.new( mod, meta )
   local obj = {}
   LoadInfo.setmeta( obj )
   if obj.__init then
      obj:__init( mod, meta )
   end
   return obj
end
function LoadInfo:__init( mod, meta )

   self.mod = mod
   self.meta = meta
end


local Front = {}
setmetatable( Front, { ifList = {frontInterface.frontInterface,} } )
function Front.new( option, bindModuleList )
   local obj = {}
   Front.setmeta( obj )
   if obj.__init then obj:__init( option, bindModuleList ); end
   return obj
end
function Front:__init(option, bindModuleList) 
   self.bindModuleSet = {}
   if bindModuleList ~= nil then
      for __index, mod in ipairs( bindModuleList ) do
         self.bindModuleSet[mod]= true
      end
      
   end
   
   self.mod2ast = Util.OrderdMap.new()
   self.gomodMap = GoMod.getGoMap(  )
   self.option = option
   self.loadedMap = {}
   self.loadedMapTest = {}
   self.loadedMetaMap = {}
   self.convertedMap = {}
   
   frontInterface.setFront( self )
   
   local loadedMap = {}
   for mod, modval in pairs( Depend.getLoadedMod(  ) ) do
      if mod == "lune.base.Testing" then
         loadedMap[mod] = modval
      end
      
      if option.testing and modval['__enableTest'] or not option.testing and modval['__enableTest'] then
         loadedMap[mod] = modval
      end
      
   end
   
   self.preloadedModMap = loadedMap
end
function Front:getLoadInfo( mod )

   if self.option.testing then
      return self.loadedMapTest[mod]
   end
   
   return self.loadedMap[mod]
end
function Front:setLoadInfo( mod, info )

   if self.option.testing then
      self.loadedMapTest[mod] = info
   end
   
   self.loadedMap[mod] = info
end
function Front.setmeta( obj )
  setmetatable( obj, { __index = Front  } )
end


function Front:error( message )

   Util.errorLog( message )
   Util.printStackTrace(  )
   os.exit( 1 )
end


function Front:loadLua( path )

   local chunk, err = loadfile( path )
   if chunk ~= nil then
      return _lune.unwrap( chunk(  ))
   end
   
   Util.errorLog( _lune.unwrapDefault( err, string.format( "load error -- %s.", path)) )
   return nil
end


local function createPaser( path, mod )

   local parser = Parser.StreamParser.create( path, false, mod )
   do
      local _exp = parser
      if _exp ~= nil then
         return _exp
      end
   end
   
   error( "failed to open " .. path )
end

function Front:scriptPath2Module( path )

   return Util.scriptPath2ModuleFromProjDir( path, self.option:get_projDir() )
end


function Front:createPaser( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   return createPaser( scriptPath, mod )
end


function Front:createAst( importModuleInfo, parser, mod, moduleId, analyzeModule, analyzeMode, pos )

   local transUnit = TransUnit.TransUnit.new(moduleId, importModuleInfo, convLua.MacroEvalImp.new(), analyzeModule, analyzeMode, pos, self.option.targetLuaVer, self.option.transCtrlInfo)
   return transUnit:createAST( parser, false, mod )
end


function Front:convert( ast, streamName, stream, metaStream, convMode, inMacro )

   local conv = convLua.createFilter( streamName, stream, metaStream, convMode, inMacro, ast:get_moduleTypeInfo(), ast:get_processInfo(), ast:get_moduleSymbolKind(), self.option.useLuneModule, self.option.targetLuaVer, self.option.testing, self.option.useIpairs )
   ast:get_node():processFilter( conv, convLua.Opt.new(ast:get_node()) )
end


local function loadFromChunk( chunk, err )

   if err ~= nil then
      Util.errorLog( err )
   end
   
   if chunk ~= nil then
      do
         local work = chunk(  )
         if work ~= nil then
            return work
         end
      end
      
      return nil
   end
   
   error( "failed to error" )
end

local function loadFromLuaTxt( txt )

   return loadFromChunk( _lune.loadstring51( txt ) )
end
_moduleObj.loadFromLuaTxt = loadFromLuaTxt

function LoadInfo.FromLuaAndMeta( lnsPath, luaCode, metaTxt )

   local meta = _lune.unwrap( loadFromLuaTxt( metaTxt ))
   local code = _lune.unwrap( loadFromLuaTxt( luaCode ))
   return LoadInfo.new(code, frontInterface.ModuleMeta.new(meta, lnsPath))
end


local function byteCompileFromLuaTxt( txt, stripDebugInfo )

   local chunk, err = _lune.loadstring51( txt )
   if chunk ~= nil then
      return string.dump( chunk, stripDebugInfo )
   end
   
   error( _lune.unwrapDefault( err, "load error") )
end

function Front:convertFromAst( ast, streamName, mode )

   local stream = Util.memStream.new()
   local metaStream = Util.memStream.new()
   
   self:convert( ast, streamName, stream, metaStream, mode, false )
   
   return metaStream:get_txt(), stream:get_txt()
end

function Front:loadFromLnsTxt( importModuleInfo, name, txt )

   local _
   local transUnit = TransUnit.TransUnit.new(frontInterface.ModuleId.tempId, importModuleInfo, convLua.MacroEvalImp.new(), nil, nil, nil, self.option.targetLuaVer, self.option.transCtrlInfo)
   local stream = Parser.TxtStream.new(txt)
   local parser = Parser.StreamParser.new(stream, name, false, nil)
   
   local ast = transUnit:createAST( parser, false, nil )
   
   local _6024, luaTxt = self:convertFromAst( ast, name, convLua.ConvMode.Exec )
   return _lune.unwrap( loadFromLuaTxt( luaTxt ))
end


local DependMetaInfo = {}
setmetatable( DependMetaInfo, { ifList = {Mapping,} } )
function DependMetaInfo.setmeta( obj )
  setmetatable( obj, { __index = DependMetaInfo  } )
end
function DependMetaInfo.new( use, buildId )
   local obj = {}
   DependMetaInfo.setmeta( obj )
   if obj.__init then
      obj:__init( use, buildId )
   end
   return obj
end
function DependMetaInfo:__init( use, buildId )

   self.use = use
   self.buildId = buildId
end
function DependMetaInfo:_toMap()
  return self
end
function DependMetaInfo._fromMap( val )
  local obj, mes = DependMetaInfo._fromMapSub( {}, val )
  if obj then
     DependMetaInfo.setmeta( obj )
  end
  return obj, mes
end
function DependMetaInfo._fromStem( val )
  return DependMetaInfo._fromMap( val )
end

function DependMetaInfo._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "use", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "buildId", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local MetaForBuildId = {}
setmetatable( MetaForBuildId, { ifList = {Mapping,} } )
function MetaForBuildId:createModuleId(  )

   return frontInterface.ModuleId.createIdFromTxt( self.__buildId )
end
function MetaForBuildId.setmeta( obj )
  setmetatable( obj, { __index = MetaForBuildId  } )
end
function MetaForBuildId.new( __buildId, __dependModuleMap, __subModuleMap, __enableTest )
   local obj = {}
   MetaForBuildId.setmeta( obj )
   if obj.__init then
      obj:__init( __buildId, __dependModuleMap, __subModuleMap, __enableTest )
   end
   return obj
end
function MetaForBuildId:__init( __buildId, __dependModuleMap, __subModuleMap, __enableTest )

   self.__buildId = __buildId
   self.__dependModuleMap = __dependModuleMap
   self.__subModuleMap = __subModuleMap
   self.__enableTest = __enableTest
end
function MetaForBuildId:_toMap()
  return self
end
function MetaForBuildId._fromMap( val )
  local obj, mes = MetaForBuildId._fromMapSub( {}, val )
  if obj then
     MetaForBuildId.setmeta( obj )
  end
  return obj, mes
end
function MetaForBuildId._fromStem( val )
  return MetaForBuildId._fromMap( val )
end

function MetaForBuildId._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "__buildId", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "__dependModuleMap", func = _lune._toMap, nilable = false, child = { { func = _lune._toStr, nilable = false, child = {} }, 
{ func = DependMetaInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "__subModuleMap", func = _lune._toList, nilable = false, child = { { func = _lune._toStr, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "__enableTest", func = _lune._toBool, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


function MetaForBuildId.LoadFromMeta( metaPath )

   do
      local fileObj = io.open( metaPath )
      if fileObj ~= nil then
         local luaCode = fileObj:read( "*a" )
         fileObj:close(  )
         if luaCode ~= nil then
            local meta = MetaForBuildId._fromStem( (loadFromLuaTxt( luaCode ) ) )
            return meta, luaCode
         end
         
      end
   end
   
   return nil, nil
end


local function getMetaInfo( lnsPath, mod, outdir )

   local moduleMetaPath = lnsPath
   if outdir ~= nil then
      moduleMetaPath = string.format( "%s/%s", outdir, (mod:gsub( "%.", "/" ) ))
   end
   
   moduleMetaPath = moduleMetaPath:gsub( "%.lns$", ".meta" )
   do
      local meta, metaCode = MetaForBuildId.LoadFromMeta( moduleMetaPath )
      if meta ~= nil and  metaCode ~= nil then
         return meta, moduleMetaPath, metaCode
      end
   end
   
   return nil, moduleMetaPath, ""
end

function Front:searchModuleFile( mod, suffix, addPath )
   local __func__ = '@lune.@base.@front.Front.searchModuleFile'

   do
      local _matchExp = self.gomodMap:convLocalModulePath( mod, suffix )
      if _matchExp[1] == GoMod.GoModResult.NotGo[1] then
      
      elseif _matchExp[1] == GoMod.GoModResult.NotFound[1] then
      
         return nil
      elseif _matchExp[1] == GoMod.GoModResult.Found[1] then
         local info = _matchExp[2][1]
      
         return info:get_path()
      end
   end
   
   
   local lnsSearchPath = package.path
   if addPath ~= nil then
      lnsSearchPath = string.format( "%s/?%s;%s", addPath, suffix, package.path )
   end
   
   lnsSearchPath = lnsSearchPath:gsub( "%.lua$", suffix )
   lnsSearchPath = lnsSearchPath:gsub( "%.lua;", suffix .. ";" )
   
   local foundPath = Depend.searchpath( mod, lnsSearchPath )
   if  nil == foundPath then
      local _foundPath = foundPath
   
      local latestProjRoot = self.gomodMap:getLatestProjRoot(  )
      if  nil == latestProjRoot then
         local _latestProjRoot = latestProjRoot
      
         return nil
      end
      
      local latestProjSearchPath = Util.pathJoin( latestProjRoot, "?" .. suffix )
      do
         local _exp = Depend.searchpath( mod, latestProjSearchPath )
         if _exp ~= nil then
            foundPath = _exp
         else
            Log.log( Log.Level.Err, __func__, 344, function (  )
            
               return string.format( "not found at %s", latestProjSearchPath)
            end )
            
            return nil
         end
      end
      
   end
   
   return (foundPath:gsub( "^./", "" ) )
end

local function getModuleId( lnsPath, mod, outdir, metaInfo )

   local buildCount = 0
   local fileTime = Depend.getFileLastModifiedTime( lnsPath )
   if  nil == fileTime then
      local _fileTime = fileTime
   
      return frontInterface.ModuleId.tempId
   end
   
   if not metaInfo then
      metaInfo = getMetaInfo( lnsPath, mod, outdir )
   end
   
   if metaInfo ~= nil then
      local buildId = metaInfo:createModuleId(  )
      buildCount = buildId:get_buildCount()
   end
   
   return frontInterface.ModuleId.createId( fileTime, buildCount )
end

local ModuleUptodate = {}
ModuleUptodate._name2Val = {}
function ModuleUptodate:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "ModuleUptodate.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function ModuleUptodate._from( val )
   return _lune._AlgeFrom( ModuleUptodate, val )
end

ModuleUptodate.NeedTouch = { "NeedTouch", {{ func=_lune._toStr, nilable=false, child={} },{ func=MetaForBuildId._fromMap, nilable=false, child={} }}}
ModuleUptodate._name2Val["NeedTouch"] = ModuleUptodate.NeedTouch
ModuleUptodate.NeedUpdate = { "NeedUpdate"}
ModuleUptodate._name2Val["NeedUpdate"] = ModuleUptodate.NeedUpdate
ModuleUptodate.Uptodate = { "Uptodate", {{ func=MetaForBuildId._fromMap, nilable=false, child={} }}}
ModuleUptodate._name2Val["Uptodate"] = ModuleUptodate.Uptodate


function Front:getModuleIdAndCheckUptodate( lnsPath, mod )
   local __func__ = '@lune.@base.@front.Front.getModuleIdAndCheckUptodate'

   local uptodate = _lune.newAlge( ModuleUptodate.NeedUpdate)
   
   do
      local _matchExp = self.option.transCtrlInfo.uptodateMode
      if _matchExp[1] == Types.CheckingUptodateMode.ForceAll[1] then
      
         return frontInterface.ModuleId.tempId, uptodate
      elseif _matchExp[1] == Types.CheckingUptodateMode.Force1[1] then
         local forceMod = _matchExp[2][1]
      
         if mod == forceMod then
            return frontInterface.ModuleId.tempId, uptodate
         end
         
      elseif _matchExp[1] == Types.CheckingUptodateMode.Normal[1] then
      
         
      elseif _matchExp[1] == Types.CheckingUptodateMode.Touch[1] then
      
         
      end
   end
   
   
   local function checkDependUptodate( metaTime, metaInfo, metaCode )
      local __func__ = '@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate'
   
      
      for depMod, dependItem in pairs( metaInfo.__dependModuleMap ) do
         
         local modMetaPath = self:searchModuleFile( depMod, ".meta", self.option.outputDir )
         if  nil == modMetaPath then
            local _modMetaPath = modMetaPath
         
            
            Log.log( Log.Level.Debug, __func__, 440, function (  )
            
               return "NeedUpdate"
            end )
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         local time = Depend.getFileLastModifiedTime( modMetaPath )
         if  nil == time then
            local _time = time
         
            
            Log.log( Log.Level.Debug, __func__, 445, function (  )
            
               return "NeedUpdate"
            end )
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         if time > metaTime then
            
            local dependMeta = MetaForBuildId.LoadFromMeta( modMetaPath )
            if  nil == dependMeta then
               local _dependMeta = dependMeta
            
               Log.log( Log.Level.Debug, __func__, 453, function (  )
               
                  return "NeedUpdate"
               end )
               
               return _lune.newAlge( ModuleUptodate.NeedUpdate)
            end
            
            local orgMetaModuleId = frontInterface.ModuleId.createIdFromTxt( dependItem.buildId )
            local metaModuleId = dependMeta:createModuleId(  )
            
            if metaModuleId:get_buildCount() ~= 0 and metaModuleId:get_buildCount() ~= orgMetaModuleId:get_buildCount() then
               
               Log.log( Log.Level.Debug, __func__, 463, function (  )
               
                  return string.format( "NeedUpdate: %s, %d, %d", modMetaPath, metaModuleId:get_buildCount(), orgMetaModuleId:get_buildCount())
               end )
               
               return _lune.newAlge( ModuleUptodate.NeedUpdate)
            end
            
         end
         
      end
      
      if self.option.transCtrlInfo.uptodateMode == _lune.newAlge( Types.CheckingUptodateMode.Touch) then
         return _lune.newAlge( ModuleUptodate.NeedTouch, {metaCode,metaInfo})
      end
      
      return _lune.newAlge( ModuleUptodate.Uptodate, {metaInfo})
   end
   
   local metaInfo, metaPath, metaCode = getMetaInfo( lnsPath, mod, self.option.outputDir )
   
   if metaInfo ~= nil then
      if metaInfo.__enableTest == self.option.testing then
         local buildId = frontInterface.ModuleId.createIdFromTxt( metaInfo.__buildId )
         if buildId ~= frontInterface.ModuleId.tempId then
            local lnsTime = Depend.getFileLastModifiedTime( lnsPath )
            local metaTime = Depend.getFileLastModifiedTime( metaPath )
            if lnsTime ~= nil and metaTime ~= nil then
               if lnsTime == buildId:get_modTime() then
                  uptodate = checkDependUptodate( metaTime, metaInfo, metaCode )
               end
               
            end
            
         end
         
      else
       
         
      end
      
   else
      Log.log( Log.Level.Debug, __func__, 501, function (  )
      
         return "not found meta"
      end )
      
   end
   
   
   local moduleId = getModuleId( lnsPath, mod, self.option.outputDir, metaInfo )
   if moduleId == frontInterface.ModuleId.tempId then
      Util.err( string.format( "not found -- %s", lnsPath) )
   end
   
   return moduleId, uptodate
end

function Front:convertLns2LuaCode( importModuleInfo, stream, streamName )

   local _
   local mod = self:scriptPath2Module( streamName )
   local ast = self:createAst( importModuleInfo, Parser.StreamParser.new(stream, streamName, false, nil), mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, TransUnit.AnalyzeMode.Compile )
   
   local _6187, luaTxt = self:convertFromAst( ast, streamName, convLua.ConvMode.Exec )
   
   return luaTxt
end


function Front:loadParserToLuaCode( importModuleInfo, parser, mod )
   local __func__ = '@lune.@base.@front.Front.loadParserToLuaCode'

   local path = parser:getStreamName(  )
   
   local ast = self:createAst( importModuleInfo, parser, mod, getModuleId( path, mod ), nil, TransUnit.AnalyzeMode.Compile, nil )
   self.mod2ast:add( mod, ast )
   
   local metaTxt, luaTxt = self:convertFromAst( ast, path, convLua.ConvMode.Exec )
   Log.log( Log.Level.Trace, __func__, 548, function (  )
   
      return string.format( "Meta = %s", metaTxt)
   end )
   
   
   if self.option.updateOnLoad then
      local function saveFile( suffix, txt, byteCompile, stripDebugInfo, checkUpdate )
      
         local newpath = ""
         do
            local dir = self.option.outputDir
            if dir ~= nil then
               newpath = string.format( "%s/%s%s", dir, mod:gsub( "%.", "/" ), suffix)
            else
               newpath = path:gsub( ".lns$", suffix )
            end
         end
         
         local saveTxt = txt
         if byteCompile then
            saveTxt = byteCompileFromLuaTxt( saveTxt, stripDebugInfo )
         end
         
         if not forceUpdateMeta and checkUpdate then
            do
               local fileObj = io.open( newpath )
               if fileObj ~= nil then
                  local oldTxt = fileObj:read( "*a" )
                  if saveTxt == oldTxt then
                     return 
                  end
                  
               end
            end
            
         end
         
         do
            local fileObj = io.open( newpath, "w" )
            if fileObj ~= nil then
               fileObj:write( saveTxt )
               fileObj:close(  )
            end
         end
         
      end
      saveFile( ".lua", luaTxt, self.option.byteCompile, self.option.stripDebugInfo, false )
      saveFile( ".meta", metaTxt, self.option.byteCompile, true, true )
   end
   
   
   local meta = frontInterface.ModuleMeta.new(_lune.unwrap( loadFromLuaTxt( metaTxt )), path)
   return meta, luaTxt
end


function Front:loadFileToLuaCode( importModuleInfo, path, mod )
   local __func__ = '@lune.@base.@front.Front.loadFileToLuaCode'

   Log.log( Log.Level.Log, __func__, 597, function (  )
      local __func__ = '@lune.@base.@front.Front.loadFileToLuaCode.<anonymous>'
   
      return string.format( "%s: %s", __func__, mod)
   end )
   
   
   return self:loadParserToLuaCode( importModuleInfo, createPaser( path, mod ), mod )
end


function Front:loadFile( importModuleInfo, path, mod )
   local __func__ = '@lune.@base.@front.Front.loadFile'

   Log.log( Log.Level.Info, __func__, 612, function (  )
      local __func__ = '@lune.@base.@front.Front.loadFile.<anonymous>'
   
      return string.format( "start %s:%s", __func__, mod)
   end )
   
   
   local meta, luaTxt = self:loadFileToLuaCode( importModuleInfo, path, mod )
   
   do
      local preLoadInfo = self.preloadedModMap[mod]
      if preLoadInfo ~= nil then
         return meta, preLoadInfo
      end
   end
   
   return meta, _lune.unwrap( loadFromLuaTxt( luaTxt ))
   
end


function Front:searchModule( mod )

   return self:searchModuleFile( mod, ".lns", nil )
end


function Front:searchLuaFile( moduleFullName, addSearchPath )

   return self:searchModuleFile( moduleFullName, ".lua", addSearchPath )
   
end


function Front:checkUptodateMeta( lnsPath, metaPath, addSearchPath )
   local __func__ = '@lune.@base.@front.Front.checkUptodateMeta'

   local metaObj = self:loadLua( metaPath )
   if  nil == metaObj then
      local _metaObj = metaObj
   
      Log.log( Log.Level.Warn, __func__, 653, function (  )
      
         return string.format( "load error -- %s", metaPath)
      end )
      
      return nil
   end
   
   local meta = metaObj
   if meta.__formatVersion ~= Ver.metaVersion then
      Log.log( Log.Level.Warn, __func__, 658, function (  )
      
         return string.format( "unmatch meta version -- %s", metaPath)
      end )
      
      return nil
   end
   
   if meta.__hasTest then
      
      if meta.__enableTest ~= self.option.testing then
         Log.log( Log.Level.Warn, __func__, 664, function (  )
         
            return string.format( "unmatch test setting -- %s", metaPath)
         end )
         
         return nil
      end
      
   end
   
   
   for moduleFullName, _6278 in pairs( meta.__dependModuleMap ) do
      do
         local moduleLnsPath = self:searchModule( moduleFullName )
         if moduleLnsPath ~= nil then
            do
               local moduleLuaPath = self:searchLuaFile( moduleFullName, addSearchPath )
               if moduleLuaPath ~= nil then
                  if not Util.getReadyCode( moduleLnsPath, metaPath ) then
                     
                     Log.log( Log.Level.Warn, __func__, 675, function (  )
                     
                        return string.format( "not ready -- %s, %s", moduleLnsPath, metaPath)
                     end )
                     
                     return nil
                  end
                  
                  local moduleMetaPath = moduleLuaPath:gsub( "%.lua$", ".meta" )
                  if Depend.existFile( moduleMetaPath ) and not Util.getReadyCode( moduleMetaPath, metaPath ) then
                     Log.log( Log.Level.Warn, __func__, 683, function (  )
                     
                        return string.format( "not ready -- %s, %s", moduleMetaPath, metaPath)
                     end )
                     
                     return nil
                  end
                  
               else
                  Log.log( Log.Level.Warn, __func__, 688, function (  )
                  
                     return string.format( "not found .lua file for -- %s", moduleFullName)
                  end )
                  
                  return nil
               end
            end
            
         else
            Log.log( Log.Level.Warn, __func__, 693, function (  )
            
               return string.format( "not found .lns file -- %s", moduleFullName)
            end )
            
            return nil
         end
      end
      
   end
   
   return frontInterface.ModuleMeta.new(meta, lnsPath)
end


function Front:loadModule( mod )
   local __func__ = '@lune.@base.@front.Front.loadModule'

   local orgMod = mod
   mod = self.gomodMap:getLuaModulePath( mod )
   
   if not self:getLoadInfo( mod ) then
      do
         local luaTxt = self.convertedMap[mod]
         if luaTxt ~= nil then
            do
               local meta = self.loadedMetaMap[mod]
               if meta ~= nil then
                  self:setLoadInfo( mod, LoadInfo.new(_lune.unwrap( loadFromLuaTxt( luaTxt )), meta) )
               else
                  error( string.format( "nothing meta -- %s", mod) )
               end
            end
            
         else
            
            do
               local lnsPath = self:searchModule( orgMod )
               if lnsPath ~= nil then
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  
                  do
                     local dir = self.option.outputDir
                     if dir ~= nil then
                        luaPath = self:searchLuaFile( orgMod, dir )
                     end
                  end
                  
                  
                  local loadVal = nil
                  if luaPath ~= nil then
                     if Util.getReadyCode( lnsPath, luaPath ) then
                        local metaPath = string.gsub( luaPath, "%.lua$", ".meta" )
                        if Util.getReadyCode( lnsPath, metaPath ) then
                           do
                              local preLoadInfo = self.preloadedModMap[mod]
                              if preLoadInfo ~= nil then
                                 loadVal = preLoadInfo
                              else
                                 loadVal = self:loadLua( luaPath )
                              end
                           end
                           
                           do
                              local _exp = loadVal
                              if _exp ~= nil then
                                 do
                                    local meta = self:checkUptodateMeta( lnsPath, metaPath, self.option.outputDir )
                                    if meta ~= nil then
                                       self:setLoadInfo( mod, LoadInfo.new(_exp, meta) )
                                    else
                                       loadVal = nil
                                    end
                                 end
                                 
                              end
                           end
                           
                        end
                        
                     end
                     
                  end
                  
                  if loadVal == nil then
                     local meta, workVal = self:loadFile( frontInterface.ImportModuleInfo.new(), lnsPath, mod )
                     self:setLoadInfo( mod, LoadInfo.new(workVal, meta) )
                  end
                  
               else
                  
                  if _lune._Set_has(self.bindModuleSet, mod ) then
                     Log.log( Log.Level.Warn, __func__, 760, function (  )
                     
                        return string.format( "load from the binding -- %s", mod)
                     end )
                     
                     local workMod = require( mod )
                     local meta = frontInterface.ModuleMeta.new(_lune.unwrap( loadFromLuaTxt( "return {}" )), mod:gsub( "%.", "/" ) .. ".lns")
                     self:setLoadInfo( mod, LoadInfo.new(workMod, meta) )
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   do
      local _exp = self:getLoadInfo( mod )
      if _exp ~= nil then
         return _exp.mod, _exp.meta
      end
   end
   
   error( string.format( "load error, %s", mod) )
end


function Front:getLuaModulePath( mod )

   return self.gomodMap:getLuaModulePath( mod )
end


function Front:loadMeta( importModuleInfo, mod )
   local __func__ = '@lune.@base.@front.Front.loadMeta'

   local orgMod = mod
   mod = self.gomodMap:getLuaModulePath( mod )
   
   if self.loadedMetaMap[mod] == nil then
      do
         local _exp = self:getLoadInfo( mod )
         if _exp ~= nil then
            self.loadedMetaMap[mod] = _exp.meta
         else
            do
               local lnsPath = self:searchModule( orgMod )
               if lnsPath ~= nil then
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  
                  do
                     local dir = self.option.outputDir
                     if dir ~= nil then
                        luaPath = self:searchLuaFile( orgMod, dir )
                     end
                  end
                  
                  
                  local meta = nil
                  if luaPath ~= nil then
                     local forceFlag
                     
                     do
                        local _matchExp = self.option.transCtrlInfo.uptodateMode
                        if _matchExp[1] == Types.CheckingUptodateMode.ForceAll[1] then
                        
                           forceFlag = true
                        elseif _matchExp[1] == Types.CheckingUptodateMode.Force1[1] then
                           local forceMod = _matchExp[2][1]
                        
                           forceFlag = mod == forceMod
                        elseif _matchExp[1] == Types.CheckingUptodateMode.Normal[1] then
                        
                           forceFlag = false
                        elseif _matchExp[1] == Types.CheckingUptodateMode.Touch[1] then
                        
                           forceFlag = false
                        end
                     end
                     
                     if not forceFlag then
                        if Util.getReadyCode( lnsPath, luaPath ) then
                           local metaPath = string.gsub( luaPath, "%.lua$", ".meta" )
                           if Util.getReadyCode( lnsPath, metaPath ) then
                              meta = self:checkUptodateMeta( lnsPath, metaPath, self.option.outputDir )
                           else
                            
                              Log.log( Log.Level.Warn, __func__, 830, function (  )
                              
                                 return string.format( "%s not ready meta %s, %s", orgMod, lnsPath, metaPath)
                              end )
                              
                           end
                           
                        else
                         
                           Log.log( Log.Level.Warn, __func__, 834, function (  )
                           
                              return string.format( "%s not ready lua %s, %s", orgMod, lnsPath, luaPath)
                           end )
                           
                        end
                        
                     else
                      
                        Log.log( Log.Level.Warn, __func__, 838, function (  )
                        
                           return string.format( "force analyze -- %s", orgMod)
                        end )
                        
                     end
                     
                  else
                     Log.log( Log.Level.Warn, __func__, 841, function (  )
                     
                        return string.format( "%s not found lua in %s", orgMod, tostring( self.option.outputDir))
                     end )
                     
                  end
                  
                  if meta ~= nil then
                     self.loadedMetaMap[mod] = meta
                  else
                     local metawork, luaTxt = self:loadFileToLuaCode( importModuleInfo, lnsPath, orgMod )
                     self.loadedMetaMap[mod] = metawork
                     self.convertedMap[mod] = luaTxt
                  end
                  
               else
                  do
                     local lnsCode = Depend.getBindLns( mod )
                     if lnsCode ~= nil then
                        local path = mod:gsub( "%.", "/" ) .. ".lns"
                        local parser = Parser.StreamParser.new(Parser.TxtStream.new(lnsCode), path, false, nil)
                        local meta, luaTxt = self:loadParserToLuaCode( importModuleInfo, parser, mod )
                        self.loadedMetaMap[mod] = meta
                        self.convertedMap[mod] = luaTxt
                     end
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   
   do
      local meta = self.loadedMetaMap[mod]
      if meta ~= nil then
         return meta
      end
   end
   
   return nil
end


function Front:dumpTokenize( scriptPath )

   local parser = self:createPaser( scriptPath )
   while true do
      local token = parser:getToken(  )
      if  nil == token then
         local _token = token
      
         break
      end
      
      print( Types.TokenKind:_getTxt( token.kind)
      , token.pos.lineNo, token.pos.column, token.txt )
   end
   
end


function Front:dumpAst( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   Depend.profile( self.option.validProf, function (  )
   
      local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
      ast:get_node():processFilter( dumpNode.createFilter( ast:get_moduleTypeInfo(), ast:get_processInfo(), io.stdout ), dumpNode.Opt.new("", 0) )
   end, scriptPath .. ".profi" )
end


function Front:format( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
   ast:get_node():processFilter( Formatter.createFilter( ast:get_moduleTypeInfo(), io.stdout ), Formatter.Opt.new(ast:get_node()) )
end


function Front:checkDiag( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   Util.setErrorCode( 0 )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), nil, TransUnit.AnalyzeMode.Diag )
end


function Front:complete( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), self.option.analyzeModule, TransUnit.AnalyzeMode.Complete, self.option.analyzePos )
end


function Front:inquire( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), self.option.analyzeModule, TransUnit.AnalyzeMode.Inquire, self.option.analyzePos )
end


function Front:createGlue( scriptPath )

   local mod = self:scriptPath2Module( scriptPath )
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser( scriptPath ), mod, getModuleId( scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
   local filter = glueFilter.createFilter( self.option.outputDir )
   ast:get_node():processFilter( filter, 0 )
end




function Front:convertLuaToStreamFromScript( parser, moduleId, uptodate, convMode, path, mod, byteCompile, stripDebugInfo, openOStream, closeOStream )

   local function outputDependInfo( stream, metaInfo )
   
      if stream ~= nil then
         if metaInfo ~= nil then
            local dependInfo = OutputDepend.DependInfo.new(mod)
            for dependMod, _6483 in pairs( metaInfo.__dependModuleMap ) do
               dependInfo:addImpotModule( dependMod )
            end
            
            for __index, subMod in ipairs( metaInfo.__subModuleMap ) do
               dependInfo:addSubMod( subMod )
            end
            
            dependInfo:output( stream )
         else
            Util.err( "metaInfo is nil" )
         end
         
      end
      
   end
   
   local streamDst, metaStreamDst, dependsStreamDst = openOStream( uptodate )
   local streamMem = Util.memStream.new()
   local metaStreamMem = Util.memStream.new()
   local dependsStreamMem = Util.memStream.new()
   
   local stream
   
   local metaStream
   
   local dependsStream
   
   
   if Str.isValidStrBuilder(  ) then
      stream = streamMem
      metaStream = metaStreamMem
      dependsStream = dependsStreamMem
   else
    
      stream = streamDst
      metaStream = metaStreamDst
      dependsStream = dependsStreamDst
   end
   
   
   local retAst = nil
   
   do
      local _matchExp = uptodate
      if _matchExp[1] == ModuleUptodate.Uptodate[1] then
         local metaInfo = _matchExp[2][1]
      
         Util.errorLog( "uptodate -- " .. path )
         outputDependInfo( dependsStream, metaInfo )
      elseif _matchExp[1] == ModuleUptodate.NeedUpdate[1] then
      
         if stream ~= nil and metaStream ~= nil then
            local ast = self:createAst( frontInterface.ImportModuleInfo.new(), createPaser( path, mod ), mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
            retAst = ast
            
            if dependsStream ~= nil then
               ast:get_node():processFilter( OutputDepend.createFilter( dependsStream ), 1 )
            end
            
            
            local outStream = stream
            local oMetaStream = metaStream
            
            local byteStream = Util.memStream.new()
            local byteMetaStream = Util.memStream.new()
            if byteCompile then
               outStream = byteStream
               oMetaStream = byteMetaStream
            end
            
            
            self:convert( ast, path, outStream, oMetaStream, convMode, false )
            
            if byteCompile then
               stream:write( byteCompileFromLuaTxt( byteStream:get_txt(), stripDebugInfo ) )
               if metaStream ~= stream then
                  metaStream:write( byteCompileFromLuaTxt( byteMetaStream:get_txt(), true ) )
               end
               
            end
            
         else
            Util.err( "failed to open lua stream or meta stream" )
         end
         
      elseif _matchExp[1] == ModuleUptodate.NeedTouch[1] then
         local metaCode = _matchExp[2][1]
         local metaInfo = _matchExp[2][2]
      
         Util.errorLog( "touch -- " .. path )
         if self.option.mode == Option.ModeKind.SaveMeta then
            if metaStream ~= nil then
               metaStream:write( metaCode )
            else
               Util.err( "failed to open meta stream" )
            end
            
         end
         
         
         outputDependInfo( dependsStream, metaInfo )
      end
   end
   
   
   if Str.isValidStrBuilder(  ) then
      
      
      if streamDst ~= nil then
         streamDst:write( streamMem:get_txt() )
      end
      
      
      
      if metaStreamDst ~= nil then
         metaStreamDst:write( metaStreamMem:get_txt() )
      end
      
      
      
      if dependsStreamDst ~= nil then
         dependsStreamDst:write( dependsStreamMem:get_txt() )
      end
      
      
   end
   
   
   if closeOStream ~= nil then
      closeOStream( stream, metaStream, dependsStream )
   end
   
   
   return retAst
end


function Front:getGoAppName(  )

   local appName = self.option.appName
   if  nil == appName then
      local _appName = appName
   
      appName = self.gomodMap:get_name()
   end
   
   return appName
end


function Front:createGoOption( scriptPath )

   local packageName
   
   do
      local _exp = self.option.packageName
      if _exp ~= nil then
         packageName = _exp
      else
         if not scriptPath:find( "/" ) then
            packageName = "main"
         else
          
            local parentPath = scriptPath:gsub( "/[^/]+$", "" ):gsub( ".*/", "" )
            if #parentPath == 0 then
               packageName = "main"
            elseif parentPath == "." then
               packageName = "main"
            elseif parentPath == ".." then
               packageName = "main"
            else
             
               packageName = parentPath:gsub( "[^%w]", "" )
            end
            
         end
         
      end
   end
   
   return convGo.Option.new(packageName, self:getGoAppName(  ), self.option.mainModule)
end


function Front:convertToLua( scriptPath, stream )

   local mod = self:scriptPath2Module( scriptPath )
   local convMode = convLua.ConvMode.Convert
   if self.option.mode == Option.ModeKind.LuaMeta then
      convMode = convLua.ConvMode.ConvMeta
   end
   
   
   local parser = createPaser( scriptPath, mod )
   local ast = self:convertLuaToStreamFromScript( parser, frontInterface.ModuleId.tempId, _lune.newAlge( ModuleUptodate.NeedUpdate), convMode, scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( mode )
   
      return stream, stream, self.option:openDepend( nil )
   end, function ( luaCodeStream, metaStream, dependStream )
   
      if dependStream ~= nil then
         dependStream:close(  )
      end
      
   end )
   
   if ast ~= nil then
      do
         local _switchExp = self.option.convTo
         if _switchExp == Types.Lang.Go then
            local conv = convGo.createFilter( self.option.testing, "stdout", stream, ast, self:createGoOption( scriptPath ) )
            ast:get_node():processFilter( conv, convGo.Opt.new(ast:get_node()) )
         end
      end
      
      self.mod2ast:add( mod, ast )
   end
   
end

function Front:saveToGo( scriptPath, ast )

   local rootNode = _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )
   if  nil == rootNode then
      local _rootNode = rootNode
   
      return 
   end
   
   for pragma, __val in pairs( rootNode:get_luneHelperInfo().pragmaSet ) do
      do
         local _matchExp = pragma
         if _matchExp[1] == LuneControl.Pragma.limit_conv_code[1] then
            local codeSet = _matchExp[2][1]
         
            if not _lune._Set_has(codeSet, LuneControl.Code.Go ) then
               return 
            end
            
         end
      end
      
   end
   
   
   local path = scriptPath:gsub( "%.lns$", ".go" )
   
   do
      local dir = self.option.outputDir
      if dir ~= nil then
         path = string.format( "%s/%s", dir, path)
      end
   end
   
   local file = io.open( path, "w" )
   if  nil == file then
      local _file = file
   
      return 
   end
   
   local memStream = Util.memStream.new()
   local dstStream
   
   if Str.isValidStrBuilder(  ) then
      dstStream = memStream
   else
    
      dstStream = file
   end
   
   
   local conv = convGo.createFilter( self.option.testing, path, dstStream, ast, self:createGoOption( scriptPath ) )
   ast:get_node():processFilter( conv, convGo.Opt.new(ast:get_node()) )
   
   if Str.isValidStrBuilder(  ) then
      file:write( memStream:get_txt() )
   end
   
   
   file:close(  )
end


function Front:saveToC( scriptPath, ast )

   local cPath = scriptPath:gsub( "%.lns$", ".c" )
   
   local file = io.open( cPath, "w" )
   if  nil == file then
      local _file = file
   
      return 
   end
   
   
   local hPath = scriptPath:gsub( "%.lns$", ".h" )
   local hFile = io.open( hPath, "w" )
   if  nil == hFile then
      local _hFile = hFile
   
      return 
   end
   
   file:close(  )
   hFile:close(  )
end


function Front:outputBuiltin( scriptPath )

   local mod = self:scriptPath2Module( "lns_builtin" )
   
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), Parser.DummyParser.new(), mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, TransUnit.AnalyzeMode.Compile )
   
   self:saveToC( scriptPath, ast )
end


local UpdateInfo = {}
function UpdateInfo.setmeta( obj )
  setmetatable( obj, { __index = UpdateInfo  } )
end
function UpdateInfo.new( scriptPath, dependsPath, parser, moduleId, uptodate )
   local obj = {}
   UpdateInfo.setmeta( obj )
   if obj.__init then
      obj:__init( scriptPath, dependsPath, parser, moduleId, uptodate )
   end
   return obj
end
function UpdateInfo:__init( scriptPath, dependsPath, parser, moduleId, uptodate )

   self.scriptPath = scriptPath
   self.dependsPath = dependsPath
   self.parser = parser
   self.moduleId = moduleId
   self.uptodate = uptodate
end
function UpdateInfo:get_scriptPath()
   return self.scriptPath
end
function UpdateInfo:get_dependsPath()
   return self.dependsPath
end
function UpdateInfo:get_parser()
   return self.parser
end
function UpdateInfo:get_moduleId()
   return self.moduleId
end
function UpdateInfo:get_uptodate()
   return self.uptodate
end




function Front:saveToLua( updateInfo )

   local scriptPath = updateInfo:get_scriptPath()
   local dependsPath = updateInfo:get_dependsPath()
   local parser = updateInfo:get_parser()
   local moduleId = updateInfo:get_moduleId()
   local uptodate = updateInfo:get_uptodate()
   
   local function txt2ModuleId( txt )
   
      local buildIdTxt = txt:gsub( "^_moduleObj.__buildId = ", "" ):gsub( '"', "" )
      return frontInterface.ModuleId.createIdFromTxt( buildIdTxt )
   end
   
   local function checkDiff( oldStream, newStream )
      local __func__ = '@lune.@base.@front.Front.saveToLua.checkDiff'
   
      
      
      
      local headEndPos = 0
      local tailBeginPos = 0
      local oldBuildIdLine = ""
      local newBuildIdLine = ""
      while true do
         local newLine = newStream:read( "*l" )
         local oldLine = oldStream:read( "*l" )
         
         if oldLine ~= nil then
            if #oldBuildIdLine == 0 then
               if oldLine:find( "^_moduleObj.__buildId" ) then
                  oldBuildIdLine = oldLine
               end
               
            end
            
         end
         
         
         
         if newLine ~= nil then
            if #newBuildIdLine == 0 then
               if newLine:find( "^_moduleObj.__buildId" ) then
                  newBuildIdLine = newLine
               end
               
            end
            
         end
         
         
         
         if newLine ~= oldLine then
            local cont = false
            if newLine ~= nil and oldLine ~= nil then
               if oldLine:find( "^_moduleObj.__buildId" ) then
                  if newLine:find( "^_moduleObj.__buildId" ) then
                     tailBeginPos = newStream:get_lineNo()
                     cont = true
                  end
                  
               elseif oldLine:find( "^__dependModuleMap.*buildId =" ) and newLine:find( "^__dependModuleMap.*buildId =" ) then
                  local oldSub = oldLine:gsub( "buildId =.*", "" )
                  local newSub = newLine:gsub( "buildId =.*", "" )
                  if oldSub == newSub then
                     cont = true
                  end
                  
               end
               
            end
            
            if not cont then
               Log.log( Log.Level.Debug, __func__, 1342, function (  )
               
                  return string.format( "<%s>, <%s>", tostring( oldLine), tostring( newLine))
               end )
               
               return false, ""
            end
            
         else
          
            if tailBeginPos == 0 then
               headEndPos = newStream:get_lineNo()
            end
            
            if not oldLine then
               if tailBeginPos == 0 then
                  return true, oldStream:get_txt()
               end
               
               
               local oldBuildId = txt2ModuleId( oldBuildIdLine )
               local newBuildId = txt2ModuleId( newBuildIdLine )
               local worlBuildId = frontInterface.ModuleId.createId( newBuildId:get_modTime(), oldBuildId:get_buildCount() )
               local buildIdLine = string.format( "_moduleObj.__buildId = %q", worlBuildId:get_idStr())
               
               local txt = string.format( "%s%s\n%s", newStream:getSubstring( 1, headEndPos ), buildIdLine, newStream:getSubstring( tailBeginPos ))
               return true, txt
            end
            
         end
         
      end
      
   end
   
   local updateFlag = true
   local ast = nil
   
   local mod = self:scriptPath2Module( scriptPath )
   local luaPath = scriptPath:gsub( "%.lns$", ".lua" )
   local metaPath = scriptPath:gsub( "%.lns$", ".meta" )
   if self.option.outputDir then
      local filename = mod:gsub( "%.", "/" )
      luaPath = string.format( "%s/%s.lua", tostring( self.option.outputDir), filename)
      metaPath = string.format( "%s/%s.meta", tostring( self.option.outputDir), filename)
   end
   
   
   if luaPath == scriptPath then
      Util.errorLog( string.format( "%s is illegal filename.", luaPath) )
   else
    
      local convMode = convLua.ConvMode.ConvMeta
      local metaMemStream = Util.memStream.new()
      local metaFileObj = nil
      local tempMetaPath = metaPath .. ".tmp"
      
      ast = self:convertLuaToStreamFromScript( parser, moduleId, uptodate, convMode, scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( mode )
      
         local function openLuaStream(  )
         
            local fileObj = io.open( luaPath, "w" )
            if  nil == fileObj then
               local _fileObj = fileObj
            
               error( string.format( "write open error -- %s", luaPath) )
            end
            
            return fileObj
         end
         local function openStreams( luaFlag )
         
            local stream = nil
            if luaFlag then
               stream = openLuaStream(  )
            end
            
            
            if self.option.mode == Option.ModeKind.SaveMeta then
               do
                  local _exp = io.open( tempMetaPath, "w+" )
                  if _exp ~= nil then
                     metaFileObj = _exp
                  else
                     error( string.format( "write open error -- %s", metaPath) )
                  end
               end
               
            end
            
            return stream, metaMemStream
         end
         local stream = nil
         local metaStream = stream
         do
            local _matchExp = mode
            if _matchExp[1] == ModuleUptodate.Uptodate[1] then
               local metaInfo = _matchExp[2][1]
            
            elseif _matchExp[1] == ModuleUptodate.NeedTouch[1] then
               local metaCode = _matchExp[2][1]
               local metaInfo = _matchExp[2][2]
            
               stream, metaStream = openStreams( false )
            else 
               
                  stream, metaStream = openStreams( true )
            end
         end
         
         
         return stream, metaStream, self.option:openDepend( dependsPath )
      end, function ( stream, metaStream, dependStream )
      
         if stream ~= nil then
            stream:close(  )
         end
         
         if dependStream ~= nil then
            dependStream:close(  )
         end
         
         
         if metaFileObj ~= nil then
            local newMetaTxt = metaMemStream:get_txt()
            metaFileObj:write( newMetaTxt )
            metaFileObj:close(  )
            local oldMetaTxt = ""
            do
               local oldFileObj = io.open( metaPath )
               if oldFileObj ~= nil then
                  oldMetaTxt = _lune.unwrapDefault( oldFileObj:read( "*a" ), "")
                  oldFileObj:close(  )
               end
            end
            
            
            local sameFlag, txt = checkDiff( Parser.TxtStream.new(oldMetaTxt), Parser.TxtStream.new(newMetaTxt) )
            if not sameFlag then
               os.rename( tempMetaPath, metaPath )
            else
             
               os.remove( tempMetaPath )
               if txt ~= "" then
                  do
                     local fileObj = io.open( metaPath, "w" )
                     if fileObj ~= nil then
                        fileObj:write( txt )
                        fileObj:close(  )
                     end
                  end
                  
               else
                
                  updateFlag = false
               end
               
            end
            
         end
         
      end )
   end
   
   
   if updateFlag then
      scriptPath:gsub( "%.lns$", ".lua" )
   end
   
   
   if ast ~= nil then
      do
         local _switchExp = self.option.convTo
         if _switchExp == Types.Lang.C then
            self:saveToC( scriptPath, ast )
         elseif _switchExp == Types.Lang.Go then
            self:saveToGo( scriptPath, ast )
         end
      end
      
      self.mod2ast:add( mod, ast )
   end
   
   
   return updateFlag
end


function Front:outputBootC( scriptPath )

   
end


local function convertLnsCode2LuaCode( lnsCode, path )

   local option = Option.Option.new()
   option.scriptPath = path
   option.useLuneModule = Option.getRuntimeModule(  )
   option.useIpairs = true
   local front = Front.new(option)
   
   return front:convertLns2LuaCode( frontInterface.ImportModuleInfo.new(), Parser.TxtStream.new(lnsCode), path )
end
_moduleObj.convertLnsCode2LuaCode = convertLnsCode2LuaCode

local BuildMode = {}
BuildMode._name2Val = {}
function BuildMode:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "BuildMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function BuildMode._from( val )
   return _lune._AlgeFrom( BuildMode, val )
end

BuildMode.CreateAst = { "CreateAst"}
BuildMode._name2Val["CreateAst"] = BuildMode.CreateAst
BuildMode.Output = { "Output", {{}}}
BuildMode._name2Val["Output"] = BuildMode.Output
BuildMode.Save = { "Save"}
BuildMode._name2Val["Save"] = BuildMode.Save


function Front:build( buildMode, astCallback )

   local function createUpdateInfo( scriptPath, dependsPath )
   
      local mod = self:scriptPath2Module( scriptPath )
      local moduleId, uptodate = self:getModuleIdAndCheckUptodate( scriptPath, mod )
      local parser
      
      do
         local _matchExp = uptodate
         if _matchExp[1] == ModuleUptodate.NeedUpdate[1] then
         
            parser = createPaser( scriptPath, mod )
         elseif _matchExp[1] == ModuleUptodate.NeedTouch[1] then
            local _ = _matchExp[2][1]
            local _ = _matchExp[2][2]
         
            parser = nil
         elseif _matchExp[1] == ModuleUptodate.Uptodate[1] then
            local _ = _matchExp[2][1]
         
            parser = nil
         end
      end
      
      return UpdateInfo.new(scriptPath, dependsPath, parser, moduleId, uptodate)
   end
   
   local function process( updateInfo )
   
      do
         local _matchExp = buildMode
         if _matchExp[1] == BuildMode.Save[1] then
         
            self:saveToLua( updateInfo )
         elseif _matchExp[1] == BuildMode.Output[1] then
            local stream = _matchExp[2][1]
         
            self:convertToLua( updateInfo:get_scriptPath(), stream )
         elseif _matchExp[1] == BuildMode.CreateAst[1] then
         
            self:convertToLua( updateInfo:get_scriptPath(), Util.NullOStream.new() )
         end
      end
      
   end
   
   Depend.profile( self.option.validProf, function (  )
   
      self.mod2ast:clear(  )
      if self.option.scriptPath == "@-" then
         local updateList = {}
         while true do
            local line = io.stdin:read( "*l" )
            if  nil == line then
               local _line = line
            
               break
            end
            
            if #line > 0 then
               table.insert( updateList, createUpdateInfo( line, (line:gsub( ".lns$", ".d" ) ) ) )
            end
            
         end
         
         
         for __index, updateInfo in ipairs( updateList ) do
            local prev = os.clock(  )
            process( updateInfo )
            print( string.format( "%s:%g", updateInfo:get_scriptPath(), os.clock(  ) - prev) )
         end
         
      else
       
         process( createUpdateInfo( self.option.scriptPath, nil ) )
      end
      
      
      if astCallback ~= nil then
         for __index, key in ipairs( self.mod2ast:get_keyList() ) do
            astCallback( _lune.unwrap( self.mod2ast:get_map()[key]) )
         end
         
      end
      
   end, self.option.scriptPath .. ".profi" )
end


local function build( option, astCallback )

   local front = Front.new(option)
   front:build( _lune.newAlge( BuildMode.CreateAst), astCallback )
end
_moduleObj.build = build

function Front:exec(  )
   local __func__ = '@lune.@base.@front.Front.exec'

   Log.log( Log.Level.Trace, __func__, 1619, function (  )
   
      return Option.ModeKind:_getTxt( self.option.mode)
      
   end )
   
   
   do
      local _switchExp = self.option.mode
      if _switchExp == Option.ModeKind.Token then
         self:dumpTokenize( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Ast then
         self:dumpAst( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Format then
         self:format( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Diag then
         self:checkDiag( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Complete then
         self:complete( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Inquire then
         self:inquire( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Glue then
         self:createGlue( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Lua or _switchExp == Option.ModeKind.LuaMeta then
         self:convertToLua( self.option.scriptPath, io.stdout )
      elseif _switchExp == Option.ModeKind.Save or _switchExp == Option.ModeKind.SaveMeta then
         self:build( _lune.newAlge( BuildMode.Save), nil )
      elseif _switchExp == Option.ModeKind.Shebang then
         do
            local modObj = self:loadModule( self:scriptPath2Module( self.option.scriptPath ) )
            if modObj ~= nil then
               local code = Depend.runMain( modObj['__main'], self.option.shebangArgList )
               os.exit( code )
            end
         end
         
      elseif _switchExp == Option.ModeKind.Exec then
         local modObj = self:loadModule( self:scriptPath2Module( self.option.scriptPath ) )
         
         if self.option.testing then
            local code = [==[
local Testing = require( "lune.base.Testing" )
return function( path )
  Testing.run( path );
  Testing.outputAllResult( io.stdout );
end
]==]
            local loaded, mess = _lune.loadstring51( code )
            if loaded ~= nil then
               do
                  local mod = loaded(  )
                  if mod ~= nil then
                     (mod )( self:scriptPath2Module( self.option.scriptPath ) )
                  end
               end
               
            else
               print( mess )
            end
            
         end
         
      elseif _switchExp == Option.ModeKind.BootC then
         self:outputBootC( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Builtin then
         self:outputBuiltin( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.MkMain then
         local mod = self:scriptPath2Module( self.option.scriptPath )
         do
            local mess = convGo.outputGoMain( self:getGoAppName(  ), mod, self.option.testing, self.option.outputPath, self.option:get_runtimeOpt() )
            if mess ~= nil then
               Util.errorLog( mess )
            end
         end
         
      else 
         
            print( "illegal mode" )
      end
   end
   
end


local function exec( args )

   local version = _lune.unwrapDefault( tonumber( Depend.getLuaVersion(  ):gsub( "^[^%d]+", "" ), nil ), 0.0)
   
   if version < 5.1 then
      io.stderr:write( string.format( "LuneScript doesn't support this lua version(%s). %s\n", tostring( version), "please use the version >= 5.1." ) )
      os.exit( 1 )
   end
   
   
   local option = Option.analyze( args )
   local front = Front.new(option)
   
   front:exec(  )
end
_moduleObj.exec = exec
local function setFront( bindModuleList )

   local option = Option.createDefaultOption( "dummy.lns", nil )
   Front.new(option, bindModuleList)
end
_moduleObj.setFront = setFront



return _moduleObj
