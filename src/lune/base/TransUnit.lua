--lune/base/TransUnit.lns
local moduleObj = {}
local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local rootTypeId = 1
moduleObj.rootTypeId = rootTypeId

local typeIdSeed = rootTypeId + 1
-- none

local rootTypeInfo = nil
local typeInfoKind = {}
moduleObj.typeInfoKind = typeInfoKind

local builtInTypeMap = {}
local builtInTypeIdSet = {}
local TypeInfoKindRoot = 0
moduleObj.TypeInfoKindRoot = TypeInfoKindRoot

local TypeInfoKindModule = 1
moduleObj.TypeInfoKindModule = TypeInfoKindModule

local TypeInfoKindPrim = 2
moduleObj.TypeInfoKindPrim = TypeInfoKindPrim

local TypeInfoKindList = 3
moduleObj.TypeInfoKindList = TypeInfoKindList

local TypeInfoKindArray = 4
moduleObj.TypeInfoKindArray = TypeInfoKindArray

local TypeInfoKindMap = 5
moduleObj.TypeInfoKindMap = TypeInfoKindMap

local TypeInfoKindClass = 6
moduleObj.TypeInfoKindClass = TypeInfoKindClass

local TypeInfoKindFunc = 7
moduleObj.TypeInfoKindFunc = TypeInfoKindFunc

local TypeInfoKindMethod = 8
moduleObj.TypeInfoKindMethod = TypeInfoKindMethod

local TypeInfoKindNilable = 9
moduleObj.TypeInfoKindNilable = TypeInfoKindNilable

local TypeInfoKindMacro = 10
moduleObj.TypeInfoKindMacro = TypeInfoKindMacro

local function isBuiltin( typeId )
  return builtInTypeIdSet[typeId]
end
moduleObj.isBuiltin = isBuiltin
local OutStream = {}
-- none
            function OutStream.new(  )
            local obj = {}
            setmetatable( obj, { __index = OutStream } )
            if obj.__init then
            obj:__init(  )
            end       
            return obj
            end         
            function OutStream:__init(  )
            
end

-- none

-- none

local TypeInfo = {}
moduleObj.TypeInfo = TypeInfo
function TypeInfo.new( baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = TypeInfo } )
  if obj.__init then obj:__init( baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, retTypeInfoList ); end
return obj
end
function TypeInfo:__init(baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, retTypeInfoList) 
  self.baseTypeInfo = baseTypeInfo
  self.autoFlag = autoFlag
  self.externalFlag = externalFlag
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.txt = txt
  self.kind = kind
  self.itemTypeInfoList = itemTypeInfoList or {}
  self.retTypeInfoList = retTypeInfoList or {}
  self.orgTypeInfo = orgTypeInfo
  self.parentInfo = parentInfo
  self.children = {}
  if rootTypeInfo and not parentInfo then
    Util.debugLog(  )
    error( "" )
  end
  if kind == TypeInfoKindRoot then
    self.typeId = typeId
    self.nilable = false
    rootTypeInfo = self
  elseif not orgTypeInfo then
    if parentInfo then
      table.insert( parentInfo.children, self )
    end
    self.typeId = typeId + 1
    self.nilable = false
    self.nilableTypeInfo = TypeInfo.new(baseTypeInfo, self, autoFlag, externalFlag, staticFlag, accessMode, "", parentInfo, typeIdSeed, TypeInfoKindNilable, itemTypeInfoList, retTypeInfoList)
    typeIdSeed = typeIdSeed + 1
  else 
    self.typeId = typeId
    self.nilable = true
    self.nilableTypeInfo = nil
  end
end
function TypeInfo:getParentId(  )
  return self.parentInfo and self.parentInfo.typeId or rootTypeId
end
function TypeInfo:get_baseId(  )
  return self.baseTypeInfo and self.baseTypeInfo.typeId or rootTypeId
end
function TypeInfo:serialize( stream )
  if self.typeId == rootTypeId then
    return nil
  end
  local parentId = self:getParentId(  )
  if self.orgTypeInfo then
    stream:write( string.format( '{ parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }', parentId, self.typeId, self.orgTypeInfo.typeId) )
    return nil
  end
  local function serializeTypeInfoList( name, list, onlyPub )
    local work = name
    for __index, typeInfo in pairs( list ) do
      if not onlyPub or typeInfo.accessMode == "pub" then
        if #work ~= #name then
          work = work .. ", "
        end
        work = string.format( "%s%d", work, typeInfo.typeId)
      end
    end
    return work .. "}, "
  end
  
  local txt = string.format( [==[{ parentId = %d, typeId = %d, baseId = %d, txt = '%s',
        staticFlag = %s, accessMode = '%s', kind = %d, ]==], parentId, self.typeId, self:get_baseId(  ), self.txt, self.staticFlag, self.accessMode, self.kind)
  stream:write( txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. serializeTypeInfoList( "children = {", self.children, true ) .. "}\n" )
end
function TypeInfo:getTxt(  )
  if self.orgTypeInfo then
    return self.orgTypeInfo:getTxt(  ) .. "!"
  end
  if self.kind == TypeInfoKindArray then
    if not self.itemTypeInfoList[1] then
      return "[@]"
    end
    return self.itemTypeInfoList[1]:getTxt(  ) .. "[@]"
  end
  if self.kind == TypeInfoKindList then
    if not self.itemTypeInfoList[1] then
      return "[]"
    end
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
function TypeInfo:equals( typeInfo, depth )
  if not typeInfo then
    return false
  end
  if not depth then
    depth = 1
  end
  if self.typeId == typeInfo.typeId then
    return true
  end
  if self.kind ~= typeInfo.kind or self.staticFlag ~= typeInfo.staticFlag or self.accessMode ~= typeInfo.accessMode or self.autoFlag ~= typeInfo.autoFlag or self.nilable ~= typeInfo.nilable then
    return false
  end
  if (not self.itemTypeInfoList and typeInfo.itemTypeInfoList or self.itemTypeInfoList and not typeInfo.itemTypeInfoList or not self.retTypeInfoList and typeInfo.retTypeInfoList or self.retTypeInfoList and not typeInfo.retTypeInfoList or not self.orgTypeInfo and typeInfo.orgTypeInfo or self.orgTypeInfo and not typeInfo.orgTypeInfo ) then
    Util.errorLog( "%s, %s", self.itemTypeInfoList, typeInfo.itemTypeInfoList )
    Util.errorLog( "%s, %s", self.retTypeInfoList, typeInfo.retTypeInfoList )
    Util.errorLog( "%s, %s", self.orgTypeInfo, typeInfo.orgTypeInfo )
    return false
  end
  if self.itemTypeInfoList then
    if #self.itemTypeInfoList ~= #typeInfo.itemTypeInfoList then
      return false
    end
    for index, item in pairs( self.itemTypeInfoList ) do
      if not item:equals( typeInfo.itemTypeInfoList[index], depth + 1 ) then
        return false
      end
    end
  end
  if self.retTypeInfoList then
    if #self.retTypeInfoList ~= #typeInfo.retTypeInfoList then
      return false
    end
    for index, item in pairs( self.retTypeInfoList ) do
      if not item:equals( typeInfo.retTypeInfoList[index], depth + 1 ) then
        return false
      end
    end
  end
  if self.orgTypeInfo and not self.orgTypeInfo:equals( typeInfo.orgTypeInfo, depth + 1 ) then
    return false
  end
  return true
end
function TypeInfo.cloneToPublic( typeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(typeInfo.baseTypeInfo, nil, typeInfo.autoFlag, typeInfo.externalFlag, typeInfo.staticFlag, "pub", typeInfo.txt, typeInfo.parentInfo, typeIdSeed, typeInfo.kind, typeInfo.itemTypeInfoList, typeInfo.retTypeInfoList)
end
function TypeInfo.create( baseInfo, parentInfo, staticFlag, kind, txt, itemTypeInfo, retTypeInfoList )
  if kind == TypeInfoKindPrim then
    return builtInTypeMap[txt]
  end
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(baseInfo, nil, false, true, staticFlag, "pub", txt, parentInfo, typeIdSeed, kind, itemTypeInfo, retTypeInfoList)
  return info
end
function TypeInfo.createBuiltin( idName, typeTxt, kind )
  local typeId = typeIdSeed + 1
  if kind == TypeInfoKindRoot then
    typeId = rootTypeId
  else 
    typeIdSeed = typeIdSeed + 1
  end
  local info = TypeInfo.new(nil, nil, false, false, false, "pub", typeTxt, rootTypeInfo, typeId, kind)
  typeInfoKind[idName] = info
  builtInTypeMap[typeTxt] = info
  builtInTypeIdSet[info.typeId] = true
  return info
end
function TypeInfo.createList( accessMode, parentInfo, itemTypeInfo )
  if not itemTypeInfo or #itemTypeInfo == 0 then
    error( string.format( "illegal list type: %s", itemTypeInfo) )
  end
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(nil, nil, false, false, false, accessMode, "", rootTypeInfo, typeIdSeed, TypeInfoKindList, itemTypeInfo)
end
function TypeInfo.createArray( accessMode, parentInfo, itemTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(nil, nil, false, false, false, accessMode, "", rootTypeInfo, typeIdSeed, TypeInfoKindArray, itemTypeInfo)
end
function TypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return TypeInfo.new(nil, nil, false, false, false, accessMode, "Map", rootTypeInfo, typeIdSeed, TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end
function TypeInfo.createClass( baseInfo, parentInfo, externalFlag, accessMode, className )
  if className == "str" then
    return builtInTypeMap[className]
  end
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(baseInfo, nil, false, externalFlag, false, accessMode, className, parentInfo, typeIdSeed, TypeInfoKindClass)
  return info
end
function TypeInfo.createFunc( kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )
  typeIdSeed = typeIdSeed + 1
  local info = TypeInfo.new(nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, typeIdSeed, kind, argTypeList or {}, retTypeInfoList or {})
  return info
end
                function TypeInfo:get_itemTypeInfoList()
                return self.itemTypeInfoList
                end
                function TypeInfo:get_retTypeInfoList()
                return self.retTypeInfoList
                end
                function TypeInfo:get_parentInfo()
                return self.parentInfo
                end
                function TypeInfo:get_typeId()
                return self.typeId
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
                function TypeInfo:get_orgTypeInfo()
                return self.orgTypeInfo
                end
                function TypeInfo:get_baseTypeInfo()
                return self.baseTypeInfo
                end
                function TypeInfo:get_nilable()
                return self.nilable
                end
                function TypeInfo:get_nilableTypeInfo()
                return self.nilableTypeInfo
                end
                function TypeInfo:get_children()
                return self.children
                end

local typeInfoRoot = TypeInfo.createBuiltin( "Root", ":", TypeInfoKindRoot )
local typeInfoNone = TypeInfo.createBuiltin( "None", "", TypeInfoKindPrim )
local typeInfoStem = TypeInfo.createBuiltin( "Stem", "stem", TypeInfoKindPrim )
local typeInfoNil = TypeInfo.createBuiltin( "Nil", "nil", TypeInfoKindPrim )
local typeInfoBool = TypeInfo.createBuiltin( "Bool", "bool", TypeInfoKindPrim )
local typeInfoInt = TypeInfo.createBuiltin( "Int", "int", TypeInfoKindPrim )
local typeInfoReal = TypeInfo.createBuiltin( "Real", "real", TypeInfoKindPrim )
local typeInfoChar = TypeInfo.createBuiltin( "char", "char", TypeInfoKindPrim )
local typeInfoString = TypeInfo.createBuiltin( "String", "str", TypeInfoKindClass )
local typeInfoMap = TypeInfo.createBuiltin( "Map", "Map", TypeInfoKindMap )
local typeInfoList = TypeInfo.createBuiltin( "List", "List", TypeInfoKindList )
local typeInfoArray = TypeInfo.createBuiltin( "Array", "Array", TypeInfoKindArray )
local typeInfoForm = TypeInfo.createBuiltin( "Form", "form", TypeInfoKindFunc )
local typeInfoSymbol = TypeInfo.createBuiltin( "Symbol", "sym", TypeInfoKindPrim )
local typeInfoStat = TypeInfo.createBuiltin( "Stat", "stat", TypeInfoKindPrim )
function TypeInfo:isSettableFrom( other )
  if not other then
    return false
  end
  if self == typeInfoStem then
    return true
  end
  if self.typeId == other.typeId then
    return true
  end
  if self.kind ~= other.kind then
    return false
  end
  return true
end

local Scope = {}
moduleObj.Scope = Scope
function Scope.new( parent, classFlag, inheritList )
  local obj = {}
  setmetatable( obj, { __index = Scope } )
  if obj.__init then obj:__init( parent, classFlag, inheritList ); end
return obj
end
function Scope:__init(parent, classFlag, inheritList) 
  self.parent = parent
  self.symbol2TypeInfoMap = {}
  self.className2ScopeMap = {}
  self.inheritList = inheritList
  self.classFlag = classFlag
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
  local typeInfo = self.symbol2TypeInfoMap[name]
  if typeInfo then
    return typeInfo
  end
  if self.inheritList then
    for __index, scope in pairs( self.inheritList ) do
      typeInfo = scope:getTypeInfo( name )
      if typeInfo then
        return typeInfo
      end
    end
  end
  if self.parent then
    return self.parent:getTypeInfo( name )
  end
  return builtInTypeMap[name]
end
function Scope:getTypeInfoMethod( name, includeSelfFlag )
  if self.classFlag then
    local typeInfo = nil
    if includeSelfFlag then
      local typeInfo = self.symbol2TypeInfoMap[name]
      if typeInfo then
        return typeInfo
      end
    end
    if self.inheritList then
      for __index, scope in pairs( self.inheritList ) do
        typeInfo = scope:getTypeInfoMethod( name, true )
        if typeInfo then
          return typeInfo
        end
      end
    end
  end
  return nil
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

local Filter = {}
moduleObj.Filter = Filter
            function Filter.new(  )
            local obj = {}
            setmetatable( obj, { __index = Filter } )
            if obj.__init then
            obj:__init(  )
            end       
            return obj
            end         
            function Filter:__init(  )
            
end

local NodePos = {}
moduleObj.NodePos = NodePos
            function NodePos.new( lineNo, column )
            local obj = {}
            setmetatable( obj, { __index = NodePos } )
            if obj.__init then
            obj:__init( lineNo, column )
            end       
            return obj
            end         
            function NodePos:__init( lineNo, column )
            
self.lineNo = lineNo
  self.column = column
  end

local Node = {}
moduleObj.Node = Node
function Node:get_expType(  )
  if not self.expTypeList then
    return typeInfoNone
  end
  return self.expTypeList[1]
end
function Node:getLiteral(  )
  return nil
end
function Node:processFilter( filter, ... )
end
            function Node.new( kind, pos, expTypeList )
            local obj = {}
            setmetatable( obj, { __index = Node } )
            if obj.__init then
            obj:__init( kind, pos, expTypeList )
            end       
            return obj
            end         
            function Node:__init( kind, pos, expTypeList )
            
self.kind = kind
  self.pos = pos
  self.expTypeList = expTypeList
  end
                function Node:get_kind()
                return self.kind
                end
                function Node:get_expTypeList()
                return self.expTypeList
                end

local NamespaceInfo = {}
moduleObj.NamespaceInfo = NamespaceInfo
            function NamespaceInfo.new( name, scope, typeInfo )
            local obj = {}
            setmetatable( obj, { __index = NamespaceInfo } )
            if obj.__init then
            obj:__init( name, scope, typeInfo )
            end       
            return obj
            end         
            function NamespaceInfo:__init( name, scope, typeInfo )
            
self.name = name
  self.scope = scope
  self.typeInfo = typeInfo
  end

-- none

local DeclMacroInfo = {}
moduleObj.DeclMacroInfo = DeclMacroInfo
            function DeclMacroInfo.new( name, argList, ast, tokenList )
            local obj = {}
            setmetatable( obj, { __index = DeclMacroInfo } )
            if obj.__init then
            obj:__init( name, argList, ast, tokenList )
            end       
            return obj
            end         
            function DeclMacroInfo:__init( name, argList, ast, tokenList )
            
self.name = name
  self.argList = argList
  self.ast = ast
  self.tokenList = tokenList
  end
                function DeclMacroInfo:get_name()
                return self.name
                end
                function DeclMacroInfo:get_argList()
                return self.argList
                end
                function DeclMacroInfo:get_ast()
                return self.ast
                end
                function DeclMacroInfo:get_tokenList()
                return self.tokenList
                end

local MacroValInfo = {}
            function MacroValInfo.new( val, typeInfo )
            local obj = {}
            setmetatable( obj, { __index = MacroValInfo } )
            if obj.__init then
            obj:__init( val, typeInfo )
            end       
            return obj
            end         
            function MacroValInfo:__init( val, typeInfo )
            
self.val = val
  self.typeInfo = typeInfo
  end

local MacroInfo = {}
            function MacroInfo.new( func, declInfo, symbol2MacroValInfoMap )
            local obj = {}
            setmetatable( obj, { __index = MacroInfo } )
            if obj.__init then
            obj:__init( func, declInfo, symbol2MacroValInfoMap )
            end       
            return obj
            end         
            function MacroInfo:__init( func, declInfo, symbol2MacroValInfoMap )
            
self.func = func
  self.declInfo = declInfo
  self.symbol2MacroValInfoMap = symbol2MacroValInfoMap
  end

local TransUnit = {}
moduleObj.TransUnit = TransUnit
function TransUnit.new( macroEval )
  local obj = {}
  setmetatable( obj, { __index = TransUnit } )
  if obj.__init then obj:__init( macroEval ); end
return obj
end
function TransUnit:__init(macroEval) 
  self.pushbackList = {}
  self.scope = Scope.new(nil, false)
  self.namespaceList = {typeInfoRoot}
  self.classList = {}
  self.typeId2ClassMap = {}
  self.typeId2Scope = {}
  self.typeInfo2ClassNode = {}
  self.currentToken = nil
  self.errMessList = {}
  self.macroEval = macroEval
  self.typeId2MacroInfo = {}
  self.macroMode = "none"
  self.symbol2ValueMapForMacro = {}
end
function TransUnit:addErrMess( pos, mess )
  table.insert( self.errMessList, string.format( "%s:%d:%d: %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, mess) )
end
function TransUnit:pushScope( classFlag, inheritList )
  self.scope = Scope.new(self.scope, classFlag, inheritList)
  return self.scope
end
function TransUnit:popScope(  )
  self.scope = self.scope:get_parent(  )
end
function TransUnit:pushNamespace( name, typeInfo, scope )
  local namespace = NamespaceInfo.new(name, scope, typeInfo)
  table.insert( self.namespaceList, namespace )
end
function TransUnit:popNamespace(  )
  table.remove( self.namespaceList )
end
function TransUnit:getCurrentClass(  )
  if #self.classList == 0 then
    return rootTypeInfo
  end
  local classInfo = self.classList[#self.classList]
  return classInfo.typeInfo
end
function TransUnit:getCurrentNamespaceTypeInfo(  )
  return self.namespaceList[#self.namespaceList].typeInfo
end
function TransUnit:pushClass( baseInfo, externalFlag, name, accessMode )
  local typeInfo = self.scope:getTypeInfoChild( name )
  if not typeInfo then
    local parentInfo = self:getCurrentNamespaceTypeInfo(  )
    typeInfo = TypeInfo.createClass( baseInfo, parentInfo, externalFlag, accessMode, name )
    local inheritList = nil
    if baseInfo then
      inheritList = {self.typeId2Scope[baseInfo:get_typeId(  )]}
    end
    local scope = self:pushScope( true, inheritList )
    scope:get_parent(  ):addClass( name, typeInfo, scope )
  else 
    self.scope = self.scope:getClassScope( name )
  end
  local namespace = NamespaceInfo.new(name, self.scope, typeInfo)
  table.insert( self.namespaceList, namespace )
  table.insert( self.classList, namespace )
  self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
  self.typeId2Scope[typeInfo:get_typeId(  )] = self.scope
  return typeInfo
end
function TransUnit:popClass(  )
  self:popScope(  )
  table.remove( self.namespaceList )
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
-- none
-- none
                function TransUnit:get_errMessList()
                return self.errMessList
                end

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
regOpLevel( 1, {"`", ",,", ",,,", ",,,,"} )
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

local function getNodeKindName( kind )
  return nodeKind2NameMap[kind]
end
moduleObj.getNodeKindName = getNodeKindName

-- none

function Filter:processNone( node, ... )
end

-- none

-- none

local nodeKindNone = regKind( [[None]] )
local NoneNode = {}
setmetatable( NoneNode, { __index = Node } )
moduleObj.NoneNode = NoneNode
function NoneNode:processFilter( filter, ... )
  local argList = {...}
  filter:processNone( self, table.unpack( argList ) )
end
function NoneNode.new( pos, typeInfoList )
  local obj = {}
  setmetatable( obj, { __index = NoneNode } )
  if obj.__init then obj:__init( pos, typeInfoList ); end
return obj
end
function NoneNode:__init(pos, typeInfoList) 
  Node.__init( self, nodeKindNone, pos, typeInfoList)
  
  -- none
  
  -- none
  
end


-- none

function Filter:processImport( node, ... )
end

-- none

-- none

local nodeKindImport = regKind( [[Import]] )
local ImportNode = {}
setmetatable( ImportNode, { __index = Node } )
moduleObj.ImportNode = ImportNode
function ImportNode:processFilter( filter, ... )
  local argList = {...}
  filter:processImport( self, table.unpack( argList ) )
end
function ImportNode.new( pos, typeInfoList, modulePath )
  local obj = {}
  setmetatable( obj, { __index = ImportNode } )
  if obj.__init then obj:__init( pos, typeInfoList, modulePath ); end
return obj
end
function ImportNode:__init(pos, typeInfoList, modulePath) 
  Node.__init( self, nodeKindImport, pos, typeInfoList)
  
  -- none
  
  self.modulePath = modulePath
  -- none
  
end
                function ImportNode:get_modulePath()
                return self.modulePath
                end


-- none

function Filter:processRoot( node, ... )
end

-- none

-- none

local nodeKindRoot = regKind( [[Root]] )
local RootNode = {}
setmetatable( RootNode, { __index = Node } )
moduleObj.RootNode = RootNode
function RootNode:processFilter( filter, ... )
  local argList = {...}
  filter:processRoot( self, table.unpack( argList ) )
end
function RootNode.new( pos, typeInfoList, children, typeId2ClassMap )
  local obj = {}
  setmetatable( obj, { __index = RootNode } )
  if obj.__init then obj:__init( pos, typeInfoList, children, typeId2ClassMap ); end
return obj
end
function RootNode:__init(pos, typeInfoList, children, typeId2ClassMap) 
  Node.__init( self, nodeKindRoot, pos, typeInfoList)
  
  -- none
  
  self.children = children
  self.typeId2ClassMap = typeId2ClassMap
  -- none
  
end
                function RootNode:get_children()
                return self.children
                end
                function RootNode:get_typeId2ClassMap()
                return self.typeId2ClassMap
                end


-- none

function Filter:processRefType( node, ... )
end

-- none

-- none

local nodeKindRefType = regKind( [[RefType]] )
local RefTypeNode = {}
setmetatable( RefTypeNode, { __index = Node } )
moduleObj.RefTypeNode = RefTypeNode
function RefTypeNode:processFilter( filter, ... )
  local argList = {...}
  filter:processRefType( self, table.unpack( argList ) )
end
function RefTypeNode.new( pos, typeInfoList, name, refFlag, mutFlag, array )
  local obj = {}
  setmetatable( obj, { __index = RefTypeNode } )
  if obj.__init then obj:__init( pos, typeInfoList, name, refFlag, mutFlag, array ); end
return obj
end
function RefTypeNode:__init(pos, typeInfoList, name, refFlag, mutFlag, array) 
  Node.__init( self, nodeKindRefType, pos, typeInfoList)
  
  -- none
  
  self.name = name
  self.refFlag = refFlag
  self.mutFlag = mutFlag
  self.array = array
  -- none
  
end
                function RefTypeNode:get_name()
                return self.name
                end
                function RefTypeNode:get_refFlag()
                return self.refFlag
                end
                function RefTypeNode:get_mutFlag()
                return self.mutFlag
                end
                function RefTypeNode:get_array()
                return self.array
                end


-- none

function Filter:processBlock( node, ... )
end

-- none

-- none

local nodeKindBlock = regKind( [[Block]] )
local BlockNode = {}
setmetatable( BlockNode, { __index = Node } )
moduleObj.BlockNode = BlockNode
function BlockNode:processFilter( filter, ... )
  local argList = {...}
  filter:processBlock( self, table.unpack( argList ) )
end
function BlockNode.new( pos, typeInfoList, blockKind, stmtList )
  local obj = {}
  setmetatable( obj, { __index = BlockNode } )
  if obj.__init then obj:__init( pos, typeInfoList, blockKind, stmtList ); end
return obj
end
function BlockNode:__init(pos, typeInfoList, blockKind, stmtList) 
  Node.__init( self, nodeKindBlock, pos, typeInfoList)
  
  -- none
  
  self.blockKind = blockKind
  self.stmtList = stmtList
  -- none
  
end
                function BlockNode:get_blockKind()
                return self.blockKind
                end
                function BlockNode:get_stmtList()
                return self.stmtList
                end


local IfStmtInfo = {}
moduleObj.IfStmtInfo = IfStmtInfo
            function IfStmtInfo.new( kind, exp, block )
            local obj = {}
            setmetatable( obj, { __index = IfStmtInfo } )
            if obj.__init then
            obj:__init( kind, exp, block )
            end       
            return obj
            end         
            function IfStmtInfo:__init( kind, exp, block )
            
self.kind = kind
  self.exp = exp
  self.block = block
  end
                function IfStmtInfo:get_kind()
                return self.kind
                end
                function IfStmtInfo:get_exp()
                return self.exp
                end
                function IfStmtInfo:get_block()
                return self.block
                end

-- none

function Filter:processIf( node, ... )
end

-- none

-- none

local nodeKindIf = regKind( [[If]] )
local IfNode = {}
setmetatable( IfNode, { __index = Node } )
moduleObj.IfNode = IfNode
function IfNode:processFilter( filter, ... )
  local argList = {...}
  filter:processIf( self, table.unpack( argList ) )
end
function IfNode.new( pos, typeInfoList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = IfNode } )
  if obj.__init then obj:__init( pos, typeInfoList, stmtList ); end
return obj
end
function IfNode:__init(pos, typeInfoList, stmtList) 
  Node.__init( self, nodeKindIf, pos, typeInfoList)
  
  -- none
  
  self.stmtList = stmtList
  -- none
  
end
                function IfNode:get_stmtList()
                return self.stmtList
                end


-- none

function Filter:processExpList( node, ... )
end

-- none

-- none

local nodeKindExpList = regKind( [[ExpList]] )
local ExpListNode = {}
setmetatable( ExpListNode, { __index = Node } )
moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpList( self, table.unpack( argList ) )
end
function ExpListNode.new( pos, typeInfoList, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpListNode } )
  if obj.__init then obj:__init( pos, typeInfoList, expList ); end
return obj
end
function ExpListNode:__init(pos, typeInfoList, expList) 
  Node.__init( self, nodeKindExpList, pos, typeInfoList)
  
  -- none
  
  self.expList = expList
  -- none
  
end
                function ExpListNode:get_expList()
                return self.expList
                end


local CaseInfo = {}
moduleObj.CaseInfo = CaseInfo
            function CaseInfo.new( expList, block )
            local obj = {}
            setmetatable( obj, { __index = CaseInfo } )
            if obj.__init then
            obj:__init( expList, block )
            end       
            return obj
            end         
            function CaseInfo:__init( expList, block )
            
self.expList = expList
  self.block = block
  end
                function CaseInfo:get_expList()
                return self.expList
                end
                function CaseInfo:get_block()
                return self.block
                end

-- none

function Filter:processSwitch( node, ... )
end

-- none

-- none

local nodeKindSwitch = regKind( [[Switch]] )
local SwitchNode = {}
setmetatable( SwitchNode, { __index = Node } )
moduleObj.SwitchNode = SwitchNode
function SwitchNode:processFilter( filter, ... )
  local argList = {...}
  filter:processSwitch( self, table.unpack( argList ) )
end
function SwitchNode.new( pos, typeInfoList, exp, caseList, default )
  local obj = {}
  setmetatable( obj, { __index = SwitchNode } )
  if obj.__init then obj:__init( pos, typeInfoList, exp, caseList, default ); end
return obj
end
function SwitchNode:__init(pos, typeInfoList, exp, caseList, default) 
  Node.__init( self, nodeKindSwitch, pos, typeInfoList)
  
  -- none
  
  self.exp = exp
  self.caseList = caseList
  self.default = default
  -- none
  
end
                function SwitchNode:get_exp()
                return self.exp
                end
                function SwitchNode:get_caseList()
                return self.caseList
                end
                function SwitchNode:get_default()
                return self.default
                end


-- none

function Filter:processWhile( node, ... )
end

-- none

-- none

local nodeKindWhile = regKind( [[While]] )
local WhileNode = {}
setmetatable( WhileNode, { __index = Node } )
moduleObj.WhileNode = WhileNode
function WhileNode:processFilter( filter, ... )
  local argList = {...}
  filter:processWhile( self, table.unpack( argList ) )
end
function WhileNode.new( pos, typeInfoList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = WhileNode } )
  if obj.__init then obj:__init( pos, typeInfoList, exp, block ); end
return obj
end
function WhileNode:__init(pos, typeInfoList, exp, block) 
  Node.__init( self, nodeKindWhile, pos, typeInfoList)
  
  -- none
  
  self.exp = exp
  self.block = block
  -- none
  
end
                function WhileNode:get_exp()
                return self.exp
                end
                function WhileNode:get_block()
                return self.block
                end


-- none

function Filter:processRepeat( node, ... )
end

-- none

-- none

local nodeKindRepeat = regKind( [[Repeat]] )
local RepeatNode = {}
setmetatable( RepeatNode, { __index = Node } )
moduleObj.RepeatNode = RepeatNode
function RepeatNode:processFilter( filter, ... )
  local argList = {...}
  filter:processRepeat( self, table.unpack( argList ) )
end
function RepeatNode.new( pos, typeInfoList, block, exp )
  local obj = {}
  setmetatable( obj, { __index = RepeatNode } )
  if obj.__init then obj:__init( pos, typeInfoList, block, exp ); end
return obj
end
function RepeatNode:__init(pos, typeInfoList, block, exp) 
  Node.__init( self, nodeKindRepeat, pos, typeInfoList)
  
  -- none
  
  self.block = block
  self.exp = exp
  -- none
  
end
                function RepeatNode:get_block()
                return self.block
                end
                function RepeatNode:get_exp()
                return self.exp
                end


-- none

function Filter:processFor( node, ... )
end

-- none

-- none

local nodeKindFor = regKind( [[For]] )
local ForNode = {}
setmetatable( ForNode, { __index = Node } )
moduleObj.ForNode = ForNode
function ForNode:processFilter( filter, ... )
  local argList = {...}
  filter:processFor( self, table.unpack( argList ) )
end
function ForNode.new( pos, typeInfoList, block, val, init, to, delta )
  local obj = {}
  setmetatable( obj, { __index = ForNode } )
  if obj.__init then obj:__init( pos, typeInfoList, block, val, init, to, delta ); end
return obj
end
function ForNode:__init(pos, typeInfoList, block, val, init, to, delta) 
  Node.__init( self, nodeKindFor, pos, typeInfoList)
  
  -- none
  
  self.block = block
  self.val = val
  self.init = init
  self.to = to
  self.delta = delta
  -- none
  
end
                function ForNode:get_block()
                return self.block
                end
                function ForNode:get_val()
                return self.val
                end
                function ForNode:get_init()
                return self.init
                end
                function ForNode:get_to()
                return self.to
                end
                function ForNode:get_delta()
                return self.delta
                end


-- none

function Filter:processApply( node, ... )
end

-- none

-- none

local nodeKindApply = regKind( [[Apply]] )
local ApplyNode = {}
setmetatable( ApplyNode, { __index = Node } )
moduleObj.ApplyNode = ApplyNode
function ApplyNode:processFilter( filter, ... )
  local argList = {...}
  filter:processApply( self, table.unpack( argList ) )
end
function ApplyNode.new( pos, typeInfoList, varList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ApplyNode } )
  if obj.__init then obj:__init( pos, typeInfoList, varList, exp, block ); end
return obj
end
function ApplyNode:__init(pos, typeInfoList, varList, exp, block) 
  Node.__init( self, nodeKindApply, pos, typeInfoList)
  
  -- none
  
  self.varList = varList
  self.exp = exp
  self.block = block
  -- none
  
end
                function ApplyNode:get_varList()
                return self.varList
                end
                function ApplyNode:get_exp()
                return self.exp
                end
                function ApplyNode:get_block()
                return self.block
                end


-- none

function Filter:processForeach( node, ... )
end

-- none

-- none

local nodeKindForeach = regKind( [[Foreach]] )
local ForeachNode = {}
setmetatable( ForeachNode, { __index = Node } )
moduleObj.ForeachNode = ForeachNode
function ForeachNode:processFilter( filter, ... )
  local argList = {...}
  filter:processForeach( self, table.unpack( argList ) )
end
function ForeachNode.new( pos, typeInfoList, val, key, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ForeachNode } )
  if obj.__init then obj:__init( pos, typeInfoList, val, key, exp, block ); end
return obj
end
function ForeachNode:__init(pos, typeInfoList, val, key, exp, block) 
  Node.__init( self, nodeKindForeach, pos, typeInfoList)
  
  -- none
  
  self.val = val
  self.key = key
  self.exp = exp
  self.block = block
  -- none
  
end
                function ForeachNode:get_val()
                return self.val
                end
                function ForeachNode:get_key()
                return self.key
                end
                function ForeachNode:get_exp()
                return self.exp
                end
                function ForeachNode:get_block()
                return self.block
                end


-- none

function Filter:processForsort( node, ... )
end

-- none

-- none

local nodeKindForsort = regKind( [[Forsort]] )
local ForsortNode = {}
setmetatable( ForsortNode, { __index = Node } )
moduleObj.ForsortNode = ForsortNode
function ForsortNode:processFilter( filter, ... )
  local argList = {...}
  filter:processForsort( self, table.unpack( argList ) )
end
function ForsortNode.new( pos, typeInfoList, val, key, exp, block, sort )
  local obj = {}
  setmetatable( obj, { __index = ForsortNode } )
  if obj.__init then obj:__init( pos, typeInfoList, val, key, exp, block, sort ); end
return obj
end
function ForsortNode:__init(pos, typeInfoList, val, key, exp, block, sort) 
  Node.__init( self, nodeKindForsort, pos, typeInfoList)
  
  -- none
  
  self.val = val
  self.key = key
  self.exp = exp
  self.block = block
  self.sort = sort
  -- none
  
end
                function ForsortNode:get_val()
                return self.val
                end
                function ForsortNode:get_key()
                return self.key
                end
                function ForsortNode:get_exp()
                return self.exp
                end
                function ForsortNode:get_block()
                return self.block
                end
                function ForsortNode:get_sort()
                return self.sort
                end


-- none

function Filter:processReturn( node, ... )
end

-- none

-- none

local nodeKindReturn = regKind( [[Return]] )
local ReturnNode = {}
setmetatable( ReturnNode, { __index = Node } )
moduleObj.ReturnNode = ReturnNode
function ReturnNode:processFilter( filter, ... )
  local argList = {...}
  filter:processReturn( self, table.unpack( argList ) )
end
function ReturnNode.new( pos, typeInfoList, expList )
  local obj = {}
  setmetatable( obj, { __index = ReturnNode } )
  if obj.__init then obj:__init( pos, typeInfoList, expList ); end
return obj
end
function ReturnNode:__init(pos, typeInfoList, expList) 
  Node.__init( self, nodeKindReturn, pos, typeInfoList)
  
  -- none
  
  self.expList = expList
  -- none
  
end
                function ReturnNode:get_expList()
                return self.expList
                end


-- none

function Filter:processBreak( node, ... )
end

-- none

-- none

local nodeKindBreak = regKind( [[Break]] )
local BreakNode = {}
setmetatable( BreakNode, { __index = Node } )
moduleObj.BreakNode = BreakNode
function BreakNode:processFilter( filter, ... )
  local argList = {...}
  filter:processBreak( self, table.unpack( argList ) )
end
function BreakNode.new( pos, typeInfoList )
  local obj = {}
  setmetatable( obj, { __index = BreakNode } )
  if obj.__init then obj:__init( pos, typeInfoList ); end
return obj
end
function BreakNode:__init(pos, typeInfoList) 
  Node.__init( self, nodeKindBreak, pos, typeInfoList)
  
  -- none
  
  -- none
  
end


-- none

function Filter:processExpNew( node, ... )
end

-- none

-- none

local nodeKindExpNew = regKind( [[ExpNew]] )
local ExpNewNode = {}
setmetatable( ExpNewNode, { __index = Node } )
moduleObj.ExpNewNode = ExpNewNode
function ExpNewNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpNew( self, table.unpack( argList ) )
end
function ExpNewNode.new( pos, typeInfoList, symbol, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpNewNode } )
  if obj.__init then obj:__init( pos, typeInfoList, symbol, argList ); end
return obj
end
function ExpNewNode:__init(pos, typeInfoList, symbol, argList) 
  Node.__init( self, nodeKindExpNew, pos, typeInfoList)
  
  -- none
  
  self.symbol = symbol
  self.argList = argList
  -- none
  
end
                function ExpNewNode:get_symbol()
                return self.symbol
                end
                function ExpNewNode:get_argList()
                return self.argList
                end


-- none

function Filter:processExpRef( node, ... )
end

-- none

-- none

local nodeKindExpRef = regKind( [[ExpRef]] )
local ExpRefNode = {}
setmetatable( ExpRefNode, { __index = Node } )
moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpRef( self, table.unpack( argList ) )
end
function ExpRefNode.new( pos, typeInfoList, token )
  local obj = {}
  setmetatable( obj, { __index = ExpRefNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token ); end
return obj
end
function ExpRefNode:__init(pos, typeInfoList, token) 
  Node.__init( self, nodeKindExpRef, pos, typeInfoList)
  
  -- none
  
  self.token = token
  -- none
  
end
                function ExpRefNode:get_token()
                return self.token
                end


-- none

function Filter:processExpOp2( node, ... )
end

-- none

-- none

local nodeKindExpOp2 = regKind( [[ExpOp2]] )
local ExpOp2Node = {}
setmetatable( ExpOp2Node, { __index = Node } )
moduleObj.ExpOp2Node = ExpOp2Node
function ExpOp2Node:processFilter( filter, ... )
  local argList = {...}
  filter:processExpOp2( self, table.unpack( argList ) )
end
function ExpOp2Node.new( pos, typeInfoList, op, exp1, exp2 )
  local obj = {}
  setmetatable( obj, { __index = ExpOp2Node } )
  if obj.__init then obj:__init( pos, typeInfoList, op, exp1, exp2 ); end
return obj
end
function ExpOp2Node:__init(pos, typeInfoList, op, exp1, exp2) 
  Node.__init( self, nodeKindExpOp2, pos, typeInfoList)
  
  -- none
  
  self.op = op
  self.exp1 = exp1
  self.exp2 = exp2
  -- none
  
end
                function ExpOp2Node:get_op()
                return self.op
                end
                function ExpOp2Node:get_exp1()
                return self.exp1
                end
                function ExpOp2Node:get_exp2()
                return self.exp2
                end


-- none

function Filter:processExpCast( node, ... )
end

-- none

-- none

local nodeKindExpCast = regKind( [[ExpCast]] )
local ExpCastNode = {}
setmetatable( ExpCastNode, { __index = Node } )
moduleObj.ExpCastNode = ExpCastNode
function ExpCastNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpCast( self, table.unpack( argList ) )
end
function ExpCastNode.new( pos, typeInfoList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpCastNode } )
  if obj.__init then obj:__init( pos, typeInfoList, exp ); end
return obj
end
function ExpCastNode:__init(pos, typeInfoList, exp) 
  Node.__init( self, nodeKindExpCast, pos, typeInfoList)
  
  -- none
  
  self.exp = exp
  -- none
  
end
                function ExpCastNode:get_exp()
                return self.exp
                end


-- none

function Filter:processExpOp1( node, ... )
end

-- none

-- none

local nodeKindExpOp1 = regKind( [[ExpOp1]] )
local ExpOp1Node = {}
setmetatable( ExpOp1Node, { __index = Node } )
moduleObj.ExpOp1Node = ExpOp1Node
function ExpOp1Node:processFilter( filter, ... )
  local argList = {...}
  filter:processExpOp1( self, table.unpack( argList ) )
end
function ExpOp1Node.new( pos, typeInfoList, op, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpOp1Node } )
  if obj.__init then obj:__init( pos, typeInfoList, op, exp ); end
return obj
end
function ExpOp1Node:__init(pos, typeInfoList, op, exp) 
  Node.__init( self, nodeKindExpOp1, pos, typeInfoList)
  
  -- none
  
  self.op = op
  self.exp = exp
  -- none
  
end
                function ExpOp1Node:get_op()
                return self.op
                end
                function ExpOp1Node:get_exp()
                return self.exp
                end


-- none

function Filter:processExpRefItem( node, ... )
end

-- none

-- none

local nodeKindExpRefItem = regKind( [[ExpRefItem]] )
local ExpRefItemNode = {}
setmetatable( ExpRefItemNode, { __index = Node } )
moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpRefItem( self, table.unpack( argList ) )
end
function ExpRefItemNode.new( pos, typeInfoList, val, index )
  local obj = {}
  setmetatable( obj, { __index = ExpRefItemNode } )
  if obj.__init then obj:__init( pos, typeInfoList, val, index ); end
return obj
end
function ExpRefItemNode:__init(pos, typeInfoList, val, index) 
  Node.__init( self, nodeKindExpRefItem, pos, typeInfoList)
  
  -- none
  
  self.val = val
  self.index = index
  -- none
  
end
                function ExpRefItemNode:get_val()
                return self.val
                end
                function ExpRefItemNode:get_index()
                return self.index
                end


-- none

function Filter:processExpCall( node, ... )
end

-- none

-- none

local nodeKindExpCall = regKind( [[ExpCall]] )
local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpCall( self, table.unpack( argList ) )
end
function ExpCallNode.new( pos, typeInfoList, func, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallNode } )
  if obj.__init then obj:__init( pos, typeInfoList, func, argList ); end
return obj
end
function ExpCallNode:__init(pos, typeInfoList, func, argList) 
  Node.__init( self, nodeKindExpCall, pos, typeInfoList)
  
  -- none
  
  self.func = func
  self.argList = argList
  -- none
  
end
                function ExpCallNode:get_func()
                return self.func
                end
                function ExpCallNode:get_argList()
                return self.argList
                end


-- none

function Filter:processExpDDD( node, ... )
end

-- none

-- none

local nodeKindExpDDD = regKind( [[ExpDDD]] )
local ExpDDDNode = {}
setmetatable( ExpDDDNode, { __index = Node } )
moduleObj.ExpDDDNode = ExpDDDNode
function ExpDDDNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpDDD( self, table.unpack( argList ) )
end
function ExpDDDNode.new( pos, typeInfoList, token )
  local obj = {}
  setmetatable( obj, { __index = ExpDDDNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token ); end
return obj
end
function ExpDDDNode:__init(pos, typeInfoList, token) 
  Node.__init( self, nodeKindExpDDD, pos, typeInfoList)
  
  -- none
  
  self.token = token
  -- none
  
end
                function ExpDDDNode:get_token()
                return self.token
                end


-- none

function Filter:processExpParen( node, ... )
end

-- none

-- none

local nodeKindExpParen = regKind( [[ExpParen]] )
local ExpParenNode = {}
setmetatable( ExpParenNode, { __index = Node } )
moduleObj.ExpParenNode = ExpParenNode
function ExpParenNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpParen( self, table.unpack( argList ) )
end
function ExpParenNode.new( pos, typeInfoList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpParenNode } )
  if obj.__init then obj:__init( pos, typeInfoList, exp ); end
return obj
end
function ExpParenNode:__init(pos, typeInfoList, exp) 
  Node.__init( self, nodeKindExpParen, pos, typeInfoList)
  
  -- none
  
  self.exp = exp
  -- none
  
end
                function ExpParenNode:get_exp()
                return self.exp
                end


-- none

function Filter:processExpMacroExp( node, ... )
end

-- none

-- none

local nodeKindExpMacroExp = regKind( [[ExpMacroExp]] )
local ExpMacroExpNode = {}
setmetatable( ExpMacroExpNode, { __index = Node } )
moduleObj.ExpMacroExpNode = ExpMacroExpNode
function ExpMacroExpNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpMacroExp( self, table.unpack( argList ) )
end
function ExpMacroExpNode.new( pos, typeInfoList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroExpNode } )
  if obj.__init then obj:__init( pos, typeInfoList, stmtList ); end
return obj
end
function ExpMacroExpNode:__init(pos, typeInfoList, stmtList) 
  Node.__init( self, nodeKindExpMacroExp, pos, typeInfoList)
  
  -- none
  
  self.stmtList = stmtList
  -- none
  
end
                function ExpMacroExpNode:get_stmtList()
                return self.stmtList
                end


-- none

function Filter:processExpMacroStat( node, ... )
end

-- none

-- none

local nodeKindExpMacroStat = regKind( [[ExpMacroStat]] )
local ExpMacroStatNode = {}
setmetatable( ExpMacroStatNode, { __index = Node } )
moduleObj.ExpMacroStatNode = ExpMacroStatNode
function ExpMacroStatNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpMacroStat( self, table.unpack( argList ) )
end
function ExpMacroStatNode.new( pos, typeInfoList, expStrList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroStatNode } )
  if obj.__init then obj:__init( pos, typeInfoList, expStrList ); end
return obj
end
function ExpMacroStatNode:__init(pos, typeInfoList, expStrList) 
  Node.__init( self, nodeKindExpMacroStat, pos, typeInfoList)
  
  -- none
  
  self.expStrList = expStrList
  -- none
  
end
                function ExpMacroStatNode:get_expStrList()
                return self.expStrList
                end


-- none

function Filter:processStmtExp( node, ... )
end

-- none

-- none

local nodeKindStmtExp = regKind( [[StmtExp]] )
local StmtExpNode = {}
setmetatable( StmtExpNode, { __index = Node } )
moduleObj.StmtExpNode = StmtExpNode
function StmtExpNode:processFilter( filter, ... )
  local argList = {...}
  filter:processStmtExp( self, table.unpack( argList ) )
end
function StmtExpNode.new( pos, typeInfoList, exp )
  local obj = {}
  setmetatable( obj, { __index = StmtExpNode } )
  if obj.__init then obj:__init( pos, typeInfoList, exp ); end
return obj
end
function StmtExpNode:__init(pos, typeInfoList, exp) 
  Node.__init( self, nodeKindStmtExp, pos, typeInfoList)
  
  -- none
  
  self.exp = exp
  -- none
  
end
                function StmtExpNode:get_exp()
                return self.exp
                end


-- none

function Filter:processRefField( node, ... )
end

-- none

-- none

local nodeKindRefField = regKind( [[RefField]] )
local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, ... )
  local argList = {...}
  filter:processRefField( self, table.unpack( argList ) )
end
function RefFieldNode.new( pos, typeInfoList, field, prefix )
  local obj = {}
  setmetatable( obj, { __index = RefFieldNode } )
  if obj.__init then obj:__init( pos, typeInfoList, field, prefix ); end
return obj
end
function RefFieldNode:__init(pos, typeInfoList, field, prefix) 
  Node.__init( self, nodeKindRefField, pos, typeInfoList)
  
  -- none
  
  self.field = field
  self.prefix = prefix
  -- none
  
end
                function RefFieldNode:get_field()
                return self.field
                end
                function RefFieldNode:get_prefix()
                return self.prefix
                end


-- none

function Filter:processGetField( node, ... )
end

-- none

-- none

local nodeKindGetField = regKind( [[GetField]] )
local GetFieldNode = {}
setmetatable( GetFieldNode, { __index = Node } )
moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, ... )
  local argList = {...}
  filter:processGetField( self, table.unpack( argList ) )
end
function GetFieldNode.new( pos, typeInfoList, field, prefix, getterTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = GetFieldNode } )
  if obj.__init then obj:__init( pos, typeInfoList, field, prefix, getterTypeInfo ); end
return obj
end
function GetFieldNode:__init(pos, typeInfoList, field, prefix, getterTypeInfo) 
  Node.__init( self, nodeKindGetField, pos, typeInfoList)
  
  -- none
  
  self.field = field
  self.prefix = prefix
  self.getterTypeInfo = getterTypeInfo
  -- none
  
end
                function GetFieldNode:get_field()
                return self.field
                end
                function GetFieldNode:get_prefix()
                return self.prefix
                end
                function GetFieldNode:get_getterTypeInfo()
                return self.getterTypeInfo
                end


local VarInfo = {}
moduleObj.VarInfo = VarInfo
            function VarInfo.new( name, refType, actualType )
            local obj = {}
            setmetatable( obj, { __index = VarInfo } )
            if obj.__init then
            obj:__init( name, refType, actualType )
            end       
            return obj
            end         
            function VarInfo:__init( name, refType, actualType )
            
self.name = name
  self.refType = refType
  self.actualType = actualType
  end
                function VarInfo:get_name()
                return self.name
                end
                function VarInfo:get_refType()
                return self.refType
                end
                function VarInfo:get_actualType()
                return self.actualType
                end

-- none

function Filter:processDeclVar( node, ... )
end

-- none

-- none

local nodeKindDeclVar = regKind( [[DeclVar]] )
local DeclVarNode = {}
setmetatable( DeclVarNode, { __index = Node } )
moduleObj.DeclVarNode = DeclVarNode
function DeclVarNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclVar( self, table.unpack( argList ) )
end
function DeclVarNode.new( pos, typeInfoList, accessMode, staticFlag, varList, expList, typeInfoList, unwrap )
  local obj = {}
  setmetatable( obj, { __index = DeclVarNode } )
  if obj.__init then obj:__init( pos, typeInfoList, accessMode, staticFlag, varList, expList, typeInfoList, unwrap ); end
return obj
end
function DeclVarNode:__init(pos, typeInfoList, accessMode, staticFlag, varList, expList, typeInfoList, unwrap) 
  Node.__init( self, nodeKindDeclVar, pos, typeInfoList)
  
  -- none
  
  self.accessMode = accessMode
  self.staticFlag = staticFlag
  self.varList = varList
  self.expList = expList
  self.typeInfoList = typeInfoList
  self.unwrap = unwrap
  -- none
  
end
                function DeclVarNode:get_accessMode()
                return self.accessMode
                end
                function DeclVarNode:get_staticFlag()
                return self.staticFlag
                end
                function DeclVarNode:get_varList()
                return self.varList
                end
                function DeclVarNode:get_expList()
                return self.expList
                end
                function DeclVarNode:get_typeInfoList()
                return self.typeInfoList
                end
                function DeclVarNode:get_unwrap()
                return self.unwrap
                end


local DeclFuncInfo = {}
moduleObj.DeclFuncInfo = DeclFuncInfo
            function DeclFuncInfo.new( className, name, argList, staticFlag, accessMode, body, retTypeList, retTypeInfoList )
            local obj = {}
            setmetatable( obj, { __index = DeclFuncInfo } )
            if obj.__init then
            obj:__init( className, name, argList, staticFlag, accessMode, body, retTypeList, retTypeInfoList )
            end       
            return obj
            end         
            function DeclFuncInfo:__init( className, name, argList, staticFlag, accessMode, body, retTypeList, retTypeInfoList )
            
self.className = className
  self.name = name
  self.argList = argList
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.body = body
  self.retTypeList = retTypeList
  self.retTypeInfoList = retTypeInfoList
  end
                function DeclFuncInfo:get_className()
                return self.className
                end
                function DeclFuncInfo:get_name()
                return self.name
                end
                function DeclFuncInfo:get_argList()
                return self.argList
                end
                function DeclFuncInfo:get_staticFlag()
                return self.staticFlag
                end
                function DeclFuncInfo:get_accessMode()
                return self.accessMode
                end
                function DeclFuncInfo:get_body()
                return self.body
                end
                function DeclFuncInfo:get_retTypeList()
                return self.retTypeList
                end
                function DeclFuncInfo:get_retTypeInfoList()
                return self.retTypeInfoList
                end

-- none

function Filter:processDeclFunc( node, ... )
end

-- none

-- none

local nodeKindDeclFunc = regKind( [[DeclFunc]] )
local DeclFuncNode = {}
setmetatable( DeclFuncNode, { __index = Node } )
moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclFunc( self, table.unpack( argList ) )
end
function DeclFuncNode.new( pos, typeInfoList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclFuncNode } )
  if obj.__init then obj:__init( pos, typeInfoList, declInfo ); end
return obj
end
function DeclFuncNode:__init(pos, typeInfoList, declInfo) 
  Node.__init( self, nodeKindDeclFunc, pos, typeInfoList)
  
  -- none
  
  self.declInfo = declInfo
  -- none
  
end
                function DeclFuncNode:get_declInfo()
                return self.declInfo
                end


-- none

function Filter:processDeclMethod( node, ... )
end

-- none

-- none

local nodeKindDeclMethod = regKind( [[DeclMethod]] )
local DeclMethodNode = {}
setmetatable( DeclMethodNode, { __index = Node } )
moduleObj.DeclMethodNode = DeclMethodNode
function DeclMethodNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclMethod( self, table.unpack( argList ) )
end
function DeclMethodNode.new( pos, typeInfoList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMethodNode } )
  if obj.__init then obj:__init( pos, typeInfoList, declInfo ); end
return obj
end
function DeclMethodNode:__init(pos, typeInfoList, declInfo) 
  Node.__init( self, nodeKindDeclMethod, pos, typeInfoList)
  
  -- none
  
  self.declInfo = declInfo
  -- none
  
end
                function DeclMethodNode:get_declInfo()
                return self.declInfo
                end


-- none

function Filter:processDeclConstr( node, ... )
end

-- none

-- none

local nodeKindDeclConstr = regKind( [[DeclConstr]] )
local DeclConstrNode = {}
setmetatable( DeclConstrNode, { __index = Node } )
moduleObj.DeclConstrNode = DeclConstrNode
function DeclConstrNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclConstr( self, table.unpack( argList ) )
end
function DeclConstrNode.new( pos, typeInfoList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclConstrNode } )
  if obj.__init then obj:__init( pos, typeInfoList, declInfo ); end
return obj
end
function DeclConstrNode:__init(pos, typeInfoList, declInfo) 
  Node.__init( self, nodeKindDeclConstr, pos, typeInfoList)
  
  -- none
  
  self.declInfo = declInfo
  -- none
  
end
                function DeclConstrNode:get_declInfo()
                return self.declInfo
                end


-- none

function Filter:processExpCallSuper( node, ... )
end

-- none

-- none

local nodeKindExpCallSuper = regKind( [[ExpCallSuper]] )
local ExpCallSuperNode = {}
setmetatable( ExpCallSuperNode, { __index = Node } )
moduleObj.ExpCallSuperNode = ExpCallSuperNode
function ExpCallSuperNode:processFilter( filter, ... )
  local argList = {...}
  filter:processExpCallSuper( self, table.unpack( argList ) )
end
function ExpCallSuperNode.new( pos, typeInfoList, superType, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallSuperNode } )
  if obj.__init then obj:__init( pos, typeInfoList, superType, expList ); end
return obj
end
function ExpCallSuperNode:__init(pos, typeInfoList, superType, expList) 
  Node.__init( self, nodeKindExpCallSuper, pos, typeInfoList)
  
  -- none
  
  self.superType = superType
  self.expList = expList
  -- none
  
end
                function ExpCallSuperNode:get_superType()
                return self.superType
                end
                function ExpCallSuperNode:get_expList()
                return self.expList
                end


-- none

function Filter:processDeclMember( node, ... )
end

-- none

-- none

local nodeKindDeclMember = regKind( [[DeclMember]] )
local DeclMemberNode = {}
setmetatable( DeclMemberNode, { __index = Node } )
moduleObj.DeclMemberNode = DeclMemberNode
function DeclMemberNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclMember( self, table.unpack( argList ) )
end
function DeclMemberNode.new( pos, typeInfoList, name, refType, staticFlag, accessMode, getterMode, setterMode )
  local obj = {}
  setmetatable( obj, { __index = DeclMemberNode } )
  if obj.__init then obj:__init( pos, typeInfoList, name, refType, staticFlag, accessMode, getterMode, setterMode ); end
return obj
end
function DeclMemberNode:__init(pos, typeInfoList, name, refType, staticFlag, accessMode, getterMode, setterMode) 
  Node.__init( self, nodeKindDeclMember, pos, typeInfoList)
  
  -- none
  
  self.name = name
  self.refType = refType
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.getterMode = getterMode
  self.setterMode = setterMode
  -- none
  
end
                function DeclMemberNode:get_name()
                return self.name
                end
                function DeclMemberNode:get_refType()
                return self.refType
                end
                function DeclMemberNode:get_staticFlag()
                return self.staticFlag
                end
                function DeclMemberNode:get_accessMode()
                return self.accessMode
                end
                function DeclMemberNode:get_getterMode()
                return self.getterMode
                end
                function DeclMemberNode:get_setterMode()
                return self.setterMode
                end


-- none

function Filter:processDeclArg( node, ... )
end

-- none

-- none

local nodeKindDeclArg = regKind( [[DeclArg]] )
local DeclArgNode = {}
setmetatable( DeclArgNode, { __index = Node } )
moduleObj.DeclArgNode = DeclArgNode
function DeclArgNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclArg( self, table.unpack( argList ) )
end
function DeclArgNode.new( pos, typeInfoList, name, argType )
  local obj = {}
  setmetatable( obj, { __index = DeclArgNode } )
  if obj.__init then obj:__init( pos, typeInfoList, name, argType ); end
return obj
end
function DeclArgNode:__init(pos, typeInfoList, name, argType) 
  Node.__init( self, nodeKindDeclArg, pos, typeInfoList)
  
  -- none
  
  self.name = name
  self.argType = argType
  -- none
  
end
                function DeclArgNode:get_name()
                return self.name
                end
                function DeclArgNode:get_argType()
                return self.argType
                end


-- none

function Filter:processDeclArgDDD( node, ... )
end

-- none

-- none

local nodeKindDeclArgDDD = regKind( [[DeclArgDDD]] )
local DeclArgDDDNode = {}
setmetatable( DeclArgDDDNode, { __index = Node } )
moduleObj.DeclArgDDDNode = DeclArgDDDNode
function DeclArgDDDNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclArgDDD( self, table.unpack( argList ) )
end
function DeclArgDDDNode.new( pos, typeInfoList )
  local obj = {}
  setmetatable( obj, { __index = DeclArgDDDNode } )
  if obj.__init then obj:__init( pos, typeInfoList ); end
return obj
end
function DeclArgDDDNode:__init(pos, typeInfoList) 
  Node.__init( self, nodeKindDeclArgDDD, pos, typeInfoList)
  
  -- none
  
  -- none
  
end


-- none

function Filter:processDeclClass( node, ... )
end

-- none

-- none

local nodeKindDeclClass = regKind( [[DeclClass]] )
local DeclClassNode = {}
setmetatable( DeclClassNode, { __index = Node } )
moduleObj.DeclClassNode = DeclClassNode
function DeclClassNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclClass( self, table.unpack( argList ) )
end
function DeclClassNode.new( pos, typeInfoList, accessMode, name, fieldList, memberList, scope, outerMethodSet )
  local obj = {}
  setmetatable( obj, { __index = DeclClassNode } )
  if obj.__init then obj:__init( pos, typeInfoList, accessMode, name, fieldList, memberList, scope, outerMethodSet ); end
return obj
end
function DeclClassNode:__init(pos, typeInfoList, accessMode, name, fieldList, memberList, scope, outerMethodSet) 
  Node.__init( self, nodeKindDeclClass, pos, typeInfoList)
  
  -- none
  
  self.accessMode = accessMode
  self.name = name
  self.fieldList = fieldList
  self.memberList = memberList
  self.scope = scope
  self.outerMethodSet = outerMethodSet
  -- none
  
end
                function DeclClassNode:get_accessMode()
                return self.accessMode
                end
                function DeclClassNode:get_name()
                return self.name
                end
                function DeclClassNode:get_fieldList()
                return self.fieldList
                end
                function DeclClassNode:get_memberList()
                return self.memberList
                end
                function DeclClassNode:get_scope()
                return self.scope
                end
                function DeclClassNode:get_outerMethodSet()
                return self.outerMethodSet
                end


-- none

function Filter:processDeclMacro( node, ... )
end

-- none

-- none

local nodeKindDeclMacro = regKind( [[DeclMacro]] )
local DeclMacroNode = {}
setmetatable( DeclMacroNode, { __index = Node } )
moduleObj.DeclMacroNode = DeclMacroNode
function DeclMacroNode:processFilter( filter, ... )
  local argList = {...}
  filter:processDeclMacro( self, table.unpack( argList ) )
end
function DeclMacroNode.new( pos, typeInfoList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMacroNode } )
  if obj.__init then obj:__init( pos, typeInfoList, declInfo ); end
return obj
end
function DeclMacroNode:__init(pos, typeInfoList, declInfo) 
  Node.__init( self, nodeKindDeclMacro, pos, typeInfoList)
  
  -- none
  
  self.declInfo = declInfo
  -- none
  
end
                function DeclMacroNode:get_declInfo()
                return self.declInfo
                end


-- none

function Filter:processLiteralNil( node, ... )
end

-- none

-- none

local nodeKindLiteralNil = regKind( [[LiteralNil]] )
local LiteralNilNode = {}
setmetatable( LiteralNilNode, { __index = Node } )
moduleObj.LiteralNilNode = LiteralNilNode
function LiteralNilNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralNil( self, table.unpack( argList ) )
end
function LiteralNilNode.new( pos, typeInfoList )
  local obj = {}
  setmetatable( obj, { __index = LiteralNilNode } )
  if obj.__init then obj:__init( pos, typeInfoList ); end
return obj
end
function LiteralNilNode:__init(pos, typeInfoList) 
  Node.__init( self, nodeKindLiteralNil, pos, typeInfoList)
  
  -- none
  
  -- none
  
end


-- none

function Filter:processLiteralChar( node, ... )
end

-- none

-- none

local nodeKindLiteralChar = regKind( [[LiteralChar]] )
local LiteralCharNode = {}
setmetatable( LiteralCharNode, { __index = Node } )
moduleObj.LiteralCharNode = LiteralCharNode
function LiteralCharNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralChar( self, table.unpack( argList ) )
end
function LiteralCharNode.new( pos, typeInfoList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralCharNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token, num ); end
return obj
end
function LiteralCharNode:__init(pos, typeInfoList, token, num) 
  Node.__init( self, nodeKindLiteralChar, pos, typeInfoList)
  
  -- none
  
  self.token = token
  self.num = num
  -- none
  
end
                function LiteralCharNode:get_token()
                return self.token
                end
                function LiteralCharNode:get_num()
                return self.num
                end


-- none

function Filter:processLiteralInt( node, ... )
end

-- none

-- none

local nodeKindLiteralInt = regKind( [[LiteralInt]] )
local LiteralIntNode = {}
setmetatable( LiteralIntNode, { __index = Node } )
moduleObj.LiteralIntNode = LiteralIntNode
function LiteralIntNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralInt( self, table.unpack( argList ) )
end
function LiteralIntNode.new( pos, typeInfoList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralIntNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token, num ); end
return obj
end
function LiteralIntNode:__init(pos, typeInfoList, token, num) 
  Node.__init( self, nodeKindLiteralInt, pos, typeInfoList)
  
  -- none
  
  self.token = token
  self.num = num
  -- none
  
end
                function LiteralIntNode:get_token()
                return self.token
                end
                function LiteralIntNode:get_num()
                return self.num
                end


-- none

function Filter:processLiteralReal( node, ... )
end

-- none

-- none

local nodeKindLiteralReal = regKind( [[LiteralReal]] )
local LiteralRealNode = {}
setmetatable( LiteralRealNode, { __index = Node } )
moduleObj.LiteralRealNode = LiteralRealNode
function LiteralRealNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralReal( self, table.unpack( argList ) )
end
function LiteralRealNode.new( pos, typeInfoList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralRealNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token, num ); end
return obj
end
function LiteralRealNode:__init(pos, typeInfoList, token, num) 
  Node.__init( self, nodeKindLiteralReal, pos, typeInfoList)
  
  -- none
  
  self.token = token
  self.num = num
  -- none
  
end
                function LiteralRealNode:get_token()
                return self.token
                end
                function LiteralRealNode:get_num()
                return self.num
                end


-- none

function Filter:processLiteralArray( node, ... )
end

-- none

-- none

local nodeKindLiteralArray = regKind( [[LiteralArray]] )
local LiteralArrayNode = {}
setmetatable( LiteralArrayNode, { __index = Node } )
moduleObj.LiteralArrayNode = LiteralArrayNode
function LiteralArrayNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralArray( self, table.unpack( argList ) )
end
function LiteralArrayNode.new( pos, typeInfoList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralArrayNode } )
  if obj.__init then obj:__init( pos, typeInfoList, expList ); end
return obj
end
function LiteralArrayNode:__init(pos, typeInfoList, expList) 
  Node.__init( self, nodeKindLiteralArray, pos, typeInfoList)
  
  -- none
  
  self.expList = expList
  -- none
  
end
                function LiteralArrayNode:get_expList()
                return self.expList
                end


-- none

function Filter:processLiteralList( node, ... )
end

-- none

-- none

local nodeKindLiteralList = regKind( [[LiteralList]] )
local LiteralListNode = {}
setmetatable( LiteralListNode, { __index = Node } )
moduleObj.LiteralListNode = LiteralListNode
function LiteralListNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralList( self, table.unpack( argList ) )
end
function LiteralListNode.new( pos, typeInfoList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralListNode } )
  if obj.__init then obj:__init( pos, typeInfoList, expList ); end
return obj
end
function LiteralListNode:__init(pos, typeInfoList, expList) 
  Node.__init( self, nodeKindLiteralList, pos, typeInfoList)
  
  -- none
  
  self.expList = expList
  -- none
  
end
                function LiteralListNode:get_expList()
                return self.expList
                end


local PairItem = {}
moduleObj.PairItem = PairItem
            function PairItem.new( key, val )
            local obj = {}
            setmetatable( obj, { __index = PairItem } )
            if obj.__init then
            obj:__init( key, val )
            end       
            return obj
            end         
            function PairItem:__init( key, val )
            
self.key = key
  self.val = val
  end
                function PairItem:get_key()
                return self.key
                end
                function PairItem:get_val()
                return self.val
                end

-- none

function Filter:processLiteralMap( node, ... )
end

-- none

-- none

local nodeKindLiteralMap = regKind( [[LiteralMap]] )
local LiteralMapNode = {}
setmetatable( LiteralMapNode, { __index = Node } )
moduleObj.LiteralMapNode = LiteralMapNode
function LiteralMapNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralMap( self, table.unpack( argList ) )
end
function LiteralMapNode.new( pos, typeInfoList, map, pairList )
  local obj = {}
  setmetatable( obj, { __index = LiteralMapNode } )
  if obj.__init then obj:__init( pos, typeInfoList, map, pairList ); end
return obj
end
function LiteralMapNode:__init(pos, typeInfoList, map, pairList) 
  Node.__init( self, nodeKindLiteralMap, pos, typeInfoList)
  
  -- none
  
  self.map = map
  self.pairList = pairList
  -- none
  
end
                function LiteralMapNode:get_map()
                return self.map
                end
                function LiteralMapNode:get_pairList()
                return self.pairList
                end


-- none

function Filter:processLiteralString( node, ... )
end

-- none

-- none

local nodeKindLiteralString = regKind( [[LiteralString]] )
local LiteralStringNode = {}
setmetatable( LiteralStringNode, { __index = Node } )
moduleObj.LiteralStringNode = LiteralStringNode
function LiteralStringNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralString( self, table.unpack( argList ) )
end
function LiteralStringNode.new( pos, typeInfoList, token, argList )
  local obj = {}
  setmetatable( obj, { __index = LiteralStringNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token, argList ); end
return obj
end
function LiteralStringNode:__init(pos, typeInfoList, token, argList) 
  Node.__init( self, nodeKindLiteralString, pos, typeInfoList)
  
  -- none
  
  self.token = token
  self.argList = argList
  -- none
  
end
                function LiteralStringNode:get_token()
                return self.token
                end
                function LiteralStringNode:get_argList()
                return self.argList
                end


-- none

function Filter:processLiteralBool( node, ... )
end

-- none

-- none

local nodeKindLiteralBool = regKind( [[LiteralBool]] )
local LiteralBoolNode = {}
setmetatable( LiteralBoolNode, { __index = Node } )
moduleObj.LiteralBoolNode = LiteralBoolNode
function LiteralBoolNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralBool( self, table.unpack( argList ) )
end
function LiteralBoolNode.new( pos, typeInfoList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralBoolNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token ); end
return obj
end
function LiteralBoolNode:__init(pos, typeInfoList, token) 
  Node.__init( self, nodeKindLiteralBool, pos, typeInfoList)
  
  -- none
  
  self.token = token
  -- none
  
end
                function LiteralBoolNode:get_token()
                return self.token
                end


-- none

function Filter:processLiteralSymbol( node, ... )
end

-- none

-- none

local nodeKindLiteralSymbol = regKind( [[LiteralSymbol]] )
local LiteralSymbolNode = {}
setmetatable( LiteralSymbolNode, { __index = Node } )
moduleObj.LiteralSymbolNode = LiteralSymbolNode
function LiteralSymbolNode:processFilter( filter, ... )
  local argList = {...}
  filter:processLiteralSymbol( self, table.unpack( argList ) )
end
function LiteralSymbolNode.new( pos, typeInfoList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralSymbolNode } )
  if obj.__init then obj:__init( pos, typeInfoList, token ); end
return obj
end
function LiteralSymbolNode:__init(pos, typeInfoList, token) 
  Node.__init( self, nodeKindLiteralSymbol, pos, typeInfoList)
  
  -- none
  
  self.token = token
  -- none
  
end
                function LiteralSymbolNode:get_token()
                return self.token
                end


function LiteralNilNode:getLiteral(  )
  return {nil}, {typeInfoNil}
end

function LiteralCharNode:getLiteral(  )
  return {self.num}, {typeInfoChar}
end

function LiteralIntNode:getLiteral(  )
  return {self.num}, {typeInfoInt}
end

function LiteralRealNode:getLiteral(  )
  return {self.num}, {typeInfoReal}
end

function LiteralArrayNode:getLiteral(  )
  local array = {}
  for __index, val in pairs( self.expList:get_expList(  ) ) do
    local txt = val:getLiteral(  )[1]
    table.insert( array, txt )
  end
  return {array}, {self:get_expType(  )}
end

function LiteralListNode:getLiteral(  )
  local list = {}
  for __index, val in pairs( self.expList:get_expList(  ) ) do
    local item = val:getLiteral(  )[1]
    table.insert( list, item )
  end
  return {list}, {self:get_expType(  )}
end

function LiteralMapNode:getLiteral(  )
  local map = {}
  for key, val in pairs( self.map ) do
    map[key:getLiteral(  )[1]] = val:getLiteral(  )[1]
  end
  return {map}, {self:get_expType(  )}
end

function LiteralStringNode:getLiteral(  )
  local txt = self.token.txt
  if string.find( txt, '^```' ) then
    txt = txt:sub( 4, -4 )
  else 
    txt = txt:sub( 2, -2 )
  end
  local argList = self:get_argList(  )
  if argList and #argList > 0 then
    local argTbl = {}
    local argValList = argList.getLiteral(  )
    for __index, arg in pairs( argValList ) do
      table.insert( argTbl, arg )
    end
    return string.format( txt, table.unpack( argTbl ) ), typeInfoString
  end
  return {txt}, {typeInfoString}
end

function LiteralBoolNode:getLiteral(  )
  return {self.token.txt == "true"}, {typeInfoBool}
end

function LiteralSymbolNode:getLiteral(  )
  return {{self.token.txt}}, {typeInfoSymbol}
end

function RefFieldNode:getLiteral(  )
  local prefix = self.prefix:getLiteral(  )[1]
  table.insert( prefix, "." )
  table.insert( prefix, self.field.txt )
  return {prefix}, {self:get_expType(  )}
end

function ExpMacroStatNode:getLiteral(  )
  local txt = ""
  for __index, token in pairs( self.expStrList ) do
    txt = string.format( "%s %s", txt, token:getLiteral(  )[1])
  end
  return {txt}, {self:get_expType(  )}
end

local MacroEval = {}
moduleObj.MacroEval = MacroEval
-- none
            function MacroEval.new(  )
            local obj = {}
            setmetatable( obj, { __index = MacroEval } )
            if obj.__init then
            obj:__init(  )
            end       
            return obj
            end         
            function MacroEval:__init(  )
            
end

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
function TransUnit:registBuiltInScope(  )
  local builtInInfo = {[""] = {["type"] = {["ret"] = {"str"}}, ["error"] = {["ret"] = {}}, ["print"] = {["ret"] = {}}, ["tonumber"] = {["ret"] = {"int"}}, ["load"] = {["ret"] = {"stem", "str"}}, ["require"] = {["ret"] = {"stem"}}, ["_fcall"] = {["ret"] = {"stem"}}}, ["io"] = {["open"] = {["ret"] = {"stem"}}}, ["os"] = {["clock"] = {["ret"] = {"int"}}}, ["string"] = {["find"] = {["ret"] = {"int", "int"}}, ["byte"] = {["ret"] = {"int"}}, ["format"] = {["ret"] = {"str"}}, ["rep"] = {["ret"] = {"str"}}, ["gmatch"] = {["ret"] = {"stem"}}, ["gsub"] = {["ret"] = {"str"}}, ["sub"] = {["ret"] = {"str"}}}, ["str"] = {["find"] = {["methodFlag"] = {}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["ret"] = {"stem"}}, ["gsub"] = {["methodFlag"] = {}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["ret"] = {"str"}}}, ["table"] = {["insert"] = {["ret"] = {""}}, ["remove"] = {["ret"] = {""}}, ["unpack"] = {["ret"] = {"stem"}}}, ["debug"] = {["getinfo"] = {["ret"] = {"stem"}}}, ["_luneScript"] = {["loadModule"] = {["ret"] = {"stem"}}}}
  do
    local __sorted = {}
    local __map = builtInTypeMap
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, name in ipairs( __sorted ) do
      typeInfo = __map[ name ]
      do
        if typeInfo.kind == TypeInfoKindClass then
          local scope = self:pushScope( true )
          scope:get_parent(  ):addClass( name, typeInfo, scope )
          self:popScope(  )
        else 
          self.scope:add( name, typeInfo )
        end
      end
    end
  end
  
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
        local parentInfo = typeInfoRoot
        if name ~= "" then
          local classTypeInfo = self:pushClass( nil, true, name, "pri" )
          parentInfo = classTypeInfo
          builtInTypeIdSet[classTypeInfo:get_typeId(  )] = true
        end
        if not parentInfo then
          error( "parentInfo is nil" )
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
              local methodFlag = info["methodFlag"]
              local typeInfo = TypeInfo.createFunc( methodFlag and TypeInfoKindMethod or TypeInfoKindFunc, parentInfo, false, true, not methodFlag, "pub", funcName, nil, retTypeList )
              builtInTypeIdSet[typeInfo:get_typeId(  )] = true
              self.scope:add( funcName, typeInfo )
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

function TransUnit:error( mess )
  local pos = {["lineNo"] = 0, ["column"] = 0}
  local txt = ""
  if self.currentToken then
    pos = self.currentToken.pos
    txt = self.currentToken.txt
  end
  error( string.format( "%s:%d:%d:(%s) %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:createNoneNode( pos )
  return NoneNode.new(pos, {typeInfoNone})
end

function TransUnit:pushbackToken( token )
  table.insert( self.pushbackList, token )
  self.currentToken = nil
end

local function expandVal( tokenList, val, pos )
  do
    local _switchExp = type( val )
    if _switchExp == "number" then
      local num = string.format( "%g", val)
      local kind = Parser.kind.Int
      if string.find( num, ".", 1, true ) then
        kind = Parser.kind.Real
      end
      table.insert( tokenList, Parser.Token.new(kind, num, pos) )
    elseif _switchExp == "string" then
      table.insert( tokenList, Parser.Token.new(Parser.kind.Str, string.format( '[[%s]]', val), pos) )
    elseif _switchExp == "table" then
      table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "{", val), pos) )
      for key, item in pairs( val ) do
        expandVal( tokenList, item, pos )
        table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( ",", val), pos) )
      end
      table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "}", val), pos) )
    end
  end
  
end

function TransUnit:newPushback( tokenList )
  for index = #tokenList, 1, -1 do
    table.insert( self.pushbackList, tokenList[index] )
  end
  self.currentToken = nil
end

function TransUnit:getTokenNoErr(  )
  if #self.pushbackList > 0 then
    self.currentToken = self.pushbackList[#self.pushbackList]
    table.remove( self.pushbackList )
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
    if self.macroMode == "expand" and token.txt == ',,' then
      token = self:getTokenNoErr(  )
      local macroVal = self.symbol2ValueMapForMacro[token.txt]
      if macroVal then
        if macroVal.typeInfo == typeInfoSymbol then
          local txtList = macroVal.val
          for index = #txtList, 1, -1 do
            token = Parser.Token.new(token.kind, txtList[index], token.pos)
            self:pushbackToken( token )
          end
        elseif macroVal.typeInfo == typeInfoStat then
          self:pushbackStr( string.format( "macroVal %s", token.txt), macroVal.val )
        elseif macroVal.typeInfo:get_kind(  ) == TypeInfoKindArray or macroVal.typeInfo:get_kind(  ) == TypeInfoKindList then
          local strList = macroVal.val
          if strList then
            for index = #strList, 1, -1 do
              self:pushbackStr( string.format( "macroVal %s[%d]", token.txt, index), strList[index] )
            end
          else 
            self:error( string.format( "macro val is nil %s", token.txt) )
          end
        else 
          local tokenList = {}
          expandVal( tokenList, macroVal.val, token.pos )
          self:newPushback( tokenList )
        end
        token = self:getTokenNoErr(  )
      else 
        self:error( string.format( "unknown macro val %s", token.txt) )
      end
    end
    token:set_commentList( commentList )
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
  table.insert( self.pushbackList, self.currentToken )
  self.currentToken = nil
end

function TransUnit:pushbackStr( name, statement )
  local memStream = Parser.TxtStream.new(statement)
  local parer = Parser.StreamParser.new(memStream, name, false)
  local list = {}
  while true do
    local token = parer:getToken(  )
    if not token then
      break
    end
    table.insert( list, token )
  end
  for index = #list, 1, -1 do
    self:pushbackToken( list[index] )
  end
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
    self:pushScope( false )
  end
  local stmtList = {}
  self:analyzeStatementList( stmtList, "}" )
  self:checkNextToken( "}" )
  if not scope then
    self:popScope(  )
  end
  local node = BlockNode.new(token.pos, {typeInfoNone}, blockKind, stmtList)
  return node
end

local _TypeInfo = {}
            function _TypeInfo.new( baseId, itemTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children )
            local obj = {}
            setmetatable( obj, { __index = _TypeInfo } )
            if obj.__init then
            obj:__init( baseId, itemTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children )
            end       
            return obj
            end         
            function _TypeInfo:__init( baseId, itemTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children )
            
self.baseId = baseId
  self.itemTypeId = itemTypeId
  self.retTypeId = retTypeId
  self.parentId = parentId
  self.typeId = typeId
  self.txt = txt
  self.kind = kind
  self.staticFlag = staticFlag
  self.nilable = nilable
  self.orgTypeId = orgTypeId
  self.children = children
  end

local _ModuleInfo = {}
            function _ModuleInfo.new( _className2InfoMap, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
            local obj = {}
            setmetatable( obj, { __index = _ModuleInfo } )
            if obj.__init then
            obj:__init( _className2InfoMap, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
            end       
            return obj
            end         
            function _ModuleInfo:__init( _className2InfoMap, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
            
self._className2InfoMap = _className2InfoMap
  self._typeId2ClassInfoMap = _typeId2ClassInfoMap
  self._typeInfoList = _typeInfoList
  self._varName2InfoMap = _varName2InfoMap
  self._funcName2InfoMap = _funcName2InfoMap
  end

function TransUnit:analyzeImport( token )
  if self.moduleScope ~= self.scope then
    self:error( "'import' must call at top scope." )
  end
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
  local typeId2TypeInfo = {}
  typeId2TypeInfo[rootTypeId] = typeInfoRoot
  local typeInfo2Scope = {}
  local moduleTypeInfo = nil
  for __index, moduleName in pairs( nameList ) do
    moduleTypeInfo = self:pushClass( nil, true, moduleName, "pub" )
    typeInfo2Scope[moduleTypeInfo] = self.scope
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  local moduleInfo = _luneScript.loadModule( modulePath )
  self.moduleName2Info[modulePath] = moduleInfo
  for __index, typeInfo in pairs( builtInTypeMap ) do
    typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
  end
  local typeId2Scope = {}
  typeId2Scope[rootTypeId] = self.scope
  local newId2OldIdMap = {}
  local function registTypeInfo( atomInfo )
    local newTypeInfo = nil
    if not builtInTypeIdSet[atomInfo.typeId] then
      if atomInfo.nilable then
        local orgTypeInfo = typeId2TypeInfo[atomInfo.orgTypeId]
        newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
        typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
      else 
        local itemTypeInfo = {}
        for __index, typeId in pairs( atomInfo.itemTypeId ) do
          table.insert( itemTypeInfo, typeId2TypeInfo[typeId] )
        end
        local retTypeInfo = {}
        for __index, typeId in pairs( atomInfo.retTypeId ) do
          table.insert( retTypeInfo, typeId2TypeInfo[typeId] )
        end
        local parentInfo = rootTypeInfo
        if atomInfo.parentId ~= rootTypeId then
          parentInfo = typeId2TypeInfo[atomInfo.parentId]
          if not parentInfo then
            error( string.format( "not found parentInfo %s %s", atomInfo.parentId, atomInfo.txt) )
          end
        end
        local baseInfo = typeId2TypeInfo[atomInfo.baseId]
        local parentScope = typeId2Scope[atomInfo.parentId]
        if not parentScope then
          error( string.format( "not found parentScope %s %s", atomInfo.parentId, atomInfo.txt) )
        end
        if atomInfo.txt ~= "" then
          newTypeInfo = parentScope:getTypeInfoChild( atomInfo.txt )
        end
        if newTypeInfo and atomInfo.kind == TypeInfoKindClass then
          typeId2Scope[atomInfo.typeId] = typeInfo2Scope[newTypeInfo]
          if not typeInfo2Scope[newTypeInfo] then
            error( string.format( "not found scope %s %s %s %s %s", parentScope, atomInfo.parentId, atomInfo.typeId, atomInfo.txt, newTypeInfo:getTxt(  )) )
          end
          typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
        else 
          if atomInfo.kind == TypeInfoKindClass then
            local baseScope = typeId2Scope[atomInfo.baseId]
            local scope = Scope.new(parentScope, true, baseScope and {baseScope} or nil)
            newTypeInfo = TypeInfo.createClass( baseInfo, parentInfo, true, "pub", atomInfo.txt )
            typeId2Scope[atomInfo.typeId] = scope
            typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
            parentScope:addClass( atomInfo.txt, newTypeInfo, scope )
          else 
            newTypeInfo = TypeInfo.create( baseInfo, parentInfo, atomInfo.staticFlag, atomInfo.kind, atomInfo.txt, itemTypeInfo, retTypeInfo )
            typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
            if atomInfo.kind == TypeInfoKindFunc or atomInfo.kind == TypeInfoKindMethod then
              typeId2Scope[atomInfo.parentId]:add( atomInfo.txt, newTypeInfo )
              local scope = Scope.new(parentScope, false)
              typeId2Scope[atomInfo.typeId] = scope
            end
          end
        end
      end
    else 
      newTypeInfo = builtInTypeMap[atomInfo.txt]
      typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
    end
    return newTypeInfo
  end
  
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    registTypeInfo( atomInfo )
  end
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    if #atomInfo.children > 0 then
      local scope = typeId2Scope[atomInfo.typeId]
      for __index, childId in pairs( atomInfo.children ) do
        local typeInfo = typeId2TypeInfo[childId]
        if typeInfo then
          scope:add( typeInfo:getTxt(  ), typeInfo )
        end
      end
    end
  end
  for typeId, typeInfo in pairs( typeId2TypeInfo ) do
    newId2OldIdMap[typeInfo:get_typeId(  )] = typeId
  end
  local function registMember( classTypeId )
    local classTypeInfo = typeId2TypeInfo[classTypeId]
    self:pushClass( nil, true, classTypeInfo:getTxt(  ), "pub" )
    local classInfo = moduleInfo._typeId2ClassInfoMap[classTypeId]
    if classInfo then
      for fieldName, fieldInfo in pairs( classInfo ) do
        local fieldTypeInfo = nil
        local typeId = fieldInfo["typeId"]
        fieldTypeInfo = typeId2TypeInfo[typeId]
        self.scope:add( fieldName, fieldTypeInfo )
      end
    end
    for __index, child in pairs( classTypeInfo:get_children(  ) ) do
      if child:get_kind(  ) == TypeInfoKindClass then
        local oldId = newId2OldIdMap[child:get_typeId(  )]
        if oldId then
          registMember( oldId )
        end
      end
    end
    self:popClass(  )
  end
  
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    if atomInfo.parentId == rootTypeId and atomInfo.kind == TypeInfoKindClass then
      registMember( atomInfo.typeId )
    end
  end
  for __index, moduleName in pairs( nameList ) do
    self:pushClass( nil, true, moduleName, "pub" )
  end
  for varName, varInfo in pairs( moduleInfo._varName2InfoMap ) do
    self.scope:add( varName, typeId2TypeInfo[varInfo["typeId"]] )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  self.scope:add( moduleToken.txt, moduleTypeInfo )
  self:checkToken( nextToken, ";" )
  return ImportNode.new(token.pos, {typeInfoNone}, modulePath)
end

function TransUnit:analyzeIf( token )
  local list = {}
  table.insert( list, IfStmtInfo.new("if", self:analyzeExp(  ), self:analyzeBlock( "if" )) )
  local nextToken = self:getToken(  )
  if nextToken.txt == "elseif" then
    while nextToken.txt == "elseif" do
      table.insert( list, IfStmtInfo.new("elseif", self:analyzeExp(  ), self:analyzeBlock( "elseif" )) )
      nextToken = self:getToken(  )
    end
  end
  if nextToken.txt == "else" then
    table.insert( list, IfStmtInfo.new("else", nil, self:analyzeBlock( "else" )) )
  else 
    self:pushback(  )
  end
  return IfNode.new(token.pos, {typeInfoNone}, list)
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
    table.insert( caseList, CaseInfo.new(condexpList, condBock) )
    nextToken = self:getToken(  )
  end
  local defaultBlock = nil
  if nextToken.txt == "default" then
    defaultBlock = self:analyzeBlock( "default" )
  else 
    self:pushback(  )
  end
  self:checkNextToken( "}" )
  return SwitchNode.new(firstToken.pos, {typeInfoNone}, exp, caseList, defaultBlock)
end

function TransUnit:analyzeWhile( token )
  return WhileNode.new(token.pos, {typeInfoNone}, self:analyzeExp(  ), self:analyzeBlock( "while" ))
end

function TransUnit:analyzeRepeat( token )
  local scope = self:pushScope( false )
  local node = RepeatNode.new(token.pos, {typeInfoNone}, self:analyzeBlock( "repeat", scope ), self:analyzeExp(  ))
  self:popScope(  )
  self:checkNextToken( ";" )
  return node
end

function TransUnit:analyzeFor( token )
  local scope = self:pushScope( false )
  local val = self:getToken(  )
  if val.kind ~= Parser.kind.Symb then
    self:error( "not symbol" )
  end
  self:checkNextToken( "=" )
  local exp1 = self:analyzeExp(  )
  self.scope:add( val.txt, exp1:get_expType() )
  self:checkNextToken( "," )
  local exp2 = self:analyzeExp(  )
  local token = self:getToken(  )
  local exp3 = nil
  if token.txt == "," then
    exp3 = self:analyzeExp(  )
  else 
    self:pushback(  )
  end
  local node = ForNode.new(token.pos, {typeInfoNone}, self:analyzeBlock( "for", scope ), val, exp1, exp2, exp3)
  self:popScope(  )
  return node
end

function TransUnit:analyzeApply( token )
  local scope = self:pushScope(  )
  local varList = {}
  local nextToken = nil
  repeat 
    local var = self:getSymbolToken(  )
    if var.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    table.insert( varList, var )
    nextToken = self:getToken(  )
    scope:add( var.txt, typeInfoStem )
  until nextToken.txt ~= ","
  self:checkToken( nextToken, "of" )
  local exp = self:analyzeExp(  )
  if exp.kind ~= nodeKindExpCall then
    self:error( "not call" )
  end
  local block = self:analyzeBlock( "apply", scope )
  self:popScope(  )
  return ApplyNode.new(token.pos, {typeInfoNone}, varList, exp, block)
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
  if not exp:get_expType() then
    self:error( string.format( "unknown type of exp -- %d:%d", token.pos.lineNo, token.pos.column) )
  else 
    local itemTypeInfoList = exp:get_expType():get_itemTypeInfoList(  )
    if exp:get_expType():get_kind(  ) == TypeInfoKindMap then
      self.scope:add( valSymbol.txt, itemTypeInfoList[2] )
      if keySymbol then
        self.scope:add( keySymbol.txt, itemTypeInfoList[1] )
      end
    elseif exp:get_expType():get_kind(  ) == TypeInfoKindList or exp:get_expType():get_kind(  ) == TypeInfoKindArray then
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
  if sortFlag then
    return ForsortNode.new(token.pos, {typeInfoNone}, valSymbol, keySymbol, exp, block, sortFlag)
  else 
    return ForeachNode.new(token.pos, {typeInfoNone}, valSymbol, keySymbol, exp, block, sortFlag)
  end
end

function TransUnit:analyzeRefType( accessMode )
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
  token = self:checkSymbol( token )
  if token then
    name = self:analyzeExpSymbol( firstToken, token, "symbol", token, true )
    typeInfo = name:get_expType()
  else 
    self:pushback(  )
  end
  token = self:getToken(  )
  if token.txt == "!" then
    typeInfo = typeInfo:get_nilableTypeInfo(  )
    token = self:getToken(  )
  end
  local arrayMode = "no"
  while true do
    if token.txt == '[' or token.txt == '[@' then
      if token.txt == '[' then
        arrayMode = "list"
        typeInfo = TypeInfo.createList( accessMode, self:getCurrentClass(  ), {typeInfo} )
      else 
        arrayMode = "array"
        typeInfo = TypeInfo.createArray( accessMode, self:getCurrentClass(  ), {typeInfo} )
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
        local typeExp = self:analyzeRefType( accessMode )
        table.insert( genericList, typeExp:get_expType() )
        nextToken = self:getToken(  )
      until nextToken.txt ~= ","
      self:checkToken( nextToken, '>' )
      if typeInfo.kind == TypeInfoKindMap then
        typeInfo = TypeInfo.createMap( accessMode, self:getCurrentClass(  ), genericList[1] or typeInfoStem, genericList[2] or typeInfoStem )
      else 
        self:error( string.format( "not support generic: %s", typeInfo:getTxt(  ) ) )
      end
    else 
      self:pushback(  )
      break
    end
    token = self:getToken(  )
  end
  return RefTypeNode.new(firstToken.pos, {typeInfo}, name, refFlag, mutFlag, arrayMode)
end

function TransUnit:analyzeDeclArgList( accessMode, argList )
  local token = Parser.noneToken
  repeat 
    local argName = self:getToken(  )
    if argName.txt == ")" then
      token = argName
      break
    elseif argName.txt == "..." then
      table.insert( argList, DeclArgDDDNode.new(argName.pos, {typeInfoNone}) )
    else 
      argName = self:checkSymbol( argName )
      self:checkNextToken( ":" )
      local refType = self:analyzeRefType( accessMode )
      local arg = DeclArgNode.new(argName.pos, refType:get_expTypeList(), argName, refType)
      self.scope:add( argName.txt, refType:get_expType() )
      table.insert( argList, arg )
    end
    token = self:getToken(  )
  until token.txt ~= ","
  self:checkToken( token, ")" )
  return token
end

local ASTInfo = {}
moduleObj.ASTInfo = ASTInfo
            function ASTInfo.new( node, moduleTypeInfo )
            local obj = {}
            setmetatable( obj, { __index = ASTInfo } )
            if obj.__init then
            obj:__init( node, moduleTypeInfo )
            end       
            return obj
            end         
            function ASTInfo:__init( node, moduleTypeInfo )
            
self.node = node
  self.moduleTypeInfo = moduleTypeInfo
  end
                function ASTInfo:get_node()
                return self.node
                end
                function ASTInfo:get_moduleTypeInfo()
                return self.moduleTypeInfo
                end

function TransUnit:createAST( parser, macroFlag, module )
  self:pushNamespace( "", typeInfoRoot, self.scope )
  self.rootScope = self.scope
  self:registBuiltInScope(  )
  local moduleTypeInfo = rootTypeInfo
  if module then
    for txt in string.gmatch( module, '[^%.]+' ) do
      moduleTypeInfo = self:pushClass( nil, false, txt, "pub" )
    end
  end
  self.moduleScope = self.scope
  self.parser = parser
  self.moduleName2Info = {}
  local ast = nil
  if macroFlag then
    ast = self:analyzeBlock( "macro" )
  else 
    local rootInfo = {}
    rootInfo.children = {}
    ast = RootNode.new(Parser.Position.new(0, 0), {typeInfoNone}, rootInfo.children, self.typeId2ClassMap)
    self:analyzeStatementList( rootInfo.children )
    local token = self:getTokenNoErr(  )
    if token then
      error( string.format( "unknown:%d:%d:(%s) %s", token.pos.lineNo, token.pos.column, Parser.getKindTxt( token.kind ), token.txt) )
    end
  end
  if module then
    for txt in string.gmatch( module, '[^%.]+' ) do
      self:popClass(  )
    end
  end
  if #self.errMessList > 0 then
    for __index, mess in pairs( self.errMessList ) do
      Util.errorLog( "error:" .. mess )
    end
    error( "has error" )
  end
  self:popNamespace(  )
  return ASTInfo.new(ast, moduleTypeInfo)
end

function TransUnit:analyzeDeclMacro( accessMode, firstToken )
  local nameToken = self:getToken(  )
  self:checkNextToken( "(" )
  local scope = self:pushScope(  )
  local argList = {}
  local nextToken = self:analyzeDeclArgList( accessMode, argList )
  self:checkToken( nextToken, ")" )
  self:checkNextToken( "{" )
  local nextToken = self:getToken(  )
  local ast = nil
  if nextToken.txt == "{" then
    local parser = Parser.WrapParser.new(self.parser, string.format( "decl macro %s", nameToken.txt))
    for symbol, typeInfo in pairs( scope:get_symbol2TypeInfoMap(  ) ) do
      scope:add( symbol, typeInfo )
    end
    self.macroScope = scope
    local bakParser = self.parser
    self.parser = parser
    local stmtList = {}
    self:analyzeStatementList( stmtList, "}" )
    self:checkNextToken( "}" )
    self.parser = bakParser
    self.macroScope = nil
    ast = BlockNode.new(firstToken.pos, {typeInfoNone}, "macro", stmtList)
  else 
    self:pushback(  )
  end
  self:popScope(  )
  local tokenList = {}
  local braceCount = 0
  while true do
    nextToken = self:getToken(  )
    if nextToken.txt == "{" then
      braceCount = braceCount + 1
    elseif nextToken.txt == "}" then
      if braceCount == 0 then
        break
      end
      braceCount = braceCount - 1
    end
    table.insert( tokenList, nextToken )
  end
  local typeInfo = TypeInfo.createFunc( TypeInfoKindMacro, self:getCurrentNamespaceTypeInfo(  ), false, false, false, accessMode, nameToken.txt, nil, nil )
  self.scope:add( nameToken.txt, typeInfo )
  local declMacroInfo = DeclMacroInfo.new(nameToken, argList, ast, tokenList)
  local node = DeclMacroNode.new(firstToken.pos, {typeInfo}, declMacroInfo)
  local macroObj = self.macroEval:eval( node )
  self.typeId2MacroInfo[typeInfo:get_typeId(  )] = MacroInfo.new(macroObj, declMacroInfo, self.symbol2ValueMapForMacro)
  self.symbol2ValueMapForMacro = {}
  return node
end

function TransUnit:analyzeDeclProto( accessMode, firstToken )
  local nextToken = self:getToken(  )
  if nextToken.txt == "class" then
    local name = self:getToken(  )
    nextToken = self:getToken(  )
    local baseRef = nil
    if nextToken.txt == "extend" then
      baseRef = self:analyzeRefType( accessMode )
      nextToken = self:getToken(  )
    end
    self:checkToken( nextToken, ";" )
    self:pushClass( baseRef and baseRef:get_expType(  ) or nil, false, name.txt, accessMode )
    self:popClass(  )
  else 
    self:error( "illegal proto" )
  end
  return self:createNoneNode( firstToken.pos )
end

function TransUnit:analyzeDecl( accessMode, staticFlag, firstToken, token )
  local staticFlag = false
  if not staticFlag then
    if token.txt == "static" then
      staticFlag = true
      token = self:getToken(  )
    end
  end
  local overrideFlag = false
  if token.txt == "override" then
    overrideFlag = true
    token = self:getToken(  )
  end
  if token.txt == "let" then
    return self:analyzeDeclVar( accessMode, staticFlag, firstToken )
  elseif token.txt == "fn" then
    return self:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, nil, firstToken, nil )
  elseif token.txt == "class" then
    return self:analyzeDeclClass( accessMode, firstToken )
  elseif token.txt == "proto" then
    return self:analyzeDeclProto( accessMode, firstToken )
  elseif token.txt == "macro" then
    return self:analyzeDeclMacro( accessMode, firstToken )
  end
  return nil
end

function TransUnit:analyzeDeclMember( accessMode, staticFlag, firstToken )
  local varName = self:getSymbolToken(  )
  local token = self:getToken(  )
  local refType = self:analyzeRefType( accessMode )
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
  self.scope:add( varName.txt, refType:get_expType() )
  return DeclMemberNode.new(firstToken.pos, refType:get_expTypeList(), varName, refType, staticFlag, accessMode, getterMode, setterMode)
end

function TransUnit:analyzeDeclMethod( overrideFlag, accessMode, staticFlag, className, firstToken, name )
  local node = self:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, className, name, name )
  return node
end

function TransUnit:analyzeDeclClass( classAccessMode, firstToken )
  local name = self:getSymbolToken(  )
  local nextToken = self:getToken(  )
  local baseRef = nil
  if nextToken.txt == "extend" then
    baseRef = self:analyzeRefType( classAccessMode )
    nextToken = self:getToken(  )
  end
  self:checkToken( nextToken, "{" )
  local classTypeInfo = self:pushClass( baseRef and baseRef:get_expType(  ) or nil, false, name.txt, classAccessMode )
  local fieldList = {}
  local memberList = {}
  local methodName2Node = {}
  local node = DeclClassNode.new(firstToken.pos, {classTypeInfo}, classAccessMode, name, fieldList, memberList, self.scope, {})
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
    local overrideFlag = false
    if token.txt == "override" then
      overrideFlag = true
      token = self:getToken(  )
    end
    if token.txt == "let" then
      local memberNode = self:analyzeDeclMember( accessMode, staticFlag, token )
      table.insert( fieldList, memberNode )
      table.insert( memberList, memberNode )
    elseif token.txt == "fn" then
      local nameToken = self:getToken(  )
      local methodNode = self:analyzeDeclMethod( overrideFlag, accessMode, staticFlag, name, token, nameToken )
      table.insert( fieldList, methodNode )
    elseif token.txt == ";" then
    else 
      self:error( "illegal field" )
    end
  end
  local parentInfo = classTypeInfo
  for __index, memberNode in pairs( memberList ) do
    local memberType = memberNode:get_expType()
    if memberNode:get_expType().accessMode ~= "pub" then
      memberType = TypeInfo.cloneToPublic( memberType )
    end
    local memberName = memberNode:get_name()
    local getterName = "get_" .. memberName.txt
    local accessMode = memberNode:get_getterMode()
    if accessMode ~= "none" and not self.scope:getTypeInfo( getterName ) then
      local retTypeInfo = TypeInfo.createFunc( TypeInfoKindMethod, parentInfo, true, false, false, "pub", getterName, nil, {memberType} )
      self.scope:add( getterName, retTypeInfo )
    end
    local setterName = "set_" .. memberName.txt
    local accessMode = memberNode:get_setterMode()
    if memberNode:get_setterMode() ~= "none" and not self.scope:getTypeInfo( setterName ) then
      self.scope:add( setterName, TypeInfo.createFunc( TypeInfoKindMethod, parentInfo, true, false, false, "pub", setterName, nil, {memberType} ) )
    end
  end
  self:popClass(  )
  return node
end

function TransUnit:addMethod( className, methodNode, name )
  local classTypeInfo = self.scope:getTypeInfo( className )
  local classNodeInfo = self.typeInfo2ClassNode[classTypeInfo]
  classNodeInfo.outerMethodSet[name] = true
  table.insert( classNodeInfo.fieldList, methodNode )
end

function TransUnit:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, classNameToken, firstToken, name )
  local argList = {}
  local token = self:getToken(  )
  if not name then
    if token.txt ~= "(" then
      name = self:checkSymbol( token )
      token = self:getToken(  )
    end
  else 
    name = self:checkSymbol( name )
  end
  local needPopFlag = false
  if token.txt == "." then
    needPopFlag = true
    classNameToken = name
    self:pushClass( nil, false, name.txt, "pub" )
    name = self:getSymbolToken(  )
    token = self:getToken(  )
  end
  local kind = nodeKindDeclConstr
  local typeKind = TypeInfoKindFunc
  if classNameToken then
    if not staticFlag then
      typeKind = TypeInfoKindMethod
    end
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
  if overrideFlag then
    local overrideType = self.scope:getTypeInfoMethod( name.txt )
    if not overrideType then
      self:error( "not found override -- " .. name.txt )
    end
    if overrideType:get_accessMode(  ) ~= accessMode then
      self:error( string.format( "missmatch override accessMode -- %s,%s,%s", name.txt, overrideType:get_accessMode(  ), accessMode) )
    end
    if overrideType:get_staticFlag(  ) ~= staticFlag then
      self:error( "missmatch override staticFlag -- " .. name.txt )
    end
    if overrideType:get_kind(  ) ~= TypeInfoKindMethod then
      self:error( string.format( "missmatch override kind -- %s, %d", name.txt, overrideType:get_kind(  )) )
    end
  elseif name and self.scope:getTypeInfoMethod( name.txt ) then
    self:error( "missmatch override --" .. name.txt )
  end
  self:checkToken( "(" )
  local scope = self:pushScope(  )
  token = self:analyzeDeclArgList( accessMode, argList )
  self:checkToken( token, ")" )
  token = self:getToken(  )
  local retTypeList = {}
  local retTypeInfoList = {}
  if token.txt == ":" then
    repeat 
      local refType = self:analyzeRefType( accessMode )
      table.insert( retTypeList, refType )
      table.insert( retTypeInfoList, refType:get_expType() )
      token = self:getToken(  )
    until token.txt ~= ","
  end
  local typeInfo = TypeInfo.createFunc( typeKind, self:getCurrentNamespaceTypeInfo(  ), false, false, staticFlag, accessMode, name and name.txt or "", nil, retTypeInfoList )
  if name then
    scope:get_parent(  ):add( name.txt, typeInfo )
  end
  if not needPopFlag then
    self:pushNamespace( name and name.txt or "", typeInfo, scope )
  end
  local node = nil
  local info = nil
  if token.txt == ";" then
    node = self:createNoneNode( firstToken.pos )
  else 
    self:pushback(  )
    local body = self:analyzeBlock( "func", scope )
    info = DeclFuncInfo.new(classNameToken, name, argList, staticFlag, accessMode, body, retTypeList, retTypeInfoList)
    do
      local _switchExp = (kind )
      if _switchExp == nodeKindDeclConstr then
        node = DeclConstrNode.new(firstToken.pos, {typeInfo}, info)
      elseif _switchExp == nodeKindDeclMethod then
        node = DeclMethodNode.new(firstToken.pos, {typeInfo}, info)
      elseif _switchExp == nodeKindDeclFunc then
        node = DeclFuncNode.new(firstToken.pos, {typeInfo}, info)
      else 
        self:error( string.format( "illegal kind -- %d", kind) )
      end
    end
    
  end
  if not needPopFlag then
    self:popNamespace(  )
  end
  self:popScope(  )
  if needPopFlag then
    self:addMethod( classNameToken.txt, node, name.txt )
    self:popClass(  )
  end
  return node
end

function TransUnit:analyzeDeclVar( accessMode, staticFlag, firstToken )
  local unwrapFlag = false
  local token = self:getToken(  )
  if token.txt == "!" then
    unwrapFlag = true
  else 
    self:pushback(  )
  end
  local typeInfoList = {}
  local varNameList = {}
  local varTypeList = {}
  repeat 
    local varName = self:getSymbolToken(  )
    token = self:getToken(  )
    local typeInfo = typeInfoNone
    local refType = nil
    if token.txt == ":" then
      refType = self:analyzeRefType( accessMode )
      table.insert( varTypeList, refType )
      typeInfo = refType:get_expType()
      token = self:getToken(  )
    else 
      table.insert( varTypeList, nil )
    end
    table.insert( varNameList, varName )
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
    local expTypeList = {}
    for index, exp in pairs( expList.expList ) do
      if index == #expList.expList then
        for __index, typeInfo in pairs( exp:get_expTypeList() ) do
          table.insert( expTypeList, typeInfo )
        end
      else 
        typeInfoList[index] = exp:get_expType()
      end
    end
    for index, typeInfo in pairs( expTypeList ) do
      if not typeInfoList[index] or typeInfoList[index] == typeInfoNone then
        typeInfoList[index] = typeInfo
      end
    end
  end
  local unwrapBlock = nil
  if unwrapFlag then
    unwrapBlock = self:analyzeBlock( "let!" )
    for index, typeInfo in pairs( typeInfoList ) do
      if typeInfo:get_nilable(  ) then
        typeInfoList[index] = typeInfo:get_orgTypeInfo(  )
      end
    end
  end
  if self.macroScope == self.scope then
    for index, varName in pairs( varNameList ) do
      local typeInfo = typeInfoList[index]
      self.symbol2ValueMapForMacro[varName.txt] = MacroValInfo.new(nil, typeInfo)
    end
  end
  self:checkNextToken( ";" )
  local varList = {}
  for index, varName in pairs( varNameList ) do
    local typeInfo = typeInfoList[index]
    table.insert( varList, VarInfo.new(varName, varTypeList[index], typeInfo) )
    if not varTypeList[index] and typeInfo == typeInfoNil then
      self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
    end
    self.scope:add( varName.txt, typeInfo )
  end
  local node = DeclVarNode.new(firstToken.pos, {typeInfoNone}, accessMode, staticFlag, varList, expList, typeInfoList, unwrapBlock)
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
  return ExpListNode.new(firstExp.pos, {typeInfoNone}, expList)
end

function TransUnit:analyzeListConst( token )
  local nextToken = self:getToken(  )
  local expList = nil
  local itemTypeInfo = typeInfoNone
  if nextToken.txt ~= "]" then
    self:pushback(  )
    expList = self:analyzeExpList(  )
    self:checkNextToken( "]" )
    local nodeList = expList.expList
    for __index, exp in pairs( nodeList ) do
      if itemTypeInfo == typeInfoNone then
        itemTypeInfo = exp:get_expType()
      elseif itemTypeInfo ~= exp:get_expType() then
        itemTypeInfo = typeInfoStem
      end
    end
  end
  local kind = nodeKindLiteralArray
  local typeInfo = typeInfoNone
  if token.txt == '[' then
    kind = nodeKindLiteralList
    typeInfo = {TypeInfo.createList( "pri", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralListNode.new(token.pos, typeInfo, expList)
  else 
    typeInfo = {TypeInfo.createArray( "pri", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralArrayNode.new(token.pos, typeInfo, expList)
  end
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
    if key:get_expType() ~= keyTypeInfo then
      if keyTypeInfo ~= typeInfoNone then
        keyTypeInfo = typeInfoStem
      else 
        keyTypeInfo = key:get_expType()
      end
    end
    self:checkNextToken( ":" )
    local val = self:analyzeExp(  )
    if val:get_expType() ~= valTypeInfo then
      if valTypeInfo ~= typeInfoNone then
        valTypeInfo = typeInfoStem
      else 
        valTypeInfo = val:get_expType()
      end
    end
    table.insert( pairList, PairItem.new(key, val) )
    map[key] = val
    nextToken = self:getToken(  )
  until nextToken.txt ~= ","
  local typeInfo = TypeInfo.createMap( "pri", self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
  self:checkToken( nextToken, "}" )
  return LiteralMapNode.new(token.pos, {typeInfo}, map, pairList)
end

function TransUnit:analyzeExpRefItem( token, exp )
  local indexExp = self:analyzeExp(  )
  self:checkNextToken( "]" )
  local info = {["val"] = exp, ["index"] = indexExp}
  local typeInfo = typeInfoStem
  local expType = exp:get_expType()
  if expType then
    if expType.kind == TypeInfoKindMap then
      typeInfo = expType:get_itemTypeInfoList(  )[2]
    elseif expType.kind == TypeInfoKindArray or expType.kind == TypeInfoKindArray then
      typeInfo = expType:get_itemTypeInfoList(  )[1]
    end
  end
  return ExpRefItemNode.new(token.pos, {typeInfo}, exp, indexExp)
end

local MacroPaser = {}
setmetatable( MacroPaser, { __index = Parser } )
function MacroPaser.new( tokenList, name )
  local obj = {}
  setmetatable( obj, { __index = MacroPaser } )
  if obj.__init then obj:__init( tokenList, name ); end
return obj
end
function MacroPaser:__init(tokenList, name) 
  self.pos = 1
  self.tokenList = tokenList
  self.name = name
  self.parser = nil
end
function MacroPaser:getToken(  )
  if #self.tokenList < self.pos then
    return nil
  end
  local token = self.tokenList[self.pos]
  self.pos = self.pos + 1
  return token
end
function MacroPaser:getStreamName(  )
  return self.name
end

function TransUnit:evalMacro( firstToken, macroTypeInfo, expList )
  if expList and expList:get_expList(  ) then
    for __index, exp in pairs( expList:get_expList(  ) ) do
      local kind = exp.kind
      if kind ~= nodeKindLiteralNil and kind ~= nodeKindLiteralChar and kind ~= nodeKindLiteralInt and kind ~= nodeKindLiteralReal and kind ~= nodeKindLiteralArray and kind ~= nodeKindLiteralList and kind ~= nodeKindLiteralMap and kind ~= nodeKindLiteralString and kind ~= nodeKindLiteralBool and kind ~= nodeKindLiteralSymbol and kind ~= nodeKindRefField and kind ~= nodeKindExpMacroStat then
        self:error( "Macro arguments must be literal value." )
      end
    end
  end
  local macroInfo = self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )]
  local argVal = {}
  local argType = {}
  if expList then
    for __index, argNode in pairs( expList:get_expList(  ) ) do
      local valList, typeInfoList = argNode:getLiteral(  )
      local val = valList[1]
      local typeInfo = typeInfoList[1]
      table.insert( argVal, val )
      table.insert( argType, typeInfo )
    end
  end
  local func = macroInfo.func
  local macroVars = func( table.unpack( argVal ) )
  for __index, name in pairs( macroVars._names ) do
    local valInfo = macroInfo.symbol2MacroValInfoMap[name]
    local typeInfo = valInfo and valInfo.typeInfo or typeInfoStem
    local val = macroVars[name]
    if typeInfo == typeInfoSymbol then
      val = {val}
    end
    self.symbol2ValueMapForMacro[name] = MacroValInfo.new(val, typeInfo)
  end
  local argList = macroInfo.declInfo:get_argList(  )
  if argList then
    for index, arg in pairs( argList ) do
      if arg:get_kind(  ) == nodeKindDeclArg then
        local argInfo = arg
        local argType = argInfo:get_argType()
        local argName = argInfo:get_name().txt
        self.symbol2ValueMapForMacro[argName] = MacroValInfo.new(argVal[index], argType:get_expType())
      else 
        self:error( "not support ... in macro" )
      end
    end
  end
  local parser = MacroPaser.new(macroInfo.declInfo.tokenList, "macro")
  local bakParser = self.parser
  self.parser = parser
  self.macroMode = "expand"
  local stmtList = {}
  self:analyzeStatementList( stmtList, "}" )
  self.macroMode = "none"
  self.parser = bakParser
  return ExpMacroExpNode.new(firstToken.pos, typeInfoNone, stmtList)
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
        local macroFlag = false
        if exp:get_expType():get_kind(  ) == TypeInfoKindMacro then
          macroFlag = true
          self.symbol2ValueMapForMacro = {}
          self.macroMode = "analyze"
        end
        matchFlag = true
        local work = self:getToken(  )
        local expList = nil
        if work.txt ~= ")" then
          self:pushback(  )
          expList = self:analyzeExpList(  )
          self:checkNextToken( ")" )
        end
        local info = {["func"] = exp, ["argList"] = expList}
        if macroFlag then
          self.macroMode = "none"
          exp = self:evalMacro( firstToken, exp:get_expType(), expList )
        else 
          exp = ExpCallNode.new(firstToken.pos, exp:get_expType():get_retTypeInfoList(  ), exp, expList)
        end
        nextToken = self:getToken(  )
      end
    until not matchFlag
  end
  if nextToken.txt == "." then
    return self:analyzeExpSymbol( firstToken, self:getToken(  ), "field", exp, skipFlag )
  elseif nextToken.txt == ".$" then
    return self:analyzeExpSymbol( firstToken, self:getToken(  ), "get", exp, skipFlag )
  end
  self:pushback(  )
  return exp
end

function TransUnit:analyzeExpSymbol( firstToken, token, mode, prefixExp, skipFlag )
  local exp = nil
  if mode == "field" or mode == "get" then
    local typeInfo = typeInfoStem
    if not prefixExp:get_expType() then
      self:error( "unknown prefix type: " .. getNodeKindName( prefixExp.kind ) )
    end
    local getterTypeInfo = nil
    if prefixExp:get_expType():get_kind(  ) == TypeInfoKindClass then
      local classScope = self.typeId2Scope[prefixExp:get_expType():get_typeId(  )]
      local className = prefixExp:get_expType():getTxt(  )
      if not classScope then
        self:error( string.format( "not found field: %s, %s", className, prefixExp:get_expType()) )
      end
      typeInfo = nil
      if mode == "get" then
        typeInfo = classScope:getTypeInfo( string.format( "get_%s", token.txt) )
        if typeInfo and (typeInfo:get_kind(  ) == TypeInfoKindFunc or typeInfo:get_kind(  ) == TypeInfoKindMethod ) then
          local retTypeList = typeInfo:get_retTypeInfoList(  )
          getterTypeInfo = typeInfo
          typeInfo = retTypeList[1]
        end
      end
      if not getterTypeInfo then
        typeInfo = classScope:getTypeInfo( token.txt )
      end
      if not typeInfo then
        print( "hoge", classScope.symbol2TypeInfoMap )
        for name, val in pairs( classScope.symbol2TypeInfoMap ) do
          print( "hoge", name, val )
        end
        self:error( string.format( "not found field typeInfo: %s.%s %s", className, token.txt, classScope ) )
      end
    end
    if not getterTypeInfo then
      exp = RefFieldNode.new(firstToken.pos, {typeInfo}, token, prefixExp)
    else 
      exp = GetFieldNode.new(firstToken.pos, {typeInfo}, token, prefixExp, getterTypeInfo)
    end
  elseif mode == "symbol" then
    if self.macroMode == "analyze" then
      exp = LiteralSymbolNode.new(firstToken.pos, {typeInfoSymbol}, token)
    else 
      local typeInfo = self.scope:getTypeInfo( token.txt )
      if not typeInfo and token.txt == "self" then
        local namespaceInfo = self.classList[#self.classList]
        typeInfo = namespaceInfo.typeInfo
      end
      if not typeInfo then
        self:error( "not found type -- " .. token.txt )
      end
      if typeInfo == typeInfoSymbol then
        skipFlag = true
      end
      exp = ExpRefNode.new(firstToken.pos, {typeInfo}, token)
    end
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( false, "pri", false, nil, token, nil )
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
      local castType = self:analyzeRefType( "pri" )
      exp = ExpCastNode.new(firstToken.pos, castType:get_expTypeList(), exp)
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
        local opTxt = nextToken.txt
        if not exp:get_expType() or not exp2:get_expType() then
          self:error( string.format( "illegal exp or exp2 %s, %s, %s , %s,%d:%d", exp:get_expType(), exp2:get_expType(), nextToken.txt, self.parser:getStreamName(  ), nextToken.pos.lineNo, nextToken.pos.column) )
        end
        local expType = typeInfoNone
        if not exp:get_expType() then
          self:error( string.format( "expType is nil %s:%d:%d", self.parser:getStreamName(  ), firstToken.pos.lineNo, firstToken.pos.column) )
        end
        do
          local _switchExp = opTxt
          if _switchExp == "or" or _switchExp == "and" then
            if exp:get_expType():equals( exp2:get_expType() ) then
              expType = exp:get_expType()
            else 
              expType = typeInfoStem
            end
          elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" or _switchExp == "~=" or _switchExp == "==" or _switchExp == "not" then
            expType = typeInfoBool
          elseif _switchExp == "^" or _switchExp == "|" or _switchExp == "~" or _switchExp == "&" or _switchExp == "<<" or _switchExp == ">>" or _switchExp == "#" then
            expType = typeInfoInt
          elseif _switchExp == ".." then
            expType = typeInfoString
          elseif _switchExp == "+" or _switchExp == "-" or _switchExp == "*" or _switchExp == "/" or _switchExp == "//" or _switchExp == "%" then
            if exp:get_expType() == typeInfoReal or exp2:get_expType() == typeInfoReal then
              expType = typeInfoReal
            else 
              expType = typeInfoInt
            end
          elseif _switchExp == "=" then
          else 
            self:error( "unknown op " .. opTxt )
          end
        end
        
        exp = ExpOp2Node.new(firstToken.pos, {expType}, nextToken, exp, exp2)
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

function TransUnit:analyzeExpMacroStat( firstToken )
  local expStrList = {}
  self:checkNextToken( "{" )
  local braceCount = 0
  while true do
    local token = self:getToken(  )
    if token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
      local exp = self:analyzeExp( true, op1levelMap[token.txt] )
      local format = token.txt == ",,," and "'%s '" or '"\'%s\'"'
      if token.txt == ",," and exp.kind == nodeKindExpRef then
        local refToken = (exp ):get_token(  )
        local macroInfo = self.symbol2ValueMapForMacro[refToken.txt]
        if macroInfo and macroInfo.typeInfo == typeInfoSymbol then
          format = "'%s '"
        end
      end
      local newToken = Parser.Token.new(Parser.kind.Str, format, token.pos)
      local literalStr = LiteralStringNode.new(token.pos, {typeInfoString}, newToken, {exp})
      table.insert( expStrList, literalStr )
    else 
      if token.txt == "{" then
        braceCount = braceCount + 1
      elseif token.txt == "}" then
        if braceCount == 0 then
          break
        end
        braceCount = braceCount - 1
      end
      local newToken = Parser.Token.new(token.kind, string.format( "'%s '", token.txt ), token.pos)
      local literalStr = LiteralStringNode.new(token.pos, {typeInfoString}, newToken, nil)
      table.insert( expStrList, literalStr )
    end
  end
  return ExpMacroStatNode.new(firstToken.pos, {typeInfoStat}, expStrList)
end

function TransUnit:analyzeSuper( firstToken )
  self:checkNextToken( "(" )
  local expList = self:analyzeExpList(  )
  self:checkNextToken( ")" )
  self:checkNextToken( ";" )
  local classType = self:getCurrentClass(  )
  local superType = classType:get_baseTypeInfo(  )
  return ExpCallSuperNode.new(firstToken.pos, {typeInfoNone}, superType, expList)
end

function TransUnit:analyzeExp( skipOp2Flag, prevOpLevel )
  local firstToken = self:getToken(  )
  local token = firstToken
  local exp = nil
  if token.kind == Parser.kind.Dlmt then
    if token.txt == "..." then
      return ExpDDDNode.new(firstToken.pos, {typeInfoNone}, token)
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
      exp = ExpParenNode.new(firstToken.pos, exp:get_expTypeList(), exp)
      exp = self:analyzeExpCont( firstToken, exp, false )
    end
  end
  if token.txt == "new" then
    exp = self:analyzeRefType( "pri" )
    self:checkNextToken( "(" )
    local nextToken = self:getToken(  )
    local argList = nil
    if nextToken.txt ~= ")" then
      self:pushback(  )
      argList = self:analyzeExpList(  )
      self:checkNextToken( ")" )
    end
    exp = ExpNewNode.new(firstToken.pos, exp:get_expTypeList(), exp, argList)
    exp = self:analyzeExpCont( firstToken, exp, false )
  end
  if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
    if token.txt == "`" then
      exp = self:analyzeExpMacroStat( token )
    else 
      exp = self:analyzeExp( true, op1levelMap[token.txt] )
      local typeInfo = typeInfoNone
      do
        local _switchExp = (token.txt )
        if _switchExp == "-" then
          if exp:get_expType() ~= typeInfoInt and exp:get_expType() ~= typeInfoReal then
            self:addErrMess( token.pos, string.format( 'unmatch type for "-" -- %s', exp:get_expType():getTxt(  )) )
          end
          typeInfo = exp:get_expType()
        elseif _switchExp == "#" then
          typeInfo = typeInfoInt
        elseif _switchExp == "not" then
          typeInfo = typeInfoBool
        elseif _switchExp == ",," then
        elseif _switchExp == ",,," then
          if exp:get_expType() ~= typeInfoString then
            self:error( "unmatch ,,, type, need string type" )
          end
          typeInfo = typeInfoSymbol
        elseif _switchExp == ",,,," then
          if exp:get_expType() ~= typeInfoSymbol then
            self:error( "unmatch ,,, type, need symbol type" )
          end
          typeInfo = typeInfoString
        elseif _switchExp == "`" then
          typeInfo = typeInfoNone
        elseif _switchExp == "not" then
          typeInfo = typeInfoBool
        else 
          self:error( "unknown op1" )
        end
      end
      
      exp = ExpOp1Node.new(firstToken.pos, {typeInfo}, token, exp)
      return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
    end
  end
  if token.kind == Parser.kind.Int then
    exp = LiteralIntNode.new(firstToken.pos, {typeInfoInt}, token, tonumber( token.txt ))
  elseif token.kind == Parser.kind.Real then
    exp = LiteralRealNode.new(firstToken.pos, {typeInfoReal}, token, tonumber( token.txt ))
  elseif token.kind == Parser.kind.Char then
    local num = 0
    if #(token.txt ) == 1 then
      num = token.txt:byte( 1 )
    else 
      num = quotedChar2Code[token.txt:sub( 2, 2 )]
    end
    exp = LiteralCharNode.new(firstToken.pos, {typeInfoChar}, token, num)
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
    exp = LiteralStringNode.new(firstToken.pos, {typeInfoString}, token, formatArgList)
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
    exp = ExpRefNode.new(firstToken.pos, {typeInfoNone}, token)
  elseif token.txt == "true" or token.txt == "false" then
    exp = LiteralBoolNode.new(firstToken.pos, {typeInfoBool}, token)
  elseif token.txt == "nil" then
    exp = LiteralNilNode.new(firstToken.pos, {typeInfoNil})
  end
  if not exp then
    self:error( "illegal exp" )
  end
  if skipOp2Flag then
    return exp
  end
  return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeStatement( termTxt )
  local token = self:getTokenNoErr(  )
  if not token then
    return nil
  end
  local statement = self:analyzeDecl( "pri", false, token, token )
  if not statement then
    if token.txt == termTxt then
      self:pushback(  )
      return nil
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
    elseif token.txt == "super" then
      statement = self:analyzeSuper( token )
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
      local nextToken = self:getToken(  )
      local expList = nil
      if nextToken.txt ~= ";" then
        self:pushback(  )
        expList = self:analyzeExpList(  )
        self:checkNextToken( ";" )
      end
      statement = ReturnNode.new(token.pos, {typeInfoNone}, expList)
    elseif token.txt == "break" then
      self:checkNextToken( ";" )
      statement = BreakNode.new(token.pos, {typeInfoNone})
    elseif token.txt == "import" then
      statement = self:analyzeImport( token )
    elseif token.txt == ";" then
      statement = self:createNoneNode( token.pos )
    elseif token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
      self:error( string.format( "illegal macro op -- %s", token.txt) )
    else 
      self:pushback(  )
      local exp = self:analyzeExp(  )
      self:checkNextToken( ";" )
      statement = StmtExpNode.new(self.currentToken.pos, {typeInfoNone}, exp)
    end
  end
  return statement
end

function TransUnit:analyzeStatementList( stmtList, termTxt )
  while true do
    local statement = self:analyzeStatement( termTxt )
    if not statement then
      break
    end
    table.insert( stmtList, statement )
  end
end

----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo2734 = {}
  _className2InfoMap.ASTInfo = _classInfo2734
  _typeId2ClassInfoMap[ 2734 ] = _classInfo2734
  end
do
  local _classInfo1236 = {}
  _className2InfoMap.ApplyNode = _classInfo1236
  _typeId2ClassInfoMap[ 1236 ] = _classInfo1236
  end
do
  local _classInfo1020 = {}
  _className2InfoMap.BlockNode = _classInfo1020
  _typeId2ClassInfoMap[ 1020 ] = _classInfo1020
  end
do
  local _classInfo1350 = {}
  _className2InfoMap.BreakNode = _classInfo1350
  _typeId2ClassInfoMap[ 1350 ] = _classInfo1350
  end
do
  local _classInfo1104 = {}
  _className2InfoMap.CaseInfo = _classInfo1104
  _typeId2ClassInfoMap[ 1104 ] = _classInfo1104
  end
do
  local _classInfo1930 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo1930
  _typeId2ClassInfoMap[ 1930 ] = _classInfo1930
  end
do
  local _classInfo1912 = {}
  _className2InfoMap.DeclArgNode = _classInfo1912
  _typeId2ClassInfoMap[ 1912 ] = _classInfo1912
  end
do
  local _classInfo538 = {}
  _className2InfoMap.DeclClassNode = _classInfo538
  _typeId2ClassInfoMap[ 538 ] = _classInfo538
  end
do
  local _classInfo1826 = {}
  _className2InfoMap.DeclConstrNode = _classInfo1826
  _typeId2ClassInfoMap[ 1826 ] = _classInfo1826
  end
do
  local _classInfo1752 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo1752
  _typeId2ClassInfoMap[ 1752 ] = _classInfo1752
  end
do
  local _classInfo1786 = {}
  _className2InfoMap.DeclFuncNode = _classInfo1786
  _typeId2ClassInfoMap[ 1786 ] = _classInfo1786
  end
do
  local _classInfo706 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo706
  _typeId2ClassInfoMap[ 706 ] = _classInfo706
  end
do
  local _classInfo2004 = {}
  _className2InfoMap.DeclMacroNode = _classInfo2004
  _typeId2ClassInfoMap[ 2004 ] = _classInfo2004
  end
do
  local _classInfo1880 = {}
  _className2InfoMap.DeclMemberNode = _classInfo1880
  _typeId2ClassInfoMap[ 1880 ] = _classInfo1880
  end
do
  local _classInfo1806 = {}
  _className2InfoMap.DeclMethodNode = _classInfo1806
  _typeId2ClassInfoMap[ 1806 ] = _classInfo1806
  end
do
  local _classInfo1714 = {}
  _className2InfoMap.DeclVarNode = _classInfo1714
  _typeId2ClassInfoMap[ 1714 ] = _classInfo1714
  end
do
  local _classInfo1510 = {}
  _className2InfoMap.ExpCallNode = _classInfo1510
  _typeId2ClassInfoMap[ 1510 ] = _classInfo1510
  end
do
  local _classInfo1848 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo1848
  _typeId2ClassInfoMap[ 1848 ] = _classInfo1848
  end
do
  local _classInfo1440 = {}
  _className2InfoMap.ExpCastNode = _classInfo1440
  _typeId2ClassInfoMap[ 1440 ] = _classInfo1440
  end
do
  local _classInfo1532 = {}
  _className2InfoMap.ExpDDDNode = _classInfo1532
  _typeId2ClassInfoMap[ 1532 ] = _classInfo1532
  end
do
  local _classInfo1082 = {}
  _className2InfoMap.ExpListNode = _classInfo1082
  _typeId2ClassInfoMap[ 1082 ] = _classInfo1082
  end
do
  local _classInfo1572 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo1572
  _typeId2ClassInfoMap[ 1572 ] = _classInfo1572
  end
do
  local _classInfo1598 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo1598
  _typeId2ClassInfoMap[ 1598 ] = _classInfo1598
  end
do
  local _classInfo1370 = {}
  _className2InfoMap.ExpNewNode = _classInfo1370
  _typeId2ClassInfoMap[ 1370 ] = _classInfo1370
  end
do
  local _classInfo1462 = {}
  _className2InfoMap.ExpOp1Node = _classInfo1462
  _typeId2ClassInfoMap[ 1462 ] = _classInfo1462
  end
do
  local _classInfo1416 = {}
  _className2InfoMap.ExpOp2Node = _classInfo1416
  _typeId2ClassInfoMap[ 1416 ] = _classInfo1416
  end
do
  local _classInfo1552 = {}
  _className2InfoMap.ExpParenNode = _classInfo1552
  _typeId2ClassInfoMap[ 1552 ] = _classInfo1552
  end
do
  local _classInfo1486 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo1486
  _typeId2ClassInfoMap[ 1486 ] = _classInfo1486
  end
do
  local _classInfo1392 = {}
  _className2InfoMap.ExpRefNode = _classInfo1392
  _typeId2ClassInfoMap[ 1392 ] = _classInfo1392
  end
do
  local _classInfo680 = {}
  _className2InfoMap.Filter = _classInfo680
  _typeId2ClassInfoMap[ 680 ] = _classInfo680
  end
do
  local _classInfo1204 = {}
  _className2InfoMap.ForNode = _classInfo1204
  _typeId2ClassInfoMap[ 1204 ] = _classInfo1204
  end
do
  local _classInfo1272 = {}
  _className2InfoMap.ForeachNode = _classInfo1272
  _typeId2ClassInfoMap[ 1272 ] = _classInfo1272
  end
do
  local _classInfo1306 = {}
  _className2InfoMap.ForsortNode = _classInfo1306
  _typeId2ClassInfoMap[ 1306 ] = _classInfo1306
  end
do
  local _classInfo1672 = {}
  _className2InfoMap.GetFieldNode = _classInfo1672
  _typeId2ClassInfoMap[ 1672 ] = _classInfo1672
  end
do
  local _classInfo1056 = {}
  _className2InfoMap.IfNode = _classInfo1056
  _typeId2ClassInfoMap[ 1056 ] = _classInfo1056
  end
do
  local _classInfo1044 = {}
  _className2InfoMap.IfStmtInfo = _classInfo1044
  _typeId2ClassInfoMap[ 1044 ] = _classInfo1044
  end
do
  local _classInfo930 = {}
  _className2InfoMap.ImportNode = _classInfo930
  _typeId2ClassInfoMap[ 930 ] = _classInfo930
  end
do
  local _classInfo2110 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo2110
  _typeId2ClassInfoMap[ 2110 ] = _classInfo2110
  end
do
  local _classInfo2222 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo2222
  _typeId2ClassInfoMap[ 2222 ] = _classInfo2222
  end
do
  local _classInfo2040 = {}
  _className2InfoMap.LiteralCharNode = _classInfo2040
  _typeId2ClassInfoMap[ 2040 ] = _classInfo2040
  end
do
  local _classInfo2064 = {}
  _className2InfoMap.LiteralIntNode = _classInfo2064
  _typeId2ClassInfoMap[ 2064 ] = _classInfo2064
  end
do
  local _classInfo2130 = {}
  _className2InfoMap.LiteralListNode = _classInfo2130
  _typeId2ClassInfoMap[ 2130 ] = _classInfo2130
  end
do
  local _classInfo2158 = {}
  _className2InfoMap.LiteralMapNode = _classInfo2158
  _typeId2ClassInfoMap[ 2158 ] = _classInfo2158
  end
do
  local _classInfo2020 = {}
  _className2InfoMap.LiteralNilNode = _classInfo2020
  _typeId2ClassInfoMap[ 2020 ] = _classInfo2020
  end
do
  local _classInfo2088 = {}
  _className2InfoMap.LiteralRealNode = _classInfo2088
  _typeId2ClassInfoMap[ 2088 ] = _classInfo2088
  end
do
  local _classInfo2194 = {}
  _className2InfoMap.LiteralStringNode = _classInfo2194
  _typeId2ClassInfoMap[ 2194 ] = _classInfo2194
  end
do
  local _classInfo2242 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo2242
  _typeId2ClassInfoMap[ 2242 ] = _classInfo2242
  end
do
  local _classInfo704 = {}
  _className2InfoMap.MacroEval = _classInfo704
  _typeId2ClassInfoMap[ 704 ] = _classInfo704
  end
do
  local _classInfo702 = {}
  _className2InfoMap.NamespaceInfo = _classInfo702
  _typeId2ClassInfoMap[ 702 ] = _classInfo702
  _classInfo702.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  _classInfo702.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 640 }
  _classInfo702.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 516 }
  end
do
  local _classInfo536 = {}
  _className2InfoMap.Node = _classInfo536
  _typeId2ClassInfoMap[ 536 ] = _classInfo536
  end
do
  local _classInfo682 = {}
  _className2InfoMap.NodePos = _classInfo682
  _typeId2ClassInfoMap[ 682 ] = _classInfo682
  end
do
  local _classInfo912 = {}
  _className2InfoMap.NoneNode = _classInfo912
  _typeId2ClassInfoMap[ 912 ] = _classInfo912
  end
do
  local _classInfo2146 = {}
  _className2InfoMap.PairItem = _classInfo2146
  _typeId2ClassInfoMap[ 2146 ] = _classInfo2146
  end
do
  local _classInfo1646 = {}
  _className2InfoMap.RefFieldNode = _classInfo1646
  _typeId2ClassInfoMap[ 1646 ] = _classInfo1646
  end
do
  local _classInfo992 = {}
  _className2InfoMap.RefTypeNode = _classInfo992
  _typeId2ClassInfoMap[ 992 ] = _classInfo992
  end
do
  local _classInfo1174 = {}
  _className2InfoMap.RepeatNode = _classInfo1174
  _typeId2ClassInfoMap[ 1174 ] = _classInfo1174
  end
do
  local _classInfo1334 = {}
  _className2InfoMap.ReturnNode = _classInfo1334
  _typeId2ClassInfoMap[ 1334 ] = _classInfo1334
  end
do
  local _classInfo952 = {}
  _className2InfoMap.RootNode = _classInfo952
  _typeId2ClassInfoMap[ 952 ] = _classInfo952
  end
do
  local _classInfo640 = {}
  _className2InfoMap.Scope = _classInfo640
  _typeId2ClassInfoMap[ 640 ] = _classInfo640
  end
do
  local _classInfo1624 = {}
  _className2InfoMap.StmtExpNode = _classInfo1624
  _typeId2ClassInfoMap[ 1624 ] = _classInfo1624
  end
do
  local _classInfo1118 = {}
  _className2InfoMap.SwitchNode = _classInfo1118
  _typeId2ClassInfoMap[ 1118 ] = _classInfo1118
  end
do
  local _classInfo730 = {}
  _className2InfoMap.TransUnit = _classInfo730
  _typeId2ClassInfoMap[ 730 ] = _classInfo730
  end
do
  local _classInfo516 = {}
  _className2InfoMap.TypeInfo = _classInfo516
  _typeId2ClassInfoMap[ 516 ] = _classInfo516
  end
do
  local _classInfo1692 = {}
  _className2InfoMap.VarInfo = _classInfo1692
  _typeId2ClassInfoMap[ 1692 ] = _classInfo1692
  end
do
  local _classInfo1150 = {}
  _className2InfoMap.WhileNode = _classInfo1150
  _typeId2ClassInfoMap[ 1150 ] = _classInfo1150
  end
do
  local _classInfo374 = {}
  _className2InfoMap.Parser = _classInfo374
  _typeId2ClassInfoMap[ 374 ] = _classInfo374
  end
do
  local _classInfo364 = {}
  _className2InfoMap.Position = _classInfo364
  _typeId2ClassInfoMap[ 364 ] = _classInfo364
  _classInfo364.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo364.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  end
do
  local _classInfo354 = {}
  _className2InfoMap.Stream = _classInfo354
  _typeId2ClassInfoMap[ 354 ] = _classInfo354
  end
do
  local _classInfo386 = {}
  _className2InfoMap.StreamParser = _classInfo386
  _typeId2ClassInfoMap[ 386 ] = _classInfo386
  _classInfo386.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 390 }
  end
do
  local _classInfo366 = {}
  _className2InfoMap.Token = _classInfo366
  _typeId2ClassInfoMap[ 366 ] = _classInfo366
  _classInfo366.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo366.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 364 }
  _classInfo366.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  end
do
  local _classInfo358 = {}
  _className2InfoMap.TxtStream = _classInfo358
  _typeId2ClassInfoMap[ 358 ] = _classInfo358
  end
do
  local _classInfo404 = {}
  _className2InfoMap.Util = _classInfo404
  _typeId2ClassInfoMap[ 404 ] = _classInfo404
  _classInfo404.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 512 }
  _classInfo404.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 510 }
  _classInfo404.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 502 }
  _classInfo404.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 498 }
  _classInfo404.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 514 }
  end
do
  local _classInfo380 = {}
  _className2InfoMap.WrapParser = _classInfo380
  _typeId2ClassInfoMap[ 380 ] = _classInfo380
  end
do
  local _classInfo108 = {}
  _className2InfoMap.base = _classInfo108
  _typeId2ClassInfoMap[ 108 ] = _classInfo108
  _classInfo108.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  _classInfo108.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 404 }
  end
do
  local _classInfo106 = {}
  _className2InfoMap.lune = _classInfo106
  _typeId2ClassInfoMap[ 106 ] = _classInfo106
  _classInfo106.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 108 }
  end
do
  local _classInfo502 = {}
  _className2InfoMap.memStream = _classInfo502
  _typeId2ClassInfoMap[ 502 ] = _classInfo502
  end
do
  local _classInfo498 = {}
  _className2InfoMap.outStream = _classInfo498
  _typeId2ClassInfoMap[ 498 ] = _classInfo498
  end
local _varName2InfoMap = {}
moduleObj._varName2InfoMap = _varName2InfoMap
_varName2InfoMap.TypeInfoKindArray = {
  name='TypeInfoKindArray', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindClass = {
  name='TypeInfoKindClass', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindFunc = {
  name='TypeInfoKindFunc', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindList = {
  name='TypeInfoKindList', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindMacro = {
  name='TypeInfoKindMacro', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindMap = {
  name='TypeInfoKindMap', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindMethod = {
  name='TypeInfoKindMethod', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindModule = {
  name='TypeInfoKindModule', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindNilable = {
  name='TypeInfoKindNilable', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindPrim = {
  name='TypeInfoKindPrim', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindRoot = {
  name='TypeInfoKindRoot', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.nodeKind = {
  name='nodeKind', accessMode = 'pub', typeId = 882 }
_varName2InfoMap.rootTypeId = {
  name='rootTypeId', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.typeInfoKind = {
  name='typeInfoKind', accessMode = 'pub', typeId = 518 }
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 100, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {102}, }
_typeInfoList[2] = { parentId = 100, typeId = 102, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[3] = { parentId = 102, typeId = 104, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {106, 516, 530, 536, 538, 640, 680, 682, 702, 704, 706, 730, 888, 912, 930, 952, 992, 1020, 1044, 1056, 1082, 1104, 1118, 1150, 1174, 1204, 1236, 1272, 1306, 1334, 1350, 1370, 1392, 1416, 1440, 1462, 1486, 1510, 1532, 1552, 1572, 1598, 1624, 1646, 1672, 1692, 1714, 1752, 1786, 1806, 1826, 1848, 1880, 1912, 1930, 2004, 2020, 2040, 2064, 2088, 2110, 2130, 2146, 2158, 2194, 2222, 2242, 2734}, }
_typeInfoList[4] = { parentId = 104, typeId = 106, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {108}, }
_typeInfoList[5] = { parentId = 106, typeId = 108, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {110, 404}, }
_typeInfoList[6] = { parentId = 108, typeId = 110, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {354, 358, 364, 366, 374, 380, 386, 394, 396, 398, 402}, }
_typeInfoList[7] = { parentId = 110, typeId = 354, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {356}, }
_typeInfoList[8] = { parentId = 354, typeId = 356, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[9] = { parentId = 110, typeId = 358, baseId = 354, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {360, 362}, }
_typeInfoList[10] = { parentId = 358, typeId = 360, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[11] = { parentId = 358, typeId = 362, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[12] = { parentId = 110, typeId = 364, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[13] = { parentId = 110, typeId = 366, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {368, 372}, }
_typeInfoList[14] = { parentId = 366, typeId = 368, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 1, typeId = 370, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {366}, retTypeId = {}, children = {}, }
_typeInfoList[16] = { parentId = 366, typeId = 372, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {370}, children = {}, }
_typeInfoList[17] = { parentId = 110, typeId = 374, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {376, 378}, }
_typeInfoList[18] = { parentId = 374, typeId = 376, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[19] = { parentId = 374, typeId = 378, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[20] = { parentId = 110, typeId = 380, baseId = 374, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {382, 384}, }
_typeInfoList[21] = { parentId = 380, typeId = 382, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[22] = { parentId = 380, typeId = 384, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[23] = { parentId = 110, typeId = 386, baseId = 374, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {388, 390, 400}, }
_typeInfoList[24] = { parentId = 386, typeId = 388, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[25] = { parentId = 386, typeId = 390, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {386}, children = {}, }
_typeInfoList[26] = { parentId = 110, typeId = 394, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[27] = { parentId = 110, typeId = 396, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[28] = { parentId = 110, typeId = 398, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[29] = { parentId = 386, typeId = 400, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[30] = { parentId = 110, typeId = 402, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {6}, children = {}, }
_typeInfoList[31] = { parentId = 108, typeId = 404, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {498, 502, 510, 512, 514}, }
_typeInfoList[32] = { parentId = 404, typeId = 498, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {500}, }
_typeInfoList[33] = { parentId = 498, typeId = 500, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[34] = { parentId = 404, typeId = 502, baseId = 498, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {504, 506, 508}, }
_typeInfoList[35] = { parentId = 502, typeId = 504, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 502, typeId = 506, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[37] = { parentId = 502, typeId = 508, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[38] = { parentId = 404, typeId = 510, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 404, typeId = 512, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[40] = { parentId = 404, typeId = 514, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 104, typeId = 516, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {558, 560, 562, 568, 570, 572, 578, 580, 584, 588, 590, 594, 600, 608, 612, 614, 616, 618, 620, 622, 624, 626, 628, 630, 632, 636, 638}, }
_typeInfoList[42] = { parentId = 1, typeId = 518, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 516}, retTypeId = {}, children = {}, }
_typeInfoList[43] = { parentId = 104, typeId = 530, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[44] = { parentId = 104, typeId = 536, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {686, 692, 694, 696, 700}, }
_typeInfoList[45] = { parentId = 104, typeId = 538, baseId = 536, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1966, 1980, 1982, 1984, 1988, 1992, 1994, 1998}, }
_typeInfoList[46] = { parentId = 516, typeId = 558, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[47] = { parentId = 516, typeId = 560, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[48] = { parentId = 516, typeId = 562, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[49] = { parentId = 516, typeId = 568, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[50] = { parentId = 516, typeId = 570, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[51] = { parentId = 516, typeId = 572, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[52] = { parentId = 516, typeId = 578, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[53] = { parentId = 516, typeId = 580, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[54] = { parentId = 516, typeId = 584, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[55] = { parentId = 516, typeId = 588, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[56] = { parentId = 516, typeId = 590, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[57] = { parentId = 516, typeId = 594, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[58] = { parentId = 516, typeId = 600, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[59] = { parentId = 1, typeId = 606, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[60] = { parentId = 516, typeId = 608, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {606}, children = {}, }
_typeInfoList[61] = { parentId = 1, typeId = 610, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[62] = { parentId = 516, typeId = 612, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {610}, children = {}, }
_typeInfoList[63] = { parentId = 516, typeId = 614, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[64] = { parentId = 516, typeId = 616, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[65] = { parentId = 516, typeId = 618, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[66] = { parentId = 516, typeId = 620, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[67] = { parentId = 516, typeId = 622, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[68] = { parentId = 516, typeId = 624, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[69] = { parentId = 516, typeId = 626, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[70] = { parentId = 516, typeId = 628, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[71] = { parentId = 516, typeId = 630, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[72] = { parentId = 516, typeId = 632, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[73] = { parentId = 1, typeId = 634, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[74] = { parentId = 516, typeId = 636, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {634}, children = {}, }
_typeInfoList[75] = { parentId = 516, typeId = 638, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[76] = { parentId = 104, typeId = 640, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {656, 658, 660, 662, 664, 666, 668, 672, 676}, }
_typeInfoList[77] = { parentId = 640, typeId = 656, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[78] = { parentId = 640, typeId = 658, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[79] = { parentId = 640, typeId = 660, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {640}, children = {}, }
_typeInfoList[80] = { parentId = 640, typeId = 662, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[81] = { parentId = 640, typeId = 664, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[82] = { parentId = 640, typeId = 666, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[83] = { parentId = 640, typeId = 668, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {640}, children = {}, }
_typeInfoList[84] = { parentId = 1, typeId = 670, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 516}, retTypeId = {}, children = {}, }
_typeInfoList[85] = { parentId = 640, typeId = 672, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {670}, children = {}, }
_typeInfoList[86] = { parentId = 1, typeId = 674, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 640}, retTypeId = {}, children = {}, }
_typeInfoList[87] = { parentId = 640, typeId = 676, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {674}, children = {}, }
_typeInfoList[88] = { parentId = 104, typeId = 680, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {914, 932, 954, 994, 1022, 1058, 1084, 1120, 1152, 1176, 1206, 1238, 1274, 1308, 1336, 1352, 1372, 1394, 1418, 1442, 1464, 1488, 1512, 1534, 1554, 1574, 1600, 1626, 1648, 1674, 1716, 1788, 1808, 1828, 1850, 1882, 1914, 1932, 1958, 2006, 2022, 2042, 2066, 2090, 2112, 2132, 2160, 2196, 2224, 2244}, }
_typeInfoList[89] = { parentId = 104, typeId = 682, baseId = 1, txt = 'NodePos',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[90] = { parentId = 536, typeId = 686, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[91] = { parentId = 1, typeId = 688, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 1, typeId = 690, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 536, typeId = 692, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {688, 690}, children = {}, }
_typeInfoList[94] = { parentId = 536, typeId = 694, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[95] = { parentId = 536, typeId = 696, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 698, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[97] = { parentId = 536, typeId = 700, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {698}, children = {}, }
_typeInfoList[98] = { parentId = 104, typeId = 702, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 104, typeId = 704, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2398}, }
_typeInfoList[100] = { parentId = 104, typeId = 706, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {712, 716, 718, 722}, }
_typeInfoList[101] = { parentId = 706, typeId = 712, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[102] = { parentId = 1, typeId = 714, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[103] = { parentId = 706, typeId = 716, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {714}, children = {}, }
_typeInfoList[104] = { parentId = 706, typeId = 718, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[105] = { parentId = 1, typeId = 720, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {366}, retTypeId = {}, children = {}, }
_typeInfoList[106] = { parentId = 706, typeId = 722, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {720}, children = {}, }
_typeInfoList[107] = { parentId = 104, typeId = 730, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {832, 2740}, }
_typeInfoList[108] = { parentId = 1, typeId = 830, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[109] = { parentId = 730, typeId = 832, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {830}, children = {}, }
_typeInfoList[110] = { parentId = 1, typeId = 882, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 12}, retTypeId = {}, children = {}, }
_typeInfoList[111] = { parentId = 104, typeId = 888, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[112] = { parentId = 104, typeId = 912, baseId = 536, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {916, 924}, }
_typeInfoList[113] = { parentId = 680, typeId = 914, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[114] = { parentId = 912, typeId = 916, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 912, typeId = 924, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[116] = { parentId = 104, typeId = 930, baseId = 536, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {934, 942, 944}, }
_typeInfoList[117] = { parentId = 680, typeId = 932, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[118] = { parentId = 930, typeId = 934, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 930, typeId = 942, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 930, typeId = 944, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[121] = { parentId = 104, typeId = 952, baseId = 536, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {960, 972, 976, 980}, }
_typeInfoList[122] = { parentId = 680, typeId = 954, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 952, typeId = 960, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[124] = { parentId = 952, typeId = 972, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[125] = { parentId = 1, typeId = 974, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[126] = { parentId = 952, typeId = 976, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {974}, children = {}, }
_typeInfoList[127] = { parentId = 1, typeId = 978, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 702}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 952, typeId = 980, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {978}, children = {}, }
_typeInfoList[129] = { parentId = 104, typeId = 992, baseId = 536, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {996, 1004, 1006, 1008, 1010, 1012}, }
_typeInfoList[130] = { parentId = 680, typeId = 994, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[131] = { parentId = 992, typeId = 996, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[132] = { parentId = 992, typeId = 1004, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 992, typeId = 1006, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[134] = { parentId = 992, typeId = 1008, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[135] = { parentId = 992, typeId = 1010, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[136] = { parentId = 992, typeId = 1012, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[137] = { parentId = 104, typeId = 1020, baseId = 536, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1026, 1036, 1038, 1042}, }
_typeInfoList[138] = { parentId = 680, typeId = 1022, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 1020, typeId = 1026, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[140] = { parentId = 1020, typeId = 1036, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 1020, typeId = 1038, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[142] = { parentId = 1, typeId = 1040, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1020, typeId = 1042, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1040}, children = {}, }
_typeInfoList[144] = { parentId = 104, typeId = 1044, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1046, 1048, 1050}, }
_typeInfoList[145] = { parentId = 1044, typeId = 1046, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[146] = { parentId = 1044, typeId = 1048, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[147] = { parentId = 1044, typeId = 1050, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[148] = { parentId = 104, typeId = 1056, baseId = 536, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1062, 1072, 1076}, }
_typeInfoList[149] = { parentId = 680, typeId = 1058, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[150] = { parentId = 1056, typeId = 1062, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 1056, typeId = 1072, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 1, typeId = 1074, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1044}, retTypeId = {}, children = {}, }
_typeInfoList[153] = { parentId = 1056, typeId = 1076, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1074}, children = {}, }
_typeInfoList[154] = { parentId = 104, typeId = 1082, baseId = 536, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1088, 1098, 1102}, }
_typeInfoList[155] = { parentId = 680, typeId = 1084, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[156] = { parentId = 1082, typeId = 1088, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[157] = { parentId = 1082, typeId = 1098, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[158] = { parentId = 1, typeId = 1100, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 1082, typeId = 1102, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1100}, children = {}, }
_typeInfoList[160] = { parentId = 104, typeId = 1104, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1106, 1108}, }
_typeInfoList[161] = { parentId = 1104, typeId = 1106, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[162] = { parentId = 1104, typeId = 1108, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[163] = { parentId = 104, typeId = 1118, baseId = 536, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1124, 1134, 1136, 1140, 1142}, }
_typeInfoList[164] = { parentId = 680, typeId = 1120, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[165] = { parentId = 1118, typeId = 1124, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[166] = { parentId = 1118, typeId = 1134, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[167] = { parentId = 1118, typeId = 1136, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[168] = { parentId = 1, typeId = 1138, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1104}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 1118, typeId = 1140, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1138}, children = {}, }
_typeInfoList[170] = { parentId = 1118, typeId = 1142, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[171] = { parentId = 104, typeId = 1150, baseId = 536, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1154, 1162, 1164, 1166}, }
_typeInfoList[172] = { parentId = 680, typeId = 1152, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[173] = { parentId = 1150, typeId = 1154, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[174] = { parentId = 1150, typeId = 1162, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[175] = { parentId = 1150, typeId = 1164, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[176] = { parentId = 1150, typeId = 1166, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[177] = { parentId = 104, typeId = 1174, baseId = 536, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1178, 1186, 1188, 1190}, }
_typeInfoList[178] = { parentId = 680, typeId = 1176, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[179] = { parentId = 1174, typeId = 1178, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[180] = { parentId = 1174, typeId = 1186, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 1174, typeId = 1188, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[182] = { parentId = 1174, typeId = 1190, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[183] = { parentId = 104, typeId = 1204, baseId = 536, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1208, 1216, 1218, 1220, 1222, 1224, 1226}, }
_typeInfoList[184] = { parentId = 680, typeId = 1206, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[185] = { parentId = 1204, typeId = 1208, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[186] = { parentId = 1204, typeId = 1216, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 1204, typeId = 1218, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[188] = { parentId = 1204, typeId = 1220, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[189] = { parentId = 1204, typeId = 1222, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[190] = { parentId = 1204, typeId = 1224, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[191] = { parentId = 1204, typeId = 1226, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[192] = { parentId = 104, typeId = 1236, baseId = 536, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1242, 1252, 1256, 1258, 1260}, }
_typeInfoList[193] = { parentId = 680, typeId = 1238, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[194] = { parentId = 1236, typeId = 1242, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 1236, typeId = 1252, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[196] = { parentId = 1, typeId = 1254, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {366}, retTypeId = {}, children = {}, }
_typeInfoList[197] = { parentId = 1236, typeId = 1256, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1254}, children = {}, }
_typeInfoList[198] = { parentId = 1236, typeId = 1258, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[199] = { parentId = 1236, typeId = 1260, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[200] = { parentId = 104, typeId = 1272, baseId = 536, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1276, 1284, 1286, 1288, 1290, 1292}, }
_typeInfoList[201] = { parentId = 680, typeId = 1274, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[202] = { parentId = 1272, typeId = 1276, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[203] = { parentId = 1272, typeId = 1284, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[204] = { parentId = 1272, typeId = 1286, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[205] = { parentId = 1272, typeId = 1288, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[206] = { parentId = 1272, typeId = 1290, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[207] = { parentId = 1272, typeId = 1292, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[208] = { parentId = 104, typeId = 1306, baseId = 536, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1310, 1318, 1320, 1322, 1324, 1326, 1328}, }
_typeInfoList[209] = { parentId = 680, typeId = 1308, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[210] = { parentId = 1306, typeId = 1310, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[211] = { parentId = 1306, typeId = 1318, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[212] = { parentId = 1306, typeId = 1320, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[213] = { parentId = 1306, typeId = 1322, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[214] = { parentId = 1306, typeId = 1324, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[215] = { parentId = 1306, typeId = 1326, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[216] = { parentId = 1306, typeId = 1328, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[217] = { parentId = 104, typeId = 1334, baseId = 536, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1338, 1346, 1348}, }
_typeInfoList[218] = { parentId = 680, typeId = 1336, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[219] = { parentId = 1334, typeId = 1338, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[220] = { parentId = 1334, typeId = 1346, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[221] = { parentId = 1334, typeId = 1348, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[222] = { parentId = 104, typeId = 1350, baseId = 536, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1354, 1362}, }
_typeInfoList[223] = { parentId = 680, typeId = 1352, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[224] = { parentId = 1350, typeId = 1354, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 1350, typeId = 1362, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[226] = { parentId = 104, typeId = 1370, baseId = 536, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1374, 1382, 1384, 1386}, }
_typeInfoList[227] = { parentId = 680, typeId = 1372, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[228] = { parentId = 1370, typeId = 1374, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[229] = { parentId = 1370, typeId = 1382, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[230] = { parentId = 1370, typeId = 1384, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[231] = { parentId = 1370, typeId = 1386, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[232] = { parentId = 104, typeId = 1392, baseId = 536, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1396, 1404, 1406}, }
_typeInfoList[233] = { parentId = 680, typeId = 1394, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[234] = { parentId = 1392, typeId = 1396, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[235] = { parentId = 1392, typeId = 1404, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[236] = { parentId = 1392, typeId = 1406, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[237] = { parentId = 104, typeId = 1416, baseId = 536, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1420, 1428, 1430, 1432, 1434}, }
_typeInfoList[238] = { parentId = 680, typeId = 1418, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[239] = { parentId = 1416, typeId = 1420, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[240] = { parentId = 1416, typeId = 1428, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 1416, typeId = 1430, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[242] = { parentId = 1416, typeId = 1432, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[243] = { parentId = 1416, typeId = 1434, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[244] = { parentId = 104, typeId = 1440, baseId = 536, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1444, 1452, 1454}, }
_typeInfoList[245] = { parentId = 680, typeId = 1442, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[246] = { parentId = 1440, typeId = 1444, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[247] = { parentId = 1440, typeId = 1452, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[248] = { parentId = 1440, typeId = 1454, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[249] = { parentId = 104, typeId = 1462, baseId = 536, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1466, 1474, 1476, 1478}, }
_typeInfoList[250] = { parentId = 680, typeId = 1464, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 1462, typeId = 1466, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[252] = { parentId = 1462, typeId = 1474, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[253] = { parentId = 1462, typeId = 1476, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[254] = { parentId = 1462, typeId = 1478, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[255] = { parentId = 104, typeId = 1486, baseId = 536, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1490, 1498, 1500, 1502}, }
_typeInfoList[256] = { parentId = 680, typeId = 1488, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[257] = { parentId = 1486, typeId = 1490, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[258] = { parentId = 1486, typeId = 1498, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[259] = { parentId = 1486, typeId = 1500, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[260] = { parentId = 1486, typeId = 1502, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[261] = { parentId = 104, typeId = 1510, baseId = 536, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1514, 1522, 1524, 1526}, }
_typeInfoList[262] = { parentId = 680, typeId = 1512, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[263] = { parentId = 1510, typeId = 1514, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[264] = { parentId = 1510, typeId = 1522, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[265] = { parentId = 1510, typeId = 1524, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[266] = { parentId = 1510, typeId = 1526, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[267] = { parentId = 104, typeId = 1532, baseId = 536, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1536, 1544, 1546}, }
_typeInfoList[268] = { parentId = 680, typeId = 1534, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[269] = { parentId = 1532, typeId = 1536, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[270] = { parentId = 1532, typeId = 1544, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[271] = { parentId = 1532, typeId = 1546, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[272] = { parentId = 104, typeId = 1552, baseId = 536, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1556, 1564, 1566}, }
_typeInfoList[273] = { parentId = 680, typeId = 1554, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[274] = { parentId = 1552, typeId = 1556, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[275] = { parentId = 1552, typeId = 1564, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[276] = { parentId = 1552, typeId = 1566, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[277] = { parentId = 104, typeId = 1572, baseId = 536, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1578, 1588, 1592}, }
_typeInfoList[278] = { parentId = 680, typeId = 1574, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[279] = { parentId = 1572, typeId = 1578, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[280] = { parentId = 1572, typeId = 1588, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[281] = { parentId = 1, typeId = 1590, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[282] = { parentId = 1572, typeId = 1592, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1590}, children = {}, }
_typeInfoList[283] = { parentId = 104, typeId = 1598, baseId = 536, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1604, 1614, 1618, 2390}, }
_typeInfoList[284] = { parentId = 680, typeId = 1600, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[285] = { parentId = 1598, typeId = 1604, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[286] = { parentId = 1598, typeId = 1614, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 1, typeId = 1616, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[288] = { parentId = 1598, typeId = 1618, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1616}, children = {}, }
_typeInfoList[289] = { parentId = 104, typeId = 1624, baseId = 536, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1628, 1636, 1638}, }
_typeInfoList[290] = { parentId = 680, typeId = 1626, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[291] = { parentId = 1624, typeId = 1628, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[292] = { parentId = 1624, typeId = 1636, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[293] = { parentId = 1624, typeId = 1638, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[294] = { parentId = 104, typeId = 1646, baseId = 536, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1650, 1658, 1660, 1662, 2380}, }
_typeInfoList[295] = { parentId = 680, typeId = 1648, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[296] = { parentId = 1646, typeId = 1650, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[297] = { parentId = 1646, typeId = 1658, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[298] = { parentId = 1646, typeId = 1660, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[299] = { parentId = 1646, typeId = 1662, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[300] = { parentId = 104, typeId = 1672, baseId = 536, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1676, 1684, 1686, 1688, 1690}, }
_typeInfoList[301] = { parentId = 680, typeId = 1674, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[302] = { parentId = 1672, typeId = 1676, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[303] = { parentId = 1672, typeId = 1684, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[304] = { parentId = 1672, typeId = 1686, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[305] = { parentId = 1672, typeId = 1688, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[306] = { parentId = 1672, typeId = 1690, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[307] = { parentId = 104, typeId = 1692, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1694, 1696, 1698}, }
_typeInfoList[308] = { parentId = 1692, typeId = 1694, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[309] = { parentId = 1692, typeId = 1696, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {992}, children = {}, }
_typeInfoList[310] = { parentId = 1692, typeId = 1698, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[311] = { parentId = 104, typeId = 1714, baseId = 536, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1722, 1734, 1736, 1738, 1742, 1744, 1748, 1750}, }
_typeInfoList[312] = { parentId = 680, typeId = 1716, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[313] = { parentId = 1714, typeId = 1722, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 1714, typeId = 1734, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[315] = { parentId = 1714, typeId = 1736, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[316] = { parentId = 1714, typeId = 1738, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[317] = { parentId = 1, typeId = 1740, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1692}, retTypeId = {}, children = {}, }
_typeInfoList[318] = { parentId = 1714, typeId = 1742, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1740}, children = {}, }
_typeInfoList[319] = { parentId = 1714, typeId = 1744, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[320] = { parentId = 1, typeId = 1746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[321] = { parentId = 1714, typeId = 1748, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1746}, children = {}, }
_typeInfoList[322] = { parentId = 1714, typeId = 1750, baseId = 1, txt = 'get_unwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1020}, children = {}, }
_typeInfoList[323] = { parentId = 104, typeId = 1752, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1760, 1762, 1766, 1768, 1770, 1772, 1776, 1780}, }
_typeInfoList[324] = { parentId = 1752, typeId = 1760, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[325] = { parentId = 1752, typeId = 1762, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[326] = { parentId = 1, typeId = 1764, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[327] = { parentId = 1752, typeId = 1766, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1764}, children = {}, }
_typeInfoList[328] = { parentId = 1752, typeId = 1768, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[329] = { parentId = 1752, typeId = 1770, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[330] = { parentId = 1752, typeId = 1772, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[331] = { parentId = 1, typeId = 1774, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[332] = { parentId = 1752, typeId = 1776, baseId = 1, txt = 'get_retTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1774}, children = {}, }
_typeInfoList[333] = { parentId = 1, typeId = 1778, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[334] = { parentId = 1752, typeId = 1780, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1778}, children = {}, }
_typeInfoList[335] = { parentId = 104, typeId = 1786, baseId = 536, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1790, 1798, 1800}, }
_typeInfoList[336] = { parentId = 680, typeId = 1788, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[337] = { parentId = 1786, typeId = 1790, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 1786, typeId = 1798, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[339] = { parentId = 1786, typeId = 1800, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1752}, children = {}, }
_typeInfoList[340] = { parentId = 104, typeId = 1806, baseId = 536, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1810, 1818, 1820}, }
_typeInfoList[341] = { parentId = 680, typeId = 1808, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[342] = { parentId = 1806, typeId = 1810, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[343] = { parentId = 1806, typeId = 1818, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[344] = { parentId = 1806, typeId = 1820, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1752}, children = {}, }
_typeInfoList[345] = { parentId = 104, typeId = 1826, baseId = 536, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1830, 1838, 1840}, }
_typeInfoList[346] = { parentId = 680, typeId = 1828, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[347] = { parentId = 1826, typeId = 1830, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[348] = { parentId = 1826, typeId = 1838, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[349] = { parentId = 1826, typeId = 1840, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1752}, children = {}, }
_typeInfoList[350] = { parentId = 104, typeId = 1848, baseId = 536, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1852, 1860, 1862, 1864}, }
_typeInfoList[351] = { parentId = 680, typeId = 1850, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[352] = { parentId = 1848, typeId = 1852, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[353] = { parentId = 1848, typeId = 1860, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[354] = { parentId = 1848, typeId = 1862, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[355] = { parentId = 1848, typeId = 1864, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[356] = { parentId = 104, typeId = 1880, baseId = 536, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1884, 1892, 1894, 1896, 1898, 1900, 1902, 1904}, }
_typeInfoList[357] = { parentId = 680, typeId = 1882, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[358] = { parentId = 1880, typeId = 1884, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[359] = { parentId = 1880, typeId = 1892, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[360] = { parentId = 1880, typeId = 1894, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[361] = { parentId = 1880, typeId = 1896, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {992}, children = {}, }
_typeInfoList[362] = { parentId = 1880, typeId = 1898, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[363] = { parentId = 1880, typeId = 1900, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[364] = { parentId = 1880, typeId = 1902, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[365] = { parentId = 1880, typeId = 1904, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[366] = { parentId = 104, typeId = 1912, baseId = 536, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1916, 1924, 1926, 1928}, }
_typeInfoList[367] = { parentId = 680, typeId = 1914, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[368] = { parentId = 1912, typeId = 1916, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 1912, typeId = 1924, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[370] = { parentId = 1912, typeId = 1926, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[371] = { parentId = 1912, typeId = 1928, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {992}, children = {}, }
_typeInfoList[372] = { parentId = 104, typeId = 1930, baseId = 536, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {1934, 1942}, }
_typeInfoList[373] = { parentId = 680, typeId = 1932, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[374] = { parentId = 1930, typeId = 1934, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[375] = { parentId = 1930, typeId = 1942, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[376] = { parentId = 680, typeId = 1958, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[377] = { parentId = 538, typeId = 1966, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[378] = { parentId = 538, typeId = 1980, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[379] = { parentId = 538, typeId = 1982, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[380] = { parentId = 538, typeId = 1984, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[381] = { parentId = 1, typeId = 1986, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[382] = { parentId = 538, typeId = 1988, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1986}, children = {}, }
_typeInfoList[383] = { parentId = 1, typeId = 1990, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1880}, retTypeId = {}, children = {}, }
_typeInfoList[384] = { parentId = 538, typeId = 1992, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1990}, children = {}, }
_typeInfoList[385] = { parentId = 538, typeId = 1994, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {640}, children = {}, }
_typeInfoList[386] = { parentId = 1, typeId = 1996, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[387] = { parentId = 538, typeId = 1998, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1996}, children = {}, }
_typeInfoList[388] = { parentId = 104, typeId = 2004, baseId = 536, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2008, 2016, 2018}, }
_typeInfoList[389] = { parentId = 680, typeId = 2006, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 2004, typeId = 2008, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[391] = { parentId = 2004, typeId = 2016, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 2004, typeId = 2018, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {706}, children = {}, }
_typeInfoList[393] = { parentId = 104, typeId = 2020, baseId = 536, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2024, 2032, 2262}, }
_typeInfoList[394] = { parentId = 680, typeId = 2022, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[395] = { parentId = 2020, typeId = 2024, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 2020, typeId = 2032, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[397] = { parentId = 104, typeId = 2040, baseId = 536, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2044, 2052, 2054, 2056, 2272}, }
_typeInfoList[398] = { parentId = 680, typeId = 2042, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[399] = { parentId = 2040, typeId = 2044, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[400] = { parentId = 2040, typeId = 2052, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[401] = { parentId = 2040, typeId = 2054, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[402] = { parentId = 2040, typeId = 2056, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[403] = { parentId = 104, typeId = 2064, baseId = 536, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2068, 2076, 2078, 2080, 2282}, }
_typeInfoList[404] = { parentId = 680, typeId = 2066, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[405] = { parentId = 2064, typeId = 2068, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[406] = { parentId = 2064, typeId = 2076, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[407] = { parentId = 2064, typeId = 2078, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[408] = { parentId = 2064, typeId = 2080, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[409] = { parentId = 104, typeId = 2088, baseId = 536, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2092, 2100, 2102, 2104, 2292}, }
_typeInfoList[410] = { parentId = 680, typeId = 2090, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[411] = { parentId = 2088, typeId = 2092, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[412] = { parentId = 2088, typeId = 2100, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[413] = { parentId = 2088, typeId = 2102, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[414] = { parentId = 2088, typeId = 2104, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[415] = { parentId = 104, typeId = 2110, baseId = 536, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2114, 2122, 2124, 2302}, }
_typeInfoList[416] = { parentId = 680, typeId = 2112, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[417] = { parentId = 2110, typeId = 2114, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[418] = { parentId = 2110, typeId = 2122, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 2110, typeId = 2124, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[420] = { parentId = 104, typeId = 2130, baseId = 536, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2134, 2142, 2144, 2316}, }
_typeInfoList[421] = { parentId = 680, typeId = 2132, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[422] = { parentId = 2130, typeId = 2134, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[423] = { parentId = 2130, typeId = 2142, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[424] = { parentId = 2130, typeId = 2144, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {1082}, children = {}, }
_typeInfoList[425] = { parentId = 104, typeId = 2146, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2148, 2150}, }
_typeInfoList[426] = { parentId = 2146, typeId = 2148, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[427] = { parentId = 2146, typeId = 2150, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[428] = { parentId = 104, typeId = 2158, baseId = 536, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2166, 2178, 2182, 2186, 2330}, }
_typeInfoList[429] = { parentId = 680, typeId = 2160, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[430] = { parentId = 2158, typeId = 2166, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[431] = { parentId = 2158, typeId = 2178, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[432] = { parentId = 1, typeId = 2180, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {536, 536}, retTypeId = {}, children = {}, }
_typeInfoList[433] = { parentId = 2158, typeId = 2182, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2180}, children = {}, }
_typeInfoList[434] = { parentId = 1, typeId = 2184, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2146}, retTypeId = {}, children = {}, }
_typeInfoList[435] = { parentId = 2158, typeId = 2186, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2184}, children = {}, }
_typeInfoList[436] = { parentId = 104, typeId = 2194, baseId = 536, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2200, 2210, 2212, 2216, 2344}, }
_typeInfoList[437] = { parentId = 680, typeId = 2196, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[438] = { parentId = 2194, typeId = 2200, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[439] = { parentId = 2194, typeId = 2210, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[440] = { parentId = 2194, typeId = 2212, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[441] = { parentId = 1, typeId = 2214, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {536}, retTypeId = {}, children = {}, }
_typeInfoList[442] = { parentId = 2194, typeId = 2216, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2214}, children = {}, }
_typeInfoList[443] = { parentId = 104, typeId = 2222, baseId = 536, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2226, 2234, 2236, 2358}, }
_typeInfoList[444] = { parentId = 680, typeId = 2224, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[445] = { parentId = 2222, typeId = 2226, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[446] = { parentId = 2222, typeId = 2234, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[447] = { parentId = 2222, typeId = 2236, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[448] = { parentId = 104, typeId = 2242, baseId = 536, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2246, 2254, 2256, 2368}, }
_typeInfoList[449] = { parentId = 680, typeId = 2244, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[450] = { parentId = 2242, typeId = 2246, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[451] = { parentId = 2242, typeId = 2254, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[452] = { parentId = 2242, typeId = 2256, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {366}, children = {}, }
_typeInfoList[453] = { parentId = 1, typeId = 2258, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[454] = { parentId = 1, typeId = 2260, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[455] = { parentId = 2020, typeId = 2262, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2258, 2260}, children = {}, }
_typeInfoList[456] = { parentId = 1, typeId = 2268, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[457] = { parentId = 1, typeId = 2270, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[458] = { parentId = 2040, typeId = 2272, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2268, 2270}, children = {}, }
_typeInfoList[459] = { parentId = 1, typeId = 2278, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[460] = { parentId = 1, typeId = 2280, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[461] = { parentId = 2064, typeId = 2282, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2278, 2280}, children = {}, }
_typeInfoList[462] = { parentId = 1, typeId = 2288, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[463] = { parentId = 1, typeId = 2290, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[464] = { parentId = 2088, typeId = 2292, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2288, 2290}, children = {}, }
_typeInfoList[465] = { parentId = 1, typeId = 2298, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[466] = { parentId = 1, typeId = 2300, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[467] = { parentId = 2110, typeId = 2302, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2298, 2300}, children = {}, }
_typeInfoList[468] = { parentId = 1, typeId = 2312, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[469] = { parentId = 1, typeId = 2314, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[470] = { parentId = 2130, typeId = 2316, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2312, 2314}, children = {}, }
_typeInfoList[471] = { parentId = 1, typeId = 2326, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[472] = { parentId = 1, typeId = 2328, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[473] = { parentId = 2158, typeId = 2330, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2326, 2328}, children = {}, }
_typeInfoList[474] = { parentId = 1, typeId = 2340, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[475] = { parentId = 1, typeId = 2342, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 2194, typeId = 2344, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2340, 2342}, children = {}, }
_typeInfoList[477] = { parentId = 1, typeId = 2354, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[478] = { parentId = 1, typeId = 2356, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 2222, typeId = 2358, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2354, 2356}, children = {}, }
_typeInfoList[480] = { parentId = 1, typeId = 2364, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[481] = { parentId = 1, typeId = 2366, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[482] = { parentId = 2242, typeId = 2368, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2364, 2366}, children = {}, }
_typeInfoList[483] = { parentId = 1, typeId = 2376, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[484] = { parentId = 1, typeId = 2378, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[485] = { parentId = 1646, typeId = 2380, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2376, 2378}, children = {}, }
_typeInfoList[486] = { parentId = 1, typeId = 2386, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, retTypeId = {}, children = {}, }
_typeInfoList[487] = { parentId = 1, typeId = 2388, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {516}, retTypeId = {}, children = {}, }
_typeInfoList[488] = { parentId = 1598, typeId = 2390, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2386, 2388}, children = {}, }
_typeInfoList[489] = { parentId = 1, typeId = 2396, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, retTypeId = {}, children = {}, }
_typeInfoList[490] = { parentId = 704, typeId = 2398, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2396}, children = {}, }
_typeInfoList[491] = { parentId = 104, typeId = 2734, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, retTypeId = {}, children = {2736, 2738}, }
_typeInfoList[492] = { parentId = 2734, typeId = 2736, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[493] = { parentId = 2734, typeId = 2738, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {516}, children = {}, }
_typeInfoList[494] = { parentId = 730, typeId = 2740, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, retTypeId = {2734}, children = {}, }
----- meta -----
return moduleObj
