--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune2 then
   _lune2 = _lune
end
local version = "1.1.10"
_moduleObj.version = version

local metaVersion = "1.0.82"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 2
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
