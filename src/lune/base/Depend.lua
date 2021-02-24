--lune/base/Depend.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Depend'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
end

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


local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local function getFileLastModifiedTime( path )

   local file = io.open( path )
   do
      local _exp = file
      if _exp ~= nil then
         _exp:close(  )
      else
         return nil
      end
   end
   
   
   local stream = io.popen( string.format( "stat -c '%%Y' %s", path) )
   if  nil == stream then
      local _stream = stream
   
      return nil
   end
   
   do
      local _exp = stream:read( '*a' )
      if _exp ~= nil then
         return tonumber( _exp )
      end
   end
   
   return nil
end
_moduleObj.getFileLastModifiedTime = getFileLastModifiedTime

local function searchpath51( mod, pathPattern )

   for path in string.gmatch( pathPattern, "[^;]+" ) do
      do
         local index = path:find( "%?.l[un][as]$" )
         if index ~= nil then
            local dir = path:sub( 1, index - 1 )
            local suffix = path:sub( #path - 3 )
            local target = dir .. mod:gsub( "%.", "/" ) .. suffix
            do
               local fileObj = io.open( target )
               if fileObj ~= nil then
                  fileObj:close(  )
                  return target
               end
            end
            
         end
      end
      
   end
   
   return nil
end

local function getLoadedMod(  )

   local loaded = _lune.nilacc( _lune.loadstring52( "return package.loaded" ), nil, 'call' )
   return (_lune.unwrap( loaded) )
end
_moduleObj.getLoadedMod = getLoadedMod

local function profile( validTest, func, path )

   if not validTest then
      return func(  )
   end
   
   
   local ProFi = require( 'ProFi' )
   
   ProFi:start(  )
   
   local result = func(  )
   
   ProFi:stop(  )
   ProFi:writeReport( path )
   
   return result
end
_moduleObj.profile = profile

local function getStackTrace(  )

   local txt = ""
   for level = 2, 10 do
      do
         local debugInfo = debug.getinfo( level )
         if debugInfo ~= nil then
            txt = txt .. string.format( "-- %s %s %s\n", debugInfo['short_src'], debugInfo['currentline'], debugInfo['name'])
         end
      end
      
   end
   
   return txt
end
_moduleObj.getStackTrace = getStackTrace

local lua_version = nil
local function getLuaVersion(  )

   if lua_version ~= nil then
      return lua_version
   end
   
   local loaded = _lune.nilacc( _lune.loadstring52( "return _VERSION" ), nil, 'call' )
   local version = (_lune.unwrap( loaded) )
   lua_version = version
   return version
end
_moduleObj.getLuaVersion = getLuaVersion

local function getCurrentVer(  )

   local luaVer = getLuaVersion(  ):gsub( "^[^%d]+", "" )
   if luaVer >= "5.3" then
      return LuaVer.ver53
   elseif luaVer >= "5.2" then
      return LuaVer.ver52
   end
   
   return LuaVer.ver51
end

local curVer = getCurrentVer(  )

local function setup( func )

   func( curVer:get_verKind() )
end
_moduleObj.setup = setup



local searchpathForm = searchpath51
if curVer:get_hasSearchPath() then
   searchpathForm = (_lune.unwrap( _lune.nilacc( _G['package'], nil, 'item', 'searchpath')) )
end


local function searchpath( mod, pathPattern )

   return searchpathForm( mod, pathPattern )
end
_moduleObj.searchpath = searchpath

local function existFile( path )

   local fileObj = io.open( path )
   if  nil == fileObj then
      local _fileObj = fileObj
   
      return false
   end
   
   fileObj:close(  )
   return true
end
_moduleObj.existFile = existFile

local function canUseChannel(  )

   return false
end
_moduleObj.canUseChannel = canUseChannel
local function canUseAsync(  )

   return false
end
_moduleObj.canUseAsync = canUseAsync

local function runMain( mainFunc, argList )

   
   if mainFunc ~= nil then
      return (mainFunc )( argList )
   end
   
   return -1
end
_moduleObj.runMain = runMain

return _moduleObj
