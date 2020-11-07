--lune/base/front.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@front'
local _lune = {}
if _lune2 then
   _lune = _lune2
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

if not _lune2 then
   _lune2 = _lune
end
local Types = _lune.loadModule( 'lune.base.Types' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local convLua = _lune.loadModule( 'lune.base.convLua' )
local convCC = _lune.loadModule( 'lune.base.convCC' )
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
local Meta = _lune.loadModule( 'lune.base.Meta' )


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
function Front.new( option )
   local obj = {}
   Front.setmeta( obj )
   if obj.__init then obj:__init( option ); end
   return obj
end
function Front:__init(option) 
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

local function scriptPath2Module( path )

   return Util.scriptPath2Module( path )
end
_moduleObj.scriptPath2Module = scriptPath2Module

function Front:createPaser(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   return createPaser( self.option.scriptPath, mod )
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
   local parser = Parser.StreamParser.new(stream, name, false)
   
   local ast = transUnit:createAST( parser, false, nil )
   
   local _5676, luaTxt = self:convertFromAst( ast, name, convLua.ConvMode.Exec )
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

   local lnsSearchPath = package.path
   if addPath ~= nil then
      lnsSearchPath = string.format( "%s/?%s;%s", addPath, suffix, package.path )
   end
   
   
   lnsSearchPath = lnsSearchPath:gsub( "%.lua$", suffix )
   lnsSearchPath = lnsSearchPath:gsub( "%.lua;", suffix .. ";" )
   local foundPath = Depend.searchpath( mod, lnsSearchPath )
   if  nil == foundPath then
      local _foundPath = foundPath
   
      return nil
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
   
   if self.option.transCtrlInfo.uptodateMode == Types.CheckingUptodateMode.Force then
      return frontInterface.ModuleId.tempId, uptodate
   end
   
   local function checkDependUptodate( metaTime, metaInfo, metaCode )
      local __func__ = '@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate'
   
      
      for depMod, dependItem in pairs( metaInfo.__dependModuleMap ) do
         
         local modMetaPath = self:searchModuleFile( depMod, ".meta", self.option.outputDir )
         if  nil == modMetaPath then
            local _modMetaPath = modMetaPath
         
            
            Log.log( Log.Level.Debug, __func__, 379, function (  )
            
               return "NeedUpdate"
            end )
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         local time = Depend.getFileLastModifiedTime( modMetaPath )
         if  nil == time then
            local _time = time
         
            
            Log.log( Log.Level.Debug, __func__, 384, function (  )
            
               return "NeedUpdate"
            end )
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         if time > metaTime then
            
            local dependMeta = MetaForBuildId.LoadFromMeta( modMetaPath )
            if  nil == dependMeta then
               local _dependMeta = dependMeta
            
               Log.log( Log.Level.Debug, __func__, 392, function (  )
               
                  return "NeedUpdate"
               end )
               
               return _lune.newAlge( ModuleUptodate.NeedUpdate)
            end
            
            local orgMetaModuleId = frontInterface.ModuleId.createIdFromTxt( dependItem.buildId )
            local metaModuleId = dependMeta:createModuleId(  )
            
            if metaModuleId:get_buildCount() ~= 0 and metaModuleId:get_buildCount() ~= orgMetaModuleId:get_buildCount() then
               
               Log.log( Log.Level.Debug, __func__, 402, function (  )
               
                  return string.format( "NeedUpdate: %s, %d, %d", modMetaPath, metaModuleId:get_buildCount(), orgMetaModuleId:get_buildCount())
               end )
               
               return _lune.newAlge( ModuleUptodate.NeedUpdate)
            end
            
         end
         
      end
      
      if self.option.transCtrlInfo.uptodateMode == Types.CheckingUptodateMode.Touch then
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
      Log.log( Log.Level.Debug, __func__, 440, function (  )
      
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
   local mod = scriptPath2Module( streamName )
   local ast = self:createAst( importModuleInfo, Parser.StreamParser.new(stream, streamName, false), mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, TransUnit.AnalyzeMode.Compile )
   
   local _5821, luaTxt = self:convertFromAst( ast, streamName, convLua.ConvMode.Exec )
   
   return luaTxt
end


function Front:loadFileToLuaCode( importModuleInfo, path, mod )

   local ast = self:createAst( importModuleInfo, createPaser( path, mod ), mod, getModuleId( path, mod ), nil, TransUnit.AnalyzeMode.Compile, nil )
   
   local metaTxt, luaTxt = self:convertFromAst( ast, path, convLua.ConvMode.Exec )
   
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
   
   
   local meta = _lune.unwrap( loadFromLuaTxt( metaTxt ))
   return meta, luaTxt
end


function Front:loadFile( importModuleInfo, path, mod )
   local __func__ = '@lune.@base.@front.Front.loadFile'

   Log.log( Log.Level.Info, __func__, 528, function (  )
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


function Front:checkUptodateMeta( metaPath, addSearchPath )
   local __func__ = '@lune.@base.@front.Front.checkUptodateMeta'

   local metaObj = self:loadLua( metaPath )
   if  nil == metaObj then
      local _metaObj = metaObj
   
      Log.log( Log.Level.Warn, __func__, 566, function (  )
      
         return string.format( "load error -- %s", metaPath)
      end )
      
      return nil
   end
   
   local meta = metaObj
   if meta.__formatVersion ~= Ver.metaVersion then
      Log.log( Log.Level.Warn, __func__, 571, function (  )
      
         return string.format( "unmatch meta version -- %s", metaPath)
      end )
      
      return nil
   end
   
   if meta.__hasTest then
      
      if meta.__enableTest ~= self.option.testing then
         Log.log( Log.Level.Warn, __func__, 577, function (  )
         
            return string.format( "unmatch test setting -- %s", metaPath)
         end )
         
         return nil
      end
      
   end
   
   
   for moduleFullName, _5899 in pairs( meta.__dependModuleMap ) do
      do
         local lnsPath = self:searchModule( moduleFullName )
         if lnsPath ~= nil then
            do
               local moduleLuaPath = self:searchLuaFile( moduleFullName, addSearchPath )
               if moduleLuaPath ~= nil then
                  if not Util.getReadyCode( lnsPath, metaPath ) then
                     
                     Log.log( Log.Level.Warn, __func__, 588, function (  )
                     
                        return string.format( "not ready -- %s, %s", lnsPath, metaPath)
                     end )
                     
                     return nil
                  end
                  
                  local moduleMetaPath = moduleLuaPath:gsub( "%.lua$", ".meta" )
                  if Depend.existFile( moduleMetaPath ) and not Util.getReadyCode( moduleMetaPath, metaPath ) then
                     Log.log( Log.Level.Warn, __func__, 596, function (  )
                     
                        return string.format( "not ready -- %s, %s", moduleMetaPath, metaPath)
                     end )
                     
                     return nil
                  end
                  
               else
                  Log.log( Log.Level.Warn, __func__, 601, function (  )
                  
                     return string.format( "not found .lua file for -- %s", moduleFullName)
                  end )
                  
                  return nil
               end
            end
            
         else
            Log.log( Log.Level.Warn, __func__, 606, function (  )
            
               return string.format( "not found .lns file -- %s", moduleFullName)
            end )
            
            return nil
         end
      end
      
   end
   
   return meta
end


function Front:loadModule( mod )

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
               local lnsPath = self:searchModule( mod )
               if lnsPath ~= nil then
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  
                  do
                     local dir = self.option.outputDir
                     if dir ~= nil then
                        luaPath = self:searchLuaFile( mod, dir )
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
                                    local meta = self:checkUptodateMeta( metaPath, self.option.outputDir )
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


function Front:loadMeta( importModuleInfo, mod )
   local __func__ = '@lune.@base.@front.Front.loadMeta'

   if self.loadedMetaMap[mod] == nil then
      do
         local _exp = self:getLoadInfo( mod )
         if _exp ~= nil then
            self.loadedMetaMap[mod] = _exp.meta
         else
            do
               local lnsPath = self:searchModule( mod )
               if lnsPath ~= nil then
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  
                  do
                     local dir = self.option.outputDir
                     if dir ~= nil then
                        luaPath = self:searchLuaFile( mod, dir )
                     end
                  end
                  
                  
                  local meta = nil
                  if luaPath ~= nil then
                     if Util.getReadyCode( lnsPath, luaPath ) then
                        local metaPath = string.gsub( luaPath, "%.lua$", ".meta" )
                        if Util.getReadyCode( lnsPath, metaPath ) then
                           meta = self:checkUptodateMeta( metaPath, self.option.outputDir )
                           if meta ~= nil then
                              self.loadedMetaMap[mod] = meta
                           end
                           
                        else
                         
                           Log.log( Log.Level.Warn, __func__, 699, function (  )
                           
                              return string.format( "%s not ready meta %s, %s", mod, lnsPath, metaPath)
                           end )
                           
                        end
                        
                     else
                      
                        Log.log( Log.Level.Warn, __func__, 703, function (  )
                        
                           return string.format( "%s not ready lua %s, %s", mod, lnsPath, luaPath)
                        end )
                        
                     end
                     
                  else
                     Log.log( Log.Level.Warn, __func__, 707, function (  )
                     
                        return string.format( "%s not found lua in %s", mod, tostring( self.option.outputDir))
                     end )
                     
                  end
                  
                  if meta == nil then
                     local metawork, luaTxt = self:loadFileToLuaCode( importModuleInfo, lnsPath, mod )
                     self.loadedMetaMap[mod] = metawork
                     self.convertedMap[mod] = luaTxt
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   return self.loadedMetaMap[mod]
end


function Front:dumpTokenize(  )

   local parser = self:createPaser(  )
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


function Front:dumpAst(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   Depend.profile( self.option.validProf, function (  )
   
      local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
      ast:get_node():processFilter( dumpNode.createFilter( ast:get_moduleTypeInfo(), ast:get_processInfo(), io.stdout ), dumpNode.Opt.new("", 0) )
   end, self.option.scriptPath .. ".profi" )
end


function Front:format(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
   ast:get_node():processFilter( Formatter.createFilter( ast:get_moduleTypeInfo(), io.stdout ), Formatter.Opt.new(ast:get_node()) )
end


function Front:checkDiag(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   Util.setErrorCode( 0 )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Diag )
end


function Front:complete(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), self.option.analyzeModule, TransUnit.AnalyzeMode.Complete, self.option.analyzePos )
end


function Front:inquire(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), self.option.analyzeModule, TransUnit.AnalyzeMode.Inquire, self.option.analyzePos )
end


function Front:createGlue(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
   local filter = glueFilter.createFilter( self.option.outputDir )
   ast:get_node():processFilter( filter, 0 )
end



function Front:convertLuaToStreamFromScript( checkUptodate, convMode, path, mod, byteCompile, stripDebugInfo, openOStream, closeOStream )

   local function outputDependInfo( stream, metaInfo )
   
      if stream ~= nil then
         if metaInfo ~= nil then
            local dependInfo = OutputDepend.DependInfo.new(mod)
            for dependMod, _6063 in pairs( metaInfo.__dependModuleMap ) do
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
   
   local retAst = nil
   
   local moduleId, uptodate
   
   if checkUptodate then
      moduleId, uptodate = self:getModuleIdAndCheckUptodate( path, mod )
   else
    
      moduleId, uptodate = frontInterface.ModuleId.tempId, _lune.newAlge( ModuleUptodate.NeedUpdate)
   end
   
   
   local stream, metaStream, dependsStream = openOStream( uptodate )
   
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
   
   
   if closeOStream ~= nil then
      closeOStream( stream, metaStream, dependsStream )
   end
   
   
   return retAst
end


function Front:convertToLua(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   local convMode = convLua.ConvMode.Convert
   if self.option.mode == Option.ModeKind.LuaMeta then
      convMode = convLua.ConvMode.ConvMeta
   end
   
   
   local ast = self:convertLuaToStreamFromScript( false, convMode, self.option.scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( mode )
   
      return io.stdout, io.stdout, self.option:openDepend(  )
   end, function ( stream, metaStream, dependStream )
   
      if dependStream ~= nil then
         dependStream:close(  )
      end
      
   end )
   
   if ast ~= nil then
      do
         local _switchExp = self.option.convTo
         if _switchExp == Types.Conv.Go then
            local conv = convGo.createFilter( self.option.testing, "stdout", io.stdout, ast )
            ast:get_node():processFilter( conv, convGo.Opt.new(ast:get_node()) )
         end
      end
      
   end
   
end

function Front:saveToGo( ast )

   local path = self.option.scriptPath:gsub( "%.lns$", ".go" )
   
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
   
   
   local conv = convGo.createFilter( self.option.testing, path, file, ast )
   ast:get_node():processFilter( conv, convGo.Opt.new(ast:get_node()) )
   
   file:close(  )
end


function Front:saveToC( ast )

   local cPath = self.option.scriptPath:gsub( "%.lns$", ".c" )
   
   local file = io.open( cPath, "w" )
   if  nil == file then
      local _file = file
   
      return 
   end
   
   
   local hPath = self.option.scriptPath:gsub( "%.lns$", ".h" )
   local hFile = io.open( hPath, "w" )
   if  nil == hFile then
      local _hFile = hFile
   
      return 
   end
   
   
   local conv = convCC.createFilter( self.option.testing, self.option.mode == Option.ModeKind.Builtin, cPath, file, hFile, ast )
   ast:get_node():processFilter( conv, convCC.Opt.new(ast:get_node()) )
   
   file:close(  )
   hFile:close(  )
end


function Front:outputBuiltin(  )

   local mod = scriptPath2Module( "lns_builtin" )
   
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), Parser.DummyParser.new(), mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, TransUnit.AnalyzeMode.Compile )
   
   self:saveToC( ast )
end


function Front:saveToLua(  )

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
               Log.log( Log.Level.Debug, __func__, 1087, function (  )
               
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
   
   local mod = scriptPath2Module( self.option.scriptPath )
   Depend.profile( self.option.validProf, function (  )
   
      local luaPath = self.option.scriptPath:gsub( "%.lns$", ".lua" )
      local metaPath = self.option.scriptPath:gsub( "%.lns$", ".meta" )
      if self.option.outputDir then
         local filename = mod:gsub( "%.", "/" )
         luaPath = string.format( "%s/%s.lua", tostring( self.option.outputDir), filename)
         metaPath = string.format( "%s/%s.meta", tostring( self.option.outputDir), filename)
      end
      
      
      if luaPath == self.option.scriptPath then
         Util.errorLog( string.format( "%s is illegal filename.", luaPath) )
      else
       
         local convMode = convLua.ConvMode.Convert
         if self.option.mode == Option.ModeKind.SaveMeta then
            convMode = convLua.ConvMode.ConvMeta
         end
         
         local metaFileObj = nil
         local tempMetaPath = metaPath .. ".tmp"
         
         ast = self:convertLuaToStreamFromScript( true, convMode, self.option.scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( mode )
         
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
               
               
               local metaStream = stream
               if self.option.mode == Option.ModeKind.SaveMeta then
                  do
                     local _exp = io.open( tempMetaPath, "w+" )
                     if _exp ~= nil then
                        metaFileObj = _exp
                        metaStream = _exp
                     else
                        error( string.format( "write open error -- %s", metaPath) )
                     end
                  end
                  
               end
               
               return stream, metaStream
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
            
            
            return stream, metaStream, self.option:openDepend(  )
         end, function ( stream, metaStream, dependStream )
         
            if stream ~= nil then
               stream:close(  )
            end
            
            if dependStream ~= nil then
               dependStream:close(  )
            end
            
            
            if metaFileObj ~= nil then
               metaFileObj:flush(  )
               metaFileObj:seek( "set", 0 )
               local newMetaTxt = metaFileObj:read( "*a" )
               if  nil == newMetaTxt then
                  local _newMetaTxt = newMetaTxt
               
                  Util.err( string.format( "faled to read meta. -- %s.", tempMetaPath) )
               end
               
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
      
   end, self.option.scriptPath .. ".profi" )
   
   if updateFlag then
      self.option.scriptPath:gsub( "%.lns$", ".lua" )
   end
   
   
   if ast ~= nil then
      do
         local _switchExp = self.option.convTo
         if _switchExp == Types.Conv.C then
            self:saveToC( ast )
         elseif _switchExp == Types.Conv.Go then
            self:saveToGo( ast )
         end
      end
      
   end
   
   
   return updateFlag
end


function Front:outputBootC(  )

   local stream
   
   do
      local path = self.option.bootPath
      if path ~= nil then
         local cPath = path:gsub( "%.lns$", ".c" )
         local file = io.open( cPath, "w" )
         if  nil == file then
            local _file = file
         
            print( string.format( "failed to open file -- %s", cPath) )
            return 
         end
         
         stream = file
      else
         stream = io.stdout
      end
   end
   
   local initModule = scriptPath2Module( self.option.scriptPath )
   convCC.outputBootcode( stream, initModule )
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

function Front:exec(  )
   local __func__ = '@lune.@base.@front.Front.exec'

   Log.log( Log.Level.Trace, __func__, 1301, function (  )
   
      return Option.ModeKind:_getTxt( self.option.mode)
      
   end )
   
   
   do
      local _switchExp = self.option.mode
      if _switchExp == Option.ModeKind.Token then
         self:dumpTokenize(  )
      elseif _switchExp == Option.ModeKind.Ast then
         self:dumpAst(  )
      elseif _switchExp == Option.ModeKind.Format then
         self:format(  )
      elseif _switchExp == Option.ModeKind.Diag then
         self:checkDiag(  )
      elseif _switchExp == Option.ModeKind.Complete then
         self:complete(  )
      elseif _switchExp == Option.ModeKind.Inquire then
         self:inquire(  )
      elseif _switchExp == Option.ModeKind.Glue then
         self:createGlue(  )
      elseif _switchExp == Option.ModeKind.Lua or _switchExp == Option.ModeKind.LuaMeta then
         self:convertToLua(  )
      elseif _switchExp == Option.ModeKind.Save or _switchExp == Option.ModeKind.SaveMeta then
         self:saveToLua(  )
      elseif _switchExp == Option.ModeKind.Exec then
         local modObj = self:loadModule( scriptPath2Module( self.option.scriptPath ) )
         
         if self.option.testing then
            Testing.run( scriptPath2Module( self.option.scriptPath ) )
            
            Testing.outputAllResult( io.stdout )
         end
         
      elseif _switchExp == Option.ModeKind.BootC then
         self:outputBootC(  )
      elseif _switchExp == Option.ModeKind.Builtin then
         self:outputBuiltin(  )
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

local function setFront(  )

   local option = Option.createDefaultOption( "dummy.lns" )
   Front.new(option)
end
_moduleObj.setFront = setFront



return _moduleObj
