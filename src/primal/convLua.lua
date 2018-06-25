local TransUnit = require( 'primal.TransUnit' )

local convLua = {}

convLua.stream = io.stdout
convLua.curLineNo = 1
convLua.indent = 0

function convLua:write( txt )
   if self.needIndent then
      self.stream:write( string.rep( " ", self.indent ) )
      self.needIndent = false
   end

   for cr in string.gmatch( txt, "\n" ) do
      self.curLineNo = self.curLineNo + 1
   end
   self.stream:write( txt )
end

function convLua:setIndent( indent )
   self.indent = indent
end

function convLua:writeln( txt, baseIndent )
   self:write( txt )
   self:write( "\n" )
   self.needIndent = true
   self.indent = baseIndent
end

convLua[ TransUnit.nodeKind.Root ] = function( self, node, baseIndent )
   for index, child in ipairs( node.info.childlen ) do
      child:filter( convLua, baseIndent )
      self:writeln( "", baseIndent )
   end
end
convLua[ TransUnit.nodeKind.Block ] = function( self, node, baseIndent )
   local word = ""
   if node.info.kind == "if" or node.info.kind == "elseif" then
      word = "then"
   elseif node.info.kind == "else" then
      word = ""
   elseif node.info.kind == "while" then
      word = "do"
   elseif node.info.kind == "repeat" then
      word = ""
   elseif node.info.kind == "for" then
      word = "do"
   elseif node.info.kind == "apply" then
      word = "do"
   elseif node.info.kind == "foreach" then
      word = "do"
   elseif node.info.kind == "func" then
      word = ""
   elseif node.info.kind == "{" then
      word = "do"
   end
   self:writeln( word, baseIndent + 4 )
   for index, statement in ipairs( node.info.stmtList ) do
      statement:filter( convLua, baseIndent + 4 )
      self:writeln( "", baseIndent + 4 )
   end

   self:setIndent( baseIndent ) 
   if node.info.kind == "{" then
      self:write( "end", baseIndent )
   end
end

convLua[ TransUnit.nodeKind.StmtExp ] = function( self, node, baseIndent )
   node.info:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.DeclVar ] = function( self, node, baseIndent )
   self:write( "let " )
   
   local varName = ""
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.name.txt )
   end

   self:write( " = " )
   
   if node.info.expList then
      node.info.expList:filter( convLua, baseIndent )
   end
end

convLua[ TransUnit.nodeKind.DeclArg ] = function( self, node, baseIndent )
   self:write( string.format( "%s: ", node.info.name.txt ) )

   node.info.argType:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.DeclArgDDD ] = function( self, node, baseIndent )
   self:write( "..." )
end

convLua[ TransUnit.nodeKind.ExpDDD ] = function( self, node, baseIndent )
   self:write( "..." )
end

convLua[ TransUnit.nodeKind.DeclFunc ] = function( self, node, baseIndent )
   self:write( string.format( "function %s( ", node.info.name.txt ) )
   
   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
      end
      arg:filter( convLua, baseIndent )
   end
   self:write( ")", baseIndent )
   -- for index, refType in ipairs( node.info.retTypeList ) do
   --    if index > 1 then
   -- 	 self:write( ", " )
   --    end
   --    refType:filter( convLua, baseIndent )
   -- end
   node.info.body:filter( convLua, baseIndent )
   self:write( "end" )
end

convLua[ TransUnit.nodeKind.RefType ] = function( self, node, baseIndent )
   self:write(
      (node.info.refFlag and "&" or "") ..
	 (node.info.mutFlag and "mut " or "") .. node.info.name.txt )
   if node.info.array == "array" then
      self:write( "[@]" )
   elseif node.info.array == "list" then
      self:write( "[]" )
   end
end

convLua[ TransUnit.nodeKind.If ] = function( self, node, baseIndent )

   for index, val in ipairs( node.info ) do
      if index == 1 then
	 self:write( "if " )
	 val.exp:filter( convLua, baseIndent )
      elseif index ~= #node.info then
	 self:write( "elseif " )
	 val.exp:filter( convLua, baseIndent )
      else
	 self:write( "else" )
      end
      self:write( " " )
      val.block:filter( convLua, baseIndent )
   end
   self:write( "end" )
end

convLua[ TransUnit.nodeKind.While ] = function( self, node, baseIndent )
   self:write( "while " )

   node.info.exp:filter( convLua, baseIndent )
   self:write( " " )
   node.info.block:filter( convLua, baseIndent )
   self:write( "end" )
end

convLua[ TransUnit.nodeKind.Repeat ] = function( self, node, baseIndent )
   self:write( "repeat " )
   node.info.block:filter( convLua, baseIndent )
   self:write( "until " )
   node.info.exp:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.For ] = function( self, node, baseIndent )
   self:write( string.format( "for %s = ", node.info.val.txt ) )
   node.info.init:filter( convLua, baseIndent )
   self:write( ", " )
   node.info.to:filter( convLua, baseIndent )
   if node.info.delta then
      self:write( ", " )
      node.info.delta:filter( convLua, baseIndent )
   end
   self:write( " " )
   node.info.block:filter( convLua, baseIndent )
   self:write( "end" )
end

convLua[ TransUnit.nodeKind.Apply ] = function( self, node, baseIndent )
   self:write( "for " )
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.txt )
   end
   self:write( " in " )
   node.info.exp:filter( convLua, baseIndent )
   self:write( " " )
   node.info.block:filter( convLua, baseIndent )
   self:write( "end" )
end

convLua[ TransUnit.nodeKind.Foreach ] = function( self, node, baseIndent )
   self:write( "foreach " )
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.txt )
   end
   self:write( " in " )
   node.info.exp:filter( convLua, baseIndent )
   self:write( " " )
   node.info.block:filter( convLua, baseIndent )
   self:write( "end" )
end


convLua[ TransUnit.nodeKind.ExpCall ] = function( self, node, baseIndent )
   node.info.func:filter( convLua, baseIndent )
   self:write( "(" )
   node.info.argList:filter( convLua, baseIndent )
   self:write( ")" )
end



convLua[ TransUnit.nodeKind.ExpList ] = function( self, node, baseIndent )
   for index, exp in ipairs( node.info ) do
      if index > 1 then
	 self:write( ", " )
      end
      exp:filter( convLua, baseIndent )
   end
end

convLua[ TransUnit.nodeKind.ExpOp1 ] = function( self, node, baseIndent )
   self:write( node.info.op.txt )
   node.info.exp:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.ExpCast ] = function( self, node, baseIndent )
   node.info.exp:filter( convLua, baseIndent )
   -- node.info.castType:filter( convLua, baseIndent )
end


convLua[ TransUnit.nodeKind.ExpOp2 ] = function( self, node, baseIndent )
   node.info.exp1:filter( convLua, baseIndent )

   self:write( node.info.op.txt )
   
   node.info.exp2:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.ExpRef ] = function( self, node, baseIndent )
   self:write( node.info.txt )
end

convLua[ TransUnit.nodeKind.ExpRefItem ] = function( self, node, baseIndent )
   node.info.val:filter( convLua, baseIndent )
   self:write( "[" )
   node.info.index:filter( convLua, baseIndent )
   self:write( "]" )
end

convLua[ TransUnit.nodeKind.RefField ] = function( self, node, baseIndent )
   node.info.prefix:filter( convLua, baseIndent )
   self:write( "." .. node.info.field.txt )
end

convLua[ TransUnit.nodeKind.Return ] = function( self, node, baseIndent )
   self:write( "return " )

   node.info:filter( convLua, baseIndent )
end

convLua[ TransUnit.nodeKind.LiteralList ] = function( self, node, baseIndent )
   self:write( "{" )

   node.info:filter( convLua, baseIndent )

   self:write( "}" )
end

convLua[ TransUnit.nodeKind.LiteralMap ] = function( self, node, baseIndent )
   self:write( "{" )
   local index = 1
   for key, val in pairs( node.info ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( "[" )
      key:filter( convLua, baseIndent )
      self:write( "] = " )
      val:filter( convLua, baseIndent )
      index = index + 1
   end

   self:write( "}" )
end


convLua[ TransUnit.nodeKind.LiteralArray ] = function( self, node, baseIndent )
   self:write( "{" )

   node.info:filter( convLua, baseIndent )

   self:write( "}" )
end


convLua[ TransUnit.nodeKind.LiteralChar ] = function( self, node, baseIndent )
   self:write( string.format( "%g", node.info.num ) )
end

convLua[ TransUnit.nodeKind.LiteralNum ] = function( self, node, baseIndent )
   self:write( string.format( "%g", node.info.num ) )
end

convLua[ TransUnit.nodeKind.LiteralString ] = function( self, node, baseIndent )
   local txt = node.info.txt
   if string.find( txt, '"""', 1, true ) then
      txt = "[==[" .. txt:sub( 3, -3 ) .. "]==]"
   end
   self:write( txt )
end

convLua[ TransUnit.nodeKind.LiteralBool ] = function( self, node, baseIndent )
   self:write( node.info.txt )
end

convLua[ TransUnit.nodeKind.Break ] = function( self, node, baseIndent )
   self:write( "break" )

end


return convLua
