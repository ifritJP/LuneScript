local TransUnit = require( 'primal.TransUnit' )

local dumpNode = {}

function dump( prefix, depth, node, txt )
   print( string.format( "%s: %s %s",
			 prefix, TransUnit:getNodeKindName( node.kind ), txt ) )
end

dumpNode[ TransUnit.nodeKind.Root ] = function( node, prefix, depth )
   dump( prefix, index, node, "" )
   for index, child in ipairs( node.info.childlen ) do
      child:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.Block ] = function( node, prefix, depth )
   dump( prefix, index, node, "" )
   for index, statement in ipairs( node.info ) do
      statement:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.StmtExp ] = function( node, prefix, depth )
   dump( prefix, index, node, "" )
   node.info:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.DeclVar ] = function( node, prefix, depth )
   local varName = ""
   for index, var in ipairs( node.info.varList ) do
      varName = varName .. " " .. var.name.txt
   end
   dump( prefix, depth, node, varName )

   if node.info.expList then
      node.info.expList:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.DeclArg ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )

   node.info.argType:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.DeclArgDDD ] = function( node, prefix, depth )
   dump( prefix, depth, node, "..." )
end

dumpNode[ TransUnit.nodeKind.ExpDDD ] = function( node, prefix, depth )
   dump( prefix, depth, node, "..." )
end

dumpNode[ TransUnit.nodeKind.DeclFunc ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.name.txt )
   for index, arg in ipairs( node.info.argList ) do
      arg:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
   for index, refType in ipairs( node.info.retTypeList ) do
      refType:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
   node.info.body:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.RefType ] = function( node, prefix, depth )
   dump( prefix, depth, node,
	 (node.info.refFlag and "&" or "") ..
	    (node.info.mutFlag and "mut " or "") ..
	    node.info.name.txt )
end

dumpNode[ TransUnit.nodeKind.If ] = function( node, prefix, depth )
   dump( prefix, depth, node, "")

   for index, val in ipairs( node.info ) do
      print( prefix .. val.kind )
      if val.exp then
	 val.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
      end
      val.block:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.While ] = function( node, prefix, depth )
   dump( prefix, depth, node, "")

   node.info.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.block:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Repeat ] = function( node, prefix, depth )
   dump( prefix, depth, node, "")

   node.info.block:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.For ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.txt )

   node.info.init:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.to:filter( dumpNode, prefix .. "  ", depth + 1 )
   if node.info.delta then
      node.info.delta:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
   node.info.block:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Apply ] = function( node, prefix, depth )
   local varNames = ""
   for index, var in ipairs( node.info.varList ) do
      varNames = varNames .. var.txt .. " "
   end
   dump( prefix, depth, node, varNames )

   node.info.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.block:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Foreach ] = function( node, prefix, depth )
   local varNames = ""
   for index, var in ipairs( node.info.varList ) do
      varNames = varNames .. var.txt .. " "
   end
   dump( prefix, depth, node, varNames )

   node.info.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.block:filter( dumpNode, prefix .. "  ", depth + 1 )
end


dumpNode[ TransUnit.nodeKind.ExpCall ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info.func:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.argList:filter( dumpNode, prefix .. "  ", depth + 1 )
end



dumpNode[ TransUnit.nodeKind.ExpList ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )

   for index, exp in ipairs( node.info ) do
      exp:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.ExpOp1 ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.op.txt )

   node.info.exp:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.ExpOp2 ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.op.txt )

   node.info.exp1:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.exp2:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.ExpRef ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.txt )
end

dumpNode[ TransUnit.nodeKind.ExpRefItem ] = function( node, prefix, depth )
   dump( prefix, depth, node, "seq[exp]" )

   node.info.val:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.index:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.RefField ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.field.txt )

   node.info.prefix:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Return ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.LiteralList ] = function( node, prefix, depth )
   dump( prefix, depth, node, "[]" )

   for index, exp in ipairs( node.info ) do
      exp:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end

dumpNode[ TransUnit.nodeKind.LiteralMap ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )

   for key, val in pairs( node.info ) do
      key:filter( dumpNode, prefix .. "  ", depth + 1 )
      val:filter( dumpNode, prefix .. "  ", depth + 1 )
   end
end


dumpNode[ TransUnit.nodeKind.LiteralArray ] = function( node, prefix, depth )
   dump( prefix, depth, node, "[@]" )

   node.info:filter( dumpNode, prefix .. "  ", depth + 1 )
end


dumpNode[ TransUnit.nodeKind.LiteralChar ] = function( node, prefix, depth )
   dump( prefix, depth, node,
	 string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

dumpNode[ TransUnit.nodeKind.LiteralNum ] = function( node, prefix, depth )
   dump( prefix, depth, node,
	 string.format( "%s(%s)", node.info.num, node.info.token.txt ) )
end

dumpNode[ TransUnit.nodeKind.LiteralString ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.txt )
end

dumpNode[ TransUnit.nodeKind.LiteralBool ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.txt == "true" )
end

dumpNode[ TransUnit.nodeKind.Break ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )
end


return dumpNode
