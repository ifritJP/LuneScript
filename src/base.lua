local Parser = require( 'primal.Parser' )
local TransUnit = require( 'primal.TransUnit' )

local parser = Parser:create( arg[ 1 ] )

TransUnit:createAST( parser )
   
