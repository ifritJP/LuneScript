--lune/base/convWasm.lns
local _moduleObj = {}
local __mod__ = '@lune.@base.@convWasm'
local _lune = {}
if _lune8 then
   _lune = _lune8
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
local isScalar
local isMethodSym
local append
local makeIfChain



local Ast = _lune.loadModule( 'lune.base.Ast' )

local Nodes = _lune.loadModule( 'lune.base.Nodes' )

local Util = _lune.loadModule( 'lune.base.Util' )

local AstInfo = _lune.loadModule( 'lune.base.AstInfo' )

local Builtin = _lune.loadModule( 'lune.base.Builtin' )

local W = _lune.loadModule( 'lune.base.WasmBin' )

local Rt = _lune.loadModule( 'lune.base.WasmRt' )


local Opt = {}
_moduleObj.Opt = Opt
function Opt._new( parent )
   local obj = {}
   Opt._setmeta( obj )
   if obj.__init then obj:__init( parent ); end
   return obj
end
function Opt:__init(parent) 
   self.parent = parent
end
function Opt._setmeta( obj )
  setmetatable( obj, { __index = Opt  } )
end


local Option = {}
_moduleObj.Option = Option
function Option._new(  )
   local obj = {}
   Option._setmeta( obj )
   if obj.__init then obj:__init(  ); end
   return obj
end
function Option:__init() 
end
function Option._setmeta( obj )
  setmetatable( obj, { __index = Option  } )
end

local strCmpOpMap = {["~="] = W.opI32Ne, ["<"] = W.opI32LtS, ["<="] = W.opI32LeS, [">"] = W.opI32GtS, [">="] = W.opI32GeS}

local i64CmpOpMap = {["=="] = W.opI64Eq, ["~="] = W.opI64Ne, ["<"] = W.opI64LtS, ["<="] = W.opI64LeS, [">"] = W.opI64GtS, [">="] = W.opI64GeS}

local f64CmpOpMap = {["=="] = W.opF64Eq, ["~="] = W.opF64Ne, ["<"] = W.opF64Lt, ["<="] = W.opF64Le, [">"] = W.opF64Gt, [">="] = W.opF64Ge}


local function lg( index )

   return W.localGet( index )
end

local function ls( index, val )

   return W.localSet( index, val )
end

local function p( code, args )

   return W.plain( code, args )
end

local VarLoc = {}
VarLoc._name2Val = {}
function VarLoc:_getTxt( val )
   local name = val[ 1 ]
   if name then
      return string.format( "VarLoc.%s", name )
   end
   return string.format( "illegal val -- %s", val )
end

function VarLoc._from( val )
   return _lune._AlgeFrom( VarLoc, val )
end

VarLoc.Captured = { "Captured", {{ func=_lune._toInt, nilable=false, child={} }}}
VarLoc._name2Val["Captured"] = VarLoc.Captured
VarLoc.CellLocal = { "CellLocal", {{ func=_lune._toInt, nilable=false, child={} }}}
VarLoc._name2Val["CellLocal"] = VarLoc.CellLocal
VarLoc.Global = { "Global", {{ func=_lune._toInt, nilable=false, child={} }}}
VarLoc._name2Val["Global"] = VarLoc.Global
VarLoc.Local = { "Local", {{ func=_lune._toInt, nilable=false, child={} }}}
VarLoc._name2Val["Local"] = VarLoc.Local

local Val = {}
function Val._setmeta( obj )
  setmetatable( obj, { __index = Val  } )
end
function Val._new( ins, rep, typeInfo )
   local obj = {}
   Val._setmeta( obj )
   if obj.__init then
      obj:__init( ins, rep, typeInfo )
   end
   return obj
end
function Val:__init( ins, rep, typeInfo )

   self.ins = ins
   self.rep = rep
   self.typeInfo = typeInfo
end
function Val:get_ins()
   return self.ins
end
function Val:get_rep()
   return self.rep
end
function Val:get_typeInfo()
   return self.typeInfo
end

local SlotInfo = {}
function SlotInfo._setmeta( obj )
  setmetatable( obj, { __index = SlotInfo  } )
end
function SlotInfo._new( name, funcType, methodType, impl )
   local obj = {}
   SlotInfo._setmeta( obj )
   if obj.__init then
      obj:__init( name, funcType, methodType, impl )
   end
   return obj
end
function SlotInfo:__init( name, funcType, methodType, impl )

   self.name = name
   self.funcType = funcType
   self.methodType = methodType
   self.impl = impl
end
function SlotInfo:get_name()
   return self.name
end
function SlotInfo:get_funcType()
   return self.funcType
end
function SlotInfo:get_methodType()
   return self.methodType
end
function SlotInfo:get_impl()
   return self.impl
end
function SlotInfo:set_impl( impl )
   self.impl = impl
end

local ClassInfo = {}
function ClassInfo._new( typeInfo, structType, vtType, vtGlobal, parent )
   local obj = {}
   ClassInfo._setmeta( obj )
   if obj.__init then obj:__init( typeInfo, structType, vtType, vtGlobal, parent ); end
   return obj
end
function ClassInfo:__init(typeInfo, structType, vtType, vtGlobal, parent) 
   self.typeInfo = typeInfo
   self.structType = structType
   self.vtType = vtType
   self.vtGlobal = vtGlobal
   self.parent = parent
   self.fieldNames = {}
   self.fieldReps = {}
   self.name2field = {}
   self.slotList = {}
   self.name2slot = {}
   self.ifList = {}
   self.ctorFunc = -1
   self.ctorFuncType = -1
   if parent ~= nil then
      for __index, name in ipairs( parent:get_fieldNames() ) do
         table.insert( self.fieldNames, name )
      end
      for __index, rep in ipairs( parent:get_fieldReps() ) do
         table.insert( self.fieldReps, rep )
      end
      for name, index in pairs( parent:get_name2field() ) do
         self.name2field[name] = index
      end
      for __index, slot in ipairs( parent:get_slotList() ) do
         table.insert( self.slotList, SlotInfo._new(slot:get_name(), slot:get_funcType(), slot:get_methodType(), slot:get_impl()) )
      end
      for name, index in pairs( parent:get_name2slot() ) do
         self.name2slot[name] = index
      end
      for __index, ifType in ipairs( parent:get_ifList() ) do
         table.insert( self.ifList, ifType )
      end
   end
end
function ClassInfo._setmeta( obj )
  setmetatable( obj, { __index = ClassInfo  } )
end
function ClassInfo:get_typeInfo()
   return self.typeInfo
end
function ClassInfo:get_structType()
   return self.structType
end
function ClassInfo:get_vtType()
   return self.vtType
end
function ClassInfo:get_vtGlobal()
   return self.vtGlobal
end
function ClassInfo:get_parent()
   return self.parent
end
function ClassInfo:get_fieldNames()
   return self.fieldNames
end
function ClassInfo:get_fieldReps()
   return self.fieldReps
end
function ClassInfo:get_name2field()
   return self.name2field
end
function ClassInfo:get_slotList()
   return self.slotList
end
function ClassInfo:get_name2slot()
   return self.name2slot
end
function ClassInfo:get_ifList()
   return self.ifList
end
function ClassInfo:get_ctorFunc()
   return self.ctorFunc
end
function ClassInfo:set_ctorFunc( ctorFunc )
   self.ctorFunc = ctorFunc
end
function ClassInfo:get_ctorFuncType()
   return self.ctorFuncType
end
function ClassInfo:set_ctorFuncType( ctorFuncType )
   self.ctorFuncType = ctorFuncType
end

local IfInfo = {}
function IfInfo._setmeta( obj )
  setmetatable( obj, { __index = IfInfo  } )
end
function IfInfo._new( typeInfo, ifId, itType, methodNames, methodFuncTypes, name2index )
   local obj = {}
   IfInfo._setmeta( obj )
   if obj.__init then
      obj:__init( typeInfo, ifId, itType, methodNames, methodFuncTypes, name2index )
   end
   return obj
end
function IfInfo:__init( typeInfo, ifId, itType, methodNames, methodFuncTypes, name2index )

   self.typeInfo = typeInfo
   self.ifId = ifId
   self.itType = itType
   self.methodNames = methodNames
   self.methodFuncTypes = methodFuncTypes
   self.name2index = name2index
end
function IfInfo:get_typeInfo()
   return self.typeInfo
end
function IfInfo:get_ifId()
   return self.ifId
end
function IfInfo:get_itType()
   return self.itType
end
function IfInfo:get_methodNames()
   return self.methodNames
end
function IfInfo:get_methodFuncTypes()
   return self.methodFuncTypes
end
function IfInfo:get_name2index()
   return self.name2index
end

local AlgeInfo = {}
function AlgeInfo._setmeta( obj )
  setmetatable( obj, { __index = AlgeInfo  } )
end
function AlgeInfo._new( typeInfo, baseType, valNames, name2tag, name2type, name2global, namesGlobal )
   local obj = {}
   AlgeInfo._setmeta( obj )
   if obj.__init then
      obj:__init( typeInfo, baseType, valNames, name2tag, name2type, name2global, namesGlobal )
   end
   return obj
end
function AlgeInfo:__init( typeInfo, baseType, valNames, name2tag, name2type, name2global, namesGlobal )

   self.typeInfo = typeInfo
   self.baseType = baseType
   self.valNames = valNames
   self.name2tag = name2tag
   self.name2type = name2type
   self.name2global = name2global
   self.namesGlobal = namesGlobal
end
function AlgeInfo:get_typeInfo()
   return self.typeInfo
end
function AlgeInfo:get_baseType()
   return self.baseType
end
function AlgeInfo:get_valNames()
   return self.valNames
end
function AlgeInfo:get_name2tag()
   return self.name2tag
end
function AlgeInfo:get_name2type()
   return self.name2type
end
function AlgeInfo:get_name2global()
   return self.name2global
end
function AlgeInfo:get_namesGlobal()
   return self.namesGlobal
end

local FuncCtx = {}
function FuncCtx._new( func, retReps, retTypes )
   local obj = {}
   FuncCtx._setmeta( obj )
   if obj.__init then obj:__init( func, retReps, retTypes ); end
   return obj
end
function FuncCtx:__init(func, retReps, retTypes) 
   self.func = func
   self.sym2loc = {}
   self.retReps = retReps
   self.retTypes = retTypes
   self.depth = 0
   self.breakDepthList = {}
   self.envLocal = -1
   self.envType = -1
   self.selfLocal = -1
   self.selfClass = nil
end
function FuncCtx:addLocal( valType )

   return self.func:addLocal( valType )
end
function FuncCtx._setmeta( obj )
  setmetatable( obj, { __index = FuncCtx  } )
end
function FuncCtx:get_func()
   return self.func
end
function FuncCtx:get_sym2loc()
   return self.sym2loc
end
function FuncCtx:get_retReps()
   return self.retReps
end
function FuncCtx:get_retTypes()
   return self.retTypes
end
function FuncCtx:get_depth()
   return self.depth
end
function FuncCtx:set_depth( depth )
   self.depth = depth
end
function FuncCtx:get_breakDepthList()
   return self.breakDepthList
end
function FuncCtx:get_envLocal()
   return self.envLocal
end
function FuncCtx:set_envLocal( envLocal )
   self.envLocal = envLocal
end
function FuncCtx:get_envType()
   return self.envType
end
function FuncCtx:set_envType( envType )
   self.envType = envType
end
function FuncCtx:get_selfLocal()
   return self.selfLocal
end
function FuncCtx:set_selfLocal( selfLocal )
   self.selfLocal = selfLocal
end
function FuncCtx:get_selfClass()
   return self.selfClass
end
function FuncCtx:set_selfClass( selfClass )
   self.selfClass = selfClass
end


local convFilter = {}
setmetatable( convFilter, { __index = Nodes.Filter } )
function convFilter._new( streamName, stream, ast )
   local obj = {}
   convFilter._setmeta( obj )
   if obj.__init then obj:__init( streamName, stream, ast ); end
   return obj
end
function convFilter:__init(streamName, stream, ast) 
   Nodes.Filter.__init( self,true, ast:get_exportInfo():get_moduleTypeInfo(), ast:get_exportInfo():get_moduleTypeInfo():get_scope())
   self.streamName = streamName
   self.stream = stream
   self.processInfo = ast:get_exportInfo():get_processInfo():clone(  )
   self.builtinFuncs = ast:get_builtinFunc()
   self.moduleTypeInfo = ast:get_exportInfo():get_moduleTypeInfo()
   self.moduleScope = _lune.unwrap( ast:get_exportInfo():get_moduleTypeInfo():get_scope())
   local mod = W.Module._new()
   self.rt = Rt.Runtime._new(mod)
   self.mod = mod
   local mainIndex, mainFunc = self.mod:addFunc( {}, {} )
   self.mainIndex = mainIndex
   self.mod:setFuncName( mainIndex, "main" )
   self.ctx = FuncCtx._new(mainFunc, {}, {})
   self.compiledFuncSet = {}
   self.globalSym2loc = {}
   self.func2index = {}
   self.func2decl = {}
   self.classType2decl = {}
   self.class2info = {}
   self.classInfoList = {}
   self.if2info = {}
   self.ifInfoList = {}
   self.algeType2decl = {}
   self.alge2info = {}
   self.algeInfoList = {}
   self.enumType2decl = {}
   self.declaredMethodSet = {}
   self.adapterMap = {}
   self.prologue = {}
   self.stdoutGlobal = -1
   self.stderrGlobal = -1
end
function convFilter._setmeta( obj )
  setmetatable( obj, { __index = convFilter  } )
end

function convFilter:error( node, mess )

   Util.err( string.format( "%s:%d:%d: wasm: %s", self.streamName, node:get_pos().lineNo, node:get_pos().column, mess) )
end


function convFilter:notSupport( node )

   self:error( node, string.format( "not support -- %s", Nodes.getNodeKindName( node:get_kind() )) )
end

function convFilter:normType( typeInfo )

   local work = typeInfo:get_aliasSrc():get_srcTypeInfo()
   if work:get_nilable() then
      work = work:get_nonnilableType():get_aliasSrc():get_srcTypeInfo()
   end
   return work
end

function convFilter:getGenSrc( typeInfo )

   local work = self:normType( typeInfo )
   local src = work:get_genSrcTypeInfo():get_aliasSrc():get_srcTypeInfo()
   if src ~= work then
      return src
   end
   return work
end


function convFilter:isType( typeInfo, builtin )

   return self:normType( typeInfo ):equals( self.processInfo, builtin )
end

function isScalar( rep )

   return not rep:isRef(  )
end

function convFilter:getRep( typeInfo )

   local orgType = typeInfo:get_aliasSrc():get_srcTypeInfo()
   if orgType:get_nilable() and orgType ~= Ast.builtinTypeNil and orgType:get_nonnilableType() ~= orgType then
      local rep = self:getRep( orgType:get_nonnilableType() )
      if isScalar( rep ) then
         return W.typeAnyRef
      end
      return rep
   end
   local work = self:normType( orgType )
   if work:equals( self.processInfo, Ast.builtinTypeString ) then
      return self.rt:get_strType()
   end
   do
      local _switchExp = work:get_kind()
      if _switchExp == Ast.TypeInfoKind.Prim then
         if work:equals( self.processInfo, Ast.builtinTypeInt ) or work:equals( self.processInfo, Ast.builtinTypeChar ) then
            return W.typeI64
         end
         if work:equals( self.processInfo, Ast.builtinTypeReal ) then
            return W.typeF64
         end
         if work:equals( self.processInfo, Ast.builtinTypeBool ) then
            return W.typeI32
         end
         return W.typeAnyRef
      elseif _switchExp == Ast.TypeInfoKind.Enum then
         do
            local enumType = _lune.__Cast( work, 3, Ast.EnumTypeInfo )
            if enumType ~= nil then
               return self:getRep( enumType:get_valTypeInfo() )
            end
         end
      elseif _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array or _switchExp == Ast.TypeInfoKind.DDD then
         return self.rt:get_listType()
      elseif _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Set then
         return self.rt:get_mapType()
      elseif _switchExp == Ast.TypeInfoKind.Class then
         do
            local info = self:getClassInfo( work )
            if info ~= nil then
               return W.refType( info:get_structType() )
            end
         end
         return self.rt:get_objType()
      elseif _switchExp == Ast.TypeInfoKind.IF then
         return self.rt:get_objType()
      elseif _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
         return W.refType( self:getFuncCloInfo( work ):get_cloType() )
      elseif _switchExp == Ast.TypeInfoKind.Alge then
         do
            local info = self:getAlgeInfo( self:getGenSrc( work ) )
            if info ~= nil then
               return W.refType( info:get_baseType() )
            end
         end
      end
   end
   return W.typeAnyRef
end


function convFilter:getRepList( typeList )

   local list = {}
   for __index, typeInfo in ipairs( typeList ) do
      table.insert( list, self:getRep( typeInfo ) )
   end
   return list
end

function convFilter:getFuncCloInfo( funcType )

   local work = self:normType( funcType )
   return self.rt:getCloInfo( self:getRepList( work:get_argTypeInfoList() ), self:getRepList( work:get_retTypeInfoList() ) )
end

function convFilter:getMethodFuncType( methodType )

   local params = {self.rt:get_objType()}
   for __index, param in ipairs( self:getRepList( methodType:get_argTypeInfoList() ) ) do
      table.insert( params, param )
   end
   return self.mod:getFuncType( params, self:getRepList( methodType:get_retTypeInfoList() ) )
end

function convFilter:getDefaultVal( rep )

   do
      local _switchExp = rep:get_kind()
      if _switchExp == W.ValKind.I64 then
         return W.i64Const( 0 )
      elseif _switchExp == W.ValKind.F64 then
         return W.f64Const( 0.0 )
      elseif _switchExp == W.ValKind.I32 then
         return W.i32Const( 0 )
      elseif _switchExp == W.ValKind.Ref then
         return W.refNull( rep:get_heapType() )
      end
   end
end

function isMethodSym( symbol )

   return symbol:get_kind() == Ast.SymbolKind.Mtd and not symbol:get_staticFlag() and symbol:get_name() ~= "__init" and symbol:get_name() ~= "__free"
end

local mappingMethodSet = {["_toMap"] = true, ["_fromMap"] = true, ["_fromStem"] = true}


function convFilter:getClassInfo( typeInfo )

   local classType = self:getGenSrc( typeInfo )
   do
      local info = self.class2info[classType]
      if info ~= nil then
         return info
      end
   end
   local declNode = self.classType2decl[classType]
   if  nil == declNode then
      local _declNode = declNode
   
      return nil
   end
   
   return self:registerClass( declNode )
end

function convFilter:registerClass( declNode )

   local classType = self:getGenSrc( declNode:get_expType() )
   do
      local info = self.class2info[classType]
      if info ~= nil then
         return info
      end
   end
   local parent = nil
   local superStruct = self.rt:get_tObj()
   local superVt = self.rt:get_tVT()
   if classType:hasBase(  ) then
      local parentInfo = self:getClassInfo( classType:get_baseTypeInfo() )
      if  nil == parentInfo then
         local _parentInfo = parentInfo
      
         self:error( declNode, string.format( "not found base class -- %s", classType:get_baseTypeInfo():getTxt(  )) )
      end
      
      parent = parentInfo
      superStruct = parentInfo:get_structType()
      superVt = parentInfo:get_vtType()
   end
   local structType = self.mod:reserveType(  )
   local vtType = self.mod:reserveType(  )
   local vtGlobal = self.mod:addGlobal( W.refType( vtType ), W.refNull( vtType ) )
   local info = ClassInfo._new(classType, structType, vtType, vtGlobal, parent)
   self.class2info[classType] = info
   table.insert( self.classInfoList, info )
   for __index, memberNode in ipairs( declNode:get_memberList() ) do
      if memberNode:get_staticFlag() then
         local symbol = memberNode:get_symbolInfo():getOrg(  )
         local rep = self:getRep( symbol:get_typeInfo() )
         self.globalSym2loc[symbol] = _lune.newAlge( VarLoc.Global, {self.mod:addGlobal( rep, self:getDefaultVal( rep ) )})
      else
       
         local name = memberNode:get_name().txt
         table.insert( info:get_fieldNames(), name )
         table.insert( info:get_fieldReps(), self:getRep( memberNode:get_symbolInfo():get_typeInfo() ) )
         info:get_name2field()[name] = #info:get_fieldNames() + 1
      end
   end
   local autoMethodSet = {}
   for __index, memberNode in ipairs( declNode:get_memberList() ) do
      if not memberNode:get_staticFlag() then
         if memberNode:get_getterMode() ~= Ast.AccessMode.None then
            autoMethodSet["get_" .. memberNode:get_name().txt]= true
         end
         if memberNode:get_setterMode() ~= Ast.AccessMode.None then
            autoMethodSet["set_" .. memberNode:get_name().txt]= true
         end
      end
   end
   local scope = _lune.unwrap( classType:get_scope())
   do
      local __sorted = {}
      local __map = scope:get_symbol2SymbolInfoMap()
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local symbol = __map[ name ]
         do
            local methodType = symbol:get_typeInfo()
            if (not _lune._Set_has(self.declaredMethodSet, methodType ) and not _lune._Set_has(autoMethodSet, name ) ) or _lune._Set_has(mappingMethodSet, name ) then
            elseif isMethodSym( symbol ) then
               local impl = -1
               if not methodType:get_abstractFlag() then
                  local _
                  local funcType
                  
                  do
                     local slotIndex = info:get_name2slot()[name]
                     if slotIndex ~= nil then
                        funcType = info:get_slotList()[slotIndex]:get_funcType()
                     else
                        funcType = self:getMethodFuncType( methodType )
                     end
                  end
                  local work, _1 = self.mod:addFuncWithType( funcType, 1 + #methodType:get_argTypeInfoList() )
                  impl = work
                  self.mod:setFuncName( impl, string.format( "%s.%s", classType:getTxt(  ), name) )
                  self.func2index[methodType] = impl
               end
               do
                  local slotIndex = info:get_name2slot()[name]
                  if slotIndex ~= nil then
                     info:get_slotList()[slotIndex]:set_impl( impl )
                  else
                     table.insert( info:get_slotList(), SlotInfo._new(name, self:getMethodFuncType( methodType ), methodType, impl) )
                     info:get_name2slot()[name] = #info:get_slotList()
                  end
               end
            elseif symbol:get_kind() == Ast.SymbolKind.Mtd and symbol:get_staticFlag() then
               local _
               local index, _2 = self.mod:addFuncWithType( self:getFuncCloInfo( methodType ):get_funcType(), 1 + #methodType:get_argTypeInfoList() )
               self.func2index[methodType] = index
               self.mod:setFuncName( index, string.format( "%s.%s", classType:getTxt(  ), name) )
            end
         end
      end
   end
   do
      local initSym = scope:get_symbol2SymbolInfoMap()["__init"]
      if initSym ~= nil then
         local _
         local ctorType = self:getMethodFuncType( initSym:get_typeInfo() )
         local ctor, _3 = self.mod:addFuncWithType( ctorType, 1 + #initSym:get_typeInfo():get_argTypeInfoList() )
         self.func2index[initSym:get_typeInfo()] = ctor
         self.mod:setFuncName( ctor, string.format( "%s.__init", classType:getTxt(  )) )
         info:set_ctorFunc( ctor )
         info:set_ctorFuncType( ctorType )
      end
   end
   for __index, ifType in ipairs( classType:get_interfaceList() ) do
      self:addInterface( info, ifType )
   end
   local fields = {W.field( W.typeI32, true ), W.field( W.refType( vtType ), false )}
   for __index, rep in ipairs( info:get_fieldReps() ) do
      table.insert( fields, W.field( rep, true ) )
   end
   self.mod:defineStructType( structType, superStruct, fields )
   local vtFields = {W.field( self.rt:get_anyArrType(), false )}
   for __index, slot in ipairs( info:get_slotList() ) do
      table.insert( vtFields, W.field( W.refType( slot:get_funcType() ), false ) )
   end
   self.mod:defineStructType( vtType, superVt, vtFields )
   return info
end


function convFilter:addInterface( info, ifType )

   local work = self:getGenSrc( ifType )
   for __index, registered in ipairs( info:get_ifList() ) do
      if registered == work then
         return 
      end
   end
   table.insert( info:get_ifList(), work )
   for __index, parentIf in ipairs( work:get_interfaceList() ) do
      self:addInterface( info, parentIf )
   end
end

function convFilter:collectIfMethods( ifType, map )

   for __index, parentIf in ipairs( ifType:get_interfaceList() ) do
      self:collectIfMethods( self:getGenSrc( parentIf ), map )
   end
   do
      local scope = ifType:get_scope()
      if scope ~= nil then
         for name, symbol in pairs( scope:get_symbol2SymbolInfoMap() ) do
            if isMethodSym( symbol ) then
               map[name] = symbol:get_typeInfo()
            end
         end
      end
   end
end


function convFilter:getIfInfo( ifType )

   local work = self:getGenSrc( ifType )
   do
      local info = self.if2info[work]
      if info ~= nil then
         return info
      end
   end
   local methodMap = {}
   self:collectIfMethods( work, methodMap )
   local names = {}
   local funcTypes = {}
   local name2index = {}
   local fields = {}
   do
      local __sorted = {}
      local __map = methodMap
      for __key in pairs( __map ) do
         table.insert( __sorted, __key )
      end
      table.sort( __sorted )
      for __index, name in ipairs( __sorted ) do
         local methodType = __map[ name ]
         do
            table.insert( names, name )
            local funcType = self:getMethodFuncType( methodType )
            table.insert( funcTypes, funcType )
            name2index[name] = #names - 1
            table.insert( fields, W.field( W.refType( funcType ), false ) )
         end
      end
   end
   local itType = self.mod:addStructType( nil, fields )
   local info = IfInfo._new(work, #self.ifInfoList, itType, names, funcTypes, name2index)
   self.if2info[work] = info
   table.insert( self.ifInfoList, info )
   return info
end

function convFilter:getFieldIndex( node, info, name )

   local index = info:get_name2field()[name]
   if  nil == index then
      local _index = index
   
      self:error( node, string.format( "not found member -- %s", name) )
   end
   
   return index
end

function convFilter:getAlgeInfo( algeType )

   do
      local info = self.alge2info[algeType]
      if info ~= nil then
         return info
      end
   end
   local work = _lune.__Cast( algeType, 3, Ast.AlgeTypeInfo )
   if  nil == work then
      local _work = work
   
      return nil
   end
   
   local valNames = {}
   do
      local declNode = self.algeType2decl[algeType]
      if declNode ~= nil then
         for __index, valInfo in ipairs( declNode:get_algeValList() ) do
            table.insert( valNames, valInfo:get_valSym():get_name() )
         end
      else
         do
            local __sorted = {}
            local __map = work:get_valInfoMap()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, name in ipairs( __sorted ) do
               local _1 = __map[ name ]
               do
                  table.insert( valNames, name )
               end
            end
         end
      end
   end
   local baseType = self.mod:addStructType( self.rt:get_tBase(), {W.field( W.typeI32, true ), W.field( W.typeI32, false )} )
   local name2tag = {}
   local name2type = {}
   local name2global = {}
   local nameInsList = {}
   for index, name in ipairs( valNames ) do
      name2tag[name] = index - 1
      local valInfo = _lune.unwrap( work:get_valInfoMap()[name])
      if #valInfo:get_typeList() == 0 then
         name2type[name] = baseType
         local globalIdx = self.mod:addGlobal( W.refType( baseType ), W.refNull( baseType ) )
         name2global[name] = globalIdx
         table.insert( self.prologue, W.globalSet( globalIdx, W.structNew( baseType, {W.i32Const( 0 ), W.i32Const( index - 1 )} ) ) )
      else
       
         local fields = {W.field( W.typeI32, true ), W.field( W.typeI32, false )}
         for __index, paramType in ipairs( valInfo:get_typeList() ) do
            table.insert( fields, W.field( self:getRep( paramType ), false ) )
         end
         name2type[name] = self.mod:addStructType( baseType, fields )
      end
      table.insert( nameInsList, self.rt:strLiteral( string.format( "%s.%s", work:get_rawTxt(), name) ) )
   end
   local namesGlobal = self.mod:addGlobal( self.rt:get_anyArrType(), W.refNull( self.rt:get_tAnyArr() ) )
   table.insert( self.prologue, W.globalSet( namesGlobal, W.arrayNewFixed( self.rt:get_tAnyArr(), nameInsList ) ) )
   local info = AlgeInfo._new(algeType, baseType, valNames, name2tag, name2type, name2global, namesGlobal)
   self.alge2info[algeType] = info
   table.insert( self.algeInfoList, info )
   return info
end

function convFilter:isGlobalSym( symbol )

   return symbol:get_scope() == self.moduleScope
end

function convFilter:cellGet( rep, cell )

   local cellType = self.rt:getCellType( rep )
   local val = W.structGet( cellType, 0, cell )
   if rep:isRef(  ) and rep:get_heapType() ~= W.htAny then
      return W.refCast( rep:get_heapType(), val )
   end
   return val
end

function convFilter:declareVar( symbol, init )

   local orgSym = symbol:getOrg(  )
   local rep = self:getRep( orgSym:get_typeInfo() )
   local val
   
   if init ~= nil then
      val = init
   else
      val = self:getDefaultVal( rep )
   end
   if self:isGlobalSym( orgSym ) then
      local globalIdx
      
      do
         local loc = self.globalSym2loc[orgSym]
         if loc ~= nil then
            do
               local _matchExp = loc
               if _matchExp[1] == VarLoc.Global[1] then
                  local index = _matchExp[2][1]
               
                  globalIdx = index
               else 
                  
                     globalIdx = self.mod:addGlobal( rep, self:getDefaultVal( rep ) )
               end
            end
         else
            globalIdx = self.mod:addGlobal( rep, self:getDefaultVal( rep ) )
            self.globalSym2loc[orgSym] = _lune.newAlge( VarLoc.Global, {globalIdx})
         end
      end
      return W.globalSet( globalIdx, val )
   end
   if orgSym:get_hasAccessFromClosure() then
      local cellType = self.rt:getCellType( rep )
      local index = self.ctx:addLocal( W.refType( cellType ) )
      self.ctx:get_sym2loc()[orgSym] = _lune.newAlge( VarLoc.CellLocal, {index})
      return ls( index, W.structNew( cellType, {val} ) )
   end
   local index = self.ctx:addLocal( rep )
   self.ctx:get_sym2loc()[orgSym] = _lune.newAlge( VarLoc.Local, {index})
   return ls( index, val )
end


function convFilter:findLoc( symbol )

   local orgSym = symbol:getOrg(  )
   do
      local loc = self.ctx:get_sym2loc()[orgSym]
      if loc ~= nil then
         return loc
      end
   end
   return self.globalSym2loc[orgSym]
end


function convFilter:envGet( field )

   return W.structGet( self.ctx:get_envType(), field, lg( self.ctx:get_envLocal() ) )
end

function convFilter:readSym( node, symbol )

   local orgSym = symbol:getOrg(  )
   local rep = self:getRep( orgSym:get_typeInfo() )
   do
      local loc = self:findLoc( orgSym )
      if loc ~= nil then
         do
            local _matchExp = loc
            if _matchExp[1] == VarLoc.Local[1] then
               local index = _matchExp[2][1]
            
               return lg( index )
            elseif _matchExp[1] == VarLoc.CellLocal[1] then
               local index = _matchExp[2][1]
            
               return self:cellGet( rep, lg( index ) )
            elseif _matchExp[1] == VarLoc.Captured[1] then
               local field = _matchExp[2][1]
            
               return self:cellGet( rep, self:envGet( field ) )
            elseif _matchExp[1] == VarLoc.Global[1] then
               local index = _matchExp[2][1]
            
               return W.globalGet( index )
            end
         end
      end
   end
   if orgSym:get_kind() == Ast.SymbolKind.Fun then
      do
         local funcIndex = self.func2index[orgSym:get_typeInfo()]
         if funcIndex ~= nil then
            local info = self:getFuncCloInfo( orgSym:get_typeInfo() )
            return W.structNew( info:get_cloType(), {W.i32Const( 0 ), self.mod:refFunc( funcIndex )} )
         end
      end
   end
   do
      local _switchExp = orgSym:get_name()
      if _switchExp == "__mod__" then
         return self.rt:strLiteral( self.moduleTypeInfo:getFullName( self:get_typeNameCtrl(), self:get_moduleInfoManager(), false ) )
      end
   end
   self:error( node, string.format( "not support symbol -- %s", orgSym:get_name()) )
end

function convFilter:writeSym( node, symbol, val )

   local orgSym = symbol:getOrg(  )
   local rep = self:getRep( orgSym:get_typeInfo() )
   local loc = self:findLoc( orgSym )
   if  nil == loc then
      local _loc = loc
   
      self:error( node, string.format( "not support symbol -- %s", orgSym:get_name()) )
   end
   
   do
      local _matchExp = loc
      if _matchExp[1] == VarLoc.Local[1] then
         local index = _matchExp[2][1]
      
         return ls( index, val )
      elseif _matchExp[1] == VarLoc.CellLocal[1] then
         local index = _matchExp[2][1]
      
         return W.structSet( self.rt:getCellType( rep ), 0, lg( index ), val )
      elseif _matchExp[1] == VarLoc.Captured[1] then
         local field = _matchExp[2][1]
      
         return W.structSet( self.rt:getCellType( rep ), 0, self:envGet( field ), val )
      elseif _matchExp[1] == VarLoc.Global[1] then
         local index = _matchExp[2][1]
      
         return W.globalSet( index, val )
      end
   end
end

function convFilter:isFuncType( typeInfo )

   do
      local _switchExp = self:normType( typeInfo ):get_kind()
      if _switchExp == Ast.TypeInfoKind.Func or _switchExp == Ast.TypeInfoKind.Form or _switchExp == Ast.TypeInfoKind.FormFunc then
         return true
      end
   end
   return false
end

function convFilter:coerceRep( ins, fromRep, toRep )

   if fromRep:equalsType( toRep ) then
      return ins
   end
   if isScalar( fromRep ) then
      if isScalar( toRep ) then
         if fromRep:get_kind() == W.ValKind.I64 and toRep:get_kind() == W.ValKind.F64 then
            return p( W.opF64ConvertI64S, {ins} )
         end
         if fromRep:get_kind() == W.ValKind.F64 and toRep:get_kind() == W.ValKind.I64 then
            return p( W.opI64TruncF64S, {ins} )
         end
         return ins
      end
      local boxed
      
      do
         local _switchExp = fromRep:get_kind()
         if _switchExp == W.ValKind.I64 then
            boxed = self.rt:boxI64( ins )
         elseif _switchExp == W.ValKind.F64 then
            boxed = self.rt:boxF64( ins )
         elseif _switchExp == W.ValKind.I32 or _switchExp == W.ValKind.Ref then
            boxed = self.rt:boxBool( ins )
         end
      end
      return boxed
   end
   if isScalar( toRep ) then
      do
         local _switchExp = toRep:get_kind()
         if _switchExp == W.ValKind.I64 then
            return self.rt:unboxI64( ins )
         elseif _switchExp == W.ValKind.F64 then
            return self.rt:unboxF64( ins )
         elseif _switchExp == W.ValKind.I32 or _switchExp == W.ValKind.Ref then
            return self.rt:unboxBool( ins )
         end
      end
   end
   if toRep:get_heapType() == W.htAny then
      return ins
   end
   return W.refCast( toRep:get_heapType(), ins )
end

function convFilter:coerce( ins, fromType, toType )

   local fromRep = self:getRep( fromType )
   local toRep = self:getRep( toType )
   if fromRep:equalsType( toRep ) then
      return ins
   end
   if self:isFuncType( fromType ) and self:isFuncType( toType ) then
      return self:adaptClosure( ins, fromType, toType )
   end
   return self:coerceRep( ins, fromRep, toRep )
end


function convFilter:coerceVal( val, toType )

   return self:coerce( val:get_ins(), val:get_typeInfo(), toType )
end

function convFilter:adaptClosure( ins, fromType, toType )

   local fromInfo = self:getFuncCloInfo( fromType )
   local toInfo = self:getFuncCloInfo( toType )
   local key = string.format( "%d:%d", fromInfo:get_funcType(), toInfo:get_funcType())
   local adaptType
   
   do
      local work = self.adapterMap[key]
      if work ~= nil then
         adaptType = work
      else
         adaptType = self.mod:addStructType( toInfo:get_cloType(), {W.field( W.typeI32, true ), W.field( W.refType( toInfo:get_funcType() ), false ), W.field( W.refType( fromInfo:get_cloType() ), false )} )
         self.adapterMap[key] = adaptType
         local funcIndex, func = self.mod:addFuncWithType( toInfo:get_funcType(), 1 + #toInfo:get_params() )
         self.adapterMap[key .. ":func"] = funcIndex
         local inner = func:addLocal( W.refType( fromInfo:get_cloType() ) )
         local body = {ls( inner, W.structGet( adaptType, 2, W.refCast( adaptType, lg( 0 ) ) ) )}
         local fromArgs = self:normType( fromType ):get_argTypeInfoList()
         local toArgs = self:normType( toType ):get_argTypeInfoList()
         local toHasDDD = #toArgs > 0 and self:normType( toArgs[#toArgs] ):get_kind() == Ast.TypeInfoKind.DDD
         local args = {lg( inner )}
         for index, fromArg in ipairs( fromArgs ) do
            local fromRep = self:getRep( fromArg )
            if self:normType( fromArg ):get_kind() == Ast.TypeInfoKind.DDD then
               if toHasDDD and index == #toArgs then
                  table.insert( args, lg( #toArgs ) )
               else
                
                  table.insert( args, W.call( self.rt:fnListNew(  ), {W.i32Const( 0 )} ) )
               end
            elseif toHasDDD and index >= #toArgs then
               table.insert( args, self:coerceRep( W.call( self.rt:fnListGet(  ), {lg( #toArgs ), W.i64Const( index - #toArgs + 1 )} ), W.typeAnyRef, fromRep ) )
            elseif index <= #toArgs then
               table.insert( args, self:coerceRep( lg( index ), self:getRep( toArgs[index] ), fromRep ) )
            else
             
               table.insert( args, self:getDefaultVal( fromRep ) )
            end
         end
         table.insert( args, W.structGet( fromInfo:get_cloType(), Rt.fCloCode, lg( inner ) ) )
         table.insert( body, W.callRef( fromInfo:get_funcType(), args ) )
         local fromRets = fromInfo:get_results()
         local toRets = toInfo:get_results()
         local retLocals = {}
         for __index, rep in ipairs( fromRets ) do
            table.insert( retLocals, func:addLocal( rep ) )
         end
         for index = #fromRets, 1, -1 do
            table.insert( body, ls( retLocals[index], W.nop(  ) ) )
         end
         local rets = {}
         local toRetTypes = self:normType( toType ):get_retTypeInfoList()
         for index, toRep in ipairs( toRets ) do
            local isDDD = self:normType( toRetTypes[index] ):get_kind() == Ast.TypeInfoKind.DDD
            if isDDD and index == #toRets then
               local list = func:addLocal( self.rt:get_listType() )
               table.insert( body, ls( list, W.call( self.rt:fnListNew(  ), {W.i32Const( 0 )} ) ) )
               for work = index, #fromRets do
                  table.insert( body, W.call( self.rt:fnListPush(  ), {lg( list ), self:coerceRep( lg( retLocals[work] ), fromRets[work], W.typeAnyRef )} ) )
               end
               table.insert( rets, lg( list ) )
            elseif index <= #fromRets then
               table.insert( rets, self:coerceRep( lg( retLocals[index] ), fromRets[index], toRep ) )
            else
             
               table.insert( rets, self:getDefaultVal( toRep ) )
            end
         end
         table.insert( body, W.seq( rets ) )
         func:setBody( body )
      end
   end
   local funcIndex = _lune.unwrap( self.adapterMap[key .. ":func"])
   return W.structNew( adaptType, {W.i32Const( 0 ), self.mod:refFunc( funcIndex ), W.refCast( fromInfo:get_cloType(), ins )} )
end

function convFilter:toAny( val )

   return self:coerceRep( val:get_ins(), val:get_rep(), W.typeAnyRef )
end

function convFilter:fnTruthy(  )

   do
      local index = self.adapterMap["truthy"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {W.typeI32} )
   self.adapterMap["truthy"] = index
   func:setBody( {W.ifElse( W.refIsNull( lg( 0 ) ), {W.ret( {W.i32Const( 0 )} )}, nil, nil ), W.ifElse( W.refTest( self.rt:get_tBoxBool(), lg( 0 ) ), {W.ret( {self.rt:unboxBool( lg( 0 ) )} )}, nil, nil ), W.i32Const( 1 )} )
   return index
end

function convFilter:toCond( val )

   do
      local _switchExp = val:get_rep():get_kind()
      if _switchExp == W.ValKind.I32 then
         return val:get_ins()
      elseif _switchExp == W.ValKind.I64 or _switchExp == W.ValKind.F64 then
         return W.seq( {W.drop( val:get_ins() ), W.i32Const( 1 )} )
      elseif _switchExp == W.ValKind.Ref then
         if val:get_rep():get_heapType() == W.htAny then
            return W.call( self:fnTruthy(  ), {val:get_ins()} )
         end
         return p( W.opI32Eqz, {W.refIsNull( val:get_ins() )} )
      end
   end
end


function convFilter:lowerCond( node )

   return self:toCond( self:lowerVal( node ) )
end

function convFilter:expTypeOf( node )

   do
      local castNode = _lune.__Cast( node, 3, Nodes.ExpCastNode )
      if castNode ~= nil then
         return castNode:get_castType()
      end
   end
   return node:get_expType()
end


function convFilter:newVal( ins, typeInfo )

   return Val._new(ins, self:getRep( typeInfo ), typeInfo)
end

function convFilter:lowerExp( node )

   return self:coerceVal( self:lowerVal( node ), self:expTypeOf( node ) )
end

function convFilter:lowerExpAs( node, typeInfo )

   return self:coerceVal( self:lowerVal( node ), typeInfo )
end

function convFilter:tempLocal( rep )

   return self.ctx:addLocal( rep )
end

function convFilter:buildList( vals )

   local list = self:tempLocal( self.rt:get_listType() )
   local body = {ls( list, W.call( self.rt:fnListNew(  ), {W.i32Const( #vals )} ) )}
   for __index, val in ipairs( vals ) do
      if self:normType( val:get_typeInfo() ):get_kind() == Ast.TypeInfoKind.DDD then
         table.insert( body, W.call( self.rt:fnListAppend(  ), {lg( list ), val:get_ins()} ) )
      else
       
         table.insert( body, W.call( self.rt:fnListPush(  ), {lg( list ), self:toAny( val )} ) )
      end
   end
   table.insert( body, lg( list ) )
   return W.seq( body )
end

function convFilter:lowerValList( expList, pre )

   local vals = {}
   local mRetIndex
   
   do
      local mRetExp = expList:get_mRetExp()
      if mRetExp ~= nil then
         mRetIndex = mRetExp:get_index()
      else
         mRetIndex = 0
      end
   end
   local castMRet = nil
   local castMRetVals = {}
   local castMRetPos = 0
   for index, exp in ipairs( expList:get_expList() ) do
      local mretInCast = nil
      local castType = exp:get_expType()
      do
         local castNode = _lune.__Cast( exp, 3, Nodes.ExpCastNode )
         if castNode ~= nil then
            if castNode:get_castKind() == Nodes.CastKind.Implicit and castNode:get_exp():get_kind() == Nodes.NodeKind.get_ExpMRet() then
               mretInCast = castNode:get_exp()
               castType = castNode:get_castType()
            end
         end
      end
      if mretInCast then
         local mret = _lune.unwrap( mretInCast)
         if castMRet ~= mret then
            castMRet = mret
            local num = 0
            for __index, other in ipairs( expList:get_expList() ) do
               do
                  local otherCast = _lune.__Cast( other, 3, Nodes.ExpCastNode )
                  if otherCast ~= nil then
                     if otherCast:get_exp() == mret then
                        num = num + 1
                     end
                  end
               end
            end
            castMRetVals = self:expandDDD( self:lowerMulti( mret, pre ), num, pre )
            castMRetPos = 0
         end
         castMRetPos = castMRetPos + 1
         local val
         
         if castMRetPos <= #castMRetVals then
            val = castMRetVals[castMRetPos]
         else
          
            val = Val._new(W.refNull( W.htAny ), W.typeAnyRef, Ast.builtinTypeNil)
         end
         table.insert( vals, Val._new(self:coerceVal( val, castType ), self:getRep( castType ), castType) )
      elseif exp:get_kind() == Nodes.NodeKind.get_ExpAccessMRet() then
      elseif index == mRetIndex or (index == #expList:get_expList() and exp:get_kind() == Nodes.NodeKind.get_ExpMRet() ) then
         for __index, val in ipairs( self:lowerMulti( exp, pre ) ) do
            table.insert( vals, val )
         end
      elseif mRetIndex > 0 then
         local val = self:lowerVal( exp )
         local work = self:tempLocal( val:get_rep() )
         table.insert( pre, ls( work, val:get_ins() ) )
         table.insert( vals, Val._new(lg( work ), val:get_rep(), val:get_typeInfo()) )
      else
       
         table.insert( vals, self:lowerVal( exp ) )
      end
   end
   return vals
end

function convFilter:lowerMulti( node, pre )

   local work = node
   do
      local mretNode = _lune.__Cast( node, 3, Nodes.ExpMRetNode )
      if mretNode ~= nil then
         work = mretNode:get_mRet()
      end
   end
   local callNode = _lune.__Cast( work, 3, Nodes.ExpCallNode )
   if  nil == callNode then
      local _callNode = callNode
   
      local val = self:lowerVal( work )
      local tmp = self:tempLocal( val:get_rep() )
      table.insert( pre, ls( tmp, val:get_ins() ) )
      return {Val._new(lg( tmp ), val:get_rep(), val:get_typeInfo())}
   end
   
   local ins, retTypes = self:lowerCall( callNode )
   table.insert( pre, ins )
   local locals = {}
   for __index, retType in ipairs( retTypes ) do
      table.insert( locals, self:tempLocal( self:getRep( retType ) ) )
   end
   for index = #locals, 1, -1 do
      table.insert( pre, ls( locals[index], W.nop(  ) ) )
   end
   local vals = {}
   for index, retType in ipairs( retTypes ) do
      table.insert( vals, Val._new(lg( locals[index] ), self:getRep( retType ), retType) )
   end
   return vals
end

function convFilter:decodeString( node, txt )

   local quote = txt:sub( 1, 1 )
   if quote == "`" then
      return txt:sub( 4, #txt - 3 )
   end
   local body = txt:sub( 2, #txt - 1 )
   local result = ""
   local pos = 1
   while pos <= #body do
      local ch = body:sub( pos, pos )
      if ch ~= "\\" then
         result = result .. ch
         pos = pos + 1
      else
       
         local next = body:sub( pos + 1, pos + 1 )
         pos = pos + 2
         do
            local _switchExp = next
            if _switchExp == "n" then
               result = result .. "\n"
            elseif _switchExp == "t" then
               result = result .. "\t"
            elseif _switchExp == "r" then
               result = result .. "\r"
            elseif _switchExp == "a" then
               result = result .. W.byte( 7 )
            elseif _switchExp == "b" then
               result = result .. W.byte( 8 )
            elseif _switchExp == "f" then
               result = result .. W.byte( 12 )
            elseif _switchExp == "v" then
               result = result .. W.byte( 11 )
            elseif _switchExp == "\\" or _switchExp == "\"" or _switchExp == "'" then
               result = result .. next
            elseif _switchExp == "x" then
               local hex = body:sub( pos, pos + 1 )
               local val = tonumber( hex, 16 )
               if  nil == val then
                  local _val = val
               
                  self:error( node, string.format( "illegal escape -- %s", txt) )
               end
               
               result = result .. W.byte( math.floor(val) )
               pos = pos + 2
            else 
               
                  local val = 0
                  local count = 0
                  pos = pos - 1
                  while count < 3 do
                     local code = body:byte( pos )
                     if  nil == code then
                        local _code = code
                     
                        break
                     end
                     
                     if code < 0x30 or code > 0x39 then
                        break
                     end
                     val = val * 10 + code - 0x30
                     count = count + 1
                     pos = pos + 1
                  end
                  if count == 0 then
                     self:error( node, string.format( "illegal escape -- %s", txt) )
                  end
                  result = result .. W.byte( val )
            end
         end
      end
   end
   return result
end


function convFilter:lowerLiteralString( node )

   local data = self:decodeString( node, node:get_token().txt )
   local fmt = self.rt:strLiteral( data )
   do
      local param = node:get_orgParam()
      if param ~= nil then
         local pre = {}
         local vals = self:lowerValList( param, pre )
         table.insert( pre, W.call( self.rt:fnFormat(  ), {fmt, self:buildList( vals )} ) )
         return W.seq( pre )
      end
   end
   return fmt
end

function convFilter:lowerEqVals( val1, val2 )

   local rep1 = val1:get_rep()
   local rep2 = val2:get_rep()
   if isScalar( rep1 ) and isScalar( rep2 ) then
      if rep1:get_kind() == W.ValKind.I32 or rep2:get_kind() == W.ValKind.I32 then
         return p( W.opI32Eq, {val1:get_ins(), val2:get_ins()} )
      end
      if rep1:get_kind() == W.ValKind.I64 and rep2:get_kind() == W.ValKind.I64 then
         return p( W.opI64Eq, {val1:get_ins(), val2:get_ins()} )
      end
      return p( W.opF64Eq, {self:coerceRep( val1:get_ins(), rep1, W.typeF64 ), self:coerceRep( val2:get_ins(), rep2, W.typeF64 )} )
   end
   if val1:get_typeInfo() == Ast.builtinTypeNil and rep2:isRef(  ) then
      return W.refIsNull( val2:get_ins() )
   end
   if val2:get_typeInfo() == Ast.builtinTypeNil and rep1:isRef(  ) then
      return W.refIsNull( val1:get_ins() )
   end
   if rep1:equalsType( self.rt:get_strType() ) and rep2:equalsType( self.rt:get_strType() ) then
      return W.call( self.rt:fnStrEq(  ), {val1:get_ins(), val2:get_ins()} )
   end
   return W.call( self.rt:fnEqAny(  ), {self:toAny( val1 ), self:toAny( val2 )} )
end

function convFilter:lowerCompare( node, opTxt, val1, val2 )

   local rep1 = val1:get_rep()
   local rep2 = val2:get_rep()
   if (rep1:get_kind() == W.ValKind.I64 or rep1:get_kind() == W.ValKind.F64 ) and (rep2:get_kind() == W.ValKind.I64 or rep2:get_kind() == W.ValKind.F64 ) then
      if rep1:get_kind() == W.ValKind.I64 and rep2:get_kind() == W.ValKind.I64 then
         return p( _lune.unwrap( i64CmpOpMap[opTxt]), {val1:get_ins(), val2:get_ins()} )
      end
      return p( _lune.unwrap( f64CmpOpMap[opTxt]), {self:coerceRep( val1:get_ins(), rep1, W.typeF64 ), self:coerceRep( val2:get_ins(), rep2, W.typeF64 )} )
   end
   if rep1:equalsType( self.rt:get_strType() ) and rep2:equalsType( self.rt:get_strType() ) then
      return p( _lune.unwrap( strCmpOpMap[opTxt]), {W.call( self.rt:fnStrCmp(  ), {val1:get_ins(), val2:get_ins()} ), W.i32Const( 0 )} )
   end
   local any1 = self:toAny( val1 )
   local any2 = self:toAny( val2 )
   local less = self.rt:fnLessAny(  )
   do
      local _switchExp = opTxt
      if _switchExp == "<" then
         return W.call( less, {any1, any2} )
      elseif _switchExp == ">" then
         return W.call( less, {any2, any1} )
      elseif _switchExp == "<=" then
         return p( W.opI32Eqz, {W.call( less, {any2, any1} )} )
      elseif _switchExp == ">=" then
         return p( W.opI32Eqz, {W.call( less, {any1, any2} )} )
      end
   end
   self:error( node, string.format( "illegal op -- %s", opTxt) )
end

function convFilter:lowerAndOr( node )

   local opTxt = node:get_op().txt
   local resultType = node:get_expType()
   local resultRep = self:getRep( resultType )
   if resultRep:get_kind() == W.ValKind.I32 then
      local cond1 = self:lowerCond( node:get_exp1() )
      local cond2 = self:lowerCond( node:get_exp2() )
      if opTxt == "and" then
         return Val._new(W.ifElse( cond1, {cond2}, {W.i32Const( 0 )}, W.typeI32 ), W.typeI32, resultType)
      end
      return Val._new(W.ifElse( cond1, {W.i32Const( 1 )}, {cond2}, W.typeI32 ), W.typeI32, resultType)
   end
   local val1 = self:lowerVal( node:get_exp1() )
   local work = self:tempLocal( val1:get_rep() )
   local cond = self:toCond( Val._new(lg( work ), val1:get_rep(), val1:get_typeInfo()) )
   local first = self:coerce( lg( work ), val1:get_typeInfo(), resultType )
   local second = self:lowerExpAs( node:get_exp2(), resultType )
   local ins
   
   if opTxt == "and" then
      ins = W.ifElse( cond, {second}, {first}, resultRep )
   else
    
      ins = W.ifElse( cond, {first}, {second}, resultRep )
   end
   return Val._new(W.seq( {ls( work, val1:get_ins() ), ins} ), resultRep, resultType)
end

function convFilter:toStr( val )

   if val:get_rep():equalsType( self.rt:get_strType() ) then
      return val:get_ins()
   end
   do
      local _switchExp = val:get_rep():get_kind()
      if _switchExp == W.ValKind.I64 then
         return W.call( self.rt:fnI64ToStr(  ), {val:get_ins()} )
      elseif _switchExp == W.ValKind.F64 then
         return W.call( self.rt:fnRealToStr(  ), {val:get_ins()} )
      elseif _switchExp == W.ValKind.I32 then
         return W.ifElse( val:get_ins(), {self.rt:strLiteral( "true" )}, {self.rt:strLiteral( "false" )}, self.rt:get_strType() )
      elseif _switchExp == W.ValKind.Ref then
         return W.call( self.rt:fnToStr(  ), {val:get_ins()} )
      end
   end
end

function convFilter:lowerOp2( node )

   local opTxt = node:get_op().txt
   local resultType = node:get_expType()
   do
      local _switchExp = opTxt
      if _switchExp == "and" or _switchExp == "or" then
         return self:lowerAndOr( node )
      elseif _switchExp == ".." then
         return self:newVal( W.call( self.rt:fnStrConcat(  ), {self:toStr( self:lowerVal( node:get_exp1() ) ), self:toStr( self:lowerVal( node:get_exp2() ) )} ), resultType )
      elseif _switchExp == "==" or _switchExp == "~=" then
         local eq = self:lowerEqVals( self:lowerVal( node:get_exp1() ), self:lowerVal( node:get_exp2() ) )
         if opTxt == "==" then
            return self:newVal( eq, resultType )
         end
         return self:newVal( p( W.opI32Eqz, {eq} ), resultType )
      elseif _switchExp == "<" or _switchExp == "<=" or _switchExp == ">" or _switchExp == ">=" then
         return self:newVal( self:lowerCompare( node, opTxt, self:lowerVal( node:get_exp1() ), self:lowerVal( node:get_exp2() ) ), resultType )
      end
   end
   local resultRep = self:getRep( resultType )
   if resultRep:get_kind() == W.ValKind.I64 then
      local exp1 = self:lowerExpAs( node:get_exp1(), Ast.builtinTypeInt )
      local exp2 = self:lowerExpAs( node:get_exp2(), Ast.builtinTypeInt )
      local ins = nil
      do
         local _switchExp = opTxt
         if _switchExp == "+" then
            ins = p( W.opI64Add, {exp1, exp2} )
         elseif _switchExp == "-" then
            ins = p( W.opI64Sub, {exp1, exp2} )
         elseif _switchExp == "*" then
            ins = p( W.opI64Mul, {exp1, exp2} )
         elseif _switchExp == "/" then
            ins = W.call( self:fnDivI64(  ), {exp1, exp2} )
         elseif _switchExp == "%" then
            ins = W.call( self:fnModI64(  ), {exp1, exp2} )
         elseif _switchExp == "&" then
            ins = p( W.opI64And, {exp1, exp2} )
         elseif _switchExp == "|" then
            ins = p( W.opI64Or, {exp1, exp2} )
         elseif _switchExp == "~" then
            ins = p( W.opI64Xor, {exp1, exp2} )
         elseif _switchExp == "|<<" then
            ins = W.call( self:fnShift( true ), {exp1, exp2} )
         elseif _switchExp == "|>>" then
            ins = W.call( self:fnShift( false ), {exp1, exp2} )
         end
      end
      if ins ~= nil then
         return Val._new(ins, resultRep, resultType)
      end
   end
   if resultRep:get_kind() == W.ValKind.F64 then
      local exp1 = self:lowerExpAs( node:get_exp1(), Ast.builtinTypeReal )
      local exp2 = self:lowerExpAs( node:get_exp2(), Ast.builtinTypeReal )
      local ins = nil
      do
         local _switchExp = opTxt
         if _switchExp == "+" then
            ins = p( W.opF64Add, {exp1, exp2} )
         elseif _switchExp == "-" then
            ins = p( W.opF64Sub, {exp1, exp2} )
         elseif _switchExp == "*" then
            ins = p( W.opF64Mul, {exp1, exp2} )
         elseif _switchExp == "/" then
            ins = p( W.opF64Div, {exp1, exp2} )
         elseif _switchExp == "%" then
            ins = W.call( self:fnModF64(  ), {exp1, exp2} )
         end
      end
      if ins ~= nil then
         return Val._new(ins, resultRep, resultType)
      end
   end
   self:error( node, string.format( "not support op -- %s (%s)", opTxt, resultType:getTxt(  )) )
end

function convFilter:lowerOp1( node )

   local opTxt = node:get_op().txt
   local resultType = node:get_expType()
   local val = self:lowerVal( node:get_exp() )
   do
      local _switchExp = opTxt
      if _switchExp == "-" then
         if val:get_rep():get_kind() == W.ValKind.I64 then
            return self:newVal( p( W.opI64Sub, {W.i64Const( 0 ), val:get_ins()} ), resultType )
         end
         if val:get_rep():get_kind() == W.ValKind.F64 then
            return self:newVal( p( W.opF64Neg, {val:get_ins()} ), resultType )
         end
      elseif _switchExp == "not" then
         return self:newVal( p( W.opI32Eqz, {self:toCond( val )} ), resultType )
      elseif _switchExp == "#" then
         if val:get_rep():equalsType( self.rt:get_strType() ) then
            return self:newVal( p( W.opI64ExtendI32U, {W.arrayLen( val:get_ins() )} ), resultType )
         end
         if val:get_rep():equalsType( self.rt:get_listType() ) then
            return self:newVal( p( W.opI64ExtendI32U, {W.structGet( self.rt:get_tList(), Rt.fListLen, val:get_ins() )} ), resultType )
         end
      elseif _switchExp == "~" then
         if val:get_rep():get_kind() == W.ValKind.I64 then
            return self:newVal( p( W.opI64Xor, {val:get_ins(), W.i64Const( -1 )} ), resultType )
         end
      end
   end
   self:error( node, string.format( "not support op -- %s (%s)", opTxt, val:get_typeInfo():getTxt(  )) )
end

function convFilter:fnDivI64(  )

   do
      local index = self.adapterMap["divI64"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeI64, W.typeI64}, {W.typeI64} )
   self.adapterMap["divI64"] = index
   local q = func:addLocal( W.typeI64 )
   func:setBody( {ls( q, p( W.opI64DivS, {lg( 0 ), lg( 1 )} ) ), W.ifElse( p( W.opI32And, {p( W.opI64Ne, {p( W.opI64RemS, {lg( 0 ), lg( 1 )} ), W.i64Const( 0 )} ), p( W.opI32Ne, {p( W.opI64LtS, {lg( 0 ), W.i64Const( 0 )} ), p( W.opI64LtS, {lg( 1 ), W.i64Const( 0 )} )} )} ), {ls( q, p( W.opI64Sub, {lg( q ), W.i64Const( 1 )} ) )}, nil, nil ), lg( q )} )
   return index
end

function convFilter:fnModI64(  )

   do
      local index = self.adapterMap["modI64"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeI64, W.typeI64}, {W.typeI64} )
   self.adapterMap["modI64"] = index
   local r = func:addLocal( W.typeI64 )
   func:setBody( {ls( r, p( W.opI64RemS, {lg( 0 ), lg( 1 )} ) ), W.ifElse( p( W.opI32And, {p( W.opI64Ne, {lg( r ), W.i64Const( 0 )} ), p( W.opI32Ne, {p( W.opI64LtS, {lg( r ), W.i64Const( 0 )} ), p( W.opI64LtS, {lg( 1 ), W.i64Const( 0 )} )} )} ), {ls( r, p( W.opI64Add, {lg( r ), lg( 1 )} ) )}, nil, nil ), lg( r )} )
   return index
end

function convFilter:fnModF64(  )

   do
      local index = self.adapterMap["modF64"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeF64, W.typeF64}, {W.typeF64} )
   self.adapterMap["modF64"] = index
   func:setBody( {p( W.opF64Sub, {lg( 0 ), p( W.opF64Mul, {p( W.opF64Floor, {p( W.opF64Div, {lg( 0 ), lg( 1 )} )} ), lg( 1 )} )} )} )
   return index
end

function convFilter:fnShift( left )

   local name
   
   local code
   
   local revCode
   
   if left then
      name = "shiftL"
      code = W.opI64Shl
      revCode = W.opI64ShrU
   else
    
      name = "shiftR"
      code = W.opI64ShrU
      revCode = W.opI64Shl
   end
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeI64, W.typeI64}, {W.typeI64} )
   self.adapterMap[name] = index
   func:setBody( {W.ifElse( p( W.opI32Or, {p( W.opI64LeS, {lg( 1 ), W.i64Const( -64 )} ), p( W.opI64GeS, {lg( 1 ), W.i64Const( 64 )} )} ), {W.ret( {W.i64Const( 0 )} )}, nil, nil ), W.ifElse( p( W.opI64LtS, {lg( 1 ), W.i64Const( 0 )} ), {W.ret( {p( revCode, {lg( 0 ), p( W.opI64Sub, {W.i64Const( 0 ), lg( 1 )} )} )} )}, nil, nil ), p( code, {lg( 0 ), lg( 1 )} )} )
   return index
end

function convFilter:lowerCast( node )

   local castType = node:get_castType()
   local val = self:lowerVal( node:get_exp() )
   local castRep = self:getRep( castType )
   if node:get_castKind() == Nodes.CastKind.Normal then
      local work = self:tempLocal( W.typeAnyRef )
      local anyVal = self:toAny( val )
      local dstType = self:normType( castType )
      local test
      
      if dstType:equals( self.processInfo, Ast.builtinTypeInt ) then
         test = W.call( self:fnIsIntVal(  ), {lg( work )} )
      elseif dstType:equals( self.processInfo, Ast.builtinTypeReal ) then
         test = p( W.opI32Or, {W.refTest( self.rt:get_tBoxF64(), lg( work ) ), W.refTest( self.rt:get_tBoxI64(), lg( work ) )} )
      elseif dstType:equals( self.processInfo, Ast.builtinTypeBool ) then
         test = W.refTest( self.rt:get_tBoxBool(), lg( work ) )
      else
       
         local dstRep = self:getRep( dstType )
         if dstRep:get_heapType() == W.htAny then
            test = p( W.opI32Eqz, {W.refIsNull( lg( work ) )} )
         else
          
            test = W.refTest( dstRep:get_heapType(), lg( work ) )
         end
      end
      local resultType = castType:get_nilableTypeInfo()
      local resultRep = self:getRep( resultType )
      local okVal = lg( work )
      if dstType:equals( self.processInfo, Ast.builtinTypeReal ) then
         okVal = self.rt:boxF64( self.rt:unboxF64( lg( work ) ) )
      elseif dstType:equals( self.processInfo, Ast.builtinTypeInt ) then
         okVal = self.rt:boxI64( p( W.opI64TruncF64S, {self.rt:unboxF64( lg( work ) )} ) )
      end
      return Val._new(W.seq( {ls( work, anyVal ), W.ifElse( test, {self:coerceRep( okVal, W.typeAnyRef, resultRep )}, {W.refNull( resultRep:get_heapType() )}, resultRep )} ), resultRep, resultType)
   end
   return Val._new(self:coerce( val:get_ins(), val:get_typeInfo(), castType ), castRep, castType)
end

function convFilter:fnIsIntVal(  )

   do
      local index = self.adapterMap["isIntVal"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {W.typeI32} )
   self.adapterMap["isIntVal"] = index
   local fval = func:addLocal( W.typeF64 )
   func:setBody( {W.ifElse( W.refTest( self.rt:get_tBoxI64(), lg( 0 ) ), {W.ret( {W.i32Const( 1 )} )}, nil, nil ), W.ifElse( p( W.opI32Eqz, {W.refTest( self.rt:get_tBoxF64(), lg( 0 ) )} ), {W.ret( {W.i32Const( 0 )} )}, nil, nil ), ls( fval, self.rt:unboxF64( lg( 0 ) ) ), p( W.opI32And, {p( W.opF64Eq, {p( W.opF64Trunc, {lg( fval )} ), lg( fval )} ), p( W.opF64Lt, {p( W.opF64Abs, {lg( fval )} ), W.f64Const( 9.2e18 )} )} )} )
   return index
end

function convFilter:lowerUnwrap( node )

   local resultType = node:get_expType()
   local resultRep = self:getRep( resultType )
   local val = self:lowerVal( node:get_exp() )
   local work = self:tempLocal( val:get_rep() )
   local elseBody
   
   do
      local defNode = node:get_default()
      if defNode ~= nil then
         elseBody = self:lowerExpAs( defNode, resultType )
      else
         elseBody = self.rt:error( self.rt:strLiteral( "unwrap nil" ) )
      end
   end
   return Val._new(W.seq( {ls( work, val:get_ins() ), W.ifElse( W.refIsNull( lg( work ) ), {elseBody}, {self:coerce( lg( work ), val:get_typeInfo(), resultType )}, resultRep )} ), resultRep, resultType)
end

function convFilter:lowerLiteralColl( node, expList, isSet )

   local pre = {}
   local vals = {}
   if expList ~= nil then
      vals = self:lowerValList( expList, pre )
   end
   if isSet then
      local set = self:tempLocal( self.rt:get_mapType() )
      table.insert( pre, ls( set, self.rt:newMap(  ) ) )
      for __index, val in ipairs( vals ) do
         local work = self:tempLocal( W.typeAnyRef )
         table.insert( pre, ls( work, self:toAny( val ) ) )
         table.insert( pre, W.call( self.rt:fnMapSet(  ), {lg( set ), lg( work ), lg( work )} ) )
      end
      table.insert( pre, lg( set ) )
   else
    
      table.insert( pre, self:buildList( vals ) )
   end
   return self:newVal( W.seq( pre ), node:get_expType() )
end


function convFilter:lowerLiteralMap( node )

   local map = self:tempLocal( self.rt:get_mapType() )
   local body = {ls( map, self.rt:newMap(  ) )}
   for __index, pair in ipairs( node:get_pairList() ) do
      table.insert( body, W.call( self.rt:fnMapSet(  ), {lg( map ), self:toAny( self:lowerVal( pair:get_key() ) ), self:toAny( self:lowerVal( pair:get_val() ) )} ) )
   end
   table.insert( body, lg( map ) )
   return self:newVal( W.seq( body ), node:get_expType() )
end

function convFilter:getObjType( typeInfo )

   local work = self:normType( typeInfo )
   if work:get_kind() == Ast.TypeInfoKind.Alternate and work:hasBase(  ) then
      return self:getObjType( work:get_baseTypeInfo() )
   end
   return work
end

function convFilter:lowerMemberRef( node, prefix, name )

   local objType = self:getObjType( prefix:get_expType() )
   local info = self:getClassInfo( objType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", objType:getTxt(  )) )
   end
   
   local fieldIndex = self:getFieldIndex( node, info, name )
   local rep = info:get_fieldReps()[fieldIndex - 1]
   local obj = self:lowerExpAs( prefix, info:get_typeInfo() )
   local memberType = self:getMemberType( info, name )
   return Val._new(W.structGet( info:get_structType(), fieldIndex, W.refCast( info:get_structType(), obj ) ), rep, memberType)
end

function convFilter:getMemberType( info, name )

   local work = info
   while true do
      local classInfo = work
      if  nil == classInfo then
         local _classInfo = classInfo
      
         break
      end
      
      do
         local scope = classInfo:get_typeInfo():get_scope()
         if scope ~= nil then
            do
               local symbol = scope:get_symbol2SymbolInfoMap()[name]
               if symbol ~= nil then
                  return symbol:get_typeInfo()
               end
            end
         end
      end
      work = classInfo:get_parent()
   end
   return Ast.builtinTypeStem_
end

function convFilter:lowerRefField( node )

   local prefix = node:get_prefix()
   local fieldName = node:get_field().txt
   local prefixType = self:normType( prefix:get_expType() )
   do
      local symbol = node:get_symbolInfo()
      if symbol ~= nil then
         local orgSym = symbol:getOrg(  )
         do
            local _exp = self.globalSym2loc[orgSym]
            if _exp ~= nil then
               return Val._new(self:readSym( node, orgSym ), self:getRep( orgSym:get_typeInfo() ), orgSym:get_typeInfo())
            end
         end
         if orgSym:get_typeInfo():get_kind() == Ast.TypeInfoKind.Enum or prefixType:get_kind() == Ast.TypeInfoKind.Enum then
            do
               local enumType = _lune.__Cast( self:normType( orgSym:get_typeInfo() ), 3, Ast.EnumTypeInfo )
               if enumType ~= nil then
                  do
                     local valInfo = enumType:get_name2EnumValInfo()[fieldName]
                     if valInfo ~= nil then
                        return self:lowerEnumVal( enumType, valInfo )
                     end
                  end
               end
            end
         end
         if orgSym:get_name() == "stdout" or orgSym:get_name() == "stderr" then
            if prefix:get_kind() == Nodes.NodeKind.get_ExpRef() then
               return Val._new(W.globalGet( self:getStdStream( orgSym:get_name() == "stdout" ) ), self.rt:get_objType(), orgSym:get_typeInfo())
            end
         end
      end
   end
   do
      local _switchExp = prefixType:get_kind()
      if _switchExp == Ast.TypeInfoKind.Map then
         return Val._new(W.call( self.rt:fnMapGet(  ), {self:lowerExp( prefix ), self.rt:strLiteral( fieldName )} ), W.typeAnyRef, Ast.builtinTypeStem_)
      elseif _switchExp == Ast.TypeInfoKind.Class or _switchExp == Ast.TypeInfoKind.Alternate then
         return self:lowerMemberRef( node, prefix, fieldName )
      end
   end
   self:error( node, string.format( "not support field -- %s.%s", prefixType:getTxt(  ), fieldName) )
end

function convFilter:lowerEnumVal( enumType, valInfo )

   do
      local _matchExp = valInfo:get_val()
      if _matchExp[1] == Ast.EnumLiteral.Int[1] then
         local val = _matchExp[2][1]
      
         return Val._new(W.i64Const( val ), W.typeI64, enumType)
      elseif _matchExp[1] == Ast.EnumLiteral.Real[1] then
         local val = _matchExp[2][1]
      
         return Val._new(W.f64Const( val ), W.typeF64, enumType)
      elseif _matchExp[1] == Ast.EnumLiteral.Str[1] then
         local val = _matchExp[2][1]
      
         return Val._new(self.rt:strLiteral( val ), self.rt:get_strType(), enumType)
      end
   end
end

function convFilter:getStdStream( isStdout )

   if self.stdoutGlobal < 0 then
      self.stdoutGlobal = self.mod:addGlobal( self.rt:get_objType(), W.refNull( self.rt:get_tObj() ) )
      self.stderrGlobal = self.mod:addGlobal( self.rt:get_objType(), W.refNull( self.rt:get_tObj() ) )
      local vt = W.structNew( self.rt:get_tVT(), {W.refNull( self.rt:get_tAnyArr() )} )
      table.insert( self.prologue, W.globalSet( self.stdoutGlobal, W.structNew( self.rt:get_tObj(), {W.i32Const( 0 ), vt} ) ) )
      table.insert( self.prologue, W.globalSet( self.stderrGlobal, W.structNew( self.rt:get_tObj(), {W.i32Const( 0 ), vt} ) ) )
   end
   if isStdout then
      return self.stdoutGlobal
   end
   return self.stderrGlobal
end

function convFilter:lowerRefItem( node )

   local valType = self:normType( node:get_val():get_expType() )
   local index = node:get_index()
   if  nil == index then
      local _index = index
   
      do
         local symbol = node:get_symbol()
         if symbol ~= nil then
            if valType:get_kind() == Ast.TypeInfoKind.Map then
               return Val._new(W.call( self.rt:fnMapGet(  ), {self:lowerExp( node:get_val() ), self.rt:strLiteral( symbol )} ), W.typeAnyRef, Ast.builtinTypeStem_)
            end
         end
      end
      self:notSupport( node )
   end
   
   do
      local _switchExp = valType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         return Val._new(W.call( self.rt:fnListGet(  ), {self:lowerExp( node:get_val() ), self:lowerExpAs( index, Ast.builtinTypeInt )} ), W.typeAnyRef, Ast.builtinTypeStem_)
      elseif _switchExp == Ast.TypeInfoKind.Map then
         return Val._new(W.call( self.rt:fnMapGet(  ), {self:lowerExp( node:get_val() ), self:toAny( self:lowerVal( index ) )} ), W.typeAnyRef, Ast.builtinTypeStem_)
      end
   end
   if self:isType( valType, Ast.builtinTypeString ) then
      local strVal = self:tempLocal( self.rt:get_strType() )
      return self:newVal( W.seq( {ls( strVal, self:lowerExp( node:get_val() ) ), p( W.opI64ExtendI32U, {W.arrayGetU( self.rt:get_tStr(), lg( strVal ), p( W.opI32WrapI64, {p( W.opI64Sub, {self:lowerExpAs( index, Ast.builtinTypeInt ), W.i64Const( 1 )} )} ) )} )} ), Ast.builtinTypeInt )
   end
   self:error( node, string.format( "not support index -- %s", valType:getTxt(  )) )
end

function convFilter:lowerNewAlgeVal( node )

   local algeType = self:getGenSrc( node:get_valInfo():get_algeType() )
   local info = self:getAlgeInfo( algeType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found alge -- %s", algeType:getTxt(  )) )
   end
   
   local name = node:get_valInfo():get_name()
   do
      local globalIdx = info:get_name2global()[name]
      if globalIdx ~= nil then
         return Val._new(W.globalGet( globalIdx ), W.refType( info:get_baseType() ), algeType)
      end
   end
   local valType = _lune.unwrap( info:get_name2type()[name])
   local args = {W.i32Const( 0 ), W.i32Const( _lune.unwrap( info:get_name2tag()[name]) )}
   for index, param in ipairs( node:get_paramList() ) do
      table.insert( args, self:lowerExpAs( param, node:get_valInfo():get_typeList()[index] ) )
   end
   return Val._new(W.structNew( valType, args ), W.refType( info:get_baseType() ), algeType)
end

function convFilter:lowerAnonymousFunc( node )

   local funcType = node:get_expType()
   local funcIndex = self:registerFunc( node:get_declInfo(), funcType )
   return self:newVal( self:compileFunc( node, node:get_declInfo(), funcType, funcIndex ), funcType )
end

function convFilter:lowerVal( node )

   local expType = node:get_expType()
   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_LiteralInt() then
         local literal = _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralIntNode ))
         return self:newVal( W.i64Const( literal:get_num() ), Ast.builtinTypeInt )
      elseif _switchExp == Nodes.NodeKind.get_LiteralChar() then
         local literal = _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralCharNode ))
         return self:newVal( W.i64Const( literal:get_num() ), Ast.builtinTypeInt )
      elseif _switchExp == Nodes.NodeKind.get_LiteralReal() then
         local literal = _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralRealNode ))
         return self:newVal( W.f64Const( literal:get_num() ), Ast.builtinTypeReal )
      elseif _switchExp == Nodes.NodeKind.get_LiteralBool() then
         local literal = _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralBoolNode ))
         if literal:get_token().txt == "true" then
            return self:newVal( W.i32Const( 1 ), Ast.builtinTypeBool )
         end
         return self:newVal( W.i32Const( 0 ), Ast.builtinTypeBool )
      elseif _switchExp == Nodes.NodeKind.get_LiteralNil() or _switchExp == Nodes.NodeKind.get_Abbr() then
         return Val._new(W.refNull( W.htAny ), W.typeAnyRef, Ast.builtinTypeNil)
      elseif _switchExp == Nodes.NodeKind.get_LiteralString() then
         return self:newVal( self:lowerLiteralString( _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralStringNode )) ), Ast.builtinTypeString )
      elseif _switchExp == Nodes.NodeKind.get_LiteralList() then
         return self:lowerLiteralColl( node, (_lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralListNode )) ):get_expList(), false )
      elseif _switchExp == Nodes.NodeKind.get_LiteralArray() then
         return self:lowerLiteralColl( node, (_lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralArrayNode )) ):get_expList(), false )
      elseif _switchExp == Nodes.NodeKind.get_LiteralSet() then
         return self:lowerLiteralColl( node, (_lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralSetNode )) ):get_expList(), true )
      elseif _switchExp == Nodes.NodeKind.get_LiteralMap() then
         return self:lowerLiteralMap( _lune.unwrap( _lune.__Cast( node, 3, Nodes.LiteralMapNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpRef() then
         local refNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpRefNode ))
         local symbol = refNode:get_symbolInfo():getOrg(  )
         return Val._new(self:readSym( node, symbol ), self:getRep( symbol:get_typeInfo() ), symbol:get_typeInfo())
      elseif _switchExp == Nodes.NodeKind.get_ExpParen() then
         return self:lowerVal( (_lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpParenNode )) ):get_exp() )
      elseif _switchExp == Nodes.NodeKind.get_ExpOp2() then
         return self:lowerOp2( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpOp2Node )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpOp1() then
         return self:lowerOp1( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpOp1Node )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpCast() then
         return self:lowerCast( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpCastNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpUnwrap() then
         return self:lowerUnwrap( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpUnwrapNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpCall() then
         local ins, retTypes = self:lowerCall( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpCallNode )) )
         if #retTypes == 0 then
            return Val._new(ins, W.typeI32, Ast.builtinTypeNone)
         end
         if #retTypes == 1 then
            if self:normType( retTypes[1] ):get_kind() == Ast.TypeInfoKind.DDD then
               return Val._new(W.call( self.rt:fnListGet(  ), {ins, W.i64Const( 1 )} ), W.typeAnyRef, Ast.builtinTypeStem_)
            end
            return self:newVal( ins, retTypes[1] )
         end
         local pre = {}
         local vals = self:lowerMulti( node, pre )
         table.insert( pre, vals[1]:get_ins() )
         return Val._new(W.seq( pre ), vals[1]:get_rep(), vals[1]:get_typeInfo())
      elseif _switchExp == Nodes.NodeKind.get_ExpMultiTo1() then
         local pre = {}
         local vals = self:lowerMulti( (_lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpMultiTo1Node )) ):get_exp(), pre )
         if self:normType( vals[1]:get_typeInfo() ):get_kind() == Ast.TypeInfoKind.DDD then
            table.insert( pre, W.call( self.rt:fnListGet(  ), {vals[1]:get_ins(), W.i64Const( 1 )} ) )
            return Val._new(W.seq( pre ), W.typeAnyRef, Ast.builtinTypeStem_)
         end
         table.insert( pre, vals[1]:get_ins() )
         return Val._new(W.seq( pre ), vals[1]:get_rep(), vals[1]:get_typeInfo())
      elseif _switchExp == Nodes.NodeKind.get_ExpToDDD() then
         local dddNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpToDDDNode ))
         local pre = {}
         local vals = self:lowerValList( dddNode:get_expList(), pre )
         table.insert( pre, self:buildList( vals ) )
         return Val._new(W.seq( pre ), self.rt:get_listType(), expType)
      elseif _switchExp == Nodes.NodeKind.get_ExpNew() then
         return self:lowerNew( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpNewNode )) )
      elseif _switchExp == Nodes.NodeKind.get_RefField() then
         return self:lowerRefField( _lune.unwrap( _lune.__Cast( node, 3, Nodes.RefFieldNode )) )
      elseif _switchExp == Nodes.NodeKind.get_GetField() then
         return self:lowerGetField( _lune.unwrap( _lune.__Cast( node, 3, Nodes.GetFieldNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpRefItem() then
         return self:lowerRefItem( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpRefItemNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpOmitEnum() then
         local omitNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpOmitEnumNode ))
         return self:lowerEnumVal( omitNode:get_enumTypeInfo(), omitNode:get_valInfo() )
      elseif _switchExp == Nodes.NodeKind.get_NewAlgeVal() then
         return self:lowerNewAlgeVal( _lune.unwrap( _lune.__Cast( node, 3, Nodes.NewAlgeValNode )) )
      elseif _switchExp == Nodes.NodeKind.get_DeclFunc() then
         return self:lowerAnonymousFunc( _lune.unwrap( _lune.__Cast( node, 3, Nodes.DeclFuncNode )) )
      elseif _switchExp == Nodes.NodeKind.get_ExpCallSuper() then
         local ins, retTypes = self:lowerCallSuper( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpCallSuperNode )) )
         if #retTypes == 0 then
            return Val._new(ins, W.typeI32, Ast.builtinTypeNone)
         end
         return self:newVal( ins, retTypes[1] )
      elseif _switchExp == Nodes.NodeKind.get_ExpMacroExp() then
         local macroNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpMacroExpNode ))
         if #macroNode:get_stmtList() == 1 then
            local stmt = macroNode:get_stmtList()[1]
            do
               local stmtExp = _lune.__Cast( stmt, 3, Nodes.StmtExpNode )
               if stmtExp ~= nil then
                  return self:lowerVal( stmtExp:get_exp() )
               end
            end
            return self:lowerVal( stmt )
         end
      end
   end
   self:notSupport( node )
end

function convFilter:lowerArgs( argList, paramTypes, pre )

   local vals = {}
   if argList ~= nil then
      vals = self:lowerValList( argList, pre )
   end
   local args = {}
   for index, paramType in ipairs( paramTypes ) do
      if self:normType( paramType ):get_kind() == Ast.TypeInfoKind.DDD then
         if #vals == index and self:normType( vals[index]:get_typeInfo() ):get_kind() == Ast.TypeInfoKind.DDD then
            table.insert( args, vals[index]:get_ins() )
         else
          
            local rest = {}
            for work = index, #vals do
               table.insert( rest, vals[work] )
            end
            table.insert( args, self:buildList( rest ) )
         end
         break
      end
      if index <= #vals then
         table.insert( args, self:coerceVal( vals[index], paramType ) )
      else
       
         table.insert( args, self:getDefaultVal( self:getRep( paramType ) ) )
      end
   end
   return args
end


function append( list, items )

   for __index, item in ipairs( items ) do
      table.insert( list, item )
   end
end

function convFilter:callClosure( cloVal, funcType, argList, pre )

   local info = self:getFuncCloInfo( funcType )
   local work = self:tempLocal( W.refType( info:get_cloType() ) )
   table.insert( pre, ls( work, W.refCast( info:get_cloType(), cloVal ) ) )
   local args = {lg( work )}
   append( args, self:lowerArgs( argList, self:normType( funcType ):get_argTypeInfoList(), pre ) )
   table.insert( args, W.structGet( info:get_cloType(), Rt.fCloCode, lg( work ) ) )
   return W.callRef( info:get_funcType(), args )
end

function convFilter:lowerPrint( node )

   local pre = {}
   local vals = {}
   do
      local argList = node:get_argList()
      if argList ~= nil then
         if #argList:get_expList() == 1 then
            do
               local dddNode = _lune.__Cast( argList:get_expList()[1], 3, Nodes.ExpToDDDNode )
               if dddNode ~= nil then
                  vals = self:lowerValList( dddNode:get_expList(), pre )
               else
                  vals = self:lowerValList( argList, pre )
               end
            end
         else
          
            vals = self:lowerValList( argList, pre )
         end
      end
   end
   table.insert( pre, W.call( self.rt:fnPrintList(  ), {self:buildList( vals )} ) )
   return W.seq( pre )
end

function convFilter:lowerCallRef( node, refNode, pre )

   local symbol = refNode:get_symbolInfo():getOrg(  )
   local funcType = symbol:get_typeInfo()
   local builtin = self.builtinFuncs
   if funcType:equals( self.processInfo, builtin.lns_print ) then
      return self:lowerPrint( node ), {}
   end
   if funcType:equals( self.processInfo, builtin.lns_error ) then
      local args = self:lowerArgs( node:get_argList(), funcType:get_argTypeInfoList(), pre )
      table.insert( pre, self.rt:error( args[1] ) )
      return W.seq( pre ), {}
   end
   if funcType:equals( self.processInfo, builtin.lns_tostring ) then
      local vals = {}
      do
         local argList = node:get_argList()
         if argList ~= nil then
            vals = self:lowerValList( argList, pre )
         end
      end
      table.insert( pre, self:toStr( vals[1] ) )
      return W.seq( pre ), {Ast.builtinTypeString}
   end
   if symbol:get_kind() == Ast.SymbolKind.Fun then
      do
         local funcIndex = self.func2index[funcType]
         if funcIndex ~= nil then
            local cloArg
            
            if self:findLoc( symbol ) then
               cloArg = self:readSym( refNode, symbol )
            else
             
               cloArg = W.refNull( self.rt:get_tClo() )
            end
            local args = {cloArg}
            append( args, self:lowerArgs( node:get_argList(), funcType:get_argTypeInfoList(), pre ) )
            table.insert( pre, W.call( funcIndex, args ) )
            return W.seq( pre ), funcType:get_retTypeInfoList()
         end
      end
   end
   table.insert( pre, self:callClosure( self:readSym( refNode, symbol ), funcType, node:get_argList(), pre ) )
   return W.seq( pre ), self:normType( funcType ):get_retTypeInfoList()
end

function convFilter:lowerArgAny( node, index )

   local argList = _lune.unwrap( node:get_argList())
   return self:toAny( self:lowerVal( argList:get_expList()[index] ) )
end

function convFilter:lowerArgAs( node, index, typeInfo )

   local argList = node:get_argList()
   if  nil == argList then
      local _argList = argList
   
      return self:getDefaultVal( self:getRep( typeInfo ) )
   end
   
   if index > #argList:get_expList() then
      return self:getDefaultVal( self:getRep( typeInfo ) )
   end
   return self:lowerExpAs( argList:get_expList()[index], typeInfo )
end

function convFilter:lowerListMethod( node, prefix, name )

   local list = self:lowerExp( prefix )
   do
      local _switchExp = name
      if _switchExp == "insert" then
         return W.call( self.rt:fnListPush(  ), {list, self:lowerArgAny( node, 1 )} ), {}
      elseif _switchExp == "remove" then
         return W.call( self.rt:fnListRemove(  ), {list, W.call( self:fnIntOrZero(  ), {self:lowerArgAs( node, 1, Ast.builtinTypeStem_ )} )} ), {Ast.builtinTypeStem_}
      elseif _switchExp == "sort" then
         local lessInfo = self.rt:getLessCloInfo(  )
         local cmp = W.refNull( lessInfo:get_cloType() )
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() > 0 then
                  local arg = argList:get_expList()[1]
                  if arg:get_kind() ~= Nodes.NodeKind.get_Abbr() then
                     local val = self:lowerVal( arg )
                     local lessType = self.rt:getCloInfo( {W.typeAnyRef, W.typeAnyRef}, {W.typeI32} )
                     if val:get_rep():equalsType( W.refType( lessType:get_cloType() ) ) then
                        cmp = val:get_ins()
                     else
                      
                        cmp = self:adaptClosure( val:get_ins(), val:get_typeInfo(), self:getLessFormType( node ) )
                     end
                  end
               end
            end
         end
         return W.call( self.rt:fnListSort(  ), {list, cmp} ), {}
      elseif _switchExp == "unpack" then
         return W.call( self.rt:fnListClone(  ), {list} ), {Ast.builtinTypeDDD}
      end
   end
   self:error( node, string.format( "not support list method -- %s", name) )
end

function convFilter:getLessFormType( node )

   local funcType = self:normType( node:get_func():get_expType() )
   do
      local fieldNode = _lune.__Cast( node:get_func(), 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         do
            local symbol = fieldNode:get_symbolInfo()
            if symbol ~= nil then
               local declType = symbol:get_typeInfo()
               if #declType:get_argTypeInfoList() > 0 then
                  return declType:get_argTypeInfoList()[1]
               end
            end
         end
      end
   end
   return funcType:get_argTypeInfoList()[1]
end

function convFilter:fnIntOrZero(  )

   do
      local index = self.adapterMap["intOrZero"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {W.typeI64} )
   self.adapterMap["intOrZero"] = index
   func:setBody( {W.ifElse( W.refIsNull( lg( 0 ) ), {W.ret( {W.i64Const( 0 )} )}, nil, nil ), self.rt:unboxI64( lg( 0 ) )} )
   return index
end

function convFilter:lowerSetMethod( node, prefix, name )

   local set = self:lowerExp( prefix )
   do
      local _switchExp = name
      if _switchExp == "add" then
         local work = self:tempLocal( W.typeAnyRef )
         return W.seq( {ls( work, self:lowerArgAny( node, 1 ) ), W.call( self.rt:fnMapSet(  ), {set, lg( work ), lg( work )} )} ), {}
      elseif _switchExp == "del" then
         return W.call( self.rt:fnMapSet(  ), {set, self:lowerArgAny( node, 1 ), W.refNull( W.htAny )} ), {}
      elseif _switchExp == "has" then
         return p( W.opI32Eqz, {W.refIsNull( W.call( self.rt:fnMapGet(  ), {set, self:lowerArgAny( node, 1 )} ) )} ), {Ast.builtinTypeBool}
      elseif _switchExp == "and" or _switchExp == "or" or _switchExp == "sub" then
         local mode = 2
         if name == "and" then
            mode = 0
         elseif name == "or" then
            mode = 1
         end
         return W.call( self.rt:fnSetOp( mode ), {set, self:lowerArgAs( node, 1, prefix:get_expType() )} ), {prefix:get_expType()}
      elseif _switchExp == "clone" then
         return W.call( self.rt:fnMapClone(  ), {set} ), {prefix:get_expType()}
      elseif _switchExp == "len" then
         return p( W.opI64ExtendI32U, {W.structGet( self.rt:get_tMap(), Rt.fMapCount, set )} ), {Ast.builtinTypeInt}
      end
   end
   self:error( node, string.format( "not support set method -- %s", name) )
end

function convFilter:lowerStrMethod( node, prefix, name )

   local strVal = self:lowerExpAs( prefix, Ast.builtinTypeString )
   do
      local _switchExp = name
      if _switchExp == "sub" then
         local endIns = W.i64Const( -1 )
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() >= 2 and argList:get_expList()[2]:get_kind() ~= Nodes.NodeKind.get_Abbr() then
                  endIns = W.call( self:fnIntOrMinus1(  ), {self:lowerArgAs( node, 2, Ast.builtinTypeStem_ )} )
               end
            end
         end
         return W.call( self.rt:fnStrSub(  ), {strVal, self:lowerArgAs( node, 1, Ast.builtinTypeInt ), endIns} ), {Ast.builtinTypeString}
      elseif _switchExp == "find" then
         local initIns = W.i64Const( 1 )
         local plainIns = W.i32Const( 0 )
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() >= 2 and argList:get_expList()[2]:get_kind() ~= Nodes.NodeKind.get_Abbr() then
                  initIns = W.call( self:fnIntOrDefault( 1 ), {self:lowerArgAs( node, 2, Ast.builtinTypeStem_ )} )
               end
               if #argList:get_expList() >= 3 and argList:get_expList()[3]:get_kind() ~= Nodes.NodeKind.get_Abbr() then
                  plainIns = self:toCond( self:lowerVal( argList:get_expList()[3] ) )
               end
            end
         end
         return W.call( self.rt:fnStrFind(  ), {strVal, self:lowerArgAs( node, 1, Ast.builtinTypeString ), initIns, plainIns} ), {Ast.builtinTypeDDD}
      elseif _switchExp == "gsub" then
         return W.call( self.rt:fnStrGsub(  ), {strVal, self:lowerArgAs( node, 1, Ast.builtinTypeString ), self:lowerArgAs( node, 2, Ast.builtinTypeString )} ), {Ast.builtinTypeString, Ast.builtinTypeInt}
      elseif _switchExp == "format" then
         local pre = {}
         local work = self:tempLocal( self.rt:get_strType() )
         table.insert( pre, ls( work, strVal ) )
         local vals = {}
         do
            local argList = node:get_argList()
            if argList ~= nil then
               if #argList:get_expList() == 1 then
                  do
                     local dddNode = _lune.__Cast( argList:get_expList()[1], 3, Nodes.ExpToDDDNode )
                     if dddNode ~= nil then
                        vals = self:lowerValList( dddNode:get_expList(), pre )
                     else
                        vals = self:lowerValList( argList, pre )
                     end
                  end
               end
            end
         end
         table.insert( pre, W.call( self.rt:fnFormat(  ), {lg( work ), self:buildList( vals )} ) )
         return W.seq( pre ), {Ast.builtinTypeString}
      end
   end
   self:error( node, string.format( "not support str method -- %s", name) )
end

function convFilter:fnIntOrDefault( defVal )

   local name = string.format( "intOr:%d", defVal)
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {W.typeI64} )
   self.adapterMap[name] = index
   func:setBody( {W.ifElse( W.refIsNull( lg( 0 ) ), {W.ret( {W.i64Const( defVal )} )}, nil, nil ), self.rt:unboxI64( lg( 0 ) )} )
   return index
end

function convFilter:fnIntOrMinus1(  )

   do
      local index = self.adapterMap["intOrMinus1"]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {W.typeI64} )
   self.adapterMap["intOrMinus1"] = index
   func:setBody( {W.ifElse( W.refIsNull( lg( 0 ) ), {W.ret( {W.i64Const( -1 )} )}, nil, nil ), self.rt:unboxI64( lg( 0 ) )} )
   return index
end

function convFilter:lowerStreamMethod( node, prefix, name )

   local obj = self:tempLocal( self.rt:get_objType() )
   local pre = {ls( obj, self:lowerExpAs( prefix, prefix:get_expType() ) )}
   do
      local _switchExp = name
      if _switchExp == "write" then
         local txt = self:lowerArgAs( node, 1, Ast.builtinTypeString )
         self:getStdStream( true )
         table.insert( pre, W.ifElse( W.refEq( lg( obj ), W.globalGet( self.stderrGlobal ) ), {W.call( self.rt:get_fnHostWriteErr(), {txt} )}, {W.call( self.rt:get_fnHostWrite(), {txt} )}, nil ) )
         table.insert( pre, lg( obj ) )
         table.insert( pre, W.refNull( self.rt:get_tStr() ) )
         return W.seq( pre ), {Ast.builtinTypeStem_, Ast.builtinTypeString:get_nilableTypeInfo()}
      elseif _switchExp == "close" or _switchExp == "flush" then
         return W.seq( pre ), {}
      end
   end
   self:error( node, string.format( "not support stream method -- %s", name) )
end

function convFilter:isStreamType( typeInfo )

   local work = self:getGenSrc( typeInfo )
   local builtin = self.builtinFuncs
   return work == self:getGenSrc( builtin.ostream_ ) or work == self:getGenSrc( builtin.istream_ ) or work == self:getGenSrc( builtin.luastream_ )
end

function convFilter:lowerMethodCall( node, prefix, name, pre )

   local objType = self:getObjType( prefix:get_expType() )
   if objType:get_kind() == Ast.TypeInfoKind.IF then
      local ifInfo = self:getIfInfo( objType )
      local methodIndex = ifInfo:get_name2index()[name]
      if  nil == methodIndex then
         local _methodIndex = methodIndex
      
         self:error( node, string.format( "not found method -- %s", name) )
      end
      
      local methodType = _lune.unwrap( self:getIfMethodType( objType, name ))
      local obj = self:tempLocal( self.rt:get_objType() )
      table.insert( pre, ls( obj, self:lowerExpAs( prefix, objType ) ) )
      local args = {lg( obj )}
      append( args, self:lowerArgs( node:get_argList(), methodType:get_argTypeInfoList(), pre ) )
      local itab = W.refCast( ifInfo:get_itType(), W.arrayGet( self.rt:get_tAnyArr(), W.structGet( self.rt:get_tVT(), Rt.fVtItabs, W.structGet( self.rt:get_tObj(), Rt.fObjVt, lg( obj ) ) ), W.i32Const( ifInfo:get_ifId() ) ) )
      table.insert( args, W.structGet( ifInfo:get_itType(), methodIndex, itab ) )
      table.insert( pre, W.callRef( ifInfo:get_methodFuncTypes()[methodIndex + 1], args ) )
      return W.seq( pre ), methodType:get_retTypeInfoList()
   end
   local info = self:getClassInfo( objType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", objType:getTxt(  )) )
   end
   
   local slotIndex = info:get_name2slot()[name]
   if  nil == slotIndex then
      local _slotIndex = slotIndex
   
      self:error( node, string.format( "not found method -- %s", name) )
   end
   
   local slot = info:get_slotList()[slotIndex]
   local obj = self:tempLocal( W.refType( info:get_structType() ) )
   table.insert( pre, ls( obj, self:lowerExpAs( prefix, info:get_typeInfo() ) ) )
   local args = {lg( obj )}
   append( args, self:lowerArgs( node:get_argList(), slot:get_methodType():get_argTypeInfoList(), pre ) )
   table.insert( args, W.structGet( info:get_vtType(), slotIndex, W.structGet( info:get_structType(), Rt.fObjVt, lg( obj ) ) ) )
   table.insert( pre, W.callRef( slot:get_funcType(), args ) )
   return W.seq( pre ), slot:get_methodType():get_retTypeInfoList()
end

function convFilter:getIfMethodType( ifType, name )

   local map = {}
   self:collectIfMethods( self:getGenSrc( ifType ), map )
   return map[name]
end

function convFilter:lowerCallField( node, fieldNode, pre )

   local prefix = fieldNode:get_prefix()
   local name = fieldNode:get_field().txt
   local prefixType = self:normType( prefix:get_expType() )
   local builtin = self.builtinFuncs
   local funcType = node:get_func():get_expType()
   if funcType:equals( self.processInfo, builtin.string_format ) then
      local vals = {}
      do
         local argList = node:get_argList()
         if argList ~= nil then
            vals = self:lowerValList( argList, pre )
         end
      end
      local rest = {}
      for index = 2, #vals do
         table.insert( rest, vals[index] )
      end
      table.insert( pre, W.call( self.rt:fnFormat(  ), {self:coerceVal( vals[1], Ast.builtinTypeString ), self:buildList( rest )} ) )
      return W.seq( pre ), {Ast.builtinTypeString}
   end
   do
      local _switchExp = prefixType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         return self:lowerListMethod( node, prefix, name )
      elseif _switchExp == Ast.TypeInfoKind.Set then
         return self:lowerSetMethod( node, prefix, name )
      end
   end
   if self:isType( prefixType, Ast.builtinTypeString ) then
      return self:lowerStrMethod( node, prefix, name )
   end
   if self:isStreamType( prefixType ) then
      return self:lowerStreamMethod( node, prefix, name )
   end
   do
      local symbol = fieldNode:get_symbolInfo()
      if symbol ~= nil then
         local methodType = symbol:getOrg(  ):get_typeInfo()
         if methodType:get_staticFlag() or symbol:get_kind() == Ast.SymbolKind.Fun then
            return self:lowerStaticCall( node, prefix, name, methodType, pre )
         end
      end
   end
   if name == "_toMap" then
      do
         local info = self:getClassInfo( self:getObjType( prefixType ) )
         if info ~= nil then
            table.insert( pre, W.call( self:fnToMap( node, info ), {self:lowerExpAs( prefix, info:get_typeInfo() )} ) )
            return W.seq( pre ), {node:get_expType()}
         end
      end
   end
   return self:lowerMethodCall( node, prefix, name, pre )
end

function convFilter:lowerStaticCall( node, prefix, name, methodType, pre )

   do
      local funcIndex = self.func2index[methodType]
      if funcIndex ~= nil then
         local args = {W.refNull( self.rt:get_tClo() )}
         append( args, self:lowerArgs( node:get_argList(), methodType:get_argTypeInfoList(), pre ) )
         table.insert( pre, W.call( funcIndex, args ) )
         return W.seq( pre ), methodType:get_retTypeInfoList()
      end
   end
   local classType = self:getGenSrc( prefix:get_expType() )
   do
      local enumType = _lune.__Cast( classType, 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         if name == "_from" then
            table.insert( pre, W.call( self:fnEnumFrom( enumType ), {self:lowerArgAs( node, 1, enumType:get_valTypeInfo() )} ) )
            return W.seq( pre ), {enumType:get_nilableTypeInfo()}
         end
      end
   end
   if name == "_fromMap" then
      do
         local info = self:getClassInfo( classType )
         if info ~= nil then
            table.insert( pre, W.call( self:fnFromMap( node, info ), {self:lowerArgAs( node, 1, Ast.builtinTypeStem_ )} ) )
            table.insert( pre, W.refNull( self.rt:get_tStr() ) )
            return W.seq( pre ), {info:get_typeInfo():get_nilableTypeInfo(), Ast.builtinTypeString:get_nilableTypeInfo()}
         end
      end
   end
   self:error( node, string.format( "not support static call -- %s", name) )
end

function convFilter:lowerCall( node )

   local pre = {}
   local funcNode = node:get_func()
   do
      local refNode = _lune.__Cast( funcNode, 3, Nodes.ExpRefNode )
      if refNode ~= nil then
         return self:lowerCallRef( node, refNode, pre )
      end
   end
   do
      local fieldNode = _lune.__Cast( funcNode, 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         return self:lowerCallField( node, fieldNode, pre )
      end
   end
   local funcVal = self:lowerVal( funcNode )
   table.insert( pre, self:callClosure( funcVal:get_ins(), funcVal:get_typeInfo(), node:get_argList(), pre ) )
   return W.seq( pre ), self:normType( funcVal:get_typeInfo() ):get_retTypeInfoList()
end

function convFilter:lowerCallSuper( node )

   local methodType = node:get_methodType()
   local funcIndex = self.func2index[methodType]
   if  nil == funcIndex then
      local _funcIndex = funcIndex
   
      self:error( node, "not found super method" )
   end
   
   local pre = {}
   local args = {lg( self.ctx:get_selfLocal() )}
   append( args, self:lowerArgs( node:get_expList(), methodType:get_argTypeInfoList(), pre ) )
   table.insert( pre, W.call( funcIndex, args ) )
   return W.seq( pre ), methodType:get_retTypeInfoList()
end

function convFilter:lowerNew( node )

   local info = self:getClassInfo( node:get_expType() )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", node:get_expType():getTxt(  )) )
   end
   
   local objRep = W.refType( info:get_structType() )
   local obj = self:tempLocal( objRep )
   local fields = {W.i32Const( 0 ), W.globalGet( info:get_vtGlobal() )}
   for __index, rep in ipairs( info:get_fieldReps() ) do
      table.insert( fields, self:getDefaultVal( rep ) )
   end
   local pre = {ls( obj, W.structNew( info:get_structType(), fields ) )}
   local ctorType = self:getCtorType( info:get_typeInfo() )
   local args = {lg( obj )}
   append( args, self:lowerArgs( node:get_argList(), ctorType:get_argTypeInfoList(), pre ) )
   table.insert( pre, W.call( info:get_ctorFunc(), args ) )
   table.insert( pre, lg( obj ) )
   return Val._new(W.seq( pre ), objRep, node:get_expType())
end

function convFilter:getCtorType( classType )

   local scope = _lune.unwrap( classType:get_scope())
   return (_lune.unwrap( scope:get_symbol2SymbolInfoMap()["__init"]) ):get_typeInfo()
end

function convFilter:lowerGetField( node )

   local prefix = node:get_prefix()
   local getterType = node:get_getterTypeInfo()
   local fieldName = node:get_field().txt:gsub( "^%$", "" )
   local getterName = "get_" .. fieldName
   local prefixType = self:normType( prefix:get_expType() )
   do
      local enumType = _lune.__Cast( prefixType, 3, Ast.EnumTypeInfo )
      if enumType ~= nil then
         do
            local _switchExp = getterName
            if _switchExp == "get__txt" then
               return self:newVal( W.call( self:fnEnumTxt( enumType ), {self:lowerExp( prefix )} ), Ast.builtinTypeString )
            elseif _switchExp == "get__allList" then
               return self:newVal( W.call( self:fnEnumAllList( enumType ), {} ), node:get_expType() )
            end
         end
      end
   end
   if prefixType:get_kind() == Ast.TypeInfoKind.Alge and getterName == "get__txt" then
      local algeType = self:getGenSrc( prefixType )
      local info = _lune.unwrap( self:getAlgeInfo( algeType ))
      return self:newVal( W.refCast( self.rt:get_tStr(), W.arrayGet( self.rt:get_tAnyArr(), W.globalGet( info:get_namesGlobal() ), W.structGet( info:get_baseType(), 1, self:lowerExpAs( prefix, algeType ) ) ) ), Ast.builtinTypeString )
   end
   if getterType:get_autoFlag() or not _lune._Set_has(self.declaredMethodSet, getterType ) then
      return self:lowerMemberRef( node, prefix, fieldName )
   end
   local objType = self:getObjType( prefix:get_expType() )
   local info = self:getClassInfo( objType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", objType:getTxt(  )) )
   end
   
   local slotIndex = info:get_name2slot()[getterName]
   if  nil == slotIndex then
      local _slotIndex = slotIndex
   
      self:error( node, string.format( "not found getter -- %s", getterName) )
   end
   
   local slot = info:get_slotList()[slotIndex]
   local obj = self:tempLocal( W.refType( info:get_structType() ) )
   local retType = slot:get_methodType():get_retTypeInfoList()[1]
   return self:newVal( W.seq( {ls( obj, self:lowerExpAs( prefix, info:get_typeInfo() ) ), W.callRef( slot:get_funcType(), {lg( obj ), W.structGet( info:get_vtType(), slotIndex, W.structGet( info:get_structType(), Rt.fObjVt, lg( obj ) ) )} )} ), retType )
end

function convFilter:getEnumValList( enumType )

   local list = {}
   do
      local declNode = self.enumType2decl[enumType]
      if declNode ~= nil then
         for __index, token in ipairs( declNode:get_valueNameList() ) do
            do
               local valInfo = enumType:get_name2EnumValInfo()[token.txt]
               if valInfo ~= nil then
                  table.insert( list, valInfo )
               end
            end
         end
      else
         do
            local __sorted = {}
            local __map = enumType:get_name2EnumValInfo()
            for __key in pairs( __map ) do
               table.insert( __sorted, __key )
            end
            table.sort( __sorted )
            for __index, __key in ipairs( __sorted ) do
               local valInfo = __map[ __key ]
               do
                  table.insert( list, valInfo )
               end
            end
         end
      end
   end
   return list
end


function convFilter:getEnumFuncName( enumType, kind )

   return string.format( "enum%s:%d", kind, enumType:get_typeId().id)
end

function convFilter:enumEq( enumType, val, valInfo )

   local literal = self:lowerEnumVal( enumType, valInfo )
   return self:lowerEqVals( Val._new(val, literal:get_rep(), enumType), literal )
end

function convFilter:fnEnumTxt( enumType )

   local name = self:getEnumFuncName( enumType, "txt" )
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local rep = self:getRep( enumType )
   local index, func = self.mod:addFunc( {rep}, {self.rt:get_strType()} )
   self.adapterMap[name] = index
   local body = {}
   for __index, valInfo in ipairs( self:getEnumValList( enumType ) ) do
      table.insert( body, W.ifElse( self:enumEq( enumType, lg( 0 ), valInfo ), {W.ret( {self.rt:strLiteral( string.format( "%s.%s", enumType:get_rawTxt(), valInfo:get_name()) )} )}, nil, nil ) )
   end
   table.insert( body, self.rt:strLiteral( "" ) )
   func:setBody( body )
   return index
end

function convFilter:fnEnumFrom( enumType )

   local name = self:getEnumFuncName( enumType, "from" )
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local rep = self:getRep( enumType )
   local index, func = self.mod:addFunc( {rep}, {W.typeAnyRef} )
   self.adapterMap[name] = index
   local body = {}
   for __index, valInfo in ipairs( self:getEnumValList( enumType ) ) do
      table.insert( body, W.ifElse( self:enumEq( enumType, lg( 0 ), valInfo ), {W.ret( {self:coerceRep( lg( 0 ), rep, W.typeAnyRef )} )}, nil, nil ) )
   end
   table.insert( body, W.refNull( W.htAny ) )
   func:setBody( body )
   return index
end

function convFilter:fnEnumAllList( enumType )

   local name = self:getEnumFuncName( enumType, "all" )
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {}, {self.rt:get_listType()} )
   self.adapterMap[name] = index
   local vals = {}
   for __index, valInfo in ipairs( self:getEnumValList( enumType ) ) do
      table.insert( vals, self:lowerEnumVal( enumType, valInfo ) )
   end
   local saved = self.ctx
   self.ctx = FuncCtx._new(func, {self.rt:get_listType()}, {})
   local ins = self:buildList( vals )
   self.ctx = saved
   func:setBody( {ins} )
   return index
end

function convFilter:fnToMap( node, info )

   local name = string.format( "toMap:%d", info:get_structType())
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local index, func = self.mod:addFunc( {W.refType( info:get_structType() )}, {self.rt:get_mapType()} )
   self.adapterMap[name] = index
   local map = func:addLocal( self.rt:get_mapType() )
   local body = {ls( map, self.rt:newMap(  ) )}
   for fieldIndex, fieldName in ipairs( info:get_fieldNames() ) do
      local rep = info:get_fieldReps()[fieldIndex]
      table.insert( body, W.call( self.rt:fnMapSet(  ), {lg( map ), self.rt:strLiteral( fieldName ), self:coerceRep( W.structGet( info:get_structType(), fieldIndex + 1, lg( 0 ) ), rep, W.typeAnyRef )} ) )
   end
   table.insert( body, lg( map ) )
   func:setBody( body )
   return index
end

function convFilter:fnFromMap( node, info )

   local name = string.format( "fromMap:%d", info:get_structType())
   do
      local index = self.adapterMap[name]
      if index ~= nil then
         return index
      end
   end
   local objRep = W.refType( info:get_structType() )
   local index, func = self.mod:addFunc( {W.typeAnyRef}, {objRep} )
   self.adapterMap[name] = index
   local map = func:addLocal( self.rt:get_mapType() )
   local fields = {W.i32Const( 0 ), W.globalGet( info:get_vtGlobal() )}
   for fieldIndex, fieldName in ipairs( info:get_fieldNames() ) do
      local rep = info:get_fieldReps()[fieldIndex]
      table.insert( fields, self:coerceRep( W.call( self.rt:fnMapGet(  ), {lg( map ), self.rt:strLiteral( fieldName )} ), W.typeAnyRef, rep ) )
   end
   func:setBody( {W.ifElse( p( W.opI32Eqz, {W.refTest( self.rt:get_tMap(), lg( 0 ) )} ), {W.ret( {W.refNull( info:get_structType() )} )}, nil, nil ), ls( map, W.refCast( self.rt:get_tMap(), lg( 0 ) ) ), W.structNew( info:get_structType(), fields )} )
   return index
end

function convFilter:lowerStmtList( stmtList )

   local list = {}
   for __index, stmt in ipairs( stmtList ) do
      self:preDeclare( stmt, list )
   end
   for __index, stmt in ipairs( stmtList ) do
      self:lowerStmt( stmt, list )
   end
   return list
end


function convFilter:lowerBlock( node )

   return self:lowerStmtList( node:get_stmtList() )
end

function convFilter:lowerBlockAt( node, depth )

   local saved = self.ctx:get_depth()
   self.ctx:set_depth( depth )
   local list = self:lowerBlock( node )
   self.ctx:set_depth( saved )
   return list
end

function convFilter:preDeclare( stmt, out )

   do
      local declNode = _lune.__Cast( stmt, 3, Nodes.DeclFuncNode )
      if declNode ~= nil then
         do
            local symbol = declNode:get_declInfo():get_symbol()
            if symbol ~= nil then
               self:registerFunc( declNode:get_declInfo(), declNode:get_expType() )
               local orgSym = symbol:getOrg(  )
               if not self:isGlobalSym( orgSym ) and not self:findLoc( orgSym ) then
                  table.insert( out, self:declareVar( orgSym, nil ) )
               end
            end
         end
      end
   end
   do
      local classNode = _lune.__Cast( stmt, 3, Nodes.DeclClassNode )
      if classNode ~= nil then
         if classNode:get_expType():get_kind() == Ast.TypeInfoKind.Class then
            self:registerClass( classNode )
         end
      end
   end
end

function convFilter:enterLoop(  )

   self.ctx:set_depth( self.ctx:get_depth() + 1 )
   table.insert( self.ctx:get_breakDepthList(), self.ctx:get_depth() )
   self.ctx:set_depth( self.ctx:get_depth() + 1 )
end

function convFilter:exitLoop( body )

   self.ctx:set_depth( self.ctx:get_depth() - 2 )
   table.remove( self.ctx:get_breakDepthList() )
   return W.block( {W.loop( body, nil )}, nil )
end

function makeIfChain( conds, bodies, elseBody )

   local acc = elseBody
   for index = #conds, 1, -1 do
      acc = {W.ifElse( conds[index], bodies[index], acc, nil )}
   end
   if acc ~= nil then
      return W.seq( acc )
   end
   return W.nop(  )
end


function convFilter:lowerIf( node )

   local base = self.ctx:get_depth()
   local conds = {}
   local bodies = {}
   local elseBody = nil
   for __index, stmt in ipairs( node:get_stmtList() ) do
      if stmt:get_kind() == Nodes.IfKind.Else then
         elseBody = self:lowerBlockAt( stmt:get_block(), base + #conds )
      else
       
         self.ctx:set_depth( base + #conds )
         table.insert( conds, self:lowerCond( stmt:get_exp() ) )
         table.insert( bodies, self:lowerBlockAt( stmt:get_block(), base + #conds ) )
      end
   end
   self.ctx:set_depth( base )
   return makeIfChain( conds, bodies, elseBody )
end


function convFilter:lowerSwitch( node )

   local base = self.ctx:get_depth()
   local val = self:lowerVal( node:get_exp() )
   local work = self:tempLocal( val:get_rep() )
   local localVal = Val._new(lg( work ), val:get_rep(), val:get_typeInfo())
   local conds = {}
   local bodies = {}
   for __index, caseInfo in ipairs( node:get_caseList() ) do
      local cond = nil
      for __index, exp in ipairs( caseInfo:get_expList():get_expList() ) do
         local eq = self:lowerEqVals( localVal, self:lowerVal( exp ) )
         do
            local prev = cond
            if prev ~= nil then
               cond = p( W.opI32Or, {prev, eq} )
            else
               cond = eq
            end
         end
      end
      table.insert( conds, _lune.unwrapDefault( cond, W.i32Const( 0 )) )
      table.insert( bodies, self:lowerBlockAt( caseInfo:get_block(), base + #conds ) )
   end
   local elseBody = nil
   do
      local defBlock = node:get_default()
      if defBlock ~= nil then
         elseBody = self:lowerBlockAt( defBlock, base + #conds )
      end
   end
   return W.seq( {ls( work, val:get_ins() ), makeIfChain( conds, bodies, elseBody )} )
end


function convFilter:lowerMatch( node )

   local base = self.ctx:get_depth()
   local algeType = self:getGenSrc( node:get_val():get_expType() )
   local info = self:getAlgeInfo( algeType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found alge -- %s", algeType:getTxt(  )) )
   end
   
   local algeVal = self:tempLocal( W.refType( info:get_baseType() ) )
   local tag = self:tempLocal( W.typeI32 )
   local conds = {}
   local bodies = {}
   for __index, caseInfo in ipairs( node:get_caseList() ) do
      local name = caseInfo:get_valInfo():get_name()
      table.insert( conds, p( W.opI32Eq, {lg( tag ), W.i32Const( _lune.unwrap( info:get_name2tag()[name]) )} ) )
      self.ctx:set_depth( base + #conds )
      local body = {}
      local valType = _lune.unwrap( info:get_name2type()[name])
      for index, symbol in ipairs( caseInfo:get_valParamNameList() ) do
         local paramType = caseInfo:get_valInfo():get_typeList()[index]
         local paramVal = W.structGet( valType, index + 1, W.refCast( valType, lg( algeVal ) ) )
         table.insert( body, self:declareVar( symbol, self:coerce( paramVal, paramType, symbol:get_typeInfo() ) ) )
      end
      append( body, self:lowerBlock( caseInfo:get_block() ) )
      table.insert( bodies, body )
   end
   local elseBody = nil
   do
      local defBlock = node:get_defaultBlock()
      if defBlock ~= nil then
         self.ctx:set_depth( base + #conds )
         local defBody = {}
         self:lowerStmt( defBlock, defBody )
         elseBody = defBody
      end
   end
   self.ctx:set_depth( base )
   return W.seq( {ls( algeVal, self:lowerExpAs( node:get_val(), algeType ) ), ls( tag, W.structGet( info:get_baseType(), 1, lg( algeVal ) ) ), makeIfChain( conds, bodies, elseBody )} )
end

function convFilter:lowerRawVal( node )

   do
      local castNode = _lune.__Cast( node, 3, Nodes.ExpCastNode )
      if castNode ~= nil then
         if castNode:get_castKind() == Nodes.CastKind.Implicit and castNode:get_exp():get_expType():get_nilable() then
            return self:lowerVal( castNode:get_exp() )
         end
      end
   end
   return self:lowerVal( node )
end

function convFilter:spillVal( val, pre )

   local work = self:tempLocal( val:get_rep() )
   table.insert( pre, ls( work, val:get_ins() ) )
   return Val._new(lg( work ), val:get_rep(), val:get_typeInfo())
end

function convFilter:expandDDD( vals, num, pre )

   if #vals == 0 or #vals > num or self:normType( vals[#vals]:get_typeInfo() ):get_kind() ~= Ast.TypeInfoKind.DDD then
      return vals
   end
   local work = {}
   for index = 1, #vals - 1 do
      table.insert( work, vals[index] )
   end
   local list = self:spillVal( vals[#vals], pre )
   for index = 1, num - #vals + 1 do
      table.insert( work, self:spillVal( Val._new(W.call( self.rt:fnListGet(  ), {list:get_ins(), W.i64Const( index )} ), W.typeAnyRef, Ast.builtinTypeStem_), pre ) )
   end
   return work
end

function convFilter:lowerRawValList( expList, pre )

   local vals = {}
   if expList:get_mRetExp() then
      for __index, val in ipairs( self:lowerValList( expList, pre ) ) do
         table.insert( vals, self:spillVal( val, pre ) )
      end
   else
    
      for __index, exp in ipairs( expList:get_expList() ) do
         table.insert( vals, self:spillVal( self:lowerRawVal( exp ), pre ) )
      end
   end
   return vals
end

function convFilter:anyNil( vals, num )

   local cond = W.i32Const( 0 )
   for index, val in ipairs( vals ) do
      if val:get_rep():isRef(  ) and index <= num then
         cond = p( W.opI32Or, {cond, W.refIsNull( val:get_ins() )} )
      end
   end
   return cond
end


function convFilter:lowerDeclVar( node, out )

   if node:get_unwrapFlag() then
      local base = self.ctx:get_depth()
      local vals = {}
      do
         local expList = node:get_expList()
         if expList ~= nil then
            vals = self:expandDDD( self:lowerRawValList( expList, out ), #node:get_symbolInfoList(), out )
         end
      end
      for __index, symbol in ipairs( node:get_symbolInfoList() ) do
         table.insert( out, self:declareVar( symbol, nil ) )
      end
      self.ctx:set_depth( base + 1 )
      local nilBody = {}
      do
         local block = node:get_unwrapBlock()
         if block ~= nil then
            append( nilBody, self:lowerBlock( block ) )
         end
      end
      local thenBody = {}
      for index, symbol in ipairs( node:get_symbolInfoList() ) do
         if index <= #vals then
            table.insert( thenBody, self:writeSym( node, symbol, self:coerceVal( vals[index], symbol:get_typeInfo() ) ) )
         end
      end
      do
         local block = node:get_thenBlock()
         if block ~= nil then
            append( thenBody, self:lowerBlock( block ) )
         end
      end
      self.ctx:set_depth( base )
      table.insert( out, W.ifElse( self:anyNil( vals, #node:get_symbolInfoList() ), nilBody, thenBody, nil ) )
      return 
   end
   local vals = {}
   do
      local expList = node:get_expList()
      if expList ~= nil then
         vals = self:expandDDD( self:lowerValList( expList, out ), #node:get_symbolInfoList(), out )
      end
   end
   for index, symbol in ipairs( node:get_symbolInfoList() ) do
      if index <= #vals then
         table.insert( out, self:declareVar( symbol, self:coerceVal( vals[index], symbol:get_typeInfo() ) ) )
      else
       
         table.insert( out, self:declareVar( symbol, nil ) )
      end
   end
end


function convFilter:lowerIfUnwrap( node )

   local base = self.ctx:get_depth()
   local pre = {}
   local vals = self:expandDDD( self:lowerRawValList( node:get_expList(), pre ), #node:get_varSymList(), pre )
   self.ctx:set_depth( base + 1 )
   local thenBody = {}
   for index, symbol in ipairs( node:get_varSymList() ) do
      if index <= #vals then
         table.insert( thenBody, self:declareVar( symbol, self:coerceVal( vals[index], symbol:get_typeInfo() ) ) )
      end
   end
   append( thenBody, self:lowerBlock( node:get_block() ) )
   local elseBody = nil
   do
      local block = node:get_nilBlock()
      if block ~= nil then
         elseBody = self:lowerBlock( block )
      end
   end
   self.ctx:set_depth( base )
   local checkNum = #node:get_varSymList()
   if checkNum == 0 then
      checkNum = #vals
   end
   table.insert( pre, W.ifElse( self:anyNil( vals, checkNum ), _lune.unwrapDefault( elseBody, {}), thenBody, nil ) )
   return W.seq( pre )
end


function convFilter:lowerWhen( node )

   local base = self.ctx:get_depth()
   local vals = {}
   for __index, pair in ipairs( node:get_symPairList() ) do
      local src = pair:get_src():getOrg(  )
      table.insert( vals, Val._new(self:readSym( node, src ), self:getRep( src:get_typeInfo() ), src:get_typeInfo()) )
   end
   self.ctx:set_depth( base + 1 )
   local thenBody = {}
   for index, pair in ipairs( node:get_symPairList() ) do
      table.insert( thenBody, self:declareVar( pair:get_dst(), self:coerceVal( vals[index], pair:get_dst():get_typeInfo() ) ) )
   end
   append( thenBody, self:lowerBlock( node:get_block() ) )
   local elseBody = nil
   do
      local block = node:get_elseBlock()
      if block ~= nil then
         elseBody = self:lowerBlock( block )
      end
   end
   self.ctx:set_depth( base )
   return W.ifElse( self:anyNil( vals, #vals ), _lune.unwrapDefault( elseBody, {}), thenBody, nil )
end


function convFilter:lowerUnwrapSet( node, out )

   local base = self.ctx:get_depth()
   local vals = self:expandDDD( self:lowerRawValList( node:get_srcExpList(), out ), #node:get_dstExpList():get_expList(), out )
   self.ctx:set_depth( base + 1 )
   local nilBody = {}
   do
      local block = node:get_unwrapBlock()
      if block ~= nil then
         append( nilBody, self:lowerBlock( block ) )
      end
   end
   local thenBody = {}
   for index, dst in ipairs( node:get_dstExpList():get_expList() ) do
      if index <= #vals then
         self:lowerAssign( dst, vals[index], thenBody )
      end
   end
   self.ctx:set_depth( base )
   table.insert( out, W.ifElse( self:anyNil( vals, #node:get_dstExpList():get_expList() ), nilBody, thenBody, nil ) )
end

function convFilter:lowerAssign( dst, val, out )

   do
      local refNode = _lune.__Cast( dst, 3, Nodes.ExpRefNode )
      if refNode ~= nil then
         local symbol = refNode:get_symbolInfo():getOrg(  )
         table.insert( out, self:writeSym( dst, symbol, self:coerceVal( val, symbol:get_typeInfo() ) ) )
         return 
      end
   end
   do
      local fieldNode = _lune.__Cast( dst, 3, Nodes.RefFieldNode )
      if fieldNode ~= nil then
         local prefixType = self:normType( fieldNode:get_prefix():get_expType() )
         local name = fieldNode:get_field().txt
         do
            local symbol = fieldNode:get_symbolInfo()
            if symbol ~= nil then
               local orgSym = symbol:getOrg(  )
               if self.globalSym2loc[orgSym] then
                  table.insert( out, self:writeSym( dst, orgSym, self:coerceVal( val, orgSym:get_typeInfo() ) ) )
                  return 
               end
            end
         end
         if prefixType:get_kind() == Ast.TypeInfoKind.Map then
            table.insert( out, W.call( self.rt:fnMapSet(  ), {self:lowerExp( fieldNode:get_prefix() ), self.rt:strLiteral( name ), self:toAny( val )} ) )
            return 
         end
         local objType = self:getObjType( prefixType )
         local info = self:getClassInfo( objType )
         if  nil == info then
            local _info = info
         
            self:error( dst, string.format( "not found class -- %s", objType:getTxt(  )) )
         end
         
         local fieldIndex = self:getFieldIndex( dst, info, name )
         table.insert( out, W.structSet( info:get_structType(), fieldIndex, W.refCast( info:get_structType(), self:lowerExpAs( fieldNode:get_prefix(), info:get_typeInfo() ) ), self:coerce( val:get_ins(), val:get_typeInfo(), self:getMemberType( info, name ) ) ) )
         return 
      end
   end
   do
      local itemNode = _lune.__Cast( dst, 3, Nodes.ExpRefItemNode )
      if itemNode ~= nil then
         local index
         
         do
            local indexNode = itemNode:get_index()
            if indexNode ~= nil then
               index = self:lowerVal( indexNode )
            else
               index = Val._new(self.rt:strLiteral( _lune.unwrap( itemNode:get_symbol()) ), self.rt:get_strType(), Ast.builtinTypeString)
            end
         end
         self:assignItem( dst, itemNode:get_val(), index, val, out )
         return 
      end
   end
   self:error( dst, string.format( "not support assign -- %s", Nodes.getNodeKindName( dst:get_kind() )) )
end

function convFilter:assignItem( node, collNode, index, val, out )

   local collType = self:normType( collNode:get_expType() )
   local coll = self:lowerExp( collNode )
   do
      local _switchExp = collType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         table.insert( out, W.call( self.rt:fnListSet(  ), {coll, self:coerceVal( index, Ast.builtinTypeInt ), self:toAny( val )} ) )
         return 
      elseif _switchExp == Ast.TypeInfoKind.Map then
         table.insert( out, W.call( self.rt:fnMapSet(  ), {coll, self:toAny( index ), self:toAny( val )} ) )
         return 
      end
   end
   self:error( node, string.format( "not support assign -- %s", collType:getTxt(  )) )
end

function convFilter:lowerSetItem( node, out )

   local index
   
   do
      local _matchExp = node:get_index()
      if _matchExp[1] == Nodes.IndexVal.NodeIdx[1] then
         local indexNode = _matchExp[2][1]
      
         index = self:lowerVal( indexNode )
      elseif _matchExp[1] == Nodes.IndexVal.SymIdx[1] then
         local symTxt = _matchExp[2][1]
      
         index = Val._new(self.rt:strLiteral( symTxt ), self.rt:get_strType(), Ast.builtinTypeString)
      end
   end
   local val
   
   do
      local listNode = _lune.__Cast( node:get_exp2(), 3, Nodes.ExpListNode )
      if listNode ~= nil then
         local vals = self:lowerValList( listNode, out )
         val = vals[1]
      else
         val = self:lowerVal( node:get_exp2() )
      end
   end
   self:assignItem( node, node:get_val(), index, val, out )
end


function convFilter:lowerSetVal( node, out )

   local dstList = {}
   do
      local listNode = _lune.__Cast( node:get_exp1(), 3, Nodes.ExpListNode )
      if listNode ~= nil then
         for __index, exp in ipairs( listNode:get_expList() ) do
            table.insert( dstList, exp )
         end
      else
         table.insert( dstList, node:get_exp1() )
      end
   end
   local vals = self:expandDDD( self:lowerValList( node:get_exp2(), out ), #dstList, out )
   if #dstList > 1 then
      local work = {}
      for __index, val in ipairs( vals ) do
         table.insert( work, self:spillVal( val, out ) )
      end
      vals = work
   end
   for index, dst in ipairs( dstList ) do
      if index <= #vals then
         self:lowerAssign( dst, vals[index], out )
      else
       
         self:lowerAssign( dst, Val._new(W.refNull( W.htAny ), W.typeAnyRef, Ast.builtinTypeNil), out )
      end
   end
end


function convFilter:lowerReturn( node, out )

   local vals = {}
   do
      local expList = node:get_expList()
      if expList ~= nil then
         vals = self:lowerValList( expList, out )
      end
   end
   local retTypes = self.ctx:get_retTypes()
   local rets = {}
   for index, retType in ipairs( retTypes ) do
      local retRep = self.ctx:get_retReps()[index]
      if self:normType( retType ):get_kind() == Ast.TypeInfoKind.DDD and index == #retTypes then
         if #vals == index and self:normType( vals[index]:get_typeInfo() ):get_kind() == Ast.TypeInfoKind.DDD then
            table.insert( rets, vals[index]:get_ins() )
         else
          
            local rest = {}
            for work = index, #vals do
               table.insert( rest, vals[work] )
            end
            table.insert( rets, self:buildList( rest ) )
         end
      elseif index <= #vals then
         table.insert( rets, self:coerceRep( self:coerceVal( vals[index], retType ), self:getRep( retType ), retRep ) )
      else
       
         table.insert( rets, self:getDefaultVal( retRep ) )
      end
   end
   table.insert( out, W.ret( rets ) )
end


function convFilter:lowerFor( node )

   local counter = self:tempLocal( W.typeI64 )
   local toIdx = self:tempLocal( W.typeI64 )
   local stepIdx = self:tempLocal( W.typeI64 )
   local list = {}
   table.insert( list, ls( counter, self:lowerExpAs( node:get_init(), Ast.builtinTypeInt ) ) )
   table.insert( list, ls( toIdx, self:lowerExpAs( node:get_to(), Ast.builtinTypeInt ) ) )
   do
      local delta = node:get_delta()
      if delta ~= nil then
         table.insert( list, ls( stepIdx, self:lowerExpAs( delta, Ast.builtinTypeInt ) ) )
      else
         table.insert( list, ls( stepIdx, W.i64Const( 1 ) ) )
      end
   end
   self:enterLoop(  )
   local body = {}
   table.insert( body, W.brIf( 1, W.ifElse( p( W.opI64GeS, {lg( stepIdx ), W.i64Const( 0 )} ), {p( W.opI64GtS, {lg( counter ), lg( toIdx )} )}, {p( W.opI64LtS, {lg( counter ), lg( toIdx )} )}, W.typeI32 ) ) )
   table.insert( body, self:declareVar( node:get_val(), self:coerceRep( lg( counter ), W.typeI64, self:getRep( node:get_val():get_typeInfo() ) ) ) )
   table.insert( body, ls( counter, p( W.opI64Add, {lg( counter ), lg( stepIdx )} ) ) )
   append( body, self:lowerBlock( node:get_block() ) )
   table.insert( body, W.br( 0 ) )
   table.insert( list, self:exitLoop( body ) )
   return W.seq( list )
end


function convFilter:lowerForeach( node, valSym, keySym, exp, block, sort )

   local expType = self:normType( exp:get_expType() )
   local list = {}
   local pos = self:tempLocal( W.typeI32 )
   table.insert( list, ls( pos, W.i32Const( 0 ) ) )
   do
      local _switchExp = expType:get_kind()
      if _switchExp == Ast.TypeInfoKind.List or _switchExp == Ast.TypeInfoKind.Array then
         local coll = self:tempLocal( self.rt:get_listType() )
         table.insert( list, ls( coll, self:lowerExp( exp ) ) )
         self:enterLoop(  )
         local body = {}
         table.insert( body, W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.structGet( self.rt:get_tList(), Rt.fListLen, lg( coll ) )} ) ) )
         local item = W.arrayGet( self.rt:get_tAnyArr(), W.structGet( self.rt:get_tList(), Rt.fListItems, lg( coll ) ), lg( pos ) )
         table.insert( body, self:declareVar( valSym, self:coerceRep( item, W.typeAnyRef, self:getRep( valSym:get_typeInfo() ) ) ) )
         table.insert( body, ls( pos, p( W.opI32Add, {lg( pos ), W.i32Const( 1 )} ) ) )
         if keySym ~= nil then
            table.insert( body, self:declareVar( keySym, p( W.opI64ExtendI32U, {lg( pos )} ) ) )
         end
         append( body, self:lowerBlock( block ) )
         table.insert( body, W.br( 0 ) )
         table.insert( list, self:exitLoop( body ) )
         return W.seq( list )
      elseif _switchExp == Ast.TypeInfoKind.Map or _switchExp == Ast.TypeInfoKind.Set then
         local isSet = expType:get_kind() == Ast.TypeInfoKind.Set
         local coll = self:tempLocal( self.rt:get_mapType() )
         table.insert( list, ls( coll, self:lowerExp( exp ) ) )
         if sort then
            local keys = self:tempLocal( self.rt:get_listType() )
            table.insert( list, ls( keys, W.call( self.rt:fnMapKeys(  ), {lg( coll )} ) ) )
            table.insert( list, W.call( self.rt:fnListSort(  ), {lg( keys ), W.refNull( self.rt:getLessCloInfo(  ):get_cloType() )} ) )
            local key = self:tempLocal( W.typeAnyRef )
            self:enterLoop(  )
            local body = {}
            table.insert( body, W.brIf( 1, p( W.opI32GeS, {lg( pos ), W.structGet( self.rt:get_tList(), Rt.fListLen, lg( keys ) )} ) ) )
            table.insert( body, ls( key, W.arrayGet( self.rt:get_tAnyArr(), W.structGet( self.rt:get_tList(), Rt.fListItems, lg( keys ) ), lg( pos ) ) ) )
            table.insert( body, ls( pos, p( W.opI32Add, {lg( pos ), W.i32Const( 1 )} ) ) )
            local valIns
            
            if isSet then
               valIns = lg( key )
            else
             
               valIns = W.call( self.rt:fnMapGet(  ), {lg( coll ), lg( key )} )
            end
            table.insert( body, self:declareVar( valSym, self:coerceRep( valIns, W.typeAnyRef, self:getRep( valSym:get_typeInfo() ) ) ) )
            if keySym ~= nil then
               table.insert( body, self:declareVar( keySym, self:coerceRep( lg( key ), W.typeAnyRef, self:getRep( keySym:get_typeInfo() ) ) ) )
            end
            append( body, self:lowerBlock( block ) )
            table.insert( body, W.br( 0 ) )
            table.insert( list, self:exitLoop( body ) )
            return W.seq( list )
         end
         self:enterLoop(  )
         local body = {}
         table.insert( body, ls( pos, W.call( self.rt:fnMapNext(  ), {lg( coll ), lg( pos )} ) ) )
         table.insert( body, W.brIf( 1, p( W.opI32LtS, {lg( pos ), W.i32Const( 0 )} ) ) )
         local keyIns = self.rt:mapKeyAt( lg( coll ), lg( pos ) )
         local valIns
         
         if isSet then
            valIns = keyIns
         else
          
            valIns = self.rt:mapValAt( lg( coll ), lg( pos ) )
         end
         table.insert( body, self:declareVar( valSym, self:coerceRep( valIns, W.typeAnyRef, self:getRep( valSym:get_typeInfo() ) ) ) )
         if keySym ~= nil then
            table.insert( body, self:declareVar( keySym, self:coerceRep( keyIns, W.typeAnyRef, self:getRep( keySym:get_typeInfo() ) ) ) )
         end
         table.insert( body, ls( pos, p( W.opI32Add, {lg( pos ), W.i32Const( 1 )} ) ) )
         append( body, self:lowerBlock( block ) )
         table.insert( body, W.br( 0 ) )
         table.insert( list, self:exitLoop( body ) )
         return W.seq( list )
      end
   end
   self:error( node, string.format( "not support foreach -- %s", expType:getTxt(  )) )
end

function convFilter:lowerStmtExp( exp, out )

   if exp:get_kind() == Nodes.NodeKind.get_ExpSetVal() or exp:get_kind() == Nodes.NodeKind.get_ExpSetItem() then
      self:lowerStmt( exp, out )
      return 
   end
   do
      local callNode = _lune.__Cast( exp, 3, Nodes.ExpCallNode )
      if callNode ~= nil then
         local ins, retTypes = self:lowerCall( callNode )
         table.insert( out, ins )
         for __index, _1 in ipairs( retTypes ) do
            table.insert( out, W.drop( W.nop(  ) ) )
         end
         return 
      end
   end
   do
      local superNode = _lune.__Cast( exp, 3, Nodes.ExpCallSuperNode )
      if superNode ~= nil then
         local ins, retTypes = self:lowerCallSuper( superNode )
         table.insert( out, ins )
         for __index, _2 in ipairs( retTypes ) do
            table.insert( out, W.drop( W.nop(  ) ) )
         end
         return 
      end
   end
   do
      local ctorNode = _lune.__Cast( exp, 3, Nodes.ExpCallSuperCtorNode )
      if ctorNode ~= nil then
         self:lowerCallSuperCtor( ctorNode, out )
         return 
      end
   end
   do
      local macroNode = _lune.__Cast( exp, 3, Nodes.ExpMacroExpNode )
      if macroNode ~= nil then
         for __index, ins in ipairs( self:lowerStmtList( macroNode:get_stmtList() ) ) do
            table.insert( out, ins )
         end
         return 
      end
   end
   local val = self:lowerVal( exp )
   if val:get_typeInfo() == Ast.builtinTypeNone then
      table.insert( out, val:get_ins() )
   else
    
      table.insert( out, W.drop( val:get_ins() ) )
   end
end

function convFilter:lowerCallSuperCtor( node, out )

   local info = self:getClassInfo( node:get_superType() )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", node:get_superType():getTxt(  )) )
   end
   
   local args = {lg( self.ctx:get_selfLocal() )}
   append( args, self:lowerArgs( node:get_expList(), self:getCtorType( info:get_typeInfo() ):get_argTypeInfoList(), out ) )
   table.insert( out, W.call( info:get_ctorFunc(), args ) )
end

function convFilter:lowerStmt( node, out )

   do
      local _switchExp = node:get_kind()
      if _switchExp == Nodes.NodeKind.get_None() or _switchExp == Nodes.NodeKind.get_BlankLine() or _switchExp == Nodes.NodeKind.get_Shebang() or _switchExp == Nodes.NodeKind.get_LuneControl() or _switchExp == Nodes.NodeKind.get_DeclMacro() or _switchExp == Nodes.NodeKind.get_TestBlock() or _switchExp == Nodes.NodeKind.get_DeclForm() or _switchExp == Nodes.NodeKind.get_ProtoClass() or _switchExp == Nodes.NodeKind.get_DeclEnum() or _switchExp == Nodes.NodeKind.get_DeclAlge() or _switchExp == Nodes.NodeKind.get_ProtoMethod() or _switchExp == Nodes.NodeKind.get_Subfile() then
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclFunc() then
         local declNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.DeclFuncNode ))
         local funcIndex = self:registerFunc( declNode:get_declInfo(), declNode:get_expType() )
         local closure = self:compileFunc( node, declNode:get_declInfo(), declNode:get_expType(), funcIndex )
         do
            local symbol = declNode:get_declInfo():get_symbol()
            if symbol ~= nil then
               if not self:isGlobalSym( symbol:getOrg(  ) ) then
                  table.insert( out, self:writeSym( node, symbol, closure ) )
               end
            end
         end
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclClass() then
         self:lowerDeclClass( _lune.unwrap( _lune.__Cast( node, 3, Nodes.DeclClassNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclMethod() then
         local declNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.DeclMethodNode ))
         self:compileMethodNode( node, declNode:get_declInfo() )
         return 
      elseif _switchExp == Nodes.NodeKind.get_StmtExp() then
         self:lowerStmtExp( (_lune.unwrap( _lune.__Cast( node, 3, Nodes.StmtExpNode )) ):get_exp(), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_ExpCall() or _switchExp == Nodes.NodeKind.get_ExpCallSuperCtor() or _switchExp == Nodes.NodeKind.get_ExpMacroExp() then
         self:lowerStmtExp( node, out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_DeclVar() then
         self:lowerDeclVar( _lune.unwrap( _lune.__Cast( node, 3, Nodes.DeclVarNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_ExpSetVal() then
         self:lowerSetVal( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpSetValNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_ExpSetItem() then
         self:lowerSetItem( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ExpSetItemNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_UnwrapSet() then
         self:lowerUnwrapSet( _lune.unwrap( _lune.__Cast( node, 3, Nodes.UnwrapSetNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_If() then
         table.insert( out, self:lowerIf( _lune.unwrap( _lune.__Cast( node, 3, Nodes.IfNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_IfUnwrap() then
         table.insert( out, self:lowerIfUnwrap( _lune.unwrap( _lune.__Cast( node, 3, Nodes.IfUnwrapNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_When() then
         table.insert( out, self:lowerWhen( _lune.unwrap( _lune.__Cast( node, 3, Nodes.WhenNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Switch() then
         table.insert( out, self:lowerSwitch( _lune.unwrap( _lune.__Cast( node, 3, Nodes.SwitchNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Match() then
         table.insert( out, self:lowerMatch( _lune.unwrap( _lune.__Cast( node, 3, Nodes.MatchNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_While() then
         local whileNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.WhileNode ))
         self:enterLoop(  )
         local body = {}
         if not whileNode:get_infinit() then
            table.insert( body, W.brIf( 1, p( W.opI32Eqz, {self:lowerCond( whileNode:get_exp() )} ) ) )
         end
         append( body, self:lowerBlock( whileNode:get_block() ) )
         table.insert( body, W.br( 0 ) )
         table.insert( out, self:exitLoop( body ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Repeat() then
         local repeatNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.RepeatNode ))
         self:enterLoop(  )
         local body = {}
         append( body, self:lowerBlock( repeatNode:get_block() ) )
         table.insert( body, W.brIf( 0, p( W.opI32Eqz, {self:lowerCond( repeatNode:get_exp() )} ) ) )
         table.insert( out, self:exitLoop( body ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_For() then
         table.insert( out, self:lowerFor( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ForNode )) ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Foreach() then
         local foreachNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ForeachNode ))
         table.insert( out, self:lowerForeach( node, foreachNode:get_val(), foreachNode:get_key(), foreachNode:get_exp(), foreachNode:get_block(), false ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Forsort() then
         local forsortNode = _lune.unwrap( _lune.__Cast( node, 3, Nodes.ForsortNode ))
         table.insert( out, self:lowerForeach( node, forsortNode:get_val(), forsortNode:get_key(), forsortNode:get_exp(), forsortNode:get_block(), true ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Break() then
         local breakList = self.ctx:get_breakDepthList()
         if #breakList == 0 then
            self:error( node, "break outside loop" )
         end
         table.insert( out, W.br( self.ctx:get_depth() - breakList[#breakList] ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Return() then
         self:lowerReturn( _lune.unwrap( _lune.__Cast( node, 3, Nodes.ReturnNode )), out )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Scope() then
         append( out, self:lowerBlock( (_lune.unwrap( _lune.__Cast( node, 3, Nodes.ScopeNode )) ):get_block() ) )
         return 
      elseif _switchExp == Nodes.NodeKind.get_Block() then
         append( out, self:lowerBlock( _lune.unwrap( _lune.__Cast( node, 3, Nodes.BlockNode )) ) )
         return 
      end
   end
   self:notSupport( node )
end

function convFilter:registerFunc( declInfo, funcType )

   local _
   do
      local index = self.func2index[funcType]
      if index ~= nil then
         return index
      end
   end
   local info = self:getFuncCloInfo( funcType )
   local index, _1 = self.mod:addFuncWithType( info:get_funcType(), 1 + #funcType:get_argTypeInfoList() )
   self.func2index[funcType] = index
   self.mod:setFuncName( index, funcType:getTxt(  ) )
   return index
end

function convFilter:getCaptureSyms( funcType )

   local list = {}
   do
      local scope = funcType:get_scope()
      if scope ~= nil then
         for __index, symbol in ipairs( scope:get_closureSymList() ) do
            local orgSym = symbol:getOrg(  )
            if not self:isGlobalSym( orgSym ) then
               table.insert( list, orgSym )
            end
         end
      end
   end
   return list
end

function convFilter:getCellRef( node, symbol )

   local loc = self:findLoc( symbol )
   if  nil == loc then
      local _loc = loc
   
      self:error( node, string.format( "not found closure symbol -- %s", symbol:get_name()) )
   end
   
   do
      local _matchExp = loc
      if _matchExp[1] == VarLoc.CellLocal[1] then
         local index = _matchExp[2][1]
      
         return lg( index )
      elseif _matchExp[1] == VarLoc.Captured[1] then
         local field = _matchExp[2][1]
      
         return self:envGet( field )
      elseif _matchExp[1] == VarLoc.Local[1] then
         local index = _matchExp[2][1]
      
         local rep = self:getRep( symbol:get_typeInfo() )
         return W.structNew( self.rt:getCellType( rep ), {lg( index )} )
      elseif _matchExp[1] == VarLoc.Global[1] then
         local _ = _matchExp[2][1]
      
         self:error( node, string.format( "global closure symbol -- %s", symbol:get_name()) )
      end
   end
end

function convFilter:setupParams( declInfo, funcType, body )

   for index, arg in ipairs( declInfo:get_argList() ) do
      if arg:get_kind() == Nodes.NodeKind.get_DeclArgDDD() then
         do
            local scope = funcType:get_scope()
            if scope ~= nil then
               do
                  local symbol = scope:get_symbol2SymbolInfoMap()["..."]
                  if symbol ~= nil then
                     self.ctx:get_sym2loc()[symbol:getOrg(  )] = _lune.newAlge( VarLoc.Local, {index})
                  end
               end
            end
         end
      end
      do
         local argNode = _lune.__Cast( arg, 3, Nodes.DeclArgNode )
         if argNode ~= nil then
            local symbol = argNode:get_symbolInfo():getOrg(  )
            if symbol:get_hasAccessFromClosure() then
               table.insert( body, self:declareVar( symbol, lg( index ) ) )
            else
             
               self.ctx:get_sym2loc()[symbol] = _lune.newAlge( VarLoc.Local, {index})
            end
         end
      end
   end
end

function convFilter:compileFunc( node, declInfo, funcType, funcIndex )

   local info = self:getFuncCloInfo( funcType )
   local capSyms = self:getCaptureSyms( funcType )
   local envType = -1
   local cloIns
   
   if #capSyms > 0 then
      local fields = {W.field( W.typeI32, true ), W.field( W.refType( info:get_funcType() ), false )}
      local vals = {W.i32Const( 0 ), self.mod:refFunc( funcIndex )}
      for __index, symbol in ipairs( capSyms ) do
         local cellType = self.rt:getCellType( self:getRep( symbol:get_typeInfo() ) )
         table.insert( fields, W.field( W.refType( cellType ), false ) )
         table.insert( vals, self:getCellRef( node, symbol ) )
      end
      envType = self.mod:addStructType( info:get_cloType(), fields )
      cloIns = W.structNew( envType, vals )
   else
    
      cloIns = W.structNew( info:get_cloType(), {W.i32Const( 0 ), self.mod:refFunc( funcIndex )} )
   end
   if _lune._Set_has(self.compiledFuncSet, funcIndex ) then
      return cloIns
   end
   self.compiledFuncSet[funcIndex]= true
   local saved = self.ctx
   local retTypes = self:normType( funcType ):get_retTypeInfoList()
   self.ctx = FuncCtx._new(self.mod:getFunc( funcIndex ), info:get_results(), retTypes)
   local body = {}
   if envType >= 0 then
      local envLocal = self.ctx:addLocal( W.refType( envType ) )
      self.ctx:set_envLocal( envLocal )
      self.ctx:set_envType( envType )
      table.insert( body, ls( envLocal, W.refCast( envType, lg( 0 ) ) ) )
      for index, symbol in ipairs( capSyms ) do
         self.ctx:get_sym2loc()[symbol] = _lune.newAlge( VarLoc.Captured, {index + 1})
      end
   end
   self:setupParams( declInfo, funcType, body )
   do
      local block = declInfo:get_body()
      if block ~= nil then
         append( body, self:lowerBlock( block ) )
      end
   end
   if #retTypes > 0 then
      table.insert( body, W.unreachable(  ) )
   end
   self.ctx:get_func():setBody( body )
   self.ctx = saved
   return cloIns
end

function convFilter:compileMethod( node, declInfo, methodType, funcIndex, info, retReps )

   if _lune._Set_has(self.compiledFuncSet, funcIndex ) then
      return 
   end
   self.compiledFuncSet[funcIndex]= true
   local saved = self.ctx
   self.ctx = FuncCtx._new(self.mod:getFunc( funcIndex ), retReps, methodType:get_retTypeInfoList())
   local body = {}
   local selfLocal = self.ctx:addLocal( W.refType( info:get_structType() ) )
   self.ctx:set_selfLocal( selfLocal )
   self.ctx:set_selfClass( info )
   table.insert( body, ls( selfLocal, W.refCast( info:get_structType(), lg( 0 ) ) ) )
   do
      local scope = methodType:get_scope()
      if scope ~= nil then
         do
            local selfSym = scope:get_symbol2SymbolInfoMap()["self"]
            if selfSym ~= nil then
               local orgSym = selfSym:getOrg(  )
               if orgSym:get_hasAccessFromClosure() then
                  table.insert( body, self:declareVar( orgSym, lg( selfLocal ) ) )
               else
                
                  self.ctx:get_sym2loc()[orgSym] = _lune.newAlge( VarLoc.Local, {selfLocal})
               end
            end
         end
      end
   end
   self:setupParams( declInfo, methodType, body )
   do
      local block = declInfo:get_body()
      if block ~= nil then
         append( body, self:lowerBlock( block ) )
      end
   end
   if #retReps > 0 then
      table.insert( body, W.unreachable(  ) )
   end
   self.ctx:get_func():setBody( body )
   self.ctx = saved
end

function convFilter:getSlotRetReps( info, methodType )

   do
      local slotIndex = info:get_name2slot()[methodType:get_rawTxt()]
      if slotIndex ~= nil then
         return self:getRepList( info:get_slotList()[slotIndex]:get_methodType():get_retTypeInfoList() )
      end
   end
   return self:getRepList( methodType:get_retTypeInfoList() )
end

function convFilter:compileMethodNode( node, declInfo )

   local classType = declInfo:get_classTypeInfo()
   if  nil == classType then
      local _classType = classType
   
      self:error( node, "no class" )
   end
   
   local info = self:getClassInfo( classType )
   if  nil == info then
      local _info = info
   
      self:error( node, string.format( "not found class -- %s", classType:getTxt(  )) )
   end
   
   local methodType = node:get_expType()
   local funcIndex = self.func2index[methodType]
   if  nil == funcIndex then
      local _funcIndex = funcIndex
   
      if declInfo:get_kind() == Nodes.FuncKind.Ctor then
         self:compileMethod( node, declInfo, methodType, info:get_ctorFunc(), info, {} )
         return 
      end
      self:error( node, string.format( "not found method -- %s", methodType:getTxt(  )) )
   end
   
   if declInfo:get_staticFlag() then
      self:compileFunc( node, declInfo, methodType, funcIndex )
      return 
   end
   if declInfo:get_kind() == Nodes.FuncKind.Ctor then
      self:compileMethod( node, declInfo, methodType, funcIndex, info, {} )
      return 
   end
   self:compileMethod( node, declInfo, methodType, funcIndex, info, self:getSlotRetReps( info, methodType ) )
end

function convFilter:compileDefaultCtor( node, info )

   if _lune._Set_has(self.compiledFuncSet, info:get_ctorFunc() ) then
      return 
   end
   self.compiledFuncSet[info:get_ctorFunc()]= true
   local ctorType = self:getCtorType( info:get_typeInfo() )
   local func = self.mod:getFunc( info:get_ctorFunc() )
   local selfLocal = func:addLocal( W.refType( info:get_structType() ) )
   local body = {ls( selfLocal, W.refCast( info:get_structType(), lg( 0 ) ) )}
   local superArgNum = 0
   do
      local parent = info:get_parent()
      if parent ~= nil then
         local parentCtorType = self:getCtorType( parent:get_typeInfo() )
         local args = {lg( selfLocal )}
         for index, argType in ipairs( parentCtorType:get_argTypeInfoList() ) do
            if node:get_hasOldCtor() then
               table.insert( args, self:getDefaultVal( self:getRep( argType ) ) )
            else
             
               table.insert( args, lg( index ) )
            end
         end
         table.insert( body, W.call( parent:get_ctorFunc(), args ) )
         if not node:get_hasOldCtor() then
            superArgNum = #parentCtorType:get_argTypeInfoList()
         end
      end
   end
   local argIndex = superArgNum + 1
   for __index, memberNode in ipairs( node:get_memberList() ) do
      if not memberNode:get_staticFlag() then
         if argIndex <= #ctorType:get_argTypeInfoList() then
            local fieldIndex = self:getFieldIndex( node, info, memberNode:get_name().txt )
            table.insert( body, W.structSet( info:get_structType(), fieldIndex, lg( selfLocal ), self:coerce( lg( argIndex ), ctorType:get_argTypeInfoList()[argIndex], memberNode:get_symbolInfo():get_typeInfo() ) ) )
         end
         argIndex = argIndex + 1
      end
   end
   func:setBody( body )
end

function convFilter:compileAccessors( node, info )

   local scope = _lune.unwrap( info:get_typeInfo():get_scope())
   for __index, memberNode in ipairs( node:get_memberList() ) do
      local member = memberNode:get_name().txt
      if not memberNode:get_staticFlag() then
         for __index, name in ipairs( {"get_" .. member, "set_" .. member} ) do
            do
               local symbol = scope:get_symbol2SymbolInfoMap()[name]
               if symbol ~= nil then
                  local methodType = symbol:get_typeInfo()
                  if not _lune._Set_has(self.declaredMethodSet, methodType ) then
                     do
                        local funcIndex = self.func2index[methodType]
                        if funcIndex ~= nil then
                           if not _lune._Set_has(self.compiledFuncSet, funcIndex ) then
                              self.compiledFuncSet[funcIndex]= true
                              local func = self.mod:getFunc( funcIndex )
                              local selfIns = W.refCast( info:get_structType(), lg( 0 ) )
                              local fieldIndex = self:getFieldIndex( node, info, member )
                              if name:find( "^get_" ) then
                                 local retReps = self:getSlotRetReps( info, methodType )
                                 func:setBody( {self:coerceRep( W.structGet( info:get_structType(), fieldIndex, selfIns ), info:get_fieldReps()[fieldIndex - 1], retReps[1] )} )
                              else
                               
                                 func:setBody( {W.structSet( info:get_structType(), fieldIndex, selfIns, self:coerceRep( lg( 1 ), self:getRep( methodType:get_argTypeInfoList()[1] ), info:get_fieldReps()[fieldIndex - 1] ) )} )
                              end
                           end
                        end
                     end
                  end
               end
            end
         end
      end
   end
end


function convFilter:lowerDeclClass( node, out )

   if node:get_expType():get_kind() == Ast.TypeInfoKind.IF then
      self:getIfInfo( node:get_expType() )
      return 
   end
   local info = self:registerClass( node )
   for __index, field in ipairs( node:get_fieldList() ) do
      do
         local _switchExp = field:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclMethod() then
            self:compileMethodNode( field, (_lune.unwrap( _lune.__Cast( field, 3, Nodes.DeclMethodNode )) ):get_declInfo() )
         elseif _switchExp == Nodes.NodeKind.get_DeclConstr() then
            self:compileMethodNode( field, (_lune.unwrap( _lune.__Cast( field, 3, Nodes.DeclConstrNode )) ):get_declInfo() )
         end
      end
   end
   if not node:hasUserInit(  ) then
      self:compileDefaultCtor( node, info )
   end
   self:compileAccessors( node, info )
   do
      local initFunc = node:get_initBlock():get_func()
      if initFunc ~= nil then
         do
            local block = initFunc:get_declInfo():get_body()
            if block ~= nil then
               append( out, self:lowerBlock( block ) )
            end
         end
      end
   end
end

function convFilter:finalizeClasses(  )

   for __index, info in ipairs( self.classInfoList ) do
      for __index, ifType in ipairs( info:get_ifList() ) do
         self:getIfInfo( ifType )
      end
   end
   for __index, info in ipairs( self.classInfoList ) do
      local itabs = {}
      for __index, ifInfo in ipairs( self.ifInfoList ) do
         local implemented = false
         for __index, ifType in ipairs( info:get_ifList() ) do
            if ifType == ifInfo:get_typeInfo() then
               implemented = true
            end
         end
         if implemented then
            local funcs = {}
            for index, name in ipairs( ifInfo:get_methodNames() ) do
               local funcType = ifInfo:get_methodFuncTypes()[index]
               local ins = W.refNull( funcType )
               do
                  local slotIndex = info:get_name2slot()[name]
                  if slotIndex ~= nil then
                     local slot = info:get_slotList()[slotIndex]
                     if slot:get_impl() >= 0 then
                        if slot:get_funcType() ~= funcType then
                           Util.err( string.format( "%s: wasm: not support interface method type -- %s.%s", self.streamName, info:get_typeInfo():getTxt(  ), name) )
                        end
                        ins = self.mod:refFunc( slot:get_impl() )
                     end
                  end
               end
               table.insert( funcs, ins )
            end
            table.insert( itabs, W.structNew( ifInfo:get_itType(), funcs ) )
         else
          
            table.insert( itabs, W.refNull( W.htAny ) )
         end
      end
      local vals = {W.arrayNewFixed( self.rt:get_tAnyArr(), itabs )}
      for __index, slot in ipairs( info:get_slotList() ) do
         if slot:get_impl() >= 0 then
            table.insert( vals, self.mod:refFunc( slot:get_impl() ) )
         else
          
            table.insert( vals, W.refNull( slot:get_funcType() ) )
         end
      end
      table.insert( self.prologue, W.globalSet( info:get_vtGlobal(), W.structNew( info:get_vtType(), vals ) ) )
   end
end

function convFilter:collectDecl( node )

   local alreadySet = {}
   node:visit( function ( workNode, parent, relation, depth )
   
      do
         local classNode = _lune.__Cast( workNode, 3, Nodes.DeclClassNode )
         if classNode ~= nil then
            self.classType2decl[self:getGenSrc( classNode:get_expType() )] = classNode
         end
      end
      do
         local algeNode = _lune.__Cast( workNode, 3, Nodes.DeclAlgeNode )
         if algeNode ~= nil then
            self.algeType2decl[algeNode:get_algeType()] = algeNode
         end
      end
      do
         local enumNode = _lune.__Cast( workNode, 3, Nodes.DeclEnumNode )
         if enumNode ~= nil then
            self.enumType2decl[enumNode:get_enumType()] = enumNode
         end
      end
      do
         local _switchExp = workNode:get_kind()
         if _switchExp == Nodes.NodeKind.get_DeclMethod() or _switchExp == Nodes.NodeKind.get_ProtoMethod() or _switchExp == Nodes.NodeKind.get_DeclConstr() then
            self.declaredMethodSet[workNode:get_expType()]= true
         end
      end
      return Nodes.NodeVisitMode.Child
   end
   , 0, alreadySet )
end


function convFilter:processRoot( node, opt )

   self:collectDecl( node )
   for __index, child in ipairs( node:get_children() ) do
      do
         local declNode = _lune.__Cast( child, 3, Nodes.DeclVarNode )
         if declNode ~= nil then
            for __index, symbol in ipairs( declNode:get_symbolInfoList() ) do
               local orgSym = symbol:getOrg(  )
               if self:isGlobalSym( orgSym ) then
                  local rep = self:getRep( orgSym:get_typeInfo() )
                  self.globalSym2loc[orgSym] = _lune.newAlge( VarLoc.Global, {self.mod:addGlobal( rep, self:getDefaultVal( rep ) )})
               end
            end
         end
      end
   end
   local mainBody = self:lowerStmtList( node:get_children() )
   self:finalizeClasses(  )
   local body = {}
   append( body, self.prologue )
   append( body, mainBody )
   local mainFunc = self.mod:getFunc( self.mainIndex )
   mainFunc:setBody( body )
   self.mod:exportFunc( "main", self.mainIndex )
   self.rt:exportHostFuncs(  )
   self.stream:write( self.mod:encode(  ) )
end


local function createFilter( streamName, stream, ast )

   return convFilter._new(streamName, stream, ast)
end
_moduleObj.createFilter = createFilter


return _moduleObj
