--lune/base/DependLuaOnLns.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@DependLuaOnLns'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.loadstring52( txt, env )
   if not env then
      return load( txt )
   end
   return load( txt, "", "bt", env )
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

if not _lune6 then
   _lune6 = _lune
end


local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )



local function runLuaOnLns( luaCode, baseDir, async )

   local newEnv = {}
   for key, val in pairs( _G ) do
      newEnv[key] = val
   end
   
   newEnv["_lnsLoad"] = function ( name, txt )
   
      local importModuleInfo = frontInterface.ImportModuleInfo.new()
      local val = frontInterface.loadFromLnsTxt( importModuleInfo, baseDir, name, txt )
      return val
   end
   
   local loaded, err
   
   if async then
      do
         loaded, err = _lune.loadstring52( luaCode, newEnv )
      end
      
   else
    
      do
         loaded, err = _lune.loadstring52( luaCode, newEnv )
      end
      
   end
   
   
   if loaded ~= nil then
      return loaded, ""
   end
   
   if err ~= nil then
      return nil, err
   end
   
   return nil, ""
end
_moduleObj.runLuaOnLns = runLuaOnLns

function __luneGetLocal( varName )

   local index = 1
   while true do
      local name, val = debug.getlocal( 3, index )
      if name == varName then
         return val
      end
      
      if not name then
         break
      end
      
      
      index = index + 1
   end
   
   error( "not found -- " .. varName )
end
_moduleObj.__luneGetLocal = __luneGetLocal

function __luneSym2Str( val )

   do
      local _exp = val
      if _exp ~= nil then
         if type( _exp ) ~= "table" then
            return string.format( "%s", _exp )
         end
         
         
         local txt = ""
         for __index, item in ipairs( _exp ) do
            txt = txt .. item
         end
         
         return txt
      end
   end
   
   return nil
end
_moduleObj.__luneSym2Str = __luneSym2Str

return _moduleObj
