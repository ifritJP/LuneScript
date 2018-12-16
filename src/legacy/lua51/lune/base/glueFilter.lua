--lune/base/glueFilter.lns
local _moduleObj = {}
local __mod__ = 'lune.base.glueFilter'
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

local Ast = _lune.loadModule( 'lune.base.Ast' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local glueGenerator = {}
function glueGenerator.setmeta( obj )
  setmetatable( obj, { __index = glueGenerator  } )
end
function glueGenerator.new( srcStream, headerStream )
   local obj = {}
   glueGenerator.setmeta( obj )
   if obj.__init then
      obj:__init( srcStream, headerStream )
   end        
   return obj 
end         
function glueGenerator:__init( srcStream, headerStream ) 

   self.srcStream = srcStream
   self.headerStream = headerStream
end

function glueGenerator:write( txt )

   self.srcStream:write( txt )
end

function glueGenerator:writeHeader( txt )

   self.headerStream:write( txt )
end

function glueGenerator:getArgInfo( argNode )

   local argType = argNode:get_expType()
   local orgType = (argType:get_nilable() and argType:get_orgTypeInfo() or argType ):get_srcTypeInfo()
   local typeTxt = ""
   local nilableTypeTxt = ""
   if orgType == Ast.builtinTypeInt then
      typeTxt = "int"
      nilableTypeTxt = typeTxt .. " *"
   elseif orgType == Ast.builtinTypeReal then
      typeTxt = "double"
      nilableTypeTxt = typeTxt .. " *"
   elseif orgType == Ast.builtinTypeString then
      typeTxt = "const char *"
      nilableTypeTxt = typeTxt
   end
   
   local argName = ""
   if argNode:get_kind() == Ast.NodeKind.get_DeclArg() then
      argName = (argNode ):get_name().txt
   end
   
   return typeTxt, argType:get_nilable() and nilableTypeTxt or typeTxt, orgType, argName
end

function glueGenerator:outputPrototype( node )

   self:write( string.format( "static int lns_glue_%s( lua_State * pLua )", node:get_expType():getTxt(  )) )
end

function glueGenerator:outputUserPrototype( node, gluePrefix )

   local expType = node:get_expType()
   self:writeHeader( string.format( "extern int %s%s( lua_State * pLua", gluePrefix, expType:getTxt(  )) )
   local declInfo = node:get_declInfo()
   for __index, argNode in pairs( declInfo:get_argList() ) do
      local typeTxt, argTypeTxt, argType, argName = self:getArgInfo( argNode )
      if typeTxt ~= "" then
         self:writeHeader( string.format( ", %s %s", argTypeTxt, argName) )
         if argType == Ast.builtinTypeString then
            self:writeHeader( string.format( ", int size_%s", argName) )
         end
         
      end
      
   end
   
   self:writeHeader( " )" )
end

function glueGenerator:outputPrototypeList( methodNodeList )

   for __index, node in pairs( methodNodeList ) do
      self:outputPrototype( node )
      self:write( ";\n" )
   end
   
end

function glueGenerator:outputUserPrototypeList( methodNodeList, gluePrefix )

   for __index, node in pairs( methodNodeList ) do
      self:outputUserPrototype( node, gluePrefix )
      self:writeHeader( ";\n" )
   end
   
end

function glueGenerator:outputFuncReg( symbolName, methodNodeList )

   self:write( string.format( "static const luaL_Reg %s[] = {\n", symbolName) )
   for __index, node in pairs( methodNodeList ) do
      do
         local nameToken = node:get_declInfo():get_name()
         if nameToken ~= nil then
            local name = nameToken.txt
            self:write( string.format( '  { "%s", lns_glue_%s },\n', name, name) )
         end
      end
      
   end
   
   self:write( '  { NULL, NULL }\n};\n' )
end

function glueGenerator:outputCommonFunc( moduleSymbolFull )

   self:writeHeader( string.format( [==[
extern int luaopen_%s( lua_State * pLua );
extern void * lns_glue_get_%s( lua_State * pLua, int index );
extern void * lns_glue_new_%s( lua_State * pLua, size_t size );
]==], moduleSymbolFull, moduleSymbolFull, moduleSymbolFull) )
   self:write( string.format( [==[
void * lns_glue_get_%s( lua_State * pLua, int index )
{
    return luaL_checkudata( pLua, index, s_full_class_name);
}

static void lns_glue_setupObjMethod(
    lua_State * pLua, const char * pName, const luaL_Reg * pReg )
{
    luaL_newmetatable(pLua, pName );
    lua_pushvalue(pLua, -1);
    lua_setfield(pLua, -2, "__index");

#if LUA_VERSION_NUM >= 502
    luaL_setfuncs(pLua, pReg, 0);

    lua_pop(pLua, 1);
#else
    luaL_register(pLua, NULL, pReg );

    lua_pop(pLua, 1);
#endif
}

void * lns_glue_new_%s( lua_State * pLua, size_t size )
{
    void * pBuf = lua_newuserdata( pLua, size );
    if ( pBuf == NULL ) {
        return NULL;
    }
    
#if LUA_VERSION_NUM >= 502
    luaL_setmetatable( pLua, s_full_class_name );
#else
    luaL_getmetatable( pLua, s_full_class_name );
    lua_setmetatable( pLua, -2 );
#endif

    return pBuf;
}

int luaopen_%s( lua_State * pLua )
{
    lns_glue_setupObjMethod( pLua, s_full_class_name, s_lua_method_info );

#if LUA_VERSION_NUM >= 502
    luaL_newlib( pLua, s_lua_func_info );
#else
    luaL_register( pLua, s_full_class_name, s_lua_func_info );
#endif
    return 1;
}
]==], moduleSymbolFull, moduleSymbolFull, moduleSymbolFull) )
end

local GlueArgInfo = {}
function GlueArgInfo.setmeta( obj )
  setmetatable( obj, { __index = GlueArgInfo  } )
end
function GlueArgInfo.new( index, argName, callArgName, callTxt, setTxt, typeInfo )
   local obj = {}
   GlueArgInfo.setmeta( obj )
   if obj.__init then
      obj:__init( index, argName, callArgName, callTxt, setTxt, typeInfo )
   end        
   return obj 
end         
function GlueArgInfo:__init( index, argName, callArgName, callTxt, setTxt, typeInfo ) 

   self.index = index
   self.argName = argName
   self.callArgName = callArgName
   self.callTxt = callTxt
   self.setTxt = setTxt
   self.typeInfo = typeInfo
end
function GlueArgInfo:get_index()       
   return self.index         
end
function GlueArgInfo:get_argName()       
   return self.argName         
end
function GlueArgInfo:get_callArgName()       
   return self.callArgName         
end
function GlueArgInfo:get_callTxt()       
   return self.callTxt         
end
function GlueArgInfo:get_setTxt()       
   return self.setTxt         
end
function GlueArgInfo:get_typeInfo()       
   return self.typeInfo         
end

function glueGenerator:outputMethod( node, gluePrefix )

   local declInfo = node:get_declInfo()
   local name = ""
   do
      local _exp = declInfo:get_name()
      if _exp ~= nil then
         name = gluePrefix .. _exp.txt
      else
         return 
      end
   end
   
   self:outputPrototype( node )
   self:write( "{\n" )
   local glueArgInfoList = {}
   for index, argNode in pairs( declInfo:get_argList() ) do
      local typeTxt, argTypeTxt, argType, argName = self:getArgInfo( argNode )
      if typeTxt ~= "" then
         local addVal = declInfo:get_staticFlag() and 0 or 1
         local callArgName = argName
         self:write( string.format( "  %s %s = ", typeTxt, argName) )
         do
            local _switchExp = argType
            if _switchExp == Ast.builtinTypeInt then
               self:write( "0;\n" )
            elseif _switchExp == Ast.builtinTypeReal then
               self:write( "0.0;\n" )
            elseif _switchExp == Ast.builtinTypeString then
               self:write( "NULL;\n" )
            end
         end
         
         if argNode:get_expType():get_nilable() then
            if argType ~= Ast.builtinTypeString then
               callArgName = string.format( "_p%s", argName)
               self:write( string.format( "  %s %s = NULL;\n", argTypeTxt, callArgName) )
            end
            
         end
         
         local setTxt = ""
         local callTxt = ""
         do
            local _switchExp = argType
            if _switchExp == Ast.builtinTypeInt then
               setTxt = string.format( "  %s = luaL_checkinteger( pLua, %d );\n", argName, index + addVal)
               callTxt = callArgName
            elseif _switchExp == Ast.builtinTypeReal then
               setTxt = string.format( "  %s = luaL_checknumber( pLua, %d );\n", argName, index + addVal)
               callTxt = callArgName
            elseif _switchExp == Ast.builtinTypeString then
               self:write( string.format( "  size_t size_%s = 0;\n", argName) )
               setTxt = string.format( "  %s = luaL_checklstring( pLua, %d, &size_%s );\n", argName, index + addVal, argName)
               callTxt = string.format( "%s, size_%s", argName, argName)
            end
         end
         
         table.insert( glueArgInfoList, GlueArgInfo.new(index + addVal, argName, callArgName, callTxt, setTxt, argNode:get_expType()) )
      end
      
   end
   
   for __index, glueArgInfo in pairs( glueArgInfoList ) do
      if glueArgInfo:get_typeInfo():get_nilable() then
         self:write( string.format( '  if ( !lua_isnoneornil( pLua, %d ) ) {\n', glueArgInfo:get_index()) )
         self:write( '  ' )
      end
      
      self:write( glueArgInfo:get_setTxt() )
      if glueArgInfo:get_typeInfo():get_nilable() then
         if glueArgInfo:get_callArgName() ~= glueArgInfo:get_argName() then
            self:write( string.format( '    %s = &%s;\n', glueArgInfo:get_callArgName(), glueArgInfo:get_argName()) )
         end
         
         self:write( '  }\n' )
      end
      
   end
   
   self:write( string.format( "  return %s( pLua", name) )
   for __index, glueArgInfo in pairs( glueArgInfoList ) do
      self:write( ", " )
      self:write( glueArgInfo:get_callTxt() )
   end
   
   self:write( ");\n" )
   self:write( "}\n" )
end

function glueGenerator:outputClass( moduleFullName, node, gluePrefix )

   local moduleSymbolFull = moduleFullName:gsub( "%.", "_" )
   local staticMethodNodeList = {}
   local methodNodeList = {}
   for __index, fieldNode in pairs( node:get_fieldList() ) do
      if fieldNode:get_kind() == Ast.NodeKind.get_DeclMethod() then
         local methodNode = fieldNode
         if methodNode:get_declInfo():get_staticFlag() then
            table.insert( staticMethodNodeList, methodNode )
         else
          
            table.insert( methodNodeList, methodNode )
         end
         
      end
      
   end
   
   self:writeHeader( '#include <lauxlib.h>\n' )
   self:outputUserPrototypeList( staticMethodNodeList, gluePrefix )
   self:outputUserPrototypeList( methodNodeList, gluePrefix )
   self:write( string.format( '#include "%s_glue.h"\n', moduleSymbolFull) )
   self:outputPrototypeList( staticMethodNodeList )
   self:outputPrototypeList( methodNodeList )
   self:write( string.format( 'static const char * s_full_class_name = "%s";\n', moduleFullName) )
   self:outputFuncReg( "s_lua_func_info", staticMethodNodeList )
   self:outputFuncReg( "s_lua_method_info", methodNodeList )
   local classSymbolFull = moduleSymbolFull .. "_" .. node:get_name().txt
   self:outputCommonFunc( moduleSymbolFull )
   for __index, methodNode in pairs( methodNodeList ) do
      self:outputMethod( methodNode, gluePrefix )
   end
   
   for __index, methodNode in pairs( staticMethodNodeList ) do
      self:outputMethod( methodNode, gluePrefix )
   end
   
end

local glueFilter = {}
setmetatable( glueFilter, { __index = Ast.Filter } )
_moduleObj.glueFilter = glueFilter
function glueFilter.setmeta( obj )
  setmetatable( obj, { __index = glueFilter  } )
end
function glueFilter.new( outputDir )
   local obj = {}
   glueFilter.setmeta( obj )
   if obj.__init then
      obj:__init( outputDir )
   end        
   return obj 
end         
function glueFilter:__init( outputDir ) 

   Ast.Filter.__init( self )
   self.outputDir = outputDir
end

function glueFilter:processRoot( node )

   local function createFile( filename )
   
      local filePath = string.format( "%s/%s", _lune.unwrapDefault( self.outputDir, "."), filename)
      do
         local _exp = io.open( filePath, "w" )
         if _exp ~= nil then
            return _exp
         end
      end
      
      error( string.format( "open error -- %s ", filePath) )
   end
   
   do
      local nodeList = node:get_nodeManager():getDeclClassNodeList(  )
      if nodeList ~= nil then
         for __index, node in pairs( nodeList ) do
            do
               local moduleName = node:get_moduleName()
               if moduleName ~= nil then
                  local moduleSymbolName = moduleName:getExcludedDelimitTxt(  ):gsub( "%.", "_" )
                  do
                     local _exp = node:get_gluePrefix()
                     if _exp ~= nil then
                        local glue = glueGenerator.new(createFile( moduleSymbolName .. "_glue.c" ), createFile( moduleSymbolName .. "_glue.h" ))
                        glue:outputClass( moduleSymbolName, node, _exp )
                     end
                  end
                  
               end
            end
            
         end
         
      end
   end
   
end

return _moduleObj
