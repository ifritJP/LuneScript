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
local nodeKindComment = regKind( 'Comment' )
local nodeKindReturn = regKind( 'Return' )
local nodeKindRefType = regKind( 'RefType' )
local nodeKindExpList = regKind( 'ExpList' )
local nodeKindExpRef = regKind( 'ExpRef' )
local nodeKindExpOp2 = regKind( 'ExpOp2' )
local nodeKindExpOp1 = regKind( 'ExpOp1' )
local nodeKindExpRefItem = regKind( 'ExpRefItem' )
local nodeKindStmtExp = regKind( 'StmtExp' )
local nodeKindVarDecl = regKind( 'VarDecl' )
local nodeKindLiteralChar = regKind( 'LiteralChar' )
local nodeKindLiteralNum = regKind( 'LiteralNum' )
local nodeKindLiteralArray = regKind( 'LiteralArray' )
local nodeKindLiteralList = regKind( 'LiteralList' )

local function nodeFilter( node, filter, ... )
   filter[ node.kind ]( node, ... )
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
      error( string.format( "%d:%d", pos.lineNo, pos.column ) )
   end
   print( self:getNodeKindName( kind ) )
   return { kind = kind, pos = pos, info = info, filter = nodeFilter }
end

function TransUnit:analyzeStatement( stmtList, termTxt )
   while true do
      local token = self:getTokenNoErr()
      if not token then
	 print( "EOF" )
	 break
      end

      local statement

      if token.txt == termTxt then
	 self:pushback()
	 break
      elseif token.txt == "let" then
	 statement = self:analyzeDeclVar( token )
      elseif token.txt == "fn" then
	 statement = self:analyzeDeclFunc( token )
      elseif token.txt == "return" then
	 local expList = self:analyzeExpList()
	 self:checkNextToken( ";" )
	 statement = self:createNode( nodeKindReturn, token.pos, expList )
      elseif token.kind == Parser.kind.Cmnt then
	 statement = self:createNode( nodeKindComment, token.pos, token.txt )
      else
	 self:pushback()
	 local exp = self:analyzeExp()
	 self:checkNextToken( ";" )
	 statement = self:createNode( nodeKindStmtExp, token.pos, exp )
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

   local token = self.parser:getToken()
   if token then
      self.currentToken = token
   end
   return token
end

function TransUnit:getSymbolToken()
   return self:checkSymbol( self:getToken() )
end

function TransUnit:checkSymbol( token )
   if token.kind ~= Parser.kind.Stmt and
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
   self:checkToken( self:getToken( mess ), txt )
end

function TransUnit:checkToken( token, txt )
   if not token or token.txt ~= txt then
      self:error( string.format( "not found -- %s", txt ) )
   end
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
   token = self:getToken( "not found type" )
   if token.txt == '[' or token.txt == '[@' then
      token = self:getToken( "not found array" )
      if token.txt ~= ']' then
	 self:pushback()
	 self:checkNextToken( ']' )
      end
   else
      self:pushback()
   end

   return { name = name, refFlag = refFlag, mutFlag = mutFlag }
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
      local argName = self:getSymbolToken()
      self:checkNextToken( ":" )
      local refType = self:analyzeRefType()
      table.insert( argList, { name = argName, refType = refType } )
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

   self:checkToken( token, "{" )

   local stmtList = {}
   self:analyzeStatement( stmtList, "}" )

   self:checkNextToken( "}" )
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
   return self:createNode( nodeKindVarDecl, letToken.pos, declVarInfo )
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

function TransUnit:analyzeExp()
   local token = self:getToken( "not found expt" )

   if token.txt == '[' or token.txt == '[@' then
      local nextToken = self:getToken( "not found next token" )
      local expList
      if nextToken.txt ~= "]" then
	 self:pushback()
	 expList = self:analyzeExpList()
      end
      self:checkNextToken( "]" )
      local kind = nodeKindLiteralArray
      if token.txt == '[' then
	 kind = nodeKindLiteralList
      end
      return self:createNode( kind, token.pos, expList )
   end

   if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
      local exp = self:analyzeExp()
      return self:createNode( nodeKindExpOp1, token.pos, { op = token.txt, exp = exp } )
   end

   
   local exp
   if token.kind == Parser.kind.Num then
      exp = self:createNode( nodeKindLiteralNum, token.pos, token.txt )
   elseif token.kind == Parser.kind.Stmt then
      exp = self:createNode( nodeKindExpRef, token.pos, token.txt )
      local nextToken = self:getToken( "illegal syntax" )
      if nextToken.txt == "[" then
	 local indexExp = self:analyzeExp()
	 self:checkNextToken( "]" )

	 local info = { array = exp, index = indexExp }
	 exp = self:createNode( nodeKindExpRefItem, token.pos, info )
      else
	 self:pushback()
      end
   elseif token.kind == Parser.kind.Type then
      exp = self:createNode( nodeKindExpRef, token.pos, token.txt )
   end

   if not exp then
      self:error( "illegal exp" )
   end

   local nextToken = self:getToken( "not found next token" )
   if nextToken.kind == Parser.kind.Ope then
      if not Parser.isOp2( nextToken.txt ) then
	 self:error( "illegal op" )
      end
      local exp2 = self:analyzeExp()
      local info = { op = nextToken.txt, exp1 = exp, exp2 = exp2 }
      return self:createNode( nodeKindExpOp2, token.pos, info )
   else
      self:pushback()
      return exp
   end
end


return TransUnit
