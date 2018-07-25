--lune/base/convLua.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local Util = require( 'lune.base.Util' )

local PubVerInfo = {}
function PubVerInfo.new( staticFlag, accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = PubVerInfo } )
  if obj.__init then
    obj:__init( staticFlag, accessMode, typeInfo )
  end        
  return obj 
 end         
function PubVerInfo:__init( staticFlag, accessMode, typeInfo ) 
            
self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.typeInfo = typeInfo
  end

-- none

local PubFuncInfo = {}
function PubFuncInfo.new( accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = PubFuncInfo } )
  if obj.__init then
    obj:__init( accessMode, typeInfo )
  end        
  return obj 
 end         
function PubFuncInfo:__init( accessMode, typeInfo ) 
            
self.accessMode = accessMode
  self.typeInfo = typeInfo
  end

local convFilter = {}
setmetatable( convFilter, { __index = Filter } )
moduleObj.convFilter = convFilter
function convFilter.new( streamName, stream, metaStream, convMode, inMacro )
  local obj = {}
  setmetatable( obj, { __index = convFilter } )
  if obj.__init then obj:__init( streamName, stream, metaStream, convMode, inMacro ); end
return obj
end
function convFilter:__init(streamName, stream, metaStream, convMode, inMacro) 
  self.macroDepth = 0
  self.streamName = streamName
  self.stream = stream
  self.metaStream = metaStream
  self.outMetaFlag = false
  self.moduleName2Info = {}
  self.convMode = convMode
  self.inMacro = inMacro
  self.indent = 0
  self.curLineNo = 1
  self.classId2TypeInfo = {}
  self.classId2MemberList = {}
  self.pubVarName2InfoMap = {}
  self.pubFuncName2InfoMap = {}
  self.needIndent = false
end

local function filter( node, filter, parent, baseIndent )
  node:processFilter( filter, parent, baseIndent )
end

local stepIndent = 2

local builtInModuleSet = {}

builtInModuleSet["io"] = true
builtInModuleSet["string"] = true
builtInModuleSet["table"] = true
builtInModuleSet["math"] = true
builtInModuleSet["debug"] = true
builtInModuleSet["_luneScript"] = true
function convFilter:write( txt )
  local stream = self.stream
  
  if self.outMetaFlag then
    stream = self.metaStream
  end
  if self.needIndent then
    stream:write( string.rep( " ", self.indent ) )
    self.needIndent = false
  end
  for cr in string.gmatch( txt, "\n" ) do
    self.curLineNo = self.curLineNo + 1
  end
  stream:write( txt )
end

function convFilter:setIndent( indent )
  self.indent = indent
end

function convFilter:writeln( txt, baseIndent )
  self:write( txt )
  self:write( "\n" )
  self.needIndent = true
  self.indent = baseIndent
end

function convFilter:processNone( node, parent, baseIndent )
  self:writeln( "-- none", baseIndent )
end

-- none

function convFilter:processImport( node, parent, baseIndent )
  local module = node:get_modulePath(  )
  
  local moduleName = string.gsub( module, ".*%.", "" )
  
  local moduleInfo = true
  
  self.moduleName2Info[moduleName] = moduleInfo
  if self.convMode == "exe" or self.convMode == "ast" then
    self:writeln( string.format( "local %s = _luneScript.loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

-- none

function convFilter:outputMeta( node, baseIndent )
  do
    local _switchExp = self.convMode
    if _switchExp == "lua" or _switchExp == "save" then
      return 
    end
  end
  
  self.outMetaFlag = true
  if self.stream ~= self.metaStream then
    self:writeln( "local moduleObj = {}", baseIndent )
  end
  self:writeln( "----- meta -----", baseIndent )
  local typeId2TypeInfo = {}
  
  local typeId2UseFlag = {}
  
  local pickupClassMap = {}
  
  local function pickupTypeId( typeInfo, forceFlag )
    if typeInfo then
      if typeInfo:get_typeId(  ) == TransUnit.rootTypeId then
        return 
      end
      if not forceFlag and typeInfo:get_accessMode(  ) ~= "pub" then
        return 
      end
      if typeId2TypeInfo[typeInfo:get_typeId(  )] then
        return 
      end
      typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
      if typeInfo:get_nilable(  ) then
        pickupTypeId( typeInfo:get_orgTypeInfo(  ), true )
      else 
        if typeInfo:get_kind() == TransUnit.TypeInfoKindClass then
          pickupClassMap[typeInfo:get_typeId()] = typeInfo
        end
        local parentInfo = typeInfo:get_parentInfo(  )
        
        pickupTypeId( parentInfo, true )
        local baseInfo = typeInfo:get_baseTypeInfo(  )
        
        if baseInfo:get_typeId() ~= TransUnit.rootTypeId then
          pickupTypeId( baseInfo, true )
        end
        local typeInfoList = typeInfo:get_itemTypeInfoList(  )
        
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo, true )
          end
        end
        typeInfoList = typeInfo:get_argTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo, true )
          end
        end
        typeInfoList = typeInfo:get_retTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo, true )
          end
        end
        typeInfoList = typeInfo:get_children(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            if itemTypeInfo:get_kind(  ) == TransUnit.TypeInfoKindClass or itemTypeInfo:get_kind(  ) == TransUnit.TypeInfoKindFunc or itemTypeInfo:get_kind(  ) == TransUnit.TypeInfoKindMethod then
              pickupTypeId( itemTypeInfo )
            end
          end
        end
        pickupTypeId( typeInfo:get_nilableTypeInfo(  ), true )
      end
    end
  end
  
  local typeId2ClassMap = node:get_typeId2ClassMap(  )
  
  for __index, namespaceInfo in pairs( typeId2ClassMap ) do
    if namespaceInfo.typeInfo:get_accessMode(  ) == "pub" then
      pickupClassMap[namespaceInfo.typeInfo:get_typeId()] = namespaceInfo.typeInfo
    end
  end
  self:writeln( "local _typeId2ClassInfoMap = {}", baseIndent )
  self:writeln( "moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap", baseIndent )
  do
    local __sorted = {}
    local __map = self.classId2TypeInfo
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, classTypeId in ipairs( __sorted ) do
      classTypeInfo = __map[ classTypeId ]
      do
        pickupTypeId( classTypeInfo )
        pickupClassMap[classTypeId] = nil
        self:writeln( "do", baseIndent + stepIndent )
        self:writeln( string.format( "local _classInfo%d = {}", classTypeId), baseIndent + stepIndent )
        self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
        for __index, memberNode in pairs( self.classId2MemberList[classTypeId] or _luneScript.error( 'unwrap val is nil' ) ) do
          if memberNode:get_accessMode() ~= "pri" then
            local memberName = memberNode:get_name().txt
            
            local memberTypeInfo = memberNode:get_expType(  )
            
            self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, memberName), baseIndent + stepIndent )
            self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberNode:get_staticFlag()) .. string.format( "accessMode = '%s', typeId = %d }", memberNode:get_accessMode(), memberTypeInfo:get_typeId(  )), baseIndent + stepIndent )
            pickupTypeId( memberTypeInfo, true )
          end
        end
        self:writeln( "end", baseIndent )
      end
    end
  end
  
  do
    local __sorted = {}
    local __map = pickupClassMap
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, classTypeId in ipairs( __sorted ) do
      classTypeInfo = __map[ classTypeId ]
      do
        local scope = classTypeInfo:get_scope() or _luneScript.error( 'unwrap val is nil' )
        
        if not TransUnit.isBuiltin( classTypeId ) then
          local className = classTypeInfo:getTxt(  )
          
          self:writeln( "do", baseIndent + stepIndent )
          self:writeln( string.format( "local _classInfo%s = {}", classTypeId), baseIndent + stepIndent )
          self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
          pickupTypeId( classTypeInfo )
          do
            local __sorted = {}
            local __map = scope:get_symbol2TypeInfoMap()
            for __key in pairs( __map ) do
              table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, fieldName in ipairs( __sorted ) do
              symbolInfo = __map[ fieldName ]
              do
                local typeInfo = symbolInfo:get_typeInfo()
                
                if typeInfo:get_kind(  ) ~= TransUnit.TypeInfoKindMethod then
                  if symbolInfo:get_accessMode() == "pub" then
                    self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, fieldName), baseIndent + stepIndent )
                    self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, typeInfo:get_staticFlag(  )) .. string.format( "accessMode = '%s', typeId = %d }", symbolInfo:get_accessMode(), typeInfo:get_typeId(  )), baseIndent + stepIndent )
                    pickupTypeId( typeInfo )
                  end
                end
              end
            end
          end
          
          self:writeln( "end", baseIndent )
        end
      end
    end
  end
  
  self:writeln( "local _varName2InfoMap = {}", baseIndent )
  self:writeln( "moduleObj._varName2InfoMap = _varName2InfoMap", baseIndent )
  do
    local __sorted = {}
    local __map = self.pubVarName2InfoMap
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, varName in ipairs( __sorted ) do
      varInfo = __map[ varName ]
      do
        self:writeln( string.format( "_varName2InfoMap.%s = {", varName ), baseIndent )
        self:writeln( string.format( "  name='%s', accessMode = '%s', typeId = %d }", varName, varInfo.accessMode, varInfo.typeInfo:get_typeId()), baseIndent )
        pickupTypeId( varInfo.typeInfo, true )
      end
    end
  end
  
  do
    local __sorted = {}
    local __map = self.pubFuncName2InfoMap
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, funcName in ipairs( __sorted ) do
      funcInfo = __map[ funcName ]
      do
        pickupTypeId( funcInfo.typeInfo, true )
      end
    end
  end
  
  self:writeln( "local _typeInfoList = {}", baseIndent )
  self:writeln( "moduleObj._typeInfoList = _typeInfoList", baseIndent )
  local listIndex = 1
  
  local wroteTypeIdSet = {}
  
  local function outputTypeInfo( typeInfo )
    local typeId = typeInfo:get_typeId(  )
    
    if wroteTypeIdSet[typeId] then
      return 
    end
    wroteTypeIdSet[typeId] = true
    if typeId2TypeInfo[typeId] and not TransUnit.isBuiltin( typeId ) then
      self:write( string.format( "_typeInfoList[%d] = ", listIndex) )
      listIndex = listIndex + 1
      typeInfo:serialize( self )
    end
  end
  
  do
    local __sorted = {}
    local __map = typeId2TypeInfo
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, typeId in ipairs( __sorted ) do
      typeInfo = __map[ typeId ]
      do
        outputTypeInfo( typeInfo )
      end
    end
  end
  
  self:writeln( "----- meta -----", baseIndent )
  if self.stream ~= self.metaStream then
    self:writeln( "return moduleObj", baseIndent )
  end
  self.outMetaFlag = false
end

function convFilter:processRoot( node, parent, baseIndent )
  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local moduleObj = {}", baseIndent )
  local children = node:get_children(  )
  
  for __index, child in pairs( children ) do
    filter( child, self, node, baseIndent )
    self:writeln( "", baseIndent )
  end
  self:outputMeta( node, baseIndent )
  self:writeln( "return moduleObj", baseIndent )
end

-- none

function convFilter:processBlock( node, parent, baseIndent )
  local word = ""
  
  if node:get_blockKind(  ) == "if" or node:get_blockKind(  ) == "elseif" then
    word = "then"
  elseif node:get_blockKind(  ) == "else" then
    word = ""
  elseif node:get_blockKind(  ) == "while" then
    word = "do"
  elseif node:get_blockKind(  ) == "repeat" then
    word = ""
  elseif node:get_blockKind(  ) == "for" then
    word = "do"
  elseif node:get_blockKind(  ) == "apply" then
    word = "do"
  elseif node:get_blockKind(  ) == "foreach" then
    word = "do"
  elseif node:get_blockKind(  ) == "macro" then
    word = ""
  elseif node:get_blockKind(  ) == "func" then
    word = ""
  elseif node:get_blockKind(  ) == "default" then
    word = ""
  elseif node:get_blockKind(  ) == "{" then
    word = "do"
  elseif node:get_blockKind(  ) == "macro" then
    word = ""
  elseif node:get_blockKind(  ) == "let!" then
    word = ""
  elseif node:get_blockKind(  ) == "if!" then
    word = ""
  end
  self:writeln( word, baseIndent + stepIndent )
  local stmtList = node:get_stmtList(  )
  
  for __index, statement in pairs( stmtList ) do
    filter( statement, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:setIndent( baseIndent )
  if node:get_blockKind(  ) == "{" then
    self:write( "end", baseIndent )
  end
end

-- none

function convFilter:processStmtExp( node, parent, baseIndent )
  filter( node:get_exp(  ), self, node, baseIndent )
end

-- none

function convFilter:processDeclClass( node, parent, baseIndent )
  local nodeInfo = node
  
  local classNameToken = nodeInfo:get_name(  )
  
  local className = classNameToken.txt
  
  local classTypeInfo = node:get_expType(  )
  
  local classTypeId = classTypeInfo:get_typeId()
  
  if nodeInfo:get_accessMode(  ) == "pub" then
    self.classId2TypeInfo[classTypeId] = classTypeInfo
  end
  self.classId2MemberList[classTypeId] = nodeInfo:get_memberList(  )
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  local baseInfo = node:get_expType(  ):get_baseTypeInfo(  )
  
  if baseInfo:get_typeId(  ) ~= TransUnit.rootTypeId then
    self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, (baseInfo or _luneScript.error( 'unwrap val is nil' ) ):getTxt(  )), baseIndent )
  end
  if nodeInfo:get_accessMode(  ) == "pub" then
    self:writeln( string.format( "moduleObj.%s = %s", className, className ), baseIndent )
  end
  local hasConstrFlag = false
  
  local memberList = {}
  
  local fieldList = nodeInfo:get_fieldList(  )
  
  local outerMethodSet = nodeInfo:get_outerMethodSet(  )
  
  for __index, field in pairs( fieldList ) do
    local ignoreFlag = false
    
    if field:get_kind() == TransUnit.nodeKind.DeclConstr then
      hasConstrFlag = true
    end
    if field:get_kind() == TransUnit.nodeKind.DeclMember then
      table.insert( memberList, field )
    end
    if field:get_kind() == TransUnit.nodeKind.DeclMethod then
      local methodNode = field
      
      local declInfo = methodNode:get_declInfo(  )
      
      local methodNameToken = declInfo:get_name(  )
      
      if outerMethodSet[(methodNameToken or _luneScript.error( 'unwrap val is nil' ) ).txt] then
        ignoreFlag = true
      end
    end
    if (not ignoreFlag ) then
      filter( field, self, node, baseIndent )
    end
  end
  if not hasConstrFlag then
    local argTxt = ""
    
    for index, member in pairs( memberList ) do
      if index > 1 then
        argTxt = argTxt .. ", "
      end
      argTxt = argTxt .. member:get_name().txt
    end
    self:writeln( string.format( [==[
function %s.new( %s )
  local obj = {}
  setmetatable( obj, { __index = %s } )
  if obj.__init then
    obj:__init( %s )
  end        
  return obj 
 end         
function %s:__init( %s ) 
            ]==], className, argTxt, className, argTxt, className, argTxt), baseIndent )
    for __index, member in pairs( memberList ) do
      local memberName = member:get_name().txt
      
      self:writeln( string.format( "self.%s = %s", memberName, memberName ), baseIndent + stepIndent )
    end
    self:writeln( 'end', baseIndent )
  end
  local scope = nodeInfo:get_scope(  )
  
  for __index, memberNode in pairs( nodeInfo:get_memberList(  ) ) do
    local memberNameToken = memberNode:get_name(  )
    
    local memberName = memberNameToken.txt
    
    local getterName = "get_" .. memberName
    
    local typeInfo = scope:getTypeInfo( getterName, scope, false )
    
    local autoFlag = not typeInfo or (typeInfo or _luneScript.error( 'unwrap val is nil' ) ):get_autoFlag(  )
    
    if memberNode:get_getterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
  return self.%s
end]==], className, getterName, memberName), baseIndent )
    end
    local setterName = "set_" .. memberName
    
    typeInfo = scope:getTypeInfo( setterName, scope, false )
    autoFlag = not typeInfo or (typeInfo or _luneScript.error( 'unwrap val is nil' ) ):get_autoFlag(  )
    if memberNode:get_setterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s( %s )
  self.%s = %s
end]==], className, setterName, memberName, memberName, memberName), baseIndent )
    end
  end
end

-- none

function convFilter:processDeclMember( node, parent, baseIndent )
end

-- none

function convFilter:processExpMacroExp( node, parent, baseIndent )
  local stmtList = node:get_stmtList(  )
  
  if stmtList then
    for __index, stmt in pairs( stmtList ) do
      filter( stmt, self, node, baseIndent )
      self:writeln( "", baseIndent )
    end
  end
end

-- none

function convFilter:processDeclMacro( node, parent, baseIndent )
  if self.inMacro then
    local nodeInfo = node:get_declInfo(  )
    
    local name = nodeInfo:get_name(  )
    
    self:write( string.format( "local function %s(", name.txt) )
    local argTxt = ""
    
    for index, arg in pairs( nodeInfo:get_argList(  ) ) do
      if index > 1 then
        self:write( ", " )
        argTxt = argTxt .. ", "
      end
      filter( arg, self, node, baseIndent )
      if arg:get_kind(  ) == TransUnit.nodeKind.DeclArg then
        argTxt = argTxt .. (arg ):get_name().txt
      else 
        error( string.format( "not support ... in macro %s", node:get_declInfo(  ):get_name().txt) )
      end
    end
    self:writeln( ")", baseIndent )
    self:writeln( "local macroVar = {}", baseIndent )
    self:writeln( "macroVar._names = {}", baseIndent )
    self.macroDepth = self.macroDepth + 1
    do
      local _exp = nodeInfo:get_ast(  )
      if _exp then
      
          filter( _exp, self, node, baseIndent )
        end
    end
    
    self.macroDepth = self.macroDepth - 1
    self:writeln( "", baseIndent )
    self:writeln( "return macroVar", baseIndent )
    self:writeln( "end", baseIndent )
    self:writeln( string.format( "return %s", name.txt), baseIndent )
  end
end

-- none

function convFilter:processExpMacroStat( node, parent, baseIndent )
  for index, token in pairs( node:get_expStrList(  ) ) do
    if index ~= 1 then
      self:write( '..' )
    end
    filter( token, self, node, baseIndent )
  end
end

-- none

function convFilter:processExpNew( node, parent, baseIndent )
  filter( node:get_symbol(  ), self, node, baseIndent )
  self:write( ".new(" )
  do
    local _exp = node:get_argList(  )
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( ")" )
end

-- none

function convFilter:processDeclConstr( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  
  local classNameToken = declInfo:get_className(  ) or _luneScript.error( 'unwrap val is nil' )
  
  local className = classNameToken.txt
  
  self:write( string.format( "function %s.new( ", className ) )
  local argTxt = ""
  
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
      argTxt = argTxt .. ", "
    end
    filter( arg, self, node, baseIndent )
    if arg:get_kind(  ) == TransUnit.nodeKind.DeclArg then
      argTxt = argTxt .. (arg ):get_name().txt
    else 
      local name = node:get_declInfo(  ):get_name() or _luneScript.error( 'unwrap val is nil' )
      
      error( string.format( "not support ... in macro -- %s", name.txt) )
    end
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ), baseIndent + stepIndent )
  self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ), baseIndent )
  self:writeln( "return obj", baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  do
    local _exp = declInfo:get_body(  )
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processExpCallSuper( node, parent, baseIndent )
  local typeInfo = node:get_superType(  )
  
  self:write( string.format( "%s.__init( self, ", typeInfo:getTxt(  )) )
  if node:get_expList(  ) then
    filter( node:get_expList(  ), self, node, baseIndent )
  end
  self:writeln( ")", baseIndent )
end

-- none

function convFilter:processDeclMethod( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  
  local delimit = ":"
  
  if declInfo:get_staticFlag(  ) then
    delimit = "."
  end
  local methodNodeToken = declInfo:get_name(  ) or _luneScript.error( 'unwrap val is nil' )
  
  local methodName = methodNodeToken.txt
  
  local classNameToken = declInfo:get_className(  ) or _luneScript.error( 'unwrap val is nil' )
  
  self:write( string.format( "function %s%s%s( ", classNameToken.txt, delimit, methodName) )
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( arg, self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  do
    local _exp = declInfo:get_body(  )
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processUnwrapSet( node, parent, baseIndent )
  local dstExpList = node:get_dstExpList()
  
  filter( dstExpList, self, node, baseIndent )
  self:write( " = " )
  filter( node:get_srcExpList(), self, node, baseIndent )
  self:writeln( "", baseIndent )
  self:write( "if " )
  for index, expNode in pairs( dstExpList:get_expList() ) do
    if index > 1 then
      self:write( " or " )
    end
    self:write( "not " )
    filter( expNode, self, node, baseIndent )
  end
  self:writeln( " then", baseIndent + stepIndent )
  for index, expNode in pairs( dstExpList:get_expList() ) do
    self:write( string.format( "local _exp%d = ", index) )
    filter( expNode, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  if node:get_unwrapBlock() then
    filter( node:get_unwrapBlock() or _luneScript.error( 'unwrap val is nil' ), self, node, baseIndent + stepIndent )
  end
  self:writeln( "end", baseIndent )
end

function convFilter:processIfUnwrap( node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:write( "local _exp = " )
  filter( node:get_exp(), self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "if _exp then", baseIndent + stepIndent )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  if node:get_nilBlock() then
    self:writeln( "else", baseIndent + stepIndent )
    filter( node:get_nilBlock() or _luneScript.error( 'unwrap val is nil' ), self, node, baseIndent + stepIndent * 2 )
  end
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

function convFilter:processDeclVar( node, parent, baseIndent )
  if node:get_syncBlock() then
    self:writeln( "do", baseIndent + stepIndent )
    for __index, varInfo in pairs( node:get_syncVarList() ) do
      self:writeln( string.format( "local _sync_%s", varInfo:get_name().txt), baseIndent + stepIndent )
    end
    self:writeln( "do", baseIndent + stepIndent * 2 )
  end
  if node:get_mode() ~= "unwrap" and node:get_accessMode(  ) ~= "global" then
    self:write( "local " )
  end
  local varList = node:get_varList(  )
  
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var:get_name().txt )
  end
  do
    local _exp = node:get_expList(  )
    if _exp then
    
        self:write( " = " )
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "", baseIndent )
  do
    local _exp = node:get_unwrapBlock()
    if _exp then
    
        self:writeln( "", baseIndent + stepIndent * 2 )
        self:write( "if " )
        for index, var in pairs( varList ) do
          if index > 1 then
            self:write( " or " )
          end
          self:write( " not " .. var:get_name().txt )
        end
        self:writeln( " then", baseIndent + stepIndent * 3 )
        for index, var in pairs( varList ) do
          self:writeln( string.format( "local _%s = %s", var:get_name().txt, var:get_name().txt), baseIndent + stepIndent * 3 )
        end
        filter( _exp, self, node, baseIndent + stepIndent * 2 )
        do
          local _exp = node:get_thenBlock()
          if _exp then
          
              self:writeln( "else", baseIndent + stepIndent * 3 )
              filter( _exp, self, node, baseIndent + stepIndent * 3 )
            end
        end
        
        self:writeln( "end", baseIndent + stepIndent * 1 )
      end
  end
  
  do
    local _exp = node:get_syncBlock()
    if _exp then
    
        filter( _exp, self, node, baseIndent + stepIndent * 1 )
        for __index, varInfo in pairs( node:get_syncVarList() ) do
          self:writeln( string.format( "_sync_%s = %s", varInfo:get_name().txt, varInfo:get_name().txt), baseIndent + stepIndent )
        end
        self:writeln( "end", baseIndent + stepIndent )
        for __index, varInfo in pairs( node:get_syncVarList() ) do
          self:writeln( string.format( "%s = _sync_%s", varInfo:get_name().txt, varInfo:get_name().txt), baseIndent + stepIndent )
        end
        self:writeln( "end", baseIndent )
      end
  end
  
  if node:get_accessMode(  ) == "pub" then
    self:writeln( "", baseIndent )
    for index, var in pairs( varList ) do
      local name = var:get_name().txt
      
      self:writeln( string.format( "moduleObj.%s = %s", name, name), baseIndent )
      self.pubVarName2InfoMap[name] = PubVerInfo.new(node:get_staticFlag(  ), node:get_accessMode(  ), node:get_typeInfoList(  )[index])
    end
  end
  if self.macroDepth > 0 then
    self:writeln( "", baseIndent )
    for index, var in pairs( varList ) do
      local varName = var:get_name().txt
      
      self:writeln( string.format( "table.insert( macroVar._names, '%s' )", varName), baseIndent )
      self:writeln( string.format( "macroVar.%s = %s", varName, varName), baseIndent )
    end
  end
end

-- none

function convFilter:processDeclArg( node, parent, baseIndent )
  self:write( string.format( "%s", node:get_name(  ).txt ) )
end

-- none

function convFilter:processDeclArgDDD( node, parent, baseIndent )
  self:write( "..." )
end

-- none

function convFilter:processExpDDD( node, parent, baseIndent )
  self:write( "..." )
end

-- none

function convFilter:processDeclFunc( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  
  local nameToken = declInfo:get_name(  )
  
  local name = ""
  
  do
    local _exp = nameToken
    if _exp then
    
        name = _exp.txt
      end
  end
  
  local letTxt = ""
  
  if declInfo:get_accessMode(  ) ~= "global" and #name ~= 0 then
    letTxt = "local "
  end
  self:write( string.format( "%sfunction %s( ", letTxt, name ) )
  local argList = declInfo:get_argList(  )
  
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( arg, self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  do
    local _exp = declInfo:get_body(  )
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:writeln( "end", baseIndent )
  local expType = node:get_expType(  )
  
  if expType:get_accessMode(  ) == "pub" then
    self:write( string.format( "moduleObj.%s = %s", name, name) )
    self.pubFuncName2InfoMap[name] = PubFuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
  end
end

-- none

function convFilter:processRefType( node, parent, baseIndent )
  self:write( (node:get_refFlag(  ) and "&" or "" ) .. (node:get_mutFlag(  ) and "mut " or "" ) )
  filter( node:get_name(  ), self, node, baseIndent )
  if node:get_array(  ) == "array" then
    self:write( "[@]" )
  elseif node:get_array(  ) == "list" then
    self:write( "[]" )
  end
end

-- none

function convFilter:processIf( node, parent, baseIndent )
  local valList = node:get_stmtList(  )
  
  for index, val in pairs( valList ) do
    if index == 1 then
      self:write( "if " )
      filter( val:get_exp(), self, node, baseIndent )
    elseif val:get_kind() == "elseif" then
      self:write( "elseif " )
      filter( val:get_exp(), self, node, baseIndent )
    else 
      self:write( "else" )
    end
    self:write( " " )
    filter( val:get_block(), self, node, baseIndent )
  end
  self:write( "end" )
end

-- none

function convFilter:processSwitch( node, parent, baseIndent )
  self:writeln( "do", baseIndent + 2 )
  self:write( "local _switchExp = " )
  filter( node:get_exp(  ), self, node, baseIndent + 2 )
  self:writeln( "", baseIndent + 2 )
  for index, caseInfo in pairs( node:get_caseList(  ) ) do
    if index == 1 then
      self:write( "if " )
    else 
      self:write( "elseif " )
    end
    local expList = caseInfo:get_expList(  )
    
    for index, expNode in pairs( expList:get_expList(  ) ) do
      if index ~= 1 then
        self:write( " or " )
      end
      self:write( "_switchExp == " )
      filter( expNode, self, node, baseIndent + 2 )
    end
    self:write( " then" )
    filter( caseInfo:get_block(), self, node, baseIndent + 2 )
  end
  do
    local _exp = node:get_default(  )
    if _exp then
    
        self:write( "else " )
        filter( _exp, self, node, baseIndent + 2 )
      end
  end
  
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processWhile( node, parent, baseIndent )
  self:write( "while " )
  filter( node:get_exp(  ), self, node, baseIndent )
  self:write( " " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processRepeat( node, parent, baseIndent )
  self:write( "repeat " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "until " )
  filter( node:get_exp(  ), self, node, baseIndent )
end

-- none

function convFilter:processFor( node, parent, baseIndent )
  self:write( string.format( "for %s = ", node:get_val(  ).txt ) )
  filter( node:get_init(  ), self, node, baseIndent )
  self:write( ", " )
  filter( node:get_to(  ), self, node, baseIndent )
  do
    local _exp = node:get_delta(  )
    if _exp then
    
        self:write( ", " )
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( " " )
  filter( node:get_block(  ), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processApply( node, parent, baseIndent )
  self:write( "for " )
  local varList = node:get_varList(  )
  
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var.txt )
  end
  self:write( " in " )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " " )
  filter( node:get_block(), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processForeach( node, parent, baseIndent )
  self:write( "for " )
  do
    local _exp = node:get_key()
    if _exp then
    
        self:write( _exp.txt )
      else
    
        self:write( "__index" )
      end
  end
  
  self:write( ", " )
  self:write( node:get_val().txt )
  self:write( " in pairs( " )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " ) " )
  filter( node:get_block(), self, node, baseIndent )
  self:write( "end" )
end

-- none

function convFilter:processForsort( node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:writeln( "local __sorted = {}", baseIndent + stepIndent )
  self:write( "local __map = " )
  filter( node:get_exp(), self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "for __key in pairs( __map ) do", baseIndent + stepIndent * 2 )
  self:writeln( "table.insert( __sorted, __key )", baseIndent + stepIndent )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "table.sort( __sorted )", baseIndent + stepIndent )
  self:write( "for __index, " )
  local key = "__key"
  
  do
    local _exp = node:get_key()
    if _exp then
    
        key = _exp.txt
      end
  end
  
  self:write( key )
  self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 )
  self:writeln( string.format( "%s = __map[ %s ]", node:get_val().txt, key ), baseIndent + stepIndent * 2 )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processExpUnwrap( node, parent, baseIndent )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " or " )
  do
    local _exp = node:get_default()
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      else
    
        self:write( "_luneScript.error( 'unwrap val is nil' )" )
      end
  end
  
end

function convFilter:processExpCall( node, parent, baseIndent )
  local wroteFuncFlag = false
  
  if node:get_func():get_kind() == TransUnit.nodeKind.RefField then
    local refField = node:get_func()
    
    if refField:get_prefix():get_expType():get_kind() == TransUnit.TypeInfoKindList then
      wroteFuncFlag = true
      self:write( string.format( "table.%s( ", refField:get_field().txt) )
      filter( refField:get_prefix(), self, refField, baseIndent )
    end
  end
  if not wroteFuncFlag then
    filter( node:get_func(), self, node, baseIndent )
    self:write( "( " )
  end
  do
    local _exp = node:get_argList()
    if _exp then
    
        if wroteFuncFlag then
          if #_exp:get_expList() > 0 then
            self:write( ", " )
          end
        end
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( " )" )
end

-- none

function convFilter:processExpList( node, parent, baseIndent )
  local expList = node:get_expList(  )
  
  for index, exp in pairs( expList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( exp, self, node, baseIndent )
  end
end

-- none

function convFilter:processExpOp1( node, parent, baseIndent )
  local op = node:get_op().txt
  
  if op == ",,," then
    filter( node:get_exp(), self, node, baseIndent )
  elseif op == ",,,," then
    if node:get_macroMode() == "expand" then
      filter( node:get_exp(), self, node, baseIndent )
    else 
      self:write( "_luneSym2Str( " )
      filter( node:get_exp(), self, node, baseIndent )
      self:write( " )" )
    end
  elseif op == ",," then
    self:write( "_luneGetLocal( " )
    filter( node:get_exp(), self, node, baseIndent )
    self:write( " )" )
  else 
    if op == "not" then
      op = op .. " "
    end
    self:write( op )
    filter( node:get_exp(), self, node, baseIndent )
  end
end

-- none

function convFilter:processExpCast( node, parent, baseIndent )
  if node:get_expType() == TransUnit.builtinTypeInt then
    self:write( "math.floor(" )
    filter( node:get_exp(), self, node, baseIndent )
    self:write( ")" )
  else 
    filter( node:get_exp(), self, node, baseIndent )
  end
end

-- none

function convFilter:processExpParen( node, parent, baseIndent )
  self:write( "(" )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " )" )
end

-- none

function convFilter:processExpOp2( node, parent, baseIndent )
  local intCast = false
  
  if node:get_expType() == TransUnit.builtinTypeInt and node:get_op().txt == "/" then
    intCast = true
    self:write( "math.floor(" )
  end
  filter( node:get_exp1(), self, node, baseIndent )
  self:write( " " .. node:get_op().txt .. " " )
  filter( node:get_exp2(), self, node, baseIndent )
  if intCast then
    self:write( ")" )
  end
end

-- none

function convFilter:processExpRef( node, parent, baseIndent )
  self:write( node:get_token().txt )
end

-- none

function convFilter:processExpRefItem( node, parent, baseIndent )
  if node:get_val():get_kind() == TransUnit.nodeKind.LiteralString then
    self:write( "string.byte( " )
    filter( node:get_val(), self, node, baseIndent )
    self:write( ", " )
    filter( node:get_index(), self, node, baseIndent )
    self:write( " )" )
  else 
    filter( node:get_val(), self, node, baseIndent )
    self:write( "[" )
    filter( node:get_index(), self, node, baseIndent )
    self:write( "]" )
  end
end

-- none

function convFilter:processRefField( node, parent, baseIndent )
  local prefix = node:get_prefix(  )
  
  filter( prefix, self, node, baseIndent )
  local delimit = "."
  
  if parent:get_kind() == TransUnit.nodeKind.ExpCall then
    if node:get_expType(  ):get_kind(  ) == TransUnit.TypeInfoKindMethod then
      delimit = ":"
    else 
      delimit = "."
    end
  end
  local fieldToken = node:get_field(  )
  
  self:write( delimit .. fieldToken.txt )
end

-- none

function convFilter:processGetField( node, parent, baseIndent )
  filter( node:get_prefix(  ), self, node, baseIndent )
  local delimit = "."
  
  if node:get_getterTypeInfo(  ):get_kind(  ) == TransUnit.TypeInfoKindMethod then
    delimit = ":"
  else 
    delimit = "."
  end
  local fieldTxt = node:get_field(  ).txt
  
  if node:get_getterTypeInfo(  ) then
    fieldTxt = string.format( "get_%s()", fieldTxt)
  end
  self:write( delimit .. fieldTxt )
end

-- none

function convFilter:processReturn( node, parent, baseIndent )
  self:write( "return " )
  do
    local _exp = node:get_expList()
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
end

-- none

function convFilter:processLiteralList( node, parent, baseIndent )
  self:write( "{" )
  do
    local _exp = node:get_expList()
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( "}" )
end

-- none

function convFilter:processLiteralMap( node, parent, baseIndent )
  self:write( "{" )
  local pairList = node:get_pairList()
  
  for index, pair in pairs( pairList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( "[" )
    filter( pair:get_key(), self, node, baseIndent )
    self:write( "] = " )
    filter( pair:get_val(), self, node, baseIndent )
    index = index + 1
  end
  self:write( "}" )
end

-- none

function convFilter:processLiteralArray( node, parent, baseIndent )
  self:write( "{" )
  do
    local _exp = node:get_expList()
    if _exp then
    
        filter( _exp, self, node, baseIndent )
      end
  end
  
  self:write( "}" )
end

-- none

function convFilter:processLiteralChar( node, parent, baseIndent )
  self:write( string.format( "%g", node:get_num() ) )
end

-- none

function convFilter:processLiteralInt( node, parent, baseIndent )
  self:write( string.format( "%d", node:get_num() ) )
end

-- none

function convFilter:processLiteralReal( node, parent, baseIndent )
  self:write( string.format( "%s", node:get_num() ) )
end

-- none

function convFilter:processLiteralString( node, parent, baseIndent )
  local txt = node:get_token(  ).txt
  
  if string.find( txt, '^```' ) then
    txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
  end
  local argList = node:get_argList(  )
  
  if #argList > 0 then
    self:write( string.format( 'string.format( %s, ', txt ) )
    for index, val in pairs( argList ) do
      if index > 1 then
        self:write( ", " )
      end
      filter( val, self, node, baseIndent )
    end
    self:write( ")" )
  else 
    self:write( txt )
  end
end

-- none

function convFilter:processLiteralBool( node, parent, baseIndent )
  self:write( node:get_token().txt )
end

-- none

function convFilter:processLiteralNil( node, parent, baseIndent )
  self:write( "nil" )
end

-- none

function convFilter:processBreak( node, parent, baseIndent )
  self:write( "break" )
end

-- none

function convFilter:processLiteralSymbol( node, parent, baseIndent )
  self:write( string.format( '%s', node:get_token().txt) )
end

-- none

local MacroEvalImp = {}
setmetatable( MacroEvalImp, { __index = MacroEval } )
moduleObj.MacroEvalImp = MacroEvalImp
function MacroEvalImp:eval( node )
  local oStream = Util.memStream.new()
  
  local conv = convFilter.new("macro", oStream, oStream, "exe", true)
  
  conv:processDeclMacro( node, node, 0 )
  local chunk, err = load( oStream:get_txt(  ) )
  
  if err then
    error( err )
  end
  do
    local _exp = chunk
    if _exp then
    
        local mod = _exp(  )
        
        if not mod then
          error( "macro load error" )
        end
        return (mod or _luneScript.error( 'unwrap val is nil' ) )
      end
  end
  
  error( "failed to load -- " .. node:get_declInfo():get_name() )
end
function MacroEvalImp.new( mode )
  local obj = {}
  setmetatable( obj, { __index = MacroEvalImp } )
  if obj.__init then
    obj:__init( mode )
  end        
  return obj 
 end         
function MacroEvalImp:__init( mode ) 
            
self.mode = mode
  end

return moduleObj
