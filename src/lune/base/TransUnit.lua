--lune/base/TransUnit.lns
local moduleObj = {}
local Parser = require( 'lune.base.Parser' )

local function debugLog(  )
  local debugInfo1 = debug.getinfo( 2 )
  local debugInfo2 = debug.getinfo( 3 )
  local debugInfo3 = debug.getinfo( 4 )
  print( "--", debugInfo1["short_src"], debugInfo1['currentline'] )
  print( "--", debugInfo2["short_src"], debugInfo2['currentline'] )
  print( "--", debugInfo3["short_src"], debugInfo3['currentline'] )
end

local typeIdSeed = 0
local typeInfo = {}
moduleObj.typeInfo = typeInfo

local builtInTypeMap = {}
local TypeInfoKindPrim = 1
moduleObj.TypeInfoKindPrim = TypeInfoKindPrim

local TypeInfoKindList = 2
moduleObj.TypeInfoKindList = TypeInfoKindList

local TypeInfoKindArray = 3
moduleObj.TypeInfoKindArray = TypeInfoKindArray

local TypeInfoKindMap = 4
moduleObj.TypeInfoKindMap = TypeInfoKindMap

local TypeInfoKindClass = 5
moduleObj.TypeInfoKindClass = TypeInfoKindClass

local TypeInfoKindFunc = 6
moduleObj.TypeInfoKindFunc = TypeInfoKindFunc

local TypeInfo = {}
function TypeInfo.new( autoFlag, externalFlag, staticFlag, accessMode, txt, typeId, kind, itemTypeInfoList, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = TypeInfo } )
  return obj.__init and obj:__init( autoFlag, externalFlag, staticFlag, accessMode, txt, typeId, kind, itemTypeInfoList, retTypeInfoList ) or nil;
end
function TypeInfo:__init(autoFlag, externalFlag, staticFlag, accessMode, txt, typeId, kind, itemTypeInfoList, retTypeInfoList) 
  self.autoFlag = autoFlag
  self.externalFlag = externalFlag
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.txt = txt
  self.typeId = typeId
  self.kind = kind
  self.itemTypeInfoList = itemTypeInfoList or {}
  self.retTypeInfoList = retTypeInfoList or {}
  return self
end
function TypeInfo:serialize(  )
  local function serializeTypeInfoList( name, list )
    local work = name
    for index, typeInfo in pairs( list ) do
      if index ~= 1 then
        work = work .. ", "
      end
      work = string.format( "%s%d", work, typeInfo.typeId)
    end
    return work .. "}, "
  end
  
  local txt = string.format( [==[{ typeId = %d, txt = '%s', staticFlag = %s, accessMode = '%s',
kind = %d, ]==], self.typeId, self.txt, self.staticFlag, self.accessMode, self.kind)
  return txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. "}"
end
function TypeInfo:getTxt(  )
  if self.kind == TypeInfoKindArray then
    return self.itemTypeInfoList[1]:getTxt(  ) .. "[@]"
  end
  if self.kind == TypeInfoKindList then
    return self.itemTypeInfoList[1]:getTxt(  ) .. "[]"
  end
  if self.itemTypeInfoList and #self.itemTypeInfoList > 0 then
    local txt = self.txt .. "<"
    for index, typeInfo in pairs( self.itemTypeInfoList ) do
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
function TypeInfo.create( staticFlag, kind, txt, itemTypeInfo, retTypeInfoList )
  if kind == TypeInfoKindPrim then
    return builtInTypeMap[txt]
  end
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(false, true, staticFlag, "pub", txt, typeIdSeed, kind, itemTypeInfo, retTypeInfoList)
  return info
end
function TypeInfo.createBuiltin( idName, typeTxt, kind )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(false, false, false, "pub", typeTxt, typeIdSeed, kind)
  typeInfo[idName] = info
  builtInTypeMap[typeTxt] = info
  return info
end
function TypeInfo.createList( itemTypeInfo )
  if not itemTypeInfo or #itemTypeInfo == 0 then
    error( string.format( "illegal list type: %s", itemTypeInfo) )
  end
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(false, false, false, "pub", nil, typeIdSeed, TypeInfoKindList, itemTypeInfo)
end
function TypeInfo.createArray( itemTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(false, false, false, "pub", nil, typeIdSeed, TypeInfoKindArray, itemTypeInfo)
end
function TypeInfo.createMap( keyTypeInfo, valTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(false, false, false, "pub", "Map", typeIdSeed, TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end
function TypeInfo.createClass( externalFlag, accessMode, className )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(false, externalFlag, false, accessMode, className, typeIdSeed, TypeInfoKindClass)
  return info
end
function TypeInfo.createFunc( autoFlag, externalFlag, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(autoFlag, externalFlag, staticFlag, accessMode, funcName, typeIdSeed, TypeInfoKindFunc, argTypeList or {}, retTypeInfoList or {})
  return info
end
function TypeInfo:get_itemTypeInfoList()
   return self.itemTypeInfoList
end
function TypeInfo:get_retTypeInfoList()
   return self.retTypeInfoList
end
function TypeInfo:get_typeId()
   return self.typeId
end
function TypeInfo:get_txt()
   return self.txt
end
function TypeInfo:get_kind()
   return self.kind
end
function TypeInfo:get_staticFlag()
   return self.staticFlag
end
function TypeInfo:get_accessMode()
   return self.accessMode
end
function TypeInfo:get_autoFlag()
   return self.autoFlag
end

local typeInfoNone = TypeInfo.createBuiltin( "None", "", TypeInfoKindPrim )
local typeInfoStem = TypeInfo.createBuiltin( "Stem", "stem", TypeInfoKindPrim )
local typeInfoNil = TypeInfo.createBuiltin( "Nil", "nil", TypeInfoKindPrim )
local typeInfoBool = TypeInfo.createBuiltin( "Bool", "bool", TypeInfoKindPrim )
local typeInfoInt = TypeInfo.createBuiltin( "Int", "int", TypeInfoKindPrim )
local typeInfoReal = TypeInfo.createBuiltin( "Real", "real", TypeInfoKindPrim )
local typeInfoChar = TypeInfo.createBuiltin( "char", "char", TypeInfoKindPrim )
local typeInfoString = TypeInfo.createBuiltin( "String", "str", TypeInfoKindPrim )
local typeInfoMap = TypeInfo.createBuiltin( "Map", "Map", TypeInfoKindMap )
local typeInfoList = TypeInfo.createBuiltin( "List", "List", TypeInfoKindList )
local typeInfoArray = TypeInfo.createBuiltin( "Array", "Array", TypeInfoKindArray )
local typeInfoForm = TypeInfo.createBuiltin( "Form", "form", TypeInfoKindFunc )
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
function Scope:get_parent()
   return self.parent
end
function Scope:get_symbol2TypeInfoMap()
   return self.symbol2TypeInfoMap
end
function Scope:get_className2ScopeMap()
   return self.className2ScopeMap
end

local NodePos = {}
moduleObj.NodePos = NodePos
function NodePos.new( lineNo, column )
  local obj = {}
  setmetatable( obj, { __index = NodePos } )
  return obj.__init and obj:__init( lineNo, column ) or nil;
end
function NodePos:__init( lineNo, column )
            
self.lineNo = lineNo
  self.column = column
    return self
end
            

local Node = {}
moduleObj.Node = Node
function Node.new( kind, pos, expType, expTypeList, info, filter )
  local obj = {}
  setmetatable( obj, { __index = Node } )
  return obj.__init and obj:__init( kind, pos, expType, expTypeList, info, filter ) or nil;
end
function Node:__init( kind, pos, expType, expTypeList, info, filter )
            
self.kind = kind
  self.pos = pos
  self.expType = expType
  self.expTypeList = expTypeList
  self.info = info
  self.filter = filter
    return self
end
            
function Node:get_kind()
   return self.kind
end
function Node:get_expType()
   return self.expType
end
function Node:get_info()
   return self.info
end

local ClassInfo = {}
function ClassInfo.new( name, scope, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = ClassInfo } )
  return obj.__init and obj:__init( name, scope, typeInfo ) or nil;
end
function ClassInfo:__init( name, scope, typeInfo )
            
self.name = name
  self.scope = scope
  self.typeInfo = typeInfo
    return self
end
            

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
  self.typeInfo2ClassNode = {}
  return self
end
function TransUnit:pushScope(  )
  self.scope = Scope.new(self.scope)
  return self.scope
end
function TransUnit:popScope(  )
  self.scope = self.scope:get_parent(  )
end
function TransUnit:pushClass( externalFlag, name )
  local typeInfo = self.scope:getTypeInfoChild( name )
  if not typeInfo then
    typeInfo = TypeInfo.createClass( externalFlag, "pub", name )
    local scope = self:pushScope(  )
    scope:get_parent(  ):addClass( name, typeInfo, scope )
  else 
    self.scope = self.scope:getClassScope( name )
  end
  table.insert( self.classList, ClassInfo.new(name, self.scope, typeInfo) )
  self.typeInfo2Scope[typeInfo] = self.scope
  return typeInfo
end
function TransUnit:popClass(  )
  self:popScope(  )
  table.remove( self.classList )
end
function TransUnit:addMethod( className, methodNode )
  local classTypeInfo = self.scope:getTypeInfo( className )
  local classNodeInfo = self.typeInfo2ClassNode[classTypeInfo].info
  classNodeInfo.outerMethodSet[methodNode.info.name.txt] = true
  table.insert( classNodeInfo.fieldList, methodNode )
end
-- none
-- none
-- none
-- none
-- none
-- none
-- none
-- none

local opLevelBase = 0
local op2levelMap = {}
local op1levelMap = {}
local function regOpLevel( opnum, opList )
  opLevelBase = opLevelBase + 1
  if opnum == 1 then
    for __index, op in pairs( opList ) do
      op1levelMap[op] = opLevelBase
    end
  else 
    for __index, op in pairs( opList ) do
      op2levelMap[op] = opLevelBase
    end
  end
end

regOpLevel( 2, {"="} )
regOpLevel( 2, {"or"} )
regOpLevel( 2, {"and"} )
regOpLevel( 2, {"<", ">", "<=", ">=", "~=", "=="} )
regOpLevel( 2, {"|"} )
regOpLevel( 2, {"~"} )
regOpLevel( 2, {"&"} )
regOpLevel( 2, {"<<", ">>"} )
regOpLevel( 2, {".."} )
regOpLevel( 2, {"+", "-"} )
regOpLevel( 2, {"*", "/", "//", "%"} )
regOpLevel( 1, {"not", "#", "-", "~"} )
regOpLevel( 1, {"^"} )
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
local nodeKindSwitch = regKind( 'Switch' )
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
local function getNodeKindName( kind )
  return nodeKind2NameMap[kind]
end
moduleObj.getNodeKindName = getNodeKindName
local function nodeFilter( node, filter, ... )
  if not filter[node.kind] then
    error( string.format( "none filter -- %s", getNodeKindName( node.kind ) ) )
  end
  return filter[node.kind]( filter, node, ... )
end
moduleObj.nodeFilter = nodeFilter
function TransUnit:registBuiltInScope(  )
  local builtInInfo = {[""] = {["error"] = {["ret"] = {}}, ["print"] = {["ret"] = {}}, ["tonumber"] = {["ret"] = {"int"}}}, ["io"] = {["open"] = {["ret"] = {"stem"}}}, ["string"] = {["find"] = {["ret"] = {"int", "int"}}, ["byte"] = {["ret"] = {"int"}}, ["format"] = {["ret"] = {"str"}}, ["rep"] = {["ret"] = {"str"}}, ["gmatch"] = {["ret"] = {"stem"}}, ["gsub"] = {["ret"] = {"str"}}}, ["str"] = {["find"] = {["methodFlag"] = {}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["ret"] = {"stem"}}, ["gsub"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["ret"] = {"str"}}}, ["table"] = {["insert"] = {["ret"] = {""}}, ["remove"] = {["ret"] = {""}}}, ["debug"] = {["getinfo"] = {["ret"] = {"stem"}}}, ["_luneScript"] = {["loadModule"] = {["ret"] = {"stem"}}}}
  do
    local __sorted = {}
    local __map = builtInInfo
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, name in ipairs( __sorted ) do
      name2FuncInfo = __map[ name ]
      do
        if name ~= "" then
          self:pushClass( true, name )
        end
        do
          local __sorted = {}
          local __map = name2FuncInfo
          for __key in pairs( __map ) do
            table.insert( __sorted, __key )
          end
          table.sort( __sorted )
          for __index, funcName in ipairs( __sorted ) do
            info = __map[ funcName ]
            do
              local retTypeList = {}
              for __index, retType in pairs( info["ret"] ) do
                table.insert( retTypeList, builtInTypeMap[retType] )
              end
              self.scope:add( funcName, TypeInfo.createFunc( false, true, not info["methodFlag"], "pub", funcName, nil, retTypeList ) )
            end
          end
        end
        
        if name ~= "" then
          self:popClass(  )
        end
      end
    end
  end
  
end

function TransUnit:createNode( kind, pos, expTypeList, info )
  if not getNodeKindName( kind ) then
    error( string.format( "%d:%d: not found nodeKind", pos.lineNo, pos.column ) )
  end
  return Node.new(kind, pos, expTypeList[1], expTypeList, info, nodeFilter)
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
  return self:createNode( nodeKindNone, {["lineNo"] = 0, ["column"] = 0}, {typeInfoNone}, {} )
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
  local node = self:createNode( nodeKindBlock, token.pos, {typeInfoNone}, {["kind"] = blockKind, ["stmtList"] = stmtList} )
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
    return self:analyzeDeclFunc( accessMode, staticFlag, nil, token, nil )
  elseif token.txt == "class" then
    return self:analyzeDeclClass( accessMode, token )
  end
  return nil
end

local _TypeInfo = {}
function _TypeInfo.new( itemTypeId, retTypeId, typeId, txt, kind, staticFlag )
  local obj = {}
  setmetatable( obj, { __index = _TypeInfo } )
  return obj.__init and obj:__init( itemTypeId, retTypeId, typeId, txt, kind, staticFlag ) or nil;
end
function _TypeInfo:__init( itemTypeId, retTypeId, typeId, txt, kind, staticFlag )
            
self.itemTypeId = itemTypeId
  self.retTypeId = retTypeId
  self.typeId = typeId
  self.txt = txt
  self.kind = kind
  self.staticFlag = staticFlag
    return self
end
            

local _ModuleInfo = {}
function _ModuleInfo.new( _className2InfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
  local obj = {}
  setmetatable( obj, { __index = _ModuleInfo } )
  return obj.__init and obj:__init( _className2InfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap ) or nil;
end
function _ModuleInfo:__init( _className2InfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
            
self._className2InfoMap = _className2InfoMap
  self._typeInfoList = _typeInfoList
  self._varName2InfoMap = _varName2InfoMap
  self._funcName2InfoMap = _funcName2InfoMap
    return self
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
  local moduleTypeInfo = self:pushClass( true, moduleToken.txt )
  local moduleInfo = _luneScript.loadModule( modulePath )
  self.moduleName2Info[modulePath] = moduleInfo
  local typeId2TypeInfo = {}
  for __index, typeInfo in pairs( moduleInfo._typeInfoList ) do
    if typeInfo.kind ~= TypeInfoKindPrim then
      local itemTypeInfo = {}
      for __index, typeId in pairs( typeInfo.itemTypeId ) do
        table.insert( itemTypeInfo, typeId2TypeInfo[typeId] )
      end
      local retTypeInfo = {}
      for __index, typeId in pairs( typeInfo.retTypeId ) do
        table.insert( retTypeInfo, typeId2TypeInfo[typeId] )
      end
      typeId2TypeInfo[typeInfo.typeId] = TypeInfo.create( typeInfo.staticFlag, typeInfo.typeId, typeInfo.txt, itemTypeInfo, retTypeInfo )
    else 
      typeId2TypeInfo[typeInfo.typeId] = builtInTypeMap[typeInfo.txt]
    end
  end
  for className, classInfo in pairs( moduleInfo._className2InfoMap ) do
    self:pushClass( true, className )
    for fieldName, fieldInfo in pairs( classInfo ) do
      local fieldTypeInfo = nil
      if fieldInfo.methodFlag then
        fieldTypeInfo = typeId2TypeInfo[fieldInfo["typeId"]]
      else 
        fieldTypeInfo = typeId2TypeInfo[fieldInfo["typeId"]]
      end
      self.scope:add( fieldName, fieldTypeInfo )
    end
    self:popClass(  )
  end
  for varName, varInfo in pairs( moduleInfo._varName2InfoMap ) do
    self.scope:add( varName, typeId2TypeInfo[varInfo["typeId"]] )
  end
  for funcName, funcInfo in pairs( moduleInfo._funcName2InfoMap ) do
    self.scope:add( funcName, typeId2TypeInfo[funcInfo["typeId"]] )
  end
  self:popClass(  )
  self:checkToken( nextToken, ";" )
  return self:createNode( nodeKindImport, token.pos, {typeInfoNone}, modulePath )
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
  return self:createNode( nodeKindIf, token.pos, {typeInfoNone}, list )
end

function TransUnit:analyzeSwitch( firstToken )
  local exp = self:analyzeExp(  )
  self:checkNextToken( "{" )
  local caseList = {}
  local nextToken = self:getToken(  )
  while (nextToken.txt == "case" ) do
    self:checkToken( nextToken, "case" )
    local condexpList = self:analyzeExpList(  )
    local condBock = self:analyzeBlock( "switch" )
    table.insert( caseList, {["expList"] = condexpList, ["block"] = condBock} )
    nextToken = self:getToken(  )
  end
  local defaultBlock = nil
  if nextToken.txt == "default" then
    defaultBlock = self:analyzeBlock( "default" )
  end
  self:checkNextToken( "}" )
  local info = {["exp"] = exp, ["caseList"] = caseList, ["default"] = defaultBlock}
  return self:createNode( nodeKindSwitch, firstToken.pos, {typeInfoNone}, info )
end

function TransUnit:analyzeWhile( token )
  local info = {["exp"] = self:analyzeExp(  ), ["block"] = self:analyzeBlock( "while" )}
  return self:createNode( nodeKindWhile, token.pos, {typeInfoNone}, info )
end

function TransUnit:analyzeRepeat( token )
  local scope = self:pushScope(  )
  local info = {["block"] = self:analyzeBlock( "repeat", scope ), ["exp"] = self:analyzeExp(  )}
  self:popScope(  )
  local node = self:createNode( nodeKindRepeat, token.pos, {typeInfoNone}, info )
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
  self.scope:add( val.txt, exp1.expType )
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
  local node = self:createNode( nodeKindFor, token.pos, {typeInfoNone}, info )
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
  return self:createNode( nodeKindApply, token.pos, {typeInfoNone}, info )
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
    self:error( "unknown type of exp" )
  else 
    local itemTypeInfoList = exp.expType:get_itemTypeInfoList(  )
    if exp.expType:get_kind(  ) == TypeInfoKindMap then
      self.scope:add( valSymbol.txt, itemTypeInfoList[2] )
      if keySymbol then
        self.scope:add( keySymbol.txt, itemTypeInfoList[1] )
      end
    elseif exp.expType:get_kind(  ) == TypeInfoKindList or exp.expType:get_kind(  ) == TypeInfoKindArray then
      self.scope:add( valSymbol.txt, itemTypeInfoList[1] )
      if keySymbol then
        self.scope:add( keySymbol.txt, typeInfoInt )
      else 
        self.scope:add( "__index", typeInfoInt )
      end
    else 
      self:error( string.format( "unknown kind type of exp for foreach-- %d:%d", exp.pos.lineNo, exp.pos.column) )
    end
  end
  local block = self:analyzeBlock( "foreach", scope )
  self:popScope(  )
  local info = {["val"] = valSymbol, ["key"] = keySymbol, ["exp"] = exp, ["block"] = block, ["sort"] = sortFlag}
  return self:createNode( sortFlag and nodeKindForsort or nodeKindForeach, token.pos, {typeInfoNone}, info )
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
  local name = nil
  local typeInfo = typeInfoStem
  if self:checkSymbol( token ) then
    name = self:analyzeExpSymbol( firstToken, token, "symbol", token, true )
    typeInfo = name.expType
  else 
    self:pushback(  )
  end
  local arrayMode = "no"
  token = self:getToken(  )
  if token.txt == '[' or token.txt == '[@' then
    if token.txt == '[' then
      arrayMode = "list"
      typeInfo = TypeInfo.createList( {typeInfo} )
    else 
      arrayMode = "array"
      typeInfo = TypeInfo.createArray( {typeInfo} )
    end
    token = self:getToken(  )
    if token.txt ~= ']' then
      self:pushback(  )
      self:checkNextToken( ']' )
    end
  elseif token.txt == "<" then
    local genericList = {}
    local nextToken = nil
    repeat 
      local typeExp = self:analyzeRefType(  )
      table.insert( genericList, typeExp.expType )
      nextToken = self:getToken(  )
    until nextToken.txt ~= ","
    self:checkToken( nextToken, '>' )
    if typeInfo.kind == TypeInfoKindMap then
      typeInfo = TypeInfo.createMap( genericList[1], genericList[2] )
    else 
      self:error( string.format( "not support generic: %s", typeInfo:getTxt(  ) ) )
    end
  else 
    self:pushback(  )
  end
  return self:createNode( nodeKindRefType, firstToken.pos, {typeInfo}, {["name"] = name, ["refFlag"] = refFlag, ["mutFlag"] = mutFlag, ["array"] = arrayMode} )
end

function TransUnit:analyzeDeclMember( accessMode, staticFlag, firstToken )
  local varName = self:getSymbolToken(  )
  local token = self:getToken(  )
  local refType = self:analyzeRefType(  )
  token = self:getToken(  )
  local getterMode = "none"
  local setterMode = "none"
  if token.txt == "{" then
    local nextToken = self:getToken(  )
    if nextToken.txt == "pub" or nextToken.txt == "pri" then
      getterMode = nextToken.txt
      nextToken = self:getToken(  )
      if nextToken.txt == "," then
        nextToken = self:getToken(  )
        if nextToken.txt == "pub" or nextToken.txt == "pri" then
          setterMode = nextToken.txt
          nextToken = self:getToken(  )
        end
      end
    end
    self:checkToken( nextToken, "}" )
    token = self:getToken(  )
  end
  self:checkToken( token, ";" )
  self.scope:add( varName.txt, refType.expType )
  local info = {["name"] = varName, ["refType"] = refType.expType, ["staticFlag"] = staticFlag, ["accessMode"] = accessMode, ["getterMode"] = getterMode, ["setterMode"] = setterMode}
  return self:createNode( nodeKindDeclMember, firstToken.pos, refType.expTypeList, info )
end

function TransUnit:analyzeDeclMethod( accessMode, staticFlag, className, firstToken, name )
  local node = self:analyzeDeclFunc( accessMode, staticFlag, className, name, name )
  return node
end

function TransUnit:analyzeDeclClass( classAccessMode, classToken )
  local name = self:getToken(  )
  self:checkNextToken( "{" )
  local classTypeInfo = self:pushClass( false, name.txt )
  local fieldList = {}
  local memberList = {}
  local methodName2Node = {}
  local node = self:createNode( nodeKindDeclClass, classToken.pos, {classTypeInfo}, {["accessMode"] = classAccessMode, ["name"] = name, ["fieldList"] = fieldList, ["memberList"] = memberList, ["scope"] = self.scope, ["outerMethodSet"] = {}} )
  self.typeInfo2ClassNode[classTypeInfo] = node
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
      local memberNode = self:analyzeDeclMember( accessMode, staticFlag, token )
      table.insert( fieldList, memberNode )
      table.insert( memberList, memberNode )
    else 
      local methodNode = self:analyzeDeclMethod( accessMode, staticFlag, name, token, token )
      table.insert( fieldList, methodNode )
    end
  end
  for __index, memberNode in pairs( memberList ) do
    local memberName = memberNode.info.name
    local getterName = "get_" .. memberName.txt
    local accessMode = memberNode.info.getterMode
    if accessMode ~= "none" and not self.scope:getTypeInfo( getterName ) then
      self.scope:add( getterName, TypeInfo.createFunc( true, false, false, accessMode, getterName, nil, {memberNode.expType} ) )
    end
    local setterName = "set_" .. memberName.txt
    local accessMode = memberNode.info.setterMode
    if memberNode.info.setterMode ~= "none" and not self.scope:getTypeInfo( setterName ) then
      self.scope:add( setterName, TypeInfo.createFunc( true, false, false, accessMode, setterName, nil, {memberNode.expType} ) )
    end
  end
  self:popClass(  )
  return node
end

function TransUnit:analyzeDeclFunc( accessMode, staticFlag, classNameToken, firstToken, name )
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
  local needPopFlag = false
  if token.txt == "." then
    needPopFlag = true
    classNameToken = name
    self:pushClass( false, name.txt )
    name = self:getSymbolToken(  )
    token = self:getToken(  )
  end
  self:checkToken( "(" )
  local kind = nodeKindDeclConstr
  if classNameToken then
    if name.txt == "__init" then
      kind = nodeKindDeclConstr
    else 
      kind = nodeKindDeclMethod
    end
  else 
    kind = nodeKindDeclFunc
    if not staticFlag then
      staticFlag = true
    end
  end
  local scope = self:pushScope(  )
  repeat 
    local argName = self:getToken(  )
    if argName.txt == ")" then
      token = argName
      break
    elseif argName.txt == "..." then
      table.insert( argList, self:createNode( nodeKindDeclArgDDD, argName.pos, {typeInfoNone}, argName ) )
    else 
      self:checkSymbol( argName )
      self:checkNextToken( ":" )
      local refType = self:analyzeRefType(  )
      local arg = self:createNode( nodeKindDeclArg, argName.pos, refType.expTypeList, {["name"] = argName, ["argType"] = refType} )
      self.scope:add( argName.txt, refType.expType )
      table.insert( argList, arg )
    end
    token = self:getToken(  )
  until token.txt ~= ","
  self:checkToken( token, ")" )
  token = self:getToken(  )
  local retTypeList = {}
  local retTypeInfoList = {}
  if token.txt == ":" then
    repeat 
      local refType = self:analyzeRefType(  )
      table.insert( retTypeList, refType )
      table.insert( retTypeInfoList, refType.expType )
      token = self:getToken(  )
    until token.txt ~= ","
  end
  local typeInfo = TypeInfo.createFunc( false, false, staticFlag, accessMode, name and name.txt or "", nil, retTypeInfoList )
  if name then
    scope:get_parent(  ):add( name.txt, typeInfo )
  end
  local node = nil
  local info = nil
  if token.txt == ";" then
    node = self:createNoneNode(  )
  else 
    self:pushback(  )
    local body = self:analyzeBlock( "func", scope )
    info = {["name"] = name, ["argList"] = argList, ["staticFlag"] = staticFlag, ["body"] = body, ["accessMode"] = accessMode, ["className"] = classNameToken, ["retTypeList"] = retTypeList, ["retTypeInfoList"] = retTypeInfoList}
    node = self:createNode( kind, firstToken.pos, {typeInfo}, info )
  end
  self:popScope(  )
  if needPopFlag then
    self:addMethod( classNameToken.txt, node )
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
    local typeInfo = typeInfoNone
    local refType = nil
    if token.txt == ":" then
      refType = self:analyzeRefType(  )
      typeInfo = refType.expType
      token = self:getToken(  )
    end
    table.insert( varList, {["name"] = varName, ["refType"] = refType} )
    table.insert( typeInfoList, typeInfo )
  until token.txt ~= ","
  local expList = nil
  if token.txt == "=" then
    expList = self:analyzeExpList(  )
    if not expList then
      self:error( "expList is nil" )
    end
  end
  if expList then
    local nodeList = expList.info
    for index, exp in pairs( nodeList ) do
      if not typeInfoList[index] or typeInfoList[index] == typeInfoNone then
        typeInfoList[index] = exp["expType"]
      end
    end
  end
  self:checkNextToken( ";" )
  local declVarInfo = {["accessMode"] = accessMode, ["varList"] = varList, ["expList"] = expList, ["typeInfoList"] = typeInfoList}
  local node = self:createNode( nodeKindDeclVar, firstToken.pos, {typeInfoNone}, declVarInfo )
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
  return self:createNode( nodeKindExpList, firstExp.pos, {typeInfoNone}, expList )
end

function TransUnit:analyzeListConst( token )
  local nextToken = self:getToken(  )
  local expList = nil
  local itemTypeInfo = typeInfoNone
  if nextToken.txt ~= "]" then
    self:pushback(  )
    expList = self:analyzeExpList(  )
    self:checkNextToken( "]" )
    local nodeList = expList.info
    for __index, exp in pairs( nodeList ) do
      if itemTypeInfo == typeInfoNone then
        itemTypeInfo = exp["expType"]
      elseif itemTypeInfo ~= exp["expType"] then
        itemTypeInfo = typeInfoStem
      end
    end
  end
  local kind = nodeKindLiteralArray
  if token.txt == '[' then
    kind = nodeKindLiteralList
  end
  return self:createNode( kind, token.pos, {TypeInfo.createList( {itemTypeInfo} )}, expList )
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
  return self:createNode( nodeKindLiteralMap, token.pos, {typeInfo}, {["map"] = map, ["pairList"] = pairList} )
end

function TransUnit:analyzeExpRefItem( token, exp )
  local indexExp = self:analyzeExp(  )
  self:checkNextToken( "]" )
  local info = {["val"] = exp, ["index"] = indexExp}
  local typeInfo = typeInfoStem
  if exp.expType then
    if exp.expType.kind == TypeInfoKindMap then
      typeInfo = exp.expType:get_itemTypeInfoList(  )[2]
    elseif exp.expType.kind == TypeInfoKindArray or exp.expType.kind == TypeInfoKindArray then
      typeInfo = exp.expType:get_itemTypeInfoList(  )[1]
    end
  end
  return self:createNode( nodeKindExpRefItem, token.pos, {typeInfo}, info )
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
        exp = self:createNode( nodeKindExpCall, firstToken.pos, exp.expType:get_retTypeInfoList(  ), info )
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
    if prefixExp.expType:get_kind(  ) == TypeInfoKindClass then
      local classScope = self.typeInfo2Scope[prefixExp.expType]
      local className = prefixExp.expType:get_txt(  )
      if not classScope then
        self:error( "not found field: " .. className )
      end
      typeInfo = classScope:getTypeInfo( token.txt )
      if not typeInfo then
        print( "hoge", classScope.symbol2TypeInfoMap )
        for __index, name in pairs( classScope.symbol2TypeInfoMap ) do
          print( "hoge", name )
        end
        self:error( string.format( "not found field typeInfo: %s.%s %s", className, token.txt, classScope ) )
      end
    end
    exp = self:createNode( nodeKindRefField, firstToken.pos, {typeInfo}, info )
  elseif mode == "symbol" then
    local typeInfo = self.scope:getTypeInfo( token.txt )
    if not typeInfo and token.txt == "self" then
      local classInfo = self.classList[#self.classList]
      typeInfo = classInfo.typeInfo
    end
    if not typeInfo then
      self:error( "not found type -- " .. token.txt )
    end
    exp = self:createNode( nodeKindExpRef, firstToken.pos, {typeInfo}, token )
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( "pri", false, nil, token, nil )
  else 
    self:error( "illegal mode", mode )
  end
  return self:analyzeExpCont( firstToken, exp, skipFlag )
end

function TransUnit:analyzeExpOp2( firstToken, exp, prevOpLevel )
  while true do
    local nextToken = self:getToken(  )
    local opLevel = prevOpLevel
    local opTxt = nextToken.txt
    if opTxt == "@" then
      local castType = self:analyzeRefType(  )
      exp = self:createNode( nodeKindExpCast, firstToken.pos, castType.expTypeList, exp )
    elseif nextToken.kind == Parser.kind.Ope then
      if Parser.isOp2( opTxt ) then
        opLevel = op2levelMap[opTxt]
        if not opLevel then
          error( string.format( "unknown op -- %s %s", opTxt, prevOpLevel ) )
        end
        if prevOpLevel and opLevel <= prevOpLevel then
          self:pushback(  )
          return exp
        end
        local exp2 = self:analyzeExp( false, opLevel )
        local info = {["op"] = nextToken, ["exp1"] = exp, ["exp2"] = exp2}
        exp = self:createNode( nodeKindExpOp2, firstToken.pos, {typeInfoNone}, info )
      else 
        self:error( "illegal op" )
      end
    else 
      self:pushback(  )
      return exp
    end
  end
  return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeExp( skipOp2Flag, prevOpLevel )
  local firstToken = self:getToken(  )
  local token = firstToken
  local exp = nil
  if token.kind == Parser.kind.Dlmt then
    if token.txt == "..." then
      return self:createNode( nodeKindExpDDD, firstToken.pos, {typeInfoNone}, token )
    end
    if token.txt == '[' or token.txt == '[@' then
      return self:analyzeListConst( token )
    end
    if token.txt == '{' then
      return self:analyzeMapConst( token )
    end
    if token.txt == "(" then
      exp = self:analyzeExp(  )
      self:checkNextToken( ")" )
      exp = self:createNode( nodeKindExpParen, firstToken.pos, exp.expTypeList, exp )
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
    exp = self:createNode( nodeKindExpNew, firstToken.pos, exp.expTypeList, {["symbol"] = exp, ["argList"] = argList} )
    exp = self:analyzeExpCont( firstToken, exp, false )
  end
  if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
    exp = self:analyzeExp( true, op1levelMap[token.txt] )
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
    exp = self:createNode( nodeKindExpOp1, firstToken.pos, {typeInfo}, {["op"] = token, ["exp"] = exp} )
    return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
  end
  if token.kind == Parser.kind.Int then
    exp = self:createNode( nodeKindLiteralInt, firstToken.pos, {typeInfoInt}, {["token"] = token, ["num"] = tonumber( token.txt )} )
  elseif token.kind == Parser.kind.Real then
    exp = self:createNode( nodeKindLiteralReal, firstToken.pos, {typeInfoReal}, {["token"] = token, ["num"] = tonumber( token.txt )} )
  elseif token.kind == Parser.kind.Char then
    local num = 0
    if #(token.txt ) == 1 then
      num = token.txt:byte( 1 )
    else 
      num = quotedChar2Code[token.txt:sub( 2, 2 )]
    end
    exp = self:createNode( nodeKindLiteralChar, firstToken.pos, {typeInfoChar}, {["token"] = token, ["num"] = num} )
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
    exp = self:createNode( nodeKindLiteralString, firstToken.os, {typeInfoString}, {["token"] = token, ["argList"] = formatArgList} )
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
    exp = self:createNode( nodeKindExpRef, firstToken.pos, {typeInfoNone}, token )
  elseif token.txt == "true" or token.txt == "false" then
    exp = self:createNode( nodeKindLiteralBool, firstToken.pos, {typeInfoBool}, token )
  elseif token.txt == "nil" then
    exp = self:createNode( nodeKindLiteralNil, firstToken.pos, {typeInfoNil}, token )
  end
  if not exp then
    self:error( "illegal exp" )
  end
  if skipOp2Flag then
    return exp
  end
  return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:createAST( parser )
  self:registBuiltInScope(  )
  local rootInfo = {}
  rootInfo.childlen = {}
  local ast = self:createNode( nodeKindRoot, {["lineNo"] = 0, ["column"] = 0}, {typeInfoNone}, rootInfo )
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
      elseif token.txt == "switch" then
        statement = self:analyzeSwitch( token )
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
        statement = self:createNode( nodeKindReturn, token.pos, {typeInfoNone}, expList )
      elseif token.txt == "break" then
        self:checkNextToken( ";" )
        statement = self:createNode( nodeKindBreak, token.pos, {typeInfoNone}, nil )
      elseif token.txt == "import" then
        statement = self:analyzeImport( token )
      else 
        self:pushback(  )
        local exp = self:analyzeExp(  )
        self:checkNextToken( ";" )
        statement = self:createNode( nodeKindStmtExp, self.currentToken.pos, {typeInfoNone}, exp )
      end
    end
    if not statement then
      break
    end
    table.insert( stmtList, statement )
  end
end

----- meta -----
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
local _classInfoClassInfo = {}
_className2InfoMap.ClassInfo = _classInfoClassInfo
local _classInfoNode = {}
_className2InfoMap.Node = _classInfoNode
_classInfoNode.filter = {
  name='filter', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 12 }
_classInfoNode.get_expType = {
  name='get_expType', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 197 }
_classInfoNode.get_info = {
  name='get_info', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 198 }
_classInfoNode.get_kind = {
  name='get_kind', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 196 }
_classInfoNode.filter = {
  name='filter', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
local _classInfoNodePos = {}
_className2InfoMap.NodePos = _classInfoNodePos
local _classInfoScope = {}
_className2InfoMap.Scope = _classInfoScope
_classInfoScope.add = {
  name='add', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 185 }
_classInfoScope.addClass = {
  name='addClass', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 186 }
_classInfoScope.getClassScope = {
  name='getClassScope', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 187 }
_classInfoScope.getTypeInfo = {
  name='getTypeInfo', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 189 }
_classInfoScope.getTypeInfoChild = {
  name='getTypeInfoChild', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 188 }
_classInfoScope.get_className2ScopeMap = {
  name='get_className2ScopeMap', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 192 }
_classInfoScope.get_parent = {
  name='get_parent', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 190 }
_classInfoScope.get_symbol2TypeInfoMap = {
  name='get_symbol2TypeInfoMap', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 191 }
local _classInfoTransUnit = {}
_className2InfoMap.TransUnit = _classInfoTransUnit
_classInfoTransUnit.createAST = {
  name='createAST', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 459 }
local _classInfoTypeInfo = {}
_className2InfoMap.TypeInfo = _classInfoTypeInfo
_classInfoTypeInfo.create = {
  name='create', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 159 }
_classInfoTypeInfo.createArray = {
  name='createArray', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 162 }
_classInfoTypeInfo.createBuiltin = {
  name='createBuiltin', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 160 }
_classInfoTypeInfo.createClass = {
  name='createClass', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 165 }
_classInfoTypeInfo.createFunc = {
  name='createFunc', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 168 }
_classInfoTypeInfo.createList = {
  name='createList', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 161 }
_classInfoTypeInfo.createMap = {
  name='createMap', staticFlag = true, accessMode = 'pub', methodFlag = true, typeId = 163 }
_classInfoTypeInfo.getTxt = {
  name='getTxt', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 156 }
_classInfoTypeInfo.get_accessMode = {
  name='get_accessMode', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 177 }
_classInfoTypeInfo.get_autoFlag = {
  name='get_autoFlag', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 178 }
_classInfoTypeInfo.get_itemTypeInfoList = {
  name='get_itemTypeInfoList', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 171 }
_classInfoTypeInfo.get_kind = {
  name='get_kind', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 175 }
_classInfoTypeInfo.get_retTypeInfoList = {
  name='get_retTypeInfoList', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 172 }
_classInfoTypeInfo.get_staticFlag = {
  name='get_staticFlag', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 176 }
_classInfoTypeInfo.get_txt = {
  name='get_txt', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 174 }
_classInfoTypeInfo.get_typeId = {
  name='get_typeId', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 173 }
_classInfoTypeInfo.serialize = {
  name='serialize', staticFlag = false, accessMode = 'pub', methodFlag = true, typeId = 153 }
local _classInfo_ModuleInfo = {}
_className2InfoMap._ModuleInfo = _classInfo_ModuleInfo
_classInfo_ModuleInfo._className2InfoMap = {
  name='_className2InfoMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 339 }
_classInfo_ModuleInfo._typeInfoList = {
  name='_typeInfoList', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 340 }
_classInfo_ModuleInfo._varName2InfoMap = {
  name='_varName2InfoMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 341 }
_classInfo_ModuleInfo._funcName2InfoMap = {
  name='_funcName2InfoMap', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 342 }
local _classInfo_TypeInfo = {}
_className2InfoMap._TypeInfo = _classInfo_TypeInfo
_classInfo_TypeInfo.itemTypeId = {
  name='itemTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 335 }
_classInfo_TypeInfo.retTypeId = {
  name='retTypeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 336 }
_classInfo_TypeInfo.typeId = {
  name='typeId', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 5 }
_classInfo_TypeInfo.txt = {
  name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 22 }
_classInfo_TypeInfo.kind = {
  name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 5 }
_classInfo_TypeInfo.staticFlag = {
  name='staticFlag', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 4 }
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
_varName2InfoMap.TypeInfoKindArray = {
  name='TypeInfoKindArray', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.TypeInfoKindClass = {
  name='TypeInfoKindClass', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.TypeInfoKindFunc = {
  name='TypeInfoKindFunc', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.TypeInfoKindList = {
  name='TypeInfoKindList', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.TypeInfoKindMap = {
  name='TypeInfoKindMap', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.TypeInfoKindPrim = {
  name='TypeInfoKindPrim', accessMode = 'pub', typeId = 5 }
_varName2InfoMap.nodeKind = {
  name='nodeKind', accessMode = 'pub', typeId = 243 }
_varName2InfoMap.typeInfo = {
  name='typeInfo', accessMode = 'pub', typeId = 143 }
local _funcName2InfoMap = {}
moduleObj._funcName2InfoMap = _funcName2InfoMap
_funcName2InfoMap.getNodeKindName = {
  accessMode = 'pub', typeId = 247 }
_funcName2InfoMap.nodeFilter = {
  accessMode = 'pub', typeId = 248 }
moduleObj._typeInfoList = {
{ typeId = 1, txt = '', staticFlag = false, accessMode = 'pub',
kind = 1, itemTypeId = {}, retTypeId = {}, },
  { typeId = 2, txt = 'stem', staticFlag = false, accessMode = 'pub',
kind = 1, itemTypeId = {}, retTypeId = {}, },
  { typeId = 4, txt = 'bool', staticFlag = false, accessMode = 'pub',
kind = 1, itemTypeId = {}, retTypeId = {}, },
  { typeId = 5, txt = 'int', staticFlag = false, accessMode = 'pub',
kind = 1, itemTypeId = {}, retTypeId = {}, },
  { typeId = 9, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {}, retTypeId = {}, },
  { typeId = 12, txt = 'form', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {}, },
  { typeId = 22, txt = 'str', staticFlag = false, accessMode = 'pub',
kind = 5, itemTypeId = {}, retTypeId = {}, },
  { typeId = 143, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {1, 1}, retTypeId = {}, },
  { typeId = 145, txt = 'TypeInfo', staticFlag = false, accessMode = 'pub',
kind = 5, itemTypeId = {}, retTypeId = {}, },
  { typeId = 146, txt = 'nil', staticFlag = false, accessMode = 'pub',
kind = 3, itemTypeId = {145}, retTypeId = {}, },
  { typeId = 147, txt = 'nil', staticFlag = false, accessMode = 'pub',
kind = 3, itemTypeId = {145}, retTypeId = {}, },
  { typeId = 153, txt = 'serialize', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {}, },
  { typeId = 156, txt = 'getTxt', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {22}, },
  { typeId = 159, txt = 'create', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {}, },
  { typeId = 160, txt = 'createBuiltin', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 161, txt = 'createList', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 162, txt = 'createArray', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 163, txt = 'createMap', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 165, txt = 'createClass', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 168, txt = 'createFunc', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 171, txt = 'get_itemTypeInfoList', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {146}, },
  { typeId = 172, txt = 'get_retTypeInfoList', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {147}, },
  { typeId = 173, txt = 'get_typeId', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {5}, },
  { typeId = 174, txt = 'get_txt', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {22}, },
  { typeId = 175, txt = 'get_kind', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {5}, },
  { typeId = 176, txt = 'get_staticFlag', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {4}, },
  { typeId = 177, txt = 'get_accessMode', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {22}, },
  { typeId = 178, txt = 'get_autoFlag', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {4}, },
  { typeId = 179, txt = 'Scope', staticFlag = false, accessMode = 'pub',
kind = 5, itemTypeId = {}, retTypeId = {}, },
  { typeId = 180, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 145}, retTypeId = {}, },
  { typeId = 181, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 179}, retTypeId = {}, },
  { typeId = 185, txt = 'add', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {}, },
  { typeId = 186, txt = 'addClass', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {}, },
  { typeId = 187, txt = 'getClassScope', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {179}, },
  { typeId = 188, txt = 'getTypeInfoChild', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 189, txt = 'getTypeInfo', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 190, txt = 'get_parent', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {179}, },
  { typeId = 191, txt = 'get_symbol2TypeInfoMap', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {180}, },
  { typeId = 192, txt = 'get_className2ScopeMap', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {181}, },
  { typeId = 196, txt = 'get_kind', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {5}, },
  { typeId = 197, txt = 'get_expType', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {145}, },
  { typeId = 198, txt = 'get_info', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {2}, },
  { typeId = 243, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 5}, retTypeId = {}, },
  { typeId = 247, txt = 'getNodeKindName', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {22}, },
  { typeId = 248, txt = 'nodeFilter', staticFlag = true, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {2}, },
  { typeId = 334, txt = '_TypeInfo', staticFlag = false, accessMode = 'pub',
kind = 5, itemTypeId = {}, retTypeId = {}, },
  { typeId = 335, txt = 'nil', staticFlag = false, accessMode = 'pub',
kind = 2, itemTypeId = {5}, retTypeId = {}, },
  { typeId = 336, txt = 'nil', staticFlag = false, accessMode = 'pub',
kind = 2, itemTypeId = {5}, retTypeId = {}, },
  { typeId = 338, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 2}, retTypeId = {}, },
  { typeId = 339, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 338}, retTypeId = {}, },
  { typeId = 340, txt = 'nil', staticFlag = false, accessMode = 'pub',
kind = 2, itemTypeId = {334}, retTypeId = {}, },
  { typeId = 341, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 2}, retTypeId = {}, },
  { typeId = 342, txt = 'Map', staticFlag = false, accessMode = 'pub',
kind = 4, itemTypeId = {22, 2}, retTypeId = {}, },
  { typeId = 459, txt = 'createAST', staticFlag = false, accessMode = 'pub',
kind = 6, itemTypeId = {}, retTypeId = {9}, },
  }
----- meta -----
return moduleObj
