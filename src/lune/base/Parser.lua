--lune/base/Parser.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Parser'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

if not _lune3 then
   _lune3 = _lune
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
      Util.err( string.format( "not support -- %s", mode) )
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


local Position = Types.Position
_moduleObj.Position = Position
local TokenKind = Types.TokenKind
_moduleObj.TokenKind = TokenKind
local Token = Types.Token
_moduleObj.Token = Token

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
   
   
   self.streamName = name
   self.pos = 1
   self.lineTokenList = {}
   
   self.asyncParser = AsyncParser.Parser.new(stream, name, luaMode)
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
   self.currentToken = Types.noneToken
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
function DefaultPushbackParser.setmeta( obj )
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
local eofToken = Token.new(Types.TokenKind.Eof, "<EOF>", Position.new(0, 0, "eof"), false, {})
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
