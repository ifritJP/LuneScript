--lune/base/Depend.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Depend'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.loadstring51( txt, env )
   local func = loadstring( txt )
   if func and env then
      setfenv( func, env )
   end
   return func
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
   local val
   
   do
      do
         local stream = io.popen( string.format( "stat -c '%%Y' %s", path) )
         if stream ~= nil then
            val = stream:read( '*a' )
            stream:close(  )
         else
            val = nil
         end
      end
   end
   do
      local _exp = val
      if _exp ~= nil then
         return tonumber( _exp )
      end
   end
   return nil
end
_moduleObj.getFileLastModifiedTime = getFileLastModifiedTime


local function searchpath51( mod, pathPattern )

   local target = ""
   do
      for path in string.gmatch( pathPattern, "[^;]+" ) do
         do
            local index = path:find( "%?.l[un][as]$" )
            if index ~= nil then
               local dir = path:sub( 1, index - 1 )
               local suffix = path:sub( #path - 3 )
               target = dir .. mod:gsub( "%.", "/" ) .. suffix
               do
                  local fileObj = io.open( target )
                  if fileObj ~= nil then
                     fileObj:close(  )
                     break
                  end
               end
            end
         end
      end
   end
   if target ~= "" then
      return target
   end
   return nil
end


local function getLoadedMod(  )

   local ret
   
   do
      local loaded = _lune.unwrap( _lune.nilacc( _lune.loadstring51( "return package.loaded" ), nil, 'call' ))
      ret = loaded
   end
   return ret
end
_moduleObj.getLoadedMod = getLoadedMod


local function profile( validTest, func, path )

   if not validTest then
      return func(  )
   end
   local ProFi = require( 'ProFi' )ProFi:start(  )
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
            txt = txt .. string.format( "-- %s %s %s\n", tostring( debugInfo['short_src']), tostring( debugInfo['currentline']), tostring( debugInfo['name']))
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
   local version
   
   do
      local loaded = _lune.nilacc( _lune.loadstring51( "return _VERSION" ), nil, 'call' )
      version = (_lune.unwrap( loaded) )
   end
   lua_version = version:gsub( "^[^%d]+", "" )
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

   local ret
   
   do
      if mainFunc ~= nil then
         ret = (mainFunc )( argList )
      else
         ret = 1
      end
   end
   return ret
end
_moduleObj.runMain = runMain


local function setupShebang(  )

   do
      local loaded, mess = _lune.loadstring51( [==[
if not _lune then
  _lune = {}
end
 _lune._shebang = true
]==] )
      if loaded ~= nil then
         do
            local mod = loaded(  )
            if mod ~= nil then
               (mod )(  )
            end
         end
      else
         print( mess )
      end
   end
end
_moduleObj.setupShebang = setupShebang


local function getGOPATH(  )

   local OS = require( "os" )do
      local path = OS.getenv( "GOPATH" )
      if path ~= nil then
         return path
      end
   end
   do
      local home = OS.getenv( "HOME" )
      if home ~= nil then
         return home .. "/go"
      end
   end
   return nil
end
_moduleObj.getGOPATH = getGOPATH


local function getBindLns( mod )

   return nil
end
_moduleObj.getBindLns = getBindLns


local function setRuntimeLog( valid )

end
_moduleObj.setRuntimeLog = setRuntimeLog


local function setRuntimeThreadLimit( limit )

end
_moduleObj.setRuntimeThreadLimit = setRuntimeThreadLimit

local function setRunnerLog( limit )

end
_moduleObj.setRunnerLog = setRunnerLog

local function dumpRunnerLog( stream )

end
_moduleObj.dumpRunnerLog = dumpRunnerLog


return _moduleObj
