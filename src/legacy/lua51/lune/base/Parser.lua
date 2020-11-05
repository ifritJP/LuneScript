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

function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
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
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
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
local Str = _lune.loadModule( 'lune.base.Str' )
local Async = _lune.loadModule( 'lune.base.Async' )
local Types = _lune.loadModule( 'lune.base.Types' )

local luaKeywordSet = {["if"] = true, ["else"] = true, ["elseif"] = true, ["while"] = true, ["for"] = true, ["in"] = true, ["return"] = true, ["break"] = true, ["nil"] = true, ["true"] = true, ["false"] = true, ["{"] = true, ["}"] = true, ["do"] = true, ["require"] = true, ["function"] = true, ["then"] = true, ["end"] = true, ["repeat"] = true, ["until"] = true, ["goto"] = true, ["local"] = true}

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
   self.lineList = Str.getLineList( self.txt )
   self.lineNo = 1
end
function TxtStream:getSubstring( fromLineNo, toLineNo )

   local txt = ""
   local to = _lune.unwrapDefault( toLineNo, #self.lineList + 1)
   for index = fromLineNo, to - 1 do
      if index < 1 or index > #self.lineList then
         break
      end
      
      txt = string.format( "%s%s", txt, self.lineList[index])
   end
   
   return txt
end
function TxtStream:read( mode )

   if mode ~= '*l' then
      Util.err( string.format( "not support -- %s", tostring( mode)) )
   end
   
   if self.lineNo > #self.lineList then
      return nil
   end
   
   self.lineNo = self.lineNo + 1
   local line = self.lineList[self.lineNo - 1]
   if Str.endsWith( line, "\n" ) then
      return line:sub( 1, #line - 1 )
   end
   
   return line
end
function TxtStream:close(  )

end
function TxtStream.setmeta( obj )
  setmetatable( obj, { __index = TxtStream  } )
end
function TxtStream:get_txt()
   return self.txt
end
function TxtStream:get_lineNo()
   return self.lineNo
end


local Position = {}
setmetatable( Position, { ifList = {Mapping,} } )
_moduleObj.Position = Position
function Position.setmeta( obj )
  setmetatable( obj, { __index = Position  } )
end
function Position.new( lineNo, column, streamName )
   local obj = {}
   Position.setmeta( obj )
   if obj.__init then
      obj:__init( lineNo, column, streamName )
   end
   return obj
end
function Position:__init( lineNo, column, streamName )

   self.lineNo = lineNo
   self.column = column
   self.streamName = streamName
end
function Position:_toMap()
  return self
end
function Position._fromMap( val )
  local obj, mes = Position._fromMapSub( {}, val )
  if obj then
     Position.setmeta( obj )
  end
  return obj, mes
end
function Position._fromStem( val )
  return Position._fromMap( val )
end

function Position._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "lineNo", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "column", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "streamName", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
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
setmetatable( Token, { ifList = {Mapping,} } )
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
function Token:getLineCount(  )

   local count = 1
   for _301 in self.txt:gmatch( "\n" ) do
      count = count + 1
   end
   
   return count
end
function Token.setmeta( obj )
  setmetatable( obj, { __index = Token  } )
end
function Token:get_commentList()
   return self.commentList
end
function Token:_toMap()
  return self
end
function Token._fromMap( val )
  local obj, mes = Token._fromMapSub( {}, val )
  if obj then
     Token.setmeta( obj )
  end
  return obj, mes
end
function Token._fromStem( val )
  return Token._fromMap( val )
end

function Token._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "kind", func = TokenKind._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "pos", func = Position._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "consecutive", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "commentList", func = _lune._toList, nilable = false, child = { { func = Token._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local Parser = {}
_moduleObj.Parser = Parser
function Parser:createPosition( lineNo, column )

   return Position.new(lineNo, column, self:getStreamName(  ))
end
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


local noneToken = Token.new(TokenKind.Eof, "", Position.new(0, -1, "eof"), false, {})
_moduleObj.noneToken = noneToken

local function convFromRawToStr( txt )

   if #txt == 0 then
      return txt
   end
   
   do
      local _switchExp = string.byte( txt, 1 )
      if _switchExp == 39 or _switchExp == 34 then
      else 
         
            return txt:sub( 4, #txt - 3 )
      end
   end
   
   local findChar = string.byte( txt, 1 )
   local workTxt = txt
   local retTxt = ""
   local workIndex = 2
   local setIndex = 2
   local workPattern = "[\"'\\]"
   while true do
      local endIndex = string.find( workTxt, workPattern, workIndex )
      if  nil == endIndex then
         local _endIndex = endIndex
      
         Util.err( string.format( "error: illegal string -- %s", workTxt) )
      end
      
      local workChar = string.byte( workTxt, endIndex )
      if workChar == findChar then
         return retTxt .. workTxt:sub( setIndex, endIndex - 1 )
      elseif workChar == 92 then
         local quote = string.byte( workTxt, endIndex + 1 )
         do
            local _switchExp = quote
            if _switchExp == 39 or _switchExp == 34 then
               retTxt = string.format( "%s%s%c", retTxt, workTxt:sub( setIndex, endIndex - 1 ), quote)
            else 
               
                  retTxt = string.format( "%s%s", retTxt, workTxt:sub( setIndex, endIndex + 1 ))
            end
         end
         
         workIndex = endIndex + 2
         setIndex = workIndex
      else
       
         workIndex = endIndex + 1
      end
      
   end
   
end
_moduleObj.convFromRawToStr = convFromRawToStr

local StreamParser = {}

local AsyncItem = {}
setmetatable( AsyncItem, { ifList = {__AsyncItem,Mapping,} } )
function AsyncItem.setmeta( obj )
  setmetatable( obj, { __index = AsyncItem  } )
end
function AsyncItem.new( list )
   local obj = {}
   AsyncItem.setmeta( obj )
   if obj.__init then
      obj:__init( list )
   end
   return obj
end
function AsyncItem:__init( list )

   self.list = list
end
function AsyncItem:_toMap()
  return self
end
function AsyncItem._fromMap( val )
  local obj, mes = AsyncItem._fromMapSub( {}, val )
  if obj then
     AsyncItem.setmeta( obj )
  end
  return obj, mes
end
function AsyncItem._fromStem( val )
  return AsyncItem._fromMap( val )
end

function AsyncItem._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "list", func = _lune._toList, nilable = false, child = { { func = Token._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local AsyncStreamParser = {}
setmetatable( AsyncStreamParser, { __index = Async.Pipe } )
function AsyncStreamParser.new( pipe, streamParser )
   local obj = {}
   AsyncStreamParser.setmeta( obj )
   if obj.__init then obj:__init( pipe, streamParser ); end
   return obj
end
function AsyncStreamParser:__init(pipe, streamParser) 
   Async.Pipe.__init( self,pipe)
   
   self.streamParser = streamParser
end
function AsyncStreamParser.setmeta( obj )
  setmetatable( obj, { __index = AsyncStreamParser  } )
end


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
   
   self.pipe = AsyncStreamParser.new(nil, self)
   
   self.eof = false
   self.stream = stream
   self.streamName = name
   self.lineNo = 0
   self.pos = 1
   self.lineTokenList = {}
   self.prevToken = _moduleObj.noneToken
   self.luaMode = _lune.unwrapDefault( luaMode, false)
   
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

   local stream
   
   if StreamParser.stdinStreamModuleName ~= moduleName then
      stream = io.open( path, "r" )
      if  nil == stream then
         local _stream = stream
      
         return nil
      end
      
   else
    
      stream = TxtStream.new(StreamParser.stdinTxt)
   end
   
   
   return StreamParser.new(stream, path, luaMode or Str.endsWith( path, ".lua" ) and true)
end
function StreamParser.setmeta( obj )
  setmetatable( obj, { __index = StreamParser  } )
end
do
   StreamParser.stdinStreamModuleName = nil
   StreamParser.stdinTxt = ""
end


function AsyncStreamParser:access(  )

   local tokenList = self.streamParser:parse(  )
   if  nil == tokenList then
      local _tokenList = tokenList
   
      return nil
   end
   
   return Async.PipeItem.new(AsyncItem.new(tokenList))
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
function DefaultPushbackParser:createPosition( lineNo, column )

   return self.parser:createPosition( lineNo, column )
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
      if #self.usedTokenList > 0 then
         local used = self.usedTokenList[#self.usedTokenList]
         if used == token then
            table.remove( self.usedTokenList )
         end
         
         if #self.usedTokenList > 0 then
            self.currentToken = self.usedTokenList[#self.usedTokenList]
         else
          
            self.currentToken = _moduleObj.noneToken
         end
         
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
function DefaultPushbackParser:getLastPos(  )

   local pos = self.parser:createPosition( 0, 0 )
   if self.currentToken.kind ~= TokenKind.Eof then
      pos = self.currentToken.pos
   else
    
      if #self.usedTokenList > 0 then
         local token = self.usedTokenList[#self.usedTokenList]
         pos = token.pos
      end
      
   end
   
   return pos
   
end
function DefaultPushbackParser:getNearCode(  )

   local code = ""
   for index = #self.usedTokenList - 30, #self.usedTokenList do
      if index > 1 then
         local token = self.usedTokenList[index]
         if token.consecutive then
            code = string.format( "%s%s", code, self.usedTokenList[index].txt)
         else
          
            code = string.format( "%s %s", code, self.usedTokenList[index].txt)
         end
         
      end
      
   end
   
   return string.format( "%s -- current '%s'", code, self.currentToken.txt)
end
function DefaultPushbackParser:getStreamName(  )

   return self.parser:getStreamName(  )
end
function DefaultPushbackParser.setmeta( obj )
  setmetatable( obj, { __index = DefaultPushbackParser  } )
end
function DefaultPushbackParser:get_currentToken()
   return self.currentToken
end


local quotedCharSet = {['a'] = true, ['b'] = true, ['f'] = true, ['n'] = true, ['r'] = true, ['t'] = true, ['v'] = true, ['\\'] = true, ['"'] = true, ["'"] = true}

local op2Set = {['+'] = true, ['-'] = true, ['*'] = true, ['/'] = true, ['^'] = true, ['%'] = true, ['&'] = true, ['~'] = true, ['|'] = true, ['|>>'] = true, ['|<<'] = true, ['..'] = true, ['<'] = true, ['<='] = true, ['>'] = true, ['>='] = true, ['=='] = true, ['~='] = true, ['and'] = true, ['or'] = true, ['@'] = true, ['@@'] = true, ['@@@'] = true, ['='] = true}

local op1Set = {['-'] = true, ['not'] = true, ['#'] = true, ['~'] = true, ['*'] = true, ['`'] = true, [',,'] = true, [',,,'] = true, [',,,,'] = true}

local function isOp2( ope )

   return _lune._Set_has(op2Set, ope )
end
_moduleObj.isOp2 = isOp2

local function isOp1( ope )

   return _lune._Set_has(op1Set, ope )
end
_moduleObj.isOp1 = isOp1

local TokenList = {}
function TokenList.setmeta( obj )
  setmetatable( obj, { __index = TokenList  } )
end
function TokenList.new( list )
   local obj = {}
   TokenList.setmeta( obj )
   if obj.__init then
      obj:__init( list )
   end
   return obj
end
function TokenList:__init( list )

   self.list = list
end


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
            local _589, termEndIndex = string.find( rawLine, termStr, searchIndex, true )
            if termEndIndex ~= nil then
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
         
         local newToken = Token.new(tokenKind, token, self:createPosition( self.lineNo, tokenColumn ), consecutive, {})
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
         local nonNumChar = string.byte( token, nonNumIndex )
         if nonNumChar == 46 then
            intFlag = false
            nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
            if  nil == nonNumIndex then
               local _nonNumIndex = nonNumIndex
            
               return #token, intFlag
            end
            
            nonNumChar = string.byte( token, nonNumIndex )
         end
         
         if nonNumChar == 88 or nonNumChar == 120 then
            nonNumIndex = token:find( '[^%da-fA-F]', nonNumIndex + 1 )
            if  nil == nonNumIndex then
               local _nonNumIndex = nonNumIndex
            
               return #token, intFlag
            end
            
            nonNumChar = string.byte( token, nonNumIndex )
         end
         
         if nonNumChar == 69 or nonNumChar == 101 then
            intFlag = false
            local nextChar = string.byte( token, nonNumIndex + 1 )
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
            while #token >= subIndex do
               if token:find( '^[%d]', subIndex ) or string.byte( token, subIndex ) == 45 and token:find( '^[%d]', subIndex + 1 ) then
                  local checkIndex = subIndex
                  if string.byte( token, checkIndex ) == 45 then
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
   
   local function getChar( index )
   
      if #rawLine >= index then
         return string.byte( rawLine, index )
      end
      
      return 0
   end
   while true do
      local syncIndexFlag = true
      local pattern = [==[[/%-%?"%'%`].]==]
      local index = string.find( rawLine, pattern, searchIndex )
      if  nil == index then
         local _index = index
      
         addVal( TokenKind.Symb, rawLine:sub( startIndex ), startIndex )
         return list
      end
      
      local findChar = getChar( index )
      local nextChar = getChar( index + 1 )
      
      if findChar == 45 and nextChar ~= 45 then
         searchIndex = index + 1
         syncIndexFlag = false
      else
       
         if startIndex < index then
            addVal( TokenKind.Symb, rawLine:sub( startIndex, index - 1 ), startIndex )
         end
         
         if findChar == 39 or findChar == 34 then
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
          
            if self.luaMode and findChar == 45 and nextChar == 45 then
               addVal( TokenKind.Cmnt, rawLine:sub( index ), index )
               searchIndex = #rawLine + 1
               
            elseif findChar == 47 then
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
               
            else
             
               Util.err( string.format( "%s:%d:%d: error: illegal syntax -- %s", self:getStreamName(  ), self.lineNo, index, rawLine) )
            end
            
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
         local pipeItem = self.pipe:getNext(  )
         if  nil == pipeItem then
            local _pipeItem = pipeItem
         
            return nil
         end
         
         self.lineTokenList = pipeItem:get_item().list
      end
      
   end
   
   
   local token = self.lineTokenList[self.pos]
   self.pos = self.pos + 1
   
   return token
end


local eofToken = Token.new(TokenKind.Eof, "<EOF>", Position.new(0, 0, "eof"), false, {})
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

   for __index, comment in ipairs( commentList ) do
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
function CommentCtrl:add( ... )
   return self.layer:add( ... )
end

function CommentCtrl:addDirect( ... )
   return self.layer:addDirect( ... )
end

function CommentCtrl:clear( ... )
   return self.layer:clear( ... )
end

function CommentCtrl:get_commentList( ... )
   return self.layer:get_commentList( ... )
end

function CommentCtrl:hasInvalidComment( ... )
   return self.layer:hasInvalidComment( ... )
end


local function quoteStr( txt )

   local work = txt
   local part = '"'
   for index = 1, #work do
      local char = string.byte( work, index )
      do
         local _switchExp = char
         if _switchExp == 10 then
            part = part .. "\\n"
         elseif _switchExp == 13 then
            part = part .. "\\r"
         elseif _switchExp == 9 then
            part = part .. "\\t"
         elseif _switchExp == 34 then
            part = part .. '\\"'
         elseif _switchExp == 92 then
            part = part .. '\\\\'
         else 
            
               part = part .. string.format( "%c", char)
         end
      end
      
   end
   
   work = part .. '"'
   return work
end
_moduleObj.quoteStr = quoteStr



return _moduleObj
