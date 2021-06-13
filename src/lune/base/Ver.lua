--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune4 then
   _lune = _lune4
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune4 then
   _lune4 = _lune
end

local version = "1.3.0"
_moduleObj.version = version

local metaVersion = "1.0.144"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 4
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
