-- test
--[[
   'test'"
]]

local kindCmnt = 0
local kindStr = 1
local kindStmt = 2
local kindDlmt = 3
local kindKywd = 4

local kind2Txt = {}
kind2Txt[ kindCmnt ] = "comment"
kind2Txt[ kindStr ] = "string"
kind2Txt[ kindStmt ] = "stmt"
kind2Txt[ kindDlmt ] = "delimit"
kind2Txt[ kindKywd ] = "keyword"


local function createReserveInfo( luaMode )
   local keywordSet = {}
   keywordSet[ "local" ] = true
   keywordSet[ "function" ] = true
   keywordSet[ "if" ] = true
   keywordSet[ "else" ] = true
   keywordSet[ "elseif" ] = true
   keywordSet[ "while" ] = true
   keywordSet[ "for" ] = true
   keywordSet[ "in" ] = true
   keywordSet[ "return" ] = true
   keywordSet[ "require" ] = true
   keywordSet[ "and" ] = true
   keywordSet[ "or" ] = true
   keywordSet[ "break" ] = true
   keywordSet[ "nil" ] = true
   keywordSet[ "true" ] = true
   keywordSet[ "false" ] = true
   keywordSet[ "self" ] = true
   keywordSet[ "not" ] = true

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
      keywordSet[ "int" ] = true
      keywordSet[ "real" ] = true
      keywordSet[ "stem" ] = true
      keywordSet[ "str" ] = true
      keywordSet[ "Map" ] = true
      keywordSet[ "each" ] = true
      keywordSet[ "form" ] = true
      keywordSet[ "class" ] = true
      keywordSet[ "super" ] = true
      keywordSet[ "static" ] = true
      keywordSet[ "advertise" ] = true
      keywordSet[ "as" ] = true
      keywordSet[ "import" ] = true
   end

   -- 2文字以上の演算子
   local multiCharDelimitMap = {}
   multiCharDelimitMap[ "=" ] = "=="
   multiCharDelimitMap[ "~" ] = "~="
   multiCharDelimitMap[ "<" ] = "<="
   multiCharDelimitMap[ ">" ] = ">="
   multiCharDelimitMap[ "." ] = ".."
   multiCharDelimitMap[ ".." ] = "..."

   return keywordSet, multiCharDelimitMap
end


local Parser = {}

function Parser.getKindTxt( kind )
   return kind2Txt[ kind ]
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

   local keywordSet, multiCharDelimitMap =
      createReserveInfo( luaMode or string.find( path, "%.lua$" ) )

   obj.keywordSet = keywordSet
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

   local multiComment = function( comIndex, pattern )
      local searchIndex = comIndex
      local comment = ""
      while true do
	 local termIndex = string.find( rawLine, ']', searchIndex, true )
	 while termIndex do
	    local termEndIndex = string.find( rawLine, ']', termIndex + 1, true )
	    if termEndIndex then
	       if pattern == rawLine:sub( termIndex, termEndIndex ) then
		  comment = comment .. rawLine:sub( searchIndex, termEndIndex )
		  return comment, termEndIndex + 1
	       end
	       termIndex = termEndIndex
	    else
	       break
	    end
	 end
	 comment = comment .. rawLine .. "\n"
	 searchIndex = 1
	 rawLine = readLine()
	 if not rawLine then
	    error( "illegal comment" )
	    return nil
	 end
      end
   end
   
   local addVal = function( kind, val, column )
      local function createInfo( tokenKind, token, tokenColumn )
	 if tokenKind == kindStmt and self.keywordSet[ token ] then
	    tokenKind = kindKywd
	 end
	 return { kind = tokenKind, lineNo = self.lineNo,
		  column = tokenColumn, txt = token }
      end
      if kind == kindStmt then
	 local searchIndex = 1
	 while true do
	    local tokenIndex, tokenEndIndex = string.find( val, "[%g]+", searchIndex )
	    if not tokenIndex then
	       break
	    end
	    local columnIndex = column + tokenIndex - 2
	    searchIndex = tokenEndIndex + 1
	    local token = val:sub( tokenIndex, tokenEndIndex )
	    local startIndex = 1
	    while true do
	       local index = string.find( token, '[^%w]', startIndex )
	       if index then
		  if index > startIndex then
		     local info = createInfo(
			kindStmt, token:sub( startIndex, index - 1 ),
			columnIndex + startIndex )
		     table.insert( list, info )
		  end
		  local delimit = token:sub( index, index )
		  local candidate = self.multiCharDelimitMap[ delimit ]
		  while candidate do
		     if candidate == token:sub( index, index + #candidate - 1 ) then
			delimit = candidate
			candidate = self.multiCharDelimitMap[ delimit ]
		     else
			break
		     end
		  end
		  startIndex = index + #delimit
		  
		  table.insert(
		     list, createInfo( kindDlmt, delimit, columnIndex + index ) )
	       else
		  if startIndex <= #token then
		     table.insert(
			list, createInfo( kindStmt, token:sub( startIndex ),
					  columnIndex + startIndex ) )
		  end
		  break
	       end
	    end
	 end
      else
	 table.insert( list, createInfo( kind, val, column ) )
      end
   end
   

   while true do
      local index = string.find( rawLine, [[[%-"%'%[].]], startIndex )

      if not index then
	 addVal( kindStmt, rawLine:sub( startIndex ), startIndex )
	 return list
      end

      if startIndex < index then
	 addVal( kindStmt, rawLine:sub( startIndex, index - 1 ), startIndex )
      end

      local findChar = string.byte( rawLine, index )
      local nextChar = string.byte( rawLine, index + 1 )
      if findChar == 45 then -- '-'
	 if nextChar == 45 then -- '--'
	    -- コメントの場合
	    if string.byte( rawLine, index + 2 ) == 91 and
	       string.byte( rawLine, index + 3 ) == 91
	    then
	       -- 複数行コメントの場合
	       local comment, nextIndex = multiComment( index + 2, "]]" )
	       addVal( kindCmnt, comment, index )
	       startIndex = nextIndex
	    else
	       addVal( kindCmnt, rawLine:sub( index ), index )
	       return list
	    end
	 else
	    addVal( kindDlmt, "-", index )
	    startIndex = index + 1
	 end
      elseif findChar == 91 then -- '['
	 if nextChar == 91 then -- '['
	    -- 文字列の場合 [[
	    local str, nextIndex = multiComment( index, "]]" )
	    addVal( kindStr, str, index )
	    startIndex = nextIndex
	 else
	    addVal( kindDlmt, "[", index )
	    startIndex = index + 1
	 end
      elseif findChar == 34 or findChar == 39   then -- '"' or "'"
	 -- 文字列の場合
	 local workIndex = index + 1
	 local pattern = findChar == 34 and '["\\]' or "['\\]"
	 while true do
	    local endIndex = string.find( rawLine, pattern, workIndex )
	    if not endIndex then
	       error( "illegal string:", rawLine )
	    end
	    if string.byte( rawLine, endIndex ) == findChar then -- -- '"' or "'"
	       addVal( kindStr, rawLine:sub( index, endIndex ), index )
	       startIndex = endIndex + 1
	       break
	    else -- '\'
	       workIndex = workIndex + 2
	    end
	 end
      end
   end
end

function ParserMtd:getToken()
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

return Parser
