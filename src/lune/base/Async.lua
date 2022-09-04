--lune/base/Async.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Async'
local _lune = {}
if _lune7 then
   _lune = _lune7
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

if not _lune7 then
   _lune7 = _lune
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


return _moduleObj
