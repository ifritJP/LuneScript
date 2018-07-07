--lune/base/convLua.lns
local moduleObj = {}
local TransUnit = require( 'lune.base.TransUnit' )

local filterObj = {}
moduleObj.filterObj = filterObj
function filterObj.new( streamName, stream, exeFlag )
  local obj = {}
  setmetatable( obj, { __index = filterObj } )
  return obj.__init and obj:__init( streamName, stream, exeFlag ) or nil;
end
function filterObj:__init(streamName, stream, exeFlag) 
  self.streamName = streamName
  self.stream = stream
  self.moduleName2Info = {}
  self.exeFlag = exeFlag
  self.indent = 0
  self.curLineNo = 1
  self.className2Scope = {}
  self.className2MemberList = {}
  self.pubVarName2InfoMap = {}
  self.pubFuncName2InfoMap = {}
  self.needIndent = false
  return self
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
  local childlen = node.info.childlen
  for __index, child in pairs( childlen ) do
    TransUnit.nodeFilter( child, self, node, baseIndent )
    self:writeln( "", baseIndent )
  end
  self:writeln( "----- meta -----", baseIndent )
  local typeId2TypeInfo = {}
  local typeId2VarInfo = {}
  local function pickupTypeId( typeInfo )
    if not typeId2TypeInfo[typeInfo:get_typeId(  )] then
      typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
      local typeInfoList = typeInfo:get_itemTypeInfoList(  )
      for __index, itemTypeInfo in pairs( typeInfoList ) do
        pickupTypeId( itemTypeInfo )
      end
      typeInfoList = typeInfo:get_retTypeInfoList(  )
      for __index, itemTypeInfo in pairs( typeInfoList ) do
        pickupTypeId( itemTypeInfo )
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
        local work = scope
        do
          local __sorted = {}
          local __map = work:get_symbol2TypeInfoMap(  )
          for __key in pairs( __map ) do
            table.insert( __sorted, __key )
          end
          table.sort( __sorted )
          for __index, declName in ipairs( __sorted ) do
            typeInfo = __map[ declName ]
            do
              if typeInfo.accessMode == "pub" and not typeInfo.externalFlag then
                if typeInfo.kind == TransUnit.TypeInfoKindFunc then
                  local fieldName = declName
                  self:writeln( string.format( "_classInfo%s.%s = {", className, fieldName), baseIndent )
                  self:writeln( string.format( "  name='%s', staticFlag = %s, ", fieldName, typeInfo.staticFlag) .. string.format( "accessMode = '%s', methodFlag = true, typeId = %d }", typeInfo.accessMode, typeInfo:get_typeId(  )), baseIndent )
                  pickupTypeId( typeInfo )
                end
              end
            end
          end
        end
        
        for __index, memberNode in pairs( self.className2MemberList[className] ) do
          local memberInfo = memberNode:get_info(  )
          if memberInfo.accessMode == "pub" then
            local memberName = memberInfo.name.txt
            local memberTypeInfo = memberName.expType
            self:writeln( string.format( "_classInfo%s.%s = {", className, memberName), baseIndent )
            self:writeln( string.format( "  name='%s', staticFlag = %s, ", memberName, memberInfo.staticFlag) .. string.format( "accessMode = '%s', methodFlag = false, typeId = %d }", memberInfo.accessMode, memberNode:get_expType(  ).typeId), baseIndent )
            pickupTypeId( memberInfo.refType )
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
        pickupTypeId( varInfo["typeInfo"] )
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
        self:writeln( string.format( "_funcName2InfoMap.%s = {", funcName ), baseIndent )
        self:writeln( string.format( "  accessMode = '%s', typeId = %d }", funcInfo["accessMode"], funcInfo["typeInfo"]:get_typeId(  )), baseIndent )
        pickupTypeId( funcInfo["typeInfo"] )
      end
    end
  end
  
  self:writeln( "moduleObj._typeInfoList = {", baseIndent )
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
        self:write( typeInfo:serialize(  ) )
        self:writeln( ",", baseIndent + 2 )
      end
    end
  end
  
  self:writeln( "}", baseIndent )
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
  self.className2Scope[className] = node.info.scope
  self.className2MemberList[className] = node.info.memberList
  self:writeln( string.format( "local %s = {}", className ), baseIndent )
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
  return obj.__init and obj:__init( %s ) or nil;
end
function %s:__init( %s )
            ]==], className, argTxt, className, argTxt, className, argTxt), baseIndent )
    for __index, member in pairs( memberList ) do
      local memberName = member["info"].name.txt
      self:writeln( string.format( "self.%s = %s", memberName, memberName ), baseIndent + stepIndent )
    end
    self:writeln( [==[
  return self
end
            ]==], baseIndent )
  end
  local scope = node.info.scope
  for __index, memberNode in pairs( node.info.memberList ) do
    local memberName = memberNode.info.name.txt
    local getterName = "get_" .. memberName
    local typeInfo = scope:getTypeInfo( getterName )
    if memberNode.info.getterMode ~= "none" and (not typeInfo or typeInfo.autoFlag ) then
      self:writeln( string.format( [==[
function %s:%s()
   return self.%s
end]==], className, getterName, memberName), baseIndent )
    end
    local setterName = "set_" .. memberName
    typeInfo = scope:getTypeInfo( setterName )
    if memberNode.info.setterMode ~= "none" and (not typeInfo or typeInfo.autoFlag ) then
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
  self:writeln( string.format( "return obj.__init and obj:__init( %s ) or nil;", argTxt ), baseIndent )
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
  local nameToken = node.info.name
  local name = nameToken and nameToken.txt or ""
  local letTxt = ""
  if node.info.accessMode ~= "global" and #name ~= 0 then
    letTxt = "local "
  end
  self:write( string.format( "%sfunction %s( ", letTxt, name ) )
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
  if node.expType:get_accessMode(  ) == "pub" then
    self:write( string.format( "moduleObj.%s = %s", name, name) )
    self.pubFuncName2InfoMap[name] = {["funcFlag"] = true, ["accessMode"] = node.info.accessMode, ["typeInfo"] = node.expType}
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
  TransUnit.nodeFilter( node.info, self, node, baseIndent )
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
  local txt = node.info.token.txt
  if string.find( txt, '^```' ) then
    txt = '[==[' .. txt:sub( 4, -4 ) .. ']==]'
  end
  if #node.info.argList > 0 then
    self:write( string.format( "string.format( %s, ", txt ) )
    local argList = node.info.argList
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
moduleObj._typeInfoList = {
}
----- meta -----
return moduleObj
