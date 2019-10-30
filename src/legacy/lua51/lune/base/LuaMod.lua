--lune/base/LuaMod.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@LuaMod'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not _lune1 then
   _lune1 = _lune
end

local CodeKind = {}
_moduleObj.CodeKind = CodeKind
CodeKind._val2NameMap = {}
function CodeKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CodeKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CodeKind._from( val )
   if CodeKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CodeKind.__allList = {}
function CodeKind.get__allList()
   return CodeKind.__allList
end

CodeKind.Init = 0
CodeKind._val2NameMap[0] = 'Init'
CodeKind.__allList[1] = CodeKind.Init
CodeKind.Unpack = 1
CodeKind._val2NameMap[1] = 'Unpack'
CodeKind.__allList[2] = CodeKind.Unpack
CodeKind.NilAcc = 2
CodeKind._val2NameMap[2] = 'NilAcc'
CodeKind.__allList[3] = CodeKind.NilAcc
CodeKind.Unwrap = 3
CodeKind._val2NameMap[3] = 'Unwrap'
CodeKind.__allList[4] = CodeKind.Unwrap
CodeKind.LoadModule = 4
CodeKind._val2NameMap[4] = 'LoadModule'
CodeKind.__allList[5] = CodeKind.LoadModule
CodeKind.LoadStr51 = 5
CodeKind._val2NameMap[5] = 'LoadStr51'
CodeKind.__allList[6] = CodeKind.LoadStr51
CodeKind.LoadStr52 = 6
CodeKind._val2NameMap[6] = 'LoadStr52'
CodeKind.__allList[7] = CodeKind.LoadStr52
CodeKind.Mapping = 7
CodeKind._val2NameMap[7] = 'Mapping'
CodeKind.__allList[8] = CodeKind.Mapping
CodeKind.Alge = 8
CodeKind._val2NameMap[8] = 'Alge'
CodeKind.__allList[9] = CodeKind.Alge
CodeKind.AlgeMapping = 9
CodeKind._val2NameMap[9] = 'AlgeMapping'
CodeKind.__allList[10] = CodeKind.AlgeMapping
CodeKind.SetMapping = 10
CodeKind._val2NameMap[10] = 'SetMapping'
CodeKind.__allList[11] = CodeKind.SetMapping
CodeKind.SetOp = 11
CodeKind._val2NameMap[11] = 'SetOp'
CodeKind.__allList[12] = CodeKind.SetOp
CodeKind.InstanceOf = 12
CodeKind._val2NameMap[12] = 'InstanceOf'
CodeKind.__allList[13] = CodeKind.InstanceOf
CodeKind.Cast = 13
CodeKind._val2NameMap[13] = 'Cast'
CodeKind.__allList[14] = CodeKind.Cast
CodeKind.Finalize = 14
CodeKind._val2NameMap[14] = 'Finalize'
CodeKind.__allList[15] = CodeKind.Finalize


local codeMap = {}

codeMap[CodeKind.Init] = [==[
local _lune = {}
]==]

codeMap[CodeKind.Unpack] = [==[
if not table.unpack then
   table.unpack = unpack
end
]==]

codeMap[CodeKind.NilAcc] = [==[
function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
end
]==]

codeMap[CodeKind.Unwrap] = [==[
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
]==]

codeMap[CodeKind.LoadModule] = [==[
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end
]==]

codeMap[CodeKind.LoadStr51] = [==[
function _lune.loadstring51( txt, env )
   local func = loadstring( txt )
   if func and env then
      setfenv( func, env )
   end
   return func
end
]==]

codeMap[CodeKind.LoadStr52] = [==[
function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end
]==]

codeMap[CodeKind.Mapping] = [==[
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
]==]

codeMap[CodeKind.SetOp] = [==[
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
]==]

codeMap[CodeKind.SetMapping] = [==[
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
]==]

codeMap[CodeKind.AlgeMapping] = [==[
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
]==]

codeMap[CodeKind.Alge] = [==[
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end
]==]

codeMap[CodeKind.InstanceOf] = [==[
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
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end
]==]

local CastKind = {}
_moduleObj.CastKind = CastKind
CastKind._val2NameMap = {}
function CastKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CastKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CastKind._from( val )
   if CastKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CastKind.__allList = {}
function CastKind.get__allList()
   return CastKind.__allList
end

CastKind.Int = 0
CastKind._val2NameMap[0] = 'Int'
CastKind.__allList[1] = CastKind.Int
CastKind.Real = 1
CastKind._val2NameMap[1] = 'Real'
CastKind.__allList[2] = CastKind.Real
CastKind.Str = 2
CastKind._val2NameMap[2] = 'Str'
CastKind.__allList[3] = CastKind.Str
CastKind.Class = 3
CastKind._val2NameMap[3] = 'Class'
CastKind.__allList[4] = CastKind.Class


codeMap[CodeKind.Cast] = string.format( [==[
function _lune.__Cast( obj, kind, class )
   if kind == %d then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == %d then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == %d then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == %d then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end
]==], CastKind.Int, CastKind.Real, CastKind.Str, CastKind.Class)

codeMap[CodeKind.Finalize] = [==[
return _lune
]==]

local function getCode( kind )

   return _lune.unwrap( codeMap[kind])
end
_moduleObj.getCode = getCode
return _moduleObj
