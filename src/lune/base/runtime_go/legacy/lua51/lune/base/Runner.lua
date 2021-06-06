--lune/base/Runner.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Runner'
local _lune = {}
if _lune4 then
   _lune = _lune4
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
    runner:run()
    return false
end

if not _lune4 then
   _lune4 = _lune
end


local RunModeWhenFull = {}
_moduleObj.RunModeWhenFull = RunModeWhenFull
RunModeWhenFull._val2NameMap = {}
function RunModeWhenFull:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "RunModeWhenFull.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function RunModeWhenFull._from( val )
   if RunModeWhenFull._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
RunModeWhenFull.__allList = {}
function RunModeWhenFull.get__allList()
   return RunModeWhenFull.__allList
end

RunModeWhenFull.Sync = 0
RunModeWhenFull._val2NameMap[0] = 'Sync'
RunModeWhenFull.__allList[1] = RunModeWhenFull.Sync
RunModeWhenFull.Queue = 1
RunModeWhenFull._val2NameMap[1] = 'Queue'
RunModeWhenFull.__allList[2] = RunModeWhenFull.Queue
RunModeWhenFull.Skip = 2
RunModeWhenFull._val2NameMap[2] = 'Skip'
RunModeWhenFull.__allList[3] = RunModeWhenFull.Skip


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
function Runner:start( mode )

   _lune._run(self, mode )
end
function Runner:join(  )

   
end
function Runner.setmeta( obj )
  setmetatable( obj, { __index = Runner  } )
end


return _moduleObj
