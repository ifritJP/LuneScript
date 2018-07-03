local Parser = {}

Parser.kind = {}

local kindSeed = 0
local kind2Txt = {}
local function regKind( name )
   local kind = kindSeed
   kindSeed = kindSeed + 1
   kind2Txt[ kind ] = name
   Parser.kind[ name ] = kind
   return kind
end

local kindCmnt = regKind( "Cmnt" )
local kindStr = regKind( "Str" )
local kindInt = regKind( "Int" )
local kindReal = regKind( "Real" )
local kindChar = regKind( "Char" )
local kindSymb = regKind( "Symb" )
local kindDlmt = regKind( "Dlmt" )
local kindKywd = regKind( "Kywd" )
local kindOpe = regKind( "Ope" )
local kindType = regKind( "Type" )
local kindEof = regKind( "EOF" )

local quotedCharSet = {}
quotedCharSet[ 'a' ] = true
quotedCharSet[ 'b' ] = true
quotedCharSet[ 'f' ] = true
quotedCharSet[ 'n' ] = true
quotedCharSet[ 'r' ] = true
quotedCharSet[ 't' ] = true
quotedCharSet[ 'v' ] = true
quotedCharSet[ '\\' ] = true
quotedCharSet[ '"' ] = true
quotedCharSet[ "'" ] = true



local op2Set = {}
op2Set[ '+' ] = true
op2Set[ '-' ] = true
op2Set[ '*' ] = true
op2Set[ '/' ] = true
op2Set[ '//' ] = true
op2Set[ '^' ] = true
op2Set[ '%' ] = true
op2Set[ '&' ] = true
op2Set[ '~' ] = true
op2Set[ '|' ] = true
op2Set[ '>>' ] = true
op2Set[ '<<' ] = true
op2Set[ '..' ] = true
op2Set[ '<' ] = true
op2Set[ '<=' ] = true
op2Set[ '>' ] = true
op2Set[ '>=' ] = true
op2Set[ '==' ] = true
op2Set[ '~=' ] = true
op2Set[ 'and' ] = true
op2Set[ 'or' ] = true
op2Set[ '@' ] = true
op2Set[ '=' ] = true

local op1Set = {}
op1Set[ '-' ] = true
op1Set[ 'not' ] = true
op1Set[ '#' ] = true
op1Set[ '~' ] = true
op1Set[ '*' ] = true



local function createReserveInfo( luaMode )
   local keywordSet = {}
   local typeSet = {}
   local builtInSet = {};
   keywordSet[ "local" ] = true
   keywordSet[ "function" ] = true
   keywordSet[ "if" ] = true
   keywordSet[ "else" ] = true
   keywordSet[ "elseif" ] = true
   keywordSet[ "while" ] = true
   keywordSet[ "for" ] = true
   keywordSet[ "in" ] = true
   keywordSet[ "return" ] = true
   keywordSet[ "break" ] = true
   keywordSet[ "nil" ] = true
   keywordSet[ "true" ] = true
   keywordSet[ "false" ] = true

   builtInSet[ "require" ] = true
   
   if luaMode then
      keywordSet[ "end" ] = true
      keywordSet[ "then" ] = true
      keywordSet[ "do" ] = true
      keywordSet[ "until" ] = true
   else
      keywordSet[ "let" ] = true
      keywordSet[ "mut" ] = true
      keywordSet[ "pub" ] = true
      keywordSet[ "pro" ] = true
      keywordSet[ "pri" ] = true
      keywordSet[ "fn" ] = true
      keywordSet[ "each" ] = true
      keywordSet[ "form" ] = true
      keywordSet[ "class" ] = true
      keywordSet[ "static" ] = true
      keywordSet[ "advertise" ] = true
      keywordSet[ "as" ] = true
      keywordSet[ "import" ] = true
      keywordSet[ "new" ] = true

      builtInSet[ "super" ] = true
      
      typeSet[ "int" ] = true
      typeSet[ "real" ] = true
      typeSet[ "stem" ] = true
      typeSet[ "str" ] = true
      typeSet[ "Map" ] = true
      typeSet[ "bool" ] = true
   end

   -- 2文字以上の演算子
   local multiCharDelimitMap = {}
   multiCharDelimitMap[ "=" ] = { "==" }
   multiCharDelimitMap[ "~" ] = { "~=" }
   multiCharDelimitMap[ "<" ] = { "<=" }
   multiCharDelimitMap[ ">" ] = { ">=" }
   multiCharDelimitMap[ "." ] = { ".." }
   multiCharDelimitMap[ "@" ] = { "@@" }
   multiCharDelimitMap[ "@@" ] = { "@@?" }
   multiCharDelimitMap[ ".." ] = { "..." }
  

   return keywordSet, typeSet, multiCharDelimitMap
end

function Parser.getKindTxt( kind )
   return kind2Txt[ kind ]
end

function Parser.isOp2( ope )
   return op2Set[ ope ]
end

function Parser.isOp1( ope )
   return op1Set[ ope ]
end

local ParserMtd = {
}

function Parser:create( path, luaMode )
   local stream = io.open( path, "r" )

   if not stream then
      return nil
   end
   
   local obj = {
      stream = stream,
      lineNo = 0,
      pos = 1,
      lineTokenList = {},
   }

   setmetatable( obj, { __index = ParserMtd } )

   local keywordSet, typeSet, multiCharDelimitMap =
      createReserveInfo( luaMode or string.find( path, "%.lua$" ) )

   obj.keywordSet = keywordSet
   obj.typeSet = typeSet
   obj.multiCharDelimitMap = multiCharDelimitMap

   return obj
end


function ParserMtd:parse()
   local function readLine()
      self.lineNo = self.lineNo + 1
      return self.stream:read( '*l' )
   end
   local rawLine = readLine()
   if not rawLine then
      return nil
   end

   local list = {}
   local startIndex = 1

   --[[
      複数行コメントの処理。

      @param comIndex 現在の解析行内の複数行コメント開始位置
      @param termStr 複数行コメントの終端文字列
   ]]
   local multiComment = function( comIndex, termStr )
      local searchIndex = comIndex
      local comment = ""
      while true do
	 local termIndex, termEndIndex = string.find(
	    rawLine, termStr, searchIndex, true )
	 if termIndex then
	    comment = comment .. rawLine:sub( searchIndex, termEndIndex )
	    return comment, termEndIndex + 1
	 end
	 comment = comment .. rawLine:sub( searchIndex ) .. "\n"
	 searchIndex = 1
	 rawLine = readLine()
	 if not rawLine then
	    error( "illegal comment" )
	 end
      end
   end

   --[[
      ソースをコメント、文字列、その他(ステートメント候補)に
      カテゴライズした結果を登録する。

      この関数内でステートメント候補の文字列をトークン毎に分割して登録する。

      @param kind カテゴライズの種類
      @param val カテゴライズした文字列
      @param column 現在の解析行内の位置
      
   ]]

   local addVal = function( kind, val, column )
      local function createInfo( tokenKind, token, tokenColumn )
	 if tokenKind == kindSymb then
	    if self.keywordSet[ token ] then
	       tokenKind = kindKywd
	    elseif self.typeSet[ token ] then
	       tokenKind = kindType
	    elseif op2Set[ token ] or op1Set[ token ] then
	       tokenKind = kindOpe
	    end
	 end
	 return { kind = tokenKind, txt = token,
		  pos = { lineNo = self.lineNo, column = tokenColumn } }
      end

      --[[
	 token の startIndex から始まる数値表現領域を特定する

	 @param token 文字列
	 @param startIndex token 内の検索開始位置。 この位置から数値表現が始まる。
	 @return 数値表現の終了位置, 整数かどうか
      ]]
      local function analyzeNumber( token, startIndex )
	 local nonNumIndex = token:find( '[^%d]', startIndex )
	 if not nonNumIndex then
	    return #token, true
	 end
	 local intFlag = true
	 local nonNumChar = token:byte( nonNumIndex )
	 if nonNumChar == 46 then -- .
	    intFlag = false
	    nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
	    nonNumChar = token:byte( nonNumIndex )
	 end
	 if nonNumChar == 120 or nonNumChar == 88 then -- X or x
	    nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
	    nonNumChar = token:byte( nonNumIndex )
	 end
	 if nonNumChar == 101 or nonNumChar == 69 then -- E or e
	    intFlag = false
	    local nextChar = token:byte( nonNumIndex + 1 )
	    if nextChar == 45 or nextChar == 43 then -- '-' or '+'
	       nonNumIndex = token:find( '[^%d]', nonNumIndex + 2 )
	    else
	       nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
	    end
	 end
	 if not nonNumIndex then
	    return #token, intFlag
	 end
	 return nonNumIndex - 1, intFlag
      end
      
      if kind == kindSymb then
	 local searchIndex = 1
	 while true do
	    -- 空白系以外の何らかの文字領域を探す
	    local tokenIndex, tokenEndIndex = string.find( val, "[%g]+", searchIndex )
	    if not tokenIndex then
	       break
	    end
	    local columnIndex = column + tokenIndex - 2
	    searchIndex = tokenEndIndex + 1
	    local token = val:sub( tokenIndex, tokenEndIndex )
	    local startIndex = 1
	    while true do
	       if token:find( '^[%d]', startIndex ) then
		  -- 数値の場合
		  local endIndex, intFlag = analyzeNumber( token, startIndex )
		  local info = createInfo(
		     intFlag and kindInt or kindReal,
		     token:sub( startIndex, endIndex ), columnIndex + startIndex )
		  table.insert( list, info )
		  startIndex = endIndex + 1
	       else
		  -- 区切り文字を探す
		  local index = string.find( token, '[^%w_]', startIndex )
		  if index then
		     if index > startIndex then
			local info = createInfo(
			   kindSymb, token:sub( startIndex, index - 1 ),
			   columnIndex + startIndex )
			table.insert( list, info )
		     end
		     local delimit = token:sub( index, index )
		     local candidateList = self.multiCharDelimitMap[ delimit ]
		     while candidateList do
			local findFlag = false
			for candIndex, candidate in ipairs( candidateList ) do
			   if candidate == token:sub( index, index + #candidate - 1 ) then
			      delimit = candidate
			      candidateList = self.multiCharDelimitMap[ delimit ]
			      findFlag = true
			      break
			   end
			end
			if not findFlag then
			   break
			end
		     end
		     startIndex = index + #delimit

		     local workKind = kindDlmt
		     if op2Set[ delimit ] or op1Set[ delimit ] then
			workKind = kindOpe
		     end
		     if delimit == "?" then
			local nextChar = token:sub( index, startIndex )
			table.insert( list, createInfo( kindChar, nextChar,
							columnIndex + startIndex ) )
			startIndex = startIndex + 1
		     else
			table.insert(
			   list, createInfo( workKind, delimit, columnIndex + index ) )
		     end
		  else
		     if startIndex <= #token then
			table.insert(
			   list, createInfo( kindSymb, token:sub( startIndex ),
					     columnIndex + startIndex ) )
		     end
		     break
		  end
	       end
	    end
	 end
      else
	 table.insert( list, createInfo( kind, val, column ) )
      end
   end
   
   -- 検索開始位置。
   -- 領域開始位置と検索開始位置が異なる場合がある。
   -- たとえば、 0.12e-2 のときに - の部分が検索開始位置、 0 の部分が領域開始位置になる
   local searchIndex = startIndex

   -- 領域をカテゴライズする
   while true do
      local syncIndexFlag = true
      local index = string.find( rawLine, [[[%-%?"%'%`%[].]], searchIndex )

      if not index then
	 addVal( kindSymb, rawLine:sub( startIndex ), startIndex )
	 return list
      end

      local findChar = string.byte( rawLine, index )
      local nextChar = string.byte( rawLine, index + 1 )

      if findChar == 45 and nextChar ~= 45 then -- --
	 searchIndex = index + 1
	 syncIndexFlag = false
      else
	 if startIndex < index then
	    addVal( kindSymb, rawLine:sub( startIndex, index - 1 ), startIndex )
	 end
	 if findChar == 39 and nextChar == 39 then -- "''"
	    if string.byte( rawLine, index + 2 ) == 39 then
	       -- 複数行コメントの場合
	       local comment, nextIndex = multiComment( index + 3, "'''" )
	       addVal( kindCmnt, "'''" .. comment, index )
	       searchIndex = nextIndex
	    else
	       addVal( kindCmnt, rawLine:sub( index ), index )
	       searchIndex = #rawLine + 1
	    end
	 elseif findChar == 91 then -- '['
	    if nextChar == 64 then -- '@'
	       addVal( kindDlmt, "[@", index )
	       searchIndex = index + 2
	    else
	       addVal( kindDlmt, "[", index )
	       searchIndex = index + 1
	    end
	 elseif findChar == 34 or findChar == 39 then -- '"' or "'"
	    -- 文字列の場合
	    local workIndex = index + 1
	    local pattern = '["\'\\]'
	    while true do
	       local endIndex = string.find( rawLine, pattern, workIndex )
	       if not endIndex then
		  error( string.format( "illegal string: %d: %s", index, rawLine ) )
	       end
	       local workChar = string.byte( rawLine, endIndex )
	       if workChar == findChar then -- '"'
		  addVal( kindStr, rawLine:sub( index, endIndex ), index )
		  searchIndex = endIndex + 1
		  break
	       elseif workChar == 92 then -- \\
		  workIndex = workIndex + 2
	       else
		  workIndex = workIndex + 1
	       end
	    end
	 elseif findChar == 96 then -- '`'
	    if nextChar == findChar and
	       string.byte( rawLine, index + 2 ) == 96
	    then -- '```'
	       -- 複数行文字列の場合
	       local str, nextIndex = multiComment( index + 3, '```' )
	       addVal( kindStr, '```' .. str, index )
	       searchIndex = nextIndex
	    else
	       addVal( kindDlmt, '`', index )
	    end
	 elseif findChar == 63 then -- ?
	    local codeChar = rawLine:sub( index + 1, index + 1 )
	    if nextChar == 92 then -- \\
	       local quoted = rawLine:sub( index + 2, index + 2 )
	       if quotedCharSet[ quoted ] then
		  codeChar = rawLine:sub( index + 1, index + 2 )
	       else
		  codeChar = quoted
	       end
	       searchIndex = index + 3
	    else
	       searchIndex = index + 2
	    end
	    addVal( kindChar, codeChar, index )
	 else
	    error( string.format( "illegal syntax:%s:%s",
				  self.lineNo, rawLine:sub( index ) ) )
	 end
      end
      if syncIndexFlag then
	 startIndex = searchIndex
      end
   end
end


function ParserMtd:getToken()
   if not self.lineTokenList then
      return 
   end
   if #self.lineTokenList < self.pos then
      self.pos = 1
      self.lineTokenList = {}
      while #self.lineTokenList == 0 do
	 self.lineTokenList = self:parse()
	 if not self.lineTokenList then
	    return nil
	 end
      end
   end

   local token = self.lineTokenList[ self.pos ]
   self.pos = self.pos + 1

   return token
end

local eofToken = { kind = kindEof, txt = "", pos = { lineNo = 0, column = 0 } }
function Parser:getEofToken()
   return eofToken
end

return Parser
