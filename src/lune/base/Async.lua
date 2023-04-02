--lune/base/Async.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Async'
local _lune = {}
if _lune8 then
   _lune = _lune8
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

if not _lune8 then
   _lune8 = _lune
end




local PipeItem = {}
_moduleObj.PipeItem = PipeItem
function PipeItem._setmeta( obj )
  setmetatable( obj, { __index = PipeItem  } )
end
function PipeItem._new( item )
   local obj = {}
   PipeItem._setmeta( obj )
   if obj.__init then
      obj:__init( item )
   end
   return obj
end
function PipeItem:__init( item )

   self.item = item
end
function PipeItem:get_item()
   return self.item
end

local Pipe = {}
_moduleObj.Pipe = Pipe
function Pipe._new( pipe )
   local obj = {}
   Pipe._setmeta( obj )
   if obj.__init then obj:__init( pipe ); end
   return obj
end
function Pipe:__init(pipe) 
   self.pipe = pipe
   self.started = false
   self.setuped = false
end
function Pipe:getNext(  )

   if self.started then
      do
         local pipe = self.pipe
         if pipe ~= nil then
            local val = pipe:get(  )
            if  nil == val then
               local _val = val
            
               return nil
            end
            
            return PipeItem._new(val)
         end
      end
      
   end
   
   return self:access(  )
end
function Pipe:run(  )

   self:setup(  )
   self.setuped = true
   local pipe = self.pipe
   if  nil == pipe then
      local _pipe = pipe
   
      return 
   end
   
   while true do
      local val = self:access(  )
      if  nil == val then
         local _val = val
      
         pipe:put( nil )
         break
      end
      
      pipe:put( val:get_item() )
   end
   
end
function Pipe:start(  )

   self.started = true
end
function Pipe:stop(  )

   self.started = false
   if not self.setuped then
      self:setup(  )
   end
   
end
function Pipe._setmeta( obj )
  setmetatable( obj, { __index = Pipe  } )
end

local RunnerBase = {}
setmetatable( RunnerBase, { ifList = {__Runner,__AsyncItem,} } )
_moduleObj.RunnerBase = RunnerBase
function RunnerBase._new( pipe )
   local obj = {}
   RunnerBase._setmeta( obj )
   if obj.__init then obj:__init( pipe ); end
   return obj
end
function RunnerBase:__init(pipe) 
   self.pipe = pipe
   self.artifact = nil
   self.ranFlag = false
end
function RunnerBase:run(  )

   self.artifact = self:runMain(  )
   self.ranFlag = true
   
   do
      local pipe = self.pipe
      if pipe ~= nil then
         pipe:put( self )
      end
   end
   
end
function RunnerBase._setmeta( obj )
  setmetatable( obj, { __index = RunnerBase  } )
end
function RunnerBase:get_artifact()
   return self.artifact
end
function RunnerBase:get_ranFlag()
   return self.ranFlag
end




local Waiter = {}
_moduleObj.Waiter = Waiter
function Waiter._new( pipeItemCount )
   local obj = {}
   Waiter._setmeta( obj )
   if obj.__init then obj:__init( pipeItemCount ); end
   return obj
end
function Waiter:__init(pipeItemCount) 
   self.pipe = nil
   self.asyncNum = 0
   self.finRunnerList = {}
end
function Waiter:startRunner( runner, mode, name )

   self.asyncNum = self.asyncNum + 1
   local result = _lune._run(runner, mode, name )
   if not self.pipe or not result then
      self.asyncNum = self.asyncNum - 1
      table.insert( self.finRunnerList, runner )
   end
   
   return result
end
function Waiter:wait( func )

   for __index, runner in ipairs( self.finRunnerList ) do
      func( runner )
   end
   
   do
      local pipe = self.pipe
      if pipe ~= nil then
         for _1 = 1, self.asyncNum do
            local runner = pipe:get(  )
            if  nil == runner then
               local _runner = runner
            
               break
            end
            
            func( runner )
         end
         
      end
   end
   
end
function Waiter._setmeta( obj )
  setmetatable( obj, { __index = Waiter  } )
end
function Waiter:get_pipe()
   return self.pipe
end


return _moduleObj
