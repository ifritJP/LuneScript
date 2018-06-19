local Parser = require( 'primal.Parser' )


local TransUnit = {}

function TransUnit:createAST( parser )
   print( "kind", "line", "column", "txt" )
   while true do
      local token = parser:getToken()
      if not token then
	 break
      end
      print( Parser.getKindTxt( token.kind ), token.lineNo, token.column, token.txt )
   end
end

return TransUnit
