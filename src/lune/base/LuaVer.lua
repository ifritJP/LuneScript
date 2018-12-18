--lune/base/LuaVer.lns
local _moduleObj = {}
local __mod__ = 'lune.base.LuaVer'
if not _lune then
   _lune = {}
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
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

local LuaVerInfo = {}
_moduleObj.LuaVerInfo = LuaVerInfo
function LuaVerInfo:isSupport( symbol )

   return not self.noSupportSymMap[symbol]
end
function LuaVerInfo:getLoadCode(  )

   return LuaMod.getCode( self.loadKind )
end
function LuaVerInfo.setmeta( obj )
  setmetatable( obj, { __index = LuaVerInfo  } )
end
function LuaVerInfo.new( hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap )
   local obj = {}
   LuaVerInfo.setmeta( obj )
   if obj.__init then
      obj:__init( hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap )
   end        
   return obj 
end         
function LuaVerInfo:__init( hasBitOp, hasTableUnpack, canFormStem2Str, hasSearchPath, loadStrFuncName, canUseMetaGc, loadKind, noSupportSymMap ) 

   self.hasBitOp = hasBitOp
   self.hasTableUnpack = hasTableUnpack
   self.canFormStem2Str = canFormStem2Str
   self.hasSearchPath = hasSearchPath
   self.loadStrFuncName = loadStrFuncName
   self.canUseMetaGc = canUseMetaGc
   self.loadKind = loadKind
   self.noSupportSymMap = noSupportSymMap
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

local function symbolList2Map( list )

   local map = {}
   for __index, name in pairs( list ) do
      map[name] = true
   end
   
   return map
end

local ver51 = LuaVerInfo.new(BitOp.Cant, false, false, false, "loadstring51", false, LuaMod.CodeKind.LoadStr51, symbolList2Map( {"package.searchpath"} ))
_moduleObj.ver51 = ver51

local ver52 = LuaVerInfo.new(BitOp.HasMod, true, true, true, "loadstring52", true, LuaMod.CodeKind.LoadStr52, {})
_moduleObj.ver52 = ver52

local ver53 = LuaVerInfo.new(BitOp.HasOp, true, true, true, "loadstring52", true, LuaMod.CodeKind.LoadStr52, {})
_moduleObj.ver53 = ver53

local function getCurrentVer(  )

   local luaVer = _VERSION:gsub( "^[^%d]+", "" )
   if luaVer >= "5.3" then
      return _moduleObj.ver53
   elseif luaVer >= "5.2" then
      return _moduleObj.ver52
   end
   
   return _moduleObj.ver51
end

local curVer = getCurrentVer(  )
_moduleObj.curVer = curVer

return _moduleObj
