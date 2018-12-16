--lune/base/Depend.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Depend'
if not _lune then
   _lune = {}
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune.nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end 
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local function getFileLastModifiedTime( path )

   local file = io.open( path )
   do
      local _exp = file
      if _exp ~= nil then
         _exp:close(  )
      else
         return nil
      end
   end
   
   local stream = io.popen( string.format( "stat -c '%%Y' %s", path) )
   do
      local _exp = _lune.nilacc( stream, 'read', 'callmtd' , '*a' )
      if _exp ~= nil then
         return tonumber( _exp )
      end
   end
   
   return nil
end
_moduleObj.getFileLastModifiedTime = getFileLastModifiedTime
local function searchpath51( mod, pathPattern )

   for path in string.gmatch( pathPattern, "[^;]+" ) do
      do
         local index = path:find( "%?.l[un][as]$" )
         if index ~= nil then
            local dir = path:sub( 1, index - 1 )
            local suffix = path:sub( #path - 3 )
            local target = dir .. mod:gsub( "%.", "/" ) .. suffix
            do
               local fileObj = io.open( target )
               if fileObj ~= nil then
                  fileObj:close(  )
                  return target
               end
            end
            
         end
      end
      
   end
   
   return nil
end


local searchpathForm = searchpath51
if LuaVer.curVer:get_hasSearchPath() then
   searchpathForm = (_lune.unwrap( _lune.nilacc( _G['package'], nil, 'item', 'searchpath')) )
end

local function searchpath( mod, pathPattern )

   return searchpathForm( mod, pathPattern )
end
_moduleObj.searchpath = searchpath
return _moduleObj
