--lune/base/Parser.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Parser'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
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

if not _lune8 then
   _lune8 = _lune
end



local Util = _lune.loadModule( 'lune.base.Util' )
local Str = _lune.loadModule( 'lune.base.Str' )
local Types = _lune.loadModule( 'lune.base.Types' )
local Async = _lune.loadModule( 'lune.base.Async' )
local AsyncTokenizer = _lune.loadModule( 'lune.base.AsyncParser' )

local function isLuaKeyword( txt )

   return AsyncTokenizer.isLuaKeyword( txt )
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

local Tokenizer = {}
_moduleObj.Tokenizer = Tokenizer
function Tokenizer._setmeta( obj )
  setmetatable( obj, { __index = Tokenizer  } )
end
function Tokenizer._new(  )
   local obj = {}
   Tokenizer._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Tokenizer:__init(  )

end


local PushbackTokenizer = {}
_moduleObj.PushbackTokenizer = PushbackTokenizer
function PushbackTokenizer._setmeta( obj )
  setmetatable( obj, { __index = PushbackTokenizer  } )
end
function PushbackTokenizer._new(  )
   local obj = {}
   PushbackTokenizer._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function PushbackTokenizer:__init(  )

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

local TokenListTokenizer = {}
setmetatable( TokenListTokenizer, { __index = Tokenizer } )
_moduleObj.TokenListTokenizer = TokenListTokenizer
function TokenListTokenizer._new( tokenList, streamName, overridePos )
   local obj = {}
   TokenListTokenizer._setmeta( obj )
   if obj.__init then obj:__init( tokenList, streamName, overridePos ); end
   return obj
end
function TokenListTokenizer:__init(tokenList, streamName, overridePos) 
   Tokenizer.__init( self)
   
   
   self.index = 1
   self.tokenList = tokenList
   self.streamName = streamName
   self.overridePos = overridePos
end
function TokenListTokenizer:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), self.overridePos )
end
function TokenListTokenizer:getStreamName(  )

   return self.streamName
end
function TokenListTokenizer:getToken(  )

   if #self.tokenList < self.index then
      return nil
   end
   
   local token = self.tokenList[self.index]
   self.index = self.index + 1
   
   return token
end
function TokenListTokenizer._setmeta( obj )
  setmetatable( obj, { __index = TokenListTokenizer  } )
end


local StreamTokenizer = {}
setmetatable( StreamTokenizer, { __index = Tokenizer } )
_moduleObj.StreamTokenizer = StreamTokenizer
function StreamTokenizer.setStdinStream( moduleName )

   StreamTokenizer.stdinStreamModuleName = moduleName
   StreamTokenizer.stdinTxt = _lune.unwrapDefault( io.stdin:read( '*a' ), "")
end
function StreamTokenizer._new( tokenizerSrc, async, stdinFile, pos )
   local obj = {}
   StreamTokenizer._setmeta( obj )
   if obj.__init then obj:__init( tokenizerSrc, async, stdinFile, pos ); end
   return obj
end
function StreamTokenizer:__init(tokenizerSrc, async, stdinFile, pos) 
   Tokenizer.__init( self)
   
   
   self.pos = 1
   self.lineTokenList = {}
   self.overridePos = pos
   
   local asyncTokenizer, errMess = AsyncTokenizer.create( tokenizerSrc, stdinFile, pos, async )
   do
      local _exp = asyncTokenizer
      if _exp ~= nil then
         self.asyncTokenizer = _exp
      else
         Util.err( errMess )
         
      end
   end
   
   self.streamName = self.asyncTokenizer:get_streamName()
end
function StreamTokenizer:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), self.overridePos )
end
function StreamTokenizer:getStreamName(  )

   return self.streamName
end
function StreamTokenizer.create( tokenizerSrc, async, stdinFile, pos )

   return StreamTokenizer._new(tokenizerSrc, async, stdinFile, pos)
end
function StreamTokenizer:getToken(  )

   if #self.lineTokenList < self.pos then
      self.pos = 1
      self.lineTokenList = {}
      while #self.lineTokenList == 0 do
         local pipeItem = self.asyncTokenizer:getNext(  )
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
function StreamTokenizer._setmeta( obj )
  setmetatable( obj, { __index = StreamTokenizer  } )
end
do
   StreamTokenizer.stdinStreamModuleName = nil
   StreamTokenizer.stdinTxt = ""
end


local DefaultPushbackTokenizer = {}
setmetatable( DefaultPushbackTokenizer, { ifList = {PushbackTokenizer,} } )
_moduleObj.DefaultPushbackTokenizer = DefaultPushbackTokenizer
function DefaultPushbackTokenizer._new( tokenizer )
   local obj = {}
   DefaultPushbackTokenizer._setmeta( obj )
   if obj.__init then obj:__init( tokenizer ); end
   return obj
end
function DefaultPushbackTokenizer:__init(tokenizer) 
   self.tokenizer = tokenizer
   self.pushbackedList = {}
   self.usedTokenList = {}
   self.currentToken = Types.noneToken
end
function DefaultPushbackTokenizer:getUsedTokenListLen(  )

   return #self.usedTokenList
end
function DefaultPushbackTokenizer.createFromLnsCode( code, name )

   return DefaultPushbackTokenizer._new(StreamTokenizer._new(_lune.newAlge( Types.TokenizerSrc.LnsCode, {code,name,nil}), false))
end
function DefaultPushbackTokenizer:createPosition( lineNo, column )

   return self.tokenizer:createPosition( lineNo, column )
end
function DefaultPushbackTokenizer:getTokenNoErr( skipFlag )

   if #self.pushbackedList > 0 then
      self.currentToken = self.pushbackedList[#self.pushbackedList]
      table.remove( self.pushbackedList )
   else
    
      do
         local token = self.tokenizer:getToken(  )
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
function DefaultPushbackTokenizer:pushbackToken( token )

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
function DefaultPushbackTokenizer:pushback(  )

   self:pushbackToken( self.currentToken )
end
function DefaultPushbackTokenizer:pushbackStr( asyncParse, name, statement, pos )

   local tokenizer = StreamTokenizer._new(_lune.newAlge( Types.TokenizerSrc.LnsCode, {statement,name,nil}), asyncParse == true, nil, pos)
   
   local list = {}
   while true do
      do
         local _exp = tokenizer:getToken(  )
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
function DefaultPushbackTokenizer:newPushback( tokenList )

   for index = #tokenList, 1, -1 do
      self:pushbackToken( tokenList[index] )
   end
   
end
function DefaultPushbackTokenizer:getLastPos(  )

   local pos = self.tokenizer:createPosition( 0, 0 )
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
function DefaultPushbackTokenizer:getNearCode(  )

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
function DefaultPushbackTokenizer:getStreamName(  )

   return self.tokenizer:getStreamName(  )
end
function DefaultPushbackTokenizer._setmeta( obj )
  setmetatable( obj, { __index = DefaultPushbackTokenizer  } )
end
function DefaultPushbackTokenizer:get_currentToken()
   return self.currentToken
end


local function isOp2( ope )

   return AsyncTokenizer.isOp2( ope )
end
_moduleObj.isOp2 = isOp2

local function isOp1( ope )

   return AsyncTokenizer.isOp1( ope )
end
_moduleObj.isOp1 = isOp1
local eofToken = Token._new(Types.TokenKind.Eof, "<EOF>", Position._new(0, 0, "eof"), false, {})
local function getEofToken(  )

   return eofToken
end
_moduleObj.getEofToken = getEofToken
local DummyTokenizer = {}
setmetatable( DummyTokenizer, { __index = Tokenizer } )
_moduleObj.DummyTokenizer = DummyTokenizer
function DummyTokenizer:getToken(  )

   return eofToken
end
function DummyTokenizer:getStreamName(  )

   return "dummy"
end
function DummyTokenizer:createPosition( lineNo, column )

   return Position.create( lineNo, column, self:getStreamName(  ), nil )
end
function DummyTokenizer._setmeta( obj )
  setmetatable( obj, { __index = DummyTokenizer  } )
end
function DummyTokenizer._new(  )
   local obj = {}
   DummyTokenizer._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function DummyTokenizer:__init(  )

   Tokenizer.__init( self)
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

local function createTokenizerFrom( src, async, stdinFile )

   return StreamTokenizer._new(src, async, stdinFile, nil)
end
_moduleObj.createTokenizerFrom = createTokenizerFrom



return _moduleObj
