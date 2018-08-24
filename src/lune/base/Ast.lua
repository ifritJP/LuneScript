--lune/base/Ast.lns
local _moduleObj = {}
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

_moduleObj.rootTypeId = rootTypeId

local typeIdSeed = _moduleObj.rootTypeId + 1

-- none

-- none

local typeInfoKind = {}

_moduleObj.typeInfoKind = typeInfoKind

local sym2builtInTypeMap = {}

_moduleObj.sym2builtInTypeMap = sym2builtInTypeMap

local builtInTypeIdSet = {}

_moduleObj.builtInTypeIdSet = builtInTypeIdSet

local TypeInfoKindRoot = 0

_moduleObj.TypeInfoKindRoot = TypeInfoKindRoot

local TypeInfoKindMacro = 1

_moduleObj.TypeInfoKindMacro = TypeInfoKindMacro

local TypeInfoKindPrim = 2

_moduleObj.TypeInfoKindPrim = TypeInfoKindPrim

local TypeInfoKindList = 3

_moduleObj.TypeInfoKindList = TypeInfoKindList

local TypeInfoKindArray = 4

_moduleObj.TypeInfoKindArray = TypeInfoKindArray

local TypeInfoKindMap = 5

_moduleObj.TypeInfoKindMap = TypeInfoKindMap

local TypeInfoKindClass = 6

_moduleObj.TypeInfoKindClass = TypeInfoKindClass

local TypeInfoKindIF = 7

_moduleObj.TypeInfoKindIF = TypeInfoKindIF

local TypeInfoKindFunc = 8

_moduleObj.TypeInfoKindFunc = TypeInfoKindFunc

local TypeInfoKindMethod = 9

_moduleObj.TypeInfoKindMethod = TypeInfoKindMethod

local TypeInfoKindNilable = 10

_moduleObj.TypeInfoKindNilable = TypeInfoKindNilable

local TypeInfoKindEnum = 11

_moduleObj.TypeInfoKindEnum = TypeInfoKindEnum

local function isBuiltin( typeId )

  return _moduleObj.builtInTypeIdSet[typeId] ~= nil
end
_moduleObj.isBuiltin = isBuiltin
local dummyList = {}

-- none


local SymbolKind = {}
_moduleObj.SymbolKind = SymbolKind
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
_moduleObj.SymbolInfo = SymbolInfo
function SymbolInfo.new( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag )
  local obj = {}
  setmetatable( obj, { __index = SymbolInfo } )
  if obj.__init then obj:__init( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag ); end
return obj
end
function SymbolInfo:__init(kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag) 
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
  self.hasValueFlag = hasValueFlag
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
function SymbolInfo:get_hasValueFlag()
  return self.hasValueFlag
end
function SymbolInfo:set_hasValueFlag( hasValueFlag )
  self.hasValueFlag = hasValueFlag
end
do
  SymbolInfo.symbolIdSeed = 0
  end

local DataOwnerInfo = {}
_moduleObj.DataOwnerInfo = DataOwnerInfo
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
_moduleObj.Scope = Scope
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
function Scope:getSymbolInfoChild( name )

  return self.symbol2TypeInfoMap[name]
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

_moduleObj.rootScope = rootScope

local TypeInfo = {}
_moduleObj.TypeInfo = TypeInfo
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

  return _moduleObj.rootTypeId
end
function TypeInfo:get_baseId(  )

  return _moduleObj.rootTypeId
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
function TypeInfo:serialize( stream, validChildrenSet )

  return 
end
function TypeInfo:get_display_stirng(  )

  return ""
end
function TypeInfo:equals( typeInfo )

  return self == typeInfo
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

  return _moduleObj.rootTypeId
end
function TypeInfo:get_kind(  )

  return _moduleObj.TypeInfoKindRoot
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

function Scope:getSymbolInfoField( name, includeSelfFlag, fromScope )

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
              
            return symbolInfo
          end
      end
      
    end
    if self.inheritList then
      for __index, scope in pairs( self.inheritList ) do
        local symbolInfo = scope:getSymbolInfoField( name, true, fromScope )
        
        if symbolInfo then
          return symbolInfo
        end
      end
    end
  end
  return nil
end

function Scope:getTypeInfoField( name, includeSelfFlag, fromScope )

  local symbolInfo = self:getSymbolInfoField( name, includeSelfFlag, fromScope )
  
  do
    local _exp = symbolInfo
    if _exp ~= nil then
    
        return _exp:get_typeInfo()
      end
  end
  
  return nil
end

function Scope:filterTypeInfoField( includeSelfFlag, fromScope, callback )

  if includeSelfFlag then
    for __index, symbolInfo in pairs( self.symbol2TypeInfoMap ) do
      if symbolInfo:canAccess( fromScope ) then
        if not callback( symbolInfo ) then
          return false
        end
      end
    end
  end
  if self.classFlag then
    if self.inheritList then
      for __index, scope in pairs( self.inheritList ) do
        if not scope:filterTypeInfoField( true, fromScope, callback ) then
          return false
        end
      end
    end
  end
  return true
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
    local _exp = _moduleObj.sym2builtInTypeMap[name]
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
    
        if _exp:get_kind() == _moduleObj.TypeInfoKindFunc or _exp:get_kind() == _moduleObj.TypeInfoKindMethod or self == moduleScope or self == _moduleObj.rootScope then
          validThisScope = true
        elseif (_exp:get_kind() == _moduleObj.TypeInfoKindIF or _exp:get_kind() == _moduleObj.TypeInfoKindClass ) and name == "self" then
          validThisScope = true
        elseif _exp:get_kind() == _moduleObj.TypeInfoKindEnum then
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
  return _moduleObj.sym2builtInTypeMap[name]
end

function Scope:filterSymbolTypeInfo( fromScope, moduleScope, callback )

  if self.classFlag then
    do
      local _exp = self.symbol2TypeInfoMap["self"]
      if _exp ~= nil then
      
          callback( _exp )
        end
    end
    
  end
  if moduleScope == fromScope or not self.classFlag then
    for __index, symbolInfo in pairs( self.symbol2TypeInfoMap ) do
      if not callback( symbolInfo ) then
        return 
      end
    end
  end
  if self.parent ~= self then
    self.parent:filterSymbolTypeInfo( fromScope, moduleScope, callback )
  end
end

function Scope:add( kind, canBeLeft, canBeRight, name, typeInfo, accessMode, staticFlag, mutable, hasValueFlag )

  local symbolInfo = SymbolInfo.new(kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag)
  
  self.symbol2TypeInfoMap[name] = symbolInfo
end

function Scope:addLocalVar( argFlag, canBeLeft, name, typeInfo, mutable )

  self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, "local", false, mutable, true )
end

function Scope:addStaticVar( argFlag, canBeLeft, name, typeInfo, mutable )

  self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, "local", true, mutable, true )
end

function Scope:addVar( accessMode, name, typeInfo, mutable )

  self:add( SymbolKind.Var, true, true, name, typeInfo, accessMode, false, mutable, true )
end

function Scope:addEnumVal( name, typeInfo )

  self:add( SymbolKind.Mbr, false, true, name, typeInfo, "pub", true, true, true )
end

function Scope:addEnum( accessMode, name, typeInfo )

  self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, true, true )
end

function Scope:addMember( name, typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Mbr, true, true, name, typeInfo, accessMode, staticFlag, mutable, false )
end

function Scope:addMethod( typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Mtd, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable, true )
end

function Scope:addFunc( typeInfo, accessMode, staticFlag, mutable )

  self:add( SymbolKind.Fun, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable, true )
end

function Scope:addClass( name, typeInfo )

  self:add( SymbolKind.Typ, false, false, name, typeInfo, typeInfo:get_accessMode(), true, true, true )
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

local rootTypeInfo = TypeInfo.new(_moduleObj.rootScope)

_moduleObj.rootTypeInfo = rootTypeInfo

function Scope:getNSTypeInfo(  )

  local scope = self
  
  while scope.ownerTypeInfo ~= _moduleObj.rootTypeInfo do
    do
      local _exp = scope.ownerTypeInfo
      if _exp ~= nil then
      
          return _exp
        end
    end
    
    scope = scope.parent
  end
  return _moduleObj.rootTypeInfo
end

function Scope:getClassTypeInfo(  )

  local scope = self
  
  while scope.ownerTypeInfo ~= _moduleObj.rootTypeInfo do
    do
      local _exp = scope.ownerTypeInfo
      if _exp ~= nil then
      
          if _exp:get_kind() == _moduleObj.TypeInfoKindClass or _exp:get_kind() == _moduleObj.TypeInfoKindIF then
            return _exp
          end
        end
    end
    
    scope = scope.parent
  end
  return _moduleObj.rootTypeInfo
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
_moduleObj.NilableTypeInfo = NilableTypeInfo
function NilableTypeInfo:get_kind(  )

  return _moduleObj.TypeInfoKindNilable
end
function NilableTypeInfo:get_nilable(  )

  return true
end
function NilableTypeInfo:getTxt(  )

  return self.orgTypeInfo:getTxt(  ) .. "!"
end
function NilableTypeInfo:get_display_stirng(  )

  return self.orgTypeInfo:get_display_stirng() .. "!"
end
function NilableTypeInfo:serialize( stream, validChildrenSet )

  local parentId = self:getParentId(  )
  
  stream:write( string.format( '{ parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', parentId, self.typeId, self.orgTypeInfo:get_typeId()) )
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

local EnumValInfo = {}
_moduleObj.EnumValInfo = EnumValInfo
function EnumValInfo.new( name, val )
  local obj = {}
  setmetatable( obj, { __index = EnumValInfo } )
  if obj.__init then
    obj:__init( name, val )
  end        
  return obj 
 end         
function EnumValInfo:__init( name, val ) 
            
self.name = name
  self.val = val
  end
function EnumValInfo:get_name()
  return self.name
end
function EnumValInfo:get_val()
  return self.val
end
do
  end

local EnumTypeInfo = {}
setmetatable( EnumTypeInfo, { __index = TypeInfo } )
_moduleObj.EnumTypeInfo = EnumTypeInfo
function EnumTypeInfo.new( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo )
  local obj = {}
  setmetatable( obj, { __index = EnumTypeInfo } )
  if obj.__init then obj:__init( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo ); end
return obj
end
function EnumTypeInfo:__init(scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo) 
  TypeInfo.__init( self, scope)
  
  self.externalFlag = externalFlag
  self.accessMode = accessMode
  self.rawTxt = txt
  self.parentInfo = _lune_unwrapDefault( parentInfo, _moduleObj.rootTypeInfo)
  self.typeId = typeId
  self.name2EnumValInfo = name2EnumValInfo
  self.valTypeInfo = valTypeInfo
  if self.parentInfo ~= _moduleObj.rootTypeInfo then
    table.insert( self.parentInfo:get_children(), self )
  end
  self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
  typeIdSeed = typeIdSeed + 1
  scope:set_ownerTypeInfo( self )
end
function EnumTypeInfo:get_kind(  )

  return _moduleObj.TypeInfoKindEnum
end
function EnumTypeInfo:getParentId(  )

  return self.parentInfo:get_typeId()
end
function EnumTypeInfo:getTxt(  )

  return self.rawTxt
end
function EnumTypeInfo:get_display_stirng(  )

  return self:getTxt(  )
end
function EnumTypeInfo:isSettableFrom( other )

  return self == other
end
function EnumTypeInfo:getEnumValInfo( name )

  return self.name2EnumValInfo[name]
end
function EnumTypeInfo:get_externalFlag()
  return self.externalFlag
end
function EnumTypeInfo:get_parentInfo()
  return self.parentInfo
end
function EnumTypeInfo:get_typeId()
  return self.typeId
end
function EnumTypeInfo:get_rawTxt()
  return self.rawTxt
end
function EnumTypeInfo:get_accessMode()
  return self.accessMode
end
function EnumTypeInfo:get_nilableTypeInfo()
  return self.nilableTypeInfo
end
function EnumTypeInfo:get_valTypeInfo()
  return self.valTypeInfo
end
function EnumTypeInfo:get_name2EnumValInfo()
  return self.name2EnumValInfo
end
do
  end

local NormalTypeInfo = {}
setmetatable( NormalTypeInfo, { __index = TypeInfo } )
_moduleObj.NormalTypeInfo = NormalTypeInfo
function NormalTypeInfo.new( abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList )
  local obj = {}
  setmetatable( obj, { __index = NormalTypeInfo } )
  if obj.__init then obj:__init( abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList ); end
return obj
end
function NormalTypeInfo:__init(abstructFlag, scope, baseTypeInfo, interfaceList, orgTypeInfo, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList) 
  TypeInfo.__init( self, scope)
  
  self.abstructFlag = abstructFlag
  self.baseTypeInfo = _lune_unwrapDefault( baseTypeInfo, _moduleObj.rootTypeInfo)
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
  self.orgTypeInfo = _lune_unwrapDefault( orgTypeInfo, _moduleObj.rootTypeInfo)
  self.parentInfo = _lune_unwrapDefault( parentInfo, _moduleObj.rootTypeInfo)
  self.children = {}
  self.typeId = typeId
  if kind == _moduleObj.TypeInfoKindRoot then
    self.nilable = false
  elseif txt == "nil" then
    self.nilable = true
    self.nilableTypeInfo = self
    self.orgTypeInfo = self
  elseif not orgTypeInfo then
    if self.parentInfo ~= _moduleObj.rootTypeInfo then
      table.insert( self.parentInfo:get_children(), self )
    end
    self.nilable = false
    local hasNilable = false
    
    do
      local _switchExp = (kind )
      if _switchExp == _moduleObj.TypeInfoKindPrim or _switchExp == _moduleObj.TypeInfoKindList or _switchExp == _moduleObj.TypeInfoKindArray or _switchExp == _moduleObj.TypeInfoKindMap or _switchExp == _moduleObj.TypeInfoKindClass or _switchExp == _moduleObj.TypeInfoKindIF then
        hasNilable = true
      elseif _switchExp == _moduleObj.TypeInfoKindFunc or _switchExp == _moduleObj.TypeInfoKindMethod then
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
      self.nilableTypeInfo = _moduleObj.rootTypeInfo
    end
    typeIdSeed = typeIdSeed + 1
  else 
    self.nilable = true
    self.nilableTypeInfo = _moduleObj.rootTypeInfo
  end
end
function NormalTypeInfo:getParentId(  )

  return self.parentInfo and self.parentInfo:get_typeId() or _moduleObj.rootTypeId
end
function NormalTypeInfo:get_baseId(  )

  return self.baseTypeInfo and self.baseTypeInfo:get_typeId() or _moduleObj.rootTypeId
end
function NormalTypeInfo:getTxt(  )

  if self.nilable and (self.nilableTypeInfo ~= self.orgTypeInfo ) then
    return (_lune_unwrap( self.orgTypeInfo) ):getTxt(  ) .. "!"
  end
  if self.kind == _moduleObj.TypeInfoKindArray then
    local _exp = self.itemTypeInfoList[1]
    
        if  nil == _exp then
          local __exp = _exp
          
          return "[@]"
        end
      
    return _exp:getTxt(  ) .. "[@]"
  end
  if self.kind == _moduleObj.TypeInfoKindList then
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

  if self.kind == _moduleObj.TypeInfoKindNilable then
    return (_lune_unwrap( self.orgTypeInfo) ):get_display_stirng(  ) .. "!"
  end
  if self.kind == _moduleObj.TypeInfoKindFunc or self.kind == _moduleObj.TypeInfoKindMethod then
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
function NormalTypeInfo:serialize( stream, validChildrenSet )

  if self.typeId == _moduleObj.rootTypeId then
    return 
  end
  local parentId = self:getParentId(  )
  
  if self.nilable then
    stream:write( string.format( '{ parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', parentId, self.typeId, self.orgTypeInfo:get_typeId()) )
    return 
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
  
  local children = self.children
  
  do
    local _exp = validChildrenSet
    if _exp ~= nil then
    
        children = {}
        for __index, child in pairs( self.children ) do
          if _exp[child:get_typeId()] then
            table.insert( children, child )
          end
        end
      end
  end
  
  stream:write( txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "ifList = {", self.interfaceList ) .. serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. serializeTypeInfoList( "children = {", children, true ) .. "}\n" )
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

  if kind == _moduleObj.TypeInfoKindPrim then
    do
      local _exp = _moduleObj.sym2builtInTypeMap[txt]
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
function NormalTypeInfo:get_externalFlag()
  return self.externalFlag
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

local typeInfoRoot = _moduleObj.rootTypeInfo

_moduleObj.typeInfoRoot = typeInfoRoot

typeIdSeed = typeIdSeed + 1
function NormalTypeInfo.createBuiltin( idName, typeTxt, kind, typeDDD )

  local typeId = typeIdSeed + 1
  
  if kind == _moduleObj.TypeInfoKindRoot then
    typeId = _moduleObj.rootTypeId
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
    if _switchExp == _moduleObj.TypeInfoKindList or _switchExp == _moduleObj.TypeInfoKindClass or _switchExp == _moduleObj.TypeInfoKindIF or _switchExp == _moduleObj.TypeInfoKindFunc or _switchExp == _moduleObj.TypeInfoKindMethod or _switchExp == _moduleObj.TypeInfoKindMacro then
      scope = Scope.new(_moduleObj.rootScope, kind == _moduleObj.TypeInfoKindClass or kind == _moduleObj.TypeInfoKindIF or kind == _moduleObj.TypeInfoKindList, {})
    end
  end
  
  local info = NormalTypeInfo.new(false, scope, nil, nil, nil, false, false, false, "pub", typeTxt, _moduleObj.typeInfoRoot, typeId, kind, {}, argTypeList, retTypeList)
  
  if scope then
    _moduleObj.rootScope:addClass( typeTxt, info )
  end
  _moduleObj.typeInfoKind[idName] = info
  _moduleObj.sym2builtInTypeMap[typeTxt] = SymbolInfo.new(SymbolKind.Typ, false, false, _moduleObj.rootScope, "pub", false, typeTxt, info, false, true)
  if info:get_nilableTypeInfo() ~= _moduleObj.rootTypeInfo then
    _moduleObj.sym2builtInTypeMap[typeTxt .. "!"] = SymbolInfo.new(SymbolKind.Typ, false, kind == _moduleObj.TypeInfoKindFunc, _moduleObj.rootScope, "pub", false, typeTxt, info:get_nilableTypeInfo(), false, true)
    _moduleObj.builtInTypeIdSet[info:get_nilableTypeInfo():get_typeId()] = info:get_nilableTypeInfo()
  end
  _moduleObj.builtInTypeIdSet[info.typeId] = info
  return info
end

function NormalTypeInfo.createList( accessMode, parentInfo, itemTypeInfo )

  if not itemTypeInfo or #itemTypeInfo == 0 then
    Util.err( string.format( "illegal list type: %s", itemTypeInfo) )
  end
  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "", _moduleObj.typeInfoRoot, typeIdSeed, _moduleObj.TypeInfoKindList, itemTypeInfo)
end

function NormalTypeInfo.createArray( accessMode, parentInfo, itemTypeInfo )

  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "", _moduleObj.typeInfoRoot, typeIdSeed, _moduleObj.TypeInfoKindArray, itemTypeInfo)
end

function NormalTypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo )

  typeIdSeed = typeIdSeed + 1
  return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "Map", _moduleObj.typeInfoRoot, typeIdSeed, _moduleObj.TypeInfoKindMap, {keyTypeInfo, valTypeInfo})
end

function NormalTypeInfo.createClass( classFlag, abstructFlag, scope, baseInfo, interfaceList, parentInfo, externalFlag, accessMode, className )

  do
    local _exp = _moduleObj.sym2builtInTypeMap[className]
    if _exp ~= nil then
    
        return _exp:get_typeInfo()
      end
  end
  
  if Parser.isLuaKeyword( className ) then
    Util.err( string.format( "This symbol can not use for a class or script file. -- %s", className) )
  end
  typeIdSeed = typeIdSeed + 1
  local info = NormalTypeInfo.new(abstructFlag, scope, baseInfo, interfaceList, nil, false, externalFlag, false, accessMode, className, parentInfo, typeIdSeed, classFlag and _moduleObj.TypeInfoKindClass or _moduleObj.TypeInfoKindIF)
  
  return info
end

function NormalTypeInfo.createEnum( scope, parentInfo, externalFlag, accessMode, enumName, valTypeInfo, name2EnumValInfo )

  if Parser.isLuaKeyword( enumName ) then
    Util.err( string.format( "This symbol can not use for a enum. -- %s", enumName) )
  end
  typeIdSeed = typeIdSeed + 1
  local info = EnumTypeInfo.new(scope, externalFlag, accessMode, enumName, parentInfo, typeIdSeed, valTypeInfo, name2EnumValInfo)
  
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

local builtinTypeNone = NormalTypeInfo.createBuiltin( "None", "", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeNone = builtinTypeNone

local builtinTypeStem = NormalTypeInfo.createBuiltin( "Stem", "stem", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeStem = builtinTypeStem

local builtinTypeNil = NormalTypeInfo.createBuiltin( "Nil", "nil", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeNil = builtinTypeNil

local builtinTypeDDD = NormalTypeInfo.createBuiltin( "DDD", "...", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeDDD = builtinTypeDDD

local builtinTypeBool = NormalTypeInfo.createBuiltin( "Bool", "bool", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeBool = builtinTypeBool

local builtinTypeInt = NormalTypeInfo.createBuiltin( "Int", "int", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeInt = builtinTypeInt

local builtinTypeReal = NormalTypeInfo.createBuiltin( "Real", "real", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeReal = builtinTypeReal

local builtinTypeChar = NormalTypeInfo.createBuiltin( "char", "char", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeChar = builtinTypeChar

local builtinTypeString = NormalTypeInfo.createBuiltin( "String", "str", _moduleObj.TypeInfoKindClass )

_moduleObj.builtinTypeString = builtinTypeString

local builtinTypeMap = NormalTypeInfo.createBuiltin( "Map", "Map", _moduleObj.TypeInfoKindMap )

_moduleObj.builtinTypeMap = builtinTypeMap

local builtinTypeList = NormalTypeInfo.createBuiltin( "List", "List", _moduleObj.TypeInfoKindList )

_moduleObj.builtinTypeList = builtinTypeList

local builtinTypeArray = NormalTypeInfo.createBuiltin( "Array", "Array", _moduleObj.TypeInfoKindArray )

_moduleObj.builtinTypeArray = builtinTypeArray

local builtinTypeForm = NormalTypeInfo.createBuiltin( "Form", "form", _moduleObj.TypeInfoKindFunc, _moduleObj.builtinTypeDDD )

_moduleObj.builtinTypeForm = builtinTypeForm

local builtinTypeSymbol = NormalTypeInfo.createBuiltin( "Symbol", "sym", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeSymbol = builtinTypeSymbol

local builtinTypeStat = NormalTypeInfo.createBuiltin( "Stat", "stat", _moduleObj.TypeInfoKindPrim )

_moduleObj.builtinTypeStat = builtinTypeStat

local builtinTypeStem_ = _lune_unwrap( _moduleObj.builtinTypeStem:get_nilableTypeInfo())

_moduleObj.builtinTypeStem_ = builtinTypeStem_

function EnumTypeInfo:serialize( stream, validChildrenSet )

  local txt = string.format( [==[{ parentId = %d, typeId = %d, txt = '%s',
accessMode = '%s', kind = %d, valTypeId = %d, ]==], self:getParentId(  ), self.typeId, self.rawTxt, self.accessMode, _moduleObj.TypeInfoKindEnum, self.valTypeInfo:get_typeId())
  
  stream:write( txt )
  stream:write( "enumValList = {" )
  do
    local __sorted = {}
    local __map = self.name2EnumValInfo
    for __key in pairs( __map ) do
      table.insert( __sorted, __key )
    end
    table.sort( __sorted )
    for __index, __key in ipairs( __sorted ) do
      enumValInfo = __map[ __key ]
      do
        if self.valTypeInfo == _moduleObj.builtinTypeString then
          stream:write( string.format( "%s = '%s',", enumValInfo:get_name(), enumValInfo:get_val()) )
        else 
          stream:write( string.format( "%s = %s,", enumValInfo:get_name(), enumValInfo:get_val()) )
        end
      end
    end
  end
  
  stream:write( "} }\n" )
end

function NilableTypeInfo:isSettableFrom( other )

  if not other then
    return false
  end
  if self == _moduleObj.builtinTypeStem_ then
    return true
  end
  if other == _moduleObj.builtinTypeNil then
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
  if (self:get_kind() ~= _moduleObj.TypeInfoKindClass and self:get_kind() ~= _moduleObj.TypeInfoKindIF ) or (other:get_kind() ~= _moduleObj.TypeInfoKindClass and other:get_kind() ~= _moduleObj.TypeInfoKindIF ) then
    return false
  end
  local baseTypeInfo = self:get_baseTypeInfo()
  
  while baseTypeInfo ~= _moduleObj.rootTypeInfo do
    if otherTypeId == baseTypeInfo:get_typeId() then
      return true
    end
    baseTypeInfo = baseTypeInfo:get_baseTypeInfo()
  end
  -- none
  
  for __index, ifType in pairs( self:get_interfaceList() ) do
    while ifType ~= _moduleObj.rootTypeInfo do
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
  if self == _moduleObj.builtinTypeStem_ or self == _moduleObj.builtinTypeDDD then
    return true
  end
  if self == _moduleObj.builtinTypeStem and not other:get_nilable() then
    return true
  end
  if self == _moduleObj.builtinTypeForm and other:get_kind() == _moduleObj.TypeInfoKindFunc then
    return true
  end
  if other == _moduleObj.builtinTypeNil then
    if self.kind ~= _moduleObj.TypeInfoKindNilable then
      return false
    end
    return true
  end
  if self.typeId == other:get_typeId() then
    return true
  end
  if self.kind ~= other:get_kind() then
    if self.kind == _moduleObj.TypeInfoKindNilable then
      if other:get_nilable() then
        return self:get_orgTypeInfo():isSettableFrom( other:get_orgTypeInfo() )
      end
      return self:get_orgTypeInfo():isSettableFrom( other )
    elseif (self:get_kind() == _moduleObj.TypeInfoKindClass or self:get_kind() == _moduleObj.TypeInfoKindIF ) and (other:get_kind() == _moduleObj.TypeInfoKindClass or other:get_kind() == _moduleObj.TypeInfoKindIF ) then
      return other:isInheritFrom( self )
    elseif other:get_kind() == _moduleObj.TypeInfoKindEnum then
      local enumTypeInfo = other
      
      return self:isSettableFrom( enumTypeInfo:get_valTypeInfo() )
    end
    return false
  end
  do
    local _switchExp = (self.kind )
    if _switchExp == _moduleObj.TypeInfoKindPrim then
      if self == _moduleObj.builtinTypeInt and other == _moduleObj.builtinTypeChar or self == _moduleObj.builtinTypeChar and other == _moduleObj.builtinTypeInt then
        return true
      end
      return false
    elseif _switchExp == _moduleObj.TypeInfoKindList or _switchExp == _moduleObj.TypeInfoKindArray then
      if other:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone then
        return true
      end
      if not (_lune_unwrap( self:get_itemTypeInfoList()[1]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[1]) ) then
        return false
      end
      
      return true
    elseif _switchExp == _moduleObj.TypeInfoKindMap then
      if other:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone and other:get_itemTypeInfoList()[2] == _moduleObj.builtinTypeNone then
        return true
      end
      if not (_lune_unwrap( self:get_itemTypeInfoList()[1]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[1]) ) then
        return false
      end
      
      if not (_lune_unwrap( self:get_itemTypeInfoList()[2]) ):isSettableFrom( _lune_unwrap( other:get_itemTypeInfoList()[2]) ) then
        return false
      end
      
      return true
    elseif _switchExp == _moduleObj.TypeInfoKindClass or _switchExp == _moduleObj.TypeInfoKindIF then
      return other:isInheritFrom( self )
    elseif _switchExp == _moduleObj.TypeInfoKindFunc then
      if self == _moduleObj.builtinTypeForm then
        return true
      end
      return false
    elseif _switchExp == _moduleObj.TypeInfoKindMethod then
      return false
    elseif _switchExp == _moduleObj.TypeInfoKindNilable then
      return self:get_orgTypeInfo():isSettableFrom( other:get_orgTypeInfo() )
    else 
      return false
    end
  end
  
  return true
end

local Filter = {}
_moduleObj.Filter = Filter
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
_moduleObj.Node = Node
function Node:get_expType(  )

  if not self.expTypeList then
    return _moduleObj.builtinTypeNone
  end
  return self.expTypeList[1]
end
function Node:getLiteral(  )

  return {nil}, {_moduleObj.builtinTypeNil}
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
_moduleObj.NamespaceInfo = NamespaceInfo
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
_moduleObj.DeclMacroInfo = DeclMacroInfo
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
_moduleObj.MacroValInfo = MacroValInfo
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
_moduleObj.MacroInfo = MacroInfo
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

_moduleObj.nodeKind = nodeKind

local function regKind( name )

  local kind = nodeKindSeed
  
  nodeKindSeed = nodeKindSeed + 1
  nodeKind2NameMap[kind] = name
  _moduleObj.nodeKind[name] = kind
  return kind
end

local function getNodeKindName( kind )

  return _lune_unwrap( nodeKind2NameMap[kind])
end
_moduleObj.getNodeKindName = getNodeKindName

-- none

function Filter:processNone( node, ... )

end

-- none

-- none

local nodeKindNone = regKind( [[None]] )

_moduleObj.nodeKindNone = nodeKindNone

local NoneNode = {}
setmetatable( NoneNode, { __index = Node } )
_moduleObj.NoneNode = NoneNode
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
function NoneNode.new( pos, typeList )
  local obj = {}
  setmetatable( obj, { __index = NoneNode } )
  if obj.__init then obj:__init( pos, typeList ); end
return obj
end
function NoneNode:__init(pos, typeList) 
  Node.__init( self, _moduleObj.nodeKindNone, pos, typeList)
  
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

_moduleObj.nodeKindSubfile = nodeKindSubfile

local SubfileNode = {}
setmetatable( SubfileNode, { __index = Node } )
_moduleObj.SubfileNode = SubfileNode
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
function SubfileNode.new( pos, typeList )
  local obj = {}
  setmetatable( obj, { __index = SubfileNode } )
  if obj.__init then obj:__init( pos, typeList ); end
return obj
end
function SubfileNode:__init(pos, typeList) 
  Node.__init( self, _moduleObj.nodeKindSubfile, pos, typeList)
  
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

_moduleObj.nodeKindImport = nodeKindImport

local ImportNode = {}
setmetatable( ImportNode, { __index = Node } )
_moduleObj.ImportNode = ImportNode
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
function ImportNode.new( pos, typeList, modulePath )
  local obj = {}
  setmetatable( obj, { __index = ImportNode } )
  if obj.__init then obj:__init( pos, typeList, modulePath ); end
return obj
end
function ImportNode:__init(pos, typeList, modulePath) 
  Node.__init( self, _moduleObj.nodeKindImport, pos, typeList)
  
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

-- none

function Filter:processRoot( node, ... )

end

-- none

-- none

local nodeKindRoot = regKind( [[Root]] )

_moduleObj.nodeKindRoot = nodeKindRoot

local RootNode = {}
setmetatable( RootNode, { __index = Node } )
_moduleObj.RootNode = RootNode
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
function RootNode.new( pos, typeList, children, provideNode, typeId2ClassMap )
  local obj = {}
  setmetatable( obj, { __index = RootNode } )
  if obj.__init then obj:__init( pos, typeList, children, provideNode, typeId2ClassMap ); end
return obj
end
function RootNode:__init(pos, typeList, children, provideNode, typeId2ClassMap) 
  Node.__init( self, _moduleObj.nodeKindRoot, pos, typeList)
  
  -- none
  
  self.children = children
  self.provideNode = provideNode
  self.typeId2ClassMap = typeId2ClassMap
  -- none
  
end
function RootNode:get_children()
  return self.children
end
function RootNode:get_provideNode()
  return self.provideNode
end
function RootNode:get_typeId2ClassMap()
  return self.typeId2ClassMap
end
do
  end


function RootNode:set_provide( node )

  self.provideNode = node
end

-- none

function Filter:processRefType( node, ... )

end

-- none

-- none

local nodeKindRefType = regKind( [[RefType]] )

_moduleObj.nodeKindRefType = nodeKindRefType

local RefTypeNode = {}
setmetatable( RefTypeNode, { __index = Node } )
_moduleObj.RefTypeNode = RefTypeNode
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
function RefTypeNode.new( pos, typeList, name, refFlag, mutFlag, array )
  local obj = {}
  setmetatable( obj, { __index = RefTypeNode } )
  if obj.__init then obj:__init( pos, typeList, name, refFlag, mutFlag, array ); end
return obj
end
function RefTypeNode:__init(pos, typeList, name, refFlag, mutFlag, array) 
  Node.__init( self, _moduleObj.nodeKindRefType, pos, typeList)
  
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

_moduleObj.nodeKindBlock = nodeKindBlock

local BlockNode = {}
setmetatable( BlockNode, { __index = Node } )
_moduleObj.BlockNode = BlockNode
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
function BlockNode.new( pos, typeList, blockKind, stmtList )
  local obj = {}
  setmetatable( obj, { __index = BlockNode } )
  if obj.__init then obj:__init( pos, typeList, blockKind, stmtList ); end
return obj
end
function BlockNode:__init(pos, typeList, blockKind, stmtList) 
  Node.__init( self, _moduleObj.nodeKindBlock, pos, typeList)
  
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
_moduleObj.IfStmtInfo = IfStmtInfo
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

_moduleObj.nodeKindIf = nodeKindIf

local IfNode = {}
setmetatable( IfNode, { __index = Node } )
_moduleObj.IfNode = IfNode
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
function IfNode.new( pos, typeList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = IfNode } )
  if obj.__init then obj:__init( pos, typeList, stmtList ); end
return obj
end
function IfNode:__init(pos, typeList, stmtList) 
  Node.__init( self, _moduleObj.nodeKindIf, pos, typeList)
  
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

_moduleObj.nodeKindExpList = nodeKindExpList

local ExpListNode = {}
setmetatable( ExpListNode, { __index = Node } )
_moduleObj.ExpListNode = ExpListNode
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
function ExpListNode.new( pos, typeList, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpListNode } )
  if obj.__init then obj:__init( pos, typeList, expList ); end
return obj
end
function ExpListNode:__init(pos, typeList, expList) 
  Node.__init( self, _moduleObj.nodeKindExpList, pos, typeList)
  
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
_moduleObj.CaseInfo = CaseInfo
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

_moduleObj.nodeKindSwitch = nodeKindSwitch

local SwitchNode = {}
setmetatable( SwitchNode, { __index = Node } )
_moduleObj.SwitchNode = SwitchNode
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
function SwitchNode.new( pos, typeList, exp, caseList, default )
  local obj = {}
  setmetatable( obj, { __index = SwitchNode } )
  if obj.__init then obj:__init( pos, typeList, exp, caseList, default ); end
return obj
end
function SwitchNode:__init(pos, typeList, exp, caseList, default) 
  Node.__init( self, _moduleObj.nodeKindSwitch, pos, typeList)
  
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

_moduleObj.nodeKindWhile = nodeKindWhile

local WhileNode = {}
setmetatable( WhileNode, { __index = Node } )
_moduleObj.WhileNode = WhileNode
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
function WhileNode.new( pos, typeList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = WhileNode } )
  if obj.__init then obj:__init( pos, typeList, exp, block ); end
return obj
end
function WhileNode:__init(pos, typeList, exp, block) 
  Node.__init( self, _moduleObj.nodeKindWhile, pos, typeList)
  
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

_moduleObj.nodeKindRepeat = nodeKindRepeat

local RepeatNode = {}
setmetatable( RepeatNode, { __index = Node } )
_moduleObj.RepeatNode = RepeatNode
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
function RepeatNode.new( pos, typeList, block, exp )
  local obj = {}
  setmetatable( obj, { __index = RepeatNode } )
  if obj.__init then obj:__init( pos, typeList, block, exp ); end
return obj
end
function RepeatNode:__init(pos, typeList, block, exp) 
  Node.__init( self, _moduleObj.nodeKindRepeat, pos, typeList)
  
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

_moduleObj.nodeKindFor = nodeKindFor

local ForNode = {}
setmetatable( ForNode, { __index = Node } )
_moduleObj.ForNode = ForNode
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
function ForNode.new( pos, typeList, block, val, init, to, delta )
  local obj = {}
  setmetatable( obj, { __index = ForNode } )
  if obj.__init then obj:__init( pos, typeList, block, val, init, to, delta ); end
return obj
end
function ForNode:__init(pos, typeList, block, val, init, to, delta) 
  Node.__init( self, _moduleObj.nodeKindFor, pos, typeList)
  
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

_moduleObj.nodeKindApply = nodeKindApply

local ApplyNode = {}
setmetatable( ApplyNode, { __index = Node } )
_moduleObj.ApplyNode = ApplyNode
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
function ApplyNode.new( pos, typeList, varList, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ApplyNode } )
  if obj.__init then obj:__init( pos, typeList, varList, exp, block ); end
return obj
end
function ApplyNode:__init(pos, typeList, varList, exp, block) 
  Node.__init( self, _moduleObj.nodeKindApply, pos, typeList)
  
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

_moduleObj.nodeKindForeach = nodeKindForeach

local ForeachNode = {}
setmetatable( ForeachNode, { __index = Node } )
_moduleObj.ForeachNode = ForeachNode
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
function ForeachNode.new( pos, typeList, val, key, exp, block )
  local obj = {}
  setmetatable( obj, { __index = ForeachNode } )
  if obj.__init then obj:__init( pos, typeList, val, key, exp, block ); end
return obj
end
function ForeachNode:__init(pos, typeList, val, key, exp, block) 
  Node.__init( self, _moduleObj.nodeKindForeach, pos, typeList)
  
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

_moduleObj.nodeKindForsort = nodeKindForsort

local ForsortNode = {}
setmetatable( ForsortNode, { __index = Node } )
_moduleObj.ForsortNode = ForsortNode
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
function ForsortNode.new( pos, typeList, val, key, exp, block, sort )
  local obj = {}
  setmetatable( obj, { __index = ForsortNode } )
  if obj.__init then obj:__init( pos, typeList, val, key, exp, block, sort ); end
return obj
end
function ForsortNode:__init(pos, typeList, val, key, exp, block, sort) 
  Node.__init( self, _moduleObj.nodeKindForsort, pos, typeList)
  
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

_moduleObj.nodeKindReturn = nodeKindReturn

local ReturnNode = {}
setmetatable( ReturnNode, { __index = Node } )
_moduleObj.ReturnNode = ReturnNode
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
function ReturnNode.new( pos, typeList, expList )
  local obj = {}
  setmetatable( obj, { __index = ReturnNode } )
  if obj.__init then obj:__init( pos, typeList, expList ); end
return obj
end
function ReturnNode:__init(pos, typeList, expList) 
  Node.__init( self, _moduleObj.nodeKindReturn, pos, typeList)
  
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

_moduleObj.nodeKindBreak = nodeKindBreak

local BreakNode = {}
setmetatable( BreakNode, { __index = Node } )
_moduleObj.BreakNode = BreakNode
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
function BreakNode.new( pos, typeList )
  local obj = {}
  setmetatable( obj, { __index = BreakNode } )
  if obj.__init then obj:__init( pos, typeList ); end
return obj
end
function BreakNode:__init(pos, typeList) 
  Node.__init( self, _moduleObj.nodeKindBreak, pos, typeList)
  
  -- none
  
  -- none
  
end
do
  end


-- none

function Filter:processProvide( node, ... )

end

-- none

-- none

local nodeKindProvide = regKind( [[Provide]] )

_moduleObj.nodeKindProvide = nodeKindProvide

local ProvideNode = {}
setmetatable( ProvideNode, { __index = Node } )
_moduleObj.ProvideNode = ProvideNode
function ProvideNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processProvide( self, table.unpack( argList ) )
end
function ProvideNode:canBeRight(  )

  return false
end
function ProvideNode:canBeLeft(  )

  return false
end
function ProvideNode.new( pos, typeList, val )
  local obj = {}
  setmetatable( obj, { __index = ProvideNode } )
  if obj.__init then obj:__init( pos, typeList, val ); end
return obj
end
function ProvideNode:__init(pos, typeList, val) 
  Node.__init( self, _moduleObj.nodeKindProvide, pos, typeList)
  
  -- none
  
  self.val = val
  -- none
  
end
function ProvideNode:get_val()
  return self.val
end
do
  end


-- none

function Filter:processExpNew( node, ... )

end

-- none

-- none

local nodeKindExpNew = regKind( [[ExpNew]] )

_moduleObj.nodeKindExpNew = nodeKindExpNew

local ExpNewNode = {}
setmetatable( ExpNewNode, { __index = Node } )
_moduleObj.ExpNewNode = ExpNewNode
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
function ExpNewNode.new( pos, typeList, symbol, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpNewNode } )
  if obj.__init then obj:__init( pos, typeList, symbol, argList ); end
return obj
end
function ExpNewNode:__init(pos, typeList, symbol, argList) 
  Node.__init( self, _moduleObj.nodeKindExpNew, pos, typeList)
  
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

_moduleObj.nodeKindExpUnwrap = nodeKindExpUnwrap

local ExpUnwrapNode = {}
setmetatable( ExpUnwrapNode, { __index = Node } )
_moduleObj.ExpUnwrapNode = ExpUnwrapNode
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
function ExpUnwrapNode.new( pos, typeList, exp, default )
  local obj = {}
  setmetatable( obj, { __index = ExpUnwrapNode } )
  if obj.__init then obj:__init( pos, typeList, exp, default ); end
return obj
end
function ExpUnwrapNode:__init(pos, typeList, exp, default) 
  Node.__init( self, _moduleObj.nodeKindExpUnwrap, pos, typeList)
  
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

_moduleObj.nodeKindExpRef = nodeKindExpRef

local ExpRefNode = {}
setmetatable( ExpRefNode, { __index = Node } )
_moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processExpRef( self, table.unpack( argList ) )
end
function ExpRefNode.new( pos, typeList, token, symbolInfo )
  local obj = {}
  setmetatable( obj, { __index = ExpRefNode } )
  if obj.__init then obj:__init( pos, typeList, token, symbolInfo ); end
return obj
end
function ExpRefNode:__init(pos, typeList, token, symbolInfo) 
  Node.__init( self, _moduleObj.nodeKindExpRef, pos, typeList)
  
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

_moduleObj.nodeKindExpOp2 = nodeKindExpOp2

local ExpOp2Node = {}
setmetatable( ExpOp2Node, { __index = Node } )
_moduleObj.ExpOp2Node = ExpOp2Node
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
function ExpOp2Node.new( pos, typeList, op, exp1, exp2 )
  local obj = {}
  setmetatable( obj, { __index = ExpOp2Node } )
  if obj.__init then obj:__init( pos, typeList, op, exp1, exp2 ); end
return obj
end
function ExpOp2Node:__init(pos, typeList, op, exp1, exp2) 
  Node.__init( self, _moduleObj.nodeKindExpOp2, pos, typeList)
  
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

_moduleObj.nodeKindUnwrapSet = nodeKindUnwrapSet

local UnwrapSetNode = {}
setmetatable( UnwrapSetNode, { __index = Node } )
_moduleObj.UnwrapSetNode = UnwrapSetNode
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
function UnwrapSetNode.new( pos, typeList, dstExpList, srcExpList, unwrapBlock )
  local obj = {}
  setmetatable( obj, { __index = UnwrapSetNode } )
  if obj.__init then obj:__init( pos, typeList, dstExpList, srcExpList, unwrapBlock ); end
return obj
end
function UnwrapSetNode:__init(pos, typeList, dstExpList, srcExpList, unwrapBlock) 
  Node.__init( self, _moduleObj.nodeKindUnwrapSet, pos, typeList)
  
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

_moduleObj.nodeKindIfUnwrap = nodeKindIfUnwrap

local IfUnwrapNode = {}
setmetatable( IfUnwrapNode, { __index = Node } )
_moduleObj.IfUnwrapNode = IfUnwrapNode
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
function IfUnwrapNode.new( pos, typeList, exp, block, nilBlock )
  local obj = {}
  setmetatable( obj, { __index = IfUnwrapNode } )
  if obj.__init then obj:__init( pos, typeList, exp, block, nilBlock ); end
return obj
end
function IfUnwrapNode:__init(pos, typeList, exp, block, nilBlock) 
  Node.__init( self, _moduleObj.nodeKindIfUnwrap, pos, typeList)
  
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

_moduleObj.nodeKindExpCast = nodeKindExpCast

local ExpCastNode = {}
setmetatable( ExpCastNode, { __index = Node } )
_moduleObj.ExpCastNode = ExpCastNode
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
function ExpCastNode.new( pos, typeList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpCastNode } )
  if obj.__init then obj:__init( pos, typeList, exp ); end
return obj
end
function ExpCastNode:__init(pos, typeList, exp) 
  Node.__init( self, _moduleObj.nodeKindExpCast, pos, typeList)
  
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

_moduleObj.nodeKindExpOp1 = nodeKindExpOp1

local ExpOp1Node = {}
setmetatable( ExpOp1Node, { __index = Node } )
_moduleObj.ExpOp1Node = ExpOp1Node
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
function ExpOp1Node.new( pos, typeList, op, macroMode, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpOp1Node } )
  if obj.__init then obj:__init( pos, typeList, op, macroMode, exp ); end
return obj
end
function ExpOp1Node:__init(pos, typeList, op, macroMode, exp) 
  Node.__init( self, _moduleObj.nodeKindExpOp1, pos, typeList)
  
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

_moduleObj.nodeKindExpRefItem = nodeKindExpRefItem

local ExpRefItemNode = {}
setmetatable( ExpRefItemNode, { __index = Node } )
_moduleObj.ExpRefItemNode = ExpRefItemNode
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
function ExpRefItemNode.new( pos, typeList, val, nilAccess, index )
  local obj = {}
  setmetatable( obj, { __index = ExpRefItemNode } )
  if obj.__init then obj:__init( pos, typeList, val, nilAccess, index ); end
return obj
end
function ExpRefItemNode:__init(pos, typeList, val, nilAccess, index) 
  Node.__init( self, _moduleObj.nodeKindExpRefItem, pos, typeList)
  
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

_moduleObj.nodeKindExpCall = nodeKindExpCall

local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
_moduleObj.ExpCallNode = ExpCallNode
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
function ExpCallNode.new( pos, typeList, func, nilAccess, argList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallNode } )
  if obj.__init then obj:__init( pos, typeList, func, nilAccess, argList ); end
return obj
end
function ExpCallNode:__init(pos, typeList, func, nilAccess, argList) 
  Node.__init( self, _moduleObj.nodeKindExpCall, pos, typeList)
  
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

_moduleObj.nodeKindExpDDD = nodeKindExpDDD

local ExpDDDNode = {}
setmetatable( ExpDDDNode, { __index = Node } )
_moduleObj.ExpDDDNode = ExpDDDNode
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
function ExpDDDNode.new( pos, typeList, token )
  local obj = {}
  setmetatable( obj, { __index = ExpDDDNode } )
  if obj.__init then obj:__init( pos, typeList, token ); end
return obj
end
function ExpDDDNode:__init(pos, typeList, token) 
  Node.__init( self, _moduleObj.nodeKindExpDDD, pos, typeList)
  
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

_moduleObj.nodeKindExpParen = nodeKindExpParen

local ExpParenNode = {}
setmetatable( ExpParenNode, { __index = Node } )
_moduleObj.ExpParenNode = ExpParenNode
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
function ExpParenNode.new( pos, typeList, exp )
  local obj = {}
  setmetatable( obj, { __index = ExpParenNode } )
  if obj.__init then obj:__init( pos, typeList, exp ); end
return obj
end
function ExpParenNode:__init(pos, typeList, exp) 
  Node.__init( self, _moduleObj.nodeKindExpParen, pos, typeList)
  
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

_moduleObj.nodeKindExpMacroExp = nodeKindExpMacroExp

local ExpMacroExpNode = {}
setmetatable( ExpMacroExpNode, { __index = Node } )
_moduleObj.ExpMacroExpNode = ExpMacroExpNode
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
function ExpMacroExpNode.new( pos, typeList, stmtList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroExpNode } )
  if obj.__init then obj:__init( pos, typeList, stmtList ); end
return obj
end
function ExpMacroExpNode:__init(pos, typeList, stmtList) 
  Node.__init( self, _moduleObj.nodeKindExpMacroExp, pos, typeList)
  
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

_moduleObj.nodeKindExpMacroStat = nodeKindExpMacroStat

local ExpMacroStatNode = {}
setmetatable( ExpMacroStatNode, { __index = Node } )
_moduleObj.ExpMacroStatNode = ExpMacroStatNode
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
function ExpMacroStatNode.new( pos, typeList, expStrList )
  local obj = {}
  setmetatable( obj, { __index = ExpMacroStatNode } )
  if obj.__init then obj:__init( pos, typeList, expStrList ); end
return obj
end
function ExpMacroStatNode:__init(pos, typeList, expStrList) 
  Node.__init( self, _moduleObj.nodeKindExpMacroStat, pos, typeList)
  
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

_moduleObj.nodeKindStmtExp = nodeKindStmtExp

local StmtExpNode = {}
setmetatable( StmtExpNode, { __index = Node } )
_moduleObj.StmtExpNode = StmtExpNode
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
function StmtExpNode.new( pos, typeList, exp )
  local obj = {}
  setmetatable( obj, { __index = StmtExpNode } )
  if obj.__init then obj:__init( pos, typeList, exp ); end
return obj
end
function StmtExpNode:__init(pos, typeList, exp) 
  Node.__init( self, _moduleObj.nodeKindStmtExp, pos, typeList)
  
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

_moduleObj.nodeKindRefField = nodeKindRefField

local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processRefField( self, table.unpack( argList ) )
end
function RefFieldNode:canBeRight(  )

  return true
end
function RefFieldNode.new( pos, typeList, field, symbolInfo, nilAccess, prefix )
  local obj = {}
  setmetatable( obj, { __index = RefFieldNode } )
  if obj.__init then obj:__init( pos, typeList, field, symbolInfo, nilAccess, prefix ); end
return obj
end
function RefFieldNode:__init(pos, typeList, field, symbolInfo, nilAccess, prefix) 
  Node.__init( self, _moduleObj.nodeKindRefField, pos, typeList)
  
  -- none
  
  self.field = field
  self.symbolInfo = symbolInfo
  self.nilAccess = nilAccess
  self.prefix = prefix
  -- none
  
end
function RefFieldNode:get_field()
  return self.field
end
function RefFieldNode:get_symbolInfo()
  return self.symbolInfo
end
function RefFieldNode:get_nilAccess()
  return self.nilAccess
end
function RefFieldNode:get_prefix()
  return self.prefix
end
do
  end


function RefFieldNode:canBeLeft(  )

  do
    local _exp = self:get_symbolInfo()
    if _exp ~= nil then
    
        return _exp:get_canBeLeft()
      end
  end
  
  return false
end

-- none

function Filter:processGetField( node, ... )

end

-- none

-- none

local nodeKindGetField = regKind( [[GetField]] )

_moduleObj.nodeKindGetField = nodeKindGetField

local GetFieldNode = {}
setmetatable( GetFieldNode, { __index = Node } )
_moduleObj.GetFieldNode = GetFieldNode
function GetFieldNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processGetField( self, table.unpack( argList ) )
end
function GetFieldNode:canBeRight(  )

  return true
end
function GetFieldNode.new( pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )
  local obj = {}
  setmetatable( obj, { __index = GetFieldNode } )
  if obj.__init then obj:__init( pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo ); end
return obj
end
function GetFieldNode:__init(pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo) 
  Node.__init( self, _moduleObj.nodeKindGetField, pos, typeList)
  
  -- none
  
  self.field = field
  self.symbolInfo = symbolInfo
  self.nilAccess = nilAccess
  self.prefix = prefix
  self.getterTypeInfo = getterTypeInfo
  -- none
  
end
function GetFieldNode:get_field()
  return self.field
end
function GetFieldNode:get_symbolInfo()
  return self.symbolInfo
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


function GetFieldNode:canBeLeft(  )

  do
    local _exp = self:get_symbolInfo()
    if _exp ~= nil then
    
        return _exp:get_canBeLeft()
      end
  end
  
  return false
end

local VarInfo = {}
_moduleObj.VarInfo = VarInfo
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

_moduleObj.nodeKindDeclVar = nodeKindDeclVar

local DeclVarNode = {}
setmetatable( DeclVarNode, { __index = Node } )
_moduleObj.DeclVarNode = DeclVarNode
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
function DeclVarNode.new( pos, typeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )
  local obj = {}
  setmetatable( obj, { __index = DeclVarNode } )
  if obj.__init then obj:__init( pos, typeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock ); end
return obj
end
function DeclVarNode:__init(pos, typeList, mode, accessMode, staticFlag, varList, expList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock) 
  Node.__init( self, _moduleObj.nodeKindDeclVar, pos, typeList)
  
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
_moduleObj.DeclFuncInfo = DeclFuncInfo
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

_moduleObj.nodeKindDeclFunc = nodeKindDeclFunc

local DeclFuncNode = {}
setmetatable( DeclFuncNode, { __index = Node } )
_moduleObj.DeclFuncNode = DeclFuncNode
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
function DeclFuncNode.new( pos, typeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclFuncNode } )
  if obj.__init then obj:__init( pos, typeList, declInfo ); end
return obj
end
function DeclFuncNode:__init(pos, typeList, declInfo) 
  Node.__init( self, _moduleObj.nodeKindDeclFunc, pos, typeList)
  
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

_moduleObj.nodeKindDeclMethod = nodeKindDeclMethod

local DeclMethodNode = {}
setmetatable( DeclMethodNode, { __index = Node } )
_moduleObj.DeclMethodNode = DeclMethodNode
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
function DeclMethodNode.new( pos, typeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMethodNode } )
  if obj.__init then obj:__init( pos, typeList, declInfo ); end
return obj
end
function DeclMethodNode:__init(pos, typeList, declInfo) 
  Node.__init( self, _moduleObj.nodeKindDeclMethod, pos, typeList)
  
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

_moduleObj.nodeKindDeclConstr = nodeKindDeclConstr

local DeclConstrNode = {}
setmetatable( DeclConstrNode, { __index = Node } )
_moduleObj.DeclConstrNode = DeclConstrNode
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
function DeclConstrNode.new( pos, typeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclConstrNode } )
  if obj.__init then obj:__init( pos, typeList, declInfo ); end
return obj
end
function DeclConstrNode:__init(pos, typeList, declInfo) 
  Node.__init( self, _moduleObj.nodeKindDeclConstr, pos, typeList)
  
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

_moduleObj.nodeKindExpCallSuper = nodeKindExpCallSuper

local ExpCallSuperNode = {}
setmetatable( ExpCallSuperNode, { __index = Node } )
_moduleObj.ExpCallSuperNode = ExpCallSuperNode
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
function ExpCallSuperNode.new( pos, typeList, superType, expList )
  local obj = {}
  setmetatable( obj, { __index = ExpCallSuperNode } )
  if obj.__init then obj:__init( pos, typeList, superType, expList ); end
return obj
end
function ExpCallSuperNode:__init(pos, typeList, superType, expList) 
  Node.__init( self, _moduleObj.nodeKindExpCallSuper, pos, typeList)
  
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

_moduleObj.nodeKindDeclMember = nodeKindDeclMember

local DeclMemberNode = {}
setmetatable( DeclMemberNode, { __index = Node } )
_moduleObj.DeclMemberNode = DeclMemberNode
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
function DeclMemberNode.new( pos, typeList, name, refType, staticFlag, accessMode, getterMode, setterMode )
  local obj = {}
  setmetatable( obj, { __index = DeclMemberNode } )
  if obj.__init then obj:__init( pos, typeList, name, refType, staticFlag, accessMode, getterMode, setterMode ); end
return obj
end
function DeclMemberNode:__init(pos, typeList, name, refType, staticFlag, accessMode, getterMode, setterMode) 
  Node.__init( self, _moduleObj.nodeKindDeclMember, pos, typeList)
  
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

_moduleObj.nodeKindDeclArg = nodeKindDeclArg

local DeclArgNode = {}
setmetatable( DeclArgNode, { __index = Node } )
_moduleObj.DeclArgNode = DeclArgNode
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
function DeclArgNode.new( pos, typeList, name, argType )
  local obj = {}
  setmetatable( obj, { __index = DeclArgNode } )
  if obj.__init then obj:__init( pos, typeList, name, argType ); end
return obj
end
function DeclArgNode:__init(pos, typeList, name, argType) 
  Node.__init( self, _moduleObj.nodeKindDeclArg, pos, typeList)
  
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

_moduleObj.nodeKindDeclArgDDD = nodeKindDeclArgDDD

local DeclArgDDDNode = {}
setmetatable( DeclArgDDDNode, { __index = Node } )
_moduleObj.DeclArgDDDNode = DeclArgDDDNode
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
function DeclArgDDDNode.new( pos, typeList )
  local obj = {}
  setmetatable( obj, { __index = DeclArgDDDNode } )
  if obj.__init then obj:__init( pos, typeList ); end
return obj
end
function DeclArgDDDNode:__init(pos, typeList) 
  Node.__init( self, _moduleObj.nodeKindDeclArgDDD, pos, typeList)
  
  -- none
  
  -- none
  
end
do
  end


local AdvertiseInfo = {}
_moduleObj.AdvertiseInfo = AdvertiseInfo
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

_moduleObj.nodeKindDeclClass = nodeKindDeclClass

local DeclClassNode = {}
setmetatable( DeclClassNode, { __index = Node } )
_moduleObj.DeclClassNode = DeclClassNode
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
function DeclClassNode.new( pos, typeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet )
  local obj = {}
  setmetatable( obj, { __index = DeclClassNode } )
  if obj.__init then obj:__init( pos, typeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet ); end
return obj
end
function DeclClassNode:__init(pos, typeList, accessMode, name, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet) 
  Node.__init( self, _moduleObj.nodeKindDeclClass, pos, typeList)
  
  -- none
  
  self.accessMode = accessMode
  self.name = name
  self.fieldList = fieldList
  self.moduleName = moduleName
  self.memberList = memberList
  self.scope = scope
  self.initStmtList = initStmtList
  self.advertiseList = advertiseList
  self.trustList = trustList
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
function DeclClassNode:get_trustList()
  return self.trustList
end
function DeclClassNode:get_outerMethodSet()
  return self.outerMethodSet
end
do
  end


-- none

function Filter:processDeclEnum( node, ... )

end

-- none

-- none

local nodeKindDeclEnum = regKind( [[DeclEnum]] )

_moduleObj.nodeKindDeclEnum = nodeKindDeclEnum

local DeclEnumNode = {}
setmetatable( DeclEnumNode, { __index = Node } )
_moduleObj.DeclEnumNode = DeclEnumNode
function DeclEnumNode:processFilter( filter, ... )

  local argList = {...}
  
  filter:processDeclEnum( self, table.unpack( argList ) )
end
function DeclEnumNode:canBeRight(  )

  return false
end
function DeclEnumNode:canBeLeft(  )

  return false
end
function DeclEnumNode.new( pos, typeList, accessMode, name, valueNameList, scope )
  local obj = {}
  setmetatable( obj, { __index = DeclEnumNode } )
  if obj.__init then obj:__init( pos, typeList, accessMode, name, valueNameList, scope ); end
return obj
end
function DeclEnumNode:__init(pos, typeList, accessMode, name, valueNameList, scope) 
  Node.__init( self, _moduleObj.nodeKindDeclEnum, pos, typeList)
  
  -- none
  
  self.accessMode = accessMode
  self.name = name
  self.valueNameList = valueNameList
  self.scope = scope
  -- none
  
end
function DeclEnumNode:get_accessMode()
  return self.accessMode
end
function DeclEnumNode:get_name()
  return self.name
end
function DeclEnumNode:get_valueNameList()
  return self.valueNameList
end
function DeclEnumNode:get_scope()
  return self.scope
end
do
  end


-- none

function Filter:processDeclMacro( node, ... )

end

-- none

-- none

local nodeKindDeclMacro = regKind( [[DeclMacro]] )

_moduleObj.nodeKindDeclMacro = nodeKindDeclMacro

local DeclMacroNode = {}
setmetatable( DeclMacroNode, { __index = Node } )
_moduleObj.DeclMacroNode = DeclMacroNode
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
function DeclMacroNode.new( pos, typeList, declInfo )
  local obj = {}
  setmetatable( obj, { __index = DeclMacroNode } )
  if obj.__init then obj:__init( pos, typeList, declInfo ); end
return obj
end
function DeclMacroNode:__init(pos, typeList, declInfo) 
  Node.__init( self, _moduleObj.nodeKindDeclMacro, pos, typeList)
  
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
_moduleObj.MacroEval = MacroEval
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

_moduleObj.nodeKindLiteralNil = nodeKindLiteralNil

local LiteralNilNode = {}
setmetatable( LiteralNilNode, { __index = Node } )
_moduleObj.LiteralNilNode = LiteralNilNode
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
function LiteralNilNode.new( pos, typeList )
  local obj = {}
  setmetatable( obj, { __index = LiteralNilNode } )
  if obj.__init then obj:__init( pos, typeList ); end
return obj
end
function LiteralNilNode:__init(pos, typeList) 
  Node.__init( self, _moduleObj.nodeKindLiteralNil, pos, typeList)
  
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

_moduleObj.nodeKindLiteralChar = nodeKindLiteralChar

local LiteralCharNode = {}
setmetatable( LiteralCharNode, { __index = Node } )
_moduleObj.LiteralCharNode = LiteralCharNode
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
function LiteralCharNode.new( pos, typeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralCharNode } )
  if obj.__init then obj:__init( pos, typeList, token, num ); end
return obj
end
function LiteralCharNode:__init(pos, typeList, token, num) 
  Node.__init( self, _moduleObj.nodeKindLiteralChar, pos, typeList)
  
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

_moduleObj.nodeKindLiteralInt = nodeKindLiteralInt

local LiteralIntNode = {}
setmetatable( LiteralIntNode, { __index = Node } )
_moduleObj.LiteralIntNode = LiteralIntNode
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
function LiteralIntNode.new( pos, typeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralIntNode } )
  if obj.__init then obj:__init( pos, typeList, token, num ); end
return obj
end
function LiteralIntNode:__init(pos, typeList, token, num) 
  Node.__init( self, _moduleObj.nodeKindLiteralInt, pos, typeList)
  
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

_moduleObj.nodeKindLiteralReal = nodeKindLiteralReal

local LiteralRealNode = {}
setmetatable( LiteralRealNode, { __index = Node } )
_moduleObj.LiteralRealNode = LiteralRealNode
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
function LiteralRealNode.new( pos, typeList, token, num )
  local obj = {}
  setmetatable( obj, { __index = LiteralRealNode } )
  if obj.__init then obj:__init( pos, typeList, token, num ); end
return obj
end
function LiteralRealNode:__init(pos, typeList, token, num) 
  Node.__init( self, _moduleObj.nodeKindLiteralReal, pos, typeList)
  
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

_moduleObj.nodeKindLiteralArray = nodeKindLiteralArray

local LiteralArrayNode = {}
setmetatable( LiteralArrayNode, { __index = Node } )
_moduleObj.LiteralArrayNode = LiteralArrayNode
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
function LiteralArrayNode.new( pos, typeList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralArrayNode } )
  if obj.__init then obj:__init( pos, typeList, expList ); end
return obj
end
function LiteralArrayNode:__init(pos, typeList, expList) 
  Node.__init( self, _moduleObj.nodeKindLiteralArray, pos, typeList)
  
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

_moduleObj.nodeKindLiteralList = nodeKindLiteralList

local LiteralListNode = {}
setmetatable( LiteralListNode, { __index = Node } )
_moduleObj.LiteralListNode = LiteralListNode
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
function LiteralListNode.new( pos, typeList, expList )
  local obj = {}
  setmetatable( obj, { __index = LiteralListNode } )
  if obj.__init then obj:__init( pos, typeList, expList ); end
return obj
end
function LiteralListNode:__init(pos, typeList, expList) 
  Node.__init( self, _moduleObj.nodeKindLiteralList, pos, typeList)
  
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
_moduleObj.PairItem = PairItem
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

_moduleObj.nodeKindLiteralMap = nodeKindLiteralMap

local LiteralMapNode = {}
setmetatable( LiteralMapNode, { __index = Node } )
_moduleObj.LiteralMapNode = LiteralMapNode
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
function LiteralMapNode.new( pos, typeList, map, pairList )
  local obj = {}
  setmetatable( obj, { __index = LiteralMapNode } )
  if obj.__init then obj:__init( pos, typeList, map, pairList ); end
return obj
end
function LiteralMapNode:__init(pos, typeList, map, pairList) 
  Node.__init( self, _moduleObj.nodeKindLiteralMap, pos, typeList)
  
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

_moduleObj.nodeKindLiteralString = nodeKindLiteralString

local LiteralStringNode = {}
setmetatable( LiteralStringNode, { __index = Node } )
_moduleObj.LiteralStringNode = LiteralStringNode
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
function LiteralStringNode.new( pos, typeList, token, argList )
  local obj = {}
  setmetatable( obj, { __index = LiteralStringNode } )
  if obj.__init then obj:__init( pos, typeList, token, argList ); end
return obj
end
function LiteralStringNode:__init(pos, typeList, token, argList) 
  Node.__init( self, _moduleObj.nodeKindLiteralString, pos, typeList)
  
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

_moduleObj.nodeKindLiteralBool = nodeKindLiteralBool

local LiteralBoolNode = {}
setmetatable( LiteralBoolNode, { __index = Node } )
_moduleObj.LiteralBoolNode = LiteralBoolNode
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
function LiteralBoolNode.new( pos, typeList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralBoolNode } )
  if obj.__init then obj:__init( pos, typeList, token ); end
return obj
end
function LiteralBoolNode:__init(pos, typeList, token) 
  Node.__init( self, _moduleObj.nodeKindLiteralBool, pos, typeList)
  
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

_moduleObj.nodeKindLiteralSymbol = nodeKindLiteralSymbol

local LiteralSymbolNode = {}
setmetatable( LiteralSymbolNode, { __index = Node } )
_moduleObj.LiteralSymbolNode = LiteralSymbolNode
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
function LiteralSymbolNode.new( pos, typeList, token )
  local obj = {}
  setmetatable( obj, { __index = LiteralSymbolNode } )
  if obj.__init then obj:__init( pos, typeList, token ); end
return obj
end
function LiteralSymbolNode:__init(pos, typeList, token) 
  Node.__init( self, _moduleObj.nodeKindLiteralSymbol, pos, typeList)
  
  -- none
  
  self.token = token
  -- none
  
end
function LiteralSymbolNode:get_token()
  return self.token
end
do
  end


function Node:getSymbolInfo(  )

  local function processExpNode( node )
  
    do
      local _switchExp = (node:get_kind() )
      if _switchExp == _moduleObj.nodeKindExpRef then
        return {(node ):get_symbolInfo()}
      elseif _switchExp == _moduleObj.nodeKindRefField then
        local refFieldNode = node
        
        do
          local _exp = refFieldNode:get_symbolInfo()
          if _exp ~= nil then
          
              return {_exp}
            end
        end
        
        return {}
      elseif _switchExp == _moduleObj.nodeKindGetField then
        local getFieldNode = node
        
        do
          local _exp = getFieldNode:get_symbolInfo()
          if _exp ~= nil then
          
              return {_exp}
            end
        end
        
        return {}
      elseif _switchExp == _moduleObj.nodeKindExpList then
        local expListNode = node
        
        local list = {}
        
        for index, expNode in pairs( expListNode:get_expList() ) do
          if index == #expListNode:get_expList() then
            for __index, symbolInfo in pairs( processExpNode( expNode ) ) do
              table.insert( list, symbolInfo )
            end
          else 
            for __index, symbolInfo in pairs( processExpNode( expNode ) ) do
              table.insert( list, symbolInfo )
              break
            end
          end
        end
        return list
      end
    end
    
    return {}
  end
  
  return processExpNode( self )
end

function LiteralNilNode:getLiteral(  )

  return {nil}, {_moduleObj.builtinTypeNil}
end

function LiteralCharNode:getLiteral(  )

  return {self.num}, {_moduleObj.builtinTypeChar}
end

function LiteralIntNode:getLiteral(  )

  return {self.num}, {_moduleObj.builtinTypeInt}
end

function LiteralRealNode:getLiteral(  )

  return {self.num}, {_moduleObj.builtinTypeReal}
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
    return {string.format( txt, table.unpack( argTbl ) )}, {_moduleObj.builtinTypeString}
  end
  return {txt}, {_moduleObj.builtinTypeString}
end

function LiteralBoolNode:getLiteral(  )

  return {self.token.txt == "true"}, {_moduleObj.builtinTypeBool}
end

function LiteralSymbolNode:getLiteral(  )

  return {{self.token.txt}}, {_moduleObj.builtinTypeSymbol}
end

function RefFieldNode:getLiteral(  )

  local prefix = (_lune_unwrap( self.prefix:getLiteral(  )[1]) )
  
  if self.nilAccess then
    table.insert( prefix, "$." )
  else 
    table.insert( prefix, "." )
  end
  table.insert( prefix, self.field.txt )
  return {prefix}, {_moduleObj.builtinTypeSymbol}
end

function ExpMacroStatNode:getLiteral(  )

  local txt = ""
  
  for __index, token in pairs( self.expStrList ) do
    txt = string.format( "%s%s", txt, token:getLiteral(  )[1])
  end
  return {txt}, {self:get_expType(  )}
end

function ExpRefNode:getLiteral(  )

  local typeInfo = self.symbolInfo:get_typeInfo()
  
  if typeInfo:get_kind() ~= _moduleObj.TypeInfoKindEnum then
    return {}, {}
  end
  local enumTypeInfo = typeInfo
  
  local val = _lune_unwrap( enumTypeInfo:getEnumValInfo( self.symbolInfo:get_name() ))
  
  return {val:get_val()}, {enumTypeInfo:get_valTypeInfo()}
end

function ExpOp2Node:getLiteral(  )

  local val1List, type1List = self:get_exp1():getLiteral(  )
  
  local val2List, type2List = self:get_exp2():getLiteral(  )
  
  if #val1List ~= 1 or #type1List ~= 1 or #val2List ~= 1 or #type2List ~= 1 then
    return {}, {}
  end
  local val1, type1, val2, type2 = _lune_unwrap( val1List[1]), type1List[1], _lune_unwrap( val2List[1]), type2List[1]
  
  if (type1 == _moduleObj.builtinTypeInt or type1 == _moduleObj.builtinTypeReal ) and (type2 == _moduleObj.builtinTypeInt or type2 == _moduleObj.builtinTypeReal ) then
    local retType = _moduleObj.builtinTypeInt
    
    if type1 == _moduleObj.builtinTypeReal or type2 == _moduleObj.builtinTypeReal then
      retType = _moduleObj.builtinTypeReal
    end
    local int1, int2 = 0, 0
    
    local real1, real2 = 0.0, 0.0
    
    if type1 == _moduleObj.builtinTypeInt then
      int1 = math.floor(val1)
      real1 = int1
    else 
      real1 = val1
    end
    if type2 == _moduleObj.builtinTypeInt then
      int2 = math.floor(val2)
      real2 = int2
    else 
      real2 = val2
    end
    do
      local _switchExp = (self.op.txt )
      if _switchExp == "+" then
        if retType == _moduleObj.builtinTypeInt then
          return {int1 + int2}, {retType}
        end
        return {real1 + real2}, {retType}
      elseif _switchExp == "-" then
        if retType == _moduleObj.builtinTypeInt then
          return {int1 - int2}, {retType}
        end
        return {real1 - real2}, {retType}
      elseif _switchExp == "*" then
        if retType == _moduleObj.builtinTypeInt then
          return {int1 * int2}, {retType}
        end
        return {real1 * real2}, {retType}
      elseif _switchExp == "/" then
        if retType == _moduleObj.builtinTypeInt then
          return {math.floor(int1 / int2)}, {retType}
        end
        return {real1 / real2}, {retType}
      end
    end
    
  elseif type1 == _moduleObj.builtinTypeString and type2 == _moduleObj.builtinTypeString then
    if self.op.txt == ".." then
      return {val1 .. val2}, {_moduleObj.builtinTypeString}
    end
  end
  return {}, {}
end

return _moduleObj
