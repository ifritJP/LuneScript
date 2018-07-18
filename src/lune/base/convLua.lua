--lune/base/convLua.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local VerInfo = {}
function VerInfo.new( staticFlag, accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = VerInfo } )
  if obj.__init then
    obj:__init( staticFlag, accessMode, typeInfo )
  end        
  return obj 
 end         
function VerInfo:__init( staticFlag, accessMode, typeInfo ) 
            
self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.typeInfo = typeInfo
  end

-- none

local FuncInfo = {}
function FuncInfo.new( accessMode, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = FuncInfo } )
  if obj.__init then
    obj:__init( accessMode, typeInfo )
  end        
  return obj 
 end         
function FuncInfo:__init( accessMode, typeInfo ) 
            
self.accessMode = accessMode
  self.typeInfo = typeInfo
  end

local convFilter = {}
setmetatable( convFilter, { __index = Filter } )
moduleObj.convFilter = convFilter
function convFilter.new( streamName, stream, exeFlag, inMacro, moduleTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = convFilter } )
  if obj.__init then obj:__init( streamName, stream, exeFlag, inMacro, moduleTypeInfo ); end
return obj
end
function convFilter:__init(streamName, stream, exeFlag, inMacro, moduleTypeInfo) 
  self.macroDepth = 0
  self.streamName = streamName
  self.stream = stream
  self.moduleName2Info = {}
  self.exeFlag = exeFlag
  self.inMacro = inMacro
  self.indent = 0
  self.curLineNo = 1
  self.className2Scope = {}
  self.className2TypeInfo = {}
  self.className2MemberList = {}
  self.pubVarName2InfoMap = {}
  self.pubFuncName2InfoMap = {}
  self.needIndent = false
  self.moduleTypeInfo = moduleTypeInfo
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
  if self.needIndent then
    self.stream:write( string.rep( " ", self.indent ) )
    self.needIndent = false
  end
  for cr in string.gmatch( txt, "\n" ) do
    self.curLineNo = self.curLineNo + 1
  end
  self.stream:write( txt )
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
  if self.exeFlag then
    self:writeln( string.format( "local %s = _luneScript.loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

-- none

function convFilter:processRoot( node, parent, baseIndent )
  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local moduleObj = {}", baseIndent )
  local children = node:get_children(  )
  local typeId2ClassMap = node:get_typeId2ClassMap(  )
  for __index, child in pairs( children ) do
    filter( child, self, node, baseIndent )
    self:writeln( "", baseIndent )
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
        pickupTypeId( typeInfo:get_orgTypeInfo(  ) )
      else 
        local parentInfo = typeInfo:get_parentInfo(  )
        pickupTypeId( parentInfo )
        local baseInfo = typeInfo:get_baseTypeInfo(  )
        if baseInfo then
          pickupTypeId( baseInfo )
        end
        local typeInfoList = typeInfo:get_itemTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo )
          end
        end
        typeInfoList = typeInfo:get_argTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo )
          end
        end
        typeInfoList = typeInfo:get_retTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo )
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
      end
    end
  end
  
  for __index, namespaceInfo in pairs( typeId2ClassMap ) do
    if namespaceInfo.typeInfo:get_accessMode(  ) == "pub" then
      pickupClassMap[namespaceInfo.name] = namespaceInfo
    end
  end
  self:writeln( "local _typeId2ClassInfoMap = {}", baseIndent )
  self:writeln( "moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap", baseIndent )
  self:writeln( "local _className2InfoMap = {}", baseIndent )
  self:writeln( "moduleObj._className2InfoMap = _className2InfoMap", baseIndent )
  do
    local __sorted = {}
    local __map = self.className2Scope
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, className in ipairs( __sorted ) do
      scope = __map[ className ]
      do
        local classTypeInfo = self.className2TypeInfo[className]
        local classTypeId = classTypeInfo:get_typeId(  )
        pickupTypeId( classTypeInfo )
        pickupClassMap[className] = nil
        self:writeln( "do", baseIndent + stepIndent )
        self:writeln( string.format( "local _classInfo%d = {}", classTypeId), baseIndent + stepIndent )
        self:writeln( string.format( "_className2InfoMap.%s = _classInfo%d", className, classTypeId), baseIndent + stepIndent )
        self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
        for __index, memberNode in pairs( self.className2MemberList[className] ) do
          if memberNode:get_accessMode() == "pub" then
            local memberName = memberNode:get_name().txt
            local memberTypeInfo = memberNode:get_expType(  )
            self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, memberName), baseIndent + stepIndent )
            self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberNode:get_staticFlag()) .. string.format( "accessMode = '%s', methodFlag = false, typeId = %d }", memberNode:get_accessMode(), memberTypeInfo:get_typeId(  )), baseIndent + stepIndent )
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
    for __index, className in ipairs( __sorted ) do
      namespaceInfo = __map[ className ]
      do
        local scope = namespaceInfo.scope
        local classTypeId = namespaceInfo.typeInfo:get_typeId(  )
        if not TransUnit.isBuiltin( classTypeId ) then
          self:writeln( "do", baseIndent + stepIndent )
          self:writeln( string.format( "local _classInfo%s = {}", classTypeId), baseIndent + stepIndent )
          self:writeln( string.format( "_className2InfoMap.%s = _classInfo%d", className, classTypeId), baseIndent + stepIndent )
          self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
          pickupTypeId( namespaceInfo.typeInfo )
          do
            local __sorted = {}
            local __map = scope:get_symbol2TypeInfoMap(  )
            for __key in pairs( __map ) do
              table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, fieldName in ipairs( __sorted ) do
              typeInfo = __map[ fieldName ]
              do
                if typeInfo:get_kind(  ) ~= TransUnit.TypeInfoKindMethod then
                  if typeInfo:get_accessMode(  ) == "pub" then
                    self:writeln( string.format( "_classInfo%d.%s = {", classTypeId, fieldName), baseIndent + stepIndent )
                    self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, typeInfo:get_staticFlag(  )) .. string.format( "accessMode = '%s', methodFlag = false, typeId = %d }", typeInfo:get_accessMode(  ), typeInfo:get_typeId(  )), baseIndent + stepIndent )
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
        self:writeln( string.format( "  name='%s', accessMode = '%s', typeId = %d }", varName, varInfo.accessMode, varInfo.typeInfo:get_typeId(  )), baseIndent )
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
        pickupTypeId( funcInfo.typeInfo )
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
  if nodeInfo:get_accessMode(  ) == "pub" then
    self.className2Scope[className] = nodeInfo:get_scope(  )
    self.className2TypeInfo[className] = node:get_expType(  )
  end
  self.className2MemberList[className] = nodeInfo:get_memberList(  )
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  local baseInfo = node:get_expType(  ):get_baseTypeInfo(  )
  if baseInfo then
    self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, baseInfo:getTxt(  )), baseIndent )
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
      if outerMethodSet[methodNameToken.txt] then
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
    local typeInfo = scope:getTypeInfo( getterName )
    local autoFlag = not typeInfo or typeInfo:get_autoFlag(  )
    if memberNode:get_getterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
  return self.%s
end]==], className, getterName, memberName), baseIndent )
    end
    local setterName = "set_" .. memberName
    typeInfo = scope:getTypeInfo( setterName )
    if memberNode:get_setterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
  return self.%s
end]==], className, setterName, memberName), baseIndent )
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
    if nodeInfo:get_ast(  ) then
      filter( nodeInfo:get_ast(  ), self, node, baseIndent )
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
  if node:get_argList(  ) then
    filter( node:get_argList(  ), self, node, baseIndent )
  end
  self:write( ")" )
end

-- none

function convFilter:processDeclConstr( node, parent, baseIndent )
  local declInfo = node:get_declInfo(  )
  local classNameToken = declInfo:get_className(  )
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
      error( string.format( "not support ... in macro %s", node:get_declInfo(  ):get_name().txt) )
    end
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ), baseIndent + stepIndent )
  self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ), baseIndent )
  self:writeln( "return obj", baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  filter( declInfo:get_body(  ), self, node, baseIndent )
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
  local methodNodeToken = declInfo:get_name(  )
  local methodName = methodNodeToken.txt
  local classNameToken = declInfo:get_className(  )
  self:write( string.format( "function %s%s%s( ", classNameToken.txt, delimit, methodName) )
  local argList = declInfo:get_argList(  )
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    filter( arg, self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  filter( declInfo:get_body(  ), self, node, baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processDeclVar( node, parent, baseIndent )
  if node:get_accessMode(  ) ~= "global" then
    self:write( "local " )
  end
  local varName = ""
  local varList = node:get_varList(  )
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var:get_name().txt )
  end
  if node:get_expList(  ) then
    self:write( " = " )
    filter( node:get_expList(  ), self, node, baseIndent )
  end
  if node:get_unwrap(  ) then
    self:writeln( "", baseIndent )
    self:write( "if " )
    for index, var in pairs( varList ) do
      if index > 1 then
        self:write( " or " )
      end
      self:write( " not " .. var:get_name().txt )
    end
    self:write( " then" )
    filter( node:get_unwrap(  ), self, node, baseIndent )
    self:writeln( "end", baseIndent )
  end
  if node:get_accessMode(  ) == "pub" then
    self:writeln( "", baseIndent )
    local varList = node:get_varList(  )
    for index, var in pairs( varList ) do
      local name = var:get_name().txt
      self:writeln( string.format( "moduleObj.%s = %s", name, name), baseIndent )
      self.pubVarName2InfoMap[name] = VerInfo.new(node:get_staticFlag(  ), node:get_accessMode(  ), node:get_typeInfoList(  )[index])
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
  local name = nameToken and nameToken.txt or ""
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
  filter( declInfo:get_body(  ), self, node, baseIndent )
  self:writeln( "end", baseIndent )
  local expType = node:get_expType(  )
  if expType:get_accessMode(  ) == "pub" then
    self:write( string.format( "moduleObj.%s = %s", name, name) )
    self.pubFuncName2InfoMap[name] = FuncInfo.new(declInfo:get_accessMode(  ), node:get_expType(  ))
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
  if node:get_default(  ) then
    self:write( "else " )
    filter( node:get_default(  ), self, node, baseIndent + 2 )
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
  if node:get_delta(  ) then
    self:write( ", " )
    filter( node:get_delta(  ), self, node, baseIndent )
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
  self:write( node:get_key() and node:get_key().txt or "__index" )
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
  local key = node:get_key() and node:get_key().txt or "__key"
  self:write( key )
  self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 )
  self:writeln( string.format( "%s = __map[ %s ]", node:get_val().txt, key ), baseIndent + stepIndent * 2 )
  filter( node:get_block(), self, node, baseIndent + stepIndent * 2 )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

-- none

function convFilter:processExpCall( node, parent, baseIndent )
  filter( node:get_func(), self, node, baseIndent )
  self:write( "( " )
  if node:get_argList() then
    filter( node:get_argList(), self, node, baseIndent )
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
  if op == ",,," or op == ",,,," then
    filter( node:get_exp(), self, node, baseIndent )
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
  filter( node:get_exp(), self, node, baseIndent )
end

-- none

function convFilter:processExpParen( node, parent, baseIndent )
  self:write( "(" )
  filter( node:get_exp(), self, node, baseIndent )
  self:write( " )" )
end

-- none

function convFilter:processExpOp2( node, parent, baseIndent )
  filter( node:get_exp1(), self, node, baseIndent )
  self:write( " " .. node:get_op().txt .. " " )
  filter( node:get_exp2(), self, node, baseIndent )
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
  filter( node:get_prefix(  ), self, node, baseIndent )
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
  if node:get_expList() then
    filter( node:get_expList(), self, node, baseIndent )
  end
end

-- none

function convFilter:processLiteralList( node, parent, baseIndent )
  self:write( "{" )
  if node:get_expList() then
    filter( node:get_expList(), self, node, baseIndent )
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
  if node:get_expList() then
    filter( node:get_expList(), self, node, baseIndent )
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
  if argList and #argList > 0 then
    self:write( string.format( 'string.format( %s, ', txt ) )
    local argList = node:get_argList(  )
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
  local conv = convFilter.new("macro", oStream, false, true)
  conv:processDeclMacro( node, nil, 0 )
  local chunk, err = load( oStream:get_txt(  ) )
  if err then
    error( err )
  end
  local mod = chunk(  )
  if not mod then
    error( "macro load error" )
  end
  return mod
end
function MacroEvalImp.new(  )
  local obj = {}
  setmetatable( obj, { __index = MacroEvalImp } )
  if obj.__init then
    obj:__init(  )
  end        
  return obj 
 end         
function MacroEvalImp:__init(  ) 
            
end

----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo4722 = {}
  _className2InfoMap.MacroEvalImp = _classInfo4722
  _typeId2ClassInfoMap[ 4722 ] = _classInfo4722
  end
do
  local _classInfo4530 = {}
  _className2InfoMap.convFilter = _classInfo4530
  _typeId2ClassInfoMap[ 4530 ] = _classInfo4530
  end
do
  local _classInfo4424 = {}
  _className2InfoMap.ASTInfo = _classInfo4424
  _typeId2ClassInfoMap[ 4424 ] = _classInfo4424
  end
do
  local _classInfo3724 = {}
  _className2InfoMap.ApplyNode = _classInfo3724
  _typeId2ClassInfoMap[ 3724 ] = _classInfo3724
  end
do
  local _classInfo3590 = {}
  _className2InfoMap.BlockNode = _classInfo3590
  _typeId2ClassInfoMap[ 3590 ] = _classInfo3590
  end
do
  local _classInfo3794 = {}
  _className2InfoMap.BreakNode = _classInfo3794
  _typeId2ClassInfoMap[ 3794 ] = _classInfo3794
  end
do
  local _classInfo3648 = {}
  _className2InfoMap.CaseInfo = _classInfo3648
  _typeId2ClassInfoMap[ 3648 ] = _classInfo3648
  end
do
  local _classInfo4146 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo4146
  _typeId2ClassInfoMap[ 4146 ] = _classInfo4146
  end
do
  local _classInfo4132 = {}
  _className2InfoMap.DeclArgNode = _classInfo4132
  _typeId2ClassInfoMap[ 4132 ] = _classInfo4132
  end
do
  local _classInfo3368 = {}
  _className2InfoMap.DeclClassNode = _classInfo3368
  _typeId2ClassInfoMap[ 3368 ] = _classInfo3368
  end
do
  local _classInfo4084 = {}
  _className2InfoMap.DeclConstrNode = _classInfo4084
  _typeId2ClassInfoMap[ 4084 ] = _classInfo4084
  end
do
  local _classInfo4038 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo4038
  _typeId2ClassInfoMap[ 4038 ] = _classInfo4038
  end
do
  local _classInfo4060 = {}
  _className2InfoMap.DeclFuncNode = _classInfo4060
  _typeId2ClassInfoMap[ 4060 ] = _classInfo4060
  end
do
  local _classInfo3502 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo3502
  _typeId2ClassInfoMap[ 3502 ] = _classInfo3502
  end
do
  local _classInfo4188 = {}
  _className2InfoMap.DeclMacroNode = _classInfo4188
  _typeId2ClassInfoMap[ 4188 ] = _classInfo4188
  end
do
  local _classInfo4110 = {}
  _className2InfoMap.DeclMemberNode = _classInfo4110
  _typeId2ClassInfoMap[ 4110 ] = _classInfo4110
  end
do
  local _classInfo4072 = {}
  _className2InfoMap.DeclMethodNode = _classInfo4072
  _typeId2ClassInfoMap[ 4072 ] = _classInfo4072
  end
do
  local _classInfo4008 = {}
  _className2InfoMap.DeclVarNode = _classInfo4008
  _typeId2ClassInfoMap[ 4008 ] = _classInfo4008
  end
do
  local _classInfo3886 = {}
  _className2InfoMap.ExpCallNode = _classInfo3886
  _typeId2ClassInfoMap[ 3886 ] = _classInfo3886
  end
do
  local _classInfo4096 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo4096
  _typeId2ClassInfoMap[ 4096 ] = _classInfo4096
  end
do
  local _classInfo3846 = {}
  _className2InfoMap.ExpCastNode = _classInfo3846
  _typeId2ClassInfoMap[ 3846 ] = _classInfo3846
  end
do
  local _classInfo3900 = {}
  _className2InfoMap.ExpDDDNode = _classInfo3900
  _typeId2ClassInfoMap[ 3900 ] = _classInfo3900
  end
do
  local _classInfo3500 = {}
  _className2InfoMap.ExpListNode = _classInfo3500
  _typeId2ClassInfoMap[ 3500 ] = _classInfo3500
  end
do
  local _classInfo3924 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo3924
  _typeId2ClassInfoMap[ 3924 ] = _classInfo3924
  end
do
  local _classInfo3940 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo3940
  _typeId2ClassInfoMap[ 3940 ] = _classInfo3940
  end
do
  local _classInfo3804 = {}
  _className2InfoMap.ExpNewNode = _classInfo3804
  _typeId2ClassInfoMap[ 3804 ] = _classInfo3804
  end
do
  local _classInfo3858 = {}
  _className2InfoMap.ExpOp1Node = _classInfo3858
  _typeId2ClassInfoMap[ 3858 ] = _classInfo3858
  end
do
  local _classInfo3830 = {}
  _className2InfoMap.ExpOp2Node = _classInfo3830
  _typeId2ClassInfoMap[ 3830 ] = _classInfo3830
  end
do
  local _classInfo3912 = {}
  _className2InfoMap.ExpParenNode = _classInfo3912
  _typeId2ClassInfoMap[ 3912 ] = _classInfo3912
  end
do
  local _classInfo3872 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo3872
  _typeId2ClassInfoMap[ 3872 ] = _classInfo3872
  end
do
  local _classInfo3818 = {}
  _className2InfoMap.ExpRefNode = _classInfo3818
  _typeId2ClassInfoMap[ 3818 ] = _classInfo3818
  end
do
  local _classInfo3472 = {}
  _className2InfoMap.Filter = _classInfo3472
  _typeId2ClassInfoMap[ 3472 ] = _classInfo3472
  end
do
  local _classInfo3704 = {}
  _className2InfoMap.ForNode = _classInfo3704
  _typeId2ClassInfoMap[ 3704 ] = _classInfo3704
  end
do
  local _classInfo3744 = {}
  _className2InfoMap.ForeachNode = _classInfo3744
  _typeId2ClassInfoMap[ 3744 ] = _classInfo3744
  end
do
  local _classInfo3762 = {}
  _className2InfoMap.ForsortNode = _classInfo3762
  _typeId2ClassInfoMap[ 3762 ] = _classInfo3762
  end
do
  local _classInfo3982 = {}
  _className2InfoMap.GetFieldNode = _classInfo3982
  _typeId2ClassInfoMap[ 3982 ] = _classInfo3982
  end
do
  local _classInfo3618 = {}
  _className2InfoMap.IfNode = _classInfo3618
  _typeId2ClassInfoMap[ 3618 ] = _classInfo3618
  end
do
  local _classInfo3608 = {}
  _className2InfoMap.IfStmtInfo = _classInfo3608
  _typeId2ClassInfoMap[ 3608 ] = _classInfo3608
  end
do
  local _classInfo3538 = {}
  _className2InfoMap.ImportNode = _classInfo3538
  _typeId2ClassInfoMap[ 3538 ] = _classInfo3538
  end
do
  local _classInfo4252 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo4252
  _typeId2ClassInfoMap[ 4252 ] = _classInfo4252
  end
do
  local _classInfo4324 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo4324
  _typeId2ClassInfoMap[ 4324 ] = _classInfo4324
  end
do
  local _classInfo4210 = {}
  _className2InfoMap.LiteralCharNode = _classInfo4210
  _typeId2ClassInfoMap[ 4210 ] = _classInfo4210
  end
do
  local _classInfo4224 = {}
  _className2InfoMap.LiteralIntNode = _classInfo4224
  _typeId2ClassInfoMap[ 4224 ] = _classInfo4224
  end
do
  local _classInfo4264 = {}
  _className2InfoMap.LiteralListNode = _classInfo4264
  _typeId2ClassInfoMap[ 4264 ] = _classInfo4264
  end
do
  local _classInfo4284 = {}
  _className2InfoMap.LiteralMapNode = _classInfo4284
  _typeId2ClassInfoMap[ 4284 ] = _classInfo4284
  end
do
  local _classInfo4200 = {}
  _className2InfoMap.LiteralNilNode = _classInfo4200
  _typeId2ClassInfoMap[ 4200 ] = _classInfo4200
  end
do
  local _classInfo4238 = {}
  _className2InfoMap.LiteralRealNode = _classInfo4238
  _typeId2ClassInfoMap[ 4238 ] = _classInfo4238
  end
do
  local _classInfo4306 = {}
  _className2InfoMap.LiteralStringNode = _classInfo4306
  _typeId2ClassInfoMap[ 4306 ] = _classInfo4306
  end
do
  local _classInfo4336 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo4336
  _typeId2ClassInfoMap[ 4336 ] = _classInfo4336
  end
do
  local _classInfo3498 = {}
  _className2InfoMap.MacroEval = _classInfo3498
  _typeId2ClassInfoMap[ 3498 ] = _classInfo3498
  end
do
  local _classInfo3494 = {}
  _className2InfoMap.NamespaceInfo = _classInfo3494
  _typeId2ClassInfoMap[ 3494 ] = _classInfo3494
  _classInfo3494.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 20 }
  _classInfo3494.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3448 }
  _classInfo3494.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3354 }
  end
do
  local _classInfo3366 = {}
  _className2InfoMap.Node = _classInfo3366
  _typeId2ClassInfoMap[ 3366 ] = _classInfo3366
  end
do
  local _classInfo3528 = {}
  _className2InfoMap.NoneNode = _classInfo3528
  _typeId2ClassInfoMap[ 3528 ] = _classInfo3528
  end
do
  local _classInfo3360 = {}
  _className2InfoMap.OutStream = _classInfo3360
  _typeId2ClassInfoMap[ 3360 ] = _classInfo3360
  end
do
  local _classInfo4276 = {}
  _className2InfoMap.PairItem = _classInfo4276
  _typeId2ClassInfoMap[ 4276 ] = _classInfo4276
  end
do
  local _classInfo3298 = {}
  _className2InfoMap.Parser = _classInfo3298
  _typeId2ClassInfoMap[ 3298 ] = _classInfo3298
  end
do
  local _classInfo3282 = {}
  _className2InfoMap.Position = _classInfo3282
  _typeId2ClassInfoMap[ 3282 ] = _classInfo3282
  _classInfo3282.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo3282.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  end
do
  local _classInfo3968 = {}
  _className2InfoMap.RefFieldNode = _classInfo3968
  _typeId2ClassInfoMap[ 3968 ] = _classInfo3968
  end
do
  local _classInfo3572 = {}
  _className2InfoMap.RefTypeNode = _classInfo3572
  _typeId2ClassInfoMap[ 3572 ] = _classInfo3572
  end
do
  local _classInfo3690 = {}
  _className2InfoMap.RepeatNode = _classInfo3690
  _typeId2ClassInfoMap[ 3690 ] = _classInfo3690
  end
do
  local _classInfo3782 = {}
  _className2InfoMap.ReturnNode = _classInfo3782
  _typeId2ClassInfoMap[ 3782 ] = _classInfo3782
  end
do
  local _classInfo3550 = {}
  _className2InfoMap.RootNode = _classInfo3550
  _typeId2ClassInfoMap[ 3550 ] = _classInfo3550
  end
do
  local _classInfo3448 = {}
  _className2InfoMap.Scope = _classInfo3448
  _typeId2ClassInfoMap[ 3448 ] = _classInfo3448
  end
do
  local _classInfo3956 = {}
  _className2InfoMap.StmtExpNode = _classInfo3956
  _typeId2ClassInfoMap[ 3956 ] = _classInfo3956
  end
do
  local _classInfo3270 = {}
  _className2InfoMap.Stream = _classInfo3270
  _typeId2ClassInfoMap[ 3270 ] = _classInfo3270
  end
do
  local _classInfo3314 = {}
  _className2InfoMap.StreamParser = _classInfo3314
  _typeId2ClassInfoMap[ 3314 ] = _classInfo3314
  _classInfo3314.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3320 }
  end
do
  local _classInfo3656 = {}
  _className2InfoMap.SwitchNode = _classInfo3656
  _typeId2ClassInfoMap[ 3656 ] = _classInfo3656
  end
do
  local _classInfo3286 = {}
  _className2InfoMap.Token = _classInfo3286
  _typeId2ClassInfoMap[ 3286 ] = _classInfo3286
  _classInfo3286.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo3286.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3282 }
  _classInfo3286.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 20 }
  end
do
  local _classInfo112 = {}
  _className2InfoMap.TransUnit = _classInfo112
  _typeId2ClassInfoMap[ 112 ] = _classInfo112
  _classInfo112.ASTInfo = {
    name='ASTInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4424 }
  _classInfo112.ApplyNode = {
    name='ApplyNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3724 }
  _classInfo112.BlockNode = {
    name='BlockNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3590 }
  _classInfo112.BreakNode = {
    name='BreakNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3794 }
  _classInfo112.CaseInfo = {
    name='CaseInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3648 }
  _classInfo112.DeclArgDDDNode = {
    name='DeclArgDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4146 }
  _classInfo112.DeclArgNode = {
    name='DeclArgNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4132 }
  _classInfo112.DeclClassNode = {
    name='DeclClassNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3368 }
  _classInfo112.DeclConstrNode = {
    name='DeclConstrNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4084 }
  _classInfo112.DeclFuncInfo = {
    name='DeclFuncInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4038 }
  _classInfo112.DeclFuncNode = {
    name='DeclFuncNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4060 }
  _classInfo112.DeclMacroInfo = {
    name='DeclMacroInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3502 }
  _classInfo112.DeclMacroNode = {
    name='DeclMacroNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4188 }
  _classInfo112.DeclMemberNode = {
    name='DeclMemberNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4110 }
  _classInfo112.DeclMethodNode = {
    name='DeclMethodNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4072 }
  _classInfo112.DeclVarNode = {
    name='DeclVarNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4008 }
  _classInfo112.ExpCallNode = {
    name='ExpCallNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3886 }
  _classInfo112.ExpCallSuperNode = {
    name='ExpCallSuperNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4096 }
  _classInfo112.ExpCastNode = {
    name='ExpCastNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3846 }
  _classInfo112.ExpDDDNode = {
    name='ExpDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3900 }
  _classInfo112.ExpListNode = {
    name='ExpListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3500 }
  _classInfo112.ExpMacroExpNode = {
    name='ExpMacroExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3924 }
  _classInfo112.ExpMacroStatNode = {
    name='ExpMacroStatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3940 }
  _classInfo112.ExpNewNode = {
    name='ExpNewNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3804 }
  _classInfo112.ExpOp1Node = {
    name='ExpOp1Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3858 }
  _classInfo112.ExpOp2Node = {
    name='ExpOp2Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3830 }
  _classInfo112.ExpParenNode = {
    name='ExpParenNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3912 }
  _classInfo112.ExpRefItemNode = {
    name='ExpRefItemNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3872 }
  _classInfo112.ExpRefNode = {
    name='ExpRefNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3818 }
  _classInfo112.Filter = {
    name='Filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3472 }
  _classInfo112.ForNode = {
    name='ForNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3704 }
  _classInfo112.ForeachNode = {
    name='ForeachNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3744 }
  _classInfo112.ForsortNode = {
    name='ForsortNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3762 }
  _classInfo112.GetFieldNode = {
    name='GetFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3982 }
  _classInfo112.IfNode = {
    name='IfNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3618 }
  _classInfo112.IfStmtInfo = {
    name='IfStmtInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3608 }
  _classInfo112.ImportNode = {
    name='ImportNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3538 }
  _classInfo112.LiteralArrayNode = {
    name='LiteralArrayNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4252 }
  _classInfo112.LiteralBoolNode = {
    name='LiteralBoolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4324 }
  _classInfo112.LiteralCharNode = {
    name='LiteralCharNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4210 }
  _classInfo112.LiteralIntNode = {
    name='LiteralIntNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4224 }
  _classInfo112.LiteralListNode = {
    name='LiteralListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4264 }
  _classInfo112.LiteralMapNode = {
    name='LiteralMapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4284 }
  _classInfo112.LiteralNilNode = {
    name='LiteralNilNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4200 }
  _classInfo112.LiteralRealNode = {
    name='LiteralRealNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4238 }
  _classInfo112.LiteralStringNode = {
    name='LiteralStringNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4306 }
  _classInfo112.LiteralSymbolNode = {
    name='LiteralSymbolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4336 }
  _classInfo112.MacroEval = {
    name='MacroEval', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3498 }
  _classInfo112.NamespaceInfo = {
    name='NamespaceInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3494 }
  _classInfo112.Node = {
    name='Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3366 }
  _classInfo112.NoneNode = {
    name='NoneNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3528 }
  _classInfo112.OutStream = {
    name='OutStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3360 }
  _classInfo112.PairItem = {
    name='PairItem', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4276 }
  _classInfo112.RefFieldNode = {
    name='RefFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3968 }
  _classInfo112.RefTypeNode = {
    name='RefTypeNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3572 }
  _classInfo112.RepeatNode = {
    name='RepeatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3690 }
  _classInfo112.ReturnNode = {
    name='ReturnNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3782 }
  _classInfo112.RootNode = {
    name='RootNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3550 }
  _classInfo112.Scope = {
    name='Scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3448 }
  _classInfo112.StmtExpNode = {
    name='StmtExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3956 }
  _classInfo112.SwitchNode = {
    name='SwitchNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3656 }
  _classInfo112.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3518 }
  _classInfo112.TypeInfo = {
    name='TypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3354 }
  _classInfo112.TypeInfoKindArray = {
    name='TypeInfoKindArray', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindClass = {
    name='TypeInfoKindClass', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindFunc = {
    name='TypeInfoKindFunc', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindList = {
    name='TypeInfoKindList', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindMacro = {
    name='TypeInfoKindMacro', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindMap = {
    name='TypeInfoKindMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindMethod = {
    name='TypeInfoKindMethod', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindModule = {
    name='TypeInfoKindModule', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindNilable = {
    name='TypeInfoKindNilable', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindPrim = {
    name='TypeInfoKindPrim', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.TypeInfoKindRoot = {
    name='TypeInfoKindRoot', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.VarInfo = {
    name='VarInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3998 }
  _classInfo112.WhileNode = {
    name='WhileNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3676 }
  _classInfo112.getNodeKindName = {
    name='getNodeKindName', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3526 }
  _classInfo112.isBuiltin = {
    name='isBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3358 }
  _classInfo112.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3264 }
  _classInfo112.nodeKind = {
    name='nodeKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3524 }
  _classInfo112.rootTypeId = {
    name='rootTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.typeInfoKind = {
    name='typeInfoKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3356 }
  end
do
  local _classInfo4442 = {}
  _className2InfoMap.TxtStream = _classInfo4442
  _typeId2ClassInfoMap[ 4442 ] = _classInfo4442
  end
do
  local _classInfo3354 = {}
  _className2InfoMap.TypeInfo = _classInfo3354
  _typeId2ClassInfoMap[ 3354 ] = _classInfo3354
  _classInfo3354.cloneToPublic = {
    name='cloneToPublic', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3380 }
  _classInfo3354.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3388 }
  _classInfo3354.createArray = {
    name='createArray', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3398 }
  _classInfo3354.createBuiltin = {
    name='createBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3390 }
  _classInfo3354.createClass = {
    name='createClass', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3402 }
  _classInfo3354.createFunc = {
    name='createFunc', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3408 }
  _classInfo3354.createList = {
    name='createList', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3394 }
  _classInfo3354.createMap = {
    name='createMap', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3400 }
  end
do
  local _classInfo4500 = {}
  _className2InfoMap.Util = _classInfo4500
  _typeId2ClassInfoMap[ 4500 ] = _classInfo4500
  _classInfo4500.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 4518 }
  _classInfo4500.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 4516 }
  _classInfo4500.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4508 }
  _classInfo4500.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4502 }
  _classInfo4500.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 4520 }
  end
do
  local _classInfo3998 = {}
  _className2InfoMap.VarInfo = _classInfo3998
  _typeId2ClassInfoMap[ 3998 ] = _classInfo3998
  end
do
  local _classInfo3676 = {}
  _className2InfoMap.WhileNode = _classInfo3676
  _typeId2ClassInfoMap[ 3676 ] = _classInfo3676
  end
do
  local _classInfo3306 = {}
  _className2InfoMap.WrapParser = _classInfo3306
  _typeId2ClassInfoMap[ 3306 ] = _classInfo3306
  end
do
  local _classInfo110 = {}
  _className2InfoMap.base = _classInfo110
  _typeId2ClassInfoMap[ 110 ] = _classInfo110
  _classInfo110.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4434 }
  _classInfo110.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 112 }
  _classInfo110.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4500 }
  end
do
  local _classInfo106 = {}
  _className2InfoMap.convLua = _classInfo106
  _typeId2ClassInfoMap[ 106 ] = _classInfo106
  _classInfo106.MacroEvalImp = {
    name='MacroEvalImp', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4722 }
  _classInfo106.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4434 }
  _classInfo106.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 112 }
  _classInfo106.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4500 }
  _classInfo106.convFilter = {
    name='convFilter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4530 }
  _classInfo106.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 108 }
  _classInfo106.stepIndent = {
    name='stepIndent', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  end
do
  local _classInfo108 = {}
  _className2InfoMap.lune = _classInfo108
  _typeId2ClassInfoMap[ 108 ] = _classInfo108
  _classInfo108.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  end
do
  local _classInfo4508 = {}
  _className2InfoMap.memStream = _classInfo4508
  _typeId2ClassInfoMap[ 4508 ] = _classInfo4508
  end
do
  local _classInfo3334 = {}
  _className2InfoMap.outStream = _classInfo3334
  _typeId2ClassInfoMap[ 3334 ] = _classInfo3334
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 102, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[2] = { parentId = 102, typeId = 104, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106}, }
_typeInfoList[3] = { parentId = 104, typeId = 106, baseId = 1, txt = 'convLua',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {108, 4530, 4722}, }
_typeInfoList[4] = { parentId = 106, typeId = 108, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {110}, }
_typeInfoList[5] = { parentId = 108, typeId = 110, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {112, 4434, 4500}, }
_typeInfoList[6] = { parentId = 110, typeId = 112, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3264, 3354, 3358, 3360, 3366, 3368, 3448, 3472, 3494, 3498, 3500, 3502, 3518, 3526, 3528, 3538, 3550, 3572, 3590, 3608, 3618, 3648, 3656, 3676, 3690, 3704, 3724, 3744, 3762, 3782, 3794, 3804, 3818, 3830, 3846, 3858, 3872, 3886, 3900, 3912, 3924, 3940, 3956, 3968, 3982, 3998, 4008, 4038, 4060, 4072, 4084, 4096, 4110, 4132, 4146, 4188, 4200, 4210, 4224, 4238, 4252, 4264, 4276, 4284, 4306, 4324, 4336, 4424}, }
_typeInfoList[7] = { parentId = 112, typeId = 3264, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3266}, }
_typeInfoList[8] = { parentId = 3264, typeId = 3266, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3268, 3332}, }
_typeInfoList[9] = { parentId = 3266, typeId = 3268, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3270, 3276, 3282, 3286, 3298, 3306, 3314, 3322, 3324, 3326, 3330}, }
_typeInfoList[10] = { parentId = 3268, typeId = 3270, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3272, 3274}, }
_typeInfoList[11] = { parentId = 3270, typeId = 3272, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {20}, children = {}, }
_typeInfoList[12] = { parentId = 3270, typeId = 3274, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[13] = { parentId = 3268, typeId = 3276, baseId = 3270, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3278, 3280}, }
_typeInfoList[14] = { parentId = 3276, typeId = 3278, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 3276, typeId = 3280, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {20}, children = {}, }
_typeInfoList[16] = { parentId = 3268, typeId = 3282, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3284}, }
_typeInfoList[17] = { parentId = 3282, typeId = 3284, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 14}, retTypeId = {}, children = {}, }
_typeInfoList[18] = { parentId = 3268, typeId = 3286, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3290, 3294, 3296}, }
_typeInfoList[19] = { parentId = 1, typeId = 3288, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3286}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[20] = { parentId = 3286, typeId = 3290, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3288}, retTypeId = {}, children = {}, }
_typeInfoList[21] = { parentId = 1, typeId = 3292, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3286}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[22] = { parentId = 3286, typeId = 3294, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3292}, children = {}, }
_typeInfoList[23] = { parentId = 3286, typeId = 3296, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 20, 3282}, retTypeId = {}, children = {}, }
_typeInfoList[24] = { parentId = 3268, typeId = 3298, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3300, 3302, 3304}, }
_typeInfoList[25] = { parentId = 3298, typeId = 3300, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[26] = { parentId = 3298, typeId = 3302, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[27] = { parentId = 3298, typeId = 3304, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[28] = { parentId = 3268, typeId = 3306, baseId = 3298, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3308, 3310, 3312}, }
_typeInfoList[29] = { parentId = 3306, typeId = 3308, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[30] = { parentId = 3306, typeId = 3310, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[31] = { parentId = 3306, typeId = 3312, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3298, 20}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 3268, typeId = 3314, baseId = 3298, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3316, 3318, 3320, 3328}, }
_typeInfoList[33] = { parentId = 3314, typeId = 3316, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3270, 20, 12}, retTypeId = {}, children = {}, }
_typeInfoList[34] = { parentId = 3314, typeId = 3318, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[35] = { parentId = 3314, typeId = 3320, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {3314}, children = {}, }
_typeInfoList[36] = { parentId = 3268, typeId = 3322, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[37] = { parentId = 3268, typeId = 3324, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {12}, children = {}, }
_typeInfoList[38] = { parentId = 3268, typeId = 3326, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {12}, children = {}, }
_typeInfoList[39] = { parentId = 3314, typeId = 3328, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[40] = { parentId = 3268, typeId = 3330, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[41] = { parentId = 3266, typeId = 3332, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3334, 3340, 3348, 3350, 3352}, }
_typeInfoList[42] = { parentId = 3332, typeId = 3334, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3336, 3338}, }
_typeInfoList[43] = { parentId = 3334, typeId = 3336, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 3334, typeId = 3338, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 3332, typeId = 3340, baseId = 3334, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3342, 3344, 3346}, }
_typeInfoList[46] = { parentId = 3340, typeId = 3342, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[47] = { parentId = 3340, typeId = 3344, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[48] = { parentId = 3340, typeId = 3346, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[49] = { parentId = 3332, typeId = 3348, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[50] = { parentId = 3332, typeId = 3350, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[51] = { parentId = 3332, typeId = 3352, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 6, 20}, retTypeId = {}, children = {}, }
_typeInfoList[52] = { parentId = 112, typeId = 3354, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3370, 3372, 3374, 3376, 3378, 3380, 3388, 3390, 3394, 3398, 3400, 3402, 3408, 3412, 3416, 3420, 3422, 3424, 3426, 3428, 3430, 3432, 3434, 3436, 3438, 3440, 3444, 3446}, }
_typeInfoList[53] = { parentId = 112, typeId = 3358, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {12}, children = {}, }
_typeInfoList[54] = { parentId = 112, typeId = 3360, baseId = 1, txt = 'OutStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3362, 3364}, }
_typeInfoList[55] = { parentId = 3360, typeId = 3362, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[56] = { parentId = 3360, typeId = 3364, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 112, typeId = 3366, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3476, 3482, 3484, 3486, 3490, 3492}, }
_typeInfoList[58] = { parentId = 112, typeId = 3368, baseId = 3366, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4158, 4168, 4170, 4172, 4176, 4180, 4182, 4186}, }
_typeInfoList[59] = { parentId = 3354, typeId = 3370, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[60] = { parentId = 3354, typeId = 3372, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[61] = { parentId = 3354, typeId = 3374, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3360}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 3354, typeId = 3376, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[63] = { parentId = 3354, typeId = 3378, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3354, 14}, retTypeId = {12}, children = {}, }
_typeInfoList[64] = { parentId = 3354, typeId = 3380, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3354}, retTypeId = {3354}, children = {}, }
_typeInfoList[65] = { parentId = 1, typeId = 3382, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[66] = { parentId = 1, typeId = 3384, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[67] = { parentId = 1, typeId = 3386, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[68] = { parentId = 3354, typeId = 3388, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3354, 3354, 12, 14, 20, 3382, 3384, 3386}, retTypeId = {3354}, children = {}, }
_typeInfoList[69] = { parentId = 3354, typeId = 3390, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 20, 14, 3354}, retTypeId = {3354}, children = {}, }
_typeInfoList[70] = { parentId = 1, typeId = 3392, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[71] = { parentId = 3354, typeId = 3394, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3354, 3392}, retTypeId = {3354}, children = {}, }
_typeInfoList[72] = { parentId = 1, typeId = 3396, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[73] = { parentId = 3354, typeId = 3398, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3354, 3396}, retTypeId = {3354}, children = {}, }
_typeInfoList[74] = { parentId = 3354, typeId = 3400, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3354, 3354, 3354}, retTypeId = {3354}, children = {}, }
_typeInfoList[75] = { parentId = 3354, typeId = 3402, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3354, 3354, 12, 20, 20}, retTypeId = {3354}, children = {}, }
_typeInfoList[76] = { parentId = 1, typeId = 3404, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 1, typeId = 3406, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 3354, typeId = 3408, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14, 3354, 12, 12, 12, 20, 20, 3404, 3406}, retTypeId = {3354}, children = {}, }
_typeInfoList[79] = { parentId = 1, typeId = 3410, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 3354, typeId = 3412, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3410}, children = {}, }
_typeInfoList[81] = { parentId = 1, typeId = 3414, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[82] = { parentId = 3354, typeId = 3416, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3414}, children = {}, }
_typeInfoList[83] = { parentId = 1, typeId = 3418, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[84] = { parentId = 3354, typeId = 3420, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3418}, children = {}, }
_typeInfoList[85] = { parentId = 3354, typeId = 3422, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[86] = { parentId = 3354, typeId = 3424, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[87] = { parentId = 3354, typeId = 3426, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[88] = { parentId = 3354, typeId = 3428, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[89] = { parentId = 3354, typeId = 3430, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[90] = { parentId = 3354, typeId = 3432, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[91] = { parentId = 3354, typeId = 3434, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[92] = { parentId = 3354, typeId = 3436, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[93] = { parentId = 3354, typeId = 3438, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[94] = { parentId = 3354, typeId = 3440, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[95] = { parentId = 1, typeId = 3442, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 3354, typeId = 3444, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3442}, children = {}, }
_typeInfoList[97] = { parentId = 3354, typeId = 3446, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3354}, retTypeId = {12}, children = {}, }
_typeInfoList[98] = { parentId = 112, typeId = 3448, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3450, 3452, 3454, 3456, 3458, 3460, 3462, 3466, 3470}, }
_typeInfoList[99] = { parentId = 3448, typeId = 3450, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3354}, retTypeId = {}, children = {}, }
_typeInfoList[100] = { parentId = 3448, typeId = 3452, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3354, 3448}, retTypeId = {}, children = {}, }
_typeInfoList[101] = { parentId = 3448, typeId = 3454, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3448}, children = {}, }
_typeInfoList[102] = { parentId = 3448, typeId = 3456, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3354}, children = {}, }
_typeInfoList[103] = { parentId = 3448, typeId = 3458, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3354}, children = {}, }
_typeInfoList[104] = { parentId = 3448, typeId = 3460, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {3354}, children = {}, }
_typeInfoList[105] = { parentId = 3448, typeId = 3462, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3448}, children = {}, }
_typeInfoList[106] = { parentId = 1, typeId = 3464, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 3448, typeId = 3466, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3464}, children = {}, }
_typeInfoList[108] = { parentId = 1, typeId = 3468, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 3448}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 3448, typeId = 3470, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3468}, children = {}, }
_typeInfoList[110] = { parentId = 112, typeId = 3472, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3474, 3530, 3540, 3552, 3574, 3592, 3620, 3634, 3658, 3678, 3692, 3706, 3726, 3746, 3764, 3784, 3796, 3806, 3820, 3832, 3848, 3860, 3874, 3888, 3902, 3914, 3926, 3942, 3958, 3970, 3984, 4010, 4062, 4074, 4086, 4098, 4112, 4134, 4148, 4156, 4190, 4202, 4212, 4226, 4240, 4254, 4266, 4286, 4308, 4326, 4338}, }
_typeInfoList[111] = { parentId = 3472, typeId = 3474, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[112] = { parentId = 3366, typeId = 3476, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[113] = { parentId = 1, typeId = 3478, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 3480, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 3366, typeId = 3482, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3478, 3480}, children = {}, }
_typeInfoList[116] = { parentId = 3366, typeId = 3484, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 3366, typeId = 3486, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[118] = { parentId = 1, typeId = 3488, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 3366, typeId = 3490, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3488}, children = {}, }
_typeInfoList[120] = { parentId = 3366, typeId = 3492, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 3282}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 112, typeId = 3494, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3496}, }
_typeInfoList[122] = { parentId = 3494, typeId = 3496, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3448, 3354}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 112, typeId = 3498, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4420, 4422}, }
_typeInfoList[124] = { parentId = 112, typeId = 3500, baseId = 3366, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3636, 3642, 3646}, }
_typeInfoList[125] = { parentId = 112, typeId = 3502, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3504, 3508, 3510, 3514, 3516}, }
_typeInfoList[126] = { parentId = 3502, typeId = 3504, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[127] = { parentId = 1, typeId = 3506, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 3502, typeId = 3508, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3506}, children = {}, }
_typeInfoList[129] = { parentId = 3502, typeId = 3510, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[130] = { parentId = 1, typeId = 3512, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3286}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 3502, typeId = 3514, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3512}, children = {}, }
_typeInfoList[132] = { parentId = 3502, typeId = 3516, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3286, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 112, typeId = 3518, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3522, 4432}, }
_typeInfoList[134] = { parentId = 1, typeId = 3520, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {20}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 3518, typeId = 3522, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3520}, children = {}, }
_typeInfoList[136] = { parentId = 112, typeId = 3526, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[137] = { parentId = 112, typeId = 3528, baseId = 3366, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3532, 3536}, }
_typeInfoList[138] = { parentId = 3472, typeId = 3530, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3528, 10}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 3528, typeId = 3532, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1, typeId = 3534, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 3528, typeId = 3536, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3534}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 112, typeId = 3538, baseId = 3366, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3542, 3546, 3548}, }
_typeInfoList[143] = { parentId = 3472, typeId = 3540, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3538, 10}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 3538, typeId = 3542, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 3544, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[146] = { parentId = 3538, typeId = 3546, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3544, 20}, retTypeId = {}, children = {}, }
_typeInfoList[147] = { parentId = 3538, typeId = 3548, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[148] = { parentId = 112, typeId = 3550, baseId = 3366, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3554, 3562, 3566, 3570}, }
_typeInfoList[149] = { parentId = 3472, typeId = 3552, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3550, 10}, retTypeId = {}, children = {}, }
_typeInfoList[150] = { parentId = 3550, typeId = 3554, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 1, typeId = 3556, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 1, typeId = 3558, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[153] = { parentId = 1, typeId = 3560, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {14, 3494}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[154] = { parentId = 3550, typeId = 3562, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3556, 3558, 3560}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 1, typeId = 3564, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 3550, typeId = 3566, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3564}, children = {}, }
_typeInfoList[157] = { parentId = 1, typeId = 3568, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {14, 3494}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[158] = { parentId = 3550, typeId = 3570, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3568}, children = {}, }
_typeInfoList[159] = { parentId = 112, typeId = 3572, baseId = 3366, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3576, 3580, 3582, 3584, 3586, 3588}, }
_typeInfoList[160] = { parentId = 3472, typeId = 3574, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3572, 10}, retTypeId = {}, children = {}, }
_typeInfoList[161] = { parentId = 3572, typeId = 3576, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[162] = { parentId = 1, typeId = 3578, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[163] = { parentId = 3572, typeId = 3580, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3578, 3366, 12, 12, 20}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 3572, typeId = 3582, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[165] = { parentId = 3572, typeId = 3584, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[166] = { parentId = 3572, typeId = 3586, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[167] = { parentId = 3572, typeId = 3588, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[168] = { parentId = 112, typeId = 3590, baseId = 3366, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3594, 3600, 3602, 3606}, }
_typeInfoList[169] = { parentId = 3472, typeId = 3592, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3590, 10}, retTypeId = {}, children = {}, }
_typeInfoList[170] = { parentId = 3590, typeId = 3594, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[171] = { parentId = 1, typeId = 3596, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[172] = { parentId = 1, typeId = 3598, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[173] = { parentId = 3590, typeId = 3600, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3596, 20, 3598}, retTypeId = {}, children = {}, }
_typeInfoList[174] = { parentId = 3590, typeId = 3602, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[175] = { parentId = 1, typeId = 3604, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[176] = { parentId = 3590, typeId = 3606, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3604}, children = {}, }
_typeInfoList[177] = { parentId = 112, typeId = 3608, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3610, 3612, 3614, 3616}, }
_typeInfoList[178] = { parentId = 3608, typeId = 3610, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[179] = { parentId = 3608, typeId = 3612, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[180] = { parentId = 3608, typeId = 3614, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[181] = { parentId = 3608, typeId = 3616, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3366, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[182] = { parentId = 112, typeId = 3618, baseId = 3366, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3622, 3628, 3632}, }
_typeInfoList[183] = { parentId = 3472, typeId = 3620, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3618, 10}, retTypeId = {}, children = {}, }
_typeInfoList[184] = { parentId = 3618, typeId = 3622, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[185] = { parentId = 1, typeId = 3624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[186] = { parentId = 1, typeId = 3626, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3608}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 3618, typeId = 3628, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3624, 3626}, retTypeId = {}, children = {}, }
_typeInfoList[188] = { parentId = 1, typeId = 3630, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3608}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[189] = { parentId = 3618, typeId = 3632, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3630}, children = {}, }
_typeInfoList[190] = { parentId = 3472, typeId = 3634, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3500, 10}, retTypeId = {}, children = {}, }
_typeInfoList[191] = { parentId = 3500, typeId = 3636, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[192] = { parentId = 1, typeId = 3638, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[193] = { parentId = 1, typeId = 3640, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[194] = { parentId = 3500, typeId = 3642, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3638, 3640}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 1, typeId = 3644, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[196] = { parentId = 3500, typeId = 3646, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3644}, children = {}, }
_typeInfoList[197] = { parentId = 112, typeId = 3648, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3650, 3652, 3654}, }
_typeInfoList[198] = { parentId = 3648, typeId = 3650, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[199] = { parentId = 3648, typeId = 3652, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[200] = { parentId = 3648, typeId = 3654, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3500, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[201] = { parentId = 112, typeId = 3656, baseId = 3366, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3660, 3666, 3668, 3672, 3674}, }
_typeInfoList[202] = { parentId = 3472, typeId = 3658, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3656, 10}, retTypeId = {}, children = {}, }
_typeInfoList[203] = { parentId = 3656, typeId = 3660, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[204] = { parentId = 1, typeId = 3662, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 1, typeId = 3664, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3648}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 3656, typeId = 3666, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3662, 3366, 3664, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[207] = { parentId = 3656, typeId = 3668, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[208] = { parentId = 1, typeId = 3670, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3648}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[209] = { parentId = 3656, typeId = 3672, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3670}, children = {}, }
_typeInfoList[210] = { parentId = 3656, typeId = 3674, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[211] = { parentId = 112, typeId = 3676, baseId = 3366, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3680, 3684, 3686, 3688}, }
_typeInfoList[212] = { parentId = 3472, typeId = 3678, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3676, 10}, retTypeId = {}, children = {}, }
_typeInfoList[213] = { parentId = 3676, typeId = 3680, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[214] = { parentId = 1, typeId = 3682, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[215] = { parentId = 3676, typeId = 3684, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3682, 3366, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[216] = { parentId = 3676, typeId = 3686, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[217] = { parentId = 3676, typeId = 3688, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[218] = { parentId = 112, typeId = 3690, baseId = 3366, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3694, 3698, 3700, 3702}, }
_typeInfoList[219] = { parentId = 3472, typeId = 3692, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3690, 10}, retTypeId = {}, children = {}, }
_typeInfoList[220] = { parentId = 3690, typeId = 3694, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[221] = { parentId = 1, typeId = 3696, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[222] = { parentId = 3690, typeId = 3698, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3696, 3590, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[223] = { parentId = 3690, typeId = 3700, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[224] = { parentId = 3690, typeId = 3702, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[225] = { parentId = 112, typeId = 3704, baseId = 3366, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3708, 3712, 3714, 3716, 3718, 3720, 3722}, }
_typeInfoList[226] = { parentId = 3472, typeId = 3706, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3704, 10}, retTypeId = {}, children = {}, }
_typeInfoList[227] = { parentId = 3704, typeId = 3708, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 1, typeId = 3710, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[229] = { parentId = 3704, typeId = 3712, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3710, 3590, 3286, 3366, 3366, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[230] = { parentId = 3704, typeId = 3714, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[231] = { parentId = 3704, typeId = 3716, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[232] = { parentId = 3704, typeId = 3718, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[233] = { parentId = 3704, typeId = 3720, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[234] = { parentId = 3704, typeId = 3722, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[235] = { parentId = 112, typeId = 3724, baseId = 3366, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3728, 3734, 3738, 3740, 3742}, }
_typeInfoList[236] = { parentId = 3472, typeId = 3726, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3724, 10}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 3724, typeId = 3728, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 1, typeId = 3730, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[239] = { parentId = 1, typeId = 3732, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3286}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[240] = { parentId = 3724, typeId = 3734, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3730, 3732, 3366, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 1, typeId = 3736, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3286}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 3724, typeId = 3738, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3736}, children = {}, }
_typeInfoList[243] = { parentId = 3724, typeId = 3740, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[244] = { parentId = 3724, typeId = 3742, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[245] = { parentId = 112, typeId = 3744, baseId = 3366, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3748, 3752, 3754, 3756, 3758, 3760}, }
_typeInfoList[246] = { parentId = 3472, typeId = 3746, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3744, 10}, retTypeId = {}, children = {}, }
_typeInfoList[247] = { parentId = 3744, typeId = 3748, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[248] = { parentId = 1, typeId = 3750, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 3744, typeId = 3752, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3750, 3286, 3286, 3366, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[250] = { parentId = 3744, typeId = 3754, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[251] = { parentId = 3744, typeId = 3756, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[252] = { parentId = 3744, typeId = 3758, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[253] = { parentId = 3744, typeId = 3760, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[254] = { parentId = 112, typeId = 3762, baseId = 3366, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3766, 3770, 3772, 3774, 3776, 3778, 3780}, }
_typeInfoList[255] = { parentId = 3472, typeId = 3764, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3762, 10}, retTypeId = {}, children = {}, }
_typeInfoList[256] = { parentId = 3762, typeId = 3766, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[257] = { parentId = 1, typeId = 3768, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[258] = { parentId = 3762, typeId = 3770, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3768, 3286, 3286, 3366, 3590, 12}, retTypeId = {}, children = {}, }
_typeInfoList[259] = { parentId = 3762, typeId = 3772, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[260] = { parentId = 3762, typeId = 3774, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[261] = { parentId = 3762, typeId = 3776, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[262] = { parentId = 3762, typeId = 3778, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[263] = { parentId = 3762, typeId = 3780, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[264] = { parentId = 112, typeId = 3782, baseId = 3366, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3786, 3790, 3792}, }
_typeInfoList[265] = { parentId = 3472, typeId = 3784, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3782, 10}, retTypeId = {}, children = {}, }
_typeInfoList[266] = { parentId = 3782, typeId = 3786, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[267] = { parentId = 1, typeId = 3788, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[268] = { parentId = 3782, typeId = 3790, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3788, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[269] = { parentId = 3782, typeId = 3792, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[270] = { parentId = 112, typeId = 3794, baseId = 3366, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3798, 3802}, }
_typeInfoList[271] = { parentId = 3472, typeId = 3796, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3794, 10}, retTypeId = {}, children = {}, }
_typeInfoList[272] = { parentId = 3794, typeId = 3798, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[273] = { parentId = 1, typeId = 3800, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[274] = { parentId = 3794, typeId = 3802, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3800}, retTypeId = {}, children = {}, }
_typeInfoList[275] = { parentId = 112, typeId = 3804, baseId = 3366, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3808, 3812, 3814, 3816}, }
_typeInfoList[276] = { parentId = 3472, typeId = 3806, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3804, 10}, retTypeId = {}, children = {}, }
_typeInfoList[277] = { parentId = 3804, typeId = 3808, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[278] = { parentId = 1, typeId = 3810, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[279] = { parentId = 3804, typeId = 3812, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3810, 3366, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[280] = { parentId = 3804, typeId = 3814, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[281] = { parentId = 3804, typeId = 3816, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[282] = { parentId = 112, typeId = 3818, baseId = 3366, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3822, 3826, 3828}, }
_typeInfoList[283] = { parentId = 3472, typeId = 3820, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3818, 10}, retTypeId = {}, children = {}, }
_typeInfoList[284] = { parentId = 3818, typeId = 3822, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[285] = { parentId = 1, typeId = 3824, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[286] = { parentId = 3818, typeId = 3826, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3824, 3286}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 3818, typeId = 3828, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[288] = { parentId = 112, typeId = 3830, baseId = 3366, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3834, 3838, 3840, 3842, 3844}, }
_typeInfoList[289] = { parentId = 3472, typeId = 3832, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3830, 10}, retTypeId = {}, children = {}, }
_typeInfoList[290] = { parentId = 3830, typeId = 3834, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[291] = { parentId = 1, typeId = 3836, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[292] = { parentId = 3830, typeId = 3838, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3836, 3286, 3366, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[293] = { parentId = 3830, typeId = 3840, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[294] = { parentId = 3830, typeId = 3842, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[295] = { parentId = 3830, typeId = 3844, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[296] = { parentId = 112, typeId = 3846, baseId = 3366, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3850, 3854, 3856}, }
_typeInfoList[297] = { parentId = 3472, typeId = 3848, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3846, 10}, retTypeId = {}, children = {}, }
_typeInfoList[298] = { parentId = 3846, typeId = 3850, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[299] = { parentId = 1, typeId = 3852, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[300] = { parentId = 3846, typeId = 3854, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3852, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[301] = { parentId = 3846, typeId = 3856, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[302] = { parentId = 112, typeId = 3858, baseId = 3366, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3862, 3866, 3868, 3870}, }
_typeInfoList[303] = { parentId = 3472, typeId = 3860, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3858, 10}, retTypeId = {}, children = {}, }
_typeInfoList[304] = { parentId = 3858, typeId = 3862, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[305] = { parentId = 1, typeId = 3864, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[306] = { parentId = 3858, typeId = 3866, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3864, 3286, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[307] = { parentId = 3858, typeId = 3868, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[308] = { parentId = 3858, typeId = 3870, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[309] = { parentId = 112, typeId = 3872, baseId = 3366, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3876, 3880, 3882, 3884}, }
_typeInfoList[310] = { parentId = 3472, typeId = 3874, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3872, 10}, retTypeId = {}, children = {}, }
_typeInfoList[311] = { parentId = 3872, typeId = 3876, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[312] = { parentId = 1, typeId = 3878, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[313] = { parentId = 3872, typeId = 3880, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3878, 3366, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 3872, typeId = 3882, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[315] = { parentId = 3872, typeId = 3884, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[316] = { parentId = 112, typeId = 3886, baseId = 3366, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3890, 3894, 3896, 3898}, }
_typeInfoList[317] = { parentId = 3472, typeId = 3888, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3886, 10}, retTypeId = {}, children = {}, }
_typeInfoList[318] = { parentId = 3886, typeId = 3890, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[319] = { parentId = 1, typeId = 3892, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[320] = { parentId = 3886, typeId = 3894, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3892, 3366, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[321] = { parentId = 3886, typeId = 3896, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[322] = { parentId = 3886, typeId = 3898, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[323] = { parentId = 112, typeId = 3900, baseId = 3366, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3904, 3908, 3910}, }
_typeInfoList[324] = { parentId = 3472, typeId = 3902, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3900, 10}, retTypeId = {}, children = {}, }
_typeInfoList[325] = { parentId = 3900, typeId = 3904, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[326] = { parentId = 1, typeId = 3906, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[327] = { parentId = 3900, typeId = 3908, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3906, 3286}, retTypeId = {}, children = {}, }
_typeInfoList[328] = { parentId = 3900, typeId = 3910, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[329] = { parentId = 112, typeId = 3912, baseId = 3366, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3916, 3920, 3922}, }
_typeInfoList[330] = { parentId = 3472, typeId = 3914, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3912, 10}, retTypeId = {}, children = {}, }
_typeInfoList[331] = { parentId = 3912, typeId = 3916, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[332] = { parentId = 1, typeId = 3918, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[333] = { parentId = 3912, typeId = 3920, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3918, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[334] = { parentId = 3912, typeId = 3922, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[335] = { parentId = 112, typeId = 3924, baseId = 3366, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3928, 3934, 3938}, }
_typeInfoList[336] = { parentId = 3472, typeId = 3926, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3924, 10}, retTypeId = {}, children = {}, }
_typeInfoList[337] = { parentId = 3924, typeId = 3928, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 1, typeId = 3930, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[339] = { parentId = 1, typeId = 3932, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[340] = { parentId = 3924, typeId = 3934, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3930, 3932}, retTypeId = {}, children = {}, }
_typeInfoList[341] = { parentId = 1, typeId = 3936, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[342] = { parentId = 3924, typeId = 3938, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3936}, children = {}, }
_typeInfoList[343] = { parentId = 112, typeId = 3940, baseId = 3366, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3944, 3950, 3954, 4418}, }
_typeInfoList[344] = { parentId = 3472, typeId = 3942, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3940, 10}, retTypeId = {}, children = {}, }
_typeInfoList[345] = { parentId = 3940, typeId = 3944, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[346] = { parentId = 1, typeId = 3946, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[347] = { parentId = 1, typeId = 3948, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[348] = { parentId = 3940, typeId = 3950, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3946, 3948}, retTypeId = {}, children = {}, }
_typeInfoList[349] = { parentId = 1, typeId = 3952, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[350] = { parentId = 3940, typeId = 3954, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3952}, children = {}, }
_typeInfoList[351] = { parentId = 112, typeId = 3956, baseId = 3366, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3960, 3964, 3966}, }
_typeInfoList[352] = { parentId = 3472, typeId = 3958, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3956, 10}, retTypeId = {}, children = {}, }
_typeInfoList[353] = { parentId = 3956, typeId = 3960, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[354] = { parentId = 1, typeId = 3962, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[355] = { parentId = 3956, typeId = 3964, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3962, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[356] = { parentId = 3956, typeId = 3966, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[357] = { parentId = 112, typeId = 3968, baseId = 3366, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3972, 3976, 3978, 3980, 4412}, }
_typeInfoList[358] = { parentId = 3472, typeId = 3970, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3968, 10}, retTypeId = {}, children = {}, }
_typeInfoList[359] = { parentId = 3968, typeId = 3972, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[360] = { parentId = 1, typeId = 3974, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[361] = { parentId = 3968, typeId = 3976, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3974, 3286, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[362] = { parentId = 3968, typeId = 3978, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[363] = { parentId = 3968, typeId = 3980, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[364] = { parentId = 112, typeId = 3982, baseId = 3366, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3986, 3990, 3992, 3994, 3996}, }
_typeInfoList[365] = { parentId = 3472, typeId = 3984, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3982, 10}, retTypeId = {}, children = {}, }
_typeInfoList[366] = { parentId = 3982, typeId = 3986, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[367] = { parentId = 1, typeId = 3988, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[368] = { parentId = 3982, typeId = 3990, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 3988, 3286, 3366, 3354}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 3982, typeId = 3992, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[370] = { parentId = 3982, typeId = 3994, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[371] = { parentId = 3982, typeId = 3996, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[372] = { parentId = 112, typeId = 3998, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4000, 4002, 4004, 4006}, }
_typeInfoList[373] = { parentId = 3998, typeId = 4000, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[374] = { parentId = 3998, typeId = 4002, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3572}, children = {}, }
_typeInfoList[375] = { parentId = 3998, typeId = 4004, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[376] = { parentId = 3998, typeId = 4006, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3286, 3572, 3354}, retTypeId = {}, children = {}, }
_typeInfoList[377] = { parentId = 112, typeId = 4008, baseId = 3366, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4012, 4020, 4022, 4024, 4028, 4030, 4034, 4036}, }
_typeInfoList[378] = { parentId = 3472, typeId = 4010, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4008, 10}, retTypeId = {}, children = {}, }
_typeInfoList[379] = { parentId = 4008, typeId = 4012, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[380] = { parentId = 1, typeId = 4014, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[381] = { parentId = 1, typeId = 4016, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3998}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[382] = { parentId = 1, typeId = 4018, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[383] = { parentId = 4008, typeId = 4020, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4014, 20, 12, 4016, 3500, 4018, 3590}, retTypeId = {}, children = {}, }
_typeInfoList[384] = { parentId = 4008, typeId = 4022, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[385] = { parentId = 4008, typeId = 4024, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[386] = { parentId = 1, typeId = 4026, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3998}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[387] = { parentId = 4008, typeId = 4028, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4026}, children = {}, }
_typeInfoList[388] = { parentId = 4008, typeId = 4030, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[389] = { parentId = 1, typeId = 4032, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 4008, typeId = 4034, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4032}, children = {}, }
_typeInfoList[391] = { parentId = 4008, typeId = 4036, baseId = 1, txt = 'get_unwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[392] = { parentId = 112, typeId = 4038, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4040, 4042, 4046, 4048, 4050, 4052, 4056, 4058}, }
_typeInfoList[393] = { parentId = 4038, typeId = 4040, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[394] = { parentId = 4038, typeId = 4042, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[395] = { parentId = 1, typeId = 4044, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 4038, typeId = 4046, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4044}, children = {}, }
_typeInfoList[397] = { parentId = 4038, typeId = 4048, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[398] = { parentId = 4038, typeId = 4050, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[399] = { parentId = 4038, typeId = 4052, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[400] = { parentId = 1, typeId = 4054, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[401] = { parentId = 4038, typeId = 4056, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4054}, children = {}, }
_typeInfoList[402] = { parentId = 4038, typeId = 4058, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3286, 3286, 12, 20, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[403] = { parentId = 112, typeId = 4060, baseId = 3366, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4064, 4068, 4070}, }
_typeInfoList[404] = { parentId = 3472, typeId = 4062, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4060, 10}, retTypeId = {}, children = {}, }
_typeInfoList[405] = { parentId = 4060, typeId = 4064, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[406] = { parentId = 1, typeId = 4066, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[407] = { parentId = 4060, typeId = 4068, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4066, 4038}, retTypeId = {}, children = {}, }
_typeInfoList[408] = { parentId = 4060, typeId = 4070, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4038}, children = {}, }
_typeInfoList[409] = { parentId = 112, typeId = 4072, baseId = 3366, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4076, 4080, 4082}, }
_typeInfoList[410] = { parentId = 3472, typeId = 4074, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4072, 10}, retTypeId = {}, children = {}, }
_typeInfoList[411] = { parentId = 4072, typeId = 4076, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[412] = { parentId = 1, typeId = 4078, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[413] = { parentId = 4072, typeId = 4080, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4078, 4038}, retTypeId = {}, children = {}, }
_typeInfoList[414] = { parentId = 4072, typeId = 4082, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4038}, children = {}, }
_typeInfoList[415] = { parentId = 112, typeId = 4084, baseId = 3366, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4088, 4092, 4094}, }
_typeInfoList[416] = { parentId = 3472, typeId = 4086, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4084, 10}, retTypeId = {}, children = {}, }
_typeInfoList[417] = { parentId = 4084, typeId = 4088, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[418] = { parentId = 1, typeId = 4090, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 4084, typeId = 4092, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4090, 4038}, retTypeId = {}, children = {}, }
_typeInfoList[420] = { parentId = 4084, typeId = 4094, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4038}, children = {}, }
_typeInfoList[421] = { parentId = 112, typeId = 4096, baseId = 3366, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4100, 4104, 4106, 4108}, }
_typeInfoList[422] = { parentId = 3472, typeId = 4098, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4096, 10}, retTypeId = {}, children = {}, }
_typeInfoList[423] = { parentId = 4096, typeId = 4100, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[424] = { parentId = 1, typeId = 4102, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[425] = { parentId = 4096, typeId = 4104, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4102, 3354, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[426] = { parentId = 4096, typeId = 4106, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[427] = { parentId = 4096, typeId = 4108, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[428] = { parentId = 112, typeId = 4110, baseId = 3366, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4114, 4118, 4120, 4122, 4124, 4126, 4128, 4130}, }
_typeInfoList[429] = { parentId = 3472, typeId = 4112, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4110, 10}, retTypeId = {}, children = {}, }
_typeInfoList[430] = { parentId = 4110, typeId = 4114, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[431] = { parentId = 1, typeId = 4116, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[432] = { parentId = 4110, typeId = 4118, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4116, 3286, 3572, 12, 20, 20, 20}, retTypeId = {}, children = {}, }
_typeInfoList[433] = { parentId = 4110, typeId = 4120, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[434] = { parentId = 4110, typeId = 4122, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3572}, children = {}, }
_typeInfoList[435] = { parentId = 4110, typeId = 4124, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[436] = { parentId = 4110, typeId = 4126, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[437] = { parentId = 4110, typeId = 4128, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[438] = { parentId = 4110, typeId = 4130, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[439] = { parentId = 112, typeId = 4132, baseId = 3366, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4136, 4140, 4142, 4144}, }
_typeInfoList[440] = { parentId = 3472, typeId = 4134, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4132, 10}, retTypeId = {}, children = {}, }
_typeInfoList[441] = { parentId = 4132, typeId = 4136, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[442] = { parentId = 1, typeId = 4138, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[443] = { parentId = 4132, typeId = 4140, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4138, 3286, 3572}, retTypeId = {}, children = {}, }
_typeInfoList[444] = { parentId = 4132, typeId = 4142, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[445] = { parentId = 4132, typeId = 4144, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3572}, children = {}, }
_typeInfoList[446] = { parentId = 112, typeId = 4146, baseId = 3366, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4150, 4154}, }
_typeInfoList[447] = { parentId = 3472, typeId = 4148, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4146, 10}, retTypeId = {}, children = {}, }
_typeInfoList[448] = { parentId = 4146, typeId = 4150, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[449] = { parentId = 1, typeId = 4152, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[450] = { parentId = 4146, typeId = 4154, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4152}, retTypeId = {}, children = {}, }
_typeInfoList[451] = { parentId = 3472, typeId = 4156, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3368, 10}, retTypeId = {}, children = {}, }
_typeInfoList[452] = { parentId = 3368, typeId = 4158, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[453] = { parentId = 1, typeId = 4160, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[454] = { parentId = 1, typeId = 4162, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[455] = { parentId = 1, typeId = 4164, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4110}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[456] = { parentId = 1, typeId = 4166, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[457] = { parentId = 3368, typeId = 4168, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4160, 20, 3286, 4162, 4164, 3448, 4166}, retTypeId = {}, children = {}, }
_typeInfoList[458] = { parentId = 3368, typeId = 4170, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[459] = { parentId = 3368, typeId = 4172, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[460] = { parentId = 1, typeId = 4174, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[461] = { parentId = 3368, typeId = 4176, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4174}, children = {}, }
_typeInfoList[462] = { parentId = 1, typeId = 4178, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4110}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[463] = { parentId = 3368, typeId = 4180, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4178}, children = {}, }
_typeInfoList[464] = { parentId = 3368, typeId = 4182, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3448}, children = {}, }
_typeInfoList[465] = { parentId = 1, typeId = 4184, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[466] = { parentId = 3368, typeId = 4186, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4184}, children = {}, }
_typeInfoList[467] = { parentId = 112, typeId = 4188, baseId = 3366, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4192, 4196, 4198}, }
_typeInfoList[468] = { parentId = 3472, typeId = 4190, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4188, 10}, retTypeId = {}, children = {}, }
_typeInfoList[469] = { parentId = 4188, typeId = 4192, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[470] = { parentId = 1, typeId = 4194, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[471] = { parentId = 4188, typeId = 4196, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4194, 3502}, retTypeId = {}, children = {}, }
_typeInfoList[472] = { parentId = 4188, typeId = 4198, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3502}, children = {}, }
_typeInfoList[473] = { parentId = 112, typeId = 4200, baseId = 3366, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4204, 4208, 4352}, }
_typeInfoList[474] = { parentId = 3472, typeId = 4202, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4200, 10}, retTypeId = {}, children = {}, }
_typeInfoList[475] = { parentId = 4200, typeId = 4204, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 1, typeId = 4206, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[477] = { parentId = 4200, typeId = 4208, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4206}, retTypeId = {}, children = {}, }
_typeInfoList[478] = { parentId = 112, typeId = 4210, baseId = 3366, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4214, 4218, 4220, 4222, 4358}, }
_typeInfoList[479] = { parentId = 3472, typeId = 4212, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4210, 10}, retTypeId = {}, children = {}, }
_typeInfoList[480] = { parentId = 4210, typeId = 4214, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[481] = { parentId = 1, typeId = 4216, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[482] = { parentId = 4210, typeId = 4218, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4216, 3286, 14}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 4210, typeId = 4220, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[484] = { parentId = 4210, typeId = 4222, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[485] = { parentId = 112, typeId = 4224, baseId = 3366, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4228, 4232, 4234, 4236, 4364}, }
_typeInfoList[486] = { parentId = 3472, typeId = 4226, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4224, 10}, retTypeId = {}, children = {}, }
_typeInfoList[487] = { parentId = 4224, typeId = 4228, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[488] = { parentId = 1, typeId = 4230, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[489] = { parentId = 4224, typeId = 4232, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4230, 3286, 14}, retTypeId = {}, children = {}, }
_typeInfoList[490] = { parentId = 4224, typeId = 4234, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[491] = { parentId = 4224, typeId = 4236, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[492] = { parentId = 112, typeId = 4238, baseId = 3366, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4242, 4246, 4248, 4250, 4370}, }
_typeInfoList[493] = { parentId = 3472, typeId = 4240, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4238, 10}, retTypeId = {}, children = {}, }
_typeInfoList[494] = { parentId = 4238, typeId = 4242, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[495] = { parentId = 1, typeId = 4244, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[496] = { parentId = 4238, typeId = 4246, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4244, 3286, 16}, retTypeId = {}, children = {}, }
_typeInfoList[497] = { parentId = 4238, typeId = 4248, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[498] = { parentId = 4238, typeId = 4250, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {16}, children = {}, }
_typeInfoList[499] = { parentId = 112, typeId = 4252, baseId = 3366, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4256, 4260, 4262, 4376}, }
_typeInfoList[500] = { parentId = 3472, typeId = 4254, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4252, 10}, retTypeId = {}, children = {}, }
_typeInfoList[501] = { parentId = 4252, typeId = 4256, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[502] = { parentId = 1, typeId = 4258, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[503] = { parentId = 4252, typeId = 4260, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4258, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[504] = { parentId = 4252, typeId = 4262, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[505] = { parentId = 112, typeId = 4264, baseId = 3366, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4268, 4272, 4274, 4382}, }
_typeInfoList[506] = { parentId = 3472, typeId = 4266, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4264, 10}, retTypeId = {}, children = {}, }
_typeInfoList[507] = { parentId = 4264, typeId = 4268, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[508] = { parentId = 1, typeId = 4270, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[509] = { parentId = 4264, typeId = 4272, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4270, 3500}, retTypeId = {}, children = {}, }
_typeInfoList[510] = { parentId = 4264, typeId = 4274, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3500}, children = {}, }
_typeInfoList[511] = { parentId = 112, typeId = 4276, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4278, 4280, 4282}, }
_typeInfoList[512] = { parentId = 4276, typeId = 4278, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[513] = { parentId = 4276, typeId = 4280, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[514] = { parentId = 4276, typeId = 4282, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3366, 3366}, retTypeId = {}, children = {}, }
_typeInfoList[515] = { parentId = 112, typeId = 4284, baseId = 3366, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4288, 4296, 4300, 4304, 4388}, }
_typeInfoList[516] = { parentId = 3472, typeId = 4286, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4284, 10}, retTypeId = {}, children = {}, }
_typeInfoList[517] = { parentId = 4284, typeId = 4288, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[518] = { parentId = 1, typeId = 4290, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[519] = { parentId = 1, typeId = 4292, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3366, 3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[520] = { parentId = 1, typeId = 4294, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4276}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[521] = { parentId = 4284, typeId = 4296, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4290, 4292, 4294}, retTypeId = {}, children = {}, }
_typeInfoList[522] = { parentId = 1, typeId = 4298, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3366, 3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[523] = { parentId = 4284, typeId = 4300, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4298}, children = {}, }
_typeInfoList[524] = { parentId = 1, typeId = 4302, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4276}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[525] = { parentId = 4284, typeId = 4304, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4302}, children = {}, }
_typeInfoList[526] = { parentId = 112, typeId = 4306, baseId = 3366, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4310, 4316, 4318, 4322, 4394}, }
_typeInfoList[527] = { parentId = 3472, typeId = 4308, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4306, 10}, retTypeId = {}, children = {}, }
_typeInfoList[528] = { parentId = 4306, typeId = 4310, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[529] = { parentId = 1, typeId = 4312, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[530] = { parentId = 1, typeId = 4314, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[531] = { parentId = 4306, typeId = 4316, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4312, 3286, 4314}, retTypeId = {}, children = {}, }
_typeInfoList[532] = { parentId = 4306, typeId = 4318, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[533] = { parentId = 1, typeId = 4320, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3366}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[534] = { parentId = 4306, typeId = 4322, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4320}, children = {}, }
_typeInfoList[535] = { parentId = 112, typeId = 4324, baseId = 3366, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4328, 4332, 4334, 4400}, }
_typeInfoList[536] = { parentId = 3472, typeId = 4326, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4324, 10}, retTypeId = {}, children = {}, }
_typeInfoList[537] = { parentId = 4324, typeId = 4328, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[538] = { parentId = 1, typeId = 4330, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[539] = { parentId = 4324, typeId = 4332, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4330, 3286}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 4324, typeId = 4334, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[541] = { parentId = 112, typeId = 4336, baseId = 3366, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4340, 4344, 4346, 4406}, }
_typeInfoList[542] = { parentId = 3472, typeId = 4338, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4336, 10}, retTypeId = {}, children = {}, }
_typeInfoList[543] = { parentId = 4336, typeId = 4340, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3472, 10}, retTypeId = {}, children = {}, }
_typeInfoList[544] = { parentId = 1, typeId = 4342, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[545] = { parentId = 4336, typeId = 4344, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3282, 4342, 3286}, retTypeId = {}, children = {}, }
_typeInfoList[546] = { parentId = 4336, typeId = 4346, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3286}, children = {}, }
_typeInfoList[547] = { parentId = 1, typeId = 4348, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[548] = { parentId = 1, typeId = 4350, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[549] = { parentId = 4200, typeId = 4352, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4348, 4350}, children = {}, }
_typeInfoList[550] = { parentId = 1, typeId = 4354, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[551] = { parentId = 1, typeId = 4356, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[552] = { parentId = 4210, typeId = 4358, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4354, 4356}, children = {}, }
_typeInfoList[553] = { parentId = 1, typeId = 4360, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[554] = { parentId = 1, typeId = 4362, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[555] = { parentId = 4224, typeId = 4364, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4360, 4362}, children = {}, }
_typeInfoList[556] = { parentId = 1, typeId = 4366, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[557] = { parentId = 1, typeId = 4368, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[558] = { parentId = 4238, typeId = 4370, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4366, 4368}, children = {}, }
_typeInfoList[559] = { parentId = 1, typeId = 4372, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[560] = { parentId = 1, typeId = 4374, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[561] = { parentId = 4252, typeId = 4376, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4372, 4374}, children = {}, }
_typeInfoList[562] = { parentId = 1, typeId = 4378, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[563] = { parentId = 1, typeId = 4380, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[564] = { parentId = 4264, typeId = 4382, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4378, 4380}, children = {}, }
_typeInfoList[565] = { parentId = 1, typeId = 4384, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[566] = { parentId = 1, typeId = 4386, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 4284, typeId = 4388, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4384, 4386}, children = {}, }
_typeInfoList[568] = { parentId = 1, typeId = 4390, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[569] = { parentId = 1, typeId = 4392, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[570] = { parentId = 4306, typeId = 4394, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4390, 4392}, children = {}, }
_typeInfoList[571] = { parentId = 1, typeId = 4396, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[572] = { parentId = 1, typeId = 4398, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[573] = { parentId = 4324, typeId = 4400, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4396, 4398}, children = {}, }
_typeInfoList[574] = { parentId = 1, typeId = 4402, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[575] = { parentId = 1, typeId = 4404, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[576] = { parentId = 4336, typeId = 4406, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4402, 4404}, children = {}, }
_typeInfoList[577] = { parentId = 1, typeId = 4408, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[578] = { parentId = 1, typeId = 4410, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[579] = { parentId = 3968, typeId = 4412, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4408, 4410}, children = {}, }
_typeInfoList[580] = { parentId = 1, typeId = 4414, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[581] = { parentId = 1, typeId = 4416, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[582] = { parentId = 3940, typeId = 4418, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4414, 4416}, children = {}, }
_typeInfoList[583] = { parentId = 3498, typeId = 4420, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4188}, retTypeId = {28}, children = {}, }
_typeInfoList[584] = { parentId = 3498, typeId = 4422, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[585] = { parentId = 112, typeId = 4424, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4426, 4428, 4430}, }
_typeInfoList[586] = { parentId = 4424, typeId = 4426, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3366}, children = {}, }
_typeInfoList[587] = { parentId = 4424, typeId = 4428, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3354}, children = {}, }
_typeInfoList[588] = { parentId = 4424, typeId = 4430, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3366, 3354}, retTypeId = {}, children = {}, }
_typeInfoList[589] = { parentId = 3518, typeId = 4432, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3298, 12, 20}, retTypeId = {4424}, children = {}, }
_typeInfoList[590] = { parentId = 110, typeId = 4434, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4436, 4442, 4448, 4452, 4464, 4472, 4480, 4490, 4492, 4494, 4498}, }
_typeInfoList[591] = { parentId = 4434, typeId = 4436, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4438, 4440}, }
_typeInfoList[592] = { parentId = 4436, typeId = 4438, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {20}, children = {}, }
_typeInfoList[593] = { parentId = 4436, typeId = 4440, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[594] = { parentId = 4434, typeId = 4442, baseId = 4436, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4444, 4446}, }
_typeInfoList[595] = { parentId = 4442, typeId = 4444, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[596] = { parentId = 4442, typeId = 4446, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {20}, children = {}, }
_typeInfoList[597] = { parentId = 4434, typeId = 4448, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4450}, }
_typeInfoList[598] = { parentId = 4448, typeId = 4450, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 14}, retTypeId = {}, children = {}, }
_typeInfoList[599] = { parentId = 4434, typeId = 4452, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4456, 4460, 4462}, }
_typeInfoList[600] = { parentId = 1, typeId = 4454, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4452}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[601] = { parentId = 4452, typeId = 4456, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4454}, retTypeId = {}, children = {}, }
_typeInfoList[602] = { parentId = 1, typeId = 4458, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4452}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[603] = { parentId = 4452, typeId = 4460, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4458}, children = {}, }
_typeInfoList[604] = { parentId = 4452, typeId = 4462, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 20, 4448}, retTypeId = {}, children = {}, }
_typeInfoList[605] = { parentId = 4434, typeId = 4464, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4466, 4468, 4470}, }
_typeInfoList[606] = { parentId = 4464, typeId = 4466, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4452}, children = {}, }
_typeInfoList[607] = { parentId = 4464, typeId = 4468, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[608] = { parentId = 4464, typeId = 4470, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[609] = { parentId = 4434, typeId = 4472, baseId = 4464, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4474, 4476, 4478}, }
_typeInfoList[610] = { parentId = 4472, typeId = 4474, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4452}, children = {}, }
_typeInfoList[611] = { parentId = 4472, typeId = 4476, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[612] = { parentId = 4472, typeId = 4478, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4464, 20}, retTypeId = {}, children = {}, }
_typeInfoList[613] = { parentId = 4434, typeId = 4480, baseId = 4464, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4482, 4484, 4486, 4496}, }
_typeInfoList[614] = { parentId = 4480, typeId = 4482, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4436, 20, 12}, retTypeId = {}, children = {}, }
_typeInfoList[615] = { parentId = 4480, typeId = 4484, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[616] = { parentId = 4480, typeId = 4486, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {4480}, children = {}, }
_typeInfoList[617] = { parentId = 4434, typeId = 4490, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[618] = { parentId = 4434, typeId = 4492, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {12}, children = {}, }
_typeInfoList[619] = { parentId = 4434, typeId = 4494, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {12}, children = {}, }
_typeInfoList[620] = { parentId = 4480, typeId = 4496, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4452}, children = {}, }
_typeInfoList[621] = { parentId = 4434, typeId = 4498, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {4452}, children = {}, }
_typeInfoList[622] = { parentId = 110, typeId = 4500, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4502, 4508, 4516, 4518, 4520}, }
_typeInfoList[623] = { parentId = 4500, typeId = 4502, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4504, 4506}, }
_typeInfoList[624] = { parentId = 4502, typeId = 4504, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[625] = { parentId = 4502, typeId = 4506, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[626] = { parentId = 4500, typeId = 4508, baseId = 4502, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4510, 4512, 4514}, }
_typeInfoList[627] = { parentId = 4508, typeId = 4510, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[628] = { parentId = 4508, typeId = 4512, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[629] = { parentId = 4508, typeId = 4514, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[630] = { parentId = 4500, typeId = 4516, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[631] = { parentId = 4500, typeId = 4518, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[632] = { parentId = 4500, typeId = 4520, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 6, 20}, retTypeId = {}, children = {}, }
_typeInfoList[633] = { parentId = 106, typeId = 4530, baseId = 3472, txt = 'convFilter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4584, 4586, 4588, 4614, 4616, 4618, 4628, 4630, 4632, 4634, 4636, 4638, 4642, 4644, 4648, 4650, 4652, 4654, 4656, 4658, 4660, 4662, 4664, 4666, 4668, 4670, 4672, 4674, 4676, 4678, 4680, 4682, 4684, 4686, 4688, 4690, 4692, 4694, 4696, 4698, 4700, 4702, 4704, 4706, 4708, 4710, 4714, 4716, 4718, 4720}, }
_typeInfoList[634] = { parentId = 4530, typeId = 4584, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3528, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[635] = { parentId = 4530, typeId = 4586, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3538, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[636] = { parentId = 4530, typeId = 4588, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3550, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[637] = { parentId = 4530, typeId = 4614, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3590, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[638] = { parentId = 4530, typeId = 4616, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3956, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[639] = { parentId = 4530, typeId = 4618, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3368, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[640] = { parentId = 4530, typeId = 4628, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4110, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[641] = { parentId = 4530, typeId = 4630, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3924, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[642] = { parentId = 4530, typeId = 4632, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4188, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[643] = { parentId = 4530, typeId = 4634, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3940, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[644] = { parentId = 4530, typeId = 4636, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3804, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[645] = { parentId = 4530, typeId = 4638, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4084, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[646] = { parentId = 4530, typeId = 4642, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4096, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[647] = { parentId = 4530, typeId = 4644, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4072, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[648] = { parentId = 4530, typeId = 4648, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4008, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[649] = { parentId = 4530, typeId = 4650, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4132, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[650] = { parentId = 4530, typeId = 4652, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4146, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[651] = { parentId = 4530, typeId = 4654, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3900, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[652] = { parentId = 4530, typeId = 4656, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4060, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[653] = { parentId = 4530, typeId = 4658, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3572, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[654] = { parentId = 4530, typeId = 4660, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3618, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[655] = { parentId = 4530, typeId = 4662, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3656, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[656] = { parentId = 4530, typeId = 4664, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3676, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[657] = { parentId = 4530, typeId = 4666, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3690, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[658] = { parentId = 4530, typeId = 4668, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3704, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[659] = { parentId = 4530, typeId = 4670, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3724, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[660] = { parentId = 4530, typeId = 4672, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3744, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[661] = { parentId = 4530, typeId = 4674, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3762, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[662] = { parentId = 4530, typeId = 4676, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3886, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[663] = { parentId = 4530, typeId = 4678, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3500, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[664] = { parentId = 4530, typeId = 4680, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3858, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[665] = { parentId = 4530, typeId = 4682, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3846, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[666] = { parentId = 4530, typeId = 4684, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3912, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[667] = { parentId = 4530, typeId = 4686, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3830, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[668] = { parentId = 4530, typeId = 4688, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3818, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[669] = { parentId = 4530, typeId = 4690, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3872, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[670] = { parentId = 4530, typeId = 4692, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3968, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[671] = { parentId = 4530, typeId = 4694, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3982, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[672] = { parentId = 4530, typeId = 4696, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3782, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[673] = { parentId = 4530, typeId = 4698, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4264, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[674] = { parentId = 4530, typeId = 4700, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4284, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[675] = { parentId = 4530, typeId = 4702, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4252, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[676] = { parentId = 4530, typeId = 4704, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4210, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[677] = { parentId = 4530, typeId = 4706, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4224, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[678] = { parentId = 4530, typeId = 4708, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4238, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[679] = { parentId = 4530, typeId = 4710, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4306, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[680] = { parentId = 4530, typeId = 4714, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4324, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[681] = { parentId = 4530, typeId = 4716, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4200, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[682] = { parentId = 4530, typeId = 4718, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3794, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[683] = { parentId = 4530, typeId = 4720, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4336, 3366, 14}, retTypeId = {}, children = {}, }
_typeInfoList[684] = { parentId = 106, typeId = 4722, baseId = 3498, txt = 'MacroEvalImp',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4724, 4726}, }
_typeInfoList[685] = { parentId = 4722, typeId = 4724, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4188}, retTypeId = {28}, children = {}, }
_typeInfoList[686] = { parentId = 4722, typeId = 4726, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
