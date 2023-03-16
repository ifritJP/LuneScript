--lune/base/TransUnitIF.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@TransUnitIF'
local _lune = {}
if _lune8 then
   _lune = _lune8
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
   if __luneScript and not package.preload[ mod ] then
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

if not _lune8 then
   _lune8 = _lune
end



local Tokenizer = _lune.loadModule( 'lune.base.Tokenizer' )
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
function Modifier._setmeta( obj )
  setmetatable( obj, { __index = Modifier  } )
end
function Modifier._new( validMutControl, processInfo )
   local obj = {}
   Modifier._setmeta( obj )
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
function IdSetInfo._new(  )
   local obj = {}
   IdSetInfo._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function IdSetInfo:__init() 
   self.anonymousFuncId = Ast.IdProvider._new(0, 10000)
   self.anonymousVarId = Ast.IdProvider._new(0, 10000)
end
function IdSetInfo:registerSym( symbol )

   if symbol:get_kind() == Ast.SymbolKind.Var then
      if symbol:get_name() == "_" then
         local id = self.anonymousVarId:getNewId(  )
         return Ast.AnonymousSymbolInfo._new(symbol, id)
      end
      
   end
   
   return symbol
end
function IdSetInfo._setmeta( obj )
  setmetatable( obj, { __index = IdSetInfo  } )
end


local LockedAsyncInfo = {}
function LockedAsyncInfo._setmeta( obj )
  setmetatable( obj, { __index = LockedAsyncInfo  } )
end
function LockedAsyncInfo._new( loopLen, lockKind )
   local obj = {}
   LockedAsyncInfo._setmeta( obj )
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


local StmtKind = {}
_moduleObj.StmtKind = StmtKind
StmtKind._val2NameMap = {}
function StmtKind:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "StmtKind.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function StmtKind._from( val )
   if StmtKind._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
StmtKind.__allList = {}
function StmtKind.get__allList()
   return StmtKind.__allList
end

StmtKind.Switch = 1
StmtKind._val2NameMap[1] = 'Switch'
StmtKind.__allList[1] = StmtKind.Switch
StmtKind.Match = 2
StmtKind._val2NameMap[2] = 'Match'
StmtKind.__allList[2] = StmtKind.Match
StmtKind.For = 3
StmtKind._val2NameMap[3] = 'For'
StmtKind.__allList[3] = StmtKind.For
StmtKind.Foreach = 4
StmtKind._val2NameMap[4] = 'Foreach'
StmtKind.__allList[4] = StmtKind.Foreach
StmtKind.Forsort = 5
StmtKind._val2NameMap[5] = 'Forsort'
StmtKind.__allList[5] = StmtKind.Forsort
StmtKind.Apply = 6
StmtKind._val2NameMap[6] = 'Apply'
StmtKind.__allList[6] = StmtKind.Apply


local NSInfo = {}
_moduleObj.NSInfo = NSInfo
function NSInfo:addStmtNum( num )

   self.stmtNum = self.stmtNum + num
end
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
function NSInfo._new( typeInfo, typeDataAccessor, pos, validAsyncCtrl )
   local obj = {}
   NSInfo._setmeta( obj )
   if obj.__init then obj:__init( typeInfo, typeDataAccessor, pos, validAsyncCtrl ); end
   return obj
end
function NSInfo:__init(typeInfo, typeDataAccessor, pos, validAsyncCtrl) 
   self.stmtIdList = {}
   for __index, _1 in ipairs( StmtKind.get__allList() ) do
      table.insert( self.stmtIdList, 0 )
   end
   
   
   self.idSetInfo = IdSetInfo._new()
   self.nobody = false
   self.lockedAsyncStack = {}
   self.loopScopeQueue = {}
   
   self.typeDataAccessor = typeDataAccessor
   self.typeInfo = typeInfo
   self.pos = pos
   self.validAsyncCtrl = validAsyncCtrl
   self.stmtNum = 0
   self.condRetNodeList = {}
   self.condRetCount = 0
end
function NSInfo:duplicate(  )

   local typeData = Ast.TypeData._new()
   local nsInfo = NSInfo._new(self.typeInfo, Ast.SimpleTypeDataAccessor._new(typeData), self.pos, self.validAsyncCtrl)
   
   typeData:addFrom( self.typeDataAccessor:get_typeData() )
   
   return nsInfo
end
function NSInfo:getNextStmtId( stmtKind )

   local id = self.stmtIdList[stmtKind]
   self.stmtIdList[stmtKind] = id + 1
   return id
end
function NSInfo:incLock( lockKind )

   table.insert( self.lockedAsyncStack, LockedAsyncInfo._new(#self.loopScopeQueue, lockKind) )
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
   if self.typeInfo:get_asyncMode() == Ast.Async.Noasync or (len > 0 and self.lockedAsyncStack[len]:get_lockKind() ~= Nodes.LockKind.LuaDepend ) then
      return true
   end
   
   return false
end
function NSInfo:canAccessLuaval(  )

   if not self.validAsyncCtrl then
      return true
   end
   
   for __index, stack in ipairs( self.lockedAsyncStack ) do
      if stack:get_lockKind() ~= Nodes.LockKind.AsyncLock then
         return true
      end
      
   end
   
   return false
end
function NSInfo:addCondRet( nodeManager, pos, inTestBlock, isInAnalyzeArgMode, expOkType, exp, kind )

   self.condRetCount = self.condRetCount + 1
   
   if kind ~= Nodes.CondRetKind.Two then
      if #exp:get_expTypeList() > 1 then
         exp = Nodes.ExpMultiTo1Node.create( nodeManager, pos, inTestBlock, isInAnalyzeArgMode, {exp:get_expType()}, exp )
      end
      
   end
   
   
   local condRetNode = Nodes.CondRetNode.create( nodeManager, pos, inTestBlock, isInAnalyzeArgMode, {expOkType}, exp, kind, self.condRetCount )
   table.insert( self.condRetNodeList, condRetNode )
   return condRetNode
end
function NSInfo:clearCondRetNodeList(  )

   if #self.condRetNodeList > 0 then
      self.condRetNodeList = {}
   end
   
end
function NSInfo._setmeta( obj )
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
function NSInfo:get_typeDataAccessor()
   return self.typeDataAccessor
end
function NSInfo:get_pos()
   return self.pos
end
function NSInfo:get_loopScopeQueue()
   return self.loopScopeQueue
end
function NSInfo:get_condRetNodeList()
   return self.condRetNodeList
end
function NSInfo:get_stmtNum()
   return self.stmtNum
end
function NSInfo:registerSym( ... )
   return self.idSetInfo:registerSym( ... )
end



local TransUnitIF = {}
_moduleObj.TransUnitIF = TransUnitIF
function TransUnitIF._setmeta( obj )
  setmetatable( obj, { __index = TransUnitIF  } )
end
function TransUnitIF._new(  )
   local obj = {}
   TransUnitIF._setmeta( obj )
   if obj.__init then
      obj:__init(  )
   end
   return obj
end
function TransUnitIF:__init(  )

end


local ErrMess = {}
_moduleObj.ErrMess = ErrMess
function ErrMess._setmeta( obj )
  setmetatable( obj, { __index = ErrMess  } )
end
function ErrMess._new( mess, pos )
   local obj = {}
   ErrMess._setmeta( obj )
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

local SetNSInfo = {}
_moduleObj.SetNSInfo = SetNSInfo
SetNSInfo._val2NameMap = {}
function SetNSInfo:_getTxt( val )
   local name = self._val2NameMap[ val ]
   if name then
      return string.format( "SetNSInfo.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end
function SetNSInfo._from( val )
   if SetNSInfo._val2NameMap[ val ] then
      return val
   end
   return nil
end
    
SetNSInfo.__allList = {}
function SetNSInfo.get__allList()
   return SetNSInfo.__allList
end

SetNSInfo.None = 0
SetNSInfo._val2NameMap[0] = 'None'
SetNSInfo.__allList[1] = SetNSInfo.None
SetNSInfo.FromScope = 1
SetNSInfo._val2NameMap[1] = 'FromScope'
SetNSInfo.__allList[2] = SetNSInfo.FromScope


local TransUnitBase = {}
setmetatable( TransUnitBase, { ifList = {TransUnitIF,} } )
_moduleObj.TransUnitBase = TransUnitBase
function TransUnitBase:getCurrentNamespaceTypeInfo(  )

   return self.scope:getNamespaceTypeInfo(  )
end
function TransUnitBase:error( mess )

   self:errorAt( self:getLatestPos(  ), mess )
end
function TransUnitBase:get_curNsInfo(  )

   do
      local _exp = self.curNsInfo
      if _exp ~= nil then
         return _exp
      end
   end
   
   self:error( string.format( "not found NSInfo for %s", self:getCurrentNamespaceTypeInfo(  )) )
end
function TransUnitBase:set_curNsInfo( nsInfo )

   self.curNsInfo = nsInfo
end
function TransUnitBase:get_scope(  )

   return self.scope
end
function TransUnitBase:get_scopeRO(  )

   return self.scope
end
function TransUnitBase:getNSInfo( typeInfo )

   local nsInfo = self.nsInfoMap[typeInfo]
   if  nil == nsInfo then
      local _nsInfo = nsInfo
   
      self:error( string.format( "not found TypeInfo -- %s", typeInfo:getTxt(  )) )
   end
   
   return nsInfo
end
function TransUnitBase:setScope( scope, setNSInfo )

   self.scope = scope
   do
      local _switchExp = setNSInfo
      if _switchExp == SetNSInfo.None then
      elseif _switchExp == SetNSInfo.FromScope then
         self.curNsInfo = self:getNSInfo( self:getCurrentNamespaceTypeInfo(  ) )
      end
   end
   
end
function TransUnitBase._new( ctrl_info, processInfo )
   local obj = {}
   TransUnitBase._setmeta( obj )
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
   self.globalScope = Ast.Scope._new(processInfo, processInfo:get_topScope(), Ast.ScopeKind.Module, nil)
   self.scope = Ast.Scope._new(processInfo, self.globalScope, Ast.ScopeKind.Module, nil)
   self.nsInfoMap = {}
   local subRootTypeInfo = self.processInfo:get_dummyParentType()
   local nsInfo = NSInfo._new(subRootTypeInfo, subRootTypeInfo, Types.Position._new(0, 0, "@builtin@"), ctrl_info.validAsyncCtrl)
   self.curNsInfo = nsInfo
   self.nsInfoMap[subRootTypeInfo] = nsInfo
end
function TransUnitBase:addErrMess( pos, mess )

   if mess:find( "type mismatch.*<- &" ) then
      mess = mess .. ". if your code is the old style, use the opiton '--legacy-mutable-control'."
   end
   
   table.insert( self.errMessList, ErrMess._new(string.format( "%s: error: %s", pos:getDisplayTxt(  ), mess), pos) )
end
function TransUnitBase:pushScope( scopeKind, baseInfo, interfaceList )

   self.scope = Ast.TypeInfo.createScope( self.processInfo, self.scope, scopeKind, baseInfo, interfaceList )
   return self.scope
end
function TransUnitBase:popScope(  )

   self.scope = self.scope:get_outerScope()
   local nsInfo = self.nsInfoMap[self:getCurrentNamespaceTypeInfo(  )]
   if nsInfo ~= self.curNsInfo then
      do
         local curNsInfo = self.curNsInfo
         if curNsInfo ~= nil then
            if #curNsInfo:get_condRetNodeList() ~= 0 then
               self:error( "internal error. condRetNodeList is not nil." )
            end
            
         end
      end
      
      self.curNsInfo = nsInfo
   end
   
end
function TransUnitBase:newNSInfoWithTypeData( typeInfo, typeDataAccessor, pos )

   local nsInfo = NSInfo._new(typeInfo, typeDataAccessor, pos, self.ctrl_info.validAsyncCtrl)
   self.nsInfoMap[typeInfo] = nsInfo
   return nsInfo
end
function TransUnitBase:newNSInfo( typeInfo, pos )

   local nsInfo = NSInfo._new(typeInfo, typeInfo, pos, self.ctrl_info.validAsyncCtrl)
   self.nsInfoMap[typeInfo] = nsInfo
   return nsInfo
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
         
         local parentNsInfo = self:get_curNsInfo()
         local parentInfo = parentNsInfo:get_typeInfo()
         local parentScope = self.scope
         local scope = self:pushScope( Ast.ScopeKind.Module )
         local newType = processInfo:createModule( scope, parentInfo, parentNsInfo:get_typeDataAccessor(), externalFlag, modName, mutable )
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
      local namespace = Nodes.NamespaceInfo._new(modName, self.scope, typeInfo)
      self.typeId2ClassMap[typeInfo:get_typeId()] = namespace
   end
   
   
   self.curNsInfo = nsInfo
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
   self.curNsInfo = self.nsInfoMap[classTypeInfo]
end
function TransUnitBase:pushClass( processInfo, errPos, mode, finalFlag, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   local nsInfo
   
   
   local typeInfo
   
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
            self:addErrMess( errPos, string.format( "mismatch class(%s) abstract for prototype", typeInfo:getTxt( self.typeNameCtrl )) )
         end
         
         if typeInfo:get_accessMode() ~= accessMode then
            self:addErrMess( errPos, string.format( "mismatch class(%s) accessmode(%s) for prototype accessmode(%s)", typeInfo:getTxt( self.typeNameCtrl ), Ast.AccessMode:_getTxt( accessMode)
            , Ast.AccessMode:_getTxt( typeInfo:get_accessMode())
            ) )
         end
         
         if baseInfo ~= nil then
            if typeInfo:get_baseTypeInfo() ~= baseInfo then
               self:addErrMess( errPos, string.format( "mismatch class(%s) base class(%s) for prototype base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), baseInfo:getTxt(  ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         else
            if typeInfo:hasBase(  ) then
               self:addErrMess( errPos, string.format( "mismatch class(%s) base class(None) for prototype base class(%s)", typeInfo:getTxt( self.typeNameCtrl ), typeInfo:get_baseTypeInfo():getTxt(  )) )
            end
            
         end
         
         
         
         
         do
            local typeList = _lune.unwrapDefault( interfaceList, {})
            if #typeInfo:get_interfaceList() == #typeList then
               for index, protoType in ipairs( typeInfo:get_interfaceList() ) do
                  if protoType ~= typeList[index] then
                     self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%s) for prototype %s(%s)", typeInfo:getTxt( self.typeNameCtrl ), "interface", typeList[index]:getTxt( self.typeNameCtrl ), "interface", protoType:getTxt(  )) )
                  end
                  
               end
               
            else
             
               self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%d) for prototype %s(%d)", typeInfo:getTxt( self.typeNameCtrl ), "interface", #typeList, "interface", #typeInfo:get_interfaceList()) )
            end
            
         end
         
         
         
         do
            local typeList = _lune.unwrapDefault( genTypeList, {})
            if #typeInfo:get_itemTypeInfoList() == #typeList then
               for index, protoType in ipairs( typeInfo:get_itemTypeInfoList() ) do
                  if protoType ~= typeList[index] then
                     self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%s) for prototype %s(%s)", typeInfo:getTxt( self.typeNameCtrl ), "generics", typeList[index]:getTxt( self.typeNameCtrl ), "generics", protoType:getTxt(  )) )
                  end
                  
               end
               
            else
             
               self:addErrMess( errPos, string.format( "mismatch class(%s) %s(%d) for prototype %s(%d)", typeInfo:getTxt( self.typeNameCtrl ), "generics", #typeList, "generics", #typeInfo:get_itemTypeInfoList()) )
            end
            
         end
         
         
         
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
         
         local parentNsInfo = self:get_curNsInfo()
         
         local parentScope = self.scope
         local scope = self:pushScope( Ast.ScopeKind.Class, baseInfo, interfaceList )
         local workGenTypeList
         
         if genTypeList ~= nil then
            workGenTypeList = genTypeList
         else
            workGenTypeList = {}
         end
         
         
         local newType = processInfo:createClassAsync( mode ~= DeclClassMode.Interface, finalFlag, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentNsInfo:get_typeInfo(), parentNsInfo:get_typeDataAccessor(), externalFlag, accessMode, name )
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
   
      namespace = Nodes.NamespaceInfo._new(name, self.scope, typeInfo)
   end
   
   self.typeId2ClassMap[typeInfo:get_typeId(  )] = namespace
   
   self.curNsInfo = nsInfo
   
   return nsInfo
end
function TransUnitBase:pushClassLow( processInfo, errPos, mode, finalFlag, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace )

   return self:pushClass( processInfo, errPos, mode, finalFlag, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace ):get_typeInfo()
end
function TransUnitBase:popClass(  )

   self:popScope(  )
end
function TransUnitBase._setmeta( obj )
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
         Util.println( "------ near code -----", self.macroMode )
         Util.println( nearCode )
         Util.println( "------" )
      end
   end
   
   
   Util.err( "has error" )
end
function SimpeTransUnit:getLatestPos(  )

   return self.latestPos
end
function SimpeTransUnit._setmeta( obj )
  setmetatable( obj, { __index = SimpeTransUnit  } )
end
function SimpeTransUnit._new( __superarg1, __superarg2,latestPos, macroMode, nearCode )
   local obj = {}
   SimpeTransUnit._setmeta( obj )
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
