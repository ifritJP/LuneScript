--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune3 then
   _lune3 = _lune
end

local version = "1.3.0"
_moduleObj.version = version

local metaVersion = "1.0.133"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 3
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
