--lune/base/OutputDepend.lns
local _moduleObj = {}
local __mod__ = 'lune.base.OutputDepend'
if not _lune then
   _lune = {}
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

local Ast = _lune.loadModule( 'lune.base.Ast' )
local Util = _lune.loadModule( 'lune.base.Util' )
local TransUnit = _lune.loadModule( 'lune.base.TransUnit' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local convFilter = {}
setmetatable( convFilter, { __index = Ast.Filter } )
function convFilter.new( stream )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( stream ); end
   return obj
end
function convFilter:__init(stream) 
   Ast.Filter.__init( self )
   
   self.stream = stream
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

function convFilter:processRoot( node, parent )

   local moduleFull = node:get_moduleTypeInfo():getFullName( {} )
   self.stream:write( string.format( "%s.meta: \\\n", moduleFull:gsub( "%.", "/" )) )
   self.stream:write( string.format( "  %s.lns \\\n", moduleFull:gsub( "%.", "/" )) )
   do
      local importList = node:get_nodeManager():getImportNodeList(  )
      if importList ~= nil then
         for __index, impNode in pairs( importList ) do
            self.stream:write( string.format( "  %s.meta \\\n", impNode:get_modulePath():gsub( "%.", "/" )) )
         end
         
      end
   end
   
   do
      local subfileList = node:get_nodeManager():getSubfileNodeList(  )
      if subfileList ~= nil then
         for __index, subfileNode in pairs( subfileList ) do
            do
               local usePath = subfileNode:get_usePath()
               if usePath ~= nil then
                  self.stream:write( string.format( "  %s.lns \\\n", usePath:gsub( "%.", "/" )) )
               end
            end
            
         end
         
      end
   end
   
end


local function createFilter( stream )

   return convFilter.new(stream)
end
_moduleObj.createFilter = createFilter
return _moduleObj
