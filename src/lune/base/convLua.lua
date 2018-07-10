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
  
  self:writeln( "local _funcName2InfoMap = {}", baseIndent )
  self:writeln( "moduleObj._funcName2InfoMap = _funcName2InfoMap", baseIndent )
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
  
  self:writeln( "moduleObj._typeInfoList = {}", baseIndent )
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
            self:writeln( "table.insert( ", baseIndent )
            self:write( "moduleObj._typeInfoList, " )
            child:serialize( self )
            self:writeln( ")", baseIndent )
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
local _funcName2InfoMap = {}
moduleObj._funcName2InfoMap = _funcName2InfoMap
moduleObj._typeInfoList = {}
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 84, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 142, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 248, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 250, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 252, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {270, 272, 328}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 274, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 300, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 302, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 304, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 336, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 338, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 340, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 342, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 344, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 346, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {360, 362, 364}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 348, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 352, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 354, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 356, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 358, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 366, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 368, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 380, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 386, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {406, 408, 410, 412, 418, 420, 422, 428, 430, 434, 438, 440, 444, 450, 458, 462, 464, 466, 468, 470, 472, 474, 476, 478, 480, 482, 486}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 456, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 460, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 484, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 488, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {504, 506, 508, 510, 512, 514, 518, 522}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 516, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 520, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 488}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 526, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 528, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {532, 534, 538}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 540, baseId = 528, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 542, baseId = 528, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 544, baseId = 528, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 546, baseId = 528, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 548, baseId = 528, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 550, baseId = 528, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 552, baseId = 528, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 554, baseId = 528, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 556, baseId = 528, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 558, baseId = 528, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 560, baseId = 528, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 562, baseId = 528, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 564, baseId = 528, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 566, baseId = 528, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 568, baseId = 528, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 570, baseId = 528, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 572, baseId = 528, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 574, baseId = 528, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 576, baseId = 528, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 578, baseId = 528, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 580, baseId = 528, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 582, baseId = 528, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 584, baseId = 528, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 586, baseId = 528, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 588, baseId = 528, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 590, baseId = 528, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 592, baseId = 528, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 594, baseId = 528, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 596, baseId = 528, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 598, baseId = 528, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 600, baseId = 528, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 602, baseId = 528, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 604, baseId = 528, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 606, baseId = 528, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 608, baseId = 528, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 610, baseId = 528, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 612, baseId = 528, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 616, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {696, 1232}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 694, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 748, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 750, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1076, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1084, 1086, 1090, 1092, 1094, 1096, 1100, 1104}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1088, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {528}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1098, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1102, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {386}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1190, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1194, 1198}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1196, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {528}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1254, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1256, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1258, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1260, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1262, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1416, 1418, 1420}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1264, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1266, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1268, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1270, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1272, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1274, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1276, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1278, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1280, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1282, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1422, 1424, 1426}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1284, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1286, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1288, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1290, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1292, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1294, baseId = 1, txt = 'errorLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1296, baseId = 1, txt = 'debugLog',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1300, baseId = 1, txt = 'isBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1302, baseId = 1, txt = 'TypeInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1428, 1430, 1432, 1434, 1436, 1438, 1440, 1442, 1444, 1446, 1448, 1450, 1452, 1454, 1456, 1458, 1460, 1462, 1464, 1466, 1468, 1470, 1472, 1474, 1476, 1478, 1480}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1304, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1306, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1308, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1310, baseId = 1, txt = 'Scope',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1482, 1484, 1486, 1488, 1490, 1492, 1494, 1496}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1312, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1314, baseId = 1, txt = 'Map',
staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18, 1310}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1316, baseId = 1, txt = 'NodePos',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1318, baseId = 1, txt = 'Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1498, 1500, 1502}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1320, baseId = 1318, txt = 'ImportNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1322, baseId = 1318, txt = 'RootNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1324, baseId = 1318, txt = 'RefTypeNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1326, baseId = 1318, txt = 'IfNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1328, baseId = 1318, txt = 'SwitchNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1330, baseId = 1318, txt = 'WhileNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1332, baseId = 1318, txt = 'RepeatNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1334, baseId = 1318, txt = 'ForNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1336, baseId = 1318, txt = 'ApplyNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1338, baseId = 1318, txt = 'ForeachNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1340, baseId = 1318, txt = 'ForsortNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1342, baseId = 1318, txt = 'ReturnNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1344, baseId = 1318, txt = 'BreakNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1346, baseId = 1318, txt = 'ExpNewNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1348, baseId = 1318, txt = 'ExpListNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1350, baseId = 1318, txt = 'ExpRefNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1352, baseId = 1318, txt = 'ExpOp2Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1354, baseId = 1318, txt = 'ExpCastNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1356, baseId = 1318, txt = 'ExpOp1Node',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1358, baseId = 1318, txt = 'ExpRefItemNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1360, baseId = 1318, txt = 'ExpCallNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1362, baseId = 1318, txt = 'ExpDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1364, baseId = 1318, txt = 'ExpParenNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1366, baseId = 1318, txt = 'BlockNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1368, baseId = 1318, txt = 'StmtExpNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1370, baseId = 1318, txt = 'RefFieldNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1372, baseId = 1318, txt = 'DeclVarNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1374, baseId = 1318, txt = 'DeclFuncNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1376, baseId = 1318, txt = 'DeclMethodNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1378, baseId = 1318, txt = 'DeclConstrNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1380, baseId = 1318, txt = 'DeclMemberNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1382, baseId = 1318, txt = 'DeclArgNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1384, baseId = 1318, txt = 'DeclArgDDDNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1386, baseId = 1318, txt = 'DeclClassNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1388, baseId = 1318, txt = 'LiteralNilNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1390, baseId = 1318, txt = 'LiteralCharNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1392, baseId = 1318, txt = 'LiteralIntNode',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1394, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1504, 1506}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1396, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {18}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1400, baseId = 1, txt = 'getNodeKindName',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1402, baseId = 1, txt = 'nodeFilter',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1404, baseId = 1, txt = 'DeclFuncInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1508, 1510, 1512, 1514, 1516, 1518, 1520, 1522}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1406, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1318}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1408, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1410, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1302}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1412, baseId = 1, txt = 'LiteralStringInfo',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1524, 1526}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1414, baseId = 1, txt = '',
staticFlag = false, accessMode = 'pub', kind = 2, itemTypeId = {1318}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1528, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1530, baseId = 1, txt = 'TransUnit',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1532, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1534, baseId = 1, txt = 'Position',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1536, baseId = 1, txt = 'Token',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1538, baseId = 1, txt = 'Parser',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {1552, 1554, 1556}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1540, baseId = 1, txt = 'Stream',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1544, baseId = 1, txt = 'getKindTxt',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1546, baseId = 1, txt = 'isOp2',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1548, baseId = 1, txt = 'isOp1',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1550, baseId = 1, txt = 'getEofToken',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1, typeId = 1558, baseId = 1, txt = 'filterObj',
staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 270, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 272, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {252}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 252, typeId = 328, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 360, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 362, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {346}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 346, typeId = 364, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 406, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 408, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 410, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 412, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 418, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 420, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 422, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 428, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 430, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 434, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 438, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 440, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 444, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 450, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 458, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {456}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 462, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {460}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 464, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 466, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 468, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 470, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 472, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 474, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 476, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 478, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 480, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 482, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 386, typeId = 486, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {484}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 504, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 506, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 508, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {488}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 510, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 512, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 514, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {488}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 518, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {516}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 488, typeId = 522, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {520}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 532, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 534, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {386}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 528, typeId = 538, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 616, typeId = 696, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {694}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 616, typeId = 1232, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1084, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1086, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1090, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1088}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1092, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1094, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1096, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {528}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1100, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1098}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1076, typeId = 1104, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1102}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1190, typeId = 1194, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {344}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1190, typeId = 1198, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1196}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1416, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1418, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1262}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1262, typeId = 1420, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1422, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1424, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1282}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1282, typeId = 1426, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1428, baseId = 1, txt = 'getParentId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1430, baseId = 1, txt = 'get_baseId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1432, baseId = 1, txt = 'addChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1434, baseId = 1, txt = 'serialize',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1436, baseId = 1, txt = 'getTxt',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1438, baseId = 1, txt = 'equals',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1440, baseId = 1, txt = 'cloneToPublic',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1442, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1444, baseId = 1, txt = 'createBuiltin',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1446, baseId = 1, txt = 'createList',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1448, baseId = 1, txt = 'createArray',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1450, baseId = 1, txt = 'createMap',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1452, baseId = 1, txt = 'createClass',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1454, baseId = 1, txt = 'createFunc',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1456, baseId = 1, txt = 'get_itemTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1304}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1458, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1306}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1460, baseId = 1, txt = 'get_parentInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1462, baseId = 1, txt = 'get_typeId',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1464, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1466, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1468, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1470, baseId = 1, txt = 'get_autoFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1472, baseId = 1, txt = 'get_orgTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1474, baseId = 1, txt = 'get_baseTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1476, baseId = 1, txt = 'get_nilable',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1478, baseId = 1, txt = 'get_nilableTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1302, typeId = 1480, baseId = 1, txt = 'get_children',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1308}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1482, baseId = 1, txt = 'add',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1484, baseId = 1, txt = 'addClass',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1486, baseId = 1, txt = 'getClassScope',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1310}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1488, baseId = 1, txt = 'getTypeInfoChild',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1490, baseId = 1, txt = 'getTypeInfo',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1492, baseId = 1, txt = 'get_parent',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1310}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1494, baseId = 1, txt = 'get_symbol2TypeInfoMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1312}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1310, typeId = 1496, baseId = 1, txt = 'get_className2ScopeMap',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1314}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1498, baseId = 1, txt = 'get_kind',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {12}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1500, baseId = 1, txt = 'get_expType',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1302}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1318, typeId = 1502, baseId = 1, txt = 'get_info',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {6}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1394, typeId = 1504, baseId = 1, txt = 'get_errMessList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1396}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1394, typeId = 1506, baseId = 1, txt = 'createAST',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {20}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1508, baseId = 1, txt = 'get_className',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1510, baseId = 1, txt = 'get_name',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1512, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1406}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1514, baseId = 1, txt = 'get_staticFlag',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {10}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1516, baseId = 1, txt = 'get_accessMode',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1518, baseId = 1, txt = 'get_body',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1318}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1520, baseId = 1, txt = 'get_retTypeList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1408}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1404, typeId = 1522, baseId = 1, txt = 'get_retTypeInfoList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1410}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1412, typeId = 1524, baseId = 1, txt = 'get_token',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1280}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1412, typeId = 1526, baseId = 1, txt = 'get_argList',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1414}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1538, typeId = 1552, baseId = 1, txt = 'getStreamName',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1538, typeId = 1554, baseId = 1, txt = 'create',
staticFlag = true, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {1538}, children = {}, }
)
table.insert( 
moduleObj._typeInfoList, { parentId = 1538, typeId = 1556, baseId = 1, txt = 'getToken',
staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {18}, children = {}, }
)
----- meta -----
return moduleObj
