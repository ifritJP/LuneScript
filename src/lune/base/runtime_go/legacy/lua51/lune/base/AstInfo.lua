--lune/base/AstInfo.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@AstInfo'
local _lune = {}
if _lune6 then
   _lune = _lune6
end
function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

if not _lune6 then
   _lune6 = _lune
end
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )
local ASTInfo = {}
_moduleObj.ASTInfo = ASTInfo
function ASTInfo.setmeta( obj )
  setmetatable( obj, { __index = ASTInfo  } )
end
function ASTInfo.new( node, exportInfo, streamName, builtinFunc )
   local obj = {}
   ASTInfo.setmeta( obj )
   if obj.__init then
      obj:__init( node, exportInfo, streamName, builtinFunc )
   end
   return obj
end
function ASTInfo:__init( node, exportInfo, streamName, builtinFunc )

   self.node = node
   self.exportInfo = exportInfo
   self.streamName = streamName
   self.builtinFunc = builtinFunc
end
function ASTInfo:get_node()
   return self.node
end
function ASTInfo:get_exportInfo()
   return self.exportInfo
end
function ASTInfo:get_streamName()
   return self.streamName
end
function ASTInfo:get_builtinFunc()
   return self.builtinFunc
end


return _moduleObj
