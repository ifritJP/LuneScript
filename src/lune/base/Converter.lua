--lune/base/Converter.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Converter'
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


local Runner = _lune.loadModule( 'lune.base.Runner' )
local Util = _lune.loadModule( 'lune.base.Util' )
local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )
local Option = _lune.loadModule( 'lune.base.Option' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Types = _lune.loadModule( 'lune.base.Types' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local convLua = _lune.loadModule( 'lune.base.convLua' )
local convGo = _lune.loadModule( 'lune.base.convGo' )
local convPython = _lune.loadModule( 'lune.base.convPython' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local LuneControl = _lune.loadModule( 'lune.base.LuneControl' )
local OutputDepend = _lune.loadModule( 'lune.base.OutputDepend' )

local function createModuleInfo( ast, mod, moduleId )

   
   local exportInfo = ast:get_exportInfo()
   return frontInterface.ModuleInfo._new(exportInfo)
end

local function ast2LuaMain( ast, streamName, stream, metaStream, convMode, inMacro, option )

   local exportInfo = ast:get_exportInfo()
   local conv = convLua.createFilter( streamName, stream, metaStream, convMode, inMacro, exportInfo:get_moduleTypeInfo(), exportInfo:get_processInfo(), exportInfo:get_provideInfo():get_symbolKind(), ast:get_builtinFunc(), option.useLuneModule, option.targetLuaVer, option.testing, option.useIpairs, convLua.Option._new(option.mainModule, option:get_legacyNewName()) )
   return conv
end
_moduleObj.ast2LuaMain = ast2LuaMain

local function byteCompileFromLuaTxt( txt, stripDebugInfo )

   local ret
   
   do
      local chunk, err = _lune.loadstring52( txt )
      if chunk ~= nil then
         ret = string.dump( chunk, stripDebugInfo )
      else
         
         error( _lune.unwrapDefault( err, "load error") )
      end
      
   end
   
   return ret
end
_moduleObj.byteCompileFromLuaTxt = byteCompileFromLuaTxt



local AstCreater = {}
setmetatable( AstCreater, { __index = Runner.Runner } )
_moduleObj.AstCreater = AstCreater
function AstCreater:createAst( importModuleInfo, parserSrc, baseDir, stdinFile, analyzeModule, analyzeMode, pos )

   local transUnit = TransUnit.TransUnitCtrl._new(self.moduleId, importModuleInfo, convLua.MacroEvalImp._new(self.builtinFunc), true, analyzeModule, analyzeMode, pos, self.option.targetLuaVer, self.option.transCtrlInfo, self.builtinFunc)
   
   return transUnit:createAST( parserSrc, true, baseDir, stdinFile, false, self.mod, function ( exportInfo )
   
      self.exportInfo = exportInfo
      
      do
         local _exp = self.exportInfoReadyFlag
         if _exp ~= nil then
            _exp:set(  )
         end
      end
      
   end )
end
function AstCreater._new( importModuleInfo, parserSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, builtinFunc, option )
   local obj = {}
   AstCreater._setmeta( obj )
   if obj.__init then obj:__init( importModuleInfo, parserSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, builtinFunc, option ); end
   return obj
end
function AstCreater:__init(importModuleInfo, parserSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, builtinFunc, option) 
   Runner.Runner.__init( self)
   
   
   self.option = option
   self.builtinFunc = builtinFunc
   self.mod = mod
   self.moduleId = moduleId
   self.moduleInfo = nil
   self.exportInfo = nil
   self.ast = nil
   self.exportInfoReadyFlag = nil
   
   self.converter = function (  )
      local __func__ = '@lune.@base.@Converter.AstCreater.__init.<anonymous>'
   
      local ast = self:createAst( importModuleInfo, parserSrc, baseDir, option:get_stdinFile(), analyzeModule, analyzeMode, pos )
      self.ast = ast
      self.moduleInfo = createModuleInfo( ast, self.mod, self.moduleId )
      Log.log( Log.Level.Log, __func__, 150, function (  )
      
         return string.format( "generated AST -- %s", mod)
      end )
      
   end
   
   local lnsPath
   
   do
      local _matchExp = parserSrc
      if _matchExp[1] == Types.ParserSrc.LnsCode[1] then
         local _ = _matchExp[2][1]
         local path = _matchExp[2][2]
         local _ = _matchExp[2][3]
      
         lnsPath = path
      elseif _matchExp[1] == Types.ParserSrc.LnsPath[1] then
         local _ = _matchExp[2][1]
         local path = _matchExp[2][2]
         local _ = _matchExp[2][3]
         local _ = _matchExp[2][4]
      
         lnsPath = path
      elseif _matchExp[1] == Types.ParserSrc.Parser[1] then
         local path = _matchExp[2][1]
         local _ = _matchExp[2][2]
         local _ = _matchExp[2][3]
         local _ = _matchExp[2][4]
      
         lnsPath = path
      end
   end
   
   
   self:start( 0, string.format( "createAst - %s", lnsPath) )
   
end
function AstCreater:runMain(  )

   self.converter(  )
end
function AstCreater:getAst(  )

   
   local exportInfo = self.exportInfo
   if  nil == exportInfo then
      local _exportInfo = exportInfo
   
      Util.err( string.format( "exportInfo is nil -- %s", self.mod) )
   end
   
   local moduleMeta = frontInterface.ModuleMeta._new(exportInfo:get_streamName(), _lune.newAlge( frontInterface.MetaOrModule.Export, {exportInfo}))
   return _lune.unwrap( self.ast), _lune.unwrap( self.moduleInfo), moduleMeta
end
function AstCreater:getExportInfo(  )
   local __func__ = '@lune.@base.@Converter.AstCreater.getExportInfo'

   do
      local _exp = self.exportInfoReadyFlag
      if _exp ~= nil then
         _exp:wait(  )
      end
   end
   
   if not self.exportInfo then
      Log.log( Log.Level.Err, __func__, 192, function (  )
      
         return string.format( "exportInfo is nil -- %s", self.mod)
      end )
      
   end
   
   
   return self.exportInfo
end
function AstCreater._setmeta( obj )
  setmetatable( obj, { __index = AstCreater  } )
end


local CreateAstResult = {}
CreateAstResult._name2Val = {}
_moduleObj.CreateAstResult = CreateAstResult
function CreateAstResult:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "CreateAstResult.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function CreateAstResult._from( val )
   return _lune._AlgeFrom( CreateAstResult, val )
end

CreateAstResult.Ast = { "Ast", {{}}}
CreateAstResult._name2Val["Ast"] = CreateAstResult.Ast
CreateAstResult.Creater = { "Creater", {{}}}
CreateAstResult._name2Val["Creater"] = CreateAstResult.Creater


local function getAstFromResult( result )

   do
      local _matchExp = result
      if _matchExp[1] == CreateAstResult.Ast[1] then
         local ast = _matchExp[2][1]
      
         return ast
      elseif _matchExp[1] == CreateAstResult.Creater[1] then
         local creater = _matchExp[2][1]
      
         return (creater:getAst(  ) )
      end
   end
   
end

local function closeStreams( stream, metaStream, dependStream, metaPath, saveMetaFlag )

   local function txt2ModuleId( txt )
   
      local buildIdTxt = txt:gsub( "^_moduleObj.__buildId = ", "" ):gsub( '"', "" )
      return frontInterface.ModuleId.createIdFromTxt( buildIdTxt )
   end
   
   local function checkDiff( oldStream, newStream )
      local __func__ = '@lune.@base.@Converter.closeStreams.checkDiff'
   
      
      
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
               Log.log( Log.Level.Debug, __func__, 288, function (  )
               
                  return string.format( "<%s>, <%s>", oldLine, newLine)
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
   
   if stream ~= nil then
      stream:close(  )
   end
   
   if dependStream ~= nil then
      dependStream:close(  )
   end
   
   
   if metaStream ~= nil then
      if saveMetaFlag then
         
         local newMetaTxt = metaStream:get_txt()
         local oldMetaTxt = ""
         do
            local oldFileObj = io.open( metaPath )
            if oldFileObj ~= nil then
               oldMetaTxt = _lune.unwrapDefault( oldFileObj:read( "*a" ), "")
               oldFileObj:close(  )
            end
         end
         
         
         local sameFlag, txt = checkDiff( Parser.TxtStream._new(oldMetaTxt), Parser.TxtStream._new(newMetaTxt) )
         
         local function saveMeta( meta )
         
            do
               local fileObj = io.open( metaPath, "w" )
               if fileObj ~= nil then
                  fileObj:write( meta )
                  fileObj:close(  )
               else
                  Util.err( string.format( "failed to open -- %s", metaPath) )
               end
            end
            
         end
         
         if not sameFlag then
            saveMeta( newMetaTxt )
         else
          
            if txt ~= "" then
               saveMeta( txt )
            end
            
         end
         
      end
      
   end
   
end
_moduleObj.closeStreams = closeStreams

local LuaConverter = {}
setmetatable( LuaConverter, { __index = Runner.Runner } )
_moduleObj.LuaConverter = LuaConverter
function LuaConverter._new( luaPath, metaPath, dependsPath, astResult, convMode, path, byteCompile, stripDebugInfo, option )
   local obj = {}
   LuaConverter._setmeta( obj )
   if obj.__init then obj:__init( luaPath, metaPath, dependsPath, astResult, convMode, path, byteCompile, stripDebugInfo, option ); end
   return obj
end
function LuaConverter:__init(luaPath, metaPath, dependsPath, astResult, convMode, path, byteCompile, stripDebugInfo, option) 
   Runner.Runner.__init( self)
   
   
   self.luaPath = luaPath
   self.metaPath = metaPath
   self.dependsPath = dependsPath
   self.option = option
   self.stripDebugInfo = stripDebugInfo
   self.byteCompile = byteCompile
   self.byteStream = Util.memStream._new()
   self.byteMetaStream = Util.memStream._new()
   self.streamMem = Util.memStream._new()
   self.metaStreamMem = Util.memStream._new()
   self.dependsStreamMem = Util.memStream._new()
   self.astResult = astResult
   self.filterInfo = nil
   
   self.converterFunc = function (  )
   
      local stream = self.streamMem
      local metaStream = self.metaStreamMem
      
      local outStream, oMetaStream = stream, metaStream
      
      local ast = getAstFromResult( self.astResult )
      
      local needDepends = option.dependsPath ~= nil
      if needDepends then
         ast:get_node():processFilter( OutputDepend.createFilter( self.dependsStreamMem ), 1 )
      end
      
      
      if byteCompile then
         outStream = self.byteStream
         oMetaStream = self.byteMetaStream
      end
      
      
      local filterInfo = ast2LuaMain( ast, path, outStream, oMetaStream, convMode, false, option )
      self.filterInfo = filterInfo
      filterInfo:outputLua( _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )) )
   end
   self:start( 1, string.format( "convlua -- %s", path) )
end
function LuaConverter:runMain(  )

   self.converterFunc(  )
end
function LuaConverter:saveLua(  )

   
   
   local ast = getAstFromResult( self.astResult )
   _lune.nilacc( self.filterInfo, 'outputMeta', 'callmtd' , _lune.unwrap( _lune.__Cast( ast:get_node(), 3, Nodes.RootNode )) )
   
   if self.byteCompile then
      
      self.streamMem:write( byteCompileFromLuaTxt( self.byteStream:get_txt(), self.stripDebugInfo ) )
      self.metaStreamMem:write( byteCompileFromLuaTxt( self.byteMetaStream:get_txt(), true ) )
   end
   
   
   local luaCode = self.streamMem:get_txt()
   local metaTxt = self.metaStreamMem:get_txt()
   local dependTxt
   
   local needDepends = self.option.dependsPath ~= nil
   if needDepends then
      dependTxt = self.dependsStreamMem:get_txt()
   else
    
      dependTxt = nil
   end
   
   
   local streamDst = io.open( self.luaPath, "w" )
   if  nil == streamDst then
      local _streamDst = streamDst
   
      error( string.format( "write open error -- %s", self.luaPath) )
   end
   
   local dependsStreamDst = self.option:openDepend( self.dependsPath )
   
   streamDst:write( luaCode )
   local metaMemStream = Util.memStream._new()
   metaMemStream:write( metaTxt )
   if dependsStreamDst ~= nil then
      dependsStreamDst:write( _lune.unwrap( dependTxt) )
   end
   
   
   closeStreams( streamDst, metaMemStream, dependsStreamDst, self.metaPath, self.option.mode == Option.ModeKind.SaveMeta )
end
function LuaConverter._setmeta( obj )
  setmetatable( obj, { __index = LuaConverter  } )
end


local GoConverter = {}
setmetatable( GoConverter, { __index = Runner.Runner } )
_moduleObj.GoConverter = GoConverter
function GoConverter._new( scriptPath, astResult, mod, option, goOpt )
   local obj = {}
   GoConverter._setmeta( obj )
   if obj.__init then obj:__init( scriptPath, astResult, mod, option, goOpt ); end
   return obj
end
function GoConverter:__init(scriptPath, astResult, mod, option, goOpt) 
   Runner.Runner.__init( self)
   
   self.validFlag = true
   
   local path = mod:gsub( "%.", "/" ) .. ".go"
   
   do
      local dir = option.outputDir
      if dir ~= nil then
         path = string.format( "%s/%s", dir, path)
      end
   end
   
   
   self.path = path
   self.memStream = Util.memStream._new()
   
   self.converter = function (  )
   
      local ast = getAstFromResult( astResult )
      
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
                  self.validFlag = false
                  return 
               end
               
            end
         end
         
      end
      
      
      local conv = convGo.createFilter( option.testing, scriptPath, self.memStream, ast, goOpt )
      ast:get_node():processFilter( conv, convGo.Opt._new(ast:get_node()) )
   end
   self:start( 1, string.format( "convgo -- %s", scriptPath) )
end
function GoConverter:runMain(  )

   self.converter(  )
end
function GoConverter:saveGo(  )

   
   
   if not self.validFlag then
      return 
   end
   
   
   local file = io.open( self.path, "w" )
   if  nil == file then
      local _file = file
   
      error( string.format( "can't open file -- %s", self.path) )
   end
   
   
   file:write( self.memStream:get_txt() )
   
   file:close(  )
end
function GoConverter._setmeta( obj )
  setmetatable( obj, { __index = GoConverter  } )
end


local PythonConverter = {}
setmetatable( PythonConverter, { __index = Runner.Runner } )
_moduleObj.PythonConverter = PythonConverter
function PythonConverter._new( scriptPath, astResult, mod, option, pythonOpt )
   local obj = {}
   PythonConverter._setmeta( obj )
   if obj.__init then obj:__init( scriptPath, astResult, mod, option, pythonOpt ); end
   return obj
end
function PythonConverter:__init(scriptPath, astResult, mod, option, pythonOpt) 
   Runner.Runner.__init( self)
   
   self.validFlag = true
   
   local path = mod:gsub( "%.", "/" ) .. ".py"
   
   do
      local dir = option.outputDir
      if dir ~= nil then
         path = string.format( "%s/%s", dir, path)
      end
   end
   
   
   self.path = path
   self.memStream = Util.memStream._new()
   
   self.converter = function (  )
   
      local ast = getAstFromResult( astResult )
      
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
            
               if not _lune._Set_has(codeSet, LuneControl.Code.Python ) then
                  self.validFlag = false
                  return 
               end
               
            end
         end
         
      end
      
      
      local conv = convPython.createFilter( option.testing, scriptPath, self.memStream, ast, pythonOpt )
      ast:get_node():processFilter( conv, convPython.Opt._new(ast:get_node()) )
   end
   self:start( 1, string.format( "convpython -- %s", scriptPath) )
end
function PythonConverter:runMain(  )

   self.converter(  )
end
function PythonConverter:savePython(  )

   
   
   if not self.validFlag then
      return 
   end
   
   
   local file = io.open( self.path, "w" )
   if  nil == file then
      local _file = file
   
      error( string.format( "can't open file -- %s", self.path) )
   end
   
   
   file:write( self.memStream:get_txt() )
   
   file:close(  )
end
function PythonConverter._setmeta( obj )
  setmetatable( obj, { __index = PythonConverter  } )
end


return _moduleObj
