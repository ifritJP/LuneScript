--lune/base/TransUnit.lns
local moduleObj = {}
local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local rootTypeId = 1

moduleObj.rootTypeId = rootTypeId

local typeIdSeed = rootTypeId + 1

-- none

local typeInfoKind = {}

moduleObj.typeInfoKind = typeInfoKind

local sym2builtInTypeMap = {}

local builtInTypeIdSet = {}

local TypeInfoKindRoot = 0

moduleObj.TypeInfoKindRoot = TypeInfoKindRoot

local TypeInfoKindMacro = 1

moduleObj.TypeInfoKindMacro = TypeInfoKindMacro

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

local function isBuiltin( typeId )
  return builtInTypeIdSet[typeId]
end
moduleObj.isBuiltin = isBuiltin
local OutStream = {}
moduleObj.OutStream = OutStream
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

local dummyList = {}

-- none

local Scope = {}
moduleObj.Scope = Scope
function Scope.new( parent, classFlag, inheritList )
  local obj = {}
  setmetatable( obj, { __index = Scope } )
  if obj.__init then obj:__init( parent, classFlag, inheritList ); end
return obj
end
function Scope:__init(parent, classFlag, inheritList) 
  self.parent = parent or self
  self.symbol2TypeInfoMap = {}
  self.className2ScopeMap = {}
  self.inheritList = inheritList
  self.classFlag = classFlag
  self.hasParentFlag = parent ~= nil
end
function Scope:set_ownerTypeInfo( owner )
  self.ownerTypeInfo = owner
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
  
  if not scope then
    if self.hasParentFlag then
      scope = self.parent:getClassScope( name )
    end
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
  if self.hasParentFlag then
    return self.parent:getTypeInfo( name )
  end
  return sym2builtInTypeMap[name]
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
function Scope:get_ownerTypeInfo()
  return self.ownerTypeInfo
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

local rootScope = Scope.new(nil, false, {})

local TypeInfo = {}
moduleObj.TypeInfo = TypeInfo
function TypeInfo.new( scope )
  local obj = {}
  setmetatable( obj, { __index = TypeInfo } )
  if obj.__init then obj:__init( scope ); end
return obj
end
function TypeInfo:__init(scope) 
  self.scope = scope
  do
    local _exp = scope
    if _exp then
    
        _exp:set_ownerTypeInfo( self )
      end
  end
  
end
function TypeInfo:getParentId(  )
  return rootTypeId
end
function TypeInfo:get_baseId(  )
  return rootTypeId
end
-- none
function TypeInfo:getTxt(  )
  return ""
end
function TypeInfo:serialize( stream )
  return 
end
function TypeInfo:equals( typeInfo, depth )
  return false
end
function TypeInfo:get_externalFlag(  )
  return false
end
function TypeInfo:get_itemTypeInfoList(  )
  return dummyList
end
function TypeInfo:get_argTypeInfoList(  )
  return dummyList
end
function TypeInfo:get_retTypeInfoList(  )
  return dummyList
end
function TypeInfo:get_parentInfo(  )
  return self
end
function TypeInfo:get_rawTxt(  )
  return ""
end
function TypeInfo:get_typeId(  )
  return rootTypeId
end
function TypeInfo:get_kind(  )
  return TypeInfoKindRoot
end
function TypeInfo:get_staticFlag(  )
  return false
end
function TypeInfo:get_accessMode(  )
  return "pri"
end
function TypeInfo:get_autoFlag(  )
  return false
end
function TypeInfo:get_orgTypeInfo(  )
  return self
end
function TypeInfo:get_baseTypeInfo(  )
  return self
end
function TypeInfo:get_nilable(  )
  return false
end
function TypeInfo:get_nilableTypeInfo(  )
  return self
end
function TypeInfo:get_children(  )
  return dummyList
end
function TypeInfo:get_children(  )
  return dummyList
end
function TypeInfo:get_scope()
  return self.scope
end

local function dumpScope( scope, prefix )
  do
    local _exp = scope
    if _exp then
    
        if #prefix > 20 then
          error( "illegal" )
        end
        do
          local __sorted = {}
          local __map = _exp:get_symbol2TypeInfoMap()
          for __key in pairs( __map ) do
            table.insert( __sorted, __key )
          end
          table.sort( __sorted )
          for __index, symbol in ipairs( __sorted ) do
            typeInfo = __map[ symbol ]
            do
              Util.errorLog( string.format( "scope: %s, %s, %s", prefix, _exp, symbol) )
              do
                local _exp = typeInfo:get_scope()
                if _exp then
                
                    dumpScope( _exp, prefix .. "  " )
                  end
              end
              
            end
          end
        end
        
      end
  end
  
end

local rootTypeInfo = TypeInfo.new(rootScope)

-- none

-- none

local NormalTypeInfo = {}
setmetatable( NormalTypeInfo, { __index = TypeInfo } )
moduleObj.NormalTypeInfo = NormalTypeInfo
function NormalTypeInfo.new( scope, baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = NormalTypeInfo } )
  if obj.__init then obj:__init( scope, baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList ); end
return obj
end
function NormalTypeInfo:__init(scope, baseTypeInfo, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList) 
  TypeInfo.__init( self, scope)
  
  self.baseTypeInfo = baseTypeInfo or rootTypeInfo
  self.autoFlag = autoFlag
  self.externalFlag = externalFlag
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.rawTxt = txt
  self.kind = kind
  self.itemTypeInfoList = itemTypeInfoList or {}
  self.argTypeInfoList = argTypeInfoList or {}
  self.retTypeInfoList = retTypeInfoList or {}
  self.orgTypeInfo = orgTypeInfo or rootTypeInfo
  self.parentInfo = parentInfo or rootTypeInfo
  self.children = {}
  self.typeId = typeId
  if kind == TypeInfoKindRoot then
    self.nilable = false
  elseif txt == "nil" or txt == "stem" then
    self.nilable = true
    self.nilableTypeInfo = self
    self.orgTypeInfo = self
  elseif not orgTypeInfo then
    if self.parentInfo ~= rootTypeInfo then
      table.insert( self.parentInfo:get_children(), self )
    end
    self.nilable = false
    do
      local _switchExp = (kind )
      if _switchExp == TypeInfoKindPrim or _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindArray or _switchExp == TypeInfoKindMap or _switchExp == TypeInfoKindClass then
        self.nilableTypeInfo = NormalTypeInfo.new(nil, baseTypeInfo, self, autoFlag, externalFlag, staticFlag, accessMode, "", parentInfo, typeId + 1, TypeInfoKindNilable, itemTypeInfoList, retTypeInfoList)
      end
    end
    
    typeIdSeed = typeIdSeed + 1
  else 
    self.nilable = true
    self.nilableTypeInfo = rootTypeInfo
  end
end
function NormalTypeInfo:getParentId(  )
  return self.parentInfo and self.parentInfo:get_typeId() or rootTypeId
end
function NormalTypeInfo:get_baseId(  )
  return self.baseTypeInfo and self.baseTypeInfo:get_typeId() or rootTypeId
end
function NormalTypeInfo:getTxt(  )
  if self.nilable and (self.nilableTypeInfo ~= self.orgTypeInfo ) then
    return (self.orgTypeInfo or _luneScript.error( 'unwrap val is nil' ) ):getTxt(  ) .. "!"
  end
  if self.kind == TypeInfoKindArray then
    local _exp = self.itemTypeInfoList[1]
    
        if  not _exp then
          local __exp = _exp
          
          return "[@]"
        end
      
    return _exp:getTxt(  ) .. "[@]"
  end
  if self.kind == TypeInfoKindList then
    local _exp = self.itemTypeInfoList[1]
    
        if  not _exp then
          local __exp = _exp
          
          return "[]"
        end
      
    return _exp:getTxt(  ) .. "[]"
  end
  if self.itemTypeInfoList and #self.itemTypeInfoList > 0 then
    local txt = self.rawTxt .. "<"
    
    for index, typeInfo in pairs( self.itemTypeInfoList ) do
      if index ~= 1 then
        txt = txt .. ","
      end
      txt = txt .. typeInfo:getTxt(  )
    end
    return txt .. ">"
  end
  if self:get_rawTxt() then
    return self:get_rawTxt()
  end
  return ""
end
function NormalTypeInfo:serialize( stream )
  if self.typeId == rootTypeId then
    return nil
  end
  local parentId = self:getParentId(  )
  
  if self.nilable then
    stream:write( string.format( '{ parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', parentId, self.typeId, self.orgTypeInfo:get_typeId()) )
    return nil
  end
  local function serializeTypeInfoList( name, list, onlyPub )
    local work = name
    
    for __index, typeInfo in pairs( list ) do
      if not onlyPub or typeInfo:get_accessMode() == "pub" then
        if #work ~= #name then
          work = work .. ", "
        end
        work = string.format( "%s%d", work, typeInfo:get_typeId())
      end
    end
    return work .. "}, "
  end
  
  local txt = string.format( [==[{ parentId = %d, typeId = %d, baseId = %d, txt = '%s',
        staticFlag = %s, accessMode = '%s', kind = %d, ]==], parentId, self.typeId, self:get_baseId(  ), self.rawTxt, self.staticFlag, self.accessMode, self.kind)
  
  stream:write( txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. serializeTypeInfoList( "children = {", self.children, true ) .. "}\n" )
end
function NormalTypeInfo:equals( typeInfo, depth )
  if not typeInfo then
    return false
  end
  if not depth then
    depth = 1
  end
  if self.typeId == typeInfo:get_typeId() then
    return true
  end
  if self.kind ~= typeInfo:get_kind() or self.staticFlag ~= typeInfo:get_staticFlag() or self.accessMode ~= typeInfo:get_accessMode() or self.autoFlag ~= typeInfo:get_autoFlag() or self.nilable ~= typeInfo:get_nilable() then
    return false
  end
  if (not self.itemTypeInfoList and typeInfo:get_itemTypeInfoList() or self.itemTypeInfoList and not typeInfo:get_itemTypeInfoList() or not self.retTypeInfoList and typeInfo:get_retTypeInfoList() or self.retTypeInfoList and not typeInfo:get_retTypeInfoList() or self.orgTypeInfo ~= typeInfo:get_orgTypeInfo() ) then
    Util.errorLog( "%s, %s", self.itemTypeInfoList, typeInfo:get_itemTypeInfoList() )
    Util.errorLog( "%s, %s", self.retTypeInfoList, typeInfo:get_retTypeInfoList() )
    Util.errorLog( "%s, %s", self.orgTypeInfo, typeInfo:get_orgTypeInfo() )
    return false
  end
  if self.itemTypeInfoList then
    if #self.itemTypeInfoList ~= #typeInfo:get_itemTypeInfoList() then
      return false
    end
    for index, item in pairs( self.itemTypeInfoList ) do
      if not item:equals( typeInfo:get_itemTypeInfoList()[index], depth + 1 ) then
        return false
      end
    end
  end
  if self.retTypeInfoList then
    if #self.retTypeInfoList ~= #typeInfo:get_retTypeInfoList() then
      return false
    end
    for index, item in pairs( self.retTypeInfoList ) do
      if not item:equals( typeInfo:get_retTypeInfoList()[index], depth + 1 ) then
        return false
      end
    end
  end
  if self.orgTypeInfo and not self.orgTypeInfo:equals( typeInfo:get_orgTypeInfo(), depth + 1 ) then
    return false
  end
  return true
end
function NormalTypeInfo.cloneToPublic( typeInfo )
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(typeInfo:get_scope(), typeInfo:get_baseTypeInfo(), nil, typeInfo:get_autoFlag(), typeInfo:get_externalFlag(), typeInfo:get_staticFlag(), "pub", typeInfo:get_rawTxt(), typeInfo:get_parentInfo(), typeIdSeed, typeInfo:get_kind(), typeInfo:get_itemTypeInfoList(), typeInfo:get_argTypeInfoList(), typeInfo:get_retTypeInfoList())
end
function NormalTypeInfo.create( scope, baseInfo, parentInfo, staticFlag, kind, txt, itemTypeInfo, argTypeInfoList, retTypeInfoList )
  if kind == TypeInfoKindPrim then
    return sym2builtInTypeMap[txt] or _luneScript.error( 'unwrap val is nil' )
  end
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(scope, baseInfo, nil, false, true, staticFlag, "pub", txt, parentInfo, typeIdSeed, kind, itemTypeInfo, argTypeInfoList, retTypeInfoList)
  
  return info
end
function NormalTypeInfo:get_itemTypeInfoList()
  return self.itemTypeInfoList
end
function NormalTypeInfo:get_argTypeInfoList()
  return self.argTypeInfoList
end
function NormalTypeInfo:get_retTypeInfoList()
  return self.retTypeInfoList
end
function NormalTypeInfo:get_parentInfo()
  return self.parentInfo
end
function NormalTypeInfo:get_typeId()
  return self.typeId
end
function NormalTypeInfo:get_rawTxt()
  return self.rawTxt
end
function NormalTypeInfo:get_kind()
  return self.kind
end
function NormalTypeInfo:get_staticFlag()
  return self.staticFlag
end
function NormalTypeInfo:get_accessMode()
  return self.accessMode
end
function NormalTypeInfo:get_autoFlag()
  return self.autoFlag
end
function NormalTypeInfo:get_orgTypeInfo()
  return self.orgTypeInfo
end
function NormalTypeInfo:get_baseTypeInfo()
  return self.baseTypeInfo
end
function NormalTypeInfo:get_nilable()
  return self.nilable
end
function NormalTypeInfo:get_nilableTypeInfo()
  return self.nilableTypeInfo
end
function NormalTypeInfo:get_children()
  return self.children
end

local typeInfoRoot = rootTypeInfo

typeIdSeed = typeIdSeed + 1
function NormalTypeInfo.createBuiltin( idName, typeTxt, kind, typeDDD )
  local typeId = typeIdSeed + 1
  
  if kind == TypeInfoKindRoot then
    typeId = rootTypeId
  else 
    typeIdSeed = typeIdSeed + 1
  end
  local argTypeList = {}
  
  local retTypeList = {}
  
  if typeTxt == "form" then
    argTypeList = {typeDDD}
    retTypeList = {typeDDD}
  end
  local scope = nil
  
  do
    local _switchExp = kind
    if _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindClass or _switchExp == TypeInfoKindFunc or _switchExp == TypeInfoKindMethod or _switchExp == TypeInfoKindMacro then
      scope = Scope.new(rootScope, kind == TypeInfoKindClass, {})
    end
  end
  
  local info = NormalTypeInfo.new(scope, nil, nil, false, false, false, "pub", typeTxt, typeInfoRoot, typeId, kind, {}, argTypeList, retTypeList)
  
  if scope then
    rootScope:add( typeTxt, info )
  end
  typeInfoKind[idName] = info
  sym2builtInTypeMap[typeTxt] = info
  if info:get_nilableTypeInfo() ~= rootTypeInfo then
    sym2builtInTypeMap[typeTxt .. "!"] = info:get_nilableTypeInfo()
  end
  builtInTypeIdSet[info.typeId] = true
  return info
end

function NormalTypeInfo.createList( accessMode, parentInfo, itemTypeInfo )
  if not itemTypeInfo or #itemTypeInfo == 0 then
    error( string.format( "illegal list type: %s", itemTypeInfo) )
  end
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(nil, nil, nil, false, false, false, accessMode, "", typeInfoRoot, typeIdSeed, TypeInfoKindList, itemTypeInfo)
end

function NormalTypeInfo.createArray( accessMode, parentInfo, itemTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(nil, nil, nil, false, false, false, accessMode, "", typeInfoRoot, typeIdSeed, TypeInfoKindArray, itemTypeInfo)
end

function NormalTypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo )
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(nil, nil, nil, false, false, false, accessMode, "Map", typeInfoRoot, typeIdSeed, TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end

function NormalTypeInfo.createClass( scope, baseInfo, parentInfo, externalFlag, accessMode, className )
  local classTypeInfo = sym2builtInTypeMap[className]
  
  do
    local _exp = classTypeInfo
    if _exp then
    
        return _exp
      end
  end
  
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(scope, baseInfo, nil, false, externalFlag, false, accessMode, className, parentInfo, typeIdSeed, TypeInfoKindClass)
  
  return info
end

function NormalTypeInfo.createFunc( scope, kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(scope, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, typeIdSeed, kind, {}, argTypeList or {}, retTypeInfoList or {})
  
  return info
end

local builtinTypeNone = NormalTypeInfo.createBuiltin( "None", "", TypeInfoKindPrim )

moduleObj.builtinTypeNone = builtinTypeNone

local builtinTypeStem = NormalTypeInfo.createBuiltin( "Stem", "stem", TypeInfoKindPrim )

moduleObj.builtinTypeStem = builtinTypeStem

local builtinTypeNil = NormalTypeInfo.createBuiltin( "Nil", "nil", TypeInfoKindPrim )

moduleObj.builtinTypeNil = builtinTypeNil

local builtinTypeDDD = NormalTypeInfo.createBuiltin( "DDD", "...", TypeInfoKindPrim )

moduleObj.builtinTypeDDD = builtinTypeDDD

local builtinTypeBool = NormalTypeInfo.createBuiltin( "Bool", "bool", TypeInfoKindPrim )

moduleObj.builtinTypeBool = builtinTypeBool

local builtinTypeInt = NormalTypeInfo.createBuiltin( "Int", "int", TypeInfoKindPrim )

moduleObj.builtinTypeInt = builtinTypeInt

local builtinTypeReal = NormalTypeInfo.createBuiltin( "Real", "real", TypeInfoKindPrim )

moduleObj.builtinTypeReal = builtinTypeReal

local builtinTypeChar = NormalTypeInfo.createBuiltin( "char", "char", TypeInfoKindPrim )

moduleObj.builtinTypeChar = builtinTypeChar

local builtinTypeString = NormalTypeInfo.createBuiltin( "String", "str", TypeInfoKindClass )

moduleObj.builtinTypeString = builtinTypeString

local builtinTypeMap = NormalTypeInfo.createBuiltin( "Map", "Map", TypeInfoKindMap )

moduleObj.builtinTypeMap = builtinTypeMap

local builtinTypeList = NormalTypeInfo.createBuiltin( "List", "List", TypeInfoKindList )

moduleObj.builtinTypeList = builtinTypeList

local builtinTypeArray = NormalTypeInfo.createBuiltin( "Array", "Array", TypeInfoKindArray )

moduleObj.builtinTypeArray = builtinTypeArray

local builtinTypeForm = NormalTypeInfo.createBuiltin( "Form", "form", TypeInfoKindFunc, builtinTypeDDD )

moduleObj.builtinTypeForm = builtinTypeForm

local builtinTypeSymbol = NormalTypeInfo.createBuiltin( "Symbol", "sym", TypeInfoKindPrim )

moduleObj.builtinTypeSymbol = builtinTypeSymbol

local builtinTypeStat = NormalTypeInfo.createBuiltin( "Stat", "stat", TypeInfoKindPrim )

moduleObj.builtinTypeStat = builtinTypeStat


-- none

local typeInfoListInsert = typeInfoRoot

moduleObj.typeInfoListInsert = typeInfoListInsert

local typeInfoListRemove = typeInfoRoot

moduleObj.typeInfoListRemove = typeInfoListRemove

function NormalTypeInfo:isSettableFrom( other )
  if not other then
    return false
  end
  if self == builtinTypeStem or self == builtinTypeDDD then
    return true
  end
  if other == builtinTypeNil then
    if self.kind ~= TypeInfoKindNilable then
      return false
    end
    return true
  end
  if self.typeId == other:get_typeId() then
    return true
  end
  if self.kind ~= other:get_kind() then
    if self.kind == TypeInfoKindNilable then
      return (self.orgTypeInfo or _luneScript.error( 'unwrap val is nil' ) ):isSettableFrom( other )
    end
    return false
  end
  do
    local _switchExp = (self.kind )
    if _switchExp == TypeInfoKindPrim then
      if self == builtinTypeInt and other == builtinTypeChar or self == builtinTypeChar and other == builtinTypeInt then
        return true
      end
      return false
    elseif _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindArray then
      if other:get_itemTypeInfoList()[1] == builtinTypeNone then
        return true
      end
      if not (self:get_itemTypeInfoList()[1] or _luneScript.error( 'unwrap val is nil' ) ):isSettableFrom( other:get_itemTypeInfoList()[1] or _luneScript.error( 'unwrap val is nil' ) ) then
        return false
      end
      
      return true
    elseif _switchExp == TypeInfoKindMap then
      if other:get_itemTypeInfoList()[1] == builtinTypeNone and other:get_itemTypeInfoList()[2] == builtinTypeNone then
        return true
      end
      if not (self:get_itemTypeInfoList()[1] or _luneScript.error( 'unwrap val is nil' ) ):isSettableFrom( other:get_itemTypeInfoList()[1] or _luneScript.error( 'unwrap val is nil' ) ) then
        return false
      end
      
      if not (self:get_itemTypeInfoList()[2] or _luneScript.error( 'unwrap val is nil' ) ):isSettableFrom( other:get_itemTypeInfoList()[2] or _luneScript.error( 'unwrap val is nil' ) ) then
        return false
      end
      
      return true
    elseif _switchExp == TypeInfoKindClass then
      local baseTypeInfo = other:get_baseTypeInfo()
      
      while baseTypeInfo ~= rootTypeInfo do
        if self:get_typeId(  ) == baseTypeInfo:get_typeId(  ) then
          return true
        end
        baseTypeInfo = baseTypeInfo:get_baseTypeInfo()
      end
      -- none
      
      return false
    elseif _switchExp == TypeInfoKindFunc then
      if self == builtinTypeForm then
        return true
      end
      return false
    elseif _switchExp == TypeInfoKindMethod then
      return false
    else 
      return false
    end
  end
  
  return true
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

local Node = {}
moduleObj.Node = Node
function Node:get_expType(  )
  if not self.expTypeList then
    return builtinTypeNone
  end
  return self.expTypeList[1]
end
function Node:getLiteral(  )
  return {nil}, {builtinTypeNil}
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
  self.scope = rootScope
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
function TransUnit:getCurrentClass(  )
  local typeInfo = rootTypeInfo
  
  local scope = self.scope
  
  repeat 
    do
      local _exp = scope:get_ownerTypeInfo()
      if _exp then
      
          if _exp:get_kind() == TypeInfoKindClass then
            return _exp
          end
        end
    end
    
    scope = scope:get_parent()
  until scope == rootScope
  return typeInfo
end
function TransUnit:getCurrentNamespaceTypeInfo(  )
  local typeInfo = rootTypeInfo
  
  local scope = self.scope
  
  repeat 
    do
      local _exp = scope:get_ownerTypeInfo()
      if _exp then
      
          return _exp
        end
    end
    
    scope = scope:get_parent()
  until scope == rootScope
  Util.errorLog( string.format( "bbb: %s", typeInfo ) )
  return typeInfo
end
function TransUnit:pushClass( baseInfo, externalFlag, name, accessMode, defNamespace )
  local typeInfo = rootTypeInfo
  
  do
    local _exp = self.scope:getTypeInfoChild( name )
    if _exp then
    
        typeInfo = _exp
        self.scope = typeInfo:get_scope() or _luneScript.error( 'unwrap val is nil' )
      else
    
        local parentInfo = self:getCurrentNamespaceTypeInfo(  )
        
        local inheritList = {}
        
        do
          local _exp = baseInfo
          if _exp then
          
              inheritList = {_exp:get_scope(  ) or _luneScript.error( 'unwrap val is nil' )}
            end
        end
        
        local scope = self:pushScope( true, inheritList )
        
        typeInfo = NormalTypeInfo.createClass( scope, baseInfo, parentInfo, externalFlag, accessMode, name )
        scope:get_parent(  ):addClass( name, typeInfo, scope )
      end
  end
  
  local namespace = defNamespace
  
      if  not namespace then
        local _namespace = namespace
        
        namespace = NamespaceInfo.new(name, self.scope, typeInfo)
      end
    
  self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
  self.typeId2Scope[typeInfo:get_typeId(  )] = self.scope
  return typeInfo
end
function TransUnit:popClass(  )
  self:popScope(  )
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
  return nodeKind2NameMap[kind] or _luneScript.error( 'unwrap val is nil' )
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
function NoneNode.new( pos, builtinTypeList )
  local obj = {}
  setmetatable( obj, { __index = NoneNode } )
  if obj.__init then obj:__init( pos, builtinTypeList ); end
return obj
end
function NoneNode:__init(pos, builtinTypeList) 
  Node.__init( self, nodeKindNone, pos, builtinTypeList)
  
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
function ImportNode.new( pos, builtinTypeList, modulePath )
  local obj = {}
  setmetatable( obj, { __index = ImportNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, modulePath ); end
return obj
end
function ImportNode:__init(pos, builtinTypeList, modulePath) 
  Node.__init( self, nodeKindImport, pos, builtinTypeList)
  
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
function RootNode.new( pos, builtinTypeList, children, typeId2ClassMap )
  local obj = {}
  setmetatable( obj, { __index = RootNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, children, typeId2ClassMap ); end
return obj
end
function RootNode:__init(pos, builtinTypeList, children, typeId2ClassMap) 
  Node.__init( self, nodeKindRoot, pos, builtinTypeList)
  
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
function RefTypeNode.new( pos, builtinTypeList, name, refFlag, mutFlag, array )
  local obj = {}
  setmetatable( obj, { __index = RefTypeNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, name, refFlag, mutFlag, array ); end
return obj
end
function RefTypeNode:__init(pos, builtinTypeList, name, refFlag, mutFlag, array) 
  Node.__init( self, nodeKindRefType, pos, builtinTypeList)
  
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
function BlockNode.new( pos, builtinTypeList, blockKind, stmtList )
  local obj = {}
  setmetatable( obj, { __index = BlockNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, blockKind, stmtList ); end
return obj
end
function BlockNode:__init(pos, builtinTypeList, blockKind, stmtList) 
  Node.__init( self, nodeKindBlock, pos, builtinTypeList)
  
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
function IfNode.new( pos, builtinTypeList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = IfNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, stmtList ); end
return obj
end
function IfNode:__init(pos, builtinTypeList, stmtList) 
  Node.__init( self, nodeKindIf, pos, builtinTypeList)
  
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
function ExpListNode.new( pos, builtinTypeList, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpListNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, expList ); end
return obj
end
function ExpListNode:__init(pos, builtinTypeList, expList) 
  Node.__init( self, nodeKindExpList, pos, builtinTypeList)
  
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
function SwitchNode.new( pos, builtinTypeList, exp, caseList, default )
  local obj = {}
  setmetatable( obj, { __index = SwitchNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp, caseList, default ); end
return obj
end
function SwitchNode:__init(pos, builtinTypeList, exp, caseList, default) 
  Node.__init( self, nodeKindSwitch, pos, builtinTypeList)
  
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
function WhileNode.new( pos, builtinTypeList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = WhileNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp, block ); end
return obj
end
function WhileNode:__init(pos, builtinTypeList, exp, block) 
  Node.__init( self, nodeKindWhile, pos, builtinTypeList)
  
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
function RepeatNode.new( pos, builtinTypeList, block, exp )
  local obj = {}
  setmetatable( obj, { __index = RepeatNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, block, exp ); end
return obj
end
function RepeatNode:__init(pos, builtinTypeList, block, exp) 
  Node.__init( self, nodeKindRepeat, pos, builtinTypeList)
  
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
function ForNode.new( pos, builtinTypeList, block, val, init, to, delta )
  local obj = {}
  setmetatable( obj, { __index = ForNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, block, val, init, to, delta ); end
return obj
end
function ForNode:__init(pos, builtinTypeList, block, val, init, to, delta) 
  Node.__init( self, nodeKindFor, pos, builtinTypeList)
  
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
function ApplyNode.new( pos, builtinTypeList, varList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ApplyNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, varList, exp, block ); end
return obj
end
function ApplyNode:__init(pos, builtinTypeList, varList, exp, block) 
  Node.__init( self, nodeKindApply, pos, builtinTypeList)
  
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
function ForeachNode.new( pos, builtinTypeList, val, key, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ForeachNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, val, key, exp, block ); end
return obj
end
function ForeachNode:__init(pos, builtinTypeList, val, key, exp, block) 
  Node.__init( self, nodeKindForeach, pos, builtinTypeList)
  
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
function ForsortNode.new( pos, builtinTypeList, val, key, exp, block, sort )
  local obj = {}
  setmetatable( obj, { __index = ForsortNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, val, key, exp, block, sort ); end
return obj
end
function ForsortNode:__init(pos, builtinTypeList, val, key, exp, block, sort) 
  Node.__init( self, nodeKindForsort, pos, builtinTypeList)
  
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
function ReturnNode.new( pos, builtinTypeList, expList )
  local obj = {}
  setmetatable( obj, { __index = ReturnNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, expList ); end
return obj
end
function ReturnNode:__init(pos, builtinTypeList, expList) 
  Node.__init( self, nodeKindReturn, pos, builtinTypeList)
  
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
function BreakNode.new( pos, builtinTypeList )
  local obj = {}
  setmetatable( obj, { __index = BreakNode } )
  if obj.__init then obj:__init( pos, builtinTypeList ); end
return obj
end
function BreakNode:__init(pos, builtinTypeList) 
  Node.__init( self, nodeKindBreak, pos, builtinTypeList)
  
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
function ExpNewNode.new( pos, builtinTypeList, symbol, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpNewNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, symbol, argList ); end
return obj
end
function ExpNewNode:__init(pos, builtinTypeList, symbol, argList) 
  Node.__init( self, nodeKindExpNew, pos, builtinTypeList)
  
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

function Filter:processExpUnwrap( node, ... )
end

-- none

-- none

local nodeKindExpUnwrap = regKind( [[ExpUnwrap]] )

local ExpUnwrapNode = {}
setmetatable( ExpUnwrapNode, { __index = Node } )
moduleObj.ExpUnwrapNode = ExpUnwrapNode
function ExpUnwrapNode:processFilter( filter, ... )
  local argList = {...}
  
  filter:processExpUnwrap( self, table.unpack( argList ) )
end
function ExpUnwrapNode.new( pos, builtinTypeList, exp, default )
  local obj = {}
  setmetatable( obj, { __index = ExpUnwrapNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp, default ); end
return obj
end
function ExpUnwrapNode:__init(pos, builtinTypeList, exp, default) 
  Node.__init( self, nodeKindExpUnwrap, pos, builtinTypeList)
  
  -- none
  
  self.exp = exp
  self.default = default
  -- none
  
end
function ExpUnwrapNode:get_exp()
  return self.exp
end
function ExpUnwrapNode:get_default()
  return self.default
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
function ExpRefNode.new( pos, builtinTypeList, token )
  local obj = {}
  setmetatable( obj, { __index = ExpRefNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token ); end
return obj
end
function ExpRefNode:__init(pos, builtinTypeList, token) 
  Node.__init( self, nodeKindExpRef, pos, builtinTypeList)
  
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
function ExpOp2Node.new( pos, builtinTypeList, op, exp1, exp2 )
  local obj = {}
  setmetatable( obj, { __index = ExpOp2Node } )
  if obj.__init then obj:__init( pos, builtinTypeList, op, exp1, exp2 ); end
return obj
end
function ExpOp2Node:__init(pos, builtinTypeList, op, exp1, exp2) 
  Node.__init( self, nodeKindExpOp2, pos, builtinTypeList)
  
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

function Filter:processUnwrapSet( node, ... )
end

-- none

-- none

local nodeKindUnwrapSet = regKind( [[UnwrapSet]] )

local UnwrapSetNode = {}
setmetatable( UnwrapSetNode, { __index = Node } )
moduleObj.UnwrapSetNode = UnwrapSetNode
function UnwrapSetNode:processFilter( filter, ... )
  local argList = {...}
  
  filter:processUnwrapSet( self, table.unpack( argList ) )
end
function UnwrapSetNode.new( pos, builtinTypeList, dstExpList, srcExpList, unwrapBlock )
  local obj = {}
  setmetatable( obj, { __index = UnwrapSetNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, dstExpList, srcExpList, unwrapBlock ); end
return obj
end
function UnwrapSetNode:__init(pos, builtinTypeList, dstExpList, srcExpList, unwrapBlock) 
  Node.__init( self, nodeKindUnwrapSet, pos, builtinTypeList)
  
  -- none
  
  self.dstExpList = dstExpList
  self.srcExpList = srcExpList
  self.unwrapBlock = unwrapBlock
  -- none
  
end
function UnwrapSetNode:get_dstExpList()
  return self.dstExpList
end
function UnwrapSetNode:get_srcExpList()
  return self.srcExpList
end
function UnwrapSetNode:get_unwrapBlock()
  return self.unwrapBlock
end


-- none

function Filter:processIfUnwrap( node, ... )
end

-- none

-- none

local nodeKindIfUnwrap = regKind( [[IfUnwrap]] )

local IfUnwrapNode = {}
setmetatable( IfUnwrapNode, { __index = Node } )
moduleObj.IfUnwrapNode = IfUnwrapNode
function IfUnwrapNode:processFilter( filter, ... )
  local argList = {...}
  
  filter:processIfUnwrap( self, table.unpack( argList ) )
end
function IfUnwrapNode.new( pos, builtinTypeList, exp, block, nilBlock )
  local obj = {}
  setmetatable( obj, { __index = IfUnwrapNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp, block, nilBlock ); end
return obj
end
function IfUnwrapNode:__init(pos, builtinTypeList, exp, block, nilBlock) 
  Node.__init( self, nodeKindIfUnwrap, pos, builtinTypeList)
  
  -- none
  
  self.exp = exp
  self.block = block
  self.nilBlock = nilBlock
  -- none
  
end
function IfUnwrapNode:get_exp()
  return self.exp
end
function IfUnwrapNode:get_block()
  return self.block
end
function IfUnwrapNode:get_nilBlock()
  return self.nilBlock
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
function ExpCastNode.new( pos, builtinTypeList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpCastNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp ); end
return obj
end
function ExpCastNode:__init(pos, builtinTypeList, exp) 
  Node.__init( self, nodeKindExpCast, pos, builtinTypeList)
  
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
function ExpOp1Node.new( pos, builtinTypeList, op, macroMode, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpOp1Node } )
  if obj.__init then obj:__init( pos, builtinTypeList, op, macroMode, exp ); end
return obj
end
function ExpOp1Node:__init(pos, builtinTypeList, op, macroMode, exp) 
  Node.__init( self, nodeKindExpOp1, pos, builtinTypeList)
  
  -- none
  
  self.op = op
  self.macroMode = macroMode
  self.exp = exp
  -- none
  
end
function ExpOp1Node:get_op()
  return self.op
end
function ExpOp1Node:get_macroMode()
  return self.macroMode
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
function ExpRefItemNode.new( pos, builtinTypeList, val, index )
  local obj = {}
  setmetatable( obj, { __index = ExpRefItemNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, val, index ); end
return obj
end
function ExpRefItemNode:__init(pos, builtinTypeList, val, index) 
  Node.__init( self, nodeKindExpRefItem, pos, builtinTypeList)
  
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
function ExpCallNode.new( pos, builtinTypeList, func, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, func, argList ); end
return obj
end
function ExpCallNode:__init(pos, builtinTypeList, func, argList) 
  Node.__init( self, nodeKindExpCall, pos, builtinTypeList)
  
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
function ExpDDDNode.new( pos, builtinTypeList, token )
  local obj = {}
  setmetatable( obj, { __index = ExpDDDNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token ); end
return obj
end
function ExpDDDNode:__init(pos, builtinTypeList, token) 
  Node.__init( self, nodeKindExpDDD, pos, builtinTypeList)
  
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
function ExpParenNode.new( pos, builtinTypeList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpParenNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp ); end
return obj
end
function ExpParenNode:__init(pos, builtinTypeList, exp) 
  Node.__init( self, nodeKindExpParen, pos, builtinTypeList)
  
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
function ExpMacroExpNode.new( pos, builtinTypeList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroExpNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, stmtList ); end
return obj
end
function ExpMacroExpNode:__init(pos, builtinTypeList, stmtList) 
  Node.__init( self, nodeKindExpMacroExp, pos, builtinTypeList)
  
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
function ExpMacroStatNode.new( pos, builtinTypeList, expStrList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroStatNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, expStrList ); end
return obj
end
function ExpMacroStatNode:__init(pos, builtinTypeList, expStrList) 
  Node.__init( self, nodeKindExpMacroStat, pos, builtinTypeList)
  
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
function StmtExpNode.new( pos, builtinTypeList, exp )
  local obj = {}
  setmetatable( obj, { __index = StmtExpNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, exp ); end
return obj
end
function StmtExpNode:__init(pos, builtinTypeList, exp) 
  Node.__init( self, nodeKindStmtExp, pos, builtinTypeList)
  
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
function RefFieldNode.new( pos, builtinTypeList, field, prefix )
  local obj = {}
  setmetatable( obj, { __index = RefFieldNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, field, prefix ); end
return obj
end
function RefFieldNode:__init(pos, builtinTypeList, field, prefix) 
  Node.__init( self, nodeKindRefField, pos, builtinTypeList)
  
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
function GetFieldNode.new( pos, builtinTypeList, field, prefix, getterTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = GetFieldNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, field, prefix, getterTypeInfo ); end
return obj
end
function GetFieldNode:__init(pos, builtinTypeList, field, prefix, getterTypeInfo) 
  Node.__init( self, nodeKindGetField, pos, builtinTypeList)
  
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
function DeclVarNode.new( pos, builtinTypeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )
  local obj = {}
  setmetatable( obj, { __index = DeclVarNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock ); end
return obj
end
function DeclVarNode:__init(pos, builtinTypeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock) 
  Node.__init( self, nodeKindDeclVar, pos, builtinTypeList)
  
  -- none
  
  self.mode = mode
  self.accessMode = accessMode
  self.staticFlag = staticFlag
  self.varList = varList
  self.expList = expList
  self.typeInfoList = typeInfoList
  self.unwrapFlag = unwrapFlag
  self.unwrapBlock = unwrapBlock
  self.thenBlock = thenBlock
  self.syncVarList = syncVarList
  self.syncBlock = syncBlock
  -- none
  
end
function DeclVarNode:get_mode()
  return self.mode
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
function DeclVarNode:get_unwrapFlag()
  return self.unwrapFlag
end
function DeclVarNode:get_unwrapBlock()
  return self.unwrapBlock
end
function DeclVarNode:get_thenBlock()
  return self.thenBlock
end
function DeclVarNode:get_syncVarList()
  return self.syncVarList
end
function DeclVarNode:get_syncBlock()
  return self.syncBlock
end


local DeclFuncInfo = {}
moduleObj.DeclFuncInfo = DeclFuncInfo
function DeclFuncInfo.new( className, name, argList, staticFlag, accessMode, body, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = DeclFuncInfo } )
  if obj.__init then
    obj:__init( className, name, argList, staticFlag, accessMode, body, retTypeInfoList )
  end        
  return obj 
 end         
function DeclFuncInfo:__init( className, name, argList, staticFlag, accessMode, body, retTypeInfoList ) 
            
self.className = className
  self.name = name
  self.argList = argList
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.body = body
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
function DeclFuncNode.new( pos, builtinTypeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclFuncNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, declInfo ); end
return obj
end
function DeclFuncNode:__init(pos, builtinTypeList, declInfo) 
  Node.__init( self, nodeKindDeclFunc, pos, builtinTypeList)
  
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
function DeclMethodNode.new( pos, builtinTypeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMethodNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, declInfo ); end
return obj
end
function DeclMethodNode:__init(pos, builtinTypeList, declInfo) 
  Node.__init( self, nodeKindDeclMethod, pos, builtinTypeList)
  
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
function DeclConstrNode.new( pos, builtinTypeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclConstrNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, declInfo ); end
return obj
end
function DeclConstrNode:__init(pos, builtinTypeList, declInfo) 
  Node.__init( self, nodeKindDeclConstr, pos, builtinTypeList)
  
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
function ExpCallSuperNode.new( pos, builtinTypeList, superType, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallSuperNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, superType, expList ); end
return obj
end
function ExpCallSuperNode:__init(pos, builtinTypeList, superType, expList) 
  Node.__init( self, nodeKindExpCallSuper, pos, builtinTypeList)
  
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
function DeclMemberNode.new( pos, builtinTypeList, name, refType, staticFlag, accessMode, getterMode, setterMode )
  local obj = {}
  setmetatable( obj, { __index = DeclMemberNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, name, refType, staticFlag, accessMode, getterMode, setterMode ); end
return obj
end
function DeclMemberNode:__init(pos, builtinTypeList, name, refType, staticFlag, accessMode, getterMode, setterMode) 
  Node.__init( self, nodeKindDeclMember, pos, builtinTypeList)
  
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
function DeclArgNode.new( pos, builtinTypeList, name, argType )
  local obj = {}
  setmetatable( obj, { __index = DeclArgNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, name, argType ); end
return obj
end
function DeclArgNode:__init(pos, builtinTypeList, name, argType) 
  Node.__init( self, nodeKindDeclArg, pos, builtinTypeList)
  
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
function DeclArgDDDNode.new( pos, builtinTypeList )
  local obj = {}
  setmetatable( obj, { __index = DeclArgDDDNode } )
  if obj.__init then obj:__init( pos, builtinTypeList ); end
return obj
end
function DeclArgDDDNode:__init(pos, builtinTypeList) 
  Node.__init( self, nodeKindDeclArgDDD, pos, builtinTypeList)
  
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
function DeclClassNode.new( pos, builtinTypeList, accessMode, name, fieldList, memberList, scope, outerMethodSet )
  local obj = {}
  setmetatable( obj, { __index = DeclClassNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, accessMode, name, fieldList, memberList, scope, outerMethodSet ); end
return obj
end
function DeclClassNode:__init(pos, builtinTypeList, accessMode, name, fieldList, memberList, scope, outerMethodSet) 
  Node.__init( self, nodeKindDeclClass, pos, builtinTypeList)
  
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
function DeclMacroNode.new( pos, builtinTypeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMacroNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, declInfo ); end
return obj
end
function DeclMacroNode:__init(pos, builtinTypeList, declInfo) 
  Node.__init( self, nodeKindDeclMacro, pos, builtinTypeList)
  
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
function LiteralNilNode.new( pos, builtinTypeList )
  local obj = {}
  setmetatable( obj, { __index = LiteralNilNode } )
  if obj.__init then obj:__init( pos, builtinTypeList ); end
return obj
end
function LiteralNilNode:__init(pos, builtinTypeList) 
  Node.__init( self, nodeKindLiteralNil, pos, builtinTypeList)
  
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
function LiteralCharNode.new( pos, builtinTypeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralCharNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token, num ); end
return obj
end
function LiteralCharNode:__init(pos, builtinTypeList, token, num) 
  Node.__init( self, nodeKindLiteralChar, pos, builtinTypeList)
  
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
function LiteralIntNode.new( pos, builtinTypeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralIntNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token, num ); end
return obj
end
function LiteralIntNode:__init(pos, builtinTypeList, token, num) 
  Node.__init( self, nodeKindLiteralInt, pos, builtinTypeList)
  
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
function LiteralRealNode.new( pos, builtinTypeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralRealNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token, num ); end
return obj
end
function LiteralRealNode:__init(pos, builtinTypeList, token, num) 
  Node.__init( self, nodeKindLiteralReal, pos, builtinTypeList)
  
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
function LiteralArrayNode.new( pos, builtinTypeList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralArrayNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, expList ); end
return obj
end
function LiteralArrayNode:__init(pos, builtinTypeList, expList) 
  Node.__init( self, nodeKindLiteralArray, pos, builtinTypeList)
  
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
function LiteralListNode.new( pos, builtinTypeList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralListNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, expList ); end
return obj
end
function LiteralListNode:__init(pos, builtinTypeList, expList) 
  Node.__init( self, nodeKindLiteralList, pos, builtinTypeList)
  
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
function LiteralMapNode.new( pos, builtinTypeList, map, pairList )
  local obj = {}
  setmetatable( obj, { __index = LiteralMapNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, map, pairList ); end
return obj
end
function LiteralMapNode:__init(pos, builtinTypeList, map, pairList) 
  Node.__init( self, nodeKindLiteralMap, pos, builtinTypeList)
  
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
function LiteralStringNode.new( pos, builtinTypeList, token, argList )
  local obj = {}
  setmetatable( obj, { __index = LiteralStringNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token, argList ); end
return obj
end
function LiteralStringNode:__init(pos, builtinTypeList, token, argList) 
  Node.__init( self, nodeKindLiteralString, pos, builtinTypeList)
  
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
function LiteralBoolNode.new( pos, builtinTypeList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralBoolNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token ); end
return obj
end
function LiteralBoolNode:__init(pos, builtinTypeList, token) 
  Node.__init( self, nodeKindLiteralBool, pos, builtinTypeList)
  
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
function LiteralSymbolNode.new( pos, builtinTypeList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralSymbolNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token ); end
return obj
end
function LiteralSymbolNode:__init(pos, builtinTypeList, token) 
  Node.__init( self, nodeKindLiteralSymbol, pos, builtinTypeList)
  
  -- none
  
  self.token = token
  -- none
  
end
function LiteralSymbolNode:get_token()
  return self.token
end


function LiteralNilNode:getLiteral(  )
  return {nil}, {builtinTypeNil}
end

function LiteralCharNode:getLiteral(  )
  return {self.num}, {builtinTypeChar}
end

function LiteralIntNode:getLiteral(  )
  return {self.num}, {builtinTypeInt}
end

function LiteralRealNode:getLiteral(  )
  return {self.num}, {builtinTypeReal}
end

function LiteralArrayNode:getLiteral(  )
  local array = {}
  
  do
    local _exp = self.expList
    if _exp then
    
        for __index, val in pairs( _exp:get_expList(  ) ) do
          local txt = val:getLiteral(  )[1]
          
          table.insert( array, txt )
        end
      end
  end
  
  return {array}, {self:get_expType(  )}
end

function LiteralListNode:getLiteral(  )
  local list = {}
  
  do
    local _exp = self.expList
    if _exp then
    
        for __index, val in pairs( _exp:get_expList(  ) ) do
          local item = val:getLiteral(  )[1]
          
          table.insert( list, item )
        end
      end
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
  local argList = self:get_argList()
  
  if #argList > 0 then
    local argTbl = {}
    
    for __index, argNode in pairs( argList ) do
      local arg = argNode:getLiteral(  )
      
      table.insert( argTbl, arg[1] )
    end
    return {string.format( txt, table.unpack( argTbl ) )}, {builtinTypeString}
  end
  return {txt}, {builtinTypeString}
end

function LiteralBoolNode:getLiteral(  )
  return {self.token.txt == "true"}, {builtinTypeBool}
end

function LiteralSymbolNode:getLiteral(  )
  return {{self.token.txt}}, {builtinTypeSymbol}
end

function RefFieldNode:getLiteral(  )
  local prefix = self.prefix:getLiteral(  )[1]
  
  table.insert( prefix, "." )
  table.insert( prefix, self.field.txt )
  return {prefix}, {builtinTypeSymbol}
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
local _TypeInfo = {}
function _TypeInfo.new( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children )
  local obj = {}
  setmetatable( obj, { __index = _TypeInfo } )
  if obj.__init then
    obj:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children )
  end        
  return obj 
 end         
function _TypeInfo:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children ) 
            
self.baseId = baseId
  self.itemTypeId = itemTypeId
  self.argTypeId = argTypeId
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

local builtinModuleName2Scope = {}

function TransUnit:registBuiltInScope(  )
  local builtInInfo = {[""] = {["type"] = {["arg"] = {"stem"}, ["ret"] = {"str"}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {}}, ["print"] = {["arg"] = {"..."}, ["ret"] = {}}, ["tonumber"] = {["arg"] = {"str"}, ["ret"] = {"real"}}, ["load"] = {["arg"] = {"str"}, ["ret"] = {"stem", "str"}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}, ["_fcall"] = {["arg"] = {"form", "..."}, ["ret"] = {""}}}, ["io"] = {["open"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}}, ["os"] = {["clock"] = {["arg"] = {}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {}}}, ["string"] = {["find"] = {["arg"] = {"str", "str", "int", "bool"}, ["ret"] = {"int", "int"}}, ["byte"] = {["arg"] = {"str", "int"}, ["ret"] = {"int"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"stem"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int"}, ["ret"] = {"str"}}}, ["str"] = {["find"] = {["methodFlag"] = {}, ["arg"] = {"str", "int", "bool"}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["arg"] = {"..."}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["arg"] = {"str"}, ["ret"] = {"stem"}}, ["gsub"] = {["methodFlag"] = {}, ["arg"] = {"str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["arg"] = {"int", "int"}, ["ret"] = {"str"}}}, ["table"] = {["unpack"] = {["arg"] = {"stem"}, ["ret"] = {"stem"}}}, ["List"] = {["insert"] = {["methodFlag"] = {}, ["arg"] = {"stem"}, ["ret"] = {""}}, ["remove"] = {["methodFlag"] = {}, ["arg"] = {"int!"}, ["ret"] = {""}}}, ["debug"] = {["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"stem"}}}, ["_luneScript"] = {["loadModule"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}}}
  
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
          parentInfo = self:pushClass( nil, true, name, "pri" )
          builtInTypeIdSet[parentInfo:get_typeId(  )] = true
        end
        if not parentInfo then
          error( "parentInfo is nil" )
        end
        if not builtinModuleName2Scope[name] then
          if name ~= "" and sym2builtInTypeMap[name] then
            builtinModuleName2Scope[name] = self.scope
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
                local argTypeList = {}
                
                for __index, argType in pairs( info["arg"] or _luneScript.error( 'unwrap val is nil' ) ) do
                  table.insert( argTypeList, self.rootScope:getTypeInfo( argType ) or _luneScript.error( 'unwrap val is nil' ) )
                end
                local retTypeList = {}
                
                for __index, retType in pairs( info["ret"] or _luneScript.error( 'unwrap val is nil' ) ) do
                  table.insert( retTypeList, sym2builtInTypeMap[retType] or _luneScript.error( 'unwrap val is nil' ) )
                end
                local methodFlag = info["methodFlag"]
                
                self:pushScope( false )
                local typeInfo = NormalTypeInfo.createFunc( self.scope, methodFlag and TypeInfoKindMethod or TypeInfoKindFunc, parentInfo, false, true, not methodFlag, "pub", funcName, argTypeList, retTypeList )
                
                self:popScope(  )
                builtInTypeIdSet[typeInfo:get_typeId(  )] = true
                self.scope:add( funcName, typeInfo )
                if methodFlag then
                  do
                    local _switchExp = (name )
                    if _switchExp == "List" then
                      do
                        local _switchExp = (funcName )
                        if _switchExp == "insert" then
                          typeInfoListInsert = typeInfo
                        elseif _switchExp == "remove" then
                          typeInfoListRemove = typeInfo
                        end
                      end
                      
                    end
                  end
                  
                end
              end
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
  local pos = Parser.Position.new(0, 0)
  
  local txt = ""
  
  do
    local _exp = self.currentToken
    if _exp then
    
        pos = _exp.pos
        txt = _exp.txt
      end
  end
  
  error( string.format( "%s:%d:%d:(%s) %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:createNoneNode( pos )
  return NoneNode.new(pos, {builtinTypeNone})
end

function TransUnit:pushbackToken( token )
  table.insert( self.pushbackList, token )
  self.currentToken = Parser.getEofToken(  )
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
  self.currentToken = Parser.getEofToken(  )
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
    do
      local _exp = token
      if _exp then
      
          if _exp.kind ~= Parser.kind.Cmnt then
            break
          end
          table.insert( commentList, _exp )
        else
      
          break
        end
    end
    
  end
  do
    local _exp = token
    if _exp then
    
        if self.macroMode == "expand" and _exp.txt == ',,' then
          local nextToken = self:getTokenNoErr(  ) or _luneScript.error( 'unwrap val is nil' )
          
          local macroVal = self.symbol2ValueMapForMacro[nextToken.txt]
          
          if macroVal then
            local macroVal = macroVal
            
                if  not macroVal then
                  local _macroVal = macroVal
                  
                end
              
            if macroVal.typeInfo == builtinTypeSymbol then
              local txtList = macroVal.val
              
              for index = #txtList, 1, -1 do
                nextToken = Parser.Token.new(nextToken.kind, txtList[index], nextToken.pos)
                self:pushbackToken( nextToken )
              end
            elseif macroVal.typeInfo == builtinTypeStat then
              self:pushbackStr( string.format( "macroVal %s", nextToken.txt), macroVal.val )
            elseif macroVal.typeInfo:get_kind(  ) == TypeInfoKindArray or macroVal.typeInfo:get_kind(  ) == TypeInfoKindList then
              local strList = macroVal.val
              
              if strList then
                for index = #strList, 1, -1 do
                  self:pushbackStr( string.format( "macroVal %s[%d]", nextToken.txt, index), strList[index] )
                end
              else 
                self:error( string.format( "macro val is nil %s", nextToken.txt) )
              end
            else 
              local tokenList = {}
              
              expandVal( tokenList, macroVal.val, nextToken.pos )
              self:newPushback( tokenList )
            end
            nextToken = self:getTokenNoErr(  ) or _luneScript.error( 'unwrap val is nil' )
          else 
            self:error( string.format( "unknown macro val %s", nextToken.txt) )
          end
          token = nextToken
        end
      end
  end
  
  do
    local _exp = token
    if _exp then
    
        _exp:set_commentList( commentList )
      end
  end
  
  self.currentToken = token
  return token
end

function TransUnit:getToken(  )
  local token = self:getTokenNoErr(  )
  
      if  not token then
        local _token = token
        
        return Parser.getEofToken(  )
      end
    
  self.currentToken = token
  return token
end

function TransUnit:pushback(  )
  table.insert( self.pushbackList, self.currentToken )
  self.currentToken = Parser.getEofToken(  )
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
  local node = BlockNode.new(token.pos, {builtinTypeNone}, blockKind, stmtList)
  
  return node
end

function TransUnit:analyzeImport( token )
  if self.moduleScope ~= self.scope then
    self:error( "'import' must call at top scope." )
  end
  local moduleToken = self:getToken(  )
  
  local modulePath = moduleToken.txt
  
  local nextToken = moduleToken
  
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
  local moduleTypeInfo = rootTypeInfo
  
  for __index, moduleName in pairs( nameList ) do
    moduleTypeInfo = self:pushClass( nil, true, moduleName, "pub" )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  local moduleInfo = _luneScript.loadModule( modulePath )
  
  self.moduleName2Info[modulePath] = moduleInfo
  for __index, typeInfo in pairs( sym2builtInTypeMap ) do
    typeId2TypeInfo[typeInfo:get_typeId(  )] = typeInfo
  end
  local typeId2Scope = {}
  
  typeId2Scope[rootTypeId] = self.scope
  local newId2OldIdMap = {}
  
  local function registTypeInfo( atomInfo )
    local newTypeInfo = nil
    
    if not builtInTypeIdSet[atomInfo.typeId] then
      if atomInfo.nilable then
        local orgTypeInfo = typeId2TypeInfo[atomInfo.orgTypeId] or _luneScript.error( 'unwrap val is nil' )
        
        newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
        typeId2TypeInfo[atomInfo.typeId] = newTypeInfo or _luneScript.error( 'unwrap val is nil' )
      else 
        local itemTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.itemTypeId ) do
          table.insert( itemTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local argTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.argTypeId ) do
          if not typeId2TypeInfo[typeId] then
            Util.errorLog( string.format( "not found -- %s,%d", atomInfo.txt, #atomInfo.argTypeId) )
          end
          table.insert( argTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local retTypeInfo = {}
        
        for __index, typeId in pairs( atomInfo.retTypeId ) do
          table.insert( retTypeInfo, typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' ) )
        end
        local parentInfo = typeInfoRoot
        
        if atomInfo.parentId ~= rootTypeId then
          local workTypeInfo = typeId2TypeInfo[atomInfo.parentId]
          
              if  not workTypeInfo then
                local _workTypeInfo = workTypeInfo
                
                error( string.format( "not found parentInfo %s %s", atomInfo.parentId, atomInfo.txt) )
              end
            
          parentInfo = workTypeInfo
        end
        local baseInfo = typeId2TypeInfo[atomInfo.baseId] or _luneScript.error( 'unwrap val is nil' )
        
        local parentScope = typeId2Scope[atomInfo.parentId]
        
            if  not parentScope then
              local _parentScope = parentScope
              
              self:error( string.format( "not found parentScope %s %s", atomInfo.parentId, atomInfo.txt) )
            end
          
        if atomInfo.txt ~= "" then
          newTypeInfo = parentScope:getTypeInfoChild( atomInfo.txt )
        end
        if newTypeInfo and atomInfo.kind == TypeInfoKindClass then
          local newTypeInfo = newTypeInfo
          
              if  not newTypeInfo then
                local _newTypeInfo = newTypeInfo
                
              end
            
          typeId2Scope[atomInfo.typeId] = newTypeInfo:get_scope()
          if not newTypeInfo:get_scope() then
            error( string.format( "not found scope %s %s %s %s %s", parentScope, atomInfo.parentId, atomInfo.typeId, atomInfo.txt, newTypeInfo:getTxt(  )) )
          end
          typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
        else 
          if atomInfo.kind == TypeInfoKindClass then
            local baseScope = typeId2Scope[atomInfo.baseId] or _luneScript.error( 'unwrap val is nil' )
            
            local scope = Scope.new(parentScope, true, baseScope and {baseScope} or nil)
            
            local workTypeInfo = NormalTypeInfo.createClass( scope, baseInfo, parentInfo, true, "pub", atomInfo.txt )
            
            newTypeInfo = workTypeInfo
            typeId2Scope[atomInfo.typeId] = scope
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            parentScope:addClass( atomInfo.txt, workTypeInfo, scope )
          else 
            local scope = nil
            
            if atomInfo.kind == TypeInfoKindFunc or atomInfo.kind == TypeInfoKindMethod then
              scope = Scope.new(parentScope, false)
            end
            local workTypeInfo = NormalTypeInfo.create( scope, baseInfo, parentInfo, atomInfo.staticFlag, atomInfo.kind, atomInfo.txt, itemTypeInfo, argTypeInfo, retTypeInfo )
            
            newTypeInfo = workTypeInfo
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            if atomInfo.kind == TypeInfoKindFunc or atomInfo.kind == TypeInfoKindMethod then
              (typeId2Scope[atomInfo.parentId] or _luneScript.error( 'unwrap val is nil' ) ):add( atomInfo.txt, workTypeInfo )
              typeId2Scope[atomInfo.typeId] = scope
            end
          end
        end
      end
    else 
      newTypeInfo = sym2builtInTypeMap[atomInfo.txt]
      typeId2TypeInfo[atomInfo.typeId] = newTypeInfo or _luneScript.error( 'unwrap val is nil' )
    end
    return newTypeInfo or _luneScript.error( 'unwrap val is nil' )
  end
  
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    registTypeInfo( atomInfo )
  end
  for __index, atomInfo in pairs( moduleInfo._typeInfoList ) do
    if atomInfo.children and #atomInfo.children > 0 then
      local scope = typeId2Scope[atomInfo.typeId] or _luneScript.error( 'unwrap val is nil' )
      
      for __index, childId in pairs( atomInfo.children ) do
        local typeInfo = typeId2TypeInfo[childId] or _luneScript.error( 'unwrap val is nil' )
        
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
    local classTypeInfo = typeId2TypeInfo[classTypeId] or _luneScript.error( 'unwrap val is nil' )
    
    self:pushClass( nil, true, classTypeInfo:getTxt(  ), "pub" )
    local classInfo = moduleInfo._typeId2ClassInfoMap[classTypeId]
    
    if classInfo then
      for fieldName, fieldInfo in pairs( classInfo ) do
        local typeId = fieldInfo.typeId
        
        local fieldTypeInfo = typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' )
        
        self.scope:add( fieldName, fieldTypeInfo )
      end
    end
    for __index, child in pairs( classTypeInfo:get_children(  ) ) do
      if child:get_kind(  ) == TypeInfoKindClass then
        local oldId = newId2OldIdMap[child:get_typeId(  )]
        
        if oldId then
          registMember( oldId or _luneScript.error( 'unwrap val is nil' ) )
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
    self.scope:add( varName, typeId2TypeInfo[varInfo.typeId] or _luneScript.error( 'unwrap val is nil' ) )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  self.scope:add( moduleToken.txt, moduleTypeInfo )
  self:checkToken( nextToken, ";" )
  return ImportNode.new(token.pos, {builtinTypeNone}, modulePath)
end

function TransUnit:analyzeIfUnwrap( firstToken )
  local exp = self:analyzeExp(  )
  
  local scope = self:pushScope(  )
  
  local expType = exp:get_expType()
  
  if not expType:get_nilable() then
    self:addErrMess( exp.pos, "this is not nilable" )
  end
  scope:add( "_exp", expType:get_orgTypeInfo() )
  local block = self:analyzeBlock( "if!", scope )
  
  self:popScope(  )
  local elseBlock = nil
  
  local nextToken = self:getToken(  )
  
  if nextToken.txt == "else" then
    elseBlock = self:analyzeBlock( "if!" )
  else 
    self:pushback(  )
  end
  return IfUnwrapNode.new(firstToken.pos, {builtinTypeNone}, exp, block, elseBlock)
end

function TransUnit:analyzeIf( token )
  local nextToken = self:getToken(  )
  
  if nextToken.txt == "!" then
    return self:analyzeIfUnwrap( token )
  end
  self:pushback(  )
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
    table.insert( list, IfStmtInfo.new("else", NoneNode.new(nextToken.pos, {builtinTypeNone}), self:analyzeBlock( "else" )) )
  else 
    self:pushback(  )
  end
  return IfNode.new(token.pos, {builtinTypeNone}, list)
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
  return SwitchNode.new(firstToken.pos, {builtinTypeNone}, exp, caseList, defaultBlock)
end

function TransUnit:analyzeWhile( token )
  return WhileNode.new(token.pos, {builtinTypeNone}, self:analyzeExp(  ), self:analyzeBlock( "while" ))
end

function TransUnit:analyzeRepeat( token )
  local scope = self:pushScope( false )
  
  local node = RepeatNode.new(token.pos, {builtinTypeNone}, self:analyzeBlock( "repeat", scope ), self:analyzeExp(  ))
  
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
  
  if exp1:get_expType() ~= builtinTypeInt then
    self:addErrMess( exp1.pos, string.format( "exp1 is not int -- %s", exp1:get_expType():getTxt(  )) )
  end
  self.scope:add( val.txt, exp1:get_expType() )
  self:checkNextToken( "," )
  local exp2 = self:analyzeExp(  )
  
  if exp2:get_expType() ~= builtinTypeInt then
    self:addErrMess( exp2.pos, string.format( "exp2 is not int -- %s", exp2:get_expType():getTxt(  )) )
  end
  local token = self:getToken(  )
  
  local exp3 = nil
  
  if token.txt == "," then
    exp3 = self:analyzeExp(  )
    do
      local _exp = exp3
      if _exp then
      
          if _exp:get_expType() ~= builtinTypeInt then
            self:addErrMess( _exp.pos, string.format( "exp is not int -- %s", _exp:get_expType():getTxt(  )) )
          end
        end
    end
    
  else 
    self:pushback(  )
  end
  local node = ForNode.new(token.pos, {builtinTypeNone}, self:analyzeBlock( "for", scope ), val, exp1, exp2, exp3)
  
  self:popScope(  )
  return node
end

function TransUnit:analyzeApply( token )
  local scope = self:pushScope(  )
  
  local varList = {}
  
  local nextToken = Parser.getEofToken(  )
  
  repeat 
    local var = self:getSymbolToken(  )
    
    if var.kind ~= Parser.kind.Symb then
      self:error( "illegal symbol" )
    end
    table.insert( varList, var )
    nextToken = self:getToken(  )
    scope:add( var.txt, builtinTypeStem )
  until nextToken.txt ~= ","
  self:checkToken( nextToken, "of" )
  local exp = self:analyzeExp(  )
  
  if exp.kind ~= nodeKindExpCall then
    self:error( "not call" )
  end
  local block = self:analyzeBlock( "apply", scope )
  
  self:popScope(  )
  return ApplyNode.new(token.pos, {builtinTypeNone}, varList, exp, block)
end

function TransUnit:analyzeForeach( token, sortFlag )
  local scope = self:pushScope(  )
  
  local valSymbol = Parser.getEofToken(  )
  
  local keySymbol = nil
  
  local nextToken = Parser.getEofToken(  )
  
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
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, itemTypeInfoList[1] )
          end
      end
      
    elseif exp:get_expType():get_kind(  ) == TypeInfoKindList or exp:get_expType():get_kind(  ) == TypeInfoKindArray then
      self.scope:add( valSymbol.txt, itemTypeInfoList[1] )
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, builtinTypeInt )
          else
        
            self.scope:add( "__index", builtinTypeInt )
          end
      end
      
    else 
      self:error( string.format( "unknown kind type of exp for foreach-- %d:%d", exp.pos.lineNo, exp.pos.column) )
    end
  end
  local block = self:analyzeBlock( "foreach", scope )
  
  self:popScope(  )
  if sortFlag then
    return ForsortNode.new(token.pos, {builtinTypeNone}, valSymbol, keySymbol, exp, block, sortFlag)
  else 
    return ForeachNode.new(token.pos, {builtinTypeNone}, valSymbol, keySymbol, exp, block, sortFlag)
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
  local name = NoneNode.new(firstToken.pos, {builtinTypeNone})
  
  -- none
  
  local typeInfo = builtinTypeStem
  
  self:checkSymbol( token )
  name = self:analyzeExpSymbol( firstToken, token, "symbol", nil, true )
  typeInfo = name:get_expType()
  token = self:getToken(  )
  if token.txt == "!" then
    typeInfo = typeInfo:get_nilableTypeInfo(  ) or _luneScript.error( 'unwrap val is nil' )
    token = self:getToken(  )
  end
  local arrayMode = "no"
  
  while true do
    if token.txt == '[' or token.txt == '[@' then
      if token.txt == '[' then
        arrayMode = "list"
        typeInfo = NormalTypeInfo.createList( accessMode, self:getCurrentClass(  ), {typeInfo} )
      else 
        arrayMode = "array"
        typeInfo = NormalTypeInfo.createArray( accessMode, self:getCurrentClass(  ), {typeInfo} )
      end
      token = self:getToken(  )
      if token.txt ~= ']' then
        self:pushback(  )
        self:checkNextToken( ']' )
      end
    elseif token.txt == "<" then
      local genericList = {}
      
      local nextToken = Parser.getEofToken(  )
      
      repeat 
        local typeExp = self:analyzeRefType( accessMode )
        
        table.insert( genericList, typeExp:get_expType() )
        nextToken = self:getToken(  )
      until nextToken.txt ~= ","
      self:checkToken( nextToken, '>' )
      if typeInfo:get_kind() == TypeInfoKindMap then
        typeInfo = NormalTypeInfo.createMap( accessMode, self:getCurrentClass(  ), genericList[1] or builtinTypeStem, genericList[2] or builtinTypeStem )
      else 
        self:error( string.format( "not support generic: %s", typeInfo:getTxt(  ) ) )
      end
    else 
      self:pushback(  )
      break
    end
    token = self:getToken(  )
  end
  if token.txt == "!" then
    typeInfo = typeInfo:get_nilableTypeInfo(  ) or _luneScript.error( 'unwrap val is nil' )
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
      table.insert( argList, DeclArgDDDNode.new(argName.pos, {builtinTypeDDD}) )
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
  self.rootScope = self.scope
  self:registBuiltInScope(  )
  for name, child in pairs( self.scope:get_symbol2TypeInfoMap() ) do
  end
  local moduleTypeInfo = typeInfoRoot
  
  if module then
    for txt in string.gmatch( module, '[^%.]+' ) do
      moduleTypeInfo = self:pushClass( nil, false, txt, "pub" ) or _luneScript.error( 'unwrap val is nil' )
    end
  end
  self.moduleScope = self.scope
  self.parser = parser
  self.moduleName2Info = {}
  local ast = nil
  
  if macroFlag then
    ast = self:analyzeBlock( "macro" )
  else 
    local children = {}
    
    ast = RootNode.new(Parser.Position.new(0, 0), {builtinTypeNone}, children, self.typeId2ClassMap)
    self:analyzeStatementList( children )
    local token = self:getTokenNoErr(  )
    
    do
      local _exp = token
      if _exp then
      
          error( string.format( "unknown:%d:%d:(%s) %s", _exp.pos.lineNo, _exp.pos.column, Parser.getKindTxt( _exp.kind ), _exp.txt) )
        end
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
  return ASTInfo.new(ast or _luneScript.error( 'unwrap val is nil' ), moduleTypeInfo)
end

function TransUnit:analyzeDeclMacro( accessMode, firstToken )
  local nameToken = self:getToken(  )
  
  self:checkNextToken( "(" )
  local scope = self:pushScope(  )
  
  local argList = {}
  
  local nextToken = self:analyzeDeclArgList( accessMode, argList )
  
  local argTypeList = {}
  
  for __index, argNode in pairs( argList ) do
    table.insert( argTypeList, argNode:get_expType() )
  end
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
    ast = BlockNode.new(firstToken.pos, {builtinTypeNone}, "macro", stmtList)
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
  local typeInfo = NormalTypeInfo.createFunc( scope, TypeInfoKindMacro, self:getCurrentNamespaceTypeInfo(  ), false, false, false, accessMode, nameToken.txt, argTypeList )
  
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
    local typeInfo = nil
    
    do
      local _exp = baseRef
      if _exp then
      
          typeInfo = _exp:get_expType(  )
        end
    end
    
    self:pushClass( typeInfo, false, name.txt, accessMode )
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
    return self:analyzeDeclVar( "let", accessMode, staticFlag, firstToken )
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
  local typeInfo = nil
  
  do
    local _exp = baseRef
    if _exp then
    
        typeInfo = _exp:get_expType(  )
      end
  end
  
  local classTypeInfo = self:pushClass( typeInfo, false, name.txt, classAccessMode )
  
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
  
  local memberTypeList = {}
  
  for __index, memberNode in pairs( memberList ) do
    local memberType = memberNode:get_expType()
    
    table.insert( memberTypeList, memberType )
    if memberNode:get_expType():get_accessMode() ~= "pub" then
      memberType = NormalTypeInfo.cloneToPublic( memberType )
    end
    local memberName = memberNode:get_name()
    
    local getterName = "get_" .. memberName.txt
    
    local accessMode = memberNode:get_getterMode()
    
    if accessMode ~= "none" and not self.scope:getTypeInfoChild( getterName ) then
      local retTypeInfo = NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, "pub", getterName, {}, {memberType} )
      
      self:popScope(  )
      self.scope:add( getterName, retTypeInfo )
    end
    local setterName = "set_" .. memberName.txt
    
    local accessMode = memberNode:get_setterMode()
    
    if memberNode:get_setterMode() ~= "none" and not self.scope:getTypeInfoChild( setterName ) then
      self.scope:add( setterName, NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, "pub", setterName, {memberType}, nil ) )
      self:popScope(  )
    end
  end
  if not self.scope:getTypeInfoChild( "__init" ) then
    local initTypeInfo = NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, "pub", "__init", memberTypeList, {} )
    
    self:popScope(  )
    self.scope:add( "__init", initTypeInfo )
  end
  self:popClass(  )
  return node
end

function TransUnit:addMethod( className, methodNode, name )
  local classTypeInfo = self.scope:getTypeInfo( className )
  
  local classNodeInfo = self.typeInfo2ClassNode[classTypeInfo] or _luneScript.error( 'unwrap val is nil' )
  
  classNodeInfo.outerMethodSet[name] = true
  table.insert( classNodeInfo.fieldList, methodNode )
end

function TransUnit:analyzeDeclFunc( overrideFlag, accessMode, staticFlag, classNameToken, firstToken, name )
  local token = self:getToken(  )
  
  do
    local _exp = name
    if _exp then
    
        name = self:checkSymbol( _exp )
      else
    
        if token.txt ~= "(" then
          name = self:checkSymbol( token )
          token = self:getToken(  )
        end
      end
  end
  
  local needPopFlag = false
  
  if token.txt == "." then
    needPopFlag = true
    classNameToken = name
    self:pushClass( nil, false, (name or _luneScript.error( 'unwrap val is nil' ) ).txt, "pub" )
    name = self:getSymbolToken(  )
    token = self:getToken(  )
  end
  local kind = nodeKindDeclConstr
  
  local typeKind = TypeInfoKindFunc
  
  if classNameToken then
    if not staticFlag then
      typeKind = TypeInfoKindMethod
    end
    if (name or _luneScript.error( 'unwrap val is nil' ) ).txt == "__init" then
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
  local funcName = ""
  
  do
    local _exp = name
    if _exp then
    
        funcName = _exp.txt
      end
  end
  
  if overrideFlag then
    local name = name
    
        if  not name then
          local _name = name
          
          self:error( "name is nil" )
        end
      
    local overrideType = self.scope:getTypeInfoMethod( funcName )
    
        if  not overrideType then
          local _overrideType = overrideType
          
          self:error( "not found override -- " .. funcName )
        end
      
    if overrideType:get_accessMode(  ) ~= accessMode then
      self:error( string.format( "missmatch override accessMode -- %s,%s,%s", funcName, overrideType:get_accessMode(  ), accessMode) )
    end
    if overrideType:get_staticFlag(  ) ~= staticFlag then
      self:error( "missmatch override staticFlag -- " .. funcName )
    end
    if overrideType:get_kind(  ) ~= TypeInfoKindMethod then
      self:error( string.format( "missmatch override kind -- %s, %d", funcName, overrideType:get_kind(  )) )
    end
  else 
    do
      local _exp = name
      if _exp then
      
          if _exp.txt ~= "__init" and self.scope:getTypeInfoMethod( _exp.txt ) then
            self:error( "missmatch override --" .. funcName )
          end
        end
    end
    
  end
  self:checkToken( token, "(" )
  local scope = self:pushScope(  )
  
  local argList = {}
  
  token = self:analyzeDeclArgList( accessMode, argList )
  local argTypeList = {}
  
  for __index, argNode in pairs( argList ) do
    table.insert( argTypeList, argNode:get_expType() )
  end
  self:checkToken( token, ")" )
  token = self:getToken(  )
  local retTypeInfoList = {}
  
  if token.txt == ":" then
    repeat 
      local refType = self:analyzeRefType( accessMode )
      
      table.insert( retTypeInfoList, refType:get_expType() )
      token = self:getToken(  )
    until token.txt ~= ","
  end
  local typeInfo = NormalTypeInfo.createFunc( scope, typeKind, self:getCurrentNamespaceTypeInfo(  ), false, false, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )
  
  do
    local _exp = name
    if _exp then
    
        scope:get_parent(  ):add( _exp.txt, typeInfo )
      end
  end
  
  local node = NoneNode.new(firstToken.pos)
  
  if token.txt == ";" then
    node = self:createNoneNode( firstToken.pos )
  else 
    self:pushback(  )
    local body = self:analyzeBlock( "func", scope )
    
    local info = DeclFuncInfo.new(classNameToken, name, argList, staticFlag, accessMode, body, retTypeInfoList)
    
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
  self:popScope(  )
  if needPopFlag then
    self:addMethod( (classNameToken or _luneScript.error( 'unwrap val is nil' ) ).txt, node, funcName )
    self:popClass(  )
  end
  return node
end

function TransUnit:analyzeDeclVar( mode, accessMode, staticFlag, firstToken )
  local unwrapFlag = false
  
  local token = self:getToken(  )
  
  if token.txt == "!" then
    unwrapFlag = true
  else 
    self:pushback(  )
    if mode ~= "let" then
      Util.errorLog( "need '!'" )
    end
  end
  local typeInfoList = {}
  
  local varNameList = {}
  
  local varTypeList = {}
  
  repeat 
    local varName = self:getSymbolToken(  )
    
    token = self:getToken(  )
    local typeInfo = builtinTypeNone
    
    if token.txt == ":" then
      local refType = self:analyzeRefType( accessMode )
      
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
  local orgExpTypeList = {}
  
  do
    local _exp = expList
    if _exp then
    
        local expTypeList = {}
        
        for index, exp in pairs( _exp.expList ) do
          if index == #_exp.expList then
            if exp:get_expType() == builtinTypeDDD then
              for subIndex = index, #varNameList do
                local argType = typeInfoList[subIndex]
                
                if argType ~= builtinTypeNone and not argType:isSettableFrom( builtinTypeStem ) then
                  self:addErrMess( firstToken.pos, string.format( "unmatch value type %s(%d) <- %s(%d)", argType:getTxt(  ), argType:get_typeId(), builtinTypeStem:getTxt(  ), builtinTypeStem:get_typeId()) )
                end
                table.insert( expTypeList, builtinTypeStem )
                table.insert( orgExpTypeList, builtinTypeStem )
              end
            else 
              for __index, typeInfo in pairs( exp:get_expTypeList() ) do
                table.insert( orgExpTypeList, typeInfo )
                if unwrapFlag and typeInfo:get_nilable() then
                  typeInfo = typeInfo:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
                end
                table.insert( expTypeList, typeInfo )
                local argType = typeInfoList[index]
                
                if not (unwrapFlag and typeInfo == builtinTypeNil ) and argType ~= builtinTypeNone and not argType:isSettableFrom( typeInfo ) then
                  self:addErrMess( firstToken.pos, string.format( "unmatch value type %s <- %s", argType:getTxt(  ), typeInfo:getTxt(  )) )
                end
              end
            end
          else 
            local expTypeInfo = builtinTypeStem
            
            if exp:get_expType() == builtinTypeDDD then
              expTypeInfo = builtinTypeStem
            else 
              expTypeInfo = exp:get_expType()
            end
            table.insert( orgExpTypeList, expTypeInfo )
            if unwrapFlag and expTypeInfo:get_nilable() then
              expTypeInfo = expTypeInfo:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
            end
            local argType = typeInfoList[index]
            
            if argType ~= builtinTypeNone and not argType:isSettableFrom( expTypeInfo ) then
              self:addErrMess( firstToken.pos, string.format( "unmatch value type %s <- %s", argType:getTxt(  ), expTypeInfo:getTxt(  )) )
            end
            table.insert( expTypeList, expTypeInfo )
          end
        end
        for index, typeInfo in pairs( expTypeList ) do
          if not typeInfoList[index] or typeInfoList[index] == builtinTypeNone then
            typeInfoList[index] = typeInfo
          end
        end
      end
  end
  
  if self.macroScope == self.scope then
    for index, varName in pairs( varNameList ) do
      local typeInfo = typeInfoList[index]
      
      self.symbol2ValueMapForMacro[varName.txt] = MacroValInfo.new(nil, typeInfo)
    end
  end
  local varList = {}
  
  local sameSymbolList = {}
  
  for index, varName in pairs( varNameList ) do
    local typeInfo = typeInfoList[index]
    
    local varInfo = VarInfo.new(varName, varTypeList[index], typeInfo)
    
    table.insert( varList, varInfo )
    if not varTypeList[index] and typeInfo == builtinTypeNil then
      self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
    end
    local sameTypeInfo = self.scope:getTypeInfo( varName.txt )
    
    do
      local _exp = sameTypeInfo
      if _exp then
      
          table.insert( sameSymbolList, varInfo )
        end
    end
    
    if mode == "let" or mode == "sync" then
      self.scope:add( varName.txt, typeInfo )
    end
  end
  local unwrapBlock = nil
  
  local thenBlock = nil
  
  if unwrapFlag then
    local scope = self:pushScope(  )
    
    for index, varName in pairs( varNameList ) do
      self.scope:add( "_" .. varName.txt, orgExpTypeList[index] )
    end
    unwrapBlock = self:analyzeBlock( "let!", scope )
    self:popScope(  )
    token = self:getToken(  )
    if token.txt == "then" then
      thenBlock = self:analyzeBlock( "let!", scope )
    else 
      self:pushback(  )
    end
  end
  local syncBlock = nil
  
  if mode == "sync" then
    local nextToken = self:getToken(  )
    
    if nextToken.txt == "do" then
      syncBlock = self:analyzeBlock( "let!", scope )
    else 
      self:pushback(  )
    end
  end
  self:checkNextToken( ";" )
  local node = DeclVarNode.new(firstToken.pos, {builtinTypeNone}, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, sameSymbolList, syncBlock)
  
  return node
end

function TransUnit:analyzeExpList( skipOp2Flag )
  local expList = {}
  
  local pos = nil
  
  local expTypeList = {}
  
  repeat 
    local exp = self:analyzeExp( skipOp2Flag )
    
    if not pos then
      pos = exp.pos
    end
    table.insert( expList, exp )
    table.insert( expTypeList, exp:get_expType() )
    local token = self:getToken(  )
    
  until token.txt ~= ","
  self:pushback(  )
  return ExpListNode.new(pos or Parser.Position.new(0, 0), expTypeList, expList)
end

function TransUnit:analyzeListConst( token )
  local nextToken = self:getToken(  )
  
  local expList = nil
  
  local itemTypeInfo = builtinTypeNone
  
  if nextToken.txt ~= "]" then
    self:pushback(  )
    expList = self:analyzeExpList(  )
    self:checkNextToken( "]" )
    local nodeList = (expList or _luneScript.error( 'unwrap val is nil' ) ).expList
    
    for __index, exp in pairs( nodeList ) do
      if itemTypeInfo == builtinTypeNone then
        itemTypeInfo = exp:get_expType()
      elseif not itemTypeInfo:isSettableFrom( exp:get_expType() ) then
        itemTypeInfo = builtinTypeStem
      end
    end
  end
  local kind = nodeKindLiteralArray
  
  local typeInfoList = {builtinTypeNone}
  
  if token.txt == '[' then
    kind = nodeKindLiteralList
    typeInfoList = {NormalTypeInfo.createList( "pri", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralListNode.new(token.pos, typeInfoList, expList)
  else 
    typeInfoList = {NormalTypeInfo.createArray( "pri", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralArrayNode.new(token.pos, typeInfoList, expList)
  end
end

function TransUnit:analyzeMapConst( token )
  local nextToken = self:getToken(  )
  
  local map = {}
  
  local pairList = {}
  
  local keyTypeInfo = builtinTypeNone
  
  local valTypeInfo = builtinTypeNone
  
  while true do
    if nextToken.txt == "}" then
      break
    end
    self:pushback(  )
    local key = self:analyzeExp(  )
    
    if not keyTypeInfo:isSettableFrom( key:get_expType() ) then
      if keyTypeInfo ~= builtinTypeNone then
        keyTypeInfo = builtinTypeStem
      else 
        keyTypeInfo = key:get_expType()
      end
    end
    self:checkNextToken( ":" )
    local val = self:analyzeExp(  )
    
    if not valTypeInfo:isSettableFrom( val:get_expType() ) then
      if valTypeInfo ~= builtinTypeNone then
        valTypeInfo = builtinTypeStem
      else 
        valTypeInfo = val:get_expType()
      end
    end
    table.insert( pairList, PairItem.new(key, val) )
    map[key] = val
    nextToken = self:getToken(  )
    if nextToken.txt ~= "," then
      break
    end
    nextToken = self:getToken(  )
  end
  local typeInfo = NormalTypeInfo.createMap( "pri", self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
  
  self:checkToken( nextToken, "}" )
  return LiteralMapNode.new(token.pos, {typeInfo}, map, pairList)
end

function TransUnit:analyzeExpRefItem( token, exp )
  local indexExp = self:analyzeExp(  )
  
  self:checkNextToken( "]" )
  local typeInfo = builtinTypeStem
  
  local expType = exp:get_expType()
  
  if expType then
    if expType:get_kind() == TypeInfoKindMap then
      typeInfo = expType:get_itemTypeInfoList(  )[2]
      if typeInfo ~= builtinTypeStem and not typeInfo:get_nilable() then
        typeInfo = typeInfo:get_nilableTypeInfo()
      end
    elseif expType:get_kind() == TypeInfoKindArray or expType:get_kind() == TypeInfoKindList then
      typeInfo = expType:get_itemTypeInfoList(  )[1]
    elseif expType == builtinTypeString then
      typeInfo = builtinTypeInt
    else 
      self:addErrMess( exp.pos, "could not access with []." )
    end
  end
  if not typeInfo then
    Util.errorLog( "illegal type" )
  end
  return ExpRefItemNode.new(token.pos, {typeInfo}, exp, indexExp)
end

function TransUnit:checkMatchValType( pos, funcTypeInfo, expList, genericTypeList )
  local argTypeList = funcTypeInfo:get_argTypeInfoList()
  
  do
    local _switchExp = funcTypeInfo
    if _switchExp == typeInfoListInsert then
      argTypeList = genericTypeList
    elseif _switchExp == typeInfoListRemove then
    end
  end
  
  do
    local _exp = expList
    if _exp then
    
        local expNodeList = _exp:get_expList()
        
        for index, expNode in pairs( expNodeList ) do
          if #argTypeList < index then
            self:addErrMess( pos, string.format( "%s: over argument. expect:%d, actual:%d", funcTypeInfo:getTxt(  ), #argTypeList, #expNodeList) )
            break
          end
          local argType = argTypeList[index]
          
          local expType = expNode:get_expType()
          
          if #argTypeList == index then
            if argType ~= builtinTypeDDD then
              if not argType:isSettableFrom( expType ) then
                self:addErrMess( expNode.pos, string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
              end
            end
            break
          elseif #expNodeList == index then
            if expType == builtinTypeDDD then
              for argIndex = index, #argTypeList do
                local workArgType = argTypeList[argIndex]
                
                if not workArgType:isSettableFrom( builtinTypeStem ) then
                  self:addErrMess( expNode.pos, string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), argIndex, workArgType:getTxt(  ), expType:getTxt(  )) )
                end
              end
            end
            break
          end
          if not argType:isSettableFrom( expType ) then
            self:addErrMess( expNode.pos, string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
          end
        end
      else
    
        if #argTypeList ~= 0 then
        end
      end
  end
  
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
  do
    local _exp = expList
    if _exp then
    
        if _exp:get_expList(  ) then
          for __index, exp in pairs( _exp:get_expList(  ) ) do
            local kind = exp.kind
            
            if kind ~= nodeKindLiteralNil and kind ~= nodeKindLiteralChar and kind ~= nodeKindLiteralInt and kind ~= nodeKindLiteralReal and kind ~= nodeKindLiteralArray and kind ~= nodeKindLiteralList and kind ~= nodeKindLiteralMap and kind ~= nodeKindLiteralString and kind ~= nodeKindLiteralBool and kind ~= nodeKindLiteralSymbol and kind ~= nodeKindRefField and kind ~= nodeKindExpMacroStat then
              self:error( "Macro arguments must be literal value." )
            end
          end
        end
      end
  end
  
  local macroInfo = self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )] or _luneScript.error( 'unwrap val is nil' )
  
  local argVal = {}
  
  local argType = {}
  
  do
    local _exp = expList
    if _exp then
    
        for __index, argNode in pairs( _exp:get_expList(  ) ) do
          local valList, typeInfoList = argNode:getLiteral(  )
          
          local val = valList[1]
          
          local typeInfo = typeInfoList[1]
          
          table.insert( argVal, val )
          table.insert( argType, typeInfo )
        end
      end
  end
  
  local func = macroInfo.func
  
  local macroVars = func( table.unpack( argVal ) )
  
  for __index, name in pairs( macroVars._names ) do
    local valInfo = macroInfo.symbol2MacroValInfoMap[name] or _luneScript.error( 'unwrap val is nil' )
    
    local typeInfo = valInfo and valInfo.typeInfo or builtinTypeStem
    
    local val = macroVars[name]
    
    if typeInfo == builtinTypeSymbol then
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
  local parser = MacroPaser.new(macroInfo.declInfo.tokenList, string.format( "macro %s", macroTypeInfo:getTxt(  )))
  
  local bakParser = self.parser
  
  self.parser = parser
  self.macroMode = "expand"
  local stmtList = {}
  
  self:analyzeStatementList( stmtList, "}" )
  self.macroMode = "none"
  self.parser = bakParser
  return ExpMacroExpNode.new(firstToken.pos, {builtinTypeNone}, stmtList)
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
        
        local funcTypeInfo = exp:get_expType()
        
        if funcTypeInfo:get_kind(  ) == TypeInfoKindMacro then
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
        local genericTypeList = funcTypeInfo:get_itemTypeInfoList()
        
        if funcTypeInfo:get_kind() == TypeInfoKindMethod and exp:get_kind() == nodeKindRefField then
          local refField = exp
          
          local classType = refField:get_prefix():get_expType()
          
          genericTypeList = classType:get_itemTypeInfoList()
        end
        self:checkMatchValType( exp.pos, funcTypeInfo, expList, genericTypeList )
        if macroFlag then
          self.macroMode = "none"
          exp = self:evalMacro( firstToken, funcTypeInfo, expList )
        else 
          exp = ExpCallNode.new(firstToken.pos, funcTypeInfo:get_retTypeInfoList(  ), exp, expList)
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
    local prefixExp = prefixExp
    
        if  not prefixExp then
          local _prefixExp = prefixExp
          
        end
      
    if self.macroMode == "analyze" then
      exp = RefFieldNode.new(firstToken.pos, {builtinTypeSymbol}, token, prefixExp or _luneScript.error( 'unwrap val is nil' ))
    else 
      local typeInfo = builtinTypeStem
      
      local prefixExpType = prefixExp:get_expType()
      
      if not prefixExpType then
        self:error( "unknown prefix type: " .. getNodeKindName( prefixExp.kind ) )
      end
      local getterTypeInfo = nil
      
      if prefixExpType:get_kind(  ) == TypeInfoKindClass or prefixExpType:get_kind(  ) == TypeInfoKindList then
        if prefixExpType:get_kind(  ) == TypeInfoKindList then
          prefixExpType = sym2builtInTypeMap["List"] or _luneScript.error( 'unwrap val is nil' )
        end
        local className = prefixExpType:getTxt(  )
        
        local classScope = prefixExpType:get_scope()
        
            if  not classScope then
              local _classScope = classScope
              
              self:error( string.format( "not found field: %s, %s, %s", classScope, className, prefixExpType) )
            end
          
        typeInfo = nil
        if mode == "get" then
          typeInfo = classScope:getTypeInfo( string.format( "get_%s", token.txt) )
          do
            local _exp = typeInfo
            if _exp then
            
                if (_exp:get_kind(  ) == TypeInfoKindFunc or _exp:get_kind(  ) == TypeInfoKindMethod ) then
                  local retTypeList = _exp:get_retTypeInfoList(  )
                  
                  getterTypeInfo = _exp
                  typeInfo = retTypeList[1]
                end
              end
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
      elseif prefixExpType:get_kind(  ) == TypeInfoKindMap then
        local work = prefixExpType:get_itemTypeInfoList()[1]
        
        if work ~= builtinTypeString then
          self:addErrMess( token.pos, string.format( "map key type is not str. (%s)", work:getTxt(  )) )
        end
        typeInfo = prefixExpType:get_itemTypeInfoList()[2]
        do
          local _exp = typeInfo
          if _exp then
          
              if _exp:get_nilable() then
                typeInfo = _exp:get_nilableTypeInfo()
              end
            end
        end
        
      elseif prefixExpType == builtinTypeStem then
      else 
        self:error( string.format( "illegal type -- %s, %d", prefixExpType:getTxt(  ), prefixExpType:get_kind(  )) )
      end
      do
        local _exp = getterTypeInfo
        if _exp then
        
            exp = GetFieldNode.new(firstToken.pos, {typeInfo or _luneScript.error( 'unwrap val is nil' )}, token, prefixExp, _exp)
          else
        
            exp = RefFieldNode.new(firstToken.pos, {typeInfo or _luneScript.error( 'unwrap val is nil' )}, token, prefixExp)
          end
      end
      
    end
  elseif mode == "symbol" then
    if self.macroMode == "analyze" then
      exp = LiteralSymbolNode.new(firstToken.pos, {builtinTypeSymbol}, token)
    else 
      local typeInfo = self.scope:getTypeInfo( token.txt )
      
      if not typeInfo and token.txt == "self" then
        typeInfo = self:getCurrentClass(  )
      end
      local typeInfo = typeInfo
      
          if  not typeInfo then
            local _typeInfo = typeInfo
            
            self:error( "not found type -- " .. token.txt )
          end
        
      if typeInfo == builtinTypeSymbol then
        skipFlag = true
      end
      exp = ExpRefNode.new(firstToken.pos, {typeInfo}, token)
    end
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( false, "pri", false, nil, token, nil )
  else 
    self:error( "illegal mode", mode )
  end
  return self:analyzeExpCont( firstToken, exp or _luneScript.error( 'unwrap val is nil' ), skipFlag )
end

function TransUnit:analyzeExpOp2( firstToken, exp, prevOpLevel )
  while true do
    local nextToken = self:getToken(  )
    
    local opTxt = nextToken.txt
    
    if opTxt == "@" then
      local castType = self:analyzeRefType( "pri" )
      
      exp = ExpCastNode.new(firstToken.pos, castType:get_expTypeList(), exp)
    elseif nextToken.kind == Parser.kind.Ope then
      if Parser.isOp2( opTxt ) then
        local opLevel = op2levelMap[opTxt]
        
            if  not opLevel then
              local _opLevel = opLevel
              
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
        local retType = builtinTypeNone
        
        local exp1Type = exp:get_expType()
        
        local exp2Type = exp2:get_expType()
        
        if not exp1Type then
          self:error( string.format( "expType is nil %s:%d:%d", self.parser:getStreamName(  ), firstToken.pos.lineNo, firstToken.pos.column) )
        end
        do
          local _switchExp = opTxt
          if _switchExp == "or" then
            if exp1Type:equals( exp2Type ) then
              retType = exp1Type
            elseif exp1Type:get_kind() == exp2Type:get_kind() then
              retType = exp1Type
            elseif exp2Type == builtinTypeNil then
              retType = exp1Type
            elseif exp1Type == builtinTypeNil then
              retType = exp2Type
            else 
              retType = builtinTypeStem
            end
          elseif _switchExp == "and" then
            retType = exp2Type
          elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" then
            if exp1Type ~= builtinTypeInt or exp2Type ~= builtinTypeInt then
              self:addErrMess( nextToken.pos, string.format( "no int type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = builtinTypeBool
          elseif _switchExp == "~=" or _switchExp == "==" then
            if (not exp1Type:isSettableFrom( exp2Type ) and not exp2Type:isSettableFrom( exp1Type ) ) then
              self:addErrMess( nextToken.pos, string.format( "not compatible type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = builtinTypeBool
          elseif _switchExp == "^" or _switchExp == "|" or _switchExp == "~" or _switchExp == "&" or _switchExp == "<<" or _switchExp == ">>" then
            if exp1Type ~= builtinTypeInt or exp2Type ~= builtinTypeInt then
              self:addErrMess( nextToken.pos, string.format( "no int type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = builtinTypeInt
          elseif _switchExp == ".." then
            if exp1Type ~= builtinTypeString or exp1Type ~= builtinTypeString then
              self:addErrMess( nextToken.pos, string.format( "no string type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            retType = builtinTypeString
          elseif _switchExp == "+" or _switchExp == "-" or _switchExp == "*" or _switchExp == "/" or _switchExp == "//" or _switchExp == "%" then
            if (exp1Type ~= builtinTypeReal and exp1Type ~= builtinTypeInt ) or (exp2Type ~= builtinTypeReal and exp2Type ~= builtinTypeInt ) then
              self:addErrMess( nextToken.pos, string.format( "no numeric type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
            if exp1Type == builtinTypeReal or exp2Type == builtinTypeReal then
              retType = builtinTypeReal
            else 
              retType = builtinTypeInt
            end
          elseif _switchExp == "=" then
            if not exp1Type:isSettableFrom( exp2Type ) then
              self:addErrMess( nextToken.pos, string.format( "unmatch type %s and %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
            end
          else 
            self:error( "unknown op " .. opTxt )
          end
        end
        
        exp = ExpOp2Node.new(firstToken.pos, {retType}, nextToken, exp, exp2)
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
      local exp = self:analyzeExp( true, op1levelMap[token.txt] or _luneScript.error( 'unwrap val is nil' ) )
      
      local nextToken = self:getToken(  )
      
      if nextToken.txt ~= "$" then
        self:pushback(  )
      end
      local format = token.txt == ",,," and "'%s '" or '"\'%s\'"'
      
      if token.txt == ",," and exp.kind == nodeKindExpRef then
        local refToken = (exp ):get_token(  )
        
        local macroInfo = self.symbol2ValueMapForMacro[refToken.txt]
        
        if macroInfo then
          if (macroInfo or _luneScript.error( 'unwrap val is nil' ) ).typeInfo == builtinTypeSymbol then
            format = "'%s '"
          end
        end
      end
      local newToken = Parser.Token.new(Parser.kind.Str, format, token.pos)
      
      local literalStr = LiteralStringNode.new(token.pos, {builtinTypeString}, newToken, {exp})
      
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
      
      local literalStr = LiteralStringNode.new(token.pos, {builtinTypeString}, newToken, {})
      
      table.insert( expStrList, literalStr )
    end
  end
  return ExpMacroStatNode.new(firstToken.pos, {builtinTypeStat}, expStrList)
end

function TransUnit:analyzeSuper( firstToken )
  self:checkNextToken( "(" )
  local expList = self:analyzeExpList(  )
  
  self:checkNextToken( ")" )
  self:checkNextToken( ";" )
  local classType = self:getCurrentClass(  )
  
  local superType = classType:get_baseTypeInfo(  )
  
  return ExpCallSuperNode.new(firstToken.pos, {builtinTypeNone}, superType, expList)
end

function TransUnit:analyzeUnwrap( firstToken )
  local nextToken = self:getToken(  )
  
  if nextToken.txt ~= "!" then
    self:pushback(  )
    self:pushbackToken( firstToken )
    local exp = self:analyzeExp(  )
    
    self:checkNextToken( ";" )
    return StmtExpNode.new(nextToken.pos, {builtinTypeNone}, exp)
  end
  self:pushback(  )
  return self:analyzeDeclVar( "unwrap", "pri", false, firstToken )
end

function TransUnit:analyzeExpUnwrap( firstToken )
  local expNode = self:analyzeExp(  )
  
  local nextToken = self:getToken(  )
  
  local insNode = nil
  
  if nextToken.txt == "default" then
    insNode = self:analyzeExp(  )
  else 
    self:pushback(  )
  end
  local unwrapType = builtinTypeStem
  
  local expType = expNode:get_expType()
  
  if not expType:get_nilable() then
    unwrapType = expType
  else 
    unwrapType = expType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
    do
      local _exp = insNode
      if _exp then
      
          local insType = _exp:get_expType()
          
          unwrapType = insType
          if not unwrapType:isSettableFrom( insType ) then
            self:addErrMess( _exp.pos, string.format( "unmatch type: %s <- %s", unwrapType:getTxt(  ), insType:getTxt(  )) )
          end
        end
    end
    
  end
  return ExpUnwrapNode.new(firstToken.pos, {unwrapType}, expNode, insNode)
end

function TransUnit:analyzeExp( skipOp2Flag, prevOpLevel )
  local firstToken = self:getToken(  )
  
  local token = firstToken
  
  local exp = NoneNode.new(firstToken.pos, {builtinTypeNone})
  
  if token.kind == Parser.kind.Dlmt then
    if token.txt == "..." then
      return ExpDDDNode.new(firstToken.pos, {builtinTypeNone}, token)
    end
    if token.txt == '[' or token.txt == '[@' then
      exp = self:analyzeListConst( token )
    end
    if token.txt == '{' then
      exp = self:analyzeMapConst( token )
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
    local classTypeInfo = exp:get_expType()
    
    local classScope = classTypeInfo:get_scope(  )
    
    local initTypeInfo = (classScope or _luneScript.error( 'unwrap val is nil' ) ):getTypeInfoChild( "__init" )
    
        if  not initTypeInfo then
          local _initTypeInfo = initTypeInfo
          
          self:error( "not found __init" )
        end
      
    self:checkMatchValType( exp.pos, initTypeInfo, argList, exp:get_expType():get_itemTypeInfoList() )
    exp = ExpNewNode.new(firstToken.pos, exp:get_expTypeList(), exp, argList)
    exp = self:analyzeExpCont( firstToken, exp, false )
  end
  if token.kind == Parser.kind.Ope and Parser.isOp1( token.txt ) then
    if token.txt == "`" then
      exp = self:analyzeExpMacroStat( token )
    else 
      exp = self:analyzeExp( true, op1levelMap[token.txt] or _luneScript.error( 'unwrap val is nil' ) )
      local typeInfo = builtinTypeNone
      
      local macroExpFlag = false
      
      do
        local _switchExp = (token.txt )
        if _switchExp == "-" then
          if exp:get_expType() ~= builtinTypeInt and exp:get_expType() ~= builtinTypeReal then
            self:addErrMess( token.pos, string.format( 'unmatch type for "-" -- %s', exp:get_expType():getTxt(  )) )
          end
          typeInfo = exp:get_expType()
        elseif _switchExp == "#" then
          if exp:get_expType():get_kind() ~= TypeInfoKindList and exp:get_expType():get_kind() ~= TypeInfoKindArray and exp:get_expType():get_kind() ~= TypeInfoKindMap and exp:get_expType() ~= builtinTypeString then
            self:addErrMess( token.pos, string.format( 'unmatch type for "#" -- %s', exp:get_expType():getTxt(  )) )
          end
          typeInfo = builtinTypeInt
        elseif _switchExp == "not" then
          typeInfo = builtinTypeBool
        elseif _switchExp == ",," then
          macroExpFlag = true
        elseif _switchExp == ",,," then
          macroExpFlag = true
          if exp:get_expType() ~= builtinTypeString then
            self:error( "unmatch ,,, type, need string type" )
          end
          typeInfo = builtinTypeSymbol
        elseif _switchExp == ",,,," then
          macroExpFlag = true
          if exp:get_expType() ~= builtinTypeSymbol then
            self:error( "unmatch ,,, type, need symbol type" )
          end
          typeInfo = builtinTypeString
        elseif _switchExp == "`" then
          typeInfo = builtinTypeNone
        elseif _switchExp == "not" then
          typeInfo = builtinTypeBool
        else 
          self:error( "unknown op1" )
        end
      end
      
      if macroExpFlag then
        local nextToken = self:getToken(  )
        
        if nextToken.txt ~= "$" then
          self:pushback(  )
        end
      end
      exp = ExpOp1Node.new(firstToken.pos, {typeInfo}, token, self.macroMode, exp)
      return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
    end
  end
  if token.kind == Parser.kind.Int then
    exp = LiteralIntNode.new(firstToken.pos, {builtinTypeInt}, token, math.floor(tonumber( token.txt )))
  elseif token.kind == Parser.kind.Real then
    exp = LiteralRealNode.new(firstToken.pos, {builtinTypeReal}, token, tonumber( token.txt ))
  elseif token.kind == Parser.kind.Char then
    local num = 0
    
    if #(token.txt ) == 1 then
      num = token.txt:byte( 1 )
    else 
      num = quotedChar2Code[token.txt:sub( 2, 2 )] or _luneScript.error( 'unwrap val is nil' )
    end
    exp = LiteralCharNode.new(firstToken.pos, {builtinTypeChar}, token, num)
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
    exp = LiteralStringNode.new(firstToken.pos, {builtinTypeString}, token, formatArgList)
    token = nextToken
    if token.txt == "[" then
      exp = self:analyzeExpRefItem( token, exp )
    else 
      self:pushback(  )
    end
  elseif token.kind == Parser.kind.Kywd and token.txt == "fn" then
    exp = self:analyzeExpSymbol( firstToken, token, "fn", nil, false )
  elseif token.kind == Parser.kind.Kywd and token.txt == "unwrap" then
    exp = self:analyzeExpUnwrap( token )
  elseif token.kind == Parser.kind.Symb then
    exp = self:analyzeExpSymbol( firstToken, token, "symbol", nil, false )
  elseif token.kind == Parser.kind.Type then
    exp = ExpRefNode.new(firstToken.pos, {builtinTypeNone}, token)
  elseif token.kind == Parser.kind.Kywd and (token.txt == "true" or token.txt == "false" ) then
    exp = LiteralBoolNode.new(firstToken.pos, {builtinTypeBool}, token)
  elseif token.kind == Parser.kind.Kywd and (token.txt == "nil" or token.txt == "null" ) then
    exp = LiteralNilNode.new(firstToken.pos, {builtinTypeNil})
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
  
      if  not token then
        local _token = token
        
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
      
      local nextToken = token
      
      if token.txt ~= "static" then
        nextToken = self:getToken(  )
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
      local funcTypeInfo = self:getCurrentNamespaceTypeInfo(  )
      
      if not funcTypeInfo then
        self:addErrMess( token.pos, "'return' could not exist here" )
      else 
        do
          local _exp = expList
          if _exp then
          
              local retTypeList = funcTypeInfo:get_retTypeInfoList()
              
              local expNodeList = _exp:get_expList()
              
              for index, retType in pairs( retTypeList ) do
                local expNode = expNodeList[index]
                
                if expNode then
                  local expType = expNode:get_expType()
                  
                  if not retType:isSettableFrom( expType ) then
                    self:addErrMess( token.pos, string.format( "return type of arg(%d) is not compatible -- %s(%d) and %s(%d)", index, retType:getTxt(  ), retType:get_typeId(  ), expType:getTxt(  ), expType:get_typeId(  )) )
                  end
                end
              end
            end
        end
        
      end
      statement = ReturnNode.new(token.pos, {builtinTypeNone}, expList)
    elseif token.txt == "break" then
      self:checkNextToken( ";" )
      statement = BreakNode.new(token.pos, {builtinTypeNone})
    elseif token.txt == "unwrap" then
      statement = self:analyzeUnwrap( token )
    elseif token.txt == "sync" then
      statement = self:analyzeDeclVar( "sync", "pri", false, token )
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
      statement = StmtExpNode.new(exp.pos, {builtinTypeNone}, exp)
    end
  end
  return statement
end

function TransUnit:analyzeStatementList( stmtList, termTxt )
  while true do
    local statement = self:analyzeStatement( termTxt )
    
    do
      local _exp = statement
      if _exp then
      
          table.insert( stmtList, _exp )
        else
      
          break
        end
    end
    
  end
end

----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo3262 = {}
  _className2InfoMap.ASTInfo = _classInfo3262
  _typeId2ClassInfoMap[ 3262 ] = _classInfo3262
  end
do
  local _classInfo1544 = {}
  _className2InfoMap.ApplyNode = _classInfo1544
  _typeId2ClassInfoMap[ 1544 ] = _classInfo1544
  end
do
  local _classInfo1326 = {}
  _className2InfoMap.BlockNode = _classInfo1326
  _typeId2ClassInfoMap[ 1326 ] = _classInfo1326
  end
do
  local _classInfo1658 = {}
  _className2InfoMap.BreakNode = _classInfo1658
  _typeId2ClassInfoMap[ 1658 ] = _classInfo1658
  end
do
  local _classInfo1410 = {}
  _className2InfoMap.CaseInfo = _classInfo1410
  _typeId2ClassInfoMap[ 1410 ] = _classInfo1410
  end
do
  local _classInfo2346 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo2346
  _typeId2ClassInfoMap[ 2346 ] = _classInfo2346
  end
do
  local _classInfo2328 = {}
  _className2InfoMap.DeclArgNode = _classInfo2328
  _typeId2ClassInfoMap[ 2328 ] = _classInfo2328
  end
do
  local _classInfo844 = {}
  _className2InfoMap.DeclClassNode = _classInfo844
  _typeId2ClassInfoMap[ 844 ] = _classInfo844
  end
do
  local _classInfo2242 = {}
  _className2InfoMap.DeclConstrNode = _classInfo2242
  _typeId2ClassInfoMap[ 2242 ] = _classInfo2242
  end
do
  local _classInfo2172 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo2172
  _typeId2ClassInfoMap[ 2172 ] = _classInfo2172
  end
do
  local _classInfo2202 = {}
  _className2InfoMap.DeclFuncNode = _classInfo2202
  _typeId2ClassInfoMap[ 2202 ] = _classInfo2202
  end
do
  local _classInfo1018 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo1018
  _typeId2ClassInfoMap[ 1018 ] = _classInfo1018
  end
do
  local _classInfo2420 = {}
  _className2InfoMap.DeclMacroNode = _classInfo2420
  _typeId2ClassInfoMap[ 2420 ] = _classInfo2420
  end
do
  local _classInfo2296 = {}
  _className2InfoMap.DeclMemberNode = _classInfo2296
  _typeId2ClassInfoMap[ 2296 ] = _classInfo2296
  end
do
  local _classInfo2222 = {}
  _className2InfoMap.DeclMethodNode = _classInfo2222
  _typeId2ClassInfoMap[ 2222 ] = _classInfo2222
  end
do
  local _classInfo2118 = {}
  _className2InfoMap.DeclVarNode = _classInfo2118
  _typeId2ClassInfoMap[ 2118 ] = _classInfo2118
  end
do
  local _classInfo1902 = {}
  _className2InfoMap.ExpCallNode = _classInfo1902
  _typeId2ClassInfoMap[ 1902 ] = _classInfo1902
  end
do
  local _classInfo2264 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo2264
  _typeId2ClassInfoMap[ 2264 ] = _classInfo2264
  end
do
  local _classInfo1828 = {}
  _className2InfoMap.ExpCastNode = _classInfo1828
  _typeId2ClassInfoMap[ 1828 ] = _classInfo1828
  end
do
  local _classInfo1924 = {}
  _className2InfoMap.ExpDDDNode = _classInfo1924
  _typeId2ClassInfoMap[ 1924 ] = _classInfo1924
  end
do
  local _classInfo1016 = {}
  _className2InfoMap.ExpListNode = _classInfo1016
  _typeId2ClassInfoMap[ 1016 ] = _classInfo1016
  end
do
  local _classInfo1964 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo1964
  _typeId2ClassInfoMap[ 1964 ] = _classInfo1964
  end
do
  local _classInfo1990 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo1990
  _typeId2ClassInfoMap[ 1990 ] = _classInfo1990
  end
do
  local _classInfo1678 = {}
  _className2InfoMap.ExpNewNode = _classInfo1678
  _typeId2ClassInfoMap[ 1678 ] = _classInfo1678
  end
do
  local _classInfo1852 = {}
  _className2InfoMap.ExpOp1Node = _classInfo1852
  _typeId2ClassInfoMap[ 1852 ] = _classInfo1852
  end
do
  local _classInfo1748 = {}
  _className2InfoMap.ExpOp2Node = _classInfo1748
  _typeId2ClassInfoMap[ 1748 ] = _classInfo1748
  end
do
  local _classInfo1944 = {}
  _className2InfoMap.ExpParenNode = _classInfo1944
  _typeId2ClassInfoMap[ 1944 ] = _classInfo1944
  end
do
  local _classInfo1878 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo1878
  _typeId2ClassInfoMap[ 1878 ] = _classInfo1878
  end
do
  local _classInfo1724 = {}
  _className2InfoMap.ExpRefNode = _classInfo1724
  _typeId2ClassInfoMap[ 1724 ] = _classInfo1724
  end
do
  local _classInfo1702 = {}
  _className2InfoMap.ExpUnwrapNode = _classInfo1702
  _typeId2ClassInfoMap[ 1702 ] = _classInfo1702
  end
do
  local _classInfo982 = {}
  _className2InfoMap.Filter = _classInfo982
  _typeId2ClassInfoMap[ 982 ] = _classInfo982
  end
do
  local _classInfo1512 = {}
  _className2InfoMap.ForNode = _classInfo1512
  _typeId2ClassInfoMap[ 1512 ] = _classInfo1512
  end
do
  local _classInfo1580 = {}
  _className2InfoMap.ForeachNode = _classInfo1580
  _typeId2ClassInfoMap[ 1580 ] = _classInfo1580
  end
do
  local _classInfo1614 = {}
  _className2InfoMap.ForsortNode = _classInfo1614
  _typeId2ClassInfoMap[ 1614 ] = _classInfo1614
  end
do
  local _classInfo2064 = {}
  _className2InfoMap.GetFieldNode = _classInfo2064
  _typeId2ClassInfoMap[ 2064 ] = _classInfo2064
  end
do
  local _classInfo1364 = {}
  _className2InfoMap.IfNode = _classInfo1364
  _typeId2ClassInfoMap[ 1364 ] = _classInfo1364
  end
do
  local _classInfo1350 = {}
  _className2InfoMap.IfStmtInfo = _classInfo1350
  _typeId2ClassInfoMap[ 1350 ] = _classInfo1350
  end
do
  local _classInfo1804 = {}
  _className2InfoMap.IfUnwrapNode = _classInfo1804
  _typeId2ClassInfoMap[ 1804 ] = _classInfo1804
  end
do
  local _classInfo1236 = {}
  _className2InfoMap.ImportNode = _classInfo1236
  _typeId2ClassInfoMap[ 1236 ] = _classInfo1236
  end
do
  local _classInfo2526 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo2526
  _typeId2ClassInfoMap[ 2526 ] = _classInfo2526
  end
do
  local _classInfo2640 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo2640
  _typeId2ClassInfoMap[ 2640 ] = _classInfo2640
  end
do
  local _classInfo2456 = {}
  _className2InfoMap.LiteralCharNode = _classInfo2456
  _typeId2ClassInfoMap[ 2456 ] = _classInfo2456
  end
do
  local _classInfo2480 = {}
  _className2InfoMap.LiteralIntNode = _classInfo2480
  _typeId2ClassInfoMap[ 2480 ] = _classInfo2480
  end
do
  local _classInfo2546 = {}
  _className2InfoMap.LiteralListNode = _classInfo2546
  _typeId2ClassInfoMap[ 2546 ] = _classInfo2546
  end
do
  local _classInfo2576 = {}
  _className2InfoMap.LiteralMapNode = _classInfo2576
  _typeId2ClassInfoMap[ 2576 ] = _classInfo2576
  end
do
  local _classInfo2436 = {}
  _className2InfoMap.LiteralNilNode = _classInfo2436
  _typeId2ClassInfoMap[ 2436 ] = _classInfo2436
  end
do
  local _classInfo2504 = {}
  _className2InfoMap.LiteralRealNode = _classInfo2504
  _typeId2ClassInfoMap[ 2504 ] = _classInfo2504
  end
do
  local _classInfo2612 = {}
  _className2InfoMap.LiteralStringNode = _classInfo2612
  _typeId2ClassInfoMap[ 2612 ] = _classInfo2612
  end
do
  local _classInfo2660 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo2660
  _typeId2ClassInfoMap[ 2660 ] = _classInfo2660
  end
do
  local _classInfo1014 = {}
  _className2InfoMap.MacroEval = _classInfo1014
  _typeId2ClassInfoMap[ 1014 ] = _classInfo1014
  end
do
  local _classInfo1010 = {}
  _className2InfoMap.NamespaceInfo = _classInfo1010
  _typeId2ClassInfoMap[ 1010 ] = _classInfo1010
  _classInfo1010.name = {
    name='name', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 18 }
  _classInfo1010.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 734 }
  _classInfo1010.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 708 }
  end
do
  local _classInfo842 = {}
  _className2InfoMap.Node = _classInfo842
  _typeId2ClassInfoMap[ 842 ] = _classInfo842
  end
do
  local _classInfo1218 = {}
  _className2InfoMap.NoneNode = _classInfo1218
  _typeId2ClassInfoMap[ 1218 ] = _classInfo1218
  end
do
  local _classInfo846 = {}
  _className2InfoMap.NormalTypeInfo = _classInfo846
  _typeId2ClassInfoMap[ 846 ] = _classInfo846
  end
do
  local _classInfo724 = {}
  _className2InfoMap.OutStream = _classInfo724
  _typeId2ClassInfoMap[ 724 ] = _classInfo724
  end
do
  local _classInfo2562 = {}
  _className2InfoMap.PairItem = _classInfo2562
  _typeId2ClassInfoMap[ 2562 ] = _classInfo2562
  end
do
  local _classInfo2038 = {}
  _className2InfoMap.RefFieldNode = _classInfo2038
  _typeId2ClassInfoMap[ 2038 ] = _classInfo2038
  end
do
  local _classInfo1298 = {}
  _className2InfoMap.RefTypeNode = _classInfo1298
  _typeId2ClassInfoMap[ 1298 ] = _classInfo1298
  end
do
  local _classInfo1482 = {}
  _className2InfoMap.RepeatNode = _classInfo1482
  _typeId2ClassInfoMap[ 1482 ] = _classInfo1482
  end
do
  local _classInfo1642 = {}
  _className2InfoMap.ReturnNode = _classInfo1642
  _typeId2ClassInfoMap[ 1642 ] = _classInfo1642
  end
do
  local _classInfo1258 = {}
  _className2InfoMap.RootNode = _classInfo1258
  _typeId2ClassInfoMap[ 1258 ] = _classInfo1258
  end
do
  local _classInfo734 = {}
  _className2InfoMap.Scope = _classInfo734
  _typeId2ClassInfoMap[ 734 ] = _classInfo734
  end
do
  local _classInfo2016 = {}
  _className2InfoMap.StmtExpNode = _classInfo2016
  _typeId2ClassInfoMap[ 2016 ] = _classInfo2016
  end
do
  local _classInfo1426 = {}
  _className2InfoMap.SwitchNode = _classInfo1426
  _typeId2ClassInfoMap[ 1426 ] = _classInfo1426
  end
do
  local _classInfo1048 = {}
  _className2InfoMap.TransUnit = _classInfo1048
  _typeId2ClassInfoMap[ 1048 ] = _classInfo1048
  end
do
  local _classInfo708 = {}
  _className2InfoMap.TypeInfo = _classInfo708
  _typeId2ClassInfoMap[ 708 ] = _classInfo708
  end
do
  local _classInfo1776 = {}
  _className2InfoMap.UnwrapSetNode = _classInfo1776
  _typeId2ClassInfoMap[ 1776 ] = _classInfo1776
  end
do
  local _classInfo2084 = {}
  _className2InfoMap.VarInfo = _classInfo2084
  _typeId2ClassInfoMap[ 2084 ] = _classInfo2084
  end
do
  local _classInfo1458 = {}
  _className2InfoMap.WhileNode = _classInfo1458
  _typeId2ClassInfoMap[ 1458 ] = _classInfo1458
  end
do
  local _classInfo636 = {}
  _className2InfoMap.Parser = _classInfo636
  _typeId2ClassInfoMap[ 636 ] = _classInfo636
  end
do
  local _classInfo372 = {}
  _className2InfoMap.Position = _classInfo372
  _typeId2ClassInfoMap[ 372 ] = _classInfo372
  _classInfo372.column = {
    name='column', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  _classInfo372.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 12 }
  end
do
  local _classInfo606 = {}
  _className2InfoMap.Stream = _classInfo606
  _typeId2ClassInfoMap[ 606 ] = _classInfo606
  end
do
  local _classInfo406 = {}
  _className2InfoMap.StreamParser = _classInfo406
  _typeId2ClassInfoMap[ 406 ] = _classInfo406
  _classInfo406.create = {
    name='create', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 558 }
  end
do
  local _classInfo622 = {}
  _className2InfoMap.Token = _classInfo622
  _typeId2ClassInfoMap[ 622 ] = _classInfo622
  end
do
  local _classInfo366 = {}
  _className2InfoMap.TxtStream = _classInfo366
  _typeId2ClassInfoMap[ 366 ] = _classInfo366
  end
do
  local _classInfo442 = {}
  _className2InfoMap.Util = _classInfo442
  _typeId2ClassInfoMap[ 442 ] = _classInfo442
  _classInfo442.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 704 }
  _classInfo442.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 702 }
  _classInfo442.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 694 }
  _classInfo442.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 688 }
  _classInfo442.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', methodFlag = false, typeId = 706 }
  end
do
  local _classInfo398 = {}
  _className2InfoMap.WrapParser = _classInfo398
  _typeId2ClassInfoMap[ 398 ] = _classInfo398
  end
do
  local _classInfo110 = {}
  _className2InfoMap.base = _classInfo110
  _typeId2ClassInfoMap[ 110 ] = _classInfo110
  _classInfo110.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 112 }
  _classInfo110.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 340 }
  _classInfo110.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 442 }
  end
do
  local _classInfo108 = {}
  _className2InfoMap.lune = _classInfo108
  _typeId2ClassInfoMap[ 108 ] = _classInfo108
  _classInfo108.base = {
    name='base', staticFlag = false, accessMode = 'pub', methodFlag = false, typeId = 110 }
  end
do
  local _classInfo694 = {}
  _className2InfoMap.memStream = _classInfo694
  _typeId2ClassInfoMap[ 694 ] = _classInfo694
  end
do
  local _classInfo688 = {}
  _className2InfoMap.outStream = _classInfo688
  _typeId2ClassInfoMap[ 688 ] = _classInfo688
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
_varName2InfoMap.TypeInfoKindNilable = {
  name='TypeInfoKindNilable', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindPrim = {
  name='TypeInfoKindPrim', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.TypeInfoKindRoot = {
  name='TypeInfoKindRoot', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.builtinTypeArray = {
  name='builtinTypeArray', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeBool = {
  name='builtinTypeBool', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeChar = {
  name='builtinTypeChar', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeDDD = {
  name='builtinTypeDDD', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeForm = {
  name='builtinTypeForm', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeInt = {
  name='builtinTypeInt', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeList = {
  name='builtinTypeList', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeMap = {
  name='builtinTypeMap', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeNil = {
  name='builtinTypeNil', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeNone = {
  name='builtinTypeNone', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeReal = {
  name='builtinTypeReal', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeStat = {
  name='builtinTypeStat', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeStem = {
  name='builtinTypeStem', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeString = {
  name='builtinTypeString', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.builtinTypeSymbol = {
  name='builtinTypeSymbol', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.nodeKind = {
  name='nodeKind', accessMode = 'pub', typeId = 1188 }
_varName2InfoMap.rootTypeId = {
  name='rootTypeId', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.typeInfoKind = {
  name='typeInfoKind', accessMode = 'pub', typeId = 710 }
_varName2InfoMap.typeInfoListInsert = {
  name='typeInfoListInsert', accessMode = 'pub', typeId = 708 }
_varName2InfoMap.typeInfoListRemove = {
  name='typeInfoListRemove', accessMode = 'pub', typeId = 708 }
local _typeInfoList = {}
moduleObj._typeInfoList = _typeInfoList
_typeInfoList[1] = { parentId = 1, typeId = 9, nilable = true, orgTypeId = 8 }
_typeInfoList[2] = { parentId = 1, typeId = 11, nilable = true, orgTypeId = 10 }
_typeInfoList[3] = { parentId = 1, typeId = 13, nilable = true, orgTypeId = 12 }
_typeInfoList[4] = { parentId = 1, typeId = 15, nilable = true, orgTypeId = 14 }
_typeInfoList[5] = { parentId = 1, typeId = 19, nilable = true, orgTypeId = 18 }
_typeInfoList[6] = { parentId = 1, typeId = 102, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {104}, }
_typeInfoList[7] = { parentId = 1, typeId = 103, nilable = true, orgTypeId = 102 }
_typeInfoList[8] = { parentId = 102, typeId = 104, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106, 154, 484}, }
_typeInfoList[9] = { parentId = 102, typeId = 105, nilable = true, orgTypeId = 104 }
_typeInfoList[10] = { parentId = 104, typeId = 106, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {108, 708, 722, 724, 734, 842, 844, 846, 982, 1010, 1014, 1016, 1018, 1048, 1194, 1218, 1236, 1258, 1298, 1326, 1350, 1364, 1410, 1426, 1458, 1482, 1512, 1544, 1580, 1614, 1642, 1658, 1678, 1702, 1724, 1748, 1776, 1804, 1828, 1852, 1878, 1902, 1924, 1944, 1964, 1990, 2016, 2038, 2064, 2084, 2118, 2172, 2202, 2222, 2242, 2264, 2296, 2328, 2346, 2420, 2436, 2456, 2480, 2504, 2526, 2546, 2562, 2576, 2612, 2640, 2660, 3262}, }
_typeInfoList[11] = { parentId = 104, typeId = 107, nilable = true, orgTypeId = 106 }
_typeInfoList[12] = { parentId = 106, typeId = 108, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {110}, }
_typeInfoList[13] = { parentId = 106, typeId = 109, nilable = true, orgTypeId = 108 }
_typeInfoList[14] = { parentId = 108, typeId = 110, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {112, 340, 442}, }
_typeInfoList[15] = { parentId = 108, typeId = 111, nilable = true, orgTypeId = 110 }
_typeInfoList[16] = { parentId = 110, typeId = 112, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {358, 360, 366, 372, 376, 390, 398, 406, 416, 418, 420, 422, 440, 518, 560, 562, 564, 566, 584}, }
_typeInfoList[17] = { parentId = 110, typeId = 113, nilable = true, orgTypeId = 112 }
_typeInfoList[18] = { parentId = 104, typeId = 154, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {210, 216, 222, 226, 240, 248, 256, 310, 312, 314, 338}, }
_typeInfoList[19] = { parentId = 104, typeId = 155, nilable = true, orgTypeId = 154 }
_typeInfoList[20] = { parentId = 1, typeId = 156, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pri', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[21] = { parentId = 1, typeId = 157, nilable = true, orgTypeId = 156 }
_typeInfoList[22] = { parentId = 1, typeId = 158, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pri', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[23] = { parentId = 1, typeId = 159, nilable = true, orgTypeId = 158 }
_typeInfoList[24] = { parentId = 1, typeId = 160, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pri', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[25] = { parentId = 1, typeId = 161, nilable = true, orgTypeId = 160 }
_typeInfoList[26] = { parentId = 1, typeId = 162, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[27] = { parentId = 1, typeId = 163, nilable = true, orgTypeId = 162 }
_typeInfoList[28] = { parentId = 1, typeId = 164, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pri', kind = 5, itemTypeId = {18, 162}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[29] = { parentId = 1, typeId = 165, nilable = true, orgTypeId = 164 }
_typeInfoList[30] = { parentId = 154, typeId = 166, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {156, 158, 160, 164}, children = {}, }
_typeInfoList[31] = { parentId = 154, typeId = 210, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {212, 214}, }
_typeInfoList[32] = { parentId = 154, typeId = 211, nilable = true, orgTypeId = 210 }
_typeInfoList[33] = { parentId = 210, typeId = 212, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[34] = { parentId = 210, typeId = 214, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[35] = { parentId = 154, typeId = 216, baseId = 210, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {218, 220}, }
_typeInfoList[36] = { parentId = 154, typeId = 217, nilable = true, orgTypeId = 216 }
_typeInfoList[37] = { parentId = 216, typeId = 218, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[38] = { parentId = 216, typeId = 220, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[39] = { parentId = 154, typeId = 222, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {224}, }
_typeInfoList[40] = { parentId = 154, typeId = 223, nilable = true, orgTypeId = 222 }
_typeInfoList[41] = { parentId = 222, typeId = 224, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 154, typeId = 226, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {232, 236, 238}, }
_typeInfoList[43] = { parentId = 154, typeId = 227, nilable = true, orgTypeId = 226 }
_typeInfoList[44] = { parentId = 1, typeId = 228, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {226}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 1, typeId = 229, nilable = true, orgTypeId = 228 }
_typeInfoList[46] = { parentId = 1, typeId = 230, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {226}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[47] = { parentId = 1, typeId = 231, nilable = true, orgTypeId = 230 }
_typeInfoList[48] = { parentId = 226, typeId = 232, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {230}, retTypeId = {}, children = {}, }
_typeInfoList[49] = { parentId = 1, typeId = 234, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {226}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[50] = { parentId = 1, typeId = 235, nilable = true, orgTypeId = 234 }
_typeInfoList[51] = { parentId = 226, typeId = 236, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {234}, children = {}, }
_typeInfoList[52] = { parentId = 226, typeId = 238, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 222, 228}, retTypeId = {}, children = {}, }
_typeInfoList[53] = { parentId = 154, typeId = 240, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {242, 244, 246}, }
_typeInfoList[54] = { parentId = 154, typeId = 241, nilable = true, orgTypeId = 240 }
_typeInfoList[55] = { parentId = 240, typeId = 242, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {226}, children = {}, }
_typeInfoList[56] = { parentId = 240, typeId = 244, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[57] = { parentId = 240, typeId = 246, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[58] = { parentId = 154, typeId = 248, baseId = 240, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {250, 252, 254}, }
_typeInfoList[59] = { parentId = 154, typeId = 249, nilable = true, orgTypeId = 248 }
_typeInfoList[60] = { parentId = 248, typeId = 250, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {226}, children = {}, }
_typeInfoList[61] = { parentId = 248, typeId = 252, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[62] = { parentId = 248, typeId = 254, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {240, 18}, retTypeId = {}, children = {}, }
_typeInfoList[63] = { parentId = 154, typeId = 256, baseId = 240, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {270, 274, 276, 334}, }
_typeInfoList[64] = { parentId = 154, typeId = 257, nilable = true, orgTypeId = 256 }
_typeInfoList[65] = { parentId = 256, typeId = 270, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {210, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[66] = { parentId = 256, typeId = 274, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[67] = { parentId = 256, typeId = 276, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {257}, children = {}, }
_typeInfoList[68] = { parentId = 154, typeId = 296, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[69] = { parentId = 154, typeId = 310, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[70] = { parentId = 154, typeId = 312, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[71] = { parentId = 154, typeId = 314, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[72] = { parentId = 1, typeId = 316, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {226}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[73] = { parentId = 1, typeId = 317, nilable = true, orgTypeId = 316 }
_typeInfoList[74] = { parentId = 256, typeId = 318, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {317}, children = {}, }
_typeInfoList[75] = { parentId = 318, typeId = 320, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[76] = { parentId = 318, typeId = 326, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[77] = { parentId = 318, typeId = 328, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {}, }
_typeInfoList[78] = { parentId = 328, typeId = 330, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {226}, children = {}, }
_typeInfoList[79] = { parentId = 328, typeId = 332, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[80] = { parentId = 256, typeId = 334, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {227}, children = {}, }
_typeInfoList[81] = { parentId = 154, typeId = 338, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {226}, children = {}, }
_typeInfoList[82] = { parentId = 110, typeId = 340, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {342}, }
_typeInfoList[83] = { parentId = 110, typeId = 341, nilable = true, orgTypeId = 340 }
_typeInfoList[84] = { parentId = 340, typeId = 342, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {344}, }
_typeInfoList[85] = { parentId = 340, typeId = 343, nilable = true, orgTypeId = 342 }
_typeInfoList[86] = { parentId = 342, typeId = 344, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {346, 586, 686}, }
_typeInfoList[87] = { parentId = 342, typeId = 345, nilable = true, orgTypeId = 344 }
_typeInfoList[88] = { parentId = 344, typeId = 346, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {604, 606, 612, 618, 622, 636, 644, 652, 660, 662, 664, 666, 684}, }
_typeInfoList[89] = { parentId = 344, typeId = 347, nilable = true, orgTypeId = 346 }
_typeInfoList[90] = { parentId = 1, typeId = 348, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[91] = { parentId = 1, typeId = 349, nilable = true, orgTypeId = 348 }
_typeInfoList[92] = { parentId = 1, typeId = 350, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 351, nilable = true, orgTypeId = 350 }
_typeInfoList[94] = { parentId = 1, typeId = 352, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[95] = { parentId = 1, typeId = 353, nilable = true, orgTypeId = 352 }
_typeInfoList[96] = { parentId = 1, typeId = 354, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[97] = { parentId = 1, typeId = 355, nilable = true, orgTypeId = 354 }
_typeInfoList[98] = { parentId = 1, typeId = 356, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 354}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 357, nilable = true, orgTypeId = 356 }
_typeInfoList[100] = { parentId = 112, typeId = 358, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {348, 350, 352, 356}, children = {}, }
_typeInfoList[101] = { parentId = 112, typeId = 360, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {362, 364, 520, 522}, }
_typeInfoList[102] = { parentId = 112, typeId = 361, nilable = true, orgTypeId = 360 }
_typeInfoList[103] = { parentId = 360, typeId = 362, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[104] = { parentId = 360, typeId = 364, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[105] = { parentId = 112, typeId = 366, baseId = 360, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {368, 370, 524, 526}, }
_typeInfoList[106] = { parentId = 112, typeId = 367, nilable = true, orgTypeId = 366 }
_typeInfoList[107] = { parentId = 366, typeId = 368, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[108] = { parentId = 366, typeId = 370, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[109] = { parentId = 112, typeId = 372, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {374, 528}, }
_typeInfoList[110] = { parentId = 112, typeId = 373, nilable = true, orgTypeId = 372 }
_typeInfoList[111] = { parentId = 372, typeId = 374, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[112] = { parentId = 112, typeId = 376, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {382, 386, 388, 534, 538, 540}, }
_typeInfoList[113] = { parentId = 112, typeId = 377, nilable = true, orgTypeId = 376 }
_typeInfoList[114] = { parentId = 1, typeId = 378, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 1, typeId = 379, nilable = true, orgTypeId = 378 }
_typeInfoList[116] = { parentId = 1, typeId = 380, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[117] = { parentId = 1, typeId = 381, nilable = true, orgTypeId = 380 }
_typeInfoList[118] = { parentId = 376, typeId = 382, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {380}, retTypeId = {}, children = {}, }
_typeInfoList[119] = { parentId = 1, typeId = 384, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[120] = { parentId = 1, typeId = 385, nilable = true, orgTypeId = 384 }
_typeInfoList[121] = { parentId = 376, typeId = 386, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {384}, children = {}, }
_typeInfoList[122] = { parentId = 376, typeId = 388, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 372, 378}, retTypeId = {}, children = {}, }
_typeInfoList[123] = { parentId = 112, typeId = 390, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {392, 394, 396, 542, 544, 546}, }
_typeInfoList[124] = { parentId = 112, typeId = 391, nilable = true, orgTypeId = 390 }
_typeInfoList[125] = { parentId = 390, typeId = 392, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[126] = { parentId = 390, typeId = 394, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[127] = { parentId = 390, typeId = 396, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[128] = { parentId = 112, typeId = 398, baseId = 390, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {400, 402, 404, 548, 550, 552}, }
_typeInfoList[129] = { parentId = 112, typeId = 399, nilable = true, orgTypeId = 398 }
_typeInfoList[130] = { parentId = 398, typeId = 400, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[131] = { parentId = 398, typeId = 402, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[132] = { parentId = 398, typeId = 404, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {390, 18}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 112, typeId = 406, baseId = 390, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {408, 410, 412, 426, 438, 554, 556, 558, 570, 582}, }
_typeInfoList[134] = { parentId = 112, typeId = 407, nilable = true, orgTypeId = 406 }
_typeInfoList[135] = { parentId = 406, typeId = 408, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {360, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 406, typeId = 410, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[137] = { parentId = 406, typeId = 412, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {407}, children = {}, }
_typeInfoList[138] = { parentId = 112, typeId = 416, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[139] = { parentId = 112, typeId = 418, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[140] = { parentId = 112, typeId = 420, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[141] = { parentId = 112, typeId = 422, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[142] = { parentId = 1, typeId = 424, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 425, nilable = true, orgTypeId = 424 }
_typeInfoList[144] = { parentId = 406, typeId = 426, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {425}, children = {428, 430, 432}, }
_typeInfoList[145] = { parentId = 426, typeId = 428, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[146] = { parentId = 426, typeId = 430, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[147] = { parentId = 426, typeId = 432, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {434, 436}, }
_typeInfoList[148] = { parentId = 432, typeId = 434, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {376}, children = {}, }
_typeInfoList[149] = { parentId = 432, typeId = 436, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[150] = { parentId = 406, typeId = 438, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[151] = { parentId = 112, typeId = 440, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[152] = { parentId = 110, typeId = 442, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {688, 694, 702, 704, 706}, }
_typeInfoList[153] = { parentId = 110, typeId = 443, nilable = true, orgTypeId = 442 }
_typeInfoList[154] = { parentId = 104, typeId = 484, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {486, 492, 500, 502, 506}, }
_typeInfoList[155] = { parentId = 104, typeId = 485, nilable = true, orgTypeId = 484 }
_typeInfoList[156] = { parentId = 484, typeId = 486, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {488, 490}, }
_typeInfoList[157] = { parentId = 484, typeId = 487, nilable = true, orgTypeId = 486 }
_typeInfoList[158] = { parentId = 486, typeId = 488, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[159] = { parentId = 486, typeId = 490, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 484, typeId = 492, baseId = 486, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {494, 496, 498}, }
_typeInfoList[161] = { parentId = 484, typeId = 493, nilable = true, orgTypeId = 492 }
_typeInfoList[162] = { parentId = 492, typeId = 494, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[163] = { parentId = 492, typeId = 496, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[164] = { parentId = 492, typeId = 498, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[165] = { parentId = 484, typeId = 500, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[166] = { parentId = 484, typeId = 502, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[167] = { parentId = 484, typeId = 506, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10, 6, 18}, retTypeId = {}, children = {}, }
_typeInfoList[168] = { parentId = 1, typeId = 508, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 1, typeId = 509, nilable = true, orgTypeId = 508 }
_typeInfoList[170] = { parentId = 1, typeId = 510, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[171] = { parentId = 1, typeId = 511, nilable = true, orgTypeId = 510 }
_typeInfoList[172] = { parentId = 1, typeId = 512, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[173] = { parentId = 1, typeId = 513, nilable = true, orgTypeId = 512 }
_typeInfoList[174] = { parentId = 1, typeId = 514, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[175] = { parentId = 1, typeId = 515, nilable = true, orgTypeId = 514 }
_typeInfoList[176] = { parentId = 1, typeId = 516, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 514}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[177] = { parentId = 1, typeId = 517, nilable = true, orgTypeId = 516 }
_typeInfoList[178] = { parentId = 112, typeId = 518, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {508, 510, 512, 516}, children = {}, }
_typeInfoList[179] = { parentId = 360, typeId = 520, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[180] = { parentId = 360, typeId = 522, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[181] = { parentId = 366, typeId = 524, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[182] = { parentId = 366, typeId = 526, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[183] = { parentId = 372, typeId = 528, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[184] = { parentId = 1, typeId = 530, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[185] = { parentId = 1, typeId = 531, nilable = true, orgTypeId = 530 }
_typeInfoList[186] = { parentId = 1, typeId = 532, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 1, typeId = 533, nilable = true, orgTypeId = 532 }
_typeInfoList[188] = { parentId = 376, typeId = 534, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {532}, retTypeId = {}, children = {}, }
_typeInfoList[189] = { parentId = 1, typeId = 536, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[190] = { parentId = 1, typeId = 537, nilable = true, orgTypeId = 536 }
_typeInfoList[191] = { parentId = 376, typeId = 538, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {536}, children = {}, }
_typeInfoList[192] = { parentId = 376, typeId = 540, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 372, 530}, retTypeId = {}, children = {}, }
_typeInfoList[193] = { parentId = 390, typeId = 542, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[194] = { parentId = 390, typeId = 544, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[195] = { parentId = 390, typeId = 546, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[196] = { parentId = 398, typeId = 548, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[197] = { parentId = 398, typeId = 550, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[198] = { parentId = 398, typeId = 552, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {390, 18}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 406, typeId = 554, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {360, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[200] = { parentId = 406, typeId = 556, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[201] = { parentId = 406, typeId = 558, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {407}, children = {}, }
_typeInfoList[202] = { parentId = 112, typeId = 560, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[203] = { parentId = 112, typeId = 562, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[204] = { parentId = 112, typeId = 564, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[205] = { parentId = 112, typeId = 566, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[206] = { parentId = 1, typeId = 568, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[207] = { parentId = 1, typeId = 569, nilable = true, orgTypeId = 568 }
_typeInfoList[208] = { parentId = 406, typeId = 570, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {569}, children = {572, 574, 576}, }
_typeInfoList[209] = { parentId = 570, typeId = 572, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[210] = { parentId = 570, typeId = 574, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[211] = { parentId = 570, typeId = 576, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {578, 580}, }
_typeInfoList[212] = { parentId = 576, typeId = 578, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {376}, children = {}, }
_typeInfoList[213] = { parentId = 576, typeId = 580, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[214] = { parentId = 406, typeId = 582, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[215] = { parentId = 112, typeId = 584, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[216] = { parentId = 344, typeId = 586, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {588}, }
_typeInfoList[217] = { parentId = 344, typeId = 587, nilable = true, orgTypeId = 586 }
_typeInfoList[218] = { parentId = 586, typeId = 588, baseId = 1, txt = 'lune',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {590}, }
_typeInfoList[219] = { parentId = 586, typeId = 589, nilable = true, orgTypeId = 588 }
_typeInfoList[220] = { parentId = 588, typeId = 590, baseId = 1, txt = 'base',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {592}, }
_typeInfoList[221] = { parentId = 588, typeId = 591, nilable = true, orgTypeId = 590 }
_typeInfoList[222] = { parentId = 590, typeId = 592, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[223] = { parentId = 590, typeId = 593, nilable = true, orgTypeId = 592 }
_typeInfoList[224] = { parentId = 1, typeId = 594, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 1, typeId = 595, nilable = true, orgTypeId = 594 }
_typeInfoList[226] = { parentId = 1, typeId = 596, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[227] = { parentId = 1, typeId = 597, nilable = true, orgTypeId = 596 }
_typeInfoList[228] = { parentId = 1, typeId = 598, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[229] = { parentId = 1, typeId = 599, nilable = true, orgTypeId = 598 }
_typeInfoList[230] = { parentId = 1, typeId = 600, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[231] = { parentId = 1, typeId = 601, nilable = true, orgTypeId = 600 }
_typeInfoList[232] = { parentId = 1, typeId = 602, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 600}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[233] = { parentId = 1, typeId = 603, nilable = true, orgTypeId = 602 }
_typeInfoList[234] = { parentId = 346, typeId = 604, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {594, 596, 598, 602}, children = {}, }
_typeInfoList[235] = { parentId = 346, typeId = 606, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {608, 610}, }
_typeInfoList[236] = { parentId = 346, typeId = 607, nilable = true, orgTypeId = 606 }
_typeInfoList[237] = { parentId = 606, typeId = 608, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[238] = { parentId = 606, typeId = 610, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[239] = { parentId = 346, typeId = 612, baseId = 606, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {614, 616}, }
_typeInfoList[240] = { parentId = 346, typeId = 613, nilable = true, orgTypeId = 612 }
_typeInfoList[241] = { parentId = 612, typeId = 614, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 612, typeId = 616, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[243] = { parentId = 346, typeId = 618, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {620}, }
_typeInfoList[244] = { parentId = 346, typeId = 619, nilable = true, orgTypeId = 618 }
_typeInfoList[245] = { parentId = 618, typeId = 620, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[246] = { parentId = 346, typeId = 622, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {628, 632, 634}, }
_typeInfoList[247] = { parentId = 346, typeId = 623, nilable = true, orgTypeId = 622 }
_typeInfoList[248] = { parentId = 1, typeId = 624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {622}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[249] = { parentId = 1, typeId = 625, nilable = true, orgTypeId = 624 }
_typeInfoList[250] = { parentId = 1, typeId = 626, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {622}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 1, typeId = 627, nilable = true, orgTypeId = 626 }
_typeInfoList[252] = { parentId = 622, typeId = 628, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {626}, retTypeId = {}, children = {}, }
_typeInfoList[253] = { parentId = 1, typeId = 630, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {622}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[254] = { parentId = 1, typeId = 631, nilable = true, orgTypeId = 630 }
_typeInfoList[255] = { parentId = 622, typeId = 632, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {630}, children = {}, }
_typeInfoList[256] = { parentId = 622, typeId = 634, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 618, 624}, retTypeId = {}, children = {}, }
_typeInfoList[257] = { parentId = 346, typeId = 636, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {638, 640, 642}, }
_typeInfoList[258] = { parentId = 346, typeId = 637, nilable = true, orgTypeId = 636 }
_typeInfoList[259] = { parentId = 636, typeId = 638, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {622}, children = {}, }
_typeInfoList[260] = { parentId = 636, typeId = 640, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[261] = { parentId = 636, typeId = 642, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[262] = { parentId = 346, typeId = 644, baseId = 636, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {646, 648, 650}, }
_typeInfoList[263] = { parentId = 346, typeId = 645, nilable = true, orgTypeId = 644 }
_typeInfoList[264] = { parentId = 644, typeId = 646, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {622}, children = {}, }
_typeInfoList[265] = { parentId = 644, typeId = 648, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[266] = { parentId = 644, typeId = 650, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {636, 18}, retTypeId = {}, children = {}, }
_typeInfoList[267] = { parentId = 346, typeId = 652, baseId = 636, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {654, 656, 658, 670, 682}, }
_typeInfoList[268] = { parentId = 346, typeId = 653, nilable = true, orgTypeId = 652 }
_typeInfoList[269] = { parentId = 652, typeId = 654, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {606, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[270] = { parentId = 652, typeId = 656, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[271] = { parentId = 652, typeId = 658, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {653}, children = {}, }
_typeInfoList[272] = { parentId = 346, typeId = 660, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[273] = { parentId = 346, typeId = 662, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[274] = { parentId = 346, typeId = 664, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[275] = { parentId = 346, typeId = 666, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[276] = { parentId = 1, typeId = 668, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {622}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[277] = { parentId = 1, typeId = 669, nilable = true, orgTypeId = 668 }
_typeInfoList[278] = { parentId = 652, typeId = 670, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {669}, children = {672, 674, 676}, }
_typeInfoList[279] = { parentId = 670, typeId = 672, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[280] = { parentId = 670, typeId = 674, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[281] = { parentId = 670, typeId = 676, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {678, 680}, }
_typeInfoList[282] = { parentId = 676, typeId = 678, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {622}, children = {}, }
_typeInfoList[283] = { parentId = 676, typeId = 680, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[284] = { parentId = 652, typeId = 682, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {623}, children = {}, }
_typeInfoList[285] = { parentId = 346, typeId = 684, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {622}, children = {}, }
_typeInfoList[286] = { parentId = 344, typeId = 686, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 344, typeId = 687, nilable = true, orgTypeId = 686 }
_typeInfoList[288] = { parentId = 442, typeId = 688, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {690, 692}, }
_typeInfoList[289] = { parentId = 442, typeId = 689, nilable = true, orgTypeId = 688 }
_typeInfoList[290] = { parentId = 688, typeId = 690, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[291] = { parentId = 688, typeId = 692, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[292] = { parentId = 442, typeId = 694, baseId = 688, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {696, 698, 700}, }
_typeInfoList[293] = { parentId = 442, typeId = 695, nilable = true, orgTypeId = 694 }
_typeInfoList[294] = { parentId = 694, typeId = 696, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[295] = { parentId = 694, typeId = 698, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[296] = { parentId = 694, typeId = 700, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[297] = { parentId = 442, typeId = 702, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[298] = { parentId = 442, typeId = 704, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[299] = { parentId = 442, typeId = 706, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10, 6, 18}, retTypeId = {}, children = {}, }
_typeInfoList[300] = { parentId = 106, typeId = 708, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {780, 782, 784, 786, 788, 790, 792, 794, 798, 802, 806, 808, 810, 812, 814, 816, 818, 820, 822, 824, 826, 828, 832, 836, 838}, }
_typeInfoList[301] = { parentId = 106, typeId = 709, nilable = true, orgTypeId = 708 }
_typeInfoList[302] = { parentId = 1, typeId = 710, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[303] = { parentId = 1, typeId = 711, nilable = true, orgTypeId = 710 }
_typeInfoList[304] = { parentId = 106, typeId = 722, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {6}, children = {}, }
_typeInfoList[305] = { parentId = 106, typeId = 724, baseId = 1, txt = 'OutStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {726, 728}, }
_typeInfoList[306] = { parentId = 106, typeId = 725, nilable = true, orgTypeId = 724 }
_typeInfoList[307] = { parentId = 724, typeId = 726, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[308] = { parentId = 724, typeId = 728, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[309] = { parentId = 106, typeId = 734, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {750, 752, 754, 756, 758, 760, 762, 764, 766, 770, 774}, }
_typeInfoList[310] = { parentId = 106, typeId = 735, nilable = true, orgTypeId = 734 }
_typeInfoList[311] = { parentId = 1, typeId = 742, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {734}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[312] = { parentId = 1, typeId = 743, nilable = true, orgTypeId = 742 }
_typeInfoList[313] = { parentId = 734, typeId = 744, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {735, 10, 742}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 734, typeId = 750, baseId = 1, txt = 'set_ownerTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {709}, retTypeId = {}, children = {}, }
_typeInfoList[315] = { parentId = 734, typeId = 752, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 708}, retTypeId = {}, children = {}, }
_typeInfoList[316] = { parentId = 734, typeId = 754, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 708, 734}, retTypeId = {}, children = {}, }
_typeInfoList[317] = { parentId = 734, typeId = 756, baseId = 1, txt = 'getClassScope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {735}, children = {}, }
_typeInfoList[318] = { parentId = 734, typeId = 758, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {709}, children = {}, }
_typeInfoList[319] = { parentId = 734, typeId = 760, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {709}, children = {}, }
_typeInfoList[320] = { parentId = 734, typeId = 762, baseId = 1, txt = 'getTypeInfoMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {709}, children = {}, }
_typeInfoList[321] = { parentId = 734, typeId = 764, baseId = 1, txt = 'get_ownerTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {709}, children = {}, }
_typeInfoList[322] = { parentId = 734, typeId = 766, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {734}, children = {}, }
_typeInfoList[323] = { parentId = 1, typeId = 768, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[324] = { parentId = 1, typeId = 769, nilable = true, orgTypeId = 768 }
_typeInfoList[325] = { parentId = 734, typeId = 770, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {768}, children = {}, }
_typeInfoList[326] = { parentId = 1, typeId = 772, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 734}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[327] = { parentId = 1, typeId = 773, nilable = true, orgTypeId = 772 }
_typeInfoList[328] = { parentId = 734, typeId = 774, baseId = 1, txt = 'get_className2ScopeMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {772}, children = {}, }
_typeInfoList[329] = { parentId = 708, typeId = 780, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {735}, retTypeId = {}, children = {}, }
_typeInfoList[330] = { parentId = 708, typeId = 782, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[331] = { parentId = 708, typeId = 784, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[332] = { parentId = 708, typeId = 786, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {708}, retTypeId = {10}, children = {}, }
_typeInfoList[333] = { parentId = 708, typeId = 788, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[334] = { parentId = 708, typeId = 790, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {724}, retTypeId = {}, children = {}, }
_typeInfoList[335] = { parentId = 708, typeId = 792, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {708, 12}, retTypeId = {10}, children = {}, }
_typeInfoList[336] = { parentId = 708, typeId = 794, baseId = 1, txt = 'get_externalFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[337] = { parentId = 1, typeId = 796, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[338] = { parentId = 1, typeId = 797, nilable = true, orgTypeId = 796 }
_typeInfoList[339] = { parentId = 708, typeId = 798, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {796}, children = {}, }
_typeInfoList[340] = { parentId = 1, typeId = 800, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[341] = { parentId = 1, typeId = 801, nilable = true, orgTypeId = 800 }
_typeInfoList[342] = { parentId = 708, typeId = 802, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {800}, children = {}, }
_typeInfoList[343] = { parentId = 1, typeId = 804, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[344] = { parentId = 1, typeId = 805, nilable = true, orgTypeId = 804 }
_typeInfoList[345] = { parentId = 708, typeId = 806, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {804}, children = {}, }
_typeInfoList[346] = { parentId = 708, typeId = 808, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[347] = { parentId = 708, typeId = 810, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[348] = { parentId = 708, typeId = 812, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[349] = { parentId = 708, typeId = 814, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[350] = { parentId = 708, typeId = 816, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[351] = { parentId = 708, typeId = 818, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[352] = { parentId = 708, typeId = 820, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[353] = { parentId = 708, typeId = 822, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[354] = { parentId = 708, typeId = 824, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[355] = { parentId = 708, typeId = 826, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[356] = { parentId = 708, typeId = 828, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[357] = { parentId = 1, typeId = 830, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[358] = { parentId = 1, typeId = 831, nilable = true, orgTypeId = 830 }
_typeInfoList[359] = { parentId = 708, typeId = 832, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {830}, children = {}, }
_typeInfoList[360] = { parentId = 1, typeId = 834, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[361] = { parentId = 1, typeId = 835, nilable = true, orgTypeId = 834 }
_typeInfoList[362] = { parentId = 708, typeId = 836, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {834}, children = {}, }
_typeInfoList[363] = { parentId = 708, typeId = 838, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {735}, children = {}, }
_typeInfoList[364] = { parentId = 106, typeId = 840, baseId = 1, txt = 'dumpScope',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {735, 18}, retTypeId = {}, children = {}, }
_typeInfoList[365] = { parentId = 106, typeId = 842, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {988, 994, 1000, 1002, 1006, 1008}, }
_typeInfoList[366] = { parentId = 106, typeId = 843, nilable = true, orgTypeId = 842 }
_typeInfoList[367] = { parentId = 106, typeId = 844, baseId = 842, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2382, 2396, 2398, 2400, 2404, 2408, 2410, 2414}, }
_typeInfoList[368] = { parentId = 106, typeId = 845, nilable = true, orgTypeId = 844 }
_typeInfoList[369] = { parentId = 106, typeId = 846, baseId = 708, txt = 'NormalTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {872, 874, 876, 878, 884, 886, 894, 898, 902, 906, 908, 910, 912, 914, 916, 918, 920, 922, 924, 926, 928, 932, 934, 954, 958, 960, 964, 970, 980}, }
_typeInfoList[370] = { parentId = 106, typeId = 847, nilable = true, orgTypeId = 846 }
_typeInfoList[371] = { parentId = 1, typeId = 856, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[372] = { parentId = 1, typeId = 857, nilable = true, orgTypeId = 856 }
_typeInfoList[373] = { parentId = 1, typeId = 858, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[374] = { parentId = 1, typeId = 859, nilable = true, orgTypeId = 858 }
_typeInfoList[375] = { parentId = 1, typeId = 860, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[376] = { parentId = 1, typeId = 861, nilable = true, orgTypeId = 860 }
_typeInfoList[377] = { parentId = 846, typeId = 862, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {735, 709, 709, 10, 10, 10, 18, 18, 709, 12, 12, 857, 859, 861}, retTypeId = {}, children = {}, }
_typeInfoList[378] = { parentId = 846, typeId = 872, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[379] = { parentId = 846, typeId = 874, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[380] = { parentId = 846, typeId = 876, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[381] = { parentId = 846, typeId = 878, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {724}, retTypeId = {}, children = {}, }
_typeInfoList[382] = { parentId = 1, typeId = 880, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[383] = { parentId = 1, typeId = 881, nilable = true, orgTypeId = 880 }
_typeInfoList[384] = { parentId = 878, typeId = 882, baseId = 1, txt = 'serializeTypeInfoList',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {18, 880, 10}, retTypeId = {18}, children = {}, }
_typeInfoList[385] = { parentId = 846, typeId = 884, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {708, 12}, retTypeId = {10}, children = {}, }
_typeInfoList[386] = { parentId = 846, typeId = 886, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {708}, retTypeId = {846}, children = {}, }
_typeInfoList[387] = { parentId = 1, typeId = 888, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[388] = { parentId = 1, typeId = 889, nilable = true, orgTypeId = 888 }
_typeInfoList[389] = { parentId = 1, typeId = 890, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 1, typeId = 891, nilable = true, orgTypeId = 890 }
_typeInfoList[391] = { parentId = 1, typeId = 892, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[392] = { parentId = 1, typeId = 893, nilable = true, orgTypeId = 892 }
_typeInfoList[393] = { parentId = 846, typeId = 894, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {735, 708, 708, 10, 12, 18, 888, 890, 892}, retTypeId = {708}, children = {}, }
_typeInfoList[394] = { parentId = 1, typeId = 896, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[395] = { parentId = 1, typeId = 897, nilable = true, orgTypeId = 896 }
_typeInfoList[396] = { parentId = 846, typeId = 898, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {896}, children = {}, }
_typeInfoList[397] = { parentId = 1, typeId = 900, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[398] = { parentId = 1, typeId = 901, nilable = true, orgTypeId = 900 }
_typeInfoList[399] = { parentId = 846, typeId = 902, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {900}, children = {}, }
_typeInfoList[400] = { parentId = 1, typeId = 904, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[401] = { parentId = 1, typeId = 905, nilable = true, orgTypeId = 904 }
_typeInfoList[402] = { parentId = 846, typeId = 906, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {904}, children = {}, }
_typeInfoList[403] = { parentId = 846, typeId = 908, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[404] = { parentId = 846, typeId = 910, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[405] = { parentId = 846, typeId = 912, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[406] = { parentId = 846, typeId = 914, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[407] = { parentId = 846, typeId = 916, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[408] = { parentId = 846, typeId = 918, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[409] = { parentId = 846, typeId = 920, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[410] = { parentId = 846, typeId = 922, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[411] = { parentId = 846, typeId = 924, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[412] = { parentId = 846, typeId = 926, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[413] = { parentId = 846, typeId = 928, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[414] = { parentId = 1, typeId = 930, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[415] = { parentId = 1, typeId = 931, nilable = true, orgTypeId = 930 }
_typeInfoList[416] = { parentId = 846, typeId = 932, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {930}, children = {}, }
_typeInfoList[417] = { parentId = 846, typeId = 934, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 18, 12, 708}, retTypeId = {708}, children = {}, }
_typeInfoList[418] = { parentId = 1, typeId = 952, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[419] = { parentId = 1, typeId = 953, nilable = true, orgTypeId = 952 }
_typeInfoList[420] = { parentId = 846, typeId = 954, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 708, 952}, retTypeId = {708}, children = {}, }
_typeInfoList[421] = { parentId = 1, typeId = 956, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[422] = { parentId = 1, typeId = 957, nilable = true, orgTypeId = 956 }
_typeInfoList[423] = { parentId = 846, typeId = 958, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 708, 956}, retTypeId = {708}, children = {}, }
_typeInfoList[424] = { parentId = 846, typeId = 960, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 708, 708, 708}, retTypeId = {708}, children = {}, }
_typeInfoList[425] = { parentId = 846, typeId = 964, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {735, 709, 708, 10, 18, 18}, retTypeId = {708}, children = {}, }
_typeInfoList[426] = { parentId = 1, typeId = 966, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[427] = { parentId = 1, typeId = 967, nilable = true, orgTypeId = 966 }
_typeInfoList[428] = { parentId = 1, typeId = 968, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[429] = { parentId = 1, typeId = 969, nilable = true, orgTypeId = 968 }
_typeInfoList[430] = { parentId = 846, typeId = 970, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {735, 12, 708, 10, 10, 10, 18, 18, 967, 969}, retTypeId = {708}, children = {}, }
_typeInfoList[431] = { parentId = 846, typeId = 980, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {708}, retTypeId = {10}, children = {}, }
_typeInfoList[432] = { parentId = 106, typeId = 982, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {984, 1220, 1238, 1260, 1300, 1328, 1366, 1390, 1428, 1460, 1484, 1514, 1546, 1582, 1616, 1644, 1660, 1680, 1704, 1726, 1750, 1778, 1806, 1830, 1854, 1880, 1904, 1926, 1946, 1966, 1992, 2018, 2040, 2066, 2120, 2204, 2224, 2244, 2266, 2298, 2330, 2348, 2374, 2422, 2438, 2458, 2482, 2506, 2528, 2548, 2578, 2614, 2642, 2662}, }
_typeInfoList[433] = { parentId = 106, typeId = 983, nilable = true, orgTypeId = 982 }
_typeInfoList[434] = { parentId = 982, typeId = 984, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[435] = { parentId = 1, typeId = 986, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[436] = { parentId = 1, typeId = 987, nilable = true, orgTypeId = 986 }
_typeInfoList[437] = { parentId = 842, typeId = 988, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[438] = { parentId = 1, typeId = 990, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[439] = { parentId = 1, typeId = 991, nilable = true, orgTypeId = 990 }
_typeInfoList[440] = { parentId = 1, typeId = 992, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[441] = { parentId = 1, typeId = 993, nilable = true, orgTypeId = 992 }
_typeInfoList[442] = { parentId = 842, typeId = 994, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {990, 992}, children = {}, }
_typeInfoList[443] = { parentId = 842, typeId = 1000, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[444] = { parentId = 842, typeId = 1002, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[445] = { parentId = 1, typeId = 1004, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[446] = { parentId = 1, typeId = 1005, nilable = true, orgTypeId = 1004 }
_typeInfoList[447] = { parentId = 842, typeId = 1006, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1004}, children = {}, }
_typeInfoList[448] = { parentId = 842, typeId = 1008, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 372, 986}, retTypeId = {}, children = {}, }
_typeInfoList[449] = { parentId = 106, typeId = 1010, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1012}, }
_typeInfoList[450] = { parentId = 106, typeId = 1011, nilable = true, orgTypeId = 1010 }
_typeInfoList[451] = { parentId = 1010, typeId = 1012, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 734, 708}, retTypeId = {}, children = {}, }
_typeInfoList[452] = { parentId = 106, typeId = 1014, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2820, 2822}, }
_typeInfoList[453] = { parentId = 106, typeId = 1015, nilable = true, orgTypeId = 1014 }
_typeInfoList[454] = { parentId = 106, typeId = 1016, baseId = 842, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1394, 1404, 1408}, }
_typeInfoList[455] = { parentId = 106, typeId = 1017, nilable = true, orgTypeId = 1016 }
_typeInfoList[456] = { parentId = 106, typeId = 1018, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1024, 1028, 1030, 1034, 1036}, }
_typeInfoList[457] = { parentId = 106, typeId = 1019, nilable = true, orgTypeId = 1018 }
_typeInfoList[458] = { parentId = 1, typeId = 1020, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[459] = { parentId = 1, typeId = 1021, nilable = true, orgTypeId = 1020 }
_typeInfoList[460] = { parentId = 1, typeId = 1022, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[461] = { parentId = 1, typeId = 1023, nilable = true, orgTypeId = 1022 }
_typeInfoList[462] = { parentId = 1018, typeId = 1024, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[463] = { parentId = 1, typeId = 1026, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[464] = { parentId = 1, typeId = 1027, nilable = true, orgTypeId = 1026 }
_typeInfoList[465] = { parentId = 1018, typeId = 1028, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1026}, children = {}, }
_typeInfoList[466] = { parentId = 1018, typeId = 1030, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {843}, children = {}, }
_typeInfoList[467] = { parentId = 1, typeId = 1032, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[468] = { parentId = 1, typeId = 1033, nilable = true, orgTypeId = 1032 }
_typeInfoList[469] = { parentId = 1018, typeId = 1034, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1032}, children = {}, }
_typeInfoList[470] = { parentId = 1018, typeId = 1036, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {376, 1020, 843, 1022}, retTypeId = {}, children = {}, }
_typeInfoList[471] = { parentId = 106, typeId = 1038, baseId = 1, txt = 'MacroValInfo',
        staticFlag = false, accessMode = 'pri', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1040}, }
_typeInfoList[472] = { parentId = 106, typeId = 1039, nilable = true, orgTypeId = 1038 }
_typeInfoList[473] = { parentId = 1038, typeId = 1040, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {6, 708}, retTypeId = {}, children = {}, }
_typeInfoList[474] = { parentId = 106, typeId = 1042, baseId = 1, txt = 'MacroInfo',
        staticFlag = false, accessMode = 'pri', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1046}, }
_typeInfoList[475] = { parentId = 106, typeId = 1043, nilable = true, orgTypeId = 1042 }
_typeInfoList[476] = { parentId = 1, typeId = 1044, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 1038}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[477] = { parentId = 1, typeId = 1045, nilable = true, orgTypeId = 1044 }
_typeInfoList[478] = { parentId = 1042, typeId = 1046, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {26, 1018, 1044}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 106, typeId = 1048, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1138, 3270}, }
_typeInfoList[480] = { parentId = 106, typeId = 1049, nilable = true, orgTypeId = 1048 }
_typeInfoList[481] = { parentId = 1048, typeId = 1066, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {1014}, retTypeId = {}, children = {}, }
_typeInfoList[482] = { parentId = 1048, typeId = 1082, baseId = 1, txt = 'addErrMess',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {372, 18}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 1, typeId = 1084, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {734}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[484] = { parentId = 1, typeId = 1085, nilable = true, orgTypeId = 1084 }
_typeInfoList[485] = { parentId = 1048, typeId = 1086, baseId = 1, txt = 'pushScope',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 1084}, retTypeId = {734}, children = {}, }
_typeInfoList[486] = { parentId = 1048, typeId = 1088, baseId = 1, txt = 'popScope',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[487] = { parentId = 1048, typeId = 1090, baseId = 1, txt = 'getCurrentClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[488] = { parentId = 1048, typeId = 1092, baseId = 1, txt = 'getCurrentNamespaceTypeInfo',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[489] = { parentId = 1048, typeId = 1094, baseId = 1, txt = 'pushClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {709, 10, 18, 18, 1011}, retTypeId = {708}, children = {}, }
_typeInfoList[490] = { parentId = 1048, typeId = 1102, baseId = 1, txt = 'popClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[491] = { parentId = 1048, typeId = 1104, baseId = 1, txt = 'pushbackStr',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[492] = { parentId = 1048, typeId = 1106, baseId = 1, txt = 'analyzeDecl',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 376, 376}, retTypeId = {843}, children = {}, }
_typeInfoList[493] = { parentId = 1048, typeId = 1108, baseId = 1, txt = 'analyzeDeclVar',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18, 10, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[494] = { parentId = 1048, typeId = 1110, baseId = 1, txt = 'analyzeDeclFunc',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 377, 376, 377}, retTypeId = {842}, children = {}, }
_typeInfoList[495] = { parentId = 1048, typeId = 1112, baseId = 1, txt = 'analyzeDeclClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[496] = { parentId = 1048, typeId = 1114, baseId = 1, txt = 'analyzeExp',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 12}, retTypeId = {842}, children = {}, }
_typeInfoList[497] = { parentId = 1, typeId = 1116, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[498] = { parentId = 1, typeId = 1117, nilable = true, orgTypeId = 1116 }
_typeInfoList[499] = { parentId = 1048, typeId = 1118, baseId = 1, txt = 'analyzeStatementList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {1116, 18}, retTypeId = {}, children = {}, }
_typeInfoList[500] = { parentId = 1048, typeId = 1120, baseId = 1, txt = 'analyzeStatement',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[501] = { parentId = 1048, typeId = 1122, baseId = 1, txt = 'analyzeExpSymbol',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 376, 18, 843, 10}, retTypeId = {842}, children = {}, }
_typeInfoList[502] = { parentId = 1048, typeId = 1124, baseId = 1, txt = 'analyzeExpList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10}, retTypeId = {1016}, children = {}, }
_typeInfoList[503] = { parentId = 1, typeId = 1136, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[504] = { parentId = 1, typeId = 1137, nilable = true, orgTypeId = 1136 }
_typeInfoList[505] = { parentId = 1048, typeId = 1138, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1136}, children = {}, }
_typeInfoList[506] = { parentId = 1, typeId = 1152, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[507] = { parentId = 1, typeId = 1153, nilable = true, orgTypeId = 1152 }
_typeInfoList[508] = { parentId = 106, typeId = 1154, baseId = 1, txt = 'regOpLevel',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {12, 1152}, retTypeId = {}, children = {}, }
_typeInfoList[509] = { parentId = 1, typeId = 1188, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[510] = { parentId = 1, typeId = 1189, nilable = true, orgTypeId = 1188 }
_typeInfoList[511] = { parentId = 106, typeId = 1192, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[512] = { parentId = 106, typeId = 1194, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[513] = { parentId = 106, typeId = 1218, baseId = 842, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1222, 1230}, }
_typeInfoList[514] = { parentId = 106, typeId = 1219, nilable = true, orgTypeId = 1218 }
_typeInfoList[515] = { parentId = 982, typeId = 1220, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1218, 8}, retTypeId = {}, children = {}, }
_typeInfoList[516] = { parentId = 1218, typeId = 1222, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[517] = { parentId = 1, typeId = 1228, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[518] = { parentId = 1, typeId = 1229, nilable = true, orgTypeId = 1228 }
_typeInfoList[519] = { parentId = 1218, typeId = 1230, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1228}, retTypeId = {}, children = {}, }
_typeInfoList[520] = { parentId = 106, typeId = 1236, baseId = 842, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1240, 1248, 1250}, }
_typeInfoList[521] = { parentId = 106, typeId = 1237, nilable = true, orgTypeId = 1236 }
_typeInfoList[522] = { parentId = 982, typeId = 1238, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1236, 8}, retTypeId = {}, children = {}, }
_typeInfoList[523] = { parentId = 1236, typeId = 1240, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[524] = { parentId = 1, typeId = 1246, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[525] = { parentId = 1, typeId = 1247, nilable = true, orgTypeId = 1246 }
_typeInfoList[526] = { parentId = 1236, typeId = 1248, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1246, 18}, retTypeId = {}, children = {}, }
_typeInfoList[527] = { parentId = 1236, typeId = 1250, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[528] = { parentId = 106, typeId = 1258, baseId = 842, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1266, 1278, 1282, 1286}, }
_typeInfoList[529] = { parentId = 106, typeId = 1259, nilable = true, orgTypeId = 1258 }
_typeInfoList[530] = { parentId = 982, typeId = 1260, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1258, 8}, retTypeId = {}, children = {}, }
_typeInfoList[531] = { parentId = 1258, typeId = 1266, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[532] = { parentId = 1, typeId = 1272, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[533] = { parentId = 1, typeId = 1273, nilable = true, orgTypeId = 1272 }
_typeInfoList[534] = { parentId = 1, typeId = 1274, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[535] = { parentId = 1, typeId = 1275, nilable = true, orgTypeId = 1274 }
_typeInfoList[536] = { parentId = 1, typeId = 1276, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 1010}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[537] = { parentId = 1, typeId = 1277, nilable = true, orgTypeId = 1276 }
_typeInfoList[538] = { parentId = 1258, typeId = 1278, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1272, 1274, 1276}, retTypeId = {}, children = {}, }
_typeInfoList[539] = { parentId = 1, typeId = 1280, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 1, typeId = 1281, nilable = true, orgTypeId = 1280 }
_typeInfoList[541] = { parentId = 1258, typeId = 1282, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1280}, children = {}, }
_typeInfoList[542] = { parentId = 1, typeId = 1284, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 1010}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[543] = { parentId = 1, typeId = 1285, nilable = true, orgTypeId = 1284 }
_typeInfoList[544] = { parentId = 1258, typeId = 1286, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1284}, children = {}, }
_typeInfoList[545] = { parentId = 106, typeId = 1298, baseId = 842, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1302, 1310, 1312, 1314, 1316, 1318}, }
_typeInfoList[546] = { parentId = 106, typeId = 1299, nilable = true, orgTypeId = 1298 }
_typeInfoList[547] = { parentId = 982, typeId = 1300, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1298, 8}, retTypeId = {}, children = {}, }
_typeInfoList[548] = { parentId = 1298, typeId = 1302, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[549] = { parentId = 1, typeId = 1308, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[550] = { parentId = 1, typeId = 1309, nilable = true, orgTypeId = 1308 }
_typeInfoList[551] = { parentId = 1298, typeId = 1310, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1308, 842, 10, 10, 18}, retTypeId = {}, children = {}, }
_typeInfoList[552] = { parentId = 1298, typeId = 1312, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[553] = { parentId = 1298, typeId = 1314, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[554] = { parentId = 1298, typeId = 1316, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[555] = { parentId = 1298, typeId = 1318, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[556] = { parentId = 106, typeId = 1326, baseId = 842, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1332, 1342, 1344, 1348}, }
_typeInfoList[557] = { parentId = 106, typeId = 1327, nilable = true, orgTypeId = 1326 }
_typeInfoList[558] = { parentId = 982, typeId = 1328, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1326, 8}, retTypeId = {}, children = {}, }
_typeInfoList[559] = { parentId = 1326, typeId = 1332, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[560] = { parentId = 1, typeId = 1338, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[561] = { parentId = 1, typeId = 1339, nilable = true, orgTypeId = 1338 }
_typeInfoList[562] = { parentId = 1, typeId = 1340, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[563] = { parentId = 1, typeId = 1341, nilable = true, orgTypeId = 1340 }
_typeInfoList[564] = { parentId = 1326, typeId = 1342, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1338, 18, 1340}, retTypeId = {}, children = {}, }
_typeInfoList[565] = { parentId = 1326, typeId = 1344, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[566] = { parentId = 1, typeId = 1346, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 1, typeId = 1347, nilable = true, orgTypeId = 1346 }
_typeInfoList[568] = { parentId = 1326, typeId = 1348, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1346}, children = {}, }
_typeInfoList[569] = { parentId = 106, typeId = 1350, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1352, 1354, 1356, 1358}, }
_typeInfoList[570] = { parentId = 106, typeId = 1351, nilable = true, orgTypeId = 1350 }
_typeInfoList[571] = { parentId = 1350, typeId = 1352, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[572] = { parentId = 1350, typeId = 1354, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[573] = { parentId = 1350, typeId = 1356, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[574] = { parentId = 1350, typeId = 1358, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 842, 1326}, retTypeId = {}, children = {}, }
_typeInfoList[575] = { parentId = 106, typeId = 1364, baseId = 842, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1370, 1380, 1384}, }
_typeInfoList[576] = { parentId = 106, typeId = 1365, nilable = true, orgTypeId = 1364 }
_typeInfoList[577] = { parentId = 982, typeId = 1366, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1364, 8}, retTypeId = {}, children = {}, }
_typeInfoList[578] = { parentId = 1364, typeId = 1370, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[579] = { parentId = 1, typeId = 1376, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[580] = { parentId = 1, typeId = 1377, nilable = true, orgTypeId = 1376 }
_typeInfoList[581] = { parentId = 1, typeId = 1378, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1350}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[582] = { parentId = 1, typeId = 1379, nilable = true, orgTypeId = 1378 }
_typeInfoList[583] = { parentId = 1364, typeId = 1380, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1376, 1378}, retTypeId = {}, children = {}, }
_typeInfoList[584] = { parentId = 1, typeId = 1382, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1350}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[585] = { parentId = 1, typeId = 1383, nilable = true, orgTypeId = 1382 }
_typeInfoList[586] = { parentId = 1364, typeId = 1384, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1382}, children = {}, }
_typeInfoList[587] = { parentId = 982, typeId = 1390, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1016, 8}, retTypeId = {}, children = {}, }
_typeInfoList[588] = { parentId = 1016, typeId = 1394, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[589] = { parentId = 1, typeId = 1400, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[590] = { parentId = 1, typeId = 1401, nilable = true, orgTypeId = 1400 }
_typeInfoList[591] = { parentId = 1, typeId = 1402, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[592] = { parentId = 1, typeId = 1403, nilable = true, orgTypeId = 1402 }
_typeInfoList[593] = { parentId = 1016, typeId = 1404, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1400, 1402}, retTypeId = {}, children = {}, }
_typeInfoList[594] = { parentId = 1, typeId = 1406, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[595] = { parentId = 1, typeId = 1407, nilable = true, orgTypeId = 1406 }
_typeInfoList[596] = { parentId = 1016, typeId = 1408, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1406}, children = {}, }
_typeInfoList[597] = { parentId = 106, typeId = 1410, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1412, 1414, 1416}, }
_typeInfoList[598] = { parentId = 106, typeId = 1411, nilable = true, orgTypeId = 1410 }
_typeInfoList[599] = { parentId = 1410, typeId = 1412, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1016}, children = {}, }
_typeInfoList[600] = { parentId = 1410, typeId = 1414, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[601] = { parentId = 1410, typeId = 1416, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1016, 1326}, retTypeId = {}, children = {}, }
_typeInfoList[602] = { parentId = 106, typeId = 1426, baseId = 842, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1432, 1442, 1444, 1448, 1450}, }
_typeInfoList[603] = { parentId = 106, typeId = 1427, nilable = true, orgTypeId = 1426 }
_typeInfoList[604] = { parentId = 982, typeId = 1428, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1426, 8}, retTypeId = {}, children = {}, }
_typeInfoList[605] = { parentId = 1426, typeId = 1432, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[606] = { parentId = 1, typeId = 1438, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[607] = { parentId = 1, typeId = 1439, nilable = true, orgTypeId = 1438 }
_typeInfoList[608] = { parentId = 1, typeId = 1440, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1410}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[609] = { parentId = 1, typeId = 1441, nilable = true, orgTypeId = 1440 }
_typeInfoList[610] = { parentId = 1426, typeId = 1442, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1438, 842, 1440, 1327}, retTypeId = {}, children = {}, }
_typeInfoList[611] = { parentId = 1426, typeId = 1444, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[612] = { parentId = 1, typeId = 1446, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1410}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[613] = { parentId = 1, typeId = 1447, nilable = true, orgTypeId = 1446 }
_typeInfoList[614] = { parentId = 1426, typeId = 1448, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1446}, children = {}, }
_typeInfoList[615] = { parentId = 1426, typeId = 1450, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[616] = { parentId = 106, typeId = 1458, baseId = 842, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1462, 1470, 1472, 1474}, }
_typeInfoList[617] = { parentId = 106, typeId = 1459, nilable = true, orgTypeId = 1458 }
_typeInfoList[618] = { parentId = 982, typeId = 1460, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1458, 8}, retTypeId = {}, children = {}, }
_typeInfoList[619] = { parentId = 1458, typeId = 1462, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[620] = { parentId = 1, typeId = 1468, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[621] = { parentId = 1, typeId = 1469, nilable = true, orgTypeId = 1468 }
_typeInfoList[622] = { parentId = 1458, typeId = 1470, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1468, 842, 1326}, retTypeId = {}, children = {}, }
_typeInfoList[623] = { parentId = 1458, typeId = 1472, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[624] = { parentId = 1458, typeId = 1474, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[625] = { parentId = 106, typeId = 1482, baseId = 842, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1486, 1494, 1496, 1498}, }
_typeInfoList[626] = { parentId = 106, typeId = 1483, nilable = true, orgTypeId = 1482 }
_typeInfoList[627] = { parentId = 982, typeId = 1484, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1482, 8}, retTypeId = {}, children = {}, }
_typeInfoList[628] = { parentId = 1482, typeId = 1486, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[629] = { parentId = 1, typeId = 1492, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[630] = { parentId = 1, typeId = 1493, nilable = true, orgTypeId = 1492 }
_typeInfoList[631] = { parentId = 1482, typeId = 1494, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1492, 1326, 842}, retTypeId = {}, children = {}, }
_typeInfoList[632] = { parentId = 1482, typeId = 1496, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[633] = { parentId = 1482, typeId = 1498, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[634] = { parentId = 106, typeId = 1512, baseId = 842, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1516, 1524, 1526, 1528, 1530, 1532, 1534}, }
_typeInfoList[635] = { parentId = 106, typeId = 1513, nilable = true, orgTypeId = 1512 }
_typeInfoList[636] = { parentId = 982, typeId = 1514, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1512, 8}, retTypeId = {}, children = {}, }
_typeInfoList[637] = { parentId = 1512, typeId = 1516, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[638] = { parentId = 1, typeId = 1522, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[639] = { parentId = 1, typeId = 1523, nilable = true, orgTypeId = 1522 }
_typeInfoList[640] = { parentId = 1512, typeId = 1524, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1522, 1326, 376, 842, 842, 843}, retTypeId = {}, children = {}, }
_typeInfoList[641] = { parentId = 1512, typeId = 1526, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[642] = { parentId = 1512, typeId = 1528, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[643] = { parentId = 1512, typeId = 1530, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[644] = { parentId = 1512, typeId = 1532, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[645] = { parentId = 1512, typeId = 1534, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {843}, children = {}, }
_typeInfoList[646] = { parentId = 106, typeId = 1544, baseId = 842, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1550, 1560, 1564, 1566, 1568}, }
_typeInfoList[647] = { parentId = 106, typeId = 1545, nilable = true, orgTypeId = 1544 }
_typeInfoList[648] = { parentId = 982, typeId = 1546, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1544, 8}, retTypeId = {}, children = {}, }
_typeInfoList[649] = { parentId = 1544, typeId = 1550, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[650] = { parentId = 1, typeId = 1556, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[651] = { parentId = 1, typeId = 1557, nilable = true, orgTypeId = 1556 }
_typeInfoList[652] = { parentId = 1, typeId = 1558, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[653] = { parentId = 1, typeId = 1559, nilable = true, orgTypeId = 1558 }
_typeInfoList[654] = { parentId = 1544, typeId = 1560, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1556, 1558, 842, 1326}, retTypeId = {}, children = {}, }
_typeInfoList[655] = { parentId = 1, typeId = 1562, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[656] = { parentId = 1, typeId = 1563, nilable = true, orgTypeId = 1562 }
_typeInfoList[657] = { parentId = 1544, typeId = 1564, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1562}, children = {}, }
_typeInfoList[658] = { parentId = 1544, typeId = 1566, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[659] = { parentId = 1544, typeId = 1568, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[660] = { parentId = 106, typeId = 1580, baseId = 842, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1584, 1592, 1594, 1596, 1598, 1600}, }
_typeInfoList[661] = { parentId = 106, typeId = 1581, nilable = true, orgTypeId = 1580 }
_typeInfoList[662] = { parentId = 982, typeId = 1582, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1580, 8}, retTypeId = {}, children = {}, }
_typeInfoList[663] = { parentId = 1580, typeId = 1584, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[664] = { parentId = 1, typeId = 1590, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[665] = { parentId = 1, typeId = 1591, nilable = true, orgTypeId = 1590 }
_typeInfoList[666] = { parentId = 1580, typeId = 1592, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1590, 376, 377, 842, 1326}, retTypeId = {}, children = {}, }
_typeInfoList[667] = { parentId = 1580, typeId = 1594, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[668] = { parentId = 1580, typeId = 1596, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[669] = { parentId = 1580, typeId = 1598, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[670] = { parentId = 1580, typeId = 1600, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[671] = { parentId = 106, typeId = 1614, baseId = 842, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1618, 1626, 1628, 1630, 1632, 1634, 1636}, }
_typeInfoList[672] = { parentId = 106, typeId = 1615, nilable = true, orgTypeId = 1614 }
_typeInfoList[673] = { parentId = 982, typeId = 1616, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1614, 8}, retTypeId = {}, children = {}, }
_typeInfoList[674] = { parentId = 1614, typeId = 1618, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[675] = { parentId = 1, typeId = 1624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[676] = { parentId = 1, typeId = 1625, nilable = true, orgTypeId = 1624 }
_typeInfoList[677] = { parentId = 1614, typeId = 1626, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1624, 376, 377, 842, 1326, 10}, retTypeId = {}, children = {}, }
_typeInfoList[678] = { parentId = 1614, typeId = 1628, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[679] = { parentId = 1614, typeId = 1630, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[680] = { parentId = 1614, typeId = 1632, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[681] = { parentId = 1614, typeId = 1634, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[682] = { parentId = 1614, typeId = 1636, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[683] = { parentId = 106, typeId = 1642, baseId = 842, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1646, 1654, 1656}, }
_typeInfoList[684] = { parentId = 106, typeId = 1643, nilable = true, orgTypeId = 1642 }
_typeInfoList[685] = { parentId = 982, typeId = 1644, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1642, 8}, retTypeId = {}, children = {}, }
_typeInfoList[686] = { parentId = 1642, typeId = 1646, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[687] = { parentId = 1, typeId = 1652, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[688] = { parentId = 1, typeId = 1653, nilable = true, orgTypeId = 1652 }
_typeInfoList[689] = { parentId = 1642, typeId = 1654, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1652, 1017}, retTypeId = {}, children = {}, }
_typeInfoList[690] = { parentId = 1642, typeId = 1656, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[691] = { parentId = 106, typeId = 1658, baseId = 842, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1662, 1670}, }
_typeInfoList[692] = { parentId = 106, typeId = 1659, nilable = true, orgTypeId = 1658 }
_typeInfoList[693] = { parentId = 982, typeId = 1660, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1658, 8}, retTypeId = {}, children = {}, }
_typeInfoList[694] = { parentId = 1658, typeId = 1662, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[695] = { parentId = 1, typeId = 1668, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[696] = { parentId = 1, typeId = 1669, nilable = true, orgTypeId = 1668 }
_typeInfoList[697] = { parentId = 1658, typeId = 1670, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1668}, retTypeId = {}, children = {}, }
_typeInfoList[698] = { parentId = 106, typeId = 1678, baseId = 842, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1682, 1690, 1692, 1694}, }
_typeInfoList[699] = { parentId = 106, typeId = 1679, nilable = true, orgTypeId = 1678 }
_typeInfoList[700] = { parentId = 982, typeId = 1680, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1678, 8}, retTypeId = {}, children = {}, }
_typeInfoList[701] = { parentId = 1678, typeId = 1682, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[702] = { parentId = 1, typeId = 1688, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[703] = { parentId = 1, typeId = 1689, nilable = true, orgTypeId = 1688 }
_typeInfoList[704] = { parentId = 1678, typeId = 1690, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1688, 842, 1017}, retTypeId = {}, children = {}, }
_typeInfoList[705] = { parentId = 1678, typeId = 1692, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[706] = { parentId = 1678, typeId = 1694, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[707] = { parentId = 106, typeId = 1702, baseId = 842, txt = 'ExpUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1706, 1714, 1716, 1718}, }
_typeInfoList[708] = { parentId = 106, typeId = 1703, nilable = true, orgTypeId = 1702 }
_typeInfoList[709] = { parentId = 982, typeId = 1704, baseId = 1, txt = 'processExpUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1702, 8}, retTypeId = {}, children = {}, }
_typeInfoList[710] = { parentId = 1702, typeId = 1706, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[711] = { parentId = 1, typeId = 1712, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[712] = { parentId = 1, typeId = 1713, nilable = true, orgTypeId = 1712 }
_typeInfoList[713] = { parentId = 1702, typeId = 1714, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1712, 842, 843}, retTypeId = {}, children = {}, }
_typeInfoList[714] = { parentId = 1702, typeId = 1716, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[715] = { parentId = 1702, typeId = 1718, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {843}, children = {}, }
_typeInfoList[716] = { parentId = 106, typeId = 1724, baseId = 842, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1728, 1736, 1738}, }
_typeInfoList[717] = { parentId = 106, typeId = 1725, nilable = true, orgTypeId = 1724 }
_typeInfoList[718] = { parentId = 982, typeId = 1726, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1724, 8}, retTypeId = {}, children = {}, }
_typeInfoList[719] = { parentId = 1724, typeId = 1728, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[720] = { parentId = 1, typeId = 1734, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[721] = { parentId = 1, typeId = 1735, nilable = true, orgTypeId = 1734 }
_typeInfoList[722] = { parentId = 1724, typeId = 1736, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1734, 376}, retTypeId = {}, children = {}, }
_typeInfoList[723] = { parentId = 1724, typeId = 1738, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[724] = { parentId = 106, typeId = 1748, baseId = 842, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1752, 1760, 1762, 1764, 1766}, }
_typeInfoList[725] = { parentId = 106, typeId = 1749, nilable = true, orgTypeId = 1748 }
_typeInfoList[726] = { parentId = 982, typeId = 1750, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1748, 8}, retTypeId = {}, children = {}, }
_typeInfoList[727] = { parentId = 1748, typeId = 1752, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[728] = { parentId = 1, typeId = 1758, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[729] = { parentId = 1, typeId = 1759, nilable = true, orgTypeId = 1758 }
_typeInfoList[730] = { parentId = 1748, typeId = 1760, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1758, 376, 842, 842}, retTypeId = {}, children = {}, }
_typeInfoList[731] = { parentId = 1748, typeId = 1762, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[732] = { parentId = 1748, typeId = 1764, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[733] = { parentId = 1748, typeId = 1766, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[734] = { parentId = 106, typeId = 1776, baseId = 842, txt = 'UnwrapSetNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1780, 1788, 1790, 1792, 1794}, }
_typeInfoList[735] = { parentId = 106, typeId = 1777, nilable = true, orgTypeId = 1776 }
_typeInfoList[736] = { parentId = 982, typeId = 1778, baseId = 1, txt = 'processUnwrapSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1776, 8}, retTypeId = {}, children = {}, }
_typeInfoList[737] = { parentId = 1776, typeId = 1780, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[738] = { parentId = 1, typeId = 1786, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[739] = { parentId = 1, typeId = 1787, nilable = true, orgTypeId = 1786 }
_typeInfoList[740] = { parentId = 1776, typeId = 1788, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1786, 1016, 1016, 1327}, retTypeId = {}, children = {}, }
_typeInfoList[741] = { parentId = 1776, typeId = 1790, baseId = 1, txt = 'get_dstExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1016}, children = {}, }
_typeInfoList[742] = { parentId = 1776, typeId = 1792, baseId = 1, txt = 'get_srcExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1016}, children = {}, }
_typeInfoList[743] = { parentId = 1776, typeId = 1794, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[744] = { parentId = 106, typeId = 1804, baseId = 842, txt = 'IfUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1808, 1816, 1818, 1820, 1822}, }
_typeInfoList[745] = { parentId = 106, typeId = 1805, nilable = true, orgTypeId = 1804 }
_typeInfoList[746] = { parentId = 982, typeId = 1806, baseId = 1, txt = 'processIfUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1804, 8}, retTypeId = {}, children = {}, }
_typeInfoList[747] = { parentId = 1804, typeId = 1808, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[748] = { parentId = 1, typeId = 1814, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[749] = { parentId = 1, typeId = 1815, nilable = true, orgTypeId = 1814 }
_typeInfoList[750] = { parentId = 1804, typeId = 1816, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1814, 842, 1326, 1327}, retTypeId = {}, children = {}, }
_typeInfoList[751] = { parentId = 1804, typeId = 1818, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[752] = { parentId = 1804, typeId = 1820, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1326}, children = {}, }
_typeInfoList[753] = { parentId = 1804, typeId = 1822, baseId = 1, txt = 'get_nilBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[754] = { parentId = 106, typeId = 1828, baseId = 842, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1832, 1840, 1842}, }
_typeInfoList[755] = { parentId = 106, typeId = 1829, nilable = true, orgTypeId = 1828 }
_typeInfoList[756] = { parentId = 982, typeId = 1830, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1828, 8}, retTypeId = {}, children = {}, }
_typeInfoList[757] = { parentId = 1828, typeId = 1832, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[758] = { parentId = 1, typeId = 1838, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[759] = { parentId = 1, typeId = 1839, nilable = true, orgTypeId = 1838 }
_typeInfoList[760] = { parentId = 1828, typeId = 1840, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1838, 842}, retTypeId = {}, children = {}, }
_typeInfoList[761] = { parentId = 1828, typeId = 1842, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[762] = { parentId = 106, typeId = 1852, baseId = 842, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1856, 1864, 1866, 1868, 1870}, }
_typeInfoList[763] = { parentId = 106, typeId = 1853, nilable = true, orgTypeId = 1852 }
_typeInfoList[764] = { parentId = 982, typeId = 1854, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1852, 8}, retTypeId = {}, children = {}, }
_typeInfoList[765] = { parentId = 1852, typeId = 1856, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[766] = { parentId = 1, typeId = 1862, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[767] = { parentId = 1, typeId = 1863, nilable = true, orgTypeId = 1862 }
_typeInfoList[768] = { parentId = 1852, typeId = 1864, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1862, 376, 18, 842}, retTypeId = {}, children = {}, }
_typeInfoList[769] = { parentId = 1852, typeId = 1866, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[770] = { parentId = 1852, typeId = 1868, baseId = 1, txt = 'get_macroMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[771] = { parentId = 1852, typeId = 1870, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[772] = { parentId = 106, typeId = 1878, baseId = 842, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1882, 1890, 1892, 1894}, }
_typeInfoList[773] = { parentId = 106, typeId = 1879, nilable = true, orgTypeId = 1878 }
_typeInfoList[774] = { parentId = 982, typeId = 1880, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1878, 8}, retTypeId = {}, children = {}, }
_typeInfoList[775] = { parentId = 1878, typeId = 1882, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[776] = { parentId = 1, typeId = 1888, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[777] = { parentId = 1, typeId = 1889, nilable = true, orgTypeId = 1888 }
_typeInfoList[778] = { parentId = 1878, typeId = 1890, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1888, 842, 842}, retTypeId = {}, children = {}, }
_typeInfoList[779] = { parentId = 1878, typeId = 1892, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[780] = { parentId = 1878, typeId = 1894, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[781] = { parentId = 106, typeId = 1902, baseId = 842, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1906, 1914, 1916, 1918}, }
_typeInfoList[782] = { parentId = 106, typeId = 1903, nilable = true, orgTypeId = 1902 }
_typeInfoList[783] = { parentId = 982, typeId = 1904, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1902, 8}, retTypeId = {}, children = {}, }
_typeInfoList[784] = { parentId = 1902, typeId = 1906, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[785] = { parentId = 1, typeId = 1912, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[786] = { parentId = 1, typeId = 1913, nilable = true, orgTypeId = 1912 }
_typeInfoList[787] = { parentId = 1902, typeId = 1914, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1912, 842, 1017}, retTypeId = {}, children = {}, }
_typeInfoList[788] = { parentId = 1902, typeId = 1916, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[789] = { parentId = 1902, typeId = 1918, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[790] = { parentId = 106, typeId = 1924, baseId = 842, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1928, 1936, 1938}, }
_typeInfoList[791] = { parentId = 106, typeId = 1925, nilable = true, orgTypeId = 1924 }
_typeInfoList[792] = { parentId = 982, typeId = 1926, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1924, 8}, retTypeId = {}, children = {}, }
_typeInfoList[793] = { parentId = 1924, typeId = 1928, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[794] = { parentId = 1, typeId = 1934, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[795] = { parentId = 1, typeId = 1935, nilable = true, orgTypeId = 1934 }
_typeInfoList[796] = { parentId = 1924, typeId = 1936, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1934, 376}, retTypeId = {}, children = {}, }
_typeInfoList[797] = { parentId = 1924, typeId = 1938, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[798] = { parentId = 106, typeId = 1944, baseId = 842, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1948, 1956, 1958}, }
_typeInfoList[799] = { parentId = 106, typeId = 1945, nilable = true, orgTypeId = 1944 }
_typeInfoList[800] = { parentId = 982, typeId = 1946, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1944, 8}, retTypeId = {}, children = {}, }
_typeInfoList[801] = { parentId = 1944, typeId = 1948, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[802] = { parentId = 1, typeId = 1954, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[803] = { parentId = 1, typeId = 1955, nilable = true, orgTypeId = 1954 }
_typeInfoList[804] = { parentId = 1944, typeId = 1956, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1954, 842}, retTypeId = {}, children = {}, }
_typeInfoList[805] = { parentId = 1944, typeId = 1958, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[806] = { parentId = 106, typeId = 1964, baseId = 842, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1970, 1980, 1984}, }
_typeInfoList[807] = { parentId = 106, typeId = 1965, nilable = true, orgTypeId = 1964 }
_typeInfoList[808] = { parentId = 982, typeId = 1966, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1964, 8}, retTypeId = {}, children = {}, }
_typeInfoList[809] = { parentId = 1964, typeId = 1970, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[810] = { parentId = 1, typeId = 1976, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[811] = { parentId = 1, typeId = 1977, nilable = true, orgTypeId = 1976 }
_typeInfoList[812] = { parentId = 1, typeId = 1978, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[813] = { parentId = 1, typeId = 1979, nilable = true, orgTypeId = 1978 }
_typeInfoList[814] = { parentId = 1964, typeId = 1980, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 1976, 1978}, retTypeId = {}, children = {}, }
_typeInfoList[815] = { parentId = 1, typeId = 1982, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[816] = { parentId = 1, typeId = 1983, nilable = true, orgTypeId = 1982 }
_typeInfoList[817] = { parentId = 1964, typeId = 1984, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1982}, children = {}, }
_typeInfoList[818] = { parentId = 106, typeId = 1990, baseId = 842, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1996, 2006, 2010, 2814}, }
_typeInfoList[819] = { parentId = 106, typeId = 1991, nilable = true, orgTypeId = 1990 }
_typeInfoList[820] = { parentId = 982, typeId = 1992, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1990, 8}, retTypeId = {}, children = {}, }
_typeInfoList[821] = { parentId = 1990, typeId = 1996, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[822] = { parentId = 1, typeId = 2002, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[823] = { parentId = 1, typeId = 2003, nilable = true, orgTypeId = 2002 }
_typeInfoList[824] = { parentId = 1, typeId = 2004, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[825] = { parentId = 1, typeId = 2005, nilable = true, orgTypeId = 2004 }
_typeInfoList[826] = { parentId = 1990, typeId = 2006, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2002, 2004}, retTypeId = {}, children = {}, }
_typeInfoList[827] = { parentId = 1, typeId = 2008, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[828] = { parentId = 1, typeId = 2009, nilable = true, orgTypeId = 2008 }
_typeInfoList[829] = { parentId = 1990, typeId = 2010, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2008}, children = {}, }
_typeInfoList[830] = { parentId = 106, typeId = 2016, baseId = 842, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2020, 2028, 2030}, }
_typeInfoList[831] = { parentId = 106, typeId = 2017, nilable = true, orgTypeId = 2016 }
_typeInfoList[832] = { parentId = 982, typeId = 2018, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2016, 8}, retTypeId = {}, children = {}, }
_typeInfoList[833] = { parentId = 2016, typeId = 2020, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[834] = { parentId = 1, typeId = 2026, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[835] = { parentId = 1, typeId = 2027, nilable = true, orgTypeId = 2026 }
_typeInfoList[836] = { parentId = 2016, typeId = 2028, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2026, 842}, retTypeId = {}, children = {}, }
_typeInfoList[837] = { parentId = 2016, typeId = 2030, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[838] = { parentId = 106, typeId = 2038, baseId = 842, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2042, 2050, 2052, 2054, 2802}, }
_typeInfoList[839] = { parentId = 106, typeId = 2039, nilable = true, orgTypeId = 2038 }
_typeInfoList[840] = { parentId = 982, typeId = 2040, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2038, 8}, retTypeId = {}, children = {}, }
_typeInfoList[841] = { parentId = 2038, typeId = 2042, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[842] = { parentId = 1, typeId = 2048, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[843] = { parentId = 1, typeId = 2049, nilable = true, orgTypeId = 2048 }
_typeInfoList[844] = { parentId = 2038, typeId = 2050, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2048, 376, 842}, retTypeId = {}, children = {}, }
_typeInfoList[845] = { parentId = 2038, typeId = 2052, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[846] = { parentId = 2038, typeId = 2054, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[847] = { parentId = 106, typeId = 2064, baseId = 842, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2068, 2076, 2078, 2080, 2082}, }
_typeInfoList[848] = { parentId = 106, typeId = 2065, nilable = true, orgTypeId = 2064 }
_typeInfoList[849] = { parentId = 982, typeId = 2066, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2064, 8}, retTypeId = {}, children = {}, }
_typeInfoList[850] = { parentId = 2064, typeId = 2068, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[851] = { parentId = 1, typeId = 2074, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[852] = { parentId = 1, typeId = 2075, nilable = true, orgTypeId = 2074 }
_typeInfoList[853] = { parentId = 2064, typeId = 2076, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2074, 376, 842, 708}, retTypeId = {}, children = {}, }
_typeInfoList[854] = { parentId = 2064, typeId = 2078, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[855] = { parentId = 2064, typeId = 2080, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[856] = { parentId = 2064, typeId = 2082, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[857] = { parentId = 106, typeId = 2084, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2086, 2088, 2090, 2092}, }
_typeInfoList[858] = { parentId = 106, typeId = 2085, nilable = true, orgTypeId = 2084 }
_typeInfoList[859] = { parentId = 2084, typeId = 2086, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[860] = { parentId = 2084, typeId = 2088, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1299}, children = {}, }
_typeInfoList[861] = { parentId = 2084, typeId = 2090, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[862] = { parentId = 2084, typeId = 2092, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {376, 1299, 708}, retTypeId = {}, children = {}, }
_typeInfoList[863] = { parentId = 106, typeId = 2118, baseId = 842, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2128, 2142, 2144, 2146, 2148, 2152, 2154, 2158, 2160, 2162, 2164, 2168, 2170}, }
_typeInfoList[864] = { parentId = 106, typeId = 2119, nilable = true, orgTypeId = 2118 }
_typeInfoList[865] = { parentId = 982, typeId = 2120, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2118, 8}, retTypeId = {}, children = {}, }
_typeInfoList[866] = { parentId = 2118, typeId = 2128, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[867] = { parentId = 1, typeId = 2134, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[868] = { parentId = 1, typeId = 2135, nilable = true, orgTypeId = 2134 }
_typeInfoList[869] = { parentId = 1, typeId = 2136, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2084}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[870] = { parentId = 1, typeId = 2137, nilable = true, orgTypeId = 2136 }
_typeInfoList[871] = { parentId = 1, typeId = 2138, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[872] = { parentId = 1, typeId = 2139, nilable = true, orgTypeId = 2138 }
_typeInfoList[873] = { parentId = 1, typeId = 2140, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2084}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[874] = { parentId = 1, typeId = 2141, nilable = true, orgTypeId = 2140 }
_typeInfoList[875] = { parentId = 2118, typeId = 2142, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2134, 18, 18, 10, 2136, 1017, 2138, 10, 1327, 1327, 2140, 1327}, retTypeId = {}, children = {}, }
_typeInfoList[876] = { parentId = 2118, typeId = 2144, baseId = 1, txt = 'get_mode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[877] = { parentId = 2118, typeId = 2146, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[878] = { parentId = 2118, typeId = 2148, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[879] = { parentId = 1, typeId = 2150, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2084}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[880] = { parentId = 1, typeId = 2151, nilable = true, orgTypeId = 2150 }
_typeInfoList[881] = { parentId = 2118, typeId = 2152, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2150}, children = {}, }
_typeInfoList[882] = { parentId = 2118, typeId = 2154, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[883] = { parentId = 1, typeId = 2156, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[884] = { parentId = 1, typeId = 2157, nilable = true, orgTypeId = 2156 }
_typeInfoList[885] = { parentId = 2118, typeId = 2158, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2156}, children = {}, }
_typeInfoList[886] = { parentId = 2118, typeId = 2160, baseId = 1, txt = 'get_unwrapFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[887] = { parentId = 2118, typeId = 2162, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[888] = { parentId = 2118, typeId = 2164, baseId = 1, txt = 'get_thenBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[889] = { parentId = 1, typeId = 2166, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2084}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[890] = { parentId = 1, typeId = 2167, nilable = true, orgTypeId = 2166 }
_typeInfoList[891] = { parentId = 2118, typeId = 2168, baseId = 1, txt = 'get_syncVarList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2166}, children = {}, }
_typeInfoList[892] = { parentId = 2118, typeId = 2170, baseId = 1, txt = 'get_syncBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1327}, children = {}, }
_typeInfoList[893] = { parentId = 106, typeId = 2172, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2178, 2180, 2184, 2186, 2188, 2190, 2194, 2196}, }
_typeInfoList[894] = { parentId = 106, typeId = 2173, nilable = true, orgTypeId = 2172 }
_typeInfoList[895] = { parentId = 1, typeId = 2174, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[896] = { parentId = 1, typeId = 2175, nilable = true, orgTypeId = 2174 }
_typeInfoList[897] = { parentId = 1, typeId = 2176, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[898] = { parentId = 1, typeId = 2177, nilable = true, orgTypeId = 2176 }
_typeInfoList[899] = { parentId = 2172, typeId = 2178, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[900] = { parentId = 2172, typeId = 2180, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[901] = { parentId = 1, typeId = 2182, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[902] = { parentId = 1, typeId = 2183, nilable = true, orgTypeId = 2182 }
_typeInfoList[903] = { parentId = 2172, typeId = 2184, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2182}, children = {}, }
_typeInfoList[904] = { parentId = 2172, typeId = 2186, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[905] = { parentId = 2172, typeId = 2188, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[906] = { parentId = 2172, typeId = 2190, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {843}, children = {}, }
_typeInfoList[907] = { parentId = 1, typeId = 2192, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[908] = { parentId = 1, typeId = 2193, nilable = true, orgTypeId = 2192 }
_typeInfoList[909] = { parentId = 2172, typeId = 2194, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2192}, children = {}, }
_typeInfoList[910] = { parentId = 2172, typeId = 2196, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {377, 377, 2174, 10, 18, 843, 2176}, retTypeId = {}, children = {}, }
_typeInfoList[911] = { parentId = 106, typeId = 2202, baseId = 842, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2206, 2214, 2216}, }
_typeInfoList[912] = { parentId = 106, typeId = 2203, nilable = true, orgTypeId = 2202 }
_typeInfoList[913] = { parentId = 982, typeId = 2204, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2202, 8}, retTypeId = {}, children = {}, }
_typeInfoList[914] = { parentId = 2202, typeId = 2206, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[915] = { parentId = 1, typeId = 2212, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[916] = { parentId = 1, typeId = 2213, nilable = true, orgTypeId = 2212 }
_typeInfoList[917] = { parentId = 2202, typeId = 2214, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2212, 2172}, retTypeId = {}, children = {}, }
_typeInfoList[918] = { parentId = 2202, typeId = 2216, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2172}, children = {}, }
_typeInfoList[919] = { parentId = 106, typeId = 2222, baseId = 842, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2226, 2234, 2236}, }
_typeInfoList[920] = { parentId = 106, typeId = 2223, nilable = true, orgTypeId = 2222 }
_typeInfoList[921] = { parentId = 982, typeId = 2224, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2222, 8}, retTypeId = {}, children = {}, }
_typeInfoList[922] = { parentId = 2222, typeId = 2226, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[923] = { parentId = 1, typeId = 2232, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[924] = { parentId = 1, typeId = 2233, nilable = true, orgTypeId = 2232 }
_typeInfoList[925] = { parentId = 2222, typeId = 2234, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2232, 2172}, retTypeId = {}, children = {}, }
_typeInfoList[926] = { parentId = 2222, typeId = 2236, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2172}, children = {}, }
_typeInfoList[927] = { parentId = 106, typeId = 2242, baseId = 842, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2246, 2254, 2256}, }
_typeInfoList[928] = { parentId = 106, typeId = 2243, nilable = true, orgTypeId = 2242 }
_typeInfoList[929] = { parentId = 982, typeId = 2244, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2242, 8}, retTypeId = {}, children = {}, }
_typeInfoList[930] = { parentId = 2242, typeId = 2246, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[931] = { parentId = 1, typeId = 2252, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[932] = { parentId = 1, typeId = 2253, nilable = true, orgTypeId = 2252 }
_typeInfoList[933] = { parentId = 2242, typeId = 2254, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2252, 2172}, retTypeId = {}, children = {}, }
_typeInfoList[934] = { parentId = 2242, typeId = 2256, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2172}, children = {}, }
_typeInfoList[935] = { parentId = 106, typeId = 2264, baseId = 842, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2268, 2276, 2278, 2280}, }
_typeInfoList[936] = { parentId = 106, typeId = 2265, nilable = true, orgTypeId = 2264 }
_typeInfoList[937] = { parentId = 982, typeId = 2266, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2264, 8}, retTypeId = {}, children = {}, }
_typeInfoList[938] = { parentId = 2264, typeId = 2268, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[939] = { parentId = 1, typeId = 2274, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[940] = { parentId = 1, typeId = 2275, nilable = true, orgTypeId = 2274 }
_typeInfoList[941] = { parentId = 2264, typeId = 2276, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2274, 708, 1016}, retTypeId = {}, children = {}, }
_typeInfoList[942] = { parentId = 2264, typeId = 2278, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[943] = { parentId = 2264, typeId = 2280, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1016}, children = {}, }
_typeInfoList[944] = { parentId = 106, typeId = 2296, baseId = 842, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2300, 2308, 2310, 2312, 2314, 2316, 2318, 2320}, }
_typeInfoList[945] = { parentId = 106, typeId = 2297, nilable = true, orgTypeId = 2296 }
_typeInfoList[946] = { parentId = 982, typeId = 2298, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2296, 8}, retTypeId = {}, children = {}, }
_typeInfoList[947] = { parentId = 2296, typeId = 2300, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[948] = { parentId = 1, typeId = 2306, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[949] = { parentId = 1, typeId = 2307, nilable = true, orgTypeId = 2306 }
_typeInfoList[950] = { parentId = 2296, typeId = 2308, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2306, 376, 1298, 10, 18, 18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[951] = { parentId = 2296, typeId = 2310, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[952] = { parentId = 2296, typeId = 2312, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[953] = { parentId = 2296, typeId = 2314, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[954] = { parentId = 2296, typeId = 2316, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[955] = { parentId = 2296, typeId = 2318, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[956] = { parentId = 2296, typeId = 2320, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[957] = { parentId = 106, typeId = 2328, baseId = 842, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2332, 2340, 2342, 2344}, }
_typeInfoList[958] = { parentId = 106, typeId = 2329, nilable = true, orgTypeId = 2328 }
_typeInfoList[959] = { parentId = 982, typeId = 2330, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2328, 8}, retTypeId = {}, children = {}, }
_typeInfoList[960] = { parentId = 2328, typeId = 2332, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[961] = { parentId = 1, typeId = 2338, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[962] = { parentId = 1, typeId = 2339, nilable = true, orgTypeId = 2338 }
_typeInfoList[963] = { parentId = 2328, typeId = 2340, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2338, 376, 1298}, retTypeId = {}, children = {}, }
_typeInfoList[964] = { parentId = 2328, typeId = 2342, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[965] = { parentId = 2328, typeId = 2344, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1298}, children = {}, }
_typeInfoList[966] = { parentId = 106, typeId = 2346, baseId = 842, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2350, 2358}, }
_typeInfoList[967] = { parentId = 106, typeId = 2347, nilable = true, orgTypeId = 2346 }
_typeInfoList[968] = { parentId = 982, typeId = 2348, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2346, 8}, retTypeId = {}, children = {}, }
_typeInfoList[969] = { parentId = 2346, typeId = 2350, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[970] = { parentId = 1, typeId = 2356, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[971] = { parentId = 1, typeId = 2357, nilable = true, orgTypeId = 2356 }
_typeInfoList[972] = { parentId = 2346, typeId = 2358, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2356}, retTypeId = {}, children = {}, }
_typeInfoList[973] = { parentId = 982, typeId = 2374, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {844, 8}, retTypeId = {}, children = {}, }
_typeInfoList[974] = { parentId = 844, typeId = 2382, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[975] = { parentId = 1, typeId = 2388, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[976] = { parentId = 1, typeId = 2389, nilable = true, orgTypeId = 2388 }
_typeInfoList[977] = { parentId = 1, typeId = 2390, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[978] = { parentId = 1, typeId = 2391, nilable = true, orgTypeId = 2390 }
_typeInfoList[979] = { parentId = 1, typeId = 2392, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2296}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[980] = { parentId = 1, typeId = 2393, nilable = true, orgTypeId = 2392 }
_typeInfoList[981] = { parentId = 1, typeId = 2394, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[982] = { parentId = 1, typeId = 2395, nilable = true, orgTypeId = 2394 }
_typeInfoList[983] = { parentId = 844, typeId = 2396, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2388, 18, 376, 2390, 2392, 734, 2394}, retTypeId = {}, children = {}, }
_typeInfoList[984] = { parentId = 844, typeId = 2398, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[985] = { parentId = 844, typeId = 2400, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[986] = { parentId = 1, typeId = 2402, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[987] = { parentId = 1, typeId = 2403, nilable = true, orgTypeId = 2402 }
_typeInfoList[988] = { parentId = 844, typeId = 2404, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2402}, children = {}, }
_typeInfoList[989] = { parentId = 1, typeId = 2406, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2296}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[990] = { parentId = 1, typeId = 2407, nilable = true, orgTypeId = 2406 }
_typeInfoList[991] = { parentId = 844, typeId = 2408, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2406}, children = {}, }
_typeInfoList[992] = { parentId = 844, typeId = 2410, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {734}, children = {}, }
_typeInfoList[993] = { parentId = 1, typeId = 2412, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[994] = { parentId = 1, typeId = 2413, nilable = true, orgTypeId = 2412 }
_typeInfoList[995] = { parentId = 844, typeId = 2414, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2412}, children = {}, }
_typeInfoList[996] = { parentId = 106, typeId = 2420, baseId = 842, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2424, 2432, 2434}, }
_typeInfoList[997] = { parentId = 106, typeId = 2421, nilable = true, orgTypeId = 2420 }
_typeInfoList[998] = { parentId = 982, typeId = 2422, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2420, 8}, retTypeId = {}, children = {}, }
_typeInfoList[999] = { parentId = 2420, typeId = 2424, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1000] = { parentId = 1, typeId = 2430, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1001] = { parentId = 1, typeId = 2431, nilable = true, orgTypeId = 2430 }
_typeInfoList[1002] = { parentId = 2420, typeId = 2432, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2430, 1018}, retTypeId = {}, children = {}, }
_typeInfoList[1003] = { parentId = 2420, typeId = 2434, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1018}, children = {}, }
_typeInfoList[1004] = { parentId = 106, typeId = 2436, baseId = 842, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2440, 2448, 2680}, }
_typeInfoList[1005] = { parentId = 106, typeId = 2437, nilable = true, orgTypeId = 2436 }
_typeInfoList[1006] = { parentId = 982, typeId = 2438, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2436, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1007] = { parentId = 2436, typeId = 2440, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1008] = { parentId = 1, typeId = 2446, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1009] = { parentId = 1, typeId = 2447, nilable = true, orgTypeId = 2446 }
_typeInfoList[1010] = { parentId = 2436, typeId = 2448, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2446}, retTypeId = {}, children = {}, }
_typeInfoList[1011] = { parentId = 106, typeId = 2456, baseId = 842, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2460, 2468, 2470, 2472, 2690}, }
_typeInfoList[1012] = { parentId = 106, typeId = 2457, nilable = true, orgTypeId = 2456 }
_typeInfoList[1013] = { parentId = 982, typeId = 2458, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2456, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1014] = { parentId = 2456, typeId = 2460, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1015] = { parentId = 1, typeId = 2466, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1016] = { parentId = 1, typeId = 2467, nilable = true, orgTypeId = 2466 }
_typeInfoList[1017] = { parentId = 2456, typeId = 2468, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2466, 376, 12}, retTypeId = {}, children = {}, }
_typeInfoList[1018] = { parentId = 2456, typeId = 2470, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1019] = { parentId = 2456, typeId = 2472, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[1020] = { parentId = 106, typeId = 2480, baseId = 842, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2484, 2492, 2494, 2496, 2700}, }
_typeInfoList[1021] = { parentId = 106, typeId = 2481, nilable = true, orgTypeId = 2480 }
_typeInfoList[1022] = { parentId = 982, typeId = 2482, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2480, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1023] = { parentId = 2480, typeId = 2484, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1024] = { parentId = 1, typeId = 2490, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1025] = { parentId = 1, typeId = 2491, nilable = true, orgTypeId = 2490 }
_typeInfoList[1026] = { parentId = 2480, typeId = 2492, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2490, 376, 12}, retTypeId = {}, children = {}, }
_typeInfoList[1027] = { parentId = 2480, typeId = 2494, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1028] = { parentId = 2480, typeId = 2496, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[1029] = { parentId = 106, typeId = 2504, baseId = 842, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2508, 2516, 2518, 2520, 2710}, }
_typeInfoList[1030] = { parentId = 106, typeId = 2505, nilable = true, orgTypeId = 2504 }
_typeInfoList[1031] = { parentId = 982, typeId = 2506, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2504, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1032] = { parentId = 2504, typeId = 2508, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1033] = { parentId = 1, typeId = 2514, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1034] = { parentId = 1, typeId = 2515, nilable = true, orgTypeId = 2514 }
_typeInfoList[1035] = { parentId = 2504, typeId = 2516, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2514, 376, 14}, retTypeId = {}, children = {}, }
_typeInfoList[1036] = { parentId = 2504, typeId = 2518, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1037] = { parentId = 2504, typeId = 2520, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[1038] = { parentId = 106, typeId = 2526, baseId = 842, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2530, 2538, 2540, 2720}, }
_typeInfoList[1039] = { parentId = 106, typeId = 2527, nilable = true, orgTypeId = 2526 }
_typeInfoList[1040] = { parentId = 982, typeId = 2528, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2526, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1041] = { parentId = 2526, typeId = 2530, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1042] = { parentId = 1, typeId = 2536, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1043] = { parentId = 1, typeId = 2537, nilable = true, orgTypeId = 2536 }
_typeInfoList[1044] = { parentId = 2526, typeId = 2538, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2536, 1017}, retTypeId = {}, children = {}, }
_typeInfoList[1045] = { parentId = 2526, typeId = 2540, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[1046] = { parentId = 106, typeId = 2546, baseId = 842, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2550, 2558, 2560, 2734}, }
_typeInfoList[1047] = { parentId = 106, typeId = 2547, nilable = true, orgTypeId = 2546 }
_typeInfoList[1048] = { parentId = 982, typeId = 2548, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2546, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1049] = { parentId = 2546, typeId = 2550, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1050] = { parentId = 1, typeId = 2556, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1051] = { parentId = 1, typeId = 2557, nilable = true, orgTypeId = 2556 }
_typeInfoList[1052] = { parentId = 2546, typeId = 2558, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2556, 1017}, retTypeId = {}, children = {}, }
_typeInfoList[1053] = { parentId = 2546, typeId = 2560, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1017}, children = {}, }
_typeInfoList[1054] = { parentId = 106, typeId = 2562, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2564, 2566, 2568}, }
_typeInfoList[1055] = { parentId = 106, typeId = 2563, nilable = true, orgTypeId = 2562 }
_typeInfoList[1056] = { parentId = 2562, typeId = 2564, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[1057] = { parentId = 2562, typeId = 2566, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[1058] = { parentId = 2562, typeId = 2568, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {842, 842}, retTypeId = {}, children = {}, }
_typeInfoList[1059] = { parentId = 106, typeId = 2576, baseId = 842, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2584, 2596, 2600, 2604, 2748}, }
_typeInfoList[1060] = { parentId = 106, typeId = 2577, nilable = true, orgTypeId = 2576 }
_typeInfoList[1061] = { parentId = 982, typeId = 2578, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2576, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1062] = { parentId = 2576, typeId = 2584, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1063] = { parentId = 1, typeId = 2590, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1064] = { parentId = 1, typeId = 2591, nilable = true, orgTypeId = 2590 }
_typeInfoList[1065] = { parentId = 1, typeId = 2592, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {842, 842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1066] = { parentId = 1, typeId = 2593, nilable = true, orgTypeId = 2592 }
_typeInfoList[1067] = { parentId = 1, typeId = 2594, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2562}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1068] = { parentId = 1, typeId = 2595, nilable = true, orgTypeId = 2594 }
_typeInfoList[1069] = { parentId = 2576, typeId = 2596, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2590, 2592, 2594}, retTypeId = {}, children = {}, }
_typeInfoList[1070] = { parentId = 1, typeId = 2598, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {842, 842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1071] = { parentId = 1, typeId = 2599, nilable = true, orgTypeId = 2598 }
_typeInfoList[1072] = { parentId = 2576, typeId = 2600, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2598}, children = {}, }
_typeInfoList[1073] = { parentId = 1, typeId = 2602, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2562}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1074] = { parentId = 1, typeId = 2603, nilable = true, orgTypeId = 2602 }
_typeInfoList[1075] = { parentId = 2576, typeId = 2604, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2602}, children = {}, }
_typeInfoList[1076] = { parentId = 106, typeId = 2612, baseId = 842, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2618, 2628, 2630, 2634, 2762}, }
_typeInfoList[1077] = { parentId = 106, typeId = 2613, nilable = true, orgTypeId = 2612 }
_typeInfoList[1078] = { parentId = 982, typeId = 2614, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2612, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1079] = { parentId = 2612, typeId = 2618, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1080] = { parentId = 1, typeId = 2624, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1081] = { parentId = 1, typeId = 2625, nilable = true, orgTypeId = 2624 }
_typeInfoList[1082] = { parentId = 1, typeId = 2626, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1083] = { parentId = 1, typeId = 2627, nilable = true, orgTypeId = 2626 }
_typeInfoList[1084] = { parentId = 2612, typeId = 2628, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2624, 376, 2626}, retTypeId = {}, children = {}, }
_typeInfoList[1085] = { parentId = 2612, typeId = 2630, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1086] = { parentId = 1, typeId = 2632, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1087] = { parentId = 1, typeId = 2633, nilable = true, orgTypeId = 2632 }
_typeInfoList[1088] = { parentId = 2612, typeId = 2634, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2632}, children = {}, }
_typeInfoList[1089] = { parentId = 106, typeId = 2640, baseId = 842, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2644, 2652, 2654, 2780}, }
_typeInfoList[1090] = { parentId = 106, typeId = 2641, nilable = true, orgTypeId = 2640 }
_typeInfoList[1091] = { parentId = 982, typeId = 2642, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2640, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1092] = { parentId = 2640, typeId = 2644, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1093] = { parentId = 1, typeId = 2650, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1094] = { parentId = 1, typeId = 2651, nilable = true, orgTypeId = 2650 }
_typeInfoList[1095] = { parentId = 2640, typeId = 2652, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2650, 376}, retTypeId = {}, children = {}, }
_typeInfoList[1096] = { parentId = 2640, typeId = 2654, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1097] = { parentId = 106, typeId = 2660, baseId = 842, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2664, 2672, 2674, 2790}, }
_typeInfoList[1098] = { parentId = 106, typeId = 2661, nilable = true, orgTypeId = 2660 }
_typeInfoList[1099] = { parentId = 982, typeId = 2662, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2660, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1100] = { parentId = 2660, typeId = 2664, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {982, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1101] = { parentId = 1, typeId = 2670, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1102] = { parentId = 1, typeId = 2671, nilable = true, orgTypeId = 2670 }
_typeInfoList[1103] = { parentId = 2660, typeId = 2672, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {372, 2670, 376}, retTypeId = {}, children = {}, }
_typeInfoList[1104] = { parentId = 2660, typeId = 2674, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1105] = { parentId = 1, typeId = 2676, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1106] = { parentId = 1, typeId = 2677, nilable = true, orgTypeId = 2676 }
_typeInfoList[1107] = { parentId = 1, typeId = 2678, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1108] = { parentId = 1, typeId = 2679, nilable = true, orgTypeId = 2678 }
_typeInfoList[1109] = { parentId = 2436, typeId = 2680, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2676, 2678}, children = {}, }
_typeInfoList[1110] = { parentId = 1, typeId = 2686, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1111] = { parentId = 1, typeId = 2687, nilable = true, orgTypeId = 2686 }
_typeInfoList[1112] = { parentId = 1, typeId = 2688, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1113] = { parentId = 1, typeId = 2689, nilable = true, orgTypeId = 2688 }
_typeInfoList[1114] = { parentId = 2456, typeId = 2690, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2686, 2688}, children = {}, }
_typeInfoList[1115] = { parentId = 1, typeId = 2696, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1116] = { parentId = 1, typeId = 2697, nilable = true, orgTypeId = 2696 }
_typeInfoList[1117] = { parentId = 1, typeId = 2698, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1118] = { parentId = 1, typeId = 2699, nilable = true, orgTypeId = 2698 }
_typeInfoList[1119] = { parentId = 2480, typeId = 2700, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2696, 2698}, children = {}, }
_typeInfoList[1120] = { parentId = 1, typeId = 2706, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1121] = { parentId = 1, typeId = 2707, nilable = true, orgTypeId = 2706 }
_typeInfoList[1122] = { parentId = 1, typeId = 2708, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1123] = { parentId = 1, typeId = 2709, nilable = true, orgTypeId = 2708 }
_typeInfoList[1124] = { parentId = 2504, typeId = 2710, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2706, 2708}, children = {}, }
_typeInfoList[1125] = { parentId = 1, typeId = 2716, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1126] = { parentId = 1, typeId = 2717, nilable = true, orgTypeId = 2716 }
_typeInfoList[1127] = { parentId = 1, typeId = 2718, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1128] = { parentId = 1, typeId = 2719, nilable = true, orgTypeId = 2718 }
_typeInfoList[1129] = { parentId = 2526, typeId = 2720, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2716, 2718}, children = {}, }
_typeInfoList[1130] = { parentId = 1, typeId = 2730, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1131] = { parentId = 1, typeId = 2731, nilable = true, orgTypeId = 2730 }
_typeInfoList[1132] = { parentId = 1, typeId = 2732, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1133] = { parentId = 1, typeId = 2733, nilable = true, orgTypeId = 2732 }
_typeInfoList[1134] = { parentId = 2546, typeId = 2734, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2730, 2732}, children = {}, }
_typeInfoList[1135] = { parentId = 1, typeId = 2744, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1136] = { parentId = 1, typeId = 2745, nilable = true, orgTypeId = 2744 }
_typeInfoList[1137] = { parentId = 1, typeId = 2746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1138] = { parentId = 1, typeId = 2747, nilable = true, orgTypeId = 2746 }
_typeInfoList[1139] = { parentId = 2576, typeId = 2748, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2744, 2746}, children = {}, }
_typeInfoList[1140] = { parentId = 1, typeId = 2758, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1141] = { parentId = 1, typeId = 2759, nilable = true, orgTypeId = 2758 }
_typeInfoList[1142] = { parentId = 1, typeId = 2760, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1143] = { parentId = 1, typeId = 2761, nilable = true, orgTypeId = 2760 }
_typeInfoList[1144] = { parentId = 2612, typeId = 2762, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2758, 2760}, children = {}, }
_typeInfoList[1145] = { parentId = 1, typeId = 2776, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1146] = { parentId = 1, typeId = 2777, nilable = true, orgTypeId = 2776 }
_typeInfoList[1147] = { parentId = 1, typeId = 2778, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1148] = { parentId = 1, typeId = 2779, nilable = true, orgTypeId = 2778 }
_typeInfoList[1149] = { parentId = 2640, typeId = 2780, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2776, 2778}, children = {}, }
_typeInfoList[1150] = { parentId = 1, typeId = 2786, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1151] = { parentId = 1, typeId = 2787, nilable = true, orgTypeId = 2786 }
_typeInfoList[1152] = { parentId = 1, typeId = 2788, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1153] = { parentId = 1, typeId = 2789, nilable = true, orgTypeId = 2788 }
_typeInfoList[1154] = { parentId = 2660, typeId = 2790, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2786, 2788}, children = {}, }
_typeInfoList[1155] = { parentId = 1, typeId = 2798, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1156] = { parentId = 1, typeId = 2799, nilable = true, orgTypeId = 2798 }
_typeInfoList[1157] = { parentId = 1, typeId = 2800, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1158] = { parentId = 1, typeId = 2801, nilable = true, orgTypeId = 2800 }
_typeInfoList[1159] = { parentId = 2038, typeId = 2802, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2798, 2800}, children = {}, }
_typeInfoList[1160] = { parentId = 1, typeId = 2810, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1161] = { parentId = 1, typeId = 2811, nilable = true, orgTypeId = 2810 }
_typeInfoList[1162] = { parentId = 1, typeId = 2812, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1163] = { parentId = 1, typeId = 2813, nilable = true, orgTypeId = 2812 }
_typeInfoList[1164] = { parentId = 1990, typeId = 2814, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2810, 2812}, children = {}, }
_typeInfoList[1165] = { parentId = 1014, typeId = 2820, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2420}, retTypeId = {26}, children = {}, }
_typeInfoList[1166] = { parentId = 1014, typeId = 2822, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1167] = { parentId = 106, typeId = 2828, baseId = 1, txt = '_TypeInfo',
        staticFlag = false, accessMode = 'pri', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2838}, }
_typeInfoList[1168] = { parentId = 106, typeId = 2829, nilable = true, orgTypeId = 2828 }
_typeInfoList[1169] = { parentId = 1, typeId = 2830, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1170] = { parentId = 1, typeId = 2831, nilable = true, orgTypeId = 2830 }
_typeInfoList[1171] = { parentId = 1, typeId = 2832, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1172] = { parentId = 1, typeId = 2833, nilable = true, orgTypeId = 2832 }
_typeInfoList[1173] = { parentId = 1, typeId = 2834, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1174] = { parentId = 1, typeId = 2835, nilable = true, orgTypeId = 2834 }
_typeInfoList[1175] = { parentId = 1, typeId = 2836, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1176] = { parentId = 1, typeId = 2837, nilable = true, orgTypeId = 2836 }
_typeInfoList[1177] = { parentId = 2828, typeId = 2838, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 2830, 2832, 2834, 12, 12, 18, 12, 10, 10, 12, 2836}, retTypeId = {}, children = {}, }
_typeInfoList[1178] = { parentId = 106, typeId = 2840, baseId = 1, txt = '_ModuleInfo',
        staticFlag = false, accessMode = 'pri', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2858}, }
_typeInfoList[1179] = { parentId = 106, typeId = 2841, nilable = true, orgTypeId = 2840 }
_typeInfoList[1180] = { parentId = 1, typeId = 2842, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1181] = { parentId = 1, typeId = 2843, nilable = true, orgTypeId = 2842 }
_typeInfoList[1182] = { parentId = 1, typeId = 2844, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1183] = { parentId = 1, typeId = 2845, nilable = true, orgTypeId = 2844 }
_typeInfoList[1184] = { parentId = 1, typeId = 2846, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1185] = { parentId = 1, typeId = 2847, nilable = true, orgTypeId = 2846 }
_typeInfoList[1186] = { parentId = 1, typeId = 2848, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 2846}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1187] = { parentId = 1, typeId = 2849, nilable = true, orgTypeId = 2848 }
_typeInfoList[1188] = { parentId = 1, typeId = 2850, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2828}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1189] = { parentId = 1, typeId = 2851, nilable = true, orgTypeId = 2850 }
_typeInfoList[1190] = { parentId = 1, typeId = 2852, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1191] = { parentId = 1, typeId = 2853, nilable = true, orgTypeId = 2852 }
_typeInfoList[1192] = { parentId = 1, typeId = 2854, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2852}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1193] = { parentId = 1, typeId = 2855, nilable = true, orgTypeId = 2854 }
_typeInfoList[1194] = { parentId = 1, typeId = 2856, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1195] = { parentId = 1, typeId = 2857, nilable = true, orgTypeId = 2856 }
_typeInfoList[1196] = { parentId = 2840, typeId = 2858, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2844, 2848, 2850, 2854, 2856}, retTypeId = {}, children = {}, }
_typeInfoList[1197] = { parentId = 1048, typeId = 2864, baseId = 1, txt = 'registBuiltInScope',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1198] = { parentId = 1048, typeId = 3094, baseId = 1, txt = 'error',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[1199] = { parentId = 1048, typeId = 3096, baseId = 1, txt = 'createNoneNode',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {372}, retTypeId = {842}, children = {}, }
_typeInfoList[1200] = { parentId = 1048, typeId = 3100, baseId = 1, txt = 'pushbackToken',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {377}, retTypeId = {}, children = {}, }
_typeInfoList[1201] = { parentId = 1, typeId = 3102, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1202] = { parentId = 1, typeId = 3103, nilable = true, orgTypeId = 3102 }
_typeInfoList[1203] = { parentId = 106, typeId = 3104, baseId = 1, txt = 'expandVal',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {3102, 6, 372}, retTypeId = {18}, children = {}, }
_typeInfoList[1204] = { parentId = 1, typeId = 3108, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1205] = { parentId = 1, typeId = 3109, nilable = true, orgTypeId = 3108 }
_typeInfoList[1206] = { parentId = 1048, typeId = 3110, baseId = 1, txt = 'newPushback',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {3108}, retTypeId = {}, children = {}, }
_typeInfoList[1207] = { parentId = 1048, typeId = 3112, baseId = 1, txt = 'getTokenNoErr',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[1208] = { parentId = 1048, typeId = 3126, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1209] = { parentId = 1048, typeId = 3128, baseId = 1, txt = 'pushback',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1210] = { parentId = 1048, typeId = 3130, baseId = 1, txt = 'pushbackStr',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1211] = { parentId = 1048, typeId = 3136, baseId = 1, txt = 'checkSymbol',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {376}, children = {}, }
_typeInfoList[1212] = { parentId = 1048, typeId = 3138, baseId = 1, txt = 'getSymbolToken',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {376}, children = {}, }
_typeInfoList[1213] = { parentId = 1048, typeId = 3140, baseId = 1, txt = 'checkToken',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 18}, retTypeId = {376}, children = {}, }
_typeInfoList[1214] = { parentId = 1048, typeId = 3142, baseId = 1, txt = 'checkNextToken',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {376}, children = {}, }
_typeInfoList[1215] = { parentId = 1048, typeId = 3144, baseId = 1, txt = 'analyzeBlock',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 734}, retTypeId = {1326}, children = {}, }
_typeInfoList[1216] = { parentId = 1048, typeId = 3152, baseId = 1, txt = 'analyzeImport',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1217] = { parentId = 3152, typeId = 3170, baseId = 1, txt = 'registTypeInfo',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {2828}, retTypeId = {708}, children = {}, }
_typeInfoList[1218] = { parentId = 3152, typeId = 3186, baseId = 1, txt = 'registMember',
        staticFlag = true, accessMode = 'pri', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {}, children = {}, }
_typeInfoList[1219] = { parentId = 1048, typeId = 3194, baseId = 1, txt = 'analyzeIfUnwrap',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1220] = { parentId = 1048, typeId = 3198, baseId = 1, txt = 'analyzeIf',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1221] = { parentId = 1048, typeId = 3208, baseId = 1, txt = 'analyzeSwitch',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1426}, children = {}, }
_typeInfoList[1222] = { parentId = 1048, typeId = 3216, baseId = 1, txt = 'analyzeWhile',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1458}, children = {}, }
_typeInfoList[1223] = { parentId = 1048, typeId = 3220, baseId = 1, txt = 'analyzeRepeat',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1482}, children = {}, }
_typeInfoList[1224] = { parentId = 1048, typeId = 3224, baseId = 1, txt = 'analyzeFor',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1512}, children = {}, }
_typeInfoList[1225] = { parentId = 1048, typeId = 3228, baseId = 1, txt = 'analyzeApply',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1544}, children = {}, }
_typeInfoList[1226] = { parentId = 1048, typeId = 3236, baseId = 1, txt = 'analyzeForeach',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 10}, retTypeId = {842}, children = {}, }
_typeInfoList[1227] = { parentId = 1048, typeId = 3242, baseId = 1, txt = 'analyzeRefType',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {1298}, children = {}, }
_typeInfoList[1228] = { parentId = 1, typeId = 3256, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1229] = { parentId = 1, typeId = 3257, nilable = true, orgTypeId = 3256 }
_typeInfoList[1230] = { parentId = 1048, typeId = 3258, baseId = 1, txt = 'analyzeDeclArgList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 3256}, retTypeId = {376}, children = {}, }
_typeInfoList[1231] = { parentId = 106, typeId = 3262, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3264, 3266, 3268}, }
_typeInfoList[1232] = { parentId = 106, typeId = 3263, nilable = true, orgTypeId = 3262 }
_typeInfoList[1233] = { parentId = 3262, typeId = 3264, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {842}, children = {}, }
_typeInfoList[1234] = { parentId = 3262, typeId = 3266, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {708}, children = {}, }
_typeInfoList[1235] = { parentId = 3262, typeId = 3268, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {842, 708}, retTypeId = {}, children = {}, }
_typeInfoList[1236] = { parentId = 1048, typeId = 3270, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {390, 10, 18}, retTypeId = {3262}, children = {}, }
_typeInfoList[1237] = { parentId = 1048, typeId = 3280, baseId = 1, txt = 'analyzeDeclMacro',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[1238] = { parentId = 1048, typeId = 3304, baseId = 1, txt = 'analyzeDeclProto',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[1239] = { parentId = 1048, typeId = 3306, baseId = 1, txt = 'analyzeDecl',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 376, 376}, retTypeId = {843}, children = {}, }
_typeInfoList[1240] = { parentId = 1048, typeId = 3308, baseId = 1, txt = 'analyzeDeclMember',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 376}, retTypeId = {2296}, children = {}, }
_typeInfoList[1241] = { parentId = 1048, typeId = 3310, baseId = 1, txt = 'analyzeDeclMethod',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 376, 376, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[1242] = { parentId = 1048, typeId = 3312, baseId = 1, txt = 'analyzeDeclClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 376}, retTypeId = {844}, children = {}, }
_typeInfoList[1243] = { parentId = 1048, typeId = 3342, baseId = 1, txt = 'addMethod',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 842, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1244] = { parentId = 1048, typeId = 3344, baseId = 1, txt = 'analyzeDeclFunc',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 377, 376, 377}, retTypeId = {842}, children = {}, }
_typeInfoList[1245] = { parentId = 1048, typeId = 3364, baseId = 1, txt = 'analyzeDeclVar',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18, 10, 376}, retTypeId = {842}, children = {}, }
_typeInfoList[1246] = { parentId = 1048, typeId = 3396, baseId = 1, txt = 'analyzeExpList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10}, retTypeId = {1016}, children = {}, }
_typeInfoList[1247] = { parentId = 1048, typeId = 3406, baseId = 1, txt = 'analyzeListConst',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1248] = { parentId = 1048, typeId = 3420, baseId = 1, txt = 'analyzeMapConst',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {2576}, children = {}, }
_typeInfoList[1249] = { parentId = 1048, typeId = 3432, baseId = 1, txt = 'analyzeExpRefItem',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 842}, retTypeId = {842}, children = {}, }
_typeInfoList[1250] = { parentId = 1, typeId = 3436, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {708}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1251] = { parentId = 1, typeId = 3437, nilable = true, orgTypeId = 3436 }
_typeInfoList[1252] = { parentId = 1048, typeId = 3438, baseId = 1, txt = 'checkMatchValType',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {372, 708, 1017, 3436}, retTypeId = {}, children = {}, }
_typeInfoList[1253] = { parentId = 106, typeId = 3440, baseId = 390, txt = 'MacroPaser',
        staticFlag = false, accessMode = 'pri', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3446, 3448, 3450}, }
_typeInfoList[1254] = { parentId = 106, typeId = 3441, nilable = true, orgTypeId = 3440 }
_typeInfoList[1255] = { parentId = 1, typeId = 3444, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {376}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1256] = { parentId = 1, typeId = 3445, nilable = true, orgTypeId = 3444 }
_typeInfoList[1257] = { parentId = 3440, typeId = 3446, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3444, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1258] = { parentId = 3440, typeId = 3448, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {377}, children = {}, }
_typeInfoList[1259] = { parentId = 3440, typeId = 3450, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[1260] = { parentId = 1048, typeId = 3454, baseId = 1, txt = 'evalMacro',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 708, 1017}, retTypeId = {1964}, children = {}, }
_typeInfoList[1261] = { parentId = 1048, typeId = 3478, baseId = 1, txt = 'analyzeExpCont',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 842, 10}, retTypeId = {842}, children = {}, }
_typeInfoList[1262] = { parentId = 1048, typeId = 3482, baseId = 1, txt = 'analyzeExpSymbol',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 376, 18, 843, 10}, retTypeId = {842}, children = {}, }
_typeInfoList[1263] = { parentId = 1048, typeId = 3494, baseId = 1, txt = 'analyzeExpOp2',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376, 842, 12}, retTypeId = {842}, children = {}, }
_typeInfoList[1264] = { parentId = 1048, typeId = 3500, baseId = 1, txt = 'analyzeExpMacroStat',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {1990}, children = {}, }
_typeInfoList[1265] = { parentId = 1048, typeId = 3516, baseId = 1, txt = 'analyzeSuper',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1266] = { parentId = 1048, typeId = 3520, baseId = 1, txt = 'analyzeUnwrap',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1267] = { parentId = 1048, typeId = 3524, baseId = 1, txt = 'analyzeExpUnwrap',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {376}, retTypeId = {842}, children = {}, }
_typeInfoList[1268] = { parentId = 1048, typeId = 3528, baseId = 1, txt = 'analyzeExp',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 12}, retTypeId = {842}, children = {}, }
_typeInfoList[1269] = { parentId = 1048, typeId = 3554, baseId = 1, txt = 'analyzeStatement',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {843}, children = {}, }
_typeInfoList[1270] = { parentId = 1, typeId = 3562, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {842}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1271] = { parentId = 1, typeId = 3563, nilable = true, orgTypeId = 3562 }
_typeInfoList[1272] = { parentId = 1048, typeId = 3564, baseId = 1, txt = 'analyzeStatementList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {3562, 18}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
