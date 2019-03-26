--lune/base/Ast.lns
local _moduleObj = {}
local __mod__ = 'lune.base.Ast'
if not _lune then
   _lune = {}
end
function _lune.newAlge( kind, vals )
   local memInfoList = kind[ 2 ]
   if not memInfoList then
      return kind
   end
   return { kind[ 1 ], vals }
end

function _lune._fromList( obj, list, memInfoList )
   if type( list ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( list[ index ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s[%s]", memInfo.name, key) or memInfo.name
      end
      obj[ index ] = val
   end
   return true
end
function _lune._AlgeFrom( Alge, val )
   local work = Alge._name2Val[ val[ 1 ] ]
   if not work then
      return nil
   end
   if #work == 1 then
     return work
   end
   local paramList = {}
   local result, mess = _lune._fromList( paramList, val[ 2 ], work[ 2 ] )
   if not result then
      return nil, mess
   end
   return { work[ 1 ], paramList }
end

function _lune._Set_or( setObj, otherSet )
   for val in pairs( otherSet ) do
      setObj[ val ] = true
   end
   return setObj
end
function _lune._Set_and( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if not otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_has( setObj, val )
   return setObj[ val ] ~= nil
end
function _lune._Set_sub( setObj, otherSet )
   local delValList = {}
   for val in pairs( setObj ) do
      if otherSet[ val ] then
         table.insert( delValList, val )
      end
   end
   for index, val in ipairs( delValList ) do
      setObj[ val ] = nil
   end
   return setObj
end
function _lune._Set_len( setObj )
   local total = 0
   for val in pairs( setObj ) do
      total = total + 1
   end
   return total
end
function _lune._Set_clone( setObj )
   local obj = {}
   for val in pairs( setObj ) do
      obj[ val ] = true
   end
   return obj
end

function _lune._toSet( val, toKeyInfo )
   if type( val ) == "table" then
      local tbl = {}
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal = _lune._toBool( mem )
         if mapKey == nil or mapVal == nil then
            if mapKey == nil then
               return nil
            end
            if keySub == nil then
               return nil, mapKey
            end
            return nil, string.format( "%s.%s", mapKey, keySub)
         end
         tbl[ mapKey ] = mapVal
      end
      return tbl
   end
   return nil
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

function _lune.__isInstanceOf( obj, class )
   while obj do
      local meta = getmetatable( obj )
      if not meta then
	 return false
      end
      local indexTbl = meta.__index
      if indexTbl == class then
	 return true
      end
      if meta.ifList then
         for index, ifType in ipairs( meta.ifList ) do
            if _lune.__isInstanceOf( ifType, class ) then
               return true
            end
         end
      end
      obj = indexTbl
   end
   return false
end

function _lune.__Cast( obj, kind, class )
   if kind == 0 then -- int
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) ~= obj then
         return nil
      end
      return obj
   elseif kind == 1 then -- real
      if type( obj ) ~= "number" then
         return nil
      end
      if math.floor( obj ) == obj then
         return nil
      end
      return obj
   elseif kind == 2 then -- str
      if type( obj ) ~= "string" then
         return nil
      end
      return obj
   elseif kind == 3 then -- class
      return _lune.__isInstanceOf( obj, class ) and obj or nil
   end
   return nil
end

local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Code = _lune.loadModule( 'lune.base.Code' )
local IdProvider = {}
_moduleObj.IdProvider = IdProvider
function IdProvider:increment(  )

   self.id = self.id + 1
   if self.id >= self.maxId then
      Util.err( "id is over" )
   end
   
end
function IdProvider:getNewId(  )

   local newId = self.id
   self.id = self.id + 1
   if self.id >= self.maxId then
      Util.err( "id is over" )
   end
   
   return newId
end
function IdProvider.setmeta( obj )
  setmetatable( obj, { __index = IdProvider  } )
end
function IdProvider.new( id, maxId )
   local obj = {}
   IdProvider.setmeta( obj )
   if obj.__init then
      obj:__init( id, maxId )
   end        
   return obj 
end         
function IdProvider:__init( id, maxId ) 

   self.id = id
   self.maxId = maxId
end
function IdProvider:get_id()       
   return self.id         
end

local extStartId = 100000
local extMaxId = 10000000
local idProvBase = IdProvider.new(1, extStartId)
local idProvExt = IdProvider.new(extStartId, extMaxId)
local idProv = idProvBase
local userStartId = 1000
local rootTypeId = idProv:getNewId(  )
_moduleObj.rootTypeId = rootTypeId



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
SerializeKind.Alias = 7
SerializeKind._val2NameMap[7] = 'Alias'
SerializeKind.__allList[8] = SerializeKind.Alias
SerializeKind.Alternate = 8
SerializeKind._val2NameMap[8] = 'Alternate'
SerializeKind.__allList[9] = SerializeKind.Alternate
SerializeKind.Generic = 9
SerializeKind._val2NameMap[9] = 'Generic'
SerializeKind.__allList[10] = SerializeKind.Generic
SerializeKind.Box = 10
SerializeKind._val2NameMap[10] = 'Box'
SerializeKind.__allList[11] = SerializeKind.Box

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
TypeInfoKind.Set = 17
TypeInfoKind._val2NameMap[17] = 'Set'
TypeInfoKind.__allList[18] = TypeInfoKind.Set
TypeInfoKind.Alternate = 18
TypeInfoKind._val2NameMap[18] = 'Alternate'
TypeInfoKind.__allList[19] = TypeInfoKind.Alternate
TypeInfoKind.Box = 19
TypeInfoKind._val2NameMap[19] = 'Box'
TypeInfoKind.__allList[20] = TypeInfoKind.Box
TypeInfoKind.CanEvalCtrl = 20
TypeInfoKind._val2NameMap[20] = 'CanEvalCtrl'
TypeInfoKind.__allList[21] = TypeInfoKind.CanEvalCtrl
TypeInfoKind.Etc = 21
TypeInfoKind._val2NameMap[21] = 'Etc'
TypeInfoKind.__allList[22] = TypeInfoKind.Etc

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
local MutMode = {}
_moduleObj.MutMode = MutMode
MutMode._val2NameMap = {}
function MutMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MutMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function MutMode._from( val )
   if MutMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
MutMode.__allList = {}
function MutMode.get__allList()
   return MutMode.__allList
end

MutMode.IMut = 0
MutMode._val2NameMap[0] = 'IMut'
MutMode.__allList[1] = MutMode.IMut
MutMode.Mut = 1
MutMode._val2NameMap[1] = 'Mut'
MutMode.__allList[2] = MutMode.Mut
MutMode.AllMut = 2
MutMode._val2NameMap[2] = 'AllMut'
MutMode.__allList[3] = MutMode.AllMut

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
function NormalSymbolInfo:get_mutable(  )

   return self.mutMode ~= MutMode.IMut
end
function NormalSymbolInfo:getOrg(  )

   return self
end
function NormalSymbolInfo.new( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutMode, hasValueFlag )
   local obj = {}
   NormalSymbolInfo.setmeta( obj )
   if obj.__init then obj:__init( kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutMode, hasValueFlag ); end
   return obj
end
function NormalSymbolInfo:__init(kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, typeInfo, mutMode, hasValueFlag) 
   SymbolInfo.__init( self)
   
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
   self.mutMode = _lune.unwrapDefault( mutMode, MutMode.IMut)
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
function NormalSymbolInfo:get_kind()       
   return self.kind         
end
function NormalSymbolInfo:get_hasValueFlag()       
   return self.hasValueFlag         
end
function NormalSymbolInfo:set_hasValueFlag( hasValueFlag )   
   self.hasValueFlag = hasValueFlag              
end
function NormalSymbolInfo:get_mutMode()       
   return self.mutMode         
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
function Scope:addSymbol( symbolInfo )

   self.symbol2SymbolInfoMap[symbolInfo:get_name()] = symbolInfo
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
function TypeInfo.new( scope )
   local obj = {}
   TypeInfo.setmeta( obj )
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
   
   self.typeData = TypeData.new({})
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
function TypeInfo:isInheritFrom( other, alt2type )

   return false
end
function TypeInfo:get_rawTxt(  )

   return ""
end
function TypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   return ""
end
function TypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function TypeInfo:canEvalWith( other, opTxt, alt2type )

   return false
end
function TypeInfo:get_abstractFlag(  )

   return false
end
function TypeInfo:serialize( stream, validChildrenSet )

   return 
end
function TypeInfo:get_display_stirng_with( raw )

   return ""
end
function TypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( "" )
end
function TypeInfo:get_srcTypeInfo(  )

   return self
end
function TypeInfo:equals( typeInfo, alt2type )

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
function TypeInfo:get_children(  )

   return self.typeData:get_children()
end
function TypeInfo:addChildren( child )

   self.typeData:addChildren( child )
end
function TypeInfo:get_mutMode(  )

   return MutMode.Mut
end
function TypeInfo.isMut( typeInfo )

   return typeInfo:get_mutMode() ~= MutMode.IMut
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
function TypeInfo:applyGeneric( alt2typeMap )

   return self
end
function TypeInfo:get_genSrcTypeInfo(  )

   return self
end
function TypeInfo:serializeTypeInfoList( name, list, onlyPub )

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
function TypeInfo.createScope( parent, classFlag, baseInfo, interfaceList )

   local inheritScope = nil
   if baseInfo ~= nil then
      inheritScope = _lune.unwrap( baseInfo.scope)
   end
   
   local ifScopeList = {}
   if interfaceList ~= nil then
      for __index, ifType in pairs( interfaceList ) do
         table.insert( ifScopeList, _lune.unwrap( ifType.scope) )
      end
      
   end
   
   return Scope.new(parent, classFlag, inheritScope, ifScopeList)
end
function TypeInfo.setmeta( obj )
  setmetatable( obj, { __index = TypeInfo  } )
end
function TypeInfo:get_scope()       
   return self.scope         
end
function TypeInfo:get_typeData()       
   return self.typeData         
end

local function getScope( typeInfo )

   return typeInfo:get_scope()
end
_moduleObj.getScope = getScope
local function isExtId( typeInfo )

   if typeInfo:get_typeId() >= extStartId then
      return true
   end
   
   return false
end
_moduleObj.isExtId = isExtId
local headTypeInfo = TypeInfo.new(_moduleObj.rootScope)
_moduleObj.headTypeInfo = headTypeInfo

function TypeInfo.isInherit( typeInfo, other, alt2type )

   local baseTypeInfo = typeInfo:get_baseTypeInfo()
   local interfaceList = typeInfo:get_interfaceList()
   local otherTypeId = other:get_typeId()
   if typeInfo:get_typeId() == otherTypeId then
      return true
   end
   
   if baseTypeInfo ~= _moduleObj.headTypeInfo then
      if baseTypeInfo:isInheritFrom( other, alt2type ) then
         return true
      end
      
   end
   
   
   for __index, ifType in pairs( typeInfo:get_interfaceList() ) do
      if ifType:isInheritFrom( other, alt2type ) then
         return true
      end
      
   end
   
   return false
end

local AutoBoxingInfo = {}
setmetatable( AutoBoxingInfo, { __index = TypeInfo } )
_moduleObj.AutoBoxingInfo = AutoBoxingInfo
function AutoBoxingInfo.new(  )
   local obj = {}
   AutoBoxingInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function AutoBoxingInfo:__init() 
   TypeInfo.__init( self,nil)
   
   self.count = 0
   AutoBoxingInfo.allObj[self] = self
end
function AutoBoxingInfo:get_kind(  )

   return TypeInfoKind.Etc
end
function AutoBoxingInfo:inc(  )

   local obj = _lune.unwrap( AutoBoxingInfo.allObj[self])
   obj.count = obj.count + 1
end
function AutoBoxingInfo:unregist(  )

   AutoBoxingInfo.allObj[self] = nil
end
function AutoBoxingInfo.setmeta( obj )
  setmetatable( obj, { __index = AutoBoxingInfo  } )
end
function AutoBoxingInfo:get_count()       
   return self.count         
end
do
   AutoBoxingInfo.allObj = {}
end

local CanEvalCtrlTypeInfo = {}
setmetatable( CanEvalCtrlTypeInfo, { __index = TypeInfo } )
_moduleObj.CanEvalCtrlTypeInfo = CanEvalCtrlTypeInfo
function CanEvalCtrlTypeInfo:get_kind(  )

   return TypeInfoKind.CanEvalCtrl
end
function CanEvalCtrlTypeInfo:get_typeId(  )

   return -1
end
function CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( detectFlag )

   if detectFlag then
      return {[CanEvalCtrlTypeInfo.detectAlt] = _moduleObj.headTypeInfo}
   end
   
   return {}
end
function CanEvalCtrlTypeInfo.isValidApply( alt2type )

   return alt2type[CanEvalCtrlTypeInfo.detectAlt] ~= nil
end
function CanEvalCtrlTypeInfo.setupNeedAutoBoxing( alt2type )

   alt2type[CanEvalCtrlTypeInfo.needAutoBoxing] = AutoBoxingInfo.new()
end
function CanEvalCtrlTypeInfo.updateNeedAutoBoxing( alt2type )

   do
      local _exp = alt2type[CanEvalCtrlTypeInfo.needAutoBoxing]
      if _exp ~= nil then
         do
            local autoBoxingInfo = _lune.__Cast( _exp, 3, AutoBoxingInfo )
            if autoBoxingInfo ~= nil then
               autoBoxingInfo:inc(  )
            end
         end
         
      else
         Util.err( "no exist needAutoBoxing" )
      end
   end
   
end
function CanEvalCtrlTypeInfo.finishNeedAutoBoxing( alt2type, count )

   do
      local _exp = alt2type[CanEvalCtrlTypeInfo.needAutoBoxing]
      if _exp ~= nil then
         do
            local autoBoxingInfo = _lune.__Cast( _exp, 3, AutoBoxingInfo )
            if autoBoxingInfo ~= nil then
               autoBoxingInfo:unregist(  )
               return autoBoxingInfo:get_count() == count
            end
         end
         
      end
   end
   
   return false
end
function CanEvalCtrlTypeInfo.canAutoBoxing( dst, src )

   local dstSrc = dst:get_srcTypeInfo():get_nonnilableType()
   if dstSrc:get_kind() ~= TypeInfoKind.Box then
      return false
   end
   
   local srcSrc = src:get_srcTypeInfo():get_nonnilableType()
   if srcSrc:get_kind() == TypeInfoKind.Box then
      return false
   end
   
   return true
end
function CanEvalCtrlTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = CanEvalCtrlTypeInfo  } )
end
function CanEvalCtrlTypeInfo.new(  )
   local obj = {}
   CanEvalCtrlTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end        
   return obj 
end         
function CanEvalCtrlTypeInfo:__init(  ) 

   TypeInfo.__init( self )
end
do
   CanEvalCtrlTypeInfo.detectAlt = CanEvalCtrlTypeInfo.new()
   CanEvalCtrlTypeInfo.needAutoBoxing = CanEvalCtrlTypeInfo.new()
   CanEvalCtrlTypeInfo.checkTypeTarget = CanEvalCtrlTypeInfo.new()
end

local AliasTypeInfo = {}
setmetatable( AliasTypeInfo, { __index = TypeInfo } )
_moduleObj.AliasTypeInfo = AliasTypeInfo
function AliasTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self.rawTxt, fullName, importInfo, localFlag )
end
function AliasTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, rawTxt = %q, srcTypeId = %d }\n', SerializeKind.Alias, parentId, self.typeId, self.rawTxt, self.aliasSrcTypeInfo:get_typeId()) )
end
function AliasTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self.rawTxt )
end
function AliasTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function AliasTypeInfo:applyGeneric( alt2typeMap )

   local typeInfo = self.aliasSrcTypeInfo:applyGeneric( alt2typeMap )
   if typeInfo == self.aliasSrcTypeInfo then
      return self
   end
   
   return nil
end
function AliasTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AliasTypeInfo  } )
end
function AliasTypeInfo.new( rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag, typeId )
   local obj = {}
   AliasTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag, typeId )
   end        
   return obj 
end         
function AliasTypeInfo:__init( rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag, typeId ) 

   TypeInfo.__init( self )
   self.rawTxt = rawTxt
   self.accessMode = accessMode
   self.parentInfo = parentInfo
   self.aliasSrcTypeInfo = aliasSrcTypeInfo
   self.externalFlag = externalFlag
   self.typeId = typeId
end
function AliasTypeInfo:get_rawTxt()       
   return self.rawTxt         
end
function AliasTypeInfo:get_accessMode()       
   return self.accessMode         
end
function AliasTypeInfo:get_parentInfo()       
   return self.parentInfo         
end
function AliasTypeInfo:get_aliasSrcTypeInfo()       
   return self.aliasSrcTypeInfo         
end
function AliasTypeInfo:get_externalFlag()       
   return self.externalFlag         
end
function AliasTypeInfo:get_typeId()       
   return self.typeId         
end
function AliasTypeInfo:isModule( ... )
   return self.aliasSrcTypeInfo:isModule( ... )
end       

function AliasTypeInfo:get_baseId( ... )
   return self.aliasSrcTypeInfo:get_baseId( ... )
end       

function AliasTypeInfo:isInheritFrom( ... )
   return self.aliasSrcTypeInfo:isInheritFrom( ... )
end       

function AliasTypeInfo:getTxtWithRaw( ... )
   return self.aliasSrcTypeInfo:getTxtWithRaw( ... )
end       

function AliasTypeInfo:canEvalWith( ... )
   return self.aliasSrcTypeInfo:canEvalWith( ... )
end       

function AliasTypeInfo:get_abstractFlag( ... )
   return self.aliasSrcTypeInfo:get_abstractFlag( ... )
end       

function AliasTypeInfo:get_display_stirng_with( ... )
   return self.aliasSrcTypeInfo:get_display_stirng_with( ... )
end       

function AliasTypeInfo:get_srcTypeInfo( ... )
   return self.aliasSrcTypeInfo:get_srcTypeInfo( ... )
end       

function AliasTypeInfo:equals( ... )
   return self.aliasSrcTypeInfo:equals( ... )
end       

function AliasTypeInfo:get_interfaceList( ... )
   return self.aliasSrcTypeInfo:get_interfaceList( ... )
end       

function AliasTypeInfo:get_itemTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_itemTypeInfoList( ... )
end       

function AliasTypeInfo:get_argTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_argTypeInfoList( ... )
end       

function AliasTypeInfo:get_retTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_retTypeInfoList( ... )
end       

function AliasTypeInfo:hasRouteNamespaceFrom( ... )
   return self.aliasSrcTypeInfo:hasRouteNamespaceFrom( ... )
end       

function AliasTypeInfo:getModule( ... )
   return self.aliasSrcTypeInfo:getModule( ... )
end       

function AliasTypeInfo:get_kind( ... )
   return self.aliasSrcTypeInfo:get_kind( ... )
end       

function AliasTypeInfo:get_staticFlag( ... )
   return self.aliasSrcTypeInfo:get_staticFlag( ... )
end       

function AliasTypeInfo:get_autoFlag( ... )
   return self.aliasSrcTypeInfo:get_autoFlag( ... )
end       

function AliasTypeInfo:get_nonnilableType( ... )
   return self.aliasSrcTypeInfo:get_nonnilableType( ... )
end       

function AliasTypeInfo:get_baseTypeInfo( ... )
   return self.aliasSrcTypeInfo:get_baseTypeInfo( ... )
end       

function AliasTypeInfo:get_nilable( ... )
   return self.aliasSrcTypeInfo:get_nilable( ... )
end       

function AliasTypeInfo:get_nilableTypeInfo( ... )
   return self.aliasSrcTypeInfo:get_nilableTypeInfo( ... )
end       

function AliasTypeInfo:get_children( ... )
   return self.aliasSrcTypeInfo:get_children( ... )
end       

function AliasTypeInfo:addChildren( ... )
   return self.aliasSrcTypeInfo:addChildren( ... )
end       

function AliasTypeInfo:get_mutMode( ... )
   return self.aliasSrcTypeInfo:get_mutMode( ... )
end       

function AliasTypeInfo:getParentFullName( ... )
   return self.aliasSrcTypeInfo:getParentFullName( ... )
end       

function AliasTypeInfo:get_genSrcTypeInfo( ... )
   return self.aliasSrcTypeInfo:get_genSrcTypeInfo( ... )
end       

function AliasTypeInfo:serializeTypeInfoList( ... )
   return self.aliasSrcTypeInfo:serializeTypeInfoList( ... )
end       

function AliasTypeInfo:get_scope( ... )
   return self.aliasSrcTypeInfo:get_scope( ... )
end       

function AliasTypeInfo:get_typeData( ... )
   return self.aliasSrcTypeInfo:get_typeData( ... )
end       

function AliasTypeInfo:createAlt2typeMap( ... )
   return self.aliasSrcTypeInfo:createAlt2typeMap( ... )
end       

function AliasTypeInfo:getFullName( ... )
   return self.aliasSrcTypeInfo:getFullName( ... )
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
   local limitSymbol = false
   do
      local _exp = self.ownerTypeInfo
      if _exp ~= nil then
         if _exp:get_kind() == TypeInfoKind.Func or _exp:get_kind() == TypeInfoKind.Method or self == moduleScope or self == _moduleObj.rootScope then
            validThisScope = true
         elseif _exp:get_kind() == TypeInfoKind.IF or _exp:get_kind() == TypeInfoKind.Class or _exp:get_kind() == TypeInfoKind.Module then
            limitSymbol = true
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
         local symbolInfo = self.symbol2SymbolInfoMap[name]
         if symbolInfo ~= nil then
            if not limitSymbol or name == "self" or (symbolInfo:get_typeInfo():get_kind() == TypeInfoKind.Alternate and symbolInfo:get_kind() == SymbolKind.Typ ) then
               return symbolInfo:canAccess( fromScope )
            end
            
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

function Scope:add( kind, canBeLeft, canBeRight, name, typeInfo, accessMode, staticFlag, mutMode, hasValueFlag )

   local symbolInfo = NormalSymbolInfo.new(kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, typeInfo, mutMode, hasValueFlag)
   self.symbol2SymbolInfoMap[name] = symbolInfo
   return symbolInfo
end

function Scope:addLocalVar( argFlag, canBeLeft, name, typeInfo, mutable )

   self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, AccessMode.Local, false, mutable and MutMode.Mut or MutMode.IMut, true )
end

function Scope:addStaticVar( argFlag, canBeLeft, name, typeInfo, mutable )

   self:add( argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, typeInfo, AccessMode.Local, true, mutable and MutMode.Mut or MutMode.IMut, true )
end

function Scope:addVar( accessMode, name, typeInfo, mutable, hasValueFlag )

   self:add( SymbolKind.Var, true, true, name, typeInfo, accessMode, false, mutable and MutMode.Mut or MutMode.IMut, hasValueFlag )
end

function Scope:addEnumVal( name, typeInfo )

   self:add( SymbolKind.Mbr, false, true, name, typeInfo, AccessMode.Pub, true, MutMode.Mut, true )
end

function Scope:addEnum( accessMode, name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, MutMode.Mut, true )
end

function Scope:addAlgeVal( name, typeInfo )

   self:add( SymbolKind.Mbr, false, true, name, typeInfo, AccessMode.Pub, true, MutMode.Mut, true )
end

function Scope:addAlge( accessMode, name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, MutMode.Mut, true )
end

function Scope:addAlternate( accessMode, name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, accessMode, true, MutMode.Mut, true )
end

function Scope:addMember( name, typeInfo, accessMode, staticFlag, mutMode )

   return self:add( SymbolKind.Mbr, true, true, name, typeInfo, accessMode, staticFlag, mutMode, true )
end

function Scope:addMethod( typeInfo, accessMode, staticFlag, mutable )

   self:add( SymbolKind.Mtd, true, staticFlag, typeInfo:get_rawTxt(), typeInfo, accessMode, staticFlag, mutable and MutMode.Mut or MutMode.IMut, true )
end

function Scope:addFunc( typeInfo, accessMode, staticFlag, mutable )

   self:add( SymbolKind.Fun, true, true, typeInfo:get_rawTxt(), typeInfo, accessMode, staticFlag, mutable and MutMode.Mut or MutMode.IMut, true )
end

function Scope:addMacro( typeInfo, accessMode )

   self:add( SymbolKind.Mac, false, false, typeInfo:get_rawTxt(), typeInfo, accessMode, true, MutMode.IMut, true )
end

function Scope:addClass( name, typeInfo )

   self:add( SymbolKind.Typ, false, false, name, typeInfo, typeInfo:get_accessMode(), true, MutMode.Mut, true )
end

local function dumpScope( workscope, workprefix )

   local function dumpScopeSub( scope, prefix, readyIdSet )
   
      do
         local _exp = scope
         if _exp ~= nil then
            if _lune._Set_has(readyIdSet, _exp ) then
               return 
            end
            
            readyIdSet[_exp]= true
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
   
   dumpScopeSub( workscope, workprefix, {} )
end
_moduleObj.dumpScope = dumpScope
function TypeInfo:createAlt2typeMap( detectFlag )

   if not detectFlag then
      return {}
   end
   
   return {[CanEvalCtrlTypeInfo.detectAlt] = _moduleObj.headTypeInfo}
end

local NilTypeInfo = {}
setmetatable( NilTypeInfo, { __index = TypeInfo } )
_moduleObj.NilTypeInfo = NilTypeInfo
function NilTypeInfo.new(  )
   local obj = {}
   NilTypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function NilTypeInfo:__init() 
   TypeInfo.__init( self,nil)
   
   idProv:increment(  )
   self.typeId = idProv:get_id()
end
function NilTypeInfo:isModule(  )

   return false
end
function NilTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function NilTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   return "nil"
end
function NilTypeInfo:canEvalWith( other, opTxt, alt2type )

   return other:get_nilable()
end
function NilTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function NilTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( "nil" )
end
function NilTypeInfo:equals( typeInfo, alt2type )

   return self == typeInfo
end
function NilTypeInfo:get_parentInfo(  )

   return _moduleObj.headTypeInfo
end
function NilTypeInfo:hasRouteNamespaceFrom( other )

   return true
end
function NilTypeInfo:get_rawTxt(  )

   return "nil"
end
function NilTypeInfo:get_kind(  )

   return TypeInfoKind.Prim
end
function NilTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function NilTypeInfo:get_nilable(  )

   return true
end
function NilTypeInfo:get_mutMode(  )

   return MutMode.IMut
end
function NilTypeInfo:getParentFullName( importInfo, localFlag )

   return "nil"
end
function NilTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = NilTypeInfo  } )
end
function NilTypeInfo:get_typeId()       
   return self.typeId         
end

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
         if fromClass:isInheritFrom( nsClass, nil ) then
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
         do
            local _switchExp = self.symbolInfo:get_mutMode()
            if _switchExp == MutMode.AllMut then
               return true
            elseif _switchExp == MutMode.IMut then
               return false
            end
         end
         
         return _exp:get_mutMode() ~= MutMode.IMut
      end
   end
   
   return self.symbolInfo:get_mutable()
end
function AccessSymbolInfo:get_canBeLeft(  )

   if not self.overrideCanBeLeft then
      return false
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

function AccessSymbolInfo:get_mutMode( ... )
   return self.symbolInfo:get_mutMode( ... )
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

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function NilableTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   return self.nonnilableType:getTxtWithRaw( raw, fullName, importInfo, localFlag ) .. "!"
end
function NilableTypeInfo:get_display_stirng_with( raw )

   return self.nonnilableType:get_display_stirng_with( raw ) .. "!"
end
function NilableTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function NilableTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, nilable = true, orgTypeId = %d }\n', SerializeKind.Nilable, parentId, self.typeId, self.nonnilableType:get_typeId()) )
end
function NilableTypeInfo:equals( typeInfo, alt2type )

   if not typeInfo:get_nilable() then
      return false
   end
   
   return self.nonnilableType:equals( typeInfo:get_nonnilableType(), alt2type )
end
function NilableTypeInfo:applyGeneric( alt2typeMap )

   local typeInfo = self.nonnilableType:applyGeneric( alt2typeMap )
   if typeInfo == self.nonnilableType then
      return self
   end
   
   if typeInfo ~= nil then
      return typeInfo:get_nilableTypeInfo()
   end
   
   return nil
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

function NilableTypeInfo:get_rawTxt( ... )
   return self.nonnilableType:get_rawTxt( ... )
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

function NilableTypeInfo:get_children( ... )
   return self.nonnilableType:get_children( ... )
end       

function NilableTypeInfo:addChildren( ... )
   return self.nonnilableType:addChildren( ... )
end       

function NilableTypeInfo:get_mutMode( ... )
   return self.nonnilableType:get_mutMode( ... )
end       

function NilableTypeInfo:getParentFullName( ... )
   return self.nonnilableType:getParentFullName( ... )
end       

function NilableTypeInfo:get_genSrcTypeInfo( ... )
   return self.nonnilableType:get_genSrcTypeInfo( ... )
end       

function NilableTypeInfo:serializeTypeInfoList( ... )
   return self.nonnilableType:serializeTypeInfoList( ... )
end       

function NilableTypeInfo:get_scope( ... )
   return self.nonnilableType:get_scope( ... )
end       

function NilableTypeInfo:get_typeData( ... )
   return self.nonnilableType:get_typeData( ... )
end       

function NilableTypeInfo:createAlt2typeMap( ... )
   return self.nonnilableType:createAlt2typeMap( ... )
end       

function NilableTypeInfo:getFullName( ... )
   return self.nonnilableType:getFullName( ... )
end       


local AlternateTypeInfo = {}
setmetatable( AlternateTypeInfo, { __index = TypeInfo } )
_moduleObj.AlternateTypeInfo = AlternateTypeInfo
function AlternateTypeInfo.new( txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList )
   local obj = {}
   AlternateTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList ); end
   return obj
end
function AlternateTypeInfo:__init(txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList) 
   TypeInfo.__init( self,TypeInfo.createScope( nil, true, baseTypeInfo, interfaceList ))
   
   idProv:increment(  )
   self.typeId = idProv:get_id()
   self.txt = txt
   self.accessMode = accessMode
   self.moduleTypeInfo = moduleTypeInfo
   self.baseTypeInfo = _lune.unwrapDefault( baseTypeInfo, _moduleObj.headTypeInfo)
   self.interfaceList = _lune.unwrapDefault( interfaceList, {})
   idProv:increment(  )
   self.nilableTypeInfo = NilableTypeInfo.new(self, idProv:get_id())
end
function AlternateTypeInfo:isModule(  )

   return false
end
function AlternateTypeInfo:getParentId(  )

   return self.moduleTypeInfo:get_typeId()
end
function AlternateTypeInfo:get_baseId(  )

   return self.baseTypeInfo:get_typeId()
end
function AlternateTypeInfo:get_parentInfo(  )

   return self.moduleTypeInfo
end
function AlternateTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function AlternateTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   return self.txt
end
function AlternateTypeInfo.getAssign( typeInfo, alt2type )

   if typeInfo:get_kind() ~= TypeInfoKind.Alternate then
      return typeInfo
   end
   
   local otherWork = typeInfo
   while true do
      do
         local _exp = alt2type[otherWork]
         if _exp ~= nil then
            otherWork = _exp
         else
            return otherWork
         end
      end
      
   end
   
end
function AlternateTypeInfo:canSetFrom( other, opTxt, alt2type )

   local otherWork = AlternateTypeInfo.getAssign( other, alt2type )
   do
      local genType = alt2type[self]
      if genType ~= nil then
         if opTxt ~= nil then
            return genType:canEvalWith( otherWork, opTxt, alt2type )
         end
         
         return genType:equals( otherWork, alt2type )
      end
   end
   
   if not CanEvalCtrlTypeInfo.isValidApply( alt2type ) then
      return false
   end
   
   if self.baseTypeInfo ~= _moduleObj.headTypeInfo then
      if not other:isInheritFrom( self.baseTypeInfo, alt2type ) then
         return false
      end
      
   end
   
   
   for __index, ifType in pairs( self.interfaceList ) do
      if not other:isInheritFrom( ifType, alt2type ) then
         return false
      end
      
   end
   
   alt2type[self] = otherWork
   return true
end
function AlternateTypeInfo:isInheritFrom( other, alt2type )

   if alt2type ~= nil then
      local otherWork = AlternateTypeInfo.getAssign( other, alt2type )
      if self == otherWork:get_srcTypeInfo() then
         return true
      end
      
      do
         local genType = alt2type[self]
         if genType ~= nil then
            return genType:isInheritFrom( otherWork, alt2type )
         end
      end
      
      if not CanEvalCtrlTypeInfo.isValidApply( alt2type ) then
         return false
      end
      
   end
   
   if self == other:get_srcTypeInfo() then
      return true
   end
   
   local function check(  )
   
      if self.baseTypeInfo ~= _moduleObj.headTypeInfo then
         if self.baseTypeInfo:isInheritFrom( other, alt2type ) then
            return true
         end
         
      end
      
      
      for __index, ifType in pairs( self.interfaceList ) do
         if ifType:isInheritFrom( other, alt2type ) then
            return true
         end
         
      end
      
      return false
   end
   
   if check(  ) then
      if alt2type ~= nil then
         alt2type[self] = other
      end
      
      return true
   end
   
   return false
end
function AlternateTypeInfo:canEvalWith( other, opTxt, alt2type )

   if self == other:get_srcTypeInfo() then
      return true
   end
   
   if other:get_nilable() then
      return false
   end
   
   return self:canSetFrom( other, opTxt, alt2type )
end
function AlternateTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function AlternateTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self.txt )
end
function AlternateTypeInfo:equals( typeInfo, alt2type )

   if self == typeInfo then
      return true
   end
   
   if alt2type ~= nil then
      return self:canSetFrom( typeInfo, nil, alt2type )
   end
   
   return false
end
function AlternateTypeInfo:hasRouteNamespaceFrom( other )

   return true
end
function AlternateTypeInfo:get_rawTxt(  )

   return self.txt
end
function AlternateTypeInfo:get_kind(  )

   return TypeInfoKind.Alternate
end
function AlternateTypeInfo:get_nilable(  )

   return false
end
function AlternateTypeInfo:get_mutMode(  )

   return MutMode.IMut
end
function AlternateTypeInfo:getParentFullName( importInfo, localFlag )

   return ""
end
function AlternateTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, txt = %q, ', SerializeKind.Alternate, parentId, self.typeId, self.txt) .. string.format( 'accessMode = %d, baseId = %d, ', self.accessMode, self:get_baseId()) )
   stream:write( self:serializeTypeInfoList( "ifList = {", self.interfaceList ) )
   stream:write( "}\n" )
end
function AlternateTypeInfo:applyGeneric( alt2typeMap )

   return AlternateTypeInfo.getAssign( self, alt2typeMap )
end
function AlternateTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AlternateTypeInfo  } )
end
function AlternateTypeInfo:get_typeId()       
   return self.typeId         
end
function AlternateTypeInfo:get_txt()       
   return self.txt         
end
function AlternateTypeInfo:get_nilableTypeInfo()       
   return self.nilableTypeInfo         
end
function AlternateTypeInfo:get_accessMode()       
   return self.accessMode         
end
function AlternateTypeInfo:get_baseTypeInfo()       
   return self.baseTypeInfo         
end
function AlternateTypeInfo:get_interfaceList()       
   return self.interfaceList         
end

local boxRootAltType = AlternateTypeInfo.new("_T", AccessMode.Pub, _moduleObj.headTypeInfo)
local boxRootScope = Scope.new(_moduleObj.rootScope, true, nil)
local BoxTypeInfo = {}
setmetatable( BoxTypeInfo, { __index = TypeInfo } )
_moduleObj.BoxTypeInfo = BoxTypeInfo
function BoxTypeInfo.new( typeId, accessMode, boxingType )
   local obj = {}
   BoxTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( typeId, accessMode, boxingType ); end
   return obj
end
function BoxTypeInfo:__init(typeId, accessMode, boxingType) 
   TypeInfo.__init( self,boxRootScope)
   
   self.boxingType = boxingType
   self.typeId = typeId
   self.itemTypeInfoList = {boxingType}
   self.accessMode = accessMode
   idProv:increment(  )
   self.nilableTypeInfo = NilableTypeInfo.new(self, idProv:get_id())
end
function BoxTypeInfo:get_scope(  )

   return TypeInfo.get_scope( self  )
end
function BoxTypeInfo:get_kind(  )

   return TypeInfoKind.Box
end
function BoxTypeInfo:get_srcTypeInfo(  )

   return self
end
function BoxTypeInfo:get_nonnilableType(  )

   return self
end
function BoxTypeInfo:get_nilable(  )

   return false
end
function BoxTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function BoxTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   return "Nilable<" .. self.boxingType:getTxtWithRaw( raw, fullName, importInfo, localFlag ) .. ">"
end
function BoxTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function BoxTypeInfo:get_display_stirng_with( raw )

   return string.format( "Nilable<%s>", self.boxingType:get_display_stirng_with( raw ))
end
function BoxTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, accessMode = %d, boxingType = %d }\n', SerializeKind.Box, parentId, self.typeId, self.accessMode, self.boxingType:get_typeId()) )
end
function BoxTypeInfo:equals( typeInfo, alt2type )

   do
      local boxType = _lune.__Cast( typeInfo, 3, BoxTypeInfo )
      if boxType ~= nil then
         return self.boxingType:equals( boxType.boxingType, alt2type )
      end
   end
   
   return false
end
function BoxTypeInfo:createAlt2typeMap( detectFlag )

   local map = CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( detectFlag )
   if self.boxingType ~= boxRootAltType then
      map[boxRootAltType] = self.boxingType
   end
   
   return map
end
function BoxTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = BoxTypeInfo  } )
end
function BoxTypeInfo:get_boxingType()       
   return self.boxingType         
end
function BoxTypeInfo:get_typeId()       
   return self.typeId         
end
function BoxTypeInfo:get_itemTypeInfoList()       
   return self.itemTypeInfoList         
end
function BoxTypeInfo:get_accessMode()       
   return self.accessMode         
end
function BoxTypeInfo:get_nilableTypeInfo()       
   return self.nilableTypeInfo         
end
function BoxTypeInfo:isModule( ... )
   return self.boxingType:isModule( ... )
end       

function BoxTypeInfo:getParentId( ... )
   return self.boxingType:getParentId( ... )
end       

function BoxTypeInfo:get_baseId( ... )
   return self.boxingType:get_baseId( ... )
end       

function BoxTypeInfo:isInheritFrom( ... )
   return self.boxingType:isInheritFrom( ... )
end       

function BoxTypeInfo:get_rawTxt( ... )
   return self.boxingType:get_rawTxt( ... )
end       

function BoxTypeInfo:get_abstractFlag( ... )
   return self.boxingType:get_abstractFlag( ... )
end       

function BoxTypeInfo:get_externalFlag( ... )
   return self.boxingType:get_externalFlag( ... )
end       

function BoxTypeInfo:get_interfaceList( ... )
   return self.boxingType:get_interfaceList( ... )
end       

function BoxTypeInfo:get_argTypeInfoList( ... )
   return self.boxingType:get_argTypeInfoList( ... )
end       

function BoxTypeInfo:get_retTypeInfoList( ... )
   return self.boxingType:get_retTypeInfoList( ... )
end       

function BoxTypeInfo:get_parentInfo( ... )
   return self.boxingType:get_parentInfo( ... )
end       

function BoxTypeInfo:hasRouteNamespaceFrom( ... )
   return self.boxingType:hasRouteNamespaceFrom( ... )
end       

function BoxTypeInfo:getModule( ... )
   return self.boxingType:getModule( ... )
end       

function BoxTypeInfo:get_staticFlag( ... )
   return self.boxingType:get_staticFlag( ... )
end       

function BoxTypeInfo:get_autoFlag( ... )
   return self.boxingType:get_autoFlag( ... )
end       

function BoxTypeInfo:get_baseTypeInfo( ... )
   return self.boxingType:get_baseTypeInfo( ... )
end       

function BoxTypeInfo:get_children( ... )
   return self.boxingType:get_children( ... )
end       

function BoxTypeInfo:addChildren( ... )
   return self.boxingType:addChildren( ... )
end       

function BoxTypeInfo:get_mutMode( ... )
   return self.boxingType:get_mutMode( ... )
end       

function BoxTypeInfo:getParentFullName( ... )
   return self.boxingType:getParentFullName( ... )
end       

function BoxTypeInfo:get_genSrcTypeInfo( ... )
   return self.boxingType:get_genSrcTypeInfo( ... )
end       

function BoxTypeInfo:serializeTypeInfoList( ... )
   return self.boxingType:serializeTypeInfoList( ... )
end       

function BoxTypeInfo:get_typeData( ... )
   return self.boxingType:get_typeData( ... )
end       

function BoxTypeInfo:getFullName( ... )
   return self.boxingType:getFullName( ... )
end       


local GenericTypeInfo = {}
setmetatable( GenericTypeInfo, { __index = TypeInfo } )
_moduleObj.GenericTypeInfo = GenericTypeInfo
function GenericTypeInfo.new( genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo )
   local obj = {}
   GenericTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo ); end
   return obj
end
function GenericTypeInfo:__init(genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo) 
   TypeInfo.__init( self,TypeInfo.createScope( (_lune.unwrap( getScope( genSrcTypeInfo )) ):get_parent(), true, genSrcTypeInfo, nil ))
   
   idProv:increment(  )
   self.typeId = idProv:get_id()
   self.moduleTypeInfo = moduleTypeInfo
   self.itemTypeInfoList = itemTypeInfoList
   self.genSrcTypeInfo = genSrcTypeInfo
   if #genSrcTypeInfo:get_itemTypeInfoList() ~= #itemTypeInfoList then
      Util.err( string.format( "unmatch generic type number -- %d, %d", #genSrcTypeInfo:get_itemTypeInfoList(), #itemTypeInfoList) )
   end
   
   local alt2typeMap = {}
   local workAlt2typeMap = CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   local hasAlter = false
   for index, altTypeInfo in pairs( genSrcTypeInfo:get_itemTypeInfoList() ) do
      local itemType = itemTypeInfoList[index]
      alt2typeMap[altTypeInfo] = itemType
      if itemType:applyGeneric( workAlt2typeMap ) ~= itemType then
         hasAlter = true
      end
      
   end
   
   self.hasAlter = hasAlter
   self.alt2typeMap = alt2typeMap
   idProv:increment(  )
   self.nilableTypeInfo = NilableTypeInfo.new(self, idProv:get_id())
end
function GenericTypeInfo:getModule(  )

   return self.moduleTypeInfo
end
function GenericTypeInfo:isInheritFrom( other, alt2type )

   local otherSrc = other:get_genSrcTypeInfo()
   if not self.genSrcTypeInfo:isInheritFrom( otherSrc, alt2type ) then
      return false
   end
   
   local genOther = _lune.__Cast( other, 3, GenericTypeInfo )
   if  nil == genOther then
      local _genOther = genOther
   
      return true
   end
   
   local workAlt2type = alt2type
   if  nil == workAlt2type then
      local _workAlt2type = workAlt2type
   
      workAlt2type = CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   end
   
   for __index, altType in pairs( otherSrc:get_itemTypeInfoList() ) do
      local genType = self.alt2typeMap[altType]
      if  nil == genType then
         local _genType = genType
      
         return false
      end
      
      local otherGenType = _lune.unwrap( genOther.alt2typeMap[altType])
      if not otherGenType:canEvalWith( genType, "=", workAlt2type ) then
         return false
      end
      
   end
   
   return true
end
function GenericTypeInfo:get_srcTypeInfo(  )

   return self
end
function GenericTypeInfo:canEvalWith( other, opTxt, alt2type )

   if TypeInfo.isMut( self ) and not TypeInfo.isMut( other ) then
      return false
   end
   
   local otherSrc = other:get_srcTypeInfo()
   if self == otherSrc then
      return true
   end
   
   local work = otherSrc
   while true do
      if work == _moduleObj.headTypeInfo then
         return false
      end
      
      for altType, genType in pairs( work:createAlt2typeMap( false ) ) do
         alt2type[altType] = genType
      end
      
      if self.genSrcTypeInfo:equals( work:get_genSrcTypeInfo(), alt2type ) then
         break
      end
      
      for __index, ifType in pairs( work:get_interfaceList() ) do
         if self:canEvalWith( ifType, opTxt, alt2type ) then
            return true
         end
         
      end
      
      work = work:get_baseTypeInfo()
   end
   
   do
      local otherGen = _lune.__Cast( work, 3, GenericTypeInfo )
      if otherGen ~= nil then
         for key, val in pairs( self.alt2typeMap ) do
            local otherType = AlternateTypeInfo.getAssign( _lune.unwrap( otherGen.alt2typeMap[key]), alt2type )
            if not val:canEvalWith( otherType, opTxt, alt2type ) then
               return false
            end
            
         end
         
      end
   end
   
   return true
end
function GenericTypeInfo:equals( other, alt2type )

   if self == other then
      return true
   end
   
   if self:get_kind() ~= self:get_kind() or #self.itemTypeInfoList ~= #other:get_itemTypeInfoList() then
      return false
   end
   
   if not (_lune.__Cast( other, 3, GenericTypeInfo ) ) then
      return false
   end
   
   if not self.genSrcTypeInfo:equals( other:get_genSrcTypeInfo(), alt2type ) then
      return false
   end
   
   for index, otherItem in pairs( other:get_itemTypeInfoList() ) do
      local typeInfo = self.itemTypeInfoList[index]
      if not typeInfo:equals( otherItem, alt2type ) then
         return false
      end
      
   end
   
   return true
end
function GenericTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, genSrcTypeId = %d, genTypeList = {', SerializeKind.Generic, parentId, self.typeId, self.genSrcTypeInfo:get_typeId()) )
   local count = 0
   for __index, genType in pairs( self.alt2typeMap ) do
      if count > 0 then
         stream:write( "," )
      end
      
      stream:write( string.format( "%d", genType:get_typeId()) )
   end
   
   stream:write( '} }\n' )
end
function GenericTypeInfo:createAlt2typeMap( detectFlag )

   local map = self.genSrcTypeInfo:createAlt2typeMap( detectFlag )
   for genType, typeInfo in pairs( self.alt2typeMap ) do
      map[genType] = typeInfo
   end
   
   return map
end
function GenericTypeInfo:applyGeneric( alt2typeMap )

   if not self.hasAlter then
      return self
   end
   
   local genSrcTypeInfo = self.genSrcTypeInfo:applyGeneric( alt2typeMap )
   if genSrcTypeInfo == self.genSrcTypeInfo then
      return self
   end
   
   Util.errorLog( string.format( "no support nest generic -- %s", self:getTxt(  )) )
   return nil
end
function GenericTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = GenericTypeInfo  } )
end
function GenericTypeInfo:get_typeId()       
   return self.typeId         
end
function GenericTypeInfo:get_itemTypeInfoList()       
   return self.itemTypeInfoList         
end
function GenericTypeInfo:get_nilableTypeInfo()       
   return self.nilableTypeInfo         
end
function GenericTypeInfo:get_genSrcTypeInfo()       
   return self.genSrcTypeInfo         
end
function GenericTypeInfo:isModule( ... )
   return self.genSrcTypeInfo:isModule( ... )
end       

function GenericTypeInfo:getParentId( ... )
   return self.genSrcTypeInfo:getParentId( ... )
end       

function GenericTypeInfo:get_baseId( ... )
   return self.genSrcTypeInfo:get_baseId( ... )
end       

function GenericTypeInfo:get_rawTxt( ... )
   return self.genSrcTypeInfo:get_rawTxt( ... )
end       

function GenericTypeInfo:getTxtWithRaw( ... )
   return self.genSrcTypeInfo:getTxtWithRaw( ... )
end       

function GenericTypeInfo:getTxt( ... )
   return self.genSrcTypeInfo:getTxt( ... )
end       

function GenericTypeInfo:get_abstractFlag( ... )
   return self.genSrcTypeInfo:get_abstractFlag( ... )
end       

function GenericTypeInfo:get_display_stirng_with( ... )
   return self.genSrcTypeInfo:get_display_stirng_with( ... )
end       

function GenericTypeInfo:get_display_stirng( ... )
   return self.genSrcTypeInfo:get_display_stirng( ... )
end       

function GenericTypeInfo:get_externalFlag( ... )
   return self.genSrcTypeInfo:get_externalFlag( ... )
end       

function GenericTypeInfo:get_interfaceList( ... )
   return self.genSrcTypeInfo:get_interfaceList( ... )
end       

function GenericTypeInfo:get_argTypeInfoList( ... )
   return self.genSrcTypeInfo:get_argTypeInfoList( ... )
end       

function GenericTypeInfo:get_retTypeInfoList( ... )
   return self.genSrcTypeInfo:get_retTypeInfoList( ... )
end       

function GenericTypeInfo:get_parentInfo( ... )
   return self.genSrcTypeInfo:get_parentInfo( ... )
end       

function GenericTypeInfo:hasRouteNamespaceFrom( ... )
   return self.genSrcTypeInfo:hasRouteNamespaceFrom( ... )
end       

function GenericTypeInfo:get_kind( ... )
   return self.genSrcTypeInfo:get_kind( ... )
end       

function GenericTypeInfo:get_staticFlag( ... )
   return self.genSrcTypeInfo:get_staticFlag( ... )
end       

function GenericTypeInfo:get_accessMode( ... )
   return self.genSrcTypeInfo:get_accessMode( ... )
end       

function GenericTypeInfo:get_autoFlag( ... )
   return self.genSrcTypeInfo:get_autoFlag( ... )
end       

function GenericTypeInfo:get_nonnilableType( ... )
   return self.genSrcTypeInfo:get_nonnilableType( ... )
end       

function GenericTypeInfo:get_baseTypeInfo( ... )
   return self.genSrcTypeInfo:get_baseTypeInfo( ... )
end       

function GenericTypeInfo:get_nilable( ... )
   return self.genSrcTypeInfo:get_nilable( ... )
end       

function GenericTypeInfo:get_children( ... )
   return self.genSrcTypeInfo:get_children( ... )
end       

function GenericTypeInfo:addChildren( ... )
   return self.genSrcTypeInfo:addChildren( ... )
end       

function GenericTypeInfo:get_mutMode( ... )
   return self.genSrcTypeInfo:get_mutMode( ... )
end       

function GenericTypeInfo:getParentFullName( ... )
   return self.genSrcTypeInfo:getParentFullName( ... )
end       

function GenericTypeInfo:serializeTypeInfoList( ... )
   return self.genSrcTypeInfo:serializeTypeInfoList( ... )
end       

function GenericTypeInfo:get_scope( ... )
   return self.genSrcTypeInfo:get_scope( ... )
end       

function GenericTypeInfo:get_typeData( ... )
   return self.genSrcTypeInfo:get_typeData( ... )
end       

function GenericTypeInfo:getFullName( ... )
   return self.genSrcTypeInfo:getFullName( ... )
end       


local function isGenericType( typeInfo )

   if _lune.__Cast( typeInfo, 3, GenericTypeInfo ) then
      return true
   end
   
   return false
end
_moduleObj.isGenericType = isGenericType
local ModifierTypeInfo = {}
setmetatable( ModifierTypeInfo, { __index = TypeInfo } )
_moduleObj.ModifierTypeInfo = ModifierTypeInfo
function ModifierTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function ModifierTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   local txt = self.srcTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )
   if self.mutMode == MutMode.IMut then
      txt = "&" .. txt
   end
   
   return txt
end
function ModifierTypeInfo:get_display_stirng_with( raw )

   local txt = self.srcTypeInfo:get_display_stirng_with( raw )
   if self.mutMode ~= MutMode.IMut then
      txt = "mut " .. txt
   end
   
   return txt
end
function ModifierTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function ModifierTypeInfo:serialize( stream, validChildrenSet )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, srcTypeId = %d, mutMode = %d }\n', SerializeKind.Modifier, parentId, self.typeId, self.srcTypeInfo:get_typeId(), self.mutMode) )
end
function ModifierTypeInfo:canEvalWith( other, opTxt, alt2type )

   return TypeInfo.canEvalWithBase( self.srcTypeInfo, TypeInfo.isMut( self ), other:get_srcTypeInfo(), opTxt, alt2type )
end
function ModifierTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = ModifierTypeInfo  } )
end
function ModifierTypeInfo.new( srcTypeInfo, typeId, mutMode )
   local obj = {}
   ModifierTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( srcTypeInfo, typeId, mutMode )
   end        
   return obj 
end         
function ModifierTypeInfo:__init( srcTypeInfo, typeId, mutMode ) 

   TypeInfo.__init( self )
   self.srcTypeInfo = srcTypeInfo
   self.typeId = typeId
   self.mutMode = mutMode
end
function ModifierTypeInfo:get_srcTypeInfo()       
   return self.srcTypeInfo         
end
function ModifierTypeInfo:get_typeId()       
   return self.typeId         
end
function ModifierTypeInfo:get_mutMode()       
   return self.mutMode         
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

function ModifierTypeInfo:get_rawTxt( ... )
   return self.srcTypeInfo:get_rawTxt( ... )
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

function ModifierTypeInfo:get_children( ... )
   return self.srcTypeInfo:get_children( ... )
end       

function ModifierTypeInfo:addChildren( ... )
   return self.srcTypeInfo:addChildren( ... )
end       

function ModifierTypeInfo:getParentFullName( ... )
   return self.srcTypeInfo:getParentFullName( ... )
end       

function ModifierTypeInfo:get_genSrcTypeInfo( ... )
   return self.srcTypeInfo:get_genSrcTypeInfo( ... )
end       

function ModifierTypeInfo:serializeTypeInfoList( ... )
   return self.srcTypeInfo:serializeTypeInfoList( ... )
end       

function ModifierTypeInfo:get_scope( ... )
   return self.srcTypeInfo:get_scope( ... )
end       

function ModifierTypeInfo:get_typeData( ... )
   return self.srcTypeInfo:get_typeData( ... )
end       

function ModifierTypeInfo:createAlt2typeMap( ... )
   return self.srcTypeInfo:createAlt2typeMap( ... )
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
   TypeInfo.__init( self,scope)
   
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

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function ModuleTypeInfo:getTxtWithRaw( rawTxt, fullName, importInfo, localFlag )

   return rawTxt
end
function ModuleTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function ModuleTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function ModuleTypeInfo:canEvalWith( other, opTxt, alt2type )

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

local EnumLiteral = {}
EnumLiteral._name2Val = {}
_moduleObj.EnumLiteral = EnumLiteral
function EnumLiteral:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "EnumLiteral.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 

function EnumLiteral._from( val )
   return _lune._AlgeFrom( EnumLiteral, val )
end

EnumLiteral.Int = { "Int", {{ func=_lune._toInt, nilable=false, child={} }}}
EnumLiteral._name2Val["Int"] = EnumLiteral.Int
EnumLiteral.Real = { "Real", {{ func=_lune._toReal, nilable=false, child={} }}}
EnumLiteral._name2Val["Real"] = EnumLiteral.Real
EnumLiteral.Str = { "Str", {{ func=_lune._toStr, nilable=false, child={} }}}
EnumLiteral._name2Val["Str"] = EnumLiteral.Str

local function getEnumLiteralVal( obj )

   do
      local _matchExp = obj
      if _matchExp[1] == EnumLiteral.Int[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == EnumLiteral.Real[1] then
         local val = _matchExp[2][1]
      
         return val
      elseif _matchExp[1] == EnumLiteral.Str[1] then
         local val = _matchExp[2][1]
      
         return val
      end
   end
   
   Util.err( "illegal enum " .. EnumLiteral:_getTxt( obj)
    )
end
_moduleObj.getEnumLiteralVal = getEnumLiteralVal
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
function EnumTypeInfo.new( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo )
   local obj = {}
   EnumTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo ); end
   return obj
end
function EnumTypeInfo:__init(scope, externalFlag, accessMode, txt, parentInfo, typeId, valTypeInfo) 
   TypeInfo.__init( self,scope)
   
   self.externalFlag = externalFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = typeId
   self.name2EnumValInfo = {}
   self.valTypeInfo = valTypeInfo
   self.val2EnumValInfo = {}
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

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function EnumTypeInfo:getTxtWithRaw( rawTxt, fullName, importInfo, localFlag )

   return rawTxt
end
function EnumTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function EnumTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function EnumTypeInfo:canEvalWith( other, opTxt, alt2type )

   return self == other:get_srcTypeInfo()
end
function EnumTypeInfo:addEnumValInfo( valInfo )

   self.name2EnumValInfo[valInfo:get_name()] = valInfo
   self.val2EnumValInfo[getEnumLiteralVal( valInfo:get_val() )] = valInfo
end
function EnumTypeInfo:getEnumValInfo( name )

   return self.name2EnumValInfo[name]
end
function EnumTypeInfo:get_mutMode(  )

   return MutMode.Mut
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
function EnumTypeInfo:get_val2EnumValInfo()       
   return self.val2EnumValInfo         
end

local AlgeValInfo = {}
_moduleObj.AlgeValInfo = AlgeValInfo
function AlgeValInfo:serialize( stream )

   stream:write( string.format( "{ name = '%s', typeList = {", self.name) )
   for index, typeInfo in pairs( self.typeList ) do
      if index > 1 then
         stream:write( ", " )
      end
      
      stream:write( string.format( "%d", typeInfo:get_typeId()) )
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
   TypeInfo.__init( self,scope)
   
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

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function AlgeTypeInfo:getTxtWithRaw( rawTxt, fullName, importInfo, localFlag )

   return rawTxt
end
function AlgeTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function AlgeTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function AlgeTypeInfo:canEvalWith( other, opTxt, alt2type )

   return self == other:get_srcTypeInfo()
end
function AlgeTypeInfo:get_mutMode(  )

   return MutMode.Mut
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
function NormalTypeInfo.new( abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode )
   local obj = {}
   NormalTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode ); end
   return obj
end
function NormalTypeInfo:__init(abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode) 
   TypeInfo.__init( self,scope)
   
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
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.mutMode = _lune.unwrapDefault( mutMode, MutMode.IMut)
   local function setupAlt2typeMap(  )
   
      if self.baseTypeInfo == _moduleObj.headTypeInfo and #self.interfaceList == 0 then
         return {}
      end
      
      local alt2typeMap = {}
      do
         local _switchExp = kind
         if _switchExp == TypeInfoKind.Set or _switchExp == TypeInfoKind.Map or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Box then
            if #self.itemTypeInfoList ~= #self.baseTypeInfo:get_itemTypeInfoList() then
               Util.err( string.format( "unmatch generic type number -- %d, %d", #self.itemTypeInfoList, #self.baseTypeInfo:get_itemTypeInfoList()) )
            end
            
            for index, appyType in pairs( self.itemTypeInfoList ) do
               local genType = self.baseTypeInfo:get_itemTypeInfoList()[index]
               alt2typeMap[genType] = appyType
            end
            
         elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
            for __index, ifType in pairs( self.interfaceList ) do
               do
                  local genericType = _lune.__Cast( ifType, 3, GenericTypeInfo )
                  if genericType ~= nil then
                     for altType, genType in pairs( genericType:createAlt2typeMap( false ) ) do
                        alt2typeMap[altType] = genType
                     end
                     
                  end
               end
               
            end
            
         end
      end
      
      return alt2typeMap
   end
   
   self.alt2typeMap = setupAlt2typeMap(  )
   self.typeId = typeId
   if kind == TypeInfoKind.Root then
   else
    
      if parentInfo ~= nil then
         parentInfo:addChildren( self )
      end
      
      local hasNilable = false
      do
         local _switchExp = (kind )
         if _switchExp == TypeInfoKind.Prim or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set or _switchExp == TypeInfoKind.Map or _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Stem or _switchExp == TypeInfoKind.Module or _switchExp == TypeInfoKind.IF then
            hasNilable = true
         elseif _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method then
            hasNilable = true
         end
      end
      
      if hasNilable then
         self.nilableTypeInfo = NilableTypeInfo.new(self, typeId + 1)
         idProv:increment(  )
      else
       
         self.nilableTypeInfo = _moduleObj.headTypeInfo
      end
      
      idProv:increment(  )
   end
   
end
function NormalTypeInfo:createAlt2typeMap( detectFlag )

   local map = self.baseTypeInfo:createAlt2typeMap( detectFlag )
   for genType, typeInfo in pairs( self.alt2typeMap ) do
      map[genType] = typeInfo
   end
   
   return map
end
function NormalTypeInfo:get_nilable(  )

   return false
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

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function NormalTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   local parentTxt = ""
   if fullName then
      parentTxt = self:getParentFullName( importInfo, localFlag )
   end
   
   if #self.itemTypeInfoList > 0 then
      local txt = raw .. "<"
      for index, typeInfo in pairs( self.itemTypeInfoList ) do
         if index ~= 1 then
            txt = txt .. ","
         end
         
         txt = txt .. typeInfo:getTxt( fullName, importInfo, localFlag )
      end
      
      return parentTxt .. txt .. ">"
   end
   
   return parentTxt .. raw
end
function NormalTypeInfo:get_display_stirng_with( raw )

   if self.kind == TypeInfoKind.Func or self.kind == TypeInfoKind.Method then
      local txt = raw .. "("
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
   
   return self:getTxtWithRaw( raw )
end
function NormalTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function NormalTypeInfo:serialize( stream, validChildrenSet )

   if self.typeId == _moduleObj.rootTypeId then
      return 
   end
   
   local parentId = self:getParentId(  )
   local txt = string.format( [==[{ skind=%d, parentId = %d, typeId = %d, baseId = %d, txt = '%s',
        abstractFlag = %s, staticFlag = %s, accessMode = %d, kind = %d, mutMode = %d, ]==], SerializeKind.Normal, parentId, self.typeId, self:get_baseId(  ), self.rawTxt, tostring( self.abstractFlag), tostring( self.staticFlag), self.accessMode, self.kind, self.mutMode)
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
   
   stream:write( txt .. self:serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. self:serializeTypeInfoList( "ifList = {", self.interfaceList ) .. self:serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. self:serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. self:serializeTypeInfoList( "children = {", children, true ) .. "}\n" )
end
function NormalTypeInfo:equalsSub( typeInfo, alt2type )

   if self.typeId == typeInfo:get_typeId() then
      return true
   end
   
   if self.kind ~= typeInfo:get_kind() or self.staticFlag ~= typeInfo:get_staticFlag() or self.accessMode ~= typeInfo:get_accessMode() or self.autoFlag ~= typeInfo:get_autoFlag() or self:get_nilable() ~= typeInfo:get_nilable() or self.rawTxt ~= typeInfo:get_rawTxt() or self.parentInfo ~= typeInfo:get_parentInfo() or self.baseTypeInfo ~= typeInfo:get_baseTypeInfo() then
      return false
   end
   
   do
      if #self.itemTypeInfoList ~= #typeInfo:get_itemTypeInfoList() then
         return false
      end
      
      for index, item in pairs( self.itemTypeInfoList ) do
         if not item:equals( typeInfo:get_itemTypeInfoList()[index], alt2type ) then
            return false
         end
         
      end
      
   end
   
   do
      if #self.retTypeInfoList ~= #typeInfo:get_retTypeInfoList() then
         return false
      end
      
      for index, item in pairs( self.retTypeInfoList ) do
         if not item:equals( typeInfo:get_retTypeInfoList()[index], alt2type ) then
            return false
         end
         
      end
      
   end
   
   return true
end
function NormalTypeInfo:equals( typeInfo, alt2type )

   return self:equalsSub( typeInfo, alt2type )
end
function NormalTypeInfo.create( accessMode, abstractFlag, scope, baseInfo, interfaceList, parentInfo, staticFlag, kind, txt, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode )

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
   local info = NormalTypeInfo.new(abstractFlag, scope, baseInfo, interfaceList, false, true, staticFlag, accessMode, txt, parentInfo, idProv:get_id(), kind, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode)
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
function NormalTypeInfo:get_baseTypeInfo()       
   return self.baseTypeInfo         
end
function NormalTypeInfo:get_interfaceList()       
   return self.interfaceList         
end
function NormalTypeInfo:get_nilableTypeInfo()       
   return self.nilableTypeInfo         
end
function NormalTypeInfo:get_mutMode()       
   return self.mutMode         
end

function NormalTypeInfo.createAlternate( txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList )

   return AlternateTypeInfo.new(txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList)
end


local TypeInfo2Map = {}
_moduleObj.TypeInfo2Map = TypeInfo2Map
function TypeInfo2Map.new(  )
   local obj = {}
   TypeInfo2Map.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function TypeInfo2Map:__init() 
   self.ImutModifierMap = {}
   self.MutModifierMap = {}
   self.BoxMap = {}
   self.DDDMap = {}
end
function TypeInfo2Map:clone(  )

   local obj = TypeInfo2Map.new()
   for key, val in pairs( self.ImutModifierMap ) do
      obj.ImutModifierMap[key] = val
   end
   
   for key, val in pairs( self.MutModifierMap ) do
      obj.MutModifierMap[key] = val
   end
   
   for key, val in pairs( self.BoxMap ) do
      obj.BoxMap[key] = val
   end
   
   for key, val in pairs( self.DDDMap ) do
      obj.DDDMap[key] = val
   end
   
   return obj
end
function TypeInfo2Map.setmeta( obj )
  setmetatable( obj, { __index = TypeInfo2Map  } )
end
function TypeInfo2Map:get_ImutModifierMap()       
   return self.ImutModifierMap         
end
function TypeInfo2Map:get_MutModifierMap()       
   return self.MutModifierMap         
end
function TypeInfo2Map:get_BoxMap()       
   return self.BoxMap         
end
function TypeInfo2Map:get_DDDMap()       
   return self.DDDMap         
end

local typeInfo2Map = TypeInfo2Map.new()
function NormalTypeInfo.createModifier( srcTypeInfo, mutMode )

   srcTypeInfo = srcTypeInfo:get_srcTypeInfo()
   do
      local _switchExp = mutMode
      if _switchExp == MutMode.IMut then
         do
            local _exp = typeInfo2Map.ImutModifierMap[srcTypeInfo]
            if _exp ~= nil then
               if _exp:get_typeId() < userStartId and srcTypeInfo:get_typeId() >= userStartId then
                  Util.err( "on cache" )
               end
               
               return _exp
            end
         end
         
      elseif _switchExp == MutMode.AllMut then
         do
            local _exp = typeInfo2Map.MutModifierMap[srcTypeInfo]
            if _exp ~= nil then
               if _exp:get_typeId() < userStartId and srcTypeInfo:get_typeId() >= userStartId then
                  Util.err( "on cache" )
               end
               
               return _exp
            end
         end
         
      end
   end
   
   idProv:increment(  )
   local modifier = ModifierTypeInfo.new(srcTypeInfo, idProv:get_id(), mutMode)
   do
      local _switchExp = mutMode
      if _switchExp == MutMode.IMut then
         typeInfo2Map.ImutModifierMap[srcTypeInfo] = modifier
      elseif _switchExp == MutMode.AllMut then
         typeInfo2Map.MutModifierMap[srcTypeInfo] = modifier
      end
   end
   
   if modifier:get_typeId() < userStartId and srcTypeInfo:get_typeId() >= userStartId then
      Util.printStackTrace(  )
      Util.err( string.format( "off cache: %s %s %s", srcTypeInfo:getTxt(  ), tostring( modifier:get_typeId()), tostring( srcTypeInfo:get_typeId())) )
   end
   
   return modifier
end

idProv:increment(  )
local function addBuiltin( typeInfo )

   _moduleObj.builtInTypeIdSet[typeInfo:get_typeId()] = typeInfo
end

local function registBuiltin( idName, typeTxt, kind, typeInfo, nilableTypeInfo, registScope )

   _moduleObj.sym2builtInTypeMap[typeTxt] = NormalSymbolInfo.new(SymbolKind.Typ, false, false, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, typeInfo, MutMode.IMut, true)
   if nilableTypeInfo ~= _moduleObj.headTypeInfo then
      _moduleObj.sym2builtInTypeMap[typeTxt .. "!"] = NormalSymbolInfo.new(SymbolKind.Typ, false, kind == TypeInfoKind.Func, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, nilableTypeInfo, MutMode.IMut, true)
   end
   
   addBuiltin( typeInfo )
   addBuiltin( NormalTypeInfo.createModifier( typeInfo, MutMode.IMut ) )
   if typeInfo:get_nilableTypeInfo() ~= _moduleObj.headTypeInfo then
      addBuiltin( typeInfo:get_nilableTypeInfo() )
      addBuiltin( NormalTypeInfo.createModifier( typeInfo:get_nilableTypeInfo(), MutMode.IMut ) )
   end
   
   if registScope then
      _moduleObj.rootScope:addClass( typeTxt, typeInfo )
   end
   
end

function NormalTypeInfo.createBuiltin( idName, typeTxt, kind, typeDDD, ifList )

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
      if _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Set or _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Module or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method or _switchExp == TypeInfoKind.Macro then
         scope = Scope.new(_moduleObj.rootScope, kind == TypeInfoKind.Class or kind == TypeInfoKind.Module or kind == TypeInfoKind.IF or kind == TypeInfoKind.List or kind == TypeInfoKind.Array or kind == TypeInfoKind.Set, nil)
      end
   end
   
   local genTypeList = {}
   do
      local _switchExp = kind
      if _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Set then
         table.insert( genTypeList, NormalTypeInfo.createAlternate( "T", AccessMode.Pri, _moduleObj.headTypeInfo ) )
      elseif _switchExp == TypeInfoKind.Map then
         table.insert( genTypeList, NormalTypeInfo.createAlternate( "K", AccessMode.Pri, _moduleObj.headTypeInfo ) )
         table.insert( genTypeList, NormalTypeInfo.createAlternate( "V", AccessMode.Pri, _moduleObj.headTypeInfo ) )
      end
   end
   
   local info = NormalTypeInfo.new(false, scope, nil, ifList, false, false, false, AccessMode.Pub, typeTxt, _moduleObj.headTypeInfo, typeId, kind, genTypeList, argTypeList, retTypeList, MutMode.Mut)
   registBuiltin( idName, typeTxt, kind, info, _moduleObj.headTypeInfo, scope ~= nil )
   return info
end

local builtinTypeNone = NormalTypeInfo.createBuiltin( "None", "", TypeInfoKind.Prim )
_moduleObj.builtinTypeNone = builtinTypeNone

local builtinTypeNeverRet = NormalTypeInfo.createBuiltin( "Error", "__", TypeInfoKind.Prim )
_moduleObj.builtinTypeNeverRet = builtinTypeNeverRet

local builtinTypeStem = NormalTypeInfo.createBuiltin( "Stem", "stem", TypeInfoKind.Stem )
_moduleObj.builtinTypeStem = builtinTypeStem

local builtinTypeStem_ = _moduleObj.builtinTypeStem:get_nilableTypeInfo()
_moduleObj.builtinTypeStem_ = builtinTypeStem_

local builtinTypeBool = NormalTypeInfo.createBuiltin( "Bool", "bool", TypeInfoKind.Prim )
_moduleObj.builtinTypeBool = builtinTypeBool

local builtinTypeInt = NormalTypeInfo.createBuiltin( "Int", "int", TypeInfoKind.Prim )
_moduleObj.builtinTypeInt = builtinTypeInt

local builtinTypeReal = NormalTypeInfo.createBuiltin( "Real", "real", TypeInfoKind.Prim )
_moduleObj.builtinTypeReal = builtinTypeReal

local builtinTypeChar = NormalTypeInfo.createBuiltin( "char", "char", TypeInfoKind.Prim )
_moduleObj.builtinTypeChar = builtinTypeChar

local builtinTypeMapping = NormalTypeInfo.createBuiltin( "Mapping", "Mapping", TypeInfoKind.IF )
_moduleObj.builtinTypeMapping = builtinTypeMapping

local builtinTypeString = NormalTypeInfo.createBuiltin( "String", "str", TypeInfoKind.Class, nil, {_moduleObj.builtinTypeMapping} )
_moduleObj.builtinTypeString = builtinTypeString

local builtinTypeMap = NormalTypeInfo.createBuiltin( "Map", "Map", TypeInfoKind.Map )
_moduleObj.builtinTypeMap = builtinTypeMap

local builtinTypeSet = NormalTypeInfo.createBuiltin( "Set", "Set", TypeInfoKind.Set )
_moduleObj.builtinTypeSet = builtinTypeSet

local builtinTypeList = NormalTypeInfo.createBuiltin( "List", "List", TypeInfoKind.List )
_moduleObj.builtinTypeList = builtinTypeList

local builtinTypeArray = NormalTypeInfo.createBuiltin( "Array", "Array", TypeInfoKind.Array )
_moduleObj.builtinTypeArray = builtinTypeArray

local function isStruct( typeInfo )

   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == TypeInfoKind.Class then
         if typeInfo == _moduleObj.builtinTypeString then
            return false
         end
         
         if typeInfo:get_baseTypeInfo() ~= _moduleObj.headTypeInfo or #typeInfo:get_interfaceList() ~= 0 or #typeInfo:get_children() ~= 1 then
            return false
         end
         
         return true
      end
   end
   
   return false
end
_moduleObj.isStruct = isStruct
idProv:increment(  )
local builtinTypeBox = BoxTypeInfo.new(idProv:get_id(), AccessMode.Pub, boxRootAltType)
_moduleObj.builtinTypeBox = builtinTypeBox

registBuiltin( "Nilable", "Nilable", TypeInfoKind.Box, _moduleObj.builtinTypeBox, _moduleObj.headTypeInfo, true )
local function isConditionalbe( typeInfo )

   if typeInfo:get_nilable() or typeInfo:equals( _moduleObj.builtinTypeBool, nil ) then
      return true
   end
   
   return false
end
_moduleObj.isConditionalbe = isConditionalbe
function NormalTypeInfo.createBox( accessMode, nonnilableType )

   do
      local boxType = typeInfo2Map.BoxMap[nonnilableType]
      if boxType ~= nil then
         return boxType
      end
   end
   
   idProv:increment(  )
   local boxType = BoxTypeInfo.new(idProv:get_id(), accessMode, nonnilableType)
   typeInfo2Map.BoxMap[nonnilableType] = boxType
   return boxType
end

function BoxTypeInfo:applyGeneric( alt2typeMap )

   local typeInfo = self.boxingType:applyGeneric( alt2typeMap )
   if typeInfo == self.boxingType then
      return self
   end
   
   if typeInfo ~= nil then
      return NormalTypeInfo.createBox( self.accessMode, typeInfo )
   end
   
   return nil
end

function NormalTypeInfo.createSet( accessMode, parentInfo, itemTypeInfo, mutMode )

   if #itemTypeInfo == 0 then
      Util.err( string.format( "illegal set type: %s", tostring( itemTypeInfo)) )
   end
   
   idProv:increment(  )
   return NormalTypeInfo.new(false, getScope( _moduleObj.builtinTypeSet ), _moduleObj.builtinTypeSet, nil, false, false, false, accessMode, "Set", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.Set, itemTypeInfo, nil, nil, mutMode)
end

function NormalTypeInfo.createList( accessMode, parentInfo, itemTypeInfo, mutMode )

   if #itemTypeInfo == 0 then
      Util.err( string.format( "illegal list type: %s", tostring( itemTypeInfo)) )
   end
   
   idProv:increment(  )
   return NormalTypeInfo.new(false, getScope( _moduleObj.builtinTypeList ), _moduleObj.builtinTypeList, nil, false, false, false, accessMode, "List", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.List, itemTypeInfo, nil, nil, mutMode)
end

function NormalTypeInfo.createArray( accessMode, parentInfo, itemTypeInfo, mutMode )

   idProv:increment(  )
   return NormalTypeInfo.new(false, getScope( _moduleObj.builtinTypeArray ), _moduleObj.builtinTypeArray, nil, false, false, false, accessMode, "Array", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.Array, itemTypeInfo, nil, nil, mutMode)
end

function NormalTypeInfo.createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo, mutMode )

   idProv:increment(  )
   return NormalTypeInfo.new(false, getScope( _moduleObj.builtinTypeMap ), _moduleObj.builtinTypeMap, nil, false, false, false, accessMode, "Map", _moduleObj.headTypeInfo, idProv:get_id(), TypeInfoKind.Map, {keyTypeInfo, valTypeInfo}, nil, nil, mutMode)
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

function NormalTypeInfo.createClass( classFlag, abstractFlag, scope, baseInfo, interfaceList, genTypeList, parentInfo, externalFlag, accessMode, className )

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
   local info = NormalTypeInfo.new(abstractFlag, scope, baseInfo, interfaceList, false, externalFlag, false, accessMode, className, parentInfo, idProv:get_id(), classFlag and TypeInfoKind.Class or TypeInfoKind.IF, genTypeList, nil, nil, MutMode.Mut)
   return info
end

function NormalTypeInfo.createFunc( abstractFlag, builtinFlag, scope, kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, altTypeList, argTypeList, retTypeInfoList, mutable )

   if not builtinFlag and Parser.isLuaKeyword( funcName ) then
      Util.err( string.format( "This symbol can not use for a function. -- %s", funcName) )
   end
   
   idProv:increment(  )
   local info = NormalTypeInfo.new(abstractFlag, scope, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, idProv:get_id(), kind, _lune.unwrapDefault( altTypeList, {}), _lune.unwrapDefault( argTypeList, {}), _lune.unwrapDefault( retTypeInfoList, {}), mutable and MutMode.Mut or MutMode.IMut)
   return info
end

function NormalTypeInfo.createAdvertiseMethodFrom( classTypeInfo, typeInfo )

   return NormalTypeInfo.createFunc( false, false, getScope( typeInfo ), typeInfo:get_kind(), classTypeInfo, true, false, false, typeInfo:get_accessMode(), typeInfo:get_rawTxt(), typeInfo:get_itemTypeInfoList(), typeInfo:get_argTypeInfoList(), typeInfo:get_retTypeInfoList(), TypeInfo.isMut( typeInfo ) )
end

function ModifierTypeInfo:get_nonnilableType(  )

   local orgType = self.srcTypeInfo:get_nonnilableType()
   if TypeInfo.isMut( self ) or not TypeInfo.isMut( orgType ) then
      return orgType
   end
   
   return NormalTypeInfo.createModifier( orgType, MutMode.IMut )
end

function ModifierTypeInfo:get_nilableTypeInfo(  )

   local orgType = self.srcTypeInfo:get_nilableTypeInfo()
   if not TypeInfo.isMut( orgType ) then
      return orgType
   end
   
   return NormalTypeInfo.createModifier( orgType, MutMode.IMut )
end

function NormalTypeInfo.createAlias( name, externalFlag, accessMode, parentInfo, typeInfo )

   idProv:increment(  )
   return AliasTypeInfo.new(name, accessMode, parentInfo, typeInfo:get_srcTypeInfo(), externalFlag, idProv:get_id())
end

function Scope:addAlias( name, externalFlag, accessMode, parentInfo, symbolInfo )

   local aliasType = NormalTypeInfo.createAlias( name, externalFlag, accessMode, parentInfo, symbolInfo:get_typeInfo():get_srcTypeInfo() )
   return self:add( symbolInfo:get_kind(), false, symbolInfo:get_canBeRight(), name, aliasType, accessMode, true, MutMode.IMut, true )
end

function Scope:addAliasForType( name, typeInfo )

   local skind = SymbolKind.Typ
   local canBeRight = false
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == TypeInfoKind.Func then
         skind = SymbolKind.Fun
         canBeRight = true
      elseif _switchExp == TypeInfoKind.Macro then
         skind = SymbolKind.Mac
      end
   end
   
   return self:add( skind, false, canBeRight, name, typeInfo, typeInfo:get_accessMode(), true, MutMode.IMut, true )
end

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
   TypeInfo.__init( self,nil)
   
   self.typeId = typeId
   self.typeInfo = typeInfo
   self.externalFlag = externalFlag
   self.itemTypeInfoList = {self.typeInfo}
   typeInfo2Map.DDDMap[typeInfo] = self
end
function DDDTypeInfo:isModule(  )

   return false
end
function DDDTypeInfo:canEvalWith( other, opTxt, alt2type )

   return self.typeInfo:canEvalWith( other, opTxt, alt2type )
end
function DDDTypeInfo:serialize( stream, validChildrenSet )

   stream:write( string.format( '{ skind=%d, typeId = %d, itemTypeId = %d, parentId = %d }\n', SerializeKind.DDD, self.typeId, self.typeInfo:get_typeId(), _moduleObj.headTypeInfo:get_typeId()) )
end
function DDDTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function DDDTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
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
function DDDTypeInfo:get_mutMode(  )

   return self.typeInfo:get_mutMode()
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
      local _exp = typeInfo2Map.DDDMap[typeInfo]
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

local builtinTypeNil = NilTypeInfo.new()
_moduleObj.builtinTypeNil = builtinTypeNil

registBuiltin( "Nil", "nil", TypeInfoKind.Prim, _moduleObj.builtinTypeNil, _moduleObj.headTypeInfo, false )
local builtinTypeDDD = NormalTypeInfo.createDDD( _moduleObj.builtinTypeStem_, true )
_moduleObj.builtinTypeDDD = builtinTypeDDD

registBuiltin( "DDD", "...", TypeInfoKind.DDD, _moduleObj.builtinTypeDDD, _moduleObj.headTypeInfo, false )
local builtinTypeForm = NormalTypeInfo.createBuiltin( "Form", "form", TypeInfoKind.Func, _moduleObj.builtinTypeDDD )
_moduleObj.builtinTypeForm = builtinTypeForm

local builtinTypeSymbol = NormalTypeInfo.createBuiltin( "Symbol", "sym", TypeInfoKind.Prim )
_moduleObj.builtinTypeSymbol = builtinTypeSymbol

local builtinTypeStat = NormalTypeInfo.createBuiltin( "Stat", "stat", TypeInfoKind.Prim )
_moduleObj.builtinTypeStat = builtinTypeStat

function DDDTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( "...", fullName, importInfo, localFlag )
end

function DDDTypeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )

   if self.typeInfo == _moduleObj.builtinTypeStem_ then
      return "..."
   end
   
   local txt = self.typeInfo:getTxtWithRaw( raw, fullName, importInfo, localFlag )
   return "...<" .. txt .. ">"
end

function NormalTypeInfo.createGeneric( genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo )

   idProv:increment(  )
   return GenericTypeInfo.new(genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo)
end

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
   TypeInfo.__init( self,nil)
   
   local typeId = idProvider:get_id() + 1
   idProvider:increment(  )
   self.typeId = typeId
   self.rawTxt = rawTxt
end
function AbbrTypeInfo:isModule(  )

   return false
end
function AbbrTypeInfo:getTxt( fullName, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), fullName, importInfo, localFlag )
end
function AbbrTypeInfo:getTxtWithRaw( rawTxt, fullName, importInfo, localFlag )

   return rawTxt
end
function AbbrTypeInfo:canEvalWith( other, opTxt, alt2type )

   return false
end
function AbbrTypeInfo:serialize( stream, validChildrenSet )

   Util.err( "illegal call" )
end
function AbbrTypeInfo:get_display_stirng_with( raw )

   return self:getTxtWithRaw( raw )
end
function AbbrTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt() )
end
function AbbrTypeInfo:getModule(  )

   return _moduleObj.headTypeInfo
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
function AbbrTypeInfo:get_mutMode(  )

   return MutMode.IMut
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
function AbbrTypeInfo:get_rawTxt()       
   return self.rawTxt         
end

local builtinTypeAbbr = AbbrTypeInfo.new(idProv, "##")
_moduleObj.builtinTypeAbbr = builtinTypeAbbr

local builtinTypeAbbrNone = AbbrTypeInfo.new(idProv, "[##]")
_moduleObj.builtinTypeAbbrNone = builtinTypeAbbrNone

local numberTypeSet = {}
numberTypeSet[_moduleObj.builtinTypeInt]= true
numberTypeSet[_moduleObj.builtinTypeChar]= true
numberTypeSet[_moduleObj.builtinTypeReal]= true
local function isNumberType( typeInfo )

   return _lune._Set_has(numberTypeSet, typeInfo:get_srcTypeInfo() )
end
_moduleObj.isNumberType = isNumberType
function NormalTypeInfo.createEnum( scope, parentInfo, externalFlag, accessMode, enumName, valTypeInfo )

   if Parser.isLuaKeyword( enumName ) then
      Util.err( string.format( "This symbol can not use for a enum. -- %s", enumName) )
   end
   
   idProv:increment(  )
   local info = EnumTypeInfo.new(scope, externalFlag, accessMode, enumName, parentInfo, idProv:get_id(), valTypeInfo)
   local getEnumName = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Method, info, true, externalFlag, false, AccessMode.Pub, "get__txt", nil, nil, {_moduleObj.builtinTypeString}, false )
   scope:addMethod( getEnumName, AccessMode.Pub, false, false )
   local fromVal = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Func, info, true, externalFlag, true, AccessMode.Pub, "_from", nil, {NormalTypeInfo.createModifier( valTypeInfo, MutMode.IMut )}, {info:get_nilableTypeInfo()}, false )
   scope:addFunc( fromVal, AccessMode.Pub, true, false )
   local allListType = NormalTypeInfo.createList( AccessMode.Pub, info, {info}, MutMode.IMut )
   local allList = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Func, info, true, externalFlag, true, AccessMode.Pub, "get__allList", nil, nil, {NormalTypeInfo.createModifier( allListType, MutMode.IMut )}, false )
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
            stream:write( string.format( "%s = ", enumValInfo:get_name()) )
            do
               local _matchExp = enumValInfo:get_val()
               if _matchExp[1] == EnumLiteral.Int[1] then
                  local val = _matchExp[2][1]
               
                  stream:write( string.format( "%d,", val) )
               elseif _matchExp[1] == EnumLiteral.Real[1] then
                  local val = _matchExp[2][1]
               
                  stream:write( string.format( "%g,", val) )
               elseif _matchExp[1] == EnumLiteral.Str[1] then
                  local val = _matchExp[2][1]
               
                  stream:write( string.format( "%q,", val) )
               end
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
   local getAlgeName = NormalTypeInfo.createFunc( false, true, nil, TypeInfoKind.Method, info, true, externalFlag, false, AccessMode.Pub, "get__txt", nil, nil, {_moduleObj.builtinTypeString}, false )
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

function BoxTypeInfo:canEvalWith( other, opTxt, alt2type )

   if self == other then
      return true
   end
   
   if opTxt ~= "=" then
      return false
   end
   
   if other:get_nilable() then
      CanEvalCtrlTypeInfo.updateNeedAutoBoxing( alt2type )
      return true
   end
   
   if self.boxingType:canEvalWith( other, opTxt, alt2type ) then
      CanEvalCtrlTypeInfo.updateNeedAutoBoxing( alt2type )
      return true
   end
   
   return false
end

function NilableTypeInfo:canEvalWith( other, opTxt, alt2type )

   local otherSrc = other
   if self == _moduleObj.builtinTypeStem_ then
      return true
   end
   
   if otherSrc == _moduleObj.builtinTypeNil or otherSrc:get_kind() == TypeInfoKind.Abbr then
      if self:get_nonnilableType():get_kind() == TypeInfoKind.Box then
         return self:get_nonnilableType():canEvalWith( otherSrc, opTxt, alt2type )
      end
      
      return true
   end
   
   if self.typeId == otherSrc:get_typeId() then
      return true
   end
   
   if otherSrc:get_nilable() then
      return self:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt, alt2type )
   end
   
   return self:get_nonnilableType():canEvalWith( otherSrc, opTxt, alt2type )
end



function NormalTypeInfo.isAvailableMapping( typeInfo, checkedTypeMap )

   local function isAvailableMappingSub(  )
   
      do
         local _switchExp = typeInfo:get_kind()
         if _switchExp == TypeInfoKind.Prim or _switchExp == TypeInfoKind.Enum then
            return true
         elseif _switchExp == TypeInfoKind.Alge then
            local algeTypeInfo = _lune.unwrap( (_lune.__Cast( typeInfo, 3, AlgeTypeInfo ) ))
            for __index, valInfo in pairs( algeTypeInfo:get_valInfoMap() ) do
               for __index, paramType in pairs( valInfo:get_typeList() ) do
                  if not NormalTypeInfo.isAvailableMapping( paramType, checkedTypeMap ) then
                     return false
                  end
                  
               end
               
            end
            
            return true
         elseif _switchExp == TypeInfoKind.Stem then
            return true
         elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
            if typeInfo:equals( _moduleObj.builtinTypeString ) then
               return true
            end
            
            return typeInfo:isInheritFrom( _moduleObj.builtinTypeMapping, nil )
         elseif _switchExp == TypeInfoKind.Alternate then
            return typeInfo:isInheritFrom( _moduleObj.builtinTypeMapping, nil )
         elseif _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set then
            return NormalTypeInfo.isAvailableMapping( typeInfo:get_itemTypeInfoList()[1], checkedTypeMap )
         elseif _switchExp == TypeInfoKind.Map then
            if NormalTypeInfo.isAvailableMapping( typeInfo:get_itemTypeInfoList()[2], checkedTypeMap ) then
               local keyType = typeInfo:get_itemTypeInfoList()[1]
               if keyType:equals( _moduleObj.builtinTypeString ) or keyType:get_kind() == TypeInfoKind.Prim or keyType:get_kind() == TypeInfoKind.Enum then
                  return true
               end
               
            end
            
            return false
         elseif _switchExp == TypeInfoKind.Nilable then
            return NormalTypeInfo.isAvailableMapping( typeInfo:get_nonnilableType(), checkedTypeMap )
         else 
            
               return false
         end
      end
      
   end
   
   typeInfo = typeInfo:get_srcTypeInfo()
   do
      local _exp = checkedTypeMap[typeInfo]
      if _exp ~= nil then
         return _exp
      end
   end
   
   checkedTypeMap[typeInfo] = true
   local result = isAvailableMappingSub(  )
   checkedTypeMap[typeInfo] = result
   return result
end

function NormalTypeInfo:isInheritFrom( other, alt2type )

   if self:get_typeId() == other:get_typeId() then
      return true
   end
   
   if (self:get_kind() ~= TypeInfoKind.Class and self:get_kind() ~= TypeInfoKind.IF ) or (other:get_kind() ~= TypeInfoKind.Class and other:get_kind() ~= TypeInfoKind.IF ) then
      if other == _moduleObj.builtinTypeMapping then
         return NormalTypeInfo.isAvailableMapping( self, {} )
      end
      
      return false
   end
   
   return TypeInfo.isInherit( self, other, alt2type )
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

function TypeInfo.checkMatchType( dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex, alt2type )

   
   local function checkDstTypeFrom( index, srcType, srcType2nd )
   
      local workExpType = srcType
      for dstIndex = index, #dstTypeList do
         local workDstType = dstTypeList[dstIndex]
         local matchResult = MatchType.Match
         if not workDstType:canEvalWith( workExpType, "=", alt2type ) then
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
         
         if not dstType:canEvalWith( checkType, "=", alt2type ) then
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
               if not dstType:canEvalWith( expType, "=", alt2type ) then
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
          
            if not dstType:canEvalWith( expType, "=", alt2type ) then
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
         if not dstType:canEvalWith( _moduleObj.builtinTypeNil, "=", alt2type ) then
            return MatchType.Error, string.format( "exp(%d) type mismatch %s <- nil", index, dstType:getTxt( true ))
         end
         
         return MatchType.Warn, Code.format( Code.ID.nothing_define_abbr, string.format( "use '##', instate of %s.", dstType:getTxt( true )) )
      end
      
   end
   
   return MatchType.Match, ""
end

function TypeInfo.canEvalWithBase( dest, destMut, other, opTxt, alt2type )

   local otherMut = TypeInfo.isMut( other )
   local otherSrc = other:get_srcTypeInfo()
   if otherSrc:get_kind() == TypeInfoKind.DDD then
      if #otherSrc:get_itemTypeInfoList() > 0 then
         otherSrc = otherSrc:get_itemTypeInfoList()[1]
      end
      
   end
   
   if opTxt == "=" and otherSrc ~= _moduleObj.builtinTypeNil and otherSrc ~= _moduleObj.builtinTypeString and otherSrc:get_kind() ~= TypeInfoKind.Prim and otherSrc:get_kind() ~= TypeInfoKind.Func and otherSrc:get_kind() ~= TypeInfoKind.Enum and otherSrc:get_kind() ~= TypeInfoKind.Abbr and otherSrc:get_kind() ~= TypeInfoKind.Alternate and otherSrc:get_kind() ~= TypeInfoKind.Box and not isGenericType( otherSrc ) and destMut and not otherMut then
      return false
   end
   
   if dest == _moduleObj.builtinTypeStem_ then
      return true
   end
   
   if dest:get_srcTypeInfo():get_kind() == TypeInfoKind.DDD then
      if #dest:get_itemTypeInfoList() > 0 then
         return dest:get_itemTypeInfoList()[1]:canEvalWith( other, opTxt, alt2type )
      end
      
      return true
   end
   
   if not dest:get_nilable() and otherSrc:get_nilable() then
      if dest:get_kind() == TypeInfoKind.Box then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      end
      
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
            return dest:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt, alt2type )
         end
         
         return dest:get_nonnilableType():canEvalWith( otherSrc, opTxt, alt2type )
      elseif isGenericType( dest ) then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      elseif (dest:get_kind() == TypeInfoKind.Class or dest:get_kind() == TypeInfoKind.IF ) and (otherSrc:get_kind() == TypeInfoKind.Class or otherSrc:get_kind() == TypeInfoKind.IF ) then
         return otherSrc:isInheritFrom( dest, alt2type )
      elseif otherSrc:get_kind() == TypeInfoKind.Enum then
         local enumTypeInfo = _lune.unwrap( (_lune.__Cast( otherSrc, 3, EnumTypeInfo ) ))
         return dest:canEvalWith( enumTypeInfo:get_valTypeInfo(), opTxt, alt2type )
      elseif dest:get_kind() == TypeInfoKind.Alternate then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      elseif dest:get_kind() == TypeInfoKind.Box then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      end
      
      return false
   end
   
   do
      local _switchExp = dest:get_kind()
      if _switchExp == TypeInfoKind.Prim then
         if dest == _moduleObj.builtinTypeInt and otherSrc == _moduleObj.builtinTypeChar or dest == _moduleObj.builtinTypeChar and otherSrc == _moduleObj.builtinTypeInt then
            return true
         end
         
         return false
      elseif _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone then
            return true
         end
         
         if #dest:get_itemTypeInfoList() >= 1 and #otherSrc:get_itemTypeInfoList() >= 1 then
            if not (dest:get_itemTypeInfoList()[1] ):canEvalWith( otherSrc:get_itemTypeInfoList()[1], "=", alt2type ) then
               return false
            end
            
         else
          
            return false
         end
         
         
         return true
      elseif _switchExp == TypeInfoKind.Map then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone and otherSrc:get_itemTypeInfoList()[2] == _moduleObj.builtinTypeNone then
            return true
         end
         
         if #dest:get_itemTypeInfoList() >= 1 and #otherSrc:get_itemTypeInfoList() >= 1 then
            if not (dest:get_itemTypeInfoList()[1] ):canEvalWith( otherSrc:get_itemTypeInfoList()[1], "=", alt2type ) then
               return false
            end
            
         else
          
            return false
         end
         
         
         if #dest:get_itemTypeInfoList() >= 2 and #otherSrc:get_itemTypeInfoList() >= 2 then
            if not (dest:get_itemTypeInfoList()[2] ):canEvalWith( otherSrc:get_itemTypeInfoList()[2], "=", alt2type ) then
               return false
            end
            
         else
          
            return false
         end
         
         
         return true
      elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
         return otherSrc:isInheritFrom( dest, alt2type )
      elseif _switchExp == TypeInfoKind.Func then
         if dest == _moduleObj.builtinTypeForm then
            return true
         end
         
         if TypeInfo.checkMatchType( dest:get_argTypeInfoList(), otherSrc:get_argTypeInfoList(), false, nil, alt2type ) == MatchType.Error or TypeInfo.checkMatchType( dest:get_retTypeInfoList(), otherSrc:get_retTypeInfoList(), false, nil, alt2type ) == MatchType.Error or #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false
         end
         
         return true
      elseif _switchExp == TypeInfoKind.Method then
         if #dest:get_argTypeInfoList() ~= #otherSrc:get_argTypeInfoList() or #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false
         end
         
         for index, argType in pairs( dest:get_argTypeInfoList() ) do
            local otherArgType = otherSrc:get_argTypeInfoList()[index]
            if not argType:equals( otherArgType, alt2type ) then
               Util.errorLog( string.format( "unmatch arg(%d) type -- %s, %s", index, argType:getTxt(  ), otherArgType:getTxt(  )) )
               return false
            end
            
         end
         
         for index, retType in pairs( dest:get_retTypeInfoList() ) do
            local otherRetType = otherSrc:get_retTypeInfoList()[index]
            if not retType:equals( otherRetType, alt2type ) then
               Util.errorLog( string.format( "unmatch ret(%d) type -- %s, %s", index, retType:getTxt(  ), otherRetType:getTxt(  )) )
               return false
            end
            
         end
         
         return true
      elseif _switchExp == TypeInfoKind.Nilable then
         return dest:get_nonnilableType():canEvalWith( otherSrc:get_nonnilableType(), opTxt, alt2type )
      elseif _switchExp == TypeInfoKind.Alternate then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      elseif _switchExp == TypeInfoKind.Box then
         return dest:canEvalWith( otherSrc, opTxt, alt2type )
      else 
         
            return false
      end
   end
   
end

function NormalTypeInfo:canEvalWith( other, opTxt, alt2type )

   return TypeInfo.canEvalWithBase( self, TypeInfo.isMut( self ), other, opTxt, alt2type )
end

function ModifierTypeInfo:applyGeneric( alt2typeMap )

   local typeInfo = self.srcTypeInfo:applyGeneric( alt2typeMap )
   if typeInfo == self.srcTypeInfo then
      return self
   end
   
   if typeInfo ~= nil then
      return NormalTypeInfo.createModifier( typeInfo, MutMode.IMut )
   end
   
   return nil
end

function NormalTypeInfo:applyGeneric( alt2typeMap )

   local needNew = false
   local fail = false
   local function createGen( typeList )
   
      local typeInfoList = {}
      for __index, srcType in pairs( typeList ) do
         do
            local typeInfo = srcType:applyGeneric( alt2typeMap )
            if typeInfo ~= nil then
               table.insert( typeInfoList, typeInfo )
               if srcType ~= typeInfo then
                  needNew = true
               end
               
            else
               fail = true
               break
            end
         end
         
      end
      
      return typeInfoList
   end
   
   local itemTypeInfoList = createGen( self.itemTypeInfoList )
   local argTypeInfoList = createGen( self.argTypeInfoList )
   local retTypeInfoList = createGen( self.retTypeInfoList )
   if fail then
      return nil
   end
   
   if not needNew then
      return self
   end
   
   do
      local _switchExp = self:get_kind()
      if _switchExp == TypeInfoKind.Set then
         return NormalTypeInfo.createSet( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.List then
         return NormalTypeInfo.createList( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.Array then
         return NormalTypeInfo.createArray( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.Map then
         return NormalTypeInfo.createMap( self.accessMode, self.parentInfo, itemTypeInfoList[1], itemTypeInfoList[2], self.mutMode )
      elseif _switchExp == TypeInfoKind.Func then
         return NormalTypeInfo.createFunc( self.abstractFlag, false, getScope( self ), self.kind, self.parentInfo, self.autoFlag, self.externalFlag, self.staticFlag, self.accessMode, self.rawTxt, itemTypeInfoList, argTypeInfoList, retTypeInfoList, TypeInfo.isMut( self ) )
      else 
         
            return nil
      end
   end
   
end

function TypeInfo:getFullName( importInfo, localFlag )

   return self:getParentFullName( importInfo, localFlag ) .. self:get_rawTxt()
end

local ProcessInfo = {}
_moduleObj.ProcessInfo = ProcessInfo
function ProcessInfo.setmeta( obj )
  setmetatable( obj, { __index = ProcessInfo  } )
end
function ProcessInfo.new( idProvier, idProvierExt, typeInfo2Map )
   local obj = {}
   ProcessInfo.setmeta( obj )
   if obj.__init then
      obj:__init( idProvier, idProvierExt, typeInfo2Map )
   end        
   return obj 
end         
function ProcessInfo:__init( idProvier, idProvierExt, typeInfo2Map ) 

   self.idProvier = idProvier
   self.idProvierExt = idProvierExt
   self.typeInfo2Map = typeInfo2Map
end
function ProcessInfo:get_idProvier()       
   return self.idProvier         
end
function ProcessInfo:get_idProvierExt()       
   return self.idProvierExt         
end
function ProcessInfo:get_typeInfo2Map()       
   return self.typeInfo2Map         
end

local processInfoQueue = {}
local IdType = {}
_moduleObj.IdType = IdType
IdType._val2NameMap = {}
function IdType:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "IdType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function IdType._from( val )
   if IdType._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
IdType.__allList = {}
function IdType.get__allList()
   return IdType.__allList
end

IdType.Base = 0
IdType._val2NameMap[0] = 'Base'
IdType.__allList[1] = IdType.Base
IdType.Ext = 1
IdType._val2NameMap[1] = 'Ext'
IdType.__allList[2] = IdType.Ext

local function switchIdProvier( idType )

   if idType == IdType.Base then
      idProv = idProvBase
   else
    
      idProv = idProvExt
   end
   
end
_moduleObj.switchIdProvier = switchIdProvier
local builtinTypeInfo2Map = typeInfo2Map:clone(  )
local function pushProcessInfo( processInfo )

   if #processInfoQueue == 0 then
      if idProv:get_id() >= userStartId then
         Util.err( "builtinId is over" )
      end
      
   end
   
   table.insert( processInfoQueue, ProcessInfo.new(idProvBase, idProvExt, typeInfo2Map) )
   if processInfo ~= nil then
      idProvBase = processInfo:get_idProvier()
      idProvExt = processInfo:get_idProvierExt()
      typeInfo2Map = processInfo:get_typeInfo2Map()
   else
      idProvBase = IdProvider.new(userStartId, extStartId)
      idProvExt = IdProvider.new(extStartId, extMaxId)
      typeInfo2Map = builtinTypeInfo2Map:clone(  )
   end
   
   idProv = idProvBase
   return ProcessInfo.new(idProvBase, idProvExt, typeInfo2Map)
end
_moduleObj.pushProcessInfo = pushProcessInfo
local function popProcessInfo(  )

   local info = processInfoQueue[#processInfoQueue]
   idProvBase = info:get_idProvier()
   idProvExt = info:get_idProvierExt()
   idProv = idProvBase
   typeInfo2Map = info:get_typeInfo2Map()
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
