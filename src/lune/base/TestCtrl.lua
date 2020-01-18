--lune/base/TestCtrl.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@TestCtrl'
local _lune = {}
if _lune2 then
   _lune = _lune2
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

if not _lune2 then
   _lune2 = _lune
end

local Meta = _lune.loadModule( 'lune.base.Meta' )



local Ctrl = {}
_moduleObj.Ctrl = Ctrl
function Ctrl:runTestcase( testcase )

   testcase(  )
end
function Ctrl:run( modObj )

   if modObj ~= nil then
      do
         local testMap = modObj['__testMap']
         if testMap ~= nil then
            for name, testcase in pairs( testMap ) do
               print( string.format( "%s:", name) )
               self:runTestcase( testcase )
            end
            
         end
      end
      
   end
   
end
function Ctrl.setmeta( obj )
  setmetatable( obj, { __index = Ctrl  } )
end
function Ctrl.new(  )
   local obj = {}
   Ctrl.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function Ctrl:__init(  )

end


return _moduleObj
