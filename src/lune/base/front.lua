--lune/base/front.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@front'
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

function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

function _lune.nilacc( val, fieldName, access, ... )
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
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
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
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
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



local Types = _lune.loadModule( 'lune.base.Types' )

local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )

local Tokenizer = _lune.loadModule( 'lune.base.Tokenizer' )

local convLua = _lune.loadModule( 'lune.base.convLua' )

local convGo = _lune.loadModule( 'lune.base.convGo' )

local convPython = _lune.loadModule( 'lune.base.convPython' )

local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )

local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )

local Util = _lune.loadModule( 'lune.base.Util' )

local Option = _lune.loadModule( 'lune.base.Option' )

local dumpNode = _lune.loadModule( 'lune.base.dumpNode' )

local glueFilter = _lune.loadModule( 'lune.base.glueFilter' )

local Depend = _lune.loadModule( 'lune.base.Depend' )

local DependLuaOnLns = _lune.loadModule( 'lune.base.DependLuaOnLns' )

local OutputDepend = _lune.loadModule( 'lune.base.OutputDepend' )

local Ver = _lune.loadModule( 'lune.base.Ver' )

local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )

local Log = _lune.loadModule( 'lune.base.Log' )

local Formatter = _lune.loadModule( 'lune.base.Formatter' )

local Testing = _lune.loadModule( 'lune.base.Testing' )

local GoMod = _lune.loadModule( 'lune.base.GoMod' )

local Meta = _lune.loadModule( 'lune.base.Meta' )

local Nodes = _lune.loadModule( 'lune.base.Nodes' )

local Builtin = _lune.loadModule( 'lune.base.Builtin' )

local NodeIndexer = _lune.loadModule( 'lune.base.NodeIndexer' )

local Converter = _lune.loadModule( 'lune.base.Converter' )





Depend.setup( function ( ver )

   LuaVer.setCurVer( ver )
end
 )


local forceUpdateMeta = true


local LoadInfo = {}
function LoadInfo._setmeta( obj )
  setmetatable( obj, { __index = LoadInfo  } )
end
function LoadInfo._new( mod, meta )
   local obj = {}
   LoadInfo._setmeta( obj )
   if obj.__init then
      obj:__init( mod, meta )
   end
   return obj
end
function LoadInfo:__init( mod, meta )

   self.mod = mod
   self.meta = meta
end




local UptodateInfo = {}
UptodateInfo._name2Val = {}
function UptodateInfo:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "UptodateInfo.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function UptodateInfo._from( val )
   return _lune._AlgeFrom( UptodateInfo, val )
end

UptodateInfo.Update = { "Update", {{},{}}}
UptodateInfo._name2Val["Update"] = UptodateInfo.Update
UptodateInfo.Uptodate = { "Uptodate", {{}}}
UptodateInfo._name2Val["Uptodate"] = UptodateInfo.Uptodate


local ModuleMgr = {}
function ModuleMgr._new(  )
   local obj = {}
   ModuleMgr._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function ModuleMgr:__init() 
   self.mod2info = Util.OrderdMap._new()
   self.loadedMetaMap = {}
end
function ModuleMgr:get( mod )

   return self.mod2info:get_map()[mod]
end
function ModuleMgr:getAst( mod )

   local info = self:get( mod )
   if  nil == info then
      local _info = info
   
      return nil
   end
   
   do
      local _matchExp = info
      if _matchExp[1] == UptodateInfo.Update[1] then
         local _ = _matchExp[2][1]
         local ast = _matchExp[2][2]
      
         return ast
      elseif _matchExp[1] == UptodateInfo.Uptodate[1] then
         local _ = _matchExp[2][1]
      
         return nil
      end
   end
end
function ModuleMgr:getModList(  )

   return self.mod2info:get_keyList()
end
function ModuleMgr:add( ast, moduleInfo )

   local mod = moduleInfo:get_fullName()
   self.mod2info:add( mod, _lune.newAlge( UptodateInfo.Update, {moduleInfo,ast}), true )
end
function ModuleMgr:addMeta( mod, meta )
   local __func__ = '@lune.@base.@front.ModuleMgr.addMeta'

   if not self:get( mod ) then
      Log.log( Log.Level.Log, __func__, 140, function (  )
      
         return mod
      end
       )
      
      
      self.mod2info:add( mod, _lune.newAlge( UptodateInfo.Uptodate, {meta}), false )
   end
   self.loadedMetaMap[mod] = meta
end
function ModuleMgr:getMeta( mod )

   return self.loadedMetaMap[mod]
end
function ModuleMgr._setmeta( obj )
  setmetatable( obj, { __index = ModuleMgr  } )
end


local Front = {}
setmetatable( Front, { ifList = {frontInterface.frontInterface,} } )
_moduleObj.Front = Front
function Front:regConvertedMap( mod, luaTxt, meta )

   self.moduleMgr:addMeta( mod, meta )
   self.convertedMap[mod] = luaTxt
end
function Front._new( option, bindModuleList )
   local obj = {}
   Front._setmeta( obj )
   if obj.__init then obj:__init( option, bindModuleList ); end
   return obj
end
function Front:__init(option, bindModuleList) 
   self.mod2loader = {}
   self.mod2astCreate = {}
   self.loadCount = 0
   self.targetSet = {}
   self.bindModuleSet = {}
   if bindModuleList ~= nil then
      for __index, mod in ipairs( bindModuleList ) do
         self.bindModuleSet[mod]= true
      end
   end
   self.moduleMgr = ModuleMgr._new()
   local gomodDir
   
   if option.scriptPath ~= "" then
      gomodDir = Util.parentPath( option.scriptPath )
   else
    
      gomodDir = ""
   end
   self.gomodMap = GoMod.getGoMap( gomodDir )
   DependLuaOnLns.addGoModPath( self.gomodMap:getModPathList(  ) )
   self.option = option
   self.loadedMap = {}
   self.loadedMapTest = {}
   self.convertedMap = {}
   do
      local builtin = Builtin.Builtin._new(self.option.targetLuaVer, option.transCtrlInfo)
      self.builtinFunc = builtin:registBuiltInScope(  )
   end
   local loadedMap = {}
   do
      for mod, modval in pairs( Depend.getLoadedMod(  ) ) do
         if mod == "lune.base.Testing" then
            loadedMap[mod] = modval
         end
         if option.testing and modval['__enableTest'] or not option.testing and modval['__enableTest'] then
            loadedMap[mod] = modval
         end
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
function Front._setmeta( obj )
  setmetatable( obj, { __index = Front  } )
end


function Front:error( message )

   Util.errorLog( message )
   Util.printStackTrace(  )
   os.exit( 1 )
end


function Front:loadLua( path )

   local ret
   
   do
      local chunk, err = loadfile( path )
      if chunk ~= nil then
         ret = _lune.unwrap( chunk(  ))
      else
         Util.errorLog( _lune.unwrapDefault( err, string.format( "load error -- %s.", path)) )
         ret = nil
      end
   end
   return ret
end


local function createPaser( baseDir, path, mod, stdinFile )

   return Tokenizer.StreamTokenizer.create( _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,path,mod,nil}), false, stdinFile, nil )
end


function Front:scriptPath2Module( path, baseDir )

   local mod = Util.scriptPath2ModuleFromProjDir( path, baseDir )
   return mod, baseDir
end


function Front:createPaser( scriptPath, baseDir )

   local mod = self:scriptPath2Module( scriptPath, baseDir )
   return createPaser( baseDir, scriptPath, mod, self.option:get_stdinFile() )
end


function Front:createAstSub( importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos )

   do
      local _exp = self.moduleMgr:get( mod )
      if _exp ~= nil then
         do
            local _matchExp = _exp
            if _matchExp[1] == UptodateInfo.Update[1] then
               local _ = _matchExp[2][1]
               local ast = _matchExp[2][2]
            
               return _lune.newAlge( Converter.CreateAstResult.Ast, {ast})
            elseif _matchExp[1] == UptodateInfo.Uptodate[1] then
               local _ = _matchExp[2][1]
            
               Util.err( string.format( "can't load the multiple module -- %s", mod) )
            end
         end
      end
   end
   do
      local creater = self.mod2astCreate[mod]
      if creater ~= nil then
         local ast = creater:getAst(  )
         return _lune.newAlge( Converter.CreateAstResult.Ast, {ast})
      end
   end
   local astCreater = Converter.AstCreater._new(frontInterface.FrontAccessor._new(self), importModuleInfo, tokenizerSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, self.builtinFunc, self.option)
   self.mod2astCreate[mod] = astCreater
   return _lune.newAlge( Converter.CreateAstResult.Creater, {astCreater})
end


function Front:applyAstResult( result )
   local __func__ = '@lune.@base.@front.Front.applyAstResult'

   do
      local _matchExp = result
      if _matchExp[1] == Converter.CreateAstResult.Ast[1] then
         local ast = _matchExp[2][1]
      
         return ast
      elseif _matchExp[1] == Converter.CreateAstResult.Creater[1] then
         local astCreater = _matchExp[2][1]
      
         local ast, moduleInfo, meta = astCreater:getAst(  )
         Log.log( Log.Level.Log, __func__, 350, function (  )
         
            return string.format( "applyAstResult -- %s", ast:get_exportInfo():get_fullName())
         end
          )
         
         
         self.moduleMgr:add( ast, moduleInfo )
         self.moduleMgr:addMeta( ast:get_exportInfo():get_fullName(), meta )
         self.mod2loader[ast:get_exportInfo():get_fullName()] = nil
         self.mod2astCreate[ast:get_exportInfo():get_fullName()] = nil
         return ast
      end
   end
end


function Front:createAst( importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos )

   local ast = self:applyAstResult( self:createAstSub( importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos ) )
   if self.option.dumpDebugAst then
      Util.println( string.format( "=== ast:%s ==========", mod) )
      ast:get_node():processFilter( dumpNode.createFilter( ast:get_exportInfo():get_moduleTypeInfo(), ast:get_exportInfo():get_processInfo(), io.stdout ), dumpNode.Opt._new("", 0) )
   end
   return ast
end


local function ast2Lua( ast, streamName, stream, metaStream, convMode, inMacro, option )

   local conv = Converter.ast2LuaMain( ast, streamName, stream, metaStream, convMode, inMacro, option )
   conv:outputLuaAndMeta( _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )) )
end


local function loadFromChunk( chunk, err )

   if err ~= nil then
      Util.errorLog( err )
   end
   local ret
   
   do
      if chunk ~= nil then
         do
            local work = chunk(  )
            if work ~= nil then
               ret = work
            else
               ret = nil
            end
         end
      else
         Util.err( "failed to error" )
      end
   end
   return ret
end


local function loadFromLuaTxt( txt )

   local chunk, err
   
   do
      chunk, err = _lune.loadstring52( txt )
   end
   return loadFromChunk( chunk, err )
end
_moduleObj.loadFromLuaTxt = loadFromLuaTxt


function Front:convertFromAst( ast, streamName, mode )

   local stream = Util.memStream._new()
   local metaStream = Util.memStream._new()
   ast2Lua( ast, streamName, stream, metaStream, mode, false, self.option )
   return metaStream:get_txt(), stream:get_txt()
end

function Front:loadFromLnsTxt( importModuleInfo, baseDir, name, txt )

   local _
   local transUnit = TransUnit.TransUnitCtrl._new(frontInterface.FrontAccessor._new(self), frontInterface.ModuleId.tempId, importModuleInfo, convLua.MacroEvalImp._new(self.builtinFunc), false, nil, nil, nil, self.option.targetLuaVer, self.option.transCtrlInfo, self.builtinFunc)
   local ast = transUnit:createAST( _lune.newAlge( Types.TokenizerSrc.LnsCode, {txt,name,nil}), false, baseDir, self.option:get_stdinFile(), false, string.format( "$load%d", self.loadCount), nil )
   self.loadCount = self.loadCount + 1
   local _1, luaTxt = self:convertFromAst( ast, name, convLua.ConvMode.ConvMeta )
   return _lune.unwrap( loadFromLuaTxt( luaTxt ))
end


local DependMetaInfo = {}
setmetatable( DependMetaInfo, { ifList = {Mapping,} } )
function DependMetaInfo._setmeta( obj )
  setmetatable( obj, { __index = DependMetaInfo  } )
end
function DependMetaInfo._new( use, buildId )
   local obj = {}
   DependMetaInfo._setmeta( obj )
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
     DependMetaInfo._setmeta( obj )
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
function MetaForBuildId._setmeta( obj )
  setmetatable( obj, { __index = MetaForBuildId  } )
end
function MetaForBuildId._new( __buildId, __dependModuleMap, __subModuleMap, __enableTest )
   local obj = {}
   MetaForBuildId._setmeta( obj )
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
     MetaForBuildId._setmeta( obj )
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
      local fileObj = Util.openRd( metaPath )
      if fileObj ~= nil then
         local luaCode = fileObj:read( "*a" )
         fileObj:close(  )
         if luaCode ~= nil then
            local meta
            
            do
               meta = MetaForBuildId._fromStem( (loadFromLuaTxt( luaCode ) ) )
            end
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


function Front:searchModuleFile( mod, suffix, baseDir, outputDir )
   local __func__ = '@lune.@base.@front.Front.searchModuleFile'

   do
      local _matchExp = self.gomodMap:convLocalModulePath( mod, suffix, baseDir )
      if _matchExp[1] == GoMod.GoModResult.NotGo[1] then
      
      elseif _matchExp[1] == GoMod.GoModResult.NotFound[1] then
      
         return nil
      elseif _matchExp[1] == GoMod.GoModResult.Found[1] then
         local info = _matchExp[2][1]
      
         return info:get_path()
      end
   end
   local lnsSearchPath = package.path
   if outputDir ~= nil then
      lnsSearchPath = string.format( "%s/?%s;%s", outputDir, suffix, package.path )
   end
   if baseDir ~= nil then
      lnsSearchPath = string.format( "%s/?%s;%s", baseDir, suffix, package.path )
   end
   do
      local projDir = self.option:get_projDir()
      if projDir ~= nil then
         lnsSearchPath = string.format( "%s;%s", Util.pathJoin( projDir, string.format( "?%s", suffix) ), package.path)
      end
   end
   lnsSearchPath = lnsSearchPath:gsub( "%.lua$", suffix )
   lnsSearchPath = lnsSearchPath:gsub( "%.lua;", suffix .. ";" )
   local foundPath = Util.searchPath( mod, lnsSearchPath )
   if  nil == foundPath then
      local _foundPath = foundPath
   
      local latestProjRoot = self.gomodMap:getLatestProjRoot(  )
      if  nil == latestProjRoot then
         local _latestProjRoot = latestProjRoot
      
         return nil
      end
      
      local latestProjSearchPath = Util.pathJoin( latestProjRoot, "?" .. suffix )
      do
         local _exp = Util.searchPath( mod, latestProjSearchPath )
         if _exp ~= nil then
            foundPath = _exp
         else
            Log.log( Log.Level.Err, __func__, 539, function (  )
            
               return string.format( "not found '%s' at %s", mod, latestProjSearchPath)
            end
             )
            
            
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
         local modMetaPath = self:searchModuleFile( depMod, ".meta", nil, self.option.outputDir )
         if  nil == modMetaPath then
            local _modMetaPath = modMetaPath
         
            Log.log( Log.Level.Debug, __func__, 636, function (  )
            
               return "NeedUpdate"
            end
             )
            
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         local time = Depend.getFileLastModifiedTime( modMetaPath )
         if  nil == time then
            local _time = time
         
            Log.log( Log.Level.Debug, __func__, 641, function (  )
            
               return "NeedUpdate"
            end
             )
            
            
            return _lune.newAlge( ModuleUptodate.NeedUpdate)
         end
         
         if time > metaTime then
            local dependMeta = MetaForBuildId.LoadFromMeta( modMetaPath )
            if  nil == dependMeta then
               local _dependMeta = dependMeta
            
               Log.log( Log.Level.Debug, __func__, 649, function (  )
               
                  return "NeedUpdate"
               end
                )
               
               
               return _lune.newAlge( ModuleUptodate.NeedUpdate)
            end
            
            local orgMetaModuleId = frontInterface.ModuleId.createIdFromTxt( dependItem.buildId )
            local metaModuleId = dependMeta:createModuleId(  )
            if metaModuleId:get_buildCount() ~= 0 and metaModuleId:get_buildCount() ~= orgMetaModuleId:get_buildCount() then
               Log.log( Log.Level.Debug, __func__, 659, function (  )
               
                  return string.format( "NeedUpdate: %s, %d, %d", modMetaPath, metaModuleId:get_buildCount(), orgMetaModuleId:get_buildCount())
               end
                )
               
               
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
      Log.log( Log.Level.Debug, __func__, 697, function (  )
      
         return "not found meta"
      end
       )
      
      
   end
   local moduleId = getModuleId( lnsPath, mod, self.option.outputDir, metaInfo )
   if moduleId == frontInterface.ModuleId.tempId then
      Util.err( string.format( "not found -- %s", lnsPath) )
   end
   return moduleId, uptodate
end

function Front:convertLns2LuaCode( importModuleInfo, analyzeMode, tokenizerSrc, baseDir, stream, streamName )

   local _
   local mod = self:scriptPath2Module( streamName, baseDir )
   local ast = self:createAst( importModuleInfo, tokenizerSrc, baseDir, mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, analyzeMode )
   local _1, luaTxt = self:convertFromAst( ast, streamName, convLua.ConvMode.ConvMeta )
   return luaTxt
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
   return convGo.Option._new(packageName, self:getGoAppName(  ), self.option.mainModule, self.option:get_addEnvArg(), self.option.convGoRunnerNum, self.option:get_sortGenerateCode())
end


function Front:createPythonOption( scriptPath )

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
   return convPython.Option._new(packageName, self:getGoAppName(  ), self.option.mainModule, false, self.option.convGoRunnerNum)
end


function Front:convertToLanguage( ast, stream, path )

   do
      local _switchExp = self.option.convTo
      if _switchExp == Types.Lang.Go then
         local conv = convGo.createFilter( self.option.testing, path, stream, ast, self:createGoOption( path ) )
         ast:get_node():processFilter( conv, convGo.Opt._new(ast:get_node()) )
      elseif _switchExp == Types.Lang.Python then
         local conv = convPython.createFilter( self.option.testing, path, stream, ast, self:createPythonOption( path ) )
         ast:get_node():processFilter( conv, convPython.Opt._new(ast:get_node()) )
      end
   end
end

function Front:loadTokenizerToLuaCode( importModuleInfo, tokenizerSrc, path, mod, baseDir )
   local __func__ = '@lune.@base.@front.Front.loadTokenizerToLuaCode'

   local moduleId = getModuleId( path, mod )
   local ast = self:createAst( importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile, nil )
   local metaTxt, luaTxt = self:convertFromAst( ast, path, convLua.ConvMode.ConvMeta )
   Log.log( Log.Level.Trace, __func__, 821, function (  )
   
      return string.format( "Meta = %s", metaTxt)
   end
    )
   
   
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
            saveTxt = Converter.byteCompileFromLuaTxt( saveTxt, stripDebugInfo )
         end
         if not forceUpdateMeta and checkUpdate then
            do
               local fileObj = Util.openRd( newpath )
               if fileObj ~= nil then
                  local oldTxt = fileObj:read( "*a" )
                  fileObj:close(  )
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
      if self.option.mode == Option.ModeKind.SaveMeta then
         saveFile( ".meta", metaTxt, self.option.byteCompile, true, true )
      end
      do
         local memStream = Util.memStream._new()
         self:convertToLanguage( ast, memStream, path )
         do
            local convTo = self.option.convTo
            if convTo ~= nil then
               local suffix
               
               do
                  local _switchExp = convTo
                  if _switchExp == Types.Lang.Same or _switchExp == Types.Lang.Lua then
                     suffix = nil
                  elseif _switchExp == Types.Lang.Go then
                     suffix = ".go"
                  elseif _switchExp == Types.Lang.Python then
                     suffix = ".py"
                  elseif _switchExp == Types.Lang.C then
                     suffix = ".c"
                  end
               end
               if suffix ~= nil then
                  saveFile( suffix, memStream:get_txt(), false, false, false )
               end
            end
         end
      end
   end
   return _lune.unwrap( self.moduleMgr:getMeta( mod )), luaTxt
end

function Front:loadFile( importModuleInfo, baseDir, path, mod )
   local __func__ = '@lune.@base.@front.Front.loadFile'

   Log.log( Log.Level.Info, __func__, 903, function (  )
      local __func__ = '@lune.@base.@front.Front.loadFile.<anonymous>'
   
      return string.format( "start %s:%s", __func__, mod)
   end
    )
   
   
   local meta, luaTxt = self:loadTokenizerToLuaCode( importModuleInfo, _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,path,mod,nil}), path, mod, baseDir )
   do
      local preLoadInfo = self.preloadedModMap[mod]
      if preLoadInfo ~= nil then
         return meta, preLoadInfo
      end
   end
   return meta, _lune.unwrap( loadFromLuaTxt( luaTxt ))
end


function Front:searchModule( mod, baseDir, addSearchPath )

   return self:searchModuleFile( mod, ".lns", baseDir, addSearchPath )
end


function Front:searchLuaFile( moduleFullName, baseDir, addSearchPath )

   return self:searchModuleFile( moduleFullName, ".lua", baseDir, addSearchPath )
end

function Front:checkUptodateMeta( lnsPath, metaPath, baseDir, addSearchPath )
   local __func__ = '@lune.@base.@front.Front.checkUptodateMeta'

   local metaObj = self:loadLua( metaPath )
   if  nil == metaObj then
      local _metaObj = metaObj
   
      Log.log( Log.Level.Warn, __func__, 938, function (  )
      
         return string.format( "load error -- %s", metaPath)
      end
       )
      
      
      return nil
   end
   
   local meta
   
   local valid = true
   do
      meta = metaObj
      if meta.__formatVersion ~= Ver.metaVersion then
         Log.log( Log.Level.Warn, __func__, 947, function (  )
         
            return string.format( "unmatch meta version -- %s", metaPath)
         end
          )
         
         
         valid = false
      end
      if valid and meta.__hasTest then
         if meta.__enableTest ~= self.option.testing then
            Log.log( Log.Level.Warn, __func__, 953, function (  )
            
               return string.format( "unmatch test setting -- %s", metaPath)
            end
             )
            
            
            valid = false
         end
      end
      if valid then
         for moduleFullName, _1 in pairs( meta.__dependModuleMap ) do
            do
               local moduleLnsPath = self:searchModule( moduleFullName, baseDir, addSearchPath )
               if moduleLnsPath ~= nil then
                  do
                     local moduleLuaPath = self:searchLuaFile( moduleFullName, baseDir, addSearchPath )
                     if moduleLuaPath ~= nil then
                        if not Util.getReadyCode( moduleLnsPath, metaPath ) then
                           Log.log( Log.Level.Warn, __func__, 967, function (  )
                           
                              return string.format( "not ready -- %s, %s", moduleLnsPath, metaPath)
                           end
                            )
                           
                           
                           valid = false
                           break
                        end
                        local moduleMetaPath = moduleLuaPath:gsub( "%.lua$", ".meta" )
                        if Depend.existFile( moduleMetaPath ) and not Util.getReadyCode( moduleMetaPath, metaPath ) then
                           Log.log( Log.Level.Warn, __func__, 976, function (  )
                           
                              return string.format( "not ready -- %s, %s", moduleMetaPath, metaPath)
                           end
                            )
                           
                           
                           valid = false
                           break
                        end
                     else
                        Log.log( Log.Level.Warn, __func__, 982, function (  )
                        
                           return string.format( "not found .lua file for -- %s", moduleFullName)
                        end
                         )
                        
                        
                        valid = false
                        break
                     end
                  end
               else
                  Log.log( Log.Level.Warn, __func__, 988, function (  )
                  
                     return string.format( "not found .lns file -- %s", moduleFullName)
                  end
                   )
                  
                  
                  valid = false
                  break
               end
            end
         end
      end
   end
   if not valid then
      return nil
   end
   return frontInterface.ModuleMeta._new(lnsPath, _lune.newAlge( frontInterface.MetaOrModule.MetaRaw, {meta}))
end

function Front:loadModuleWithBaseDir( orgMod, baseDir )
   local __func__ = '@lune.@base.@front.Front.loadModuleWithBaseDir'

   local _
   Log.log( Log.Level.Info, __func__, 1013, function (  )
   
      return string.format( "loadModule  -- %s:%s", orgMod, baseDir)
   end
    )
   
   
   local _1, _2, mod = self.gomodMap:getLuaModulePath( orgMod, baseDir )
   if not self:getLoadInfo( mod ) then
      do
         local luaTxt = self.convertedMap[mod]
         if luaTxt ~= nil then
            do
               local meta = self.moduleMgr:getMeta( mod )
               if meta ~= nil then
                  self:setLoadInfo( mod, LoadInfo._new(_lune.unwrap( loadFromLuaTxt( luaTxt )), meta) )
               else
                  Util.err( string.format( "nothing meta -- %s", mod) )
               end
            end
         else
            do
               local lnsPath = self:searchModule( orgMod, baseDir, nil )
               if lnsPath ~= nil then
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  do
                     local dir = self.option.outputDir
                     if dir ~= nil then
                        luaPath = self:searchLuaFile( orgMod, nil, dir )
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
                                    local meta = self:checkUptodateMeta( lnsPath, metaPath, baseDir, self.option.outputDir )
                                    if meta ~= nil then
                                       self:setLoadInfo( mod, LoadInfo._new(_exp, meta) )
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
                     local meta, workVal = self:loadFile( frontInterface.ImportModuleInfo._new(), baseDir, lnsPath, mod )
                     self:setLoadInfo( mod, LoadInfo._new(workVal, meta) )
                  end
               else
                  if _lune._Set_has(self.bindModuleSet, mod ) then
                     Log.log( Log.Level.Warn, __func__, 1070, function (  )
                     
                        return string.format( "load from the binding -- %s", mod)
                     end
                      )
                     
                     
                     local workMod
                     
                     do
                        workMod = require( mod )
                     end
                     local meta = frontInterface.ModuleMeta._new(mod:gsub( "%.", "/" ) .. ".lns", _lune.newAlge( frontInterface.MetaOrModule.MetaRaw, {_lune.unwrap( loadFromLuaTxt( "return {}" ))}))
                     self:setLoadInfo( mod, LoadInfo._new(workMod, meta) )
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
   Util.err( string.format( "load error, %s", mod) )
end

function Front:loadModule( mod )

   return self:loadModuleWithBaseDir( mod, nil )
end


function Front:getLuaModulePath( mod, baseDir )

   return self.gomodMap:getLuaModulePath( mod, baseDir )
end


function Front:loadMeta( importModuleInfo, mod, orgMod, baseDir, loader )
   local __func__ = '@lune.@base.@front.Front.loadMeta'

   do
      local creater = self.mod2astCreate[orgMod]
      if creater ~= nil then
         local exportInfo = creater:getExportInfo(  )
         if  nil == exportInfo then
            local _exportInfo = exportInfo
         
            return nil
         end
         
         local meta = frontInterface.ModuleMeta._new(exportInfo:get_streamName(), _lune.newAlge( frontInterface.MetaOrModule.Export, {exportInfo}))
         self.moduleMgr:addMeta( orgMod, meta )
      end
   end
   do
      local workLoader = self.mod2loader[orgMod]
      if workLoader ~= nil then
         local exportInfo = workLoader:getExportInfo(  )
         if  nil == exportInfo then
            local _exportInfo = exportInfo
         
            return nil
         end
         
         local meta = frontInterface.ModuleMeta._new(exportInfo:get_streamName(), _lune.newAlge( frontInterface.MetaOrModule.Export, {exportInfo}))
         self.moduleMgr:addMeta( orgMod, meta )
         return meta
      end
   end
   if not self.moduleMgr:getMeta( orgMod ) then
      self.mod2loader[orgMod] = loader
      do
         local _exp = self:getLoadInfo( orgMod )
         if _exp ~= nil then
            self.moduleMgr:addMeta( orgMod, _exp.meta )
         else
            do
               local lnsPath = self:searchModule( mod, baseDir, nil )
               if lnsPath ~= nil then
                  local meta = nil
                  local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                  if not _lune._Set_has(self.targetSet, orgMod ) then
                     luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
                     if not baseDir then
                        do
                           local dir = self.option.outputDir
                           if dir ~= nil then
                              luaPath = self:searchLuaFile( orgMod, nil, dir )
                           end
                        end
                     end
                     if luaPath ~= nil then
                        local forceFlag
                        
                        do
                           local _matchExp = self.option.transCtrlInfo.uptodateMode
                           if _matchExp[1] == Types.CheckingUptodateMode.ForceAll[1] then
                           
                              forceFlag = true
                           elseif _matchExp[1] == Types.CheckingUptodateMode.Force1[1] then
                              local forceMod = _matchExp[2][1]
                           
                              forceFlag = orgMod == forceMod
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
                                 meta = self:checkUptodateMeta( lnsPath, metaPath, baseDir, self.option.outputDir )
                              else
                               
                                 Log.log( Log.Level.Warn, __func__, 1185, function (  )
                                 
                                    return string.format( "%s not ready meta %s, %s", orgMod, lnsPath, metaPath)
                                 end
                                  )
                                 
                                 
                              end
                           else
                            
                              Log.log( Log.Level.Warn, __func__, 1189, function (  )
                              
                                 return string.format( "%s not ready lua %s, %s", orgMod, lnsPath, luaPath)
                              end
                               )
                              
                              
                           end
                        else
                         
                           Log.log( Log.Level.Warn, __func__, 1193, function (  )
                           
                              return string.format( "force analyze -- %s", orgMod)
                           end
                            )
                           
                           
                        end
                     else
                        Log.log( Log.Level.Warn, __func__, 1197, function (  )
                        
                           return string.format( "%s not found lua in %s", orgMod, self.option.outputDir)
                        end
                         )
                        
                        
                     end
                  end
                  if meta ~= nil then
                     self.moduleMgr:addMeta( orgMod, meta )
                     do
                        local _switchExp = self.option.mode
                        if _switchExp == Option.ModeKind.Exec or _switchExp == Option.ModeKind.Shebang then
                           if luaPath ~= nil then
                              do
                                 local luaCode = Util.readFile( luaPath )
                                 if luaCode ~= nil then
                                    self:regConvertedMap( orgMod, luaCode, meta )
                                 end
                              end
                           end
                        end
                     end
                  else
                     local metawork, luaTxt = self:loadTokenizerToLuaCode( importModuleInfo, _lune.newAlge( Types.TokenizerSrc.LnsPath, {nil,lnsPath,orgMod,nil}), lnsPath, orgMod, baseDir )
                     self:regConvertedMap( orgMod, luaTxt, metawork )
                  end
               else
                  do
                     local lnsCode = Depend.getBindLns( orgMod )
                     if lnsCode ~= nil then
                        local path = mod:gsub( "%.", "/" ) .. ".lns"
                        local meta, luaTxt = self:loadTokenizerToLuaCode( importModuleInfo, _lune.newAlge( Types.TokenizerSrc.LnsCode, {lnsCode,orgMod,nil}), path, orgMod, baseDir )
                        self:regConvertedMap( orgMod, luaTxt, meta )
                     end
                  end
               end
            end
         end
      end
   end
   do
      local meta = self.moduleMgr:getMeta( orgMod )
      if meta ~= nil then
         return meta
      end
   end
   return nil
end


function Front:dumpTokenize( scriptPath, baseDir )

   local tokenizer = self:createPaser( scriptPath, baseDir )
   while true do
      local token = tokenizer:getToken(  )
      if  nil == token then
         local _token = token
      
         break
      end
      
      Util.println( Types.TokenKind:_getTxt( token.kind)
      , token.consecutive, token.pos.lineNo, token.pos.column, token.txt )
   end
end

function Front:dumpAst( tokenizerSrc, mod, moduleId, baseDir )

   Depend.profile( self.option.validProf, function (  )
   
      local ast = self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
      ast:get_node():processFilter( dumpNode.createFilter( ast:get_exportInfo():get_moduleTypeInfo(), ast:get_exportInfo():get_processInfo(), io.stdout ), dumpNode.Opt._new("", 0) )
   end
   , mod .. ".profi" )
end

function Front:format( tokenizerSrc, mod, moduleId, baseDir )

   local ast = self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
   ast:get_node():processFilter( Formatter.createFilter( ast:get_exportInfo():get_moduleTypeInfo(), io.stdout, false ), Formatter.Opt._new(ast:get_node()) )
end


function Front:checkDiag( tokenizerSrc, mod, moduleId, baseDir )

   Util.setErrorCode( 0 )
   self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Diag )
end


function Front:complete( tokenizerSrc, mod, moduleId, baseDir )

   self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, self.option.analyzeModule, TransUnit.AnalyzeMode.Complete, self.option.analyzePos )
end


function Front:completeFromCode( lnsCode, mod, baseDir )

   self:complete( _lune.newAlge( Types.TokenizerSrc.LnsCode, {lnsCode,mod,nil}), mod, frontInterface.ModuleId.createId( 0.0, 0 ), baseDir )
end


function Front:inquire( tokenizerSrc, mod, moduleId, baseDir )

   self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, self.option.analyzeModule, TransUnit.AnalyzeMode.Inquire, self.option.analyzePos )
end


function Front:createGlue( tokenizerSrc, mod, moduleId, baseDir )

   local ast = self:createAst( frontInterface.ImportModuleInfo._new(), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
   local filter = glueFilter.createFilter( self.option.outputDir )
   ast:get_node():processFilter( filter, 0 )
end


local function outputDependInfo( stream, metaInfo, mod )

   if stream ~= nil then
      local dependInfo = OutputDepend.DependInfo._new(mod)
      for dependMod, _1 in pairs( metaInfo.__dependModuleMap ) do
         dependInfo:addImpotModule( dependMod )
      end
      for __index, subMod in ipairs( metaInfo.__subModuleMap ) do
         dependInfo:addSubMod( subMod )
      end
      dependInfo:output( stream )
   end
end


function Front:convertToLua( scriptPath, baseDir, convMode, streamLua, streamMeta )

   local mod = self:scriptPath2Module( scriptPath, baseDir )
   local moduleId = getModuleId( scriptPath, mod )
   local ast = self:createAst( frontInterface.ImportModuleInfo._new(), _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,scriptPath,mod,nil}), baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile, nil )
   local metaTxt, luaTxt = self:convertFromAst( ast, scriptPath, convMode )
   streamLua:write( luaTxt )
   streamMeta:write( metaTxt )
   self:convertToLanguage( ast, streamLua, scriptPath )
   return ast
end

function Front:saveToGo( scriptPath, astResult, mod )

   return Converter.GoConverter._new(scriptPath, astResult, mod, self.option, self:createGoOption( scriptPath ))
end

function Front:saveToPython( scriptPath, astResult, mod )

   return Converter.PythonConverter._new(scriptPath, astResult, mod, self.option, self:createPythonOption( scriptPath ))
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

   local mod, baseDir = self:scriptPath2Module( "lns_builtin", nil )
   local ast = self:createAst( frontInterface.ImportModuleInfo._new(), _lune.newAlge( Types.TokenizerSrc.LnsCode, {"",mod,nil}), baseDir, mod, frontInterface.ModuleId.createId( 0.0, 0 ), nil, TransUnit.AnalyzeMode.Compile )
   self:saveToC( scriptPath, ast )
end


local UpdateInfo = {}
function UpdateInfo._setmeta( obj )
  setmetatable( obj, { __index = UpdateInfo  } )
end
function UpdateInfo._new( scriptPath, dependsPath, moduleId, uptodate )
   local obj = {}
   UpdateInfo._setmeta( obj )
   if obj.__init then
      obj:__init( scriptPath, dependsPath, moduleId, uptodate )
   end
   return obj
end
function UpdateInfo:__init( scriptPath, dependsPath, moduleId, uptodate )

   self.scriptPath = scriptPath
   self.dependsPath = dependsPath
   self.moduleId = moduleId
   self.uptodate = uptodate
end
function UpdateInfo:get_scriptPath()
   return self.scriptPath
end
function UpdateInfo:get_dependsPath()
   return self.dependsPath
end
function UpdateInfo:get_moduleId()
   return self.moduleId
end
function UpdateInfo:get_uptodate()
   return self.uptodate
end

function Front:saveToLua( updateInfo, baseDir )

   local scriptPath = updateInfo:get_scriptPath()
   local dependsPath = updateInfo:get_dependsPath()
   local moduleId = updateInfo:get_moduleId()
   local uptodate = updateInfo:get_uptodate()
   local mod = self:scriptPath2Module( scriptPath, baseDir )
   local luaPath = scriptPath:gsub( "%.lns$", ".lua" )
   local metaPath = scriptPath:gsub( "%.lns$", ".meta" )
   if self.option.outputDir then
      local filename = mod:gsub( "%.", "/" )
      luaPath = string.format( "%s/%s.lua", self.option.outputDir, filename)
      metaPath = string.format( "%s/%s.meta", self.option.outputDir, filename)
   end
   local convMode = convLua.ConvMode.ConvMeta
   if luaPath == scriptPath then
      Util.errorLog( string.format( "%s is illegal filename.", luaPath) )
      return nil
   end
   do
      local _matchExp = uptodate
      if _matchExp[1] == ModuleUptodate.NeedUpdate[1] then
      
         local result = self:createAstSub( frontInterface.ImportModuleInfo._new(), _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,scriptPath,mod,nil}), baseDir, mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
         local luaConv
         
         if self.option.noLua then
            luaConv = nil
         else
          
            luaConv = Converter.LuaConverter._new(luaPath, metaPath, dependsPath, result, convMode, scriptPath, self.option.byteCompile, self.option.stripDebugInfo, self.option)
         end
         local goConv = nil
         local pyConv = nil
         do
            local _switchExp = self.option.convTo
            if _switchExp == Types.Lang.Python then
               pyConv = self:saveToPython( scriptPath, result, mod )
            elseif _switchExp == Types.Lang.Go then
               goConv = self:saveToGo( scriptPath, result, mod )
            end
         end
         return function (  )
         
            _lune.nilacc( luaConv, 'saveLua', 'callmtd'  )
            _lune.nilacc( goConv, 'saveGo', 'callmtd'  )
            _lune.nilacc( pyConv, 'savePython', 'callmtd'  )
         end
         
      elseif _matchExp[1] == ModuleUptodate.Uptodate[1] then
         local metaInfo = _matchExp[2][1]
      
         Util.errorLog( "uptodate -- " .. scriptPath )
         local dependsStream = self.option:openDepend( dependsPath )
         outputDependInfo( dependsStream, metaInfo, mod )
         Converter.closeStreams( nil, nil, dependsStream, metaPath, false )
         return nil
      elseif _matchExp[1] == ModuleUptodate.NeedTouch[1] then
         local metaCode = _matchExp[2][1]
         local metaInfo = _matchExp[2][2]
      
         Util.errorLog( "touch -- " .. scriptPath )
         local dependsStream = self.option:openDepend( dependsPath )
         local metaMemStream = Util.memStream._new()
         if self.option.mode == Option.ModeKind.SaveMeta then
            metaMemStream:write( metaCode )
         end
         outputDependInfo( dependsStream, metaInfo, mod )
         Converter.closeStreams( nil, metaMemStream, dependsStream, metaPath, self.option.mode == Option.ModeKind.SaveMeta )
         return nil
      end
   end
end


function Front:outputBootC( scriptPath )

end


function Front:convertLnsCode2LuaCodeWithOpt( lnsCode, path, baseDir )

   return self:convertLns2LuaCode( frontInterface.ImportModuleInfo._new(), TransUnit.AnalyzeMode.Compile, _lune.newAlge( Types.TokenizerSrc.LnsCode, {lnsCode,path,nil}), baseDir, Tokenizer.TxtStream._new(lnsCode), path )
end


local function convertLnsCode2LuaCode( lnsCode, path, baseDir )

   local option = Option.Option._new()
   option.scriptPath = path
   option.useLuneModule = Option.getRuntimeModule(  )
   option.useIpairs = true
   local front = Front._new(option)
   return front:convertLnsCode2LuaCodeWithOpt( lnsCode, path, baseDir )
end
_moduleObj.convertLnsCode2LuaCode = convertLnsCode2LuaCode


local function getBaseDir( path, userProjDir )

   local projDir
   
   if userProjDir ~= nil then
      projDir = userProjDir
   else
      if path:find( "^/" ) or path:find( "^%.%./" ) then
         local parentPath = Util.parentPath( path )
         do
            local workDir = Util.searchProjDir( parentPath )
            if workDir ~= nil then
               projDir = workDir
            else
               projDir = parentPath
            end
         end
      else
       
         return path, nil
      end
   end
   local scriptPath
   
   if #projDir == 1 then
      scriptPath = path:sub( 2 )
   else
    
      scriptPath = path:sub( #projDir + 2 )
   end
   return scriptPath, projDir
end


local function getParseSrcAndBaseDir( path, userProjDir )

   local workPath, projDir = getBaseDir( path, userProjDir )
   local mod = Util.scriptPath2ModuleFromProjDir( workPath, projDir )
   local moduleId = getModuleId( path, mod )
   return _lune.newAlge( Types.TokenizerSrc.LnsPath, {projDir,workPath,mod,nil}), mod, moduleId, projDir
end


local BuildMode = {}
BuildMode._name2Val = {}
_moduleObj.BuildMode = BuildMode
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
BuildMode.Output = { "Output", {{},{}}}
BuildMode._name2Val["Output"] = BuildMode.Output
BuildMode.Save = { "Save"}
BuildMode._name2Val["Save"] = BuildMode.Save


function Front:build( buildMode, astCallback )

   local function createUpdateInfo( scriptPath, dependsPath, baseDir )
   
      local mod = self:scriptPath2Module( scriptPath, baseDir )
      local moduleId, uptodate = self:getModuleIdAndCheckUptodate( scriptPath, mod )
      return UpdateInfo._new(scriptPath, dependsPath, moduleId, uptodate)
   end
   local function process( oneShot, updateInfo, baseDir )
   
      local mod = self:scriptPath2Module( updateInfo:get_scriptPath(), baseDir )
      do
         local _matchExp = buildMode
         if _matchExp[1] == BuildMode.Save[1] then
         
            return self:saveToLua( updateInfo, baseDir )
         elseif _matchExp[1] == BuildMode.Output[1] then
            local streamLua = _matchExp[2][1]
            local streamMeta = _matchExp[2][2]
         
            self:convertToLua( updateInfo:get_scriptPath(), baseDir, convLua.ConvMode.ConvMeta, streamLua, streamMeta )
         elseif _matchExp[1] == BuildMode.CreateAst[1] then
         
            if not self.mod2astCreate[mod] and not self.moduleMgr:getAst( mod ) then
               local result = self:createAstSub( frontInterface.ImportModuleInfo._new(), _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,updateInfo:get_scriptPath(),mod,nil}), baseDir, mod, updateInfo:get_moduleId(), nil, TransUnit.AnalyzeMode.Compile )
               return function (  )
               
                  self:applyAstResult( result )
               end
               
            end
         end
      end
      return nil
   end
   Depend.profile( self.option.validProf, function (  )
      local __func__ = '@lune.@base.@front.Front.build.<anonymous>'
   
      if self.option.scriptPath == "@-" then
         local baseDir = self.option:get_projDir()
         for __index, path in ipairs( self.option.batchList ) do
            self.targetSet[(self:scriptPath2Module( path, baseDir ) )]= true
         end
         local postProcessMap = {}
         for index, path in ipairs( self.option.batchList ) do
            local updateInfo = createUpdateInfo( path, (path:gsub( ".lns$", ".d" ) ), baseDir )
            Util.println( string.format( "%s: start...", updateInfo:get_scriptPath()) )
            do
               local _exp = process( false, updateInfo, nil )
               if _exp ~= nil then
                  if self.option:get_validPostBuild() then
                     postProcessMap[index] = _exp
                  else
                   
                     _exp(  )
                  end
               end
            end
         end
         do
            local __sorted = {}
            local __map = postProcessMap
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, index in ipairs( __sorted ) do
               local postProcess = __map[ index ]
               do
                  local prev = os.clock(  )
                  local path = self.option.batchList[index]
                  Util.println( string.format( "%s: waiting...", path) )
                  postProcess(  )
                  Util.println( string.format( "%s: done %g msec", path, (os.clock(  ) - prev ) * 1000) )
               end
            end
         end
      else
       
         local _
         local _1, baseDir = getBaseDir( self.option.scriptPath, self.option:get_projDir() )
         do
            local postProcess = process( true, createUpdateInfo( self.option.scriptPath, nil, baseDir ), baseDir )
            if postProcess ~= nil then
               postProcess(  )
            end
         end
      end
      if astCallback ~= nil then
         for __index, mod in ipairs( self.moduleMgr:getModList(  ) ) do
            do
               local _exp = self.moduleMgr:getAst( mod )
               if _exp ~= nil then
                  astCallback( _exp )
               else
                  Log.log( Log.Level.Err, __func__, 1700, function (  )
                  
                     return string.format( "not found AST -- %s", mod)
                  end
                   )
                  
                  
               end
            end
         end
      end
      
   end
   , self.option.scriptPath .. ".profi" )
end


local function buildWithBuildMode( option, buildMode, astCallback )

   local front
   
   do
      front = Front._new(option)
   end
   front:build( buildMode, astCallback )
end
_moduleObj.buildWithBuildMode = buildWithBuildMode


local function build( option, astCallback )

   buildWithBuildMode( option, _lune.newAlge( BuildMode.CreateAst), astCallback )
end
_moduleObj.build = build


function Front:setupPreloadWithImportedModules( asyncFlag )

   if asyncFlag then
      do
         do
            local subModPreLoad = [==[
return function( submod2Code, dumpCode )
   local preloadFunc = function( mod )
      code = submod2Code[ mod ]
      local loadFunc = loadstring or load -- lua5.1 and lua5.2
      local loaded, mess = loadFunc( code )
      if not loaded then
         error( mess )
      end
      return loaded()
   end
   for mod, code in pairs( submod2Code ) do
      if dumpCode then
         print( string.format( "mod: %s %s", mod, code ) )
      end
      package.preload[ mod ] = preloadFunc
   end
end
]==]
            local loaded = _lune.loadstring52( subModPreLoad )
            if  nil == loaded then
               local _loaded = loaded
            
               Util.err( "failed to subModPreLoad" )
            end
            
            local preloadFunc = loaded(  )
            if  nil == preloadFunc then
               local _preloadFunc = preloadFunc
            
               Util.err( "failed to preloadFunc" )
            end
            
            (preloadFunc )( self.convertedMap, Log.getLevel(  ) >= Log.Level.Debug )
         end
      end
   else
    
      do
         do
            local subModPreLoad = [==[
return function( submod2Code, dumpCode )
   local preloadFunc = function( mod )
      code = submod2Code[ mod ]
      local loadFunc = loadstring or load -- lua5.1 and lua5.2
      local loaded, mess = loadFunc( code )
      if not loaded then
         error( mess )
      end
      return loaded()
   end
   for mod, code in pairs( submod2Code ) do
      if dumpCode then
         print( string.format( "mod: %s %s", mod, code ) )
      end
      package.preload[ mod ] = preloadFunc
   end
end
]==]
            local loaded = _lune.loadstring52( subModPreLoad )
            if  nil == loaded then
               local _loaded = loaded
            
               Util.err( "failed to subModPreLoad" )
            end
            
            local preloadFunc = loaded(  )
            if  nil == preloadFunc then
               local _preloadFunc = preloadFunc
            
               Util.err( "failed to preloadFunc" )
            end
            
            (preloadFunc )( self.convertedMap, Log.getLevel(  ) >= Log.Level.Debug )
         end
      end
   end
   
   
end


function Front:executeLns( path, baseDir )
   local __func__ = '@lune.@base.@front.Front.executeLns'

   local _
   local mod = Util.scriptPath2ModuleFromProjDir( path, baseDir )
   local tokenizerSrc = _lune.newAlge( Types.TokenizerSrc.LnsPath, {baseDir,path,mod,nil})
   local _1, luaCode = self:loadTokenizerToLuaCode( frontInterface.ImportModuleInfo._new(), tokenizerSrc, path, mod, baseDir )
   Log.log( Log.Level.Debug, __func__, 1780, function (  )
   
      return "luacode: " .. luaCode
   end
    )
   
   
   self:setupPreloadWithImportedModules( false )
   return loadFromLuaTxt( luaCode ), mod
end


function Front:executeLnsAndTest( path, baseDir )

   local _
   local _1, mod = self:executeLns( path, baseDir )
   if self.option.testing then
      local code = [==[
local Testing = require( "lune.base.Testing" )
return function( path )
  Testing.run( path );
  Testing.outputAllResult( io.stdout );
end
]==]
      do
         local loaded, mess = _lune.loadstring52( code )
         if loaded ~= nil then
            do
               local testFunc = loaded(  )
               if testFunc ~= nil then
                  (testFunc )( mod )
               end
            end
         else
            Util.println( mess )
         end
      end
   end
end


function Front:exec(  )
   local __func__ = '@lune.@base.@front.Front.exec'

   Log.log( Log.Level.Trace, __func__, 1815, function (  )
   
      return Option.ModeKind:_getTxt( self.option.mode)
      
   end
    )
   
   
   do
      local _switchExp = self.option.mode
      if _switchExp == Option.ModeKind.Token then
         self:dumpTokenize( getBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Ast then
         self:dumpAst( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Format then
         self:format( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Diag then
         self:checkDiag( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Complete then
         self:complete( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Inquire then
         self:inquire( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Glue then
         self:createGlue( getParseSrcAndBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.Lua or _switchExp == Option.ModeKind.LuaMeta then
         local convMode
         
         if self.option.mode == Option.ModeKind.Lua then
            convMode = convLua.ConvMode.Convert
         else
          
            convMode = convLua.ConvMode.ConvMeta
         end
         local scriptPath, baseDir = getBaseDir( self.option.scriptPath, self.option:get_projDir() )
         self:convertToLua( scriptPath, baseDir, convMode, io.stdout, io.stdout )
      elseif _switchExp == Option.ModeKind.Save or _switchExp == Option.ModeKind.SaveMeta then
         self:build( _lune.newAlge( BuildMode.Save), nil )
      elseif _switchExp == Option.ModeKind.BuildAst then
         self:build( _lune.newAlge( BuildMode.CreateAst), function ( ast )
         
            Util.println( ast:get_streamName() )
         end
          )
      elseif _switchExp == Option.ModeKind.Shebang then
         Depend.setupShebang(  )
         do
            local modObj = self:executeLns( getBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
            if  nil == modObj then
               local _modObj = modObj
            
               os.exit( 1 )
            end
            
            local code = Depend.runMain( modObj['__main'], self.option.shebangArgList )
            os.exit( code )
         end
      elseif _switchExp == Option.ModeKind.GoMod then
         for __index, path in ipairs( self.gomodMap:getModPathList(  ) ) do
            Util.println( path )
         end
      elseif _switchExp == Option.ModeKind.Exec then
         self:executeLnsAndTest( getBaseDir( self.option.scriptPath, self.option:get_projDir() ) )
      elseif _switchExp == Option.ModeKind.BootC then
         self:outputBootC( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.Builtin then
         self:outputBuiltin( self.option.scriptPath )
      elseif _switchExp == Option.ModeKind.MkMain then
         local mod = self:scriptPath2Module( self.option.scriptPath, self.option:get_projDir() )
         do
            local mess = convGo.outputGoMain( self:getGoAppName(  ), mod, self.option.testing, self.option.outputDir, self.option:get_runtimeOpt() )
            if mess ~= nil then
               Util.errorLog( mess )
            end
         end
      elseif _switchExp == Option.ModeKind.Indexer then
         self:build( _lune.newAlge( BuildMode.CreateAst), function ( ast )
         
            local indexer = NodeIndexer.Indexer._new(ast:get_exportInfo():get_processInfo())
            indexer:start( ast:get_node(), {[Nodes.NodeKind.get_Switch()] = true, [Nodes.NodeKind.get_Match()] = true, [Nodes.NodeKind.get_For()] = true, [Nodes.NodeKind.get_Apply()] = true} )
            indexer:dump(  )
         end
          )
      else 
         
            Util.println( "illegal mode" )
      end
   end
end


local function exec( args )

   local version = _lune.unwrapDefault( tonumber( Depend.getLuaVersion(  ):gsub( "^[^%d]+", "" ), nil ), 0.0)
   if version < 5.1 then
      Util.errorLog( string.format( "LuneScript doesn't support this lua version(%s). %s\n", version, "please use the version >= 5.1." ) )
      os.exit( 1 )
   end
   local option = Option.analyze( args )
   local front = Front._new(option)
   front:exec(  )
   if option.validProf then
      Depend.setRuntimeLog( true )
   end
end
_moduleObj.exec = exec


local function __main( argList )

   local list = {}
   for index, arg in ipairs( argList ) do
      if index > 1 then
         table.insert( list, arg )
      end
   end
   exec( list )
   return 0
end
_moduleObj.__main = __main




return _moduleObj
