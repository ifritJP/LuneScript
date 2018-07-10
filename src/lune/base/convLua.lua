--lune/base/convLua.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local Parser = require( 'lune.base.Parser' )

local filterObj = {}
moduleObj.filterObj = filterObj
function filterObj.new( streamName, stream, exeFlag )
  local obj = {}
  setmetatable( obj, { __index = filterObj } )
  if obj.__init then obj:__init( streamName, stream, exeFlag ); end
return obj
end
function filterObj:__init(streamName, stream, exeFlag) 
  self.streamName = streamName
  self.stream = stream
  self.moduleName2Info = {}
  self.exeFlag = exeFlag
  self.indent = 0
  self.curLineNo = 1
  self.className2Scope = {}
  self.className2TypeInfo = {}
  self.className2MemberList = {}
  self.pubVarName2InfoMap = {}
  self.pubFuncName2InfoMap = {}
  self.needIndent = false
end

local stepIndent = 2
local builtInModuleSet = {}
builtInModuleSet["io"] = true
builtInModuleSet["string"] = true
builtInModuleSet["table"] = true
builtInModuleSet["math"] = true
builtInModuleSet["debug"] = true
builtInModuleSet["_luneScript"] = true
function filterObj:write( txt )
  if self.needIndent then
    self.stream:write( string.rep( " ", self.indent ) )
    self.needIndent = false
  end
  for cr in string.gmatch( txt, "\n" ) do
    self.curLineNo = self.curLineNo + 1
  end
  self.stream:write( txt )
end

function filterObj:setIndent( indent )
  self.indent = indent
end

function filterObj:writeln( txt, baseIndent )
  self:write( txt )
  self:write( "\n" )
  self.needIndent = true
  self.indent = baseIndent
end

filterObj[TransUnit.nodeKind.None] = function ( self, node, parent, baseIndent )
  self:writeln( "-- none", baseIndent )
end

filterObj[TransUnit.nodeKind.Import] = function ( self, node, parent, baseIndent )
  local module = node.info
  local moduleName = string.gsub( module, ".*%.", "" )
  local moduleInfo = true
  self.moduleName2Info[moduleName] = moduleInfo
  if self.exeFlag then
    self:writeln( string.format( "local %s = _luneScript.loadModule( '%s' )", moduleName, module), baseIndent )
  else 
    self:writeln( string.format( "local %s = require( '%s' )", moduleName, module), baseIndent )
  end
end

filterObj[TransUnit.nodeKind.Root] = function ( self, node, parent, baseIndent )
  self:writeln( string.format( "--%s", self.streamName), baseIndent )
  self:writeln( "local moduleObj = {}", baseIndent )
  local children = node.info.children
  for __index, child in pairs( children ) do
    TransUnit.nodeFilter( child, self, node, baseIndent )
    self:writeln( "", baseIndent )
  end
  self:writeln( "----- meta -----", baseIndent )
  local typeId2TypeInfo = {}
  local typeId2UseFlag = {}
  local function pickupTypeId( typeInfo, forceFlag )
    if typeInfo then
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
        typeInfoList = typeInfo:get_retTypeInfoList(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            pickupTypeId( itemTypeInfo )
          end
        end
        typeInfoList = typeInfo:get_children(  )
        if typeInfoList then
          for __index, itemTypeInfo in pairs( typeInfoList ) do
            if itemTypeInfo:get_kind(  ) == TransUnit.TypeInfoKindClass or itemTypeInfo:get_kind(  ) == TransUnit.TypeInfoKindFunc then
              pickupTypeId( itemTypeInfo )
            end
          end
        end
      end
    end
  end
  
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
        self:writeln( string.format( "local _classInfo%s = {}", className), baseIndent )
        self:writeln( string.format( "_className2InfoMap.%s = _classInfo%s", className, className), baseIndent )
        pickupTypeId( self.className2TypeInfo[className] )
        for __index, memberNode in pairs( self.className2MemberList[className] ) do
          local memberInfo = memberNode:get_info(  )
          if memberInfo.accessMode == "pub" then
            local memberName = memberInfo.name.txt
            local memberTypeInfo = memberNode:get_expType(  )
            self:writeln( string.format( "_classInfo%s.%s = {", className, memberName), baseIndent )
            self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberInfo.staticFlag) .. string.format( "accessMode = '%s', methodFlag = false, typeId = %d }", memberInfo.accessMode, memberTypeInfo:get_typeId(  )), baseIndent )
          end
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
        self:writeln( string.format( "  name='%s', accessMode = '%s', typeId = %d }", varName, varInfo["accessMode"], varInfo["typeInfo"]:get_typeId(  )), baseIndent )
        pickupTypeId( varInfo["typeInfo"], true )
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
        pickupTypeId( funcInfo["typeInfo"] )
      end
    end
  end
  
  self:writeln( "local _typeInfoList = {}", baseIndent )
  self:writeln( "moduleObj._typeInfoList = _typeInfoList", baseIndent )
  local listIndex = 1
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
        for __index, child in pairs( typeInfo:get_children(  ) ) do
          local child = child
          local childTypeId = child:get_typeId(  )
          if typeId2TypeInfo[childTypeId] and not TransUnit.isBuiltin( childTypeId ) then
            self:write( string.format( "_typeInfoList[%d] = ", listIndex) )
            listIndex = listIndex + 1
            child:serialize( self )
          end
        end
      end
    end
  end
  
  self:writeln( "----- meta -----", baseIndent )
  self:writeln( "return moduleObj", baseIndent )
end

filterObj[TransUnit.nodeKind.Block] = function ( self, node, parent, baseIndent )
  local word = ""
  if node.info.kind == "if" or node.info.kind == "elseif" then
    word = "then"
  elseif node.info.kind == "else" then
    word = ""
  elseif node.info.kind == "while" then
    word = "do"
  elseif node.info.kind == "repeat" then
    word = ""
  elseif node.info.kind == "for" then
    word = "do"
  elseif node.info.kind == "apply" then
    word = "do"
  elseif node.info.kind == "foreach" then
    word = "do"
  elseif node.info.kind == "func" then
    word = ""
  elseif node.info.kind == "default" then
    word = ""
  elseif node.info.kind == "{" then
    word = "do"
  end
  self:writeln( word, baseIndent + stepIndent )
  local stmtList = node.info.stmtList
  for __index, statement in pairs( stmtList ) do
    TransUnit.nodeFilter( statement, self, node, baseIndent + stepIndent )
    self:writeln( "", baseIndent + stepIndent )
  end
  self:setIndent( baseIndent )
  if node.info.kind == "{" then
    self:write( "end", baseIndent )
  end
end

filterObj[TransUnit.nodeKind.StmtExp] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info, self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.DeclClass] = function ( self, node, parent, baseIndent )
  local className = node.info.name.txt
  if node.info.accessMode == "pub" then
    self.className2Scope[className] = node.info.scope
    self.className2TypeInfo[className] = node.expType
  end
  self.className2MemberList[className] = node.info.memberList
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
  local baseInfo = node.expType:get_baseTypeInfo(  )
  if baseInfo then
    self:writeln( string.format( "setmetatable( %s, { __index = %s } )", className, baseInfo:getTxt(  )), baseIndent )
  end
  if node.info.accessMode == "pub" then
    self:writeln( string.format( "moduleObj.%s = %s", className, className ), baseIndent )
  end
  local hasConstrFlag = false
  local memberList = {}
  local fieldList = node.info.fieldList
  local outerMethodSet = node.info.outerMethodSet
  for __index, field in pairs( fieldList ) do
    local ignoreFlag = false
    if field["kind"] == TransUnit.nodeKind.DeclConstr then
      hasConstrFlag = true
    end
    if field["kind"] == TransUnit.nodeKind.DeclMember then
      table.insert( memberList, field )
    end
    if field["kind"] == TransUnit.nodeKind.DeclMethod then
      if outerMethodSet[field.info.name.txt] then
        ignoreFlag = true
      end
    end
    if (not ignoreFlag ) then
      TransUnit.nodeFilter( field, self, node, baseIndent )
    end
  end
  if not hasConstrFlag then
    local argTxt = ""
    for index, member in pairs( memberList ) do
      if index > 1 then
        argTxt = argTxt .. ", "
      end
      argTxt = argTxt .. member["info"].name.txt
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
      local memberName = member["info"].name.txt
      self:writeln( string.format( "self.%s = %s", memberName, memberName ), baseIndent + stepIndent )
    end
    self:writeln( 'end', baseIndent )
  end
  local scope = node.info.scope
  for __index, memberNode in pairs( node.info.memberList ) do
    local memberName = memberNode.info.name.txt
    local getterName = "get_" .. memberName
    local typeInfo = scope:getTypeInfo( getterName )
    local autoFlag = not typeInfo or typeInfo:get_autoFlag(  )
    if memberNode.info.getterMode ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
   return self.%s
end]==], className, getterName, memberName), baseIndent )
    end
    local setterName = "set_" .. memberName
    typeInfo = scope:getTypeInfo( setterName )
    if memberNode.info.setterMode ~= "none" and autoFlag then
      self:writeln( string.format( [==[
function %s:%s()
   return self.%s
end]==], className, setterName, memberName), baseIndent )
    end
  end
end

filterObj[TransUnit.nodeKind.DeclMember] = function ( self, node, parent, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpNew] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info.symbol, self, node, baseIndent )
  self:write( ".new(" )
  if node.info.argList then
    TransUnit.nodeFilter( node.info.argList, self, node, baseIndent )
  end
  self:write( ")" )
end

filterObj[TransUnit.nodeKind.DeclConstr] = function ( self, node, parent, baseIndent )
  local className = node.info.className.txt
  self:write( string.format( "function %s.new( ", className ) )
  local argTxt = ""
  local argList = node.info.argList
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
      argTxt = argTxt .. ", "
    end
    TransUnit.nodeFilter( arg, self, node, baseIndent )
    argTxt = argTxt .. arg["info"].name.txt
  end
  self:writeln( " )", baseIndent + stepIndent )
  self:writeln( "local obj = {}", baseIndent + stepIndent )
  self:writeln( string.format( "setmetatable( obj, { __index = %s } )", className ), baseIndent + stepIndent )
  self:writeln( string.format( "if obj.__init then obj:__init( %s ); end", argTxt ), baseIndent )
  self:writeln( "return obj", baseIndent )
  self:writeln( "end", baseIndent )
  self:write( string.format( "function %s:__init(%s) ", className, argTxt ) )
  TransUnit.nodeFilter( node.info.body, self, node, baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.DeclMethod] = function ( self, node, parent, baseIndent )
  local delimit = ":"
  if node.info.staticFlag then
    delimit = "."
  end
  local methodName = node.info.name.txt
  self:write( string.format( "function %s%s%s( ", node.info.className.txt, delimit, methodName) )
  local argList = node.info.argList
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    TransUnit.nodeFilter( arg, self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  TransUnit.nodeFilter( node.info.body, self, node, baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.DeclVar] = function ( self, node, parent, baseIndent )
  if node.info.accessMode ~= "global" then
    self:write( "local " )
  end
  local varName = ""
  local varList = node.info.varList
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var["name"].txt )
  end
  if node.info.expList then
    self:write( " = " )
    TransUnit.nodeFilter( node.info.expList, self, node, baseIndent )
  end
  if node.info.unwrap then
    self:writeln( "", baseIndent )
    self:write( "if " )
    for index, var in pairs( varList ) do
      if index > 1 then
        self:write( " or " )
      end
      self:write( " not " .. var["name"].txt )
    end
    self:write( " then" )
    TransUnit.nodeFilter( node.info.unwrap, self, node, baseIndent )
    self:writeln( "end", baseIndent )
  end
  if node.info.accessMode == "pub" then
    self:writeln( "", baseIndent )
    local varList = node.info.varList
    for index, var in pairs( varList ) do
      local name = var["name"].txt
      self:writeln( string.format( "moduleObj.%s = %s", name, name), baseIndent )
      self.pubVarName2InfoMap[name] = {["funcFlag"] = false, ["staticFlag"] = node.info.staticFlag, ["accessMode"] = node.info.accessMode, ["typeInfo"] = node.info.typeInfoList[index]}
    end
  end
end

filterObj[TransUnit.nodeKind.DeclArg] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%s", node.info.name.txt ) )
end

filterObj[TransUnit.nodeKind.DeclArgDDD] = function ( self, node, parent, baseIndent )
  self:write( "..." )
end

filterObj[TransUnit.nodeKind.ExpDDD] = function ( self, node, parent, baseIndent )
  self:write( "..." )
end

filterObj[TransUnit.nodeKind.DeclFunc] = function ( self, node, parent, baseIndent )
  local nodeInfo = node:get_info(  )
  local nameToken = nodeInfo:get_name(  )
  local name = nameToken and nameToken.txt or ""
  local letTxt = ""
  if nodeInfo:get_accessMode(  ) ~= "global" and #name ~= 0 then
    letTxt = "local "
  end
  self:write( string.format( "%sfunction %s( ", letTxt, name ) )
  local argList = nodeInfo:get_argList(  )
  for index, arg in pairs( argList ) do
    if index > 1 then
      self:write( ", " )
    end
    TransUnit.nodeFilter( arg, self, node, baseIndent )
  end
  self:write( " )", baseIndent )
  TransUnit.nodeFilter( nodeInfo:get_body(  ), self, node, baseIndent )
  self:writeln( "end", baseIndent )
  local expType = node:get_expType(  )
  if expType:get_accessMode(  ) == "pub" then
    self:write( string.format( "moduleObj.%s = %s", name, name) )
    self.pubFuncName2InfoMap[name] = {["funcFlag"] = true, ["accessMode"] = nodeInfo:get_accessMode(  ), ["typeInfo"] = node:get_expType(  )}
  end
end

filterObj[TransUnit.nodeKind.RefType] = function ( self, node, parent, baseIndent )
  self:write( (node.info.refFlag and "&" or "" ) .. (node.info.mutFlag and "mut " or "" ) )
  TransUnit.nodeFilter( node.info.name, self, node, baseIndent )
  if node.info.array == "array" then
    self:write( "[@]" )
  elseif node.info.array == "list" then
    self:write( "[]" )
  end
end

filterObj[TransUnit.nodeKind.If] = function ( self, node, parent, baseIndent )
  local valList = node.info
  for index, val in pairs( valList ) do
    if index == 1 then
      self:write( "if " )
      TransUnit.nodeFilter( val["exp"], self, node, baseIndent )
    elseif val["kind"] == "elseif" then
      self:write( "elseif " )
      TransUnit.nodeFilter( val["exp"], self, node, baseIndent )
    else 
      self:write( "else" )
    end
    self:write( " " )
    TransUnit.nodeFilter( val["block"], self, node, baseIndent )
  end
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Switch] = function ( self, node, parent, baseIndent )
  self:writeln( "do", baseIndent + 2 )
  self:write( "local _switchExp = " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent + 2 )
  self:writeln( "", baseIndent + 2 )
  for index, caseInfo in pairs( node.info.caseList ) do
    if index == 1 then
      self:write( "if " )
    else 
      self:write( "elseif " )
    end
    local expList = caseInfo.expList
    for index, expNode in pairs( expList.info ) do
      if index ~= 1 then
        self:write( " or " )
      end
      self:write( "_switchExp == " )
      TransUnit.nodeFilter( expNode, self, node, baseIndent + 2 )
    end
    self:write( " then" )
    TransUnit.nodeFilter( caseInfo.block, self, node, baseIndent + 2 )
  end
  if node.info.default then
    self:write( "else " )
    TransUnit.nodeFilter( node.info.default, self, node, baseIndent + 2 )
    self:writeln( "end", baseIndent )
  end
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.While] = function ( self, node, parent, baseIndent )
  self:write( "while " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent )
  self:write( " " )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Repeat] = function ( self, node, parent, baseIndent )
  self:write( "repeat " )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent )
  self:write( "until " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.For] = function ( self, node, parent, baseIndent )
  self:write( string.format( "for %s = ", node.info.val.txt ) )
  TransUnit.nodeFilter( node.info.init, self, node, baseIndent )
  self:write( ", " )
  TransUnit.nodeFilter( node.info.to, self, node, baseIndent )
  if node.info.delta then
    self:write( ", " )
    TransUnit.nodeFilter( node.info.delta, self, node, baseIndent )
  end
  self:write( " " )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Apply] = function ( self, node, parent, baseIndent )
  self:write( "for " )
  local varList = node.info.varList
  for index, var in pairs( varList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( var["txt"] )
  end
  self:write( " in " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent )
  self:write( " " )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Foreach] = function ( self, node, parent, baseIndent )
  self:write( "for " )
  self:write( node.info.key and node.info.key.txt or "__index" )
  self:write( ", " )
  self:write( node.info.val.txt )
  self:write( " in pairs( " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent )
  self:write( " ) " )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent )
  self:write( "end" )
end

filterObj[TransUnit.nodeKind.Forsort] = function ( self, node, parent, baseIndent )
  self:writeln( "do", baseIndent + stepIndent )
  self:writeln( "local __sorted = {}", baseIndent + stepIndent )
  self:write( "local __map = " )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent + stepIndent )
  self:writeln( "", baseIndent + stepIndent )
  self:writeln( "for __key in pairs( __map ) do", baseIndent + stepIndent * 2 )
  self:writeln( "table.insert( __sorted, __key )", baseIndent + stepIndent )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "table.sort( __sorted )", baseIndent + stepIndent )
  self:write( "for __index, " )
  local key = node.info.key and node.info.key.txt or "__key"
  self:write( key )
  self:writeln( " in ipairs( __sorted ) do", baseIndent + stepIndent * 2 )
  self:writeln( string.format( "%s = __map[ %s ]", node.info.val.txt, key ), baseIndent + stepIndent * 2 )
  TransUnit.nodeFilter( node.info.block, self, node, baseIndent + stepIndent * 2 )
  self:writeln( "end", baseIndent + stepIndent )
  self:writeln( "end", baseIndent )
  self:writeln( "end", baseIndent )
end

filterObj[TransUnit.nodeKind.ExpCall] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info.func, self, node, baseIndent )
  self:write( "( " )
  if node.info.argList then
    TransUnit.nodeFilter( node.info.argList, self, node, baseIndent )
  end
  self:write( " )" )
end

filterObj[TransUnit.nodeKind.ExpList] = function ( self, node, parent, baseIndent )
  local expList = node.info
  for index, exp in pairs( expList ) do
    if index > 1 then
      self:write( ", " )
    end
    TransUnit.nodeFilter( exp, self, node, baseIndent )
  end
end

filterObj[TransUnit.nodeKind.ExpOp1] = function ( self, node, parent, baseIndent )
  local op = node.info.op.txt
  if op == "not" then
    op = op .. " "
  end
  self:write( op )
  TransUnit.nodeFilter( node.info.exp, self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpCast] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info, self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpParen] = function ( self, node, parent, baseIndent )
  self:write( "(" )
  TransUnit.nodeFilter( node.info, self, node, baseIndent )
  self:write( " )" )
end

filterObj[TransUnit.nodeKind.ExpOp2] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info.exp1, self, node, baseIndent )
  self:write( " " .. node.info.op.txt .. " " )
  TransUnit.nodeFilter( node.info.exp2, self, node, baseIndent )
end

filterObj[TransUnit.nodeKind.ExpRef] = function ( self, node, parent, baseIndent )
  self:write( node.info.txt )
end

filterObj[TransUnit.nodeKind.ExpRefItem] = function ( self, node, parent, baseIndent )
  if node.info.val.kind == TransUnit.nodeKind.LiteralString then
    self:write( "string.byte( " )
    TransUnit.nodeFilter( node.info.val, self, node, baseIndent )
    self:write( ", " )
    TransUnit.nodeFilter( node.info.index, self, node, baseIndent )
    self:write( " )" )
  else 
    TransUnit.nodeFilter( node.info.val, self, node, baseIndent )
    self:write( "[" )
    TransUnit.nodeFilter( node.info.index, self, node, baseIndent )
    self:write( "]" )
  end
end

filterObj[TransUnit.nodeKind.RefField] = function ( self, node, parent, baseIndent )
  TransUnit.nodeFilter( node.info.prefix, self, node, baseIndent )
  local delimit = "."
  if parent.kind == TransUnit.nodeKind.ExpCall then
    local prefixSymbol = node.info.prefix.info.txt
    if node.expType:get_staticFlag(  ) then
      delimit = "."
    else 
      delimit = ":"
    end
  end
  self:write( delimit .. node.info.field.txt )
end

filterObj[TransUnit.nodeKind.Return] = function ( self, node, parent, baseIndent )
  self:write( "return " )
  if node.info then
    TransUnit.nodeFilter( node.info, self, node, baseIndent )
  end
end

filterObj[TransUnit.nodeKind.LiteralList] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  if node.info then
    TransUnit.nodeFilter( node.info, self, node, baseIndent )
  end
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralMap] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  local pairList = node.info.pairList
  for index, pair in pairs( pairList ) do
    if index > 1 then
      self:write( ", " )
    end
    self:write( "[" )
    TransUnit.nodeFilter( pair["key"], self, node, baseIndent )
    self:write( "] = " )
    TransUnit.nodeFilter( pair["val"], self, node, baseIndent )
    index = index + 1
  end
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralArray] = function ( self, node, parent, baseIndent )
  self:write( "{" )
  if node.info then
    TransUnit.nodeFilter( node.info, self, node, baseIndent )
  end
  self:write( "}" )
end

filterObj[TransUnit.nodeKind.LiteralChar] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%g", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralInt] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%d", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralReal] = function ( self, node, parent, baseIndent )
  self:write( string.format( "%s", node.info.num ) )
end

filterObj[TransUnit.nodeKind.LiteralString] = function ( self, node, parent, baseIndent )
  local nodeInfo = node:get_info(  )
  local txt = (nodeInfo:get_token(  ) ).txt
  if string.find( txt, '^```' ) then
    txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
  end
  if #nodeInfo:get_argList(  ) > 0 then
    self:write( string.format( "string.format( %s, ", txt ) )
    local argList = nodeInfo:get_argList(  )
    for index, val in pairs( argList ) do
      if index > 1 then
        self:write( ", " )
      end
      TransUnit.nodeFilter( val, self, node, baseIndent )
    end
    self:write( ")" )
  else 
    self:write( txt )
  end
end

filterObj[TransUnit.nodeKind.LiteralBool] = function ( self, node, parent, baseIndent )
  self:write( node.info.txt )
end

filterObj[TransUnit.nodeKind.LiteralNil] = function ( self, node, parent, baseIndent )
  self:write( "nil" )
end

filterObj[TransUnit.nodeKind.Break] = function ( self, node, parent, baseIndent )
  self:write( "break" )
end

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfofilterObj = {}
_className2InfoMap.filterObj = _classInfofilterObj
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 88, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[2] = { parentId = 1, typeId = 150, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[3] = { parentId = 1, typeId = 260, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[4] = { parentId = 1, typeId = 262, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[5] = { parentId = 1, typeId = 264, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {282, 284, 340}, }
_typeInfoList[6] = { parentId = 1, typeId = 286, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[7] = { parentId = 1, typeId = 312, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[8] = { parentId = 1, typeId = 314, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[9] = { parentId = 1, typeId = 316, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[10] = { parentId = 1, typeId = 348, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[11] = { parentId = 1, typeId = 350, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[12] = { parentId = 1, typeId = 352, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[13] = { parentId = 1, typeId = 354, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[14] = { parentId = 1, typeId = 356, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 1, typeId = 358, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {372, 374, 376}, }
_typeInfoList[16] = { parentId = 1, typeId = 360, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[17] = { parentId = 1, typeId = 364, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[18] = { parentId = 1, typeId = 366, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[19] = { parentId = 1, typeId = 368, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[20] = { parentId = 1, typeId = 370, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[21] = { parentId = 1, typeId = 378, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[22] = { parentId = 1, typeId = 380, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[23] = { parentId = 1, typeId = 392, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[24] = { parentId = 1, typeId = 398, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {418, 420, 422, 424, 430, 432, 434, 440, 442, 446, 450, 452, 456, 462, 470, 474, 476, 478, 480, 482, 484, 486, 488, 490, 492, 494, 498}, }
_typeInfoList[25] = { parentId = 1, typeId = 468, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[26] = { parentId = 1, typeId = 472, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[27] = { parentId = 1, typeId = 496, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[28] = { parentId = 1, typeId = 500, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {516, 518, 520, 522, 524, 526, 530, 534}, }
_typeInfoList[29] = { parentId = 1, typeId = 528, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 398}, retTypeId = {}, children = {}, }
_typeInfoList[30] = { parentId = 1, typeId = 532, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 500}, retTypeId = {}, children = {}, }
_typeInfoList[31] = { parentId = 1, typeId = 538, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 1, typeId = 540, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {544, 546, 550}, }
_typeInfoList[33] = { parentId = 1, typeId = 552, baseId = 540, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[34] = { parentId = 1, typeId = 554, baseId = 540, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[35] = { parentId = 1, typeId = 556, baseId = 540, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 1, typeId = 558, baseId = 540, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 1, typeId = 560, baseId = 540, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[38] = { parentId = 1, typeId = 562, baseId = 540, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 1, typeId = 564, baseId = 540, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[40] = { parentId = 1, typeId = 566, baseId = 540, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 1, typeId = 568, baseId = 540, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 1, typeId = 570, baseId = 540, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[43] = { parentId = 1, typeId = 572, baseId = 540, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 1, typeId = 574, baseId = 540, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 1, typeId = 576, baseId = 540, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[46] = { parentId = 1, typeId = 578, baseId = 540, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[47] = { parentId = 1, typeId = 580, baseId = 540, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[48] = { parentId = 1, typeId = 582, baseId = 540, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[49] = { parentId = 1, typeId = 584, baseId = 540, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[50] = { parentId = 1, typeId = 586, baseId = 540, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[51] = { parentId = 1, typeId = 588, baseId = 540, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[52] = { parentId = 1, typeId = 590, baseId = 540, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[53] = { parentId = 1, typeId = 592, baseId = 540, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[54] = { parentId = 1, typeId = 594, baseId = 540, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[55] = { parentId = 1, typeId = 596, baseId = 540, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[56] = { parentId = 1, typeId = 598, baseId = 540, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 1, typeId = 600, baseId = 540, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[58] = { parentId = 1, typeId = 602, baseId = 540, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[59] = { parentId = 1, typeId = 604, baseId = 540, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[60] = { parentId = 1, typeId = 606, baseId = 540, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[61] = { parentId = 1, typeId = 608, baseId = 540, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 1, typeId = 610, baseId = 540, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[63] = { parentId = 1, typeId = 612, baseId = 540, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[64] = { parentId = 1, typeId = 614, baseId = 540, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[65] = { parentId = 1, typeId = 616, baseId = 540, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[66] = { parentId = 1, typeId = 618, baseId = 540, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[67] = { parentId = 1, typeId = 620, baseId = 540, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[68] = { parentId = 1, typeId = 622, baseId = 540, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[69] = { parentId = 1, typeId = 624, baseId = 540, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[70] = { parentId = 1, typeId = 628, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {708, 1250}, }
_typeInfoList[71] = { parentId = 1, typeId = 706, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[72] = { parentId = 1, typeId = 760, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[73] = { parentId = 1, typeId = 762, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[74] = { parentId = 1, typeId = 1094, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1102, 1104, 1108, 1110, 1112, 1114, 1118, 1122}, }
_typeInfoList[75] = { parentId = 1, typeId = 1106, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {540}, retTypeId = {}, children = {}, }
_typeInfoList[76] = { parentId = 1, typeId = 1116, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 1, typeId = 1120, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {398}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 1, typeId = 1208, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1212, 1216}, }
_typeInfoList[79] = { parentId = 1, typeId = 1214, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {540}, retTypeId = {}, children = {}, }
_typeInfoList[80] = { parentId = 1, typeId = 1272, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 1, typeId = 1274, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[82] = { parentId = 1, typeId = 1276, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[83] = { parentId = 1, typeId = 1278, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[84] = { parentId = 1, typeId = 1280, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1434, 1436, 1438}, }
_typeInfoList[85] = { parentId = 1, typeId = 1282, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[86] = { parentId = 1, typeId = 1284, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[87] = { parentId = 1, typeId = 1286, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[88] = { parentId = 1, typeId = 1288, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[89] = { parentId = 1, typeId = 1290, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[90] = { parentId = 1, typeId = 1292, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 1, typeId = 1294, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 1, typeId = 1296, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 1298, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[94] = { parentId = 1, typeId = 1300, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1440, 1442, 1444}, }
_typeInfoList[95] = { parentId = 1, typeId = 1302, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 1304, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[97] = { parentId = 1, typeId = 1306, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[98] = { parentId = 1, typeId = 1308, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 1310, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[100] = { parentId = 1, typeId = 1312, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[101] = { parentId = 1, typeId = 1314, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[102] = { parentId = 1, typeId = 1318, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[103] = { parentId = 1, typeId = 1320, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1446, 1448, 1450, 1452, 1454, 1456, 1458, 1460, 1462, 1464, 1466, 1468, 1470, 1472, 1474, 1476, 1478, 1480, 1482, 1484, 1486, 1488, 1490, 1492, 1494, 1496, 1498}, }
_typeInfoList[104] = { parentId = 1, typeId = 1322, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[105] = { parentId = 1, typeId = 1324, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[106] = { parentId = 1, typeId = 1326, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 1, typeId = 1328, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1500, 1502, 1504, 1506, 1508, 1510, 1512, 1514}, }
_typeInfoList[108] = { parentId = 1, typeId = 1330, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1320}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 1, typeId = 1332, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1328}, retTypeId = {}, children = {}, }
_typeInfoList[110] = { parentId = 1, typeId = 1334, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[111] = { parentId = 1, typeId = 1336, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1516, 1518, 1520}, }
_typeInfoList[112] = { parentId = 1, typeId = 1338, baseId = 1336, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[113] = { parentId = 1, typeId = 1340, baseId = 1336, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 1342, baseId = 1336, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 1, typeId = 1344, baseId = 1336, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[116] = { parentId = 1, typeId = 1346, baseId = 1336, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 1, typeId = 1348, baseId = 1336, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[118] = { parentId = 1, typeId = 1350, baseId = 1336, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 1, typeId = 1352, baseId = 1336, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 1, typeId = 1354, baseId = 1336, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[121] = { parentId = 1, typeId = 1356, baseId = 1336, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[122] = { parentId = 1, typeId = 1358, baseId = 1336, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 1, typeId = 1360, baseId = 1336, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[124] = { parentId = 1, typeId = 1362, baseId = 1336, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[125] = { parentId = 1, typeId = 1364, baseId = 1336, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 1, typeId = 1366, baseId = 1336, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[127] = { parentId = 1, typeId = 1368, baseId = 1336, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 1, typeId = 1370, baseId = 1336, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[129] = { parentId = 1, typeId = 1372, baseId = 1336, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[130] = { parentId = 1, typeId = 1374, baseId = 1336, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 1, typeId = 1376, baseId = 1336, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[132] = { parentId = 1, typeId = 1378, baseId = 1336, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 1, typeId = 1380, baseId = 1336, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[134] = { parentId = 1, typeId = 1382, baseId = 1336, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[135] = { parentId = 1, typeId = 1384, baseId = 1336, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 1, typeId = 1386, baseId = 1336, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[137] = { parentId = 1, typeId = 1388, baseId = 1336, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[138] = { parentId = 1, typeId = 1390, baseId = 1336, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 1, typeId = 1392, baseId = 1336, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1, typeId = 1394, baseId = 1336, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 1, typeId = 1396, baseId = 1336, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[142] = { parentId = 1, typeId = 1398, baseId = 1336, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 1400, baseId = 1336, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[144] = { parentId = 1, typeId = 1402, baseId = 1336, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 1404, baseId = 1336, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[146] = { parentId = 1, typeId = 1406, baseId = 1336, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[147] = { parentId = 1, typeId = 1408, baseId = 1336, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[148] = { parentId = 1, typeId = 1410, baseId = 1336, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[149] = { parentId = 1, typeId = 1412, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1522, 1524}, }
_typeInfoList[150] = { parentId = 1, typeId = 1414, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 1, typeId = 1418, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[152] = { parentId = 1, typeId = 1420, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[153] = { parentId = 1, typeId = 1422, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1526, 1528, 1530, 1532, 1534, 1536, 1538, 1540}, }
_typeInfoList[154] = { parentId = 1, typeId = 1424, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1336}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 1, typeId = 1426, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 1, typeId = 1428, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1320}, retTypeId = {}, children = {}, }
_typeInfoList[157] = { parentId = 1, typeId = 1430, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1542, 1544}, }
_typeInfoList[158] = { parentId = 1, typeId = 1432, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1336}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 1, typeId = 1546, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 1, typeId = 1548, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[161] = { parentId = 1, typeId = 1550, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[162] = { parentId = 1, typeId = 1552, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[163] = { parentId = 1, typeId = 1554, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 1, typeId = 1556, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1570, 1572, 1574}, }
_typeInfoList[165] = { parentId = 1, typeId = 1558, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[166] = { parentId = 1, typeId = 1562, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[167] = { parentId = 1, typeId = 1564, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[168] = { parentId = 1, typeId = 1566, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[169] = { parentId = 1, typeId = 1568, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[170] = { parentId = 1, typeId = 1576, baseId = 1, txt = 'filterObj',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[171] = { parentId = 264, typeId = 282, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[172] = { parentId = 264, typeId = 284, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {264}, children = {}, }
_typeInfoList[173] = { parentId = 264, typeId = 340, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[174] = { parentId = 358, typeId = 372, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[175] = { parentId = 358, typeId = 374, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {358}, children = {}, }
_typeInfoList[176] = { parentId = 358, typeId = 376, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[177] = { parentId = 398, typeId = 418, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[178] = { parentId = 398, typeId = 420, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[179] = { parentId = 398, typeId = 422, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[180] = { parentId = 398, typeId = 424, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 398, typeId = 430, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[182] = { parentId = 398, typeId = 432, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[183] = { parentId = 398, typeId = 434, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[184] = { parentId = 398, typeId = 440, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[185] = { parentId = 398, typeId = 442, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[186] = { parentId = 398, typeId = 446, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[187] = { parentId = 398, typeId = 450, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[188] = { parentId = 398, typeId = 452, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[189] = { parentId = 398, typeId = 456, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[190] = { parentId = 398, typeId = 462, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[191] = { parentId = 398, typeId = 470, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {468}, children = {}, }
_typeInfoList[192] = { parentId = 398, typeId = 474, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {472}, children = {}, }
_typeInfoList[193] = { parentId = 398, typeId = 476, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[194] = { parentId = 398, typeId = 478, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[195] = { parentId = 398, typeId = 480, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[196] = { parentId = 398, typeId = 482, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[197] = { parentId = 398, typeId = 484, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[198] = { parentId = 398, typeId = 486, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[199] = { parentId = 398, typeId = 488, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[200] = { parentId = 398, typeId = 490, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[201] = { parentId = 398, typeId = 492, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[202] = { parentId = 398, typeId = 494, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[203] = { parentId = 398, typeId = 498, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {496}, children = {}, }
_typeInfoList[204] = { parentId = 500, typeId = 516, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 500, typeId = 518, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 500, typeId = 520, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {500}, children = {}, }
_typeInfoList[207] = { parentId = 500, typeId = 522, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[208] = { parentId = 500, typeId = 524, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[209] = { parentId = 500, typeId = 526, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {500}, children = {}, }
_typeInfoList[210] = { parentId = 500, typeId = 530, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {528}, children = {}, }
_typeInfoList[211] = { parentId = 500, typeId = 534, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {532}, children = {}, }
_typeInfoList[212] = { parentId = 540, typeId = 544, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[213] = { parentId = 540, typeId = 546, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {398}, children = {}, }
_typeInfoList[214] = { parentId = 540, typeId = 550, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[215] = { parentId = 628, typeId = 708, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {706}, children = {}, }
_typeInfoList[216] = { parentId = 628, typeId = 1250, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[217] = { parentId = 1094, typeId = 1102, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[218] = { parentId = 1094, typeId = 1104, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[219] = { parentId = 1094, typeId = 1108, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1106}, children = {}, }
_typeInfoList[220] = { parentId = 1094, typeId = 1110, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[221] = { parentId = 1094, typeId = 1112, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[222] = { parentId = 1094, typeId = 1114, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {540}, children = {}, }
_typeInfoList[223] = { parentId = 1094, typeId = 1118, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1116}, children = {}, }
_typeInfoList[224] = { parentId = 1094, typeId = 1122, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1120}, children = {}, }
_typeInfoList[225] = { parentId = 1208, typeId = 1212, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {356}, children = {}, }
_typeInfoList[226] = { parentId = 1208, typeId = 1216, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1214}, children = {}, }
_typeInfoList[227] = { parentId = 1280, typeId = 1434, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[228] = { parentId = 1280, typeId = 1436, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
_typeInfoList[229] = { parentId = 1280, typeId = 1438, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[230] = { parentId = 1300, typeId = 1440, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[231] = { parentId = 1300, typeId = 1442, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1300}, children = {}, }
_typeInfoList[232] = { parentId = 1300, typeId = 1444, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[233] = { parentId = 1320, typeId = 1446, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[234] = { parentId = 1320, typeId = 1448, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[235] = { parentId = 1320, typeId = 1450, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[236] = { parentId = 1320, typeId = 1452, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 1320, typeId = 1454, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[238] = { parentId = 1320, typeId = 1456, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[239] = { parentId = 1320, typeId = 1458, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[240] = { parentId = 1320, typeId = 1460, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[241] = { parentId = 1320, typeId = 1462, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[242] = { parentId = 1320, typeId = 1464, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[243] = { parentId = 1320, typeId = 1466, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[244] = { parentId = 1320, typeId = 1468, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[245] = { parentId = 1320, typeId = 1470, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[246] = { parentId = 1320, typeId = 1472, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[247] = { parentId = 1320, typeId = 1474, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1322}, children = {}, }
_typeInfoList[248] = { parentId = 1320, typeId = 1476, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1324}, children = {}, }
_typeInfoList[249] = { parentId = 1320, typeId = 1478, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[250] = { parentId = 1320, typeId = 1480, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[251] = { parentId = 1320, typeId = 1482, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[252] = { parentId = 1320, typeId = 1484, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[253] = { parentId = 1320, typeId = 1486, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[254] = { parentId = 1320, typeId = 1488, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[255] = { parentId = 1320, typeId = 1490, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[256] = { parentId = 1320, typeId = 1492, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[257] = { parentId = 1320, typeId = 1494, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[258] = { parentId = 1320, typeId = 1496, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[259] = { parentId = 1320, typeId = 1498, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[260] = { parentId = 1328, typeId = 1500, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[261] = { parentId = 1328, typeId = 1502, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[262] = { parentId = 1328, typeId = 1504, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1328}, children = {}, }
_typeInfoList[263] = { parentId = 1328, typeId = 1506, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[264] = { parentId = 1328, typeId = 1508, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[265] = { parentId = 1328, typeId = 1510, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1328}, children = {}, }
_typeInfoList[266] = { parentId = 1328, typeId = 1512, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1330}, children = {}, }
_typeInfoList[267] = { parentId = 1328, typeId = 1514, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1332}, children = {}, }
_typeInfoList[268] = { parentId = 1336, typeId = 1516, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[269] = { parentId = 1336, typeId = 1518, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1320}, children = {}, }
_typeInfoList[270] = { parentId = 1336, typeId = 1520, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[271] = { parentId = 1412, typeId = 1522, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1414}, children = {}, }
_typeInfoList[272] = { parentId = 1412, typeId = 1524, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
_typeInfoList[273] = { parentId = 1422, typeId = 1526, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[274] = { parentId = 1422, typeId = 1528, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[275] = { parentId = 1422, typeId = 1530, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1424}, children = {}, }
_typeInfoList[276] = { parentId = 1422, typeId = 1532, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[277] = { parentId = 1422, typeId = 1534, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[278] = { parentId = 1422, typeId = 1536, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1336}, children = {}, }
_typeInfoList[279] = { parentId = 1422, typeId = 1538, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1426}, children = {}, }
_typeInfoList[280] = { parentId = 1422, typeId = 1540, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1428}, children = {}, }
_typeInfoList[281] = { parentId = 1430, typeId = 1542, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[282] = { parentId = 1430, typeId = 1544, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1432}, children = {}, }
_typeInfoList[283] = { parentId = 1556, typeId = 1570, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[284] = { parentId = 1556, typeId = 1572, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1556}, children = {}, }
_typeInfoList[285] = { parentId = 1556, typeId = 1574, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
----- meta -----
return moduleObj
