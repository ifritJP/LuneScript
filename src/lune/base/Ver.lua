--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune7 then
   _lune = _lune7
end
function _lune.loadModule( mod )
   if __luneScript and not package.preload[ mod ] then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune7 then
   _lune7 = _lune
end

local version = "1.5.0"
_moduleObj.version = version

local metaVersion = "1.0.158"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 7
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
