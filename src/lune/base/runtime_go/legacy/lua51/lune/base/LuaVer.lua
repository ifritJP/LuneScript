--lune/base/LuaVer.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@LuaVer'
local _lune = {}
if _lune6 then
   _lune = _lune6
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

if not _lune6 then
   _lune6 = _lune
end


local LuaMod = _lune.loadModule( 'lune.base.LuaMod' )

local BitOp = {}
_moduleObj.BitOp = BitOp
BitOp._val2NameMap = {}
function BitOp:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BitOp.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function BitOp._from( val )
   if BitOp._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
BitOp.__allList = {}
function BitOp.get__allList()
   return BitOp.__allList
end

BitOp.HasOp = 0
BitOp._val2NameMap[0] = 'HasOp'
BitOp.__allList[1] = BitOp.HasOp
BitOp.HasMod = 1
BitOp._val2NameMap[1] = 'HasMod'
BitOp.__allList[2] = BitOp.HasMod
BitOp.Cant = 2
BitOp._val2NameMap[2] = 'Cant'
BitOp.__allList[3] = BitOp.Cant


local VerKind = {}
_moduleObj.VerKind = VerKind
VerKind._val2NameMap = {}
function VerKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "VerKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function VerKind._from( val )
   if VerKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
VerKind.__allList = {}
function VerKind.get__allList()
   return VerKind.__allList
end

VerKind.v51 = 51
VerKind._val2NameMap[51] = 'v51'
VerKind.__allList[1] = VerKind.v51
VerKind.v52 = 52
VerKind._val2NameMap[52] = 'v52'
VerKind.__allList[2] = VerKind.v52
VerKind.v53 = 53
VerKind._val2NameMap[53] = 'v53'
VerKind.__allList[3] = VerKind.v53


local LuaVerInfo = {}
_moduleObj.LuaVerInfo = LuaVerInfo
function LuaVerInfo:isSupport( symbol )

   return not _lune._Set_has(self.noSupportSymMap, symbol )
end
function LuaVerInfo:getLoadCode(  )

   return LuaMod.getCode( self.loadKind )
end
function LuaVerInfo.setmeta( obj )
  setmetatable( obj, { __index = LuaVerInfo  } )
end
function LuaVerInfo.new( verKind, hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap )
   local obj = {}
   LuaVerInfo.setmeta( obj )
   if obj.__init then
      obj:__init( verKind, hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap )
   end
   return obj
end
function LuaVerInfo:__init( verKind, hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap )

   self.verKind = verKind
   self.hasBitOp = hasBitOp
   self.hasTableUnpack = hasTableUnpack
   self.canFormStem2Str = canFormStem2Str
   self.hasSearchPath = hasSearchPath
   self.loadStrFuncName = loadStrFuncName
   self.canUseMetaGc = canUseMetaGc
   self.loadKind = loadKind
   self.noSupportSymMap = noSupportSymMap
end
function LuaVerInfo:get_verKind()
   return self.verKind
end
function LuaVerInfo:get_hasBitOp()
   return self.hasBitOp
end
function LuaVerInfo:get_hasTableUnpack()
   return self.hasTableUnpack
end
function LuaVerInfo:get_canFormStem2Str()
   return self.canFormStem2Str
end
function LuaVerInfo:get_hasSearchPath()
   return self.hasSearchPath
end
function LuaVerInfo:get_loadStrFuncName()
   return self.loadStrFuncName
end
function LuaVerInfo:get_canUseMetaGc()
   return self.canUseMetaGc
end


local ver51 = LuaVerInfo.new(VerKind.v51, BitOp.Cant, false, false, false, "loadstring51", false, LuaMod.CodeKind.LoadStr51, {["package.searchpath"] = true})
_moduleObj.ver51 = ver51

local ver52 = LuaVerInfo.new(VerKind.v52, BitOp.HasMod, true, true, true, "loadstring52", true, LuaMod.CodeKind.LoadStr52, {})
_moduleObj.ver52 = ver52

local ver53 = LuaVerInfo.new(VerKind.v53, BitOp.HasOp, true, true, true, "loadstring52", true, LuaMod.CodeKind.LoadStr52, {})
_moduleObj.ver53 = ver53

local kind2verMap = {[VerKind.v51] = _moduleObj.ver51, [VerKind.v52] = _moduleObj.ver52, [VerKind.v53] = _moduleObj.ver53}

local curVer = nil
local function setCurVer( ver )

   local verKind = VerKind._from( ver )
   if  nil == verKind then
      local _verKind = verKind
   
      return 
   end
   
   curVer = _lune.unwrap( kind2verMap[verKind])
end
_moduleObj.setCurVer = setCurVer
local function getCurVer(  )

   return _lune.unwrap( curVer)
end
_moduleObj.getCurVer = getCurVer

return _moduleObj
