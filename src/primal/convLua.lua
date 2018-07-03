local TransUnit = require( 'primal.TransUnit' )
local Parser = require( 'primal.Parser' )

local filterObj = {}

filterObj.curLineNo = 1
filterObj.indent = 0

local stepIndent = 2

local builtInModuleSet = {}
builtInModuleSet[ "io" ] = true
builtInModuleSet[ "string" ] = true
builtInModuleSet[ "table" ] = true
builtInModuleSet[ "math" ] = true
builtInModuleSet[ "debug" ] = true
builtInModuleSet[ "_luneScript" ] = true


-- クラス名 → クラス情報
local className2InfoMap = {}
-- public 変数名 → 変数情報
local pubVarName2InfoMap = {}

function filterObj:new( streamName, stream )
   self.streamName = streamName;
   self.stream = stream
   self.moduleName2Info = {}
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

filterObj[ TransUnit.nodeKind.Import ] = function( self, node, parent, baseIndent )
   local path = node.info
   local moduleName = string.gsub( path, ".*%.", "" )
   local moduleInfo = require( path )
   self.moduleName2Info[ moduleName ] = moduleInfo;
   self:writeln( string.format( "local %s = require( '%s' )", moduleName, node.info ),
		 baseIndent );
end

filterObj[ TransUnit.nodeKind.Root ] = function( self, node, parent, baseIndent )
   self:writeln( "--" .. filterObj.streamName, baseIndent )
   self:writeln( "local moduleObj = {}", baseIndent )
   
   for index, child in ipairs( node.info.childlen ) do
      child:filter( filterObj, node, baseIndent )
      self:writeln( "", baseIndent )
   end


   self:writeln( "local _className2InfoMap = {}", baseIndent )
   self:writeln( "moduleObj._className2InfoMap = _className2InfoMap", baseIndent )

   local keyList = {}
   for className, classInfo in pairs( className2InfoMap ) do
      table.insert( keyList, className )
   end
   table.sort( keyList )

   for index, className in ipairs( keyList ) do
      local classInfo = className2InfoMap[ className ];
      self:writeln( string.format( "local _classInfo%s = {}", className), baseIndent )
      self:writeln( string.format( "_className2InfoMap.%s = _classInfo%s",
				   className, className), baseIndent )

      local keyList2 = {}
      for methodName, methodInfo in pairs( classInfo ) do
	 table.insert( keyList2, methodName )
      end
      table.sort( keyList2 )
      
      for index2, methodName in ipairs( keyList2 ) do
	 local methodInfo = classInfo[ methodName ]
	 self:writeln( string.format( "_classInfo%s.%s = {", className, methodName ),
		       baseIndent )
	 self:writeln(
	    string.format( "  name='%s', staticFlag = %s, accessMode = '%s' }",
			   methodName, methodInfo.staticFlag, methodInfo.accessMode ),
	    baseIndent )
      end
   end

   self:writeln( "local _varName2InfoMap = {}", baseIndent )
   self:writeln( "moduleObj._varName2InfoMap = _varName2InfoMap", baseIndent )
   
   local keyList3 = {}
   for varName, varInfo in pairs( pubVarName2InfoMap ) do
      table.insert( keyList3, varName )
   end
   table.sort( keyList3 )
   for index, varName in ipairs( keyList3 ) do
      local varInfo = pubVarName2InfoMap[ varName ]
      self:writeln( string.format( "_varName2InfoMap.%s = {", varName ),
		    baseIndent )
      self:writeln(
	 string.format( "  name='%s', accessMode = '%s' }",
			varName, varInfo.accessMode ),
	 baseIndent )
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
   local classInfo = {}
   local className = node.info.name.txt
   className2InfoMap[ className ] = classInfo
   
   self:writeln( string.format( "local %s = {}", className ), baseIndent )
   if node.info.accessMode == "pub" then
      self:writeln( string.format( "moduleObj.%s = %s", className, className ),
		    baseIndent )
   end
   for index, field in ipairs( node.info.fieldList ) do
      field:filter( filterObj, node, baseIndent )
   end
end

filterObj[ TransUnit.nodeKind.DeclMember ] = function( self, node, parent, baseIndent )
   -- dump( baseIndent, node, node.info.name.txt )
   -- node.info.refType:filter( filterObj, prefix .. "  ", depth + 1 )
end

filterObj[ TransUnit.nodeKind.ExpNew ] = function( self, node, parent, baseIndent )
   node.info.symbol:filter( filterObj, node, baseIndent )
   self:write( ".new(" )
   if node.info.argList then
      node.info.argList:filter( filterObj, node, baseIndent )
   end
   self:write( ")" )
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
   self:writeln( " )", baseIndent + stepIndent )
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

   local classInfo = className2InfoMap[ node.info.className.txt ]
   
   local delimit = ":"
   if node.info.staticFlag then
      delimit = "."
   end
   local methodName = node.info.name.txt
   self:write( string.format( "function %s%s%s( ",
			      node.info.className.txt,
			      delimit, methodName ) )

   classInfo[ methodName ] = {
      funcFlag = true,
      staticFlag = node.info.staticFlag, accessMode = node.info.accessMode
   }
   

   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
      end
      arg:filter( filterObj, node, baseIndent )
   end
   self:write( " )", baseIndent )
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

   if node.info.accessMode ~= "global" then
      self:write( "local " )
   end
   
   local varName = ""
   for index, var in ipairs( node.info.varList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( var.name.txt )
   end

   
   if node.info.expList then
      self:write( " = " )
      node.info.expList:filter( filterObj, node, baseIndent )
   end

   if node.info.accessMode == "pub" then
      self:writeln( "", baseIndent )
      for index, var in ipairs( node.info.varList ) do
	 self:writeln( string.format( "moduleObj.%s = %s", var.name.txt, var.name.txt ),
		       baseIndent )
	 pubVarName2InfoMap[ var.name.txt ] =  {
	    funcFlag = false,
	    staticFlag = node.info.staticFlag, accessMode = node.info.accessMode
	 }
      end

   end
end

filterObj[ TransUnit.nodeKind.DeclArg ] = function( self, node, parent, baseIndent )
   self:write( string.format( "%s", node.info.name.txt ) )

   -- node.info.argType:filter( filterObj, node, baseIndent )
end

filterObj[ TransUnit.nodeKind.DeclArgDDD ] = function( self, node, parent, baseIndent )
   self:write( "..." )
end

filterObj[ TransUnit.nodeKind.ExpDDD ] = function( self, node, parent, baseIndent )
   self:write( "..." )
end

filterObj[ TransUnit.nodeKind.DeclFunc ] = function( self, node, parent, baseIndent )
   local nameToken = node.info.name
   local name = nameToken and nameToken.txt or ""
   local localTxt = ""
   if node.info.accessMode ~= "global" and #name ~= 0 then
      localTxt = "local "
   end
   self:write( string.format( "%sfunction %s( ", localTxt, name ) )
   
   for index, arg in ipairs( node.info.argList ) do
      if index > 1 then
	 self:write( ", " )
      end
      arg:filter( filterObj, node, baseIndent )
   end
   self:write( " )", baseIndent )
   -- for index, refType in ipairs( node.info.retTypeList ) do
   --    if index > 1 then
   -- 	 self:write( ", " )
   --    end
   --    refType:filter( filterObj, node, baseIndent )
   -- end
   node.info.body:filter( filterObj, node, baseIndent )

   self:writeln( "end", baseIndent )

   if node.info.accessMode == "pub" then
      self:write( string.format( "moduleObj.%s = %s", name, name ) )
   end
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


filterObj[ TransUnit.nodeKind.Forsort ] = function( self, node, parent, baseIndent )
   self:writeln( "do", baseIndent + stepIndent );
   self:writeln( "local __sorted = {}", baseIndent + stepIndent );
   self:write( "local __map = " );
   node.info.exp:filter( filterObj, node, baseIndent + stepIndent )
   self:writeln( "", baseIndent + stepIndent );
   self:writeln( "for __key in pairs( __map ) do", baseIndent + stepIndent * 2 );
   self:writeln( "table.insert( __sorted, __key )", baseIndent + stepIndent );
   self:writeln( "end", baseIndent + stepIndent );

   self:writeln( "table.sort( __sorted )", baseIndent + stepIndent );

   
   self:write( "for __index, " );
   local key = node.info.key and node.info.key.txt or "__key"
   self:write( key );
   self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 );
   self:writeln( string.format( "%s = __map[ %s ]", node.info.val.txt, key ),
		 baseIndent + stepIndent * 2 );
   node.info.block:filter( filterObj, node, baseIndent + stepIndent * 2 );
   self:writeln( "end", baseIndent + stepIndent );
   self:writeln( "end", baseIndent );
   self:writeln( "end", baseIndent );
end


filterObj[ TransUnit.nodeKind.ExpCall ] = function( self, node, parent, baseIndent )
   node.info.func:filter( filterObj, node, baseIndent )
   self:write( "( " )
   if node.info.argList then
      node.info.argList:filter( filterObj, node, baseIndent )
   end
   self:write( " )" )
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


filterObj[ TransUnit.nodeKind.ExpParen ] = function( self, node, parent, baseIndent )
   self:write( "(" )
   node.info:filter( filterObj, node, baseIndent )
   self:write( " )" )
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
      local prefixSymbol = node.info.prefix.info.txt
      if node.info.prefix.kind == TransUnit.nodeKind.ExpRef and
	 ( builtInModuleSet[ prefixSymbol ] or
	      self.moduleName2Info[ prefixSymbol ] or
	      className2InfoMap[ prefixSymbol ] )
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

   if node.info then
      node.info:filter( filterObj, node, baseIndent )
   end

   self:write( "}" )
end

filterObj[ TransUnit.nodeKind.LiteralMap ] = function( self, node, parent, baseIndent )
   self:write( "{" )
   for index, pair in pairs( node.info.pairList ) do
      if index > 1 then
	 self:write( ", " )
      end
      self:write( "[" )
      pair.key:filter( filterObj, node, baseIndent )
      self:write( "] = " )
      pair.val:filter( filterObj, node, baseIndent )
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

filterObj[ TransUnit.nodeKind.LiteralNil ] = function( self, node, parent, baseIndent )
   self:write( "nil" )
end

filterObj[ TransUnit.nodeKind.Break ] = function( self, node, parent, baseIndent )
   self:write( "break" )

end


return filterObj
