local TransUnit = require( 'primal.TransUnit' )

local filterObj = {}

function dump( prefix, depth, node, txt )
   print( string.format( "%s: %s %s",
			 prefix, TransUnit:getNodeKindName( node.kind ), txt ) )
end

filterObj[ TransUnit.nodeKind.None ] = function( self, node, prefix, depth )
   dump( prefix, index, node, "" )
end

filterObj[ TransUnit.nodeKind.Import ] = function( self, node, prefix, depth )
   dump( prefix, index, node, node.info )
end

filterObj[ TransUnit.nodeKind.Root ] = function( self, node, prefix, depth )
   dump( prefix, index, node, "" )
   for index, child in ipairs( node.info.childlen ) do
      child:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.Block ] = function( self, node, prefix, depth )
   dump( prefix, index, node, "" )
   for index, statement in ipairs( node.info.stmtList ) do
      statement:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.StmtExp ] = function( self, node, prefix, depth )
   dump( prefix, index, node, "" )
   node.info:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.DeclClass ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )

   for index, field in ipairs( node.info.fieldList ) do
      field:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.DeclMember ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )
   node.info.refType:filter( filterObj, prefix .. "  ", depth + 1 )
end


filterObj[ TransUnit.nodeKind.DeclVar ] = function( self, node, prefix, depth )
   local varName = ""
   for index, var in ipairs( node.info.varList ) do
      varName = varName .. " " .. var.name.txt
   end
   dump( prefix, depth, node, varName )

   if node.info.expList then
      node.info.expList:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.DeclArg ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )

   node.info.argType:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.DeclArgDDD ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "..." )
end

filterObj[ TransUnit.nodeKind.ExpDDD ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "..." )
end

filterObj[ TransUnit.nodeKind.DeclFunc ] = function( self, node, prefix, depth )
   local name = node.info.name
   dump( prefix, depth, node, name and name.txt or "<anonymous>" )
   for index, arg in ipairs( node.info.argList ) do
      arg:filter( filterObj, prefix .. "  ", depth + 1 )
   end
   for index, refType in ipairs( node.info.retTypeList ) do
      refType:filter( filterObj, prefix .. "  ", depth + 1 )
   end
   node.info.body:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.DeclMethod ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )
   for index, arg in ipairs( node.info.argList ) do
      arg:filter( filterObj, prefix .. "  ", depth + 1 )
   end
   for index, refType in ipairs( node.info.retTypeList ) do
      refType:filter( filterObj, prefix .. "  ", depth + 1 )
   end
   node.info.body:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.DeclConstr ] = function( self, node, prefix, depth )
   filterObj[ TransUnit.nodeKind.DeclMethod ]( self, node, prefix, depth )
end

filterObj[ TransUnit.nodeKind.RefType ] = function( self, node, prefix, depth )
   dump( prefix, depth, node,
	 (node.info.refFlag and "&" or "") ..
	    (node.info.mutFlag and "mut " or "") ..
	    node.info.name.txt )
end

filterObj[ TransUnit.nodeKind.If ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "")

   for index, val in ipairs( node.info ) do
      print( prefix .. val.kind )
      if val.exp then
	 val.exp:filter( filterObj, prefix .. "  ", depth + 1 )
      end
      val.block:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.While ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "")

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.Repeat ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "")

   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.For ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.val.txt )

   node.info.init:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.to:filter( filterObj, prefix .. "  ", depth + 1 )
   if node.info.delta then
      node.info.delta:filter( filterObj, prefix .. "  ", depth + 1 )
   end
   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.Apply ] = function( self, node, prefix, depth )
   local varNames = ""
   for index, var in ipairs( node.info.varList ) do
      varNames = varNames .. var.txt .. " "
   end
   dump( prefix, depth, node, varNames )

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.Foreach ] = function( self, node, prefix, depth )
   local index = node.info.key and node.info.key.txt or ""
   dump( prefix, depth, node, node.info.val.txt .. " " .. index )

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.Forsort ] = function( self, node, prefix, depth )
   local index = node.info.key and node.info.key.txt or ""
   dump( prefix, depth, node, node.info.val.txt .. " " .. index )

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.block:filter( filterObj, prefix .. "  ", depth + 1 )
end


filterObj[ TransUnit.nodeKind.ExpCall ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info.func:filter( filterObj, prefix .. "  ", depth + 1 )
   if node.info.argList then
      node.info.argList:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end



filterObj[ TransUnit.nodeKind.ExpList ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   for index, exp in ipairs( node.info ) do
      exp:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.ExpOp1 ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.op.txt )

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.ExpCast ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info.exp:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.castType:filter( filterObj, prefix .. "  ", depth + 1 )
end


filterObj[ TransUnit.nodeKind.ExpParen ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "()" )

   node.info:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.ExpOp2 ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.op.txt )

   node.info.exp1:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.exp2:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.ExpNew ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info.symbol:filter( filterObj, prefix .. "  ", depth + 1 )
   if node.info.argList then
      node.info.argList:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end

filterObj[ TransUnit.nodeKind.ExpRef ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.txt )
end

filterObj[ TransUnit.nodeKind.ExpRefItem ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "seq[exp]" )

   node.info.val:filter( filterObj, prefix .. "  ", depth + 1 )
   node.info.index:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.RefField ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.field.txt )

   node.info.prefix:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.Return ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.LiteralList ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "[]" )

   node.info:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.LiteralMap ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )

   for index, pair in ipairs( node.info.pairList ) do
      pair.key:filter( filterObj, prefix .. "  ", depth + 1 )
      pair.val:filter( filterObj, prefix .. "  ", depth + 1 )
   end
end


filterObj[ TransUnit.nodeKind.LiteralArray ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "[@]" )

   node.info:filter( filterObj, prefix .. "  ", depth + 1 )
end


filterObj[ TransUnit.nodeKind.LiteralChar ] = function( self, node, prefix, depth )
   dump( prefix, depth, node,
	 string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[ TransUnit.nodeKind.LiteralInt ] = function( self, node, prefix, depth )
   dump( prefix, depth, node,
	 string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[ TransUnit.nodeKind.LiteralReal ] = function( self, node, prefix, depth )
   dump( prefix, depth, node,
	 string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

filterObj[ TransUnit.nodeKind.LiteralString ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.token.txt )
end

filterObj[ TransUnit.nodeKind.LiteralBool ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, node.info.txt == "true" )
end

filterObj[ TransUnit.nodeKind.LiteralNil ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )
end

filterObj[ TransUnit.nodeKind.Break ] = function( self, node, prefix, depth )
   dump( prefix, depth, node, "" )
end


return filterObj
