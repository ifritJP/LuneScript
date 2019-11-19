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
function Ctrl:outputResult( stream )

   stream:write( string.format( "test total: %d (OK:%d, NG:%d)\n", self.okNum + self.ngNum, self.okNum, self.ngNum) )
end
function Ctrl:err( mess, mod, lineNo )

   self.ngNum = self.ngNum + 1
   io.stderr:write( string.format( "error: %s:%d: %s\n", mod, lineNo, mess) )
end
function Ctrl:isTrue( val1, val1txt, mod, lineNo )

   if val1 == true then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "not true -- %s:[%s]\n", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Ctrl:isNotTrue( val1, val1txt, mod, lineNo )

   if not val1 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "is true -- %s:[%s]\n", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Ctrl:isNil( val1, val1txt, mod, lineNo )

   if val1 == nil then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "is not nil -- %s:[%s]\n", val1txt, tostring( val1)), mod, lineNo )
   return false
end
function Ctrl:checkEq( val1, val2, val1txt, val2txt, mod, lineNo )

   if val1 == val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "not equal -- %s:[%s] != %s:[%s]\n", val1txt, tostring( val1), val2txt, tostring( val2)), mod, lineNo )
   return false
end
function Ctrl:checkNotEq( val1, val2, val1txt, val2txt, mod, lineNo )

   if val1 ~= val2 then
      self.okNum = self.okNum + 1
      return true
   end
   
   self:err( string.format( "equal -- %s:[%s] == %s:[%s]\n", val1txt, tostring( val1), val2txt, tostring( val2)), mod, lineNo )
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


local function outputAllResult( stream )

   for __index, ctrl in pairs( ctrlList ) do
      ctrl:outputResult( stream )
   end
   
end
_moduleObj.outputAllResult = outputAllResult













return _moduleObj
