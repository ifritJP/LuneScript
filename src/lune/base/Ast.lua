--lune/base/Ast.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Ast'
local _lune = {}
if _lune3 then
   _lune = _lune3
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

function _lune.nilacc( val, fieldName, access, ... )
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
            if ifType == class then
               return true
            end
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

if not _lune3 then
   _lune3 = _lune
end
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Code = _lune.loadModule( 'lune.base.Code' )
local Log = _lune.loadModule( 'lune.base.Log' )
local Types = _lune.loadModule( 'lune.base.Types' )

local IdProvider = {}
_moduleObj.IdProvider = IdProvider
function IdProvider:increment(  )

   self.id = self.id + 1
   if self.id >= self.maxId then
      Util.err( "id is over" )
   end
   
end
function IdProvider:getNewId(  )

   self.id = self.id + 1
   if self.id >= self.maxId then
      Util.err( "id is over" )
   end
   
   return self.id
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

local TypeInfo2Map = {}

local extStartId = 100000
local extMaxId = 10000000
local userStartId = 1000

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

local ProcessInfo = {}
_moduleObj.ProcessInfo = ProcessInfo
function ProcessInfo:get_typeInfo2Map(  )

   return _lune.unwrap( self.typeInfo2Map)
end
function ProcessInfo:set_typeInfo2Map( typeInfo2Map )

   self.typeInfo2Map = typeInfo2Map
end
function ProcessInfo.new( validCheckingMutable, idProvBase, validExtType, typeInfo2Map )
   local obj = {}
   ProcessInfo.setmeta( obj )
   if obj.__init then obj:__init( validCheckingMutable, idProvBase, validExtType, typeInfo2Map ); end
   return obj
end
function ProcessInfo:__init(validCheckingMutable, idProvBase, validExtType, typeInfo2Map) 
   self.validCheckingMutable = validCheckingMutable
   self.validExtType = validExtType
   self.idProvBase = idProvBase
   self.idProvExt = IdProvider.new(extStartId, extMaxId)
   self.idProv = self.idProvBase
   self.idProvSym = IdProvider.new(0, extMaxId)
   self.idProvScope = IdProvider.new(0, extMaxId)
   self.typeInfo2Map = typeInfo2Map
end
function ProcessInfo.createRoot(  )

   return ProcessInfo.new(true, IdProvider.new(0, userStartId), true, nil)
end
function ProcessInfo.createUser( validCheckingMutable, validExtType, typeInfo2Map )

   return ProcessInfo.new(validCheckingMutable, IdProvider.new(userStartId, extStartId), validExtType, typeInfo2Map)
end
function ProcessInfo:switchIdProvier( idType )
   local __func__ = '@lune.@base.@Ast.ProcessInfo.switchIdProvier'

   Log.log( Log.Level.Trace, __func__, 112, function (  )
   
      return "start"
   end )
   
   if idType == IdType.Base then
      self.idProv = self.idProvBase
   else
    
      self.idProv = self.idProvExt
   end
   
end
function ProcessInfo.setmeta( obj )
  setmetatable( obj, { __index = ProcessInfo  } )
end
function ProcessInfo:get_idProvSym()
   return self.idProvSym
end
function ProcessInfo:get_idProvScope()
   return self.idProvScope
end
function ProcessInfo:get_idProv()
   return self.idProv
end
function ProcessInfo:get_idProvBase()
   return self.idProvBase
end
function ProcessInfo:get_idProvExt()
   return self.idProvExt
end
function ProcessInfo:get_validExtType()
   return self.validExtType
end
function ProcessInfo:get_validCheckingMutable()
   return self.validCheckingMutable
end


local IdInfo = {}
_moduleObj.IdInfo = IdInfo
function IdInfo:equals( idInfo )

   return (self.id == idInfo.id ) and (self.processInfo == idInfo.processInfo )
end
function IdInfo.setmeta( obj )
  setmetatable( obj, { __index = IdInfo  } )
end
function IdInfo.new( id, processInfo )
   local obj = {}
   IdInfo.setmeta( obj )
   if obj.__init then
      obj:__init( id, processInfo )
   end
   return obj
end
function IdInfo:__init( id, processInfo )

   self.id = id
   self.processInfo = processInfo
end


local rootProcessInfo = ProcessInfo.createRoot(  )
local function getRootProcessInfo(  )

   return rootProcessInfo
end
_moduleObj.getRootProcessInfo = getRootProcessInfo

function ProcessInfo:newId(  )

   local id = self.idProv:getNewId(  )
   return IdInfo.new(id, self)
end


local rootTypeIdInfo = rootProcessInfo:newId(  )
_moduleObj.rootTypeIdInfo = rootTypeIdInfo

local rootTypeId = _moduleObj.rootTypeIdInfo.id
_moduleObj.rootTypeId = rootTypeId


local TypeInfo = {}

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


local SimpleModuleInfo = {}
setmetatable( SimpleModuleInfo, { ifList = {ModuleInfoIF,} } )
_moduleObj.SimpleModuleInfo = SimpleModuleInfo
function SimpleModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = SimpleModuleInfo  } )
end
function SimpleModuleInfo.new( assignName, modulePath )
   local obj = {}
   SimpleModuleInfo.setmeta( obj )
   if obj.__init then
      obj:__init( assignName, modulePath )
   end
   return obj
end
function SimpleModuleInfo:__init( assignName, modulePath )

   self.assignName = assignName
   self.modulePath = modulePath
end
function SimpleModuleInfo:get_assignName()
   return self.assignName
end
function SimpleModuleInfo:get_modulePath()
   return self.modulePath
end


local ModuleInfoManager = {}
_moduleObj.ModuleInfoManager = ModuleInfoManager
function ModuleInfoManager.setmeta( obj )
  setmetatable( obj, { __index = ModuleInfoManager  } )
end
function ModuleInfoManager.new(  )
   local obj = {}
   ModuleInfoManager.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function ModuleInfoManager:__init(  )

end


local DummyModuleInfoManager = {}
setmetatable( DummyModuleInfoManager, { ifList = {ModuleInfoManager,} } )
_moduleObj.DummyModuleInfoManager = DummyModuleInfoManager
function DummyModuleInfoManager.new(  )
   local obj = {}
   DummyModuleInfoManager.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function DummyModuleInfoManager:__init() 
end
function DummyModuleInfoManager:getModuleInfo( typeInfo )

   return nil
end
function DummyModuleInfoManager.setmeta( obj )
  setmetatable( obj, { __index = DummyModuleInfoManager  } )
end
function DummyModuleInfoManager.get_instance()
   return DummyModuleInfoManager.instance
end
do
   DummyModuleInfoManager.instance = DummyModuleInfoManager.new()
end


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


local Scope = {}

local LowSymbol = {}
_moduleObj.LowSymbol = LowSymbol
function LowSymbol.setmeta( obj )
  setmetatable( obj, { __index = LowSymbol  } )
end
function LowSymbol.new(  )
   local obj = {}
   LowSymbol.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function LowSymbol:__init(  )

end


local SymbolInfo = {}

local sym2builtInTypeMap = {}
local function getSym2builtInTypeMap(  )

   return sym2builtInTypeMap
end
_moduleObj.getSym2builtInTypeMap = getSym2builtInTypeMap
local builtInTypeIdSet = {}
local function getBuiltInTypeIdMap(  )

   return builtInTypeIdSet
end
_moduleObj.getBuiltInTypeIdMap = getBuiltInTypeIdMap

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
SerializeKind.Ext = 11
SerializeKind._val2NameMap[11] = 'Ext'
SerializeKind.__allList[12] = SerializeKind.Ext


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
TypeInfoKind.Form = 22
TypeInfoKind._val2NameMap[22] = 'Form'
TypeInfoKind.__allList[23] = TypeInfoKind.Form
TypeInfoKind.FormFunc = 23
TypeInfoKind._val2NameMap[23] = 'FormFunc'
TypeInfoKind.__allList[24] = TypeInfoKind.FormFunc
TypeInfoKind.Ext = 24
TypeInfoKind._val2NameMap[24] = 'Ext'
TypeInfoKind.__allList[25] = TypeInfoKind.Ext
TypeInfoKind.CombineIF = 25
TypeInfoKind._val2NameMap[25] = 'CombineIF'
TypeInfoKind.__allList[26] = TypeInfoKind.CombineIF
TypeInfoKind.ExtModule = 26
TypeInfoKind._val2NameMap[26] = 'ExtModule'
TypeInfoKind.__allList[27] = TypeInfoKind.ExtModule


local function isBuiltin( typeId )

   return builtInTypeIdSet[typeId] ~= nil
end
_moduleObj.isBuiltin = isBuiltin

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

local accessMode2txtMap = {}
accessMode2txtMap[AccessMode.None] = "none"
accessMode2txtMap[AccessMode.Pub] = "pub"
accessMode2txtMap[AccessMode.Pro] = "pro"
accessMode2txtMap[AccessMode.Pri] = "pri"
accessMode2txtMap[AccessMode.Local] = "local"
accessMode2txtMap[AccessMode.Global] = "global"
local function accessMode2txt( accessMode )

   return _lune.unwrap( accessMode2txtMap[accessMode])
end
_moduleObj.accessMode2txt = accessMode2txt

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
MutMode.IMutRe = 1
MutMode._val2NameMap[1] = 'IMutRe'
MutMode.__allList[2] = MutMode.IMutRe
MutMode.Mut = 2
MutMode._val2NameMap[2] = 'Mut'
MutMode.__allList[3] = MutMode.Mut
MutMode.AllMut = 3
MutMode._val2NameMap[3] = 'AllMut'
MutMode.__allList[4] = MutMode.AllMut

local function isMutable( mode )

   do
      local _switchExp = mode
      if _switchExp == MutMode.AllMut or _switchExp == MutMode.Mut then
         return true
      end
   end
   
   return false
end
_moduleObj.isMutable = isMutable
local TypeNameCtrl = {}
_moduleObj.TypeNameCtrl = TypeNameCtrl
function TypeNameCtrl.setmeta( obj )
  setmetatable( obj, { __index = TypeNameCtrl  } )
end
function TypeNameCtrl.new( moduleTypeInfo )
   local obj = {}
   TypeNameCtrl.setmeta( obj )
   if obj.__init then
      obj:__init( moduleTypeInfo )
   end
   return obj
end
function TypeNameCtrl:__init( moduleTypeInfo )

   self.moduleTypeInfo = moduleTypeInfo
end
function TypeNameCtrl:get_moduleTypeInfo()
   return self.moduleTypeInfo
end
function TypeNameCtrl:set_moduleTypeInfo( moduleTypeInfo )
   self.moduleTypeInfo = moduleTypeInfo
end


local ScopeAccess = {}
_moduleObj.ScopeAccess = ScopeAccess
ScopeAccess._val2NameMap = {}
function ScopeAccess:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ScopeAccess.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function ScopeAccess._from( val )
   if ScopeAccess._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
ScopeAccess.__allList = {}
function ScopeAccess.get__allList()
   return ScopeAccess.__allList
end

ScopeAccess.Normal = 0
ScopeAccess._val2NameMap[0] = 'Normal'
ScopeAccess.__allList[1] = ScopeAccess.Normal
ScopeAccess.Full = 1
ScopeAccess._val2NameMap[1] = 'Full'
ScopeAccess.__allList[2] = ScopeAccess.Full

setmetatable( SymbolInfo, { ifList = {LowSymbol,} } )
_moduleObj.SymbolInfo = SymbolInfo
function SymbolInfo.new(  )
   local obj = {}
   SymbolInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function SymbolInfo:__init() 
   self.namespaceTypeInfo = nil
end
function SymbolInfo:hasAccess(  )

   if self:get_posForModToRef() or self:get_posForLatestMod() ~= self:get_pos() then
      return true
   end
   
   return false
end
function SymbolInfo:updateValue( pos )

   self:set_hasValueFlag( true )
   self:set_posForLatestMod( pos )
end
function SymbolInfo:clearValue(  )

   self:set_hasValueFlag( false )
end
function SymbolInfo.setmeta( obj )
  setmetatable( obj, { __index = SymbolInfo  } )
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


setmetatable( Scope, { ifList = {ModuleInfoManager,} } )
_moduleObj.Scope = Scope
function Scope.new( processInfo, parent, classFlag, inherit, ifScopeList )
   local obj = {}
   Scope.setmeta( obj )
   if obj.__init then obj:__init( processInfo, parent, classFlag, inherit, ifScopeList ); end
   return obj
end
function Scope:__init(processInfo, parent, classFlag, inherit, ifScopeList) 
   self.scopeId = processInfo:get_idProvScope():getNewId(  )
   self.typeInfo2ModuleInfoMap = {}
   self.closureSymMap = {}
   self.closureSym2NumMap = {}
   self.closureSymList = {}
   self.parent = _lune.unwrapDefault( parent, self)
   self.symbol2SymbolInfoMap = {}
   self.inherit = inherit
   self.classFlag = classFlag
   self.symbolId2DataOwnerInfo = {}
   self.ifScopeList = _lune.unwrapDefault( ifScopeList, {})
   self.ownerTypeInfo = nil
   self.validCheckingUnaccess = true
end
function Scope:isRoot(  )

   return self.parent == self
end
function Scope:set_ownerTypeInfo( owner )

   if not self.ownerTypeInfo then
      self.ownerTypeInfo = owner
   end
   
end
function Scope:switchOwnerTypeInfo( owner )

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
function Scope:addModule( typeInfo, moduleInfo )

   self.typeInfo2ModuleInfoMap[typeInfo] = moduleInfo
end
function Scope:getModuleInfo( typeInfo )

   do
      local moduleInfo = self.typeInfo2ModuleInfoMap[typeInfo]
      if moduleInfo ~= nil then
         return moduleInfo
      end
   end
   
   if self.parent ~= self then
      return self.parent:getModuleInfo( typeInfo )
   end
   
   return nil
end
function Scope.setmeta( obj )
  setmetatable( obj, { __index = Scope  } )
end
function Scope:get_scopeId()
   return self.scopeId
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
function Scope:get_inherit()
   return self.inherit
end
function Scope:get_closureSymMap()
   return self.closureSymMap
end
function Scope:get_closureSymList()
   return self.closureSymList
end
function Scope:get_closureSym2NumMap()
   return self.closureSym2NumMap
end
function Scope:get_validCheckingUnaccess()
   return self.validCheckingUnaccess
end
function Scope:set_validCheckingUnaccess( validCheckingUnaccess )
   self.validCheckingUnaccess = validCheckingUnaccess
end


function SymbolInfo:get_namespaceTypeInfo(  )

   do
      local _exp = self.namespaceTypeInfo
      if _exp ~= nil then
         return _exp
      end
   end
   
   local work = self:get_scope():getNamespaceTypeInfo(  )
   self.namespaceTypeInfo = work
   return work
end

local rootScope = Scope.new(rootProcessInfo, nil, false, nil)
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

local CanEvalType = {}
_moduleObj.CanEvalType = CanEvalType
CanEvalType._val2NameMap = {}
function CanEvalType:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "CanEvalType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function CanEvalType._from( val )
   if CanEvalType._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
CanEvalType.__allList = {}
function CanEvalType.get__allList()
   return CanEvalType.__allList
end

CanEvalType.SetOpIMut = 0
CanEvalType._val2NameMap[0] = 'SetOpIMut'
CanEvalType.__allList[1] = CanEvalType.SetOpIMut
CanEvalType.SetOp = 1
CanEvalType._val2NameMap[1] = 'SetOp'
CanEvalType.__allList[2] = CanEvalType.SetOp
CanEvalType.SetEq = 2
CanEvalType._val2NameMap[2] = 'SetEq'
CanEvalType.__allList[3] = CanEvalType.SetEq
CanEvalType.Equal = 3
CanEvalType._val2NameMap[3] = 'Equal'
CanEvalType.__allList[4] = CanEvalType.Equal
CanEvalType.Math = 4
CanEvalType._val2NameMap[4] = 'Math'
CanEvalType.__allList[5] = CanEvalType.Math
CanEvalType.Comp = 5
CanEvalType._val2NameMap[5] = 'Comp'
CanEvalType.__allList[6] = CanEvalType.Comp
CanEvalType.Logical = 6
CanEvalType._val2NameMap[6] = 'Logical'
CanEvalType.__allList[7] = CanEvalType.Logical


local SerializeInfo = {}
_moduleObj.SerializeInfo = SerializeInfo
function SerializeInfo:isValidChildren( idInfo )

   return not self.validChildrenMap[idInfo]
end
function SerializeInfo:serializeId( idInfo )

   local processId = _lune.unwrap( self.processInfo2Id[idInfo.processInfo])
   return string.format( "{ id = %d, mod = %d }", idInfo.id, processId)
end
function SerializeInfo.setmeta( obj )
  setmetatable( obj, { __index = SerializeInfo  } )
end
function SerializeInfo.new( processInfo2Id, validChildrenMap )
   local obj = {}
   SerializeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( processInfo2Id, validChildrenMap )
   end
   return obj
end
function SerializeInfo:__init( processInfo2Id, validChildrenMap )

   self.processInfo2Id = processInfo2Id
   self.validChildrenMap = validChildrenMap
end
function SerializeInfo:get_validChildrenMap()
   return self.validChildrenMap
end


_moduleObj.TypeInfo = TypeInfo
function TypeInfo:getProcessInfo(  )

   return self.processInfo
end
function TypeInfo:switchScope( scope )

   self.scope = scope
   scope:switchOwnerTypeInfo( self )
end
function TypeInfo:getOverridingType(  )

   return nil
end
function TypeInfo.new( scope, processInfo )
   local obj = {}
   TypeInfo.setmeta( obj )
   if obj.__init then obj:__init( scope, processInfo ); end
   return obj
end
function TypeInfo:__init(scope, processInfo) 
   self.scope = scope
   do
      local _exp = scope
      if _exp ~= nil then
         _exp:set_ownerTypeInfo( self )
      end
   end
   
   self.typeData = TypeData.new({})
   self.processInfo = processInfo
   
end
function TypeInfo:get_aliasSrc(  )

   return self
end
function TypeInfo:get_extedType(  )

   return self
end
function TypeInfo.getModulePath( fullname )

   return (fullname:gsub( "@", "" ) )
end
function TypeInfo:isModule(  )

   return true
end
function TypeInfo:getParentId(  )

   return _moduleObj.rootTypeIdInfo
end
function TypeInfo:get_baseId(  )

   return _moduleObj.rootTypeIdInfo
end
function TypeInfo:isInheritFrom( processInfo, other, alt2type )

   return false
end
function TypeInfo:get_rawTxt(  )

   return ""
end
function TypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   return ""
end
function TypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function TypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return false, nil
end
function TypeInfo:get_abstractFlag(  )

   return false
end
function TypeInfo:serialize( stream, serializeInfo )

   return 
end
function TypeInfo:get_display_stirng_with( raw, alt2type )

   return ""
end
function TypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( "", nil )
end
function TypeInfo:get_srcTypeInfo(  )

   return self
end
function TypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   if checkModifer then
      return self == typeInfo
   end
   
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

   return _moduleObj.rootTypeIdInfo
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

   return isMutable( typeInfo:get_mutMode() )
end
function TypeInfo:getParentFullName( typeNameCtrl, importInfo, localFlag )

   return typeNameCtrl:getParentFullName( self, importInfo, localFlag )
end
function TypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   return self
end
function TypeInfo:get_genSrcTypeInfo(  )

   return self
end
function TypeInfo:serializeTypeInfoList( name, list, onlyPub )

   local work = name
   for __index, typeInfo in ipairs( list ) do
      if not onlyPub or typeInfo:get_accessMode() == AccessMode.Pub then
         if #work ~= #name then
            work = work .. ", "
         end
         
         work = string.format( "%s%d", work, typeInfo:get_typeId().id)
      end
      
   end
   
   return work .. "}, "
end
function TypeInfo.createScope( processInfo, parent, classFlag, baseInfo, interfaceList )

   local inheritScope = nil
   if baseInfo ~= nil then
      inheritScope = _lune.unwrap( baseInfo:get_scope())
   end
   
   local ifScopeList = {}
   if interfaceList ~= nil then
      for __index, ifType in ipairs( interfaceList ) do
         table.insert( ifScopeList, _lune.unwrap( ifType:get_scope()) )
      end
      
   end
   
   return Scope.new(processInfo, parent, classFlag, inheritScope, ifScopeList)
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
function TypeInfo:get_processInfo()
   return self.processInfo
end


local LuavalResult = {}
LuavalResult._name2Val = {}
_moduleObj.LuavalResult = LuavalResult
function LuavalResult:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "LuavalResult.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function LuavalResult._from( val )
   return _lune._AlgeFrom( LuavalResult, val )
end

LuavalResult.Err = { "Err", {{}}}
LuavalResult._name2Val["Err"] = LuavalResult.Err
LuavalResult.OK = { "OK", {{},{}}}
LuavalResult._name2Val["OK"] = LuavalResult.OK



local function applyGenericDefault( typeInfo, alt2typeMap, moduleTypeInfo )

   do
      local genType = typeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )
      if genType ~= nil then
         return genType
      end
   end
   
   return typeInfo
end
_moduleObj.applyGenericDefault = applyGenericDefault

function SymbolInfo:getModule(  )

   return self:get_namespaceTypeInfo(  ):getModule(  )
end


local MethodKind = {}
_moduleObj.MethodKind = MethodKind
MethodKind._val2NameMap = {}
function MethodKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "MethodKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function MethodKind._from( val )
   if MethodKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
MethodKind.__allList = {}
function MethodKind.get__allList()
   return MethodKind.__allList
end

MethodKind.All = 0
MethodKind._val2NameMap[0] = 'All'
MethodKind.__allList[1] = MethodKind.All
MethodKind.Static = 1
MethodKind._val2NameMap[1] = 'Static'
MethodKind.__allList[2] = MethodKind.Static
MethodKind.Object = 2
MethodKind._val2NameMap[2] = 'Object'
MethodKind.__allList[3] = MethodKind.Object

local function getAllNameForKind( classInfo, kind, symbolKind )

   local nameSet = Util.OrderedSet.new()
   local function process( scope )
   
      do
         local inherit = scope:get_inherit()
         if inherit ~= nil then
            process( inherit )
         end
      end
      
      do
         local __sorted = {}
         local __map = scope:get_symbol2SymbolInfoMap()
         for __key in pairs( __map ) do
            table.insert( __sorted, __key )
         end
         table.sort( __sorted )
         for __index, __key in ipairs( __sorted ) do
            local symbolInfo = __map[ __key ]
            do
               do
                  local _switchExp = symbolInfo:get_kind()
                  if _switchExp == symbolKind then
                     if symbolKind == SymbolKind.Mtd and symbolInfo:get_name() == "__init" then
                        
                     else
                      
                        local staticFlag = symbolInfo:get_staticFlag()
                        if kind == MethodKind.All or kind == MethodKind.Static and staticFlag or kind == MethodKind.Object and not staticFlag then
                           nameSet:add( symbolInfo:get_name() )
                        end
                        
                     end
                     
                  end
               end
               
            end
         end
      end
      
   end
   
   do
      local scope = classInfo:get_scope()
      if scope ~= nil then
         process( scope )
      end
   end
   
   return nameSet
end
_moduleObj.getAllNameForKind = getAllNameForKind

local function getAllMethodName( classInfo, kind )

   return getAllNameForKind( classInfo, kind, SymbolKind.Mtd )
end
_moduleObj.getAllMethodName = getAllMethodName

function TypeNameCtrl:getModuleName( workTypeInfo, name, moduleInfoMan )

   do
      local moduleInfo = moduleInfoMan:getModuleInfo( workTypeInfo )
      if moduleInfo ~= nil then
         local txt = moduleInfo:get_assignName()
         return txt .. "." .. name
      else
         if self.moduleTypeInfo ~= workTypeInfo then
            return workTypeInfo:get_rawTxt() .. "." .. name
         end
         
      end
   end
   
   return name
end


function TypeNameCtrl:getParentFullName( typeInfo, importInfo, localFlag )

   local workTypeInfo = typeInfo
   local name = ""
   local moduleInfoMan = importInfo
   if  nil == moduleInfoMan then
      local _moduleInfoMan = moduleInfoMan
   
      moduleInfoMan = DummyModuleInfoManager.get_instance()
   end
   
   while true do
      workTypeInfo = workTypeInfo:get_parentInfo()
      local txt = workTypeInfo:get_rawTxt()
      if workTypeInfo == workTypeInfo:get_parentInfo() then
         break
      end
      
      if localFlag then
         if workTypeInfo:isModule(  ) then
            name = self:getModuleName( workTypeInfo, name, moduleInfoMan )
            break
         end
         
      end
      
      name = txt .. "." .. name
   end
   
   return name
end


local function getScope( typeInfo )

   return typeInfo:get_scope()
   
end
_moduleObj.getScope = getScope

local function isExtId( typeInfo )

   if typeInfo:get_typeId().id >= extStartId then
      return true
   end
   
   return false
end
_moduleObj.isExtId = isExtId

local RootTypeInfo = {}
setmetatable( RootTypeInfo, { __index = TypeInfo } )
_moduleObj.RootTypeInfo = RootTypeInfo
function RootTypeInfo.new(  )
   local obj = {}
   RootTypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function RootTypeInfo:__init() 
   TypeInfo.__init( self,_moduleObj.rootScope, rootProcessInfo)
   
   self.typeId = _moduleObj.rootTypeIdInfo
end
function RootTypeInfo:get_baseTypeInfo(  )

   return self
end
function RootTypeInfo:get_parentInfo(  )

   return self
end
function RootTypeInfo.create(  )

   return RootTypeInfo.new()
end
function RootTypeInfo:get_rawTxt(  )

   return "<head>"
end
function RootTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return "<head>"
end
function RootTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = RootTypeInfo  } )
end
function RootTypeInfo:get_typeId()
   return self.typeId
end


local headTypeInfo = RootTypeInfo.create(  )
_moduleObj.headTypeInfo = headTypeInfo

local defaultTypeNameCtrl = TypeNameCtrl.new(_moduleObj.headTypeInfo)
_moduleObj.defaultTypeNameCtrl = defaultTypeNameCtrl


function TypeInfo:hasBase(  )

   return self:get_baseTypeInfo() ~= _moduleObj.headTypeInfo
end


function Scope:getNamespaceTypeInfo(  )

   local typeInfo = _moduleObj.headTypeInfo
   local scope = self
   repeat 
      do
         local _exp = scope:get_ownerTypeInfo()
         if _exp ~= nil then
            return _exp
         end
      end
      
      scope = scope:get_parent()
   until scope:isRoot(  )
   return typeInfo
end


function Scope:getModule(  )

   return self:getNamespaceTypeInfo(  ):getModule(  )
end


function Scope:getProcessInfo(  )

   return self:getModule(  ):getProcessInfo(  )
end


local NormalSymbolInfo = {}
setmetatable( NormalSymbolInfo, { __index = SymbolInfo } )
_moduleObj.NormalSymbolInfo = NormalSymbolInfo
function NormalSymbolInfo:get_mutable(  )

   return isMutable( self.mutMode )
end
function NormalSymbolInfo:getOrg(  )

   return self
end
function NormalSymbolInfo.new( processInfo, kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad )
   local obj = {}
   NormalSymbolInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad ); end
   return obj
end
function NormalSymbolInfo:__init(processInfo, kind, canBeLeft, canBeRight, scope, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad) 
   SymbolInfo.__init( self)
   
   self.convModuleParam = nil
   self.hasAccessFromClosure = false
   if hasValueFlag then
      self.posForLatestMod = pos
   else
    
      self.posForLatestMod = nil
   end
   
   self.posForModToRef = nil
   self.kind = kind
   self.canBeLeft = canBeLeft
   self.canBeRight = canBeRight
   self.symbolId = processInfo:get_idProvSym():getNewId(  )
   self.scope = scope
   self.accessMode = accessMode
   self.staticFlag = staticFlag
   self.name = name
   self.pos = pos
   self.typeInfo = typeInfo
   self.mutMode = _lune.unwrapDefault( mutMode, MutMode.IMut)
   self.hasValueFlag = hasValueFlag
   self.isLazyLoad = isLazyLoad
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
function NormalSymbolInfo:get_isLazyLoad()
   return self.isLazyLoad
end
function NormalSymbolInfo:get_name()
   return self.name
end
function NormalSymbolInfo:get_pos()
   return self.pos
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
function NormalSymbolInfo:get_hasAccessFromClosure()
   return self.hasAccessFromClosure
end
function NormalSymbolInfo:set_hasAccessFromClosure( hasAccessFromClosure )
   self.hasAccessFromClosure = hasAccessFromClosure
end
function NormalSymbolInfo:get_posForLatestMod()
   return self.posForLatestMod
end
function NormalSymbolInfo:set_posForLatestMod( posForLatestMod )
   self.posForLatestMod = posForLatestMod
end
function NormalSymbolInfo:get_posForModToRef()
   return self.posForModToRef
end
function NormalSymbolInfo:set_posForModToRef( posForModToRef )
   self.posForModToRef = posForModToRef
end
function NormalSymbolInfo:get_convModuleParam()
   return self.convModuleParam
end
function NormalSymbolInfo:set_convModuleParam( convModuleParam )
   self.convModuleParam = convModuleParam
end


function TypeInfo.isInherit( processInfo, typeInfo, other, alt2type )

   local baseTypeInfo = typeInfo:get_baseTypeInfo()
   local otherTypeId = other:get_typeId()
   if typeInfo:get_typeId() == otherTypeId then
      return true
   end
   
   if baseTypeInfo ~= _moduleObj.headTypeInfo then
      if baseTypeInfo:isInheritFrom( processInfo, other, alt2type ) then
         return true
      end
      
   end
   
   
   for __index, ifType in ipairs( typeInfo:get_interfaceList() ) do
      if ifType:isInheritFrom( processInfo, other, alt2type ) then
         return true
      end
      
   end
   
   return false
end


local AutoBoxingInfo = {}
setmetatable( AutoBoxingInfo, { __index = TypeInfo } )
_moduleObj.AutoBoxingInfo = AutoBoxingInfo
function AutoBoxingInfo.new( processInfo )
   local obj = {}
   AutoBoxingInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo ); end
   return obj
end
function AutoBoxingInfo:__init(processInfo) 
   TypeInfo.__init( self,nil, processInfo)
   
   self.count = 0
   AutoBoxingInfo.allObj[self] = self
end
function AutoBoxingInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function AutoBoxingInfo:get_parentInfo(  )

   return _moduleObj.headTypeInfo
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


local dummyIdInfo = IdInfo.new(1, rootProcessInfo)
_moduleObj.dummyIdInfo = dummyIdInfo

local CanEvalCtrlTypeInfo = {}
setmetatable( CanEvalCtrlTypeInfo, { __index = TypeInfo } )
_moduleObj.CanEvalCtrlTypeInfo = CanEvalCtrlTypeInfo
function CanEvalCtrlTypeInfo.new(  )
   local obj = {}
   CanEvalCtrlTypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function CanEvalCtrlTypeInfo:__init() 
   TypeInfo.__init( self,nil, rootProcessInfo)
   
end
function CanEvalCtrlTypeInfo:get_kind(  )

   return TypeInfoKind.CanEvalCtrl
end
function CanEvalCtrlTypeInfo:get_typeId(  )

   return _moduleObj.dummyIdInfo
   
end
function CanEvalCtrlTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function CanEvalCtrlTypeInfo:get_parentInfo(  )

   return _moduleObj.headTypeInfo
end
function CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( detectFlag )

   if detectFlag then
      local map = {}
      map[CanEvalCtrlTypeInfo.detectAlt] = _moduleObj.headTypeInfo
      return map
   end
   
   return {}
end
function CanEvalCtrlTypeInfo.isValidApply( alt2type )

   return alt2type[CanEvalCtrlTypeInfo.detectAlt] ~= nil
end
function CanEvalCtrlTypeInfo.setupNeedAutoBoxing( alt2type, processInfo )

   alt2type[CanEvalCtrlTypeInfo.needAutoBoxing] = AutoBoxingInfo.new(processInfo)
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
function CanEvalCtrlTypeInfo.hasNeedAutoBoxing( alt2type )

   do
      local _exp = alt2type[CanEvalCtrlTypeInfo.needAutoBoxing]
      if _exp ~= nil then
         do
            local autoBoxingInfo = _lune.__Cast( _exp, 3, AutoBoxingInfo )
            if autoBoxingInfo ~= nil then
               return autoBoxingInfo:get_count() ~= 0
            end
         end
         
      end
   end
   
   return false
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
do
   CanEvalCtrlTypeInfo.detectAlt = CanEvalCtrlTypeInfo.new()
   CanEvalCtrlTypeInfo.needAutoBoxing = CanEvalCtrlTypeInfo.new()
   CanEvalCtrlTypeInfo.checkTypeTarget = CanEvalCtrlTypeInfo.new()
end


local NilableTypeInfo = {}
setmetatable( NilableTypeInfo, { __index = TypeInfo } )
_moduleObj.NilableTypeInfo = NilableTypeInfo
function NilableTypeInfo:get_kind(  )

   return TypeInfoKind.Nilable
end
function NilableTypeInfo:get_aliasSrc(  )

   return self
end
function NilableTypeInfo:get_srcTypeInfo(  )

   return self
end
function NilableTypeInfo:get_nilable(  )

   return true
end
function NilableTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function NilableTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   return self.nonnilableType:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag ) .. "!"
end
function NilableTypeInfo:get_display_stirng_with( raw, alt2type )

   if self.nonnilableType:get_kind() == TypeInfoKind.FormFunc then
      return self.nonnilableType:get_display_stirng_with( raw .. "!", alt2type )
   end
   
   return self.nonnilableType:get_display_stirng_with( raw, alt2type ) .. "!"
end
function NilableTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function NilableTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind = %d, typeId = %d, orgTypeId = %d }\n', SerializeKind.Nilable, self.typeId.id, self.nonnilableType:get_typeId().id) )
end
function NilableTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   if not typeInfo:get_nilable() then
      return false
   end
   
   return self.nonnilableType:equals( processInfo, typeInfo:get_nonnilableType(), alt2type, checkModifer )
end
function NilableTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local typeInfo = self.nonnilableType:applyGeneric( alt2typeMap, moduleTypeInfo )
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
function NilableTypeInfo.new( __superarg1, __superarg2,nonnilableType, typeId )
   local obj = {}
   NilableTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2,nonnilableType, typeId )
   end
   return obj
end
function NilableTypeInfo:__init( __superarg1, __superarg2,nonnilableType, typeId )

   TypeInfo.__init( self, __superarg1, __superarg2 )
   self.nonnilableType = nonnilableType
   self.typeId = typeId
end
function NilableTypeInfo:get_nonnilableType()
   return self.nonnilableType
end
function NilableTypeInfo:get_typeId()
   return self.typeId
end
function NilableTypeInfo:addChildren( ... )
   return self.nonnilableType:addChildren( ... )
end

function NilableTypeInfo:createAlt2typeMap( ... )
   return self.nonnilableType:createAlt2typeMap( ... )
end

function NilableTypeInfo:getFullName( ... )
   return self.nonnilableType:getFullName( ... )
end

function NilableTypeInfo:getModule( ... )
   return self.nonnilableType:getModule( ... )
end

function NilableTypeInfo:getOverridingType( ... )
   return self.nonnilableType:getOverridingType( ... )
end

function NilableTypeInfo:getParentFullName( ... )
   return self.nonnilableType:getParentFullName( ... )
end

function NilableTypeInfo:getParentId( ... )
   return self.nonnilableType:getParentId( ... )
end

function NilableTypeInfo:getProcessInfo( ... )
   return self.nonnilableType:getProcessInfo( ... )
end

function NilableTypeInfo:get_abstractFlag( ... )
   return self.nonnilableType:get_abstractFlag( ... )
end

function NilableTypeInfo:get_accessMode( ... )
   return self.nonnilableType:get_accessMode( ... )
end

function NilableTypeInfo:get_argTypeInfoList( ... )
   return self.nonnilableType:get_argTypeInfoList( ... )
end

function NilableTypeInfo:get_autoFlag( ... )
   return self.nonnilableType:get_autoFlag( ... )
end

function NilableTypeInfo:get_baseId( ... )
   return self.nonnilableType:get_baseId( ... )
end

function NilableTypeInfo:get_baseTypeInfo( ... )
   return self.nonnilableType:get_baseTypeInfo( ... )
end

function NilableTypeInfo:get_children( ... )
   return self.nonnilableType:get_children( ... )
end

function NilableTypeInfo:get_extedType( ... )
   return self.nonnilableType:get_extedType( ... )
end

function NilableTypeInfo:get_externalFlag( ... )
   return self.nonnilableType:get_externalFlag( ... )
end

function NilableTypeInfo:get_genSrcTypeInfo( ... )
   return self.nonnilableType:get_genSrcTypeInfo( ... )
end

function NilableTypeInfo:get_interfaceList( ... )
   return self.nonnilableType:get_interfaceList( ... )
end

function NilableTypeInfo:get_itemTypeInfoList( ... )
   return self.nonnilableType:get_itemTypeInfoList( ... )
end

function NilableTypeInfo:get_mutMode( ... )
   return self.nonnilableType:get_mutMode( ... )
end

function NilableTypeInfo:get_nilableTypeInfo( ... )
   return self.nonnilableType:get_nilableTypeInfo( ... )
end

function NilableTypeInfo:get_parentInfo( ... )
   return self.nonnilableType:get_parentInfo( ... )
end

function NilableTypeInfo:get_processInfo( ... )
   return self.nonnilableType:get_processInfo( ... )
end

function NilableTypeInfo:get_rawTxt( ... )
   return self.nonnilableType:get_rawTxt( ... )
end

function NilableTypeInfo:get_retTypeInfoList( ... )
   return self.nonnilableType:get_retTypeInfoList( ... )
end

function NilableTypeInfo:get_scope( ... )
   return self.nonnilableType:get_scope( ... )
end

function NilableTypeInfo:get_staticFlag( ... )
   return self.nonnilableType:get_staticFlag( ... )
end

function NilableTypeInfo:get_typeData( ... )
   return self.nonnilableType:get_typeData( ... )
end

function NilableTypeInfo:hasBase( ... )
   return self.nonnilableType:hasBase( ... )
end

function NilableTypeInfo:hasRouteNamespaceFrom( ... )
   return self.nonnilableType:hasRouteNamespaceFrom( ... )
end

function NilableTypeInfo:isInheritFrom( ... )
   return self.nonnilableType:isInheritFrom( ... )
end

function NilableTypeInfo:isModule( ... )
   return self.nonnilableType:isModule( ... )
end

function NilableTypeInfo:serializeTypeInfoList( ... )
   return self.nonnilableType:serializeTypeInfoList( ... )
end

function NilableTypeInfo:switchScope( ... )
   return self.nonnilableType:switchScope( ... )
end



local AliasTypeInfo = {}
setmetatable( AliasTypeInfo, { __index = TypeInfo } )
_moduleObj.AliasTypeInfo = AliasTypeInfo
function AliasTypeInfo.new( processInfo, rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag )
   local obj = {}
   AliasTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag ); end
   return obj
end
function AliasTypeInfo:__init(processInfo, rawTxt, accessMode, parentInfo, aliasSrcTypeInfo, externalFlag) 
   TypeInfo.__init( self,nil, processInfo)
   
   self.rawTxt = rawTxt
   self.accessMode = accessMode
   self.parentInfo = parentInfo
   self.aliasSrcTypeInfo = aliasSrcTypeInfo
   self.externalFlag = externalFlag
   self.typeId = processInfo:newId(  )
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
end
function AliasTypeInfo:getParentFullName( typeNameCtrl, importInfo, localFlag )

   return typeNameCtrl:getParentFullName( self, importInfo, localFlag )
end
function AliasTypeInfo:getFullName( typeNameCtrl, importInfo, localFlag )

   return self:getParentFullName( typeNameCtrl, importInfo, localFlag ) .. self:get_rawTxt()
end
function AliasTypeInfo:get_aliasSrc(  )

   return self.aliasSrcTypeInfo
end
function AliasTypeInfo:get_nonnilableType(  )

   return self
end
function AliasTypeInfo:get_srcTypeInfo(  )

   return self
end
function AliasTypeInfo:get_genSrcTypeInfo(  )

   return self
end
function AliasTypeInfo:getModule(  )

   return self:get_parentInfo():getModule(  )
end
function AliasTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self.rawTxt, typeNameCtrl, importInfo, localFlag )
end
function AliasTypeInfo:serialize( stream, serializeInfo )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, rawTxt = %q, srcTypeId = %d }\n', SerializeKind.Alias, parentId.id, self.typeId.id, self.rawTxt, self.aliasSrcTypeInfo:get_typeId().id) )
end
function AliasTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self.rawTxt, nil )
end
function AliasTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function AliasTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local typeInfo = self.aliasSrcTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )
   if typeInfo == self.aliasSrcTypeInfo then
      return self
   end
   
   return nil
end
function AliasTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return self.aliasSrcTypeInfo:canEvalWith( processInfo, other:get_aliasSrc(), canEvalType, alt2type )
end
function AliasTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   return self.aliasSrcTypeInfo:equals( processInfo, typeInfo:get_aliasSrc(), alt2type, checkModifer )
end
function AliasTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AliasTypeInfo  } )
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
function AliasTypeInfo:get_nilableTypeInfo()
   return self.nilableTypeInfo
end
function AliasTypeInfo:addChildren( ... )
   return self.aliasSrcTypeInfo:addChildren( ... )
end

function AliasTypeInfo:createAlt2typeMap( ... )
   return self.aliasSrcTypeInfo:createAlt2typeMap( ... )
end

function AliasTypeInfo:getOverridingType( ... )
   return self.aliasSrcTypeInfo:getOverridingType( ... )
end

function AliasTypeInfo:getProcessInfo( ... )
   return self.aliasSrcTypeInfo:getProcessInfo( ... )
end

function AliasTypeInfo:getTxtWithRaw( ... )
   return self.aliasSrcTypeInfo:getTxtWithRaw( ... )
end

function AliasTypeInfo:get_abstractFlag( ... )
   return self.aliasSrcTypeInfo:get_abstractFlag( ... )
end

function AliasTypeInfo:get_argTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_argTypeInfoList( ... )
end

function AliasTypeInfo:get_autoFlag( ... )
   return self.aliasSrcTypeInfo:get_autoFlag( ... )
end

function AliasTypeInfo:get_baseId( ... )
   return self.aliasSrcTypeInfo:get_baseId( ... )
end

function AliasTypeInfo:get_baseTypeInfo( ... )
   return self.aliasSrcTypeInfo:get_baseTypeInfo( ... )
end

function AliasTypeInfo:get_children( ... )
   return self.aliasSrcTypeInfo:get_children( ... )
end

function AliasTypeInfo:get_display_stirng_with( ... )
   return self.aliasSrcTypeInfo:get_display_stirng_with( ... )
end

function AliasTypeInfo:get_extedType( ... )
   return self.aliasSrcTypeInfo:get_extedType( ... )
end

function AliasTypeInfo:get_interfaceList( ... )
   return self.aliasSrcTypeInfo:get_interfaceList( ... )
end

function AliasTypeInfo:get_itemTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_itemTypeInfoList( ... )
end

function AliasTypeInfo:get_kind( ... )
   return self.aliasSrcTypeInfo:get_kind( ... )
end

function AliasTypeInfo:get_mutMode( ... )
   return self.aliasSrcTypeInfo:get_mutMode( ... )
end

function AliasTypeInfo:get_nilable( ... )
   return self.aliasSrcTypeInfo:get_nilable( ... )
end

function AliasTypeInfo:get_processInfo( ... )
   return self.aliasSrcTypeInfo:get_processInfo( ... )
end

function AliasTypeInfo:get_retTypeInfoList( ... )
   return self.aliasSrcTypeInfo:get_retTypeInfoList( ... )
end

function AliasTypeInfo:get_scope( ... )
   return self.aliasSrcTypeInfo:get_scope( ... )
end

function AliasTypeInfo:get_staticFlag( ... )
   return self.aliasSrcTypeInfo:get_staticFlag( ... )
end

function AliasTypeInfo:get_typeData( ... )
   return self.aliasSrcTypeInfo:get_typeData( ... )
end

function AliasTypeInfo:hasBase( ... )
   return self.aliasSrcTypeInfo:hasBase( ... )
end

function AliasTypeInfo:hasRouteNamespaceFrom( ... )
   return self.aliasSrcTypeInfo:hasRouteNamespaceFrom( ... )
end

function AliasTypeInfo:isInheritFrom( ... )
   return self.aliasSrcTypeInfo:isInheritFrom( ... )
end

function AliasTypeInfo:isModule( ... )
   return self.aliasSrcTypeInfo:isModule( ... )
end

function AliasTypeInfo:serializeTypeInfoList( ... )
   return self.aliasSrcTypeInfo:serializeTypeInfoList( ... )
end

function AliasTypeInfo:switchScope( ... )
   return self.aliasSrcTypeInfo:switchScope( ... )
end




function Scope:filterTypeInfoField( includeSelfFlag, fromScope, access, callback )

   if self.classFlag then
      if includeSelfFlag then
         do
            local __sorted = {}
            local __map = self.symbol2SymbolInfoMap
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local symbolInfo = __map[ __key ]
               do
                  if symbolInfo:canAccess( fromScope, access ) then
                     if not callback( symbolInfo ) then
                        return false
                     end
                     
                  end
                  
               end
            end
         end
         
      end
      
      do
         local scope = self.inherit
         if scope ~= nil then
            if not scope:filterTypeInfoField( true, fromScope, access, callback ) then
               return false
            end
            
         end
      end
      
   end
   
   
   return true
end

function Scope:getSymbolInfoField( name, includeSelfFlag, fromScope, access )

   if self.classFlag then
      if includeSelfFlag then
         do
            local _exp = self.symbol2SymbolInfoMap[name]
            if _exp ~= nil then
               local symbolInfo = _exp:canAccess( fromScope, access )
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
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope, access )
            if symbolInfo then
               return symbolInfo
            end
            
         end
      end
      
   end
   
   
   return nil
end

function Scope:getSymbolInfoIfField( name, fromScope, access )

   if self.classFlag then
      for __index, scope in ipairs( self.ifScopeList ) do
         do
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope, access )
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
            local symbolInfo = scope:getSymbolInfoIfField( name, fromScope, access )
            if symbolInfo ~= nil then
               return symbolInfo
            end
         end
         
      end
   end
   
   
   return nil
end

function Scope:filterSymbolInfoIfField( fromScope, access, callback )

   for __index, scope in ipairs( self.ifScopeList ) do
      if not scope:filterTypeInfoField( true, fromScope, access, callback ) then
         return false
      end
      
   end
   
   
   do
      local scope = self.inherit
      if scope ~= nil then
         if not scope:filterSymbolInfoIfField( fromScope, access, callback ) then
            return false
         end
         
      end
   end
   
   
   return true
end


function Scope:getTypeInfoField( name, includeSelfFlag, fromScope, access )

   local symbolInfo = self:getSymbolInfoField( name, includeSelfFlag, fromScope, access )
   do
      local _exp = symbolInfo
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   return nil
end


function Scope:filterTypeInfoFieldAndIF( includeSelfFlag, fromScope, access, callback )

   if self.classFlag then
      if includeSelfFlag then
         do
            local __sorted = {}
            local __map = self.symbol2SymbolInfoMap
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local symbolInfo = __map[ __key ]
               do
                  if symbolInfo:canAccess( fromScope, access ) then
                     if not callback( symbolInfo ) then
                        return false
                     end
                     
                  end
                  
               end
            end
         end
         
      end
      
      do
         local scope = self.inherit
         if scope ~= nil then
            if not scope:filterTypeInfoField( true, fromScope, access, callback ) then
               return false
            end
            
         end
      end
      
   end
   
   
   for __index, scope in ipairs( self.ifScopeList ) do
      if not scope:filterTypeInfoField( true, fromScope, access, callback ) then
         return false
      end
      
   end
   
   
   do
      local scope = self.inherit
      if scope ~= nil then
         if not scope:filterSymbolInfoIfField( fromScope, access, callback ) then
            return false
         end
         
      end
   end
   
   
   return true
end


function Scope:getSymbolInfo( name, fromScope, onlySameNsFlag, access )

   do
      local _exp = self.symbol2SymbolInfoMap[name]
      if _exp ~= nil then
         local symbolInfo = _exp:canAccess( fromScope, access )
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
            local symbolInfo = scope:getSymbolInfoField( name, true, fromScope, access )
            if symbolInfo then
               return symbolInfo
            end
            
         end
      end
      
   end
   
   
   if not onlySameNsFlag or not self.ownerTypeInfo then
      if self.parent ~= self then
         return self.parent:getSymbolInfo( name, fromScope, onlySameNsFlag, access )
      end
      
   else
    
      local workScope = self.parent
      while workScope.parent ~= workScope do
         if _lune.nilacc( workScope.ownerTypeInfo, 'get_kind', 'callmtd' ) ~= TypeInfoKind.Class and _lune.nilacc( workScope.ownerTypeInfo, 'get_kind', 'callmtd' ) ~= TypeInfoKind.Module then
            return workScope:getSymbolInfo( name, fromScope, onlySameNsFlag, access )
         end
         
         workScope = workScope.parent
      end
      
   end
   
   do
      local _exp = sym2builtInTypeMap[name]
      if _exp ~= nil then
         return _exp
      end
   end
   
   return nil
   
end


function Scope:getTypeInfo( name, fromScope, onlySameNsFlag, access )

   local symbolInfo = self:getSymbolInfo( name, fromScope, onlySameNsFlag, access )
   if  nil == symbolInfo then
      local _symbolInfo = symbolInfo
   
      return nil
   end
   
   return symbolInfo:get_typeInfo()
end


function Scope:getSymbolTypeInfo( name, fromScope, moduleScope, access )

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
               return symbolInfo:canAccess( fromScope, access )
            end
            
         end
      end
      
   end
   
   if self.parent ~= self then
      return self.parent:getSymbolTypeInfo( name, fromScope, moduleScope, access )
   end
   
   return sym2builtInTypeMap[name]
end


function Scope:filterSymbolTypeInfo( fromScope, moduleScope, access, callback )

   if self.classFlag then
      do
         local _exp = self.symbol2SymbolInfoMap["self"]
         if _exp ~= nil then
            callback( _exp )
         end
      end
      
   end
   
   if moduleScope == fromScope or moduleScope == self or not self.classFlag then
      for __index, symbolInfo in pairs( self.symbol2SymbolInfoMap ) do
         if not callback( symbolInfo ) then
            return 
         end
         
      end
      
   end
   
   
   if self.parent ~= self then
      self.parent:filterSymbolTypeInfo( fromScope, moduleScope, access, callback )
   end
   
end


function Scope:add( processInfo, kind, canBeLeft, canBeRight, name, pos, typeInfo, accessMode, staticFlag, mutMode, hasValueFlag, isLazyLoad )

   do
      local _switchExp = kind
      if _switchExp == SymbolKind.Typ or _switchExp == SymbolKind.Fun or _switchExp == SymbolKind.Mac then
         local existSymbol
         
         do
            local _switchExp = typeInfo:get_kind()
            if _switchExp == TypeInfoKind.Enum then
               if _lune.nilacc( self.ownerTypeInfo, 'get_kind', 'callmtd' ) == TypeInfoKind.Class then
                  existSymbol = self:getSymbolInfoField( name, true, self, ScopeAccess.Full )
               else
                
                  existSymbol = self:getSymbolInfo( name, self, true, ScopeAccess.Full )
               end
               
            elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Module then
               existSymbol = self:getSymbolInfoChild( name )
            else 
               
                  existSymbol = self:getSymbolInfo( name, self, true, ScopeAccess.Full )
            end
         end
         
         if existSymbol ~= nil then
            if typeInfo:get_kind() ~= existSymbol:get_typeInfo():get_kind() or not isBuiltin( existSymbol:get_typeInfo():get_typeId().id ) then
               return nil, existSymbol
            end
            
         end
         
      end
   end
   
   
   local symbolInfo = NormalSymbolInfo.new(processInfo, kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad)
   self.symbol2SymbolInfoMap[name] = symbolInfo
   return symbolInfo, nil
end


function Scope:addLocalVar( processInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable )

   return self:add( processInfo, argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, name ~= "_", name, pos, typeInfo, AccessMode.Local, false, mutable, true, false )
end


local dummySymbol = _lune.unwrap( _moduleObj.rootScope:addLocalVar( rootProcessInfo, false, false, "$$", nil, _moduleObj.headTypeInfo, MutMode.IMut ))
_moduleObj.dummySymbol = dummySymbol


function Scope:addUnwrapedVar( processInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable )

   return self:add( processInfo, argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, pos, typeInfo, AccessMode.Local, false, mutable, true, false )
end


function Scope:addStaticVar( processInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable )

   return self:add( processInfo, argFlag and SymbolKind.Arg or SymbolKind.Var, canBeLeft, true, name, pos, typeInfo, AccessMode.Pub, true, mutable, true, false )
end


function Scope:addVar( processInfo, accessMode, name, pos, typeInfo, mutable, hasValueFlag )

   return self:add( processInfo, SymbolKind.Var, true, true, name, pos, typeInfo, accessMode, false, mutable, hasValueFlag, false )
end


function Scope:addEnumVal( processInfo, name, pos, typeInfo )

   return self:add( processInfo, SymbolKind.Mbr, false, true, name, pos, typeInfo, AccessMode.Pub, true, MutMode.Mut, true, false )
end


function Scope:addEnum( processInfo, accessMode, name, pos, typeInfo )

   return self:add( processInfo, SymbolKind.Typ, false, false, name, pos, typeInfo, accessMode, true, MutMode.Mut, true, false )
end


function Scope:addAlgeVal( processInfo, name, pos, typeInfo )

   return self:add( processInfo, SymbolKind.Mbr, false, true, name, pos, typeInfo, AccessMode.Pub, true, MutMode.Mut, true, false )
end


function Scope:addAlge( processInfo, accessMode, name, pos, typeInfo )

   return self:add( processInfo, SymbolKind.Typ, false, false, name, pos, typeInfo, accessMode, true, MutMode.Mut, true, false )
end


function Scope:addAlternate( processInfo, accessMode, name, pos, typeInfo )

   self:add( processInfo, SymbolKind.Typ, false, false, name, pos, typeInfo, accessMode, true, MutMode.Mut, true, false )
end


function Scope:addMember( processInfo, name, pos, typeInfo, accessMode, staticFlag, mutMode )

   return self:add( processInfo, SymbolKind.Mbr, true, true, name, pos, typeInfo, accessMode, staticFlag, mutMode, true, false )
end


function Scope:addMethod( processInfo, pos, typeInfo, accessMode, staticFlag, mutable )

   return self:add( processInfo, SymbolKind.Mtd, true, staticFlag, typeInfo:get_rawTxt(), pos, typeInfo, accessMode, staticFlag, mutable and MutMode.Mut or MutMode.IMut, true, false )
end


function Scope:addFunc( processInfo, pos, typeInfo, accessMode, staticFlag, mutable )

   return self:add( processInfo, SymbolKind.Fun, true, true, typeInfo:get_rawTxt(), pos, typeInfo, accessMode, staticFlag, mutable and MutMode.Mut or MutMode.IMut, true, false )
end


function Scope:addForm( processInfo, pos, typeInfo, accessMode )

   return self:add( processInfo, SymbolKind.Typ, false, false, typeInfo:get_rawTxt(), pos, typeInfo, accessMode, true, MutMode.IMut, false, false )
end


function Scope:addMacro( processInfo, pos, typeInfo, accessMode )

   return self:add( processInfo, SymbolKind.Mac, false, false, typeInfo:get_rawTxt(), pos, typeInfo, accessMode, true, MutMode.IMut, true, false )
end


function Scope:addClassLazy( processInfo, name, pos, typeInfo, lazyLoad )

   return self:add( processInfo, SymbolKind.Typ, false, false, name, pos, typeInfo, typeInfo:get_accessMode(), true, MutMode.Mut, true, lazyLoad )
end


function Scope:addClass( processInfo, name, pos, typeInfo )

   return self:addClassLazy( processInfo, name, pos, typeInfo, false )
end


function Scope:addExtModule( processInfo, name, pos, typeInfo, lazy, lang )

   if lang ~= Types.Lang.Same then
      do
         local _matchExp = processInfo:createLuaval( typeInfo, true )
         if _matchExp[1] == LuavalResult.Err[1] then
            local mess = _matchExp[2][1]
         
            Util.err( mess )
         elseif _matchExp[1] == LuavalResult.OK[1] then
            local luavalTypeInfo = _matchExp[2][1]
            local _ = _matchExp[2][2]
         
            typeInfo = luavalTypeInfo
         end
      end
      
   end
   
   return self:add( processInfo, SymbolKind.Typ, false, false, name, pos, typeInfo, typeInfo:get_accessMode(), false, MutMode.Mut, true, lazy )
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
                     Util.log( string.format( "scope: %s, %s, %s", prefix, _exp, symbol) )
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
function Scope:setClosure( workSymbol )

   local symbol = workSymbol:getOrg(  )
   local targetFuncType = symbol:get_namespaceTypeInfo()
   local funcType = self:getNamespaceTypeInfo(  )
   
   while true do
      local funcScope = _lune.unwrap( funcType:get_scope())
      if not funcScope.closureSymMap[symbol:get_symbolId()] then
         funcScope.closureSymMap[symbol:get_symbolId()] = symbol
         funcScope.closureSym2NumMap[symbol] = #funcScope.closureSymList
         table.insert( funcScope.closureSymList, symbol )
         funcType = funcScope.parent:getNamespaceTypeInfo(  )
      else
       
         break
      end
      
      if funcType == targetFuncType then
         break
      end
      
   end
   
   if not symbol:get_hasAccessFromClosure() then
      symbol:set_hasAccessFromClosure( true )
   end
   
end

function Scope:isClosureAccess( moduleScope, symbol )

   local processInfo = moduleScope:getModule(  ):getProcessInfo(  )
   do
      local _switchExp = symbol:get_kind()
      if _switchExp == SymbolKind.Var or _switchExp == SymbolKind.Arg or _switchExp == SymbolKind.Fun then
         if symbol:get_scope() == moduleScope or symbol:get_scope() == _moduleObj.rootScope then
            
         elseif symbol:get_name() == "self" then
            local funcType = self:getNamespaceTypeInfo(  )
            if funcType:get_parentInfo():isInheritFrom( processInfo, symbol:get_namespaceTypeInfo():get_parentInfo() ) then
               
            else
             
               return true
            end
            
         else
          
            local funcType = self:getNamespaceTypeInfo(  )
            if funcType ~= symbol:get_namespaceTypeInfo() then
               return true
            end
            
         end
         
      end
   end
   
   return false
end


function Scope:accessSymbol( moduleScope, symbol )

   if symbol:get_kind() == SymbolKind.Fun and self:getNamespaceTypeInfo(  ) == symbol:get_typeInfo() then
      return 
   end
   
   
   if self:isClosureAccess( moduleScope, symbol ) then
      self:setClosure( symbol )
   end
   
end


function TypeInfo:createAlt2typeMap( detectFlag )

   return CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( detectFlag )
end


local NilTypeInfo = {}
setmetatable( NilTypeInfo, { __index = TypeInfo } )
_moduleObj.NilTypeInfo = NilTypeInfo
function NilTypeInfo.new( processInfo )
   local obj = {}
   NilTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo ); end
   return obj
end
function NilTypeInfo:__init(processInfo) 
   TypeInfo.__init( self,nil, processInfo)
   
   
   self.typeId = processInfo:newId(  )
end
function NilTypeInfo:isModule(  )

   return false
end
function NilTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function NilTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   return "nil"
end
function NilTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return other:get_nilable(), nil
end
function NilTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function NilTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( "nil", nil )
end
function NilTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

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
function NilTypeInfo:getParentFullName( typeNameCtrl, importInfo, localFlag )

   return ""
end
function NilTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = NilTypeInfo  } )
end
function NilTypeInfo:get_typeId()
   return self.typeId
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


function NormalSymbolInfo:canAccess( fromScope, access )

   if access == ScopeAccess.Full then
      return self
   end
   
   if self.scope == fromScope then
      return self
   end
   
   
   local processInfo = fromScope:getProcessInfo(  )
   
   do
      local _switchExp = self:get_accessMode()
      if _switchExp == AccessMode.Pub or _switchExp == AccessMode.Global then
         return self
      elseif _switchExp == AccessMode.Pro then
         local nsClass = self.scope:getClassTypeInfo(  )
         local fromClass = fromScope:getClassTypeInfo(  )
         if fromClass:isInheritFrom( processInfo, nsClass, nil ) then
            return self
         end
         
         return nil
      elseif _switchExp == AccessMode.Local then
         if isBuiltin( self:get_typeInfo():get_typeId().id ) or self:getModule(  ) == fromScope:getModule(  ) then
            return self
         end
         
         return nil
      elseif _switchExp == AccessMode.Pri then
         if fromScope:isInnerOf( self.scope ) then
            return self
         end
         
         return nil
      end
   end
   
   Util.err( string.format( "illegl accessmode -- %s, %s", self:get_accessMode(), self:get_name()) )
end


local OverrideMut = {}
OverrideMut._name2Val = {}
_moduleObj.OverrideMut = OverrideMut
function OverrideMut:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "OverrideMut.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function OverrideMut._from( val )
   return _lune._AlgeFrom( OverrideMut, val )
end

OverrideMut.IMut = { "IMut", {{}}}
OverrideMut._name2Val["IMut"] = OverrideMut.IMut
OverrideMut.None = { "None"}
OverrideMut._name2Val["None"] = OverrideMut.None
OverrideMut.Prefix = { "Prefix", {{}}}
OverrideMut._name2Val["Prefix"] = OverrideMut.Prefix


local AccessSymbolInfo = {}
setmetatable( AccessSymbolInfo, { __index = SymbolInfo } )
_moduleObj.AccessSymbolInfo = AccessSymbolInfo
function AccessSymbolInfo.new( symbolInfo, overrideMut, overrideCanBeLeft )
   local obj = {}
   AccessSymbolInfo.setmeta( obj )
   if obj.__init then obj:__init( symbolInfo, overrideMut, overrideCanBeLeft ); end
   return obj
end
function AccessSymbolInfo:__init(symbolInfo, overrideMut, overrideCanBeLeft) 
   SymbolInfo.__init( self)
   
   self.symbolInfo = symbolInfo
   self.overrideMut = overrideMut
   self.overrideCanBeLeft = overrideCanBeLeft
   local symType = symbolInfo:get_typeInfo()
   local work
   
   do
      local _matchExp = self.overrideMut
      if _matchExp[1] == OverrideMut.None[1] then
      
         work = symType
      elseif _matchExp[1] == OverrideMut.Prefix[1] then
         local prefixTypeInfo = _matchExp[2][1]
      
         if self.symbolInfo:get_kind() == SymbolKind.Mbr and symType:get_kind() == TypeInfoKind.Alternate and prefixTypeInfo:get_kind() == TypeInfoKind.Class and #prefixTypeInfo:get_itemTypeInfoList() > 0 then
            local alt2TypeMap = prefixTypeInfo:createAlt2typeMap( false )
            local typeInfo = symType:applyGeneric( alt2TypeMap, symType:getModule(  ) )
            if typeInfo ~= nil then
               work = typeInfo
            else
               work = symType
            end
            
         else
          
            work = symType
         end
         
      elseif _matchExp[1] == OverrideMut.IMut[1] then
         local typeInfo = _matchExp[2][1]
      
         work = typeInfo
      end
   end
   
   self.overrideTypeInfo = work
end
function AccessSymbolInfo:getOrg(  )

   return self.symbolInfo:getOrg(  )
end
function AccessSymbolInfo:canAccess( fromScope, access )

   if self.symbolInfo:canAccess( fromScope, access ) then
      return self
   end
   
   return nil
end
function AccessSymbolInfo:get_typeInfo(  )

   return self.overrideTypeInfo
end
function AccessSymbolInfo:get_mutMode(  )

   do
      local _matchExp = self.overrideMut
      if _matchExp[1] == OverrideMut.None[1] then
      
         
      elseif _matchExp[1] == OverrideMut.Prefix[1] then
         local prefixTypeInfo = _matchExp[2][1]
      
         do
            local _switchExp = self.symbolInfo:get_mutMode()
            if _switchExp == MutMode.AllMut or _switchExp == MutMode.IMut or _switchExp == MutMode.IMutRe then
               return self.symbolInfo:get_mutMode()
            elseif _switchExp == MutMode.Mut then
               do
                  local _switchExp = prefixTypeInfo:get_mutMode()
                  if _switchExp == MutMode.AllMut then
                     return MutMode.Mut
                  elseif _switchExp == MutMode.Mut or _switchExp == MutMode.IMut or _switchExp == MutMode.IMutRe then
                     return prefixTypeInfo:get_mutMode()
                  end
               end
               
            end
         end
         
      elseif _matchExp[1] == OverrideMut.IMut[1] then
         local _ = _matchExp[2][1]
      
         return MutMode.IMut
      end
   end
   
   return self.symbolInfo:get_mutMode()
end
function AccessSymbolInfo:get_mutable(  )

   return isMutable( self:get_mutMode(  ) )
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
function AccessSymbolInfo:get_symbolInfo()
   return self.symbolInfo
end
function AccessSymbolInfo:clearValue( ... )
   return self.symbolInfo:clearValue( ... )
end

function AccessSymbolInfo:getModule( ... )
   return self.symbolInfo:getModule( ... )
end

function AccessSymbolInfo:get_accessMode( ... )
   return self.symbolInfo:get_accessMode( ... )
end

function AccessSymbolInfo:get_canBeRight( ... )
   return self.symbolInfo:get_canBeRight( ... )
end

function AccessSymbolInfo:get_convModuleParam( ... )
   return self.symbolInfo:get_convModuleParam( ... )
end

function AccessSymbolInfo:get_hasAccessFromClosure( ... )
   return self.symbolInfo:get_hasAccessFromClosure( ... )
end

function AccessSymbolInfo:get_hasValueFlag( ... )
   return self.symbolInfo:get_hasValueFlag( ... )
end

function AccessSymbolInfo:get_isLazyLoad( ... )
   return self.symbolInfo:get_isLazyLoad( ... )
end

function AccessSymbolInfo:get_kind( ... )
   return self.symbolInfo:get_kind( ... )
end

function AccessSymbolInfo:get_name( ... )
   return self.symbolInfo:get_name( ... )
end

function AccessSymbolInfo:get_namespaceTypeInfo( ... )
   return self.symbolInfo:get_namespaceTypeInfo( ... )
end

function AccessSymbolInfo:get_pos( ... )
   return self.symbolInfo:get_pos( ... )
end

function AccessSymbolInfo:get_posForLatestMod( ... )
   return self.symbolInfo:get_posForLatestMod( ... )
end

function AccessSymbolInfo:get_posForModToRef( ... )
   return self.symbolInfo:get_posForModToRef( ... )
end

function AccessSymbolInfo:get_scope( ... )
   return self.symbolInfo:get_scope( ... )
end

function AccessSymbolInfo:get_staticFlag( ... )
   return self.symbolInfo:get_staticFlag( ... )
end

function AccessSymbolInfo:get_symbolId( ... )
   return self.symbolInfo:get_symbolId( ... )
end

function AccessSymbolInfo:hasAccess( ... )
   return self.symbolInfo:hasAccess( ... )
end

function AccessSymbolInfo:set_convModuleParam( ... )
   return self.symbolInfo:set_convModuleParam( ... )
end

function AccessSymbolInfo:set_hasAccessFromClosure( ... )
   return self.symbolInfo:set_hasAccessFromClosure( ... )
end

function AccessSymbolInfo:set_hasValueFlag( ... )
   return self.symbolInfo:set_hasValueFlag( ... )
end

function AccessSymbolInfo:set_posForLatestMod( ... )
   return self.symbolInfo:set_posForLatestMod( ... )
end

function AccessSymbolInfo:set_posForModToRef( ... )
   return self.symbolInfo:set_posForModToRef( ... )
end

function AccessSymbolInfo:set_typeInfo( ... )
   return self.symbolInfo:set_typeInfo( ... )
end

function AccessSymbolInfo:updateValue( ... )
   return self.symbolInfo:updateValue( ... )
end



local AlternateTypeInfo = {}
setmetatable( AlternateTypeInfo, { __index = TypeInfo } )
_moduleObj.AlternateTypeInfo = AlternateTypeInfo
function AlternateTypeInfo.new( processInfo, belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList )
   local obj = {}
   AlternateTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList ); end
   return obj
end
function AlternateTypeInfo:__init(processInfo, belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList) 
   TypeInfo.__init( self,TypeInfo.createScope( processInfo, nil, true, baseTypeInfo, interfaceList ), processInfo)
   
   self.typeId = processInfo:newId(  )
   
   self.txt = txt
   self.accessMode = accessMode
   self.parentInfo = parentInfo
   self.baseTypeInfo = _lune.unwrapDefault( baseTypeInfo, _moduleObj.headTypeInfo)
   self.interfaceList = _lune.unwrapDefault( interfaceList, {})
   self.belongClassFlag = belongClassFlag
   self.altIndex = altIndex
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
end
function AlternateTypeInfo:updateParentInfo( typeInfo )

   self.parentInfo = typeInfo
end
function AlternateTypeInfo:isModule(  )

   return false
end
function AlternateTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function AlternateTypeInfo:get_baseId(  )

   return self.baseTypeInfo:get_typeId()
end
function AlternateTypeInfo:get_parentInfo(  )

   return self.parentInfo
end
function AlternateTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function AlternateTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   return self.txt
end
function AlternateTypeInfo:isInheritFrom( processInfo, other, alt2type )

   local workAlt2type
   
   if alt2type ~= nil then
      local otherWork = AlternateTypeInfo.getAssign( other, alt2type )
      if self == otherWork:get_srcTypeInfo() then
         return true
      end
      
      
      do
         local genType = alt2type[self]
         if genType ~= nil then
            return genType:isInheritFrom( processInfo, otherWork, alt2type )
         end
      end
      
      if not CanEvalCtrlTypeInfo.isValidApply( alt2type ) then
         workAlt2type = nil
      else
       
         workAlt2type = alt2type
      end
      
   else
      workAlt2type = nil
   end
   
   if self == other:get_srcTypeInfo() then
      return true
   end
   
   
   local function check(  )
   
      if self:hasBase(  ) then
         if self.baseTypeInfo:isInheritFrom( processInfo, other, workAlt2type ) then
            return true
         end
         
      end
      
      
      for __index, ifType in ipairs( self.interfaceList ) do
         if ifType:isInheritFrom( processInfo, other, workAlt2type ) then
            return true
         end
         
      end
      
      return false
   end
   
   if check(  ) then
      if workAlt2type ~= nil then
         workAlt2type[self] = other
      end
      
      return true
   end
   
   return false
end
function AlternateTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   if self == other:get_srcTypeInfo() then
      return true, nil
   end
   
   if other:get_nilable() then
      return false, nil
   end
   
   if not self.belongClassFlag then
      do
         local altType = _lune.__Cast( other, 3, AlternateTypeInfo )
         if altType ~= nil then
            if not altType.belongClassFlag and altType.altIndex == self.altIndex then
               return true, nil
            end
            
         end
      end
      
   end
   
   return self:canSetFrom( processInfo, other, canEvalType, alt2type ), nil
end
function AlternateTypeInfo:get_display_stirng_with( raw, alt2type )

   if alt2type ~= nil then
      do
         local genType = alt2type[self]
         if genType ~= nil then
            return genType:get_display_stirng_with( genType:get_rawTxt(), alt2type )
         end
      end
      
   end
   
   return self:getTxtWithRaw( raw )
end
function AlternateTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self.txt, nil )
end
function AlternateTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   if self == typeInfo then
      return true
   end
   
   if not self.belongClassFlag then
      do
         local altType = _lune.__Cast( typeInfo, 3, AlternateTypeInfo )
         if altType ~= nil then
            if not altType.belongClassFlag and altType.altIndex == self.altIndex then
               return true
            end
            
         end
      end
      
   end
   
   
   if alt2type ~= nil then
      return self:canSetFrom( processInfo, typeInfo, nil, alt2type )
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

   return MutMode.Mut
end
function AlternateTypeInfo:serialize( stream, serializeInfo )

   local parentId = self:getParentId(  )
   stream:write( string.format( '{ skind = %d, parentId = %d, typeId = %d, txt = %q, ', SerializeKind.Alternate, parentId.id, self.typeId.id, self.txt) .. string.format( 'accessMode = %d, baseId = %d, ', self.accessMode, self:get_baseId().id) .. string.format( 'belongClassFlag = %s, altIndex = %d, ', self.belongClassFlag, self.altIndex) )
   
   stream:write( self:serializeTypeInfoList( "ifList = {", self.interfaceList ) )
   stream:write( "}\n" )
end
function AlternateTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

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
function AlternateTypeInfo:get_altIndex()
   return self.altIndex
end


local boxRootAltType = AlternateTypeInfo.new(rootProcessInfo, true, 1, "_T", AccessMode.Pub, _moduleObj.headTypeInfo)
local boxRootScope = Scope.new(rootProcessInfo, _moduleObj.rootScope, true, nil)

local BoxTypeInfo = {}
setmetatable( BoxTypeInfo, { __index = TypeInfo } )
_moduleObj.BoxTypeInfo = BoxTypeInfo
function BoxTypeInfo.new( processInfo, typeId, accessMode, boxingType )
   local obj = {}
   BoxTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, typeId, accessMode, boxingType ); end
   return obj
end
function BoxTypeInfo:__init(processInfo, typeId, accessMode, boxingType) 
   TypeInfo.__init( self,boxRootScope, processInfo)
   
   self.boxingType = boxingType
   self.typeId = typeId
   self.itemTypeInfoList = {boxingType}
   self.accessMode = accessMode
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
end
function BoxTypeInfo:get_scope(  )

   return TypeInfo.get_scope( self)
end
function BoxTypeInfo:get_kind(  )

   return TypeInfoKind.Box
end
function BoxTypeInfo:get_aliasSrc(  )

   return self
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
function BoxTypeInfo:get_extedType(  )

   return self
end
function BoxTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function BoxTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   return "Nilable<" .. self.boxingType:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag ) .. ">"
end
function BoxTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function BoxTypeInfo:get_display_stirng_with( raw, alt2type )

   return string.format( "Nilable<%s>", self.boxingType:get_display_stirng_with( raw, alt2type ))
end
function BoxTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind = %d, typeId = %d, accessMode = %d, boxingType = %d }\n', SerializeKind.Box, self.typeId.id, self.accessMode, self.boxingType:get_typeId().id) )
end
function BoxTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   do
      local boxType = _lune.__Cast( typeInfo, 3, BoxTypeInfo )
      if boxType ~= nil then
         return self.boxingType:equals( processInfo, boxType.boxingType, alt2type, checkModifer )
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
function BoxTypeInfo:addChildren( ... )
   return self.boxingType:addChildren( ... )
end

function BoxTypeInfo:getFullName( ... )
   return self.boxingType:getFullName( ... )
end

function BoxTypeInfo:getModule( ... )
   return self.boxingType:getModule( ... )
end

function BoxTypeInfo:getOverridingType( ... )
   return self.boxingType:getOverridingType( ... )
end

function BoxTypeInfo:getParentFullName( ... )
   return self.boxingType:getParentFullName( ... )
end

function BoxTypeInfo:getParentId( ... )
   return self.boxingType:getParentId( ... )
end

function BoxTypeInfo:getProcessInfo( ... )
   return self.boxingType:getProcessInfo( ... )
end

function BoxTypeInfo:get_abstractFlag( ... )
   return self.boxingType:get_abstractFlag( ... )
end

function BoxTypeInfo:get_argTypeInfoList( ... )
   return self.boxingType:get_argTypeInfoList( ... )
end

function BoxTypeInfo:get_autoFlag( ... )
   return self.boxingType:get_autoFlag( ... )
end

function BoxTypeInfo:get_baseId( ... )
   return self.boxingType:get_baseId( ... )
end

function BoxTypeInfo:get_baseTypeInfo( ... )
   return self.boxingType:get_baseTypeInfo( ... )
end

function BoxTypeInfo:get_children( ... )
   return self.boxingType:get_children( ... )
end

function BoxTypeInfo:get_externalFlag( ... )
   return self.boxingType:get_externalFlag( ... )
end

function BoxTypeInfo:get_genSrcTypeInfo( ... )
   return self.boxingType:get_genSrcTypeInfo( ... )
end

function BoxTypeInfo:get_interfaceList( ... )
   return self.boxingType:get_interfaceList( ... )
end

function BoxTypeInfo:get_mutMode( ... )
   return self.boxingType:get_mutMode( ... )
end

function BoxTypeInfo:get_parentInfo( ... )
   return self.boxingType:get_parentInfo( ... )
end

function BoxTypeInfo:get_processInfo( ... )
   return self.boxingType:get_processInfo( ... )
end

function BoxTypeInfo:get_rawTxt( ... )
   return self.boxingType:get_rawTxt( ... )
end

function BoxTypeInfo:get_retTypeInfoList( ... )
   return self.boxingType:get_retTypeInfoList( ... )
end

function BoxTypeInfo:get_staticFlag( ... )
   return self.boxingType:get_staticFlag( ... )
end

function BoxTypeInfo:get_typeData( ... )
   return self.boxingType:get_typeData( ... )
end

function BoxTypeInfo:hasBase( ... )
   return self.boxingType:hasBase( ... )
end

function BoxTypeInfo:hasRouteNamespaceFrom( ... )
   return self.boxingType:hasRouteNamespaceFrom( ... )
end

function BoxTypeInfo:isInheritFrom( ... )
   return self.boxingType:isInheritFrom( ... )
end

function BoxTypeInfo:isModule( ... )
   return self.boxingType:isModule( ... )
end

function BoxTypeInfo:serializeTypeInfoList( ... )
   return self.boxingType:serializeTypeInfoList( ... )
end

function BoxTypeInfo:switchScope( ... )
   return self.boxingType:switchScope( ... )
end



local GenericTypeInfo = {}
setmetatable( GenericTypeInfo, { __index = TypeInfo } )
_moduleObj.GenericTypeInfo = GenericTypeInfo
function GenericTypeInfo:get_display_stirng_with( raw, alt2type )

   return self.genSrcTypeInfo:get_display_stirng_with( raw, self.alt2typeMap )
end
function GenericTypeInfo.new( processInfo, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo )
   local obj = {}
   GenericTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo ); end
   return obj
end
function GenericTypeInfo:__init(processInfo, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo) 
   TypeInfo.__init( self,TypeInfo.createScope( processInfo, (_lune.unwrap( genSrcTypeInfo:get_scope()) ):get_parent(), true, genSrcTypeInfo, nil ), processInfo)
   
   self.typeId = processInfo:newId(  )
   self.moduleTypeInfo = moduleTypeInfo
   
   self.itemTypeInfoList = itemTypeInfoList
   self.genSrcTypeInfo = genSrcTypeInfo
   
   if #genSrcTypeInfo:get_itemTypeInfoList() ~= #itemTypeInfoList then
      Util.err( string.format( "unmatch generic type number -- %d, %d", #genSrcTypeInfo:get_itemTypeInfoList(), #itemTypeInfoList) )
   end
   
   local alt2typeMap = {}
   local workAlt2typeMap = CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   local hasAlter = false
   for index, altTypeInfo in ipairs( genSrcTypeInfo:get_itemTypeInfoList() ) do
      local itemType = itemTypeInfoList[index]
      alt2typeMap[altTypeInfo] = itemType
      if itemType:applyGeneric( workAlt2typeMap, moduleTypeInfo ) ~= itemType then
         hasAlter = true
      end
      
   end
   
   self.hasAlter = hasAlter
   self.alt2typeMap = alt2typeMap
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
end
function GenericTypeInfo:getModule(  )

   return self.moduleTypeInfo
end
function GenericTypeInfo:isInheritFrom( processInfo, other, alt2type )

   local otherSrc = other:get_genSrcTypeInfo()
   
   if not self.genSrcTypeInfo:isInheritFrom( processInfo, otherSrc, alt2type ) then
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
   
   for __index, altType in ipairs( otherSrc:get_itemTypeInfoList() ) do
      local genType = self.alt2typeMap[altType]
      if  nil == genType then
         local _genType = genType
      
         return false
      end
      
      local otherGenType = _lune.unwrap( genOther.alt2typeMap[altType])
      if not otherGenType:canEvalWith( processInfo, genType, CanEvalType.SetEq, workAlt2type ) then
         return false
      end
      
   end
   
   return true
end
function GenericTypeInfo:get_aliasSrc(  )

   return self
end
function GenericTypeInfo:get_srcTypeInfo(  )

   return self
end
function GenericTypeInfo:get_extedType(  )

   return self
end
function GenericTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   if other:get_nilable() then
      return false, nil
   end
   
   
   if TypeInfo.isMut( self ) and not TypeInfo.isMut( other ) then
      return false, nil
   end
   
   
   local otherSrc = other:get_srcTypeInfo()
   if self == otherSrc then
      return true, nil
   end
   
   local work = otherSrc
   while true do
      if work == _moduleObj.headTypeInfo then
         return false, nil
      end
      
      for altType, genType in pairs( work:createAlt2typeMap( false ) ) do
         alt2type[altType] = genType
      end
      
      if self.genSrcTypeInfo:equals( processInfo, work:get_genSrcTypeInfo(), alt2type ) then
         break
      end
      
      for __index, ifType in ipairs( work:get_interfaceList() ) do
         if self:canEvalWith( processInfo, ifType, canEvalType, alt2type ) then
            return true, nil
         end
         
      end
      
      work = work:get_baseTypeInfo()
   end
   
   do
      local otherGen = _lune.__Cast( work, 3, GenericTypeInfo )
      if otherGen ~= nil then
         
         local evalType
         
         if canEvalType == CanEvalType.SetOp then
            evalType = CanEvalType.SetEq
         else
          
            evalType = canEvalType
         end
         
         
         for key, val in pairs( self.alt2typeMap ) do
            local otherType = AlternateTypeInfo.getAssign( _lune.unwrap( otherGen.alt2typeMap[key]), alt2type )
            local ret, mess = val:canEvalWith( processInfo, otherType, evalType, alt2type )
            if not ret then
               return false, mess
            end
            
         end
         
      end
   end
   
   return true, nil
end
function GenericTypeInfo:equals( processInfo, other, alt2type, checkModifer )

   if self == other then
      return true
   end
   
   if self:get_kind() ~= self:get_kind() or #self.itemTypeInfoList ~= #other:get_itemTypeInfoList() then
      return false
   end
   
   if not (_lune.__Cast( other, 3, GenericTypeInfo ) ) then
      return false
   end
   
   
   if not self.genSrcTypeInfo:equals( processInfo, other:get_genSrcTypeInfo(), alt2type, checkModifer ) then
      return false
   end
   
   
   for index, otherItem in ipairs( other:get_itemTypeInfoList() ) do
      local typeInfo = self.itemTypeInfoList[index]
      if not typeInfo:equals( processInfo, otherItem, alt2type, checkModifer ) then
         return false
      end
      
   end
   
   return true
end
function GenericTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind = %d, typeId = %d, genSrcTypeId = %d, genTypeList = {', SerializeKind.Generic, self.typeId.id, self.genSrcTypeInfo:get_typeId().id) )
   local count = 0
   for __index, genType in pairs( self.alt2typeMap ) do
      if count > 0 then
         stream:write( "," )
      end
      
      stream:write( string.format( "%d", genType:get_typeId().id) )
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
function GenericTypeInfo:addChildren( ... )
   return self.genSrcTypeInfo:addChildren( ... )
end

function GenericTypeInfo:getFullName( ... )
   return self.genSrcTypeInfo:getFullName( ... )
end

function GenericTypeInfo:getOverridingType( ... )
   return self.genSrcTypeInfo:getOverridingType( ... )
end

function GenericTypeInfo:getParentFullName( ... )
   return self.genSrcTypeInfo:getParentFullName( ... )
end

function GenericTypeInfo:getParentId( ... )
   return self.genSrcTypeInfo:getParentId( ... )
end

function GenericTypeInfo:getProcessInfo( ... )
   return self.genSrcTypeInfo:getProcessInfo( ... )
end

function GenericTypeInfo:getTxt( ... )
   return self.genSrcTypeInfo:getTxt( ... )
end

function GenericTypeInfo:getTxtWithRaw( ... )
   return self.genSrcTypeInfo:getTxtWithRaw( ... )
end

function GenericTypeInfo:get_abstractFlag( ... )
   return self.genSrcTypeInfo:get_abstractFlag( ... )
end

function GenericTypeInfo:get_accessMode( ... )
   return self.genSrcTypeInfo:get_accessMode( ... )
end

function GenericTypeInfo:get_argTypeInfoList( ... )
   return self.genSrcTypeInfo:get_argTypeInfoList( ... )
end

function GenericTypeInfo:get_autoFlag( ... )
   return self.genSrcTypeInfo:get_autoFlag( ... )
end

function GenericTypeInfo:get_baseId( ... )
   return self.genSrcTypeInfo:get_baseId( ... )
end

function GenericTypeInfo:get_baseTypeInfo( ... )
   return self.genSrcTypeInfo:get_baseTypeInfo( ... )
end

function GenericTypeInfo:get_children( ... )
   return self.genSrcTypeInfo:get_children( ... )
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

function GenericTypeInfo:get_kind( ... )
   return self.genSrcTypeInfo:get_kind( ... )
end

function GenericTypeInfo:get_mutMode( ... )
   return self.genSrcTypeInfo:get_mutMode( ... )
end

function GenericTypeInfo:get_nilable( ... )
   return self.genSrcTypeInfo:get_nilable( ... )
end

function GenericTypeInfo:get_nonnilableType( ... )
   return self.genSrcTypeInfo:get_nonnilableType( ... )
end

function GenericTypeInfo:get_parentInfo( ... )
   return self.genSrcTypeInfo:get_parentInfo( ... )
end

function GenericTypeInfo:get_processInfo( ... )
   return self.genSrcTypeInfo:get_processInfo( ... )
end

function GenericTypeInfo:get_rawTxt( ... )
   return self.genSrcTypeInfo:get_rawTxt( ... )
end

function GenericTypeInfo:get_retTypeInfoList( ... )
   return self.genSrcTypeInfo:get_retTypeInfoList( ... )
end

function GenericTypeInfo:get_scope( ... )
   return self.genSrcTypeInfo:get_scope( ... )
end

function GenericTypeInfo:get_staticFlag( ... )
   return self.genSrcTypeInfo:get_staticFlag( ... )
end

function GenericTypeInfo:get_typeData( ... )
   return self.genSrcTypeInfo:get_typeData( ... )
end

function GenericTypeInfo:hasBase( ... )
   return self.genSrcTypeInfo:hasBase( ... )
end

function GenericTypeInfo:hasRouteNamespaceFrom( ... )
   return self.genSrcTypeInfo:hasRouteNamespaceFrom( ... )
end

function GenericTypeInfo:isModule( ... )
   return self.genSrcTypeInfo:isModule( ... )
end

function GenericTypeInfo:serializeTypeInfoList( ... )
   return self.genSrcTypeInfo:serializeTypeInfoList( ... )
end

function GenericTypeInfo:switchScope( ... )
   return self.genSrcTypeInfo:switchScope( ... )
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
function ModifierTypeInfo:get_extedType(  )

   return self
end
function ModifierTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function ModifierTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   local txt = self.srcTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )
   if not isMutable( self.mutMode ) then
      txt = "&" .. txt
   end
   
   return txt
end
function ModifierTypeInfo:get_display_stirng_with( raw, alt2type )

   local txt = self.srcTypeInfo:get_display_stirng_with( raw, alt2type )
   if isMutable( self.mutMode ) then
      txt = "mut " .. txt
   end
   
   return txt
end
function ModifierTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function ModifierTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind = %d, typeId = %d, srcTypeId = %d, mutMode = %d }\n', SerializeKind.Modifier, self.typeId.id, self.srcTypeInfo:get_typeId().id, self.mutMode) )
end
function ModifierTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   local evalType
   
   if #self.srcTypeInfo:get_itemTypeInfoList() >= 1 then
      do
         local _switchExp = canEvalType
         if _switchExp == CanEvalType.SetEq or _switchExp == CanEvalType.SetOp then
            evalType = CanEvalType.SetOpIMut
         else 
            
               evalType = canEvalType
         end
      end
      
   else
    
      evalType = canEvalType
   end
   
   
   return TypeInfo.canEvalWithBase( processInfo, self.srcTypeInfo, TypeInfo.isMut( self ), other:get_srcTypeInfo(), evalType, alt2type )
end
function ModifierTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   if checkModifer then
      if TypeInfo.isMut( self ) ~= TypeInfo.isMut( typeInfo ) then
         return false
      end
      
   end
   
   return self.srcTypeInfo:equals( processInfo, typeInfo:get_srcTypeInfo(), alt2type, checkModifer )
end
function ModifierTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = ModifierTypeInfo  } )
end
function ModifierTypeInfo.new( __superarg1, __superarg2,srcTypeInfo, typeId, mutMode )
   local obj = {}
   ModifierTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2,srcTypeInfo, typeId, mutMode )
   end
   return obj
end
function ModifierTypeInfo:__init( __superarg1, __superarg2,srcTypeInfo, typeId, mutMode )

   TypeInfo.__init( self, __superarg1, __superarg2 )
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
function ModifierTypeInfo:addChildren( ... )
   return self.srcTypeInfo:addChildren( ... )
end

function ModifierTypeInfo:createAlt2typeMap( ... )
   return self.srcTypeInfo:createAlt2typeMap( ... )
end

function ModifierTypeInfo:getFullName( ... )
   return self.srcTypeInfo:getFullName( ... )
end

function ModifierTypeInfo:getModule( ... )
   return self.srcTypeInfo:getModule( ... )
end

function ModifierTypeInfo:getOverridingType( ... )
   return self.srcTypeInfo:getOverridingType( ... )
end

function ModifierTypeInfo:getParentFullName( ... )
   return self.srcTypeInfo:getParentFullName( ... )
end

function ModifierTypeInfo:getParentId( ... )
   return self.srcTypeInfo:getParentId( ... )
end

function ModifierTypeInfo:getProcessInfo( ... )
   return self.srcTypeInfo:getProcessInfo( ... )
end

function ModifierTypeInfo:get_abstractFlag( ... )
   return self.srcTypeInfo:get_abstractFlag( ... )
end

function ModifierTypeInfo:get_accessMode( ... )
   return self.srcTypeInfo:get_accessMode( ... )
end

function ModifierTypeInfo:get_aliasSrc( ... )
   return self.srcTypeInfo:get_aliasSrc( ... )
end

function ModifierTypeInfo:get_argTypeInfoList( ... )
   return self.srcTypeInfo:get_argTypeInfoList( ... )
end

function ModifierTypeInfo:get_autoFlag( ... )
   return self.srcTypeInfo:get_autoFlag( ... )
end

function ModifierTypeInfo:get_baseId( ... )
   return self.srcTypeInfo:get_baseId( ... )
end

function ModifierTypeInfo:get_baseTypeInfo( ... )
   return self.srcTypeInfo:get_baseTypeInfo( ... )
end

function ModifierTypeInfo:get_children( ... )
   return self.srcTypeInfo:get_children( ... )
end

function ModifierTypeInfo:get_externalFlag( ... )
   return self.srcTypeInfo:get_externalFlag( ... )
end

function ModifierTypeInfo:get_genSrcTypeInfo( ... )
   return self.srcTypeInfo:get_genSrcTypeInfo( ... )
end

function ModifierTypeInfo:get_interfaceList( ... )
   return self.srcTypeInfo:get_interfaceList( ... )
end

function ModifierTypeInfo:get_itemTypeInfoList( ... )
   return self.srcTypeInfo:get_itemTypeInfoList( ... )
end

function ModifierTypeInfo:get_kind( ... )
   return self.srcTypeInfo:get_kind( ... )
end

function ModifierTypeInfo:get_nilable( ... )
   return self.srcTypeInfo:get_nilable( ... )
end

function ModifierTypeInfo:get_parentInfo( ... )
   return self.srcTypeInfo:get_parentInfo( ... )
end

function ModifierTypeInfo:get_processInfo( ... )
   return self.srcTypeInfo:get_processInfo( ... )
end

function ModifierTypeInfo:get_rawTxt( ... )
   return self.srcTypeInfo:get_rawTxt( ... )
end

function ModifierTypeInfo:get_retTypeInfoList( ... )
   return self.srcTypeInfo:get_retTypeInfoList( ... )
end

function ModifierTypeInfo:get_scope( ... )
   return self.srcTypeInfo:get_scope( ... )
end

function ModifierTypeInfo:get_staticFlag( ... )
   return self.srcTypeInfo:get_staticFlag( ... )
end

function ModifierTypeInfo:get_typeData( ... )
   return self.srcTypeInfo:get_typeData( ... )
end

function ModifierTypeInfo:hasBase( ... )
   return self.srcTypeInfo:hasBase( ... )
end

function ModifierTypeInfo:hasRouteNamespaceFrom( ... )
   return self.srcTypeInfo:hasRouteNamespaceFrom( ... )
end

function ModifierTypeInfo:isInheritFrom( ... )
   return self.srcTypeInfo:isInheritFrom( ... )
end

function ModifierTypeInfo:isModule( ... )
   return self.srcTypeInfo:isModule( ... )
end

function ModifierTypeInfo:serializeTypeInfoList( ... )
   return self.srcTypeInfo:serializeTypeInfoList( ... )
end

function ModifierTypeInfo:switchScope( ... )
   return self.srcTypeInfo:switchScope( ... )
end



local ModuleTypeInfo = {}
setmetatable( ModuleTypeInfo, { __index = TypeInfo } )
_moduleObj.ModuleTypeInfo = ModuleTypeInfo
function ModuleTypeInfo.new( processInfo, scope, externalFlag, txt, parentInfo, typeId, mutable )
   local obj = {}
   ModuleTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, scope, externalFlag, txt, parentInfo, typeId, mutable ); end
   return obj
end
function ModuleTypeInfo:__init(processInfo, scope, externalFlag, txt, parentInfo, typeId, mutable) 
   TypeInfo.__init( self,scope, processInfo)
   
   
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
   
   
   processInfo:get_idProv():increment(  )
   
   scope:set_ownerTypeInfo( self )
end
function ModuleTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
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
function ModuleTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function ModuleTypeInfo:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag )

   return rawTxt
end
function ModuleTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function ModuleTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function ModuleTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return false, nil
end
function ModuleTypeInfo:serialize( stream, serializeInfo )

   local txt = string.format( "{ skind = %d, parentId = %d, typeId = %d, txt = '%s', ", SerializeKind.Module, self:getParentId(  ).id, self.typeId.id, self.rawTxt)
   stream:write( txt .. '\n' )
   stream:write( "}\n" )
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
   
end
_moduleObj.getEnumLiteralVal = getEnumLiteralVal

local EnumValInfo = {}
_moduleObj.EnumValInfo = EnumValInfo
function EnumValInfo.setmeta( obj )
  setmetatable( obj, { __index = EnumValInfo  } )
end
function EnumValInfo.new( name, val, symbolInfo )
   local obj = {}
   EnumValInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, val, symbolInfo )
   end
   return obj
end
function EnumValInfo:__init( name, val, symbolInfo )

   self.name = name
   self.val = val
   self.symbolInfo = symbolInfo
end
function EnumValInfo:get_name()
   return self.name
end
function EnumValInfo:get_val()
   return self.val
end
function EnumValInfo:get_symbolInfo()
   return self.symbolInfo
end


local EnumTypeInfo = {}
setmetatable( EnumTypeInfo, { __index = TypeInfo } )
_moduleObj.EnumTypeInfo = EnumTypeInfo
function EnumTypeInfo.new( processInfo, scope, externalFlag, accessMode, txt, parentInfo, valTypeInfo )
   local obj = {}
   EnumTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, scope, externalFlag, accessMode, txt, parentInfo, valTypeInfo ); end
   return obj
end
function EnumTypeInfo:__init(processInfo, scope, externalFlag, accessMode, txt, parentInfo, valTypeInfo) 
   TypeInfo.__init( self,scope, processInfo)
   
   
   self.externalFlag = externalFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = processInfo:newId(  )
   self.name2EnumValInfo = {}
   self.valTypeInfo = valTypeInfo
   
   self.val2EnumValInfo = {}
   
   do
      local _exp = parentInfo
      if _exp ~= nil then
         _exp:addChildren( self )
      end
   end
   
   
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
   
   scope:set_ownerTypeInfo( self )
end
function EnumTypeInfo:isModule(  )

   return false
end
function EnumTypeInfo:get_kind(  )

   return TypeInfoKind.Enum
end
function EnumTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function EnumTypeInfo:getParentId(  )

   return self.parentInfo:get_typeId()
end
function EnumTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function EnumTypeInfo:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag )

   return rawTxt
end
function EnumTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function EnumTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function EnumTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   if self == other:get_srcTypeInfo():get_aliasSrc() then
      return true, nil
   end
   
   return false, string.format( "%d != %d", self:get_typeId().id, other:get_srcTypeInfo():get_aliasSrc():get_typeId().id)
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

local AlgeTypeInfo = {}
setmetatable( AlgeTypeInfo, { __index = TypeInfo } )
_moduleObj.AlgeTypeInfo = AlgeTypeInfo
function AlgeTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function AlgeTypeInfo.new( processInfo, scope, externalFlag, accessMode, txt, parentInfo )
   local obj = {}
   AlgeTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, scope, externalFlag, accessMode, txt, parentInfo ); end
   return obj
end
function AlgeTypeInfo:__init(processInfo, scope, externalFlag, accessMode, txt, parentInfo) 
   TypeInfo.__init( self,scope, processInfo)
   
   
   self.externalFlag = externalFlag
   self.accessMode = accessMode
   self.rawTxt = txt
   self.parentInfo = _lune.unwrapDefault( parentInfo, _moduleObj.headTypeInfo)
   self.typeId = processInfo:newId(  )
   self.valInfoMap = {}
   self.valInfoNum = 0
   
   do
      local _exp = parentInfo
      if _exp ~= nil then
         _exp:addChildren( self )
      end
   end
   
   
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
   
   scope:set_ownerTypeInfo( self )
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
function AlgeTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function AlgeTypeInfo:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag )

   return rawTxt
end
function AlgeTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function AlgeTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function AlgeTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return self == other:get_srcTypeInfo():get_aliasSrc(), nil
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
function AlgeTypeInfo:get_valInfoNum()
   return self.valInfoNum
end


_moduleObj.AlgeValInfo = AlgeValInfo
function AlgeValInfo:serialize( stream )

   stream:write( string.format( "{ name = '%s', typeList = {", self.name) )
   for index, typeInfo in ipairs( self.typeList ) do
      if index > 1 then
         stream:write( ", " )
      end
      
      stream:write( string.format( "%d", typeInfo:get_typeId().id) )
   end
   
   stream:write( "} }" )
end
function AlgeValInfo.setmeta( obj )
  setmetatable( obj, { __index = AlgeValInfo  } )
end
function AlgeValInfo.new( name, typeList, algeTpye, symbolInfo )
   local obj = {}
   AlgeValInfo.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeList, algeTpye, symbolInfo )
   end
   return obj
end
function AlgeValInfo:__init( name, typeList, algeTpye, symbolInfo )

   self.name = name
   self.typeList = typeList
   self.algeTpye = algeTpye
   self.symbolInfo = symbolInfo
end
function AlgeValInfo:get_name()
   return self.name
end
function AlgeValInfo:get_typeList()
   return self.typeList
end
function AlgeValInfo:get_algeTpye()
   return self.algeTpye
end
function AlgeValInfo:get_symbolInfo()
   return self.symbolInfo
end


function AlgeTypeInfo:addValInfo( valInfo )

   self.valInfoMap[valInfo:get_name()] = valInfo
   self.valInfoNum = self.valInfoNum + 1
end


local OverridingType = {}
OverridingType._name2Val = {}
function OverridingType:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "OverridingType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function OverridingType._from( val )
   return _lune._AlgeFrom( OverridingType, val )
end

OverridingType.NoReady = { "NoReady"}
OverridingType._name2Val["NoReady"] = OverridingType.NoReady
OverridingType.NotOverride = { "NotOverride"}
OverridingType._name2Val["NotOverride"] = OverridingType.NotOverride
OverridingType.Override = { "Override", {{}}}
OverridingType._name2Val["Override"] = OverridingType.Override


local NormalTypeInfo = {}
setmetatable( NormalTypeInfo, { __index = TypeInfo } )
_moduleObj.NormalTypeInfo = NormalTypeInfo
function NormalTypeInfo:getOverridingType(  )

   do
      local _matchExp = self.overridingType
      if _matchExp[1] == OverridingType.NotOverride[1] then
      
         return nil
      elseif _matchExp[1] == OverridingType.Override[1] then
         local typeInfo = _matchExp[2][1]
      
         return typeInfo
      elseif _matchExp[1] == OverridingType.NoReady[1] then
      
         local scope = _lune.unwrap( self.parentInfo:get_scope())
         do
            local typeInfo = scope:getTypeInfoField( self.rawTxt, false, scope, ScopeAccess.Normal )
            if typeInfo ~= nil then
               do
                  local workType = typeInfo:getOverridingType(  )
                  if workType ~= nil then
                     self.overridingType = _lune.newAlge( OverridingType.Override, {workType})
                     return workType
                  else
                     self.overridingType = _lune.newAlge( OverridingType.Override, {typeInfo})
                     return typeInfo
                  end
               end
               
            else
               self.overridingType = _lune.newAlge( OverridingType.NotOverride)
               return nil
            end
         end
         
      end
   end
   
end
function NormalTypeInfo:switchScopeTo( scope )

   self:switchScope( scope )
end
function NormalTypeInfo.new( processInfo, abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode, moduleLang )
   local obj = {}
   NormalTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode, moduleLang ); end
   return obj
end
function NormalTypeInfo:__init(processInfo, abstractFlag, scope, baseTypeInfo, interfaceList, autoFlag, externalFlag, staticFlag, accessMode, txt, parentInfo, typeId, kind, itemTypeInfoList, argTypeInfoList, retTypeInfoList, mutMode, moduleLang) 
   TypeInfo.__init( self,scope, processInfo)
   
   
   if type( kind ) ~= "number" then
      Util.printStackTrace(  )
   end
   
   
   if kind == TypeInfoKind.Method and _lune.nilacc( parentInfo, 'hasBase', 'callmtd'  ) then
      self.overridingType = _lune.newAlge( OverridingType.NoReady)
   else
    
      self.overridingType = _lune.newAlge( OverridingType.NotOverride)
   end
   
   
   self.requirePath = ""
   self.moduleLang = moduleLang
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
            
            for index, appyType in ipairs( self.itemTypeInfoList ) do
               local genType = self.baseTypeInfo:get_itemTypeInfoList()[index]
               alt2typeMap[genType] = appyType
            end
            
         elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
            for __index, ifType in ipairs( self.interfaceList ) do
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
         elseif _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method or _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.FormFunc then
            hasNilable = true
         end
      end
      
      if hasNilable then
         self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
      else
       
         self.nilableTypeInfo = _moduleObj.headTypeInfo
      end
      
      processInfo:get_idProv():increment(  )
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

   return self.parentInfo:get_typeId()
end
function NormalTypeInfo:get_baseId(  )

   return self.baseTypeInfo:get_typeId()
end
function NormalTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function NormalTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   local parentTxt = ""
   if typeNameCtrl ~= nil then
      parentTxt = self:getParentFullName( typeNameCtrl, importInfo, localFlag )
   end
   
   local name
   
   if #self.itemTypeInfoList > 0 then
      local txt = raw .. "<"
      for index, typeInfo in ipairs( self.itemTypeInfoList ) do
         if index ~= 1 then
            txt = txt .. ","
         end
         
         txt = txt .. typeInfo:getTxt( typeNameCtrl, importInfo, localFlag )
      end
      
      
      name = parentTxt .. txt .. ">"
   else
    
      name = parentTxt .. raw
   end
   
   
   return name
end
function NormalTypeInfo:get_display_stirng_with( raw, alt2type )

   do
      local _switchExp = self.kind
      if _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Method or _switchExp == TypeInfoKind.Macro then
         local txt = raw .. "("
         for index, argType in ipairs( self.argTypeInfoList ) do
            if index ~= 1 then
               txt = txt .. ", "
            end
            
            txt = txt .. argType:get_display_stirng(  )
         end
         
         txt = txt .. ")"
         for index, retType in ipairs( self.retTypeInfoList ) do
            if index == 1 then
               txt = txt .. ": "
            else
             
               txt = txt .. ", "
            end
            
            txt = txt .. retType:get_display_stirng(  )
         end
         
         return txt
      end
   end
   
   
   local parentTxt = ""
   local name
   
   if #self.itemTypeInfoList > 0 then
      local txt = raw .. "<"
      for index, typeInfo in ipairs( self.itemTypeInfoList ) do
         if index ~= 1 then
            txt = txt .. ","
         end
         
         txt = txt .. typeInfo:get_display_stirng_with( typeInfo:get_rawTxt(), alt2type )
      end
      
      
      name = parentTxt .. txt .. ">"
   else
    
      name = parentTxt .. raw
   end
   
   
   return name
end
function NormalTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function NormalTypeInfo:serialize( stream, serializeInfo )

   if self.typeId.id == _moduleObj.rootTypeId then
      return 
   end
   
   
   local parentId = self:getParentId(  )
   
   local txt = string.format( [==[{ skind=%d, parentId = %d, typeId = %d, baseId = %d, txt = '%s',
        abstractFlag = %s, staticFlag = %s, accessMode = %d, kind = %d, mutMode = %d, ]==], SerializeKind.Normal, parentId.id, self.typeId.id, self:get_baseId().id, self.rawTxt, self.abstractFlag, self.staticFlag, self.accessMode, self.kind, self.mutMode)
   do
      local _exp = self.moduleLang
      if _exp ~= nil then
         txt = txt .. string.format( 'moduleLang = %d, ', _exp)
      end
   end
   
   if self.requirePath ~= "" then
      txt = txt .. string.format( 'requirePath = "%s", ', self.requirePath)
   end
   
   
   local children = {}
   for __index, child in ipairs( self:get_children() ) do
      if serializeInfo:isValidChildren( child:get_typeId() ) then
         table.insert( children, child )
      end
      
   end
   
   
   stream:write( txt .. self:serializeTypeInfoList( "itemTypeId = {", self.itemTypeInfoList ) .. self:serializeTypeInfoList( "ifList = {", self.interfaceList ) .. self:serializeTypeInfoList( "argTypeId = {", self.argTypeInfoList ) .. self:serializeTypeInfoList( "retTypeId = {", self.retTypeInfoList ) .. self:serializeTypeInfoList( "children = {", children, true ) .. "}\n" )
end
function NormalTypeInfo:equalsSub( processInfo, typeInfo, alt2type, checkModifer )

   if self.typeId == typeInfo:get_typeId() then
      return true
   end
   
   
   if typeInfo:get_kind() == TypeInfoKind.Alternate then
      return typeInfo:equals( processInfo, self, alt2type, checkModifer )
   end
   
   
   do
      local aliasType = _lune.__Cast( typeInfo, 3, AliasTypeInfo )
      if aliasType ~= nil then
         return aliasType:equals( processInfo, self, alt2type, checkModifer )
      end
   end
   
   
   if self.kind ~= typeInfo:get_kind() or self.staticFlag ~= typeInfo:get_staticFlag() or self.autoFlag ~= typeInfo:get_autoFlag() or self:get_nilable() ~= typeInfo:get_nilable() or self.rawTxt ~= typeInfo:get_rawTxt() or self.parentInfo ~= typeInfo:get_parentInfo() or self.baseTypeInfo ~= typeInfo:get_baseTypeInfo() then
      return false
   end
   
   
   if self.accessMode ~= typeInfo:get_accessMode() then
      do
         local _switchExp = self.kind
         if _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Map or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set then
            
         else 
            
               return false
         end
      end
      
   end
   
   
   do
      if #self.itemTypeInfoList ~= #typeInfo:get_itemTypeInfoList() then
         return false
      end
      
      for index, item in ipairs( self.itemTypeInfoList ) do
         if not item:equals( processInfo, typeInfo:get_itemTypeInfoList()[index], alt2type, checkModifer ) then
            return false
         end
         
      end
      
   end
   
   
   do
      if #self.retTypeInfoList ~= #typeInfo:get_retTypeInfoList() then
         return false
      end
      
      for index, item in ipairs( self.retTypeInfoList ) do
         if not item:equals( processInfo, typeInfo:get_retTypeInfoList()[index], alt2type, checkModifer ) then
            return false
         end
         
      end
      
   end
   
   
   return true
end
function NormalTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   return self:equalsSub( processInfo, typeInfo, alt2type, checkModifer )
end
function NormalTypeInfo.create( processInfo, accessMode, abstractFlag, scope, baseInfo, parentInfo, staticFlag, kind, txt, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode )

   do
      local _switchExp = kind
      if _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.ExtModule or _switchExp == TypeInfoKind.IF then
         Util.err( "can't use create() method. use createClass(), createExtModule()." )
      end
   end
   
   if kind == TypeInfoKind.Prim then
      do
         local _exp = sym2builtInTypeMap[txt]
         if _exp ~= nil then
            return _exp:get_typeInfo()
         end
      end
      
      Util.err( string.format( "not found symbol -- %s", txt) )
   end
   
   local info = NormalTypeInfo.new(processInfo, abstractFlag, scope, baseInfo, nil, false, true, staticFlag, accessMode, txt, parentInfo, processInfo:newId(  ), kind, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode, nil)
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
function NormalTypeInfo:set_mutMode( mutMode )
   self.mutMode = mutMode
end
function NormalTypeInfo:get_requirePath()
   return self.requirePath
end
function NormalTypeInfo:set_requirePath( requirePath )
   self.requirePath = requirePath
end


function ProcessInfo:createAlternate( belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList )

   return AlternateTypeInfo.new(self, belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList)
end


local DDDTypeInfo = {}

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
   self.ExtDDDMap = {}
   self.ExtMap = {}
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
   
   
   
   for key, val in pairs( self.ExtDDDMap ) do
      obj.ExtDDDMap[key] = val
   end
   
   
   
   for key, val in pairs( self.ExtMap ) do
      obj.ExtMap[key] = val
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
function TypeInfo2Map:get_ExtDDDMap()
   return self.ExtDDDMap
end
function TypeInfo2Map:get_ExtMap()
   return self.ExtMap
end


rootProcessInfo:set_typeInfo2Map( TypeInfo2Map.new() )
local function isExtType( typeInfo )

   return typeInfo:get_kind() == TypeInfoKind.Ext or (typeInfo:get_kind() == TypeInfoKind.DDD and typeInfo:get_extedType() ~= typeInfo )
end
_moduleObj.isExtType = isExtType

local immutableTypeSet = {}
local function isMutableType( typeInfo )

   typeInfo = typeInfo:get_nonnilableType()
   if _lune._Set_has(immutableTypeSet, typeInfo ) then
      return false
   end
   
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Enum then
         return false
      end
   end
   
   return true
end
_moduleObj.isMutableType = isMutableType

function ProcessInfo:createModifier( srcTypeInfo, mutMode )

   srcTypeInfo = srcTypeInfo:get_srcTypeInfo()
   if not isMutableType( srcTypeInfo ) then
      return srcTypeInfo
   end
   
   
   do
      local _switchExp = mutMode
      if _switchExp == MutMode.IMut or _switchExp == MutMode.IMutRe then
         do
            local _exp = self:get_typeInfo2Map().ImutModifierMap[srcTypeInfo]
            if _exp ~= nil then
               return _exp
            end
         end
         
      elseif _switchExp == MutMode.AllMut then
         do
            local _exp = self:get_typeInfo2Map().MutModifierMap[srcTypeInfo]
            if _exp ~= nil then
               return _exp
            end
         end
         
      end
   end
   
   
   local modifier
   
   if srcTypeInfo:get_nonnilableType():get_kind() == TypeInfoKind.Ext then
      do
         local _matchExp = self:createLuaval( self:createModifier( srcTypeInfo:get_extedType(), mutMode ), false )
         if _matchExp[1] == LuavalResult.OK[1] then
            local workType = _matchExp[2][1]
            local _ = _matchExp[2][2]
         
            if srcTypeInfo:get_nilable() then
               modifier = workType:get_nilableTypeInfo()
            else
             
               modifier = workType
            end
            
         elseif _matchExp[1] == LuavalResult.Err[1] then
            local err = _matchExp[2][1]
         
            Util.err( err )
         end
      end
      
   else
    
      modifier = ModifierTypeInfo.new(nil, self, srcTypeInfo, self:newId(  ), mutMode)
   end
   
   do
      local _switchExp = mutMode
      if _switchExp == MutMode.IMut or _switchExp == MutMode.IMutRe then
         self:get_typeInfo2Map().ImutModifierMap[srcTypeInfo] = modifier
      elseif _switchExp == MutMode.AllMut then
         self:get_typeInfo2Map().MutModifierMap[srcTypeInfo] = modifier
      end
   end
   
   return modifier
end


function Scope:addOverrideImut( symbolInfo )

   local typeInfo
   
   if TypeInfo.isMut( symbolInfo:get_typeInfo() ) then
      local processInfo = self:getProcessInfo(  )
      typeInfo = processInfo:createModifier( symbolInfo:get_typeInfo(), MutMode.IMut )
   else
    
      typeInfo = symbolInfo:get_typeInfo()
   end
   
   
   self.symbol2SymbolInfoMap[symbolInfo:get_name()] = AccessSymbolInfo.new(symbolInfo, _lune.newAlge( OverrideMut.IMut, {typeInfo}), false)
end


rootProcessInfo:get_idProv():increment(  )

local function addBuiltin( typeInfo )

   builtInTypeIdSet[typeInfo:get_typeId().id] = typeInfo
end
_moduleObj.addBuiltin = addBuiltin
addBuiltin( _moduleObj.headTypeInfo )

local function registBuiltin( idName, typeTxt, kind, typeInfo, nilableTypeInfo, registScope )

   sym2builtInTypeMap[typeTxt] = NormalSymbolInfo.new(rootProcessInfo, SymbolKind.Typ, false, false, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, nil, typeInfo, MutMode.IMut, true, false)
   if nilableTypeInfo ~= _moduleObj.headTypeInfo then
      sym2builtInTypeMap[typeTxt .. "!"] = NormalSymbolInfo.new(rootProcessInfo, SymbolKind.Typ, false, kind == TypeInfoKind.Func, _moduleObj.rootScope, AccessMode.Pub, false, typeTxt, nil, nilableTypeInfo, MutMode.IMut, true, false)
   end
   
   addBuiltin( typeInfo )
   addBuiltin( rootProcessInfo:createModifier( typeInfo, MutMode.IMut ) )
   
   if typeInfo:get_nilableTypeInfo() ~= _moduleObj.headTypeInfo then
      addBuiltin( typeInfo:get_nilableTypeInfo() )
      addBuiltin( rootProcessInfo:createModifier( typeInfo:get_nilableTypeInfo(), MutMode.IMut ) )
   end
   
   
   if registScope then
      _moduleObj.rootScope:addClass( rootProcessInfo, typeTxt, nil, typeInfo )
   end
   
end

function NormalTypeInfo.createBuiltin( idName, typeTxt, kind, typeDDD, ifList )

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
      if _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Set or _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.Module or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Method or _switchExp == TypeInfoKind.Macro then
         scope = Scope.new(rootProcessInfo, _moduleObj.rootScope, kind == TypeInfoKind.Class or kind == TypeInfoKind.Module or kind == TypeInfoKind.IF or kind == TypeInfoKind.List or kind == TypeInfoKind.Array or kind == TypeInfoKind.Set, nil)
      end
   end
   
   
   local genTypeList = {}
   do
      local _switchExp = kind
      if _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Set then
         table.insert( genTypeList, rootProcessInfo:createAlternate( true, 1, "T", AccessMode.Pri, _moduleObj.headTypeInfo ) )
      elseif _switchExp == TypeInfoKind.Map then
         table.insert( genTypeList, rootProcessInfo:createAlternate( true, 1, "K", AccessMode.Pri, _moduleObj.headTypeInfo ) )
         table.insert( genTypeList, rootProcessInfo:createAlternate( true, 2, "V", AccessMode.Pri, _moduleObj.headTypeInfo ) )
      end
   end
   
   local info = NormalTypeInfo.new(rootProcessInfo, false, scope, nil, ifList, false, false, false, AccessMode.Pub, typeTxt, _moduleObj.headTypeInfo, rootProcessInfo:newId(  ), kind, genTypeList, argTypeList, retTypeList, MutMode.Mut, nil)
   
   registBuiltin( idName, typeTxt, kind, info, _moduleObj.headTypeInfo, scope ~= nil )
   return info
end


local builtinTypeNone = NormalTypeInfo.createBuiltin( "__None", "", TypeInfoKind.Prim )
_moduleObj.builtinTypeNone = builtinTypeNone

local builtinTypeEmpty = NormalTypeInfo.createBuiltin( "__Empty", "::", TypeInfoKind.Prim )
_moduleObj.builtinTypeEmpty = builtinTypeEmpty

local builtinTypeNeverRet = NormalTypeInfo.createBuiltin( "__NRet", "__", TypeInfoKind.Prim )
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

local builtinTypeChar = NormalTypeInfo.createBuiltin( "char", "__char", TypeInfoKind.Prim )
_moduleObj.builtinTypeChar = builtinTypeChar

local builtinTypeMapping = NormalTypeInfo.createBuiltin( "Mapping", "Mapping", TypeInfoKind.IF )
_moduleObj.builtinTypeMapping = builtinTypeMapping

local builtinTypeAsyncItem = NormalTypeInfo.createBuiltin( "__AsyncItem", "__AsyncItem", TypeInfoKind.IF )
_moduleObj.builtinTypeAsyncItem = builtinTypeAsyncItem

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


immutableTypeSet[_moduleObj.builtinTypeBool]= true
immutableTypeSet[_moduleObj.builtinTypeInt]= true
immutableTypeSet[_moduleObj.builtinTypeReal]= true
immutableTypeSet[_moduleObj.builtinTypeChar]= true
immutableTypeSet[_moduleObj.builtinTypeString]= true

local function isClass( typeInfo )

   return typeInfo:get_kind() == TypeInfoKind.Class and typeInfo ~= _moduleObj.builtinTypeString
end
_moduleObj.isClass = isClass

function Scope:addIgnoredVar( processInfo )

   self:addLocalVar( processInfo, false, true, "_", nil, _moduleObj.builtinTypeEmpty, MutMode.Mut )
end


function AlternateTypeInfo:canSetFrom( processInfo, other, canEvalType, alt2type )

   local otherWork = AlternateTypeInfo.getAssign( other, alt2type )
   if self == otherWork then
      return true
   end
   
   
   do
      local genType = alt2type[self]
      if genType ~= nil then
         if canEvalType ~= nil then
            return (genType:canEvalWith( processInfo, otherWork, canEvalType, alt2type ) )
         end
         
         return genType:equals( processInfo, otherWork, alt2type )
      end
   end
   
   local workAlt2type
   
   if not CanEvalCtrlTypeInfo.isValidApply( alt2type ) then
      if not isClass( otherWork ) and otherWork:get_kind() ~= TypeInfoKind.IF then
         return false
      end
      
      workAlt2type = CanEvalCtrlTypeInfo.createDefaultAlt2typeMap( false )
   else
    
      workAlt2type = alt2type
   end
   
   
   if self:hasBase(  ) then
      if not other:isInheritFrom( processInfo, self.baseTypeInfo, workAlt2type ) then
         return false
      end
      
   end
   
   
   for __index, ifType in ipairs( self.interfaceList ) do
      if not other:isInheritFrom( processInfo, ifType, workAlt2type ) then
         return false
      end
      
   end
   
   
   workAlt2type[self] = otherWork
   return true
end


function NormalSymbolInfo:set_typeInfo( typeInfo )

   if self.name == "_" then
      return 
   end
   
   self.typeInfo = typeInfo
end


local LuavalConvKind = {}
LuavalConvKind._val2NameMap = {}
function LuavalConvKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "LuavalConvKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function LuavalConvKind._from( val )
   if LuavalConvKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
LuavalConvKind.__allList = {}
function LuavalConvKind.get__allList()
   return LuavalConvKind.__allList
end

LuavalConvKind.InLua = 0
LuavalConvKind._val2NameMap[0] = 'InLua'
LuavalConvKind.__allList[1] = LuavalConvKind.InLua
LuavalConvKind.ToLua = 1
LuavalConvKind._val2NameMap[1] = 'ToLua'
LuavalConvKind.__allList[2] = LuavalConvKind.ToLua


local function failCreateLuavalWith( typeInfo, convFlag, validToCheck )

   
   if isExtType( typeInfo ) then
      return nil, true
   end
   
   local mess = string.format( "not support to use the type as Luaval -- %s", typeInfo:getTxt(  ))
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == TypeInfoKind.Nilable then
         return failCreateLuavalWith( typeInfo:get_nonnilableType(), convFlag, validToCheck )
      elseif _switchExp == TypeInfoKind.Prim then
         return nil, true
      elseif _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.DDD or _switchExp == TypeInfoKind.ExtModule then
         if not validToCheck then
            return nil, false
         end
         
         if convFlag ~= LuavalConvKind.InLua then
            return mess, false
         end
         
         return nil, false
      elseif _switchExp == TypeInfoKind.Stem then
         return nil, false
      elseif _switchExp == TypeInfoKind.Class then
         if typeInfo ~= _moduleObj.builtinTypeString then
            if not validToCheck then
               return nil, false
            end
            
            if convFlag ~= LuavalConvKind.InLua then
               return mess, false
            end
            
            return nil, false
         end
         
         return nil, true
      elseif _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Map then
         if not validToCheck then
            return nil, false
         end
         
         if convFlag ~= LuavalConvKind.ToLua and isMutable( typeInfo:get_mutMode() ) then
            return "not support mutable collecion. " .. mess, false
         end
         
         local canConv = true
         
         for __index, itemType in ipairs( typeInfo:get_itemTypeInfoList() ) do
            local err, work = failCreateLuavalWith( itemType, convFlag, validToCheck )
            if err ~= nil then
               return err, false
            end
            
            if not work then
               canConv = false
            end
            
         end
         
         
         canConv = false
         return nil, canConv
      elseif _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Func then
         if not validToCheck then
            return nil, false
         end
         
         if convFlag ~= LuavalConvKind.InLua then
            return mess, false
         end
         
         if #typeInfo:get_itemTypeInfoList() ~= 0 then
            return mess, false
         end
         
         local canConv = true
         
         for __index, itemType in ipairs( typeInfo:get_argTypeInfoList() ) do
            local err, work = failCreateLuavalWith( itemType, convFlag, validToCheck )
            if err ~= nil then
               return err, false
            end
            
            if not work then
               canConv = false
            end
            
         end
         
         
         
         for __index, itemType in ipairs( typeInfo:get_retTypeInfoList() ) do
            local err, work = failCreateLuavalWith( itemType, convFlag, validToCheck )
            if err ~= nil then
               return err, false
            end
            
            if not work then
               canConv = false
            end
            
         end
         
         
         canConv = false
         return nil, canConv
      end
   end
   
   return string.format( "not support -- %s:%s", typeInfo:getTxt(  ), TypeInfoKind:_getTxt( typeInfo:get_kind())
   ), false
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
            if _exp ~= otherWork then
               otherWork = _exp
            else
             
               return otherWork
            end
            
         else
            return otherWork
         end
      end
      
   end
   
end


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

local builtinTypeBox = BoxTypeInfo.new(rootProcessInfo, rootProcessInfo:newId(  ), AccessMode.Pub, boxRootAltType)
_moduleObj.builtinTypeBox = builtinTypeBox

registBuiltin( "Nilable", "Nilable", TypeInfoKind.Box, _moduleObj.builtinTypeBox, _moduleObj.headTypeInfo, true )

local function isConditionalbe( processInfo, typeInfo )

   if typeInfo:get_nilable() or typeInfo:equals( processInfo, _moduleObj.builtinTypeBool, nil ) then
      return true
   end
   
   return false
end
_moduleObj.isConditionalbe = isConditionalbe

function ProcessInfo:createBox( accessMode, nonnilableType )

   do
      local boxType = self:get_typeInfo2Map().BoxMap[nonnilableType]
      if boxType ~= nil then
         return boxType
      end
   end
   
   
   local boxType = BoxTypeInfo.new(self, self:newId(  ), accessMode, nonnilableType)
   self:get_typeInfo2Map().BoxMap[nonnilableType] = boxType
   return boxType
end


function BoxTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local typeInfo = self.boxingType:applyGeneric( alt2typeMap, moduleTypeInfo )
   if typeInfo == self.boxingType then
      return self
   end
   
   if typeInfo ~= nil then
      return moduleTypeInfo:getProcessInfo(  ):createBox( self.accessMode, typeInfo )
   end
   
   return nil
end




function ProcessInfo:createSet( accessMode, parentInfo, itemTypeInfo, mutMode )

   
   local tmpMutMode
   
   if isMutable( mutMode ) then
      tmpMutMode = mutMode
   else
    
      tmpMutMode = MutMode.Mut
   end
   
   local function newTypeFunc( workMutMode )
   
      return NormalTypeInfo.new(self, false, getScope( _moduleObj.builtinTypeSet ), _moduleObj.builtinTypeSet, nil, false, false, false, AccessMode.Pub, "Set", _moduleObj.headTypeInfo, self:newId(  ), TypeInfoKind.Set, itemTypeInfo, nil, nil, workMutMode, nil)
   end
   
   
   local typeInfo = newTypeFunc( tmpMutMode )
   
   if isMutable( mutMode ) then
      return typeInfo
   end
   
   return self:createModifier( typeInfo, mutMode )
   
end


function ProcessInfo:createList( accessMode, parentInfo, itemTypeInfo, mutMode )

   
   local tmpMutMode
   
   if isMutable( mutMode ) then
      tmpMutMode = mutMode
   else
    
      tmpMutMode = MutMode.Mut
   end
   
   local function newTypeFunc( workMutMode )
   
      return NormalTypeInfo.new(self, false, getScope( _moduleObj.builtinTypeList ), _moduleObj.builtinTypeList, nil, false, false, false, AccessMode.Pub, "List", _moduleObj.headTypeInfo, self:newId(  ), TypeInfoKind.List, itemTypeInfo, nil, nil, workMutMode, nil)
   end
   
   
   local typeInfo = newTypeFunc( tmpMutMode )
   
   if isMutable( mutMode ) then
      return typeInfo
   end
   
   return self:createModifier( typeInfo, mutMode )
   
end


function ProcessInfo:createArray( accessMode, parentInfo, itemTypeInfo, mutMode )

   
   local tmpMutMode
   
   if isMutable( mutMode ) then
      tmpMutMode = mutMode
   else
    
      tmpMutMode = MutMode.Mut
   end
   
   local function newTypeFunc( workMutMode )
   
      return NormalTypeInfo.new(self, false, getScope( _moduleObj.builtinTypeArray ), _moduleObj.builtinTypeArray, nil, false, false, false, AccessMode.Pub, "Array", _moduleObj.headTypeInfo, self:newId(  ), TypeInfoKind.Array, itemTypeInfo, nil, nil, workMutMode, nil)
   end
   
   
   local typeInfo = newTypeFunc( tmpMutMode )
   
   if isMutable( mutMode ) then
      return typeInfo
   end
   
   return self:createModifier( typeInfo, mutMode )
   
end


function ProcessInfo:createMap( accessMode, parentInfo, keyTypeInfo, valTypeInfo, mutMode )

   
   local tmpMutMode
   
   if isMutable( mutMode ) then
      tmpMutMode = mutMode
   else
    
      tmpMutMode = MutMode.Mut
   end
   
   local function newTypeFunc( workMutMode )
   
      return NormalTypeInfo.new(self, false, getScope( _moduleObj.builtinTypeMap ), _moduleObj.builtinTypeMap, nil, false, false, false, AccessMode.Pub, "Map", _moduleObj.headTypeInfo, self:newId(  ), TypeInfoKind.Map, {keyTypeInfo, valTypeInfo}, nil, nil, workMutMode, nil)
   end
   
   
   local typeInfo = newTypeFunc( tmpMutMode )
   
   if isMutable( mutMode ) then
      return typeInfo
   end
   
   return self:createModifier( typeInfo, mutMode )
   
end


function ProcessInfo:createModule( scope, parentInfo, externalFlag, moduleName, mutable )

   do
      local _exp = sym2builtInTypeMap[moduleName]
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   
   if Parser.isLuaKeyword( moduleName ) then
      Util.err( string.format( "This symbol can not use for a class or script file. -- %s", moduleName) )
   end
   
   
   local info = ModuleTypeInfo.new(self, scope, externalFlag, moduleName, parentInfo, self:newId(  ), mutable)
   return info
end


function ProcessInfo:createClass( classFlag, abstractFlag, scope, baseInfo, interfaceList, genTypeList, parentInfo, externalFlag, accessMode, className )

   do
      local _exp = sym2builtInTypeMap[className]
      if _exp ~= nil then
         return _exp:get_typeInfo()
         
      end
   end
   
   
   if Parser.isLuaKeyword( className ) then
      Util.err( string.format( "This symbol can not use for a class or script file. -- %s", className) )
   end
   
   
   local info = NormalTypeInfo.new(self, abstractFlag, scope, baseInfo, interfaceList, false, externalFlag, false, accessMode, className, parentInfo, self:newId(  ), classFlag and TypeInfoKind.Class or TypeInfoKind.IF, genTypeList, nil, nil, MutMode.Mut, nil)
   
   for __index, genType in ipairs( genTypeList ) do
      genType:updateParentInfo( info )
   end
   
   
   return info
end


function ProcessInfo:createExtModule( scope, parentInfo, externalFlag, accessMode, className, moduleLang, requirePath )

   do
      local _exp = sym2builtInTypeMap[className]
      if _exp ~= nil then
         return _exp:get_typeInfo()
      end
   end
   
   
   if Parser.isLuaKeyword( className ) then
      Util.err( string.format( "This symbol can not use for a class or script file. -- %s", className) )
   end
   
   
   local info = NormalTypeInfo.new(self, false, scope, nil, nil, false, externalFlag, false, accessMode, className, parentInfo, self:newId(  ), TypeInfoKind.ExtModule, nil, nil, nil, MutMode.Mut, moduleLang)
   info:set_requirePath( requirePath )
   return info
end


function ProcessInfo:createFunc( abstractFlag, builtinFlag, scope, kind, parentInfo, autoFlag, externalFlag, staticFlag, accessMode, funcName, altTypeList, argTypeList, retTypeInfoList, mutable )

   if not builtinFlag and Parser.isLuaKeyword( funcName ) then
      Util.err( string.format( "This symbol can not use for a function. -- %s", funcName) )
   end
   
   
   local info = NormalTypeInfo.new(self, abstractFlag, scope, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, self:newId(  ), kind, _lune.unwrapDefault( altTypeList, {}), _lune.unwrapDefault( argTypeList, {}), _lune.unwrapDefault( retTypeInfoList, {}), mutable and MutMode.Mut or MutMode.IMut, nil)
   
   if altTypeList ~= nil then
      for __index, genType in ipairs( altTypeList ) do
         do
            local _exp = _lune.__Cast( genType, 3, AlternateTypeInfo )
            if _exp ~= nil then
               _exp:updateParentInfo( info )
            end
         end
         
      end
      
   end
   
   
   return info
end

function ProcessInfo:createAdvertiseMethodFrom( classTypeInfo, typeInfo )

   return self:createFunc( false, false, getScope( typeInfo ), typeInfo:get_kind(), classTypeInfo, true, false, false, typeInfo:get_accessMode(), typeInfo:get_rawTxt(), typeInfo:get_itemTypeInfoList(), typeInfo:get_argTypeInfoList(), typeInfo:get_retTypeInfoList(), TypeInfo.isMut( typeInfo ) )
end


function ModifierTypeInfo:get_nonnilableType(  )

   local orgType = self.srcTypeInfo:get_nonnilableType()
   if TypeInfo.isMut( self ) or not TypeInfo.isMut( orgType ) then
      return orgType
   end
   
   return self:getProcessInfo(  ):createModifier( orgType, MutMode.IMut )
end


function ModifierTypeInfo:get_nilableTypeInfo(  )

   local orgType = self.srcTypeInfo:get_nilableTypeInfo()
   if not TypeInfo.isMut( orgType ) then
      return orgType
   end
   
   return self:getProcessInfo(  ):createModifier( orgType, MutMode.IMut )
end


function ProcessInfo:createAlias( processInfo, name, externalFlag, accessMode, parentInfo, typeInfo )

   return AliasTypeInfo.new(processInfo, name, accessMode, parentInfo, typeInfo:get_srcTypeInfo(), externalFlag)
end


function Scope:addAlias( processInfo, name, pos, externalFlag, accessMode, parentInfo, symbolInfo )

   local aliasType = self:getProcessInfo(  ):createAlias( processInfo, name, externalFlag, accessMode, parentInfo, symbolInfo:get_typeInfo():get_srcTypeInfo() )
   return self:add( processInfo, symbolInfo:get_kind(), false, symbolInfo:get_canBeRight(), name, pos, aliasType, accessMode, true, MutMode.IMut, true, false )
end


function Scope:addAliasForType( processInfo, name, pos, typeInfo )

   local skind = SymbolKind.Typ
   local canBeRight = false
   do
      local _switchExp = typeInfo:get_kind()
      if _switchExp == TypeInfoKind.Func then
         skind = SymbolKind.Fun
         canBeRight = true
      elseif _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.FormFunc then
         canBeRight = true
      elseif _switchExp == TypeInfoKind.Macro then
         skind = SymbolKind.Mac
      end
   end
   
   
   return self:add( processInfo, skind, false, canBeRight, name, pos, typeInfo, typeInfo:get_accessMode(), true, MutMode.IMut, true, false )
end


setmetatable( DDDTypeInfo, { __index = TypeInfo } )
_moduleObj.DDDTypeInfo = DDDTypeInfo
function DDDTypeInfo:get_extTypeFlag(  )

   return self.extedType ~= self
end
function DDDTypeInfo:get_scope(  )

   return nil
end
function DDDTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function DDDTypeInfo:get_parentInfo(  )

   return _moduleObj.headTypeInfo
end
function DDDTypeInfo.new( processInfo, typeId, typeInfo, externalFlag, extOrgDDType )
   local obj = {}
   DDDTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, typeId, typeInfo, externalFlag, extOrgDDType ); end
   return obj
end
function DDDTypeInfo:__init(processInfo, typeId, typeInfo, externalFlag, extOrgDDType) 
   TypeInfo.__init( self,nil, processInfo)
   
   self.typeId = typeId
   self.typeInfo = typeInfo
   self.externalFlag = externalFlag
   self.itemTypeInfoList = {self.typeInfo}
   local extOrgType
   
   if extOrgDDType ~= nil then
      extOrgType = extOrgDDType
      processInfo:get_typeInfo2Map().ExtDDDMap[typeInfo] = self
   else
      extOrgType = self
      processInfo:get_typeInfo2Map().DDDMap[typeInfo] = self
   end
   
   self.extedType = extOrgType
end
function DDDTypeInfo:isModule(  )

   return false
end
function DDDTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return self.typeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )
end
function DDDTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind=%d, typeId = %d, itemTypeId = %d, parentId = %d, extTypeFlag = %s }\n', SerializeKind.DDD, self.typeId.id, self.typeInfo:get_typeId().id, _moduleObj.headTypeInfo:get_typeId().id, self:get_extTypeFlag()) )
end
function DDDTypeInfo:get_display_stirng_with( raw, alt2type )

   local txt = self:getTxtWithRaw( raw )
   return txt
end
function DDDTypeInfo:get_display_stirng(  )

   local txt = self:get_display_stirng_with( self:get_rawTxt(), nil )
   return txt
end
function DDDTypeInfo:getModule(  )

   return self:get_typeInfo():getModule(  )
end
function DDDTypeInfo:get_rawTxt(  )

   return self:getTxt(  )
end
function DDDTypeInfo:get_kind(  )

   return TypeInfoKind.DDD
end
function DDDTypeInfo:get_nilable(  )

   return true
end
function DDDTypeInfo:get_nilableTypeInfo(  )

   return self
end
function DDDTypeInfo:get_mutMode(  )

   return self:get_typeInfo():get_mutMode()
end
function DDDTypeInfo:get_aliasSrc(  )

   return self
end
function DDDTypeInfo:get_srcTypeInfo(  )

   return self
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
function DDDTypeInfo:get_extedType()
   return self.extedType
end


function ProcessInfo:createDDD( typeInfo, externalFlag, extTypeFlag )

   if typeInfo:get_kind() == TypeInfoKind.DDD then
      typeInfo = typeInfo:get_itemTypeInfoList()[1]
   end
   
   if not failCreateLuavalWith( typeInfo, LuavalConvKind.InLua, true ) and extTypeFlag then
      extTypeFlag = false
   end
   
   
   if typeInfo:get_nonnilableType():get_kind() ~= TypeInfoKind.Ext and extTypeFlag then
      do
         local _matchExp = self:createLuaval( typeInfo, true )
         if _matchExp[1] == LuavalResult.OK[1] then
            local work = _matchExp[2][1]
            local _ = _matchExp[2][2]
         
            typeInfo = work
         elseif _matchExp[1] == LuavalResult.Err[1] then
            local mess = _matchExp[2][1]
         
            Util.err( mess )
         end
      end
      
   end
   
   
   local dddMap
   
   if extTypeFlag then
      dddMap = self:get_typeInfo2Map().ExtDDDMap
   else
    
      dddMap = self:get_typeInfo2Map().DDDMap
   end
   
   do
      local _exp = dddMap[typeInfo]
      if _exp ~= nil then
         if _exp:get_typeId().id < userStartId and typeInfo:get_typeId().id >= userStartId then
            Util.err( "on cache" )
         end
         
         return _exp
      end
   end
   
   
   local dddType = DDDTypeInfo.new(self, self:newId(  ), typeInfo, externalFlag, nil)
   
   if failCreateLuavalWith( typeInfo, LuavalConvKind.InLua, true ) then
      local extDDDType = DDDTypeInfo.new(self, self:newId(  ), typeInfo, externalFlag, dddType)
      
      if extTypeFlag then
         return extDDDType
      end
      
   end
   
   
   return dddType
end


local builtinTypeNil = NilTypeInfo.new(rootProcessInfo)
_moduleObj.builtinTypeNil = builtinTypeNil

registBuiltin( "Nil", "nil", TypeInfoKind.Prim, _moduleObj.builtinTypeNil, _moduleObj.headTypeInfo, false )

local builtinTypeDDD = rootProcessInfo:createDDD( _moduleObj.builtinTypeStem_, true, false )
_moduleObj.builtinTypeDDD = builtinTypeDDD

registBuiltin( "DDD", "...", TypeInfoKind.DDD, _moduleObj.builtinTypeDDD, _moduleObj.headTypeInfo, false )

local builtinTypeForm = NormalTypeInfo.createBuiltin( "Form", "form", TypeInfoKind.Form, _moduleObj.builtinTypeDDD )
_moduleObj.builtinTypeForm = builtinTypeForm

immutableTypeSet[_moduleObj.builtinTypeForm]= true

local builtinTypeSymbol = NormalTypeInfo.createBuiltin( "Symbol", "sym", TypeInfoKind.Prim )
_moduleObj.builtinTypeSymbol = builtinTypeSymbol

local builtinTypeStat = NormalTypeInfo.createBuiltin( "Stat", "stat", TypeInfoKind.Prim )
_moduleObj.builtinTypeStat = builtinTypeStat

local builtinTypeExp = NormalTypeInfo.createBuiltin( "Exp", "__exp", TypeInfoKind.Prim )
_moduleObj.builtinTypeExp = builtinTypeExp

local builtinTypeMultiExp = NormalTypeInfo.createBuiltin( "Exps", "__exps", TypeInfoKind.Prim )
_moduleObj.builtinTypeMultiExp = builtinTypeMultiExp


local CombineType = {}
_moduleObj.CombineType = CombineType
function CombineType.new( typeInfo )
   local obj = {}
   CombineType.setmeta( obj )
   if obj.__init then obj:__init( typeInfo ); end
   return obj
end
function CombineType:__init(typeInfo) 
   self.ifSet = {}
   for __index, iftype in ipairs( typeInfo:get_interfaceList() ) do
      self.ifSet[iftype]= true
   end
   
   self.nilable = typeInfo:get_nilable()
   self.mutMode = typeInfo:get_mutMode()
end
function CombineType:isInheritFrom( processInfo, other, alt2type )

   for ifType, __val in pairs( self.ifSet ) do
      if ifType:isInheritFrom( processInfo, other, alt2type ) then
         return true
      end
      
   end
   
   return false
end
function CombineType:andIfSet( processInfo, ifSet, alt2type )

   local workSet = {}
   for other, __val in pairs( ifSet ) do
      if self:isInheritFrom( processInfo, other, alt2type ) then
         workSet[other]= true
      else
       
         for ifType, __val in pairs( self.ifSet ) do
            if other:isInheritFrom( processInfo, ifType, alt2type ) then
               workSet[ifType]= true
            end
            
         end
         
      end
      
   end
   
   self.ifSet = workSet
end
function CombineType:createStem( processInfo )

   local retType
   
   if self.nilable then
      retType = _moduleObj.builtinTypeStem_
   else
    
      retType = _moduleObj.builtinTypeStem
   end
   
   if isMutable( self.mutMode ) then
      return retType
   end
   
   return processInfo:createModifier( retType, self.mutMode )
end
function CombineType:get_typeInfo( processInfo )

   if _lune._Set_len(self.ifSet ) ~= 1 then
      return self:createStem( processInfo )
   end
   
   for ifType, __val in pairs( self.ifSet ) do
      local work = ifType
      if self.nilable then
         work = work:get_nilableTypeInfo()
      end
      
      if isMutable( self.mutMode ) then
         return work
      end
      
      return processInfo:createModifier( work, self.mutMode )
   end
   
   error( "illegal" )
end
function CombineType.setmeta( obj )
  setmetatable( obj, { __index = CombineType  } )
end


local CommonType = {}
CommonType._name2Val = {}
_moduleObj.CommonType = CommonType
function CommonType:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "CommonType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function CommonType._from( val )
   return _lune._AlgeFrom( CommonType, val )
end

CommonType.Combine = { "Combine", {{}}}
CommonType._name2Val["Combine"] = CommonType.Combine
CommonType.Normal = { "Normal", {{}}}
CommonType._name2Val["Normal"] = CommonType.Normal


function CombineType:andType( processInfo, other, alt2type )

   do
      local _matchExp = other
      if _matchExp[1] == CommonType.Combine[1] then
         local comboInfo = _matchExp[2][1]
      
         self:andIfSet( processInfo, comboInfo.ifSet, alt2type )
         if not isMutable( comboInfo.mutMode ) then
            self.mutMode = comboInfo.mutMode
         end
         
         return _lune.newAlge( CommonType.Combine, {self})
      elseif _matchExp[1] == CommonType.Normal[1] then
         local typeInfo = _matchExp[2][1]
      
         if not isMutable( typeInfo:get_mutMode() ) then
            self.mutMode = typeInfo:get_mutMode()
         end
         
         local ifSet = {}
         if typeInfo:get_kind() == TypeInfoKind.IF then
            ifSet[typeInfo]= true
         else
          
            for __index, iftype in ipairs( typeInfo:get_interfaceList() ) do
               ifSet[iftype]= true
            end
            
         end
         
         self:andIfSet( processInfo, ifSet, alt2type )
         
         if _lune._Set_len(self.ifSet ) ~= 0 then
            return _lune.newAlge( CommonType.Combine, {self})
         end
         
         return _lune.newAlge( CommonType.Normal, {self:createStem( processInfo )})
      end
   end
   
end


function TypeInfo.getCommonTypeCombo( processInfo, commonType, otherType, alt2type )

   local typeInfo = _moduleObj.builtinTypeNone
   
   do
      local _matchExp = commonType
      if _matchExp[1] == CommonType.Combine[1] then
         local comb = _matchExp[2][1]
      
         return comb:andType( processInfo, otherType, alt2type )
      elseif _matchExp[1] == CommonType.Normal[1] then
         local workTypeInfo = _matchExp[2][1]
      
         typeInfo = workTypeInfo
      end
   end
   
   
   local other = _moduleObj.builtinTypeNone
   do
      local _matchExp = otherType
      if _matchExp[1] == CommonType.Combine[1] then
         local comb = _matchExp[2][1]
      
         return comb:andType( processInfo, commonType, alt2type )
      elseif _matchExp[1] == CommonType.Normal[1] then
         local workTypeInfo = _matchExp[2][1]
      
         other = workTypeInfo
      end
   end
   
   
   local function getType( workType )
   
      if typeInfo:get_nilable() or other:get_nilable() then
         workType = workType:get_nilableTypeInfo()
      end
      
      if not TypeInfo.isMut( typeInfo ) or not TypeInfo.isMut( other ) then
         workType = processInfo:createModifier( workType, MutMode.IMut )
      end
      
      return _lune.newAlge( CommonType.Normal, {workType})
   end
   
   local type1 = typeInfo:get_nonnilableType():get_srcTypeInfo()
   local type2 = other:get_nonnilableType():get_srcTypeInfo()
   
   if type1 == _moduleObj.builtinTypeNone then
      return otherType
   end
   
   if type2 == _moduleObj.builtinTypeNone then
      return commonType
   end
   
   
   if type1 == _moduleObj.builtinTypeNil then
      return _lune.newAlge( CommonType.Normal, {other:get_nilableTypeInfo()})
   end
   
   if type2 == _moduleObj.builtinTypeNil then
      return _lune.newAlge( CommonType.Normal, {typeInfo:get_nilableTypeInfo()})
   end
   
   
   if type1:canEvalWith( processInfo, type2, CanEvalType.SetOp, alt2type ) then
      return getType( type1 )
   end
   
   if type2:canEvalWith( processInfo, type1, CanEvalType.SetOp, alt2type ) then
      return getType( type2 )
   end
   
   
   local mutMode
   
   if TypeInfo.isMut( typeInfo ) and TypeInfo.isMut( other ) then
      mutMode = MutMode.Mut
   else
    
      mutMode = MutMode.IMut
   end
   
   
   if type1:get_kind() == type2:get_kind() then
      
      local function getCommon( workTypeInfo, workOther, workAlt2type )
      
         do
            local _matchExp = TypeInfo.getCommonTypeCombo( processInfo, _lune.newAlge( CommonType.Normal, {workTypeInfo}), _lune.newAlge( CommonType.Normal, {workOther}), workAlt2type )
            if _matchExp[1] == CommonType.Normal[1] then
               local info = _matchExp[2][1]
            
               return info
            elseif _matchExp[1] == CommonType.Combine[1] then
               local combine = _matchExp[2][1]
            
               return combine:get_typeInfo( processInfo )
            end
         end
         
      end
      
      do
         local _switchExp = type1:get_kind()
         if _switchExp == TypeInfoKind.List then
            return getType( processInfo:createList( AccessMode.Local, _moduleObj.headTypeInfo, {getCommon( type1:get_itemTypeInfoList()[1], type2:get_itemTypeInfoList()[1], alt2type )}, mutMode ) )
         elseif _switchExp == TypeInfoKind.Array then
            return getType( processInfo:createArray( AccessMode.Local, _moduleObj.headTypeInfo, {getCommon( type1:get_itemTypeInfoList()[1], type2:get_itemTypeInfoList()[1], alt2type )}, mutMode ) )
         elseif _switchExp == TypeInfoKind.Set then
            return getType( processInfo:createSet( AccessMode.Local, _moduleObj.headTypeInfo, {getCommon( type1:get_itemTypeInfoList()[1], type2:get_itemTypeInfoList()[1], alt2type )}, mutMode ) )
         elseif _switchExp == TypeInfoKind.Map then
            return getType( processInfo:createMap( AccessMode.Local, _moduleObj.headTypeInfo, getCommon( type1:get_itemTypeInfoList()[1], type2:get_itemTypeInfoList()[1], alt2type ), getCommon( type1:get_itemTypeInfoList()[2], type2:get_itemTypeInfoList()[2], alt2type ), mutMode ) )
         end
      end
      
   end
   
   
   local work = type1:get_baseTypeInfo()
   
   while work ~= _moduleObj.headTypeInfo do
      if work:canEvalWith( processInfo, type2, CanEvalType.SetOp, alt2type ) then
         if typeInfo:get_nilable() or other:get_nilable() then
            work = work:get_nilableTypeInfo()
         end
         
         if not isMutable( mutMode ) then
            work = processInfo:createModifier( work, mutMode )
         end
         
         return _lune.newAlge( CommonType.Normal, {work})
      end
      
      work = work:get_baseTypeInfo()
   end
   
   local combine = CombineType.new(typeInfo)
   return combine:andType( processInfo, _lune.newAlge( CommonType.Normal, {other}), alt2type )
end


function TypeInfo.getCommonType( processInfo, typeInfo, other, alt2type )

   do
      local _matchExp = TypeInfo.getCommonTypeCombo( processInfo, _lune.newAlge( CommonType.Normal, {typeInfo}), _lune.newAlge( CommonType.Normal, {other}), alt2type )
      if _matchExp[1] == CommonType.Normal[1] then
         local info = _matchExp[2][1]
      
         return info
      elseif _matchExp[1] == CommonType.Combine[1] then
         local combine = _matchExp[2][1]
      
         return combine:get_typeInfo( processInfo )
      end
   end
   
end


function DDDTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( "...", typeNameCtrl, importInfo, localFlag )
end


function DDDTypeInfo:getTxtWithRaw( raw, typeNameCtrl, importInfo, localFlag )

   if self.typeInfo == _moduleObj.builtinTypeStem_ then
      if self:get_extTypeFlag() then
         return "Luaval<...>"
      end
      
      return "..."
   end
   
   local typeInfo
   
   if self:get_extTypeFlag() then
      typeInfo = self.typeInfo:get_extedType()
   else
    
      typeInfo = self.typeInfo
   end
   
   local txt = string.format( "...<%s>", typeInfo:getTxt( typeNameCtrl, importInfo, localFlag ))
   if self:get_extTypeFlag() then
      return string.format( "Luaval<%s>", txt)
   end
   
   return txt
end

function ProcessInfo:createGeneric( genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo )

   self:get_idProv():increment(  )
   return GenericTypeInfo.new(self, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo)
end


local function applyGenericList( typeList, alt2typeMap, moduleTypeInfo )

   local typeInfoList = {}
   local needNew = false
   for __index, srcType in ipairs( typeList ) do
      do
         local typeInfo = srcType:applyGeneric( alt2typeMap, moduleTypeInfo )
         if typeInfo ~= nil then
            table.insert( typeInfoList, typeInfo )
            if srcType ~= typeInfo then
               needNew = true
            end
            
         else
            return nil, false
         end
      end
      
   end
   
   return typeInfoList, needNew
end

function GenericTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   if self.genSrcTypeInfo:get_kind() == TypeInfoKind.Class then
      local itemTypeInfoList, newFlag = applyGenericList( self:get_itemTypeInfoList(), alt2typeMap, moduleTypeInfo )
      if itemTypeInfoList ~= nil then
         if newFlag then
            return moduleTypeInfo:getProcessInfo(  ):createGeneric( self.genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo )
         end
         
      end
      
   end
   
   
   local genSrcTypeInfo = self.genSrcTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )
   if genSrcTypeInfo == self.genSrcTypeInfo then
      return self
   end
   
   if not self.hasAlter then
      return self
   end
   
   
   Util.errorLog( string.format( "no support nest generic -- %s", self:getTxt(  )) )
   return nil
end


local AbbrTypeInfo = {}
setmetatable( AbbrTypeInfo, { __index = TypeInfo } )
_moduleObj.AbbrTypeInfo = AbbrTypeInfo
function AbbrTypeInfo:get_scope(  )

   return nil
end
function AbbrTypeInfo:get_baseTypeInfo(  )

   return _moduleObj.headTypeInfo
end
function AbbrTypeInfo:get_parentInfo(  )

   return _moduleObj.headTypeInfo
end
function AbbrTypeInfo.new( processInfo, rawTxt )
   local obj = {}
   AbbrTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, rawTxt ); end
   return obj
end
function AbbrTypeInfo:__init(processInfo, rawTxt) 
   TypeInfo.__init( self,nil, processInfo)
   
   
   self.typeId = processInfo:newId(  )
   self.rawTxt = rawTxt
end
function AbbrTypeInfo:isModule(  )

   return false
end
function AbbrTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function AbbrTypeInfo:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag )

   return rawTxt
end
function AbbrTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return false, nil
end
function AbbrTypeInfo:serialize( stream, serializeInfo )

   Util.err( "illegal call" )
end
function AbbrTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function AbbrTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
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


local builtinTypeAbbr = AbbrTypeInfo.new(rootProcessInfo, "##")
_moduleObj.builtinTypeAbbr = builtinTypeAbbr

local builtinTypeAbbrNone = AbbrTypeInfo.new(rootProcessInfo, "[##]")
_moduleObj.builtinTypeAbbrNone = builtinTypeAbbrNone


local ExtTypeInfo = {}
setmetatable( ExtTypeInfo, { __index = TypeInfo } )
_moduleObj.ExtTypeInfo = ExtTypeInfo
function ExtTypeInfo.new( processInfo, extedType )
   local obj = {}
   ExtTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, extedType ); end
   return obj
end
function ExtTypeInfo:__init(processInfo, extedType) 
   TypeInfo.__init( self,extedType:get_scope(), processInfo)
   
   
   self.typeId = processInfo:newId(  )
   self.extedType = extedType
   
   self.nilableTypeInfo = NilableTypeInfo.new(nil, processInfo, self, processInfo:newId(  ))
end
function ExtTypeInfo:getTxt( typeNameCtrl, importInfo, localFlag )

   return self:getTxtWithRaw( self:get_rawTxt(), typeNameCtrl, importInfo, localFlag )
end
function ExtTypeInfo:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag )

   return string.format( "Luaval<%s>", self.extedType:getTxtWithRaw( rawTxt, typeNameCtrl, importInfo, localFlag ))
end
function ExtTypeInfo:equals( processInfo, typeInfo, alt2type, checkModifer )

   do
      local extTypeInfo = _lune.__Cast( typeInfo, 3, ExtTypeInfo )
      if extTypeInfo ~= nil then
         return self.extedType:equals( processInfo, extTypeInfo.extedType, alt2type, checkModifer )
      end
   end
   
   if failCreateLuavalWith( self.extedType, LuavalConvKind.InLua, false ) then
      return false
   end
   
   return self.extedType:equals( processInfo, typeInfo, alt2type, checkModifer )
end
function ExtTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   do
      local extTypeInfo = _lune.__Cast( other:get_nonnilableType(), 3, ExtTypeInfo )
      if extTypeInfo ~= nil then
         local otherExtedType
         
         if other:get_nilable() then
            otherExtedType = extTypeInfo.extedType:get_nilableTypeInfo()
         else
          
            otherExtedType = extTypeInfo.extedType
         end
         
         return self.extedType:canEvalWith( processInfo, otherExtedType, canEvalType, alt2type )
      end
   end
   
   do
      local _exp = failCreateLuavalWith( other, LuavalConvKind.ToLua, true )
      if _exp ~= nil then
         return false, _exp
      end
   end
   
   return true, nil
   
end
function ExtTypeInfo:serialize( stream, serializeInfo )

   stream:write( string.format( '{ skind = %d, typeId = %d, extedTypeId = %d }\n', SerializeKind.Ext, self.typeId.id, self.extedType:get_typeId().id) )
end
function ExtTypeInfo:get_display_stirng_with( raw, alt2type )

   return self:getTxtWithRaw( raw )
end
function ExtTypeInfo:get_display_stirng(  )

   return self:get_display_stirng_with( self:get_rawTxt(), nil )
end
function ExtTypeInfo:getModule(  )

   return _moduleObj.headTypeInfo
end
function ExtTypeInfo:get_kind(  )

   return TypeInfoKind.Ext
end
function ExtTypeInfo:get_aliasSrc(  )

   return self
end
function ExtTypeInfo:get_srcTypeInfo(  )

   return self
end
function ExtTypeInfo:get_nonnilableType(  )

   return self
end
function ExtTypeInfo:get_nilable(  )

   return false
end
function ExtTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local typeInfo = self.extedType:applyGeneric( alt2typeMap, moduleTypeInfo )
   if typeInfo ~= self.extedType then
      Util.err( string.format( "not support generics -- %s", self.extedType:getTxt(  )) )
   end
   
   return self
end
function ExtTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = ExtTypeInfo  } )
end
function ExtTypeInfo:get_typeId()
   return self.typeId
end
function ExtTypeInfo:get_extedType()
   return self.extedType
end
function ExtTypeInfo:get_nilableTypeInfo()
   return self.nilableTypeInfo
end
function ExtTypeInfo:addChildren( ... )
   return self.extedType:addChildren( ... )
end

function ExtTypeInfo:createAlt2typeMap( ... )
   return self.extedType:createAlt2typeMap( ... )
end

function ExtTypeInfo:getFullName( ... )
   return self.extedType:getFullName( ... )
end

function ExtTypeInfo:getOverridingType( ... )
   return self.extedType:getOverridingType( ... )
end

function ExtTypeInfo:getParentFullName( ... )
   return self.extedType:getParentFullName( ... )
end

function ExtTypeInfo:getParentId( ... )
   return self.extedType:getParentId( ... )
end

function ExtTypeInfo:getProcessInfo( ... )
   return self.extedType:getProcessInfo( ... )
end

function ExtTypeInfo:get_abstractFlag( ... )
   return self.extedType:get_abstractFlag( ... )
end

function ExtTypeInfo:get_accessMode( ... )
   return self.extedType:get_accessMode( ... )
end

function ExtTypeInfo:get_argTypeInfoList( ... )
   return self.extedType:get_argTypeInfoList( ... )
end

function ExtTypeInfo:get_autoFlag( ... )
   return self.extedType:get_autoFlag( ... )
end

function ExtTypeInfo:get_baseId( ... )
   return self.extedType:get_baseId( ... )
end

function ExtTypeInfo:get_baseTypeInfo( ... )
   return self.extedType:get_baseTypeInfo( ... )
end

function ExtTypeInfo:get_children( ... )
   return self.extedType:get_children( ... )
end

function ExtTypeInfo:get_externalFlag( ... )
   return self.extedType:get_externalFlag( ... )
end

function ExtTypeInfo:get_genSrcTypeInfo( ... )
   return self.extedType:get_genSrcTypeInfo( ... )
end

function ExtTypeInfo:get_interfaceList( ... )
   return self.extedType:get_interfaceList( ... )
end

function ExtTypeInfo:get_itemTypeInfoList( ... )
   return self.extedType:get_itemTypeInfoList( ... )
end

function ExtTypeInfo:get_mutMode( ... )
   return self.extedType:get_mutMode( ... )
end

function ExtTypeInfo:get_parentInfo( ... )
   return self.extedType:get_parentInfo( ... )
end

function ExtTypeInfo:get_processInfo( ... )
   return self.extedType:get_processInfo( ... )
end

function ExtTypeInfo:get_rawTxt( ... )
   return self.extedType:get_rawTxt( ... )
end

function ExtTypeInfo:get_retTypeInfoList( ... )
   return self.extedType:get_retTypeInfoList( ... )
end

function ExtTypeInfo:get_scope( ... )
   return self.extedType:get_scope( ... )
end

function ExtTypeInfo:get_staticFlag( ... )
   return self.extedType:get_staticFlag( ... )
end

function ExtTypeInfo:get_typeData( ... )
   return self.extedType:get_typeData( ... )
end

function ExtTypeInfo:hasBase( ... )
   return self.extedType:hasBase( ... )
end

function ExtTypeInfo:hasRouteNamespaceFrom( ... )
   return self.extedType:hasRouteNamespaceFrom( ... )
end

function ExtTypeInfo:isInheritFrom( ... )
   return self.extedType:isInheritFrom( ... )
end

function ExtTypeInfo:isModule( ... )
   return self.extedType:isModule( ... )
end

function ExtTypeInfo:serializeTypeInfoList( ... )
   return self.extedType:serializeTypeInfoList( ... )
end

function ExtTypeInfo:switchScope( ... )
   return self.extedType:switchScope( ... )
end



function ProcessInfo:createLuaval( luneType, validToCheck )

   if not self.validExtType then
      return _lune.newAlge( LuavalResult.OK, {luneType,true})
   end
   
   
   if luneType:get_kind() == TypeInfoKind.Method then
      return _lune.newAlge( LuavalResult.OK, {luneType,true})
   end
   
   if isExtType( luneType ) then
      return _lune.newAlge( LuavalResult.OK, {luneType,true})
   end
   
   do
      local extType = self:get_typeInfo2Map().ExtMap[luneType]
      if extType ~= nil then
         if extType:get_typeId().id < userStartId and luneType:get_typeId().id >= userStartId then
            Util.err( "on cache" )
         end
         
         if extType:get_kind() == TypeInfoKind.Ext then
            return _lune.newAlge( LuavalResult.OK, {extType,false})
         end
         
         return _lune.newAlge( LuavalResult.OK, {extType,true})
      end
   end
   
   
   local function process(  )
   
      do
         local dddType = _lune.__Cast( luneType, 3, DDDTypeInfo )
         if dddType ~= nil then
            do
               local _matchExp = self:createLuaval( dddType:get_typeInfo(), validToCheck )
               if _matchExp[1] == LuavalResult.Err[1] then
                  local mess = _matchExp[2][1]
               
                  Util.err( mess )
               elseif _matchExp[1] == LuavalResult.OK[1] then
                  local workType = _matchExp[2][1]
                  local _ = _matchExp[2][2]
               
                  return _lune.newAlge( LuavalResult.OK, {self:createDDD( workType, dddType:get_externalFlag(), true ),false})
               end
            end
            
         end
      end
      
      local err, canConv
      
      err, canConv = failCreateLuavalWith( luneType, LuavalConvKind.InLua, validToCheck )
      if err ~= nil then
         return _lune.newAlge( LuavalResult.Err, {err})
      end
      
      if canConv then
         return _lune.newAlge( LuavalResult.OK, {luneType,true})
      end
      
      self:get_idProv():increment(  )
      local extType = ExtTypeInfo.new(self, luneType:get_nonnilableType())
      if luneType:get_nilable() then
         return _lune.newAlge( LuavalResult.OK, {extType:get_nilableTypeInfo(),false})
      end
      
      return _lune.newAlge( LuavalResult.OK, {extType,false})
   end
   local result = process(  )
   do
      local _matchExp = result
      if _matchExp[1] == LuavalResult.OK[1] then
         local typeInfo = _matchExp[2][1]
         local _ = _matchExp[2][2]
      
         self:get_typeInfo2Map().ExtMap[luneType:get_nonnilableType()] = typeInfo:get_nonnilableType()
         self:get_typeInfo2Map().ExtMap[luneType:get_nilableTypeInfo()] = typeInfo:get_nilableTypeInfo()
      end
   end
   
   return result
end


local builtinTypeLua

_moduleObj.builtinTypeLua = builtinTypeLua

do
   local _matchExp = rootProcessInfo:createLuaval( _moduleObj.builtinTypeStem, true )
   if _matchExp[1] == LuavalResult.OK[1] then
      local typeInfo = _matchExp[2][1]
      local _ = _matchExp[2][2]
   
      _moduleObj.builtinTypeLua = typeInfo
   else 
      
         Util.err( "illegal" )
   end
end

registBuiltin( "Luaval", "Luaval", TypeInfoKind.Ext, _moduleObj.builtinTypeLua, _moduleObj.headTypeInfo, false )
local builtinTypeDDDLua = rootProcessInfo:createDDD( _moduleObj.builtinTypeStem_, true, true )
_moduleObj.builtinTypeDDDLua = builtinTypeDDDLua

registBuiltin( "__LuaDDD", "__LuaDDD", TypeInfoKind.Ext, _moduleObj.builtinTypeDDDLua, _moduleObj.headTypeInfo, false )

local function convToExtTypeList( processInfo, list )

   local extList = {}
   for __index, typeInfo in ipairs( list ) do
      do
         local _matchExp = processInfo:createLuaval( typeInfo, true )
         if _matchExp[1] == LuavalResult.OK[1] then
            local extType = _matchExp[2][1]
            local pass = _matchExp[2][2]
         
            table.insert( extList, extType )
         elseif _matchExp[1] == LuavalResult.Err[1] then
            local err = _matchExp[2][1]
         
            return nil, err
         end
      end
      
   end
   
   return extList, ""
end
_moduleObj.convToExtTypeList = convToExtTypeList
local AndExpTypeInfo = {}
setmetatable( AndExpTypeInfo, { __index = TypeInfo } )
_moduleObj.AndExpTypeInfo = AndExpTypeInfo
function AndExpTypeInfo.new( processInfo, exp1, exp2, result )
   local obj = {}
   AndExpTypeInfo.setmeta( obj )
   if obj.__init then obj:__init( processInfo, exp1, exp2, result ); end
   return obj
end
function AndExpTypeInfo:__init(processInfo, exp1, exp2, result) 
   TypeInfo.__init( self,result:get_scope(), processInfo)
   
   self.exp1 = exp1
   self.exp2 = exp2
   self.result = result
end
function AndExpTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = AndExpTypeInfo  } )
end
function AndExpTypeInfo:get_exp1()
   return self.exp1
end
function AndExpTypeInfo:get_exp2()
   return self.exp2
end
function AndExpTypeInfo:get_result()
   return self.result
end
function AndExpTypeInfo:addChildren( ... )
   return self.result:addChildren( ... )
end

function AndExpTypeInfo:applyGeneric( ... )
   return self.result:applyGeneric( ... )
end

function AndExpTypeInfo:canEvalWith( ... )
   return self.result:canEvalWith( ... )
end

function AndExpTypeInfo:createAlt2typeMap( ... )
   return self.result:createAlt2typeMap( ... )
end

function AndExpTypeInfo:equals( ... )
   return self.result:equals( ... )
end

function AndExpTypeInfo:getFullName( ... )
   return self.result:getFullName( ... )
end

function AndExpTypeInfo:getModule( ... )
   return self.result:getModule( ... )
end

function AndExpTypeInfo:getOverridingType( ... )
   return self.result:getOverridingType( ... )
end

function AndExpTypeInfo:getParentFullName( ... )
   return self.result:getParentFullName( ... )
end

function AndExpTypeInfo:getParentId( ... )
   return self.result:getParentId( ... )
end

function AndExpTypeInfo:getProcessInfo( ... )
   return self.result:getProcessInfo( ... )
end

function AndExpTypeInfo:getTxt( ... )
   return self.result:getTxt( ... )
end

function AndExpTypeInfo:getTxtWithRaw( ... )
   return self.result:getTxtWithRaw( ... )
end

function AndExpTypeInfo:get_abstractFlag( ... )
   return self.result:get_abstractFlag( ... )
end

function AndExpTypeInfo:get_accessMode( ... )
   return self.result:get_accessMode( ... )
end

function AndExpTypeInfo:get_aliasSrc( ... )
   return self.result:get_aliasSrc( ... )
end

function AndExpTypeInfo:get_argTypeInfoList( ... )
   return self.result:get_argTypeInfoList( ... )
end

function AndExpTypeInfo:get_autoFlag( ... )
   return self.result:get_autoFlag( ... )
end

function AndExpTypeInfo:get_baseId( ... )
   return self.result:get_baseId( ... )
end

function AndExpTypeInfo:get_baseTypeInfo( ... )
   return self.result:get_baseTypeInfo( ... )
end

function AndExpTypeInfo:get_children( ... )
   return self.result:get_children( ... )
end

function AndExpTypeInfo:get_display_stirng( ... )
   return self.result:get_display_stirng( ... )
end

function AndExpTypeInfo:get_display_stirng_with( ... )
   return self.result:get_display_stirng_with( ... )
end

function AndExpTypeInfo:get_extedType( ... )
   return self.result:get_extedType( ... )
end

function AndExpTypeInfo:get_externalFlag( ... )
   return self.result:get_externalFlag( ... )
end

function AndExpTypeInfo:get_genSrcTypeInfo( ... )
   return self.result:get_genSrcTypeInfo( ... )
end

function AndExpTypeInfo:get_interfaceList( ... )
   return self.result:get_interfaceList( ... )
end

function AndExpTypeInfo:get_itemTypeInfoList( ... )
   return self.result:get_itemTypeInfoList( ... )
end

function AndExpTypeInfo:get_kind( ... )
   return self.result:get_kind( ... )
end

function AndExpTypeInfo:get_mutMode( ... )
   return self.result:get_mutMode( ... )
end

function AndExpTypeInfo:get_nilable( ... )
   return self.result:get_nilable( ... )
end

function AndExpTypeInfo:get_nilableTypeInfo( ... )
   return self.result:get_nilableTypeInfo( ... )
end

function AndExpTypeInfo:get_nonnilableType( ... )
   return self.result:get_nonnilableType( ... )
end

function AndExpTypeInfo:get_parentInfo( ... )
   return self.result:get_parentInfo( ... )
end

function AndExpTypeInfo:get_processInfo( ... )
   return self.result:get_processInfo( ... )
end

function AndExpTypeInfo:get_rawTxt( ... )
   return self.result:get_rawTxt( ... )
end

function AndExpTypeInfo:get_retTypeInfoList( ... )
   return self.result:get_retTypeInfoList( ... )
end

function AndExpTypeInfo:get_scope( ... )
   return self.result:get_scope( ... )
end

function AndExpTypeInfo:get_srcTypeInfo( ... )
   return self.result:get_srcTypeInfo( ... )
end

function AndExpTypeInfo:get_staticFlag( ... )
   return self.result:get_staticFlag( ... )
end

function AndExpTypeInfo:get_typeData( ... )
   return self.result:get_typeData( ... )
end

function AndExpTypeInfo:get_typeId( ... )
   return self.result:get_typeId( ... )
end

function AndExpTypeInfo:hasBase( ... )
   return self.result:hasBase( ... )
end

function AndExpTypeInfo:hasRouteNamespaceFrom( ... )
   return self.result:hasRouteNamespaceFrom( ... )
end

function AndExpTypeInfo:isInheritFrom( ... )
   return self.result:isInheritFrom( ... )
end

function AndExpTypeInfo:isModule( ... )
   return self.result:isModule( ... )
end

function AndExpTypeInfo:serialize( ... )
   return self.result:serialize( ... )
end

function AndExpTypeInfo:serializeTypeInfoList( ... )
   return self.result:serializeTypeInfoList( ... )
end

function AndExpTypeInfo:switchScope( ... )
   return self.result:switchScope( ... )
end



local numberTypeSet = {}
numberTypeSet[_moduleObj.builtinTypeInt]= true
numberTypeSet[_moduleObj.builtinTypeChar]= true
numberTypeSet[_moduleObj.builtinTypeReal]= true

local function isNumberType( typeInfo )

   return _lune._Set_has(numberTypeSet, typeInfo:get_srcTypeInfo() )
end
_moduleObj.isNumberType = isNumberType

function ProcessInfo:createEnum( scope, parentInfo, externalFlag, accessMode, enumName, valTypeInfo )

   if Parser.isLuaKeyword( enumName ) then
      Util.err( string.format( "This symbol can not use for a enum. -- %s", enumName) )
   end
   
   
   local info = EnumTypeInfo.new(self, scope, externalFlag, accessMode, enumName, parentInfo, valTypeInfo)
   
   local getEnumName = self:createFunc( false, true, nil, TypeInfoKind.Method, info, true, externalFlag, false, AccessMode.Pub, "get__txt", nil, nil, {_moduleObj.builtinTypeString}, false )
   scope:addMethod( self, nil, getEnumName, AccessMode.Pub, false, false )
   
   local fromVal = self:createFunc( false, true, nil, TypeInfoKind.Func, info, true, externalFlag, true, AccessMode.Pub, "_from", nil, {self:createModifier( valTypeInfo, MutMode.IMut )}, {info:get_nilableTypeInfo()}, false )
   scope:addMethod( self, nil, fromVal, AccessMode.Pub, true, false )
   
   local allListType = self:createList( AccessMode.Pub, info, {info}, MutMode.IMut )
   local allList = self:createFunc( false, true, nil, TypeInfoKind.Func, info, true, externalFlag, true, AccessMode.Pub, "get__allList", nil, nil, {self:createModifier( allListType, MutMode.IMut )}, false )
   scope:addMethod( self, nil, allList, AccessMode.Pub, true, false )
   
   return info
end


function EnumTypeInfo:serialize( stream, serializeInfo )

   local txt = string.format( [==[{ skind = %d, parentId = %d, typeId = %d, txt = '%s',
accessMode = %d, kind = %d, valTypeId = %d, ]==], SerializeKind.Enum, self:getParentId(  ).id, self.typeId.id, self.rawTxt, self.accessMode, TypeInfoKind.Enum, self.valTypeInfo:get_typeId().id)
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


function ProcessInfo:createAlge( scope, parentInfo, externalFlag, accessMode, algeName )

   if Parser.isLuaKeyword( algeName ) then
      Util.err( string.format( "This symbol can not use for a alge. -- %s", algeName) )
   end
   
   
   local info = AlgeTypeInfo.new(self, scope, externalFlag, accessMode, algeName, parentInfo)
   
   local getAlgeName = self:createFunc( false, true, nil, TypeInfoKind.Method, info, true, externalFlag, false, AccessMode.Pub, "get__txt", nil, nil, {_moduleObj.builtinTypeString}, false )
   scope:addMethod( self, nil, getAlgeName, AccessMode.Pub, false, false )
   
   return info
end


function AlgeTypeInfo:serialize( stream, serializeInfo )

   local txt = string.format( [==[{ skind = %d, parentId = %d, typeId = %d, txt = '%s',
accessMode = %d, kind = %d, ]==], SerializeKind.Alge, self:getParentId(  ).id, self.typeId.id, self.rawTxt, self.accessMode, TypeInfoKind.Alge)
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


function BoxTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   if self == other then
      return true, nil
   end
   
   do
      local _switchExp = canEvalType
      if _switchExp == CanEvalType.SetOp or _switchExp == CanEvalType.SetOpIMut or _switchExp == CanEvalType.SetEq then
      else 
         
            return false, nil
      end
   end
   
   
   if other:get_nilable() then
      CanEvalCtrlTypeInfo.updateNeedAutoBoxing( alt2type )
      return true, nil
   end
   
   
   do
      local otherBoxType = _lune.__Cast( other, 3, BoxTypeInfo )
      if otherBoxType ~= nil then
         return self.boxingType:canEvalWith( processInfo, otherBoxType.boxingType, canEvalType, alt2type )
      end
   end
   
   
   if self.boxingType:canEvalWith( processInfo, other, canEvalType, alt2type ) then
      CanEvalCtrlTypeInfo.updateNeedAutoBoxing( alt2type )
      return true, nil
   end
   
   return false, nil
end


function NilableTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   local otherSrc = other
   
   if self == _moduleObj.builtinTypeStem_ then
      return true, nil
   end
   
   
   if otherSrc == _moduleObj.builtinTypeNil or otherSrc:get_kind() == TypeInfoKind.Abbr then
      if self:get_nonnilableType():get_kind() == TypeInfoKind.Box then
         return self:get_nonnilableType():canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      end
      
      return true, nil
   end
   
   if self.typeId == otherSrc:get_typeId() then
      return true, nil
   end
   
   if otherSrc:get_nilable() then
      if otherSrc:get_kind() == TypeInfoKind.DDD then
         return self:get_nonnilableType():canEvalWith( processInfo, otherSrc:get_itemTypeInfoList()[1], canEvalType, alt2type )
      end
      
      
      return self:get_nonnilableType():canEvalWith( processInfo, otherSrc:get_nonnilableType(), canEvalType, alt2type )
   end
   
   return self:get_nonnilableType():canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
end





function NormalTypeInfo.isAvailableMapping( processInfo, typeInfo, checkedTypeMap )

   local function isAvailableMappingSub(  )
   
      do
         local _switchExp = typeInfo:get_kind()
         if _switchExp == TypeInfoKind.Prim or _switchExp == TypeInfoKind.Enum then
            return true
         elseif _switchExp == TypeInfoKind.Alge then
            local algeTypeInfo = _lune.unwrap( (_lune.__Cast( typeInfo, 3, AlgeTypeInfo ) ))
            for __index, valInfo in pairs( algeTypeInfo:get_valInfoMap() ) do
               for __index, paramType in ipairs( valInfo:get_typeList() ) do
                  if not NormalTypeInfo.isAvailableMapping( processInfo, paramType, checkedTypeMap ) then
                     return false
                  end
                  
               end
               
            end
            
            return true
         elseif _switchExp == TypeInfoKind.Stem then
            return true
         elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
            if typeInfo:equals( processInfo, _moduleObj.builtinTypeString ) then
               return true
            end
            
            return typeInfo:isInheritFrom( processInfo, _moduleObj.builtinTypeMapping, nil )
         elseif _switchExp == TypeInfoKind.Alternate then
            return typeInfo:isInheritFrom( processInfo, _moduleObj.builtinTypeMapping, nil )
         elseif _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set then
            return NormalTypeInfo.isAvailableMapping( processInfo, typeInfo:get_itemTypeInfoList()[1], checkedTypeMap )
         elseif _switchExp == TypeInfoKind.Map then
            if NormalTypeInfo.isAvailableMapping( processInfo, typeInfo:get_itemTypeInfoList()[2], checkedTypeMap ) then
               local keyType = typeInfo:get_itemTypeInfoList()[1]
               if keyType:equals( processInfo, _moduleObj.builtinTypeString ) or keyType:get_kind() == TypeInfoKind.Prim or keyType:get_kind() == TypeInfoKind.Enum then
                  return true
               end
               
            end
            
            return false
         elseif _switchExp == TypeInfoKind.Nilable then
            return NormalTypeInfo.isAvailableMapping( processInfo, typeInfo:get_nonnilableType(), checkedTypeMap )
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

function NormalTypeInfo:isInheritFrom( processInfo, other, alt2type )

   if self:get_typeId() == other:get_typeId() then
      return true
   end
   
   if (self:get_kind() ~= TypeInfoKind.Class and self:get_kind() ~= TypeInfoKind.IF ) or (other:get_kind() ~= TypeInfoKind.Class and other:get_kind() ~= TypeInfoKind.IF ) then
      if other == _moduleObj.builtinTypeMapping then
         return NormalTypeInfo.isAvailableMapping( processInfo, self, {} )
      end
      
      return false
   end
   
   return TypeInfo.isInherit( processInfo, self, other, alt2type )
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


function TypeInfo.checkMatchType( processInfo, dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex, alt2type )

   
   
   
   
   
   local warnMess = nil
   
   local function checkDstTypeFrom( index, srcType, srcType2nd )
   
      local workExpType = srcType
      for dstIndex = index, #dstTypeList do
         local workDstType = dstTypeList[dstIndex]
         local canEval, evalMess = workDstType:canEvalWith( processInfo, workExpType, CanEvalType.SetOp, alt2type )
         if not canEval then
            local message = string.format( "exp(%d) type mismatch %s <- %s: dst %d%s", dstIndex, workDstType:getTxt( _moduleObj.defaultTypeNameCtrl ), workExpType:getTxt( _moduleObj.defaultTypeNameCtrl ), dstIndex, (evalMess and string.format( " -- %s", evalMess) or "" )
            )
            return MatchType.Error, message
         elseif workExpType == _moduleObj.builtinTypeAbbrNone then
            return MatchType.Warn, Code.format( Code.ID.nothing_define_abbr, string.format( "use '##', instate of '%s'.", workDstType:getTxt( _moduleObj.defaultTypeNameCtrl )) )
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
            
         elseif srcIndex > index and expType:get_kind() == TypeInfoKind.Abbr then
            return MatchType.Error, "must not use '##'"
         end
         
         
         local canEval, evalMess = dstType:canEvalWith( processInfo, checkType, CanEvalType.SetOp, alt2type )
         if not canEval then
            return MatchType.Error, string.format( "exp(%d) type mismatch %s <- %s: src: %d%s", srcIndex, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ), srcIndex, (evalMess and string.format( " -- %s", evalMess) or "" )
            )
         end
         
         
         
         if warnForFollowSrcIndex ~= nil then
            
            if warnForFollowSrcIndex <= srcIndex then
               local workMess = string.format( "use '**' at arg(%d). %s <- %s", srcIndex, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ))
               return MatchType.Warn, workMess
            end
            
         end
         
         if dstType ~= _moduleObj.builtinTypeEmpty and not isExtType( dstType:get_srcTypeInfo():get_nonnilableType() ) and isExtType( expType:get_srcTypeInfo():get_nonnilableType() ) then
            warnMess = string.format( "exp(%d) luaval mismatch %s <- %s", index, dstType:getTxt(  ), expType:getTxt(  ))
         end
         
         
      end
      
      return MatchType.Match, ""
   end
   
   if #expTypeList > 0 then
      for index, expType in ipairs( expTypeList ) do
         if #dstTypeList == 0 then
            return MatchType.Error, string.format( "over exp. expect:0, actual:%d", #expTypeList)
         end
         
         local dstType = dstTypeList[index]
         if #dstTypeList == index then
            if dstType:get_srcTypeInfo():get_kind() ~= TypeInfoKind.DDD then
               local isMatch, msg = dstType:canEvalWith( processInfo, expType, CanEvalType.SetOp, alt2type )
               if not isMatch then
                  return MatchType.Error, string.format( "exp(%d) type mismatch %s <- %s: index %d%s", index, applyGenericDefault( dstType, alt2type, dstType:getModule(  ) ):getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ), index, msg and string.format( " -- %s", msg) or "")
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
               
               if warnForFollowSrcIndex <= index then
                  local workMess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ))
                  return MatchType.Warn, workMess
               end
               
            end
            
            if dstType ~= _moduleObj.builtinTypeEmpty and not isExtType( dstType:get_srcTypeInfo():get_nonnilableType() ) and isExtType( expType:get_srcTypeInfo():get_nonnilableType() ) then
               warnMess = string.format( "exp(%d) luaval mismatch %s <- %s", index, dstType:getTxt(  ), expType:getTxt(  ))
            end
            
            
            break
         elseif #expTypeList == index then
            local srcType = expType
            local srcType2nd = _moduleObj.builtinTypeAbbrNone
            if expType:get_kind() == TypeInfoKind.DDD then
               srcType = _moduleObj.builtinTypeStem_
               srcType2nd = _moduleObj.builtinTypeStem_
               if #expType:get_itemTypeInfoList() > 0 then
                  srcType = expType:get_itemTypeInfoList()[1]:get_nilableTypeInfo()
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
               
               if warnForFollowSrcIndex <= index then
                  local workMess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ))
                  return MatchType.Warn, workMess
               end
               
            end
            
            if dstType ~= _moduleObj.builtinTypeEmpty and not isExtType( dstType:get_srcTypeInfo():get_nonnilableType() ) and isExtType( expType:get_srcTypeInfo():get_nonnilableType() ) then
               warnMess = string.format( "exp(%d) luaval mismatch %s <- %s", index, dstType:getTxt(  ), expType:getTxt(  ))
            end
            
            
            break
         else
          
            local canEval, evalMess = dstType:canEvalWith( processInfo, expType, CanEvalType.SetOp, alt2type )
            
            if not canEval then
               return MatchType.Error, string.format( "exp(%d) type mismatch %s(%d) <- %s(%d)%s", index, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), dstType:get_typeId().id, expType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:get_typeId().id, (evalMess and string.format( " -- %s", evalMess) or "" )
               )
            end
            
            
            if warnForFollowSrcIndex ~= nil then
               
               if warnForFollowSrcIndex <= index then
                  local workMess = string.format( "use '**' at arg(%d). %s <- %s", index, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ), expType:getTxt( _moduleObj.defaultTypeNameCtrl ))
                  return MatchType.Warn, workMess
               end
               
            end
            
            if dstType ~= _moduleObj.builtinTypeEmpty and not isExtType( dstType:get_srcTypeInfo():get_nonnilableType() ) and isExtType( expType:get_srcTypeInfo():get_nonnilableType() ) then
               warnMess = string.format( "exp(%d) luaval mismatch %s <- %s", index, dstType:getTxt(  ), expType:getTxt(  ))
            end
            
            
         end
         
      end
      
   elseif not allowDstShort then
      for index, dstType in ipairs( dstTypeList ) do
         if not dstType:canEvalWith( processInfo, _moduleObj.builtinTypeNil, CanEvalType.SetOp, alt2type ) then
            return MatchType.Error, string.format( "exp(%d) type mismatch %s <- nil: short", index, dstType:getTxt( _moduleObj.defaultTypeNameCtrl ))
         end
         
         return MatchType.Warn, Code.format( Code.ID.nothing_define_abbr, string.format( "use '##', instate of '%s'.", dstType:getTxt( _moduleObj.defaultTypeNameCtrl )) )
      end
      
   end
   
   
   if warnMess ~= nil then
      return MatchType.Warn, warnMess
   end
   
   
   return MatchType.Match, ""
end

local function isSettableToForm( processInfo, typeInfo )

   if #typeInfo:get_argTypeInfoList() > 0 then
      for __index, argType in ipairs( typeInfo:get_argTypeInfoList() ) do
         do
            local dddType = _lune.__Cast( argType, 3, DDDTypeInfo )
            if dddType ~= nil then
               if not dddType:get_typeInfo():get_nilableTypeInfo():equals( processInfo, _moduleObj.builtinTypeStem_ ) then
                  return false
               end
               
            else
               if not argType:get_srcTypeInfo():equals( processInfo, _moduleObj.builtinTypeStem_ ) then
                  return false
               end
               
            end
         end
         
      end
      
   end
   
   return true
end

function TypeInfo.canEvalWithBase( processInfo, dest, destMut, other, canEvalType, alt2type )

   if dest ~= dest:get_aliasSrc() then
      return dest:get_aliasSrc():canEvalWith( processInfo, other, canEvalType, alt2type )
   end
   
   
   if dest == _moduleObj.builtinTypeExp or dest == _moduleObj.builtinTypeMultiExp then
      if other == _moduleObj.builtinTypeMultiExp and dest ~= _moduleObj.builtinTypeMultiExp then
         return false, "can't eval from '__exp' to '__exps'."
      end
      
      if other:equals( processInfo, _moduleObj.builtinTypeAbbr ) then
         return false, nil
      end
      
      return true, nil
   end
   
   local otherMut = TypeInfo.isMut( other )
   local otherSrc = other:get_srcTypeInfo():get_aliasSrc()
   if otherSrc:get_kind() == TypeInfoKind.DDD then
      if #otherSrc:get_itemTypeInfoList() > 0 then
         otherSrc = otherSrc:get_itemTypeInfoList()[1]:get_nilableTypeInfo()
      else
       
         otherSrc = _moduleObj.builtinTypeStem_
      end
      
   end
   
   
   if otherSrc:get_kind() == TypeInfoKind.Alternate then
      if dest:get_kind() ~= TypeInfoKind.Alternate then
         otherSrc = AlternateTypeInfo.getAssign( otherSrc, alt2type ):get_srcTypeInfo()
      end
      
   end
   
   
   do
      local _switchExp = canEvalType
      if _switchExp == CanEvalType.SetEq or _switchExp == CanEvalType.SetOp or _switchExp == CanEvalType.SetOpIMut then
         if dest == _moduleObj.builtinTypeEmpty then
            do
               local _switchExp = otherSrc
               if _switchExp == _moduleObj.builtinTypeAbbr or _switchExp == _moduleObj.builtinTypeAbbrNone then
                  return false, nil
               end
            end
            
            if otherSrc:get_kind() == TypeInfoKind.Func then
               return false, nil
            end
            
            return true, nil
         elseif not otherMut and destMut then
            if processInfo:get_validCheckingMutable() then
               local nonNilOtherType = otherSrc:get_nonnilableType()
               do
                  local _switchExp = nonNilOtherType:get_kind()
                  if _switchExp == TypeInfoKind.Set or _switchExp == TypeInfoKind.Map or _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.IF or _switchExp == TypeInfoKind.Alternate then
                     return false, nil
                  elseif _switchExp == TypeInfoKind.Class then
                     if _moduleObj.builtinTypeString ~= nonNilOtherType then
                        return false, nil
                     end
                     
                  elseif _switchExp == TypeInfoKind.Prim then
                     if _moduleObj.builtinTypeStem == nonNilOtherType then
                        return false, nil
                     end
                     
                  end
               end
               
            end
            
         end
         
         
         if otherSrc ~= _moduleObj.builtinTypeNil and otherSrc ~= _moduleObj.builtinTypeString and otherSrc:get_kind() ~= TypeInfoKind.Prim and otherSrc:get_kind() ~= TypeInfoKind.Func and otherSrc:get_kind() ~= TypeInfoKind.Method and otherSrc:get_kind() ~= TypeInfoKind.Form and otherSrc:get_kind() ~= TypeInfoKind.FormFunc and otherSrc:get_kind() ~= TypeInfoKind.Enum and otherSrc:get_kind() ~= TypeInfoKind.Abbr and otherSrc:get_kind() ~= TypeInfoKind.Alternate and otherSrc:get_kind() ~= TypeInfoKind.Box and not isGenericType( otherSrc ) and destMut and not otherMut then
            if processInfo:get_validCheckingMutable() then
               return false, nil
            end
            
         end
         
      end
   end
   
   
   if dest == _moduleObj.builtinTypeStem_ then
      return true, nil
   end
   
   
   if dest:get_srcTypeInfo():get_kind() == TypeInfoKind.DDD then
      if #dest:get_itemTypeInfoList() > 0 then
         return dest:get_itemTypeInfoList()[1]:canEvalWith( processInfo, other, canEvalType, alt2type )
      end
      
      return true, nil
   end
   
   if not dest:get_nilable() and otherSrc:get_nilable() then
      if canEvalType ~= CanEvalType.Equal then
         if dest:get_kind() == TypeInfoKind.Box then
            return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
         end
         
         return false, nil
      else
       
         otherSrc = otherSrc:get_nonnilableType()
      end
      
   end
   
   if dest == _moduleObj.builtinTypeStem and not otherSrc:get_nilable() then
      return true, nil
   end
   
   if dest == _moduleObj.builtinTypeForm and (otherSrc:get_kind() == TypeInfoKind.Func or otherSrc:get_kind() == TypeInfoKind.Form or otherSrc:get_kind() == TypeInfoKind.FormFunc ) then
      return isSettableToForm( processInfo, otherSrc ), nil
   end
   
   if otherSrc == _moduleObj.builtinTypeNil or otherSrc:get_kind() == TypeInfoKind.Abbr then
      if dest:get_kind() ~= TypeInfoKind.Nilable then
         return false, nil
      end
      
      return true, nil
   end
   
   if dest:get_typeId() == otherSrc:get_typeId() then
      return true, nil
   end
   
   
   if dest:get_kind() == TypeInfoKind.Ext then
      return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
   end
   
   do
      local extTypeInfo = _lune.__Cast( otherSrc, 3, ExtTypeInfo )
      if extTypeInfo ~= nil then
         if canEvalType ~= CanEvalType.SetEq and not failCreateLuavalWith( extTypeInfo:get_extedType(), LuavalConvKind.ToLua, false ) then
            otherSrc = extTypeInfo:get_extedType()
         end
         
      end
   end
   
   
   if dest:get_kind() ~= otherSrc:get_kind() then
      if otherSrc:get_kind() == TypeInfoKind.Alternate and otherSrc:hasBase(  ) then
         return TypeInfo.canEvalWithBase( processInfo, dest, destMut, otherSrc:get_baseTypeInfo(), canEvalType, alt2type )
      end
      
      
      if dest:get_kind() == TypeInfoKind.Nilable then
         local dstNonNil
         
         if destMut then
            dstNonNil = dest:get_nonnilableType()
         else
          
            dstNonNil = processInfo:createModifier( dest:get_nonnilableType(), MutMode.IMut )
         end
         
         
         if otherSrc:get_nilable() then
            if otherSrc:get_kind() == TypeInfoKind.DDD then
               return dstNonNil:canEvalWith( processInfo, otherSrc:get_itemTypeInfoList()[1], canEvalType, alt2type )
            end
            
            return dstNonNil:canEvalWith( processInfo, otherSrc:get_nonnilableType(), canEvalType, alt2type )
         end
         
         return dstNonNil:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      elseif isGenericType( dest ) then
         return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      elseif (dest:get_kind() == TypeInfoKind.Class or dest:get_kind() == TypeInfoKind.IF ) and (otherSrc:get_kind() == TypeInfoKind.Class or otherSrc:get_kind() == TypeInfoKind.IF ) then
         if canEvalType == CanEvalType.SetOp or canEvalType == CanEvalType.SetOpIMut then
            return otherSrc:isInheritFrom( processInfo, dest, alt2type ), nil
         end
         
         return false, nil
      elseif otherSrc:get_kind() == TypeInfoKind.Enum then
         local enumTypeInfo = _lune.unwrap( (_lune.__Cast( otherSrc, 3, EnumTypeInfo ) ))
         return dest:canEvalWith( processInfo, enumTypeInfo:get_valTypeInfo(), canEvalType, alt2type )
      elseif dest:get_kind() == TypeInfoKind.Alternate then
         return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      elseif dest:get_kind() == TypeInfoKind.Box then
         return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      elseif dest:get_kind() == TypeInfoKind.Form then
         do
            local _switchExp = otherSrc:get_kind()
            if _switchExp == TypeInfoKind.Form then
               return true, nil
            elseif _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Func then
               return isSettableToForm( processInfo, otherSrc ), nil
            end
         end
         
      elseif dest:get_kind() == TypeInfoKind.FormFunc then
         do
            local _switchExp = otherSrc:get_kind()
            if _switchExp == TypeInfoKind.FormFunc or _switchExp == TypeInfoKind.Func then
               do
                  local result, mess = TypeInfo.checkMatchType( processInfo, dest:get_argTypeInfoList(), otherSrc:get_argTypeInfoList(), false, nil, alt2type )
                  if result == MatchType.Error then
                     return false, mess
                  end
                  
               end
               
               do
                  local result, mess = TypeInfo.checkMatchType( processInfo, otherSrc:get_argTypeInfoList(), dest:get_argTypeInfoList(), false, nil, alt2type )
                  if result == MatchType.Error then
                     return false, mess
                  end
                  
               end
               
               do
                  local result, mess = TypeInfo.checkMatchType( processInfo, dest:get_retTypeInfoList(), otherSrc:get_retTypeInfoList(), false, nil, alt2type )
                  if result == MatchType.Error then
                     return false, mess
                  end
                  
               end
               
               do
                  local result, mess = TypeInfo.checkMatchType( processInfo, otherSrc:get_retTypeInfoList(), dest:get_retTypeInfoList(), false, nil, alt2type )
                  if result == MatchType.Error then
                     return false, mess
                  end
                  
               end
               
               if #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
                  return false, string.format( "argNum %d != %d", #dest:get_retTypeInfoList(), #otherSrc:get_retTypeInfoList())
               end
               
               
               return true, nil
            end
         end
         
      end
      
      return false, nil
   end
   
   do
      local _switchExp = dest:get_kind()
      if _switchExp == TypeInfoKind.Prim then
         if dest == _moduleObj.builtinTypeInt and otherSrc == _moduleObj.builtinTypeChar or dest == _moduleObj.builtinTypeChar and otherSrc == _moduleObj.builtinTypeInt then
            return true, nil
         end
         
         return false, nil
      elseif _switchExp == TypeInfoKind.List or _switchExp == TypeInfoKind.Array or _switchExp == TypeInfoKind.Set then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone then
            return true, nil
         end
         
         
         if #dest:get_itemTypeInfoList() >= 1 and #otherSrc:get_itemTypeInfoList() >= 1 then
            
            local ret, mess = (dest:get_itemTypeInfoList()[1] ):canEvalWith( processInfo, otherSrc:get_itemTypeInfoList()[1], destMut and CanEvalType.SetEq or CanEvalType.SetOpIMut, alt2type )
            if not ret then
               return false, mess
            end
            
         else
          
            return false, nil
         end
         
         
         return true, nil
      elseif _switchExp == TypeInfoKind.Map then
         if otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone and otherSrc:get_itemTypeInfoList()[2] == _moduleObj.builtinTypeNone then
            return true, nil
         end
         
         local function check1(  )
         
            
            if #dest:get_itemTypeInfoList() >= 1 and #otherSrc:get_itemTypeInfoList() >= 1 then
               
               local ret, mess = (dest:get_itemTypeInfoList()[1] ):canEvalWith( processInfo, otherSrc:get_itemTypeInfoList()[1], destMut and CanEvalType.SetEq or CanEvalType.SetOpIMut, alt2type )
               if not ret then
                  return false
               end
               
            else
             
               return false
            end
            
            
            return true
         end
         local function check2(  )
         
            
            if #dest:get_itemTypeInfoList() >= 2 and #otherSrc:get_itemTypeInfoList() >= 2 then
               
               local ret, mess = (dest:get_itemTypeInfoList()[2] ):canEvalWith( processInfo, otherSrc:get_itemTypeInfoList()[2], destMut and CanEvalType.SetEq or CanEvalType.SetOpIMut, alt2type )
               if not ret then
                  return false
               end
               
            else
             
               return false
            end
            
            
            return true
         end
         local result1 = check1(  )
         local result2 = check2(  )
         if result1 and result2 then
            return true, nil
         end
         
         if not result1 and otherSrc:get_itemTypeInfoList()[1] == _moduleObj.builtinTypeNone or not result2 and otherSrc:get_itemTypeInfoList()[2] == _moduleObj.builtinTypeNone then
            return true, nil
         end
         
         return false, nil
      elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
         if isGenericType( dest ) then
            return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
         end
         
         if canEvalType == CanEvalType.SetOp or canEvalType == CanEvalType.SetOpIMut then
            return otherSrc:isInheritFrom( processInfo, dest, alt2type ), nil
         end
         
         return false, nil
         
      elseif _switchExp == TypeInfoKind.Form then
         return isSettableToForm( processInfo, otherSrc ), nil
      elseif _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.FormFunc then
         if #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false, string.format( "argNum %d != %d", #dest:get_retTypeInfoList(), #otherSrc:get_retTypeInfoList())
         end
         
         local argCheck, argMess = TypeInfo.checkMatchType( processInfo, dest:get_argTypeInfoList(), otherSrc:get_argTypeInfoList(), false, nil, alt2type )
         if argCheck == MatchType.Error then
            return false, argMess
         end
         
         local retCheck, retMess = TypeInfo.checkMatchType( processInfo, dest:get_retTypeInfoList(), otherSrc:get_retTypeInfoList(), false, nil, alt2type )
         if retCheck == MatchType.Error then
            return false, retMess
         end
         
         return true, nil
      elseif _switchExp == TypeInfoKind.Method then
         if #dest:get_argTypeInfoList() ~= #otherSrc:get_argTypeInfoList() or #dest:get_retTypeInfoList() ~= #otherSrc:get_retTypeInfoList() then
            return false, nil
         end
         
         for index, argType in ipairs( dest:get_argTypeInfoList() ) do
            local otherArgType = otherSrc:get_argTypeInfoList()[index]
            if not argType:equals( processInfo, otherArgType, alt2type, true ) then
               local mess = string.format( "unmatch arg(%d) type -- %s, %s", index, argType:getTxt(  ), otherArgType:getTxt(  ))
               return false, mess
            end
            
         end
         
         for index, retType in ipairs( dest:get_retTypeInfoList() ) do
            local otherRetType = otherSrc:get_retTypeInfoList()[index]
            if not retType:equals( processInfo, otherRetType, alt2type, true ) then
               local mess = string.format( "unmatch ret(%d) type -- %s, %s, %s", index, retType:getTxt(  ), otherRetType:getTxt(  ), dest:getTxt(  ))
               return false, mess
            end
            
         end
         
         if dest:get_mutMode() ~= otherSrc:get_mutMode() then
            local mess = string.format( "unmatch mutable mode -- %s, %s", MutMode:_getTxt( dest:get_mutMode())
            , MutMode:_getTxt( otherSrc:get_mutMode())
            )
            return false, mess
         end
         
         return true, nil
      elseif _switchExp == TypeInfoKind.Nilable then
         local dstNonNil
         
         if destMut then
            dstNonNil = dest:get_nonnilableType()
         else
          
            dstNonNil = processInfo:createModifier( dest:get_nonnilableType(), MutMode.IMut )
         end
         
         return dstNonNil:canEvalWith( processInfo, otherSrc:get_nonnilableType(), canEvalType, alt2type )
      elseif _switchExp == TypeInfoKind.Alternate then
         return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      elseif _switchExp == TypeInfoKind.Box then
         return dest:canEvalWith( processInfo, otherSrc, canEvalType, alt2type )
      else 
         
            return false, nil
      end
   end
   
end


function NormalTypeInfo:canEvalWith( processInfo, other, canEvalType, alt2type )

   return TypeInfo.canEvalWithBase( processInfo, self, TypeInfo.isMut( self ), other, canEvalType, alt2type )
end


function ModifierTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local typeInfo = self.srcTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )
   if typeInfo == self.srcTypeInfo then
      return self
   end
   
   if typeInfo ~= nil then
      return moduleTypeInfo:getProcessInfo(  ):createModifier( typeInfo, MutMode.IMut )
   end
   
   return nil
end


function NormalTypeInfo:applyGeneric( alt2typeMap, moduleTypeInfo )

   local itemTypeInfoList, needNew = applyGenericList( self.itemTypeInfoList, alt2typeMap, moduleTypeInfo )
   if  nil == itemTypeInfoList or  nil == needNew then
      local _itemTypeInfoList = itemTypeInfoList
      local _needNew = needNew
   
      return nil
   end
   
   local processInfo = moduleTypeInfo:getProcessInfo(  )
   do
      local _switchExp = self:get_kind()
      if _switchExp == TypeInfoKind.Set then
         if not needNew then
            return self
         end
         
         return processInfo:createSet( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.List then
         if not needNew then
            return self
         end
         
         return processInfo:createList( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.Array then
         if not needNew then
            return self
         end
         
         return processInfo:createArray( self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode )
      elseif _switchExp == TypeInfoKind.Map then
         if not needNew then
            return self
         end
         
         return processInfo:createMap( self.accessMode, self.parentInfo, itemTypeInfoList[1], itemTypeInfoList[2], self.mutMode )
      elseif _switchExp == TypeInfoKind.Func or _switchExp == TypeInfoKind.Form or _switchExp == TypeInfoKind.FormFunc then
         local argTypeInfoList, workArg = applyGenericList( self.argTypeInfoList, alt2typeMap, moduleTypeInfo )
         if  nil == argTypeInfoList or  nil == workArg then
            local _argTypeInfoList = argTypeInfoList
            local _workArg = workArg
         
            return nil
         end
         
         local retTypeInfoList, workRet = applyGenericList( self.retTypeInfoList, alt2typeMap, moduleTypeInfo )
         if  nil == retTypeInfoList or  nil == workRet then
            local _retTypeInfoList = retTypeInfoList
            local _workRet = workRet
         
            return nil
         end
         
         if needNew or workArg or workRet then
            return processInfo:createFunc( self.abstractFlag, false, getScope( self ), self.kind, self.parentInfo, self.autoFlag, self.externalFlag, self.staticFlag, self.accessMode, self.rawTxt, itemTypeInfoList, argTypeInfoList, retTypeInfoList, TypeInfo.isMut( self ) )
         end
         
         return self
      else 
         
            if #self.itemTypeInfoList == 0 then
               return self
            end
            
            return nil
      end
   end
   
end


function TypeInfo:getFullName( typeNameCtrl, importInfo, localFlag )

   if localFlag and self:isModule(  ) then
      return typeNameCtrl:getModuleName( self, "", importInfo )
   end
   
   return self:getParentFullName( typeNameCtrl, importInfo, localFlag ) .. self:get_rawTxt()
end

local builtinTypeInfo2Map = rootProcessInfo:get_typeInfo2Map():clone(  )

local function createProcessInfo( validCheckingMutable, validExtType )

   return ProcessInfo.createUser( validCheckingMutable, validExtType, builtinTypeInfo2Map:clone(  ) )
end
_moduleObj.createProcessInfo = createProcessInfo

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


local compOpSet = {["=="] = true, ["~="] = true}
_moduleObj.compOpSet = compOpSet

local mathCompOpSet = {["<"] = true, ["<="] = true, [">"] = true, [">="] = true}
_moduleObj.mathCompOpSet = mathCompOpSet


local RefTypeInfo = {}
_moduleObj.RefTypeInfo = RefTypeInfo
function RefTypeInfo.setmeta( obj )
  setmetatable( obj, { __index = RefTypeInfo  } )
end
function RefTypeInfo.new( pos, itemRefTypeList, typeInfo )
   local obj = {}
   RefTypeInfo.setmeta( obj )
   if obj.__init then
      obj:__init( pos, itemRefTypeList, typeInfo )
   end
   return obj
end
function RefTypeInfo:__init( pos, itemRefTypeList, typeInfo )

   self.pos = pos
   self.itemRefTypeList = itemRefTypeList
   self.typeInfo = typeInfo
end
function RefTypeInfo:get_pos()
   return self.pos
end
function RefTypeInfo:get_itemRefTypeList()
   return self.itemRefTypeList
end
function RefTypeInfo:get_typeInfo()
   return self.typeInfo
end


local TypeAnalyzer = {}
function TypeAnalyzer.new( processInfo, parentInfo, moduleType, moduleScope, scopeAccess, validMutControl )
   local obj = {}
   TypeAnalyzer.setmeta( obj )
   if obj.__init then obj:__init( processInfo, parentInfo, moduleType, moduleScope, scopeAccess, validMutControl ); end
   return obj
end
function TypeAnalyzer:__init(processInfo, parentInfo, moduleType, moduleScope, scopeAccess, validMutControl) 
   self.processInfo = processInfo
   self.parentInfo = parentInfo
   self.moduleType = moduleType
   self.moduleScope = moduleScope
   self.scopeAccess = scopeAccess
   self.validMutControl = validMutControl
   
   self.scope = _moduleObj.rootScope
   self.accessMode = AccessMode.Local
   self.parentPub = false
   self.parser = Parser.DefaultPushbackParser.new(Parser.DummyParser.new())
end
function TypeAnalyzer:analyzeType( scope, parser, accessMode, allowDDD, parentPub )

   self.scope = scope
   self.parser = parser
   self.accessMode = accessMode
   self.parentPub = parentPub
   return self:analyzeTypeSub( allowDDD )
end
function TypeAnalyzer:analyzeTypeFromTxt( txt, scope, accessMode, parentPub )

   local stream = Parser.TxtStream.new(txt)
   local parser = Parser.DefaultPushbackParser.new(Parser.StreamParser.new(stream, "test"))
   return self:analyzeType( scope, parser, accessMode, true, parentPub )
end
function TypeAnalyzer.setmeta( obj )
  setmetatable( obj, { __index = TypeAnalyzer  } )
end


function TypeAnalyzer:analyzeTypeSub( allowDDD )

   local firstToken = self.parser:getTokenNoErr(  )
   local token = firstToken
   local refFlag = false
   if token.txt == "&" then
      refFlag = true
      token = self.parser:getTokenNoErr(  )
   end
   
   local mutFlag = false
   if token.txt == "mut" then
      mutFlag = true
      token = self.parser:getTokenNoErr(  )
   end
   
   
   local typeInfo
   
   if token.txt == "..." then
      typeInfo = _moduleObj.builtinTypeDDD
   else
    
      local symbol = self.scope:getSymbolTypeInfo( token.txt, self.scope, self.moduleScope, self.scopeAccess )
      if  nil == symbol then
         local _symbol = symbol
      
         return nil, token.pos, "not found type -- " .. token.txt
      end
      
      if symbol:get_kind() ~= SymbolKind.Typ then
         return nil, token.pos, string.format( "illegal type -- %s", symbol:get_name())
      end
      
      typeInfo = symbol:get_typeInfo()
   end
   
   return self:analyzeTypeItemList( allowDDD, refFlag, mutFlag, typeInfo, token.pos )
end


function TypeAnalyzer:analyzeTypeItemList( allowDDD, refFlag, mutFlag, typeInfo, pos )

   if self.parentPub and isPubToExternal( self.accessMode ) and not isPubToExternal( typeInfo:get_accessMode() ) then
      return nil, pos, string.format( "This type must be public. -- %s", typeInfo:getTxt(  ))
   end
   
   
   local token = self.parser:getTokenNoErr(  )
   
   if token.consecutive and token.txt == "!" then
      typeInfo = typeInfo:get_nilableTypeInfo()
      token = self.parser:getTokenNoErr(  )
   end
   
   
   
   
   local genericRefList = {}
   while true do
      if token.txt == '[' or token.txt == '[@' then
         if token.txt == '[' then
            typeInfo = self.processInfo:createList( self.accessMode, self.parentInfo, {typeInfo}, MutMode.Mut )
         else
          
            typeInfo = self.processInfo:createArray( self.accessMode, self.parentInfo, {typeInfo}, MutMode.Mut )
         end
         
         token = self.parser:getTokenNoErr(  )
         
         if token.txt ~= ']' then
            return nil, token.pos, "not found -- ']'"
         end
         
         
      elseif token.txt == "<" then
         local genericList = {}
         local nextToken = Parser.getEofToken(  )
         repeat 
            local refType = self:analyzeTypeSub( false )
            if refType ~= nil then
               table.insert( genericRefList, refType )
               table.insert( genericList, refType:get_typeInfo() )
            end
            
            nextToken = self.parser:getTokenNoErr(  )
         until nextToken.txt ~= ","
         
         if nextToken.txt ~= '>' then
            return nil, nextToken.pos, "not found -- ']'"
         end
         
         
         
         
         
         do
            local _switchExp = typeInfo:get_kind()
            if _switchExp == TypeInfoKind.Map then
               if #genericList ~= 2 then
                  return nil, pos, "Key or value type is unknown"
               else
                
                  typeInfo = self.processInfo:createMap( self.accessMode, self.parentInfo, genericList[1], genericList[2], MutMode.Mut )
               end
               
            elseif _switchExp == TypeInfoKind.List then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               typeInfo = self.processInfo:createList( self.accessMode, self.parentInfo, genericList, MutMode.Mut )
            elseif _switchExp == TypeInfoKind.Array then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               typeInfo = self.processInfo:createArray( self.accessMode, self.parentInfo, genericList, MutMode.Mut )
            elseif _switchExp == TypeInfoKind.Set then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               typeInfo = self.processInfo:createSet( self.accessMode, self.parentInfo, genericList, MutMode.Mut )
            elseif _switchExp == TypeInfoKind.DDD then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               typeInfo = self.processInfo:createDDD( genericList[1], false, false )
            elseif _switchExp == TypeInfoKind.Class or _switchExp == TypeInfoKind.IF then
               
               if #genericList ~= #typeInfo:get_itemTypeInfoList() then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               for __index, itemType in ipairs( genericList ) do
                  if itemType:get_nilable() then
                     local mess = string.format( "can't use nilable type -- %s", itemType:getTxt(  ))
                     return nil, pos, mess
                  end
                  
               end
               
               typeInfo = self.processInfo:createGeneric( typeInfo, genericList, self.moduleType )
            elseif _switchExp == TypeInfoKind.Box then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               typeInfo = self.processInfo:createBox( self.accessMode, genericList[1] )
            elseif _switchExp == TypeInfoKind.Ext then
               
               if #genericList ~= 1 then
                  return nil, pos, string.format( "generic type count is unmatch. -- %d", #genericList)
               end
               
               
               do
                  local _matchExp = self.processInfo:createLuaval( genericList[1], true )
                  if _matchExp[1] == LuavalResult.OK[1] then
                     local extTypeInfo = _matchExp[2][1]
                     local _ = _matchExp[2][2]
                  
                     typeInfo = extTypeInfo
                  elseif _matchExp[1] == LuavalResult.Err[1] then
                     local err = _matchExp[2][1]
                  
                     return nil, pos, err
                  end
               end
               
            else 
               
                  return nil, pos, string.format( "not support generic: %s", typeInfo:getTxt(  ))
            end
         end
         
      else
       
         self.parser:pushback(  )
         break
      end
      
      token = self.parser:getTokenNoErr(  )
   end
   
   if token.txt == "!" then
      typeInfo = typeInfo:get_nilableTypeInfo(  )
      self.parser:getTokenNoErr(  )
   end
   
   
   if not allowDDD then
      if typeInfo:get_kind() == TypeInfoKind.DDD then
         return nil, pos, string.format( "invalid type. -- '%s'", typeInfo:getTxt(  ))
      end
      
   end
   
   
   if refFlag then
      if self.validMutControl then
         typeInfo = self.processInfo:createModifier( typeInfo, MutMode.IMut )
      end
      
   end
   
   
   return RefTypeInfo.new(pos, genericRefList, typeInfo), nil, nil
end




return _moduleObj
