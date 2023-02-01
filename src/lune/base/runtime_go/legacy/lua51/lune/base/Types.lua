--lune/base/Types.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Types'
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

if not _lune8 then
   _lune8 = _lune
end


local Util = _lune.loadModule( 'lune.base.Util' )


local AltBase = {}
_moduleObj.AltBase = AltBase
function AltBase._setmeta( obj )
  setmetatable( obj, { __index = AltBase  } )
end
function AltBase._new( val, mess )
   local obj = {}
   AltBase._setmeta( obj )
   if obj.__init then
      obj:__init( val, mess )
   end
   return obj
end
function AltBase:__init( val, mess )

   self.val = val
   self.mess = mess
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
Lang.Python = 3
Lang._val2NameMap[3] = 'Python'
Lang.__allList[4] = Lang.Python
Lang.C = 4
Lang._val2NameMap[4] = 'C'
Lang.__allList[5] = Lang.C

local CheckingUptodateMode = {}
CheckingUptodateMode._name2Val = {}
_moduleObj.CheckingUptodateMode = CheckingUptodateMode
function CheckingUptodateMode:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "CheckingUptodateMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function CheckingUptodateMode._from( val )
   return _lune._AlgeFrom( CheckingUptodateMode, val )
end

CheckingUptodateMode.Force1 = { "Force1", {{ func=_lune._toStr, nilable=false, child={} }}}
CheckingUptodateMode._name2Val["Force1"] = CheckingUptodateMode.Force1
CheckingUptodateMode.ForceAll = { "ForceAll"}
CheckingUptodateMode._name2Val["ForceAll"] = CheckingUptodateMode.ForceAll
CheckingUptodateMode.Normal = { "Normal"}
CheckingUptodateMode._name2Val["Normal"] = CheckingUptodateMode.Normal
CheckingUptodateMode.Touch = { "Touch"}
CheckingUptodateMode._name2Val["Touch"] = CheckingUptodateMode.Touch


local TransCtrlInfo = {}
_moduleObj.TransCtrlInfo = TransCtrlInfo
function TransCtrlInfo._new(  )
   local obj = {}
   TransCtrlInfo._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function TransCtrlInfo:__init() 
   self.useWaiter = true
   self.macroAsyncParseStmtLen = 500
   self.warningShadowing = false
   self.compatComment = false
   self.checkingDefineAbbr = true
   self.stopByWarning = false
   self.uptodateMode = _lune.newAlge( CheckingUptodateMode.Touch)
   self.validLuaval = false
   self.defaultLazy = false
   self.validCheckingMutable = true
   self.legacyMutableControl = false
   self.validAstDetailError = false
   self.validAsyncCtrl = false
   self.defaultAsync = false
   self.testing = false
   self.validMultiPhaseTransUnit = true
   self.threadPerUnitThread = false and 5 or 0
   self.validMacroAsync = false
end
function TransCtrlInfo._setmeta( obj )
  setmetatable( obj, { __index = TransCtrlInfo  } )
end


function TransCtrlInfo.create_normal(  )

   return TransCtrlInfo._new()
end


local Position = {}
setmetatable( Position, { ifList = {Mapping,__absimmut,} } )
_moduleObj.Position = Position
function Position._new( lineNo, column, streamName, orgPos )
   local obj = {}
   Position._setmeta( obj )
   if obj.__init then obj:__init( lineNo, column, streamName, orgPos ); end
   return obj
end
function Position:__init(lineNo, column, streamName, orgPos) 
   self.lineNo = lineNo
   self.column = column
   self.streamName = streamName
   self.orgPos = orgPos
end
function Position:get_orgPos(  )

   do
      local _exp = self.orgPos
      if _exp ~= nil then
         return _exp:get_orgPos()
      end
   end
   
   return self
end
function Position:get_RawOrgPos(  )

   return self.orgPos
end
function Position.create( lineNo, column, streamName, orgPos )

   return Position._new(lineNo, column, streamName, orgPos)
end
function Position:getDisplayTxt(  )

   local txt = string.format( "%s:%d:%d", self.streamName, self.lineNo, self.column)
   local orgPos = self:get_orgPos()
   if self ~= orgPos then
      local txt2 = string.format( "%s:%d:%d", orgPos.streamName, orgPos.lineNo, orgPos.column)
      return string.format( "%s: (%s)", txt2, txt)
   end
   
   return txt
end
function Position:comp( other )

   if self.streamName < other.streamName then
      return -1
   end
   
   if self.streamName > other.streamName then
      return 1
   end
   
   if self.lineNo < other.lineNo then
      return -1
   end
   
   if self.lineNo > other.lineNo then
      return 1
   end
   
   if self.column < other.column then
      return -1
   end
   
   if self.column > other.column then
      return 1
   end
   
   local orgPos, otherOrgPos = self.orgPos, other.orgPos
   if  nil == orgPos or  nil == otherOrgPos then
      local _orgPos = orgPos
      local _otherOrgPos = otherOrgPos
   
      return 0
   end
   
   return orgPos:comp( otherOrgPos )
end
function Position._setmeta( obj )
  setmetatable( obj, { __index = Position  } )
end
function Position:_toMap()
  return self
end
function Position._fromMap( val )
  local obj, mes = Position._fromMapSub( {}, val )
  if obj then
     Position._setmeta( obj )
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
   table.insert( memInfo, { name = "orgPos", func = Position._fromMap, nilable = true, child = {} } )
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
function Token._new( kind, txt, pos, consecutive, commentList )
   local obj = {}
   Token._setmeta( obj )
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
   
   
   Util.err( string.format( "illegal delimit -- %s", self.txt) )
end
function Token:getLineCount(  )

   local count = 1
   do
      for _1 in self.txt:gmatch( "\n" ) do
         count = count + 1
      end
      
   end
   
   return count
end
function Token._setmeta( obj )
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
     Token._setmeta( obj )
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


local nonePos = Position._new(0, -1, "eof")
_moduleObj.nonePos = nonePos

local noneToken = Token._new(TokenKind.Eof, "", _moduleObj.nonePos, false, {})
_moduleObj.noneToken = noneToken


local StdinFile = {}
_moduleObj.StdinFile = StdinFile
function StdinFile._setmeta( obj )
  setmetatable( obj, { __index = StdinFile  } )
end
function StdinFile._new( mod, txt )
   local obj = {}
   StdinFile._setmeta( obj )
   if obj.__init then
      obj:__init( mod, txt )
   end
   return obj
end
function StdinFile:__init( mod, txt )

   self.mod = mod
   self.txt = txt
end
function StdinFile:get_mod()
   return self.mod
end
function StdinFile:get_txt()
   return self.txt
end


local ParserSrc = {}
ParserSrc._name2Val = {}
_moduleObj.ParserSrc = ParserSrc
function ParserSrc:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "ParserSrc.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function ParserSrc._from( val )
   return _lune._AlgeFrom( ParserSrc, val )
end

ParserSrc.LnsCode = { "LnsCode", {{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toInt, nilable=true, child={} }}}
ParserSrc._name2Val["LnsCode"] = ParserSrc.LnsCode
ParserSrc.LnsPath = { "LnsPath", {{ func=_lune._toStr, nilable=true, child={} },{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toInt, nilable=true, child={} }}}
ParserSrc._name2Val["LnsPath"] = ParserSrc.LnsPath
ParserSrc.Parser = { "Parser", {{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toBool, nilable=false, child={} },{ func=_lune._toStr, nilable=false, child={} },{ func=_lune._toInt, nilable=true, child={} }}}
ParserSrc._name2Val["Parser"] = ParserSrc.Parser


local defaultParserPipeSize = 100
_moduleObj.defaultParserPipeSize = defaultParserPipeSize


return _moduleObj
