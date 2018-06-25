local Parser = require( 'primal.Parser' )

local TransUnit = {}


local nodeKind2NameMap = {}
local nodeKindSeed = 0
local nodeKind = {}

TransUnit.nodeKind = nodeKind

local function regKind( name )
   local kind = nodeKindSeed
   nodeKindSeed = nodeKindSeed + 1
   nodeKind2NameMap[ kind ] = name
   nodeKind[ name ] = kind
   return kind
end

local nodeKindRoot = regKind( 'Root' )
local nodeKindRefType = regKind( 'RefType' )
local nodeKindIf = regKind( 'If' )
local nodeKindWhile = regKind( 'While' )
local nodeKindRepeat = regKind( 'Repeat' )
local nodeKindFor = regKind( 'For' )
local nodeKindApply = regKind( 'Apply' )
local nodeKindForeach = regKind( 'Foreach' )
local nodeKindReturn = regKind( 'Return' )
local nodeKindBreak = regKind( 'Break' )
local nodeKindExpList = regKind( 'ExpList' )
local nodeKindExpRef = regKind( 'ExpRef' )
local nodeKindExpOp2 = regKind( 'ExpOp2' )
local nodeKindExpCast = regKind( 'ExpCast' )
local nodeKindExpOp1 = regKind( 'ExpOp1' )
local nodeKindExpRefItem = regKind( 'ExpRefItem' )
local nodeKindExpCall = regKind( 'ExpCall' )
local nodeKindExpDDD = regKind( 'ExpDDD' )
local nodeKindBlock = regKind( 'Block' )
local nodeKindStmtExp = regKind( 'StmtExp' )
local nodeKindRefField = regKind( 'RefField' )
local nodeKindDeclVar = regKind( 'DeclVar' )
local nodeKindDeclFunc = regKind( 'DeclFunc' )
local nodeKindDeclArg = regKind( 'DeclArg' )
local nodeKindDeclArgDDD = regKind( 'DeclArgDDD' )
local nodeKindLiteralChar = regKind( 'LiteralChar' )
local nodeKindLiteralNum = regKind( 'LiteralNum' )
local nodeKindLiteralArray = regKind( 'LiteralArray' )
local nodeKindLiteralList = regKind( 'LiteralList' )
local nodeKindLiteralMap = regKind( 'LiteralMap' )
local nodeKindLiteralString = regKind( 'LiteralString' )
local nodeKindLiteralBool = regKind( 'LiteralBool' )

local function nodeFilter( node, filter, ... )
   if not filter[ node.kind ] then
      error( string.format( "none filter -- %s",
			    TransUnit:getNodeKindName( node.kind ) ))
   end
   filter[ node.kind ]( filter, node, ... )
end

function TransUnit:getNodeKindName( kind )
   return nodeKind2NameMap[ kind ]
end


function TransUnit:createAST( parser )

   local rootInfo = {}
   rootInfo.childlen = {}
   self.ast = self:createNode(
      nodeKindRoot, { lineNo = 0, column = 0 }, rootInfo )
   self.parser = parser
   
   TransUnit:analyzeStatement( rootInfo.childlen )

   local token = self:getTokenNoErr()
   if token then
      error( string.format( "unknown:%d:%d:(%s) %s",
			    token.pos.lineNo, token.pos.column,
			    Parser.getKindTxt( token.kind ), token.txt ) )
   end

   return self.ast
end

function TransUnit:createNode( kind, pos, info )
   if not self:getNodeKindName( kind ) then
      error( string.format( "%d:%d: not found nodeKind", pos.lineNo, pos.column ) )
   end
   return { kind = kind, pos = pos, info = info, filter = nodeFilter }
end

function TransUnit:analyzeDecl( firstToken, token )
   if token.txt == "let" then
      return self:analyzeDeclVar( token )
   elseif token.txt == "fn" then
      return self:analyzeDeclFunc( token )
   end

   return nil
end

function TransUnit:analyzeStatement( stmtList, termTxt )
   while true do
      local token = self:getTokenNoErr()
      if not token then
	 print( "EOF" )
	 break
      end

      local statement = self:analyzeDecl( token, token )

      if not statement then
	 if token.txt == termTxt then
	    self:pushback()
	    break
	 elseif token.txt == "pub" or token.txt == "pro" or
	    token.txt == "pri" or token.txt == "global"
	 then
	    statement = self:analyzeDecl( token, self:getToken( "illegal syntax" ) )
	 elseif token.txt == "{" then
	    self:pushback()
	    statement = self:analyzeBlock( "{" )
	 elseif token.txt == "if" then
	    statement = self:analyzeIf( token )
	 elseif token.txt == "while" then
	    statement = self:analyzeWhile( token )
	 elseif token.txt == "repeat" then
	    statement = self:analyzeRepeat( token )
	 elseif token.txt == "for" then
	    statement = self:analyzeFor( token )
	 elseif token.txt == "apply" then
	    statement = self:analyzeApply( token )
	 elseif token.txt == "foreach" then
	    statement = self:analyzeForeach( token )
	 elseif token.txt == "return" then
	    local expList = self:analyzeExpList()
	    self:checkNextToken( ";" )
	    statement = self:createNode( nodeKindReturn, token.pos, expList )
	 elseif token.txt == "break" then
	    self:checkNextToken( ";" )
	    statement = self:createNode( nodeKindBreak, token.pos, nil )
	 else
	    self:pushback()
	    local exp = self:analyzeExp()
	    self:checkNextToken( ";" )
	    statement = self:createNode( nodeKindStmtExp, token.pos, exp )
	 end
      end

      if not statement then
	 break
      end
      table.insert( stmtList, statement )
   end
end

function TransUnit:pushback()
   if self.pushbackToken then
      error( string.format( "multiple pushback:%d:%d: %s, %s",
			    self.currentToken.pos.lineNo,
			    self.currentToken.pos.column,
			    self.pushbackToken.txt, self.currentToken.txt ) )
   end
   self.pushbackToken = self.currentToken
   self.currentToken = nil
end

function TransUnit:getToken( mess )
   local token = self:getTokenNoErr()
   if not token then
      self:error( mess )
   end
   self.currentToken = token
   return self.currentToken
end

function TransUnit:getTokenNoErr()
   if self.pushbackToken then
      self.currentToken = self.pushbackToken
      self.pushbackToken = nil
      return self.currentToken
   end

   local commentList = {}
   local token
   while true do
      token = self.parser:getToken()
      if not token then
	 break
      end
      if token.kind ~= Parser.kind.Cmnt then
	 break
      end
      table.insert( commentList, token )
   end

   if token then
      token.commentList = commentList
      self.currentToken = token
   end
   return token
end

function TransUnit:getSymbolToken()
   return self:checkSymbol( self:getToken() )
end

function TransUnit:checkSymbol( token )
   if token.kind ~= Parser.kind.Symb and
      token.kind ~= Parser.kind.Kywd and
      token.kind ~= Parser.kind.Type
   then
      self:error( "illegal symbol" )
   end
   return token
end


function TransUnit:error( mess )
   local pos = { lineNo = 0, column = 0 }
   local txt = ""
   if self.currentToken then
      pos = self.currentToken.pos
      txt = self.currentToken.txt
   end
   error( string.format( "%d:%d:(%s) %s", pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:checkNextToken( txt )
   return self:checkToken( self:getToken( mess ), txt )
end

function TransUnit:checkToken( token, txt )
   if not token or token.txt ~= txt then
      self:error( string.format( "not found -- %s", txt ) )
   end
   return token
end

function TransUnit:analyzeIf( token )
   local list = {}
   table.insert(
      list, { kind = "if", exp = self:analyzeExp(), block = self:analyzeBlock( "if" ) } )

   local token = self:getToken( "illegal if syntax" )
   while token.txt == "elseif" do
      table.insert(
	 list, { kind = "elseif", exp = self:analyzeExp(),
		 block = self:analyzeBlock( "elseif" ) } )
      token = self:getToken( "illegal if syntax" )
   end

   table.insert(
      list, { kind = "else", block = self:analyzeBlock( "else" ) } )

   return self:createNode( nodeKindIf, token.pos, list )
end

function TransUnit:analyzeWhile( token )
   local info = { exp = self:analyzeExp(), block = self:analyzeBlock( "while" ) }
   return self:createNode( nodeKindWhile, token.pos, info )
end

function TransUnit:analyzeRepeat( token )
   local info = { block = self:analyzeBlock(), exp = self:analyzeExp( "repeat" ) }
   local node = self:createNode( nodeKindRepeat, token.pos, info )
   self:checkNextToken( ";" )
   return node
end

function TransUnit:analyzeFor( token )
   local val = self:getToken()
   if val.kind ~= Parser.kind.Symb then
      self:error( "not symbol" )
   end
   self:checkNextToken( "=" )
   local exp1 = self:analyzeExp()
   self:checkNextToken( "," )
   local exp2 = self:analyzeExp()
   local token = self:getToken( "illegal syntax for" )
   local exp3
   if token.txt == "," then
      exp3 = self:analyzeExp()
   else
      self:pushback()
   end
   
   local info = { block = self:analyzeBlock( "for" ), val = val,
		  init = exp1, to = exp2, delta = exp3 }
   local node = self:createNode( nodeKindFor, token.pos, info )
   return node
end

function TransUnit:analyzeApply( token )
   local varList = {}
   local nextToken
   repeat
      local var = self:getToken( "illegal syntax apply" )
      if var.kind ~= Parser.kind.Symb then
	 self:error( "illegal symbol" )
      end
      table.insert( varList, var )
      nextToken = self:getToken( "illegal symbol apply" )
   until nextToken.txt ~= ","
   self:checkToken( nextToken, "of" )

   local exp = self:analyzeExp()
   if exp.kind ~= nodeKindExpCall then
      self:error( "not call" )
   end

   local block = self:analyzeBlock( "apply" )

   local info = { varList = varList, exp = exp, block = block }
   return self:createNode( nodeKindApply, token.pos, info )
end

function TransUnit:analyzeForeach( token )
   local varList = {}
   local nextToken
   repeat
      local var = self:getToken( "illegal syntax apply" )
      if var.kind ~= Parser.kind.Symb then
	 self:error( "illegal symbol" )
      end
      table.insert( varList, var )
      nextToken = self:getToken( "illegal symbol apply" )
   until nextToken.txt ~= ","
   self:checkToken( nextToken, "in" )

   local exp = self:analyzeExp()

   local block = self:analyzeBlock( "foreach" )

   local info = { varList = varList, exp = exp, block = block }
   return self:createNode( nodeKindForeach, token.pos, info )
end


function TransUnit:analyzeRefType()
   local firstToken = self:getToken( "not found type" )
   local token = firstToken
   local refFlag = false
   if token.txt == "&" then
      refFlag = true
      token = self:getToken( "not found type" )
   end
   local mutFlag = false
   if token.txt == "mut" then
      mutFlag = true
      token = self:getToken( "not found type" )
   end
   local name = self:checkSymbol( token )
   local arrayMode = "no"
   token = self:getToken( "not found type" )
   if token.txt == '[' or token.txt == '[@' then
      if token.txt == '[' then
	 arrayMode = "list"
      else
	 arrayMode = "array"
      end
      token = self:getToken( "not found array" )
      if token.txt ~= ']' then
	 self:pushback()
	 self:checkNextToken( ']' )
      end
   else
      self:pushback()
   end

   return self:createNode(
      nodeKindRefType, firstToken.pos,
      { name = name, refFlag = refFlag, mutFlag = mutFlag, array = arrayMode } )
end

function TransUnit:analyzeDeclFunc( fnToken )
   local argList = {}
   local token = self:getToken( "not found fund name" )
   local name = nil
   if token ~= "(" then
      name = self:checkSymbol( token )
      self:checkNextToken( "(" )
   end
   repeat
      local argName = self:getToken( "illegal symbol" )
      if argName.txt == "..." then
	 table.insert( argList, self:createNode( nodeKindDeclArgDDD,
						 argName.pos, argName ) )
      else
	 self:checkSymbol( argName )
	 
	 self:checkNextToken( ":" )
	 local refType = self:analyzeRefType()
	 local arg = self:createNode( nodeKindDeclArg, argName.pos,
				      { name = argName, argType = refType } )
	 table.insert( argList, arg )
      end
      token = self:getToken( "illegal fn syntax" )
   until token.txt ~= ","

   self:checkToken( token, ")" )

   token = self:getToken( "illegal fn syntax" )
   local typeList = {}
   if token.txt == ":" then
      repeat
	 table.insert( typeList, self:analyzeRefType() )
	 token = self:getToken( "illegal fn syntax" )
      until token.txt ~= ","
   end

   self:pushback()
   local body = self:analyzeBlock( "func" )
   local info = { name = name, argList = argList,
		  retTypeList = typeList, body = body }
   return self:createNode( nodeKindDeclFunc, fnToken.pos, info )
end

function TransUnit:analyzeBlock( blockKind )
   local token = self:checkNextToken( "{" )

   local stmtList = {}
   self:analyzeStatement( stmtList, "}" )

   self:checkNextToken( "}" )

   return self:createNode( nodeKindBlock, token.pos,
			   { kind = blockKind, stmtList = stmtList } )
end

function TransUnit:analyzeDeclVar( letToken )
   local varList = {}
   local token
   repeat
      local varName = self:getSymbolToken()
      token = self:getToken( "illegal let syntax" )
      if token.txt == ":" then
	 local refType = self:analyzeRefType()
	 token = self:getToken( "illegal let syntax" )
      end
      table.insert( varList, { name = varName, refType = refType } )
   until token.txt ~= ","
   
   local expList
   if token.txt == "=" then
      expList = self:analyzeExpList()
   end

   self:checkNextToken( ";" )

   local declVarInfo = { varList = varList, expList = expList }
   return self:createNode( nodeKindDeclVar, letToken.pos, declVarInfo )
end

function TransUnit:analyzeExpList()
   local expList = {}
   local firstExp = nil
   repeat
      local exp = self:analyzeExp()
      if not firstExp then
	 firstExp = exp
      end
      table.insert( expList, exp )
      local token = self:getToken( "not found expList" )
   until token.txt ~= ","

   self:pushback()

   return self:createNode( nodeKindExpList, firstExp.pos, expList )
end

function TransUnit:analyzeListConst( token )
   local nextToken = self:getToken( "not found next token" )
   local expList
   if nextToken.txt ~= "]" then
      self:pushback()
      expList = self:analyzeExpList()
      self:checkNextToken( "]" )
   end
   local kind = nodeKindLiteralArray
   if token.txt == '[' then
      kind = nodeKindLiteralList
   end
   return self:createNode( kind, token.pos, expList )
end

function TransUnit:analyzeMapConst( token )
   local nextToken
   local map = {}
   repeat
      nextToken = self:getToken( "" )
      if nextToken.txt == "}" then
	 break
      end
      self:pushback()
      
      local key = self:analyzeExp()
      self:checkNextToken( ":" )
      local val = self:analyzeExp()
      map[ key ] = val
      nextToken = self:getToken( "illegal map constructor" )
   until nextToken.txt ~= ","

   self:checkToken( nextToken, "}" )
   return self:createNode( nodeKindLiteralMap, token.pos, map )
end

function TransUnit:analyzeExpRefItem( token, exp )
   local indexExp = self:analyzeExp()
   self:checkNextToken( "]" )

   local info = { val = exp, index = indexExp }
   return self:createNode( nodeKindExpRefItem, token.pos, info )
end   

function TransUnit:analyzeExpSymbol( firstToken, token, fieldFlag, prefixExp )
   local exp

   if fieldFlag then
      local info = { field = token, prefix = prefixExp }
      exp = self:createNode( nodeKindRefField, firstToken.pos, info )
   else
      exp = self:createNode( nodeKindExpRef, firstToken.pos, token )
   end

   local nextToken = self:getToken( "illegal syntax" )
   repeat
      local matchFlag = false
      if nextToken.txt == "[" then
	 matchFlag = true
	 exp = self:analyzeExpRefItem( nextToken, exp )
	 nextToken = self:getToken( "illegal syntax" )
      end
      if nextToken.txt == "(" then
	 matchFlag = true
	 local expList = self:analyzeExpList()
	 self:checkNextToken( ")" )
	 local info = { func = exp, argList = expList }

	 exp = self:createNode( nodeKindExpCall, token.pos, info )
	 nextToken = self:getToken( "illegal syntax" )
      end
   until not matchFlag

   if nextToken.txt == "." then
      return self:analyzeExpSymbol(
	 firstToken, self:getToken( "illegal field" ), true, exp )
   end
   
   self:pushback()
   return exp
end


function TransUnit:analyzeExp( skipOp2Flag )
   local token = self:getToken( "not found exp" )

   if token.txt == "..." then
      return self:createNode( nodeKindExpDDD, token.pos, token )
   end
   
   if token.txt == '[' or token.txt == '[@' then
      return self:analyzeListConst( token )
   end
   if token.txt == '{' then
      return self:analyzeMapConst( token )
   end
   
   local exp
   if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
      -- 単項演算
      exp = self:analyzeExp( true )
      exp = self:createNode( nodeKindExpOp1, token.pos, { op = token, exp = exp } )
      return self:analyzeExpOp2( token, exp )
   end

   if token.kind == Parser.kind.Num then
      exp = self:createNode( nodeKindLiteralNum, token.pos,
			     { token = token, num = tonumber( token.txt ) } )
   elseif token.kind == Parser.kind.Char then
      exp = self:createNode( nodeKindLiteralChar, token.pos,
			     { token = token, num = token.txt:byte( 2 ) } )
   elseif token.kind == Parser.kind.Str then
      exp = self:createNode( nodeKindLiteralString, token.pos, token )
      local nextToken = self:getToken( "illegal syntax" )
      if nextToken.txt == "[" then
	 exp = self:analyzeExpRefItem( nextToken, exp )
      else
	 self:pushback()
      end
   elseif token.kind == Parser.kind.Symb then
      exp = self:analyzeExpSymbol( token, token, false, token )
   elseif token.kind == Parser.kind.Type then
      exp = self:createNode( nodeKindExpRef, token.pos, token )
   elseif token.txt == "true" or token.txt == "false" then
      exp = self:createNode( nodeKindLiteralBool, token.pos, token )
   end

   if not exp then
      self:error( "illegal exp" )
   end

   if skipOp2Flag then
      return exp
   end
   
   return self:analyzeExpOp2( token, exp )
end

function TransUnit:analyzeExpOp2( token, exp )
   local nextToken = self:getToken( "not found next token" )
   while true do
      if nextToken.txt == "@" then
	 local castType = self:analyzeRefType()
	 local info = { exp = exp, castType = castType }
	 exp = self:createNode( nodeKindExpCast, token.pos, info )
      elseif nextToken.kind == Parser.kind.Ope then
	 if Parser.isOp2( nextToken.txt ) then
	    local exp2 = self:analyzeExp()
	    local info = { op = nextToken, exp1 = exp, exp2 = exp2 }
	    exp = self:createNode( nodeKindExpOp2, token.pos, info )
	 else
	    self:error( "illegal op" )
	 end
      else
	 self:pushback()
	 return exp
      end
      nextToken = self:getToken( "not found next token" )
   end
end

return TransUnit
