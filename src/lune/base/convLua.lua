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
        if baseInfo:get_typeId() ~= TransUnit.rootTypeId then
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
        pickupTypeId( typeInfo:get_nilableTypeInfo(  ) )
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
        local classTypeInfo = self.className2TypeInfo[className] or _luneScript.error( 'unwrap val is nil' )
        local classTypeId = classTypeInfo:get_typeId(  )
        pickupTypeId( classTypeInfo )
        pickupClassMap[className] = nil
        self:writeln( "do", baseIndent + stepIndent )
        self:writeln( string.format( "local _classInfo%d = {}", classTypeId), baseIndent + stepIndent )
        self:writeln( string.format( "_className2InfoMap.%s = _classInfo%d", className, classTypeId), baseIndent + stepIndent )
        self:writeln( string.format( "_typeId2ClassInfoMap[ %d ] = _classInfo%d", classTypeId, classTypeId), baseIndent + stepIndent )
        for __index, memberNode in pairs( self.className2MemberList[className] or _luneScript.error( 'unwrap val is nil' ) ) do
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
  if nodeInfo:get_accessMode(  ) == "pub" then
    self.className2Scope[className] = nodeInfo:get_scope(  )
    self.className2TypeInfo[className] = node:get_expType(  )
  end
  self.className2MemberList[className] = nodeInfo:get_memberList(  )
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
    local autoFlag = not typeInfo or (typeInfo or _luneScript.error( 'unwrap val is nil' ) ):get_autoFlag(  )
    if memberNode:get_getterMode(  ) ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
  return self.%s
end]==], className, getterName, memberName), baseIndent )
    end
    local setterName = "set_" .. memberName
    typeInfo = scope:getTypeInfo( setterName )
    autoFlag = not typeInfo or (typeInfo or _luneScript.error( 'unwrap val is nil' ) ):get_autoFlag(  )
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
  if node:get_unwrapBlock() then
    self:writeln( "", baseIndent )
    self:write( "if " )
    for index, var in pairs( varList ) do
      if index > 1 then
        self:write( " or " )
      end
      self:write( " not " .. var:get_name().txt )
    end
    self:writeln( " then", baseIndent )
    for index, var in pairs( varList ) do
      self:writeln( string.format( "local _%s = %s", var:get_name().txt, var:get_name().txt), baseIndent + stepIndent )
    end
    if node:get_unwrapBlock() then
      filter( node:get_unwrapBlock() or _luneScript.error( 'unwrap val is nil' ), self, node, baseIndent )
    end
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

function convFilter:processExpUnwrap( node, parent, baseIndent )
  filter( node:get_exp(), self, node )
  self:write( " or " )
  if node:get_instead() then
    filter( node:get_instead(), self, node )
  else 
    self:write( "_luneScript.error( 'unwrap val is nil' )" )
  end
end

function convFilter:processExpCall( node, parent, baseIndent )
  filter( node:get_func(), self, node, baseIndent )
  self:write( "( " )
  do
    local _exp = node:get_argList()
    if _exp then
    
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
  local _classInfo5010 = {}
  _className2InfoMap.MacroEvalImp = _classInfo5010
  _typeId2ClassInfoMap[ 5010 ] = _classInfo5010
  end
do
  local _classInfo4812 = {}
  _className2InfoMap.convFilter = _classInfo4812
  _typeId2ClassInfoMap[ 4812 ] = _classInfo4812
  end
do
  local _classInfo4706 = {}
  _className2InfoMap.ASTInfo = _classInfo4706
  _typeId2ClassInfoMap[ 4706 ] = _classInfo4706
  end
do
  local _classInfo3958 = {}
  _className2InfoMap.ApplyNode = _classInfo3958
  _typeId2ClassInfoMap[ 3958 ] = _classInfo3958
  end
do
  local _classInfo3824 = {}
  _className2InfoMap.BlockNode = _classInfo3824
  _typeId2ClassInfoMap[ 3824 ] = _classInfo3824
  end
do
  local _classInfo4028 = {}
  _className2InfoMap.BreakNode = _classInfo4028
  _typeId2ClassInfoMap[ 4028 ] = _classInfo4028
  end
do
  local _classInfo3882 = {}
  _className2InfoMap.CaseInfo = _classInfo3882
  _typeId2ClassInfoMap[ 3882 ] = _classInfo3882
  end
do
  local _classInfo4428 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo4428
  _typeId2ClassInfoMap[ 4428 ] = _classInfo4428
  end
do
  local _classInfo4414 = {}
  _className2InfoMap.DeclArgNode = _classInfo4414
  _typeId2ClassInfoMap[ 4414 ] = _classInfo4414
  end
do
  local _classInfo3598 = {}
  _className2InfoMap.DeclClassNode = _classInfo3598
  _typeId2ClassInfoMap[ 3598 ] = _classInfo3598
  end
do
  local _classInfo4366 = {}
  _className2InfoMap.DeclConstrNode = _classInfo4366
  _typeId2ClassInfoMap[ 4366 ] = _classInfo4366
  end
do
  local _classInfo4320 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo4320
  _typeId2ClassInfoMap[ 4320 ] = _classInfo4320
  end
do
  local _classInfo4342 = {}
  _className2InfoMap.DeclFuncNode = _classInfo4342
  _typeId2ClassInfoMap[ 4342 ] = _classInfo4342
  end
do
  local _classInfo3736 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo3736
  _typeId2ClassInfoMap[ 3736 ] = _classInfo3736
  end
do
  local _classInfo4470 = {}
  _className2InfoMap.DeclMacroNode = _classInfo4470
  _typeId2ClassInfoMap[ 4470 ] = _classInfo4470
  end
do
  local _classInfo4392 = {}
  _className2InfoMap.DeclMemberNode = _classInfo4392
  _typeId2ClassInfoMap[ 4392 ] = _classInfo4392
  end
do
  local _classInfo4354 = {}
  _className2InfoMap.DeclMethodNode = _classInfo4354
  _typeId2ClassInfoMap[ 4354 ] = _classInfo4354
  end
do
  local _classInfo4288 = {}
  _className2InfoMap.DeclVarNode = _classInfo4288
  _typeId2ClassInfoMap[ 4288 ] = _classInfo4288
  end
do
  local _classInfo4166 = {}
  _className2InfoMap.ExpCallNode = _classInfo4166
  _typeId2ClassInfoMap[ 4166 ] = _classInfo4166
  end
do
  local _classInfo4378 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo4378
  _typeId2ClassInfoMap[ 4378 ] = _classInfo4378
  end
do
  local _classInfo4126 = {}
  _className2InfoMap.ExpCastNode = _classInfo4126
  _typeId2ClassInfoMap[ 4126 ] = _classInfo4126
  end
do
  local _classInfo4180 = {}
  _className2InfoMap.ExpDDDNode = _classInfo4180
  _typeId2ClassInfoMap[ 4180 ] = _classInfo4180
  end
do
  local _classInfo3734 = {}
  _className2InfoMap.ExpListNode = _classInfo3734
  _typeId2ClassInfoMap[ 3734 ] = _classInfo3734
  end
do
  local _classInfo4204 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo4204
  _typeId2ClassInfoMap[ 4204 ] = _classInfo4204
  end
do
  local _classInfo4220 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo4220
  _typeId2ClassInfoMap[ 4220 ] = _classInfo4220
  end
do
  local _classInfo4038 = {}
  _className2InfoMap.ExpNewNode = _classInfo4038
  _typeId2ClassInfoMap[ 4038 ] = _classInfo4038
  end
do
  local _classInfo4138 = {}
  _className2InfoMap.ExpOp1Node = _classInfo4138
  _typeId2ClassInfoMap[ 4138 ] = _classInfo4138
  end
do
  local _classInfo4078 = {}
  _className2InfoMap.ExpOp2Node = _classInfo4078
  _typeId2ClassInfoMap[ 4078 ] = _classInfo4078
  end
do
  local _classInfo4192 = {}
  _className2InfoMap.ExpParenNode = _classInfo4192
  _typeId2ClassInfoMap[ 4192 ] = _classInfo4192
  end
do
  local _classInfo4152 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo4152
  _typeId2ClassInfoMap[ 4152 ] = _classInfo4152
  end
do
  local _classInfo4066 = {}
  _className2InfoMap.ExpRefNode = _classInfo4066
  _typeId2ClassInfoMap[ 4066 ] = _classInfo4066
  end
do
  local _classInfo4052 = {}
  _className2InfoMap.ExpUnwrapNode = _classInfo4052
  _typeId2ClassInfoMap[ 4052 ] = _classInfo4052
  end
do
  local _classInfo3706 = {}
  _className2InfoMap.Filter = _classInfo3706
  _typeId2ClassInfoMap[ 3706 ] = _classInfo3706
  end
do
  local _classInfo3938 = {}
  _className2InfoMap.ForNode = _classInfo3938
  _typeId2ClassInfoMap[ 3938 ] = _classInfo3938
  end
do
  local _classInfo3978 = {}
  _className2InfoMap.ForeachNode = _classInfo3978
  _typeId2ClassInfoMap[ 3978 ] = _classInfo3978
  end
do
  local _classInfo3996 = {}
  _className2InfoMap.ForsortNode = _classInfo3996
  _typeId2ClassInfoMap[ 3996 ] = _classInfo3996
  end
do
  local _classInfo4262 = {}
  _className2InfoMap.GetFieldNode = _classInfo4262
  _typeId2ClassInfoMap[ 4262 ] = _classInfo4262
  end
do
  local _classInfo3852 = {}
  _className2InfoMap.IfNode = _classInfo3852
  _typeId2ClassInfoMap[ 3852 ] = _classInfo3852
  end
do
  local _classInfo3842 = {}
  _className2InfoMap.IfStmtInfo = _classInfo3842
  _typeId2ClassInfoMap[ 3842 ] = _classInfo3842
  end
do
  local _classInfo4110 = {}
  _className2InfoMap.IfUnwrapNode = _classInfo4110
  _typeId2ClassInfoMap[ 4110 ] = _classInfo4110
  end
do
  local _classInfo3772 = {}
  _className2InfoMap.ImportNode = _classInfo3772
  _typeId2ClassInfoMap[ 3772 ] = _classInfo3772
  end
do
  local _classInfo4534 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo4534
  _typeId2ClassInfoMap[ 4534 ] = _classInfo4534
  end
do
  local _classInfo4606 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo4606
  _typeId2ClassInfoMap[ 4606 ] = _classInfo4606
  end
do
  local _classInfo4492 = {}
  _className2InfoMap.LiteralCharNode = _classInfo4492
  _typeId2ClassInfoMap[ 4492 ] = _classInfo4492
  end
do
  local _classInfo4506 = {}
  _className2InfoMap.LiteralIntNode = _classInfo4506
  _typeId2ClassInfoMap[ 4506 ] = _classInfo4506
  end
do
  local _classInfo4546 = {}
  _className2InfoMap.LiteralListNode = _classInfo4546
  _typeId2ClassInfoMap[ 4546 ] = _classInfo4546
  end
do
  local _classInfo4566 = {}
  _className2InfoMap.LiteralMapNode = _classInfo4566
  _typeId2ClassInfoMap[ 4566 ] = _classInfo4566
  end
do
  local _classInfo4482 = {}
  _className2InfoMap.LiteralNilNode = _classInfo4482
  _typeId2ClassInfoMap[ 4482 ] = _classInfo4482
  end
do
  local _classInfo4520 = {}
  _className2InfoMap.LiteralRealNode = _classInfo4520
  _typeId2ClassInfoMap[ 4520 ] = _classInfo4520
  end
do
  local _classInfo4588 = {}
  _className2InfoMap.LiteralStringNode = _classInfo4588
  _typeId2ClassInfoMap[ 4588 ] = _classInfo4588
  end
do
  local _classInfo4618 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo4618
  _typeId2ClassInfoMap[ 4618 ] = _classInfo4618
  end
do
  local _classInfo3732 = {}
  _className2InfoMap.MacroEval = _classInfo3732
  _typeId2ClassInfoMap[ 3732 ] = _classInfo3732
  end
do
  local _classInfo3728 = {}
  _className2InfoMap.NamespaceInfo = _classInfo3728
  _typeId2ClassInfoMap[ 3728 ] = _classInfo3728
  _classInfo3728.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 20 }
  _classInfo3728.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3682 }
  _classInfo3728.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3526 }
  end
do
  local _classInfo3596 = {}
  _className2InfoMap.Node = _classInfo3596
  _typeId2ClassInfoMap[ 3596 ] = _classInfo3596
  end
do
  local _classInfo3762 = {}
  _className2InfoMap.NoneNode = _classInfo3762
  _typeId2ClassInfoMap[ 3762 ] = _classInfo3762
  end
do
  local _classInfo3600 = {}
  _className2InfoMap.NormalTypeInfo = _classInfo3600
  _typeId2ClassInfoMap[ 3600 ] = _classInfo3600
  _classInfo3600.cloneToPublic = {
    name='cloneToPublic', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3612 }
  _classInfo3600.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3620 }
  _classInfo3600.createArray = {
    name='createArray', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3668 }
  _classInfo3600.createBuiltin = {
    name='createBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3660 }
  _classInfo3600.createClass = {
    name='createClass', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3672 }
  _classInfo3600.createFunc = {
    name='createFunc', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3678 }
  _classInfo3600.createList = {
    name='createList', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3664 }
  _classInfo3600.createMap = {
    name='createMap', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3670 }
  end
do
  local _classInfo3532 = {}
  _className2InfoMap.OutStream = _classInfo3532
  _typeId2ClassInfoMap[ 3532 ] = _classInfo3532
  end
do
  local _classInfo4558 = {}
  _className2InfoMap.PairItem = _classInfo4558
  _typeId2ClassInfoMap[ 4558 ] = _classInfo4558
  end
do
  local _classInfo3440 = {}
  _className2InfoMap.Parser = _classInfo3440
  _typeId2ClassInfoMap[ 3440 ] = _classInfo3440
  _classInfo3440.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3470 }
  _classInfo3440.Position = {
    name='Position', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3454 }
  _classInfo3440.Stream = {
    name='Stream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3442 }
  _classInfo3440.StreamParser = {
    name='StreamParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3486 }
  _classInfo3440.Token = {
    name='Token', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3458 }
  _classInfo3440.TxtStream = {
    name='TxtStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3448 }
  _classInfo3440.WrapParser = {
    name='WrapParser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3478 }
  _classInfo3440.getEofToken = {
    name='getEofToken', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3502 }
  _classInfo3440.getKindTxt = {
    name='getKindTxt', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3494 }
  _classInfo3440.isOp1 = {
    name='isOp1', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3498 }
  _classInfo3440.isOp2 = {
    name='isOp2', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3496 }
  end
do
  local _classInfo3454 = {}
  _className2InfoMap.Position = _classInfo3454
  _typeId2ClassInfoMap[ 3454 ] = _classInfo3454
  _classInfo3454.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo3454.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  end
do
  local _classInfo4248 = {}
  _className2InfoMap.RefFieldNode = _classInfo4248
  _typeId2ClassInfoMap[ 4248 ] = _classInfo4248
  end
do
  local _classInfo3806 = {}
  _className2InfoMap.RefTypeNode = _classInfo3806
  _typeId2ClassInfoMap[ 3806 ] = _classInfo3806
  end
do
  local _classInfo3924 = {}
  _className2InfoMap.RepeatNode = _classInfo3924
  _typeId2ClassInfoMap[ 3924 ] = _classInfo3924
  end
do
  local _classInfo4016 = {}
  _className2InfoMap.ReturnNode = _classInfo4016
  _typeId2ClassInfoMap[ 4016 ] = _classInfo4016
  end
do
  local _classInfo3784 = {}
  _className2InfoMap.RootNode = _classInfo3784
  _typeId2ClassInfoMap[ 3784 ] = _classInfo3784
  end
do
  local _classInfo3682 = {}
  _className2InfoMap.Scope = _classInfo3682
  _typeId2ClassInfoMap[ 3682 ] = _classInfo3682
  end
do
  local _classInfo4236 = {}
  _className2InfoMap.StmtExpNode = _classInfo4236
  _typeId2ClassInfoMap[ 4236 ] = _classInfo4236
  end
do
  local _classInfo3442 = {}
  _className2InfoMap.Stream = _classInfo3442
  _typeId2ClassInfoMap[ 3442 ] = _classInfo3442
  end
do
  local _classInfo4762 = {}
  _className2InfoMap.StreamParser = _classInfo4762
  _typeId2ClassInfoMap[ 4762 ] = _classInfo4762
  _classInfo4762.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 4768 }
  end
do
  local _classInfo3890 = {}
  _className2InfoMap.SwitchNode = _classInfo3890
  _typeId2ClassInfoMap[ 3890 ] = _classInfo3890
  end
do
  local _classInfo3458 = {}
  _className2InfoMap.Token = _classInfo3458
  _typeId2ClassInfoMap[ 3458 ] = _classInfo3458
  _classInfo3458.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo3458.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3454 }
  _classInfo3458.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 20 }
  end
do
  local _classInfo112 = {}
  _className2InfoMap.TransUnit = _classInfo112
  _typeId2ClassInfoMap[ 112 ] = _classInfo112
  _classInfo112.ASTInfo = {
    name='ASTInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4706 }
  _classInfo112.ApplyNode = {
    name='ApplyNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3958 }
  _classInfo112.BlockNode = {
    name='BlockNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3824 }
  _classInfo112.BreakNode = {
    name='BreakNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4028 }
  _classInfo112.CaseInfo = {
    name='CaseInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3882 }
  _classInfo112.DeclArgDDDNode = {
    name='DeclArgDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4428 }
  _classInfo112.DeclArgNode = {
    name='DeclArgNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4414 }
  _classInfo112.DeclClassNode = {
    name='DeclClassNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3598 }
  _classInfo112.DeclConstrNode = {
    name='DeclConstrNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4366 }
  _classInfo112.DeclFuncInfo = {
    name='DeclFuncInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4320 }
  _classInfo112.DeclFuncNode = {
    name='DeclFuncNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4342 }
  _classInfo112.DeclMacroInfo = {
    name='DeclMacroInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3736 }
  _classInfo112.DeclMacroNode = {
    name='DeclMacroNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4470 }
  _classInfo112.DeclMemberNode = {
    name='DeclMemberNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4392 }
  _classInfo112.DeclMethodNode = {
    name='DeclMethodNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4354 }
  _classInfo112.DeclVarNode = {
    name='DeclVarNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4288 }
  _classInfo112.ExpCallNode = {
    name='ExpCallNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4166 }
  _classInfo112.ExpCallSuperNode = {
    name='ExpCallSuperNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4378 }
  _classInfo112.ExpCastNode = {
    name='ExpCastNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4126 }
  _classInfo112.ExpDDDNode = {
    name='ExpDDDNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4180 }
  _classInfo112.ExpListNode = {
    name='ExpListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3734 }
  _classInfo112.ExpMacroExpNode = {
    name='ExpMacroExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4204 }
  _classInfo112.ExpMacroStatNode = {
    name='ExpMacroStatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4220 }
  _classInfo112.ExpNewNode = {
    name='ExpNewNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4038 }
  _classInfo112.ExpOp1Node = {
    name='ExpOp1Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4138 }
  _classInfo112.ExpOp2Node = {
    name='ExpOp2Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4078 }
  _classInfo112.ExpParenNode = {
    name='ExpParenNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4192 }
  _classInfo112.ExpRefItemNode = {
    name='ExpRefItemNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4152 }
  _classInfo112.ExpRefNode = {
    name='ExpRefNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4066 }
  _classInfo112.ExpUnwrapNode = {
    name='ExpUnwrapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4052 }
  _classInfo112.Filter = {
    name='Filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3706 }
  _classInfo112.ForNode = {
    name='ForNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3938 }
  _classInfo112.ForeachNode = {
    name='ForeachNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3978 }
  _classInfo112.ForsortNode = {
    name='ForsortNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3996 }
  _classInfo112.GetFieldNode = {
    name='GetFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4262 }
  _classInfo112.IfNode = {
    name='IfNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3852 }
  _classInfo112.IfStmtInfo = {
    name='IfStmtInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3842 }
  _classInfo112.IfUnwrapNode = {
    name='IfUnwrapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4110 }
  _classInfo112.ImportNode = {
    name='ImportNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3772 }
  _classInfo112.LiteralArrayNode = {
    name='LiteralArrayNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4534 }
  _classInfo112.LiteralBoolNode = {
    name='LiteralBoolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4606 }
  _classInfo112.LiteralCharNode = {
    name='LiteralCharNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4492 }
  _classInfo112.LiteralIntNode = {
    name='LiteralIntNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4506 }
  _classInfo112.LiteralListNode = {
    name='LiteralListNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4546 }
  _classInfo112.LiteralMapNode = {
    name='LiteralMapNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4566 }
  _classInfo112.LiteralNilNode = {
    name='LiteralNilNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4482 }
  _classInfo112.LiteralRealNode = {
    name='LiteralRealNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4520 }
  _classInfo112.LiteralStringNode = {
    name='LiteralStringNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4588 }
  _classInfo112.LiteralSymbolNode = {
    name='LiteralSymbolNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4618 }
  _classInfo112.MacroEval = {
    name='MacroEval', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3732 }
  _classInfo112.NamespaceInfo = {
    name='NamespaceInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3728 }
  _classInfo112.Node = {
    name='Node', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3596 }
  _classInfo112.NoneNode = {
    name='NoneNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3762 }
  _classInfo112.NormalTypeInfo = {
    name='NormalTypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3600 }
  _classInfo112.OutStream = {
    name='OutStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3532 }
  _classInfo112.PairItem = {
    name='PairItem', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4558 }
  _classInfo112.RefFieldNode = {
    name='RefFieldNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4248 }
  _classInfo112.RefTypeNode = {
    name='RefTypeNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3806 }
  _classInfo112.RepeatNode = {
    name='RepeatNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3924 }
  _classInfo112.ReturnNode = {
    name='ReturnNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4016 }
  _classInfo112.RootNode = {
    name='RootNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3784 }
  _classInfo112.Scope = {
    name='Scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3682 }
  _classInfo112.StmtExpNode = {
    name='StmtExpNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4236 }
  _classInfo112.SwitchNode = {
    name='SwitchNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3890 }
  _classInfo112.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3752 }
  _classInfo112.TypeInfo = {
    name='TypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3526 }
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
  _classInfo112.UnwrapSetNode = {
    name='UnwrapSetNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4094 }
  _classInfo112.VarInfo = {
    name='VarInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4278 }
  _classInfo112.WhileNode = {
    name='WhileNode', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3910 }
  _classInfo112.getNodeKindName = {
    name='getNodeKindName', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3760 }
  _classInfo112.isBuiltin = {
    name='isBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3530 }
  _classInfo112.lune = {
    name='lune', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3436 }
  _classInfo112.nodeKind = {
    name='nodeKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3758 }
  _classInfo112.rootTypeId = {
    name='rootTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 14 }
  _classInfo112.typeInfoKind = {
    name='typeInfoKind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3528 }
  end
do
  local _classInfo3448 = {}
  _className2InfoMap.TxtStream = _classInfo3448
  _typeId2ClassInfoMap[ 3448 ] = _classInfo3448
  end
do
  local _classInfo3526 = {}
  _className2InfoMap.TypeInfo = _classInfo3526
  _typeId2ClassInfoMap[ 3526 ] = _classInfo3526
  end
do
  local _classInfo4094 = {}
  _className2InfoMap.UnwrapSetNode = _classInfo4094
  _typeId2ClassInfoMap[ 4094 ] = _classInfo4094
  end
do
  local _classInfo3504 = {}
  _className2InfoMap.Util = _classInfo3504
  _typeId2ClassInfoMap[ 3504 ] = _classInfo3504
  _classInfo3504.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3522 }
  _classInfo3504.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3520 }
  _classInfo3504.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3512 }
  _classInfo3504.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 3506 }
  _classInfo3504.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 3524 }
  end
do
  local _classInfo4278 = {}
  _className2InfoMap.VarInfo = _classInfo4278
  _typeId2ClassInfoMap[ 4278 ] = _classInfo4278
  end
do
  local _classInfo3910 = {}
  _className2InfoMap.WhileNode = _classInfo3910
  _typeId2ClassInfoMap[ 3910 ] = _classInfo3910
  end
do
  local _classInfo3478 = {}
  _className2InfoMap.WrapParser = _classInfo3478
  _typeId2ClassInfoMap[ 3478 ] = _classInfo3478
  end
do
  local _classInfo110 = {}
  _className2InfoMap.base = _classInfo110
  _typeId2ClassInfoMap[ 110 ] = _classInfo110
  _classInfo110.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4716 }
  _classInfo110.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 112 }
  _classInfo110.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4782 }
  end
do
  local _classInfo106 = {}
  _className2InfoMap.convLua = _classInfo106
  _typeId2ClassInfoMap[ 106 ] = _classInfo106
  _classInfo106.MacroEvalImp = {
    name='MacroEvalImp', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 5010 }
  _classInfo106.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4716 }
  _classInfo106.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 112 }
  _classInfo106.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4782 }
  _classInfo106.convFilter = {
    name='convFilter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4812 }
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
  local _classInfo3512 = {}
  _className2InfoMap.memStream = _classInfo3512
  _typeId2ClassInfoMap[ 3512 ] = _classInfo3512
  end
do
  local _classInfo3506 = {}
  _className2InfoMap.outStream = _classInfo3506
  _typeId2ClassInfoMap[ 3506 ] = _classInfo3506
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 7, nilable = true, orgTypeId = 6 }_typeInfoList[2] = { parentId = 1, typeId = 11, nilable = true, orgTypeId = 10 }_typeInfoList[3] = { parentId = 1, typeId = 13, nilable = true, orgTypeId = 12 }_typeInfoList[4] = { parentId = 1, typeId = 15, nilable = true, orgTypeId = 14 }_typeInfoList[5] = { parentId = 1, typeId = 17, nilable = true, orgTypeId = 16 }_typeInfoList[6] = { parentId = 1, typeId = 21, nilable = true, orgTypeId = 20 }_typeInfoList[7] = { parentId = 1, typeId = 102, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[8] = { parentId = 1, typeId = 103, nilable = true, orgTypeId = 102 }_typeInfoList[9] = { parentId = 102, typeId = 104, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106}, }
_typeInfoList[10] = { parentId = 102, typeId = 105, nilable = true, orgTypeId = 104 }_typeInfoList[11] = { parentId = 104, typeId = 106, baseId = 1, txt = 'convLua',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {108, 4812, 5010}, }
_typeInfoList[12] = { parentId = 104, typeId = 107, nilable = true, orgTypeId = 106 }_typeInfoList[13] = { parentId = 106, typeId = 108, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {110}, }
_typeInfoList[14] = { parentId = 106, typeId = 109, nilable = true, orgTypeId = 108 }_typeInfoList[15] = { parentId = 108, typeId = 110, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {112, 4716, 4782}, }
_typeInfoList[16] = { parentId = 108, typeId = 111, nilable = true, orgTypeId = 110 }_typeInfoList[17] = { parentId = 110, typeId = 112, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3436, 3526, 3530, 3532, 3596, 3598, 3600, 3682, 3706, 3728, 3732, 3734, 3736, 3752, 3760, 3762, 3772, 3784, 3806, 3824, 3842, 3852, 3882, 3890, 3910, 3924, 3938, 3958, 3978, 3996, 4016, 4028, 4038, 4052, 4066, 4078, 4094, 4110, 4126, 4138, 4152, 4166, 4180, 4192, 4204, 4220, 4236, 4248, 4262, 4278, 4288, 4320, 4342, 4354, 4366, 4378, 4392, 4414, 4428, 4470, 4482, 4492, 4506, 4520, 4534, 4546, 4558, 4566, 4588, 4606, 4618, 4706}, }
_typeInfoList[18] = { parentId = 110, typeId = 113, nilable = true, orgTypeId = 112 }_typeInfoList[19] = { parentId = 112, typeId = 3436, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3438}, }
_typeInfoList[20] = { parentId = 112, typeId = 3437, nilable = true, orgTypeId = 3436 }_typeInfoList[21] = { parentId = 3436, typeId = 3438, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3440, 3504}, }
_typeInfoList[22] = { parentId = 3436, typeId = 3439, nilable = true, orgTypeId = 3438 }_typeInfoList[23] = { parentId = 3438, typeId = 3440, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3442, 3448, 3454, 3458, 3470, 3478, 3486, 3494, 3496, 3498, 3502}, }
_typeInfoList[24] = { parentId = 3438, typeId = 3441, nilable = true, orgTypeId = 3440 }_typeInfoList[25] = { parentId = 3440, typeId = 3442, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3444, 3446}, }
_typeInfoList[26] = { parentId = 3440, typeId = 3443, nilable = true, orgTypeId = 3442 }_typeInfoList[27] = { parentId = 3442, typeId = 3444, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[28] = { parentId = 3442, typeId = 3446, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[29] = { parentId = 3440, typeId = 3448, baseId = 3442, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3450, 3452}, }
_typeInfoList[30] = { parentId = 3440, typeId = 3449, nilable = true, orgTypeId = 3448 }_typeInfoList[31] = { parentId = 3448, typeId = 3450, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 3448, typeId = 3452, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[33] = { parentId = 3440, typeId = 3454, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3456}, }
_typeInfoList[34] = { parentId = 3440, typeId = 3455, nilable = true, orgTypeId = 3454 }_typeInfoList[35] = { parentId = 3454, typeId = 3456, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 14}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 3440, typeId = 3458, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3462, 3466, 3468}, }
_typeInfoList[37] = { parentId = 3440, typeId = 3459, nilable = true, orgTypeId = 3458 }_typeInfoList[38] = { parentId = 1, typeId = 3460, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3458}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 1, typeId = 3461, nilable = true, orgTypeId = 3460 }_typeInfoList[40] = { parentId = 3458, typeId = 3462, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3460}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 1, typeId = 3464, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3458}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 1, typeId = 3465, nilable = true, orgTypeId = 3464 }_typeInfoList[43] = { parentId = 3458, typeId = 3466, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3464}, children = {}, }
_typeInfoList[44] = { parentId = 3458, typeId = 3468, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 20, 3454}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 3440, typeId = 3470, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3472, 3474, 3476}, }
_typeInfoList[46] = { parentId = 3440, typeId = 3471, nilable = true, orgTypeId = 3470 }_typeInfoList[47] = { parentId = 3470, typeId = 3472, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[48] = { parentId = 3470, typeId = 3474, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[49] = { parentId = 3470, typeId = 3476, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[50] = { parentId = 3440, typeId = 3478, baseId = 3470, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3480, 3482, 3484}, }
_typeInfoList[51] = { parentId = 3440, typeId = 3479, nilable = true, orgTypeId = 3478 }_typeInfoList[52] = { parentId = 3478, typeId = 3480, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[53] = { parentId = 3478, typeId = 3482, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[54] = { parentId = 3478, typeId = 3484, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3470, 20}, retTypeId = {}, children = {}, }
_typeInfoList[55] = { parentId = 3440, typeId = 3486, baseId = 3470, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3488, 3490, 3492, 3500}, }
_typeInfoList[56] = { parentId = 3440, typeId = 3487, nilable = true, orgTypeId = 3486 }_typeInfoList[57] = { parentId = 3486, typeId = 3488, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3442, 20, 12}, retTypeId = {}, children = {}, }
_typeInfoList[58] = { parentId = 3486, typeId = 3490, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[59] = { parentId = 3486, typeId = 3492, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {3487}, children = {}, }
_typeInfoList[60] = { parentId = 3440, typeId = 3494, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[61] = { parentId = 3440, typeId = 3496, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[62] = { parentId = 3440, typeId = 3498, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[63] = { parentId = 3486, typeId = 3500, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3459}, children = {}, }
_typeInfoList[64] = { parentId = 3440, typeId = 3502, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[65] = { parentId = 3438, typeId = 3504, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3506, 3512, 3520, 3522, 3524}, }
_typeInfoList[66] = { parentId = 3438, typeId = 3505, nilable = true, orgTypeId = 3504 }_typeInfoList[67] = { parentId = 3504, typeId = 3506, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3508, 3510}, }
_typeInfoList[68] = { parentId = 3504, typeId = 3507, nilable = true, orgTypeId = 3506 }_typeInfoList[69] = { parentId = 3506, typeId = 3508, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[70] = { parentId = 3506, typeId = 3510, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[71] = { parentId = 3504, typeId = 3512, baseId = 3506, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3514, 3516, 3518}, }
_typeInfoList[72] = { parentId = 3504, typeId = 3513, nilable = true, orgTypeId = 3512 }_typeInfoList[73] = { parentId = 3512, typeId = 3514, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[74] = { parentId = 3512, typeId = 3516, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[75] = { parentId = 3512, typeId = 3518, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[76] = { parentId = 3504, typeId = 3520, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 3504, typeId = 3522, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 3504, typeId = 3524, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 6, 20}, retTypeId = {}, children = {}, }
_typeInfoList[79] = { parentId = 112, typeId = 3526, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3538, 3540, 3542, 3544, 3546, 3548, 3550, 3554, 3558, 3562, 3564, 3566, 3568, 3570, 3572, 3574, 3576, 3578, 3580, 3582, 3584, 3588, 3592, 3594}, }
_typeInfoList[80] = { parentId = 112, typeId = 3527, nilable = true, orgTypeId = 3526 }_typeInfoList[81] = { parentId = 112, typeId = 3530, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {6}, children = {}, }
_typeInfoList[82] = { parentId = 112, typeId = 3532, baseId = 1, txt = 'OutStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3534, 3536}, }
_typeInfoList[83] = { parentId = 112, typeId = 3533, nilable = true, orgTypeId = 3532 }_typeInfoList[84] = { parentId = 3532, typeId = 3534, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[85] = { parentId = 3532, typeId = 3536, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[86] = { parentId = 3526, typeId = 3538, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[87] = { parentId = 3526, typeId = 3540, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[88] = { parentId = 3526, typeId = 3542, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3526}, retTypeId = {12}, children = {}, }
_typeInfoList[89] = { parentId = 3526, typeId = 3544, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[90] = { parentId = 3526, typeId = 3546, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3532}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 3526, typeId = 3548, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3526, 14}, retTypeId = {12}, children = {}, }
_typeInfoList[92] = { parentId = 3526, typeId = 3550, baseId = 1, txt = 'get_externalFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 3552, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 1, typeId = 3553, nilable = true, orgTypeId = 3552 }_typeInfoList[95] = { parentId = 3526, typeId = 3554, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3552}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 3556, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[97] = { parentId = 1, typeId = 3557, nilable = true, orgTypeId = 3556 }_typeInfoList[98] = { parentId = 3526, typeId = 3558, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3556}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 3560, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[100] = { parentId = 1, typeId = 3561, nilable = true, orgTypeId = 3560 }_typeInfoList[101] = { parentId = 3526, typeId = 3562, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3560}, children = {}, }
_typeInfoList[102] = { parentId = 3526, typeId = 3564, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[103] = { parentId = 3526, typeId = 3566, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[104] = { parentId = 3526, typeId = 3568, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[105] = { parentId = 3526, typeId = 3570, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[106] = { parentId = 3526, typeId = 3572, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[107] = { parentId = 3526, typeId = 3574, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[108] = { parentId = 3526, typeId = 3576, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[109] = { parentId = 3526, typeId = 3578, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[110] = { parentId = 3526, typeId = 3580, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[111] = { parentId = 3526, typeId = 3582, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[112] = { parentId = 3526, typeId = 3584, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[113] = { parentId = 1, typeId = 3586, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 3587, nilable = true, orgTypeId = 3586 }_typeInfoList[115] = { parentId = 3526, typeId = 3588, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3586}, children = {}, }
_typeInfoList[116] = { parentId = 1, typeId = 3590, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 1, typeId = 3591, nilable = true, orgTypeId = 3590 }_typeInfoList[118] = { parentId = 3526, typeId = 3592, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3590}, children = {}, }
_typeInfoList[119] = { parentId = 3526, typeId = 3594, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 112, typeId = 3596, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3710, 3716, 3718, 3720, 3724, 3726}, }
_typeInfoList[121] = { parentId = 112, typeId = 3597, nilable = true, orgTypeId = 3596 }_typeInfoList[122] = { parentId = 112, typeId = 3598, baseId = 3596, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4440, 4450, 4452, 4454, 4458, 4462, 4464, 4468}, }
_typeInfoList[123] = { parentId = 112, typeId = 3599, nilable = true, orgTypeId = 3598 }_typeInfoList[124] = { parentId = 112, typeId = 3600, baseId = 3526, txt = 'NormalTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3602, 3604, 3606, 3608, 3610, 3612, 3620, 3624, 3628, 3632, 3634, 3636, 3638, 3640, 3642, 3644, 3646, 3648, 3650, 3652, 3654, 3658, 3660, 3664, 3668, 3670, 3672, 3678, 3680}, }
_typeInfoList[125] = { parentId = 112, typeId = 3601, nilable = true, orgTypeId = 3600 }_typeInfoList[126] = { parentId = 3600, typeId = 3602, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[127] = { parentId = 3600, typeId = 3604, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[128] = { parentId = 3600, typeId = 3606, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[129] = { parentId = 3600, typeId = 3608, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3532}, retTypeId = {}, children = {}, }
_typeInfoList[130] = { parentId = 3600, typeId = 3610, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3526, 14}, retTypeId = {12}, children = {}, }
_typeInfoList[131] = { parentId = 3600, typeId = 3612, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3526}, retTypeId = {3600}, children = {}, }
_typeInfoList[132] = { parentId = 1, typeId = 3614, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 1, typeId = 3615, nilable = true, orgTypeId = 3614 }_typeInfoList[134] = { parentId = 1, typeId = 3616, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 1, typeId = 3617, nilable = true, orgTypeId = 3616 }_typeInfoList[136] = { parentId = 1, typeId = 3618, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[137] = { parentId = 1, typeId = 3619, nilable = true, orgTypeId = 3618 }_typeInfoList[138] = { parentId = 3600, typeId = 3620, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3526, 3526, 12, 14, 20, 3614, 3616, 3618}, retTypeId = {3526}, children = {}, }
_typeInfoList[139] = { parentId = 1, typeId = 3622, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1, typeId = 3623, nilable = true, orgTypeId = 3622 }_typeInfoList[141] = { parentId = 3600, typeId = 3624, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3622}, children = {}, }
_typeInfoList[142] = { parentId = 1, typeId = 3626, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 3627, nilable = true, orgTypeId = 3626 }_typeInfoList[144] = { parentId = 3600, typeId = 3628, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3626}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 3630, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[146] = { parentId = 1, typeId = 3631, nilable = true, orgTypeId = 3630 }_typeInfoList[147] = { parentId = 3600, typeId = 3632, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3630}, children = {}, }
_typeInfoList[148] = { parentId = 3600, typeId = 3634, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[149] = { parentId = 3600, typeId = 3636, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[150] = { parentId = 3600, typeId = 3638, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[151] = { parentId = 3600, typeId = 3640, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[152] = { parentId = 3600, typeId = 3642, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[153] = { parentId = 3600, typeId = 3644, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[154] = { parentId = 3600, typeId = 3646, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[155] = { parentId = 3600, typeId = 3648, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[156] = { parentId = 3600, typeId = 3650, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[157] = { parentId = 3600, typeId = 3652, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[158] = { parentId = 3600, typeId = 3654, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[159] = { parentId = 1, typeId = 3656, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 1, typeId = 3657, nilable = true, orgTypeId = 3656 }_typeInfoList[161] = { parentId = 3600, typeId = 3658, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3656}, children = {}, }
_typeInfoList[162] = { parentId = 3600, typeId = 3660, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 20, 14, 3526}, retTypeId = {3526}, children = {}, }
_typeInfoList[163] = { parentId = 1, typeId = 3662, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 1, typeId = 3663, nilable = true, orgTypeId = 3662 }_typeInfoList[165] = { parentId = 3600, typeId = 3664, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3526, 3662}, retTypeId = {3526}, children = {}, }
_typeInfoList[166] = { parentId = 1, typeId = 3666, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[167] = { parentId = 1, typeId = 3667, nilable = true, orgTypeId = 3666 }_typeInfoList[168] = { parentId = 3600, typeId = 3668, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3526, 3666}, retTypeId = {3526}, children = {}, }
_typeInfoList[169] = { parentId = 3600, typeId = 3670, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 3526, 3526, 3526}, retTypeId = {3526}, children = {}, }
_typeInfoList[170] = { parentId = 3600, typeId = 3672, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {3526, 3526, 12, 20, 20}, retTypeId = {3526}, children = {}, }
_typeInfoList[171] = { parentId = 1, typeId = 3674, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[172] = { parentId = 1, typeId = 3675, nilable = true, orgTypeId = 3674 }_typeInfoList[173] = { parentId = 1, typeId = 3676, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[174] = { parentId = 1, typeId = 3677, nilable = true, orgTypeId = 3676 }_typeInfoList[175] = { parentId = 3600, typeId = 3678, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14, 3526, 12, 12, 12, 20, 20, 3674, 3676}, retTypeId = {3526}, children = {}, }
_typeInfoList[176] = { parentId = 3600, typeId = 3680, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3526}, retTypeId = {12}, children = {}, }
_typeInfoList[177] = { parentId = 112, typeId = 3682, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3684, 3686, 3688, 3690, 3692, 3694, 3696, 3700, 3704}, }
_typeInfoList[178] = { parentId = 112, typeId = 3683, nilable = true, orgTypeId = 3682 }_typeInfoList[179] = { parentId = 3682, typeId = 3684, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3526}, retTypeId = {}, children = {}, }
_typeInfoList[180] = { parentId = 3682, typeId = 3686, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3526, 3682}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 3682, typeId = 3688, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3683}, children = {}, }
_typeInfoList[182] = { parentId = 3682, typeId = 3690, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3527}, children = {}, }
_typeInfoList[183] = { parentId = 3682, typeId = 3692, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {3527}, children = {}, }
_typeInfoList[184] = { parentId = 3682, typeId = 3694, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {3527}, children = {}, }
_typeInfoList[185] = { parentId = 3682, typeId = 3696, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3682}, children = {}, }
_typeInfoList[186] = { parentId = 1, typeId = 3698, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 1, typeId = 3699, nilable = true, orgTypeId = 3698 }_typeInfoList[188] = { parentId = 3682, typeId = 3700, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3698}, children = {}, }
_typeInfoList[189] = { parentId = 1, typeId = 3702, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 3682}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[190] = { parentId = 1, typeId = 3703, nilable = true, orgTypeId = 3702 }_typeInfoList[191] = { parentId = 3682, typeId = 3704, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3702}, children = {}, }
_typeInfoList[192] = { parentId = 112, typeId = 3706, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3708, 3764, 3774, 3786, 3808, 3826, 3854, 3868, 3892, 3912, 3926, 3940, 3960, 3980, 3998, 4018, 4030, 4040, 4054, 4068, 4080, 4096, 4112, 4128, 4140, 4154, 4168, 4182, 4194, 4206, 4222, 4238, 4250, 4264, 4290, 4344, 4356, 4368, 4380, 4394, 4416, 4430, 4438, 4472, 4484, 4494, 4508, 4522, 4536, 4548, 4568, 4590, 4608, 4620}, }
_typeInfoList[193] = { parentId = 112, typeId = 3707, nilable = true, orgTypeId = 3706 }_typeInfoList[194] = { parentId = 3706, typeId = 3708, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 3596, typeId = 3710, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[196] = { parentId = 1, typeId = 3712, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[197] = { parentId = 1, typeId = 3713, nilable = true, orgTypeId = 3712 }_typeInfoList[198] = { parentId = 1, typeId = 3714, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 1, typeId = 3715, nilable = true, orgTypeId = 3714 }_typeInfoList[200] = { parentId = 3596, typeId = 3716, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3712, 3714}, children = {}, }
_typeInfoList[201] = { parentId = 3596, typeId = 3718, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[202] = { parentId = 3596, typeId = 3720, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[203] = { parentId = 1, typeId = 3722, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[204] = { parentId = 1, typeId = 3723, nilable = true, orgTypeId = 3722 }_typeInfoList[205] = { parentId = 3596, typeId = 3724, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3722}, children = {}, }
_typeInfoList[206] = { parentId = 3596, typeId = 3726, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 3454}, retTypeId = {}, children = {}, }
_typeInfoList[207] = { parentId = 112, typeId = 3728, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3730}, }
_typeInfoList[208] = { parentId = 112, typeId = 3729, nilable = true, orgTypeId = 3728 }_typeInfoList[209] = { parentId = 3728, typeId = 3730, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3682, 3526}, retTypeId = {}, children = {}, }
_typeInfoList[210] = { parentId = 112, typeId = 3732, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4702, 4704}, }
_typeInfoList[211] = { parentId = 112, typeId = 3733, nilable = true, orgTypeId = 3732 }_typeInfoList[212] = { parentId = 112, typeId = 3734, baseId = 3596, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3870, 3876, 3880}, }
_typeInfoList[213] = { parentId = 112, typeId = 3735, nilable = true, orgTypeId = 3734 }_typeInfoList[214] = { parentId = 112, typeId = 3736, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3738, 3742, 3744, 3748, 3750}, }
_typeInfoList[215] = { parentId = 112, typeId = 3737, nilable = true, orgTypeId = 3736 }_typeInfoList[216] = { parentId = 3736, typeId = 3738, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[217] = { parentId = 1, typeId = 3740, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[218] = { parentId = 1, typeId = 3741, nilable = true, orgTypeId = 3740 }_typeInfoList[219] = { parentId = 3736, typeId = 3742, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3740}, children = {}, }
_typeInfoList[220] = { parentId = 3736, typeId = 3744, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[221] = { parentId = 1, typeId = 3746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3458}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[222] = { parentId = 1, typeId = 3747, nilable = true, orgTypeId = 3746 }_typeInfoList[223] = { parentId = 3736, typeId = 3748, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3746}, children = {}, }
_typeInfoList[224] = { parentId = 3736, typeId = 3750, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3458, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 112, typeId = 3752, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3756, 4714}, }
_typeInfoList[226] = { parentId = 112, typeId = 3753, nilable = true, orgTypeId = 3752 }_typeInfoList[227] = { parentId = 1, typeId = 3754, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {20}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 1, typeId = 3755, nilable = true, orgTypeId = 3754 }_typeInfoList[229] = { parentId = 3752, typeId = 3756, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3754}, children = {}, }
_typeInfoList[230] = { parentId = 112, typeId = 3760, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[231] = { parentId = 112, typeId = 3762, baseId = 3596, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3766, 3770}, }
_typeInfoList[232] = { parentId = 112, typeId = 3763, nilable = true, orgTypeId = 3762 }_typeInfoList[233] = { parentId = 3706, typeId = 3764, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3762, 10}, retTypeId = {}, children = {}, }
_typeInfoList[234] = { parentId = 3762, typeId = 3766, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[235] = { parentId = 1, typeId = 3768, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[236] = { parentId = 1, typeId = 3769, nilable = true, orgTypeId = 3768 }_typeInfoList[237] = { parentId = 3762, typeId = 3770, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3768}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 112, typeId = 3772, baseId = 3596, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3776, 3780, 3782}, }
_typeInfoList[239] = { parentId = 112, typeId = 3773, nilable = true, orgTypeId = 3772 }_typeInfoList[240] = { parentId = 3706, typeId = 3774, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3772, 10}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 3772, typeId = 3776, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 1, typeId = 3778, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[243] = { parentId = 1, typeId = 3779, nilable = true, orgTypeId = 3778 }_typeInfoList[244] = { parentId = 3772, typeId = 3780, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3778, 20}, retTypeId = {}, children = {}, }
_typeInfoList[245] = { parentId = 3772, typeId = 3782, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[246] = { parentId = 112, typeId = 3784, baseId = 3596, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3788, 3796, 3800, 3804}, }
_typeInfoList[247] = { parentId = 112, typeId = 3785, nilable = true, orgTypeId = 3784 }_typeInfoList[248] = { parentId = 3706, typeId = 3786, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3784, 10}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 3784, typeId = 3788, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[250] = { parentId = 1, typeId = 3790, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 1, typeId = 3791, nilable = true, orgTypeId = 3790 }_typeInfoList[252] = { parentId = 1, typeId = 3792, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[253] = { parentId = 1, typeId = 3793, nilable = true, orgTypeId = 3792 }_typeInfoList[254] = { parentId = 1, typeId = 3794, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {14, 3728}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[255] = { parentId = 1, typeId = 3795, nilable = true, orgTypeId = 3794 }_typeInfoList[256] = { parentId = 3784, typeId = 3796, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3790, 3792, 3794}, retTypeId = {}, children = {}, }
_typeInfoList[257] = { parentId = 1, typeId = 3798, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[258] = { parentId = 1, typeId = 3799, nilable = true, orgTypeId = 3798 }_typeInfoList[259] = { parentId = 3784, typeId = 3800, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3798}, children = {}, }
_typeInfoList[260] = { parentId = 1, typeId = 3802, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {14, 3728}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[261] = { parentId = 1, typeId = 3803, nilable = true, orgTypeId = 3802 }_typeInfoList[262] = { parentId = 3784, typeId = 3804, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3802}, children = {}, }
_typeInfoList[263] = { parentId = 112, typeId = 3806, baseId = 3596, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3810, 3814, 3816, 3818, 3820, 3822}, }
_typeInfoList[264] = { parentId = 112, typeId = 3807, nilable = true, orgTypeId = 3806 }_typeInfoList[265] = { parentId = 3706, typeId = 3808, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3806, 10}, retTypeId = {}, children = {}, }
_typeInfoList[266] = { parentId = 3806, typeId = 3810, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[267] = { parentId = 1, typeId = 3812, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[268] = { parentId = 1, typeId = 3813, nilable = true, orgTypeId = 3812 }_typeInfoList[269] = { parentId = 3806, typeId = 3814, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3812, 3596, 12, 12, 20}, retTypeId = {}, children = {}, }
_typeInfoList[270] = { parentId = 3806, typeId = 3816, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[271] = { parentId = 3806, typeId = 3818, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[272] = { parentId = 3806, typeId = 3820, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[273] = { parentId = 3806, typeId = 3822, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[274] = { parentId = 112, typeId = 3824, baseId = 3596, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3828, 3834, 3836, 3840}, }
_typeInfoList[275] = { parentId = 112, typeId = 3825, nilable = true, orgTypeId = 3824 }_typeInfoList[276] = { parentId = 3706, typeId = 3826, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3824, 10}, retTypeId = {}, children = {}, }
_typeInfoList[277] = { parentId = 3824, typeId = 3828, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[278] = { parentId = 1, typeId = 3830, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[279] = { parentId = 1, typeId = 3831, nilable = true, orgTypeId = 3830 }_typeInfoList[280] = { parentId = 1, typeId = 3832, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[281] = { parentId = 1, typeId = 3833, nilable = true, orgTypeId = 3832 }_typeInfoList[282] = { parentId = 3824, typeId = 3834, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3830, 20, 3832}, retTypeId = {}, children = {}, }
_typeInfoList[283] = { parentId = 3824, typeId = 3836, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[284] = { parentId = 1, typeId = 3838, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[285] = { parentId = 1, typeId = 3839, nilable = true, orgTypeId = 3838 }_typeInfoList[286] = { parentId = 3824, typeId = 3840, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3838}, children = {}, }
_typeInfoList[287] = { parentId = 112, typeId = 3842, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3844, 3846, 3848, 3850}, }
_typeInfoList[288] = { parentId = 112, typeId = 3843, nilable = true, orgTypeId = 3842 }_typeInfoList[289] = { parentId = 3842, typeId = 3844, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[290] = { parentId = 3842, typeId = 3846, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[291] = { parentId = 3842, typeId = 3848, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[292] = { parentId = 3842, typeId = 3850, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20, 3596, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[293] = { parentId = 112, typeId = 3852, baseId = 3596, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3856, 3862, 3866}, }
_typeInfoList[294] = { parentId = 112, typeId = 3853, nilable = true, orgTypeId = 3852 }_typeInfoList[295] = { parentId = 3706, typeId = 3854, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3852, 10}, retTypeId = {}, children = {}, }
_typeInfoList[296] = { parentId = 3852, typeId = 3856, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[297] = { parentId = 1, typeId = 3858, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[298] = { parentId = 1, typeId = 3859, nilable = true, orgTypeId = 3858 }_typeInfoList[299] = { parentId = 1, typeId = 3860, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[300] = { parentId = 1, typeId = 3861, nilable = true, orgTypeId = 3860 }_typeInfoList[301] = { parentId = 3852, typeId = 3862, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3858, 3860}, retTypeId = {}, children = {}, }
_typeInfoList[302] = { parentId = 1, typeId = 3864, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[303] = { parentId = 1, typeId = 3865, nilable = true, orgTypeId = 3864 }_typeInfoList[304] = { parentId = 3852, typeId = 3866, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3864}, children = {}, }
_typeInfoList[305] = { parentId = 3706, typeId = 3868, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3734, 10}, retTypeId = {}, children = {}, }
_typeInfoList[306] = { parentId = 3734, typeId = 3870, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[307] = { parentId = 1, typeId = 3872, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[308] = { parentId = 1, typeId = 3873, nilable = true, orgTypeId = 3872 }_typeInfoList[309] = { parentId = 1, typeId = 3874, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[310] = { parentId = 1, typeId = 3875, nilable = true, orgTypeId = 3874 }_typeInfoList[311] = { parentId = 3734, typeId = 3876, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3872, 3874}, retTypeId = {}, children = {}, }
_typeInfoList[312] = { parentId = 1, typeId = 3878, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[313] = { parentId = 1, typeId = 3879, nilable = true, orgTypeId = 3878 }_typeInfoList[314] = { parentId = 3734, typeId = 3880, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3878}, children = {}, }
_typeInfoList[315] = { parentId = 112, typeId = 3882, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3884, 3886, 3888}, }
_typeInfoList[316] = { parentId = 112, typeId = 3883, nilable = true, orgTypeId = 3882 }_typeInfoList[317] = { parentId = 3882, typeId = 3884, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[318] = { parentId = 3882, typeId = 3886, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[319] = { parentId = 3882, typeId = 3888, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3734, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[320] = { parentId = 112, typeId = 3890, baseId = 3596, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3894, 3900, 3902, 3906, 3908}, }
_typeInfoList[321] = { parentId = 112, typeId = 3891, nilable = true, orgTypeId = 3890 }_typeInfoList[322] = { parentId = 3706, typeId = 3892, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3890, 10}, retTypeId = {}, children = {}, }
_typeInfoList[323] = { parentId = 3890, typeId = 3894, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[324] = { parentId = 1, typeId = 3896, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[325] = { parentId = 1, typeId = 3897, nilable = true, orgTypeId = 3896 }_typeInfoList[326] = { parentId = 1, typeId = 3898, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3882}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[327] = { parentId = 1, typeId = 3899, nilable = true, orgTypeId = 3898 }_typeInfoList[328] = { parentId = 3890, typeId = 3900, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3896, 3596, 3898, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[329] = { parentId = 3890, typeId = 3902, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[330] = { parentId = 1, typeId = 3904, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3882}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[331] = { parentId = 1, typeId = 3905, nilable = true, orgTypeId = 3904 }_typeInfoList[332] = { parentId = 3890, typeId = 3906, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3904}, children = {}, }
_typeInfoList[333] = { parentId = 3890, typeId = 3908, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[334] = { parentId = 112, typeId = 3910, baseId = 3596, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3914, 3918, 3920, 3922}, }
_typeInfoList[335] = { parentId = 112, typeId = 3911, nilable = true, orgTypeId = 3910 }_typeInfoList[336] = { parentId = 3706, typeId = 3912, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3910, 10}, retTypeId = {}, children = {}, }
_typeInfoList[337] = { parentId = 3910, typeId = 3914, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 1, typeId = 3916, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[339] = { parentId = 1, typeId = 3917, nilable = true, orgTypeId = 3916 }_typeInfoList[340] = { parentId = 3910, typeId = 3918, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3916, 3596, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[341] = { parentId = 3910, typeId = 3920, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[342] = { parentId = 3910, typeId = 3922, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[343] = { parentId = 112, typeId = 3924, baseId = 3596, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3928, 3932, 3934, 3936}, }
_typeInfoList[344] = { parentId = 112, typeId = 3925, nilable = true, orgTypeId = 3924 }_typeInfoList[345] = { parentId = 3706, typeId = 3926, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3924, 10}, retTypeId = {}, children = {}, }
_typeInfoList[346] = { parentId = 3924, typeId = 3928, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[347] = { parentId = 1, typeId = 3930, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[348] = { parentId = 1, typeId = 3931, nilable = true, orgTypeId = 3930 }_typeInfoList[349] = { parentId = 3924, typeId = 3932, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3930, 3824, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[350] = { parentId = 3924, typeId = 3934, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[351] = { parentId = 3924, typeId = 3936, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[352] = { parentId = 112, typeId = 3938, baseId = 3596, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3942, 3946, 3948, 3950, 3952, 3954, 3956}, }
_typeInfoList[353] = { parentId = 112, typeId = 3939, nilable = true, orgTypeId = 3938 }_typeInfoList[354] = { parentId = 3706, typeId = 3940, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3938, 10}, retTypeId = {}, children = {}, }
_typeInfoList[355] = { parentId = 3938, typeId = 3942, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[356] = { parentId = 1, typeId = 3944, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[357] = { parentId = 1, typeId = 3945, nilable = true, orgTypeId = 3944 }_typeInfoList[358] = { parentId = 3938, typeId = 3946, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3944, 3824, 3458, 3596, 3596, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[359] = { parentId = 3938, typeId = 3948, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[360] = { parentId = 3938, typeId = 3950, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[361] = { parentId = 3938, typeId = 3952, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[362] = { parentId = 3938, typeId = 3954, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[363] = { parentId = 3938, typeId = 3956, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[364] = { parentId = 112, typeId = 3958, baseId = 3596, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3962, 3968, 3972, 3974, 3976}, }
_typeInfoList[365] = { parentId = 112, typeId = 3959, nilable = true, orgTypeId = 3958 }_typeInfoList[366] = { parentId = 3706, typeId = 3960, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3958, 10}, retTypeId = {}, children = {}, }
_typeInfoList[367] = { parentId = 3958, typeId = 3962, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[368] = { parentId = 1, typeId = 3964, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 1, typeId = 3965, nilable = true, orgTypeId = 3964 }_typeInfoList[370] = { parentId = 1, typeId = 3966, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3458}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[371] = { parentId = 1, typeId = 3967, nilable = true, orgTypeId = 3966 }_typeInfoList[372] = { parentId = 3958, typeId = 3968, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3964, 3966, 3596, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[373] = { parentId = 1, typeId = 3970, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3458}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[374] = { parentId = 1, typeId = 3971, nilable = true, orgTypeId = 3970 }_typeInfoList[375] = { parentId = 3958, typeId = 3972, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3970}, children = {}, }
_typeInfoList[376] = { parentId = 3958, typeId = 3974, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[377] = { parentId = 3958, typeId = 3976, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[378] = { parentId = 112, typeId = 3978, baseId = 3596, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3982, 3986, 3988, 3990, 3992, 3994}, }
_typeInfoList[379] = { parentId = 112, typeId = 3979, nilable = true, orgTypeId = 3978 }_typeInfoList[380] = { parentId = 3706, typeId = 3980, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3978, 10}, retTypeId = {}, children = {}, }
_typeInfoList[381] = { parentId = 3978, typeId = 3982, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[382] = { parentId = 1, typeId = 3984, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[383] = { parentId = 1, typeId = 3985, nilable = true, orgTypeId = 3984 }_typeInfoList[384] = { parentId = 3978, typeId = 3986, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 3984, 3458, 3458, 3596, 3824}, retTypeId = {}, children = {}, }
_typeInfoList[385] = { parentId = 3978, typeId = 3988, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[386] = { parentId = 3978, typeId = 3990, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[387] = { parentId = 3978, typeId = 3992, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[388] = { parentId = 3978, typeId = 3994, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[389] = { parentId = 112, typeId = 3996, baseId = 3596, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4000, 4004, 4006, 4008, 4010, 4012, 4014}, }
_typeInfoList[390] = { parentId = 112, typeId = 3997, nilable = true, orgTypeId = 3996 }_typeInfoList[391] = { parentId = 3706, typeId = 3998, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3996, 10}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 3996, typeId = 4000, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[393] = { parentId = 1, typeId = 4002, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[394] = { parentId = 1, typeId = 4003, nilable = true, orgTypeId = 4002 }_typeInfoList[395] = { parentId = 3996, typeId = 4004, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4002, 3458, 3458, 3596, 3824, 12}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 3996, typeId = 4006, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[397] = { parentId = 3996, typeId = 4008, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[398] = { parentId = 3996, typeId = 4010, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[399] = { parentId = 3996, typeId = 4012, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[400] = { parentId = 3996, typeId = 4014, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[401] = { parentId = 112, typeId = 4016, baseId = 3596, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4020, 4024, 4026}, }
_typeInfoList[402] = { parentId = 112, typeId = 4017, nilable = true, orgTypeId = 4016 }_typeInfoList[403] = { parentId = 3706, typeId = 4018, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4016, 10}, retTypeId = {}, children = {}, }
_typeInfoList[404] = { parentId = 4016, typeId = 4020, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[405] = { parentId = 1, typeId = 4022, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[406] = { parentId = 1, typeId = 4023, nilable = true, orgTypeId = 4022 }_typeInfoList[407] = { parentId = 4016, typeId = 4024, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4022, 3734}, retTypeId = {}, children = {}, }
_typeInfoList[408] = { parentId = 4016, typeId = 4026, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[409] = { parentId = 112, typeId = 4028, baseId = 3596, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4032, 4036}, }
_typeInfoList[410] = { parentId = 112, typeId = 4029, nilable = true, orgTypeId = 4028 }_typeInfoList[411] = { parentId = 3706, typeId = 4030, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4028, 10}, retTypeId = {}, children = {}, }
_typeInfoList[412] = { parentId = 4028, typeId = 4032, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[413] = { parentId = 1, typeId = 4034, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[414] = { parentId = 1, typeId = 4035, nilable = true, orgTypeId = 4034 }_typeInfoList[415] = { parentId = 4028, typeId = 4036, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4034}, retTypeId = {}, children = {}, }
_typeInfoList[416] = { parentId = 112, typeId = 4038, baseId = 3596, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4042, 4046, 4048, 4050}, }
_typeInfoList[417] = { parentId = 112, typeId = 4039, nilable = true, orgTypeId = 4038 }_typeInfoList[418] = { parentId = 3706, typeId = 4040, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4038, 10}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 4038, typeId = 4042, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[420] = { parentId = 1, typeId = 4044, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[421] = { parentId = 1, typeId = 4045, nilable = true, orgTypeId = 4044 }_typeInfoList[422] = { parentId = 4038, typeId = 4046, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4044, 3596, 3734}, retTypeId = {}, children = {}, }
_typeInfoList[423] = { parentId = 4038, typeId = 4048, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[424] = { parentId = 4038, typeId = 4050, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[425] = { parentId = 112, typeId = 4052, baseId = 3596, txt = 'ExpUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4056, 4060, 4062, 4064}, }
_typeInfoList[426] = { parentId = 112, typeId = 4053, nilable = true, orgTypeId = 4052 }_typeInfoList[427] = { parentId = 3706, typeId = 4054, baseId = 1, txt = 'processExpUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4052, 10}, retTypeId = {}, children = {}, }
_typeInfoList[428] = { parentId = 4052, typeId = 4056, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[429] = { parentId = 1, typeId = 4058, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[430] = { parentId = 1, typeId = 4059, nilable = true, orgTypeId = 4058 }_typeInfoList[431] = { parentId = 4052, typeId = 4060, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4058, 3596, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[432] = { parentId = 4052, typeId = 4062, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[433] = { parentId = 4052, typeId = 4064, baseId = 1, txt = 'get_instead',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[434] = { parentId = 112, typeId = 4066, baseId = 3596, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4070, 4074, 4076}, }
_typeInfoList[435] = { parentId = 112, typeId = 4067, nilable = true, orgTypeId = 4066 }_typeInfoList[436] = { parentId = 3706, typeId = 4068, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4066, 10}, retTypeId = {}, children = {}, }
_typeInfoList[437] = { parentId = 4066, typeId = 4070, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[438] = { parentId = 1, typeId = 4072, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[439] = { parentId = 1, typeId = 4073, nilable = true, orgTypeId = 4072 }_typeInfoList[440] = { parentId = 4066, typeId = 4074, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4072, 3458}, retTypeId = {}, children = {}, }
_typeInfoList[441] = { parentId = 4066, typeId = 4076, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[442] = { parentId = 112, typeId = 4078, baseId = 3596, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4082, 4086, 4088, 4090, 4092}, }
_typeInfoList[443] = { parentId = 112, typeId = 4079, nilable = true, orgTypeId = 4078 }_typeInfoList[444] = { parentId = 3706, typeId = 4080, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4078, 10}, retTypeId = {}, children = {}, }
_typeInfoList[445] = { parentId = 4078, typeId = 4082, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[446] = { parentId = 1, typeId = 4084, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[447] = { parentId = 1, typeId = 4085, nilable = true, orgTypeId = 4084 }_typeInfoList[448] = { parentId = 4078, typeId = 4086, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4084, 3458, 3596, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[449] = { parentId = 4078, typeId = 4088, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[450] = { parentId = 4078, typeId = 4090, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[451] = { parentId = 4078, typeId = 4092, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[452] = { parentId = 112, typeId = 4094, baseId = 3596, txt = 'UnwrapSetNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4098, 4102, 4104, 4106, 4108}, }
_typeInfoList[453] = { parentId = 112, typeId = 4095, nilable = true, orgTypeId = 4094 }_typeInfoList[454] = { parentId = 3706, typeId = 4096, baseId = 1, txt = 'processUnwrapSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4094, 10}, retTypeId = {}, children = {}, }
_typeInfoList[455] = { parentId = 4094, typeId = 4098, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[456] = { parentId = 1, typeId = 4100, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[457] = { parentId = 1, typeId = 4101, nilable = true, orgTypeId = 4100 }_typeInfoList[458] = { parentId = 4094, typeId = 4102, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4100, 3734, 3734, 3825}, retTypeId = {}, children = {}, }
_typeInfoList[459] = { parentId = 4094, typeId = 4104, baseId = 1, txt = 'get_dstExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[460] = { parentId = 4094, typeId = 4106, baseId = 1, txt = 'get_srcExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[461] = { parentId = 4094, typeId = 4108, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3825}, children = {}, }
_typeInfoList[462] = { parentId = 112, typeId = 4110, baseId = 3596, txt = 'IfUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4114, 4118, 4120, 4122, 4124}, }
_typeInfoList[463] = { parentId = 112, typeId = 4111, nilable = true, orgTypeId = 4110 }_typeInfoList[464] = { parentId = 3706, typeId = 4112, baseId = 1, txt = 'processIfUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4110, 10}, retTypeId = {}, children = {}, }
_typeInfoList[465] = { parentId = 4110, typeId = 4114, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[466] = { parentId = 1, typeId = 4116, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[467] = { parentId = 1, typeId = 4117, nilable = true, orgTypeId = 4116 }_typeInfoList[468] = { parentId = 4110, typeId = 4118, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4116, 3596, 3824, 3825}, retTypeId = {}, children = {}, }
_typeInfoList[469] = { parentId = 4110, typeId = 4120, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[470] = { parentId = 4110, typeId = 4122, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3824}, children = {}, }
_typeInfoList[471] = { parentId = 4110, typeId = 4124, baseId = 1, txt = 'get_nilBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3825}, children = {}, }
_typeInfoList[472] = { parentId = 112, typeId = 4126, baseId = 3596, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4130, 4134, 4136}, }
_typeInfoList[473] = { parentId = 112, typeId = 4127, nilable = true, orgTypeId = 4126 }_typeInfoList[474] = { parentId = 3706, typeId = 4128, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4126, 10}, retTypeId = {}, children = {}, }
_typeInfoList[475] = { parentId = 4126, typeId = 4130, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 1, typeId = 4132, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[477] = { parentId = 1, typeId = 4133, nilable = true, orgTypeId = 4132 }_typeInfoList[478] = { parentId = 4126, typeId = 4134, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4132, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 4126, typeId = 4136, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[480] = { parentId = 112, typeId = 4138, baseId = 3596, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4142, 4146, 4148, 4150}, }
_typeInfoList[481] = { parentId = 112, typeId = 4139, nilable = true, orgTypeId = 4138 }_typeInfoList[482] = { parentId = 3706, typeId = 4140, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4138, 10}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 4138, typeId = 4142, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[484] = { parentId = 1, typeId = 4144, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[485] = { parentId = 1, typeId = 4145, nilable = true, orgTypeId = 4144 }_typeInfoList[486] = { parentId = 4138, typeId = 4146, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4144, 3458, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[487] = { parentId = 4138, typeId = 4148, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[488] = { parentId = 4138, typeId = 4150, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[489] = { parentId = 112, typeId = 4152, baseId = 3596, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4156, 4160, 4162, 4164}, }
_typeInfoList[490] = { parentId = 112, typeId = 4153, nilable = true, orgTypeId = 4152 }_typeInfoList[491] = { parentId = 3706, typeId = 4154, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4152, 10}, retTypeId = {}, children = {}, }
_typeInfoList[492] = { parentId = 4152, typeId = 4156, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[493] = { parentId = 1, typeId = 4158, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[494] = { parentId = 1, typeId = 4159, nilable = true, orgTypeId = 4158 }_typeInfoList[495] = { parentId = 4152, typeId = 4160, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4158, 3596, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[496] = { parentId = 4152, typeId = 4162, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[497] = { parentId = 4152, typeId = 4164, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[498] = { parentId = 112, typeId = 4166, baseId = 3596, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4170, 4174, 4176, 4178}, }
_typeInfoList[499] = { parentId = 112, typeId = 4167, nilable = true, orgTypeId = 4166 }_typeInfoList[500] = { parentId = 3706, typeId = 4168, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4166, 10}, retTypeId = {}, children = {}, }
_typeInfoList[501] = { parentId = 4166, typeId = 4170, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[502] = { parentId = 1, typeId = 4172, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[503] = { parentId = 1, typeId = 4173, nilable = true, orgTypeId = 4172 }_typeInfoList[504] = { parentId = 4166, typeId = 4174, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4172, 3596, 3735}, retTypeId = {}, children = {}, }
_typeInfoList[505] = { parentId = 4166, typeId = 4176, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[506] = { parentId = 4166, typeId = 4178, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3735}, children = {}, }
_typeInfoList[507] = { parentId = 112, typeId = 4180, baseId = 3596, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4184, 4188, 4190}, }
_typeInfoList[508] = { parentId = 112, typeId = 4181, nilable = true, orgTypeId = 4180 }_typeInfoList[509] = { parentId = 3706, typeId = 4182, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4180, 10}, retTypeId = {}, children = {}, }
_typeInfoList[510] = { parentId = 4180, typeId = 4184, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[511] = { parentId = 1, typeId = 4186, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[512] = { parentId = 1, typeId = 4187, nilable = true, orgTypeId = 4186 }_typeInfoList[513] = { parentId = 4180, typeId = 4188, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4186, 3458}, retTypeId = {}, children = {}, }
_typeInfoList[514] = { parentId = 4180, typeId = 4190, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[515] = { parentId = 112, typeId = 4192, baseId = 3596, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4196, 4200, 4202}, }
_typeInfoList[516] = { parentId = 112, typeId = 4193, nilable = true, orgTypeId = 4192 }_typeInfoList[517] = { parentId = 3706, typeId = 4194, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4192, 10}, retTypeId = {}, children = {}, }
_typeInfoList[518] = { parentId = 4192, typeId = 4196, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[519] = { parentId = 1, typeId = 4198, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[520] = { parentId = 1, typeId = 4199, nilable = true, orgTypeId = 4198 }_typeInfoList[521] = { parentId = 4192, typeId = 4200, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4198, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[522] = { parentId = 4192, typeId = 4202, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[523] = { parentId = 112, typeId = 4204, baseId = 3596, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4208, 4214, 4218}, }
_typeInfoList[524] = { parentId = 112, typeId = 4205, nilable = true, orgTypeId = 4204 }_typeInfoList[525] = { parentId = 3706, typeId = 4206, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4204, 10}, retTypeId = {}, children = {}, }
_typeInfoList[526] = { parentId = 4204, typeId = 4208, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[527] = { parentId = 1, typeId = 4210, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[528] = { parentId = 1, typeId = 4211, nilable = true, orgTypeId = 4210 }_typeInfoList[529] = { parentId = 1, typeId = 4212, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[530] = { parentId = 1, typeId = 4213, nilable = true, orgTypeId = 4212 }_typeInfoList[531] = { parentId = 4204, typeId = 4214, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4210, 4212}, retTypeId = {}, children = {}, }
_typeInfoList[532] = { parentId = 1, typeId = 4216, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[533] = { parentId = 1, typeId = 4217, nilable = true, orgTypeId = 4216 }_typeInfoList[534] = { parentId = 4204, typeId = 4218, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4216}, children = {}, }
_typeInfoList[535] = { parentId = 112, typeId = 4220, baseId = 3596, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4224, 4230, 4234, 4700}, }
_typeInfoList[536] = { parentId = 112, typeId = 4221, nilable = true, orgTypeId = 4220 }_typeInfoList[537] = { parentId = 3706, typeId = 4222, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4220, 10}, retTypeId = {}, children = {}, }
_typeInfoList[538] = { parentId = 4220, typeId = 4224, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[539] = { parentId = 1, typeId = 4226, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 1, typeId = 4227, nilable = true, orgTypeId = 4226 }_typeInfoList[541] = { parentId = 1, typeId = 4228, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[542] = { parentId = 1, typeId = 4229, nilable = true, orgTypeId = 4228 }_typeInfoList[543] = { parentId = 4220, typeId = 4230, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4226, 4228}, retTypeId = {}, children = {}, }
_typeInfoList[544] = { parentId = 1, typeId = 4232, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[545] = { parentId = 1, typeId = 4233, nilable = true, orgTypeId = 4232 }_typeInfoList[546] = { parentId = 4220, typeId = 4234, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4232}, children = {}, }
_typeInfoList[547] = { parentId = 112, typeId = 4236, baseId = 3596, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4240, 4244, 4246}, }
_typeInfoList[548] = { parentId = 112, typeId = 4237, nilable = true, orgTypeId = 4236 }_typeInfoList[549] = { parentId = 3706, typeId = 4238, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4236, 10}, retTypeId = {}, children = {}, }
_typeInfoList[550] = { parentId = 4236, typeId = 4240, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[551] = { parentId = 1, typeId = 4242, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[552] = { parentId = 1, typeId = 4243, nilable = true, orgTypeId = 4242 }_typeInfoList[553] = { parentId = 4236, typeId = 4244, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4242, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[554] = { parentId = 4236, typeId = 4246, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[555] = { parentId = 112, typeId = 4248, baseId = 3596, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4252, 4256, 4258, 4260, 4694}, }
_typeInfoList[556] = { parentId = 112, typeId = 4249, nilable = true, orgTypeId = 4248 }_typeInfoList[557] = { parentId = 3706, typeId = 4250, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4248, 10}, retTypeId = {}, children = {}, }
_typeInfoList[558] = { parentId = 4248, typeId = 4252, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[559] = { parentId = 1, typeId = 4254, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[560] = { parentId = 1, typeId = 4255, nilable = true, orgTypeId = 4254 }_typeInfoList[561] = { parentId = 4248, typeId = 4256, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4254, 3458, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[562] = { parentId = 4248, typeId = 4258, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[563] = { parentId = 4248, typeId = 4260, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[564] = { parentId = 112, typeId = 4262, baseId = 3596, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4266, 4270, 4272, 4274, 4276}, }
_typeInfoList[565] = { parentId = 112, typeId = 4263, nilable = true, orgTypeId = 4262 }_typeInfoList[566] = { parentId = 3706, typeId = 4264, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4262, 10}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 4262, typeId = 4266, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[568] = { parentId = 1, typeId = 4268, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[569] = { parentId = 1, typeId = 4269, nilable = true, orgTypeId = 4268 }_typeInfoList[570] = { parentId = 4262, typeId = 4270, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4268, 3458, 3596, 3526}, retTypeId = {}, children = {}, }
_typeInfoList[571] = { parentId = 4262, typeId = 4272, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[572] = { parentId = 4262, typeId = 4274, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[573] = { parentId = 4262, typeId = 4276, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[574] = { parentId = 112, typeId = 4278, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4280, 4282, 4284, 4286}, }
_typeInfoList[575] = { parentId = 112, typeId = 4279, nilable = true, orgTypeId = 4278 }_typeInfoList[576] = { parentId = 4278, typeId = 4280, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[577] = { parentId = 4278, typeId = 4282, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3806}, children = {}, }
_typeInfoList[578] = { parentId = 4278, typeId = 4284, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[579] = { parentId = 4278, typeId = 4286, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3458, 3806, 3526}, retTypeId = {}, children = {}, }
_typeInfoList[580] = { parentId = 112, typeId = 4288, baseId = 3596, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4292, 4300, 4302, 4304, 4308, 4310, 4314, 4316, 4318}, }
_typeInfoList[581] = { parentId = 112, typeId = 4289, nilable = true, orgTypeId = 4288 }_typeInfoList[582] = { parentId = 3706, typeId = 4290, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4288, 10}, retTypeId = {}, children = {}, }
_typeInfoList[583] = { parentId = 4288, typeId = 4292, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[584] = { parentId = 1, typeId = 4294, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[585] = { parentId = 1, typeId = 4295, nilable = true, orgTypeId = 4294 }_typeInfoList[586] = { parentId = 1, typeId = 4296, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4278}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[587] = { parentId = 1, typeId = 4297, nilable = true, orgTypeId = 4296 }_typeInfoList[588] = { parentId = 1, typeId = 4298, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[589] = { parentId = 1, typeId = 4299, nilable = true, orgTypeId = 4298 }_typeInfoList[590] = { parentId = 4288, typeId = 4300, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4294, 20, 12, 4296, 3734, 4298, 12, 3825}, retTypeId = {}, children = {}, }
_typeInfoList[591] = { parentId = 4288, typeId = 4302, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[592] = { parentId = 4288, typeId = 4304, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[593] = { parentId = 1, typeId = 4306, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4278}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[594] = { parentId = 1, typeId = 4307, nilable = true, orgTypeId = 4306 }_typeInfoList[595] = { parentId = 4288, typeId = 4308, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4306}, children = {}, }
_typeInfoList[596] = { parentId = 4288, typeId = 4310, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[597] = { parentId = 1, typeId = 4312, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[598] = { parentId = 1, typeId = 4313, nilable = true, orgTypeId = 4312 }_typeInfoList[599] = { parentId = 4288, typeId = 4314, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4312}, children = {}, }
_typeInfoList[600] = { parentId = 4288, typeId = 4316, baseId = 1, txt = 'get_unwrapFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[601] = { parentId = 4288, typeId = 4318, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3825}, children = {}, }
_typeInfoList[602] = { parentId = 112, typeId = 4320, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4322, 4324, 4328, 4330, 4332, 4334, 4338, 4340}, }
_typeInfoList[603] = { parentId = 112, typeId = 4321, nilable = true, orgTypeId = 4320 }_typeInfoList[604] = { parentId = 4320, typeId = 4322, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[605] = { parentId = 4320, typeId = 4324, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[606] = { parentId = 1, typeId = 4326, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[607] = { parentId = 1, typeId = 4327, nilable = true, orgTypeId = 4326 }_typeInfoList[608] = { parentId = 4320, typeId = 4328, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4326}, children = {}, }
_typeInfoList[609] = { parentId = 4320, typeId = 4330, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[610] = { parentId = 4320, typeId = 4332, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[611] = { parentId = 4320, typeId = 4334, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[612] = { parentId = 1, typeId = 4336, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[613] = { parentId = 1, typeId = 4337, nilable = true, orgTypeId = 4336 }_typeInfoList[614] = { parentId = 4320, typeId = 4338, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4336}, children = {}, }
_typeInfoList[615] = { parentId = 4320, typeId = 4340, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3458, 3458, 12, 20, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[616] = { parentId = 112, typeId = 4342, baseId = 3596, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4346, 4350, 4352}, }
_typeInfoList[617] = { parentId = 112, typeId = 4343, nilable = true, orgTypeId = 4342 }_typeInfoList[618] = { parentId = 3706, typeId = 4344, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4342, 10}, retTypeId = {}, children = {}, }
_typeInfoList[619] = { parentId = 4342, typeId = 4346, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[620] = { parentId = 1, typeId = 4348, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[621] = { parentId = 1, typeId = 4349, nilable = true, orgTypeId = 4348 }_typeInfoList[622] = { parentId = 4342, typeId = 4350, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4348, 4320}, retTypeId = {}, children = {}, }
_typeInfoList[623] = { parentId = 4342, typeId = 4352, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4320}, children = {}, }
_typeInfoList[624] = { parentId = 112, typeId = 4354, baseId = 3596, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4358, 4362, 4364}, }
_typeInfoList[625] = { parentId = 112, typeId = 4355, nilable = true, orgTypeId = 4354 }_typeInfoList[626] = { parentId = 3706, typeId = 4356, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4354, 10}, retTypeId = {}, children = {}, }
_typeInfoList[627] = { parentId = 4354, typeId = 4358, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[628] = { parentId = 1, typeId = 4360, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[629] = { parentId = 1, typeId = 4361, nilable = true, orgTypeId = 4360 }_typeInfoList[630] = { parentId = 4354, typeId = 4362, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4360, 4320}, retTypeId = {}, children = {}, }
_typeInfoList[631] = { parentId = 4354, typeId = 4364, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4320}, children = {}, }
_typeInfoList[632] = { parentId = 112, typeId = 4366, baseId = 3596, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4370, 4374, 4376}, }
_typeInfoList[633] = { parentId = 112, typeId = 4367, nilable = true, orgTypeId = 4366 }_typeInfoList[634] = { parentId = 3706, typeId = 4368, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4366, 10}, retTypeId = {}, children = {}, }
_typeInfoList[635] = { parentId = 4366, typeId = 4370, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[636] = { parentId = 1, typeId = 4372, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[637] = { parentId = 1, typeId = 4373, nilable = true, orgTypeId = 4372 }_typeInfoList[638] = { parentId = 4366, typeId = 4374, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4372, 4320}, retTypeId = {}, children = {}, }
_typeInfoList[639] = { parentId = 4366, typeId = 4376, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4320}, children = {}, }
_typeInfoList[640] = { parentId = 112, typeId = 4378, baseId = 3596, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4382, 4386, 4388, 4390}, }
_typeInfoList[641] = { parentId = 112, typeId = 4379, nilable = true, orgTypeId = 4378 }_typeInfoList[642] = { parentId = 3706, typeId = 4380, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4378, 10}, retTypeId = {}, children = {}, }
_typeInfoList[643] = { parentId = 4378, typeId = 4382, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[644] = { parentId = 1, typeId = 4384, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[645] = { parentId = 1, typeId = 4385, nilable = true, orgTypeId = 4384 }_typeInfoList[646] = { parentId = 4378, typeId = 4386, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4384, 3526, 3734}, retTypeId = {}, children = {}, }
_typeInfoList[647] = { parentId = 4378, typeId = 4388, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[648] = { parentId = 4378, typeId = 4390, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[649] = { parentId = 112, typeId = 4392, baseId = 3596, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4396, 4400, 4402, 4404, 4406, 4408, 4410, 4412}, }
_typeInfoList[650] = { parentId = 112, typeId = 4393, nilable = true, orgTypeId = 4392 }_typeInfoList[651] = { parentId = 3706, typeId = 4394, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4392, 10}, retTypeId = {}, children = {}, }
_typeInfoList[652] = { parentId = 4392, typeId = 4396, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[653] = { parentId = 1, typeId = 4398, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[654] = { parentId = 1, typeId = 4399, nilable = true, orgTypeId = 4398 }_typeInfoList[655] = { parentId = 4392, typeId = 4400, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4398, 3458, 3806, 12, 20, 20, 20}, retTypeId = {}, children = {}, }
_typeInfoList[656] = { parentId = 4392, typeId = 4402, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[657] = { parentId = 4392, typeId = 4404, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3806}, children = {}, }
_typeInfoList[658] = { parentId = 4392, typeId = 4406, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[659] = { parentId = 4392, typeId = 4408, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[660] = { parentId = 4392, typeId = 4410, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[661] = { parentId = 4392, typeId = 4412, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[662] = { parentId = 112, typeId = 4414, baseId = 3596, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4418, 4422, 4424, 4426}, }
_typeInfoList[663] = { parentId = 112, typeId = 4415, nilable = true, orgTypeId = 4414 }_typeInfoList[664] = { parentId = 3706, typeId = 4416, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4414, 10}, retTypeId = {}, children = {}, }
_typeInfoList[665] = { parentId = 4414, typeId = 4418, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[666] = { parentId = 1, typeId = 4420, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[667] = { parentId = 1, typeId = 4421, nilable = true, orgTypeId = 4420 }_typeInfoList[668] = { parentId = 4414, typeId = 4422, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4420, 3458, 3806}, retTypeId = {}, children = {}, }
_typeInfoList[669] = { parentId = 4414, typeId = 4424, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[670] = { parentId = 4414, typeId = 4426, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3806}, children = {}, }
_typeInfoList[671] = { parentId = 112, typeId = 4428, baseId = 3596, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4432, 4436}, }
_typeInfoList[672] = { parentId = 112, typeId = 4429, nilable = true, orgTypeId = 4428 }_typeInfoList[673] = { parentId = 3706, typeId = 4430, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4428, 10}, retTypeId = {}, children = {}, }
_typeInfoList[674] = { parentId = 4428, typeId = 4432, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[675] = { parentId = 1, typeId = 4434, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[676] = { parentId = 1, typeId = 4435, nilable = true, orgTypeId = 4434 }_typeInfoList[677] = { parentId = 4428, typeId = 4436, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4434}, retTypeId = {}, children = {}, }
_typeInfoList[678] = { parentId = 3706, typeId = 4438, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3598, 10}, retTypeId = {}, children = {}, }
_typeInfoList[679] = { parentId = 3598, typeId = 4440, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[680] = { parentId = 1, typeId = 4442, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[681] = { parentId = 1, typeId = 4443, nilable = true, orgTypeId = 4442 }_typeInfoList[682] = { parentId = 1, typeId = 4444, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[683] = { parentId = 1, typeId = 4445, nilable = true, orgTypeId = 4444 }_typeInfoList[684] = { parentId = 1, typeId = 4446, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4392}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[685] = { parentId = 1, typeId = 4447, nilable = true, orgTypeId = 4446 }_typeInfoList[686] = { parentId = 1, typeId = 4448, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[687] = { parentId = 1, typeId = 4449, nilable = true, orgTypeId = 4448 }_typeInfoList[688] = { parentId = 3598, typeId = 4450, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4442, 20, 3458, 4444, 4446, 3682, 4448}, retTypeId = {}, children = {}, }
_typeInfoList[689] = { parentId = 3598, typeId = 4452, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[690] = { parentId = 3598, typeId = 4454, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[691] = { parentId = 1, typeId = 4456, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[692] = { parentId = 1, typeId = 4457, nilable = true, orgTypeId = 4456 }_typeInfoList[693] = { parentId = 3598, typeId = 4458, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4456}, children = {}, }
_typeInfoList[694] = { parentId = 1, typeId = 4460, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4392}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[695] = { parentId = 1, typeId = 4461, nilable = true, orgTypeId = 4460 }_typeInfoList[696] = { parentId = 3598, typeId = 4462, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4460}, children = {}, }
_typeInfoList[697] = { parentId = 3598, typeId = 4464, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3682}, children = {}, }
_typeInfoList[698] = { parentId = 1, typeId = 4466, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {20, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[699] = { parentId = 1, typeId = 4467, nilable = true, orgTypeId = 4466 }_typeInfoList[700] = { parentId = 3598, typeId = 4468, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4466}, children = {}, }
_typeInfoList[701] = { parentId = 112, typeId = 4470, baseId = 3596, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4474, 4478, 4480}, }
_typeInfoList[702] = { parentId = 112, typeId = 4471, nilable = true, orgTypeId = 4470 }_typeInfoList[703] = { parentId = 3706, typeId = 4472, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4470, 10}, retTypeId = {}, children = {}, }
_typeInfoList[704] = { parentId = 4470, typeId = 4474, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[705] = { parentId = 1, typeId = 4476, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[706] = { parentId = 1, typeId = 4477, nilable = true, orgTypeId = 4476 }_typeInfoList[707] = { parentId = 4470, typeId = 4478, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4476, 3736}, retTypeId = {}, children = {}, }
_typeInfoList[708] = { parentId = 4470, typeId = 4480, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3736}, children = {}, }
_typeInfoList[709] = { parentId = 112, typeId = 4482, baseId = 3596, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4486, 4490, 4634}, }
_typeInfoList[710] = { parentId = 112, typeId = 4483, nilable = true, orgTypeId = 4482 }_typeInfoList[711] = { parentId = 3706, typeId = 4484, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4482, 10}, retTypeId = {}, children = {}, }
_typeInfoList[712] = { parentId = 4482, typeId = 4486, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[713] = { parentId = 1, typeId = 4488, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[714] = { parentId = 1, typeId = 4489, nilable = true, orgTypeId = 4488 }_typeInfoList[715] = { parentId = 4482, typeId = 4490, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4488}, retTypeId = {}, children = {}, }
_typeInfoList[716] = { parentId = 112, typeId = 4492, baseId = 3596, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4496, 4500, 4502, 4504, 4640}, }
_typeInfoList[717] = { parentId = 112, typeId = 4493, nilable = true, orgTypeId = 4492 }_typeInfoList[718] = { parentId = 3706, typeId = 4494, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4492, 10}, retTypeId = {}, children = {}, }
_typeInfoList[719] = { parentId = 4492, typeId = 4496, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[720] = { parentId = 1, typeId = 4498, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[721] = { parentId = 1, typeId = 4499, nilable = true, orgTypeId = 4498 }_typeInfoList[722] = { parentId = 4492, typeId = 4500, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4498, 3458, 14}, retTypeId = {}, children = {}, }
_typeInfoList[723] = { parentId = 4492, typeId = 4502, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[724] = { parentId = 4492, typeId = 4504, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[725] = { parentId = 112, typeId = 4506, baseId = 3596, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4510, 4514, 4516, 4518, 4646}, }
_typeInfoList[726] = { parentId = 112, typeId = 4507, nilable = true, orgTypeId = 4506 }_typeInfoList[727] = { parentId = 3706, typeId = 4508, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4506, 10}, retTypeId = {}, children = {}, }
_typeInfoList[728] = { parentId = 4506, typeId = 4510, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[729] = { parentId = 1, typeId = 4512, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[730] = { parentId = 1, typeId = 4513, nilable = true, orgTypeId = 4512 }_typeInfoList[731] = { parentId = 4506, typeId = 4514, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4512, 3458, 14}, retTypeId = {}, children = {}, }
_typeInfoList[732] = { parentId = 4506, typeId = 4516, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[733] = { parentId = 4506, typeId = 4518, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[734] = { parentId = 112, typeId = 4520, baseId = 3596, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4524, 4528, 4530, 4532, 4652}, }
_typeInfoList[735] = { parentId = 112, typeId = 4521, nilable = true, orgTypeId = 4520 }_typeInfoList[736] = { parentId = 3706, typeId = 4522, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4520, 10}, retTypeId = {}, children = {}, }
_typeInfoList[737] = { parentId = 4520, typeId = 4524, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[738] = { parentId = 1, typeId = 4526, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[739] = { parentId = 1, typeId = 4527, nilable = true, orgTypeId = 4526 }_typeInfoList[740] = { parentId = 4520, typeId = 4528, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4526, 3458, 16}, retTypeId = {}, children = {}, }
_typeInfoList[741] = { parentId = 4520, typeId = 4530, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[742] = { parentId = 4520, typeId = 4532, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {16}, children = {}, }
_typeInfoList[743] = { parentId = 112, typeId = 4534, baseId = 3596, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4538, 4542, 4544, 4658}, }
_typeInfoList[744] = { parentId = 112, typeId = 4535, nilable = true, orgTypeId = 4534 }_typeInfoList[745] = { parentId = 3706, typeId = 4536, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4534, 10}, retTypeId = {}, children = {}, }
_typeInfoList[746] = { parentId = 4534, typeId = 4538, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[747] = { parentId = 1, typeId = 4540, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[748] = { parentId = 1, typeId = 4541, nilable = true, orgTypeId = 4540 }_typeInfoList[749] = { parentId = 4534, typeId = 4542, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4540, 3734}, retTypeId = {}, children = {}, }
_typeInfoList[750] = { parentId = 4534, typeId = 4544, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[751] = { parentId = 112, typeId = 4546, baseId = 3596, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4550, 4554, 4556, 4664}, }
_typeInfoList[752] = { parentId = 112, typeId = 4547, nilable = true, orgTypeId = 4546 }_typeInfoList[753] = { parentId = 3706, typeId = 4548, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4546, 10}, retTypeId = {}, children = {}, }
_typeInfoList[754] = { parentId = 4546, typeId = 4550, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[755] = { parentId = 1, typeId = 4552, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[756] = { parentId = 1, typeId = 4553, nilable = true, orgTypeId = 4552 }_typeInfoList[757] = { parentId = 4546, typeId = 4554, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4552, 3734}, retTypeId = {}, children = {}, }
_typeInfoList[758] = { parentId = 4546, typeId = 4556, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3734}, children = {}, }
_typeInfoList[759] = { parentId = 112, typeId = 4558, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4560, 4562, 4564}, }
_typeInfoList[760] = { parentId = 112, typeId = 4559, nilable = true, orgTypeId = 4558 }_typeInfoList[761] = { parentId = 4558, typeId = 4560, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[762] = { parentId = 4558, typeId = 4562, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[763] = { parentId = 4558, typeId = 4564, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3596, 3596}, retTypeId = {}, children = {}, }
_typeInfoList[764] = { parentId = 112, typeId = 4566, baseId = 3596, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4570, 4578, 4582, 4586, 4670}, }
_typeInfoList[765] = { parentId = 112, typeId = 4567, nilable = true, orgTypeId = 4566 }_typeInfoList[766] = { parentId = 3706, typeId = 4568, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4566, 10}, retTypeId = {}, children = {}, }
_typeInfoList[767] = { parentId = 4566, typeId = 4570, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[768] = { parentId = 1, typeId = 4572, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[769] = { parentId = 1, typeId = 4573, nilable = true, orgTypeId = 4572 }_typeInfoList[770] = { parentId = 1, typeId = 4574, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3596, 3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[771] = { parentId = 1, typeId = 4575, nilable = true, orgTypeId = 4574 }_typeInfoList[772] = { parentId = 1, typeId = 4576, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4558}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[773] = { parentId = 1, typeId = 4577, nilable = true, orgTypeId = 4576 }_typeInfoList[774] = { parentId = 4566, typeId = 4578, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4572, 4574, 4576}, retTypeId = {}, children = {}, }
_typeInfoList[775] = { parentId = 1, typeId = 4580, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {3596, 3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[776] = { parentId = 1, typeId = 4581, nilable = true, orgTypeId = 4580 }_typeInfoList[777] = { parentId = 4566, typeId = 4582, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4580}, children = {}, }
_typeInfoList[778] = { parentId = 1, typeId = 4584, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4558}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[779] = { parentId = 1, typeId = 4585, nilable = true, orgTypeId = 4584 }_typeInfoList[780] = { parentId = 4566, typeId = 4586, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4584}, children = {}, }
_typeInfoList[781] = { parentId = 112, typeId = 4588, baseId = 3596, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4592, 4598, 4600, 4604, 4676}, }
_typeInfoList[782] = { parentId = 112, typeId = 4589, nilable = true, orgTypeId = 4588 }_typeInfoList[783] = { parentId = 3706, typeId = 4590, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4588, 10}, retTypeId = {}, children = {}, }
_typeInfoList[784] = { parentId = 4588, typeId = 4592, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[785] = { parentId = 1, typeId = 4594, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[786] = { parentId = 1, typeId = 4595, nilable = true, orgTypeId = 4594 }_typeInfoList[787] = { parentId = 1, typeId = 4596, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[788] = { parentId = 1, typeId = 4597, nilable = true, orgTypeId = 4596 }_typeInfoList[789] = { parentId = 4588, typeId = 4598, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4594, 3458, 4596}, retTypeId = {}, children = {}, }
_typeInfoList[790] = { parentId = 4588, typeId = 4600, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[791] = { parentId = 1, typeId = 4602, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3596}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[792] = { parentId = 1, typeId = 4603, nilable = true, orgTypeId = 4602 }_typeInfoList[793] = { parentId = 4588, typeId = 4604, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4602}, children = {}, }
_typeInfoList[794] = { parentId = 112, typeId = 4606, baseId = 3596, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4610, 4614, 4616, 4682}, }
_typeInfoList[795] = { parentId = 112, typeId = 4607, nilable = true, orgTypeId = 4606 }_typeInfoList[796] = { parentId = 3706, typeId = 4608, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4606, 10}, retTypeId = {}, children = {}, }
_typeInfoList[797] = { parentId = 4606, typeId = 4610, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[798] = { parentId = 1, typeId = 4612, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[799] = { parentId = 1, typeId = 4613, nilable = true, orgTypeId = 4612 }_typeInfoList[800] = { parentId = 4606, typeId = 4614, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4612, 3458}, retTypeId = {}, children = {}, }
_typeInfoList[801] = { parentId = 4606, typeId = 4616, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[802] = { parentId = 112, typeId = 4618, baseId = 3596, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4622, 4626, 4628, 4688}, }
_typeInfoList[803] = { parentId = 112, typeId = 4619, nilable = true, orgTypeId = 4618 }_typeInfoList[804] = { parentId = 3706, typeId = 4620, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4618, 10}, retTypeId = {}, children = {}, }
_typeInfoList[805] = { parentId = 4618, typeId = 4622, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3706, 10}, retTypeId = {}, children = {}, }
_typeInfoList[806] = { parentId = 1, typeId = 4624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[807] = { parentId = 1, typeId = 4625, nilable = true, orgTypeId = 4624 }_typeInfoList[808] = { parentId = 4618, typeId = 4626, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3454, 4624, 3458}, retTypeId = {}, children = {}, }
_typeInfoList[809] = { parentId = 4618, typeId = 4628, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3458}, children = {}, }
_typeInfoList[810] = { parentId = 1, typeId = 4630, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[811] = { parentId = 1, typeId = 4631, nilable = true, orgTypeId = 4630 }_typeInfoList[812] = { parentId = 1, typeId = 4632, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[813] = { parentId = 1, typeId = 4633, nilable = true, orgTypeId = 4632 }_typeInfoList[814] = { parentId = 4482, typeId = 4634, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4630, 4632}, children = {}, }
_typeInfoList[815] = { parentId = 1, typeId = 4636, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[816] = { parentId = 1, typeId = 4637, nilable = true, orgTypeId = 4636 }_typeInfoList[817] = { parentId = 1, typeId = 4638, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[818] = { parentId = 1, typeId = 4639, nilable = true, orgTypeId = 4638 }_typeInfoList[819] = { parentId = 4492, typeId = 4640, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4636, 4638}, children = {}, }
_typeInfoList[820] = { parentId = 1, typeId = 4642, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[821] = { parentId = 1, typeId = 4643, nilable = true, orgTypeId = 4642 }_typeInfoList[822] = { parentId = 1, typeId = 4644, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[823] = { parentId = 1, typeId = 4645, nilable = true, orgTypeId = 4644 }_typeInfoList[824] = { parentId = 4506, typeId = 4646, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4642, 4644}, children = {}, }
_typeInfoList[825] = { parentId = 1, typeId = 4648, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[826] = { parentId = 1, typeId = 4649, nilable = true, orgTypeId = 4648 }_typeInfoList[827] = { parentId = 1, typeId = 4650, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[828] = { parentId = 1, typeId = 4651, nilable = true, orgTypeId = 4650 }_typeInfoList[829] = { parentId = 4520, typeId = 4652, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4648, 4650}, children = {}, }
_typeInfoList[830] = { parentId = 1, typeId = 4654, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[831] = { parentId = 1, typeId = 4655, nilable = true, orgTypeId = 4654 }_typeInfoList[832] = { parentId = 1, typeId = 4656, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[833] = { parentId = 1, typeId = 4657, nilable = true, orgTypeId = 4656 }_typeInfoList[834] = { parentId = 4534, typeId = 4658, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4654, 4656}, children = {}, }
_typeInfoList[835] = { parentId = 1, typeId = 4660, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[836] = { parentId = 1, typeId = 4661, nilable = true, orgTypeId = 4660 }_typeInfoList[837] = { parentId = 1, typeId = 4662, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[838] = { parentId = 1, typeId = 4663, nilable = true, orgTypeId = 4662 }_typeInfoList[839] = { parentId = 4546, typeId = 4664, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4660, 4662}, children = {}, }
_typeInfoList[840] = { parentId = 1, typeId = 4666, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[841] = { parentId = 1, typeId = 4667, nilable = true, orgTypeId = 4666 }_typeInfoList[842] = { parentId = 1, typeId = 4668, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[843] = { parentId = 1, typeId = 4669, nilable = true, orgTypeId = 4668 }_typeInfoList[844] = { parentId = 4566, typeId = 4670, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4666, 4668}, children = {}, }
_typeInfoList[845] = { parentId = 1, typeId = 4672, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[846] = { parentId = 1, typeId = 4673, nilable = true, orgTypeId = 4672 }_typeInfoList[847] = { parentId = 1, typeId = 4674, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[848] = { parentId = 1, typeId = 4675, nilable = true, orgTypeId = 4674 }_typeInfoList[849] = { parentId = 4588, typeId = 4676, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4672, 4674}, children = {}, }
_typeInfoList[850] = { parentId = 1, typeId = 4678, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[851] = { parentId = 1, typeId = 4679, nilable = true, orgTypeId = 4678 }_typeInfoList[852] = { parentId = 1, typeId = 4680, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[853] = { parentId = 1, typeId = 4681, nilable = true, orgTypeId = 4680 }_typeInfoList[854] = { parentId = 4606, typeId = 4682, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4678, 4680}, children = {}, }
_typeInfoList[855] = { parentId = 1, typeId = 4684, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[856] = { parentId = 1, typeId = 4685, nilable = true, orgTypeId = 4684 }_typeInfoList[857] = { parentId = 1, typeId = 4686, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[858] = { parentId = 1, typeId = 4687, nilable = true, orgTypeId = 4686 }_typeInfoList[859] = { parentId = 4618, typeId = 4688, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4684, 4686}, children = {}, }
_typeInfoList[860] = { parentId = 1, typeId = 4690, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[861] = { parentId = 1, typeId = 4691, nilable = true, orgTypeId = 4690 }_typeInfoList[862] = { parentId = 1, typeId = 4692, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[863] = { parentId = 1, typeId = 4693, nilable = true, orgTypeId = 4692 }_typeInfoList[864] = { parentId = 4248, typeId = 4694, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4690, 4692}, children = {}, }
_typeInfoList[865] = { parentId = 1, typeId = 4696, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[866] = { parentId = 1, typeId = 4697, nilable = true, orgTypeId = 4696 }_typeInfoList[867] = { parentId = 1, typeId = 4698, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {3526}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[868] = { parentId = 1, typeId = 4699, nilable = true, orgTypeId = 4698 }_typeInfoList[869] = { parentId = 4220, typeId = 4700, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4696, 4698}, children = {}, }
_typeInfoList[870] = { parentId = 3732, typeId = 4702, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4470}, retTypeId = {28}, children = {}, }
_typeInfoList[871] = { parentId = 3732, typeId = 4704, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[872] = { parentId = 112, typeId = 4706, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4708, 4710, 4712}, }
_typeInfoList[873] = { parentId = 112, typeId = 4707, nilable = true, orgTypeId = 4706 }_typeInfoList[874] = { parentId = 4706, typeId = 4708, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3596}, children = {}, }
_typeInfoList[875] = { parentId = 4706, typeId = 4710, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {3526}, children = {}, }
_typeInfoList[876] = { parentId = 4706, typeId = 4712, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3596, 3526}, retTypeId = {}, children = {}, }
_typeInfoList[877] = { parentId = 3752, typeId = 4714, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3470, 12, 20}, retTypeId = {4706}, children = {}, }
_typeInfoList[878] = { parentId = 110, typeId = 4716, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4718, 4724, 4730, 4734, 4746, 4754, 4762, 4772, 4774, 4776, 4780}, }
_typeInfoList[879] = { parentId = 110, typeId = 4717, nilable = true, orgTypeId = 4716 }_typeInfoList[880] = { parentId = 4716, typeId = 4718, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4720, 4722}, }
_typeInfoList[881] = { parentId = 4716, typeId = 4719, nilable = true, orgTypeId = 4718 }_typeInfoList[882] = { parentId = 4718, typeId = 4720, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[883] = { parentId = 4718, typeId = 4722, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[884] = { parentId = 4716, typeId = 4724, baseId = 4718, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4726, 4728}, }
_typeInfoList[885] = { parentId = 4716, typeId = 4725, nilable = true, orgTypeId = 4724 }_typeInfoList[886] = { parentId = 4724, typeId = 4726, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[887] = { parentId = 4724, typeId = 4728, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {21}, children = {}, }
_typeInfoList[888] = { parentId = 4716, typeId = 4730, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4732}, }
_typeInfoList[889] = { parentId = 4716, typeId = 4731, nilable = true, orgTypeId = 4730 }_typeInfoList[890] = { parentId = 4730, typeId = 4732, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 14}, retTypeId = {}, children = {}, }
_typeInfoList[891] = { parentId = 4716, typeId = 4734, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4738, 4742, 4744}, }
_typeInfoList[892] = { parentId = 4716, typeId = 4735, nilable = true, orgTypeId = 4734 }_typeInfoList[893] = { parentId = 1, typeId = 4736, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4734}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[894] = { parentId = 1, typeId = 4737, nilable = true, orgTypeId = 4736 }_typeInfoList[895] = { parentId = 4734, typeId = 4738, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4736}, retTypeId = {}, children = {}, }
_typeInfoList[896] = { parentId = 1, typeId = 4740, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {4734}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[897] = { parentId = 1, typeId = 4741, nilable = true, orgTypeId = 4740 }_typeInfoList[898] = { parentId = 4734, typeId = 4742, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4740}, children = {}, }
_typeInfoList[899] = { parentId = 4734, typeId = 4744, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {14, 20, 4730}, retTypeId = {}, children = {}, }
_typeInfoList[900] = { parentId = 4716, typeId = 4746, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4748, 4750, 4752}, }
_typeInfoList[901] = { parentId = 4716, typeId = 4747, nilable = true, orgTypeId = 4746 }_typeInfoList[902] = { parentId = 4746, typeId = 4748, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4734}, children = {}, }
_typeInfoList[903] = { parentId = 4746, typeId = 4750, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[904] = { parentId = 4746, typeId = 4752, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[905] = { parentId = 4716, typeId = 4754, baseId = 4746, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4756, 4758, 4760}, }
_typeInfoList[906] = { parentId = 4716, typeId = 4755, nilable = true, orgTypeId = 4754 }_typeInfoList[907] = { parentId = 4754, typeId = 4756, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4734}, children = {}, }
_typeInfoList[908] = { parentId = 4754, typeId = 4758, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[909] = { parentId = 4754, typeId = 4760, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4746, 20}, retTypeId = {}, children = {}, }
_typeInfoList[910] = { parentId = 4716, typeId = 4762, baseId = 4746, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4764, 4766, 4768, 4778}, }
_typeInfoList[911] = { parentId = 4716, typeId = 4763, nilable = true, orgTypeId = 4762 }_typeInfoList[912] = { parentId = 4762, typeId = 4764, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4718, 20, 12}, retTypeId = {}, children = {}, }
_typeInfoList[913] = { parentId = 4762, typeId = 4766, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[914] = { parentId = 4762, typeId = 4768, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20, 12}, retTypeId = {4763}, children = {}, }
_typeInfoList[915] = { parentId = 4716, typeId = 4772, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {14}, retTypeId = {20}, children = {}, }
_typeInfoList[916] = { parentId = 4716, typeId = 4774, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[917] = { parentId = 4716, typeId = 4776, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {6}, children = {}, }
_typeInfoList[918] = { parentId = 4762, typeId = 4778, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {4735}, children = {}, }
_typeInfoList[919] = { parentId = 4716, typeId = 4780, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {4734}, children = {}, }
_typeInfoList[920] = { parentId = 110, typeId = 4782, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4784, 4790, 4798, 4800, 4802}, }
_typeInfoList[921] = { parentId = 110, typeId = 4783, nilable = true, orgTypeId = 4782 }_typeInfoList[922] = { parentId = 4782, typeId = 4784, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4786, 4788}, }
_typeInfoList[923] = { parentId = 4782, typeId = 4785, nilable = true, orgTypeId = 4784 }_typeInfoList[924] = { parentId = 4784, typeId = 4786, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[925] = { parentId = 4784, typeId = 4788, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[926] = { parentId = 4782, typeId = 4790, baseId = 4784, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4792, 4794, 4796}, }
_typeInfoList[927] = { parentId = 4782, typeId = 4791, nilable = true, orgTypeId = 4790 }_typeInfoList[928] = { parentId = 4790, typeId = 4792, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[929] = { parentId = 4790, typeId = 4794, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[930] = { parentId = 4790, typeId = 4796, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[931] = { parentId = 4782, typeId = 4798, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {20}, retTypeId = {}, children = {}, }
_typeInfoList[932] = { parentId = 4782, typeId = 4800, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[933] = { parentId = 4782, typeId = 4802, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 6, 20}, retTypeId = {}, children = {}, }
_typeInfoList[934] = { parentId = 106, typeId = 4812, baseId = 3706, txt = 'convFilter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {4866, 4868, 4870, 4896, 4898, 4900, 4910, 4912, 4914, 4916, 4918, 4920, 4924, 4926, 4930, 4932, 4934, 4936, 4938, 4940, 4942, 4944, 4946, 4948, 4950, 4952, 4954, 4956, 4958, 4960, 4962, 4964, 4966, 4968, 4970, 4972, 4974, 4976, 4978, 4980, 4982, 4984, 4986, 4988, 4990, 4992, 4994, 4996, 4998, 5002, 5004, 5006, 5008}, }
_typeInfoList[935] = { parentId = 106, typeId = 4813, nilable = true, orgTypeId = 4812 }_typeInfoList[936] = { parentId = 4812, typeId = 4866, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3762, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[937] = { parentId = 4812, typeId = 4868, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3772, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[938] = { parentId = 4812, typeId = 4870, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3784, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[939] = { parentId = 4812, typeId = 4896, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3824, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[940] = { parentId = 4812, typeId = 4898, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4236, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[941] = { parentId = 4812, typeId = 4900, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3598, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[942] = { parentId = 4812, typeId = 4910, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4392, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[943] = { parentId = 4812, typeId = 4912, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4204, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[944] = { parentId = 4812, typeId = 4914, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4470, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[945] = { parentId = 4812, typeId = 4916, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4220, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[946] = { parentId = 4812, typeId = 4918, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4038, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[947] = { parentId = 4812, typeId = 4920, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4366, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[948] = { parentId = 4812, typeId = 4924, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4378, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[949] = { parentId = 4812, typeId = 4926, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4354, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[950] = { parentId = 4812, typeId = 4930, baseId = 1, txt = 'processUnwrapSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4094, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[951] = { parentId = 4812, typeId = 4932, baseId = 1, txt = 'processIfUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4110, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[952] = { parentId = 4812, typeId = 4934, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4288, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[953] = { parentId = 4812, typeId = 4936, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4414, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[954] = { parentId = 4812, typeId = 4938, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4428, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[955] = { parentId = 4812, typeId = 4940, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4180, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[956] = { parentId = 4812, typeId = 4942, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4342, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[957] = { parentId = 4812, typeId = 4944, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3806, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[958] = { parentId = 4812, typeId = 4946, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3852, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[959] = { parentId = 4812, typeId = 4948, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3890, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[960] = { parentId = 4812, typeId = 4950, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3910, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[961] = { parentId = 4812, typeId = 4952, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3924, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[962] = { parentId = 4812, typeId = 4954, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3938, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[963] = { parentId = 4812, typeId = 4956, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3958, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[964] = { parentId = 4812, typeId = 4958, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3978, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[965] = { parentId = 4812, typeId = 4960, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3996, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[966] = { parentId = 4812, typeId = 4962, baseId = 1, txt = 'processExpUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4052, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[967] = { parentId = 4812, typeId = 4964, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4166, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[968] = { parentId = 4812, typeId = 4966, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3734, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[969] = { parentId = 4812, typeId = 4968, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4138, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[970] = { parentId = 4812, typeId = 4970, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4126, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[971] = { parentId = 4812, typeId = 4972, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4192, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[972] = { parentId = 4812, typeId = 4974, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4078, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[973] = { parentId = 4812, typeId = 4976, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4066, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[974] = { parentId = 4812, typeId = 4978, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4152, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[975] = { parentId = 4812, typeId = 4980, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4248, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[976] = { parentId = 4812, typeId = 4982, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4262, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[977] = { parentId = 4812, typeId = 4984, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4016, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[978] = { parentId = 4812, typeId = 4986, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4546, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[979] = { parentId = 4812, typeId = 4988, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4566, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[980] = { parentId = 4812, typeId = 4990, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4534, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[981] = { parentId = 4812, typeId = 4992, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4492, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[982] = { parentId = 4812, typeId = 4994, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4506, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[983] = { parentId = 4812, typeId = 4996, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4520, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[984] = { parentId = 4812, typeId = 4998, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4588, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[985] = { parentId = 4812, typeId = 5002, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4606, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[986] = { parentId = 4812, typeId = 5004, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4482, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[987] = { parentId = 4812, typeId = 5006, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4028, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[988] = { parentId = 4812, typeId = 5008, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4618, 3596, 14}, retTypeId = {}, children = {}, }
_typeInfoList[989] = { parentId = 106, typeId = 5010, baseId = 3732, txt = 'MacroEvalImp',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {5012, 5014}, }
_typeInfoList[990] = { parentId = 106, typeId = 5011, nilable = true, orgTypeId = 5010 }_typeInfoList[991] = { parentId = 5010, typeId = 5012, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {4470}, retTypeId = {28}, children = {}, }
_typeInfoList[992] = { parentId = 5010, typeId = 5014, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
