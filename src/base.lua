-- ajiopjiojio
--[[ fjiaojfeap ]]
local Parser = require( 'primal.Parser' )

local parser = Parser:create( arg[ 1 ] )

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
      ast:filter( require( 'primal.dumpNode' ), "", 0 )
   elseif mode == "lua" then
      ast:filter( require( 'primal.convLua' ), 0 )
   else
      print( "illegal mode" )
   end
   
end
