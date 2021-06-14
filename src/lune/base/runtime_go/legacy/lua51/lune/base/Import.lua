--lune/base/Import.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@Import'
local _lune = {}
if _lune4 then
   _lune = _lune4
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

if not _lune4 then
   _lune4 = _lune
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
local Runner = _lune.loadModule( 'lune.base.Runner' )

local TransUnitIF = _lune.loadModule( 'lune.base.TransUnitIF' )



local ModuleLoader = {}

local _TypeInfo = {}

local ImportParam = {}
function ImportParam.setmeta( obj )
  setmetatable( obj, { __index = ImportParam  } )
end
function ImportParam.new( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, typeId2TypeInfoMut, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap, dependLibId2DependInfo )
   local obj = {}
   ImportParam.setmeta( obj )
   if obj.__init then
      obj:__init( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, typeId2TypeInfoMut, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap, dependLibId2DependInfo )
   end
   return obj
end
function ImportParam:__init( pos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, typeId2TypeInfoMut, importedAliasMap, lazyModuleSet, metaInfo, scope, moduleTypeInfo, scopeAccess, typeId2AtomMap, dependLibId2DependInfo )

   self.pos = pos
   self.modifier = modifier
   self.processInfo = processInfo
   self.typeId2Scope = typeId2Scope
   self.typeId2TypeInfo = typeId2TypeInfo
   self.typeId2TypeInfoMut = typeId2TypeInfoMut
   self.importedAliasMap = importedAliasMap
   self.lazyModuleSet = lazyModuleSet
   self.metaInfo = metaInfo
   self.scope = scope
   self.moduleTypeInfo = moduleTypeInfo
   self.scopeAccess = scopeAccess
   self.typeId2AtomMap = typeId2AtomMap
   self.dependLibId2DependInfo = dependLibId2DependInfo
end


setmetatable( _TypeInfo, { ifList = {Mapping,} } )
function _TypeInfo.new(  )
   local obj = {}
   _TypeInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function _TypeInfo:__init() 
   self.typeId = Ast.userRootId
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
      typeInfo:get_typeId():set_orgId( self.typeId )
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


function ImportParam:getTypeInfoMut( typeId )

   local typeInfo, mess = self:getTypeInfo( typeId )
   if typeInfo ~= nil then
      local typeInfoMut = self.typeId2TypeInfoMut[typeId]
      if  nil == typeInfoMut then
         local _typeInfoMut = typeInfoMut
      
         Util.err( string.format( "not found TypeInfoMut for %d: %s", typeId, typeInfo:getTxt(  )) )
      end
      
      return typeInfoMut
   end
   
   Util.err( string.format( "not found TypeInfo for %d: %s", typeId, mess or "") )
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

   if typeId.mod == 0 then
      return self:getTypeInfo( typeId.id )
   end
   
   if typeId.mod == frontInterface.getRootDependModId(  ) then
      return Ast.getRootProcessInfoRo(  ):getTypeInfo( typeId.id ), nil
   end
   
   local exportInfo = self.dependLibId2DependInfo[typeId.mod]
   if  nil == exportInfo then
      local _exportInfo = exportInfo
   
      Util.err( string.format( "%s, %d, %d", self.moduleTypeInfo:getTxt(  ), typeId.mod, typeId.id) )
   end
   
   do
      local typeInfo = exportInfo:get_importId2localTypeInfoMap()[typeId.id]
      if typeInfo ~= nil then
         return typeInfo, nil
      end
   end
   
   do
      local typeInfo = exportInfo:get_processInfo():getTypeInfo( typeId.id )
      if typeInfo ~= nil then
         return typeInfo, nil
      end
   end
   
   return nil, string.format( "not found type -- %s, %d, %d", self.moduleTypeInfo:getTxt(  ), typeId.mod, typeId.id)
end


local _TypeInfoNilable = {}
setmetatable( _TypeInfoNilable, { __index = _TypeInfo } )
function _TypeInfoNilable:createTypeInfo( param )

   local orgTypeInfo = param:getTypeInfoFrom( self.orgTypeId )
   if  nil == orgTypeInfo then
      local _orgTypeInfo = orgTypeInfo
   
      Util.err( string.format( "failed to createTypeInfo -- self.orgTypeId = (%d,%d)", self.orgTypeId.mod, self.orgTypeId.id) )
   end
   
   local newTypeInfo = orgTypeInfo:get_nilableTypeInfo(  )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   table.insert( memInfo, { name = "orgTypeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
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
   param.typeId2TypeInfoMut[self.typeId] = newTypeInfo
   
   newTypeInfo:get_typeId():set_orgId( self.typeId )
   
   local _1 = param:getTypeInfo( self.parentId )
   if  nil == _1 then
      local __1 = _1
   
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
   
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   param.typeId2TypeInfoMut[self.typeId] = newTypeInfo
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   
   local newTypeInfo, scope = param.processInfo:createGeneric( genSrcTypeInfo, genTypeList, param.moduleTypeInfo )
   param.typeId2TypeInfo[self.typeId] = newTypeInfo
   param.typeId2TypeInfoMut[self.typeId] = newTypeInfo
   newTypeInfo:get_typeId():set_orgId( self.typeId )
   param.typeId2Scope[self.typeId] = scope
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
   
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   newTypeInfo:get_typeId():set_orgId( self.typeId )
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
   if self.parentId ~= Ast.userRootId then
      
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
         
         error( "internal error" )
      else
         local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
         
         local mutable = false
         do
            if self.typeId == param.metaInfo.__moduleTypeId then
               mutable = param.metaInfo.__moduleMutable
            end
            
         end
         
         local parentInfoMut
         
         
         parentInfoMut = param:getTypeInfoMut( parentInfo:get_typeId().id )
         local workTypeInfo = param.processInfo:createModule( scope, parentInfoMut, true, self.txt, mutable )
         
         newTypeInfo = workTypeInfo
         param.typeId2Scope[self.typeId] = scope
         param.typeId2TypeInfo[self.typeId] = workTypeInfo
         workTypeInfo:get_typeId():set_orgId( self.typeId )
         parentScope:addClass( param.processInfo, self.txt, nil, workTypeInfo )
         
         Log.log( Log.Level.Info, __func__, 376, function (  )
         
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
   if self.parentId ~= Ast.userRootId or not Ast.getBuiltInTypeIdMap(  )[self.typeId] or self.kind == Ast.TypeInfoKind.List or self.kind == Ast.TypeInfoKind.Array or self.kind == Ast.TypeInfoKind.Map or self.kind == Ast.TypeInfoKind.Set then
      local parentInfo = Ast.headTypeInfo
      if self.parentId ~= Ast.userRootId then
         
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
         error( "internal error" )
      else
       
         local function postProcess( workTypeInfo, scope )
         
            newTypeInfo = workTypeInfo
            
            if scope ~= nil then
               param.typeId2Scope[self.typeId] = scope
            end
            
            param.typeId2TypeInfo[self.typeId] = workTypeInfo
            
            workTypeInfo:get_typeId():set_orgId( self.typeId )
         end
         
         do
            local _switchExp = self.kind
            if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.IF then
               Log.log( Log.Level.Debug, __func__, 486, function (  )
               
                  return string.format( "new type -- %d, %s -- %s, %d", self.parentId, self.txt, _lune.nilacc( parentScope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , Ast.defaultTypeNameCtrl, parentScope, false ) or "nil", _lune.nilacc( _lune.nilacc( parentScope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1)
               end )
               
               
               local baseScope = _lune.unwrap( baseInfo:get_scope())
               local ifScopeList = {}
               for __index, ifType in ipairs( interfaceList ) do
                  table.insert( ifScopeList, _lune.unwrap( ifType:get_scope()) )
               end
               
               
               local scope = Ast.Scope.new(param.processInfo, parentScope, true, baseScope, ifScopeList)
               
               local altTypeList = {}
               for __index, itemType in ipairs( itemTypeInfo ) do
                  table.insert( altTypeList, _lune.unwrap( (_lune.__Cast( itemType, 3, Ast.AlternateTypeInfo ) )) )
               end
               
               
               local parentInfoMut = param:getTypeInfoMut( self.parentId )
               local workTypeInfo = param.processInfo:createClassAsync( self.kind == Ast.TypeInfoKind.Class, self.abstractFlag, scope, baseInfo, interfaceList, altTypeList, parentInfoMut, true, Ast.AccessMode.Pub, self.txt )
               parentScope:addClassLazy( param.processInfo, self.txt, nil, workTypeInfo, _lune._Set_has(param.lazyModuleSet, self.typeId ) )
               
               postProcess( workTypeInfo, scope )
               
               param.typeId2TypeInfoMut[self.typeId] = workTypeInfo
            elseif _switchExp == Ast.TypeInfoKind.ExtModule then
               Log.log( Log.Level.Debug, __func__, 523, function (  )
               
                  return string.format( "new type -- %d, %s -- %s, %d", self.parentId, self.txt, _lune.nilacc( parentScope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , Ast.defaultTypeNameCtrl, parentScope, false ) or "nil", _lune.nilacc( _lune.nilacc( parentScope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1)
               end )
               
               
               local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil, {})
               
               local parentInfoMut = param:getTypeInfoMut( self.parentId )
               local workTypeInfo = param.processInfo:createExtModule( scope, parentInfoMut, true, Ast.AccessMode.Pub, self.txt, _lune.unwrap( self.moduleLang), _lune.unwrap( self.requirePath) )
               parentScope:addExtModule( param.processInfo, self.txt, nil, workTypeInfo, _lune._Set_has(param.lazyModuleSet, self.typeId ), _lune.unwrap( self.moduleLang) )
               
               postProcess( workTypeInfo, scope )
               
               param.typeId2TypeInfoMut[self.typeId] = workTypeInfo
            elseif _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.FormFunc or _switchExp == Ast.TypeInfoKind.Macro then
               local typeInfoKind = self.kind
               local accessMode = self.accessMode
               
               local workTypeInfo
               
               local scope = nil
               if self.kind ~= Ast.TypeInfoKind.FormFunc then
                  scope = Ast.Scope.new(param.processInfo, parentScope, false, nil)
               end
               
               
               local parentInfoMut = param:getTypeInfoMut( self.parentId )
               local workTypeInfoMut = param.processInfo:createFuncAsync( self.abstractFlag, false, scope, typeInfoKind, parentInfoMut, false, true, self.staticFlag, accessMode, self.txt, self.asyncMode, itemTypeInfo, argTypeInfo, retTypeInfo, Ast.isMutable( self.mutMode ) )
               param.typeId2TypeInfoMut[self.typeId] = workTypeInfoMut
               
               postProcess( workTypeInfoMut, scope )
               
               do
                  local _switchExp = self.kind
                  if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Method or _switchExp == Ast.TypeInfoKind.Macro or _switchExp == Ast.TypeInfoKind.FormFunc then
                     local symbolKind = Ast.SymbolKind.Fun
                     do
                        local _switchExp = self.kind
                        if _switchExp == Ast.TypeInfoKind.Method then
                           symbolKind = Ast.SymbolKind.Mtd
                        elseif _switchExp == Ast.TypeInfoKind.Macro then
                           symbolKind = Ast.SymbolKind.Mac
                        elseif _switchExp == Ast.TypeInfoKind.FormFunc then
                           symbolKind = Ast.SymbolKind.Typ
                        end
                     end
                     
                     local workParentScope = _lune.unwrap( param.typeId2Scope[self.parentId])
                     workParentScope:add( param.processInfo, symbolKind, false, self.kind == Ast.TypeInfoKind.Func, self.txt, nil, workTypeInfoMut, accessMode, self.staticFlag, Ast.MutMode.IMut, true, false )
                  end
               end
               
            elseif _switchExp == Ast.TypeInfoKind.Set then
               local workTypeInfo = param.processInfo:createSet( self.accessMode, parentInfo, itemTypeInfo, self.mutMode )
               postProcess( workTypeInfo, nil )
            elseif _switchExp == Ast.TypeInfoKind.List then
               local workTypeInfo = param.processInfo:createList( self.accessMode, parentInfo, itemTypeInfo, self.mutMode )
               postProcess( workTypeInfo, nil )
            elseif _switchExp == Ast.TypeInfoKind.Array then
               local workTypeInfo = param.processInfo:createArray( self.accessMode, parentInfo, itemTypeInfo, self.mutMode )
               postProcess( workTypeInfo, nil )
            elseif _switchExp == Ast.TypeInfoKind.Map then
               local workTypeInfo = param.processInfo:createMap( self.accessMode, parentInfo, itemTypeInfo[1], itemTypeInfo[2], self.mutMode )
               postProcess( workTypeInfo, nil )
            else 
               
                  Util.err( string.format( "illegal kind -- %s", Ast.TypeInfoKind:_getTxt( self.kind)
                  ) )
            end
         end
         
      end
      
   else
    
      newTypeInfo = param.scope:getTypeInfo( self.txt, param.scope, false, param.scopeAccess )
      if newTypeInfo ~= nil then
         param.typeId2TypeInfo[self.typeId] = newTypeInfo
         newTypeInfo:get_typeId():set_orgId( self.typeId )
      else
         for key, val in pairs( self:_toMap(  ) ) do
            Util.errorLog( string.format( "error: illegal self %s:%s", key, tostring( val)) )
         end
         
      end
      
   end
   
   return newTypeInfo, nil
end
function _TypeInfoNormal.setmeta( obj )
  setmetatable( obj, { __index = _TypeInfoNormal  } )
end
function _TypeInfoNormal.new( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, asyncMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )
   local obj = {}
   _TypeInfoNormal.setmeta( obj )
   if obj.__init then
      obj:__init( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, asyncMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )
   end
   return obj
end
function _TypeInfoNormal:__init( parentId, abstractFlag, baseId, txt, staticFlag, accessMode, kind, mutMode, asyncMode, ifList, itemTypeId, argTypeId, retTypeId, children, moduleLang, requirePath )

   _TypeInfo.__init( self)
   self.parentId = parentId
   self.abstractFlag = abstractFlag
   self.baseId = baseId
   self.txt = txt
   self.staticFlag = staticFlag
   self.accessMode = accessMode
   self.kind = kind
   self.mutMode = mutMode
   self.asyncMode = asyncMode
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
   table.insert( memInfo, { name = "asyncMode", func = Ast.Async._from, nilable = false, child = {} } )
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
   
   local parentInfo = param:getTypeInfoMut( self.parentId )
   local parentScope = _lune.unwrap( param.typeId2Scope[self.parentId])
   local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
   
   param.typeId2Scope[self.typeId] = scope
   local valTypeInfo = _lune.unwrap( param:getTypeInfo( self.valTypeId ))
   local enumTypeInfo = param.processInfo:createEnum( scope, parentInfo, true, accessMode, self.txt, valTypeInfo )
   local newTypeInfo = enumTypeInfo
   param.typeId2TypeInfo[self.typeId] = enumTypeInfo
   param.typeId2TypeInfoMut[self.typeId] = enumTypeInfo
   enumTypeInfo:get_typeId():set_orgId( self.typeId )
   
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
   
   local parentInfo = param:getTypeInfoMut( self.parentId )
   
   local parentScope = _lune.unwrap( param.typeId2Scope[self.parentId])
   local scope = Ast.Scope.new(param.processInfo, parentScope, true, nil)
   
   param.typeId2Scope[self.typeId] = scope
   local algeTypeInfo = param.processInfo:createAlge( scope, parentInfo, true, accessMode, self.txt )
   local newTypeInfo = algeTypeInfo
   param.typeId2TypeInfo[self.typeId] = algeTypeInfo
   param.typeId2TypeInfoMut[self.typeId] = algeTypeInfo
   
   algeTypeInfo:get_typeId():set_orgId( self.typeId )
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


local ModuleLoaderParam = {}
_moduleObj.ModuleLoaderParam = ModuleLoaderParam
function ModuleLoaderParam.setmeta( obj )
  setmetatable( obj, { __index = ModuleLoaderParam  } )
end
function ModuleLoaderParam.new( ctrl_info, processInfo, latestPos, macroMode, nearCode, validMutControl, macroEval )
   local obj = {}
   ModuleLoaderParam.setmeta( obj )
   if obj.__init then
      obj:__init( ctrl_info, processInfo, latestPos, macroMode, nearCode, validMutControl, macroEval )
   end
   return obj
end
function ModuleLoaderParam:__init( ctrl_info, processInfo, latestPos, macroMode, nearCode, validMutControl, macroEval )

   self.ctrl_info = ctrl_info
   self.processInfo = processInfo
   self.latestPos = latestPos
   self.macroMode = macroMode
   self.nearCode = nearCode
   self.validMutControl = validMutControl
   self.macroEval = macroEval
end
function ModuleLoaderParam:get_ctrl_info()
   return self.ctrl_info
end
function ModuleLoaderParam:get_processInfo()
   return self.processInfo
end
function ModuleLoaderParam:get_latestPos()
   return self.latestPos
end
function ModuleLoaderParam:get_macroMode()
   return self.macroMode
end
function ModuleLoaderParam:get_nearCode()
   return self.nearCode
end
function ModuleLoaderParam:get_validMutControl()
   return self.validMutControl
end
function ModuleLoaderParam:get_macroEval()
   return self.macroEval
end



local ModuleLoaderResult = {}
function ModuleLoaderResult.setmeta( obj )
  setmetatable( obj, { __index = ModuleLoaderResult  } )
end
function ModuleLoaderResult.new( exportInfo, modulePath, fullModulePath, baseDir, err, depth, importedAliasMap )
   local obj = {}
   ModuleLoaderResult.setmeta( obj )
   if obj.__init then
      obj:__init( exportInfo, modulePath, fullModulePath, baseDir, err, depth, importedAliasMap )
   end
   return obj
end
function ModuleLoaderResult:__init( exportInfo, modulePath, fullModulePath, baseDir, err, depth, importedAliasMap )

   self.exportInfo = exportInfo
   self.modulePath = modulePath
   self.fullModulePath = fullModulePath
   self.baseDir = baseDir
   self.err = err
   self.depth = depth
   self.importedAliasMap = importedAliasMap
end
function ModuleLoaderResult:get_exportInfo()
   return self.exportInfo
end


setmetatable( ModuleLoader, { __index = Runner.Runner,ifList = {frontInterface.ModuleLoader,} } )
_moduleObj.ModuleLoader = ModuleLoader
function ModuleLoader:applyExportInfo( exportInfo )

   if exportInfo ~= nil then
      do
         local work = _lune.__Cast( exportInfo, 3, Nodes.ExportInfo )
         if work ~= nil then
            self.macroCtrl:importMacroInfo( work:get_typeId2DefMacroInfo() )
         end
      end
      
      for key, val in pairs( exportInfo:get_importedAliasMap() ) do
         self.result.importedAliasMap[key] = val
      end
      
   end
   
end
function ModuleLoader:craeteModuleInfo( moduleMeta )

   do
      local _matchExp = moduleMeta:get_metaOrModule()
      if _matchExp[1] == frontInterface.MetaOrModule.Module[1] then
         local moduleInfo = _matchExp[2][1]
      
         self:applyExportInfo( moduleInfo:get_exportInfo() )
         return moduleInfo:get_exportInfo()
      elseif _matchExp[1] == frontInterface.MetaOrModule.Export[1] then
         local exportInfo = _matchExp[2][1]
      
         self:applyExportInfo( exportInfo )
         return exportInfo
      elseif _matchExp[1] == frontInterface.MetaOrModule.MetaRaw[1] then
         local metaInfo = _matchExp[2][1]
      
         self.importModuleInfo:add( self.result.modulePath )
         self.importProcessInfo:switchIdProvier( Ast.IdType.Ext )
         
         local nameList = Util.splitStr( self.result.modulePath, '[^%./:]+' )
         
         local moduleInfo = self:processImportFromFile( self.importProcessInfo, moduleMeta:get_lnsPath(), metaInfo, self.result.fullModulePath, self.result.modulePath, nameList, self.result.baseDir, self.result.depth )
         
         self.importProcessInfo:switchIdProvier( Ast.IdType.Base )
         self.importModuleInfo:remove(  )
         
         moduleMeta:set_metaOrModule( _lune.newAlge( frontInterface.MetaOrModule.Module, {moduleInfo}) )
         
         return moduleInfo:get_exportInfo()
      end
   end
   
end
function ModuleLoader.new( exportInfo, workImportModuleInfo, modulePath, fullModulePath, baseDir, moduleLoaderParam, depth )
   local obj = {}
   ModuleLoader.setmeta( obj )
   if obj.__init then obj:__init( exportInfo, workImportModuleInfo, modulePath, fullModulePath, baseDir, moduleLoaderParam, depth ); end
   return obj
end
function ModuleLoader:__init(exportInfo, workImportModuleInfo, modulePath, fullModulePath, baseDir, moduleLoaderParam, depth) 
   Runner.Runner.__init( self)
   
   
   self.moduleLoaderParam = moduleLoaderParam
   self.result = ModuleLoaderResult.new(exportInfo, modulePath, fullModulePath, baseDir, "", depth, {})
   
   self.moduleMeta = nil
   self.validMutControl = moduleLoaderParam:get_validMutControl()
   self.curPos = moduleLoaderParam:get_latestPos()
   self.macroCtrl = Macro.MacroCtrl.new(moduleLoaderParam:get_macroEval())
   self.importModuleInfo = workImportModuleInfo:clone(  )
   
   self.importProcessInfo = moduleLoaderParam:get_processInfo():newUser(  )
   
   local simpleTransUnit = TransUnitIF.SimpeTransUnit.new(moduleLoaderParam:get_ctrl_info(), self.importProcessInfo, moduleLoaderParam:get_latestPos(), moduleLoaderParam:get_macroMode(), moduleLoaderParam:get_nearCode())
   
   self.transUnitIF = simpleTransUnit
   self.globalScope = simpleTransUnit:get_globalScope()
   
   self.loaderFunc = function (  )
   
      if not self.result.exportInfo then
         if not self.importModuleInfo:add( fullModulePath ) then
            self.result.err = string.format( "recursive import: %s -> %s", self.importModuleInfo:getFull(  ), fullModulePath)
         else
          
            do
               do
                  local _exp = frontInterface.loadMeta( self.importModuleInfo:clone(  ), modulePath, fullModulePath, baseDir, self )
                  if _exp ~= nil then
                     local moduleMeta = _exp
                     self.result.exportInfo = self:craeteModuleInfo( moduleMeta )
                  else
                     self.result.err = string.format( "failed to load meta -- %s on %s", fullModulePath, baseDir or "./")
                  end
               end
               
            end
            
            self.importModuleInfo:remove(  )
         end
         
      end
      
   end
   self:start( 0, string.format( "ModuleLoader - %s", fullModulePath) )
end
function ModuleLoader:runMain(  )

   self.loaderFunc(  )
end
function ModuleLoader:getResult(  )

   
   return self.result
end
function ModuleLoader:getExportInfo(  )

   
   return self.result:get_exportInfo()
end
function ModuleLoader.setmeta( obj )
  setmetatable( obj, { __index = ModuleLoader  } )
end


function ModuleLoader:processImportFromFile( processInfo, lnsPath, metaInfoStem, fullModulePath, modulePath, nameList, baseDir, depth )
   local __func__ = '@lune.@base.@Import.ModuleLoader.processImportFromFile'

   local moduleInfo
   
   do
      local metaInfo = metaInfoStem
      Log.log( Log.Level.Info, __func__, 943, function (  )
      
         return string.format( "%s processing", fullModulePath)
      end )
      
      
      local dependLibId2DependInfo = {}
      do
         local loaderMap = {}
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
                  
                  local workProcessInfo = processInfo:newUser(  )
                  
                  local moduleLoader = self:processImportMain( workProcessInfo, baseDir, dependName, depth + 1 )
                  
                  local typeId = math.floor((_lune.unwrap( dependInfo['typeId']) ))
                  loaderMap[moduleLoader] = typeId
               end
            end
         end
         
         for moduleLoader, typeId in pairs( loaderMap ) do
            local result = moduleLoader:getResult(  )
            do
               local _exp = result.exportInfo
               if _exp ~= nil then
                  self:applyExportInfo( _exp )
                  dependLibId2DependInfo[typeId] = _exp
               else
                  self.transUnitIF:error( result.err )
               end
            end
            
            
         end
         
      end
      
      
      local typeId2TypeInfo = {}
      local typeId2TypeInfoMut = {}
      
      typeId2TypeInfo[Ast.userRootId] = processInfo:get_dummyParentType()
      local typeId2Scope = {}
      typeId2Scope[Ast.userRootId] = processInfo:get_topScope()
      
      for typeId, dependIdInfo in pairs( metaInfo.__dependIdMap ) do
         local dependInfo = _lune.unwrap( dependLibId2DependInfo[_lune.unwrap( dependIdInfo[1])])
         local typeInfo = _lune.unwrap( dependInfo:getTypeInfo( _lune.unwrap( dependIdInfo[2]) ))
         typeId2TypeInfo[typeId] = typeInfo
      end
      
      
      local moduleTypeInfo = Ast.headTypeInfo
      for index, moduleName in ipairs( nameList ) do
         local mutable = false
         if index == #nameList then
            mutable = metaInfo.__moduleMutable
         end
         
         local nsInfo = self.transUnitIF:pushModule( processInfo, true, moduleName, mutable )
         moduleTypeInfo = nsInfo:get_typeInfo()
         local typeId = _lune.unwrap( metaInfo.__moduleHierarchy[#nameList - index + 1])
         typeId2TypeInfo[typeId] = moduleTypeInfo
         typeId2TypeInfoMut[typeId] = nsInfo:get_typeInfo()
         typeId2Scope[typeId] = self.transUnitIF:get_scope()
      end
      
      for __index, _1 in ipairs( nameList ) do
         self.transUnitIF:popModule(  )
      end
      
      
      for __index, symbolInfo in pairs( Ast.getSym2builtInTypeMap(  ) ) do
         typeId2TypeInfo[symbolInfo:get_typeInfo():get_typeId(  ).id] = symbolInfo:get_typeInfo()
      end
      
      for __index, builtinTypeInfo in pairs( Ast.getBuiltInTypeIdMap(  ) ) do
         local typeInfo = builtinTypeInfo:get_typeInfo()
         typeId2TypeInfo[typeInfo:get_typeId().id] = typeInfo
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
      
      
      local modifier = TransUnitIF.Modifier.new(self.validMutControl, processInfo)
      
      local importParam = ImportParam.new(self.curPos, modifier, processInfo, typeId2Scope, typeId2TypeInfo, typeId2TypeInfoMut, {}, lazyModuleSet, metaInfo, self.transUnitIF:get_scope(), moduleTypeInfo, Ast.ScopeAccess.Normal, id2atomMap, dependLibId2DependInfo)
      
      for __index, atomInfo in ipairs( _typeInfoList ) do
         local newTypeInfo, errMess = atomInfo:createTypeInfoCache( importParam )
         do
            local _exp = errMess
            if _exp ~= nil then
               Util.err( string.format( "Failed to createType -- %s: %s(%d): %s", fullModulePath, Ast.SerializeKind:_getTxt( atomInfo.skind)
               , atomInfo.typeId, _exp) )
            end
         end
         
         if newTypeInfo ~= nil then
            if newTypeInfo:get_kind() == Ast.TypeInfoKind.Macro then
               orgId2MacroTypeInfo[atomInfo.typeId] = newTypeInfo
            end
            
            if newTypeInfo:get_kind() == Ast.TypeInfoKind.Set then
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
            importParam:getTypeInfo( atomInfo.typeId )
            local scope = _lune.unwrap( typeId2Scope[atomInfo.typeId])
            for __index, childId in ipairs( atomInfo.children ) do
               local typeInfo = importParam:getTypeInfoFrom( childId )
               if  nil == typeInfo then
                  local _typeInfo = typeInfo
               
                  Util.err( string.format( "not found childId -- %s, %d, %s(%d)", fullModulePath, childId.id, atomInfo.txt, atomInfo.typeId) )
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
         local __func__ = '@lune.@base.@Import.ModuleLoader.processImportFromFile.registMember'
      
         local skip = false
         do
            if metaInfo.__dependIdMap[classTypeId] then
               skip = true
            end
            
         end
         
         if skip then
            return 
         end
         
         
         do
            local classTypeInfo = _lune.unwrap( typeId2TypeInfo[classTypeId])
            
            do
               local _switchExp = (classTypeInfo:get_kind() )
               if _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.ExtModule then
                  local scope = _lune.unwrap( typeId2Scope[classTypeId])
                  self.transUnitIF:pushClassScope( self.curPos, classTypeInfo, scope )
                  
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
                              local typeId = _IdInfo._fromStem( (fieldInfo['typeId'] ) )
                              if typeId ~= nil then
                                 local fieldTypeInfo = _lune.unwrap( importParam:getTypeInfoFrom( typeId ))
                                 local symbolInfo = self.transUnitIF:get_scope():addMember( processInfo, fieldName, nil, fieldTypeInfo, _lune.unwrap( Ast.AccessMode._from( math.floor((_lune.unwrap( fieldInfo['accessMode']) )) )), fieldInfo['staticFlag'] and true or false, _lune.unwrap( Ast.MutMode._from( math.floor((_lune.unwrap( fieldInfo['mutMode']) )) )) )
                              else
                                 self.transUnitIF:error( "not found fieldInfo.typeId" )
                              end
                           end
                           
                        end
                        
                     else
                        self.transUnitIF:error( string.format( "not found class -- %s: %d, %s", fullModulePath, classTypeId, classTypeInfo:getTxt(  )) )
                     end
                  end
                  
               elseif _switchExp == Ast.TypeInfoKind.Module then
                  self.transUnitIF:pushModuleLow( processInfo, true, classTypeInfo:getTxt(  ), Ast.TypeInfo.isMut( classTypeInfo ) )
                  Log.log( Log.Level.Debug, __func__, 1261, function (  )
                  
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
         
      end
      for __index, atomInfo in ipairs( _typeInfoList ) do
         do
            local workInfo = _lune.__Cast( atomInfo, 3, _TypeInfoNormal )
            if workInfo ~= nil then
               if workInfo.parentId == Ast.userRootId then
                  registMember( atomInfo.typeId )
               end
               
            else
               do
                  local workInfo = _lune.__Cast( atomInfo, 3, _TypeInfoModule )
                  if workInfo ~= nil then
                     if workInfo.parentId == Ast.userRootId then
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
         
         self.transUnitIF:pushModuleLow( processInfo, true, moduleName, mutable )
      end
      
      
      local VarNameInfo = {}
      setmetatable( VarNameInfo, { ifList = {Mapping,} } )
      function VarNameInfo.setmeta( obj )
  setmetatable( obj, { __index = VarNameInfo  } )
end
      function VarNameInfo.new( typeId, accessMode, mutable )
   local obj = {}
   VarNameInfo.setmeta( obj )
   if obj.__init then
      obj:__init( typeId, accessMode, mutable )
   end
   return obj
end
function VarNameInfo:__init( typeId, accessMode, mutable )

         self.typeId = typeId
         self.accessMode = accessMode
         self.mutable = mutable
      end
      function VarNameInfo:_toMap()
  return self
end
function VarNameInfo._fromMap( val )
  local obj, mes = VarNameInfo._fromMapSub( {}, val )
  if obj then
     VarNameInfo.setmeta( obj )
  end
  return obj, mes
end
function VarNameInfo._fromStem( val )
  return VarNameInfo._fromMap( val )
end

      function VarNameInfo._fromMapSub( obj, val )
         local memInfo = {}
         table.insert( memInfo, { name = "typeId", func = _IdInfo._fromMap, nilable = false, child = {} } )
         table.insert( memInfo, { name = "accessMode", func = Ast.AccessMode._from, nilable = false, child = {} } )
         table.insert( memInfo, { name = "mutable", func = _lune._toBool, nilable = false, child = {} } )
         local result, mess = _lune._fromMap( obj, val, memInfo )
   if not result then
      return nil, mess
   end
   return obj
end
      
      
      for varName, varInfo in pairs( metaInfo.__varName2InfoMap ) do
         
         do
            local varNameInfo = VarNameInfo._fromStem( (varInfo ) )
            if varNameInfo ~= nil then
               local typeId = varNameInfo.typeId
               local scope
               
               if varNameInfo.accessMode == Ast.AccessMode.Global then
                  scope = self.globalScope
               else
                
                  scope = self.transUnitIF:get_scope()
               end
               
               
               scope:addExportedVar( processInfo, varNameInfo.mutable, varNameInfo.accessMode, varName, nil, _lune.unwrap( importParam:getTypeInfoFrom( typeId )), varNameInfo.mutable and Ast.MutMode.Mut or Ast.MutMode.IMut )
            else
               self.transUnitIF:error( "illegal varInfo.typeId" )
            end
         end
         
      end
      
      
      local importedMacroInfoMap = {}
      for orgTypeId, macroInfoStem in pairs( metaInfo.__macroName2InfoMap ) do
         self.macroCtrl:importMacro( processInfo, lnsPath, (macroInfoStem ), _lune.unwrap( orgId2MacroTypeInfo[orgTypeId]), typeId2TypeInfo, importedMacroInfoMap, baseDir )
      end
      
      
      local globalSymbolList = {}
      
      for __index, symbolInfo in pairs( self.globalScope:get_symbol2SymbolInfoMap() ) do
         if symbolInfo:get_accessMode() == Ast.AccessMode.Global then
            table.insert( globalSymbolList, symbolInfo )
         end
         
      end
      
      
      for __index, _2 in ipairs( nameList ) do
         self.transUnitIF:popModule(  )
      end
      
      
      if depth == 1 then
         
         for key, val in pairs( importParam.importedAliasMap ) do
            self.result.importedAliasMap[key] = val
         end
         
      end
      
      
      local moduleProvideInfo = frontInterface.ModuleProvideInfo.new(_lune.unwrap( typeId2TypeInfo[metaInfo.__moduleTypeId]), _lune.unwrap( Ast.SymbolKind._from( metaInfo.__moduleSymbolKind )), metaInfo.__moduleMutable)
      
      local exportInfo = Nodes.ExportInfo.new(moduleTypeInfo, moduleProvideInfo, processInfo, globalSymbolList, importParam.importedAliasMap, frontInterface.ModuleId.createIdFromTxt( metaInfo.__buildId ), fullModulePath, nameList[#nameList], lnsPath, newId2OldIdMap, importedMacroInfoMap)
      
      moduleInfo = frontInterface.ModuleInfo.new(exportInfo)
   end
   
   return moduleInfo
end


local Import = {}
_moduleObj.Import = Import
function Import.new( curPos, importModuleInfo, moduleType, macroCtrl, typeNameCtrl, importedAliasMap, baseDir, validMutControl )
   local obj = {}
   Import.setmeta( obj )
   if obj.__init then obj:__init( curPos, importModuleInfo, moduleType, macroCtrl, typeNameCtrl, importedAliasMap, baseDir, validMutControl ); end
   return obj
end
function Import:__init(curPos, importModuleInfo, moduleType, macroCtrl, typeNameCtrl, importedAliasMap, baseDir, validMutControl) 
   self.baseDir = baseDir
   self.importModuleInfo = importModuleInfo
   self.moduleType = moduleType
   self.macroCtrl = macroCtrl
   self.typeNameCtrl = typeNameCtrl
   self.importedAliasMap = importedAliasMap
   
   self.importModule2ExportInfo = {}
   self.importModuleName2ModuleInfo = {}
end
function Import.setmeta( obj )
  setmetatable( obj, { __index = Import  } )
end
function Import:get_importModule2ExportInfo()
   return self.importModule2ExportInfo
end


function Import:createModuleLoader( baseDir, modulePath, moduleLoaderParam, depth )
   local __func__ = '@lune.@base.@Import.Import.createModuleLoader'

   local fullModulePath
   
   
   do
      modulePath, baseDir, fullModulePath = frontInterface.getLuaModulePath( modulePath, baseDir )
   end
   
   
   Log.log( Log.Level.Info, __func__, 1459, function (  )
   
      return string.format( "%s -> %s start on %s", self.moduleType:getTxt( self.typeNameCtrl ), fullModulePath, tostring( baseDir))
   end )
   
   
   local exportInfo = self.importModuleName2ModuleInfo[fullModulePath]
   
   if exportInfo ~= nil then
      Log.log( Log.Level.Info, __func__, 1467, function (  )
      
         return string.format( "%s already", fullModulePath)
      end )
      
      
      if depth == 1 then
         self.importModule2ExportInfo[exportInfo:get_moduleTypeInfo()] = exportInfo
      end
      
      
      for key, val in pairs( exportInfo:get_importedAliasMap() ) do
         self.importedAliasMap[key] = val
      end
      
   end
   
   
   return ModuleLoader.new(exportInfo, self.importModuleInfo, modulePath, fullModulePath, baseDir, moduleLoaderParam, depth)
end


function Import:loadModuleInfo( moduleLoader )
   local __func__ = '@lune.@base.@Import.Import.loadModuleInfo'

   local result = moduleLoader:getResult(  )
   
   local exportInfo = result.exportInfo
   if  nil == exportInfo then
      local _exportInfo = exportInfo
   
      return nil, result.err
   end
   
   
   local fullModulePath = result.fullModulePath
   local depth = result.depth
   
   do
      local work = _lune.__Cast( exportInfo, 3, Nodes.ExportInfo )
      if work ~= nil then
         self.macroCtrl:importMacroInfo( work:get_typeId2DefMacroInfo() )
      end
   end
   
   for key, val in pairs( exportInfo:get_importedAliasMap() ) do
      self.importedAliasMap[key] = val
   end
   
   for key, val in pairs( result.importedAliasMap ) do
      self.importedAliasMap[key] = val
   end
   
   
   if depth == 1 then
      self.importModule2ExportInfo[exportInfo:get_moduleTypeInfo()] = exportInfo
   end
   
   self.importModuleName2ModuleInfo[fullModulePath] = exportInfo
   
   Log.log( Log.Level.Info, __func__, 1518, function (  )
   
      return string.format( "%s complete", fullModulePath)
   end )
   
   
   return exportInfo, ""
end

function ModuleLoader:processImportMain( processInfo, baseDir, modulePath, depth )
   local __func__ = '@lune.@base.@Import.ModuleLoader.processImportMain'

   local fullModulePath
   
   modulePath, baseDir, fullModulePath = frontInterface.getLuaModulePath( modulePath, baseDir )
   
   Log.log( Log.Level.Info, __func__, 1532, function (  )
   
      return string.format( "%s -> %s start on %s", self.result.fullModulePath, fullModulePath, tostring( baseDir))
   end )
   
   
   local moduleLoader = ModuleLoader.new(nil, self.importModuleInfo, modulePath, fullModulePath, baseDir, self.moduleLoaderParam, depth)
   return moduleLoader
end


function Import:processImport( modulePath, moduleLoaderParam )

   local moduleLoader = self:createModuleLoader( self.baseDir, modulePath, moduleLoaderParam, 1 )
   
   return moduleLoader
end


return _moduleObj
