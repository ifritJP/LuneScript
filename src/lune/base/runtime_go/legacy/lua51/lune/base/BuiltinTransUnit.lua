--lune/base/BuiltinTransUnit.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@BuiltinTransUnit'
local _lune = {}
if _lune6 then
   _lune = _lune6
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

if not _lune6 then
   _lune6 = _lune
end
local Types = _lune.loadModule( 'lune.base.Types' )
local Parser = _lune.loadModule( 'lune.base.Parser' )
local Util = _lune.loadModule( 'lune.base.Util' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local TransUnitIF = _lune.loadModule( 'lune.base.TransUnitIF' )

local TransUnit = {}
_moduleObj.TransUnit = TransUnit
function TransUnit:get_scope(  )

   return self.scope
end
function TransUnit._new( ctrl_info, processInfo )
   local obj = {}
   TransUnit._setmeta( obj )
   if obj.__init then obj:__init( ctrl_info, processInfo ); end
   return obj
end
function TransUnit:__init(ctrl_info, processInfo) 
   self.namespace2Scope = {}
   self.processInfo = processInfo
   self.scope = processInfo:get_topScope()
end
function TransUnit:error( mess )

   Util.err( mess )
end
function TransUnit:pushScope( scopeKind, baseInfo, interfaceList )

   self.scope = Ast.TypeInfo.createScope( self.processInfo, self.scope, scopeKind, baseInfo, interfaceList )
   return self.scope
end
function TransUnit:popScope(  )

   self.scope = self.scope:get_outerScope()
end
function TransUnit:getCurrentNamespaceTypeInfo(  )

   return Ast.getBuiltinMut( self.scope:getNamespaceTypeInfo(  ) )
end
function TransUnit:pushModuleLow( processInfo, externalFlag, name, mutable )

   local typeInfo = Ast.headTypeInfo
   
   local modName
   
   if name:find( "^@" ) then
      modName = name
   else
    
      modName = string.format( "@%s", name)
   end
   
   
   do
      local _exp = self.scope:getTypeInfoChild( modName )
      if _exp ~= nil then
         typeInfo = _exp
         do
            local scope = self.namespace2Scope[typeInfo]
            if scope ~= nil then
               self.scope = scope
            else
               self:error( string.format( "not found scope -- %s", name) )
            end
         end
         
      else
         local _
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         local parentScope = self.scope
         local scope = self:pushScope( Ast.ScopeKind.Module )
         local newType = processInfo:createModule( scope, parentInfo, parentInfo, externalFlag, modName, mutable )
         typeInfo = newType
         self.namespace2Scope[typeInfo] = scope
         Ast.addBuiltinMut( newType, scope )
         
         local _1, existSym = parentScope:addClass( processInfo, modName, nil, typeInfo )
         if existSym ~= nil then
            self:error( string.format( "module symbols exist -- %s.%s -- %s.%s", existSym:get_namespaceTypeInfo():getTxt(  ), existSym:get_name(), parentInfo:getTxt(  ), modName) )
         end
         
      end
   end
   
   return typeInfo
end
function TransUnit:popModule(  )

   self:popScope(  )
end
function TransUnit:pushClassScope( errPos, classTypeInfo, scope )

   if self.scope ~= _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ) then
      local classParentName
      
      local classParentTypeId
      
      do
         local parentType = _lune.nilacc( _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ), 'get_ownerTypeInfo', 'callmtd' )
         if parentType ~= nil then
            classParentName = parentType:getTxt(  )
            classParentTypeId = parentType:get_typeId()
         else
            classParentName = "nil"
            classParentTypeId = Ast.dummyIdInfo
         end
      end
      
      self:error( string.format( "This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", classTypeInfo:getTxt(  ), tostring( _lune.nilacc( self.scope:get_ownerTypeInfo(), 'getTxt', 'callmtd'  )), _lune.nilacc( _lune.nilacc( self.scope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1, classParentName, classParentTypeId.id) )
   end
   
   self.scope = scope
end
function TransUnit:pushClassLow( processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   local typeInfo
   
   do
      local _exp = self.scope:getTypeInfo( name, self.scope, true, Ast.ScopeAccess.Normal )
      if _exp ~= nil then
         
         typeInfo = _exp
         
         if typeInfo:get_abstractFlag() ~= abstractFlag then
            self:error( string.format( "mismatch class(%s) abstract for prototpye", typeInfo:getTxt(  )) )
         end
         
         if typeInfo:get_accessMode() ~= accessMode then
            self:error( string.format( "mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", typeInfo:getTxt(  ), Ast.AccessMode:_getTxt( accessMode)
            , Ast.AccessMode:_getTxt( typeInfo:get_accessMode())
            ) )
         end
         
         if baseInfo ~= nil then
            if typeInfo:get_baseTypeInfo() ~= baseInfo then
               self:error( string.format( "mismatch class(%s) base class(%s) for prototpye base class(%s)", typeInfo:getTxt(  ), baseInfo:getTxt(  ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         else
            if typeInfo:hasBase(  ) then
               self:error( string.format( "mismatch class(%s) base class(None) for prototpye base class(%s)", typeInfo:getTxt(  ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         end
         
         
         local function compareList( protoList, typeList, message )
         
            if #protoList == #typeList then
               for index, protoType in ipairs( protoList ) do
                  if protoType ~= typeList[index] then
                     self:error( string.format( "mismatch class(%s) %s(%s) for prototpye %s(%s)", typeInfo:getTxt(  ), message, typeList[index]:getTxt(  ), message, protoType:getTxt(  )) )
                  end
                  
               end
               
            else
             
               self:error( string.format( "mismatch class(%s) %s(%d) for prototpye %s(%d)", typeInfo:getTxt(  ), message, #typeList, message, #protoList) )
            end
            
         end
         
         compareList( typeInfo:get_interfaceList(), _lune.unwrapDefault( interfaceList, {}), "interface" )
         
         compareList( typeInfo:get_itemTypeInfoList(), _lune.unwrapDefault( genTypeList, {}), "generics" )
         
         do
            local scope = self.namespace2Scope[typeInfo]
            if scope ~= nil then
               self.scope = scope
            else
               do
                  local scope = Ast.TypeInfo.getBuiltinInfo( typeInfo ):get_scope()
                  if scope ~= nil then
                     self.scope = scope
                  else
                     self:error( string.format( "not find scope -- %s", name) )
                  end
               end
               
            end
         end
         
         do
            local _switchExp = (typeInfo:get_kind() )
            if _switchExp == Ast.TypeInfoKind.Class then
               if mode == TransUnitIF.DeclClassMode.Interface then
                  self:error( string.format( "define interface already -- %s", name) )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.IF then
               if mode ~= TransUnitIF.DeclClassMode.Interface then
                  self:error( string.format( "define class already -- %s", name) )
               end
               
            end
         end
         
      else
         local parentInfo = self:getCurrentNamespaceTypeInfo(  )
         
         local parentScope = self.scope
         local scope = self:pushScope( Ast.ScopeKind.Class, baseInfo, interfaceList )
         local workGenTypeList
         
         if genTypeList ~= nil then
            workGenTypeList = genTypeList
         else
            workGenTypeList = {}
         end
         
         
         local newType = processInfo:createClassAsync( mode ~= TransUnitIF.DeclClassMode.Interface, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, parentInfo, externalFlag, accessMode, name )
         typeInfo = newType
         self.namespace2Scope[typeInfo] = scope
         Ast.addBuiltinMut( newType, scope )
         
         parentScope:addClassLazy( processInfo, name, errPos, typeInfo, mode == TransUnitIF.DeclClassMode.LazyModule )
      end
   end
   
   if genTypeList ~= nil then
      for __index, genType in ipairs( genTypeList ) do
         self.scope:addAlternate( processInfo, accessMode, genType:get_txt(), errPos, genType )
      end
      
   end
   
   return typeInfo
end
function TransUnit:popClass(  )

   self:popScope(  )
end
function TransUnit._setmeta( obj )
  setmetatable( obj, { __index = TransUnit  } )
end


return _moduleObj
