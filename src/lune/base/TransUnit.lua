--lune/base/TransUnit.lns
local _moduleObj = {}
local __mod__ = 'lune.base.TransUnit'
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

      
function _lune._toStem( val )
   return val
end
function _lune._toInt( val )
   if type( val ) == "number" then
      return math.floor( val )
   end
   return nil
end
function _lune._toReal( val )
   if type( val ) == "number" then
      return val
   end
   return nil
end
function _lune._toBool( val )
   if type( val ) == "boolean" then
      return val
   end
   return nil
end
function _lune._toStr( val )
   if type( val ) == "string" then
      return val
   end
   return nil
end
function _lune._toList( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toValInfo = toValInfoList[ 1 ]
      for index, mem in ipairs( val ) do
         local memval, mess = toValInfo.func( mem, toValInfo.child )
         if memval == nil and not toValInfo.nilable then
            if mess then
              return nil, string.format( "%d.%s", index, mess )
            end
            return nil, index
         end
         tbl[ index ] = memval
      end
      return tbl
   end
   return nil
end
function _lune._toMap( val, toValInfoList )
   if type( val ) == "table" then
      local tbl = {}
      local toKeyInfo = toValInfoList[ 1 ]
      local toValInfo = toValInfoList[ 2 ]
      for key, mem in pairs( val ) do
         local mapKey, keySub = toKeyInfo.func( key, toKeyInfo.child )
         local mapVal, valSub = toValInfo.func( mem, toValInfo.child )
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
function _lune._fromMap( obj, map, memInfoList )
   if type( map ) ~= "table" then
      return false
   end
   for index, memInfo in ipairs( memInfoList ) do
      local val, key = memInfo.func( map[ memInfo.name ], memInfo.child )
      if val == nil and not memInfo.nilable then
         return false, key and string.format( "%s.%s", memInfo.name, key) or memInfo.name
      end
      obj[ memInfo.name ] = val
   end
   return true
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

function _lune.loadModule( mod )
   if __luneScript then
      return  __luneScript:loadModule( mod )
   end
   return require( mod )
end




local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Writer = _lune.loadModule( 'lune.base.Writer' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local LuaVer = _lune.loadModule( 'lune.base.LuaVer' )
local DeclClassMode = {}
DeclClassMode._val2NameMap = {}
function DeclClassMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclClassMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function DeclClassMode._from( val )
   if DeclClassMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
DeclClassMode.__allList = {}
function DeclClassMode.get__allList()
   return DeclClassMode.__allList
end

DeclClassMode.Class = 0
DeclClassMode._val2NameMap[0] = 'Class'
DeclClassMode.__allList[1] = DeclClassMode.Class
DeclClassMode.Interface = 1
DeclClassMode._val2NameMap[1] = 'Interface'
DeclClassMode.__allList[2] = DeclClassMode.Interface
DeclClassMode.Module = 2
DeclClassMode._val2NameMap[2] = 'Module'
DeclClassMode.__allList[3] = DeclClassMode.Module

local DeclFuncMode = {}
DeclFuncMode._val2NameMap = {}
function DeclFuncMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "DeclFuncMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function DeclFuncMode._from( val )
   if DeclFuncMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
DeclFuncMode.__allList = {}
function DeclFuncMode.get__allList()
   return DeclFuncMode.__allList
end

DeclFuncMode.Func = 0
DeclFuncMode._val2NameMap[0] = 'Func'
DeclFuncMode.__allList[1] = DeclFuncMode.Func
DeclFuncMode.Class = 1
DeclFuncMode._val2NameMap[1] = 'Class'
DeclFuncMode.__allList[2] = DeclFuncMode.Class
DeclFuncMode.Module = 2
DeclFuncMode._val2NameMap[2] = 'Module'
DeclFuncMode.__allList[3] = DeclFuncMode.Module
DeclFuncMode.Glue = 3
DeclFuncMode._val2NameMap[3] = 'Glue'
DeclFuncMode.__allList[4] = DeclFuncMode.Glue

local ExpSymbolMode = {}
ExpSymbolMode._val2NameMap = {}
function ExpSymbolMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "ExpSymbolMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function ExpSymbolMode._from( val )
   if ExpSymbolMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
ExpSymbolMode.__allList = {}
function ExpSymbolMode.get__allList()
   return ExpSymbolMode.__allList
end

ExpSymbolMode.Symbol = 0
ExpSymbolMode._val2NameMap[0] = 'Symbol'
ExpSymbolMode.__allList[1] = ExpSymbolMode.Symbol
ExpSymbolMode.Fn = 1
ExpSymbolMode._val2NameMap[1] = 'Fn'
ExpSymbolMode.__allList[2] = ExpSymbolMode.Fn
ExpSymbolMode.Field = 2
ExpSymbolMode._val2NameMap[2] = 'Field'
ExpSymbolMode.__allList[3] = ExpSymbolMode.Field
ExpSymbolMode.FieldNil = 3
ExpSymbolMode._val2NameMap[3] = 'FieldNil'
ExpSymbolMode.__allList[4] = ExpSymbolMode.FieldNil
ExpSymbolMode.Get = 4
ExpSymbolMode._val2NameMap[4] = 'Get'
ExpSymbolMode.__allList[5] = ExpSymbolMode.Get
ExpSymbolMode.GetNil = 5
ExpSymbolMode._val2NameMap[5] = 'GetNil'
ExpSymbolMode.__allList[6] = ExpSymbolMode.GetNil

local AnalyzeMode = {}
_moduleObj.AnalyzeMode = AnalyzeMode
AnalyzeMode._val2NameMap = {}
function AnalyzeMode:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "AnalyzeMode.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function AnalyzeMode._from( val )
   if AnalyzeMode._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
AnalyzeMode.__allList = {}
function AnalyzeMode.get__allList()
   return AnalyzeMode.__allList
end

AnalyzeMode.Compile = 0
AnalyzeMode._val2NameMap[0] = 'Compile'
AnalyzeMode.__allList[1] = AnalyzeMode.Compile
AnalyzeMode.Diag = 1
AnalyzeMode._val2NameMap[1] = 'Diag'
AnalyzeMode.__allList[2] = AnalyzeMode.Diag
AnalyzeMode.Complete = 2
AnalyzeMode._val2NameMap[2] = 'Complete'
AnalyzeMode.__allList[3] = AnalyzeMode.Complete

local TransUnit = {}
_moduleObj.TransUnit = TransUnit
function TransUnit.new( importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer )
   local obj = {}
   TransUnit.setmeta( obj )
   if obj.__init then obj:__init( importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer ); end
   return obj
end
function TransUnit:__init(importModuleInfo, macroEval, analyzeModule, mode, pos, targetLuaVer) 
   self.helperInfo = Ast.LuneHelperInfo.new(false, false, false, false, false, false)
   self.targetLuaVer = targetLuaVer
   self.importModuleInfo = importModuleInfo
   self.protoFuncMap = {}
   self.loopScopeQueue = {}
   self.has__func__Symbol = false
   self.nodeManager = Ast.NodeManager.new()
   self.importModuleName2ModuleInfo = {}
   self.importModule2ModuleInfoCurrent = {}
   self.importModule2ModuleInfo = {}
   self.macroScope = nil
   self.validMutControl = true
   self.moduleName = ""
   self.moduleType = Ast.headTypeInfo
   self.parser = Parser.DummyParser.new()
   self.subfileList = {}
   self.pushbackList = {}
   self.usedTokenList = {}
   self.globalScope = Ast.Scope.new(Ast.rootScope, false, nil)
   self.scope = Ast.Scope.new(self.globalScope, true, nil)
   self.topScope = self.scope
   self.moduleScope = self.scope
   self.typeId2ClassMap = {}
   self.typeInfo2ClassNode = {}
   self.currentToken = Parser.getEofToken(  )
   self.errMessList = {}
   self.warnMessList = {}
   self.macroEval = macroEval
   self.typeId2MacroInfo = {}
   self.macroMode = Ast.MacroMode.None
   self.symbol2ValueMapForMacro = {}
   self.analyzeMode = _lune.unwrapDefault( mode, AnalyzeMode.Compile)
   self.analyzePos = _lune.unwrapDefault( pos, Parser.Position.new(0, 0))
   self.analyzeModule = _lune.unwrapDefault( analyzeModule, "")
   self.provideNode = nil
end
function TransUnit:addErrMess( pos, mess )

   table.insert( self.errMessList, string.format( "%s:%d:%d: error: %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, mess) )
end
function TransUnit:addWarnMess( pos, mess )

   table.insert( self.warnMessList, string.format( "%s:%d:%d: warning: %s", self.parser:getStreamName(  ), pos.lineNo, pos.column, mess) )
end
function TransUnit:pushScope( classFlag, inheritScope, ifScopeList )

   self.scope = Ast.Scope.new(self.scope, classFlag, inheritScope, ifScopeList)
   return self.scope
end
function TransUnit:popScope(  )

   self.scope = self.scope:get_parent(  )
end
function TransUnit:getCurrentClass(  )

   local typeInfo = Ast.headTypeInfo
   local scope = self.scope
   repeat 
      do
         local _exp = scope:get_ownerTypeInfo()
         if _exp ~= nil then
            if _exp:get_kind() == Ast.TypeInfoKind.Class or _exp:get_kind() == Ast.TypeInfoKind.Module or _exp:get_kind() == Ast.TypeInfoKind.IF then
               return _exp
            end
            
         end
      end
      
      scope = scope:get_parent()
   until scope:isRoot(  )
   return typeInfo
end
function TransUnit:getCurrentNamespaceTypeInfo(  )

   local typeInfo = Ast.headTypeInfo
   local scope = self.scope
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
function TransUnit:pushModule( externalFlag, name, mutable )

   local typeInfo = Ast.headTypeInfo
   do
      local _exp = self.scope:getTypeInfoChild( name )
      if _exp ~= nil then
         typeInfo = _exp
         self.scope = _lune.unwrap( Ast.getScope( typeInfo ))
      else
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         local parentScope = self.scope
         local scope = self:pushScope( true )
         typeInfo = Ast.NormalTypeInfo.createModule( scope, parentInfo, externalFlag, name, mutable )
         parentScope:addClass( name, typeInfo )
      end
   end
   
   if not self.typeId2ClassMap[typeInfo:get_typeId(  )] then
      local namespace = Ast.NamespaceInfo.new(name, self.scope, typeInfo)
      self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
   end
   
   return typeInfo
end
function TransUnit:popModule(  )

   self:popScope(  )
end
function TransUnit:pushClass( classFlag, abstractFlag, baseInfo, interfaceList, externalFlag, name, accessMode, defNamespace )

   local typeInfo = Ast.headTypeInfo
   do
      local _exp = self.scope:getTypeInfoChild( name )
      if _exp ~= nil then
         typeInfo = _exp
         if typeInfo:get_abstractFlag() ~= abstractFlag then
            self:addErrMess( self.currentToken.pos, "mismatch class abstract for prototpye" )
         end
         
         if typeInfo:get_accessMode() ~= accessMode then
            self:addErrMess( self.currentToken.pos, string.format( "mismatch class accessmode(%s) for prototpye accessmode(%s)", Ast.AccessMode:_getTxt( accessMode)
            , Ast.AccessMode:_getTxt( typeInfo:get_accessMode())
            ) )
         end
         
         self.scope = _lune.unwrap( Ast.getScope( typeInfo ))
         do
            local _switchExp = (typeInfo:get_kind() )
            if _switchExp == Ast.TypeInfoKind.Class then
               if not classFlag then
                  self:addErrMess( self.currentToken.pos, string.format( "define interface already -- %s", name) )
                  Util.printStackTrace(  )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.IF then
               if classFlag then
                  self:addErrMess( self.currentToken.pos, string.format( "define class already -- %s", name) )
                  Util.printStackTrace(  )
               end
               
            end
         end
         
      else
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         local inheritScope = nil
         if baseInfo ~= nil then
            inheritScope = _lune.unwrap( baseInfo:get_scope())
         end
         
         local ifScopeList = {}
         if interfaceList ~= nil then
            for __index, ifType in pairs( interfaceList ) do
               table.insert( ifScopeList, _lune.unwrap( ifType:get_scope(  )) )
            end
            
         end
         
         local parentScope = self.scope
         local scope = self:pushScope( true, inheritScope, ifScopeList )
         typeInfo = Ast.NormalTypeInfo.createClass( classFlag, abstractFlag, scope, baseInfo, interfaceList, parentInfo, externalFlag, accessMode, name )
         parentScope:addClass( name, typeInfo )
      end
   end
   
   local namespace = defNamespace
   if  nil == namespace then
      local _namespace = namespace
   
      namespace = Ast.NamespaceInfo.new(name, self.scope, typeInfo)
   end
   
   self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
   return typeInfo
end
function TransUnit:popClass(  )

   self:popScope(  )
end
function TransUnit:addLocalVar( pos, argFlag, canBeLeft, name, typeInfo, mutable, allowShadow )

   if not allowShadow then
      if self.scope:getSymbolTypeInfo( name, self.scope, self.moduleScope ) then
         self:addErrMess( pos, string.format( "shadowing variable -- %s", name) )
      end
      
   end
   
   self.scope:addLocalVar( argFlag, canBeLeft, name, typeInfo, mutable )
end
function TransUnit.setmeta( obj )
  setmetatable( obj, { __index = TransUnit  } )
end
function TransUnit:get_errMessList()       
   return self.errMessList         
end
function TransUnit:get_warnMessList()       
   return self.warnMessList         
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
regOpLevel( 2, {"|<<", "|>>"} )
regOpLevel( 2, {".."} )
regOpLevel( 2, {"+", "-"} )
regOpLevel( 2, {"*", "/", "//", "%"} )
regOpLevel( 1, {"`", ",,", ",,,", ",,,,"} )
regOpLevel( 1, {"not", "#", "-", "~"} )
regOpLevel( 1, {"^"} )
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
function TransUnit:createModifier( typeInfo, mutable )

   if not self.validMutControl then
      return typeInfo
   end
   
   return Ast.NormalTypeInfo.createModifier( typeInfo, mutable )
end

local _MetaInfo = {}
_moduleObj._MetaInfo = _MetaInfo
function _MetaInfo.setmeta( obj )
  setmetatable( obj, { __index = _MetaInfo  } )
end
function _MetaInfo.new( __formatVersion, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap, _moduleTypeId, _moduleSymbolKind, _moduleMutable, _dependModuleMap, _dependIdMap )
   local obj = {}
   _MetaInfo.setmeta( obj )
   if obj.__init then
      obj:__init( __formatVersion, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap, _moduleTypeId, _moduleSymbolKind, _moduleMutable, _dependModuleMap, _dependIdMap )
   end        
   return obj 
end         
function _MetaInfo:__init( __formatVersion, _typeId2ClassInfoMap, _typeInfoList, _varName2InfoMap, _funcName2InfoMap, _moduleTypeId, _moduleSymbolKind, _moduleMutable, _dependModuleMap, _dependIdMap ) 

   self.__formatVersion = __formatVersion
   self._typeId2ClassInfoMap = _typeId2ClassInfoMap
   self._typeInfoList = _typeInfoList
   self._varName2InfoMap = _varName2InfoMap
   self._funcName2InfoMap = _funcName2InfoMap
   self._moduleTypeId = _moduleTypeId
   self._moduleSymbolKind = _moduleSymbolKind
   self._moduleMutable = _moduleMutable
   self._dependModuleMap = _dependModuleMap
   self._dependIdMap = _dependIdMap
end

local ImportParam = {}
function ImportParam.setmeta( obj )
  setmetatable( obj, { __index = ImportParam  } )
end
function ImportParam.new( transUnit, typeId2Scope, typeId2TypeInfo, metaInfo, scope )
   local obj = {}
   ImportParam.setmeta( obj )
   if obj.__init then
      obj:__init( transUnit, typeId2Scope, typeId2TypeInfo, metaInfo, scope )
   end        
   return obj 
end         
function ImportParam:__init( transUnit, typeId2Scope, typeId2TypeInfo, metaInfo, scope ) 

   self.transUnit = transUnit
   self.typeId2Scope = typeId2Scope
   self.typeId2TypeInfo = typeId2TypeInfo
   self.metaInfo = metaInfo
   self.scope = scope
end

local _TypeInfo = {}
function _TypeInfo.new(  )
   local obj = {}
   _TypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function _TypeInfo:__init() 
   self.parentId = Ast.rootTypeId
   self.typeId = Ast.rootTypeId
   self.skind = Ast.SerializeKind.Normal
end
function _TypeInfo.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfo  } )
end
function _TypeInfo:_toMap()
  return self
end
function _TypeInfo._fromMap( val )
  local obj, mes = _TypeInfo._fromMapSub( {}, val )
  if obj then
     _TypeInfo.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfo._fromStem( val )
  return _TypeInfo._fromMap( val )
end

function _TypeInfo._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "skind", func = Ast.SerializeKind._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "typeId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoNilable = {}
setmetatable( _TypeInfoNilable, { __index = _TypeInfo } )
function _TypeInfoNilable:createTypeInfo( param )

   local orgTypeInfo = _lune.unwrap( param.typeId2TypeInfo[self.orgTypeId])
   local newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
   param.typeId2TypeInfo[self.typeId] = _lune.unwrap( newTypeInfo)
   return newTypeInfo, nil
end
function _TypeInfoNilable.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoNilable  } )
end
function _TypeInfoNilable.new( nilable, orgTypeId )
   local obj = {}
   _TypeInfoNilable.setmeta( obj )
   if obj.__init then
      obj:__init( nilable, orgTypeId )
   end        
   return obj 
end         
function _TypeInfoNilable:__init( nilable, orgTypeId ) 

   _TypeInfo.__init( self )
   self.nilable = nilable
   self.orgTypeId = orgTypeId
end
function _TypeInfoNilable:_toMap()
  return self
end
function _TypeInfoNilable._fromMap( val )
  local obj, mes = _TypeInfoNilable._fromMapSub( {}, val )
  if obj then
     _TypeInfoNilable.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoNilable._fromStem( val )
  return _TypeInfoNilable._fromMap( val )
end

function _TypeInfoNilable._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "nilable", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "orgTypeId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoModifier = {}
setmetatable( _TypeInfoModifier, { __index = _TypeInfo } )
function _TypeInfoModifier:createTypeInfo( param )

   local srcTypeInfo = param.typeId2TypeInfo[self.srcTypeId]
   if  nil == srcTypeInfo then
      local _srcTypeInfo = srcTypeInfo
   
      return nil, string.format( "not found srcType -- %d, %d", self.parentId, self.srcTypeId)
   end
   
   local newTypeInfo = param.transUnit:createModifier( srcTypeInfo, _lune.unwrapDefault( self.mutable, false) )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoModifier.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoModifier  } )
end
function _TypeInfoModifier.new( srcTypeId, mutable )
   local obj = {}
   _TypeInfoModifier.setmeta( obj )
   if obj.__init then
      obj:__init( srcTypeId, mutable )
   end        
   return obj 
end         
function _TypeInfoModifier:__init( srcTypeId, mutable ) 

   _TypeInfo.__init( self )
   self.srcTypeId = srcTypeId
   self.mutable = mutable
end
function _TypeInfoModifier:_toMap()
  return self
end
function _TypeInfoModifier._fromMap( val )
  local obj, mes = _TypeInfoModifier._fromMapSub( {}, val )
  if obj then
     _TypeInfoModifier.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoModifier._fromStem( val )
  return _TypeInfoModifier._fromMap( val )
end

function _TypeInfoModifier._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "srcTypeId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "mutable", func = _lune._toBool, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoModule = {}
setmetatable( _TypeInfoModule, { __index = _TypeInfo } )
function _TypeInfoModule:createTypeInfo( param )

   local parentInfo = Ast.headTypeInfo
   if self.parentId ~= Ast.rootTypeId then
      local workTypeInfo = param.typeId2TypeInfo[self.parentId]
      if  nil == workTypeInfo then
         local _workTypeInfo = workTypeInfo
      
         Util.err( string.format( "not found parentInfo %s %s", self.parentId, self.txt) )
      end
      
      parentInfo = workTypeInfo
   end
   
   local parentScope = param.typeId2Scope[self.parentId]
   if  nil == parentScope then
      local _parentScope = parentScope
   
      return nil, string.format( "not found parentScope %s %s", self.parentId, self.txt)
   end
   
   local newTypeInfo = parentScope:getTypeInfoChild( self.txt )
   do
      local _exp = newTypeInfo
      if _exp ~= nil then
         param.typeId2Scope[self.typeId] = _exp:get_scope()
         if not _exp:get_scope() then
            return nil, string.format( "not found scope %s %s %s %s %s", parentScope, self.parentId, self.typeId, self.txt, _exp:getTxt(  ))
         end
         
         param.typeId2TypeInfo[self.typeId] = _exp
      else
         local scope = Ast.Scope.new(parentScope, true, nil)
         local mutable = false
         if self.typeId == param.metaInfo._moduleTypeId then
            mutable = param.metaInfo._moduleMutable
         end
         
         local workTypeInfo = Ast.NormalTypeInfo.createModule( scope, parentInfo, true, self.txt, mutable )
         newTypeInfo = workTypeInfo
         param.typeId2Scope[self.typeId] = scope
         param.typeId2TypeInfo[self.typeId] = workTypeInfo
         parentScope:addClass( self.txt, workTypeInfo )
      end
   end
   
   return newTypeInfo, nil
end
function _TypeInfoModule.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoModule  } )
end
function _TypeInfoModule.new( txt )
   local obj = {}
   _TypeInfoModule.setmeta( obj )
   if obj.__init then
      obj:__init( txt )
   end        
   return obj 
end         
function _TypeInfoModule:__init( txt ) 

   _TypeInfo.__init( self )
   self.txt = txt
end
function _TypeInfoModule:_toMap()
  return self
end
function _TypeInfoModule._fromMap( val )
  local obj, mes = _TypeInfoModule._fromMapSub( {}, val )
  if obj then
     _TypeInfoModule.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoModule._fromStem( val )
  return _TypeInfoModule._fromMap( val )
end

function _TypeInfoModule._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoNormal = {}
setmetatable( _TypeInfoNormal, { __index = _TypeInfo } )
function _TypeInfoNormal:createTypeInfo( param )

   local newTypeInfo = nil
   if self.parentId ~= Ast.rootTypeId or not Ast.builtInTypeIdSet[self.typeId] or self.kind == Ast.TypeInfoKind.List or self.kind == Ast.TypeInfoKind.Array or self.kind == Ast.TypeInfoKind.Map then
      local parentInfo = Ast.headTypeInfo
      if self.parentId ~= Ast.rootTypeId then
         local workTypeInfo = param.typeId2TypeInfo[self.parentId]
         if  nil == workTypeInfo then
            local _workTypeInfo = workTypeInfo
         
            return nil, string.format( "not found parentInfo %s %s", self.parentId, self.txt)
         end
         
         parentInfo = workTypeInfo
      end
      
      local itemTypeInfo = {}
      for __index, typeId in pairs( self.itemTypeId ) do
         table.insert( itemTypeInfo, _lune.unwrap( param.typeId2TypeInfo[typeId]) )
      end
      
      local argTypeInfo = {}
      for index, typeId in pairs( self.argTypeId ) do
         do
            local argType = param.typeId2TypeInfo[typeId]
            if argType ~= nil then
               table.insert( argTypeInfo, argType )
            else
               local mess = string.format( "not found arg (index:%d) -- %s.%s, %d, %d", index, parentInfo:getTxt(  ), self.txt, typeId, #self.argTypeId)
               return nil, mess
            end
         end
         
      end
      
      local retTypeInfo = {}
      for __index, typeId in pairs( self.retTypeId ) do
         table.insert( retTypeInfo, _lune.unwrap( param.typeId2TypeInfo[typeId]) )
      end
      
      local baseInfo = _lune.unwrap( param.typeId2TypeInfo[self.baseId])
      local interfaceList = {}
      for __index, ifTypeId in pairs( self.ifList ) do
         table.insert( interfaceList, _lune.unwrap( param.typeId2TypeInfo[ifTypeId]) )
      end
      
      local parentScope = param.typeId2Scope[self.parentId]
      if  nil == parentScope then
         local _parentScope = parentScope
      
         return nil, string.format( "not found parentScope %s %s", self.parentId, self.txt)
      end
      
      if self.txt ~= "" then
         newTypeInfo = parentScope:getTypeInfoChild( self.txt )
      end
      
      if newTypeInfo and (self.kind == Ast.TypeInfoKind.Class or self.kind == Ast.TypeInfoKind.IF ) then
         do
            local _exp = newTypeInfo
            if _exp ~= nil then
               param.typeId2Scope[self.typeId] = _exp:get_scope()
               if not _exp:get_scope() then
                  Util.err( string.format( "not found scope %s %s %s %s %s", parentScope, self.parentId, self.typeId, self.txt, _exp:getTxt(  )) )
               end
               
               param.typeId2TypeInfo[self.typeId] = _exp
            end
         end
         
         
      else
       
         if self.kind == Ast.TypeInfoKind.Class or self.kind == Ast.TypeInfoKind.IF then
            local baseScope = _lune.unwrap( param.typeId2Scope[self.baseId])
            local scope = Ast.Scope.new(parentScope, true, baseScope)
            local workTypeInfo = Ast.NormalTypeInfo.createClass( self.kind == Ast.TypeInfoKind.Class, self.abstractFlag, scope, baseInfo, interfaceList, parentInfo, true, Ast.AccessMode.Pub, self.txt )
            newTypeInfo = workTypeInfo
            param.typeId2Scope[self.typeId] = scope
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
            parentScope:addClass( self.txt, workTypeInfo )
         else
          
            local scope = nil
            if self.kind == Ast.TypeInfoKind.Func or self.kind == Ast.TypeInfoKind.Method then
               scope = Ast.Scope.new(parentScope, false, nil)
            end
            
            local typeInfoKind = _lune.unwrap( Ast.TypeInfoKind._from( self.kind ))
            local mutable = true
            do
               local _exp = self.mutable
               if _exp ~= nil then
                  if not _exp then
                     mutable = false
                  end
                  
               end
            end
            
            local accessMode = _lune.unwrap( Ast.AccessMode._from( self.accessMode ))
            local workTypeInfo = Ast.NormalTypeInfo.create( accessMode, self.abstractFlag, scope, baseInfo, interfaceList, parentInfo, self.staticFlag, typeInfoKind, self.txt, itemTypeInfo, argTypeInfo, retTypeInfo, mutable )
            newTypeInfo = workTypeInfo
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
            if self.kind == Ast.TypeInfoKind.Func or self.kind == Ast.TypeInfoKind.Method then
               local symbolKind = Ast.SymbolKind.Fun
               if self.kind == Ast.TypeInfoKind.Method then
                  symbolKind = Ast.SymbolKind.Mtd
               end
               
               local workParentScope = _lune.unwrap( param.typeId2Scope[self.parentId])
               workParentScope:add( symbolKind, false, self.kind == Ast.TypeInfoKind.Func, self.txt, workTypeInfo, accessMode, self.staticFlag, false, true )
               param.typeId2Scope[self.typeId] = scope
            end
            
         end
         
      end
      
   else
    
      newTypeInfo = param.scope:getTypeInfo( self.txt, param.scope, false )
      if not newTypeInfo then
         for key, val in pairs( self:_toMap(  ) ) do
            Util.errorLog( string.format( "error: illegal self %s:%s", key, val) )
         end
         
      end
      
      param.typeId2TypeInfo[self.typeId] = _lune.unwrap( newTypeInfo)
   end
   
   return newTypeInfo, nil
end
function _TypeInfoNormal.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoNormal  } )
end
function _TypeInfoNormal.new( abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutable, ifList, itemTypeId, argTypeId, retTypeId, children )
   local obj = {}
   _TypeInfoNormal.setmeta( obj )
   if obj.__init then
      obj:__init( abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutable, ifList, itemTypeId, argTypeId, retTypeId, children )
   end        
   return obj 
end         
function _TypeInfoNormal:__init( abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutable, ifList, itemTypeId, argTypeId, retTypeId, children ) 

   _TypeInfo.__init( self )
   self.abstractFlag = abstractFlag
   self.baseId = baseId
   self.txt = txt
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.kind = kind
   self.mutable = mutable
   self.ifList = ifList
   self.itemTypeId = itemTypeId
   self.argTypeId = argTypeId
   self.retTypeId = retTypeId
   self.children = children
end
function _TypeInfoNormal:_toMap()
  return self
end
function _TypeInfoNormal._fromMap( val )
  local obj, mes = _TypeInfoNormal._fromMapSub( {}, val )
  if obj then
     _TypeInfoNormal.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoNormal._fromStem( val )
  return _TypeInfoNormal._fromMap( val )
end

function _TypeInfoNormal._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "abstractFlag", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "baseId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "staticFlag", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "kind", func = Ast.TypeInfoKind._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "mutable", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "ifList", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "itemTypeId", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "argTypeId", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "retTypeId", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "children", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoEnum = {}
setmetatable( _TypeInfoEnum, { __index = _TypeInfo } )
function _TypeInfoEnum:createTypeInfo( param )

   local accessMode = _lune.unwrap( Ast.AccessMode._from( self.accessMode ))
   local parentInfo = _lune.unwrap( param.typeId2TypeInfo[self.parentId])
   local name2EnumValInfo = {}
   local parentScope = _lune.unwrap( Ast.getScope( parentInfo ))
   local scope = Ast.Scope.new(parentScope, true, nil)
   param.typeId2Scope[self.typeId] = scope
   local enumTypeInfo = Ast.NormalTypeInfo.createEnum( scope, _lune.unwrap( parentInfo), true, accessMode, self.txt, _lune.unwrap( param.typeId2TypeInfo[self.valTypeId]), name2EnumValInfo )
   local newTypeInfo = enumTypeInfo
   param.typeId2TypeInfo[self.typeId] = enumTypeInfo
   for valName, valData in pairs( self.enumValList ) do
      name2EnumValInfo[valName] = Ast.EnumValInfo.new(valName, valData)
      scope:addEnumVal( valName, enumTypeInfo )
   end
   
   parentScope:addEnum( accessMode, self.txt, enumTypeInfo )
   return newTypeInfo, nil
end
function _TypeInfoEnum.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoEnum  } )
end
function _TypeInfoEnum.new( txt, accessMode, valTypeId, enumValList )
   local obj = {}
   _TypeInfoEnum.setmeta( obj )
   if obj.__init then
      obj:__init( txt, accessMode, valTypeId, enumValList )
   end        
   return obj 
end         
function _TypeInfoEnum:__init( txt, accessMode, valTypeId, enumValList ) 

   _TypeInfo.__init( self )
   self.txt = txt
   self.accessMode = accessMode
   self.valTypeId = valTypeId
   self.enumValList = enumValList
end
function _TypeInfoEnum:_toMap()
  return self
end
function _TypeInfoEnum._fromMap( val )
  local obj, mes = _TypeInfoEnum._fromMapSub( {}, val )
  if obj then
     _TypeInfoEnum.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoEnum._fromStem( val )
  return _TypeInfoEnum._fromMap( val )
end

function _TypeInfoEnum._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "valTypeId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "enumValList", func = _lune._toMap, nilable = false, child = { { func = _lune._toStr, nilable = false, child = {} }, 
{ func = _lune._toStem, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoAlgeVal = {}
function _TypeInfoAlgeVal.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoAlgeVal  } )
end
function _TypeInfoAlgeVal.new( name, typeList )
   local obj = {}
   _TypeInfoAlgeVal.setmeta( obj )
   if obj.__init then
      obj:__init( name, typeList )
   end        
   return obj 
end         
function _TypeInfoAlgeVal:__init( name, typeList ) 

   self.name = name
   self.typeList = typeList
end
function _TypeInfoAlgeVal:_toMap()
  return self
end
function _TypeInfoAlgeVal._fromMap( val )
  local obj, mes = _TypeInfoAlgeVal._fromMapSub( {}, val )
  if obj then
     _TypeInfoAlgeVal.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoAlgeVal._fromStem( val )
  return _TypeInfoAlgeVal._fromMap( val )
end

function _TypeInfoAlgeVal._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "name", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "typeList", func = _lune._toList, nilable = false, child = { { func = _lune._toInt, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local _TypeInfoAlge = {}
setmetatable( _TypeInfoAlge, { __index = _TypeInfo } )
function _TypeInfoAlge:createTypeInfo( param )

   local accessMode = _lune.unwrap( Ast.AccessMode._from( self.accessMode ))
   local parentInfo = _lune.unwrap( param.typeId2TypeInfo[self.parentId])
   local name2AlgeValInfo = {}
   local parentScope = _lune.unwrap( Ast.getScope( parentInfo ))
   local scope = Ast.Scope.new(parentScope, true, nil)
   param.typeId2Scope[self.typeId] = scope
   local algeTypeInfo = Ast.NormalTypeInfo.createAlge( scope, _lune.unwrap( parentInfo), true, accessMode, self.txt )
   local newTypeInfo = algeTypeInfo
   param.typeId2TypeInfo[self.typeId] = algeTypeInfo
   for __index, valInfo in pairs( self.algeValList ) do
      local typeInfoList = {}
      for __index, orgTypeId in pairs( valInfo.typeList ) do
         table.insert( typeInfoList, _lune.unwrap( param.typeId2TypeInfo[orgTypeId]) )
      end
      
      local algeVal = Ast.AlgeValInfo.new(valInfo.name, typeInfoList)
      scope:addAlgeVal( valInfo.name, algeTypeInfo )
      algeTypeInfo:addValInfo( algeVal )
   end
   
   parentScope:addAlge( accessMode, self.txt, algeTypeInfo )
   return newTypeInfo, nil
end
function _TypeInfoAlge.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoAlge  } )
end
function _TypeInfoAlge.new( txt, accessMode, algeValList )
   local obj = {}
   _TypeInfoAlge.setmeta( obj )
   if obj.__init then
      obj:__init( txt, accessMode, algeValList )
   end        
   return obj 
end         
function _TypeInfoAlge:__init( txt, accessMode, algeValList ) 

   _TypeInfo.__init( self )
   self.txt = txt
   self.accessMode = accessMode
   self.algeValList = algeValList
end
function _TypeInfoAlge:_toMap()
  return self
end
function _TypeInfoAlge._fromMap( val )
  local obj, mes = _TypeInfoAlge._fromMapSub( {}, val )
  if obj then
     _TypeInfoAlge.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoAlge._fromStem( val )
  return _TypeInfoAlge._fromMap( val )
end

function _TypeInfoAlge._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "algeValList", func = _lune._toList, nilable = false, child = { { func = _TypeInfoAlgeVal._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end

local typeInfoLuneLoad = Ast.headTypeInfo
_moduleObj.typeInfoLuneLoad = typeInfoLuneLoad

local typeInfoListInsert = Ast.headTypeInfo
_moduleObj.typeInfoListInsert = typeInfoListInsert

local typeInfoListRemove = Ast.headTypeInfo
_moduleObj.typeInfoListRemove = typeInfoListRemove

local typeInfoListUnpack = Ast.headTypeInfo
_moduleObj.typeInfoListUnpack = typeInfoListUnpack

local typeInfoArrayUnpack = Ast.headTypeInfo
_moduleObj.typeInfoArrayUnpack = typeInfoArrayUnpack

local typeInfoMappingIF = Ast.headTypeInfo
_moduleObj.typeInfoMappingIF = typeInfoMappingIF

local typeInfoMappingToMap = Ast.headTypeInfo
_moduleObj.typeInfoMappingToMap = typeInfoMappingToMap

local typeInfoStrGMatch = Ast.headTypeInfo
local typeInfoStringGMatch = Ast.headTypeInfo
local typeInfoStringForm = Ast.headTypeInfo
local function isStrFormFunc( typeInfo )

   if typeInfo:equals( typeInfoStringForm ) then
      return true
   end
   
   return false
end
_moduleObj.isStrFormFunc = isStrFormFunc
local function setupBuiltinTypeInfo( name, fieldName, typeInfo )

   do
      local _switchExp = name
      if _switchExp == "" then
         do
            local _switchExp = fieldName
            if _switchExp == "_load" then
               _moduleObj.typeInfoLuneLoad = typeInfo
            end
         end
         
      elseif _switchExp == "List" then
         do
            local _switchExp = fieldName
            if _switchExp == "insert" then
               _moduleObj.typeInfoListInsert = typeInfo
            elseif _switchExp == "remove" then
               _moduleObj.typeInfoListRemove = typeInfo
            elseif _switchExp == "unpack" then
               _moduleObj.typeInfoListUnpack = typeInfo
            end
         end
         
      elseif _switchExp == "Array" then
         do
            local _switchExp = fieldName
            if _switchExp == "unpack" then
               _moduleObj.typeInfoArrayUnpack = typeInfo
            end
         end
         
      elseif _switchExp == "Mapping" then
         do
            local _switchExp = fieldName
            if _switchExp == "_toMap" then
               _moduleObj.typeInfoMappingToMap = typeInfo
            end
         end
         
      elseif _switchExp == "str" then
         do
            local _switchExp = fieldName
            if _switchExp == "gmatch" then
               typeInfoStrGMatch = typeInfo
            end
         end
         
      elseif _switchExp == "string" then
         do
            local _switchExp = fieldName
            if _switchExp == "gmatch" then
               typeInfoStringGMatch = typeInfo
            elseif _switchExp == "format" then
               typeInfoStringForm = typeInfo
            end
         end
         
      end
   end
   
end

local readyBuiltin = false
function TransUnit:registBuiltInScope(  )

   if readyBuiltin then
      return 
   end
   
   readyBuiltin = true
   local builtInInfo = {{[""] = {["type"] = {["arg"] = {"&stem!"}, ["ret"] = {"str"}}, ["error"] = {["arg"] = {"str"}, ["ret"] = {"__"}}, ["print"] = {["arg"] = {"&..."}, ["ret"] = {}}, ["tonumber"] = {["arg"] = {"str", "int!"}, ["ret"] = {"real"}}, ["tostring"] = {["arg"] = {"&stem"}, ["ret"] = {"str"}}, ["load"] = {["arg"] = {"str", "str!", "str!", "stem!"}, ["ret"] = {"form!", "str!"}}, ["loadfile"] = {["arg"] = {"str"}, ["ret"] = {"form!", "str!"}}, ["require"] = {["arg"] = {"str"}, ["ret"] = {"stem!"}}, ["collectgarbage"] = {["arg"] = {}, ["ret"] = {}}, ["_fcall"] = {["arg"] = {"form", "&..."}, ["ret"] = {""}}, ["_load"] = {["arg"] = {"str", "stem!"}, ["ret"] = {"form!", "str!"}}}}, {["iStream"] = {["__attrib"] = {["type"] = {"interface"}}, ["read"] = {["type"] = {"mut"}, ["arg"] = {"&stem!"}, ["ret"] = {"str!"}}, ["close"] = {["type"] = {"mut"}, ["arg"] = {}, ["ret"] = {}}}}, {["oStream"] = {["__attrib"] = {["type"] = {"interface"}}, ["write"] = {["type"] = {"mut"}, ["arg"] = {"str"}, ["ret"] = {"stem!", "str!"}}, ["close"] = {["type"] = {"mut"}, ["arg"] = {}, ["ret"] = {}}, ["flush"] = {["type"] = {"mut"}, ["arg"] = {}, ["ret"] = {}}}}, {["luaStream"] = {["__attrib"] = {["inplements"] = {"iStream", "oStream"}}, ["read"] = {["type"] = {"mut"}, ["arg"] = {"&stem!"}, ["ret"] = {"str!"}}, ["write"] = {["type"] = {"mut"}, ["arg"] = {"str"}, ["ret"] = {"stem!", "str!"}}, ["close"] = {["type"] = {"mut"}, ["arg"] = {}, ["ret"] = {}}, ["flush"] = {["type"] = {"mut"}, ["arg"] = {}, ["ret"] = {}}, ["seek"] = {["type"] = {"mut"}, ["arg"] = {"str", "int"}, ["ret"] = {"int!", "str!"}}}}, {["Mapping"] = {["__attrib"] = {["type"] = {"interface"}}, ["_toMap"] = {["type"] = {"method"}, ["arg"] = {}, ["ret"] = {}}}}, {["io"] = {["stdin"] = {["type"] = {"member"}, ["typeInfo"] = {"iStream"}}, ["stdout"] = {["type"] = {"member"}, ["typeInfo"] = {"oStream"}}, ["stderr"] = {["type"] = {"member"}, ["typeInfo"] = {"oStream"}}, ["open"] = {["arg"] = {"str", "str!"}, ["ret"] = {"luaStream!"}}, ["popen"] = {["arg"] = {"str"}, ["ret"] = {"luaStream!"}}}}, {["package"] = {["path"] = {["type"] = {"member"}, ["typeInfo"] = {"str"}}, ["searchpath"] = {["arg"] = {"str", "str"}, ["ret"] = {"str!"}}}}, {["os"] = {["clock"] = {["arg"] = {}, ["ret"] = {"int"}}, ["exit"] = {["arg"] = {"int!"}, ["ret"] = {"__"}}, ["remove"] = {["arg"] = {"str"}, ["ret"] = {"bool!", "str!"}}, ["date"] = {["arg"] = {"str!", "stem!"}, ["ret"] = {"stem!"}}, ["time"] = {["arg"] = {"stem!"}, ["ret"] = {"stem!"}}, ["difftime"] = {["arg"] = {"stem", "stem"}, ["ret"] = {"int"}}}}, {["string"] = {["find"] = {["arg"] = {"str", "str", "int!", "bool!"}, ["ret"] = {"int!", "int!"}}, ["byte"] = {["arg"] = {"str", "int"}, ["ret"] = {"int"}}, ["format"] = {["arg"] = {"str", "..."}, ["ret"] = {"str"}}, ["rep"] = {["arg"] = {"str", "int"}, ["ret"] = {"str"}}, ["gmatch"] = {["arg"] = {"str", "str"}, ["ret"] = {"form", "stem!", "stem!"}}, ["gsub"] = {["arg"] = {"str", "str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["arg"] = {"str", "int", "int!"}, ["ret"] = {"str"}}}}, {["str"] = {["find"] = {["type"] = {"method"}, ["arg"] = {"str", "int!", "bool!"}, ["ret"] = {"int!", "int!"}}, ["byte"] = {["type"] = {"method"}, ["arg"] = {"int"}, ["ret"] = {"int"}}, ["format"] = {["type"] = {"method"}, ["arg"] = {"&..."}, ["ret"] = {"str"}}, ["rep"] = {["type"] = {"method"}, ["arg"] = {"int"}, ["ret"] = {"str"}}, ["gmatch"] = {["type"] = {"method"}, ["arg"] = {"str"}, ["ret"] = {"form", "stem!", "stem!"}}, ["gsub"] = {["type"] = {"method"}, ["arg"] = {"str", "str"}, ["ret"] = {"str"}}, ["sub"] = {["type"] = {"method"}, ["arg"] = {"int", "int!"}, ["ret"] = {"str"}}}}, {["List"] = {["insert"] = {["type"] = {"mut"}, ["arg"] = {"&stem"}, ["ret"] = {""}}, ["remove"] = {["type"] = {"mut"}, ["arg"] = {"int!"}, ["ret"] = {""}}, ["unpack"] = {["type"] = {"method"}, ["arg"] = {}, ["ret"] = {"..."}}}}, {["Array"] = {["unpack"] = {["type"] = {"method"}, ["arg"] = {}, ["ret"] = {"..."}}}}, {["debug"] = {["getinfo"] = {["arg"] = {"int"}, ["ret"] = {"stem!"}}, ["getlocal"] = {["arg"] = {"int", "int"}, ["ret"] = {"str!", "stem!"}}}}}
   local function getTypeInfo( typeName )
   
      local mutable = true
      if typeName:find( "^&" ) then
         mutable = false
         typeName = typeName:gsub( "^&", "" )
      end
      
      local typeInfo = Ast.headTypeInfo
      if typeName:find( "!$" ) then
         local orgTypeName = typeName:gsub( "!$", "" )
         typeInfo = _lune.unwrap( self.scope:getTypeInfo( orgTypeName, self.scope, false ))
         typeInfo = typeInfo:get_nilableTypeInfo()
      else
       
         typeInfo = _lune.unwrap( self.scope:getTypeInfo( typeName, self.scope, false ))
      end
      
      if mutable then
         return typeInfo
      end
      
      typeInfo = self:createModifier( typeInfo, false )
      return typeInfo
   end
   
   self.scope = Ast.rootScope
   local builtinModuleName2Scope = {}
   local mapType = Ast.NormalTypeInfo.createMap( Ast.AccessMode.Pub, Ast.headTypeInfo, Ast.builtinTypeString, Ast.builtinTypeStem )
   self.scope:addVar( Ast.AccessMode.Global, "_ENV", mapType, false, true )
   self.scope:addVar( Ast.AccessMode.Global, "_G", mapType, false, true )
   self.scope:addVar( Ast.AccessMode.Global, "_VERSION", Ast.builtinTypeString, false, true )
   self.scope:addVar( Ast.AccessMode.Global, "__mod__", Ast.builtinTypeString, false, true )
   self.scope:addVar( Ast.AccessMode.Global, "__line__", Ast.builtinTypeInt, false, true )
   self.scope:addVar( Ast.AccessMode.Global, "__func__", Ast.builtinTypeString, false, true )
   for __index, builtinClassInfo in pairs( builtInInfo ) do
      for name, name2FieldInfo in pairs( builtinClassInfo ) do
         local parentInfo = Ast.headTypeInfo
         if name ~= "" then
            local classFlag = true
            if _lune.nilacc( _lune.nilacc( name2FieldInfo['__attrib'], nil, 'item', 'type'), nil, 'item', 1) == "interface" then
               classFlag = false
            end
            
            local interfaceList = {}
            do
               local _exp = _lune.nilacc( name2FieldInfo['__attrib'], nil, 'item', 'inplements')
               if _exp ~= nil then
                  for __index, ifname in pairs( _exp ) do
                     local ifType = getTypeInfo( ifname )
                     table.insert( interfaceList, ifType )
                  end
                  
               end
            end
            
            parentInfo = self:pushClass( classFlag, false, nil, interfaceList, true, name, Ast.AccessMode.Pub )
            Ast.builtInTypeIdSet[parentInfo:get_typeId(  )] = parentInfo
            Ast.builtInTypeIdSet[parentInfo:get_nilableTypeInfo():get_typeId()] = parentInfo:get_nilableTypeInfo()
            if name == "Mapping" then
               _moduleObj.typeInfoMappingIF = parentInfo
            end
            
         end
         
         if not builtinModuleName2Scope[name] then
            if name ~= "" and getTypeInfo( name ) then
               builtinModuleName2Scope[name] = self.scope
            end
            
            do
               local __sorted = {}
               local __map = name2FieldInfo
               for __key in pairs( __map ) do
                  table.insert( __sorted, __key )
               end
               table.sort( __sorted )
               for __index, fieldName in ipairs( __sorted ) do
                  local info = __map[ fieldName ]
                  do
                     if fieldName ~= "__attrib" then
                        if self.targetLuaVer:isSupport( string.format( "%s.%s", name, fieldName) ) then
                           if _lune.nilacc( info['type'], nil, 'item', 1) == "member" then
                              self.scope:addMember( fieldName, getTypeInfo( _lune.unwrap( _lune.nilacc( info['typeInfo'], nil, 'item', 1)) ), Ast.AccessMode.Pub, true, true )
                           else
                            
                              local argTypeList = {}
                              for __index, argType in pairs( _lune.unwrap( info["arg"]) ) do
                                 table.insert( argTypeList, getTypeInfo( argType ) )
                              end
                              
                              local retTypeList = {}
                              for __index, retType in pairs( _lune.unwrap( info["ret"]) ) do
                                 local retTypeInfo = getTypeInfo( retType )
                                 table.insert( retTypeList, retTypeInfo )
                              end
                              
                              local funcType = _lune.nilacc( info['type'], nil, 'item', 1)
                              local methodFlag = funcType == "method" or funcType == "mut"
                              local mutable = funcType == "mut"
                              self:pushScope( false )
                              local typeInfo = Ast.NormalTypeInfo.createFunc( false, true, self.scope, methodFlag and Ast.TypeInfoKind.Method or Ast.TypeInfoKind.Func, parentInfo, false, true, not methodFlag, Ast.AccessMode.Pub, fieldName, argTypeList, retTypeList, mutable )
                              self:popScope(  )
                              Ast.builtInTypeIdSet[typeInfo:get_typeId(  )] = typeInfo
                              if typeInfo:get_nilableTypeInfo() ~= Ast.headTypeInfo then
                                 Ast.builtInTypeIdSet[typeInfo:get_nilableTypeInfo():get_typeId()] = typeInfo:get_nilableTypeInfo()
                              end
                              
                              self.scope:add( methodFlag and Ast.SymbolKind.Mtd or Ast.SymbolKind.Fun, not methodFlag, not methodFlag, fieldName, typeInfo, Ast.AccessMode.Pub, not methodFlag, mutable, true )
                              setupBuiltinTypeInfo( name, fieldName, typeInfo )
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
   
   self.scope = self.topScope
end

function TransUnit:error( mess )

   local pos = Parser.Position.new(0, 0)
   local txt = ""
   if self.currentToken ~= Parser.getEofToken(  ) then
      pos = self.currentToken.pos
      txt = self.currentToken.txt
   else
    
      if #self.usedTokenList > 0 then
         local token = self.usedTokenList[#self.usedTokenList]
         pos = token.pos
         txt = token.txt
      end
      
   end
   
   self:addErrMess( pos, mess )
   for __index, mess in pairs( self.errMessList ) do
      Util.errorLog( mess )
   end
   
   for __index, mess in pairs( self.warnMessList ) do
      Util.errorLog( mess )
   end
   
   Util.err( "has error" )
end

function TransUnit:createNoneNode( pos )

   return Ast.NoneNode.create( self.nodeManager, pos, {Ast.builtinTypeNone} )
end

function TransUnit:pushbackToken( token )

   if token ~= Parser.getEofToken(  ) then
      table.insert( self.pushbackList, token )
   end
   
   self.currentToken = self.usedTokenList[#self.usedTokenList]
end

local function expandVal( tokenList, val, pos )

   do
      local _exp = val
      if _exp ~= nil then
         do
            local _switchExp = type( _exp )
            if _switchExp == "boolean" then
               local token = string.format( "%s", _exp)
               local kind = Parser.TokenKind.Kywd
               table.insert( tokenList, Parser.Token.new(kind, token, pos) )
            elseif _switchExp == "number" then
               local num = string.format( "%g", _exp)
               local kind = Parser.TokenKind.Int
               if string.find( num, ".", 1, true ) then
                  kind = Parser.TokenKind.Real
               end
               
               table.insert( tokenList, Parser.Token.new(kind, num, pos) )
            elseif _switchExp == "string" then
               table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Str, string.format( '[[%s]]', _exp), pos) )
            elseif _switchExp == "table" then
               table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Dlmt, "{", pos) )
               for key, item in pairs( _exp ) do
                  expandVal( tokenList, item, pos )
                  table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Dlmt, ",", pos) )
               end
               
               table.insert( tokenList, Parser.Token.new(Parser.TokenKind.Dlmt, "}", pos) )
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
      if self.currentToken ~= Parser.getEofToken(  ) then
         table.insert( self.usedTokenList, self.currentToken )
      end
      
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
         if _exp ~= nil then
            if _exp.kind ~= Parser.TokenKind.Cmnt then
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
      if _exp ~= nil then
         if self.macroMode == Ast.MacroMode.Expand then
            local tokenTxt = _exp.txt
            if tokenTxt == ',,' or tokenTxt == ',,,,' then
               local nextToken = self:getTokenNoErr(  )
               local macroVal = self.symbol2ValueMapForMacro[nextToken.txt]
               if  nil == macroVal then
                  local _macroVal = macroVal
               
                  self:error( string.format( "unknown macro val %s", nextToken.txt) )
               end
               
               if tokenTxt == ',,' then
                  if macroVal.typeInfo:equals( Ast.builtinTypeSymbol ) then
                     local txtList = (_lune.unwrap( macroVal.val) )
                     for index = #txtList, 1, -1 do
                        nextToken = Parser.Token.new(nextToken.kind, txtList[index], nextToken.pos)
                        self:pushbackToken( nextToken )
                     end
                     
                  elseif macroVal.typeInfo:equals( Ast.builtinTypeStat ) then
                     self:pushbackStr( string.format( "macroVal %s", nextToken.txt), (_lune.unwrap( macroVal.val) ) )
                  elseif macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKind.Array or macroVal.typeInfo:get_kind(  ) == Ast.TypeInfoKind.List then
                     local strList = (_lune.unwrap( macroVal.val) )
                     for index = #strList, 1, -1 do
                        self:pushbackStr( string.format( "macroVal %s[%d]", nextToken.txt, index), strList[index] )
                     end
                     
                  else
                   
                     local tokenList = {}
                     expandVal( tokenList, macroVal.val, nextToken.pos )
                     self:newPushback( tokenList )
                  end
                  
               elseif tokenTxt == ',,,,' then
                  if macroVal.typeInfo:equals( Ast.builtinTypeSymbol ) then
                     local txtList = (_lune.unwrap( macroVal.val) )
                     local newToken = ""
                     for __index, txt in pairs( txtList ) do
                        newToken = string.format( "%s%s", newToken, txt)
                     end
                     
                     nextToken = Parser.Token.new(Parser.TokenKind.Str, string.format( "'%s'", newToken), nextToken.pos)
                     self:pushbackToken( nextToken )
                  elseif macroVal.typeInfo:equals( Ast.builtinTypeStat ) then
                     nextToken = Parser.Token.new(Parser.TokenKind.Str, string.format( "'%s'", _lune.unwrap( macroVal.val)), nextToken.pos)
                     self:pushbackToken( nextToken )
                  else
                   
                     self:error( string.format( "not support this symbol -- %s%s", tokenTxt, nextToken.txt) )
                  end
                  
               end
               
               nextToken = self:getTokenNoErr(  )
               token = nextToken
            end
            
         end
         
      end
   end
   
   table.insert( self.usedTokenList, self.currentToken )
   self.currentToken = _lune.unwrapDefault( token, Parser.getEofToken(  ))
   return self.currentToken
end

function TransUnit:getToken( allowEof )

   local token = self:getTokenNoErr(  )
   if token == Parser.getEofToken(  ) then
      if allowEof then
         return Parser.getEofToken(  )
      end
      
      self:error( "EOF" )
   end
   
   
   self.currentToken = token
   return token
end

function TransUnit:pushback(  )

   if self.currentToken ~= Parser.getEofToken(  ) then
      table.insert( self.pushbackList, self.currentToken )
   end
   
   self.currentToken = self.usedTokenList[#self.usedTokenList]
   table.remove( self.usedTokenList )
end

function TransUnit:pushbackStr( name, statement )

   local memStream = Parser.TxtStream.new(statement)
   local parser = Parser.StreamParser.new(memStream, name, false)
   local list = {}
   while true do
      do
         local _exp = parser:getToken(  )
         if _exp ~= nil then
            table.insert( list, _exp )
         else
            break
         end
      end
      
   end
   
   for index = #list, 1, -1 do
      self:pushbackToken( list[index] )
   end
   
end

function TransUnit:checkSymbol( token )

   if token.kind ~= Parser.TokenKind.Symb and token.kind ~= Parser.TokenKind.Kywd and token.kind ~= Parser.TokenKind.Type then
      self:addErrMess( token.pos, string.format( "illegal symbol -- '%s'", token.txt) )
   end
   
   return token
end

function TransUnit:getSymbolToken(  )

   return self:checkSymbol( self:getToken(  ) )
end

function TransUnit:checkToken( token, txt )

   if token.txt ~= txt then
      self:error( string.format( "not found -- %s", txt) )
   end
   
   return token
end

function TransUnit:checkNextToken( txt )

   return self:checkToken( self:getToken(  ), txt )
end

function TransUnit:getContinueToken(  )

   local prevToken = self.currentToken
   if  nil == prevToken then
      local _prevToken = prevToken
   
      return self:getToken(  ), false
   end
   
   local token = self:getToken(  )
   if prevToken.pos.lineNo ~= token.pos.lineNo or prevToken.pos.column + #prevToken.txt ~= token.pos.column then
      return token, false
   end
   
   return token, true
end

function TransUnit:analyzeStatementList( stmtList, termTxt )

   local breakKind = Ast.BreakKind.None
   if #stmtList > 0 then
      breakKind = stmtList[#stmtList]:getBreakKind( Ast.CheckBreakMode.Normal )
   end
   
   local lastStatement = nil
   while true do
      local statement = self:analyzeStatement( termTxt )
      do
         local _exp = statement
         if _exp ~= nil then
            if breakKind ~= Ast.BreakKind.None then
               self:addErrMess( _exp:get_pos(), string.format( "This statement is not reached -- %s", Ast.BreakKind:_getTxt( breakKind)
               ) )
            end
            
            table.insert( stmtList, _exp )
            lastStatement = statement
            breakKind = _exp:getBreakKind( Ast.CheckBreakMode.Normal )
         else
            break
         end
      end
      
   end
   
   return lastStatement
end

function TransUnit:analyzeStatementListSubfile( stmtList )

   local statement = self:analyzeStatement(  )
   do
      local _exp = statement
      if _exp ~= nil then
         if _exp:get_kind() ~= Ast.NodeKind.get_Subfile() then
            self:error( "subfile must have 'subfile' declaration at top." )
         end
         
      else
         self:error( "subfile must have 'subfile' declaration at top." )
      end
   end
   
   return self:analyzeStatementList( stmtList )
end

function TransUnit:analyzeLuneControl( firstToken )

   local nextToken = self:getToken(  )
   do
      local _switchExp = (nextToken.txt )
      if _switchExp == "disable_mut_control" then
         self.validMutControl = false
      else 
         
            self:addErrMess( nextToken.pos, string.format( "unknown option -- %s", nextToken.txt) )
      end
   end
   
   self:checkNextToken( ";" )
end

function TransUnit:analyzeBlock( blockKind, scope )

   local token = self:checkNextToken( "{" )
   if not scope then
      self:pushScope( false )
   end
   
   local loopFlag = false
   do
      local _switchExp = blockKind
      if _switchExp == Ast.BlockKind.For or _switchExp == Ast.BlockKind.Apply or _switchExp == Ast.BlockKind.While or _switchExp == Ast.BlockKind.Repeat or _switchExp == Ast.BlockKind.Foreach then
         loopFlag = true
         table.insert( self.loopScopeQueue, self.scope )
      end
   end
   
   local stmtList = {}
   self:analyzeStatementList( stmtList, "}" )
   self:checkNextToken( "}" )
   if loopFlag then
      table.remove( self.loopScopeQueue )
   end
   
   if not scope then
      self:popScope(  )
   end
   
   local node = Ast.BlockNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, blockKind, stmtList )
   return node
end

local DependModuleInfo = {}
function DependModuleInfo:getTypeInfo( metaTypeId )

   return _lune.unwrap( self.metaTypeId2TypeInfoMap[metaTypeId])
end
function DependModuleInfo.setmeta( obj )
  setmetatable( obj, { __index = DependModuleInfo  } )
end
function DependModuleInfo.new( id, metaTypeId2TypeInfoMap )
   local obj = {}
   DependModuleInfo.setmeta( obj )
   if obj.__init then
      obj:__init( id, metaTypeId2TypeInfoMap )
   end        
   return obj 
end         
function DependModuleInfo:__init( id, metaTypeId2TypeInfoMap ) 

   self.id = id
   self.metaTypeId2TypeInfoMap = metaTypeId2TypeInfoMap
end

function TransUnit:processImport( modulePath, moduleInfoMap )

   if not self.importModuleInfo:add( modulePath ) then
      self:error( string.format( "recursive import: %s -> %s", self.importModuleInfo:getFull(  ), modulePath) )
   end
   
   do
      local moduleInfo = self.importModuleName2ModuleInfo[modulePath]
      if moduleInfo ~= nil then
         local metaInfo = frontInterface.loadMeta( self.importModuleInfo, modulePath )
         local typeId2TypeInfo = moduleInfo:get_importId2localTypeInfoMap()
         local moduleTypeInfo = _lune.unwrap( typeId2TypeInfo[metaInfo._moduleTypeId])
         moduleInfoMap[moduleTypeInfo] = moduleInfo
         self.importModuleInfo:remove(  )
         return metaInfo, typeId2TypeInfo
      end
   end
   
   local nameList = {}
   for txt in string.gmatch( modulePath, '[^%.]+' ) do
      table.insert( nameList, txt )
   end
   
   local metaInfo = frontInterface.loadMeta( self.importModuleInfo, modulePath )
   local dependLibId2DependInfo = {}
   do
      local __sorted = {}
      local __map = metaInfo._dependModuleMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, dependName in ipairs( __sorted ) do
         local dependInfo = __map[ dependName ]
         do
            if dependInfo['use'] then
               local workModuleInfo, metaTypeId2TypeInfoMap = self:processImport( dependName, {} )
               local id = math.floor((_lune.unwrap( dependInfo['id']) ))
               dependLibId2DependInfo[id] = DependModuleInfo.new(id, metaTypeId2TypeInfoMap)
            end
            
         end
      end
   end
   
   local typeId2TypeInfo = {}
   typeId2TypeInfo[Ast.rootTypeId] = Ast.headTypeInfo
   local typeId2Scope = {}
   typeId2Scope[Ast.rootTypeId] = self.scope
   for typeId, dependIdInfo in pairs( metaInfo._dependIdMap ) do
      local dependInfo = _lune.unwrap( dependLibId2DependInfo[dependIdInfo[1]])
      local typeInfo = dependInfo:getTypeInfo( dependIdInfo[2] )
      typeId2TypeInfo[typeId] = typeInfo
      do
         local _exp = typeInfo:get_scope()
         if _exp ~= nil then
            typeId2Scope[typeId] = _exp
         end
      end
      
   end
   
   local moduleTypeInfo = Ast.headTypeInfo
   for index, moduleName in pairs( nameList ) do
      local mutable = false
      if index == #nameList then
         mutable = metaInfo._moduleMutable
      end
      
      moduleTypeInfo = self:pushModule( true, moduleName, mutable )
   end
   
   for __index, moduleName in pairs( nameList ) do
      self:popModule(  )
   end
   
   for __index, symbolInfo in pairs( Ast.sym2builtInTypeMap ) do
      typeId2TypeInfo[symbolInfo:get_typeInfo():get_typeId(  )] = symbolInfo:get_typeInfo()
   end
   
   for __index, builtinTypeInfo in pairs( Ast.builtInTypeIdSet ) do
      typeId2TypeInfo[builtinTypeInfo:get_typeId()] = builtinTypeInfo
   end
   
   local newId2OldIdMap = {}
   local _typeInfoList = {}
   local _typeInfoNormalList = {}
   for __index, atomInfo in pairs( metaInfo._typeInfoList ) do
      do
         local skind = atomInfo['skind']
         if skind ~= nil then
            local actInfo = nil
            local kind = _lune.unwrap( Ast.SerializeKind._from( math.floor(skind) ))
            do
               local _switchExp = kind
               if _switchExp == Ast.SerializeKind.Enum then
                  actInfo = _TypeInfoEnum._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Alge then
                  actInfo = _TypeInfoAlge._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Module then
                  actInfo = _TypeInfoModule._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Normal then
                  do
                     local _exp = _TypeInfoNormal._fromMap( atomInfo )
                     if _exp ~= nil then
                        table.insert( _typeInfoNormalList, _exp )
                        actInfo = _exp
                     end
                  end
                  
               elseif _switchExp == Ast.SerializeKind.Nilable then
                  actInfo = _TypeInfoNilable._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Modifier then
                  actInfo = _TypeInfoModifier._fromMap( atomInfo )
               else 
                  
                     Util.err( string.format( "unknown skind -- %d", math.floor(skind)) )
               end
            end
            
            do
               local _exp = actInfo
               if _exp ~= nil then
                  table.insert( _typeInfoList, _exp )
               else
                  for key, val in pairs( atomInfo ) do
                     Util.errorLog( string.format( "table: %s:%s", key, val) )
                  end
                  
                  Util.err( string.format( "_TypeInfo%s._fromMap error", Ast.SerializeKind:_getTxt( kind)
                  ) )
               end
            end
            
         end
      end
      
   end
   
   local importParam = ImportParam.new(self, typeId2Scope, typeId2TypeInfo, metaInfo, self.scope)
   for __index, atomInfo in pairs( _typeInfoList ) do
      local newTypeInfo, errMess = atomInfo:createTypeInfo( importParam )
      do
         local _exp = errMess
         if _exp ~= nil then
            Util.err( string.format( "%s: %s", modulePath, _exp) )
         end
      end
      
      if newTypeInfo ~= nil then
         if newTypeInfo:get_accessMode() == Ast.AccessMode.Global then
            do
               local _switchExp = newTypeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class then
                  self.globalScope:addClass( newTypeInfo:get_rawTxt(), newTypeInfo )
               elseif _switchExp == Ast.TypeInfoKind.Func then
                  self.globalScope:addFunc( newTypeInfo, Ast.AccessMode.Global, newTypeInfo:get_staticFlag(), newTypeInfo:get_mutable() )
               elseif _switchExp == Ast.TypeInfoKind.Enum then
                  self.globalScope:addEnum( Ast.AccessMode.Global, newTypeInfo:get_rawTxt(), newTypeInfo )
               elseif _switchExp == Ast.TypeInfoKind.Nilable then
                  
               else 
                  
                     Util.err( string.format( "%s: not support kind -- %s", __func__, Ast.TypeInfoKind:_getTxt( newTypeInfo:get_kind())
                     ) )
               end
            end
            
         end
         
      end
      
   end
   
   for __index, atomInfo in pairs( _typeInfoNormalList ) do
      if #atomInfo.children > 0 then
         local scope = _lune.unwrap( typeId2Scope[atomInfo.typeId])
         for __index, childId in pairs( atomInfo.children ) do
            local typeInfo = typeId2TypeInfo[childId]
            if  nil == typeInfo then
               local _typeInfo = typeInfo
            
               Util.err( string.format( "not found childId -- %s, %d, %s(%d)", modulePath, childId, atomInfo.txt, atomInfo.typeId) )
            end
            
            local symbolKind = Ast.SymbolKind.Typ
            local addFlag = true
            do
               local _switchExp = typeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.Func then
                  symbolKind = Ast.SymbolKind.Fun
               elseif _switchExp == Ast.TypeInfoKind.Method then
                  symbolKind = Ast.SymbolKind.Mtd
               elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Module then
                  symbolKind = Ast.SymbolKind.Typ
               elseif _switchExp == Ast.TypeInfoKind.Enum then
                  addFlag = false
               end
            end
            
            if addFlag then
               scope:add( symbolKind, false, typeInfo:get_kind() == Ast.TypeInfoKind.Func, typeInfo:getTxt(  ), typeInfo, typeInfo:get_accessMode(), typeInfo:get_staticFlag(), typeInfo:get_mutable(), true )
            end
            
         end
         
      end
      
   end
   
   for typeId, typeInfo in pairs( typeId2TypeInfo ) do
      newId2OldIdMap[typeInfo] = typeId
   end
   
   local function registMember( classTypeId )
   
      if metaInfo._dependIdMap[classTypeId] then
         return 
      end
      
      local classTypeInfo = _lune.unwrap( typeId2TypeInfo[classTypeId])
      do
         local _switchExp = (classTypeInfo:get_kind() )
         if _switchExp == Ast.TypeInfoKind.Class then
            self:pushClass( true, classTypeInfo:get_abstractFlag(), nil, nil, true, classTypeInfo:getTxt(  ), Ast.AccessMode.Pub )
            do
               local _exp = metaInfo._typeId2ClassInfoMap[classTypeId]
               if _exp ~= nil then
                  local classInfo = _exp
                  for fieldName, fieldInfo in pairs( classInfo ) do
                     local typeId = fieldInfo['typeId']
                     local fieldTypeInfo = _lune.unwrap( typeId2TypeInfo[typeId])
                     self.scope:addMember( fieldName, fieldTypeInfo, _lune.unwrap( Ast.AccessMode._from( math.floor((_lune.unwrap( fieldInfo['accessMode']) )) )), _lune.unwrapDefault( fieldInfo['staticFlag'], false), _lune.unwrapDefault( fieldInfo['mutable'], false) )
                  end
                  
               else
                  self:error( string.format( "not found class -- %s: %d, %s", modulePath, classTypeId, classTypeInfo:getTxt(  )) )
               end
            end
            
         elseif _switchExp == Ast.TypeInfoKind.Module then
            self:pushModule( true, classTypeInfo:getTxt(  ), classTypeInfo:get_mutable() )
         end
      end
      
      for __index, child in pairs( classTypeInfo:get_children(  ) ) do
         if child:get_kind(  ) == Ast.TypeInfoKind.Class or child:get_kind(  ) == Ast.TypeInfoKind.Module or child:get_kind(  ) == Ast.TypeInfoKind.IF then
            local oldId = newId2OldIdMap[child]
            if oldId then
               registMember( _lune.unwrap( oldId) )
            end
            
         end
         
      end
      
      if classTypeInfo:get_kind() == Ast.TypeInfoKind.Class then
         self:popClass(  )
      elseif classTypeInfo:get_kind() == Ast.TypeInfoKind.Module then
         self:popModule(  )
      end
      
   end
   
   for __index, atomInfo in pairs( _typeInfoList ) do
      if atomInfo.parentId == Ast.rootTypeId and (atomInfo.skind == Ast.SerializeKind.Normal or atomInfo.skind == Ast.SerializeKind.Module ) then
         registMember( atomInfo.typeId )
      end
      
   end
   
   for index, moduleName in pairs( nameList ) do
      local mutable = false
      if index == #nameList then
         mutable = metaInfo._moduleMutable
      end
      
      self:pushModule( true, moduleName, mutable )
   end
   
   for varName, varInfo in pairs( metaInfo._varName2InfoMap ) do
      self.scope:addStaticVar( false, true, varName, _lune.unwrap( typeId2TypeInfo[varInfo['typeId']]), _lune.unwrapDefault( varInfo['mutable'], false) )
   end
   
   for __index, moduleName in pairs( nameList ) do
      self:popModule(  )
   end
   
   local moduleInfo = Ast.ModuleInfo.new(modulePath, nameList[#nameList], newId2OldIdMap)
   self.importModule2ModuleInfo[moduleTypeInfo] = moduleInfo
   self.importModuleName2ModuleInfo[modulePath] = moduleInfo
   moduleInfoMap[moduleTypeInfo] = moduleInfo
   self.importModuleInfo:remove(  )
   return metaInfo, typeId2TypeInfo
end

function TransUnit:analyzeImport( token )

   if self.moduleScope ~= self.scope then
      self:error( "'import' must call at top scope." )
   end
   
   self.scope = self.topScope
   local moduleToken = self:getToken(  )
   local modulePath = moduleToken.txt
   local nextToken = moduleToken
   while true do
      nextToken = self:getToken(  )
      if nextToken.txt == "." then
         nextToken = self:getToken(  )
         moduleToken = nextToken
         modulePath = string.format( "%s.%s", modulePath, moduleToken.txt)
      else
       
         break
      end
      
   end
   
   local metaInfo, typeId2TypeInfo = self:processImport( modulePath, self.importModule2ModuleInfoCurrent )
   self.scope = self.moduleScope
   local assignName = moduleToken
   if nextToken.txt == "as" then
      assignName = self:getSymbolToken(  )
      nextToken = self:getToken(  )
   end
   
   local moduleTypeInfo = _lune.unwrap( typeId2TypeInfo[metaInfo._moduleTypeId])
   local moduleSymbolKind = _lune.unwrap( Ast.SymbolKind._from( metaInfo._moduleSymbolKind ))
   self.scope:add( moduleSymbolKind, false, false, assignName.txt, moduleTypeInfo, Ast.AccessMode.Local, true, metaInfo._moduleMutable, true )
   self:checkToken( nextToken, ";" )
   if self.moduleScope ~= self.scope then
      self:error( "illegal top scope." )
   end
   
   return Ast.ImportNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, modulePath, assignName.txt, moduleTypeInfo )
end

function TransUnit:analyzeSubfile( token )

   if self.scope ~= self.moduleScope then
      self:error( "'module' must be top scope." )
   end
   
   local mode = self:getToken(  )
   local moduleName = ""
   while true do
      local nextToken = self:getToken(  )
      if nextToken.txt == ";" then
         break
      end
      
      if moduleName == "" then
         moduleName = nextToken.txt
      else
       
         moduleName = string.format( "%s%s", moduleName, nextToken.txt)
      end
      
   end
   
   local usePath = nil
   if moduleName == "" then
      self:addErrMess( token.pos, "illegal subfile" )
   else
    
      if mode.txt == "use" then
         usePath = moduleName
         if frontInterface.searchModule( moduleName ) then
            table.insert( self.subfileList, moduleName )
         else
          
            self:addErrMess( token.pos, string.format( "not found subfile -- %s", moduleName) )
         end
         
      elseif mode.txt == "owner" then
         if self.moduleName ~= moduleName then
            self:addErrMess( token.pos, string.format( "illegal owner module -- %s, %s", moduleName, self.moduleName) )
         end
         
      else
       
         self:addErrMess( mode.pos, string.format( "illegal module mode -- %s", mode.txt) )
      end
      
   end
   
   return Ast.SubfileNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, usePath )
end

function TransUnit:analyzeIf( token )

   local nextToken, continueFlag = self:getContinueToken(  )
   if continueFlag and nextToken.txt == "!" then
      return self:analyzeIfUnwrap( token )
   end
   
   self:pushback(  )
   local list = {}
   local ifExp = self:analyzeExp( false )
   table.insert( list, Ast.IfStmtInfo.new(Ast.IfKind.If, ifExp, self:analyzeBlock( Ast.BlockKind.If )) )
   local function checkCond( condExp )
   
      do
         local _switchExp = condExp:get_expType():get_kind()
         if _switchExp == Ast.TypeInfoKind.Nilable or _switchExp == Ast.TypeInfoKind.Stem then
            
         elseif _switchExp == Ast.TypeInfoKind.Prim then
            if not condExp:get_expType():equals( Ast.builtinTypeBool ) then
               self:addErrMess( condExp:get_pos(), string.format( "This exp never be false -- %s", condExp:get_expType():getTxt(  )) )
            end
            
         else 
            
               self:addErrMess( condExp:get_pos(), string.format( "This exp never be false -- %s", condExp:get_expType():getTxt(  )) )
         end
      end
      
   end
   
   checkCond( ifExp )
   nextToken = self:getToken( true )
   if nextToken.txt == "elseif" then
      while nextToken.txt == "elseif" do
         local condExp = self:analyzeExp( false )
         table.insert( list, Ast.IfStmtInfo.new(Ast.IfKind.ElseIf, condExp, self:analyzeBlock( Ast.BlockKind.Elseif )) )
         checkCond( condExp )
         nextToken = self:getToken( true )
      end
      
   end
   
   if nextToken.txt == "else" then
      table.insert( list, Ast.IfStmtInfo.new(Ast.IfKind.Else, self:createNoneNode( nextToken.pos ), self:analyzeBlock( Ast.BlockKind.Else )) )
   else
    
      self:pushback(  )
   end
   
   return Ast.IfNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, list )
end

function TransUnit:analyzeSwitch( firstToken )

   local exp = self:analyzeExp( false )
   self:checkNextToken( "{" )
   local caseList = {}
   local nextToken = self:getToken(  )
   while (nextToken.txt == "case" ) do
      self:checkToken( nextToken, "case" )
      local condexpList = self:analyzeExpList( false, nil, {exp:get_expType()}, true )
      local condBock = self:analyzeBlock( Ast.BlockKind.Switch )
      table.insert( caseList, Ast.CaseInfo.new(condexpList, condBock) )
      nextToken = self:getToken(  )
   end
   
   local defaultBlock = nil
   if nextToken.txt == "default" then
      defaultBlock = self:analyzeBlock( Ast.BlockKind.Default )
   else
    
      self:pushback(  )
   end
   
   self:checkNextToken( "}" )
   return Ast.SwitchNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, exp, caseList, defaultBlock )
end

function TransUnit:analyzeMatch( firstToken )

   local exp = self:analyzeExp( false )
   if exp:get_expType():get_kind() ~= Ast.TypeInfoKind.Alge then
      self:error( "match must have alge value" )
   end
   
   local algeTypeInfo = exp:get_expType():get_srcTypeInfo()
   self:checkNextToken( "{" )
   local caseList = {}
   local nextToken = self:getToken(  )
   while (nextToken.txt == "case" ) do
      self:checkNextToken( "." )
      local valNameToken = self:getToken(  )
      local valInfo = algeTypeInfo:getValInfo( valNameToken.txt )
      if  nil == valInfo then
         local _valInfo = valInfo
      
         self:error( string.format( "not found val -- %s", valNameToken.txt) )
      end
      
      local valParamNameList = {}
      nextToken = self:getToken(  )
      local blockScope = self:pushScope( false )
      if nextToken.txt == "(" then
         for __index, paramType in pairs( valInfo:get_typeList() ) do
            local paramName = self:getSymbolToken(  )
            if self.scope:getTypeInfo( paramName.txt, self.scope, true ) then
               self:addErrMess( paramName.pos, string.format( "shadowing variable -- %s", paramName.txt) )
            end
            
            local workType = paramType
            if paramType:get_mutable() and not exp:get_expType():get_mutable() then
               workType = self:createModifier( workType, false )
            end
            
            blockScope:addLocalVar( false, false, paramName.txt, workType, false )
            table.insert( valParamNameList, paramName.txt )
            nextToken = self:getToken(  )
            if nextToken.txt ~= "," then
               break
            end
            
         end
         
         self:checkToken( nextToken, ")" )
      else
       
         self:pushback(  )
      end
      
      if #valParamNameList ~= #valInfo:get_typeList() then
         self:addErrMess( valNameToken.pos, string.format( "unmatch param -- %d != %d", #valParamNameList, #valInfo:get_typeList()) )
      end
      
      local block = self:analyzeBlock( Ast.BlockKind.Match, blockScope )
      self:popScope(  )
      local matchCase = Ast.MatchCase.new(valInfo, valParamNameList, block)
      table.insert( caseList, matchCase )
      nextToken = self:getToken(  )
   end
   
   local defaultBlock = nil
   if nextToken.txt == "default" then
      defaultBlock = self:analyzeBlock( Ast.BlockKind.Block )
      nextToken = self:getToken(  )
   end
   
   self:checkToken( nextToken, "}" )
   return Ast.MatchNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, exp, algeTypeInfo, caseList, defaultBlock )
end

function TransUnit:analyzeWhile( token )

   return Ast.WhileNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, self:analyzeExp( false ), self:analyzeBlock( Ast.BlockKind.While ) )
end

function TransUnit:analyzeRepeat( token )

   local scope = self:pushScope( false )
   local node = Ast.RepeatNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, self:analyzeBlock( Ast.BlockKind.Repeat, scope ), self:analyzeExp( false ) )
   self:popScope(  )
   self:checkNextToken( ";" )
   return node
end

function TransUnit:analyzeFor( firstToken )

   local scope = self:pushScope( false )
   local val = self:getToken(  )
   if val.kind ~= Parser.TokenKind.Symb then
      self:error( "not symbol" )
   end
   
   self:checkNextToken( "=" )
   local exp1 = self:analyzeExp( false )
   if not exp1:get_expType():equals( Ast.builtinTypeInt ) then
      self:addErrMess( exp1:get_pos(), string.format( "exp1 is not int -- %s", exp1:get_expType():getTxt(  )) )
   end
   
   self:addLocalVar( exp1:get_pos(), false, true, val.txt, exp1:get_expType(), false )
   self:checkNextToken( "," )
   local exp2 = self:analyzeExp( false )
   if not exp2:get_expType():equals( Ast.builtinTypeInt ) then
      self:addErrMess( exp2:get_pos(), string.format( "exp2 is not int -- %s", exp2:get_expType():getTxt(  )) )
   end
   
   local token = self:getToken(  )
   local exp3 = nil
   if token.txt == "," then
      exp3 = self:analyzeExp( false )
      do
         local _exp = exp3
         if _exp ~= nil then
            if not _exp:get_expType():equals( Ast.builtinTypeInt ) then
               self:addErrMess( _exp:get_pos(), string.format( "exp is not int -- %s", _exp:get_expType():getTxt(  )) )
            end
            
         end
      end
      
   else
    
      self:pushback(  )
   end
   
   local node = Ast.ForNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, self:analyzeBlock( Ast.BlockKind.For, scope ), val, exp1, exp2, exp3 )
   self:popScope(  )
   return node
end

function TransUnit:analyzeApply( token )

   local scope = self:pushScope( false )
   local varList = {}
   local nextToken = Parser.getEofToken(  )
   repeat 
      local var = self:getSymbolToken(  )
      if var.kind ~= Parser.TokenKind.Symb then
         self:error( "illegal symbol" )
      end
      
      table.insert( varList, var )
      nextToken = self:getToken(  )
   until nextToken.txt ~= ","
   self:checkToken( nextToken, "of" )
   local exp = self:analyzeExp( false )
   local expTypeList = exp:get_expTypeList()
   if #expTypeList < 3 then
      self:addErrMess( exp:get_pos(), string.format( "apply must have 3 values -- %s", #expTypeList) )
   end
   
   local itemTypeList = {}
   local defaultItemType = Ast.builtinTypeStem_
   if exp:get_kind() == Ast.nodeKind['ExpCall'] then
      local callNode = exp
      local callFuncType = callNode:get_func():get_expType()
      if callFuncType:equals( typeInfoStrGMatch ) or callFuncType:equals( typeInfoStringGMatch ) then
         table.insert( itemTypeList, Ast.builtinTypeString )
         defaultItemType = Ast.builtinTypeString:get_nilableTypeInfo()
      else
       
         if #callFuncType:get_retTypeInfoList() == 0 then
            self:addErrMess( exp:get_pos(), "apply value must return iterator function." )
         end
         
         local iteFunc = callFuncType:get_retTypeInfoList()[1]
         for index, itemType in pairs( iteFunc:get_retTypeInfoList() ) do
            local workType = itemType
            if index == 1 then
               if itemType:get_nilable() then
                  workType = workType:get_orgTypeInfo()
               end
               
            end
            
            table.insert( itemTypeList, workType )
         end
         
      end
      
   end
   
   for index, var in pairs( varList ) do
      local itemType = defaultItemType
      if index <= #itemTypeList then
         itemType = itemTypeList[index]
      end
      
      self:addLocalVar( var.pos, false, true, var.txt, itemType, false )
   end
   
   local block = self:analyzeBlock( Ast.BlockKind.Apply, scope )
   self:popScope(  )
   return Ast.ApplyNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, varList, exp, block )
end

function TransUnit:analyzeForeach( token, sortFlag )

   local scope = self:pushScope( false )
   local valSymbol = Parser.getEofToken(  )
   local keySymbol = nil
   local nextToken = Parser.getEofToken(  )
   for index = 1, 2 do
      local symbol = self:getToken(  )
      if symbol.kind ~= Parser.TokenKind.Symb then
         self:error( "illegal symbol" )
      end
      
      if index == 1 then
         valSymbol = symbol
      else
       
         keySymbol = symbol
      end
      
      nextToken = self:getToken(  )
      if nextToken.txt ~= "," then
         break
      end
      
   end
   
   self:checkToken( nextToken, "in" )
   local exp = self:analyzeExp( false )
   local itemTypeInfoList = exp:get_expType():get_itemTypeInfoList(  )
   if exp:get_expType():get_kind(  ) == Ast.TypeInfoKind.Map then
      self:addLocalVar( exp:get_pos(), false, true, valSymbol.txt, itemTypeInfoList[2], false )
      do
         local _exp = keySymbol
         if _exp ~= nil then
            self:addLocalVar( _exp.pos, false, true, _exp.txt, itemTypeInfoList[1], false )
         end
      end
      
   elseif exp:get_expType():get_kind(  ) == Ast.TypeInfoKind.List or exp:get_expType():get_kind(  ) == Ast.TypeInfoKind.Array then
      self.scope:addLocalVar( false, true, valSymbol.txt, itemTypeInfoList[1], false )
      do
         local _exp = keySymbol
         if _exp ~= nil then
            self:addLocalVar( _exp.pos, false, false, _exp.txt, Ast.builtinTypeInt, false )
         else
            self.scope:addLocalVar( false, false, "__index", Ast.builtinTypeInt, false )
         end
      end
      
   else
    
      self:error( string.format( "unknown kind type of exp for foreach-- %s(%d:%d)", exp:get_expType():getTxt(  ), exp:get_pos().lineNo, exp:get_pos().column) )
   end
   
   local block = self:analyzeBlock( Ast.BlockKind.Foreach, scope )
   self:popScope(  )
   if sortFlag then
      return Ast.ForsortNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, valSymbol, keySymbol, exp, block, sortFlag )
   else
    
      return Ast.ForeachNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, valSymbol, keySymbol, exp, block )
   end
   
end

function TransUnit:analyzeProvide( firstToken )

   local token = self:getSymbolToken(  )
   local symbolNode = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, true )
   self:checkNextToken( ";" )
   local symbolInfoList = symbolNode:getSymbolInfo(  )
   if #symbolInfoList ~= 1 then
      self:error( "'provide' must be symbol." )
   end
   
   local symbolInfo = symbolInfoList[1]
   local node = Ast.ProvideNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, symbolInfo )
   if self.provideNode then
      self:addErrMess( firstToken.pos, "multiple provide" )
   end
   
   self.provideNode = node
   if symbolInfo:get_accessMode() ~= Ast.AccessMode.Pub then
      self:addErrMess( firstToken.pos, string.format( "provide variable must be 'pub'.  -- %s", symbolInfo:get_accessMode()) )
   end
   
   return node
end

function TransUnit:analyzeRefType( accessMode, allowDDD )

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
   
   local typeInfo = Ast.builtinTypeStem_
   self:checkSymbol( token )
   local name = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, true )
   typeInfo = name:get_expType()
   local continueToken, continueFlag = self:getContinueToken(  )
   token = continueToken
   if continueFlag and token.txt == "!" then
      typeInfo = _lune.unwrap( typeInfo:get_nilableTypeInfo(  ))
      token = self:getToken(  )
   end
   
   local arrayMode = "no"
   while true do
      if token.txt == '[' or token.txt == '[@' then
         if token.txt == '[' then
            arrayMode = "list"
            typeInfo = Ast.NormalTypeInfo.createList( accessMode, self:getCurrentClass(  ), {typeInfo} )
         else
          
            arrayMode = "array"
            typeInfo = Ast.NormalTypeInfo.createArray( accessMode, self:getCurrentClass(  ), {typeInfo} )
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
            local typeExp = self:analyzeRefType( accessMode, false )
            table.insert( genericList, typeExp:get_expType() )
            nextToken = self:getToken(  )
         until nextToken.txt ~= ","
         self:checkToken( nextToken, '>' )
         if typeInfo:get_kind() == Ast.TypeInfoKind.Map then
            typeInfo = Ast.NormalTypeInfo.createMap( accessMode, self:getCurrentClass(  ), genericList[1] or Ast.builtinTypeStem, genericList[2] or Ast.builtinTypeStem )
         elseif typeInfo:get_kind() == Ast.TypeInfoKind.List then
            typeInfo = Ast.NormalTypeInfo.createList( accessMode, self:getCurrentClass(  ), {genericList[1]} or {Ast.builtinTypeStem} )
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
      typeInfo = _lune.unwrap( typeInfo:get_nilableTypeInfo(  ))
      token = self:getToken(  )
   end
   
   if not allowDDD and typeInfo:equals( Ast.builtinTypeDDD ) then
      self:addErrMess( firstToken.pos, string.format( "invalid type. -- '%s'", typeInfo:getTxt(  )) )
   end
   
   if refFlag then
      typeInfo = self:createModifier( typeInfo, false )
   end
   
   return Ast.RefTypeNode.create( self.nodeManager, firstToken.pos, {typeInfo}, name, refFlag, mutFlag, arrayMode )
end

function TransUnit:analyzeDeclArgList( accessMode, argList )

   local nextToken = Parser.noneToken
   repeat 
      nextToken = self:getToken(  )
      if nextToken.txt == ")" then
         break
      end
      
      local mutable = false
      if nextToken.txt == "mut" then
         mutable = true
         nextToken = self:getToken(  )
      end
      
      local argName = nextToken
      if argName.txt == "..." then
         table.insert( argList, Ast.DeclArgDDDNode.create( self.nodeManager, argName.pos, {Ast.builtinTypeDDD} ) )
         self.scope:addLocalVar( false, true, argName.txt, Ast.builtinTypeDDD, false )
      else
       
         argName = self:checkSymbol( argName )
         if self.scope:getSymbolTypeInfo( argName.txt, self.scope, self.moduleScope ) then
            self:addErrMess( argName.pos, string.format( "shadowing variable -- %s", argName.txt) )
         end
         
         self:checkNextToken( ":" )
         local refType = self:analyzeRefType( accessMode, false )
         local arg = Ast.DeclArgNode.create( self.nodeManager, argName.pos, refType:get_expTypeList(), argName, refType )
         self.scope:addLocalVar( false, true, argName.txt, refType:get_expType(), mutable )
         table.insert( argList, arg )
      end
      
      nextToken = self:getToken(  )
   until nextToken.txt ~= ","
   self:checkToken( nextToken, ")" )
   return nextToken
end

local ASTInfo = {}
_moduleObj.ASTInfo = ASTInfo
function ASTInfo.setmeta( obj )
  setmetatable( obj, { __index = ASTInfo  } )
end
function ASTInfo.new( node, moduleTypeInfo, moduleSymbolKind )
   local obj = {}
   ASTInfo.setmeta( obj )
   if obj.__init then
      obj:__init( node, moduleTypeInfo, moduleSymbolKind )
   end        
   return obj 
end         
function ASTInfo:__init( node, moduleTypeInfo, moduleSymbolKind ) 

   self.node = node
   self.moduleTypeInfo = moduleTypeInfo
   self.moduleSymbolKind = moduleSymbolKind
end
function ASTInfo:get_node()       
   return self.node         
end
function ASTInfo:get_moduleTypeInfo()       
   return self.moduleTypeInfo         
end
function ASTInfo:get_moduleSymbolKind()       
   return self.moduleSymbolKind         
end

function TransUnit:createAST( parser, macroFlag, moduleName )

   self.moduleName = _lune.unwrapDefault( moduleName, "")
   self:registBuiltInScope(  )
   local processInfo = Ast.pushProcessInfo(  )
   local moduleTypeInfo = Ast.headTypeInfo
   local moduleSymbolKind = Ast.SymbolKind.Typ
   if moduleName ~= nil then
      for txt in string.gmatch( moduleName, '[^%.]+' ) do
         moduleTypeInfo = self:pushModule( false, txt, true )
      end
      
   end
   
   self.moduleScope = self.scope
   self.moduleType = moduleTypeInfo
   self.parser = parser
   local ast = nil
   local lastStatement = nil
   if macroFlag then
      ast = self:analyzeBlock( Ast.BlockKind.Macro )
   else
    
      local children = {}
      lastStatement = self:analyzeStatementList( children )
      local token = self:getTokenNoErr(  )
      if token ~= Parser.getEofToken(  ) then
         self:error( string.format( "%s:%d:%d:(%s) not eof -- %s", self.parser:getStreamName(  ), token.pos.lineNo, token.pos.column, Parser.TokenKind:_getTxt( token.kind)
         , token.txt) )
      end
      
      for __index, subModule in pairs( self.subfileList ) do
         local file = frontInterface.searchModule( subModule )
         if  nil == file then
            local _file = file
         
            self:error( string.format( "not found subfile -- %s", subModule) )
         end
         
         if self.scope ~= self.moduleScope then
            self:error( "scope does not close" )
         end
         
         local subParser = Parser.StreamParser.create( file, false, subModule )
         if  nil == subParser then
            local _subParser = subParser
         
            self:error( string.format( "open error -- %s", file) )
         end
         
         self.parser = subParser
         lastStatement = self:analyzeStatementListSubfile( children )
         token = self:getTokenNoErr(  )
         if token ~= Parser.getEofToken(  ) then
            Util.err( string.format( "unknown:%d:%d:(%s) %s", token.pos.lineNo, token.pos.column, Parser.TokenKind:_getTxt( token.kind)
            , token.txt) )
         end
         
      end
      
      local rootNode = Ast.RootNode.create( self.nodeManager, Parser.Position.new(0, 0), {Ast.builtinTypeNone}, children, processInfo, moduleTypeInfo, nil, self.helperInfo, self.nodeManager, self.importModule2ModuleInfo, self.typeId2ClassMap )
      ast = rootNode
      do
         local _exp = self.provideNode
         if _exp ~= nil then
            if lastStatement ~= _exp then
               self:addErrMess( _exp:get_pos(), "'provide' must be last." )
            end
            
            rootNode:set_provide( _exp )
            moduleSymbolKind = _exp:get_symbol():get_kind()
         end
      end
      
   end
   
   if moduleName ~= nil then
      for txt in string.gmatch( moduleName, '[^%.]+' ) do
         self:popModule(  )
      end
      
   end
   
   Ast.popProcessInfo(  )
   for protoType, pos in pairs( self.protoFuncMap ) do
      self:addErrMess( pos, string.format( "This function doesn't have body. -- %s", protoType:getTxt(  )) )
   end
   
   for __index, mess in pairs( self.warnMessList ) do
      Util.errorLog( mess )
   end
   
   if #self.errMessList > 0 then
      for __index, mess in pairs( self.errMessList ) do
         Util.errorLog( mess )
      end
      
      Util.err( "has error" )
   end
   
   if self.analyzeMode == AnalyzeMode.Diag or self.analyzeMode == AnalyzeMode.Complete then
      os.exit( 0 )
   end
   
   return ASTInfo.new(_lune.unwrap( ast), moduleTypeInfo, moduleSymbolKind)
end

function TransUnit:analyzeDeclMacro( accessMode, firstToken )

   local nameToken = self:getToken(  )
   self:checkNextToken( "(" )
   local scope = self:pushScope( false )
   local workArgList = {}
   local argList = {}
   local nextToken = self:analyzeDeclArgList( accessMode, workArgList )
   local argTypeList = {}
   for index, argNode in pairs( workArgList ) do
      if argNode:get_kind() == Ast.NodeKind.get_DeclArg() then
         table.insert( argList, argNode )
      else
       
         self:error( "macro argument can not use '...'." )
      end
      
      local argType = argNode:get_expType()
      table.insert( argTypeList, argType )
   end
   
   self:checkNextToken( "{" )
   nextToken = self:getToken(  )
   local ast = nil
   if nextToken.txt == "{" then
      local parser = Parser.WrapParser.new(self.parser, string.format( "decl macro %s", nameToken.txt))
      self.macroScope = scope
      local funcType = Ast.NormalTypeInfo.createFunc( false, true, nil, Ast.TypeInfoKind.Func, Ast.headTypeInfo, false, true, true, Ast.AccessMode.Global, "_lnsLoad", {Ast.builtinTypeString, Ast.builtinTypeString}, {Ast.builtinTypeStem}, false )
      scope:addLocalVar( false, false, "_lnsLoad", funcType, false )
      local bakParser = self.parser
      self.parser = parser
      local stmtList = {}
      self:analyzeStatementList( stmtList, "}" )
      self:checkNextToken( "}" )
      self.parser = bakParser
      self.macroScope = nil
      ast = Ast.BlockNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, Ast.BlockKind.Macro, stmtList )
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
   
   local typeInfo = Ast.NormalTypeInfo.createFunc( false, false, scope, Ast.TypeInfoKind.Macro, self:getCurrentNamespaceTypeInfo(  ), false, false, false, accessMode, nameToken.txt, argTypeList )
   self.scope:addLocalVar( false, false, nameToken.txt, typeInfo, false )
   local declMacroInfo = Ast.DeclMacroInfo.new(nameToken, argList, ast, tokenList)
   local node = Ast.DeclMacroNode.create( self.nodeManager, firstToken.pos, {typeInfo}, declMacroInfo )
   local macroObj = self.macroEval:eval( node )
   self.typeId2MacroInfo[typeInfo:get_typeId(  )] = Ast.MacroInfo.new(macroObj, declMacroInfo, self.symbol2ValueMapForMacro)
   self.symbol2ValueMapForMacro = {}
   return node
end

function TransUnit:analyzePushClass( classFlag, abstractFlag, firstToken, name, accessMode )

   local nextToken = self:getToken(  )
   local baseRef = nil
   local interfaceList = {}
   if nextToken.txt == "extend" then
      nextToken = self:getToken(  )
      if nextToken.txt ~= "(" then
         self:pushback(  )
         do
            local _sync_baseRef
            do
               local baseRef = self:analyzeRefType( accessMode, false )
               if  nil == baseRef then
                  local _baseRef = baseRef
               
               end
               
                  if baseRef:get_expType():get_kind() ~= Ast.TypeInfoKind.Class then
                     self:addErrMess( baseRef:get_pos(), string.format( "%s is not class.", baseRef:get_expType():getTxt(  )) )
                  end
                  
               _sync_baseRef = baseRef
            end
            baseRef = _sync_baseRef
         end
         
         nextToken = self:getToken(  )
      end
      
      if nextToken.txt == "(" then
         while true do
            nextToken = self:getToken(  )
            if nextToken.txt == ")" then
               break
            end
            
            self:pushback(  )
            local ifType = self:analyzeRefType( accessMode, false )
            if ifType:get_expType():get_kind() ~= Ast.TypeInfoKind.IF then
               self:error( string.format( "%s is not interface -- %d", ifType:get_expType():getTxt(  ), ifType:get_expType():get_kind()) )
            end
            
            table.insert( interfaceList, ifType:get_expType() )
            nextToken = self:getToken(  )
            if nextToken.txt ~= "," then
               if nextToken.txt == ")" then
                  break
               end
               
               self:error( "illegal token" )
            end
            
         end
         
         nextToken = self:getToken(  )
      end
      
   end
   
   local typeInfo = nil
   if baseRef ~= nil then
      local baseTypeInfo = baseRef:get_expType(  )
      typeInfo = baseTypeInfo
      local initTypeInfo = _lune.nilacc( baseTypeInfo:get_scope(), 'getTypeInfoChild', 'callmtd' , "__init" )
      if  nil == initTypeInfo then
         local _initTypeInfo = initTypeInfo
      
      else
         
            if initTypeInfo:get_accessMode() == Ast.AccessMode.Pri then
               self:addErrMess( firstToken.pos, "The access mode of '__init' is 'pri'." )
            end
            
      end
      
   end
   
   local symbol2TypeInfo = {}
   for __index, ifType in pairs( interfaceList ) do
      _lune.nilacc( ifType:get_scope(), 'filterTypeInfoField', 'callmtd' , true, self.scope, function ( symbolInfo )
      
         do
            local ifFuncType = symbol2TypeInfo[symbolInfo:get_name()]
            if ifFuncType ~= nil then
               if not ifFuncType:canEvalWith( symbolInfo:get_typeInfo(), "=" ) then
                  self:addErrMess( firstToken.pos, string.format( "mismatch method type -- %s.%s, %s.%s", symbolInfo:get_typeInfo():get_parentInfo():getTxt(  ), symbolInfo:get_name(), ifFuncType:get_parentInfo():getTxt(  ), ifFuncType:getTxt(  )) )
               end
               
            else
               symbol2TypeInfo[symbolInfo:get_name()] = symbolInfo:get_typeInfo()
            end
         end
         
         return true
      end
       )
   end
   
   local classTypeInfo = self:pushClass( classFlag, abstractFlag, typeInfo, interfaceList, false, name.txt, accessMode )
   return nextToken, classTypeInfo
end

function TransUnit:analyzeDeclProto( accessMode, firstToken )

   local nextToken = self:getToken(  )
   local abstractFlag = false
   if nextToken.txt == "abstract" then
      abstractFlag = true
      nextToken = self:getToken(  )
   end
   
   if nextToken.txt == "class" or nextToken.txt == "interface" then
      local name = self:getSymbolToken(  )
      nextToken = self:analyzePushClass( nextToken.txt ~= "interface", abstractFlag, firstToken, name, accessMode )
      self:popClass(  )
      self:checkToken( nextToken, ";" )
   else
    
      self:error( "illegal proto" )
   end
   
   return self:createNoneNode( firstToken.pos )
end

function TransUnit:analyzeDeclEnum( accessMode, firstToken )

   local name = self:getSymbolToken(  )
   self:checkNextToken( "{" )
   local valueList = {}
   local valueName2Info = {}
   local scope = self:pushScope( true )
   local enumTypeInfo = Ast.headTypeInfo
   local nextToken = self:getToken(  )
   local number = 0.0
   local prevValTypeInfo = Ast.headTypeInfo
   local valTypeInfo = Ast.headTypeInfo
   while nextToken.txt ~= "}" do
      local valName = self:checkSymbol( nextToken )
      nextToken = self:getToken(  )
      local enumVal = number
      do
         local _switchExp = (prevValTypeInfo )
         if _switchExp == Ast.builtinTypeReal then
         elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.headTypeInfo then
            enumVal = math.floor(number)
         end
      end
      
      if nextToken.txt == "=" then
         local exp = self:analyzeExp( false )
         local valList, typeInfoList = exp:getLiteral(  )
         if #valList ~= 1 or #typeInfoList ~= 1 then
            self:error( string.format( "illegal enum val -- %d %d", #valList, #typeInfoList) )
         end
         
         valTypeInfo = typeInfoList[1]
         do
            local _switchExp = (valTypeInfo )
            if _switchExp == Ast.builtinTypeString then
               enumVal = (_lune.unwrap( valList[1]) )
            elseif _switchExp == Ast.builtinTypeInt then
               local val = math.floor((_lune.unwrap( valList[1]) ))
               enumVal = val
               number = val
            elseif _switchExp == Ast.builtinTypeReal then
               number = (_lune.unwrap( valList[1]) )
               enumVal = number
            else 
               
                  self:error( string.format( "illegal enum val type -- %s", valTypeInfo:getTxt(  )) )
            end
         end
         
         nextToken = self:getToken(  )
      else
       
         do
            local _switchExp = (prevValTypeInfo )
            if _switchExp == Ast.headTypeInfo then
               valTypeInfo = Ast.builtinTypeInt
            elseif _switchExp == Ast.builtinTypeInt or _switchExp == Ast.builtinTypeReal then
               valTypeInfo = prevValTypeInfo
            else 
               
                  self:addErrMess( valName.pos, string.format( "illegal enum val type -- %s", valTypeInfo:getTxt(  )) )
            end
         end
         
      end
      
      if prevValTypeInfo ~= Ast.headTypeInfo and prevValTypeInfo ~= valTypeInfo then
         self:addErrMess( valName.pos, string.format( "multiple enum val type. %s, %s", valTypeInfo:getTxt(  ), prevValTypeInfo:getTxt(  )) )
      end
      
      prevValTypeInfo = valTypeInfo
      if enumTypeInfo == Ast.headTypeInfo then
         enumTypeInfo = Ast.NormalTypeInfo.createEnum( scope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt, valTypeInfo, valueName2Info )
      end
      
      scope:addEnumVal( valName.txt, enumTypeInfo )
      local enumValInfo = Ast.EnumValInfo.new(valName.txt, enumVal)
      table.insert( valueList, valName )
      valueName2Info[valName.txt] = enumValInfo
      if nextToken.txt == "}" then
         break
      end
      
      self:checkToken( nextToken, "," )
      nextToken = self:getToken(  )
      number = number + 1
   end
   
   if enumTypeInfo == Ast.headTypeInfo then
      enumTypeInfo = Ast.NormalTypeInfo.createEnum( scope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt, Ast.builtinTypeNone, valueName2Info )
   end
   
   self:popScope(  )
   self.scope:addEnum( accessMode, name.txt, enumTypeInfo )
   return Ast.DeclEnumNode.create( self.nodeManager, firstToken.pos, {enumTypeInfo}, accessMode, name, valueList, scope )
end

function TransUnit:analyzeDeclAlge( accessMode, firstToken )

   local name = self:getSymbolToken(  )
   self:checkNextToken( "{" )
   local scope = self.scope
   local algeScope = self:pushScope( true )
   local algeTypeInfo = Ast.NormalTypeInfo.createAlge( algeScope, self:getCurrentNamespaceTypeInfo(  ), false, accessMode, name.txt )
   scope:addAlge( accessMode, name.txt, algeTypeInfo )
   local nextToken = self:getToken(  )
   while nextToken.txt ~= "}" do
      local valName = self:checkSymbol( nextToken )
      nextToken = self:getToken(  )
      local typeInfoList = {}
      if nextToken.txt == "(" then
         while true do
            local typeNode = self:analyzeRefType( Ast.AccessMode.Pub, false )
            table.insert( typeInfoList, typeNode:get_expType() )
            nextToken = self:getToken(  )
            if nextToken.txt ~= "," then
               self:checkToken( nextToken, ")" )
               nextToken = self:getToken(  )
               break
            end
            
         end
         
      end
      
      algeScope:addAlgeVal( valName.txt, algeTypeInfo )
      local algeValInfo = Ast.AlgeValInfo.new(valName.txt, typeInfoList)
      algeTypeInfo:addValInfo( algeValInfo )
      if nextToken.txt == "}" then
         break
      end
      
      self:checkToken( nextToken, "," )
      nextToken = self:getToken(  )
   end
   
   self:popScope(  )
   return Ast.DeclAlgeNode.create( self.nodeManager, firstToken.pos, {algeTypeInfo}, accessMode, algeTypeInfo, algeScope )
end

function TransUnit:analyzeRetTypeList( pubToExtFlag, accessMode, token )

   local retTypeInfoList = {}
   if token.txt == ":" then
      while true do
         local refTypeNode = self:analyzeRefType( accessMode, true )
         local retType = refTypeNode:get_expType()
         if pubToExtFlag and not Ast.isPubToExternal( retType:get_accessMode() ) then
            self:addErrMess( refTypeNode:get_pos(), string.format( "this is not public type -- %s", retType:getTxt(  )) )
         end
         
         table.insert( retTypeInfoList, retType )
         token = self:getToken(  )
         if token.txt ~= "," then
            break
         end
         
      end
      
   end
   
   return retTypeInfoList, token
end

function TransUnit:analyzeDeclForm( accessMode, firstToken )

   local name = self:getSymbolToken(  )
   self:checkNextToken( "(" )
   local argList = {}
   local funcBodyScope = self:pushScope( false )
   local nextToken = self:analyzeDeclArgList( accessMode, argList )
   self:checkToken( nextToken, ")" )
   local retTypeList = {}
   nextToken = self:getToken(  )
   retTypeList, nextToken = self:analyzeRetTypeList( Ast.isPubToExternal( accessMode ), accessMode, nextToken )
   self:checkToken( nextToken, ";" )
   self:popScope(  )
   local argTypeInfoList = {}
   for __index, argNode in pairs( argList ) do
      table.insert( argTypeInfoList, argNode:get_expType() )
   end
   
   local formType = Ast.NormalTypeInfo.createFunc( false, false, nil, Ast.TypeInfoKind.Func, self:getCurrentNamespaceTypeInfo(  ), false, false, true, accessMode, name.txt, argTypeInfoList, retTypeList, false )
   self.scope:add( Ast.SymbolKind.Fun, false, false, name.txt, formType, accessMode, true, false, false )
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
   
   local abstractFlag = false
   if token.txt == "abstract" then
      abstractFlag = true
      token = self:getToken(  )
   end
   
   if token.txt == "let" then
      return self:analyzeDeclVar( Ast.DeclVarMode.Let, accessMode, firstToken )
   elseif token.txt == "fn" then
      return self:analyzeDeclFunc( DeclFuncMode.Func, abstractFlag, overrideFlag, accessMode, staticFlag, nil, firstToken, nil )
   elseif token.txt == "class" then
      return self:analyzeDeclClass( abstractFlag, accessMode, firstToken, DeclClassMode.Class )
   elseif token.txt == "interface" then
      return self:analyzeDeclClass( true, accessMode, firstToken, DeclClassMode.Interface )
   elseif token.txt == "module" then
      return self:analyzeDeclClass( false, accessMode, firstToken, DeclClassMode.Module )
   elseif token.txt == "proto" then
      return self:analyzeDeclProto( accessMode, firstToken )
   elseif token.txt == "macro" then
      return self:analyzeDeclMacro( accessMode, firstToken )
   elseif token.txt == "enum" then
      return self:analyzeDeclEnum( accessMode, firstToken )
   elseif token.txt == "alge" then
      return self:analyzeDeclAlge( accessMode, firstToken )
   elseif token.txt == "form" then
      self:analyzeDeclForm( accessMode, firstToken )
      return self:createNoneNode( firstToken.pos )
   end
   
   return nil
end

function TransUnit:checkPublic( pos, typeInfo )

   local checkedTypeSet = {}
   local function checkPub( workType )
   
      if checkedTypeSet[workType] then
         return 
      end
      
      checkedTypeSet[workType] = true
      if workType:get_kind() ~= Ast.TypeInfoKind.Array and workType:get_kind() ~= Ast.TypeInfoKind.List and workType:get_kind() ~= Ast.TypeInfoKind.Map and not Ast.isPubToExternal( workType:get_accessMode() ) then
         self:addErrMess( pos, string.format( "not public this type -- %s", workType:getTxt(  )) )
      else
       
         for __index, itemType in pairs( workType:get_itemTypeInfoList() ) do
            checkPub( itemType )
         end
         
      end
      
   end
   
   checkPub( typeInfo )
end

function TransUnit:analyzeDeclMember( classTypeInfo, accessMode, staticFlag, firstToken )

   local nextToken = self:getToken(  )
   local mutable = false
   if nextToken.txt == "mut" then
      mutable = true
      nextToken = self:getToken(  )
   end
   
   local varName = self:checkSymbol( nextToken )
   local token = self:getToken(  )
   local refType = self:analyzeRefType( accessMode, false )
   token = self:getToken(  )
   local getterMode = Ast.AccessMode.None
   local getterMutable = true
   local setterMode = Ast.AccessMode.None
   if token.txt == "{" then
      local function analyzeAccessorMode(  )
      
         local mode = Ast.AccessMode.None
         local workToken = self:getToken(  )
         do
            local _switchExp = workToken.txt
            if _switchExp == "pub" or _switchExp == "pri" or _switchExp == "pro" then
               mode = _lune.unwrap( Ast.txt2AccessMode( workToken.txt ))
               workToken = self:getToken(  )
               if workToken.txt == "&" then
                  getterMutable = false
                  workToken = self:getToken(  )
               end
               
            end
         end
         
         return mode, workToken
      end
      
      getterMode, nextToken = analyzeAccessorMode(  )
      if nextToken.txt == "," then
         setterMode, nextToken = analyzeAccessorMode(  )
      end
      
      self:checkToken( nextToken, "}" )
      token = self:getToken(  )
   end
   
   self:checkToken( token, ";" )
   local typeInfo = refType:get_expType()
   if typeInfo:get_mutable() and not mutable then
      typeInfo = self:createModifier( typeInfo, false )
   end
   
   if Ast.isPubToExternal( classTypeInfo:get_accessMode() ) then
      if Ast.isPubToExternal( accessMode ) or Ast.isPubToExternal( getterMode ) or Ast.isPubToExternal( setterMode ) then
         self:checkPublic( refType:get_pos(), typeInfo )
      end
      
   end
   
   local symbolInfo = self.scope:addMember( varName.txt, typeInfo, accessMode, staticFlag, mutable )
   return Ast.DeclMemberNode.create( self.nodeManager, firstToken.pos, {typeInfo}, varName, refType, symbolInfo, staticFlag, accessMode, getterMutable, getterMode, setterMode )
end

function TransUnit:analyzeDeclMethod( classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, firstToken, name )

   local node = self:analyzeDeclFunc( declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, name, name )
   return node
end

function TransUnit:analyzeClassBody( classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, name, moduleName, nextToken )

   local memberName2Node = {}
   local declStmtList = {}
   local fieldList = {}
   local memberList = {}
   local methodNameSet = {}
   local initStmtList = {}
   local advertiseList = {}
   local trustList = {}
   local node = Ast.DeclClassNode.create( self.nodeManager, firstToken.pos, {classTypeInfo}, classAccessMode, name, gluePrefix, declStmtList, fieldList, moduleName, memberList, self.scope, initStmtList, advertiseList, trustList, {} )
   self.typeInfo2ClassNode[classTypeInfo] = node
   local declCtorNode = nil
   local hasInitBlock = false
   local hasStaticMember = false
   while true do
      local token = self:getToken(  )
      if token.txt == "}" then
         break
      end
      
      local accessMode = Ast.txt2AccessMode( token.txt )
      if  nil == accessMode then
         local _accessMode = accessMode
      
         accessMode = Ast.AccessMode.Pri
      else
         
            token = self:getToken(  )
      end
      
      if mode == DeclClassMode.Interface and accessMode ~= Ast.AccessMode.Pub then
         self:addErrMess( token.pos, "interface's fields must be 'pub'." )
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
      
      local abstractFlag = false
      if token.txt == "abstract" then
         abstractFlag = true
         token = self:getToken(  )
      elseif mode == DeclClassMode.Interface then
         abstractFlag = true
      end
      
      if token.txt == "let" then
         if staticFlag then
            hasStaticMember = true
         end
         
         if mode == DeclClassMode.Interface then
            self:addErrMess( token.pos, "interface can not have member" )
         end
         
         if not staticFlag and declCtorNode then
            self:addErrMess( token.pos, "member can't declare after '__init' method." )
         elseif staticFlag and hasInitBlock then
            self:addErrMess( token.pos, "static member can't declare after '__init' block." )
         end
         
         local memberNode = self:analyzeDeclMember( classTypeInfo, accessMode, staticFlag, token )
         table.insert( fieldList, memberNode )
         table.insert( memberList, memberNode )
         memberName2Node[memberNode:get_name().txt] = memberNode
      elseif token.txt == "fn" then
         local nameToken = self:getSymbolToken(  )
         local declFuncMode = DeclFuncMode.Class
         if mode == DeclClassMode.Module then
            if gluePrefix then
               declFuncMode = DeclFuncMode.Glue
            else
             
               declFuncMode = DeclFuncMode.Module
            end
            
         end
         
         local methodNode = self:analyzeDeclMethod( classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, token, nameToken )
         table.insert( fieldList, methodNode )
         methodNameSet[nameToken.txt] = true
         if nameToken.txt == "__init" then
            declCtorNode = methodNode
         end
         
      elseif token.txt == "__init" then
         if mode ~= DeclClassMode.Class then
            self:error( string.format( "%s can not have __init method", mode) )
         end
         
         hasInitBlock = true
         for symbolName, symbolInfo in pairs( self.scope:get_symbol2SymbolInfoMap() ) do
            if symbolInfo:get_staticFlag() then
               symbolInfo:set_hasValueFlag( false )
            end
            
         end
         
         self:checkNextToken( "{" )
         self:analyzeStatementList( initStmtList, "}" )
         self:checkNextToken( "}" )
      elseif token.txt == "advertise" then
         local memberToken = self:getSymbolToken(  )
         nextToken = self:getToken(  )
         local prefix = ""
         if nextToken.txt ~= ";" and nextToken.txt ~= "{" then
            prefix = nextToken.txt
            nextToken = self:getToken(  )
         end
         
         self:checkToken( nextToken, ";" )
         local memberNode = memberName2Node[memberToken.txt]
         if  nil == memberNode then
            local _memberNode = memberNode
         
            self:error( string.format( "not found member -- %s", memberToken.txt) )
         end
         
         table.insert( advertiseList, Ast.AdvertiseInfo.new(memberNode, prefix) )
      elseif token.txt == ";" then
      elseif token.txt == "enum" then
         if accessMode ~= Ast.AccessMode.Pri and (classAccessMode == Ast.AccessMode.Pri or classAccessMode == Ast.AccessMode.Local ) then
            self:addErrMess( token.pos, string.format( "unmatch access mode, class('%s') and enum('%s')", Ast.AccessMode:_getTxt( classAccessMode)
            , Ast.AccessMode:_getTxt( accessMode)
            ) )
         end
         
         table.insert( declStmtList, self:analyzeDeclEnum( accessMode, token ) )
      else
       
         self:error( "illegal field" )
      end
      
   end
   
   if mode ~= DeclClassMode.Module then
      if declCtorNode ~= nil then
         for memberName, memberNode in pairs( memberName2Node ) do
            if not memberNode:get_staticFlag() then
               local symbolInfo = _lune.unwrap( self.scope:getSymbolInfoChild( memberName ))
               local typeInfo = symbolInfo:get_typeInfo()
               if not symbolInfo:get_hasValueFlag() and not typeInfo:get_nilable() then
                  self:addErrMess( declCtorNode:get_pos(), string.format( "does not set member -- %s.%s", name.txt, memberName) )
               end
               
            end
            
         end
         
      end
      
      if hasStaticMember and not hasInitBlock then
         self:addErrMess( node:get_pos(), string.format( "This class (%s) need __init block for initialize static members.", name.txt) )
      end
      
      for memberName, memberNode in pairs( memberName2Node ) do
         if memberNode:get_staticFlag() then
            local symbolInfo = _lune.unwrap( self.scope:getSymbolInfoChild( memberName ))
            local typeInfo = symbolInfo:get_typeInfo()
            if not symbolInfo:get_hasValueFlag() and not typeInfo:get_nilable() then
               self:addErrMess( memberNode:get_pos(), string.format( "does not set member -- %s", memberName) )
            end
            
         end
         
      end
      
   end
   
   return node, nextToken, methodNameSet
end

function TransUnit:analyzeDeclClass( classAbstructFlag, classAccessMode, firstToken, mode )

   local name = self:getSymbolToken(  )
   if classAccessMode == Ast.AccessMode.Local then
      classAccessMode = Ast.AccessMode.Pri
   end
   
   local moduleName = nil
   local gluePrefix = nil
   if mode == DeclClassMode.Module then
      self:checkNextToken( "require" )
      moduleName = self:getToken(  )
      local nextToken = self:getToken(  )
      if nextToken.txt == "glue" then
         gluePrefix = self:getToken(  ):getExcludedDelimitTxt(  )
      else
       
         self:pushback(  )
      end
      
   end
   
   local nextToken, classTypeInfo = self:analyzePushClass( mode ~= DeclClassMode.Interface, classAbstructFlag, firstToken, name, classAccessMode )
   local classScope = self.scope
   self:checkToken( nextToken, "{" )
   local mapType = self:createModifier( Ast.NormalTypeInfo.createMap( Ast.AccessMode.Pub, classTypeInfo, Ast.builtinTypeString, self:createModifier( Ast.builtinTypeStem, false ) ), false )
   if classTypeInfo:isInheritFrom( _moduleObj.typeInfoMappingIF ) then
      self.helperInfo.hasMappingClassDef = true
      if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo and not classTypeInfo:get_baseTypeInfo():isInheritFrom( _moduleObj.typeInfoMappingIF ) then
         self:addErrMess( firstToken.pos, string.format( "must extend Mapping at %s", classTypeInfo:get_baseTypeInfo():getTxt(  )) )
      end
      
      local toMapFuncTypeInfo = Ast.NormalTypeInfo.createFunc( false, false, nil, Ast.TypeInfoKind.Method, classTypeInfo, true, false, false, Ast.AccessMode.Pub, "_toMap", {}, {mapType}, false )
      classScope:addMethod( toMapFuncTypeInfo, Ast.AccessMode.Pub, false, false )
   end
   
   local node, workNextToken, methodNameSet = self:analyzeClassBody( classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, name, moduleName, nextToken )
   nextToken = workNextToken
   local parentInfo = classTypeInfo
   local memberTypeList = {}
   for __index, memberNode in pairs( node:get_memberList() ) do
      local memberType = memberNode:get_expType()
      if not memberNode:get_staticFlag() then
         table.insert( memberTypeList, memberType )
      end
      
      local memberName = memberNode:get_name()
      local getterName = "get_" .. memberName.txt
      local accessMode = memberNode:get_getterMode()
      if accessMode ~= Ast.AccessMode.None and not classScope:getTypeInfoChild( getterName ) then
         local mutable = memberNode:get_getterMutable()
         local getterMemberType = memberType
         if memberType:get_mutable() and not mutable then
            getterMemberType = self:createModifier( memberType, false )
         end
         
         local retTypeInfo = Ast.NormalTypeInfo.createFunc( false, false, self:pushScope( false ), Ast.TypeInfoKind.Method, parentInfo, true, false, memberNode:get_staticFlag(), accessMode, getterName, {}, {getterMemberType} )
         self:popScope(  )
         classScope:addMethod( retTypeInfo, accessMode, memberNode:get_staticFlag(), false )
         methodNameSet[getterName] = true
      end
      
      local setterName = "set_" .. memberName.txt
      accessMode = memberNode:get_setterMode()
      if memberNode:get_setterMode() ~= Ast.AccessMode.None and not classScope:getTypeInfoChild( setterName ) then
         classScope:addMethod( Ast.NormalTypeInfo.createFunc( false, false, self:pushScope( false ), Ast.TypeInfoKind.Method, parentInfo, true, false, memberNode:get_staticFlag(), accessMode, setterName, {memberType}, nil, true ), accessMode, memberNode:get_staticFlag(), true )
         self:popScope(  )
         methodNameSet[setterName] = true
      end
      
   end
   
   local ctorAccessMode = Ast.AccessMode.Pub
   do
      local ctorTypeInfo = classScope:getTypeInfoChild( "__init" )
      if ctorTypeInfo ~= nil then
         ctorAccessMode = ctorTypeInfo:get_accessMode()
      else
         if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
            local superScope = _lune.unwrap( classTypeInfo:get_baseTypeInfo():get_scope())
            local superTypeInfo = _lune.unwrap( superScope:getTypeInfoChild( "__init" ))
            for __index, argType in pairs( superTypeInfo:get_argTypeInfoList() ) do
               if not argType:get_nilable() then
                  self:addErrMess( firstToken.pos, "not found '__init' decl." )
               end
               
            end
            
         end
         
         local initTypeInfo = Ast.NormalTypeInfo.createFunc( false, false, self:pushScope( false ), Ast.TypeInfoKind.Method, parentInfo, true, false, false, Ast.AccessMode.Pub, "__init", memberTypeList, {} )
         self:popScope(  )
         classScope:addMethod( initTypeInfo, Ast.AccessMode.Pub, false, false )
         methodNameSet["__init"] = true
         for __index, memberNode in pairs( node:get_memberList() ) do
            if not memberNode:get_staticFlag() then
               memberNode:get_symbolInfo():set_hasValueFlag( true )
            end
            
         end
         
      end
   end
   
   for __index, advertiseInfo in pairs( node:get_advertiseList() ) do
      local memberType = advertiseInfo:get_member():get_expType()
      do
         local _switchExp = memberType:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
            for __index, child in pairs( memberType:get_children() ) do
               if child:get_kind() == Ast.TypeInfoKind.Method and child:get_accessMode() ~= Ast.AccessMode.Pri and not child:get_staticFlag() then
                  local childName = advertiseInfo:get_prefix() .. child:getTxt(  )
                  if not methodNameSet[childName] then
                     local impMtdType = Ast.NormalTypeInfo.createAdvertiseMethodFrom( classTypeInfo, child )
                     classScope:addMethod( impMtdType, child:get_accessMode(), child:get_staticFlag(), false )
                  end
                  
               end
               
            end
            
         else 
            
               self:error( string.format( "advertise member type is illegal -- %s", advertiseInfo:get_member():get_name()) )
         end
      end
      
   end
   
   if classTypeInfo:isInheritFrom( _moduleObj.typeInfoMappingIF ) then
      local function isAvailableMapping( typeInfo, checkedTypeMap )
      
         local function isAvailableMappingSub(  )
         
            do
               local _switchExp = typeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.Prim or _switchExp == Ast.TypeInfoKind.Enum then
                  return true
               elseif _switchExp == Ast.TypeInfoKind.Alge then
                  local algeTypeInfo = typeInfo
                  for __index, valInfo in pairs( algeTypeInfo:get_valInfoMap() ) do
                     for __index, paramType in pairs( valInfo:get_typeList() ) do
                        if not isAvailableMapping( paramType, checkedTypeMap ) then
                           return false
                        end
                        
                     end
                     
                  end
                  
                  return true
               elseif _switchExp == Ast.TypeInfoKind.Stem then
                  return true
               elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
                  if typeInfo:equals( Ast.builtinTypeString ) then
                     return true
                  end
                  
                  return typeInfo:isInheritFrom( _moduleObj.typeInfoMappingIF )
               elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
                  return isAvailableMapping( typeInfo:get_itemTypeInfoList()[1], checkedTypeMap )
               elseif _switchExp == Ast.TypeInfoKind.Map then
                  if isAvailableMapping( typeInfo:get_itemTypeInfoList()[2], checkedTypeMap ) then
                     local keyType = typeInfo:get_itemTypeInfoList()[1]
                     if keyType:equals( Ast.builtinTypeString ) or keyType:get_kind() == Ast.TypeInfoKind.Prim or keyType:get_kind() == Ast.TypeInfoKind.Enum then
                        return true
                     end
                     
                  end
                  
                  return false
               elseif _switchExp == Ast.TypeInfoKind.Nilable then
                  return isAvailableMapping( typeInfo:get_orgTypeInfo(), checkedTypeMap )
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
      
      local checkedTypeMap = {}
      for __index, memberNode in pairs( node:get_memberList() ) do
         local memberType = memberNode:get_expType()
         if not isAvailableMapping( memberType, checkedTypeMap ) then
            self:addErrMess( memberNode:get_pos(), string.format( "member type is not Mapping -- %s", memberType:getTxt(  )) )
         end
         
      end
      
      local fromMapFuncTypeInfo = Ast.NormalTypeInfo.createFunc( false, false, nil, Ast.TypeInfoKind.Func, classTypeInfo, true, false, true, Ast.AccessMode.Pub, "_fromMap", {mapType:get_nilableTypeInfo()}, {classTypeInfo:get_nilableTypeInfo(), Ast.builtinTypeString:get_nilableTypeInfo()}, true )
      classScope:addMethod( fromMapFuncTypeInfo, ctorAccessMode, true, false )
      local fromStemFuncTypeInfo = Ast.NormalTypeInfo.createFunc( false, false, nil, Ast.TypeInfoKind.Func, classTypeInfo, true, false, true, Ast.AccessMode.Pub, "_fromStem", {Ast.builtinTypeStem_}, {classTypeInfo:get_nilableTypeInfo(), Ast.builtinTypeString:get_nilableTypeInfo()}, true )
      classScope:addMethod( fromStemFuncTypeInfo, ctorAccessMode, true, false )
   end
   
   local function checkOverrideMethod( scope )
   
      scope:filterTypeInfoField( true, classScope, function ( symbolInfo )
      
         if symbolInfo:get_kind() == Ast.SymbolKind.Mtd then
            local noImp = false
            do
               local impMethodType = classScope:getTypeInfoField( symbolInfo:get_name(), true, classScope )
               if impMethodType ~= nil then
                  if impMethodType:get_abstractFlag() then
                     noImp = true
                  end
                  
               else
                  noImp = true
               end
            end
            
            if noImp then
               self:addErrMess( firstToken.pos, "not implements method -- " .. symbolInfo:get_name() )
            end
            
         end
         
         return true
      end
       )
   end
   
   if not classAbstructFlag then
      local workTypeInfo = classTypeInfo
      repeat 
         if workTypeInfo ~= Ast.headTypeInfo then
            checkOverrideMethod( _lune.unwrap( workTypeInfo:get_scope()) )
         end
         
         for __index, ifType in pairs( workTypeInfo:get_interfaceList() ) do
            checkOverrideMethod( _lune.unwrap( ifType:get_scope()) )
         end
         
         workTypeInfo = workTypeInfo:get_baseTypeInfo()
      until workTypeInfo == Ast.headTypeInfo
   end
   
   self:popClass(  )
   return node
end

function TransUnit:addMethod( classTypeInfo, methodNode, name )

   local classNodeInfo = _lune.unwrap( self.typeInfo2ClassNode[classTypeInfo])
   classNodeInfo:get_outerMethodSet()[name] = true
   table.insert( classNodeInfo:get_fieldList(), methodNode )
end

function TransUnit:analyzeDeclFunc( declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, firstToken, name )

   local token = self:getToken(  )
   do
      local _exp = name
      if _exp ~= nil then
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
      local className = (_lune.unwrap( name) ).txt
      classTypeInfo = self.scope:getTypeInfoChild( className )
      if classTypeInfo ~= nil then
         self:pushClass( classTypeInfo:get_kind() == Ast.TypeInfoKind.Class, classTypeInfo:get_abstractFlag(), nil, nil, false, className, classTypeInfo:get_accessMode() )
      else
         self:error( string.format( "not found class -- %s", className) )
      end
      
      name = self:getSymbolToken(  )
      token = self:getToken(  )
      if accessMode == Ast.AccessMode.Local then
         accessMode = Ast.AccessMode.Pri
      end
      
   end
   
   local isCtorFlag = false
   local kind = Ast.NodeKind.get_DeclConstr()
   local typeKind = Ast.TypeInfoKind.Func
   if classTypeInfo then
      if not staticFlag then
         typeKind = Ast.TypeInfoKind.Method
      end
      
      do
         local _switchExp = (_lune.unwrap( name) ).txt
         if _switchExp == "__init" then
            isCtorFlag = true
            kind = Ast.NodeKind.get_DeclConstr()
            for symbolName, symbolInfo in pairs( self.scope:get_symbol2SymbolInfoMap() ) do
               if not symbolInfo:get_staticFlag() then
                  symbolInfo:set_hasValueFlag( false )
               end
               
            end
            
         elseif _switchExp == "__free" then
            kind = Ast.NodeKind.get_DeclDestr()
            if not self.targetLuaVer:get_canUseMetaGc() then
               self:addErrMess( firstToken.pos, "this lua version is not support __free." )
            end
            
         else 
            
               kind = Ast.NodeKind.get_DeclMethod()
         end
      end
      
   else
    
      kind = Ast.NodeKind.get_DeclFunc()
      if not staticFlag then
         staticFlag = true
      end
      
   end
   
   local orgStaticFlag = staticFlag
   if declFuncMode == DeclFuncMode.Module then
      staticFlag = true
   end
   
   local funcName = ""
   if name ~= nil then
      funcName = name.txt
      if kind == Ast.NodeKind.get_DeclFunc() then
         do
            local _switchExp = accessMode
            if _switchExp == Ast.AccessMode.Pub or _switchExp == Ast.AccessMode.Global then
               if self.scope ~= self.moduleScope then
                  self:addErrMess( firstToken.pos, "'global' or 'pub' function must exist top scope." )
               end
               
            end
         end
         
      end
      
   end
   
   self:checkToken( token, "(" )
   local funcBodyScope = self:pushScope( false )
   local argList = {}
   token = self:analyzeDeclArgList( accessMode, argList )
   local argTypeList = {}
   for __index, argNode in pairs( argList ) do
      table.insert( argTypeList, argNode:get_expType() )
   end
   
   self:checkToken( token, ")" )
   token = self:getToken(  )
   local mutable = false
   if token.txt == "mut" then
      token = self:getToken(  )
      mutable = true
   end
   
   local pubToExtFlag = Ast.isPubToExternal( accessMode )
   if classTypeInfo ~= nil then
      if kind == Ast.NodeKind.get_DeclMethod() or kind == Ast.NodeKind.get_DeclConstr() or kind == Ast.NodeKind.get_DeclDestr() then
         local workClass = classTypeInfo
         if kind == Ast.NodeKind.get_DeclConstr() or kind == Ast.NodeKind.get_DeclDestr() then
            mutable = true
         end
         
         if not Ast.isPubToExternal( workClass:get_accessMode() ) then
            pubToExtFlag = false
         end
         
         if workClass:get_mutable() and not mutable then
            workClass = self:createModifier( workClass, false )
         end
         
         if not staticFlag then
            self.scope:add( Ast.SymbolKind.Var, false, true, "self", workClass, Ast.AccessMode.Pri, false, mutable, true )
         end
         
         if not workClass:get_abstractFlag() and abstractFlag then
            self:addErrMess( firstToken.pos, "no abstract class does not have abstract method" )
         end
         
      end
      
   end
   
   local retTypeInfoList = {}
   retTypeInfoList, token = self:analyzeRetTypeList( pubToExtFlag, accessMode, token )
   local namespaceInfo = self:getCurrentNamespaceTypeInfo(  )
   local typeInfo = Ast.NormalTypeInfo.createFunc( abstractFlag, false, funcBodyScope, typeKind, namespaceInfo, false, false, staticFlag, accessMode, funcName, argTypeList, retTypeInfoList, mutable )
   if name ~= nil then
      local parentScope = funcBodyScope:get_parent(  )
      if accessMode == Ast.AccessMode.Global then
         parentScope = self.globalScope
      end
      
      do
         local prottype = parentScope:getTypeInfoChild( typeInfo:get_rawTxt() )
         if prottype ~= nil then
            local matchFlag, err = Ast.TypeInfo.checkMatchType( prottype:get_argTypeInfoList(), argTypeList, false )
            if not matchFlag then
               self:addErrMess( name.pos, err )
            end
            
            if self.protoFuncMap[prottype] then
               self.protoFuncMap[prottype] = nil
            else
             
               if not prottype:get_autoFlag() then
                  self:addErrMess( token.pos, string.format( "multiple define -- %s", name.txt) )
               end
               
            end
            
         end
      end
      
      if kind == Ast.NodeKind.get_DeclFunc() then
         parentScope:addFunc( typeInfo, accessMode, staticFlag, mutable )
      else
       
         parentScope:addMethod( typeInfo, accessMode, staticFlag, mutable )
      end
      
   end
   
   if overrideFlag then
      if not name then
         self:addErrMess( firstToken.pos, "can't override anonymous func" )
      end
      
      
      local overrideType = self.scope:get_parent():getTypeInfoField( funcName, false, funcBodyScope )
      if  nil == overrideType then
         local _overrideType = overrideType
      
         self:addErrMess( firstToken.pos, "not found override -- " .. funcName )
      else
         
            if overrideType:get_accessMode(  ) ~= accessMode then
               self:addErrMess( firstToken.pos, string.format( "mismatch override accessMode -- %s,%s,%s", funcName, Ast.AccessMode:_getTxt( overrideType:get_accessMode(  ))
               , Ast.AccessMode:_getTxt( accessMode)
               ) )
            end
            
            if overrideType:get_staticFlag(  ) ~= staticFlag then
               self:addErrMess( firstToken.pos, "mismatch override staticFlag -- " .. funcName )
            end
            
            if overrideType:get_kind(  ) ~= Ast.TypeInfoKind.Method then
               self:addErrMess( firstToken.pos, string.format( "mismatch override kind -- %s, %d", funcName, overrideType:get_kind(  )) )
            end
            
            if overrideType:get_mutable() ~= typeInfo:get_mutable() then
               self:addErrMess( firstToken.pos, string.format( "mismatch mutable -- %s", funcName) )
            end
            
            if not overrideType:canEvalWith( typeInfo, "=" ) then
               self:addErrMess( firstToken.pos, string.format( "mismatch method type -- %s", funcName) )
            end
            
      end
      
   else
    
      if name ~= nil then
         if name.txt ~= "__init" then
            if self.scope:get_parent():getTypeInfoField( name.txt, false, funcBodyScope ) then
               self:addErrMess( firstToken.pos, "mismatch override --" .. funcName )
            else
             
               do
                  local ifFunc = self.scope:get_parent():getSymbolInfoIfField( name.txt, funcBodyScope )
                  if ifFunc ~= nil then
                     if not ifFunc:get_typeInfo():canEvalWith( typeInfo, "=" ) then
                        self:addErrMess( firstToken.pos, string.format( "mismatch method type -- %s", funcName) )
                     end
                     
                  end
               end
               
            end
            
         end
         
      end
      
   end
   
   local node = self:createNoneNode( firstToken.pos )
   local needNode = false
   local body = nil
   if token.txt == ";" then
      if declFuncMode == DeclFuncMode.Module or declFuncMode == DeclFuncMode.Glue then
         needNode = true
      else
       
         if not abstractFlag then
            self.protoFuncMap[typeInfo] = firstToken.pos
         end
         
      end
      
   else
    
      needNode = true
      if abstractFlag then
         self:addErrMess( token.pos, "abstract method can't have body." )
      end
      
      local __func__Symbol = funcBodyScope:addLocalVar( false, false, "__func__", Ast.builtinTypeString, false )
      self.has__func__Symbol = false
      if classTypeInfo ~= nil then
         do
            local overrideType = self.scope:get_parent():getTypeInfoField( funcName, false, funcBodyScope )
            if overrideType ~= nil then
               if not overrideType:get_abstractFlag() then
                  funcBodyScope:addLocalVar( false, false, "super", overrideType, false )
               end
               
            end
         end
         
      end
      
      self:pushback(  )
      body = self:analyzeBlock( Ast.BlockKind.Func, funcBodyScope )
      if body ~= nil then
         if #retTypeInfoList ~= 0 then
            local breakKind = body:getBreakKind( Ast.CheckBreakMode.Return )
            if retTypeInfoList[1] ~= Ast.builtinTypeNeverRet then
               do
                  local _switchExp = breakKind
                  if _switchExp == Ast.BreakKind.Return or _switchExp == Ast.BreakKind.NeverRet then
                  else 
                     
                        self:addErrMess( firstToken.pos, "This funcion doesn't have return." )
                  end
               end
               
            else
             
               if breakKind ~= Ast.BreakKind.NeverRet then
                  self:addErrMess( firstToken.pos, string.format( "This funcion must be never return. -- %s", Ast.BreakKind:_getTxt( breakKind)
                  ) )
               end
               
            end
            
         end
         
         if isCtorFlag then
            if classTypeInfo ~= nil then
               if classTypeInfo:get_baseTypeInfo() ~= Ast.headTypeInfo then
                  if #body:get_stmtList() == 0 or body:get_stmtList()[1]:get_kind() ~= Ast.nodeKind['ExpCallSuper'] then
                     self:addErrMess( body:get_pos(), "__init must call super() with first." )
                  end
                  
               end
               
            end
            
         end
         
      end
      
   end
   
   if needNode then
      local info = Ast.DeclFuncInfo.new(classTypeInfo, name, argList, orgStaticFlag, accessMode, body, retTypeInfoList, self.has__func__Symbol)
      do
         local _switchExp = (kind )
         if _switchExp == Ast.NodeKind.get_DeclConstr() then
            node = Ast.DeclConstrNode.create( self.nodeManager, firstToken.pos, {typeInfo}, info )
         elseif _switchExp == Ast.NodeKind.get_DeclDestr() then
            node = Ast.DeclDestrNode.create( self.nodeManager, firstToken.pos, {typeInfo}, info )
         elseif _switchExp == Ast.NodeKind.get_DeclMethod() then
            node = Ast.DeclMethodNode.create( self.nodeManager, firstToken.pos, {typeInfo}, info )
         elseif _switchExp == Ast.NodeKind.get_DeclFunc() then
            node = Ast.DeclFuncNode.create( self.nodeManager, firstToken.pos, {typeInfo}, info )
         else 
            
               self:error( string.format( "illegal kind -- %d", kind) )
         end
      end
      
   end
   
   self:popScope(  )
   if needPopFlag then
      self:addMethod( _lune.unwrap( classTypeInfo), node, funcName )
      self:popClass(  )
   end
   
   return node
end

local LetVarInfo = {}
function LetVarInfo.setmeta( obj )
  setmetatable( obj, { __index = LetVarInfo  } )
end
function LetVarInfo.new( mutable, varName, varType )
   local obj = {}
   LetVarInfo.setmeta( obj )
   if obj.__init then
      obj:__init( mutable, varName, varType )
   end        
   return obj 
end         
function LetVarInfo:__init( mutable, varName, varType ) 

   self.mutable = mutable
   self.varName = varName
   self.varType = varType
end

function TransUnit:analyzeLetAndInitExp( firstPos, initMutable, accessMode, unwrapFlag )

   local typeInfoList = {}
   local letVarList = {}
   local nextToken = Parser.getEofToken(  )
   repeat 
      local mutable = initMutable
      nextToken = self:getToken(  )
      if nextToken.txt == "mut" then
         mutable = true
         nextToken = self:getToken(  )
      end
      
      local varName = self:checkSymbol( nextToken )
      nextToken = self:getToken(  )
      local typeInfo = Ast.builtinTypeNone
      if nextToken.txt == ":" then
         local refType = self:analyzeRefType( accessMode, false )
         table.insert( letVarList, LetVarInfo.new(mutable, varName, refType) )
         typeInfo = refType:get_expType()
         nextToken = self:getToken(  )
      else
       
         table.insert( letVarList, LetVarInfo.new(mutable, varName, nil) )
      end
      
      if not typeInfo:equals( Ast.builtinTypeNone ) and typeInfo:get_mutable() and not mutable then
         typeInfo = self:createModifier( typeInfo, false )
      end
      
      table.insert( typeInfoList, typeInfo )
   until nextToken.txt ~= ","
   local expList = nil
   if nextToken.txt == "=" then
      expList = self:analyzeExpList( false )
      if not expList then
         self:error( "expList is nil" )
      end
      
   end
   
   local orgExpTypeList = {}
   if expList ~= nil then
      for index, exp in pairs( expList:get_expList() ) do
         if not exp:canBeRight(  ) then
            self:addErrMess( exp:get_pos(), string.format( "this node can not be r-value. -- %s", Ast.getNodeKindName( exp:get_kind() )) )
         end
         
      end
      
      local expTypeList = {}
      for index, expType in pairs( expList:get_expTypeList() ) do
         local processedFlag = false
         if index == #expList:get_expTypeList() and expList:get_expTypeList()[index]:equals( Ast.builtinTypeDDD ) then
            for subIndex = index, #letVarList do
               local argType = typeInfoList[subIndex]
               local checkType = Ast.builtinTypeStem_
               if unwrapFlag then
                  checkType = Ast.builtinTypeStem
               end
               
               if not argType:equals( Ast.builtinTypeNone ) and not argType:canEvalWith( checkType, "=" ) then
                  self:addErrMess( firstPos, string.format( "unmatch value type (index = %d) %s(%d) <- %s(%d)", subIndex, argType:getTxt( true ), argType:get_typeId(), Ast.builtinTypeStem_:getTxt(  ), Ast.builtinTypeStem_:get_typeId()) )
               end
               
               table.insert( expTypeList, checkType )
               table.insert( orgExpTypeList, Ast.builtinTypeStem_ )
            end
            
         else
          
            local expTypeInfo = expType
            if expType:equals( Ast.builtinTypeDDD ) then
               expTypeInfo = Ast.builtinTypeStem_
            end
            
            table.insert( orgExpTypeList, expTypeInfo )
            if unwrapFlag and expTypeInfo:get_nilable() then
               expTypeInfo = _lune.unwrap( expTypeInfo:get_orgTypeInfo())
            end
            
            if index <= #typeInfoList then
               local argType = typeInfoList[index]
               if not argType:equals( Ast.builtinTypeNone ) and not argType:canEvalWith( expTypeInfo, "=" ) and not (unwrapFlag and expTypeInfo:equals( Ast.builtinTypeNil ) ) then
                  self:addErrMess( firstPos, string.format( "unmatch value type (index:%d) %s <- %s", index, argType:getTxt( true ), expTypeInfo:getTxt( true )) )
               end
               
            end
            
            table.insert( expTypeList, expTypeInfo )
         end
         
      end
      
      for index, varType in pairs( typeInfoList ) do
         if index > #expTypeList then
            if not varType:get_nilable() then
               self:addErrMess( firstPos, string.format( "unmatch value type (index:%d) %s <- nil", index, varType:getTxt( true )) )
            end
            
         end
         
      end
      
      for index, typeInfo in pairs( expTypeList ) do
         if #typeInfoList < index or typeInfoList[index]:equals( Ast.builtinTypeNone ) then
            if typeInfo:get_mutable() and index <= #letVarList and not letVarList[index].mutable then
               typeInfoList[index] = self:createModifier( typeInfo, false )
            else
             
               typeInfoList[index] = typeInfo
            end
            
         end
         
      end
      
   end
   
   return typeInfoList, letVarList, orgExpTypeList, expList
end

function TransUnit:analyzeDeclVar( mode, accessMode, firstToken )

   local unwrapFlag = false
   local token, continueFlag = self:getContinueToken(  )
   if continueFlag and token.txt == "!" then
      unwrapFlag = true
   else
    
      self:pushback(  )
      if mode ~= Ast.DeclVarMode.Let then
         Util.log( "need '!'" )
      end
      
   end
   
   if accessMode == Ast.AccessMode.Pub then
      if self.scope ~= self.moduleScope then
         self:addErrMess( firstToken.pos, "'pub' variable must exist top scope." )
      end
      
   end
   
   local typeInfoList, letVarList, orgExpTypeList, expList = self:analyzeLetAndInitExp( firstToken.pos, mode == Ast.DeclVarMode.Sync, accessMode, unwrapFlag )
   if mode ~= Ast.DeclVarMode.Sync and self.macroScope then
      for index, letVarInfo in pairs( letVarList ) do
         local typeInfo = typeInfoList[index]
         self.symbol2ValueMapForMacro[letVarInfo.varName.txt] = Ast.MacroValInfo.new(nil, typeInfo)
      end
      
   end
   
   local syncScope = self.scope
   if mode == Ast.DeclVarMode.Sync then
      syncScope = self:pushScope( false )
   end
   
   local symbolInfoList = {}
   local varList = {}
   local syncSymbolList = {}
   for index, letVarInfo in pairs( letVarList ) do
      local varName = letVarInfo.varName
      local typeInfo = typeInfoList[index]
      local varInfo = Ast.VarInfo.new(varName, letVarInfo.varType, typeInfo)
      table.insert( varList, varInfo )
      if Ast.isPubToExternal( accessMode ) then
         self:checkPublic( varName.pos, typeInfo )
      end
      
      if not letVarInfo.varType and typeInfo:equals( Ast.builtinTypeNil ) then
         self:addErrMess( varName.pos, string.format( 'need type -- %s', varName.txt) )
      end
      
      if mode == Ast.DeclVarMode.Sync then
         if self.scope:getTypeInfo( varName.txt, self.scope, true ) then
            table.insert( syncSymbolList, varInfo )
         end
         
      end
      
      if mode == Ast.DeclVarMode.Let or mode == Ast.DeclVarMode.Sync then
         if mode == Ast.DeclVarMode.Let then
            if self.scope:getTypeInfo( varName.txt, self.scope, true ) then
               self:addErrMess( varName.pos, string.format( "shadowing variable -- %s", varName.txt) )
            end
            
         end
         
         self.scope:addVar( accessMode, varName.txt, typeInfo, letVarInfo.mutable, not unwrapFlag )
      end
      
      table.insert( symbolInfoList, _lune.unwrap( self.scope:getSymbolInfo( varName.txt, self.scope, true )) )
   end
   
   local unwrapBlock = nil
   local thenBlock = nil
   if unwrapFlag then
      local scope = self:pushScope( false )
      for index, letVarInfo in pairs( letVarList ) do
         self:addLocalVar( letVarInfo.varName.pos, false, true, "_" .. letVarInfo.varName.txt, orgExpTypeList[index], false )
      end
      
      unwrapBlock = self:analyzeBlock( Ast.BlockKind.LetUnwrap, scope )
      self:popScope(  )
      if mode == Ast.DeclVarMode.Let or mode == Ast.DeclVarMode.Sync then
         for index, letVarInfo in pairs( letVarList ) do
            local symbolInfo = _lune.unwrap( self.scope:getSymbolInfoChild( letVarInfo.varName.txt ))
            symbolInfo:set_hasValueFlag( true )
         end
         
      end
      
      token = self:getToken( true )
      if token.txt == "then" then
         thenBlock = self:analyzeBlock( Ast.BlockKind.LetUnwrap, scope )
      else
       
         self:pushback(  )
      end
      
   end
   
   local syncBlock = nil
   if mode == Ast.DeclVarMode.Sync then
      self:checkNextToken( "do" )
      syncBlock = self:analyzeBlock( Ast.BlockKind.LetUnwrap, syncScope )
      self:popScope(  )
   end
   
   self:checkNextToken( ";" )
   local node = Ast.DeclVarNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, mode, accessMode, false, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncSymbolList, syncBlock )
   return node
end

function TransUnit:analyzeIfUnwrap( firstToken )

   local nextToken = self:getToken(  )
   local typeInfoList = {}
   local varNameList = {}
   local expNodeList = {}
   if nextToken.txt == "let" then
      local workTypeInfoList, letVarList, orgExpTypeList, expList = self:analyzeLetAndInitExp( firstToken.pos, false, Ast.AccessMode.Local, true )
      typeInfoList = workTypeInfoList
      for __index, exp in pairs( (_lune.unwrap( expList) ):get_expList() ) do
         table.insert( expNodeList, exp )
      end
      
      for __index, varInfo in pairs( letVarList ) do
         table.insert( varNameList, varInfo.varName.txt )
      end
      
   else
    
      self:pushback(  )
      local exp = self:analyzeExp( false )
      table.insert( expNodeList, exp )
      if exp:get_expType():get_nilable() then
         table.insert( typeInfoList, exp:get_expType():get_orgTypeInfo() )
      else
       
         table.insert( typeInfoList, exp:get_expType() )
      end
      
      table.insert( varNameList, "_exp" )
   end
   
   local scope = self:pushScope( false )
   for index, expType in pairs( typeInfoList ) do
      if index > #varNameList then
         break
      end
      
      local varName = varNameList[index]
      self:addLocalVar( firstToken.pos, false, true, varName, expType, false )
   end
   
   local block = self:analyzeBlock( Ast.BlockKind.IfUnwrap, scope )
   self:popScope(  )
   local elseBlock = nil
   nextToken = self:getToken( true )
   if nextToken.txt == "else" then
      elseBlock = self:analyzeBlock( Ast.BlockKind.IfUnwrap )
   else
    
      self:pushback(  )
   end
   
   return Ast.IfUnwrapNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, varNameList, expNodeList, block, elseBlock )
end

function TransUnit:analyzeWhen( firstToken )

   local nextToken, continueFlag = self:getContinueToken(  )
   local varNameList = {}
   if not (continueFlag and nextToken.txt == "!" ) then
      self:pushback(  )
      self:addErrMess( nextToken.pos, "'when' need '!'" )
   end
   
   local symListNode = self:analyzeExpList( false )
   local scope = self:pushScope( false )
   local expNodeList = {}
   for __index, expNode in pairs( symListNode:get_expList() ) do
      table.insert( expNodeList, expNode )
      if expNode:get_kind() ~= Ast.NodeKind.get_ExpRef() then
         self:addErrMess( expNode:get_pos(), "'when' support only local variables or arguments." )
      else
       
         if expNode:get_expType():get_nilable() then
            local refNode = expNode
            local symbolInfo = refNode:get_symbolInfo()
            table.insert( varNameList, refNode:get_token().txt )
            self:addLocalVar( firstToken.pos, false, expNode:canBeLeft(  ), refNode:get_token().txt, expNode:get_expType():get_orgTypeInfo(), symbolInfo:get_mutable(), true )
         else
          
            self:addErrMess( expNode:get_pos(), string.format( "This type isn't nilable. -- %s", expNode:get_expType():getTxt(  )) )
         end
         
      end
      
   end
   
   local block = self:analyzeBlock( Ast.BlockKind.IfUnwrap, scope )
   self:popScope(  )
   local elseBlock = nil
   nextToken = self:getToken( true )
   if nextToken.txt == "else" then
      elseBlock = self:analyzeBlock( Ast.BlockKind.When )
   else
    
      self:pushback(  )
   end
   
   return Ast.WhenNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, varNameList, expNodeList, block, elseBlock )
end

function TransUnit:analyzeExpList( skipOp2Flag, expNode, expectTypeList, contExpect )

   local expList = {}
   local pos = nil
   local expTypeList = {}
   if expNode ~= nil then
      pos = expNode:get_pos()
      table.insert( expList, expNode )
      table.insert( expTypeList, expNode:get_expType() )
   end
   
   local index = 1
   repeat 
      local expectType = nil
      if expectTypeList ~= nil then
         if #expectTypeList > 0 then
            local checkIndex = index
            if index > #expectTypeList and contExpect then
               checkIndex = #expectTypeList
            end
            
            if checkIndex <= #expectTypeList and expectTypeList[checkIndex] ~= Ast.builtinTypeNone then
               expectType = expectTypeList[checkIndex]
            end
            
         end
         
      end
      
      local exp = self:analyzeExp( skipOp2Flag, 0, expectType )
      if not pos then
         pos = exp:get_pos()
      end
      
      table.insert( expList, exp )
      table.insert( expTypeList, exp:get_expType() )
      local token = self:getToken(  )
      index = index + 1
   until token.txt ~= ","
   for listIndex, expType in pairs( expList[#expList]:get_expTypeList() ) do
      if listIndex ~= 1 then
         table.insert( expTypeList, expType )
      end
      
   end
   
   self:pushback(  )
   return Ast.ExpListNode.create( self.nodeManager, _lune.unwrapDefault( pos, Parser.Position.new(0, 0)), expTypeList, expList )
end

function TransUnit:analyzeListConst( token )

   local nextToken = self:getToken(  )
   local expList = nil
   local itemTypeInfo = Ast.builtinTypeNone
   if nextToken.txt ~= "]" then
      self:pushback(  )
      expList = self:analyzeExpList( false )
      self:checkNextToken( "]" )
      local nodeList = (_lune.unwrap( expList) ):get_expList()
      for __index, exp in pairs( nodeList ) do
         local expType = exp:get_expType()
         if itemTypeInfo:equals( Ast.builtinTypeNone ) then
            itemTypeInfo = expType
         elseif not itemTypeInfo:canEvalWith( expType, "=" ) then
            if expType:equals( Ast.builtinTypeNil ) then
               itemTypeInfo = _lune.unwrap( itemTypeInfo:get_nilableTypeInfo())
            elseif expType:get_nilable() then
               itemTypeInfo = Ast.builtinTypeStem_
            else
             
               itemTypeInfo = Ast.builtinTypeStem
            end
            
         end
         
      end
      
   end
   
   if itemTypeInfo:equals( Ast.builtinTypeDDD ) then
      itemTypeInfo = Ast.builtinTypeStem_
   end
   
   local kind = Ast.NodeKind.get_LiteralArray()
   local typeInfoList = {Ast.builtinTypeNone}
   if token.txt == '[' then
      kind = Ast.NodeKind.get_LiteralList()
      typeInfoList = {Ast.NormalTypeInfo.createList( Ast.AccessMode.Local, self:getCurrentClass(  ), {itemTypeInfo} )}
      return Ast.LiteralListNode.create( self.nodeManager, token.pos, typeInfoList, expList )
   else
    
      typeInfoList = {Ast.NormalTypeInfo.createArray( Ast.AccessMode.Local, self:getCurrentClass(  ), {itemTypeInfo} )}
      return Ast.LiteralArrayNode.create( self.nodeManager, token.pos, typeInfoList, expList )
   end
   
end

function TransUnit:analyzeMapConst( token )

   local nextToken = self:getToken(  )
   local map = {}
   local pairList = {}
   local keyTypeInfo = Ast.builtinTypeNone
   local valTypeInfo = Ast.builtinTypeNone
   local function getMapKeyValType( pos, keyFlag, typeInfo, expType )
   
      if expType:get_nilable() then
         if keyFlag then
            self:addErrMess( pos, string.format( "map key can't set a nilable -- %s", expType:getTxt(  )) )
         end
         
         if expType:equals( Ast.builtinTypeNil ) then
            return typeInfo
         end
         
         expType = _lune.unwrap( expType:get_orgTypeInfo())
      end
      
      if not typeInfo:canEvalWith( expType, "=" ) then
         if not typeInfo:equals( Ast.builtinTypeNone ) then
            typeInfo = Ast.builtinTypeStem
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
      local key = self:analyzeExp( false )
      keyTypeInfo = getMapKeyValType( key:get_pos(), true, keyTypeInfo, key:get_expType() )
      self:checkNextToken( ":" )
      local val = self:analyzeExp( false )
      valTypeInfo = getMapKeyValType( val:get_pos(), false, valTypeInfo, val:get_expType() )
      table.insert( pairList, Ast.PairItem.new(key, val) )
      map[key] = val
      nextToken = self:getToken(  )
      if nextToken.txt ~= "," then
         break
      end
      
      nextToken = self:getToken(  )
   end
   
   local typeInfo = Ast.NormalTypeInfo.createMap( Ast.AccessMode.Local, self:getCurrentClass(  ), keyTypeInfo, valTypeInfo )
   self:checkToken( nextToken, "}" )
   return Ast.LiteralMapNode.create( self.nodeManager, token.pos, {typeInfo}, map, pairList )
end

function TransUnit:analyzeExpRefItem( token, exp, nilAccess )

   local expType = exp:get_expType()
   if nilAccess then
      if not expType:get_nilable() then
         nilAccess = false
      else
       
         expType = _lune.unwrap( expType:get_orgTypeInfo())
      end
      
   end
   
   local expectItemType = nil
   local typeInfo = Ast.builtinTypeStem_
   if expType:get_kind() == Ast.TypeInfoKind.Map then
      local itemTypeList = expType:get_itemTypeInfoList(  )
      typeInfo = itemTypeList[2]
      expectItemType = itemTypeList[1]
      if not typeInfo:equals( Ast.builtinTypeStem_ ) and not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
   elseif expType:get_kind() == Ast.TypeInfoKind.Array or expType:get_kind() == Ast.TypeInfoKind.List then
      typeInfo = expType:get_itemTypeInfoList(  )[1]
   elseif expType:equals( Ast.builtinTypeString ) then
      typeInfo = Ast.builtinTypeInt
   elseif expType:equals( Ast.builtinTypeStem ) then
      typeInfo = Ast.builtinTypeStem
   else
    
      self:addErrMess( exp:get_pos(), string.format( "could not access with []. -- %s", expType:getTxt(  )) )
   end
   
   if nilAccess then
      self.helperInfo.useNilAccess = true
      if not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
   end
   
   if typeInfo:get_mutable() and not expType:get_mutable() then
      typeInfo = self:createModifier( typeInfo, false )
   end
   
   local indexExp = self:analyzeExp( false, nil, expectItemType )
   self:checkNextToken( "]" )
   return Ast.ExpRefItemNode.create( self.nodeManager, token.pos, {typeInfo}, exp, nilAccess, nil, indexExp )
end

function TransUnit:checkMatchType( message, pos, dstTypeList, expNodeList, allowDstShort )

   local expTypeList = {}
   for index, expNode in pairs( expNodeList ) do
      if index == #expNodeList then
         for __index, expType in pairs( expNode:get_expTypeList() ) do
            table.insert( expTypeList, expType )
         end
         
      else
       
         table.insert( expTypeList, expNode:get_expType() )
      end
      
   end
   
   local match, mess = Ast.TypeInfo.checkMatchType( dstTypeList, expTypeList, allowDstShort )
   if not match then
      self:addErrMess( pos, string.format( "%s: %s", message, mess) )
   end
   
end

function TransUnit:checkMatchValType( pos, funcTypeInfo, expList, genericTypeList )

   local argTypeList = funcTypeInfo:get_argTypeInfoList()
   do
      local _switchExp = funcTypeInfo
      if _switchExp == _moduleObj.typeInfoListInsert then
         argTypeList = genericTypeList
      elseif _switchExp == _moduleObj.typeInfoListRemove then
      end
   end
   
   local expNodeList = {}
   if expList ~= nil then
      for __index, node in pairs( expList:get_expList() ) do
         table.insert( expNodeList, node )
      end
      
   end
   
   self:checkMatchType( funcTypeInfo:getTxt(  ), pos, argTypeList, expNodeList, false )
end

local MacroPaser = {}
setmetatable( MacroPaser, { __index = Parser.Parser } )
function MacroPaser.new( tokenList, name )
   local obj = {}
   MacroPaser.setmeta( obj )
   if obj.__init then obj:__init( tokenList, name ); end
   return obj
end
function MacroPaser:__init(tokenList, name) 
   Parser.Parser.__init( self )
   
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
function MacroPaser.setmeta( obj )
  setmetatable( obj, { __index = MacroPaser  } )
end

function TransUnit:evalMacro( firstToken, macroTypeInfo, expList )

   if expList ~= nil then
      for __index, exp in pairs( expList:get_expList(  ) ) do
         local kind = exp:get_kind()
         if kind ~= Ast.NodeKind.get_LiteralNil() and kind ~= Ast.NodeKind.get_LiteralChar() and kind ~= Ast.NodeKind.get_LiteralInt() and kind ~= Ast.NodeKind.get_LiteralReal() and kind ~= Ast.NodeKind.get_LiteralArray() and kind ~= Ast.NodeKind.get_LiteralList() and kind ~= Ast.NodeKind.get_LiteralMap() and kind ~= Ast.NodeKind.get_LiteralString() and kind ~= Ast.NodeKind.get_LiteralBool() and kind ~= Ast.NodeKind.get_LiteralSymbol() and kind ~= Ast.NodeKind.get_RefField() and kind ~= Ast.NodeKind.get_ExpMacroStat() then
            self:error( "Macro arguments must be literal value." )
         end
         
      end
      
   end
   
   local macroInfo = _lune.unwrap( self.typeId2MacroInfo[macroTypeInfo:get_typeId(  )])
   local argValMap = {}
   local macroArgValMap = {}
   local macroArgNodeList = macroInfo.declInfo:get_argList()
   if expList ~= nil then
      for index, argNode in pairs( expList:get_expList(  ) ) do
         local valList, typeInfoList = argNode:getLiteral(  )
         local typeInfo = typeInfoList[1]
         local val = valList[1]
         if  nil == val then
            local _val = val
         
         else
            
               argValMap[index] = val
               local declArgNode = macroArgNodeList[index]
               macroArgValMap[declArgNode:get_name().txt] = val
         end
         
      end
      
   end
   
   local func = macroInfo.func
   local macroVars = {}
   do
      local macroRet = func( macroArgValMap )
      if macroRet ~= nil then
         macroVars = macroRet
      end
   end
   
   for __index, name in pairs( (_lune.unwrap( macroVars['_names']) ) ) do
      local valInfo = _lune.unwrap( macroInfo.symbol2MacroValInfoMap[name])
      local typeInfo = valInfo.typeInfo or Ast.builtinTypeStem_
      local val = macroVars[name]
      if typeInfo:equals( Ast.builtinTypeSymbol ) then
         val = {val}
      end
      
      self.symbol2ValueMapForMacro[name] = Ast.MacroValInfo.new(val, typeInfo)
   end
   
   for index, arg in pairs( macroInfo.declInfo:get_argList() ) do
      if arg:get_kind(  ) == Ast.NodeKind.get_DeclArg() then
         local argInfo = arg
         local argType = argInfo:get_argType()
         local argName = argInfo:get_name().txt
         self.symbol2ValueMapForMacro[argName] = Ast.MacroValInfo.new(argValMap[index], argType:get_expType())
      else
       
         self:error( "not support ... in macro" )
      end
      
   end
   
   local parser = MacroPaser.new(macroInfo.declInfo:get_tokenList(), string.format( "macro %s", macroTypeInfo:getTxt(  )))
   local bakParser = self.parser
   self.parser = parser
   self.macroMode = Ast.MacroMode.Expand
   local stmtList = {}
   self:analyzeStatementList( stmtList, "}" )
   self.macroMode = Ast.MacroMode.None
   self.parser = bakParser
   return Ast.ExpMacroExpNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, stmtList )
end

local function findForm( format )

   local remain = format
   local opList = {}
   while true do
      local pos, endPos = nil, nil
      do
         local index, endIndex = remain:find( "^%%[%d]*%a" )
         if index ~= nil and endIndex ~= nil then
            pos, endPos = index, endIndex
         else
            do
               local index, endIndex = remain:find( "[^%%]%%[%d]*%a" )
               if index ~= nil and endIndex ~= nil then
                  pos, endPos = index + 1, endIndex
               end
            end
            
         end
      end
      
      if pos ~= nil and endPos ~= nil then
         local op = remain:sub( pos, endPos )
         table.insert( opList, op )
         remain = remain:sub( endPos + 1 )
      else
         break
      end
      
   end
   
   return opList
end
_moduleObj.findForm = findForm
local FormType = {}
_moduleObj.FormType = FormType
FormType._val2NameMap = {}
function FormType:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "FormType.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end 
function FormType._from( val )
   if FormType._val2NameMap[ val ] then
      return val
   end
   return nil
end 
    
FormType.__allList = {}
function FormType.get__allList()
   return FormType.__allList
end

FormType.Match = 0
FormType._val2NameMap[0] = 'Match'
FormType.__allList[1] = FormType.Match
FormType.NeedConv = 1
FormType._val2NameMap[1] = 'NeedConv'
FormType.__allList[2] = FormType.NeedConv
FormType.Unmatch = 2
FormType._val2NameMap[2] = 'Unmatch'
FormType.__allList[3] = FormType.Unmatch

local function isMatchStringFormatType( opKind, argType, luaVer )

   if argType:get_kind() == Ast.TypeInfoKind.Enum then
      local enumType = argType:get_srcTypeInfo()
      argType = enumType:get_valTypeInfo()
   end
   
   do
      local _switchExp = string.byte( opKind, #opKind )
      if _switchExp == 115 or _switchExp == 113 then
         if not argType:equals( Ast.builtinTypeString ) then
            if not luaVer:get_canFormStem2Str() then
               return FormType.NeedConv, Ast.builtinTypeString
            end
            
         end
         
      elseif _switchExp == 65 or _switchExp == 97 or _switchExp == 69 or _switchExp == 101 or _switchExp == 102 or _switchExp == 71 or _switchExp == 103 then
         if not argType:equals( Ast.builtinTypeReal ) then
            return FormType.Unmatch, Ast.builtinTypeReal
         end
         
      else 
         
            if not argType:equals( Ast.builtinTypeInt ) then
               return FormType.Unmatch, Ast.builtinTypeInt
            end
            
      end
   end
   
   return FormType.Match, Ast.builtinTypeNone
end
_moduleObj.isMatchStringFormatType = isMatchStringFormatType
function TransUnit:checkStringFormat( pos, formatTxt, argTypeList )

   local opList = findForm( formatTxt )
   local dstTypeList = {}
   if #opList > #argTypeList then
      self:addErrMess( pos, string.format( "argument number is mismatch -- %d > %d", #opList, #argTypeList) )
      return 
   end
   
   for index, op in pairs( opList ) do
      local argType = argTypeList[index]
      local match, reqType = isMatchStringFormatType( op, argType, self.targetLuaVer )
      if match == FormType.Unmatch then
         local mess = string.format( "type must be %s -- %s", reqType:getTxt(  ), argType:getTxt(  ))
         self:addErrMess( pos, string.format( "argument(%d) %s", index, mess) )
      end
      
   end
   
end

function TransUnit:analyzeExpCall( firstToken, exp, nextToken )

   local macroFlag = false
   local funcTypeInfo = exp:get_expType()
   local nilAccess = nextToken.txt == "$("
   if nilAccess then
      if funcTypeInfo:get_nilable() then
         funcTypeInfo = funcTypeInfo:get_orgTypeInfo()
      else
       
         nilAccess = false
      end
      
   end
   
   if funcTypeInfo:get_kind(  ) == Ast.TypeInfoKind.Macro then
      macroFlag = true
      self.symbol2ValueMapForMacro = {}
      self.macroMode = Ast.MacroMode.Analyze
   end
   
   local work = self:getToken(  )
   local argList = nil
   if work.txt ~= ")" then
      self:pushback(  )
      argList = self:analyzeExpList( false, nil, funcTypeInfo:get_argTypeInfoList() )
      self:checkNextToken( ")" )
   end
   
   local genericTypeList = funcTypeInfo:get_itemTypeInfoList()
   if funcTypeInfo:get_kind() == Ast.TypeInfoKind.Method and exp:get_kind() == Ast.NodeKind.get_RefField() then
      local refField = exp
      local classType = refField:get_prefix():get_expType()
      genericTypeList = classType:get_itemTypeInfoList()
   end
   
   self:checkMatchValType( exp:get_pos(), funcTypeInfo, argList, genericTypeList )
   if funcTypeInfo:equals( _moduleObj.typeInfoListInsert ) then
      if argList ~= nil then
         if argList:get_expType():get_nilable() then
            self:addErrMess( argList:get_pos(), "list can't insert nilable" )
         end
         
      end
      
   end
   
   if macroFlag then
      self.macroMode = Ast.MacroMode.None
      exp = self:evalMacro( firstToken, funcTypeInfo, argList )
   else
    
      do
         local _switchExp = (funcTypeInfo:get_kind() )
         if _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.Func then
         else 
            
               self:error( string.format( "can't call the type -- %s, %s", funcTypeInfo:getTxt(  ), Ast.TypeInfoKind:_getTxt( funcTypeInfo:get_kind())
               ) )
         end
      end
      
      local retTypeInfoList = funcTypeInfo:get_retTypeInfoList(  )
      if nilAccess then
         local retList = {}
         for __index, retType in pairs( funcTypeInfo:get_retTypeInfoList(  ) ) do
            if retType:get_nilable() then
               table.insert( retList, retType )
            else
             
               table.insert( retList, retType:get_nilableTypeInfo() )
            end
            
         end
         
         retTypeInfoList = retList
         self.helperInfo.useNilAccess = true
      end
      
      local errorFuncFlag = false
      if #retTypeInfoList > 0 then
         local retType = retTypeInfoList[1]
         if retType == Ast.builtinTypeNeverRet then
            errorFuncFlag = true
         end
         
      end
      
      do
         local _switchExp = funcTypeInfo
         if _switchExp == typeInfoStringForm then
            if argList ~= nil then
               local formArgTypeList = {}
               local formatTxt = ""
               if #argList:get_expList() > 0 then
                  local argNode = argList:get_expList()[1]
                  if argNode:get_kind() == Ast.NodeKind.get_LiteralString() then
                     local workForm = argNode:getLiteral(  )
                     if #workForm >= 1 then
                        formatTxt = tostring( _lune.unwrapDefault( workForm[1], "") )
                     end
                     
                  end
                  
               end
               
               for index, argType in pairs( argList:get_expTypeList() ) do
                  if index ~= 1 then
                     table.insert( formArgTypeList, argType )
                  end
                  
               end
               
               self:checkStringFormat( firstToken.pos, formatTxt, formArgTypeList )
            end
            
         end
      end
      
      exp = Ast.ExpCallNode.create( self.nodeManager, firstToken.pos, retTypeInfoList, exp, errorFuncFlag, nilAccess, argList )
   end
   
   return exp, self:getToken(  )
end

function TransUnit:analyzeExpCont( firstToken, exp, skipFlag )

   local nextToken = self:getToken(  )
   if not skipFlag then
      repeat 
         local matchFlag = false
         if nextToken.txt == "[" or nextToken.txt == "$[" then
            matchFlag = true
            exp = self:analyzeExpRefItem( nextToken, exp, nextToken.txt == "$[" )
            nextToken = self:getToken(  )
         end
         
         if nextToken.txt == "(" or nextToken.txt == "$(" then
            matchFlag = true
            exp, nextToken = self:analyzeExpCall( firstToken, exp, nextToken )
         end
         
      until not matchFlag
   end
   
   do
      local _switchExp = nextToken.txt
      if _switchExp == "." then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.Field, exp, skipFlag )
      elseif _switchExp == "$." then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.FieldNil, exp, skipFlag )
      elseif _switchExp == ".$" then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.Get, exp, skipFlag )
      elseif _switchExp == "$.$" then
         return self:analyzeExpSymbol( firstToken, self:getToken(  ), ExpSymbolMode.GetNil, exp, skipFlag )
      end
   end
   
   self:pushback(  )
   return exp
end

function TransUnit:analyzeAccessClassField( classTypeInfo, mode, token )

   do
      local _switchExp = classTypeInfo:get_kind(  )
      if _switchExp == Ast.TypeInfoKind.List then
         classTypeInfo = Ast.builtinTypeList
      elseif _switchExp == Ast.TypeInfoKind.Array then
         classTypeInfo = Ast.builtinTypeArray
      end
   end
   
   local className = classTypeInfo:getTxt(  )
   local classScope = classTypeInfo:get_scope()
   if  nil == classScope then
      local _classScope = classScope
   
      self:error( string.format( "not found field: %s, %s, %s", classScope, className, classTypeInfo) )
   end
   
   local symbolInfo = nil
   local fieldTypeInfo = nil
   local getterFlag = false
   if mode == ExpSymbolMode.Get or mode == ExpSymbolMode.GetNil then
      local fieldSymbolInfo = classScope:getSymbolInfo( string.format( "get_%s", token.txt), self.scope, false )
      if fieldSymbolInfo ~= nil then
         if (fieldSymbolInfo:get_kind(  ) == Ast.SymbolKind.Mtd or fieldSymbolInfo:get_kind(  ) == Ast.SymbolKind.Fun ) then
            local retTypeList = fieldSymbolInfo:get_typeInfo():get_retTypeInfoList(  )
            symbolInfo = fieldSymbolInfo
            if #retTypeList > 0 then
               fieldTypeInfo = retTypeList[1]
            end
            
            getterFlag = true
         end
         
      end
      
   end
   
   if not symbolInfo then
      symbolInfo = classScope:getSymbolInfoField( token.txt, true, self.scope )
      if symbolInfo ~= nil then
         fieldTypeInfo = symbolInfo:get_typeInfo()
      end
      
   end
   
   if not fieldTypeInfo then
      for name, val in pairs( classScope:get_symbol2SymbolInfoMap() ) do
         Util.log( string.format( "debug: %s, %s", name, val) )
      end
      
      self:error( string.format( "not found field typeInfo: %s.%s", className, token.txt) )
   end
   
   local typeInfo = _lune.unwrapDefault( fieldTypeInfo, Ast.builtinTypeNone)
   return typeInfo, symbolInfo, getterFlag
end

function TransUnit:dumpComp( writer, pattern, symbolInfo, getterFlag )

   local symbol = symbolInfo:get_name()
   if pattern == "" or symbol:find( pattern ) then
      if getterFlag then
         writer:startParent( "candidate", false )
         local typeInfo = symbolInfo:get_typeInfo()
         writer:write( "type", string.format( "%s", Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
         ) )
         do
            local _switchExp = (symbolInfo:get_kind() )
            if _switchExp == Ast.SymbolKind.Mtd or _switchExp == Ast.SymbolKind.Fun then
               writer:write( "displayTxt", string.format( "$%s", typeInfo:get_rawTxt():gsub( "^get_", "" )) )
            elseif _switchExp == Ast.SymbolKind.Mbr then
               writer:write( "displayTxt", string.format( "$%s: %s", symbolInfo:get_name(), typeInfo:getTxt(  )) )
            end
         end
         
      else
       
         writer:startParent( "candidate", false )
         local typeInfo = symbolInfo:get_typeInfo()
         writer:write( "type", string.format( "%s", Ast.SymbolKind:_getTxt( symbolInfo:get_kind())
         ) )
         do
            local _switchExp = (symbolInfo:get_kind() )
            if _switchExp == Ast.SymbolKind.Fun or _switchExp == Ast.SymbolKind.Mtd then
               writer:write( "displayTxt", string.format( "%s", typeInfo:get_display_stirng()) )
            elseif _switchExp == Ast.SymbolKind.Mbr or _switchExp == Ast.SymbolKind.Var then
               writer:write( "displayTxt", string.format( "%s: %s", symbolInfo:get_name(), typeInfo:get_display_stirng()) )
            elseif _switchExp == Ast.SymbolKind.Typ then
               writer:write( "displayTxt", string.format( "%s", typeInfo:get_display_stirng()) )
            end
         end
         
      end
      
      writer:endElement(  )
   end
   
   return true
end

function TransUnit:dumpFieldComp( writer, isPrefixType, prefixTypeInfo, pattern, getterPattern )

   local typeInfo = prefixTypeInfo
   local scope = typeInfo:get_scope()
   if  nil == scope then
      local _scope = scope
   
      return 
   end
   
   scope:filterTypeInfoField( true, self.scope, function ( symbolInfo )
   
      if (isPrefixType ) then
         if not symbolInfo:get_staticFlag() and not symbolInfo:get_typeInfo():get_staticFlag() and symbolInfo:get_kind() ~= Ast.SymbolKind.Typ then
            return true
         end
         
      elseif symbolInfo:get_staticFlag() then
         return true
      end
      
      local symbol = symbolInfo:get_name()
      if symbol ~= "__init" and symbol ~= "__free" and symbol ~= "self" then
         if getterPattern ~= nil then
            if symbolInfo:get_kind() == Ast.SymbolKind.Mtd or symbolInfo:get_kind() == Ast.SymbolKind.Fun then
               local typeInfo = symbolInfo:get_typeInfo()
               local retList = typeInfo:get_retTypeInfoList()
               if #retList == 1 then
                  return self:dumpComp( writer, getterPattern, symbolInfo, true )
               end
               
            end
            
            return true
         end
         
         return self:dumpComp( writer, pattern, symbolInfo, false )
      end
      
      return true
   end
    )
end

function TransUnit:dumpSymbolComp( writer, scope, pattern )

   scope:filterSymbolTypeInfo( scope, self.moduleScope, function ( symbolInfo )
   
      return self:dumpComp( writer, pattern, symbolInfo, false )
   end
    )
end


function TransUnit:checkComp( token, callback )

   if self.analyzeMode == AnalyzeMode.Complete and self.analyzePos.lineNo == token.pos.lineNo and self.analyzePos.column >= token.pos.column and self.analyzePos.column <= token.pos.column + #token.txt then
      local currentModule = self.parser:getStreamName(  ):gsub( "%.lns", "" )
      currentModule = currentModule:gsub( ".*/", "" )
      local target = self.analyzeModule:gsub( "[^%.]+%.", "" )
      if currentModule == target then
         local jsonWriter = Writer.JSON.new(io.stdout)
         jsonWriter:startParent( "lunescript", false )
         local prefix = token.txt:gsub( "lune$", "" )
         jsonWriter:write( "prefix", prefix )
         jsonWriter:startParent( "candidateList", true )
         callback( jsonWriter, prefix )
         jsonWriter:endElement(  )
         jsonWriter:endElement(  )
         jsonWriter:fin(  )
         os.exit( 0 )
      end
      
   end
   
end

function TransUnit:checkFieldComp( getterFlag, token, prefixExp )

   if self.analyzeMode ~= AnalyzeMode.Complete then
      return 
   end
   
   local prefixSymbolInfoList = prefixExp:getSymbolInfo(  )
   local prefixSymbolInfo = nil
   if #prefixSymbolInfoList == 1 then
      prefixSymbolInfo = prefixSymbolInfoList[1]
   end
   
   self:checkComp( token, function ( jsonWriter, prefix )
   
      local getterPattern = nil
      if getterFlag then
         getterPattern = "^get_" .. prefix
      end
      
      local isPrefixType = false
      do
         local _exp = prefixSymbolInfo
         if _exp ~= nil then
            isPrefixType = _exp:get_kind() == Ast.SymbolKind.Typ
         end
      end
      
      self:dumpFieldComp( jsonWriter, isPrefixType, prefixExp:get_expType(), prefix == "" and "" or "^" .. prefix, getterPattern )
   end
    )
end

function TransUnit:checkEnumComp( token, enumTypeInfo )

   if self.analyzeMode ~= AnalyzeMode.Complete then
      return 
   end
   
   self:checkComp( token, function ( jsonWriter, prefix )
   
      self:dumpFieldComp( jsonWriter, true, enumTypeInfo, prefix == "" and "" or "^" .. prefix, nil )
   end
    )
end

function TransUnit:checkSymbolComp( token )

   self:checkComp( token, function ( jsonWriter, prefix )
   
      self:dumpSymbolComp( jsonWriter, self.scope, prefix == "" and "" or "^" .. prefix )
   end
    )
end

function TransUnit:analyzeExpField( firstToken, token, mode, prefixExp )

   local accessNil = false
   if mode == ExpSymbolMode.FieldNil or mode == ExpSymbolMode.GetNil then
      accessNil = true
   end
   
   if self.macroMode == Ast.MacroMode.Analyze then
      if accessNil then
         self.helperInfo.useNilAccess = true
      end
      
      return Ast.RefFieldNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeSymbol}, token, nil, accessNil, _lune.unwrap( prefixExp) )
   end
   
   local typeInfo = Ast.builtinTypeStem_
   local prefixExpType = prefixExp:get_expType()
   self:checkFieldComp( mode == ExpSymbolMode.Get or mode == ExpSymbolMode.GetNil, token, prefixExp )
   if accessNil then
      if prefixExpType:get_nilable() then
         prefixExpType = _lune.unwrap( prefixExpType:get_orgTypeInfo())
      else
       
         accessNil = false
      end
      
   end
   
   if accessNil then
      self.helperInfo.useNilAccess = true
   end
   
   local getterTypeInfo = nil
   local symbolInfo = nil
   if prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Class or prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Module or prefixExpType:get_kind(  ) == Ast.TypeInfoKind.IF or prefixExpType:get_kind(  ) == Ast.TypeInfoKind.List or prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Array then
      local getterFlag = false
      typeInfo, symbolInfo, getterFlag = self:analyzeAccessClassField( prefixExpType, mode, token )
      if getterFlag then
         do
            local _exp = symbolInfo
            if _exp ~= nil then
               getterTypeInfo = _exp:get_typeInfo()
            end
         end
         
      end
      
   elseif prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Enum or prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Alge then
      local scope = _lune.unwrap( prefixExpType:get_scope())
      local fieldName = token.txt
      if mode == ExpSymbolMode.Get then
         local moduleType = prefixExpType:getModule(  )
         if not moduleType:equals( self.moduleType ) and not self.importModule2ModuleInfoCurrent[moduleType] then
            self:addErrMess( token.pos, string.format( "need to import module -- %s", prefixExpType:getModule(  ):getTxt(  )) )
         end
         
         fieldName = "get_" .. fieldName
         do
            local funcType = scope:getTypeInfoChild( fieldName )
            if funcType ~= nil then
               local retTypeList = funcType:get_retTypeInfoList()
               if #retTypeList == 0 then
                  self:addErrMess( token.pos, string.format( "The func (%s) doesn't return value.", funcType:getTxt(  )) )
               else
                
                  typeInfo = retTypeList[1]
               end
               
            else
               self:addErrMess( token.pos, string.format( "not found -- %s.", fieldName) )
               typeInfo = Ast.builtinTypeNone
            end
         end
         
         getterTypeInfo = Ast.headTypeInfo
      else
       
         do
            local _exp = scope:getTypeInfoChild( fieldName )
            if _exp ~= nil then
               typeInfo = _exp
            else
               self:addErrMess( token.pos, string.format( "not found enum field -- %s", token.txt) )
               typeInfo = Ast.builtinTypeInt
            end
         end
         
      end
      
   elseif prefixExpType:get_kind(  ) == Ast.TypeInfoKind.Map then
      local work = prefixExpType:get_itemTypeInfoList()[1]
      if not work:equals( Ast.builtinTypeString ) then
         self:addErrMess( token.pos, string.format( "map key type is not str. (%s)", work:getTxt(  )) )
      end
      
      typeInfo = prefixExpType:get_itemTypeInfoList()[2]
      if not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
      return Ast.ExpRefItemNode.create( self.nodeManager, token.pos, {typeInfo}, prefixExp, accessNil, token.txt, nil )
   elseif prefixExpType:equals( Ast.builtinTypeStem ) then
      return Ast.ExpRefItemNode.create( self.nodeManager, token.pos, {Ast.builtinTypeStem_}, prefixExp, accessNil, token.txt, nil )
   else
    
      self:error( string.format( "illegal type -- %s, %d", prefixExpType:getTxt(  ), prefixExpType:get_kind(  )) )
   end
   
   if not symbolInfo then
      local prefixScope = prefixExpType:get_scope()
      do
         local _exp = prefixScope
         if _exp ~= nil then
            symbolInfo = _exp:getSymbolInfoField( token.txt, true, self.scope )
         end
      end
      
   end
   
   local prefixSymbolInfoList = prefixExp:getSymbolInfo(  )
   do
      local _exp = symbolInfo
      if _exp ~= nil then
         if #prefixSymbolInfoList == 1 then
            local prefixSymbolInfo = prefixSymbolInfoList[1]
            if prefixSymbolInfo:get_kind() == Ast.SymbolKind.Typ then
               if not _exp:get_staticFlag() and _exp:get_kind() ~= Ast.SymbolKind.Typ then
                  self:addErrMess( token.pos, string.format( "Type can't access this symbol. -- %s", _exp:get_name()) )
               end
               
            elseif _exp:get_staticFlag() and _exp:get_typeInfo():get_kind() ~= Ast.TypeInfoKind.Method then
               self:addErrMess( token.pos, string.format( "can't access this symbol. -- %s", token.txt) )
            end
            
         end
         
         
         if not prefixExpType:get_mutable() and not _exp:get_staticFlag() and _exp:get_kind() == Ast.SymbolKind.Mtd and _exp:get_mutable() then
            self:addErrMess( token.pos, string.format( "can't access mutable method. -- %s", token.txt) )
         end
         
      end
   end
   
   if accessNil then
      if not typeInfo:get_nilable() then
         typeInfo = typeInfo:get_nilableTypeInfo()
      end
      
      self.helperInfo.useNilAccess = true
   end
   
   local accessSymbolInfo = nil
   do
      local _exp = symbolInfo
      if _exp ~= nil then
         accessSymbolInfo = Ast.AccessSymbolInfo.new(_exp, prefixExpType, not accessNil)
      end
   end
   
   
   if not prefixExpType:get_mutable() and typeInfo:get_mutable() then
      typeInfo = self:createModifier( typeInfo, false )
   end
   
   if typeInfo:equals( _moduleObj.typeInfoListUnpack ) or typeInfo:equals( _moduleObj.typeInfoArrayUnpack ) then
      self.helperInfo.useUnpack = true
   end
   
   do
      local _exp = getterTypeInfo
      if _exp ~= nil then
         return Ast.GetFieldNode.create( self.nodeManager, firstToken.pos, {typeInfo}, token, accessSymbolInfo, accessNil, prefixExp, _exp )
      else
         return Ast.RefFieldNode.create( self.nodeManager, firstToken.pos, {typeInfo}, token, accessSymbolInfo, accessNil, prefixExp )
      end
   end
   
end

function TransUnit:analyzeNewAlge( firstToken, algeTypeInfo, prefix )

   local symbolToken = self:getSymbolToken(  )
   do
      local valInfo = algeTypeInfo:getValInfo( symbolToken.txt )
      if valInfo ~= nil then
         local argList = {}
         if #valInfo:get_typeList() > 0 then
            self:checkNextToken( "(" )
            for index, typeInfo in pairs( valInfo:get_typeList() ) do
               local argExp = self:analyzeExp( false )
               table.insert( argList, argExp )
               if index ~= #valInfo:get_typeList() then
                  self:checkNextToken( "," )
               end
               
            end
            
            self:checkNextToken( ")" )
         end
         
         self:checkMatchType( "call", symbolToken.pos, valInfo:get_typeList(), argList, false )
         return Ast.NewAlgeValNode.create( self.nodeManager, firstToken.pos, {algeTypeInfo}, symbolToken, prefix, algeTypeInfo, valInfo, argList )
      else
         self:addErrMess( symbolToken.pos, string.format( "not found Alge -- %s", symbolToken.txt) )
         return Ast.NewAlgeValNode.create( self.nodeManager, firstToken.pos, {algeTypeInfo}, symbolToken, prefix, algeTypeInfo, Ast.AlgeValInfo.new("", {}), {} )
      end
   end
   
end

function TransUnit:analyzeExpSymbol( firstToken, token, mode, prefixExp, skipFlag )

   local exp = nil
   if mode == ExpSymbolMode.Field or mode == ExpSymbolMode.Get or mode == ExpSymbolMode.FieldNil or mode == ExpSymbolMode.GetNil then
      if prefixExp ~= nil then
         exp = self:analyzeExpField( firstToken, token, mode, prefixExp )
         do
            local expType = exp:get_expType()
            if expType ~= nil then
               if expType:get_kind() == Ast.TypeInfoKind.Alge and prefixExp:get_expType():isModule(  ) then
                  local nextToken = self:getToken(  )
                  if nextToken.txt == "." then
                     return self:analyzeNewAlge( firstToken, expType, exp )
                  end
                  
                  self:pushback(  )
               end
               
            end
         end
         
      end
      
   elseif mode == ExpSymbolMode.Symbol then
      if self.macroMode == Ast.MacroMode.Analyze then
         exp = Ast.LiteralSymbolNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeSymbol}, token )
      else
       
         self:checkSymbolComp( token )
         local symbolInfo = self.scope:getSymbolTypeInfo( token.txt, self.scope, self.moduleScope )
         if  nil == symbolInfo then
            local _symbolInfo = symbolInfo
         
            local work = self.scope
            while true do
               print( work, self.globalScope, Ast.rootScope )
               if work == work:get_parent() then
                  break
               end
               
               work = work:get_parent()
            end
            
            self.scope:filterSymbolTypeInfo( self.scope, self.moduleScope, function ( workSymbolInfo )
            
               print( "sym", workSymbolInfo:get_name() )
               return true
            end
             )
            self:error( "not found type -- " .. token.txt )
         end
         
         local typeInfo = symbolInfo:get_typeInfo()
         if typeInfo:get_kind() == Ast.TypeInfoKind.Alge and symbolInfo:get_kind() == Ast.SymbolKind.Typ then
            local nextToken = self:getToken(  )
            if nextToken.txt == "." then
               return self:analyzeNewAlge( firstToken, typeInfo, nil )
            end
            
            self:pushback(  )
         end
         
         if typeInfo:equals( Ast.builtinTypeSymbol ) then
            skipFlag = true
         end
         
         if typeInfo:equals( _moduleObj.typeInfoLuneLoad ) then
            self.helperInfo.useLoad = true
         end
         
         if token.txt == "__func__" then
            self.has__func__Symbol = true
         end
         
         exp = Ast.ExpRefNode.create( self.nodeManager, firstToken.pos, {typeInfo}, token, Ast.AccessSymbolInfo.new(symbolInfo, nil, true) )
      end
      
   elseif mode == ExpSymbolMode.Fn then
      exp = self:analyzeDeclFunc( DeclFuncMode.Func, false, false, Ast.AccessMode.Local, false, nil, token, nil )
   else
    
      self:error( string.format( "illegal mode -- %s", mode) )
   end
   
   return self:analyzeExpCont( firstToken, _lune.unwrap( exp), skipFlag )
end

function TransUnit:analyzeExpOpSet( exp, opeToken, exp2NodeList )

   if not exp:canBeLeft(  ) then
      self:addErrMess( exp:get_pos(), string.format( "this node can not be l-value. -- %s", Ast.getNodeKindName( exp:get_kind() )) )
   end
   
   self:checkMatchType( "= operator", opeToken.pos, exp:get_expTypeList(), exp2NodeList, true )
   for __index, symbolInfo in pairs( exp:getSymbolInfo(  ) ) do
      if not symbolInfo:get_mutable() and symbolInfo:get_hasValueFlag() then
         if self.validMutControl then
            self:addErrMess( opeToken.pos, string.format( "this is not mutable variable. -- %s", symbolInfo:get_name()) )
         end
         
      end
      
      symbolInfo:set_hasValueFlag( true )
   end
   
end

function TransUnit:analyzeExpOp2( firstToken, exp, prevOpLevel )

   while true do
      local nextToken = self:getToken(  )
      local opTxt = nextToken.txt
      if opTxt == "@@" then
         local castTypeNode = self:analyzeRefType( Ast.AccessMode.Local, false )
         local castType = castTypeNode:get_expType()
         local expType = exp:get_expType()
         if expType:get_nilable() and not castType:get_nilable() then
            self:addErrMess( firstToken.pos, string.format( "can't cast from nilable to not nilable  -- %s->%s", expType:getTxt(  ), castType:getTxt(  )) )
         elseif not expType:get_mutable() and castType:get_mutable() then
            castType = self:createModifier( castType, false )
         end
         
         if castType:canEvalWith( expType, "=" ) then
            self:addWarnMess( castTypeNode:get_pos(), string.format( "This cast doesn't need. (%s <- %s)", castType:getTxt( true ), expType:getTxt( true )) )
         elseif not expType:canEvalWith( castType, "=" ) then
            if not Ast.isNumberType( expType ) and not Ast.isNumberType( castType ) then
               self:addErrMess( castTypeNode:get_pos(), string.format( "This type can't cast. (%s <- %s)", castType:getTxt( true ), expType:getTxt( true )) )
            end
            
         end
         
         exp = Ast.ExpCastNode.create( self.nodeManager, firstToken.pos, {castType}, exp )
      elseif nextToken.kind == Parser.TokenKind.Ope then
         if Parser.isOp2( opTxt ) then
            if not exp:canBeRight(  ) then
               self:addErrMess( exp:get_pos(), string.format( "This can't evaluate for '%s' -- %s", opTxt, Ast.getNodeKindName( exp:get_kind() )) )
            end
            
            local opLevel = op2levelMap[opTxt]
            if  nil == opLevel then
               local _opLevel = opLevel
            
               self:error( string.format( "unknown op -- %s %s", opTxt, prevOpLevel) )
            end
            
            do
               local _exp = prevOpLevel
               if _exp ~= nil then
                  if opLevel <= _exp then
                     self:pushback(  )
                     return exp
                  end
                  
               end
            end
            
            local enumTypeInfo = nil
            do
               local prefixExpType = exp:get_expType()
               if prefixExpType:get_nilable() then
                  prefixExpType = prefixExpType:get_orgTypeInfo()
               end
               
               if prefixExpType:get_kind() == Ast.TypeInfoKind.Enum then
                  enumTypeInfo = prefixExpType:get_srcTypeInfo()
               end
               
            end
            
            local exp2 = self:analyzeExp( false, opLevel, enumTypeInfo )
            local exp2NodeList = {exp2}
            if not exp2:canBeRight(  ) then
               self:addErrMess( exp2:get_pos(), string.format( "This can't evaluate for '%s' -- %s", opTxt, Ast.getNodeKindName( exp2:get_kind() )) )
            end
            
            if opTxt == "=" then
               local workToken = self:getToken(  )
               if workToken.txt == "," then
                  local expListNode = self:analyzeExpList( false, exp2 )
                  exp2 = expListNode
                  exp2NodeList = expListNode:get_expList()
               else
                
                  self:pushback(  )
               end
               
            end
            
            local info = {["op"] = nextToken, ["exp1"] = exp, ["exp2"] = exp2}
            local retType = Ast.builtinTypeNone
            if not exp2:canBeRight(  ) then
               self:addErrMess( exp2:get_pos(), string.format( "this node can not be r-value. -- %s", Ast.getNodeKindName( exp2:get_kind() )) )
            end
            
            local exp1Type = exp:get_expType()
            local exp2Type = exp2:get_expType()
            do
               local _switchExp = opTxt
               if _switchExp == "or" then
                  if exp1Type:equals( exp2Type ) then
                     retType = exp1Type
                  elseif exp1Type:canEvalWith( exp2Type, "=" ) then
                     retType = exp1Type
                  elseif exp2Type:canEvalWith( exp1Type, "=" ) then
                     retType = exp2Type
                  elseif exp2Type:equals( Ast.builtinTypeNil ) then
                     retType = exp1Type
                  elseif exp1Type:equals( Ast.builtinTypeNil ) then
                     retType = exp2Type
                  else
                   
                     if exp1Type:get_nilable() or exp2Type:get_nilable() then
                        retType = Ast.builtinTypeStem_
                     else
                      
                        retType = Ast.builtinTypeStem
                     end
                     
                  end
                  
               elseif _switchExp == "and" then
                  local workToken = self:getToken(  )
                  self:pushback(  )
                  if not exp1Type:equals( Ast.builtinTypeBool ) and not exp1Type:equals( Ast.builtinTypeStem ) and not exp1Type:get_nilable() then
                     self:addWarnMess( exp:get_pos(), "this value never be 'false'" )
                  elseif exp2:get_kind() == Ast.NodeKind.get_LiteralBool() then
                     local valList = exp2:getLiteral(  )
                     if not valList[1] then
                        self:addErrMess( exp2:get_pos(), "this value never be 'true'" )
                     end
                     
                  end
                  
                  if workToken.txt == "or" then
                     retType = exp2Type
                  else
                   
                     if exp1Type:get_nilable() then
                        if exp2Type:get_nilable() then
                           retType = exp2Type
                        else
                         
                           retType = exp2Type:get_nilableTypeInfo()
                        end
                        
                     elseif exp1Type:equals( Ast.builtinTypeBool ) or exp2Type:equals( Ast.builtinTypeBool ) then
                        if exp1Type:canEvalWith( exp2Type, "=" ) then
                           retType = exp1Type
                        elseif exp2Type:canEvalWith( exp1Type, "=" ) then
                           retType = exp2Type
                        else
                         
                           if exp2Type:get_nilable() then
                              retType = Ast.builtinTypeStem_
                           else
                            
                              retType = Ast.builtinTypeStem
                           end
                           
                        end
                        
                     else
                      
                        retType = exp2Type
                     end
                     
                  end
                  
               elseif _switchExp == "<" or _switchExp == ">" or _switchExp == "<=" or _switchExp == ">=" then
                  if Ast.builtinTypeString:canEvalWith( exp1Type, "=" ) and Ast.builtinTypeString:canEvalWith( exp2Type, "=" ) or (Ast.builtinTypeInt:canEvalWith( exp1Type, opTxt ) or Ast.builtinTypeReal:canEvalWith( exp1Type, opTxt ) ) and (Ast.builtinTypeInt:canEvalWith( exp2Type, opTxt ) or Ast.builtinTypeReal:canEvalWith( exp2Type, opTxt ) ) then
                     
                  else
                   
                     self:addErrMess( nextToken.pos, string.format( "no numeric type %s or %s", exp1Type:getTxt( true ), exp2Type:getTxt( true )) )
                  end
                  
                  retType = Ast.builtinTypeBool
               elseif _switchExp == "~=" or _switchExp == "==" then
                  if (not exp1Type:canEvalWith( exp2Type, opTxt ) and not exp2Type:canEvalWith( exp1Type, opTxt ) ) then
                     self:addErrMess( nextToken.pos, string.format( "not compatible type %s or %s", exp1Type:getTxt( true ), exp2Type:getTxt( true )) )
                  end
                  
                  if exp1Type:equals( Ast.builtinTypeBool ) and exp2Type:equals( Ast.builtinTypeBool ) and (exp:get_kind() == Ast.NodeKind.get_LiteralBool() or exp2:get_kind() == Ast.NodeKind.get_LiteralBool() ) then
                     self:addWarnMess( exp:get_pos(), "this operation is deprecated." )
                  end
                  
                  retType = Ast.builtinTypeBool
               elseif _switchExp == "^" or _switchExp == "|" or _switchExp == "~" or _switchExp == "&" or _switchExp == "|<<" or _switchExp == "|>>" then
                  if self.targetLuaVer:get_hasBitOp() == LuaVer.BitOp.Cant then
                     self:addErrMess( nextToken.pos, "this lua version can't use bit operand." )
                  end
                  
                  if not Ast.builtinTypeInt:canEvalWith( exp1Type, opTxt ) or not Ast.builtinTypeInt:canEvalWith( exp2Type, opTxt ) then
                     self:addErrMess( nextToken.pos, string.format( "no int type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  retType = Ast.builtinTypeInt
               elseif _switchExp == ".." then
                  if not exp1Type:equals( Ast.builtinTypeString ) or not exp1Type:equals( Ast.builtinTypeString ) then
                     self:addErrMess( nextToken.pos, string.format( "no string type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  retType = Ast.builtinTypeString
               elseif _switchExp == "+" or _switchExp == "-" or _switchExp == "*" or _switchExp == "/" or _switchExp == "%" then
                  if (not Ast.builtinTypeInt:canEvalWith( exp1Type, opTxt ) and not Ast.builtinTypeReal:canEvalWith( exp1Type, opTxt ) ) or (not Ast.builtinTypeInt:canEvalWith( exp2Type, opTxt ) and not Ast.builtinTypeReal:canEvalWith( exp2Type, opTxt ) ) then
                     self:addErrMess( nextToken.pos, string.format( "no numeric type %s or %s", exp1Type:getTxt(  ), exp2Type:getTxt(  )) )
                  end
                  
                  if exp1Type:equals( Ast.builtinTypeReal ) or exp2Type:equals( Ast.builtinTypeReal ) then
                     retType = Ast.builtinTypeReal
                  else
                   
                     retType = Ast.builtinTypeInt
                  end
                  
               elseif _switchExp == "=" then
                  self:analyzeExpOpSet( exp, nextToken, exp2NodeList )
               else 
                  
                     self:error( "unknown op " .. opTxt )
               end
            end
            
            exp = Ast.ExpOp2Node.create( self.nodeManager, firstToken.pos, {retType}, nextToken, exp, exp2 )
         else
          
            self:error( "illegal op" )
         end
         
      else
       
         self:pushback(  )
         return exp
      end
      
   end
   
end

function TransUnit:analyzeExpMacroStat( firstToken )

   local expStrList = {}
   self:checkNextToken( "{" )
   local braceCount = 0
   local prevToken = firstToken
   while true do
      local token = self:getToken(  )
      if token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
         local exp = self:analyzeExp( true, _lune.unwrap( op1levelMap[token.txt]) )
         local nextToken = self:getToken(  )
         if nextToken.txt ~= "~~" then
            self:pushback(  )
         end
         
         local format = token.txt == ",,," and "'%s '" or '"\'%s\'"'
         if token.txt == ",," and exp:get_kind() == Ast.NodeKind.get_ExpRef() then
            local refToken = (exp ):get_token(  )
            local macroInfo = self.symbol2ValueMapForMacro[refToken.txt]
            if macroInfo then
               if (_lune.unwrap( macroInfo) ).typeInfo:equals( Ast.builtinTypeSymbol ) then
                  format = "'%s '"
               end
               
            else
             
               if exp:get_expType():equals( Ast.builtinTypeInt ) or exp:get_expType():equals( Ast.builtinTypeReal ) then
                  format = "'%s' "
               end
               
            end
            
         end
         
         local newToken = Parser.Token.new(Parser.TokenKind.Str, format, token.pos)
         local literalStr = Ast.LiteralStringNode.create( self.nodeManager, token.pos, {Ast.builtinTypeString}, newToken, {exp} )
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
         
         local format = "' %s '"
         if prevToken == firstToken or (prevToken.pos.lineNo == token.pos.lineNo and prevToken.pos.column + #prevToken.txt == token.pos.column ) then
            format = "'%s'"
         end
         
         local newToken = Parser.Token.new(token.kind, string.format( format, token.txt ), token.pos)
         local literalStr = Ast.LiteralStringNode.create( self.nodeManager, token.pos, {Ast.builtinTypeString}, newToken, {} )
         table.insert( expStrList, literalStr )
      end
      
      prevToken = token
   end
   
   return Ast.ExpMacroStatNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeStat}, expStrList )
end

function TransUnit:analyzeSuper( firstToken )

   self:checkNextToken( "(" )
   local nextToken = self:getToken(  )
   local expList = nil
   if nextToken.txt ~= ")" then
      self:pushback(  )
      expList = self:analyzeExpList( false )
      self:checkNextToken( ")" )
   end
   
   self:checkNextToken( ";" )
   local classType = self:getCurrentClass(  )
   local currentFunc = self:getCurrentNamespaceTypeInfo(  )
   if currentFunc:get_kind() == Ast.TypeInfoKind.Method then
      local superType = classType:get_baseTypeInfo(  )
      if superType:equals( Ast.headTypeInfo ) then
         self:addErrMess( firstToken.pos, "This class doesn't have super-class." )
      else
       
         if currentFunc:get_rawTxt() == "__init" then
            local superScope = superType:get_scope()
            if  nil == superScope then
               local _superScope = superScope
            
               self:error( "not found super scope" )
            end
            
            local superCtorType = superScope:getTypeInfoChild( "__init" )
            if  nil == superCtorType then
               local _superCtorType = superCtorType
            
               self:error( "not found super '__init'" )
            end
            
            self:checkMatchValType( firstToken.pos, superCtorType, expList, {} )
            return Ast.ExpCallSuperNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, superType, superCtorType, expList )
         else
          
            do
               local superFunc = (_lune.unwrap( superType:get_scope()) ):getTypeInfoField( currentFunc:get_rawTxt(), true, self.scope )
               if superFunc ~= nil then
                  if superFunc:get_abstractFlag() then
                     self:addErrMess( firstToken.pos, "super is abstract." )
                  end
                  
                  self:checkMatchValType( firstToken.pos, superFunc, expList, {} )
                  return Ast.ExpCallSuperNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, superType, superFunc, expList )
               end
            end
            
            self:addErrMess( firstToken.pos, "this is not override method." )
            return self:createNoneNode( firstToken.pos )
         end
         
      end
      
   end
   
   self:addErrMess( firstToken.pos, "super can't call here." )
   return self:createNoneNode( firstToken.pos )
end

function TransUnit:analyzeUnwrap( firstToken )

   local nextToken, continueFlag = self:getContinueToken(  )
   if not continueFlag or nextToken.txt ~= "!" then
      self:pushback(  )
      self:pushbackToken( firstToken )
      local exp = self:analyzeExp( false )
      self:checkNextToken( ";" )
      if not exp:get_expType():get_nilable() then
         self:addErrMess( exp:get_pos(), "this value is not nilable." )
      end
      
      return Ast.StmtExpNode.create( self.nodeManager, nextToken.pos, {Ast.builtinTypeNone}, exp )
   end
   
   self:pushback(  )
   return self:analyzeDeclVar( Ast.DeclVarMode.Unwrap, Ast.AccessMode.Local, firstToken )
end

function TransUnit:analyzeExpUnwrap( firstToken )

   local expNode = self:analyzeExp( true )
   local nextToken = self:getToken(  )
   local insNode = nil
   if nextToken.txt == "default" then
      insNode = self:analyzeExp( false )
   else
    
      self:pushback(  )
   end
   
   local unwrapType = Ast.builtinTypeStem_
   local expType = expNode:get_expType()
   if not expType:get_nilable() then
      unwrapType = expType
   elseif expType:equals( Ast.builtinTypeDDD ) then
      unwrapType = Ast.builtinTypeStem
   else
    
      unwrapType = _lune.unwrap( expType:get_orgTypeInfo())
      do
         local _exp = insNode
         if _exp ~= nil then
            local insType = _exp:get_expType()
            unwrapType = insType
            if not unwrapType:canEvalWith( insType, "=" ) then
               self:addErrMess( _exp:get_pos(), string.format( "unmatch type: %s <- %s", unwrapType:getTxt( true ), insType:getTxt( true )) )
            end
            
         end
      end
      
   end
   
   self.helperInfo.useUnwrapExp = true
   return Ast.ExpUnwrapNode.create( self.nodeManager, firstToken.pos, {unwrapType}, expNode, insNode )
end

function TransUnit:analyzeExp( skipOp2Flag, prevOpLevel, expectType )

   local firstToken = self:getToken(  )
   local token = firstToken
   local exp = self:createNoneNode( firstToken.pos )
   if token.kind == Parser.TokenKind.Dlmt then
      if token.txt == "." then
         do
            local _exp = expectType
            if _exp ~= nil then
               local orgExpectType = _exp
               if orgExpectType:get_nilable() then
                  orgExpectType = orgExpectType:get_orgTypeInfo()
               end
               
               if orgExpectType:get_kind() == Ast.TypeInfoKind.Enum then
                  local enumTyepInfo = orgExpectType:get_srcTypeInfo()
                  local nextToken = self:getToken(  )
                  self:checkEnumComp( nextToken, enumTyepInfo )
                  if enumTyepInfo:getEnumValInfo( nextToken.txt ) then
                     if orgExpectType:get_externalFlag() and not self.importModule2ModuleInfoCurrent[orgExpectType:getModule(  ):get_srcTypeInfo()] then
                        local fullname = orgExpectType:getFullName( self.importModule2ModuleInfo, true )
                        self:addErrMess( token.pos, string.format( "This module not import -- %s", fullname) )
                     end
                     
                     exp = Ast.ExpOmitEnumNode.create( self.nodeManager, token.pos, {enumTyepInfo}, nextToken, enumTyepInfo )
                     exp = self:analyzeExpCont( firstToken, exp, false )
                  else
                   
                     self:error( string.format( "illegal enum val -- %s.%s", orgExpectType:getTxt(  ), nextToken.txt) )
                  end
                  
               else
                
                  self:error( string.format( "illegal type for '.' -- %s", orgExpectType:getTxt(  )) )
               end
               
            else
               self:error( "illegal '.'" )
            end
         end
         
      elseif token.txt == "..." then
         return Ast.ExpDDDNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, token )
      elseif token.txt == '[' or token.txt == '[@' then
         exp = self:analyzeListConst( token )
      elseif token.txt == '{' then
         exp = self:analyzeMapConst( token )
      elseif token.txt == "(" then
         exp = self:analyzeExp( false )
         self:checkNextToken( ")" )
         exp = Ast.ExpParenNode.create( self.nodeManager, firstToken.pos, {exp:get_expType()}, exp )
         exp = self:analyzeExpCont( firstToken, exp, false )
      end
      
   end
   
   if token.txt == "new" then
      exp = self:analyzeRefType( Ast.AccessMode.Local, false )
      local classTypeInfo = exp:get_expType()
      if classTypeInfo:get_externalFlag() then
         do
            local _switchExp = classTypeInfo:get_accessMode()
            if _switchExp == Ast.AccessMode.Pri or _switchExp == Ast.AccessMode.Local then
               self:addErrMess( token.pos, string.format( "Can't access -- %s", Ast.AccessMode:_getTxt( classTypeInfo:get_accessMode())
               ) )
            end
         end
         
      end
      
      if classTypeInfo:get_abstractFlag() then
         self:addErrMess( token.pos, "abstract class can't new" )
      end
      
      local classScope = classTypeInfo:get_scope(  )
      local initTypeInfo = (_lune.unwrap( classScope) ):getTypeInfoChild( "__init" )
      if  nil == initTypeInfo then
         local _initTypeInfo = initTypeInfo
      
         self:error( "not found __init" )
      end
      
      self:checkNextToken( "(" )
      local nextToken = self:getToken(  )
      local argList = nil
      if nextToken.txt ~= ")" then
         self:pushback(  )
         argList = self:analyzeExpList( false, nil, initTypeInfo:get_argTypeInfoList() )
         self:checkNextToken( ")" )
      end
      
      if initTypeInfo:get_accessMode() == Ast.AccessMode.Pub or (initTypeInfo:get_accessMode() == Ast.AccessMode.Pro and self.scope:getClassTypeInfo(  ):isInheritFrom( classTypeInfo ) ) or (self.scope:getClassTypeInfo(  ) == classTypeInfo ) then
      else
       
         self:addErrMess( token.pos, string.format( "can't access to __init of %s", classTypeInfo:getTxt(  )) )
      end
      
      self:checkMatchValType( exp:get_pos(), initTypeInfo, argList, exp:get_expType():get_itemTypeInfoList() )
      exp = Ast.ExpNewNode.create( self.nodeManager, firstToken.pos, exp:get_expTypeList(), exp, argList )
      exp = self:analyzeExpCont( firstToken, exp, false )
   end
   
   if token.kind == Parser.TokenKind.Ope and Parser.isOp1( token.txt ) then
      if token.txt == "`" then
         exp = self:analyzeExpMacroStat( token )
      else
       
         exp = self:analyzeExp( true, _lune.unwrap( op1levelMap[token.txt]) )
         local typeInfo = Ast.builtinTypeNone
         local macroExpFlag = false
         local expType = exp:get_expType()
         do
            local _switchExp = (token.txt )
            if _switchExp == "-" then
               if not expType:equals( Ast.builtinTypeInt ) and not expType:equals( Ast.builtinTypeReal ) then
                  self:addErrMess( token.pos, string.format( 'unmatch type for "-" -- %s', expType:getTxt(  )) )
               end
               
               typeInfo = expType
            elseif _switchExp == "#" then
               if expType:get_kind() ~= Ast.TypeInfoKind.List and expType:get_kind() ~= Ast.TypeInfoKind.Array and not Ast.builtinTypeString:canEvalWith( expType, "=" ) then
                  self:addErrMess( token.pos, string.format( 'unmatch type for "#" -- %s', expType:getTxt(  )) )
               end
               
               typeInfo = Ast.builtinTypeInt
            elseif _switchExp == "not" then
               typeInfo = Ast.builtinTypeBool
               if not expType:get_nilable() and not expType:equals( Ast.builtinTypeBool ) and not expType:equals( Ast.builtinTypeStem ) and not expType:equals( Ast.builtinTypeDDD ) then
                  self:addErrMess( token.pos, "this 'not' operand never be false" )
               end
               
            elseif _switchExp == ",," then
               macroExpFlag = true
            elseif _switchExp == ",,," then
               macroExpFlag = true
               if not expType:equals( Ast.builtinTypeString ) then
                  self:error( "unmatch ,,, type, need string type" )
               end
               
               typeInfo = Ast.builtinTypeSymbol
            elseif _switchExp == ",,,," then
               macroExpFlag = true
               if not expType:equals( Ast.builtinTypeSymbol ) then
                  self:error( "unmatch ,,, type, need symbol type" )
               end
               
               typeInfo = Ast.builtinTypeString
            elseif _switchExp == "`" then
               typeInfo = Ast.builtinTypeNone
            elseif _switchExp == "~" then
               if not expType:equals( Ast.builtinTypeInt ) then
                  self:addErrMess( token.pos, string.format( 'unmatch type for "~" -- %s', expType:getTxt(  )) )
               end
               
               typeInfo = Ast.builtinTypeInt
            else 
               
                  self:error( "unknown op1" )
            end
         end
         
         if macroExpFlag then
            local nextToken = self:getToken(  )
            if nextToken.txt ~= "~~" then
               self:pushback(  )
            end
            
         end
         
         exp = Ast.ExpOp1Node.create( self.nodeManager, firstToken.pos, {typeInfo}, token, self.macroMode, exp )
         return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
      end
      
   end
   
   if token.kind == Parser.TokenKind.Int then
      exp = Ast.LiteralIntNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeInt}, token, math.floor(tonumber( token.txt )) )
   elseif token.kind == Parser.TokenKind.Real then
      exp = Ast.LiteralRealNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeReal}, token, tonumber( token.txt ) )
   elseif token.kind == Parser.TokenKind.Char then
      local num = 0
      if #(token.txt ) == 1 then
         num = token.txt:byte( 1 )
      else
       
         num = _lune.unwrap( quotedChar2Code[token.txt:sub( 2, 2 )])
      end
      
      exp = Ast.LiteralCharNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeChar}, token, num )
   elseif token.kind == Parser.TokenKind.Str then
      local nextToken = self:getToken(  )
      local formatArgList = {}
      if nextToken.txt == "(" then
         local argNodeList = self:analyzeExpList( false, nil )
         self:checkNextToken( ")" )
         nextToken = self:getToken(  )
         formatArgList = argNodeList:get_expList()
         self:checkStringFormat( token.pos, token.txt, argNodeList:get_expTypeList() )
      end
      
      exp = Ast.LiteralStringNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeString}, token, formatArgList )
      token = nextToken
      if token.txt == "[" or token.txt == "$[" then
         exp = self:analyzeExpRefItem( token, exp, token.txt == "$[" )
      else
       
         self:pushback(  )
      end
      
   elseif token.kind == Parser.TokenKind.Symb and token.txt == "__line__" then
      exp = Ast.LiteralIntNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeInt}, Parser.Token.new(Parser.TokenKind.Int, string.format( "%d", token.pos.lineNo), token.pos, nil), token.pos.lineNo )
   elseif token.kind == Parser.TokenKind.Kywd and token.txt == "fn" then
      exp = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Fn, nil, false )
   elseif token.kind == Parser.TokenKind.Kywd and token.txt == "unwrap" then
      exp = self:analyzeExpUnwrap( token )
   elseif token.kind == Parser.TokenKind.Symb then
      exp = self:analyzeExpSymbol( firstToken, token, ExpSymbolMode.Symbol, nil, false )
   elseif token.kind == Parser.TokenKind.Type then
      local symbolTypeInfo = Ast.sym2builtInTypeMap[token.txt]
      if  nil == symbolTypeInfo then
         local _symbolTypeInfo = symbolTypeInfo
      
         self:error( string.format( "unknown type -- %s", token.txt) )
      end
      
      exp = Ast.ExpRefNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNone}, token, Ast.AccessSymbolInfo.new(symbolTypeInfo, nil, false) )
   elseif token.kind == Parser.TokenKind.Kywd and (token.txt == "true" or token.txt == "false" ) then
      exp = Ast.LiteralBoolNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeBool}, token )
   elseif token.kind == Parser.TokenKind.Kywd and (token.txt == "nil" or token.txt == "null" ) then
      exp = Ast.LiteralNilNode.create( self.nodeManager, firstToken.pos, {Ast.builtinTypeNil} )
   end
   
   if exp:get_kind() == Ast.NodeKind.get_None() then
      self:error( "illegal exp" )
   end
   
   if skipOp2Flag then
      return exp
   end
   
   return self:analyzeExpOp2( firstToken, exp, prevOpLevel )
end

function TransUnit:analyzeReturn( token )

   local expList = nil
   local funcTypeInfo = self:getCurrentNamespaceTypeInfo(  )
   if funcTypeInfo == Ast.headTypeInfo or (funcTypeInfo:get_kind() ~= Ast.TypeInfoKind.Func and funcTypeInfo:get_kind() ~= Ast.TypeInfoKind.Method ) then
      self:addErrMess( token.pos, "'return' could not use here" )
   end
   
   local nextToken = self:getToken(  )
   local retTypeList = funcTypeInfo:get_retTypeInfoList()
   if nextToken.txt ~= ";" then
      self:pushback(  )
      expList = self:analyzeExpList( false, nil, retTypeList )
      self:checkNextToken( ";" )
   end
   
   do
      local _exp = expList
      if _exp ~= nil then
         local expTypeList = _exp:get_expTypeList()
         if #retTypeList == 0 and #expTypeList > 0 then
            self:addErrMess( token.pos, "this function can't return value." )
         end
         
         for index, retType in pairs( retTypeList ) do
            local expType = expTypeList[index]
            if not retType:canEvalWith( expType, "=" ) then
               self:addErrMess( token.pos, string.format( "return type of arg(%d) is not compatible -- %s(%d) and %s(%d)", index, retType:getTxt(  ), retType:get_typeId(  ), expType:getTxt(  ), expType:get_typeId(  )) )
            end
            
            if index == #retTypeList then
               if #retTypeList < #expTypeList and not retType:equals( Ast.builtinTypeDDD ) then
                  self:addErrMess( token.pos, "over return value" )
               end
               
            elseif index == #expTypeList then
               if expType:equals( Ast.builtinTypeDDD ) then
                  for retIndex = index, #retTypeList do
                     local workRetType = retTypeList[retIndex]
                     if not workRetType:canEvalWith( Ast.builtinTypeStem_, "=" ) then
                        self:addErrMess( token.pos, string.format( "return type of arg(%d) is not compatible -- %s(%d) and %s(%d)", retIndex, workRetType:getTxt(  ), workRetType:get_typeId(  ), expType:getTxt(  ), expType:get_typeId(  )) )
                     end
                     
                  end
                  
               else
                
                  self:addErrMess( token.pos, string.format( "short return value -- %d < %d", #expTypeList, #retTypeList) )
               end
               
               break
            end
            
         end
         
      else
         if funcTypeInfo:getTxt(  ) == "__init" then
            self:addErrMess( token.pos, "__init method can't return" )
         end
         
         if #retTypeList ~= 0 then
            self:addErrMess( token.pos, "no return value" )
         end
         
      end
   end
   
   return Ast.ReturnNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone}, expList )
end

function TransUnit:analyzeStatement( termTxt )

   local token = self:getTokenNoErr(  )
   if token == Parser.getEofToken(  ) then
      return nil
   end
   
   
   local statement = self:analyzeDecl( Ast.AccessMode.Local, false, token, token )
   if not statement then
      if token.txt == termTxt then
         self:pushback(  )
         return nil
      elseif token.txt == "pub" or token.txt == "pro" or token.txt == "pri" or token.txt == "global" or token.txt == "static" then
         local accessMode = Ast.txt2AccessMode( token.txt )
         if  nil == accessMode then
            local _accessMode = accessMode
         
            accessMode = Ast.AccessMode.Pri
         end
         
         local staticFlag = (token.txt == "static" )
         local nextToken = token
         if token.txt ~= "static" then
            nextToken = self:getToken(  )
         end
         
         statement = self:analyzeDecl( accessMode, staticFlag, token, nextToken )
         if not statement then
            self:addErrMess( nextToken.pos, string.format( "This token is illegal -- %s", nextToken.txt) )
         end
         
      elseif token.txt == "{" then
         self:pushback(  )
         statement = self:analyzeBlock( Ast.BlockKind.Block )
      elseif token.txt == "super" then
         statement = self:analyzeSuper( token )
      elseif token.txt == "if" then
         statement = self:analyzeIf( token )
      elseif token.txt == "when" then
         statement = self:analyzeWhen( token )
      elseif token.txt == "switch" then
         statement = self:analyzeSwitch( token )
      elseif token.txt == "match" then
         statement = self:analyzeMatch( token )
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
         statement = self:analyzeReturn( token )
      elseif token.txt == "break" then
         self:checkNextToken( ";" )
         statement = Ast.BreakNode.create( self.nodeManager, token.pos, {Ast.builtinTypeNone} )
         if #self.loopScopeQueue == 0 then
            self:addErrMess( token.pos, "no loop syntax." )
         end
         
      elseif token.txt == "unwrap" then
         statement = self:analyzeUnwrap( token )
      elseif token.txt == "sync" then
         statement = self:analyzeDeclVar( Ast.DeclVarMode.Sync, Ast.AccessMode.Local, token )
      elseif token.txt == "import" then
         statement = self:analyzeImport( token )
      elseif token.txt == "subfile" then
         statement = self:analyzeSubfile( token )
      elseif token.txt == "_lune_control" then
         self:analyzeLuneControl( token )
         statement = self:createNoneNode( token.pos )
      elseif token.txt == "provide" then
         statement = self:analyzeProvide( token )
      elseif token.txt == ";" then
         statement = self:createNoneNode( token.pos )
      elseif token.txt == ",," or token.txt == ",,," or token.txt == ",,,," then
         self:error( string.format( "illegal macro op -- %s", token.txt) )
      else
       
         self:pushback(  )
         local exp = self:analyzeExp( false )
         local nextToken = self:getToken(  )
         if nextToken.txt == "," then
            exp = self:analyzeExpList( true, exp )
            exp = self:analyzeExpOp2( token, exp, nil )
            nextToken = self:getToken(  )
         end
         
         self:checkToken( nextToken, ";" )
         statement = Ast.StmtExpNode.create( self.nodeManager, exp:get_pos(), {Ast.builtinTypeNone}, exp )
      end
      
   end
   
   do
      local _exp = statement
      if _exp ~= nil then
         if not _exp:canBeStatement(  ) then
            self:addErrMess( _exp:get_pos(), string.format( "This node can't be statement. -- %s", Ast.getNodeKindName( _exp:get_kind() )) )
         end
         
      end
   end
   
   return statement
end

return _moduleObj
