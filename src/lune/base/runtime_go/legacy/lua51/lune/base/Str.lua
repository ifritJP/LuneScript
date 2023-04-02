--lune/base/Str.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Str'
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

if not _lune8 then
   _lune8 = _lune
end



local function startsWith( txt, ptn )

   return txt:find( ptn, 1, true ) ~= nil
end
_moduleObj.startsWith = startsWith
local function endsWith( txt, ptn )

   return txt:find( ptn, #txt - #ptn + 1, true ) ~= nil
end
_moduleObj.endsWith = endsWith
local function getLineList( txt )

   local list = {}
   
   do
      
      for line in string.gmatch( txt, "[^\n]*\n" ) do
         table.insert( list, line )
      end
      
      for last in string.gmatch( txt, "[^\n]+$" ) do
         table.insert( list, last )
      end
      
   end
   
   return list
end
_moduleObj.getLineList = getLineList
local function replace( txt, src, dst )

   return (txt:gsub( src, dst ) )
end
_moduleObj.replace = replace

local Builder = {}
_moduleObj.Builder = Builder
function Builder:get_txt(  )

   if #self.progress == 0 then
      return self.txt
   end
   
   return self.txt .. self.progress
end
function Builder._new(  )
   local obj = {}
   Builder._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Builder:__init() 
   self.txt = ""
   self.progress = ""
end
function Builder:add( val )

   if #self.progress + #val > 1000 then
      self.txt = string.format( "%s%s%s", self.txt, self.progress, val)
      self.progress = ""
   else
    
      self.progress = self.progress .. val
   end
   
end
function Builder:len(  )

   return #self.txt + #self.progress
end
function Builder:clear(  )

   self.txt = ""
   self.progress = ""
end
function Builder:flush(  )

   self.txt = self.txt .. self.progress
   self.progress = ""
end
function Builder._setmeta( obj )
  setmetatable( obj, { __index = Builder  } )
end


local function isValidStrBuilder(  )

   return false
end
_moduleObj.isValidStrBuilder = isValidStrBuilder

return _moduleObj
