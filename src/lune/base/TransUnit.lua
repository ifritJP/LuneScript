--lune/base/TransUnit.lns
local moduleObj = {}
local Parser = require( 'lune.base.Parser' )

local function debugLog(  )
  local debugInfo1 = debug.getinfo( 2 )
  local debugInfo2 = debug.getinfo( 3 )
  local debugInfo3 = debug.getinfo( 4 )
  print( debugInfo1.short_src, debugInfo1.currentline )
  print( debugInfo2.short_src, debugInfo2.currentline )
  print( debugInfo3.short_src, debugInfo3.currentline )
end

local typeIdSeed = 1
local typeInfo = {}
moduleObj.typeInfo = typeInfo

local builtInTypeMap = {}
local TypeInfoKindPrim = 1
local TypeInfoKindList = 2
local TypeInfoKindArray = 3
local TypeInfoKindMap = 4
local TypeInfoKindClass = 5
local TypeInfoKindFunc = 6
local TypeInfo = {}
function TypeInfo.new( txt, typeId, kind, itemTypeInfo, retTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = TypeInfo } )
  return obj.__init and obj:__init( txt, typeId, kind, itemTypeInfo, retTypeInfo ) or nil;
end
function TypeInfo:__init(txt, typeId, kind, itemTypeInfo, retTypeInfo) 
  self.txt = txt
  self.typeId = typeId
  self.kind = kind
  self.itemTypeInfo = itemTypeInfo or {}
  self.retTypeInfo = retTypeInfo or {}
  return self
end
function TypeInfo:serialize(  )
  local txt = '{ itemTypeId = { '
  for index, typeInfo in pairs( self.itemTypeInfo ) do
    if index ~= 1 then
      txt = txt .. ", "
    end
    txt = string.format( "%s%d", txt, typeInfo.typeId)
  end
  txt = txt .. string.format( '}, typeId = %d, txt = "%s", kind = %d }', self.typeId, self.txt, self.kind)
  return txt
end
function TypeInfo:getItemTypeInfoList(  )
  return self.itemTypeInfo
end
function TypeInfo:getRetTypeInfoList(  )
  return self.retTypeInfo
end
function TypeInfo:getKind(  )
  return self.kind
end
function TypeInfo:getTypeId(  )
  return self.typeId
end
function TypeInfo:getTxt(  )
  if self.kind == TypeInfoKindArray then
    return self.itemTypeInfo[1]:getTxt(  ) .. "[@]"
  end
  if self.kind == TypeInfoKindList then
    return self.itemTypeInfo[1]:getTxt(  ) .. "[]"
  end
  if self.itemTypeInfo and #self.itemTypeInfo > 0 then
    local txt = self.txt .. "<"
    for index, typeInfo in pairs( self.itemTypeInfo ) do
      if index ~= 1 then
        txt = txt .. ","
      end
      txt = txt .. typeInfo:getTxt(  )
    end
    return txt .. ">"
  end
  if self.txt then
    return self.txt
  end
  return ""
end
function TypeInfo.create( kind, txt, itemTypeInfo, retTypeInfo )
  if kind == TypeInfoKindPrim then
    return builtInTypeMap[txt]
  end
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(txt, typeIdSeed, kind, itemTypeInfo, retTypeInfo)
  return info
end
function TypeInfo.createBuiltin( idName, typeTxt )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(typeTxt, typeIdSeed, TypeInfoKindPrim)
  typeInfo[idName] = info
  builtInTypeMap[typeTxt] = info
  return info
end
function TypeInfo.createList( itemTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(nil, typeIdSeed, TypeInfoKindList, itemTypeInfo)
end
function TypeInfo.createArray( itemTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(nil, typeIdSeed, TypeInfoKindArray, itemTypeInfo)
end
function TypeInfo.createMap( keyTypeInfo, valTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new("Map", typeIdSeed, TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end
function TypeInfo.createClass( className )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(className, typeIdSeed, TypeInfoKindClass)
  return info
end
function TypeInfo.createFunc( funcName, argTypeList, retTypeInfo )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(funcName, typeIdSeed, TypeInfoKindFunc, argTypeList, retTypeInfo)
  return info
end

local typeInfoNone = TypeInfo.createBuiltin( "None", "" )
local typeInfoStem = TypeInfo.createBuiltin( "Stem", "stem" )
local typeInfoNil = TypeInfo.createBuiltin( "Nil", "nil" )
local typeInfoBool = TypeInfo.createBuiltin( "Bool", "bool" )
local typeInfoInt = TypeInfo.createBuiltin( "Int", "int" )
local typeInfoReal = TypeInfo.createBuiltin( "Real", "real" )
local typeInfoChar = TypeInfo.createBuiltin( "char", "char" )
local typeInfoString = TypeInfo.createBuiltin( "String", "str" )
local typeInfoMap = TypeInfo.createBuiltin( "Map", "Map" )
local typeInfoList = TypeInfo.createBuiltin( "List", "List" )
local typeInfoArray = TypeInfo.createBuiltin( "Array", "Array" )
local typeInfoForm = TypeInfo.createBuiltin( "Form", "form" )
local Scope = {}
function Scope.new( parent )
  local obj = {}
  setmetatable( obj, { __index = Scope } )
  return obj.__init and obj:__init( parent ) or nil;
end
function Scope:__init(parent) 
  self.parent = parent
  self.symbol2TypeInfoMap = {}
  self.className2ScopeMap = {}
  return self
end
function Scope:add( name, typeInfo )
  self.symbol2TypeInfoMap[name] = typeInfo
end
function Scope:addClass( name, typeInfo, scope )
  self:add( name, typeInfo )
  self.className2ScopeMap[name] = scope
end
function Scope:getClassScope( name )
  local scope = self.className2ScopeMap[name]
  if not scope and self.parent then
    scope = self.parent:getClassScope( name )
  end
  return scope
end
function Scope:getTypeInfoChild( name )
  return self.symbol2TypeInfoMap[name]
end
function Scope:getTypeInfo( name )
  local node = self.symbol2TypeInfoMap[name]
  if node then
    return node
  end
  if self.parent then
    return self.parent:getTypeInfo( name )
  end
  return builtInTypeMap[name]
end
function Scope:getParent(  )
  return self.parent
end

local Node = {}
moduleObj.Node = Node

local TransUnit = {}
moduleObj.TransUnit = TransUnit
function TransUnit.new(  )
  local obj = {}
  setmetatable( obj, { __index = TransUnit } )
  return obj.__init and obj:__init(  ) or nil;
end
function TransUnit:__init() 
  self.scope = Scope.new(nil)
  self.classList = {}
  self.typeInfo2Scope = {}
  return self
end
function TransUnit:pushScope(  )
  self.scope = Scope.new(self.scope)
  return self.scope
end
function TransUnit:popScope(  )
  self.scope = self.scope:getParent(  )
end
function TransUnit:pushClass( name )
  local typeInfo = self.scope:getTypeInfoChild( name )
  if not typeInfo then
    typeInfo = TypeInfo.createClass( name )
    local scope = self:pushScope(  )
    scope:getParent(  ):addClass( name, typeInfo, scope )
  else 
    self.scope = self.scope:getClassScope( name )
  end
  table.insert( self.classList, {["name"] = name, ["scope"] = self.scope, ["typeInfo"] = typeInfo} )
  self.typeInfo2Scope[typeInfo] = self.scope
  return self.scope
end
function TransUnit:popClass(  )
  self:popScope(  )
  table.remove( self.classList )
end
-- none
-- none
-- none
-- none
-- none
-- none
-- none
-- none

local nodeKind2NameMap = {}
local nodeKindSeed = 1
local nodeKind = {}
moduleObj.nodeKind = nodeKind

local function regKind( name )
  local kind = nodeKindSeed
  nodeKindSeed = nodeKindSeed + 1
  nodeKind2NameMap[kind] = name
  nodeKind[name] = kind
  return kind
end

local nodeKindNone = regKind( 'None' )
local nodeKindImport = regKind( 'Import' )
local nodeKindRoot = regKind( 'Root' )
local nodeKindRefType = regKind( 'RefType' )
local nodeKindIf = regKind( 'If' )
local nodeKindWhile = regKind( 'While' )
local nodeKindRepeat = regKind( 'Repeat' )
local nodeKindFor = regKind( 'For' )
local nodeKindApply = regKind( 'Apply' )
local nodeKindForeach = regKind( 'Foreach' )
local nodeKindForsort = regKind( 'Forsort' )
local nodeKindReturn = regKind( 'Return' )
local nodeKindBreak = regKind( 'Break' )
local nodeKindExpNew = regKind( 'ExpNew' )
local nodeKindExpList = regKind( 'ExpList' )
local nodeKindExpRef = regKind( 'ExpRef' )
local nodeKindExpOp2 = regKind( 'ExpOp2' )
local nodeKindExpCast = regKind( 'ExpCast' )
local nodeKindExpOp1 = regKind( 'ExpOp1' )
local nodeKindExpRefItem = regKind( 'ExpRefItem' )
local nodeKindExpCall = regKind( 'ExpCall' )
local nodeKindExpDDD = regKind( 'ExpDDD' )
local nodeKindExpParen = regKind( 'ExpParen' )
local nodeKindBlock = regKind( 'Block' )
local nodeKindStmtExp = regKind( 'StmtExp' )
local nodeKindRefField = regKind( 'RefField' )
local nodeKindDeclVar = regKind( 'DeclVar' )
local nodeKindDeclFunc = regKind( 'DeclFunc' )
local nodeKindDeclMethod = regKind( 'DeclMethod' )
local nodeKindDeclConstr = regKind( 'DeclConstr' )
local nodeKindDeclMember = regKind( 'DeclMember' )
local nodeKindDeclArg = regKind( 'DeclArg' )
local nodeKindDeclArgDDD = regKind( 'DeclArgDDD' )
local nodeKindDeclClass = regKind( 'DeclClass' )
local nodeKindLiteralNil = regKind( 'LiteralNil' )
local nodeKindLiteralChar = regKind( 'LiteralChar' )
local nodeKindLiteralInt = regKind( 'LiteralInt' )
local nodeKindLiteralReal = regKind( 'LiteralReal' )
local nodeKindLiteralArray = regKind( 'LiteralArray' )
local nodeKindLiteralList = regKind( 'LiteralList' )
local nodeKindLiteralMap = regKind( 'LiteralMap' )
local nodeKindLiteralString = regKind( 'LiteralString' )
local nodeKindLiteralBool = regKind( 'LiteralBool' )
local quotedChar2Code = {}
quotedChar2Code['a'] = 7
quotedChar2Code['b'] = 8
quotedChar2Code['t'] = 9
quotedChar2Code['n'] = 10
quotedChar2Code['v'] = 11
quotedChar2Code['f'] = 12
quotedChar2Code['r'] = 13
quotedChar2Code['\\'] = 92
quotedChar2Code['"'] = 34
quotedChar2Code["'"] = 39
local function nodeFilter( node, filter, ... )
  if not filter[node.kind] then
    error( string.format( "none filter -- %s", getNodeKindName( node.kind ) ) )
  end
  return filter[node.kind]( filter, node, ... )
end
moduleObj.nodeFilter = nodeFilter
local function getNodeKindName( kind )
  return nodeKind2NameMap[kind]
end
moduleObj.getNodeKindName = getNodeKindName
function TransUnit:registBuiltInScope(  )
  local builtInInfo = {["io"] = {"open"}, ["string"] = {"find", "byte", "format", "rep", "gmatch", "gsub"}, ["table"] = {"insert", "remove"}, ["debug"] = {"getinfo"}, ["_luneScript"] = {"loadModule"}}
  for name, infoList in pairs( builtInInfo ) do
    local childScope = self:pushClass( name )
    for __index, info in pairs( infoList ) do
      childScope:add( info, TypeInfo.createFunc( info ) )
    end
    self:popClass(  )
  end
end

function TransUnit:createNode( kind, pos, expType, info )
  if not getNodeKindName( kind ) then
    error( string.format( "%d:%d: not found nodeKind", pos.lineNo, pos.column ) )
  end
  return {["kind"] = kind, ["pos"] = pos, ["expType"] = expType, ["info"] = info, ["filter"] = nodeFilter}
end

function TransUnit:error( mess )
  local pos = {["lineNo"] = 0, ["column"] = 0}
  local txt = ""
  if self.currentToken then
    pos = self.currentToken.pos
    txt = self.currentToken.txt
  end
  error( string.format( "%d:%d:(%s) %s", pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:createNoneNode(  )
  return self:createNode( nodeKindNone, {["lineNo"] = 0, ["column"] = 0}, typeInfoNone, {} )
end

function TransUnit:getTokenNoErr(  )
  if self.pushbackToken then
    self.currentToken = self.pushbackToken
    self.pushbackToken = nil
    return self.currentToken
  end
  local commentList = {}
  local token = nil
  while true do
    token = self.parser:getToken(  )
    if not token then
      break
    end
    if token.kind ~= Parser.kind.Cmnt then
      break
    end
    table.insert( commentList, token )
  end
  if token then
    token.commentList = commentList
  end
  self.currentToken = token
  return token
end

function TransUnit:getToken( mess )
  local token = self:getTokenNoErr(  )
  if not token then
    return Parser.getEofToken(  )
  end
  self.currentToken = token
  return self.currentToken
end

function TransUnit:pushback(  )
  if self.pushbackToken then
    error( string.format( "multiple pushback:%d:%d: %s, %s", self.currentToken.pos.lineNo, self.currentToken.pos.column, self.pushbackToken.txt, self.currentToken.txt ) )
  end
  self.pushbackToken = self.currentToken
  self.currentToken = nil
end

function TransUnit:checkSymbol( token )
  if token.kind ~= Parser.kind.Symb and token.kind ~= Parser.kind.Kywd and token.kind ~= Parser.kind.Type then
    self:error( "illegal symbol" )
  end
  return token
end

function TransUnit:getSymbolToken(  )
  return self:checkSymbol( self:getToken(  ) )
end

function TransUnit:checkToken( token, txt )
  if not token or token.txt ~= txt then
    self:error( string.format( "not found -- %s", txt ) )
  end
  return token
end

function TransUnit:checkNextToken( txt )
  return self:checkToken( self:getToken(  ), txt )
end

function TransUnit:analyzeBlock( blockKind, scope )
  local token = self:checkNextToken( "{" )
  if not scope then
    self:pushScope(  )
  end
  local stmtList = {}
  self:analyzeStatement( stmtList, "}" )
  self:checkNextToken( "}" )
  if not scope then
    self:popScope(  )
  end
  local node = self:createNode( nodeKindBlock, token.pos, typeInfoNone, {["kind"] = blockKind, ["stmtList"] = stmtList} )
  return node
end

function TransUnit:analyzeDecl( accessMode, staticFlag, firstToken, token )
  local staticFlag = false
  if not staticFlag then
    if token.txt == "static" then
      staticFlag = true
      token = self:getToken(  )
    end
  end
  if token.txt == "let" then
    return self:analyzeDeclVar( accessMode, staticFlag, firstToken )
  elseif token.txt == "fn" then
    return self:analyzeDeclFunc( accessMode, staticFlag, false, token, nil )
  elseif token.txt == "class" then
    return self:analyzeDeclClass( accessMode, token )
  end
  return nil
end

function TransUnit:analyzeImport( token )
  local moduleToken = self:getToken(  )
  local modulePath = moduleToken.txt
  local nextToken = {}
  local nameList = {moduleToken.txt}
  while true do
    nextToken = self:getToken(  )
    if nextToken.txt == "." then
      nextToken = self:getToken(  )
      moduleToken = nextToken
      modulePath = string.format( "%s.%s", modulePath, moduleToken.txt)
      table.insert( nameList, moduleToken.txt )
    else 
      break
    end
  end
  self:pushClass( moduleToken.txt )
  local moduleInfo = _luneScript.loadModule( modulePath )
  self.moduleName2Info[modulePath] = moduleInfo
  for className, classInfo in pairs( moduleInfo._className2InfoMap ) do
    self:pushClass( className )
    for fieldName, fieldInfo in pairs( classInfo ) do
      self.scope:add( fieldName, TypeInfo.createFunc( fieldName ) )
    end
    self:popClass(  )
  end
  local typeId2TypeInfo = {}
  for __index, typeInfo in pairs( moduleInfo._typeInfoList ) do
    local itemTypeInfo = {}
    for __index, typeId in pairs( typeInfo.itemTypeId ) do
      table.insert( itemTypeInfo, typeId2TypeInfo[typeId] )
    end
    typeId2TypeInfo[typeInfo.typeId] = TypeInfo.create( typeInfo.typeId, typeInfo.txt, itemTypeInfo )
  end
  for varName, varInfo in pairs( moduleInfo._varName2InfoMap ) do
    self.scope:add( varName, typeId2TypeInfo[varInfo["typeId"]] )
  end
  for funcName, funcInfo in pairs( moduleInfo._funcName2InfoMap ) do
    self.scope:add( funcName, typeId2TypeInfo[funcInfo["typeId"]] )
  end
  self:popClass(  )
  self:checkToken( nextToken, ";" )
  return self:createNode( nodeKindImport, token.pos, typeInfoNone, modulePath )
end

function TransUnit:analyzeIf( token )
  local list = {}
  table.insert( list, {["kind"] = "if", ["exp"] = self:analyzeExp(  ), ["block"] = self:analyzeBlock( "if" )} )
  local nextToken = self:getToken(  )
  if nextToken.txt == "elseif" then
    while nextToken.txt == "elseif" do
      table.insert( list, {["kind"] = "elseif", ["exp"] = self:analyzeExp(  ), ["block"] = self:analyzeBlock( "elseif" )} )
      nextToken = self:getToken(  )
    end
  end
  if nextToken.txt == "else" then
    table.insert( list, {["kind"] = "else", ["block"] = self:analyzeBlock( "else" )} )
  else 
    self:pushback(  )
  end
  return self:createNode( nodeKindIf, token.pos, typeInfoNone, list )
end

function TransUnit:analyzeWhile( token )
  local info = {["exp"] = self:analyzeExp(  ), ["block"] = self:analyzeBlock( "while" )}
  return self:createNode( nodeKindWhile, token.pos, typeInfoNone, info )
end

function TransUnit:analyzeRepeat( token )
  local scope = self:pushScope(  )
  local info = {["block"] = self:analyzeBlock( "repeat", scope ), ["exp"] = self:analyzeExp(  )}
  self:popScope(  )
  local node = self:createNode( nodeKindRepeat, token.pos, typeInfoNone, info )
  self:checkNextToken( ";" )
  return node
end

function TransUnit:analyzeFor( token )
  local scope = self:pushScope(  )
  local val = self:getToken(  )
  if val.kind ~= Parser.kind.Symb then
    self:error( "not symbol" )
  end
  self:checkNextToken( "=" )
  local exp1 = self:analyzeExp(  )
  self:checkNextToken( "," )
  local exp2 = self:analyzeExp(  )
  local token = self:getToken(  )
  local exp3 = nil
  if token.txt == "," then
    exp3 = self:analyzeExp(  )
  else 
    self:pushback(  )
  end
  local info = {["block"] = self:analyzeBlock( "for", scope ), ["val"] = val, ["init"] = exp1, ["to"] = exp2, ["delta"] = exp3}
  self:popScope(  )
  local node = self:createNode( nodeKindFor, token.pos, typeInfoNone, info )
  return node
end

function TransUnit:analyzeApply( token )
  local scope = self:pushScope(  )
  local varList = {}
  local nextToken = nil
  repeat 
    local var = self:getToken(  )
    if var.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    table.insert( varList, var )
    nextToken = self:getToken(  )
  until nextToken.txt ~= ","
  self:checkToken( nextToken, "of" )
  local exp = self:analyzeExp(  )
  if exp.kind ~= nodeKindExpCall then
    self:error( "not call" )
  end
  local block = self:analyzeBlock( "apply", scope )
  self:popScope(  )
  local info = {["varList"] = varList, ["exp"] = exp, ["block"] = block}
  return self:createNode( nodeKindApply, token.pos, typeInfoNone, info )
end

function TransUnit:analyzeForeach( token, sortFlag )
  local scope = self:pushScope(  )
  local valSymbol = nil
  local keySymbol = nil
  local nextToken = nil
  for index = 1, 2 do
    local sym = self:getToken(  )
    if sym.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    if index == 1 then
      valSymbol = sym
    else 
      keySymbol = sym
    end
    nextToken = self:getToken(  )
    if nextToken.txt ~= "," then
      break
    end
  end
  self:checkToken( nextToken, "in" )
  local exp = self:analyzeExp(  )
  if not exp.expType then
  else 
    local itemTypeInfoList = exp.expType:getItemTypeInfoList(  )
    if exp.expType:getKind(  ) == TypeInfoKindMap then
      self.scope:add( valSymbol.txt, itemTypeInfoList[2] )
      if keySymbol then
        self.scope:add( keySymbol.txt, itemTypeInfoList[1] )
      end
    elseif exp.expType:getKind(  ) == TypeInfoKindList or exp.expType:getKind(  ) == TypeInfoKindArray then
      self.scope:add( valSymbol.txt, itemTypeInfoList[1] )
      if keySymbol then
        self.scope:add( "__index", typeInfoInt )
      end
    end
  end
  local block = self:analyzeBlock( "foreach", scope )
  self:popScope(  )
  local info = {["val"] = valSymbol, ["key"] = keySymbol, ["exp"] = exp, ["block"] = block, ["sort"] = sortFlag}
  return self:createNode( sortFlag and nodeKindForsort or nodeKindForeach, token.pos, typeInfoNone, info )
end

function TransUnit:analyzeRefType(  )
  local firstToken = self:getToken(  )
  local token = firstToken
  local refFlag = false
  if token.txt == "&" then
    refFlag = true
    token = self:getToken(  )
  end
  local mutFlag = false
  if token.txt == "mut" then
    mutFlag = true
    token = self:getToken(  )
  end
  local name = self:analyzeExpSymbol( firstToken, token, "symbol", token, true )
  local arrayMode = "no"
  token = self:getToken(  )
  if token.txt == '[' or token.txt == '[@' then
    if token.txt == '[' then
      arrayMode = "list"
    else 
      arrayMode = "array"
    end
    token = self:getToken(  )
    if token.txt ~= ']' then
      self:pushback(  )
      self:checkNextToken( ']' )
    end
  elseif token.txt == "<" then
    local nextToken = nil
    repeat 
      self:getSymbolToken(  )
      nextToken = self:getToken(  )
    until nextToken.txt ~= ","
    self:checkToken( nextToken, '>' )
  else 
    self:pushback(  )
  end
  return self:createNode( nodeKindRefType, firstToken.pos, name.expType, {["name"] = name, ["refFlag"] = refFlag, ["mutFlag"] = mutFlag, ["array"] = arrayMode} )
end

function TransUnit:analyzeDeclMember( accessMode, staticFlag, firstToken )
  local varName = self:getSymbolToken(  )
  token = self:getToken(  )
  local refType = self:analyzeRefType(  )
  token = self:getToken(  )
  self:checkToken( token, ";" )
  self.scope:add( varName.txt, refType.expType )
  return self:createNode( nodeKindDeclMember, firstToken.pos, refType.expType, {["name"] = varName, ["refType"] = refType, ["staticFlag"] = staticFlag, ["accessMode"] = accessMode} )
end

function TransUnit:analyzeDeclMethod( accessMode, staticFlag, className, firstToken, name )
  local node = self:analyzeDeclFunc( accessMode, staticFlag, true, name, name )
  node.info.className = className
  return node
end

function TransUnit:analyzeDeclClass( classAccessMode, classToken )
  local name = self:getToken(  )
  self:checkNextToken( "{" )
  local fieldList = {}
  local typeInfo = TypeInfo.createClass( name )
  local node = self:createNode( nodeKindDeclClass, classToken.pos, typeInfo, {["accessMode"] = classAccessMode, ["name"] = name, ["fieldList"] = fieldList} )
  self:pushClass( name.txt )
  while true do
    local token = self:getToken(  )
    if token.txt == "}" then
      break
    end
    local accessMode = "pri"
    if token.txt == "pub" or token.txt == "pro" or token.txt == "pri" or token.txt == "global" then
      accessMode = token.txt
      token = self:getToken(  )
    end
    local staticFlag = false
    if token.txt == "static" then
      staticFlag = true
      token = self:getToken(  )
    end
    if token.txt == "let" then
      table.insert( fieldList, self:analyzeDeclMember( accessMode, staticFlag, token ) )
    else 
      table.insert( fieldList, self:analyzeDeclMethod( accessMode, staticFlag, name, token, token ) )
    end
  end
  self:popClass(  )
  return node
end

function TransUnit:analyzeDeclFunc( accessMode, staticFlag, methodFlag, firstToken, name )
  local argList = {}
  local token = self:getToken(  )
  if not name then
    if token.txt ~= "(" then
      name = self:checkSymbol( token )
      token = self:getToken(  )
    end
  else 
    self:checkSymbol( name )
  end
  local className = nil
  if token.txt == "." then
    methodFlag = true
    className = name
    self:pushClass( name.txt )
    name = self:getSymbolToken(  )
    token = self:getToken(  )
  end
  self:checkToken( "(" )
  local kind = nodeKindDeclConstr
  if methodFlag then
    if name.txt == "__init" then
      kind = nodeKindDeclConstr
    else 
      kind = nodeKindDeclMethod
    end
  else 
    kind = nodeKindDeclFunc
  end
  local scope = self:pushScope(  )
  repeat 
    local argName = self:getToken(  )
    if argName.txt == ")" then
      token = argName
      break
    elseif argName.txt == "..." then
      table.insert( argList, self:createNode( nodeKindDeclArgDDD, argName.pos, typeInfoNone, argName ) )
    else 
      self:checkSymbol( argName )
      self:checkNextToken( ":" )
      local refType = self:analyzeRefType(  )
      local arg = self:createNode( nodeKindDeclArg, argName.pos, refType.expType, {["name"] = argName, ["argType"] = refType} )
      self.scope:add( argName.txt, refType.expType )
      table.insert( argList, arg )
    end
    token = self:getToken(  )
  until token.txt ~= ","
  self:checkToken( token, ")" )
  token = self:getToken(  )
  local typeList = {}
  local retTypeList = {}
  if token.txt == ":" then
    repeat 
      local refType = self:analyzeRefType(  )
      table.insert( typeList, refType )
      table.insert( retTypeList, refType.expType )
      token = self:getToken(  )
    until token.txt ~= ","
  end
  local typeInfo = TypeInfo.createFunc( name and name.txt or "", nil, retTypeList )
  if name then
    scope:getParent(  ):add( name.txt, typeInfo )
  end
  if token.txt == ";" then
    return self:createNoneNode(  )
  end
  self:pushback(  )
  local body = self:analyzeBlock( "func", scope )
  self:popScope(  )
  local info = {["name"] = name, ["argList"] = argList, ["staticFlag"] = staticFlag, ["body"] = body, ["accessMode"] = accessMode, ["retTypeList"] = retTypeList}
  local node = self:createNode( kind, firstToken.pos, typeInfo, info )
  if className then
    info.className = className
    self:popClass(  )
  end
  return node
end

function TransUnit:analyzeDeclVar( accessMode, staticFlag, firstToken )
  local typeInfoList = {}
  local varList = {}
  local token = nil
  repeat 
    local varName = self:getSymbolToken(  )
    token = self:getToken(  )
    local refType = typeInfoNone
    if token.txt == ":" then
      refType = self:analyzeRefType(  )
      token = self:getToken(  )
    end
    table.insert( varList, {["name"] = varName, ["refType"] = refType} )
    table.insert( typeInfoList, refType )
  until token.txt ~= ","
  local expList = nil
  if token.txt == "=" then
    expList = self:analyzeExpList(  )
  end
  for index, exp in pairs( expList.info ) do
    typeInfoList[index] = exp["expType"]
  end
  self:checkNextToken( ";" )
  local declVarInfo = {["accessMode"] = accessMode, ["varList"] = varList, ["expList"] = expList, ["typeInfoList"] = typeInfoList}
  local node = self:createNode( nodeKindDeclVar, firstToken.pos, typeInfoNone, declVarInfo )
  for index, typeInfo in pairs( typeInfoList ) do
    self.scope:add( varList[index].name.txt, typeInfo )
  end
  return node
end

function TransUnit:analyzeExpList(  )
  local expList = {}
  local firstExp = nil
  repeat 
    local exp = self:analyzeExp(  )
    if not firstExp then
      firstExp = exp
    end
    table.insert( expList, exp )
    local token = self:getToken(  )
  until token.txt ~= ","
  self:pushback(  )
  return self:createNode( nodeKindExpList, firstExp.pos, typeInfoNone, expList )
end

function TransUnit:analyzeListConst( token )
  local nextToken = self:getToken(  )
  local expList = nil
  if nextToken.txt ~= "]" then
    self:pushback(  )
    expList = self:analyzeExpList(  )
    self:checkNextToken( "]" )
  end
  local kind = nodeKindLiteralArray
  if token.txt == '[' then
    kind = nodeKindLiteralList
  end
  return self:createNode( kind, token.pos, typeInfoList, expList )
end

function TransUnit:analyzeMapConst( token )
  local nextToken = nil
  local map = {}
  local pairList = {}
  local keyTypeInfo = typeInfoNone
  local valTypeInfo = typeInfoNone
  repeat 
    nextToken = self:getToken(  )
    if nextToken.txt == "}" then
      break
    end
    self:pushback(  )
    local key = self:analyzeExp(  )
    if key.expType ~= keyTypeInfo then
      if keyTypeInfo ~= typeInfoNone then
        keyTypeInfo = typeInfoStem
      else 
        keyTypeInfo = key.expType
      end
    end
    self:checkNextToken( ":" )
    local val = self:analyzeExp(  )
    if val.expType ~= valTypeInfo then
      if valTypeInfo ~= typeInfoNone then
        valTypeInfo = typeInfoStem
      else 
        valTypeInfo = val.expType
      end
    end
    table.insert( pairList, {["key"] = key, ["val"] = val} )
    map[key] = val
    nextToken = self:getToken(  )
  until nextToken.txt ~= ","
  local typeInfo = TypeInfo.createMap( keyTypeInfo, valTypeInfo )
  self:checkToken( nextToken, "}" )
  return self:createNode( nodeKindLiteralMap, token.pos, typeInfo, {["map"] = map, ["pairList"] = pairList} )
end

function TransUnit:analyzeExpRefItem( token, exp )
  local indexExp = self:analyzeExp(  )
  self:checkNextToken( "]" )
  local info = {["val"] = exp, ["index"] = indexExp}
  return self:createNode( nodeKindExpRefItem, token.pos, typeInfoNone, info )
end

function TransUnit:analyzeExpCont( firstToken, exp, skipFlag )
  local nextToken = self:getToken(  )
  if not skipFlag then
    repeat 
      local matchFlag = false
      if nextToken.txt == "[" then
        matchFlag = true
        exp = self:analyzeExpRefItem( nextToken, exp )
        nextToken = self:getToken(  )
      end
      if nextToken.txt == "(" then
        matchFlag = true
        local work = self:getToken(  )
        local expList = nil
        if work.txt ~= ")" then
          self:pushback(  )
          expList = self:analyzeExpList(  )
          self:checkNextToken( ")" )
        end
        local info = {["func"] = exp, ["argList"] = expList}
        exp = self:createNode( nodeKindExpCall, firstToken.pos, typeInfoNone, info )
        nextToken = self:getToken(  )
      end
    until not matchFlag
  end
  if nextToken.txt == "." then
    return self:analyzeExpSymbol( firstToken, self:getToken(  ), "field", exp, skipFlag )
  end
  self:pushback(  )
  return exp
end

function TransUnit:analyzeExpSymbol( firstToken, token, mode, prefixExp, skipFlag )
  local exp = nil
  if mode == "field" then
    local info = {["field"] = token, ["prefix"] = prefixExp}
    local typeInfo = typeInfoNone
    if not prefixExp.expType then
      self:error( "unknown prefix type: " .. getNodeKindName( prefixExp.kind ) )
    end
    if prefixExp.expType:getKind(  ) == TypeInfoKindClass then
      local classScope = self.typeInfo2Scope[prefixExp.expType]
      local className = prefixExp.expType.txt
      if not classScope then
        self:error( "not found field: " .. className )
      end
      typeInfo = classScope:getTypeInfo( token.txt )
      if not typeInfo then
        self:error( string.format( "not found field typeInfo: %s.%s %s", className, token.txt, classScope ) )
      end
    end
    exp = self:createNode( nodeKindRefField, firstToken.pos, typeInfo, info )
  elseif mode == "symbol" then
    local typeInfo = self.scope:getTypeInfo( token.txt )
    if not typeInfo and token.txt == "self" then
      local classInfo = self.classList[#self.classList]
      typeInfo = classInfo.typeInfo
    end
    exp = self:createNode( nodeKindExpRef, firstToken.pos, typeInfo, token )
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( "pri", false, false, token, nil )
  else 
    self:error( "illegal mode", mode )
  end
  return self:analyzeExpCont( firstToken, exp, skipFlag )
end

function TransUnit:analyzeExpOp2( firstToken, exp )
  local nextToken = self:getToken(  )
  while true do
    if nextToken.txt == "@" then
      local castType = self:analyzeRefType(  )
      local info = {["exp"] = exp, ["castType"] = castType}
      exp = self:createNode( nodeKindExpCast, firstToken.pos, typeInfoNone, info )
    elseif nextToken.kind == Parser.kind.Ope then
      if Parser.isOp2( nextToken.txt ) then
        local exp2 = self:analyzeExp( (nextToken.txt == "and" ) or (nextToken.txt == "*" ) )
        local info = {["op"] = nextToken, ["exp1"] = exp, ["exp2"] = exp2}
        exp = self:createNode( nodeKindExpOp2, firstToken.pos, typeInfoNone, info )
      else 
        self:error( "illegal op" )
      end
    else 
      self:pushback(  )
      return exp
    end
    nextToken = self:getToken(  )
  end
end

function TransUnit:analyzeExp( skipOp2Flag )
  local firstToken = self:getToken(  )
  local token = firstToken
  local exp = nil
  if token.kind == Parser.kind.Dlmt then
    if token.txt == "..." then
      return self:createNode( nodeKindExpDDD, firstToken.pos, typeInfoNone, token )
    end
    if token.txt == '[' or token.txt == '[@' then
      return self:analyzeListConst( token )
    end
    if token.txt == '{' then
      return self:analyzeMapConst( token )
    end
    if token.txt == "(" then
      exp = self:analyzeExp( false )
      self:checkNextToken( ")" )
      exp = self:createNode( nodeKindExpParen, firstToken.pos, typeInfoNone, exp )
      exp = self:analyzeExpCont( firstToken, exp, false )
    end
  end
  if token.txt == "new" then
    local nextToken = self:getToken(  )
    exp = self:analyzeExpSymbol( firstToken, nextToken, "symbol", nextToken, true )
    self:checkNextToken( "(" )
    nextToken = self:getToken(  )
    local argList = nil
    if nextToken.txt ~= ")" then
      self:pushback(  )
      argList = self:analyzeExpList(  )
      self:checkNextToken( ")" )
    end
    exp = self:createNode( nodeKindExpNew, firstToken.pos, exp.expType, {["symbol"] = exp, ["argList"] = argList} )
    exp = self:analyzeExpCont( firstToken, exp, false )
  end
  if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
    exp = self:analyzeExp( true )
    local typeInfo = typeInfoNone
    if token.txt == "-" then
      typeInfo = exp.expType
    elseif token.txt == "#" then
      typeInfo = typeInfoInt
    elseif token.txt == "not" then
      typeInfo = typeInfoBool
    else 
      self:error( "unknown op1" )
    end
    exp = self:createNode( nodeKindExpOp1, firstToken.pos, typeInfo, {["op"] = token, ["exp"] = exp} )
    return self:analyzeExpOp2( firstToken, exp )
  end
  if token.kind == Parser.kind.Int then
    exp = self:createNode( nodeKindLiteralInt, firstToken.pos, typeInfoInt, {["token"] = token, ["num"] = tonumber( token.txt )} )
  elseif token.kind == Parser.kind.Real then
    exp = self:createNode( nodeKindLiteralReal, firstToken.pos, typeInfoReal, {["token"] = token, ["num"] = tonumber( token.txt )} )
  elseif token.kind == Parser.kind.Char then
    local num = 0
    if #(token.txt ) == 1 then
      num = token.txt:byte( 1 )
    else 
      num = quotedChar2Code[token.txt:sub( 2, 2 )]
    end
    exp = self:createNode( nodeKindLiteralChar, firstToken.pos, typeInfoChar, {["token"] = token, ["num"] = num} )
  elseif token.kind == Parser.kind.Str then
    local nextToken = self:getToken(  )
    local formatArgList = {}
    if nextToken.txt == "(" then
      repeat 
        local arg = self:analyzeExp(  )
        table.insert( formatArgList, arg )
        nextToken = self:getToken(  )
      until nextToken.txt ~= ","
      self:checkToken( nextToken, ")" )
      nextToken = self:getToken(  )
    end
    exp = self:createNode( nodeKindLiteralString, firstToken.pos, typeInfoString, {["token"] = token, ["argList"] = formatArgList} )
    token = nextToken
    if token.txt == "[" then
      exp = self:analyzeExpRefItem( token, exp )
    else 
      self:pushback(  )
    end
  elseif token.txt == "fn" then
    exp = self:analyzeExpSymbol( firstToken, token, "fn", token, false )
  elseif token.kind == Parser.kind.Symb then
    exp = self:analyzeExpSymbol( firstToken, token, "symbol", token, false )
  elseif token.kind == Parser.kind.Type then
    exp = self:createNode( nodeKindExpRef, firstToken.pos, typeInfoNone, token )
  elseif token.txt == "true" or token.txt == "false" then
    exp = self:createNode( nodeKindLiteralBool, firstToken.pos, typeInfoBool, token )
  elseif token.txt == "nil" then
    exp = self:createNode( nodeKindLiteralNil, firstToken.pos, typeInfoNil, token )
  end
  if not exp then
    self:error( "illegal exp" )
  end
  if skipOp2Flag then
    return exp
  end
  return self:analyzeExpOp2( firstToken, exp )
end

function TransUnit:createAST( parser )
  self:registBuiltInScope(  )
  local rootInfo = {}
  rootInfo.childlen = {}
  local ast = self:createNode( nodeKindRoot, {["lineNo"] = 0, ["column"] = 0}, typeInfoNone, rootInfo )
  self.parser = parser
  self.moduleName2Info = {}
  self:analyzeStatement( rootInfo.childlen )
  local token = self:getTokenNoErr(  )
  if token then
    error( string.format( "unknown:%d:%d:(%s) %s", token.pos.lineNo, token.pos.column, Parser.getKindTxt( token.kind ), token.txt ) )
  end
  return ast
end

function TransUnit:analyzeStatement( stmtList, termTxt )
  while true do
    local token = self:getTokenNoErr(  )
    if not token then
      break
    end
    local statement = self:analyzeDecl( "pri", false, token, token )
    if not statement then
      if token.txt == termTxt then
        self:pushback(  )
        break
      elseif token.txt == "pub" or token.txt == "pro" or token.txt == "pri" or token.txt == "global" or token.txt == "static" then
        local accessMode = (token.txt ~= "static" ) and token.txt or "pri"
        local staticFlag = (token.txt == "static" )
        local nextToken = nil
        if token.txt ~= "static" then
          nextToken = self:getToken(  )
        else 
          nextToken = token
        end
        statement = self:analyzeDecl( accessMode, staticFlag, token, nextToken )
      elseif token.txt == "{" then
        self:pushback(  )
        statement = self:analyzeBlock( "{" )
      elseif token.txt == "if" then
        statement = self:analyzeIf( token )
      elseif token.txt == "while" then
        statement = self:analyzeWhile( token )
      elseif token.txt == "repeat" then
        statement = self:analyzeRepeat( token )
      elseif token.txt == "for" then
        statement = self:analyzeFor( token )
      elseif token.txt == "apply" then
        statement = self:analyzeApply( token )
      elseif token.txt == "foreach" then
        statement = self:analyzeForeach( token, false )
      elseif token.txt == "forsort" then
        statement = self:analyzeForeach( token, true )
      elseif token.txt == "return" then
        local expList = self:analyzeExpList(  )
        self:checkNextToken( ";" )
        statement = self:createNode( nodeKindReturn, token.pos, typeInfoNone, expList )
      elseif token.txt == "break" then
        self:checkNextToken( ";" )
        statement = self:createNode( nodeKindBreak, token.pos, typeInfoNone, nil )
      elseif token.txt == "import" then
        statement = self:analyzeImport( token )
      else 
        self:pushback(  )
        local exp = self:analyzeExp(  )
        self:checkNextToken( ";" )
        statement = self:createNode( nodeKindStmtExp, token.pos, typeInfoNone, exp )
      end
    end
    if not statement then
      break
    end
    table.insert( stmtList, statement )
  end
end

----- meta -----
moduleObj._typeInfoList = {
{ itemTypeId = { }, typeId = 2, txt = "", kind = 1 },
  { itemTypeId = { }, typeId = 3, txt = "stem", kind = 1 },
  { itemTypeId = { }, typeId = 9, txt = "str", kind = 1 },
  { itemTypeId = { 2, 2}, typeId = 97, txt = "Map", kind = 4 },
  { itemTypeId = { 2, 2}, typeId = 147, txt = "Map", kind = 4 },
  { itemTypeId = { }, typeId = 150, txt = "nodeFilter", kind = 6 },
  { itemTypeId = { }, typeId = 151, txt = "getNodeKindName", kind = 6 },
  }
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfoNode = {}
_className2InfoMap.Node = _classInfoNode
local _classInfoScope = {}
_className2InfoMap.Scope = _classInfoScope
_classInfoScope.add = {
  name='add', staticFlag = false, accessMode = 'pub' }
_classInfoScope.addClass = {
  name='addClass', staticFlag = false, accessMode = 'pub' }
_classInfoScope.getClassScope = {
  name='getClassScope', staticFlag = false, accessMode = 'pub' }
_classInfoScope.getParent = {
  name='getParent', staticFlag = false, accessMode = 'pub' }
_classInfoScope.getTypeInfo = {
  name='getTypeInfo', staticFlag = false, accessMode = 'pub' }
_classInfoScope.getTypeInfoChild = {
  name='getTypeInfoChild', staticFlag = false, accessMode = 'pub' }
local _classInfoTransUnit = {}
_className2InfoMap.TransUnit = _classInfoTransUnit
_classInfoTransUnit.analyzeApply = {
  name='analyzeApply', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeBlock = {
  name='analyzeBlock', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDecl = {
  name='analyzeDecl', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDeclClass = {
  name='analyzeDeclClass', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDeclFunc = {
  name='analyzeDeclFunc', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDeclMember = {
  name='analyzeDeclMember', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDeclMethod = {
  name='analyzeDeclMethod', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeDeclVar = {
  name='analyzeDeclVar', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExp = {
  name='analyzeExp', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExpCont = {
  name='analyzeExpCont', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExpList = {
  name='analyzeExpList', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExpOp2 = {
  name='analyzeExpOp2', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExpRefItem = {
  name='analyzeExpRefItem', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeExpSymbol = {
  name='analyzeExpSymbol', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeFor = {
  name='analyzeFor', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeForeach = {
  name='analyzeForeach', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeIf = {
  name='analyzeIf', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeImport = {
  name='analyzeImport', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeListConst = {
  name='analyzeListConst', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeMapConst = {
  name='analyzeMapConst', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeRefType = {
  name='analyzeRefType', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeRepeat = {
  name='analyzeRepeat', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeStatement = {
  name='analyzeStatement', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.analyzeWhile = {
  name='analyzeWhile', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.checkNextToken = {
  name='checkNextToken', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.checkSymbol = {
  name='checkSymbol', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.checkToken = {
  name='checkToken', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.createAST = {
  name='createAST', staticFlag = false, accessMode = 'pub' }
_classInfoTransUnit.createNode = {
  name='createNode', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.createNoneNode = {
  name='createNoneNode', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.error = {
  name='error', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.getSymbolToken = {
  name='getSymbolToken', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.getToken = {
  name='getToken', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.getTokenNoErr = {
  name='getTokenNoErr', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.popClass = {
  name='popClass', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.popScope = {
  name='popScope', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.pushClass = {
  name='pushClass', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.pushScope = {
  name='pushScope', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.pushback = {
  name='pushback', staticFlag = false, accessMode = 'pri' }
_classInfoTransUnit.registBuiltInScope = {
  name='registBuiltInScope', staticFlag = false, accessMode = 'pri' }
local _classInfoTypeInfo = {}
_className2InfoMap.TypeInfo = _classInfoTypeInfo
_classInfoTypeInfo.create = {
  name='create', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createArray = {
  name='createArray', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createBuiltin = {
  name='createBuiltin', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createClass = {
  name='createClass', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createFunc = {
  name='createFunc', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createList = {
  name='createList', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.createMap = {
  name='createMap', staticFlag = true, accessMode = 'pub' }
_classInfoTypeInfo.getItemTypeInfoList = {
  name='getItemTypeInfoList', staticFlag = false, accessMode = 'pub' }
_classInfoTypeInfo.getKind = {
  name='getKind', staticFlag = false, accessMode = 'pub' }
_classInfoTypeInfo.getRetTypeInfoList = {
  name='getRetTypeInfoList', staticFlag = false, accessMode = 'pub' }
_classInfoTypeInfo.getTxt = {
  name='getTxt', staticFlag = false, accessMode = 'pub' }
_classInfoTypeInfo.getTypeId = {
  name='getTypeId', staticFlag = false, accessMode = 'pub' }
_classInfoTypeInfo.serialize = {
  name='serialize', staticFlag = false, accessMode = 'pub' }
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
_varName2InfoMap.nodeKind = {
  name='nodeKind', accessMode = 'pub', typeId = 147 }
_varName2InfoMap.typeInfo = {
  name='typeInfo', accessMode = 'pub', typeId = 97 }
local _funcName2InfoMap = {}
moduleObj._funcName2InfoMap = _funcName2InfoMap
_funcName2InfoMap.getNodeKindName = {
  accessMode = 'pub', typeId = 151 }
_funcName2InfoMap.nodeFilter = {
  accessMode = 'pub', typeId = 150 }
----- meta -----
return moduleObj
