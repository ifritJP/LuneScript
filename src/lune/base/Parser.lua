--lune/base/Parser.lns
local moduleObj = {}
local function createReserveInfo( luaMode )
  local keywordSet = {}
  local typeSet = {}
  local builtInSet = {}
  keywordSet["let"] = true
  keywordSet["if"] = true
  keywordSet["else"] = true
  keywordSet["elseif"] = true
  keywordSet["while"] = true
  keywordSet["for"] = true
  keywordSet["in"] = true
  keywordSet["return"] = true
  keywordSet["break"] = true
  keywordSet["nil"] = true
  keywordSet["true"] = true
  keywordSet["false"] = true
  keywordSet["{"] = true
  keywordSet["}"] = true
  builtInSet["require"] = true
  if luaMode then
    keywordSet["function"] = true
    keywordSet["then"] = true
    keywordSet["do"] = true
    keywordSet["until"] = true
  else 
    keywordSet["let"] = true
    keywordSet["mut"] = true
    keywordSet["pub"] = true
    keywordSet["pro"] = true
    keywordSet["pri"] = true
    keywordSet["fn"] = true
    keywordSet["each"] = true
    keywordSet["form"] = true
    keywordSet["class"] = true
    builtInSet["super"] = true
    keywordSet["static"] = true
    keywordSet["advertise"] = true
    keywordSet["import"] = true
    keywordSet["new"] = true
    keywordSet["!"] = true
    keywordSet["unwrap"] = true
    typeSet["int"] = true
    typeSet["real"] = true
    typeSet["stem"] = true
    typeSet["str"] = true
    typeSet["Map"] = true
    typeSet["bool"] = true
  end
  local multiCharDelimitMap = {}
  multiCharDelimitMap["="] = {"=="}
  multiCharDelimitMap["~"] = {"~="}
  multiCharDelimitMap["<"] = {"<="}
  multiCharDelimitMap[">"] = {">="}
  if not luaMode then
    multiCharDelimitMap["."] = {"..", ".$"}
    multiCharDelimitMap[".."] = {"..."}
    multiCharDelimitMap[","] = {",,"}
    multiCharDelimitMap[",,"] = {",,,"}
    multiCharDelimitMap[",,,"] = {",,,,"}
    multiCharDelimitMap["@"] = {"@@"}
    multiCharDelimitMap["@@"] = {"@@?"}
  else 
    multiCharDelimitMap["."] = {".."}
  end
  return keywordSet, typeSet, builtInSet, multiCharDelimitMap
end

local Stream = {}
moduleObj.Stream = Stream
-- none
function Stream.new(  )
  local obj = {}
  setmetatable( obj, { __index = Stream } )
  if obj.__init then
    obj:__init(  )
  end        
  return obj 
 end         
function Stream:__init(  ) 
            
end

local TxtStream = {}
setmetatable( TxtStream, { __index = Stream } )
moduleObj.TxtStream = TxtStream
function TxtStream.new( txt )
  local obj = {}
  setmetatable( obj, { __index = TxtStream } )
  if obj.__init then obj:__init( txt ); end
return obj
end
function TxtStream:__init(txt) 
  self.txt = txt
  self.start = 1
  if not txt then
    error( "txt is nil" )
  end
  self.eof = false
end
function TxtStream:read( mode )
  if self.eof then
    return nil
  end
  local index = self.txt:find( "\n", self.start, true )
  if index then
    local txt = self.txt:sub( self.start, index - 1 )
    self.start = index + 1
    return txt
  end
  self.eof = true
  return self.txt:sub( self.start )
end

local Position = {}
moduleObj.Position = Position
function Position.new( lineNo, column )
  local obj = {}
  setmetatable( obj, { __index = Position } )
  if obj.__init then
    obj:__init( lineNo, column )
  end        
  return obj 
 end         
function Position:__init( lineNo, column ) 
            
self.lineNo = lineNo
  self.column = column
  end

local Token = {}
moduleObj.Token = Token
function Token:set_commentList( commentList )
  self.commentList = commentList
end
function Token.new( kind, txt, pos, commentList )
  local obj = {}
  setmetatable( obj, { __index = Token } )
  if obj.__init then
    obj:__init( kind, txt, pos, commentList )
  end        
  return obj 
 end         
function Token:__init( kind, txt, pos, commentList ) 
            
self.kind = kind
  self.txt = txt
  self.pos = pos
  self.commentList = commentList
  end
function Token:get_commentList()
  return self.commentList
end

local Parser = {}
moduleObj.Parser = Parser
-- none
-- none
function Parser.new(  )
  local obj = {}
  setmetatable( obj, { __index = Parser } )
  if obj.__init then
    obj:__init(  )
  end        
  return obj 
 end         
function Parser:__init(  ) 
            
end

local WrapParser = {}
setmetatable( WrapParser, { __index = Parser } )
moduleObj.WrapParser = WrapParser
function WrapParser:getToken(  )
  local token = self.parser:getToken(  )
  return token
end
function WrapParser:getStreamName(  )
  return self.name
end
function WrapParser.new( parser, name )
  local obj = {}
  setmetatable( obj, { __index = WrapParser } )
  if obj.__init then
    obj:__init( parser, name )
  end        
  return obj 
 end         
function WrapParser:__init( parser, name ) 
            
self.parser = parser
  self.name = name
  end

local StreamParser = {}
setmetatable( StreamParser, { __index = Parser } )
moduleObj.StreamParser = StreamParser
function StreamParser.new( stream, name, luaMode )
  local obj = {}
  setmetatable( obj, { __index = StreamParser } )
  if obj.__init then obj:__init( stream, name, luaMode ); end
return obj
end
function StreamParser:__init(stream, name, luaMode) 
  self.stream = stream
  self.streamName = name
  self.lineNo = 0
  self.pos = 1
  self.lineTokenList = {}
  local keywordSet, typeSet, builtInSet, multiCharDelimitMap = createReserveInfo( luaMode )
  self.keywordSet = keywordSet
  self.typeSet = typeSet
  self.builtInSet = builtInSet
  self.multiCharDelimitMap = multiCharDelimitMap
end
function StreamParser:getStreamName(  )
  return self.streamName
end
function StreamParser.create( path, luaMode )
  local stream = io.open( path, "r" )
  if not stream then
    return nil
  end
  return StreamParser.new(stream, path, luaMode or string.find( path, "%.lua$" ))
end

local kind = {}
moduleObj.kind = kind

local kindSeed = 0
local kind2Txt = {}
local function regKind( name )
  local assignKind = kindSeed
  kindSeed = kindSeed + 1
  kind2Txt[assignKind] = name
  kind[name] = assignKind
  return assignKind
end

local kindCmnt = regKind( "Cmnt" )
moduleObj.kindCmnt = kindCmnt

local kindStr = regKind( "Str" )
moduleObj.kindStr = kindStr

local kindInt = regKind( "Int" )
moduleObj.kindInt = kindInt

local kindReal = regKind( "Real" )
moduleObj.kindReal = kindReal

local kindChar = regKind( "Char" )
moduleObj.kindChar = kindChar

local kindSymb = regKind( "Symb" )
moduleObj.kindSymb = kindSymb

local kindDlmt = regKind( "Dlmt" )
moduleObj.kindDlmt = kindDlmt

local kindKywd = regKind( "Kywd" )
moduleObj.kindKywd = kindKywd

local kindOpe = regKind( "Ope" )
moduleObj.kindOpe = kindOpe

local kindType = regKind( "Type" )
moduleObj.kindType = kindType

local kindEof = regKind( "Eof" )
moduleObj.kindEof = kindEof

local noneToken = Token.new(kindEof, "", Position.new(0, 0))
moduleObj.noneToken = noneToken

local quotedCharSet = {}
quotedCharSet['a'] = true
quotedCharSet['b'] = true
quotedCharSet['f'] = true
quotedCharSet['n'] = true
quotedCharSet['r'] = true
quotedCharSet['t'] = true
quotedCharSet['v'] = true
quotedCharSet['\\'] = true
quotedCharSet['"'] = true
quotedCharSet["'"] = true
local op2Set = {}
op2Set['+'] = true
op2Set['-'] = true
op2Set['*'] = true
op2Set['/'] = true
op2Set['//'] = true
op2Set['^'] = true
op2Set['%'] = true
op2Set['&'] = true
op2Set['~'] = true
op2Set['|'] = true
op2Set['>>'] = true
op2Set['<<'] = true
op2Set['..'] = true
op2Set['<'] = true
op2Set['<='] = true
op2Set['>'] = true
op2Set['>='] = true
op2Set['=='] = true
op2Set['~='] = true
op2Set['and'] = true
op2Set['or'] = true
op2Set['@'] = true
op2Set['='] = true
local op1Set = {}
op1Set['-'] = true
op1Set['not'] = true
op1Set['#'] = true
op1Set['~'] = true
op1Set['*'] = true
op1Set['`'] = true
op1Set[',,'] = true
op1Set[',,,'] = true
op1Set[',,,,'] = true
local function getKindTxt( kind )
  return kind2Txt[kind] or _luneScript.error( 'unwrap val is nil' )
end
moduleObj.getKindTxt = getKindTxt
local function isOp2( ope )
  return op2Set[ope]
end
moduleObj.isOp2 = isOp2
local function isOp1( ope )
  return op1Set[ope]
end
moduleObj.isOp1 = isOp1
function StreamParser:parse(  )
  local function readLine(  )
    self.lineNo = self.lineNo + 1
    return self.stream:read( '*l' )
  end
  
  local rawLine = readLine(  )
  if  not rawLine then
  local _rawLine = rawLine
    
    return nil
  end
  
  local list = {}
  local startIndex = 1
  local multiComment = function ( comIndex, termStr )
    local searchIndex = comIndex
    local comment = ""
    while true do
      local termIndex, termEndIndex = string.find( rawLine, termStr, searchIndex, true )
      if termIndex then
        comment = comment .. rawLine:sub( searchIndex, termEndIndex )
        return comment, termEndIndex + 1
      end
      comment = comment .. rawLine:sub( searchIndex ) .. "\n"
      searchIndex = 1
      rawLine = readLine(  )
      if not rawLine then
        local _exp1 = rawLine
        
          error( "illegal comment" )
        end
      
      -- none
      
    end
  end
  
  local addVal = function ( kind, val, column )
    local function createInfo( tokenKind, token, tokenColumn )
      if tokenKind == kindSymb then
        if self.keywordSet[token] then
          tokenKind = kindKywd
        elseif self.typeSet[token] then
          tokenKind = kindType
        elseif op2Set[token] or op1Set[token] then
          tokenKind = kindOpe
        end
      end
      return Token.new(tokenKind, token, Position.new(self.lineNo, tokenColumn))
    end
    
    local function analyzeNumber( token, startIndex )
      local nonNumIndex = token:find( '[^%d]', startIndex )
      if not nonNumIndex then
        return #token, true
      end
      local intFlag = true
      local nonNumChar = token:byte( nonNumIndex )
      if nonNumChar == 46 then
        intFlag = false
        nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
        nonNumChar = token:byte( nonNumIndex )
      end
      if nonNumChar == 120 or nonNumChar == 88 then
        nonNumIndex = token:find( '[^%d]', nonNumIndex + 1 )
        nonNumChar = token:byte( nonNumIndex )
      end
      if nonNumChar == 101 or nonNumChar == 69 then
        intFlag = false
        local nextChar = token:byte( nonNumIndex + 1 )
        if nextChar == 45 or nextChar == 43 then
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
            local endIndex, intFlag = analyzeNumber( token, startIndex )
            local info = createInfo( intFlag and kindInt or kindReal, token:sub( startIndex, endIndex ), columnIndex + startIndex )
            table.insert( list, info )
            startIndex = endIndex + 1
          else 
            local index = string.find( token, '[^%w_]', startIndex )
            if index then
              if index > startIndex then
                local info = createInfo( kindSymb, token:sub( startIndex, index - 1 ), columnIndex + startIndex )
                table.insert( list, info )
              end
              local delimit = token:sub( index, index )
              local candidateList = self.multiCharDelimitMap[delimit]
              while candidateList do
                local findFlag = false
                for __index, candidate in pairs( (candidateList ) or _luneScript.error( 'unwrap val is nil' ) ) do
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
              startIndex = index + #delimit
              local workKind = kindDlmt
              if op2Set[delimit] or op1Set[delimit] then
                workKind = kindOpe
              end
              if delimit == "..." then
                workKind = kindSymb
              end
              if delimit == "?" then
                local nextChar = token:sub( index, startIndex )
                table.insert( list, createInfo( kindChar, nextChar, columnIndex + startIndex ) )
                startIndex = startIndex + 1
              else 
                table.insert( list, createInfo( workKind, delimit, columnIndex + index ) )
              end
            else 
              if startIndex <= #token then
                table.insert( list, createInfo( kindSymb, token:sub( startIndex ), columnIndex + startIndex ) )
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
  
  local searchIndex = startIndex
  while true do
    local syncIndexFlag = true
    local pattern = [==[[%-%?"%'%`%[].]==]
    local index = string.find( rawLine, pattern, searchIndex )
    if not index then
      addVal( kindSymb, rawLine:sub( startIndex ), startIndex )
      return list
    end
    local findChar = string.byte( rawLine, index )
    local nextChar = string.byte( rawLine, index + 1 )
    if findChar == 45 and nextChar ~= 45 then
      searchIndex = index + 1
      syncIndexFlag = false
    else 
      if startIndex < index then
        addVal( kindSymb, rawLine:sub( startIndex, index - 1 ), startIndex )
      end
      if findChar == 39 and nextChar == 39 then
        if string.byte( rawLine, index + 2 ) == 39 then
          local comment, nextIndex = multiComment( index + 3, "'''" )
          addVal( kindCmnt, "'''" .. comment, index )
          searchIndex = nextIndex
        else 
          addVal( kindCmnt, rawLine:sub( index ), index )
          searchIndex = #rawLine + 1
        end
      elseif findChar == 91 then
        if nextChar == 64 then
          addVal( kindDlmt, "[@", index )
          searchIndex = index + 2
        else 
          addVal( kindDlmt, "[", index )
          searchIndex = index + 1
        end
      elseif findChar == 39 or findChar == 34 then
        local workIndex = index + 1
        local pattern = '["\'\\]'
        while true do
          local endIndex = string.find( rawLine, pattern, workIndex )
          if not endIndex then
            error( string.format( "illegal string: %d: %s", index, rawLine ) )
          end
          local workChar = string.byte( rawLine, endIndex )
          if workChar == findChar then
            addVal( kindStr, rawLine:sub( index, endIndex ), index )
            searchIndex = endIndex + 1
            break
          elseif workChar == 92 then
            workIndex = workIndex + 2
          else 
            workIndex = workIndex + 1
          end
        end
      elseif findChar == 96 then
        if (nextChar == findChar and string.byte( rawLine, index + 2 ) == 96 ) then
          local str, nextIndex = multiComment( index + 3, '```' )
          addVal( kindStr, '```' .. str, index )
          searchIndex = nextIndex
        else 
          addVal( kindOpe, '`', index )
          searchIndex = index + 1
        end
      elseif findChar == 63 then
        local codeChar = rawLine:sub( index + 1, index + 1 )
        if nextChar == 92 then
          local quoted = rawLine:sub( index + 2, index + 2 )
          if quotedCharSet[quoted] then
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
        error( string.format( "illegal syntax:%s:%s", self.lineNo, rawLine:sub( index ) ) )
      end
    end
    if syncIndexFlag then
      startIndex = searchIndex
    end
  end
end

function StreamParser:getToken(  )
  if not self.lineTokenList then
    return nil
  end
  if #self.lineTokenList < self.pos then
    self.pos = 1
    self.lineTokenList = {}
    while #self.lineTokenList == 0 do
      local workList = self:parse(  )
      if  not workList then
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

local eofToken = Token.new(kindEof, "", Position.new(0, 0))
local function getEofToken(  )
  return eofToken
end
moduleObj.getEofToken = getEofToken
----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo192 = {}
  _className2InfoMap.Parser = _classInfo192
  _typeId2ClassInfoMap[ 192 ] = _classInfo192
  end
do
  local _classInfo174 = {}
  _className2InfoMap.Position = _classInfo174
  _typeId2ClassInfoMap[ 174 ] = _classInfo174
  _classInfo174.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo174.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  end
do
  local _classInfo162 = {}
  _className2InfoMap.Stream = _classInfo162
  _typeId2ClassInfoMap[ 162 ] = _classInfo162
  end
do
  local _classInfo208 = {}
  _className2InfoMap.StreamParser = _classInfo208
  _typeId2ClassInfoMap[ 208 ] = _classInfo208
  end
do
  local _classInfo178 = {}
  _className2InfoMap.Token = _classInfo178
  _typeId2ClassInfoMap[ 178 ] = _classInfo178
  _classInfo178.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo178.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 20 }
  _classInfo178.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 174 }
  end
do
  local _classInfo168 = {}
  _className2InfoMap.TxtStream = _classInfo168
  _typeId2ClassInfoMap[ 168 ] = _classInfo168
  end
do
  local _classInfo200 = {}
  _className2InfoMap.WrapParser = _classInfo200
  _typeId2ClassInfoMap[ 200 ] = _classInfo200
  end
do
  local _classInfo104 = {}
  _className2InfoMap.base = _classInfo104
  _typeId2ClassInfoMap[ 104 ] = _classInfo104
  _classInfo104.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 106 }
  end
do
  local _classInfo102 = {}
  _className2InfoMap.lune = _classInfo102
  _typeId2ClassInfoMap[ 102 ] = _classInfo102
  _classInfo102.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 104 }
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
_varName2InfoMap.kind = {
  name='kind', accessMode = 'pub', typeId = 240 }
_varName2InfoMap.kindChar = {
  name='kindChar', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindCmnt = {
  name='kindCmnt', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindDlmt = {
  name='kindDlmt', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindEof = {
  name='kindEof', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindInt = {
  name='kindInt', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindKywd = {
  name='kindKywd', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindOpe = {
  name='kindOpe', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindReal = {
  name='kindReal', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindStr = {
  name='kindStr', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindSymb = {
  name='kindSymb', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.kindType = {
  name='kindType', accessMode = 'pub', typeId = 14 }
_varName2InfoMap.noneToken = {
  name='noneToken', accessMode = 'pub', typeId = 178 }
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 7, nilable = true, orgTypeId = 6 }_typeInfoList[2] = { parentId = 1, typeId = 11, nilable = true, orgTypeId = 10 }_typeInfoList[3] = { parentId = 1, typeId = 13, nilable = true, orgTypeId = 12 }_typeInfoList[4] = { parentId = 1, typeId = 15, nilable = true, orgTypeId = 14 }_typeInfoList[5] = { parentId = 1, typeId = 21, nilable = true, orgTypeId = 20 }_typeInfoList[6] = { parentId = 1, typeId = 102, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[7] = { parentId = 1, typeId = 103, nilable = true, orgTypeId = 102 }_typeInfoList[8] = { parentId = 102, typeId = 104, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106}, }
_typeInfoList[9] = { parentId = 102, typeId = 105, nilable = true, orgTypeId = 104 }_typeInfoList[10] = { parentId = 104, typeId = 106, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {162, 168, 174, 178, 192, 200, 208, 262, 264, 266, 290}, }
_typeInfoList[11] = { parentId = 104, typeId = 107, nilable = true, orgTypeId = 106 }_typeInfoList[12] = { parentId = 106, typeId = 162, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {164, 166}, }
_typeInfoList[13] = { parentId = 106, typeId = 163, nilable = true, orgTypeId = 162 }_typeInfoList[14] = { parentId = 162, typeId = 164, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[15] = { parentId = 162, typeId = 166, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 106, typeId = 168, baseId = 162, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {170, 172}, }
_typeInfoList[17] = { parentId = 106, typeId = 169, nilable = true, orgTypeId = 168 }_typeInfoList[18] = { parentId = 168, typeId = 170, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[19] = { parentId = 168, typeId = 172, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[20] = { parentId = 106, typeId = 174, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {176}, }
_typeInfoList[21] = { parentId = 106, typeId = 175, nilable = true, orgTypeId = 174 }_typeInfoList[22] = { parentId = 174, typeId = 176, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 14}, retTypeId = {}, children = {}, }
_typeInfoList[23] = { parentId = 106, typeId = 178, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {184, 188, 190}, }
_typeInfoList[24] = { parentId = 106, typeId = 179, nilable = true, orgTypeId = 178 }_typeInfoList[25] = { parentId = 1, typeId = 182, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {178}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[26] = { parentId = 1, typeId = 183, nilable = true, orgTypeId = 182 }_typeInfoList[27] = { parentId = 178, typeId = 184, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {182}, retTypeId = {}, children = {}, }
_typeInfoList[28] = { parentId = 1, typeId = 186, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {178}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[29] = { parentId = 1, typeId = 187, nilable = true, orgTypeId = 186 }_typeInfoList[30] = { parentId = 178, typeId = 188, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {186}, children = {}, }
_typeInfoList[31] = { parentId = 178, typeId = 190, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 20, 174, 180}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 106, typeId = 192, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {194, 196, 198}, }
_typeInfoList[33] = { parentId = 106, typeId = 193, nilable = true, orgTypeId = 192 }_typeInfoList[34] = { parentId = 192, typeId = 194, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {178}, children = {}, }
_typeInfoList[35] = { parentId = 192, typeId = 196, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[36] = { parentId = 192, typeId = 198, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 106, typeId = 200, baseId = 192, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {202, 204, 206}, }
_typeInfoList[38] = { parentId = 106, typeId = 201, nilable = true, orgTypeId = 200 }_typeInfoList[39] = { parentId = 200, typeId = 202, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {178}, children = {}, }
_typeInfoList[40] = { parentId = 200, typeId = 204, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[41] = { parentId = 200, typeId = 206, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {192, 20}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 106, typeId = 208, baseId = 192, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {222, 226, 228, 286}, }
_typeInfoList[43] = { parentId = 106, typeId = 209, nilable = true, orgTypeId = 208 }_typeInfoList[44] = { parentId = 208, typeId = 222, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {162, 20, 12}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 208, typeId = 226, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[46] = { parentId = 208, typeId = 228, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {209}, children = {}, }
_typeInfoList[47] = { parentId = 1, typeId = 240, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 14}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[48] = { parentId = 1, typeId = 241, nilable = true, orgTypeId = 240 }_typeInfoList[49] = { parentId = 106, typeId = 262, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[50] = { parentId = 106, typeId = 264, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[51] = { parentId = 106, typeId = 266, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[52] = { parentId = 208, typeId = 286, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {179}, children = {}, }
_typeInfoList[53] = { parentId = 106, typeId = 290, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {178}, children = {}, }
----- meta -----
return moduleObj
