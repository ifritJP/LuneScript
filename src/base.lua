-- ajiopjiojio
--[[ fjiaojfeap ]]
local Parser = require( 'primal.Parser' )
local TransUnit = require( 'primal.TransUnit' )

local parser = Parser:create( arg[ 1 ] )

if arg[ 2 ] then
   while true do
      local token = parser:getToken()
      if not token then
	 break
      end
      print( token.kind, token.pos.lineNo, token.pos.column, token.txt )
   end
end

local ast = TransUnit:createAST( parser )

ast:filter( require( 'primal.dumpNode' ), "", 0 )
