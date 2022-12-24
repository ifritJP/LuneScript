--lune/base/AsyncParser.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@AsyncParser'
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

function _lune._run( runner, mod )
    if mod == 2 then
      return false
    end
    runner:run()
    return true
end

if not _lune7 then
   _lune7 = _lune
end


local Util = _lune.loadModule( 'lune.base.Util' )
local Types = _lune.loadModule( 'lune.base.Types' )
local Async = _lune.loadModule( 'lune.base.Async' )

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
      keywordSet["__request"]= true
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

local quotedCharSet = {['a'] = true, ['b'] = true, ['f'] = true, ['n'] = true, ['r'] = true, ['t'] = true, ['v'] = true, ['\\'] = true, ['"'] = true, ["'"] = true}

local op2Set = {['+'] = true, ['-'] = true, ['*'] = true, ['/'] = true, ['^'] = true, ['%'] = true, ['&'] = true, ['~'] = true, ['|'] = true, ['|>>'] = true, ['|<<'] = true, ['..'] = true, ['<'] = true, ['<='] = true, ['>'] = true, ['>='] = true, ['=='] = true, ['~='] = true, ['and'] = true, ['or'] = true, ['@'] = true, ['@@'] = true, ['@@@'] = true, ['='] = true}

local op1Set = {['-'] = true, ['not'] = true, ['#'] = true, ['~'] = true, ['`{'] = true, [',,'] = true, [',,,'] = true, [',,,,'] = true}

local function isOp2( ope )

   return _lune._Set_has(op2Set, ope )
end
_moduleObj.isOp2 = isOp2

local function isOp1( ope )

   return _lune._Set_has(op1Set, ope )
end
_moduleObj.isOp1 = isOp1

local AsyncItem = {}
setmetatable( AsyncItem, { ifList = {__AsyncItem,Mapping,} } )
_moduleObj.AsyncItem = AsyncItem
function AsyncItem._setmeta( obj )
  setmetatable( obj, { __index = AsyncItem  } )
end
function AsyncItem._new( list )
   local obj = {}
   AsyncItem._setmeta( obj )
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
     AsyncItem._setmeta( obj )
  end
  return obj, mes
end
function AsyncItem._fromStem( val )
  return AsyncItem._fromMap( val )
end

function AsyncItem._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "list", func = _lune._toList, nilable = false, child = { { func = Types.Token._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local defaultPipeSize = 100
local function setDefaultPipeSize( size )

   defaultPipeSize = size
end
_moduleObj.setDefaultPipeSize = setDefaultPipeSize

local Parser = {}
setmetatable( Parser, { __index = Async.Pipe } )
_moduleObj.Parser = Parser
function Parser._new( streamName, stream, luaMode, overridePos, pipeSize )
   local obj = {}
   Parser._setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, luaMode, overridePos, pipeSize ); end
   return obj
end
function Parser:__init(streamName, stream, luaMode, overridePos, pipeSize) 
   local _
   Async.Pipe.__init( self,nil)
   
   
   self.stream = stream
   self.lineList = {}
   self.streamName = streamName
   self.overridePos = overridePos
   self.firstLine = true
   self.lineNo = 0
   self.prevToken = Types.noneToken
   self.luaMode = luaMode
   
   local keywordSet, typeSet, _1, multiCharDelimitMap = createReserveInfo( luaMode )
   
   self.keywordSet = keywordSet
   self.typeSet = typeSet
   self.multiCharDelimitMap = multiCharDelimitMap
end
function Parser:setup(  )

   local lineList = {}
   while true do
      local line = self.stream:read( '*l' )
      if  nil == line then
         local _line = line
      
         break
      end
      
      table.insert( lineList, line )
   end
   
   self.lineList = lineList
   self.stream:close(  )
end
function Parser.create( parserSrc, stdinFile, overridePos )

   local function createStream( mod, path )
   
      if stdinFile ~= nil then
         if stdinFile:get_mod() == mod then
            return Util.TxtStream._new(stdinFile:get_txt()), ""
         end
         
      end
      
      do
         local _exp = io.open( path, "r" )
         if _exp ~= nil then
            return _exp, ""
         end
      end
      
      return nil, string.format( "failed to open -- %s", path)
   end
   
   local function createStreamFrom(  )
   
      do
         local _matchExp = parserSrc
         if _matchExp[1] == Types.ParserSrc.LnsCode[1] then
            local txt = _matchExp[2][1]
            local path = _matchExp[2][2]
            local pipeSize = _matchExp[2][3]
         
            return path, false, Util.TxtStream._new(txt), "", pipeSize
         elseif _matchExp[1] == Types.ParserSrc.LnsPath[1] then
            local path = _matchExp[2][1]
            local mod = _matchExp[2][2]
            local pipeSize = _matchExp[2][3]
         
            local stream, mess = createStream( mod, path )
            return path, false, stream, mess, pipeSize
         elseif _matchExp[1] == Types.ParserSrc.Parser[1] then
            local path = _matchExp[2][1]
            local luaMode = _matchExp[2][2]
            local mod = _matchExp[2][3]
            local pipeSize = _matchExp[2][4]
         
            local stream, mess = createStream( mod, path )
            return path, luaMode, stream, mess, pipeSize
         end
      end
      
   end
   
   local streamName, luaMode, stream, mess, pipeSize = createStreamFrom(  )
   if stream ~= nil then
      return Parser._new(streamName, stream, luaMode, overridePos, pipeSize), ""
   end
   
   return nil, mess
end
function Parser:access(  )

   local tokenList = self:parse(  )
   if  nil == tokenList then
      local _tokenList = tokenList
   
      return nil
   end
   
   return Async.PipeItem._new(AsyncItem._new(tokenList))
end
function Parser._setmeta( obj )
  setmetatable( obj, { __index = Parser  } )
end
function Parser:get_streamName()
   return self.streamName
end


local Runner = {}
setmetatable( Runner, { ifList = {__Runner,} } )
function Runner._new( parserSrc, stdinFile, overridePos )
   local obj = {}
   Runner._setmeta( obj )
   if obj.__init then obj:__init( parserSrc, stdinFile, overridePos ); end
   return obj
end
function Runner:__init(parserSrc, stdinFile, overridePos) 
   self.parser, self.errMess = Parser.create( parserSrc, stdinFile, overridePos )
   
   do
      local _exp = self.parser
      if _exp ~= nil then
         _exp:start(  )
         if not _lune._run(self, 2, string.format( "parser - %s", _exp:get_streamName()) ) then
            _exp:stop(  )
         end
         
         
      end
   end
   
end
function Runner:run(  )

   do
      local _exp = self.parser
      if _exp ~= nil then
         _exp:run(  )
      end
   end
   
end
function Runner._setmeta( obj )
  setmetatable( obj, { __index = Runner  } )
end
function Runner:get_parser()
   return self.parser
end
function Runner:get_errMess()
   return self.errMess
end


local function create( parserSrc, stdinFile, overridePos, async )

   if async then
      local runner = Runner._new(parserSrc, stdinFile, overridePos)
      return runner:get_parser(), runner:get_errMess()
   end
   
   local parser, mess = Parser.create( parserSrc, stdinFile, overridePos )
   if parser ~= nil then
      
      parser:stop(  )
   end
   
   return parser, mess
end
_moduleObj.create = create

function Parser:createInfo( tokenKind, token, tokenColumn )

   if tokenKind == Types.TokenKind.Symb then
      if _lune._Set_has(self.keywordSet, token ) then
         tokenKind = Types.TokenKind.Kywd
      elseif _lune._Set_has(self.typeSet, token ) then
         tokenKind = Types.TokenKind.Type
      elseif _lune._Set_has(op2Set, token ) or _lune._Set_has(op1Set, token ) then
         tokenKind = Types.TokenKind.Ope
      end
      
   end
   
   local consecutive = false
   if self.prevToken.pos.lineNo == self.lineNo and self.prevToken.pos.column + #self.prevToken.txt == tokenColumn then
      consecutive = true
   end
   
   local newToken = Types.Token._new(tokenKind, token, Types.Position.create( self.lineNo, tokenColumn, self.streamName, self.overridePos ), consecutive, {})
   self.prevToken = newToken
   return newToken
end


function Parser:analyzeNumber( token, beginIndex )

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


function Parser:readLine(  )

   if self.lineNo >= #self.lineList then
      return nil
   end
   
   self.lineNo = self.lineNo + 1
   return self.lineList[self.lineNo]
end

function Parser:addVal( list, kind, val, column )

   if kind ~= Types.TokenKind.Symb then
      table.insert( list, self:createInfo( kind, val, column ) )
      return 
   end
   
   
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
            
            local endIndex, intFlag = self:analyzeNumber( token, checkIndex )
            local info = self:createInfo( intFlag and Types.TokenKind.Int or Types.TokenKind.Real, token:sub( subIndex, endIndex ), columnIndex + subIndex )
            table.insert( list, info )
            subIndex = endIndex + 1
         else
          
            
            do
               local _exp = string.find( token, '[^%w_]', subIndex )
               if _exp ~= nil then
                  local index = _exp
                  if index > subIndex then
                     local info = self:createInfo( Types.TokenKind.Symb, token:sub( subIndex, index - 1 ), columnIndex + subIndex )
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
                  
                  local workKind = Types.TokenKind.Dlmt
                  if _lune._Set_has(op2Set, delimit ) or _lune._Set_has(op1Set, delimit ) then
                     workKind = Types.TokenKind.Ope
                  end
                  
                  if delimit == "..." then
                     workKind = Types.TokenKind.Symb
                  end
                  
                  if delimit == "?" then
                     local nextChar = token:sub( index, subIndex )
                     table.insert( list, self:createInfo( Types.TokenKind.Char, nextChar, columnIndex + subIndex ) )
                     subIndex = subIndex + 1
                  else
                   
                     table.insert( list, self:createInfo( workKind, delimit, columnIndex + index ) )
                  end
                  
               else
                  if subIndex <= #token then
                     table.insert( list, self:createInfo( Types.TokenKind.Symb, token:sub( subIndex ), columnIndex + subIndex ) )
                  end
                  
                  break
               end
            end
            
         end
         
      end
      
   end
   
end



function Parser:parse(  )

   local rawLine = self:readLine(  )
   if  nil == rawLine then
      local _rawLine = rawLine
   
      return nil
   end
   
   if self.firstLine then
      self.firstLine = false
      if rawLine:find( "^#!" ) then
         local token = Types.Token._new(Types.TokenKind.Sheb, rawLine, Types.Position._new(self.lineNo, 1, self.streamName), false, {})
         return {token}
      end
      
   end
   
   
   local function multiComment( comIndex, termStr )
   
      local searchIndex = comIndex
      local comment = ""
      while true do
         do
            local _1, termEndIndex = string.find( rawLine, termStr, searchIndex, true )
            if termEndIndex ~= nil then
               comment = comment .. rawLine:sub( searchIndex, termEndIndex )
               return comment, termEndIndex + 1
            end
         end
         
         comment = comment .. rawLine:sub( searchIndex ) .. "\n"
         searchIndex = 1
         rawLine = _lune.unwrap( self:readLine(  ))
      end
      
   end
   
   
   local startIndex = 1
   
   local searchIndex = startIndex
   
   local function getChar( index )
   
      if #rawLine >= index then
         return string.byte( rawLine, index )
      end
      
      return 0
   end
   
   local list = {}
   
   while true do
      local syncIndexFlag = true
      local pattern = [==[[/%-%?"%'%`].]==]
      local index = string.find( rawLine, pattern, searchIndex )
      if  nil == index then
         local _index = index
      
         
         self:addVal( list, Types.TokenKind.Symb, rawLine:sub( startIndex ), startIndex )
         return list
      end
      
      
      local findChar = getChar( index )
      local nextChar = getChar( index + 1 )
      
      if findChar == 45 and nextChar ~= 45 then
         searchIndex = index + 1
         syncIndexFlag = false
      else
       
         if startIndex < index then
            self:addVal( list, Types.TokenKind.Symb, rawLine:sub( startIndex, index - 1 ), startIndex )
         end
         
         if findChar == 39 or findChar == 34 then
            
            local workIndex = index + 1
            local workPattern = '["\'\\]'
            while true do
               local endIndex = string.find( rawLine, workPattern, workIndex )
               if  nil == endIndex then
                  local _endIndex = endIndex
               
                  
                  Util.err( string.format( "%s:%d:%d: error: illegal string -- %s", self.streamName, self.lineNo, index, rawLine) )
               end
               
               local workChar = string.byte( rawLine, endIndex )
               if workChar == findChar then
                  
                  self:addVal( list, Types.TokenKind.Str, rawLine:sub( index, endIndex ), index )
                  searchIndex = endIndex + 1
                  break
               elseif workChar == 92 then
                  
                  workIndex = endIndex + 2
               else
                
                  workIndex = endIndex + 1
               end
               
            end
            
         elseif findChar == 96 then
            if (nextChar == findChar and getChar( index + 2 ) == 96 ) then
               
               local txt, nextIndex = multiComment( index + 3, '```' )
               self:addVal( list, Types.TokenKind.Str, '```' .. txt, index )
               searchIndex = nextIndex
            elseif nextChar == 123 then
               self:addVal( list, Types.TokenKind.Ope, '`{', index )
               searchIndex = index + 2
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
            
            self:addVal( list, Types.TokenKind.Char, codeChar, index )
         else
          
            if self.luaMode and findChar == 45 and nextChar == 45 then
               
               self:addVal( list, Types.TokenKind.Cmnt, rawLine:sub( index ), index )
               searchIndex = #rawLine + 1
            elseif findChar == 47 then
               
               if nextChar == 42 then
                  
                  local comment, nextIndex = multiComment( index + 2, "*/" )
                  self:addVal( list, Types.TokenKind.Cmnt, "/*" .. comment, index )
                  searchIndex = nextIndex
               elseif nextChar == 47 then
                  
                  self:addVal( list, Types.TokenKind.Cmnt, rawLine:sub( index ), index )
                  searchIndex = #rawLine + 1
               else
                
                  self:addVal( list, Types.TokenKind.Ope, "/", index )
                  searchIndex = index + 1
               end
               
            else
             
               Util.err( string.format( "%s:%d:%d: error: illegal syntax -- %s", self.streamName, self.lineNo, index, rawLine) )
            end
            
         end
         
      end
      
      if syncIndexFlag then
         startIndex = searchIndex
      end
      
   end
   
end


return _moduleObj
