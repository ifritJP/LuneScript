--lune/base/TransUnitIF.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@TransUnitIF'
local _lune = {}
if _lune4 then
   _lune = _lune4
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

if not _lune4 then
   _lune4 = _lune
end


local Parser = _lune.loadModule( 'lune.base.Parser' )
local Ast = _lune.loadModule( 'lune.base.Ast' )
local Nodes = _lune.loadModule( 'lune.base.Nodes' )
local Util = _lune.loadModule( 'lune.base.Util' )

local Types = _lune.loadModule( 'lune.base.Types' )

local DeclClassMode = {}
_moduleObj.DeclClassMode = DeclClassMode
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
DeclClassMode.LazyModule = 3
DeclClassMode._val2NameMap[3] = 'LazyModule'
DeclClassMode.__allList[4] = DeclClassMode.LazyModule


local Modifier = {}
_moduleObj.Modifier = Modifier
function Modifier:createModifier( typeInfo, mutMode )

   if not self.validMutControl then
      return typeInfo
   end
   
   return self.processInfo:createModifier( typeInfo, mutMode )
end
function Modifier.setmeta( obj )
  setmetatable( obj, { __index = Modifier  } )
end
function Modifier.new( validMutControl, processInfo )
   local obj = {}
   Modifier.setmeta( obj )
   if obj.__init then
      obj:__init( validMutControl, processInfo )
   end
   return obj
end
function Modifier:__init( validMutControl, processInfo )

   self.validMutControl = validMutControl
   self.processInfo = processInfo
end
function Modifier:set_validMutControl( validMutControl )
   self.validMutControl = validMutControl
end


local IdSetInfo = {}
_moduleObj.IdSetInfo = IdSetInfo
function IdSetInfo.new(  )
   local obj = {}
   IdSetInfo.setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function IdSetInfo:__init() 
   self.anonymousFuncId = Ast.IdProvider.new(0, 10000)
   self.anonymousVarId = Ast.IdProvider.new(0, 10000)
end
function IdSetInfo:registerSym( symbol )

   if symbol:get_kind() == Ast.SymbolKind.Var then
      if symbol:get_name() == "_" then
         local id = self.anonymousVarId:getNewId(  )
         return Ast.AnonymousSymbolInfo.new(symbol, id)
      end
      
   end
   
   return symbol
end
function IdSetInfo.setmeta( obj )
  setmetatable( obj, { __index = IdSetInfo  } )
end


local LockedAsyncInfo = {}
function LockedAsyncInfo.setmeta( obj )
  setmetatable( obj, { __index = LockedAsyncInfo  } )
end
function LockedAsyncInfo.new( loopLen, lockKind )
   local obj = {}
   LockedAsyncInfo.setmeta( obj )
   if obj.__init then
      obj:__init( loopLen, lockKind )
   end
   return obj
end
function LockedAsyncInfo:__init( loopLen, lockKind )

   self.loopLen = loopLen
   self.lockKind = lockKind
end
function LockedAsyncInfo:get_loopLen()
   return self.loopLen
end
function LockedAsyncInfo:get_lockKind()
   return self.lockKind
end


local NSInfo = {}
_moduleObj.NSInfo = NSInfo
function NSInfo:isLockedAsync(  )

   return #self.lockedAsyncStack > 0
end
function NSInfo:isNoasync(  )

   if self.typeInfo:get_asyncMode() == Ast.Async.Noasync then
      return true
   end
   
   for __index, info in ipairs( self.lockedAsyncStack ) do
      do
         local _switchExp = info:get_lockKind()
         if _switchExp == Nodes.LockKind.AsyncLock or _switchExp == Nodes.LockKind.LuaLock then
            return true
         end
      end
      
   end
   
   return false
end
function NSInfo.new( typeInfo, pos, validAsyncCtrl )
   local obj = {}
   NSInfo.setmeta( obj )
   if obj.__init then obj:__init( typeInfo, pos, validAsyncCtrl ); end
   return obj
end
function NSInfo:__init(typeInfo, pos, validAsyncCtrl) 
   self.idSetInfo = IdSetInfo.new()
   self.nobody = false
   self.lockedAsyncStack = {}
   self.loopScopeQueue = {}
   
   self.typeInfo = typeInfo
   self.pos = pos
   self.validAsyncCtrl = validAsyncCtrl
end
function NSInfo:incLock( lockKind )

   table.insert( self.lockedAsyncStack, LockedAsyncInfo.new(#self.loopScopeQueue, lockKind) )
end
function NSInfo:decLock(  )

   table.remove( self.lockedAsyncStack )
end
function NSInfo:canBreak(  )

   local len = #self.lockedAsyncStack
   local loopQueueLen = #self.loopScopeQueue
   if len == 0 then
      return loopQueueLen > 0
   end
   
   return self.lockedAsyncStack[len]:get_loopLen() < loopQueueLen
end
function NSInfo:canAccessNoasync(  )

   local len = #self.lockedAsyncStack
   if self.typeInfo:get_asyncMode() == Ast.Async.Noasync or (len > 0 and self.lockedAsyncStack[len]:get_lockKind() ~= Nodes.LockKind.Unsafe ) then
      return true
   end
   
   return false
end
function NSInfo:canAccessLuaval(  )

   if not self.validAsyncCtrl then
      return true
   end
   
   if #self.lockedAsyncStack > 0 then
      return true
   end
   
   return false
end
function NSInfo.setmeta( obj )
  setmetatable( obj, { __index = NSInfo  } )
end
function NSInfo:get_nobody()
   return self.nobody
end
function NSInfo:set_nobody( nobody )
   self.nobody = nobody
end
function NSInfo:get_typeInfo()
   return self.typeInfo
end
function NSInfo:get_pos()
   return self.pos
end
function NSInfo:get_loopScopeQueue()
   return self.loopScopeQueue
end
function NSInfo:registerSym( ... )
   return self.idSetInfo:registerSym( ... )
end



local TransUnitIF = {}
_moduleObj.TransUnitIF = TransUnitIF
function TransUnitIF.setmeta( obj )
  setmetatable( obj, { __index = TransUnitIF  } )
end
function TransUnitIF.new(  )
   local obj = {}
   TransUnitIF.setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function TransUnitIF:__init(  )

end


local ErrMess = {}
_moduleObj.ErrMess = ErrMess
function ErrMess.setmeta( obj )
  setmetatable( obj, { __index = ErrMess  } )
end
function ErrMess.new( mess, pos )
   local obj = {}
   ErrMess.setmeta( obj )
   if obj.__init then
      obj:__init( mess, pos )
   end
   return obj
end
function ErrMess:__init( mess, pos )

   self.mess = mess
   self.pos = pos
end
function ErrMess:get_mess()
   return self.mess
end
function ErrMess:get_pos()
   return self.pos
end


local function sortMess( list )

   table.sort( list, function ( mess1, mess2 )
   
      local pos1 = mess1:get_pos():get_orgPos()
      local pos2 = mess2:get_pos():get_orgPos()
      if pos1.streamName < pos2.streamName then
         return true
      end
      
      if pos1.streamName > pos2.streamName then
         return false
      end
      
      if pos1.lineNo < pos2.lineNo then
         return true
      end
      
      if pos1.lineNo > pos2.lineNo then
         return false
      end
      
      if pos1.column < pos2.column then
         return true
      end
      
      if pos1.column > pos2.column then
         return false
      end
      
      return mess1:get_mess() < mess2:get_mess()
   end )
   return list
end
_moduleObj.sortMess = sortMess

local TransUnitBase = {}
setmetatable( TransUnitBase, { ifList = {TransUnitIF,} } )
_moduleObj.TransUnitBase = TransUnitBase
function TransUnitBase:get_scope(  )

   return self.scope
end
function TransUnitBase.new( ctrl_info, processInfo )
   local obj = {}
   TransUnitBase.setmeta( obj )
   if obj.__init then obj:__init( ctrl_info, processInfo ); end
   return obj
end
function TransUnitBase:__init(ctrl_info, processInfo) 
   self.ctrl_info = ctrl_info
   self.typeId2ClassMap = {}
   self.typeNameCtrl = Ast.defaultTypeNameCtrl
   self.errMessList = {}
   self.namespace2Scope = {}
   self.processInfo = processInfo
   self.globalScope = Ast.Scope.new(processInfo, processInfo:get_topScope(), true, nil)
   self.scope = Ast.Scope.new(processInfo, self.globalScope, true, nil)
   self.nsInfoMap = {}
   local subRootTypeInfo = self.processInfo:get_dummyParentType()
   self.nsInfoMap[subRootTypeInfo] = NSInfo.new(subRootTypeInfo, Types.Position.new(0, 0, "@builtin@"), ctrl_info.validAsyncCtrl)
end
function TransUnitBase:addErrMess( pos, mess )

   if mess:find( "type mismatch.*<- &" ) then
      mess = mess .. ". if your code is the old style, use the opiton '--legacy-mutable-control'."
   end
   
   table.insert( self.errMessList, ErrMess.new(string.format( "%s: error: %s", pos:getDisplayTxt(  ), mess), pos) )
end
function TransUnitBase:error( mess )

   self:errorAt( self:getLatestPos(  ), mess )
end
function TransUnitBase:pushScope( classFlag, baseInfo, interfaceList )

   self.scope = Ast.TypeInfo.createScope( self.processInfo, self.scope, classFlag, baseInfo, interfaceList )
   return self.scope
end
function TransUnitBase:popScope(  )

   self.scope = self.scope:get_parent(  )
end
function TransUnitBase:newNSInfo( typeInfo, pos )

   local nsInfo = NSInfo.new(typeInfo, pos, self.ctrl_info.validAsyncCtrl)
   self.nsInfoMap[typeInfo] = nsInfo
   return nsInfo
end
function TransUnitBase:getCurrentNamespaceTypeInfoMut(  )

   local typeInfo = self.scope:getNamespaceTypeInfo(  )
   local nsInfo = self.nsInfoMap[typeInfo]
   if  nil == nsInfo then
      local _nsInfo = nsInfo
   
      self:error( string.format( "not found nsInfo -- %s (%d)", typeInfo:getTxt(  ), typeInfo:get_typeId().id) )
   end
   
   return nsInfo:get_typeInfo()
end
function TransUnitBase:getCurrentNamespaceTypeInfo(  )

   return self.scope:getNamespaceTypeInfo(  )
end
function TransUnitBase:pushModule( processInfo, externalFlag, name, mutable )

   local typeInfo = Ast.headTypeInfo
   
   local modName
   
   if name:find( "^@" ) then
      modName = name
   else
    
      modName = string.format( "@%s", name)
   end
   
   
   local nsInfo
   
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
         
         nsInfo = _lune.unwrap( self.nsInfoMap[typeInfo])
      else
         local _
         local parentInfo = self:getCurrentNamespaceTypeInfoMut(  )
         local parentScope = self.scope
         local scope = self:pushScope( true )
         local newType = processInfo:createModule( scope, parentInfo, externalFlag, modName, mutable )
         typeInfo = newType
         self.namespace2Scope[typeInfo] = scope
         nsInfo = self:newNSInfo( newType, self:getLatestPos(  ) )
         
         local _1, existSym = parentScope:addClass( processInfo, modName, nil, typeInfo )
         if existSym ~= nil then
            self:addErrMess( self:getLatestPos(  ), string.format( "module symbols exist -- %s.%s -- %s.%s", existSym:get_namespaceTypeInfo():getFullName( self.typeNameCtrl, parentScope, false ), existSym:get_name(), parentInfo:getFullName( self.typeNameCtrl, parentScope, false ), modName) )
         end
         
      end
   end
   
   if not self.typeId2ClassMap[typeInfo:get_typeId()] then
      local namespace = Nodes.NamespaceInfo.new(modName, self.scope, typeInfo)
      self.typeId2ClassMap[typeInfo:get_typeId()] = namespace
   end
   
   return nsInfo
end
function TransUnitBase:pushModuleLow( processInfo, externalFlag, name, mutable )

   return self:pushModule( processInfo, externalFlag, name, mutable ):get_typeInfo()
end
function TransUnitBase:popModule(  )

   self:popScope(  )
end
function TransUnitBase:pushClassScope( errPos, classTypeInfo, scope )

   if self.scope ~= _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ) then
      local classParentName
      
      local classParentTypeId
      
      do
         local parentType = _lune.nilacc( _lune.nilacc( classTypeInfo:get_scope(), 'get_parent', 'callmtd' ), 'get_ownerTypeInfo', 'callmtd' )
         if parentType ~= nil then
            classParentName = parentType:getFullName( self.typeNameCtrl, _lune.unwrap( parentType:get_scope()), false )
            classParentTypeId = parentType:get_typeId()
         else
            classParentName = "nil"
            classParentTypeId = Ast.dummyIdInfo
         end
      end
      
      self:addErrMess( errPos, string.format( "This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", classTypeInfo:getTxt(  ), _lune.nilacc( self.scope:get_ownerTypeInfo(), 'getFullName', 'callmtd' , self.typeNameCtrl, self.scope, false ), _lune.nilacc( _lune.nilacc( self.scope:get_ownerTypeInfo(), 'get_typeId', 'callmtd' ), "id" ) or -1, classParentName, classParentTypeId.id) )
   end
   
   self.scope = scope
end
function TransUnitBase:pushClass( processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   local typeInfo
   
   local nsInfo
   
   do
      local _exp = self.scope:getTypeInfo( name, self.scope, true, Ast.ScopeAccess.Normal )
      if _exp ~= nil then
         
         typeInfo = _exp
         nsInfo = _lune.unwrap( self.nsInfoMap[typeInfo])
         
         if _lune.nilacc( typeInfo:get_scope(), 'get_parent', 'callmtd' ) ~= self.scope then
            self:addErrMess( errPos, string.format( "multiple class(%s)", typeInfo:getTxt( self.typeNameCtrl )) )
            
            self:error( "stop by error" )
         end
         
         
         if typeInfo:get_abstractFlag() ~= abstractFlag then
            self:addErrMess( errPos, string.format( "mismatch class(%s) abstract for prototpye", typeInfo:getTxt( self.typeNameCtrl )) )
         end
         
         if typeInfo:get_accessMode() ~= accessMode then
            self:addErrMess( errPos, string.format( "mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", typeInfo:getTxt( self.typeNameCtrl ), Ast.AccessMode:_getTxt( accessMode)
            , Ast.AccessMode:_getTxt( typeInfo:get_accessMode())
            ) )
         end
         
         if baseInfo ~= nil then
            if typeInfo:get_baseTypeInfo() ~= baseInfo then
               self:addErrMess( errPos, string.format( "mismatch class(%s) base class(%s) for prototpye base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), baseInfo:getTxt(  ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         else
            if typeInfo:hasBase(  ) then
               self:addErrMess( errPos, string.format( "mismatch class(%s) base class(None) for prototpye base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         end
         
         
         local function compareList( protoList, typeList, message )
         
            if #protoList == #typeList then
               for index, protoType in ipairs( protoList ) do
                  if protoType ~= typeList[index] then
                     self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%s) for prototpye %s(%s)", typeInfo:getTxt( self.typeNameCtrl ), message, typeList[index]:getTxt( self.typeNameCtrl ), message, protoType:getTxt(  )) )
                  end
                  
               end
               
            else
             
               self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%d) for prototpye %s(%d)", typeInfo:getTxt( self.typeNameCtrl ), message, #typeList, message, #protoList) )
            end
            
         end
         
         compareList( typeInfo:get_interfaceList(), _lune.unwrapDefault( interfaceList, {}), "interface" )
         
         compareList( typeInfo:get_itemTypeInfoList(), _lune.unwrapDefault( genTypeList, {}), "generics" )
         
         do
            local scope = self.namespace2Scope[typeInfo]
            if scope ~= nil then
               self.scope = scope
            else
               self:errorAt( errPos, string.format( "not find scope -- %s", name) )
            end
         end
         
         do
            local _switchExp = (typeInfo:get_kind() )
            if _switchExp == Ast.TypeInfoKind.Class then
               if mode == DeclClassMode.Interface then
                  self:addErrMess( errPos, string.format( "define interface already -- %s", name) )
                  Util.printStackTrace(  )
               end
               
            elseif _switchExp == Ast.TypeInfoKind.IF then
               if mode ~= DeclClassMode.Interface then
                  self:addErrMess( errPos, string.format( "define class already -- %s", name) )
                  Util.printStackTrace(  )
               end
               
            end
         end
         
      else
         local parentInfo = self:getCurrentNamespaceTypeInfoMut(  )
         
         local parentScope = self.scope
         local scope = self:pushScope( true, baseInfo, interfaceList )
         local workGenTypeList
         
         if genTypeList ~= nil then
            workGenTypeList = genTypeList
         else
            workGenTypeList = {}
         end
         
         
         local newType = processInfo:createClassAsync( mode ~= DeclClassMode.Interface, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, externalFlag, accessMode, name )
         typeInfo = newType
         self.namespace2Scope[typeInfo] = scope
         nsInfo = self:newNSInfo( newType, errPos )
         
         parentScope:addClassLazy( processInfo, name, errPos, typeInfo, mode == DeclClassMode.LazyModule )
      end
   end
   
   if genTypeList ~= nil then
      for __index, genType in ipairs( genTypeList ) do
         self.scope:addAlternate( processInfo, accessMode, genType:get_txt(), errPos, genType )
      end
      
   end
   
   
   local namespace = defNamespace
   if  nil == namespace then
      local _namespace = namespace
   
      namespace = Nodes.NamespaceInfo.new(name, self.scope, typeInfo)
   end
   
   self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
   
   return nsInfo
end
function TransUnitBase:pushClassLow( processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   return self:pushClass( processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace ):get_typeInfo()
end
function TransUnitBase:popClass(  )

   self:popScope(  )
end
function TransUnitBase.setmeta( obj )
  setmetatable( obj, { __index = TransUnitBase  } )
end
function TransUnitBase:get_globalScope()
   return self.globalScope
end
function TransUnitBase:get_errMessList()
   return self.errMessList
end


local SimpeTransUnit = {}
setmetatable( SimpeTransUnit, { __index = TransUnitBase } )
_moduleObj.SimpeTransUnit = SimpeTransUnit
function SimpeTransUnit:errorAt( pos, mess )

   self:addErrMess( pos, mess )
   for __index, errmess in ipairs( self.errMessList ) do
      Util.errorLog( errmess:get_mess() )
   end
   
   do
      local nearCode = self.nearCode
      if nearCode ~= nil then
         print( "------ near code -----", self.macroMode )
         print( nearCode )
         print( "------" )
      end
   end
   
   
   Util.err( "has error" )
end
function SimpeTransUnit:getLatestPos(  )

   return self.latestPos
end
function SimpeTransUnit.setmeta( obj )
  setmetatable( obj, { __index = SimpeTransUnit  } )
end
function SimpeTransUnit.new( __superarg1, __superarg2,latestPos, macroMode, nearCode )
   local obj = {}
   SimpeTransUnit.setmeta( obj )
   if obj.__init then
      obj:__init( __superarg1, __superarg2,latestPos, macroMode, nearCode )
   end
   return obj
end
function SimpeTransUnit:__init( __superarg1, __superarg2,latestPos, macroMode, nearCode )

   TransUnitBase.__init( self, __superarg1, __superarg2 )
   self.latestPos = latestPos
   self.macroMode = macroMode
   self.nearCode = nearCode
end


return _moduleObj
