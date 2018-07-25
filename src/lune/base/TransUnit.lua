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

local SymbolInfo = {}
moduleObj.SymbolInfo = SymbolInfo
-- none
function SymbolInfo.new( accessMode, name, typeInfo )
  local obj = {}
  setmetatable( obj, { __index = SymbolInfo } )
  if obj.__init then
    obj:__init( accessMode, name, typeInfo )
  end        
  return obj 
 end         
function SymbolInfo:__init( accessMode, name, typeInfo ) 
            
self.accessMode = accessMode
  self.name = name
  self.typeInfo = typeInfo
  end
function SymbolInfo:get_accessMode()
  return self.accessMode
end
function SymbolInfo:get_name()
  return self.name
end
function SymbolInfo:get_typeInfo()
  return self.typeInfo
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
  self.parent = parent or self
  self.symbol2TypeInfoMap = {}
  self.inheritList = inheritList
  self.classFlag = classFlag
end
function Scope:set_ownerTypeInfo( owner )
  self.ownerTypeInfo = owner
end
function Scope:getTypeInfoChild( name )
  do
    local _exp = self.symbol2TypeInfoMap[name]
    if _exp then
    
        return _exp:get_typeInfo()
      end
  end
  
  return nil
end
-- none
function Scope:get_ownerTypeInfo()
  return self.ownerTypeInfo
end
function Scope:get_parent()
  return self.parent
end
function Scope:get_symbol2TypeInfoMap()
  return self.symbol2TypeInfoMap
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
-- none
function TypeInfo:getTxt(  )
  return ""
end
function TypeInfo:serialize( stream )
  return 
end
function TypeInfo:equals( typeInfo )
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

function Scope:getTypeInfoField( name, includeSelfFlag, fromScope )
  if self.classFlag then
    if includeSelfFlag then
      do
        local _exp = self.symbol2TypeInfoMap[name]
        if _exp then
        
            return _exp:canAccess( self, fromScope )
          end
      end
      
    end
    if self.inheritList then
      for __index, scope in pairs( self.inheritList ) do
        local typeInfo = scope:getTypeInfoField( name, true, fromScope )
        
        if typeInfo then
          return typeInfo
        end
      end
    end
  end
  return nil
end

function Scope:getTypeInfo( name, fromScope, onlySameNsFlag )
  local typeInfo = nil
  
  do
    local _exp = self.symbol2TypeInfoMap[name]
    if _exp then
    
        return _exp:canAccess( self, fromScope )
      end
  end
  
  if not onlySameNsFlag then
    if self.inheritList then
      for __index, scope in pairs( self.inheritList ) do
        typeInfo = scope:getTypeInfoField( name, true, fromScope )
        if typeInfo then
          return typeInfo
        end
      end
    end
  end
  if not onlySameNsFlag or not self.ownerTypeInfo then
    if self.parent ~= self then
      return self.parent:getTypeInfo( name, fromScope, onlySameNsFlag )
    end
  end
  if onlySameNsFlag then
    return nil
  end
  return sym2builtInTypeMap[name]
end

function Scope:add( name, typeInfo, accessMode )
  self.symbol2TypeInfoMap[name] = SymbolInfo.new(accessMode, name, typeInfo)
end

function Scope:addClass( name, typeInfo, scope )
  self:add( name, typeInfo, typeInfo:get_accessMode() )
end

local function dumpScopeSub( scope, prefix, readyIdSet )
  do
    local _exp = scope
    if _exp then
    
        if readyIdSet[scope] then
          return 
        end
        readyIdSet[scope] = true
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
            symbolInfo = __map[ symbol ]
            do
              Util.errorLog( string.format( "scope: %s, %s, %s", prefix, _exp, symbol) )
              do
                local _exp = symbolInfo:get_typeInfo():get_scope()
                if _exp then
                
                    dumpScopeSub( _exp, prefix .. "  ", readyIdSet )
                  end
              end
              
            end
          end
        end
        
      end
  end
  
end

local function dumpScope( scope, prefix )
  dumpScopeSub( scope, prefix, {} )
end

local rootTypeInfo = TypeInfo.new(rootScope)

function Scope:getNSTypeInfo(  )
  local scope = self
  
  while scope.ownerTypeInfo ~= rootTypeInfo do
    do
      local _exp = scope.ownerTypeInfo
      if _exp then
      
          return _exp
        end
    end
    
    scope = scope.parent
  end
  return rootTypeInfo
end

function Scope:getClassTypeInfo(  )
  local scope = self
  
  while scope.ownerTypeInfo ~= rootTypeInfo do
    do
      local _exp = scope.ownerTypeInfo
      if _exp then
      
          if _exp:get_kind() == TypeInfoKindClass then
            return _exp
          end
        end
    end
    
    scope = scope.parent
  end
  return rootTypeInfo
end

function SymbolInfo:canAccess( nsScope, fromScope )
  local typeInfo = self:get_typeInfo()
  
  if nsScope == fromScope then
    return typeInfo
  end
  do
    local _switchExp = self:get_accessMode()
    if _switchExp == "pub" or _switchExp == "global" then
      return typeInfo
    elseif _switchExp == "pro" then
      local nsClass = nsScope:getClassTypeInfo(  )
      
      local fromClass = fromScope:getClassTypeInfo(  )
      
      if fromClass:isInheritFrom( nsClass ) then
        return typeInfo
      end
      return nil
    elseif _switchExp == "local" then
      return typeInfo
    elseif _switchExp == "pri" then
      local nsClass = nsScope:getClassTypeInfo(  )
      
      local fromClass = fromScope:getClassTypeInfo(  )
      
      if nsClass == fromClass then
        return typeInfo
      end
      return nil
    end
  end
  
  error( string.format( "illegl accessmode -- %s, %s", self:get_accessMode(), self:get_name()) )
end

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
  elseif txt == "nil" then
    self.nilable = true
    self.nilableTypeInfo = self
    self.orgTypeInfo = self
  elseif not orgTypeInfo then
    if self.parentInfo ~= rootTypeInfo then
      table.insert( self.parentInfo:get_children(), self )
    end
    self.nilable = false
    local hasNilable = false
    
    do
      local _switchExp = (kind )
      if _switchExp == TypeInfoKindPrim or _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindArray or _switchExp == TypeInfoKindMap or _switchExp == TypeInfoKindClass then
        hasNilable = true
      elseif _switchExp == TypeInfoKindFunc then
        if txt == "form" then
          hasNilable = true
        end
      end
    end
    
    if hasNilable then
      self.nilableTypeInfo = NormalTypeInfo.new(nil, baseTypeInfo, self, autoFlag, externalFlag, staticFlag, accessMode, "", parentInfo, typeId + 1, TypeInfoKindNilable, itemTypeInfoList, argTypeInfoList, retTypeInfoList)
    else 
      self.nilableTypeInfo = rootTypeInfo
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
function NormalTypeInfo:equalsSub( typeInfo, depth )
  if not typeInfo then
    return false
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
function NormalTypeInfo:equals( typeInfo )
  return self:equalsSub( typeInfo, 1 )
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
    do
      local _exp = typeDDD
      if _exp then
      
          argTypeList = {_exp}
          retTypeList = {_exp}
        end
    end
    
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
    rootScope:add( typeTxt, info, "pub" )
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

local builtinTypeStem_ = builtinTypeStem:get_nilableTypeInfo() or _luneScript.error( 'unwrap val is nil' )

moduleObj.builtinTypeStem_ = builtinTypeStem_


-- none

local typeInfoListInsert = typeInfoRoot

moduleObj.typeInfoListInsert = typeInfoListInsert

local typeInfoListRemove = typeInfoRoot

moduleObj.typeInfoListRemove = typeInfoListRemove

function NormalTypeInfo:isInheritFrom( other )
  local otherTypeId = other:get_typeId()
  
  if self:get_typeId() == otherTypeId then
    return true
  end
  if self:get_kind() ~= self:get_kind() or self:get_kind() ~= TypeInfoKindClass then
    return false
  end
  local baseTypeInfo = self:get_baseTypeInfo()
  
  while baseTypeInfo ~= rootTypeInfo do
    if otherTypeId == baseTypeInfo:get_typeId() then
      return true
    end
    baseTypeInfo = baseTypeInfo:get_baseTypeInfo()
  end
  -- none
  
  return false
end

function NormalTypeInfo:isSettableFrom( other )
  if not other then
    return false
  end
  if self == builtinTypeStem_ or self == builtinTypeDDD then
    return true
  end
  if self == builtinTypeStem and not other:get_nilable() then
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
      return other:isInheritFrom( self )
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
function Node:get_pos()
  return self.pos
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
  self.usedTokenList = {}
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
  self.scope = Scope.new(self.scope, classFlag, inheritList or {})
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
        local parentScope = scope:get_parent(  )
        
        parentScope:addClass( name, typeInfo, scope )
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
  local prefix = (self.prefix:getLiteral(  )[1] or _luneScript.error( 'unwrap val is nil' ) )
  
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
function _TypeInfo.new( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode )
  local obj = {}
  setmetatable( obj, { __index = _TypeInfo } )
  if obj.__init then
    obj:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode )
  end        
  return obj 
 end         
function _TypeInfo:__init( baseId, itemTypeId, argTypeId, retTypeId, parentId, typeId, txt, kind, staticFlag, nilable, orgTypeId, children, accessMode ) 
            
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
  self.accessMode = accessMode
  end

local _ModuleInfo = {}
function _ModuleInfo.new( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
  local obj = {}
  setmetatable( obj, { __index = _ModuleInfo } )
  if obj.__init then
    obj:__init( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap )
  end        
  return obj 
 end         
function _ModuleInfo:__init( _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap ) 
            
self._typeId2ClassInfoMap = _typeId2ClassInfoMap
  self._typeInfoList = _typeInfoList
  self._varName2InfoMap = _varName2InfoMap
  self._funcName2InfoMap = _funcName2InfoMap
  end

local builtinModuleName2Scope = {}

function TransUnit:registBuiltInScope(  )
  local builtInInfo = {{[""] = {["type"] = {["arg"] = {"stem!"}, ["ret"] = {"str"}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {}}, ["print"] = {["arg"] = {"..."}, ["ret"] = {}}, ["tonumber"] = {["arg"] = {"str"}, ["ret"] = {"real"}}, ["load"] = {["arg"] = {"str"}, ["ret"] = {"form!", "str"}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"stem!"}}, ["_fcall"] = {["arg"] = {"form", "..."}, ["ret"] = {""}}}}, {["ioStream"] = {["read"] = {["methodFlag"] = {}, ["arg"] = {"stem!"}, ["ret"] = {"str!"}}, ["close"] = {["methodFlag"] = {}, ["arg"] = {}, ["ret"] = {}}}}, {["io"] = {["open"] = {["arg"] = {"str", "str!"}, ["ret"] = {"ioStream!"}}, ["popen"] = {["arg"] = {"str"}, ["ret"] = {"ioStream!"}}}}, {["os"] = {["clock"] = {["arg"] = {}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {}}}}, {["string"] = {["find"] = {["arg"] = {"str", "str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["arg"] = {"str", "int"}, ["ret"] = {"int"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"stem!"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int!"}, ["ret"] = {"str"}}}}, {["str"] = {["find"] = {["methodFlag"] = {}, ["arg"] = {"str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["arg"] = {"..."}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["arg"] = {"str"}, ["ret"] = {"stem!"}}, ["gsub"] = {["methodFlag"] = {}, ["arg"] = {"str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["arg"] = {"int", "int!"}, ["ret"] = {"str"}}}}, {["table"] = {["unpack"] = {["arg"] = {"stem"}, ["ret"] = {"..."}}}}, {["List"] = {["insert"] = {["methodFlag"] = {}, ["arg"] = {"stem!"}, ["ret"] = {""}}, ["remove"] = {["methodFlag"] = {}, ["arg"] = {"int!"}, ["ret"] = {""}}}}, {["debug"] = {["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"stem"}}}}, {["_luneScript"] = {["loadModule"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}}}}
  
  local function getTypeInfo( typeName )
    if typeName:find( "!$" ) then
      local orgTypeName = typeName:gsub( "!$", "" )
      
      local typeInfo = self.rootScope:getTypeInfo( orgTypeName, self.rootScope, false ) or _luneScript.error( 'unwrap val is nil' )
      
      return typeInfo:get_nilableTypeInfo()
    end
    return self.rootScope:getTypeInfo( typeName, self.rootScope, false ) or _luneScript.error( 'unwrap val is nil' )
  end
  
  for __index, builtinClassInfo in pairs( builtInInfo ) do
    for name, name2FuncInfo in pairs( builtinClassInfo ) do
      local parentInfo = typeInfoRoot
      
      if name ~= "" then
        parentInfo = self:pushClass( nil, true, name, "pub" )
        builtInTypeIdSet[parentInfo:get_typeId(  )] = true
      end
      if not parentInfo then
        error( "parentInfo is nil" )
      end
      if not builtinModuleName2Scope[name] then
        if name ~= "" and getTypeInfo( name ) then
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
                table.insert( argTypeList, getTypeInfo( argType ) )
              end
              local retTypeList = {}
              
              for __index, retType in pairs( info["ret"] or _luneScript.error( 'unwrap val is nil' ) ) do
                local retTypeInfo = getTypeInfo( retType )
                
                table.insert( retTypeList, retTypeInfo )
              end
              local methodFlag = info["methodFlag"]
              
              self:pushScope( false )
              local typeInfo = NormalTypeInfo.createFunc( self.scope, methodFlag and TypeInfoKindMethod or TypeInfoKindFunc, parentInfo, false, true, not methodFlag, "pub", funcName, argTypeList, retTypeList )
              
              self:popScope(  )
              builtInTypeIdSet[typeInfo:get_typeId(  )] = true
              self.scope:add( funcName, typeInfo, "pub" )
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
  
  error( string.format( "error:%s:%d:%d:(%s) %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, txt, mess ) )
end

function TransUnit:createNoneNode( pos )
  return NoneNode.new(pos, {builtinTypeNone})
end

function TransUnit:pushbackToken( token )
  table.insert( self.pushbackList, token )
  self.currentToken = self.usedTokenList[#self.usedTokenList]
end

local function expandVal( tokenList, val, pos )
  do
    local _exp = val
    if _exp then
    
        do
          local _switchExp = type( _exp )
          if _switchExp == "number" then
            local num = string.format( "%g", _exp)
            
            local kind = Parser.kind.Int
            
            if string.find( num, ".", 1, true ) then
              kind = Parser.kind.Real
            end
            table.insert( tokenList, Parser.Token.new(kind, num, pos) )
          elseif _switchExp == "string" then
            table.insert( tokenList, Parser.Token.new(Parser.kind.Str, string.format( '[[%s]]', _exp), pos) )
          elseif _switchExp == "table" then
            table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "{", _exp), pos) )
            for key, item in pairs( _exp ) do
              expandVal( tokenList, item, pos )
              table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( ",", _exp), pos) )
            end
            table.insert( tokenList, Parser.Token.new(Parser.kind.Dlmt, string.format( "}", _exp), pos) )
          end
        end
        
      end
  end
  
end

function TransUnit:newPushback( tokenList )
  for index = #tokenList, 1, -1 do
    table.insert( self.pushbackList, tokenList[index] )
  end
  self.currentToken = self.usedTokenList[#self.usedTokenList]
end

function TransUnit:getTokenNoErr(  )
  if #self.pushbackList > 0 then
    table.insert( self.usedTokenList, self.currentToken )
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
          
              if  not macroVal then
                local _macroVal = macroVal
                
                self:error( string.format( "unknown macro val %s", nextToken.txt) )
              end
            
          if macroVal.typeInfo == builtinTypeSymbol then
            local txtList = (macroVal.val or _luneScript.error( 'unwrap val is nil' ) )
            
            for index = #txtList, 1, -1 do
              nextToken = Parser.Token.new(nextToken.kind, txtList[index], nextToken.pos)
              self:pushbackToken( nextToken )
            end
          elseif macroVal.typeInfo == builtinTypeStat then
            self:pushbackStr( string.format( "macroVal %s", nextToken.txt), (macroVal.val or _luneScript.error( 'unwrap val is nil' ) ) )
          elseif macroVal.typeInfo:get_kind(  ) == TypeInfoKindArray or macroVal.typeInfo:get_kind(  ) == TypeInfoKindList then
            local strList = (macroVal.val or _luneScript.error( 'unwrap val is nil' ) )
            
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
  
  table.insert( self.usedTokenList, self.currentToken )
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
  self.currentToken = self.usedTokenList[#self.usedTokenList]
  table.remove( self.usedTokenList )
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

function TransUnit:getContinueToken(  )
  local prevToken = self.currentToken
  
      if  not prevToken then
        local _prevToken = prevToken
        
        return self:getToken(  ), false
      end
    
  local token = self:getToken(  )
  
  if prevToken.pos.lineNo ~= token.pos.lineNo or prevToken.pos.column + #prevToken.txt ~= token.pos.column then
    return token, false
  end
  return token, true
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
  self.scope = self.rootScope
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
                
              else
                
                  typeId2Scope[atomInfo.typeId] = newTypeInfo:get_scope()
                  if not newTypeInfo:get_scope() then
                    error( string.format( "not found scope %s %s %s %s %s", parentScope, atomInfo.parentId, atomInfo.typeId, atomInfo.txt, newTypeInfo:getTxt(  )) )
                  end
                  typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
                end
            
        else 
          if atomInfo.kind == TypeInfoKindClass then
            local baseScope = typeId2Scope[atomInfo.baseId] or _luneScript.error( 'unwrap val is nil' )
            
            local scope = Scope.new(parentScope, true, baseScope and {baseScope} or {})
            
            local workTypeInfo = NormalTypeInfo.createClass( scope, baseInfo, parentInfo, true, "pub", atomInfo.txt )
            
            newTypeInfo = workTypeInfo
            typeId2Scope[atomInfo.typeId] = scope
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            parentScope:addClass( atomInfo.txt, workTypeInfo, scope )
          else 
            local scope = nil
            
            if atomInfo.kind == TypeInfoKindFunc or atomInfo.kind == TypeInfoKindMethod then
              scope = Scope.new(parentScope, false, {})
            end
            local workTypeInfo = NormalTypeInfo.create( scope, baseInfo, parentInfo, atomInfo.staticFlag, atomInfo.kind, atomInfo.txt, itemTypeInfo, argTypeInfo, retTypeInfo )
            
            newTypeInfo = workTypeInfo
            typeId2TypeInfo[atomInfo.typeId] = workTypeInfo
            if atomInfo.kind == TypeInfoKindFunc or atomInfo.kind == TypeInfoKindMethod then
              (typeId2Scope[atomInfo.parentId] or _luneScript.error( 'unwrap val is nil' ) ):add( atomInfo.txt, workTypeInfo, atomInfo.accessMode )
              typeId2Scope[atomInfo.typeId] = scope
            end
          end
        end
      end
    else 
      newTypeInfo = self.rootScope:getTypeInfo( atomInfo.txt, self.rootScope, false )
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
        
        scope:add( typeInfo:getTxt(  ), typeInfo, typeInfo:get_accessMode() )
      end
    end
  end
  for typeId, typeInfo in pairs( typeId2TypeInfo ) do
    newId2OldIdMap[typeInfo:get_typeId(  )] = typeId
  end
  local function registMember( classTypeId )
    local classTypeInfo = typeId2TypeInfo[classTypeId] or _luneScript.error( 'unwrap val is nil' )
    
    self:pushClass( nil, true, classTypeInfo:getTxt(  ), "pub" )
    do
      local _exp = moduleInfo._typeId2ClassInfoMap[classTypeId]
      if _exp then
      
          local classInfo = _exp
          
          for fieldName, fieldInfo in pairs( classInfo ) do
            local typeId = fieldInfo.typeId
            
            local fieldTypeInfo = typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' )
            
            self.scope:add( fieldName, fieldTypeInfo, fieldInfo.accessMode )
          end
          for __index, child in pairs( classTypeInfo:get_children(  ) ) do
            if child:get_kind(  ) == TypeInfoKindClass then
              local oldId = newId2OldIdMap[child:get_typeId(  )]
              
              if oldId then
                registMember( oldId or _luneScript.error( 'unwrap val is nil' ) )
              end
            end
          end
        else
      
          self:error( string.format( "not found class -- %d, %s", classTypeId, classTypeInfo:getTxt(  )) )
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
    self.scope:add( varName, typeId2TypeInfo[varInfo.typeId] or _luneScript.error( 'unwrap val is nil' ), "pub" )
  end
  for __index, moduleName in pairs( nameList ) do
    self:popClass(  )
  end
  self.scope = self.moduleScope
  self.scope:add( moduleToken.txt, moduleTypeInfo, "local" )
  self:checkToken( nextToken, ";" )
  if self.moduleScope ~= self.scope then
    self:error( "illegal top scope." )
  end
  return ImportNode.new(token.pos, {builtinTypeNone}, modulePath)
end

function TransUnit:analyzeIfUnwrap( firstToken )
  local exp = self:analyzeExp(  )
  
  local scope = self:pushScope(  )
  
  local expType = exp:get_expType()
  
  if not expType:get_nilable() then
    self:addErrMess( exp:get_pos(), "this is not nilable" )
    scope:add( "_exp", expType, "local" )
  else 
    scope:add( "_exp", expType:get_orgTypeInfo(), "local" )
  end
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
  local nextToken, continueFlag = self:getContinueToken(  )
  
  if continueFlag and nextToken.txt == "!" then
    return self:analyzeIfUnwrap( token )
  end
  self:pushback(  )
  local list = {}
  
  table.insert( list, IfStmtInfo.new("if", self:analyzeExp(  ), self:analyzeBlock( "if" )) )
  nextToken = self:getToken(  )
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

function TransUnit:analyzeFor( firstToken )
  local scope = self:pushScope( false )
  
  local val = self:getToken(  )
  
  if val.kind ~= Parser.kind.Symb then
    self:error( "not symbol" )
  end
  self:checkNextToken( "=" )
  local exp1 = self:analyzeExp(  )
  
  if exp1:get_expType() ~= builtinTypeInt then
    self:addErrMess( exp1:get_pos(), string.format( "exp1 is not int -- %s", exp1:get_expType():getTxt(  )) )
  end
  self.scope:add( val.txt, exp1:get_expType(), "local" )
  self:checkNextToken( "," )
  local exp2 = self:analyzeExp(  )
  
  if exp2:get_expType() ~= builtinTypeInt then
    self:addErrMess( exp2:get_pos(), string.format( "exp2 is not int -- %s", exp2:get_expType():getTxt(  )) )
  end
  local token = self:getToken(  )
  
  local exp3 = nil
  
  if token.txt == "," then
    exp3 = self:analyzeExp(  )
    do
      local _exp = exp3
      if _exp then
      
          if _exp:get_expType() ~= builtinTypeInt then
            self:addErrMess( _exp:get_pos(), string.format( "exp is not int -- %s", _exp:get_expType():getTxt(  )) )
          end
        end
    end
    
  else 
    self:pushback(  )
  end
  local node = ForNode.new(firstToken.pos, {builtinTypeNone}, self:analyzeBlock( "for", scope ), val, exp1, exp2, exp3)
  
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
    scope:add( var.txt, builtinTypeStem, "local" )
  until nextToken.txt ~= ","
  self:checkToken( nextToken, "of" )
  local exp = self:analyzeExp(  )
  
  if exp:get_kind() ~= nodeKindExpCall then
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
      self.scope:add( valSymbol.txt, itemTypeInfoList[2], "local" )
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, itemTypeInfoList[1], "local" )
          end
      end
      
    elseif exp:get_expType():get_kind(  ) == TypeInfoKindList or exp:get_expType():get_kind(  ) == TypeInfoKindArray then
      self.scope:add( valSymbol.txt, itemTypeInfoList[1], "local" )
      do
        local _exp = keySymbol
        if _exp then
        
            self.scope:add( _exp.txt, builtinTypeInt, "local" )
          else
        
            self.scope:add( "__index", builtinTypeInt, "local" )
          end
      end
      
    else 
      self:error( string.format( "unknown kind type of exp for foreach-- %d:%d", exp:get_pos().lineNo, exp:get_pos().column) )
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
  local typeInfo = builtinTypeStem_
  
  self:checkSymbol( token )
  local name = self:analyzeExpSymbol( firstToken, token, "symbol", nil, true )
  
  typeInfo = name:get_expType()
  local continueToken, continueFlag = self:getContinueToken(  )
  
  token = continueToken
  if continueFlag and token.txt == "!" then
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
      
      self.scope:add( argName.txt, refType:get_expType(), "local" )
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
  nextToken = self:getToken(  )
  local ast = nil
  
  if nextToken.txt == "{" then
    local parser = Parser.WrapParser.new(self.parser, string.format( "decl macro %s", nameToken.txt))
    
    for symbol, symbolInfo in pairs( scope:get_symbol2TypeInfoMap() ) do
      scope:add( symbol, symbolInfo:get_typeInfo(), "local" )
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
  
  self.scope:add( nameToken.txt, typeInfo, "local" )
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
  self.scope:add( varName.txt, refType:get_expType(), accessMode )
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
      local retTypeInfo = NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, accessMode, getterName, {}, {memberType} )
      
      self:popScope(  )
      self.scope:add( getterName, retTypeInfo, accessMode )
    end
    local setterName = "set_" .. memberName.txt
    
    accessMode = memberNode:get_setterMode()
    if memberNode:get_setterMode() ~= "none" and not self.scope:getTypeInfoChild( setterName ) then
      self.scope:add( setterName, NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, accessMode, setterName, {memberType}, nil ), accessMode )
      self:popScope(  )
    end
  end
  if not self.scope:getTypeInfoChild( "__init" ) then
    local initTypeInfo = NormalTypeInfo.createFunc( self:pushScope( false ), TypeInfoKindMethod, parentInfo, true, false, false, "pub", "__init", memberTypeList, {} )
    
    self:popScope(  )
    self.scope:add( "__init", initTypeInfo, "pub" )
  end
  self:popClass(  )
  return node
end

function TransUnit:addMethod( className, methodNode, name )
  local classTypeInfo = self.scope:getTypeInfo( className, self.scope, false )
  
  local classNodeInfo = self.typeInfo2ClassNode[classTypeInfo] or _luneScript.error( 'unwrap val is nil' )
  
  classNodeInfo:get_outerMethodSet()[name] = true
  table.insert( classNodeInfo:get_fieldList(), methodNode )
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
    if not name then
      self:error( "name is nil" )
    end
    -- none
    
    local overrideType = self.scope:getTypeInfoField( funcName, false, self.scope )
    
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
      
          if _exp.txt ~= "__init" and self.scope:getTypeInfoField( _exp.txt, false, self.scope ) then
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
    
        scope:get_parent(  ):add( _exp.txt, typeInfo, accessMode )
      end
  end
  
  local node = self:createNoneNode( firstToken.pos )
  
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
  
  local token, continueFlag = self:getContinueToken(  )
  
  if continueFlag and token.txt == "!" then
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
        
        for index, exp in pairs( _exp:get_expList() ) do
          if index == #_exp:get_expList() then
            if exp:get_expType() == builtinTypeDDD then
              for subIndex = index, #varNameList do
                local argType = typeInfoList[subIndex]
                
                if argType ~= builtinTypeNone and not argType:isSettableFrom( builtinTypeStem_ ) then
                  self:addErrMess( firstToken.pos, string.format( "unmatch value type %s(%d) <- %s(%d)", argType:getTxt(  ), argType:get_typeId(), builtinTypeStem_:getTxt(  ), builtinTypeStem_:get_typeId()) )
                end
                if unwrapFlag then
                  table.insert( expTypeList, builtinTypeStem )
                else 
                  table.insert( expTypeList, builtinTypeStem_ )
                end
                table.insert( orgExpTypeList, builtinTypeStem_ )
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
            local expTypeInfo = builtinTypeStem_
            
            if exp:get_expType() == builtinTypeDDD then
              expTypeInfo = builtinTypeStem_
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
  
  local currentClass = self:getCurrentClass(  )
  
  for index, varName in pairs( varNameList ) do
    local typeInfo = typeInfoList[index]
    
    local varInfo = VarInfo.new(varName, varTypeList[index], typeInfo)
    
    table.insert( varList, varInfo )
    if not varTypeList[index] and typeInfo == builtinTypeNil then
      self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
    end
    local sameTypeInfo = self.scope:getTypeInfo( varName.txt, self.scope, true )
    
    do
      local _exp = sameTypeInfo
      if _exp then
      
          table.insert( sameSymbolList, varInfo )
        end
    end
    
    if mode == "let" or mode == "sync" then
      if mode == "let" then
        if self.scope:getTypeInfo( varName.txt, self.scope, true ) then
          self:addErrMess( varName.pos, string.format( "shadowing variable -- %s", varName.txt) )
        end
      end
      self.scope:add( varName.txt, typeInfo, "local" )
    end
  end
  local unwrapBlock = nil
  
  local thenBlock = nil
  
  if unwrapFlag then
    local scope = self:pushScope(  )
    
    for index, varName in pairs( varNameList ) do
      self.scope:add( "_" .. varName.txt, orgExpTypeList[index], "local" )
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
    local exp = self:analyzeExp( skipOp2Flag, 0 )
    
    if not pos then
      pos = exp:get_pos()
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
    local nodeList = (expList or _luneScript.error( 'unwrap val is nil' ) ):get_expList()
    
    for __index, exp in pairs( nodeList ) do
      local expType = exp:get_expType()
      
      if itemTypeInfo == builtinTypeNone then
        itemTypeInfo = expType
      elseif not itemTypeInfo:isSettableFrom( expType ) then
        if expType == builtinTypeNil then
          itemTypeInfo = itemTypeInfo:get_nilableTypeInfo() or _luneScript.error( 'unwrap val is nil' )
        elseif expType:get_nilable() then
          itemTypeInfo = builtinTypeStem_
        else 
          itemTypeInfo = builtinTypeStem
        end
      end
    end
  end
  local kind = nodeKindLiteralArray
  
  local typeInfoList = {builtinTypeNone}
  
  if token.txt == '[' then
    kind = nodeKindLiteralList
    typeInfoList = {NormalTypeInfo.createList( "local", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralListNode.new(token.pos, typeInfoList, expList)
  else 
    typeInfoList = {NormalTypeInfo.createArray( "local", self:getCurrentClass(  ), {itemTypeInfo} )}
    return LiteralArrayNode.new(token.pos, typeInfoList, expList)
  end
end

function TransUnit:analyzeMapConst( token )
  local nextToken = self:getToken(  )
  
  local map = {}
  
  local pairList = {}
  
  local keyTypeInfo = builtinTypeNone
  
  local valTypeInfo = builtinTypeNone
  
  local function getMapKeyValType( pos, keyFlag, typeInfo, expType )
    if expType:get_nilable() then
      if keyFlag then
        self:addErrMess( pos, string.format( "map key can't set a nilable -- %s", expType:getTxt(  )) )
      end
      if expType == builtinTypeNil then
        return typeInfo
      end
      expType = expType:get_orgTypeInfo() or _luneScript.error( 'unwrap val is nil' )
    end
    if not typeInfo:isSettableFrom( expType ) then
      if typeInfo ~= builtinTypeNone then
        typeInfo = builtinTypeStem
      else 
        typeInfo = expType
      end
    end
    return typeInfo
  end
  
  while true do
    if nextToken.txt == "}" then
      break
    end
    self:pushback(  )
    local key = self:analyzeExp(  )
    
    keyTypeInfo = getMapKeyValType( key:get_pos(), true, keyTypeInfo, key:get_expType() )
    self:checkNextToken( ":" )
    local val = self:analyzeExp(  )
    
    valTypeInfo = getMapKeyValType( val:get_pos(), false, valTypeInfo, val:get_expType() )
    table.insert( pairList, PairItem.new(key, val) )
    map[key] = val
    nextToken = self:getToken(  )
    if nextToken.txt ~= "," then
      break
    end
    nextToken = self:getToken(  )
  end
  local typeInfo = NormalTypeInfo.createMap( "local", self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
  
  self:checkToken( nextToken, "}" )
  return LiteralMapNode.new(token.pos, {typeInfo}, map, pairList)
end

function TransUnit:analyzeExpRefItem( token, exp )
  local indexExp = self:analyzeExp(  )
  
  self:checkNextToken( "]" )
  local typeInfo = builtinTypeStem_
  
  local expType = exp:get_expType()
  
  if expType then
    if expType:get_kind() == TypeInfoKindMap then
      typeInfo = expType:get_itemTypeInfoList(  )[2]
      if typeInfo ~= builtinTypeStem_ and not typeInfo:get_nilable() then
        typeInfo = typeInfo:get_nilableTypeInfo()
      end
    elseif expType:get_kind() == TypeInfoKindArray or expType:get_kind() == TypeInfoKindList then
      typeInfo = expType:get_itemTypeInfoList(  )[1]
    elseif expType == builtinTypeString then
      typeInfo = builtinTypeInt
    else 
      self:addErrMess( exp:get_pos(), "could not access with []." )
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
                self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
              end
            end
            break
          elseif #expNodeList == index then
            if expType == builtinTypeDDD then
              for argIndex = index, #argTypeList do
                local workArgType = argTypeList[argIndex]
                
                if not workArgType:isSettableFrom( builtinTypeStem_ ) then
                  self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), argIndex, workArgType:getTxt(  ), expType:getTxt(  )) )
                end
              end
            else 
              for argIndex = index, #argTypeList do
                local argTypeInfo = argTypeList[argIndex]
                
                if not argTypeInfo:isSettableFrom( expType ) then
                  self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), argIndex, argTypeInfo:getTxt(  ), expType:getTxt(  )) )
                end
                expType = builtinTypeNil
              end
            end
            break
          end
          if not argType:isSettableFrom( expType ) then
            self:addErrMess( expNode:get_pos(), string.format( "%s: argument(%d) type mismatch %s <- %s", funcTypeInfo:getTxt(  ), index, argType:getTxt(  ), expType:getTxt(  )) )
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
            local kind = exp:get_kind()
            
            if kind ~= nodeKindLiteralNil and kind ~= nodeKindLiteralChar and kind ~= nodeKindLiteralInt and kind ~= nodeKindLiteralReal and kind ~= nodeKindLiteralArray and kind ~= nodeKindLiteralList and kind ~= nodeKindLiteralMap and kind ~= nodeKindLiteralString and kind ~= nodeKindLiteralBool and kind ~= nodeKindLiteralSymbol and kind ~= nodeKindRefField and kind ~= nodeKindExpMacroStat then
              self:error( "Macro arguments must be literal value." )
            end
          end
        end
      end
  end
  
  local macroInfo = self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )] or _luneScript.error( 'unwrap val is nil' )
  
  local argVal = {}
  
  do
    local _exp = expList
    if _exp then
    
        for __index, argNode in pairs( _exp:get_expList(  ) ) do
          local valList, typeInfoList = argNode:getLiteral(  )
          
          local val = valList[1]
          
          local typeInfo = typeInfoList[1]
          
          table.insert( argVal, val )
        end
      end
  end
  
  local func = macroInfo.func
  
  local macroVars = func( table.unpack( argVal ) )
  
  for __index, name in pairs( macroVars._names ) do
    local valInfo = macroInfo.symbol2MacroValInfoMap[name] or _luneScript.error( 'unwrap val is nil' )
    
    local typeInfo = valInfo and valInfo.typeInfo or builtinTypeStem_
    
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
  local parser = MacroPaser.new(macroInfo.declInfo:get_tokenList(), string.format( "macro %s", macroTypeInfo:getTxt(  )))
  
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
        self:checkMatchValType( exp:get_pos(), funcTypeInfo, expList, genericTypeList )
        if macroFlag then
          self.macroMode = "none"
          exp = self:evalMacro( firstToken, funcTypeInfo, expList )
        else 
          do
            local _switchExp = (exp:get_expType():get_kind() )
            if _switchExp == TypeInfoKindMethod or _switchExp == TypeInfoKindFunc then
            else 
              self:error( string.format( "can't call the type -- %s", exp:get_expType():getTxt(  )) )
            end
          end
          
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
          
          self:error( "prefix is nil" )
        else
          
            if self.macroMode == "analyze" then
              exp = RefFieldNode.new(firstToken.pos, {builtinTypeSymbol}, token, prefixExp or _luneScript.error( 'unwrap val is nil' ))
            else 
              local typeInfo = builtinTypeStem_
              
              local prefixExpType = prefixExp:get_expType()
              
              if not prefixExpType then
                self:error( "unknown prefix type: " .. getNodeKindName( prefixExp:get_kind() ) )
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
                  typeInfo = classScope:getTypeInfo( string.format( "get_%s", token.txt), self.scope, false )
                  do
                    local _exp = typeInfo
                    if _exp then
                    
                        if (_exp:get_kind(  ) == TypeInfoKindMethod ) then
                          local retTypeList = _exp:get_retTypeInfoList(  )
                          
                          getterTypeInfo = _exp
                          typeInfo = retTypeList[1]
                        end
                      end
                  end
                  
                end
                if not getterTypeInfo then
                  typeInfo = classScope:getTypeInfo( token.txt, self.scope, false )
                end
                if not typeInfo then
                  for name, val in pairs( classScope:get_symbol2TypeInfoMap() ) do
                    Util.errorLog( string.format( "debug: %s, %s", name, val) )
                  end
                  self:error( string.format( "not found field typeInfo: %s.%s", className, token.txt ) )
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
          end
      
  elseif mode == "symbol" then
    if self.macroMode == "analyze" then
      exp = LiteralSymbolNode.new(firstToken.pos, {builtinTypeSymbol}, token)
    else 
      local typeInfo = self.scope:getTypeInfo( token.txt, self.scope, false )
      
      if not typeInfo and token.txt == "self" then
        typeInfo = self:getCurrentClass(  )
      end
      do
        local _exp = typeInfo
        if _exp then
        
            if _exp == builtinTypeSymbol then
              skipFlag = true
            end
            exp = ExpRefNode.new(firstToken.pos, {_exp}, token)
          else
        
            self:error( "not found type -- " .. token.txt )
          end
      end
      
    end
  elseif mode == "fn" then
    exp = self:analyzeDeclFunc( false, "local", false, nil, token, nil )
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
      local castType = self:analyzeRefType( "local" )
      
      if exp:get_expType():get_nilable() and not castType:get_expType():get_nilable() then
        self:addErrMess( firstToken.pos, string.format( "can't cast from nilable to not nilable  -- %s->%s", exp:get_expType():getTxt(  ), castType:get_expType():getTxt(  )) )
      end
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
              retType = builtinTypeStem_
            end
          elseif _switchExp == "and" then
            retType = exp2Type
          elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" then
            if exp1Type ~= builtinTypeInt and exp1Type ~= builtinTypeReal or exp2Type ~= builtinTypeInt and exp2Type ~= builtinTypeReal then
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
      
      if token.txt == ",," and exp:get_kind() == nodeKindExpRef then
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
  local nextToken, continueFlag = self:getContinueToken(  )
  
  if not continueFlag or nextToken.txt ~= "!" then
    self:pushback(  )
    self:pushbackToken( firstToken )
    local exp = self:analyzeExp(  )
    
    self:checkNextToken( ";" )
    return StmtExpNode.new(nextToken.pos, {builtinTypeNone}, exp)
  end
  self:pushback(  )
  return self:analyzeDeclVar( "unwrap", "local", false, firstToken )
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
  local unwrapType = builtinTypeStem_
  
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
            self:addErrMess( _exp:get_pos(), string.format( "unmatch type: %s <- %s", unwrapType:getTxt(  ), insType:getTxt(  )) )
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
    exp = self:analyzeRefType( "local" )
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
      
    self:checkMatchValType( exp:get_pos(), initTypeInfo, argList, exp:get_expType():get_itemTypeInfoList() )
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
    
  local statement = self:analyzeDecl( "local", false, token, token )
  
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
      statement = self:analyzeDeclVar( "sync", "local", false, token )
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
      statement = StmtExpNode.new(exp:get_pos(), {builtinTypeNone}, exp)
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

return moduleObj
