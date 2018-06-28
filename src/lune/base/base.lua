-- ajiopjiojio
--[[ fjiaojfeap ]]
local Parser = require( 'lune.base.Parser' ).Parser

local parser = Parser.new( arg[ 1 ] )

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
   local TransUnit = require( 'primal.TransUnit' )
   local ast = TransUnit:createAST( parser )
   if mode == "ast" then
      ast:filter( require( 'lune.base.dumpNode' ), "", 0 )
   elseif mode == "lua" then
      ast:filter( require( 'primal.convLua' ), nil, 0 )
   elseif mode == "exe" then
      local convLua = require( 'primal.convLua' )

      local stream = { txt = "" }
      stream.write = function( self, txt )
	 self.txt = self.txt .. txt
      end

      ast:filter( convLua:new( stream ), nil, 0 )

      load( stream.txt )()
   else
      print( "illegal mode" )
   end
   
end
