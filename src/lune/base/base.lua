local Parser = require( 'lune.base.Parser' )
local convLua = require( 'lune.base.convLua' )
local TransUnit = require( 'lune.base.TransUnit' ).TransUnit
local Util = require( 'lune.base.Util' );


local scriptPath = arg[ 1 ]
local mode = arg[ 2 ]
local validProf = arg[ 3 ]

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

function _fcall( func, ... )
   return func( ... )
end

local function createStream( val, writeFunc )
   local stream = { val = val }
   stream.write = writeFunc
   return stream
end


_luneScript = {}
_luneScript.loadedMap = {}
_luneScript.loadedMetaMap = {}


function _luneScript.loadModule( module )
   if not _luneScript.loadedMap[ module ] then
      local searchPath = package.path
      searchPath = string.gsub( searchPath, "%.lua", ".lns" )
      local path = package.searchpath( module, searchPath )

      _luneScript.loadedMap[ module ] = _luneScript.loadFile( path, module )
   end
   return _luneScript.loadedMap[ module ]
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
   return TransUnit.new( convLua.MacroEvalImp.new() )
end

local function createAst( path, module )
   local transUnit = newTransUnit()
   return transUnit:createAST( createPaser( path ), nil, module )
end

local function getNode( ast )
   return ast.node
end

local function convert( ast, streamName, stream, exeFlag, inMacro, moduleTypeInfo )
   if convLua.Filter then
      local conv = convLua.Filter.new(
	 streamName, stream, exeFlag, inMacro, moduleTypeInfo );
      -- for key, val in pairs( ast ) do
      -- 	 Util.errorLog( string.format( "%s %s", key, val ) )
      -- end
      if ast.node.filterObj then
	 ast.node:filterObj( conv, nil, 0 )
      else
	 ast.node:filter( conv, nil, 0 )
      end
   else
      local conv = convLua.convFilter.new(
	 streamName, stream, exeFlag, inMacro, moduleTypeInfo )
      getNode( ast ):processFilter( conv, nil, 0 )
   end
end

function _luneScript.loadFile( path, module )

   Util.errorLog( "loadFile " .. path )
   
   local ast = createAst( path, module )
   
   local func = function( self, txt )
      self.val = self.val .. txt
   end
   local stream = createStream( "", func )
   convert( ast, path, stream, true, nil, ast.moduleTypeInfo )

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
   elseif mode == "lua" then
      local ast = createAst( scriptPath, module )
      convert( ast, scriptPath, io.stdout )
   elseif mode == "save" then

      Util.profile(
	 validProf,
      	 function()
	    local ast = createAst( scriptPath, module )
	    local func = function( self, txt )
	       self.val:write( txt )
	    end
	    local luaPath = scriptPath:gsub( ".lns$", ".lua" )
	    if luaPath ~= scriptPath then
	       local fileObj = io.open( luaPath, "w" )
	       local stream = createStream( fileObj, func )
	       convert( ast, scriptPath, stream )
	       fileObj:close()
	    end
      	 end, scriptPath .. ".profi" )
   elseif mode == "exe" then
      _luneScript.loadModule( module )
   else
      print( "illegal mode" )
   end
   
end
