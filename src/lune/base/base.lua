local Parser = require( 'lune.base.Parser' ).Parser
local convLua = require( 'lune.base.convLua' ).filterObj
local TransUnit = require( 'lune.base.TransUnit' ).TransUnit

local scriptPath = arg[ 1 ]

local function createStream( val, writeFunc )
   local stream = { val = val }
   stream.write = writeFunc
   return stream
end


_luneScript = {}
_luneScript.loadedMap = {}


function _luneScript.loadModule( module )
   if not _luneScript.loadedMap[ module ] then
      local searchPath = package.path
      searchPath = string.gsub( searchPath, "%.lua", ".lns" )
      local path = package.searchpath( module, searchPath )

      _luneScript.loadedMap[ module ] = _luneScript.loadFile( path )
   end
   return _luneScript.loadedMap[ module ]
end

function _luneScript.loadFile( path )
   local transUnit = TransUnit.new()
   local ast = transUnit:createAST( Parser.new( path ) )
   
   local func = function( self, txt )
      self.val = self.val .. txt
   end
   local stream = createStream( "", func )
   ast:filter( convLua.new( path, stream, true ), nil, 0 )

   local chunk, err = load( stream.val )
   if err then
      print( err )
   end
   return chunk()
end


local parser = Parser.new( scriptPath )

local mode = arg[ 2 ]
if mode == "token" then
   while true do
      local token = parser:getToken()
      if not token then
	 break
      end
      print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
   end
else
   if mode == "ast" then
      local ast = TransUnit:createAST( parser )
      ast:filter( require( 'lune.base.dumpNode' ).filterObj, "", 0 )
   elseif mode == "lua" then
      local ast = TransUnit:createAST( parser )
      ast:filter( convLua.new( scriptPath, io.stdout ), nil, 0 )
   elseif mode == "save" then
      local ast = TransUnit:createAST( parser )
      local func = function( self, txt )
	 self.val:write( txt )
      end
      local luaPath = scriptPath:gsub( ".lns$", ".lua" )
      if luaPath ~= scriptPath then
	 local fileObj = io.open( luaPath, "w" )
	 local stream = createStream( fileObj, func )
	 ast:filter( convLua.new( scriptPath, stream ), nil, 0 )
	 fileObj:close()
      end
   elseif mode == "exe" then
      local module = string.gsub( scriptPath, "/", "." )
      module = string.gsub( module, "%.lns$", "" )
      _luneScript.loadModule( module )
   else
      print( "illegal mode" )
   end
   
end
