local Parser = require( 'lune.base.Parser' )
local convLua = require( 'lune.base.convLua' )
local TransUnit = require( 'lune.base.TransUnit' ).TransUnit
local Util = require( 'lune.base.Util' );
local Option = require( 'lune.base.Option' );

local option = Option.analyze( arg );

function _luneGetLocal( varName )
   local index = 1
   while true do
      local name, val = debug.getlocal( 3, index )
      if name == varName then
	 return val
      end
      if not name then
	 break
      end
      print( name, val )
      index = index + 1
   end
   error( "not found -- " .. varName )
end

function _luneSym2Str( val )
   if not val then
      return nil
   end

   if type( val ) ~= "table" then
      return string.format( "%s", val )
   end

   local txt = ""
   for key, val in pairs( val ) do
      txt = txt .. val
   end
   return txt
end


local function createStream( val, writeFunc )
   local stream = { val = val }
   stream.write = writeFunc
   return stream
end


_luneScript = {}
_luneScript.loadedMap = {}
_luneScript.loadedMetaMap = {}

function _luneScript.error( message ) 
      Util.errorLog( message );
      Util.printStackTrace()
      os.exit( 1 );
end

function _luneScript.loadLua( path )
   local chunk, err = loadfile( path )
   if err then
      Util.errorLog( err )
   end
   if not chunk then
      error( "failed to error" )
   end
   return chunk()
end

function _luneScript.searchModule( module )
   local lnsSearchPath = package.path
   lnsSearchPath = string.gsub( lnsSearchPath, "%.lua", ".lns" )
   return package.searchpath( module, lnsSearchPath )
end
      

function _luneScript.loadModule( module )
   if not _luneScript.loadedMap[ module ] then
      local lnsPath = _luneScript.searchModule( module )
      local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )
      local luaSearchPath = package.path
      local bakSearchPath = package.path
	    
      if option.outputDir then
	 luaSearchPath = string.format( "%s/?.lua;%s", option.outputDir, package.path )
	 luaPath = package.searchpath( module, luaSearchPath )
	 package.path = luaSearchPath
      end

      local mod = nil
      if luaPath and Util.getReadyCode( lnsPath, luaPath ) then
	 local metaPath = string.gsub( luaPath, "%.lua$", ".meta" )
	 if Util.getReadyCode( lnsPath, metaPath ) then
	    mod = _luneScript.loadLua( luaPath )
	    local meta = _luneScript.loadLua( metaPath )

	    -- for key, val in pairs( meta ) do
	    --    mod[ key ] = val
	    -- end
	    
	    -- _luneScript.loadedMap[ module ] = mod
	    local info = {}
	    info.mod = mod
	    info.meta = meta
	    _luneScript.loadedMap[ module ] = info
	    
	 end
      end
      if option.outputDir then
	 package.path = bakSearchPath
      end
      if not mod then
	 local mod, meta = _luneScript.loadFile( lnsPath, module )
	 local info = {}
	 info.mod = mod
	 info.meta = meta

	 -- for key, val in pairs( meta ) do
	 --    mod[ key ] = val
	 -- end
	 
	 _luneScript.loadedMap[ module ] = info
      end
   end
   local ret = _luneScript.loadedMap[ module ]
   if ret then
      return ret.mod, ret.meta
   end
   error( "load error", module )
end

local function createPaser( path, module )
   local parser = Parser.StreamParser.create( path, false, module )
   if not parser then
      error( "failed to open " .. path );
   end
   return parser
end


local function newTransUnit( analyzeModule, analyzeMode, pos )
   return TransUnit.new( convLua.MacroEvalImp.new( option.mode ),
			 analyzeModule, analyzeMode, pos )
end

local function createAst( path, module, analyzeModule, analyzeMode, pos )
   local transUnit = newTransUnit( analyzeModule, analyzeMode, pos )
   return transUnit:createAST( createPaser( path, module ), nil, module )
end

local function getNode( ast )
   return ast.node
end

local function convert( ast, streamName, stream, metaStream,
			convMode, inMacro )
   local conv = convLua.convFilter.new(
      streamName, stream, metaStream, convMode, inMacro, ast.moduleTypeInfo )
   getNode( ast ):processFilter( conv, nil, 0 )
end

function _luneScript.loadFile( path, module, analyzeMode, pos )

   local ast = createAst( path, module, module, analyzeMode, pos )
   
   local func = function( self, txt )
      self.val = self.val .. txt
   end
   local stream = createStream( "", func )
   local metaStream = createStream( "", func )

   convert( ast, path, stream, metaStream, "exe", nil )
   

   local loadFunc = function( txt )
      --print( txt )
      local chunk, err = load( txt )
      if err then
	 print( err )
      end
      if not chunk then
	 error( "failed to error" )
      end
      return chunk()
   end
   return loadFunc( stream.val ), loadFunc( metaStream.val )
end



local module = string.gsub( option.scriptPath, "/", "." )
if option.mode == "token" then
   local parser = createPaser( option.scriptPath, module )
   while true do
      local token = parser:getToken()
      if not token then
	 break
      end
      print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
   end
else
   module = string.gsub( module, "%.lns$", "" )
   if option.mode == "ast" then
      Util.profile(
	 option.validProf,
	 function()
	    local ast = createAst( option.scriptPath, module )
	    local dumpNode = require( 'lune.base.dumpNode' ).dumpFilter;
	    getNode( ast ):processFilter( dumpNode, "", 0 )
	 end, option.scriptPath .. ".profi" )
   elseif option.mode == "diag" then
      Util.setErrorCode( 0 );
      createAst( option.scriptPath, module, nil, "diag" )
   elseif option.mode == "comp" then
      createAst( option.scriptPath, module,
		 option.analyzeModule, "comp", option.analyzePos )
   elseif option.mode == "lua" or option.mode == "LUA" then
      local ast = createAst( option.scriptPath, module )
      convert( ast, option.scriptPath, io.stdout, io.stdout, option.mode )
   elseif option.mode == "save" or option.mode == "SAVE" then
      Util.profile(
	 option.validProf,
      	 function()
	    local ast = createAst( option.scriptPath, module )
	    local func = function( self, txt )
	       self.val:write( txt )
	    end
	    local luaPath = option.scriptPath:gsub( "%.lns$", ".lua" )
	    local metaPath = option.scriptPath:gsub( "%.lns$", ".meta" )
	    if option.outputDir then
	       local filename = module:gsub( "%.", "/" )
	       luaPath = string.format( "%s/%s.lua", option.outputDir, filename )
	       metaPath = string.format( "%s/%s.meta", option.outputDir, filename )
	    end

	    
	    if luaPath ~= option.scriptPath then
	       local fileObj = io.open( luaPath, "w" )
	       local stream = createStream( fileObj, func )

	       local metaFileObj = nil
	       local metaStream = stream
	       if option.mode == "SAVE" then
		  metaFileObj = io.open( metaPath, "w" )
		  metaStream = createStream( metaFileObj, func )
	       end
	       
	       convert( ast, option.scriptPath, stream, metaStream, option.mode )
	       fileObj:close()
	       if metaFileObj then
		  metaFileObj:close()
	       end
	    end
      	 end, option.scriptPath .. ".profi" )
   elseif option.mode == "exe" then
      _luneScript.loadModule( module )
   else
      print( "illegal mode" )
   end
   
end
