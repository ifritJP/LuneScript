local TransUnit = require( 'primal.TransUnit' )
local Parser = require( 'primal.Parser' )

local filterObj = {}

filterObj.stream = io.stdout
filterObj.curLineNo = 1
filterObj.indent = 0

local stepIndent = 2

local builtInModuleSet = {}
builtInModuleSet[ "io" ] = true
builtInModuleSet[ "string" ] = true
builtInModuleSet[ "table" ] = true
builtInModuleSet[ "math" ] = true


function filterObj:new( stream )
   self.stream = stream
   return self
end

function filterObj:write( txt )
   if self.needIndent then
      self.stream:write( string.rep( " ", self.indent ) )
      self.needIndent = false
   end

   for cr in string.gmatch( txt, "\n" ) do
      self.curLineNo = self.curLineNo + 1
   end
   self.stream:write( txt )
end

function filterObj:setIndent( indent )
   self.indent = indent
end

function filterObj:writeln( txt, baseIndent )
   self:write( txt )
   self:write( "\n" )
   self.needIndent = true
   self.indent = baseIndent
end

filterObj[ TransUnit.nodeKind.None ] = function( self, node, parent, baseIndent )
   self:writeln( "-- none", baseIndent );
end


filterObj[ TransUnit.nodeKind.Root ] = function( self, node, parent, baseIndent )
   self:writeln( "local moduleObj = {}", baseIndent )
   
   for index, child in ipairs( node.info.childlen ) do
      child:filter( filterObj, node, baseIndent )
      self:writeln( "", baseIndent )
   end

   self:writeln( "return moduleObj", baseIndent )
end
filterObj[ TransUnit.nodeKind.Block ] = function( self, node, parent, baseIndent )
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
   self:writeln( word, baseIndent + stepIndent )
   for index, statement in ipairs( node.info.stmtList ) do
      statement:filter( filterObj, node, baseIndent + stepIndent )
      self:writeln( "", baseIndent + stepIndent )
   end

   self:setIndent( baseIndent ) 
   if node.info.kind == "{" then
      self:write( "end", baseIndent )
   end
end

filterObj[ TransUnit.nodeKind.StmtExp ] = function( self, node, parent, baseIndent )
   node.info:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.DeclClass ] = function( self, node, parent, baseIndent )
   local className = node.info.name.txt
   self:writeln( string.format( "local %s = {}", className ), baseIndent )
   self:writeln( string.format( "moduleObj.%s = %s", className, className ),
		 baseIndent )
   for index, field in ipairs( node.info.fieldList ) do
      field:filter( filterObj, node, baseIndent )
   end
end

filterObj[ TransUnit.nodeKind.DeclMember ] = function( self, node, parent, baseIndent )
   -- dump( baseIndent, node, node.info.name.txt )
   -- node.info.refType:filter( filterObj, prefix .. "  ", depth + 1 )
end


filterObj[ TransUnit.nodeKind.DeclConstr ] = function( self, node, parent, baseIndent )
   local className = node.info.className.txt
   self:write( string.format( "function %s.new( ", className ) )

   local argTxt = ""
   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
	 argTxt = argTxt .. ", "
      end
      arg:filter( filterObj, node, baseIndent )
      argTxt = argTxt .. arg.info.name.txt
   end
   self:writeln( ")", baseIndent + stepIndent )
   self:writeln( "local obj = {}", baseIndent + stepIndent )
   self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ),
		 baseIndent + stepIndent )
   self:writeln( string.format( "return obj.__init and obj:__init( %s ) or nil;",
				argTxt ), baseIndent )
   self:writeln( "end", baseIndent )

   
   -- for index, refType in ipairs( node.info.retTypeList ) do
   --    if index > 1 then
   -- 	 self:write( ", " )
   --    end
   --    refType:filter( filterObj, node, baseIndent )
   -- end
   self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
   node.info.body:filter( filterObj, node, baseIndent )
   self:writeln( "end", baseIndent )
end


filterObj[ TransUnit.nodeKind.DeclMethod ] = function( self, node, parent, baseIndent )
   local delimit = ":"
   if node.info.staticFlag then
      delimit = "."
   end
   self:write( string.format( "function %s%s%s( ",
			      node.info.className.txt,
			      delimit, node.info.name.txt ) )

   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
      end
      arg:filter( filterObj, node, baseIndent )
   end
   self:write( ")", baseIndent )
   -- for index, refType in ipairs( node.info.retTypeList ) do
   --    if index > 1 then
   -- 	 self:write( ", " )
   --    end
   --    refType:filter( filterObj, node, baseIndent )
   -- end
   node.info.body:filter( filterObj, node, baseIndent )
   self:writeln( "end", baseIndent )
end


filterObj[ TransUnit.nodeKind.DeclVar ] = function( self, node, parent, baseIndent )
   self:write( "local " )
   
   local varName = ""
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.name.txt )
   end

   self:write( " = " )
   
   if node.info.expList then
      node.info.expList:filter( filterObj, node, baseIndent )
   end
end

filterObj[ TransUnit.nodeKind.DeclArg ] = function( self, node, parent, baseIndent )
   self:write( string.format( "%s ", node.info.name.txt ) )

   -- node.info.argType:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.DeclArgDDD ] = function( self, node, parent, baseIndent )
   self:write( "..." )
end

filterObj[ TransUnit.nodeKind.ExpDDD ] = function( self, node, parent, baseIndent )
   self:write( "..." )
end

filterObj[ TransUnit.nodeKind.DeclFunc ] = function( self, node, parent, baseIndent )
   local name = node.info.name
   self:write( string.format( "function %s( ", name and name.txt or "" ) )
   
   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
      end
      arg:filter( filterObj, node, baseIndent )
   end
   self:write( ")", baseIndent )
   -- for index, refType in ipairs( node.info.retTypeList ) do
   --    if index > 1 then
   -- 	 self:write( ", " )
   --    end
   --    refType:filter( filterObj, node, baseIndent )
   -- end
   node.info.body:filter( filterObj, node, baseIndent )
   self:write( "end" )
end

filterObj[ TransUnit.nodeKind.RefType ] = function( self, node, parent, baseIndent )
   self:write(
      (node.info.refFlag and "&" or "") ..
	 (node.info.mutFlag and "mut " or "") .. node.info.name.txt )
   if node.info.array == "array" then
      self:write( "[@]" )
   elseif node.info.array == "list" then
      self:write( "[]" )
   end
end

filterObj[ TransUnit.nodeKind.If ] = function( self, node, parent, baseIndent )

   for index, val in ipairs( node.info ) do
      if index == 1 then
	 self:write( "if " )
	 val.exp:filter( filterObj, node, baseIndent )
      elseif val.kind == "elseif" then
	 self:write( "elseif " )
	 val.exp:filter( filterObj, node, baseIndent )
      else
	 self:write( "else" )
      end
      self:write( " " )
      val.block:filter( filterObj, node, baseIndent )
   end
   self:write( "end" )
end

filterObj[ TransUnit.nodeKind.While ] = function( self, node, parent, baseIndent )
   self:write( "while " )

   node.info.exp:filter( filterObj, node, baseIndent )
   self:write( " " )
   node.info.block:filter( filterObj, node, baseIndent )
   self:write( "end" )
end

filterObj[ TransUnit.nodeKind.Repeat ] = function( self, node, parent, baseIndent )
   self:write( "repeat " )
   node.info.block:filter( filterObj, node, baseIndent )
   self:write( "until " )
   node.info.exp:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.For ] = function( self, node, parent, baseIndent )
   self:write( string.format( "for %s = ", node.info.val.txt ) )
   node.info.init:filter( filterObj, node, baseIndent )
   self:write( ", " )
   node.info.to:filter( filterObj, node, baseIndent )
   if node.info.delta then
      self:write( ", " )
      node.info.delta:filter( filterObj, node, baseIndent )
   end
   self:write( " " )
   node.info.block:filter( filterObj, node, baseIndent )
   self:write( "end" )
end

filterObj[ TransUnit.nodeKind.Apply ] = function( self, node, parent, baseIndent )
   self:write( "for " )
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.txt )
   end
   self:write( " in " )
   node.info.exp:filter( filterObj, node, baseIndent )
   self:write( " " )
   node.info.block:filter( filterObj, node, baseIndent )
   self:write( "end" )
end

filterObj[ TransUnit.nodeKind.Foreach ] = function( self, node, parent, baseIndent )
   self:write( "for " )
   self:write( node.info.key and node.info.key.txt or "__index" )
   self:write( ", " )
   self:write( node.info.val.txt )

   self:write( " in pairs( " )
   node.info.exp:filter( filterObj, node, baseIndent )
   self:write( " ) " )
   node.info.block:filter( filterObj, node, baseIndent )
   self:write( "end" )
end


filterObj[ TransUnit.nodeKind.ExpCall ] = function( self, node, parent, baseIndent )
   node.info.func:filter( filterObj, node, baseIndent )
   self:write( "(" )
   if node.info.argList then
      node.info.argList:filter( filterObj, node, baseIndent )
   end
   self:write( ")" )
end



filterObj[ TransUnit.nodeKind.ExpList ] = function( self, node, parent, baseIndent )
   for index, exp in ipairs( node.info ) do
      if index > 1 then
	 self:write( ", " )
      end
      exp:filter( filterObj, node, baseIndent )
   end
end

filterObj[ TransUnit.nodeKind.ExpOp1 ] = function( self, node, parent, baseIndent )
   local op = node.info.op.txt
   if op == "not" then
      op = op .. " "
   end
   self:write( op )
   node.info.exp:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.ExpCast ] = function( self, node, parent, baseIndent )
   node.info.exp:filter( filterObj, node, baseIndent )
   -- node.info.castType:filter( filterObj, node, baseIndent )
end


filterObj[ TransUnit.nodeKind.ExpParen ] = function( self, node, prefix, depth )
   self:write( "(" )
   node.info:filter( filterObj, node, baseIndent )
   self:write( ")" )
end

filterObj[ TransUnit.nodeKind.ExpOp2 ] = function( self, node, parent, baseIndent )
   node.info.exp1:filter( filterObj, node, baseIndent )

   self:write( " " .. node.info.op.txt .. " " )
   
   node.info.exp2:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.ExpRef ] = function( self, node, parent, baseIndent )
   self:write( node.info.txt )
end

filterObj[ TransUnit.nodeKind.ExpRefItem ] = function( self, node, parent, baseIndent )
   if node.info.val.kind == TransUnit.nodeKind.LiteralString then
      self:write( "string.byte( " )
      node.info.val:filter( filterObj, node, baseIndent )
      self:write( ", " )
      node.info.index:filter( filterObj, node, baseIndent )
      self:write( " )" )
   else
      node.info.val:filter( filterObj, node, baseIndent )
      self:write( "[" )
      node.info.index:filter( filterObj, node, baseIndent )
      self:write( "]" )
   end
end

filterObj[ TransUnit.nodeKind.RefField ] = function( self, node, parent, baseIndent )
   node.info.prefix:filter( filterObj, node, baseIndent )
   local delimit = "."
   if parent.kind == TransUnit.nodeKind.ExpCall then
      if node.info.prefix.kind == TransUnit.nodeKind.ExpRef and
	 builtInModuleSet[ node.info.prefix.info.txt ]
      then
	 delimit = "."
      else
	 delimit = ":"
      end
   end
   self:write( delimit .. node.info.field.txt )
end

filterObj[ TransUnit.nodeKind.Return ] = function( self, node, parent, baseIndent )
   self:write( "return " )

   node.info:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.LiteralList ] = function( self, node, parent, baseIndent )
   self:write( "{" )

   node.info:filter( filterObj, node, baseIndent )

   self:write( "}" )
end

filterObj[ TransUnit.nodeKind.LiteralMap ] = function( self, node, parent, baseIndent )
   self:write( "{" )
   local index = 1
   for key, val in pairs( node.info ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( "[" )
      key:filter( filterObj, node, baseIndent )
      self:write( "] = " )
      val:filter( filterObj, node, baseIndent )
      index = index + 1
   end

   self:write( "}" )
end


filterObj[ TransUnit.nodeKind.LiteralArray ] = function( self, node, parent, baseIndent )
   self:write( "{" )

   node.info:filter( filterObj, node, baseIndent )

   self:write( "}" )
end


filterObj[ TransUnit.nodeKind.LiteralChar ] = function( self, node, parent, baseIndent )
   self:write( string.format( "%g", node.info.num ) )
end

filterObj[ TransUnit.nodeKind.LiteralInt ] = function( self, node, parent, baseIndent )
   self:write( string.format( "%d", node.info.num ) )
end

filterObj[ TransUnit.nodeKind.LiteralReal ] = function( self, node, parent, baseIndent )
   self:write( string.format( "%s", node.info.num ) )
end

filterObj[ TransUnit.nodeKind.LiteralString ] = function( self, node, parent, baseIndent )
   local txt = node.info.token.txt
   if string.find( txt, '^```' ) then
      txt = "[==[" .. txt:sub( 4, -4 ) .. "]==]"
   end
   if #node.info.argList > 0 then
      self:write( string.format( "string.format( %s, ", txt ) )
      for index, val in ipairs( node.info.argList ) do
	 if index > 1 then
	    self:write( ", " )
	 end
	 val:filter( filterObj, node, baseIndent )
      end
      self:write( ")" )
   else
      self:write( txt )
   end
end

filterObj[ TransUnit.nodeKind.LiteralBool ] = function( self, node, parent, baseIndent )
   self:write( node.info.txt )
end

filterObj[ TransUnit.nodeKind.LiteralNil ] = function( self, node, prefix, depth )
   self:write( "nil" )
end

filterObj[ TransUnit.nodeKind.Break ] = function( self, node, parent, baseIndent )
   self:write( "break" )

end


return filterObj
