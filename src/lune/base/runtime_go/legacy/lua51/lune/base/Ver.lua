--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune6 then
   _lune6 = _lune
end

local version = "1.3.0"
_moduleObj.version = version

local metaVersion = "1.0.151"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 6
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
