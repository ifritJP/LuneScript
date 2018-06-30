-- ajiopjiojio
--[[ fjiaojfeap ]]
local Parser = require( 'lune.base.Parser' ).Parser
--local convLua = require( 'primal.convLua' )
local convLua = require( 'lune.base.convLua' ).filterObj

local path = arg[ 1 ]
local parser = Parser.new( path )

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
   local function createStream( val, writeFunc )
      local stream = { val = val }
      stream.write = writeFunc
      return stream
   end
   --local TransUnit = require( 'primal.TransUnit' )
   local TransUnit = require( 'lune.base.TransUnit' ).TransUnit
   local ast = TransUnit:createAST( parser )
   if mode == "ast" then
      ast:filter( require( 'lune.base.dumpNode' ).filterObj, "", 0 )
   elseif mode == "lua" then
      ast:filter( convLua:new( io.stdout ), nil, 0 )
   elseif mode == "save" then
      local func = function( self, txt )
	 self.val:write( txt )
      end
      local luaPath = path:gsub( ".lns$", ".lua" )
      if luaPath ~= path then
	 local fileObj = io.open( luaPath, "w" )
	 local stream = createStream( fileObj, func )
	 ast:filter( convLua:new( stream ), nil, 0 )
	 fileObj:close()
      end
   elseif mode == "exe" then
      local func = function( self, txt )
	 self.val = self.val .. txt
      end
      local stream = createStream( "", func )
      ast:filter( convLua:new( stream ), nil, 0 )

      local chunk, err = load( stream.txt )
      if err then
	 print( err )
      end
      chunk()
   else
      print( "illegal mode" )
   end
   
end
