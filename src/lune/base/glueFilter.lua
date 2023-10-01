--lune/base/glueFilter.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@glueFilter'
local _lune = {}
if _lune8 then
   _lune = _lune8
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
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end



local Ast = _lune.loadModule( 'lune.base.Ast' )

local Nodes = _lune.loadModule( 'lune.base.Nodes' )

local Util = _lune.loadModule( 'lune.base.Util' )


local glueGenerator = {}
function glueGenerator._setmeta( obj )
  setmetatable( obj, { __index = glueGenerator  } )
end
function glueGenerator._new( srcStream, headerStream )
   local obj = {}
   glueGenerator._setmeta( obj )
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
   local orgType = (argType:get_nilable() and argType:get_nonnilableType() or argType ):get_srcTypeInfo()
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
   do
      local _exp = _lune.__Cast( argNode, 3, Nodes.DeclArgNode )
      if _exp ~= nil then
         argName = _exp:get_name().txt
      end
   end
   return typeTxt, argType:get_nilable() and nilableTypeTxt or typeTxt, orgType, argName
end


local function getDeclFuncInfo( node )

   do
      local work = _lune.__Cast( node, 3, Nodes.DeclConstrNode )
      if work ~= nil then
         return work:get_declInfo()
      end
   end
   do
      local work = _lune.__Cast( node, 3, Nodes.DeclDestrNode )
      if work ~= nil then
         return work:get_declInfo()
      end
   end
   do
      local work = _lune.__Cast( node, 3, Nodes.DeclMethodNode )
      if work ~= nil then
         return work:get_declInfo()
      end
   end
   do
      local work = _lune.__Cast( node, 3, Nodes.DeclFuncNode )
      if work ~= nil then
         return work:get_declInfo()
      end
   end
   Util.err( "failed to get DeclFuncInfo" )
end


local function getFuncName( name )

   if name == "__free" then
      return "__gc"
   end
   return name
end


function glueGenerator:outputPrototype( node )

   local name = getFuncName( node:get_expType():get_rawTxt() )
   self:write( string.format( "static int lns_glue_%s( lua_State * pLua )", name) )
end


function glueGenerator:outputUserPrototype( node, gluePrefix )

   local expType = node:get_expType()
   self:writeHeader( string.format( "extern int %s%s( lua_State * pLua", gluePrefix, getFuncName( expType:get_rawTxt() )) )
   local declInfo = getDeclFuncInfo( node )
   for __index, argNode in ipairs( declInfo:get_argList() ) do
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

   for __index, node in ipairs( methodNodeList ) do
      self:outputPrototype( node )
      self:write( ";\n" )
   end
end


function glueGenerator:outputUserPrototypeList( methodNodeList, gluePrefix )

   for __index, node in ipairs( methodNodeList ) do
      self:outputUserPrototype( node, gluePrefix )
      self:writeHeader( ";\n" )
   end
end


function glueGenerator:outputFuncReg( symbolName, methodNodeList )

   self:write( string.format( "static const luaL_Reg %s[] = {\n", symbolName) )
   for __index, node in ipairs( methodNodeList ) do
      do
         local nameToken = getDeclFuncInfo( node ):get_name()
         if nameToken ~= nil then
            local name = getFuncName( nameToken.txt )
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
function GlueArgInfo._setmeta( obj )
  setmetatable( obj, { __index = GlueArgInfo  } )
end
function GlueArgInfo._new( index, argName, callArgName, callTxt, setTxt, typeInfo )
   local obj = {}
   GlueArgInfo._setmeta( obj )
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

   local declInfo = getDeclFuncInfo( node )
   local name
   
   do
      local _exp = declInfo:get_name()
      if _exp ~= nil then
         name = gluePrefix .. getFuncName( _exp.txt )
      else
         return 
      end
   end
   self:outputPrototype( node )
   self:write( "{\n" )
   local glueArgInfoList = {}
   for index, argNode in ipairs( declInfo:get_argList() ) do
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
         table.insert( glueArgInfoList, GlueArgInfo._new(index + addVal, argName, callArgName, callTxt, setTxt, argNode:get_expType()) )
      end
   end
   for __index, glueArgInfo in ipairs( glueArgInfoList ) do
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
   for __index, glueArgInfo in ipairs( glueArgInfoList ) do
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
   for __index, fieldNode in ipairs( node:get_fieldList() ) do
      do
         local methodNode = _lune.__Cast( fieldNode, 3, Nodes.DeclMethodNode )
         if methodNode ~= nil then
            if methodNode:get_declInfo():get_staticFlag() then
               table.insert( staticMethodNodeList, methodNode )
            else
             
               table.insert( methodNodeList, methodNode )
            end
         else
            do
               local methodNode = _lune.__Cast( fieldNode, 3, Nodes.DeclDestrNode )
               if methodNode ~= nil then
                  table.insert( methodNodeList, methodNode )
               end
            end
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
   self:outputCommonFunc( moduleSymbolFull )
   for __index, methodNode in ipairs( methodNodeList ) do
      self:outputMethod( methodNode, gluePrefix )
   end
   for __index, methodNode in ipairs( staticMethodNodeList ) do
      self:outputMethod( methodNode, gluePrefix )
   end
end


local glueFilter = {}
setmetatable( glueFilter, { __index = Nodes.Filter } )
function glueFilter._new( outputDir )
   local obj = {}
   glueFilter._setmeta( obj )
   if obj.__init then obj:__init( outputDir ); end
   return obj
end
function glueFilter:__init(outputDir) 
   Nodes.Filter.__init( self,false, nil, nil)
   self.outputDir = outputDir
end
function glueFilter._setmeta( obj )
  setmetatable( obj, { __index = glueFilter  } )
end


local function createFilter( outputDir )

   return glueFilter._new(outputDir)
end
_moduleObj.createFilter = createFilter


function glueFilter:processRoot( node, dummy )

   local function createFile( filename )
   
      local filePath = string.format( "%s/%s", _lune.unwrapDefault( self.outputDir, "."), filename)
      do
         local _exp = io.open( filePath, "w" )
         if _exp ~= nil then
            return _exp
         end
      end
      Util.err( string.format( "open error -- %s ", filePath) )
   end
   for __index, declClassNode in ipairs( node:get_nodeManager():getDeclClassNodeList(  ) ) do
      do
         local moduleName = declClassNode:get_moduleName()
         if moduleName ~= nil then
            local moduleSymbolName = moduleName:getExcludedDelimitTxt(  ):gsub( "%.", "_" )
            do
               local _exp = declClassNode:get_gluePrefix()
               if _exp ~= nil then
                  local cFile = createFile( moduleSymbolName .. "_glue.c" )
                  local hFile = createFile( moduleSymbolName .. "_glue.h" )
                  local glue = glueGenerator._new(cFile, hFile)
                  glue:outputClass( moduleSymbolName, declClassNode, _exp )
                  cFile:close(  )
                  hFile:close(  )
               end
            end
         end
      end
   end
end


return _moduleObj
