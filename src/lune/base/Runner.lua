--lune/base/Runner.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Runner'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if ifType == class then
               return true
            end
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

function _lune._run( runner, mod )
    if mod == 2 then
      return false
    end
    runner:run()
    return true
end

if not _lune6 then
   _lune6 = _lune
end


local Runner = {}
setmetatable( Runner, { ifList = {__Runner,} } )
_moduleObj.Runner = Runner
function Runner.new(  )
   local obj = {}
   Runner.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Runner:__init() 
end
function Runner:run(  )

   self:runMain(  )
end
function Runner:start( mode, name )

   _lune._run(self, mode, name )
end
function Runner.setmeta( obj )
  setmetatable( obj, { __index = Runner  } )
end


return _moduleObj
