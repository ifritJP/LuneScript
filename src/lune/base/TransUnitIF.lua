--lune/base/TransUnitIF.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@TransUnitIF'
local _lune = {}
if _lune3 then
   _lune = _lune3
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
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

if not _lune3 then
   _lune3 = _lune
end
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )

local Types = _lune.loadModule( 'lune.base.Types' )

local DeclClassMode = {}
_moduleObj.DeclClassMode = DeclClassMode
DeclClassMode._val2NameMap = {}
function DeclClassMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclClassMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function DeclClassMode._from( val )
   if DeclClassMode._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
DeclClassMode.__allList = {}
function DeclClassMode.get__allList()
   return DeclClassMode.__allList
end

DeclClassMode.Class = 0
DeclClassMode._val2NameMap[0] = 'Class'
DeclClassMode.__allList[1] = DeclClassMode.Class
DeclClassMode.Interface = 1
DeclClassMode._val2NameMap[1] = 'Interface'
DeclClassMode.__allList[2] = DeclClassMode.Interface
DeclClassMode.Module = 2
DeclClassMode._val2NameMap[2] = 'Module'
DeclClassMode.__allList[3] = DeclClassMode.Module
DeclClassMode.LazyModule = 3
DeclClassMode._val2NameMap[3] = 'LazyModule'
DeclClassMode.__allList[4] = DeclClassMode.LazyModule


local Modifier = {}
_moduleObj.Modifier = Modifier
function Modifier:createModifier( typeInfo, mutMode )

   if not self.validMutControl then
      return typeInfo
   end
   
   return self.processInfo:createModifier( typeInfo, mutMode )
end
function Modifier.setmeta( obj )
  setmetatable( obj, { __index = Modifier  } )
end
function Modifier.new( validMutControl, processInfo )
   local obj = {}
   Modifier.setmeta( obj )
   if obj.__init then
      obj:__init( validMutControl, processInfo )
   end
   return obj
end
function Modifier:__init( validMutControl, processInfo )

   self.validMutControl = validMutControl
   self.processInfo = processInfo
end
function Modifier:set_validMutControl( validMutControl )
   self.validMutControl = validMutControl
end


local NSInfo = {}
_moduleObj.NSInfo = NSInfo
function NSInfo:isLockedAsync(  )

   return #self.lockedAsyncStack > 0
end
function NSInfo.new( typeInfo, pos )
   local obj = {}
   NSInfo.setmeta( obj )
   if obj.__init then obj:__init( typeInfo, pos ); end
   return obj
end
function NSInfo:__init(typeInfo, pos) 
   self.nobody = false
   self.lockedAsyncStack = {}
   self.loopScopeQueue = {}
   
   self.typeInfo = typeInfo
   self.pos = pos
end
function NSInfo:incLock(  )

   table.insert( self.lockedAsyncStack, #self.loopScopeQueue )
end
function NSInfo:decLock(  )

   table.remove( self.lockedAsyncStack )
end
function NSInfo:canBreak(  )

   local len = #self.lockedAsyncStack
   local loopQueueLen = #self.loopScopeQueue
   if len == 0 then
      return loopQueueLen > 0
   end
   
   return self.lockedAsyncStack[len] < loopQueueLen
end
function NSInfo.setmeta( obj )
  setmetatable( obj, { __index = NSInfo  } )
end
function NSInfo:get_nobody()
   return self.nobody
end
function NSInfo:set_nobody( nobody )
   self.nobody = nobody
end
function NSInfo:get_typeInfo()
   return self.typeInfo
end
function NSInfo:get_pos()
   return self.pos
end
function NSInfo:get_loopScopeQueue()
   return self.loopScopeQueue
end


local TransUnitIF = {}
_moduleObj.TransUnitIF = TransUnitIF
function TransUnitIF.setmeta( obj )
  setmetatable( obj, { __index = TransUnitIF  } )
end
function TransUnitIF.new(  )
   local obj = {}
   TransUnitIF.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function TransUnitIF:__init(  )

end


return _moduleObj
