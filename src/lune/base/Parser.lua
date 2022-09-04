--lune/base/Parser.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Parser'
local _lune = {}
if _lune7 then
   _lune = _lune7
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
   if __luneScript and not package.preload[ mod ] then
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

if not _lune7 then
   _lune7 = _lune
end


local Util = _lune.loadModule( 'lune.base.Util' )
local Str = _lune.loadModule( 'lune.base.Str' )
local Types = _lune.loadModule( 'lune.base.Types' )
local Async = _lune.loadModule( 'lune.base.Async' )
local AsyncParser = _lune.loadModule( 'lune.base.AsyncParser' )

local function isLuaKeyword( txt )

   return AsyncParser.isLuaKeyword( txt )
end
_moduleObj.isLuaKeyword = isLuaKeyword

local TxtStream = Util.TxtStream
_moduleObj.TxtStream = TxtStream
local Position = Types.Position
_moduleObj.Position = Position
local TokenKind = Types.TokenKind
_moduleObj.TokenKind = TokenKind
local Token = Types.Token
_moduleObj.Token = Token

local Parser = {}
_moduleObj.Parser = Parser
function Parser._setmeta( obj )
  setmetatable( obj, { __index = Parser  } )
end
function Parser._new(  )
   local obj = {}
   Parser._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Parser:__init(  )

end


local PushbackParser = {}
_moduleObj.PushbackParser = PushbackParser
function PushbackParser._setmeta( obj )
  setmetatable( obj, { __index = PushbackParser  } )
end
function PushbackParser._new(  )
   local obj = {}
   PushbackParser._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function PushbackParser:__init(  )

end


local noneToken = Types.noneToken
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

local TokenListParser = {}
setmetatable( TokenListParser, { __index = Parser } )
_moduleObj.TokenListParser = TokenListParser
function TokenListParser._new( tokenList, streamName, overridePos )
   local obj = {}
   TokenListParser._setmeta( obj )
   if obj.__init then obj:__init( tokenList, streamName, overridePos ); end
   return obj
end
function TokenListParser:__init(tokenList, streamName, overridePos) 
   Parser.__init( self)
   
   
   self.index = 1
   self.tokenList = tokenList
   self.streamName = streamName
   self.overridePos = overridePos
end
function TokenListParser:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), self.overridePos )
end
function TokenListParser:getStreamName(  )

   return self.streamName
end
function TokenListParser:getToken(  )

   if #self.tokenList < self.index then
      return nil
   end
   
   local token = self.tokenList[self.index]
   self.index = self.index + 1
   
   return token
end
function TokenListParser._setmeta( obj )
  setmetatable( obj, { __index = TokenListParser  } )
end


local StreamParser = {}
setmetatable( StreamParser, { __index = Parser } )
_moduleObj.StreamParser = StreamParser
function StreamParser.setStdinStream( moduleName )

   StreamParser.stdinStreamModuleName = moduleName
   StreamParser.stdinTxt = _lune.unwrapDefault( io.stdin:read( '*a' ), "")
end
function StreamParser._new( parserSrc, async, stdinFile, pos )
   local obj = {}
   StreamParser._setmeta( obj )
   if obj.__init then obj:__init( parserSrc, async, stdinFile, pos ); end
   return obj
end
function StreamParser:__init(parserSrc, async, stdinFile, pos) 
   Parser.__init( self)
   
   
   self.pos = 1
   self.lineTokenList = {}
   self.overridePos = pos
   
   local asyncParser, errMess = AsyncParser.create( parserSrc, stdinFile, pos, async )
   do
      local _exp = asyncParser
      if _exp ~= nil then
         self.asyncParser = _exp
      else
         Util.err( errMess )
      end
   end
   
   self.streamName = self.asyncParser:get_streamName()
end
function StreamParser:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), self.overridePos )
end
function StreamParser:getStreamName(  )

   return self.streamName
end
function StreamParser.create( parserSrc, async, stdinFile, pos )

   return StreamParser._new(parserSrc, async, stdinFile, pos)
end
function StreamParser:getToken(  )

   if #self.lineTokenList < self.pos then
      self.pos = 1
      self.lineTokenList = {}
      while #self.lineTokenList == 0 do
         local pipeItem = self.asyncParser:getNext(  )
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
function StreamParser._setmeta( obj )
  setmetatable( obj, { __index = StreamParser  } )
end
do
   StreamParser.stdinStreamModuleName = nil
   StreamParser.stdinTxt = ""
end


local DefaultPushbackParser = {}
setmetatable( DefaultPushbackParser, { ifList = {PushbackParser,} } )
_moduleObj.DefaultPushbackParser = DefaultPushbackParser
function DefaultPushbackParser._new( parser )
   local obj = {}
   DefaultPushbackParser._setmeta( obj )
   if obj.__init then obj:__init( parser ); end
   return obj
end
function DefaultPushbackParser:__init(parser) 
   self.parser = parser
   self.pushbackedList = {}
   self.usedTokenList = {}
   self.currentToken = Types.noneToken
end
function DefaultPushbackParser:getUsedTokenListLen(  )

   return #self.usedTokenList
end
function DefaultPushbackParser.createFromLnsCode( code, name )

   return DefaultPushbackParser._new(StreamParser._new(_lune.newAlge( Types.ParserSrc.LnsCode, {code,name,nil}), false))
end
function DefaultPushbackParser:createPosition( lineNo, column )

   return self.parser:createPosition( lineNo, column )
end
function DefaultPushbackParser:getTokenNoErr( skipFlag )

   if #self.pushbackedList > 0 then
      self.currentToken = self.pushbackedList[#self.pushbackedList]
      table.remove( self.pushbackedList )
   else
    
      do
         local token = self.parser:getToken(  )
         if token ~= nil then
            self.currentToken = token
         else
            self.currentToken = Types.noneToken
         end
      end
      
      
   end
   
   if self.currentToken.kind ~= Types.TokenKind.Eof then
      table.insert( self.usedTokenList, self.currentToken )
   end
   
   return self.currentToken
end
function DefaultPushbackParser:pushbackToken( token )

   if token.kind ~= Types.TokenKind.Eof then
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
          
            self.currentToken = Types.noneToken
         end
         
      else
       
         self.currentToken = Types.noneToken
      end
      
   end
   
end
function DefaultPushbackParser:pushback(  )

   self:pushbackToken( self.currentToken )
end
function DefaultPushbackParser:pushbackStr( asyncParse, name, statement, pos )

   local parser = StreamParser._new(_lune.newAlge( Types.ParserSrc.LnsCode, {statement,name,nil}), asyncParse == true, nil, pos)
   
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
   if self:get_currentToken().kind ~= Types.TokenKind.Eof then
      pos = self:get_currentToken().pos
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
function DefaultPushbackParser._setmeta( obj )
  setmetatable( obj, { __index = DefaultPushbackParser  } )
end
function DefaultPushbackParser:get_currentToken()
   return self.currentToken
end


local function isOp2( ope )

   return AsyncParser.isOp2( ope )
end
_moduleObj.isOp2 = isOp2

local function isOp1( ope )

   return AsyncParser.isOp1( ope )
end
_moduleObj.isOp1 = isOp1
local eofToken = Token._new(Types.TokenKind.Eof, "<EOF>", Position._new(0, 0, "eof"), false, {})
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
function DummyParser:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), nil )
end
function DummyParser._setmeta( obj )
  setmetatable( obj, { __index = DummyParser  } )
end
function DummyParser._new(  )
   local obj = {}
   DummyParser._setmeta( obj )
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
function CommentLayer._new(  )
   local obj = {}
   CommentLayer._setmeta( obj )
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
function CommentLayer._setmeta( obj )
  setmetatable( obj, { __index = CommentLayer  } )
end
function CommentLayer:get_commentList()
   return self.commentList
end


local CommentCtrl = {}
_moduleObj.CommentCtrl = CommentCtrl
function CommentCtrl._new(  )
   local obj = {}
   CommentCtrl._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function CommentCtrl:__init() 
   self.layer = CommentLayer._new()
   self.layerStack = {self.layer}
end
function CommentCtrl:push(  )

   self.layer = CommentLayer._new()
   table.insert( self.layerStack, self.layer )
end
function CommentCtrl:pop(  )

   self.layer = self.layerStack[#self.layerStack]
   table.remove( self.layerStack )
end
function CommentCtrl._setmeta( obj )
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

local function createParserFrom( src, async, stdinFile )

   return StreamParser._new(src, async, stdinFile, nil)
end
_moduleObj.createParserFrom = createParserFrom



return _moduleObj
