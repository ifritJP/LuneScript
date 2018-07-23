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

function Scope:getTypeInfo( name, fromScope )
  local typeInfo = nil
  
  do
    local _exp = self.symbol2TypeInfoMap[name]
    if _exp then
    
        return _exp:canAccess( self, fromScope )
      end
  end
  
  if self.inheritList then
    for __index, scope in pairs( self.inheritList ) do
      typeInfo = scope:getTypeInfoField( name, true, fromScope )
      if typeInfo then
        return typeInfo
      end
    end
  end
  if self.parent ~= self then
    return self.parent:getTypeInfo( name, fromScope )
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
        self.nilableTypeInfo = NormalTypeInfo.new(nil, baseTypeInfo, self, autoFlag, externalFlag, staticFlag, accessMode, "", parentInfo, typeId + 1, TypeInfoKindNilable, itemTypeInfoList, argTypeInfoList, retTypeInfoList)
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
  local builtInInfo = {[""] = {["type"] = {["arg"] = {"stem"}, ["ret"] = {"str"}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {}}, ["print"] = {["arg"] = {"..."}, ["ret"] = {}}, ["tonumber"] = {["arg"] = {"str"}, ["ret"] = {"real"}}, ["load"] = {["arg"] = {"str"}, ["ret"] = {"stem", "str"}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}, ["_fcall"] = {["arg"] = {"form", "..."}, ["ret"] = {""}}}, ["io"] = {["open"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}}, ["os"] = {["clock"] = {["arg"] = {}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {}}}, ["string"] = {["find"] = {["arg"] = {"str", "str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["arg"] = {"str", "int"}, ["ret"] = {"int"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"stem"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int!"}, ["ret"] = {"str"}}}, ["str"] = {["find"] = {["methodFlag"] = {}, ["arg"] = {"str", "int!", "bool!"}, ["ret"] = {"int", "int"}}, ["byte"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"int"}}, ["format"] = {["methodFlag"] = {}, ["arg"] = {"..."}, ["ret"] = {"str"}}, ["rep"] = {["methodFlag"] = {}, ["arg"] = {"int"}, ["ret"] = {"str"}}, ["gmatch"] = {["methodFlag"] = {}, ["arg"] = {"str"}, ["ret"] = {"stem"}}, ["gsub"] = {["methodFlag"] = {}, ["arg"] = {"str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["methodFlag"] = {}, ["arg"] = {"int", "int!"}, ["ret"] = {"str"}}}, ["table"] = {["unpack"] = {["arg"] = {"stem"}, ["ret"] = {"stem"}}}, ["List"] = {["insert"] = {["methodFlag"] = {}, ["arg"] = {"stem"}, ["ret"] = {""}}, ["remove"] = {["methodFlag"] = {}, ["arg"] = {"int!"}, ["ret"] = {""}}}, ["debug"] = {["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"stem"}}}, ["_luneScript"] = {["loadModule"] = {["arg"] = {"str"}, ["ret"] = {"stem"}}}}
  
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
          parentInfo = self:pushClass( nil, true, name, "pub" )
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
                  table.insert( argTypeList, self.rootScope:getTypeInfo( argType, self.rootScope ) or _luneScript.error( 'unwrap val is nil' ) )
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
                
              end
            
          typeId2Scope[atomInfo.typeId] = newTypeInfo:get_scope()
          if not newTypeInfo:get_scope() then
            error( string.format( "not found scope %s %s %s %s %s", parentScope, atomInfo.parentId, atomInfo.typeId, atomInfo.txt, newTypeInfo:getTxt(  )) )
          end
          typeId2TypeInfo[atomInfo.typeId] = newTypeInfo
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
    local classInfo = moduleInfo._typeId2ClassInfoMap[classTypeId]
    
    if classInfo then
      for fieldName, fieldInfo in pairs( classInfo ) do
        local typeId = fieldInfo.typeId
        
        local fieldTypeInfo = typeId2TypeInfo[typeId] or _luneScript.error( 'unwrap val is nil' )
        
        self.scope:add( fieldName, fieldTypeInfo, fieldInfo.accessMode )
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
  local nextToken = self:getToken(  )
  
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
  local classTypeInfo = self.scope:getTypeInfo( className, self.scope )
  
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
    local name = name
    
        if  not name then
          local _name = name
          
          self:error( "name is nil" )
        end
      
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
        
        for index, exp in pairs( _exp:get_expList() ) do
          if index == #_exp:get_expList() then
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
  
  local currentClass = self:getCurrentClass(  )
  
  for index, varName in pairs( varNameList ) do
    local typeInfo = typeInfoList[index]
    
    local varInfo = VarInfo.new(varName, varTypeList[index], typeInfo)
    
    table.insert( varList, varInfo )
    if not varTypeList[index] and typeInfo == builtinTypeNil then
      self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
    end
    local sameTypeInfo = self.scope:getTypeInfo( varName.txt, self.scope )
    
    do
      local _exp = sameTypeInfo
      if _exp then
      
          table.insert( sameSymbolList, varInfo )
        end
    end
    
    if mode == "let" or mode == "sync" then
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
  local typeInfo = NormalTypeInfo.createMap( "local", self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
  
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
                
                if not workArgType:isSettableFrom( builtinTypeStem ) then
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
          typeInfo = classScope:getTypeInfo( string.format( "get_%s", token.txt), self.scope )
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
          typeInfo = classScope:getTypeInfo( token.txt, self.scope )
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
  elseif mode == "symbol" then
    if self.macroMode == "analyze" then
      exp = LiteralSymbolNode.new(firstToken.pos, {builtinTypeSymbol}, token)
    else 
      local typeInfo = self.scope:getTypeInfo( token.txt, self.scope )
      
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
  local nextToken = self:getToken(  )
  
  if nextToken.txt ~= "!" then
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

----- meta -----
local _typeId2ClassInfoMap = {}
moduleObj._typeId2ClassInfoMap = _typeId2ClassInfoMap
local _className2InfoMap = {}
moduleObj._className2InfoMap = _className2InfoMap
do
  local _classInfo3254 = {}
  _className2InfoMap.ASTInfo = _classInfo3254
  _typeId2ClassInfoMap[ 3254 ] = _classInfo3254
  end
do
  local _classInfo1526 = {}
  _className2InfoMap.ApplyNode = _classInfo1526
  _typeId2ClassInfoMap[ 1526 ] = _classInfo1526
  end
do
  local _classInfo1308 = {}
  _className2InfoMap.BlockNode = _classInfo1308
  _typeId2ClassInfoMap[ 1308 ] = _classInfo1308
  end
do
  local _classInfo1642 = {}
  _className2InfoMap.BreakNode = _classInfo1642
  _typeId2ClassInfoMap[ 1642 ] = _classInfo1642
  end
do
  local _classInfo1392 = {}
  _className2InfoMap.CaseInfo = _classInfo1392
  _typeId2ClassInfoMap[ 1392 ] = _classInfo1392
  end
do
  local _classInfo2332 = {}
  _className2InfoMap.DeclArgDDDNode = _classInfo2332
  _typeId2ClassInfoMap[ 2332 ] = _classInfo2332
  end
do
  local _classInfo2312 = {}
  _className2InfoMap.DeclArgNode = _classInfo2312
  _typeId2ClassInfoMap[ 2312 ] = _classInfo2312
  end
do
  local _classInfo816 = {}
  _className2InfoMap.DeclClassNode = _classInfo816
  _typeId2ClassInfoMap[ 816 ] = _classInfo816
  end
do
  local _classInfo2226 = {}
  _className2InfoMap.DeclConstrNode = _classInfo2226
  _typeId2ClassInfoMap[ 2226 ] = _classInfo2226
  end
do
  local _classInfo2156 = {}
  _className2InfoMap.DeclFuncInfo = _classInfo2156
  _typeId2ClassInfoMap[ 2156 ] = _classInfo2156
  end
do
  local _classInfo2186 = {}
  _className2InfoMap.DeclFuncNode = _classInfo2186
  _typeId2ClassInfoMap[ 2186 ] = _classInfo2186
  end
do
  local _classInfo996 = {}
  _className2InfoMap.DeclMacroInfo = _classInfo996
  _typeId2ClassInfoMap[ 996 ] = _classInfo996
  end
do
  local _classInfo2406 = {}
  _className2InfoMap.DeclMacroNode = _classInfo2406
  _typeId2ClassInfoMap[ 2406 ] = _classInfo2406
  end
do
  local _classInfo2280 = {}
  _className2InfoMap.DeclMemberNode = _classInfo2280
  _typeId2ClassInfoMap[ 2280 ] = _classInfo2280
  end
do
  local _classInfo2206 = {}
  _className2InfoMap.DeclMethodNode = _classInfo2206
  _typeId2ClassInfoMap[ 2206 ] = _classInfo2206
  end
do
  local _classInfo2102 = {}
  _className2InfoMap.DeclVarNode = _classInfo2102
  _typeId2ClassInfoMap[ 2102 ] = _classInfo2102
  end
do
  local _classInfo1886 = {}
  _className2InfoMap.ExpCallNode = _classInfo1886
  _typeId2ClassInfoMap[ 1886 ] = _classInfo1886
  end
do
  local _classInfo2248 = {}
  _className2InfoMap.ExpCallSuperNode = _classInfo2248
  _typeId2ClassInfoMap[ 2248 ] = _classInfo2248
  end
do
  local _classInfo1812 = {}
  _className2InfoMap.ExpCastNode = _classInfo1812
  _typeId2ClassInfoMap[ 1812 ] = _classInfo1812
  end
do
  local _classInfo1908 = {}
  _className2InfoMap.ExpDDDNode = _classInfo1908
  _typeId2ClassInfoMap[ 1908 ] = _classInfo1908
  end
do
  local _classInfo994 = {}
  _className2InfoMap.ExpListNode = _classInfo994
  _typeId2ClassInfoMap[ 994 ] = _classInfo994
  end
do
  local _classInfo1948 = {}
  _className2InfoMap.ExpMacroExpNode = _classInfo1948
  _typeId2ClassInfoMap[ 1948 ] = _classInfo1948
  end
do
  local _classInfo1974 = {}
  _className2InfoMap.ExpMacroStatNode = _classInfo1974
  _typeId2ClassInfoMap[ 1974 ] = _classInfo1974
  end
do
  local _classInfo1662 = {}
  _className2InfoMap.ExpNewNode = _classInfo1662
  _typeId2ClassInfoMap[ 1662 ] = _classInfo1662
  end
do
  local _classInfo1836 = {}
  _className2InfoMap.ExpOp1Node = _classInfo1836
  _typeId2ClassInfoMap[ 1836 ] = _classInfo1836
  end
do
  local _classInfo1732 = {}
  _className2InfoMap.ExpOp2Node = _classInfo1732
  _typeId2ClassInfoMap[ 1732 ] = _classInfo1732
  end
do
  local _classInfo1928 = {}
  _className2InfoMap.ExpParenNode = _classInfo1928
  _typeId2ClassInfoMap[ 1928 ] = _classInfo1928
  end
do
  local _classInfo1862 = {}
  _className2InfoMap.ExpRefItemNode = _classInfo1862
  _typeId2ClassInfoMap[ 1862 ] = _classInfo1862
  end
do
  local _classInfo1708 = {}
  _className2InfoMap.ExpRefNode = _classInfo1708
  _typeId2ClassInfoMap[ 1708 ] = _classInfo1708
  end
do
  local _classInfo1686 = {}
  _className2InfoMap.ExpUnwrapNode = _classInfo1686
  _typeId2ClassInfoMap[ 1686 ] = _classInfo1686
  end
do
  local _classInfo958 = {}
  _className2InfoMap.Filter = _classInfo958
  _typeId2ClassInfoMap[ 958 ] = _classInfo958
  end
do
  local _classInfo1494 = {}
  _className2InfoMap.ForNode = _classInfo1494
  _typeId2ClassInfoMap[ 1494 ] = _classInfo1494
  end
do
  local _classInfo1562 = {}
  _className2InfoMap.ForeachNode = _classInfo1562
  _typeId2ClassInfoMap[ 1562 ] = _classInfo1562
  end
do
  local _classInfo1596 = {}
  _className2InfoMap.ForsortNode = _classInfo1596
  _typeId2ClassInfoMap[ 1596 ] = _classInfo1596
  end
do
  local _classInfo2048 = {}
  _className2InfoMap.GetFieldNode = _classInfo2048
  _typeId2ClassInfoMap[ 2048 ] = _classInfo2048
  end
do
  local _classInfo1346 = {}
  _className2InfoMap.IfNode = _classInfo1346
  _typeId2ClassInfoMap[ 1346 ] = _classInfo1346
  end
do
  local _classInfo1332 = {}
  _className2InfoMap.IfStmtInfo = _classInfo1332
  _typeId2ClassInfoMap[ 1332 ] = _classInfo1332
  end
do
  local _classInfo1788 = {}
  _className2InfoMap.IfUnwrapNode = _classInfo1788
  _typeId2ClassInfoMap[ 1788 ] = _classInfo1788
  end
do
  local _classInfo1218 = {}
  _className2InfoMap.ImportNode = _classInfo1218
  _typeId2ClassInfoMap[ 1218 ] = _classInfo1218
  end
do
  local _classInfo2514 = {}
  _className2InfoMap.LiteralArrayNode = _classInfo2514
  _typeId2ClassInfoMap[ 2514 ] = _classInfo2514
  end
do
  local _classInfo2628 = {}
  _className2InfoMap.LiteralBoolNode = _classInfo2628
  _typeId2ClassInfoMap[ 2628 ] = _classInfo2628
  end
do
  local _classInfo2444 = {}
  _className2InfoMap.LiteralCharNode = _classInfo2444
  _typeId2ClassInfoMap[ 2444 ] = _classInfo2444
  end
do
  local _classInfo2468 = {}
  _className2InfoMap.LiteralIntNode = _classInfo2468
  _typeId2ClassInfoMap[ 2468 ] = _classInfo2468
  end
do
  local _classInfo2534 = {}
  _className2InfoMap.LiteralListNode = _classInfo2534
  _typeId2ClassInfoMap[ 2534 ] = _classInfo2534
  end
do
  local _classInfo2564 = {}
  _className2InfoMap.LiteralMapNode = _classInfo2564
  _typeId2ClassInfoMap[ 2564 ] = _classInfo2564
  end
do
  local _classInfo2424 = {}
  _className2InfoMap.LiteralNilNode = _classInfo2424
  _typeId2ClassInfoMap[ 2424 ] = _classInfo2424
  end
do
  local _classInfo2492 = {}
  _className2InfoMap.LiteralRealNode = _classInfo2492
  _typeId2ClassInfoMap[ 2492 ] = _classInfo2492
  end
do
  local _classInfo2600 = {}
  _className2InfoMap.LiteralStringNode = _classInfo2600
  _typeId2ClassInfoMap[ 2600 ] = _classInfo2600
  end
do
  local _classInfo2648 = {}
  _className2InfoMap.LiteralSymbolNode = _classInfo2648
  _typeId2ClassInfoMap[ 2648 ] = _classInfo2648
  end
do
  local _classInfo992 = {}
  _className2InfoMap.MacroEval = _classInfo992
  _typeId2ClassInfoMap[ 992 ] = _classInfo992
  end
do
  local _classInfo988 = {}
  _className2InfoMap.NamespaceInfo = _classInfo988
  _typeId2ClassInfoMap[ 988 ] = _classInfo988
  _classInfo988.name = {
    name='name', staticFlag = false, accessMode = 'pub', typeId = 18 }
  _classInfo988.scope = {
    name='scope', staticFlag = false, accessMode = 'pub', typeId = 686 }
  _classInfo988.typeInfo = {
    name='typeInfo', staticFlag = false, accessMode = 'pub', typeId = 660 }
  end
do
  local _classInfo814 = {}
  _className2InfoMap.Node = _classInfo814
  _typeId2ClassInfoMap[ 814 ] = _classInfo814
  end
do
  local _classInfo1200 = {}
  _className2InfoMap.NoneNode = _classInfo1200
  _typeId2ClassInfoMap[ 1200 ] = _classInfo1200
  end
do
  local _classInfo818 = {}
  _className2InfoMap.NormalTypeInfo = _classInfo818
  _typeId2ClassInfoMap[ 818 ] = _classInfo818
  end
do
  local _classInfo676 = {}
  _className2InfoMap.OutStream = _classInfo676
  _typeId2ClassInfoMap[ 676 ] = _classInfo676
  end
do
  local _classInfo2550 = {}
  _className2InfoMap.PairItem = _classInfo2550
  _typeId2ClassInfoMap[ 2550 ] = _classInfo2550
  end
do
  local _classInfo2022 = {}
  _className2InfoMap.RefFieldNode = _classInfo2022
  _typeId2ClassInfoMap[ 2022 ] = _classInfo2022
  end
do
  local _classInfo1280 = {}
  _className2InfoMap.RefTypeNode = _classInfo1280
  _typeId2ClassInfoMap[ 1280 ] = _classInfo1280
  end
do
  local _classInfo1464 = {}
  _className2InfoMap.RepeatNode = _classInfo1464
  _typeId2ClassInfoMap[ 1464 ] = _classInfo1464
  end
do
  local _classInfo1624 = {}
  _className2InfoMap.ReturnNode = _classInfo1624
  _typeId2ClassInfoMap[ 1624 ] = _classInfo1624
  end
do
  local _classInfo1240 = {}
  _className2InfoMap.RootNode = _classInfo1240
  _typeId2ClassInfoMap[ 1240 ] = _classInfo1240
  end
do
  local _classInfo686 = {}
  _className2InfoMap.Scope = _classInfo686
  _typeId2ClassInfoMap[ 686 ] = _classInfo686
  end
do
  local _classInfo2000 = {}
  _className2InfoMap.StmtExpNode = _classInfo2000
  _typeId2ClassInfoMap[ 2000 ] = _classInfo2000
  end
do
  local _classInfo1408 = {}
  _className2InfoMap.SwitchNode = _classInfo1408
  _typeId2ClassInfoMap[ 1408 ] = _classInfo1408
  end
do
  local _classInfo688 = {}
  _className2InfoMap.SymbolInfo = _classInfo688
  _typeId2ClassInfoMap[ 688 ] = _classInfo688
  end
do
  local _classInfo1026 = {}
  _className2InfoMap.TransUnit = _classInfo1026
  _typeId2ClassInfoMap[ 1026 ] = _classInfo1026
  end
do
  local _classInfo660 = {}
  _className2InfoMap.TypeInfo = _classInfo660
  _typeId2ClassInfoMap[ 660 ] = _classInfo660
  end
do
  local _classInfo1760 = {}
  _className2InfoMap.UnwrapSetNode = _classInfo1760
  _typeId2ClassInfoMap[ 1760 ] = _classInfo1760
  end
do
  local _classInfo2068 = {}
  _className2InfoMap.VarInfo = _classInfo2068
  _typeId2ClassInfoMap[ 2068 ] = _classInfo2068
  end
do
  local _classInfo1440 = {}
  _className2InfoMap.WhileNode = _classInfo1440
  _typeId2ClassInfoMap[ 1440 ] = _classInfo1440
  end
do
  local _classInfo238 = {}
  _className2InfoMap.Parser = _classInfo238
  _typeId2ClassInfoMap[ 238 ] = _classInfo238
  end
do
  local _classInfo216 = {}
  _className2InfoMap.Position = _classInfo216
  _typeId2ClassInfoMap[ 216 ] = _classInfo216
  _classInfo216.column = {
    name='column', staticFlag = false, accessMode = 'pub', typeId = 12 }
  _classInfo216.lineNo = {
    name='lineNo', staticFlag = false, accessMode = 'pub', typeId = 12 }
  end
do
  local _classInfo204 = {}
  _className2InfoMap.Stream = _classInfo204
  _typeId2ClassInfoMap[ 204 ] = _classInfo204
  end
do
  local _classInfo254 = {}
  _className2InfoMap.StreamParser = _classInfo254
  _typeId2ClassInfoMap[ 254 ] = _classInfo254
  _classInfo254.create = {
    name='create', staticFlag = true, accessMode = 'pub', typeId = 616 }
  end
do
  local _classInfo220 = {}
  _className2InfoMap.Token = _classInfo220
  _typeId2ClassInfoMap[ 220 ] = _classInfo220
  _classInfo220.kind = {
    name='kind', staticFlag = false, accessMode = 'pub', typeId = 12 }
  _classInfo220.pos = {
    name='pos', staticFlag = false, accessMode = 'pub', typeId = 216 }
  _classInfo220.txt = {
    name='txt', staticFlag = false, accessMode = 'pub', typeId = 18 }
  end
do
  local _classInfo210 = {}
  _className2InfoMap.TxtStream = _classInfo210
  _typeId2ClassInfoMap[ 210 ] = _classInfo210
  end
do
  local _classInfo424 = {}
  _className2InfoMap.Util = _classInfo424
  _typeId2ClassInfoMap[ 424 ] = _classInfo424
  _classInfo424.debugLog = {
    name='debugLog', staticFlag = true, accessMode = 'pub', typeId = 656 }
  _classInfo424.errorLog = {
    name='errorLog', staticFlag = true, accessMode = 'pub', typeId = 654 }
  _classInfo424.memStream = {
    name='memStream', staticFlag = false, accessMode = 'pub', typeId = 472 }
  _classInfo424.outStream = {
    name='outStream', staticFlag = false, accessMode = 'pub', typeId = 466 }
  _classInfo424.profile = {
    name='profile', staticFlag = true, accessMode = 'pub', typeId = 658 }
  end
do
  local _classInfo246 = {}
  _className2InfoMap.WrapParser = _classInfo246
  _typeId2ClassInfoMap[ 246 ] = _classInfo246
  end
do
  local _classInfo104 = {}
  _className2InfoMap.base = _classInfo104
  _typeId2ClassInfoMap[ 104 ] = _classInfo104
  _classInfo104.Parser = {
    name='Parser', staticFlag = false, accessMode = 'pub', typeId = 108 }
  _classInfo104.TransUnit = {
    name='TransUnit', staticFlag = false, accessMode = 'pub', typeId = 106 }
  _classInfo104.Util = {
    name='Util', staticFlag = false, accessMode = 'pub', typeId = 424 }
  end
do
  local _classInfo102 = {}
  _className2InfoMap.lune = _classInfo102
  _typeId2ClassInfoMap[ 102 ] = _classInfo102
  _classInfo102.base = {
    name='base', staticFlag = false, accessMode = 'pub', typeId = 104 }
  end
do
  local _classInfo472 = {}
  _className2InfoMap.memStream = _classInfo472
  _typeId2ClassInfoMap[ 472 ] = _classInfo472
  end
do
  local _classInfo466 = {}
  _className2InfoMap.outStream = _classInfo466
  _typeId2ClassInfoMap[ 466 ] = _classInfo466
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
  name='builtinTypeArray', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeBool = {
  name='builtinTypeBool', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeChar = {
  name='builtinTypeChar', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeDDD = {
  name='builtinTypeDDD', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeForm = {
  name='builtinTypeForm', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeInt = {
  name='builtinTypeInt', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeList = {
  name='builtinTypeList', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeMap = {
  name='builtinTypeMap', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeNil = {
  name='builtinTypeNil', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeNone = {
  name='builtinTypeNone', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeReal = {
  name='builtinTypeReal', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeStat = {
  name='builtinTypeStat', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeStem = {
  name='builtinTypeStem', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeString = {
  name='builtinTypeString', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.builtinTypeSymbol = {
  name='builtinTypeSymbol', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.nodeKind = {
  name='nodeKind', accessMode = 'pub', typeId = 1168 }
_varName2InfoMap.rootTypeId = {
  name='rootTypeId', accessMode = 'pub', typeId = 12 }
_varName2InfoMap.typeInfoKind = {
  name='typeInfoKind', accessMode = 'pub', typeId = 662 }
_varName2InfoMap.typeInfoListInsert = {
  name='typeInfoListInsert', accessMode = 'pub', typeId = 660 }
_varName2InfoMap.typeInfoListRemove = {
  name='typeInfoListRemove', accessMode = 'pub', typeId = 660 }
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
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {106, 108, 424}, }
_typeInfoList[9] = { parentId = 102, typeId = 105, nilable = true, orgTypeId = 104 }
_typeInfoList[10] = { parentId = 104, typeId = 106, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {660, 674, 676, 686, 688, 814, 816, 818, 958, 988, 992, 994, 996, 1026, 1174, 1200, 1218, 1240, 1280, 1308, 1332, 1346, 1392, 1408, 1440, 1464, 1494, 1526, 1562, 1596, 1624, 1642, 1662, 1686, 1708, 1732, 1760, 1788, 1812, 1836, 1862, 1886, 1908, 1928, 1948, 1974, 2000, 2022, 2048, 2068, 2102, 2156, 2186, 2206, 2226, 2248, 2280, 2312, 2332, 2406, 2424, 2444, 2468, 2492, 2514, 2534, 2550, 2564, 2600, 2628, 2648, 3254}, }
_typeInfoList[11] = { parentId = 104, typeId = 107, nilable = true, orgTypeId = 106 }
_typeInfoList[12] = { parentId = 104, typeId = 108, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {204, 210, 216, 220, 238, 246, 254, 310, 312, 314, 342, 354, 398, 400, 402, 404, 422, 498, 540, 542, 544, 546, 564, 576, 618, 620, 622, 624, 642}, }
_typeInfoList[13] = { parentId = 104, typeId = 109, nilable = true, orgTypeId = 108 }
_typeInfoList[14] = { parentId = 1, typeId = 150, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'local', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[15] = { parentId = 1, typeId = 151, nilable = true, orgTypeId = 150 }
_typeInfoList[16] = { parentId = 1, typeId = 152, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'local', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[17] = { parentId = 1, typeId = 153, nilable = true, orgTypeId = 152 }
_typeInfoList[18] = { parentId = 1, typeId = 154, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'local', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[19] = { parentId = 1, typeId = 155, nilable = true, orgTypeId = 154 }
_typeInfoList[20] = { parentId = 1, typeId = 156, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[21] = { parentId = 1, typeId = 157, nilable = true, orgTypeId = 156 }
_typeInfoList[22] = { parentId = 1, typeId = 158, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'local', kind = 5, itemTypeId = {18, 156}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[23] = { parentId = 1, typeId = 159, nilable = true, orgTypeId = 158 }
_typeInfoList[24] = { parentId = 108, typeId = 160, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {150, 152, 154, 158}, children = {}, }
_typeInfoList[25] = { parentId = 108, typeId = 204, baseId = 1, txt = 'Stream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {206, 208, 356, 358, 500, 502, 578, 580}, }
_typeInfoList[26] = { parentId = 108, typeId = 205, nilable = true, orgTypeId = 204 }
_typeInfoList[27] = { parentId = 204, typeId = 206, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[28] = { parentId = 204, typeId = 208, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[29] = { parentId = 108, typeId = 210, baseId = 204, txt = 'TxtStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {212, 214, 360, 362, 504, 506, 582, 584}, }
_typeInfoList[30] = { parentId = 108, typeId = 211, nilable = true, orgTypeId = 210 }
_typeInfoList[31] = { parentId = 210, typeId = 212, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[32] = { parentId = 210, typeId = 214, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[33] = { parentId = 108, typeId = 216, baseId = 1, txt = 'Position',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {218, 364, 508, 586}, }
_typeInfoList[34] = { parentId = 108, typeId = 217, nilable = true, orgTypeId = 216 }
_typeInfoList[35] = { parentId = 216, typeId = 218, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[36] = { parentId = 108, typeId = 220, baseId = 1, txt = 'Token',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {226, 232, 236, 368, 372, 376, 512, 516, 520, 590, 594, 598}, }
_typeInfoList[37] = { parentId = 108, typeId = 221, nilable = true, orgTypeId = 220 }
_typeInfoList[38] = { parentId = 1, typeId = 224, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[39] = { parentId = 1, typeId = 225, nilable = true, orgTypeId = 224 }
_typeInfoList[40] = { parentId = 220, typeId = 226, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 216, 225}, retTypeId = {}, children = {}, }
_typeInfoList[41] = { parentId = 1, typeId = 230, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[42] = { parentId = 1, typeId = 231, nilable = true, orgTypeId = 230 }
_typeInfoList[43] = { parentId = 220, typeId = 232, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {230}, retTypeId = {}, children = {}, }
_typeInfoList[44] = { parentId = 1, typeId = 234, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[45] = { parentId = 1, typeId = 235, nilable = true, orgTypeId = 234 }
_typeInfoList[46] = { parentId = 220, typeId = 236, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {234}, children = {}, }
_typeInfoList[47] = { parentId = 108, typeId = 238, baseId = 1, txt = 'Parser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {240, 242, 244, 378, 380, 382, 522, 524, 526, 600, 602, 604}, }
_typeInfoList[48] = { parentId = 108, typeId = 239, nilable = true, orgTypeId = 238 }
_typeInfoList[49] = { parentId = 238, typeId = 240, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[50] = { parentId = 238, typeId = 242, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[51] = { parentId = 238, typeId = 244, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[52] = { parentId = 108, typeId = 246, baseId = 238, txt = 'WrapParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {248, 250, 252, 384, 386, 388, 528, 530, 532, 606, 608, 610}, }
_typeInfoList[53] = { parentId = 108, typeId = 247, nilable = true, orgTypeId = 246 }
_typeInfoList[54] = { parentId = 246, typeId = 248, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[55] = { parentId = 246, typeId = 250, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[56] = { parentId = 246, typeId = 252, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {238, 18}, retTypeId = {}, children = {}, }
_typeInfoList[57] = { parentId = 108, typeId = 254, baseId = 238, txt = 'StreamParser',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {268, 272, 274, 336, 390, 392, 394, 408, 420, 534, 536, 538, 550, 562, 612, 614, 616, 628, 640}, }
_typeInfoList[58] = { parentId = 108, typeId = 255, nilable = true, orgTypeId = 254 }
_typeInfoList[59] = { parentId = 254, typeId = 268, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {204, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[60] = { parentId = 254, typeId = 272, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[61] = { parentId = 254, typeId = 274, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {255}, children = {}, }
_typeInfoList[62] = { parentId = 108, typeId = 294, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[63] = { parentId = 108, typeId = 310, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[64] = { parentId = 108, typeId = 312, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[65] = { parentId = 108, typeId = 314, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[66] = { parentId = 1, typeId = 316, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[67] = { parentId = 1, typeId = 317, nilable = true, orgTypeId = 316 }
_typeInfoList[68] = { parentId = 254, typeId = 318, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {317}, children = {}, }
_typeInfoList[69] = { parentId = 318, typeId = 320, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[70] = { parentId = 318, typeId = 326, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[71] = { parentId = 318, typeId = 328, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {}, }
_typeInfoList[72] = { parentId = 328, typeId = 330, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {220}, children = {}, }
_typeInfoList[73] = { parentId = 328, typeId = 334, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[74] = { parentId = 254, typeId = 336, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[75] = { parentId = 108, typeId = 342, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[76] = { parentId = 1, typeId = 344, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[77] = { parentId = 1, typeId = 345, nilable = true, orgTypeId = 344 }
_typeInfoList[78] = { parentId = 1, typeId = 346, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[79] = { parentId = 1, typeId = 347, nilable = true, orgTypeId = 346 }
_typeInfoList[80] = { parentId = 1, typeId = 348, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[81] = { parentId = 1, typeId = 349, nilable = true, orgTypeId = 348 }
_typeInfoList[82] = { parentId = 1, typeId = 350, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[83] = { parentId = 1, typeId = 351, nilable = true, orgTypeId = 350 }
_typeInfoList[84] = { parentId = 1, typeId = 352, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 350}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[85] = { parentId = 1, typeId = 353, nilable = true, orgTypeId = 352 }
_typeInfoList[86] = { parentId = 108, typeId = 354, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {344, 346, 348, 352}, children = {}, }
_typeInfoList[87] = { parentId = 204, typeId = 356, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[88] = { parentId = 204, typeId = 358, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[89] = { parentId = 210, typeId = 360, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[90] = { parentId = 210, typeId = 362, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[91] = { parentId = 216, typeId = 364, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[92] = { parentId = 1, typeId = 366, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[93] = { parentId = 1, typeId = 367, nilable = true, orgTypeId = 366 }
_typeInfoList[94] = { parentId = 220, typeId = 368, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 216, 367}, retTypeId = {}, children = {}, }
_typeInfoList[95] = { parentId = 1, typeId = 370, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[96] = { parentId = 1, typeId = 371, nilable = true, orgTypeId = 370 }
_typeInfoList[97] = { parentId = 220, typeId = 372, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {370}, retTypeId = {}, children = {}, }
_typeInfoList[98] = { parentId = 1, typeId = 374, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[99] = { parentId = 1, typeId = 375, nilable = true, orgTypeId = 374 }
_typeInfoList[100] = { parentId = 220, typeId = 376, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {374}, children = {}, }
_typeInfoList[101] = { parentId = 238, typeId = 378, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[102] = { parentId = 238, typeId = 380, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[103] = { parentId = 238, typeId = 382, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[104] = { parentId = 246, typeId = 384, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[105] = { parentId = 246, typeId = 386, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[106] = { parentId = 246, typeId = 388, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {238, 18}, retTypeId = {}, children = {}, }
_typeInfoList[107] = { parentId = 254, typeId = 390, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {204, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[108] = { parentId = 254, typeId = 392, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[109] = { parentId = 254, typeId = 394, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {255}, children = {}, }
_typeInfoList[110] = { parentId = 108, typeId = 398, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[111] = { parentId = 108, typeId = 400, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[112] = { parentId = 108, typeId = 402, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[113] = { parentId = 108, typeId = 404, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[114] = { parentId = 1, typeId = 406, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[115] = { parentId = 1, typeId = 407, nilable = true, orgTypeId = 406 }
_typeInfoList[116] = { parentId = 254, typeId = 408, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {407}, children = {410, 412, 414}, }
_typeInfoList[117] = { parentId = 408, typeId = 410, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[118] = { parentId = 408, typeId = 412, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[119] = { parentId = 408, typeId = 414, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {416, 418}, }
_typeInfoList[120] = { parentId = 414, typeId = 416, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {220}, children = {}, }
_typeInfoList[121] = { parentId = 414, typeId = 418, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[122] = { parentId = 254, typeId = 420, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[123] = { parentId = 108, typeId = 422, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[124] = { parentId = 104, typeId = 424, baseId = 1, txt = 'Util',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {466, 472, 480, 482, 486, 654, 656, 658}, }
_typeInfoList[125] = { parentId = 104, typeId = 425, nilable = true, orgTypeId = 424 }
_typeInfoList[126] = { parentId = 424, typeId = 466, baseId = 1, txt = 'outStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {468, 470, 644, 646}, }
_typeInfoList[127] = { parentId = 424, typeId = 467, nilable = true, orgTypeId = 466 }
_typeInfoList[128] = { parentId = 466, typeId = 468, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[129] = { parentId = 466, typeId = 470, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[130] = { parentId = 424, typeId = 472, baseId = 466, txt = 'memStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {474, 476, 478, 648, 650, 652}, }
_typeInfoList[131] = { parentId = 424, typeId = 473, nilable = true, orgTypeId = 472 }
_typeInfoList[132] = { parentId = 472, typeId = 474, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[133] = { parentId = 472, typeId = 476, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[134] = { parentId = 472, typeId = 478, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[135] = { parentId = 424, typeId = 480, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[136] = { parentId = 424, typeId = 482, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[137] = { parentId = 424, typeId = 486, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10, 6, 18}, retTypeId = {}, children = {}, }
_typeInfoList[138] = { parentId = 1, typeId = 488, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[139] = { parentId = 1, typeId = 489, nilable = true, orgTypeId = 488 }
_typeInfoList[140] = { parentId = 1, typeId = 490, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[141] = { parentId = 1, typeId = 491, nilable = true, orgTypeId = 490 }
_typeInfoList[142] = { parentId = 1, typeId = 492, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[143] = { parentId = 1, typeId = 493, nilable = true, orgTypeId = 492 }
_typeInfoList[144] = { parentId = 1, typeId = 494, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[145] = { parentId = 1, typeId = 495, nilable = true, orgTypeId = 494 }
_typeInfoList[146] = { parentId = 1, typeId = 496, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 494}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[147] = { parentId = 1, typeId = 497, nilable = true, orgTypeId = 496 }
_typeInfoList[148] = { parentId = 108, typeId = 498, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {488, 490, 492, 496}, children = {}, }
_typeInfoList[149] = { parentId = 204, typeId = 500, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[150] = { parentId = 204, typeId = 502, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[151] = { parentId = 210, typeId = 504, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[152] = { parentId = 210, typeId = 506, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[153] = { parentId = 216, typeId = 508, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[154] = { parentId = 1, typeId = 510, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[155] = { parentId = 1, typeId = 511, nilable = true, orgTypeId = 510 }
_typeInfoList[156] = { parentId = 220, typeId = 512, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 216, 511}, retTypeId = {}, children = {}, }
_typeInfoList[157] = { parentId = 1, typeId = 514, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[158] = { parentId = 1, typeId = 515, nilable = true, orgTypeId = 514 }
_typeInfoList[159] = { parentId = 220, typeId = 516, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {514}, retTypeId = {}, children = {}, }
_typeInfoList[160] = { parentId = 1, typeId = 518, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[161] = { parentId = 1, typeId = 519, nilable = true, orgTypeId = 518 }
_typeInfoList[162] = { parentId = 220, typeId = 520, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {518}, children = {}, }
_typeInfoList[163] = { parentId = 238, typeId = 522, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[164] = { parentId = 238, typeId = 524, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[165] = { parentId = 238, typeId = 526, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[166] = { parentId = 246, typeId = 528, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[167] = { parentId = 246, typeId = 530, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[168] = { parentId = 246, typeId = 532, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {238, 18}, retTypeId = {}, children = {}, }
_typeInfoList[169] = { parentId = 254, typeId = 534, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {204, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[170] = { parentId = 254, typeId = 536, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[171] = { parentId = 254, typeId = 538, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {255}, children = {}, }
_typeInfoList[172] = { parentId = 108, typeId = 540, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[173] = { parentId = 108, typeId = 542, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[174] = { parentId = 108, typeId = 544, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[175] = { parentId = 108, typeId = 546, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[176] = { parentId = 1, typeId = 548, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[177] = { parentId = 1, typeId = 549, nilable = true, orgTypeId = 548 }
_typeInfoList[178] = { parentId = 254, typeId = 550, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {549}, children = {552, 554, 556}, }
_typeInfoList[179] = { parentId = 550, typeId = 552, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[180] = { parentId = 550, typeId = 554, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[181] = { parentId = 550, typeId = 556, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {558, 560}, }
_typeInfoList[182] = { parentId = 556, typeId = 558, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {220}, children = {}, }
_typeInfoList[183] = { parentId = 556, typeId = 560, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[184] = { parentId = 254, typeId = 562, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[185] = { parentId = 108, typeId = 564, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[186] = { parentId = 1, typeId = 566, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[187] = { parentId = 1, typeId = 567, nilable = true, orgTypeId = 566 }
_typeInfoList[188] = { parentId = 1, typeId = 568, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[189] = { parentId = 1, typeId = 569, nilable = true, orgTypeId = 568 }
_typeInfoList[190] = { parentId = 1, typeId = 570, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[191] = { parentId = 1, typeId = 571, nilable = true, orgTypeId = 570 }
_typeInfoList[192] = { parentId = 1, typeId = 572, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[193] = { parentId = 1, typeId = 573, nilable = true, orgTypeId = 572 }
_typeInfoList[194] = { parentId = 1, typeId = 574, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 572}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[195] = { parentId = 1, typeId = 575, nilable = true, orgTypeId = 574 }
_typeInfoList[196] = { parentId = 108, typeId = 576, baseId = 1, txt = 'createReserveInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10}, retTypeId = {566, 568, 570, 574}, children = {}, }
_typeInfoList[197] = { parentId = 204, typeId = 578, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[198] = { parentId = 204, typeId = 580, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[199] = { parentId = 210, typeId = 582, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[200] = { parentId = 210, typeId = 584, baseId = 1, txt = 'read',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {19}, children = {}, }
_typeInfoList[201] = { parentId = 216, typeId = 586, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 12}, retTypeId = {}, children = {}, }
_typeInfoList[202] = { parentId = 1, typeId = 588, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[203] = { parentId = 1, typeId = 589, nilable = true, orgTypeId = 588 }
_typeInfoList[204] = { parentId = 220, typeId = 590, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 18, 216, 589}, retTypeId = {}, children = {}, }
_typeInfoList[205] = { parentId = 1, typeId = 592, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[206] = { parentId = 1, typeId = 593, nilable = true, orgTypeId = 592 }
_typeInfoList[207] = { parentId = 220, typeId = 594, baseId = 1, txt = 'set_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {592}, retTypeId = {}, children = {}, }
_typeInfoList[208] = { parentId = 1, typeId = 596, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[209] = { parentId = 1, typeId = 597, nilable = true, orgTypeId = 596 }
_typeInfoList[210] = { parentId = 220, typeId = 598, baseId = 1, txt = 'get_commentList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {596}, children = {}, }
_typeInfoList[211] = { parentId = 238, typeId = 600, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[212] = { parentId = 238, typeId = 602, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[213] = { parentId = 238, typeId = 604, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[214] = { parentId = 246, typeId = 606, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[215] = { parentId = 246, typeId = 608, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[216] = { parentId = 246, typeId = 610, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {238, 18}, retTypeId = {}, children = {}, }
_typeInfoList[217] = { parentId = 254, typeId = 612, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {204, 18, 10}, retTypeId = {}, children = {}, }
_typeInfoList[218] = { parentId = 254, typeId = 614, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[219] = { parentId = 254, typeId = 616, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 10}, retTypeId = {255}, children = {}, }
_typeInfoList[220] = { parentId = 108, typeId = 618, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[221] = { parentId = 108, typeId = 620, baseId = 1, txt = 'getKindTxt',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[222] = { parentId = 108, typeId = 622, baseId = 1, txt = 'isOp2',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[223] = { parentId = 108, typeId = 624, baseId = 1, txt = 'isOp1',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {6}, children = {}, }
_typeInfoList[224] = { parentId = 1, typeId = 626, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[225] = { parentId = 1, typeId = 627, nilable = true, orgTypeId = 626 }
_typeInfoList[226] = { parentId = 254, typeId = 628, baseId = 1, txt = 'parse',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {627}, children = {630, 632, 634}, }
_typeInfoList[227] = { parentId = 628, typeId = 630, baseId = 1, txt = 'readLine',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {19}, children = {}, }
_typeInfoList[228] = { parentId = 628, typeId = 632, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18}, retTypeId = {18, 12}, children = {}, }
_typeInfoList[229] = { parentId = 628, typeId = 634, baseId = 1, txt = '',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {6}, children = {636, 638}, }
_typeInfoList[230] = { parentId = 634, typeId = 636, baseId = 1, txt = 'createInfo',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12, 18, 12}, retTypeId = {220}, children = {}, }
_typeInfoList[231] = { parentId = 634, typeId = 638, baseId = 1, txt = 'analyzeNumber',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 12}, retTypeId = {12, 10}, children = {}, }
_typeInfoList[232] = { parentId = 254, typeId = 640, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[233] = { parentId = 108, typeId = 642, baseId = 1, txt = 'getEofToken',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[234] = { parentId = 466, typeId = 644, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[235] = { parentId = 466, typeId = 646, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[236] = { parentId = 472, typeId = 648, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[237] = { parentId = 472, typeId = 650, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[238] = { parentId = 472, typeId = 652, baseId = 1, txt = 'get_txt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[239] = { parentId = 424, typeId = 654, baseId = 1, txt = 'errorLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[240] = { parentId = 424, typeId = 656, baseId = 1, txt = 'debugLog',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[241] = { parentId = 424, typeId = 658, baseId = 1, txt = 'profile',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {10, 6, 18}, retTypeId = {}, children = {}, }
_typeInfoList[242] = { parentId = 106, typeId = 660, baseId = 1, txt = 'TypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {728, 730, 732, 734, 736, 738, 740, 742, 744, 748, 752, 756, 758, 760, 762, 764, 766, 768, 770, 772, 774, 776, 778, 782, 786, 788}, }
_typeInfoList[243] = { parentId = 106, typeId = 661, nilable = true, orgTypeId = 660 }
_typeInfoList[244] = { parentId = 1, typeId = 662, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[245] = { parentId = 1, typeId = 663, nilable = true, orgTypeId = 662 }
_typeInfoList[246] = { parentId = 106, typeId = 674, baseId = 1, txt = 'isBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {6}, children = {}, }
_typeInfoList[247] = { parentId = 106, typeId = 676, baseId = 1, txt = 'OutStream',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {678, 680}, }
_typeInfoList[248] = { parentId = 106, typeId = 677, nilable = true, orgTypeId = 676 }
_typeInfoList[249] = { parentId = 676, typeId = 678, baseId = 1, txt = 'write',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[250] = { parentId = 676, typeId = 680, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[251] = { parentId = 106, typeId = 686, baseId = 1, txt = 'Scope',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {710, 712, 714, 716, 718, 722, 790, 792, 794, 796, 808, 810}, }
_typeInfoList[252] = { parentId = 106, typeId = 687, nilable = true, orgTypeId = 686 }
_typeInfoList[253] = { parentId = 106, typeId = 688, baseId = 1, txt = 'SymbolInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {690, 692, 694, 696, 698, 812}, }
_typeInfoList[254] = { parentId = 106, typeId = 689, nilable = true, orgTypeId = 688 }
_typeInfoList[255] = { parentId = 688, typeId = 690, baseId = 1, txt = 'canAccess',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {686, 686}, retTypeId = {661}, children = {}, }
_typeInfoList[256] = { parentId = 688, typeId = 692, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[257] = { parentId = 688, typeId = 694, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[258] = { parentId = 688, typeId = 696, baseId = 1, txt = 'get_typeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[259] = { parentId = 688, typeId = 698, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 18, 660}, retTypeId = {}, children = {}, }
_typeInfoList[260] = { parentId = 1, typeId = 704, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {686}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[261] = { parentId = 1, typeId = 705, nilable = true, orgTypeId = 704 }
_typeInfoList[262] = { parentId = 686, typeId = 706, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {687, 10, 704}, retTypeId = {}, children = {}, }
_typeInfoList[263] = { parentId = 686, typeId = 710, baseId = 1, txt = 'set_ownerTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {661}, retTypeId = {}, children = {}, }
_typeInfoList[264] = { parentId = 686, typeId = 712, baseId = 1, txt = 'getTypeInfoChild',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {661}, children = {}, }
_typeInfoList[265] = { parentId = 686, typeId = 714, baseId = 1, txt = 'getNSTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[266] = { parentId = 686, typeId = 716, baseId = 1, txt = 'get_ownerTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {661}, children = {}, }
_typeInfoList[267] = { parentId = 686, typeId = 718, baseId = 1, txt = 'get_parent',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {686}, children = {}, }
_typeInfoList[268] = { parentId = 1, typeId = 720, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 688}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[269] = { parentId = 1, typeId = 721, nilable = true, orgTypeId = 720 }
_typeInfoList[270] = { parentId = 686, typeId = 722, baseId = 1, txt = 'get_symbol2TypeInfoMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {720}, children = {}, }
_typeInfoList[271] = { parentId = 660, typeId = 728, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {687}, retTypeId = {}, children = {}, }
_typeInfoList[272] = { parentId = 660, typeId = 730, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[273] = { parentId = 660, typeId = 732, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[274] = { parentId = 660, typeId = 734, baseId = 1, txt = 'isInheritFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[275] = { parentId = 660, typeId = 736, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[276] = { parentId = 660, typeId = 738, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[277] = { parentId = 660, typeId = 740, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {676}, retTypeId = {}, children = {}, }
_typeInfoList[278] = { parentId = 660, typeId = 742, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[279] = { parentId = 660, typeId = 744, baseId = 1, txt = 'get_externalFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[280] = { parentId = 1, typeId = 746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[281] = { parentId = 1, typeId = 747, nilable = true, orgTypeId = 746 }
_typeInfoList[282] = { parentId = 660, typeId = 748, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {746}, children = {}, }
_typeInfoList[283] = { parentId = 1, typeId = 750, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[284] = { parentId = 1, typeId = 751, nilable = true, orgTypeId = 750 }
_typeInfoList[285] = { parentId = 660, typeId = 752, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {750}, children = {}, }
_typeInfoList[286] = { parentId = 1, typeId = 754, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[287] = { parentId = 1, typeId = 755, nilable = true, orgTypeId = 754 }
_typeInfoList[288] = { parentId = 660, typeId = 756, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {754}, children = {}, }
_typeInfoList[289] = { parentId = 660, typeId = 758, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[290] = { parentId = 660, typeId = 760, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[291] = { parentId = 660, typeId = 762, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[292] = { parentId = 660, typeId = 764, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[293] = { parentId = 660, typeId = 766, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[294] = { parentId = 660, typeId = 768, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[295] = { parentId = 660, typeId = 770, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[296] = { parentId = 660, typeId = 772, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[297] = { parentId = 660, typeId = 774, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[298] = { parentId = 660, typeId = 776, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[299] = { parentId = 660, typeId = 778, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[300] = { parentId = 1, typeId = 780, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[301] = { parentId = 1, typeId = 781, nilable = true, orgTypeId = 780 }
_typeInfoList[302] = { parentId = 660, typeId = 782, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {780}, children = {}, }
_typeInfoList[303] = { parentId = 1, typeId = 784, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[304] = { parentId = 1, typeId = 785, nilable = true, orgTypeId = 784 }
_typeInfoList[305] = { parentId = 660, typeId = 786, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {784}, children = {}, }
_typeInfoList[306] = { parentId = 660, typeId = 788, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {687}, children = {}, }
_typeInfoList[307] = { parentId = 686, typeId = 790, baseId = 1, txt = 'getTypeInfoField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 11, 686}, retTypeId = {661}, children = {}, }
_typeInfoList[308] = { parentId = 686, typeId = 792, baseId = 1, txt = 'getTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 686}, retTypeId = {661}, children = {}, }
_typeInfoList[309] = { parentId = 686, typeId = 794, baseId = 1, txt = 'add',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 660, 18}, retTypeId = {}, children = {}, }
_typeInfoList[310] = { parentId = 686, typeId = 796, baseId = 1, txt = 'addClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 660, 686}, retTypeId = {}, children = {}, }
_typeInfoList[311] = { parentId = 1, typeId = 798, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'local', kind = 5, itemTypeId = {686, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[312] = { parentId = 1, typeId = 799, nilable = true, orgTypeId = 798 }
_typeInfoList[313] = { parentId = 106, typeId = 800, baseId = 1, txt = 'dumpScopeSub',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {687, 18, 798}, retTypeId = {}, children = {}, }
_typeInfoList[314] = { parentId = 106, typeId = 802, baseId = 1, txt = 'dumpScope',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {687, 18}, retTypeId = {}, children = {}, }
_typeInfoList[315] = { parentId = 686, typeId = 808, baseId = 1, txt = 'getNSTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[316] = { parentId = 686, typeId = 810, baseId = 1, txt = 'getClassTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[317] = { parentId = 688, typeId = 812, baseId = 1, txt = 'canAccess',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {686, 686}, retTypeId = {661}, children = {}, }
_typeInfoList[318] = { parentId = 106, typeId = 814, baseId = 1, txt = 'Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {964, 970, 976, 978, 980, 984, 986}, }
_typeInfoList[319] = { parentId = 106, typeId = 815, nilable = true, orgTypeId = 814 }
_typeInfoList[320] = { parentId = 106, typeId = 816, baseId = 814, txt = 'DeclClassNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2368, 2382, 2384, 2386, 2390, 2394, 2396, 2400}, }
_typeInfoList[321] = { parentId = 106, typeId = 817, nilable = true, orgTypeId = 816 }
_typeInfoList[322] = { parentId = 106, typeId = 818, baseId = 660, txt = 'NormalTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {844, 846, 848, 850, 856, 858, 860, 868, 872, 876, 880, 882, 884, 886, 888, 890, 892, 894, 896, 898, 900, 902, 906, 908, 928, 932, 934, 938, 944, 954, 956}, }
_typeInfoList[323] = { parentId = 106, typeId = 819, nilable = true, orgTypeId = 818 }
_typeInfoList[324] = { parentId = 1, typeId = 828, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[325] = { parentId = 1, typeId = 829, nilable = true, orgTypeId = 828 }
_typeInfoList[326] = { parentId = 1, typeId = 830, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[327] = { parentId = 1, typeId = 831, nilable = true, orgTypeId = 830 }
_typeInfoList[328] = { parentId = 1, typeId = 832, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[329] = { parentId = 1, typeId = 833, nilable = true, orgTypeId = 832 }
_typeInfoList[330] = { parentId = 818, typeId = 834, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {687, 661, 661, 10, 10, 10, 18, 18, 661, 12, 12, 829, 831, 833}, retTypeId = {}, children = {}, }
_typeInfoList[331] = { parentId = 818, typeId = 844, baseId = 1, txt = 'getParentId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[332] = { parentId = 818, typeId = 846, baseId = 1, txt = 'get_baseId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[333] = { parentId = 818, typeId = 848, baseId = 1, txt = 'getTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[334] = { parentId = 818, typeId = 850, baseId = 1, txt = 'serialize',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {676}, retTypeId = {}, children = {}, }
_typeInfoList[335] = { parentId = 1, typeId = 852, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[336] = { parentId = 1, typeId = 853, nilable = true, orgTypeId = 852 }
_typeInfoList[337] = { parentId = 850, typeId = 854, baseId = 1, txt = 'serializeTypeInfoList',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {18, 852, 11}, retTypeId = {18}, children = {}, }
_typeInfoList[338] = { parentId = 818, typeId = 856, baseId = 1, txt = 'equalsSub',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660, 12}, retTypeId = {10}, children = {}, }
_typeInfoList[339] = { parentId = 818, typeId = 858, baseId = 1, txt = 'equals',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[340] = { parentId = 818, typeId = 860, baseId = 1, txt = 'cloneToPublic',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {660}, retTypeId = {818}, children = {}, }
_typeInfoList[341] = { parentId = 1, typeId = 862, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[342] = { parentId = 1, typeId = 863, nilable = true, orgTypeId = 862 }
_typeInfoList[343] = { parentId = 1, typeId = 864, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[344] = { parentId = 1, typeId = 865, nilable = true, orgTypeId = 864 }
_typeInfoList[345] = { parentId = 1, typeId = 866, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[346] = { parentId = 1, typeId = 867, nilable = true, orgTypeId = 866 }
_typeInfoList[347] = { parentId = 818, typeId = 868, baseId = 1, txt = 'create',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {687, 660, 660, 10, 12, 18, 862, 864, 866}, retTypeId = {660}, children = {}, }
_typeInfoList[348] = { parentId = 1, typeId = 870, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[349] = { parentId = 1, typeId = 871, nilable = true, orgTypeId = 870 }
_typeInfoList[350] = { parentId = 818, typeId = 872, baseId = 1, txt = 'get_itemTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {870}, children = {}, }
_typeInfoList[351] = { parentId = 1, typeId = 874, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[352] = { parentId = 1, typeId = 875, nilable = true, orgTypeId = 874 }
_typeInfoList[353] = { parentId = 818, typeId = 876, baseId = 1, txt = 'get_argTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {874}, children = {}, }
_typeInfoList[354] = { parentId = 1, typeId = 878, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[355] = { parentId = 1, typeId = 879, nilable = true, orgTypeId = 878 }
_typeInfoList[356] = { parentId = 818, typeId = 880, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {878}, children = {}, }
_typeInfoList[357] = { parentId = 818, typeId = 882, baseId = 1, txt = 'get_parentInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[358] = { parentId = 818, typeId = 884, baseId = 1, txt = 'get_typeId',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[359] = { parentId = 818, typeId = 886, baseId = 1, txt = 'get_rawTxt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[360] = { parentId = 818, typeId = 888, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[361] = { parentId = 818, typeId = 890, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[362] = { parentId = 818, typeId = 892, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[363] = { parentId = 818, typeId = 894, baseId = 1, txt = 'get_autoFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[364] = { parentId = 818, typeId = 896, baseId = 1, txt = 'get_orgTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[365] = { parentId = 818, typeId = 898, baseId = 1, txt = 'get_baseTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[366] = { parentId = 818, typeId = 900, baseId = 1, txt = 'get_nilable',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[367] = { parentId = 818, typeId = 902, baseId = 1, txt = 'get_nilableTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[368] = { parentId = 1, typeId = 904, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[369] = { parentId = 1, typeId = 905, nilable = true, orgTypeId = 904 }
_typeInfoList[370] = { parentId = 818, typeId = 906, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {904}, children = {}, }
_typeInfoList[371] = { parentId = 818, typeId = 908, baseId = 1, txt = 'createBuiltin',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 18, 12, 661}, retTypeId = {660}, children = {}, }
_typeInfoList[372] = { parentId = 1, typeId = 926, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[373] = { parentId = 1, typeId = 927, nilable = true, orgTypeId = 926 }
_typeInfoList[374] = { parentId = 818, typeId = 928, baseId = 1, txt = 'createList',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 660, 926}, retTypeId = {660}, children = {}, }
_typeInfoList[375] = { parentId = 1, typeId = 930, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[376] = { parentId = 1, typeId = 931, nilable = true, orgTypeId = 930 }
_typeInfoList[377] = { parentId = 818, typeId = 932, baseId = 1, txt = 'createArray',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 660, 930}, retTypeId = {660}, children = {}, }
_typeInfoList[378] = { parentId = 818, typeId = 934, baseId = 1, txt = 'createMap',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {18, 660, 660, 660}, retTypeId = {660}, children = {}, }
_typeInfoList[379] = { parentId = 818, typeId = 938, baseId = 1, txt = 'createClass',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {687, 661, 660, 10, 18, 18}, retTypeId = {660}, children = {}, }
_typeInfoList[380] = { parentId = 1, typeId = 940, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[381] = { parentId = 1, typeId = 941, nilable = true, orgTypeId = 940 }
_typeInfoList[382] = { parentId = 1, typeId = 942, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[383] = { parentId = 1, typeId = 943, nilable = true, orgTypeId = 942 }
_typeInfoList[384] = { parentId = 818, typeId = 944, baseId = 1, txt = 'createFunc',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {687, 12, 660, 10, 10, 10, 18, 18, 941, 943}, retTypeId = {660}, children = {}, }
_typeInfoList[385] = { parentId = 818, typeId = 954, baseId = 1, txt = 'isInheritFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[386] = { parentId = 818, typeId = 956, baseId = 1, txt = 'isSettableFrom',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {660}, retTypeId = {10}, children = {}, }
_typeInfoList[387] = { parentId = 106, typeId = 958, baseId = 1, txt = 'Filter',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {960, 1202, 1220, 1242, 1282, 1310, 1348, 1372, 1410, 1442, 1466, 1496, 1528, 1564, 1598, 1626, 1644, 1664, 1688, 1710, 1734, 1762, 1790, 1814, 1838, 1864, 1888, 1910, 1930, 1950, 1976, 2002, 2024, 2050, 2104, 2188, 2208, 2228, 2250, 2282, 2314, 2334, 2360, 2408, 2426, 2446, 2470, 2494, 2516, 2536, 2566, 2602, 2630, 2650}, }
_typeInfoList[388] = { parentId = 106, typeId = 959, nilable = true, orgTypeId = 958 }
_typeInfoList[389] = { parentId = 958, typeId = 960, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[390] = { parentId = 1, typeId = 962, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[391] = { parentId = 1, typeId = 963, nilable = true, orgTypeId = 962 }
_typeInfoList[392] = { parentId = 814, typeId = 964, baseId = 1, txt = 'get_expType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[393] = { parentId = 1, typeId = 966, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[394] = { parentId = 1, typeId = 967, nilable = true, orgTypeId = 966 }
_typeInfoList[395] = { parentId = 1, typeId = 968, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[396] = { parentId = 1, typeId = 969, nilable = true, orgTypeId = 968 }
_typeInfoList[397] = { parentId = 814, typeId = 970, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {966, 968}, children = {}, }
_typeInfoList[398] = { parentId = 814, typeId = 976, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[399] = { parentId = 814, typeId = 978, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[400] = { parentId = 814, typeId = 980, baseId = 1, txt = 'get_pos',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {216}, children = {}, }
_typeInfoList[401] = { parentId = 1, typeId = 982, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[402] = { parentId = 1, typeId = 983, nilable = true, orgTypeId = 982 }
_typeInfoList[403] = { parentId = 814, typeId = 984, baseId = 1, txt = 'get_expTypeList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {982}, children = {}, }
_typeInfoList[404] = { parentId = 814, typeId = 986, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 216, 962}, retTypeId = {}, children = {}, }
_typeInfoList[405] = { parentId = 106, typeId = 988, baseId = 1, txt = 'NamespaceInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {990}, }
_typeInfoList[406] = { parentId = 106, typeId = 989, nilable = true, orgTypeId = 988 }
_typeInfoList[407] = { parentId = 988, typeId = 990, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 686, 660}, retTypeId = {}, children = {}, }
_typeInfoList[408] = { parentId = 106, typeId = 992, baseId = 1, txt = 'MacroEval',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2808, 2810}, }
_typeInfoList[409] = { parentId = 106, typeId = 993, nilable = true, orgTypeId = 992 }
_typeInfoList[410] = { parentId = 106, typeId = 994, baseId = 814, txt = 'ExpListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1376, 1386, 1390}, }
_typeInfoList[411] = { parentId = 106, typeId = 995, nilable = true, orgTypeId = 994 }
_typeInfoList[412] = { parentId = 106, typeId = 996, baseId = 1, txt = 'DeclMacroInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1002, 1006, 1008, 1012, 1014}, }
_typeInfoList[413] = { parentId = 106, typeId = 997, nilable = true, orgTypeId = 996 }
_typeInfoList[414] = { parentId = 1, typeId = 998, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[415] = { parentId = 1, typeId = 999, nilable = true, orgTypeId = 998 }
_typeInfoList[416] = { parentId = 1, typeId = 1000, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[417] = { parentId = 1, typeId = 1001, nilable = true, orgTypeId = 1000 }
_typeInfoList[418] = { parentId = 996, typeId = 1002, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[419] = { parentId = 1, typeId = 1004, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[420] = { parentId = 1, typeId = 1005, nilable = true, orgTypeId = 1004 }
_typeInfoList[421] = { parentId = 996, typeId = 1006, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1004}, children = {}, }
_typeInfoList[422] = { parentId = 996, typeId = 1008, baseId = 1, txt = 'get_ast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {815}, children = {}, }
_typeInfoList[423] = { parentId = 1, typeId = 1010, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[424] = { parentId = 1, typeId = 1011, nilable = true, orgTypeId = 1010 }
_typeInfoList[425] = { parentId = 996, typeId = 1012, baseId = 1, txt = 'get_tokenList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1010}, children = {}, }
_typeInfoList[426] = { parentId = 996, typeId = 1014, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {220, 998, 815, 1000}, retTypeId = {}, children = {}, }
_typeInfoList[427] = { parentId = 106, typeId = 1016, baseId = 1, txt = 'MacroValInfo',
        staticFlag = false, accessMode = 'local', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1018}, }
_typeInfoList[428] = { parentId = 106, typeId = 1017, nilable = true, orgTypeId = 1016 }
_typeInfoList[429] = { parentId = 1016, typeId = 1018, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {6, 660}, retTypeId = {}, children = {}, }
_typeInfoList[430] = { parentId = 106, typeId = 1020, baseId = 1, txt = 'MacroInfo',
        staticFlag = false, accessMode = 'local', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1024}, }
_typeInfoList[431] = { parentId = 106, typeId = 1021, nilable = true, orgTypeId = 1020 }
_typeInfoList[432] = { parentId = 1, typeId = 1022, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 1016}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[433] = { parentId = 1, typeId = 1023, nilable = true, orgTypeId = 1022 }
_typeInfoList[434] = { parentId = 1020, typeId = 1024, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {26, 996, 1022}, retTypeId = {}, children = {}, }
_typeInfoList[435] = { parentId = 106, typeId = 1026, baseId = 1, txt = 'TransUnit',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1118, 3262}, }
_typeInfoList[436] = { parentId = 106, typeId = 1027, nilable = true, orgTypeId = 1026 }
_typeInfoList[437] = { parentId = 1026, typeId = 1044, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {992}, retTypeId = {}, children = {}, }
_typeInfoList[438] = { parentId = 1026, typeId = 1060, baseId = 1, txt = 'addErrMess',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {216, 18}, retTypeId = {}, children = {}, }
_typeInfoList[439] = { parentId = 1, typeId = 1062, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {686}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[440] = { parentId = 1, typeId = 1063, nilable = true, orgTypeId = 1062 }
_typeInfoList[441] = { parentId = 1026, typeId = 1064, baseId = 1, txt = 'pushScope',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 1063}, retTypeId = {686}, children = {}, }
_typeInfoList[442] = { parentId = 1026, typeId = 1068, baseId = 1, txt = 'popScope',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[443] = { parentId = 1026, typeId = 1070, baseId = 1, txt = 'getCurrentClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[444] = { parentId = 1026, typeId = 1072, baseId = 1, txt = 'getCurrentNamespaceTypeInfo',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[445] = { parentId = 1026, typeId = 1074, baseId = 1, txt = 'pushClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {661, 10, 18, 18, 989}, retTypeId = {660}, children = {}, }
_typeInfoList[446] = { parentId = 1026, typeId = 1082, baseId = 1, txt = 'popClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[447] = { parentId = 1026, typeId = 1084, baseId = 1, txt = 'pushbackStr',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[448] = { parentId = 1026, typeId = 1086, baseId = 1, txt = 'analyzeDecl',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 220, 220}, retTypeId = {815}, children = {}, }
_typeInfoList[449] = { parentId = 1026, typeId = 1088, baseId = 1, txt = 'analyzeDeclVar',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 18, 10, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[450] = { parentId = 1026, typeId = 1090, baseId = 1, txt = 'analyzeDeclFunc',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 221, 220, 221}, retTypeId = {814}, children = {}, }
_typeInfoList[451] = { parentId = 1026, typeId = 1092, baseId = 1, txt = 'analyzeDeclClass',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[452] = { parentId = 1026, typeId = 1094, baseId = 1, txt = 'analyzeExp',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10, 12}, retTypeId = {814}, children = {}, }
_typeInfoList[453] = { parentId = 1, typeId = 1096, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[454] = { parentId = 1, typeId = 1097, nilable = true, orgTypeId = 1096 }
_typeInfoList[455] = { parentId = 1026, typeId = 1098, baseId = 1, txt = 'analyzeStatementList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {1096, 19}, retTypeId = {}, children = {}, }
_typeInfoList[456] = { parentId = 1026, typeId = 1100, baseId = 1, txt = 'analyzeStatement',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {19}, retTypeId = {}, children = {}, }
_typeInfoList[457] = { parentId = 1026, typeId = 1102, baseId = 1, txt = 'analyzeExpSymbol',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {220, 220, 18, 815, 10}, retTypeId = {814}, children = {}, }
_typeInfoList[458] = { parentId = 1026, typeId = 1104, baseId = 1, txt = 'analyzeExpList',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {10}, retTypeId = {994}, children = {}, }
_typeInfoList[459] = { parentId = 1, typeId = 1116, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[460] = { parentId = 1, typeId = 1117, nilable = true, orgTypeId = 1116 }
_typeInfoList[461] = { parentId = 1026, typeId = 1118, baseId = 1, txt = 'get_errMessList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1116}, children = {}, }
_typeInfoList[462] = { parentId = 1, typeId = 1132, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 4, itemTypeId = {18}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[463] = { parentId = 1, typeId = 1133, nilable = true, orgTypeId = 1132 }
_typeInfoList[464] = { parentId = 106, typeId = 1134, baseId = 1, txt = 'regOpLevel',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {12, 1132}, retTypeId = {}, children = {}, }
_typeInfoList[465] = { parentId = 1, typeId = 1168, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[466] = { parentId = 1, typeId = 1169, nilable = true, orgTypeId = 1168 }
_typeInfoList[467] = { parentId = 106, typeId = 1172, baseId = 1, txt = 'regKind',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {18}, retTypeId = {12}, children = {}, }
_typeInfoList[468] = { parentId = 106, typeId = 1174, baseId = 1, txt = 'getNodeKindName',
        staticFlag = true, accessMode = 'pub', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {18}, children = {}, }
_typeInfoList[469] = { parentId = 106, typeId = 1200, baseId = 814, txt = 'NoneNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1204, 1212}, }
_typeInfoList[470] = { parentId = 106, typeId = 1201, nilable = true, orgTypeId = 1200 }
_typeInfoList[471] = { parentId = 958, typeId = 1202, baseId = 1, txt = 'processNone',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1200, 8}, retTypeId = {}, children = {}, }
_typeInfoList[472] = { parentId = 1200, typeId = 1204, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[473] = { parentId = 1, typeId = 1210, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[474] = { parentId = 1, typeId = 1211, nilable = true, orgTypeId = 1210 }
_typeInfoList[475] = { parentId = 1200, typeId = 1212, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1210}, retTypeId = {}, children = {}, }
_typeInfoList[476] = { parentId = 106, typeId = 1218, baseId = 814, txt = 'ImportNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1222, 1230, 1232}, }
_typeInfoList[477] = { parentId = 106, typeId = 1219, nilable = true, orgTypeId = 1218 }
_typeInfoList[478] = { parentId = 958, typeId = 1220, baseId = 1, txt = 'processImport',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1218, 8}, retTypeId = {}, children = {}, }
_typeInfoList[479] = { parentId = 1218, typeId = 1222, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[480] = { parentId = 1, typeId = 1228, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[481] = { parentId = 1, typeId = 1229, nilable = true, orgTypeId = 1228 }
_typeInfoList[482] = { parentId = 1218, typeId = 1230, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1228, 18}, retTypeId = {}, children = {}, }
_typeInfoList[483] = { parentId = 1218, typeId = 1232, baseId = 1, txt = 'get_modulePath',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[484] = { parentId = 106, typeId = 1240, baseId = 814, txt = 'RootNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1248, 1260, 1264, 1268}, }
_typeInfoList[485] = { parentId = 106, typeId = 1241, nilable = true, orgTypeId = 1240 }
_typeInfoList[486] = { parentId = 958, typeId = 1242, baseId = 1, txt = 'processRoot',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1240, 8}, retTypeId = {}, children = {}, }
_typeInfoList[487] = { parentId = 1240, typeId = 1248, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[488] = { parentId = 1, typeId = 1254, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[489] = { parentId = 1, typeId = 1255, nilable = true, orgTypeId = 1254 }
_typeInfoList[490] = { parentId = 1, typeId = 1256, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[491] = { parentId = 1, typeId = 1257, nilable = true, orgTypeId = 1256 }
_typeInfoList[492] = { parentId = 1, typeId = 1258, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 988}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[493] = { parentId = 1, typeId = 1259, nilable = true, orgTypeId = 1258 }
_typeInfoList[494] = { parentId = 1240, typeId = 1260, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1254, 1256, 1258}, retTypeId = {}, children = {}, }
_typeInfoList[495] = { parentId = 1, typeId = 1262, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[496] = { parentId = 1, typeId = 1263, nilable = true, orgTypeId = 1262 }
_typeInfoList[497] = { parentId = 1240, typeId = 1264, baseId = 1, txt = 'get_children',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1262}, children = {}, }
_typeInfoList[498] = { parentId = 1, typeId = 1266, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 988}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[499] = { parentId = 1, typeId = 1267, nilable = true, orgTypeId = 1266 }
_typeInfoList[500] = { parentId = 1240, typeId = 1268, baseId = 1, txt = 'get_typeId2ClassMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1266}, children = {}, }
_typeInfoList[501] = { parentId = 106, typeId = 1280, baseId = 814, txt = 'RefTypeNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1284, 1292, 1294, 1296, 1298, 1300}, }
_typeInfoList[502] = { parentId = 106, typeId = 1281, nilable = true, orgTypeId = 1280 }
_typeInfoList[503] = { parentId = 958, typeId = 1282, baseId = 1, txt = 'processRefType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1280, 8}, retTypeId = {}, children = {}, }
_typeInfoList[504] = { parentId = 1280, typeId = 1284, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[505] = { parentId = 1, typeId = 1290, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[506] = { parentId = 1, typeId = 1291, nilable = true, orgTypeId = 1290 }
_typeInfoList[507] = { parentId = 1280, typeId = 1292, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1290, 814, 10, 10, 18}, retTypeId = {}, children = {}, }
_typeInfoList[508] = { parentId = 1280, typeId = 1294, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[509] = { parentId = 1280, typeId = 1296, baseId = 1, txt = 'get_refFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[510] = { parentId = 1280, typeId = 1298, baseId = 1, txt = 'get_mutFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[511] = { parentId = 1280, typeId = 1300, baseId = 1, txt = 'get_array',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[512] = { parentId = 106, typeId = 1308, baseId = 814, txt = 'BlockNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1314, 1324, 1326, 1330}, }
_typeInfoList[513] = { parentId = 106, typeId = 1309, nilable = true, orgTypeId = 1308 }
_typeInfoList[514] = { parentId = 958, typeId = 1310, baseId = 1, txt = 'processBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1308, 8}, retTypeId = {}, children = {}, }
_typeInfoList[515] = { parentId = 1308, typeId = 1314, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[516] = { parentId = 1, typeId = 1320, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[517] = { parentId = 1, typeId = 1321, nilable = true, orgTypeId = 1320 }
_typeInfoList[518] = { parentId = 1, typeId = 1322, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[519] = { parentId = 1, typeId = 1323, nilable = true, orgTypeId = 1322 }
_typeInfoList[520] = { parentId = 1308, typeId = 1324, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1320, 18, 1322}, retTypeId = {}, children = {}, }
_typeInfoList[521] = { parentId = 1308, typeId = 1326, baseId = 1, txt = 'get_blockKind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[522] = { parentId = 1, typeId = 1328, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[523] = { parentId = 1, typeId = 1329, nilable = true, orgTypeId = 1328 }
_typeInfoList[524] = { parentId = 1308, typeId = 1330, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1328}, children = {}, }
_typeInfoList[525] = { parentId = 106, typeId = 1332, baseId = 1, txt = 'IfStmtInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1334, 1336, 1338, 1340}, }
_typeInfoList[526] = { parentId = 106, typeId = 1333, nilable = true, orgTypeId = 1332 }
_typeInfoList[527] = { parentId = 1332, typeId = 1334, baseId = 1, txt = 'get_kind',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[528] = { parentId = 1332, typeId = 1336, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[529] = { parentId = 1332, typeId = 1338, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[530] = { parentId = 1332, typeId = 1340, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {18, 814, 1308}, retTypeId = {}, children = {}, }
_typeInfoList[531] = { parentId = 106, typeId = 1346, baseId = 814, txt = 'IfNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1352, 1362, 1366}, }
_typeInfoList[532] = { parentId = 106, typeId = 1347, nilable = true, orgTypeId = 1346 }
_typeInfoList[533] = { parentId = 958, typeId = 1348, baseId = 1, txt = 'processIf',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1346, 8}, retTypeId = {}, children = {}, }
_typeInfoList[534] = { parentId = 1346, typeId = 1352, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[535] = { parentId = 1, typeId = 1358, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[536] = { parentId = 1, typeId = 1359, nilable = true, orgTypeId = 1358 }
_typeInfoList[537] = { parentId = 1, typeId = 1360, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1332}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[538] = { parentId = 1, typeId = 1361, nilable = true, orgTypeId = 1360 }
_typeInfoList[539] = { parentId = 1346, typeId = 1362, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1358, 1360}, retTypeId = {}, children = {}, }
_typeInfoList[540] = { parentId = 1, typeId = 1364, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1332}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[541] = { parentId = 1, typeId = 1365, nilable = true, orgTypeId = 1364 }
_typeInfoList[542] = { parentId = 1346, typeId = 1366, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1364}, children = {}, }
_typeInfoList[543] = { parentId = 958, typeId = 1372, baseId = 1, txt = 'processExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {994, 8}, retTypeId = {}, children = {}, }
_typeInfoList[544] = { parentId = 994, typeId = 1376, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[545] = { parentId = 1, typeId = 1382, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[546] = { parentId = 1, typeId = 1383, nilable = true, orgTypeId = 1382 }
_typeInfoList[547] = { parentId = 1, typeId = 1384, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[548] = { parentId = 1, typeId = 1385, nilable = true, orgTypeId = 1384 }
_typeInfoList[549] = { parentId = 994, typeId = 1386, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1382, 1384}, retTypeId = {}, children = {}, }
_typeInfoList[550] = { parentId = 1, typeId = 1388, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[551] = { parentId = 1, typeId = 1389, nilable = true, orgTypeId = 1388 }
_typeInfoList[552] = { parentId = 994, typeId = 1390, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1388}, children = {}, }
_typeInfoList[553] = { parentId = 106, typeId = 1392, baseId = 1, txt = 'CaseInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1394, 1396, 1398}, }
_typeInfoList[554] = { parentId = 106, typeId = 1393, nilable = true, orgTypeId = 1392 }
_typeInfoList[555] = { parentId = 1392, typeId = 1394, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {994}, children = {}, }
_typeInfoList[556] = { parentId = 1392, typeId = 1396, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[557] = { parentId = 1392, typeId = 1398, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {994, 1308}, retTypeId = {}, children = {}, }
_typeInfoList[558] = { parentId = 106, typeId = 1408, baseId = 814, txt = 'SwitchNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1414, 1424, 1426, 1430, 1432}, }
_typeInfoList[559] = { parentId = 106, typeId = 1409, nilable = true, orgTypeId = 1408 }
_typeInfoList[560] = { parentId = 958, typeId = 1410, baseId = 1, txt = 'processSwitch',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1408, 8}, retTypeId = {}, children = {}, }
_typeInfoList[561] = { parentId = 1408, typeId = 1414, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[562] = { parentId = 1, typeId = 1420, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[563] = { parentId = 1, typeId = 1421, nilable = true, orgTypeId = 1420 }
_typeInfoList[564] = { parentId = 1, typeId = 1422, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1392}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[565] = { parentId = 1, typeId = 1423, nilable = true, orgTypeId = 1422 }
_typeInfoList[566] = { parentId = 1408, typeId = 1424, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1420, 814, 1422, 1309}, retTypeId = {}, children = {}, }
_typeInfoList[567] = { parentId = 1408, typeId = 1426, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[568] = { parentId = 1, typeId = 1428, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {1392}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[569] = { parentId = 1, typeId = 1429, nilable = true, orgTypeId = 1428 }
_typeInfoList[570] = { parentId = 1408, typeId = 1430, baseId = 1, txt = 'get_caseList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1428}, children = {}, }
_typeInfoList[571] = { parentId = 1408, typeId = 1432, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[572] = { parentId = 106, typeId = 1440, baseId = 814, txt = 'WhileNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1444, 1452, 1454, 1456}, }
_typeInfoList[573] = { parentId = 106, typeId = 1441, nilable = true, orgTypeId = 1440 }
_typeInfoList[574] = { parentId = 958, typeId = 1442, baseId = 1, txt = 'processWhile',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1440, 8}, retTypeId = {}, children = {}, }
_typeInfoList[575] = { parentId = 1440, typeId = 1444, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[576] = { parentId = 1, typeId = 1450, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[577] = { parentId = 1, typeId = 1451, nilable = true, orgTypeId = 1450 }
_typeInfoList[578] = { parentId = 1440, typeId = 1452, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1450, 814, 1308}, retTypeId = {}, children = {}, }
_typeInfoList[579] = { parentId = 1440, typeId = 1454, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[580] = { parentId = 1440, typeId = 1456, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[581] = { parentId = 106, typeId = 1464, baseId = 814, txt = 'RepeatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1468, 1476, 1478, 1480}, }
_typeInfoList[582] = { parentId = 106, typeId = 1465, nilable = true, orgTypeId = 1464 }
_typeInfoList[583] = { parentId = 958, typeId = 1466, baseId = 1, txt = 'processRepeat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1464, 8}, retTypeId = {}, children = {}, }
_typeInfoList[584] = { parentId = 1464, typeId = 1468, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[585] = { parentId = 1, typeId = 1474, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[586] = { parentId = 1, typeId = 1475, nilable = true, orgTypeId = 1474 }
_typeInfoList[587] = { parentId = 1464, typeId = 1476, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1474, 1308, 814}, retTypeId = {}, children = {}, }
_typeInfoList[588] = { parentId = 1464, typeId = 1478, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[589] = { parentId = 1464, typeId = 1480, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[590] = { parentId = 106, typeId = 1494, baseId = 814, txt = 'ForNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1498, 1506, 1508, 1510, 1512, 1514, 1516}, }
_typeInfoList[591] = { parentId = 106, typeId = 1495, nilable = true, orgTypeId = 1494 }
_typeInfoList[592] = { parentId = 958, typeId = 1496, baseId = 1, txt = 'processFor',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1494, 8}, retTypeId = {}, children = {}, }
_typeInfoList[593] = { parentId = 1494, typeId = 1498, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[594] = { parentId = 1, typeId = 1504, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[595] = { parentId = 1, typeId = 1505, nilable = true, orgTypeId = 1504 }
_typeInfoList[596] = { parentId = 1494, typeId = 1506, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1504, 1308, 220, 814, 814, 815}, retTypeId = {}, children = {}, }
_typeInfoList[597] = { parentId = 1494, typeId = 1508, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[598] = { parentId = 1494, typeId = 1510, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[599] = { parentId = 1494, typeId = 1512, baseId = 1, txt = 'get_init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[600] = { parentId = 1494, typeId = 1514, baseId = 1, txt = 'get_to',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[601] = { parentId = 1494, typeId = 1516, baseId = 1, txt = 'get_delta',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {815}, children = {}, }
_typeInfoList[602] = { parentId = 106, typeId = 1526, baseId = 814, txt = 'ApplyNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1532, 1542, 1546, 1548, 1550}, }
_typeInfoList[603] = { parentId = 106, typeId = 1527, nilable = true, orgTypeId = 1526 }
_typeInfoList[604] = { parentId = 958, typeId = 1528, baseId = 1, txt = 'processApply',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1526, 8}, retTypeId = {}, children = {}, }
_typeInfoList[605] = { parentId = 1526, typeId = 1532, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[606] = { parentId = 1, typeId = 1538, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[607] = { parentId = 1, typeId = 1539, nilable = true, orgTypeId = 1538 }
_typeInfoList[608] = { parentId = 1, typeId = 1540, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[609] = { parentId = 1, typeId = 1541, nilable = true, orgTypeId = 1540 }
_typeInfoList[610] = { parentId = 1526, typeId = 1542, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1538, 1540, 814, 1308}, retTypeId = {}, children = {}, }
_typeInfoList[611] = { parentId = 1, typeId = 1544, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[612] = { parentId = 1, typeId = 1545, nilable = true, orgTypeId = 1544 }
_typeInfoList[613] = { parentId = 1526, typeId = 1546, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1544}, children = {}, }
_typeInfoList[614] = { parentId = 1526, typeId = 1548, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[615] = { parentId = 1526, typeId = 1550, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[616] = { parentId = 106, typeId = 1562, baseId = 814, txt = 'ForeachNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1566, 1574, 1576, 1578, 1580, 1582}, }
_typeInfoList[617] = { parentId = 106, typeId = 1563, nilable = true, orgTypeId = 1562 }
_typeInfoList[618] = { parentId = 958, typeId = 1564, baseId = 1, txt = 'processForeach',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1562, 8}, retTypeId = {}, children = {}, }
_typeInfoList[619] = { parentId = 1562, typeId = 1566, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[620] = { parentId = 1, typeId = 1572, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[621] = { parentId = 1, typeId = 1573, nilable = true, orgTypeId = 1572 }
_typeInfoList[622] = { parentId = 1562, typeId = 1574, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1572, 220, 221, 814, 1308}, retTypeId = {}, children = {}, }
_typeInfoList[623] = { parentId = 1562, typeId = 1576, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[624] = { parentId = 1562, typeId = 1578, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[625] = { parentId = 1562, typeId = 1580, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[626] = { parentId = 1562, typeId = 1582, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[627] = { parentId = 106, typeId = 1596, baseId = 814, txt = 'ForsortNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1600, 1608, 1610, 1612, 1614, 1616, 1618}, }
_typeInfoList[628] = { parentId = 106, typeId = 1597, nilable = true, orgTypeId = 1596 }
_typeInfoList[629] = { parentId = 958, typeId = 1598, baseId = 1, txt = 'processForsort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1596, 8}, retTypeId = {}, children = {}, }
_typeInfoList[630] = { parentId = 1596, typeId = 1600, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[631] = { parentId = 1, typeId = 1606, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[632] = { parentId = 1, typeId = 1607, nilable = true, orgTypeId = 1606 }
_typeInfoList[633] = { parentId = 1596, typeId = 1608, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1606, 220, 221, 814, 1308, 10}, retTypeId = {}, children = {}, }
_typeInfoList[634] = { parentId = 1596, typeId = 1610, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[635] = { parentId = 1596, typeId = 1612, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[636] = { parentId = 1596, typeId = 1614, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[637] = { parentId = 1596, typeId = 1616, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[638] = { parentId = 1596, typeId = 1618, baseId = 1, txt = 'get_sort',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[639] = { parentId = 106, typeId = 1624, baseId = 814, txt = 'ReturnNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1628, 1636, 1638}, }
_typeInfoList[640] = { parentId = 106, typeId = 1625, nilable = true, orgTypeId = 1624 }
_typeInfoList[641] = { parentId = 958, typeId = 1626, baseId = 1, txt = 'processReturn',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1624, 8}, retTypeId = {}, children = {}, }
_typeInfoList[642] = { parentId = 1624, typeId = 1628, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[643] = { parentId = 1, typeId = 1634, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[644] = { parentId = 1, typeId = 1635, nilable = true, orgTypeId = 1634 }
_typeInfoList[645] = { parentId = 1624, typeId = 1636, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1634, 995}, retTypeId = {}, children = {}, }
_typeInfoList[646] = { parentId = 1624, typeId = 1638, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[647] = { parentId = 106, typeId = 1642, baseId = 814, txt = 'BreakNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1646, 1654}, }
_typeInfoList[648] = { parentId = 106, typeId = 1643, nilable = true, orgTypeId = 1642 }
_typeInfoList[649] = { parentId = 958, typeId = 1644, baseId = 1, txt = 'processBreak',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1642, 8}, retTypeId = {}, children = {}, }
_typeInfoList[650] = { parentId = 1642, typeId = 1646, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[651] = { parentId = 1, typeId = 1652, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[652] = { parentId = 1, typeId = 1653, nilable = true, orgTypeId = 1652 }
_typeInfoList[653] = { parentId = 1642, typeId = 1654, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1652}, retTypeId = {}, children = {}, }
_typeInfoList[654] = { parentId = 106, typeId = 1662, baseId = 814, txt = 'ExpNewNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1666, 1674, 1676, 1678}, }
_typeInfoList[655] = { parentId = 106, typeId = 1663, nilable = true, orgTypeId = 1662 }
_typeInfoList[656] = { parentId = 958, typeId = 1664, baseId = 1, txt = 'processExpNew',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1662, 8}, retTypeId = {}, children = {}, }
_typeInfoList[657] = { parentId = 1662, typeId = 1666, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[658] = { parentId = 1, typeId = 1672, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[659] = { parentId = 1, typeId = 1673, nilable = true, orgTypeId = 1672 }
_typeInfoList[660] = { parentId = 1662, typeId = 1674, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1672, 814, 995}, retTypeId = {}, children = {}, }
_typeInfoList[661] = { parentId = 1662, typeId = 1676, baseId = 1, txt = 'get_symbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[662] = { parentId = 1662, typeId = 1678, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[663] = { parentId = 106, typeId = 1686, baseId = 814, txt = 'ExpUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1690, 1698, 1700, 1702}, }
_typeInfoList[664] = { parentId = 106, typeId = 1687, nilable = true, orgTypeId = 1686 }
_typeInfoList[665] = { parentId = 958, typeId = 1688, baseId = 1, txt = 'processExpUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1686, 8}, retTypeId = {}, children = {}, }
_typeInfoList[666] = { parentId = 1686, typeId = 1690, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[667] = { parentId = 1, typeId = 1696, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[668] = { parentId = 1, typeId = 1697, nilable = true, orgTypeId = 1696 }
_typeInfoList[669] = { parentId = 1686, typeId = 1698, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1696, 814, 815}, retTypeId = {}, children = {}, }
_typeInfoList[670] = { parentId = 1686, typeId = 1700, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[671] = { parentId = 1686, typeId = 1702, baseId = 1, txt = 'get_default',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {815}, children = {}, }
_typeInfoList[672] = { parentId = 106, typeId = 1708, baseId = 814, txt = 'ExpRefNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1712, 1720, 1722}, }
_typeInfoList[673] = { parentId = 106, typeId = 1709, nilable = true, orgTypeId = 1708 }
_typeInfoList[674] = { parentId = 958, typeId = 1710, baseId = 1, txt = 'processExpRef',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1708, 8}, retTypeId = {}, children = {}, }
_typeInfoList[675] = { parentId = 1708, typeId = 1712, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[676] = { parentId = 1, typeId = 1718, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[677] = { parentId = 1, typeId = 1719, nilable = true, orgTypeId = 1718 }
_typeInfoList[678] = { parentId = 1708, typeId = 1720, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1718, 220}, retTypeId = {}, children = {}, }
_typeInfoList[679] = { parentId = 1708, typeId = 1722, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[680] = { parentId = 106, typeId = 1732, baseId = 814, txt = 'ExpOp2Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1736, 1744, 1746, 1748, 1750}, }
_typeInfoList[681] = { parentId = 106, typeId = 1733, nilable = true, orgTypeId = 1732 }
_typeInfoList[682] = { parentId = 958, typeId = 1734, baseId = 1, txt = 'processExpOp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1732, 8}, retTypeId = {}, children = {}, }
_typeInfoList[683] = { parentId = 1732, typeId = 1736, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[684] = { parentId = 1, typeId = 1742, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[685] = { parentId = 1, typeId = 1743, nilable = true, orgTypeId = 1742 }
_typeInfoList[686] = { parentId = 1732, typeId = 1744, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1742, 220, 814, 814}, retTypeId = {}, children = {}, }
_typeInfoList[687] = { parentId = 1732, typeId = 1746, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[688] = { parentId = 1732, typeId = 1748, baseId = 1, txt = 'get_exp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[689] = { parentId = 1732, typeId = 1750, baseId = 1, txt = 'get_exp2',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[690] = { parentId = 106, typeId = 1760, baseId = 814, txt = 'UnwrapSetNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1764, 1772, 1774, 1776, 1778}, }
_typeInfoList[691] = { parentId = 106, typeId = 1761, nilable = true, orgTypeId = 1760 }
_typeInfoList[692] = { parentId = 958, typeId = 1762, baseId = 1, txt = 'processUnwrapSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1760, 8}, retTypeId = {}, children = {}, }
_typeInfoList[693] = { parentId = 1760, typeId = 1764, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[694] = { parentId = 1, typeId = 1770, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[695] = { parentId = 1, typeId = 1771, nilable = true, orgTypeId = 1770 }
_typeInfoList[696] = { parentId = 1760, typeId = 1772, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1770, 994, 994, 1309}, retTypeId = {}, children = {}, }
_typeInfoList[697] = { parentId = 1760, typeId = 1774, baseId = 1, txt = 'get_dstExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {994}, children = {}, }
_typeInfoList[698] = { parentId = 1760, typeId = 1776, baseId = 1, txt = 'get_srcExpList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {994}, children = {}, }
_typeInfoList[699] = { parentId = 1760, typeId = 1778, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[700] = { parentId = 106, typeId = 1788, baseId = 814, txt = 'IfUnwrapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1792, 1800, 1802, 1804, 1806}, }
_typeInfoList[701] = { parentId = 106, typeId = 1789, nilable = true, orgTypeId = 1788 }
_typeInfoList[702] = { parentId = 958, typeId = 1790, baseId = 1, txt = 'processIfUnwrap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1788, 8}, retTypeId = {}, children = {}, }
_typeInfoList[703] = { parentId = 1788, typeId = 1792, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[704] = { parentId = 1, typeId = 1798, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[705] = { parentId = 1, typeId = 1799, nilable = true, orgTypeId = 1798 }
_typeInfoList[706] = { parentId = 1788, typeId = 1800, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1798, 814, 1308, 1309}, retTypeId = {}, children = {}, }
_typeInfoList[707] = { parentId = 1788, typeId = 1802, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[708] = { parentId = 1788, typeId = 1804, baseId = 1, txt = 'get_block',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1308}, children = {}, }
_typeInfoList[709] = { parentId = 1788, typeId = 1806, baseId = 1, txt = 'get_nilBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[710] = { parentId = 106, typeId = 1812, baseId = 814, txt = 'ExpCastNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1816, 1824, 1826}, }
_typeInfoList[711] = { parentId = 106, typeId = 1813, nilable = true, orgTypeId = 1812 }
_typeInfoList[712] = { parentId = 958, typeId = 1814, baseId = 1, txt = 'processExpCast',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1812, 8}, retTypeId = {}, children = {}, }
_typeInfoList[713] = { parentId = 1812, typeId = 1816, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[714] = { parentId = 1, typeId = 1822, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[715] = { parentId = 1, typeId = 1823, nilable = true, orgTypeId = 1822 }
_typeInfoList[716] = { parentId = 1812, typeId = 1824, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1822, 814}, retTypeId = {}, children = {}, }
_typeInfoList[717] = { parentId = 1812, typeId = 1826, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[718] = { parentId = 106, typeId = 1836, baseId = 814, txt = 'ExpOp1Node',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1840, 1848, 1850, 1852, 1854}, }
_typeInfoList[719] = { parentId = 106, typeId = 1837, nilable = true, orgTypeId = 1836 }
_typeInfoList[720] = { parentId = 958, typeId = 1838, baseId = 1, txt = 'processExpOp1',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1836, 8}, retTypeId = {}, children = {}, }
_typeInfoList[721] = { parentId = 1836, typeId = 1840, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[722] = { parentId = 1, typeId = 1846, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[723] = { parentId = 1, typeId = 1847, nilable = true, orgTypeId = 1846 }
_typeInfoList[724] = { parentId = 1836, typeId = 1848, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1846, 220, 18, 814}, retTypeId = {}, children = {}, }
_typeInfoList[725] = { parentId = 1836, typeId = 1850, baseId = 1, txt = 'get_op',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[726] = { parentId = 1836, typeId = 1852, baseId = 1, txt = 'get_macroMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[727] = { parentId = 1836, typeId = 1854, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[728] = { parentId = 106, typeId = 1862, baseId = 814, txt = 'ExpRefItemNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1866, 1874, 1876, 1878}, }
_typeInfoList[729] = { parentId = 106, typeId = 1863, nilable = true, orgTypeId = 1862 }
_typeInfoList[730] = { parentId = 958, typeId = 1864, baseId = 1, txt = 'processExpRefItem',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1862, 8}, retTypeId = {}, children = {}, }
_typeInfoList[731] = { parentId = 1862, typeId = 1866, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[732] = { parentId = 1, typeId = 1872, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[733] = { parentId = 1, typeId = 1873, nilable = true, orgTypeId = 1872 }
_typeInfoList[734] = { parentId = 1862, typeId = 1874, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1872, 814, 814}, retTypeId = {}, children = {}, }
_typeInfoList[735] = { parentId = 1862, typeId = 1876, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[736] = { parentId = 1862, typeId = 1878, baseId = 1, txt = 'get_index',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[737] = { parentId = 106, typeId = 1886, baseId = 814, txt = 'ExpCallNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1890, 1898, 1900, 1902}, }
_typeInfoList[738] = { parentId = 106, typeId = 1887, nilable = true, orgTypeId = 1886 }
_typeInfoList[739] = { parentId = 958, typeId = 1888, baseId = 1, txt = 'processExpCall',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1886, 8}, retTypeId = {}, children = {}, }
_typeInfoList[740] = { parentId = 1886, typeId = 1890, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[741] = { parentId = 1, typeId = 1896, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[742] = { parentId = 1, typeId = 1897, nilable = true, orgTypeId = 1896 }
_typeInfoList[743] = { parentId = 1886, typeId = 1898, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1896, 814, 995}, retTypeId = {}, children = {}, }
_typeInfoList[744] = { parentId = 1886, typeId = 1900, baseId = 1, txt = 'get_func',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[745] = { parentId = 1886, typeId = 1902, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[746] = { parentId = 106, typeId = 1908, baseId = 814, txt = 'ExpDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1912, 1920, 1922}, }
_typeInfoList[747] = { parentId = 106, typeId = 1909, nilable = true, orgTypeId = 1908 }
_typeInfoList[748] = { parentId = 958, typeId = 1910, baseId = 1, txt = 'processExpDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1908, 8}, retTypeId = {}, children = {}, }
_typeInfoList[749] = { parentId = 1908, typeId = 1912, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[750] = { parentId = 1, typeId = 1918, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[751] = { parentId = 1, typeId = 1919, nilable = true, orgTypeId = 1918 }
_typeInfoList[752] = { parentId = 1908, typeId = 1920, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1918, 220}, retTypeId = {}, children = {}, }
_typeInfoList[753] = { parentId = 1908, typeId = 1922, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[754] = { parentId = 106, typeId = 1928, baseId = 814, txt = 'ExpParenNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1932, 1940, 1942}, }
_typeInfoList[755] = { parentId = 106, typeId = 1929, nilable = true, orgTypeId = 1928 }
_typeInfoList[756] = { parentId = 958, typeId = 1930, baseId = 1, txt = 'processExpParen',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1928, 8}, retTypeId = {}, children = {}, }
_typeInfoList[757] = { parentId = 1928, typeId = 1932, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[758] = { parentId = 1, typeId = 1938, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[759] = { parentId = 1, typeId = 1939, nilable = true, orgTypeId = 1938 }
_typeInfoList[760] = { parentId = 1928, typeId = 1940, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1938, 814}, retTypeId = {}, children = {}, }
_typeInfoList[761] = { parentId = 1928, typeId = 1942, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[762] = { parentId = 106, typeId = 1948, baseId = 814, txt = 'ExpMacroExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1954, 1964, 1968}, }
_typeInfoList[763] = { parentId = 106, typeId = 1949, nilable = true, orgTypeId = 1948 }
_typeInfoList[764] = { parentId = 958, typeId = 1950, baseId = 1, txt = 'processExpMacroExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1948, 8}, retTypeId = {}, children = {}, }
_typeInfoList[765] = { parentId = 1948, typeId = 1954, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[766] = { parentId = 1, typeId = 1960, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[767] = { parentId = 1, typeId = 1961, nilable = true, orgTypeId = 1960 }
_typeInfoList[768] = { parentId = 1, typeId = 1962, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[769] = { parentId = 1, typeId = 1963, nilable = true, orgTypeId = 1962 }
_typeInfoList[770] = { parentId = 1948, typeId = 1964, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1960, 1962}, retTypeId = {}, children = {}, }
_typeInfoList[771] = { parentId = 1, typeId = 1966, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[772] = { parentId = 1, typeId = 1967, nilable = true, orgTypeId = 1966 }
_typeInfoList[773] = { parentId = 1948, typeId = 1968, baseId = 1, txt = 'get_stmtList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1966}, children = {}, }
_typeInfoList[774] = { parentId = 106, typeId = 1974, baseId = 814, txt = 'ExpMacroStatNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {1980, 1990, 1994, 2802}, }
_typeInfoList[775] = { parentId = 106, typeId = 1975, nilable = true, orgTypeId = 1974 }
_typeInfoList[776] = { parentId = 958, typeId = 1976, baseId = 1, txt = 'processExpMacroStat',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {1974, 8}, retTypeId = {}, children = {}, }
_typeInfoList[777] = { parentId = 1974, typeId = 1980, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[778] = { parentId = 1, typeId = 1986, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[779] = { parentId = 1, typeId = 1987, nilable = true, orgTypeId = 1986 }
_typeInfoList[780] = { parentId = 1, typeId = 1988, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[781] = { parentId = 1, typeId = 1989, nilable = true, orgTypeId = 1988 }
_typeInfoList[782] = { parentId = 1974, typeId = 1990, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 1986, 1988}, retTypeId = {}, children = {}, }
_typeInfoList[783] = { parentId = 1, typeId = 1992, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[784] = { parentId = 1, typeId = 1993, nilable = true, orgTypeId = 1992 }
_typeInfoList[785] = { parentId = 1974, typeId = 1994, baseId = 1, txt = 'get_expStrList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1992}, children = {}, }
_typeInfoList[786] = { parentId = 106, typeId = 2000, baseId = 814, txt = 'StmtExpNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2004, 2012, 2014}, }
_typeInfoList[787] = { parentId = 106, typeId = 2001, nilable = true, orgTypeId = 2000 }
_typeInfoList[788] = { parentId = 958, typeId = 2002, baseId = 1, txt = 'processStmtExp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2000, 8}, retTypeId = {}, children = {}, }
_typeInfoList[789] = { parentId = 2000, typeId = 2004, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[790] = { parentId = 1, typeId = 2010, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[791] = { parentId = 1, typeId = 2011, nilable = true, orgTypeId = 2010 }
_typeInfoList[792] = { parentId = 2000, typeId = 2012, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2010, 814}, retTypeId = {}, children = {}, }
_typeInfoList[793] = { parentId = 2000, typeId = 2014, baseId = 1, txt = 'get_exp',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[794] = { parentId = 106, typeId = 2022, baseId = 814, txt = 'RefFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2026, 2034, 2036, 2038, 2790}, }
_typeInfoList[795] = { parentId = 106, typeId = 2023, nilable = true, orgTypeId = 2022 }
_typeInfoList[796] = { parentId = 958, typeId = 2024, baseId = 1, txt = 'processRefField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2022, 8}, retTypeId = {}, children = {}, }
_typeInfoList[797] = { parentId = 2022, typeId = 2026, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[798] = { parentId = 1, typeId = 2032, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[799] = { parentId = 1, typeId = 2033, nilable = true, orgTypeId = 2032 }
_typeInfoList[800] = { parentId = 2022, typeId = 2034, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2032, 220, 814}, retTypeId = {}, children = {}, }
_typeInfoList[801] = { parentId = 2022, typeId = 2036, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[802] = { parentId = 2022, typeId = 2038, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[803] = { parentId = 106, typeId = 2048, baseId = 814, txt = 'GetFieldNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2052, 2060, 2062, 2064, 2066}, }
_typeInfoList[804] = { parentId = 106, typeId = 2049, nilable = true, orgTypeId = 2048 }
_typeInfoList[805] = { parentId = 958, typeId = 2050, baseId = 1, txt = 'processGetField',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2048, 8}, retTypeId = {}, children = {}, }
_typeInfoList[806] = { parentId = 2048, typeId = 2052, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[807] = { parentId = 1, typeId = 2058, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[808] = { parentId = 1, typeId = 2059, nilable = true, orgTypeId = 2058 }
_typeInfoList[809] = { parentId = 2048, typeId = 2060, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2058, 220, 814, 660}, retTypeId = {}, children = {}, }
_typeInfoList[810] = { parentId = 2048, typeId = 2062, baseId = 1, txt = 'get_field',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[811] = { parentId = 2048, typeId = 2064, baseId = 1, txt = 'get_prefix',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[812] = { parentId = 2048, typeId = 2066, baseId = 1, txt = 'get_getterTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[813] = { parentId = 106, typeId = 2068, baseId = 1, txt = 'VarInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2070, 2072, 2074, 2076}, }
_typeInfoList[814] = { parentId = 106, typeId = 2069, nilable = true, orgTypeId = 2068 }
_typeInfoList[815] = { parentId = 2068, typeId = 2070, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[816] = { parentId = 2068, typeId = 2072, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1281}, children = {}, }
_typeInfoList[817] = { parentId = 2068, typeId = 2074, baseId = 1, txt = 'get_actualType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[818] = { parentId = 2068, typeId = 2076, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {220, 1281, 660}, retTypeId = {}, children = {}, }
_typeInfoList[819] = { parentId = 106, typeId = 2102, baseId = 814, txt = 'DeclVarNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2112, 2126, 2128, 2130, 2132, 2136, 2138, 2142, 2144, 2146, 2148, 2152, 2154}, }
_typeInfoList[820] = { parentId = 106, typeId = 2103, nilable = true, orgTypeId = 2102 }
_typeInfoList[821] = { parentId = 958, typeId = 2104, baseId = 1, txt = 'processDeclVar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2102, 8}, retTypeId = {}, children = {}, }
_typeInfoList[822] = { parentId = 2102, typeId = 2112, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[823] = { parentId = 1, typeId = 2118, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[824] = { parentId = 1, typeId = 2119, nilable = true, orgTypeId = 2118 }
_typeInfoList[825] = { parentId = 1, typeId = 2120, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2068}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[826] = { parentId = 1, typeId = 2121, nilable = true, orgTypeId = 2120 }
_typeInfoList[827] = { parentId = 1, typeId = 2122, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[828] = { parentId = 1, typeId = 2123, nilable = true, orgTypeId = 2122 }
_typeInfoList[829] = { parentId = 1, typeId = 2124, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2068}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[830] = { parentId = 1, typeId = 2125, nilable = true, orgTypeId = 2124 }
_typeInfoList[831] = { parentId = 2102, typeId = 2126, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2118, 18, 18, 10, 2120, 995, 2122, 10, 1309, 1309, 2124, 1309}, retTypeId = {}, children = {}, }
_typeInfoList[832] = { parentId = 2102, typeId = 2128, baseId = 1, txt = 'get_mode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[833] = { parentId = 2102, typeId = 2130, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[834] = { parentId = 2102, typeId = 2132, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[835] = { parentId = 1, typeId = 2134, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2068}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[836] = { parentId = 1, typeId = 2135, nilable = true, orgTypeId = 2134 }
_typeInfoList[837] = { parentId = 2102, typeId = 2136, baseId = 1, txt = 'get_varList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2134}, children = {}, }
_typeInfoList[838] = { parentId = 2102, typeId = 2138, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[839] = { parentId = 1, typeId = 2140, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[840] = { parentId = 1, typeId = 2141, nilable = true, orgTypeId = 2140 }
_typeInfoList[841] = { parentId = 2102, typeId = 2142, baseId = 1, txt = 'get_typeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2140}, children = {}, }
_typeInfoList[842] = { parentId = 2102, typeId = 2144, baseId = 1, txt = 'get_unwrapFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[843] = { parentId = 2102, typeId = 2146, baseId = 1, txt = 'get_unwrapBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[844] = { parentId = 2102, typeId = 2148, baseId = 1, txt = 'get_thenBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[845] = { parentId = 1, typeId = 2150, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2068}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[846] = { parentId = 1, typeId = 2151, nilable = true, orgTypeId = 2150 }
_typeInfoList[847] = { parentId = 2102, typeId = 2152, baseId = 1, txt = 'get_syncVarList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2150}, children = {}, }
_typeInfoList[848] = { parentId = 2102, typeId = 2154, baseId = 1, txt = 'get_syncBlock',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1309}, children = {}, }
_typeInfoList[849] = { parentId = 106, typeId = 2156, baseId = 1, txt = 'DeclFuncInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2162, 2164, 2168, 2170, 2172, 2174, 2178, 2180}, }
_typeInfoList[850] = { parentId = 106, typeId = 2157, nilable = true, orgTypeId = 2156 }
_typeInfoList[851] = { parentId = 1, typeId = 2158, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[852] = { parentId = 1, typeId = 2159, nilable = true, orgTypeId = 2158 }
_typeInfoList[853] = { parentId = 1, typeId = 2160, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pri', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[854] = { parentId = 1, typeId = 2161, nilable = true, orgTypeId = 2160 }
_typeInfoList[855] = { parentId = 2156, typeId = 2162, baseId = 1, txt = 'get_className',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[856] = { parentId = 2156, typeId = 2164, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[857] = { parentId = 1, typeId = 2166, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[858] = { parentId = 1, typeId = 2167, nilable = true, orgTypeId = 2166 }
_typeInfoList[859] = { parentId = 2156, typeId = 2168, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2166}, children = {}, }
_typeInfoList[860] = { parentId = 2156, typeId = 2170, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[861] = { parentId = 2156, typeId = 2172, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[862] = { parentId = 2156, typeId = 2174, baseId = 1, txt = 'get_body',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {815}, children = {}, }
_typeInfoList[863] = { parentId = 1, typeId = 2176, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[864] = { parentId = 1, typeId = 2177, nilable = true, orgTypeId = 2176 }
_typeInfoList[865] = { parentId = 2156, typeId = 2178, baseId = 1, txt = 'get_retTypeInfoList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2176}, children = {}, }
_typeInfoList[866] = { parentId = 2156, typeId = 2180, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {221, 221, 2158, 10, 18, 815, 2160}, retTypeId = {}, children = {}, }
_typeInfoList[867] = { parentId = 106, typeId = 2186, baseId = 814, txt = 'DeclFuncNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2190, 2198, 2200}, }
_typeInfoList[868] = { parentId = 106, typeId = 2187, nilable = true, orgTypeId = 2186 }
_typeInfoList[869] = { parentId = 958, typeId = 2188, baseId = 1, txt = 'processDeclFunc',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2186, 8}, retTypeId = {}, children = {}, }
_typeInfoList[870] = { parentId = 2186, typeId = 2190, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[871] = { parentId = 1, typeId = 2196, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[872] = { parentId = 1, typeId = 2197, nilable = true, orgTypeId = 2196 }
_typeInfoList[873] = { parentId = 2186, typeId = 2198, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2196, 2156}, retTypeId = {}, children = {}, }
_typeInfoList[874] = { parentId = 2186, typeId = 2200, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2156}, children = {}, }
_typeInfoList[875] = { parentId = 106, typeId = 2206, baseId = 814, txt = 'DeclMethodNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2210, 2218, 2220}, }
_typeInfoList[876] = { parentId = 106, typeId = 2207, nilable = true, orgTypeId = 2206 }
_typeInfoList[877] = { parentId = 958, typeId = 2208, baseId = 1, txt = 'processDeclMethod',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2206, 8}, retTypeId = {}, children = {}, }
_typeInfoList[878] = { parentId = 2206, typeId = 2210, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[879] = { parentId = 1, typeId = 2216, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[880] = { parentId = 1, typeId = 2217, nilable = true, orgTypeId = 2216 }
_typeInfoList[881] = { parentId = 2206, typeId = 2218, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2216, 2156}, retTypeId = {}, children = {}, }
_typeInfoList[882] = { parentId = 2206, typeId = 2220, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2156}, children = {}, }
_typeInfoList[883] = { parentId = 106, typeId = 2226, baseId = 814, txt = 'DeclConstrNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2230, 2238, 2240}, }
_typeInfoList[884] = { parentId = 106, typeId = 2227, nilable = true, orgTypeId = 2226 }
_typeInfoList[885] = { parentId = 958, typeId = 2228, baseId = 1, txt = 'processDeclConstr',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2226, 8}, retTypeId = {}, children = {}, }
_typeInfoList[886] = { parentId = 2226, typeId = 2230, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[887] = { parentId = 1, typeId = 2236, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[888] = { parentId = 1, typeId = 2237, nilable = true, orgTypeId = 2236 }
_typeInfoList[889] = { parentId = 2226, typeId = 2238, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2236, 2156}, retTypeId = {}, children = {}, }
_typeInfoList[890] = { parentId = 2226, typeId = 2240, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2156}, children = {}, }
_typeInfoList[891] = { parentId = 106, typeId = 2248, baseId = 814, txt = 'ExpCallSuperNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2252, 2260, 2262, 2264}, }
_typeInfoList[892] = { parentId = 106, typeId = 2249, nilable = true, orgTypeId = 2248 }
_typeInfoList[893] = { parentId = 958, typeId = 2250, baseId = 1, txt = 'processExpCallSuper',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2248, 8}, retTypeId = {}, children = {}, }
_typeInfoList[894] = { parentId = 2248, typeId = 2252, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[895] = { parentId = 1, typeId = 2258, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[896] = { parentId = 1, typeId = 2259, nilable = true, orgTypeId = 2258 }
_typeInfoList[897] = { parentId = 2248, typeId = 2260, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2258, 660, 994}, retTypeId = {}, children = {}, }
_typeInfoList[898] = { parentId = 2248, typeId = 2262, baseId = 1, txt = 'get_superType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[899] = { parentId = 2248, typeId = 2264, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {994}, children = {}, }
_typeInfoList[900] = { parentId = 106, typeId = 2280, baseId = 814, txt = 'DeclMemberNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2284, 2292, 2294, 2296, 2298, 2300, 2302, 2304}, }
_typeInfoList[901] = { parentId = 106, typeId = 2281, nilable = true, orgTypeId = 2280 }
_typeInfoList[902] = { parentId = 958, typeId = 2282, baseId = 1, txt = 'processDeclMember',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2280, 8}, retTypeId = {}, children = {}, }
_typeInfoList[903] = { parentId = 2280, typeId = 2284, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[904] = { parentId = 1, typeId = 2290, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[905] = { parentId = 1, typeId = 2291, nilable = true, orgTypeId = 2290 }
_typeInfoList[906] = { parentId = 2280, typeId = 2292, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2290, 220, 1280, 10, 18, 18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[907] = { parentId = 2280, typeId = 2294, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[908] = { parentId = 2280, typeId = 2296, baseId = 1, txt = 'get_refType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1280}, children = {}, }
_typeInfoList[909] = { parentId = 2280, typeId = 2298, baseId = 1, txt = 'get_staticFlag',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {10}, children = {}, }
_typeInfoList[910] = { parentId = 2280, typeId = 2300, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[911] = { parentId = 2280, typeId = 2302, baseId = 1, txt = 'get_getterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[912] = { parentId = 2280, typeId = 2304, baseId = 1, txt = 'get_setterMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[913] = { parentId = 106, typeId = 2312, baseId = 814, txt = 'DeclArgNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2316, 2324, 2326, 2328}, }
_typeInfoList[914] = { parentId = 106, typeId = 2313, nilable = true, orgTypeId = 2312 }
_typeInfoList[915] = { parentId = 958, typeId = 2314, baseId = 1, txt = 'processDeclArg',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2312, 8}, retTypeId = {}, children = {}, }
_typeInfoList[916] = { parentId = 2312, typeId = 2316, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[917] = { parentId = 1, typeId = 2322, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[918] = { parentId = 1, typeId = 2323, nilable = true, orgTypeId = 2322 }
_typeInfoList[919] = { parentId = 2312, typeId = 2324, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2322, 220, 1280}, retTypeId = {}, children = {}, }
_typeInfoList[920] = { parentId = 2312, typeId = 2326, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[921] = { parentId = 2312, typeId = 2328, baseId = 1, txt = 'get_argType',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {1280}, children = {}, }
_typeInfoList[922] = { parentId = 106, typeId = 2332, baseId = 814, txt = 'DeclArgDDDNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2336, 2344}, }
_typeInfoList[923] = { parentId = 106, typeId = 2333, nilable = true, orgTypeId = 2332 }
_typeInfoList[924] = { parentId = 958, typeId = 2334, baseId = 1, txt = 'processDeclArgDDD',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2332, 8}, retTypeId = {}, children = {}, }
_typeInfoList[925] = { parentId = 2332, typeId = 2336, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[926] = { parentId = 1, typeId = 2342, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[927] = { parentId = 1, typeId = 2343, nilable = true, orgTypeId = 2342 }
_typeInfoList[928] = { parentId = 2332, typeId = 2344, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2342}, retTypeId = {}, children = {}, }
_typeInfoList[929] = { parentId = 958, typeId = 2360, baseId = 1, txt = 'processDeclClass',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {816, 8}, retTypeId = {}, children = {}, }
_typeInfoList[930] = { parentId = 816, typeId = 2368, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[931] = { parentId = 1, typeId = 2374, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[932] = { parentId = 1, typeId = 2375, nilable = true, orgTypeId = 2374 }
_typeInfoList[933] = { parentId = 1, typeId = 2376, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[934] = { parentId = 1, typeId = 2377, nilable = true, orgTypeId = 2376 }
_typeInfoList[935] = { parentId = 1, typeId = 2378, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2280}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[936] = { parentId = 1, typeId = 2379, nilable = true, orgTypeId = 2378 }
_typeInfoList[937] = { parentId = 1, typeId = 2380, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[938] = { parentId = 1, typeId = 2381, nilable = true, orgTypeId = 2380 }
_typeInfoList[939] = { parentId = 816, typeId = 2382, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2374, 18, 220, 2376, 2378, 686, 2380}, retTypeId = {}, children = {}, }
_typeInfoList[940] = { parentId = 816, typeId = 2384, baseId = 1, txt = 'get_accessMode',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[941] = { parentId = 816, typeId = 2386, baseId = 1, txt = 'get_name',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[942] = { parentId = 1, typeId = 2388, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[943] = { parentId = 1, typeId = 2389, nilable = true, orgTypeId = 2388 }
_typeInfoList[944] = { parentId = 816, typeId = 2390, baseId = 1, txt = 'get_fieldList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2388}, children = {}, }
_typeInfoList[945] = { parentId = 1, typeId = 2392, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2280}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[946] = { parentId = 1, typeId = 2393, nilable = true, orgTypeId = 2392 }
_typeInfoList[947] = { parentId = 816, typeId = 2394, baseId = 1, txt = 'get_memberList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2392}, children = {}, }
_typeInfoList[948] = { parentId = 816, typeId = 2396, baseId = 1, txt = 'get_scope',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {686}, children = {}, }
_typeInfoList[949] = { parentId = 1, typeId = 2398, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 10}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[950] = { parentId = 1, typeId = 2399, nilable = true, orgTypeId = 2398 }
_typeInfoList[951] = { parentId = 816, typeId = 2400, baseId = 1, txt = 'get_outerMethodSet',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2398}, children = {}, }
_typeInfoList[952] = { parentId = 106, typeId = 2406, baseId = 814, txt = 'DeclMacroNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2410, 2418, 2420}, }
_typeInfoList[953] = { parentId = 106, typeId = 2407, nilable = true, orgTypeId = 2406 }
_typeInfoList[954] = { parentId = 958, typeId = 2408, baseId = 1, txt = 'processDeclMacro',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2406, 8}, retTypeId = {}, children = {}, }
_typeInfoList[955] = { parentId = 2406, typeId = 2410, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[956] = { parentId = 1, typeId = 2416, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[957] = { parentId = 1, typeId = 2417, nilable = true, orgTypeId = 2416 }
_typeInfoList[958] = { parentId = 2406, typeId = 2418, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2416, 996}, retTypeId = {}, children = {}, }
_typeInfoList[959] = { parentId = 2406, typeId = 2420, baseId = 1, txt = 'get_declInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {996}, children = {}, }
_typeInfoList[960] = { parentId = 106, typeId = 2424, baseId = 814, txt = 'LiteralNilNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2428, 2436, 2668}, }
_typeInfoList[961] = { parentId = 106, typeId = 2425, nilable = true, orgTypeId = 2424 }
_typeInfoList[962] = { parentId = 958, typeId = 2426, baseId = 1, txt = 'processLiteralNil',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2424, 8}, retTypeId = {}, children = {}, }
_typeInfoList[963] = { parentId = 2424, typeId = 2428, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[964] = { parentId = 1, typeId = 2434, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[965] = { parentId = 1, typeId = 2435, nilable = true, orgTypeId = 2434 }
_typeInfoList[966] = { parentId = 2424, typeId = 2436, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2434}, retTypeId = {}, children = {}, }
_typeInfoList[967] = { parentId = 106, typeId = 2444, baseId = 814, txt = 'LiteralCharNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2448, 2456, 2458, 2460, 2678}, }
_typeInfoList[968] = { parentId = 106, typeId = 2445, nilable = true, orgTypeId = 2444 }
_typeInfoList[969] = { parentId = 958, typeId = 2446, baseId = 1, txt = 'processLiteralChar',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2444, 8}, retTypeId = {}, children = {}, }
_typeInfoList[970] = { parentId = 2444, typeId = 2448, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[971] = { parentId = 1, typeId = 2454, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[972] = { parentId = 1, typeId = 2455, nilable = true, orgTypeId = 2454 }
_typeInfoList[973] = { parentId = 2444, typeId = 2456, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2454, 220, 12}, retTypeId = {}, children = {}, }
_typeInfoList[974] = { parentId = 2444, typeId = 2458, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[975] = { parentId = 2444, typeId = 2460, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[976] = { parentId = 106, typeId = 2468, baseId = 814, txt = 'LiteralIntNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2472, 2480, 2482, 2484, 2688}, }
_typeInfoList[977] = { parentId = 106, typeId = 2469, nilable = true, orgTypeId = 2468 }
_typeInfoList[978] = { parentId = 958, typeId = 2470, baseId = 1, txt = 'processLiteralInt',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2468, 8}, retTypeId = {}, children = {}, }
_typeInfoList[979] = { parentId = 2468, typeId = 2472, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[980] = { parentId = 1, typeId = 2478, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[981] = { parentId = 1, typeId = 2479, nilable = true, orgTypeId = 2478 }
_typeInfoList[982] = { parentId = 2468, typeId = 2480, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2478, 220, 12}, retTypeId = {}, children = {}, }
_typeInfoList[983] = { parentId = 2468, typeId = 2482, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[984] = { parentId = 2468, typeId = 2484, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {12}, children = {}, }
_typeInfoList[985] = { parentId = 106, typeId = 2492, baseId = 814, txt = 'LiteralRealNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2496, 2504, 2506, 2508, 2698}, }
_typeInfoList[986] = { parentId = 106, typeId = 2493, nilable = true, orgTypeId = 2492 }
_typeInfoList[987] = { parentId = 958, typeId = 2494, baseId = 1, txt = 'processLiteralReal',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2492, 8}, retTypeId = {}, children = {}, }
_typeInfoList[988] = { parentId = 2492, typeId = 2496, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[989] = { parentId = 1, typeId = 2502, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[990] = { parentId = 1, typeId = 2503, nilable = true, orgTypeId = 2502 }
_typeInfoList[991] = { parentId = 2492, typeId = 2504, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2502, 220, 14}, retTypeId = {}, children = {}, }
_typeInfoList[992] = { parentId = 2492, typeId = 2506, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[993] = { parentId = 2492, typeId = 2508, baseId = 1, txt = 'get_num',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {14}, children = {}, }
_typeInfoList[994] = { parentId = 106, typeId = 2514, baseId = 814, txt = 'LiteralArrayNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2518, 2526, 2528, 2708}, }
_typeInfoList[995] = { parentId = 106, typeId = 2515, nilable = true, orgTypeId = 2514 }
_typeInfoList[996] = { parentId = 958, typeId = 2516, baseId = 1, txt = 'processLiteralArray',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2514, 8}, retTypeId = {}, children = {}, }
_typeInfoList[997] = { parentId = 2514, typeId = 2518, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[998] = { parentId = 1, typeId = 2524, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[999] = { parentId = 1, typeId = 2525, nilable = true, orgTypeId = 2524 }
_typeInfoList[1000] = { parentId = 2514, typeId = 2526, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2524, 995}, retTypeId = {}, children = {}, }
_typeInfoList[1001] = { parentId = 2514, typeId = 2528, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[1002] = { parentId = 106, typeId = 2534, baseId = 814, txt = 'LiteralListNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2538, 2546, 2548, 2722}, }
_typeInfoList[1003] = { parentId = 106, typeId = 2535, nilable = true, orgTypeId = 2534 }
_typeInfoList[1004] = { parentId = 958, typeId = 2536, baseId = 1, txt = 'processLiteralList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2534, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1005] = { parentId = 2534, typeId = 2538, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1006] = { parentId = 1, typeId = 2544, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1007] = { parentId = 1, typeId = 2545, nilable = true, orgTypeId = 2544 }
_typeInfoList[1008] = { parentId = 2534, typeId = 2546, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2544, 995}, retTypeId = {}, children = {}, }
_typeInfoList[1009] = { parentId = 2534, typeId = 2548, baseId = 1, txt = 'get_expList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {995}, children = {}, }
_typeInfoList[1010] = { parentId = 106, typeId = 2550, baseId = 1, txt = 'PairItem',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2552, 2554, 2556}, }
_typeInfoList[1011] = { parentId = 106, typeId = 2551, nilable = true, orgTypeId = 2550 }
_typeInfoList[1012] = { parentId = 2550, typeId = 2552, baseId = 1, txt = 'get_key',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[1013] = { parentId = 2550, typeId = 2554, baseId = 1, txt = 'get_val',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[1014] = { parentId = 2550, typeId = 2556, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {814, 814}, retTypeId = {}, children = {}, }
_typeInfoList[1015] = { parentId = 106, typeId = 2564, baseId = 814, txt = 'LiteralMapNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2572, 2584, 2588, 2592, 2736}, }
_typeInfoList[1016] = { parentId = 106, typeId = 2565, nilable = true, orgTypeId = 2564 }
_typeInfoList[1017] = { parentId = 958, typeId = 2566, baseId = 1, txt = 'processLiteralMap',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2564, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1018] = { parentId = 2564, typeId = 2572, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1019] = { parentId = 1, typeId = 2578, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1020] = { parentId = 1, typeId = 2579, nilable = true, orgTypeId = 2578 }
_typeInfoList[1021] = { parentId = 1, typeId = 2580, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {814, 814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1022] = { parentId = 1, typeId = 2581, nilable = true, orgTypeId = 2580 }
_typeInfoList[1023] = { parentId = 1, typeId = 2582, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2550}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1024] = { parentId = 1, typeId = 2583, nilable = true, orgTypeId = 2582 }
_typeInfoList[1025] = { parentId = 2564, typeId = 2584, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2578, 2580, 2582}, retTypeId = {}, children = {}, }
_typeInfoList[1026] = { parentId = 1, typeId = 2586, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {814, 814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1027] = { parentId = 1, typeId = 2587, nilable = true, orgTypeId = 2586 }
_typeInfoList[1028] = { parentId = 2564, typeId = 2588, baseId = 1, txt = 'get_map',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2586}, children = {}, }
_typeInfoList[1029] = { parentId = 1, typeId = 2590, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2550}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1030] = { parentId = 1, typeId = 2591, nilable = true, orgTypeId = 2590 }
_typeInfoList[1031] = { parentId = 2564, typeId = 2592, baseId = 1, txt = 'get_pairList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2590}, children = {}, }
_typeInfoList[1032] = { parentId = 106, typeId = 2600, baseId = 814, txt = 'LiteralStringNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2606, 2616, 2618, 2622, 2750}, }
_typeInfoList[1033] = { parentId = 106, typeId = 2601, nilable = true, orgTypeId = 2600 }
_typeInfoList[1034] = { parentId = 958, typeId = 2602, baseId = 1, txt = 'processLiteralString',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2600, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1035] = { parentId = 2600, typeId = 2606, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1036] = { parentId = 1, typeId = 2612, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1037] = { parentId = 1, typeId = 2613, nilable = true, orgTypeId = 2612 }
_typeInfoList[1038] = { parentId = 1, typeId = 2614, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1039] = { parentId = 1, typeId = 2615, nilable = true, orgTypeId = 2614 }
_typeInfoList[1040] = { parentId = 2600, typeId = 2616, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2612, 220, 2614}, retTypeId = {}, children = {}, }
_typeInfoList[1041] = { parentId = 2600, typeId = 2618, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[1042] = { parentId = 1, typeId = 2620, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1043] = { parentId = 1, typeId = 2621, nilable = true, orgTypeId = 2620 }
_typeInfoList[1044] = { parentId = 2600, typeId = 2622, baseId = 1, txt = 'get_argList',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2620}, children = {}, }
_typeInfoList[1045] = { parentId = 106, typeId = 2628, baseId = 814, txt = 'LiteralBoolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2632, 2640, 2642, 2768}, }
_typeInfoList[1046] = { parentId = 106, typeId = 2629, nilable = true, orgTypeId = 2628 }
_typeInfoList[1047] = { parentId = 958, typeId = 2630, baseId = 1, txt = 'processLiteralBool',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2628, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1048] = { parentId = 2628, typeId = 2632, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1049] = { parentId = 1, typeId = 2638, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1050] = { parentId = 1, typeId = 2639, nilable = true, orgTypeId = 2638 }
_typeInfoList[1051] = { parentId = 2628, typeId = 2640, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2638, 220}, retTypeId = {}, children = {}, }
_typeInfoList[1052] = { parentId = 2628, typeId = 2642, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[1053] = { parentId = 106, typeId = 2648, baseId = 814, txt = 'LiteralSymbolNode',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2652, 2660, 2662, 2778}, }
_typeInfoList[1054] = { parentId = 106, typeId = 2649, nilable = true, orgTypeId = 2648 }
_typeInfoList[1055] = { parentId = 958, typeId = 2650, baseId = 1, txt = 'processLiteralSymbol',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2648, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1056] = { parentId = 2648, typeId = 2652, baseId = 1, txt = 'processFilter',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {958, 8}, retTypeId = {}, children = {}, }
_typeInfoList[1057] = { parentId = 1, typeId = 2658, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1058] = { parentId = 1, typeId = 2659, nilable = true, orgTypeId = 2658 }
_typeInfoList[1059] = { parentId = 2648, typeId = 2660, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {216, 2658, 220}, retTypeId = {}, children = {}, }
_typeInfoList[1060] = { parentId = 2648, typeId = 2662, baseId = 1, txt = 'get_token',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[1061] = { parentId = 1, typeId = 2664, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1062] = { parentId = 1, typeId = 2665, nilable = true, orgTypeId = 2664 }
_typeInfoList[1063] = { parentId = 1, typeId = 2666, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1064] = { parentId = 1, typeId = 2667, nilable = true, orgTypeId = 2666 }
_typeInfoList[1065] = { parentId = 2424, typeId = 2668, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2664, 2666}, children = {}, }
_typeInfoList[1066] = { parentId = 1, typeId = 2674, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1067] = { parentId = 1, typeId = 2675, nilable = true, orgTypeId = 2674 }
_typeInfoList[1068] = { parentId = 1, typeId = 2676, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1069] = { parentId = 1, typeId = 2677, nilable = true, orgTypeId = 2676 }
_typeInfoList[1070] = { parentId = 2444, typeId = 2678, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2674, 2676}, children = {}, }
_typeInfoList[1071] = { parentId = 1, typeId = 2684, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1072] = { parentId = 1, typeId = 2685, nilable = true, orgTypeId = 2684 }
_typeInfoList[1073] = { parentId = 1, typeId = 2686, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1074] = { parentId = 1, typeId = 2687, nilable = true, orgTypeId = 2686 }
_typeInfoList[1075] = { parentId = 2468, typeId = 2688, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2684, 2686}, children = {}, }
_typeInfoList[1076] = { parentId = 1, typeId = 2694, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1077] = { parentId = 1, typeId = 2695, nilable = true, orgTypeId = 2694 }
_typeInfoList[1078] = { parentId = 1, typeId = 2696, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1079] = { parentId = 1, typeId = 2697, nilable = true, orgTypeId = 2696 }
_typeInfoList[1080] = { parentId = 2492, typeId = 2698, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2694, 2696}, children = {}, }
_typeInfoList[1081] = { parentId = 1, typeId = 2704, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1082] = { parentId = 1, typeId = 2705, nilable = true, orgTypeId = 2704 }
_typeInfoList[1083] = { parentId = 1, typeId = 2706, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1084] = { parentId = 1, typeId = 2707, nilable = true, orgTypeId = 2706 }
_typeInfoList[1085] = { parentId = 2514, typeId = 2708, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2704, 2706}, children = {}, }
_typeInfoList[1086] = { parentId = 1, typeId = 2718, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1087] = { parentId = 1, typeId = 2719, nilable = true, orgTypeId = 2718 }
_typeInfoList[1088] = { parentId = 1, typeId = 2720, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1089] = { parentId = 1, typeId = 2721, nilable = true, orgTypeId = 2720 }
_typeInfoList[1090] = { parentId = 2534, typeId = 2722, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2718, 2720}, children = {}, }
_typeInfoList[1091] = { parentId = 1, typeId = 2732, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1092] = { parentId = 1, typeId = 2733, nilable = true, orgTypeId = 2732 }
_typeInfoList[1093] = { parentId = 1, typeId = 2734, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1094] = { parentId = 1, typeId = 2735, nilable = true, orgTypeId = 2734 }
_typeInfoList[1095] = { parentId = 2564, typeId = 2736, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2732, 2734}, children = {}, }
_typeInfoList[1096] = { parentId = 1, typeId = 2746, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1097] = { parentId = 1, typeId = 2747, nilable = true, orgTypeId = 2746 }
_typeInfoList[1098] = { parentId = 1, typeId = 2748, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1099] = { parentId = 1, typeId = 2749, nilable = true, orgTypeId = 2748 }
_typeInfoList[1100] = { parentId = 2600, typeId = 2750, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2746, 2748}, children = {}, }
_typeInfoList[1101] = { parentId = 1, typeId = 2764, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1102] = { parentId = 1, typeId = 2765, nilable = true, orgTypeId = 2764 }
_typeInfoList[1103] = { parentId = 1, typeId = 2766, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1104] = { parentId = 1, typeId = 2767, nilable = true, orgTypeId = 2766 }
_typeInfoList[1105] = { parentId = 2628, typeId = 2768, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2764, 2766}, children = {}, }
_typeInfoList[1106] = { parentId = 1, typeId = 2774, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1107] = { parentId = 1, typeId = 2775, nilable = true, orgTypeId = 2774 }
_typeInfoList[1108] = { parentId = 1, typeId = 2776, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1109] = { parentId = 1, typeId = 2777, nilable = true, orgTypeId = 2776 }
_typeInfoList[1110] = { parentId = 2648, typeId = 2778, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2774, 2776}, children = {}, }
_typeInfoList[1111] = { parentId = 1, typeId = 2786, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1112] = { parentId = 1, typeId = 2787, nilable = true, orgTypeId = 2786 }
_typeInfoList[1113] = { parentId = 1, typeId = 2788, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1114] = { parentId = 1, typeId = 2789, nilable = true, orgTypeId = 2788 }
_typeInfoList[1115] = { parentId = 2022, typeId = 2790, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2786, 2788}, children = {}, }
_typeInfoList[1116] = { parentId = 1, typeId = 2798, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1117] = { parentId = 1, typeId = 2799, nilable = true, orgTypeId = 2798 }
_typeInfoList[1118] = { parentId = 1, typeId = 2800, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1119] = { parentId = 1, typeId = 2801, nilable = true, orgTypeId = 2800 }
_typeInfoList[1120] = { parentId = 1974, typeId = 2802, baseId = 1, txt = 'getLiteral',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {2798, 2800}, children = {}, }
_typeInfoList[1121] = { parentId = 992, typeId = 2808, baseId = 1, txt = 'eval',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2406}, retTypeId = {26}, children = {}, }
_typeInfoList[1122] = { parentId = 992, typeId = 2810, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1123] = { parentId = 106, typeId = 2816, baseId = 1, txt = '_TypeInfo',
        staticFlag = false, accessMode = 'local', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2826}, }
_typeInfoList[1124] = { parentId = 106, typeId = 2817, nilable = true, orgTypeId = 2816 }
_typeInfoList[1125] = { parentId = 1, typeId = 2818, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1126] = { parentId = 1, typeId = 2819, nilable = true, orgTypeId = 2818 }
_typeInfoList[1127] = { parentId = 1, typeId = 2820, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1128] = { parentId = 1, typeId = 2821, nilable = true, orgTypeId = 2820 }
_typeInfoList[1129] = { parentId = 1, typeId = 2822, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1130] = { parentId = 1, typeId = 2823, nilable = true, orgTypeId = 2822 }
_typeInfoList[1131] = { parentId = 1, typeId = 2824, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {12}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1132] = { parentId = 1, typeId = 2825, nilable = true, orgTypeId = 2824 }
_typeInfoList[1133] = { parentId = 2816, typeId = 2826, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {12, 2818, 2820, 2822, 12, 12, 18, 12, 10, 10, 12, 2824, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1134] = { parentId = 106, typeId = 2828, baseId = 1, txt = '_ModuleInfo',
        staticFlag = false, accessMode = 'local', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {2846}, }
_typeInfoList[1135] = { parentId = 106, typeId = 2829, nilable = true, orgTypeId = 2828 }
_typeInfoList[1136] = { parentId = 1, typeId = 2830, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1137] = { parentId = 1, typeId = 2831, nilable = true, orgTypeId = 2830 }
_typeInfoList[1138] = { parentId = 1, typeId = 2832, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2830}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1139] = { parentId = 1, typeId = 2833, nilable = true, orgTypeId = 2832 }
_typeInfoList[1140] = { parentId = 1, typeId = 2834, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1141] = { parentId = 1, typeId = 2835, nilable = true, orgTypeId = 2834 }
_typeInfoList[1142] = { parentId = 1, typeId = 2836, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {12, 2834}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1143] = { parentId = 1, typeId = 2837, nilable = true, orgTypeId = 2836 }
_typeInfoList[1144] = { parentId = 1, typeId = 2838, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {2816}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1145] = { parentId = 1, typeId = 2839, nilable = true, orgTypeId = 2838 }
_typeInfoList[1146] = { parentId = 1, typeId = 2840, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1147] = { parentId = 1, typeId = 2841, nilable = true, orgTypeId = 2840 }
_typeInfoList[1148] = { parentId = 1, typeId = 2842, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 2840}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1149] = { parentId = 1, typeId = 2843, nilable = true, orgTypeId = 2842 }
_typeInfoList[1150] = { parentId = 1, typeId = 2844, baseId = 1, txt = 'Map',
        staticFlag = false, accessMode = 'pub', kind = 5, itemTypeId = {18, 6}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1151] = { parentId = 1, typeId = 2845, nilable = true, orgTypeId = 2844 }
_typeInfoList[1152] = { parentId = 2828, typeId = 2846, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {2832, 2836, 2838, 2842, 2844}, retTypeId = {}, children = {}, }
_typeInfoList[1153] = { parentId = 1026, typeId = 2852, baseId = 1, txt = 'registBuiltInScope',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1154] = { parentId = 1026, typeId = 3082, baseId = 1, txt = 'error',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {}, children = {}, }
_typeInfoList[1155] = { parentId = 1026, typeId = 3084, baseId = 1, txt = 'createNoneNode',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {216}, retTypeId = {814}, children = {}, }
_typeInfoList[1156] = { parentId = 1026, typeId = 3088, baseId = 1, txt = 'pushbackToken',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {221}, retTypeId = {}, children = {}, }
_typeInfoList[1157] = { parentId = 1, typeId = 3090, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1158] = { parentId = 1, typeId = 3091, nilable = true, orgTypeId = 3090 }
_typeInfoList[1159] = { parentId = 106, typeId = 3092, baseId = 1, txt = 'expandVal',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {3090, 6, 216}, retTypeId = {18}, children = {}, }
_typeInfoList[1160] = { parentId = 1, typeId = 3096, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1161] = { parentId = 1, typeId = 3097, nilable = true, orgTypeId = 3096 }
_typeInfoList[1162] = { parentId = 1026, typeId = 3098, baseId = 1, txt = 'newPushback',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {3096}, retTypeId = {}, children = {}, }
_typeInfoList[1163] = { parentId = 1026, typeId = 3100, baseId = 1, txt = 'getTokenNoErr',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[1164] = { parentId = 1026, typeId = 3114, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[1165] = { parentId = 1026, typeId = 3116, baseId = 1, txt = 'pushback',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1166] = { parentId = 1026, typeId = 3118, baseId = 1, txt = 'pushbackStr',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1167] = { parentId = 1026, typeId = 3124, baseId = 1, txt = 'checkSymbol',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {220}, children = {}, }
_typeInfoList[1168] = { parentId = 1026, typeId = 3126, baseId = 1, txt = 'getSymbolToken',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {220}, children = {}, }
_typeInfoList[1169] = { parentId = 1026, typeId = 3128, baseId = 1, txt = 'checkToken',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 18}, retTypeId = {220}, children = {}, }
_typeInfoList[1170] = { parentId = 1026, typeId = 3130, baseId = 1, txt = 'checkNextToken',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {220}, children = {}, }
_typeInfoList[1171] = { parentId = 1026, typeId = 3132, baseId = 1, txt = 'analyzeBlock',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 687}, retTypeId = {1308}, children = {}, }
_typeInfoList[1172] = { parentId = 1026, typeId = 3140, baseId = 1, txt = 'analyzeImport',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1173] = { parentId = 3140, typeId = 3158, baseId = 1, txt = 'registTypeInfo',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {2816}, retTypeId = {660}, children = {}, }
_typeInfoList[1174] = { parentId = 3140, typeId = 3178, baseId = 1, txt = 'registMember',
        staticFlag = true, accessMode = 'local', kind = 7, itemTypeId = {}, argTypeId = {12}, retTypeId = {}, children = {}, }
_typeInfoList[1175] = { parentId = 1026, typeId = 3186, baseId = 1, txt = 'analyzeIfUnwrap',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1176] = { parentId = 1026, typeId = 3190, baseId = 1, txt = 'analyzeIf',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1177] = { parentId = 1026, typeId = 3200, baseId = 1, txt = 'analyzeSwitch',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1408}, children = {}, }
_typeInfoList[1178] = { parentId = 1026, typeId = 3208, baseId = 1, txt = 'analyzeWhile',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1440}, children = {}, }
_typeInfoList[1179] = { parentId = 1026, typeId = 3212, baseId = 1, txt = 'analyzeRepeat',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1464}, children = {}, }
_typeInfoList[1180] = { parentId = 1026, typeId = 3216, baseId = 1, txt = 'analyzeFor',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1494}, children = {}, }
_typeInfoList[1181] = { parentId = 1026, typeId = 3220, baseId = 1, txt = 'analyzeApply',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1526}, children = {}, }
_typeInfoList[1182] = { parentId = 1026, typeId = 3228, baseId = 1, txt = 'analyzeForeach',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 10}, retTypeId = {814}, children = {}, }
_typeInfoList[1183] = { parentId = 1026, typeId = 3234, baseId = 1, txt = 'analyzeRefType',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18}, retTypeId = {1280}, children = {}, }
_typeInfoList[1184] = { parentId = 1, typeId = 3248, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1185] = { parentId = 1, typeId = 3249, nilable = true, orgTypeId = 3248 }
_typeInfoList[1186] = { parentId = 1026, typeId = 3250, baseId = 1, txt = 'analyzeDeclArgList',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 3248}, retTypeId = {220}, children = {}, }
_typeInfoList[1187] = { parentId = 106, typeId = 3254, baseId = 1, txt = 'ASTInfo',
        staticFlag = false, accessMode = 'pub', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3256, 3258, 3260}, }
_typeInfoList[1188] = { parentId = 106, typeId = 3255, nilable = true, orgTypeId = 3254 }
_typeInfoList[1189] = { parentId = 3254, typeId = 3256, baseId = 1, txt = 'get_node',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {814}, children = {}, }
_typeInfoList[1190] = { parentId = 3254, typeId = 3258, baseId = 1, txt = 'get_moduleTypeInfo',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {660}, children = {}, }
_typeInfoList[1191] = { parentId = 3254, typeId = 3260, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {814, 660}, retTypeId = {}, children = {}, }
_typeInfoList[1192] = { parentId = 1026, typeId = 3262, baseId = 1, txt = 'createAST',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {238, 10, 18}, retTypeId = {3254}, children = {}, }
_typeInfoList[1193] = { parentId = 1026, typeId = 3272, baseId = 1, txt = 'analyzeDeclMacro',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[1194] = { parentId = 1026, typeId = 3296, baseId = 1, txt = 'analyzeDeclProto',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[1195] = { parentId = 1026, typeId = 3298, baseId = 1, txt = 'analyzeDecl',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 220, 220}, retTypeId = {815}, children = {}, }
_typeInfoList[1196] = { parentId = 1026, typeId = 3300, baseId = 1, txt = 'analyzeDeclMember',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 10, 220}, retTypeId = {2280}, children = {}, }
_typeInfoList[1197] = { parentId = 1026, typeId = 3302, baseId = 1, txt = 'analyzeDeclMethod',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 220, 220, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[1198] = { parentId = 1026, typeId = 3304, baseId = 1, txt = 'analyzeDeclClass',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 220}, retTypeId = {816}, children = {}, }
_typeInfoList[1199] = { parentId = 1026, typeId = 3334, baseId = 1, txt = 'addMethod',
        staticFlag = false, accessMode = 'pri', kind = 8, itemTypeId = {}, argTypeId = {18, 814, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1200] = { parentId = 1026, typeId = 3336, baseId = 1, txt = 'analyzeDeclFunc',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {10, 18, 10, 221, 220, 221}, retTypeId = {814}, children = {}, }
_typeInfoList[1201] = { parentId = 1026, typeId = 3356, baseId = 1, txt = 'analyzeDeclVar',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {18, 18, 10, 220}, retTypeId = {814}, children = {}, }
_typeInfoList[1202] = { parentId = 1026, typeId = 3388, baseId = 1, txt = 'analyzeExpList',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {10}, retTypeId = {994}, children = {}, }
_typeInfoList[1203] = { parentId = 1026, typeId = 3398, baseId = 1, txt = 'analyzeListConst',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1204] = { parentId = 1026, typeId = 3412, baseId = 1, txt = 'analyzeMapConst',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {2564}, children = {}, }
_typeInfoList[1205] = { parentId = 1026, typeId = 3424, baseId = 1, txt = 'analyzeExpRefItem',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 814}, retTypeId = {814}, children = {}, }
_typeInfoList[1206] = { parentId = 1, typeId = 3428, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {660}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1207] = { parentId = 1, typeId = 3429, nilable = true, orgTypeId = 3428 }
_typeInfoList[1208] = { parentId = 1026, typeId = 3430, baseId = 1, txt = 'checkMatchValType',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {216, 660, 995, 3428}, retTypeId = {}, children = {}, }
_typeInfoList[1209] = { parentId = 106, typeId = 3432, baseId = 238, txt = 'MacroPaser',
        staticFlag = false, accessMode = 'local', kind = 6, itemTypeId = {}, argTypeId = {}, retTypeId = {}, children = {3438, 3440, 3442}, }
_typeInfoList[1210] = { parentId = 106, typeId = 3433, nilable = true, orgTypeId = 3432 }
_typeInfoList[1211] = { parentId = 1, typeId = 3436, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'pub', kind = 3, itemTypeId = {220}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1212] = { parentId = 1, typeId = 3437, nilable = true, orgTypeId = 3436 }
_typeInfoList[1213] = { parentId = 3432, typeId = 3438, baseId = 1, txt = '__init',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {3436, 18}, retTypeId = {}, children = {}, }
_typeInfoList[1214] = { parentId = 3432, typeId = 3440, baseId = 1, txt = 'getToken',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {221}, children = {}, }
_typeInfoList[1215] = { parentId = 3432, typeId = 3442, baseId = 1, txt = 'getStreamName',
        staticFlag = false, accessMode = 'pub', kind = 8, itemTypeId = {}, argTypeId = {}, retTypeId = {18}, children = {}, }
_typeInfoList[1216] = { parentId = 1026, typeId = 3446, baseId = 1, txt = 'evalMacro',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 660, 995}, retTypeId = {1948}, children = {}, }
_typeInfoList[1217] = { parentId = 1026, typeId = 3470, baseId = 1, txt = 'analyzeExpCont',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 814, 10}, retTypeId = {814}, children = {}, }
_typeInfoList[1218] = { parentId = 1026, typeId = 3474, baseId = 1, txt = 'analyzeExpSymbol',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 220, 18, 815, 10}, retTypeId = {814}, children = {}, }
_typeInfoList[1219] = { parentId = 1026, typeId = 3486, baseId = 1, txt = 'analyzeExpOp2',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220, 814, 12}, retTypeId = {814}, children = {}, }
_typeInfoList[1220] = { parentId = 1026, typeId = 3492, baseId = 1, txt = 'analyzeExpMacroStat',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {1974}, children = {}, }
_typeInfoList[1221] = { parentId = 1026, typeId = 3508, baseId = 1, txt = 'analyzeSuper',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1222] = { parentId = 1026, typeId = 3512, baseId = 1, txt = 'analyzeUnwrap',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1223] = { parentId = 1026, typeId = 3516, baseId = 1, txt = 'analyzeExpUnwrap',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {220}, retTypeId = {814}, children = {}, }
_typeInfoList[1224] = { parentId = 1026, typeId = 3520, baseId = 1, txt = 'analyzeExp',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {10, 12}, retTypeId = {814}, children = {}, }
_typeInfoList[1225] = { parentId = 1026, typeId = 3546, baseId = 1, txt = 'analyzeStatement',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {19}, retTypeId = {815}, children = {}, }
_typeInfoList[1226] = { parentId = 1, typeId = 3554, baseId = 1, txt = '',
        staticFlag = false, accessMode = 'local', kind = 3, itemTypeId = {814}, argTypeId = {}, retTypeId = {}, children = {}, }
_typeInfoList[1227] = { parentId = 1, typeId = 3555, nilable = true, orgTypeId = 3554 }
_typeInfoList[1228] = { parentId = 1026, typeId = 3556, baseId = 1, txt = 'analyzeStatementList',
        staticFlag = false, accessMode = 'local', kind = 8, itemTypeId = {}, argTypeId = {3554, 19}, retTypeId = {}, children = {}, }
----- meta -----
return moduleObj
