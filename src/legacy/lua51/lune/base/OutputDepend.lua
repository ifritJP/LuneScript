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
local DependInfo = {}
_moduleObj.DependInfo = DependInfo
function DependInfo.new( targetModule )
   local obj = {}
   DependInfo.setmeta( obj )
   if obj.__init then obj:__init( targetModule ); end
   return obj
end
function DependInfo:__init(targetModule) 
   self.targetModule = targetModule
   self.importModuleList = {}
   self.subModList = {}
end
function DependInfo:addImpotModule( mod )

   table.insert( self.importModuleList, mod )
end
function DependInfo:addSubMod( path )

   table.insert( self.subModList, path )
end
function DependInfo:output( stream )

   stream:write( string.format( "%s.meta: \\\n", self.targetModule:gsub( "%.", "/" )) )
   stream:write( string.format( "  %s.lns \\\n", self.targetModule:gsub( "%.", "/" )) )
   for __index, mod in pairs( self.importModuleList ) do
      stream:write( string.format( "  %s.meta \\\n", mod:gsub( "%.", "/" )) )
   end
   
   for __index, path in pairs( self.subModList ) do
      stream:write( string.format( "  %s.lns \\\n", path:gsub( "%.", "/" )) )
   end
   
end
function DependInfo.setmeta( obj )
  setmetatable( obj, { __index = DependInfo  } )
end

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
   local dependInfo = DependInfo.new(moduleFull)
   do
      local importList = node:get_nodeManager():getImportNodeList(  )
      if importList ~= nil then
         for __index, impNode in pairs( importList ) do
            dependInfo:addImpotModule( impNode:get_modulePath() )
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
                  dependInfo:addSubMod( usePath )
               end
            end
            
         end
         
      end
   end
   
   dependInfo:output( self.stream )
end


local function createFilter( stream )

   return convFilter.new(stream)
end
_moduleObj.createFilter = createFilter
return _moduleObj
