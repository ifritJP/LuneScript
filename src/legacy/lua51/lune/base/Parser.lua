--lune/base/Parser.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Parser'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

if not _lune2 then
   _lune2 = _lune
end

local Util = _lune.loadModule( 'lune.base.Util' )

local luaKeywordSet = {}
luaKeywordSet["let"]= true
luaKeywordSet["if"]= true
luaKeywordSet["else"]= true
luaKeywordSet["elseif"]= true
luaKeywordSet["while"]= true
luaKeywordSet["for"]= true
luaKeywordSet["in"]= true
luaKeywordSet["return"]= true
luaKeywordSet["break"]= true
luaKeywordSet["nil"]= true
luaKeywordSet["true"]= true
luaKeywordSet["false"]= true
luaKeywordSet["{"]= true
luaKeywordSet["}"]= true
luaKeywordSet["do"]= true
luaKeywordSet["require"]= true
luaKeywordSet["function"]= true
luaKeywordSet["then"]= true
luaKeywordSet["until"]= true

local function isLuaKeyword( txt )

   return _lune._Set_has(luaKeywordSet, txt )
end
_moduleObj.isLuaKeyword = isLuaKeyword

local function createReserveInfo( luaMode )

   local keywordSet = {}
   local typeSet = {}
   local builtInSet = {}
   
   builtInSet["require"]= true
   for key, __val in pairs( luaKeywordSet ) do
      if not _lune._Set_has(builtInSet, key ) then
         keywordSet[key]= true
      end
      
   end
   
   
   if not luaMode then
      keywordSet["null"]= true
      keywordSet["let"]= true
      keywordSet["mut"]= true
      keywordSet["pub"]= true
      keywordSet["pro"]= true
      keywordSet["pri"]= true
      keywordSet["fn"]= true
      keywordSet["each"]= true
      keywordSet["form"]= true
      keywordSet["class"]= true
      builtInSet["super"]= true
      keywordSet["static"]= true
      keywordSet["advertise"]= true
      keywordSet["import"]= true
      keywordSet["new"]= true
      keywordSet["!"]= true
      keywordSet["unwrap"]= true
      keywordSet["sync"]= true
      
      typeSet["int"]= true
      typeSet["real"]= true
      typeSet["stem"]= true
      typeSet["str"]= true
      typeSet["Map"]= true
      typeSet["bool"]= true
   end
   
   
   local multiCharDelimitMap = {}
   multiCharDelimitMap["="] = {"=="}
   multiCharDelimitMap["<"] = {"<="}
   multiCharDelimitMap[">"] = {">="}
   
   if not luaMode then
      multiCharDelimitMap["|"] = {"|<", "|>"}
      multiCharDelimitMap["|<"] = {"|<<"}
      multiCharDelimitMap["|>"] = {"|>>"}
      multiCharDelimitMap["["] = {"[@"}
      multiCharDelimitMap["("] = {"(@"}
      multiCharDelimitMap["~"] = {"~=", "~~"}
      multiCharDelimitMap["$"] = {"$[", "$.", "$("}
      multiCharDelimitMap["$."] = {"$.$"}
      multiCharDelimitMap["."] = {"..", ".$"}
      multiCharDelimitMap[".."] = {"..."}
      multiCharDelimitMap[","] = {",,"}
      multiCharDelimitMap[",,"] = {",,,"}
      multiCharDelimitMap[",,,"] = {",,,,"}
      multiCharDelimitMap["@"] = {"@@"}
      multiCharDelimitMap["@@"] = {"@@@", "@@="}
      multiCharDelimitMap["#"] = {"##"}
      multiCharDelimitMap["*"] = {"**"}
   else
    
      multiCharDelimitMap["."] = {".."}
      multiCharDelimitMap["~"] = {"~="}
   end
   
   
   return keywordSet, typeSet, builtInSet, multiCharDelimitMap
end

local TxtStream = {}
setmetatable( TxtStream, { ifList = {iStream,} } )
_moduleObj.TxtStream = TxtStream
function TxtStream.new( txt )
   local obj = {}
   TxtStream.setmeta( obj )
   if obj.__init then obj:__init( txt ); end
   return obj
end
function TxtStream:__init(txt) 
   self.txt = txt
   self.start = 1
   self.eof = false
end
function TxtStream:get_pos(  )

   return self.start
end
function TxtStream:read( mode )

   
   if self.eof then
      return nil
   end
   
   local index = self.txt:find( "\n", self.start, true )
   do
      local _exp = index
      if _exp ~= nil then
         local txt = self.txt:sub( self.start, _exp - 1 )
         self.start = _exp + 1
         return txt
      end
   end
   
   self.eof = true
   return self.txt:sub( self.start )
end
function TxtStream:close(  )

end
function TxtStream:getHead(  )

   return self.txt:sub( 1, self.start )
end
function TxtStream:getTail(  )

   return self.txt:sub( self.start )
end
function TxtStream.setmeta( obj )
  setmetatable( obj, { __index = TxtStream  } )
end
function TxtStream:get_txt()
   return self.txt
end


local Position = {}
_moduleObj.Position = Position
function Position.setmeta( obj )
  setmetatable( obj, { __index = Position  } )
end
function Position.new( lineNo, column )
   local obj = {}
   Position.setmeta( obj )
   if obj.__init then
      obj:__init( lineNo, column )
   end
   return obj
end
function Position:__init( lineNo, column )

   self.lineNo = lineNo
   self.column = column
end


local TokenKind = {}
_moduleObj.TokenKind = TokenKind
TokenKind._val2NameMap = {}
function TokenKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "TokenKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function TokenKind._from( val )
   if TokenKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
TokenKind.__allList = {}
function TokenKind.get__allList()
   return TokenKind.__allList
end

TokenKind.Cmnt = 0
TokenKind._val2NameMap[0] = 'Cmnt'
TokenKind.__allList[1] = TokenKind.Cmnt
TokenKind.Str = 1
TokenKind._val2NameMap[1] = 'Str'
TokenKind.__allList[2] = TokenKind.Str
TokenKind.Int = 2
TokenKind._val2NameMap[2] = 'Int'
TokenKind.__allList[3] = TokenKind.Int
TokenKind.Real = 3
TokenKind._val2NameMap[3] = 'Real'
TokenKind.__allList[4] = TokenKind.Real
TokenKind.Char = 4
TokenKind._val2NameMap[4] = 'Char'
TokenKind.__allList[5] = TokenKind.Char
TokenKind.Symb = 5
TokenKind._val2NameMap[5] = 'Symb'
TokenKind.__allList[6] = TokenKind.Symb
TokenKind.Dlmt = 6
TokenKind._val2NameMap[6] = 'Dlmt'
TokenKind.__allList[7] = TokenKind.Dlmt
TokenKind.Kywd = 7
TokenKind._val2NameMap[7] = 'Kywd'
TokenKind.__allList[8] = TokenKind.Kywd
TokenKind.Ope = 8
TokenKind._val2NameMap[8] = 'Ope'
TokenKind.__allList[9] = TokenKind.Ope
TokenKind.Type = 9
TokenKind._val2NameMap[9] = 'Type'
TokenKind.__allList[10] = TokenKind.Type
TokenKind.Eof = 10
TokenKind._val2NameMap[10] = 'Eof'
TokenKind.__allList[11] = TokenKind.Eof


local Token = {}
_moduleObj.Token = Token
function Token.new( kind, txt, pos, consecutive, commentList )
   local obj = {}
   Token.setmeta( obj )
   if obj.__init then obj:__init( kind, txt, pos, consecutive, commentList ); end
   return obj
end
function Token:__init(kind, txt, pos, consecutive, commentList) 
   self.kind = kind
   self.txt = txt
   self.pos = pos
   self.consecutive = consecutive
   self.commentList = _lune.unwrapDefault( commentList, {})
end
function Token:getExcludedDelimitTxt(  )

   if self.kind ~= TokenKind.Str then
      return self.txt
   end
   
   do
      local _switchExp = string.byte( self.txt, 1 )
      if _switchExp == 39 or _switchExp == 34 then
         return self.txt:sub( 2, #self.txt - 1 )
      elseif _switchExp == 96 then
         return self.txt:sub( 1 + 3, #self.txt - 3 )
      end
   end
   
   error( string.format( "illegal delimit -- %s", self.txt) )
end
function Token:set_commentList( commentList )

   self.commentList = commentList
end
function Token.setmeta( obj )
  setmetatable( obj, { __index = Token  } )
end
function Token:get_commentList()
   return self.commentList
end


local Parser = {}
_moduleObj.Parser = Parser
function Parser.setmeta( obj )
  setmetatable( obj, { __index = Parser  } )
end
function Parser.new(  )
   local obj = {}
   Parser.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Parser:__init(  )

end


local PushbackParser = {}
_moduleObj.PushbackParser = PushbackParser
function PushbackParser.setmeta( obj )
  setmetatable( obj, { __index = PushbackParser  } )
end
function PushbackParser.new(  )
   local obj = {}
   PushbackParser.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function PushbackParser:__init(  )

end


local WrapParser = {}
setmetatable( WrapParser, { __index = Parser } )
_moduleObj.WrapParser = WrapParser
function WrapParser:getToken(  )

   local token = self.parser:getToken(  )
   return token
end
function WrapParser:getStreamName(  )

   return self.name
end
function WrapParser.setmeta( obj )
  setmetatable( obj, { __index = WrapParser  } )
end
function WrapParser.new( parser, name )
   local obj = {}
   WrapParser.setmeta( obj )
   if obj.__init then
      obj:__init( parser, name )
   end
   return obj
end
function WrapParser:__init( parser, name )

   Parser.__init( self)
   self.parser = parser
   self.name = name
end


local noneToken = Token.new(TokenKind.Eof, "", Position.new(0, -1), false, {})
_moduleObj.noneToken = noneToken


local StreamParser = {}
setmetatable( StreamParser, { __index = Parser } )
_moduleObj.StreamParser = StreamParser
function StreamParser.setStdinStream( moduleName )

   StreamParser.stdinStreamModuleName = moduleName
   StreamParser.stdinTxt = _lune.unwrapDefault( io.stdin:read( '*a' ), "")
end
function StreamParser.new( stream, name, luaMode )
   local obj = {}
   StreamParser.setmeta( obj )
   if obj.__init then obj:__init( stream, name, luaMode ); end
   return obj
end
function StreamParser:__init(stream, name, luaMode) 
   Parser.__init( self)
   
   self.eof = false
   self.stream = stream
   self.streamName = name
   self.lineNo = 0
   self.pos = 1
   self.lineTokenList = {}
   self.prevToken = _moduleObj.noneToken
   
   local keywordSet, typeSet, builtInSet, multiCharDelimitMap = createReserveInfo( luaMode )
   
   self.keywordSet = keywordSet
   self.typeSet = typeSet
   self.builtInSet = builtInSet
   self.multiCharDelimitMap = multiCharDelimitMap
end
function StreamParser:getStreamName(  )

   return self.streamName
end
function StreamParser.create( path, luaMode, moduleName )

   local stream = TxtStream.new(StreamParser.stdinTxt)
   if StreamParser.stdinStreamModuleName ~= moduleName then
      stream = io.open( path, "r" )
      if  nil == stream then
         local _stream = stream
      
         return nil
      end
      
   end
   
   
   return StreamParser.new(stream, path, luaMode or string.find( path, "%.lua$" ) and true)
end
function StreamParser.setmeta( obj )
  setmetatable( obj, { __index = StreamParser  } )
end
do
   StreamParser.stdinStreamModuleName = nil
   StreamParser.stdinTxt = ""
end


local DefaultPushbackParser = {}
setmetatable( DefaultPushbackParser, { ifList = {PushbackParser,} } )
_moduleObj.DefaultPushbackParser = DefaultPushbackParser
function DefaultPushbackParser.new( parser )
   local obj = {}
   DefaultPushbackParser.setmeta( obj )
   if obj.__init then obj:__init( parser ); end
   return obj
end
function DefaultPushbackParser:__init(parser) 
   self.parser = parser
   self.pushbackedList = {}
   self.usedTokenList = {}
   self.currentToken = _moduleObj.noneToken
end
function DefaultPushbackParser:getTokenNoErr(  )

   if #self.pushbackedList > 0 then
      self.currentToken = self.pushbackedList[#self.pushbackedList]
      table.remove( self.pushbackedList )
   else
    
      do
         local token = self.parser:getToken(  )
         if token ~= nil then
            self.currentToken = token
         else
            self.currentToken = _moduleObj.noneToken
         end
      end
      
      
   end
   
   if self.currentToken.kind ~= TokenKind.Eof then
      table.insert( self.usedTokenList, self.currentToken )
   end
   
   return self.currentToken
end
function DefaultPushbackParser:pushbackToken( token )

   if token.kind ~= TokenKind.Eof then
      table.insert( self.pushbackedList, token )
   end
   
   if token == self.currentToken then
      table.remove( self.usedTokenList )
      if #self.usedTokenList > 0 then
         self.currentToken = self.usedTokenList[#self.usedTokenList]
      else
       
         self.currentToken = _moduleObj.noneToken
      end
      
   end
   
end
function DefaultPushbackParser:pushback(  )

   self:pushbackToken( self.currentToken )
end
function DefaultPushbackParser:pushbackStr( name, statement )

   local memStream = TxtStream.new(statement)
   local parser = StreamParser.new(memStream, name, false)
   
   local list = {}
   while true do
      do
         local _exp = parser:getToken(  )
         if _exp ~= nil then
            table.insert( list, _exp )
         else
            break
         end
      end
      
   end
   
   for index = #list, 1, -1 do
      self:pushbackToken( list[index] )
   end
   
end
function DefaultPushbackParser:newPushback( tokenList )

   for index = #tokenList, 1, -1 do
      self:pushbackToken( tokenList[index] )
   end
   
end
function DefaultPushbackParser:error( message )

   Util.err( message )
end
function DefaultPushbackParser.setmeta( obj )
  setmetatable( obj, { __index = DefaultPushbackParser  } )
end


local quotedCharSet = {}
quotedCharSet['a']= true
quotedCharSet['b']= true
quotedCharSet['f']= true
quotedCharSet['n']= true
quotedCharSet['r']= true
quotedCharSet['t']= true
quotedCharSet['v']= true
quotedCharSet['\\']= true
quotedCharSet['"']= true
quotedCharSet["'"]= true

local op2Set = {}
op2Set['+']= true
op2Set['-']= true
op2Set['*']= true
op2Set['/']= true
op2Set['^']= true
op2Set['%']= true
op2Set['&']= true
op2Set['~']= true
op2Set['|']= true
op2Set['|>>']= true
op2Set['|<<']= true
op2Set['..']= true
op2Set['<']= true
op2Set['<=']= true
op2Set['>']= true
op2Set['>=']= true
op2Set['==']= true
op2Set['~=']= true
op2Set['and']= true
op2Set['or']= true
op2Set['@']= true
op2Set['@@']= true
op2Set['@@@']= true
op2Set['=']= true

local op1Set = {}
op1Set['-']= true
op1Set['not']= true
op1Set['#']= true
op1Set['~']= true
op1Set['*']= true
op1Set['`']= true
op1Set[',,']= true
op1Set[',,,']= true
op1Set[',,,,']= true

local function isOp2( ope )

   return _lune._Set_has(op2Set, ope )
end
_moduleObj.isOp2 = isOp2

local function isOp1( ope )

   return _lune._Set_has(op1Set, ope )
end
_moduleObj.isOp1 = isOp1

function StreamParser:parse(  )

   local function readLine(  )
   
      if self.eof then
         return nil
      end
      
      self.lineNo = self.lineNo + 1
      local line = self.stream:read( '*l' )
      if not line then
         self.eof = true
      end
      
      return line
   end
   local rawLine = readLine(  )
   if  nil == rawLine then
      local _rawLine = rawLine
   
      return nil
   end
   
   
   local list = {}
   local startIndex = 1
   
   local function multiComment( comIndex, termStr )
   
      local searchIndex = comIndex
      local comment = ""
      while true do
         do
            local termIndex, termEndIndex = string.find( rawLine, termStr, searchIndex, true )
            if termIndex ~= nil and termEndIndex ~= nil then
               comment = comment .. rawLine:sub( searchIndex, termEndIndex )
               return comment, termEndIndex + 1
            end
         end
         
         comment = comment .. rawLine:sub( searchIndex ) .. "\n"
         searchIndex = 1
         rawLine = _lune.unwrap( readLine(  ))
      end
      
   end
   
   
   local function addVal( kind, val, column )
   
      local function createInfo( tokenKind, token, tokenColumn )
      
         if tokenKind == TokenKind.Symb then
            if _lune._Set_has(self.keywordSet, token ) then
               tokenKind = TokenKind.Kywd
            elseif _lune._Set_has(self.typeSet, token ) then
               tokenKind = TokenKind.Type
            elseif _lune._Set_has(op2Set, token ) or _lune._Set_has(op1Set, token ) then
               tokenKind = TokenKind.Ope
            end
            
         end
         
         local consecutive = false
         if self.prevToken.pos.lineNo == self.lineNo and self.prevToken.pos.column + #self.prevToken.txt == tokenColumn then
            consecutive = true
         end
         
         local newToken = Token.new(tokenKind, token, Position.new(self.lineNo, tokenColumn), consecutive, {})
         self.prevToken = newToken
         return newToken
      end
      
      local function analyzeNumber( token, beginIndex )
      
         local nonNumIndex = token:find( '[^%d]', beginIndex )
         if  nil == nonNumIndex then
            local _nonNumIndex = nonNumIndex
         
            return #token, true
         end
         
         local intFlag = true
         local nonNumChar = token:byte( nonNumIndex )
         if nonNumChar == 46 then
            intFlag = false
            nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
            if  nil == nonNumIndex then
               local _nonNumIndex = nonNumIndex
            
               return #token, intFlag
            end
            
            nonNumChar = token:byte( nonNumIndex )
         end
         
         if nonNumChar == 88 or nonNumChar == 120 then
            nonNumIndex = token:find( '[^%da-fA-F]', nonNumIndex + 1 )
            if  nil == nonNumIndex then
               local _nonNumIndex = nonNumIndex
            
               return #token, intFlag
            end
            
            nonNumChar = token:byte( nonNumIndex )
         end
         
         if nonNumChar == 69 or nonNumChar == 101 then
            intFlag = false
            local nextChar = token:byte( nonNumIndex + 1 )
            if nextChar == 45 or nextChar == 43 then
               nonNumIndex = token:find( '[^%d]', nonNumIndex + 2 )
               if  nil == nonNumIndex then
                  local _nonNumIndex = nonNumIndex
               
                  return #token, intFlag
               end
               
            else
             
               nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
               if  nil == nonNumIndex then
                  local _nonNumIndex = nonNumIndex
               
                  return #token, intFlag
               end
               
            end
            
         end
         
         return nonNumIndex - 1, intFlag
      end
      
      if kind == TokenKind.Symb then
         local searchIndex = 1
         while true do
            
            local tokenIndex, tokenEndIndex = string.find( val, "[%p%w]+", searchIndex )
            if  nil == tokenIndex or  nil == tokenEndIndex then
               local _tokenIndex = tokenIndex
               local _tokenEndIndex = tokenEndIndex
            
               break
            end
            
            local columnIndex = column + tokenIndex - 2
            searchIndex = tokenEndIndex + 1
            local token = val:sub( tokenIndex, tokenEndIndex )
            local subIndex = 1
            while true do
               if token:find( '^[%d]', subIndex ) or token:find( '^-[%d]', subIndex ) then
                  
                  local checkIndex = subIndex
                  if string.byte( token, 1 ) == 45 then
                     checkIndex = checkIndex + 1
                  end
                  
                  local endIndex, intFlag = analyzeNumber( token, checkIndex )
                  local info = createInfo( intFlag and TokenKind.Int or TokenKind.Real, token:sub( subIndex, endIndex ), columnIndex + subIndex )
                  table.insert( list, info )
                  subIndex = endIndex + 1
               else
                
                  
                  do
                     local _exp = string.find( token, '[^%w_]', subIndex )
                     if _exp ~= nil then
                        local index = _exp
                        if index > subIndex then
                           local info = createInfo( TokenKind.Symb, token:sub( subIndex, index - 1 ), columnIndex + subIndex )
                           table.insert( list, info )
                        end
                        
                        local delimit = token:sub( index, index )
                        local candidateList = self.multiCharDelimitMap[delimit]
                        while candidateList do
                           local findFlag = false
                           for __index, candidate in pairs( _lune.unwrap( (candidateList )) ) do
                              if candidate == token:sub( index, index + #candidate - 1 ) then
                                 delimit = candidate
                                 candidateList = self.multiCharDelimitMap[delimit]
                                 findFlag = true
                                 break
                              end
                              
                           end
                           
                           if not findFlag then
                              break
                           end
                           
                        end
                        
                        subIndex = index + #delimit
                        
                        local workKind = TokenKind.Dlmt
                        if _lune._Set_has(op2Set, delimit ) or _lune._Set_has(op1Set, delimit ) then
                           workKind = TokenKind.Ope
                        end
                        
                        if delimit == "..." then
                           workKind = TokenKind.Symb
                        end
                        
                        if delimit == "?" then
                           local nextChar = token:sub( index, subIndex )
                           table.insert( list, createInfo( TokenKind.Char, nextChar, columnIndex + subIndex ) )
                           subIndex = subIndex + 1
                        else
                         
                           table.insert( list, createInfo( workKind, delimit, columnIndex + index ) )
                        end
                        
                     else
                        if subIndex <= #token then
                           table.insert( list, createInfo( TokenKind.Symb, token:sub( subIndex ), columnIndex + subIndex ) )
                        end
                        
                        break
                     end
                  end
                  
               end
               
            end
            
         end
         
      else
       
         table.insert( list, createInfo( kind, val, column ) )
      end
      
   end
   
   
   local searchIndex = startIndex
   
   while true do
      local syncIndexFlag = true
      
      local pattern = [==[[/%-%?"%'%`].]==]
      local index = string.find( rawLine, pattern, searchIndex )
      if  nil == index then
         local _index = index
      
         addVal( TokenKind.Symb, rawLine:sub( startIndex ), startIndex )
         return list
      end
      
      
      local findChar = string.byte( rawLine, index )
      local nextChar = string.byte( rawLine, index + 1 )
      
      if findChar == 45 and nextChar ~= 45 then
         searchIndex = index + 1
         syncIndexFlag = false
      else
       
         if startIndex < index then
            addVal( TokenKind.Symb, rawLine:sub( startIndex, index - 1 ), startIndex )
         end
         
         if findChar == 47 then
            
            if nextChar == 42 then
               
               local comment, nextIndex = multiComment( index + 2, "*/" )
               addVal( TokenKind.Cmnt, "/*" .. comment, index )
               searchIndex = nextIndex
            elseif nextChar == 47 then
               
               addVal( TokenKind.Cmnt, rawLine:sub( index ), index )
               searchIndex = #rawLine + 1
            else
             
               addVal( TokenKind.Ope, "/", index )
               searchIndex = index + 1
            end
            
         elseif findChar == 39 or findChar == 34 then
            
            local workIndex = index + 1
            local workPattern = '["\'\\]'
            while true do
               local endIndex = string.find( rawLine, workPattern, workIndex )
               if  nil == endIndex then
                  local _endIndex = endIndex
               
                  
                  Util.err( string.format( "%s:%d:%d: error: illegal string -- %s", self:getStreamName(  ), self.lineNo, index, rawLine) )
               end
               
               local workChar = string.byte( rawLine, endIndex )
               if workChar == findChar then
                  
                  addVal( TokenKind.Str, rawLine:sub( index, endIndex ), index )
                  searchIndex = endIndex + 1
                  break
               elseif workChar == 92 then
                  
                  workIndex = endIndex + 2
               else
                
                  workIndex = endIndex + 1
               end
               
            end
            
         elseif findChar == 96 then
            if (nextChar == findChar and string.byte( rawLine, index + 2 ) == 96 ) then
               
               local txt, nextIndex = multiComment( index + 3, '```' )
               addVal( TokenKind.Str, '```' .. txt, index )
               searchIndex = nextIndex
            else
             
               addVal( TokenKind.Ope, '`', index )
               searchIndex = index + 1
            end
            
         elseif findChar == 63 then
            local codeChar = rawLine:sub( index + 1, index + 1 )
            if nextChar == 92 then
               local quoted = rawLine:sub( index + 2, index + 2 )
               if _lune._Set_has(quotedCharSet, quoted ) then
                  codeChar = rawLine:sub( index + 1, index + 2 )
               else
                
                  codeChar = quoted
               end
               
               searchIndex = index + 3
            else
             
               searchIndex = index + 2
            end
            
            addVal( TokenKind.Char, codeChar, index )
         else
          
            Util.err( string.format( "%s:%d:%d: error: illegal syntax -- %s", self:getStreamName(  ), self.lineNo, index, rawLine) )
         end
         
      end
      
      if syncIndexFlag then
         startIndex = searchIndex
      end
      
   end
   
end


function StreamParser:getToken(  )

   if #self.lineTokenList < self.pos then
      self.pos = 1
      self.lineTokenList = {}
      while #self.lineTokenList == 0 do
         local workList = self:parse(  )
         if  nil == workList then
            local _workList = workList
         
            return nil
         end
         
         self.lineTokenList = workList
      end
      
   end
   
   
   local token = self.lineTokenList[self.pos]
   self.pos = self.pos + 1
   
   return token
end


local eofToken = Token.new(TokenKind.Eof, "<EOF>", Position.new(0, 0), false, {})
local function getEofToken(  )

   return eofToken
end
_moduleObj.getEofToken = getEofToken
local DummyParser = {}
setmetatable( DummyParser, { __index = Parser } )
_moduleObj.DummyParser = DummyParser
function DummyParser:getToken(  )

   return eofToken
end
function DummyParser:getStreamName(  )

   return "dummy"
end
function DummyParser.setmeta( obj )
  setmetatable( obj, { __index = DummyParser  } )
end
function DummyParser.new(  )
   local obj = {}
   DummyParser.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function DummyParser:__init(  )

   Parser.__init( self)
end


local CommentLayer = {}
_moduleObj.CommentLayer = CommentLayer
function CommentLayer.new(  )
   local obj = {}
   CommentLayer.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function CommentLayer:__init() 
   self.commentList = {}
   self.tokenSet = {}
   self.tokenList = {}
end
function CommentLayer:addDirect( commentList )

   for __index, comment in pairs( commentList ) do
      table.insert( self.commentList, comment )
   end
   
end
function CommentLayer:add( token )

   if not _lune._Set_has(self.tokenSet, token ) then
      self.tokenSet[token]= true
      table.insert( self.tokenList, token )
      
      self:addDirect( token:get_commentList() )
   end
   
end
function CommentLayer:clear(  )

   if #self.commentList ~= 0 then
      self.commentList = {}
      self.tokenSet = {}
      self.tokenList = {}
   end
   
end
function CommentLayer:hasInvalidComment(  )

   return #self.tokenList > 1 and self.tokenList[2]:get_commentList()[1] or nil
end
function CommentLayer.setmeta( obj )
  setmetatable( obj, { __index = CommentLayer  } )
end
function CommentLayer:get_commentList()
   return self.commentList
end


local CommentCtrl = {}
_moduleObj.CommentCtrl = CommentCtrl
function CommentCtrl.new(  )
   local obj = {}
   CommentCtrl.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function CommentCtrl:__init() 
   self.layer = CommentLayer.new()
   self.layerStack = {self.layer}
end
function CommentCtrl:push(  )

   self.layer = CommentLayer.new()
   table.insert( self.layerStack, self.layer )
end
function CommentCtrl:pop(  )

   self.layer = self.layerStack[#self.layerStack]
   table.remove( self.layerStack )
end
function CommentCtrl.setmeta( obj )
  setmetatable( obj, { __index = CommentCtrl  } )
end
function CommentCtrl:addDirect( ... )
   return self.layer:addDirect( ... )
end

function CommentCtrl:add( ... )
   return self.layer:add( ... )
end

function CommentCtrl:clear( ... )
   return self.layer:clear( ... )
end

function CommentCtrl:hasInvalidComment( ... )
   return self.layer:hasInvalidComment( ... )
end

function CommentCtrl:get_commentList( ... )
   return self.layer:get_commentList( ... )
end





return _moduleObj
