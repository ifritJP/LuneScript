--lune/base/Testing.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Testing'
local _lune = {}
if _lune1 then
   _lune = _lune1
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

if not _lune1 then
   _lune1 = _lune
end


local ctrlList = {}

local function getCtrlList(  )

   return ctrlList
end
_moduleObj.getCtrlList = getCtrlList

local Ctrl = {}
_moduleObj.Ctrl = Ctrl
function Ctrl.new(  )
   local obj = {}
   Ctrl.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Ctrl:__init() 
   self.okNum = 0
   self.ngNum = 0
   table.insert( ctrlList, self )
end
function Ctrl:checkEq( val1, val2, val1txt, val2txt, mod, lineNo )

   if val1 == val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self.ngNum = self.ngNum + 1
   io.stderr:write( string.format( "error:%s:%d: not equal -- %s:[%s] != %s:[%s]\n", mod, lineNo, val1txt, val1, val2txt, val2) )
   return false
end
function Ctrl:checkNotEq( val1, val2, val1txt, val2txt, mod, lineNo )

   if val1 ~= val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self.ngNum = self.ngNum + 1
   io.stderr:write( string.format( "error:%s:%d: equal -- %s:[%s] == %s:[%s]\n", mod, lineNo, val1txt, val1, val2txt, val2) )
   return false
end
function Ctrl.setmeta( obj )
  setmetatable( obj, { __index = Ctrl  } )
end
function Ctrl:get_okNum()
   return self.okNum
end
function Ctrl:get_ngNum()
   return self.ngNum
end






return _moduleObj
