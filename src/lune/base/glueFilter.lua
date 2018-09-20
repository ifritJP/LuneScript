--lune/base/glueFilter.lns
local _moduleObj = {}
if not _ENV._lune then
   _lune = {}
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

local Ast = require( 'lune.base.Ast' )

local Parser = require( 'lune.base.Parser' )

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
do
  end

function glueGenerator:write( txt )
  self.srcStream:write( txt )
end

function glueGenerator:writeHeader( txt )
  self.headerStream:write( txt )
end

function glueGenerator:outputMethod( classFullName, methodNode )
  local declInfo = methodNode:get_declInfo()
  
  do
    local _exp = declInfo:get_name()
    if _exp ~= nil then
    
        if declInfo:get_staticFlag() then
          self:write( string.format( "static method %s.%s", classFullName, _exp.txt) )
        else 
          self:write( string.format( "method %s.%s", classFullName, _exp.txt) )
        end
      end
  end
  
end

function glueGenerator:getArgInfo( argNode )
  local argType = argNode:get_expType()
  
  local orgType = (argType:get_nilable() and argType:get_orgTypeInfo() or argType ):get_srcTypeInfo()
  
  local typeTxt = ""
  
  if orgType == Ast.builtinTypeInt then
    typeTxt = "int"
  elseif orgType == Ast.builtinTypeReal then
    typeTxt = "double"
  elseif orgType == Ast.builtinTypeString then
    typeTxt = "const char *"
  end
  local argName = ""
  
  if argNode:get_kind() == Ast.nodeKindDeclArg then
    argName = (argNode ):get_name().txt
  end
  return typeTxt, orgType, argName
end

function glueGenerator:outputPrototype( node )
  self:write( string.format( "static int lns_glue_%s( lua_State * pLua )", node:get_expType():getTxt(  )) )
end

function glueGenerator:outputUserPrototype( node, gluePrefix )
  local expType = node:get_expType()
  
  self:writeHeader( string.format( "extern int %s%s( lua_State * pLua", gluePrefix, expType:getTxt(  )) )
  local declInfo = node:get_declInfo()
  
  local addFlag = true
  
  for __index, argNode in pairs( declInfo:get_argList() ) do
    local typeTxt, argType, argName = self:getArgInfo( argNode )
    
    if typeTxt ~= "" then
      if addFlag then
        self:writeHeader( ", " )
      end
      addFlag = true
      self:writeHeader( string.format( "%s %s", typeTxt, argName) )
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
extern void * lns_glue_get_%s( lua_State * pLua, int index );
extern void * lns_glue_new_%s( lua_State * pLua, size_t size );
]==], moduleSymbolFull, moduleSymbolFull) )
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
  local callTxtList = {}
  
  for index, argNode in pairs( declInfo:get_argList() ) do
    local typeTxt, argType, argName = self:getArgInfo( argNode )
    
    local addVal = declInfo:get_staticFlag() and 0 or 1
    
    if typeTxt ~= "" then
      do
        local _switchExp = argType
        if _switchExp == Ast.builtinTypeInt then
          self:write( string.format( "  %s %s = luaL_checkinteger( pLua, %d );\n", typeTxt, argName, index + addVal) )
          table.insert( callTxtList, argName )
        elseif _switchExp == Ast.builtinTypeReal then
          self:write( string.format( "  %s %s = luaL_checknumber( pLua, %d );\n", typeTxt, argName, index + addVal) )
          table.insert( callTxtList, argName )
        elseif _switchExp == Ast.builtinTypeString then
          self:write( string.format( "  size_t size_%s = 0;\n", argName) .. string.format( "  const char * %s = luaL_checklstring( pLua, %d, &size_%s );\n", argName, index + addVal, argName) )
          table.insert( callTxtList, argName )
          table.insert( callTxtList, "size_" .. argName )
        end
      end
      
    end
  end
  self:write( string.format( "  return %s( pLua", name) )
  for index, argTxt in pairs( callTxtList ) do
    self:write( ", " )
    self:write( argTxt )
  end
  self:write( ");\n" )
  self:write( "}\n" )
end

function glueGenerator:outputClass( moduleFullName, node, gluePrefix )
  local moduleSymbolFull = moduleFullName:gsub( "%.", "_" )
  
  local staticMethodNodeList = {}
  
  local methodNodeList = {}
  
  for __index, fieldNode in pairs( node:get_fieldList() ) do
    if fieldNode:get_kind() == Ast.nodeKindDeclMethod then
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
setmetatable( glueFilter, { __index = Filter } )
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

self.outputDir = outputDir
  end
do
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
                
                local glue = glueGenerator.new(createFile( moduleSymbolName .. "_glue.c" ), createFile( moduleSymbolName .. "_glue.h" ))
                
                do
                  local _exp = node:get_gluePrefix()
                  if _exp ~= nil then
                  
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
