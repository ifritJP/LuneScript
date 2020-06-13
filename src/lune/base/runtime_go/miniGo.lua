--lune/base/runtime_go/miniGo.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@runtime_go.@miniGo'
local _lune = {}
if _lune2 then
   _lune = _lune2
end
if not _lune2 then
   _lune2 = _lune
end
local function sub(  )

   print( "hello world", "abc", "xyz", 1, 1.5, 10 + 100 )
end
local function sub2(  )

   return 1
end
local function sub3(  )

   return 100, "abc"
end
sub(  )
print( sub2(  ), sub3(  ) )

return _moduleObj
