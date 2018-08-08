--lune/base/Ast.lns
local moduleObj = {}
local function _lune_nilacc( val, fieldName, access, ... )
   if not val then
      return nil
   end
   if fieldName then
      local field = val[ fieldName ]
      if not field then
         return nil
      end
      if access == "item" then
         local typeId = type( field )
         if typeId == "table" then
            return field[ ... ]
         elseif typeId == "string" then
            return string.byte( field, ... )
         end
      elseif access == "call" then
         return field( ... )
      elseif access == "callmtd" then
         return field( val, ... )
      end
      return field
   end
   if access == "item" then
      local typeId = type( val )
      if typeId == "table" then
         return val[ ... ]
      elseif typeId == "string" then
         return string.byte( val, ... )
      end
   elseif access == "call" then
      return val( ... )
   elseif access == "list" then
      local list, arg = ...
      if not list then
         return nil
      end
      return val( list, arg )
   end
   error( string.format( "illegal access -- %s", access ) )
end
function _lune_unwrap( val )
  if val == nil then
     _luneScript.error( 'unwrap val is nil' )
  end
  return val
end
function _lune_unwrapDefault( val, defval )
  if val == nil then
     return defval
  end
  return val
end

local Parser = require( 'lune.base.Parser' )

local Util = require( 'lune.base.Util' )

local rootTypeId = 1

moduleObj.rootTypeId = rootTypeId

local typeIdSeed = rootTypeId + 1

-- none

-- none

local typeInfoKind = {}

moduleObj.typeInfoKind = typeInfoKind

local sym2builtInTypeMap = {}

moduleObj.sym2builtInTypeMap = sym2builtInTypeMap

local builtInTypeIdSet = {}

moduleObj.builtInTypeIdSet = builtInTypeIdSet

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

local TypeInfoKindIF = 7

moduleObj.TypeInfoKindIF = TypeInfoKindIF

local TypeInfoKindFunc = 8

moduleObj.TypeInfoKindFunc = TypeInfoKindFunc

local TypeInfoKindMethod = 9

moduleObj.TypeInfoKindMethod = TypeInfoKindMethod

local TypeInfoKindNilable = 10

moduleObj.TypeInfoKindNilable = TypeInfoKindNilable

local function isBuiltin( typeId )

  return builtInTypeIdSet[typeId] ~= nil
end
moduleObj.isBuiltin = isBuiltin
local dummyList = {}

-- none


local SymbolKind = {}
moduleObj.SymbolKind = SymbolKind
function SymbolKind.getTxt( val )

  return _lune_unwrap( SymbolKind.val2txt[val])
end
function SymbolKind.new(  )
  local obj = {}
  setmetatable( obj, { __index = SymbolKind } )
  if obj.__init then
    obj:__init(  )
  end        
  return obj 
 end         
function SymbolKind:__init(  ) 
            
end
do
  SymbolKind.seed = 0
  SymbolKind.val2txt = {}
  SymbolKind.val2txt[SymbolKind.seed] = 'Typ'
  SymbolKind.Typ = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  SymbolKind.val2txt[SymbolKind.seed] = 'Mbr'
  SymbolKind.Mbr = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  SymbolKind.val2txt[SymbolKind.seed] = 'Mtd'
  SymbolKind.Mtd = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  SymbolKind.val2txt[SymbolKind.seed] = 'Fun'
  SymbolKind.Fun = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  SymbolKind.val2txt[SymbolKind.seed] = 'Var'
  SymbolKind.Var = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  SymbolKind.val2txt[SymbolKind.seed] = 'Arg'
  SymbolKind.Arg = SymbolKind.seed
  SymbolKind.seed = SymbolKind.seed + 1
  
  end

-- none

local SymbolInfo = {}
moduleObj.SymbolInfo = SymbolInfo
function SymbolInfo.new( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable )
  local obj = {}
  setmetatable( obj, { __index = SymbolInfo } )
  if obj.__init then obj:__init( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable ); end
return obj
end
function SymbolInfo:__init(kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable) 
  SymbolInfo.symbolIdSeed = SymbolInfo.symbolIdSeed + 1
  self.kind = kind
  self.canBeLeft = canBeLeft
  self.canBeRight = canBeRight
  self.symbolId = SymbolInfo.symbolIdSeed
  self.scope = scope
  self.accessMode = accessMode
  self.staticFlag = staticFlag
  self.name = name
  self.typeInfo = typeInfo
  self.mutable = mutable
end
-- none
function SymbolInfo:get_canBeLeft()
  return self.canBeLeft
end
function SymbolInfo:get_canBeRight()
  return self.canBeRight
end
function SymbolInfo:get_symbolId()
  return self.symbolId
end
function SymbolInfo:get_scope()
  return self.scope
end
function SymbolInfo:get_accessMode()
  return self.accessMode
end
function SymbolInfo:get_staticFlag()
  return self.staticFlag
end
function SymbolInfo:get_name()
  return self.name
end
function SymbolInfo:get_typeInfo()
  return self.typeInfo
end
function SymbolInfo:get_mutable()
  return self.mutable
end
function SymbolInfo:get_kind()
  return self.kind
end
do
  SymbolInfo.symbolIdSeed = 0
  end

local DataOwnerInfo = {}
moduleObj.DataOwnerInfo = DataOwnerInfo
function DataOwnerInfo.new( hasData, symbolInfo )
  local obj = {}
  setmetatable( obj, { __index = DataOwnerInfo } )
  if obj.__init then
    obj:__init( hasData, symbolInfo )
  end        
  return obj 
 end         
function DataOwnerInfo:__init( hasData, symbolInfo ) 
            
self.hasData = hasData
  self.symbolInfo = symbolInfo
  end
do
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
  self.parent = _lune_unwrapDefault( parent, self)
  self.symbol2TypeInfoMap = {}
  self.inheritList = inheritList
  self.classFlag = classFlag
  self.symbolId2DataOwnerInfo = {}
end
function Scope:set_ownerTypeInfo( owner )

  self.ownerTypeInfo = owner
end
function Scope:getTypeInfoChild( name )

  do
    local _exp = self.symbol2TypeInfoMap[name]
    if _exp ~= nil then
    
        return _exp:get_typeInfo()
      end
  end
  
  return nil
end
function Scope:setData( symbolInfo )

  self.symbolId2DataOwnerInfo[symbolInfo:get_symbolId()] = DataOwnerInfo.new(true, symbolInfo)
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
function Scope:get_inheritList()
  return self.inheritList
end
do
  end

local rootScope = Scope.new(nil, false, {})

moduleObj.rootScope = rootScope

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
    if _exp ~= nil then
    
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
function TypeInfo:isInheritFrom( other )

  return false
end
function TypeInfo:isSettableFrom( other )

  return false
end
function TypeInfo:get_abstructFlag(  )

  return false
end
function TypeInfo:getTxt(  )

  return ""
end
function TypeInfo:serialize( stream )

  return 
end
function TypeInfo:get_display_stirng(  )

  return ""
end
function TypeInfo:equals( typeInfo )

  return false
end
function TypeInfo:get_externalFlag(  )

  return false
end
function TypeInfo:get_interfaceList(  )

  return dummyList
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
function TypeInfo:get_mutable(  )

  return false
end
function TypeInfo:get_scope()
  return self.scope
end
do
  end

function Scope:getTypeInfoField( name, includeSelfFlag, fromScope )

  if self.classFlag then
    if includeSelfFlag then
      do
        local _exp = self.symbol2TypeInfoMap[name]
        if _exp ~= nil then
        
            local symbolInfo = _exp:canAccess( fromScope )
            
                if  nil == symbolInfo then
                  local _symbolInfo = symbolInfo
                  
                  return nil
                end
              
            return symbolInfo:get_typeInfo()
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
    if _exp ~= nil then
    
        local symbolInfo = _exp:canAccess( fromScope )
        
            if  nil == symbolInfo then
              local _symbolInfo = symbolInfo
              
              return nil
            end
          
        return symbolInfo:get_typeInfo()
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
  do
    local _exp = sym2builtInTypeMap[name]
    if _exp ~= nil then
    
        return _exp:get_typeInfo()
      end
  end
  
  return nil
end

function Scope:getSymbolTypeInfo( name, fromScope, moduleScope )

  local typeInfo = nil
  
  local validThisScope = false
  
  do
    local _exp = self.ownerTypeInfo
    if _exp ~= nil then
    
        if _exp:get_kind() == TypeInfoKindFunc or _exp:get_kind() == TypeInfoKindMethod or self == moduleScope or self == rootScope then
          validThisScope = true
        end
        if (_exp:get_kind() == TypeInfoKindIF or _exp:get_kind() == TypeInfoKindClass ) and name == "self" then
          validThisScope = true
        end
      else
    
        validThisScope = true
      end
  end
  
  if validThisScope then
    do
      local _exp = self.symbol2TypeInfoMap[name]
      if _exp ~= nil then
      
          return _exp:canAccess( fromScope )
        end
    end
    
  end
  if self.parent ~= self then
    return self.parent:getSymbolTypeInfo( name, fromScope, moduleScope )
  end
  return sym2builtInTypeMap[name]
end

function Scope:add( kind, canBeLeft, canBeRight, name, typeInfo, accessMode, staticFlag, mutable )

  self.symbol2TypeInfoMap[name] = SymbolInfo.new(kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, typeInfo, mutable)
end

function Scope:addVar( argFlag, canBeLeft, name, typeInfo, mutable )

  self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, "local", false, mutable )
end

function Scope:addMember( name, typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Mbr, true, true, name, typeInfo, accessMode, staticFlag, mutable )
end

function Scope:addMethod( typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Mtd, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable )
end

function Scope:addFunc( typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Fun, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable )
end

function Scope:addClass( name, typeInfo )

  self:add( SymbolKind.Typ, false, false, name, typeInfo, typeInfo:get_accessMode(), true, true )
end

local function dumpScopeSub( scope, prefix, readyIdSet )

  local scope = scope
  
      if  nil == scope then
        local _scope = scope
        
        return 
      else
        
          if readyIdSet[scope] then
            return 
          end
          readyIdSet[scope] = true
          if #prefix > 20 then
            Util.err( "illegal" )
          end
          do
            local __sorted = {}
            local __map = scope:get_symbol2TypeInfoMap()
            for __key in pairs( __map ) do
              table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, symbol in ipairs( __sorted ) do
              symbolInfo = __map[ symbol ]
              do
                Util.log( string.format( "scope: %s, %s, %s", prefix, scope, symbol) )
                do
                  local _exp = symbolInfo:get_typeInfo():get_scope()
                  if _exp ~= nil then
                  
                      dumpScopeSub( _exp, prefix .. "  ", readyIdSet )
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

moduleObj.rootTypeInfo = rootTypeInfo

function Scope:getNSTypeInfo(  )

  local scope = self
  
  while scope.ownerTypeInfo ~= rootTypeInfo do
    do
      local _exp = scope.ownerTypeInfo
      if _exp ~= nil then
      
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
      if _exp ~= nil then
      
          if _exp:get_kind() == TypeInfoKindClass or _exp:get_kind() == TypeInfoKindIF then
            return _exp
          end
        end
    end
    
    scope = scope.parent
  end
  return rootTypeInfo
end

function SymbolInfo:canAccess( fromScope )

  local typeInfo = self:get_typeInfo()
  
  if self.scope == fromScope then
    return self
  end
  do
    local _switchExp = self:get_accessMode()
    if _switchExp == "pub" or _switchExp == "global" then
      return self
    elseif _switchExp == "pro" then
      local nsClass = self.scope:getClassTypeInfo(  )
      
      local fromClass = fromScope:getClassTypeInfo(  )
      
      if fromClass:isInheritFrom( nsClass ) then
        return self
      end
      return nil
    elseif _switchExp == "local" then
      return self
    elseif _switchExp == "pri" then
      local nsClass = self.scope:getClassTypeInfo(  )
      
      local fromClass = fromScope:getClassTypeInfo(  )
      
      if nsClass == fromClass then
        return self
      end
      return nil
    end
  end
  
  Util.err( string.format( "illegl accessmode -- %s, %s", self:get_accessMode(), self:get_name()) )
end

-- none

-- none

local NilableTypeInfo = {}
setmetatable( NilableTypeInfo, { __index = TypeInfo } )
moduleObj.NilableTypeInfo = NilableTypeInfo
function NilableTypeInfo:get_kind(  )

  return TypeInfoKindNilable
end
function NilableTypeInfo:get_nilable(  )

  return true
end
function NilableTypeInfo:getTxt(  )

  return self.orgTypeInfo:getTxt(  ) .. "!"
end
function NilableTypeInfo:serialize( stream )

  local parentId = self:getParentId(  )
  
  stream:write( string.format( '{ parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', parentId, self.typeId, self.orgTypeInfo:get_typeId()) )
  return nil
end
function NilableTypeInfo.new( orgTypeInfo, typeId )
  local obj = {}
  setmetatable( obj, { __index = NilableTypeInfo } )
  if obj.__init then
    obj:__init( orgTypeInfo, typeId )
  end        
  return obj 
 end         
function NilableTypeInfo:__init( orgTypeInfo, typeId ) 
            
self.orgTypeInfo = orgTypeInfo
  self.typeId = typeId
  end
function NilableTypeInfo:get_orgTypeInfo()
  return self.orgTypeInfo
end
function NilableTypeInfo:get_typeId()
  return self.typeId
end
function NilableTypeInfo:getParentId( ... )
   return self.orgTypeInfo:getParentId( ... )
end

function NilableTypeInfo:get_baseId( ... )
   return self.orgTypeInfo:get_baseId( ... )
end

function NilableTypeInfo:isInheritFrom( ... )
   return self.orgTypeInfo:isInheritFrom( ... )
end

function NilableTypeInfo:get_abstructFlag( ... )
   return self.orgTypeInfo:get_abstructFlag( ... )
end

function NilableTypeInfo:get_display_stirng( ... )
   return self.orgTypeInfo:get_display_stirng( ... )
end

function NilableTypeInfo:equals( ... )
   return self.orgTypeInfo:equals( ... )
end

function NilableTypeInfo:get_externalFlag( ... )
   return self.orgTypeInfo:get_externalFlag( ... )
end

function NilableTypeInfo:get_interfaceList( ... )
   return self.orgTypeInfo:get_interfaceList( ... )
end

function NilableTypeInfo:get_itemTypeInfoList( ... )
   return self.orgTypeInfo:get_itemTypeInfoList( ... )
end

function NilableTypeInfo:get_argTypeInfoList( ... )
   return self.orgTypeInfo:get_argTypeInfoList( ... )
end

function NilableTypeInfo:get_retTypeInfoList( ... )
   return self.orgTypeInfo:get_retTypeInfoList( ... )
end

function NilableTypeInfo:get_parentInfo( ... )
   return self.orgTypeInfo:get_parentInfo( ... )
end

function NilableTypeInfo:get_rawTxt( ... )
   return self.orgTypeInfo:get_rawTxt( ... )
end

function NilableTypeInfo:get_staticFlag( ... )
   return self.orgTypeInfo:get_staticFlag( ... )
end

function NilableTypeInfo:get_accessMode( ... )
   return self.orgTypeInfo:get_accessMode( ... )
end

function NilableTypeInfo:get_autoFlag( ... )
   return self.orgTypeInfo:get_autoFlag( ... )
end

function NilableTypeInfo:get_baseTypeInfo( ... )
   return self.orgTypeInfo:get_baseTypeInfo( ... )
end

function NilableTypeInfo:get_nilableTypeInfo( ... )
   return self.orgTypeInfo:get_nilableTypeInfo( ... )
end

function NilableTypeInfo:get_children( ... )
   return self.orgTypeInfo:get_children( ... )
end

function NilableTypeInfo:get_mutable( ... )
   return self.orgTypeInfo:get_mutable( ... )
end

function NilableTypeInfo:get_scope( ... )
   return self.orgTypeInfo:get_scope( ... )
end

do
  end

local NormalTypeInfo = {}
setmetatable( NormalTypeInfo, { __index = TypeInfo } )
moduleObj.NormalTypeInfo = NormalTypeInfo
function NormalTypeInfo.new( abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = NormalTypeInfo } )
  if obj.__init then obj:__init( abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList ); end
return obj
end
function NormalTypeInfo:__init(abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList) 
  TypeInfo.__init( self, scope)
  
  self.abstructFlag = abstructFlag
  self.baseTypeInfo = _lune_unwrapDefault( baseTypeInfo, rootTypeInfo)
  self.interfaceList = _lune_unwrapDefault( interfaceList, {})
  self.autoFlag = autoFlag
  self.externalFlag = externalFlag
  self.staticFlag = staticFlag
  self.accessMode = accessMode
  self.rawTxt = txt
  self.kind = kind
  self.itemTypeInfoList = _lune_unwrapDefault( itemTypeInfoList, {})
  self.argTypeInfoList = _lune_unwrapDefault( argTypeInfoList, {})
  self.retTypeInfoList = _lune_unwrapDefault( retTypeInfoList, {})
  self.orgTypeInfo = _lune_unwrapDefault( orgTypeInfo, rootTypeInfo)
  self.parentInfo = _lune_unwrapDefault( parentInfo, rootTypeInfo)
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
      if _switchExp == TypeInfoKindPrim or _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindArray or _switchExp == TypeInfoKindMap or _switchExp == TypeInfoKindClass or _switchExp == TypeInfoKindIF then
        hasNilable = true
      elseif _switchExp == TypeInfoKindFunc or _switchExp == TypeInfoKindMethod then
        hasNilable = true
      end
    end
    
    if hasNilable then
      if txt == "..." then
        self.nilableTypeInfo = self
      else 
        self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
        typeIdSeed = typeIdSeed + 1
      end
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
    return (_lune_unwrap( self.orgTypeInfo) ):getTxt(  ) .. "!"
  end
  if self.kind == TypeInfoKindArray then
    local _exp = self.itemTypeInfoList[1]
    
        if  nil == _exp then
          local __exp = _exp
          
          return "[@]"
        end
      
    return _exp:getTxt(  ) .. "[@]"
  end
  if self.kind == TypeInfoKindList then
    local _exp = self.itemTypeInfoList[1]
    
        if  nil == _exp then
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
function NormalTypeInfo:get_display_stirng(  )

  if self.kind == TypeInfoKindNilable then
    return (_lune_unwrap( self.orgTypeInfo) ):get_display_stirng(  ) .. "!"
  end
  if self.kind == TypeInfoKindFunc or self.kind == TypeInfoKindMethod then
    local txt = self:get_rawTxt() .. "("
    
    for index, argType in pairs( self.argTypeInfoList ) do
      if index ~= 1 then
        txt = txt .. ", "
      end
      txt = txt .. argType:get_display_stirng(  )
    end
    txt = txt .. ")"
    for index, retType in pairs( self.retTypeInfoList ) do
      if index == 1 then
        txt = txt .. ": "
      else 
        txt = txt .. ", "
      end
      txt = txt .. retType:get_display_stirng(  )
    end
    return txt
  end
  return self:getTxt(  )
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
  
  stream:write( txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "ifList = {", self.interfaceList ) .. serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. serializeTypeInfoList( "children = {", self.children, true ) .. "}\n" )
end
function NormalTypeInfo:equalsSub( typeInfo )

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
    Util.log( string.format( "%s, %s", self.itemTypeInfoList, typeInfo:get_itemTypeInfoList()) )
    Util.log( string.format( "%s, %s", self.retTypeInfoList, typeInfo:get_retTypeInfoList()) )
    Util.log( string.format( "%s, %s", self.orgTypeInfo, typeInfo:get_orgTypeInfo()) )
    return false
  end
  if self.itemTypeInfoList then
    if #self.itemTypeInfoList ~= #typeInfo:get_itemTypeInfoList() then
      return false
    end
    for index, item in pairs( self.itemTypeInfoList ) do
      if not item:equals( typeInfo:get_itemTypeInfoList()[index] ) then
        return false
      end
    end
  end
  if self.retTypeInfoList then
    if #self.retTypeInfoList ~= #typeInfo:get_retTypeInfoList() then
      return false
    end
    for index, item in pairs( self.retTypeInfoList ) do
      if not item:equals( typeInfo:get_retTypeInfoList()[index] ) then
        return false
      end
    end
  end
  if self.orgTypeInfo and not self.orgTypeInfo:equals( typeInfo:get_orgTypeInfo() ) then
    return false
  end
  return true
end
function NormalTypeInfo:equals( typeInfo )

  return self:equalsSub( typeInfo )
end
function NormalTypeInfo.create( abstructFlag, scope, baseInfo, interfaceList, parentInfo, staticFlag, kind, txt, itemTypeInfo, argTypeInfoList, retTypeInfoList )

  if kind == TypeInfoKindPrim then
    do
      local _exp = sym2builtInTypeMap[txt]
      if _exp ~= nil then
      
          return _exp:get_typeInfo()
        end
    end
    
    Util.err( string.format( "not found symbol -- %s", txt) )
  end
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(abstructFlag, scope, baseInfo, interfaceList, nil, false, true, staticFlag, "pub", txt, parentInfo, typeIdSeed, kind, itemTypeInfo, argTypeInfoList, retTypeInfoList)
  
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
function NormalTypeInfo:get_abstructFlag()
  return self.abstructFlag
end
function NormalTypeInfo:get_orgTypeInfo()
  return self.orgTypeInfo
end
function NormalTypeInfo:get_baseTypeInfo()
  return self.baseTypeInfo
end
function NormalTypeInfo:get_interfaceList()
  return self.interfaceList
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
do
  end

local typeInfoRoot = rootTypeInfo

moduleObj.typeInfoRoot = typeInfoRoot

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
      if _exp ~= nil then
      
          argTypeList = {_exp}
          retTypeList = {_exp}
        end
    end
    
  end
  local scope = nil
  
  do
    local _switchExp = kind
    if _switchExp == TypeInfoKindList or _switchExp == TypeInfoKindClass or _switchExp == TypeInfoKindIF or _switchExp == TypeInfoKindFunc or _switchExp == TypeInfoKindMethod or _switchExp == TypeInfoKindMacro then
      scope = Scope.new(rootScope, kind == TypeInfoKindClass or kind == TypeInfoKindIF or kind == TypeInfoKindList, {})
    end
  end
  
  local info = NormalTypeInfo.new(false, scope, nil, nil, nil, false, false, false, "pub", typeTxt, typeInfoRoot, typeId, kind, {}, argTypeList, retTypeList)
  
  if scope then
    rootScope:addClass( typeTxt, info )
  end
  typeInfoKind[idName] = info
  sym2builtInTypeMap[typeTxt] = SymbolInfo.new(SymbolKind.Typ, false, false, rootScope, "pub", false, typeTxt, info, false)
  if info:get_nilableTypeInfo() ~= rootTypeInfo then
    sym2builtInTypeMap[typeTxt .. "!"] = SymbolInfo.new(SymbolKind.Typ, false, kind == TypeInfoKindFunc, rootScope, "pub", false, typeTxt, info:get_nilableTypeInfo(), false)
    builtInTypeIdSet[info:get_nilableTypeInfo():get_typeId()] = info:get_nilableTypeInfo()
  end
  builtInTypeIdSet[info.typeId] = info
  return info
end

function NormalTypeInfo.createList( accessMode, parentInfo, itemTypeInfo )

  if not itemTypeInfo or #itemTypeInfo == 0 then
    Util.err( string.format( "illegal list type: %s", itemTypeInfo) )
  end
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "", typeInfoRoot, typeIdSeed, TypeInfoKindList, itemTypeInfo)
end

function NormalTypeInfo.createArray( accessMode, parentInfo, itemTypeInfo )

  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "", typeInfoRoot, typeIdSeed, TypeInfoKindArray, itemTypeInfo)
end

function NormalTypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo )

  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "Map", typeInfoRoot, typeIdSeed, TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end

function NormalTypeInfo.createClass( classFlag, abstructFlag, scope, baseInfo, interfaceList, parentInfo, externalFlag, accessMode, className )

  do
    local _exp = sym2builtInTypeMap[className]
    if _exp ~= nil then
    
        return _exp:get_typeInfo()
      end
  end
  
  if Parser.isLuaKeyword( className ) then
    Util.err( string.format( "This symbol can not use for a class or script file. -- %s", className) )
  end
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(abstructFlag, scope, baseInfo, interfaceList, nil, false, externalFlag, false, accessMode, className, parentInfo, typeIdSeed, classFlag and TypeInfoKindClass or TypeInfoKindIF)
  
  return info
end

function NormalTypeInfo.createFunc( abstructFlag, builtinFlag, scope, kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList )

  if not builtinFlag and Parser.isLuaKeyword( funcName ) then
    Util.err( string.format( "This symbol can not use for a function. -- %s", funcName) )
  end
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(abstructFlag, scope, nil, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, typeIdSeed, kind, {}, _lune_unwrapDefault( argTypeList, {}), _lune_unwrapDefault( retTypeInfoList, {}))
  
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

local builtinTypeStem_ = _lune_unwrap( builtinTypeStem:get_nilableTypeInfo())

moduleObj.builtinTypeStem_ = builtinTypeStem_

function NilableTypeInfo:isSettableFrom( other )

  if not other then
    return false
  end
  if self == builtinTypeStem_ then
    return true
  end
  if other == builtinTypeNil then
    return true
  end
  if self.typeId == other:get_typeId() then
    return true
  end
  if other:get_nilable() then
    return self:get_orgTypeInfo():isSettableFrom( other:get_orgTypeInfo() )
  end
  return self:get_orgTypeInfo():isSettableFrom( other )
end


-- none

function NormalTypeInfo:isInheritFrom( other )

  local otherTypeId = other:get_typeId()
  
  if self:get_typeId() == otherTypeId then
    return true
  end
  if (self:get_kind() ~= TypeInfoKindClass and self:get_kind() ~= TypeInfoKindIF ) or (other:get_kind() ~= TypeInfoKindClass and other:get_kind() ~= TypeInfoKindIF ) then
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
  
  for __index, ifType in pairs( self:get_interfaceList() ) do
    while ifType ~= rootTypeInfo do
      if otherTypeId == ifType:get_typeId() then
        return true
      end
      ifType = ifType:get_baseTypeInfo()
    end
    -- none
    
  end
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
  if self == builtinTypeForm and other:get_kind() == TypeInfoKindFunc then
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
      if other:get_nilable() then
        return self:get_orgTypeInfo():isSettableFrom( other:get_orgTypeInfo() )
      end
      return self:get_orgTypeInfo():isSettableFrom( other )
    end
    if (self:get_kind() == TypeInfoKindClass or self:get_kind() == TypeInfoKindIF ) and (other:get_kind() == TypeInfoKindClass or other:get_kind() == TypeInfoKindIF ) then
      return other:isInheritFrom( self )
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
      if not (_lune_unwrap( self:get_itemTypeInfoList()[1]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[1]) ) then
        return false
      end
      
      return true
    elseif _switchExp == TypeInfoKindMap then
      if other:get_itemTypeInfoList()[1] == builtinTypeNone and other:get_itemTypeInfoList()[2] == builtinTypeNone then
        return true
      end
      if not (_lune_unwrap( self:get_itemTypeInfoList()[1]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[1]) ) then
        return false
      end
      
      if not (_lune_unwrap( self:get_itemTypeInfoList()[2]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[2]) ) then
        return false
      end
      
      return true
    elseif _switchExp == TypeInfoKindClass or _switchExp == TypeInfoKindIF then
      return other:isInheritFrom( self )
    elseif _switchExp == TypeInfoKindFunc then
      if self == builtinTypeForm then
        return true
      end
      return false
    elseif _switchExp == TypeInfoKindMethod then
      return false
    elseif _switchExp == TypeInfoKindNilable then
      return self:get_orgTypeInfo():isSettableFrom( other:get_orgTypeInfo() )
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
do
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
function Node:canBeLeft(  )

  return false
end
function Node:canBeRight(  )

  return false
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
do
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
do
  end

-- none

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
do
  end

local MacroValInfo = {}
moduleObj.MacroValInfo = MacroValInfo
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
do
  end

local MacroInfo = {}
moduleObj.MacroInfo = MacroInfo
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
do
  end

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

  return _lune_unwrap( nodeKind2NameMap[kind])
end
moduleObj.getNodeKindName = getNodeKindName

-- none

function Filter:processNone( node, ... )

end

-- none

-- none

local nodeKindNone = regKind( [[None]] )

moduleObj.nodeKindNone = nodeKindNone

local NoneNode = {}
setmetatable( NoneNode, { __index = Node } )
moduleObj.NoneNode = NoneNode
function NoneNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processNone( self, table.unpack( argList ) )
end
function NoneNode:canBeRight(  )

  return false
end
function NoneNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processSubfile( node, ... )

end

-- none

-- none

local nodeKindSubfile = regKind( [[Subfile]] )

moduleObj.nodeKindSubfile = nodeKindSubfile

local SubfileNode = {}
setmetatable( SubfileNode, { __index = Node } )
moduleObj.SubfileNode = SubfileNode
function SubfileNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processSubfile( self, table.unpack( argList ) )
end
function SubfileNode:canBeRight(  )

  return false
end
function SubfileNode:canBeLeft(  )

  return false
end
function SubfileNode.new( pos, builtinTypeList )
  local obj = {}
  setmetatable( obj, { __index = SubfileNode } )
  if obj.__init then obj:__init( pos, builtinTypeList ); end
return obj
end
function SubfileNode:__init(pos, builtinTypeList) 
  Node.__init( self, nodeKindSubfile, pos, builtinTypeList)
  
  -- none
  
  -- none
  
end
do
  end


-- none

function Filter:processImport( node, ... )

end

-- none

-- none

local nodeKindImport = regKind( [[Import]] )

moduleObj.nodeKindImport = nodeKindImport

local ImportNode = {}
setmetatable( ImportNode, { __index = Node } )
moduleObj.ImportNode = ImportNode
function ImportNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processImport( self, table.unpack( argList ) )
end
function ImportNode:canBeRight(  )

  return false
end
function ImportNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processRoot( node, ... )

end

-- none

-- none

local nodeKindRoot = regKind( [[Root]] )

moduleObj.nodeKindRoot = nodeKindRoot

local RootNode = {}
setmetatable( RootNode, { __index = Node } )
moduleObj.RootNode = RootNode
function RootNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processRoot( self, table.unpack( argList ) )
end
function RootNode:canBeRight(  )

  return false
end
function RootNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processRefType( node, ... )

end

-- none

-- none

local nodeKindRefType = regKind( [[RefType]] )

moduleObj.nodeKindRefType = nodeKindRefType

local RefTypeNode = {}
setmetatable( RefTypeNode, { __index = Node } )
moduleObj.RefTypeNode = RefTypeNode
function RefTypeNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processRefType( self, table.unpack( argList ) )
end
function RefTypeNode:canBeRight(  )

  return false
end
function RefTypeNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processBlock( node, ... )

end

-- none

-- none

local nodeKindBlock = regKind( [[Block]] )

moduleObj.nodeKindBlock = nodeKindBlock

local BlockNode = {}
setmetatable( BlockNode, { __index = Node } )
moduleObj.BlockNode = BlockNode
function BlockNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processBlock( self, table.unpack( argList ) )
end
function BlockNode:canBeRight(  )

  return false
end
function BlockNode:canBeLeft(  )

  return false
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
do
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
do
  end

-- none

function Filter:processIf( node, ... )

end

-- none

-- none

local nodeKindIf = regKind( [[If]] )

moduleObj.nodeKindIf = nodeKindIf

local IfNode = {}
setmetatable( IfNode, { __index = Node } )
moduleObj.IfNode = IfNode
function IfNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processIf( self, table.unpack( argList ) )
end
function IfNode:canBeRight(  )

  return false
end
function IfNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpList( node, ... )

end

-- none

-- none

local nodeKindExpList = regKind( [[ExpList]] )

moduleObj.nodeKindExpList = nodeKindExpList

local ExpListNode = {}
setmetatable( ExpListNode, { __index = Node } )
moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpList( self, table.unpack( argList ) )
end
function ExpListNode:canBeRight(  )

  return true
end
function ExpListNode:canBeLeft(  )

  return true
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
do
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
do
  end

-- none

function Filter:processSwitch( node, ... )

end

-- none

-- none

local nodeKindSwitch = regKind( [[Switch]] )

moduleObj.nodeKindSwitch = nodeKindSwitch

local SwitchNode = {}
setmetatable( SwitchNode, { __index = Node } )
moduleObj.SwitchNode = SwitchNode
function SwitchNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processSwitch( self, table.unpack( argList ) )
end
function SwitchNode:canBeRight(  )

  return false
end
function SwitchNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processWhile( node, ... )

end

-- none

-- none

local nodeKindWhile = regKind( [[While]] )

moduleObj.nodeKindWhile = nodeKindWhile

local WhileNode = {}
setmetatable( WhileNode, { __index = Node } )
moduleObj.WhileNode = WhileNode
function WhileNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processWhile( self, table.unpack( argList ) )
end
function WhileNode:canBeRight(  )

  return false
end
function WhileNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processRepeat( node, ... )

end

-- none

-- none

local nodeKindRepeat = regKind( [[Repeat]] )

moduleObj.nodeKindRepeat = nodeKindRepeat

local RepeatNode = {}
setmetatable( RepeatNode, { __index = Node } )
moduleObj.RepeatNode = RepeatNode
function RepeatNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processRepeat( self, table.unpack( argList ) )
end
function RepeatNode:canBeRight(  )

  return false
end
function RepeatNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processFor( node, ... )

end

-- none

-- none

local nodeKindFor = regKind( [[For]] )

moduleObj.nodeKindFor = nodeKindFor

local ForNode = {}
setmetatable( ForNode, { __index = Node } )
moduleObj.ForNode = ForNode
function ForNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processFor( self, table.unpack( argList ) )
end
function ForNode:canBeRight(  )

  return false
end
function ForNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processApply( node, ... )

end

-- none

-- none

local nodeKindApply = regKind( [[Apply]] )

moduleObj.nodeKindApply = nodeKindApply

local ApplyNode = {}
setmetatable( ApplyNode, { __index = Node } )
moduleObj.ApplyNode = ApplyNode
function ApplyNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processApply( self, table.unpack( argList ) )
end
function ApplyNode:canBeRight(  )

  return false
end
function ApplyNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processForeach( node, ... )

end

-- none

-- none

local nodeKindForeach = regKind( [[Foreach]] )

moduleObj.nodeKindForeach = nodeKindForeach

local ForeachNode = {}
setmetatable( ForeachNode, { __index = Node } )
moduleObj.ForeachNode = ForeachNode
function ForeachNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processForeach( self, table.unpack( argList ) )
end
function ForeachNode:canBeRight(  )

  return false
end
function ForeachNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processForsort( node, ... )

end

-- none

-- none

local nodeKindForsort = regKind( [[Forsort]] )

moduleObj.nodeKindForsort = nodeKindForsort

local ForsortNode = {}
setmetatable( ForsortNode, { __index = Node } )
moduleObj.ForsortNode = ForsortNode
function ForsortNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processForsort( self, table.unpack( argList ) )
end
function ForsortNode:canBeRight(  )

  return false
end
function ForsortNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processReturn( node, ... )

end

-- none

-- none

local nodeKindReturn = regKind( [[Return]] )

moduleObj.nodeKindReturn = nodeKindReturn

local ReturnNode = {}
setmetatable( ReturnNode, { __index = Node } )
moduleObj.ReturnNode = ReturnNode
function ReturnNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processReturn( self, table.unpack( argList ) )
end
function ReturnNode:canBeRight(  )

  return false
end
function ReturnNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processBreak( node, ... )

end

-- none

-- none

local nodeKindBreak = regKind( [[Break]] )

moduleObj.nodeKindBreak = nodeKindBreak

local BreakNode = {}
setmetatable( BreakNode, { __index = Node } )
moduleObj.BreakNode = BreakNode
function BreakNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processBreak( self, table.unpack( argList ) )
end
function BreakNode:canBeRight(  )

  return false
end
function BreakNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpNew( node, ... )

end

-- none

-- none

local nodeKindExpNew = regKind( [[ExpNew]] )

moduleObj.nodeKindExpNew = nodeKindExpNew

local ExpNewNode = {}
setmetatable( ExpNewNode, { __index = Node } )
moduleObj.ExpNewNode = ExpNewNode
function ExpNewNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpNew( self, table.unpack( argList ) )
end
function ExpNewNode:canBeRight(  )

  return true
end
function ExpNewNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpUnwrap( node, ... )

end

-- none

-- none

local nodeKindExpUnwrap = regKind( [[ExpUnwrap]] )

moduleObj.nodeKindExpUnwrap = nodeKindExpUnwrap

local ExpUnwrapNode = {}
setmetatable( ExpUnwrapNode, { __index = Node } )
moduleObj.ExpUnwrapNode = ExpUnwrapNode
function ExpUnwrapNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpUnwrap( self, table.unpack( argList ) )
end
function ExpUnwrapNode:canBeRight(  )

  return true
end
function ExpUnwrapNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpRef( node, ... )

end

-- none

-- none

local nodeKindExpRef = regKind( [[ExpRef]] )

moduleObj.nodeKindExpRef = nodeKindExpRef

local ExpRefNode = {}
setmetatable( ExpRefNode, { __index = Node } )
moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpRef( self, table.unpack( argList ) )
end
function ExpRefNode.new( pos, builtinTypeList, token, symbolInfo )
  local obj = {}
  setmetatable( obj, { __index = ExpRefNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, token, symbolInfo ); end
return obj
end
function ExpRefNode:__init(pos, builtinTypeList, token, symbolInfo) 
  Node.__init( self, nodeKindExpRef, pos, builtinTypeList)
  
  -- none
  
  self.token = token
  self.symbolInfo = symbolInfo
  -- none
  
end
function ExpRefNode:get_token()
  return self.token
end
function ExpRefNode:get_symbolInfo()
  return self.symbolInfo
end
do
  end


function ExpRefNode:canBeLeft(  )

  return self:get_symbolInfo():get_canBeLeft()
end

function ExpRefNode:canBeRight(  )

  return self:get_symbolInfo():get_canBeRight()
end

-- none

function Filter:processExpOp2( node, ... )

end

-- none

-- none

local nodeKindExpOp2 = regKind( [[ExpOp2]] )

moduleObj.nodeKindExpOp2 = nodeKindExpOp2

local ExpOp2Node = {}
setmetatable( ExpOp2Node, { __index = Node } )
moduleObj.ExpOp2Node = ExpOp2Node
function ExpOp2Node:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpOp2( self, table.unpack( argList ) )
end
function ExpOp2Node:canBeRight(  )

  return true
end
function ExpOp2Node:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processUnwrapSet( node, ... )

end

-- none

-- none

local nodeKindUnwrapSet = regKind( [[UnwrapSet]] )

moduleObj.nodeKindUnwrapSet = nodeKindUnwrapSet

local UnwrapSetNode = {}
setmetatable( UnwrapSetNode, { __index = Node } )
moduleObj.UnwrapSetNode = UnwrapSetNode
function UnwrapSetNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processUnwrapSet( self, table.unpack( argList ) )
end
function UnwrapSetNode:canBeRight(  )

  return false
end
function UnwrapSetNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processIfUnwrap( node, ... )

end

-- none

-- none

local nodeKindIfUnwrap = regKind( [[IfUnwrap]] )

moduleObj.nodeKindIfUnwrap = nodeKindIfUnwrap

local IfUnwrapNode = {}
setmetatable( IfUnwrapNode, { __index = Node } )
moduleObj.IfUnwrapNode = IfUnwrapNode
function IfUnwrapNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processIfUnwrap( self, table.unpack( argList ) )
end
function IfUnwrapNode:canBeRight(  )

  return false
end
function IfUnwrapNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpCast( node, ... )

end

-- none

-- none

local nodeKindExpCast = regKind( [[ExpCast]] )

moduleObj.nodeKindExpCast = nodeKindExpCast

local ExpCastNode = {}
setmetatable( ExpCastNode, { __index = Node } )
moduleObj.ExpCastNode = ExpCastNode
function ExpCastNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpCast( self, table.unpack( argList ) )
end
function ExpCastNode:canBeRight(  )

  return true
end
function ExpCastNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpOp1( node, ... )

end

-- none

-- none

local nodeKindExpOp1 = regKind( [[ExpOp1]] )

moduleObj.nodeKindExpOp1 = nodeKindExpOp1

local ExpOp1Node = {}
setmetatable( ExpOp1Node, { __index = Node } )
moduleObj.ExpOp1Node = ExpOp1Node
function ExpOp1Node:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpOp1( self, table.unpack( argList ) )
end
function ExpOp1Node:canBeRight(  )

  return true
end
function ExpOp1Node:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpRefItem( node, ... )

end

-- none

-- none

local nodeKindExpRefItem = regKind( [[ExpRefItem]] )

moduleObj.nodeKindExpRefItem = nodeKindExpRefItem

local ExpRefItemNode = {}
setmetatable( ExpRefItemNode, { __index = Node } )
moduleObj.ExpRefItemNode = ExpRefItemNode
function ExpRefItemNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpRefItem( self, table.unpack( argList ) )
end
function ExpRefItemNode:canBeRight(  )

  return true
end
function ExpRefItemNode:canBeLeft(  )

  return true
end
function ExpRefItemNode.new( pos, builtinTypeList, val, nilAccess, index )
  local obj = {}
  setmetatable( obj, { __index = ExpRefItemNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, val, nilAccess, index ); end
return obj
end
function ExpRefItemNode:__init(pos, builtinTypeList, val, nilAccess, index) 
  Node.__init( self, nodeKindExpRefItem, pos, builtinTypeList)
  
  -- none
  
  self.val = val
  self.nilAccess = nilAccess
  self.index = index
  -- none
  
end
function ExpRefItemNode:get_val()
  return self.val
end
function ExpRefItemNode:get_nilAccess()
  return self.nilAccess
end
function ExpRefItemNode:get_index()
  return self.index
end
do
  end


-- none

function Filter:processExpCall( node, ... )

end

-- none

-- none

local nodeKindExpCall = regKind( [[ExpCall]] )

moduleObj.nodeKindExpCall = nodeKindExpCall

local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpCall( self, table.unpack( argList ) )
end
function ExpCallNode:canBeRight(  )

  return true
end
function ExpCallNode:canBeLeft(  )

  return false
end
function ExpCallNode.new( pos, builtinTypeList, func, nilAccess, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, func, nilAccess, argList ); end
return obj
end
function ExpCallNode:__init(pos, builtinTypeList, func, nilAccess, argList) 
  Node.__init( self, nodeKindExpCall, pos, builtinTypeList)
  
  -- none
  
  self.func = func
  self.nilAccess = nilAccess
  self.argList = argList
  -- none
  
end
function ExpCallNode:get_func()
  return self.func
end
function ExpCallNode:get_nilAccess()
  return self.nilAccess
end
function ExpCallNode:get_argList()
  return self.argList
end
do
  end


-- none

function Filter:processExpDDD( node, ... )

end

-- none

-- none

local nodeKindExpDDD = regKind( [[ExpDDD]] )

moduleObj.nodeKindExpDDD = nodeKindExpDDD

local ExpDDDNode = {}
setmetatable( ExpDDDNode, { __index = Node } )
moduleObj.ExpDDDNode = ExpDDDNode
function ExpDDDNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpDDD( self, table.unpack( argList ) )
end
function ExpDDDNode:canBeRight(  )

  return true
end
function ExpDDDNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpParen( node, ... )

end

-- none

-- none

local nodeKindExpParen = regKind( [[ExpParen]] )

moduleObj.nodeKindExpParen = nodeKindExpParen

local ExpParenNode = {}
setmetatable( ExpParenNode, { __index = Node } )
moduleObj.ExpParenNode = ExpParenNode
function ExpParenNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpParen( self, table.unpack( argList ) )
end
function ExpParenNode:canBeRight(  )

  return true
end
function ExpParenNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpMacroExp( node, ... )

end

-- none

-- none

local nodeKindExpMacroExp = regKind( [[ExpMacroExp]] )

moduleObj.nodeKindExpMacroExp = nodeKindExpMacroExp

local ExpMacroExpNode = {}
setmetatable( ExpMacroExpNode, { __index = Node } )
moduleObj.ExpMacroExpNode = ExpMacroExpNode
function ExpMacroExpNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpMacroExp( self, table.unpack( argList ) )
end
function ExpMacroExpNode:canBeRight(  )

  return false
end
function ExpMacroExpNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpMacroStat( node, ... )

end

-- none

-- none

local nodeKindExpMacroStat = regKind( [[ExpMacroStat]] )

moduleObj.nodeKindExpMacroStat = nodeKindExpMacroStat

local ExpMacroStatNode = {}
setmetatable( ExpMacroStatNode, { __index = Node } )
moduleObj.ExpMacroStatNode = ExpMacroStatNode
function ExpMacroStatNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpMacroStat( self, table.unpack( argList ) )
end
function ExpMacroStatNode:canBeRight(  )

  return true
end
function ExpMacroStatNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processStmtExp( node, ... )

end

-- none

-- none

local nodeKindStmtExp = regKind( [[StmtExp]] )

moduleObj.nodeKindStmtExp = nodeKindStmtExp

local StmtExpNode = {}
setmetatable( StmtExpNode, { __index = Node } )
moduleObj.StmtExpNode = StmtExpNode
function StmtExpNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processStmtExp( self, table.unpack( argList ) )
end
function StmtExpNode:canBeRight(  )

  return true
end
function StmtExpNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processRefField( node, ... )

end

-- none

-- none

local nodeKindRefField = regKind( [[RefField]] )

moduleObj.nodeKindRefField = nodeKindRefField

local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processRefField( self, table.unpack( argList ) )
end
function RefFieldNode:canBeRight(  )

  return true
end
function RefFieldNode:canBeLeft(  )

  return true
end
function RefFieldNode.new( pos, builtinTypeList, field, nilAccess, prefix )
  local obj = {}
  setmetatable( obj, { __index = RefFieldNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, field, nilAccess, prefix ); end
return obj
end
function RefFieldNode:__init(pos, builtinTypeList, field, nilAccess, prefix) 
  Node.__init( self, nodeKindRefField, pos, builtinTypeList)
  
  -- none
  
  self.field = field
  self.nilAccess = nilAccess
  self.prefix = prefix
  -- none
  
end
function RefFieldNode:get_field()
  return self.field
end
function RefFieldNode:get_nilAccess()
  return self.nilAccess
end
function RefFieldNode:get_prefix()
  return self.prefix
end
do
  end


-- none

function Filter:processGetField( node, ... )

end

-- none

-- none

local nodeKindGetField = regKind( [[GetField]] )

moduleObj.nodeKindGetField = nodeKindGetField

local GetFieldNode = {}
setmetatable( GetFieldNode, { __index = Node } )
moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processGetField( self, table.unpack( argList ) )
end
function GetFieldNode:canBeRight(  )

  return true
end
function GetFieldNode:canBeLeft(  )

  return false
end
function GetFieldNode.new( pos, builtinTypeList, field, nilAccess, prefix, getterTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = GetFieldNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, field, nilAccess, prefix, getterTypeInfo ); end
return obj
end
function GetFieldNode:__init(pos, builtinTypeList, field, nilAccess, prefix, getterTypeInfo) 
  Node.__init( self, nodeKindGetField, pos, builtinTypeList)
  
  -- none
  
  self.field = field
  self.nilAccess = nilAccess
  self.prefix = prefix
  self.getterTypeInfo = getterTypeInfo
  -- none
  
end
function GetFieldNode:get_field()
  return self.field
end
function GetFieldNode:get_nilAccess()
  return self.nilAccess
end
function GetFieldNode:get_prefix()
  return self.prefix
end
function GetFieldNode:get_getterTypeInfo()
  return self.getterTypeInfo
end
do
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
do
  end

-- none

function Filter:processDeclVar( node, ... )

end

-- none

-- none

local nodeKindDeclVar = regKind( [[DeclVar]] )

moduleObj.nodeKindDeclVar = nodeKindDeclVar

local DeclVarNode = {}
setmetatable( DeclVarNode, { __index = Node } )
moduleObj.DeclVarNode = DeclVarNode
function DeclVarNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclVar( self, table.unpack( argList ) )
end
function DeclVarNode:canBeRight(  )

  return false
end
function DeclVarNode:canBeLeft(  )

  return false
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
do
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
do
  end

-- none

function Filter:processDeclFunc( node, ... )

end

-- none

-- none

local nodeKindDeclFunc = regKind( [[DeclFunc]] )

moduleObj.nodeKindDeclFunc = nodeKindDeclFunc

local DeclFuncNode = {}
setmetatable( DeclFuncNode, { __index = Node } )
moduleObj.DeclFuncNode = DeclFuncNode
function DeclFuncNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclFunc( self, table.unpack( argList ) )
end
function DeclFuncNode:canBeRight(  )

  return true
end
function DeclFuncNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processDeclMethod( node, ... )

end

-- none

-- none

local nodeKindDeclMethod = regKind( [[DeclMethod]] )

moduleObj.nodeKindDeclMethod = nodeKindDeclMethod

local DeclMethodNode = {}
setmetatable( DeclMethodNode, { __index = Node } )
moduleObj.DeclMethodNode = DeclMethodNode
function DeclMethodNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclMethod( self, table.unpack( argList ) )
end
function DeclMethodNode:canBeRight(  )

  return false
end
function DeclMethodNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processDeclConstr( node, ... )

end

-- none

-- none

local nodeKindDeclConstr = regKind( [[DeclConstr]] )

moduleObj.nodeKindDeclConstr = nodeKindDeclConstr

local DeclConstrNode = {}
setmetatable( DeclConstrNode, { __index = Node } )
moduleObj.DeclConstrNode = DeclConstrNode
function DeclConstrNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclConstr( self, table.unpack( argList ) )
end
function DeclConstrNode:canBeRight(  )

  return false
end
function DeclConstrNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processExpCallSuper( node, ... )

end

-- none

-- none

local nodeKindExpCallSuper = regKind( [[ExpCallSuper]] )

moduleObj.nodeKindExpCallSuper = nodeKindExpCallSuper

local ExpCallSuperNode = {}
setmetatable( ExpCallSuperNode, { __index = Node } )
moduleObj.ExpCallSuperNode = ExpCallSuperNode
function ExpCallSuperNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpCallSuper( self, table.unpack( argList ) )
end
function ExpCallSuperNode:canBeRight(  )

  return false
end
function ExpCallSuperNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processDeclMember( node, ... )

end

-- none

-- none

local nodeKindDeclMember = regKind( [[DeclMember]] )

moduleObj.nodeKindDeclMember = nodeKindDeclMember

local DeclMemberNode = {}
setmetatable( DeclMemberNode, { __index = Node } )
moduleObj.DeclMemberNode = DeclMemberNode
function DeclMemberNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclMember( self, table.unpack( argList ) )
end
function DeclMemberNode:canBeRight(  )

  return false
end
function DeclMemberNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processDeclArg( node, ... )

end

-- none

-- none

local nodeKindDeclArg = regKind( [[DeclArg]] )

moduleObj.nodeKindDeclArg = nodeKindDeclArg

local DeclArgNode = {}
setmetatable( DeclArgNode, { __index = Node } )
moduleObj.DeclArgNode = DeclArgNode
function DeclArgNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclArg( self, table.unpack( argList ) )
end
function DeclArgNode:canBeRight(  )

  return false
end
function DeclArgNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processDeclArgDDD( node, ... )

end

-- none

-- none

local nodeKindDeclArgDDD = regKind( [[DeclArgDDD]] )

moduleObj.nodeKindDeclArgDDD = nodeKindDeclArgDDD

local DeclArgDDDNode = {}
setmetatable( DeclArgDDDNode, { __index = Node } )
moduleObj.DeclArgDDDNode = DeclArgDDDNode
function DeclArgDDDNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclArgDDD( self, table.unpack( argList ) )
end
function DeclArgDDDNode:canBeRight(  )

  return false
end
function DeclArgDDDNode:canBeLeft(  )

  return false
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
do
  end


local AdvertiseInfo = {}
moduleObj.AdvertiseInfo = AdvertiseInfo
function AdvertiseInfo.new( member, prefix )
  local obj = {}
  setmetatable( obj, { __index = AdvertiseInfo } )
  if obj.__init then
    obj:__init( member, prefix )
  end        
  return obj 
 end         
function AdvertiseInfo:__init( member, prefix ) 
            
self.member = member
  self.prefix = prefix
  end
function AdvertiseInfo:get_member()
  return self.member
end
function AdvertiseInfo:get_prefix()
  return self.prefix
end
do
  end

-- none

-- none

function Filter:processDeclClass( node, ... )

end

-- none

-- none

local nodeKindDeclClass = regKind( [[DeclClass]] )

moduleObj.nodeKindDeclClass = nodeKindDeclClass

local DeclClassNode = {}
setmetatable( DeclClassNode, { __index = Node } )
moduleObj.DeclClassNode = DeclClassNode
function DeclClassNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclClass( self, table.unpack( argList ) )
end
function DeclClassNode:canBeRight(  )

  return false
end
function DeclClassNode:canBeLeft(  )

  return false
end
function DeclClassNode.new( pos, builtinTypeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, outerMethodSet )
  local obj = {}
  setmetatable( obj, { __index = DeclClassNode } )
  if obj.__init then obj:__init( pos, builtinTypeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, outerMethodSet ); end
return obj
end
function DeclClassNode:__init(pos, builtinTypeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, outerMethodSet) 
  Node.__init( self, nodeKindDeclClass, pos, builtinTypeList)
  
  -- none
  
  self.accessMode = accessMode
  self.name = name
  self.fieldList = fieldList
  self.moduleName = moduleName
  self.memberList = memberList
  self.scope = scope
  self.initStmtList = initStmtList
  self.advertiseList = advertiseList
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
function DeclClassNode:get_moduleName()
  return self.moduleName
end
function DeclClassNode:get_memberList()
  return self.memberList
end
function DeclClassNode:get_scope()
  return self.scope
end
function DeclClassNode:get_initStmtList()
  return self.initStmtList
end
function DeclClassNode:get_advertiseList()
  return self.advertiseList
end
function DeclClassNode:get_outerMethodSet()
  return self.outerMethodSet
end
do
  end


-- none

function Filter:processDeclMacro( node, ... )

end

-- none

-- none

local nodeKindDeclMacro = regKind( [[DeclMacro]] )

moduleObj.nodeKindDeclMacro = nodeKindDeclMacro

local DeclMacroNode = {}
setmetatable( DeclMacroNode, { __index = Node } )
moduleObj.DeclMacroNode = DeclMacroNode
function DeclMacroNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclMacro( self, table.unpack( argList ) )
end
function DeclMacroNode:canBeRight(  )

  return false
end
function DeclMacroNode:canBeLeft(  )

  return false
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
do
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
do
  end

-- none

function Filter:processLiteralNil( node, ... )

end

-- none

-- none

local nodeKindLiteralNil = regKind( [[LiteralNil]] )

moduleObj.nodeKindLiteralNil = nodeKindLiteralNil

local LiteralNilNode = {}
setmetatable( LiteralNilNode, { __index = Node } )
moduleObj.LiteralNilNode = LiteralNilNode
function LiteralNilNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralNil( self, table.unpack( argList ) )
end
function LiteralNilNode:canBeRight(  )

  return true
end
function LiteralNilNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralChar( node, ... )

end

-- none

-- none

local nodeKindLiteralChar = regKind( [[LiteralChar]] )

moduleObj.nodeKindLiteralChar = nodeKindLiteralChar

local LiteralCharNode = {}
setmetatable( LiteralCharNode, { __index = Node } )
moduleObj.LiteralCharNode = LiteralCharNode
function LiteralCharNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralChar( self, table.unpack( argList ) )
end
function LiteralCharNode:canBeRight(  )

  return true
end
function LiteralCharNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralInt( node, ... )

end

-- none

-- none

local nodeKindLiteralInt = regKind( [[LiteralInt]] )

moduleObj.nodeKindLiteralInt = nodeKindLiteralInt

local LiteralIntNode = {}
setmetatable( LiteralIntNode, { __index = Node } )
moduleObj.LiteralIntNode = LiteralIntNode
function LiteralIntNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralInt( self, table.unpack( argList ) )
end
function LiteralIntNode:canBeRight(  )

  return true
end
function LiteralIntNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralReal( node, ... )

end

-- none

-- none

local nodeKindLiteralReal = regKind( [[LiteralReal]] )

moduleObj.nodeKindLiteralReal = nodeKindLiteralReal

local LiteralRealNode = {}
setmetatable( LiteralRealNode, { __index = Node } )
moduleObj.LiteralRealNode = LiteralRealNode
function LiteralRealNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralReal( self, table.unpack( argList ) )
end
function LiteralRealNode:canBeRight(  )

  return true
end
function LiteralRealNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralArray( node, ... )

end

-- none

-- none

local nodeKindLiteralArray = regKind( [[LiteralArray]] )

moduleObj.nodeKindLiteralArray = nodeKindLiteralArray

local LiteralArrayNode = {}
setmetatable( LiteralArrayNode, { __index = Node } )
moduleObj.LiteralArrayNode = LiteralArrayNode
function LiteralArrayNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralArray( self, table.unpack( argList ) )
end
function LiteralArrayNode:canBeRight(  )

  return true
end
function LiteralArrayNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralList( node, ... )

end

-- none

-- none

local nodeKindLiteralList = regKind( [[LiteralList]] )

moduleObj.nodeKindLiteralList = nodeKindLiteralList

local LiteralListNode = {}
setmetatable( LiteralListNode, { __index = Node } )
moduleObj.LiteralListNode = LiteralListNode
function LiteralListNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralList( self, table.unpack( argList ) )
end
function LiteralListNode:canBeRight(  )

  return true
end
function LiteralListNode:canBeLeft(  )

  return false
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
do
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
do
  end

-- none

function Filter:processLiteralMap( node, ... )

end

-- none

-- none

local nodeKindLiteralMap = regKind( [[LiteralMap]] )

moduleObj.nodeKindLiteralMap = nodeKindLiteralMap

local LiteralMapNode = {}
setmetatable( LiteralMapNode, { __index = Node } )
moduleObj.LiteralMapNode = LiteralMapNode
function LiteralMapNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralMap( self, table.unpack( argList ) )
end
function LiteralMapNode:canBeRight(  )

  return true
end
function LiteralMapNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralString( node, ... )

end

-- none

-- none

local nodeKindLiteralString = regKind( [[LiteralString]] )

moduleObj.nodeKindLiteralString = nodeKindLiteralString

local LiteralStringNode = {}
setmetatable( LiteralStringNode, { __index = Node } )
moduleObj.LiteralStringNode = LiteralStringNode
function LiteralStringNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralString( self, table.unpack( argList ) )
end
function LiteralStringNode:canBeRight(  )

  return true
end
function LiteralStringNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralBool( node, ... )

end

-- none

-- none

local nodeKindLiteralBool = regKind( [[LiteralBool]] )

moduleObj.nodeKindLiteralBool = nodeKindLiteralBool

local LiteralBoolNode = {}
setmetatable( LiteralBoolNode, { __index = Node } )
moduleObj.LiteralBoolNode = LiteralBoolNode
function LiteralBoolNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralBool( self, table.unpack( argList ) )
end
function LiteralBoolNode:canBeRight(  )

  return true
end
function LiteralBoolNode:canBeLeft(  )

  return false
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
do
  end


-- none

function Filter:processLiteralSymbol( node, ... )

end

-- none

-- none

local nodeKindLiteralSymbol = regKind( [[LiteralSymbol]] )

moduleObj.nodeKindLiteralSymbol = nodeKindLiteralSymbol

local LiteralSymbolNode = {}
setmetatable( LiteralSymbolNode, { __index = Node } )
moduleObj.LiteralSymbolNode = LiteralSymbolNode
function LiteralSymbolNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processLiteralSymbol( self, table.unpack( argList ) )
end
function LiteralSymbolNode:canBeRight(  )

  return true
end
function LiteralSymbolNode:canBeLeft(  )

  return false
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
do
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
    if _exp ~= nil then
    
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
    if _exp ~= nil then
    
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

  local prefix = (_lune_unwrap( self.prefix:getLiteral(  )[1]) )
  
  if self.nilAccess then
    table.insert( prefix, "$." )
  else 
    table.insert( prefix, "." )
  end
  table.insert( prefix, self.field.txt )
  return {prefix}, {builtinTypeSymbol}
end

function ExpMacroStatNode:getLiteral(  )

  local txt = ""
  
  for __index, token in pairs( self.expStrList ) do
    txt = string.format( "%s%s", txt, token:getLiteral(  )[1])
  end
  return {txt}, {self:get_expType(  )}
end

return moduleObj
