local Parser = require( 'lune.base.Parser' )
local convLua = require( 'lune.base.convLua' )
local TransUnit = require( 'lune.base.TransUnit' ).TransUnit
local Util = require( 'lune.base.Util' );

local scriptPath = arg[ 1 ]
local mode = arg[ 2 ]
local outputDir = arg[ 3 ]
local validProf = arg[ 4 ]
local outputMetaFlag = true

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
      Util.debugLog()
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

function _luneScript.loadModule( module )
   if not _luneScript.loadedMap[ module ] then
      local lnsSearchPath = package.path
      lnsSearchPath = string.gsub( lnsSearchPath, "%.lua", ".lns" )
      local lnsPath = package.searchpath( module, lnsSearchPath )
      local luaPath = string.gsub( lnsPath, "%.lns$", ".lua" )

      if outputDir then
	 local luaSearchPath =
	    string.format( "%s/?.lua;%s", outputDir, package.path )
	 luaPath = package.searchpath( module, luaSearchPath )
      end

      local mod = nil
      if luaPath and Util.getReadyCode( lnsPath, luaPath ) then
	 local metaPath = string.gsub( luaPath, "%.lua$", ".meta" )
	 if Util.getReadyCode( lnsPath, metaPath ) then
	    mod = _luneScript.loadLua( luaPath )
	    local meta = _luneScript.loadLua( metaPath )

	    for key, val in pairs( meta ) do
	       mod[ key ] = val
	    end
	    
	    _luneScript.loadedMap[ module ] = mod
	 end
      end
      if not mod then
	 _luneScript.loadedMap[ module ] = _luneScript.loadFile( lnsPath, module )
      end
   end
   local ret = _luneScript.loadedMap[ module ]
   if ret then
      return ret
   end
   error( "load error", module )
end

function _luneScript.loadMeta( module )
   if not _luneScript.loadedMetaMap[ module ] then
      local searchPath = package.path
      searchPath = string.gsub( searchPath, "%.lua", ".lnm" )
      local path = package.searchpath( module, searchPath )

      _luneScript.loadedMetaMap[ module ] = _luneScript.loadFile( path, module )
   end
   return _luneScript.loadedMetaMap[ module ]
end

local function createPaser( path )
   local parser = Parser.StreamParser.create( path )
   if not parser then
      error( "failed to open " .. path );
   end
   return parser
end


local function newTransUnit()
   return TransUnit.new( convLua.MacroEvalImp.new( mode ) )
end

local function createAst( path, module )
   local transUnit = newTransUnit()
   return transUnit:createAST( createPaser( path ), nil, module )
end

local function getNode( ast )
   return ast.node
end

local function convert( ast, streamName, stream, metaStream,
			convMode, inMacro, moduleTypeInfo )
   local conv = convLua.convFilter.new(
      streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo )
   getNode( ast ):processFilter( conv, nil, 0 )
end

function _luneScript.loadFile( path, module )

   local ast = createAst( path, module )
   
   local func = function( self, txt )
      self.val = self.val .. txt
   end
   local stream = createStream( "", func )

   convert( ast, path, stream, stream, "exe", nil, ast.moduleTypeInfo )

   local chunk, err = load( stream.val )
   if err then
      print( err )
   end
   if not chunk then
      error( "failed to error" )
   end
   return chunk()
end



if mode == "token" then
   local parser = createPaser( scriptPath )
   while true do
      local token = parser:getToken()
      if not token then
	 break
      end
      print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
   end
else
   local module = string.gsub( scriptPath, "/", "." )
   module = string.gsub( module, "%.lns$", "" )
   if mode == "ast" then
      Util.profile(
	 validProf,
	 function()
	    local ast = createAst( scriptPath, module )
	    local dumpNode = require( 'lune.base.dumpNode' ).dumpFilter;
	    getNode( ast ):processFilter( dumpNode, "", 0 )
	 end, scriptPath .. ".profi" )
   elseif mode == "lua" or mode == "LUA" then
      local ast = createAst( scriptPath, module )
      convert( ast, scriptPath, io.stdout, io.stdout, mode )
   elseif mode == "save" or mode == "SAVE" then
      Util.profile(
	 validProf,
      	 function()
	    local ast = createAst( scriptPath, module )
	    local func = function( self, txt )
	       self.val:write( txt )
	    end
	    local luaPath = scriptPath:gsub( "%.lns$", ".lua" )
	    local metaPath = scriptPath:gsub( "%.lns$", ".meta" )
	    if outputDir then
	       local filename = module:gsub( "%.", "/" )
	       luaPath = string.format( "%s/%s.lua", outputDir, filename )
	       metaPath = string.format( "%s/%s.meta", outputDir, filename )
	    end

	    
	    if luaPath ~= scriptPath then
	       local fileObj = io.open( luaPath, "w" )
	       local stream = createStream( fileObj, func )

	       local metaFileObj = nil
	       local metaStream = stream
	       if mode == "SAVE" then
		  metaFileObj = io.open( metaPath, "w" )
		  metaStream = createStream( metaFileObj, func )
	       end
	       
	       convert( ast, scriptPath, stream, metaStream, mode )
	       fileObj:close()
	       if metaFileObj then
		  metaFileObj:close()
	       end
	    end
      	 end, scriptPath .. ".profi" )
   elseif mode == "exe" then
      _luneScript.loadModule( module )
   else
      print( "illegal mode" )
   end
   
end
