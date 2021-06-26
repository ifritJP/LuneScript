--lune/base/OutputDepend.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@OutputDepend'
local _lune = {}
if _lune5 then
   _lune = _lune5
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

if not _lune5 then
   _lune5 = _lune
end


local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )

local DependInfo = {}
_moduleObj.DependInfo = DependInfo
function DependInfo.new( targetModule )
   local obj = {}
   DependInfo.setmeta( obj )
   if obj.__init then obj:__init( targetModule ); end
   return obj
end
function DependInfo:__init(targetModule) 
   self.targetModule = Ast.TypeInfo.getModulePath( targetModule )
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

   stream:write( string.format( "%s.meta: \\\n", (self.targetModule:gsub( "%.", "/" ) )) )
   stream:write( string.format( "  %s.lns \\\n", (self.targetModule:gsub( "%.", "/" ) )) )
   for __index, mod in ipairs( self.importModuleList ) do
      stream:write( string.format( "  %s.meta \\\n", (mod:gsub( "%.", "/" ) )) )
   end
   
   for __index, path in ipairs( self.subModList ) do
      stream:write( string.format( "  %s.lns \\\n", (path:gsub( "%.", "/" ) )) )
   end
   
end
function DependInfo.setmeta( obj )
  setmetatable( obj, { __index = DependInfo  } )
end


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter.new( stream )
   local obj = {}
   convFilter.setmeta( obj )
   if obj.__init then obj:__init( stream ); end
   return obj
end
function convFilter:__init(stream) 
   Nodes.Filter.__init( self,false, nil, nil)
   
   self.stream = stream
end
function convFilter.setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end


function convFilter:processRoot( node, dummy )

   local moduleFull = node:get_moduleTypeInfo():getFullName( self:get_typeNameCtrl(), Ast.DummyModuleInfoManager.get_instance() )
   local dependInfo = DependInfo.new(moduleFull)
   
   for __index, impNode in ipairs( node:get_nodeManager():getImportNodeList(  ) ) do
      dependInfo:addImpotModule( impNode:get_info():get_modulePath() )
   end
   
   for __index, subfileNode in ipairs( node:get_nodeManager():getSubfileNodeList(  ) ) do
      do
         local usePath = subfileNode:get_usePath()
         if usePath ~= nil then
            dependInfo:addSubMod( usePath )
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
