local _moduleObj = {}
local __mod__ = '@test.@Sub9'
print( __mod__ )
local function func( val )
   local __func__ = '@test.@Sub9.func'

   print( __func__, val )
end
_moduleObj.func = func

return _moduleObj
