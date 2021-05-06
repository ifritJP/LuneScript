--lune/base/Import.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Import'
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
local Types = _lune.loadModule( 'lune.base.Types' )
local Meta = _lune.loadModule( 'lune.base.Meta' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Macro = _lune.loadModule( 'lune.base.Macro' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local frontInterface = _lune.loadModule( 'lune.base.frontInterface' )
local Log = _lune.loadModule( 'lune.base.Log' )

local TransUnitIF = _lune.loadModule( 'lune.base.TransUnitIF' )
local Builtin = _lune.loadModule( 'lune.base.Builtin' )



local Import = {}
_moduleObj.Import = Import
function Import.new( transUnitIF, importModuleInfo, moduleType, builtinFunc, globalScope, macroCtrl, typeNameCtrl, helperInfo, importedAliasMap, validMutControl )
   local obj = {}
   Import.setmeta( obj )
   if obj.__init then obj:__init( transUnitIF, importModuleInfo, moduleType, builtinFunc, globalScope, macroCtrl, typeNameCtrl, helperInfo, importedAliasMap, validMutControl ); end
   return obj
end
function Import:__init(transUnitIF, importModuleInfo, moduleType, builtinFunc, globalScope, macroCtrl, typeNameCtrl, helperInfo, importedAliasMap, validMutControl) 
   self.validMutControl = validMutControl
   self.transUnitIF = transUnitIF
   self.importModuleInfo = importModuleInfo
   self.moduleType = moduleType
   self.builtinFunc = builtinFunc
   self.globalScope = globalScope
   self.macroCtrl = macroCtrl
   self.typeNameCtrl = typeNameCtrl
   self.helperInfo = helperInfo
   self.importedAliasMap = importedAliasMap
   self.importModule2ModuleInfo = {}
   self.importModuleName2ModuleInfo = {}
   
end
function Import.setmeta( obj )
  setmetatable( obj, { __index = Import  } )
end
function Import:get_importModule2ModuleInfo()
   return self.importModule2ModuleInfo
end


local _TypeInfo = {}

local ImportParam = {}
function ImportParam.setmeta( obj )
  setmetatable( obj, { __index = ImportParam  } )
end
function ImportParam.new( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap )
   local obj = {}
   ImportParam.setmeta( obj )
   if obj.__init then
      obj:__init( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap )
   end
   return obj
end
function ImportParam:__init( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap )

   self.pos = pos
   self.modifier = modifier
   self.processInfo = processInfo
   self.typeId2Scope = typeId2Scope
   self.typeId2TypeInfo = typeId2TypeInfo
   self.importedAliasMap = importedAliasMap
   self.lazyModuleSet = lazyModuleSet
   self.metaInfo = metaInfo
   self.scope = scope
   self.moduleTypeInfo = moduleTypeInfo
   self.scopeAccess = scopeAccess
   self.typeId2AtomMap = typeId2AtomMap
end


setmetatable( _TypeInfo, { ifList = {Mapping,} } )
function _TypeInfo.new(  )
   local obj = {}
   _TypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function _TypeInfo:__init() 
   self.typeId = Ast.rootTypeId
   self.skind = Ast.SerializeKind.Normal
end
function _TypeInfo:createTypeInfoCache( param )

   do
      local typeInfo = param.typeId2TypeInfo[self.typeId]
      if typeInfo ~= nil then
         return typeInfo, nil
      end
   end
   
   local typeInfo, mess = self:createTypeInfo( param )
   if typeInfo ~= nil then
      param.typeId2TypeInfo[self.typeId] = typeInfo
   end
   
   return typeInfo, mess
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
   table.insert( memInfo, { name = "typeId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


function ImportParam:getTypeInfo( typeId )

   do
      local typeInfo = self.typeId2TypeInfo[typeId]
      if typeInfo ~= nil then
         return typeInfo, nil
      end
   end
   
   do
      local atom = self.typeId2AtomMap[typeId]
      if atom ~= nil then
         local typeInfo, mess = atom:createTypeInfoCache( self )
         if typeInfo ~= nil then
            self.typeId2TypeInfo[typeId] = typeInfo
         end
         
         return typeInfo, mess
      end
   end
   
   return nil, nil
end


local _IdInfo = {}
setmetatable( _IdInfo, { ifList = {Mapping,} } )
function _IdInfo.setmeta( obj )
  setmetatable( obj, { __index = _IdInfo  } )
end
function _IdInfo.new( id, mod )
   local obj = {}
   _IdInfo.setmeta( obj )
   if obj.__init then
      obj:__init( id, mod )
   end
   return obj
end
function _IdInfo:__init( id, mod )

   self.id = id
   self.mod = mod
end
function _IdInfo:_toMap()
  return self
end
function _IdInfo._fromMap( val )
  local obj, mes = _IdInfo._fromMapSub( {}, val )
  if obj then
     _IdInfo.setmeta( obj )
  end
  return obj, mes
end
function _IdInfo._fromStem( val )
  return _IdInfo._fromMap( val )
end

function _IdInfo._fromMapSub( obj, val )
   local memInfo = {}
   table.insert( memInfo, { name = "id", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "mod", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


function ImportParam:getTypeInfoFrom( typeId )

   return self:getTypeInfo( typeId.id )
end


local _TypeInfoNilable = {}
setmetatable( _TypeInfoNilable, { __index = _TypeInfo } )
function _TypeInfoNilable:createTypeInfo( param )

   local orgTypeInfo = param:getTypeInfo( self.orgTypeId )
   if  nil == orgTypeInfo then
      local _orgTypeInfo = orgTypeInfo
   
      Util.err( string.format( "failed to createTypeInfo -- self.orgTypeId = %d", self.orgTypeId) )
   end
   
   local newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoNilable.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoNilable  } )
end
function _TypeInfoNilable.new( orgTypeId )
   local obj = {}
   _TypeInfoNilable.setmeta( obj )
   if obj.__init then
      obj:__init( orgTypeId )
   end
   return obj
end
function _TypeInfoNilable:__init( orgTypeId )

   _TypeInfo.__init( self)
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
   table.insert( memInfo, { name = "orgTypeId", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoAlias = {}
setmetatable( _TypeInfoAlias, { __index = _TypeInfo } )
function _TypeInfoAlias:createTypeInfo( param )
   local __func__ = '@lune.@base.@Import._TypeInfoAlias.createTypeInfo'

   local _
   local srcTypeInfo = _lune.unwrap( param:getTypeInfoFrom( self.srcTypeId ))
   local newTypeInfo = param.processInfo:createAlias( param.processInfo, self.rawTxt, true, Ast.AccessMode.Pub, param.moduleTypeInfo, srcTypeInfo )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   local _3702 = param:getTypeInfo( self.parentId )
   if  nil == _3702 then
      local __3702 = _3702
   
      return nil, string.format( "%s: not found parentInfo %d %s", __func__, self.parentId, self.rawTxt)
   end
   
   
   local parentScope = param.typeId2Scope[self.parentId]
   if  nil == parentScope then
      local _parentScope = parentScope
   
      return nil, string.format( "%s: not found parentScope %s %s", __func__, tostring( self.parentId), self.rawTxt)
   end
   
   parentScope:addAliasForType( param.processInfo, self.rawTxt, nil, newTypeInfo )
   param.importedAliasMap[srcTypeInfo] = newTypeInfo
   
   return newTypeInfo, nil
end
function _TypeInfoAlias.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoAlias  } )
end
function _TypeInfoAlias.new( parentId, rawTxt, srcTypeId )
   local obj = {}
   _TypeInfoAlias.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, rawTxt, srcTypeId )
   end
   return obj
end
function _TypeInfoAlias:__init( parentId, rawTxt, srcTypeId )

   _TypeInfo.__init( self)
   self.parentId = parentId
   self.rawTxt = rawTxt
   self.srcTypeId = srcTypeId
end
function _TypeInfoAlias:_toMap()
  return self
end
function _TypeInfoAlias._fromMap( val )
  local obj, mes = _TypeInfoAlias._fromMapSub( {}, val )
  if obj then
     _TypeInfoAlias.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoAlias._fromStem( val )
  return _TypeInfoAlias._fromMap( val )
end

function _TypeInfoAlias._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "rawTxt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "srcTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoDDD = {}
setmetatable( _TypeInfoDDD, { __index = _TypeInfo } )
function _TypeInfoDDD:createTypeInfo( param )

   local itemTypeInfo = _lune.unwrap( param:getTypeInfoFrom( self.itemTypeId ))
   local newTypeInfo = param.processInfo:createDDD( itemTypeInfo, true, self.extTypeFlag )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoDDD.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoDDD  } )
end
function _TypeInfoDDD.new( parentId, itemTypeId, extTypeFlag )
   local obj = {}
   _TypeInfoDDD.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, itemTypeId, extTypeFlag )
   end
   return obj
end
function _TypeInfoDDD:__init( parentId, itemTypeId, extTypeFlag )

   _TypeInfo.__init( self)
   self.parentId = parentId
   self.itemTypeId = itemTypeId
   self.extTypeFlag = extTypeFlag
end
function _TypeInfoDDD:_toMap()
  return self
end
function _TypeInfoDDD._fromMap( val )
  local obj, mes = _TypeInfoDDD._fromMapSub( {}, val )
  if obj then
     _TypeInfoDDD.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoDDD._fromStem( val )
  return _TypeInfoDDD._fromMap( val )
end

function _TypeInfoDDD._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "itemTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "extTypeFlag", func = _lune._toBool, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoAlternate = {}
setmetatable( _TypeInfoAlternate, { __index = _TypeInfo } )
function _TypeInfoAlternate:createTypeInfo( param )

   local baseInfo = _lune.unwrap( param:getTypeInfoFrom( self.baseId ))
   local interfaceList = {}
   for __index, ifTypeId in ipairs( self.ifList ) do
      table.insert( interfaceList, _lune.unwrap( param:getTypeInfoFrom( ifTypeId )) )
   end
   
   local newTypeInfo = param.processInfo:createAlternate( self.belongClassFlag, self.altIndex, self.txt, self.accessMode, param.moduleTypeInfo, baseInfo, interfaceList )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoAlternate.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoAlternate  } )
end
function _TypeInfoAlternate.new( parentId, txt, accessMode, baseId, ifList, belongClassFlag, altIndex )
   local obj = {}
   _TypeInfoAlternate.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, txt, accessMode, baseId, ifList, belongClassFlag, altIndex )
   end
   return obj
end
function _TypeInfoAlternate:__init( parentId, txt, accessMode, baseId, ifList, belongClassFlag, altIndex )

   _TypeInfo.__init( self)
   self.parentId = parentId
   self.txt = txt
   self.accessMode = accessMode
   self.baseId = baseId
   self.ifList = ifList
   self.belongClassFlag = belongClassFlag
   self.altIndex = altIndex
end
function _TypeInfoAlternate:_toMap()
  return self
end
function _TypeInfoAlternate._fromMap( val )
  local obj, mes = _TypeInfoAlternate._fromMapSub( {}, val )
  if obj then
     _TypeInfoAlternate.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoAlternate._fromStem( val )
  return _TypeInfoAlternate._fromMap( val )
end

function _TypeInfoAlternate._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "baseId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "ifList", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "belongClassFlag", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "altIndex", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoGeneric = {}
setmetatable( _TypeInfoGeneric, { __index = _TypeInfo } )
function _TypeInfoGeneric:createTypeInfo( param )

   local genSrcTypeInfo = _lune.unwrap( param:getTypeInfoFrom( self.genSrcTypeId ))
   local genTypeList = {}
   for __index, typeId in ipairs( self.genTypeList ) do
      table.insert( genTypeList, _lune.unwrap( param:getTypeInfoFrom( typeId )) )
   end
   
   local newTypeInfo = param.processInfo:createGeneric( genSrcTypeInfo, genTypeList, param.moduleTypeInfo )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   param.typeId2Scope[self.typeId] = Ast.getScope( newTypeInfo )
   return newTypeInfo, nil
end
function _TypeInfoGeneric.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoGeneric  } )
end
function _TypeInfoGeneric.new( genSrcTypeId, genTypeList )
   local obj = {}
   _TypeInfoGeneric.setmeta( obj )
   if obj.__init then
      obj:__init( genSrcTypeId, genTypeList )
   end
   return obj
end
function _TypeInfoGeneric:__init( genSrcTypeId, genTypeList )

   _TypeInfo.__init( self)
   self.genSrcTypeId = genSrcTypeId
   self.genTypeList = genTypeList
end
function _TypeInfoGeneric:_toMap()
  return self
end
function _TypeInfoGeneric._fromMap( val )
  local obj, mes = _TypeInfoGeneric._fromMapSub( {}, val )
  if obj then
     _TypeInfoGeneric.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoGeneric._fromStem( val )
  return _TypeInfoGeneric._fromMap( val )
end

function _TypeInfoGeneric._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "genSrcTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "genTypeList", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoBox = {}
setmetatable( _TypeInfoBox, { __index = _TypeInfo } )
function _TypeInfoBox:createTypeInfo( param )

   local boxingType = _lune.unwrap( param:getTypeInfo( self.boxingType ))
   local newTypeInfo = param.processInfo:createBox( self.accessMode, boxingType )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoBox.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoBox  } )
end
function _TypeInfoBox.new( accessMode, boxingType )
   local obj = {}
   _TypeInfoBox.setmeta( obj )
   if obj.__init then
      obj:__init( accessMode, boxingType )
   end
   return obj
end
function _TypeInfoBox:__init( accessMode, boxingType )

   _TypeInfo.__init( self)
   self.accessMode = accessMode
   self.boxingType = boxingType
end
function _TypeInfoBox:_toMap()
  return self
end
function _TypeInfoBox._fromMap( val )
  local obj, mes = _TypeInfoBox._fromMapSub( {}, val )
  if obj then
     _TypeInfoBox.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoBox._fromStem( val )
  return _TypeInfoBox._fromMap( val )
end

function _TypeInfoBox._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "boxingType", func = _lune._toInt, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoExt = {}
setmetatable( _TypeInfoExt, { __index = _TypeInfo } )
function _TypeInfoExt:createTypeInfo( param )

   local extedType = _lune.unwrap( param:getTypeInfoFrom( self.extedTypeId ))
   local newTypeInfo
   
   do
      local _matchExp = param.processInfo:createLuaval( extedType, true )
      if _matchExp[1] == Ast.LuavalResult.OK[1] then
         local extType = _matchExp[2][1]
         local _ = _matchExp[2][2]
      
         newTypeInfo = extType
      elseif _matchExp[1] == Ast.LuavalResult.Err[1] then
         local mess = _matchExp[2][1]
      
         Util.err( mess )
      end
   end
   
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoExt.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoExt  } )
end
function _TypeInfoExt.new( extedTypeId )
   local obj = {}
   _TypeInfoExt.setmeta( obj )
   if obj.__init then
      obj:__init( extedTypeId )
   end
   return obj
end
function _TypeInfoExt:__init( extedTypeId )

   _TypeInfo.__init( self)
   self.extedTypeId = extedTypeId
end
function _TypeInfoExt:_toMap()
  return self
end
function _TypeInfoExt._fromMap( val )
  local obj, mes = _TypeInfoExt._fromMapSub( {}, val )
  if obj then
     _TypeInfoExt.setmeta( obj )
  end
  return obj, mes
end
function _TypeInfoExt._fromStem( val )
  return _TypeInfoExt._fromMap( val )
end

function _TypeInfoExt._fromMapSub( obj, val )
   local result, mes = _TypeInfo._fromMapSub( obj, val )
   if not result then
      return nil, mes
   end

   local memInfo = {}
   table.insert( memInfo, { name = "extedTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoModifier = {}
setmetatable( _TypeInfoModifier, { __index = _TypeInfo } )
function _TypeInfoModifier:createTypeInfo( param )

   local srcTypeInfo = param:getTypeInfoFrom( self.srcTypeId )
   if  nil == srcTypeInfo then
      local _srcTypeInfo = srcTypeInfo
   
      return nil, string.format( "not found srcType -- %d", self.srcTypeId.id)
   end
   
   local newTypeInfo = param.modifier:createModifier( srcTypeInfo, self.mutMode )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   return newTypeInfo, nil
end
function _TypeInfoModifier.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoModifier  } )
end
function _TypeInfoModifier.new( srcTypeId, mutMode )
   local obj = {}
   _TypeInfoModifier.setmeta( obj )
   if obj.__init then
      obj:__init( srcTypeId, mutMode )
   end
   return obj
end
function _TypeInfoModifier:__init( srcTypeId, mutMode )

   _TypeInfo.__init( self)
   self.srcTypeId = srcTypeId
   self.mutMode = mutMode
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
   table.insert( memInfo, { name = "srcTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "mutMode", func = Ast.MutMode._from, nilable = false, child = {} } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end


local _TypeInfoModule = {}
setmetatable( _TypeInfoModule, { __index = _TypeInfo } )
function _TypeInfoModule:createTypeInfo( param )
   local __func__ = '@lune.@base.@Import._TypeInfoModule.createTypeInfo'

   local parentInfo = Ast.headTypeInfo
   if self.parentId ~= Ast.rootTypeId then
      local workTypeInfo = param:getTypeInfo( self.parentId )
      if  nil == workTypeInfo then
         local _workTypeInfo = workTypeInfo
      
         Util.err( string.format( "not found parentInfo %d %s", self.parentId, self.txt) )
      end
      
      parentInfo = workTypeInfo
   end
   
   local parentScope = param.typeId2Scope[self.parentId]
   if  nil == parentScope then
      local _parentScope = parentScope
   
      return nil, string.format( "%s: not found parentScope %s %s", __func__, tostring( self.parentId), self.txt)
   end
   
   
   local newTypeInfo = parentScope:getTypeInfoChild( self.txt )
   do
      local _exp = newTypeInfo
      if _exp ~= nil then
         param.typeId2Scope[self.typeId] = Ast.getScope( _exp )
         if not _exp:get_scope() then
            return nil, string.format( "not found scope %s %d %s %s %s", tostring( parentScope), self.parentId, tostring( self.typeId), self.txt, _exp:getTxt(  ))
         end
         
         param.typeId2TypeInfo[self.typeId] = _exp
      else
         local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
         
         local mutable = false
         if self.typeId == param.metaInfo.__moduleTypeId then
            mutable = param.metaInfo.__moduleMutable
         end
         
         local workTypeInfo = param.processInfo:createModule( scope, parentInfo, true, self.txt, mutable )
         
         newTypeInfo = workTypeInfo
         param.typeId2Scope[self.typeId] = scope
         param.typeId2TypeInfo[self.typeId] = workTypeInfo
         parentScope:addClass( param.processInfo, self.txt, nil, workTypeInfo )
         
         Log.log( Log.Level.Info, __func__, 379, function (  )
         
            return string.format( "new module -- %s, %s, %d, %d, %d", self.txt, workTypeInfo:getFullName( Ast.defaultTypeNameCtrl, parentScope, false ), self.typeId, workTypeInfo:get_typeId().id, parentScope:get_scopeId())
         end )
         
      end
   end
   
   return newTypeInfo, nil
end
function _TypeInfoModule.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoModule  } )
end
function _TypeInfoModule.new( parentId, txt )
   local obj = {}
   _TypeInfoModule.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, txt )
   end
   return obj
end
function _TypeInfoModule:__init( parentId, txt )

   _TypeInfo.__init( self)
   self.parentId = parentId
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
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
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
   local __func__ = '@lune.@base.@Import._TypeInfoNormal.createTypeInfo'

   local newTypeInfo = nil
   if self.parentId ~= Ast.rootTypeId or not Ast.getBuiltInTypeIdMap(  )[self.typeId] or self.kind == Ast.TypeInfoKind.List or self.kind == Ast.TypeInfoKind.Array or self.kind == Ast.TypeInfoKind.Map or self.kind == Ast.TypeInfoKind.Set then
      local parentInfo = Ast.headTypeInfo
      if self.parentId ~= Ast.rootTypeId then
         local workTypeInfo = param:getTypeInfo( self.parentId )
         if  nil == workTypeInfo then
            local _workTypeInfo = workTypeInfo
         
            return nil, string.format( "not found parentInfo %d %s", self.parentId, self.txt)
         end
         
         parentInfo = workTypeInfo
      end
      
      
      local itemTypeInfo = {}
      for __index, typeId in ipairs( self.itemTypeId ) do
         table.insert( itemTypeInfo, _lune.unwrap( param:getTypeInfoFrom( typeId )) )
      end
      
      local argTypeInfo = {}
      for index, typeId in ipairs( self.argTypeId ) do
         local argType, mess = param:getTypeInfoFrom( typeId )
         if argType ~= nil then
            table.insert( argTypeInfo, argType )
         else
            local errmess = string.format( "not found arg (index:%d) -- %s.%s, %d, %d. %s", index, parentInfo:getTxt(  ), self.txt, typeId.id, #self.argTypeId, tostring( mess))
            return nil, errmess
         end
         
      end
      
      local retTypeInfo = {}
      for __index, typeId in ipairs( self.retTypeId ) do
         table.insert( retTypeInfo, _lune.unwrap( param:getTypeInfoFrom( typeId )) )
      end
      
      local baseInfo = _lune.unwrap( param:getTypeInfoFrom( self.baseId ))
      local interfaceList = {}
      for __index, ifTypeId in ipairs( self.ifList ) do
         table.insert( interfaceList, _lune.unwrap( param:getTypeInfoFrom( ifTypeId )) )
      end
      
      
      local parentScope = param.typeId2Scope[self.parentId]
      if  nil == parentScope then
         local _parentScope = parentScope
      
         return nil, string.format( "%s: not found parentScope %s %s", __func__, tostring( self.parentId), self.txt)
      end
      
      
      if self.txt ~= "" then
         newTypeInfo = parentScope:getTypeInfoChild( self.txt )
      end
      
      if newTypeInfo and (self.kind == Ast.TypeInfoKind.Class or self.kind == Ast.TypeInfoKind.ExtModule or self.kind == Ast.TypeInfoKind.IF ) then
         do
            local _exp = newTypeInfo
            if _exp ~= nil then
               param.typeId2Scope[self.typeId] = Ast.getScope( _exp )
               if not _exp:get_scope() then
                  Util.err( string.format( "not found scope %s %s %s %s %s", tostring( parentScope), tostring( self.parentId), tostring( self.typeId), self.txt, _exp:getTxt(  )) )
               end
               
               param.typeId2TypeInfo[self.typeId] = _exp
            end
         end
         
         
      else
       
         if self.kind == Ast.TypeInfoKind.Class or self.kind == Ast.TypeInfoKind.IF then
            Log.log( Log.Level.Debug, __func__, 486, function (  )
            
               return string.format( "new type -- %d, %s -- %s, %d", self.parentId, self.txt, _lune.nilacc( parentScope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , Ast.defaultTypeNameCtrl, parentScope, false ) or "nil", _lune.nilacc( _lune.nilacc( parentScope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1)
            end )
            
            
            local baseScope = _lune.unwrap( Ast.getScope( baseInfo ))
            local ifScopeList = {}
            for __index, ifType in ipairs( interfaceList ) do
               table.insert( ifScopeList, _lune.unwrap( ifType:get_scope()) )
            end
            
            
            local scope = Ast.Scope.new(param.processInfo, parentScope, true, baseScope, ifScopeList)
            
            local altTypeList = {}
            for __index, itemType in ipairs( itemTypeInfo ) do
               table.insert( altTypeList, _lune.unwrap( (_lune.__Cast( itemType, 3, Ast.AlternateTypeInfo ) )) )
            end
            
            
            local workTypeInfo = param.processInfo:createClass( self.kind == Ast.TypeInfoKind.Class, self.abstractFlag, scope, baseInfo, interfaceList, altTypeList, parentInfo, true, Ast.AccessMode.Pub, self.txt )
            parentScope:addClassLazy( param.processInfo, self.txt, nil, workTypeInfo, _lune._Set_has(param.lazyModuleSet, self.typeId ) )
            
            newTypeInfo = workTypeInfo
            
            param.typeId2Scope[self.typeId] = scope
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
         elseif self.kind == Ast.TypeInfoKind.ExtModule then
            Log.log( Log.Level.Debug, __func__, 523, function (  )
            
               return string.format( "new type -- %d, %s -- %s, %d", self.parentId, self.txt, _lune.nilacc( parentScope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , Ast.defaultTypeNameCtrl, parentScope, false ) or "nil", _lune.nilacc( _lune.nilacc( parentScope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1)
            end )
            
            
            local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil, {})
            
            local workTypeInfo = param.processInfo:createExtModule( scope, parentInfo, true, Ast.AccessMode.Pub, self.txt, _lune.unwrap( self.moduleLang), _lune.unwrap( self.requirePath) )
            parentScope:addExtModule( param.processInfo, self.txt, nil, workTypeInfo, _lune._Set_has(param.lazyModuleSet, self.typeId ), _lune.unwrap( self.moduleLang) )
            
            newTypeInfo = workTypeInfo
            
            param.typeId2Scope[self.typeId] = scope
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
         else
          
            local scope = nil
            
            if self.kind == Ast.TypeInfoKind.Func or self.kind == Ast.TypeInfoKind.Method then
               scope = Ast.Scope.new(param.processInfo, parentScope, false, nil)
            end
            
            
            local typeInfoKind = self.kind
            local accessMode = self.accessMode
            local workTypeInfo = Ast.NormalTypeInfo.create( param.processInfo, accessMode, self.abstractFlag, scope, baseInfo, parentInfo, self.staticFlag, typeInfoKind, self.txt, itemTypeInfo, argTypeInfo, retTypeInfo, self.mutMode )
            newTypeInfo = workTypeInfo
            
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
            
            do
               local _switchExp = self.kind
               if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.Macro or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
                  local symbolKind = Ast.SymbolKind.Fun
                  do
                     local _switchExp = self.kind
                     if _switchExp == Ast.TypeInfoKind.Method then
                        symbolKind = Ast.SymbolKind.Mtd
                     elseif _switchExp == Ast.TypeInfoKind.Macro then
                        symbolKind = Ast.SymbolKind.Mac
                     elseif _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
                        symbolKind = Ast.SymbolKind.Typ
                     end
                  end
                  
                  local workParentScope = _lune.unwrap( param.typeId2Scope[self.parentId])
                  workParentScope:add( param.processInfo, symbolKind, false, self.kind == Ast.TypeInfoKind.Func, self.txt, nil, workTypeInfo, accessMode, self.staticFlag, Ast.MutMode.IMut, true, false )
                  param.typeId2Scope[self.typeId] = scope
               end
            end
            
         end
         
      end
      
   else
    
      newTypeInfo = param.scope:getTypeInfo( self.txt, param.scope, false, param.scopeAccess )
      if not newTypeInfo then
         for key, val in pairs( self:_toMap(  ) ) do
            Util.errorLog( string.format( "error: illegal self %s:%s", key, tostring( val)) )
         end
         
      end
      
      param.typeId2TypeInfo[self.typeId] = _lune.unwrap( newTypeInfo)
   end
   
   return newTypeInfo, nil
end
function _TypeInfoNormal.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoNormal  } )
end
function _TypeInfoNormal.new( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )
   local obj = {}
   _TypeInfoNormal.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )
   end
   return obj
end
function _TypeInfoNormal:__init( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )

   _TypeInfo.__init( self)
   self.parentId = parentId
   self.abstractFlag = abstractFlag
   self.baseId = baseId
   self.txt = txt
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.kind = kind
   self.mutMode = mutMode
   self.ifList = ifList
   self.itemTypeId = itemTypeId
   self.argTypeId = argTypeId
   self.retTypeId = retTypeId
   self.children = children
   self.moduleLang = moduleLang
   self.requirePath = requirePath
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
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "abstractFlag", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "baseId", func = _IdInfo._fromMap, nilable = false, child = {} } )
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "staticFlag", func = _lune._toBool, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "kind", func = Ast.TypeInfoKind._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "mutMode", func = Ast.MutMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "ifList", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "itemTypeId", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "argTypeId", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "retTypeId", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "children", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
   table.insert( memInfo, { name = "moduleLang", func = Types.Lang._from, nilable = true, child = {} } )
   table.insert( memInfo, { name = "requirePath", func = _lune._toStr, nilable = true, child = {} } )
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
   local parentInfo = _lune.unwrap( param:getTypeInfo( self.parentId ))
   local parentScope = _lune.unwrap( Ast.getScope( parentInfo ))
   local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
   
   param.typeId2Scope[self.typeId] = scope
   local valTypeInfo = _lune.unwrap( param:getTypeInfo( self.valTypeId ))
   local enumTypeInfo = param.processInfo:createEnum( scope, parentInfo, true, accessMode, self.txt, valTypeInfo )
   local newTypeInfo = enumTypeInfo
   param.typeId2TypeInfo[self.typeId] = enumTypeInfo
   
   local function getEnumLiteral( val )
   
      do
         local _switchExp = valTypeInfo
         if _switchExp == Ast.builtinTypeInt then
            return _lune.newAlge( Ast.EnumLiteral.Int, {math.floor(val)})
         elseif _switchExp == Ast.builtinTypeReal then
            return _lune.newAlge( Ast.EnumLiteral.Real, {val * 1.0})
         elseif _switchExp == Ast.builtinTypeString then
            return _lune.newAlge( Ast.EnumLiteral.Str, {val})
         end
      end
      
      return nil
   end
   for valName, valData in pairs( self.enumValList ) do
      local val = getEnumLiteral( valData )
      if  nil == val then
         local _val = val
      
         return nil, string.format( "unknown enum val type -- %s", valTypeInfo:getTxt(  ))
      end
      
      local evalValSym = _lune.unwrap( scope:addEnumVal( param.processInfo, valName, nil, enumTypeInfo ))
      enumTypeInfo:addEnumValInfo( Ast.EnumValInfo.new(valName, val, evalValSym) )
   end
   
   parentScope:addEnum( param.processInfo, accessMode, self.txt, nil, enumTypeInfo )
   return newTypeInfo, nil
end
function _TypeInfoEnum.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoEnum  } )
end
function _TypeInfoEnum.new( parentId, txt, accessMode, valTypeId, enumValList )
   local obj = {}
   _TypeInfoEnum.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, txt, accessMode, valTypeId, enumValList )
   end
   return obj
end
function _TypeInfoEnum:__init( parentId, txt, accessMode, valTypeId, enumValList )

   _TypeInfo.__init( self)
   self.parentId = parentId
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
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
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
setmetatable( _TypeInfoAlgeVal, { ifList = {Mapping,} } )
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
   table.insert( memInfo, { name = "typeList", func = _lune._toList, nilable = false, child = { { func = _IdInfo._fromMap, nilable = false, child = {} } } } )
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
   local parentInfo = _lune.unwrap( param:getTypeInfo( self.parentId ))
   local parentScope = _lune.unwrap( Ast.getScope( parentInfo ))
   local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
   
   param.typeId2Scope[self.typeId] = scope
   local algeTypeInfo = param.processInfo:createAlge( scope, parentInfo, true, accessMode, self.txt )
   local newTypeInfo = algeTypeInfo
   param.typeId2TypeInfo[self.typeId] = algeTypeInfo
   for __index, valInfo in ipairs( self.algeValList ) do
      local typeInfoList = {}
      for __index, orgTypeId in ipairs( valInfo.typeList ) do
         table.insert( typeInfoList, _lune.unwrap( param:getTypeInfoFrom( orgTypeId )) )
      end
      
      local algeValSym = scope:addAlgeVal( param.processInfo, valInfo.name, nil, algeTypeInfo )
      local algeVal = Ast.AlgeValInfo.new(valInfo.name, typeInfoList, algeTypeInfo, _lune.unwrap( algeValSym))
      algeTypeInfo:addValInfo( algeVal )
   end
   
   parentScope:addAlge( param.processInfo, accessMode, self.txt, nil, algeTypeInfo )
   return newTypeInfo, nil
end
function _TypeInfoAlge.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoAlge  } )
end
function _TypeInfoAlge.new( parentId, txt, accessMode, algeValList )
   local obj = {}
   _TypeInfoAlge.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, txt, accessMode, algeValList )
   end
   return obj
end
function _TypeInfoAlge:__init( parentId, txt, accessMode, algeValList )

   _TypeInfo.__init( self)
   self.parentId = parentId
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
   table.insert( memInfo, { name = "parentId", func = _lune._toInt, nilable = false, child = {} } )
   table.insert( memInfo, { name = "txt", func = _lune._toStr, nilable = false, child = {} } )
   table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
   table.insert( memInfo, { name = "algeValList", func = _lune._toList, nilable = false, child = { { func = _TypeInfoAlgeVal._fromMap, nilable = false, child = {} } } } )
   local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
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


function Import:processImport( processInfo, modulePath, depth )
   local __func__ = '@lune.@base.@Import.Import.processImport'

   local modifier = TransUnitIF.Modifier.new(self.validMutControl, processInfo)
   
   local orgModulePath = modulePath
   modulePath = frontInterface.getLuaModulePath( modulePath )
   
   Log.log( Log.Level.Info, __func__, 732, function (  )
   
      return string.format( "%s -> %s start", self.moduleType:getTxt( self.typeNameCtrl ), orgModulePath)
   end )
   
   
   if not self.importModuleInfo:add( orgModulePath ) then
      self.transUnitIF:error( string.format( "recursive import: %s -> %s", self.importModuleInfo:getFull(  ), orgModulePath) )
   end
   
   
   do
      local moduleInfo = self.importModuleName2ModuleInfo[modulePath]
      if moduleInfo ~= nil then
         do
            local moduleMeta = frontInterface.loadMeta( self.importModuleInfo, orgModulePath )
            if moduleMeta ~= nil then
               Log.log( Log.Level.Info, __func__, 747, function (  )
               
                  return string.format( "%s already", orgModulePath)
               end )
               
               
               local metaInfo = moduleMeta:get_metaInfo()
               local typeId2TypeInfo = moduleInfo:get_importId2localTypeInfoMap()
               self.importModuleInfo:remove(  )
               
               for key, val in pairs( moduleInfo:get_importedAliasMap() ) do
                  self.importedAliasMap[key] = val
               end
               
               return metaInfo, typeId2TypeInfo, moduleInfo
            end
         end
         
         self.transUnitIF:error( "failed to load meta -- " .. orgModulePath )
      end
   end
   
   
   local nameList = {}
   for txt in string.gmatch( modulePath, '[^%./:]+' ) do
      table.insert( nameList, txt )
   end
   
   
   local moduleMeta = frontInterface.loadMeta( self.importModuleInfo, orgModulePath )
   if  nil == moduleMeta then
      local _moduleMeta = moduleMeta
   
      self.transUnitIF:error( "failed to load meta -- " .. orgModulePath )
   end
   
   local importedProcessInfo = moduleMeta:get_processInfo()
   local metaInfo = moduleMeta:get_metaInfo()
   Log.log( Log.Level.Info, __func__, 775, function (  )
   
      return string.format( "%s processing -- %s", orgModulePath, tostring( importedProcessInfo))
   end )
   
   
   local dependLibId2DependInfo = {}
   do
      local __sorted = {}
      local __map = metaInfo.__dependModuleMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, dependName in ipairs( __sorted ) do
         local dependInfo = __map[ dependName ]
         do
            if dependInfo['use'] then
               local _
               local _4055, metaTypeId2TypeInfoMap = self:processImport( processInfo, dependName, depth + 1 )
               local typeId = math.floor((_lune.unwrap( dependInfo['typeId']) ))
               dependLibId2DependInfo[typeId] = DependModuleInfo.new(typeId, metaTypeId2TypeInfoMap)
            end
            
         end
      end
   end
   
   local typeId2TypeInfo = {}
   typeId2TypeInfo[Ast.rootTypeId] = Ast.headTypeInfo
   local typeId2Scope = {}
   typeId2Scope[Ast.rootTypeId] = self.transUnitIF:get_scope()
   
   typeId2TypeInfo[self.builtinFunc.lnsthread_:get_typeId().id] = self.builtinFunc.lnsthread_
   typeId2Scope[self.builtinFunc.lnsthread_:get_typeId().id] = self.builtinFunc.lnsthread_:get_scope()
   for typeId, dependIdInfo in pairs( metaInfo.__dependIdMap ) do
      local dependInfo = _lune.unwrap( dependLibId2DependInfo[_lune.unwrap( dependIdInfo[1])])
      local typeInfo = dependInfo:getTypeInfo( _lune.unwrap( dependIdInfo[2]) )
      typeId2TypeInfo[typeId] = typeInfo
      do
         local _exp = Ast.getScope( typeInfo )
         if _exp ~= nil then
            typeId2Scope[typeId] = _exp
         end
      end
      
   end
   
   
   local moduleTypeInfo = Ast.headTypeInfo
   for index, moduleName in ipairs( nameList ) do
      local mutable = false
      if index == #nameList then
         mutable = metaInfo.__moduleMutable
      end
      
      moduleTypeInfo = self.transUnitIF:pushModule( processInfo, true, moduleName, mutable )
   end
   
   for __index, _4073 in ipairs( nameList ) do
      self.transUnitIF:popModule(  )
   end
   
   for __index, symbolInfo in pairs( Ast.getSym2builtInTypeMap(  ) ) do
      typeId2TypeInfo[symbolInfo:get_typeInfo():get_typeId(  ).id] = symbolInfo:get_typeInfo()
   end
   
   for __index, builtinTypeInfo in pairs( Ast.getBuiltInTypeIdMap(  ) ) do
      typeId2TypeInfo[builtinTypeInfo:get_typeId().id] = builtinTypeInfo
   end
   
   
   local newId2OldIdMap = {}
   
   local _typeInfoList = {}
   local id2atomMap = {}
   local _typeInfoNormalList = {}
   for __index, atomInfoLua in pairs( metaInfo.__typeInfoList ) do
      local workAtomInfo = (atomInfoLua )
      if  nil == workAtomInfo then
         local _workAtomInfo = workAtomInfo
      
         self.transUnitIF:error( "illegal atomInfo" )
      end
      
      local atomInfo = workAtomInfo
      do
         local skind = atomInfo['skind']
         if skind ~= nil then
            local actInfo = nil
            local mess = nil
            local kind = _lune.unwrap( Ast.SerializeKind._from( math.floor(skind) ))
            do
               local _switchExp = kind
               if _switchExp == Ast.SerializeKind.Enum then
                  actInfo, mess = _TypeInfoEnum._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Alge then
                  actInfo, mess = _TypeInfoAlge._fromMap( atomInfo )
                  self.helperInfo.useAlge = true
               elseif _switchExp == Ast.SerializeKind.Module then
                  actInfo, mess = _TypeInfoModule._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Normal then
                  local workInfo
                  
                  workInfo, mess = _TypeInfoNormal._fromMap( atomInfo )
                  if workInfo ~= nil then
                     table.insert( _typeInfoNormalList, workInfo )
                  end
                  
                  actInfo = workInfo
               elseif _switchExp == Ast.SerializeKind.Nilable then
                  actInfo, mess = _TypeInfoNilable._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Alias then
                  actInfo, mess = _TypeInfoAlias._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.DDD then
                  actInfo, mess = _TypeInfoDDD._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Alternate then
                  actInfo, mess = _TypeInfoAlternate._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Generic then
                  actInfo, mess = _TypeInfoGeneric._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Modifier then
                  actInfo, mess = _TypeInfoModifier._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Box then
                  actInfo, mess = _TypeInfoBox._fromMap( atomInfo )
               elseif _switchExp == Ast.SerializeKind.Ext then
                  actInfo, mess = _TypeInfoExt._fromMap( atomInfo )
               end
            end
            
            if actInfo ~= nil then
               table.insert( _typeInfoList, actInfo )
               id2atomMap[actInfo.typeId] = actInfo
            else
               for key, val in pairs( atomInfo ) do
                  Util.errorLog( string.format( "table: %s:%s", key, tostring( val)) )
               end
               
               if mess ~= nil then
                  Util.errorLog( mess )
               end
               
               Util.err( string.format( "_TypeInfo.%s._fromMap error", Ast.SerializeKind:_getTxt( kind)
               ) )
            end
            
         end
      end
      
   end
   
   
   local orgId2MacroTypeInfo = {}
   
   local lazyModuleSet = {}
   for __index, typeId in pairs( metaInfo.__lazyModuleList ) do
      lazyModuleSet[typeId]= true
   end
   
   
   local importParam = ImportParam.new(self.transUnitIF:getLatestPos(  ), modifier, processInfo, typeId2Scope, typeId2TypeInfo, {}, lazyModuleSet, metaInfo, self.transUnitIF:get_scope(), moduleTypeInfo, Ast.ScopeAccess.Normal, id2atomMap)
   
   for __index, atomInfo in ipairs( _typeInfoList ) do
      local newTypeInfo, errMess = atomInfo:createTypeInfoCache( importParam )
      do
         local _exp = errMess
         if _exp ~= nil then
            Util.err( string.format( "Failed to createType -- %s: %s(%d): %s", orgModulePath, Ast.SerializeKind:_getTxt( atomInfo.skind)
            , atomInfo.typeId, _exp) )
         end
      end
      
      if newTypeInfo ~= nil then
         if newTypeInfo:get_kind() == Ast.TypeInfoKind.Macro then
            orgId2MacroTypeInfo[atomInfo.typeId] = newTypeInfo
         end
         
         if newTypeInfo:get_kind() == Ast.TypeInfoKind.Set then
            self.helperInfo.useSet = true
         end
         
         if newTypeInfo:get_accessMode() == Ast.AccessMode.Global then
            do
               local _switchExp = newTypeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.IF or _switchExp == Ast.TypeInfoKind.Class then
                  self.globalScope:addClass( processInfo, newTypeInfo:get_rawTxt(), nil, newTypeInfo )
               elseif _switchExp == Ast.TypeInfoKind.Func then
                  self.globalScope:addFunc( processInfo, nil, newTypeInfo, Ast.AccessMode.Global, newTypeInfo:get_staticFlag(), Ast.TypeInfo.isMut( newTypeInfo ) )
               elseif _switchExp == Ast.TypeInfoKind.Enum then
                  self.globalScope:addEnum( processInfo, Ast.AccessMode.Global, newTypeInfo:get_rawTxt(), nil, newTypeInfo )
               elseif _switchExp == Ast.TypeInfoKind.Nilable then
                  
                  
               else 
                  
                     Util.err( string.format( "%s: not support kind -- %s", __func__, Ast.TypeInfoKind:_getTxt( newTypeInfo:get_kind())
                     ) )
               end
            end
            
         end
         
      end
      
   end
   
   for __index, atomInfo in ipairs( _typeInfoNormalList ) do
      if #atomInfo.children > 0 then
         local scope = _lune.unwrap( typeId2Scope[atomInfo.typeId])
         for __index, childId in ipairs( atomInfo.children ) do
            local typeInfo = typeId2TypeInfo[childId.id]
            if  nil == typeInfo then
               local _typeInfo = typeInfo
            
               Util.err( string.format( "not found childId -- %s, %d, %s(%d)", orgModulePath, childId.id, atomInfo.txt, atomInfo.typeId) )
            end
            
            local symbolKind = Ast.SymbolKind.Typ
            local addFlag = true
            do
               local _switchExp = typeInfo:get_kind()
               if _switchExp == Ast.TypeInfoKind.Func then
                  symbolKind = Ast.SymbolKind.Fun
               elseif _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
                  symbolKind = Ast.SymbolKind.Typ
               elseif _switchExp == Ast.TypeInfoKind.Method then
                  symbolKind = Ast.SymbolKind.Mtd
               elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Module then
                  symbolKind = Ast.SymbolKind.Typ
               elseif _switchExp == Ast.TypeInfoKind.Enum then
                  addFlag = false
               end
            end
            
            
            if addFlag then
               scope:add( processInfo, symbolKind, false, typeInfo:get_kind() == Ast.TypeInfoKind.Func, typeInfo:getTxt(  ), nil, typeInfo, typeInfo:get_accessMode(), typeInfo:get_staticFlag(), typeInfo:get_mutMode(), true, false )
            end
            
         end
         
      end
      
   end
   
   for typeId, typeInfo in pairs( typeId2TypeInfo ) do
      newId2OldIdMap[typeInfo] = typeId
   end
   
   
   local function registMember( classTypeId )
      local __func__ = '@lune.@base.@Import.Import.processImport.registMember'
   
      if metaInfo.__dependIdMap[classTypeId] then
         return 
      end
      
      local classTypeInfo = _lune.unwrap( typeId2TypeInfo[classTypeId])
      
      do
         local _switchExp = (classTypeInfo:get_kind() )
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.ExtModule then
            self.transUnitIF:pushClassScope( self.transUnitIF:getLatestPos(  ), classTypeInfo )
            
            do
               local _exp = metaInfo.__typeId2ClassInfoMap[classTypeId]
               if _exp ~= nil then
                  local classInfo = (_exp )
                  if  nil == classInfo then
                     local _classInfo = classInfo
                  
                     self.transUnitIF:error( "illegal val" )
                  end
                  
                  for fieldName, fieldInfo in pairs( classInfo ) do
                     do
                        local typeId = fieldInfo['typeId']
                        if typeId ~= nil then
                           local fieldTypeInfo = _lune.unwrap( typeId2TypeInfo[math.floor(typeId)])
                           local symbolInfo = self.transUnitIF:get_scope():addMember( processInfo, fieldName, nil, fieldTypeInfo, _lune.unwrap( Ast.AccessMode._from( math.floor((_lune.unwrap( fieldInfo['accessMode']) )) )), fieldInfo['staticFlag'] and true or false, _lune.unwrap( Ast.MutMode._from( math.floor((_lune.unwrap( fieldInfo['mutMode']) )) )) )
                        else
                           self.transUnitIF:error( "not found fieldInfo.typeId" )
                        end
                     end
                     
                  end
                  
               else
                  self.transUnitIF:error( string.format( "not found class -- %s: %d, %s", orgModulePath, classTypeId, classTypeInfo:getTxt(  )) )
               end
            end
            
         elseif _switchExp == Ast.TypeInfoKind.Module then
            self.transUnitIF:pushModule( processInfo, true, classTypeInfo:getTxt(  ), Ast.TypeInfo.isMut( classTypeInfo ) )
            Log.log( Log.Level.Debug, __func__, 1056, function (  )
            
               return string.format( "push module -- %s, %s, %d, %d, %d", classTypeInfo:getTxt(  ), _lune.nilacc( self.transUnitIF:get_scope():get_ownerTypeInfo(), 'getFullName', 'callmtd' , Ast.defaultTypeNameCtrl, self.transUnitIF:get_scope(), false ) or "nil", _lune.nilacc( _lune.nilacc( self.transUnitIF:get_scope():get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1, classTypeInfo:get_typeId().id, self.transUnitIF:get_scope():get_parent():get_scopeId())
            end )
            
         end
      end
      
      
      for __index, child in ipairs( classTypeInfo:get_children(  ) ) do
         if child:get_kind(  ) == Ast.TypeInfoKind.Class or child:get_kind(  ) == Ast.TypeInfoKind.ExtModule or child:get_kind(  ) == Ast.TypeInfoKind.Module or child:get_kind(  ) == Ast.TypeInfoKind.IF then
            local oldId = newId2OldIdMap[child]
            if oldId then
               registMember( _lune.unwrap( oldId) )
            end
            
         end
         
      end
      
      
      do
         local _switchExp = classTypeInfo:get_kind()
         if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.ExtModule then
            self.transUnitIF:popClass(  )
         elseif _switchExp == Ast.TypeInfoKind.Module then
            self.transUnitIF:popModule(  )
         end
      end
      
   end
   for __index, atomInfo in ipairs( _typeInfoList ) do
      do
         local workInfo = _lune.__Cast( atomInfo, 3, _TypeInfoNormal )
         if workInfo ~= nil then
            if workInfo.parentId == Ast.rootTypeId then
               registMember( atomInfo.typeId )
            end
            
         else
            do
               local workInfo = _lune.__Cast( atomInfo, 3, _TypeInfoModule )
               if workInfo ~= nil then
                  if workInfo.parentId == Ast.rootTypeId then
                     registMember( atomInfo.typeId )
                  end
                  
               end
            end
            
         end
      end
      
   end
   
   
   for index, moduleName in ipairs( nameList ) do
      local mutable = false
      if index == #nameList then
         mutable = metaInfo.__moduleMutable
      end
      
      self.transUnitIF:pushModule( processInfo, true, moduleName, mutable )
   end
   
   
   for varName, varInfo in pairs( metaInfo.__varName2InfoMap ) do
      do
         local typeId = varInfo['typeId']
         if typeId ~= nil then
            self.transUnitIF:get_scope():addStaticVar( processInfo, false, true, varName, nil, _lune.unwrap( typeId2TypeInfo[math.floor(typeId)]), varInfo['mutable'] and Ast.MutMode.Mut or Ast.MutMode.IMut )
         else
            self.transUnitIF:error( "illegal varInfo.typeId" )
         end
      end
      
   end
   
   for orgTypeId, macroInfoStem in pairs( metaInfo.__macroName2InfoMap ) do
      
      self.macroCtrl:importMacro( processInfo, moduleMeta:get_lnsPath(), (macroInfoStem ), _lune.unwrap( orgId2MacroTypeInfo[orgTypeId]), typeId2TypeInfo )
   end
   
   
   for __index, _4212 in ipairs( nameList ) do
      self.transUnitIF:popModule(  )
   end
   
   
   if depth == 1 then
      for key, val in pairs( importParam.importedAliasMap ) do
         self.importedAliasMap[key] = val
      end
      
   end
   
   
   local moduleInfo = frontInterface.ModuleInfo.new(orgModulePath, nameList[#nameList], newId2OldIdMap, frontInterface.ModuleId.createIdFromTxt( metaInfo.__buildId ), importParam.importedAliasMap)
   self.importModule2ModuleInfo[moduleTypeInfo] = moduleInfo
   self.importModuleName2ModuleInfo[modulePath] = moduleInfo
   
   self.importModuleInfo:remove(  )
   
   Log.log( Log.Level.Info, __func__, 1158, function (  )
   
      return string.format( "%s complete", orgModulePath)
   end )
   
   
   return metaInfo, typeId2TypeInfo, moduleInfo
end


return _moduleObj
