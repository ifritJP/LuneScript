--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Ver'
if not _lune then
   _lune = {}
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

local version = "1.0.2"
_moduleObj.version = version

local metaVersion = "1.0.0"
_moduleObj.metaVersion = metaVersion

return _moduleObj
