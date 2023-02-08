--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ver'
local _lune = {}
if _lune8 then
   _lune = _lune8
end
function _lune.loadModule( mod )
   if __luneScript and not package.preload[ mod ] then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune8 then
   _lune8 = _lune
end

local version = "1.5.3"
_moduleObj.version = version

local metaVersion = "1.0.164"
_moduleObj.metaVersion = metaVersion

local luaModVersion = 8
_moduleObj.luaModVersion = luaModVersion




return _moduleObj
