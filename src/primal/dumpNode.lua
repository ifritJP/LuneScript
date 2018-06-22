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
	    (node.info.mutFlag and "mut" or "") ..
	    node.info.name.txt )
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

   node.info.array:filter( dumpNode, prefix .. "  ", depth + 1 )
   node.info.index:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Return ] = function( node, prefix, depth )
   dump( prefix, depth, node, "" )

   node.info:filter( dumpNode, prefix .. "  ", depth + 1 )
end

dumpNode[ TransUnit.nodeKind.Comment ] = function( node, prefix, depth )
   dump( prefix, depth, node, node.info.txt )
end

dumpNode[ TransUnit.nodeKind.LiteralList ] = function( node, prefix, depth )
   dump( prefix, depth, node, "[]" )

   for index, exp in ipairs( node.info ) do
      exp:filter( dumpNode, prefix .. "  ", depth + 1 )
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


return dumpNode
