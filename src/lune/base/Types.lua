--lune/base/Types.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Types'
local _lune = {}
if _lune3 then
   _lune = _lune3
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
local Lang = {}
_moduleObj.Lang = Lang
Lang._val2NameMap = {}
function Lang:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "Lang.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function Lang._from( val )
   if Lang._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
Lang.__allList = {}
function Lang.get__allList()
   return Lang.__allList
end

Lang.Same = 0
Lang._val2NameMap[0] = 'Same'
Lang.__allList[1] = Lang.Same
Lang.Lua = 1
Lang._val2NameMap[1] = 'Lua'
Lang.__allList[2] = Lang.Lua
Lang.Go = 2
Lang._val2NameMap[2] = 'Go'
Lang.__allList[3] = Lang.Go
Lang.C = 3
Lang._val2NameMap[3] = 'C'
Lang.__allList[4] = Lang.C

local CheckingUptodateMode = {}
_moduleObj.CheckingUptodateMode = CheckingUptodateMode
CheckingUptodateMode._val2NameMap = {}
function CheckingUptodateMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CheckingUptodateMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CheckingUptodateMode._from( val )
   if CheckingUptodateMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CheckingUptodateMode.__allList = {}
function CheckingUptodateMode.get__allList()
   return CheckingUptodateMode.__allList
end

CheckingUptodateMode.Force = 'force'
CheckingUptodateMode._val2NameMap['force'] = 'Force'
CheckingUptodateMode.__allList[1] = CheckingUptodateMode.Force
CheckingUptodateMode.Normal = 'none'
CheckingUptodateMode._val2NameMap['none'] = 'Normal'
CheckingUptodateMode.__allList[2] = CheckingUptodateMode.Normal
CheckingUptodateMode.Touch = 'touch'
CheckingUptodateMode._val2NameMap['touch'] = 'Touch'
CheckingUptodateMode.__allList[3] = CheckingUptodateMode.Touch


local TransCtrlInfo = {}
_moduleObj.TransCtrlInfo = TransCtrlInfo
function TransCtrlInfo.setmeta( obj )
  setmetatable( obj, { __index = TransCtrlInfo  } )
end
function TransCtrlInfo.new( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval, defaultLazy )
   local obj = {}
   TransCtrlInfo.setmeta( obj )
   if obj.__init then
      obj:__init( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval, defaultLazy )
   end
   return obj
end
function TransCtrlInfo:__init( warningShadowing, compatComment, checkingDefineAbbr, stopByWarning, uptodateMode, validLuaval, defaultLazy )

   self.warningShadowing = warningShadowing
   self.compatComment = compatComment
   self.checkingDefineAbbr = checkingDefineAbbr
   self.stopByWarning = stopByWarning
   self.uptodateMode = uptodateMode
   self.validLuaval = validLuaval
   self.defaultLazy = defaultLazy
end


function TransCtrlInfo.create_normal(  )

   return TransCtrlInfo.new(false, false, true, false, CheckingUptodateMode.Touch, false, false)
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
TokenKind.Sheb = 10
TokenKind._val2NameMap[10] = 'Sheb'
TokenKind.__allList[11] = TokenKind.Sheb
TokenKind.Eof = 11
TokenKind._val2NameMap[11] = 'Eof'
TokenKind.__allList[12] = TokenKind.Eof


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
   for _91 in self.txt:gmatch( "\n" ) do
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


local noneToken = Token.new(TokenKind.Eof, "", Position.new(0, -1, "eof"), false, {})
_moduleObj.noneToken = noneToken


return _moduleObj
