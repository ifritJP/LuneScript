--lune/base/Ast.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Ast'
if not _lune then
   _lune = {}
end
if not table.unpack then
   table.unpack = unpack
end

function _lune.unwrap( val )
   if val == nil then
      __luneScript:error( 'unwrap val is nil' )
   end
   return val
end 
function _lune.unwrapDefault( val, defval )
   if val == nil then
      return defval
   end
   return val
end

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end

local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Code = _lune.loadModule( 'lune.base.Code' )
local IdProvider = {}
_moduleObj.IdProvider = IdProvider
function IdProvider:increment(  )

   self.id = self.id + 1
end
function IdProvider:getNewId(  )

   local newId = self.id
   self.id = self.id + 1
   return newId
end
function IdProvider.setmeta( obj )
  setmetatable( obj, { __index = IdProvider  } )
end
function IdProvider.new( id )
   local obj = {}
   IdProvider.setmeta( obj )
   if obj.__init then
      obj:__init( id )
   end        
   return obj 
end         
function IdProvider:__init( id ) 

   self.id = id
end
function IdProvider:get_id()       
   return self.id         
end

local idProv = IdProvider.new(1)
local userStartId = 1000
local rootTypeId = idProv:getNewId(  )
_moduleObj.rootTypeId = rootTypeId



local typeInfoKind = {}
_moduleObj.typeInfoKind = typeInfoKind

local sym2builtInTypeMap = {}
_moduleObj.sym2builtInTypeMap = sym2builtInTypeMap

local builtInTypeIdSet = {}
_moduleObj.builtInTypeIdSet = builtInTypeIdSet

local SerializeKind = {}
_moduleObj.SerializeKind = SerializeKind
SerializeKind._val2NameMap = {}
function SerializeKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "SerializeKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function SerializeKind._from( val )
   if SerializeKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
SerializeKind.__allList = {}
function SerializeKind.get__allList()
   return SerializeKind.__allList
end

SerializeKind.Nilable = 0
SerializeKind._val2NameMap[0] = 'Nilable'
SerializeKind.__allList[1] = SerializeKind.Nilable
SerializeKind.Modifier = 1
SerializeKind._val2NameMap[1] = 'Modifier'
SerializeKind.__allList[2] = SerializeKind.Modifier
SerializeKind.Module = 2
SerializeKind._val2NameMap[2] = 'Module'
SerializeKind.__allList[3] = SerializeKind.Module
SerializeKind.Normal = 3
SerializeKind._val2NameMap[3] = 'Normal'
SerializeKind.__allList[4] = SerializeKind.Normal
SerializeKind.Enum = 4
SerializeKind._val2NameMap[4] = 'Enum'
SerializeKind.__allList[5] = SerializeKind.Enum
SerializeKind.Alge = 5
SerializeKind._val2NameMap[5] = 'Alge'
SerializeKind.__allList[6] = SerializeKind.Alge
SerializeKind.DDD = 6
SerializeKind._val2NameMap[6] = 'DDD'
SerializeKind.__allList[7] = SerializeKind.DDD

local TypeInfoKind = {}
_moduleObj.TypeInfoKind = TypeInfoKind
TypeInfoKind._val2NameMap = {}
function TypeInfoKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "TypeInfoKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function TypeInfoKind._from( val )
   if TypeInfoKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
TypeInfoKind.__allList = {}
function TypeInfoKind.get__allList()
   return TypeInfoKind.__allList
end

TypeInfoKind.Root = 0
TypeInfoKind._val2NameMap[0] = 'Root'
TypeInfoKind.__allList[1] = TypeInfoKind.Root
TypeInfoKind.Macro = 1
TypeInfoKind._val2NameMap[1] = 'Macro'
TypeInfoKind.__allList[2] = TypeInfoKind.Macro
TypeInfoKind.Prim = 2
TypeInfoKind._val2NameMap[2] = 'Prim'
TypeInfoKind.__allList[3] = TypeInfoKind.Prim
TypeInfoKind.List = 3
TypeInfoKind._val2NameMap[3] = 'List'
TypeInfoKind.__allList[4] = TypeInfoKind.List
TypeInfoKind.Array = 4
TypeInfoKind._val2NameMap[4] = 'Array'
TypeInfoKind.__allList[5] = TypeInfoKind.Array
TypeInfoKind.Map = 5
TypeInfoKind._val2NameMap[5] = 'Map'
TypeInfoKind.__allList[6] = TypeInfoKind.Map
TypeInfoKind.Class = 6
TypeInfoKind._val2NameMap[6] = 'Class'
TypeInfoKind.__allList[7] = TypeInfoKind.Class
TypeInfoKind.IF = 7
TypeInfoKind._val2NameMap[7] = 'IF'
TypeInfoKind.__allList[8] = TypeInfoKind.IF
TypeInfoKind.Func = 8
TypeInfoKind._val2NameMap[8] = 'Func'
TypeInfoKind.__allList[9] = TypeInfoKind.Func
TypeInfoKind.Method = 9
TypeInfoKind._val2NameMap[9] = 'Method'
TypeInfoKind.__allList[10] = TypeInfoKind.Method
TypeInfoKind.Nilable = 10
TypeInfoKind._val2NameMap[10] = 'Nilable'
TypeInfoKind.__allList[11] = TypeInfoKind.Nilable
TypeInfoKind.Enum = 11
TypeInfoKind._val2NameMap[11] = 'Enum'
TypeInfoKind.__allList[12] = TypeInfoKind.Enum
TypeInfoKind.Module = 12
TypeInfoKind._val2NameMap[12] = 'Module'
TypeInfoKind.__allList[13] = TypeInfoKind.Module
TypeInfoKind.Stem = 13
TypeInfoKind._val2NameMap[13] = 'Stem'
TypeInfoKind.__allList[14] = TypeInfoKind.Stem
TypeInfoKind.Alge = 14
TypeInfoKind._val2NameMap[14] = 'Alge'
TypeInfoKind.__allList[15] = TypeInfoKind.Alge
TypeInfoKind.DDD = 15
TypeInfoKind._val2NameMap[15] = 'DDD'
TypeInfoKind.__allList[16] = TypeInfoKind.DDD
TypeInfoKind.Abbr = 16
TypeInfoKind._val2NameMap[16] = 'Abbr'
TypeInfoKind.__allList[17] = TypeInfoKind.Abbr

local function isBuiltin( typeId )

   return _moduleObj.builtInTypeIdSet[typeId] ~= nil
end
_moduleObj.isBuiltin = isBuiltin

local SymbolKind = {}
_moduleObj.SymbolKind = SymbolKind
SymbolKind._val2NameMap = {}
function SymbolKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "SymbolKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function SymbolKind._from( val )
   if SymbolKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
SymbolKind.__allList = {}
function SymbolKind.get__allList()
   return SymbolKind.__allList
end

SymbolKind.Typ = 0
SymbolKind._val2NameMap[0] = 'Typ'
SymbolKind.__allList[1] = SymbolKind.Typ
SymbolKind.Mbr = 1
SymbolKind._val2NameMap[1] = 'Mbr'
SymbolKind.__allList[2] = SymbolKind.Mbr
SymbolKind.Mtd = 2
SymbolKind._val2NameMap[2] = 'Mtd'
SymbolKind.__allList[3] = SymbolKind.Mtd
SymbolKind.Fun = 3
SymbolKind._val2NameMap[3] = 'Fun'
SymbolKind.__allList[4] = SymbolKind.Fun
SymbolKind.Var = 4
SymbolKind._val2NameMap[4] = 'Var'
SymbolKind.__allList[5] = SymbolKind.Var
SymbolKind.Arg = 5
SymbolKind._val2NameMap[5] = 'Arg'
SymbolKind.__allList[6] = SymbolKind.Arg
SymbolKind.Mac = 6
SymbolKind._val2NameMap[6] = 'Mac'
SymbolKind.__allList[7] = SymbolKind.Mac

local AccessMode = {}
_moduleObj.AccessMode = AccessMode
AccessMode._val2NameMap = {}
function AccessMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "AccessMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function AccessMode._from( val )
   if AccessMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
AccessMode.__allList = {}
function AccessMode.get__allList()
   return AccessMode.__allList
end

AccessMode.None = 0
AccessMode._val2NameMap[0] = 'None'
AccessMode.__allList[1] = AccessMode.None
AccessMode.Pub = 1
AccessMode._val2NameMap[1] = 'Pub'
AccessMode.__allList[2] = AccessMode.Pub
AccessMode.Pro = 2
AccessMode._val2NameMap[2] = 'Pro'
AccessMode.__allList[3] = AccessMode.Pro
AccessMode.Pri = 3
AccessMode._val2NameMap[3] = 'Pri'
AccessMode.__allList[4] = AccessMode.Pri
AccessMode.Local = 4
AccessMode._val2NameMap[4] = 'Local'
AccessMode.__allList[5] = AccessMode.Local
AccessMode.Global = 5
AccessMode._val2NameMap[5] = 'Global'
AccessMode.__allList[6] = AccessMode.Global

local function isPubToExternal( mode )

   do
      local _switchExp = mode
      if _switchExp == AccessMode.Pub or _switchExp == AccessMode.Pro or _switchExp == AccessMode.Global then
         return true
      end
   end
   
   return false
end
_moduleObj.isPubToExternal = isPubToExternal
local txt2AccessModeMap = {}
txt2AccessModeMap["none"] = AccessMode.None
txt2AccessModeMap["pub"] = AccessMode.Pub
txt2AccessModeMap["pro"] = AccessMode.Pro
txt2AccessModeMap["pri"] = AccessMode.Pri
txt2AccessModeMap["local"] = AccessMode.Local
txt2AccessModeMap["global"] = AccessMode.Global
local function txt2AccessMode( accessMode )

   return txt2AccessModeMap[accessMode]
end
_moduleObj.txt2AccessMode = txt2AccessMode
local SymbolInfo = {}
_moduleObj.SymbolInfo = SymbolInfo
function SymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = SymbolInfo  } )
end
function SymbolInfo.new(  )
   local obj = {}
   SymbolInfo.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function SymbolInfo:__init(  ) 

end

local NormalSymbolInfo = {}
setmetatable( NormalSymbolInfo, { __index = SymbolInfo } )
_moduleObj.NormalSymbolInfo = NormalSymbolInfo
function NormalSymbolInfo:getOrg(  )

   return self
end
function NormalSymbolInfo.new( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag )
   local obj = {}
   NormalSymbolInfo.setmeta( obj )
   if obj.__init then obj:__init( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag ); end
   return obj
end
function NormalSymbolInfo:__init(kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag) 
   SymbolInfo.__init( self )
   
   NormalSymbolInfo.symbolIdSeed = NormalSymbolInfo.symbolIdSeed + 1
   self.kind = kind
   self.canBeLeft = canBeLeft
   self.canBeRight = canBeRight
   self.symbolId = NormalSymbolInfo.symbolIdSeed
   self.scope = scope
   self.accessMode = accessMode
   self.staticFlag = staticFlag
   self.name = name
   self.typeInfo = typeInfo
   self.mutable = mutable and true or false
   self.hasValueFlag = hasValueFlag
end
function NormalSymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = NormalSymbolInfo  } )
end
function NormalSymbolInfo:get_canBeLeft()       
   return self.canBeLeft         
end
function NormalSymbolInfo:get_canBeRight()       
   return self.canBeRight         
end
function NormalSymbolInfo:get_symbolId()       
   return self.symbolId         
end
function NormalSymbolInfo:get_scope()       
   return self.scope         
end
function NormalSymbolInfo:get_accessMode()       
   return self.accessMode         
end
function NormalSymbolInfo:get_staticFlag()       
   return self.staticFlag         
end
function NormalSymbolInfo:get_name()       
   return self.name         
end
function NormalSymbolInfo:get_typeInfo()       
   return self.typeInfo         
end
function NormalSymbolInfo:get_mutable()       
   return self.mutable         
end
function NormalSymbolInfo:get_kind()       
   return self.kind         
end
function NormalSymbolInfo:get_hasValueFlag()       
   return self.hasValueFlag         
end
function NormalSymbolInfo:set_hasValueFlag( hasValueFlag )   
   self.hasValueFlag = hasValueFlag              
end
do
   NormalSymbolInfo.symbolIdSeed = 0
end

local DataOwnerInfo = {}
_moduleObj.DataOwnerInfo = DataOwnerInfo
function DataOwnerInfo.setmeta( obj )
  setmetatable( obj, { __index = DataOwnerInfo  } )
end
function DataOwnerInfo.new( hasData, symbolInfo )
   local obj = {}
   DataOwnerInfo.setmeta( obj )
   if obj.__init then
      obj:__init( hasData, symbolInfo )
   end        
   return obj 
end         
function DataOwnerInfo:__init( hasData, symbolInfo ) 

   self.hasData = hasData
   self.symbolInfo = symbolInfo
end

local Scope = {}
_moduleObj.Scope = Scope
function Scope.new( parent, classFlag, inherit, ifScopeList )
   local obj = {}
   Scope.setmeta( obj )
   if obj.__init then obj:__init( parent, classFlag, inherit, ifScopeList ); end
   return obj
end
function Scope:__init(parent, classFlag, inherit, ifScopeList) 
   self.parent = _lune.unwrapDefault( parent, self)
   self.symbol2SymbolInfoMap = {}
   self.inherit = inherit
   self.classFlag = classFlag
   self.symbolId2DataOwnerInfo = {}
   self.ifScopeList = _lune.unwrapDefault( ifScopeList, {})
end
function Scope:isRoot(  )

   return self.parent == self
end
function Scope:set_ownerTypeInfo( owner )

   self.ownerTypeInfo = owner
end
function Scope:getTypeInfoChild( name )

   do
      local _exp = self.symbol2SymbolInfoMap[name]
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   return nil
end
function Scope:getSymbolInfoChild( name )

   return self.symbol2SymbolInfoMap[name]
end
function Scope:setData( symbolInfo )

   self.symbolId2DataOwnerInfo[symbolInfo:get_symbolId()] = DataOwnerInfo.new(true, symbolInfo)
end
function Scope:remove( name )

   self.symbol2SymbolInfoMap[name] = nil
end
function Scope.setmeta( obj )
  setmetatable( obj, { __index = Scope  } )
end
function Scope:get_ownerTypeInfo()       
   return self.ownerTypeInfo         
end
function Scope:get_parent()       
   return self.parent         
end
function Scope:get_symbol2SymbolInfoMap()       
   return self.symbol2SymbolInfoMap         
end

local rootScope = Scope.new(nil, false, nil)
_moduleObj.rootScope = rootScope

function Scope:isInnerOf( scope )

   local workScope = self
   while workScope ~= _moduleObj.rootScope do
      if workScope == scope then
         return true
      end
      
      workScope = workScope.parent
   end
   
   return false
end

local dummyList = {}
local rootChildren = {}
local TypeData = {}
_moduleObj.TypeData = TypeData
function TypeData:addChildren( child )

   table.insert( self.children, child )
end
function TypeData.setmeta( obj )
  setmetatable( obj, { __index = TypeData  } )
end
function TypeData.new( children )
   local obj = {}
   TypeData.setmeta( obj )
   if obj.__init then
      obj:__init( children )
   end        
   return obj 
end         
function TypeData:__init( children ) 

   self.children = children
end
function TypeData:get_children()       
   return self.children         
end

local TypeManager = {}
function TypeManager.add( typeInfo )

   TypeManager.info2Data[typeInfo] = TypeData.new({})
end
function TypeManager.getData( typeInfo )

   return TypeManager.info2Data[typeInfo]
end
function TypeManager.setmeta( obj )
  setmetatable( obj, { __index = TypeManager  } )
end
function TypeManager.new(  )
   local obj = {}
   TypeManager.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function TypeManager:__init(  ) 

end
do
   TypeManager.info2Data = {}
end

local typeInfo2ScopeMap = {}
local function getScope( typeInfo )

   return typeInfo2ScopeMap[typeInfo]
end
_moduleObj.getScope = getScope
local ModuleInfoIF = {}
_moduleObj.ModuleInfoIF = ModuleInfoIF
function ModuleInfoIF.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfoIF  } )
end
function ModuleInfoIF.new(  )
   local obj = {}
   ModuleInfoIF.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function ModuleInfoIF:__init(  ) 

end

local TypeInfo = {}
_moduleObj.TypeInfo = TypeInfo
function TypeInfo:get_scope(  )

   return typeInfo2ScopeMap[self]
end
function TypeInfo.new( scope )
   local obj = {}
   TypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope ); end
   return obj
end
function TypeInfo:__init(scope) 
   typeInfo2ScopeMap[self] = scope
   self.scope = scope
   do
      local _exp = scope
      if _exp ~= nil then
         _exp:set_ownerTypeInfo( self )
      end
   end
   
   TypeManager.add( self )
end
function TypeInfo:isModule(  )

   return true
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
function TypeInfo:getTxt( fullName, importInfo, localFlag )

   return ""
end
function TypeInfo:canEvalWith( other, opTxt )

   return false
end
function TypeInfo:get_abstractFlag(  )

   return false
end
function TypeInfo:serialize( stream, validChildrenSet )

   return 
end
function TypeInfo:get_display_stirng(  )

   return ""
end
function TypeInfo:get_srcTypeInfo(  )

   return self
end
function TypeInfo:equals( typeInfo )

   return self == typeInfo:get_srcTypeInfo()
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
function TypeInfo:hasRouteNamespaceFrom( other )

   while true do
      if other == self then
         return true
      end
      
      if other:get_parentInfo() == other then
         break
      end
      
      other = other:get_parentInfo()
   end
   
   return false
end
function TypeInfo:getModule(  )

   if self:isModule(  ) then
      return self
   end
   
   return self:get_parentInfo():getModule(  )
end
function TypeInfo:get_rawTxt(  )

   return ""
end
function TypeInfo:get_typeId(  )

   return _moduleObj.rootTypeId
end
function TypeInfo:get_kind(  )

   return TypeInfoKind.Root
end
function TypeInfo:get_staticFlag(  )

   return false
end
function TypeInfo:get_accessMode(  )

   return AccessMode.Pri
end
function TypeInfo:get_autoFlag(  )

   return false
end
function TypeInfo:get_nonnilableType(  )

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
function TypeInfo:get_typeData(  )

   return _lune.unwrap( TypeManager.getData( self ))
end
function TypeInfo:get_children(  )

   return self:get_typeData():get_children()
end
function TypeInfo:addChildren( child )

   (_lune.unwrap( TypeManager.getData( self )) ):addChildren( child )
end
function TypeInfo:get_mutable(  )

   return true
end
function TypeInfo:getParentFullName( importInfo, localFlag )

   local typeInfo = self
   local name = ""
   local infoMap = importInfo
   if  nil == infoMap then
      local _infoMap = infoMap
   
      infoMap = {}
   end
   
   while not infoMap[typeInfo] do
      typeInfo = typeInfo:get_parentInfo()
      if typeInfo == typeInfo:get_parentInfo() then
         break
      end
      
      local txt = typeInfo:get_rawTxt()
      if localFlag then
         do
            local moduleInfo = infoMap[typeInfo]
            if moduleInfo ~= nil then
               txt = moduleInfo:get_assignName()
            else
               if typeInfo:isModule(  ) then
                  break
               end
               
            end
         end
         
      end
      
      name = txt .. "." .. name
   end
   
   return name
end
function TypeInfo.setmeta( obj )
  setmetatable( obj, { __index = TypeInfo  } )
end


function Scope:filterTypeInfoField( includeSelfFlag, fromScope, callback )

   if self.classFlag then
      if includeSelfFlag then
         for __index, symbolInfo in pairs( self.symbol2SymbolInfoMap ) do
            if symbolInfo:canAccess( fromScope ) then
               if not callback( symbolInfo ) then
                  return false
               end
               
            end
            
         end
         
      end
      
      do
         local scope = self.inherit
         if scope ~= nil then
            if not scope:filterTypeInfoField( true, fromScope, callback ) then
               return false
            end
            
         end
      end
      
   end
   
   return true
end

function Scope:getSymbolInfoField( name, includeSelfFlag, fromScope )

   if self.classFlag then
      if includeSelfFlag then
         do
            local _exp = self.symbol2SymbolInfoMap[name]
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
      
      do
         local scope = self.inherit
         if scope ~= nil then
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope )
            if symbolInfo then
               return symbolInfo
            end
            
         end
      end
      
   end
   
   return nil
end

function Scope:getSymbolInfoIfField( name, fromScope )

   if self.classFlag then
      for __index, scope in pairs( self.ifScopeList ) do
         do
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope )
            if symbolInfo ~= nil then
               return symbolInfo
            end
         end
         
      end
      
   end
   
   do
      local scope = self.inherit
      if scope ~= nil then
         do
            local symbolInfo = scope:getSymbolInfoIfField( name, fromScope )
            if symbolInfo ~= nil then
               return symbolInfo
            end
         end
         
      end
   end
   
   return nil
end

function Scope:filterSymbolInfoIfField( fromScope, callback )

   for __index, scope in pairs( self.ifScopeList ) do
      if not scope:filterTypeInfoField( true, fromScope, callback ) then
         return false
      end
      
   end
   
   do
      local scope = self.inherit
      if scope ~= nil then
         if not scope:filterSymbolInfoIfField( fromScope, callback ) then
            return false
         end
         
      end
   end
   
   return true
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

function Scope:getSymbolInfo( name, fromScope, onlySameNsFlag )

   do
      local _exp = self.symbol2SymbolInfoMap[name]
      if _exp ~= nil then
         local symbolInfo = _exp:canAccess( fromScope )
         if  nil == symbolInfo then
            local _symbolInfo = symbolInfo
         
            return nil
         end
         
         return symbolInfo
      end
   end
   
   if not onlySameNsFlag then
      do
         local scope = self.inherit
         if scope ~= nil then
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope )
            if symbolInfo then
               return symbolInfo
            end
            
         end
      end
      
   end
   
   if not onlySameNsFlag or not self.ownerTypeInfo then
      if self.parent ~= self then
         return self.parent:getSymbolInfo( name, fromScope, onlySameNsFlag )
      end
      
   end
   
   if onlySameNsFlag then
      return nil
   end
   
   do
      local _exp = _moduleObj.sym2builtInTypeMap[name]
      if _exp ~= nil then
         return _exp
      end
   end
   
   return nil
end

function Scope:getTypeInfo( name, fromScope, onlySameNsFlag )

   local symbolInfo = self:getSymbolInfo( name, fromScope, onlySameNsFlag )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      return nil
   end
   
   return symbolInfo:get_typeInfo()
end

function Scope:getSymbolTypeInfo( name, fromScope, moduleScope )

   local typeInfo = nil
   local validThisScope = false
   do
      local _exp = self.ownerTypeInfo
      if _exp ~= nil then
         if _exp:get_kind() == TypeInfoKind.Func or _exp:get_kind() == TypeInfoKind.Method or self == moduleScope or self == _moduleObj.rootScope then
            validThisScope = true
         elseif (_exp:get_kind() == TypeInfoKind.IF or _exp:get_kind() == TypeInfoKind.Class or _exp:get_kind() == TypeInfoKind.Module ) and name == "self" then
            validThisScope = true
         elseif _exp:get_kind() == TypeInfoKind.Enum or _exp:get_kind() == TypeInfoKind.Alge then
            validThisScope = true
         end
         
      else
         validThisScope = true
      end
   end
   
   if validThisScope then
      do
         local _exp = self.symbol2SymbolInfoMap[name]
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
         local _exp = self.symbol2SymbolInfoMap["self"]
         if _exp ~= nil then
            callback( _exp )
         end
      end
      
   end
   
   if moduleScope == fromScope or not self.classFlag then
      for __index, symbolInfo in pairs( self.symbol2SymbolInfoMap ) do
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

   local symbolInfo = NormalSymbolInfo.new(kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, typeInfo, mutable, hasValueFlag)
   self.symbol2SymbolInfoMap[name] = symbolInfo
   return symbolInfo
end

function Scope:addLocalVar( argFlag, canBeLeft, name, typeInfo, mutable )

   self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, AccessMode.Local, false, mutable, true )
end

function Scope:addStaticVar( argFlag, canBeLeft, name, typeInfo, mutable )

   self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, AccessMode.Local, true, mutable, true )
end

function Scope:addVar( accessMode, name, typeInfo, mutable, hasValueFlag )

   self:add( SymbolKind.Var, true, true, name, typeInfo, accessMode, false, mutable, hasValueFlag )
end

function Scope:addEnumVal( name, typeInfo )

   self:add( SymbolKind.Mbr, false, true, name, typeInfo, AccessMode.Pub, true, true, true )
end

function Scope:addEnum( accessMode, name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, true, true )
end

function Scope:addAlgeVal( name, typeInfo )

   self:add( SymbolKind.Mbr, false, true, name, typeInfo, AccessMode.Pub, true, true, true )
end

function Scope:addAlge( accessMode, name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, true, true )
end

function Scope:addMember( name, typeInfo, accessMode, staticFlag, mutable )

   return self:add( SymbolKind.Mbr, true, true, name, typeInfo, accessMode, staticFlag, mutable, true )
end

function Scope:addMethod( typeInfo, accessMode, staticFlag, mutable )

   self:add( SymbolKind.Mtd, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable, true )
end

function Scope:addFunc( typeInfo, accessMode, staticFlag, mutable )

   self:add( SymbolKind.Fun, true, true, typeInfo:getTxt(  ), typeInfo, accessMode, staticFlag, mutable, true )
end

function Scope:addMacro( typeInfo, accessMode )

   self:add( SymbolKind.Mac, false, false, typeInfo:getTxt(  ), typeInfo, accessMode, true, false, true )
end

function Scope:addClass( name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, typeInfo:get_accessMode(), true, true, true )
end

local function dumpScopeSub( scope, prefix, readyIdSet )

   do
      local _exp = scope
      if _exp ~= nil then
         if readyIdSet[_exp] then
            return 
         end
         
         readyIdSet[_exp] = true
         if #prefix > 20 then
            Util.err( "illegal" )
         end
         
         do
            local __sorted = {}
            local __map = _exp:get_symbol2SymbolInfoMap()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, symbol in ipairs( __sorted ) do
               local symbolInfo = __map[ symbol ]
               do
                  Util.log( string.format( "scope: %s, %s, %s", prefix, tostring( _exp), symbol) )
                  do
                     local subScope = symbolInfo:get_typeInfo():get_scope()
                     if subScope ~= nil then
                        dumpScopeSub( subScope, prefix .. "  ", readyIdSet )
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

local headTypeInfo = TypeInfo.new(_moduleObj.rootScope)
_moduleObj.headTypeInfo = headTypeInfo

function Scope:getNSTypeInfo(  )

   local scope = self
   while true do
      do
         local owner = scope.ownerTypeInfo
         if owner ~= nil then
            if owner:get_kind() == TypeInfoKind.Root then
               return owner
            end
            
         end
      end
      
      if scope.parent == scope then
         break
      end
      
      scope = scope.parent
   end
   
   return _moduleObj.headTypeInfo
end

function Scope:getClassTypeInfo(  )

   local scope = self
   while true do
      do
         local owner = scope.ownerTypeInfo
         if owner ~= nil then
            do
               local _switchExp = owner:get_kind()
               if _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.Module then
                  return owner
               end
            end
            
         end
      end
      
      if scope.parent == scope then
         break
      end
      
      scope = scope.parent
   end
   
   return _moduleObj.headTypeInfo
end

function NormalSymbolInfo:canAccess( fromScope )

   local typeInfo = self:get_typeInfo()
   if self.scope == fromScope then
      return self
   end
   
   do
      local _switchExp = self:get_accessMode()
      if _switchExp == AccessMode.Pub or _switchExp == AccessMode.Global then
         return self
      elseif _switchExp == AccessMode.Pro then
         local nsClass = self.scope:getClassTypeInfo(  )
         local fromClass = fromScope:getClassTypeInfo(  )
         if fromClass:isInheritFrom( nsClass ) then
            return self
         end
         
         return nil
      elseif _switchExp == AccessMode.Local then
         return self
      elseif _switchExp == AccessMode.Pri then
         if fromScope:isInnerOf( self.scope ) then
            return self
         end
         
         return nil
      end
   end
   
   Util.err( string.format( "illegl accessmode -- %s, %s", tostring( self:get_accessMode()), self:get_name()) )
end

local AccessSymbolInfo = {}
setmetatable( AccessSymbolInfo, { __index = SymbolInfo } )
_moduleObj.AccessSymbolInfo = AccessSymbolInfo
function AccessSymbolInfo:getOrg(  )

   return self.symbolInfo:getOrg(  )
end
function AccessSymbolInfo:get_mutable(  )

   do
      local _exp = self.prefixTypeInfo
      if _exp ~= nil then
         return _exp:get_mutable() and self.symbolInfo:get_mutable()
      end
   end
   
   return self.symbolInfo:get_mutable()
end
function AccessSymbolInfo:get_canBeLeft(  )

   if not self.overrideCanBeLeft then
      return false
   end
   
   do
      local _exp = self.prefixTypeInfo
      if _exp ~= nil then
         if not _exp:get_mutable() then
            return false
         end
         
      end
   end
   
   return self.symbolInfo:get_canBeLeft()
end
function AccessSymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = AccessSymbolInfo  } )
end
function AccessSymbolInfo.new( symbolInfo, prefixTypeInfo, overrideCanBeLeft )
   local obj = {}
   AccessSymbolInfo.setmeta( obj )
   if obj.__init then
      obj:__init( symbolInfo, prefixTypeInfo, overrideCanBeLeft )
   end        
   return obj 
end         
function AccessSymbolInfo:__init( symbolInfo, prefixTypeInfo, overrideCanBeLeft ) 

   SymbolInfo.__init( self )
   self.symbolInfo = symbolInfo
   self.prefixTypeInfo = prefixTypeInfo
   self.overrideCanBeLeft = overrideCanBeLeft
end
function AccessSymbolInfo:get_symbolInfo()       
   return self.symbolInfo         
end
function AccessSymbolInfo:get_prefixTypeInfo()       
   return self.prefixTypeInfo         
end
function AccessSymbolInfo:get_canBeRight( ... )
   return self.symbolInfo:get_canBeRight( ... )
end       

function AccessSymbolInfo:get_symbolId( ... )
   return self.symbolInfo:get_symbolId( ... )
end       

function AccessSymbolInfo:get_scope( ... )
   return self.symbolInfo:get_scope( ... )
end       

function AccessSymbolInfo:get_accessMode( ... )
   return self.symbolInfo:get_accessMode( ... )
end       

function AccessSymbolInfo:get_staticFlag( ... )
   return self.symbolInfo:get_staticFlag( ... )
end       

function AccessSymbolInfo:get_name( ... )
   return self.symbolInfo:get_name( ... )
end       

function AccessSymbolInfo:get_typeInfo( ... )
   return self.symbolInfo:get_typeInfo( ... )
end       

function AccessSymbolInfo:get_kind( ... )
   return self.symbolInfo:get_kind( ... )
end       

function AccessSymbolInfo:get_hasValueFlag( ... )
   return self.symbolInfo:get_hasValueFlag( ... )
end       

function AccessSymbolInfo:set_hasValueFlag( ... )
   return self.symbolInfo:set_hasValueFlag( ... )
end       

function AccessSymbolInfo:canAccess( ... )
   return self.symbolInfo:canAccess( ... )
end       




local NilableTypeInfo = {}
setmetatable( NilableTypeInfo, { __index = TypeInfo } )
_moduleObj.NilableTypeInfo = NilableTypeInfo
function NilableTypeInfo:get_kind(  )

   return TypeInfoKind.Nilable
end
function NilableTypeInfo:get_srcTypeInfo(  )

   return self
end
function NilableTypeInfo:get_nilable(  )

   return true
end
function NilableTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self.nonnilableType:getTxt( fullName, importInfo, localFlag ) .. "!"
end
function NilableTypeInfo:get_display_stirng(  )

   return self.nonnilableType:get_display_stirng() .. "!"
end
function NilableTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', SerializeKind.Nilable, parentId, self.typeId, self.nonnilableType:get_typeId()) )
end
function NilableTypeInfo:equals( typeInfo )

   if not typeInfo:get_nilable() then
      return false
   end
   
   return self.nonnilableType:equals( typeInfo )
end
function NilableTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = NilableTypeInfo  } )
end
function NilableTypeInfo.new( nonnilableType, typeId )
   local obj = {}
   NilableTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( nonnilableType, typeId )
   end        
   return obj 
end         
function NilableTypeInfo:__init( nonnilableType, typeId ) 

   TypeInfo.__init( self )
   self.nonnilableType = nonnilableType
   self.typeId = typeId
end
function NilableTypeInfo:get_nonnilableType()       
   return self.nonnilableType         
end
function NilableTypeInfo:get_typeId()       
   return self.typeId         
end
function NilableTypeInfo:get_scope( ... )
   return self.nonnilableType:get_scope( ... )
end       

function NilableTypeInfo:isModule( ... )
   return self.nonnilableType:isModule( ... )
end       

function NilableTypeInfo:getParentId( ... )
   return self.nonnilableType:getParentId( ... )
end       

function NilableTypeInfo:get_baseId( ... )
   return self.nonnilableType:get_baseId( ... )
end       

function NilableTypeInfo:isInheritFrom( ... )
   return self.nonnilableType:isInheritFrom( ... )
end       

function NilableTypeInfo:get_abstractFlag( ... )
   return self.nonnilableType:get_abstractFlag( ... )
end       

function NilableTypeInfo:get_externalFlag( ... )
   return self.nonnilableType:get_externalFlag( ... )
end       

function NilableTypeInfo:get_interfaceList( ... )
   return self.nonnilableType:get_interfaceList( ... )
end       

function NilableTypeInfo:get_itemTypeInfoList( ... )
   return self.nonnilableType:get_itemTypeInfoList( ... )
end       

function NilableTypeInfo:get_argTypeInfoList( ... )
   return self.nonnilableType:get_argTypeInfoList( ... )
end       

function NilableTypeInfo:get_retTypeInfoList( ... )
   return self.nonnilableType:get_retTypeInfoList( ... )
end       

function NilableTypeInfo:get_parentInfo( ... )
   return self.nonnilableType:get_parentInfo( ... )
end       

function NilableTypeInfo:hasRouteNamespaceFrom( ... )
   return self.nonnilableType:hasRouteNamespaceFrom( ... )
end       

function NilableTypeInfo:getModule( ... )
   return self.nonnilableType:getModule( ... )
end       

function NilableTypeInfo:get_rawTxt( ... )
   return self.nonnilableType:get_rawTxt( ... )
end       

function NilableTypeInfo:get_staticFlag( ... )
   return self.nonnilableType:get_staticFlag( ... )
end       

function NilableTypeInfo:get_accessMode( ... )
   return self.nonnilableType:get_accessMode( ... )
end       

function NilableTypeInfo:get_autoFlag( ... )
   return self.nonnilableType:get_autoFlag( ... )
end       

function NilableTypeInfo:get_baseTypeInfo( ... )
   return self.nonnilableType:get_baseTypeInfo( ... )
end       

function NilableTypeInfo:get_nilableTypeInfo( ... )
   return self.nonnilableType:get_nilableTypeInfo( ... )
end       

function NilableTypeInfo:get_typeData( ... )
   return self.nonnilableType:get_typeData( ... )
end       

function NilableTypeInfo:get_children( ... )
   return self.nonnilableType:get_children( ... )
end       

function NilableTypeInfo:addChildren( ... )
   return self.nonnilableType:addChildren( ... )
end       

function NilableTypeInfo:get_mutable( ... )
   return self.nonnilableType:get_mutable( ... )
end       

function NilableTypeInfo:getParentFullName( ... )
   return self.nonnilableType:getParentFullName( ... )
end       

function NilableTypeInfo:getFullName( ... )
   return self.nonnilableType:getFullName( ... )
end       


local ModifierTypeInfo = {}
setmetatable( ModifierTypeInfo, { __index = TypeInfo } )
_moduleObj.ModifierTypeInfo = ModifierTypeInfo
function ModifierTypeInfo:getTxt( fullName, importInfo, localFlag )

   local txt = self.srcTypeInfo:getTxt( fullName, importInfo, localFlag )
   if not self.mutable then
      txt = "&" .. txt
   end
   
   return txt
end
function ModifierTypeInfo:get_display_stirng(  )

   local txt = self.srcTypeInfo:get_display_stirng(  )
   if self.mutable then
      txt = "mut " .. txt
   end
   
   return txt
end
function ModifierTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, srcTypeId = %d, mutable = %s }\n', SerializeKind.Modifier, parentId, self.typeId, self.srcTypeInfo:get_typeId(), tostring( self.mutable and true or false)) )
end
function ModifierTypeInfo:canEvalWith( other, opTxt )

   return TypeInfo.canEvalWithBase( self.srcTypeInfo, self.mutable, other, opTxt )
end
function ModifierTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = ModifierTypeInfo  } )
end
function ModifierTypeInfo.new( srcTypeInfo, typeId, mutable )
   local obj = {}
   ModifierTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( srcTypeInfo, typeId, mutable )
   end        
   return obj 
end         
function ModifierTypeInfo:__init( srcTypeInfo, typeId, mutable ) 

   TypeInfo.__init( self )
   self.srcTypeInfo = srcTypeInfo
   self.typeId = typeId
   self.mutable = mutable
end
function ModifierTypeInfo:get_srcTypeInfo()       
   return self.srcTypeInfo         
end
function ModifierTypeInfo:get_typeId()       
   return self.typeId         
end
function ModifierTypeInfo:get_mutable()       
   return self.mutable         
end
function ModifierTypeInfo:get_scope( ... )
   return self.srcTypeInfo:get_scope( ... )
end       

function ModifierTypeInfo:isModule( ... )
   return self.srcTypeInfo:isModule( ... )
end       

function ModifierTypeInfo:getParentId( ... )
   return self.srcTypeInfo:getParentId( ... )
end       

function ModifierTypeInfo:get_baseId( ... )
   return self.srcTypeInfo:get_baseId( ... )
end       

function ModifierTypeInfo:isInheritFrom( ... )
   return self.srcTypeInfo:isInheritFrom( ... )
end       

function ModifierTypeInfo:get_abstractFlag( ... )
   return self.srcTypeInfo:get_abstractFlag( ... )
end       

function ModifierTypeInfo:equals( ... )
   return self.srcTypeInfo:equals( ... )
end       

function ModifierTypeInfo:get_externalFlag( ... )
   return self.srcTypeInfo:get_externalFlag( ... )
end       

function ModifierTypeInfo:get_interfaceList( ... )
   return self.srcTypeInfo:get_interfaceList( ... )
end       

function ModifierTypeInfo:get_itemTypeInfoList( ... )
   return self.srcTypeInfo:get_itemTypeInfoList( ... )
end       

function ModifierTypeInfo:get_argTypeInfoList( ... )
   return self.srcTypeInfo:get_argTypeInfoList( ... )
end       

function ModifierTypeInfo:get_retTypeInfoList( ... )
   return self.srcTypeInfo:get_retTypeInfoList( ... )
end       

function ModifierTypeInfo:get_parentInfo( ... )
   return self.srcTypeInfo:get_parentInfo( ... )
end       

function ModifierTypeInfo:hasRouteNamespaceFrom( ... )
   return self.srcTypeInfo:hasRouteNamespaceFrom( ... )
end       

function ModifierTypeInfo:getModule( ... )
   return self.srcTypeInfo:getModule( ... )
end       

function ModifierTypeInfo:get_rawTxt( ... )
   return self.srcTypeInfo:get_rawTxt( ... )
end       

function ModifierTypeInfo:get_kind( ... )
   return self.srcTypeInfo:get_kind( ... )
end       

function ModifierTypeInfo:get_staticFlag( ... )
   return self.srcTypeInfo:get_staticFlag( ... )
end       

function ModifierTypeInfo:get_accessMode( ... )
   return self.srcTypeInfo:get_accessMode( ... )
end       

function ModifierTypeInfo:get_autoFlag( ... )
   return self.srcTypeInfo:get_autoFlag( ... )
end       

function ModifierTypeInfo:get_baseTypeInfo( ... )
   return self.srcTypeInfo:get_baseTypeInfo( ... )
end       

function ModifierTypeInfo:get_nilable( ... )
   return self.srcTypeInfo:get_nilable( ... )
end       

function ModifierTypeInfo:get_nilableTypeInfo( ... )
   return self.srcTypeInfo:get_nilableTypeInfo( ... )
end       

function ModifierTypeInfo:get_typeData( ... )
   return self.srcTypeInfo:get_typeData( ... )
end       

function ModifierTypeInfo:get_children( ... )
   return self.srcTypeInfo:get_children( ... )
end       

function ModifierTypeInfo:addChildren( ... )
   return self.srcTypeInfo:addChildren( ... )
end       

function ModifierTypeInfo:getParentFullName( ... )
   return self.srcTypeInfo:getParentFullName( ... )
end       

function ModifierTypeInfo:getFullName( ... )
   return self.srcTypeInfo:getFullName( ... )
end       


local ModuleTypeInfo = {}
setmetatable( ModuleTypeInfo, { __index = TypeInfo } )
_moduleObj.ModuleTypeInfo = ModuleTypeInfo
function ModuleTypeInfo.new( scope, externalFlag, txt, parentInfo, typeId, mutable )
   local obj = {}
   ModuleTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope, externalFlag, txt, parentInfo, typeId, mutable ); end
   return obj
end
function ModuleTypeInfo:__init(scope, externalFlag, txt, parentInfo, typeId, mutable) 
   TypeInfo.__init( self ,scope)
   
   self.externalFlag = externalFlag
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = typeId
   self.mutable = mutable
   do
      local _exp = parentInfo
      if _exp ~= nil then
         _exp:addChildren( self )
      end
   end
   
   idProv:increment(  )
   scope:set_ownerTypeInfo( self )
end
function ModuleTypeInfo:isModule(  )

   return true
end
function ModuleTypeInfo:get_accessMode(  )

   return AccessMode.Pub
end
function ModuleTypeInfo:get_kind(  )

   return TypeInfoKind.Module
end
function ModuleTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function ModuleTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self.rawTxt
end
function ModuleTypeInfo:get_display_stirng(  )

   return self:getTxt(  )
end
function ModuleTypeInfo:canEvalWith( other, opTxt )

   return false
end
function ModuleTypeInfo:serialize( stream, validChildrenSet )

   local txt = string.format( "{ skind = %d, parentId = %d, typeId = %d, txt = '%s', kind = %d, ", SerializeKind.Module, self:getParentId(  ), self.typeId, self.rawTxt, TypeInfoKind.Module)
   stream:write( txt .. '\n' )
   stream:write( "children = {" )
   local set = validChildrenSet
   if  nil == set then
      local _set = set
   
      set = {}
   end
   
   do
      local _exp = validChildrenSet
      if _exp ~= nil then
         for __index, child in pairs( self:get_children() ) do
            if set[child:get_typeId()] and (child:get_accessMode() == AccessMode.Pub or child:get_accessMode() == AccessMode.Global ) then
               stream:write( string.format( "%d, ", child:get_typeId()) )
            end
            
         end
         
      end
   end
   
   stream:write( "} }\n" )
end
function ModuleTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleTypeInfo  } )
end
function ModuleTypeInfo:get_externalFlag()       
   return self.externalFlag         
end
function ModuleTypeInfo:get_parentInfo()       
   return self.parentInfo         
end
function ModuleTypeInfo:get_typeId()       
   return self.typeId         
end
function ModuleTypeInfo:get_rawTxt()       
   return self.rawTxt         
end
function ModuleTypeInfo:get_mutable()       
   return self.mutable         
end

local EnumValInfo = {}
_moduleObj.EnumValInfo = EnumValInfo
function EnumValInfo.setmeta( obj )
  setmetatable( obj, { __index = EnumValInfo  } )
end
function EnumValInfo.new( name, val )
   local obj = {}
   EnumValInfo.setmeta( obj )
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

local EnumTypeInfo = {}
setmetatable( EnumTypeInfo, { __index = TypeInfo } )
_moduleObj.EnumTypeInfo = EnumTypeInfo
function EnumTypeInfo.new( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo )
   local obj = {}
   EnumTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo ); end
   return obj
end
function EnumTypeInfo:__init(scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo, name2EnumValInfo) 
   TypeInfo.__init( self ,scope)
   
   self.externalFlag = externalFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = typeId
   self.name2EnumValInfo = name2EnumValInfo
   self.valTypeInfo = valTypeInfo
   do
      local _exp = parentInfo
      if _exp ~= nil then
         _exp:addChildren( self )
      end
   end
   
   self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
   idProv:increment(  )
   scope:set_ownerTypeInfo( self )
end
function EnumTypeInfo:isModule(  )

   return false
end
function EnumTypeInfo:get_kind(  )

   return TypeInfoKind.Enum
end
function EnumTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function EnumTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self.rawTxt
end
function EnumTypeInfo:get_display_stirng(  )

   return self:getTxt(  )
end
function EnumTypeInfo:canEvalWith( other, opTxt )

   return self == other:get_srcTypeInfo()
end
function EnumTypeInfo:getEnumValInfo( name )

   return self.name2EnumValInfo[name]
end
function EnumTypeInfo:get_mutable(  )

   return true
end
function EnumTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = EnumTypeInfo  } )
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

local AlgeValInfo = {}
_moduleObj.AlgeValInfo = AlgeValInfo
function AlgeValInfo:serialize( stream )

   stream:write( string.format( "{ name = '%s', typeList = {", self.name) )
   do
      local __sorted = {}
      local __map = self.typeList
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, index in ipairs( __sorted ) do
         local typeInfo = __map[ index ]
         do
            if index > 1 then
               stream:write( ", " )
            end
            
            stream:write( string.format( "%d", typeInfo:get_typeId()) )
         end
      end
   end
   
   stream:write( "} }" )
end
function AlgeValInfo.setmeta( obj )
  setmetatable( obj, { __index = AlgeValInfo  } )
end
function AlgeValInfo.new( name, typeList )
   local obj = {}
   AlgeValInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeList )
   end        
   return obj 
end         
function AlgeValInfo:__init( name, typeList ) 

   self.name = name
   self.typeList = typeList
end
function AlgeValInfo:get_name()       
   return self.name         
end
function AlgeValInfo:get_typeList()       
   return self.typeList         
end

local AlgeTypeInfo = {}
setmetatable( AlgeTypeInfo, { __index = TypeInfo } )
_moduleObj.AlgeTypeInfo = AlgeTypeInfo
function AlgeTypeInfo.new( scope, externalFlag, accessMode, txt, parentInfo, typeId )
   local obj = {}
   AlgeTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope, externalFlag, accessMode, txt, parentInfo, typeId ); end
   return obj
end
function AlgeTypeInfo:__init(scope, externalFlag, accessMode, txt, parentInfo, typeId) 
   TypeInfo.__init( self ,scope)
   
   self.externalFlag = externalFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = typeId
   self.valInfoMap = {}
   do
      local _exp = parentInfo
      if _exp ~= nil then
         _exp:addChildren( self )
      end
   end
   
   self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
   idProv:increment(  )
   scope:set_ownerTypeInfo( self )
end
function AlgeTypeInfo:addValInfo( valInfo )

   self.valInfoMap[valInfo:get_name()] = valInfo
end
function AlgeTypeInfo:getValInfo( name )

   return self.valInfoMap[name]
end
function AlgeTypeInfo:isModule(  )

   return false
end
function AlgeTypeInfo:get_kind(  )

   return TypeInfoKind.Alge
end
function AlgeTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function AlgeTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self.rawTxt
end
function AlgeTypeInfo:get_display_stirng(  )

   return self:getTxt(  )
end
function AlgeTypeInfo:canEvalWith( other, opTxt )

   return self == other:get_srcTypeInfo()
end
function AlgeTypeInfo:get_mutable(  )

   return true
end
function AlgeTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AlgeTypeInfo  } )
end
function AlgeTypeInfo:get_externalFlag()       
   return self.externalFlag         
end
function AlgeTypeInfo:get_parentInfo()       
   return self.parentInfo         
end
function AlgeTypeInfo:get_typeId()       
   return self.typeId         
end
function AlgeTypeInfo:get_rawTxt()       
   return self.rawTxt         
end
function AlgeTypeInfo:get_accessMode()       
   return self.accessMode         
end
function AlgeTypeInfo:get_nilableTypeInfo()       
   return self.nilableTypeInfo         
end
function AlgeTypeInfo:get_valInfoMap()       
   return self.valInfoMap         
end

local NormalTypeInfo = {}
setmetatable( NormalTypeInfo, { __index = TypeInfo } )
_moduleObj.NormalTypeInfo = NormalTypeInfo
function NormalTypeInfo.new( abstractFlag, scope, baseTypeInfo, interfaceList, nonnilableType, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutable )
   local obj = {}
   NormalTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( abstractFlag, scope, baseTypeInfo, interfaceList, nonnilableType, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutable ); end
   return obj
end
function NormalTypeInfo:__init(abstractFlag, scope, baseTypeInfo, interfaceList, nonnilableType, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutable) 
   TypeInfo.__init( self ,scope)
   
   if type( kind ) ~= "number" then
      Util.printStackTrace(  )
   end
   
   self.abstractFlag = abstractFlag
   self.baseTypeInfo = _lune.unwrapDefault( baseTypeInfo, _moduleObj.headTypeInfo)
   self.interfaceList = _lune.unwrapDefault( interfaceList, {})
   self.autoFlag = autoFlag
   self.externalFlag = externalFlag
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.kind = kind
   self.itemTypeInfoList = _lune.unwrapDefault( itemTypeInfoList, {})
   self.argTypeInfoList = _lune.unwrapDefault( argTypeInfoList, {})
   self.retTypeInfoList = _lune.unwrapDefault( retTypeInfoList, {})
   self.nonnilableType = _lune.unwrapDefault( nonnilableType, _moduleObj.headTypeInfo)
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.mutable = mutable and true or false
   self.typeId = typeId
   if kind == TypeInfoKind.Root then
      self.nilable = false
   elseif txt == "nil" or txt == "..." then
      self.nilable = true
      self.nilableTypeInfo = self
      self.nonnilableType = self
   elseif not nonnilableType then
      do
         local _exp = parentInfo
         if _exp ~= nil then
            _exp:addChildren( self )
         end
      end
      
      self.nilable = false
      local hasNilable = false
      do
         local _switchExp = (kind )
         if _switchExp == TypeInfoKind.Prim or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Map or _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Stem or _switchExp == TypeInfoKind.Module or _switchExp == TypeInfoKind.IF then
            hasNilable = true
         elseif _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method then
            hasNilable = true
         end
      end
      
      if hasNilable then
         if txt == "..." then
            self.nilableTypeInfo = self
         else
          
            self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
            idProv:increment(  )
         end
         
      else
       
         self.nilableTypeInfo = _moduleObj.headTypeInfo
      end
      
      idProv:increment(  )
   else
    
      self.nilable = true
      self.nilableTypeInfo = _moduleObj.headTypeInfo
   end
   
end
function NormalTypeInfo:isModule(  )

   return false
end
function NormalTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId() or _moduleObj.rootTypeId
end
function NormalTypeInfo:get_baseId(  )

   return self.baseTypeInfo:get_typeId() or _moduleObj.rootTypeId
end
function NormalTypeInfo:getTxt( fullName, importInfo, localFlag )

   local parentTxt = ""
   if fullName then
      parentTxt = self:getParentFullName( importInfo, localFlag )
   end
   
   if self.nilable and (self.nilableTypeInfo ~= self.nonnilableType ) then
      return parentTxt .. (_lune.unwrap( self.nonnilableType) ):getTxt( fullName, importInfo, localFlag ) .. "!"
   end
   
   if #self.itemTypeInfoList > 0 then
      local txt = self.rawTxt .. "<"
      for index, typeInfo in pairs( self.itemTypeInfoList ) do
         if index ~= 1 then
            txt = txt .. ","
         end
         
         txt = txt .. typeInfo:getTxt( fullName, importInfo, localFlag )
      end
      
      return parentTxt .. txt .. ">"
   end
   
   return parentTxt .. self:get_rawTxt()
end
function NormalTypeInfo:get_display_stirng(  )

   if self.kind == TypeInfoKind.Nilable then
      return (_lune.unwrap( self.nonnilableType) ):get_display_stirng(  ) .. "!"
   end
   
   if self.kind == TypeInfoKind.Func or self.kind == TypeInfoKind.Method then
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
      stream:write( string.format( '{ parentId=%d, typeId = %d, nilable = true, orgTypeId = %d }\n', parentId, self.typeId, self.nonnilableType:get_typeId()) )
      return 
   end
   
   local function serializeTypeInfoList( name, list, onlyPub )
   
      local work = name
      for __index, typeInfo in pairs( list ) do
         if not onlyPub or typeInfo:get_accessMode() == AccessMode.Pub then
            if #work ~= #name then
               work = work .. ", "
            end
            
            work = string.format( "%s%d", work, typeInfo:get_typeId())
         end
         
      end
      
      return work .. "}, "
   end
   
   local txt = string.format( [==[{ skind=%d, parentId = %d, typeId = %d, baseId = %d, txt = '%s',
        abstractFlag = %s, staticFlag = %s, accessMode = %d, kind = %d, mutable = %s, ]==], SerializeKind.Normal, parentId, self.typeId, self:get_baseId(  ), self.rawTxt, tostring( self.abstractFlag), tostring( self.staticFlag), self.accessMode, self.kind, tostring( self.mutable))
   local children = {}
   local set = validChildrenSet
   if  nil == set then
      local _set = set
   
      set = {}
   end
   
   for __index, child in pairs( self:get_children() ) do
      if set[child:get_typeId()] then
         table.insert( children, child )
      end
      
   end
   
   stream:write( txt .. serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. serializeTypeInfoList( "ifList = {", self.interfaceList ) .. serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. serializeTypeInfoList( "children = {", children, true ) .. "}\n" )
end
function NormalTypeInfo:equalsSub( typeInfo )

   typeInfo = typeInfo:get_srcTypeInfo()
   if self.typeId == typeInfo:get_typeId() then
      return true
   end
   
   if self.kind ~= typeInfo:get_kind() or self.staticFlag ~= typeInfo:get_staticFlag() or self.accessMode ~= typeInfo:get_accessMode() or self.autoFlag ~= typeInfo:get_autoFlag() or self.nilable ~= typeInfo:get_nilable() or self.rawTxt ~= typeInfo:get_rawTxt() or self.parentInfo ~= typeInfo:get_parentInfo() or self.baseTypeInfo ~= typeInfo:get_baseTypeInfo() or self ~= typeInfo:get_srcTypeInfo() then
      return false
   end
   
   if (self.nonnilableType ~= typeInfo:get_nonnilableType() ) then
      Util.log( string.format( "%s, %s", tostring( self.nonnilableType), tostring( typeInfo:get_nonnilableType())) )
      return false
   end
   
   do
      if #self.itemTypeInfoList ~= #typeInfo:get_itemTypeInfoList() then
         return false
      end
      
      for index, item in pairs( self.itemTypeInfoList ) do
         if not item:equals( typeInfo:get_itemTypeInfoList()[index] ) then
            return false
         end
         
      end
      
   end
   
   do
      if #self.retTypeInfoList ~= #typeInfo:get_retTypeInfoList() then
         return false
      end
      
      for index, item in pairs( self.retTypeInfoList ) do
         if not item:equals( typeInfo:get_retTypeInfoList()[index] ) then
            return false
         end
         
      end
      
   end
   
   if self.nonnilableType ~= _moduleObj.headTypeInfo and not self.nonnilableType:equals( typeInfo:get_nonnilableType() ) then
      error( string.format( "illegal %s:%d %s:%d", self:getTxt(  ), self.typeId, typeInfo:getTxt(  ), typeInfo:get_typeId()) )
   end
   
   return true
end
function NormalTypeInfo:equals( typeInfo )

   return self:equalsSub( typeInfo )
end
function NormalTypeInfo.create( accessMode, abstractFlag, scope, baseInfo, interfaceList, parentInfo, staticFlag, kind, txt, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutable )

   if kind == TypeInfoKind.Prim then
      do
         local _exp = _moduleObj.sym2builtInTypeMap[txt]
         if _exp ~= nil then
            return _exp:get_typeInfo()
         end
      end
      
      Util.err( string.format( "not found symbol -- %s", txt) )
   end
   
   idProv:increment(  )
   local info = NormalTypeInfo.new(abstractFlag, scope, baseInfo, interfaceList, nil, false, true, staticFlag, accessMode, txt, parentInfo, idProv:get_id(), kind, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutable)
   return info
end
function NormalTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = NormalTypeInfo  } )
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
function NormalTypeInfo:get_abstractFlag()       
   return self.abstractFlag         
end
function NormalTypeInfo:get_nonnilableType()       
   return self.nonnilableType         
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
function NormalTypeInfo:get_mutable()       
   return self.mutable         
end

idProv:increment(  )
function NormalTypeInfo.createBuiltin( idName, typeTxt, kind, typeDDD )

   local typeId = idProv:get_id() + 1
   if kind == TypeInfoKind.Root then
      typeId = _moduleObj.rootTypeId
   else
    
      idProv:increment(  )
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
      if _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Module or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method or _switchExp == TypeInfoKind.Macro then
         scope = Scope.new(_moduleObj.rootScope, kind == TypeInfoKind.Class or kind == TypeInfoKind.Module or kind == TypeInfoKind.IF or kind == TypeInfoKind.List or kind == TypeInfoKind.Array, nil)
      end
   end
   
   local info = NormalTypeInfo.new(false, scope, nil, nil, nil, false, false, false, AccessMode.Pub, typeTxt, _moduleObj.headTypeInfo, typeId, kind, {}, argTypeList, retTypeList, true)
   if scope then
      _moduleObj.rootScope:addClass( typeTxt, info )
   end
   
   _moduleObj.typeInfoKind[idName] = info
   _moduleObj.sym2builtInTypeMap[typeTxt] = NormalSymbolInfo.new(SymbolKind.Typ, false, false, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, info, false, true)
   if info:get_nilableTypeInfo() ~= _moduleObj.headTypeInfo then
      _moduleObj.sym2builtInTypeMap[typeTxt .. "!"] = NormalSymbolInfo.new(SymbolKind.Typ, false, kind == TypeInfoKind.Func, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, info:get_nilableTypeInfo(), false, true)
      _moduleObj.builtInTypeIdSet[info:get_nilableTypeInfo():get_typeId()] = info:get_nilableTypeInfo()
   end
   
   _moduleObj.builtInTypeIdSet[info.typeId] = info
   return info
end

function NormalTypeInfo.createList( accessMode, parentInfo, itemTypeInfo )

   if #itemTypeInfo == 0 then
      Util.err( string.format( "illegal list type: %s", tostring( itemTypeInfo)) )
   end
   
   idProv:increment(  )
   return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "List", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.List, itemTypeInfo, nil, nil, true)
end

function NormalTypeInfo.createArray( accessMode, parentInfo, itemTypeInfo )

   idProv:increment(  )
   return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "Array", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.Array, itemTypeInfo, nil, nil, true)
end

function NormalTypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo )

   idProv:increment(  )
   return NormalTypeInfo.new(false, nil, nil, nil, nil, false, false, false, accessMode, "Map", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.Map, {keyTypeInfo, valTypeInfo}, nil, nil, true)
end

function NormalTypeInfo.createModule( scope, parentInfo, externalFlag, moduleName, mutable )

   do
      local _exp = _moduleObj.sym2builtInTypeMap[moduleName]
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   if Parser.isLuaKeyword( moduleName ) then
      Util.err( string.format( "This symbol can not use for a class or script file. -- %s", moduleName) )
   end
   
   idProv:increment(  )
   local info = ModuleTypeInfo.new(scope, externalFlag, moduleName, parentInfo, idProv:get_id(), mutable)
   return info
end

function NormalTypeInfo.createClass( classFlag, abstractFlag, scope, baseInfo, interfaceList, parentInfo, externalFlag, accessMode, className )

   do
      local _exp = _moduleObj.sym2builtInTypeMap[className]
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   if Parser.isLuaKeyword( className ) then
      Util.err( string.format( "This symbol can not use for a class or script file. -- %s", className) )
   end
   
   idProv:increment(  )
   local info = NormalTypeInfo.new(abstractFlag, scope, baseInfo, interfaceList, nil, false, externalFlag, false, accessMode, className, parentInfo, idProv:get_id(), classFlag and TypeInfoKind.Class or TypeInfoKind.IF, nil, nil, nil, true)
   return info
end

function NormalTypeInfo.createFunc( abstractFlag, builtinFlag, scope, kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList, mutable )

   if not builtinFlag and Parser.isLuaKeyword( funcName ) then
      Util.err( string.format( "This symbol can not use for a function. -- %s", funcName) )
   end
   
   idProv:increment(  )
   local info = NormalTypeInfo.new(abstractFlag, scope, nil, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, idProv:get_id(), kind, {}, _lune.unwrapDefault( argTypeList, {}), _lune.unwrapDefault( retTypeInfoList, {}), mutable)
   return info
end

function NormalTypeInfo.createAdvertiseMethodFrom( classTypeInfo, typeInfo )

   return NormalTypeInfo.createFunc( false, false, typeInfo:get_scope(), typeInfo:get_kind(), classTypeInfo, true, false, false, typeInfo:get_accessMode(), typeInfo:get_rawTxt(), typeInfo:get_argTypeInfoList(), typeInfo:get_retTypeInfoList(), typeInfo:get_mutable() )
end

local typeInfo2ModifierMap = {}
function NormalTypeInfo.createModifier( srcTypeInfo, mutable )

   srcTypeInfo = srcTypeInfo:get_srcTypeInfo()
   do
      local _exp = typeInfo2ModifierMap[srcTypeInfo]
      if _exp ~= nil then
         if _exp:get_typeId() < userStartId and srcTypeInfo:get_typeId() >= userStartId then
            Util.err( "on cache" )
         end
         
         return _exp
      end
   end
   
   idProv:increment(  )
   local modifier = ModifierTypeInfo.new(srcTypeInfo, idProv:get_id(), mutable)
   typeInfo2ModifierMap[srcTypeInfo] = modifier
   if modifier:get_typeId() < userStartId and srcTypeInfo:get_typeId() >= userStartId then
      Util.printStackTrace(  )
      Util.err( string.format( "off cache: %s %s %s", srcTypeInfo:getTxt(  ), tostring( modifier:get_typeId()), tostring( srcTypeInfo:get_typeId())) )
   end
   
   return modifier
end

function ModifierTypeInfo:get_nonnilableType(  )

   local orgType = self.srcTypeInfo:get_nonnilableType()
   if self.mutable or not orgType:get_mutable() then
      return orgType
   end
   
   return NormalTypeInfo.createModifier( orgType, false )
end

local builtinTypeNone = NormalTypeInfo.createBuiltin( "None", "", TypeInfoKind.Prim )
_moduleObj.builtinTypeNone = builtinTypeNone

local builtinTypeNeverRet = NormalTypeInfo.createBuiltin( "Error", "__", TypeInfoKind.Prim )
_moduleObj.builtinTypeNeverRet = builtinTypeNeverRet

local builtinTypeStem = NormalTypeInfo.createBuiltin( "Stem", "stem", TypeInfoKind.Stem )
_moduleObj.builtinTypeStem = builtinTypeStem

local builtinTypeNil = NormalTypeInfo.createBuiltin( "Nil", "nil", TypeInfoKind.Prim )
_moduleObj.builtinTypeNil = builtinTypeNil

local builtinTypeDDD = NormalTypeInfo.createBuiltin( "DDD", "...", TypeInfoKind.DDD )
_moduleObj.builtinTypeDDD = builtinTypeDDD

local builtinTypeBool = NormalTypeInfo.createBuiltin( "Bool", "bool", TypeInfoKind.Prim )
_moduleObj.builtinTypeBool = builtinTypeBool

local builtinTypeInt = NormalTypeInfo.createBuiltin( "Int", "int", TypeInfoKind.Prim )
_moduleObj.builtinTypeInt = builtinTypeInt

local builtinTypeReal = NormalTypeInfo.createBuiltin( "Real", "real", TypeInfoKind.Prim )
_moduleObj.builtinTypeReal = builtinTypeReal

local builtinTypeChar = NormalTypeInfo.createBuiltin( "char", "char", TypeInfoKind.Prim )
_moduleObj.builtinTypeChar = builtinTypeChar

local builtinTypeString = NormalTypeInfo.createBuiltin( "String", "str", TypeInfoKind.Class )
_moduleObj.builtinTypeString = builtinTypeString

local builtinTypeMap = NormalTypeInfo.createBuiltin( "Map", "Map", TypeInfoKind.Map )
_moduleObj.builtinTypeMap = builtinTypeMap

local builtinTypeList = NormalTypeInfo.createBuiltin( "List", "List", TypeInfoKind.List )
_moduleObj.builtinTypeList = builtinTypeList

local builtinTypeArray = NormalTypeInfo.createBuiltin( "Array", "Array", TypeInfoKind.Array )
_moduleObj.builtinTypeArray = builtinTypeArray

local builtinTypeForm = NormalTypeInfo.createBuiltin( "Form", "form", TypeInfoKind.Func, _moduleObj.builtinTypeDDD )
_moduleObj.builtinTypeForm = builtinTypeForm

local builtinTypeSymbol = NormalTypeInfo.createBuiltin( "Symbol", "sym", TypeInfoKind.Prim )
_moduleObj.builtinTypeSymbol = builtinTypeSymbol

local builtinTypeStat = NormalTypeInfo.createBuiltin( "Stat", "stat", TypeInfoKind.Prim )
_moduleObj.builtinTypeStat = builtinTypeStat

local builtinTypeStem_ = _lune.unwrap( _moduleObj.builtinTypeStem:get_nilableTypeInfo())
_moduleObj.builtinTypeStem_ = builtinTypeStem_

local AbbrTypeInfo = {}
setmetatable( AbbrTypeInfo, { __index = TypeInfo } )
_moduleObj.AbbrTypeInfo = AbbrTypeInfo
function AbbrTypeInfo:get_scope(  )

   return nil
end
function AbbrTypeInfo.new( idProvider, rawTxt )
   local obj = {}
   AbbrTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( idProvider, rawTxt ); end
   return obj
end
function AbbrTypeInfo:__init(idProvider, rawTxt) 
   TypeInfo.__init( self ,nil)
   
   local typeId = idProvider:get_id() + 1
   idProvider:increment(  )
   self.typeId = typeId
   self.rawTxt = rawTxt
end
function AbbrTypeInfo:isModule(  )

   return false
end
function AbbrTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self.rawTxt
end
function AbbrTypeInfo:canEvalWith( other, opTxt )

   return false
end
function AbbrTypeInfo:serialize( stream, validChildrenSet )

   Util.err( "illegal call" )
end
function AbbrTypeInfo:get_display_stirng(  )

   return self:getTxt(  )
end
function AbbrTypeInfo:getModule(  )

   return _moduleObj.headTypeInfo
end
function AbbrTypeInfo:get_rawTxt(  )

   return self:getTxt(  )
end
function AbbrTypeInfo:get_kind(  )

   return TypeInfoKind.Abbr
end
function AbbrTypeInfo:get_nilable(  )

   return true
end
function AbbrTypeInfo:get_nilableTypeInfo(  )

   return self
end
function AbbrTypeInfo:get_mutable(  )

   return false
end
function AbbrTypeInfo:get_accessMode(  )

   return AccessMode.Local
end
function AbbrTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AbbrTypeInfo  } )
end
function AbbrTypeInfo:get_typeId()       
   return self.typeId         
end

local builtinTypeAbbr = AbbrTypeInfo.new(idProv, "##")
_moduleObj.builtinTypeAbbr = builtinTypeAbbr

local builtinTypeAbbrNone = AbbrTypeInfo.new(idProv, "[##]")
_moduleObj.builtinTypeAbbrNone = builtinTypeAbbrNone


local typeInfo2DDDMap = {}
local DDDTypeInfo = {}
setmetatable( DDDTypeInfo, { __index = TypeInfo } )
_moduleObj.DDDTypeInfo = DDDTypeInfo
function DDDTypeInfo:get_scope(  )

   return nil
end
function DDDTypeInfo.new( typeId, typeInfo, externalFlag )
   local obj = {}
   DDDTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( typeId, typeInfo, externalFlag ); end
   return obj
end
function DDDTypeInfo:__init(typeId, typeInfo, externalFlag) 
   TypeInfo.__init( self ,nil)
   
   self.typeId = typeId
   self.typeInfo = typeInfo
   self.externalFlag = externalFlag
   self.itemTypeInfoList = {self.typeInfo}
   typeInfo2DDDMap[typeInfo] = self
end
function DDDTypeInfo:isModule(  )

   return false
end
function DDDTypeInfo:getTxt( fullName, importInfo, localFlag )

   return "...<" .. self.typeInfo:getTxt( fullName, importInfo, localFlag ) .. ">"
end
function DDDTypeInfo:canEvalWith( other, opTxt )

   return self.typeInfo:canEvalWith( other, opTxt )
end
function DDDTypeInfo:serialize( stream, validChildrenSet )

   stream:write( string.format( '{ skind=%d, typeId = %d, itemTypeId = %d, parentId = %d }\n', SerializeKind.DDD, self.typeId, self.typeInfo:get_typeId(), _moduleObj.headTypeInfo:get_typeId()) )
end
function DDDTypeInfo:get_display_stirng(  )

   return self:getTxt(  )
end
function DDDTypeInfo:getModule(  )

   return self.typeInfo:getModule(  )
end
function DDDTypeInfo:get_rawTxt(  )

   return self:getTxt(  )
end
function DDDTypeInfo:get_kind(  )

   return TypeInfoKind.DDD
end
function DDDTypeInfo:get_nilable(  )

   return self.typeInfo:get_nilable()
end
function DDDTypeInfo:get_nilableTypeInfo(  )

   return self
end
function DDDTypeInfo:get_mutable(  )

   return self.typeInfo:get_mutable()
end
function DDDTypeInfo:get_accessMode(  )

   return AccessMode.Pub
end
function DDDTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = DDDTypeInfo  } )
end
function DDDTypeInfo:get_typeInfo()       
   return self.typeInfo         
end
function DDDTypeInfo:get_typeId()       
   return self.typeId         
end
function DDDTypeInfo:get_externalFlag()       
   return self.externalFlag         
end
function DDDTypeInfo:get_itemTypeInfoList()       
   return self.itemTypeInfoList         
end

function NormalTypeInfo.createDDD( typeInfo, externalFlag )

   do
      local _exp = typeInfo2DDDMap[typeInfo]
      if _exp ~= nil then
         if _exp:get_typeId() < userStartId and typeInfo:get_typeId() >= userStartId then
            Util.err( "on cache" )
         end
         
         return _exp
      end
   end
   
   idProv:increment(  )
   return DDDTypeInfo.new(idProv:get_id(), typeInfo, externalFlag)
end

local numberTypeMap = {}
numberTypeMap[_moduleObj.builtinTypeInt] = true
numberTypeMap[_moduleObj.builtinTypeChar] = true
numberTypeMap[_moduleObj.builtinTypeReal] = true
local function isNumberType( typeInfo )

   return numberTypeMap[typeInfo:get_srcTypeInfo()] and true or false
end
_moduleObj.isNumberType = isNumberType
function NormalTypeInfo.createEnum( scope, parentInfo, externalFlag, accessMode, enumName, valTypeInfo, name2EnumValInfo )

   if Parser.isLuaKeyword( enumName ) then
      Util.err( string.format( "This symbol can not use for a enum. -- %s", enumName) )
   end
   
   idProv:increment(  )
   local info = EnumTypeInfo.new(scope, externalFlag, accessMode, enumName, parentInfo, idProv:get_id(), valTypeInfo, name2EnumValInfo)
   local getEnumName = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Method, info, true, true, false, AccessMode.Pub, "get__txt", nil, {_moduleObj.builtinTypeString}, false )
   scope:addMethod( getEnumName, AccessMode.Pub, false, false )
   local fromVal = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Func, info, true, true, true, AccessMode.Pub, "_from", {NormalTypeInfo.createModifier( valTypeInfo, false )}, {info:get_nilableTypeInfo()}, false )
   scope:addFunc( fromVal, AccessMode.Pub, true, false )
   local allListType = NormalTypeInfo.createList( AccessMode.Pub, info, {info} )
   local allList = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Func, info, true, true, true, AccessMode.Pub, "get__allList", {}, {NormalTypeInfo.createModifier( allListType, false )}, false )
   scope:addFunc( allList, AccessMode.Pub, true, false )
   return info
end

function EnumTypeInfo:serialize( stream, validChildrenSet )

   local txt = string.format( [==[{ skind = %d, parentId = %d, typeId = %d, txt = '%s',
accessMode = %d, kind = %d, valTypeId = %d, ]==], SerializeKind.Enum, self:getParentId(  ), self.typeId, self.rawTxt, self.accessMode, TypeInfoKind.Enum, self.valTypeInfo:get_typeId())
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
         local enumValInfo = __map[ __key ]
         do
            if self.valTypeInfo:equals( _moduleObj.builtinTypeString ) then
               stream:write( string.format( "%s = '%s',", enumValInfo:get_name(), tostring( enumValInfo:get_val())) )
            else
             
               stream:write( string.format( "%s = %s,", enumValInfo:get_name(), tostring( enumValInfo:get_val())) )
            end
            
         end
      end
   end
   
   stream:write( "} }\n" )
end

function NormalTypeInfo.createAlge( scope, parentInfo, externalFlag, accessMode, algeName )

   if Parser.isLuaKeyword( algeName ) then
      Util.err( string.format( "This symbol can not use for a alge. -- %s", algeName) )
   end
   
   idProv:increment(  )
   local info = AlgeTypeInfo.new(scope, externalFlag, accessMode, algeName, parentInfo, idProv:get_id())
   local getAlgeName = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Method, info, true, true, false, AccessMode.Pub, "get__txt", nil, {_moduleObj.builtinTypeString}, false )
   scope:addMethod( getAlgeName, AccessMode.Pub, false, false )
   return info
end

function AlgeTypeInfo:serialize( stream, validChildrenSet )

   local txt = string.format( [==[{ skind = %d, parentId = %d, typeId = %d, txt = '%s',
accessMode = %d, kind = %d, ]==], SerializeKind.Alge, self:getParentId(  ), self.typeId, self.rawTxt, self.accessMode, TypeInfoKind.Alge)
   stream:write( txt )
   stream:write( "algeValList = {" )
   local firstFlag = true
   do
      local __sorted = {}
      local __map = self.valInfoMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, __key in ipairs( __sorted ) do
         local algeValInfo = __map[ __key ]
         do
            if not firstFlag then
               stream:write( "," )
            else
             
               firstFlag = false
            end
            
            algeValInfo:serialize( stream )
         end
      end
   end
   
   stream:write( "} }\n" )
end

function NilableTypeInfo:canEvalWith( other, opTxt )

   local otherSrc = other:get_srcTypeInfo()
   if self == _moduleObj.builtinTypeStem_ then
      return true
   end
   
   if otherSrc == _moduleObj.builtinTypeNil or otherSrc:get_kind() == TypeInfoKind.Abbr then
      return true
   end
   
   if self.typeId == otherSrc:get_typeId() then
      return true
   end
   
   if otherSrc:get_nilable() then
      return self:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt )
   end
   
   return self:get_nonnilableType():canEvalWith( otherSrc, opTxt )
end



function NormalTypeInfo:isInheritFrom( other )

   local otherTypeId = other:get_typeId()
   if self:get_typeId() == otherTypeId then
      return true
   end
   
   if (self:get_kind() ~= TypeInfoKind.Class and self:get_kind() ~= TypeInfoKind.IF ) or (other:get_kind() ~= TypeInfoKind.Class and other:get_kind() ~= TypeInfoKind.IF ) then
      return false
   end
   
   local baseTypeInfo = self:get_baseTypeInfo()
   if baseTypeInfo ~= _moduleObj.headTypeInfo then
      if baseTypeInfo:isInheritFrom( other ) then
         return true
      end
      
   end
   
   
   for __index, ifType in pairs( self:get_interfaceList() ) do
      if ifType:isInheritFrom( other ) then
         return true
      end
      
   end
   
   return false
end

local MatchType = {}
_moduleObj.MatchType = MatchType
MatchType._val2NameMap = {}
function MatchType:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MatchType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function MatchType._from( val )
   if MatchType._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
MatchType.__allList = {}
function MatchType.get__allList()
   return MatchType.__allList
end

MatchType.Match = 0
MatchType._val2NameMap[0] = 'Match'
MatchType.__allList[1] = MatchType.Match
MatchType.Warn = 1
MatchType._val2NameMap[1] = 'Warn'
MatchType.__allList[2] = MatchType.Warn
MatchType.Error = 2
MatchType._val2NameMap[2] = 'Error'
MatchType.__allList[3] = MatchType.Error

function TypeInfo.checkMatchType( dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex )

   
   local function checkDstTypeFrom( index, srcType, srcType2nd )
   
      local workExpType = srcType
      for dstIndex = index, #dstTypeList do
         local workDstType = dstTypeList[dstIndex]
         local matchResult = MatchType.Match
         if not workDstType:canEvalWith( workExpType, "=" ) then
            local message = string.format( "exp(%d) type mismatch %s <- %s", dstIndex, workDstType:getTxt( true ), workExpType:getTxt( true ))
            return MatchType.Error, message
         elseif workExpType == _moduleObj.builtinTypeAbbrNone then
            return MatchType.Warn, Code.format( Code.ID.nothing_define_abbr, string.format( "use '##', instate of %s.", workDstType:getTxt( true )) )
         end
         
         if matchResult ~= MatchType.Match then
         end
         
         workExpType = srcType2nd
      end
      
      return MatchType.Match, ""
   end
   
   local function checkSrcTypeFrom( index, dstType )
   
      for srcIndex = index, #expTypeList do
         local expType = expTypeList[srcIndex]
         local checkType = expType
         if expType:get_kind() == TypeInfoKind.DDD then
            if #expType:get_itemTypeInfoList() > 0 then
               checkType = expType:get_itemTypeInfoList()[1]
            else
             
               checkType = _moduleObj.builtinTypeStem_
            end
            
         end
         
         if not dstType:canEvalWith( checkType, "=" ) then
            return MatchType.Error, string.format( "exp(%d) type mismatch %s <- %s", srcIndex, dstType:getTxt( true ), expType:getTxt( true ))
         end
         
         if warnForFollowSrcIndex ~= nil then
            if warnForFollowSrcIndex <= srcIndex and dstType:get_nilable() then
               local mess = string.format( "use '**' at arg(%d). %s <- %s", srcIndex, dstType:getTxt( true ), expType:getTxt( true ))
               return MatchType.Warn, mess
            end
            
         end
         
         
      end
      
      return MatchType.Match, ""
   end
   
   if #expTypeList > 0 then
      for index, expType in pairs( expTypeList ) do
         if #dstTypeList == 0 then
            return MatchType.Error, string.format( "over exp. expect:0, actual:%d", #expTypeList)
         end
         
         local dstType = dstTypeList[index]
         if #dstTypeList == index then
            if dstType:get_srcTypeInfo():get_kind() ~= TypeInfoKind.DDD then
               if not dstType:canEvalWith( expType, "=" ) then
                  return MatchType.Error, string.format( "exp(%d) type mismatch %s <- %s", index, dstType:getTxt( true ), expType:getTxt( true ))
               end
               
               if not allowDstShort and #dstTypeList < #expTypeList then
                  return MatchType.Error, string.format( "over exp. expect: %d: actual: %d", #dstTypeList, #expTypeList)
               end
               
            else
             
               local dddItemType = _moduleObj.builtinTypeStem_
               if #dstType:get_itemTypeInfoList() > 0 then
                  dddItemType = dstType:get_itemTypeInfoList()[1]
               end
               
               local result, mess = checkSrcTypeFrom( index, dddItemType )
               if result ~= MatchType.Match then
                  return result, mess
               end
               
            end
            
            if warnForFollowSrcIndex ~= nil then
               if warnForFollowSrcIndex <= index and dstType:get_nilable() then
                  local mess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( true ), expType:getTxt( true ))
                  return MatchType.Warn, mess
               end
               
            end
            
            break
         elseif #expTypeList == index then
            local srcType = expType
            local srcType2nd = _moduleObj.builtinTypeAbbrNone
            if expType:get_kind() == TypeInfoKind.DDD then
               srcType = _moduleObj.builtinTypeStem_
               srcType2nd = _moduleObj.builtinTypeStem_
               if #expType:get_itemTypeInfoList() > 0 then
                  srcType = expType:get_itemTypeInfoList()[1]
                  srcType2nd = srcType
               end
               
            elseif expType == _moduleObj.builtinTypeAbbr then
               srcType2nd = _moduleObj.builtinTypeAbbr
            end
            
            local result, mess = checkDstTypeFrom( index, srcType, srcType2nd )
            if result ~= MatchType.Match then
               return result, mess
            end
            
            if warnForFollowSrcIndex ~= nil then
               if warnForFollowSrcIndex <= index and dstType:get_nilable() then
                  local warnMess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( true ), expType:getTxt( true ))
                  return MatchType.Warn, warnMess
               end
               
            end
            
            break
         else
          
            if not dstType:canEvalWith( expType, "=" ) then
               return MatchType.Error, string.format( "exp(%d) type mismatch %s <- %s", index, dstType:getTxt( true ), expType:getTxt( true ))
            end
            
            if warnForFollowSrcIndex ~= nil then
               if warnForFollowSrcIndex <= index and dstType:get_nilable() then
                  local mess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( true ), expType:getTxt( true ))
                  return MatchType.Warn, mess
               end
               
            end
            
         end
         
      end
      
   elseif not allowDstShort then
      for index, dstType in pairs( dstTypeList ) do
         if not dstType:canEvalWith( _moduleObj.builtinTypeNil, "=" ) then
            return MatchType.Error, string.format( "exp(%d) type mismatch %s <- nil", index, dstType:getTxt( true ))
         end
         
         return MatchType.Warn, Code.format( Code.ID.nothing_define_abbr, string.format( "use '##', instate of %s.", dstType:getTxt( true )) )
      end
      
   end
   
   return MatchType.Match, ""
end

function TypeInfo.canEvalWithBase( dest, destMut, other, opTxt )

   local otherMut = other:get_mutable()
   local otherSrc = other:get_srcTypeInfo()
   if otherSrc:get_kind() == TypeInfoKind.DDD then
      if #otherSrc:get_itemTypeInfoList() > 0 then
         otherSrc = otherSrc:get_itemTypeInfoList()[1]
      end
      
   end
   
   if opTxt == "=" and otherSrc ~= _moduleObj.builtinTypeNil and otherSrc ~= _moduleObj.builtinTypeString and otherSrc:get_kind() ~= TypeInfoKind.Prim and otherSrc:get_kind() ~= TypeInfoKind.Func and otherSrc:get_kind() ~= TypeInfoKind.Enum and otherSrc:get_kind() ~= TypeInfoKind.Abbr and destMut and not otherMut then
      return false
   end
   
   if dest == _moduleObj.builtinTypeStem_ then
      return true
   end
   
   if dest:get_srcTypeInfo():get_kind() == TypeInfoKind.DDD then
      if #dest:get_itemTypeInfoList() > 0 then
         return dest:get_itemTypeInfoList()[1]:canEvalWith( other, opTxt )
      end
      
      return true
   end
   
   if not dest:get_nilable() and otherSrc:get_nilable() then
      return false
   end
   
   if dest == _moduleObj.builtinTypeStem and not otherSrc:get_nilable() then
      return true
   end
   
   if dest == _moduleObj.builtinTypeForm and otherSrc:get_kind() == TypeInfoKind.Func then
      return true
   end
   
   if otherSrc == _moduleObj.builtinTypeNil or otherSrc:get_kind() == TypeInfoKind.Abbr then
      if dest:get_kind() ~= TypeInfoKind.Nilable then
         return false
      end
      
      return true
   end
   
   if dest:get_typeId() == otherSrc:get_typeId() then
      return true
   end
   
   if dest:get_kind() ~= otherSrc:get_kind() then
      if dest:get_kind() == TypeInfoKind.Nilable then
         if otherSrc:get_nilable() then
            return dest:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt )
         end
         
         return dest:get_nonnilableType():canEvalWith( otherSrc, opTxt )
      elseif (dest:get_kind() == TypeInfoKind.Class or dest:get_kind() == TypeInfoKind.IF ) and (otherSrc:get_kind() == TypeInfoKind.Class or otherSrc:get_kind() == TypeInfoKind.IF ) then
         return otherSrc:isInheritFrom( dest )
      elseif otherSrc:get_kind() == TypeInfoKind.Enum then
         local enumTypeInfo = otherSrc
         return dest:canEvalWith( enumTypeInfo:get_valTypeInfo(), opTxt )
      end
      
      return false
   end
   
   do
      local _switchExp = (dest:get_kind() )
      if _switchExp == TypeInfoKind.Prim then
         if dest == _moduleObj.builtinTypeInt and otherSrc == _moduleObj.builtinTypeChar or dest == _moduleObj.builtinTypeChar and otherSrc == _moduleObj.builtinTypeInt then
            return true
         end
         
         return false
      elseif _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone then
            return true
         end
         
         if not (_lune.unwrap( dest:get_itemTypeInfoList()[1]) ):canEvalWith( _lune.unwrap( otherSrc:get_itemTypeInfoList()[1]), "=" ) then
            return false
         end
         
         
         return true
      elseif _switchExp == TypeInfoKind.Map then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone and otherSrc:get_itemTypeInfoList()[2] == _moduleObj.builtinTypeNone then
            return true
         end
         
         if not (_lune.unwrap( dest:get_itemTypeInfoList()[1]) ):canEvalWith( _lune.unwrap( otherSrc:get_itemTypeInfoList()[1]), "=" ) then
            return false
         end
         
         
         if not (_lune.unwrap( dest:get_itemTypeInfoList()[2]) ):canEvalWith( _lune.unwrap( otherSrc:get_itemTypeInfoList()[2]), "=" ) then
            return false
         end
         
         
         return true
      elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
         return otherSrc:isInheritFrom( dest )
      elseif _switchExp == TypeInfoKind.Func then
         if dest == _moduleObj.builtinTypeForm then
            return true
         end
         
         if TypeInfo.checkMatchType( dest:get_argTypeInfoList(), otherSrc:get_argTypeInfoList(), false, nil ) == MatchType.Error or TypeInfo.checkMatchType( dest:get_retTypeInfoList(), otherSrc:get_retTypeInfoList(), false, nil ) == MatchType.Error or #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false
         end
         
         return true
      elseif _switchExp == TypeInfoKind.Method then
         if TypeInfo.checkMatchType( dest:get_argTypeInfoList(), otherSrc:get_argTypeInfoList(), false, nil ) == MatchType.Error or TypeInfo.checkMatchType( dest:get_retTypeInfoList(), otherSrc:get_retTypeInfoList(), false, nil ) == MatchType.Error or #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false
         end
         
         return true
      elseif _switchExp == TypeInfoKind.Nilable then
         return dest:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt )
      else 
         
            return false
      end
   end
   
end

function NormalTypeInfo:canEvalWith( other, opTxt )

   return TypeInfo.canEvalWithBase( self, self:get_mutable(), other, opTxt )
end

local Filter = {}
_moduleObj.Filter = Filter
function Filter.setmeta( obj )
  setmetatable( obj, { __index = Filter  } )
end
function Filter.new(  )
   local obj = {}
   Filter.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function Filter:__init(  ) 

end

local BreakKind = {}
_moduleObj.BreakKind = BreakKind
BreakKind._val2NameMap = {}
function BreakKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BreakKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function BreakKind._from( val )
   if BreakKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
BreakKind.__allList = {}
function BreakKind.get__allList()
   return BreakKind.__allList
end

BreakKind.None = 0
BreakKind._val2NameMap[0] = 'None'
BreakKind.__allList[1] = BreakKind.None
BreakKind.Break = 1
BreakKind._val2NameMap[1] = 'Break'
BreakKind.__allList[2] = BreakKind.Break
BreakKind.Return = 2
BreakKind._val2NameMap[2] = 'Return'
BreakKind.__allList[3] = BreakKind.Return
BreakKind.NeverRet = 3
BreakKind._val2NameMap[3] = 'NeverRet'
BreakKind.__allList[4] = BreakKind.NeverRet

local CheckBreakMode = {}
_moduleObj.CheckBreakMode = CheckBreakMode
CheckBreakMode._val2NameMap = {}
function CheckBreakMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CheckBreakMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function CheckBreakMode._from( val )
   if CheckBreakMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
CheckBreakMode.__allList = {}
function CheckBreakMode.get__allList()
   return CheckBreakMode.__allList
end

CheckBreakMode.Normal = 0
CheckBreakMode._val2NameMap[0] = 'Normal'
CheckBreakMode.__allList[1] = CheckBreakMode.Normal
CheckBreakMode.Return = 1
CheckBreakMode._val2NameMap[1] = 'Return'
CheckBreakMode.__allList[2] = CheckBreakMode.Return
CheckBreakMode.IgnoreFlow = 2
CheckBreakMode._val2NameMap[2] = 'IgnoreFlow'
CheckBreakMode.__allList[3] = CheckBreakMode.IgnoreFlow
CheckBreakMode.IgnoreFlowReturn = 3
CheckBreakMode._val2NameMap[3] = 'IgnoreFlowReturn'
CheckBreakMode.__allList[4] = CheckBreakMode.IgnoreFlowReturn

local Node = {}
_moduleObj.Node = Node
function Node:get_expType(  )

   if #self.expTypeList == 0 then
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
function Node:canBeStatement(  )

   return false
end
function Node:getBreakKind( checkMode )

   return BreakKind.None
end
function Node.setmeta( obj )
  setmetatable( obj, { __index = Node  } )
end
function Node.new( kind, pos, expTypeList )
   local obj = {}
   Node.setmeta( obj )
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
_moduleObj.NamespaceInfo = NamespaceInfo
function NamespaceInfo.setmeta( obj )
  setmetatable( obj, { __index = NamespaceInfo  } )
end
function NamespaceInfo.new( name, scope, typeInfo )
   local obj = {}
   NamespaceInfo.setmeta( obj )
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





local DeclMacroInfo = {}
_moduleObj.DeclMacroInfo = DeclMacroInfo
function DeclMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = DeclMacroInfo  } )
end
function DeclMacroInfo.new( pubFlag, name, argList, stmtBlock, tokenList )
   local obj = {}
   DeclMacroInfo.setmeta( obj )
   if obj.__init then
      obj:__init( pubFlag, name, argList, stmtBlock, tokenList )
   end        
   return obj 
end         
function DeclMacroInfo:__init( pubFlag, name, argList, stmtBlock, tokenList ) 

   self.pubFlag = pubFlag
   self.name = name
   self.argList = argList
   self.stmtBlock = stmtBlock
   self.tokenList = tokenList
end
function DeclMacroInfo:get_pubFlag()       
   return self.pubFlag         
end
function DeclMacroInfo:get_name()       
   return self.name         
end
function DeclMacroInfo:get_argList()       
   return self.argList         
end
function DeclMacroInfo:get_stmtBlock()       
   return self.stmtBlock         
end
function DeclMacroInfo:get_tokenList()       
   return self.tokenList         
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

   return _lune.unwrap( nodeKind2NameMap[kind])
end
_moduleObj.getNodeKindName = getNodeKindName
local NodeManager = {}
_moduleObj.NodeManager = NodeManager
function NodeManager.new(  )
   local obj = {}
   NodeManager.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function NodeManager:__init() 
   self.nodeKind2NodeList = {}
end
function NodeManager:getList( kind )

   return self.nodeKind2NodeList[kind]
end
function NodeManager:addNode( node )

   local list = self.nodeKind2NodeList[node:get_kind()]
   if  nil == list then
      local _list = list
   
      list = {}
      self.nodeKind2NodeList[node:get_kind()] = list
   end
   
   table.insert( list, node )
end
function NodeManager.setmeta( obj )
  setmetatable( obj, { __index = NodeManager  } )
end

local NodeKind = {}
_moduleObj.NodeKind = NodeKind
function NodeKind.setmeta( obj )
  setmetatable( obj, { __index = NodeKind  } )
end
function NodeKind.new(  )
   local obj = {}
   NodeKind.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function NodeKind:__init(  ) 

end


function NodeKind.get_None(  )

   return _lune.unwrap( _moduleObj.nodeKind['None'])
end


regKind( [[None]] )
function Filter:processNone( node, ... )

end


function NodeManager:getNoneNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['None']) )
end


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
function NoneNode:canBeStatement(  )

   return true
end
function NoneNode.new( pos, typeList )
   local obj = {}
   NoneNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList ); end
   return obj
end
function NoneNode:__init(pos, typeList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['None']), pos, typeList)
   
   
   
end
function NoneNode.create( nodeMan, pos, typeList )

   local node = NoneNode.new(pos, typeList)
   nodeMan:addNode( node )
   return node
end
function NoneNode.setmeta( obj )
  setmetatable( obj, { __index = NoneNode  } )
end


function NodeKind.get_Subfile(  )

   return _lune.unwrap( _moduleObj.nodeKind['Subfile'])
end


regKind( [[Subfile]] )
function Filter:processSubfile( node, ... )

end


function NodeManager:getSubfileNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Subfile']) )
end


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
function SubfileNode:canBeStatement(  )

   return true
end
function SubfileNode.new( pos, typeList, usePath )
   local obj = {}
   SubfileNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, usePath ); end
   return obj
end
function SubfileNode:__init(pos, typeList, usePath) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Subfile']), pos, typeList)
   
   
   self.usePath = usePath
   
end
function SubfileNode.create( nodeMan, pos, typeList, usePath )

   local node = SubfileNode.new(pos, typeList, usePath)
   nodeMan:addNode( node )
   return node
end
function SubfileNode.setmeta( obj )
  setmetatable( obj, { __index = SubfileNode  } )
end
function SubfileNode:get_usePath()       
   return self.usePath         
end


function NodeKind.get_Import(  )

   return _lune.unwrap( _moduleObj.nodeKind['Import'])
end


regKind( [[Import]] )
function Filter:processImport( node, ... )

end


function NodeManager:getImportNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Import']) )
end


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
function ImportNode:canBeStatement(  )

   return true
end
function ImportNode.new( pos, typeList, modulePath, assignName, moduleTypeInfo )
   local obj = {}
   ImportNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, modulePath, assignName, moduleTypeInfo ); end
   return obj
end
function ImportNode:__init(pos, typeList, modulePath, assignName, moduleTypeInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Import']), pos, typeList)
   
   
   self.modulePath = modulePath
   self.assignName = assignName
   self.moduleTypeInfo = moduleTypeInfo
   
end
function ImportNode.create( nodeMan, pos, typeList, modulePath, assignName, moduleTypeInfo )

   local node = ImportNode.new(pos, typeList, modulePath, assignName, moduleTypeInfo)
   nodeMan:addNode( node )
   return node
end
function ImportNode.setmeta( obj )
  setmetatable( obj, { __index = ImportNode  } )
end
function ImportNode:get_modulePath()       
   return self.modulePath         
end
function ImportNode:get_assignName()       
   return self.assignName         
end
function ImportNode:get_moduleTypeInfo()       
   return self.moduleTypeInfo         
end



local LuneHelperInfo = {}
_moduleObj.LuneHelperInfo = LuneHelperInfo
function LuneHelperInfo.setmeta( obj )
  setmetatable( obj, { __index = LuneHelperInfo  } )
end
function LuneHelperInfo.new( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge )
   local obj = {}
   LuneHelperInfo.setmeta( obj )
   if obj.__init then
      obj:__init( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge )
   end        
   return obj 
end         
function LuneHelperInfo:__init( useNilAccess, useUnwrapExp, hasMappingClassDef, useLoad, useUnpack, useAlge ) 

   self.useNilAccess = useNilAccess
   self.useUnwrapExp = useUnwrapExp
   self.hasMappingClassDef = hasMappingClassDef
   self.useLoad = useLoad
   self.useUnpack = useUnpack
   self.useAlge = useAlge
end

local ModuleInfo = {}
_moduleObj.ModuleInfo = ModuleInfo
function ModuleInfo.new( fullName, assignName, idMap, moduleId )
   local obj = {}
   ModuleInfo.setmeta( obj )
   if obj.__init then obj:__init( fullName, assignName, idMap, moduleId ); end
   return obj
end
function ModuleInfo:__init(fullName, assignName, idMap, moduleId) 
   self.moduleId = moduleId
   self.fullName = fullName
   self.assignName = assignName
   self.localTypeInfo2importIdMap = idMap
   self.importId2localTypeInfoMap = {}
   for typeInfo, importId in pairs( idMap ) do
      self.importId2localTypeInfoMap[importId] = typeInfo
   end
   
end
function ModuleInfo:get_modulePath(  )

   return self.fullName
end
function ModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfo  } )
end
function ModuleInfo:get_fullName()       
   return self.fullName         
end
function ModuleInfo:get_localTypeInfo2importIdMap()       
   return self.localTypeInfo2importIdMap         
end
function ModuleInfo:get_importId2localTypeInfoMap()       
   return self.importId2localTypeInfoMap         
end
function ModuleInfo:get_assignName()       
   return self.assignName         
end
function ModuleInfo:get_moduleId()       
   return self.moduleId         
end

function TypeInfo:getFullName( importInfo, localFlag )

   return self:getParentFullName( importInfo, localFlag ) .. self:get_rawTxt()
end

local ProcessInfo = {}
_moduleObj.ProcessInfo = ProcessInfo
function ProcessInfo.setmeta( obj )
  setmetatable( obj, { __index = ProcessInfo  } )
end
function ProcessInfo.new( idProvier, typeInfo2ModifierMap, typeInfo2DDDMap )
   local obj = {}
   ProcessInfo.setmeta( obj )
   if obj.__init then
      obj:__init( idProvier, typeInfo2ModifierMap, typeInfo2DDDMap )
   end        
   return obj 
end         
function ProcessInfo:__init( idProvier, typeInfo2ModifierMap, typeInfo2DDDMap ) 

   self.idProvier = idProvier
   self.typeInfo2ModifierMap = typeInfo2ModifierMap
   self.typeInfo2DDDMap = typeInfo2DDDMap
end
function ProcessInfo:get_idProvier()       
   return self.idProvier         
end
function ProcessInfo:get_typeInfo2ModifierMap()       
   return self.typeInfo2ModifierMap         
end
function ProcessInfo:get_typeInfo2DDDMap()       
   return self.typeInfo2DDDMap         
end

local MacroValInfo = {}
_moduleObj.MacroValInfo = MacroValInfo
function MacroValInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroValInfo  } )
end
function MacroValInfo.new( val, typeInfo )
   local obj = {}
   MacroValInfo.setmeta( obj )
   if obj.__init then
      obj:__init( val, typeInfo )
   end        
   return obj 
end         
function MacroValInfo:__init( val, typeInfo ) 

   self.val = val
   self.typeInfo = typeInfo
end

local MacroArgInfo = {}
_moduleObj.MacroArgInfo = MacroArgInfo
function MacroArgInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroArgInfo  } )
end
function MacroArgInfo.new( name, typeInfo )
   local obj = {}
   MacroArgInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeInfo )
   end        
   return obj 
end         
function MacroArgInfo:__init( name, typeInfo ) 

   self.name = name
   self.typeInfo = typeInfo
end
function MacroArgInfo:get_name()       
   return self.name         
end
function MacroArgInfo:get_typeInfo()       
   return self.typeInfo         
end

local MacroInfo = {}
_moduleObj.MacroInfo = MacroInfo
function MacroInfo.setmeta( obj )
  setmetatable( obj, { __index = MacroInfo  } )
end
function MacroInfo.new( func, symbol2MacroValInfoMap )
   local obj = {}
   MacroInfo.setmeta( obj )
   if obj.__init then
      obj:__init( func, symbol2MacroValInfoMap )
   end        
   return obj 
end         
function MacroInfo:__init( func, symbol2MacroValInfoMap ) 

   self.func = func
   self.symbol2MacroValInfoMap = symbol2MacroValInfoMap
end

function NodeKind.get_Root(  )

   return _lune.unwrap( _moduleObj.nodeKind['Root'])
end


regKind( [[Root]] )
function Filter:processRoot( node, ... )

end


function NodeManager:getRootNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Root']) )
end


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
function RootNode:canBeStatement(  )

   return false
end
function RootNode.new( pos, typeList, children, useModuleMacroMap, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )
   local obj = {}
   RootNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, children, useModuleMacroMap, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap ); end
   return obj
end
function RootNode:__init(pos, typeList, children, useModuleMacroMap, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Root']), pos, typeList)
   
   
   self.children = children
   self.useModuleMacroMap = useModuleMacroMap
   self.moduleId = moduleId
   self.processInfo = processInfo
   self.moduleTypeInfo = moduleTypeInfo
   self.provideNode = provideNode
   self.luneHelperInfo = luneHelperInfo
   self.nodeManager = nodeManager
   self.importModule2moduleInfo = importModule2moduleInfo
   self.typeId2MacroInfo = typeId2MacroInfo
   self.typeId2ClassMap = typeId2ClassMap
   
end
function RootNode.create( nodeMan, pos, typeList, children, useModuleMacroMap, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap )

   local node = RootNode.new(pos, typeList, children, useModuleMacroMap, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap)
   nodeMan:addNode( node )
   return node
end
function RootNode.setmeta( obj )
  setmetatable( obj, { __index = RootNode  } )
end
function RootNode:get_children()       
   return self.children         
end
function RootNode:get_useModuleMacroMap()       
   return self.useModuleMacroMap         
end
function RootNode:get_moduleId()       
   return self.moduleId         
end
function RootNode:get_processInfo()       
   return self.processInfo         
end
function RootNode:get_moduleTypeInfo()       
   return self.moduleTypeInfo         
end
function RootNode:get_provideNode()       
   return self.provideNode         
end
function RootNode:get_luneHelperInfo()       
   return self.luneHelperInfo         
end
function RootNode:get_nodeManager()       
   return self.nodeManager         
end
function RootNode:get_importModule2moduleInfo()       
   return self.importModule2moduleInfo         
end
function RootNode:get_typeId2MacroInfo()       
   return self.typeId2MacroInfo         
end
function RootNode:get_typeId2ClassMap()       
   return self.typeId2ClassMap         
end


function RootNode:set_provide( node )

   self.provideNode = node
end

function NodeKind.get_RefType(  )

   return _lune.unwrap( _moduleObj.nodeKind['RefType'])
end


regKind( [[RefType]] )
function Filter:processRefType( node, ... )

end


function NodeManager:getRefTypeNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['RefType']) )
end


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
function RefTypeNode:canBeStatement(  )

   return false
end
function RefTypeNode.new( pos, typeList, name, refFlag, mutFlag, array )
   local obj = {}
   RefTypeNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, name, refFlag, mutFlag, array ); end
   return obj
end
function RefTypeNode:__init(pos, typeList, name, refFlag, mutFlag, array) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['RefType']), pos, typeList)
   
   
   self.name = name
   self.refFlag = refFlag
   self.mutFlag = mutFlag
   self.array = array
   
end
function RefTypeNode.create( nodeMan, pos, typeList, name, refFlag, mutFlag, array )

   local node = RefTypeNode.new(pos, typeList, name, refFlag, mutFlag, array)
   nodeMan:addNode( node )
   return node
end
function RefTypeNode.setmeta( obj )
  setmetatable( obj, { __index = RefTypeNode  } )
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


local BlockKind = {}
_moduleObj.BlockKind = BlockKind
BlockKind._val2NameMap = {}
function BlockKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BlockKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function BlockKind._from( val )
   if BlockKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
BlockKind.__allList = {}
function BlockKind.get__allList()
   return BlockKind.__allList
end

BlockKind.If = 0
BlockKind._val2NameMap[0] = 'If'
BlockKind.__allList[1] = BlockKind.If
BlockKind.Elseif = 1
BlockKind._val2NameMap[1] = 'Elseif'
BlockKind.__allList[2] = BlockKind.Elseif
BlockKind.Else = 2
BlockKind._val2NameMap[2] = 'Else'
BlockKind.__allList[3] = BlockKind.Else
BlockKind.While = 3
BlockKind._val2NameMap[3] = 'While'
BlockKind.__allList[4] = BlockKind.While
BlockKind.Switch = 4
BlockKind._val2NameMap[4] = 'Switch'
BlockKind.__allList[5] = BlockKind.Switch
BlockKind.Match = 5
BlockKind._val2NameMap[5] = 'Match'
BlockKind.__allList[6] = BlockKind.Match
BlockKind.Repeat = 6
BlockKind._val2NameMap[6] = 'Repeat'
BlockKind.__allList[7] = BlockKind.Repeat
BlockKind.For = 7
BlockKind._val2NameMap[7] = 'For'
BlockKind.__allList[8] = BlockKind.For
BlockKind.Apply = 8
BlockKind._val2NameMap[8] = 'Apply'
BlockKind.__allList[9] = BlockKind.Apply
BlockKind.Foreach = 9
BlockKind._val2NameMap[9] = 'Foreach'
BlockKind.__allList[10] = BlockKind.Foreach
BlockKind.Macro = 14
BlockKind._val2NameMap[14] = 'Macro'
BlockKind.__allList[11] = BlockKind.Macro
BlockKind.Func = 11
BlockKind._val2NameMap[11] = 'Func'
BlockKind.__allList[12] = BlockKind.Func
BlockKind.Default = 12
BlockKind._val2NameMap[12] = 'Default'
BlockKind.__allList[13] = BlockKind.Default
BlockKind.Block = 13
BlockKind._val2NameMap[13] = 'Block'
BlockKind.__allList[14] = BlockKind.Block
BlockKind.Macro = 14
BlockKind._val2NameMap[14] = 'Macro'
BlockKind.__allList[15] = BlockKind.Macro
BlockKind.LetUnwrap = 15
BlockKind._val2NameMap[15] = 'LetUnwrap'
BlockKind.__allList[16] = BlockKind.LetUnwrap
BlockKind.IfUnwrap = 16
BlockKind._val2NameMap[16] = 'IfUnwrap'
BlockKind.__allList[17] = BlockKind.IfUnwrap
BlockKind.When = 17
BlockKind._val2NameMap[17] = 'When'
BlockKind.__allList[18] = BlockKind.When

function NodeKind.get_Block(  )

   return _lune.unwrap( _moduleObj.nodeKind['Block'])
end


regKind( [[Block]] )
function Filter:processBlock( node, ... )

end


function NodeManager:getBlockNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Block']) )
end


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
function BlockNode:canBeStatement(  )

   return true
end
function BlockNode.new( pos, typeList, blockKind, stmtList )
   local obj = {}
   BlockNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, blockKind, stmtList ); end
   return obj
end
function BlockNode:__init(pos, typeList, blockKind, stmtList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Block']), pos, typeList)
   
   
   self.blockKind = blockKind
   self.stmtList = stmtList
   
end
function BlockNode.create( nodeMan, pos, typeList, blockKind, stmtList )

   local node = BlockNode.new(pos, typeList, blockKind, stmtList)
   nodeMan:addNode( node )
   return node
end
function BlockNode.setmeta( obj )
  setmetatable( obj, { __index = BlockNode  } )
end
function BlockNode:get_blockKind()       
   return self.blockKind         
end
function BlockNode:get_stmtList()       
   return self.stmtList         
end



function BlockNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in pairs( self.stmtList ) do
         local work = stmt:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      return kind
   else
    
      if #self.stmtList > 0 then
         local node = self.stmtList[#self.stmtList]
         return node:getBreakKind( checkMode )
      end
      
   end
   
   return BreakKind.None
end

local IfKind = {}
_moduleObj.IfKind = IfKind
IfKind._val2NameMap = {}
function IfKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "IfKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function IfKind._from( val )
   if IfKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
IfKind.__allList = {}
function IfKind.get__allList()
   return IfKind.__allList
end

IfKind.If = 0
IfKind._val2NameMap[0] = 'If'
IfKind.__allList[1] = IfKind.If
IfKind.ElseIf = 1
IfKind._val2NameMap[1] = 'ElseIf'
IfKind.__allList[2] = IfKind.ElseIf
IfKind.Else = 2
IfKind._val2NameMap[2] = 'Else'
IfKind.__allList[3] = IfKind.Else

local IfStmtInfo = {}
_moduleObj.IfStmtInfo = IfStmtInfo
function IfStmtInfo.setmeta( obj )
  setmetatable( obj, { __index = IfStmtInfo  } )
end
function IfStmtInfo.new( kind, exp, block )
   local obj = {}
   IfStmtInfo.setmeta( obj )
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

function NodeKind.get_If(  )

   return _lune.unwrap( _moduleObj.nodeKind['If'])
end


regKind( [[If]] )
function Filter:processIf( node, ... )

end


function NodeManager:getIfNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['If']) )
end


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
function IfNode:canBeStatement(  )

   return true
end
function IfNode.new( pos, typeList, stmtList )
   local obj = {}
   IfNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, stmtList ); end
   return obj
end
function IfNode:__init(pos, typeList, stmtList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['If']), pos, typeList)
   
   
   self.stmtList = stmtList
   
end
function IfNode.create( nodeMan, pos, typeList, stmtList )

   local node = IfNode.new(pos, typeList, stmtList)
   nodeMan:addNode( node )
   return node
end
function IfNode.setmeta( obj )
  setmetatable( obj, { __index = IfNode  } )
end
function IfNode:get_stmtList()       
   return self.stmtList         
end


function IfNode:getBreakKind( checkMode )

   local hasElseFlag = false
   local kind = BreakKind.None
   for __index, stmtInfo in pairs( self.stmtList ) do
      local work = stmtInfo:get_block():getBreakKind( checkMode )
      if checkMode == CheckBreakMode.IgnoreFlowReturn then
         if work == BreakKind.Return then
            return BreakKind.Return
         end
         
         if work == BreakKind.NeverRet then
            return BreakKind.NeverRet
         end
         
      else
       
         do
            local _switchExp = work
            if _switchExp == BreakKind.None then
               if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                  return BreakKind.None
               end
               
            else 
               
                  if kind == BreakKind.None or kind > work then
                     kind = work
                  end
                  
            end
         end
         
      end
      
      
      if stmtInfo:get_kind() == IfKind.Else then
         hasElseFlag = true
      end
      
   end
   
   if hasElseFlag or (checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return ) then
      return kind
   end
   
   return BreakKind.None
end

function NodeKind.get_ExpList(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpList'])
end


regKind( [[ExpList]] )
function Filter:processExpList( node, ... )

end


function NodeManager:getExpListNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpList']) )
end


local ExpListNode = {}
setmetatable( ExpListNode, { __index = Node } )
_moduleObj.ExpListNode = ExpListNode
function ExpListNode:processFilter( filter, ... )

   local argList = {...}
   filter:processExpList( self, table.unpack( argList ) )
end
function ExpListNode:canBeStatement(  )

   return false
end
function ExpListNode.new( pos, typeList, expList, followOn )
   local obj = {}
   ExpListNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, expList, followOn ); end
   return obj
end
function ExpListNode:__init(pos, typeList, expList, followOn) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpList']), pos, typeList)
   
   
   self.expList = expList
   self.followOn = followOn
   
end
function ExpListNode.create( nodeMan, pos, typeList, expList, followOn )

   local node = ExpListNode.new(pos, typeList, expList, followOn)
   nodeMan:addNode( node )
   return node
end
function ExpListNode.setmeta( obj )
  setmetatable( obj, { __index = ExpListNode  } )
end
function ExpListNode:get_expList()       
   return self.expList         
end
function ExpListNode:get_followOn()       
   return self.followOn         
end


function ExpListNode:canBeLeft(  )

   for __index, expNode in pairs( self:get_expList() ) do
      if not expNode:canBeLeft(  ) then
         return false
      end
      
   end
   
   return true
end

function ExpListNode:canBeRight(  )

   for __index, expNode in pairs( self:get_expList() ) do
      if not expNode:canBeRight(  ) then
         return false
      end
      
   end
   
   return true
end

local CaseInfo = {}
_moduleObj.CaseInfo = CaseInfo
function CaseInfo.setmeta( obj )
  setmetatable( obj, { __index = CaseInfo  } )
end
function CaseInfo.new( expList, block )
   local obj = {}
   CaseInfo.setmeta( obj )
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

function NodeKind.get_Switch(  )

   return _lune.unwrap( _moduleObj.nodeKind['Switch'])
end


regKind( [[Switch]] )
function Filter:processSwitch( node, ... )

end


function NodeManager:getSwitchNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Switch']) )
end


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
function SwitchNode:canBeStatement(  )

   return true
end
function SwitchNode.new( pos, typeList, exp, caseList, default )
   local obj = {}
   SwitchNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp, caseList, default ); end
   return obj
end
function SwitchNode:__init(pos, typeList, exp, caseList, default) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Switch']), pos, typeList)
   
   
   self.exp = exp
   self.caseList = caseList
   self.default = default
   
end
function SwitchNode.create( nodeMan, pos, typeList, exp, caseList, default )

   local node = SwitchNode.new(pos, typeList, exp, caseList, default)
   nodeMan:addNode( node )
   return node
end
function SwitchNode.setmeta( obj )
  setmetatable( obj, { __index = SwitchNode  } )
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


function SwitchNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   for __index, caseInfo in pairs( self.caseList ) do
      local work = caseInfo:get_block():getBreakKind( checkMode )
      if checkMode == CheckBreakMode.IgnoreFlowReturn then
         if work == BreakKind.Return then
            return BreakKind.Return
         end
         
         if work == BreakKind.NeverRet then
            return BreakKind.NeverRet
         end
         
      else
       
         do
            local _switchExp = work
            if _switchExp == BreakKind.None then
               if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                  return BreakKind.None
               end
               
            else 
               
                  if kind == BreakKind.None or kind > work then
                     kind = work
                  end
                  
            end
         end
         
      end
      
      
   end
   
   do
      local block = self.default
      if block ~= nil then
         local work = block:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     return BreakKind.None
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   return BreakKind.None
end


function NodeKind.get_While(  )

   return _lune.unwrap( _moduleObj.nodeKind['While'])
end


regKind( [[While]] )
function Filter:processWhile( node, ... )

end


function NodeManager:getWhileNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['While']) )
end


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
function WhileNode:canBeStatement(  )

   return true
end
function WhileNode.new( pos, typeList, exp, block )
   local obj = {}
   WhileNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp, block ); end
   return obj
end
function WhileNode:__init(pos, typeList, exp, block) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['While']), pos, typeList)
   
   
   self.exp = exp
   self.block = block
   
end
function WhileNode.create( nodeMan, pos, typeList, exp, block )

   local node = WhileNode.new(pos, typeList, exp, block)
   nodeMan:addNode( node )
   return node
end
function WhileNode.setmeta( obj )
  setmetatable( obj, { __index = WhileNode  } )
end
function WhileNode:get_exp()       
   return self.exp         
end
function WhileNode:get_block()       
   return self.block         
end


function NodeKind.get_Repeat(  )

   return _lune.unwrap( _moduleObj.nodeKind['Repeat'])
end


regKind( [[Repeat]] )
function Filter:processRepeat( node, ... )

end


function NodeManager:getRepeatNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Repeat']) )
end


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
function RepeatNode:canBeStatement(  )

   return true
end
function RepeatNode.new( pos, typeList, block, exp )
   local obj = {}
   RepeatNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, block, exp ); end
   return obj
end
function RepeatNode:__init(pos, typeList, block, exp) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Repeat']), pos, typeList)
   
   
   self.block = block
   self.exp = exp
   
end
function RepeatNode.create( nodeMan, pos, typeList, block, exp )

   local node = RepeatNode.new(pos, typeList, block, exp)
   nodeMan:addNode( node )
   return node
end
function RepeatNode.setmeta( obj )
  setmetatable( obj, { __index = RepeatNode  } )
end
function RepeatNode:get_block()       
   return self.block         
end
function RepeatNode:get_exp()       
   return self.exp         
end


function RepeatNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_For(  )

   return _lune.unwrap( _moduleObj.nodeKind['For'])
end


regKind( [[For]] )
function Filter:processFor( node, ... )

end


function NodeManager:getForNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['For']) )
end


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
function ForNode:canBeStatement(  )

   return true
end
function ForNode.new( pos, typeList, block, val, init, to, delta )
   local obj = {}
   ForNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, block, val, init, to, delta ); end
   return obj
end
function ForNode:__init(pos, typeList, block, val, init, to, delta) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['For']), pos, typeList)
   
   
   self.block = block
   self.val = val
   self.init = init
   self.to = to
   self.delta = delta
   
end
function ForNode.create( nodeMan, pos, typeList, block, val, init, to, delta )

   local node = ForNode.new(pos, typeList, block, val, init, to, delta)
   nodeMan:addNode( node )
   return node
end
function ForNode.setmeta( obj )
  setmetatable( obj, { __index = ForNode  } )
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


function ForNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Apply(  )

   return _lune.unwrap( _moduleObj.nodeKind['Apply'])
end


regKind( [[Apply]] )
function Filter:processApply( node, ... )

end


function NodeManager:getApplyNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Apply']) )
end


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
function ApplyNode:canBeStatement(  )

   return true
end
function ApplyNode.new( pos, typeList, varList, exp, block )
   local obj = {}
   ApplyNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, varList, exp, block ); end
   return obj
end
function ApplyNode:__init(pos, typeList, varList, exp, block) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Apply']), pos, typeList)
   
   
   self.varList = varList
   self.exp = exp
   self.block = block
   
end
function ApplyNode.create( nodeMan, pos, typeList, varList, exp, block )

   local node = ApplyNode.new(pos, typeList, varList, exp, block)
   nodeMan:addNode( node )
   return node
end
function ApplyNode.setmeta( obj )
  setmetatable( obj, { __index = ApplyNode  } )
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


function ApplyNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Foreach(  )

   return _lune.unwrap( _moduleObj.nodeKind['Foreach'])
end


regKind( [[Foreach]] )
function Filter:processForeach( node, ... )

end


function NodeManager:getForeachNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Foreach']) )
end


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
function ForeachNode:canBeStatement(  )

   return true
end
function ForeachNode.new( pos, typeList, val, key, exp, block )
   local obj = {}
   ForeachNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, val, key, exp, block ); end
   return obj
end
function ForeachNode:__init(pos, typeList, val, key, exp, block) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Foreach']), pos, typeList)
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   
end
function ForeachNode.create( nodeMan, pos, typeList, val, key, exp, block )

   local node = ForeachNode.new(pos, typeList, val, key, exp, block)
   nodeMan:addNode( node )
   return node
end
function ForeachNode.setmeta( obj )
  setmetatable( obj, { __index = ForeachNode  } )
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


function ForeachNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Forsort(  )

   return _lune.unwrap( _moduleObj.nodeKind['Forsort'])
end


regKind( [[Forsort]] )
function Filter:processForsort( node, ... )

end


function NodeManager:getForsortNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Forsort']) )
end


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
function ForsortNode:canBeStatement(  )

   return true
end
function ForsortNode.new( pos, typeList, val, key, exp, block, sort )
   local obj = {}
   ForsortNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, val, key, exp, block, sort ); end
   return obj
end
function ForsortNode:__init(pos, typeList, val, key, exp, block, sort) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Forsort']), pos, typeList)
   
   
   self.val = val
   self.key = key
   self.exp = exp
   self.block = block
   self.sort = sort
   
end
function ForsortNode.create( nodeMan, pos, typeList, val, key, exp, block, sort )

   local node = ForsortNode.new(pos, typeList, val, key, exp, block, sort)
   nodeMan:addNode( node )
   return node
end
function ForsortNode.setmeta( obj )
  setmetatable( obj, { __index = ForsortNode  } )
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


function ForsortNode:getBreakKind( checkMode )

   local kind = BreakKind.None
   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      return self.block:getBreakKind( checkMode )
   end
   
   return BreakKind.None
end


function NodeKind.get_Return(  )

   return _lune.unwrap( _moduleObj.nodeKind['Return'])
end


regKind( [[Return]] )
function Filter:processReturn( node, ... )

end


function NodeManager:getReturnNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Return']) )
end


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
function ReturnNode:canBeStatement(  )

   return true
end
function ReturnNode.new( pos, typeList, expList )
   local obj = {}
   ReturnNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, expList ); end
   return obj
end
function ReturnNode:__init(pos, typeList, expList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Return']), pos, typeList)
   
   
   self.expList = expList
   
end
function ReturnNode.create( nodeMan, pos, typeList, expList )

   local node = ReturnNode.new(pos, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function ReturnNode.setmeta( obj )
  setmetatable( obj, { __index = ReturnNode  } )
end
function ReturnNode:get_expList()       
   return self.expList         
end


function ReturnNode:getBreakKind( checkMode )

   return BreakKind.Return
end

function NodeKind.get_Break(  )

   return _lune.unwrap( _moduleObj.nodeKind['Break'])
end


regKind( [[Break]] )
function Filter:processBreak( node, ... )

end


function NodeManager:getBreakNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Break']) )
end


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
function BreakNode:canBeStatement(  )

   return true
end
function BreakNode.new( pos, typeList )
   local obj = {}
   BreakNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList ); end
   return obj
end
function BreakNode:__init(pos, typeList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Break']), pos, typeList)
   
   
   
end
function BreakNode.create( nodeMan, pos, typeList )

   local node = BreakNode.new(pos, typeList)
   nodeMan:addNode( node )
   return node
end
function BreakNode.setmeta( obj )
  setmetatable( obj, { __index = BreakNode  } )
end


function BreakNode:getBreakKind( checkMode )

   return BreakKind.Break
end

function NodeKind.get_Provide(  )

   return _lune.unwrap( _moduleObj.nodeKind['Provide'])
end


regKind( [[Provide]] )
function Filter:processProvide( node, ... )

end


function NodeManager:getProvideNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Provide']) )
end


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
function ProvideNode:canBeStatement(  )

   return true
end
function ProvideNode.new( pos, typeList, symbol )
   local obj = {}
   ProvideNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, symbol ); end
   return obj
end
function ProvideNode:__init(pos, typeList, symbol) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Provide']), pos, typeList)
   
   
   self.symbol = symbol
   
end
function ProvideNode.create( nodeMan, pos, typeList, symbol )

   local node = ProvideNode.new(pos, typeList, symbol)
   nodeMan:addNode( node )
   return node
end
function ProvideNode.setmeta( obj )
  setmetatable( obj, { __index = ProvideNode  } )
end
function ProvideNode:get_symbol()       
   return self.symbol         
end


function NodeKind.get_ExpNew(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpNew'])
end


regKind( [[ExpNew]] )
function Filter:processExpNew( node, ... )

end


function NodeManager:getExpNewNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpNew']) )
end


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
function ExpNewNode:canBeStatement(  )

   return true
end
function ExpNewNode.new( pos, typeList, symbol, argList )
   local obj = {}
   ExpNewNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, symbol, argList ); end
   return obj
end
function ExpNewNode:__init(pos, typeList, symbol, argList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpNew']), pos, typeList)
   
   
   self.symbol = symbol
   self.argList = argList
   
end
function ExpNewNode.create( nodeMan, pos, typeList, symbol, argList )

   local node = ExpNewNode.new(pos, typeList, symbol, argList)
   nodeMan:addNode( node )
   return node
end
function ExpNewNode.setmeta( obj )
  setmetatable( obj, { __index = ExpNewNode  } )
end
function ExpNewNode:get_symbol()       
   return self.symbol         
end
function ExpNewNode:get_argList()       
   return self.argList         
end


function NodeKind.get_ExpUnwrap(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpUnwrap'])
end


regKind( [[ExpUnwrap]] )
function Filter:processExpUnwrap( node, ... )

end


function NodeManager:getExpUnwrapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpUnwrap']) )
end


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
function ExpUnwrapNode:canBeStatement(  )

   return false
end
function ExpUnwrapNode.new( pos, typeList, exp, default )
   local obj = {}
   ExpUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp, default ); end
   return obj
end
function ExpUnwrapNode:__init(pos, typeList, exp, default) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpUnwrap']), pos, typeList)
   
   
   self.exp = exp
   self.default = default
   
end
function ExpUnwrapNode.create( nodeMan, pos, typeList, exp, default )

   local node = ExpUnwrapNode.new(pos, typeList, exp, default)
   nodeMan:addNode( node )
   return node
end
function ExpUnwrapNode.setmeta( obj )
  setmetatable( obj, { __index = ExpUnwrapNode  } )
end
function ExpUnwrapNode:get_exp()       
   return self.exp         
end
function ExpUnwrapNode:get_default()       
   return self.default         
end


function NodeKind.get_ExpRef(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpRef'])
end


regKind( [[ExpRef]] )
function Filter:processExpRef( node, ... )

end


function NodeManager:getExpRefNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpRef']) )
end


local ExpRefNode = {}
setmetatable( ExpRefNode, { __index = Node } )
_moduleObj.ExpRefNode = ExpRefNode
function ExpRefNode:processFilter( filter, ... )

   local argList = {...}
   filter:processExpRef( self, table.unpack( argList ) )
end
function ExpRefNode:canBeStatement(  )

   return false
end
function ExpRefNode.new( pos, typeList, token, symbolInfo )
   local obj = {}
   ExpRefNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token, symbolInfo ); end
   return obj
end
function ExpRefNode:__init(pos, typeList, token, symbolInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpRef']), pos, typeList)
   
   
   self.token = token
   self.symbolInfo = symbolInfo
   
end
function ExpRefNode.create( nodeMan, pos, typeList, token, symbolInfo )

   local node = ExpRefNode.new(pos, typeList, token, symbolInfo)
   nodeMan:addNode( node )
   return node
end
function ExpRefNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefNode  } )
end
function ExpRefNode:get_token()       
   return self.token         
end
function ExpRefNode:get_symbolInfo()       
   return self.symbolInfo         
end


function ExpRefNode:canBeLeft(  )

   return self:get_symbolInfo():get_canBeLeft()
end

function ExpRefNode:canBeRight(  )

   return self:get_symbolInfo():get_canBeRight()
end

function NodeKind.get_ExpOp2(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpOp2'])
end


regKind( [[ExpOp2]] )
function Filter:processExpOp2( node, ... )

end


function NodeManager:getExpOp2NodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOp2']) )
end


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
   ExpOp2Node.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, op, exp1, exp2 ); end
   return obj
end
function ExpOp2Node:__init(pos, typeList, op, exp1, exp2) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpOp2']), pos, typeList)
   
   
   self.op = op
   self.exp1 = exp1
   self.exp2 = exp2
   
end
function ExpOp2Node.create( nodeMan, pos, typeList, op, exp1, exp2 )

   local node = ExpOp2Node.new(pos, typeList, op, exp1, exp2)
   nodeMan:addNode( node )
   return node
end
function ExpOp2Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp2Node  } )
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


function ExpOp2Node:canBeStatement(  )

   return self:get_op().txt == '='
end

function NodeKind.get_UnwrapSet(  )

   return _lune.unwrap( _moduleObj.nodeKind['UnwrapSet'])
end


regKind( [[UnwrapSet]] )
function Filter:processUnwrapSet( node, ... )

end


function NodeManager:getUnwrapSetNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['UnwrapSet']) )
end


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
function UnwrapSetNode:canBeStatement(  )

   return true
end
function UnwrapSetNode.new( pos, typeList, dstExpList, srcExpList, unwrapBlock )
   local obj = {}
   UnwrapSetNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, dstExpList, srcExpList, unwrapBlock ); end
   return obj
end
function UnwrapSetNode:__init(pos, typeList, dstExpList, srcExpList, unwrapBlock) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['UnwrapSet']), pos, typeList)
   
   
   self.dstExpList = dstExpList
   self.srcExpList = srcExpList
   self.unwrapBlock = unwrapBlock
   
end
function UnwrapSetNode.create( nodeMan, pos, typeList, dstExpList, srcExpList, unwrapBlock )

   local node = UnwrapSetNode.new(pos, typeList, dstExpList, srcExpList, unwrapBlock)
   nodeMan:addNode( node )
   return node
end
function UnwrapSetNode.setmeta( obj )
  setmetatable( obj, { __index = UnwrapSetNode  } )
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


function NodeKind.get_IfUnwrap(  )

   return _lune.unwrap( _moduleObj.nodeKind['IfUnwrap'])
end


regKind( [[IfUnwrap]] )
function Filter:processIfUnwrap( node, ... )

end


function NodeManager:getIfUnwrapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['IfUnwrap']) )
end


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
function IfUnwrapNode:canBeStatement(  )

   return true
end
function IfUnwrapNode.new( pos, typeList, varNameList, expNodeList, block, nilBlock )
   local obj = {}
   IfUnwrapNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, varNameList, expNodeList, block, nilBlock ); end
   return obj
end
function IfUnwrapNode:__init(pos, typeList, varNameList, expNodeList, block, nilBlock) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['IfUnwrap']), pos, typeList)
   
   
   self.varNameList = varNameList
   self.expNodeList = expNodeList
   self.block = block
   self.nilBlock = nilBlock
   
end
function IfUnwrapNode.create( nodeMan, pos, typeList, varNameList, expNodeList, block, nilBlock )

   local node = IfUnwrapNode.new(pos, typeList, varNameList, expNodeList, block, nilBlock)
   nodeMan:addNode( node )
   return node
end
function IfUnwrapNode.setmeta( obj )
  setmetatable( obj, { __index = IfUnwrapNode  } )
end
function IfUnwrapNode:get_varNameList()       
   return self.varNameList         
end
function IfUnwrapNode:get_expNodeList()       
   return self.expNodeList         
end
function IfUnwrapNode:get_block()       
   return self.block         
end
function IfUnwrapNode:get_nilBlock()       
   return self.nilBlock         
end


function IfUnwrapNode:getBreakKind( checkMode )

   local kind = self.block:getBreakKind( checkMode )
   local work = kind
   if checkMode == CheckBreakMode.IgnoreFlowReturn then
      if work == BreakKind.Return then
         return BreakKind.Return
      end
      
      if work == BreakKind.NeverRet then
         return BreakKind.NeverRet
      end
      
   else
    
      do
         local _switchExp = work
         if _switchExp == BreakKind.None then
            if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
               return BreakKind.None
            end
            
         else 
            
               if kind == BreakKind.None or kind > work then
                  kind = work
               end
               
         end
      end
      
   end
   
   
   do
      local block = self.nilBlock
      if block ~= nil then
         work = block:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     return BreakKind.None
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   return BreakKind.None
end

function NodeKind.get_When(  )

   return _lune.unwrap( _moduleObj.nodeKind['When'])
end


regKind( [[When]] )
function Filter:processWhen( node, ... )

end


function NodeManager:getWhenNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['When']) )
end


local WhenNode = {}
setmetatable( WhenNode, { __index = Node } )
_moduleObj.WhenNode = WhenNode
function WhenNode:processFilter( filter, ... )

   local argList = {...}
   filter:processWhen( self, table.unpack( argList ) )
end
function WhenNode:canBeRight(  )

   return false
end
function WhenNode:canBeLeft(  )

   return false
end
function WhenNode:canBeStatement(  )

   return true
end
function WhenNode.new( pos, typeList, varNameList, expNodeList, block, elseBlock )
   local obj = {}
   WhenNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, varNameList, expNodeList, block, elseBlock ); end
   return obj
end
function WhenNode:__init(pos, typeList, varNameList, expNodeList, block, elseBlock) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['When']), pos, typeList)
   
   
   self.varNameList = varNameList
   self.expNodeList = expNodeList
   self.block = block
   self.elseBlock = elseBlock
   
end
function WhenNode.create( nodeMan, pos, typeList, varNameList, expNodeList, block, elseBlock )

   local node = WhenNode.new(pos, typeList, varNameList, expNodeList, block, elseBlock)
   nodeMan:addNode( node )
   return node
end
function WhenNode.setmeta( obj )
  setmetatable( obj, { __index = WhenNode  } )
end
function WhenNode:get_varNameList()       
   return self.varNameList         
end
function WhenNode:get_expNodeList()       
   return self.expNodeList         
end
function WhenNode:get_block()       
   return self.block         
end
function WhenNode:get_elseBlock()       
   return self.elseBlock         
end


function WhenNode:getBreakKind( checkMode )

   local kind = self.block:getBreakKind( checkMode )
   local work = kind
   if checkMode == CheckBreakMode.IgnoreFlowReturn then
      if work == BreakKind.Return then
         return BreakKind.Return
      end
      
      if work == BreakKind.NeverRet then
         return BreakKind.NeverRet
      end
      
   else
    
      do
         local _switchExp = work
         if _switchExp == BreakKind.None then
            if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
               return BreakKind.None
            end
            
         else 
            
               if kind == BreakKind.None or kind > work then
                  kind = work
               end
               
         end
      end
      
   end
   
   
   do
      local block = self.elseBlock
      if block ~= nil then
         work = block:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     return BreakKind.None
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
         return kind
      end
   end
   
   return BreakKind.None
end

function NodeKind.get_ExpCast(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpCast'])
end


regKind( [[ExpCast]] )
function Filter:processExpCast( node, ... )

end


function NodeManager:getExpCastNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCast']) )
end


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
function ExpCastNode:canBeStatement(  )

   return false
end
function ExpCastNode.new( pos, typeList, exp )
   local obj = {}
   ExpCastNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp ); end
   return obj
end
function ExpCastNode:__init(pos, typeList, exp) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpCast']), pos, typeList)
   
   
   self.exp = exp
   
end
function ExpCastNode.create( nodeMan, pos, typeList, exp )

   local node = ExpCastNode.new(pos, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function ExpCastNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCastNode  } )
end
function ExpCastNode:get_exp()       
   return self.exp         
end


local MacroMode = {}
_moduleObj.MacroMode = MacroMode
MacroMode._val2NameMap = {}
function MacroMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MacroMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function MacroMode._from( val )
   if MacroMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
MacroMode.__allList = {}
function MacroMode.get__allList()
   return MacroMode.__allList
end

MacroMode.None = 0
MacroMode._val2NameMap[0] = 'None'
MacroMode.__allList[1] = MacroMode.None
MacroMode.Expand = 1
MacroMode._val2NameMap[1] = 'Expand'
MacroMode.__allList[2] = MacroMode.Expand
MacroMode.Analyze = 2
MacroMode._val2NameMap[2] = 'Analyze'
MacroMode.__allList[3] = MacroMode.Analyze

function NodeKind.get_ExpOp1(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpOp1'])
end


regKind( [[ExpOp1]] )
function Filter:processExpOp1( node, ... )

end


function NodeManager:getExpOp1NodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOp1']) )
end


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
function ExpOp1Node:canBeStatement(  )

   return false
end
function ExpOp1Node.new( pos, typeList, op, macroMode, exp )
   local obj = {}
   ExpOp1Node.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, op, macroMode, exp ); end
   return obj
end
function ExpOp1Node:__init(pos, typeList, op, macroMode, exp) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpOp1']), pos, typeList)
   
   
   self.op = op
   self.macroMode = macroMode
   self.exp = exp
   
end
function ExpOp1Node.create( nodeMan, pos, typeList, op, macroMode, exp )

   local node = ExpOp1Node.new(pos, typeList, op, macroMode, exp)
   nodeMan:addNode( node )
   return node
end
function ExpOp1Node.setmeta( obj )
  setmetatable( obj, { __index = ExpOp1Node  } )
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


function NodeKind.get_ExpRefItem(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpRefItem'])
end


regKind( [[ExpRefItem]] )
function Filter:processExpRefItem( node, ... )

end


function NodeManager:getExpRefItemNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpRefItem']) )
end


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
function ExpRefItemNode:canBeStatement(  )

   return false
end
function ExpRefItemNode.new( pos, typeList, val, nilAccess, symbol, index )
   local obj = {}
   ExpRefItemNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, val, nilAccess, symbol, index ); end
   return obj
end
function ExpRefItemNode:__init(pos, typeList, val, nilAccess, symbol, index) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpRefItem']), pos, typeList)
   
   
   self.val = val
   self.nilAccess = nilAccess
   self.symbol = symbol
   self.index = index
   
end
function ExpRefItemNode.create( nodeMan, pos, typeList, val, nilAccess, symbol, index )

   local node = ExpRefItemNode.new(pos, typeList, val, nilAccess, symbol, index)
   nodeMan:addNode( node )
   return node
end
function ExpRefItemNode.setmeta( obj )
  setmetatable( obj, { __index = ExpRefItemNode  } )
end
function ExpRefItemNode:get_val()       
   return self.val         
end
function ExpRefItemNode:get_nilAccess()       
   return self.nilAccess         
end
function ExpRefItemNode:get_symbol()       
   return self.symbol         
end
function ExpRefItemNode:get_index()       
   return self.index         
end


function ExpRefItemNode:canBeLeft(  )

   if self.val:get_expType() == _moduleObj.builtinTypeStem then
      return false
   end
   
   return self:get_val():get_expType():get_mutable() and not self.nilAccess
end

function NodeKind.get_ExpCall(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpCall'])
end


regKind( [[ExpCall]] )
function Filter:processExpCall( node, ... )

end


function NodeManager:getExpCallNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCall']) )
end


local ExpCallNode = {}
setmetatable( ExpCallNode, { __index = Node } )
_moduleObj.ExpCallNode = ExpCallNode
function ExpCallNode:processFilter( filter, ... )

   local argList = {...}
   filter:processExpCall( self, table.unpack( argList ) )
end
function ExpCallNode:canBeLeft(  )

   return false
end
function ExpCallNode:canBeStatement(  )

   return true
end
function ExpCallNode.new( pos, typeList, func, errorFunc, nilAccess, argList )
   local obj = {}
   ExpCallNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, func, errorFunc, nilAccess, argList ); end
   return obj
end
function ExpCallNode:__init(pos, typeList, func, errorFunc, nilAccess, argList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpCall']), pos, typeList)
   
   
   self.func = func
   self.errorFunc = errorFunc
   self.nilAccess = nilAccess
   self.argList = argList
   
end
function ExpCallNode.create( nodeMan, pos, typeList, func, errorFunc, nilAccess, argList )

   local node = ExpCallNode.new(pos, typeList, func, errorFunc, nilAccess, argList)
   nodeMan:addNode( node )
   return node
end
function ExpCallNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCallNode  } )
end
function ExpCallNode:get_func()       
   return self.func         
end
function ExpCallNode:get_errorFunc()       
   return self.errorFunc         
end
function ExpCallNode:get_nilAccess()       
   return self.nilAccess         
end
function ExpCallNode:get_argList()       
   return self.argList         
end


function ExpCallNode:canBeRight(  )

   local expType = self:get_expType()
   if expType:equals( _moduleObj.builtinTypeNone ) or expType:equals( _moduleObj.builtinTypeNeverRet ) then
      return false
   end
   
   return true
end

function ExpCallNode:getBreakKind( checkMode )

   if self.errorFunc then
      return BreakKind.NeverRet
   end
   
   return BreakKind.None
end

function NodeKind.get_ExpDDD(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpDDD'])
end


regKind( [[ExpDDD]] )
function Filter:processExpDDD( node, ... )

end


function NodeManager:getExpDDDNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpDDD']) )
end


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
function ExpDDDNode:canBeStatement(  )

   return false
end
function ExpDDDNode.new( pos, typeList, token )
   local obj = {}
   ExpDDDNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token ); end
   return obj
end
function ExpDDDNode:__init(pos, typeList, token) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpDDD']), pos, typeList)
   
   
   self.token = token
   
end
function ExpDDDNode.create( nodeMan, pos, typeList, token )

   local node = ExpDDDNode.new(pos, typeList, token)
   nodeMan:addNode( node )
   return node
end
function ExpDDDNode.setmeta( obj )
  setmetatable( obj, { __index = ExpDDDNode  } )
end
function ExpDDDNode:get_token()       
   return self.token         
end


function NodeKind.get_ExpParen(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpParen'])
end


regKind( [[ExpParen]] )
function Filter:processExpParen( node, ... )

end


function NodeManager:getExpParenNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpParen']) )
end


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
function ExpParenNode:canBeStatement(  )

   return false
end
function ExpParenNode.new( pos, typeList, exp )
   local obj = {}
   ExpParenNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp ); end
   return obj
end
function ExpParenNode:__init(pos, typeList, exp) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpParen']), pos, typeList)
   
   
   self.exp = exp
   
end
function ExpParenNode.create( nodeMan, pos, typeList, exp )

   local node = ExpParenNode.new(pos, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function ExpParenNode.setmeta( obj )
  setmetatable( obj, { __index = ExpParenNode  } )
end
function ExpParenNode:get_exp()       
   return self.exp         
end


function NodeKind.get_ExpMacroExp(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpMacroExp'])
end


regKind( [[ExpMacroExp]] )
function Filter:processExpMacroExp( node, ... )

end


function NodeManager:getExpMacroExpNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpMacroExp']) )
end


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
function ExpMacroExpNode:canBeStatement(  )

   return true
end
function ExpMacroExpNode.new( pos, typeList, stmtList )
   local obj = {}
   ExpMacroExpNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, stmtList ); end
   return obj
end
function ExpMacroExpNode:__init(pos, typeList, stmtList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpMacroExp']), pos, typeList)
   
   
   self.stmtList = stmtList
   
end
function ExpMacroExpNode.create( nodeMan, pos, typeList, stmtList )

   local node = ExpMacroExpNode.new(pos, typeList, stmtList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroExpNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroExpNode  } )
end
function ExpMacroExpNode:get_stmtList()       
   return self.stmtList         
end


function ExpMacroExpNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in pairs( self.stmtList ) do
         local work = stmt:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      return kind
   else
    
      if #self.stmtList > 0 then
         return self.stmtList[#self.stmtList]:getBreakKind( checkMode )
      end
      
   end
   
   return BreakKind.None
end

function NodeKind.get_ExpMacroStat(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpMacroStat'])
end


regKind( [[ExpMacroStat]] )
function Filter:processExpMacroStat( node, ... )

end


function NodeManager:getExpMacroStatNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpMacroStat']) )
end


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
function ExpMacroStatNode:canBeStatement(  )

   return false
end
function ExpMacroStatNode.new( pos, typeList, expStrList )
   local obj = {}
   ExpMacroStatNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, expStrList ); end
   return obj
end
function ExpMacroStatNode:__init(pos, typeList, expStrList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpMacroStat']), pos, typeList)
   
   
   self.expStrList = expStrList
   
end
function ExpMacroStatNode.create( nodeMan, pos, typeList, expStrList )

   local node = ExpMacroStatNode.new(pos, typeList, expStrList)
   nodeMan:addNode( node )
   return node
end
function ExpMacroStatNode.setmeta( obj )
  setmetatable( obj, { __index = ExpMacroStatNode  } )
end
function ExpMacroStatNode:get_expStrList()       
   return self.expStrList         
end


function NodeKind.get_StmtExp(  )

   return _lune.unwrap( _moduleObj.nodeKind['StmtExp'])
end


regKind( [[StmtExp]] )
function Filter:processStmtExp( node, ... )

end


function NodeManager:getStmtExpNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['StmtExp']) )
end


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
   StmtExpNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, exp ); end
   return obj
end
function StmtExpNode:__init(pos, typeList, exp) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['StmtExp']), pos, typeList)
   
   
   self.exp = exp
   
end
function StmtExpNode.create( nodeMan, pos, typeList, exp )

   local node = StmtExpNode.new(pos, typeList, exp)
   nodeMan:addNode( node )
   return node
end
function StmtExpNode.setmeta( obj )
  setmetatable( obj, { __index = StmtExpNode  } )
end
function StmtExpNode:get_exp()       
   return self.exp         
end


function StmtExpNode:canBeStatement(  )

   return self:get_exp():canBeStatement(  )
end

function StmtExpNode:getBreakKind( checkMode )

   return self:get_exp():getBreakKind( checkMode )
end

function NodeKind.get_ExpOmitEnum(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum'])
end


regKind( [[ExpOmitEnum]] )
function Filter:processExpOmitEnum( node, ... )

end


function NodeManager:getExpOmitEnumNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum']) )
end


local ExpOmitEnumNode = {}
setmetatable( ExpOmitEnumNode, { __index = Node } )
_moduleObj.ExpOmitEnumNode = ExpOmitEnumNode
function ExpOmitEnumNode:processFilter( filter, ... )

   local argList = {...}
   filter:processExpOmitEnum( self, table.unpack( argList ) )
end
function ExpOmitEnumNode:canBeRight(  )

   return true
end
function ExpOmitEnumNode:canBeLeft(  )

   return true
end
function ExpOmitEnumNode:canBeStatement(  )

   return false
end
function ExpOmitEnumNode.new( pos, typeList, valToken, enumTypeInfo )
   local obj = {}
   ExpOmitEnumNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, valToken, enumTypeInfo ); end
   return obj
end
function ExpOmitEnumNode:__init(pos, typeList, valToken, enumTypeInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpOmitEnum']), pos, typeList)
   
   
   self.valToken = valToken
   self.enumTypeInfo = enumTypeInfo
   
end
function ExpOmitEnumNode.create( nodeMan, pos, typeList, valToken, enumTypeInfo )

   local node = ExpOmitEnumNode.new(pos, typeList, valToken, enumTypeInfo)
   nodeMan:addNode( node )
   return node
end
function ExpOmitEnumNode.setmeta( obj )
  setmetatable( obj, { __index = ExpOmitEnumNode  } )
end
function ExpOmitEnumNode:get_valToken()       
   return self.valToken         
end
function ExpOmitEnumNode:get_enumTypeInfo()       
   return self.enumTypeInfo         
end


function NodeKind.get_RefField(  )

   return _lune.unwrap( _moduleObj.nodeKind['RefField'])
end


regKind( [[RefField]] )
function Filter:processRefField( node, ... )

end


function NodeManager:getRefFieldNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['RefField']) )
end


local RefFieldNode = {}
setmetatable( RefFieldNode, { __index = Node } )
_moduleObj.RefFieldNode = RefFieldNode
function RefFieldNode:processFilter( filter, ... )

   local argList = {...}
   filter:processRefField( self, table.unpack( argList ) )
end
function RefFieldNode:canBeStatement(  )

   return false
end
function RefFieldNode.new( pos, typeList, field, symbolInfo, nilAccess, prefix )
   local obj = {}
   RefFieldNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, field, symbolInfo, nilAccess, prefix ); end
   return obj
end
function RefFieldNode:__init(pos, typeList, field, symbolInfo, nilAccess, prefix) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['RefField']), pos, typeList)
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   
end
function RefFieldNode.create( nodeMan, pos, typeList, field, symbolInfo, nilAccess, prefix )

   local node = RefFieldNode.new(pos, typeList, field, symbolInfo, nilAccess, prefix)
   nodeMan:addNode( node )
   return node
end
function RefFieldNode.setmeta( obj )
  setmetatable( obj, { __index = RefFieldNode  } )
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


function RefFieldNode:canBeLeft(  )

   do
      local _exp = self:get_symbolInfo()
      if _exp ~= nil then
         return _exp:get_canBeLeft()
      end
   end
   
   return false
end

function RefFieldNode:canBeRight(  )

   do
      local _exp = self:get_symbolInfo()
      if _exp ~= nil then
         return _exp:get_canBeRight()
      end
   end
   
   return true
end

function NodeKind.get_GetField(  )

   return _lune.unwrap( _moduleObj.nodeKind['GetField'])
end


regKind( [[GetField]] )
function Filter:processGetField( node, ... )

end


function NodeManager:getGetFieldNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['GetField']) )
end


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
function GetFieldNode:canBeStatement(  )

   return false
end
function GetFieldNode.new( pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )
   local obj = {}
   GetFieldNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo ); end
   return obj
end
function GetFieldNode:__init(pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['GetField']), pos, typeList)
   
   
   self.field = field
   self.symbolInfo = symbolInfo
   self.nilAccess = nilAccess
   self.prefix = prefix
   self.getterTypeInfo = getterTypeInfo
   
end
function GetFieldNode.create( nodeMan, pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo )

   local node = GetFieldNode.new(pos, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo)
   nodeMan:addNode( node )
   return node
end
function GetFieldNode.setmeta( obj )
  setmetatable( obj, { __index = GetFieldNode  } )
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
function VarInfo.setmeta( obj )
  setmetatable( obj, { __index = VarInfo  } )
end
function VarInfo.new( name, refType, actualType )
   local obj = {}
   VarInfo.setmeta( obj )
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

local DeclVarMode = {}
_moduleObj.DeclVarMode = DeclVarMode
DeclVarMode._val2NameMap = {}
function DeclVarMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclVarMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function DeclVarMode._from( val )
   if DeclVarMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
DeclVarMode.__allList = {}
function DeclVarMode.get__allList()
   return DeclVarMode.__allList
end

DeclVarMode.Let = 0
DeclVarMode._val2NameMap[0] = 'Let'
DeclVarMode.__allList[1] = DeclVarMode.Let
DeclVarMode.Sync = 1
DeclVarMode._val2NameMap[1] = 'Sync'
DeclVarMode.__allList[2] = DeclVarMode.Sync
DeclVarMode.Unwrap = 2
DeclVarMode._val2NameMap[2] = 'Unwrap'
DeclVarMode.__allList[3] = DeclVarMode.Unwrap

function NodeKind.get_DeclVar(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclVar'])
end


regKind( [[DeclVar]] )
function Filter:processDeclVar( node, ... )

end


function NodeManager:getDeclVarNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclVar']) )
end


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
function DeclVarNode:canBeStatement(  )

   return true
end
function DeclVarNode.new( pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )
   local obj = {}
   DeclVarNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock ); end
   return obj
end
function DeclVarNode:__init(pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclVar']), pos, typeList)
   
   
   self.mode = mode
   self.accessMode = accessMode
   self.staticFlag = staticFlag
   self.varList = varList
   self.expList = expList
   self.symbolInfoList = symbolInfoList
   self.typeInfoList = typeInfoList
   self.unwrapFlag = unwrapFlag
   self.unwrapBlock = unwrapBlock
   self.thenBlock = thenBlock
   self.syncVarList = syncVarList
   self.syncBlock = syncBlock
   
end
function DeclVarNode.create( nodeMan, pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock )

   local node = DeclVarNode.new(pos, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock)
   nodeMan:addNode( node )
   return node
end
function DeclVarNode.setmeta( obj )
  setmetatable( obj, { __index = DeclVarNode  } )
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
function DeclVarNode:get_symbolInfoList()       
   return self.symbolInfoList         
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


function DeclVarNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      do
         local block = self.unwrapBlock
         if block ~= nil then
            local kind = block:getBreakKind( checkMode )
            do
               local _switchExp = kind
               if _switchExp == BreakKind.Return or _switchExp == BreakKind.NeverRet then
                  return kind
               end
            end
            
         end
      end
      
      do
         local block = self.thenBlock
         if block ~= nil then
            local kind = block:getBreakKind( checkMode )
            do
               local _switchExp = kind
               if _switchExp == BreakKind.Return or _switchExp == BreakKind.NeverRet then
                  return kind
               end
            end
            
         end
      end
      
      do
         local block = self.syncBlock
         if block ~= nil then
            local kind = block:getBreakKind( checkMode )
            do
               local _switchExp = kind
               if _switchExp == BreakKind.Return or _switchExp == BreakKind.NeverRet then
                  return kind
               end
            end
            
         end
      end
      
      return BreakKind.None
   else
    
      local kind = BreakKind.None
      local work = BreakKind.None
      do
         local block = self.unwrapBlock
         if block ~= nil then
            work = block:getBreakKind( checkMode )
            if checkMode == CheckBreakMode.IgnoreFlowReturn then
               if work == BreakKind.Return then
                  return BreakKind.Return
               end
               
               if work == BreakKind.NeverRet then
                  return BreakKind.NeverRet
               end
               
            else
             
               do
                  local _switchExp = work
                  if _switchExp == BreakKind.None then
                     if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                        return BreakKind.None
                     end
                     
                  else 
                     
                        if kind == BreakKind.None or kind > work then
                           kind = work
                        end
                        
                  end
               end
               
            end
            
            
            do
               local thenBlock = self.thenBlock
               if thenBlock ~= nil then
                  work = thenBlock:getBreakKind( checkMode )
                  if checkMode == CheckBreakMode.IgnoreFlowReturn then
                     if work == BreakKind.Return then
                        return BreakKind.Return
                     end
                     
                     if work == BreakKind.NeverRet then
                        return BreakKind.NeverRet
                     end
                     
                  else
                   
                     do
                        local _switchExp = work
                        if _switchExp == BreakKind.None then
                           if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                              return BreakKind.None
                           end
                           
                        else 
                           
                              if kind == BreakKind.None or kind > work then
                                 kind = work
                              end
                              
                        end
                     end
                     
                  end
                  
                  
                  do
                     local syncBlock = self.syncBlock
                     if syncBlock ~= nil then
                        work = syncBlock:getBreakKind( checkMode )
                        if checkMode == CheckBreakMode.IgnoreFlowReturn then
                           if work == BreakKind.Return then
                              return BreakKind.Return
                           end
                           
                           if work == BreakKind.NeverRet then
                              return BreakKind.NeverRet
                           end
                           
                        else
                         
                           do
                              local _switchExp = work
                              if _switchExp == BreakKind.None then
                                 if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                                    return BreakKind.None
                                 end
                                 
                              else 
                                 
                                    if kind == BreakKind.None or kind > work then
                                       kind = work
                                    end
                                    
                              end
                           end
                           
                        end
                        
                        
                        return kind
                     end
                  end
                  
               end
            end
            
         end
      end
      
      return BreakKind.None
   end
   
end

local DeclFuncInfo = {}
_moduleObj.DeclFuncInfo = DeclFuncInfo
function DeclFuncInfo.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncInfo  } )
end
function DeclFuncInfo.new( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol )
   local obj = {}
   DeclFuncInfo.setmeta( obj )
   if obj.__init then
      obj:__init( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol )
   end        
   return obj 
end         
function DeclFuncInfo:__init( classTypeInfo, name, argList, staticFlag, accessMode, body, retTypeInfoList, has__func__Symbol ) 

   self.classTypeInfo = classTypeInfo
   self.name = name
   self.argList = argList
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.body = body
   self.retTypeInfoList = retTypeInfoList
   self.has__func__Symbol = has__func__Symbol
end
function DeclFuncInfo:get_classTypeInfo()       
   return self.classTypeInfo         
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
function DeclFuncInfo:get_has__func__Symbol()       
   return self.has__func__Symbol         
end

function NodeKind.get_DeclFunc(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclFunc'])
end


regKind( [[DeclFunc]] )
function Filter:processDeclFunc( node, ... )

end


function NodeManager:getDeclFuncNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclFunc']) )
end


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
function DeclFuncNode:canBeStatement(  )

   return true
end
function DeclFuncNode.new( pos, typeList, declInfo )
   local obj = {}
   DeclFuncNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, declInfo ); end
   return obj
end
function DeclFuncNode:__init(pos, typeList, declInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclFunc']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclFuncNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclFuncNode.new(pos, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclFuncNode.setmeta( obj )
  setmetatable( obj, { __index = DeclFuncNode  } )
end
function DeclFuncNode:get_declInfo()       
   return self.declInfo         
end


function NodeKind.get_DeclMethod(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclMethod'])
end


regKind( [[DeclMethod]] )
function Filter:processDeclMethod( node, ... )

end


function NodeManager:getDeclMethodNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMethod']) )
end


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
function DeclMethodNode:canBeStatement(  )

   return true
end
function DeclMethodNode.new( pos, typeList, declInfo )
   local obj = {}
   DeclMethodNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, declInfo ); end
   return obj
end
function DeclMethodNode:__init(pos, typeList, declInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclMethod']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclMethodNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclMethodNode.new(pos, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclMethodNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMethodNode  } )
end
function DeclMethodNode:get_declInfo()       
   return self.declInfo         
end


function NodeKind.get_DeclConstr(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclConstr'])
end


regKind( [[DeclConstr]] )
function Filter:processDeclConstr( node, ... )

end


function NodeManager:getDeclConstrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclConstr']) )
end


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
function DeclConstrNode:canBeStatement(  )

   return true
end
function DeclConstrNode.new( pos, typeList, declInfo )
   local obj = {}
   DeclConstrNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, declInfo ); end
   return obj
end
function DeclConstrNode:__init(pos, typeList, declInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclConstr']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclConstrNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclConstrNode.new(pos, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclConstrNode.setmeta( obj )
  setmetatable( obj, { __index = DeclConstrNode  } )
end
function DeclConstrNode:get_declInfo()       
   return self.declInfo         
end


function NodeKind.get_DeclDestr(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclDestr'])
end


regKind( [[DeclDestr]] )
function Filter:processDeclDestr( node, ... )

end


function NodeManager:getDeclDestrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclDestr']) )
end


local DeclDestrNode = {}
setmetatable( DeclDestrNode, { __index = Node } )
_moduleObj.DeclDestrNode = DeclDestrNode
function DeclDestrNode:processFilter( filter, ... )

   local argList = {...}
   filter:processDeclDestr( self, table.unpack( argList ) )
end
function DeclDestrNode:canBeRight(  )

   return false
end
function DeclDestrNode:canBeLeft(  )

   return false
end
function DeclDestrNode:canBeStatement(  )

   return true
end
function DeclDestrNode.new( pos, typeList, declInfo )
   local obj = {}
   DeclDestrNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, declInfo ); end
   return obj
end
function DeclDestrNode:__init(pos, typeList, declInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclDestr']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclDestrNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclDestrNode.new(pos, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclDestrNode.setmeta( obj )
  setmetatable( obj, { __index = DeclDestrNode  } )
end
function DeclDestrNode:get_declInfo()       
   return self.declInfo         
end


function NodeKind.get_ExpCallSuper(  )

   return _lune.unwrap( _moduleObj.nodeKind['ExpCallSuper'])
end


regKind( [[ExpCallSuper]] )
function Filter:processExpCallSuper( node, ... )

end


function NodeManager:getExpCallSuperNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['ExpCallSuper']) )
end


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
function ExpCallSuperNode:canBeStatement(  )

   return true
end
function ExpCallSuperNode.new( pos, typeList, superType, methodType, expList )
   local obj = {}
   ExpCallSuperNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, superType, methodType, expList ); end
   return obj
end
function ExpCallSuperNode:__init(pos, typeList, superType, methodType, expList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['ExpCallSuper']), pos, typeList)
   
   
   self.superType = superType
   self.methodType = methodType
   self.expList = expList
   
end
function ExpCallSuperNode.create( nodeMan, pos, typeList, superType, methodType, expList )

   local node = ExpCallSuperNode.new(pos, typeList, superType, methodType, expList)
   nodeMan:addNode( node )
   return node
end
function ExpCallSuperNode.setmeta( obj )
  setmetatable( obj, { __index = ExpCallSuperNode  } )
end
function ExpCallSuperNode:get_superType()       
   return self.superType         
end
function ExpCallSuperNode:get_methodType()       
   return self.methodType         
end
function ExpCallSuperNode:get_expList()       
   return self.expList         
end


function NodeKind.get_DeclMember(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclMember'])
end


regKind( [[DeclMember]] )
function Filter:processDeclMember( node, ... )

end


function NodeManager:getDeclMemberNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMember']) )
end


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
function DeclMemberNode:canBeStatement(  )

   return true
end
function DeclMemberNode.new( pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode )
   local obj = {}
   DeclMemberNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode ); end
   return obj
end
function DeclMemberNode:__init(pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclMember']), pos, typeList)
   
   
   self.name = name
   self.refType = refType
   self.symbolInfo = symbolInfo
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.getterMutable = getterMutable
   self.getterMode = getterMode
   self.setterMode = setterMode
   
end
function DeclMemberNode.create( nodeMan, pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode )

   local node = DeclMemberNode.new(pos, typeList, name, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode)
   nodeMan:addNode( node )
   return node
end
function DeclMemberNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMemberNode  } )
end
function DeclMemberNode:get_name()       
   return self.name         
end
function DeclMemberNode:get_refType()       
   return self.refType         
end
function DeclMemberNode:get_symbolInfo()       
   return self.symbolInfo         
end
function DeclMemberNode:get_staticFlag()       
   return self.staticFlag         
end
function DeclMemberNode:get_accessMode()       
   return self.accessMode         
end
function DeclMemberNode:get_getterMutable()       
   return self.getterMutable         
end
function DeclMemberNode:get_getterMode()       
   return self.getterMode         
end
function DeclMemberNode:get_setterMode()       
   return self.setterMode         
end


function NodeKind.get_DeclArg(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclArg'])
end


regKind( [[DeclArg]] )
function Filter:processDeclArg( node, ... )

end


function NodeManager:getDeclArgNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclArg']) )
end


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
function DeclArgNode:canBeStatement(  )

   return false
end
function DeclArgNode.new( pos, typeList, name, argType )
   local obj = {}
   DeclArgNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, name, argType ); end
   return obj
end
function DeclArgNode:__init(pos, typeList, name, argType) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclArg']), pos, typeList)
   
   
   self.name = name
   self.argType = argType
   
end
function DeclArgNode.create( nodeMan, pos, typeList, name, argType )

   local node = DeclArgNode.new(pos, typeList, name, argType)
   nodeMan:addNode( node )
   return node
end
function DeclArgNode.setmeta( obj )
  setmetatable( obj, { __index = DeclArgNode  } )
end
function DeclArgNode:get_name()       
   return self.name         
end
function DeclArgNode:get_argType()       
   return self.argType         
end


function NodeKind.get_DeclArgDDD(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclArgDDD'])
end


regKind( [[DeclArgDDD]] )
function Filter:processDeclArgDDD( node, ... )

end


function NodeManager:getDeclArgDDDNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclArgDDD']) )
end


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
function DeclArgDDDNode:canBeStatement(  )

   return false
end
function DeclArgDDDNode.new( pos, typeList )
   local obj = {}
   DeclArgDDDNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList ); end
   return obj
end
function DeclArgDDDNode:__init(pos, typeList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclArgDDD']), pos, typeList)
   
   
   
end
function DeclArgDDDNode.create( nodeMan, pos, typeList )

   local node = DeclArgDDDNode.new(pos, typeList)
   nodeMan:addNode( node )
   return node
end
function DeclArgDDDNode.setmeta( obj )
  setmetatable( obj, { __index = DeclArgDDDNode  } )
end


local AdvertiseInfo = {}
_moduleObj.AdvertiseInfo = AdvertiseInfo
function AdvertiseInfo.setmeta( obj )
  setmetatable( obj, { __index = AdvertiseInfo  } )
end
function AdvertiseInfo.new( member, prefix )
   local obj = {}
   AdvertiseInfo.setmeta( obj )
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


function NodeKind.get_DeclClass(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclClass'])
end


regKind( [[DeclClass]] )
function Filter:processDeclClass( node, ... )

end


function NodeManager:getDeclClassNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclClass']) )
end


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
function DeclClassNode:canBeStatement(  )

   return true
end
function DeclClassNode.new( pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet )
   local obj = {}
   DeclClassNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet ); end
   return obj
end
function DeclClassNode:__init(pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclClass']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.name = name
   self.gluePrefix = gluePrefix
   self.declStmtList = declStmtList
   self.fieldList = fieldList
   self.moduleName = moduleName
   self.memberList = memberList
   self.scope = scope
   self.initStmtList = initStmtList
   self.advertiseList = advertiseList
   self.trustList = trustList
   self.outerMethodSet = outerMethodSet
   
end
function DeclClassNode.create( nodeMan, pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet )

   local node = DeclClassNode.new(pos, typeList, accessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, scope, initStmtList, advertiseList, trustList, outerMethodSet)
   nodeMan:addNode( node )
   return node
end
function DeclClassNode.setmeta( obj )
  setmetatable( obj, { __index = DeclClassNode  } )
end
function DeclClassNode:get_accessMode()       
   return self.accessMode         
end
function DeclClassNode:get_name()       
   return self.name         
end
function DeclClassNode:get_gluePrefix()       
   return self.gluePrefix         
end
function DeclClassNode:get_declStmtList()       
   return self.declStmtList         
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


function NodeKind.get_DeclEnum(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclEnum'])
end


regKind( [[DeclEnum]] )
function Filter:processDeclEnum( node, ... )

end


function NodeManager:getDeclEnumNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclEnum']) )
end


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
function DeclEnumNode:canBeStatement(  )

   return true
end
function DeclEnumNode.new( pos, typeList, accessMode, name, valueNameList, scope )
   local obj = {}
   DeclEnumNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, accessMode, name, valueNameList, scope ); end
   return obj
end
function DeclEnumNode:__init(pos, typeList, accessMode, name, valueNameList, scope) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclEnum']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.name = name
   self.valueNameList = valueNameList
   self.scope = scope
   
end
function DeclEnumNode.create( nodeMan, pos, typeList, accessMode, name, valueNameList, scope )

   local node = DeclEnumNode.new(pos, typeList, accessMode, name, valueNameList, scope)
   nodeMan:addNode( node )
   return node
end
function DeclEnumNode.setmeta( obj )
  setmetatable( obj, { __index = DeclEnumNode  } )
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


function NodeKind.get_DeclAlge(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclAlge'])
end


regKind( [[DeclAlge]] )
function Filter:processDeclAlge( node, ... )

end


function NodeManager:getDeclAlgeNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclAlge']) )
end


local DeclAlgeNode = {}
setmetatable( DeclAlgeNode, { __index = Node } )
_moduleObj.DeclAlgeNode = DeclAlgeNode
function DeclAlgeNode:processFilter( filter, ... )

   local argList = {...}
   filter:processDeclAlge( self, table.unpack( argList ) )
end
function DeclAlgeNode:canBeRight(  )

   return false
end
function DeclAlgeNode:canBeLeft(  )

   return false
end
function DeclAlgeNode:canBeStatement(  )

   return true
end
function DeclAlgeNode.new( pos, typeList, accessMode, algeType, scope )
   local obj = {}
   DeclAlgeNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, accessMode, algeType, scope ); end
   return obj
end
function DeclAlgeNode:__init(pos, typeList, accessMode, algeType, scope) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclAlge']), pos, typeList)
   
   
   self.accessMode = accessMode
   self.algeType = algeType
   self.scope = scope
   
end
function DeclAlgeNode.create( nodeMan, pos, typeList, accessMode, algeType, scope )

   local node = DeclAlgeNode.new(pos, typeList, accessMode, algeType, scope)
   nodeMan:addNode( node )
   return node
end
function DeclAlgeNode.setmeta( obj )
  setmetatable( obj, { __index = DeclAlgeNode  } )
end
function DeclAlgeNode:get_accessMode()       
   return self.accessMode         
end
function DeclAlgeNode:get_algeType()       
   return self.algeType         
end
function DeclAlgeNode:get_scope()       
   return self.scope         
end


function NodeKind.get_NewAlgeVal(  )

   return _lune.unwrap( _moduleObj.nodeKind['NewAlgeVal'])
end


regKind( [[NewAlgeVal]] )
function Filter:processNewAlgeVal( node, ... )

end


function NodeManager:getNewAlgeValNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['NewAlgeVal']) )
end


local NewAlgeValNode = {}
setmetatable( NewAlgeValNode, { __index = Node } )
_moduleObj.NewAlgeValNode = NewAlgeValNode
function NewAlgeValNode:processFilter( filter, ... )

   local argList = {...}
   filter:processNewAlgeVal( self, table.unpack( argList ) )
end
function NewAlgeValNode:canBeRight(  )

   return true
end
function NewAlgeValNode:canBeLeft(  )

   return false
end
function NewAlgeValNode:canBeStatement(  )

   return false
end
function NewAlgeValNode.new( pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList )
   local obj = {}
   NewAlgeValNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList ); end
   return obj
end
function NewAlgeValNode:__init(pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['NewAlgeVal']), pos, typeList)
   
   
   self.name = name
   self.prefix = prefix
   self.algeTypeInfo = algeTypeInfo
   self.valInfo = valInfo
   self.paramList = paramList
   
end
function NewAlgeValNode.create( nodeMan, pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList )

   local node = NewAlgeValNode.new(pos, typeList, name, prefix, algeTypeInfo, valInfo, paramList)
   nodeMan:addNode( node )
   return node
end
function NewAlgeValNode.setmeta( obj )
  setmetatable( obj, { __index = NewAlgeValNode  } )
end
function NewAlgeValNode:get_name()       
   return self.name         
end
function NewAlgeValNode:get_prefix()       
   return self.prefix         
end
function NewAlgeValNode:get_algeTypeInfo()       
   return self.algeTypeInfo         
end
function NewAlgeValNode:get_valInfo()       
   return self.valInfo         
end
function NewAlgeValNode:get_paramList()       
   return self.paramList         
end


local MatchCase = {}
_moduleObj.MatchCase = MatchCase
function MatchCase.setmeta( obj )
  setmetatable( obj, { __index = MatchCase  } )
end
function MatchCase.new( valInfo, valParamNameList, block )
   local obj = {}
   MatchCase.setmeta( obj )
   if obj.__init then
      obj:__init( valInfo, valParamNameList, block )
   end        
   return obj 
end         
function MatchCase:__init( valInfo, valParamNameList, block ) 

   self.valInfo = valInfo
   self.valParamNameList = valParamNameList
   self.block = block
end
function MatchCase:get_valInfo()       
   return self.valInfo         
end
function MatchCase:get_valParamNameList()       
   return self.valParamNameList         
end
function MatchCase:get_block()       
   return self.block         
end

function NodeKind.get_Match(  )

   return _lune.unwrap( _moduleObj.nodeKind['Match'])
end


regKind( [[Match]] )
function Filter:processMatch( node, ... )

end


function NodeManager:getMatchNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Match']) )
end


local MatchNode = {}
setmetatable( MatchNode, { __index = Node } )
_moduleObj.MatchNode = MatchNode
function MatchNode:processFilter( filter, ... )

   local argList = {...}
   filter:processMatch( self, table.unpack( argList ) )
end
function MatchNode:canBeRight(  )

   return false
end
function MatchNode:canBeLeft(  )

   return false
end
function MatchNode:canBeStatement(  )

   return true
end
function MatchNode.new( pos, typeList, val, algeTypeInfo, caseList, defaultBlock )
   local obj = {}
   MatchNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, val, algeTypeInfo, caseList, defaultBlock ); end
   return obj
end
function MatchNode:__init(pos, typeList, val, algeTypeInfo, caseList, defaultBlock) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Match']), pos, typeList)
   
   
   self.val = val
   self.algeTypeInfo = algeTypeInfo
   self.caseList = caseList
   self.defaultBlock = defaultBlock
   
end
function MatchNode.create( nodeMan, pos, typeList, val, algeTypeInfo, caseList, defaultBlock )

   local node = MatchNode.new(pos, typeList, val, algeTypeInfo, caseList, defaultBlock)
   nodeMan:addNode( node )
   return node
end
function MatchNode.setmeta( obj )
  setmetatable( obj, { __index = MatchNode  } )
end
function MatchNode:get_val()       
   return self.val         
end
function MatchNode:get_algeTypeInfo()       
   return self.algeTypeInfo         
end
function MatchNode:get_caseList()       
   return self.caseList         
end
function MatchNode:get_defaultBlock()       
   return self.defaultBlock         
end


function NodeKind.get_DeclMacro(  )

   return _lune.unwrap( _moduleObj.nodeKind['DeclMacro'])
end


regKind( [[DeclMacro]] )
function Filter:processDeclMacro( node, ... )

end


function NodeManager:getDeclMacroNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['DeclMacro']) )
end


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
function DeclMacroNode:canBeStatement(  )

   return true
end
function DeclMacroNode.new( pos, typeList, declInfo )
   local obj = {}
   DeclMacroNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, declInfo ); end
   return obj
end
function DeclMacroNode:__init(pos, typeList, declInfo) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['DeclMacro']), pos, typeList)
   
   
   self.declInfo = declInfo
   
end
function DeclMacroNode.create( nodeMan, pos, typeList, declInfo )

   local node = DeclMacroNode.new(pos, typeList, declInfo)
   nodeMan:addNode( node )
   return node
end
function DeclMacroNode.setmeta( obj )
  setmetatable( obj, { __index = DeclMacroNode  } )
end
function DeclMacroNode:get_declInfo()       
   return self.declInfo         
end


local MacroEval = {}
_moduleObj.MacroEval = MacroEval
function MacroEval.setmeta( obj )
  setmetatable( obj, { __index = MacroEval  } )
end
function MacroEval.new(  )
   local obj = {}
   MacroEval.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function MacroEval:__init(  ) 

end

function NodeKind.get_Abbr(  )

   return _lune.unwrap( _moduleObj.nodeKind['Abbr'])
end


regKind( [[Abbr]] )
function Filter:processAbbr( node, ... )

end


function NodeManager:getAbbrNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['Abbr']) )
end


local AbbrNode = {}
setmetatable( AbbrNode, { __index = Node } )
_moduleObj.AbbrNode = AbbrNode
function AbbrNode:processFilter( filter, ... )

   local argList = {...}
   filter:processAbbr( self, table.unpack( argList ) )
end
function AbbrNode:canBeRight(  )

   return false
end
function AbbrNode:canBeLeft(  )

   return false
end
function AbbrNode:canBeStatement(  )

   return false
end
function AbbrNode.new( pos, typeList )
   local obj = {}
   AbbrNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList ); end
   return obj
end
function AbbrNode:__init(pos, typeList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['Abbr']), pos, typeList)
   
   
   
end
function AbbrNode.create( nodeMan, pos, typeList )

   local node = AbbrNode.new(pos, typeList)
   nodeMan:addNode( node )
   return node
end
function AbbrNode.setmeta( obj )
  setmetatable( obj, { __index = AbbrNode  } )
end


function NodeKind.get_LiteralNil(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralNil'])
end


regKind( [[LiteralNil]] )
function Filter:processLiteralNil( node, ... )

end


function NodeManager:getLiteralNilNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralNil']) )
end


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
function LiteralNilNode:canBeStatement(  )

   return false
end
function LiteralNilNode.new( pos, typeList )
   local obj = {}
   LiteralNilNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList ); end
   return obj
end
function LiteralNilNode:__init(pos, typeList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralNil']), pos, typeList)
   
   
   
end
function LiteralNilNode.create( nodeMan, pos, typeList )

   local node = LiteralNilNode.new(pos, typeList)
   nodeMan:addNode( node )
   return node
end
function LiteralNilNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralNilNode  } )
end


function NodeKind.get_LiteralChar(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralChar'])
end


regKind( [[LiteralChar]] )
function Filter:processLiteralChar( node, ... )

end


function NodeManager:getLiteralCharNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralChar']) )
end


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
function LiteralCharNode:canBeStatement(  )

   return false
end
function LiteralCharNode.new( pos, typeList, token, num )
   local obj = {}
   LiteralCharNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token, num ); end
   return obj
end
function LiteralCharNode:__init(pos, typeList, token, num) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralChar']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralCharNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralCharNode.new(pos, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralCharNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralCharNode  } )
end
function LiteralCharNode:get_token()       
   return self.token         
end
function LiteralCharNode:get_num()       
   return self.num         
end


function NodeKind.get_LiteralInt(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralInt'])
end


regKind( [[LiteralInt]] )
function Filter:processLiteralInt( node, ... )

end


function NodeManager:getLiteralIntNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralInt']) )
end


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
function LiteralIntNode:canBeStatement(  )

   return false
end
function LiteralIntNode.new( pos, typeList, token, num )
   local obj = {}
   LiteralIntNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token, num ); end
   return obj
end
function LiteralIntNode:__init(pos, typeList, token, num) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralInt']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralIntNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralIntNode.new(pos, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralIntNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralIntNode  } )
end
function LiteralIntNode:get_token()       
   return self.token         
end
function LiteralIntNode:get_num()       
   return self.num         
end


function NodeKind.get_LiteralReal(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralReal'])
end


regKind( [[LiteralReal]] )
function Filter:processLiteralReal( node, ... )

end


function NodeManager:getLiteralRealNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralReal']) )
end


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
function LiteralRealNode:canBeStatement(  )

   return false
end
function LiteralRealNode.new( pos, typeList, token, num )
   local obj = {}
   LiteralRealNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token, num ); end
   return obj
end
function LiteralRealNode:__init(pos, typeList, token, num) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralReal']), pos, typeList)
   
   
   self.token = token
   self.num = num
   
end
function LiteralRealNode.create( nodeMan, pos, typeList, token, num )

   local node = LiteralRealNode.new(pos, typeList, token, num)
   nodeMan:addNode( node )
   return node
end
function LiteralRealNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralRealNode  } )
end
function LiteralRealNode:get_token()       
   return self.token         
end
function LiteralRealNode:get_num()       
   return self.num         
end


function NodeKind.get_LiteralArray(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralArray'])
end


regKind( [[LiteralArray]] )
function Filter:processLiteralArray( node, ... )

end


function NodeManager:getLiteralArrayNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralArray']) )
end


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
function LiteralArrayNode:canBeStatement(  )

   return false
end
function LiteralArrayNode.new( pos, typeList, expList )
   local obj = {}
   LiteralArrayNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, expList ); end
   return obj
end
function LiteralArrayNode:__init(pos, typeList, expList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralArray']), pos, typeList)
   
   
   self.expList = expList
   
end
function LiteralArrayNode.create( nodeMan, pos, typeList, expList )

   local node = LiteralArrayNode.new(pos, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function LiteralArrayNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralArrayNode  } )
end
function LiteralArrayNode:get_expList()       
   return self.expList         
end


function NodeKind.get_LiteralList(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralList'])
end


regKind( [[LiteralList]] )
function Filter:processLiteralList( node, ... )

end


function NodeManager:getLiteralListNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralList']) )
end


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
function LiteralListNode:canBeStatement(  )

   return false
end
function LiteralListNode.new( pos, typeList, expList )
   local obj = {}
   LiteralListNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, expList ); end
   return obj
end
function LiteralListNode:__init(pos, typeList, expList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralList']), pos, typeList)
   
   
   self.expList = expList
   
end
function LiteralListNode.create( nodeMan, pos, typeList, expList )

   local node = LiteralListNode.new(pos, typeList, expList)
   nodeMan:addNode( node )
   return node
end
function LiteralListNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralListNode  } )
end
function LiteralListNode:get_expList()       
   return self.expList         
end


local PairItem = {}
_moduleObj.PairItem = PairItem
function PairItem.setmeta( obj )
  setmetatable( obj, { __index = PairItem  } )
end
function PairItem.new( key, val )
   local obj = {}
   PairItem.setmeta( obj )
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

function NodeKind.get_LiteralMap(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralMap'])
end


regKind( [[LiteralMap]] )
function Filter:processLiteralMap( node, ... )

end


function NodeManager:getLiteralMapNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralMap']) )
end


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
function LiteralMapNode:canBeStatement(  )

   return false
end
function LiteralMapNode.new( pos, typeList, map, pairList )
   local obj = {}
   LiteralMapNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, map, pairList ); end
   return obj
end
function LiteralMapNode:__init(pos, typeList, map, pairList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralMap']), pos, typeList)
   
   
   self.map = map
   self.pairList = pairList
   
end
function LiteralMapNode.create( nodeMan, pos, typeList, map, pairList )

   local node = LiteralMapNode.new(pos, typeList, map, pairList)
   nodeMan:addNode( node )
   return node
end
function LiteralMapNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralMapNode  } )
end
function LiteralMapNode:get_map()       
   return self.map         
end
function LiteralMapNode:get_pairList()       
   return self.pairList         
end


function NodeKind.get_LiteralString(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralString'])
end


regKind( [[LiteralString]] )
function Filter:processLiteralString( node, ... )

end


function NodeManager:getLiteralStringNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralString']) )
end


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
function LiteralStringNode:canBeStatement(  )

   return false
end
function LiteralStringNode.new( pos, typeList, token, argList )
   local obj = {}
   LiteralStringNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token, argList ); end
   return obj
end
function LiteralStringNode:__init(pos, typeList, token, argList) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralString']), pos, typeList)
   
   
   self.token = token
   self.argList = argList
   
end
function LiteralStringNode.create( nodeMan, pos, typeList, token, argList )

   local node = LiteralStringNode.new(pos, typeList, token, argList)
   nodeMan:addNode( node )
   return node
end
function LiteralStringNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralStringNode  } )
end
function LiteralStringNode:get_token()       
   return self.token         
end
function LiteralStringNode:get_argList()       
   return self.argList         
end


function NodeKind.get_LiteralBool(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralBool'])
end


regKind( [[LiteralBool]] )
function Filter:processLiteralBool( node, ... )

end


function NodeManager:getLiteralBoolNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralBool']) )
end


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
function LiteralBoolNode:canBeStatement(  )

   return false
end
function LiteralBoolNode.new( pos, typeList, token )
   local obj = {}
   LiteralBoolNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token ); end
   return obj
end
function LiteralBoolNode:__init(pos, typeList, token) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralBool']), pos, typeList)
   
   
   self.token = token
   
end
function LiteralBoolNode.create( nodeMan, pos, typeList, token )

   local node = LiteralBoolNode.new(pos, typeList, token)
   nodeMan:addNode( node )
   return node
end
function LiteralBoolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralBoolNode  } )
end
function LiteralBoolNode:get_token()       
   return self.token         
end


function NodeKind.get_LiteralSymbol(  )

   return _lune.unwrap( _moduleObj.nodeKind['LiteralSymbol'])
end


regKind( [[LiteralSymbol]] )
function Filter:processLiteralSymbol( node, ... )

end


function NodeManager:getLiteralSymbolNodeList(  )

   return self:getList( _lune.unwrap( _moduleObj.nodeKind['LiteralSymbol']) )
end


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
function LiteralSymbolNode:canBeStatement(  )

   return false
end
function LiteralSymbolNode.new( pos, typeList, token )
   local obj = {}
   LiteralSymbolNode.setmeta( obj )
   if obj.__init then obj:__init( pos, typeList, token ); end
   return obj
end
function LiteralSymbolNode:__init(pos, typeList, token) 
   Node.__init( self ,_lune.unwrap( _moduleObj.nodeKind['LiteralSymbol']), pos, typeList)
   
   
   self.token = token
   
end
function LiteralSymbolNode.create( nodeMan, pos, typeList, token )

   local node = LiteralSymbolNode.new(pos, typeList, token)
   nodeMan:addNode( node )
   return node
end
function LiteralSymbolNode.setmeta( obj )
  setmetatable( obj, { __index = LiteralSymbolNode  } )
end
function LiteralSymbolNode:get_token()       
   return self.token         
end


function Node:getSymbolInfo(  )

   local function processExpNode( node )
   
      do
         local _switchExp = (node:get_kind() )
         if _switchExp == NodeKind.get_ExpRef() then
            return {(node ):get_symbolInfo()}
         elseif _switchExp == NodeKind.get_RefField() then
            local refFieldNode = node
            do
               local _exp = refFieldNode:get_symbolInfo()
               if _exp ~= nil then
                  return {_exp}
               end
            end
            
            return {}
         elseif _switchExp == NodeKind.get_GetField() then
            local getFieldNode = node
            do
               local _exp = getFieldNode:get_symbolInfo()
               if _exp ~= nil then
                  return {_exp}
               end
            end
            
            return {}
         elseif _switchExp == NodeKind.get_ExpList() then
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

function WhileNode:getBreakKind( checkMode )

   if checkMode ~= CheckBreakMode.Normal and checkMode ~= CheckBreakMode.Return then
      local kind = BreakKind.None
      for __index, stmt in pairs( self.block:get_stmtList() ) do
         local work = stmt:getBreakKind( checkMode )
         if checkMode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if checkMode == CheckBreakMode.Normal or checkMode == CheckBreakMode.Return then
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      if kind == BreakKind.Break then
         return BreakKind.None
      end
      
      return kind
   else
    
      if self.exp:get_expType():get_nilable() then
         return BreakKind.None
      end
      
      if self.exp:get_expType():equals( _moduleObj.builtinTypeBool ) then
         if self.exp:get_kind() == NodeKind.get_LiteralBool() then
            local boolNode = self.exp
            if boolNode:get_token().txt == "false" then
               return BreakKind.None
            end
            
         else
          
            return BreakKind.None
         end
         
      end
      
      local mode = CheckBreakMode.IgnoreFlow
      local kind = BreakKind.None
      for __index, stmt in pairs( self.block:get_stmtList() ) do
         local work = stmt:getBreakKind( mode )
         if mode == CheckBreakMode.IgnoreFlowReturn then
            if work == BreakKind.Return then
               return BreakKind.Return
            end
            
            if work == BreakKind.NeverRet then
               return BreakKind.NeverRet
            end
            
         else
          
            do
               local _switchExp = work
               if _switchExp == BreakKind.None then
                  if mode == CheckBreakMode.Normal or mode == CheckBreakMode.Return then
                     
                  end
                  
               else 
                  
                     if kind == BreakKind.None or kind > work then
                        kind = work
                     end
                     
               end
            end
            
         end
         
         
      end
      
      if kind == BreakKind.Break then
         return BreakKind.None
      end
      
      if kind == BreakKind.Return then
         return BreakKind.Return
      end
      
      return BreakKind.NeverRet
   end
   
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
            if  nil == txt then
               local _txt = txt
            
               return {}, {}
            end
            
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
            if  nil == item then
               local _item = item
            
               return {}, {}
            end
            
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
         if #arg > 1 then
            local argTxt = arg[1]
            if  nil == argTxt then
               local _argTxt = argTxt
            
               return {}, {}
            end
            
            table.insert( argTbl, argTxt )
         end
         
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

   local prefix = (_lune.unwrap( self.prefix:getLiteral(  )[1]) )
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
      txt = string.format( "%s%s", txt, tostring( token:getLiteral(  )[1]))
   end
   
   return {txt}, {self:get_expType(  )}
end

function ExpRefNode:getLiteral(  )

   local typeInfo = self.symbolInfo:get_typeInfo()
   if typeInfo:get_kind() ~= TypeInfoKind.Enum then
      return {}, {}
   end
   
   local enumTypeInfo = typeInfo
   local val = _lune.unwrap( enumTypeInfo:getEnumValInfo( self.symbolInfo:get_name() ))
   return {val:get_val()}, {enumTypeInfo:get_valTypeInfo()}
end

function ExpOp2Node:getLiteral(  )

   local val1List, type1List = self:get_exp1():getLiteral(  )
   local val2List, type2List = self:get_exp2():getLiteral(  )
   if #val1List ~= 1 or #type1List ~= 1 or #val2List ~= 1 or #type2List ~= 1 then
      return {}, {}
   end
   
   local val1, type1, val2, type2 = _lune.unwrap( val1List[1]), type1List[1]:get_srcTypeInfo(), _lune.unwrap( val2List[1]), type2List[1]:get_srcTypeInfo()
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

local DefMacroInfo = {}
setmetatable( DefMacroInfo, { __index = MacroInfo } )
_moduleObj.DefMacroInfo = DefMacroInfo
function DefMacroInfo:get_name(  )

   return self.declInfo:get_name().txt
end
function DefMacroInfo:getArgList(  )

   return self.argList
end
function DefMacroInfo:getTokenList(  )

   return self.declInfo:get_tokenList()
end
function DefMacroInfo.new( func, declInfo, symbol2MacroValInfoMap )
   local obj = {}
   DefMacroInfo.setmeta( obj )
   if obj.__init then obj:__init( func, declInfo, symbol2MacroValInfoMap ); end
   return obj
end
function DefMacroInfo:__init(func, declInfo, symbol2MacroValInfoMap) 
   MacroInfo.__init( self ,func, symbol2MacroValInfoMap)
   
   self.declInfo = declInfo
   self.argList = {}
   for __index, argNode in pairs( declInfo:get_argList() ) do
      if argNode:get_kind(  ) == NodeKind.get_DeclArg() then
         local argType = argNode:get_argType():get_expType()
         local argName = argNode:get_name().txt
         table.insert( self.argList, MacroArgInfo.new(argName, argType) )
      end
      
   end
   
end
function DefMacroInfo.setmeta( obj )
  setmetatable( obj, { __index = DefMacroInfo  } )
end

local processInfoQueue = {}
local function pushProcessInfo( processInfo )

   if #processInfoQueue == 0 then
      if idProv:get_id() >= userStartId then
         Util.err( "builtinId is over" )
      end
      
   end
   
   table.insert( processInfoQueue, ProcessInfo.new(idProv, typeInfo2ModifierMap, typeInfo2DDDMap) )
   if processInfo ~= nil then
      idProv = processInfo:get_idProvier()
      typeInfo2ModifierMap = processInfo:get_typeInfo2ModifierMap()
      typeInfo2DDDMap = processInfo:get_typeInfo2DDDMap()
   else
      idProv = IdProvider.new(userStartId)
      typeInfo2ModifierMap = {}
      typeInfo2DDDMap = {}
   end
   
   return ProcessInfo.new(idProv, typeInfo2ModifierMap, typeInfo2DDDMap)
end
_moduleObj.pushProcessInfo = pushProcessInfo
local function popProcessInfo(  )

   local info = processInfoQueue[#processInfoQueue]
   idProv = info:get_idProvier()
   typeInfo2ModifierMap = info:get_typeInfo2ModifierMap()
   table.remove( processInfoQueue )
end
_moduleObj.popProcessInfo = popProcessInfo
local BitOpKind = {}
_moduleObj.BitOpKind = BitOpKind
BitOpKind._val2NameMap = {}
function BitOpKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "BitOpKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function BitOpKind._from( val )
   if BitOpKind._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
BitOpKind.__allList = {}
function BitOpKind.get__allList()
   return BitOpKind.__allList
end

BitOpKind.And = 0
BitOpKind._val2NameMap[0] = 'And'
BitOpKind.__allList[1] = BitOpKind.And
BitOpKind.Or = 1
BitOpKind._val2NameMap[1] = 'Or'
BitOpKind.__allList[2] = BitOpKind.Or
BitOpKind.Xor = 2
BitOpKind._val2NameMap[2] = 'Xor'
BitOpKind.__allList[3] = BitOpKind.Xor
BitOpKind.LShift = 3
BitOpKind._val2NameMap[3] = 'LShift'
BitOpKind.__allList[4] = BitOpKind.LShift
BitOpKind.RShift = 4
BitOpKind._val2NameMap[4] = 'RShift'
BitOpKind.__allList[5] = BitOpKind.RShift

local bitBinOpMap = {["&"] = BitOpKind.And, ["|"] = BitOpKind.Or, ["~"] = BitOpKind.Xor, ["|>>"] = BitOpKind.RShift, ["|<<"] = BitOpKind.LShift}
_moduleObj.bitBinOpMap = bitBinOpMap

return _moduleObj
