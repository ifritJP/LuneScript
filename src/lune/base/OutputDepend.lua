--lune/base/OutputDepend.lns
local _moduleObj = {}
local __mod__ = 'lune.base.OutputDepend'
if not _lune then
   _lune = {}
end
local Ast = require( 'lune.base.Ast' )
local Util = require( 'lune.base.Util' )
local TransUnit = require( 'lune.base.TransUnit' )
local frontInterface = require( 'lune.base.frontInterface' )
local convFilter = {}
setmetatable( convFilter, { __index = Ast.Filter } )
function convFilter.new( stream )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( stream ); end
   return obj
end
function convFilter:__init(stream) 
   self.stream = stream
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

function convFilter:processRoot( node, parent )

   local moduleFull = node:get_moduleTypeInfo():getFullName( {} )
   self.stream:write( string.format( "%s.lua: \\\n", moduleFull:gsub( "%.", "/" )) )
   do
      local importList = node:get_nodeManager():getImportNodeList(  )
      if importList ~= nil then
         for __index, impNode in pairs( importList ) do
            self.stream:write( string.format( "  %s.lns \\\n", impNode:get_modulePath():gsub( "%.", "/" )) )
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
