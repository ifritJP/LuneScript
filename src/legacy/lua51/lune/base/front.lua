--lune/base/front.lns
local _moduleObj = {}
local __mod__ = 'lune.base.front'
if not _lune then
   _lune = {}
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

local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local convLua = _lune.loadModule( 'lune.base.convLua' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Option = _lune.loadModule( 'lune.base.Option' )
local dumpNode = _lune.loadModule( 'lune.base.dumpNode' )
local glueFilter = _lune.loadModule( 'lune.base.glueFilter' )
local Depend = _lune.loadModule( 'lune.base.Depend' )
local OutputDepend = _lune.loadModule( 'lune.base.OutputDepend' )
local Ver = _lune.loadModule( 'lune.base.Ver' )

local forceUpdateMeta = true
function __luneGetLocal( varName )

   local index = 1
   while true do
      local name, val = debug.getlocal( 3, index )
      if name == varName then
         return val
      end
      
      if not name then
         break
      end
      
      index = index + 1
   end
   
   error( "not found -- " .. varName )
end

function __luneSym2Str( val )

   do
      local _exp = val
      if _exp ~= nil then
         if type( _exp ) ~= "table" then
            return string.format( "%s", tostring( _exp) )
         end
         
         local txt = ""
         for __index, item in pairs( _exp ) do
            txt = txt .. item
         end
         
         return txt
      end
   end
   
   return nil
end

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
function Front.new( option )
   local obj = {}
   Front.setmeta( obj )
   if obj.__init then obj:__init( option ); end
   return obj
end
function Front:__init(option) 
   self.option = option
   self.loadedMap = {}
   self.loadedMetaMap = {}
   self.convertedMap = {}
   frontInterface.setFront( self )
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

   local mod = string.gsub( path, "/", "." )
   return (string.gsub( mod, "%.lns$", "" ) )
end
_moduleObj.scriptPath2Module = scriptPath2Module
function Front:createPaser(  )

   local mod = scriptPath2Module( self.option.scriptPath )
   return createPaser( self.option.scriptPath, mod )
end

function Front:createAst( importModuleInfo, parser, mod, moduleId, analyzeModule, analyzeMode, pos )

   local transUnit = TransUnit.TransUnit.new(moduleId, importModuleInfo, convLua.MacroEvalImp.new(self.option.mode), analyzeModule, analyzeMode, pos, self.option.targetLuaVer, self.option.transCtrlInfo)
   return transUnit:createAST( parser, false, mod )
end

function Front:convert( ast, streamName, stream, metaStream, convMode, inMacro )

   local conv = convLua.createFilter( streamName, stream, metaStream, convMode, inMacro, ast:get_moduleTypeInfo(), ast:get_moduleSymbolKind(), self.option.useLuneModule, self.option.targetLuaVer )
   ast:get_node():processFilter( conv, nil, 0 )
end

local function loadFromChunk( chunk, err )

   if err ~= nil then
      Util.errorLog( err )
   end
   
   if chunk ~= nil then
      return _lune.unwrap( chunk(  ))
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

   local transUnit = TransUnit.TransUnit.new(frontInterface.ModuleId.tempId, importModuleInfo, convLua.MacroEvalImp.new(self.option.mode), nil, nil, nil, self.option.targetLuaVer, self.option.transCtrlInfo)
   local stream = Parser.TxtStream.new(txt)
   local parser = Parser.StreamParser.new(stream, name, false)
   local ast = transUnit:createAST( parser, false, nil )
   local metaTxt, luaTxt = self:convertFromAst( ast, name, convLua.ConvMode.Exec )
   return loadFromLuaTxt( luaTxt )
end

local DependMetaInfo = {}
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
function MetaForBuildId.setmeta( obj )
  setmetatable( obj, { __index = MetaForBuildId  } )
end
function MetaForBuildId.new( __buildId, __dependModuleMap, __subModuleMap )
   local obj = {}
   MetaForBuildId.setmeta( obj )
   if obj.__init then
      obj:__init( __buildId, __dependModuleMap, __subModuleMap )
   end        
   return obj 
end         
function MetaForBuildId:__init( __buildId, __dependModuleMap, __subModuleMap ) 

   self.__buildId = __buildId
   self.__dependModuleMap = __dependModuleMap
   self.__subModuleMap = __subModuleMap
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
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local function getMetaInfo( lnsPath, mod, outdir )

   local moduleMetaPath = lnsPath
   if outdir ~= nil then
      moduleMetaPath = string.format( "%s/%s", outdir, mod:gsub( "%.", "/" ))
   end
   
   moduleMetaPath = moduleMetaPath:gsub( "%.lns$", ".meta" )
   do
      local fileObj = io.open( moduleMetaPath )
      if fileObj ~= nil then
         do
            local luaCode = fileObj:read( "*a" )
            if luaCode ~= nil then
               local meta = MetaForBuildId._fromStem( loadFromLuaTxt( luaCode ) )
               return meta, moduleMetaPath, luaCode
            end
         end
         
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
      local buildId = frontInterface.ModuleId.createIdFromTxt( metaInfo.__buildId )
      buildCount = buildId:get_buildCount()
      if fileTime ~= buildId:get_modTime() then
         buildCount = 0
      end
      
   end
   
   return frontInterface.ModuleId.createId( fileTime, buildCount )
end

function Front:getModuleIdAndCheckUptodate( lnsPath, mod )

   local uptodate = false
   local metaInfo, metaPath, metaCode = getMetaInfo( lnsPath, mod, self.option.outputDir )
   local function checkDependUptodate( metaTime, dependModuleMap )
   
      for depMod, dependMap in pairs( dependModuleMap ) do
         if dependMap.use then
            local modMetaPath = self:searchModuleFile( depMod, ".meta", self.option.outputDir )
            if  nil == modMetaPath then
               local _modMetaPath = modMetaPath
            
               return false
            end
            
            local time = Depend.getFileLastModifiedTime( modMetaPath )
            if  nil == time then
               local _time = time
            
               return false
            end
            
            if time > metaTime then
               return false
            end
            
         end
         
      end
      
      return true
   end
   
   if metaInfo ~= nil then
      local buildId = frontInterface.ModuleId.createIdFromTxt( metaInfo.__buildId )
      if buildId ~= frontInterface.ModuleId.tempId then
         local lnsTime = Depend.getFileLastModifiedTime( lnsPath )
         local metaTime = Depend.getFileLastModifiedTime( metaPath )
         if lnsTime ~= nil and metaTime ~= nil then
            if lnsTime == buildId:get_modTime() then
               uptodate = checkDependUptodate( metaTime, metaInfo.__dependModuleMap )
            end
            
         end
         
      end
      
   end
   
   local moduleId = getModuleId( lnsPath, mod, self.option.outputDir, metaInfo )
   return moduleId, uptodate, metaInfo, metaCode
end

function Front:loadFile( importModuleInfo, path, mod, onlyMeta )

   local ast = self:createAst( importModuleInfo, createPaser( path, mod ), mod, getModuleId( path, mod ), nil, TransUnit.AnalyzeMode.Compile, nil )
   local convMode = convLua.ConvMode.Exec
   local metaTxt, luaTxt = self:convertFromAst( ast, path, convMode )
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
   
   local meta = loadFromLuaTxt( metaTxt )
   if onlyMeta then
      return meta, luaTxt
   end
   
   return meta, loadFromLuaTxt( luaTxt )
end

function Front:searchModule( mod )

   return self:searchModuleFile( mod, ".lns", nil )
end

function Front:searchLuaFile( moduleFullName, addSearchPath )

   local luaSearchPath = package.path
   do
      local _exp = addSearchPath
      if _exp ~= nil then
         luaSearchPath = string.format( "%s/?.lua;%s", tostring( addSearchPath), package.path )
      end
   end
   
   local foundPath = Depend.searchpath( moduleFullName, luaSearchPath )
   if  nil == foundPath then
      local _foundPath = foundPath
   
      return nil
   end
   
   return (foundPath:gsub( "^./", "" ) )
end

function Front:checkUptodateMeta( metaPath, addSearchPath )

   local metaObj = self:loadLua( metaPath )
   if  nil == metaObj then
      local _metaObj = metaObj
   
      return nil
   end
   
   local meta = metaObj
   if meta.__formatVersion ~= Ver.metaVersion then
      return nil
   end
   
   for moduleFullName, dependInfo in pairs( meta.__dependModuleMap ) do
      do
         local moduleLuaPath = self:searchLuaFile( moduleFullName, addSearchPath )
         if moduleLuaPath ~= nil then
            local moduleLnsPath = moduleLuaPath:gsub( "%.lua$", ".lns" )
            if not Util.getReadyCode( moduleLnsPath, metaPath ) then
               return nil
            end
            
            local moduleMetaPath = moduleLuaPath:gsub( "%.lua$", ".meta" )
            if Util.existFile( moduleMetaPath ) and not Util.getReadyCode( moduleMetaPath, metaPath ) then
               return nil
            end
            
         end
      end
      
   end
   
   return meta
end

function Front:loadModule( mod )

   if self.loadedMap[mod] == nil then
      do
         local luaTxt = self.convertedMap[mod]
         if luaTxt ~= nil then
            do
               local meta = self.loadedMetaMap[mod]
               if meta ~= nil then
                  self.loadedMap[mod] = LoadInfo.new(loadFromLuaTxt( luaTxt ), meta)
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
                           loadVal = self:loadLua( luaPath )
                           do
                              local _exp = loadVal
                              if _exp ~= nil then
                                 do
                                    local meta = self:checkUptodateMeta( metaPath, self.option.outputDir )
                                    if meta ~= nil then
                                       self.loadedMap[mod] = LoadInfo.new(_exp, meta)
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
                     local meta, workVal = self:loadFile( frontInterface.ImportModuleInfo.new(), lnsPath, mod, false )
                     self.loadedMap[mod] = LoadInfo.new(workVal, meta)
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   do
      local _exp = self.loadedMap[mod]
      if _exp ~= nil then
         return _lune.unwrap( _exp.mod), _lune.unwrap( _exp.meta)
      end
   end
   
   error( string.format( "load error, %s", mod) )
end

function Front:loadMeta( importModuleInfo, mod )

   if self.loadedMetaMap[mod] == nil then
      do
         local _exp = self.loadedMap[mod]
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
                           
                        end
                        
                     end
                     
                  end
                  
                  if meta == nil then
                     local metawork, luaTxt = self:loadFile( importModuleInfo, lnsPath, mod, true )
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

   frontInterface.setFront( self )
   local parser = self:createPaser(  )
   while true do
      local token = parser:getToken(  )
      if  nil == token then
         local _token = token
      
         break
      end
      
      print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
   end
   
end

function Front:dumpAst(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   Util.profile( self.option.validProf, function (  )
   
      local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
      ast:get_node():processFilter( dumpNode.dumpFilter.new(), "", 0 )
   end
   , self.option.scriptPath .. ".profi" )
end

function Front:checkDiag(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   Util.setErrorCode( 0 )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Diag )
end

function Front:complete(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), self.option.analyzeModule, TransUnit.AnalyzeMode.Complete, self.option.analyzePos )
end

function Front:createGlue(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   local ast = self:createAst( frontInterface.ImportModuleInfo.new(), self:createPaser(  ), mod, getModuleId( self.option.scriptPath, mod ), nil, TransUnit.AnalyzeMode.Compile )
   local glue = glueFilter.glueFilter.new(self.option.outputDir)
   ast:get_node():processFilter( glue )
end



function Front:convertToLuaFromScript( convMode, path, mod, byteCompile, stripDebugInfo, openOStream, closeOStream )

   local moduleId, uptodate, metaInfo, metaCode = self:getModuleIdAndCheckUptodate( path, mod )
   if moduleId == frontInterface.ModuleId.tempId then
      Util.err( string.format( "not found -- %s", path) )
   end
   
   local stream, metaStream, dependsStream = openOStream( not uptodate )
   if not uptodate then
      if stream ~= nil then
         local ast = self:createAst( frontInterface.ImportModuleInfo.new(), createPaser( path, mod ), mod, moduleId, nil, TransUnit.AnalyzeMode.Compile )
         if dependsStream ~= nil then
            ast:get_node():processFilter( OutputDepend.createFilter( dependsStream ) )
         end
         
         local oStream = stream
         local oMetaStream = metaStream
         local byteStream = Util.memStream.new()
         local byteMetaStream = Util.memStream.new()
         if byteCompile then
            oStream = byteStream
            oMetaStream = byteMetaStream
         end
         
         self:convert( ast, path, oStream, oMetaStream, convMode, false )
         if byteCompile then
            stream:write( byteCompileFromLuaTxt( byteStream:get_txt(), stripDebugInfo ) )
            if metaStream ~= stream then
               metaStream:write( byteCompileFromLuaTxt( byteMetaStream:get_txt(), true ) )
            end
            
         end
         
      end
      
   else
    
      Util.errorLog( "uptodate -- " .. path )
      metaStream:write( metaCode )
      if dependsStream ~= nil and metaInfo ~= nil then
         local dependInfo = OutputDepend.DependInfo.new(mod)
         for dependMod, moduleInfo in pairs( metaInfo.__dependModuleMap ) do
            dependInfo:addImpotModule( dependMod )
         end
         
         for __index, subMod in pairs( metaInfo.__subModuleMap ) do
            dependInfo:addSubMod( subMod )
         end
         
         dependInfo:output( dependsStream )
      end
      
   end
   
   if closeOStream ~= nil then
      closeOStream( stream, metaStream, dependsStream )
   end
   
end

function Front:convertToLua(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   local convMode = convLua.ConvMode.Convert
   if self.option.mode == Option.ModeKind.LuaMeta then
      convMode = convLua.ConvMode.ConvMeta
   end
   
   self:convertToLuaFromScript( convMode, self.option.scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( validLuaStream )
   
      return io.stdout, io.stdout, self.option.dependsStream
   end
   , nil )
end

function Front:saveToLua(  )

   frontInterface.setFront( self )
   local mod = scriptPath2Module( self.option.scriptPath )
   Util.profile( self.option.validProf, function (  )
   
      local luaPath = self.option.scriptPath:gsub( "%.lns$", ".lua" )
      local metaPath = self.option.scriptPath:gsub( "%.lns$", ".meta" )
      if self.option.outputDir then
         local filename = mod:gsub( "%.", "/" )
         luaPath = string.format( "%s/%s.lua", tostring( self.option.outputDir), filename)
         metaPath = string.format( "%s/%s.meta", tostring( self.option.outputDir), filename)
      end
      
      if luaPath == self.option.scriptPath then
         Util.errorLog( "%s is illegal filename." )
      else
       
         local convMode = convLua.ConvMode.Convert
         if self.option.mode == "SAVE" then
            convMode = convLua.ConvMode.ConvMeta
         end
         
         local metaFileObj = nil
         local tempMetaPath = metaPath .. ".tmp"
         self:convertToLuaFromScript( convMode, self.option.scriptPath, mod, self.option.byteCompile, self.option.stripDebugInfo, function ( validLuaStream )
         
            local stream = nil
            if self.option.mode ~= "SAVE" or validLuaStream then
               local fileObj = io.open( luaPath, "w" )
               if  nil == fileObj then
                  local _fileObj = fileObj
               
                  error( string.format( "write open error -- %s", luaPath) )
               end
               
               stream = fileObj
            end
            
            local metaStream = stream
            local convMode = convLua.ConvMode.Convert
            if self.option.mode == "SAVE" then
               convMode = convLua.ConvMode.ConvMeta
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
            
            return stream, _lune.unwrap( metaStream), self.option.dependsStream
         end
         , function ( stream, metaStream, dependStream )
         
            if stream ~= nil then
               stream:close(  )
            end
            
            if metaFileObj ~= nil then
               metaFileObj:flush(  )
               metaFileObj:seek( "set", 0 )
               local oldMetaTxt = metaFileObj:read( "*a" )
               metaFileObj:close(  )
               local newMetaTxt = ""
               do
                  local oldFileObj = io.open( metaPath )
                  if oldFileObj ~= nil then
                     newMetaTxt = _lune.unwrapDefault( oldFileObj:read( "*a" ), "")
                     oldFileObj:close(  )
                  end
               end
               
               if forceUpdateMeta or newMetaTxt ~= oldMetaTxt then
                  os.rename( tempMetaPath, metaPath )
               else
                
                  os.remove( tempMetaPath )
               end
               
            end
            
         end
          )
      end
      
   end
   , self.option.scriptPath .. ".profi" )
end

function Front:exec(  )

   do
      local _switchExp = self.option.mode
      if _switchExp == Option.ModeKind.Token then
         self:dumpTokenize(  )
      elseif _switchExp == Option.ModeKind.Ast then
         self:dumpAst(  )
      elseif _switchExp == Option.ModeKind.Diag then
         self:checkDiag(  )
      elseif _switchExp == Option.ModeKind.Complete then
         self:complete(  )
      elseif _switchExp == Option.ModeKind.Glue then
         self:createGlue(  )
      elseif _switchExp == Option.ModeKind.Lua or _switchExp == Option.ModeKind.LuaMeta then
         self:convertToLua(  )
      elseif _switchExp == Option.ModeKind.Save or _switchExp == Option.ModeKind.SaveMeta then
         self:saveToLua(  )
      elseif _switchExp == Option.ModeKind.Exec then
         frontInterface.setFront( self )
         self:loadModule( scriptPath2Module( self.option.scriptPath ) )
      else 
         
            print( "illegal mode" )
      end
   end
   
end

local function exec( args )

   local version = tonumber( _VERSION:gsub( "^[^%d]+", "" ), nil )
   if version < 5.1 then
      io.stderr:write( string.format( "LuneScript doesn't support this lua version(%s). %s\n", tostring( version), "please use the version >= 5.1." ) )
      os.exit( 1 )
   end
   
   local option = Option.analyze( args )
   local front = Front.new(option)
   front:exec(  )
end
_moduleObj.exec = exec
return _moduleObj
