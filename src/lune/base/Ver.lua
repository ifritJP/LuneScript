--lune/base/Ver.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Ver'
if not _lune then
   _lune = {}
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

local version = "1.0.3"
_moduleObj.version = version

local metaVersion = "1.0.7"
_moduleObj.metaVersion = metaVersion

return _moduleObj
